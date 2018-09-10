ORG_PATH=github.com/rite2nikhil
PROJECT_NAME := kubernetes-node-scaler
REPO_PATH="$(ORG_PATH)/$(PROJECT_NAME)"
DEMO_BINARY_NAME := k8s-node-scaler-demo
DEMO_VERSION=1.0

VERSION_VAR := $(REPO_PATH)/version.Version
GIT_VAR := $(REPO_PATH)/version.GitCommit
BUILD_DATE_VAR := $(REPO_PATH)/version.BuildDate
BUILD_DATE := $$(date +%Y-%m-%d-%H:%M)
GIT_HASH := $$(git rev-parse --short HEAD)

ifeq ($(OS),Windows_NT)
	GO_BUILD_MODE = default
else
	UNAME_S := $(shell uname -s)
	ifeq ($(UNAME_S), Linux)
		GO_BUILD_MODE = pie
	endif
	ifeq ($(UNAME_S), Darwin)
		GO_BUILD_MODE = default
	endif
endif

GO_BUILD_OPTIONS := -buildmode=${GO_BUILD_MODE} -ldflags "-s -X $(VERSION_VAR)=$(DEMO_VERSION) -X $(GIT_VAR)=$(GIT_HASH) -X $(BUILD_DATE_VAR)=$(BUILD_DATE)"

# useful for other docker repos
REGISTRY ?= nikhilbh
DEMO_IMAGE_NAME := $(REGISTRY)/$(DEMO_BINARY_NAME)

clean-demo:
	rm -rf bin/$(PROJECT_NAME)/$(DEMO_BINARY_NAME)

clean:
	rm -rf bin/$(PROJECT_NAME)

build-demo:clean-demo
	go build -o bin/$(PROJECT_NAME)/$(DEMO_BINARY_NAME) $(GO_BUILD_OPTIONS) $(ORG_PATH)/$(PROJECT_NAME)/cmd

build:clean
	go build -o bin/$(PROJECT_NAME)/$(DEMO_BINARY_NAME) $(GO_BUILD_OPTIONS) $(ORG_PATH)/$(PROJECT_NAME)/cmd

.PHONY: build