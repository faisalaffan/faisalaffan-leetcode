# 3444 — Minimum Increments For Target Multiples In An Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func lcm(a, b int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming, GCD / Matematika, Bitmask

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3444: Minimum Increments for Target Multiples in an Array
// https://leetcode.com/problems/minimum-increments-for-target-multiples-in-an-array/
// Difficulty: Hard
//
// DP over bitmask of targets. For each array element, compute cost to make it
// divisible by each subset S of targets (cost = nearest multiple of lcm(S)).
// Then dp[mask] = min cost to cover mask using processed elements (0/1 knapSack).

import "fmt"

const INF = 1 << 60

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

func minIncrementsForTargetMultiples(nums []int, target []int) int {
	n := len(nums)
	m := len(target)
	M := 1 << m

	// Precompute LCM for each subset mask
  // Alokasi slice integer
	lcmMask := make([]int, M)
	lcmMask[0] = 1
	for mask := 1; mask < M; mask++ {
		lsb := mask & -mask
		bit := 0
		for lsb>>bit != 1 {
			bit++
		}
		prev := mask ^ lsb
		if prev == 0 {
			lcmMask[mask] = target[bit]
		} else {
			l := lcm(lcmMask[prev], target[bit])
			if l > 1_000_000_000 {
				l = 1_000_000_001
			}
			lcmMask[mask] = l
		}
	}

	// For each element, min increment to cover each mask
	// cost[i][mask] = min increment to make nums[i] divisible by lcmMask[mask]
  // Membuat matriks/slice 2D untuk DP
	elemCost := make([][]int, n)
	for i, x := range nums {
		elemCost[i] = make([]int, M)
		elemCost[i][0] = 0
		for mask := 1; mask < M; mask++ {
			l := lcmMask[mask]
			if l > 1_000_000_000 {
				elemCost[i][mask] = INF
				continue
			}
			rem := x % l
			if rem == 0 {
				elemCost[i][mask] = 0
			} else {
				elemCost[i][mask] = l - rem
			}
		}
	}

	// 0/1 knapSack DP over elements
  // Alokasi slice integer
	dp := make([]int, M)
	for mask := 1; mask < M; mask++ {
		dp[mask] = INF
	}

	for _, cost := range elemCost {
  // Alokasi slice integer
		ndp := make([]int, M)
		copy(ndp, dp)
		for oldMask := 0; oldMask < M; oldMask++ {
			if dp[oldMask] == INF {
				continue
			}
			for s := 1; s < M; s++ {
				if cost[s] == INF {
					continue
				}
				newMask := oldMask | s
				cand := dp[oldMask] + cost[s]
				if cand < ndp[newMask] {
					ndp[newMask] = cand
				}
			}
		}
		dp = ndp
	}

	return dp[M-1]
}

func main() {
	fmt.Printf("[1,2,3] target=[4,2,6] -> %d\n",
		minIncrementsForTargetMultiples([]int{1, 2, 3}, []int{4, 2, 6}))

	fmt.Printf("[2,3,5] target=[3,5] -> %d\n",
		minIncrementsForTargetMultiples([]int{2, 3, 5}, []int{3, 5}))

	fmt.Printf("[1] target=[2] -> %d\n",
		minIncrementsForTargetMultiples([]int{1}, []int{2}))

	fmt.Printf("[4,8,12] target=[3] -> %d\n",
		minIncrementsForTargetMultiples([]int{4, 8, 12}, []int{3}))

	fmt.Printf("[2,5] target=[4,6,8] -> %d\n",
		minIncrementsForTargetMultiples([]int{2, 5}, []int{4, 6, 8}))
}
```
