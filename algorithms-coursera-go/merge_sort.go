package main

func mergeSort(unsorted []int) []int {
	if len(unsorted) <= 1 {
		return unsorted
	}

	n := len(unsorted)

	mid := n / 2

	left := unsorted[0:mid]
	right := unsorted[mid:]

	leftSorted := mergeSort(left)
	rightSorted := mergeSort(right)

	return merge(leftSorted, rightSorted)
}

func merge(left, right []int) []int {
	i, j := 0, 0
	var sorted []int

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			sorted = append(sorted, left[i])
			i++
		} else {
			sorted = append(sorted, right[j])
			j++
		}
	}

	if i < len(left) {
		sorted = append(sorted, left[i:]...)
	}

	if j < len(right) {
		sorted = append(sorted, right[j:]...)
	}

	return sorted
}
