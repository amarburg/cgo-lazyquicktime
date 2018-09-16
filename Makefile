
current_dir = $(shell pwd)
UNAME_S := $(shell uname -s)

ifeq ($(UNAME_S),Darwin)
			export DYLD_LIBRARY_PATH=gtest/lib
endif

LIB=libmovieset.so


default: test

${LIB}: *.go  ${GOPATH}/src/github.com/amarburg/go-movieset/*.go
	go build -buildmode=c-shared -o ${LIB}
ifeq ($(UNAME_S),Darwin)
	    install_name_tool -id ${current_dir}/$@ $@
endif


test: ${LIB}
	mkdir -p test_c/build
	go get -u github.com/amarburg/go-lazyfs-testfiles
	cd test_c/build && \
			cmake .. && \
			make gtest_ext && \
			make && \
	 		make CTEST_OUTPUT_ON_FAILURE=TRUE test


cmd: ${LIB}
		mkdir -p cmd/build
		cd cmd/build && cmake .. && make


clean:
	rm -f ${LIB} ${LIB:.so=.h}
	rm -rf test_c/build/

.PHONY: build test clean cmd default
