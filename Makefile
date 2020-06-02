PROTOS_PATH := protos

PROTO_SRC_FILES := $(shell find $(PROTOS_PATH) -name '*.proto')
PROTO_GEN_FILES := $(PROTO_SRC_FILES:$(PROTOS_PATH)/%.proto=%.pb.go)
PROTO_SRC_TARGETS := $(PROTO_GEN_FILES:%.pb.go=$(PROTOS_PATH)/%.proto)
proto_gen_target = $(addprefix $(dir $(proto_gen_file)),%.pb.go)
PROTO_GEN_TARGETS := $(foreach proto_gen_file,$(PROTO_GEN_FILES),$(proto_gen_target))

CMD_SRC_FILES := $(wildcard cmd/**/main.go)
CMD_SRC_TARGETS := cmd/%/main.go
CMD_BIN_FILES := $(CMD_SRC_FILES:cmd/%/main.go=build/%)
CMD_BIN_TARGETS := build/%

GO_SRC_FILES := $(shell find . -name '*.go' -not -path './cmd/*')

.PHONY: all
all: $(CMD_BIN_FILES)

.SILENT: $(PROTO_GEN_FILES)
$(PROTO_GEN_TARGETS): $(PROTO_SRC_TARGETS)
	mkdir -p $(dir $@)
	echo "Making $@"
	protoc --proto_path=$(PROTOS_PATH) --go_out=plugins=grpc:. --go_opt=paths=source_relative $<

.SILENT: $(CMD_BIN_FILES)
$(CMD_BIN_TARGETS): $(CMD_SRC_TARGETS) $(GO_SRC_FILES) $(PROTO_GEN_FILES)
	mkdir -p build
	echo "Making $@"
	go build -o $@ $<
