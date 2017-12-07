#pragma once

typedef struct MovieInfo {
  float duration;
  int num_frames;
  unsigned char valid;
} MovieInfo;


// Aligns with OpenCV convention
typedef enum {
  IMG_8U = 0,
  IMG_16U = 2,
  IMG_32F = 5
} ImageDepth;

// I wonder about the wisdom of defining a new interchange format
typedef struct ImageBuffer {
  void *data;
  int width, height;
  int channels;
  ImageDepth depth;
  unsigned char valid;
} ImageBuffer;
