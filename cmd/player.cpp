
#include <iostream>

#include <stdlib.h>
#include <opencv2/highgui.hpp>
#include <opencv2/imgproc.hpp>

#include <boost/program_options.hpp>
namespace po = boost::program_options;
using namespace std;


#include "liblazyquicktime.h"

int main( int argc, char **argv )
{
  // Declare the supported options.
  po::options_description desc("Allowed options");
  desc.add_options()
      ("help", "produce help message")
      ("input-file", po::value<string>(), "Name of input file");

  po::positional_options_description p;
  p.add("input-file", -1);

  po::variables_map vm;
  po::store(po::command_line_parser(argc, argv).
          options(desc).positional(p).run(), vm);
  po::notify(vm);

  if (vm.count("help")) {
      cout << desc << endl;
      return 1;
  }

  if (vm.count("input-file") != 1 ) {
      cout << "Program takes exactly one input file" << endl;
      return 1;
  }

  string inputFile(  vm["input-file"].as<string>() );
  cout << "Opening input file: " << inputFile << endl;

  std::vector<char> chars(inputFile.c_str(), inputFile.c_str() + inputFile.size() + 1u);
  int id = OpenFrameSource( &chars[0] );

  int count = 0;
  ImageBuffer buffer;
  while( FrameSourceNext(id, &buffer) > 0 && count < 200 ) {
    cout << "Read frame " << count << " as " << buffer.width << " x " << buffer.height << endl;

    cv::Mat mat( buffer.height, buffer.width, CV_MAKETYPE(buffer.depth, buffer.channels), buffer.data );

    cv::cvtColor( mat, mat, cv::COLOR_BGR2RGB );

    cv::imshow("Window", mat);
    cv::waitKey(1);

      ++count;
  }


  CloseFrameSource( id );

  return 0;
}
