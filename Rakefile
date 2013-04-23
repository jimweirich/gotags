require 'rake/clean'

CLOBBER.include("bin", "TAGS", "dist")

PROG = "bin/gotags"

VERSION_DATA = `egrep "^ *var +version +=" src/onestepback.org/gotags/gotags.go`
if /"([0-9.]+)"/ =~ VERSION_DATA
  VERSION = $1
else
  VERSION = 'unknown'
end

task :default => :check

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

TEST_FILES = FileList['testdata/**/*.rb']

file "TAGS" => ["bin/gotags"] + TEST_FILES do
  sh "bin/gotags testdata"
end

desc "Check that we produce a compatible TAGS file"
task :check => ["TAGS"] do
  sh "diff -u testdata/expected_tags.out TAGS"
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

desc "Run tags on all gems in Gem HOME"
task :taghome do
  sh "time bin/gotags '#{`gem env gemhome`.strip}'"
end

directory "dist"

PLATFORM = `uname -s -m`.strip.downcase.gsub(/[^a-z0-9_]+/, '-')

EXE_FILE = "gotags-#{VERSION}-#{PLATFORM}"
TAR_FILE = "gotags-#{VERSION}-#{PLATFORM}.tgz"

EXE_PATH = "dist/#{EXE_FILE}"
TAR_PATH = "dist/#{TAR_FILE}"

desc "Make a platform binary executable"
task :executable => EXE_PATH

desc "Make a distribution zip file"
task :tar => TAR_PATH

task :need_version do
  if VERSION == 'unknown'
    puts "Unknown version"
    exit(-1)
  end
  this_version = `bin/gotags -v`.strip
  if VERSION != this_version
    puts "Mismatched version (expected #{this_version}, got #{VERSION})"
    puts "Run 'rake build' first"
    exit(-1)
  end
end

file EXE_PATH => [:build, :need_version, "dist"] do
  cp "bin/gotags", EXE_PATH
end

file TAR_PATH => [EXE_PATH] do
  Dir.chdir("dist") do
    sh "tar cvfz #{TAR_FILE} #{EXE_FILE}"
  end
end

task :upload => TAR_PATH do
  sh "scp #{TAR_PATH} linode:sites/onestepback.org/download"
end
