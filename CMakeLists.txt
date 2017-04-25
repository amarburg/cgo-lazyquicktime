cmake_minimum_required(VERSION 2.8.7)
project(cgo-lazyquicktime-test)


include_directories( ${CMAKE_SOURCE_DIR} )
link_directories( ${CMAKE_SOURCE_DIR} )

set( TEST_SRCS
test_c/test_sum.c)

add_executable(cgo-lazyquicktime-test ${TEST_SRCS} )
target_link_libraries( cgo-lazyquicktime-test
  lazyquicktime )
