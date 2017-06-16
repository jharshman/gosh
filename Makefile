#
# GoSH Makefile
#

.PHONY: all check-path get-deps lint test proto

all: check-path get-deps test

check-path:
	@echo "[*] checking path..."
ifndef GOPATH
	@echo "[!] FATAL: GOPATH not defined"
	@echo "Fix your Go Installation and try again"
	@echo "For more information: https://golang.org/doc/install"
	@echo "exit 1"
endif
ifneq ($(subst ~,$(HOME),$(GOPATH))/src/github.com/jharshman/gosh, $(PWD))
	@echo "[!] FATAL: source not in GOPATH"
	@echo "go get github.com/jharshman/gosh"
	@echo "exit 1"
endif
	@echo "all good - exit 0"

get-deps: check-path
	@echo "[*] getting dependencies..."
	go get ./...

lint: get-deps
	@echo "[*] linting..."
	golint ./...

proto:
	@echo "[*] protoc... "
	while read file; do 																\
		echo "[*] compiling ${file}"; 										\
		protoc -I=${file%\/*} --go_out=${file%\/*} $file; \
	done < <(find . -name "*.proto");
