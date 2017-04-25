#include <gtest/gtest.h>
#include <string>
#include <iostream>

#include "liblazyquicktime.h"
#include "types.h"

using std::endl;
using std::cout;

TEST( MovieInfo, FirstTest )
{
  char url[] = "";
  MovieInfo info = MovInfo(url);

  ASSERT_EQ( info.valid, 0 );
}

// Test against a real movie
TEST( MovieInfo, LazyFsTestFiles )
{
  char url[] =  "../go-lazyfs-testfiles/CamHD_Vent_Short.mov";

  // TODO: How to set this path..?
  MovieInfo info = MovInfo(url);

  ASSERT_EQ( info.valid, 1 );

  // These values are known apriori
  ASSERT_EQ( info.num_frames, 31 );
  ASSERT_FLOAT_EQ( info.duration, 1.0343666 );
}
