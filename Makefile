PROTO_PATH := proto

PROTO_SRC_FILES := $(shell find $(PROTO_PATH) -name '*.proto')
PROTO_GEN_FILES := $(PROTO_SRC_FILES:$(PROTO_PATH)/%.proto=%.pb.go)

CMD_SRC_FILES := $(wildcard cmd/**/main.go)
CMD_BIN_FILES := $(CMD_SRC_FILES:cmd/%/main.go=build/%)

GO_SRC_FILES := $(shell find . -name '*.go' -not -path './cmd/*')

.PHONY: all
all: $(CMD_BIN_FILES)

.SILENT: $(PROTO_GEN_FILES)
$(PROTO_GEN_FILES): INCLUDE = -I$(PROTO_PATH) -I/usr/local/include -I${GOPATH}/src
$(PROTO_GEN_FILES): INCLUDE += -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway
$(PROTO_GEN_FILES): INCLUDE += -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis
$(PROTO_GEN_FILES): HOOKS = $(shell jq -rcM --arg proto $(<F) '.hooks[$$proto]? | @sh' $(PROTO_PATH)/$(@D)/metadata.json)
$(PROTO_GEN_FILES): %.pb.go: $(PROTO_PATH)/%.proto
	mkdir -p $(@D)
	echo "Making $@"
	protoc $(INCLUDE) --go_out=plugins=grpc:. --go_opt=paths=source_relative $<
	if [ -n "$(findstring gateway,$(HOOKS))" ]; then \
		echo "Making ${@:%.go=%.gw.go}"; \
		protoc $(INCLUDE) --grpc-gateway_out=logtostderr=false,paths=source_relative:. $<; \
	fi
	if [ -n "$(findstring swagger,$(HOOKS))" ]; then \
		echo "Making ${@:%.pb.go=%.swagger.json}"; \
		protoc $(INCLUDE) --swagger_out=logtostderr=false:. $<; \
	fi

.SILENT: $(CMD_BIN_FILES)
$(CMD_BIN_FILES): build/%: cmd/%/main.go $(GO_SRC_FILES) $(PROTO_GEN_FILES)
	mkdir -p build
	echo "Making $@"
	go build -o $@ $<

.PHONY: protobuf
protobuf: $(PROTO_GEN_FILES)
