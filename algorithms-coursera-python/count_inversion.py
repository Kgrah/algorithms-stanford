def merge(left, right, inv):
    i, j = 0, 0
    sorted_list = []

    while i < len(left) and j < len(right):
        if left[i] <= right[j]:
            sorted_list.append(left[i])
            i += 1
        else:
            sorted_list.append(right[j])
            j += 1
            # Count inversions: all remaining elements in left are greater than right[j]
            inv += len(left) - i

    # Append any remaining elements
    sorted_list.extend(left[i:])
    sorted_list.extend(right[j:])

    return sorted_list, inv


def merge_sort(unsorted, inv_count):
    if len(unsorted) <= 1:
        return unsorted, inv_count

    mid = len(unsorted) // 2
    left, right = unsorted[:mid], unsorted[mid:]

    left_sorted, inv_count = merge_sort(left, inv_count)
    right_sorted, inv_count = merge_sort(right, inv_count)

    return merge(left_sorted, right_sorted, inv_count)


def count_inversions(nums):
    _, inv_count = merge_sort(nums, 0)
    return inv_count


def main():
    nums = [6, 5, 4, 3, 2, 1]
    result = count_inversions(nums)
    print("Number of inversions:", result)


if __name__ == "__main__":
    main()
