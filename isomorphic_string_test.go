package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func isIsomorphic(s string, t string) bool {
	return validatePair(s, t) && validatePair(t, s)
}

func validatePair(s string, t string) bool {
	mp := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		a, b := s[i], t[i]
		val, ok := mp[a]
		if !ok {
			mp[a] = b
		} else {
			if val != b {
				return false
			}
		}
	}
	return true
}

func TestIsomorphic1st(t *testing.T) {
	assert.True(t, isIsomorphic("egg", "add"))
}

func TestIsomorphic2nd(t *testing.T) {
	assert.False(t, isIsomorphic("foo", "bar"))
}

func TestIsomorphic3rd(t *testing.T) {
	assert.True(t, isIsomorphic("title", "paper"))
}

func TestIsomorphic4th(t *testing.T) {
	assert.False(t, isIsomorphic("badc", "baba"))
}
