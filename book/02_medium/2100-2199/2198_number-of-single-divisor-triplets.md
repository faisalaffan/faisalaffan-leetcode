# 2198 — Number Of Single Divisor Triplets

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func singleDivisorTriplet(nums []int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n + m^3)  
**Kompleksitas Ruang:** O(m) where m = max value

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2198: Number of Single Divisor Triplets
// https://leetcode.com/problems/number-of-single-divisor-triplets/
// Difficulty: Medium [Paid]
// Time: O(n + m^3) | Space: O(m) where m = max value

import "fmt"

func singleDivisorTriplet(nums []int) int64 {
  // Alokasi slice integer
	count := make([]int, 101)
	for _, v := range nums {
		count[v]++
	}

	var res int64 = 0

	for a := 1; a <= 100; a++ {
		if count[a] == 0 {
			continue
		}
		for b := a; b <= 100; b++ {
			if count[b] == 0 {
				continue
			}
			for c := b; c <= 100; c++ {
				if count[c] == 0 {
					continue
				}

				s := a + b + c
				div := 0
				if s%a == 0 {
					div++
				}
				if s%b == 0 {
					div++
				}
				if s%c == 0 {
					div++
				}
				if div != 1 {
					continue
				}

				if a == b && b == c {
					res += int64(count[a]) * int64(count[a]-1) * int64(count[a]-2) / 6
				} else if a == b {
					res += int64(count[a]) * int64(count[a]-1) / 2 * int64(count[c])
				} else if b == c {
					res += int64(count[a]) * int64(count[b]) * int64(count[b]-1) / 2
				} else {
					res += int64(count[a]) * int64(count[b]) * int64(count[c])
				}
			}
		}
	}
	return res
}

func main() {
	// Test case 1
	fmt.Println(singleDivisorTriplet([]int{4, 6, 7, 3, 2}))
	// Expected: 12

	// Test case 2
	fmt.Println(singleDivisorTriplet([]int{1, 2, 2}))
	// Expected: 0
}
```
