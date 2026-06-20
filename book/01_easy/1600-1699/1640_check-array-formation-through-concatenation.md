# 1640 — Check Array Formation Through Concatenation

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func CanFormArray(arr []int, pieces [][]int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1640: Check Array Formation Through Concatenation
// https://leetcode.com/problems/check-array-formation-through-concatenation/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(n)
func CanFormArray(arr []int, pieces [][]int) bool {
  // Membuat map (HashMap) — pencarian O(1)
	pos := make(map[int]int)
	for i, num := range arr {
		pos[num] = i
	}
	for _, piece := range pieces {
		first := piece[0]
		idx, ok := pos[first]
		if !ok {
			return false
		}
		for j := 1; j < len(piece); j++ {
			if idx+j >= len(arr) || arr[idx+j] != piece[j] {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(CanFormArray([]int{15, 88}, [][]int{{88}, {15}}))
	fmt.Println(CanFormArray([]int{49, 18, 16}, [][]int{{16, 18, 49}}))
	fmt.Println(CanFormArray([]int{91, 4, 64, 78}, [][]int{{78}, {4, 64}, {91}}))
}
```
