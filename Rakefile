
task :default => :build_so

task :build_so do
  sh "go build -buildmode=c-shared -o liblazyquicktime.so"
end

task :test => :build_so do
  Dir.chdir "test_c" do
    mkdir "build" unless FileTest.directory?("build")
    Dir.chdir "build" do
      sh "cmake .."
      sh "make gtest_ext" unless FileTest.exists?("gtest/include/gtest.h")
      sh "make"
      sh "make CTEST_OUTPUT_ON_FAILURE=TRUE test"
    end
  end
end
