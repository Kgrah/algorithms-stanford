mod merge_sort;

fn main() {
    let unsorted = vec![7, 2, 9, 4, 1, 6, 3, 0, 8, 5];

    let sorted = merge_sort::merge_sort(unsorted.clone());

    println!("len of sorted: {}", sorted.len());

    for n in sorted {
        println!("{}", n);
    }
}
