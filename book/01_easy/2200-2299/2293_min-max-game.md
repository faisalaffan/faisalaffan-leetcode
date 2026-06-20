# 2293 — Min Max Game

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan aturan permainan dan harus menentukan siapa yang menang atau berapa skor maksimal. Tugasmu adalah menganalisis permainan dan membuat keputusan optimal di setiap langkah.

Soal game theory menguji kemampuanmu berpikir beberapa langkah ke depan (minimax). Seringkali diselesaikan dengan DP (Dynamic Programming) untuk menyimpan hasil subproblem.

**Konsep kunci:** minimax, optimal play, game state, DP memoization, win/lose positions.

**Fungsi yang perlu kamu implementasikan:**
```go
func MinMaxGame(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2293: Min Max Game
// https://leetcode.com/problems/min-max-game/
// Difficulty: Easy
// Time O(n) | Space O(n)

import "fmt"

func main() {
	fmt.Println(MinMaxGame([]int{1, 3, 5, 2, 4, 8, 2, 2})) // 1
	fmt.Println(MinMaxGame([]int{3}))                        // 3
}

func MinMaxGame(nums []int) int {
	for len(nums) > 1 {
  // Alokasi slice integer
		next := make([]int, len(nums)/2)
  // Loop linear O(n): iterasi setiap elemen
		for i := 0; i < len(next); i++ {
			if i%2 == 0 {
				next[i] = min(nums[2*i], nums[2*i+1])
			} else {
				next[i] = max(nums[2*i], nums[2*i+1])
			}
		}
		nums = next
	}
	return nums[0]
}
```
