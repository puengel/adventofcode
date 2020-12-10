#!/bin/bash

adapters=(0)

for line in  $(<input.txt); do
    adapters+=($line)
done


# for a in ${adapters[@]}; do
    # echo $a
# done

# https://stackoverflow.com/questions/7442417/how-to-sort-an-array-in-bash

IFS=$'\n' sorted=($(sort -n <<<"${adapters[*]}")); unset IFS

sorted+=($((sorted[-1]+3)))

let before=0 diff1=0 diff3=0
for n in ${sorted[@]}; do
    if [[ $((n-before)) -eq 1 ]]
    then
        let diff1+=1
    elif [[ $((n-before)) -eq 3 ]]
    then
        let diff3+=1
    fi
    before=$n
done

echo "Part 1: $(($diff1*$diff3))"

# https://www.geeksforgeeks.org/count-ways-reach-nth-stair/

res=(1)
for i in $(seq 1 $((${#sorted[@]}-1))); do
    ans=0
    for j in $(seq 0 $i); do
        if [[ $((sorted[j]+3)) -ge ${sorted[$i]} ]]
        then
            let ans+=res[j]
        fi
    done
    res+=($ans)
done

echo "Part 2: ${res[-1]}"