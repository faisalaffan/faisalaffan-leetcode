# 3416 — Subsequences With A Unique Middle Mode Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func SubsequencesWithAUniqueMiddleModeIi(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Two Pointer, Binary Search

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3416: Subsequences with a Unique Middle Mode II
// https://leetcode.com/problems/subsequences-with-a-unique-middle-mode-ii/
// Difficulty: Hard [Paid]
//
// Combinatorics with inclusion-exclusion. Fix middle index,
// count C(left,2)*C(right,2), subtract invalid cases.
// Uses MOD = 1e9+7.

import "fmt"

func main() {
	fmt.Println(SubsequencesWithAUniqueMiddleModeIi([]int{1, 2, 2, 3, 3, 4}))
}

const MOD3416 = 1000000007

func SubsequencesWithAUniqueMiddleModeIi(nums []int) int {
	n := len(nums)
	if n < 5 {
		return 0
	}

	ans := int64(0)
	for mid := 2; mid <= n-3; mid++ {
  // Membuat map (HashMap) — pencarian O(1)
		left := make(map[int]int)
  // Membuat map (HashMap) — pencarian O(1)
		right := make(map[int]int)
		for i := 0; i < mid; i++ {
			left[nums[i]]++
		}
		for i := mid + 1; i < n; i++ {
			right[nums[i]]++
		}

		val := nums[mid]
		lv := left[val]
		rv := right[val]

		// Total pairs from left and right
		leftLen := mid
		rightLen := n - mid - 1

		totalLeft := int64(leftLen * (leftLen - 1) / 2)
		totalRight := int64(rightLen * (rightLen - 1) / 2)
		total := totalLeft * totalRight % MOD3416

		// Subtract: both left and right pairs use val
		if lv >= 2 && rv >= 2 {
			lp := int64(lv * (lv - 1) / 2)
			rp := int64(rv * (rv - 1) / 2)
			total = (total - lp*rp%MOD3416 + MOD3416) % MOD3416
		}

		// For each other value, subtract invalid contributions
		for x, lc := range left {
			if x == val {
				continue
			}
			rc := right[x]
			invalid := int64(0)

			// x appears 2+ in left
			if lc >= 2 {
				invalid = (invalid + int64(lc*(lc-1)/2)*totalRight) % MOD3416
			}
			// x appears 2+ in right
			if rc >= 2 {
				invalid = (invalid + totalLeft*int64(rc*(rc-1)/2)) % MOD3416
			}
			// x appears 1 on each side
			if lc >= 1 && rc >= 1 {
				leftRest := leftLen - 1
				rightRest := rightLen - 1
				if leftRest >= 1 && rightRest >= 1 {
					invalid = (invalid + int64(lc)*int64(rc)%MOD3416*int64(leftRest)%MOD3416*int64(rightRest)) % MOD3416
				}
			}

			total = (total - invalid + MOD3416) % MOD3416
		}

		ans = (ans + total) % MOD3416
	}
	return int(ans)
}
```
