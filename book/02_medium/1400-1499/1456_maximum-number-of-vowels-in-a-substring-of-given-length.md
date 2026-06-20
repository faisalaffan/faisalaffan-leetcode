# 1456 — Maximum Number Of Vowels In A Substring Of Given Length

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func isVowel(ch byte) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sliding Window

**Waktu:** O(n) where n = length of string  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Sliding Window** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1456: Maximum Number of Vowels in a Substring of Given Length
// https://leetcode.com/problems/maximum-number-of-vowels-in-a-substring-of-given-length/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(maxVowels("abciiidef", 3)) // 3

	// Test case 2
	fmt.Println(maxVowels("aeiou", 2)) // 2

	// Test case 3
	fmt.Println(maxVowels("leetcode", 3)) // 2

	// Test case 4
	fmt.Println(maxVowels("rhythms", 4)) // 0
}

func isVowel(ch byte) bool {
	return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u'
}

// Time: O(n) where n = length of string
// Space: O(1)
func maxVowels(s string, k int) int {
	// Count vowels in first window
	count := 0
	for i := 0; i < k; i++ {
		if isVowel(s[i]) {
			count++
		}
	}

	maxCount := count

	// Sliding window
	for i := k; i < len(s); i++ {
		if isVowel(s[i-k]) {
			count--
		}
		if isVowel(s[i]) {
			count++
		}
		if count > maxCount {
			maxCount = count
		}
	}

	return maxCount
}
```
