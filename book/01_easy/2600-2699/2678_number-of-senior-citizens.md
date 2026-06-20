# 2678 — Number Of Senior Citizens

## Deskripsi

**Soal:** [2678. Number Of Senior Citizens](https://leetcode.com/problems/number-of-senior-citizens/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2678: Number of Senior Citizens
// https://leetcode.com/problems/number-of-senior-citizens/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(NumberOfSeniorCitizens([]string{"7868190130M7522", "5303914400F9211", "9273338290F4010"}))
	fmt.Println(NumberOfSeniorCitizens([]string{"1313579440F2036", "2921522980M5644"}))
}

func NumberOfSeniorCitizens(details []string) int {
	count := 0
	for _, d := range details {
		age, _ := strconv.Atoi(d[11:13])
		if age > 60 {
			count++
		}
	}
	return count
}
```
