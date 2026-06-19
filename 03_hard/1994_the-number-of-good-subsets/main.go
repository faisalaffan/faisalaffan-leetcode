package main

// LeetCode #1994: The Number of Good Subsets
// https://leetcode.com/problems/the-number-of-good-subsets/
// Difficulty: Hard
// Approach: DP bitmask over the 10 primes up to 30.
// A "good" subset has no repeated prime factors in its product,
// so each number's prime factorization must use each prime at most once.
// "1" is special (can be included any number of times, factor is 2^count).

import "fmt"

const MOD1994 = 1000000007

var primes1994 = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}

// primeMask returns the bitmask of prime factors for a number.
// Returns -1 if the number has a repeated prime factor (e.g., 4, 8, 9, 12).
func primeMask(num int) int {
	mask := 0
	for i, p := range primes1994 {
		if num%p == 0 {
			num /= p
			if num%p == 0 {
				return -1 // repeated prime factor
			}
			mask |= 1 << i
		}
	}
	if num > 1 {
		return -1 // has a prime factor > 30
	}
	return mask
}

func powMod(a, e, mod int) int {
	res := 1
	for e > 0 {
		if e&1 == 1 {
			res = (res * a) % mod
		}
		a = (a * a) % mod
		e >>= 1
	}
	return res
}

func numberOfGoodSubsets(nums []int) int {
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}

	masks := make(map[int]int) // num -> mask (-1 if invalid)
	for k := range freq {
		if k == 1 {
			continue
		}
		masks[k] = primeMask(k)
	}

	dp := make([]int, 1<<10)
	dp[0] = 1

	for num, mask := range masks {
		if mask == -1 || freq[num] == 0 {
			continue
		}
		count := freq[num]
		// For each existing mask, try adding this number
		// Iterate in reverse to avoid using the same number multiple times
		for m := (1 << 10) - 1; m >= 0; m-- {
			if dp[m] == 0 {
				continue
			}
			if m&mask == 0 {
				dp[m|mask] = (dp[m|mask] + dp[m]*count) % MOD1994
			}
		}
	}

	ans := 0
	for m := 1; m < 1<<10; m++ {
		ans = (ans + dp[m]) % MOD1994
	}

	// Multiply by 2^freq[1] (each "1" can be independently included or not)
	ans = ans * powMod(2, freq[1], MOD1994) % MOD1994

	return ans
}

func main() {
	// Example: [1,2,3,4] -> 6
	fmt.Println(numberOfGoodSubsets([]int{1, 2, 3, 4}))

	// Additional tests
	fmt.Println(numberOfGoodSubsets([]int{1, 1, 2, 3, 4}))
	fmt.Println(numberOfGoodSubsets([]int{2, 3, 5}))
	fmt.Println(numberOfGoodSubsets([]int{4, 8, 9}))
}
