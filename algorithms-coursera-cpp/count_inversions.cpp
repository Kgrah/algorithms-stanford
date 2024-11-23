#include <vector>;
#include <tuple>;
using namespace std;

int count_inversions(vector<int> nums) {
    auto [_, inv_count] = merge_sort(nums, 0);

    return inv_count;
}

tuple<vector<int>, int> merge_sort(vector<int> unsorted, int inv_count) {
    vector<int> sorted;

    if (unsorted.size() == 1) {
        return make_tuple(unsorted, inv_count);
    }

    int n = unsorted.size();
    int mid = n / 2;

    vector<int> left(unsorted.begin(), unsorted.begin() + mid);
    vector<int> right(unsorted.begin() + mid, unsorted.end());

    auto [left_sorted, inv_left] = merge_sort(left, inv_count);
    auto [right_sorted, inv_right] = merge_sort(right, inv_left);



    return merge(left_sorted, right_sorted, inv_right);
}

tuple<vector<int>, int> merge(vector<int> left, vector<int> right, int inv) {
    vector<int> merged;
    int i = 0;
    int j = 0;

    while (i < left.size() && j < right.size()) {
        if (left[i] < right[j]) {
            merged.push_back(left[i]);
            i++;
        } else {
            merged.push_back(right[j]);
            j++;
            inv += left.size() - i;
        }
    }

    merged.insert(merged.end(), left.begin() + i, left.end());
    merged.insert(merged.end(), right.begin() + j, right.end());

    return make_tuple(merged, inv);
}

