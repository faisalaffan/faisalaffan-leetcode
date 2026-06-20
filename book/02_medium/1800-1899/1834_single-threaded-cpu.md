# 1834 — Single Threaded Cpu

## Deskripsi

**Soal:** [1834. Single Threaded Cpu](https://leetcode.com/problems/single-threaded-cpu/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Queue (antrian FIFO), Heap (priority queue)

**Fungsi Solusi:** `func getOrder(tasks [][]int) []int`

## Solusi Go

```go
package main

// LeetCode #1834: Single-Threaded CPU
// https://leetcode.com/problems/single-threaded-cpu/
// Difficulty: Medium
// Time: O(n log n), Space: O(n)

import (
	"container/heap"
	"fmt"
	"sort"
)

type Task struct {
	index       int
	enqueueTime int
	processTime int
}

type MinHeap []Task

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool {
	if h[i].processTime != h[j].processTime {
		return h[i].processTime < h[j].processTime
	}
	return h[i].index < h[j].index
}
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(Task)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func getOrder(tasks [][]int) []int {
	n := len(tasks)
  // Membuat slice untuk menyimpan hasil
	taskList := make([]Task, n)
	for i, t := range tasks {
		taskList[i] = Task{index: i, enqueueTime: t[0], processTime: t[1]}
	}

	sort.Slice(taskList, func(i, j int) bool {
		return taskList[i].enqueueTime < taskList[j].enqueueTime
	})

  // Membuat slice untuk menyimpan hasil
	result := make([]int, 0, n)
	pq := &MinHeap{}
	heap.Init(pq)
	time := 0
	i := 0

	for i < n || pq.Len() > 0 {
		// Add all available tasks
		for i < n && taskList[i].enqueueTime <= time {
			heap.Push(pq, taskList[i])
			i++
		}
		if pq.Len() == 0 {
			time = taskList[i].enqueueTime
			continue
		}
		t := heap.Pop(pq).(Task)
		result = append(result, t.index)
		time += t.processTime
	}
	return result
}

func main() {
	fmt.Println(getOrder([][]int{{1, 2}, {2, 4}, {3, 2}, {4, 1}})) // Expected: [0, 2, 3, 1]
	fmt.Println(getOrder([][]int{{7, 10}, {7, 12}, {7, 5}, {7, 4}, {7, 2}})) // Expected: [4, 3, 2, 0, 1]
	fmt.Println(getOrder([][]int{{5, 2}, {7, 2}, {9, 4}, {6, 3}, {5, 10}, {1, 1}})) // Expected: [5, 0, 1, 3, 2, 4]
}
```
