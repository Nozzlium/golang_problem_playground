package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func plusOne(digits []int) []int {
	carryover := 1
	for i := len(digits) - 1; i >= 0; i-- {
		if carryover == 0 {
			return digits
		} else {
			num := digits[i] + carryover
			if num > 9 {
				digits[i] = 0
			} else {
				digits[i] = num
				carryover = 0
			}
		}
	}
	if carryover == 1 {
		return append([]int{1}, digits...)
	}
	return digits
}

func TestPlusOneFirst(t *testing.T) {
	res := plusOne([]int{1, 2, 3})
	assert.Equal(t, []int{1, 2, 4}, res)
}

func TestPlusOneSecond(t *testing.T) {
	res := plusOne([]int{4, 3, 2, 1})
	assert.Equal(t, []int{4, 3, 2, 2}, res)
}

func TestPlusOneThird(t *testing.T) {
	res := plusOne([]int{9, 9, 9})
	assert.Equal(t, []int{1, 0, 0, 0}, res)
}
