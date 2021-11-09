#!/bin/bash

export SOURCE_CODE_PATH=${HOME}/go_proto_files;
export PWD=$(pwd);
tag='';
SOURCE_CODE_PATH=gen/out;

pre_build(){
  if [ -d $SOURCE_CODE_PATH ]; then rm -Rf gen; fi
  for dir in */; do
    cd $dir;
    echo "buf mod initializing in $dir";
    buf mod update;
    echo "done updating buf mod in $dir";
    cd ..;
  done
}

build() {
  # shellcheck disable=SC2046
  echo "${PWD}"
  buf generate -v --include-imports;
  cd "${SOURCE_CODE_PATH}" || exit;
  go mod init "${GO_CODE_PATH}";
  go mod tidy;
}

fill_git_tag() {
  tag=$(git describe --always --tags || echo pre-commit);
}

push() {
  git init --initial-branch="latest-${1}";
  git remote add origin "${ORIGIN_PATH}";
  git config user.name "Kumar D";
  git config user.email kumar.d@byteintellect.com;
  git add .;
  git commit -m "compiled protobuf code";
  git fetch --all;
  git tag "${1}";
  git push -u --atomic origin "latest-${1}" "${1}";
}

build_and_push(){
  fill_git_tag;
  echo "${tag}";
  pre_build;
  build;
  push "${tag}";
}

compile(){
  fill_git_tag;
  echo "${tag}";
  pre_build;
  build;
}
