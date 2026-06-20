# 0128 — Longest Consecutive Sequence

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func longestConsecutive(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #128: Longest Consecutive Sequence
// https://leetcode.com/problems/longest-consecutive-sequence/
// Difficulty: Medium

import "fmt"

func longestConsecutive(nums []int) int {
  // Membuat map (HashMap) — pencarian O(1)
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	maxLen := 0
	for num := range numSet {
		if !numSet[num-1] {
			curr := num
			length := 1
			for numSet[curr+1] {
				curr++
				length++
			}
			if length > maxLen {
				maxLen = length
			}
		}
	}

	return maxLen
}

func main() {
	// Test case 1
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2})) // 4

	// Test case 2
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1})) // 9

	// Test case 3
	fmt.Println(longestConsecutive([]int{})) // 0
}

// Time: O(n) | Space: O(n)
```
