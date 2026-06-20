# 0507 — Perfect Number

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func PerfectNumber(num int) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(sqrt(n)), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #507: Perfect Number
// https://leetcode.com/problems/perfect-number/
// Difficulty: Easy

import "fmt"

// Time: O(sqrt(n)), Space: O(1)
func PerfectNumber(num int) bool {
	if num <= 1 {
		return false
	}
	sum := 1
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			sum += i
			if i != num/i {
				sum += num / i
			}
		}
	}
	return sum == num
}

func main() {
	fmt.Println(PerfectNumber(28))
	fmt.Println(PerfectNumber(7))
	fmt.Println(PerfectNumber(6))
}
```
