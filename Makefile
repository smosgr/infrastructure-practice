
.PHONY: docker clean

bin/web-server: $(shell find . -iname '*.go')
	go build -o bin/web-server main.go

bin/web-server-dist: $(shell find . -iname '*.go')
	GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -o bin/web-server-dist main.go

docker: Dockerfile bin/web-server-dist
	@eval $$(minikube docker-env); \
	docker image build -t greeter:latest .

clean:
	rm -rf bin
