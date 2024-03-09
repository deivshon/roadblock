MAIN_FILES=cmd/*
OUT_DIR=out
BINARY_NAME=roadblock

all:
	mkdir -p out

	go mod tidy
	golangci-lint run $(MAIN_FILES)
	go build -o $(OUT_DIR)/$(BINARY_NAME) $(MAIN_FILES)

clean:
	rm -rf $(OUT_DIR)
