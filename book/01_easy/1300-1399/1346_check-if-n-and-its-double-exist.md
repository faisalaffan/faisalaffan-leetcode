# 1346 — Check If N And Its Double Exist

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func checkIfExist(arr []int) bool

import "fmt"

func main()
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1346: Check If N and Its Double Exist
// https://leetcode.com/problems/check-if-n-and-its-double-exist/
// Difficulty: Easy
//
// LeetCode submission: func checkIfExist(arr []int) bool

import "fmt"

func main() {
	fmt.Println(CheckIfNAndItsDoubleExist([]int{10, 2, 5, 3}))  // true (10 = 2*5)
	fmt.Println(CheckIfNAndItsDoubleExist([]int{3, 1, 7, 11}))   // false
	fmt.Println(CheckIfNAndItsDoubleExist([]int{7, 1, 14, 11}))  // true (14 = 2*7)
}

// Time: O(n), Space: O(n)
func CheckIfNAndItsDoubleExist(arr []int) bool {
  // Membuat map (HashMap) — pencarian O(1)
	seen := make(map[int]bool, len(arr))
	for _, v := range arr {
		if seen[v*2] || (v%2 == 0 && seen[v/2]) {
			return true
		}
		seen[v] = true
	}
	return false
}
```
