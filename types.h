#pragma once

typedef struct {
  float duration;
  int num_frames;
  unsigned char valid;
} MovieInfo;


typedef enum {
  IMG_8U = 0,
  IMG_16U,
  IMG_32F
} ImageDepth;

// I wonder about the wisdom of defining a new interchange format
typedef struct  {
  void *data;
  int width, rows;
  int channels;
  ImageDepth depth;
} ImageBuffer;
