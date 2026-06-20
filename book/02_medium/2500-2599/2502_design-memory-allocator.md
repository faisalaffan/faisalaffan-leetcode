# 2502 — Design Memory Allocator

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diminta mendesain struktur data kustom dengan operasi spesifik (insert, delete, search). Target: O(1) atau O(log n) per operasi.

**Cara berpikir:** Kombinasikan HashMap + Heap + Linked List sesuai kebutuhan.

**Fungsi Solusi:** `func Constructor(n int) Allocator`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n) per allocate  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #2502: Design Memory Allocator
// https://leetcode.com/problems/design-memory-allocator/
// Difficulty: Medium
// Time: O(n) per allocate | O(1) per free
// Array-based allocation: find first free block of size.

import "fmt"

type Allocator struct {
	mem []int
}

func main() {
	alloc := Constructor(10)
	fmt.Println(alloc.Allocate(1, 1))  // 0
	fmt.Println(alloc.Allocate(1, 2))  // 1
	fmt.Println(alloc.Allocate(1, 1))  // 2
	fmt.Println(alloc.Free(1))          // 2
	fmt.Println(alloc.Allocate(2, 1))  // 0 (freed 0 and 2)
	fmt.Println(alloc.Free(2))          // 1
	fmt.Println(alloc.Allocate(3, 1))  // -1 (need 3, only 1 free block of size 1)
}

func Constructor(n int) Allocator {
	return Allocator{mem: make([]int, n)}
}

func (a *Allocator) Allocate(size int, mID int) int {
	n := len(a.mem)
	for i := 0; i < n; {
		if a.mem[i] != 0 {
			i++
			continue
		}
		j := i
		for j < n && a.mem[j] == 0 {
			j++
		}
		if j-i >= size {
			for k := i; k < i+size; k++ {
				a.mem[k] = mID
			}
			return i
		}
		i = j
	}
	return -1
}

func (a *Allocator) Free(mID int) int {
	count := 0
  // Range loop
	for i := range a.mem {
		if a.mem[i] == mID {
			a.mem[i] = 0
			count++
		}
	}
	return count
}
```
