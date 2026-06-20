# 2680 — Maximum Or

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maximumOr(nums []int, k int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2680: Maximum OR
// https://leetcode.com/problems/maximum-or/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func maximumOr(nums []int, k int) int64 {
	n := len(nums)

	// suffix[i] = OR of nums from i to n-1
  // Alokasi slice integer
	suffix := make([]int64, n+1)
	for i := n - 1; i >= 0; i-- {
		suffix[i] = suffix[i+1] | int64(nums[i])
	}

	var prefix int64
	var ans int64

	for i := 0; i < n; i++ {
		candidate := prefix | (int64(nums[i]) << uint(k)) | suffix[i+1]
		if candidate > ans {
			ans = candidate
		}
		prefix |= int64(nums[i])
	}
	return ans
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", maximumOr([]int{12, 9}, 1))
	// Expected: 30

	// Test case 2
	fmt.Println("Test 2:", maximumOr([]int{8, 1, 2}, 2))
	// Expected: 35

	// Test case 3
	fmt.Println("Test 3:", maximumOr([]int{10, 8, 4}, 1))
	// Expected: 30
}
```
