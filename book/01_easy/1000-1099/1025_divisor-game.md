# 1025 — Divisor Game

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func divisorGame(n int) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1025: Divisor Game
// https://leetcode.com/problems/divisor-game/
// Difficulty: Easy
// Time: O(1) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(divisorGame(2)) // true
	fmt.Println(divisorGame(3)) // false
	fmt.Println(divisorGame(4)) // true
}

// LeetCode submission: divisorGame
func divisorGame(n int) bool {
	return n%2 == 0
}
```
