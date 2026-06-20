# 0641 — Design Circular Deque

## Deskripsi

**Soal:** [0641. Design Circular Deque](https://leetcode.com/problems/design-circular-deque/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(1) for all operations  
**Kompleksitas Ruang:** O(k) where k is capacity

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #641: Design Circular Deque
// https://leetcode.com/problems/design-circular-deque/
// Difficulty: Medium
// Time: O(1) for all operations
// Space: O(k) where k is capacity

import "fmt"

func main() {
	deque := Constructor(3)
	fmt.Println(deque.InsertLast(1))
	fmt.Println(deque.InsertLast(2))
	fmt.Println(deque.InsertFront(3))
	fmt.Println(deque.InsertFront(4))
	fmt.Println(deque.GetRear())
	fmt.Println(deque.IsFull())
	fmt.Println(deque.DeleteLast())
	fmt.Println(deque.InsertFront(4))
	fmt.Println(deque.GetFront())
}

type MyCircularDeque struct {
	data  []int
	front int
	rear  int
	size  int
	cap   int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		data:  make([]int, k),
		front: 0,
		rear:  0,
		size:  0,
		cap:   k,
	}
}

func (d *MyCircularDeque) InsertFront(value int) bool {
	if d.IsFull() {
		return false
	}
	d.front = (d.front - 1 + d.cap) % d.cap
	d.data[d.front] = value
	d.size++
	return true
}

func (d *MyCircularDeque) InsertLast(value int) bool {
	if d.IsFull() {
		return false
	}
	d.data[d.rear] = value
	d.rear = (d.rear + 1) % d.cap
	d.size++
	return true
}

func (d *MyCircularDeque) DeleteFront() bool {
	if d.IsEmpty() {
		return false
	}
	d.front = (d.front + 1) % d.cap
	d.size--
	return true
}

func (d *MyCircularDeque) DeleteLast() bool {
	if d.IsEmpty() {
		return false
	}
	d.rear = (d.rear - 1 + d.cap) % d.cap
	d.size--
	return true
}

func (d *MyCircularDeque) GetFront() int {
	if d.IsEmpty() {
		return -1
	}
	return d.data[d.front]
}

func (d *MyCircularDeque) GetRear() int {
	if d.IsEmpty() {
		return -1
	}
	return d.data[(d.rear-1+d.cap)%d.cap]
}

func (d *MyCircularDeque) IsEmpty() bool {
	return d.size == 0
}

func (d *MyCircularDeque) IsFull() bool {
	return d.size == d.cap
}
```
