# 0339 — Nested List Weight Sum

## Deskripsi

**Soal:** [0339. Nested List Weight Sum](https://leetcode.com/problems/nested-list-weight-sum/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(d)

**Algoritma:** LIS (Longest Increasing Subsequence)

**Fungsi Solusi:** `func NewInt(val int) *NestedInteger`

## Solusi Go

```go
package main

// LeetCode #339: Nested List Weight Sum
// https://leetcode.com/problems/nested-list-weight-sum/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(d)

import "fmt"

// NestedInteger represents a nested integer structure (LeetCode interface)
type NestedInteger struct {
	integer int
	list    []*NestedInteger
	isInt   bool
}

func NewInt(val int) *NestedInteger {
	return &NestedInteger{integer: val, isInt: true}
}

func NewList(items ...*NestedInteger) *NestedInteger {
	return &NestedInteger{list: items, isInt: false}
}

func (n NestedInteger) IsInteger() bool           { return n.isInt }
func (n NestedInteger) GetInteger() int            { return n.integer }
func (n NestedInteger) GetList() []*NestedInteger  { return n.list }

func depthSum(nestedList []*NestedInteger) int {
	var dfs func(list []*NestedInteger, depth int) int
	dfs = func(list []*NestedInteger, depth int) int {
		total := 0
		for _, ni := range list {
			if ni.IsInteger() {
				total += ni.GetInteger() * depth
			} else {
				total += dfs(ni.GetList(), depth+1)
			}
		}
		return total
	}
	return dfs(nestedList, 1)
}

func main() {
	// Test case 1: [[1,1],2,[1,1]]
	n1 := NewList(NewInt(1), NewInt(1))
	n2 := NewInt(2)
	n3 := NewList(NewInt(1), NewInt(1))
	fmt.Println("Test 1:", depthSum([]*NestedInteger{n1, n2, n3}))
	// Expected: 10 (2*1 + 1*2 + 1*2 + 1*2 + 1*2)

	// Test case 2: [1,[4,[6]]]
	inner := NewList(NewInt(6))
	mid := NewList(NewInt(4), inner)
	fmt.Println("Test 2:", depthSum([]*NestedInteger{NewInt(1), mid}))
	// Expected: 27 (1*1 + 4*2 + 6*3)

	// Test case 3: Single integer
	fmt.Println("Test 3:", depthSum([]*NestedInteger{NewInt(5)}))
	// Expected: 5
}
```
