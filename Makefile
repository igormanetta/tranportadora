text=

run:
	go run ./cmd/transportadora/main.go	

test:
	go test ./...

cover:
	go test ./... -cover

sqlcgen:
	cd infra/pgstore && sqlc generate

migratecreate:
	migrate create -ext sql -dir infra/pgstore/migrations -seq $(text) 	

build: 
	docker build -t transportadora .
