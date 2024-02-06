package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func climbStairs(n int) int {
	res := []int{1, 2}
	for i := 2; i < n; i++ {
		res = append(res, res[i-1]+res[i-2])
	}
	return res[n-1]
}

func TestStairsFirst(t *testing.T) {
	out := climbStairs(2)
	assert.Equal(t, 2, out)
}

func TestStairsSecond(t *testing.T) {
	out := climbStairs(3)
	assert.Equal(t, 3, out)
}

func TestStairsThird(t *testing.T) {
	out := climbStairs(4)
	assert.Equal(t, 5, out)
}
