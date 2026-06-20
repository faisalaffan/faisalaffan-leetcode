# 0254 — Factor Combinations

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func getFactors(n int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Backtracking

**Kompleksitas Waktu:** O(n log n), Space: O(log n)  
**Kompleksitas Ruang:** O(log n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Backtracking** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #254: Factor Combinations
// https://leetcode.com/problems/factor-combinations/
// Difficulty: Medium [Paid]
// Time: O(n log n), Space: O(log n)

import "fmt"

func getFactors(n int) [][]int {
	result := [][]int{}
	var backtrack func(start, remaining int, path []int)
	backtrack = func(start, remaining int, path []int) {
		if len(path) > 0 {
  // Alokasi slice integer
			combo := make([]int, len(path)+1)
			copy(combo, path)
			combo[len(combo)-1] = remaining
			result = append(result, combo)
		}

		for i := start; i*i <= remaining; i++ {
			if remaining%i == 0 {
				path = append(path, i)
				backtrack(i, remaining/i, path)
				path = path[:len(path)-1]
			}
		}
	}

	backtrack(2, n, []int{})
	return result
}

func main() {
	fmt.Println(getFactors(12))
	fmt.Println(getFactors(37))
	fmt.Println(getFactors(32))
}
```
