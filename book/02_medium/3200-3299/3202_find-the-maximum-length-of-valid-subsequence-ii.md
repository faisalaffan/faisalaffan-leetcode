# 3202 — Find The Maximum Length Of Valid Subsequence Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string berisi tanda kurung: `()`, `[]`, `{}`. Tugasmu adalah memeriksa apakah string tersebut **valid** — setiap kurung buka harus ditutup oleh kurung yang sesuai dalam urutan benar.

Contoh valid: `()[]{}`, `({[]})`. Tidak valid: `(]`, `([)]`.

**Cara berpikir:** Gunakan Stack. Kurung buka → push. Kurung tutup → pop dan cek kecocokan. Di akhir, stack harus kosong.

**Fungsi Solusi:** `func maximumLength(nums []int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** O(n * k)  |  **Ruang:** O(k)

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3202: Find the Maximum Length of Valid Subsequence II
// https://leetcode.com/problems/find-the-maximum-length-of-valid-subsequence-ii/
// Difficulty: Medium
// Time: O(n * k) | Space: O(k)

import "fmt"

func maximumLength(nums []int, k int) int {
  // Matriks 2D
	dp := make([][]int, k)
  // Range loop
	for i := range dp {
		dp[i] = make([]int, k)
	}

	ans := 0
	for _, v := range nums {
		cur := v % k
		for j := 0; j < k; j++ {
			need := (j - cur%k + k) % k
			dp[cur][j] = dp[need][j] + 1
			if dp[cur][j] > ans {
				ans = dp[cur][j]
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(maximumLength([]int{1, 4, 2, 3, 1, 4}, 3)) // Expected: 4
	fmt.Println(maximumLength([]int{1, 2, 3, 4, 5}, 2))     // Expected: 3
}
```
