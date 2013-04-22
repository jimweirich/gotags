task :default => :build

task :printenv do
  puts "export GOPATH=\"#{ENV['PWD']}\""
end

task :env do
  ENV['GOPATH'] = ENV['PWD']
end

task :build => :env do
  sh "go install onestepback.org/gotags"
end

task :run, [:c] => :build do |t, args|
  c = args.c || 10
  puts "Concurrency is #{c}"
  sh "time -p " +
    "bin/gotags " +
    "-c #{c} " +
    "testdata"
end

BINDIR = "#{ENV['HOME']}/local/bin"

task :deploy => :build do
  cp "bin/gotags", BINDIR
end
