# 1860 — Incremental Memory Leak

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MemLeak(memory1 int, memory2 int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(sqrt(memory1+memory2)), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1860: Incremental Memory Leak
// https://leetcode.com/problems/incremental-memory-leak/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MemLeak(2, 2))
	fmt.Println(MemLeak(8, 11))
	fmt.Println(MemLeak(1, 1))
}

// Time: O(sqrt(memory1+memory2)), Space: O(1)
func MemLeak(memory1 int, memory2 int) []int {
	t := 1
	for memory1 >= t || memory2 >= t {
		if memory1 >= memory2 {
			memory1 -= t
		} else {
			memory2 -= t
		}
		t++
	}
	return []int{t, memory1, memory2}
}
```
