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

add_executable(cgo-lazyquicktime-test ${TEST_SRCS} )
target_link_libraries( cgo-lazyquicktime-test
  lazyquicktime
  ${GTEST_BOTH_LIBRARIES}
  ${CMAKE_THREAD_LIBS_INIT} )

## Compile in paths to ExternalData
## Who the hell designed this API?
include(ExternalData)
set(ExternalData_URL_TEMPLATES "file://$ENV{GOPATH}/src/github.com/amarburg/go-lazyfs-testfiles/hashed/%(algo)/%(hash)"
                                "https://github.com/amarburg/go-lazyfs-testfiles/raw/master/hashed/%(algo)/%(hash)" )

ExternalData_Expand_Arguments( LocalTestMov
            LOCAL_TEST_MOV
            DATA{CamHD_Vent_Short.mov}  )

ExternalData_Add_Target( LocalTestMov )

target_compile_definitions( cgo-lazyquicktime-test PUBLIC
                    CMAKE_BINARY_DIR="${CMAKE_BINARY_DIR}"
                    CMAKE_SOURCE_DIR="${CMAKE_SOURCE_DIR}")

enable_testing()
externaldata_add_test( NAME cgo-lazyquicktime-test
         COMMAND ./cgo-lazyquicktime-test
         DIRECTORY ${CMAKE_BINARY_DIR} )
add_dependencies( cgo-lazyquicktime-test LocalTestMov )
