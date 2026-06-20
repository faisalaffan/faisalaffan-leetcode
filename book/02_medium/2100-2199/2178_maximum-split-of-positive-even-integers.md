# 2178 — Maximum Split Of Positive Even Integers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func maximumEvenSplit(finalSum int64) []int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(sqrt(finalSum))  
**Kompleksitas Ruang:** O(sqrt(finalSum))

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2178: Maximum Split of Positive Even Integers
// https://leetcode.com/problems/maximum-split-of-positive-even-integers/
// Difficulty: Medium
// Time: O(sqrt(finalSum)) | Space: O(sqrt(finalSum))

import "fmt"

func maximumEvenSplit(finalSum int64) []int64 {
	if finalSum%2 != 0 {
		return []int64{}
	}

	result := []int64{}
	cur := int64(2)

	for finalSum >= cur {
		result = append(result, cur)
		finalSum -= cur
		cur += 2
	}

	// Add remaining to last element
	if finalSum > 0 {
		result[len(result)-1] += finalSum
	}

	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", maximumEvenSplit(12))
	// Expected: [2, 4, 6] or [2, 10] — both valid, but greedy gives max length

	// Test case 2
	fmt.Println("Test 2:", maximumEvenSplit(7))
	// Expected: []

	// Test case 3
	fmt.Println("Test 3:", maximumEvenSplit(28))
	// Expected: [2, 4, 6, 16] (greedy fills smallest unique evens, remainder added to last)
}
```
