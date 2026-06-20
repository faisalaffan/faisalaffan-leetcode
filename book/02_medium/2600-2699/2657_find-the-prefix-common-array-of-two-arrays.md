# 2657 — Find The Prefix Common Array Of Two Arrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func findThePrefixCommonArray(A []int, B []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Prefix Sum

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2657: Find the Prefix Common Array of Two Arrays
// https://leetcode.com/problems/find-the-prefix-common-array-of-two-arrays/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func findThePrefixCommonArray(A []int, B []int) []int {
	n := len(A)
  // Alokasi slice integer
	ans := make([]int, n)
  // Membuat map (HashMap) — pencarian O(1)
	seen := make(map[int]bool)
	count := 0

	for i := 0; i < n; i++ {
		if seen[A[i]] {
			count++
		} else {
			seen[A[i]] = true
		}
		if seen[B[i]] {
			count++
		} else {
			seen[B[i]] = true
		}
		ans[i] = count
	}
	return ans
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", findThePrefixCommonArray([]int{1, 3, 2, 4}, []int{3, 1, 2, 4}))
	// Expected: [0,2,3,4]

	// Test case 2
	fmt.Println("Test 2:", findThePrefixCommonArray([]int{1, 2, 3}, []int{1, 2, 3}))
	// Expected: [1,2,3]

	// Test case 3
	fmt.Println("Test 3:", findThePrefixCommonArray([]int{1, 2, 3}, []int{3, 1, 2}))
	// Expected: [0,1,3]
}
```
