# 0077 — Combinations

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func combine(n int, k int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Backtracking

**Kompleksitas Waktu:** O(C(n,k) * k)  
**Kompleksitas Ruang:** O(k)

> **Untuk fresh graduate:** Kuasai dulu teknik **Backtracking** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #77: Combinations
// https://leetcode.com/problems/combinations/
// Difficulty: Medium

import "fmt"

func combine(n int, k int) [][]int {
	result := [][]int{}
	var backtrack func(start int, path []int)
	backtrack = func(start int, path []int) {
		if len(path) == k {
  // Alokasi slice integer
			comb := make([]int, k)
			copy(comb, path)
			result = append(result, comb)
			return
		}
		// Prune: if remaining numbers are not enough, skip
		for i := start; i <= n-(k-len(path))+1; i++ {
			path = append(path, i)
			backtrack(i+1, path)
			path = path[:len(path)-1]
		}
	}
	backtrack(1, []int{})
	return result
}

func main() {
	// Test case 1
	fmt.Println(combine(4, 2)) // [[1 2] [1 3] [1 4] [2 3] [2 4] [3 4]]

	// Test case 2
	fmt.Println(combine(1, 1)) // [[1]]

	// Test case 3
	fmt.Println(combine(4, 4)) // [[1 2 3 4]]
}

// Time: O(C(n,k) * k) | Space: O(k)
```
