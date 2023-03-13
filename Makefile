build_and_run:
	go build && ./service-pattern-go
build_and_run_migrate:
	go build && ./service-pattern-go -migrateDB=true
build_and_run_background:
	go build && ./service-pattern-go &
run:
	go run .
unit-test:
	go test ./...
format:
	go fmt ./...
vet:
	go vet
tidy:
	go mod tidy
list_processes:
	ps
kill_process:
	kill 12345678
