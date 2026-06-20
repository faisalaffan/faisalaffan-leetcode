# 0622 — Design Circular Queue

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diminta untuk mendesain (merancang) sebuah struktur data kustom dengan operasi tertentu (insert, delete, search, update). Tugasmu adalah memilih representasi data yang tepat agar setiap operasi berjalan efisien — biasanya O(1) atau O(log n).

Ini adalah soal yang paling sering muncul di interview sistem desain. Kamu perlu memilih kombinasi struktur data yang tepat (HashMap + Heap + LinkedList) untuk mencapai kompleksitas yang diminta.

**Konsep kunci:** HashMap (O(1) lookup), Heap (priority), Doubly Linked List (O(1) remove), TreeMap (ordered keys).

**Fungsi yang perlu kamu implementasikan:**
```go
func Constructor(k int) MyCircularQueue
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** BFS

**Kompleksitas Waktu:** O(1) for all operations  
**Kompleksitas Ruang:** O(k)

> **Untuk fresh graduate:** Kuasai dulu teknik **BFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #622: Design Circular Queue
// https://leetcode.com/problems/design-circular-queue/
// Difficulty: Medium
// Time: O(1) for all operations
// Space: O(k)

import "fmt"

type MyCircularQueue struct {
	data  []int
	front int
	rear  int
	size  int
	cap   int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		data: make([]int, k),
		front: 0,
		rear:  -1,
		size:  0,
		cap:   k,
	}
}

func (q *MyCircularQueue) EnQueue(value int) bool {
	if q.IsFull() {
		return false
	}
	q.rear = (q.rear + 1) % q.cap
	q.data[q.rear] = value
	q.size++
	return true
}

func (q *MyCircularQueue) DeQueue() bool {
	if q.IsEmpty() {
		return false
	}
	q.front = (q.front + 1) % q.cap
	q.size--
	return true
}

func (q *MyCircularQueue) Front() int {
	if q.IsEmpty() {
		return -1
	}
	return q.data[q.front]
}

func (q *MyCircularQueue) Rear() int {
	if q.IsEmpty() {
		return -1
	}
	return q.data[q.rear]
}

func (q *MyCircularQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *MyCircularQueue) IsFull() bool {
	return q.size == q.cap
}

func main() {
	q := Constructor(3)
	fmt.Println(q.EnQueue(1))
	fmt.Println(q.EnQueue(2))
	fmt.Println(q.EnQueue(3))
	fmt.Println(q.EnQueue(4))
	fmt.Println(q.Rear())
	fmt.Println(q.IsFull())
	fmt.Println(q.DeQueue())
	fmt.Println(q.EnQueue(4))
	fmt.Println(q.Rear())
}
```
