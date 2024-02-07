package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func maxProfit(prices []int) int {
	res := 0
	lowest, highest := prices[0], prices[0]
	for i := 1; i < len(prices); i++ {
		price := prices[i]
		if price < lowest {
			temp := highest - lowest
			if temp > res {
				res = temp
			}
			lowest = price
			highest = price
		} else {
			if price > highest {
				highest = price
			}
		}
	}
	if highest-lowest > res {
		return highest - lowest
	}
	return res
}

func TestStock1st(t *testing.T) {
	assert.Equal(t, 5, maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

func TestStock2nd(t *testing.T) {
	assert.Equal(t, 0, maxProfit([]int{7, 6, 4, 3, 1}))
}

func TestStock3rd(t *testing.T) {
	assert.Equal(t, 2, maxProfit([]int{2, 4, 1}))
}
