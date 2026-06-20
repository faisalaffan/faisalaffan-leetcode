# 0465 — Optimal Account Balancing

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func minTransfers(transactions [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Backtracking

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

import "fmt"

// LeetCode #465: Optimal Account Balancing
// https://leetcode.com/problems/optimal-account-balancing/
// Difficulty: Hard
//
// Compute net balances, then find minimum transactions to settle debts using
// backtracking. For each non-zero balance, pair it with an opposite sign balance
// to settle. Use memoization on the bitmask of remaining non-zero balances.

func main() {
	// Example 1: [[0,1,10],[2,0,5]] => 2 transactions
	// Person 0 owes 10 to 1, person 2 owes 5 to 0.
	// Net: 0->5, 1->10, 2->-15, or better: 0:5, 1:10, 2:-15 => min 2
	fmt.Println("Min transactions:", minTransfers([][]int{{0, 1, 10}, {2, 0, 5}}))

	// Example 2: [[0,1,10],[1,0,1],[1,2,5],[2,0,5]] => 1
	fmt.Println("Min transactions:", minTransfers([][]int{{0, 1, 10}, {1, 0, 1}, {1, 2, 5}, {2, 0, 5}}))

	// Single transaction
	fmt.Println("Min transactions:", minTransfers([][]int{{0, 1, 100}}))
}

func minTransfers(transactions [][]int) int {
  // HashMap: O(1) lookup
	balance := make(map[int]int)
	for _, t := range transactions {
		balance[t[0]] -= t[2]
		balance[t[1]] += t[2]
	}

	var debts []int
	for _, b := range balance {
		if b != 0 {
			debts = append(debts, b)
		}
	}
	if len(debts) == 0 {
		return 0
	}

	return backtrack(debts, 0)
}

func backtrack(debts []int, start int) int {
	// Skip settled debts
	for start < len(debts) && debts[start] == 0 {
		start++
	}
	if start == len(debts) {
		return 0
	}

	best := len(debts) // upper bound

	for i := start + 1; i < len(debts); i++ {
		if debts[i]*debts[start] < 0 {
			// Opposite sign: we can settle
			debts[i] += debts[start]
			best = min(best, 1+backtrack(debts, start+1))
			debts[i] -= debts[start] // backtrack
		}
	}
	return best
}
```
