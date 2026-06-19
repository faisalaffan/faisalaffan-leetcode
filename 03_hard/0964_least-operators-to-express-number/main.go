package main

// LeetCode #964: Least Operators to Express Number
// https://leetcode.com/problems/least-operators-to-express-number/
// Difficulty: Hard

import "fmt"

func leastOpsExpressTarget(x int, target int) int {
	// Convert target to base x (least significant first)
	var digits []int
	for t := target; t > 0; t /= x {
		digits = append(digits, t%x)
	}
	n := len(digits)

	// cost[pos] = operators needed to produce x^pos as a term
	// cost[0] = 2 (x/x = 1)
	// cost[pos] = pos (x repeated pos times with pos-1 multiplications)
	cost := make([]int, n+1)
	cost[0] = 2
	for i := 1; i <= n; i++ {
		cost[i] = i
	}

	memo := make([][]int, n+1)
	for i := range memo {
		memo[i] = []int{-1, -1}
	}

	// For safety, process up to n (positions beyond the original digits have digit=0)
	var dfs func(pos int, carry int) int
	dfs = func(pos int, carry int) int {
		if pos > n {
			// If we have carry beyond all digit positions
			if carry == 0 {
				return 0
			}
			return carry * cost[pos]
		}

		if pos == n {
			if carry == 0 {
				return 0
			}
			return carry * cost[n]
		}

		if memo[pos][carry] != -1 {
			return memo[pos][carry]
		}

		// For positions beyond len(digits), digit is 0
		d := carry
		if pos < len(digits) {
			d = digits[pos] + carry
		}
		digit := d % x
		newCarry := d / x

		// Option A: use digit additions of x^pos
		best := digit*cost[pos] + dfs(pos+1, newCarry)

		// Option B: borrow from next power
		// Add (x-digit) copies of x^pos, then subtract x^(pos+1)
		// The expression: ... + (x-digit)*x^pos - x^(pos+1) + ...
		// Cost: (x-digit) * cost[pos]  (for the (x-digit)*x^pos terms)
		//      + 1                     (for the '-' operator)
		//      + cost[pos+1]           (for the x^(pos+1) term itself)
		//      + dfs(pos+1, newCarry+1) (for remaining digits)
		borrowCost := (x-digit)*cost[pos] + 1 + cost[pos+1] + dfs(pos+1, newCarry+1)
		if borrowCost < best {
			best = borrowCost
		}

		memo[pos][carry] = best
		return best
	}

	// Subtract 1 because the first term doesn't have a leading '+' or '-'
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
	// Expected: 0 (edge case)

	fmt.Println("Example 5:")
	fmt.Println(leastOpsExpressTarget(2, 1))
	// Expected: 1 (x/x = 1, that's 2 operators? Let's check... actually 1 operator for the /)
}
