SERVER_OUT := "bin/server"
API_OUT := "api/api.pb.go"
PKG := "github.com/adamo57/contentAPI"
SERVER_PKG_BUILD := "${PKG}/server"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)

.PHONY: api build_server 

api/api.pb.go: api/api.proto
	@protoc -I api/ \
		--go_out=plugins=grpc:api \
		api/api.proto


api: api/api.pb.go ## Auto-generate grpc go sources

dep: ## Get the dependencies
	@go get -v -d ./...

build_server: dep api ## Build the binary file for server
	@go build -i -v -o $(SERVER_OUT) $(SERVER_PKG_BUILD)

clean: ## Remove previous builds
	@rm $(SERVER_OUT) $(API_OUT)

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
