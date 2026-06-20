# 3270 — Find The Key Of The Numbers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func minDigit(num1, num2, num3, place int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(1). Space: O(1).  |  **Ruang:** O(1).


## 💻 Solusi Go

```go
package main

// LeetCode #3270: Find the Key of the Numbers
// https://leetcode.com/problems/find-the-key-of-the-numbers/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindTheKeyOfTheNumbers(1, 10, 1000))
	fmt.Println(FindTheKeyOfTheNumbers(987, 879, 798))
	fmt.Println(FindTheKeyOfTheNumbers(123, 456, 789))
}

// minDigit returns the minimum digit in a number at a given decimal place (1, 10, 100, 1000).
func minDigit(num1, num2, num3, place int) int {
	d1 := (num1 / place) % 10
	d2 := (num2 / place) % 10
	d3 := (num3 / place) % 10
	minD := d1
	if d2 < minD {
		minD = d2
	}
	if d3 < minD {
		minD = d3
	}
	return minD
}

// FindTheKeyOfTheNumbers returns a 4-digit key by taking the minimum digit at each position across three numbers.
// Time: O(1). Space: O(1).
func FindTheKeyOfTheNumbers(num1 int, num2 int, num3 int) int {
	key := 0
	places := []int{1000, 100, 10, 1}
	for _, p := range places {
		key = key*10 + minDigit(num1, num2, num3, p)
	}
	return key
}
```
