package web

import (
	"context"
	"net/http"
	"os"
	"syscall"
	"time"

	"github.com/dimfeld/httptreemux/v5"
	"github.com/google/uuid"
)

type ctxKey int

const KeyValues ctxKey = 1

type Values struct {
	TraceID    string
	Now        time.Time
	StatusCode int
}

// handler ...
type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

// App ...
type App struct {
	*httptreemux.ContextMux
	shutdown chan os.Signal
	mw       []Middleware
}

// NewApp ...
func NewApp(shutdown chan os.Signal, mw ...Middleware) *App {
	app := App{
		ContextMux: httptreemux.NewContextMux(),
		shutdown:   shutdown,
		mw:         mw,
	}

	return &app
}

// Handle ...
func (a *App) Handle(method string, path string, handler Handler, mw ...Middleware) {
	// first wrap handler specific middleware around this handler.
	handler = wrapMiddleware(mw, handler)

	// add the application's ggeneral middleware to the handler chain.
	handler = wrapMiddleware(a.mw, handler)

	h := func(w http.ResponseWriter, r *http.Request) {
		v := Values{
			TraceID: uuid.New().String(),
			Now:     time.Now(),
		}
		ctx := context.WithValue(r.Context(), KeyValues, &v)

		if err := handler(ctx, w, r); err != nil {
			// Handle error
			a.SignalShutdown()
			return
		}

		// boilerplate
	}

	a.ContextMux.Handle(method, path, h)
}

func (a *App) SignalShutdown() {
	a.shutdown <- syscall.SIGTERM
}
