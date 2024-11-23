pub fn count_inversion(nums: Vec<i32>) -> i32 {
    let (_, inv_count) = merge_sort(nums);

    inv_count
}

fn merge_sort(unsorted: Vec<i32>, inv_count: i32) -> (Vec<i32>, i32) {
    if unsorted.len() == 1 {
        return (unsorted, inv_count);
    }

    let n = unsorted.len();
    let mid = n / 2;

    let (left, right) = (unsorted[0..mid].to_vec(), unsorted[mid..n].to_vec());

    let (left_sorted, inv_left) = merge_sort(left, inv_count);
    let (right_sorted, inv_right) = merge_sort(right, inv_left);

    merge(left_sorted, right_sorted, inv_right)
}

fn merge(left: Vec<i32>, right: Vec<i32>, inv: i32) -> (Vec<i32>, i32) {
    let (mut i, mut j) = (0, 0);

    let mut sorted = vec![];
    while i < left.len() && j < right.len() {
        if left[i] < right[j] {
            sorted.push(left[i]);
            i += 1;
        } else {
            sorted.push(right[j]);
            j += 1;
            inv += (left.len() - i) as i32;
        }
    }

    sorted.extend(&left[i..]);
    sorted.extend(&right[j..]);

    (sorted, inv)
}
