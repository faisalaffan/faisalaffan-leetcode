# 1381 — Design A Stack With Increment Operation

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diminta untuk mendesain (merancang) sebuah struktur data kustom dengan operasi tertentu (insert, delete, search, update). Tugasmu adalah memilih representasi data yang tepat agar setiap operasi berjalan efisien — biasanya O(1) atau O(log n).

Ini adalah soal yang paling sering muncul di interview sistem desain. Kamu perlu memilih kombinasi struktur data yang tepat (HashMap + Heap + LinkedList) untuk mencapai kompleksitas yang diminta.

**Konsep kunci:** HashMap (O(1) lookup), Heap (priority), Doubly Linked List (O(1) remove), TreeMap (ordered keys).

**Fungsi yang perlu kamu implementasikan:**
```go
func Constructor(maxSize int) CustomStack
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Stack

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Stack** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1381: Design a Stack With Increment Operation
// https://leetcode.com/problems/design-a-stack-with-increment-operation/
// Difficulty: Medium

import "fmt"

type CustomStack struct {
	stack []int
	inc   []int // lazy increment array
}

func main() {
	cs := Constructor(3)
	cs.Push(1)
	cs.Push(2)
	fmt.Println(cs.Pop()) // 2
	cs.Push(2)
	cs.Push(3)
	cs.Push(4)
	cs.Increment(5, 100)
	cs.Increment(2, 100)
	fmt.Println(cs.Pop()) // 103
	fmt.Println(cs.Pop()) // 202
	fmt.Println(cs.Pop()) // 201
	fmt.Println(cs.Pop()) // -1

	cs2 := Constructor(2)
	cs2.Push(1)
	cs2.Increment(1, 100)
	cs2.Increment(1, 100)
	fmt.Println(cs2.Pop()) // 201
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{
		stack: make([]int, 0, maxSize),
		inc:   make([]int, 0, maxSize),
	}
}

// Time: O(1)
func (this *CustomStack) Push(x int) {
	if len(this.stack) < cap(this.stack) {
		this.stack = append(this.stack, x)
		this.inc = append(this.inc, 0)
	}
}

// Time: O(1)
func (this *CustomStack) Pop() int {
	if len(this.stack) == 0 {
		return -1
	}
	n := len(this.stack)
	top := this.stack[n-1] + this.inc[n-1]
	if n > 1 {
		this.inc[n-2] += this.inc[n-1]
	}
	this.stack = this.stack[:n-1]
	this.inc = this.inc[:n-1]
	return top
}

// Time: O(1) using lazy increment
func (this *CustomStack) Increment(k int, val int) {
	if len(this.stack) == 0 {
		return
	}
	idx := k - 1
	if idx >= len(this.stack) {
		idx = len(this.stack) - 1
	}
	this.inc[idx] += val
}
```
