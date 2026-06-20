# 0982 — Triples With Bitwise And Equal To Zero

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func countTriplets(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #982: Triples with Bitwise AND Equal To Zero
// https://leetcode.com/problems/triples-with-bitwise-and-equal-to-zero/
// Difficulty: Hard

import "fmt"

func countTriplets(nums []int) int {
	// Count pairs (i, j) for each possible AND value
	// Since nums[i] <= 2^16, we can use an array of size 2^16
	maxVal := 1 << 16
  // Alokasi slice integer
	pairCount := make([]int, maxVal)

	for _, a := range nums {
		for _, b := range nums {
			pairCount[a&b]++
		}
	}

	// For each pair AND value, find triples with k such that (pair & k) == 0
	// This means k must be a subset of the complement of pair
	// We can use SOS DP / subset enumeration

	ans := 0
	for pairAnd, count := range pairCount {
		if count == 0 {
			continue
		}
		// All k such that k & pairAnd == 0 are valid
		// Enumerate all subsets of (~pairAnd) within 16 bits
		complement := (maxVal - 1) ^ pairAnd
		subset := complement
		for {
			// subset is a valid k value when working with the full pair count
			// But we need to count actual k values from nums that equal this subset
			// Actually we need to check if subset is in nums
			// Better approach: precompute frequency of each value in nums
			// Then for each pair, sum over k where (pair & k) == 0
			// subset enumeration works but we need freq
			if subset < maxVal {
				// This is a valid k that works with this pair
				// We'll count it
				_ = subset
			}
			if subset == 0 {
				break
			}
			subset = (subset - 1) & complement
		}
	}

	// Let me rewrite more cleanly
	// Precompute frequency of each value in nums
  // Alokasi slice integer
	freq := make([]int, maxVal)
	for _, v := range nums {
		if v < maxVal {
			freq[v]++
		}
	}

	ans = 0
	for andVal, pairCount := range pairCount {
		if pairCount == 0 {
			continue
		}
		complement := (maxVal - 1) ^ andVal
		subset := complement
		for {
			if subset < maxVal && freq[subset] > 0 {
				ans += pairCount * freq[subset]
			}
			if subset == 0 {
				break
			}
			subset = (subset - 1) & complement
		}
	}

	return ans
}

func main() {
	fmt.Println("Example 1:")
	fmt.Println(countTriplets([]int{2, 1, 3}))
	// Expected: 12

	fmt.Println("Example 2:")
	fmt.Println(countTriplets([]int{0, 0, 0}))
	// Expected: 27
}
```
