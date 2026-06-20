# 0342 — Power Of Four

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func PowerOfFour(n int) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(1), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #342: Power of Four
// https://leetcode.com/problems/power-of-four/
// Difficulty: Easy

import "fmt"

// Time: O(1), Space: O(1)
func PowerOfFour(n int) bool {
	return n > 0 && n&(n-1) == 0 && (n-1)%3 == 0
}

func main() {
	fmt.Println(PowerOfFour(16))
	fmt.Println(PowerOfFour(5))
	fmt.Println(PowerOfFour(1))
}
```
