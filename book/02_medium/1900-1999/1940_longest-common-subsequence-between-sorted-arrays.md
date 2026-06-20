# 1940 — Longest Common Subsequence Between Sorted Arrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func LongestCommonSubsequenceBetweenSortedArrays(arrs [][]int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(total elements), Space: O(unique elements)  
**Kompleksitas Ruang:** O(unique elements)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1940: Longest Common Subsequence Between Sorted Arrays
// https://leetcode.com/problems/longest-common-subsequence-between-sorted-arrays/
// Difficulty: Medium [Paid]

import "fmt"

func main() {
	fmt.Println(LongestCommonSubsequenceBetweenSortedArrays([][]int{{1, 3, 4}, {1, 4, 7, 9}}))
	fmt.Println(LongestCommonSubsequenceBetweenSortedArrays([][]int{{2, 3, 6, 8}, {1, 2, 3, 5, 6, 7, 10}, {2, 3, 4, 6, 9}}))
}

// Time: O(total elements), Space: O(unique elements)
func LongestCommonSubsequenceBetweenSortedArrays(arrs [][]int) []int {
	// Since arrays are sorted, use frequency counting
	// Numbers appearing in ALL arrays are the answer
  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[int]int)
	for _, arr := range arrs {
		for _, v := range arr {
			freq[v]++
		}
	}

	n := len(arrs)
  // Alokasi slice integer
	result := make([]int, 0)
	for _, v := range arrs[0] {
		if freq[v] == n {
			result = append(result, v)
		}
	}
	return result
}
```
