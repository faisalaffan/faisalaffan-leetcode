package main

// LeetCode #364: Nested List Weight Sum II
// https://leetcode.com/problems/nested-list-weight-sum-ii/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(d)

import "fmt"

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

func (n NestedInteger) IsInteger() bool          { return n.isInt }
func (n NestedInteger) GetInteger() int           { return n.integer }
func (n NestedInteger) GetList() []*NestedInteger { return n.list }

func depthSumInverse(nestedList []*NestedInteger) int {
	// First pass: find max depth
	var maxDepth func(list []*NestedInteger, depth int) int
	maxDepth = func(list []*NestedInteger, depth int) int {
		maxD := depth
		for _, ni := range list {
			if !ni.IsInteger() {
				if d := maxDepth(ni.GetList(), depth+1); d > maxD {
					maxD = d
				}
			}
		}
		return maxD
	}

	md := maxDepth(nestedList, 1)

	// Second pass: compute weighted sum
	var dfs func(list []*NestedInteger, depth int) int
	dfs = func(list []*NestedInteger, depth int) int {
		total := 0
		for _, ni := range list {
			if ni.IsInteger() {
				total += ni.GetInteger() * (md - depth + 1)
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
	fmt.Println("Test 1:", depthSumInverse([]*NestedInteger{n1, n2, n3}))
	// Expected: 8 (deepest=2, 1*1+1*1 + 2*2 + 1*1+1*1)

	// Test case 2: [1,[4,[6]]]
	inner := NewList(NewInt(6))
	mid := NewList(NewInt(4), inner)
	fmt.Println("Test 2:", depthSumInverse([]*NestedInteger{NewInt(1), mid}))
	// Expected: 17 (1*3 + 4*2 + 6*1)

	// Test case 3: Single integer
	fmt.Println("Test 3:", depthSumInverse([]*NestedInteger{NewInt(5)}))
	// Expected: 5
}
