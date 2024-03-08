run:
	go run ./cmd/
migrate:
	psql -U postgres -c "CREATE DATABASE mobydev;"
	