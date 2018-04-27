#include <gtest/gtest.h>
#include <string>
#include <iostream>

#include "liblazyquicktime.h"
#include "types.h"

#include "test_data_paths.h"

using std::endl;
using std::cout;

TEST( MovieInfo, FirstTest )
{
  char url[] = "";
  MovieInfo info;

  // Should fail
  ASSERT_EQ( MovInfo(url, &info), -1);

  ASSERT_EQ( info.valid, 0 );
}

// Create from a real movie
TEST( MovieInfo, MovieInfoLocalMov )
{

  // TODO: How to set this path..?
  MovieInfo info;
  ASSERT_EQ( MovInfo( LOCAL_TEST_MOV, &info ), 0 );

  ASSERT_EQ( info.valid, 1 );

  // These values are known apriori
  ASSERT_EQ( info.num_frames, 42 );
  ASSERT_FLOAT_EQ( info.duration, 1.4014001 );
}

// Create from a multimov
TEST( MovieInfo, MovieInfoLocalMultimov )
{
  // TODO: How to set this path..?
  MovieInfo info;
  ASSERT_EQ( MovInfo( LOCAL_TEST_MULTIMOV, &info ), 0 );

  ASSERT_EQ( info.valid, 1 );

  // These values are known apriori
  // ASSERT_EQ( info.num_frames, 31 );
  // ASSERT_FLOAT_EQ( info.duration, 1.0343666 );
}
