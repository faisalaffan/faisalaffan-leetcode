# 3935 — Power Update After K Th Largest Insertion I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func NewBIT(size int) *BIT
```

> **💡 Hint:** Maintain sorted multiset via Fenwick tree. For each query,

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Binary Search, Prefix Sum, Fenwick Tree (BIT)

**Kompleksitas Waktu:** O(N log M + Q log M)  
**Kompleksitas Ruang:** O(M) where M = max value

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3935: Power Update After K-th Largest Insertion I
// https://leetcode.com/problems/power-update-after-k-th-largest-insertion-i/
// Difficulty: Medium [Paid]
// Time: O(N log M + Q log M) | Space: O(M) where M = max value
// Approach: Maintain sorted multiset via Fenwick tree. For each query,
// insert val, find k-th largest via binary search on BIT prefix sums,
// update p = p ^ kth % MOD.

import (
	"fmt"
	"sort"
)

const MOD = 1000000007

type BIT struct {
	tree []int
	size int
}

func NewBIT(size int) *BIT {
	return &BIT{tree: make([]int, size+2), size: size}
}

func (b *BIT) Update(idx int, delta int) {
	idx++
	for idx <= b.size+1 {
		b.tree[idx] += delta
		idx += idx & -idx
	}
}

func (b *BIT) Query(idx int) int {
	idx++
	sum := 0
	for idx > 0 {
		sum += b.tree[idx]
		idx -= idx & -idx
	}
	return sum
}

func (b *BIT) KthLargest(k int) int {
	// Find smallest idx such that suffix sum (total - prefix) >= k
	// i.e., total - Query(idx-1) >= k
	// i.e., Query(idx-1) <= total - k
	total := b.Query(b.size - 1)
	target := total - k // number of elements strictly less than the kth largest

	// Binary search for first index with prefix sum > target
	lo, hi := 0, b.size-1
	for lo < hi {
		mid := (lo + hi) / 2
		if b.Query(mid) > target {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func PowerUpdateAfterKThLargestInsertionI(nums []int, p int, queries [][]int) []int {
	// Coordinate compress all values
  // Alokasi slice integer
	allVals := make([]int, 0, len(nums)+len(queries))
	allVals = append(allVals, nums...)
	for _, q := range queries {
		allVals = append(allVals, q[0])
	}
  // Urutkan secara ascending — O(n log n)
	sort.Ints(allVals)
	uniq := []int{allVals[0]}
	for i := 1; i < len(allVals); i++ {
		if allVals[i] != allVals[i-1] {
			uniq = append(uniq, allVals[i])
		}
	}

  // Membuat map (HashMap) — pencarian O(1)
	rank := make(map[int]int)
	for i, v := range uniq {
		rank[v] = i
	}
	m := len(uniq)

	bit := NewBIT(m)
	for _, v := range nums {
		bit.Update(rank[v], 1)
	}

	cur := int64(p)
  // Alokasi slice integer
	ans := make([]int, len(queries))

	for idx, q := range queries {
		val, k := q[0], q[1]
		bit.Update(rank[val], 1)

		kthVal := uniq[bit.KthLargest(k)]

		cur = (cur ^ int64(kthVal)) % MOD
		ans[idx] = int(cur)
	}

	return ans
}

func main() {
	// Example 1
	fmt.Println(PowerUpdateAfterKThLargestInsertionI(
		[]int{2}, 4, [][]int{{3, 1}, {1, 2}},
	)) // Expected: [64, 4096]

	// Example 2
	fmt.Println(PowerUpdateAfterKThLargestInsertionI(
		[]int{7, 5}, 6, [][]int{{4, 3}, {7, 2}},
	)) // Expected: [1296, 220296870]
}
```
