package main

// #include "types.h"
// #include <errno.h>
import "C"

import (
	"fmt"
	"github.com/amarburg/go-frameset/frameset"
	"github.com/amarburg/go-frameset/framesource"
)

var FSMap FrameSetMap

func init() {
	FSMap.internal = make(map[int]*frameset.FrameSet)
}

//export OpenFrameSet
func OpenFrameSet(path *C.char) C.int {

	// Todo, look for duplicates

	goPath := C.GoString(path)
	source, err := frameset.LoadFrameSet(goPath)

	if err != nil {
		fmt.Printf("Error loading frame set: %s", err.Error())
		return -1
	}

	return C.int(FSMap.Add(source))
}

//export CloseFrameSet
func CloseFrameSet(id C.int) {
	FSMap.Delete(int(id))
}

//export FrameSourceFromFrameSet
func FrameSourceFromFrameSet(id C.int) C.int {

	set, has := FSMap.Load(int(id))

	if !has {
		fmt.Printf("FrameSet %d doesn't exist", id)
		return -1
	}

	extractor, err := set.MovieExtractor()

	if err != nil {
		fmt.Printf("Couldn't create movie extractor")
		return -1
	}

	source,err := framesource.MakeMovieExtractorFrameSource( extractor )

	if err != nil {
		fmt.Printf("Couldn't convert movie extractor to frame source")
		return -1
	}

	return C.int(IdMap.Add(source))
}
