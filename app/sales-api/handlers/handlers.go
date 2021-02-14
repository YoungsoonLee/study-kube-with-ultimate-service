package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/YoungsoonLee/study-kube-with-ultimate-service/business/mid"
	"github.com/YoungsoonLee/study-kube-with-ultimate-service/foundation/web"
)

func API(build string, shutdown chan os.Signal, log *log.Logger) *web.App {
	app := web.NewApp(shutdown, mid.Logger(log), mid.Errors(log), mid.Metrics(), mid.Panics(log))

	check := check{
		build: build,
		log:   log,
	}
	app.Handle(http.MethodGet, "/readiness", check.readiness)
	app.Handle(http.MethodGet, "/liveness", check.liveness)

	return app
}
