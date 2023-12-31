.PHONY: server client create-network docker-server docker-client test bench lint

server:
	go run cmd/server/main.go

client:
	go run cmd/client/main.go

create-network:
	-docker network create quotes

docker-server: create-network
	-docker rm -f quotes_server
	docker build -t quotes_server -f deployments/Dockerfile.server .
	docker run --rm --network=quotes --name=quotes_server -e LISTEN=:8099 -p 8099:8099 quotes_server

docker-client: create-network
	docker build -t quotes_client -f deployments/Dockerfile.client .
	docker run --rm --network=quotes -e ADDRESS=quotes_server:8099 quotes_client

test: 
	go test ./...

bench:
	go test -bench=. -count 4

lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.54.2
	$(shell go env GOPATH)/bin/golangci-lint run -v