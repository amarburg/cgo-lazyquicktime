#include <gtest/gtest.h>
#include <string>
#include <iostream>

#include "liblazyquicktime.h"
#include "types.h"

using std::endl;
using std::cout;

TEST( GetFrame, FirstTest )
{
  char url[] = "";
  ImageBuffer img = GetFrame(url, 1);

  ASSERT_EQ( img.valid, 0 );
}

// TEST( GetFrame, LocalFileZeroFrame )
// {
//   // TODO: How to set this path..?
//   ImageBuffer img = GetFrame( localFilePath, 0 );
//   ASSERT_EQ( img.valid, 0 );
// }

// Test against a real movie
TEST( GetFrame, LocalFileValidFrame )
{
  // TODO: How to set this path..?
  ImageBuffer img = GetFrame( LOCAL_TEST_MOV, 1 );
  ASSERT_EQ( img.valid, 1 );

  ASSERT_EQ( img.width, 1920 );
  ASSERT_EQ( img.height, 1080 );
  ASSERT_TRUE( img.data != nullptr );
  ASSERT_EQ( img.channels, 4 );
  ASSERT_EQ( img.depth, IMG_8U );

  free( img.data );
}
