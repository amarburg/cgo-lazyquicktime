package main

// #include "types.h"
import "C"

import (
	"fmt"
	"github.com/amarburg/go-frameset/multimov"
)

//== These are the "oneshot" versions

//export GetFrame
func GetFrame(path *C.char, frameNum C.int, out *C.ImageBuffer) int {

	goPath := C.GoString(path)
	ext, err := multimov.MovieExtractorFromPath(goPath)

	if err != nil {
		fmt.Printf("Error extracting image: %s", err.Error())
		out.valid = C.uchar(0)
		return -1
	}

	img, err := ext.ExtractFrame(uint64(frameNum))

	if err != nil {
		fmt.Printf("Error extracting image: %s", err.Error())
		out.valid = C.uchar(0)
		return -1
	}

	imageToImageBuffer( img, out )
	return 0
}

//export MovInfo
func MovInfo(path *C.char, info *C.MovieInfo) int {

	goPath := C.GoString(path)
	ext, err := multimov.MovieExtractorFromPath(goPath)

	if err != nil || ext == nil {
		info.valid = C.uchar(0)
		return -1
	}

	info.duration =C.float(ext.Duration().Seconds())
	info.num_frames = C.int(ext.NumFrames())
	info.valid = C.uchar(1)

	return 0
}
