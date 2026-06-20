# 3697 — Compute Decimal Representation

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func ComputeDecimalRepresentation(n int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log n)  
**Kompleksitas Ruang:** O(log n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3697: Compute Decimal Representation
// https://leetcode.com/problems/compute-decimal-representation/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ComputeDecimalRepresentation(537))
	fmt.Println(ComputeDecimalRepresentation(102))
	fmt.Println(ComputeDecimalRepresentation(6))
}

// Time: O(log n)
// Space: O(log n)
func ComputeDecimalRepresentation(n int) []int {
  // Alokasi slice integer
	res := make([]int, 0)
	place := 1
	for n > 0 {
		d := n % 10
		if d != 0 {
			res = append(res, d*place)
		}
		place *= 10
		n /= 10
	}

	// Reverse to descending order
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
```
