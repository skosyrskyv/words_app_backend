# TEXT STYLE CONSTS
RED=[31m
GREEN=[32m
YELLOW=[33m
BLUE=[34m
MAGENTA=[35m
CYAN=[36m
WHITE=[37m
RESET=[0m
BOLD=[1m

# FUNCS
define infos
	@echo $(BLUE)$(1)$(RESET)
endef

define success
	@echo $(GREEN)$(1)$(RESET)
endef

define warning
	@echo $(YELLOW)$(1)$(RESET)
endef

define error
	@echo $(RED)$(1)$(RESET)
endef

define header
	@echo $(CYAN)=========================================$(RESET)
	@echo $(BOLD)$(CYAN)          $(1)$(RESET)
	@echo $(CYAN)=========================================$(RESET)
endef

# PATHS
BINARY_NAME=server.exe
MAIN_PATH=./cmd/api/
TMP_DIR=./tmp

.PHONY: build run clean

dev:
	$(call header, DEVELOPMENT SERVER)
	$(call infos,Starting with hot reload...)
	@CompileDaemon \
		-build="go build -o $(TMP_DIR)/$(BINARY_NAME) $(MAIN_PATH)" \
		-command="$(TMP_DIR)/$(BINARY_NAME)" \
		-log-prefix=false \
		-color=true

run:
	$(call header,       RUNNING)
	$(call infos, Running server...)
	@go run $(MAIN_PATH)/main.go

build:
	$(call header,      BUILDING)
	$(call infos,Building binary...)
	@go build -o ./bin/$(BINARY_NAME) $(MAIN_PATH)

clean:
	$(call header,      CLEANING)
	$(call warning, Removing temporary files...)
	@rm -rf $(TMP_DIR)
	@rm -rf ./bin
	@go clean
	$(call success, Clean complete!)
	
deps:
	@go mod tidy
	@go mod verify
