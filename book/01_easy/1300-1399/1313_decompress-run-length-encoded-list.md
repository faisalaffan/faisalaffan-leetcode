# 1313 — Decompress Run Length Encoded List

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func decompressRLElist(nums []int) []int

import "fmt"

func main()
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n + totalLen), Space: O(totalLen)  
**Kompleksitas Ruang:** O(totalLen)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1313: Decompress Run-Length Encoded List
// https://leetcode.com/problems/decompress-run-length-encoded-list/
// Difficulty: Easy
//
// LeetCode submission: func decompressRLElist(nums []int) []int

import "fmt"

func main() {
	fmt.Println(DecompressRunLengthEncodedList([]int{1, 2, 3, 4}))       // [2 4 4 4]
	fmt.Println(DecompressRunLengthEncodedList([]int{1, 1, 2, 3}))       // [1 3 3]
	fmt.Println(DecompressRunLengthEncodedList([]int{2, 5, 1, 7, 3, 9})) // [5 5 7 9 9 9]
}

// Time: O(n + totalLen), Space: O(totalLen)
func DecompressRunLengthEncodedList(nums []int) []int {
	totalLen := 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(nums); i += 2 {
		totalLen += nums[i]
	}
  // Alokasi slice integer
	res := make([]int, 0, totalLen)
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(nums); i += 2 {
		freq, val := nums[i], nums[i+1]
		for j := 0; j < freq; j++ {
			res = append(res, val)
		}
	}
	return res
}
```
