#include <iostream>
#include <vector>
#include <span>
using namespace std;

vector<int> merge(vector<int> left, vector<int> right) {
    int i = 0;
    int j = 0;

    vector<int> sorted;
    while (i < left.size() && j < right.size()) {
        if (left[i] < right[j]) {
            sorted.push_back(left[i]);
            i++;
        } else {
            sorted.push_back(right[j]);
            j++;
        }
    }

    if (i < left.size()) {
        sorted.insert(sorted.end(), left.begin() + i, left.end());
    }

    if (j < right.size()) {
        sorted.insert(sorted.end(), right.begin() + j, right.end());
    }

    return sorted;
}

vector<int> merge_sort(vector<int> unsorted) {
    int n = unsorted.size();
    if (n <= 1) {
        return unsorted;
    }

    int mid = n / 2;
    vector<int> left(unsorted.begin(), unsorted.begin() + mid);
    vector<int> right(unsorted.begin() + mid, unsorted.end());

    left = merge_sort(left);
    right = merge_sort(right);

    return merge(left, right);
}

int main() {
    vector<int> numbers = {0, 1, 2, 3, 4, 5, 6, 7, 8, 9};
    vector<int> sorted = merge_sort(numbers);
    for (int num : sorted) {
        cout << num << " ";
    }

    return 0;
}