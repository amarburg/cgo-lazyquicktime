
require 'os'
require 'pathname'

task :default => :build_so

task :build_so do
  sh "go build -buildmode=c-shared -o liblazyquicktime.so"

  ## If OSX, set library to identify itself with absolute path
  if OS.mac?
    ldPath = Pathname.new(__FILE__).parent.join('liblazyquicktime.so')
    sh "install_name_tool -id #{ldPath} liblazyquicktime.so"
  end
end

task :test => :build_so do
  Dir.chdir "test_c" do
    mkdir "build" unless FileTest.directory?("build")
    Dir.chdir "build" do
      sh "cmake .."
      sh "make gtest_ext" unless FileTest.exists?("gtest/include/gtest.h")
      sh "make"

      env = "DYLD_LIBRARY_PATH=gtest/lib" if OS.mac?

      sh "#{env} make CTEST_OUTPUT_ON_FAILURE=TRUE test"
    end
  end
end
