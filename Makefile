compose-up:
	docker-compose up -d

schema-up:
	go run ./cmd/dbschema/main.go -objects up

seed-up:
	go run ./cmd/dbschema/main.go -seed up

seed-down:
	go run ./cmd/dbschema/main.go -seed down
	

dbfull = schema-up seed-down seed-up

db-up: $(dbfull) 

api-serve:
	go run ./cmd/server/main.go

api-test:
	go test ./test/apitests/...