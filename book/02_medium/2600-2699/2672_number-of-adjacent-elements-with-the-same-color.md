# 2672 — Number Of Adjacent Elements With The Same Color

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func colorTheArray(n int, queries [][]int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n + q)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2672: Number of Adjacent Elements With the Same Color
// https://leetcode.com/problems/number-of-adjacent-elements-with-the-same-color/
// Difficulty: Medium
// Time: O(n + q) | Space: O(n)

import "fmt"

func colorTheArray(n int, queries [][]int) []int {
  // Alokasi slice integer
	colors := make([]int, n)
  // Alokasi slice integer
	ans := make([]int, len(queries))
	pairs := 0

	for i, q := range queries {
		idx, color := q[0], q[1]

		// Remove existing pairs
		if colors[idx] != 0 {
			if idx > 0 && colors[idx-1] == colors[idx] {
				pairs--
			}
			if idx < n-1 && colors[idx+1] == colors[idx] {
				pairs--
			}
		}

		colors[idx] = color

		// Add new pairs
		if idx > 0 && colors[idx-1] == colors[idx] {
			pairs++
		}
		if idx < n-1 && colors[idx+1] == colors[idx] {
			pairs++
		}

		ans[i] = pairs
	}
	return ans
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", colorTheArray(4, [][]int{{0, 2}, {1, 2}, {3, 1}, {1, 1}, {2, 1}}))
	// Expected: [0,1,1,0,2]

	// Test case 2
	fmt.Println("Test 2:", colorTheArray(1, [][]int{{0, 1}}))
	// Expected: [0]

	// Test case 3
	fmt.Println("Test 3:", colorTheArray(3, [][]int{{0, 1}, {1, 1}, {2, 1}}))
	// Expected: [0,1,2]
}
```
