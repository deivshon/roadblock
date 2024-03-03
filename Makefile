ENTRY_POINT=cmd/main.go
BINARY_NAME=roadblock

all:
	go mod tidy
	golangci-lint run $(ENTRY_POINT)
	go build -o $(BINARY_NAME) $(ENTRY_POINT)
