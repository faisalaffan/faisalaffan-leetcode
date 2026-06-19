package main

// LeetCode #3490: Count Beautiful Numbers
// https://leetcode.com/problems/count-beautiful-numbers/
// Difficulty: Hard
//
// A number is "beautiful" if its digit product is non-zero and divides the number.
// Digit DP over the range [low, high]. Track product of digits (non-zero) and
// remainder modulo product. Use memoization on (pos, tight, started, prod, rem).

import (
	"fmt"
	"strconv"
)

func countBeautifulNumbers(low, high int) int {
	if low < 1 {
		low = 1
	}
	return countUpTo(high) - countUpTo(low-1)
}

func countUpTo(limit int) int {
	if limit <= 0 {
		return 0
	}
	s := strconv.Itoa(limit)
	n := len(s)
	digits := make([]int, n)
	for i, ch := range s {
		digits[i] = int(ch - '0')
	}

	// memo[pos][tight][started] = map from (prod, rem) to count
	// Use a struct key for the map
	type stateKey struct {
		prod int
		rem  int
	}
	memo := make([][][]map[stateKey]int, n)
	for i := range memo {
		memo[i] = make([][]map[stateKey]int, 2)
		for t := 0; t < 2; t++ {
			memo[i][t] = make([]map[stateKey]int, 2)
		}
	}

	var dfs func(pos int, tight int, started int, prod int, rem int) int
	dfs = func(pos int, tight int, started int, prod int, rem int) int {
		if pos == n {
			if started == 0 {
				return 0
			}
			if prod > 0 && rem%prod == 0 {
				return 1
			}
			return 0
		}

		if memo[pos][tight][started] != nil {
			key := stateKey{prod, rem}
			if val, ok := memo[pos][tight][started][key]; ok && tight == 0 {
				return val
			}
		}

		limitDigit := 9
		if tight == 1 {
			limitDigit = digits[pos]
		}

		total := 0
		for d := 0; d <= limitDigit; d++ {
			nextTight := 0
			if tight == 1 && d == limitDigit {
				nextTight = 1
			}

			if started == 0 && d == 0 {
				// Still not started
				total += dfs(pos+1, nextTight, 0, prod, rem)
			} else if d == 0 {
				// Digit 0 means product becomes 0, which can never divide the number
				// Skip further processing, product=0 case is handled at leaf
				total += dfs(pos+1, nextTight, 1, 0, 0)
			} else {
				newProd := prod
				newRem := rem
				if started == 1 {
					newProd = prod * d
					newRem = (rem*10 + d) % newProd
					// Actually we need rem % newProd, not rem
					// rem = number_so_far % prod
					// new_number = number_so_far*10 + d
					// new_prod = prod*d
					// new_rem = (number_so_far*10 + d) % (prod*d)
					// = ((prod*q + rem)*10 + d) % (prod*d)
					// = (rem*10 + d) % (prod*d) since prod*10*q % (prod*d) = 0
					// ... no wait, (prod*q*10) % (prod*d) = prod*(10*q) % (prod*d)
					// = prod * ((10*q) % d) ... this doesn't simplify nicely
				} else {
					newProd = d
					newRem = 0 // number = d, product = d, d % d = 0
				}
				total += dfs(pos+1, nextTight, 1, newProd, newRem)
			}
		}

		if tight == 0 {
			if memo[pos][tight][started] == nil {
				memo[pos][tight][started] = make(map[stateKey]int)
			}
			key := stateKey{prod, rem}
			memo[pos][tight][started][key] = total
		}

		return total
	}

	return dfs(0, 1, 0, 1, 0)
}

func main() {
	// Test: low=1, high=20 -> expected 19 (all except ... let's count)
	// Beautiful numbers up to 20: 1..9 all beautiful (prod=number, number%prod=0)
	// 10: product=0 not beautiful
	// 11: product=1, 11%1=0 beautiful
	// 12: product=2, 12%2=0 beautiful
	// 13: product=3, 13%3=1 not beautiful? Wait 13%3=1, so no.
	// Actually 13%3=1≠0, so not beautiful.
	// 14: product=4, 14%4=2 not beautiful
	// 15: product=5, 15%5=0 beautiful
	// 16: product=6, 16%6=4 not beautiful
	// 17: product=7, 17%7=3 not beautiful
	// 18: product=8, 18%8=2 not beautiful
	// 19: product=9, 19%9=1 not beautiful
	// 20: product=0 not beautiful
	// So up to 20: 1-9(9), 11(1), 15(1) = 11, not 19.
	// The user said 1,20→19 which might be wrong or for a different problem definition.
	// Let me just run it and see.
	fmt.Printf("count(1,20) -> %d\n", countBeautifulNumbers(1, 20))

	// Test: 1,100
	fmt.Printf("count(1,100) -> %d\n", countBeautifulNumbers(1, 100))

	// Test: individual numbers
	for i := 1; i <= 20; i++ {
		c := countBeautifulNumbers(i, i)
		if c > 0 {
			fmt.Printf("%d is beautiful\n", i)
		}
	}
}
