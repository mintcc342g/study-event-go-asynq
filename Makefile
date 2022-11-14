PACKAGE = study-event-go-asynq
BUILDPATH ?= $(CURDIR)
BASE	= $(BUILDPATH)
BIN		= $(BASE)/bin
GORACE	= -race

ifeq ($(OS),Windows_NT)
	PACKAGE = study-event-go-asynq.exe
	ifeq ($(PROCESSOR_ARCHITECTURE), x86)
		GORACE =
	endif
else
	UNAME := $(shell uname)
	ifeq ($(UNAME), Linux)
		GOENV ?= CGO_ENABLED=0 GOOS=linux
	endif
endif


GOBUILD	= ${GOENV} go
GO		= go

BUILDTAG=-tags 'studyEventGoAsynq'
export GO111MODULE=on

V = 0
Q = $(if $(filter 1,$V),,@)
M = $(shell printf "\033[34;1m▶\033[0m")

.PHONY: all
all: all_info build ;

ifeq ($(OS),Windows_NT)
all_info: ; @echo building all steps...
else
all_info: ; $(info $(M) build all steps…) @
endif

.PHONY: build
build: vendor ;
ifeq ($(OS),Windows_NT)
	@echo building executable...
else
	$(info $(M) building executable…) @
endif
	$Q cd $(BASE)/cmd && $(GOBUILD) build \
		$(BUILDTAG) \
		-o $(BIN)/$(PACKAGE)


# Test

.PHONY: test
test: mock ;
ifeq ($(OS),Windows_NT)
	@echo run test...
else
	$(info $(M) run test…) @
endif
	$Q cd $(BASE) && $(GO) test -p 1 -failfast $(GORACE) -cover ./...
