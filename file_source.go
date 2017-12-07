package main

import "C"
import (
	"fmt"
	"github.com/amarburg/go-lazyfs"
	"github.com/amarburg/go-lazyquicktime"
	"github.com/amarburg/go-multimov"
	"net/url"
	"path/filepath"
	"regexp"
)

func sourceFromPath(path string) (lazyfs.FileSource, error) {

	// TODO.  Additional validatation of path
	if len(path) == 0 {
		return nil, fmt.Errorf("Zero length filename: %s", path)
	}

	var file lazyfs.FileSource
	var err error

	match, _ := regexp.MatchString("^http", path)
	if match {
		uri, err := url.Parse(path)
		file, err = lazyfs.OpenHttpSource(*uri)
		if err != nil {
			return nil, fmt.Errorf("Error opening URL: %s", err.Error())
		}
	} else {
		file, err = lazyfs.OpenLocalFile(path)
		if err != nil {
			return nil, fmt.Errorf("Error opening file: %s", err.Error())
		}
	}

	if file == nil {
		return nil, fmt.Errorf("Error creating file")
	}

	return file, nil
}

func movieExtractorFromPath(path string) (lazyquicktime.MovieExtractor, error) {

	if filepath.Ext(path) == ".mov" {

		file, err := sourceFromPath(path)

		if err != nil {
			return nil, err
		}

		qtInfo, err := lazyquicktime.LoadMovMetadata(file)

		if err != nil {
			return nil, err
		}

		return qtInfo, nil

	} else if filepath.Ext(path) == ".json" {

		fmt.Println(path)

		mm, err := multimov.LoadMultiMov(path)
		if err != nil {
			return nil, err
		}

		return mm, nil

	}

	return nil, fmt.Errorf("Can't make a movie extractor from file %s", path)

}
