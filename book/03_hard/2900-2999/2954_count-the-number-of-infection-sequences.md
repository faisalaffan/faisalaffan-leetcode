# 2954 — Count The Number Of Infection Sequences

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func powMod(a, e int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #2954: Count the Number of Infection Sequences
// https://leetcode.com/problems/count-the-number-of-infection-sequences/
// Difficulty: Hard
//
// Initially some houses are infected. Each day, the infection spreads to one
// adjacent uninfected house. Count number of possible infection sequences.
//
// For each gap of uninfected houses between infected ones:
//   - Internal gap of size g (bounded by two infected houses): 2^(g-1) ways
//   - End gap of size g (only one adjacent infected house): 1 way
//
// The total = (sum(g))! / prod(g_i!) * prod(internal_ways)
// where internal_ways = 2^(g_i-1) for internal gaps.
// Equivalent to: total_uninfected! * prod(2^(g_i-1) * inv_fact[g_i]) % MOD
// for internal gaps, and * inv_fact[g_i] % MOD for end gaps.

import (
	"fmt"
)

const mod = 1_000_000_007

func powMod(a, e int) int {
	res := 1
	for e > 0 {
		if e&1 == 1 {
			res = (res * a) % mod
		}
		a = (a * a) % mod
		e >>= 1
	}
	return res
}

func numberOfInfectionSequences(n int, infected []int) int {
	m := len(infected)

	// Build gap list
	// Gap before first infected house (end gap)
  // Alokasi slice
	gaps := make([]int, 0)
	totalInfected := 0

	// Gap before first infected
	firstGap := infected[0]
	if firstGap > 0 {
		gaps = append(gaps, firstGap)
		totalInfected += firstGap
	}

	// Internal gaps
	for i := 1; i < m; i++ {
		gapSize := infected[i] - infected[i-1] - 1
		if gapSize > 0 {
			gaps = append(gaps, gapSize)
			totalInfected += gapSize
		}
	}

	// Gap after last infected (end gap)
	lastGap := n - 1 - infected[m-1]
	if lastGap > 0 {
		gaps = append(gaps, lastGap)
		totalInfected += lastGap
	}

	if totalInfected == 0 {
		return 1
	}

	// Precompute factorials up to totalInfected
  // Alokasi slice
	fact := make([]int, totalInfected+1)
  // Alokasi slice
	invFact := make([]int, totalInfected+1)
	fact[0] = 1
	for i := 1; i <= totalInfected; i++ {
		fact[i] = (fact[i-1] * i) % mod
	}
	invFact[totalInfected] = powMod(fact[totalInfected], mod-2)
	for i := totalInfected - 1; i >= 0; i-- {
		invFact[i] = (invFact[i+1] * (i + 1)) % mod
	}

	// Total = totalInfected! / prod(gap_i!) * prod(factor_i)
	// where factor_i = 2^(g_i-1) for internal gaps, 1 for end gaps
	ans := fact[totalInfected]

	// First and last gaps in the list are end gaps (if they exist)
	// Internal gaps: those not at position 0 (if firstGap > 0) or last (if lastGap > 0)
	// Actually, the gaps list is: [firstGap?, internal gaps..., lastGap?]
	// We know which are internal: those between firstGap and lastGap.
	// Since firstGap is at index 0 (if present), and lastGap is at last index (if present).
	// Internal gaps: index from `hasFirst` to `len(gaps)-1-hasLast`.

	hasFirst := 0
	if firstGap > 0 {
		hasFirst = 1
	}
	hasLast := 0
	if lastGap > 0 {
		hasLast = 1
	}

	for i, g := range gaps {
		ans = (ans * invFact[g]) % mod
		// Internal gap gets factor 2^(g-1)
		if i >= hasFirst && i < len(gaps)-hasLast {
			ans = (ans * powMod(2, g-1)) % mod
		}
	}

	return ans
}

func main() {
	// Example: n=5, infected=[0,4]
	// Houses: [0=infected, 1, 2, 3, 4=infected]
	// Internal gap of size 3 between 0 and 4
	// Ways: fact[3] * inv_fact[3] * 2^(3-1) = 6 * inv(6) * 4 = 4
	fmt.Println(numberOfInfectionSequences(5, []int{0, 4}))

	// n=4, infected=[1]
	// End gaps: [0] of size 1, [2,3] of size 2
	// Total: fact[3] * inv_fact[1] * inv_fact[2] = 6 * 1 * inv(2) = 3
	// (end gaps, both factor 1)
	fmt.Println(numberOfInfectionSequences(4, []int{1}))

	// All houses initially infected: only 1 sequence
	fmt.Println(numberOfInfectionSequences(3, []int{0, 1, 2}))

	// No houses infected initially (should handle)
	// This case: n houses, no initially infected.
	// Then the first house to get infected can be any of the n houses.
	// After that, the infection spreads from that house outward.
	// Actually the problem seems to assume at least one initially infected.
	// Let's just test 2 houses with 1 infected.
	fmt.Println(numberOfInfectionSequences(2, []int{0}))
}
```
