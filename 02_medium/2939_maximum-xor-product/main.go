package main

// LeetCode #2939: Maximum Xor Product
// https://leetcode.com/problems/maximum-xor-product/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(maximumXorProduct(12, 5, 4))
	fmt.Println(maximumXorProduct(6, 7, 5))
	fmt.Println(maximumXorProduct(1, 6, 3))
}

func maximumXorProduct(a int64, b int64, n int) int {
	const mod int64 = 1e9 + 7
	ax := (a >> n) << n
	bx := (b >> n) << n
	for i := n - 1; i >= 0; i-- {
		x, y := (a>>i)&1, (b>>i)&1
		if x == y {
			ax |= 1 << i
			bx |= 1 << i
		} else if ax < bx {
			ax |= 1 << i
		} else {
			bx |= 1 << i
		}
	}
	ax %= mod
	bx %= mod
	return int(ax * bx % mod)
}
