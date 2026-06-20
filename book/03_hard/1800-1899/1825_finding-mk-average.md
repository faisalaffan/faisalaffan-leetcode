# 1825 — Finding Mk Average

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func NewFenwick(n int) *Fenwick
```

> **💡 Hint:** Fenwick Tree (Binary Indexed Tree) + Circular Buffer.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Sliding Window, BFS, Prefix Sum, Bitmask, Fenwick Tree (BIT)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Sliding Window** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1825: Finding MK Average
// https://leetcode.com/problems/finding-mk-average/
// Difficulty: Hard
//
// Approach: Fenwick Tree (Binary Indexed Tree) + Circular Buffer.
//   Maintain a sliding window of the last m elements.
//   Use two BITs: one for element counts, one for element sums.
//   - addElement: add to BITs + queue; if queue full, evict oldest.
//   - calculateMKAverage: use order statistics to find the kth and
//     (m-k)th smallest values, then compute:
//       midSum = totalSum - sum(k smallest) - sum(k largest)
//       return midSum / (m - 2*k)

import "fmt"

func main() {
	// Example: MKAverage(3, 1)
	mk := Constructor(3, 1)
	mk.AddElement(3)
	mk.AddElement(1)
	fmt.Println("After [3,1]:", mk.CalculateMKAverage()) // not enough elements -> -1
	mk.AddElement(4)
	fmt.Println("After [3,1,4]:", mk.CalculateMKAverage()) // window=[3,1,4], remove k=1 smallest(1) and 1 largest(4), mid=[3], avg=3
	mk.AddElement(2)
	fmt.Println("After [3,1,4,2]:", mk.CalculateMKAverage()) // window=[1,4,2], remove 1 and 4, mid=[2], avg=2
	mk.AddElement(5)
	fmt.Println("After [3,1,4,2,5]:", mk.CalculateMKAverage()) // window=[4,2,5], remove 2 and 5, mid=[4], avg=4

	// Edge: larger test
	fmt.Println("---")
	mk2 := Constructor(5, 2)
	for _, v := range []int{10, 20, 30, 40, 50, 5, 15, 25, 35, 45} {
		mk2.AddElement(v)
	}
	fmt.Println("MKAverage after 10 elements:", mk2.CalculateMKAverage())
}

// --- Fenwick Tree (Binary Indexed Tree) ---

type Fenwick struct {
	tree []int
	n    int
}

func NewFenwick(n int) *Fenwick {
	return &Fenwick{
		tree: make([]int, n+1), // 1-indexed internally
		n:    n,
	}
}

func (f *Fenwick) add(idx, delta int) {
	i := idx + 1 // convert to 1-indexed
	for i <= f.n {
		f.tree[i] += delta
		i += i & -i
	}
}

func (f *Fenwick) sum(idx int) int {
	if idx < 0 {
		return 0
	}
	if idx >= f.n {
		idx = f.n - 1
	}
	i := idx + 1 // convert to 1-indexed
	res := 0
	for i > 0 {
		res += f.tree[i]
		i -= i & -i
	}
	return res
}

// kth returns the smallest 0-indexed value such that prefix sum >= k.
// k is 1-indexed (1-based rank).
func (f *Fenwick) kth(k int) int {
	idx := 0
	bitMask := 1
	for bitMask <= f.n {
		bitMask <<= 1
	}
	bitMask >>= 1
	for bitMask > 0 {
		next := idx + bitMask
		if next <= f.n && f.tree[next] < k {
			k -= f.tree[next]
			idx = next
		}
		bitMask >>= 1
	}
	// idx is the largest internal index with prefix < k.
	// The answer in 0-indexed is idx.
	return idx
}

// --- MKAverage ---

const maxVal = 100001

type MKAverage struct {
	m          int
	k          int
	cnt        *Fenwick
	sum        *Fenwick
	queue      []int
	writeIdx   int
	elemCount  int
}

func Constructor(m int, k int) MKAverage {
	return MKAverage{
		m:     m,
		k:     k,
		cnt:   NewFenwick(maxVal),
		sum:   NewFenwick(maxVal),
		queue: make([]int, m),
	}
}

func (mk *MKAverage) AddElement(num int) {
	// Remove oldest if window is full
	if mk.elemCount >= mk.m {
		old := mk.queue[mk.writeIdx]
		mk.cnt.add(old, -1)
		mk.sum.add(old, -old)
	}

	// Add new element
	mk.queue[mk.writeIdx] = num
	mk.writeIdx = (mk.writeIdx + 1) % mk.m
	mk.elemCount++
	mk.cnt.add(num, 1)
	mk.sum.add(num, num)
}

func (mk *MKAverage) CalculateMKAverage() int {
	if mk.elemCount < mk.m {
		return -1
	}

	totalSum := mk.sum.sum(maxVal - 1)

	// Sum of k smallest elements
	vk := mk.cnt.kth(mk.k)
	sumBeforeK := mk.sum.sum(vk - 1)
	cntBeforeK := mk.cnt.sum(vk - 1)
	sumSmallestK := sumBeforeK + (mk.k-cntBeforeK)*vk

	// Sum of m-k smallest elements (to derive sum of k largest)
	vm := mk.cnt.kth(mk.m - mk.k)
	sumBeforeM := mk.sum.sum(vm - 1)
	cntBeforeM := mk.cnt.sum(vm - 1)
	sumFirstMK := sumBeforeM + (mk.m-mk.k-cntBeforeM)*vm
	sumLargestK := totalSum - sumFirstMK

	midSum := totalSum - sumSmallestK - sumLargestK
	return midSum / (mk.m - 2*mk.k)
}

// Stub kept for compatibility with the repo scaffold.
func FindingMkAverage() any {
	mk := Constructor(3, 1)
	mk.AddElement(3)
	mk.AddElement(1)
	mk.AddElement(4)
	return mk.CalculateMKAverage()
}
```
