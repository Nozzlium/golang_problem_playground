package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func twoSum(nums []int, target int) []int {
	res := []int{}
	mp := make(map[int]int)
	for index, num := range nums {
		firstIndex, ok := mp[num]
		if ok && 2*num == target {
			res = append(res, firstIndex)
			res = append(res, index)
			return res
		} else {
			mp[num] = index
		}
	}
	for _, num := range nums {
		firstIndex := mp[num]
		secondIndex, ok := mp[target-num]
		if ok && secondIndex != firstIndex {
			res = append(res, firstIndex)
			res = append(res, secondIndex)
			return res
		}
	}
	return res
}

func TestFirst(t *testing.T) {
	nums := []int{3, 3}
	target := 6
	assert.Equal(t, []int{0, 1}, twoSum(nums, target))
}

func TestSecond(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6
	assert.Equal(t, []int{1, 2}, twoSum(nums, target))
}
