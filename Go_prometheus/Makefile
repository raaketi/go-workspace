## LOCAL -----------------------------------------------------------------------
 
.PHONY: run
run:
	go build -race -ldflags "-s -w" -o bin/main main.go
	bin/main
 
.PHONY: test
test:
	curl --request GET http://localhost:3000/api/v1/leagues
 
## DOCKER ----------------------------------------------------------------------
 
.PHONY: docker-push
docker-push:
	docker build -t adi463/football:latest -f ./docker/Dockerfile .
	docker push adi463/football:latest
	docker rmi adi463/football:latest
	docker system prune --volumes --force
 
## KUBE ------------------------------------------------------------------------
 
.PHONY: kube-deploy
kube-deploy:
	kubectl apply -f deploy/k8s/deployment.yaml
	kubectl apply -f deploy/k8s/service.yaml
 
.PHONY: kube-test
kube-test:
	curl --request GET http://localhost:3000/api/v1/leagues

.PHONY: deploy-roles
deploy-roles:
	kubectl apply -f deploy/k8s/monitor/cluster-role.yaml
	kubectl apply -f deploy/k8s/monitor/service-account.yaml
	kubectl apply -f deploy/k8s/monitor/cluster-role-binding.yaml
 
.PHONY: deploy-prometheus
deploy-prometheus:
	kubectl apply -f deploy/k8s/prometheus/config-map.yaml
	kubectl apply -f deploy/k8s/prometheus/deployment.yaml
	kubectl apply -f deploy/k8s/prometheus/service.yaml