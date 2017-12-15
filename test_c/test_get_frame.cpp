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
  ImageBuffer img;
  ASSERT_EQ( GetFrameQt( fd, 1, &img ), 0 );
  validateCamHDFrame( img );

  ASSERT_EQ( GetFrameQt( fd, 5, &img ), 0 );
  validateCamHDFrame( img );

  ASSERT_EQ( GetFrameQt( fd, 10, &img ), 0 );
  validateCamHDFrame( img );

  MovieInfo info = GetMovieInfoQt( fd );

  CloseQt( fd );

  free( img.data );
}
