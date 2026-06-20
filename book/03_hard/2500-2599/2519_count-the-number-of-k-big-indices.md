# 2519 — Count The Number Of K Big Indices

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func countKBigIndices(nums []int, k int) int
```

> **💡 Hint:** Coordinate compression + Fenwick Tree (BIT).

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Two Pointer, Fenwick Tree (BIT)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2519: Count the Number of K-Big Indices
// https://leetcode.com/problems/count-the-number-of-k-big-indices/
// Difficulty: Hard [Paid]

import (
	"fmt"
	"sort"
)

// countKBigIndices counts indices i where at least k elements before i
// are strictly less than nums[i] AND at least k elements after i
// are strictly less than nums[i].
//
// Approach: Coordinate compression + Fenwick Tree (BIT).
// Scan left-to-right to count less-than-curr before i,
// scan right-to-left to count less-than-curr after i.
//
// Complexity: O(n log n) time, O(n) space
func countKBigIndices(nums []int, k int) int {
	n := len(nums)

	// Coordinate compression
  // Alokasi slice integer
	sorted := make([]int, n)
	copy(sorted, nums)
  // Urutkan secara ascending — O(n log n)
	sort.Ints(sorted)
  // Membuat map (HashMap) — pencarian O(1)
	rank := make(map[int]int)
	for _, v := range sorted {
		if _, ok := rank[v]; !ok {
			rank[v] = len(rank) + 1 // 1-indexed for BIT
		}
	}
	m := len(rank)

	// leftLess[i] = count of elements before i with value < nums[i]
  // Alokasi slice integer
	bit := make([]int, m+2)
  // Alokasi slice integer
	leftLess := make([]int, n)
	for i := 0; i < n; i++ {
		r := rank[nums[i]]
		leftLess[i] = bitQuery(bit, r-1)
		bitUpdate(bit, r, 1)
	}

	// rightLess[i] = count of elements after i with value < nums[i]
	bit = make([]int, m+2)
  // Alokasi slice integer
	rightLess := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		r := rank[nums[i]]
		rightLess[i] = bitQuery(bit, r-1)
		bitUpdate(bit, r, 1)
	}

	ans := 0
	for i := 0; i < n; i++ {
		if leftLess[i] >= k && rightLess[i] >= k {
			ans++
		}
	}
	return ans
}

func bitUpdate(bit []int, idx, delta int) {
	for idx < len(bit) {
		bit[idx] += delta
		idx += idx & -idx
	}
}

func bitQuery(bit []int, idx int) int {
	sum := 0
	for idx > 0 {
		sum += bit[idx]
		idx -= idx & -idx
	}
	return sum
}

func main() {
	// Test cases
	fmt.Println("Test 1: nums=[2,3,1,2,1], k=2 ->", countKBigIndices([]int{2, 3, 1, 2, 1}, 2))
	fmt.Println("Test 2: nums=[1,2,3,4,5], k=2 ->", countKBigIndices([]int{1, 2, 3, 4, 5}, 2))
	fmt.Println("Test 3: nums=[5,4,3,2,1], k=1 ->", countKBigIndices([]int{5, 4, 3, 2, 1}, 1))
	fmt.Println("Test 4: nums=[1,1,1], k=1 ->", countKBigIndices([]int{1, 1, 1}, 1))
	fmt.Println("Test 5: nums=[1], k=1 ->", countKBigIndices([]int{1}, 1))
	fmt.Println("Test 6: nums=[3,2,1,2,3], k=2 ->", countKBigIndices([]int{3, 2, 1, 2, 3}, 2))
}
```
