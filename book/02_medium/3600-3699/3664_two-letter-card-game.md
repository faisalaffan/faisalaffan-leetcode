# 3664 — Two Letter Card Game

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func twoLetterCardGame(cards []string, x byte) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer

**Waktu:** O(n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3664: Two-Letter Card Game
// https://leetcode.com/problems/two-letter-card-game/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func twoLetterCardGame(cards []string, x byte) int {
	var left [26]int
	var right [26]int
	xx := 0

	for _, c := range cards {
		a, b := c[0], c[1]
		if a == x && b == x {
			xx++
		} else if a == x {
			left[b-'a']++
		} else if b == x {
			right[a-'a']++
		}
	}

	pairsLeft := computePairs(left[:])
	pairsRight := computePairs(right[:])

	leftoverL := sumArr(left[:]) - 2*pairsLeft
	leftoverR := sumArr(right[:]) - 2*pairsRight
	leftovers := leftoverL + leftoverR

	useWithXX := xx
	if useWithXX > leftovers {
		useWithXX = leftovers
	}
	xxLeft := xx - useWithXX

	extra := xxLeft / 2
	totalPairs := pairsLeft + pairsRight
	if extra > totalPairs {
		extra = totalPairs
	}

	return pairsLeft + pairsRight + useWithXX + extra
}

func computePairs(buckets []int) int {
	total := 0
	maxBucket := 0
	for _, v := range buckets {
		total += v
		if v > maxBucket {
			maxBucket = v
		}
	}
	if total < 2 {
		return 0
	}
	half := total / 2
	diff := total - maxBucket
	if half < diff {
		return half
	}
	return diff
}

func sumArr(arr []int) int {
	s := 0
	for _, v := range arr {
		s += v
	}
	return s
}

func main() {
	fmt.Println(twoLetterCardGame([]string{"aa", "ab", "ba", "ac"}, 'a'))
	fmt.Println(twoLetterCardGame([]string{"ab", "bc", "cd"}, 'a'))
	fmt.Println(twoLetterCardGame([]string{"xx", "xa", "xb", "ax", "bx"}, 'x'))
}
```
