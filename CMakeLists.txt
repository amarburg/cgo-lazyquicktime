cmake_minimum_required(VERSION 2.8.7)
project(cgo-lazyquicktime-test)

# Per best practice, make a local copy of GTest
include( ${CMAKE_SOURCE_DIR}/cmake/BuildGTestLibraries.cmake )

include_directories( ${CMAKE_SOURCE_DIR}
                      ${GTEST_INCLUDE_DIRS} )
link_directories( ${CMAKE_SOURCE_DIR} )

add_definitions( -std=c++11 )

## Build as CPP to using GoogleTest framework
set( TEST_SRCS
  test_c/test_sum.cpp
  test_c/test_simple_api.cpp )

add_executable(cgo-lazyquicktime-test ${TEST_SRCS} )
target_link_libraries( cgo-lazyquicktime-test
  lazyquicktime
  ${GTEST_BOTH_LIBRARIES} )
