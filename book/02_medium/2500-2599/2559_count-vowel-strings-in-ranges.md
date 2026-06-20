# 2559 — Count Vowel Strings In Ranges

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func vowelStrings(words []string, queries [][]int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Prefix Sum

**Waktu:** O(n + q)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Prefix Sum** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2559: Count Vowel Strings in Ranges
// https://leetcode.com/problems/count-vowel-strings-in-ranges/
// Difficulty: Medium
// Time: O(n + q) | Space: O(n)

import "fmt"

func vowelStrings(words []string, queries [][]int) []int {
	isVowel := func(ch byte) bool {
		return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u'
	}

	n := len(words)
  // Alokasi slice
	prefix := make([]int, n+1)
	for i, w := range words {
		prefix[i+1] = prefix[i]
		if len(w) > 0 && isVowel(w[0]) && isVowel(w[len(w)-1]) {
			prefix[i+1]++
		}
	}

  // Alokasi slice
	ans := make([]int, len(queries))
	for i, q := range queries {
		ans[i] = prefix[q[1]+1] - prefix[q[0]]
	}
	return ans
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", vowelStrings([]string{"aba", "bcb", "ece", "aa", "e"}, [][]int{{0, 2}, {1, 4}, {1, 1}}))
	// Expected: [2,3,0]

	// Test case 2
	fmt.Println("Test 2:", vowelStrings([]string{"a", "e", "i"}, [][]int{{0, 2}, {0, 0}, {2, 2}}))
	// Expected: [3,1,1]

	// Test case 3: empty range
	fmt.Println("Test 3:", vowelStrings([]string{"a", "b", "c"}, [][]int{{1, 1}}))
	// Expected: [0]
}
```
