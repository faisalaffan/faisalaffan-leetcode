# 1756 — Design Most Recently Used Queue

## Deskripsi

**Soal:** [1756. Design Most Recently Used Queue](https://leetcode.com/problems/design-most-recently-used-queue/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n) per operation, Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Queue (antrian FIFO)

**Fungsi Solusi:** `func Constructor(n int) MRUQueue`

## Solusi Go

```go
package main

// LeetCode #1756: Design Most Recently Used Queue
// https://leetcode.com/problems/design-most-recently-used-queue/
// Difficulty: Medium [Paid]
// Time: O(n) per operation, Space: O(n)

import "fmt"

type MRUQueue struct {
	data []int
}

func Constructor(n int) MRUQueue {
  // Membuat slice untuk menyimpan hasil
	data := make([]int, n)
	for i := 0; i < n; i++ {
		data[i] = i + 1
	}
	return MRUQueue{data}
}

func (q *MRUQueue) Fetch(k int) int {
	// 1-indexed, fetch kth element and move to end
	val := q.data[k-1]
	// Remove
	q.data = append(q.data[:k-1], q.data[k:]...)
	// Move to end
	q.data = append(q.data, val)
	return val
}

func main() {
	q := Constructor(8)
	fmt.Println(q.Fetch(3)) // Expected: 3
	fmt.Println(q.Fetch(5)) // Expected: 6
	fmt.Println(q.Fetch(2)) // Expected: 2
	fmt.Println(q.Fetch(8)) // Expected: 3
}
```
