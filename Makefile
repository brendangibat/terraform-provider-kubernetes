# All directories with go source
SRCDIRS=$(shell ./scripts/find-project-files *.go)
TESTDIRS=$(shell ./scripts/find-project-files *_test.go)

GO15VENDOREXPERIMENT=$(shell echo 1)

GIT_SHA= $(shell git rev-parse --short HEAD)

ifeq ($(shell uname -s), Linux)
# Do a build without dependencies for docker image.  Much slower.
GO_BUILD = CGO_ENABLED=0 go build -a -installsuffix cgo
else
GO_BUILD = go build -a
endif

.PHONY: all
all : show test build

.PHONY: restore-deps
restore-deps :
	command -v glide >/dev/null 2>&1 || { echo >&2 "Error: glide (https://github.com/Masterminds/glide) is not installed.  Please install.  Aborting."; exit 1; }
	rm -rf vendor/
	glide up

.PHONY: build
build :
	$(GO_BUILD) -o bin/terraform-provider-kubernetes .

.PHONY: install
install : restore-deps build
	cp terraform-provider-kubernetes $GOPATH/bin/terraform-provider-kubernetes

.PHONY: test
test : prepare-test
	go test -v $(TESTDIRS) | tee go.out
	if which go2xunit >/dev/null; \
	then mkdir -p /report; go2xunit -fail -input go.out -output /report/test_results.xml; \
	fi

# Use the 'go vet' tool to examine Go source code and report suspicious constructs for each package in project
.PHONY: vet
vet :
	find . ! -path "./vendor/*" -a -name "*.go" -print0 | xargs -0 -n1 go vet

.PHONY: gofmt
gofmt:
	find . ! -path "./vendor/*" -a -name "*.go" -print0 | xargs -0 -n1 gofmt -w

.PHONY: show
show :
	@echo "========================================"
	@echo "GOPATH=${GOPATH}"
	@echo "GOBIN=${GOBIN}"
	@echo "SRCDIRS=${SRCDIRS}"
	@echo "TESTDIRS=${TESTDIRS}"
	@go version
	@echo "========================================"

# List all available argets/tasks
.PHONY: no_targets__ list
no_targets__:
list:
	sh -c "$(MAKE) -p no_targets__ | awk -F':' '/^[a-zA-Z0-9][^\$$#\/\\t=]*:([^=]|$$)/ {split(\$$1,A,/ /);for(i in A)print A[i]}' | grep -v '__\$$' | sort"

.PHONY: clean
clean :
	rm -f *.out
	rm -rf bin
	find . -iname *.test | xargs -I LALA rm LALA
