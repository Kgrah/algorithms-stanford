package main

func countInversion(nums []int) int {
	_, count := mergeSortInv(nums, 0)
	return count
}

func mergeSortInv(unsorted []int, invCount int) ([]int, int) {
	if len(unsorted) == 1 {
		return unsorted, invCount
	}

	n := len(unsorted)
	mid := n / 2

	left, right := unsorted[:mid], unsorted[mid:]

	var leftSorted []int
	leftSorted, invCount = mergeSortInv(left, invCount)

	var rightSorted []int
	rightSorted, invCount = mergeSortInv(right, invCount)

	return mergeInv(leftSorted, rightSorted, invCount)
}

func mergeInv(left, right []int, inv int) ([]int, int) {
	i, j := 0, 0

	var sorted []int
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			sorted = append(sorted, left[i])
			i++
		} else {
			sorted = append(sorted, right[j])
			j++
			inv += len(left) - i
		}
	}

	sorted = append(sorted, left[i:]...)
	sorted = append(sorted, right[j:]...)

	return sorted, inv
}
