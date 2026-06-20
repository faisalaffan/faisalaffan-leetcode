# 0288 — Unique Word Abbreviation

## Deskripsi

**Soal:** [0288. Unique Word Abbreviation](https://leetcode.com/problems/unique-word-abbreviation/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n) for init, O(1) for isUnique, Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func Constructor(dictionary []string) ValidWordAbbr`

## Solusi Go

```go
package main

// LeetCode #288: Unique Word Abbreviation
// https://leetcode.com/problems/unique-word-abbreviation/
// Difficulty: Medium [Paid]
// Time: O(n) for init, O(1) for isUnique, Space: O(n)

import (
	"fmt"
	"strconv"
)

type ValidWordAbbr struct {
	abbrMap map[string]string
}

func Constructor(dictionary []string) ValidWordAbbr {
  // Membuat map untuk pencarian O(1): key → value
	abbrMap := make(map[string]string)
	for _, word := range dictionary {
		abbr := getAbbr(word)
		if existing, ok := abbrMap[abbr]; ok {
			if existing != word {
				abbrMap[abbr] = ""
			}
		} else {
			abbrMap[abbr] = word
		}
	}
	return ValidWordAbbr{abbrMap}
}

func (this *ValidWordAbbr) IsUnique(word string) bool {
	abbr := getAbbr(word)
	val, ok := this.abbrMap[abbr]
	return !ok || val == word
}

func getAbbr(s string) string {
	if len(s) <= 2 {
		return s
	}
	return string(s[0]) + strconv.Itoa(len(s)-2) + string(s[len(s)-1])
}

func main() {
	vwa := Constructor([]string{"deer", "door", "cake", "card"})
	fmt.Println(vwa.IsUnique("dear"))
	fmt.Println(vwa.IsUnique("cart"))
	fmt.Println(vwa.IsUnique("cane"))
	fmt.Println(vwa.IsUnique("make"))

	vwa2 := Constructor([]string{"a", "a"})
	fmt.Println(vwa2.IsUnique("a"))
}
```
