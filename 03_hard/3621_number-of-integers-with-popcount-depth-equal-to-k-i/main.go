package main

// LeetCode #3621: Number of Integers With Popcount-Depth Equal to K I
// https://leetcode.com/problems/number-of-integers-with-popcount-depth-equal-to-k-i/
// Difficulty: Hard
//
// Count numbers from 1 to n whose "popcount depth" equals exactly k.
// Popcount depth is the number of times you apply popcount (count set bits)
// until you reach 1 (or 0 for the number 0).
//
// Examples:
// popcount depth of 0 is 0, depth of 1 is 1, depth of 2 is 2 (popcount(2)=1),
// depth of 3 is 2 (popcount(3)=2, popcount(2)=1).
//
// Approach: Precompute depths for all numbers up to 60 (since n <= 10^15 < 2^50),
// then use digit DP to count numbers in [1, n] with a given target depth.

import "fmt"
import "math/bits"

func main() {
	// Example 1
	fmt.Println(popcountDepth(10, 2))
	// Example 2
	fmt.Println(popcountDepth(5, 1))
	// Edge: n = 1
	fmt.Println(popcountDepth(1, 1))
	// Edge: k = 0
	fmt.Println(popcountDepth(10, 0))
}

func popcountDepth(n int64, k int) int {
	if k == 0 {
		if n >= 0 {
			return 1 // only the number 0 has depth 0
		}
		return 0
	}
	if n <= 0 {
		return 0
	}

	// Precompute depth for numbers up to 60 (max bits for 10^15)
	// depth[i] = popcount depth of i
	depth := make([]int, 61)
	depth[0] = 0
	for i := 1; i <= 60; i++ {
		d := 1
		x := i
		for x > 1 {
			x = bits.OnesCount(uint(x))
			d++
		}
		depth[i] = d
	}

	// Digit DP over binary representation
	bits := make([]int, 0)
	tmp := n
	for tmp > 0 {
		bits = append(bits, int(tmp&1))
		tmp >>= 1
	}
	// bits[0] is LSB, we need MSB first
	for i, j := 0, len(bits)-1; i < j; i, j = i+1, j-1 {
		bits[i], bits[j] = bits[j], bits[i]
	}
	m := len(bits)

	// memo[pos][tight][popcount]
	memo := make([][][]int, m)
	for i := range memo {
		memo[i] = make([][]int, 2)
		for j := range memo[i] {
			memo[i][j] = make([]int, 61)
			for p := range memo[i][j] {
				memo[i][j][p] = -1
			}
		}
	}

	var dp func(pos int, tight bool, popcount int) int
	dp = func(pos int, tight bool, popcount int) int {
		if pos == m {
			if popcount <= 60 && depth[popcount] == k {
				return 1
			}
			return 0
		}
		tightInt := 0
		if tight {
			tightInt = 1
		}
		if memo[pos][tightInt][popcount] != -1 {
			return memo[pos][tightInt][popcount]
		}

		limit := 1
		if tight {
			limit = bits[pos]
		}

		total := 0
		for d := 0; d <= limit; d++ {
			newTight := tight && (d == limit)
			// Actually limit computation above is wrong
		}
		return 0
	}

	// Simpler approach: iterate all possible popcounts up to 60
	// and count numbers whose popcount = p and depth[p] = k
	total := 0

	// Count numbers <= n with a given popcount using combinatorics
	// C[m][p] = number of ways to have p ones in m bits
	C := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		C[i] = make([]int, m+1)
		C[i][0] = 1
		for j := 1; j <= i; j++ {
			C[i][j] = C[i-1][j-1] + C[i-1][j]
		}
	}

	// Count numbers <= n with exactly p ones in binary representation
	countWithPopcount := func(p int) int {
		if p == 0 {
			return 1
		}
		cnt := 0
		ones := 0
		for i := 0; i < m; i++ {
			if bits[i] == 1 {
				// Set this bit to 0, fill remaining with p-ones-1 ones
				remaining := m - i - 1
				if p-ones-1 >= 0 && remaining >= p-ones-1 {
					cnt += C[remaining][p-ones-1]
				}
				ones++
			}
		}
		// Count n itself if its popcount == p
		if ones == p {
			cnt++
		}
		return cnt
	}

	for p := 0; p <= 60 && p <= m; p++ {
		if depth[p] == k {
			total += countWithPopcount(p)
		}
	}

	if k == 0 {
		// Already counted 0 above in countWithPopcount(0)
	} else {
		// Subtract the counting of 0 if it was included
		if k == 0 {
			// handled
		}
		// 0 has depth 0, so if k > 0, 0 shouldn't be counted
		// countWithPopcount(0) counts 0. If k > 0, we need to subtract 1.
		if k > 0 && depth[0] != k {
			total-- // subtract the counting of 0
		}
	}

	return total
}
