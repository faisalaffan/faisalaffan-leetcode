# 0180 — Consecutive Numbers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func consecutiveNumbers(nums []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #180: Consecutive Numbers
// https://leetcode.com/problems/consecutive-numbers/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

func consecutiveNumbers(nums []int) []int {
	if len(nums) < 3 {
		return nil
	}

	result := []int{}
  // Membuat map (HashMap) — pencarian O(1)
	seen := make(map[int]bool)

  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] == nums[i+1] && nums[i] == nums[i+2] && !seen[nums[i]] {
			result = append(result, nums[i])
			seen[nums[i]] = true
		}
	}

	return result
}

func main() {
	fmt.Println(consecutiveNumbers([]int{1, 1, 1, 2, 2, 3, 3, 3}))
	fmt.Println(consecutiveNumbers([]int{1, 2, 3, 4}))
	fmt.Println(consecutiveNumbers([]int{1, 1, 1, 1, 2, 2, 2}))
}
```
