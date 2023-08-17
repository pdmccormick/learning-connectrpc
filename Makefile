GO              ?= go
BIN             ?= $(abspath ./bin)
BUILD           ?= build
CMDS            := $(patsubst cmd/%/main.go,%,$(wildcard cmd/*/main.go))
BUILDS          = $(foreach cmd,$(CMDS),$(BUILD)/$(cmd))

GO_BUILD_FLAGS  ?= -trimpath -ldflags='-extldflags=-static'
GO_TAGS         ?= netgo

export PATH         := $(abspath $(BIN)):$(abspath $(BUILD)):$(PATH)
export GOBIN        := $(abspath $(BIN))
export CGO_ENABLED  := 0

ifdef V
	ifeq ("$(origin V)", "command line")
		VERBOSE=$(V)
	endif
endif

ifndef VERBOSE
	VERBOSE=0
endif

ifeq ($(VERBOSE),1)
	Q=
else
	Q=@
endif

.PHONY: all
all: $(BUILDS)

.PHONY: example
example: $(BUILD)/example

$(BUILD)/%: generate
	@echo "     BUILD       $@"
	$(Q)$(GO) build $(GO_BUILD_FLAGS) -o $@ -tags $(GO_TAGS),$* ./cmd/$*

$(BUILD)/%-amd64: generate cmd/%/*.go
	@echo "     BUILD       $@ [amd64]"
	$(Q)GOARCH=amd64 $(GO) build $(GO_BUILD_FLAGS) -o $@ -tags $(GO_TAGS),$* ./cmd/$*

$(BUILD)/%-arm64: generate cmd/%/*.go
	@echo "     BUILD       $@ [arm64]"
	$(Q)GOARCH=arm64 $(GO) build $(GO_BUILD_FLAGS) -o $@ -tags $(GO_TAGS),$* ./cmd/$*

.PHONY: generate
generate:
	@echo "     GENERATE"
	$(Q)$(GO) generate ./...

.PHONY: test
test:
	@echo "     TEST"
	$(Q)$(GO) test ./...

.PHONY: mod-tidy
mod-tidy:
	@echo "     GO MOD TIDY"
	$(Q)$(GO) mod tidy

.PHONY: mod-download
mod-download:
	@echo "     GO MOD DOWNLOAD"
	$(Q)$(GO) mod download

.PHONY: mod-vendor
mod-vendor:
	@echo "     GO MOD VENDOR"
	$(Q)$(GO) mod vendor

.PHONY: clean
clean:
	@echo "     CLEAN"
	$(Q)rm -rf $(BUILDS)

print-%: ; @echo $*=$($*)
