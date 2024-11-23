pub fn merge_sort(unsorted: Vec<i32>) -> Vec<i32> {
    let n = unsorted.len();
    if n <= 1 {
        return unsorted;
    }

    let mid = n / 2;
    let left = unsorted[0..mid].to_vec();
    let right = unsorted[mid..n].to_vec();

    let sorted_left = merge_sort(left);
    let sorted_right = merge_sort(right);

    merge(sorted_left, sorted_right)
}

fn merge(left: Vec<i32>, right: Vec<i32>) -> Vec<i32> {
    let (mut i, mut j) = (0, 0);
    let mut sorted = vec![];

    while i < left.len() && j < right.len() {
        if left[i] < right[j] {
            sorted.push(left[i]);
            i += 1;
        } else {
            sorted.push(right[j]);
            j += 1;
        }
    }

    if i < left.len() {
        sorted.extend_from_slice(&left[i..]);
    }

    if j < right.len() {
        sorted.extend_from_slice(&right[i..]);
    }

    sorted
}
