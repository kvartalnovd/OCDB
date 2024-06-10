TARGET=ocdb
PROJECT_NAME=OCDB
LOCAL_CONFIG=config/local.yml

START_TIME="`date +'%y.%m.%d %H:%M:%S'`"

.PHONY: tag
tag:
	./scripts/tag.sh

.PHONY: build
build:
	@go build -o bin/${TARGET} cmd/${TARGET}/main.go 

.PHONY: run
run: build
	@./bin/${TARGET} --config "${LOCAL_CONFIG}"

xrun:
	@echo "[$(PROJECT_NAME)][$(START_TIME)] Starting Contracted DataBase"
	@make run
	@echo "[$(PROJECT_NAME)]["`date +'%y.%m.%d %H:%M:%S'`"] Close the OCDB"

.PHONY: test
test:
	@go test ./...
