# 2179 — Count Good Triplets In An Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func countGoodTriplets(nums1 []int, nums2 []int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Fenwick Tree (BIT)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2179: Count Good Triplets in an Array
// https://leetcode.com/problems/count-good-triplets-in-an-array/
// Difficulty: Hard
//
// Map each value to its index in nums1. Then transform nums2 into positions.
// Count increasing triplets in the transformed array using BIT (Fenwick tree).

import "fmt"

func main() {
	fmt.Println(countGoodTriplets([]int{2, 0, 1, 3}, []int{0, 1, 2, 3}))       // 1
	fmt.Println(countGoodTriplets([]int{4, 0, 1, 3, 2}, []int{4, 1, 0, 2, 3})) // 4
	fmt.Println(countGoodTriplets([]int{0, 1, 2, 3}, []int{0, 1, 2, 3}))       // 4
	fmt.Println(countGoodTriplets([]int{0, 1, 2}, []int{2, 1, 0}))             // 0
}

func countGoodTriplets(nums1 []int, nums2 []int) int64 {
	n := len(nums1)
  // Alokasi slice integer
	pos1 := make([]int, n)
	for i, v := range nums1 {
		pos1[v] = i
	}

  // Alokasi slice integer
	arr := make([]int, n)
	for i, v := range nums2 {
		arr[i] = pos1[v]
	}

	// leftLess[i] = count of j < i with arr[j] < arr[i]
  // Alokasi slice integer
	leftLess := make([]int, n)
	bit := newFenwick(n)
	for i, v := range arr {
		leftLess[i] = bit.query(v - 1)
		bit.add(v, 1)
	}

	// rightGreater[i] = count of j > i with arr[j] > arr[i]
  // Alokasi slice integer
	rightGreater := make([]int, n)
	bit = newFenwick(n)
	for i := n - 1; i >= 0; i-- {
		rightGreater[i] = bit.query(n-1) - bit.query(arr[i])
		bit.add(arr[i], 1)
	}

	var ans int64
	for i := 0; i < n; i++ {
		ans += int64(leftLess[i]) * int64(rightGreater[i])
	}
	return ans
}

type fenwick struct {
	tree []int
	n    int
}

func newFenwick(n int) *fenwick {
	return &fenwick{tree: make([]int, n+1), n: n}
}

func (f *fenwick) add(idx, val int) {
	for i := idx + 1; i <= f.n; i += i & -i {
		f.tree[i] += val
	}
}

func (f *fenwick) query(idx int) int {
	if idx < 0 {
		return 0
	}
	res := 0
	for i := idx + 1; i > 0; i -= i & -i {
		res += f.tree[i]
	}
	return res
}

func CountGoodTripletsInAnArray() any {
	return countGoodTriplets([]int{2, 0, 1, 3}, []int{0, 1, 2, 3})
}
```
