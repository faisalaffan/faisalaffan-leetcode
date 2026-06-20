# 0414 — Third Maximum Number

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func ThirdMaximumNumber(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #414: Third Maximum Number
// https://leetcode.com/problems/third-maximum-number/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func ThirdMaximumNumber(nums []int) int {
	var max1, max2, max3 *int
	for _, v := range nums {
		val := v
		if max1 != nil && val == *max1 {
			continue
		}
		if max1 == nil || val > *max1 {
			max3 = max2
			max2 = max1
			max1 = &val
		} else if max2 == nil || val > *max2 {
			max3 = max2
			max2 = &val
		} else if max3 == nil || val > *max3 {
			max3 = &val
		}
	}
	if max3 != nil {
		return *max3
	}
	return *max1
}

func main() {
	fmt.Println(ThirdMaximumNumber([]int{3, 2, 1}))
	fmt.Println(ThirdMaximumNumber([]int{1, 2}))
	fmt.Println(ThirdMaximumNumber([]int{2, 2, 3, 1}))
}
```
