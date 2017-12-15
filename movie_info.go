package main

// #include "types.h"
import "C"
import (
	"github.com/amarburg/go-multimov"
)

//export MovInfo
func MovInfo(path *C.char) C.MovieInfo {

	goPath := C.GoString(path)
	ext, err := multimov.MovieExtractorFromPath(goPath)

	if err != nil || ext == nil {
		return C.MovieInfo{
			valid: 0,
		}

	}

	return C.MovieInfo{
		duration:   C.float(ext.Duration().Seconds()),
		num_frames: C.int(ext.NumFrames()),
		valid:      1,
	}

}
