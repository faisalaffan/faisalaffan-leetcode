# 0691 — Stickers To Spell Word

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func minStickers(stickers []string, target string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, DP, Bitmask

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
	"math"
)

// LeetCode #691: Stickers to Spell Word
// https://leetcode.com/problems/stickers-to-spell-word/
// Difficulty: Hard
//
// DP with bitmask. For each state of covered target characters, try each
// sticker. Memoized DFS to find min stickers.

func main() {
	// ["with","example","science"], "thehat" => 3
	fmt.Println(minStickers([]string{"with", "example", "science"}, "thehat"))
	// ["notice","possible"], "basicbasic" => -1
	fmt.Println(minStickers([]string{"notice", "possible"}, "basicbasic"))
	// ["a"], "aa" => 2
	fmt.Println(minStickers([]string{"a"}, "aa"))
	// ["these","guess","about","garden","him"], "atomher" => 3
	fmt.Println(minStickers([]string{"these", "guess", "about", "garden", "him"}, "atomher"))
	// Single sticker multiple chars
	fmt.Println(minStickers([]string{"abc"}, "abc"))
}

func minStickers(stickers []string, target string) int {
	t := len(target)
  // Alokasi slice
	targetCount := make([]int, 26)
	for _, ch := range target {
		targetCount[ch-'a']++
	}

	var freqList [][]int
	for _, s := range stickers {
  // Alokasi slice
		freq := make([]int, 26)
		for _, ch := range s {
			freq[ch-'a']++
		}
		useful := false
		for i := 0; i < 26; i++ {
			if freq[i] > 0 && targetCount[i] > 0 {
				useful = true
				break
			}
		}
		if useful {
			freqList = append(freqList, freq)
		}
	}

  // HashMap: O(1) lookup
	memo := make(map[int]int)
	fullMask := (1 << t) - 1

	var dp func(mask int) int
	dp = func(mask int) int {
		if mask == fullMask {
			return 0
		}
		if v, ok := memo[mask]; ok {
			return v
		}

		first := -1
		for i := 0; i < t; i++ {
			if mask&(1<<i) == 0 {
				first = i
				break
			}
		}

		ans := math.MaxInt32
		for _, freq := range freqList {
			ch := target[first] - 'a'
			if freq[ch] == 0 {
				continue
			}
			newMask := mask
  // Alokasi slice
			remain := make([]int, 26)
			copy(remain, freq)
			for i := first; i < t; i++ {
				if newMask&(1<<i) == 0 {
					ch2 := target[i] - 'a'
					if remain[ch2] > 0 {
						remain[ch2]--
						newMask |= (1 << i)
					}
				}
			}
			sub := dp(newMask)
			if sub != -1 && sub+1 < ans {
				ans = sub + 1
			}
		}

		if ans == math.MaxInt32 {
			ans = -1
		}
		memo[mask] = ans
		return ans
	}

	return dp(0)
}
```
