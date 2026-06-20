# 0231 — Power Of Two

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func IsPowerOfTwo(n int) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #231: Power of Two
// https://leetcode.com/problems/power-of-two/
// Difficulty: Easy

import "fmt"

// Time: O(1) | Space: O(1)
func IsPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}

func main() {
	fmt.Println(IsPowerOfTwo(1))
	fmt.Println(IsPowerOfTwo(16))
	fmt.Println(IsPowerOfTwo(3))
}
```
