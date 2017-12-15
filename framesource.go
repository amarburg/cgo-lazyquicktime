package main

// #include "types.h"
import "C"

import (
	"fmt"
	"github.com/amarburg/go-frameset/framesource"
	"sync"
	"io"
)

//TODO.  is Printf on error the best way to get error information out?

type FrameSourceMap struct {
	sync.RWMutex
	nextId				int
	internal			map[int]framesource.FrameSource
}

func (rm *FrameSourceMap) Load(key int) (framesource.FrameSource, bool) {
	rm.RLock()
	defer rm.RUnlock()

	value,ok := rm.internal[key]
	return value,ok
}

func (rm *FrameSourceMap) Delete(key int) {
	rm.Lock()
	defer rm.Unlock()
	delete(rm.internal, key)
}

func (rm *FrameSourceMap) Add(value framesource.FrameSource) int {
	rm.Lock()
	defer rm.Unlock()
	id := rm.nextId
	rm.internal[id] = value
	rm.nextId++
return id
}


var IdMap FrameSourceMap

//export OpenFrameSource
func OpenFrameSource(path *C.char) C.int {

	// Todo, look for duplicates

	goPath := C.GoString(path)
	source, err := framesource.MakeFrameSourceFromPath(goPath)

	if err != nil {
		fmt.Printf("Error extracting image: %s", err.Error())
		return -1
	}

	// todo.  Make threadsafe
	return C.int( IdMap.Add( source ) )
}

//export CloseFrameSource
func CloseFrameSource(id C.int) {
	IdMap.Delete(int(id))
}

//export FrameSourceNext
func FrameSourceNext(id C.int) C.ImageBuffer {

	source, has := IdMap.Load(int(id))

	if !has {
		fmt.Printf("Id doesn't exist")
		return C.ImageBuffer{}
	}

	img, err := source.Next()

	switch err {
	case nil:
		return imageBufferFromImage(img)
	case io.EOF:
		return C.ImageBuffer{}
	default:
		fmt.Printf("Error extracting image: %s", err.Error())
		return C.ImageBuffer{}
	}
}
