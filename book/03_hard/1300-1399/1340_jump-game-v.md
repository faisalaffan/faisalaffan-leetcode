# 1340 — Jump Game V

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan aturan permainan dan harus menentukan siapa yang menang atau berapa skor maksimal. Tugasmu adalah menganalisis permainan dan membuat keputusan optimal di setiap langkah.

Soal game theory menguji kemampuanmu berpikir beberapa langkah ke depan (minimax). Seringkali diselesaikan dengan DP (Dynamic Programming) untuk menyimpan hasil subproblem.

**Konsep kunci:** minimax, optimal play, game state, DP memoization, win/lose positions.

**Fungsi yang perlu kamu implementasikan:**
```go
func maxJumps(arr []int, d int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, DFS, Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1340: Jump Game V
// https://leetcode.com/problems/jump-game-v/
// Difficulty: Hard

import "fmt"

func maxJumps(arr []int, d int) int {
	n := len(arr)
  // Alokasi slice integer
	memo := make([]int, n)

	var dfs func(i int) int
	dfs = func(i int) int {
		if memo[i] > 0 {
			return memo[i]
		}
		res := 1 // at least we can stay here

		// Jump to the left
		for j := i - 1; j >= 0 && i-j <= d && arr[j] < arr[i]; j-- {
			if cand := 1 + dfs(j); cand > res {
				res = cand
			}
		}

		// Jump to the right
		for j := i + 1; j < n && j-i <= d && arr[j] < arr[i]; j++ {
			if cand := 1 + dfs(j); cand > res {
				res = cand
			}
		}

		memo[i] = res
		return res
	}

	ans := 0
	for i := 0; i < n; i++ {
		if cand := dfs(i); cand > ans {
			ans = cand
		}
	}
	return ans
}

func main() {
	// Example 1
	fmt.Println(maxJumps([]int{6, 4, 14, 6, 8, 13, 9, 7, 10, 6, 12}, 2))
	// Expected: 4

	// Example 2
	fmt.Println(maxJumps([]int{3, 3, 3, 3, 3}, 3))
	// Expected: 1

	// Example 3
	fmt.Println(maxJumps([]int{7, 6, 5, 4, 3, 2, 1}, 1))
	// Expected: 7
}
```
