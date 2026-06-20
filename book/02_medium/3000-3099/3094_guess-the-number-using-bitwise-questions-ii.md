# 3094 — Guess The Number Using Bitwise Questions Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func commonBits(num int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(log n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3094: Guess the Number Using Bitwise Questions II
// https://leetcode.com/problems/generate-binary-strings-without-adjacent-zeros/
// https://leetcode.com/problems/guess-the-number-using-bitwise-questions-ii/
// Difficulty: Medium [Paid]
// Time: O(log n) | Space: O(1)

import "fmt"

var hiddenNumber int

func commonBits(num int) int {
	count := 0
	x := hiddenNumber ^ num
	for x > 0 {
		count += x & 1
		x >>= 1
	}
	return 30 - count
}

func findNumber() int {
	n := 0
	for i := 0; i <= 30; i++ {
		bit1 := commonBits(n | (1 << i))
		bit0 := commonBits(n)
		if bit1 > bit0 {
			n |= (1 << i)
		}
	}
	return n
}

func main() {
	hiddenNumber = 42
	fmt.Println(findNumber())

	hiddenNumber = 100
	fmt.Println(findNumber())
}
```
