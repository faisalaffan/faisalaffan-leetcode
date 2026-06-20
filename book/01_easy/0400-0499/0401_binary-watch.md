# 0401 — Binary Watch

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func countBits(n int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(1), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #401: Binary Watch
// https://leetcode.com/problems/binary-watch/
// Difficulty: Easy

import (
	"fmt"
)

func countBits(n int) int {
	count := 0
	for n > 0 {
		n &= n - 1
		count++
	}
	return count
}

// Time: O(1), Space: O(1)
func BinaryWatch(turnedOn int) []string {
	var result []string
	for h := 0; h < 12; h++ {
		for m := 0; m < 60; m++ {
			if countBits(h)+countBits(m) == turnedOn {
				result = append(result, fmt.Sprintf("%d:%02d", h, m))
			}
		}
	}
	return result
}

func main() {
	fmt.Println(BinaryWatch(1))
	fmt.Println(BinaryWatch(9))
}
```
