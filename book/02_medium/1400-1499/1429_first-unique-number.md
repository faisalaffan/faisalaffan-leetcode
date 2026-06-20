# 1429 — First Unique Number

## Deskripsi

**Soal:** [1429. First Unique Number](https://leetcode.com/problems/first-unique-number/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(1) amortized  
**Kompleksitas Ruang:** —

**Algoritma:** Queue (antrian FIFO)

## Solusi Go

```go
package main

// LeetCode #1429: First Unique Number
// https://leetcode.com/problems/first-unique-number/
// Difficulty: Medium

import "fmt"

type FirstUnique struct {
	queue []int
	count map[int]int
}

func main() {
	fu := NewFirstUnique([]int{2, 3, 5})
	fmt.Println(fu.ShowFirstUnique()) // 2
	fu.Add(5)
	fmt.Println(fu.ShowFirstUnique()) // 2
	fu.Add(2)
	fmt.Println(fu.ShowFirstUnique()) // 3
	fu.Add(3)
	fmt.Println(fu.ShowFirstUnique()) // -1

	fu2 := NewFirstUnique([]int{7, 7, 7, 7, 7, 7})
	fmt.Println(fu2.ShowFirstUnique()) // -1
	fu2.Add(7)
	fu2.Add(3)
	fmt.Println(fu2.ShowFirstUnique()) // 3
}

func NewFirstUnique(nums []int) FirstUnique {
	fu := FirstUnique{count: make(map[int]int)}
	for _, n := range nums {
		fu.Add(n)
	}
	return fu
}

func (this *FirstUnique) Add(value int) {
	this.count[value]++
	if this.count[value] == 1 {
		this.queue = append(this.queue, value)
	}
}

// Time: O(1) amortized
func (this *FirstUnique) ShowFirstUnique() int {
	for len(this.queue) > 0 && this.count[this.queue[0]] > 1 {
		this.queue = this.queue[1:]
	}
	if len(this.queue) == 0 {
		return -1
	}
	return this.queue[0]
}
```
