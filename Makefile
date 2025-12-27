APP=cldo
BIN_DIR=bin

build:
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(APP) ./cmd/$(APP)

run:
	go run ./cmd/$(APP)

install:
	cp $(BIN_DIR)/$(APP) /usr/local/bin/$(APP)

clean:
	rm -rf $(BIN_DIR)
