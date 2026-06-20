# 0641 — Design Circular Deque

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diminta untuk mendesain (merancang) sebuah struktur data kustom dengan operasi tertentu (insert, delete, search, update). Tugasmu adalah memilih representasi data yang tepat agar setiap operasi berjalan efisien — biasanya O(1) atau O(log n).

Ini adalah soal yang paling sering muncul di interview sistem desain. Kamu perlu memilih kombinasi struktur data yang tepat (HashMap + Heap + LinkedList) untuk mencapai kompleksitas yang diminta.

**Konsep kunci:** HashMap (O(1) lookup), Heap (priority), Doubly Linked List (O(1) remove), TreeMap (ordered keys).

**Fungsi yang perlu kamu implementasikan:**
```go
func Constructor(k int) MyCircularDeque
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1) for all operations  
**Kompleksitas Ruang:** O(k) where k is capacity

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

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
