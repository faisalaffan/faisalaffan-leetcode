# 0964 — Least Operators To Express Number

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func leastOpsExpressTarget(x int, target int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #964: Least Operators to Express Number
// https://leetcode.com/problems/least-operators-to-express-number/
// Difficulty: Hard

import "fmt"

func leastOpsExpressTarget(x int, target int) int {
	if target == 0 {
		return 0
	}

	// Convert target to base x (least significant first)
	var digits []int
	for t := target; t > 0; t /= x {
		digits = append(digits, t%x)
	}
	n := len(digits)

	// cost[pos] = operators needed to produce x^pos
	// cost[0] = 2 (x/x = 1)
	// cost[pos] = pos for pos >= 1 (x * x * ... * x, pos-1 multiplications)
  // Alokasi slice
	cost := make([]int, n+2)
	cost[0] = 2
	for i := 1; i <= n+1; i++ {
		cost[i] = i
	}

  // Matriks 2D
	memo := make([][]int, n+2)
  // Range loop
	for i := range memo {
		memo[i] = []int{-1, -1}
	}

	var dfs func(pos int, carry int) int
	dfs = func(pos int, carry int) int {
		if pos == n {
			if carry == 0 {
				return 0
			}
			return carry * cost[n]
		}

		if memo[pos][carry] != -1 {
			return memo[pos][carry]
		}

		d := carry
		if pos < len(digits) {
			d = digits[pos] + carry
		}
		digit := d % x
		newCarry := d / x

		// Option A: use digit additions of x^pos
		best := digit*cost[pos] + dfs(pos+1, newCarry)

		// Option B: borrow from next power (only when digit > 0)
		if digit > 0 {
			// Add (x-digit) copies of x^pos and subtract x^(pos+1)
			// cost: (x-digit)*cost[pos] + '-' operator + cost[pos+1]
			borrowCost := (x-digit)*cost[pos] + cost[pos+1] + dfs(pos+1, newCarry)
			if borrowCost < best {
				best = borrowCost
			}
		}

		memo[pos][carry] = best
		return best
	}

	return dfs(0, 0) - 1
}

func main() {
	fmt.Println("Example 1:")
	fmt.Println(leastOpsExpressTarget(3, 19))
	// Expected: 5

	fmt.Println("Example 2:")
	fmt.Println(leastOpsExpressTarget(5, 501))
	// Expected: 8

	fmt.Println("Example 3:")
	fmt.Println(leastOpsExpressTarget(100, 100000000))
	// Expected: 3

	fmt.Println("Example 4:")
	fmt.Println(leastOpsExpressTarget(2, 0))
	// Expected: 0

	fmt.Println("Example 5:")
	fmt.Println(leastOpsExpressTarget(2, 1))
	// Expected: 1 (x/x = one operator)
}
```
