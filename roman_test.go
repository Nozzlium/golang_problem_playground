package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getDec(b byte) int {
	switch b {
	case 73:
		return 1
	case 86:
		return 5
	case 88:
		return 10
	case 76:
		return 50
	case 67:
		return 100
	case 68:
		return 500
	case 77:
		return 1000
	}
	return 0
}

func romanToInt(s string) int {
	largest := 0
	res := 0
	// Kita akan melakukan iterasi string bilangan romawi dari belakang
	for i := len(s) - 1; i >= 0; i-- {
		char := getDec(s[i])
		/*
			Angka romawi kalau diiterasi dari belakang ke depan pasti akan selalu
			naik jumlahnya (contoh: VII) kecuali dalam kasus seperti IV dan IX. Itulah
			gunanya if clause ini. Ketika nilai pada suatu iterasi lebih kecil dari
			iterasi sebelumnya, bisa dipastikan kita sedang bertemu dengan kasus seperti
			IV dan IX tadi.
		*/
		if char >= largest {
			// Selama masih naik nilainya ketika diiterasi terus jumlahkan nilainya
			// contoh: VIII, kita akan selalu menjumlah kan 1 + 1 + 1 + 5 sampai dapat nilai 8
			largest = char
			res += char
		} else {
			// Contoh kasus adalah pada bilangan "IX", ketika iterasi sampai pada "I" setelah sebelumnya
			// kita sampai pada "X", maka yang perlu kita lakukan adalah mengurangi total perhitungan dengan 1
			res -= char
		}
	}
	return res
}

func TestRomanFirst(t *testing.T) {
	assert.Equal(t, 3, romanToInt("III"))
}

func TestRomanSecond(t *testing.T) {
	assert.Equal(t, 58, romanToInt("LVIII"))
}

func TestRomanThird(t *testing.T) {
	assert.Equal(t, 1994, romanToInt("MCMXCIV"))
}
