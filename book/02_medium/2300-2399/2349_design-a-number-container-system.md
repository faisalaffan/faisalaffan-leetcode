# 2349 — Design A Number Container System

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func Constructor2() NumberContainers
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Heap / Priority Queue, Stack

**Kompleksitas Waktu:** O(log n) per operation  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2349: Design a Number Container System
// https://leetcode.com/problems/design-a-number-container-system/
// Difficulty: Medium
// Time: O(log n) per operation | Space: O(n)

import (
	"container/heap"
	"fmt"
)

type NumberContainers struct {
	indexToNum map[int]int
	numToIndex map[int]*minHeap3
}

type minHeap3 []int

func (h minHeap3) Len() int           { return len(h) }
func (h minHeap3) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap3) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap3) Push(x any)        { *h = append(*h, x.(int)) }
func (h *minHeap3) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func Constructor2() NumberContainers {
	return NumberContainers{
		indexToNum: make(map[int]int),
		numToIndex: make(map[int]*minHeap3),
	}
}

func (this *NumberContainers) Change(index int, number int) {
	this.indexToNum[index] = number
	if this.numToIndex[number] == nil {
		this.numToIndex[number] = &minHeap3{}
	}
  // Masukkan elemen ke priority queue
	heap.Push(this.numToIndex[number], index)
}

func (this *NumberContainers) Find(number int) int {
	if this.numToIndex[number] == nil {
		return -1
	}
	for this.numToIndex[number].Len() > 0 {
		idx := (*this.numToIndex[number])[0]
		if this.indexToNum[idx] == number {
			return idx
		}
  // Ambil elemen terkecil/terbesar dari heap
		heap.Pop(this.numToIndex[number])
	}
	return -1
}

func main() {
	nc := Constructor2()
	nc.Change(1, 10)
	fmt.Println(nc.Find(10)) // 1
	nc.Change(1, 20)
	fmt.Println(nc.Find(10)) // -1
	fmt.Println(nc.Find(20)) // 1
	fmt.Println(nc.Find(30)) // -1
}
```
