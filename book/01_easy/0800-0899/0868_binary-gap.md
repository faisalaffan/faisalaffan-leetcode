# 0868 — Binary Gap

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func binaryGap(n int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(log n). Space: O(1).  |  **Ruang:** O(1).


## 💻 Solusi Go

```go
package main

// LeetCode #868: Binary Gap
// https://leetcode.com/problems/binary-gap/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(binaryGap(22))  // 2 (10110)
	fmt.Println(binaryGap(8))   // 0 (1000)
	fmt.Println(binaryGap(5))   // 2 (101)
	fmt.Println(binaryGap(6))   // 1 (110)
}

// binaryGap finds the longest distance between two consecutive 1s in binary representation.
// Time: O(log n). Space: O(1).
func binaryGap(n int) int {
	last := -1
	maxDist := 0
	for i := 0; n > 0; i++ {
		if n&1 == 1 {
			if last != -1 {
				if i-last > maxDist {
					maxDist = i - last
				}
			}
			last = i
		}
		n >>= 1
	}
	return maxDist
}
```
