# 1670 — Design Front Middle Back Queue

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diminta untuk mendesain (merancang) sebuah struktur data kustom dengan operasi tertentu (insert, delete, search, update). Tugasmu adalah memilih representasi data yang tepat agar setiap operasi berjalan efisien — biasanya O(1) atau O(log n).

Ini adalah soal yang paling sering muncul di interview sistem desain. Kamu perlu memilih kombinasi struktur data yang tepat (HashMap + Heap + LinkedList) untuk mencapai kompleksitas yang diminta.

**Konsep kunci:** HashMap (O(1) lookup), Heap (priority), Doubly Linked List (O(1) remove), TreeMap (ordered keys).

**Fungsi yang perlu kamu implementasikan:**
```go
func Constructor() FrontMiddleBackQueue
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** BFS

**Kompleksitas Waktu:** O(1) per operation (amortized), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **BFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1670: Design Front Middle Back Queue
// https://leetcode.com/problems/design-front-middle-back-queue/
// Difficulty: Medium
// Time: O(1) per operation (amortized), Space: O(n)

import "fmt"

type FrontMiddleBackQueue struct {
	left  []int
	right []int
}

func Constructor() FrontMiddleBackQueue {
	return FrontMiddleBackQueue{}
}

func (q *FrontMiddleBackQueue) balance() {
	// Keep left size >= right size, and left size - right size <= 1
	if len(q.left) > len(q.right)+1 {
		// Move last of left to front of right
		v := q.left[len(q.left)-1]
		q.left = q.left[:len(q.left)-1]
		q.right = append([]int{v}, q.right...)
	} else if len(q.right) > len(q.left) {
		// Move first of right to back of left
		v := q.right[0]
		q.right = q.right[1:]
		q.left = append(q.left, v)
	}
}

func (q *FrontMiddleBackQueue) PushFront(val int) {
	q.left = append([]int{val}, q.left...)
	q.balance()
}

func (q *FrontMiddleBackQueue) PushMiddle(val int) {
	q.left = append(q.left, val)
	q.balance()
}

func (q *FrontMiddleBackQueue) PushBack(val int) {
	q.right = append(q.right, val)
	q.balance()
}

func (q *FrontMiddleBackQueue) PopFront() int {
	if len(q.left) == 0 {
		return -1
	}
	val := q.left[0]
	q.left = q.left[1:]
	q.balance()
	return val
}

func (q *FrontMiddleBackQueue) PopMiddle() int {
	if len(q.left) == 0 {
		return -1
	}
	val := q.left[len(q.left)-1]
	q.left = q.left[:len(q.left)-1]
	q.balance()
	return val
}

func (q *FrontMiddleBackQueue) PopBack() int {
	if len(q.right) == 0 {
		if len(q.left) == 0 {
			return -1
		}
		val := q.left[len(q.left)-1]
		q.left = q.left[:len(q.left)-1]
		return val
	}
	val := q.right[len(q.right)-1]
	q.right = q.right[:len(q.right)-1]
	q.balance()
	return val
}

func main() {
	q := Constructor()
	q.PushFront(1)
	q.PushBack(2)
	q.PushMiddle(3)
	q.PushMiddle(4)

	fmt.Println("PopFront:", q.PopFront()) // 1
	fmt.Println("PopMiddle:", q.PopMiddle()) // 3
	fmt.Println("PopMiddle:", q.PopMiddle()) // 4
	fmt.Println("PopBack:", q.PopBack()) // 2
	fmt.Println("PopFront:", q.PopFront()) // -1 (empty)

	// Test case 2
	q2 := Constructor()
	q2.PushFront(1)
	q2.PushFront(2)
	q2.PushBack(3)
	q2.PushBack(4)
	fmt.Println("PopFront:", q2.PopFront()) // 2
	fmt.Println("PopBack:", q2.PopBack())   // 4
	fmt.Println("PopFront:", q2.PopFront()) // 1
}
```
