cmake_minimum_required(VERSION 2.8.7)
project(cgo-lazyquicktime-cmd)

find_package( OpenCV REQUIRED core highgui imgproc )
find_package( Threads REQUIRED )
find_package( Boost REQUIRED program_options )

if(APPLE)
  ## Force absolute paths to shared libraries
  cmake_policy(SET CMP0060 NEW)
  set( MACOSX_RPATH OFF )
  set(CMAKE_SKIP_RPATH 1)
endif()

add_definitions( -std=c++11 )

include_directories( ${CMAKE_SOURCE_DIR}/..
                      ${OpenCV_INCLUDE_DIRS} )

link_directories( ${CMAKE_SOURCE_DIR}/.. )

add_executable(player
              player.cpp )

target_link_libraries( player
              lazyquicktime
              ${OpenCV_LIBS}
              ${Boost_LIBRARIES} )

## Build as C++ becuase the GTest framework is C++
file( GLOB TEST_SRCS *.cpp )
