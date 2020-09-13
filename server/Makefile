GO_CMD=go
GO_BUILD=$(GO_CMD) build
GO_CLEAN=$(GO_CMD) clean
GO_TEST=$(GO_CMD) test
GO_MOD=$(GO_CMD) mod

PKG_NAME=user_management.be_webapi

CMD_DIR=cmd
MAIN_FILE=main.go
MAIN_PATH=$(CMD_DIR)/$(MAIN_FILE)

BIN_DIR=bin
BINARY_NAME=build
BUILD_PATH=$(BIN_DIR)/$(BINARY_NAME)

MOD_PATH=go.mod
SUM_PATH=go.sum

.PHONY: build
build: clean
	$(GO_MOD) init $(PKG_NAME)
	$(GO_BUILD) -o $(BUILD_PATH) -v $(MAIN_PATH)

.PHONY: clean
clean: 
	rm -f $(MOD_PATH)
	rm -f $(SUM_PATH)
	rm -f $(BUILD_PATH)
	
.PHONY: run
run: build
	$(BUILD_PATH)

.PHONY: dev
dev: 
	nodemon --exec go run $(MAIN_PATH) --signal SIGTERM
