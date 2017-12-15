package main

// #include "types.h"
import "C"

import (
	"fmt"
	"github.com/amarburg/go-multimov"
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
