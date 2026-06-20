# 2283 — Check If Number Has Equal Digit Count And Digit Value

## Deskripsi

**Soal:** [2283. Check If Number Has Equal Digit Count And Digit Value](https://leetcode.com/problems/check-if-number-has-equal-digit-count-and-digit-value/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2283: Check if Number Has Equal Digit Count and Digit Value
// https://leetcode.com/problems/check-if-number-has-equal-digit-count-and-digit-value/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(CheckIfNumberHasEqualDigitCountAndDigitValue("1210")) // true
	fmt.Println(CheckIfNumberHasEqualDigitCountAndDigitValue("030"))  // false
}

func CheckIfNumberHasEqualDigitCountAndDigitValue(num string) bool {
	count := [10]int{}
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(num); i++ {
		count[num[i]-'0']++
	}
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(num); i++ {
		if count[i] != int(num[i]-'0') {
			return false
		}
	}
	return true
}
```
