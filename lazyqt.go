package main

// #include "types.h"
import "C"

import (
	"fmt"
	"github.com/amarburg/go-lazyquicktime"
	"github.com/amarburg/go-multimov"
	"image"
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
func GetFrameQt(id C.int, frameNum C.int) C.ImageBuffer {

	val, has := QTIds.Load(int(id))

	if !has {
		fmt.Printf("Id doesn't exist")
		return C.ImageBuffer{}
	}

	ext := val.(lazyquicktime.MovieExtractor)

	img, perf, err := ext.ExtractFramePerf(uint64(frameNum))

	fmt.Printf("Read took %g; decode took %g\n", perf.Read.Seconds(), perf.Decode.Seconds())

	if err != nil {
		fmt.Printf("Error extracting image: %s", err.Error())
		return C.ImageBuffer{}
	}

	return imageBufferFromImage(img)

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

func imageBufferFromImage(img image.Image) C.ImageBuffer {
	out := C.ImageBuffer{
		width:  C.int(img.Bounds().Max.X - img.Bounds().Min.X),
		height: C.int(img.Bounds().Max.Y - img.Bounds().Min.Y),
	}

	// TODO.   Almost certainly a more idiomatic way to do this
	ok := false

	//fmt.Printf("Image is type %T\n", img)

	switch t := img.(type) {
	default:
		fmt.Printf("Unexpected type %T\n", t)
	case *image.Gray:
		g := img.(*image.Gray)
		// if( ok ) {
		out.channels = 1
		out.depth = C.IMG_8U
		out.data = C.CBytes(g.Pix)
		// }
		ok = true
	case *image.RGBA:
		g := img.(*image.RGBA)
		// fmt.Printf("%T: %v\n", g, g)

		// if( ok ) {
		out.channels = 4
		out.depth = C.IMG_8U
		out.data = C.CBytes(g.Pix)
		// }
		ok = true
	case *image.NRGBA:
		g := img.(*image.NRGBA)
		// fmt.Printf("%T: %v\n", g, g)

		// if( ok ) {
		out.channels = 4
		out.depth = C.IMG_8U
		out.data = C.CBytes(g.Pix)
		// }
		ok = true
	}

	if ok {
		out.valid = C.uchar(1)
	}
	return out
}
