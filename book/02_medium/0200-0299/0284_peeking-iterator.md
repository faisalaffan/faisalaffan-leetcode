# 0284 — Peeking Iterator

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func Constructor(iter *Iterator) *PeekingIterator
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1) per operation, Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #284: Peeking Iterator
// https://leetcode.com/problems/peeking-iterator/
// Difficulty: Medium
// Time: O(1) per operation, Space: O(1)

import "fmt"

type Iterator struct {
	data []int
	pos  int
}

func (this *Iterator) hasNext() bool {
	return this.pos < len(this.data)
}

func (this *Iterator) next() int {
	val := this.data[this.pos]
	this.pos++
	return val
}

type PeekingIterator struct {
	iter    *Iterator
	hasPeek bool
	peekVal int
}

func Constructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{iter, false, 0}
}

func (this *PeekingIterator) hasNext() bool {
	return this.hasPeek || this.iter.hasNext()
}

func (this *PeekingIterator) next() int {
	if this.hasPeek {
		this.hasPeek = false
		return this.peekVal
	}
	return this.iter.next()
}

func (this *PeekingIterator) peek() int {
	if !this.hasPeek {
		this.peekVal = this.iter.next()
		this.hasPeek = true
	}
	return this.peekVal
}

func main() {
	iter := &Iterator{[]int{1, 2, 3}, 0}
	pIter := Constructor(iter)
	fmt.Println(pIter.peek())
	fmt.Println(pIter.next())
	fmt.Println(pIter.peek())
	fmt.Println(pIter.hasNext())
	fmt.Println(pIter.next())
	fmt.Println(pIter.hasNext())
	fmt.Println(pIter.next())
	fmt.Println(pIter.hasNext())
}
```
