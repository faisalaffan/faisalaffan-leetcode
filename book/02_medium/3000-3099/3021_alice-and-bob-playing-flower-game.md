# 3021 — Alice And Bob Playing Flower Game

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func flowerGame(n int, m int) int64`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3021: Alice and Bob Playing Flower Game
// https://leetcode.com/problems/alice-and-bob-playing-flower-game/
// Difficulty: Medium
// Time: O(1) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(flowerGame(3, 2))
	fmt.Println(flowerGame(1, 1))
	fmt.Println(flowerGame(4, 4))
}

func flowerGame(n int, m int) int64 {
	oddN := int64((n + 1) / 2)
	evenN := int64(n / 2)
	oddM := int64((m + 1) / 2)
	evenM := int64(m / 2)
	return oddN*evenM + evenN*oddM
}
```
