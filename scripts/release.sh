#!/usr/bin/env bash

name=$1
targets_dir="$PWD/$2"
release_dir="$PWD/$3"

echo "start: zip all targets in $targets_dir to $release_dir."
mkdir -p "$release_dir"
cd "$targets_dir" || exit

targets=$(ls "$targets_dir")
for target in $targets
do
    echo "zipping $target..."
    cd "$target" 
    tar -czf "$name-$target.tar.gz" ./* 
    mv "$name-$target.tar.gz" "$release_dir" 
    cd ..
done

echo "done."