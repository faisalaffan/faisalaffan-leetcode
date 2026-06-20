# 0967 — Numbers With Same Consecutive Differences

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func numsSameConsecDiff(n int, k int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DFS

**Waktu:** O(n * 2^n)  |  **Ruang:** O(n * 2^n)

> 🎓 **Fresh Grad Tips:** Kuasai **DFS** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #967: Numbers With Same Consecutive Differences
// https://leetcode.com/problems/numbers-with-same-consecutive-differences/
// Difficulty: Medium

import "fmt"

// Time: O(n * 2^n) | Space: O(n * 2^n)
func numsSameConsecDiff(n int, k int) []int {
	if n == 1 {
		return []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	}

  // Alokasi slice
	ans := make([]int, 0)
	for d := 1; d <= 9; d++ {
		dfs(n, k, d, &ans)
	}
	return ans
}

func dfs(n, k, cur int, ans *[]int) {
	if n == 1 {
		*ans = append(*ans, cur)
		return
	}
	last := cur % 10
	if last+k <= 9 {
		dfs(n-1, k, cur*10+last+k, ans)
	}
	if k != 0 && last-k >= 0 {
		dfs(n-1, k, cur*10+last-k, ans)
	}
}

func main() {
	fmt.Println(numsSameConsecDiff(3, 7))
	fmt.Println(numsSameConsecDiff(2, 1))
	fmt.Println(numsSameConsecDiff(2, 0))
}
```
