# 0779 — K Th Symbol In Grammar

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func kthGrammar(n int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #779: K-th Symbol in Grammar
// https://leetcode.com/problems/k-th-symbol-in-grammar/
// Difficulty: Medium
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(kthGrammar(1, 1))
	fmt.Println(kthGrammar(2, 1))
	fmt.Println(kthGrammar(2, 2))
}

func kthGrammar(n int, k int) int {
	if n == 1 {
		return 0
	}

	parent := kthGrammar(n-1, (k+1)/2)
	if k%2 == 0 {
		return 1 - parent
	}
	return parent
}
```
