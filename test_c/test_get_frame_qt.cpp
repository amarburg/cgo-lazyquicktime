#include <gtest/gtest.h>
#include <string>
#include <iostream>

#include "liblazyquicktime.h"
#include "types.h"

#include "test_data_paths.h"

using std::endl;
using std::cout;

void validateCamHDFrame( ImageBuffer &img );

// Test against a real movie
TEST( GetFrame, GetFrameQtFromLocalFile )
{
  int fd = OpenQt( LOCAL_TEST_MOV );

  // TODO: How to set this path..?
  ImageBuffer img = GetFrameQt( fd, 1 );
  validateCamHDFrame( img );

  img = GetFrameQt( fd, 5 );
  validateCamHDFrame( img );

  img = GetFrameQt( fd, 10 );
  validateCamHDFrame( img );

  MovieInfo info = GetMovieInfoQt( fd );

  CloseQt( fd );

  free( img.data );
}
//
// // Test against a real movie
// TEST( GetFrame, GetFrameFromHTTPFile )
// {
//   // TODO: How to set this path..?
//   ImageBuffer img = GetFrame( CI_TEST_MOVIE_URL, 1 );
//
//   validateCamHDFrame( img );
//
//   free( img.data );
// }
//
//
// // Test against a multimov
// TEST( GetFrame, GetFrameFromLocalMultimov )
// {
//   // TODO: How to set this path..?
//   ImageBuffer img = GetFrame( LOCAL_TEST_MULTIMOV, 1 );
//
//   validateCamHDFrame( img );
//
//   free( img.data );
// }
