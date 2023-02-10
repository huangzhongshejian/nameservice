LEDGER_ENABLED ?= true

build_tags =
ifeq ($(LEDGER_ENABLED),true)
  ifeq ($(OS),Windows_NT)
    GCCEXE = $(shell where gcc.exe 2> NUL)
    ifeq ($(GCCEXE),)
      $(error gcc.exe not installed for ledger support, please install or set LEDGER_ENABLED=false)
    else
      build_tags += ledger
    endif
  else
    UNAME_S = $(shell uname -s)
    ifeq ($(UNAME_S),OpenBSD)
      $(warning OpenBSD detected, disabling ledger support (https://github.com/cosmos/cosmos-sdk/issues/1988))
    else
      GCC = $(shell command -v gcc 2> /dev/null)
      ifeq ($(GCC),)
        $(error gcc not installed for ledger support, please install or set LEDGER_ENABLED=false)
      else
        build_tags += ledger
      endif
    endif
  endif
endif

PACKAGES=$(shell go list ./... | grep -v '/simulation')

VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=NameService \
	-X github.com/cosmos/cosmos-sdk/version.ServerName=nsd \
	-X github.com/cosmos/cosmos-sdk/version.ClientName=nscli \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) 

BUILD_FLAGS := -ldflags '$(ldflags)'

include Makefile.ledger
all: install

install: go.sum
		@echo "--> Installing nsd & nscli"
		@go install -mod=readonly $(BUILD_FLAGS) ./cmd/nsd
		@go install -mod=readonly $(BUILD_FLAGS) ./cmd/nscli

go.sum: go.mod
		@echo "--> Ensure dependencies have not been modified"
		GO111MODULE=on go mod verify

test:
	@go test -mod=readonly $(PACKAGES)

# PACKAGES=$(shell go list ./... | grep -v '/simulation')

# VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
# COMMIT := $(shell git log -1 --format='%H')

# # TODO: Update the ldflags with the app, client & server names
# ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=NewApp \
# 	-X github.com/cosmos/cosmos-sdk/version.ServerName=appd \
# 	-X github.com/cosmos/cosmos-sdk/version.ClientName=appcli \
# 	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
# 	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) 

# BUILD_FLAGS := -ldflags '$(ldflags)'

# all: install

# install: go.sum
# 		go install -mod=readonly $(BUILD_FLAGS) ./cmd/appd
# 		go install -mod=readonly $(BUILD_FLAGS) ./cmd/appcli

# go.sum: go.mod
# 		@echo "--> Ensure dependencies have not been modified"
# 		GO111MODULE=on go mod verify

# # Uncomment when you have some tests
# # test:
# # 	@go test -mod=readonly $(PACKAGES)

# # look into .golangci.yml for enabling / disabling linters
# lint:
# 	@echo "--> Running linter"
# 	@golangci-lint run
# 	@go mod verify
