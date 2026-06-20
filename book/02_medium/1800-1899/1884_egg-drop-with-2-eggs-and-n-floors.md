# 1884 — Egg Drop With 2 Eggs And N Floors

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func TwoEggDrop(n int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(1), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1884: Egg Drop With 2 Eggs and N Floors
// https://leetcode.com/problems/egg-drop-with-2-eggs-and-n-floors/
// Difficulty: Medium

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(TwoEggDrop(2))
	fmt.Println(TwoEggDrop(100))
	fmt.Println(TwoEggDrop(10))
}

// Time: O(1), Space: O(1)
func TwoEggDrop(n int) int {
	// Solve x(x+1)/2 >= n
	// x = ceil((-1 + sqrt(1 + 8n)) / 2)
	return int(math.Ceil((-1 + math.Sqrt(1+8*float64(n))) / 2))
}
```
