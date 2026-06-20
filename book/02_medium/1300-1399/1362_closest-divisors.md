# 1362 — Closest Divisors

## Deskripsi

**Soal:** [1362. Closest Divisors](https://leetcode.com/problems/closest-divisors/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(sqrt(num)) for finding divisors  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1362: Closest Divisors
// https://leetcode.com/problems/closest-divisors/
// Difficulty: Medium

import "fmt"
import "math"

func main() {
	// Test case 1
	fmt.Println(closestDivisors(8)) // [3,3]

	// Test case 2
	fmt.Println(closestDivisors(123)) // [5,25]

	// Test case 3
	fmt.Println(closestDivisors(999)) // [25,40]

	// Test case 4
	fmt.Println(closestDivisors(1)) // [1,2]
}

// Time: O(sqrt(num)) for finding divisors
// Space: O(1)
func closestDivisors(num int) []int {
	// Check num+1 and num+2 for closest divisor pair
	result := []int{0, 0}
	minDiff := math.MaxInt32

	for n := num + 1; n <= num+2; n++ {
		for i := int(math.Sqrt(float64(n))); i >= 1; i-- {
			if n%i == 0 {
				j := n / i
				diff := j - i
				if diff < minDiff {
					minDiff = diff
					result[0] = i
					result[1] = j
				}
				break
			}
		}
	}

	return result
}
```
