# 3696 — Maximum Distance Between Unequal Words In Array I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func MaximumDistanceBetweenUnequalWordsInArrayI(words []string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3696: Maximum Distance Between Unequal Words in Array I
// https://leetcode.com/problems/maximum-distance-between-unequal-words-in-array-i/
// Difficulty: Easy [Paid]

import "fmt"

func main() {
	fmt.Println(MaximumDistanceBetweenUnequalWordsInArrayI([]string{"leetcode", "leetcode", "codeforces"}))
	fmt.Println(MaximumDistanceBetweenUnequalWordsInArrayI([]string{"a", "b", "c", "a", "a"}))
	fmt.Println(MaximumDistanceBetweenUnequalWordsInArrayI([]string{"z", "z", "z"}))
}

// Time: O(n)
// Space: O(1)
func MaximumDistanceBetweenUnequalWordsInArrayI(words []string) int {
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
```
