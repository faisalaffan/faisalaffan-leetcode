# 0771 — Jewels And Stones

## Deskripsi

**Soal:** [0771. Jewels And Stones](https://leetcode.com/problems/jewels-and-stones/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(j + s). Space: O(j).  
**Kompleksitas Ruang:** O(j).

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #771: Jewels and Stones
// https://leetcode.com/problems/jewels-and-stones/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(numJewelsInStones("aA", "aAAbbbb")) // 3
	fmt.Println(numJewelsInStones("z", "ZZ"))       // 0
	fmt.Println(numJewelsInStones("", "abc"))       // 0
}

// numJewelsInStones counts how many stones are also jewels.
// Time: O(j + s). Space: O(j).
func numJewelsInStones(jewels string, stones string) int {
  // Membuat map untuk pencarian O(1): key → value
	jSet := make(map[rune]bool)
	for _, c := range jewels {
		jSet[c] = true
	}
	count := 0
	for _, c := range stones {
		if jSet[c] {
			count++
		}
	}
	return count
}
```
