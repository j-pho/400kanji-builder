#!/bin/zsh

# need system audio default to be set to something that knows japanese...
# e.g., add the japanese siri voice and set the default to that

# INP should be one word per line

INP='merged-kanji-cards.txt'
OUTP='./trimmed-cards-out'

split -l 1000 ${INP}
files=(x*)
for item in ${files[*]}
do
  mv $item "${OUTP}/cards-${item}.txt"
done
