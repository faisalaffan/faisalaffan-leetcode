# 2677 — Chunk Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func ChunkArray(arr []int, size int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2677: Chunk Array
// https://leetcode.com/problems/chunk-array/
// Difficulty: Easy
// Time: O(n) | Space: O(n)
// Note: JavaScript problem, adapted to Go. Splits array into chunks of given size.

import "fmt"

func main() {
	fmt.Println(ChunkArray([]int{1, 2, 3, 4, 5}, 1))
	fmt.Println(ChunkArray([]int{1, 9, 6, 3, 2}, 3))
}

func ChunkArray(arr []int, size int) [][]int {
	var result [][]int
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(arr); i += size {
		end := i + size
		if end > len(arr) {
			end = len(arr)
		}
		result = append(result, arr[i:end])
	}
	return result
}
```
