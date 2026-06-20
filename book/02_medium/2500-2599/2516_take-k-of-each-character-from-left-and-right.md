# 2516 — Take K Of Each Character From Left And Right

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func takeCharacters(s string, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, Sliding Window

**Waktu:** O(n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2516: Take K of Each Character From Left and Right
// https://leetcode.com/problems/take-k-of-each-character-from-left-and-right/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// Sliding window: find max middle substring that doesn't exceed total-k per char.
// Answer = n - len(max window).

import "fmt"

func main() {
	fmt.Println(takeCharacters("aabaaaacaabc", 2)) // 8
	fmt.Println(takeCharacters("a", 1))            // 1
}

func takeCharacters(s string, k int) int {
	n := len(s)
	if k == 0 {
		return 0
	}

	// Total counts
  // Alokasi slice
	total := make([]int, 3)
	for _, ch := range s {
		total[ch-'a']++
	}
	for _, c := range total {
		if c < k {
			return -1
		}
	}

	// Max middle window where each char <= total-chars - k
	need := []int{total[0] - k, total[1] - k, total[2] - k}
  // Alokasi slice
	cnt := make([]int, 3)
	left, maxWindow := 0, 0

	for right := 0; right < n; right++ {
		cnt[s[right]-'a']++
		for cnt[0] > need[0] || cnt[1] > need[1] || cnt[2] > need[2] {
			cnt[s[left]-'a']--
			left++
		}
		if right-left+1 > maxWindow {
			maxWindow = right - left + 1
		}
	}
	return n - maxWindow
}
```
