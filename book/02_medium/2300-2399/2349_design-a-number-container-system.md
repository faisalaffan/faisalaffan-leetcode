# 2349 — Design A Number Container System

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diminta mendesain struktur data kustom dengan operasi spesifik (insert, delete, search). Target: O(1) atau O(log n) per operasi.

**Cara berpikir:** Kombinasikan HashMap + Heap + Linked List sesuai kebutuhan.

**Fungsi Solusi:** `func Constructor2() NumberContainers`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Heap

**Waktu:** O(log n) per operation  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

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
  // Push ke priority queue
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
  // Pop dari priority queue
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
