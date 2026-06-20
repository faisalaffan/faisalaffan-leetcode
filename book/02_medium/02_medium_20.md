# Medium (Sedang) — Problem 3650–3828

## 3650 — Minimum Cost Path With Edge Reversals

```go
package main

// LeetCode #3650: Minimum Cost Path with Edge Reversals
// https://leetcode.com/problems/minimum-cost-path-with-edge-reversals/
// Difficulty: Medium
// Time: O((n+m) log n) | Space: O(n+m)

import (
	"container/heap"
	"fmt"
	"math"
)

type edge struct {
	to, weight int
}

type minHeap []struct{ node, dist int }

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i].dist < h[j].dist }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(struct{ node, dist int })) }
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func minimumCostPathWithEdgeReversals(n int, edges [][]int) int {
	adj := make([][]edge, n)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		adj[u] = append(adj[u], edge{v, w})
		adj[v] = append(adj[v], edge{u, 2 * w})
	}

	dist := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[0] = 0

	pq := &minHeap{}
	heap.Push(pq, struct{ node, dist int }{0, 0})

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(struct{ node, dist int })
		u, d := cur.node, cur.dist
		if d > dist[u] {
			continue
		}
		if u == n-1 {
			return d
		}
		for _, e := range adj[u] {
			nd := d + e.weight
			if nd < dist[e.to] {
				dist[e.to] = nd
				heap.Push(pq, struct{ node, dist int }{e.to, nd})
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(minimumCostPathWithEdgeReversals(4, [][]int{{0, 1, 5}, {1, 2, 3}, {2, 3, 2}}))
	fmt.Println(minimumCostPathWithEdgeReversals(3, [][]int{{0, 1, 10}, {1, 2, 5}}))
	fmt.Println(minimumCostPathWithEdgeReversals(4, [][]int{{0, 1, 1}, {0, 2, 4}, {1, 2, 1}, {2, 3, 2}}))
}
```

## 3652 — Best Time To Buy And Sell Stock Using Strategy

```go
package main

// LeetCode #3652: Best Time to Buy and Sell Stock using Strategy
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-using-strategy/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func bestTimeToBuyAndSellStockUsingStrategy(prices []int, strategy []int, k int) int64 {
	n := len(prices)
	half := k / 2

	var base int64 = 0
	for i := 0; i < n; i++ {
		base += int64(prices[i] * strategy[i])
	}

	// sm = sum of prices in right half of window
	// so = sum of strategy[i]*prices[i] in full window
	var sm, so, diff int64 = 0, 0, 0
	for i := 0; i < n; i++ {
		sm += int64(prices[i])
		if i >= half {
			sm -= int64(prices[i-half])
		}

		so += int64(prices[i] * strategy[i])
		if i >= k {
			so -= int64(prices[i-k] * strategy[i-k])
		}

		if i+1 >= k {
			delta := sm - so
			if delta > diff {
				diff = delta
			}
		}
	}

	return base + diff
}

func main() {
	fmt.Println(bestTimeToBuyAndSellStockUsingStrategy([]int{1, 2, 3, 4}, []int{-1, 0, 1, 1}, 2))
	fmt.Println(bestTimeToBuyAndSellStockUsingStrategy([]int{5, 3, 2, 6}, []int{-1, 0, 1, -1}, 2))
	fmt.Println(bestTimeToBuyAndSellStockUsingStrategy([]int{1, 2, 3}, []int{-1, 0, 1}, 2))
}
```

## 3653 — Xor After Range Multiplication Queries I

```go
package main

// LeetCode #3653: XOR After Range Multiplication Queries I
// https://leetcode.com/problems/xor-after-range-multiplication-queries-i/
// Difficulty: Medium
// Time: O(n + q) | Space: O(1)

import "fmt"

func xorAfterRangeMultiplicationQueriesI(arr []int, queries [][]int) []int {
	n := len(arr)
	pref := make([]int, n+1)
	for i, v := range arr {
		pref[i+1] = pref[i] ^ v
	}

	ans := make([]int, len(queries))
	for qi, q := range queries {
		l, r, mul := q[0], q[1], q[2]
		// Apply multiplication to the range [l, r]
		for i := l; i <= r; i++ {
			arr[i] *= mul
		}
		// Recompute prefix XOR
		for i := l; i <= r; i++ {
			pref[i+1] = pref[i] ^ arr[i]
		}
		// XOR of range [l, r]
		ans[qi] = pref[r+1] ^ pref[l]
	}
	return ans
}

func main() {
	fmt.Println(xorAfterRangeMultiplicationQueriesI([]int{1, 2, 3}, [][]int{{0, 1, 2}, {1, 2, 3}}))
	fmt.Println(xorAfterRangeMultiplicationQueriesI([]int{5, 7}, [][]int{{0, 0, 2}, {0, 1, 1}}))
}
```

## 3654 — Minimum Sum After Divisible Sum Deletions

```go
package main

// LeetCode #3654: Minimum Sum After Divisible Sum Deletions
// https://leetcode.com/problems/minimum-sum-after-divisible-sum-deletions/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func minimumSumAfterDivisibleSumDeletions(nums []int, k int) int {
	sort.Ints(nums)
	n := len(nums)
	sum := 0
	for i := n/2; i < n; i++ {
		if nums[i]%k == 0 {
			continue
		}
		sum += nums[i]
	}
	for i := 0; i < n/2; i++ {
		if nums[i]%k != 0 {
			sum += nums[i]
		}
	}
	return sum
}

func main() {
	fmt.Println(minimumSumAfterDivisibleSumDeletions([]int{1, 2, 3, 4, 5}, 2))
	fmt.Println(minimumSumAfterDivisibleSumDeletions([]int{2, 4, 6, 8}, 2))
	fmt.Println(minimumSumAfterDivisibleSumDeletions([]int{3, 1, 4, 2}, 3))
}
```

## 3656 — Determine If A Simple Graph Exists

```go
package main

// LeetCode #3656: Determine if a Simple Graph Exists
// https://leetcode.com/problems/determine-if-a-simple-graph-exists/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func determineIfASimpleGraphExists(degrees []int) bool {
	n := len(degrees)
	arr := make([]int, n)
	copy(arr, degrees)
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))

	for i := 0; i < n; i++ {
		if arr[i] == 0 {
			break
		}
		if arr[i] > n-i-1 {
			return false
		}
		for j := i + 1; j <= i+arr[i]; j++ {
			arr[j]--
			if arr[j] < 0 {
				return false
			}
		}
		arr[i] = 0
		sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	}

	return true
}

func main() {
	fmt.Println(determineIfASimpleGraphExists([]int{3, 3, 3, 3}))
	fmt.Println(determineIfASimpleGraphExists([]int{1, 1, 0}))
	fmt.Println(determineIfASimpleGraphExists([]int{1, 1, 1}))
}
```

## 3657 — Find Loyal Customers

```go
package main

// LeetCode #3657: Find Loyal Customers
// https://leetcode.com/problems/find-loyal-customers/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func findLoyalCustomers(purchases [][]int, minPurchases int, minAmount float64) []int {
	customerTotals := make(map[int]int)
	customerCounts := make(map[int]int)

	for _, p := range purchases {
		custID, amount := p[0], p[1]
		customerTotals[custID] += amount
		customerCounts[custID]++
	}

	var loyal []int
	for id := range customerTotals {
		if customerCounts[id] >= minPurchases && float64(customerTotals[id]) >= minAmount {
			loyal = append(loyal, id)
		}
	}

	sort.Ints(loyal)
	return loyal
}

func main() {
	fmt.Println(findLoyalCustomers([][]int{{1, 100}, {2, 50}, {1, 200}, {3, 300}, {2, 150}, {1, 50}}, 2, 200))
	fmt.Println(findLoyalCustomers([][]int{{1, 50}, {1, 50}}, 2, 100))
	fmt.Println(findLoyalCustomers([][]int{{1, 100}, {2, 200}}, 2, 100))
}
```

## 3659 — Partition Array Into K Distinct Groups

```go
package main

// LeetCode #3659: Partition Array Into K-Distinct Groups
// https://leetcode.com/problems/partition-array-into-k-distinct-groups/
// Difficulty: Medium
// Time: O(n) | Space: O(max(nums))

import (
	"fmt"
	"slices"
)

func partitionArrayIntoKDistinctGroups(nums []int, k int) bool {
	n := len(nums)
	if n%k != 0 {
		return false
	}

	maxVal := slices.Max(nums)
	cnt := make([]int, maxVal+1)
	for _, x := range nums {
		cnt[x]++
		if cnt[x] > n/k {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(partitionArrayIntoKDistinctGroups([]int{1, 2, 3, 4}, 2))
	fmt.Println(partitionArrayIntoKDistinctGroups([]int{1, 1, 1, 1}, 2))
	fmt.Println(partitionArrayIntoKDistinctGroups([]int{1, 2, 2, 3, 3, 4}, 3))
}
```

## 3660 — Jump Game Ix

```go
package main

// LeetCode #3660: Jump Game IX
// https://leetcode.com/problems/jump-game-ix/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func jumpGameIx(nums []int) []int {
	n := len(nums)
	pre := make([]int, n)
	pre[0] = nums[0]
	for i := 1; i < n; i++ {
		if nums[i] > pre[i-1] {
			pre[i] = nums[i]
		} else {
			pre[i] = pre[i-1]
		}
	}

	ans := make([]int, n)
	ans[n-1] = pre[n-1]
	minVal := nums[n-1]

	for i := n - 2; i >= 0; i-- {
		if pre[i] > minVal {
			ans[i] = ans[i+1]
		} else {
			ans[i] = pre[i]
		}
		if nums[i] < minVal {
			minVal = nums[i]
		}
	}
	return ans
}

func main() {
	fmt.Println(jumpGameIx([]int{3, 1, 4, 1, 5}))
	fmt.Println(jumpGameIx([]int{1, 2, 3, 4}))
	fmt.Println(jumpGameIx([]int{5, 4, 3, 2, 1}))
}
```

## 3664 — Two Letter Card Game

```go
package main

// LeetCode #3664: Two-Letter Card Game
// https://leetcode.com/problems/two-letter-card-game/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func twoLetterCardGame(cards []string, x byte) int {
	var left [26]int
	var right [26]int
	xx := 0

	for _, c := range cards {
		a, b := c[0], c[1]
		if a == x && b == x {
			xx++
		} else if a == x {
			left[b-'a']++
		} else if b == x {
			right[a-'a']++
		}
	}

	pairsLeft := computePairs(left[:])
	pairsRight := computePairs(right[:])

	leftoverL := sumArr(left[:]) - 2*pairsLeft
	leftoverR := sumArr(right[:]) - 2*pairsRight
	leftovers := leftoverL + leftoverR

	useWithXX := xx
	if useWithXX > leftovers {
		useWithXX = leftovers
	}
	xxLeft := xx - useWithXX

	extra := xxLeft / 2
	totalPairs := pairsLeft + pairsRight
	if extra > totalPairs {
		extra = totalPairs
	}

	return pairsLeft + pairsRight + useWithXX + extra
}

func computePairs(buckets []int) int {
	total := 0
	maxBucket := 0
	for _, v := range buckets {
		total += v
		if v > maxBucket {
			maxBucket = v
		}
	}
	if total < 2 {
		return 0
	}
	half := total / 2
	diff := total - maxBucket
	if half < diff {
		return half
	}
	return diff
}

func sumArr(arr []int) int {
	s := 0
	for _, v := range arr {
		s += v
	}
	return s
}

func main() {
	fmt.Println(twoLetterCardGame([]string{"aa", "ab", "ba", "ac"}, 'a'))
	fmt.Println(twoLetterCardGame([]string{"ab", "bc", "cd"}, 'a'))
	fmt.Println(twoLetterCardGame([]string{"xx", "xa", "xb", "ax", "bx"}, 'x'))
}
```

## 3665 — Twisted Mirror Path Count

```go
package main

// LeetCode #3665: Twisted Mirror Path Count
// https://leetcode.com/problems/twisted-mirror-path-count/
// Difficulty: Medium
// Time: O(m*n) | Space: O(m*n)

import "fmt"

func twistedMirrorPathCount(grid [][]int) int {
	mod := int(1e9 + 7)
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dp[i][j] == 0 {
				continue
			}
			cur := dp[i][j]

			// Try moving right to (i, j+1)
			if j+1 < n {
				if grid[i][j+1] == 1 {
					// Mirror at (i, j+1): reflect down to (i+1, j+1)
					if i+1 < m {
						dp[i+1][j+1] = (dp[i+1][j+1] + cur) % mod
					}
				} else {
					dp[i][j+1] = (dp[i][j+1] + cur) % mod
				}
			}

			// Try moving down to (i+1, j)
			if i+1 < m {
				if grid[i+1][j] == 1 {
					// Mirror at (i+1, j): reflect right to (i+1, j+1)
					if j+1 < n {
						dp[i+1][j+1] = (dp[i+1][j+1] + cur) % mod
					}
				} else {
					dp[i+1][j] = (dp[i+1][j] + cur) % mod
				}
			}
		}
	}

	return dp[m-1][n-1]
}

func main() {
	fmt.Println(twistedMirrorPathCount([][]int{{0, 1, 0}, {0, 0, 1}, {1, 0, 0}}))
	fmt.Println(twistedMirrorPathCount([][]int{{0, 0}, {0, 0}}))
	fmt.Println(twistedMirrorPathCount([][]int{{0, 1}, {1, 0}}))
}
```

## 3669 — Balanced K Factor Decomposition

```go
package main

// LeetCode #3669: Balanced K-Factor Decomposition
// https://leetcode.com/problems/balanced-k-factor-decomposition/
// Difficulty: Medium
// Time: O(d^k) worst case with pruning | Space: O(k)

import (
	"fmt"
	"math"
)

func balancedKFactorDecomposition(n int, k int) []int {
	bestDiff := math.MaxInt32
	var best []int
	path := make([]int, k)

	var dfs func(rem int, start int, depth int)
	dfs = func(rem int, start int, depth int) {
		if depth == k-1 {
			if rem >= start {
				path[depth] = rem
				mn, mx := path[0], path[0]
				for _, v := range path {
					if v < mn {
						mn = v
					}
					if v > mx {
						mx = v
					}
				}
				diff := mx - mn
				if diff < bestDiff {
					bestDiff = diff
					best = make([]int, k)
					copy(best, path)
				}
			}
			return
		}

		for d := start; d*d <= rem; d++ {
			if rem%d == 0 {
				path[depth] = d
				dfs(rem/d, d, depth+1)
			}
		}
	}

	dfs(n, 1, 0)
	return best
}

func main() {
	fmt.Println(balancedKFactorDecomposition(100, 2))
	fmt.Println(balancedKFactorDecomposition(44, 3))
	fmt.Println(balancedKFactorDecomposition(12, 2))
}
```

## 3670 — Maximum Product Of Two Integers With No Common Bits

```go
package main

// LeetCode #3670: Maximum Product of Two Integers With No Common Bits
// https://leetcode.com/problems/maximum-product-of-two-integers-with-no-common-bits/
// Difficulty: Medium
// Time: O(n + B*2^B) | Space: O(2^B)

import (
	"fmt"
	"math/bits"
)

func maximumProductOfTwoIntegersWithNoCommonBits(nums []int) int64 {
	maxVal := 0
	for _, v := range nums {
		if v > maxVal {
			maxVal = v
		}
	}

	maxBits := 32 - bits.LeadingZeros32(uint32(maxVal))
	if maxBits == 0 {
		maxBits = 1
	}
	size := 1 << maxBits

	dp := make([]int, size)
	for _, x := range nums {
		if dp[x] < x {
			dp[x] = x
		}
	}

	// SOS DP
	for b := 0; b < maxBits; b++ {
		half := 1 << b
		step := half << 1
		for base := 0; base < size; base += step {
			upper := base + half
			for m := 0; m < half; m++ {
				u := upper + m
				l := base + m
				if dp[u] < dp[l] {
					dp[u] = dp[l]
				}
			}
		}
	}

	var ans int64 = 0
	full := size - 1
	for _, x := range nums {
		complement := (^x) & full
		y := dp[complement]
		if y > 0 {
			prod := int64(x) * int64(y)
			if prod > ans {
				ans = prod
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(maximumProductOfTwoIntegersWithNoCommonBits([]int{3, 5, 2}))
	fmt.Println(maximumProductOfTwoIntegersWithNoCommonBits([]int{1, 2, 4, 8}))
	fmt.Println(maximumProductOfTwoIntegersWithNoCommonBits([]int{5, 10, 3, 6}))
}
```

## 3672 — Sum Of Weighted Modes In Subarrays

```go
package main

// LeetCode #3672: Sum of Weighted Modes in Subarrays
// https://leetcode.com/problems/sum-of-weighted-modes-in-subarrays/
// Difficulty: Medium [Paid]
// Time: O(n log k) | Space: O(k)

import (
	"container/heap"
	"fmt"
)

type pair struct {
	freq int
	val  int
	idx  int // unique index to break ties in heap
}

type maxHeap []pair

func (h maxHeap) Len() int      { return len(h) }
func (h maxHeap) Less(i, j int) bool {
	if h[i].freq != h[j].freq {
		return h[i].freq > h[j].freq
	}
	return h[i].val < h[j].val
}
func (h maxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x interface{}) { *h = append(*h, x.(pair)) }
func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func sumOfWeightedModesInSubarrays(nums []int, k int) int64 {
	n := len(nums)
	cnt := make(map[int]int)
	pq := &maxHeap{}
	heap.Init(pq)

	add := func(val int, idx int) {
		cnt[val]++
		heap.Push(pq, pair{freq: cnt[val], val: val, idx: idx})
	}

	remove := func(val int) {
		cnt[val]--
	}

	getMode := func() (int, int) {
		for pq.Len() > 0 {
			p := (*pq)[0]
			if cnt[p.val] == p.freq {
				return p.val, p.freq
			}
			heap.Pop(pq)
		}
		return 0, 0
	}

	var ans int64 = 0
	globalIdx := 0

	for i := 0; i < n; i++ {
		add(nums[i], globalIdx)
		globalIdx++

		if i >= k {
			remove(nums[i-k])
		}

		if i >= k-1 {
			val, freq := getMode()
			ans += int64(val) * int64(freq)
		}
	}

	return ans
}

func main() {
	fmt.Println(sumOfWeightedModesInSubarrays([]int{1, 2, 2, 3}, 3))
	fmt.Println(sumOfWeightedModesInSubarrays([]int{1, 1, 1, 1}, 2))
	fmt.Println(sumOfWeightedModesInSubarrays([]int{1, 2, 3}, 1))
}
```

## 3675 — Minimum Operations To Transform String

```go
package main

// LeetCode #3675: Minimum Operations to Transform String
// https://leetcode.com/problems/minimum-operations-to-transform-string/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func minimumOperationsToTransformString(s string) int {
	minChar := byte('z' + 1)
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c != 'a' && c < minChar {
			minChar = c
			if minChar == 'b' {
				break
			}
		}
	}
	if minChar > 'z' {
		return 0
	}
	return int('z' + 1 - minChar)
}

func main() {
	fmt.Println(minimumOperationsToTransformString("yz"))
	fmt.Println(minimumOperationsToTransformString("a"))
	fmt.Println(minimumOperationsToTransformString("abc"))
}
```

## 3676 — Count Bowl Subarrays

```go
package main

// LeetCode #3676: Count Bowl Subarrays
// https://leetcode.com/problems/count-bowl-subarrays/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func countBowlSubarrays(nums []int) int64 {
	var ans int64 = 0
	var stack []int

	for _, num := range nums {
		for len(stack) > 0 && stack[len(stack)-1] < num {
			if len(stack) >= 2 {
				ans++
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, num)
	}

	return ans
}

func main() {
	fmt.Println(countBowlSubarrays([]int{1, 3, 5, 4, 2}))
	fmt.Println(countBowlSubarrays([]int{3, 1, 2, 4}))
	fmt.Println(countBowlSubarrays([]int{1, 2, 3, 4}))
}
```

## 3679 — Minimum Discards To Balance Inventory

```go
package main

// LeetCode #3679: Minimum Discards to Balance Inventory
// https://leetcode.com/problems/minimum-discards-to-balance-inventory/
// Difficulty: Medium
// Time: O(n) | Space: O(max(arrivals))

import "fmt"

func minimumDiscardsToBalanceInventory(arrivals []int, w int, m int) int {
	maxVal := 0
	for _, v := range arrivals {
		if v > maxVal {
			maxVal = v
		}
	}

	cnt := make([]int, maxVal+1)
	ans := 0

	for i, x := range arrivals {
		if cnt[x] == m {
			arrivals[i] = 0
			ans++
		} else {
			cnt[x]++
		}

		left := i + 1 - w
		if left >= 0 {
			cnt[arrivals[left]]--
		}
	}

	return ans
}

func main() {
	fmt.Println(minimumDiscardsToBalanceInventory([]int{1, 2, 3, 3, 3, 4}, 3, 2))
	fmt.Println(minimumDiscardsToBalanceInventory([]int{1, 1, 1, 2, 2}, 2, 2))
	fmt.Println(minimumDiscardsToBalanceInventory([]int{1, 2, 3}, 3, 1))
}
```

## 3680 — Generate Schedule

```go
package main

// LeetCode #3680: Generate Schedule
// https://leetcode.com/problems/generate-schedule/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n^2)

import "fmt"

func generateSchedule(n int) [][]int {
	if n < 5 {
		return [][]int{}
	}

	totalMatches := n * (n - 1)
	schedule := make([][]int, 0, totalMatches)

	// Phase 1: offset 2 to n-2
	for offset := 2; offset < n-1; offset++ {
		for team := 0; team < n; team++ {
			schedule = append(schedule, []int{team, (team + offset) % n})
		}
	}

	// Phase 2: wrap-around pairs
	for team := 0; team < n; team++ {
		schedule = append(schedule, []int{team, (team + 1) % n})
		schedule = append(schedule, []int{(team + 4) % n, (team + 3) % n})
	}

	return schedule
}

func main() {
	fmt.Println(len(generateSchedule(5)))
	fmt.Println(len(generateSchedule(3)))
	fmt.Println(len(generateSchedule(6)))
}
```

## 3682 — Minimum Index Sum Of Common Elements

```go
package main

// LeetCode #3682: Minimum Index Sum of Common Elements
// https://leetcode.com/problems/minimum-index-sum-of-common-elements/
// Difficulty: Medium [Paid]
// Time: O(n + m) | Space: O(n)

import "fmt"

func minimumIndexSumOfCommonElements(nums1 []int, nums2 []int) int {
	idxMap := make(map[int]int)
	for i, v := range nums2 {
		idxMap[v] = i
	}

	minSum := int(^uint(0) >> 1)
	found := false

	for i, v := range nums1 {
		if j, ok := idxMap[v]; ok {
			sum := i + j
			if !found || sum < minSum {
				minSum = sum
				found = true
			}
		}
	}

	if !found {
		return -1
	}
	return minSum
}

func main() {
	fmt.Println(minimumIndexSumOfCommonElements([]int{3, 2, 1}, []int{1, 3, 1}))
	fmt.Println(minimumIndexSumOfCommonElements([]int{5, 1, 2}, []int{2, 1, 3}))
	fmt.Println(minimumIndexSumOfCommonElements([]int{6, 4}, []int{7, 8}))
}
```

## 3685 — Subsequence Sum After Capping Elements

```go
package main

// LeetCode #3685: Subsequence Sum After Capping Elements
// https://leetcode.com/problems/subsequence-sum-after-capping-elements/
// Difficulty: Medium
// Time: O(n*k + n^2) | Space: O(k + n)

import "fmt"

func subsequenceSumAfterCappingElements(nums []int, k int) []bool {
	n := len(nums)
	ans := make([]bool, n)

	dp := make([]bool, k+1)
	dp[0] = true

	maxVal := n
	cnt := make([]int, maxVal+1)
	for _, v := range nums {
		if v <= maxVal {
			cnt[v]++
		} else {
			cnt[maxVal]++
		}
	}

	cntGe := make([]int, maxVal+2)
	for x := maxVal; x >= 1; x-- {
		cntGe[x] = cntGe[x+1] + cnt[x]
	}

	for x := 1; x <= n; x++ {
		ge := cntGe[x]
		ok := false
		maxM := ge
		if maxM > k/x {
			maxM = k / x
		}
		for m := 0; m <= maxM; m++ {
			if dp[k-m*x] {
				ok = true
				break
			}
		}
		ans[x-1] = ok

		c := cnt[x]
		if c == 0 {
			continue
		}
		p := 1
		for c > 0 {
			take := p
			if c < take {
				take = c
			}
			w := take * x
			for s := k; s >= w; s-- {
				if !dp[s] && dp[s-w] {
					dp[s] = true
				}
			}
			c -= take
			p <<= 1
		}
	}

	return ans
}

func main() {
	fmt.Println(subsequenceSumAfterCappingElements([]int{1, 2, 3}, 3))
	fmt.Println(subsequenceSumAfterCappingElements([]int{2, 4, 6}, 6))
	fmt.Println(subsequenceSumAfterCappingElements([]int{1, 1, 1}, 2))
}
```

## 3689 — Maximum Total Subarray Value I

```go
package main

// LeetCode #3689: Maximum Total Subarray Value I
// https://leetcode.com/problems/maximum-total-subarray-value-i/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"
import "math"

func maximumTotalSubarrayValueI(nums []int, k int) int64 {
	minVal := math.MaxInt32
	maxVal := math.MinInt32
	for _, v := range nums {
		if v < minVal {
			minVal = v
		}
		if v > maxVal {
			maxVal = v
		}
	}
	return int64(maxVal-minVal) * int64(k)
}

func main() {
	fmt.Println(maximumTotalSubarrayValueI([]int{1, 3, 2}, 2))
	fmt.Println(maximumTotalSubarrayValueI([]int{4, 2, 5, 1}, 3))
	fmt.Println(maximumTotalSubarrayValueI([]int{1, 1, 1}, 5))
}
```

## 3690 — Split And Merge Array Transformation

```go
package main

// LeetCode #3690: Split and Merge Array Transformation
// https://leetcode.com/problems/split-and-merge-array-transformation/
// Difficulty: Medium
// Time: O(n! * n^4) | Space: O(n! * n)

import "fmt"

func splitAndMergeArrayTransformation(nums1 []int, nums2 []int) int {
	n := len(nums1)
	target := make([]int, n)
	copy(target, nums2)

	type state struct {
		arr   []int
		steps int
	}

	queue := []state{{arr: append([]int(nil), nums1...), steps: 0}}
	visited := make(map[string]bool)

	key := func(arr []int) string {
		b := make([]byte, len(arr)*4)
		for i, v := range arr {
			b[i*4] = byte(v >> 24)
			b[i*4+1] = byte(v >> 16)
			b[i*4+2] = byte(v >> 8)
			b[i*4+3] = byte(v)
		}
		return string(b)
	}

	visited[key(nums1)] = true

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if equal(cur.arr, target) {
			return cur.steps
		}

		// Try all subarrays [L, R]
		for L := 0; L < n; L++ {
			for R := L; R < n; R++ {
				sub := make([]int, R-L+1)
				copy(sub, cur.arr[L:R+1])

				remain := make([]int, 0, n-(R-L+1))
				remain = append(remain, cur.arr[:L]...)
				remain = append(remain, cur.arr[R+1:]...)

				// Insert sub at all positions in remain
				for pos := 0; pos <= len(remain); pos++ {
					next := make([]int, 0, n)
					next = append(next, remain[:pos]...)
					next = append(next, sub...)
					next = append(next, remain[pos:]...)

					k := key(next)
					if !visited[k] {
						visited[k] = true
						queue = append(queue, state{arr: next, steps: cur.steps + 1})
					}
				}
			}
		}
	}

	return -1
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(splitAndMergeArrayTransformation([]int{3, 1, 2}, []int{1, 2, 3}))
	fmt.Println(splitAndMergeArrayTransformation([]int{1, 2, 3}, []int{1, 2, 3}))
	fmt.Println(splitAndMergeArrayTransformation([]int{2, 1}, []int{1, 2}))
}
```

## 3693 — Climbing Stairs Ii

```go
package main

// LeetCode #3693: Climbing Stairs II
// https://leetcode.com/problems/climbing-stairs-ii/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func climbingStairsIi(n int, costs []int) int {
	dp0, dp1, dp2 := 0, 0, 0
	for j := 0; j < n; j++ {
		cur := dp2 + costs[j] + 1
		if j >= 1 {
			cand := dp1 + costs[j] + 4
			if cand < cur {
				cur = cand
			}
		}
		if j >= 2 {
			cand := dp0 + costs[j] + 9
			if cand < cur {
				cur = cand
			}
		}
		dp0, dp1, dp2 = dp1, dp2, cur
	}
	return dp2
}

func main() {
	fmt.Println(climbingStairsIi(3, []int{1, 2, 3}))
	fmt.Println(climbingStairsIi(2, []int{5, 10}))
	fmt.Println(climbingStairsIi(5, []int{1, 1, 1, 1, 1}))
}
```

## 3694 — Distinct Points Reachable After Substring Removal

```go
package main

// LeetCode #3694: Distinct Points Reachable After Substring Removal
// https://leetcode.com/problems/distinct-points-reachable-after-substring-removal/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func distinctPointsReachableAfterSubstringRemoval(s string, k int) int {
	n := len(s)
	// prefix sums for x and y
	fx := make([]int, n+1)
	fy := make([]int, n+1)
	for i, ch := range s {
		fx[i+1] = fx[i]
		fy[i+1] = fy[i]
		switch ch {
		case 'U':
			fy[i+1]++
		case 'D':
			fy[i+1]--
		case 'L':
			fx[i+1]--
		case 'R':
			fx[i+1]++
		}
	}

	seen := make(map[[2]int]bool)
	for i := k; i <= n; i++ {
		x := fx[n] - (fx[i] - fx[i-k])
		y := fy[n] - (fy[i] - fy[i-k])
		seen[[2]int{x, y}] = true
	}

	return len(seen)
}

func main() {
	fmt.Println(distinctPointsReachableAfterSubstringRemoval("LUL", 1))
	fmt.Println(distinctPointsReachableAfterSubstringRemoval("UDLR", 4))
	fmt.Println(distinctPointsReachableAfterSubstringRemoval("UU", 1))
}
```

## 3698 — Split Array With Minimum Difference

```go
package main

// LeetCode #3698: Split Array With Minimum Difference
// https://leetcode.com/problems/split-array-with-minimum-difference/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"math"
)

func splitArrayWithMinimumDifference(nums []int) int64 {
	n := len(nums)
	prefix := make([]int64, n)
	prefix[0] = int64(nums[0])
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] + int64(nums[i])
	}
	total := prefix[n-1]

	inc := make([]bool, n)
	inc[0] = true
	for i := 1; i < n; i++ {
		inc[i] = inc[i-1] && nums[i] > nums[i-1]
	}

	dec := make([]bool, n)
	dec[n-1] = true
	for i := n - 2; i >= 0; i-- {
		dec[i] = dec[i+1] && nums[i] > nums[i+1]
	}

	minDiff := int64(math.MaxInt64)
	for i := 0; i < n-1; i++ {
		if inc[i] && dec[i+1] {
			leftSum := prefix[i]
			rightSum := total - prefix[i]
			diff := leftSum - rightSum
			if diff < 0 {
				diff = -diff
			}
			if diff < minDiff {
				minDiff = diff
			}
		}
	}

	if minDiff == int64(math.MaxInt64) {
		return -1
	}
	return minDiff
}

func main() {
	fmt.Println(splitArrayWithMinimumDifference([]int{1, 3, 2}))
	fmt.Println(splitArrayWithMinimumDifference([]int{1, 2, 4, 3}))
	fmt.Println(splitArrayWithMinimumDifference([]int{3, 1, 2}))
}
```

## 3702 — Longest Subsequence With Non Zero Bitwise Xor

```go
package main

// LeetCode #3702: Longest Subsequence With Non-Zero Bitwise XOR
// https://leetcode.com/problems/longest-subsequence-with-non-zero-bitwise-xor/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func longestSubsequenceWithNonZeroBitwiseXor(nums []int) int {
	xorSum := 0
	allZero := true
	for _, num := range nums {
		xorSum ^= num
		if num != 0 {
			allZero = false
		}
	}
	if allZero {
		return 0
	}
	if xorSum != 0 {
		return len(nums)
	}
	return len(nums) - 1
}

func main() {
	fmt.Println(longestSubsequenceWithNonZeroBitwiseXor([]int{1, 2, 3}))
	fmt.Println(longestSubsequenceWithNonZeroBitwiseXor([]int{2, 3, 4}))
	fmt.Println(longestSubsequenceWithNonZeroBitwiseXor([]int{0, 0, 0}))
}
```

## 3703 — Remove K Balanced Substrings

```go
package main

// LeetCode #3703: Remove K-Balanced Substrings
// https://leetcode.com/problems/remove-k-balanced-substrings/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

type pair struct {
	ch    byte
	count int
}

func removeKBalancedSubstrings(s string, k int) string {
	var stack []pair

	for i := 0; i < len(s); i++ {
		c := s[i]
		if len(stack) > 0 && stack[len(stack)-1].ch == c {
			stack[len(stack)-1].count++
		} else {
			stack = append(stack, pair{ch: c, count: 1})
		}

		// Check for k-balanced pattern
		if c == ')' && stack[len(stack)-1].count == k {
			if len(stack) >= 2 {
				prev := &stack[len(stack)-2]
				if prev.ch == '(' && prev.count >= k {
					// Remove closing part
					stack = stack[:len(stack)-1]
					// Remove or reduce opening part
					if prev.count == k {
						stack = stack[:len(stack)-1]
					} else {
						prev.count -= k
					}
				}
			}
		}
	}

	res := make([]byte, 0, len(s))
	for _, p := range stack {
		for j := 0; j < p.count; j++ {
			res = append(res, p.ch)
		}
	}
	return string(res)
}

func main() {
	fmt.Println(removeKBalancedSubstrings("(())", 1))
	fmt.Println(removeKBalancedSubstrings("(()(", 1))
	fmt.Println(removeKBalancedSubstrings("((()))()()()", 3))
}
```

## 3705 — Find Golden Hour Customers

```go
package main

// LeetCode #3705: Find Golden Hour Customers
// https://leetcode.com/problems/find-golden-hour-customers/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

type order struct {
	customerID int
	hour       int
	rating     int // 0 means no rating
	hasRating  bool
}

type customerResult struct {
	customerID         int
	totalOrders        int
	peakHourPercentage int
	averageRating      float64
}

func findGoldenHourCustomers(orders []order) []customerResult {
	type stats struct {
		total     int
		peakCount int
		ratingSum int
		rated     int
	}

	custStats := make(map[int]*stats)
	for _, o := range orders {
		if _, ok := custStats[o.customerID]; !ok {
			custStats[o.customerID] = &stats{}
		}
		s := custStats[o.customerID]
		s.total++
		if (o.hour >= 11 && o.hour <= 13) || (o.hour >= 18 && o.hour <= 20) {
			s.peakCount++
		}
		if o.hasRating {
			s.ratingSum += o.rating
			s.rated++
		}
	}

	var results []customerResult
	for id, s := range custStats {
		if s.total < 3 {
			continue
		}
		peakPct := s.peakCount * 100 / s.total
		if peakPct < 60 {
			continue
		}
		if s.rated == 0 || s.rated*2 < s.total {
			continue
		}
		avgRating := float64(s.ratingSum) / float64(s.rated)
		if avgRating < 4.0 {
			continue
		}
		avgRounded := float64(int(avgRating*100+0.5)) / 100
		results = append(results, customerResult{
			customerID:         id,
			totalOrders:        s.total,
			peakHourPercentage: peakPct,
			averageRating:      avgRounded,
		})
	}

	sort.Slice(results, func(i, j int) bool {
		if results[i].averageRating != results[j].averageRating {
			return results[i].averageRating > results[j].averageRating
		}
		return results[i].customerID > results[j].customerID
	})

	return results
}

func main() {
	orders := []order{
		{101, 12, 5, true},
		{101, 13, 4, true},
		{101, 19, 5, true},
		{102, 10, 3, true},
		{102, 14, 4, true},
		{103, 11, 5, true},
		{103, 12, 4, true},
		{103, 18, 5, true},
		{105, 11, 4, true},
		{105, 19, 5, true},
		{105, 20, 4, true},
	}
	results := findGoldenHourCustomers(orders)
	for _, r := range results {
		fmt.Printf("Customer %d: orders=%d peak=%d%% rating=%.2f\n", r.customerID, r.totalOrders, r.peakHourPercentage, r.averageRating)
	}
	if len(results) == 0 {
		fmt.Println("[]")
	}
}
```

## 3706 — Maximum Distance Between Unequal Words In Array Ii

```go
package main

// LeetCode #3706: Maximum Distance Between Unequal Words in Array II
// https://leetcode.com/problems/maximum-distance-between-unequal-words-in-array-ii/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func maximumDistanceBetweenUnequalWordsInArrayIi(words []string) int {
	n := len(words)
	ans := 0
	for i := 0; i < n; i++ {
		if words[i] != words[0] {
			if i+1 > ans {
				ans = i + 1
			}
		}
		if words[i] != words[n-1] {
			if n-i > ans {
				ans = n - i
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(maximumDistanceBetweenUnequalWordsInArrayIi([]string{"leetcode", "leetcode", "codeforces"}))
	fmt.Println(maximumDistanceBetweenUnequalWordsInArrayIi([]string{"a", "b", "c", "a", "a"}))
	fmt.Println(maximumDistanceBetweenUnequalWordsInArrayIi([]string{"z", "z", "z"}))
}
```

## 3708 — Longest Fibonacci Subarray

```go
package main

// LeetCode #3708: Longest Fibonacci Subarray
// https://leetcode.com/problems/longest-fibonacci-subarray/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func longestFibonacciSubarray(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}
	ans := 2
	cnt := 2
	for i := 2; i < n; i++ {
		if nums[i] == nums[i-1]+nums[i-2] {
			cnt++
		} else {
			cnt = 2
		}
		if cnt > ans {
			ans = cnt
		}
	}
	return ans
}

func main() {
	fmt.Println(longestFibonacciSubarray([]int{1, 1, 2, 3, 5, 8}))
	fmt.Println(longestFibonacciSubarray([]int{1, 3, 5, 8}))
	fmt.Println(longestFibonacciSubarray([]int{1, 2, 3, 4, 5, 6}))
}
```

## 3709 — Design Exam Scores Tracker

```go
package main

// LeetCode #3709: Design Exam Scores Tracker
// https://leetcode.com/problems/design-exam-scores-tracker/
// Difficulty: Medium
// Time: O(log n) per query | Space: O(n)

import (
	"fmt"
	"sort"
)

type ExamScoresTracker struct {
	times []int
	pref  []int64
}

func NewExamScoresTracker() *ExamScoresTracker {
	return &ExamScoresTracker{
		times: []int{0},
		pref:  []int64{0},
	}
}

func (t *ExamScoresTracker) Record(time int, score int) {
	t.times = append(t.times, time)
	t.pref = append(t.pref, t.pref[len(t.pref)-1]+int64(score))
}

func (t *ExamScoresTracker) TotalScore(startTime int, endTime int) int64 {
	l := sort.SearchInts(t.times, startTime)
	if l < len(t.times) && t.times[l] < startTime {
		l++
	}
	l--

	r := sort.SearchInts(t.times, endTime+1) - 1

	if l < 0 {
		l = 0
	}
	if r < 0 {
		return 0
	}
	if r >= len(t.pref) {
		r = len(t.pref) - 1
	}
	return t.pref[r] - t.pref[l]
}

func main() {
	tracker := NewExamScoresTracker()
	tracker.Record(10, 5)
	tracker.Record(20, 3)
	tracker.Record(30, 7)
	fmt.Println(tracker.TotalScore(10, 20))
	fmt.Println(tracker.TotalScore(15, 25))
	fmt.Println(tracker.TotalScore(5, 35))

	tracker2 := NewExamScoresTracker()
	fmt.Println(tracker2.TotalScore(0, 100))
}
```

## 3711 — Maximum Transactions Without Negative Balance

```go
package main

// LeetCode #3711: Maximum Transactions Without Negative Balance
// https://leetcode.com/problems/maximum-transactions-without-negative-balance/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"container/heap"
	"fmt"
)

type minHeapInt []int

func (h minHeapInt) Len() int           { return len(h) }
func (h minHeapInt) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeapInt) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeapInt) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *minHeapInt) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func maximumTransactionsWithoutNegativeBalance(transactions []int) int {
	h := &minHeapInt{}
	heap.Init(h)
	s := 0
	ans := len(transactions)

	for _, x := range transactions {
		s += x
		heap.Push(h, x)
		for s < 0 {
			y := heap.Pop(h).(int)
			s -= y
			ans--
		}
	}
	return ans
}

func main() {
	fmt.Println(maximumTransactionsWithoutNegativeBalance([]int{2, -5, 3, -1, -2}))
	fmt.Println(maximumTransactionsWithoutNegativeBalance([]int{-1, -2, -3}))
	fmt.Println(maximumTransactionsWithoutNegativeBalance([]int{3, -2, 3, -2, 1, -1}))
}
```

## 3713 — Longest Balanced Substring I

```go
package main

// LeetCode #3713: Longest Balanced Substring I
// https://leetcode.com/problems/longest-balanced-substring-i/
// Difficulty: Medium
// Time: O(n^2) | Space: O(1)

import "fmt"

func longestBalancedSubstringI(s string) int {
	n := len(s)
	ans := 0
	for i := 0; i < n; i++ {
		var cnt [26]int
		distinct := 0
		maxFreq := 0
		for j := i; j < n; j++ {
			idx := s[j] - 'a'
			cnt[idx]++
			if cnt[idx] == 1 {
				distinct++
			}
			if cnt[idx] > maxFreq {
				maxFreq = cnt[idx]
			}
			if maxFreq*distinct == j-i+1 {
				if j-i+1 > ans {
					ans = j - i + 1
				}
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(longestBalancedSubstringI("abbac"))
	fmt.Println(longestBalancedSubstringI("zzabccy"))
	fmt.Println(longestBalancedSubstringI("aba"))
}
```

## 3714 — Longest Balanced Substring Ii

```go
package main

// LeetCode #3714: Longest Balanced Substring II
// https://leetcode.com/problems/longest-balanced-substring-ii/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func longestBalancedSubstringIi(s string) int {
	n := len(s)
	ans := 0

	// Case 1: single char run
	i := 0
	for i < n {
		start := i
		for i < n && s[i] == s[start] {
			i++
		}
		if i-start > ans {
			ans = i - start
		}
	}

	// Case 2: exactly two chars
	solvePair := func(x, y byte) {
		i := 0
		for i < n {
			pos := map[int]int{0: i - 1}
			d := 0
			for i < n && (s[i] == x || s[i] == y) {
				if s[i] == x {
					d++
				} else {
					d--
				}
				if firstIdx, ok := pos[d]; ok {
					if i-firstIdx > ans {
						ans = i - firstIdx
					}
				} else {
					pos[d] = i
				}
				i++
			}
			if i < n {
				i++ // skip third char
			}
		}
	}

	solvePair('a', 'b')
	solvePair('a', 'c')
	solvePair('b', 'c')

	// Case 3: all three chars
	pos3 := make(map[[2]int]int)
	pos3[[2]int{0, 0}] = -1
	cnt := [3]int{} // a, b, c
	for i := 0; i < n; i++ {
		switch s[i] {
		case 'a':
			cnt[0]++
		case 'b':
			cnt[1]++
		case 'c':
			cnt[2]++
		}
		state := [2]int{cnt[0] - cnt[1], cnt[1] - cnt[2]}
		if firstIdx, ok := pos3[state]; ok {
			if i-firstIdx > ans {
				ans = i - firstIdx
			}
		} else {
			pos3[state] = i
		}
	}

	return ans
}

func main() {
	fmt.Println(longestBalancedSubstringIi("abcabc"))
	fmt.Println(longestBalancedSubstringIi("aabbcc"))
	fmt.Println(longestBalancedSubstringIi("aaa"))
}
```

## 3716 — Find Churn Risk Customers

```go
package main

// LeetCode #3716: Find Churn Risk Customers
// https://leetcode.com/problems/find-churn-risk-customers/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

type subEvent struct {
	userID        int
	eventDate     int
	eventType     string // "start", "cancel", "downgrade"
	planName      string
	monthlyAmount float64
}

type churnResult struct {
	userID               int
	currentPlan          string
	currentMonthlyAmount float64
	maxHistoricalAmount  float64
	daysAsSubscriber     int
}

func findChurnRiskCustomers(events []subEvent) []churnResult {
	// Group events by user
	type userData struct {
		events         []subEvent
		startDate      int
		lastDate       int
		maxAmount      float64
		downgradeCount int
	}

	userMap := make(map[int]*userData)
	for _, e := range events {
		if _, ok := userMap[e.userID]; !ok {
			userMap[e.userID] = &userData{}
		}
		u := userMap[e.userID]
		u.events = append(u.events, e)
		if e.eventType == "start" && (u.startDate == 0 || e.eventDate < u.startDate) {
			u.startDate = e.eventDate
		}
		if e.eventDate > u.lastDate {
			u.lastDate = e.eventDate
		}
		if e.monthlyAmount > u.maxAmount {
			u.maxAmount = e.monthlyAmount
		}
		if e.eventType == "downgrade" {
			u.downgradeCount++
		}
	}

	var results []churnResult
	for uid, u := range userMap {
		// Sort events by date
		sort.Slice(u.events, func(i, j int) bool {
			return u.events[i].eventDate < u.events[j].eventDate
		})

		lastEvent := u.events[len(u.events)-1]
		if lastEvent.eventType == "cancel" {
			continue
		}
		if u.downgradeCount == 0 {
			continue
		}
		if lastEvent.monthlyAmount*2 >= u.maxAmount {
			continue
		}
		days := u.lastDate - u.startDate
		if days < 60 {
			continue
		}

		results = append(results, churnResult{
			userID:               uid,
			currentPlan:          lastEvent.planName,
			currentMonthlyAmount: lastEvent.monthlyAmount,
			maxHistoricalAmount:  u.maxAmount,
			daysAsSubscriber:     days,
		})
	}

	sort.Slice(results, func(i, j int) bool {
		if results[i].daysAsSubscriber != results[j].daysAsSubscriber {
			return results[i].daysAsSubscriber > results[j].daysAsSubscriber
		}
		return results[i].userID < results[j].userID
	})

	return results
}

func main() {
	events := []subEvent{
		{501, 1, "start", "Basic", 9.99},
		{501, 30, "downgrade", "Lite", 4.99},
		{501, 60, "downgrade", "Free", 0},
		{502, 10, "start", "Pro", 29.99},
		{502, 40, "downgrade", "Basic", 9.99},
		{503, 5, "start", "Pro", 29.99},
		{504, 1, "start", "Basic", 9.99},
		{504, 30, "cancel", "Basic", 0},
		{506, 1, "start", "Basic", 9.99},
	}
	results := findChurnRiskCustomers(events)
	for _, r := range results {
		fmt.Printf("User %d: plan=%s current=%.2f max=%.2f days=%d\n",
			r.userID, r.currentPlan, r.currentMonthlyAmount, r.maxHistoricalAmount, r.daysAsSubscriber)
	}
	if len(results) == 0 {
		fmt.Println("[]")
	}
}
```

## 3717 — Minimum Operations To Make The Array Beautiful

```go
package main

// LeetCode #3717: Minimum Operations to Make the Array Beautiful
// https://leetcode.com/problems/minimum-operations-to-make-the-array-beautiful/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func minimumOperationsToMakeTheArrayBeautiful(nums []int) int {
	ops := 0
	prev := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]%prev != 0 {
			target := ((nums[i] / prev) + 1) * prev
			ops += target - nums[i]
			prev = target
		} else {
			prev = nums[i]
		}
	}
	return ops
}

func main() {
	fmt.Println(minimumOperationsToMakeTheArrayBeautiful([]int{3, 7, 9}))
	fmt.Println(minimumOperationsToMakeTheArrayBeautiful([]int{1, 1, 1}))
	fmt.Println(minimumOperationsToMakeTheArrayBeautiful([]int{2, 3, 5}))
}
```

## 3719 — Longest Balanced Subarray I

```go
package main

// LeetCode #3719: Longest Balanced Subarray I
// https://leetcode.com/problems/longest-balanced-subarray-i/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import "fmt"

func longestBalancedSubarrayI(nums []int) int {
	n := len(nums)
	ans := 0

	for i := 0; i < n; i++ {
		if n-i <= ans {
			break
		}
		evenVisited := make(map[int]bool)
		oddVisited := make(map[int]bool)
		evenCount := 0
		oddCount := 0

		for j := i; j < n; j++ {
			if nums[j]%2 == 0 {
				if !evenVisited[nums[j]] {
					evenVisited[nums[j]] = true
					evenCount++
				}
			} else {
				if !oddVisited[nums[j]] {
					oddVisited[nums[j]] = true
					oddCount++
				}
			}
			if evenCount == oddCount {
				if j-i+1 > ans {
					ans = j - i + 1
				}
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(longestBalancedSubarrayI([]int{1, 2, 3, 4}))
	fmt.Println(longestBalancedSubarrayI([]int{2, 4, 6, 8}))
	fmt.Println(longestBalancedSubarrayI([]int{1, 1, 2, 2, 3, 3}))
}
```

## 3720 — Lexicographically Smallest Permutation Greater Than Target

```go
package main

// LeetCode #3720: Lexicographically Smallest Permutation Greater Than Target
// https://leetcode.com/problems/lexicographically-smallest-permutation-greater-than-target/
// Difficulty: Medium
// Time: O(n*26) | Space: O(26)

import "fmt"

func lexicographicallySmallestPermutationGreaterThanTarget(s string, target string) string {
	n := len(s)
	var freq [26]int
	for i := 0; i < n; i++ {
		freq[s[i]-'a']++
	}

	var ans []byte

	var dfs func(idx int, check bool) bool
	dfs = func(idx int, check bool) bool {
		if idx == n {
			return check
		}
		for ch := 0; ch < 26; ch++ {
			if freq[ch] == 0 {
				continue
			}
			if !check && byte(ch)+'a' < target[idx] {
				continue
			}
			freq[ch]--
			ans = append(ans, byte(ch)+'a')
			nextCheck := check || byte(ch)+'a' > target[idx]
			if dfs(idx+1, nextCheck) {
				return true
			}
			ans = ans[:len(ans)-1]
			freq[ch]++
		}
		return false
	}

	if dfs(0, false) {
		return string(ans)
	}
	return ""
}

func main() {
	fmt.Println(lexicographicallySmallestPermutationGreaterThanTarget("abc", "bba"))
	fmt.Println(lexicographicallySmallestPermutationGreaterThanTarget("leet", "code"))
	fmt.Println(lexicographicallySmallestPermutationGreaterThanTarget("baba", "bbaa"))
}
```

## 3722 — Lexicographically Smallest String After Reverse

```go
package main

// LeetCode #3722: Lexicographically Smallest String After Reverse
// https://leetcode.com/problems/lexicographically-smallest-string-after-reverse/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import "fmt"

func reverseSubstring(s string, start int, end int) string {
	b := []byte(s)
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func lexicographicallySmallestStringAfterReverse(s string) string {
	n := len(s)
	best := s

	// Reverse first k
	for k := 1; k <= n; k++ {
		candidate := reverseSubstring(s, 0, k-1)
		if candidate < best {
			best = candidate
		}
	}

	// Reverse last k
	for k := 1; k <= n; k++ {
		candidate := reverseSubstring(s, n-k, n-1)
		if candidate < best {
			best = candidate
		}
	}

	return best
}

func main() {
	fmt.Println(lexicographicallySmallestStringAfterReverse("dcab"))
	fmt.Println(lexicographicallySmallestStringAfterReverse("abba"))
	fmt.Println(lexicographicallySmallestStringAfterReverse("zxy"))
}
```

## 3723 — Maximize Sum Of Squares Of Digits

```go
package main

// LeetCode #3723: Maximize Sum of Squares of Digits
// https://leetcode.com/problems/maximize-sum-of-squares-of-digits/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"strings"
)

func maximizeSumOfSquaresOfDigits(num int, total int) string {
	if num*9 < total {
		return ""
	}
	nines := total / 9
	rem := total % 9
	var sb strings.Builder
	sb.WriteString(strings.Repeat("9", nines))
	if rem > 0 {
		sb.WriteByte(byte(rem) + '0')
	}
	for sb.Len() < num {
		sb.WriteByte('0')
	}
	return sb.String()
}

func main() {
	fmt.Println(maximizeSumOfSquaresOfDigits(2, 3))
	fmt.Println(maximizeSumOfSquaresOfDigits(2, 17))
	fmt.Println(maximizeSumOfSquaresOfDigits(1, 10))
}
```

## 3724 — Minimum Operations To Transform Array

```go
package main

// LeetCode #3724: Minimum Operations to Transform Array
// https://leetcode.com/problems/minimum-operations-to-transform-array/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func minimumOperationsToTransformArray(nums1 []int, nums2 []int) int64 {
	n := len(nums1)
	last := nums2[n]
	var ops int64 = 1
	var extra int64 = 1 << 60

	for i := 0; i < n; i++ {
		lo, hi := nums1[i], nums2[i]
		if lo > hi {
			lo, hi = hi, lo
		}
		ops += int64(hi - lo)

		if lo <= last && last <= hi {
			extra = 0
		} else if last < lo {
			if int64(1+lo-last) < extra {
				extra = int64(1 + lo - last)
			}
		} else {
			if int64(1+last-hi) < extra {
				extra = int64(1 + last - hi)
			}
		}
	}

	return ops + extra
}

func main() {
	fmt.Println(minimumOperationsToTransformArray([]int{2, 8}, []int{1, 7, 3}))
	fmt.Println(minimumOperationsToTransformArray([]int{1, 2}, []int{3, 4, 5}))
	fmt.Println(minimumOperationsToTransformArray([]int{5, 5}, []int{5, 5, 5}))
}
```

## 3727 — Maximum Alternating Sum Of Squares

```go
package main

// LeetCode #3727: Maximum Alternating Sum of Squares
// https://leetcode.com/problems/maximum-alternating-sum-of-squares/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func maximumAlternatingSumOfSquares(nums []int) int64 {
	sort.Slice(nums, func(i, j int) bool {
		return abs(nums[i]) > abs(nums[j])
	})

	var ans int64
	for i, v := range nums {
		sq := int64(v) * int64(v)
		if i%2 == 0 {
			ans += sq
		} else {
			ans -= sq
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(maximumAlternatingSumOfSquares([]int{1, 2, 3}))
	fmt.Println(maximumAlternatingSumOfSquares([]int{1, -1, 2, -2, 3, -3}))
	fmt.Println(maximumAlternatingSumOfSquares([]int{0, 0, 0}))
}
```

## 3728 — Stable Subarrays With Equal Boundary And Interior Sum

```go
package main

// LeetCode #3728: Stable Subarrays With Equal Boundary and Interior Sum
// https://leetcode.com/problems/stable-subarrays-with-equal-boundary-and-interior-sum/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func stableSubarraysWithEqualBoundaryAndInteriorSum(capacity []int) int64 {
	n := len(capacity)
	s := make([]int64, n+1)
	for i := 0; i < n; i++ {
		s[i+1] = s[i] + int64(capacity[i])
	}

	type pair struct {
		val int
		sum int64
	}
	cnt := make(map[pair]int)
	var ans int64

	for r := 0; r < n; r++ {
		// Query for valid left boundaries
		key := pair{val: capacity[r], sum: s[r] - int64(capacity[r])}
		ans += int64(cnt[key])

		// Insert position r-1 as future left boundary (delayed by 1 for length >= 3)
		if r >= 1 {
			ins := pair{val: capacity[r-1], sum: s[r]}
			cnt[ins]++
		}
	}

	return ans
}

func main() {
	fmt.Println(stableSubarraysWithEqualBoundaryAndInteriorSum([]int{9, 3, 3, 3, 9}))
	fmt.Println(stableSubarraysWithEqualBoundaryAndInteriorSum([]int{1, 2, 3, 4, 5}))
	fmt.Println(stableSubarraysWithEqualBoundaryAndInteriorSum([]int{5, 2, 1, 2, 5}))
}
```

## 3730 — Maximum Calories Burnt From Jumps

```go
package main

// LeetCode #3730: Maximum Calories Burnt from Jumps
// https://leetcode.com/problems/maximum-calories-burnt-from-jumps/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func maximumCaloriesBurntFromJumps(heights []int) int64 {
	n := len(heights)
	sort.Ints(heights)

	seq := make([]int, 0, n)
	l, r := 0, n-1
	for l <= r {
		seq = append(seq, heights[r])
		r--
		if l <= r {
			seq = append(seq, heights[l])
			l++
		}
	}

	total := int64(seq[0]) * int64(seq[0])
	for i := 1; i < n; i++ {
		diff := seq[i] - seq[i-1]
		total += int64(diff) * int64(diff)
	}
	return total
}

func main() {
	fmt.Println(maximumCaloriesBurntFromJumps([]int{1, 7, 9}))
	fmt.Println(maximumCaloriesBurntFromJumps([]int{5, 2, 4}))
	fmt.Println(maximumCaloriesBurntFromJumps([]int{3, 3}))
}
```

## 3732 — Maximum Product Of Three Elements After One Replacement

```go
package main

// LeetCode #3732: Maximum Product of Three Elements After One Replacement
// https://leetcode.com/problems/maximum-product-of-three-elements-after-one-replacement/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func maximumProductOfThreeElementsAfterOneReplacement(nums []int) int64 {
	var first, second int64 = 0, 0
	for _, v := range nums {
		val := int64(v)
		if val < 0 {
			val = -val
		}
		if val > first {
			second = first
			first = val
		} else if val > second {
			second = val
		}
	}
	return 100000 * first * second
}

func main() {
	fmt.Println(maximumProductOfThreeElementsAfterOneReplacement([]int{1, 2, 3, 4}))
	fmt.Println(maximumProductOfThreeElementsAfterOneReplacement([]int{-4, -2, -1, -3}))
	fmt.Println(maximumProductOfThreeElementsAfterOneReplacement([]int{-5, 7, 0}))
}
```

## 3733 — Minimum Time To Complete All Deliveries

```go
package main

// LeetCode #3733: Minimum Time to Complete All Deliveries
// https://leetcode.com/problems/minimum-time-to-complete-all-deliveries/
// Difficulty: Medium
// Time: O(log(maxTime)) | Space: O(1)

import "fmt"

func minimumTimeToCompleteAllDeliveries(d []int, r []int) int64 {
	a, b := int64(d[0]), int64(d[1])
	x, y := int64(r[0]), int64(r[1])

	// Check if t hours is enough
	check := func(t int64) bool {
		// Hours drone 2 cannot work (drone 1 only)
		only1 := t / y - t/lcm(x, y)
		// Hours drone 1 cannot work (drone 2 only)
		only2 := t / x - t/lcm(x, y)
		// Hours both can work
		both := t - t/x - t/y + t/lcm(x, y)

		needA := a - only1
		if needA < 0 {
			needA = 0
		}
		needB := b - only2
		if needB < 0 {
			needB = 0
		}
		return needA+needB <= both
	}

	lo, hi := a+b, (a+b)*max(x, y)
	for lo < hi {
		mid := lo + (hi-lo)/2
		if check(mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int64) int64 {
	return a / gcd(a, b) * b
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minimumTimeToCompleteAllDeliveries([]int{3, 1}, []int{2, 3}))
	fmt.Println(minimumTimeToCompleteAllDeliveries([]int{1, 3}, []int{2, 2}))
	fmt.Println(minimumTimeToCompleteAllDeliveries([]int{2, 1}, []int{3, 4}))
}
```

## 3737 — Count Subarrays With Majority Element I

```go
package main

// LeetCode #3737: Count Subarrays With Majority Element I
// https://leetcode.com/problems/count-subarrays-with-majority-element-i/
// Difficulty: Medium
// Time: O(n^2) | Space: O(1)

import "fmt"

func countSubarraysWithMajorityElementI(nums []int, target int) int {
	n := len(nums)
	ans := 0
	for i := 0; i < n; i++ {
		cnt := 0
		for j := i; j < n; j++ {
			if nums[j] == target {
				cnt++
			}
			if cnt*2 > j-i+1 {
				ans++
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(countSubarraysWithMajorityElementI([]int{1, 2, 1, 2, 1}, 1))
	fmt.Println(countSubarraysWithMajorityElementI([]int{3, 1, 2, 3}, 3))
	fmt.Println(countSubarraysWithMajorityElementI([]int{1, 2, 3, 4}, 1))
}
```

## 3738 — Longest Non Decreasing Subarray After Replacing At Most One Element

```go
package main

// LeetCode #3738: Longest Non-Decreasing Subarray After Replacing at Most One Element
// https://leetcode.com/problems/longest-non-decreasing-subarray-after-replacing-at-most-one-element/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func longestNonDecreasingSubarrayAfterReplacingAtMostOneElement(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 1
	}

	left := make([]int, n)
	right := make([]int, n)
	left[0] = 1
	right[n-1] = 1

	for i := 1; i < n; i++ {
		if nums[i] >= nums[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}
	for i := n - 2; i >= 0; i-- {
		if nums[i] <= nums[i+1] {
			right[i] = right[i+1] + 1
		} else {
			right[i] = 1
		}
	}

	ans := 1
	for i := 0; i < n; i++ {
		if i == 0 {
			if 1+right[i+1] > ans {
				ans = 1 + right[i+1]
			}
		} else if i == n-1 {
			if 1+left[i-1] > ans {
				ans = 1 + left[i-1]
			}
		} else if nums[i-1] <= nums[i+1] {
			cand := left[i-1] + 1 + right[i+1]
			if cand > ans {
				ans = cand
			}
		} else {
			cand := left[i-1]
			if right[i+1] > cand {
				cand = right[i+1]
			}
			cand++
			if cand > ans {
				ans = cand
			}
		}
	}

	return ans
}

func main() {
	fmt.Println(longestNonDecreasingSubarrayAfterReplacingAtMostOneElement([]int{1, 2, 3, 1, 2}))
	fmt.Println(longestNonDecreasingSubarrayAfterReplacingAtMostOneElement([]int{5, 4, 3, 2, 1}))
	fmt.Println(longestNonDecreasingSubarrayAfterReplacingAtMostOneElement([]int{1, 2, 3, 4}))
}
```

## 3741 — Minimum Distance Between Three Equal Elements Ii

```go
package main

// LeetCode #3741: Minimum Distance Between Three Equal Elements II
// https://leetcode.com/problems/minimum-distance-between-three-equal-elements-ii/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func minimumDistanceBetweenThreeEqualElementsIi(nums []int) int {
	ans := -1
	prev1 := make(map[int]int)
	prev2 := make(map[int]int)

	// prev1[v] = last index where v appeared
	// prev2[v] = second-last index where v appeared

	for i, v := range nums {
		if p2, ok := prev2[v]; ok {
			dist := 2 * (i - p2)
			if ans == -1 || dist < ans {
				ans = dist
			}
		}
		// Shift: prev2 gets prev1, prev1 gets current
		prev2[v] = prev1[v]
		prev1[v] = i
	}

	return ans
}

func main() {
	fmt.Println(minimumDistanceBetweenThreeEqualElementsIi([]int{1, 3, 1, 1, 2, 1}))
	fmt.Println(minimumDistanceBetweenThreeEqualElementsIi([]int{1, 2, 3, 4}))
	fmt.Println(minimumDistanceBetweenThreeEqualElementsIi([]int{1, 1, 1}))
}
```

## 3742 — Maximum Path Score In A Grid

```go
package main

// LeetCode #3742: Maximum Path Score in a Grid
// https://leetcode.com/problems/maximum-path-score-in-a-grid/
// Difficulty: Medium
// Time: O(m*n*k) | Space: O(m*n*k)

import "fmt"

func maximumPathScoreInAGrid(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	maxCost := k + 1
	if m+n+1 < maxCost {
		maxCost = m + n + 1
	}

	// dp[i][j][c] = max score at (i,j) with cost c, -1 = unreachable
	dp := make([][][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, maxCost)
			for c := 0; c < maxCost; c++ {
				dp[i][j][c] = -1
			}
		}
	}

	dp[0][0][0] = 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for c := 0; c < maxCost; c++ {
				cur := dp[i][j][c]
				if cur == -1 {
					continue
				}
				// Move right
				if j+1 < n {
					add := 0
					if grid[i][j+1] > 0 {
						add = 1
					}
					if c+add < maxCost {
						val := cur + grid[i][j+1]
						if val > dp[i][j+1][c+add] {
							dp[i][j+1][c+add] = val
						}
					}
				}
				// Move down
				if i+1 < m {
					add := 0
					if grid[i+1][j] > 0 {
						add = 1
					}
					if c+add < maxCost {
						val := cur + grid[i+1][j]
						if val > dp[i+1][j][c+add] {
							dp[i+1][j][c+add] = val
						}
					}
				}
			}
		}
	}

	ans := -1
	for c := 0; c < maxCost; c++ {
		if dp[m-1][n-1][c] > ans {
			ans = dp[m-1][n-1][c]
		}
	}
	return ans
}

func main() {
	fmt.Println(maximumPathScoreInAGrid([][]int{{0, 1, 2}, {1, 0, 1}, {2, 1, 0}}, 3))
	fmt.Println(maximumPathScoreInAGrid([][]int{{0, 0}, {0, 0}}, 1))
	fmt.Println(maximumPathScoreInAGrid([][]int{{0, 2}, {2, 0}}, 1))
}
```

## 3744 — Find Kth Character In Expanded String

```go
package main

// LeetCode #3744: Find Kth Character in Expanded String
// https://leetcode.com/problems/find-kth-character-in-expanded-string/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"strings"
)

func findKthCharacterInExpandedString(s string, k int) byte {
	words := strings.Fields(s)

	for _, word := range words {
		l := len(word)
		m := l * (l + 1) / 2

		if k == m {
			return ' '
		} else if k > m {
			k -= (m + 1)
			continue
		} else {
			cur := 0
			for i, ch := range word {
				cur += (i + 1)
				if k < cur {
					return byte(ch)
				}
			}
			return ' '
		}
	}
	return ' '
}

func main() {
	fmt.Printf("%c\n", findKthCharacterInExpandedString("hello world", 0))
	fmt.Printf("%c\n", findKthCharacterInExpandedString("hello world", 15))
	fmt.Printf("%c\n", findKthCharacterInExpandedString("hello world", 20))
}
```

## 3746 — Minimum String Length After Balanced Removals

```go
package main

// LeetCode #3746: Minimum String Length After Balanced Removals
// https://leetcode.com/problems/minimum-string-length-after-balanced-removals/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func minimumStringLengthAfterBalancedRemovals(s string) int {
	a, b := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' {
			a++
		} else {
			b++
		}
	}
	if a > b {
		return a - b
	}
	return b - a
}

func main() {
	fmt.Println(minimumStringLengthAfterBalancedRemovals("aabbab"))
	fmt.Println(minimumStringLengthAfterBalancedRemovals("aaaa"))
	fmt.Println(minimumStringLengthAfterBalancedRemovals("aaabb"))
}
```

## 3747 — Count Distinct Integers After Removing Zeros

```go
package main

// LeetCode #3747: Count Distinct Integers After Removing Zeros
// https://leetcode.com/problems/count-distinct-integers-after-removing-zeros/
// Difficulty: Medium
// Time: O(log n) | Space: O(log n)

import "fmt"

func countDistinctIntegersAfterRemovingZeros(n int64) int64 {
	s := fmt.Sprintf("%d", n)
	m := len(s)

	// Precompute powers of 9
	pow9 := make([]int64, m+1)
	pow9[0] = 1
	for i := 1; i <= m; i++ {
		pow9[i] = pow9[i-1] * 9
	}

	// Count numbers with fewer digits (all non-zero digits)
	var ans int64
	for length := 1; length < m; length++ {
		ans += pow9[length]
	}

	// Count numbers with same length as n, but <= n
	for idx := 0; idx < m; idx++ {
		d := int(s[idx] - '0')
		if d == 0 {
			return ans
		}
		ans += int64(d-1) * pow9[m-idx-1]
	}
	return ans + 1
}

func main() {
	fmt.Println(countDistinctIntegersAfterRemovingZeros(10))
	fmt.Println(countDistinctIntegersAfterRemovingZeros(100))
	fmt.Println(countDistinctIntegersAfterRemovingZeros(1))
}
```

## 3751 — Total Waviness Of Numbers In Range I

```go
package main

// LeetCode #3751: Total Waviness of Numbers in Range I
// https://leetcode.com/problems/total-waviness-of-numbers-in-range-i/
// Difficulty: Medium
// Time: O(N * D) | Space: O(1)

import (
	"fmt"
	"strconv"
)

func totalWavinessOfNumbersInRangeI(num1 int, num2 int) int {
	ans := 0
	for num := num1; num <= num2; num++ {
		s := strconv.Itoa(num)
		for i := 1; i < len(s)-1; i++ {
			if (s[i] > s[i-1] && s[i] > s[i+1]) || (s[i] < s[i-1] && s[i] < s[i+1]) {
				ans++
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(totalWavinessOfNumbersInRangeI(1, 100))
	fmt.Println(totalWavinessOfNumbersInRangeI(100, 200))
	fmt.Println(totalWavinessOfNumbersInRangeI(1000, 1050))
}
```

## 3752 — Lexicographically Smallest Negated Permutation That Sums To Target

```go
package main

// LeetCode #3752: Lexicographically Smallest Negated Permutation that Sums to Target
// https://leetcode.com/problems/lexicographically-smallest-negated-permutation-that-sums-to-target/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func lexicographicallySmallestNegatedPermutationThatSumsToTarget(n int, target int64) []int {
	s := int64(n) * int64(n+1) / 2
	drop := s - target
	if drop < 0 || drop%2 != 0 {
		return []int{}
	}

	delta := drop / 2
	used := make([]bool, n+1)
	negatedSum := int64(0)

	// Greedily pick largest numbers to negate
	for i := n; i > 0; i-- {
		if negatedSum+int64(i) <= delta {
			used[i] = true
			negatedSum += int64(i)
		}
	}

	if negatedSum != delta {
		return []int{}
	}

	ans := make([]int, 0, n)
	// Negated numbers first (from largest to smallest = lexicographically smallest)
	for i := n; i > 0; i-- {
		if used[i] {
			ans = append(ans, -i)
		}
	}
	// Then positive numbers in ascending order
	for i := 1; i <= n; i++ {
		if !used[i] {
			ans = append(ans, i)
		}
	}
	return ans
}

func main() {
	fmt.Println(lexicographicallySmallestNegatedPermutationThatSumsToTarget(3, 0))
	fmt.Println(lexicographicallySmallestNegatedPermutationThatSumsToTarget(4, 10))
	fmt.Println(lexicographicallySmallestNegatedPermutationThatSumsToTarget(5, 15))
}
```

## 3755 — Find Maximum Balanced Xor Subarray Length

```go
package main

// LeetCode #3755: Find Maximum Balanced XOR Subarray Length
// https://leetcode.com/problems/find-maximum-balanced-xor-subarray-length/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func findMaximumBalancedXorSubarrayLength(nums []int) int {
	type state struct {
		xor  int
		diff int
	}
	first := make(map[state]int)
	// Initial state before first element
	first[state{xor: 0, diff: 0}] = -1

	prefixXor := 0
	diff := 0
	maxLen := 0

	for i, v := range nums {
		prefixXor ^= v
		if v%2 == 1 {
			diff++
		} else {
			diff--
		}

		key := state{xor: prefixXor, diff: diff}
		if pos, ok := first[key]; ok {
			if i-pos > maxLen {
				maxLen = i - pos
			}
		} else {
			first[key] = i
		}
	}

	return maxLen
}

func main() {
	fmt.Println(findMaximumBalancedXorSubarrayLength([]int{3, 1, 3, 2, 0}))
	fmt.Println(findMaximumBalancedXorSubarrayLength([]int{3, 2, 8, 5, 4, 14, 9, 15}))
	fmt.Println(findMaximumBalancedXorSubarrayLength([]int{1, 2, 3, 4, 5}))
}
```

## 3756 — Concatenate Non Zero Digits And Multiply By Sum Ii

```go
package main

// LeetCode #3756: Concatenate Non-Zero Digits and Multiply by Sum II
// https://leetcode.com/problems/concatenate-non-zero-digits-and-multiply-by-sum-ii/
// Difficulty: Medium
// Time: O(n + q) | Space: O(n)

import "fmt"

const mod3756 = 1000000007

func concatenateNonZeroDigitsAndMultiplyBySumIi(s string, queries [][]int) []int {
	n := len(s)
	prefixCnt := make([]int, n+1)
	prefixSum := make([]int, n+1)
	prefixNum := make([]int64, n+1)

	for i, ch := range s {
		d := int(ch - '0')
		prefixSum[i+1] = prefixSum[i] + d
		if d > 0 {
			prefixCnt[i+1] = prefixCnt[i] + 1
			prefixNum[i+1] = (prefixNum[i]*10 + int64(d)) % mod3756
		} else {
			prefixCnt[i+1] = prefixCnt[i]
			prefixNum[i+1] = prefixNum[i]
		}
	}

	ans := make([]int, len(queries))
	for qi, q := range queries {
		l, r := q[0], q[1]
		subLen := prefixCnt[r+1] - prefixCnt[l]
		// x = (prefixNum[r+1] - prefixNum[l] * 10^subLen) % mod
		x := (prefixNum[r+1] - prefixNum[l]*pow10(int64(subLen))) % mod3756
		if x < 0 {
			x += mod3756
		}
		digitSum := prefixSum[r+1] - prefixSum[l]
		ans[qi] = int((x * int64(digitSum)) % mod3756)
	}
	return ans
}

func pow10(exp int64) int64 {
	if exp == 0 {
		return 1
	}
	if exp == 1 {
		return 10
	}
	half := pow10(exp / 2)
	half = (half * half) % mod3756
	if exp%2 == 1 {
		half = (half * 10) % mod3756
	}
	return half
}

func main() {
	fmt.Println(concatenateNonZeroDigitsAndMultiplyBySumIi("10203004", [][]int{{0, 7}, {1, 3}, {4, 6}}))
	fmt.Println(concatenateNonZeroDigitsAndMultiplyBySumIi("1000", [][]int{{0, 3}, {1, 1}}))
	fmt.Println(concatenateNonZeroDigitsAndMultiplyBySumIi("9876543210", [][]int{{0, 9}}))
}
```

## 3758 — Convert Number Words To Digits

```go
package main

// LeetCode #3758: Convert Number Words to Digits
// https://leetcode.com/problems/convert-number-words-to-digits/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import (
	"fmt"
	"strings"
)

func convertNumberWordsToDigits(s string) string {
	words := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	var res strings.Builder
	i := 0
	for i < len(s) {
		found := false
		for d, word := range words {
			if i+len(word) <= len(s) && s[i:i+len(word)] == word {
				res.WriteByte(byte(d) + '0')
				i += len(word)
				found = true
				break
			}
		}
		if !found {
			i++
		}
	}
	return res.String()
}

func main() {
	fmt.Println(convertNumberWordsToDigits("onefourthree"))
	fmt.Println(convertNumberWordsToDigits("ninexsix"))
	fmt.Println(convertNumberWordsToDigits("zeero"))
}
```

## 3759 — Count Elements With At Least K Greater Values

```go
package main

// LeetCode #3759: Count Elements With at Least K Greater Values
// https://leetcode.com/problems/count-elements-with-at-least-k-greater-values/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func countElementsWithAtLeastKGreaterValues(nums []int, k int) int {
	if k == 0 {
		return len(nums)
	}
	n := len(nums)
	sort.Ints(nums)
	left := n - k
	// Skip duplicates of threshold value
	for left-1 >= 0 && nums[left-1] == nums[left] {
		left--
	}
	return left
}

func main() {
	fmt.Println(countElementsWithAtLeastKGreaterValues([]int{3, 1, 2}, 1))
	fmt.Println(countElementsWithAtLeastKGreaterValues([]int{5, 5, 5}, 2))
	fmt.Println(countElementsWithAtLeastKGreaterValues([]int{1, 2, 3, 4, 5}, 2))
}
```

## 3760 — Maximum Substrings With Distinct Start

```go
package main

// LeetCode #3760: Maximum Substrings With Distinct Start
// https://leetcode.com/problems/maximum-substrings-with-distinct-start/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func maximumSubstringsWithDistinctStart(s string) int {
	seen := [26]bool{}
	ans := 0
	for _, ch := range s {
		idx := ch - 'a'
		if !seen[idx] {
			seen[idx] = true
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(maximumSubstringsWithDistinctStart("abacaba"))
	fmt.Println(maximumSubstringsWithDistinctStart("aaaa"))
	fmt.Println(maximumSubstringsWithDistinctStart("abc"))
}
```

## 3761 — Minimum Absolute Distance Between Mirror Pairs

```go
package main

// LeetCode #3761: Minimum Absolute Distance Between Mirror Pairs
// https://leetcode.com/problems/minimum-absolute-distance-between-mirror-pairs/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func minimumAbsoluteDistanceBetweenMirrorPairs(nums []int) int {
	prev := make(map[int]int)
	ans := -1

	for j, v := range nums {
		if pos, ok := prev[v]; ok {
			dist := j - pos
			if ans == -1 || dist < ans {
				ans = dist
			}
		}
		// Store reversed number
		rev := 0
		for x := v; x > 0; x /= 10 {
			rev = rev*10 + x%10
		}
		prev[rev] = j
	}

	return ans
}

func main() {
	fmt.Println(minimumAbsoluteDistanceBetweenMirrorPairs([]int{12, 21, 45, 33, 54}))
	fmt.Println(minimumAbsoluteDistanceBetweenMirrorPairs([]int{1, 2, 3, 4}))
	fmt.Println(minimumAbsoluteDistanceBetweenMirrorPairs([]int{11, 22, 11}))
}
```

## 3763 — Maximum Total Sum With Threshold Constraints

```go
package main

// LeetCode #3763: Maximum Total Sum with Threshold Constraints
// https://leetcode.com/problems/maximum-total-sum-with-threshold-constraints/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"container/heap"
	"fmt"
	"sort"
)

type maxHeap []int

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *maxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *maxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func maximumTotalSumWithThresholdConstraints(nums []int, threshold []int) int {
	n := len(nums)
	items := make([][2]int, n)
	for i := 0; i < n; i++ {
		items[i] = [2]int{threshold[i], nums[i]}
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i][0] < items[j][0]
	})

	h := &maxHeap{}
	heap.Init(h)
	idx := 0
	total := 0

	for step := 1; step <= n; step++ {
		for idx < n && items[idx][0] <= step {
			heap.Push(h, items[idx][1])
			idx++
		}
		if h.Len() == 0 {
			break
		}
		total += heap.Pop(h).(int)
	}
	return total
}

func main() {
	fmt.Println(maximumTotalSumWithThresholdConstraints([]int{5, 3, 4}, []int{3, 1, 2}))
	fmt.Println(maximumTotalSumWithThresholdConstraints([]int{10, 20, 30}, []int{3, 2, 1}))
	fmt.Println(maximumTotalSumWithThresholdConstraints([]int{1, 2}, []int{1, 1}))
}
```

## 3765 — Complete Prime Number

```go
package main

// LeetCode #3765: Complete Prime Number
// https://leetcode.com/problems/complete-prime-number/
// Difficulty: Medium
// Time: O(d * sqrt(n)) | Space: O(1)

import "fmt"

func completePrimeNumber(num int) bool {
	if num <= 1 {
		return false
	}

	suffix := 0
	power := 1
	x := num

	for x > 0 {
		// Build suffix: prepend last digit
		suffix = power*(x%10) + suffix
		power *= 10

		if !isPrime(suffix) {
			return false
		}
		if !isPrime(x) {
			return false
		}
		x /= 10
	}
	return true
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(completePrimeNumber(23))
	fmt.Println(completePrimeNumber(39))
	fmt.Println(completePrimeNumber(7))
}
```

## 3766 — Minimum Operations To Make Binary Palindrome

```go
package main

// LeetCode #3766: Minimum Operations to Make Binary Palindrome
// https://leetcode.com/problems/minimum-operations-to-make-binary-palindrome/
// Difficulty: Medium
// Time: O(n * log M) | Space: O(M)

import (
	"fmt"
	"sort"
	"strconv"
)

var binaryPalindromes []int

func init() {
	for i := 0; i < (1 << 14); i++ {
		s := strconv.FormatInt(int64(i), 2)
		if isPalindromeStr(s) {
			binaryPalindromes = append(binaryPalindromes, i)
		}
	}
}

func isPalindromeStr(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func minimumOperationsToMakeBinaryPalindrome(nums []int) []int {
	ans := make([]int, len(nums))
	for idx, x := range nums {
		pos := sort.SearchInts(binaryPalindromes, x)
		best := 1 << 30
		if pos < len(binaryPalindromes) {
			if binaryPalindromes[pos]-x < best {
				best = binaryPalindromes[pos] - x
			}
		}
		if pos > 0 {
			if x-binaryPalindromes[pos-1] < best {
				best = x - binaryPalindromes[pos-1]
			}
		}
		ans[idx] = best
	}
	return ans
}

func main() {
	fmt.Println(minimumOperationsToMakeBinaryPalindrome([]int{1, 2, 3, 4, 5}))
	fmt.Println(minimumOperationsToMakeBinaryPalindrome([]int{10, 20}))
	fmt.Println(minimumOperationsToMakeBinaryPalindrome([]int{7}))
}
```

## 3767 — Maximize Points After Choosing K Tasks

```go
package main

// LeetCode #3767: Maximize Points After Choosing K Tasks
// https://leetcode.com/problems/maximize-points-after-choosing-k-tasks/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func maximizePointsAfterChoosingKTasks(technique1 []int, technique2 []int, k int) int64 {
	n := len(technique1)
	type task struct {
		diff   int
		t1, t2 int
	}
	tasks := make([]task, n)
	for i := 0; i < n; i++ {
		tasks[i] = task{
			diff: technique1[i] - technique2[i],
			t1:   technique1[i],
			t2:   technique2[i],
		}
	}

	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].diff > tasks[j].diff
	})

	var ans int64
	posCount := 0
	for i := 0; i < n; i++ {
		if i < k {
			ans += int64(tasks[i].t1)
			posCount++
		} else if tasks[i].diff >= 0 {
			ans += int64(tasks[i].t1)
			posCount++
		} else {
			ans += int64(tasks[i].t2)
		}
	}
	return ans
}

func main() {
	fmt.Println(maximizePointsAfterChoosingKTasks([]int{5, 3, 4}, []int{2, 6, 1}, 2))
	fmt.Println(maximizePointsAfterChoosingKTasks([]int{1, 2, 3}, []int{4, 5, 6}, 1))
	fmt.Println(maximizePointsAfterChoosingKTasks([]int{10, 20}, []int{5, 15}, 1))
}
```

## 3770 — Largest Prime From Consecutive Prime Sum

```go
package main

// LeetCode #3770: Largest Prime from Consecutive Prime Sum
// https://leetcode.com/problems/largest-prime-from-consecutive-prime-sum/
// Difficulty: Medium
// Time: O(n log log n) | Space: O(n)

import "fmt"

func largestPrimeFromConsecutivePrimeSum(n int) int {
	if n < 2 {
		return 0
	}
	// Sieve up to n
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i <= n; i++ {
		if isPrime[i] {
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}

	// Collect primes
	primes := make([]int, 0)
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}

	// Prefix sums
	pref := make([]int, len(primes)+1)
	for i, p := range primes {
		pref[i+1] = pref[i] + p
	}

	ans := 0
	for i := 0; i < len(primes); i++ {
		for j := i; j < len(primes); j++ {
			sum := pref[j+1] - pref[i]
			if sum > n {
				break
			}
			if isPrime[sum] && sum > ans {
				ans = sum
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(largestPrimeFromConsecutivePrimeSum(20))
	fmt.Println(largestPrimeFromConsecutivePrimeSum(2))
	fmt.Println(largestPrimeFromConsecutivePrimeSum(50))
}
```

## 3771 — Total Score Of Dungeon Runs

```go
package main

// LeetCode #3771: Total Score of Dungeon Runs
// https://leetcode.com/problems/total-score-of-dungeon-runs/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func totalScoreOfDungeonRuns(hp int, damage []int, requirement []int) int64 {
	n := len(damage)
	// suffix cumulative damage
	suf := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		suf[i] = suf[i+1] + damage[i]
	}

	// For each starting index i, binary search how many rooms we can pass
	// Condition: hp - (damage[i] + ... + damage[j]) >= requirement[j]
	// => hp - requirement[j] >= suf[i] - suf[j+1]
	// => suf[j+1] >= suf[i] - (hp - requirement[j])
	// => if suf[i] - suf[j+1] <= hp - requirement[j]

	// Need to find j >= i where we have enough hp after damage to meet requirement
	// Equivalent: suf[i] - suf[j+1] <= hp - requirement[j]
	// => suf[j+1] >= suf[i] - (hp - requirement[j])

	// Binary search approach: for each i, find valid j range
	// Since suffix sums are decreasing as j increases, we can track valid range

	var ans int64
	for i := 0; i < n; i++ {
		// Binary search for furthest j we can reach
		lo, hi := i, n-1
		best := -1
		for lo <= hi {
			mid := lo + (hi-lo)/2
			damageSum := suf[i] - suf[mid+1]
			if hp-damageSum >= requirement[mid] {
				best = mid
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
		if best != -1 {
			// We can pass rooms i through best
			// This contributes to scores at start positions <= i
			// Specifically, for start position s where s <= i <= best:
			// score(s) includes this room
			// Number of start positions that include room i = i - 0 + 1 = i+1
			// But only if best >= i (always true here)

			// Actually simpler: for each start i, we pass (best - i + 1) rooms
			ans += int64(best - i + 1)
		}
	}
	return ans
}

func binarySearch(arr []int, target int) int {
	return sort.SearchInts(arr, target)
}

func main() {
	fmt.Println(totalScoreOfDungeonRuns(11, []int{3, 6, 7}, []int{4, 2, 5}))
	fmt.Println(totalScoreOfDungeonRuns(5, []int{1, 2, 3}, []int{1, 1, 1}))
	fmt.Println(totalScoreOfDungeonRuns(1, []int{5}, []int{10}))
}
```

## 3773 — Maximum Number Of Equal Length Runs

```go
package main

// LeetCode #3773: Maximum Number of Equal Length Runs
// https://leetcode.com/problems/maximum-number-of-equal-length-runs/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func maximumNumberOfEqualLengthRuns(s string) int {
	cnt := make(map[int]int)
	maxCount := 0
	n := len(s)
	i := 0
	for i < n {
		j := i
		for j < n && s[j] == s[i] {
			j++
		}
		runLen := j - i
		cnt[runLen]++
		if cnt[runLen] > maxCount {
			maxCount = cnt[runLen]
		}
		i = j
	}
	return maxCount
}

func main() {
	fmt.Println(maximumNumberOfEqualLengthRuns("hello"))
	fmt.Println(maximumNumberOfEqualLengthRuns("aaabaaa"))
	fmt.Println(maximumNumberOfEqualLengthRuns("aabbcc"))
}
```

## 3775 — Reverse Words With Same Vowel Count

```go
package main

// LeetCode #3775: Reverse Words With Same Vowel Count
// https://leetcode.com/problems/reverse-words-with-same-vowel-count/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"strings"
)

func reverseWordsWithSameVowelCount(s string) string {
	words := strings.Fields(s)
	if len(words) == 0 {
		return s
	}

	vowels := "aeiou"
	countVowels := func(w string) int {
		c := 0
		for _, ch := range w {
			if strings.ContainsRune(vowels, ch) {
				c++
			}
		}
		return c
	}

	firstCnt := countVowels(words[0])
	for i := 1; i < len(words); i++ {
		if countVowels(words[i]) == firstCnt {
			// Reverse word
			runes := []rune(words[i])
			for l, r := 0, len(runes)-1; l < r; l, r = l+1, r-1 {
				runes[l], runes[r] = runes[r], runes[l]
			}
			words[i] = string(runes)
		}
	}

	return strings.Join(words, " ")
}

func main() {
	fmt.Println(reverseWordsWithSameVowelCount("cat and mice"))
	fmt.Println(reverseWordsWithSameVowelCount("book is nice"))
	fmt.Println(reverseWordsWithSameVowelCount("banana healthy"))
}
```

## 3776 — Minimum Moves To Balance Circular Array

```go
package main

// LeetCode #3776: Minimum Moves to Balance Circular Array
// https://leetcode.com/problems/minimum-moves-to-balance-circular-array/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func minimumMovesToBalanceCircularArray(balance []int) int64 {
	sum := 0
	n := len(balance)
	negIdx := -1
	for i, v := range balance {
		sum += v
		if v < 0 {
			negIdx = i
		}
	}
	if sum < 0 {
		return -1
	}
	if negIdx == -1 {
		return 0
	}

	need := -balance[negIdx]
	var ans int64
	for d := 1; d < n && need > 0; d++ {
		left := balance[(negIdx-d+n)%n]
		right := balance[(negIdx+d)%n]

		if left > 0 {
			take := left
			if take > need {
				take = need
			}
			need -= take
			ans += int64(take) * int64(d)
		}
		if need > 0 && right > 0 {
			take := right
			if take > need {
				take = need
			}
			need -= take
			ans += int64(take) * int64(d)
		}
	}
	return ans
}

func main() {
	fmt.Println(minimumMovesToBalanceCircularArray([]int{-2, 1, 1}))
	fmt.Println(minimumMovesToBalanceCircularArray([]int{1, -1, 0}))
	fmt.Println(minimumMovesToBalanceCircularArray([]int{0, 0, 0}))
}
```

## 3778 — Minimum Distance Excluding One Maximum Weighted Edge

```go
package main

// LeetCode #3778: Minimum Distance Excluding One Maximum Weighted Edge
// https://leetcode.com/problems/minimum-distance-excluding-one-maximum-weighted-edge/
// Difficulty: Medium [Paid]
// Time: O(E log V) | Space: O(V + E)

import (
	"container/heap"
	"fmt"
	"math"
)

type edge struct {
	to, w int
}

type state struct {
	node     int
	used     bool // whether max edge already excluded
	total    int
	maxEdge  int
}

type pqItem struct {
	total   int
	node    int
	used    bool
	maxEdge int
	index   int
}

type priorityQueue []*pqItem

func (pq priorityQueue) Len() int { return len(pq) }
func (pq priorityQueue) Less(i, j int) bool { return pq[i].total-pq[i].maxEdge < pq[j].total-pq[j].maxEdge }
func (pq priorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i]; pq[i].index = i; pq[j].index = j }
func (pq *priorityQueue) Push(x any) {
	item := x.(*pqItem)
	item.index = len(*pq)
	*pq = append(*pq, item)
}
func (pq *priorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[:n-1]
	return item
}

func minimumDistanceExcludingOneMaximumWeightedEdge(n int, edges [][]int) int {
	graph := make([][]edge, n)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		graph[u] = append(graph[u], edge{v, w})
		graph[v] = append(graph[v], edge{u, w})
	}

	// dist[node][used] = min effective cost (sum - maxEdge)
	dist := make([][2]int, n)
	for i := 0; i < n; i++ {
		dist[i][0] = math.MaxInt32
		dist[i][1] = math.MaxInt32
	}
	dist[0][0] = 0

	pq := &priorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &pqItem{node: 0, total: 0, maxEdge: 0, used: false})

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(*pqItem)
		effective := cur.total - cur.maxEdge
		if effective != dist[cur.node][btoi(cur.used)] {
			continue
		}
		if cur.node == n-1 {
			return effective
		}

		for _, e := range graph[cur.node] {
			// Option 1: Don't exclude this edge
			newTotal := cur.total + e.w
			newMax := cur.maxEdge
			if e.w > newMax {
				newMax = e.w
			}
			eff := newTotal - newMax
			if eff < dist[e.to][btoi(cur.used)] {
				dist[e.to][btoi(cur.used)] = eff
				heap.Push(pq, &pqItem{
					node: e.to, total: newTotal, maxEdge: newMax, used: cur.used,
				})
			}

			// Option 2: Exclude this edge (if not already excluded)
			if !cur.used {
				// Excluding means we don't add e.w to total, and track it as maxEdge
				// Actually: we just skip this edge's weight entirely
				if cur.total < dist[e.to][1] {
					dist[e.to][1] = cur.total
					heap.Push(pq, &pqItem{
						node: e.to, total: cur.total, maxEdge: cur.total, used: true,
					})
				}
			}
		}
	}
	return -1
}

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

func main() {
	fmt.Println(minimumDistanceExcludingOneMaximumWeightedEdge(4, [][]int{{0, 1, 2}, {1, 2, 3}, {2, 3, 1}, {0, 3, 10}}))
	fmt.Println(minimumDistanceExcludingOneMaximumWeightedEdge(3, [][]int{{0, 1, 5}, {1, 2, 5}, {0, 2, 10}}))
	fmt.Println(minimumDistanceExcludingOneMaximumWeightedEdge(2, [][]int{{0, 1, 100}}))
}
```

## 3779 — Minimum Number Of Operations To Have Distinct Elements

```go
package main

// LeetCode #3779: Minimum Number of Operations to Have Distinct Elements
// https://leetcode.com/problems/minimum-number-of-operations-to-have-distinct-elements/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func minimumNumberOfOperationsToHaveDistinctElements(nums []int) int {
	seen := make(map[int]bool)
	for i := len(nums) - 1; i >= 0; i-- {
		if seen[nums[i]] {
			return (i + 3) / 3
		}
		seen[nums[i]] = true
	}
	return 0
}

func main() {
	fmt.Println(minimumNumberOfOperationsToHaveDistinctElements([]int{3, 8, 3, 6, 5, 8}))
	fmt.Println(minimumNumberOfOperationsToHaveDistinctElements([]int{2, 2}))
	fmt.Println(minimumNumberOfOperationsToHaveDistinctElements([]int{4, 3, 5, 1, 2}))
}
```

## 3780 — Maximum Sum Of Three Numbers Divisible By Three

```go
package main

// LeetCode #3780: Maximum Sum of Three Numbers Divisible by Three
// https://leetcode.com/problems/maximum-sum-of-three-numbers-divisible-by-three/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func maximumSumOfThreeNumbersDivisibleByThree(nums []int) int {
	groups := [3][]int{}
	for _, v := range nums {
		r := v % 3
		groups[r] = append(groups[r], v)
	}

	for r := 0; r < 3; r++ {
		sort.Slice(groups[r], func(i, j int) bool {
			return groups[r][i] > groups[r][j]
		})
	}

	ans := 0

	// (0,0,0)
	if len(groups[0]) >= 3 {
		sum := groups[0][0] + groups[0][1] + groups[0][2]
		if sum > ans {
			ans = sum
		}
	}

	// (1,1,1)
	if len(groups[1]) >= 3 {
		sum := groups[1][0] + groups[1][1] + groups[1][2]
		if sum > ans {
			ans = sum
		}
	}

	// (2,2,2)
	if len(groups[2]) >= 3 {
		sum := groups[2][0] + groups[2][1] + groups[2][2]
		if sum > ans {
			ans = sum
		}
	}

	// (0,1,2)
	if len(groups[0]) >= 1 && len(groups[1]) >= 1 && len(groups[2]) >= 1 {
		sum := groups[0][0] + groups[1][0] + groups[2][0]
		if sum > ans {
			ans = sum
		}
	}

	return ans
}

func main() {
	fmt.Println(maximumSumOfThreeNumbersDivisibleByThree([]int{4, 2, 3, 1}))
	fmt.Println(maximumSumOfThreeNumbersDivisibleByThree([]int{1, 2, 3, 4, 5}))
	fmt.Println(maximumSumOfThreeNumbersDivisibleByThree([]int{1, 1, 1}))
}
```

## 3781 — Maximum Score After Binary Swaps

```go
package main

// LeetCode #3781: Maximum Score After Binary Swaps
// https://leetcode.com/problems/maximum-score-after-binary-swaps/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"container/heap"
	"fmt"
)

type maxHeap3781 []int

func (h maxHeap3781) Len() int           { return len(h) }
func (h maxHeap3781) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap3781) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *maxHeap3781) Push(x any)        { *h = append(*h, x.(int)) }
func (h *maxHeap3781) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func maximumScoreAfterBinarySwaps(nums []int, s string) int64 {
	h := &maxHeap3781{}
	heap.Init(h)
	var ans int64

	for i, ch := range s {
		heap.Push(h, nums[i])
		if ch == '1' {
			ans += int64(heap.Pop(h).(int))
		}
	}
	return ans
}

func main() {
	fmt.Println(maximumScoreAfterBinarySwaps([]int{3, 1, 4, 2}, "1010"))
	fmt.Println(maximumScoreAfterBinarySwaps([]int{5, 2, 8, 1}, "1001"))
	fmt.Println(maximumScoreAfterBinarySwaps([]int{10, 20, 30}, "111"))
}
```

## 3784 — Minimum Deletion Cost To Make All Characters Equal

```go
package main

// LeetCode #3784: Minimum Deletion Cost to Make All Characters Equal
// https://leetcode.com/problems/minimum-deletion-cost-to-make-all-characters-equal/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func minimumDeletionCostToMakeAllCharactersEqual(s string, cost []int) int64 {
	total := int64(0)
	charCost := [26]int64{}
	for i, ch := range s {
		v := int64(cost[i])
		total += v
		idx := ch - 'a'
		charCost[idx] += v
	}

	maxCost := int64(0)
	for _, v := range charCost {
		if v > maxCost {
			maxCost = v
		}
	}
	return total - maxCost
}

func main() {
	fmt.Println(minimumDeletionCostToMakeAllCharactersEqual("aabaac", []int{1, 2, 3, 4, 1, 10}))
	fmt.Println(minimumDeletionCostToMakeAllCharactersEqual("abc", []int{10, 5, 8}))
	fmt.Println(minimumDeletionCostToMakeAllCharactersEqual("zzzzz", []int{67, 67, 67, 67, 67}))
}
```

## 3787 — Find Diameter Endpoints Of A Tree

```go
package main

// LeetCode #3787: Find Diameter Endpoints of a Tree
// https://leetcode.com/problems/find-diameter-endpoints-of-a-tree/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func findDiameterEndpointsOfATree(n int, edges [][]int) string {
	g := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}

	bfs := func(start int) (int, []int) {
		dist := make([]int, n)
		for i := 0; i < n; i++ {
			dist[i] = -1
		}
		dist[start] = 0
		q := []int{start}
		far := start
		for len(q) > 0 {
			u := q[0]
			q = q[1:]
			if dist[u] > dist[far] {
				far = u
			}
			for _, v := range g[u] {
				if dist[v] == -1 {
					dist[v] = dist[u] + 1
					q = append(q, v)
				}
			}
		}
		return far, dist
	}

	a, _ := bfs(0)
	b, distA := bfs(a)
	_, distB := bfs(b)
	diameter := distA[b]

	ans := make([]byte, n)
	for i := 0; i < n; i++ {
		if distA[i] == diameter || distB[i] == diameter {
			ans[i] = '1'
		} else {
			ans[i] = '0'
		}
	}
	return string(ans)
}

func main() {
	fmt.Println(findDiameterEndpointsOfATree(7, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {3, 5}, {1, 6}}))
	fmt.Println(findDiameterEndpointsOfATree(3, [][]int{{0, 1}, {1, 2}}))
	fmt.Println(findDiameterEndpointsOfATree(1, [][]int{}))
}
```

## 3788 — Maximum Score Of A Split

```go
package main

// LeetCode #3788: Maximum Score of a Split
// https://leetcode.com/problems/maximum-score-of-a-split/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"math"
)

func maximumScoreOfASplit(nums []int) int64 {
	n := len(nums)
	suf := make([]int64, n)
	suf[n-1] = int64(nums[n-1])
	for i := n - 2; i >= 0; i-- {
		suf[i] = int64(nums[i])
		if suf[i+1] < suf[i] {
			suf[i] = suf[i+1]
		}
	}

	var pre int64
	var ans int64 = math.MinInt64
	for i := 0; i < n-1; i++ {
		pre += int64(nums[i])
		score := pre - suf[i+1]
		if score > ans {
			ans = score
		}
	}
	return ans
}

func main() {
	fmt.Println(maximumScoreOfASplit([]int{10, -1, 3, -4, -5}))
	fmt.Println(maximumScoreOfASplit([]int{1, 2, 3, 4}))
	fmt.Println(maximumScoreOfASplit([]int{-5, -3, -1}))
}
```

## 3789 — Minimum Cost To Acquire Required Items

```go
package main

// LeetCode #3789: Minimum Cost to Acquire Required Items
// https://leetcode.com/problems/minimum-cost-to-acquire-required-items/
// Difficulty: Medium
// Time: O(1) | Space: O(1)

import "fmt"

func minimumCostToAcquireRequiredItems(cost1 int, cost2 int, costBoth int, need1 int, need2 int) int64 {
	a := int64(need1)*int64(cost1) + int64(need2)*int64(cost2)
	b := int64(costBoth) * int64(max(need1, need2))
	mn := min(need1, need2)
	c := int64(costBoth)*int64(mn) + int64(need1-mn)*int64(cost1) + int64(need2-mn)*int64(cost2)

	ans := a
	if b < ans {
		ans = b
	}
	if c < ans {
		ans = c
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minimumCostToAcquireRequiredItems(3, 2, 1, 3, 2))
	fmt.Println(minimumCostToAcquireRequiredItems(5, 4, 15, 2, 3))
	fmt.Println(minimumCostToAcquireRequiredItems(10, 10, 5, 5, 5))
}
```

## 3790 — Smallest All Ones Multiple

```go
package main

// LeetCode #3790: Smallest All-Ones Multiple
// https://leetcode.com/problems/smallest-all-ones-multiple/
// Difficulty: Medium
// Time: O(k) | Space: O(1)

import "fmt"

func smallestAllOnesMultiple(k int) int {
	if k%2 == 0 || k%5 == 0 {
		return -1
	}
	rem := 0
	for n := 1; n <= k; n++ {
		rem = (rem*10 + 1) % k
		if rem == 0 {
			return n
		}
	}
	return -1
}

func main() {
	fmt.Println(smallestAllOnesMultiple(3))
	fmt.Println(smallestAllOnesMultiple(7))
	fmt.Println(smallestAllOnesMultiple(2))
}
```

## 3792 — Sum Of Increasing Product Blocks

```go
package main

// LeetCode #3792: Sum of Increasing Product Blocks
// https://leetcode.com/problems/sum-of-increasing-product-blocks/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

const mod3792 = 1000000007

func sumOfIncreasingProductBlocks(n int) int {
	ans := 0
	k := 1
	for i := 1; i <= n; i++ {
		prod := 1
		for j := k; j < k+i; j++ {
			prod = (prod * j) % mod3792
		}
		ans = (ans + prod) % mod3792
		k += i
	}
	return ans
}

func main() {
	fmt.Println(sumOfIncreasingProductBlocks(3))
	fmt.Println(sumOfIncreasingProductBlocks(7))
	fmt.Println(sumOfIncreasingProductBlocks(1))
}
```

## 3795 — Minimum Subarray Length With Distinct Sum At Least K

```go
package main

// LeetCode #3795: Minimum Subarray Length With Distinct Sum At Least K
// https://leetcode.com/problems/minimum-subarray-length-with-distinct-sum-at-least-k/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"math"
)

func minimumSubarrayLengthWithDistinctSumAtLeastK(nums []int, k int) int {
	n := len(nums)
	freq := make(map[int]int)
	left := 0
	distinctSum := 0
	ans := math.MaxInt32

	for right := 0; right < n; right++ {
		freq[nums[right]]++
		if freq[nums[right]] == 1 {
			distinctSum += nums[right]
		}

		for distinctSum >= k {
			length := right - left + 1
			if length < ans {
				ans = length
			}
			freq[nums[left]]--
			if freq[nums[left]] == 0 {
				distinctSum -= nums[left]
			}
			left++
		}
	}

	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}

func main() {
	fmt.Println(minimumSubarrayLengthWithDistinctSumAtLeastK([]int{2, 2, 3, 1}, 4))
	fmt.Println(minimumSubarrayLengthWithDistinctSumAtLeastK([]int{3, 2, 3, 4}, 5))
	fmt.Println(minimumSubarrayLengthWithDistinctSumAtLeastK([]int{5, 5, 4}, 5))
}
```

## 3796 — Find Maximum Value In A Constrained Sequence

```go
package main

// LeetCode #3796: Find Maximum Value in a Constrained Sequence
// https://leetcode.com/problems/find-maximum-value-in-a-constrained-sequence/
// Difficulty: Medium
// Time: O(n) | Space: O(n)
// Approach: Two-pass greedy constraint propagation. Forward pass applies
// constraints from left to right, backward pass propagates from right to left.

import "fmt"

func FindMaximumValueInAConstrainedSequence(n int, restrictions [][]int, diff []int) int {
	const INF = 1 << 60
	a := make([]int, n)
	for i := range a {
		a[i] = INF
	}
	a[0] = 0

	// Apply restrictions
	for _, r := range restrictions {
		idx, maxVal := r[0], r[1]
		if a[idx] > maxVal {
			a[idx] = maxVal
		}
	}

	// Forward pass: propagate constraints left to right
	for i := 1; i < n; i++ {
		limit := a[i-1] + diff[i-1]
		if a[i] > limit {
			a[i] = limit
		}
	}

	// Backward pass: propagate constraints right to left
	for i := n - 2; i >= 0; i-- {
		limit := a[i+1] + diff[i]
		if a[i] > limit {
			a[i] = limit
		}
	}

	// Find maximum value
	ans := 0
	for _, v := range a {
		if v > ans {
			ans = v
		}
	}
	return ans
}

func main() {
	// Example 1
	n1 := 10
	restrictions1 := [][]int{{3, 1}, {8, 1}}
	diff1 := []int{2, 2, 3, 1, 4, 5, 1, 1, 2}
	fmt.Println(FindMaximumValueInAConstrainedSequence(n1, restrictions1, diff1)) // Expected: 6

	// Example 2
	n2 := 8
	restrictions2 := [][]int{{3, 2}}
	diff2 := []int{3, 5, 2, 4, 2, 3, 1}
	fmt.Println(FindMaximumValueInAConstrainedSequence(n2, restrictions2, diff2)) // Expected: 12
}
```

## 3799 — Word Squares Ii

```go
package main

// LeetCode #3799: Word Squares II
// https://leetcode.com/problems/word-squares-ii/
// Difficulty: Medium
// Time: O(N^4 * L) | Space: O(N)
// Approach: Generate all valid 4-word squares [top, left, right, bottom]
// satisfying corner constraints: top[0]==left[0], top[3]==right[0],
// bottom[0]==left[3], bottom[3]==right[3]. All 4 words must be distinct.

import (
	"fmt"
	"sort"
)

func WordSquaresIi(words []string) [][]string {
	n := len(words)
	result := [][]string{}

	for ti := 0; ti < n; ti++ {
		top := words[ti]
		for li := 0; li < n; li++ {
			if li == ti {
				continue
			}
			left := words[li]
			if left[0] != top[0] {
				continue
			}
			for ri := 0; ri < n; ri++ {
				if ri == ti || ri == li {
					continue
				}
				right := words[ri]
				if right[0] != top[3] {
					continue
				}
				for bi := 0; bi < n; bi++ {
					if bi == ti || bi == li || bi == ri {
						continue
					}
					bottom := words[bi]
					if bottom[0] != left[3] || bottom[3] != right[3] {
						continue
					}
					square := []string{top, left, right, bottom}
					result = append(result, square)
				}
			}
		}
	}

	// Sort by (top, left, right, bottom) lexicographically
	sort.Slice(result, func(i, j int) bool {
		for k := 0; k < 4; k++ {
			if result[i][k] != result[j][k] {
				return result[i][k] < result[j][k]
			}
		}
		return false
	})

	return result
}

func main() {
	// Example 1
	words1 := []string{"able", "area", "echo", "also"}
	result1 := WordSquaresIi(words1)
	fmt.Println("Result 1:")
	for _, sq := range result1 {
		fmt.Printf("  %v\n", sq)
	}

	// Example 2
	words2 := []string{"ball", "area", "lead", "lady"}
	result2 := WordSquaresIi(words2)
	fmt.Println("Result 2:")
	for _, sq := range result2 {
		fmt.Printf("  %v\n", sq)
	}
}
```

## 3800 — Minimum Cost To Make Two Binary Strings Equal

```go
package main

// LeetCode #3800: Minimum Cost to Make Two Binary Strings Equal
// https://leetcode.com/problems/minimum-cost-to-make-two-binary-strings-equal/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// Approach: Count mismatches and evaluate three strategies: all flips,
// swaps+flips, or cross-swaps+swaps+flips. Take the minimum.

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MinimumCostToMakeTwoBinaryStringsEqual(s string, t string, flipCost int, swapCost int, crossCost int) int {
	n := len(s)
	diff := []int{0, 0} // diff[0] = mismatches where s[i]=='0', diff[1] = mismatches where s[i]=='1'

	for i := 0; i < n; i++ {
		if s[i] != t[i] {
			diff[int(s[i]-'0')]++
		}
	}

	totalDiff := diff[0] + diff[1]
	if totalDiff == 0 {
		return 0
	}

	ans := totalDiff * flipCost // Strategy A: flip all

	// Strategy B: pair mismatches within each s-value, swap them (costs swapCost per pair), flip remainder
	mx, mn := diff[0], diff[1]
	if mx < mn {
		mx, mn = mn, mx
	}
	ans = min(ans, mn*swapCost+(mx-mn)*flipCost)

	// Strategy C: cross-swaps (swap s[i] and t[i] at mismatched positions) + flips
	// Each cross-swap fixes both bits at a mismatched position (s[i] and t[i] trade values)
	crossPairs := totalDiff / 2
	remaining := totalDiff % 2
	ans = min(ans, crossPairs*crossCost+remaining*flipCost)

	return ans
}

func main() {
	// Example 1
	fmt.Println(MinimumCostToMakeTwoBinaryStringsEqual("01000", "10111", 10, 2, 2)) // Expected: 16

	// Example 2
	fmt.Println(MinimumCostToMakeTwoBinaryStringsEqual("001", "110", 2, 100, 100)) // Expected: 6

	// Example 3
	fmt.Println(MinimumCostToMakeTwoBinaryStringsEqual("1010", "1010", 5, 5, 5)) // Expected: 0
}
```

## 3804 — Number Of Centered Subarrays

```go
package main

// LeetCode #3804: Number of Centered Subarrays
// https://leetcode.com/problems/number-of-centered-subarrays/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)
// Approach: For each start index, expand subarrays and track running sum
// with a set of seen elements. A subarray is centered if its sum equals
// at least one element within it.

import "fmt"

func NumberOfCenteredSubarrays(nums []int) int {
	n := len(nums)
	ans := 0

	for i := 0; i < n; i++ {
		sum := 0
		seen := make(map[int]bool)
		for j := i; j < n; j++ {
			seen[nums[j]] = true
			sum += nums[j]
			if seen[sum] {
				ans++
			}
		}
	}

	return ans
}

func main() {
	// Example 1
	fmt.Println(NumberOfCenteredSubarrays([]int{-1, 1, 0})) // Expected: 5

	// Example 2
	fmt.Println(NumberOfCenteredSubarrays([]int{2, -3})) // Expected: 2

	// Example 3
	fmt.Println(NumberOfCenteredSubarrays([]int{1, 2, 3})) // Expected: 3 (all single elements are centered)
}
```

## 3805 — Count Caesar Cipher Pairs

```go
package main

// LeetCode #3805: Count Caesar Cipher Pairs
// https://leetcode.com/problems/count-caesar-cipher-pairs/
// Difficulty: Medium
// Time: O(N * M) | Space: O(N * M)
// Approach: Normalize each string by shifting so that its first character
// becomes 'a'. Strings in the same Caesar-shift equivalence class will have
// the same normalized form. Count pairs using hash map.

import "fmt"

func CountCaesarCipherPairs(words []string) int {
	normalize := func(s string) string {
		shift := int(s[0] - 'a')
		res := make([]byte, len(s))
		for i := 0; i < len(s); i++ {
			res[i] = byte((int(s[i]-'a')-shift+26)%26 + 'a')
		}
		return string(res)
	}

	count := make(map[string]int)
	ans := 0

	for _, w := range words {
		norm := normalize(w)
		ans += count[norm]
		count[norm]++
	}

	return ans
}

func main() {
	// Example 1
	fmt.Println(CountCaesarCipherPairs([]string{"fusion", "layout"})) // Expected: 1

	// Example 2
	fmt.Println(CountCaesarCipherPairs([]string{"ab", "aa", "za", "aa"})) // Expected: 2

	// Example 3
	fmt.Println(CountCaesarCipherPairs([]string{"abc", "bcd", "cde", "xyz"})) // Expected: 3
}
```

## 3807 — Minimum Cost To Repair Edges To Traverse A Graph

```go
package main

// LeetCode #3807: Minimum Cost to Repair Edges to Traverse a Graph
// https://leetcode.com/problems/minimum-cost-to-repair-edges-to-traverse-a-graph/
// Difficulty: Medium [Paid]
// Time: O((N+M) * log M) | Space: O(N+M)
// Approach: Binary search on cost + BFS to check reachability within k edges.

import (
	"container/list"
	"fmt"
	"sort"
)

func MinimumCostToRepairEdgesToTraverseAGraph(n int, edges [][]int, k int) int {
	// Sort edges by repair cost
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] < edges[j][2]
	})

	m := len(edges)

	// Binary search on edge cost threshold
	lo, hi := 0, m
	ans := -1

	for lo <= hi {
		mid := (lo + hi) / 2
		if mid >= m {
			break // can't use edges beyond available
		}
		costLimit := edges[mid][2]

		// Build graph with edges <= costLimit
		adj := make([][]int, n)
		for _, e := range edges {
			if e[2] <= costLimit {
				u, v := e[0], e[1]
				adj[u] = append(adj[u], v)
				adj[v] = append(adj[v], u)
			}
		}

		// BFS to find shortest path from 0 to n-1
		dist := make([]int, n)
		for i := range dist {
			dist[i] = -1
		}
		dist[0] = 0
		q := list.New()
		q.PushBack(0)

		for q.Len() > 0 {
			u := q.Remove(q.Front()).(int)
			if u == n-1 {
				break
			}
			for _, v := range adj[u] {
				if dist[v] == -1 {
					dist[v] = dist[u] + 1
					q.PushBack(v)
				}
			}
		}

		if dist[n-1] != -1 && dist[n-1] <= k {
			ans = costLimit
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}

	return ans
}

func main() {
	// Example 1
	fmt.Println(MinimumCostToRepairEdgesToTraverseAGraph(3, [][]int{{0, 1, 10}, {1, 2, 10}, {0, 2, 100}}, 1)) // Expected: 100

	// Example 2
	edges2 := [][]int{{0, 2, 5}, {2, 3, 6}, {3, 4, 7}, {4, 5, 5}, {0, 1, 10}, {1, 5, 12}, {0, 3, 9}, {1, 2, 8}, {2, 4, 11}}
	fmt.Println(MinimumCostToRepairEdgesToTraverseAGraph(6, edges2, 2)) // Expected: 12

	// Example 3
	fmt.Println(MinimumCostToRepairEdgesToTraverseAGraph(3, [][]int{{0, 1, 1}}, 1)) // Expected: -1
}
```

## 3808 — Find Emotionally Consistent Users

```go
package main

// LeetCode #3808: Find Emotionally Consistent Users
// https://leetcode.com/problems/find-emotionally-consistent-users/
// Difficulty: Medium (SQL problem — implemented in Go)

import (
	"fmt"
	"sort"
)

type Reaction struct {
	UserID    int
	ContentID int
	Reaction  string
}

type ConsistentUser struct {
	UserID           int
	DominantReaction string
	ReactionRatio    float64
}

func main() {
	reactions := []Reaction{
		{1, 1, "like"}, {1, 2, "like"}, {1, 3, "like"}, {1, 4, "like"}, {1, 5, "love"},
		{2, 1, "like"}, {2, 2, "love"}, {2, 3, "wow"}, {2, 4, "sad"}, {2, 5, "angry"},
		{3, 1, "love"}, {3, 2, "love"}, {3, 3, "love"}, {3, 4, "love"}, {3, 5, "love"},
	}

	result := FindEmotionallyConsistentUsers(reactions)
	for _, r := range result {
		fmt.Printf("%d %s %.2f\n", r.UserID, r.DominantReaction, r.ReactionRatio)
	}
}

// Time: O(N + U log U) where N = reactions, U = unique users
// Space: O(N)
func FindEmotionallyConsistentUsers(reactions []Reaction) []ConsistentUser {
	// userID -> reaction -> count
	userReactionCounts := make(map[int]map[string]int)
	userTotal := make(map[int]int)

	for _, r := range reactions {
		if userReactionCounts[r.UserID] == nil {
			userReactionCounts[r.UserID] = make(map[string]int)
		}
		userReactionCounts[r.UserID][r.Reaction]++
		userTotal[r.UserID]++
	}

	var result []ConsistentUser
	for userID, reactionCounts := range userReactionCounts {
		total := userTotal[userID]
		if total < 5 {
			continue
		}

		maxCnt := 0
		dominant := ""
		for reaction, cnt := range reactionCounts {
			if cnt > maxCnt || (cnt == maxCnt && reaction < dominant) {
				maxCnt = cnt
				dominant = reaction
			}
		}

		ratio := float64(maxCnt) / float64(total)
		if ratio >= 0.60 {
			result = append(result, ConsistentUser{
				UserID:           userID,
				DominantReaction: dominant,
				ReactionRatio:    ratio,
			})
		}
	}

	sort.Slice(result, func(i, j int) bool {
		if result[i].ReactionRatio != result[j].ReactionRatio {
			return result[i].ReactionRatio > result[j].ReactionRatio
		}
		return result[i].UserID < result[j].UserID
	})

	return result
}
```

## 3809 — Best Reachable Tower

```go
package main

// LeetCode #3809: Best Reachable Tower
// https://leetcode.com/problems/best-reachable-tower/
// Difficulty: Medium
// Time: O(N) | Space: O(1)
// Approach: One-pass scan checking Manhattan distance and comparing quality.

import "fmt"

func BestReachableTower(towers [][]int, center []int, radius int) []int {
	cx, cy := center[0], center[1]
	bestQuality := -1
	bestX, bestY := -1, -1

	for _, t := range towers {
		x, y, q := t[0], t[1], t[2]
		dist := abs(x-cx) + abs(y-cy)
		if dist > radius {
			continue
		}
		if q > bestQuality || (q == bestQuality && (x < bestX || (x == bestX && y < bestY))) {
			bestQuality = q
			bestX, bestY = x, y
		}
	}

	if bestQuality == -1 {
		return []int{-1, -1}
	}
	return []int{bestX, bestY}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	// Example 1
	fmt.Println(BestReachableTower([][]int{{1, 2, 5}, {2, 1, 7}, {3, 1, 9}}, []int{1, 1}, 2)) // Expected: [3, 1]

	// Example 2
	fmt.Println(BestReachableTower([][]int{{1, 3, 4}, {2, 2, 4}, {4, 4, 7}}, []int{0, 0}, 5)) // Expected: [1, 3]

	// Example 3
	fmt.Println(BestReachableTower([][]int{{5, 6, 8}, {0, 3, 5}}, []int{1, 2}, 1)) // Expected: [-1, -1]
}
```

## 3810 — Minimum Operations To Reach Target Array

```go
package main

// LeetCode #3810: Minimum Operations to Reach Target Array
// https://leetcode.com/problems/minimum-operations-to-reach-target-array/
// Difficulty: Medium
// Time: O(N) | Space: O(max(nums[i]))
// Approach: Track which values need updating. Each operation picks a value x,
// finds all maximal contiguous segments where nums[i]==x, and simultaneously
// updates them to target values.

import "fmt"

func MinimumOperationsToReachTargetArray(nums []int, target []int) int {
	n := len(nums)
	// Map from value to list of positions
	valToPos := make(map[int][]int)
	for i, v := range nums {
		valToPos[v] = append(valToPos[v], i)
	}

	// Track positions that still need changes
	needChange := make(map[int]bool)
	for i := 0; i < n; i++ {
		if nums[i] != target[i] {
			needChange[i] = true
		}
	}

	ops := 0
	for len(needChange) > 0 {
		// Find a value that appears at positions needing change
		var chosenVal int
		found := false
		for v, pos := range valToPos {
			for _, p := range pos {
				if needChange[p] {
					chosenVal = v
					found = true
					break
				}
			}
			if found {
				break
			}
		}
		if !found {
			break
		}

		// Update all maximal contiguous segments of chosenVal
		changed := make(map[int]bool)
		for _, p := range valToPos[chosenVal] {
			if needChange[p] {
				changed[p] = true
			}
		}

		// Process each changed position
		for p := range changed {
			if !needChange[p] {
				continue
			}
			// Expand to full contiguous segment
			l, r := p, p
			for l-1 >= 0 && nums[l-1] == chosenVal {
				l--
			}
			for r+1 < n && nums[r+1] == chosenVal {
				r++
			}
			// Update the segment
			for i := l; i <= r; i++ {
				if needChange[i] {
					nums[i] = target[i]
					delete(needChange, i)
				}
			}
		}
		ops++
	}

	return ops
}

func main() {
	// Example 1
	fmt.Println(MinimumOperationsToReachTargetArray([]int{1, 2, 3}, []int{2, 1, 3})) // Expected: 2

	// Example 2
	fmt.Println(MinimumOperationsToReachTargetArray([]int{4, 1, 4}, []int{5, 1, 4})) // Expected: 1

	// Example 3
	fmt.Println(MinimumOperationsToReachTargetArray([]int{7, 3, 7}, []int{5, 5, 9})) // Expected: 2
}
```

## 3811 — Number Of Alternating Xor Partitions

```go
package main

// LeetCode #3811: Number of Alternating XOR Partitions
// https://leetcode.com/problems/number-of-alternating-xor-partitions/
// Difficulty: Medium
// Time: O(N) | Space: O(N)
// Approach: DP with prefix XOR and two hash maps to track alternating pattern.

import "fmt"

const MOD = 1000000007

func NumberOfAlternatingXorPartitions(nums []int, target1 int, target2 int) int {
	cnt1 := make(map[int]int) // ends with target1
	cnt2 := make(map[int]int) // ends with target2

	cnt2[0] = 1
	pre := 0
	ans := 0

	for _, x := range nums {
		pre ^= x

		a := cnt2[pre^target1] // ways block XOR = target1
		b := cnt1[pre^target2] // ways block XOR = target2

		ans = (a + b) % MOD

		cnt1[pre] = (cnt1[pre] + a) % MOD
		cnt2[pre] = (cnt2[pre] + b) % MOD
	}

	return ans
}

func main() {
	// Example 1
	fmt.Println(NumberOfAlternatingXorPartitions([]int{2, 3, 1, 4}, 1, 5)) // Expected: 1

	// Example 2
	fmt.Println(NumberOfAlternatingXorPartitions([]int{1, 0, 0}, 1, 0)) // Expected: 3

	// Example 3
	fmt.Println(NumberOfAlternatingXorPartitions([]int{1, 2, 3}, 1, 2))
}
```

## 3814 — Maximum Capacity Within Budget

```go
package main

// LeetCode #3814: Maximum Capacity Within Budget
// https://leetcode.com/problems/maximum-capacity-within-budget/
// Difficulty: Medium
// Time: O(N log N) | Space: O(N)
// Approach: Sort by cost, use two pointers to find best pair with cost < budget.

import (
	"fmt"
	"sort"
)

func MaximumCapacityWithinBudget(costs []int, capacity []int, budget int) int {
	n := len(costs)
	type machine struct {
		cost int
		cap  int
	}
	machines := make([]machine, n)
	for i := 0; i < n; i++ {
		machines[i] = machine{costs[i], capacity[i]}
	}

	// Sort by cost
	sort.Slice(machines, func(i, j int) bool {
		if machines[i].cost != machines[j].cost {
			return machines[i].cost < machines[j].cost
		}
		return machines[i].cap > machines[j].cap
	})

	ans := 0

	// Try single machine
	for _, m := range machines {
		if m.cost < budget && m.cap > ans {
			ans = m.cap
		}
	}

	// Try two machines using two pointers
	// For each machine, find best capacity at index < j with costs[i] + costs[j] < budget
	// Track max capacity seen so far for each cost
	maxCapAtCost := make(map[int]int)
	for _, m := range machines {
		prevMax := 0
		if v, ok := maxCapAtCost[m.cost]; ok {
			prevMax = v
		}
		if m.cap > prevMax {
			maxCapAtCost[m.cost] = m.cap
		}
	}

	bestCap := make([]int, n)
	bestCap[0] = machines[0].cap
	for i := 1; i < n; i++ {
		if machines[i].cap > bestCap[i-1] {
			bestCap[i] = machines[i].cap
		} else {
			bestCap[i] = bestCap[i-1]
		}
	}

	for i := 1; i < n; i++ {
		// find a machine j < i with costs[j] + costs[i] < budget
		// binary search for largest cost < budget - costs[i]
		target := budget - machines[i].cost
		if target <= 0 {
			continue
		}
		lo, hi := 0, i-1
		best := -1
		for lo <= hi {
			mid := (lo + hi) / 2
			if machines[mid].cost < target {
				best = mid
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
		if best != -1 {
			total := machines[i].cap + bestCap[best]
			if total > ans {
				ans = total
			}
		}
	}

	return ans
}

func main() {
	// Example 1
	fmt.Println(MaximumCapacityWithinBudget([]int{4, 8, 5, 3}, []int{1, 5, 2, 7}, 8)) // Expected: 8

	// Example 2
	fmt.Println(MaximumCapacityWithinBudget([]int{3, 5, 7, 4}, []int{2, 4, 3, 6}, 7)) // Expected: 6

	// Example 3
	fmt.Println(MaximumCapacityWithinBudget([]int{2, 2, 2}, []int{3, 5, 4}, 5)) // Expected: 9
}
```

## 3815 — Design Auction System

```go
package main

// LeetCode #3815: Design Auction System
// https://leetcode.com/problems/design-auction-system/
// Difficulty: Medium
// Time: O(log N) per operation | Space: O(N)
// Approach: Hash map for user bids, sorted set for highest bid per item.

import (
	"fmt"
)

// AuctionSystem handles bids from users on items.
type AuctionSystem struct {
	// users[userId][itemId] = bidAmount
	users map[int]map[int]int
	// items[itemId] -> sorted list of (bidAmount, userId) using slice
	items map[int][][2]int
}

func Constructor() AuctionSystem {
	return AuctionSystem{
		users: make(map[int]map[int]int),
		items: make(map[int][][2]int),
	}
}

func (as *AuctionSystem) AddBid(userId int, itemId int, bidAmount int) {
	if _, ok := as.users[userId]; !ok {
		as.users[userId] = make(map[int]int)
	}
	// If user already has a bid on this item, remove it first
	if oldBid, ok := as.users[userId][itemId]; ok {
		as.removeBidFromItem(itemId, oldBid, userId)
	}
	as.users[userId][itemId] = bidAmount
	as.items[itemId] = append(as.items[itemId], [2]int{bidAmount, userId})
}

func (as *AuctionSystem) removeBidFromItem(itemId int, bidAmount int, userId int) {
	bids := as.items[itemId]
	for i, b := range bids {
		if b[0] == bidAmount && b[1] == userId {
			as.items[itemId] = append(bids[:i], bids[i+1:]...)
			break
		}
	}
}

func (as *AuctionSystem) UpdateBid(userId int, itemId int, newAmount int) {
	if oldBid, ok := as.users[userId][itemId]; ok {
		as.removeBidFromItem(itemId, oldBid, userId)
		as.users[userId][itemId] = newAmount
		as.items[itemId] = append(as.items[itemId], [2]int{newAmount, userId})
	}
}

func (as *AuctionSystem) RemoveBid(userId int, itemId int) {
	if bidAmount, ok := as.users[userId][itemId]; ok {
		as.removeBidFromItem(itemId, bidAmount, userId)
		delete(as.users[userId], itemId)
	}
}

func (as *AuctionSystem) GetHighestBidder(itemId int) int {
	bids, ok := as.items[itemId]
	if !ok || len(bids) == 0 {
		return -1
	}
	bestAmount, bestUser := -1, -1
	for _, b := range bids {
		if b[0] > bestAmount || (b[0] == bestAmount && b[1] > bestUser) {
			bestAmount = b[0]
			bestUser = b[1]
		}
	}
	return bestUser
}

func main() {
	as := Constructor()

	as.AddBid(1, 7, 5)
	as.AddBid(2, 7, 6)
	fmt.Println(as.GetHighestBidder(7)) // Expected: 2

	as.UpdateBid(1, 7, 8)
	fmt.Println(as.GetHighestBidder(7)) // Expected: 1

	as.RemoveBid(2, 7)
	fmt.Println(as.GetHighestBidder(7)) // Expected: 1

	fmt.Println(as.GetHighestBidder(3)) // Expected: -1
}
```

## 3817 — Good Indices In A Digit String

```go
package main

// LeetCode #3817: Good Indices in a Digit String
// https://leetcode.com/problems/good-indices-in-a-digit-string/
// Difficulty: Medium [Paid]
// Time: O(N * L) where L <= 6 | Space: O(1)
// Approach: For each index, check if a substring ending at i equals decimal representation of i.

import (
	"fmt"
	"strconv"
)

func GoodIndicesInADigitString(s string) []int {
	n := len(s)
	ans := []int{}

	for i := 0; i < n; i++ {
		rep := strconv.Itoa(i)
		l := len(rep)
		if i-l+1 >= 0 && s[i-l+1:i+1] == rep {
			ans = append(ans, i)
		}
	}

	return ans
}

func main() {
	// Example 1
	fmt.Println(GoodIndicesInADigitString("0234567890112")) // Expected: [0 11 12]

	// Example 2
	fmt.Println(GoodIndicesInADigitString("01234")) // Expected: [0 1 2 3 4]

	// Example 3
	fmt.Println(GoodIndicesInADigitString("12345")) // Expected: []
}
```

## 3818 — Minimum Prefix Removal To Make Array Strictly Increasing

```go
package main

// LeetCode #3818: Minimum Prefix Removal to Make Array Strictly Increasing
// https://leetcode.com/problems/minimum-prefix-removal-to-make-array-strictly-increasing/
// Difficulty: Medium
// Time: O(N) | Space: O(1)
// Approach: Scan from right to left to find the longest strictly increasing suffix.

import "fmt"

func MinimumPrefixRemovalToMakeArrayStrictlyIncreasing(nums []int) int {
	n := len(nums)
	for i := n - 1; i > 0; i-- {
		if nums[i-1] >= nums[i] {
			return i
		}
	}
	return 0
}

func main() {
	// Example 1
	fmt.Println(MinimumPrefixRemovalToMakeArrayStrictlyIncreasing([]int{1, -1, 2, 3, 3, 4, 5})) // Expected: 4

	// Example 2
	fmt.Println(MinimumPrefixRemovalToMakeArrayStrictlyIncreasing([]int{4, 3, -2, -5})) // Expected: 3

	// Example 3
	fmt.Println(MinimumPrefixRemovalToMakeArrayStrictlyIncreasing([]int{1, 2, 3, 4})) // Expected: 0
}
```

## 3819 — Rotate Non Negative Elements

```go
package main

// LeetCode #3819: Rotate Non Negative Elements
// https://leetcode.com/problems/rotate-non-negative-elements/
// Difficulty: Medium
// Time: O(N) | Space: O(N)
// Approach: Collect non-negative elements, rotate left by k cyclically, place back.

import "fmt"

func RotateNonNegativeElements(nums []int, k int) []int {
	n := len(nums)
	type pair struct {
		idx int
		val int
	}
	var nonNeg []pair
	for i, v := range nums {
		if v >= 0 {
			nonNeg = append(nonNeg, pair{i, v})
		}
	}

	m := len(nonNeg)
	if m == 0 {
		res := make([]int, n)
		copy(res, nums)
		return res
	}

	result := make([]int, n)
	copy(result, nums)

	// For each position that had a non-negative, put the rotated value
	for i, p := range nonNeg {
		srcIdx := (i + k) % m
		result[p.idx] = nonNeg[srcIdx].val
	}

	return result
}

func main() {
	// Example 1
	fmt.Println(RotateNonNegativeElements([]int{1, -2, 3, -4}, 3)) // Expected: [3 -2 1 -4]

	// Example 2
	fmt.Println(RotateNonNegativeElements([]int{-3, -2, 7}, 1)) // Expected: [-3 -2 7]

	// Example 3
	fmt.Println(RotateNonNegativeElements([]int{5, 4, -9, 6}, 2)) // Expected: [6 5 -9 4]
}
```

## 3820 — Pythagorean Distance Nodes In A Tree

```go
package main

// LeetCode #3820: Pythagorean Distance Nodes in a Tree
// https://leetcode.com/problems/pythagorean-distance-nodes-in-a-tree/
// Difficulty: Medium
// Time: O(N) | Space: O(N)
// Approach: BFS from each target node x, y, z to compute distances,
// then count nodes where sorted distances form a Pythagorean triple.

import (
	"fmt"
	"sort"
)

func PythagoreanDistanceNodesInATree(n int, edges [][]int, x int, y int, z int) int {
	// Build adjacency list
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// BFS to compute distances from a source
	bfs := func(src int) []int {
		dist := make([]int, n)
		for i := range dist {
			dist[i] = -1
		}
		dist[src] = 0
		q := []int{src}
		for len(q) > 0 {
			u := q[0]
			q = q[1:]
			for _, v := range adj[u] {
				if dist[v] == -1 {
					dist[v] = dist[u] + 1
					q = append(q, v)
				}
			}
		}
		return dist
	}

	distX := bfs(x)
	distY := bfs(y)
	distZ := bfs(z)

	ans := 0
	for i := 0; i < n; i++ {
		d := []int{distX[i], distY[i], distZ[i]}
		sort.Ints(d)
		a, b, c := d[0], d[1], d[2]
		if a*a+b*b == c*c {
			ans++
		}
	}

	return ans
}

func main() {
	// Example 1
	fmt.Println(PythagoreanDistanceNodesInATree(4, [][]int{{0, 1}, {0, 2}, {0, 3}}, 1, 2, 3)) // Expected: 3

	// Example 2
	fmt.Println(PythagoreanDistanceNodesInATree(5, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}, 0, 2, 4)) // Expected: ?

	// Example 3
	fmt.Println(PythagoreanDistanceNodesInATree(3, [][]int{{0, 1}, {1, 2}}, 0, 1, 2))
}
```

## 3822 — Design Order Management System

```go
package main

// LeetCode #3822: Design Order Management System
// https://leetcode.com/problems/design-order-management-system/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

type Order struct {
	ID    int
	Type  string
	Price int
	Active bool
}

type OrderManagementSystem struct {
	orders      map[int]*Order
	priceLookup map[int]map[string]map[int]struct{}
}

func Constructor() OrderManagementSystem {
	return OrderManagementSystem{
		orders:      make(map[int]*Order),
		priceLookup: make(map[int]map[string]map[int]struct{}),
	}
}

func (o *OrderManagementSystem) AddOrder(orderID int, orderType string, price int) {
	o.orders[orderID] = &Order{ID: orderID, Type: orderType, Price: price, Active: true}

	if o.priceLookup[price] == nil {
		o.priceLookup[price] = make(map[string]map[int]struct{})
	}
	if o.priceLookup[price][orderType] == nil {
		o.priceLookup[price][orderType] = make(map[int]struct{})
	}
	o.priceLookup[price][orderType][orderID] = struct{}{}
}

func (o *OrderManagementSystem) ModifyOrder(orderID int, newPrice int) {
	order := o.orders[orderID]
	// Remove from old price bucket
	delete(o.priceLookup[order.Price][order.Type], orderID)
	if len(o.priceLookup[order.Price][order.Type]) == 0 {
		delete(o.priceLookup[order.Price], order.Type)
	}
	if len(o.priceLookup[order.Price]) == 0 {
		delete(o.priceLookup, order.Price)
	}

	// Add to new price bucket
	order.Price = newPrice
	if o.priceLookup[newPrice] == nil {
		o.priceLookup[newPrice] = make(map[string]map[int]struct{})
	}
	if o.priceLookup[newPrice][order.Type] == nil {
		o.priceLookup[newPrice][order.Type] = make(map[int]struct{})
	}
	o.priceLookup[newPrice][order.Type][orderID] = struct{}{}
}

func (o *OrderManagementSystem) CancelOrder(orderID int) {
	order := o.orders[orderID]
	order.Active = false

	delete(o.priceLookup[order.Price][order.Type], orderID)
	if len(o.priceLookup[order.Price][order.Type]) == 0 {
		delete(o.priceLookup[order.Price], order.Type)
	}
	if len(o.priceLookup[order.Price]) == 0 {
		delete(o.priceLookup, order.Price)
	}
}

func (o *OrderManagementSystem) GetOrdersAtPrice(orderType string, price int) []int {
	if o.priceLookup[price] == nil || o.priceLookup[price][orderType] == nil {
		return []int{}
	}
	result := make([]int, 0, len(o.priceLookup[price][orderType]))
	for id := range o.priceLookup[price][orderType] {
		result = append(result, id)
	}
	sort.Ints(result)
	return result
}

func main() {
	oms := Constructor()
	oms.AddOrder(1, "buy", 1)
	oms.AddOrder(2, "buy", 1)
	oms.AddOrder(3, "sell", 2)

	fmt.Println(oms.GetOrdersAtPrice("buy", 1)) // [1, 2]

	oms.ModifyOrder(1, 3)
	oms.ModifyOrder(2, 1)
	fmt.Println(oms.GetOrdersAtPrice("buy", 1)) // [2]

	oms.CancelOrder(3)
	oms.CancelOrder(2)
	fmt.Println(oms.GetOrdersAtPrice("buy", 1)) // []
}
```

## 3824 — Minimum K To Reduce Array Within Limit

```go
package main

// LeetCode #3824: Minimum K to Reduce Array Within Limit
// https://leetcode.com/problems/minimum-k-to-reduce-array-within-limit/
// Difficulty: Medium
// Time: O(N log M) where M = max(nums) | Space: O(1)
// Approach: Binary search on k. Check if total operations <= k^2.

import "fmt"

func MinimumKToReduceArrayWithinLimit(nums []int) int {
	// Binary search on k
	maxVal := 0
	for _, v := range nums {
		if v > maxVal {
			maxVal = v
		}
	}

	lo, hi := 1, maxVal
	ans := maxVal

	for lo <= hi {
		mid := (lo + hi) / 2
		ops := 0
		for _, v := range nums {
			ops += (v + mid - 1) / mid // ceil(v / mid)
		}
		if ops <= mid*mid {
			ans = mid
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}

	return ans
}

func main() {
	// Example 1
	fmt.Println(MinimumKToReduceArrayWithinLimit([]int{3, 7, 5})) // Expected: 3

	// Example 2
	fmt.Println(MinimumKToReduceArrayWithinLimit([]int{1})) // Expected: 1

	// Example 3
	fmt.Println(MinimumKToReduceArrayWithinLimit([]int{10, 10, 10}))
}
```

## 3825 — Longest Strictly Increasing Subsequence With Non Zero Bitwise And

```go
package main

// LeetCode #3825: Longest Strictly Increasing Subsequence With Non-Zero Bitwise AND
// https://leetcode.com/problems/longest-strictly-increasing-subsequence-with-non-zero-bitwise-and/
// Difficulty: Medium
// Time: O(log M * N log N) | Space: O(N)
// Approach: For each bit position, filter nums that have that bit set,
// compute LIS on the filtered list. Take max across all bits.

import (
	"fmt"
	"sort"
)

func LongestStrictlyIncreasingSubsequenceWithNonZeroBitwiseAnd(nums []int) int {
	lis := func(arr []int) int {
		tails := []int{}
		for _, x := range arr {
			j := sort.Search(len(tails), func(i int) bool { return tails[i] >= x })
			if j == len(tails) {
				tails = append(tails, x)
			} else {
				tails[j] = x
			}
		}
		return len(tails)
	}

	ans := 0
	// Check up to 31 bits (since nums[i] <= 1e9)
	for bit := 0; bit < 31; bit++ {
		arr := []int{}
		for _, x := range nums {
			if (x>>bit)&1 == 1 {
				arr = append(arr, x)
			}
		}
		if len(arr) > 0 {
			l := lis(arr)
			if l > ans {
				ans = l
			}
		}
	}

	return ans
}

func main() {
	// Example 1
	fmt.Println(LongestStrictlyIncreasingSubsequenceWithNonZeroBitwiseAnd([]int{5, 4, 7})) // Expected: 2

	// Example 2
	fmt.Println(LongestStrictlyIncreasingSubsequenceWithNonZeroBitwiseAnd([]int{2, 3, 6})) // Expected: 3

	// Example 3
	fmt.Println(LongestStrictlyIncreasingSubsequenceWithNonZeroBitwiseAnd([]int{0, 1})) // Expected: 1
}
```

## 3828 — Final Element After Subarray Deletions

```go
package main

// LeetCode #3828: Final Element After Subarray Deletions
// https://leetcode.com/problems/final-element-after-subarray-deletions/
// Difficulty: Medium
// Time: O(1) | Space: O(1)
// Approach: Alice (first) can always keep first or last. Optimal play yields max(first, last).

import "fmt"

func FinalElementAfterSubarrayDeletions(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if nums[0] > nums[len(nums)-1] {
		return nums[0]
	}
	return nums[len(nums)-1]
}

func main() {
	// Example 1
	fmt.Println(FinalElementAfterSubarrayDeletions([]int{1, 5, 2})) // Expected: 2

	// Example 2
	fmt.Println(FinalElementAfterSubarrayDeletions([]int{3, 7})) // Expected: 7

	// Example 3
	fmt.Println(FinalElementAfterSubarrayDeletions([]int{5})) // Expected: 5
}
```

