# 3594 — Minimum Time To Transport All Individuals

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minTime(n int, k int, m int, time []int, mul []float64) float64
```

> **💡 Hint:** Bitmask DP. State = (mask of people on left, current stage).

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS, Dynamic Programming, Bitmask

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3594: Minimum Time to Transport All Individuals
// https://leetcode.com/problems/minimum-time-to-transport-all-individuals/
// Difficulty: Hard
//
// Transport n people across a river with a boat of capacity k. The crossing
// time depends on the slowest person in the boat and the current stage multiplier.
// Find minimum total time to transport all people.
//
// Approach: Bitmask DP. State = (mask of people on left, current stage).
// For each state, try all valid boatloads and returners.

import "fmt"
import "math"

func main() {
	// Example 1
	fmt.Println(minTime(2, 1, 2, []int{1, 2}, []float64{1.0, 0.5}))
	// Example 2
	fmt.Println(minTime(3, 2, 1, []int{2, 3, 4}, []float64{1.5}))
	// Edge: single person
	fmt.Println(minTime(1, 1, 1, []int{5}, []float64{1.0}))
	// Edge: no one
	fmt.Println(minTime(0, 1, 1, []int{}, []float64{1.0}))
}

func minTime(n int, k int, m int, time []int, mul []float64) float64 {
  // Edge case: input kosong — langsung return
	if n == 0 {
		return 0
	}

	// memo[mask][stage]
  // Membuat matriks/slice 2D untuk DP
	memo := make([][]float64, 1<<n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range memo {
		memo[i] = make([]float64, m)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(mask int, stage int) float64
	dfs = func(mask int, stage int) float64 {
		if mask == 0 {
			return 0
		}
		if memo[mask][stage] != -1 {
			return memo[mask][stage]
		}

		best := math.Inf(1)

		// Try all subsets of mask with size <= k
		sub := mask
		for sub > 0 {
			// Count bits
			cnt := 0
			for tmp := sub; tmp > 0; tmp &= tmp - 1 {
				cnt++
			}
			if cnt <= k {
				// Compute crossing time
				maxTime := 0
				for i := 0; i < n; i++ {
					if sub&(1<<i) != 0 {
						if time[i] > maxTime {
							maxTime = time[i]
						}
					}
				}
				crossingTime := float64(maxTime) * mul[stage]
				newStage := (stage + int(math.Floor(crossingTime))) % m
				newMask := mask ^ sub

				if newMask == 0 {
					// All transported
					if crossingTime < best {
						best = crossingTime
					}
				} else {
					// Need to return one person
					// Try each person in the boat as returner
					for i := 0; i < n; i++ {
						if sub&(1<<i) != 0 {
							returnTime := float64(time[i]) * mul[newStage]
							newStage2 := (newStage + int(math.Floor(returnTime))) % m
							finalMask := newMask | (1 << i)
							total := crossingTime + returnTime + dfs(finalMask, newStage2)
							if total < best {
								best = total
							}
						}
					}
				}
			}
			sub = (sub - 1) & mask
		}

		memo[mask][stage] = best
		return best
	}

	ans := dfs((1<<n)-1, 0)
	if math.IsInf(ans, 1) {
		return -1
	}
	return ans
}
```
