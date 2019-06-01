#!/bin/bash
if [ $# -lt 2 ]; then
  echo "Use: build.sh source_dir git_repo"
  echo "Sample: build.sh /home/my/source git@github.com:leimark/fibonacci.git"
  echo "Or: build.sh /home/my/source https://github.com/leimark/fibonacci.git"
  exit 1
fi

baseProjectDir=$1
git_repo=$2

if [ ! -d $baseProjectDir ]; then
   mkdir $baseProjectDir
fi

cd $baseProjectDir

#clone source from git repo
git clone $git_repo 

if [ $? -ne 0 ]; then
   echo "Clone source from $git_repo ran into problems"
   exit 1
fi

#build the source
cd fibonacci
go build main.go

#Run the unit test
go test -v main.go main_test.go
