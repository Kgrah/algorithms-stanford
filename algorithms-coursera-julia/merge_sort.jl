function merge_sort(unsorted)
    if length(unsorted) == 1
        return unsorted
    end

    n = length(unsorted)
    mid = div(n, 2)

    left = unsorted[1:mid]
    right = unsorted[mid+1:end]

    left_sorted = merge_sort(left)
    right_sorted = merge_sort(right)

    return merge(left_sorted, right_sorted)
end

function merge(left, right)
    i, j = 1, 1
    sorted = []
    while i <= length(left) && j <= length(right)
        if left[i] < right[j]
            push!(sorted, left[i])
            i += 1
        else
            push!(sorted, right[j])
            j += 1
        end
    end

    append!(sorted, left[i:end])
    append!(sorted, right[j:end])

    return sorted
end

nums = [6,5,4,3,2,1]

sorted = merge_sort(nums)
println(sorted)