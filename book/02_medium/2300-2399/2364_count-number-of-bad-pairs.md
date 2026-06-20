# 2364 — Count Number Of Bad Pairs

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func countBadPairs(nums []int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2364: Count Number of Bad Pairs
// https://leetcode.com/problems/count-number-of-bad-pairs/
// Difficulty: Medium
// Time: O(n) | Space: O(n)
// Transform nums[i] into nums[i] - i. Good pairs have same transformed value.
// Bad pairs = total pairs - good pairs.

import "fmt"

func main() {
	fmt.Println(countBadPairs([]int{4, 1, 3, 3})) // 5
	fmt.Println(countBadPairs([]int{1, 2, 3, 4, 5})) // 0
}

func countBadPairs(nums []int) int64 {
	n := len(nums)
  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[int]int64)
	var good int64
	for i, v := range nums {
		key := v - i
		good += freq[key]
		freq[key]++
	}
	total := int64(n) * int64(n-1) / 2
	return total - good
}
```
