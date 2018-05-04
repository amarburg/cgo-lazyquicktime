package main

// #include "types.h"
// #include <errno.h>
import "C"

import (
	"fmt"
	"github.com/amarburg/go-movieset"
	"io"
)

var IdMap SequentialMap

func init() {
	IdMap.internal = make(map[int]movieset.Sequential)
}

//export OpenSequential
func OpenSequential(path *C.char) C.int {

	// Todo, look for duplicates

	goPath := C.GoString(path)
	source, err := movieset.OpenSequential(goPath)

	if err != nil {
		fmt.Printf("Error extracting image: %s", err.Error())
		return -1
	}

	// todo.  Make threadsafe
	return C.int(IdMap.Add(source))
}

//export CloseSequential
func CloseSequential(id C.int) {
	IdMap.Delete(int(id))
}

//export SequentialNext
func SequentialNext(id C.int, buffer *C.ImageBuffer) int64 {

	source, has := IdMap.Load(int(id))

	if !has {
		fmt.Printf("Id doesn't exist")
		return -1
	}

	img, frameNum, err := source.Next()

	switch err {
	case nil:
		imageToImageBuffer(img, buffer)
		return int64(frameNum)
	case io.EOF:
		return 0
	default:
		fmt.Printf("Error extracting image: %s", err.Error())
		return -1
	}
}
