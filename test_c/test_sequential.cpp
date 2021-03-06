#include <gtest/gtest.h>
#include <string>
#include <iostream>

#include "libmovieset.h"
#include "types.h"

#include "test_data_paths.h"
#include "validate_camhd_frame.h"

using std::endl;
using std::cout;


// Test against a real movie
TEST( FrameSource, FromFrameSet )
{
  int fd = OpenSequential( LOCAL_TEST_FRAMESET );

  // TODO: How to set this path..?
  ImageBuffer img;

  for( int i = 0; i < LOCAL_TEST_FRAMESET_NUMFRAMES; i++ ) {
    int err = SequentialNext( fd, &img );

    ASSERT_EQ( err, LOCAL_TEST_FRAMESET_FRAMES[i] );

    validateCamHDFrame( img );
  }

  CloseSequential( fd );

  free( img.data );
}
