package main

// #include "types.h"
import "C"

import (
	"fmt"
	"github.com/amarburg/go-frameset/multimov"
)

//== These are the "oneshot" versions

//export GetFrame
func GetFrame(path *C.char, frameNum C.int) C.ImageBuffer {

	goPath := C.GoString(path)
	ext, err := multimov.MovieExtractorFromPath(goPath)

	if err != nil {
		fmt.Printf("Error extracting image: %s", err.Error())
		return C.ImageBuffer{}
	}

	img, err := ext.ExtractFrame(uint64(frameNum))

	if err != nil {
		fmt.Printf("Error extracting image: %s", err.Error())
		return C.ImageBuffer{}
	}

	return imageBufferFromImage(img)
}

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
