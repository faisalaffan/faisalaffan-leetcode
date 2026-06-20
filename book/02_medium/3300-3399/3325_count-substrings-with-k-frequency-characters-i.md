# 3325 — Count Substrings With K Frequency Characters I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func numberOfSubstrings(s string, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer

**Waktu:** O(n) Space: O(1)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3325: Count Substrings With K-Frequency Characters I
// https://leetcode.com/problems/count-substrings-with-k-frequency-characters-i/
// Difficulty: Medium
// Time: O(n) Space: O(1)

import "fmt"

func main() {
	fmt.Println(numberOfSubstrings("abacb", 2)) // 4
	fmt.Println(numberOfSubstrings("abcde", 1)) // 15
	fmt.Println(numberOfSubstrings("aaaa", 2))  // 3
}

func numberOfSubstrings(s string, k int) int {
	n := len(s)
	cnt := [26]int{}
	ans, left := 0, 0

	for right := 0; right < n; right++ {
		idx := s[right] - 'a'
		cnt[idx]++

		for cnt[idx] >= k {
			ans += n - right
			cnt[s[left]-'a']--
			left++
		}
	}

	return ans
}
```
