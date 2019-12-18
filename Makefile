SHELL := /bin/bash

deploy:
# 	./build/distribute-image.sh go-chassis/pilotv2server:latest
	- kubectl --kubeconfig $(KUBECONFIG_TEST) delete -f ./deploy/k8s/
	kubectl --kubeconfig $(KUBECONFIG_TEST) apply -f ./deploy/k8s/

binary:
	- rm ./client
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags '-w' -o ./client .

docker: binary
	docker build -t swr.cn-east-2.myhuaweicloud.com/yb7/istio-demo-client:rest-v1 .