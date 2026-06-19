package main

// LeetCode #2117: Abbreviating the Product of a Range
// https://leetcode.com/problems/abbreviating-the-product-of-a-range/
// Difficulty: Hard
//
// Compute the product of [left, right], then abbreviate as "first5...last5eZ".
// Count trailing zeros by tracking 2s and 5s. Use modular arithmetic for last
// 5 digits and logarithms for the first 5 digits.

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	// Example 1
	fmt.Println(abbreviateProduct(1, 4))

	// Example 2
	fmt.Println(abbreviateProduct(2, 11))

	// Example 3
	fmt.Println(abbreviateProduct(999998, 1000000))

	// Single number
	fmt.Println(abbreviateProduct(5, 5))

	// No trailing zeros
	fmt.Println(abbreviateProduct(1, 3))
}

func abbreviateProduct(left int, right int) string {
	count2 := 0
	count5 := 0

	// Count factors of 2 and 5 in the range
	for i := left; i <= right; i++ {
		x := i
		for x%2 == 0 {
			count2++
			x /= 2
		}
		for x%5 == 0 {
			count5++
			x /= 5
		}
	}

	zeros := count2
	if count5 < zeros {
		zeros = count5
	}
	extra2 := count2 - zeros
	extra5 := count5 - zeros

	// Compute last 5 digits (mod 100000) after removing all 2s and 5s
	last5 := 1
	for i := left; i <= right; i++ {
		x := i
		for x%2 == 0 {
			x /= 2
		}
		for x%5 == 0 {
			x /= 5
		}
		last5 = (last5 * (x % 100000)) % 100000
	}

	// Multiply back remaining 2s or 5s
	for i := 0; i < extra2; i++ {
		last5 = (last5 * 2) % 100000
	}
	for i := 0; i < extra5; i++ {
		last5 = (last5 * 5) % 100000
	}

	// Compute log10 of product to get first 5 digits
	var logSum float64 = 0
	for i := left; i <= right; i++ {
		logSum += math.Log10(float64(i))
	}

	// Remove trailing zeros from log
	productLog := logSum - float64(zeros)
	totalDigits := int(math.Floor(productLog)) + 1

	if totalDigits <= 5 {
		// The product without trailing zeros is small enough to show fully
		fullNum := 1
		for i := left; i <= right; i++ {
			fullNum *= i
		}
		// Remove trailing zeros
		for i := 0; i < zeros; i++ {
			fullNum /= 10
		}
		if zeros > 0 {
			return fmt.Sprintf("%de%d", fullNum, zeros)
		}
		return fmt.Sprintf("%d", fullNum)
	}

	// First 5 digits using log
	frac := productLog - math.Floor(productLog)
	first5 := int(math.Pow(10, frac+4))

	// Handle the "last 5" formatting (pad with leading zeros if needed)
	last5Str := fmt.Sprintf("%05d", last5)

	return fmt.Sprintf("%d...%se%d", first5, last5Str, zeros)
}

// Remove unused import
var _ = strings.Builder{}
