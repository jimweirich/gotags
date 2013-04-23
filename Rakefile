require 'rake/clean'

CLOBBER.include("bin", "TAGS")

task :default => :build

desc "Print the expected GOPATH"
task :printenv do
  puts "export GOPATH=\"#{ENV['PWD']}\""
end

desc "Set the GOPATH environment variable"
task :env do
  ENV['GOPATH'] = ENV['PWD']
end

desc "Build the gotags program"
task :build => :env do
  sh "go install onestepback.org/gotags"
end

desc "Run gotags on the test data"
task :run => :build do
  sh "time -p bin/gotags testdata"
end

file "bin/gotags" => :build

file "TAGS" => ["bin/gotags"] do
  sh "bin/gotags testdata"
end

desc "Check that we produce a compatible TAGS file"
task :check => ["TAGS"] do
  sh "diff -u TAGS testdata/expected_tags.out"
end

namespace "check" do
  task :update => ["TAGS"] do
    cp "TAGS", "testdata/expected_tags.out"
  end
end

BINDIR = "#{ENV['HOME']}/local/bin"

desc "Deploy the program to a bindir"
task :deploy, [:bindir] => :build do |t, args|
  bindir = args[:bindir] || BINDIR
  mkdir_p bindir unless File.exist?(bindir)
  cp "bin/gotags", bindir
end
