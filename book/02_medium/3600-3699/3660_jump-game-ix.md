# 3660 — Jump Game Ix

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan aturan permainan dan harus menentukan siapa yang menang atau berapa skor maksimal. Tugasmu adalah menganalisis permainan dan membuat keputusan optimal di setiap langkah.

Soal game theory menguji kemampuanmu berpikir beberapa langkah ke depan (minimax). Seringkali diselesaikan dengan DP (Dynamic Programming) untuk menyimpan hasil subproblem.

**Konsep kunci:** minimax, optimal play, game state, DP memoization, win/lose positions.

**Fungsi yang perlu kamu implementasikan:**
```go
func jumpGameIx(nums []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3660: Jump Game IX
// https://leetcode.com/problems/jump-game-ix/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func jumpGameIx(nums []int) []int {
	n := len(nums)
  // Alokasi slice integer
	pre := make([]int, n)
	pre[0] = nums[0]
	for i := 1; i < n; i++ {
		if nums[i] > pre[i-1] {
			pre[i] = nums[i]
		} else {
			pre[i] = pre[i-1]
		}
	}

  // Alokasi slice integer
	ans := make([]int, n)
	ans[n-1] = pre[n-1]
	minVal := nums[n-1]

	for i := n - 2; i >= 0; i-- {
		if pre[i] > minVal {
			ans[i] = ans[i+1]
		} else {
			ans[i] = pre[i]
		}
		if nums[i] < minVal {
			minVal = nums[i]
		}
	}
	return ans
}

func main() {
	fmt.Println(jumpGameIx([]int{3, 1, 4, 1, 5}))
	fmt.Println(jumpGameIx([]int{1, 2, 3, 4}))
	fmt.Println(jumpGameIx([]int{5, 4, 3, 2, 1}))
}
```
