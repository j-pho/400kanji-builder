#!/bin/zsh

# need system audio default to be set to something that knows japanese...
# e.g., add the japanese siri voice and set the default to that

# example: ./single-word.sh 0581 開催

SENT=$2
NN=$1

say -o tmp.m4a ${SENT}
ffmpeg -i tmp.m4a k-${NN}_${SENT}.mp3
rm tmp.m4a
