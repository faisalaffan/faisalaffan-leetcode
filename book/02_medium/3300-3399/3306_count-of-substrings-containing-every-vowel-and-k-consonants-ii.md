# 3306 — Count Of Substrings Containing Every Vowel And K Consonants Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func countOfSubstringsII(word string, k int) int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Two Pointer

**Waktu:** O(n) Space: O(1)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

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

  // HashMap: O(1) lookup
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
