package main

// #include "types.h"
import "C"

import (
	"fmt"
	"image"
)

//TODO.  Lots of repetition between this API and movie_info.go
// room for DRY?

//TODO.  is Printf on error the best way to get error information out?


//export GetFrame
func GetFrame(path *C.char, frameNum C.int) C.ImageBuffer {

	goPath := C.GoString(path)
	ext,err := movieExtractorFromPath( goPath )

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
		width:   C.int(img.Bounds().Max.X - img.Bounds().Min.X),
		height:  C.int(img.Bounds().Max.Y - img.Bounds().Min.Y),
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
