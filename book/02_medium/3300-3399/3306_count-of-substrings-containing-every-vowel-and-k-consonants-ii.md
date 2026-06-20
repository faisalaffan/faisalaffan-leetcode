# 3306 — Count Of Substrings Containing Every Vowel And K Consonants Ii

## Deskripsi

**Soal:** [3306. Count Of Substrings Containing Every Vowel And K Consonants Ii](https://leetcode.com/problems/count-of-substrings-containing-every-vowel-and-k-consonants-ii/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n) Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3306: Count of Substrings Containing Every Vowel and K Consonants II
// https://leetcode.com/problems/count-of-substrings-containing-every-vowel-and-k-consonants-ii/
// Difficulty: Medium
// Time: O(n) Space: O(1)

import "fmt"

func main() {
	fmt.Println(countOfSubstringsII("aeioqq", 1))         // 0
	fmt.Println(countOfSubstringsII("ieaouqqieaouqq", 1)) // 3
	fmt.Println(countOfSubstringsII("aeiou", 0))          // 1
}

func countOfSubstringsII(word string, k int) int64 {
	return atLeastK(word, k) - atLeastK(word, k+1)
}

func atLeastK(word string, k int) int64 {
	n := len(word)
	isVowel := func(c byte) bool {
		return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
	}

  // Membuat map untuk pencarian O(1): key → value
	vowelCnt := make(map[byte]int)
	cons := 0
	var ans int64
	left := 0

	for right := 0; right < n; right++ {
		c := word[right]
		if isVowel(c) {
			vowelCnt[c]++
		} else {
			cons++
		}

		for len(vowelCnt) == 5 && cons >= k {
			out := word[left]
			if isVowel(out) {
				vowelCnt[out]--
				if vowelCnt[out] == 0 {
					delete(vowelCnt, out)
				}
			} else {
				cons--
			}
			left++
		}
		ans += int64(left)
	}
	return ans
}
```
