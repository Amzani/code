package main

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
    testCases := []struct {
        nums1    []int
        m        int
        nums2    []int
        n        int
        expected []int
    }{
        {[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
        {[]int{1}, 1, []int{}, 0, []int{1}},
        {[]int{0}, 0, []int{1}, 1, []int{1}},
        {[]int{4, 5, 6, 0, 0, 0}, 3, []int{1, 2, 3}, 3, []int{1, 2, 3, 4, 5, 6}},
        {[]int{1, 2, 3, 0, 0, 0}, 3, []int{4, 5, 6}, 3, []int{1, 2, 3, 4, 5, 6}},
        {[]int{2, 0}, 1, []int{1}, 1, []int{1, 2}},
        {[]int{0, 0, 0, 0, 0}, 0, []int{1, 2, 3, 4, 5}, 5, []int{1, 2, 3, 4, 5}},
        {[]int{1, 2, 4, 5, 6, 0}, 5, []int{3}, 1, []int{1, 2, 3, 4, 5, 6}},
        {[]int{1, 2, 3, 0, 0, 0}, 3, []int{1, 2, 3}, 3, []int{1, 1, 2, 2, 3, 3}},
        {[]int{-1, 0, 1, 0, 0, 0}, 3, []int{-3, -2, -1}, 3, []int{-3, -2, -1, -1, 0, 1}},
    }

    for _, tc := range testCases {
        nums1Copy := make([]int, len(tc.nums1))
        copy(nums1Copy, tc.nums1)
        merge(nums1Copy, tc.m, tc.nums2, tc.n)
        if !reflect.DeepEqual(nums1Copy, tc.expected) {
            t.Errorf("merge(%v, %d, %v, %d) = %v; expected %v", tc.nums1, tc.m, tc.nums2, tc.n, nums1Copy, tc.expected)
        }
    }
}
