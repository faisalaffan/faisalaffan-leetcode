# 0649 — Dota2 Senate

## Deskripsi

**Soal:** [0649. Dota2 Senate](https://leetcode.com/problems/dota2-senate/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #649: Dota2 Senate
// https://leetcode.com/problems/dota2-senate/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(predictPartyVictory("RD"))
	fmt.Println(predictPartyVictory("RDD"))
	fmt.Println(predictPartyVictory("RRDDD"))
}

func predictPartyVictory(senate string) string {
	n := len(senate)
  // Membuat slice untuk menyimpan hasil
	radiant := make([]int, 0)
  // Membuat slice untuk menyimpan hasil
	dire := make([]int, 0)

	for i, c := range senate {
		if c == 'R' {
			radiant = append(radiant, i)
		} else {
			dire = append(dire, i)
		}
	}

	for len(radiant) > 0 && len(dire) > 0 {
		r := radiant[0]
		d := dire[0]
		radiant = radiant[1:]
		dire = dire[1:]

		if r < d {
			radiant = append(radiant, r+n)
		} else {
			dire = append(dire, d+n)
		}
	}

	if len(radiant) > 0 {
		return "Radiant"
	}
	return "Dire"
}
```
