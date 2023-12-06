serve:
	docker compose exec -it girack go run ./cmd/girack/main.go

dup:
	docker compose up -d

ddown:
	docker compose down

rebuild:
	docker compose rm db && docker compose rm girack && docker compose build --no-cache && docker compose up -d
	 
db:
	docker compose exec -it db sh