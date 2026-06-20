# 0284 — Peeking Iterator

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func Constructor(iter *Iterator) *PeekingIterator`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(1) per operation, Space: O(1)  |  **Ruang:** O(1)


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
