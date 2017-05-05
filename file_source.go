package main

import "C"
import (
	"fmt"
	"github.com/amarburg/go-lazyfs"
	"net/url"
	"regexp"
)

func sourceFromCPath(path *C.char) (lazyfs.FileSource, error) {

	goPath := C.GoString(path)

	// TODO.  Additional validatation of path
	if len(goPath) == 0 {
		return nil, fmt.Errorf("Zero length filename: %s", goPath)
	}

	var file lazyfs.FileSource
	var err error

	match, _ := regexp.MatchString("^http", goPath);
	if  match {
		uri, err := url.Parse(goPath)
		file, err = lazyfs.OpenHttpSource(*uri)
		if err != nil {
			return nil, fmt.Errorf("Error opening URL: %s", err.Error())
		}
	} else {
		file, err = lazyfs.OpenLocalFile(C.GoString(path))
		if err != nil {
			return nil, fmt.Errorf("Error opening file: %s", err.Error())
		}
	}

	if file == nil {
		return nil, fmt.Errorf("Error creating file")
	}

	return file, nil
}
