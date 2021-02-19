package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/YoungsoonLee/study-kube-with-ultimate-service/foundation/database"
	"github.com/YoungsoonLee/study-kube-with-ultimate-service/foundation/web"
	"github.com/jmoiron/sqlx"
)

type checkGroup struct {
	build string
	log   *log.Logger
	db    *sqlx.DB
}

func (cg checkGroup) readiness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	status := "ok"
	statusCode := http.StatusOK
	if err := database.StatusCheck(ctx, cg.db); err != nil {
		status = "db not ready"
		statusCode = http.StatusInternalServerError
	}

	health := struct {
		Status string `json:"status"`
	}{
		Status: status,
	}

	return web.Respond(ctx, w, health, statusCode)
}

// liveness returns simple status info if the service is alive.
// if the app is deployed to a kubernates cluster, it will also return pod, node, and
// namespace detiails via the Downward API. The kubernetes environment variables
// need to be set within your Pod/Deployment manifest.
func (cg checkGroup) liveness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	return web.Respond(ctx, w, nil, http.StatusOK)
}
