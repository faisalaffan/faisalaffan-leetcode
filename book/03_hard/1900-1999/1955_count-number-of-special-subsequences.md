# 1955 — Count Number Of Special Subsequences

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func countSpecialSubsequences(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1955: Count Number of Special Subsequences
// https://leetcode.com/problems/count-number-of-special-subsequences/
// Difficulty: Hard
// DP[0] = count of subsequences matching pattern "0"
// DP[1] = count of subsequences matching pattern "0...1"
// DP[2] = count of subsequences matching pattern "0...1...2"

import "fmt"

const mod = 1_000_000_007

func countSpecialSubsequences(nums []int) int {
	dp0, dp1, dp2 := 0, 0, 0
	for _, x := range nums {
		switch x {
		case 0:
			dp0 = (dp0 + dp0 + 1) % mod
		case 1:
			dp1 = (dp1 + dp1 + dp0) % mod
		case 2:
			dp2 = (dp2 + dp2 + dp1) % mod
		}
	}
	return dp2
}

func main() {
	fmt.Println(countSpecialSubsequences([]int{0, 1, 2, 2}))       // Expected: 3
	fmt.Println(countSpecialSubsequences([]int{0, 1, 2, 0, 1, 2})) // Expected: 7
	fmt.Println(countSpecialSubsequences([]int{2, 2, 0, 0}))       // Expected: 0 (no 1s)
}
```
