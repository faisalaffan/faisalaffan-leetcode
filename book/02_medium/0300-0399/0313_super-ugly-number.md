# 0313 — Super Ugly Number

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func nthSuperUglyNumber(n int, primes []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n * len(primes)), Space: O(n + len(primes))  |  **Ruang:** O(n + len(primes))


## 💻 Solusi Go

```go
package main

// LeetCode #313: Super Ugly Number
// https://leetcode.com/problems/super-ugly-number/
// Difficulty: Medium
// Time: O(n * len(primes)), Space: O(n + len(primes))

import "fmt"

func nthSuperUglyNumber(n int, primes []int) int {
  // Alokasi slice
	ugly := make([]int, n)
	ugly[0] = 1

  // Alokasi slice
	pointers := make([]int, len(primes))
  // Alokasi slice
	values := make([]int, len(primes))
	for i, p := range primes {
		values[i] = p
	}

	for i := 1; i < n; i++ {
		minVal := values[0]
		for _, v := range values {
			if v < minVal {
				minVal = v
			}
		}
		ugly[i] = minVal

		for j := range values {
			if values[j] == minVal {
				pointers[j]++
				values[j] = ugly[pointers[j]] * primes[j]
			}
		}
	}

	return ugly[n-1]
}

func main() {
	fmt.Println(nthSuperUglyNumber(12, []int{2, 7, 13, 19}))
	fmt.Println(nthSuperUglyNumber(1, []int{2, 3, 5}))
	fmt.Println(nthSuperUglyNumber(6, []int{2, 3, 5}))
}
```
