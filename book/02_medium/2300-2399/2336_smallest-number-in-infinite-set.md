# 2336 — Smallest Number In Infinite Set

## Deskripsi

**Soal:** [2336. Smallest Number In Infinite Set](https://leetcode.com/problems/smallest-number-in-infinite-set/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(log n) for pop, O(1) for addBack  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Heap (priority queue)

**Fungsi Solusi:** `func Constructor() SmallestInfiniteSet`

## Solusi Go

```go
package main

// LeetCode #2336: Smallest Number in Infinite Set
// https://leetcode.com/problems/smallest-number-in-infinite-set/
// Difficulty: Medium
// Time: O(log n) for pop, O(1) for addBack | Space: O(n)

import (
	"container/heap"
	"fmt"
)

type SmallestInfiniteSet struct {
	added     *minHeap2
	addedSet  map[int]bool
	smallest  int
}

type minHeap2 []int

func (h minHeap2) Len() int           { return len(h) }
func (h minHeap2) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap2) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap2) Push(x any)        { *h = append(*h, x.(int)) }
func (h *minHeap2) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func Constructor() SmallestInfiniteSet {
	return SmallestInfiniteSet{
		added:    &minHeap2{},
		addedSet: make(map[int]bool),
		smallest: 1,
	}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	if this.added.Len() > 0 {
		val := heap.Pop(this.added).(int)
		delete(this.addedSet, val)
		return val
	}
	val := this.smallest
	this.smallest++
	return val
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if num >= this.smallest || this.addedSet[num] {
		return
	}
	this.addedSet[num] = true
	heap.Push(this.added, num)
}

func main() {
	set := Constructor()
	set.AddBack(2)
	fmt.Println(set.PopSmallest()) // 1
	fmt.Println(set.PopSmallest()) // 2
	fmt.Println(set.PopSmallest()) // 3
	set.AddBack(1)
	fmt.Println(set.PopSmallest()) // 1
	fmt.Println(set.PopSmallest()) // 4
}
```
