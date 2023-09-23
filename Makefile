server:
	go run cmd/server/main.go

client:
	go run cmd/client/main.go

create-network:
	-docker network create qoutes

docker-server: create-network
	-docker rm -f qoutes_server
	docker build -t qoutes_server -f deployments/Dockerfile.server .
	docker run --rm --network=qoutes --name=qoutes_server -e LISTEN=:8099 -p 8099:8099 qoutes_server

docker-client: create-network
	docker build -t quotes_client -f deployments/Dockerfile.client .
	docker run --rm --network=qoutes -e ADDRESS=qoutes_server:8099 quotes_client
