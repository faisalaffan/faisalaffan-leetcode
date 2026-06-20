# 0526 — Beautiful Arrangement

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func CountArrangement(n int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Backtracking

**Waktu:** O(k) where k = number of valid permutations  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Backtracking** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #526: Beautiful Arrangement
// https://leetcode.com/problems/beautiful-arrangement/
// Difficulty: Medium
// Time: O(k) where k = number of valid permutations
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(CountArrangement(2))
	fmt.Println(CountArrangement(1))
}

func CountArrangement(n int) int {
	used := make([]bool, n+1)
	count := 0

	var backtrack func(pos int)
	backtrack = func(pos int) {
		if pos > n {
			count++
			return
		}
		for i := 1; i <= n; i++ {
			if !used[i] && (i%pos == 0 || pos%i == 0) {
				used[i] = true
				backtrack(pos + 1)
				used[i] = false
			}
		}
	}

	backtrack(1)
	return count
}
```
