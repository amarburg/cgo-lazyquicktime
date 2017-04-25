package main

// #include "types.h"
import "C"
import (
	"fmt"
	"github.com/amarburg/go-lazyquicktime"
)

//export MovInfo
func MovInfo(path *C.char) C.MovieInfo {

	file, err := sourceFromCPath(path)

	if err != nil {
		fmt.Printf("Error opening path: %s", err.Error())
		return C.MovieInfo{}
	}

	qtInfo, err := lazyquicktime.LoadMovMetadata(file)

	if err != nil {
		fmt.Printf("Error getting metadata: %s", err.Error())
		return C.MovieInfo{}
	}

	return qtInfoToMovieInfo(qtInfo)
}

func qtInfoToMovieInfo(qtInfo *lazyquicktime.LazyQuicktime) C.MovieInfo {
	return C.MovieInfo{
		duration:   C.float(qtInfo.Duration()),
		num_frames: C.int(qtInfo.NumFrames()),
		valid:      1,
	}
}
