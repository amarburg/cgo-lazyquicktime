package main

// #include "types.h"
import "C"

//export MovInfo
func MovInfo( url *C.char )  C.MovieInfo {
  return C.MovieInfo{
    num_frames: 1234,
    duration: 0.1234,
  }
}
