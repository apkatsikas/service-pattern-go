build_and_run:
	go build && ./service-pattern-go
build_and_run_migrate:
	go build && ./service-pattern-go -migrateDB=true
build_and_run_background:
	go build && ./service-pattern-go &
build_and_run_docker:
	docker run -p 8080:8080 --volume $$PWD:/app --rm -it $$(docker build -q .)
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
find_go_symlink:
	readlink -f /usr/bin/go
