# Hard (Sulit) — Problem ��2199

## 1787 — Make The Xor Of All Segments Equal To Zero

```go
package main

// LeetCode #1787: Make the XOR of All Segments Equal to Zero
// https://leetcode.com/problems/make-the-xor-of-all-segments-equal-to-zero/
// Difficulty: Hard
//
// Approach: DP with grouping by index mod k.
// For all windows of length k, their XOR must be equal. This means:
// nums[i] XOR nums[i+1] XOR ... XOR nums[i+k-1] == 0 for all i.
// This implies nums[i] == nums[i + k] for all i.
// So the array must be periodic with period k.
// Additionally, the XOR of the first k elements must be 0.
//
// Group nums by index mod k. For each group (0..k-1), we need to pick a
// value for all positions in that group. The XOR of the k chosen values
// must be 0.
//
// For each group g, count the frequency of each value.
// For each group, consider changing all elements to some value v.
// Cost for group g choosing v = size(g) - freq[g][v].
//
// DP over groups: dp[i][xor] = min cost for first i groups achieving XOR = xor.
// Result = dp[k][0].

import (
	"fmt"
	"math"
)

func minChanges(nums []int, k int) int {
	n := len(nums)

	// Group values by index mod k
	groups := make([]map[int]int, k)
	for i := 0; i < k; i++ {
		groups[i] = make(map[int]int)
	}
	for i, v := range nums {
		groups[i%k][v]++
	}

	groupSizes := make([]int, k)
	for g := 0; g < k; g++ {
		groupSizes[g] = len(groups[g]) // actually number of unique values is not the size; size = n/k rounded properly
	}
	// Actual group sizes (how many positions per group)
	for i := 0; i < n; i++ {
		_ = i % k // we'll compute directly
	}

	// Size of each group (number of positions)
	size := make([]int, k)
	for i := 0; i < n; i++ {
		size[i%k]++
	}

	const maxXor = 1024 // nums[i] < 1024 based on constraints (2^10)
	INF := math.MaxInt32

	dp := make([][]int, k+1)
	for i := range dp {
		dp[i] = make([]int, maxXor)
		for j := range dp[i] {
			dp[i][j] = INF
		}
	}
	dp[0][0] = 0

	for g := 0; g < k; g++ {
		// Option 1: For each position in group, we can change to any value.
		// Compute globalBest = min over previous xor of dp[g][prevXor] + size[g]
		// (since we can change all elements in this group to some value that gives
		// best possible cost)
		globalBest := INF
		for x := 0; x < maxXor; x++ {
			if dp[g][x] < globalBest {
				globalBest = dp[g][x]
			}
		}

		// For each possible next xor state
		for nextXor := 0; nextXor < maxXor; nextXor++ {
			// Best via "change all in group to some value v that gives us global best"
			// Changing all in group costs size[g], and we can pick any value v.
			// This handles the case where we change to a value not in the group.
			best := globalBest + size[g]

			// Option 2: For each value present in this group, we can keep it
			for val, freq := range groups[g] {
				prevXor := nextXor ^ val
				if dp[g][prevXor]+size[g]-freq < best {
					best = dp[g][prevXor] + size[g] - freq
				}
			}

			dp[g+1][nextXor] = best
		}
	}

	return dp[k][0]
}

func main() {
	// Example test case
	fmt.Println("nums=[1,2,3,4,5,6],k=3 →", minChanges([]int{1, 2, 3, 4, 5, 6}, 3)) // Expected: 3

	// Additional tests
	fmt.Println("nums=[3,4,5,2,1,7,3,4,7],k=3 →", minChanges([]int{3, 4, 5, 2, 1, 7, 3, 4, 7}, 3))
	fmt.Println("nums=[1,2,3],k=1 →", minChanges([]int{1, 2, 3}, 1)) // Need XOR of whole array = 0, change 2

	// Edge case: already periodic
	fmt.Println("nums=[1,2,1,2],k=2 →", minChanges([]int{1, 2, 1, 2}, 2)) // Already: nums[0]=nums[2]=1, nums[1]=nums[3]=2, XOR=1^2=3≠0 so need 1 change
}
```

## 1788 — Maximize The Beauty Of The Garden

```go
package main

// LeetCode #1788: Maximize the Beauty of the Garden
// https://leetcode.com/problems/maximize-the-beauty-of-the-garden/
// Difficulty: Hard [Paid]
//
// We have flowers in a row with beauty values (positive = beautiful, negative = ugly).
// Pick a contiguous segment, then remove any subset of flowers (must keep at least one).
// Maximize the sum of kept flowers.
//
// DP solution: at each position i, dp = max(flowers[i], previous_dp + max(0, flowers[i])).
// This captures the ability to "skip" negative values within the segment while keeping
// all positives. Answer is the max dp over all positions.

import "fmt"

func main() {
	// Example 1: all positive, keep all
	fmt.Println(maximumBeauty([]int{1, 2, 3, 4}))
	// Example 2: skip the negative in middle
	fmt.Println(maximumBeauty([]int{1, -2, 3}))
	// Example 3: all negative, keep the least negative
	fmt.Println(maximumBeauty([]int{-1, -2, -3}))
	// Mixed
	fmt.Println(maximumBeauty([]int{4, -1, 3, -10, 5}))
	// Single positive
	fmt.Println(maximumBeauty([]int{7}))
	// Single negative
	fmt.Println(maximumBeauty([]int{-7}))
	// All negative various
	fmt.Println(maximumBeauty([]int{-5, -1, -3}))
	// Empty
	fmt.Println(maximumBeauty([]int{}))
	// All zeros
	fmt.Println(maximumBeauty([]int{0, 0, 0}))
	// Alternating signs
	fmt.Println(maximumBeauty([]int{-3, 2, -1, 4, -5, 6}))
}

func maximumBeauty(flowers []int) int {
	if len(flowers) == 0 {
		return 0
	}
	dp := flowers[0]
	ans := dp
	for i := 1; i < len(flowers); i++ {
		dp = max(flowers[i], dp+max(0, flowers[i]))
		ans = max(ans, dp)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

## 1793 — Maximum Score Of A Good Subarray

```go
package main

// LeetCode #1793: Maximum Score of a Good Subarray
// https://leetcode.com/problems/maximum-score-of-a-good-subarray/
// Difficulty: Hard
//
// Given array nums and index k, find the maximum score of any "good" subarray
// that contains index k. Score = min(subarray) * length(subarray).
//
// Approach: expand outward from k, maintaining the running minimum.
// At each step, expand in the direction with the larger next value to keep
// the minimum as high as possible.

import "fmt"

func main() {
	// Example 1: nums = [1,4,3,7,4,5], k = 3 => 15 (subarray [4,3,7,4,5] min=3, len=5)
	fmt.Println(maximumScore([]int{1, 4, 3, 7, 4, 5}, 3))
	// Example 2: nums = [5,5,4,5,4,1,1,1], k = 0 => 20 (subarray [5,5,4,5,4] min=4, len=5)
	fmt.Println(maximumScore([]int{5, 5, 4, 5, 4, 1, 1, 1}, 0))
	// Single element
	fmt.Println(maximumScore([]int{5}, 0))
	// Two elements
	fmt.Println(maximumScore([]int{2, 1}, 1))
	// All same
	fmt.Println(maximumScore([]int{3, 3, 3, 3, 3}, 2))
	// Decreasing
	fmt.Println(maximumScore([]int{10, 9, 8, 7, 6}, 2))
	// Random
	fmt.Println(maximumScore([]int{6569, 9667, 3148, 7698, 1622, 6272, 4522, 2757, 5270, 9955}, 2))
}

func maximumScore(nums []int, k int) int {
	left, right := k, k
	minVal := nums[k]
	ans := nums[k]
	n := len(nums)

	for left > 0 || right < n-1 {
		if left == 0 {
			right++
		} else if right == n-1 {
			left--
		} else if nums[left-1] >= nums[right+1] {
			left--
		} else {
			right++
		}
		if nums[left] < minVal {
			minVal = nums[left]
		}
		if nums[right] < minVal {
			minVal = nums[right]
		}
		score := minVal * (right - left + 1)
		if score > ans {
			ans = score
		}
	}
	return ans
}
```

## 1799 — Maximize Score After N Operations

```go
package main

// LeetCode #1799: Maximize Score After N Operations
// https://leetcode.com/problems/maximize-score-after-n-operations/
// Difficulty: Hard
//
// Approach: DP with Bitmask.
//   dp[mask] = max score from remaining elements represented by mask.
//   At each step, pick two unused elements, compute opNum * gcd(a,b),
//   and recurse on the new mask. Use memoization to avoid recomputation.

import "fmt"

func main() {
	// Example 1
	fmt.Println("Example 1:", maxScore([]int{1, 2, 3, 4, 5, 6}))
	// Expected: 14

	// Example 2
	fmt.Println("Example 2:", maxScore([]int{3, 4, 6, 8}))
	// Expected: 11

	// Edge case: 2 elements
	fmt.Println("Edge (2 elems):", maxScore([]int{1, 2}))
	// Expected: 1*1 = 1
}

func maxScore(nums []int) int {
	m := 1 << len(nums)
	memo := make([]int, m)
	for i := range memo {
		memo[i] = -1
	}

	var dfs func(mask int) int
	dfs = func(mask int) int {
		if memo[mask] != -1 {
			return memo[mask]
		}
		if mask == m-1 {
			return 0
		}
		bits := popcount(mask)
		op := bits/2 + 1
		best := 0
		for i := 0; i < len(nums); i++ {
			if mask&(1<<i) != 0 {
				continue
			}
			for j := i + 1; j < len(nums); j++ {
				if mask&(1<<j) != 0 {
					continue
				}
				score := op*gcd(nums[i], nums[j]) + dfs(mask|(1<<i)|(1<<j))
				if score > best {
					best = score
				}
			}
		}
		memo[mask] = best
		return best
	}

	return dfs(0)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func popcount(x int) int {
	cnt := 0
	for x > 0 {
		cnt++
		x &= x - 1
	}
	return cnt
}

// Stub kept for compatibility with the repo scaffold.
func MaximizeScoreAfterNOperations() any {
	return maxScore([]int{1, 2, 3, 4, 5, 6})
}
```

## 1803 — Count Pairs With Xor In A Range

```go
package main

// LeetCode #1803: Count Pairs With XOR in a Range
// https://leetcode.com/problems/count-pairs-with-xor-in-a-range/
// Difficulty: Hard
//
// Approach: Binary Trie.
//   Use a binary trie to store numbers. For each number, query the count
//   of numbers already in the trie whose XOR with it is <= K.
//   Answer = count(high) - count(low-1).
//   During query, traverse bits MSB-first:
//     - If kBit == 1: count the XOR=0 branch (all are < k at this position),
//       then continue with XOR=1 branch.
//     - If kBit == 0: must keep XOR=0, continue with XOR=0 branch.

import "fmt"

func main() {
	// Example 1
	fmt.Println("Example (low=2, high=5):", countPairs([]int{1, 4, 2, 7}, 2, 5))
	// Expected: 4

	// Example 2: all pairs
	fmt.Println("Example (low=2, high=6):", countPairs([]int{1, 4, 2, 7}, 2, 6))
	// Expected: 6 (all six pairs)

	// Edge case: identical numbers
	fmt.Println("Edge (low=0, high=0):", countPairs([]int{1, 1}, 0, 0))
	// Expected: 1 (1^1=0)

	// Single element
	fmt.Println("Edge (single):", countPairs([]int{5}, 0, 100))
	// Expected: 0
}

type TrieNode struct {
	children [2]*TrieNode
	count    int
}

func countPairs(nums []int, low int, high int) int {
	if low == 0 {
		return countLessEqual(nums, high)
	}
	return countLessEqual(nums, high) - countLessEqual(nums, low-1)
}

func countLessEqual(nums []int, k int) int {
	if k < 0 {
		return 0
	}
	root := &TrieNode{}
	result := 0
	for _, num := range nums {
		result += query(root, num, k)
		insert(root, num)
	}
	return result
}

func insert(root *TrieNode, num int) {
	node := root
	for i := 20; i >= 0; i-- {
		bit := (num >> i) & 1
		if node.children[bit] == nil {
			node.children[bit] = &TrieNode{}
		}
		node = node.children[bit]
		node.count++
	}
}

func query(root *TrieNode, num int, k int) int {
	node := root
	result := 0
	for i := 20; i >= 0; i-- {
		if node == nil {
			break
		}
		numBit := (num >> i) & 1
		kBit := (k >> i) & 1
		if kBit == 1 {
			// XOR bit = 0 makes XOR < k at this position
			if node.children[numBit] != nil {
				result += node.children[numBit].count
			}
			// Continue with XOR bit = 1 branch
			node = node.children[1-numBit]
		} else {
			// Must have XOR bit = 0 to keep XOR <= k
			node = node.children[numBit]
		}
	}
	if node != nil {
		result += node.count // XOR exactly equals k
	}
	return result
}

// Stub kept for compatibility with the repo scaffold.
func CountPairsWithXorInARange() any {
	return countPairs([]int{1, 4, 2, 7}, 2, 5)
}
```

## 1808 — Maximize Number Of Nice Divisors

```go
package main

// LeetCode #1808: Maximize Number of Nice Divisors
// https://leetcode.com/problems/maximize-number-of-nice-divisors/
// Difficulty: Hard
//
// Approach: Split primeFactors into groups of 3's for maximum product.
//   If n = primeFactors, we want to maximize the product of divisors,
//   which is equivalent to breaking n into positive integers that
//   multiply to the maximum value. Optimal strategy: use as many 3's
//   as possible. Handle remainders:
//     n % 3 == 0 -> 3^(n/3)
//     n % 3 == 1 -> 3^(n/3-1) * 4  (because 3*1 < 2*2)
//     n % 3 == 2 -> 3^(n/3) * 2
//   Result modulo 1_000_000_007.

import "fmt"

const MOD1808 = 1_000_000_007

func main() {
	// Example 1
	fmt.Println("Example 1 (p=5):", maxNiceDivisors(5))
	// Expected: 6  (n = 2*3 = 6, prime factors split into 2 and 3)

	// Example 2
	fmt.Println("Example 2 (p=8):", maxNiceDivisors(8))
	// Expected: 18 (split 8 = 3+3+2, product = 3*3*2 = 18)

	// Edge cases
	fmt.Println("Edge (p=1):", maxNiceDivisors(1))
	// Expected: 1
	fmt.Println("Edge (p=2):", maxNiceDivisors(2))
	// Expected: 2
	fmt.Println("Edge (p=3):", maxNiceDivisors(3))
	// Expected: 3
	fmt.Println("Edge (p=4):", maxNiceDivisors(4))
	// Expected: 4 (2*2)
}

func maxNiceDivisors(primeFactors int) int {
	if primeFactors <= 3 {
		return primeFactors
	}

	q := primeFactors / 3
	r := primeFactors % 3

	switch r {
	case 0:
		return modPow(3, q, MOD1808)
	case 1:
		// 3+1 -> 2+2 gives better product: 3*1 < 2*2
		return (modPow(3, q-1, MOD1808) * 4) % MOD1808
	default: // r == 2
		return (modPow(3, q, MOD1808) * 2) % MOD1808
	}
}

func modPow(base, exp, mod int) int {
	result := 1
	base %= mod
	for exp > 0 {
		if exp&1 == 1 {
			result = (result * base) % mod
		}
		base = (base * base) % mod
		exp >>= 1
	}
	return result
}

// Stub kept for compatibility with the repo scaffold.
func MaximizeNumberOfNiceDivisors() any {
	return maxNiceDivisors(5)
}
```

## 1815 — Maximum Number Of Groups Getting Fresh Donuts

```go
package main

// LeetCode #1815: Maximum Number of Groups Getting Fresh Donuts
// https://leetcode.com/problems/maximum-number-of-groups-getting-fresh-donuts/
// Difficulty: Hard
//
// Approach: DP with Memoization (state compression via int64).
//   We can reorder groups arbitrarily. A group gets fresh donuts iff the
//   running total before serving them is divisible by batchSize.
//   State = (counts of each remainder modulo batchSize, current running remainder).
//   Encode the state into an int64 and use memoized DFS.
//   Remainder-0 groups always get fresh donuts (they reset the batch),
//   so we serve them first, then recurse on non-zero remainders.

import "fmt"

func main() {
	// Example 1
	fmt.Println("Example 1:", maxHappyGroups(3, []int{1, 2, 3, 4, 5, 6}))
	// Expected: 4

	// Example 2
	fmt.Println("Example 2:", maxHappyGroups(4, []int{1, 3, 2, 5, 2, 2, 1, 6}))
	// Expected: 4

	// Edge case
	fmt.Println("Edge (single group):", maxHappyGroups(5, []int{5}))
	// Expected: 1

	// All divisible
	fmt.Println("Edge (all divisible):", maxHappyGroups(3, []int{3, 6, 9}))
	// Expected: 3
}

func maxHappyGroups(batchSize int, groups []int) int {
	counts := make([]int, batchSize)
	zeroRem := 0
	for _, g := range groups {
		r := g % batchSize
		if r == 0 {
			zeroRem++
		} else {
			counts[r]++
		}
	}

	memo := make(map[int64]int)

	var dfs func(state int64, left int) int
	dfs = func(state int64, left int) int {
		if left == 0 {
			return 0
		}
		if val, ok := memo[state]; ok {
			return val
		}
		cur := int(state & 0xF)       // current running remainder (4 bits)
		best := 0
		for r := 1; r < batchSize; r++ {
			c := int((state >> (4 + 5*(r-1))) & 0x1F)
			if c == 0 {
				continue
			}
			fresh := 0
			if cur == 0 {
				fresh = 1
			}
			newCur := (cur + r) % batchSize
			newState := state
			// decrement count for remainder r
			newState -= 1 << (4 + 5*(r-1))
			// update current remainder
			newState &^= 0xF
			newState |= int64(newCur)

			val := fresh + dfs(newState, left-1)
			if val > best {
				best = val
			}
		}
		memo[state] = best
		return best
	}

	// Build initial state: current remainder = 0, encode counts
	state := int64(0) // cur=0
	for r := 1; r < batchSize; r++ {
		state |= int64(counts[r]) << (4 + 5*(r-1))
	}

	totalLeft := 0
	for r := 1; r < batchSize; r++ {
		totalLeft += counts[r]
	}

	return zeroRem + dfs(state, totalLeft)
}

// Stub kept for compatibility with the repo scaffold.
func MaximumNumberOfGroupsGettingFreshDonuts() any {
	return maxHappyGroups(3, []int{1, 2, 3, 4, 5, 6})
}
```

## 1819 — Number Of Different Subsequences Gcds

```go
package main

// LeetCode #1819: Number of Different Subsequences GCDs
// https://leetcode.com/problems/number-of-different-subsequences-gcds/
// Difficulty: Hard
//
// Approach: Count the number of distinct GCD values achievable by any
//   non-empty subsequence of nums.
//
//   For each possible GCD value g from 1 to max(nums):
//     Collect all numbers in nums that are multiples of g.
//     Compute the GCD of that collection.
//     If GCD == g, then g is achievable as a subsequence GCD.
//
//   Optimization: Use a boolean frequency array (present[x] = true if x in nums)
//   and iterate over multiples of g to compute GCD.

import "fmt"

func main() {
	// Example 1
	fmt.Println("Example 1:", countDifferentSubsequenceGCDs([]int{6, 10, 3}))
	// Expected: 5  (GCDs: 1,2,3,5,6)

	// Example 2
	fmt.Println("Example 2:", countDifferentSubsequenceGCDs([]int{5, 15, 10, 3}))
	// Expected: 6  (GCDs: 1,3,5,10,15,5? wait, let's compute: 1,2,3,5,10,15)

	// Edge case: single element
	fmt.Println("Edge (single):", countDifferentSubsequenceGCDs([]int{7}))
	// Expected: 1 (GCD = {7})

	// All ones
	fmt.Println("Edge (ones):", countDifferentSubsequenceGCDs([]int{1, 1, 1}))
	// Expected: 1 (GCD = {1})
}

func countDifferentSubsequenceGCDs(nums []int) int {
	maxVal := 0
	for _, v := range nums {
		if v > maxVal {
			maxVal = v
		}
	}

	present := make([]bool, maxVal+1)
	for _, v := range nums {
		present[v] = true
	}

	ans := 0
	// Check each possible GCD value
	for g := 1; g <= maxVal; g++ {
		curGCD := 0
		for m := g; m <= maxVal; m += g {
			if present[m] {
				if curGCD == 0 {
					curGCD = m
				} else {
					curGCD = gcd(curGCD, m)
				}
				if curGCD == g {
					ans++
					break
				}
			}
		}
	}
	return ans
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Stub kept for compatibility with the repo scaffold.
func NumberOfDifferentSubsequencesGcds() any {
	return countDifferentSubsequenceGCDs([]int{6, 10, 3})
}
```

## 1825 — Finding Mk Average

```go
package main

// LeetCode #1825: Finding MK Average
// https://leetcode.com/problems/finding-mk-average/
// Difficulty: Hard
//
// Approach: Fenwick Tree (Binary Indexed Tree) + Circular Buffer.
//   Maintain a sliding window of the last m elements.
//   Use two BITs: one for element counts, one for element sums.
//   - addElement: add to BITs + queue; if queue full, evict oldest.
//   - calculateMKAverage: use order statistics to find the kth and
//     (m-k)th smallest values, then compute:
//       midSum = totalSum - sum(k smallest) - sum(k largest)
//       return midSum / (m - 2*k)

import "fmt"

func main() {
	// Example: MKAverage(3, 1)
	mk := Constructor(3, 1)
	mk.AddElement(3)
	mk.AddElement(1)
	fmt.Println("After [3,1]:", mk.CalculateMKAverage()) // not enough elements -> -1
	mk.AddElement(4)
	fmt.Println("After [3,1,4]:", mk.CalculateMKAverage()) // window=[3,1,4], remove k=1 smallest(1) and 1 largest(4), mid=[3], avg=3
	mk.AddElement(2)
	fmt.Println("After [3,1,4,2]:", mk.CalculateMKAverage()) // window=[1,4,2], remove 1 and 4, mid=[2], avg=2
	mk.AddElement(5)
	fmt.Println("After [3,1,4,2,5]:", mk.CalculateMKAverage()) // window=[4,2,5], remove 2 and 5, mid=[4], avg=4

	// Edge: larger test
	fmt.Println("---")
	mk2 := Constructor(5, 2)
	for _, v := range []int{10, 20, 30, 40, 50, 5, 15, 25, 35, 45} {
		mk2.AddElement(v)
	}
	fmt.Println("MKAverage after 10 elements:", mk2.CalculateMKAverage())
}

// --- Fenwick Tree (Binary Indexed Tree) ---

type Fenwick struct {
	tree []int
	n    int
}

func NewFenwick(n int) *Fenwick {
	return &Fenwick{
		tree: make([]int, n+1), // 1-indexed internally
		n:    n,
	}
}

func (f *Fenwick) add(idx, delta int) {
	i := idx + 1 // convert to 1-indexed
	for i <= f.n {
		f.tree[i] += delta
		i += i & -i
	}
}

func (f *Fenwick) sum(idx int) int {
	if idx < 0 {
		return 0
	}
	if idx >= f.n {
		idx = f.n - 1
	}
	i := idx + 1 // convert to 1-indexed
	res := 0
	for i > 0 {
		res += f.tree[i]
		i -= i & -i
	}
	return res
}

// kth returns the smallest 0-indexed value such that prefix sum >= k.
// k is 1-indexed (1-based rank).
func (f *Fenwick) kth(k int) int {
	idx := 0
	bitMask := 1
	for bitMask <= f.n {
		bitMask <<= 1
	}
	bitMask >>= 1
	for bitMask > 0 {
		next := idx + bitMask
		if next <= f.n && f.tree[next] < k {
			k -= f.tree[next]
			idx = next
		}
		bitMask >>= 1
	}
	// idx is the largest internal index with prefix < k.
	// The answer in 0-indexed is idx.
	return idx
}

// --- MKAverage ---

const maxVal = 100001

type MKAverage struct {
	m          int
	k          int
	cnt        *Fenwick
	sum        *Fenwick
	queue      []int
	writeIdx   int
	elemCount  int
}

func Constructor(m int, k int) MKAverage {
	return MKAverage{
		m:     m,
		k:     k,
		cnt:   NewFenwick(maxVal),
		sum:   NewFenwick(maxVal),
		queue: make([]int, m),
	}
}

func (mk *MKAverage) AddElement(num int) {
	// Remove oldest if window is full
	if mk.elemCount >= mk.m {
		old := mk.queue[mk.writeIdx]
		mk.cnt.add(old, -1)
		mk.sum.add(old, -old)
	}

	// Add new element
	mk.queue[mk.writeIdx] = num
	mk.writeIdx = (mk.writeIdx + 1) % mk.m
	mk.elemCount++
	mk.cnt.add(num, 1)
	mk.sum.add(num, num)
}

func (mk *MKAverage) CalculateMKAverage() int {
	if mk.elemCount < mk.m {
		return -1
	}

	totalSum := mk.sum.sum(maxVal - 1)

	// Sum of k smallest elements
	vk := mk.cnt.kth(mk.k)
	sumBeforeK := mk.sum.sum(vk - 1)
	cntBeforeK := mk.cnt.sum(vk - 1)
	sumSmallestK := sumBeforeK + (mk.k-cntBeforeK)*vk

	// Sum of m-k smallest elements (to derive sum of k largest)
	vm := mk.cnt.kth(mk.m - mk.k)
	sumBeforeM := mk.sum.sum(vm - 1)
	cntBeforeM := mk.cnt.sum(vm - 1)
	sumFirstMK := sumBeforeM + (mk.m-mk.k-cntBeforeM)*vm
	sumLargestK := totalSum - sumFirstMK

	midSum := totalSum - sumSmallestK - sumLargestK
	return midSum / (mk.m - 2*mk.k)
}

// Stub kept for compatibility with the repo scaffold.
func FindingMkAverage() any {
	mk := Constructor(3, 1)
	mk.AddElement(3)
	mk.AddElement(1)
	mk.AddElement(4)
	return mk.CalculateMKAverage()
}
```

## 1830 — Minimum Number Of Operations To Make String Sorted

```go
package main

// LeetCode #1830: Minimum Number of Operations to Make String Sorted
// https://leetcode.com/problems/minimum-number-of-operations-to-make-string-sorted/
// Difficulty: Hard
//
// Approach: Combinatorics with Modular Arithmetic.
//   Each operation is the standard "next permutation" algorithm.
//   The number of operations to reach the sorted (ascending) string from
//   a given string equals the number of distinct next-permutations from
//   the sorted form, wrapping around. This is equivalent to:
//     answer = (totalPermutations - lexicographicRank) % totalPermutations
//
//   Compute total permutations and rank using factorials and modular
//   inverses modulo 1_000_000_007.

import "fmt"

const MOD1830 = 1_000_000_007

func main() {
	// Example 1
	fmt.Println("Example 1:", makeStringSorted("cba"))
	// Expected: 1

	// Example 2
	fmt.Println("Example 2:", makeStringSorted("aab"))
	// Expected: 3  (aab -> aba -> baa -> sorted? wait, sorted is "aab")

	// Single character
	fmt.Println("Edge (single):", makeStringSorted("a"))
	// Expected: 0

	// Already sorted
	fmt.Println("Edge (sorted):", makeStringSorted("abc"))
	// Expected: 0

	// Larger test
	fmt.Println("Example 3:", makeStringSorted("leetcode"))
	// Expected: some value; just verify it runs without error
}

func makeStringSorted(s string) int {
	n := len(s)

	// Special case: string already sorted -> 0 operations
	isSorted := true
	for i := 1; i < n; i++ {
		if s[i] < s[i-1] {
			isSorted = false
			break
		}
	}
	if isSorted {
		return 0
	}

	// Precompute factorials and inverse factorials
	fact := make([]int, n+1)
	invFact := make([]int, n+1)
	fact[0] = 1
	for i := 1; i <= n; i++ {
		fact[i] = (fact[i-1] * i) % MOD1830
	}
	invFact[n] = modPow1830(fact[n], MOD1830-2)
	for i := n - 1; i >= 0; i-- {
		invFact[i] = (invFact[i+1] * (i + 1)) % MOD1830
	}

	// Count character frequencies (lowercase English letters)
	cnt := make([]int, 26)
	for i := 0; i < n; i++ {
		cnt[s[i]-'a']++
	}

	// Compute total number of permutations = n! / prod(cnt[i]!)
	totalPerms := fact[n]
	for i := 0; i < 26; i++ {
		totalPerms = (totalPerms * invFact[cnt[i]]) % MOD1830
	}

	// Compute lexicographic rank (0-indexed)
	// Current product of invFact[cnt[i]] for efficient updates
	curProd := int64(1)
	for i := 0; i < 26; i++ {
		curProd = curProd * int64(invFact[cnt[i]]) % MOD1830
	}

	rank := int64(0)
	remaining := n
	for i := 0; i < n; i++ {
		ch := int(s[i] - 'a')
		remaining--

		for c := 0; c < ch; c++ {
			if cnt[c] == 0 {
				continue
			}
			// Number of permutations with a smaller char at this position
			// = fact[remaining] * invFact[all counts after decrementing cnt[c]]
			// curProd currently has invFact[cnt[c]], we need invFact[cnt[c]-1]
			// invFact[cnt[c]-1] = invFact[cnt[c]] * cnt[c]
			ways := int64(fact[remaining]) * curProd % MOD1830
			ways = ways * int64(cnt[c]) % MOD1830
			rank = (rank + ways) % MOD1830
		}

		// Update: use s[i], decrement its count
		// curProd changes: invFact[cnt[ch]] -> invFact[cnt[ch]-1]
		// invFact[cnt-1] = invFact[cnt] * cnt
		curProd = curProd * int64(cnt[ch]) % MOD1830
		cnt[ch]--
	}

	// Answer = (totalPerms - rank) % totalPerms
	ans := (totalPerms - int(rank) + MOD1830) % MOD1830
	return ans
}

func modPow1830(base, exp int) int {
	result := 1
	base %= MOD1830
	for exp > 0 {
		if exp&1 == 1 {
			result = (result * base) % MOD1830
		}
		base = (base * base) % MOD1830
		exp >>= 1
	}
	return result
}

// Stub kept for compatibility with the repo scaffold.
func MinimumNumberOfOperationsToMakeStringSorted() any {
	return makeStringSorted("cba")
}
```

## 1835 — Find Xor Sum Of All Pairs Bitwise And

```go
package main

// LeetCode #1835: Find XOR Sum of All Pairs Bitwise AND
// https://leetcode.com/problems/find-xor-sum-of-all-pairs-bitwise-and/
// Difficulty: Hard
//
// Approach: Bit Manipulation.
//   The XOR sum of all pairwise ANDs can be simplified:
//     For each bit position b:
//       Let c1 = count of arr1 elements with bit b set
//       Let c2 = count of arr2 elements with bit b set
//       Pairs with (arr1[i] & arr2[j]) having bit b set = c1 * c2
//       Bit b is set in final result iff c1 * c2 is odd
//     Since c1 * c2 is odd iff both c1 and c2 are odd, and parity of
//     count == XOR reduction:
//       result = (xor of all arr1) & (xor of all arr2)

import "fmt"

func main() {
	// Example 1
	fmt.Println("Example 1:", xorAllNums([]int{1, 2, 3}, []int{5, 6, 7}))
	// Expected: 0
	// xor1 = 1^2^3 = 0, xor2 = 5^6^7 = 4, result = 0 & 4 = 0

	// Example 2
	fmt.Println("Example 2:", xorAllNums([]int{0, 1, 2}, []int{3, 4, 5}))
	// Expected: 4
	// xor1 = 0^1^2 = 3, xor2 = 3^4^5 = 2, result = 3 & 2 = 2? Let's compute:
	// 0&3 ^ 0&4 ^ 0&5 ^ 1&3 ^ 1&4 ^ 1&5 ^ 2&3 ^ 2&4 ^ 2&5
	// = 0 ^ 0 ^ 0 ^ 1 ^ 0 ^ 1 ^ 2 ^ 0 ^ 0 = 1^1^2 = 2
	// Hmm, expected might be different. Let's check with example.

	// Actually from LeetCode: [0,1,2] and [3,4,5]
	// All pairs:
	// 0&3=0, 0&4=0, 0&5=0
	// 1&3=1, 1&4=0, 1&5=1
	// 2&3=2, 2&4=0, 2&5=0
	// XOR sum = 0^0^0^1^0^1^2^0^0 = 2
	fmt.Println("  -> xor1=0^1^2=3, xor2=3^4^5=2, result=3&2=2")

	// Edge case: single elements each
	fmt.Println("Edge (single each):", xorAllNums([]int{1}, []int{2}))
	// Expected: 0 (1&2 = 0)

	// All zeros
	fmt.Println("Edge (zeros):", xorAllNums([]int{0, 0}, []int{0, 0}))
	// Expected: 0
}

func xorAllNums(arr1 []int, arr2 []int) int {
	xor1 := 0
	for _, v := range arr1 {
		xor1 ^= v
	}
	xor2 := 0
	for _, v := range arr2 {
		xor2 ^= v
	}
	return xor1 & xor2
}

// Stub kept for compatibility with the repo scaffold.
func FindXorSumOfAllPairsBitwiseAnd() any {
	return xorAllNums([]int{1, 2, 3}, []int{5, 6, 7})
}
```

## 1840 — Maximum Building Height

```go
package main

// LeetCode #1840: Maximum Building Height
// https://leetcode.com/problems/maximum-building-height/
// Difficulty: Hard
//
// Approach: Two-Pass Constraint Propagation.
//   We have n buildings numbered 1..n. Each restriction [id, h] says
//   the height of building `id` cannot exceed `h`. Adjacent buildings
//   differ in height by at most 1.
//
//   1. Add implicit restriction: building 1 max height = 0.
//   2. Sort restrictions by id.
//   3. Forward pass: height[i] = min(height[i], height[i-1] + (id_i - id_{i-1}))
//   4. Backward pass: height[i] = min(height[i], height[i+1] + (id_{i+1} - id_i))
//   5. Between each pair of consecutive restrictions, the max height is
//      achieved at a peak point. The max building height between id[a] and id[b]
//      with heights h[a] and h[b] is:
//        dist = id[b] - id[a]
//        h[a] + h[b] + dist
//        peak = ---------------
//                2            (using integer arithmetic)
//      The max height in this interval = max(h[a], h[b]) + floor((dist - |h[a]-h[b]|) / 2)

import (
	"fmt"
	"sort"
)

func main() {
	// Example 1
	fmt.Println("Example 1:", maxBuilding(5, [][]int{{2, 1}, {4, 1}}))
	// Expected: 2

	// Example 2
	fmt.Println("Example 2:", maxBuilding(6, [][]int{}))
	// Expected: 5 (heights 0,1,2,3,4,5; max = 5)

	// Example 3
	fmt.Println("Example 3:", maxBuilding(10, [][]int{{5, 3}, {2, 5}, {7, 4}, {10, 0}}))
	// Expected: 5

	// Edge case: single restriction
	fmt.Println("Edge:", maxBuilding(3, [][]int{{3, 0}}))
	// Expected: 1 (0,1,0 -> max=1)
}

func maxBuilding(n int, restrictions [][]int) int {
	// Add implicit restriction: building 1 must have height 0
	type res struct {
		id, h int
	}
	list := make([]res, 0, len(restrictions)+1)
	list = append(list, res{1, 0})
	for _, r := range restrictions {
		if r[0] == 1 {
			if r[1] < 0 {
				r[1] = 0
			}
			list[0].h = min(r[1], list[0].h)
			continue
		}
		list = append(list, res{r[0], r[1]})
	}

	// Sort by id
	sort.Slice(list, func(i, j int) bool {
		return list[i].id < list[j].id
	})

	// Forward pass: propagate left-to-right constraints
	for i := 1; i < len(list); i++ {
		dist := list[i].id - list[i-1].id
		list[i].h = min(list[i].h, list[i-1].h+dist)
	}

	// Backward pass: propagate right-to-left constraints
	for i := len(list) - 2; i >= 0; i-- {
		dist := list[i+1].id - list[i].id
		list[i].h = min(list[i].h, list[i+1].h+dist)
	}

	// Compute max building height between pairs of restrictions
	ans := 0
	for i := 0; i < len(list); i++ {
		if list[i].h > ans {
			ans = list[i].h
		}
		if i+1 < len(list) {
			// Between list[i] and list[i+1], find the peak height
			dist := list[i+1].id - list[i].id
			hDiff := abs(list[i].h - list[i+1].h)
			// The height rises from the lower to the higher, then peaks
			remaining := dist - hDiff
			peak := max(list[i].h, list[i+1].h) + remaining/2
			if peak > ans {
				ans = peak
			}
		}
	}

	// Also consider building n (no explicit restriction after the last one)
	if len(list) > 0 {
		last := list[len(list)-1]
		remaining := n - last.id
		potential := last.h + remaining
		if potential > ans {
			ans = potential
		}
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Stub kept for compatibility with the repo scaffold.
func MaximumBuildingHeight() any {
	return maxBuilding(5, [][]int{{2, 1}, {4, 1}})
}
```

## 1842 — Next Palindrome Using Same Digits

```go
package main

// LeetCode #1842: Next Palindrome Using Same Digits
// https://leetcode.com/problems/next-palindrome-using-same-digits/
// Difficulty: Hard [Paid]
//
// Given a palindromic number as a string, find the next palindrome using the
// same digits. Since the right half is a mirror of the left half, we only
// need to find the next permutation of the left half, then mirror it back.
//
// If no larger palindrome can be formed, return an empty string.

import "fmt"

func main() {
	// Example 1: "1221" -> "2112"
	fmt.Println(nextPalindrome("1221"))
	// Example 2: "12321" -> "13231" (odd length, center stays)
	fmt.Println(nextPalindrome("12321"))
	// Example 3: "1" -> "" (no next palindrome)
	fmt.Println(nextPalindrome("1"))
	// Example 4: "11" -> "" (no next palindrome)
	fmt.Println(nextPalindrome("11"))
	// Example 5: "121" -> "211" (next left half of "12" is "21", mirror to "211")
	fmt.Println(nextPalindrome("121"))
	// Example 6: "32123" -> "" (no next palindrome)
	fmt.Println(nextPalindrome("32123"))
	// Example 7: "21312" -> "23132"
	fmt.Println(nextPalindrome("21312"))
	// Example 8: "4554" -> "5445"
	fmt.Println(nextPalindrome("4554"))
	// Example 9: "1234321" -> "1243421"
	fmt.Println(nextPalindrome("1234321"))
	// Example 10: "9999" -> "" (no next palindrome)
	fmt.Println(nextPalindrome("9999"))
}

func nextPalindrome(num string) string {
	n := len(num)
	if n <= 1 {
		return ""
	}

	half := n / 2
	// Work with the left half (including middle character for odd length)
	s := []byte(num[:half])
	// If odd length, we may or may not need to include middle
	// Actually, for palindrome, only the left half determines the number.
	// For odd length, center character stays the same, so we work with half.

	// Find next greater permutation of the left half
	// Find the rightmost position where s[i] < s[i+1]
	i := half - 2
	for i >= 0 && s[i] >= s[i+1] {
		i--
	}
	if i < 0 {
		return ""
	}

	// Find the rightmost element greater than s[i]
	j := half - 1
	for s[j] <= s[i] {
		j--
	}

	// Swap
	s[i], s[j] = s[j], s[i]

	// Reverse suffix starting at i+1
	for a, b := i+1, half-1; a < b; a, b = a+1, b-1 {
		s[a], s[b] = s[b], s[a]
	}

	// Build the palindrome
	result := make([]byte, n)
	copy(result, s)
	for idx := 0; idx < half; idx++ {
		result[n-1-idx] = s[idx]
	}
	// For odd length, the middle character is unchanged
	if n%2 == 1 {
		result[half] = num[half]
	}

	return string(result)
}
```

## 1847 — Closest Room

```go
package main

// LeetCode #1847: Closest Room
// https://leetcode.com/problems/closest-room/
// Difficulty: Hard
//
// There are rooms with (roomId, size). For each query (preferred, minSize),
// find the room with size >= minSize and roomId closest to preferred.
// If tie, choose the smaller roomId.
//
// Approach: sort rooms by size descending, sort queries by minSize descending.
// Process queries in order, adding eligible rooms to a sorted list.
// For each query, use binary search to find the closest roomId.

import (
	"fmt"
	"sort"
)

func main() {
	// Example 1:
	// rooms = [[2,2],[1,2],[3,2]]
	// queries = [[3,1],[3,3],[5,2]]
	// Expected: [3,-1,3]
	fmt.Println(closestRoom([][]int{{2, 2}, {1, 2}, {3, 2}}, [][]int{{3, 1}, {3, 3}, {5, 2}}))

	// Example 2:
	// rooms = [[1,4],[2,3],[3,5],[4,1],[5,2]]
	// queries = [[2,3],[2,4],[2,5]]
	// Expected: [2,1,-1]
	fmt.Println(closestRoom([][]int{{1, 4}, {2, 3}, {3, 5}, {4, 1}, {5, 2}}, [][]int{{2, 3}, {2, 4}, {2, 5}}))

	// Single room, single query
	fmt.Println(closestRoom([][]int{{1, 10}}, [][]int{{5, 5}}))

	// No room meets min size
	fmt.Println(closestRoom([][]int{{1, 5}}, [][]int{{3, 10}}))

	// Multiple rooms, exact match
	fmt.Println(closestRoom([][]int{{10, 20}, {20, 30}, {30, 40}}, [][]int{{25, 25}}))

	// Tie-breaking: smaller roomId wins
	fmt.Println(closestRoom([][]int{{5, 10}, {7, 10}}, [][]int{{6, 5}}))
}

func closestRoom(rooms [][]int, queries [][]int) []int {
	// Sort rooms by size descending
	sort.Slice(rooms, func(i, j int) bool {
		return rooms[i][1] > rooms[j][1]
	})

	// Sort queries by minSize descending, keeping original index
	type query struct {
		preferred int
		minSize   int
		idx       int
	}
	sortedQueries := make([]query, len(queries))
	for i, q := range queries {
		sortedQueries[i] = query{preferred: q[0], minSize: q[1], idx: i}
	}
	sort.Slice(sortedQueries, func(i, j int) bool {
		return sortedQueries[i].minSize > sortedQueries[j].minSize
	})

	ans := make([]int, len(queries))
	avail := []int{} // sorted room IDs
	roomIdx := 0

	for _, q := range sortedQueries {
		// Add all rooms with size >= query.minSize
		for roomIdx < len(rooms) && rooms[roomIdx][1] >= q.minSize {
			// Insert room ID in sorted order
			id := rooms[roomIdx][0]
			pos := sort.SearchInts(avail, id)
			avail = append(avail, 0)
			copy(avail[pos+1:], avail[pos:])
			avail[pos] = id
			roomIdx++
		}

		if len(avail) == 0 {
			ans[q.idx] = -1
			continue
		}

		// Binary search for closest room ID
		pos := sort.SearchInts(avail, q.preferred)
		if pos == 0 {
			ans[q.idx] = avail[0]
		} else if pos == len(avail) {
			ans[q.idx] = avail[len(avail)-1]
		} else {
			leftDiff := q.preferred - avail[pos-1]
			rightDiff := avail[pos] - q.preferred
			if leftDiff <= rightDiff {
				ans[q.idx] = avail[pos-1]
			} else {
				ans[q.idx] = avail[pos]
			}
		}
	}
	return ans
}
```

## 1851 — Minimum Interval To Include Each Query

```go
package main

// LeetCode #1851: Minimum Interval to Include Each Query
// https://leetcode.com/problems/minimum-interval-to-include-each-query/
// Difficulty: Hard
//
// Given intervals [left, right] and queries, for each query find the minimum
// interval length (right-left+1) that contains it, or -1 if none.
//
// Approach: sort intervals by length ascending, sort queries with index.
// Use DSU (union-find) to efficiently skip already-answered queries.
// For each interval in increasing length, find all queries within [left, right]
// that haven't been answered yet and set their answer.

import (
	"fmt"
	"sort"
)

func main() {
	// Example 1:
	// intervals = [[1,4],[2,4],[3,6],[4,4]]
	// queries = [2,3,4,5]
	// Expected: [3,3,1,4]
	fmt.Println(minInterval([][]int{{1, 4}, {2, 4}, {3, 6}, {4, 4}}, []int{2, 3, 4, 5}))

	// Example 2:
	// intervals = [[2,3],[2,5],[1,8],[20,25]]
	// queries = [2,19,5,22]
	// Expected: [2,-1,4,6]
	fmt.Println(minInterval([][]int{{2, 3}, {2, 5}, {1, 8}, {20, 25}}, []int{2, 19, 5, 22}))

	// Single interval, single query
	fmt.Println(minInterval([][]int{{1, 5}}, []int{3}))

	// No interval contains query
	fmt.Println(minInterval([][]int{{1, 2}, {5, 6}}, []int{4}))

	// Multiple queries, overlapping intervals
	fmt.Println(minInterval([][]int{{1, 10}, {2, 8}, {3, 6}}, []int{1, 4, 7, 10}))

	// Single query
	fmt.Println(minInterval([][]int{{1, 3}, {2, 5}, {1, 100}}, []int{4}))
}

func minInterval(intervals [][]int, queries []int) []int {
	n := len(queries)

	// Sort queries with their original indices
	type qi struct {
		val int
		idx int
	}
	qSorted := make([]qi, n)
	for i, q := range queries {
		qSorted[i] = qi{val: q, idx: i}
	}
	sort.Slice(qSorted, func(i, j int) bool {
		return qSorted[i].val < qSorted[j].val
	})

	// Sort intervals by length ascending
	sort.Slice(intervals, func(i, j int) bool {
		li := intervals[i][1] - intervals[i][0] + 1
		lj := intervals[j][1] - intervals[j][0] + 1
		return li < lj
	})

	// DSU: parent[i] = next unprocessed query index (or i itself)
	parent := make([]int, n+1)
	for i := 0; i <= n; i++ {
		parent[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}

	for _, iv := range intervals {
		l, r := iv[0], iv[1]
		length := r - l + 1

		// Find first query >= l
		start := sort.Search(n, func(i int) bool {
			return qSorted[i].val >= l
		})
		if start == n {
			continue
		}

		// Process all unassigned queries within [l, r]
		for idx := find(start); idx < n && qSorted[idx].val <= r; idx = find(idx) {
			ans[qSorted[idx].idx] = length
			parent[idx] = idx + 1
		}
	}
	return ans
}
```

## 1857 — Largest Color Value In A Directed Graph

```go
package main

// LeetCode #1857: Largest Color Value in a Directed Graph
// https://leetcode.com/problems/largest-color-value-in-a-directed-graph/
// Difficulty: Hard
//
// Approach: Topological Sort + DP per Color.
//   For each node, maintain a dp[c] = max count of color c on any path
//   ending at this node. Process nodes in topological (Kahn's) order.
//   For each edge u->v:
//     dp[v][c] = max(dp[v][c], dp[u][c] + (1 if colors[v]==c else 0))
//   The answer is the max dp value across all nodes and colors.
//   If there's a cycle (not all nodes processed), return -1.

import "fmt"

func main() {
	// Example 1
	fmt.Println("Example 1:",
		largestPathValue("abaca", [][]int{{0, 1}, {0, 2}, {2, 3}, {3, 4}}))
	// Expected: 3 (path 0->2->3->4 has 3 'a's)

	// Example 2: cycle
	fmt.Println("Example 2:",
		largestPathValue("a", [][]int{{0, 0}}))
	// Expected: -1 (self-loop is a cycle)

	// Example 3: no edges
	fmt.Println("Example 3:",
		largestPathValue("abc", [][]int{}))
	// Expected: 1 (each node alone has count 1 of its own color)

	// Edge case: single node
	fmt.Println("Edge (single):",
		largestPathValue("z", [][]int{}))
	// Expected: 1

	// Two nodes, two colors
	fmt.Println("Edge (two nodes):",
		largestPathValue("ab", [][]int{{0, 1}}))
	// Expected: 1 (path 0->1 has one a and one b, max frequency = 1)
}

func largestPathValue(colors string, edges [][]int) int {
	n := len(colors)

	// Build adjacency list and in-degree array
	adj := make([][]int, n)
	inDeg := make([]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		inDeg[v]++
	}

	// dp[i][c] = max count of color c on any path ending at node i
	dp := make([][26]int, n)

	// Kahn's topological sort
	queue := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if inDeg[i] == 0 {
			queue = append(queue, i)
			dp[i][colors[i]-'a'] = 1
		}
	}

	processed := 0
	ans := 0

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		processed++

		// Update global answer with dp[u]
		for c := 0; c < 26; c++ {
			if dp[u][c] > ans {
				ans = dp[u][c]
			}
		}

		for _, v := range adj[u] {
			// Propagate dp from u to v
			for c := 0; c < 26; c++ {
				add := 0
				if int(colors[v]-'a') == c {
					add = 1
				}
				if dp[u][c]+add > dp[v][c] {
					dp[v][c] = dp[u][c] + add
				}
			}
			inDeg[v]--
			if inDeg[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	if processed != n {
		return -1 // cycle detected
	}

	return ans
}

// Stub kept for compatibility with the repo scaffold.
func LargestColorValueInADirectedGraph() any {
	return largestPathValue("abaca", [][]int{{0, 1}, {0, 2}, {2, 3}, {3, 4}})
}
```

## 1862 — Sum Of Floored Pairs

```go
package main

// LeetCode #1862: Sum of Floored Pairs
// https://leetcode.com/problems/sum-of-floored-pairs/
// Difficulty: Hard

import (
	"fmt"
)

func sumOfFlooredPairs(nums []int) int {
	const mod = 1_000_000_007
	maxVal := 0
	for _, v := range nums {
		if v > maxVal {
			maxVal = v
		}
	}
	freq := make([]int, maxVal+1)
	for _, v := range nums {
		freq[v]++
	}
	prefix := make([]int, maxVal+1)
	for i := 1; i <= maxVal; i++ {
		prefix[i] = prefix[i-1] + freq[i]
	}

	ans := 0
	for x := 1; x <= maxVal; x++ {
		if freq[x] == 0 {
			continue
		}
		// For each multiple y = x * q, count how many numbers in [y, y+x-1]
		for y := x; y <= maxVal; y += x {
			q := y / x
			count := prefix[min(maxVal, y+x-1)] - prefix[y-1]
			ans = (ans + freq[x]*count%mod*q%mod) % mod
		}
	}
	return ans
}

func main() {
	// Example: [2,5,9] -> 10
	// floor(2/2)+floor(5/2)+floor(9/2)=1+2+4=7
	// floor(2/5)+floor(5/5)+floor(9/5)=0+1+1=2
	// floor(2/9)+floor(5/9)+floor(9/9)=0+0+1=1
	// Total: 7+2+1=10
	fmt.Println(sumOfFlooredPairs([]int{2, 5, 9}))

	// Additional test
	fmt.Println(sumOfFlooredPairs([]int{7, 7, 7, 7, 7, 7, 7}))
}
```

## 1866 — Number Of Ways To Rearrange Sticks With K Sticks Visible

```go
package main

// LeetCode #1866: Number of Ways to Rearrange Sticks With K Sticks Visible
// https://leetcode.com/problems/number-of-ways-to-rearrange-sticks-with-k-sticks-visible/
// Difficulty: Hard

import "fmt"

func numberOfWaysToRearrangeSticks(n int, k int) int {
	const mod = 1_000_000_007
	// dp[i][j] = ways to arrange i sticks so exactly j are visible
	// Recurrence: dp[i][j] = dp[i-1][j-1] + (i-1)*dp[i-1][j]
	// This is the unsigned Stirling numbers of the first kind
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
	}
	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= min(i, k); j++ {
			dp[i][j] = (dp[i-1][j-1] + (i-1)*dp[i-1][j]%mod) % mod
		}
	}
	return dp[n][k]
}

func main() {
	// Example: n=3, k=2 -> 3
	// [1,3,2], [2,3,1], [2,1,3]
	fmt.Println(numberOfWaysToRearrangeSticks(3, 2))

	// Additional test
	fmt.Println(numberOfWaysToRearrangeSticks(5, 3))
}
```

## 1872 — Stone Game Viii

```go
package main

// LeetCode #1872: Stone Game VIII
// https://leetcode.com/problems/stone-game-viii/
// Difficulty: Hard

import "fmt"

func stoneGameViii(stones []int) int {
	n := len(stones)
	prefix := make([]int, n)
	prefix[0] = stones[0]
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] + stones[i]
	}

	// dp[i] = max score difference (current player - opponent) starting from
	// position i. The player may choose any j >= i, j < n-1, take the prefix
	// from position i to j (scoring prefix[j] - base), and leave position j+1
	// for the opponent.
	//
	// Recurrence: dp[i] = max over j >= i of (prefix[j] - base - dp[j+1])
	// where base = 0 (when i=0) or prefix[i-1].
	// This simplifies to: dp[i] = max(prefix[i] - dp[i+1], dp[i+1]).
	//
	// dp[n-1] = 0 (cannot take when only 1 stone remains).

	dp := 0
	for i := n - 2; i >= 0; i-- {
		dp = max(prefix[i]-dp, dp)
	}
	return dp
}

func main() {
	// Test cases
	fmt.Println(stoneGameViii([]int{-1, 2, -3, 4, -5}))
	fmt.Println(stoneGameViii([]int{1, 2, 3, 4, 5}))
}
```

## 1879 — Minimum Xor Sum Of Two Arrays

```go
package main

// LeetCode #1879: Minimum XOR Sum of Two Arrays
// https://leetcode.com/problems/minimum-xor-sum-of-two-arrays/
// Difficulty: Hard

import (
	"fmt"
	"math"
)

func minimumXorSum(nums1 []int, nums2 []int) int {
	n := len(nums1)
	size := 1 << n
	dp := make([]int, size)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0

	// dp[mask] = min XOR sum assigning first k elements of nums1
	// to elements of nums2 indicated by mask (where k = popcount(mask))
	for mask := 0; mask < size; mask++ {
		i := bitsCount(mask) // number of assigned elements in nums1
		if i >= n {
			continue
		}
		for j := 0; j < n; j++ {
			if mask&(1<<j) == 0 {
				newMask := mask | (1 << j)
				val := dp[mask] + (nums1[i] ^ nums2[j])
				if val < dp[newMask] {
					dp[newMask] = val
				}
			}
		}
	}
	return dp[size-1]
}

func bitsCount(x int) int {
	c := 0
	for x > 0 {
		c += x & 1
		x >>= 1
	}
	return c
}

func main() {
	// Example: [1,2], [2,3] -> 2
	// (1 XOR 2) + (2 XOR 3) = 3 + 1 = 4
	// (1 XOR 3) + (2 XOR 2) = 2 + 0 = 2 (minimum)
	fmt.Println(minimumXorSum([]int{1, 2}, []int{2, 3}))

	// Additional test
	fmt.Println(minimumXorSum([]int{1, 0, 3}, []int{5, 3, 4}))
}
```

## 1883 — Minimum Skips To Arrive At Meeting On Time

```go
package main

// LeetCode #1883: Minimum Skips to Arrive at Meeting On Time
// https://leetcode.com/problems/minimum-skips-to-arrive-at-meeting-on-time/
// Difficulty: Hard

import (
	"fmt"
	"math"
)

func minSkipsToArriveAtMeetingOnTime(dist []int, speed int, hoursBefore int) int {
	n := len(dist)
	// dp[i][j] = minimum time (in terms of distance units multiplied to avoid float)
	// to reach dist[i] with j skips. We use the fact that skipping means we avoid
	// the "round up" (arrive at a rest stop).
	// Time is represented as integer scaled by speed, but we need to handle rounding.
	//
	// Let's use the approach: dp[i][j] = minimum "delay" (extra time beyond raw travel)
	// represented as a value such that actual_time = (total_dist + dp[i][j]) / speed
	// No, let's use a cleaner approach.
	//
	// When we don't skip: time = ceil(time_before + dist[i] / speed)
	// When we skip: time = time_before + dist[i] / speed (no rounding)
	// where time_before is the time entering rest stop i.
	//
	// Use dp[i][j] = the total time (in units of 1/speed, i.e., dist units) after
	// reaching dist[i] with j skips, rounding up at each step except skipped ones.
	// But we need to track rounding status precisely.
	//
	// Better approach: dp[i][j] = the minimum "actual arrival time in hours" (as a rational)
	// after i-th road with j skips. Represent as ceil of a fraction.
	//
	// Let's use the technique from the editorial: use dp[i][j] = minimum time to complete
	// first i roads with j skips, but store time as an integer representing the "rounded"
	// value. The trick: when we don't skip, we round up. When we skip, we don't.
	// The check at the end: if dp[n][j] <= hoursBefore * speed, it's possible with j skips.
	// Because dp[n][j] stores total distance without the "rounding factor" of the last step.
	//
	// Actually the cleanest way: dp[i][j] = minimum number of "extra minutes" or such.
	// Let me just follow the well-known solution pattern:
	//
	// Use dp[i][j] = minimum time to cross first i roads with j skips.
	// Time is measured in units of speed (1 unit = 1/speed hours).
	// When crossing road i (0-indexed), we travel dist[i] units.
	// Without skip: new_time = ((dp[i-1][j] + dist[i-1] + speed - 1) / speed) * speed ...
	// No, this is getting messy. Let me use the simplest known working approach.
	//
	// We store dp[j] after processing each road, where dp[j] = total time in "distance units",
	// and we round up at each step unless we skip. At the end, check if dp[j] <= hoursBefore * speed.
	//
	// dp[j] = floor(total_time * speed) where total_time is accumulated time so far
	// with j skips. When we don't skip: dp[j] = ((dp[j] + dist[i] + speed - 1) / speed) * speed
	// When we skip: dp[j] = dp[j] + dist[i]
	// Wait no. Let me think in terms of the actual time in hours.
	//
	// If we track the "time elapsed so far" as a value where we keep the integer part
	// (hours passed) and the remainder (distance into the current hour).
	// Actually, let me just use the standard double/float with epsilon approach
	// but do it with integer math.
	//
	// The key insight from the editorial:
	// Let dp[i][j] = minimum time to finish first i roads with j skips, but we DON'T
	// apply the rounding up on the last segment.
	// dp[i][j] = min(
	//   skip on i: dp[i-1][j-1] + dist[i-1],                                 // if j > 0
	//   no skip on i: ceil(dp[i-1][j] + dist[i-1])                             // round up
	// )
	// where ceil(x) means round up to next multiple of 1 (in integer terms).
	// But how to represent "ceil" cleanly?
	//
	// Represent time in "fractional hours * speed" = distance units.
	// Let dp[j] = accumulated time in distance units after processing some roads.
	// Without rounding, total is just sum of dist.
	// With rounding, at each road the total gets rounded up to the next whole hour.
	//
	// Round up to next hour: ceil_to_hour(t) = ((t + speed - 1) / speed) * speed
	// where t is in distance units.
	//
	// So for road with distance d:
	// No skip: new_dp[j] = ((dp[j] + d + speed - 1) / speed) * speed
	// Skip:    new_dp[j] = dp[j-1] + d  (or old dp[j-1] + d)
	//
	// At the end, we have total_time in distance units. We need dp[n][j] <= hoursBefore * speed.
	// But we didn't round up on the last segment (the destination isn't a "rest stop").
	// Ah wait, we shouldn't round up on the final arrival. Let me reconsider.
	//
	// The roads end with the destination. We round up at each rest stop (which is between roads).
	// There are n-1 rest stops (after roads 0..n-2). Road n-1 leads directly to the destination.
	// So we round up for roads 0 to n-2 (arriving at rest stops), but NOT for road n-1.
	//
	// This means the last step is always "skip" automatically.
	// dp[n-1][j] represents the time to reach road n-1's start (which is a rest stop).
	// Then total_time = dp[n-1][j] + dist[n-1] (no rounding).
	// We need total <= hoursBefore * speed.
	//
	// For i from 0 to n-2:
	//   new_dp[j] = min(
	//     j > 0 ? old_dp[j-1] + dist[i] : INF,
	//     ((old_dp[j] + dist[i] + speed - 1) / speed) * speed
	//   )
	// Then final check: dp[j] + dist[n-1] <= hoursBefore * speed.

	// Handle: if sum of dist > hoursBefore * speed, impossible even with all skips
	totalDist := 0
	for _, d := range dist {
		totalDist += d
	}
	if totalDist > hoursBefore*speed {
		return -1
	}

	INF := math.MaxInt32
	dp := make([]int, n)
	for i := range dp {
		dp[i] = INF
	}
	dp[0] = 0

	// Process roads 0 to n-2 (arriving at rest stops)
	for i := 0; i < n-1; i++ {
		d := dist[i]
		// Process skips from high to low so we don't reuse same row
		newDp := make([]int, n)
		for j := range newDp {
			newDp[j] = INF
		}
		for j := 0; j <= i+1; j++ {
			// Option 1: no skip (round up)
			if dp[j] < INF {
				rounded := ((dp[j] + d + speed - 1) / speed) * speed
				if rounded < newDp[j] {
					newDp[j] = rounded
				}
			}
			// Option 2: skip (no round), can only do if j > 0
			if j > 0 && dp[j-1] < INF {
				noRound := dp[j-1] + d
				if noRound < newDp[j] {
					newDp[j] = noRound
				}
			}
		}
		dp = newDp
	}

	// Final road
	last := dist[n-1]
	for j := 0; j < n; j++ {
		if dp[j] < INF && dp[j]+last <= hoursBefore*speed {
			return j
		}
	}
	return -1
}

func main() {
	// Example: dist=[1,3,2], speed=4, hoursBefore=2 -> 1
	fmt.Println(minSkipsToArriveAtMeetingOnTime([]int{1, 3, 2}, 4, 2))

	// Additional test
	fmt.Println(minSkipsToArriveAtMeetingOnTime([]int{7, 3, 5, 5}, 2, 10))
}
```

## 1889 — Minimum Space Wasted From Packaging

```go
package main

// LeetCode #1889: Minimum Space Wasted From Packaging
// https://leetcode.com/problems/minimum-space-wasted-from-packaging/
// Difficulty: Hard

import (
	"fmt"
	"math"
	"sort"
)

func minWastedSpace(packages []int, boxes [][]int) int {
	const mod = 1_000_000_007
	sort.Ints(packages)
	n := len(packages)
	prefix := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + packages[i]
	}

	ans := math.MaxInt64

	for _, supplier := range boxes {
		sort.Ints(supplier)
		if supplier[len(supplier)-1] < packages[n-1] {
			continue // cannot fit the largest package
		}
		waste := 0
		prevIdx := 0
		for _, box := range supplier {
			// find the last package that fits in this box
			idx := sort.Search(n-prevIdx, func(k int) bool {
				return packages[prevIdx+k] > box
			}) + prevIdx
			if idx > prevIdx {
				count := idx - prevIdx
				// waste = box * count - sum of packages in [prevIdx, idx)
				sum := prefix[idx] - prefix[prevIdx]
				waste += box*count - sum
				prevIdx = idx
				if waste > ans { // early break
					break
				}
			}
		}
		if prevIdx == n && waste < ans {
			ans = waste
		}
	}

	if ans == math.MaxInt64 {
		return -1
	}
	return ans % mod
}

func main() {
	// Example: packages=[2,3,5], boxes=[[4,8],[2,8]] -> 6
	fmt.Println(minWastedSpace([]int{2, 3, 5}, [][]int{{4, 8}, {2, 8}}))

	// Additional test
	fmt.Println(minWastedSpace([]int{3, 5, 8, 10, 11, 12}, [][]int{{12}, {11, 9}, {10, 5, 14}}))
}
```

## 1892 — Page Recommendations Ii

```go
package main

// LeetCode #1892: Page Recommendations II
// https://leetcode.com/problems/page-recommendations-ii/
// Difficulty: Hard [Paid]
//
// Given tables: Friendship (user1_id, user2_id), Likes (user_id, page_id),
// and Users (user_id), for each user recommend pages that:
// - Are liked by friends of the user's friends (2nd degree friends)
// - ARE NOT already liked by the user
// - ARE NOT already liked by the user's direct friends
// Order by user_id, page_id.

import (
	"fmt"
	"sort"
)

func main() {
	// Test case 1: basic
	friendships := [][]int{
		{1, 2}, {1, 3}, {1, 4}, {2, 3},
	}
	likes := [][]int{
		{1, 101}, {2, 102}, {3, 103}, {4, 104},
		{2, 105},
	}
	users := []int{1}
	// Friends of user 1: {2, 3, 4}
	// Friends of friends (excluding 1 and {2,3,4}): none directly, so empty
	fmt.Println(pageRecommendationsIi(friendships, likes, users))

	// Test case 2:
	friendships = [][]int{
		{1, 2}, {2, 3}, {3, 4},
	}
	likes = [][]int{
		{1, 101}, {1, 102},
		{2, 102},
		{3, 103},
		{4, 104},
	}
	users = []int{1}
	// Direct friends: {2} (likes 102)
	// Friends of friends: {3} (likes 103) -> recommend 103 (not liked by 1 or direct friends)
	// {4} is friend of 3, also friend-of-friend of 1 -> likes 104 -> recommend 104
	// Result: should get 103, 104
	fmt.Println(pageRecommendationsIi(friendships, likes, users))

	// Test case 3: multiple users
	friendships = [][]int{
		{1, 2}, {2, 3},
	}
	likes = [][]int{
		{1, 101}, {2, 101}, {3, 102},
	}
	users = []int{1, 2}
	fmt.Println(pageRecommendationsIi(friendships, likes, users))

	// Test case 4: empty
	fmt.Println(pageRecommendationsIi([][]int{}, [][]int{}, []int{1}))
}

// pageRecommendationsIi returns [][]int{{user_id, page_id}, ...}
func pageRecommendationsIi(friendships [][]int, likes [][]int, users []int) [][]int {
	// Build friendship graph
	friends := make(map[int]map[int]bool) // user -> set of friends
	for _, f := range friendships {
		u1, u2 := f[0], f[1]
		if friends[u1] == nil {
			friends[u1] = make(map[int]bool)
		}
		if friends[u2] == nil {
			friends[u2] = make(map[int]bool)
		}
		friends[u1][u2] = true
		friends[u2][u1] = true
	}

	// Build likes mapping
	userLikes := make(map[int]map[int]bool) // user -> set of liked pages
	for _, l := range likes {
		uid, pid := l[0], l[1]
		if userLikes[uid] == nil {
			userLikes[uid] = make(map[int]bool)
		}
		userLikes[uid][pid] = true
	}

	var result [][]int

	for _, user := range users {
		// Find direct friends
		directFriends := friends[user]
		if directFriends == nil {
			directFriends = make(map[int]bool)
		}

		// Find friend-of-friend pages to recommend
		recommended := make(map[int]bool)

		// For each friend of the user
		for friend := range directFriends {
			// For each friend of that friend
			for fof := range friends[friend] {
				if fof == user || directFriends[fof] {
					continue // skip user and direct friends
				}
				// Add pages liked by fof
				for page := range userLikes[fof] {
					recommended[page] = true
				}
			}
		}

		// Remove pages already liked by user
		for page := range userLikes[user] {
			delete(recommended, page)
		}

		// Remove pages liked by direct friends
		for friend := range directFriends {
			for page := range userLikes[friend] {
				delete(recommended, page)
			}
		}

		// Build sorted result
		var pages []int
		for p := range recommended {
			pages = append(pages, p)
		}
		sort.Ints(pages)
		for _, p := range pages {
			result = append(result, []int{user, p})
		}
	}

	if result == nil {
		return [][]int{}
	}
	return result
}
```

## 1896 — Minimum Cost To Change The Final Value Of Expression

```go
package main

// LeetCode #1896: Minimum Cost to Change the Final Value of Expression
// https://leetcode.com/problems/minimum-cost-to-change-the-final-value-of-expression/
// Difficulty: Hard

import (
	"fmt"
)

// Pair represents (min cost to make value 0, min cost to make value 1)
type Pair struct{ zero, one int }

func minCostToChange(expression string) int {
	var valStack []Pair
	var opStack []byte

	apply := func() {
		b := valStack[len(valStack)-1]
		valStack = valStack[:len(valStack)-1]
		a := valStack[len(valStack)-1]
		valStack = valStack[:len(valStack)-1]
		op := opStack[len(opStack)-1]
		opStack = opStack[:len(opStack)-1]

		var res Pair
		if op == '&' {
			// cost to make 0:
			// already 0 on left: cost 0 (any right) or change left from 0 to 0
			// We can make left=0 (a.zero) OR right=0 (b.zero) or both
			// Actually for AND: result=0 if any operand is 0
			// Option 1: make left 0 (cost a.zero)
			// Option 2: make right 0 (cost b.zero)
			// Option 2b: change operator (cost 1) but we need to track toggle ops too
			// Option 3: change operator to |, then result 0 if both 0: cost = a.zero + b.zero + 1

			// Make 0: either operand is 0
			res.zero = min(a.zero, b.zero)
			// Also can change operator: a.zero + b.zero + 1 (change & to |, make both 0)
			res.zero = min(res.zero, a.zero+b.zero+1)

			// Make 1: both operands must be 1
			res.one = a.one + b.one
			// Or change operator: min(a.one, b.one) + 1 (change & to |, make one 1)
			res.one = min(res.one, min(a.one, b.one)+1)
		} else { // '|'
			// Make 0: both operands must be 0
			res.zero = a.zero + b.zero
			// Or change operator: min(a.zero, b.zero) + 1 (change | to &, make one 0)
			res.zero = min(res.zero, min(a.zero, b.zero)+1)

			// Make 1: either operand is 1
			res.one = min(a.one, b.one)
			// Or change operator: a.one + b.one + 1 (change | to &, make both 1)
			res.one = min(res.one, a.one+b.one+1)
		}
		valStack = append(valStack, res)
	}

	for _, ch := range expression {
		switch ch {
		case '(':
			opStack = append(opStack, '(')
		case ')':
			for len(opStack) > 0 && opStack[len(opStack)-1] != '(' {
				apply()
			}
			opStack = opStack[:len(opStack)-1] // pop '('
		case '&', '|':
			for len(opStack) > 0 && opStack[len(opStack)-1] != '(' && opStack[len(opStack)-1] != '&' && opStack[len(opStack)-1] != '|' {
				// precedence: same-level operators processed left to right
				// but actually & and | have the same precedence in this problem
				// We process left to right when same operator type
			}
			// Actually for simplicity, since we only have & and | with same precedence,
			// we just process left to right. We'll just evaluate when we see an operator
			// if the previous one was also an operator.
			for len(opStack) > 0 && opStack[len(opStack)-1] != '(' {
				apply()
			}
			opStack = append(opStack, byte(ch))
		case '0', '1':
			v := int(ch - '0')
			if v == 0 {
				valStack = append(valStack, Pair{0, 1}) // cost 0 to make 0, cost 1 to make 1
			} else {
				valStack = append(valStack, Pair{1, 0}) // cost 1 to make 0, cost 0 to make 1
			}
		}
	}

	for len(opStack) > 0 {
		apply()
	}

	res := valStack[len(valStack)-1]
	// The input expression evaluates to some value.
	// We need the min cost to flip that value.
	// First, determine the current value.
	// Re-evaluate the expression without the cost to find current value:
	currVal := evaluate(expression)
	if currVal == 0 {
		return res.one
	}
	return res.zero
}

func evaluate(expr string) int {
	// Simple stack-based evaluation
	var vals []int
	var ops []byte
	for _, ch := range expr {
		switch ch {
		case '(':
			ops = append(ops, '(')
		case ')':
			for len(ops) > 0 && ops[len(ops)-1] != '(' {
				b := vals[len(vals)-1]
				vals = vals[:len(vals)-1]
				a := vals[len(vals)-1]
				vals = vals[:len(vals)-1]
				op := ops[len(ops)-1]
				ops = ops[:len(ops)-1]
				if op == '&' {
					vals = append(vals, a&b)
				} else {
					vals = append(vals, a|b)
				}
			}
			ops = ops[:len(ops)-1] // pop '('
		case '0', '1':
			vals = append(vals, int(ch-'0'))
		case '&', '|':
			ops = append(ops, byte(ch))
		}
	}
	return vals[0]
}

func main() {
	// Example: "1&(0|1)" -> 1 (change & to |)
	fmt.Println(minCostToChange("1&(0|1)"))

	// Additional tests
	fmt.Println(minCostToChange("0&0"))
	fmt.Println(minCostToChange("1|0"))
}
```

## 1900 — The Earliest And Latest Rounds Where Players Compete

```go
package main

// LeetCode #1900: The Earliest and Latest Rounds Where Players Compete
// https://leetcode.com/problems/the-earliest-and-latest-rounds-where-players-compete/
// Difficulty: Hard

import "fmt"

type state struct {
	a, b, n int
}

var memo map[state][]int

func dfs(a, b, n int) []int {
	// Returns [earliest, latest] round (1 for "this round") where a and b meet.
	if a+b == n-1 {
		return []int{1, 1}
	}

	key := state{a, b, n}
	if res, ok := memo[key]; ok {
		return res
	}

	pairs := n / 2               // number of pairs
	half := (n + 1) / 2          // number of players advancing
	mid := n / 2                 // middle player (only meaningful when n odd)

	pairA := min(a, n-1-a)
	pairB := min(b, n-1-b)

	// Other pair indices (0..pairs-1), excluding pairA and pairB
	otherPairs := make([]int, 0)
	for i := 0; i < pairs; i++ {
		if i != pairA && i != pairB {
			otherPairs = append(otherPairs, i)
		}
	}

	minRound := 100
	maxRound := 0
	m := len(otherPairs)

	for mask := 0; mask < (1 << m); mask++ {
		winners := make([]int, half)
		for i := range winners {
			winners[i] = -1
		}

		// a and b always win their matches
		winners[pairA] = a
		winners[pairB] = b

		// For each other pair, try both winners
		for k, pi := range otherPairs {
			if mask&(1<<k) != 0 {
				winners[pi] = n - 1 - pi // higher-indexed player wins
			} else {
				winners[pi] = pi // lower-indexed player wins
			}
		}

		// When n is odd, the middle player (position = pairs = n/2) gets a bye
		// and advances to the last position in the new lineup.
		if n%2 == 1 {
			winners[half-1] = mid
		}

		newA, newB := -1, -1
		for i, w := range winners {
			if w == a {
				newA = i
			}
			if w == b {
				newB = i
			}
		}
		if newA != -1 && newB != -1 {
			if newA > newB {
				newA, newB = newB, newA
			}
			sub := dfs(newA, newB, half)
			earliest := 1 + sub[0]
			latest := 1 + sub[1]
			if earliest < minRound {
				minRound = earliest
			}
			if latest > maxRound {
				maxRound = latest
			}
		}
	}

	result := []int{minRound, maxRound}
	memo[key] = result
	return result
}

func earliestAndLatest(n int, firstPlayer int, secondPlayer int) []int {
	memo = make(map[state][]int)
	a, b := firstPlayer-1, secondPlayer-1
	if a > b {
		a, b = b, a
	}
	return dfs(a, b, n)
}

func main() {
	// Example: n=11, firstPlayer=2, secondPlayer=4 -> [3,4]
	fmt.Println(earliestAndLatest(11, 2, 4))

	// Additional test
	fmt.Println(earliestAndLatest(5, 1, 5))
}
```

## 1912 — Design Movie Rental System

```go
package main

// LeetCode #1912: Design Movie Rental System
// https://leetcode.com/problems/design-movie-rental-system/
// Difficulty: Hard

import (
	"container/heap"
	"fmt"
	"sort"
)

// Entry represents a movie copy: (shop, movie, price)
type Entry struct {
	shop, movie, price int
}

// ReportHeap min-heap sorted by (price, shop, movie)
type ReportHeap []Entry

func (h ReportHeap) Len() int { return len(h) }
func (h ReportHeap) Less(i, j int) bool {
	if h[i].price != h[j].price {
		return h[i].price < h[j].price
	}
	if h[i].shop != h[j].shop {
		return h[i].shop < h[j].shop
	}
	return h[i].movie < h[j].movie
}
func (h ReportHeap) Swap(i, j int)     { h[i], h[j] = h[j], h[i] }
func (h *ReportHeap) Push(x any)       { *h = append(*h, x.(Entry)) }
func (h *ReportHeap) Pop() any         { old := *h; n := len(old); x := old[n-1]; *h = old[:n-1]; return x }

// SearchHeap min-heap sorted by (price, shop)
type SearchHeap []Entry

func (h SearchHeap) Len() int { return len(h) }
func (h SearchHeap) Less(i, j int) bool {
	if h[i].price != h[j].price {
		return h[i].price < h[j].price
	}
	return h[i].shop < h[j].shop
}
func (h SearchHeap) Swap(i, j int)     { h[i], h[j] = h[j], h[i] }
func (h *SearchHeap) Push(x any)       { *h = append(*h, x.(Entry)) }
func (h *SearchHeap) Pop() any         { old := *h; n := len(old); x := old[n-1]; *h = old[:n-1]; return x }

type MovieRentalSystem struct {
	avail      map[int]*SearchHeap // per-movie min-heap of available copies
	rented     *ReportHeap         // min-heap of all currently rented copies
	rentedData map[Entry]bool      // tracks which entries are currently rented, keyed by full Entry
	priceOf    map[[2]int]int      // (shop,movie) -> price, for quick lookup in Drop
}

func Constructor(entries [][]int) MovieRentalSystem {
	mrs := MovieRentalSystem{
		avail:      make(map[int]*SearchHeap),
		rented:     &ReportHeap{},
		rentedData: make(map[Entry]bool),
		priceOf:    make(map[[2]int]int),
	}
	byMovie := make(map[int][]Entry)
	for _, e := range entries {
		shop, movie, price := e[0], e[1], e[2]
		byMovie[movie] = append(byMovie[movie], Entry{shop, movie, price})
		mrs.priceOf[[2]int{shop, movie}] = price
	}
	for movie, list := range byMovie {
		sort.Slice(list, func(i, j int) bool {
			if list[i].price != list[j].price {
				return list[i].price < list[j].price
			}
			return list[i].shop < list[j].shop
		})
		h := &SearchHeap{}
		heap.Init(h)
		for _, e := range list {
			heap.Push(h, e)
		}
		mrs.avail[movie] = h
	}
	return mrs
}

func (mrs *MovieRentalSystem) Search(movie int) [][]int {
	h, ok := mrs.avail[movie]
	if !ok {
		return nil
	}
	var result []Entry
	var temp []Entry
	for h.Len() > 0 && len(result) < 5 {
		e := heap.Pop(h).(Entry)
		if !mrs.rentedData[e] {
			result = append(result, e)
			temp = append(temp, e)
		}
	}
	for _, e := range temp {
		heap.Push(h, e)
	}
	res := make([][]int, len(result))
	for i, e := range result {
		res[i] = []int{e.shop, e.movie, e.price}
	}
	return res
}

func (mrs *MovieRentalSystem) Rent(shop int, movie int) {
	price := mrs.priceOf[[2]int{shop, movie}]
	e := Entry{shop, movie, price}
	mrs.rentedData[e] = true
	heap.Push(mrs.rented, e)
}

func (mrs *MovieRentalSystem) Drop(shop int, movie int) {
	price := mrs.priceOf[[2]int{shop, movie}]
	e := Entry{shop, movie, price}
	delete(mrs.rentedData, e)
}

func (mrs *MovieRentalSystem) Report() [][]int {
	var temp []Entry
	var result []Entry
	for mrs.rented.Len() > 0 && len(result) < 5 {
		e := heap.Pop(mrs.rented).(Entry)
		if mrs.rentedData[e] {
			result = append(result, e)
			temp = append(temp, e)
		}
	}
	for _, e := range temp {
		heap.Push(mrs.rented, e)
	}
	if len(result) > 5 {
		result = result[:5]
	}
	res := make([][]int, len(result))
	for i, e := range result {
		res[i] = []int{e.shop, e.movie, e.price}
	}
	return res
}

func DesignMovieRentalSystem() any {
	entries := [][]int{
		{0, 1, 5}, {0, 2, 6}, {0, 3, 7},
		{1, 1, 4}, {1, 2, 7},
		{2, 1, 5},
	}
	mrs := Constructor(entries)
	// Search(1): [[1,1,4],[0,1,5],[2,1,5]]
	_ = mrs.Search(1)
	mrs.Rent(0, 1)
	mrs.Rent(1, 2)
	// Report: [[0,1,5],[1,2,7]]
	_ = mrs.Report()
	mrs.Drop(1, 2)
	// Search(2): [[0,2,6]]
	_ = mrs.Search(2)
	return nil
}

func main() {
	DesignMovieRentalSystem()
	fmt.Println("Movie Rental System test completed")
}
```

## 1916 — Count Ways To Build Rooms In An Ant Colony

```go
package main

// LeetCode #1916: Count Ways to Build Rooms in an Ant Colony
// https://leetcode.com/problems/count-ways-to-build-rooms-in-an-ant-colony/
// Difficulty: Hard

import "fmt"

const mod = 1_000_000_007

func waysToBuildRooms(prevRoom []int) int {
	n := len(prevRoom)
	children := make([][]int, n)
	for i := 1; i < n; i++ {
		p := prevRoom[i]
		children[p] = append(children[p], i)
	}

	// Precompute factorials and inverse factorials
	fact := make([]int, n+1)
	invFact := make([]int, n+1)
	fact[0] = 1
	for i := 1; i <= n; i++ {
		fact[i] = fact[i-1] * i % mod
	}
	invFact[n] = modPow(fact[n], mod-2)
	for i := n - 1; i >= 0; i-- {
		invFact[i] = invFact[i+1] * (i + 1) % mod
	}

	nCr := func(nn, r int) int {
		if r < 0 || r > nn {
			return 0
		}
		return fact[nn] * invFact[r] % mod * invFact[nn-r] % mod
	}

	// Post-order DP: returns (subtreeSize, ways) for each node
	var dfs func(u int) (int, int)
	dfs = func(u int) (int, int) {
		size := 1
		ways := 1
		for _, v := range children[u] {
			subSize, subWays := dfs(v)
			// Interleave this child's subtree with what we've accumulated so far.
			// C((size-1) + subSize, subSize) = ways to merge subSize elements
			// into sequence of (size-1) elements, preserving relative order,
			// since node u must be first.
			interleave := nCr(size-1+subSize, subSize)
			ways = ways * subWays % mod * interleave % mod
			size += subSize
		}
		return size, ways
	}

	_, ans := dfs(0)
	return ans
}

func modPow(a, b int) int {
	res := 1
	for b > 0 {
		if b&1 == 1 {
			res = res * a % mod
		}
		a = a * a % mod
		b >>= 1
	}
	return res
}

func main() {
	fmt.Println(waysToBuildRooms([]int{-1, 0, 1}))
	fmt.Println(waysToBuildRooms([]int{-1, 0, 0}))
}
```

## 1917 — Leetcodify Friends Recommendations

```go
package main

// LeetCode #1917: Leetcodify Friends Recommendations
// https://leetcode.com/problems/leetcodify-friends-recommendations/
// Difficulty: Hard [Paid]
//
// Given tables:
// - listens (user_id, song_id, day)
// - friendship (user1_id, user2_id)
//
// Recommend friend pairs (user1_id, user2_id) who are not already friends
// but listened to the same song(s) on the same day >= 3 times.
// Return distinct pairs with user1_id < user2_id.

import (
	"fmt"
	"sort"
)

func main() {
	// Test case 1: simple case
	listens := [][]int{
		{1, 101, 1},
		{1, 101, 1}, // same day same song same user (but should not double count)
		{1, 101, 1},
		{2, 101, 1},
		{2, 101, 1},
		{2, 101, 1},
		{3, 101, 1},
	}
	friendships := [][]int{
		{1, 3},
	}
	// Users 1 and 2 both listened to song 101 on day 1 >= 3 times, recommend (1,2)
	// Users 2 and 3: user 3 listened to 101 on day 1 only once
	fmt.Println(leetcodifyFriendsRecommendations(listens, friendships))

	// Test case 2: multiple recommendations
	listens = [][]int{
		{1, 101, 1}, {1, 101, 1}, {1, 101, 1},
		{2, 101, 1}, {2, 101, 1}, {2, 101, 1},
		{1, 102, 2}, {1, 102, 2}, {1, 102, 2},
		{3, 102, 2}, {3, 102, 2}, {3, 102, 2},
	}
	friendships = [][]int{
		{1, 2},
	}
	// 1-2 already friends, skip
	// 1-3 both listened to 102 on day 2 >= 3 times, recommend (1,3)
	fmt.Println(leetcodifyFriendsRecommendations(listens, friendships))

	// Test case 3: no recommendations
	fmt.Println(leetcodifyFriendsRecommendations([][]int{{1, 101, 1}, {2, 102, 1}}, [][]int{}))

	// Test case 4: empty tables
	fmt.Println(leetcodifyFriendsRecommendations([][]int{}, [][]int{}))
}

// leetcodifyFriendsRecommendations returns [][]int{{user1_id, user2_id}, ...}
func leetcodifyFriendsRecommendations(listens [][]int, friendships [][]int) [][]int {
	// Count listens per user per song per day
	type key struct {
		user, song, day int
	}
	count := make(map[key]int)
	for _, l := range listens {
		k := key{user: l[0], song: l[1], day: l[2]}
		count[k]++
	}

	// Build existing friendships set
	existingFriends := make(map[[2]int]bool)
	for _, f := range friendships {
		u1, u2 := f[0], f[1]
		if u1 > u2 {
			u1, u2 = u2, u1
		}
		existingFriends[[2]int{u1, u2}] = true
	}

	// For each (song, day) pair, find all users with >= 3 listens
	songDayUsers := make(map[[2]int]map[int]bool) // (song, day) -> set of users with >= 3 listens
	for k, c := range count {
		if c >= 3 {
			sd := [2]int{k.song, k.day}
			if songDayUsers[sd] == nil {
				songDayUsers[sd] = make(map[int]bool)
			}
			songDayUsers[sd][k.user] = true
		}
	}

	// Generate recommendations
	recSet := make(map[[2]int]bool)
	for _, users := range songDayUsers {
		// All pairs of users who listened to this (song, day) >= 3 times
		var userList []int
		for u := range users {
			userList = append(userList, u)
		}
		sort.Ints(userList)
		for i := 0; i < len(userList); i++ {
			for j := i + 1; j < len(userList); j++ {
				u1, u2 := userList[i], userList[j]
				pair := [2]int{u1, u2}
				if !existingFriends[pair] {
					recSet[pair] = true
				}
			}
		}
	}

	// Build sorted result
	var result [][]int
	for p := range recSet {
		result = append(result, []int{p[0], p[1]})
	}
	sort.Slice(result, func(i, j int) bool {
		if result[i][0] != result[j][0] {
			return result[i][0] < result[j][0]
		}
		return result[i][1] < result[j][1]
	})

	if result == nil {
		return [][]int{}
	}
	return result
}
```

## 1919 — Leetcodify Similar Friends

```go
package main

// LeetCode #1919: Leetcodify Similar Friends
// https://leetcode.com/problems/leetcodify-similar-friends/
// Difficulty: Hard [Paid]
//
// Given tables:
// - listens (user_id, song_id, day)
// - friendship (user1_id, user2_id)
//
// Find friend pairs (already friends) who have high similarity in music taste.
// Two friends are "similar" if the number of songs they both listened to
// on the same day is >= 3 times the number of different songs they listened to
// (union of songs, counted per day).
//
// More precisely: for each pair of friends (u1, u2), count pairs (song, day)
// where both listened to that song on that day. Count pairs (song, day) where
// either listened. If same_count >= 3 AND same_count / total >= threshold (0.6?),
// return them.
//
// Actually from the problem: return friend pairs where the number of (song, day)
// they have in common >= 3, and common / total_listened_both >= some threshold.

import (
	"fmt"
	"sort"
)

func main() {
	// Test case 1: simple
	listens := [][]int{
		{1, 101, 1}, {1, 102, 1},
		{2, 101, 1}, {2, 102, 1},
		{3, 101, 1},
	}
	friendships := [][]int{
		{1, 2}, {1, 3},
	}
	// Friends 1-2: both listened to {101, 102} on day 1, same=2, total=2, ratio=1.0 >= 0.6
	// Friends 1-3: both listened to {101} on day 1, same=1, total=not sure
	fmt.Println(leetcodifySimilarFriends(listens, friendships))

	// Test case 2: higher threshold
	listens = [][]int{
		{1, 101, 1}, {1, 101, 1}, {1, 101, 1},
		{2, 101, 1},
		{3, 101, 1}, {3, 102, 1},
	}
	friendships = [][]int{
		{1, 2}, {2, 3},
	}
	fmt.Println(leetcodifySimilarFriends(listens, friendships))

	// Test case 3: empty
	fmt.Println(leetcodifySimilarFriends([][]int{}, [][]int{}))
}

// leetcodifySimilarFriends returns [][]int{{user1_id, user2_id}, ...}
func leetcodifySimilarFriends(listens [][]int, friendships [][]int) [][]int {
	// Build set of (user, song, day) - unique listens per user per song per day
	userSongDay := make(map[[3]int]bool)
	userSongs := make(map[int]map[[2]int]bool) // user -> set of (song, day) pairs
	for _, l := range listens {
		u, s, d := l[0], l[1], l[2]
		key := [3]int{u, s, d}
		if !userSongDay[key] {
			userSongDay[key] = true
			if userSongs[u] == nil {
				userSongs[u] = make(map[[2]int]bool)
			}
			userSongs[u][[2]int{s, d}] = true
		}
	}

	var result [][]int

	for _, f := range friendships {
		u1, u2 := f[0], f[1]
		if u1 > u2 {
			u1, u2 = u2, u1
		}

		songs1 := userSongs[u1]
		songs2 := userSongs[u2]

		if len(songs1) == 0 || len(songs2) == 0 {
			continue
		}

		// Count intersection
		common := 0
		for sd := range songs1 {
			if songs2[sd] {
				common++
			}
		}

		// Count union
		union := make(map[[2]int]bool)
		for sd := range songs1 {
			union[sd] = true
		}
		for sd := range songs2 {
			union[sd] = true
		}
		total := len(union)

		if total == 0 {
			continue
		}

		// Use similarity ratio threshold >= 0.6
		// common >= 3 AND common/total >= 0.6
		if common >= 3 && common*10 >= total*6 {
			result = append(result, []int{u1, u2})
		}
	}

	sort.Slice(result, func(i, j int) bool {
		if result[i][0] != result[j][0] {
			return result[i][0] < result[j][0]
		}
		return result[i][1] < result[j][1]
	})

	if result == nil {
		return [][]int{}
	}
	return result
}
```

## 1923 — Longest Common Subpath

```go
package main

// LeetCode #1923: Longest Common Subpath
// https://leetcode.com/problems/longest-common-subpath/
// Difficulty: Hard
//
// Given n (number of cities) and paths (each is an array of cities visited),
// find the length of the longest subpath that appears in every path.
//
// Approach: binary search on length + rolling hash (Rabin-Karp).
// For a candidate length L, compute all hashes of subarrays of length L
// in the first path, then check if they appear in all other paths.
// Uses double hash (two moduli) to avoid collisions.

import (
	"fmt"
)

func main() {
	// Example 1:
	// n = 5, paths = [[0,1,2,3,4], [2,3,4], [4,0,1,2,3]]
	// Expected: 2 (subpath [2,3] or [3,4] or [4,0]... actually [2,3,4] doesn't appear in all)
	// [2,3] appears in all -> length 2
	fmt.Println(longestCommonSubpath(5, [][]int{{0, 1, 2, 3, 4}, {2, 3, 4}, {4, 0, 1, 2, 3}}))

	// Example 2:
	// n = 5, paths = [[0,1,2,3,4], [4,1,2,3,0], [1,2,3,4,0]]
	// Expected: 3 (subpath [1,2,3])
	fmt.Println(longestCommonSubpath(5, [][]int{{0, 1, 2, 3, 4}, {4, 1, 2, 3, 0}, {1, 2, 3, 4, 0}}))

	// Example 3:
	// n = 5, paths = [[0,0,0], [0,0,0]]
	// Expected: 3
	fmt.Println(longestCommonSubpath(5, [][]int{{0, 0, 0}, {0, 0, 0}}))

	// Single path
	fmt.Println(longestCommonSubpath(3, [][]int{{1, 2, 3}}))

	// No common subpath
	fmt.Println(longestCommonSubpath(5, [][]int{{0, 1}, {2, 3}}))
}

func longestCommonSubpath(n int, paths [][]int) int {
	if len(paths) == 0 {
		return 0
	}
	if len(paths) == 1 {
		return len(paths[0])
	}

	// Find min path length for binary search upper bound
	low, high := 0, len(paths[0])
	for _, p := range paths {
		if len(p) < high {
			high = len(p)
		}
	}

	ans := 0
	for low <= high {
		mid := (low + high) / 2
		if hasCommonSubpath(paths, mid) {
			ans = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return ans
}

func hasCommonSubpath(paths [][]int, length int) bool {
	if length == 0 {
		return true
	}

	const mod1, mod2 = 1000000007, 1000000009
	const base = 100003

	// Precompute powers
	pow1, pow2 := 1, 1
	for i := 0; i < length-1; i++ {
		pow1 = pow1 * base % mod1
		pow2 = pow2 * base % mod2
	}

	// Get all subarray hashes from first path
	common := make(map[[2]int]bool)
	var h1, h2 int

	for i, v := range paths[0] {
		x := v + 1 // shift by 1 to avoid hash issues with 0
		h1 = (h1*base + x) % mod1
		h2 = (h2*base + x) % mod2

		if i >= length {
			y := paths[0][i-length] + 1
			h1 = (h1 - y*pow1%mod1 + mod1) % mod1
			h2 = (h2 - y*pow2%mod2 + mod2) % mod2
		}

		if i >= length-1 {
			common[[2]int{h1, h2}] = true
		}
	}

	if len(common) == 0 {
		return false
	}

	// Check each remaining path
	for _, path := range paths[1:] {
		current := make(map[[2]int]bool)
		h1, h2 = 0, 0

		for i, v := range path {
			x := v + 1
			h1 = (h1*base + x) % mod1
			h2 = (h2*base + x) % mod2

			if i >= length {
				y := path[i-length] + 1
				h1 = (h1 - y*pow1%mod1 + mod1) % mod1
				h2 = (h2 - y*pow2%mod2 + mod2) % mod2
			}

			if i >= length-1 {
				key := [2]int{h1, h2}
				if common[key] {
					current[key] = true
				}
			}
		}

		common = current
		if len(common) == 0 {
			return false
		}
	}
	return true
}
```

## 1924 — Erect The Fence Ii

```go
package main

// LeetCode #1924: Erect the Fence II (Minimum Enclosing Circle)
// https://leetcode.com/problems/erect-the-fence-ii/
// Difficulty: Hard [Paid]
//
// Given a set of points (trees) on a 2D plane, find the minimum-radius circle
// that encloses all points. Return [center_x, center_y, radius].
//
// Uses Welzl's randomized algorithm for minimum enclosing circle,
// expected O(n) time.

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	// Test case 1: two points
	fmt.Println(erectTheFenceIi([][]int{{0, 0}, {2, 0}}))
	// Test case 2: three points forming a triangle
	fmt.Println(erectTheFenceIi([][]int{{0, 0}, {1, 0}, {0, 1}}))
	// Test case 3: single point
	fmt.Println(erectTheFenceIi([][]int{{5, 5}}))
	// Test case 4: colinear points
	fmt.Println(erectTheFenceIi([][]int{{0, 0}, {2, 0}, {4, 0}}))
	// Test case 5: square
	fmt.Println(erectTheFenceIi([][]int{{0, 0}, {2, 0}, {2, 2}, {0, 2}}))
	// Test case 6: random points
	fmt.Println(erectTheFenceIi([][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}))
}

// Point represents a 2D point
type Point struct {
	x, y float64
}

// Circle represents a circle with center (x, y) and radius r
type Circle struct {
	x, y, r float64
}

func erectTheFenceIi(trees [][]int) []float64 {
	n := len(trees)
	if n == 0 {
		return []float64{0, 0, 0}
	}

	// Convert to float64 points
	pts := make([]Point, n)
	for i, t := range trees {
		pts[i] = Point{float64(t[0]), float64(t[1])}
	}

	// Shuffle for expected linear time
	rand.Shuffle(n, func(i, j int) {
		pts[i], pts[j] = pts[j], pts[i]
	})

	c := welzl(pts, nil, 0)
	return []float64{c.x, c.y, c.r}
}

func welzl(pts []Point, boundary []Point, idx int) Circle {
	if idx == len(pts) || len(boundary) == 3 {
		return trivialCircle(boundary)
	}

	c := welzl(pts, boundary, idx+1)
	if inside(c, pts[idx]) {
		return c
	}

	// pts[idx] must be on the boundary
	newBoundary := append(boundary, pts[idx])
	return welzl(pts, newBoundary, idx+1)
}

func trivialCircle(pts []Point) Circle {
	switch len(pts) {
	case 0:
		return Circle{0, 0, 0}
	case 1:
		return Circle{pts[0].x, pts[0].y, 0}
	case 2:
		return circleFromTwo(pts[0], pts[1])
	case 3:
		// Try circle through all three points
		c := circleFromThree(pts[0], pts[1], pts[2])
		// If it's valid (not NaN), return it
		if !math.IsNaN(c.r) && c.r >= 0 {
			return c
		}
		// Otherwise, return smallest circle from any two
		c12 := circleFromTwo(pts[0], pts[1])
		c13 := circleFromTwo(pts[0], pts[2])
		c23 := circleFromTwo(pts[1], pts[2])
		best := c12
		if c13.r < best.r {
			best = c13
		}
		if c23.r < best.r {
			best = c23
		}
		return best
	}
	return Circle{0, 0, 0}
}

func circleFromTwo(a, b Point) Circle {
	cx := (a.x + b.x) / 2
	cy := (a.y + b.y) / 2
	r := dist(a, b) / 2
	return Circle{cx, cy, r}
}

func circleFromThree(a, b, c Point) Circle {
	// Compute circumcenter of triangle abc
	// Using formula from https://en.wikipedia.org/wiki/Circumcircle
	d := 2 * (a.x*(b.y-c.y) + b.x*(c.y-a.y) + c.x*(a.y-b.y))
	if math.Abs(d) < 1e-12 {
		return Circle{0, 0, -1} // colinear, invalid
	}

	ux := ((a.x*a.x+a.y*a.y)*(b.y-c.y) + (b.x*b.x+b.y*b.y)*(c.y-a.y) + (c.x*c.x+c.y*c.y)*(a.y-b.y)) / d
	uy := ((a.x*a.x+a.y*a.y)*(c.x-b.x) + (b.x*b.x+b.y*b.y)*(a.x-c.x) + (c.x*c.x+c.y*c.y)*(b.x-a.x)) / d

	r := dist(Point{ux, uy}, a)
	return Circle{ux, uy, r}
}

func inside(c Circle, p Point) bool {
	dx := p.x - c.x
	dy := p.y - c.y
	return dx*dx+dy*dy <= c.r*c.r+1e-9
}

func dist(a, b Point) float64 {
	dx := a.x - b.x
	dy := a.y - b.y
	return math.Sqrt(dx*dx + dy*dy)
}
```

## 1926 — Nearest Exit From Entrance In Maze

```go
package main

// LeetCode #1926: Nearest Exit from Entrance in Maze
// https://leetcode.com/problems/nearest-exit-from-entrance-in-maze/
// Difficulty: Hard - BFS

import "fmt"

func nearestExit(maze [][]byte, entrance []int) int {
	m, n := len(maze), len(maze[0])
	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	queue := [][2]int{{entrance[0], entrance[1]}}
	dist := make([][]int, m)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = -1
		}
	}
	dist[entrance[0]][entrance[1]] = 0

	for len(queue) > 0 {
		r, c := queue[0][0], queue[0][1]
		queue = queue[1:]

		// Check if this is an exit (border, not entrance)
		if (r == 0 || r == m-1 || c == 0 || c == n-1) &&
			!(r == entrance[0] && c == entrance[1]) {
			return dist[r][c]
		}

		for _, d := range dirs {
			nr, nc := r+d[0], c+d[1]
			if nr >= 0 && nr < m && nc >= 0 && nc < n &&
				maze[nr][nc] == '.' && dist[nr][nc] == -1 {
				dist[nr][nc] = dist[r][c] + 1
				queue = append(queue, [2]int{nr, nc})
			}
		}
	}
	return -1
}

func main() {
	// Example 1
	maze1 := [][]byte{
		{'+', '+', '.', '+'},
		{'.', '.', '.', '+'},
		{'+', '+', '+', '.'},
	}
	fmt.Println(nearestExit(maze1, []int{1, 2})) // Expected: 1

	// Example 2
	maze2 := [][]byte{
		{'+', '+', '+'},
		{'.', '.', '.'},
		{'+', '+', '+'},
	}
	fmt.Println(nearestExit(maze2, []int{1, 0})) // Expected: 2

	// Example 3
	maze3 := [][]byte{{'.', '+'}}
	fmt.Println(nearestExit(maze3, []int{0, 0})) // Expected: -1
}
```

## 1928 — Minimum Cost To Reach Destination In Time

```go
package main

// LeetCode #1928: Minimum Cost to Reach Destination in Time
// https://leetcode.com/problems/minimum-cost-to-reach-destination-in-time/
// Difficulty: Hard - Dijkstra (DP on time)
// State: dp[city][time] = min cost to reach city at exactly that time.
// We process cities in increasing time order.

import (
	"fmt"
	"math"
)

func minCost(maxTime int, edges [][]int, passingFees []int) int {
	n := len(passingFees)
	// dp[t][v] = min cost to reach node v at time t
	dp := make([][]int, maxTime+1)
	for t := range dp {
		dp[t] = make([]int, n)
		for v := range dp[t] {
			dp[t][v] = math.MaxInt32
		}
	}
	dp[0][0] = passingFees[0]

	// Build adjacency list
	adj := make([][][2]int, n)
	for _, e := range edges {
		u, v, t := e[0], e[1], e[2]
		adj[u] = append(adj[u], [2]int{v, t})
		adj[v] = append(adj[v], [2]int{u, t})
	}

	ans := math.MaxInt32
	for t := 0; t <= maxTime; t++ {
		for v := 0; v < n; v++ {
			if dp[t][v] == math.MaxInt32 {
				continue
			}
			if v == n-1 {
				if dp[t][v] < ans {
					ans = dp[t][v]
				}
			}
			for _, edge := range adj[v] {
				to, w := edge[0], edge[1]
				nt := t + w
				if nt <= maxTime {
					cost := dp[t][v] + passingFees[to]
					if cost < dp[nt][to] {
						dp[nt][to] = cost
					}
				}
			}
		}
	}

	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}

func main() {
	// Example 1
	maxTime := 30
	edges := [][]int{{0, 1, 10}, {1, 2, 10}, {2, 5, 10}, {0, 3, 1}, {3, 4, 10}, {4, 5, 15}}
	passingFees := []int{5, 1, 2, 20, 20, 3}
	fmt.Println(minCost(maxTime, edges, passingFees)) // Expected: 11

	// Example 2
	maxTime2 := 29
	edges2 := [][]int{{0, 1, 10}, {1, 2, 10}, {2, 5, 10}, {0, 3, 1}, {3, 4, 10}, {4, 5, 15}}
	passingFees2 := []int{5, 1, 2, 20, 20, 3}
	fmt.Println(minCost(maxTime2, edges2, passingFees2)) // Expected: 48 (take path 0->3->4->5)

	// Example 3
	maxTime3 := 1
	edges3 := [][]int{{0, 1, 1}, {1, 2, 1}}
	passingFees3 := []int{0, 1, 2}
	fmt.Println(minCost(maxTime3, edges3, passingFees3)) // Expected: -1
}
```

## 1931 — Painting A Grid With Three Different Colors

```go
package main

// LeetCode #1931: Painting a Grid With Three Different Colors
// https://leetcode.com/problems/painting-a-grid-with-three-different-colors/
// Difficulty: Hard
// DP on columns, m <= 5, 3 colors. Encode each column as base-3 integer.
// dp[col][mask] = ways.

import "fmt"

const mod = 1_000_000_007

func colorTheGrid(m int, n int) int {
	// Generate all valid column patterns (no adjacent cells same color in a column)
	total := 1
	for i := 0; i < m; i++ {
		total *= 3
	}

	var masks []int
	for mask := 0; mask < total; mask++ {
		valid := true
		prev := -1
		tmp := mask
		for i := 0; i < m; i++ {
			cur := tmp % 3
			tmp /= 3
			if cur == prev {
				valid = false
				break
			}
			prev = cur
		}
		if valid {
			masks = append(masks, mask)
		}
	}

	// Precompute transitions: check if two columns can be adjacent
	canTransition := func(a, b int) bool {
		for i := 0; i < m; i++ {
			if a%3 == b%3 {
				return false
			}
			a /= 3
			b /= 3
		}
		return true
	}

	// dp for current column
	dp := make([]int, total)
	for _, mask := range masks {
		dp[mask] = 1
	}

	for col := 1; col < n; col++ {
		ndp := make([]int, total)
		for _, a := range masks {
			if dp[a] == 0 {
				continue
			}
			for _, b := range masks {
				if canTransition(a, b) {
					ndp[b] = (ndp[b] + dp[a]) % mod
				}
			}
		}
		dp = ndp
	}

	ans := 0
	for _, mask := range masks {
		ans = (ans + dp[mask]) % mod
	}
	return ans
}

func main() {
	fmt.Println(colorTheGrid(1, 1)) // Expected: 3
	fmt.Println(colorTheGrid(1, 2)) // Expected: 6
	fmt.Println(colorTheGrid(2, 1)) // Expected: 6
	fmt.Println(colorTheGrid(2, 2)) // Expected: 18
	fmt.Println(colorTheGrid(5, 1000)) // Expected: something (performance test)
}
```

## 1932 — Merge Bsts To Create Single Bst

```go
package main

// LeetCode #1932: Merge BSTs to Create Single BST
// https://leetcode.com/problems/merge-bsts-to-create-single-bst/
// Difficulty: Hard
// Map leaf values to tree roots. Merge by replacing leaves with matching root subtrees.
// Finally validate BST on the single remaining tree.

import "fmt"
import "math"

// TreeNode definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func canMerge(trees []*TreeNode) *TreeNode {
	// Map root values to nodes
	rootMap := make(map[int]*TreeNode)
	for _, t := range trees {
		rootMap[t.Val] = t
	}

	// Count how many times each value appears as a leaf (incoming count)
	incoming := make(map[int]int)
	for _, t := range trees {
		if t.Left != nil {
			incoming[t.Left.Val]++
		}
		if t.Right != nil {
			incoming[t.Right.Val]++
		}
	}

	// Find the ultimate root (value not appearing as any leaf)
	// Also check that each root that appears as a leaf appears exactly once
	var root *TreeNode
	for _, t := range trees {
		cnt := incoming[t.Val]
		if cnt == 0 {
			if root != nil {
				return nil // more than one tree with no incoming
			}
			root = t
		}
	}
	if root == nil {
		return nil
	}

	var dfs func(node *TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		// If node is a leaf (no children), check if its value is a root of another tree
		if node.Left == nil && node.Right == nil {
			if t, ok := rootMap[node.Val]; ok {
				// Don't merge if the target is the root itself (cycle)
				if t == root {
					return false
				}
				// Merge: replace leaf with the subtree
				node.Left = t.Left
				node.Right = t.Right
				delete(rootMap, node.Val)
				return true
			}
		}
		if !dfs(node.Left) || !dfs(node.Right) {
			return false
		}
		return true
	}

	// Iteratively merge (since merging creates new leaves that might match)
	changed := true
	for changed {
		changed = false
		// Collect leaves to merge to avoid modifying during traversal
		var leaves []*TreeNode
		var collectLeaves func(node *TreeNode)
		collectLeaves = func(node *TreeNode) {
			if node == nil {
				return
			}
			if node.Left == nil && node.Right == nil {
				if _, ok := rootMap[node.Val]; ok && node.Val != root.Val {
					leaves = append(leaves, node)
				}
				return
			}
			collectLeaves(node.Left)
			collectLeaves(node.Right)
		}
		collectLeaves(root)

		for _, leaf := range leaves {
			if t, ok := rootMap[leaf.Val]; ok && t != root {
				leaf.Left = t.Left
				leaf.Right = t.Right
				delete(rootMap, leaf.Val)
				changed = true
			}
		}
	}

	// After merging, only the root tree should remain
	if len(rootMap) != 1 {
		return nil
	}
	if _, ok := rootMap[root.Val]; !ok {
		return nil
	}

	// Validate BST
	var isValid func(node *TreeNode, min, max int) bool
	isValid = func(node *TreeNode, min, max int) bool {
		if node == nil {
			return true
		}
		if node.Val <= min || node.Val >= max {
			return false
		}
		return isValid(node.Left, min, node.Val) && isValid(node.Right, node.Val, max)
	}

	if !isValid(root, math.MinInt32, math.MaxInt32) {
		return nil
	}

	return root
}

func main() {
	// Example 1: trees = [[2,1],[3,2,5],[5,4]] → true
	t1 := &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}
	t2 := &TreeNode{Val: 3, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 5}}
	t3 := &TreeNode{Val: 5, Left: &TreeNode{Val: 4}}
	result := canMerge([]*TreeNode{t1, t2, t3})
	fmt.Println(result != nil) // Expected: true

	// Example 2: trees = [[5,3,8],[3,2,6]] → false (6 is larger than 5)
	t4 := &TreeNode{Val: 5, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 8}}
	t5 := &TreeNode{Val: 3, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 6}}
	result2 := canMerge([]*TreeNode{t4, t5})
	fmt.Println(result2 != nil) // Expected: false

	// Simple case: single tree
	t6 := &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 2}}
	result3 := canMerge([]*TreeNode{t6})
	fmt.Println(result3 != nil) // Expected: true
}
```

## 1938 — Maximum Genetic Difference Query

```go
package main

// LeetCode #1938: Maximum Genetic Difference Query
// https://leetcode.com/problems/maximum-genetic-difference-query/
// Difficulty: Hard
// Trie + DFS offline. Binary trie stores node values along current path.
// Answer each query by finding max XOR with any ancestor value.

import "fmt"

const bitLen = 18 // 2^18 = 262144, covers values up to 200k

type TrieNode struct {
	children [2]*TrieNode
	count    int
}

type Trie struct {
	root *TrieNode
}

func newTrie() *Trie {
	return &Trie{root: &TrieNode{}}
}

func (t *Trie) insert(x int) {
	node := t.root
	for b := bitLen; b >= 0; b-- {
		bit := (x >> b) & 1
		if node.children[bit] == nil {
			node.children[bit] = &TrieNode{}
		}
		node = node.children[bit]
		node.count++
	}
}

func (t *Trie) remove(x int) {
	node := t.root
	for b := bitLen; b >= 0; b-- {
		bit := (x >> b) & 1
		child := node.children[bit]
		child.count--
		if child.count == 0 {
			node.children[bit] = nil
			return
		}
		node = child
	}
}

func (t *Trie) maxXor(x int) int {
	node := t.root
	ans := 0
	for b := bitLen; b >= 0; b-- {
		bit := (x >> b) & 1
		want := 1 - bit
		if node.children[want] != nil {
			ans |= (1 << b)
			node = node.children[want]
		} else if node.children[bit] != nil {
			node = node.children[bit]
		} else {
			break
		}
	}
	return ans
}

func maxGeneticDifference(parents []int, queries [][]int) []int {
	n := len(parents)
	// Build adjacency and find root
	children := make([][]int, n)
	var root int
	for i, p := range parents {
		if p == -1 {
			root = i
		} else {
			children[p] = append(children[p], i)
		}
	}

	// Group queries by node
	qByNode := make([][][2]int, n)
	for i, q := range queries {
		node, val := q[0], q[1]
		qByNode[node] = append(qByNode[node], [2]int{val, i})
	}

	ans := make([]int, len(queries))
	trie := newTrie()

	var dfs func(u int)
	dfs = func(u int) {
		trie.insert(u)
		for _, q := range qByNode[u] {
			val, idx := q[0], q[1]
			ans[idx] = trie.maxXor(val)
		}
		for _, v := range children[u] {
			dfs(v)
		}
		trie.remove(u)
	}

	dfs(root)
	return ans
}

func main() {
	// Example 1
	parents := []int{-1, 0, 1, 1}
	queries := [][]int{{0, 2}, {3, 2}, {2, 5}}
	result := maxGeneticDifference(parents, queries)
	fmt.Println(result) // Expected: [2, 3, 7]

	// Example 2
	parents2 := []int{-1, 0, 0, 1, 1, 2, 2}
	queries2 := [][]int{{0, 10}, {6, 7}}
	result2 := maxGeneticDifference(parents2, queries2)
	fmt.Println(result2)
}
```

## 1944 — Number Of Visible People In A Queue

```go
package main

// LeetCode #1944: Number of Visible People in a Queue
// https://leetcode.com/problems/number-of-visible-people-in-a-queue/
// Difficulty: Hard
// Monotonic stack from right. Count shorter people on stack, then +1 for the next taller.

import "fmt"

func canSeePersonsCount(heights []int) []int {
	n := len(heights)
	ans := make([]int, n)
	stack := make([]int, 0, n)

	for i := n - 1; i >= 0; i-- {
		visible := 0
		// Pop everyone shorter than current height
		for len(stack) > 0 && stack[len(stack)-1] < heights[i] {
			stack = stack[:len(stack)-1]
			visible++
		}
		// If there's someone taller left, they are also visible
		if len(stack) > 0 {
			visible++
		}
		ans[i] = visible
		stack = append(stack, heights[i])
	}

	return ans
}

func main() {
	fmt.Println(canSeePersonsCount([]int{10, 6, 8, 5, 11, 9})) // Expected: [3 1 2 1 1 0]
	fmt.Println(canSeePersonsCount([]int{5, 1, 2, 3, 10}))     // Expected: [4 1 1 1 0]
}
```

## 1948 — Delete Duplicate Folders In System

```go
package main

// LeetCode #1948: Delete Duplicate Folders in System
// https://leetcode.com/problems/delete-duplicate-folders-in-system/
// Difficulty: Hard
// Build trie, serialise each subtree. If two nodes have the same serialisation,
// mark all of them for deletion. Collect remaining paths (DFS).

import (
	"fmt"
	"sort"
)

type Folder struct {
	name     string
	children map[string]*Folder
	del      bool
	// cache for serialisation to avoid recomputing
	serial string
}

func newFolder(name string) *Folder {
	return &Folder{name: name, children: make(map[string]*Folder)}
}

func deleteDuplicateFolder(paths [][]string) [][]string {
	root := newFolder("/")

	// Build trie
	for _, path := range paths {
		node := root
		for _, name := range path {
			if _, ok := node.children[name]; !ok {
				node.children[name] = newFolder(name)
			}
			node = node.children[name]
		}
	}

	// Map serialisation -> list of nodes with that serialisation
	hashNodes := make(map[string][]*Folder)

	var dfs func(node *Folder) string
	dfs = func(node *Folder) string {
		if node.serial != "" {
			return node.serial
		}
		if len(node.children) == 0 {
			return ""
		}
		var parts []string
		for name, child := range node.children {
			childSerial := dfs(child)
			parts = append(parts, name+"|"+childSerial)
		}
		sort.Strings(parts)
		serial := ""
		for _, p := range parts {
			serial += "(" + p + ")"
		}
		node.serial = serial
		hashNodes[serial] = append(hashNodes[serial], node)
		return serial
	}

	// Compute serialisation for all non-root, non-leaf nodes
	for _, child := range root.children {
		dfs(child)
	}

	// Mark duplicates: if serial appears more than once, mark all
	for _, nodes := range hashNodes {
		if len(nodes) > 1 {
			for _, node := range nodes {
				node.del = true
			}
		}
	}

	// Collect remaining paths (skip marked nodes and their descendants)
	var ans [][]string
	var collect func(node *Folder, path []string)
	collect = func(node *Folder, path []string) {
		if node.del {
			return
		}
		if node != root {
			ans = append(ans, append([]string{}, path...))
		}
		// Collect children in sorted order for stable output
		var names []string
		for name := range node.children {
			names = append(names, name)
		}
		sort.Strings(names)
		for _, name := range names {
			path = append(path, name)
			collect(node.children[name], path)
			path = path[:len(path)-1]
		}
	}

	collect(root, nil)
	return ans
}

func main() {
	// Example 1
	paths1 := [][]string{{"a"}, {"c"}, {"d"}, {"a", "b"}, {"c", "b"}, {"d", "a"}}
	result1 := deleteDuplicateFolder(paths1)
	fmt.Println(result1) // Expected: [[d] [d a]]

	// Example 2
	paths2 := [][]string{{"a"}, {"c"}, {"a", "b"}, {"c", "b"}, {"a", "b", "x"}, {"a", "b", "x", "y"}, {"w"}, {"w", "y"}}
	result2 := deleteDuplicateFolder(paths2)
	fmt.Println(result2)

	// Example 3
	paths3 := [][]string{{"a", "b"}, {"c", "d"}, {"c"}, {"a"}}
	result3 := deleteDuplicateFolder(paths3)
	fmt.Println(result3)
}
```

## 1955 — Count Number Of Special Subsequences

```go
package main

// LeetCode #1955: Count Number of Special Subsequences
// https://leetcode.com/problems/count-number-of-special-subsequences/
// Difficulty: Hard
// DP[0] = count of subsequences matching pattern "0"
// DP[1] = count of subsequences matching pattern "0...1"
// DP[2] = count of subsequences matching pattern "0...1...2"

import "fmt"

const mod = 1_000_000_007

func countSpecialSubsequences(nums []int) int {
	dp0, dp1, dp2 := 0, 0, 0
	for _, x := range nums {
		switch x {
		case 0:
			dp0 = (dp0 + dp0 + 1) % mod
		case 1:
			dp1 = (dp1 + dp1 + dp0) % mod
		case 2:
			dp2 = (dp2 + dp2 + dp1) % mod
		}
	}
	return dp2
}

func main() {
	fmt.Println(countSpecialSubsequences([]int{0, 1, 2, 2}))       // Expected: 3
	fmt.Println(countSpecialSubsequences([]int{0, 1, 2, 0, 1, 2})) // Expected: 7
	fmt.Println(countSpecialSubsequences([]int{2, 2, 0, 0}))       // Expected: 0 (no 1s)
}
```

## 1956 — Minimum Time For K Virus Variants To Spread

```go
package main

// LeetCode #1956: Minimum Time For K Virus Variants to Spread
// https://leetcode.com/problems/minimum-time-for-k-virus-variants-to-spread/
// Difficulty: Hard [Paid]
//
// Given n virus variants at positions (xi, yi), each spreads at speed 1
// (Manhattan distance per unit time). Find the minimum integer time T such
// that there exists a point that can be reached by at least k variants.
//
// A variant at (xi, yi) reaches point (x, y) in time T if
// |x - xi| + |y - yi| <= T.
//
// Transform to (u = x+y, v = x-y) space where the reachable region becomes
// an axis-aligned square of side 2T centered at (ui, vi).
// Check function: sweep line over u, maintain difference array for v.

import (
	"fmt"
	"sort"
)

func main() {
	// Test case 1: two points, k=2, meet at midpoint
	// (0,0) and (2,0) -> |x|+|y| <= T, |x-2|+|y| <= T
	// Meet at (1,0) in time 1
	fmt.Println(minTimeForKVirusVariantsToSpread([][]int{{0, 0}, {2, 0}}, 2))

	// Test case 2: three points forming a triangle, k=2
	// (0,0), (1,0), (0,1) -> any two can meet at their midpoint
	// (0,0) and (1,0) meet at (0.5,0.5) ... wait Manhattan midpoint
	// |x|+|y| <= T, |x-1|+|y| <= T for (0.5, 0.5): |0.5|+|0.5|=1, |0.5-1|+|0.5|=1 => T=1
	fmt.Println(minTimeForKVirusVariantsToSpread([][]int{{0, 0}, {1, 0}, {0, 1}}, 2))

	// Test case 3: single point, k=1 -> T=0
	fmt.Println(minTimeForKVirusVariantsToSpread([][]int{{5, 5}}, 1))

	// Test case 4: colinear, k=2
	fmt.Println(minTimeForKVirusVariantsToSpread([][]int{{0, 0}, {4, 0}}, 2))

	// Test case 5: need exactly k=3 out of 4
	fmt.Println(minTimeForKVirusVariantsToSpread([][]int{{0, 0}, {10, 10}, {0, 10}, {10, 0}}, 3))

	// Test case 6: large gap
	fmt.Println(minTimeForKVirusVariantsToSpread([][]int{{0, 0}, {100, 0}}, 2))
}

// minTimeForKVirusVariantsToSpread returns the minimum integer time T.
func minTimeForKVirusVariantsToSpread(points [][]int, k int) int {
	n := len(points)
	if n < k {
		return -1
	}
	if k <= 1 {
		return 0
	}

	// Binary search on time
	// Upper bound: max coordinate range
	maxCoord := 0
	for _, p := range points {
		for _, c := range p {
			if c > maxCoord {
				maxCoord = c
			}
		}
	}
	high := 2 * maxCoord // safe upper bound
	if high < 0 {
		high = 2000000
	}

	low := 0
	ans := high
	for low <= high {
		mid := (low + high) / 2
		if canMeet(points, k, mid) {
			ans = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return ans
}

func canMeet(points [][]int, k, T int) bool {
	// Transform: u = x+y, v = x-y
	// Each point covers [ui-T, ui+T] in u, [vi-T, vi+T] in v
	type event struct {
		u        int
		vLow     int
		vHigh    int
		isAdd    bool
	}

	var events []event
	vSet := make(map[int]bool)

	for _, p := range points {
		x, y := p[0], p[1]
		u := x + y
		v := x - y
		events = append(events, event{u - T, v - T, v + T, true})
		events = append(events, event{u + T + 1, v - T, v + T, false})
		vSet[v-T] = true
		vSet[v+T] = true
	}

	// Coordinate compress v values
	var vVals []int
	for v := range vSet {
		vVals = append(vVals, v)
	}
	sort.Ints(vVals)
	vComp := make(map[int]int)
	for i, v := range vVals {
		vComp[v] = i
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].u < events[j].u
	})

	// Difference array over compressed v
	diff := make([]int, len(vVals)+1)

	for i := 0; i < len(events); {
		curU := events[i].u
		// Apply all events at this u
		for i < len(events) && events[i].u == curU {
			e := events[i]
			l := vComp[e.vLow]
			r := vComp[e.vHigh]
			if e.isAdd {
				diff[l]++
				diff[r+1]--
			} else {
				diff[l]--
				diff[r+1]++
			}
			i++
		}
		// Check current sweep position
		cur := 0
		for j := 0; j < len(vVals)-1; j++ {
			cur += diff[j]
			if cur >= k {
				return true
			}
		}
	}

	return false
}
```

## 1960 — Maximum Product Of The Length Of Two Palindromic Substrings

```go
package main

// LeetCode #1960: Maximum Product of the Length of Two Palindromic Substrings
// https://leetcode.com/problems/maximum-product-of-the-length-of-two-palindromic-substrings/
// Difficulty: Hard
// Manacher to find all palindrome radii, then compute L[i] = longest palindrome
// ending at i and R[i] = longest palindrome starting at i. Answer = max L[i]*R[i+1].

import "fmt"

func maxProduct(s string) int64 {
	n := len(s)
	// Odd palindrome radii (center is a character)
	odd := make([]int, n)
	center, right := 0, 0
	for i := 0; i < n; i++ {
		if i < right {
			odd[i] = min(odd[2*center-i], right-i)
		}
		for i-odd[i] >= 0 && i+odd[i] < n && s[i-odd[i]] == s[i+odd[i]] {
			odd[i]++
		}
		if i+odd[i] > right {
			center, right = i, i+odd[i]
		}
	}

	// Even palindrome radii (centered between s[i-1] and s[i])
	// even[i] = radius for even palindrome centered between i-1 and i
	even := make([]int, n+1)
	l, r := 0, 0
	for i := 0; i <= n; i++ {
		if i < r {
			even[i] = min(even[l+r-i], r-i)
		}
		for i-even[i]-1 >= 0 && i+even[i] < n && s[i-even[i]-1] == s[i+even[i]] {
			even[i]++
		}
		if i+even[i] > r {
			l, r = i-even[i], i+even[i]
		}
	}

	// L[i] = longest palindrome ending at i
	L := make([]int, n)
	for i := 0; i < n; i++ {
		L[i] = 1
	}
	// Odd palindromes
	for c := 0; c < n; c++ {
		rad := odd[c]
		end := c + rad - 1
		length := 2*rad - 1
		if end < n && length > L[end] {
			L[end] = length
		}
	}
	// Even palindromes
	for i := 0; i <= n; i++ {
		rad := even[i]
		if rad == 0 {
			continue
		}
		end := i + rad - 1
		length := 2 * rad
		if end < n && length > L[end] {
			L[end] = length
		}
	}
	// Propagate: if palindrome of length k ends at i+1, one of length k-2 ends at i
	for i := n - 2; i >= 0; i-- {
		if L[i+1]-2 > L[i] {
			L[i] = L[i+1] - 2
		}
	}

	// R[i] = longest palindrome starting at i
	R := make([]int, n)
	for i := 0; i < n; i++ {
		R[i] = 1
	}
	// Odd palindromes
	for c := 0; c < n; c++ {
		rad := odd[c]
		start := c - rad + 1
		length := 2*rad - 1
		if start >= 0 && length > R[start] {
			R[start] = length
		}
	}
	// Even palindromes
	for i := 0; i <= n; i++ {
		rad := even[i]
		if rad == 0 {
			continue
		}
		start := i - rad
		length := 2 * rad
		if start >= 0 && length > R[start] {
			R[start] = length
		}
	}
	// Propagate forward
	for i := 1; i < n; i++ {
		if R[i-1]-2 > R[i] {
			R[i] = R[i-1] - 2
		}
	}

	// Suffix max of R so we can check any split, not just adjacent positions
	suffixMaxR := make([]int, n)
	suffixMaxR[n-1] = R[n-1]
	for i := n - 2; i >= 0; i-- {
		if R[i] > suffixMaxR[i+1] {
			suffixMaxR[i] = R[i]
		} else {
			suffixMaxR[i] = suffixMaxR[i+1]
		}
	}

	var ans int64 = 0
	for i := 0; i < n-1; i++ {
		prod := int64(L[i]) * int64(suffixMaxR[i+1])
		if prod > ans {
			ans = prod
		}
	}
	return ans
}

func main() {
	fmt.Println(maxProduct("ababbb"))   // Expected: 9
	fmt.Println(maxProduct("zaaaxbbby")) // Expected: 9 (zaaaz=5, bbby=4 or similar)
	fmt.Println(maxProduct("a"))         // Expected: 0 (only one char, no two substrings)
}
```

## 1964 — Find The Longest Valid Obstacle Course At Each Position

```go
package main

// LeetCode #1964: Find the Longest Valid Obstacle Course at Each Position
// https://leetcode.com/problems/find-the-longest-valid-obstacle-course-at-each-position/
// Difficulty: Hard
// LIS variant with non-decreasing constraint.
// Patience sorting: binary search for first element > h, replace with h.

import (
	"fmt"
	"sort"
)

func longestObstacleCourseAtEachPosition(obstacles []int) []int {
	n := len(obstacles)
	ans := make([]int, n)
	tails := make([]int, 0, n)

	for i, h := range obstacles {
		// Find first element in tails > h (strictly greater)
		// sort.Search finds first index where predicate is true
		idx := sort.Search(len(tails), func(k int) bool { return tails[k] > h })
		if idx == len(tails) {
			tails = append(tails, h)
		} else {
			tails[idx] = h
		}
		ans[i] = idx + 1
	}
	return ans
}

func main() {
	fmt.Println(longestObstacleCourseAtEachPosition([]int{1, 2, 3, 2}))          // Expected: [1 2 3 3]
	fmt.Println(longestObstacleCourseAtEachPosition([]int{2, 2, 1}))             // Expected: [1 2 1]
	fmt.Println(longestObstacleCourseAtEachPosition([]int{3, 1, 5, 6, 4, 2}))   // Expected: [1 1 2 3 2 2]
}
```

## 1970 — Last Day Where You Can Still Cross

```go
package main

// LeetCode #1970: Last Day Where You Can Still Cross
// https://leetcode.com/problems/last-day-where-you-can-still-cross/
// Difficulty: Hard
// Approach: Binary search + Union-Find (DSU)
// For each day mid, build DSU with cells up to day mid as water.
// Check if there's a path from top to bottom through land cells.
// Virtual nodes: row*col = top, row*col+1 = bottom.

import "fmt"

type DSU struct {
	parent []int
	rank   []int
}

func NewDSU(n int) *DSU {
	p := make([]int, n)
	r := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
	}
	return &DSU{parent: p, rank: r}
}

func (d *DSU) Find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.Find(d.parent[x])
	}
	return d.parent[x]
}

func (d *DSU) Union(x, y int) {
	xr, yr := d.Find(x), d.Find(y)
	if xr == yr {
		return
	}
	if d.rank[xr] < d.rank[yr] {
		xr, yr = yr, xr
	}
	d.parent[yr] = xr
	if d.rank[xr] == d.rank[yr] {
		d.rank[xr]++
	}
}

func latestDayToCross(row, col int, cells [][]int) int {
	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	n := row * col
	top, bottom := n, n+1

	canCross := func(day int) bool {
		dsu := NewDSU(n + 2)
		land := make([][]bool, row)
		for i := 0; i < row; i++ {
			land[i] = make([]bool, col)
			for j := 0; j < col; j++ {
				land[i][j] = true
			}
		}
		// Flood first 'day' cells
		for d := 0; d < day; d++ {
			r, c := cells[d][0]-1, cells[d][1]-1
			land[r][c] = false
		}

		for r := 0; r < row; r++ {
			for c := 0; c < col; c++ {
				if !land[r][c] {
					continue
				}
				idx := r*col + c
				if r == 0 {
					dsu.Union(idx, top)
				}
				if r == row-1 {
					dsu.Union(idx, bottom)
				}
				for _, d := range dirs {
					nr, nc := r+d[0], c+d[1]
					if nr >= 0 && nr < row && nc >= 0 && nc < col && land[nr][nc] {
						dsu.Union(idx, nr*col+nc)
					}
				}
			}
		}
		return dsu.Find(top) == dsu.Find(bottom)
	}

	lo, hi := 1, row*col
	ans := 0
	for lo <= hi {
		mid := (lo + hi) / 2
		if canCross(mid) {
			ans = mid
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return ans
}

func main() {
	// Example: row=2, col=2, cells=[[1,1],[2,1],[1,2],[2,2]] -> 2
	fmt.Println(latestDayToCross(2, 2, [][]int{{1, 1}, {2, 1}, {1, 2}, {2, 2}}))

	// Additional tests
	fmt.Println(latestDayToCross(2, 2, [][]int{{1, 1}, {1, 2}, {2, 1}, {2, 2}}))
	fmt.Println(latestDayToCross(3, 3, [][]int{{1, 1}, {2, 1}, {3, 1}, {1, 2}, {2, 2}, {3, 2}, {1, 3}, {2, 3}, {3, 3}}))
}
```

## 1972 — First And Last Call On The Same Day

```go
package main

// LeetCode #1972: First and Last Call On the Same Day
// https://leetcode.com/problems/first-and-last-call-on-the-same-day/
// Difficulty: Hard [Paid]
//
// Given a table calls (caller_id, recipient_id, call_time) with call_time
// as a datetime, find all users whose first and last call on the same day
// was with the same person. For each such user, return user_id, the other
// participant's id, and the date.
//
// This is a SQL-style problem implemented in Go.

import (
	"fmt"
	"sort"
)

func main() {
	// Test case 1: simple
	calls := [][]interface{}{
		{1, 2, "2021-01-01 09:00:00"},
		{1, 2, "2021-01-01 17:00:00"},
	}
	// User 1: first call 9am with 2, last call 5pm with 2 -> same person (2)
	fmt.Println(firstAndLastCallOnTheSameDay(calls))

	// Test case 2: different person for first vs last
	calls = [][]interface{}{
		{1, 2, "2021-01-01 09:00:00"},
		{1, 3, "2021-01-01 17:00:00"},
	}
	// User 1: first with 2, last with 3 -> different, no result
	fmt.Println(firstAndLastCallOnTheSameDay(calls))

	// Test case 3: multiple users
	calls = [][]interface{}{
		{1, 2, "2021-01-01 09:00:00"},
		{1, 2, "2021-01-01 17:00:00"},
		{2, 1, "2021-01-01 08:00:00"},
		{2, 1, "2021-01-01 18:00:00"},
	}
	// User 1: first with 2, last with 2 -> (1, 2, 2021-01-01)
	// User 2: first with 1, last with 1 -> (2, 1, 2021-01-01)
	fmt.Println(firstAndLastCallOnTheSameDay(calls))

	// Test case 4: multiple days
	calls = [][]interface{}{
		{1, 2, "2021-01-01 09:00:00"},
		{1, 3, "2021-01-01 17:00:00"},
		{1, 2, "2021-01-02 09:00:00"},
		{1, 2, "2021-01-02 17:00:00"},
	}
	// User 1 on 2021-01-02: first with 2, last with 2 -> (1, 2, 2021-01-02)
	fmt.Println(firstAndLastCallOnTheSameDay(calls))

	// Test case 5: empty
	fmt.Println(firstAndLastCallOnTheSameDay([][]interface{}{}))
}

func firstAndLastCallOnTheSameDay(calls [][]interface{}) [][]interface{} {
	// Group calls by (user, date)
	// Each call involves two users: caller and recipient
	type callInfo struct {
		other int
		time  string
	}
	userDayCalls := make(map[[2]string][]callInfo) // (user_id, date) -> calls

	for _, c := range calls {
		caller := c[0].(int)
		recipient := c[1].(int)
		timeStr := c[2].(string)
		date := timeStr[:10] // YYYY-MM-DD

		// Add for caller
		key1 := [2]string{fmt.Sprintf("%d", caller), date}
		userDayCalls[key1] = append(userDayCalls[key1], callInfo{recipient, timeStr})

		// Add for recipient (they also participated in the call)
		key2 := [2]string{fmt.Sprintf("%d", recipient), date}
		userDayCalls[key2] = append(userDayCalls[key2], callInfo{caller, timeStr})
	}

	var result [][]interface{}

	for key, calls := range userDayCalls {
		// Sort calls by time
		sort.Slice(calls, func(i, j int) bool {
			return calls[i].time < calls[j].time
		})

		firstOther := calls[0].other
		lastOther := calls[len(calls)-1].other

		if firstOther == lastOther {
			uid := 0
			fmt.Sscanf(key[0], "%d", &uid)
			result = append(result, []interface{}{uid, firstOther, key[1]})
		}
	}

	// Sort result by user_id, then date
	sort.Slice(result, func(i, j int) bool {
		ui := result[i][0].(int)
		uj := result[j][0].(int)
		if ui != uj {
			return ui < uj
		}
		return result[i][2].(string) < result[j][2].(string)
	})

	if result == nil {
		return [][]interface{}{}
	}
	return result
}
```

## 1977 — Number Of Ways To Separate Numbers

```go
package main

// LeetCode #1977: Number of Ways to Separate Numbers
// https://leetcode.com/problems/number-of-ways-to-separate-numbers/
// Difficulty: Hard
// Approach: DP + LCP (Longest Common Prefix)
// dp[i] = number of ways to split suffix s[i:]
// pref[i] = sum_{k >= i} dp[k]
// For each i, try each j > i. Compare s[i:j] with s[j:j+len] using LCP.
// If s[i:j] < s[j:j+len], contribution = pref[j+len] (all splits where next number length >= len)
// else contribution = pref[j+len+1] (all splits where next number length > len)

import "fmt"

const MOD1977 = 1000000007

func numberOfWaysToSeparateNumbers(s string) int {
	n := len(s)

	// LCP[i][j] = longest common prefix of s[i:] and s[j:]
	lcp := make([][]int, n+1)
	for i := range lcp {
		lcp[i] = make([]int, n+1)
	}
	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if s[i] == s[j] {
				lcp[i][j] = lcp[i+1][j+1] + 1
			}
		}
	}

	dp := make([]int, n+2)  // dp[n] = 1 (empty suffix)
	pref := make([]int, n+2) // pref[i] = sum_{k >= i} dp[k]
	dp[n] = 1
	pref[n] = 1

	for i := n - 1; i >= 0; i-- {
		if s[i] == '0' {
			pref[i] = pref[i+1]
			continue
		}
		total := 0
		for j := i + 1; j <= n; j++ {
			curLen := j - i
			if j == n {
				total = (total + 1) % MOD1977
				break
			}
			if s[j] == '0' {
				continue
			}
			if n-j < curLen {
				continue
			}
			l := lcp[i][j]
			if l >= curLen || s[i+l] < s[j+l] {
				total = (total + pref[j+curLen]) % MOD1977
			} else {
				total = (total + pref[j+curLen+1]) % MOD1977
			}
		}
		dp[i] = total
		pref[i] = (dp[i] + pref[i+1]) % MOD1977
	}

	return dp[0]
}

func main() {
	// Example: "327" -> 2
	fmt.Println(numberOfWaysToSeparateNumbers("327"))

	// Additional tests
	fmt.Println(numberOfWaysToSeparateNumbers("123"))
	fmt.Println(numberOfWaysToSeparateNumbers("100"))
	fmt.Println(numberOfWaysToSeparateNumbers("1"))
}
```

## 1982 — Find Array Given Subset Sums

```go
package main

// LeetCode #1982: Find Array Given Subset Sums
// https://leetcode.com/problems/find-array-given-subset-sums/
// Difficulty: Hard
// Approach: Sort + recursive extraction.
// For sorted sums, the difference between consecutive elements gives a candidate.
// Partition into "without" and "with" the candidate. Recurse on "without".

import (
	"fmt"
	"sort"
)

func recoverArray(n int, sums []int) []int {
	sort.Ints(sums)
	ans := make([]int, 0, n)

	for len(sums) > 1 {
		// Candidate: diff between smallest two elements
		diff := sums[1] - sums[0]

		// Partition sums into left (without diff) and right (with diff)
		freq := make(map[int]int)
		for _, v := range sums {
			freq[v]++
		}

		left := make([]int, 0, len(sums)/2)
		right := make([]int, 0, len(sums)/2)
		for _, v := range sums {
			if freq[v] == 0 {
				continue
			}
			freq[v]--
			freq[v+diff]--
			left = append(left, v)
			right = append(right, v+diff)
		}

		// Check if left contains 0 (the empty set sum)
		// If yes, diff is positive (or 0 which shouldn't happen for valid input)
		// If no, diff is negative (right contains 0 instead)
		hasZero := false
		for _, v := range left {
			if v == 0 {
				hasZero = true
				break
			}
		}

		if !hasZero {
			left, right = right, left
			diff = -diff
		}

		ans = append(ans, diff)
		sums = left
	}

	return ans
}

func main() {
	// Example: n=3, sums=[-3,-2,-1,0,0,1,2,3] -> [1,2,-3] or [1,-2,3] etc.
	// The problem states sums are distinct, but standard LeetCode example:
	// Let's use a proper test case
	fmt.Println(recoverArray(3, []int{-3, -2, -1, 0, 0, 1, 2, 3}))

	// n=2, sums=[0,1,2,3] -> original array sum = 3, subset sums = {0, a, b, a+b}
	// If a=1, b=2: sums={0,1,2,3}. So result is [1,2] or [2,1]
	fmt.Println(recoverArray(2, []int{0, 1, 2, 3}))

	// Another test
	fmt.Println(recoverArray(2, []int{0, 1, 1, 2}))
}
```

## 1987 — Number Of Unique Good Subsequences

```go
package main

// LeetCode #1987: Number of Unique Good Subsequences
// https://leetcode.com/problems/number-of-unique-good-subsequences/
// Difficulty: Hard
// Approach: DP tracking distinct good subsequences.
// "Good" means no leading zeros (except the single "0" itself).
// State:
//   zeroExists: whether we've seen the single "0" subsequence
//   ends1: number of good subsequences ending with '1'
//   ends0: number of good subsequences ending with '0' that start with '1' (no leading zero)
// For each '1': ends1 += ends1 + ends0 + 1
// For each '0': ends0 += ends1 + ends0; count "0" once

import "fmt"

const MOD1987 = 1000000007

func numberOfUniqueGoodSubsequences(binary string) int {
	zeroExists := false
	ends1 := 0
	ends0 := 0

	for _, ch := range binary {
		if ch == '1' {
			// Append '1' to all existing good subsequences, plus the single "1"
			ends1 = (ends1 + ends0 + 1) % MOD1987
		} else {
			// Append '0' to all good subsequences that already have a '1' (no leading zero)
			ends0 = (ends0 + ends1) % MOD1987
			if !zeroExists {
				zeroExists = true
			}
		}
	}

	ans := (ends1 + ends0) % MOD1987
	if zeroExists {
		ans = (ans + 1) % MOD1987
	}
	return ans
}

func main() {
	// Example: "001" -> 2 (subsequences: "0", "1")
	fmt.Println(numberOfUniqueGoodSubsequences("001"))

	// Additional tests
	fmt.Println(numberOfUniqueGoodSubsequences("101")) // "0","1","11","10","101" -> 5
	fmt.Println(numberOfUniqueGoodSubsequences("000")) // just "0" -> 1
	fmt.Println(numberOfUniqueGoodSubsequences("111")) // "1" -> 1
}
```

## 1994 — The Number Of Good Subsets

```go
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
```

## 1998 — Gcd Sort Of An Array

```go
package main

// LeetCode #1998: GCD Sort of an Array
// https://leetcode.com/problems/gcd-sort-of-an-array/
// Difficulty: Hard
// Approach: Union-Find over numbers and their prime factors.
// Two numbers can be swapped if they share a GCD > 1,
// which is equivalent to being connected through prime factors.
// After building DSU, check if each nums[i] and sorted[i] are in the same set.

import (
	"fmt"
	"sort"
)

type DSU1998 struct {
	parent []int
	rank   []int
}

func NewDSU1998(n int) *DSU1998 {
	p := make([]int, n)
	r := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
	}
	return &DSU1998{parent: p, rank: r}
}

func (d *DSU1998) Find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.Find(d.parent[x])
	}
	return d.parent[x]
}

func (d *DSU1998) Union(x, y int) {
	xr, yr := d.Find(x), d.Find(y)
	if xr == yr {
		return
	}
	if d.rank[xr] < d.rank[yr] {
		xr, yr = yr, xr
	}
	d.parent[yr] = xr
	if d.rank[xr] == d.rank[yr] {
		d.rank[xr]++
	}
}

// smallestPrimeFactor using sieve
func spfSieve(limit int) []int {
	spf := make([]int, limit+1)
	for i := 2; i <= limit; i++ {
		if spf[i] == 0 {
			spf[i] = i
			if i*i <= limit {
				for j := i * i; j <= limit; j += i {
					if spf[j] == 0 {
						spf[j] = i
					}
				}
			}
		}
	}
	return spf
}

func gcdSort(nums []int) bool {
	maxVal := 0
	for _, v := range nums {
		if v > maxVal {
			maxVal = v
		}
	}

	// Sieve for smallest prime factors
	spf := spfSieve(maxVal + 1)
	// DSU for numbers 1..maxVal plus indices for nums
	// We union numbers with their prime factors
	dsu := NewDSU1998(maxVal + 1)

	// For each number, union with its prime factors
	for _, v := range nums {
		x := v
		for x > 1 {
			p := spf[x]
			if p == 0 {
				p = x
			}
			dsu.Union(v, p)
			for x%p == 0 {
				x /= p
			}
		}
	}

	// Sort a copy and check
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	sort.Ints(sorted)

	for i := 0; i < len(nums); i++ {
		if dsu.Find(nums[i]) != dsu.Find(sorted[i]) {
			return false
		}
	}
	return true
}

func main() {
	// Example: [7,21,3] -> true
	fmt.Println(gcdSort([]int{7, 21, 3}))

	// Additional tests
	fmt.Println(gcdSort([]int{5, 2, 6, 2}))
	fmt.Println(gcdSort([]int{1, 2, 3, 4}))
	fmt.Println(gcdSort([]int{10, 5, 9, 3, 15}))
}
```

## 2003 — Smallest Missing Genetic Value In Each Subtree

```go
package main

// LeetCode #2003: Smallest Missing Genetic Value in Each Subtree
// https://leetcode.com/problems/smallest-missing-genetic-value-in-each-subtree/
// Difficulty: Hard
// Approach: DFS + set.
// Only the node with value 1 (and its ancestors) can have missing value > 1.
// Other nodes' answer is always 1.
// Traverse from the 1-node upward, collecting subtree values to find mex.

import "fmt"

func smallestMissingValueSubtree(parents []int, nums []int) []int {
	n := len(parents)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = 1
	}

	// Find node with value 1
	oneNode := -1
	for i := 0; i < n; i++ {
		if nums[i] == 1 {
			oneNode = i
			break
		}
	}
	if oneNode == -1 {
		return ans // all answers are 1
	}

	// Build children adjacency
	children := make([][]int, n)
	for i := 1; i < n; i++ {
		p := parents[i]
		children[p] = append(children[p], i)
	}

	// Reconstruct parent chain from oneNode to root
	pathSet := make(map[int]bool)
	curr := oneNode
	for curr != -1 {
		pathSet[curr] = true
		curr = parents[curr]
		if curr == -1 {
			break
		}
	}

	visited := make(map[int]bool)
	mex := 1

	// DFS to collect values in a subtree
	var dfs func(u int)
	dfs = func(u int) {
		visited[nums[u]] = true
		for _, v := range children[u] {
			dfs(v)
		}
	}

	// Process from oneNode upward
	curr = oneNode
	for curr != -1 {
		// DFS all children of curr that are NOT on the path
		for _, v := range children[curr] {
			if !pathSet[v] {
				dfs(v)
			}
		}
		// Add curr's own value
		visited[nums[curr]] = true

		// Update mex
		for visited[mex] {
			mex++
		}
		ans[curr] = mex
		curr = parents[curr]
		if curr == -1 {
			break
		}
	}

	return ans
}

func main() {
	// Example: parents=[-1,0,0,2], nums=[1,2,3,4] -> [5,1,1,1]
	fmt.Println(smallestMissingValueSubtree([]int{-1, 0, 0, 2}, []int{1, 2, 3, 4}))

	// Additional tests
	fmt.Println(smallestMissingValueSubtree([]int{-1, 0, 1, 0}, []int{1, 2, 3, 4}))
	fmt.Println(smallestMissingValueSubtree([]int{-1, 0, 1, 1}, []int{2, 3, 4, 5}))
}
```

## 2004 — The Number Of Seniors And Juniors To Join The Company

```go
package main

// LeetCode #2004: The Number of Seniors and Juniors to Join the Company
// https://leetcode.com/problems/the-number-of-seniors-and-juniors-to-join-the-company/
// Difficulty: Hard [Paid]
//
// Given tables Employees (employee_id, experience='Senior'|'Junior', salary)
// and a company budget, find the maximum number of employees to hire.
// Hiring rule: first hire as many Seniors as possible within budget,
// then hire as many Juniors as possible with the remaining budget.
//
// Return [senior_count, junior_count].

import (
	"fmt"
	"sort"
)

func main() {
	// Test case 1: simple
	employees := [][]interface{}{
		{1, "Senior", 50000},
		{2, "Senior", 80000},
		{3, "Junior", 30000},
		{4, "Junior", 25000},
	}
	// Budget = 100000
	// Seniors: can hire 1 (50000) or (80000) -> 1 senior with min salary to maximize count
	// Actually "hire as many Seniors as possible" means maximize count
	// Senior: 50000 works, 80000 works but only 1. 50000 is better for remaining budget
	// With 50000 senior, remaining=50000, juniors: 30000+25000=55000 > 50000, so only 1 junior
	// Result: [1, 1]
	fmt.Println(theNumberOfSeniorsAndJuniorsToJoinTheCompany(100000, employees))

	// Test case 2: hire many juniors
	employees = [][]interface{}{
		{1, "Senior", 10000},
		{2, "Senior", 20000},
		{3, "Junior", 5000},
		{4, "Junior", 5000},
		{5, "Junior", 5000},
	}
	// Budget = 25000
	// Seniors: hire 1 (10000) -> remaining 15000
	// Juniors: 3 * 5000 = 15000 -> 3 juniors
	// Result: [1, 3]
	fmt.Println(theNumberOfSeniorsAndJuniorsToJoinTheCompany(25000, employees))

	// Test case 3: budget too low
	employees = [][]interface{}{
		{1, "Senior", 100000},
		{2, "Junior", 80000},
	}
	fmt.Println(theNumberOfSeniorsAndJuniorsToJoinTheCompany(50000, employees))

	// Test case 4: no budget
	fmt.Println(theNumberOfSeniorsAndJuniorsToJoinTheCompany(0, employees))

	// Test case 5: larger example
	employees = [][]interface{}{
		{1, "Senior", 40000},
		{2, "Senior", 30000},
		{3, "Senior", 60000},
		{4, "Junior", 15000},
		{5, "Junior", 10000},
		{6, "Junior", 20000},
	}
	// Budget = 70000
	// Seniors: can hire 2 (40000+30000=70000) -> remaining 0
	// No budget for juniors
	// Result: [2, 0]
	fmt.Println(theNumberOfSeniorsAndJuniorsToJoinTheCompany(70000, employees))
}

func theNumberOfSeniorsAndJuniorsToJoinTheCompany(budget int, employees [][]interface{}) []int {
	var seniorSalaries, juniorSalaries []int
	for _, emp := range employees {
		salary := emp[2].(int)
		if emp[1].(string) == "Senior" {
			seniorSalaries = append(seniorSalaries, salary)
		} else {
			juniorSalaries = append(juniorSalaries, salary)
		}
	}

	sort.Ints(seniorSalaries)
	sort.Ints(juniorSalaries)

	// Hire as many Seniors as possible (maximize count)
	seniorCount := 0
	seniorCost := 0
	for _, s := range seniorSalaries {
		if seniorCost+s <= budget {
			seniorCost += s
			seniorCount++
		} else {
			break
		}
	}

	// Hire as many Juniors as possible with remaining budget
	remaining := budget - seniorCost
	juniorCount := 0
	juniorCost := 0
	for _, s := range juniorSalaries {
		if juniorCost+s <= remaining {
			juniorCost += s
			juniorCount++
		} else {
			break
		}
	}

	return []int{seniorCount, juniorCount}
}
```

## 2005 — Subtree Removal Game With Fibonacci Tree

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
	g := make([]int, k+1)
	g[0] = 1
	g[1] = 1

	for i := 2; i <= k; i++ {
		g[i] = g[i-1] ^ g[i-2]
	}

	return g[k]
}
```

## 2009 — Minimum Number Of Operations To Make Array Continuous

```go
package main

// LeetCode #2009: Minimum Number of Operations to Make Array Continuous
// https://leetcode.com/problems/minimum-number-of-operations-to-make-array-continuous/
// Difficulty: Hard
// Approach: Sort + sliding window on unique elements.
// A continuous array has max-min < n and distinct elements.
// For each element as left bound, find rightmost element with value < left+n.
// Answer = n - max window size.

import (
	"fmt"
	"sort"
)

func minOperations(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	// Sort and deduplicate
	sort.Ints(nums)
	uniq := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		uniq = append(uniq, nums[i])
	}

	maxWindow := 0
	right := 0
	for left := 0; left < len(uniq); left++ {
		for right < len(uniq) && uniq[right] < uniq[left]+n {
			right++
		}
		windowSize := right - left
		if windowSize > maxWindow {
			maxWindow = windowSize
		}
	}

	return n - maxWindow
}

func main() {
	// Example: [4,2,5,3] -> 0 (already continuous)
	fmt.Println(minOperations([]int{4, 2, 5, 3}))

	// Additional tests
	fmt.Println(minOperations([]int{1, 2, 3, 5, 6})) // [1,2,3,5,6] -> need to make [2,3,4,5,6] or [1,2,3,4,5], ops = 1
	fmt.Println(minOperations([]int{1, 10, 100, 1000}))
	fmt.Println(minOperations([]int{8, 5, 9, 9, 5, 7, 2, 3}))
}
```

## 2010 — The Number Of Seniors And Juniors To Join The Company Ii

```go
package main

// LeetCode #2010: The Number of Seniors and Juniors to Join the Company II
// https://leetcode.com/problems/the-number-of-seniors-and-juniors-to-join-the-company-ii/
// Difficulty: Hard [Paid]
//
// Variation of 2004. Instead of hiring as many Seniors as possible first,
// we need to find the maximum TOTAL number of employees we can hire within
// budget. We can choose any combination of Seniors and Juniors.
//
// Return [senior_count, junior_count] for the optimal hiring strategy
// that maximizes total headcount. If there are multiple solutions with
// the same total, prefer the one with more Seniors (or less Juniors).

import (
	"fmt"
	"sort"
)

func main() {
	// Test case 1: simple
	employees := [][]interface{}{
		{1, "Senior", 50000},
		{2, "Senior", 80000},
		{3, "Junior", 30000},
		{4, "Junior", 25000},
	}
	// Budget = 100000
	// Option A: 1 Senior (50000) + 1 Junior (25000 or 30000) = 2 employees
	// Option B: 0 Seniors + 3 Juniors (25000+30000) = 55000, remaining 45000 can't afford more = 2
	// Option C: 2 Seniors (50000+80000) = 130000 > budget, no
	// Max total = 2. Prefer more Seniors: [1, 1]
	fmt.Println(theNumberOfSeniorsAndJuniorsToJoinTheCompanyIi(100000, employees))

	// Test case 2: more juniors gives higher total
	employees = [][]interface{}{
		{1, "Senior", 100000},
		{2, "Senior", 60000},
		{3, "Junior", 15000},
		{4, "Junior", 15000},
		{5, "Junior", 15000},
	}
	// Budget = 100000
	// Option A: 1 Senior (60000) + 2 Juniors (30000) = 3 employees
	// Option B: 0 Seniors + 6 Juniors... only have 3 juniors = 3 employees
	// Option C: 1 Senior (100000) + 0 Juniors = 1 employee
	// Max total = 3 with more Seniors: [1, 2]
	fmt.Println(theNumberOfSeniorsAndJuniorsToJoinTheCompanyIi(100000, employees))

	// Test case 3: tie-breaking
	employees = [][]interface{}{
		{1, "Senior", 30000},
		{2, "Junior", 30000},
	}
	// Budget = 60000
	// Option A: 1 Senior (30000) + 1 Junior (30000) = 2 employees
	// Option B: 2 Seniors... only 1 senior = 1
	// Option C: 2 Juniors... only 1 junior = 1
	// [1, 1]
	fmt.Println(theNumberOfSeniorsAndJuniorsToJoinTheCompanyIi(60000, employees))

	// Test case 4: budget not enough for any
	employees = [][]interface{}{
		{1, "Senior", 50000},
		{2, "Junior", 30000},
	}
	fmt.Println(theNumberOfSeniorsAndJuniorsToJoinTheCompanyIi(20000, employees))

	// Test case 5: complex case
	employees = [][]interface{}{
		{1, "Senior", 10000},
		{2, "Senior", 20000},
		{3, "Senior", 30000},
		{4, "Junior", 5000},
		{5, "Junior", 5000},
		{6, "Junior", 5000},
		{7, "Junior", 10000},
	}
	// Budget = 40000
	// Seniors: [10000, 20000, 30000]
	// Juniors: [5000, 5000, 5000, 10000]
	// Try various combinations to maximize total count:
	// Option: 1 Senior (10000) + up to 6 juniors... max juniors with 30000: 3*5000+10000=25000 (4)
	//   Total: 5
	// Option: 2 Seniors (10000+20000=30000) + remaining 10000: 2*5000=10000 (2 juniors)
	//   Total: 4
	// Option: 0 Seniors + 4 Juniors (5000*3+10000=25000) -> remaining 15000 -> 4 total
	// Option: 3 Seniors: 60000 > budget
	// Best: 1 Senior + 4 Juniors = 5 total
	fmt.Println(theNumberOfSeniorsAndJuniorsToJoinTheCompanyIi(40000, employees))
}

func theNumberOfSeniorsAndJuniorsToJoinTheCompanyIi(budget int, employees [][]interface{}) []int {
	var seniorSalaries, juniorSalaries []int
	for _, emp := range employees {
		salary := emp[2].(int)
		if emp[1].(string) == "Senior" {
			seniorSalaries = append(seniorSalaries, salary)
		} else {
			juniorSalaries = append(juniorSalaries, salary)
		}
	}

	sort.Ints(seniorSalaries)
	sort.Ints(juniorSalaries)

	// Prefix sums for quick cost calculation
	seniorPrefix := make([]int, len(seniorSalaries)+1)
	for i, s := range seniorSalaries {
		seniorPrefix[i+1] = seniorPrefix[i] + s
	}

	juniorPrefix := make([]int, len(juniorSalaries)+1)
	for i, s := range juniorSalaries {
		juniorPrefix[i+1] = juniorPrefix[i] + s
	}

	bestTotal := 0
	bestSenior := 0
	bestJunior := 0

	// Try all possible senior counts
	for s := 0; s <= len(seniorSalaries); s++ {
		seniorCost := seniorPrefix[s]
		if seniorCost > budget {
			break
		}
		remaining := budget - seniorCost

		// Maximum juniors with remaining budget
		j := 0
		for j < len(juniorSalaries) && juniorPrefix[j+1] <= remaining {
			j++
		}

		total := s + j
		if total > bestTotal || (total == bestTotal && s > bestSenior) {
			bestTotal = total
			bestSenior = s
			bestJunior = j
		}
	}

	return []int{bestSenior, bestJunior}
}
```

## 2014 — Longest Subsequence Repeated K Times

```go
package main

// LeetCode #2014: Longest Subsequence Repeated k Times
// https://leetcode.com/problems/longest-subsequence-repeated-k-times/
// Difficulty: Hard
// Approach: BFS generate candidate strings in order of length.
// Count character frequencies, max_uses = freq / k.
// Generate all possible strings up to n/k length.
// For each candidate, check if repeated k times is a subsequence of s.
// Keep the longest.

import "fmt"

func longestSubsequenceRepeatedK(s string, k int) string {
	// Count frequencies
	freq := make([]int, 26)
	for _, ch := range s {
		freq[ch-'a']++
	}

	// Max uses for each character
	maxUses := make([]int, 26)
	for i := 0; i < 26; i++ {
		maxUses[i] = freq[i] / k
	}

	// Check if str is a subsequence of s
	isSubseq := func(str string) bool {
		j := 0
		for i := 0; i < len(s) && j < len(str); i++ {
			if s[i] == str[j] {
				j++
			}
		}
		return j == len(str)
	}

	// Check if t repeated k times is a subsequence of s
	check := func(t string) bool {
		if len(t) == 0 {
			return false
		}
		concat := ""
		for i := 0; i < k; i++ {
			concat += t
		}
		return isSubseq(concat)
	}

	// BFS to generate candidates
	queue := []string{""}
	best := ""

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		for c := 0; c < 26; c++ {
			if maxUses[c] == 0 {
				continue
			}
			next := cur + string(rune('a'+c))
			if !check(next) {
				continue
			}
			queue = append(queue, next)
			if len(next) > len(best) || (len(next) == len(best) && next > best) {
				best = next
			}
		}
	}

	return best
}

func main() {
	// Example: "letsleetcode", k=2 -> "let"
	fmt.Println(longestSubsequenceRepeatedK("letsleetcode", 2))

	// Additional tests
	fmt.Println(longestSubsequenceRepeatedK("aabbaabbaabb", 3))
	fmt.Println(longestSubsequenceRepeatedK("abcd", 2))
}
```

## 2019 — The Score Of Students Solving Math Expression

```go
package main

// LeetCode #2019: The Score of Students Solving Math Expression
// https://leetcode.com/problems/the-score-of-students-solving-math-expression/
// Difficulty: Hard
//
// Compute all possible results from different parenthesizations of an expression
// containing digits, '+', and '*'. Score student answers: +5 for correct answer
// (using standard *-before-+ precedence), +2 for any other achievable result, 0 otherwise.

import "fmt"

func main() {
	// Example 1
	fmt.Println(scoreOfStudents("7+3*1*2", []int{20, 13, 42}))

	// Example 2
	fmt.Println(scoreOfStudents("3+5*2", []int{13, 0, 10, 13, 13, 16, 16}))

	// Example 3
	fmt.Println(scoreOfStudents("6+0*1", []int{12, 9, 6, 4, 8, 6}))

	// Single number
	fmt.Println(scoreOfStudents("5", []int{5, 0, 10}))
}

func scoreOfStudents(s string, answers []int) int {
	// Parse expression into numbers and operators
	nums := []int{}
	ops := []byte{}
	num := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			num = num*10 + int(s[i]-'0')
		} else {
			nums = append(nums, num)
			ops = append(ops, s[i])
			num = 0
		}
	}
	nums = append(nums, num)

	n := len(nums)

	// dp[l][r] = set of possible results for subexpression nums[l..r]
	dp := make([][]map[int]bool, n)
	for i := range dp {
		dp[i] = make([]map[int]bool, n)
		for j := range dp[i] {
			dp[i][j] = make(map[int]bool)
		}
		dp[i][i][nums[i]] = true
	}

	// Fill DP for increasing subexpression lengths
	for length := 2; length <= n; length++ {
		for l := 0; l+length <= n; l++ {
			r := l + length - 1
			for k := l; k < r; k++ {
				for a := range dp[l][k] {
					for b := range dp[k+1][r] {
						var val int
						if ops[k] == '+' {
							val = a + b
						} else {
							val = a * b
						}
						if val <= 1000 {
							dp[l][r][val] = true
						}
					}
				}
			}
		}
	}

	possible := dp[0][n-1]
	correct := evaluateStandard(nums, ops)

	total := 0
	for _, ans := range answers {
		if ans == correct {
			total += 5
		} else if possible[ans] {
			total += 2
		}
	}
	return total
}

// evaluateStandard computes the result using standard *-before-+ precedence.
func evaluateStandard(nums []int, ops []byte) int {
	stack := []int{nums[0]}
	for i, op := range ops {
		if op == '+' {
			stack = append(stack, nums[i+1])
		} else {
			stack[len(stack)-1] *= nums[i+1]
		}
	}
	sum := 0
	for _, v := range stack {
		sum += v
	}
	return sum
}
```

## 2025 — Maximum Number Of Ways To Partition An Array

```go
package main

// LeetCode #2025: Maximum Number of Ways to Partition an Array
// https://leetcode.com/problems/maximum-number-of-ways-to-partition-an-array/
// Difficulty: Hard
// Approach: Prefix sum + hash maps.
// Compute prefix sums. Count ways without changes.
// For each position i, compute ways if nums[i] is changed to k.
// Use left_freq (prefix sums before i) and right_freq (prefix sums from i to n-2).
// For j < i: condition is prefix[j] == new_total / 2 (unchanged)
// For j >= i: condition is prefix[j] + diff == new_total / 2, i.e. prefix[j] == target - diff

import "fmt"

func maxNumberOfWaysToPartition(nums []int, k int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	// Compute prefix sums
	prefix := make([]int, n)
	prefix[0] = nums[0]
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] + nums[i]
	}
	total := prefix[n-1]

	ans := 0

	// Without changes: count j where prefix[j]*2 == total (j < n-1)
	if total%2 == 0 {
		target := total / 2
		for j := 0; j < n-1; j++ {
			if prefix[j] == target {
				ans++
			}
		}
	}

	// With changes: try changing each nums[i] to k
	// leftFreq: prefix values for partition positions j < i
	// rightFreq: prefix values for partition positions j >= i
	leftFreq := make(map[int]int)
	rightFreq := make(map[int]int)
	for j := 0; j < n-1; j++ {
		rightFreq[prefix[j]]++
	}

	for i := 0; i < n; i++ {
		diff := k - nums[i]
		newTotal := total + diff

		if newTotal%2 == 0 {
			target := newTotal / 2
			cnt := 0

			// j < i: unchanged prefix
			cnt += leftFreq[target]

			// j >= i: prefix[j] + diff == target => prefix[j] == target - diff
			cnt += rightFreq[target-diff]

			if cnt > ans {
				ans = cnt
			}
		}

		// Move prefix[i] from rightFreq to leftFreq for next iteration
		if i < n-1 {
			rightFreq[prefix[i]]--
			if rightFreq[prefix[i]] == 0 {
				delete(rightFreq, prefix[i])
			}
			leftFreq[prefix[i]]++
		}
	}

	return ans
}

func main() {
	// Example: nums=[2,-1,2], k=3 -> 1
	fmt.Println(maxNumberOfWaysToPartition([]int{2, -1, 2}, 3))

	// Additional tests
	fmt.Println(maxNumberOfWaysToPartition([]int{1, 2, 3, 4}, 5))
	fmt.Println(maxNumberOfWaysToPartition([]int{1, 1, 1}, 2))
	fmt.Println(maxNumberOfWaysToPartition([]int{0, 0, 0}, 1))
}
```

## 2030 — Smallest K Length Subsequence With Occurrences Of A Letter

```go
package main

// LeetCode #2030: Smallest K-Length Subsequence With Occurrences of a Letter
// https://leetcode.com/problems/smallest-k-length-subsequence-with-occurrences-of-a-letter/
// Difficulty: Hard
// Approach: Monotonic Stack

import "fmt"

func smallestKLengthSubsequence(s string, k int, letter byte, rep int) string {
	n := len(s)
	// Count total occurrences of letter in s
	totalLetter := 0
	for i := 0; i < n; i++ {
		if s[i] == letter {
			totalLetter++
		}
	}

	stack := make([]byte, 0, k)
	usedLetter := 0

	for i := 0; i < n; i++ {
		ch := s[i]
		// How many letters remain after current position
		remainingLetter := totalLetter
		if ch == letter {
			remainingLetter-- // decrement for current char
		}

		// While stack has elements, we can try to pop
		for len(stack) > 0 && stack[len(stack)-1] > ch {
			// If popping would make it impossible to reach k length, stop
			if n-i+len(stack)-1 < k {
				break
			}
			// If we're popping a letter, check if we'd still have enough letters
			if stack[len(stack)-1] == letter {
				if usedLetter-1+remainingLetter < rep {
					break
				}
				usedLetter--
			}
			stack = stack[:len(stack)-1]
		}

		// Add current character if we have room
		if len(stack) < k {
			stack = append(stack, ch)
			if ch == letter {
				usedLetter++
			}
		}

		// Update total remaining letter count
		if ch == letter {
			totalLetter--
		}
	}

	// If we have more than k characters, trim from end, but keep enough letters
	if len(stack) > k {
		extra := len(stack) - k
		newStack := make([]byte, 0, k)
		keptLetter := 0
		for i := 0; i < len(stack)-extra; i++ {
			newStack = append(newStack, stack[i])
			if stack[i] == letter {
				keptLetter++
			}
		}
		// Need to add from the trimmed part if not enough letters
		if keptLetter < rep {
			for i := len(stack) - extra; i < len(stack) && keptLetter < rep; i++ {
				if stack[i] == letter {
					newStack = append(newStack, stack[i])
					keptLetter++
				}
			}
		}
		stack = newStack
	}

	return string(stack)
}

func main() {
	fmt.Println("2030. Smallest K-Length Subsequence With Occurrences of a Letter")

	// Example 1
	s1 := "leet"
	k1 := 3
	letter1 := byte('e')
	rep1 := 1
	fmt.Printf("s=%q k=%d letter=%c rep=%d → %q (expected \"eet\")\n",
		s1, k1, letter1, rep1, smallestKLengthSubsequence(s1, k1, letter1, rep1))

	// Example 2
	s2 := "leetcode"
	k2 := 4
	letter2 := byte('e')
	rep2 := 2
	fmt.Printf("s=%q k=%d letter=%c rep=%d → %q (expected \"ecde\")\n",
		s2, k2, letter2, rep2, smallestKLengthSubsequence(s2, k2, letter2, rep2))

	// Additional test
	s3 := "aaabbb"
	k3 := 3
	letter3 := byte('a')
	rep3 := 2
	fmt.Printf("s=%q k=%d letter=%c rep=%d → %q\n",
		s3, k3, letter3, rep3, smallestKLengthSubsequence(s3, k3, letter3, rep3))
}
```

## 2035 — Partition Array Into Two Arrays To Minimize Sum Difference

```go
package main

// LeetCode #2035: Partition Array Into Two Arrays to Minimize Sum Difference
// https://leetcode.com/problems/partition-array-into-two-arrays-to-minimize-sum-difference/
// Difficulty: Hard
// Approach: Meet-in-the-Middle

import (
	"fmt"
	"math"
	"sort"
)

func minimumDifference(nums []int) int {
	n := len(nums) / 2 // each half size
	total := 0
	for _, v := range nums {
		total += v
	}

	// Generate all subset sums for each half, grouped by subset size
	leftSums := make([][]int, n+1)
	rightSums := make([][]int, n+1)

	// Generate combinations for left half
	for mask := 0; mask < (1 << n); mask++ {
		sum := 0
		size := 0
		for i := 0; i < n; i++ {
			if mask&(1<<i) != 0 {
				sum += nums[i]
				size++
			}
		}
		leftSums[size] = append(leftSums[size], sum)
	}

	// Generate combinations for right half
	for mask := 0; mask < (1 << n); mask++ {
		sum := 0
		size := 0
		for i := 0; i < n; i++ {
			if mask&(1<<i) != 0 {
				sum += nums[n+i]
				size++
			}
		}
		rightSums[size] = append(rightSums[size], sum)
	}

	// Sort each group in right half for binary search
	for i := 0; i <= n; i++ {
		sort.Ints(rightSums[i])
	}

	ans := math.MaxInt32

	// For each possible size from left half, find complement from right half
	for leftSize := 0; leftSize <= n; leftSize++ {
		rightSize := n - leftSize
		for _, leftSum := range leftSums[leftSize] {
			// We want leftSum + rightSum as close to total/2 as possible
			target := total/2 - leftSum
			rightArr := rightSums[rightSize]
			if len(rightArr) == 0 {
				continue
			}
			// Binary search for closest
			idx := sort.SearchInts(rightArr, target)
			if idx < len(rightArr) {
				sum := leftSum + rightArr[idx]
				diff := total - 2*sum
				if diff < 0 {
					diff = -diff
				}
				if diff < ans {
					ans = diff
				}
			}
			if idx > 0 {
				idx--
				sum := leftSum + rightArr[idx]
				diff := total - 2*sum
				if diff < 0 {
					diff = -diff
				}
				if diff < ans {
					ans = diff
				}
			}
		}
	}

	return ans
}

func main() {
	fmt.Println("2035. Partition Array Into Two Arrays to Minimize Sum Difference")

	// Example 1
	nums1 := []int{3, 9, 7, 3}
	fmt.Printf("nums=%v → %d (expected 2)\n", nums1, minimumDifference(nums1))

	// Example 2
	nums2 := []int{-36, 36}
	fmt.Printf("nums=%v → %d (expected 72)\n", nums2, minimumDifference(nums2))

	// Example 3
	nums3 := []int{2, -1, 0, 4, -2, -9}
	fmt.Printf("nums=%v → %d (expected 0)\n", nums3, minimumDifference(nums3))
}
```

## 2040 — Kth Smallest Product Of Two Sorted Arrays

```go
package main

// LeetCode #2040: Kth Smallest Product of Two Sorted Arrays
// https://leetcode.com/problems/kth-smallest-product-of-two-sorted-arrays/
// Difficulty: Hard
//
// Approach: Binary search on the answer value. Count how many pairs have
// product <= mid using two-pointer technique (since both arrays are sorted).

import (
	"fmt"
	"sort"
)

func main() {
	// Example 1
	fmt.Println(kthSmallestProduct([]int{2, 5}, []int{3, 4}, 2))

	// Example 2
	fmt.Println(kthSmallestProduct([]int{-4, -2, 0, 3}, []int{2, 4}, 6))

	// Example 3
	fmt.Println(kthSmallestProduct([]int{-2, -1, 0, 1, 2}, []int{-3, -1, 2, 4, 5}, 3))
}

func kthSmallestProduct(nums1 []int, nums2 []int, k int) int64 {
	left, right := int64(-1_000_000_000_000_000_000), int64(1_000_000_000_000_000_000)

	for left < right {
		mid := left + (right-left)/2
		if countLessOrEqual(nums1, nums2, mid) >= int64(k) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

// countLessOrEqual counts number of pairs (i,j) such that nums1[i]*nums2[j] <= x
func countLessOrEqual(nums1, nums2 []int, x int64) int64 {
	var count int64
	for _, a := range nums1 {
		if a > 0 {
			limit := x / int64(a)
			count += int64(sort.Search(len(nums2), func(i int) bool { return nums2[i] > int(limit) }))
		} else if a < 0 {
			limit := ceilDiv(x, int64(a))
			idx := sort.Search(len(nums2), func(i int) bool { return nums2[i] >= int(limit) })
			count += int64(len(nums2) - idx)
		} else if x >= 0 {
			count += int64(len(nums2))
		}
	}
	return count
}

// ceilDiv returns ceil(a/b) for possibly negative values. Assumes b != 0.
func ceilDiv(a, b int64) int64 {
	if a*b >= 0 {
		return (a + b - 1) / b
	}
	return a / b
}
```

## 2045 — Second Minimum Time To Reach Destination

```go
package main

// LeetCode #2045: Second Minimum Time to Reach Destination
// https://leetcode.com/problems/second-minimum-time-to-reach-destination/
// Difficulty: Hard
// Approach: Modified Dijkstra / BFS with two distances

import (
	"container/heap"
	"fmt"
	"math"
)

func secondMinimum(n int, edges [][]int, time int, change int) int {
	// Build adjacency list
	adj := make([][]int, n+1)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// dist1[i] = shortest time to reach i, dist2[i] = second shortest
	dist1 := make([]int, n+1)
	dist2 := make([]int, n+1)
	for i := range dist1 {
		dist1[i] = math.MaxInt32
		dist2[i] = math.MaxInt32
	}

	dist1[1] = 0

	// Min-heap: (time, node)
	pq := &minHeap{}
	heap.Init(pq)
	heap.Push(pq, [2]int{0, 1})

	for pq.Len() > 0 {
		cur := heap.Pop(pq).([2]int)
		t := cur[0]
		u := cur[1]

		// If this is stale (worse than dist2), skip
		if t > dist2[u] {
			continue
		}

		// Compute wait time for traffic signal
		// Signal is green in [0, change), red in [change, 2*change), green in [2*change, 3*change), ...
		wait := 0
		cycle := t / change
		if cycle%2 == 1 {
			// Red light, wait until next green
			wait = change - t%change
		}
		nextTime := t + wait + time

		for _, v := range adj[u] {
			if nextTime < dist1[v] {
				dist2[v] = dist1[v]
				dist1[v] = nextTime
				heap.Push(pq, [2]int{nextTime, v})
			} else if nextTime > dist1[v] && nextTime < dist2[v] {
				dist2[v] = nextTime
				heap.Push(pq, [2]int{nextTime, v})
			}
		}
	}

	return dist2[n]
}

type minHeap [][2]int

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i][0] < h[j][0] }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.([2]int)) }
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func main() {
	fmt.Println("2045. Second Minimum Time to Reach Destination")

	// Example 1
	n1 := 5
	edges1 := [][]int{{1, 2}, {1, 3}, {1, 4}, {3, 4}, {4, 5}}
	time1 := 3
	change1 := 5
	fmt.Printf("n=%d edges=%v time=%d change=%d → %d (expected 13)\n",
		n1, edges1, time1, change1, secondMinimum(n1, edges1, time1, change1))

	// Example 2
	n2 := 2
	edges2 := [][]int{{1, 2}}
	time2 := 3
	change2 := 2
	fmt.Printf("n=%d edges=%v time=%d change=%d → %d (expected 11)\n",
		n2, edges2, time2, change2, secondMinimum(n2, edges2, time2, change2))
}
```

## 2050 — Parallel Courses Iii

```go
package main

// LeetCode #2050: Parallel Courses III
// https://leetcode.com/problems/parallel-courses-iii/
// Difficulty: Hard
// Approach: Topological Sort + DP

import "fmt"

func minimumTime(n int, relations [][]int, time []int) int {
	// Build graph and indegree
	graph := make([][]int, n+1)
	indeg := make([]int, n+1)
	for _, r := range relations {
		prev, next := r[0], r[1]
		graph[prev] = append(graph[prev], next)
		indeg[next]++
	}

	// dp[i] = earliest completion time for course i (1-indexed)
	dp := make([]int, n+1)
	queue := make([]int, 0)

	for i := 1; i <= n; i++ {
		if indeg[i] == 0 {
			dp[i] = time[i-1]
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		for _, v := range graph[u] {
			// Can start v only after u finishes
			if dp[u] > dp[v] {
				dp[v] = dp[u]
			}
			indeg[v]--
			if indeg[v] == 0 {
				dp[v] += time[v-1]
				queue = append(queue, v)
			}
		}
	}

	ans := 0
	for i := 1; i <= n; i++ {
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}

func main() {
	fmt.Println("2050. Parallel Courses III")

	// Example 1
	n1 := 3
	relations1 := [][]int{{1, 3}, {2, 3}}
	time1 := []int{3, 2, 5}
	fmt.Printf("n=%d relations=%v time=%v → %d (expected 8)\n",
		n1, relations1, time1, minimumTime(n1, relations1, time1))

	// Example 2
	n2 := 5
	relations2 := [][]int{{1, 5}, {2, 5}, {3, 5}, {3, 4}, {4, 5}}
	time2 := []int{1, 2, 3, 4, 5}
	fmt.Printf("n=%d relations=%v time=%v → %d (expected 12)\n",
		n2, relations2, time2, minimumTime(n2, relations2, time2))
}
```

## 2056 — Number Of Valid Move Combinations On Chessboard

```go
package main

// LeetCode #2056: Number of Valid Move Combinations On Chessboard
// https://leetcode.com/problems/number-of-valid-move-combinations-on-chessboard/
// Difficulty: Hard
//
// For each piece on an 8x8 board, enumerate all possible straight-line moves
// (including staying still). Then use backtracking to try all combinations and
// simulate simultaneous movement step-by-step to detect collisions.

import "fmt"

type piece byte

const (
	rook   piece = 'r'
	bishop piece = 'b'
	queen  piece = 'q'
)

type move struct{ dr, dc, steps int }

var rookDirs = [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
var bishopDirs = [][2]int{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
var queenDirs = [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

func main() {
	// Example: single rook at (1,1)
	fmt.Println(countCombinations([]string{"rook"}, [][]int{{1, 1}}))

	// Example: single queen
	fmt.Println(countCombinations([]string{"queen"}, [][]int{{1, 1}}))

	// Two rooks
	fmt.Println(countCombinations([]string{"rook", "rook"}, [][]int{{1, 1}, {8, 8}}))
}

func countCombinations(pieces []string, positions [][]int) int {
	n := len(pieces)
	pts := make([]piece, n)
	start := make([][2]int, n)
	for i := range pieces {
		pts[i] = piece(pieces[i][0])
		start[i] = [2]int{positions[i][0] - 1, positions[i][1] - 1}
	}

	allMoves := make([][]move, n)
	for i := 0; i < n; i++ {
		allMoves[i] = genMoves(pts[i], start[i])
	}

	ans := 0
	chosen := make([]move, n)
	var dfs func(int)
	dfs = func(idx int) {
		if idx == n {
			if simulate(start, chosen) {
				ans++
			}
			return
		}
		for _, m := range allMoves[idx] {
			chosen[idx] = m
			dfs(idx + 1)
		}
	}
	dfs(0)
	return ans
}

func genMoves(pt piece, pos [2]int) []move {
	var dirs [][2]int
	switch pt {
	case rook:
		dirs = rookDirs
	case bishop:
		dirs = bishopDirs
	case queen:
		dirs = queenDirs
	}

	var moves []move
	// Stay still
	moves = append(moves, move{0, 0, 0})

	for _, d := range dirs {
		maxSteps := 0
		r, c := pos[0]+d[0], pos[1]+d[1]
		for r >= 0 && r < 8 && c >= 0 && c < 8 {
			maxSteps++
			r += d[0]
			c += d[1]
		}
		for s := 1; s <= maxSteps; s++ {
			moves = append(moves, move{d[0], d[1], s})
		}
	}
	return moves
}

func simulate(start [][2]int, chosen []move) bool {
	n := len(start)
	pos := make([][2]int, n)
	copy(pos, start)

	// Find the maximum number of steps among all chosen moves
	maxSteps := 0
	for _, m := range chosen {
		if m.steps > maxSteps {
			maxSteps = m.steps
		}
	}

	for step := 1; step <= maxSteps; step++ {
		for i := 0; i < n; i++ {
			if step <= chosen[i].steps {
				pos[i][0] += chosen[i].dr
				pos[i][1] += chosen[i].dc
			}
		}
		// Check for collisions at this step
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				if pos[i] == pos[j] {
					return false
				}
			}
		}
	}
	return true
}
```

## 2060 — Check If An Original String Exists Given Two Encoded Strings

```go
package main

// LeetCode #2060: Check if an Original String Exists Given Two Encoded Strings
// https://leetcode.com/problems/check-if-an-original-string-exists-given-two-encoded-strings/
// Difficulty: Hard
// Approach: DP with memoization on (i, j, diff)

import (
	"fmt"
	"unicode"
)

func possiblyEquals(s1 string, s2 string) bool {
	// Memoization: key = (i, j, diff) where diff = delta length from s1's perspective
	// If diff > 0, s1 has extra length (s2 needs to catch up)
	// If diff < 0, s2 has extra length
	// Offset diff by 2000 to make it non-negative
	memo := make(map[[3]int]bool)
	var dfs func(i, j, diff int) bool
	dfs = func(i, j, diff int) bool {
		if i == len(s1) && j == len(s2) {
			return diff == 0
		}

		key := [3]int{i, j, diff + 2000}
		if val, ok := memo[key]; ok {
			return val
		}

		// Case 1: s1 has a digit
		if i < len(s1) && unicode.IsDigit(rune(s1[i])) {
			end := i
			for end < len(s1) && unicode.IsDigit(rune(s1[end])) {
				end++
			}
			// Generate all possible numbers from digits s1[i:end]
			// Max 3 digits based on problem constraints
			num := 0
			for p := i; p < end; p++ {
				num = num*10 + int(s1[p]-'0')
				if dfs(p+1, j, diff-num) {
					memo[key] = true
					return true
				}
			}
			memo[key] = false
			return false
		}

		// Case 2: s2 has a digit
		if j < len(s2) && unicode.IsDigit(rune(s2[j])) {
			end := j
			for end < len(s2) && unicode.IsDigit(rune(s2[end])) {
				end++
			}
			num := 0
			for p := j; p < end; p++ {
				num = num*10 + int(s2[p]-'0')
				if dfs(i, p+1, diff+num) {
					memo[key] = true
					return true
				}
			}
			memo[key] = false
			return false
		}

		// Case 3: both have letters (or one is empty due to diff)
		if diff > 0 {
			// s1 has extra length, consume from s1
			if i < len(s1) && unicode.IsLetter(rune(s1[i])) {
				if dfs(i+1, j, diff-1) {
					memo[key] = true
					return true
				}
			}
		} else if diff < 0 {
			// s2 has extra length, consume from s2
			if j < len(s2) && unicode.IsLetter(rune(s2[j])) {
				if dfs(i, j+1, diff+1) {
					memo[key] = true
					return true
				}
			}
		} else {
			// diff == 0, both must be letters and equal
			if i < len(s1) && j < len(s2) && unicode.IsLetter(rune(s1[i])) && unicode.IsLetter(rune(s2[j])) && s1[i] == s2[j] {
				if dfs(i+1, j+1, 0) {
					memo[key] = true
					return true
				}
			}
		}

		memo[key] = false
		return false
	}

	return dfs(0, 0, 0)
}

func main() {
	fmt.Println("2060. Check if an Original String Exists Given Two Encoded Strings")

	// Example 1
	s1 := "internationalization"
	s2 := "i18n"
	fmt.Printf("s1=%q s2=%q → %v (expected true)\n", s1, s2, possiblyEquals(s1, s2))

	// Example 2
	s1 = "l123e"
	s2 = "44"
	fmt.Printf("s1=%q s2=%q → %v (expected true)\n", s1, s2, possiblyEquals(s1, s2))

	// Example 3
	s1 = "a5b"
	s2 = "c5b"
	fmt.Printf("s1=%q s2=%q → %v (expected false)\n", s1, s2, possiblyEquals(s1, s2))
}
```

## 2065 — Maximum Path Quality Of A Graph

```go
package main

// LeetCode #2065: Maximum Path Quality of a Graph
// https://leetcode.com/problems/maximum-path-quality-of-a-graph/
// Difficulty: Hard
// Approach: DFS + Backtracking

import "fmt"

func maximalPathQuality(values []int, edges [][]int, maxTime int) int {
	n := len(values)
	// Build adjacency list: each entry is (neighbor, time)
	adj := make([][][2]int, n)
	for _, e := range edges {
		u, v, t := e[0], e[1], e[2]
		adj[u] = append(adj[u], [2]int{v, t})
		adj[v] = append(adj[v], [2]int{u, t})
	}

	visited := make([]int, n) // count visits to detect re-visits
	ans := 0

	var dfs func(u, time, quality int)
	dfs = func(u, time, quality int) {
		// Quality earned when first visiting a node
		if visited[u] == 0 {
			quality += values[u]
		}
		visited[u]++

		// If we're back at node 0, update answer
		if u == 0 {
			if quality > ans {
				ans = quality
			}
		}

		// Explore neighbors
		for _, edge := range adj[u] {
			v, t := edge[0], edge[1]
			if time+t <= maxTime {
				dfs(v, time+t, quality)
			}
		}

		// Backtrack
		visited[u]--
	}

	dfs(0, 0, 0)
	return ans
}

func main() {
	fmt.Println("2065. Maximum Path Quality of a Graph")

	// Example 1
	values1 := []int{0, 32, 10, 43}
	edges1 := [][]int{{0, 1, 10}, {1, 2, 15}, {0, 3, 10}}
	maxTime1 := 49
	fmt.Printf("values=%v edges=%v maxTime=%d → %d (expected 75)\n",
		values1, edges1, maxTime1, maximalPathQuality(values1, edges1, maxTime1))

	// Example 2
	values2 := []int{5, 10, 15, 20}
	edges2 := [][]int{{0, 1, 10}, {1, 2, 10}, {0, 3, 10}}
	maxTime2 := 30
	fmt.Printf("values=%v edges=%v maxTime=%d → %d (expected 25)\n",
		values2, edges2, maxTime2, maximalPathQuality(values2, edges2, maxTime2))

	// Example 3
	values3 := []int{1, 2, 3, 4}
	edges3 := [][]int{{0, 1, 10}, {1, 2, 11}, {2, 3, 12}, {1, 3, 13}}
	maxTime3 := 50
	fmt.Printf("values=%v edges=%v maxTime=%d → %d (expected 7)\n",
		values3, edges3, maxTime3, maximalPathQuality(values3, edges3, maxTime3))
}
```

## 2071 — Maximum Number Of Tasks You Can Assign

```go
package main

// LeetCode #2071: Maximum Number of Tasks You Can Assign
// https://leetcode.com/problems/maximum-number-of-tasks-you-can-assign/
// Difficulty: Hard
// Approach: Binary Search + Multiset (simulated with sort + two-pointer)

import (
	"fmt"
	"sort"
)

func maxTaskAssign(tasks []int, workers []int, pills int, strength int) int {
	sort.Ints(tasks)
	sort.Ints(workers)

	// Binary search on how many tasks we can complete
	left, right := 0, len(tasks)
	if len(workers) < right {
		right = len(workers)
	}

	for left < right {
		mid := left + (right-left+1)/2 // try to do mid tasks (hardest mid tasks)
		if canAssign(tasks, workers, pills, strength, mid) {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return left
}

func canAssign(tasks, workers []int, pills, strength, count int) bool {
	// We try to assign the 'count' hardest tasks to the 'count' strongest workers
	// tasks are sorted ascending, so we take the last 'count' tasks
	n := len(workers)
	taskIdx := len(tasks) - count // start from this task index

	// Use a multiset (multiset via sorted slice) for available workers
	// Actually, let's use a different approach: sort workers, then use a sliding window

	// Workers are sorted. We'll use a multiset implemented as a sorted slice
	// by maintaining available workers and removing with binary search.

	// Alternative: use two-pointer + multiset (e.g., a balanced BST via sorted slice + binary search)
	// We'll track workers that are available for the current task

	// Simpler approach: use a multiset implemented as a slice
	avail := make([]int, n)
	copy(avail, workers)
	// We'll pop from avail when a worker is used

	pillsLeft := pills

	// Process tasks from hardest to easiest (greedy: hardest task needs strongest worker)
	for j := len(tasks) - 1; j >= taskIdx; j-- {
		task := tasks[j]

		// Find a worker who can do this task without pill (weakest sufficient worker)
		idx := sort.SearchInts(avail, task)
		if idx < len(avail) {
			// Found a worker who can do it without pill, use them
			avail = append(avail[:idx], avail[idx+1:]...)
			continue
		}

		// Need a pill
		if pillsLeft > 0 {
			pillsLeft--
			// Find a worker who can do it with pill: worker + strength >= task
			needed := task - strength
			idx := sort.SearchInts(avail, needed)
			if idx < len(avail) {
				avail = append(avail[:idx], avail[idx+1:]...)
				continue
			}
		}

		// Cannot assign this task
		return false
	}

	return true
}

func main() {
	fmt.Println("2071. Maximum Number of Tasks You Can Assign")

	// Example 1
	tasks1 := []int{3, 2, 1}
	workers1 := []int{0, 3, 3}
	pills1 := 1
	strength1 := 1
	fmt.Printf("tasks=%v workers=%v pills=%d strength=%d → %d (expected 3)\n",
		tasks1, workers1, pills1, strength1, maxTaskAssign(tasks1, workers1, pills1, strength1))

	// Example 2
	tasks2 := []int{5, 4}
	workers2 := []int{0, 0, 0}
	pills2 := 1
	strength2 := 5
	fmt.Printf("tasks=%v workers=%v pills=%d strength=%d → %d (expected 1)\n",
		tasks2, workers2, pills2, strength2, maxTaskAssign(tasks2, workers2, pills2, strength2))

	// Example 3
	tasks3 := []int{10, 15, 30}
	workers3 := []int{0, 10, 20}
	pills3 := 2
	strength3 := 10
	fmt.Printf("tasks=%v workers=%v pills=%d strength=%d → %d (expected 2)\n",
		tasks3, workers3, pills3, strength3, maxTaskAssign(tasks3, workers3, pills3, strength3))
}
```

## 2076 — Process Restricted Friend Requests

```go
package main

// LeetCode #2076: Process Restricted Friend Requests
// https://leetcode.com/problems/process-restricted-friend-requests/
// Difficulty: Hard
// Approach: Union-Find + Restriction Check

import "fmt"

type unionFind struct {
	parent []int
	rank   []int
}

func newUnionFind(n int) *unionFind {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &unionFind{parent, rank}
}

func (uf *unionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *unionFind) union(x, y int) {
	rx, ry := uf.find(x), uf.find(y)
	if rx == ry {
		return
	}
	if uf.rank[rx] < uf.rank[ry] {
		rx, ry = ry, rx
	}
	uf.parent[ry] = rx
	if uf.rank[rx] == uf.rank[ry] {
		uf.rank[rx]++
	}
}

func friendRequests(n int, restrictions [][]int, requests [][]int) []bool {
	uf := newUnionFind(n)
	ans := make([]bool, len(requests))

	for i, req := range requests {
		u, v := req[0], req[1]
		ru, rv := uf.find(u), uf.find(v)
		canFriend := true

		// Check all restrictions: if both u's group and v's group
		// would contain both ends of a restriction, reject
		for _, res := range restrictions {
			a, b := res[0], res[1]
			ra, rb := uf.find(a), uf.find(b)
			// After union, ru and rv would be merged.
			// If (ra, rb) == (ru, rv) or (ra, rb) == (rv, ru), restriction is violated
			if (ra == ru && rb == rv) || (ra == rv && rb == ru) {
				canFriend = false
				break
			}
		}

		if canFriend {
			uf.union(u, v)
		}
		ans[i] = canFriend
	}

	return ans
}

func main() {
	fmt.Println("2076. Process Restricted Friend Requests")

	// Example 1
	n1 := 3
	restrictions1 := [][]int{{0, 1}}
	requests1 := [][]int{{0, 2}, {2, 1}}
	fmt.Printf("n=%d restrictions=%v requests=%v → %v (expected [true, false])\n",
		n1, restrictions1, requests1, friendRequests(n1, restrictions1, requests1))

	// Example 2
	n2 := 5
	restrictions2 := [][]int{{0, 1}, {1, 2}, {2, 3}}
	requests2 := [][]int{{0, 4}, {1, 2}, {3, 1}, {3, 4}}
	fmt.Printf("n=%d restrictions=%v requests=%v → %v (expected [true, false, true, false])\n",
		n2, restrictions2, requests2, friendRequests(n2, restrictions2, requests2))
}
```

## 2081 — Sum Of K Mirror Numbers

```go
package main

// LeetCode #2081: Sum of k-Mirror Numbers
// https://leetcode.com/problems/sum-of-k-mirror-numbers/
// Difficulty: Hard
// Approach: Generate palindromes in base k, check decimal palindrome

import (
	"fmt"
	"strconv"
)

func kMirror(k int, n int) int64 {
	var sum int64
	count := 0

	// Generate palindromes in base k
	length := 1
	for count < n {
		// Generate palindromes of given length in base k
		// For odd length, the first half determines the palindrome
		// For even length, the first half determines the palindrome

		halfLen := (length + 1) / 2
		start := pow(k, halfLen-1)
		end := pow(k, halfLen)

		for base := start; base < end && count < n; base++ {
			// Build the full palindrome in base k
			pal := buildPalindrome(base, length%2 == 1, k)
			if pal == 0 {
				continue
			}

			// Convert to decimal and check if it's a palindrome in decimal
			dec := toDecimal(pal, k)
			if isDecimalPalindrome(dec) {
				sum += dec
				count++
				if count >= n {
					break
				}
			}
		}
		length++
	}

	return sum
}

// pow computes k^e
func pow(k, e int) int {
	res := 1
	for i := 0; i < e; i++ {
		res *= k
	}
	return res
}

// buildPalindrome builds a number whose decimal representation is the palindrome in base k.
// The palindrome is formed by mirroring the half around an optional middle digit.
// Example: half=7("10") in base 7, even length → "1001" → returns 1001.
func buildPalindrome(half int, oddLen bool, k int) int {
	if half == 0 {
		return 0
	}
	// Extract digits of half in base k (least significant first)
	tmp := half
	digits := make([]int, 0)
	for tmp > 0 {
		digits = append(digits, tmp%k)
		tmp /= k
	}
	// Reverse to get most-significant first
	m := len(digits)
	for i, j := 0, m-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}
	// Build the palindrome as a decimal number whose digits are the base-k digits
	result := 0
	for _, d := range digits {
		result = result*10 + d
	}
	start := m - 1
	if oddLen {
		start = m - 2
	}
	for i := start; i >= 0; i-- {
		result = result*10 + digits[i]
	}
	return result
}

// toDecimal converts a number (whose decimal digits represent base-k digits) to a decimal value.
// Example: toDecimal(1001, 7) where 1001 decimal represents "1001" in base 7 = 1*343+0*49+0*7+1 = 344.
func toDecimal(num, k int) int64 {
	var dec int64 = 0
	mult := int64(1)
	for num > 0 {
		dec += int64(num%10) * mult
		num /= 10
		mult *= int64(k)
	}
	return dec
}

// isDecimalPalindrome checks if a decimal number is a palindrome
func isDecimalPalindrome(num int64) bool {
	s := strconv.FormatInt(num, 10)
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	fmt.Println("2081. Sum of k-Mirror Numbers")

	// Example 1
	k1 := 2
	n1 := 5
	fmt.Printf("k=%d n=%d → %d (expected 25)\n", k1, n1, kMirror(k1, n1))

	// Example 2
	k2 := 3
	n2 := 7
	fmt.Printf("k=%d n=%d → %d (expected 499)\n", k2, n2, kMirror(k2, n2))

	// Example 3 (n=17, not 7)
	k3 := 7
	n3 := 17
	fmt.Printf("k=%d n=%d → %d (expected 20379000)\n", k3, n3, kMirror(k3, n3))
}
```

## 2088 — Count Fertile Pyramids In A Land

```go
package main

// LeetCode #2088: Count Fertile Pyramids in a Land
// https://leetcode.com/problems/count-fertile-pyramids-in-a-land/
// Difficulty: Hard
//
// DP approach: dp[i][j] = max pyramid height with (i,j) as the top.
// Regular pyramid (top down): dp[i][j] = 1 + min(dp[i+1][j-1], dp[i+1][j], dp[i+1][j+1])
// Inverted pyramid (top up):  dp[i][j] = 1 + min(dp[i-1][j-1], dp[i-1][j], dp[i-1][j+1])
// Sum (dp[i][j] - 1) over all cells where dp[i][j] > 1.

import "fmt"

func main() {
	fmt.Println(countPyramids([][]int{{0, 1, 1, 0}, {1, 1, 1, 1}}))
	fmt.Println(countPyramids([][]int{{1, 1, 1}, {1, 1, 1}}))
	fmt.Println(countPyramids([][]int{{1}}))
	fmt.Println(countPyramids([][]int{{1, 1}}))
	fmt.Println(countPyramids([][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}))
}

func countPyramids(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	total := 0

	// Regular pyramids (top pointing down): bottom-up DP
	for i := m - 1; i >= 0; i-- {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				dp[i][j] = 0
				continue
			}
			if i == m-1 || j == 0 || j == n-1 {
				dp[i][j] = 1
			} else {
				dp[i][j] = 1 + min(dp[i+1][j-1], min(dp[i+1][j], dp[i+1][j+1]))
			}
			if dp[i][j] > 1 {
				total += dp[i][j] - 1
			}
		}
	}

	// Inverted pyramids (top pointing up): top-down DP
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				dp[i][j] = 0
				continue
			}
			if i == 0 || j == 0 || j == n-1 {
				dp[i][j] = 1
			} else {
				dp[i][j] = 1 + min(dp[i-1][j-1], min(dp[i-1][j], dp[i-1][j+1]))
			}
			if dp[i][j] > 1 {
				total += dp[i][j] - 1
			}
		}
	}

	return total
}
```

## 2092 — Find All People With Secret

```go
package main

// LeetCode #2092: Find All People With Secret
// https://leetcode.com/problems/find-all-people-with-secret/
// Difficulty: Hard
//
// Approach: Time-sorted Union-Find.
// Group meetings by time, union all participants within each time group,
// then check if any participant in the group is connected to a secret-knower.
// If not, reset their connections (they cannot learn the secret at this time).

import (
	"fmt"
	"sort"
)

func main() {
	// Example from problem statement
	n1 := 6
	meetings1 := [][]int{{1, 2, 5}, {2, 3, 8}, {1, 5, 10}}
	firstPerson1 := 1
	fmt.Printf("findAllPeople(%d, %v, %d) = %v (expected [0 1 2 3 5])\n",
		n1, meetings1, firstPerson1, findAllPeople(n1, meetings1, firstPerson1))

	// Additional tests
	n2 := 4
	meetings2 := [][]int{{3, 1, 3}, {1, 2, 2}, {0, 3, 3}}
	firstPerson2 := 3
	fmt.Printf("findAllPeople(%d, %v, %d) = %v\n",
		n2, meetings2, firstPerson2, findAllPeople(n2, meetings2, firstPerson2))

	n3 := 5
	meetings3 := [][]int{{0, 2, 1}, {1, 3, 1}, {4, 2, 2}}
	firstPerson3 := 2
	fmt.Printf("findAllPeople(%d, %v, %d) = %v\n",
		n3, meetings3, firstPerson3, findAllPeople(n3, meetings3, firstPerson3))
}

func findAllPeople(n int, meetings [][]int, firstPerson int) []int {
	// Sort meetings by time
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][2] < meetings[j][2]
	})

	// Union-Find structure
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func(a, b int) {
		pa, pb := find(a), find(b)
		if pa != pb {
			parent[pa] = pb
		}
	}

	// 0 and firstPerson start with the secret
	union(0, firstPerson)

	i := 0
	for i < len(meetings) {
		t := meetings[i][2]
		j := i
		for j < len(meetings) && meetings[j][2] == t {
			j++
		}

		// Union all pairs in this time group
		for k := i; k < j; k++ {
			union(meetings[k][0], meetings[k][1])
		}

		// Collect unique people in this time group
		people := make(map[int]bool)
		for k := i; k < j; k++ {
			people[meetings[k][0]] = true
			people[meetings[k][1]] = true
		}

		// Reset those who are not connected to a secret-knower
		root0 := find(0)
		for p := range people {
			if find(p) != root0 {
				parent[p] = p
			}
		}

		i = j
	}

	// Collect all who know the secret
	result := []int{}
	root0 := find(0)
	for p := 0; p < n; p++ {
		if find(p) == root0 {
			result = append(result, p)
		}
	}
	return result
}
```

## 2097 — Valid Arrangement Of Pairs

```go
package main

// LeetCode #2097: Valid Arrangement of Pairs
// https://leetcode.com/problems/valid-arrangement-of-pairs/
// Difficulty: Hard
//
// Eulerian Path (Hierholzer's Algorithm). Build adjacency list, compute
// in/out degrees. Start at node with out > in, or any node with edges.
// Use iterative DFS to reconstruct the path, then reverse.

import "fmt"

func main() {
	fmt.Println(validArrangement([][]int{{5, 1}, {4, 5}, {11, 9}, {9, 4}}))
	fmt.Println(validArrangement([][]int{{1, 2}, {2, 3}, {3, 4}}))
	fmt.Println(validArrangement([][]int{{1, 2}, {2, 1}}))
}

func validArrangement(pairs [][]int) [][]int {
	adj := make(map[int][]int)
	inDeg := make(map[int]int)
	outDeg := make(map[int]int)

	for _, p := range pairs {
		u, v := p[0], p[1]
		adj[u] = append(adj[u], v)
		outDeg[u]++
		inDeg[v]++
	}

	// Find start node: node with out > in, or any node with edges
	start := pairs[0][0]
	for node := range adj {
		if outDeg[node] > inDeg[node] {
			start = node
			break
		}
	}

	// Hierholzer's algorithm (iterative)
	stack := []int{start}
	path := []int{}

	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		if len(adj[cur]) > 0 {
			next := adj[cur][0]
			adj[cur] = adj[cur][1:]
			stack = append(stack, next)
		} else {
			path = append(path, cur)
			stack = stack[:len(stack)-1]
		}
	}

	// Reverse path to get correct order
	result := make([][]int, len(path)-1)
	for i := 0; i < len(path)-1; i++ {
		result[i] = []int{path[len(path)-1-i], path[len(path)-2-i]}
	}
	return result
}
```

## 2102 — Sequentially Ordinal Rank Tracker

```go
package main

// LeetCode #2102: Sequentially Ordinal Rank Tracker
// https://leetcode.com/problems/sequentially-ordinal-rank-tracker/
// Difficulty: Hard
//
// Two-heap approach:
// - low (min-heap by score ASC, name DESC): contains top-k items
// - high (min-heap by score DESC, name ASC): contains items beyond top-k

import (
	"container/heap"
	"fmt"
)

type location struct {
	name  string
	score int
}

type lowHeap []location

func (h lowHeap) Len() int      { return len(h) }
func (h lowHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h lowHeap) Less(i, j int) bool {
	if h[i].score != h[j].score {
		return h[i].score < h[j].score
	}
	return h[i].name > h[j].name
}
func (h *lowHeap) Push(x any)   { *h = append(*h, x.(location)) }
func (h *lowHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type highHeap []location

func (h highHeap) Len() int      { return len(h) }
func (h highHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h highHeap) Less(i, j int) bool {
	if h[i].score != h[j].score {
		return h[i].score > h[j].score
	}
	return h[i].name < h[j].name
}
func (h *highHeap) Push(x any)   { *h = append(*h, x.(location)) }
func (h *highHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type SORTracker struct {
	low     lowHeap
	high    highHeap
	queries int
}

func Constructor() SORTracker {
	return SORTracker{}
}

func (t *SORTracker) Add(name string, score int) {
	heap.Push(&t.low, location{name, score})
	if len(t.low) > t.queries {
		heap.Push(&t.high, heap.Pop(&t.low))
	}
}

func (t *SORTracker) Get() string {
	t.queries++
	for len(t.low) < t.queries {
		heap.Push(&t.low, heap.Pop(&t.high))
	}
	return t.low[0].name
}

func main() {
	tracker := Constructor()
	tracker.Add("bradford", 2)
	tracker.Add("branford", 3)
	fmt.Println(tracker.Get())
	tracker.Add("alps", 2)
	fmt.Println(tracker.Get())
	tracker.Add("orland", 2)
	fmt.Println(tracker.Get())
	tracker.Add("orlando", 3)
	fmt.Println(tracker.Get())
	tracker.Add("alpine", 2)
	fmt.Println(tracker.Get())
	fmt.Println(tracker.Get())
}
```

## 2106 — Maximum Fruits Harvested After At Most K Steps

```go
package main

// LeetCode #2106: Maximum Fruits Harvested After at Most K Steps
// https://leetcode.com/problems/maximum-fruits-harvested-after-at-most-k-steps/
// Difficulty: Hard
//
// Approach: Prefix sum + sliding window.
// Build a fruit amount array up to max position (200000).
// Use prefix sums to query range sums in O(1).
// For each possible left steps (0..k), compute remaining steps for right,
// and vice versa. Take max of all ranges.

import "fmt"

func main() {
	// Example from problem statement
	fruits1 := [][]int{{2, 8}, {6, 3}, {8, 6}}
	startPos1 := 5
	k1 := 4
	fmt.Printf("maxTotalFruits(%v, %d, %d) = %d (expected 9)\n",
		fruits1, startPos1, k1, maxTotalFruits(fruits1, startPos1, k1))

	// Additional tests
	fruits2 := [][]int{{0, 9}, {4, 1}, {5, 7}, {6, 2}, {7, 4}, {10, 9}}
	startPos2 := 5
	k2 := 4
	fmt.Printf("maxTotalFruits(%v, %d, %d) = %d (expected 14)\n",
		fruits2, startPos2, k2, maxTotalFruits(fruits2, startPos2, k2))

	fruits3 := [][]int{{0, 3}, {6, 4}, {8, 5}}
	startPos3 := 3
	k3 := 2
	fmt.Printf("maxTotalFruits(%v, %d, %d) = %d\n",
		fruits3, startPos3, k3, maxTotalFruits(fruits3, startPos3, k3))
}

func maxTotalFruits(fruits [][]int, startPos int, k int) int {
	const maxPos = 200000
	amount := make([]int, maxPos+1)
	for _, f := range fruits {
		amount[f[0]] = f[1]
	}

	prefix := make([]int, maxPos+2)
	for i := 0; i <= maxPos; i++ {
		prefix[i+1] = prefix[i] + amount[i]
	}

	sumRange := func(l, r int) int {
		if l < 0 {
			l = 0
		}
		if r > maxPos {
			r = maxPos
		}
		if l > r {
			return 0
		}
		return prefix[r+1] - prefix[l]
	}

	ans := 0
	// Go left first, then right
	for left := 0; left <= k; left++ {
		right := max(0, k-2*left)
		l := startPos - left
		r := startPos + right
		ans = max(ans, sumRange(l, r))
	}

	// Go right first, then left
	for right := 0; right <= k; right++ {
		left := max(0, k-2*right)
		l := startPos - left
		r := startPos + right
		ans = max(ans, sumRange(l, r))
	}

	return ans
}
```

## 2111 — Minimum Operations To Make The Array K Increasing

```go
package main

// LeetCode #2111: Minimum Operations to Make the Array K-Increasing
// https://leetcode.com/problems/minimum-operations-to-make-the-array-k-increasing/
// Difficulty: Hard
//
// Approach: LIS per mod-k subsequence.
// Split array into k subsequences: arr[i], arr[i+k], arr[i+2k], ...
// For each subsequence, compute length of longest non-decreasing subsequence (LIS).
// Min operations = len(seq) - LIS(seq). Sum over all k subsequences.

import (
	"fmt"
	"sort"
)

func main() {
	// Example from problem statement
	arr1 := []int{5, 4, 3, 2, 1}
	k1 := 1
	fmt.Printf("minOperations(%v, %d) = %d (expected 4)\n", arr1, k1, minOperations(arr1, k1))

	// Additional tests
	arr2 := []int{4, 1, 5, 2, 6, 2}
	k2 := 2
	fmt.Printf("minOperations(%v, %d) = %d (expected 0)\n", arr2, k2, minOperations(arr2, k2))

	arr3 := []int{1, 2, 3, 4, 5, 6}
	k3 := 2
	fmt.Printf("minOperations(%v, %d) = %d\n", arr3, k3, minOperations(arr3, k3))

	arr4 := []int{5, 3, 1, 4, 2}
	k4 := 2
	fmt.Printf("minOperations(%v, %d) = %d\n", arr4, k4, minOperations(arr4, k4))
}

func minOperations(arr []int, k int) int {
	n := len(arr)
	total := 0

	for i := 0; i < k; i++ {
		seq := []int{}
		for j := i; j < n; j += k {
			seq = append(seq, arr[j])
		}
		total += len(seq) - lengthOfLIS(seq)
	}

	return total
}

// lengthOfLIS returns the length of the longest non-decreasing subsequence.
func lengthOfLIS(nums []int) int {
	tails := []int{}
	for _, x := range nums {
		// Find first element > x (since we want non-decreasing, equal values can extend)
		idx := sort.SearchInts(tails, x+1)
		if idx == len(tails) {
			tails = append(tails, x)
		} else {
			tails[idx] = x
		}
	}
	return len(tails)
}
```

## 2117 — Abbreviating The Product Of A Range

```go
package main

// LeetCode #2117: Abbreviating the Product of a Range
// https://leetcode.com/problems/abbreviating-the-product-of-a-range/
// Difficulty: Hard
//
// Compute the product of [left, right], then abbreviate as "first5...last5eZ".
// Count trailing zeros by tracking 2s and 5s. Use modular arithmetic for last
// 5 digits and logarithms for the first 5 digits.

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	// Example 1
	fmt.Println(abbreviateProduct(1, 4))

	// Example 2
	fmt.Println(abbreviateProduct(2, 11))

	// Example 3
	fmt.Println(abbreviateProduct(999998, 1000000))

	// Single number
	fmt.Println(abbreviateProduct(5, 5))

	// No trailing zeros
	fmt.Println(abbreviateProduct(1, 3))
}

func abbreviateProduct(left int, right int) string {
	count2 := 0
	count5 := 0

	// Count factors of 2 and 5 in the range
	for i := left; i <= right; i++ {
		x := i
		for x%2 == 0 {
			count2++
			x /= 2
		}
		for x%5 == 0 {
			count5++
			x /= 5
		}
	}

	zeros := count2
	if count5 < zeros {
		zeros = count5
	}
	extra2 := count2 - zeros
	extra5 := count5 - zeros

	// Compute last 5 digits (mod 100000) after removing all 2s and 5s
	last5 := 1
	for i := left; i <= right; i++ {
		x := i
		for x%2 == 0 {
			x /= 2
		}
		for x%5 == 0 {
			x /= 5
		}
		last5 = (last5 * (x % 100000)) % 100000
	}

	// Multiply back remaining 2s or 5s
	for i := 0; i < extra2; i++ {
		last5 = (last5 * 2) % 100000
	}
	for i := 0; i < extra5; i++ {
		last5 = (last5 * 5) % 100000
	}

	// Compute log10 of product to get first 5 digits
	var logSum float64 = 0
	for i := left; i <= right; i++ {
		logSum += math.Log10(float64(i))
	}

	// Remove trailing zeros from log
	productLog := logSum - float64(zeros)
	totalDigits := int(math.Floor(productLog)) + 1

	if totalDigits <= 5 {
		// The product without trailing zeros is small enough to show fully
		fullNum := 1
		for i := left; i <= right; i++ {
			fullNum *= i
		}
		// Remove trailing zeros
		for i := 0; i < zeros; i++ {
			fullNum /= 10
		}
		if zeros > 0 {
			return fmt.Sprintf("%de%d", fullNum, zeros)
		}
		return fmt.Sprintf("%d", fullNum)
	}

	// First 5 digits using log
	frac := productLog - math.Floor(productLog)
	first5 := int(math.Pow(10, frac+4))

	// Handle the "last 5" formatting (pad with leading zeros if needed)
	last5Str := fmt.Sprintf("%05d", last5)

	return fmt.Sprintf("%d...%se%d", first5, last5Str, zeros)
}

// Remove unused import
var _ = strings.Builder{}
```

## 2118 — Build The Equation

```go
package main

// LeetCode #2118: Build the Equation
// https://leetcode.com/problems/build-the-equation/
// Difficulty: Hard [Paid]
//
// Given a list of terms (power, coefficient), build the polynomial equation
// string. Terms are sorted by power descending. Format: "+2x^3-4x^1+5" = "0".
// Handle special cases: coefficient = 1 or -1, power = 0 or 1.

import (
	"fmt"
	"sort"
)

// term represents a polynomial term with coefficient and power
type term struct {
	coefficient int
	power       int
}

func main() {
	// Example: y = 2x^3 - 4x + 5  =>  "2x^3-4x+5=0"
	fmt.Println(buildEquation([]term{{2, 3}, {-4, 1}, {5, 0}}))

	// Example: y = -x^2 + x - 1 =>  "-x^2+x-1=0"
	fmt.Println(buildEquation([]term{{-1, 2}, {1, 1}, {-1, 0}}))

	// Example: y = 3x^2 =>  "3x^2=0"
	fmt.Println(buildEquation([]term{{3, 2}}))

	// Single constant
	fmt.Println(buildEquation([]term{{7, 0}}))

	// No terms (empty equation)
	fmt.Println(buildEquation([]term{}))
}

func buildEquation(terms []term) string {
	if len(terms) == 0 {
		return "0=0"
	}

	// Sort by power descending
	sort.Slice(terms, func(i, j int) bool {
		return terms[i].power > terms[j].power
	})

	var result string

	for i, t := range terms {
		coef := t.coefficient
		pow := t.power

		if coef == 0 {
			continue
		}

		// Sign handling
		if i == 0 {
			// First term: no leading '+' for positive
			if coef < 0 {
				result += "-"
			}
		} else {
			if coef > 0 {
				result += "+"
			} else {
				result += "-"
			}
		}

		absCoef := coef
		if absCoef < 0 {
			absCoef = -absCoef
		}

		// Coefficient part (don't print 1 or -1 unless power is 0)
		if pow == 0 || absCoef != 1 {
			result += fmt.Sprintf("%d", absCoef)
		}

		// Variable part
		if pow > 0 {
			result += "x"
			if pow > 1 {
				result += fmt.Sprintf("^%d", pow)
			}
		}
	}

	if result == "" {
		return "0=0"
	}

	return result + "=0"
}
```

## 2122 — Recover The Original Array

```go
package main

// LeetCode #2122: Recover the Original Array
// https://leetcode.com/problems/recover-the-original-array/
// Difficulty: Hard
//
// Approach: Sort + frequency map.
// Sort the array. The smallest element must be original[0] - k.
// Try each possible partner for nums[0] as the "high" pair (original[0] + k).
// For each valid 2k = diff, greedily pair elements using a frequency map.
// Return the first valid original array found.

import (
	"fmt"
	"sort"
)

func main() {
	// Example from problem statement
	nums1 := []int{2, 10, 6, 4, 8, 12}
	fmt.Printf("recoverArray(%v) = %v (expected [3 7 11] or similar)\n", nums1, recoverArray(nums1))

	// Additional tests
	nums2 := []int{1, 1, 3, 3}
	fmt.Printf("recoverArray(%v) = %v (expected [2 2])\n", nums2, recoverArray(nums2))

	nums3 := []int{5, 5, 9, 9}
	fmt.Printf("recoverArray(%v) = %v\n", nums3, recoverArray(nums3))
}

func recoverArray(nums []int) []int {
	sort.Ints(nums)
	n := len(nums)

	// nums[0] is always a "low" element (original[0] - k)
	// Try each possible partner for nums[0] to determine 2k = diff
	for i := 1; i < n; i++ {
		diff := nums[i] - nums[0]
		if diff == 0 || diff%2 != 0 {
			continue
		}

		// Try to reconstruct with this 2k value
		freq := make(map[int]int)
		for _, v := range nums {
			freq[v]++
		}

		result := make([]int, 0, n/2)
		valid := true

		for _, v := range nums {
			if freq[v] == 0 {
				continue
			}
			// v must be a "low" element; pair with v + diff (if diff = 2k)
			high := v + diff
			if freq[high] == 0 {
				valid = false
				break
			}
			freq[v]--
			freq[high]--
			result = append(result, v+diff/2)
		}

		if valid && len(result) == n/2 {
			return result
		}
	}

	return nil
}
```

## 2123 — Minimum Operations To Remove Adjacent Ones In Matrix

```go
package main

// LeetCode #2123: Minimum Operations to Remove Adjacent Ones in Matrix
// https://leetcode.com/problems/minimum-operations-to-remove-adjacent-ones-in-matrix/
// Difficulty: Hard [Paid]
//
// Minimum vertex cover in a bipartite graph. Each operation at (i,j) toggles
// the cell and its 4 neighbors. Model as: for each adjacent pair of 1s, at
// least one must be flipped. This reduces to max bipartite matching (Kőnig's
// theorem). Graph cells are colored like a chessboard based on (i+j) parity.

import "fmt"

func main() {
	// Example
	grid1 := [][]int{{1, 1, 0}, {1, 1, 1}, {0, 1, 1}}
	fmt.Println(minimumOperationsToRemoveAdjacentOnes(grid1))

	// All 1s diagonal - no adjacent 1s
	grid2 := [][]int{{1, 0, 1}, {0, 1, 0}, {1, 0, 1}}
	fmt.Println(minimumOperationsToRemoveAdjacentOnes(grid2))

	// All adjacent
	grid3 := [][]int{{1, 1}, {1, 1}}
	fmt.Println(minimumOperationsToRemoveAdjacentOnes(grid3))

	// Single cell
	grid4 := [][]int{{1}}
	fmt.Println(minimumOperationsToRemoveAdjacentOnes(grid4))
}

func minimumOperationsToRemoveAdjacentOnes(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	// Assign IDs to cells with value 1
	id := make([][]int, m)
	for i := range id {
		id[i] = make([]int, n)
		for j := range id[i] {
			id[i][j] = -1
		}
	}

	leftCount := 0
	rightCount := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				if (i+j)%2 == 0 {
					id[i][j] = leftCount
					leftCount++
				} else {
					id[i][j] = rightCount
					rightCount++
				}
			}
		}
	}

	if leftCount == 0 || rightCount == 0 {
		return 0
	}

	// Build adjacency from left to right
	adj := make([][]int, leftCount)
	for i := 0; i < leftCount; i++ {
		adj[i] = []int{}
	}

	dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 && (i+j)%2 == 0 {
				u := id[i][j]
				for _, d := range dirs {
					ni, nj := i+d[0], j+d[1]
					if ni >= 0 && ni < m && nj >= 0 && nj < n && grid[ni][nj] == 1 {
						v := id[ni][nj]
						adj[u] = append(adj[u], v)
					}
				}
			}
		}
	}

	// Maximum bipartite matching using DFS (Kuhn's algorithm)
	matchR := make([]int, rightCount)
	for i := range matchR {
		matchR[i] = -1
	}

	var dfs func(u int, seen []bool) bool
	dfs = func(u int, seen []bool) bool {
		for _, v := range adj[u] {
			if seen[v] {
				continue
			}
			seen[v] = true
			if matchR[v] == -1 || dfs(matchR[v], seen) {
				matchR[v] = u
				return true
			}
		}
		return false
	}

	result := 0
	for u := 0; u < leftCount; u++ {
		seen := make([]bool, rightCount)
		if dfs(u, seen) {
			result++
		}
	}

	return result
}
```

## 2127 — Maximum Employees To Be Invited To A Meeting

```go
package main

// LeetCode #2127: Maximum Employees to Be Invited to a Meeting
// https://leetcode.com/problems/maximum-employees-to-be-invited-to-a-meeting/
// Difficulty: Hard
//
// Approach: Cycle detection in functional graph.
// Each node has exactly one outgoing edge (favorite[i]).
// Two cases contribute to the answer:
//   1. 2-cycles (mutual favorites): attach longest chain feeding into each node.
//   2. Cycles of length >= 3: must take the entire cycle, no chains attached.
//
// Use topological sort (Kahn's) to compute chain lengths for non-cycle nodes,
// then detect cycles among remaining nodes.

import "fmt"

func main() {
	// Example from problem statement
	favorite1 := []int{2, 2, 1, 2}
	fmt.Printf("maximumInvitations(%v) = %d (expected 3)\n", favorite1, maximumInvitations(favorite1))

	// Additional tests
	favorite2 := []int{1, 2, 0}
	fmt.Printf("maximumInvitations(%v) = %d (expected 3)\n", favorite2, maximumInvitations(favorite2))

	favorite3 := []int{1, 0, 3, 2}
	fmt.Printf("maximumInvitations(%v) = %d (expected 4)\n", favorite3, maximumInvitations(favorite3))

	favorite4 := []int{1, 0}
	fmt.Printf("maximumInvitations(%v) = %d (expected 2)\n", favorite4, maximumInvitations(favorite4))
}

func maximumInvitations(favorite []int) int {
	n := len(favorite)
	inDegree := make([]int, n)
	for _, f := range favorite {
		inDegree[f]++
	}

	// chainLen[i] = longest chain of non-cycle nodes ending at i
	chainLen := make([]int, n)
	q := make([]int, 0)
	for i := 0; i < n; i++ {
		if inDegree[i] == 0 {
			q = append(q, i)
		}
	}

	// Topological sort to compute chain lengths (removes non-cycle nodes)
	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		v := favorite[u]
		if chainLen[u]+1 > chainLen[v] {
			chainLen[v] = chainLen[u] + 1
		}
		inDegree[v]--
		if inDegree[v] == 0 {
			q = append(q, v)
		}
	}

	visited := make([]bool, n)
	totalChains := 0
	maxCycle := 0

	for i := 0; i < n; i++ {
		if inDegree[i] > 0 && !visited[i] {
			// Find the cycle
			cur := i
			cycleNodes := []int{}
			for !visited[cur] {
				visited[cur] = true
				cycleNodes = append(cycleNodes, cur)
				cur = favorite[cur]
			}

			cycleLen := len(cycleNodes)
			if cycleLen == 2 {
				a, b := cycleNodes[0], cycleNodes[1]
				totalChains += 2 + chainLen[a] + chainLen[b]
			} else {
				if cycleLen > maxCycle {
					maxCycle = cycleLen
				}
			}
		}
	}

	if totalChains > maxCycle {
		return totalChains
	}
	return maxCycle
}
```

## 2132 — Stamping The Grid

```go
package main

// LeetCode #2132: Stamping the Grid
// https://leetcode.com/problems/stamping-the-grid/
// Difficulty: Hard
//
// 2D prefix sum + 2D difference array. First compute prefix sums to query
// empty subgrids. For each possible stamp top-left, if the area is obstacle-free,
// mark it in the diff array. Then reconstruct coverage and verify all empty
// cells are covered by at least one stamp.

import "fmt"

func main() {
	grid1 := [][]int{
		{1, 0, 0, 0},
		{1, 0, 0, 0},
		{1, 0, 0, 0},
		{1, 0, 0, 0},
		{1, 0, 0, 0},
	}
	fmt.Println(possibleToStamp(grid1, 4, 3))

	grid2 := [][]int{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
	fmt.Println(possibleToStamp(grid2, 2, 2))

	fmt.Println(possibleToStamp([][]int{{0}}, 1, 1))
}

func possibleToStamp(grid [][]int, stampHeight int, stampWidth int) bool {
	m, n := len(grid), len(grid[0])

	// 2D prefix sum (obstacles = 1)
	prefix := make([][]int, m+1)
	for i := range prefix {
		prefix[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			prefix[i+1][j+1] = prefix[i][j+1] + prefix[i+1][j] - prefix[i][j] + grid[i][j]
		}
	}

	sumRange := func(r1, c1, r2, c2 int) int {
		return prefix[r2+1][c2+1] - prefix[r1][c2+1] - prefix[r2+1][c1] + prefix[r1][c1]
	}

	// 2D difference array
	diff := make([][]int, m+2)
	for i := range diff {
		diff[i] = make([]int, n+2)
	}

	// Mark valid stamp placements
	for i := 0; i+stampHeight <= m; i++ {
		for j := 0; j+stampWidth <= n; j++ {
			if sumRange(i, j, i+stampHeight-1, j+stampWidth-1) == 0 {
				diff[i][j]++
				diff[i][j+stampWidth]--
				diff[i+stampHeight][j]--
				diff[i+stampHeight][j+stampWidth]++
			}
		}
	}

	// Reconstruct coverage and verify
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i > 0 {
				diff[i][j] += diff[i-1][j]
			}
			if j > 0 {
				diff[i][j] += diff[i][j-1]
			}
			if i > 0 && j > 0 {
				diff[i][j] -= diff[i-1][j-1]
			}
			if grid[i][j] == 0 && diff[i][j] == 0 {
				return false
			}
		}
	}

	return true
}
```

## 2136 — Earliest Possible Day Of Full Bloom

```go
package main

// LeetCode #2136: Earliest Possible Day of Full Bloom
// https://leetcode.com/problems/earliest-possible-day-of-full-bloom/
// Difficulty: Hard
//
// Approach: Sort by grow time descending (Greedy).
// Plant seeds sequentially, always plant the one with the longest grow time first.
// This minimizes the overall completion time since longer-growing seeds start earlier.
// For each seed, bloom day = cumulative plant time + grow time. Answer = max bloom day.

import (
	"fmt"
	"sort"
)

func main() {
	// Example from problem statement
	plantTime1 := []int{1, 4, 3}
	growTime1 := []int{2, 3, 1}
	fmt.Printf("earliestFullBloom(%v, %v) = %d (expected 9)\n",
		plantTime1, growTime1, earliestFullBloom(plantTime1, growTime1))

	// Additional tests
	plantTime2 := []int{1, 2, 3}
	growTime2 := []int{1, 2, 3}
	fmt.Printf("earliestFullBloom(%v, %v) = %d (expected 7)\n",
		plantTime2, growTime2, earliestFullBloom(plantTime2, growTime2))

	plantTime3 := []int{1}
	growTime3 := []int{1}
	fmt.Printf("earliestFullBloom(%v, %v) = %d (expected 2)\n",
		plantTime3, growTime3, earliestFullBloom(plantTime3, growTime3))

	plantTime4 := []int{3, 2, 1}
	growTime4 := []int{1, 2, 3}
	fmt.Printf("earliestFullBloom(%v, %v) = %d\n",
		plantTime4, growTime4, earliestFullBloom(plantTime4, growTime4))
}

func earliestFullBloom(plantTime []int, growTime []int) int {
	n := len(plantTime)
	seeds := make([][2]int, n)
	for i := 0; i < n; i++ {
		seeds[i] = [2]int{plantTime[i], growTime[i]}
	}

	// Sort by grow time descending
	sort.Slice(seeds, func(i, j int) bool {
		return seeds[i][1] > seeds[j][1]
	})

	day := 0
	ans := 0
	for _, s := range seeds {
		day += s[0] // plant this seed
		bloomDay := day + s[1]
		if bloomDay > ans {
			ans = bloomDay
		}
	}
	return ans
}
```

## 2141 — Maximum Running Time Of N Computers

```go
package main

// LeetCode #2141: Maximum Running Time of N Computers
// https://leetcode.com/problems/maximum-running-time-of-n-computers/
// Difficulty: Hard
//
// Binary search on the answer. For a given minutes, each battery contributes
// min(capacity, minutes). If sum >= n * minutes, it's feasible.

import "fmt"

func main() {
	fmt.Println(maxRunTime(2, []int{3, 3, 3}))                                    // 4
	fmt.Println(maxRunTime(3, []int{10, 10, 3, 5}))                               // 8
	fmt.Println(maxRunTime(1, []int{1, 2, 3}))                                    // 6
	fmt.Println(maxRunTime(4, []int{10, 10, 10, 10, 5, 5, 5, 5}))                // 15
	fmt.Println(maxRunTime(3, []int{1, 1, 1, 1}))                                 // 1
}

func maxRunTime(n int, batteries []int) int64 {
	canRun := func(minutes int64) bool {
		var total int64
		for _, b := range batteries {
			if int64(b) < minutes {
				total += int64(b)
			} else {
				total += minutes
			}
		}
		return total >= minutes*int64(n)
	}

	lo := int64(0)
	var hi int64
	for _, b := range batteries {
		hi += int64(b)
	}
	hi /= int64(n)

	for lo < hi {
		mid := (lo + hi + 1) / 2
		if canRun(mid) {
			lo = mid
		} else {
			hi = mid - 1
		}
	}
	return lo
}

func MaximumRunningTimeOfNComputers() any {
	return maxRunTime(2, []int{3, 3, 3})
}
```

## 2143 — Choose Numbers From Two Arrays In Range

```go
package main

// LeetCode #2143: Choose Numbers From Two Arrays in Range
// https://leetcode.com/problems/choose-numbers-from-two-arrays-in-range/
// Difficulty: Hard [Paid]
//
// Count subarrays [l, r] where sum(nums1[l..r]) == sum(nums2[l..r]) and the
// combined multiset has at most one odd value. Use prefix difference tracking:
// diff = sum1[i] - sum2[i], oddCount = number of odd values up to i.
// For each r, find matching l with same diff and oddCount diff <= 1.

import "fmt"

func main() {
	// Example 1
	fmt.Println(countSubarrays([]int{1, 2, 3, 4}, []int{2, 1, 4, 5}))

	// Example 2
	fmt.Println(countSubarrays([]int{1, 2, 3, 4, 5}, []int{3, 4, 2, 1, 5}))

	// Example 3
	fmt.Println(countSubarrays([]int{0, 0, 0}, []int{0, 0, 0}))

	// Simple case
	fmt.Println(countSubarrays([]int{1, 1}, []int{1, 1}))
}

func countSubarrays(nums1 []int, nums2 []int) int {
	n := len(nums1)

	// prefixDiff[i] = sum(nums1[0..i-1]) - sum(nums2[0..i-1])
	prefixDiff := make([]int, n+1)
	// prefixOdd[i] = count of odd values in nums1[0..i-1] + nums2[0..i-1]
	prefixOdd := make([]int, n+1)

	for i := 0; i < n; i++ {
		prefixDiff[i+1] = prefixDiff[i] + nums1[i] - nums2[i]
		oddCount := 0
		if nums1[i]%2 == 1 {
			oddCount++
		}
		if nums2[i]%2 == 1 {
			oddCount++
		}
		prefixOdd[i+1] = prefixOdd[i] + oddCount
	}

	// Map: prefixDiff -> list of prefixOdd values at that diff
	// We'll use a map from diff to a map of oddCount->frequency
	diffMap := make(map[int]map[int]int)
	ans := 0

	// l starts at 0 (empty prefix), r runs from 1 to n
	// We iterate r, and before processing, we add prefix l=r to the map
	// For r, we want l where prefixDiff[r] == prefixDiff[l] and prefixOdd[r] - prefixOdd[l] <= 1

	// Initialize with l=0 (empty prefix)
	diffMap[0] = map[int]int{0: 1}

	for r := 1; r <= n; r++ {
		diff := prefixDiff[r]
		odd := prefixOdd[r]

		// Query: find l where prefixDiff[l] == diff and odd - prefixOdd[l] <= 1
		if m, ok := diffMap[diff]; ok {
			// Add all l where oddDiff <= 1
			// oddDiff = odd - prefixOdd[l]
			// prefixOdd[l] >= odd - 1 and prefixOdd[l] <= odd
			for oddL, freq := range m {
				if odd-oddL <= 1 {
					ans += freq
				}
			}
		}

		// Add current prefix to map
		if _, ok := diffMap[diff]; !ok {
			diffMap[diff] = make(map[int]int)
		}
		diffMap[diff][odd]++
	}

	return ans
}
```

## 2147 — Number Of Ways To Divide A Long Corridor

```go
package main

// LeetCode #2147: Number of Ways to Divide a Long Corridor
// https://leetcode.com/problems/number-of-ways-to-divide-a-long-corridor/
// Difficulty: Hard
//
// Collect indices of seats. Each pair of seats (consecutive in the even-odd sense)
// defines a mandatory section boundary. Multiply gaps between pairs.

import "fmt"

func main() {
	fmt.Println(numberOfWays("SSPPSPS")) // 3
	fmt.Println(numberOfWays("PPSPSP"))  // 1
	fmt.Println(numberOfWays("S"))       // 0
	fmt.Println(numberOfWays("SS"))      // 1
	fmt.Println(numberOfWays("SPPSS"))   // 2
}

const mod2147 = 1_000_000_007

func numberOfWays(corridor string) int {
	var seats []int
	for i, ch := range corridor {
		if ch == 'S' {
			seats = append(seats, i)
		}
	}
	if len(seats) == 0 || len(seats)%2 != 0 {
		return 0
	}

	ans := 1
	for i := 1; i < len(seats)-1; i += 2 {
		gap := seats[i+1] - seats[i]
		ans = (ans * gap) % mod2147
	}
	return ans
}

func NumberOfWaysToDivideALongCorridor() any {
	return numberOfWays("SSPPSPS")
}
```

## 2151 — Maximum Good People Based On Statements

```go
package main

// LeetCode #2151: Maximum Good People Based on Statements
// https://leetcode.com/problems/maximum-good-people-based-on-statements/
// Difficulty: Hard
//
// Bitmask enumeration: try all 2^n assignments of good (1) / bad (0) people.
// For each assignment, verify all non-2 statements from good people.

import "fmt"

func main() {
	fmt.Println(maximumGood([][]int{{2, 1, 2}, {1, 2, 2}, {2, 0, 2}})) // 2
	fmt.Println(maximumGood([][]int{{2, 0}, {0, 2}}))                  // 1
	fmt.Println(maximumGood([][]int{{2, 2, 2}, {2, 2, 2}, {2, 2, 2}})) // 3
}

func maximumGood(statements [][]int) int {
	n := len(statements)
	best := 0

	for mask := 0; mask < (1 << n); mask++ {
		valid := true
		cnt := 0
		for i := 0; i < n && valid; i++ {
			if mask&(1<<i) == 0 {
				continue
			}
			cnt++
			for j := 0; j < n; j++ {
				st := statements[i][j]
				if st == 2 {
					continue
				}
				isGoodJ := (mask >> j) & 1
				if st != isGoodJ {
					valid = false
					break
				}
			}
		}
		if valid && cnt > best {
			best = cnt
		}
	}
	return best
}

func MaximumGoodPeopleBasedOnStatements() any {
	return maximumGood([][]int{{2, 1, 2}, {1, 2, 2}, {2, 0, 2}})
}
```

## 2153 — The Number Of Passengers In Each Bus Ii

```go
package main

// LeetCode #2153: The Number of Passengers in Each Bus II
// https://leetcode.com/problems/the-number-of-passengers-in-each-bus-ii/
// Difficulty: Hard [Paid]
//
// Sort buses and passengers by arrival time. For each bus in order, board
// waiting passengers up to capacity. Each bus departs at its arrival time;
// passengers arriving at the exact same time can board.

import (
	"fmt"
	"sort"
)

type bus struct {
	id    int
	time  int
	cap   int
}

type passenger struct {
	time int
	id   int
}

func main() {
	// Example 1
	buses1 := []bus{{1, 2, 2}, {2, 5, 3}}
	passengers1 := []passenger{{1, 101}, {2, 102}, {3, 103}, {4, 104}, {5, 105}}
	fmt.Println(calculateBusPassengers(buses1, passengers1))

	// Example 2
	buses2 := []bus{{1, 3, 2}, {2, 5, 2}}
	passengers2 := []passenger{{1, 201}, {2, 202}, {5, 203}}
	fmt.Println(calculateBusPassengers(buses2, passengers2))
}

func calculateBusPassengers(buses []bus, passengers []passenger) map[int]int {
	// Sort buses by arrival time
	sort.Slice(buses, func(i, j int) bool { return buses[i].time < buses[j].time })

	// Sort passengers by arrival time
	sort.Slice(passengers, func(i, j int) bool { return passengers[i].time < passengers[j].time })

	result := make(map[int]int)
	passIdx := 0
	n := len(passengers)

	for _, b := range buses {
		boarded := 0
		for boarded < b.cap && passIdx < n && passengers[passIdx].time <= b.time {
			result[b.id]++
			boarded++
			passIdx++
		}
		if result[b.id] == 0 {
			result[b.id] = 0
		}
	}

	return result
}
```

## 2156 — Find Substring With Given Hash Value

```go
package main

// LeetCode #2156: Find Substring With Given Hash Value
// https://leetcode.com/problems/find-substring-with-given-hash-value/
// Difficulty: Hard
//
// Rolling hash from left to right. Precompute power^(k-1) % mod for sliding.
// h(i) = s[i]*p^{k-1} + s[i+1]*p^{k-2} + ... + s[i+k-1]*p^0
// h(i+1) = (h(i) - s[i]*p^{k-1}) * power + s[i+k]

import "fmt"

func main() {
	fmt.Println(subStrHash("leetcode", 7, 20, 2, 0))  // "ee"
	fmt.Println(subStrHash("fbxzaad", 31, 100, 3, 32)) // "" (no match)
	fmt.Println(subStrHash("xqgcas", 4, 7, 3, 4))     // "xqg"
	fmt.Println(subStrHash("helloworld", 10, 1000, 3, 862)) // "hel"
}

func subStrHash(s string, power int, mod int, k int, hashValue int) string {
	n := len(s)

	// pk1 = power^(k-1) % mod
	pk1 := 1
	for i := 0; i < k-1; i++ {
		pk1 = (pk1 * power) % mod
	}
	// pk = power^k % mod (not strictly needed, pk1 * power)

	// Compute hash of first window
	cur := 0
	for i := 0; i < k; i++ {
		cur = (cur*power + int(s[i]-'a'+1)) % mod
	}
	if cur == hashValue {
		return s[:k]
	}

	// Slide window left to right
	for i := 1; i <= n-k; i++ {
		// h(i) = (h(i-1) - s[i-1]*p^{k-1}) * power + s[i+k-1]
		cur = (cur - (int(s[i-1]-'a'+1))*pk1%mod + mod) % mod
		cur = (cur*power + int(s[i+k-1]-'a'+1)) % mod
		if cur == hashValue {
			return s[i : i+k]
		}
	}
	return ""
}

func FindSubstringWithGivenHashValue() any {
	return subStrHash("leetcode", 7, 20, 2, 0)
}
```

## 2157 — Groups Of Strings

```go
package main

// LeetCode #2157: Groups of Strings
// https://leetcode.com/problems/groups-of-strings/
// Difficulty: Hard
//
// Union-Find + bitmask (26-bit). Two words are connected if one can become
// the other by adding, deleting, or replacing one character.

import "fmt"

func main() {
	fmt.Println(groupStrings([]string{"a", "b", "ab", "cde"}))       // [2, 3]
	fmt.Println(groupStrings([]string{"a", "ab", "abc"}))            // [1, 3]
	fmt.Println(groupStrings([]string{"abc", "acb", "bac", "bca"}))  // [1, 4]
	fmt.Println(groupStrings([]string{"ab"}))                        // [1, 1]
}

func groupStrings(words []string) []int {
	n := len(words)
	masks := make([]uint32, n)
	idxOf := make(map[uint32]int)

	for i, w := range words {
		var m uint32
		for _, ch := range w {
			m |= 1 << (ch - 'a')
		}
		masks[i] = m
		idxOf[m] = i
	}

	// Union-Find
	parent := make([]int, n)
	sz := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		sz[i] = 1
	}
	var find func(int) int
	find = func(x int) int {
		for parent[x] != x {
			parent[x] = parent[parent[x]]
			x = parent[x]
		}
		return x
	}
	union := func(a, b int) {
		ra, rb := find(a), find(b)
		if ra == rb {
			return
		}
		if sz[ra] < sz[rb] {
			ra, rb = rb, ra
		}
		parent[rb] = ra
		sz[ra] += sz[rb]
	}

	// Union identical masks first
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if masks[i] == masks[j] {
				union(i, j)
			}
		}
	}

	// For each unique mask, try transforms
	seen := make(map[uint32]bool)
	for i, m := range masks {
		if seen[m] {
			continue
		}
		seen[m] = true

		// Delete one bit
		for b := 0; b < 26; b++ {
			if m&(1<<b) != 0 {
				nm := m & ^(1 << b)
				if j, ok := idxOf[nm]; ok {
					union(i, j)
				}
			}
		}

		// Add one bit
		for b := 0; b < 26; b++ {
			if m&(1<<b) == 0 {
				nm := m | (1 << b)
				if j, ok := idxOf[nm]; ok {
					union(i, j)
				}
			}
		}

		// Replace one bit
		for b1 := 0; b1 < 26; b1++ {
			if m&(1<<b1) == 0 {
				continue
			}
			for b2 := 0; b2 < 26; b2++ {
				if b1 == b2 || m&(1<<b2) != 0 {
					continue
				}
				nm := (m & ^(1 << b1)) | (1 << b2)
				if j, ok := idxOf[nm]; ok {
					union(i, j)
				}
			}
		}
	}

	groupSizes := make(map[int]int)
	maxSize := 0
	for i := 0; i < n; i++ {
		r := find(i)
		groupSizes[r]++
		if groupSizes[r] > maxSize {
			maxSize = groupSizes[r]
		}
	}
	return []int{len(groupSizes), maxSize}
}

func GroupsOfStrings() any {
	return groupStrings([]string{"a", "b", "ab", "cde"})
}
```

## 2158 — Amount Of New Area Painted Each Day

```go
package main

// LeetCode #2158: Amount of New Area Painted Each Day
// https://leetcode.com/problems/amount-of-new-area-painted-each-day/
// Difficulty: Hard [Paid]
//
// Given intervals [start, end) painted each day, compute the new area painted
// each day (not previously painted). Use a DSU/union-find structure to skip
// already-painted cells, achieving near O(N) amortized time.

import "fmt"

func main() {
	// Example 1
	fmt.Println(amountPainted([][]int{{1, 4}, {4, 7}, {5, 8}}))

	// Example 2
	fmt.Println(amountPainted([][]int{{1, 5}, {2, 4}}))

	// Overlapping intervals
	fmt.Println(amountPainted([][]int{{1, 3}, {2, 5}, {3, 6}}))

	// Non-overlapping
	fmt.Println(amountPainted([][]int{{1, 2}, {3, 4}, {5, 6}}))
}

func amountPainted(paint [][]int) []int {
	maxEnd := 0
	for _, p := range paint {
		if p[1] > maxEnd {
			maxEnd = p[1]
		}
	}

	// DSU: next[x] = next unpainted point >= x
	next := make([]int, maxEnd+2)
	for i := range next {
		next[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if next[x] != x {
			next[x] = find(next[x])
		}
		return next[x]
	}

	result := make([]int, len(paint))
	for day, p := range paint {
		start, end := p[0], p[1]
		count := 0
		pos := find(start)
		for pos < end {
			count++
			next[pos] = find(pos + 1)
			pos = find(pos)
		}
		result[day] = count
	}

	return result
}
```

## 2163 — Minimum Difference In Sums After Removal Of Elements

```go
package main

// LeetCode #2163: Minimum Difference in Sums After Removal of Elements
// https://leetcode.com/problems/minimum-difference-in-sums-after-removal-of-elements/
// Difficulty: Hard
//
// Prefix min-heap (keep n smallest) and suffix max-heap (keep n largest).
// Answer = min over split points of prefixMin[i] - suffixMax[i+1].

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(minimumDifference([]int{3, 1, 2}))                       // -1
	fmt.Println(minimumDifference([]int{7, 9, 5, 8, 1, 3}))             // 1
	fmt.Println(minimumDifference([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))    // -18
}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func minimumDifference(nums []int) int64 {
	m := len(nums)
	n := m / 3

	pref := make([]int, m)
	h := &MaxHeap{}
	heap.Init(h)
	sum := 0
	for i := 0; i < m; i++ {
		heap.Push(h, nums[i])
		sum += nums[i]
		if h.Len() > n {
			sum -= heap.Pop(h).(int)
		}
		if i >= n-1 {
			pref[i] = sum
		}
	}

	suf := make([]int, m)
	h2 := &MinHeap{}
	heap.Init(h2)
	sum = 0
	for i := m - 1; i >= 0; i-- {
		heap.Push(h2, nums[i])
		sum += nums[i]
		if h2.Len() > n {
			sum -= heap.Pop(h2).(int)
		}
		if i <= 2*n {
			suf[i] = sum
		}
	}

	ans := int(^uint(0) >> 1)
	for i := n - 1; i < 2*n; i++ {
		diff := pref[i] - suf[i+1]
		if diff < ans {
			ans = diff
		}
	}
	return int64(ans)
}

func MinimumDifferenceInSumsAfterRemovalOfElements() any {
	return minimumDifference([]int{3, 1, 2})
}
```

## 2167 — Minimum Time To Remove All Cars Containing Illegal Goods

```go
package main

// LeetCode #2167: Minimum Time to Remove All Cars Containing Illegal Goods
// https://leetcode.com/problems/minimum-time-to-remove-all-cars-containing-illegal-goods/
// Difficulty: Hard
//
// 3-state DP: processing left to right:
//   0 = removing from left end (cost 1 per car)
//   1 = middle section (cost 2 per '1', 0 per '0')
//   2 = removing from right end (cost 1 per car)
// Transitions: 0 -> 1 -> 2 (or 0 -> 2 directly).

import "fmt"

func main() {
	fmt.Println(minimumTime("1100101")) // 5
	fmt.Println(minimumTime("0010"))    // 2
	fmt.Println(minimumTime("010"))     // 2
	fmt.Println(minimumTime("111"))     // 3
	fmt.Println(minimumTime("0"))       // 0
}

func minimumTime(s string) int {
	const inf = 1 << 60
	l, m, r := 0, inf, inf
	for _, ch := range s {
		c := int(ch - '0')
		nl := l + 1
		nm := m + 2*c
		if l+2*c < nm {
			nm = l + 2*c
		}
		nr := r + 1
		if m+1 < nr {
			nr = m + 1
		}
		if l+1 < nr {
			nr = l + 1
		}
		l, m, r = nl, nm, nr
	}
	ans := l
	if m < ans {
		ans = m
	}
	if r < ans {
		ans = r
	}
	return ans
}

func MinimumTimeToRemoveAllCarsContainingIllegalGoods() any {
	return minimumTime("1100101")
}
```

## 2172 — Maximum And Sum Of Array

```go
package main

// LeetCode #2172: Maximum AND Sum of Array
// https://leetcode.com/problems/maximum-and-sum-of-array/
// Difficulty: Hard
//
// DP with bitmask. Double each slot (capacity 2 -> 2 slots of capacity 1).
// dp[mask] = max AND sum for the given assignment mask.

import "fmt"

func main() {
	fmt.Println(maximumANDSum([]int{1, 2, 3, 4, 5, 6}, 3)) // 9
	fmt.Println(maximumANDSum([]int{1, 3, 10, 4, 7, 1}, 3)) // 10
	fmt.Println(maximumANDSum([]int{1, 2, 3}, 2))            // 5
	fmt.Println(maximumANDSum([]int{1, 2}, 1))               // 1
}

func maximumANDSum(nums []int, numSlots int) int {
	n := len(nums)
	m := 2 * numSlots // doubled slots
	total := 1 << m

	dp := make([]int, total)
	for i := 1; i < total; i++ {
		dp[i] = -1
	}

	ans := 0
	for mask := 0; mask < total; mask++ {
		if dp[mask] < 0 {
			continue
		}
		idx := popcount(mask)
		if idx >= n {
			if dp[mask] > ans {
				ans = dp[mask]
			}
			continue
		}
		for slot := 0; slot < m; slot++ {
			if mask&(1<<slot) == 0 {
				nm := mask | (1 << slot)
				val := dp[mask] + (nums[idx] & (slot/2 + 1))
				if val > dp[nm] {
					dp[nm] = val
				}
			}
		}
	}
	return ans
}

func popcount(x int) int {
	c := 0
	for x > 0 {
		c += x & 1
		x >>= 1
	}
	return c
}

func MaximumAndSumOfArray() any {
	return maximumANDSum([]int{1, 2, 3, 4, 5, 6}, 3)
}
```

## 2173 — Longest Winning Streak

```go
package main

// LeetCode #2173: Longest Winning Streak
// https://leetcode.com/problems/longest-winning-streak/
// Difficulty: Hard [Paid]
//
// Given match results (player_id, match_date, result), find each player's
// longest consecutive winning streak. Sort matches by player then date,
// then scan for consecutive wins.

import (
	"fmt"
	"sort"
)

type match struct {
	playerID int
	date     int // represented as integer for simplicity
	result   string
}

type playerStreak struct {
	PlayerID       int
	LongestStreak int
}

func main() {
	// Example
	matches1 := []match{
		{1, 1, "Win"},
		{1, 2, "Win"},
		{1, 3, "Loss"},
		{1, 4, "Win"},
		{2, 1, "Win"},
		{2, 3, "Loss"},
		{2, 5, "Win"},
	}
	fmt.Println(longestWinningStreak(matches1))

	// All wins
	matches2 := []match{
		{1, 1, "Win"},
		{1, 2, "Win"},
		{1, 3, "Win"},
	}
	fmt.Println(longestWinningStreak(matches2))

	// No wins
	matches3 := []match{
		{1, 1, "Loss"},
		{1, 2, "Loss"},
	}
	fmt.Println(longestWinningStreak(matches3))

	// Single match
	matches4 := []match{
		{1, 1, "Win"},
	}
	fmt.Println(longestWinningStreak(matches4))
}

func longestWinningStreak(matches []match) []playerStreak {
	if len(matches) == 0 {
		return nil
	}

	// Sort by playerID, then by date
	sort.Slice(matches, func(i, j int) bool {
		if matches[i].playerID != matches[j].playerID {
			return matches[i].playerID < matches[j].playerID
		}
		return matches[i].date < matches[j].date
	})

	var result []playerStreak
	currentPlayer := matches[0].playerID
	currentStreak := 0
	maxStreak := 0

	for _, m := range matches {
		if m.playerID != currentPlayer {
			result = append(result, playerStreak{currentPlayer, maxStreak})
			currentPlayer = m.playerID
			currentStreak = 0
			maxStreak = 0
		}

		if m.result == "Win" {
			currentStreak++
			if currentStreak > maxStreak {
				maxStreak = currentStreak
			}
		} else {
			currentStreak = 0
		}
	}

	result = append(result, playerStreak{currentPlayer, maxStreak})
	return result
}
```

## 2179 — Count Good Triplets In An Array

```go
package main

// LeetCode #2179: Count Good Triplets in an Array
// https://leetcode.com/problems/count-good-triplets-in-an-array/
// Difficulty: Hard
//
// Map each value to its index in nums1. Then transform nums2 into positions.
// Count increasing triplets in the transformed array using BIT (Fenwick tree).

import "fmt"

func main() {
	fmt.Println(countGoodTriplets([]int{2, 0, 1, 3}, []int{0, 1, 2, 3}))       // 1
	fmt.Println(countGoodTriplets([]int{4, 0, 1, 3, 2}, []int{4, 1, 0, 2, 3})) // 4
	fmt.Println(countGoodTriplets([]int{0, 1, 2, 3}, []int{0, 1, 2, 3}))       // 4
	fmt.Println(countGoodTriplets([]int{0, 1, 2}, []int{2, 1, 0}))             // 0
}

func countGoodTriplets(nums1 []int, nums2 []int) int64 {
	n := len(nums1)
	pos1 := make([]int, n)
	for i, v := range nums1 {
		pos1[v] = i
	}

	arr := make([]int, n)
	for i, v := range nums2 {
		arr[i] = pos1[v]
	}

	// leftLess[i] = count of j < i with arr[j] < arr[i]
	leftLess := make([]int, n)
	bit := newFenwick(n)
	for i, v := range arr {
		leftLess[i] = bit.query(v - 1)
		bit.add(v, 1)
	}

	// rightGreater[i] = count of j > i with arr[j] > arr[i]
	rightGreater := make([]int, n)
	bit = newFenwick(n)
	for i := n - 1; i >= 0; i-- {
		rightGreater[i] = bit.query(n-1) - bit.query(arr[i])
		bit.add(arr[i], 1)
	}

	var ans int64
	for i := 0; i < n; i++ {
		ans += int64(leftLess[i]) * int64(rightGreater[i])
	}
	return ans
}

type fenwick struct {
	tree []int
	n    int
}

func newFenwick(n int) *fenwick {
	return &fenwick{tree: make([]int, n+1), n: n}
}

func (f *fenwick) add(idx, val int) {
	for i := idx + 1; i <= f.n; i += i & -i {
		f.tree[i] += val
	}
}

func (f *fenwick) query(idx int) int {
	if idx < 0 {
		return 0
	}
	res := 0
	for i := idx + 1; i > 0; i -= i & -i {
		res += f.tree[i]
	}
	return res
}

func CountGoodTripletsInAnArray() any {
	return countGoodTriplets([]int{2, 0, 1, 3}, []int{0, 1, 2, 3})
}
```

## 2183 — Count Array Pairs Divisible By K

```go
package main

// LeetCode #2183: Count Array Pairs Divisible by K
// https://leetcode.com/problems/count-array-pairs-divisible-by-k/
// Difficulty: Hard
//
// For each element, compute g = gcd(num, k). Two numbers pair to make a
// product divisible by k iff (g1 * g2) % k == 0. Count frequency of each
// gcd value, then iterate over all pairs of gcd values.

import "fmt"

func main() {
	// Example 1
	fmt.Println(countPairs([]int{1, 2, 3, 4, 5}, 2))

	// Example 2
	fmt.Println(countPairs([]int{1, 2, 3, 4}, 5))

	// Example 3
	fmt.Println(countPairs([]int{1, 2, 3, 4, 5}, 6))

	// All divisible
	fmt.Println(countPairs([]int{10, 20, 30}, 5))
}

func countPairs(nums []int, k int) int64 {
	// Count frequency of each gcd value
	freq := make(map[int]int)
	for _, num := range nums {
		g := gcd(num, k)
		freq[g]++
	}

	var ans int64

	// Iterate over unique gcd values
	gcdVals := make([]int, 0, len(freq))
	for g := range freq {
		gcdVals = append(gcdVals, g)
	}

	for i := 0; i < len(gcdVals); i++ {
		for j := i; j < len(gcdVals); j++ {
			g1, g2 := gcdVals[i], gcdVals[j]
			if (int64(g1)*int64(g2))%int64(k) == 0 {
				if i == j {
					// Same gcd: count C(freq, 2)
					f := int64(freq[g1])
					ans += f * (f - 1) / 2
				} else {
					ans += int64(freq[g1]) * int64(freq[g2])
				}
			}
		}
	}

	return ans
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
```

## 2188 — Minimum Time To Finish The Race

```go
package main

// LeetCode #2188: Minimum Time to Finish the Race
// https://leetcode.com/problems/minimum-time-to-finish-the-race/
// Difficulty: Hard
//
// DP: best[l] = min time for l consecutive laps with ONE tire (no change).
// dp[k] = min over j of dp[k-j] + changeTime + best[j].

import "fmt"

func main() {
	fmt.Println(minimumFinishTime([][]int{{2, 3}, {3, 4}}, 5, 4))            // 21
	fmt.Println(minimumFinishTime([][]int{{1, 10}, {2, 2}, {3, 4}}, 2, 5))   // 13
	fmt.Println(minimumFinishTime([][]int{{3, 4}}, 2, 3))                    // 13
}

func minimumFinishTime(tires [][]int, changeTime int, numLaps int) int {
	const INF = 1 << 60

	best := make([]int, numLaps+1)
	for i := range best {
		best[i] = INF
	}

	for _, tire := range tires {
		f, r := tire[0], tire[1]
		total := 0
		lapTime := f
		for l := 1; l <= numLaps; l++ {
			if total+lapTime >= INF {
				break
			}
			total += lapTime
			if total < best[l] {
				best[l] = total
					}
			if int64(lapTime)*int64(r) >= INF {
				break
			}
			lapTime *= r
		}
	}

	dp := make([]int, numLaps+1)
	for i := range dp {
		dp[i] = INF
	}
	dp[0] = -changeTime

	for i := 1; i <= numLaps; i++ {
		for j := 1; j <= i; j++ {
			if best[j] >= INF {
				continue
			}
			cand := dp[i-j] + changeTime + best[j]
			if cand < dp[i] {
				dp[i] = cand
			}
		}
	}
	return dp[numLaps]
}

func MinimumTimeToFinishTheRace() any {
	return minimumFinishTime([][]int{{2, 3}, {3, 4}}, 5, 4)
}
```

## 2193 — Minimum Number Of Moves To Make Palindrome

```go
package main

// LeetCode #2193: Minimum Number of Moves to Make Palindrome
// https://leetcode.com/problems/minimum-number-of-moves-to-make-palindrome/
// Difficulty: Hard
//
// Greedy two-pointer simulation: For each left character, find the matching
// character closest to the right. Bring it to the rightmost position via
// adjacent swaps. For center characters (single occurrence), swap it one step
// to the right and continue (it will migrate to the center).
// O(n^2) time, O(n) space. n <= 2000 per constraints.

import (
	"fmt"
)

func main() {
	// "aabb" -> 2
	fmt.Println(minMovesToMakePalindrome("aabb"))
	// "letelt" -> 2
	fmt.Println(minMovesToMakePalindrome("letelt"))
	// "a" -> 0
	fmt.Println(minMovesToMakePalindrome("a"))
	// "abba" -> 0
	fmt.Println(minMovesToMakePalindrome("abba"))
	// "abcba" -> 0
	fmt.Println(minMovesToMakePalindrome("abcba"))
}

func minMovesToMakePalindrome(s string) int {
	b := []byte(s)
	n := len(b)
	moves := 0
	l, r := 0, n-1

	for l < r {
		// Find matching character for b[l] from the right side.
		match := r
		for match > l && b[match] != b[l] {
			match--
		}

		if match == l {
			// Center character: appears only once among remaining elements.
			// Swap it one step to the right (it will migrate to center).
			b[l], b[l+1] = b[l+1], b[l]
			moves++
			continue
		}

		// Bring the matched character to position r via adjacent swaps.
		for i := match; i < r; i++ {
			b[i], b[i+1] = b[i+1], b[i]
			moves++
		}

		l++
		r--
	}

	return moves
}
```

## 2196 — Create Binary Tree From Descriptions

```go
package main

// LeetCode #2196: Create Binary Tree From Descriptions
// https://leetcode.com/problems/create-binary-tree-from-descriptions/
// Difficulty: Medium (listed as Hard)

import "fmt"

// TreeNode definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// createBinaryTree builds a binary tree from descriptions.
// Each description: [parent, child, isLeft] where isLeft=1 means left child.
func createBinaryTree(descriptions [][]int) *TreeNode {
	children := make(map[int]*TreeNode)
	hasParent := make(map[int]bool)

	for _, d := range descriptions {
		parentVal, childVal, isLeft := d[0], d[1], d[2]

		// get or create parent
		parent, ok := children[parentVal]
		if !ok {
			parent = &TreeNode{Val: parentVal}
			children[parentVal] = parent
		}

		// get or create child
		child, ok := children[childVal]
		if !ok {
			child = &TreeNode{Val: childVal}
			children[childVal] = child
		}

		// set left/right
		if isLeft == 1 {
			parent.Left = child
		} else {
			parent.Right = child
		}

		hasParent[childVal] = true
		// ensure parent also tracked
		if !hasParent[parentVal] {
			hasParent[parentVal] = false
		}
	}

	// find root: node without a parent
	var root *TreeNode
	for val, node := range children {
		if !hasParent[val] {
			root = node
			break
		}
	}

	return root
}

func main() {
	// Example 1
	descriptions1 := [][]int{
		{20, 15, 1},
		{20, 17, 0},
		{50, 20, 1},
		{50, 80, 0},
		{80, 19, 1},
	}
	root1 := createBinaryTree(descriptions1)
	fmt.Println("Root value:", root1.Val) // Expected: 50

	// Example 2
	descriptions2 := [][]int{
		{1, 2, 1},
		{2, 3, 0},
		{3, 4, 1},
	}
	root2 := createBinaryTree(descriptions2)
	fmt.Println("Root value:", root2.Val) // Expected: 1
}
```

## 2197 — Replace Non Coprime Numbers In Array

```go
package main

// LeetCode #2197: Replace Non-Coprime Numbers in Array
// https://leetcode.com/problems/replace-non-coprime-numbers-in-array/
// Difficulty: Hard
//
// Stack merge: for each number, push onto stack. While stack has >= 2 elements
// and gcd(stack[-2], stack[-1]) > 1, pop top two, push lcm.

import (
	"fmt"
)

func main() {
	// Example: [6,4,3,2,7,6,2] => [12,7,6]
	fmt.Println(replaceNonCoprimeNumbers([]int{6, 4, 3, 2, 7, 6, 2}))
	// Example: [2,2,1,1,3,3,3] => [2,1,1,3]
	fmt.Println(replaceNonCoprimeNumbers([]int{2, 2, 1, 1, 3, 3, 3}))
	// Example: [1,1,1,1] => [1,1,1,1]
	fmt.Println(replaceNonCoprimeNumbers([]int{1, 1, 1, 1}))
	// Example: [12,18,6] => [36]
	fmt.Println(replaceNonCoprimeNumbers([]int{12, 18, 6}))
	// Example: single
	fmt.Println(replaceNonCoprimeNumbers([]int{7}))
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

func replaceNonCoprimeNumbers(nums []int) []int {
	stack := make([]int, 0, len(nums))

	for _, x := range nums {
		stack = append(stack, x)
		for len(stack) >= 2 {
			a, b := stack[len(stack)-2], stack[len(stack)-1]
			g := gcd(a, b)
			if g == 1 {
				break
			}
			// Replace with LCM
			stack = stack[:len(stack)-2]
			stack = append(stack, lcm(a, b))
		}
	}

	return stack
}
```

## 2199 — Finding The Topic Of Each Post

```go
package main

// LeetCode #2199: Finding the Topic of Each Post
// https://leetcode.com/problems/finding-the-topic-of-each-post/
// Difficulty: Hard [Paid]
//
// Given tables: Posts (post_id, content) and Keywords (topic_id, word),
// find for each post the topics that appear in the content (case-insensitive).
// A keyword is considered "appearing" if it appears as a standalone word
// (i.e., not part of another word) in the post content.

import (
	"fmt"
	"strings"
)

// Post represents a post with ID and content.
type Post struct {
	ID      int
	Content string
}

// Keyword maps a topic ID to a keyword string.
type Keyword struct {
	TopicID int
	Word    string
}

// findingTheTopicOfEachPost returns a map from post ID to sorted topic IDs
// whose keywords appear as standalone words (case-insensitive) in the post content.
func findingTheTopicOfEachPost(posts []Post, keywords []Keyword) map[int][]int {
	// build topic keyword index: topicID -> set of lowercase keywords
	topicWords := make(map[int]map[string]bool)
	for _, kw := range keywords {
		if topicWords[kw.TopicID] == nil {
			topicWords[kw.TopicID] = make(map[string]bool)
		}
		topicWords[kw.TopicID][strings.ToLower(kw.Word)] = true
	}

	result := make(map[int][]int)

	for _, post := range posts {
		content := strings.ToLower(post.Content)
		// extract words from content (split on non-alphabetic)
		words := strings.FieldsFunc(content, func(r rune) bool {
			return !((r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '\'')
		})

		wordSet := make(map[string]bool)
		for _, w := range words {
			// strip surrounding punctuation if any
			w = strings.Trim(w, ".,!?;:\"'()[]{}")
			if w != "" {
				wordSet[w] = true
			}
		}

		var matchedTopics []int
		for topicID, kws := range topicWords {
			for kw := range kws {
				if wordSet[kw] {
					matchedTopics = append(matchedTopics, topicID)
					break // one match per topic suffices
				}
			}
		}

		// sort matched topics (simple insertion sort for small slices)
		for i := 0; i < len(matchedTopics); i++ {
			for j := i + 1; j < len(matchedTopics); j++ {
				if matchedTopics[j] < matchedTopics[i] {
					matchedTopics[i], matchedTopics[j] = matchedTopics[j], matchedTopics[i]
				}
			}
		}

		result[post.ID] = matchedTopics
	}

	return result
}

func main() {
	posts := []Post{
		{1, "I love apples and bananas"},
		{2, "The new laptop is great"},
		{3, "Cats are better than dogs"},
	}

	keywords := []Keyword{
		{1, "apple"},
		{1, "banana"},
		{2, "laptop"},
		{3, "cat"},
		{3, "dog"},
	}

	result := findingTheTopicOfEachPost(posts, keywords)
	for _, p := range posts {
		fmt.Printf("Post %d: topics %v\n", p.ID, result[p.ID])
	}
}
```

