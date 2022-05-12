#update path
PATH := $(PATH):$(GOPATH)/bin
GOROOT := $(GOROOT)
GOPATH := $(GOPATH)

.PHONY: INFO
info-run:
	@echo "PATH:"
	@echo $(PATH)
	@echo "GoROOT:"
	@echo $(GOROOT)
	@echo "GoPATH:"
	@echo $(GOPATH)

.PHONY: RUN
run: copy-proto go-dep

.PHONY: dependency
go-dep:
	go mod download

.PHONY: generate and copy proto files
copy-proto:
	make -C ../../protolib run
	cp -r ../../protolib/gen/proto/go/* ./proto


