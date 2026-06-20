# 0331 — Verify Preorder Serialization Of A Binary Tree

## Deskripsi

**Soal:** [0331. Verify Preorder Serialization Of A Binary Tree](https://leetcode.com/problems/verify-preorder-serialization-of-a-binary-tree/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func isValidSerialization(preorder string) bool`

## Solusi Go

```go
package main

// LeetCode #331: Verify Preorder Serialization of a Binary Tree
// https://leetcode.com/problems/verify-preorder-serialization-of-a-binary-tree/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"strings"
)

func isValidSerialization(preorder string) bool {
	nodes := strings.Split(preorder, ",")
	// Each non-null node consumes 1 slot and creates 2 new slots
	// Each null node consumes 1 slot
	slots := 1
	for _, node := range nodes {
		slots--
		if slots < 0 {
			return false
		}
		if node != "#" {
			slots += 2
		}
	}
	return slots == 0
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", isValidSerialization("9,3,4,#,#,1,#,#,2,#,6,#,#"))
	// Expected: true

	// Test case 2
	fmt.Println("Test 2:", isValidSerialization("1,#"))
	// Expected: false

	// Test case 3
	fmt.Println("Test 3:", isValidSerialization("9,#,#,1"))
	// Expected: false
}
```
