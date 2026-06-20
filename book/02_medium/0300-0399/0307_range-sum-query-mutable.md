# 0307 — Range Sum Query Mutable

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func Constructor(nums []int) NumArray
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log n) per operation, Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #307: Range Sum Query - Mutable
// https://leetcode.com/problems/range-sum-query-mutable/
// Difficulty: Medium
// Time: O(log n) per operation, Space: O(n)

import "fmt"

type NumArray struct {
	nums []int
	tree []int
	n    int
}

func Constructor(nums []int) NumArray {
	n := len(nums)
  // Alokasi slice integer
	tree := make([]int, n+1)
	na := NumArray{nums, tree, n}
	for i, val := range nums {
		na.add(i+1, val)
	}
	return na
}

func (this *NumArray) add(idx, val int) {
	for idx <= this.n {
		this.tree[idx] += val
		idx += idx & -idx
	}
}

func (this *NumArray) sum(idx int) int {
	res := 0
	for idx > 0 {
		res += this.tree[idx]
		idx -= idx & -idx
	}
	return res
}

func (this *NumArray) Update(index int, val int) {
	diff := val - this.nums[index]
	this.nums[index] = val
	this.add(index+1, diff)
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.sum(right+1) - this.sum(left)
}

func main() {
	na := Constructor([]int{1, 3, 5})
	fmt.Println(na.SumRange(0, 2))
	na.Update(1, 2)
	fmt.Println(na.SumRange(0, 2))

	na2 := Constructor([]int{1})
	fmt.Println(na2.SumRange(0, 0))
	na2.Update(0, 5)
	fmt.Println(na2.SumRange(0, 0))
}
```
