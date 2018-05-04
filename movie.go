package main

// #include "types.h"
import "C"

import (
	"fmt"
	"github.com/amarburg/go-movieset"
	"sync"
)

var QTIds sync.Map
var NextId = 0

//export OpenMovie
func OpenMovie(path *C.char) C.int {

	// Todo, look for duplicates

	goPath := C.GoString(path)
	ext, err := movieset.OpenMovieExtractor(goPath)

	if err != nil {
		fmt.Printf("Error extracting image: %s", err.Error())
		return -1
	}

	// todo.  Make threadsafe
	id := NextId
	QTIds.Store(id, ext)
	NextId++

	return C.int(id)

}

//export CloseMovie
func CloseMovie(id C.int) {
	QTIds.Delete(int(id))
}

//export GetMovieFrame
func GetMovieFrame(id C.int, frameNum C.int, out *C.ImageBuffer) int {

	val, has := QTIds.Load(int(id))

	if !has {
		fmt.Printf("Id doesn't exist")
		return -1
	}

	ext := val.(movieset.MovieExtractor)

	img, err := ext.ExtractFrame(uint64(frameNum))

	// img, perf, err := ext.ExtractFramePerf(uint64(frameNum))
	//
	// fmt.Printf("Read took %g; decode took %g\n", perf.Read.Seconds(), perf.Decode.Seconds())

	if err != nil {
		fmt.Printf("Error extracting image: %s", err.Error())
		return -1
	}

	imageToImageBuffer(img, out)
	return 0

}

//export GetMovieInfo
func GetMovieInfo(id C.int, info *C.MovieInfo) int {

	val, has := QTIds.Load(int(id))
	if !has {
		info.valid = C.uchar(0)
		return -1
	}

	ext := val.(movieset.MovieExtractor)

	info.duration = C.float(ext.Duration().Seconds())
	info.num_frames = C.int(ext.NumFrames())
	info.valid = C.uchar(1)

	return 0

}
