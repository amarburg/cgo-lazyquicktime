
current_dir = $(shell pwd)
UNAME_S := $(shell uname -s)

ifeq ($(UNAME_S),Darwin)
			export DYLD_LIBRARY_PATH=gtest/lib
endif

liblazyquicktime.so: *.go
	go build -buildmode=c-shared -o liblazyquicktime.so
ifeq ($(UNAME_S),Darwin)
	    install_name_tool -id ${current_dir}/liblazyquicktime.so liblazyquicktime.so
endif


test: liblazyquicktime.so
	mkdir -p test_c/build
	cd test_c/build && \
			cmake .. && \
			make gtest_ext && \
			make && \
	 		make CTEST_OUTPUT_ON_FAILURE=TRUE test


clean:
	rm -f liblazyquicktime.so
	rm -rf test_c/build/

PHONY: build test clean
