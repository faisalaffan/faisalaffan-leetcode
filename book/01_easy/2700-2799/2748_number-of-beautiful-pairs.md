# 2748 — Number Of Beautiful Pairs

## Deskripsi

**Soal:** [2748. Number Of Beautiful Pairs](https://leetcode.com/problems/number-of-beautiful-pairs/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2748: Number of Beautiful Pairs
// https://leetcode.com/problems/number-of-beautiful-pairs/
// Difficulty: Easy
// Time: O(n^2) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(NumberOfBeautifulPairs([]int{2, 5, 1, 4}))
	fmt.Println(NumberOfBeautifulPairs([]int{11, 21, 12}))
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func firstDigit(n int) int {
	for n >= 10 {
		n /= 10
	}
	return n
}

func lastDigit(n int) int {
	return n % 10
}

func NumberOfBeautifulPairs(nums []int) int {
	count := 0
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if gcd(firstDigit(nums[i]), lastDigit(nums[j])) == 1 {
				count++
			}
		}
	}
	return count
}
```
