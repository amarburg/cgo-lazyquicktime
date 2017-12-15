package main

// #include "types.h"
import "C"

import (
	"fmt"
	"github.com/amarburg/go-frameset/multimov"
	"github.com/amarburg/go-lazyquicktime"
	"sync"
)

//TODO.  is Printf on error the best way to get error information out?

var QTIds sync.Map
var NextId = 0

//export OpenQt
func OpenQt(path *C.char) C.int {

	// Todo, look for duplicates

	goPath := C.GoString(path)
	ext, err := multimov.MovieExtractorFromPath(goPath)

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

//export CloseQt
func CloseQt(id C.int) {
	QTIds.Delete(int(id))
}

//export GetFrameQt
func GetFrameQt(id C.int, frameNum C.int, out *C.ImageBuffer) int {

	val, has := QTIds.Load(int(id))

	if !has {
		fmt.Printf("Id doesn't exist")
		return -1
	}

	ext := val.(lazyquicktime.MovieExtractor)

	img, perf, err := ext.ExtractFramePerf(uint64(frameNum))

	fmt.Printf("Read took %g; decode took %g\n", perf.Read.Seconds(), perf.Decode.Seconds())

	if err != nil {
		fmt.Printf("Error extracting image: %s", err.Error())
		return -1
	}

	imageToImageBuffer(img, out)
	return 0

}

//export GetMovieInfoQt
func GetMovieInfoQt(id C.int) C.MovieInfo {

	val, has := QTIds.Load(int(id))
	if !has {
		fmt.Printf("Id doesn't exist")
		return C.MovieInfo{
			valid: 0,
		}
	}

	ext := val.(lazyquicktime.MovieExtractor)
	return C.MovieInfo{
		duration:   C.float(ext.Duration().Seconds()),
		num_frames: C.int(ext.NumFrames()),
		valid:      1,
	}

}
