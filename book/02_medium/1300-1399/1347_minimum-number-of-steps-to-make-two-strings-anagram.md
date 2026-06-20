# 1347 — Minimum Number Of Steps To Make Two Strings Anagram

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan dua string. Tugasmu adalah menentukan apakah keduanya **anagram** — mengandung huruf yang sama dengan jumlah sama, hanya urutan berbeda.

**Cara berpikir:** Hitung frekuensi huruf string pertama, kurangi dengan string kedua. Semua harus nol.

**Fungsi Solusi:** `func minSteps(s string, t string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n) where n = length of strings  |  **Ruang:** O(1) - fixed size array of 26


## 💻 Solusi Go

```go
package main

// LeetCode #1347: Minimum Number of Steps to Make Two Strings Anagram
// https://leetcode.com/problems/minimum-number-of-steps-to-make-two-strings-anagram/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(minSteps("bab", "aba")) // 1

	// Test case 2
	fmt.Println(minSteps("leetcode", "practice")) // 5

	// Test case 3
	fmt.Println(minSteps("anagram", "mangaar")) // 0
}

// Time: O(n) where n = length of strings
// Space: O(1) - fixed size array of 26
func minSteps(s string, t string) int {
  // Alokasi slice
	freq := make([]int, 26)
  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
		freq[t[i]-'a']--
	}

	steps := 0
	for _, f := range freq {
		if f > 0 {
			steps += f
		}
	}
	return steps
}
```
