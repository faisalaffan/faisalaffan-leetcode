# 1431 — Kids With The Greatest Number Of Candies

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func kidsWithCandies(candies []int, extraCandies int) []bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

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
	res := make([]bool, len(candies))
	for i, c := range candies {
		res[i] = c+extraCandies >= maxCandy
	}
	return res
}
```
