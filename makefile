SHELL := /bin/bash

# build containers
all: sales-api

sales-api:
	docker build \
		-f zarf/docker/dockerfile.sales-api \
		-t sales-api-amd64:1.0 \
		--build-arg VCS_REF=`git rev-parse HEAD` \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
		.
# running from within k8s/dev
kind-up:
	kind create cluster --image kindest/node:v1.19.1 --name ardan-starter-cluster --config zarf/k8s/dev/kind-config.yaml

kind-load:
	kind load docker-image sales-api-amd64:1.0 --name ardan-starter-cluster
	# kind load docker-image metrics-amd64:1.0 --name ardan-starter-cluster

kind-service:
	kustomize build zarf/k8s/dev | kubectl apply -f -

kind-status:
	kubectl get nodes
	kubectl get pods --watch

kind-status-full:
	kubectl describe pod -lapp=sales-api

kind-logs:
	kubectl logs -lapp=sales-api --all-containers=true -f

kind-down:
	kind delete cluster --name ardan-starter-cluster

kind-sales-api: sales-api
	kind load docker-image sales-api-amd64:1.0 --name ardan-starter-cluster
	kubectl delete pods -lapp=sales-api

run:
	go run app/sales-api/main.go

runa:
	go run app/admin/main.go

tidy:
	go mod tidy
	go mod vendor

test:
	go test -v ./... -count=1
	staticcheck ./...