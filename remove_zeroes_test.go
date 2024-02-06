package main

import (
	"fmt"
	"testing"
)

func moveZeroes(nums []int) {
	queue := []int{}
	count := 0
	for i, num := range nums {
		if num == 0 {
			queue = append(queue, i)
			count++
		} else {
			if len(queue) > 0 {
				j := queue[0]
				queue = queue[1:]
				nums[j] = num
				queue = append(queue, i)
			}
		}
	}
	for i := (len(nums) - 1); i >= len(nums)-count; i-- {
		nums[i] = 0
	}
}

func TestRemoveZeroesFirst(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}

func TestRemoveZeroesThird(t *testing.T) {
	nums := []int{0}
	moveZeroes(nums)
	fmt.Println(nums)
}

func TestRemoveZeroesFourth(t *testing.T) {
	nums := []int{1, 0, 6, 0, 0, 0, 4, 3, 0, 3, 1, 0}
	moveZeroes(nums)
	fmt.Println(nums)
}
