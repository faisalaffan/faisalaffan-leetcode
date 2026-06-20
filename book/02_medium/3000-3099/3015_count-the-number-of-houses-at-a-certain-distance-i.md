# 3015 — Count The Number Of Houses At A Certain Distance I

## Deskripsi

**Soal:** [3015. Count The Number Of Houses At A Certain Distance I](https://leetcode.com/problems/count-the-number-of-houses-at-a-certain-distance-i/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3015: Count the Number of Houses at a Certain Distance I
// https://leetcode.com/problems/count-the-number-of-houses-at-a-certain-distance-i/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(countOfPairs(3, 1, 3))
	fmt.Println(countOfPairs(5, 2, 4))
	fmt.Println(countOfPairs(4, 1, 1))
}

func countOfPairs(n int, x int, y int) []int {
  // Membuat slice untuk menyimpan hasil
	ans := make([]int, n)
	if x > y {
		x, y = y, x
	}
	for i := 1; i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			dist := j - i
			via := abs(i-x) + 1 + abs(y-j)
			if via < dist {
				dist = via
			}
			ans[dist-1] += 2
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```
