# 1952 — Three Divisors

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func ThreeDivisors(n int) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(sqrt(n)), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1952: Three Divisors
// https://leetcode.com/problems/three-divisors/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ThreeDivisors(2))  // false
	fmt.Println(ThreeDivisors(4))  // true
	fmt.Println(ThreeDivisors(81)) // false
}

// Time: O(sqrt(n)), Space: O(1)
func ThreeDivisors(n int) bool {
	// n has exactly 3 divisors iff n is a perfect square of a prime
	if n < 4 {
		return false
	}

	// Check if sqrt(n) is integer
	root := 1
	for root*root < n {
		root++
	}
	if root*root != n {
		return false
	}

	// Check if root is prime
	for i := 2; i*i <= root; i++ {
		if root%i == 0 {
			return false
		}
	}
	return root > 1
}
```
