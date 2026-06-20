# 1995 — Count Special Quadruplets

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func CountSpecialQuadruplets(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^3), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1995: Count Special Quadruplets
// https://leetcode.com/problems/count-special-quadruplets/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountSpecialQuadruplets([]int{1, 2, 3, 6}))   // 1
	fmt.Println(CountSpecialQuadruplets([]int{3, 3, 6, 4, 5})) // 0
	fmt.Println(CountSpecialQuadruplets([]int{1, 1, 1, 3, 5})) // 4
}

// Time: O(n^3), Space: O(1)
func CountSpecialQuadruplets(nums []int) int {
	n := len(nums)
	count := 0
	for a := 0; a < n; a++ {
		for b := a + 1; b < n; b++ {
			for c := b + 1; c < n; c++ {
				for d := c + 1; d < n; d++ {
					if nums[a]+nums[b]+nums[c] == nums[d] {
						count++
					}
				}
			}
		}
	}
	return count
}
```
