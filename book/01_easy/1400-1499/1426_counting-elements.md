# 1426 — Counting Elements

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func countElements(arr []int) int

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

// LeetCode #1426: Counting Elements
// https://leetcode.com/problems/counting-elements/
// Difficulty: Easy [Paid]
//
// LeetCode submission: func countElements(arr []int) int

import "fmt"

func main() {
	fmt.Println(CountingElements([]int{1, 2, 3}))       // 2
	fmt.Println(CountingElements([]int{1, 1, 3, 3, 5, 5, 7, 7})) // 0
	fmt.Println(CountingElements([]int{1, 1, 2, 2}))    // 2
}

// Time: O(n), Space: O(n)
func CountingElements(arr []int) int {
  // Membuat map (HashMap) — pencarian O(1)
	seen := make(map[int]bool, len(arr))
	for _, v := range arr {
		seen[v] = true
	}
	count := 0
	for _, v := range arr {
		if seen[v+1] {
			count++
		}
	}
	return count
}
```
