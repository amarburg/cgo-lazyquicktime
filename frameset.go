package main

// #include "types.h"
// #include <errno.h>
import "C"

import (
	"fmt"
	"github.com/amarburg/go-frameset/frameset"
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

//export
