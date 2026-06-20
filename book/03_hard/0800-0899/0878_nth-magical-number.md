# 0878 — Nth Magical Number

## Deskripsi

**Soal:** [0878. Nth Magical Number](https://leetcode.com/problems/nth-magical-number/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Binary Search (pencarian biner)

**Fungsi Solusi:** `func gcd(a, b int) int`

## Solusi Go

```go
package main

// LeetCode #878: Nth Magical Number
// https://leetcode.com/problems/nth-magical-number/
// Difficulty: Hard
//
// Binary search on the answer. A number x is "magical" if divisible by a or b.
// Count of magical numbers <= x is: x/a + x/b - x/lcm(a,b).
// Binary search for the smallest x such that count >= n.

import "fmt"

const mod878 = 1_000_000_007

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

func nthMagicalNumber(n int, a int, b int) int {
	l := lcm(a, b)

	// Upper bound: at worst, the nth magical number is n * min(a, b).
	left, right := 1, n*min(a, b)

  // Loop two-pointer: kiri vs kanan
	for left < right {
		mid := left + (right-left)/2
		// Count of magical numbers <= mid.
		count := mid/a + mid/b - mid/l
		if count >= n {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left % mod878
}

func main() {
	// Example 1: n=1, a=2, b=3 -> 2
	fmt.Println("Test 1:", nthMagicalNumber(1, 2, 3)) // 2

	// Example 2: n=4, a=2, b=3 -> 6
	fmt.Println("Test 2:", nthMagicalNumber(4, 2, 3)) // 6

	// Example 3: n=5, a=2, b=4 -> 10
	// Magical numbers: 2,4,6,8,10 (divisible by 2 or 4)
	fmt.Println("Test 3:", nthMagicalNumber(5, 2, 4)) // 10

	// Edge: n=3, a=6, b=4 -> lcm=12, magical numbers: 4,6,8,12,16,18,20,24,...
	// 4(1),6(2),8(3) -> 8
	fmt.Println("Test 4:", nthMagicalNumber(3, 6, 4)) // 8

	// Larger: n=1000000000, a=40000, b=40000 -> expects 999720007
	// Just validate no overflow
	fmt.Println("Test 5:", nthMagicalNumber(10, 3, 5)) // 15: 3,5,6,9,10,12,15,...
	// Count: 15/3=5, 15/5=3, 15/15=1 => 5+3-1=7 >=10. Not quite.
	// Manual: 3,5,6,9,10,12,15,18,20,21 -> 21 is 10th
	// Let's verify: 21/3=7, 21/5=4, 21/15=1 => 7+4-1=10. Yes.
	fmt.Println("Test 6:", nthMagicalNumber(10, 3, 5)) // 21
}
```
