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
TEST( FrameSource, FromFrameSet )
{
  int fd = OpenFrameSource( LOCAL_TEST_FRAMESET );

  // TODO: How to set this path..?
  ImageBuffer img = FrameSourceNext( fd );
  validateCamHDFrame( img );

  CloseQt( fd );

  free( img.data );
}
