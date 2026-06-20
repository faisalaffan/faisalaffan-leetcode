# 0055 — Jump Game

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan aturan permainan dan harus menentukan siapa yang menang atau berapa skor maksimal. Tugasmu adalah menganalisis permainan dan membuat keputusan optimal di setiap langkah.

Soal game theory menguji kemampuanmu berpikir beberapa langkah ke depan (minimax). Seringkali diselesaikan dengan DP (Dynamic Programming) untuk menyimpan hasil subproblem.

**Konsep kunci:** minimax, optimal play, game state, DP memoization, win/lose positions.

**Fungsi yang perlu kamu implementasikan:**
```go
func canJump(nums []int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #55: Jump Game
// https://leetcode.com/problems/jump-game/
// Difficulty: Medium

import "fmt"

func canJump(nums []int) bool {
	reachable := 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(nums); i++ {
		if i > reachable {
			return false
		}
		if i+nums[i] > reachable {
			reachable = i + nums[i]
		}
	}
	return true
}

func main() {
	// Test case 1
	fmt.Println(canJump([]int{2, 3, 1, 1, 4})) // true

	// Test case 2
	fmt.Println(canJump([]int{3, 2, 1, 0, 4})) // false

	// Test case 3
	fmt.Println(canJump([]int{0})) // true
}

// Time: O(n) | Space: O(1)
```
