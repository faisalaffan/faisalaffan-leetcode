# 0225 — Implement Stack Using Queues

## Deskripsi

**Soal:** [0225. Implement Stack Using Queues](https://leetcode.com/problems/implement-stack-using-queues/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Stack (tumpukan LIFO), Queue (antrian FIFO)

**Fungsi Solusi:** `func Constructor() MyStack`

## Solusi Go

```go
package main

// LeetCode #225: Implement Stack using Queues
// https://leetcode.com/problems/implement-stack-using-queues/
// Difficulty: Easy

import "fmt"

type MyStack struct {
	q []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (s *MyStack) Push(x int) {
	s.q = append(s.q, x)
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(s.q)-1; i++ {
		s.q = append(s.q, s.q[0])
		s.q = s.q[1:]
	}
}

func (s *MyStack) Pop() int {
	x := s.q[0]
	s.q = s.q[1:]
	return x
}

func (s *MyStack) Top() int {
	return s.q[0]
}

func (s *MyStack) Empty() bool {
	return len(s.q) == 0
}

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	fmt.Println(obj.Top())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Empty())
}
```
