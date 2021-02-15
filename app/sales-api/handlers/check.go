package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/YoungsoonLee/study-kube-with-ultimate-service/foundation/web"
)

type check struct {
	build string
	log   *log.Logger
}

func (c check) readiness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	status := struct {
		Status string
	}{
		Status: "OK",
	}

	return web.Respond(ctx, w, status, http.StatusOK)
}

// liveness returns simple status info if the service is alive.
// if the app is deployed to a kubernates cluster, it will also return pod, node, and
// namespace detiails via the Downward API. The kubernetes environment variables
// need to be set within your Pod/Deployment manifest.
func (c check) liveness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	return web.Respond(ctx, w, nil, http.StatusOK)
}
