# 1881 — Maximum Value After Insertion

## Deskripsi

**Soal:** [1881. Maximum Value After Insertion](https://leetcode.com/problems/maximum-value-after-insertion/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1881: Maximum Value After Insertion
// https://leetcode.com/problems/maximum-value-after-insertion/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MaxValue("99", 9))
	fmt.Println(MaxValue("-13", 2))
	fmt.Println(MaxValue("73", 6))
}

// Time: O(n), Space: O(n)
func MaxValue(n string, x int) string {
  // Membuat slice untuk menyimpan hasil
	result := make([]byte, 0, len(n)+1)
	negative := n[0] == '-'

	if negative {
		result = append(result, '-')
		inserted := false
		for i := 1; i < len(n); i++ {
			digit := int(n[i] - '0')
			if !inserted && x < digit {
				result = append(result, byte(x+'0'))
				inserted = true
			}
			result = append(result, n[i])
		}
		if !inserted {
			result = append(result, byte(x+'0'))
		}
	} else {
		inserted := false
  // Loop standar: indeks 0 sampai n-1
		for i := 0; i < len(n); i++ {
			digit := int(n[i] - '0')
			if !inserted && x > digit {
				result = append(result, byte(x+'0'))
				inserted = true
			}
			result = append(result, n[i])
		}
		if !inserted {
			result = append(result, byte(x+'0'))
		}
	}
	return string(result)
}
```
