# 3706 — Maximum Distance Between Unequal Words In Array Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func maximumDistanceBetweenUnequalWordsInArrayIi(words []string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3706: Maximum Distance Between Unequal Words in Array II
// https://leetcode.com/problems/maximum-distance-between-unequal-words-in-array-ii/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func maximumDistanceBetweenUnequalWordsInArrayIi(words []string) int {
	n := len(words)
	ans := 0
	for i := 0; i < n; i++ {
		if words[i] != words[0] {
			if i+1 > ans {
				ans = i + 1
			}
		}
		if words[i] != words[n-1] {
			if n-i > ans {
				ans = n - i
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(maximumDistanceBetweenUnequalWordsInArrayIi([]string{"leetcode", "leetcode", "codeforces"}))
	fmt.Println(maximumDistanceBetweenUnequalWordsInArrayIi([]string{"a", "b", "c", "a", "a"}))
	fmt.Println(maximumDistanceBetweenUnequalWordsInArrayIi([]string{"z", "z", "z"}))
}
```
