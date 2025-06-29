#!/bin/zsh

# need system audio default to be set to something that knows japanese...
# e.g., add the japanese siri voice and set the default to that

# INP should be one word per line

INP='sentences_for_conversion.txt'
PREFIX='4c-s'
declare -i i=0

while read SENT; do
  i+=1
  DG=$(printf "%04d" i)
  say -o ${DG}_${SENT}.m4a ${SENT}
  ffmpeg -i ${DG}_${SENT}.m4a ${PREFIX}-${SENT}.mp3
  rm ${DG}_${SENT}.m4a
  mv ${PREFIX}-${SENT}.mp3 ../collection.media/${PREFIX}-${SENT}.mp3
done < ${INP}
