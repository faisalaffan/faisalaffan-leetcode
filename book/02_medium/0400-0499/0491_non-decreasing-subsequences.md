# 0491 — Non Decreasing Subsequences

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func NonDecreasingSubsequences(nums []int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Backtracking

**Kompleksitas Waktu:** O(2^n * n) worst case  
**Kompleksitas Ruang:** O(2^n * n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #491: Non-decreasing Subsequences
// https://leetcode.com/problems/non-decreasing-subsequences/
// Difficulty: Medium
// Time: O(2^n * n) worst case
// Space: O(2^n * n)

import "fmt"

func main() {
	fmt.Println(NonDecreasingSubsequences([]int{4, 6, 7, 7}))
	fmt.Println(NonDecreasingSubsequences([]int{4, 4, 3, 2, 1}))
}

func NonDecreasingSubsequences(nums []int) [][]int {
	result := [][]int{}
	var backtrack func(start int, path []int)
	backtrack = func(start int, path []int) {
		if len(path) >= 2 {
  // Alokasi slice integer
			cp := make([]int, len(path))
			copy(cp, path)
			result = append(result, cp)
		}
  // Membuat map (HashMap) — pencarian O(1)
		used := make(map[int]bool)
		for i := start; i < len(nums); i++ {
			if used[nums[i]] {
				continue
			}
			if len(path) == 0 || nums[i] >= path[len(path)-1] {
				used[nums[i]] = true
				backtrack(i+1, append(path, nums[i]))
			}
		}
	}
	backtrack(0, []int{})
	return result
}
```
