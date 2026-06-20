# 2595 — Number Of Even And Odd Bits

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func NumberOfEvenAndOddBits(n int) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #2595: Number of Even and Odd Bits
// https://leetcode.com/problems/number-of-even-and-odd-bits/
// Difficulty: Easy
// Time O(log n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(NumberOfEvenAndOddBits(17)) // [2,0]
	fmt.Println(NumberOfEvenAndOddBits(2))  // [0,1]
}

func NumberOfEvenAndOddBits(n int) []int {
	even, odd := 0, 0
	idx := 0
	for n > 0 {
		if n&1 == 1 {
			if idx%2 == 0 {
				even++
			} else {
				odd++
			}
		}
		n >>= 1
		idx++
	}
	return []int{even, odd}
}
```
