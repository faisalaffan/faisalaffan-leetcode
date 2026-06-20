# 2206 — Divide Array Into Equal Pairs

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func DivideArrayIntoEqualPairs(nums []int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2206: Divide Array Into Equal Pairs
// https://leetcode.com/problems/divide-array-into-equal-pairs/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(DivideArrayIntoEqualPairs([]int{3, 2, 3, 2, 2, 2})) // true
	fmt.Println(DivideArrayIntoEqualPairs([]int{1, 2, 3, 4}))       // false
}

// Time: O(n), Space: O(n)
func DivideArrayIntoEqualPairs(nums []int) bool {
  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}
	for _, c := range freq {
		if c%2 != 0 {
			return false
		}
	}
	return true
}
```
