#include <gtest/gtest.h>

#include "liblazyquicktime.h"


TEST( TestSum, ConstantAddition )
{
    ASSERT_EQ( Sum(2, 40), 42 );
}
