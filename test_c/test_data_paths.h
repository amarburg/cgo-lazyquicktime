#pragma once

#include <string>

static char CI_TEST_MOVIE_URL[] = "https://rawdata.oceanobservatories.org/files/RS03ASHS/PN03B/06-CAMHDA301/2016/01/01/CAMHDA301-20160101T000000Z.mov";


static char LOCAL_TEST_MOV[] = TESTFILE_DIR"/CamHD_Vent_Short.mov";

static char LOCAL_TEST_MULTIMOV[] = TESTFILE_DIR"/multimov/four_mov.json";

static char LOCAL_TEST_FRAMESET[] = TESTFILE_DIR"/frameset/good_frameset.json";
static long unsigned LOCAL_TEST_FRAMESET_FRAMES[] = {1,5,10};
static int LOCAL_TEST_FRAMESET_NUMFRAMES = sizeof(LOCAL_TEST_FRAMESET_FRAMES)/sizeof(long unsigned);
