setup-db:
	sqlite3 /var/tmp/tennis.db < setup.sql
build_and_run:
	go1.20.1 build && ./service-pattern-go
build_and_run_background:
	go1.20.1 build && ./service-pattern-go &
run:
	go1.20.1 run .
unit-test:
	go1.20.1 test ./...
format:
	go1.20.1 fmt ./...
vet:
	go1.20.1 vet
tidy:
	go1.20.1 mod tidy
