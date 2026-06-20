# 2433 — Find The Original Array Of Prefix Xor

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func findArray(pref []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1) (excluding output)

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2433: Find The Original Array of Prefix Xor
// https://leetcode.com/problems/find-the-original-array-of-prefix-xor/
// Difficulty: Medium
// Time: O(n) | Space: O(1) (excluding output)
// Given pref[i] = XOR(arr[0..i]), find arr. arr[i] = pref[i] ^ pref[i-1].

import "fmt"

func main() {
	fmt.Println(findArray([]int{5, 2, 0, 3, 1})) // [5, 7, 2, 3, 2]
	fmt.Println(findArray([]int{13}))             // [13]
}

func findArray(pref []int) []int {
	n := len(pref)
  // Alokasi slice integer
	arr := make([]int, n)
	arr[0] = pref[0]
	for i := 1; i < n; i++ {
		arr[i] = pref[i] ^ pref[i-1]
	}
	return arr
}
```
