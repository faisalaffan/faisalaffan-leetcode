# 2868 — The Wording Game

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func canAliceWin(aliceWords, bobWords []string) bool
```

> **💡 Hint:** Process words grouped by length (descending). Maintain 26-bit bitmasks

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Dynamic Programming, Bitmask

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2868: The Wording Game
// https://leetcode.com/problems/the-wording-game/
// Difficulty: Hard [Paid]
//
// Alice and Bob pick words alternately from their respective lists. Alice goes first.
// Rules:
//   1. Alice's first word must start with 'a'.
//   2. Words must be picked in strictly increasing length.
//   3. Each word must start with the same letter as the previous word, or a letter
//      that comes later in the alphabet.
// A player who cannot make a valid move loses. Determine if Alice can force a win.
//
// Approach: Process words grouped by length (descending). Maintain 26-bit bitmasks
// for each player indicating which letters are available at or above the current length.
// DP state: dp[letter][playerTurn] = can this player force a win from this state?
// A player wins if there exists a word they own of length >= current length with
// letter >= current letter such that the resulting state is losing for the opponent.

import (
	"fmt"
	"sort"
)

func canAliceWin(aliceWords, bobWords []string) bool {
	// Collect word lengths and group by letter for each player
	// map[length] -> 26-bit mask for each player
  // Membuat map (HashMap) — pencarian O(1)
	wordsByLen := make(map[int][2]int) // [aliceMask, bobMask]

	for _, w := range aliceWords {
		l := len(w)
		mask := wordsByLen[l]
		mask[0] |= 1 << (w[0] - 'a')
		wordsByLen[l] = mask
	}
	for _, w := range bobWords {
		l := len(w)
		mask := wordsByLen[l]
		mask[1] |= 1 << (w[0] - 'a')
		wordsByLen[l] = mask
	}

	// Collect and sort unique lengths
  // Alokasi slice integer
	lengths := make([]int, 0, len(wordsByLen))
	for l := range wordsByLen {
		lengths = append(lengths, l)
	}
  // Urutkan secara ascending — O(n log n)
	sort.Ints(lengths)

	// dp[c][p] = can player p force a win when it's their turn and
	// the minimum letter they can use is 'a' + c.
	// We process lengths from largest to smallest.
	dp := [2][26]bool{} // dp[player][letter]

	for i := len(lengths) - 1; i >= 0; i-- {
		ln := lengths[i]
		mask := wordsByLen[ln]

		// newdp for this length
		var ndp [2][26]bool

		for p := 0; p < 2; p++ {
			playerMask := mask[p]
			oppMask := mask[1-p]
			_ = oppMask

			for c := 0; c < 26; c++ {
				// Option 1: skip this length (use a longer word later)
				// dp[p][c] = dp[p][c] from previous iteration (already copied? No)
				// Actually we need to consider: if we skip, the state stays the same,
				// and the next length (shorter) will be processed later.
				// Since we process lengths descending, dp[p][c] currently has the
				// result for lengths > ln. If we skip length ln, we use dp[p][c].

				// Check if player p can pick a word of this length with letter >= c
				canWin := false
				for letter := c; letter < 26; letter++ {
					if playerMask&(1<<letter) != 0 {
						// Player picks this word. Next player's turn with letter = letter,
						// needing a word of length > ln.
						nextState := dp[1-p][letter]
						if !nextState {
							canWin = true
							break
						}
					}
				}
				if !canWin {
					// Player can't win at this length; they may still win at longer lengths
					canWin = dp[p][c]
				}
				ndp[p][c] = canWin
			}
		}
		dp = ndp
	}

	return dp[0][0] // Alice's turn, needs letter 'a'
}

func main() {
	// Example: ["ab","abc"], Bob: ["b","c"] => Alice wins
	fmt.Println(canAliceWin([]string{"ab", "abc"}, []string{"b", "c"}))

	// Alice has 'a' words, Bob none
	fmt.Println(canAliceWin([]string{"a"}, []string{}))

	// Alice has no 'a' words => loses immediately
	fmt.Println(canAliceWin([]string{"b"}, []string{"a"}))

	// More complex
	fmt.Println(canAliceWin([]string{"a", "aa", "aaa"}, []string{"b", "bb"}))

	// Equal footing
	fmt.Println(canAliceWin([]string{"a", "aa"}, []string{"b", "bb"}))

	// Bob has advantage
	fmt.Println(canAliceWin([]string{"a"}, []string{"b", "bb", "bbb"}))
}
```
