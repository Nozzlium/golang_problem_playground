package main

import (
	"fmt"
	"testing"
)

func generate(numRows int) [][]int {
	res := [][]int{
		{1},
		{1, 1},
	}
	for i := 2; i < numRows; i++ {
		nums := res[i-1]
		middlePart := []int{}
		for j := 0; j < i-1; j++ {
			middlePart = append(middlePart, nums[j]+nums[j+1])
		}
		middlePart = append(middlePart, 1)
		res = append(res, append([]int{1}, middlePart...))
	}
	return res[:numRows]
}

func TestPascalFirst(t *testing.T) {
	arr := generate(5)
	fmt.Println(arr)
}

func TestPascalSecond(t *testing.T) {
	arr := generate(1)
	fmt.Println(arr)
}
