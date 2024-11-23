def main():
    print("start")

    numbers = [7, 2, 9, 4, 1, 6, 3, 0, 8, 5]

    print(merge_sort(numbers))

def merge_sort(unsorted):
    if len(unsorted) <= 1:
        return unsorted

    n = len(unsorted)
    mid = n // 2

    left = unsorted[0:mid]
    right = unsorted[mid:]

    left_sorted = merge_sort(left)
    right_sorted = merge_sort(right)

    return merge(left_sorted, right_sorted)

def merge(left, right):
    i, j = 0, 0
    sorted = []

    while i < len(left) and j < len(right):
        if left[i] < right[j]:
            sorted.append(left[i])
            i += 1
        else:
            sorted.append(right[j])
            j += 1
            
    if i < len(left):
        sorted.extend(left[i:])
        
    if j < len(right):
        sorted.extend(right[j:])

    return sorted

if __name__ == "__main__":
    main()