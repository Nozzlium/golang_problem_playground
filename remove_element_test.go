package main

import (
	"fmt"
	"testing"
)

func removeElement(nums []int, val int) int {
	queue := []int{}
	res := 0
	for i, num := range nums {
		if num == val {
			queue = append(queue, i)
		} else {
			if len(queue) > 0 {
				j := queue[0]
				queue = queue[1:]
				nums[j] = num
				queue = append(queue, i)
			}
			res++
		}
	}
	return res
}

func TestRemoveElementFirst(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	k := removeElement(nums, 3)
	fmt.Println(nums, k)
}

func TestRemoveElementSecond(t *testing.T) {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	k := removeElement(nums, 2)
	fmt.Println(nums, k)
}
