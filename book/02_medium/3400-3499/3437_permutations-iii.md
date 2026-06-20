# 3437 — Permutations Iii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func permute(n int) [][]int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DFS

**Waktu:** O(n!) Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **DFS** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3437: Permutations III
// https://leetcode.com/problems/permutations-iii/
// Difficulty: Medium [Paid]
// Time: O(n!) Space: O(n)

import "fmt"

func permute(n int) [][]int {
	var ans [][]int
	used := make([]bool, n+1)
  // Alokasi slice
	cur := make([]int, 0, n)

	var dfs func()
	dfs = func() {
		if len(cur) == n {
  // Alokasi slice
			tmp := make([]int, n)
			copy(tmp, cur)
			ans = append(ans, tmp)
			return
		}
		start := len(cur)%2 + 1
		for i := start; i <= n; i += 2 {
			if !used[i] {
				used[i] = true
				cur = append(cur, i)
				dfs()
				cur = cur[:len(cur)-1]
				used[i] = false
			}
		}
	}
	dfs()
	return ans
}

func main() {
	fmt.Println(len(permute(3))) // 2
	fmt.Println(len(permute(4))) // 4
	for _, p := range permute(3) {
		fmt.Println(p)
	}
}
```
