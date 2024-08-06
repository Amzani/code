package main

import "fmt"

/* 
 The technique used in the merge function involves merging from the back of the arrays to the front.
 This allows the function to utilize the extra space at the end of nums1 without the need for additional memory
 In-place Merge: The function merges the arrays in place without requiring additional memory, making it space-efficient.
 Efficient: The function runs in O(m + n) time, where m and n are the lengths of the input arrays, ensuring that it processes each element exactly once.
*/

func merge(nums1 []int, m int, nums2 []int, n int)  {
	k := (n + m) - 1
	i := m - 1
	j := n - 1
	for i >= 0 && j >= 0 {
		if nums1[i] >= nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k -
	}

	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}
}

func main() {
	nums1 := []int{2, 4, 7, 0, 0, 0}
	nums2 := []int{1, 3, 5}
	merge(nums1, 3, nums2, 3)
	fmt.Println(nums1)
}
