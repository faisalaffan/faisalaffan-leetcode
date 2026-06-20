# 0232 — Implement Queue Using Stacks

## Deskripsi

**Soal:** [0232. Implement Queue Using Stacks](https://leetcode.com/problems/implement-queue-using-stacks/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Stack (tumpukan LIFO), Queue (antrian FIFO)

**Fungsi Solusi:** `func Constructor() MyQueue`

## Solusi Go

```go
package main

// LeetCode #232: Implement Queue using Stacks
// https://leetcode.com/problems/implement-queue-using-stacks/
// Difficulty: Easy

import "fmt"

type MyQueue struct {
	in  []int
	out []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (q *MyQueue) Push(x int) {
	q.in = append(q.in, x)
}

func (q *MyQueue) transfer() {
	if len(q.out) == 0 {
		for len(q.in) > 0 {
			q.out = append(q.out, q.in[len(q.in)-1])
			q.in = q.in[:len(q.in)-1]
		}
	}
}

func (q *MyQueue) Pop() int {
	q.transfer()
	x := q.out[len(q.out)-1]
	q.out = q.out[:len(q.out)-1]
	return x
}

func (q *MyQueue) Peek() int {
	q.transfer()
	return q.out[len(q.out)-1]
}

func (q *MyQueue) Empty() bool {
	return len(q.in) == 0 && len(q.out) == 0
}

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	fmt.Println(obj.Peek())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Empty())
}
```
