#!/bin/bash

merge_sort() {
    unsorted=("$@")
    n=${#unsorted[@]}
    if [ $n -le 1 ]; then
        echo "${unsorted[@]}"
        return
    fi

    mid=$((n / 2))
    left=("${unsorted[@]:0:$mid}")
    right=("${unsorted[@]:$mid}")

    left_sorted=($(merge_sort "${left[@]}"))
    right_sorted=($(merge_sort "${right[@]}"))

    merge "${left_sorted[@]}" "${right_sorted[@]}"
}

merge() {
    left=("$@")
    right=("${@:$((${#left[@]} + 1))}")

    i=0
    j=0
    results=()

    while [ $i -lt ${#left[@]} ] && [ $j -lt ${#right[@]} ]; do
        if [ ${left[$i]} -lt ${right[$j]} ]; then
            results+=("${left[$i]}")
            ((i++))
        else
            results+=("${right[$j]}")
            ((j++))
        fi
    done

    if [ $i -lt ${#left[@]} ]; then
        results+=( "${left[@]:$i}" )
    fi
    if [ $j -lt ${#right[@]} ]; then
        results+=( "${right[@]:$j}" )
    fi 

    echo "${results[@]}"
}

unsorted=(3 1 4 1 5 9 2 6 5 3 5)
sorted=($(merge_sort "${unsorted[@]}"))
echo "Sorted: ${sorted[@]}"
