# 0281 — Zigzag Iterator

## Deskripsi

**Soal:** [0281. Zigzag Iterator](https://leetcode.com/problems/zigzag-iterator/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(1) amortized per next/hasNext, Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func Constructor(v1, v2 []int) *ZigzagIterator`

## Solusi Go

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
  // Membuat slice untuk menyimpan hasil
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
