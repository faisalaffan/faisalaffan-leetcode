# 0594 — Longest Harmonious Subsequence

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func LongestHarmoniousSubsequence(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #594: Longest Harmonious Subsequence
// https://leetcode.com/problems/longest-harmonious-subsequence/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(n)
func LongestHarmoniousSubsequence(nums []int) int {
  // Membuat map (HashMap) — pencarian O(1)
	count := make(map[int]int)
	for _, v := range nums {
		count[v]++
	}
	maxLen := 0
	for v, c := range count {
		if c2, ok := count[v+1]; ok {
			if c+c2 > maxLen {
				maxLen = c + c2
			}
		}
	}
	return maxLen
}

func main() {
	fmt.Println(LongestHarmoniousSubsequence([]int{1, 3, 2, 2, 5, 2, 3, 7}))
	fmt.Println(LongestHarmoniousSubsequence([]int{1, 2, 3, 4}))
	fmt.Println(LongestHarmoniousSubsequence([]int{1, 1, 1, 1}))
}
```
