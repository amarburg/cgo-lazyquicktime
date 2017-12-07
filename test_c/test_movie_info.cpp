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
  MovieInfo info = MovInfo(url);

  ASSERT_EQ( info.valid, 0 );
}

// Test against a real movie
TEST( MovieInfo, MovieInfoLocalMov )
{

  // TODO: How to set this path..?
  MovieInfo info = MovInfo( LOCAL_TEST_MOV );

  ASSERT_EQ( info.valid, 1 );

  // These values are known apriori
  ASSERT_EQ( info.num_frames, 31 );
  ASSERT_FLOAT_EQ( info.duration, 1.0343666 );
}

// Test against a real movie
TEST( MovieInfo, MovieInfoLocalMultimov )
{

  // TODO: How to set this path..?
  MovieInfo info = MovInfo( LOCAL_TEST_MULTIMOV );

  ASSERT_EQ( info.valid, 1 );

  // These values are known apriori
  // ASSERT_EQ( info.num_frames, 31 );
  // ASSERT_FLOAT_EQ( info.duration, 1.0343666 );
}
