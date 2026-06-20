# 1188 — Design Bounded Blocking Queue

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diminta untuk mendesain (merancang) sebuah struktur data kustom dengan operasi tertentu (insert, delete, search, update). Tugasmu adalah memilih representasi data yang tepat agar setiap operasi berjalan efisien — biasanya O(1) atau O(log n).

Ini adalah soal yang paling sering muncul di interview sistem desain. Kamu perlu memilih kombinasi struktur data yang tepat (HashMap + Heap + LinkedList) untuk mencapai kompleksitas yang diminta.

**Konsep kunci:** HashMap (O(1) lookup), Heap (priority), Doubly Linked List (O(1) remove), TreeMap (ordered keys).

**Fungsi yang perlu kamu implementasikan:**
```go
func NewBoundedBlockingQueue(capacity int) *BoundedBlockingQueue
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** BFS

**Kompleksitas Waktu:** O(1) per operation  
**Kompleksitas Ruang:** O(capacity)

> **Untuk fresh graduate:** Kuasai dulu teknik **BFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
	"sync"
)

// LeetCode #1188: Design Bounded Blocking Queue
// https://leetcode.com/problems/design-bounded-blocking-queue/
// Difficulty: Medium [Paid]

// Thread-safe bounded queue using mutex and condition variable.

// Time: O(1) per operation
// Space: O(capacity)

type BoundedBlockingQueue struct {
	capacity int
	queue    []int
	mu       sync.Mutex
	notFull  *sync.Cond
	notEmpty *sync.Cond
}

func NewBoundedBlockingQueue(capacity int) *BoundedBlockingQueue {
	bq := &BoundedBlockingQueue{
		capacity: capacity,
		queue:    make([]int, 0),
	}
	bq.notFull = sync.NewCond(&bq.mu)
	bq.notEmpty = sync.NewCond(&bq.mu)
	return bq
}

func (bq *BoundedBlockingQueue) Enqueue(element int) {
	bq.mu.Lock()
	defer bq.mu.Unlock()

	for len(bq.queue) >= bq.capacity {
		bq.notFull.Wait()
	}
	bq.queue = append(bq.queue, element)
	bq.notEmpty.Signal()
}

func (bq *BoundedBlockingQueue) Dequeue() int {
	bq.mu.Lock()
	defer bq.mu.Unlock()

	for len(bq.queue) == 0 {
		bq.notEmpty.Wait()
	}
	val := bq.queue[0]
	bq.queue = bq.queue[1:]
	bq.notFull.Signal()
	return val
}

func (bq *BoundedBlockingQueue) Size() int {
	bq.mu.Lock()
	defer bq.mu.Unlock()
	return len(bq.queue)
}

func main() {
	bq := NewBoundedBlockingQueue(2)
	bq.Enqueue(1)
	bq.Enqueue(2)
	fmt.Printf("size: %d (expected: 2)\n", bq.Size())
	fmt.Printf("dequeue: %d (expected: 1)\n", bq.Dequeue())
	fmt.Printf("size: %d (expected: 1)\n", bq.Size())
	bq.Enqueue(3)
	fmt.Printf("dequeue: %d (expected: 2)\n", bq.Dequeue())
	fmt.Printf("dequeue: %d (expected: 3)\n", bq.Dequeue())
}
```
