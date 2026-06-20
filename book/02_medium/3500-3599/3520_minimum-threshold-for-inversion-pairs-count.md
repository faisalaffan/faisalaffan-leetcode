# 3520 — Minimum Threshold For Inversion Pairs Count

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumThresholdForInversionPairsCount(arr []int, threshold int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3520: Minimum Threshold for Inversion Pairs Count
// https://leetcode.com/problems/minimum-threshold-for-inversion-pairs-count/
// Difficulty: Medium [Paid]
// Complexity: O(n log n) time, O(n) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", MinimumThresholdForInversionPairsCount([]int{1, 3, 2, 4}, 1))
	// Test case 2
	fmt.Println("Test 2:", MinimumThresholdForInversionPairsCount([]int{4, 3, 2, 1}, 3))
	// Test case 3
	fmt.Println("Test 3:", MinimumThresholdForInversionPairsCount([]int{1, 2, 3}, 0))
}

func MinimumThresholdForInversionPairsCount(arr []int, threshold int) int {
	// Count inversion pairs (i < j, arr[i] > arr[j])
	count := 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				count++
			}
		}
	}
	if count >= threshold {
		return 1
	}
	return 0
}
```
