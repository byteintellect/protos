#!/bin/bash

export SOURCE_CODE_PATH=${HOME}/go_proto_files;
export PWD=$(pwd);
tag='';

build() {
  [ ! -d "${SOURCE_CODE_PATH}" ] && mkdir -p "${SOURCE_CODE_PATH}";
  # shellcheck disable=SC2046
  docker run --rm --user $(id -u):$(id -g) -v"${PWD}":"${PWD}" -v"${SOURCE_CODE_PATH}":"${SOURCE_CODE_PATH}" -w"${SOURCE_CODE_PATH}" thethingsindustries/protoc:latest --proto_path="${PWD}" --go_out="${SOURCE_CODE_PATH}" --go-grpc_out="${SOURCE_CODE_PATH}" --grpc-gateway_out="logtostderr=true,allow_delete_body=true:." -I/usr/include/github.com/gogo/protobuf "${PWD}"/core_v1/core_v1.proto "${PWD}"/users_v1/users_v1.proto;
  cd "${SOURCE_CODE_PATH}"/"${GO_CODE_PATH}" || exit;
  go mod init "${GO_CODE_PATH}";
  go mod tidy;
}

fill_git_tag() {
  tag=$(git describe --always --tags || echo pre-commit);
}

push() {
  cd "${SOURCE_CODE_PATH}"/"${GO_CODE_PATH}" || exit;
  git init --initial-branch=master;
  git remote add origin "${ORIGIN_PATH}";
  git config user.name kutty-kumar;
  git config user.email kumar.varalakshmi@outlook.com;
  git add .;
  git commit -m "compiled protobuf code";
  git fetch --all;
  git branch --set-upstream-to=origin/master master;
  git pull --rebase;
  git tag "${1}";
  git push --atomic origin HEAD "${1}";
}

build_and_push(){
  fill_git_tag;
  echo "${tag}";
  build;
  push "${tag}";
}