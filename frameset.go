package main

// #include "types.h"
// #include <errno.h>
import "C"

import (
	"fmt"
	"github.com/amarburg/go-movieset"
)

var FSMap FrameSetMap

func init() {
	FSMap.internal = make(map[int]*movieset.FrameSet)
}

//export OpenFrameSet
func OpenFrameSet(path *C.char) C.int {

	// Todo, look for duplicates

	goPath := C.GoString(path)
	source, err := movieset.LoadFrameSet(goPath)

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

// //export FrameSourceFromFrameSet
// func FrameSourceFromFrameSet(id C.int) C.int {
//
// 	set, has := FSMap.Load(int(id))
//
// 	if !has {
// 		fmt.Printf("FrameSet %d doesn't exist", id)
// 		return -1
// 	}
//
// 	extractor, err := set.MovieExtractor()
//
// 	if err != nil {
// 		fmt.Printf("Couldn't create movie extractor")
// 		return -1
// 	}
//
// 	source,err := movieset.MakeMovieSequential( extractor )
//
// 	if err != nil {
// 		fmt.Printf("Couldn't convert movie extractor to frame source")
// 		return -1
// 	}
//
// 	return C.int(IdMap.Add(source))
// }


//export FrameSetToSequential
func FrameSetToSequential( id C.int ) C.int {

	set, has := FSMap.Load(int(id))

	if !has {
		fmt.Printf("FrameSet %d doesn't exist", id)
		return -1
	}

	seq,err := movieset.MakeFrameSetSequential( set )

	if err != nil {
		fmt.Printf("Error turning frame set into sequential: %s", err);
		return -1
	}

	return C.int(IdMap.Add(seq))
}



//export OpenFrameSetChunk
func OpenFrameSetChunk(id C.int, chunkName *C.char) C.int {
	chunk := C.GoString(chunkName)
	set, has := FSMap.Load(int(id))

	if !has {
		fmt.Printf("FrameSet %d doesn't exist", id)
		return -1
	}

	chunkmov,err := set.MovFromChunk( chunk )
	if err != nil {
		fmt.Printf("Can't convert chunk to movie: %s\n", err)
		return -1
	}

	mov,err := movieset.MakeMovieSequential(chunkmov)
	if err != nil {
		fmt.Printf("Can't make sequential from movie for chunk \"%s\": %s\n", chunk, err )
		return -1
	}

	return C.int(IdMap.Add(mov))
}
