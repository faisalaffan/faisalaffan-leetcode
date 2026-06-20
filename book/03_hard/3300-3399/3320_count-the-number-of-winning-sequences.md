# 3320 — Count The Number Of Winning Sequences

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func countWinningSequences(s string) int
```

> **💡 Hint:** DP with memoization. State: (index, scoreDiff, lastMove).

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3320: Count The Number of Winning Sequences
// https://leetcode.com/problems/count-the-number-of-winning-sequences/
// Difficulty: Hard
//
// Alice and Bob play a game with creatures: F (Fire Dragon), W (Water Serpent),
// E (Earth Golem). The win rule: F beats E, E beats W, W beats F.
// Alice's moves are given. Bob cannot repeat the same creature twice in a row.
// Count the number of sequences Bob can use to have a strictly higher total
// score than Alice after n rounds.
//
// Approach: DP with memoization. State: (index, scoreDiff, lastMove).
// Score diff can range from -n to n, use offset.

import "fmt"

func main() {
	// Example 1
	fmt.Println(countWinningSequences("FFF"))
	// Example 2
	fmt.Println(countWinningSequences("FWEFW"))
	// Edge: single round where Bob can win
	fmt.Println(countWinningSequences("F"))
	// Edge: Alice always wins
	fmt.Println(countWinningSequences("W"))
}

const MOD = 1000000007

var scoreMap = map[byte]int{'F': 0, 'W': 1, 'E': 2}

func countWinningSequences(s string) int {
	n := len(s)
	// dp[i][diff][last] where diff is offset by n
	offset := n
  // Membuat matriks/slice 2D untuk DP
	dp := make([][][]int, n+1)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = make([][]int, 2*n+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, 4) // 0=no prev, 1=F, 2=W, 3=E
		}
	}
	dp[0][offset][0] = 1 // no last move

	for i := 0; i < n; i++ {
		alice := scoreMap[s[i]]
		for diff := 0; diff <= 2*n; diff++ {
			for last := 0; last <= 3; last++ {
				cur := dp[i][diff][last]
				if cur == 0 {
					continue
				}
				for bob := 0; bob < 3; bob++ {
					if bob+1 == last {
						continue // cannot repeat
					}
					newDiff := diff
					if bob == alice {
						// tie, no change
					} else if (bob == 0 && alice == 2) || (bob == 1 && alice == 0) || (bob == 2 && alice == 1) {
						newDiff++ // Bob wins
					} else {
						newDiff-- // Alice wins
					}
					dp[i+1][newDiff][bob+1] = (dp[i+1][newDiff][bob+1] + cur) % MOD
				}
			}
		}
	}

	var ans int
	for diff := offset + 1; diff <= 2*n; diff++ {
		for last := 1; last <= 3; last++ {
			ans = (ans + dp[n][diff][last]) % MOD
		}
	}
	return ans
}
```
