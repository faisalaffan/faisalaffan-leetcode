# 2502 — Design Memory Allocator

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diminta untuk mendesain (merancang) sebuah struktur data kustom dengan operasi tertentu (insert, delete, search, update). Tugasmu adalah memilih representasi data yang tepat agar setiap operasi berjalan efisien — biasanya O(1) atau O(log n).

Ini adalah soal yang paling sering muncul di interview sistem desain. Kamu perlu memilih kombinasi struktur data yang tepat (HashMap + Heap + LinkedList) untuk mencapai kompleksitas yang diminta.

**Konsep kunci:** HashMap (O(1) lookup), Heap (priority), Doubly Linked List (O(1) remove), TreeMap (ordered keys).

**Fungsi yang perlu kamu implementasikan:**
```go
func Constructor(n int) Allocator
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n) per allocate  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

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
  // Range loop: iterasi dengan indeks + nilai
	for i := range a.mem {
		if a.mem[i] == mID {
			a.mem[i] = 0
			count++
		}
	}
	return count
}
```
