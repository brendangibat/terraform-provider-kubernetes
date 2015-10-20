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

# Initialize container workspace, test, build
.PHONY: container-init
container-init: test build

# Build and run a docker container (dev)
.PHONY: container-run
container-run : container-build
	docker run -p 8080:8080 -e Udh_GitSha=$(GIT_SHA) -it docker.moveaws.com/udh:dev-latest

# Build and run a docker container if using docker-machine (dev)
.PHONY: container-run-machine
container-run-machine : container-build
	docker run -p 8080:8080 -e Udh_Swagger_Host=`docker-machine ip default` -e Udh_GitSha=$(GIT_SHA) -it docker.moveaws.com/udh:dev-latest

# Build and run a docker container if using docker-machine (dev)
.PHONY: container-run-machine-trace
container-run-machine-trace : container-build
	docker run -p 8080:8080 -e Udh_Trace=true -e Udh_Swagger_Host=`docker-machine ip default` -e Udh_GitSha=$(GIT_SHA) -it docker.moveaws.com/udh:dev-latest

# The make command on jenkins executes things differently,
# building the image twice is the only way I've found to be consistent for
# local pushing and on jenkins. This could be easily done without a rebuild
# if I could get variable substitution from shell outs working correctly on
# jenkins for its make version
.PHONY: build-push-dev-images
build-push-dev-images :
	docker build --tag=docker.moveaws.com/udh:dev-latest .
	docker build --tag=docker.moveaws.com/udh:$(GIT_SHA) .
	docker push docker.moveaws.com/udh:dev-latest
	docker push docker.moveaws.com/udh:$(GIT_SHA)

.PHONY: build
build :
	$(GO_BUILD) -o bin/server .

.PHONY: prepare-test
prepare-test :
	cp udh.yaml /tmp/udh-test.yaml

.PHONY: test
test : prepare-test
	go test -v $(TESTDIRS) | tee go.out
	if which go2xunit >/dev/null; \
	then mkdir -p /report; go2xunit -fail -input go.out -output /report/test_results.xml; \
	fi

.PHONY: container-build
container-build :
	docker build --tag=docker.moveaws.com/udh:dev-latest .

.PHONY: container-test
container-test :
	docker ps -q -f name=udh-test | xargs -I LALA docker stop LALA
	docker ps -q -f status=exited -f name=udh-test | xargs -I LALA docker rm LALA
	docker build --tag=docker.moveaws.com/udh:test .
	docker run -i --name udh-test docker.moveaws.com/udh:test make test coverage
	docker stop udh-test
	docker rm udh-test

# Use the 'go vet' tool to examine Go source code and report suspicious constructs for each package in project
.PHONY: vet
vet :
	find . ! -path "./vendor/*" -a -name "*.go" -print0 | xargs -0 -n1 go vet

# Determine the test coverage for each package in the project
.PHONY: coverage
coverage : prepare-test
	scripts/coverage

# Create a heat map for the test coverage for the project and open in the default browser.
.PHONY: coverage-report
coverage-report : prepare-test
	scripts/coverage --html

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
