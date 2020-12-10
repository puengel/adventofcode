#!/bin/bash

adapters=()

while read line; do adapters+=($line); done < input.txt

# for a in ${adapters[@]}; do
    # echo $a
# done

# https://stackoverflow.com/questions/7442417/how-to-sort-an-array-in-bash

IFS=$'\n' sorted=($(sort -n <<<"${adapters[*]}")); unset IFS

echo ${sorted[@]}