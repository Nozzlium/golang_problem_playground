package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func isPalindrome(x int) bool {
	// Jika angka negatif, sudah bisa dipastikan bukan palindrom
	// cth: -101, dibalik akan menjadi 101-
	if x < 0 {
		return false
	}
	slc := []int{}
	for x > 0 {
		one := x % 10
		slc = append(slc, one)
		x = x / 10
	}
	i, j := 0, len(slc)-1
	for i < j {
		if slc[i] != slc[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func TestPalindromeFirst(t *testing.T) {
	isPalindrome := isPalindrome(121)
	assert.True(t, isPalindrome)
}

func TestPalindromeSecond(t *testing.T) {
	isPalindrome := isPalindrome(-121)
	assert.False(t, isPalindrome)
}

func TestPalindromeThird(t *testing.T) {
	isPalindrome := isPalindrome(10)
	assert.False(t, isPalindrome)
}
