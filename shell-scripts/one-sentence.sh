#!/bin/zsh

# need system audio default to be set to something that knows japanese...
# e.g., add the japanese siri voice and set the default to that

# example: ./one-sentence.sh 0563 当初の計画とは違う方向に進みました。

SENT=$2
NN=$1

say -o tmp.m4a ${SENT}
ffmpeg -i tmp.m4a s-${NN}_${SENT}.mp3
rm tmp.m4a
