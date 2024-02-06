package main

import (
	"fmt"
	"testing"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j := 0, 0
	res := []int{}
	for i < m && j < n {
		num1 := nums1[i]
		num2 := nums2[j]

		if num1 == num2 {
			res = append(res, num1, num2)
		}

		if num1 < num2 {
			res = append(res, num1)
			i++
		}

		if num2 < num1 {
			res = append(res, num2)
			j++
		}
	}
	if i < m {
		res = append(res, nums1[i:m]...)
	}
	if j < n {
		res = append(res, nums2[j:n]...)
	}
}

func TestMergeFirst(t *testing.T) {
	a := []int{2, 3, 4, 2, 1}
	fmt.Println(a[1:5])
}
