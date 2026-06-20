# 2086 — Minimum Number Of Food Buckets To Feed The Hamsters

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumBuckets(hamsters string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2086: Minimum Number of Food Buckets to Feed the Hamsters
// https://leetcode.com/problems/minimum-number-of-food-buckets-to-feed-the-hamsters/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func minimumBuckets(hamsters string) int {
	n := len(hamsters)
	buckets := make([]bool, n)
	count := 0

	for i := 0; i < n; i++ {
		if hamsters[i] == 'H' {
			if i > 0 && buckets[i-1] {
				continue
			}
			if i+1 < n && hamsters[i+1] == '.' {
				buckets[i+1] = true
				count++
			} else if i > 0 && hamsters[i-1] == '.' {
				buckets[i-1] = true
				count++
			} else {
				return -1
			}
		}
	}
	return count
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minimumBuckets("H..H"))
	// Expected: 2

	// Test case 2
	fmt.Println("Test 2:", minimumBuckets(".H.H."))
	// Expected: 1

	// Test case 3
	fmt.Println("Test 3:", minimumBuckets("HH"))
	// Expected: -1
}
```
