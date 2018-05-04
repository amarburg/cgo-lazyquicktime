#include <gtest/gtest.h>
#include <string>
#include <iostream>

#include "libmovieextractor.h"
#include "types.h"

#include "test_data_paths.h"
#include "validate_camhd_frame.h"

using std::endl;
using std::cout;


// Test against a real movie
TEST( GetFrame, GetFrameQtFromLocalFile )
{
  int fd = OpenMovie( LOCAL_TEST_MOV );

  // TODO: How to set this path..?
  ImageBuffer img;
  ASSERT_EQ( GetMovieFrame( fd, 1, &img ), 0 );
  validateCamHDFrame( img );

  ASSERT_EQ( GetMovieFrame( fd, 5, &img ), 0 );
  validateCamHDFrame( img );

  ASSERT_EQ( GetMovieFrame( fd, 10, &img ), 0 );
  validateCamHDFrame( img );

  MovieInfo info;
  ASSERT_EQ( GetMovieInfo( fd, &info ), 0 );

  CloseMovie( fd );

  free( img.data );
}
