build: bin/doing
	go build -o bin/doing github.com/cuichenli/doing/cmd/doing

test:
	go test ./model