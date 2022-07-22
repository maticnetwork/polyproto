GOBINREL = build/bin
GOBIN = $(CURDIR)/$(GOBINREL)
GOBUILD = env GO111MODULE=on go build -trimpath
OS = $(shell uname -s)
ARCH = $(shell uname -m)

ifeq ($(ARCH),arm64)
ARCH = aarch_64
endif

ifeq ($(OS),Darwin)
PROTOC_OS := osx
endif
ifeq ($(OS),Linux)
PROTOC_OS = linux
endif

PROTOC_INCLUDE = build/include/google

$(GOBINREL):
	mkdir -p "$(GOBIN)"

$(GOBINREL)/protoc: | $(GOBINREL)
	$(eval PROTOC_TMP := $(shell mktemp -d))
	curl -sSL https://github.com/protocolbuffers/protobuf/releases/download/v21.2/protoc-21.2-$(PROTOC_OS)-$(ARCH).zip -o "$(PROTOC_TMP)/protoc.zip"
	cd "$(PROTOC_TMP)" && unzip protoc.zip
	cp "$(PROTOC_TMP)/bin/protoc" "$(GOBIN)"
	mkdir -p "$(PROTOC_INCLUDE)"
	cp -R "$(PROTOC_TMP)/include/google/" "$(PROTOC_INCLUDE)"
	rm -rf "$(PROTOC_TMP)"

# 'protoc-gen-go' tool generates proto messages
$(GOBINREL)/protoc-gen-go: | $(GOBINREL)
	$(GOBUILD) -o "$(GOBIN)/protoc-gen-go" google.golang.org/protobuf/cmd/protoc-gen-go

# 'protoc-gen-go-grpc' tool generates grpc services
$(GOBINREL)/protoc-gen-go-grpc: | $(GOBINREL)
	$(GOBUILD) -o "$(GOBIN)/protoc-gen-go-grpc" google.golang.org/grpc/cmd/protoc-gen-go-grpc

protoc-all: $(GOBINREL)/protoc $(PROTOC_INCLUDE) $(GOBINREL)/protoc-gen-go $(GOBINREL)/protoc-gen-go-grpc

protoc-clean:
	rm -f "$(GOBIN)/protoc"*
	rm -rf "$(PROTOC_INCLUDE)"

grpc: protoc-all
	PATH="$(GOBIN):$(PATH)" protoc --proto_path=. --go_out=. --go-grpc_out=. -I=$(PROTOC_INCLUDE) \
		heimdallclient/heimdall_client.proto
