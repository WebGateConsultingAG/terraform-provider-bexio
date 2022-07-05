TEST?=$$(go list ./... |grep -v 'vendor'|grep -v 'examples')
HOSTNAME=hashicorp.com
NAMESPACE=wgc
NAME=bexio
VERSION=0.4.0
OS := $(if $(GOOS),$(GOOS),$(shell go env GOOS))
ARCH := $(if $(GOARCH),$(GOARCH),$(shell go env GOARCH))

default: build

build:
ifeq ($(OS),windows)
	go build -o terraform-provider-${NAME}.exe
else
	go build -o terraform-provider-${NAME}
endif

release:
	goreleaser release --rm-dist --snapshot --skip-publish  --skip-sign

install: build
ifeq ($(OS),windows)
	if not exist "%APPDATA%\terraform.d\plugins\${HOSTNAME}\${NAMESPACE}\${NAME}\${VERSION}\${OS}_${ARCH}" mkdir %APPDATA%\terraform.d\plugins\${HOSTNAME}\${NAMESPACE}\${NAME}\${VERSION}\${OS}_${ARCH}
	move terraform-provider-${NAME}.exe %APPDATA%\terraform.d\plugins\${HOSTNAME}\${NAMESPACE}\${NAME}\${VERSION}\${OS}_${ARCH}
	if exist ".\.terraform" rmdir /s /q .\.terraform
	if exist ".\.terraform.lock.hcl" del .\.terraform.lock.hcl
	if exist ".\terraform.tfstate" del .\terraform.tfstate
else
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS}_${ARCH}
	mv terraform-provider-${NAME} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS}_${ARCH}
	rm -rf ./.terraform 
	rm -rf ./.terraform.lock.hcl
	rm -rf ./terraform.tfstate
endif

fmt:
	@echo "==> Format source code with gofmt..."
	find . -name '*.go' | grep -v vendor | xargs gofmt -s -w

test: 
	go test $(TEST) || exit 1                                                   
	echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4                    

testacc: 
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m   
