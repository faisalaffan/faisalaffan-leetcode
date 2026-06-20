# 0281 — Zigzag Iterator

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func Constructor(v1, v2 []int) *ZigzagIterator
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1) amortized per next/hasNext, Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #281: Zigzag Iterator
// https://leetcode.com/problems/zigzag-iterator/
// Difficulty: Medium [Paid]
// Time: O(1) amortized per next/hasNext, Space: O(n)

import "fmt"

type ZigzagIterator struct {
	vectors [][]int
	indices []int
	curr    int
	total   int
}

func Constructor(v1, v2 []int) *ZigzagIterator {
	vectors := [][]int{v1, v2}
  // Alokasi slice integer
	indices := make([]int, 2)
	total := 0
	for _, v := range vectors {
		total += len(v)
	}
	return &ZigzagIterator{vectors, indices, 0, total}
}

func (this *ZigzagIterator) next() int {
	for this.indices[this.curr] >= len(this.vectors[this.curr]) {
		this.curr = (this.curr + 1) % 2
	}
	val := this.vectors[this.curr][this.indices[this.curr]]
	this.indices[this.curr]++
	this.curr = (this.curr + 1) % 2
	this.total--
	return val
}

func (this *ZigzagIterator) hasNext() bool {
	return this.total > 0
}

func main() {
	iter := Constructor([]int{1, 2, 3}, []int{4, 5, 6, 7})
	for iter.hasNext() {
		fmt.Print(iter.next(), " ")
	}
	fmt.Println()

	iter2 := Constructor([]int{1}, []int{})
	for iter2.hasNext() {
		fmt.Print(iter2.next(), " ")
	}
	fmt.Println()
}
```
