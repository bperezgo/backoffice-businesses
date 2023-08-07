.PHONY:

build:
	go build -o main ./cmd/main.go

docker-build:
	docker build -t backoffice-business-app:v1 -f docker/Dockerfile .