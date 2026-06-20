# 0443 — String Compression

## Deskripsi

**Soal:** [0443. String Compression](https://leetcode.com/problems/string-compression/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func compress(chars []byte) int`

## Solusi Go

```go
package main

// LeetCode #443: String Compression
// https://leetcode.com/problems/string-compression/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {
	write := 0
	read := 0

	for read < len(chars) {
		ch := chars[read]
		count := 0

		// Count consecutive chars
		for read < len(chars) && chars[read] == ch {
			read++
			count++
		}

		// Write char
		chars[write] = ch
		write++

		// Write count
		if count > 1 {
			for _, d := range strconv.Itoa(count) {
				chars[write] = byte(d)
				write++
			}
		}
	}
	return write
}

func main() {
	// Test case 1
	chars1 := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	l1 := compress(chars1)
	fmt.Println("Test 1:", l1, string(chars1[:l1]))
	// Expected: 6, "a2b2c3"

	// Test case 2
	chars2 := []byte{'a'}
	l2 := compress(chars2)
	fmt.Println("Test 2:", l2, string(chars2[:l2]))
	// Expected: 1, "a"

	// Test case 3
	chars3 := []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}
	l3 := compress(chars3)
	fmt.Println("Test 3:", l3, string(chars3[:l3]))
	// Expected: 4, "ab12"
}
```
