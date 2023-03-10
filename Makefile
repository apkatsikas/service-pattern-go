setup-db:
	sqlite3 /var/tmp/tennis.db < setup.sql
# TODO - build the binary and run that instead
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
