# 2843 — Count Symmetric Integers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func CountSymmetricIntegers(low int, high int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(high - low)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2843: Count Symmetric Integers
// https://leetcode.com/problems/count-symmetric-integers/
// Difficulty: Easy
// Time: O(high - low) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(CountSymmetricIntegers(1, 100))
	fmt.Println(CountSymmetricIntegers(1200, 1230))
}

func CountSymmetricIntegers(low int, high int) int {
	count := 0
	for n := low; n <= high; n++ {
		if isSymmetric(n) {
			count++
		}
	}
	return count
}

func isSymmetric(n int) bool {
	s := fmt.Sprintf("%d", n)
	if len(s)%2 == 1 {
		return false
	}
	half := len(s) / 2
	sum1, sum2 := 0, 0
	for i := 0; i < half; i++ {
		sum1 += int(s[i] - '0')
		sum2 += int(s[i+half] - '0')
	}
	return sum1 == sum2
}
```
