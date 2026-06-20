# 0326 — Power Of Three

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func IsPowerOfThree(n int) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(log n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #326: Power of Three
// https://leetcode.com/problems/power-of-three/
// Difficulty: Easy

import "fmt"

// Time: O(log n) | Space: O(1)
func IsPowerOfThree(n int) bool {
	if n < 1 {
		return false
	}
	for n%3 == 0 {
		n /= 3
	}
	return n == 1
}

func main() {
	fmt.Println(IsPowerOfThree(27))
	fmt.Println(IsPowerOfThree(0))
	fmt.Println(IsPowerOfThree(-1))
}
```
