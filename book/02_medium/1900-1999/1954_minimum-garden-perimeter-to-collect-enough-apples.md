# 1954 — Minimum Garden Perimeter To Collect Enough Apples

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func MinimumPerimeter(neededApples int64) int64`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(cuberoot(n)), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1954: Minimum Garden Perimeter to Collect Enough Apples
// https://leetcode.com/problems/minimum-garden-perimeter-to-collect-enough-apples/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MinimumPerimeter(1))
	fmt.Println(MinimumPerimeter(13))
	fmt.Println(MinimumPerimeter(1000000000))
}

// Time: O(cuberoot(n)), Space: O(1)
func MinimumPerimeter(neededApples int64) int64 {
	// For a garden with side length 2n (total apples = 2n(n+1)(2n+1))
	// Apples = 2 * n * (n+1) * (2n+1)
	// Perimeter = 8 * n

	n := int64(1)
	for {
		apples := 2 * n * (n + 1) * (2*n + 1)
		if apples >= neededApples {
			return 8 * n
		}
		n++
	}
}
```
