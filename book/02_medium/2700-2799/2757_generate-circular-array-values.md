# 2757 — Generate Circular Array Values

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func GenerateCircularArrayValues(arr []int, start int, count int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2757: Generate Circular Array Values
// https://leetcode.com/problems/generate-circular-array-values/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func GenerateCircularArrayValues(arr []int, start int, count int) []int {
	n := len(arr)
  // Edge case: input kosong — langsung return
	if n == 0 {
		return []int{}
	}
  // Alokasi slice integer
	result := make([]int, count)
	for i := 0; i < count; i++ {
		result[i] = arr[(start+i)%n]
	}
	return result
}

func main() {
	fmt.Println(GenerateCircularArrayValues([]int{1, 2, 3, 4}, 2, 6))
	fmt.Println(GenerateCircularArrayValues([]int{10, 20}, 1, 3))
}
```
