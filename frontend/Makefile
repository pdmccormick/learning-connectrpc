NPM         ?= npm
SVELTEKIT   ?= $(NPM) exec -- svelte-kit
VITE        ?= $(NPM) exec -- vite

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
all: build

.PHONY: svelte-kit-sync
svelte-kit-sync:
	@echo "     SVELTE-KIT SYNC"
	$(Q)$(SVELTEKIT) sync

.PHONY: dev
dev: svelte-kit-sync
	@echo "     VITE DEV"
	$(Q)$(VITE) dev

.PHONY: build
build: svelte-kit-sync
	@echo "     VITE BUILD"
	$(Q)$(VITE) build

.PHONY: build-watch
build-watch: svelte-kit-sync
	@echo "     VITE BUILD-WATCH"
	$(Q)$(SVELTEKIT) sync
	$(Q)$(VITE) build --watch

.PHONY: preview
preview: svelte-kit-sync
	@echo "     VITE PREVIEW"
	$(Q)$(VITE) preview

.PHONY: test
test:
	@echo "     NPM RUN TEST"
	$(Q)$(NPM) run test

.PHONY: check
check:
	@echo "     NPM RUN CHECK"
	$(Q)$(NPM) run check

.PHONY: lint
lint:
	@echo "     NPM RUN LINT"
	$(Q)$(NPM) run lint

.PHONY: format
format:
	@echo "     NPM RUN FORMAT"
	$(Q)$(NPM) run format

.PHONY: generate-buf
generate-buf:
	@echo "     BUF GENERATE"
	$(Q)$(NPM) run buf:generate

.PHONY: clean
clean:
	@echo "     CLEAN"
	$(Q)rm -rf ./build/*

print-%: ; @echo $*=$($*)
