MAIN_FILES=cmd/*
OUT_DIR=out
BINARY_NAME=roadblock
INSTALL_DIR=/usr/local/bin

all: tidy lint build

release: tidy build

install:
	cp $(OUT_DIR)/$(BINARY_NAME) $(INSTALL_DIR)
	chmod 755 $(INSTALL_DIR)/$(BINARY_NAME)

lint:
	golangci-lint run $(MAIN_FILES)

tidy:
	go mod tidy

build:
	mkdir -p out
	go build -o $(OUT_DIR)/$(BINARY_NAME) $(MAIN_FILES)

uninstall:
	rm -f $(INSTALL_DIR)/$(BINARY_NAME)

clean:
	rm -rf $(OUT_DIR)
