# 2064 — Minimized Maximum Of Products Distributed To Any Store

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minimizedMaximum(n int, quantities []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log max(quantities))  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2064: Minimized Maximum of Products Distributed to Any Store
// https://leetcode.com/problems/minimized-maximum-of-products-distributed-to-any-store/
// Difficulty: Medium
// Time: O(n log max(quantities)) | Space: O(1)

import "fmt"

func minimizedMaximum(n int, quantities []int) int {
	canDistribute := func(maxProducts int) bool {
		stores := 0
		for _, q := range quantities {
			stores += (q + maxProducts - 1) / maxProducts
			if stores > n {
				return false
			}
		}
		return stores <= n
	}

	left, right := 1, 0
	for _, q := range quantities {
		if q > right {
			right = q
		}
	}

	result := right
	for left <= right {
		mid := left + (right-left)/2
		if canDistribute(mid) {
			result = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minimizedMaximum(6, []int{11, 6}))
	// Expected: 3

	// Test case 2
	fmt.Println("Test 2:", minimizedMaximum(7, []int{15, 10, 10}))
	// Expected: 5

	// Test case 3
	fmt.Println("Test 3:", minimizedMaximum(1, []int{100000}))
	// Expected: 100000
}
```
