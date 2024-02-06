package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func longestCommonPrefix(strs []string) string {
	// Potongan kode di bawah bertujuan untuk menentukan string
	// paling pendek di antara string-string dalam array.
	// Karena iterasi string-string use case akan mentok pada string terpendek.
	shortest := len(strs[0])
	for i := 1; i < len(strs); i++ {
		temp := len(strs[i])
		if temp < shortest {
			shortest = temp
		}
	}
	count := 0
	// variable untuk menyimpan prefix untuk nantinya di-return
	res := ""
	// loop untuk mengiterasi karakter tiap string.
	for count < shortest {
		var temp byte
		for _, str := range strs {
			char := str[count]
			// inisiasi karakter yang akan kita bandingkan dengan karakter
			// dari string lainnya pada index yang sama
			if temp == 0 {
				temp = char
			} else {
				// Ketika ditemukan ada karakter yang berbeda pada index yang sama
				// berarti telah tertutup kemungkinan ada prefix yang sama
				// dari index tersebut ke belakang.
				if char != temp {
					return res
				}
			}
		}
		// Jika sudah keluar dari loop tersebut, bisa dipastikan semua karakter
		// dari semua string pada suatu index sama. Maka bisa kita concat ke dalam
		// variable sebagai string yang berpotensi sebagai hasil return
		res += string(temp)
		count++
	}
	return res
}

func TestPrefixFirst(t *testing.T) {
	strs := []string{"flower", "flow", "flight"}
	assert.Equal(t, "fl", longestCommonPrefix(strs))
}

func TestPrefixSecond(t *testing.T) {
	strs := []string{"dog", "racecar", "car"}
	assert.Equal(t, "", longestCommonPrefix(strs))
}

func TestPrefixThird(t *testing.T) {
	strs := []string{"", "car"}
	assert.Equal(t, "", longestCommonPrefix(strs))
}
