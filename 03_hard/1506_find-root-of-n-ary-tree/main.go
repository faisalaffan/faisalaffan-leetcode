package main

// LeetCode #1506: Find Root of N-Ary Tree
// https://leetcode.com/problems/find-root-of-n-ary-tree/
// Difficulty: Medium (listed here as Hard)
//
// You are given a list of nodes from an N-ary tree. Each node has a value
// and a list of children. Find the root node of the tree.

import "fmt"

// Node represents an N-ary tree node.
type Node struct {
	Val      int
	Children []*Node
}

// findRoot finds the root node from a list of tree nodes.
// The root is the node that never appears as a child.
//
// Approach 1: Use indegree count. The root has indegree 0, all others have indegree >= 1.
// Approach 2: XOR all node values + all child values. The result is the root's value.
//             Then find the node with that value.
//
// We use Approach 1 for clarity and correctness (handles duplicate values).

func findRoot(tree []*Node) *Node {
	if len(tree) == 0 {
		return nil
	}
	if len(tree) == 1 {
		return tree[0]
	}

	// Count indegree: how many times each node appears as a child
	indegree := make(map[*Node]int)
	for _, node := range tree {
		for _, child := range node.Children {
			indegree[child]++
		}
	}

	// The root never appears as a child
	for _, node := range tree {
		if indegree[node] == 0 {
			return node
		}
	}

	return nil // should not happen for a valid tree
}

// findRootByXOR finds the root using XOR of all node values and child values.
// Assumes all node values are unique.
func findRootByXOR(tree []*Node) *Node {
	if len(tree) == 0 {
		return nil
	}
	if len(tree) == 1 {
		return tree[0]
	}

	var xorSum int
	valueToNode := make(map[int]*Node)

	for _, node := range tree {
		valueToNode[node.Val] = node
		xorSum ^= node.Val
		for _, child := range node.Children {
			xorSum ^= child.Val
		}
	}

	// xorSum now equals the root's value (all non-root values cancel out)
	return valueToNode[xorSum]
}

func main() {
	// Build a tree:
	//       1
	//     / | \
	//    2  3  4
	//   /
	//  5
	child5 := &Node{Val: 5}
	child2 := &Node{Val: 2, Children: []*Node{child5}}
	child3 := &Node{Val: 3}
	child4 := &Node{Val: 4}
	root1 := &Node{Val: 1, Children: []*Node{child2, child3, child4}}

	// Shuffle the list (simulates the problem input)
	tree1 := []*Node{child2, root1, child4, child3, child5}

	// Test 1: findRoot
	result1 := findRoot(tree1)
	fmt.Printf("Test 1 - findRoot: Val=%d (expected 1)\n", result1.Val)

	// Test 2: findRootByXOR
	result2 := findRootByXOR(tree1)
	fmt.Printf("Test 2 - findRootByXOR: Val=%d (expected 1)\n", result2.Val)

	// Test 3: Larger tree
	//       10
	//     /    \
	//    20    30
	//   /  \     \
	//  40  50    60
	n40 := &Node{Val: 40}
	n50 := &Node{Val: 50}
	n60 := &Node{Val: 60}
	n20 := &Node{Val: 20, Children: []*Node{n40, n50}}
	n30 := &Node{Val: 30, Children: []*Node{n60}}
	n10 := &Node{Val: 10, Children: []*Node{n20, n30}}

	tree2 := []*Node{n30, n10, n40, n60, n20, n50}
	result3 := findRoot(tree2)
	fmt.Printf("Test 3 - findRoot (larger): Val=%d (expected 10)\n", result3.Val)

	// Test 4: Single node
	tree3 := []*Node{{Val: 42}}
	result4 := findRoot(tree3)
	fmt.Printf("Test 4 - Single node: Val=%d (expected 42)\n", result4.Val)

	// Test 5: Empty
	result5 := findRoot(nil)
	fmt.Printf("Test 5 - Empty: %v (expected nil)\n", result5)

	// Test 6: Two nodes (root and child)
	c := &Node{Val: 2}
	r := &Node{Val: 1, Children: []*Node{c}}
	tree4 := []*Node{c, r}
	result6 := findRoot(tree4)
	fmt.Printf("Test 6 - Two nodes: Val=%d (expected 1)\n", result6.Val)
}
