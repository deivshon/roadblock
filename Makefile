ENTRY_POINT=cmd/main.go
OUT_DIR=out
BINARY_NAME=roadblock

all:
	mkdir -p out

	go mod tidy
	golangci-lint run $(ENTRY_POINT)
	go build -o $(OUT_DIR)/$(BINARY_NAME) $(ENTRY_POINT)

clean:
	rm -rf $(OUT_DIR)
