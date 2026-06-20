# 1856 — Maximum Subarray Min Product

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MaxSumMinProduct(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Stack, Prefix Sum, Monotonic Stack/Queue

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1856: Maximum Subarray Min-Product
// https://leetcode.com/problems/maximum-subarray-min-product/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MaxSumMinProduct([]int{1, 2, 3, 2}))
	fmt.Println(MaxSumMinProduct([]int{2, 3, 3, 1, 2}))
	fmt.Println(MaxSumMinProduct([]int{3, 1, 5, 6, 4, 2}))
}

const mod = 1000000007

// Time: O(n), Space: O(n)
func MaxSumMinProduct(nums []int) int {
	n := len(nums)
  // Alokasi slice integer
	prefix := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + nums[i]
	}

	// Monotonic stack to find previous smaller and next smaller elements
  // Alokasi slice integer
	left := make([]int, n)  // left[i] = index of previous smaller element
  // Alokasi slice integer
	right := make([]int, n) // right[i] = index of next smaller element

  // Alokasi slice integer
	stack := make([]int, 0)
	for i := 0; i < n; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] >= nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	stack = make([]int, 0)
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && nums[stack[len(stack)-1]] >= nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			right[i] = n
		} else {
			right[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	maxProd := int64(0)
	for i := 0; i < n; i++ {
		sum := int64(prefix[right[i]] - prefix[left[i]+1])
		prod := sum * int64(nums[i])
		if prod > maxProd {
			maxProd = prod
		}
	}
	return int(maxProd % mod)
}
```
