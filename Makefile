###############################################################################
###                                Build Control                            ###
###############################################################################

# entry point
MAIN_ENTRY=cmd/app/main.go

# project info
NAME=NEU_IPGW
REPO=github.com/DoraTiger/NEU_IPGW

# build version
ifneq ($(shell git symbolic-ref -q --short HEAD),)
BUILD_VERSION := unreleased-$(shell git symbolic-ref -q --short HEAD)-$(shell git rev-parse HEAD)
else
BUILD_VERSION := $(shell git describe --tags --always)
endif

LD_FLAGS += -X '$(REPO)/version.BuildVersion=${BUILD_VERSION}'

# build time
BUILD_TIME=$(shell date +%FT%T%z)
LD_FLAGS += -X "${REPO}/version.BuildTime=${BUILD_TIME}"

# build repo
BuildRepo=$(shell git config --get remote.origin.url)
LD_FLAGS += -X "${REPO}/version.BuildRepo=${BuildRepo}"

# additional flags
ifeq (vendor,$(findstring vendor,$(BUILD_OPTIONS)))
  BUILD_FLAGS += -mod=vendor
else
  BUILD_FLAGS += -mod=readonly
endif

ifeq (,$(findstring nostrip,$(BUILD_OPTIONS)))
  BUILD_FLAGS += -trimpath
  LD_FLAGS += -s -w
endif

ifeq (race,$(findstring race,$(BUILD_OPTIONS)))
  CGO_ENABLED=1
  BUILD_FLAGS += -race
endif



# disable cgo by default, make sure the binary is statically linked
CGO_ENABLED ?= 0

# output directories
BUILD_DIR=build
RELEASE_DIR=release

ifeq ($(OS),Windows_NT)
  BINARY_EXT := .exe
else
  BINARY_EXT :=
endif

###############################################################################
###                                Build                                    ###
###############################################################################

GO_BUILD = CGO_ENABLED=$(CGO_ENABLED) go build $(BUILD_FLAGS) -ldflags '$(LD_FLAGS)'

build:
	$(GO_BUILD) -o $(BUILD_DIR)/$(NAME)$(BINARY_EXT) ${MAIN_ENTRY}
.PHONY: build

release: clean build-all
	@bash scripts/release.sh $(NAME) $(BUILD_DIR) $(RELEASE_DIR)
.PHONY: release

clean:
	rm -rf $(BUILD_DIR)/*
	rm -rf $(RELEASE_DIR)/*
.PHONY: clean

# Build for a specific platform
define build-platform
GOOS=$(1) GOARCH=$(2) $(GO_BUILD) -o $(BUILD_DIR)/$(1)-$(2)/$(NAME)$(if $(findstring windows,$(1)),.exe) ${MAIN_ENTRY}
endef

# define a rule for each platform
define build-platform-rule
$(1): ; $$(call build-platform,$$(word 1,$$(subst -, ,$(1))),$$(word 2,$$(subst -, ,$(1))))
.PHONY: $(1)
endef

PLATFORM_LIST = \
	darwin-amd64 \
	darwin-arm64 \
	linux-386 \
	linux-amd64 \
	linux-arm \
	linux-mips \
	linux-mipsle \
	linux-mips64 \
	linux-mips64le \
	freebsd-386 \
	freebsd-amd64 \
	windows-386 \
	windows-amd64 \
	windows-arm

# auto-generate rules for each platform
$(foreach p,$(PLATFORM_LIST),$(eval $(call build-platform-rule,$(p))))

build-all: $(PLATFORM_LIST)
.PHONY: build-all
