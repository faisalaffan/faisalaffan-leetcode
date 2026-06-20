# 2361 — Minimum Costs Using The Train Line

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** —

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func minimumCosts(regular, express []int, expressCost int) []int64`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

import (
	"fmt"
	"math"
)

// 2361. Minimum Costs Using the Train Line
// ----------------------------------------------------------------
// You start at station 0 on the regular line.
// From station i to i+1:
//   regular[i] = cost to travel on regular line
//   express[i] = cost to travel on express line
// expressCost = one‑time fee to switch FROM regular TO express.
// Switching back (express → regular) is free.
//
// dpReg[i] = min cost to reach station i on regular line
// dpExp[i] = min cost to reach station i on express line
//
// Transition (i → i+1):
//   stay regular:      dpReg[i] + regular[i]
//   switch→regular:    dpExp[i] + 0 + regular[i]   (free switch)
//   stay express:      dpExp[i] + express[i]
//   switch→express:    dpReg[i] + expressCost + express[i]
//
// Result: min(dpReg[n], dpExp[n]).

func minimumCosts(regular, express []int, expressCost int) []int64 {
	n := len(regular)
	dpReg := int64(0)
	dpExp := int64(math.MaxInt64)

  // Alokasi slice
	ans := make([]int64, n)
	for i := 0; i < n; i++ {
		newReg := int64(math.MaxInt64)
		newExp := int64(math.MaxInt64)

		// Stay on regular.
		if dpReg != math.MaxInt64 {
			newReg = min64(newReg, dpReg+int64(regular[i]))
		}
		// Switch from express to regular (FREE).
		if dpExp != math.MaxInt64 {
			newReg = min64(newReg, dpExp+int64(regular[i]))
		}

		// Stay on express.
		if dpExp != math.MaxInt64 {
			newExp = min64(newExp, dpExp+int64(express[i]))
		}
		// Switch from regular to express (pay expressCost).
		if dpReg != math.MaxInt64 {
			newExp = min64(newExp, dpReg+int64(expressCost)+int64(express[i]))
		}

		dpReg = newReg
		dpExp = newExp

		ans[i] = min64(dpReg, dpExp)
	}
	return ans
}

func min64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

// ---------------------------------------------------------------------------
//  Wrapper

func MinimumCostsUsingTheTrainLine() interface{} {
	return minimumCosts([]int{1, 6, 9, 5}, []int{5, 2, 3, 10}, 8)
}

func main() {
	fmt.Println(MinimumCostsUsingTheTrainLine())

	got := minimumCosts([]int{1, 6, 9, 5}, []int{5, 2, 3, 10}, 8)
	want := []int64{1, 7, 14, 19}
  // Range loop
	for i := range got {
		if got[i] != want[i] {
			fmt.Printf("FAIL [%d]: got %d, want %d\n", i, got[i], want[i])
		}
	}

	got2 := minimumCosts([]int{11, 5, 13}, []int{7, 10, 6}, 3)
	want2 := []int64{10, 15, 24}
  // Range loop
	for i := range got2 {
		if got2[i] != want2[i] {
			fmt.Printf("FAIL2 [%d]: got %d, want %d\n", i, got2[i], want2[i])
		}
	}
	fmt.Println("Done testing 2361.")
}
```
