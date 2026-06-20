# 1570 — Dot Product Of Two Sparse Vectors

## Deskripsi

**Soal:** [1570. Dot Product Of Two Sparse Vectors](https://leetcode.com/problems/dot-product-of-two-sparse-vectors/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(N+M), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** Two Pointer (penunjuk kiri & kanan)

## Solusi Go

```go
package main

// LeetCode #1570: Dot Product of Two Sparse Vectors
// https://leetcode.com/problems/dot-product-of-two-sparse-vectors/
// Difficulty: Medium [Paid]

import "fmt"

func main() {
	v1 := ConstructorSparse([]int{1, 0, 0, 2, 3})
	v2 := ConstructorSparse([]int{0, 3, 0, 4, 0})

	dotProduct := v1.dotProduct(v2)
	fmt.Println("Dot product:", dotProduct)

	// Additional test
	v3 := ConstructorSparse([]int{0, 1, 0, 0, 0})
	v4 := ConstructorSparse([]int{0, 0, 0, 0, 2})
	fmt.Println("Dot product (no overlap):", v3.dotProduct(v4))
}

// SparseVector stores non-zero elements with their indices.
type SparseVector struct {
	pairs [][2]int // [index, value]
}

func ConstructorSparse(nums []int) SparseVector {
  // Membuat slice untuk menyimpan hasil
	pairs := make([][2]int, 0)
	for i, num := range nums {
		if num != 0 {
			pairs = append(pairs, [2]int{i, num})
		}
	}
	return SparseVector{pairs: pairs}
}

func (sv *SparseVector) dotProduct(vec SparseVector) int {
	// Time: O(N+M), Space: O(1)
	// Two pointer approach on sorted pairs
	result := 0
	i, j := 0, 0

	for i < len(sv.pairs) && j < len(vec.pairs) {
		if sv.pairs[i][0] == vec.pairs[j][0] {
			result += sv.pairs[i][1] * vec.pairs[j][1]
			i++
			j++
		} else if sv.pairs[i][0] < vec.pairs[j][0] {
			i++
		} else {
			j++
		}
	}

	return result
}
```
