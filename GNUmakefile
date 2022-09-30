GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)
PKG_NAME=outscale
TEST?=./...
PROVIDER_VERSION=0.5.4 #TODO get version automaticly
PROVIDER_NAME=terraform-provider-outscale_v$(PROVIDER_VERSION)
INSTALL_DIR=~/.terraform.d/plugins/registry.terraform.io/outscale-dev/outscale/$(PROVIDER_VERSION)/linux_amd64/
WEBSITE_REPO=github.com/hashicorp/terraform-website

.PHONY: default
default: build

.PHONY: build
build: fmtcheck
	go build -v -o $(PROVIDER_NAME)

.PHONY: doc
doc:
	@sh -c "'$(CURDIR)/scripts/generate-doc.sh'"

.PHONY: fmt
fmt:
	echo "==> Fixing source code with gofmt..."
	gofmt -s -w ./main.go
	gofmt -s -w ./$(PKG_NAME)

.PHONY: fmtcheck
fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

.PHONY: install
install: build
	mkdir -p $(INSTALL_DIR)
	mv $(PROVIDER_NAME) $(INSTALL_DIR)

.PHONY: lint
lint:
	echo "==> Checking source code against linters..."
	@GOGC=30 golangci-lint run ./$(PKG_NAME)  --deadline=30m

.PHONY: fmtcheck
test: fmtcheck
	go test $(TEST) -count 1 -timeout=30s -parallel=4

.PHONY: testacc
testacc: fmtcheck
	TF_ACC=1 go test $(TEST) -count 1 -v -parallel 4 $(TESTARGS) -timeout 240m -cover

.PHONY: test-compile
test-compile:
	if [ "$(TEST)" = "./..." ]; then \
		echo "ERROR: Set TEST to a specific package. For example,"; \
		echo "  make test-compile TEST=./$(PKG_NAME)"; \
		exit 1; \
	fi
	go test -c $(TEST) $(TESTARGS)

.PHONY: test-examples
test-examples: build
	@sh -c "'$(CURDIR)/scripts/test-examples.sh'"

.PHONY: test-integration
test-integration: build
	@sh -c "'$(CURDIR)/scripts/integration.sh'"

.PHONY: tools
tools:
	GO111MODULE=off go get -u github.com/client9/misspell/cmd/misspell
	GO111MODULE=off go get -u github.com/golangci/golangci-lint/cmd/golangci-lint

.PHONY: website
website:
ifeq (,$(wildcard $(GOPATH)/src/$(WEBSITE_REPO)))
	echo "$(WEBSITE_REPO) not found in your GOPATH (necessary for layouts and assets), get-ting..."
	@git clone https://$(WEBSITE_REPO) $(GOPATH)/src/$(WEBSITE_REPO)
endif
	@$(MAKE) -C $(GOPATH)/src/$(WEBSITE_REPO) website-provider PROVIDER_PATH=$(shell pwd) PROVIDER_NAME=$(PKG_NAME)

.PHONY: websitefmtcheck
websitefmtcheck:
	@sh -c "'$(CURDIR)/scripts/websitefmtcheck.sh'"

.PHONY: website-lint
website-lint:
	echo "==> Checking website against linters..."
	@misspell -error -source=text website/

.PHONY: website-local
website-local:
	@sh -c "'$(CURDIR)/scripts/test-doc.sh'"

.PHONY: website-test
website-test:
ifeq (,$(wildcard $(GOPATH)/src/$(WEBSITE_REPO)))
	echo "$(WEBSITE_REPO) not found in your GOPATH (necessary for layouts and assets), get-ting..."
	@git clone https://$(WEBSITE_REPO) $(GOPATH)/src/$(WEBSITE_REPO)
endif
	@$(MAKE) -C $(GOPATH)/src/$(WEBSITE_REPO) website-provider-test PROVIDER_PATH=$(shell pwd) PROVIDER_NAME=$(PKG_NAME)
