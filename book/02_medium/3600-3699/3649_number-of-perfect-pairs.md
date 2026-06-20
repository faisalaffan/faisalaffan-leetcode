# 3649 — Number Of Perfect Pairs

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func numberOfPerfectPairs(nums []int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3649: Number of Perfect Pairs
// https://leetcode.com/problems/number-of-perfect-pairs/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func numberOfPerfectPairs(nums []int) int64 {
	n := len(nums)
  // Alokasi slice integer
	arr := make([]int, n)
	for i, v := range nums {
		if v < 0 {
			arr[i] = -v
		} else {
			arr[i] = v
		}
	}

  // Urutkan secara ascending — O(n log n)
	sort.Ints(arr)

	var ans int64 = 0
	j := 0
	for i := 1; i < n; i++ {
		for 2*arr[j] < arr[i] {
			j++
		}
		ans += int64(i - j)
	}
	return ans
}

func main() {
	fmt.Println(numberOfPerfectPairs([]int{1, 2, 3, 4}))
	fmt.Println(numberOfPerfectPairs([]int{-1, 1, -2, 2}))
	fmt.Println(numberOfPerfectPairs([]int{5, 1, 2}))
}
```
