package main

import (
	"sort"
)

func frequencySort(s string) string {
	mpCharFreq := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		char := s[i]
		val, ok := mpCharFreq[char]
		if ok {
			mpCharFreq[char] = val + 1
		} else {
			mpCharFreq[char] = 1
		}
	}

	mpFreqChar := make(map[int][]byte)
	for key, val := range mpCharFreq {
		slc, ok := mpFreqChar[val]
		if ok {
			mpFreqChar[val] = append(slc, key)
		} else {
			mpFreqChar[val] = []byte{key}
		}
	}

	keys := []int{}
	for key := range mpFreqChar {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	res := []byte{}
	for i := len(keys) - 1; i >= 0; i-- {
		freq := keys[i]
		chars := mpFreqChar[freq]
		for _, v := range chars {
			tempRes := make([]byte, 0, freq)
			for i := 0; i < freq; i++ {
				tempRes = append(tempRes, v)
			}
			res = append(res, tempRes...)
		}
	}
	return string(res)
}
