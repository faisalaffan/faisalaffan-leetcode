# 3870 — Count Commas In Range

## Deskripsi

**Soal:** [3870. Count Commas In Range](https://leetcode.com/problems/count-commas-in-range/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(log n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3870: Count Commas in Range
// https://leetcode.com/problems/count-commas-in-range/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountCommasInRange(1002))
	fmt.Println(CountCommasInRange(998))
	fmt.Println(CountCommasInRange(1500000))
}

// Time: O(log n)
// Space: O(1)
func CountCommasInRange(n int) int {
	total := 0
	threshold := 1000
	for n >= threshold {
		total += n - threshold + 1
		threshold *= 1000
	}
	return total
}
```
