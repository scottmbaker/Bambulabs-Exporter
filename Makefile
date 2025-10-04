REPOSITORY="smbaker/bambulabs-exporter"
VERSION="0.2.0"

build:
	go build -o bambulabs-exporter

docker-build:
	docker build -t $(REPOSITORY):$(VERSION) .

docker-push:
	docker push $(REPOSITORY):$(VERSION)

tidy:
	go fmt main.go
	go fmt schema.go
	go mod tidy
