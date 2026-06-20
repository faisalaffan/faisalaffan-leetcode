# 0045 — Jump Game Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan aturan permainan dan harus menentukan siapa yang menang atau berapa skor maksimal. Tugasmu adalah menganalisis permainan dan membuat keputusan optimal di setiap langkah.

Soal game theory menguji kemampuanmu berpikir beberapa langkah ke depan (minimax). Seringkali diselesaikan dengan DP (Dynamic Programming) untuk menyimpan hasil subproblem.

**Konsep kunci:** minimax, optimal play, game state, DP memoization, win/lose positions.

**Fungsi yang perlu kamu implementasikan:**
```go
func jump(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #45: Jump Game II
// https://leetcode.com/problems/jump-game-ii/
// Difficulty: Medium

import "fmt"

func jump(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}

	jumps := 0
	currentEnd := 0
	farthest := 0

	for i := 0; i < n-1; i++ {
		if i+nums[i] > farthest {
			farthest = i + nums[i]
		}
		if i == currentEnd {
			jumps++
			currentEnd = farthest
			if currentEnd >= n-1 {
				break
			}
		}
	}

	return jumps
}

func main() {
	// Test case 1
	fmt.Println(jump([]int{2, 3, 1, 1, 4})) // 2

	// Test case 2
	fmt.Println(jump([]int{2, 3, 0, 1, 4})) // 2

	// Test case 3
	fmt.Println(jump([]int{0})) // 0
}

// Time: O(n) | Space: O(1)
```
