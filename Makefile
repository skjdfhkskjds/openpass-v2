#!/usr/bin/make -f

# Specify the default target if none is provided
.DEFAULT_GOAL := build

###############################################################################
###                                  Build                                  ###
###############################################################################

# Build the project
build:
	@echo "Building the project..."
	@go build -o bin/ ./...

clean:
	@rm -rf .tmp/


###############################################################################
###                                 Testing                                 ###
###############################################################################

test-unit:
	@$(MAKE)
	@echo "Running unit tests..."
	go test ./...


###############################################################################
###                                 Linting                                 ###
###############################################################################

format:
	@$(MAKE) license-fix golangci-fix

lint:
	@$(MAKE) license golangci


MODULES := $(shell find . -type f -name 'go.mod' -exec dirname {} \;)
# Exclude root module
MODULES := $(filter-out ./,$(MODULES))

#################
# golangci-lint #
#################

golangci-install:
	@echo "--> Installing golangci-lint"
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint

golangci:
	@$(MAKE) golangci-install
	@echo "--> Running linter"
	@go list -f '{{.Dir}}/...' -m | grep -v '**/contracts' | xargs golangci-lint run  --timeout=10m --concurrency 8 -v 

golangci-fix:
	@$(MAKE) golangci-install
	@echo "--> Running linter"
	@go list -f '{{.Dir}}/...' -m | grep -v '**/contracts' | xargs golangci-lint run  --timeout=10m --fix --concurrency 8 -v 


#################
#    license    #
#################

license-install:
	@echo "--> Installing google/addlicense"
	@go install github.com/google/addlicense

license:
	@$(MAKE) license-install
	@echo "--> Running addlicense with -check"
	@for module in $(MODULES); do \
		(cd $$module && addlicense -check -v -f ./LICENSE.header ./.) || exit 1; \
	done

license-fix:
	@$(MAKE) license-install
	@echo "--> Running addlicense"
	@for module in $(MODULES); do \
		(cd $$module && addlicense -v -f ./LICENSE.header ./.) || exit 1; \
	done

#################
#     gosec     #
#################

gosec-install:
	@echo "--> Installing gosec"
	@go install github.com/cosmos/gosec/v2/cmd/gosec 

gosec:
	@$(MAKE) gosec-install
	@echo "--> Running gosec"
	@gosec -exclude G702 ./...


###############################################################################
###                                 CodeGen                                 ###
###############################################################################

generate:
	@$(MAKE) proto

#################
#     proto     #
#################

modulesProtoDir := "proto"

proto:
	@$(MAKE) buf-lint-fix buf-lint proto-build

proto-build:
	@echo "--> Generating proto files"
	@protoc --go_out=$(modulesProtoDir)/out --go_opt=paths=source_relative ./$(modulesProtoDir)/*.proto

proto-clean:
	@find . -name '*.pb.go' -delete
	@find . -name '*.pb.gw.go' -delete

buf-lint-fix:
	@$(MAKE) buf-install 
	@echo "--> Running buf format"
	@buf format -w --error-format=json $(modulesProtoDir)

buf-lint:
	@$(MAKE) buf-install 
	@echo "--> Running buf lint"
	@buf lint --error-format=json $(modulesProtoDir)

###############################################################################
###                                 Golang                                  ###
###############################################################################

tidy: |
	go mod tidy