proto: 
	protoc **/pb/*.proto --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative

shell:
	bash script.sh

postgres:
	docker run --name grpc-commerce -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine

server:
	go run main.go