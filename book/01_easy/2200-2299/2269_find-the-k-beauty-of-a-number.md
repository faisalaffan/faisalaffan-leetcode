# 2269 — Find The K Beauty Of A Number

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func FindTheKBeautyOfANumber(num int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #2269: Find the K-Beauty of a Number
// https://leetcode.com/problems/find-the-k-beauty-of-a-number/
// Difficulty: Easy
// Time O(n * k) | Space O(1)

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(FindTheKBeautyOfANumber(240, 2))  // 2
	fmt.Println(FindTheKBeautyOfANumber(430043, 2)) // 2
}

func FindTheKBeautyOfANumber(num int, k int) int {
	s := strconv.Itoa(num)
	count := 0
	for i := 0; i <= len(s)-k; i++ {
		sub, _ := strconv.Atoi(s[i : i+k])
		if sub != 0 && num%sub == 0 {
			count++
		}
	}
	return count
}
```
