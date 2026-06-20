# 3046 — Split The Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func SplitTheArray(nums []int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3046: Split the Array
// https://leetcode.com/problems/split-the-array/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: isPossibleToSplit
	fmt.Println(SplitTheArray([]int{1, 1, 2, 2, 3, 4})) // true
	fmt.Println(SplitTheArray([]int{1, 1, 1, 1}))       // false
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: isPossibleToSplit
// Each number can appear at most twice (once in each half of the split)
func SplitTheArray(nums []int) bool {
  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
		if freq[v] > 2 {
			return false
		}
	}
	return true
}
```
