# 0251 — Flatten 2D Vector

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func Constructor(vec [][]int) Vector2D
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1) amortized per next/hasNext, Space: O(1) excluding input  
**Kompleksitas Ruang:** O(1) excluding input

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #251: Flatten 2D Vector
// https://leetcode.com/problems/flatten-2d-vector/
// Difficulty: Medium [Paid]
// Time: O(1) amortized per next/hasNext, Space: O(1) excluding input

import "fmt"

type Vector2D struct {
	vec    [][]int
	row    int
	col    int
}

func Constructor(vec [][]int) Vector2D {
	return Vector2D{vec, 0, 0}
}

func (this *Vector2D) advance() {
	for this.row < len(this.vec) && this.col >= len(this.vec[this.row]) {
		this.row++
		this.col = 0
	}
}

func (this *Vector2D) Next() int {
	this.advance()
	val := this.vec[this.row][this.col]
	this.col++
	return val
}

func (this *Vector2D) HasNext() bool {
	this.advance()
	return this.row < len(this.vec)
}

func main() {
	iter := Constructor([][]int{{1, 2}, {3}, {4, 5, 6}})
	for iter.HasNext() {
		fmt.Print(iter.Next(), " ")
	}
	fmt.Println()

	iter2 := Constructor([][]int{{}, {1}, {}})
	for iter2.HasNext() {
		fmt.Print(iter2.Next(), " ")
	}
	fmt.Println()

	iter3 := Constructor([][]int{{}})
	fmt.Println(iter3.HasNext())
}
```
