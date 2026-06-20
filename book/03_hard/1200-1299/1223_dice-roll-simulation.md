# 1223 — Dice Roll Simulation

## Deskripsi

**Soal:** [1223. Dice Roll Simulation](https://leetcode.com/problems/dice-roll-simulation/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP)

## Solusi Go

```go
package main

// LeetCode #1223: Dice Roll Simulation
// https://leetcode.com/problems/dice-roll-simulation/
// Difficulty: Hard

import "fmt"

func main() {
	fmt.Println("1223. Dice Roll Simulation")
	fmt.Println("n=2, rollMax=[1,1,2,2,2,3]:", dieSimulator(2, []int{1, 1, 2, 2, 2, 3}), "(expected 34)")
	fmt.Println("n=1, rollMax=[1,1,2,2,2,3]:", dieSimulator(1, []int{1, 1, 2, 2, 2, 3}), "(expected 6)")
	fmt.Println("n=3, rollMax=[1,1,1,1,1,1]:", dieSimulator(3, []int{1, 1, 1, 1, 1, 1}), "(expected 150)")
}

func dieSimulator(n int, rollMax []int) int {
	const MOD = 1000000007
	// dp[face][cnt] for current position
	var dp [6][16]int
	for f := 0; f < 6; f++ {
		dp[f][1] = 1
	}

	for i := 1; i < n; i++ {
		var ndp [6][16]int
		for last := 0; last < 6; last++ {
			for cnt := 1; cnt <= rollMax[last]; cnt++ {
				if dp[last][cnt] == 0 {
					continue
				}
				for nxt := 0; nxt < 6; nxt++ {
					if nxt == last {
						if cnt+1 <= rollMax[nxt] {
							ndp[nxt][cnt+1] = (ndp[nxt][cnt+1] + dp[last][cnt]) % MOD
						}
					} else {
						ndp[nxt][1] = (ndp[nxt][1] + dp[last][cnt]) % MOD
					}
				}
			}
		}
		dp = ndp
	}

	result := 0
	for f := 0; f < 6; f++ {
		for cnt := 1; cnt <= rollMax[f]; cnt++ {
			result = (result + dp[f][cnt]) % MOD
		}
	}
	return result
}
```
