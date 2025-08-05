#
ARCH=$(shell uname -m)
SERVERS:=$(shell ls app)

WORK_DIR:=$(shell pwd)

PB_VERSION:=30.2
PB_FILES = $(shell find . -name "*.proto")
PB_DIRS = $(sort $(dir $(PB_FILES)))
PB_GO_FILES = $(shell find . -name "*.pb.go")
PB_DIR_TGTS = $(addprefix _PB, $(PB_DIRS))

.PHONY: servers
servers: $(SERVERS)

.PHONY: $(SERVERS)
$(SERVERS):
	@echo "Building server: $@..."
	@go mod tidy && go build ./app/$@/
	@mv $@ bin/ && echo Done

.PHONY: pb
pb: $(PB_DIR_TGTS)

.PHONY: $(PB_DIR_TGTS)
$(PB_DIR_TGTS):
	@for dir in $(subst _PB,, $@); do \
		echo Now Build proto in directory: $$dir; \
		cd $$dir; rm -rf mock; \
		export PATH=$(PATH); \
		rm -f *.pb.go; rm -f *.trpc.go; \
		find . -name "*.proto" | xargs -I DD \
			trpc create -f --protofile=DD --protocol=trpc --rpconly --nogomod --alias --mock=false --protodir=$(WORK_DIR)/proto; \
		ls *.trpc.go | xargs -I DD mockgen -source=DD -destination=mock/DD -package=mock ; \
		find `pwd` -name '*.pb.go'; \
	done