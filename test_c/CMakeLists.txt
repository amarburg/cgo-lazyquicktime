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

add_definitions( -std=c++11 )

## Build as C++ becuase the GTest framework is C++
set( TEST_SRCS
  test_movie_info.cpp
  test_get_frame.cpp )

add_executable(cgo-lazyquicktime-test ${TEST_SRCS} )
target_link_libraries( cgo-lazyquicktime-test
  lazyquicktime
  ${GTEST_BOTH_LIBRARIES}
  ${CMAKE_THREAD_LIBS_INIT} )

## Compile in paths to ExternalData
## Who the hell designed this API?
##
## (Hell no, this make no sense)
# include(ExternalData)
# set(ExternalData_URL_TEMPLATES
#     "https://github.com/amarburg/go-lazyfs-testfiles/raw/master/CamHD_Vent_Short.mov" )
#
# ExternalData_Add_Target( LocalTestMov )
# ExternalData_Expand_Arguments( LocalTestMov
#             LOCAL_TEST_MOV
#             DATA{CamHD_Vent_Short.mov}  )
target_compile_definitions( cgo-lazyquicktime-test PUBLIC LOCAL_TEST_MOV="${CMAKE_CURRENT_SOURCE_DIR}/CamHD_Vent_Short.mov" )
#

enable_testing()
add_test( NAME cgo-lazyquicktime-test
         COMMAND cgo-lazyquicktime-test )
