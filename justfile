build:
    go build -o uts

run:
    go run .

install:
    go install .

tidy:
    go mod tidy

test:
    go test -v ./...