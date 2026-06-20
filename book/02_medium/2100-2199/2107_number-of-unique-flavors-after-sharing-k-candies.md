# 2107 — Number Of Unique Flavors After Sharing K Candies

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func shareCandies(candies []int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2107: Number of Unique Flavors After Sharing K Candies
// https://leetcode.com/problems/number-of-unique-flavors-after-sharing-k-candies/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func shareCandies(candies []int, k int) int {
	if k >= len(candies) {
		return 0
	}

  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[int]int)
	for i := k; i < len(candies); i++ {
		freq[candies[i]]++
	}

	maxUnique := len(freq)
	for i := k; i < len(candies); i++ {
		// Add candies[i-k] back (give it away)
		freq[candies[i-k]]++
		// Remove candies[i] from the kept set
		freq[candies[i]]--
		if freq[candies[i]] == 0 {
			delete(freq, candies[i])
		}
		if len(freq) > maxUnique {
			maxUnique = len(freq)
		}
	}

	return maxUnique
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", shareCandies([]int{1, 2, 2, 3, 4, 3}, 3))
	// Expected: 3

	// Test case 2
	fmt.Println("Test 2:", shareCandies([]int{1, 1, 2, 3}, 2))
	// Expected: 2

	// Test case 3
	fmt.Println("Test 3:", shareCandies([]int{1, 1, 1, 1}, 2))
	// Expected: 1
}
```
