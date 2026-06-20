# 0046 — Permutations

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func permute(nums []int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Backtracking

**Kompleksitas Waktu:** O(n * n!)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Backtracking** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #46: Permutations
// https://leetcode.com/problems/permutations/
// Difficulty: Medium

import "fmt"

func permute(nums []int) [][]int {
	result := [][]int{}
	n := len(nums)
	var backtrack func(path []int, used []bool)
	backtrack = func(path []int, used []bool) {
		if len(path) == n {
  // Alokasi slice integer
			perm := make([]int, n)
			copy(perm, path)
			result = append(result, perm)
			return
		}
		for i := 0; i < n; i++ {
			if used[i] {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			backtrack(path, used)
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	backtrack([]int{}, make([]bool, n))
	return result
}

func main() {
	// Test case 1
	fmt.Println(permute([]int{1, 2, 3})) // [[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 1 2] [3 2 1]]

	// Test case 2
	fmt.Println(permute([]int{0, 1})) // [[0 1] [1 0]]

	// Test case 3
	fmt.Println(permute([]int{1})) // [[1]]
}

// Time: O(n * n!) | Space: O(n)
```
