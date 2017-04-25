package main

// #include "types.h"
import "C"
import (
    "github.com/amarburg/go-lazyquicktime"
    "github.com/amarburg/go-lazyfs"
    "fmt"
    "net/url"
    "regexp"
)

//export MovInfo
func MovInfo( path *C.char ) C.MovieInfo {


  goPath := C.GoString(path)

  // TODO.  Validate path
  if( len(goPath) == 0 ) {
    fmt.Printf( "Zero length filename: %s", goPath )
    return C.MovieInfo{}
  }

  if match,_ := regexp.MatchString( "/^{http}/", goPath ); match {
    return MovInfoFromURL( path )
  }

  return MovInfoFromFile( path )
}

//export MovInfoFromFile
func MovInfoFromFile( path *C.char ) C.MovieInfo {

  file, err := lazyfs.OpenLocalFile( C.GoString(path) )
  if( err != nil ) {
    fmt.Printf( "Error opening file: %s", err.Error() )
    return C.MovieInfo{}
  }

  qtInfo,err := lazyquicktime.LoadMovMetadata( file )

  if( err != nil ) {
    fmt.Printf( "Error getting metadata: %s", err.Error() )
    return C.MovieInfo{}
  }

  return qtInfoToMovieInfo( qtInfo )
}

//export MovInfoFromURL
func MovInfoFromURL( path *C.char ) C.MovieInfo {

  uri,_ := url.Parse( C.GoString(path) )

  file, err := lazyfs.OpenHttpSource( *uri )
  if( err != nil ) {
    fmt.Printf( "Error opening file: %s", err.Error() )
    return C.MovieInfo{}
  }

  qtInfo,err := lazyquicktime.LoadMovMetadata( file )

  if( err != nil ) {
    fmt.Printf( "Error getting metadata: %s", err.Error() )
    return C.MovieInfo{}
  }

  return qtInfoToMovieInfo( qtInfo )
}



func qtInfoToMovieInfo( qtInfo *lazyquicktime.LazyQuicktime ) C.MovieInfo {
  return C.MovieInfo {
    duration:   C.float(qtInfo.Duration()),
    num_frames: C.int(qtInfo.NumFrames()),
    valid: 1,
  }
}
