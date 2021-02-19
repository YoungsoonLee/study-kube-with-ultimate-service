package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/YoungsoonLee/study-kube-with-ultimate-service/business/mid"
	"github.com/YoungsoonLee/study-kube-with-ultimate-service/foundation/web"
	"github.com/jmoiron/sqlx"
)

func API(build string, shutdown chan os.Signal, log *log.Logger, db *sqlx.DB) *web.App {
	app := web.NewApp(shutdown, mid.Logger(log), mid.Errors(log), mid.Metrics(), mid.Panics(log))

	cg := checkGroup{
		build: build,
		log:   log,
		db:    db,
	}
	app.Handle(http.MethodGet, "/readiness", cg.readiness)
	app.Handle(http.MethodGet, "/liveness", cg.liveness)

	return app
}
