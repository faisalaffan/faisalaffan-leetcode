# 1431 — Kids With The Greatest Number Of Candies

## Deskripsi

**Soal:** [1431. Kids With The Greatest Number Of Candies](https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func kidsWithCandies(candies []int, extraCandies int) []bool`

## Solusi Go

```go
package main

// LeetCode #1431: Kids With the Greatest Number of Candies
// https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/
// Difficulty: Easy
//
// LeetCode submission: func kidsWithCandies(candies []int, extraCandies int) []bool

import "fmt"

func main() {
	fmt.Println(KidsWithTheGreatestNumberOfCandies([]int{2, 3, 5, 1, 3}, 3)) // [true true true false true]
	fmt.Println(KidsWithTheGreatestNumberOfCandies([]int{4, 2, 1, 1, 2}, 1)) // [true false false false false]
}

// Time: O(n), Space: O(n)
func KidsWithTheGreatestNumberOfCandies(candies []int, extraCandies int) []bool {
	maxCandy := 0
	for _, c := range candies {
		if c > maxCandy {
			maxCandy = c
		}
	}
  // Membuat slice untuk menyimpan hasil
	res := make([]bool, len(candies))
	for i, c := range candies {
		res[i] = c+extraCandies >= maxCandy
	}
	return res
}
```
