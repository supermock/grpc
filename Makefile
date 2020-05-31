PROTO_SRC_FILES := $(wildcard *.proto)
PROTO_SRC_TARGETS := %.proto
proto_gen_path = $(addprefix $(proto_src:%.proto=%)/, $(proto_src:%.proto=%.pb.go))
PROTO_GEN_FILES := $(foreach proto_src,$(PROTO_SRC_FILES),$(proto_gen_path))
PROTO_GEN_TARGETS := $(PROTO_SRC_FILES:%.proto=%/%.pb.go)

CMD_SRC_FILES := $(wildcard cmd/**/main.go)
CMD_SRC_TARGETS := cmd/%/main.go
CMD_BIN_FILES := $(CMD_SRC_FILES:cmd/%/main.go=build/%)
CMD_BIN_TARGETS := build/%

GO_SRC_FILES := $(shell find $(CUR) -name '*.go' -not -path './cmd/*')

.PHONY: all
all: $(CMD_BIN_FILES)

.SILENT: $(PROTO_GEN_FILES)
$(PROTO_GEN_TARGETS): $(PROTO_SRC_TARGETS)
	$(eval $@_name := $(<:$(PROTO_SRC_TARGETS)=%))
	mkdir -p $($@_name)
	echo "Making $@"
	protoc --go_out=plugins=grpc:$($@_name) $<

.SILENT: $(CMD_BIN_FILES)
$(CMD_BIN_TARGETS): $(CMD_SRC_TARGETS) $(GO_SRC_FILES) $(PROTO_GEN_FILES)
	mkdir -p build
	echo "Making $@"
	go build -o $@ $<
