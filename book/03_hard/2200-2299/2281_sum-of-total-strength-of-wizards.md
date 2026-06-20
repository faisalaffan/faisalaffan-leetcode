# 2281 — Sum Of Total Strength Of Wizards

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func totalStrength(strength []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Stack, Prefix Sum, Monotonic Stack/Queue

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2281: Sum of Total Strength of Wizards
// https://leetcode.com/problems/sum-of-total-strength-of-wizards/
// Difficulty: Hard
//
// Monotonic stack + prefix sums of prefix sums:
// For each element as minimum, find its range of dominance via next smaller
// element on left/right. Use prefix-of-prefix sums for O(1) range sum queries.

import (
	"fmt"
)

const mod = 1000000007

func main() {
	// [1,3,1,2] => 44
	fmt.Println(totalStrength([]int{1, 3, 1, 2}))
	// [5] => 5
	fmt.Println(totalStrength([]int{5}))
	// [1,2,3] => 33
	fmt.Println(totalStrength([]int{1, 2, 3}))
	// [2,2,2]
	fmt.Println(totalStrength([]int{2, 2, 2}))
}

func totalStrength(strength []int) int {
	n := len(strength)

	// Previous smaller element (strictly smaller, index).
  // Alokasi slice integer
	left := make([]int, n)
  // Alokasi slice integer
	stack := make([]int, 0, n)
	for i := 0; i < n; i++ {
		for len(stack) > 0 && strength[stack[len(stack)-1]] >= strength[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	// Next smaller element (strictly smaller, index).
  // Alokasi slice integer
	right := make([]int, n)
	stack = stack[:0]
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && strength[stack[len(stack)-1]] > strength[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			right[i] = n
		} else {
			right[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	// Prefix sums.
  // Alokasi slice integer
	pref := make([]int, n+1)
	for i := 0; i < n; i++ {
		pref[i+1] = (pref[i] + strength[i]) % mod
	}

	// Prefix sums of prefix sums.
  // Alokasi slice integer
	pref2 := make([]int, n+2)
	for i := 0; i <= n; i++ {
		pref2[i+1] = (pref2[i] + pref[i]) % mod
	}

	rangePref := func(l, r int) int {
		if l > r {
			return 0
		}
		return (pref2[r+1] - pref2[l] + mod) % mod
	}

	ans := 0
	for i := 0; i < n; i++ {
		l := left[i]
		r := right[i]

		// Number of subarrays where arr[i] is the minimum:
		// Left choices = i - l, Right choices = r - i
		leftLen := i - l
		rightLen := r - i

		// Sum of subarray sums where arr[i] is minimum:
		// = leftLen * sum(pref[i+1..r]) - rightLen * sum(pref[l+1..i])
		sumRight := rangePref(i+1, r)
		sumLeft := rangePref(l+1, i)

		total := (leftLen*sumRight%mod - rightLen*sumLeft%mod + mod) % mod
		contrib := (strength[i] % mod) * total % mod
		ans = (ans + contrib) % mod
	}

	return ans
}
```
