# 2219 — Maximum Sum Score Of Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maximumSumScore(nums []int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2219: Maximum Sum Score of Array
// https://leetcode.com/problems/maximum-sum-score-of-array/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func maximumSumScore(nums []int) int64 {
	n := len(nums)
	total := int64(0)
	for _, v := range nums {
		total += int64(v)
	}

	maxScore := int64(-1 << 63)
	var prefix int64 = 0
	for i := 0; i < n; i++ {
		prefix += int64(nums[i])
		suffix := total - prefix + int64(nums[i])
		score := prefix
		if suffix > score {
			score = suffix
		}
		if score > maxScore {
			maxScore = score
		}
	}
	return maxScore
}

func main() {
	// Test case 1
	fmt.Println(maximumSumScore([]int{1, 2, 3, 4, 5}))
	// Expected: 15

	// Test case 2
	fmt.Println(maximumSumScore([]int{-5, 10, -3, 4}))
	// Expected: 11

	// Test case 3
	fmt.Println(maximumSumScore([]int{-1, -2, -3}))
	// Expected: -1
}
```
