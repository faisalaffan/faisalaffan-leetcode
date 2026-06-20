# 0914 — X Of A Kind In A Deck Of Cards

## Deskripsi

**Soal:** [0914. X Of A Kind In A Deck Of Cards](https://leetcode.com/problems/x-of-a-kind-in-a-deck-of-cards/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n log m) where m is max count. Space: O(n).  
**Kompleksitas Ruang:** O(n).

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #914: X of a Kind in a Deck of Cards
// https://leetcode.com/problems/x-of-a-kind-in-a-deck-of-cards/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(hasGroupsSizeX([]int{1, 2, 3, 4, 4, 3, 2, 1})) // true
	fmt.Println(hasGroupsSizeX([]int{1, 1, 1, 2, 2, 2, 3, 3})) // false
	fmt.Println(hasGroupsSizeX([]int{1}))                        // false
}

// hasGroupsSizeX checks if the deck can be split into groups of equal size with same values.
// Time: O(n log m) where m is max count. Space: O(n).
func hasGroupsSizeX(deck []int) bool {
  // Membuat map untuk pencarian O(1): key → value
	counts := make(map[int]int)
	for _, v := range deck {
		counts[v]++
	}
	g := -1
	for _, c := range counts {
		if g == -1 {
			g = c
		} else {
			g = gcd(g, c)
		}
	}
	return g >= 2
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
```
