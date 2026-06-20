# 2344 — Minimum Deletions To Make Array Divisible

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minDeletions(nums []int, numsDivide []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** GCD / Matematika

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **GCD / Matematika** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
	"sort"
)

// 2344. Minimum Deletions to Make Array Divisible
// ----------------------------------------------------------------
// Given nums (integers) and numsDivide (array of divisors), we can delete
// elements from nums.  Return the minimum number of deletions so that the
// smallest remaining element divides every element of numsDivide.
//
// Strategy:
//   1. Compute g = gcd of all elements in numsDivide.
//   2. Sort nums.
//   3. Scan from the smallest: find the first element that divides g.
//      Return its index (number of deletions).  If none found, return -1.

func minDeletions(nums []int, numsDivide []int) int {
	// GCD of all numsDivide.
	g := 0
	for _, v := range numsDivide {
		g = gcd(g, v)
	}

  // Urutkan secara ascending — O(n log n)
	sort.Ints(nums)
	for i, v := range nums {
		if g%v == 0 {
			return i
		}
	}
	return -1
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// ---------------------------------------------------------------------------
//  Wrapper

func MinimumDeletionsToMakeArrayDivisible() interface{} {
	return minDeletions([]int{4, 3, 6}, []int{8, 2, 10})
}

func main() {
	fmt.Println(MinimumDeletionsToMakeArrayDivisible())

	tests := []struct {
		nums      []int
		numsDiv   []int
		want      int
	}{
		{[]int{4, 3, 6}, []int{8, 2, 10}, -1},
		{[]int{3, 2, 6}, []int{9, 6, 12}, 1},
		{[]int{5, 7, 11}, []int{2, 3, 4}, -1},
		{[]int{2, 4, 8}, []int{4, 8, 16}, 0},
		{[]int{8, 12, 6, 4}, []int{9, 15}, -1},
	}
	for _, tc := range tests {
		got := minDeletions(tc.nums, tc.numsDiv)
		if got != tc.want {
			fmt.Printf("FAIL nums=%v numsDiv=%v: got %d, want %d\n",
				tc.nums, tc.numsDiv, got, tc.want)
		}
	}
	fmt.Println("Done testing 2344.")
}
```
