# 1390 — Four Divisors

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func sumFourDivisors(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n * sqrt(m)) where n = len(nums), m = max value in nums  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1390: Four Divisors
// https://leetcode.com/problems/four-divisors/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(sumFourDivisors([]int{21, 4, 7})) // 32

	// Test case 2
	fmt.Println(sumFourDivisors([]int{21, 21})) // 64

	// Test case 3
	fmt.Println(sumFourDivisors([]int{1, 2, 3, 4, 5})) // 0
}

// Time: O(n * sqrt(m)) where n = len(nums), m = max value in nums
// Space: O(1)
func sumFourDivisors(nums []int) int {
	total := 0

	for _, num := range nums {
		divCount := 0
		divSum := 0

		for i := 1; i*i <= num; i++ {
			if num%i == 0 {
				divCount++
				divSum += i

				if i*i != num {
					divCount++
					divSum += num / i
				}
			}
			if divCount > 4 {
				break
			}
		}

		if divCount == 4 {
			total += divSum
		}
	}

	return total
}
```
