cmake_minimum_required(VERSION 2.8.7)
project(cgo-lazyquicktime-test)

# Per best practice, make a local copy of GTest
include( ${CMAKE_SOURCE_DIR}/cmake/BuildGTestLibraries.cmake )

include_directories( ${CMAKE_SOURCE_DIR}/..
                      ${GTEST_INCLUDE_DIRS} )
link_directories( ${CMAKE_SOURCE_DIR}/.. )

#if(LINUX)
  find_package( Threads REQUIRED )
#endif()

if(APPLE)
  ## Force absolute paths to shared libraries
  cmake_policy(SET CMP0060 NEW)
  set( MACOSX_RPATH OFF )
  set(CMAKE_SKIP_RPATH 1)
endif()

add_definitions( -std=c++11 )

## Build as C++ becuase the GTest framework is C++
file( GLOB TEST_SRCS *.cpp )

add_executable(cgo-movieextractor-test ${TEST_SRCS} )
target_link_libraries( cgo-movieextractor-test
  movieextractor
  ${GTEST_BOTH_LIBRARIES}
  ${CMAKE_THREAD_LIBS_INIT} )

set( TESTFILE_DIR "${CMAKE_SOURCE_DIR}/../../go-lazyfs-testfiles")

target_compile_definitions( cgo-movieextractor-test PUBLIC
                    CMAKE_BINARY_DIR="${CMAKE_BINARY_DIR}"
                    CMAKE_SOURCE_DIR="${CMAKE_SOURCE_DIR}"
                    TESTFILE_DIR="${TESTFILE_DIR}")

enable_testing()
add_test( NAME cgo-movieextractor-test
         COMMAND ./cgo-movieextractor-test
         DIRECTORY ${CMAKE_BINARY_DIR} )
#add_dependencies( cgo-movieextractor-test LocalTestMov )
