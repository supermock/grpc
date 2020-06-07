GO      = go
PROTOC  = protoc

PROTO_PATH                := proto
PROTO_SRC_FILES           := $(shell find $(PROTO_PATH) -name '*.proto')
PROTO_GEN_FILES           := $(PROTO_SRC_FILES:$(PROTO_PATH)/%.proto=%.pb.go)
PROTO_GEN_FILES_RM_TARGET := $(addprefix rm/,$(PROTO_GEN_FILES))

proto_gw_file = $(1:%.go=%.gw.go)
proto_sw_file = $(1:%.pb.go=%.swagger.json)

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
	$(PROTOC) $(INCLUDE) --go_out=plugins=grpc:. --go_opt=paths=source_relative $<
	if [ -n "$(findstring gateway,$(HOOKS))" ]; then \
		echo "Making $(call proto_gw_file,$(@))"; \
		$(PROTOC) $(INCLUDE) --grpc-gateway_out=logtostderr=false,paths=source_relative:. $<; \
	fi
	if [ -n "$(findstring swagger,$(HOOKS))" ]; then \
		echo "Making $(call proto_sw_file,$(@))"; \
		$(PROTOC) $(INCLUDE) --swagger_out=logtostderr=false:. $<; \
	fi

.SILENT: $(CMD_BIN_FILES)
$(CMD_BIN_FILES): build/%: cmd/%/main.go $(GO_SRC_FILES) $(PROTO_GEN_FILES)
	mkdir -p build
	echo "Making $@"
	$(GO) build -o $@ $<

.PHONY: $(PROTO_GEN_FILES_RM_TARGET)
.SILENT: $(PROTO_GEN_FILES_RM_TARGET)
$(PROTO_GEN_FILES_RM_TARGET):	PROTO_GEN_PB_FILE = $(subst rm/,,$(@))
$(PROTO_GEN_FILES_RM_TARGET):	PROTO_GEN_GW_FILE = $(call proto_gw_file,$(PROTO_GEN_PB_FILE))
$(PROTO_GEN_FILES_RM_TARGET):	PROTO_GEN_SW_FILE = $(call proto_sw_file,$(PROTO_GEN_PB_FILE))
$(PROTO_GEN_FILES_RM_TARGET):
	rm -f $(PROTO_GEN_PB_FILE)
	rm -f $(PROTO_GEN_GW_FILE)
	rm -f $(PROTO_GEN_SW_FILE)

.PHONY: protobuf
protobuf: $(PROTO_GEN_FILES)

.PHONY: protobuf-clean
.SILENT: protobuf-clean
protobuf-clean: $(PROTO_GEN_FILES_RM_TARGET)
	echo "Cleanup completed"
