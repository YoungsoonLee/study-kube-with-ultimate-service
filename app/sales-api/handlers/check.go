package handlers

import (
	"context"
	"log"
	"net/http"
	"os"

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
	host, err := os.Hostname()
	if err != nil {
		host = "unavailable"
	}

	info := struct {
		Status    string `json:"status,omitempty"`
		Build     string `json:"build,ommitempty"`
		Host      string `json:"host,omitemmpty"`
		Pod       string `json:"pod,omitempty"`
		PodIP     string `json:"podIP,omitempty"`
		Node      string `json:"node,omitempty"`
		Namespace string `json:"namespace,omitempty"`
	}{
		Status:    "up",
		Build:     c.build,
		Host:      host,
		Pod:       os.Getenv("KUBERNATES_PODNAME"),
		PodIP:     os.Getenv("KUBERNATES_NAMESPACE_POD_IP"),
		Node:      os.Getenv("KUBERBATES_NODENAME"),
		Namespace: os.Getenv("KUBERNATES_NAMESPACE"),
	}

	return web.Respond(ctx, w, info, http.StatusOK)
}
