
task :default => :test


task :test => :build_so do
  mkdir "build" unless FileTest.directory?("build")
  Dir.chdir "build" do
    sh "cmake .."
    sh "make"
  end

  sh "build/cgo-lazyquicktime-test"
end

task :build_so do
  sh "go build -buildmode=c-shared -o liblazyquicktime.so"
end
