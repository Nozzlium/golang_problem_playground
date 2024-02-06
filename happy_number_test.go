package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func isHappy(n int) bool {
	cal := n
	dict := make(map[int]bool)
	for cal > 1 {
		_, ok := dict[cal]
		if ok {
			return false
		}
		dict[cal] = true
		cal = process(cal)
	}
	return true
}

func process(n int) int {
	res := 0
	cal := n
	for cal > 0 {
		one := cal % 10
		res += one * one
		cal /= 10
	}
	return res
}

func TestIsHappy1st(t *testing.T) {
	n := isHappy(19)
	assert.True(t, n)
}

func TestIsHappy2nd(t *testing.T) {
	n := isHappy(2)
	assert.False(t, n)
}
