# 0155 — Min Stack

## Deskripsi

**Soal:** [0155. Min Stack](https://leetcode.com/problems/min-stack/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(1) per operation, Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Stack (tumpukan LIFO)

**Fungsi Solusi:** `func Constructor() MinStack`

## Solusi Go

```go
package main

// LeetCode #155: Min Stack
// https://leetcode.com/problems/min-stack/
// Difficulty: Medium
// Time: O(1) per operation, Space: O(n)

import "fmt"

type MinStack struct {
	stack []int
	min   []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if len(this.min) == 0 || val <= this.GetMin() {
		this.min = append(this.min, val)
	}
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	if this.Top() == this.GetMin() {
		this.min = this.min[:len(this.min)-1]
	}
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}

func main() {
	ms := Constructor()
	ms.Push(-2)
	ms.Push(0)
	ms.Push(-3)
	fmt.Println(ms.GetMin())
	ms.Pop()
	fmt.Println(ms.Top())
	fmt.Println(ms.GetMin())
}
```
