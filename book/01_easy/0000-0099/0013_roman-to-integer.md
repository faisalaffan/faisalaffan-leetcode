# 0013 — Roman To Integer

## Deskripsi

**Soal:** [0013. Roman To Integer](https://leetcode.com/problems/roman-to-integer/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func RomanToInt(s string) int`

## Solusi Go

```go
package main

// LeetCode #13: Roman to Integer
// https://leetcode.com/problems/roman-to-integer/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(1)
func RomanToInt(s string) int {
	vals := map[byte]int{
		'I': 1, 'V': 5, 'X': 10, 'L': 50,
		'C': 100, 'D': 500, 'M': 1000,
	}
	sum, prev := 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		cur := vals[s[i]]
		if cur < prev {
			sum -= cur
		} else {
			sum += cur
		}
		prev = cur
	}
	return sum
}

func main() {
	fmt.Println(RomanToInt("III"))
	fmt.Println(RomanToInt("LVIII"))
	fmt.Println(RomanToInt("MCMXCIV"))
}
```
