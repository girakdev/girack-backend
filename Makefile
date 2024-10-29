serve:
	docker compose exec -it girack go run ./cmd/girack/main.go

dup:
	docker compose up -d

ddown:
	docker compose down

db:
	docker compose exec -it db bash -c 'psql -h db -p 5432 -U postgres -d girack'

migration:
	go generate ./ent
	docker compose exec -it girack go run ./cmd/migration/main.go

mock-gen:
	go generate ./...

swag-gen:
	swag init --dir ./cmd/girack,./controller,./application/model -o ./docs/

test:
	go test -v ./... -p 4