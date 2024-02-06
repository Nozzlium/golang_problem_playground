package main

import (
	"fmt"
	"testing"
)

func removeDuplicates(nums []int) int {
	x := nums[0]
	queue := []int{}
	res := 1
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num != x {
			if len(queue) > 0 {
				j := queue[0]
				queue = queue[1:]
				nums[j] = num
				queue = append(queue, i)
			}
			x = num
			res++
		} else {
			queue = append(queue, i)
		}
	}
	return res
}

func TestDupeSecond(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	k := removeDuplicates(nums)
	fmt.Println(k, nums)
}

func TestDupeFirst(t *testing.T) {
	nums := []int{1, 1, 2}
	k := removeDuplicates(nums)
	fmt.Println(k, nums)
}

func TestDupeThird(t *testing.T) {
	nums := []int{1, 2}
	k := removeDuplicates(nums)
	fmt.Println(k, nums)
}

func TestDupeFourth(t *testing.T) {
	nums := []int{1, 2, 3}
	k := removeDuplicates(nums)
	fmt.Println(k, nums)
}
