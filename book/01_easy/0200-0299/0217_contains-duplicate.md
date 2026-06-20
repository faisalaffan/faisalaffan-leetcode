# 0217 — Contains Duplicate

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func ContainsDuplicate(nums []int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #217: Contains Duplicate
// https://leetcode.com/problems/contains-duplicate/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(n)
func ContainsDuplicate(nums []int) bool {
  // Membuat map (HashMap) — pencarian O(1)
	seen := make(map[int]struct{}, len(nums))
	for _, n := range nums {
		if _, ok := seen[n]; ok {
			return true
		}
		seen[n] = struct{}{}
	}
	return false
}

func main() {
	fmt.Println(ContainsDuplicate([]int{1, 2, 3, 1}))
	fmt.Println(ContainsDuplicate([]int{1, 2, 3, 4}))
	fmt.Println(ContainsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
}
```
