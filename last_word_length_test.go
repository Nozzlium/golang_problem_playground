package main

import (
	"fmt"
	"testing"
)

func lengthOfLastWord(s string) int {
	space := false
	count := 0
	for i := 0; i < len(s); i++ {
		char := s[i]
		if char != 32 {
			if space {
				count = 1
				space = false
			} else {
				count++
			}
		} else {
			if !space {
				space = true
			}
		}
	}
	return count
}

func TestWordLengthFirst(t *testing.T) {
	a := lengthOfLastWord("Hello World")
	fmt.Println(a)
}

func TestWordLengthSecond(t *testing.T) {
	a := lengthOfLastWord("   fly me   to   the moon  ")
	fmt.Println(a)
}

func TestWordLengthThird(t *testing.T) {
	a := lengthOfLastWord("luffy is still joyboy")
	fmt.Println(a)
}
