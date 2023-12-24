#!/usr/bin/env bash

BIN_NAME="server"

script_dir=$(dirname "$(readlink -f "$0")")

current_dir="$PWD"

base_dir=$script_dir

version=$(git describe --tags --always)


CreateDirectory() {
    mkdir -p $base_dir/output/bin $base_dir/output/config
}

CopyFile(){
    cp -rf $script_dir/script/bootstrap.sh $base_dir/output/
    cp -rf $script_dir/config/* $base_dir/output/config

    chmod +x $base_dir/output/bootstrap.sh
}


GoBuild(){
    GO111MODULE=on
    go mod tidy
    go mod download
    go build -ldflags "-X main.Version=$version"  -o $base_dir/output/bin/$BIN_NAME $script_dir/cmd/server/main.go
}

GenerateVersion(){
    echo $1 >> $base_dir/output/version
}

main(){
    echo "Git version:"$version

    echo "output director:"$base_dir

    CreateDirectory

    CopyFile

    GenerateVersion $version

    GoBuild
}

main