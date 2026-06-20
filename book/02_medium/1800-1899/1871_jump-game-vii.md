# 1871 — Jump Game Vii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan aturan permainan dan harus menentukan siapa yang menang atau berapa skor maksimal. Tugasmu adalah menganalisis permainan dan membuat keputusan optimal di setiap langkah.

Soal game theory menguji kemampuanmu berpikir beberapa langkah ke depan (minimax). Seringkali diselesaikan dengan DP (Dynamic Programming) untuk menyimpan hasil subproblem.

**Konsep kunci:** minimax, optimal play, game state, DP memoization, win/lose positions.

**Fungsi yang perlu kamu implementasikan:**
```go
func CanReach(s string, minJump int, maxJump int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Dynamic Programming, Prefix Sum

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1871: Jump Game VII
// https://leetcode.com/problems/jump-game-vii/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(CanReach("011010", 2, 3))
	fmt.Println(CanReach("01101110", 2, 3))
	fmt.Println(CanReach("00", 1, 1))
}

// Time: O(n), Space: O(n)
func CanReach(s string, minJump int, maxJump int) bool {
	n := len(s)
	if s[n-1] != '0' {
		return false
	}

	dp := make([]bool, n)
  // Alokasi slice integer
	prefix := make([]int, n+1)
	dp[0] = true
	prefix[1] = 1

	for i := 1; i < n; i++ {
		if s[i] == '1' {
			prefix[i+1] = prefix[i]
			continue
		}
		left := max(0, i-maxJump)
		right := i - minJump
		if right >= left {
			reachable := prefix[right+1] - prefix[left] > 0
			if reachable {
				dp[i] = true
			}
		}
		prefix[i+1] = prefix[i]
		if dp[i] {
			prefix[i+1]++
		}
	}
	return dp[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```
