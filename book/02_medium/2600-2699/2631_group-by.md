# 2631 — Group By

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2631: Group By
// https://leetcode.com/problems/group-by/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func groupBy[T any, K comparable](items []T, keyFn func(T) K) map[K][]T {
  // Membuat map (HashMap) — pencarian O(1)
	result := make(map[K][]T)
	for _, item := range items {
		key := keyFn(item)
		result[key] = append(result[key], item)
	}
	return result
}

func main() {
	// Test case 1: group integers by even/odd
	nums := []int{1, 2, 3, 4, 5, 6}
	grouped := groupBy(nums, func(n int) string {
		if n%2 == 0 {
			return "even"
		}
		return "odd"
	})
	fmt.Println("Test 1:", grouped)
	// Expected: map[even:[2 4 6] odd:[1 3 5]]

	// Test case 2: group strings by length
	words := []string{"one", "two", "three", "four"}
	byLen := groupBy(words, func(s string) int {
		return len(s)
	})
	fmt.Println("Test 2:", byLen)
	// Expected: map[3:[one two] 5:[three four]]

	// Test case 3
	single := []int{1}
	result := groupBy(single, func(n int) string { return "a" })
	fmt.Println("Test 3:", result)
}
```
