# 3852 — Smallest Pair With Different Frequencies

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func SmallestPairWithDifferentFrequencies(nums []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3852: Smallest Pair With Different Frequencies
// https://leetcode.com/problems/smallest-pair-with-different-frequencies/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(SmallestPairWithDifferentFrequencies([]int{1, 1, 2, 2, 3, 4}))
	fmt.Println(SmallestPairWithDifferentFrequencies([]int{1, 5}))
	fmt.Println(SmallestPairWithDifferentFrequencies([]int{7}))
}

// Time: O(n)
// Space: O(n)
func SmallestPairWithDifferentFrequencies(nums []int) []int {
  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}
	minVal := nums[0]
	for _, v := range nums {
		if v < minVal {
			minVal = v
		}
	}
	minFreq := freq[minVal]
  // Membuat map (HashMap) — pencarian O(1)
	seen := make(map[int]bool)
	for _, v := range nums {
		if seen[v] {
			continue
		}
		seen[v] = true
		if v > minVal && freq[v] != minFreq {
			return []int{minVal, v}
		}
	}
	return []int{-1, -1}
}
```
