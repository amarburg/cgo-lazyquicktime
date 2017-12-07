#pragma once

#include <string>

static char CI_TEST_MOVIE_URL[] = "https://rawdata.oceanobservatories.org/files/RS03ASHS/PN03B/06-CAMHDA301/2016/01/01/CAMHDA301-20160101T000000Z.mov";


// These files are brought in by CMake's ExternalData system and
// installed in the build directory...
static char LOCAL_TEST_MOV[] = CMAKE_BINARY_DIR"/CamHD_Vent_Short.mov";

// This exists as a copy in this repo..
static char LOCAL_TEST_MULTIMOV[] = CMAKE_SOURCE_DIR"/four_mov.json";
