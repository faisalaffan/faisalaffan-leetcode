# 3895 — Count Digit Appearances

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func CountDigitAppearances(nums []int, digit int) int
```

> **💡 Hint:** For each number, extract digits and count matches to target digit.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(N * log M)  
**Kompleksitas Ruang:** O(1) where M = max value in nums

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3895: Count Digit Appearances
// https://leetcode.com/problems/count-digit-appearances/
// Difficulty: Medium
// Time: O(N * log M) | Space: O(1) where M = max value in nums
// Approach: For each number, extract digits and count matches to target digit.

import "fmt"

func CountDigitAppearances(nums []int, digit int) int {
	ans := 0
	for _, v := range nums {
		for v > 0 {
			if v%10 == digit {
				ans++
			}
			v /= 10
		}
	}
	return ans
}

func main() {
	// Example 1
	fmt.Println(CountDigitAppearances([]int{12, 54, 32, 22}, 2)) // Expected: 4

	// Example 2
	fmt.Println(CountDigitAppearances([]int{1, 34, 7}, 9)) // Expected: 0
}
```
