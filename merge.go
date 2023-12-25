package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}

	results, i, j := make([]int, 0), 0, 0

	for i < m || j < n {
		if j >= n {
			results = append(results, nums1[i:]...)
			break
		}

		if i >= m {
			results = append(results, nums2[j:]...)
			break
		}

		if nums1[i] < nums2[j] {
			results = append(results, nums1[i])
			i++
		} else {
			results = append(results, nums2[j])
			j++
		}
	}

	copy(nums1, results)
}

func main() {
	nums1 := []int{2,0}
	nums2 := []int{1}
	merge(nums1, 1, nums2, 1)
	fmt.Println(nums1)
}
