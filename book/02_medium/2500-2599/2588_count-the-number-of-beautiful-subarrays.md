# 2588 — Count The Number Of Beautiful Subarrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func beautifulSubarrays(nums []int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Prefix Sum

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2588: Count the Number of Beautiful Subarrays
// https://leetcode.com/problems/count-the-number-of-beautiful-subarrays/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func beautifulSubarrays(nums []int) int64 {
  // Membuat map (HashMap) — pencarian O(1)
	prefixXor := make(map[int]int)
	prefixXor[0] = 1
	xor := 0
	var ans int64

	for _, v := range nums {
		xor ^= v
		ans += int64(prefixXor[xor])
		prefixXor[xor]++
	}
	return ans
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", beautifulSubarrays([]int{4, 3, 1, 2, 4}))
	// Expected: 2

	// Test case 2
	fmt.Println("Test 2:", beautifulSubarrays([]int{1, 10, 4}))
	// Expected: 0

	// Test case 3
	fmt.Println("Test 3:", beautifulSubarrays([]int{0, 0, 0}))
	// Expected: 6
}
```
