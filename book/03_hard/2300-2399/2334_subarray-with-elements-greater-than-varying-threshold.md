# 2334 — Subarray With Elements Greater Than Varying Threshold

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func validSubarraySize(nums []int, threshold int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Stack, Monotonic Stack/Queue

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
)

// 2334. Subarray With Elements Greater Than Varying Threshold
// ----------------------------------------------------------------
// Find the longest subarray such that every element is strictly greater than
// threshold / len(subarray).  Equivalently, if the minimum element of the
// subarray is m and length is L, we need m * L > threshold.
//
// For each element nums[i], find the maximal subarray where nums[i] is the
// *minimum* (strictly smaller than any element outside).  The left bound is
// given by the previous element strictly smaller than nums[i];
// the right bound by the next element strictly smaller.
//
//   left  = index of previous smaller element + 1
//   right = index of next smaller element - 1
//   maxLen = right - left + 1
//
// Then check whether nums[i] * maxLen > threshold.  If yes, maxLen is a
// candidate answer (we can always shorten the subarray since every element
// is still ≥ nums[i] which is > threshold/maxLen ≥ threshold/shorterLen).
// Wait — if we shorten the subarray, the condition becomes easier because
// threshold/shorterLen is larger, but the minimum may increase if we cut off
// elements.  Since nums[i] is the *global* minimum for the maximal subarray,
// any subarray still containing nums[i] has minimum ≤ nums[i].  The condition
// nums[i] * L > threshold is monotonic in L for fixed i: larger L is better.
// So checking only maxLen is sufficient.

func validSubarraySize(nums []int, threshold int) int {
	n := len(nums)
  // Alokasi slice integer
	prevSmaller := make([]int, n)
  // Alokasi slice integer
	nextSmaller := make([]int, n)

	// Previous strictly smaller element.
  // Alokasi slice integer
	stack := make([]int, 0)
	for i := 0; i < n; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] >= nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			prevSmaller[i] = -1
		} else {
			prevSmaller[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	// Next strictly smaller element.
	stack = stack[:0]
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && nums[stack[len(stack)-1]] >= nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			nextSmaller[i] = n
		} else {
			nextSmaller[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	best := -1
	for i := 0; i < n; i++ {
		left := prevSmaller[i] + 1
		right := nextSmaller[i] - 1
		maxLen := right - left + 1
		// Use int64 to avoid overflow.
		if int64(nums[i])*int64(maxLen) > int64(threshold) && maxLen > best {
			best = maxLen
		}
	}
	return best
}

// ---------------------------------------------------------------------------
//  Wrapper

func SubarrayWithElementsGreaterThanVaryingThreshold() interface{} {
	return validSubarraySize([]int{1, 3, 4, 3, 1}, 6)
}

func main() {
	fmt.Println(SubarrayWithElementsGreaterThanVaryingThreshold())

	tests := []struct {
		nums      []int
		threshold int
		want      int
	}{
		{[]int{1, 3, 4, 3, 1}, 6, 3},
		{[]int{6, 5, 6, 5, 8}, 7, 5},
		{[]int{1, 2, 3, 4, 5}, 100, -1},
		{[]int{5, 5, 5, 5}, 10, 4},
		{[]int{2, 1, 2}, 1, 3},
	}
	for _, tc := range tests {
		got := validSubarraySize(tc.nums, tc.threshold)
		if got != tc.want {
			fmt.Printf("FAIL nums=%v threshold=%d: got %d, want %d\n",
				tc.nums, tc.threshold, got, tc.want)
		}
	}
	fmt.Println("Done testing 2334.")
}
```
