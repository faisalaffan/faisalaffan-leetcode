# 1432 — Max Difference You Can Get From Changing An Integer

## Deskripsi

**Soal:** [1432. Max Difference You Can Get From Changing An Integer](https://leetcode.com/problems/max-difference-you-can-get-from-changing-an-integer/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n) where n = number of digits  
**Kompleksitas Ruang:** O(n) for string conversion

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1432: Max Difference You Can Get From Changing an Integer
// https://leetcode.com/problems/max-difference-you-can-get-from-changing-an-integer/
// Difficulty: Medium

import "fmt"
import "strconv"

func main() {
	// Test case 1
	fmt.Println(maxDiff(555)) // 888

	// Test case 2
	fmt.Println(maxDiff(9)) // 8

	// Test case 3
	fmt.Println(maxDiff(123456)) // 820000

	// Test case 4
	fmt.Println(maxDiff(10000)) // 20000

	// Test case 5
	fmt.Println(maxDiff(9288)) // 8700
}

// Time: O(n) where n = number of digits
// Space: O(n) for string conversion
func maxDiff(num int) int {
	s := strconv.Itoa(num)
	digits := []byte(s)

	// Find max: replace first non-'9' with '9'
  // Membuat slice untuk menyimpan hasil
	maxDigits := make([]byte, len(digits))
	copy(maxDigits, digits)
	targetMax := byte(0)
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(maxDigits); i++ {
		if maxDigits[i] != '9' {
			targetMax = maxDigits[i]
			break
		}
	}
	if targetMax != 0 {
  // Loop standar: indeks 0 sampai n-1
		for i := 0; i < len(maxDigits); i++ {
			if maxDigits[i] == targetMax {
				maxDigits[i] = '9'
			}
		}
	}
	maxVal, _ := strconv.Atoi(string(maxDigits))

	// Find min: replace first non-'0'/'1' appropriately
  // Membuat slice untuk menyimpan hasil
	minDigits := make([]byte, len(digits))
	copy(minDigits, digits)

	// If first digit is not '1', replace it with '1'
	// Otherwise find first digit > '1' to replace with '0' (but not leading)
	targetMin := byte(0)
	replacementMin := byte(0)

	if minDigits[0] != '1' {
		targetMin = minDigits[0]
		replacementMin = '1'
	} else {
		for i := 1; i < len(minDigits); i++ {
			if minDigits[i] != '0' && minDigits[i] != '1' {
				targetMin = minDigits[i]
				replacementMin = '0'
				break
			}
		}
	}

	if targetMin != 0 {
  // Loop standar: indeks 0 sampai n-1
		for i := 0; i < len(minDigits); i++ {
			if minDigits[i] == targetMin {
				minDigits[i] = replacementMin
			}
		}
	}
	minVal, _ := strconv.Atoi(string(minDigits))

	return maxVal - minVal
}
```
