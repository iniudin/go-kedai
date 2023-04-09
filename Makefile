.PHONY: run tidy compose-up compose-down db-up db-down sql-gen

run:
	air

tidy:
	go mod tidy

compose-up:
	docker compose up -d

compose-down:
	docker compose down
db-up:
	migrate -path database/migrations -database "postgresql://katabe:PalingKatabe@localhost:5432/katabe?sslmode=disable" -verbose up
db-down:
	migrate -path database/migrations -database "postgresql://katabe:PalingKatabe@localhost:5432/katabe?sslmode=disable" -verbose down
sql-gen:
	sqlc generate