# 2154 — Keep Multiplying Found Values By Two

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func KeepMultiplyingFoundValuesByTwo(nums []int, original int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2154: Keep Multiplying Found Values by Two
// https://leetcode.com/problems/keep-multiplying-found-values-by-two/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(KeepMultiplyingFoundValuesByTwo([]int{5, 3, 6, 1, 12}, 3))  // 24
	fmt.Println(KeepMultiplyingFoundValuesByTwo([]int{2, 7, 9}, 4))          // 4
}

// Time: O(n), Space: O(n)
func KeepMultiplyingFoundValuesByTwo(nums []int, original int) int {
  // Membuat map (HashMap) — pencarian O(1)
	set := make(map[int]bool)
	for _, v := range nums {
		set[v] = true
	}

	for set[original] {
		original *= 2
	}
	return original
}
```
