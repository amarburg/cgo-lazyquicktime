
#include <gtest/gtest.h>

#include "validate_camhd_frame.h"

void validateCamHDFrame( ImageBuffer &img )
{
  ASSERT_EQ( img.valid, 1 );

  ASSERT_EQ( img.width, 1920 );
  ASSERT_EQ( img.height, 1080 );
  ASSERT_TRUE( img.data != nullptr );
  ASSERT_EQ( img.channels, 4 );
  ASSERT_EQ( img.depth, IMG_8U );
}
