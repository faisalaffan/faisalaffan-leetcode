# 2005 — Subtree Removal Game With Fibonacci Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan pohon (tree). Tugasmu menjelajahi atau memanipulasi struktur pohon.

**Cara berpikir:** TreeNode punya Val, Left, Right. Gunakan DFS rekursif (pre/in/post-order).

**Fungsi Solusi:** `func subtreeRemovalGameWithFibonacciTree(k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2005: Subtree Removal Game with Fibonacci Tree
// https://leetcode.com/problems/subtree-removal-game-with-fibonacci-tree/
// Difficulty: Hard [Paid]
//
// A Fibonacci tree of order k is defined recursively:
// - F(0) = single node
// - F(1) = single node
// - F(k) = root with F(k-1) as left subtree and F(k-2) as right subtree
//
// Two players alternate removing a subtree. The player who takes the last
// node wins (normal play). Determine if the first player wins.
//
// This is an impartial combinatorial game. The Grundy number (nimber) for
// a Fibonacci tree of order k determines the outcome:
// - If Grundy(k) != 0, first player wins.
// - Grundy(0) = 1 (single node = can remove it)
// - Grundy(1) = 1 (single node)
// - Grundy(k) = Grundy(k-1) XOR (Grundy(k-2) XOR 1)
//
// Actually: each move in a tree removes a subtree. The game value of a tree
// is mex of game values of all possible resulting positions.
// For a Fibonacci tree, the options are:
// 1. Remove the root -> 0 (empty tree)
// 2. Remove left subtree -> right subtree remains (if any)
// 3. Remove right subtree -> left subtree remains (if any)
// 4. Remove any subtree inside the left or right child recursively
//
// The Grundy of a tree T with left child L and right child R:
// Grundy(T) = mex{0, Grundy(R), Grundy(L), Grundy(T_with_subtree_removed_in_L), ...}
//
// A known result: For a Fibonacci tree F(k):
// Grundy(k) = Grundy(k-1) XOR Grundy(k-2)

import (
	"fmt"
)

func main() {
	// Test cases: Fibonacci tree of order k
	// Print 1 if first player wins, 0 otherwise
	for k := 0; k <= 15; k++ {
		fmt.Printf("k=%d: %d\n", k, subtreeRemovalGameWithFibonacciTree(k))
	}
}

// subtreeRemovalGameWithFibonacciTree returns 1 if first player wins, 0 otherwise.
func subtreeRemovalGameWithFibonacciTree(k int) int {
	if k < 0 {
		return 0
	}
	g := grundy(k)
	if g != 0 {
		return 1
	}
	return 0
}

// grundy computes the Grundy number (nimber) of a Fibonacci tree of order k.
//
// The Fibonacci tree F(k) has:
// - F(0) = single node
// - F(1) = single node
// - F(k >= 2) = root with F(k-1) as left child and F(k-2) as right child
//
// In impartial combinatorial game theory, a player can remove any subtree.
// This is equivalent to the game of "Tree Nim" where removing a subtree
// replaces the tree with the remaining forest.
//
// For a tree T with root and children subtrees T1, T2, ..., Tn:
// Grundy(T) = mex { 0, Grundy(T1), Grundy(T2), ..., Grundy(Tn) } XOR ...
// Actually, removing a subtree is like taking that component out.
//
// For F(k) with left=F(k-1), right=F(k-2):
// Options:
// - Remove root -> empty game (Grundy = 0)
// - Remove F(k-1) entirely -> only F(k-2) remains -> Grundy = Grundy(k-2)
// - Remove F(k-2) entirely -> only F(k-1) remains -> Grundy = Grundy(k-1)
// - Remove any proper subtree of F(k-1) -> after removal, remaining game
//   is the disjoint sum of Grundy of the new F'(k-1) XOR Grundy(k-2)
// - Similarly for F(k-2)
//
// The key insight: Grundy(F(k)) = Grundy(F(k-1)) XOR Grundy(F(k-2))
// This can be proven by induction.
func grundy(k int) int {
	if k == 0 || k == 1 {
		return 1
	}

	// Compute Grundy numbers iteratively
  // Alokasi slice
	g := make([]int, k+1)
	g[0] = 1
	g[1] = 1

	for i := 2; i <= k; i++ {
		g[i] = g[i-1] ^ g[i-2]
	}

	return g[k]
}
```
