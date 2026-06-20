# Medium (Sedang) — Problem ��2627

## 2434 — Using A Robot To Print The Lexicographically Smallest String

```go
package main

// LeetCode #2434: Using a Robot to Print the Lexicographically Smallest String
// https://leetcode.com/problems/using-a-robot-to-print-the-lexicographically-smallest-string/
// Difficulty: Medium
// Time: O(n * 26) | Space: O(n)
// Track suffix minimum char. Push to stack if remaining suffix has smaller char.
// Otherwise pop from stack.

import "fmt"

func main() {
	fmt.Println(robotWithString("bac"))   // "abc"
	fmt.Println(robotWithString("bdda"))  // "addb"
}

func robotWithString(s string) string {
	n := len(s)
	suffixMin := make([]byte, n+1)
	suffixMin[n] = 'z' + 1
	for i := n - 1; i >= 0; i-- {
		if s[i] < suffixMin[i+1] {
			suffixMin[i] = s[i]
		} else {
			suffixMin[i] = suffixMin[i+1]
		}
	}

	stack := make([]byte, 0, n)
	res := make([]byte, 0, n)
	for i := 0; i < n; i++ {
		stack = append(stack, s[i])
		for len(stack) > 0 && stack[len(stack)-1] <= suffixMin[i+1] {
			res = append(res, stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}
	}
	for len(stack) > 0 {
		res = append(res, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	return string(res)
}
```

## 2436 — Minimum Split Into Subarrays With Gcd Greater Than One

```go
package main

// LeetCode #2436: Minimum Split Into Subarrays With GCD Greater Than One
// https://leetcode.com/problems/minimum-split-into-subarrays-with-gcd-greater-than-one/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// Greedy: start new subarray when current GCD becomes 1.

import "fmt"

func main() {
	fmt.Println(minimumSplits([]int{12, 6, 3, 14, 8})) // 2
	fmt.Println(minimumSplits([]int{4, 12, 6, 14}))    // 1
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func minimumSplits(nums []int) int {
	ans := 1
	cur := 0
	for _, v := range nums {
		cur = gcd(cur, v)
		if cur == 1 {
			ans++
			cur = v
		}
	}
	return ans
}
```

## 2438 — Range Product Queries Of Powers

```go
package main

// LeetCode #2438: Range Product Queries of Powers
// https://leetcode.com/problems/range-product-queries-of-powers/
// Difficulty: Medium
// Time: O(n + q * n) | Space: O(log n)
// Decompose n into powers of 2. For each query, multiply range.

import "fmt"

func main() {
	fmt.Println(productQueries(15, [][]int{{0, 1}, {2, 2}, {0, 3}})) // [2,4,64]
	fmt.Println(productQueries(2, [][]int{{0, 0}}))                  // [2]
}

const MOD = 1000000007

func productQueries(n int, queries [][]int) []int {
	powers := make([]int, 0)
	pow := 1
	for n > 0 {
		if n&1 == 1 {
			powers = append(powers, pow)
		}
		n >>= 1
		pow <<= 1
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		prod := 1
		for j := q[0]; j <= q[1]; j++ {
			prod = (prod * powers[j]) % MOD
		}
		ans[i] = prod
	}
	return ans
}
```

## 2439 — Minimize Maximum Of Array

```go
package main

// LeetCode #2439: Minimize Maximum of Array
// https://leetcode.com/problems/minimize-maximum-of-array/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// Prefix average approach: we can distribute value to the left.
// The min possible max is the max prefix average (ceil).

import "fmt"

func main() {
	fmt.Println(minimizeArrayValue([]int{3, 7, 1, 6})) // 5
	fmt.Println(minimizeArrayValue([]int{10, 1}))      // 10
}

func minimizeArrayValue(nums []int) int {
	var sum int64
	ans := 0
	for i, v := range nums {
		sum += int64(v)
		avg := int((sum + int64(i)) / int64(i+1)) // ceil division
		if avg > ans {
			ans = avg
		}
	}
	return ans
}
```

## 2442 — Count Number Of Distinct Integers After Reverse Operations

```go
package main

// LeetCode #2442: Count Number of Distinct Integers After Reverse Operations
// https://leetcode.com/problems/count-number-of-distinct-integers-after-reverse-operations/
// Difficulty: Medium
// Time: O(n * log(max)) | Space: O(n)
// For each num, add num and its reverse to a set.

import "fmt"

func main() {
	fmt.Println(countDistinctIntegers([]int{1, 13, 10, 12, 31})) // 6
	fmt.Println(countDistinctIntegers([]int{2, 2, 2}))           // 1
}

func countDistinctIntegers(nums []int) int {
	set := make(map[int]bool)
	for _, v := range nums {
		set[v] = true
		set[reverse(v)] = true
	}
	return len(set)
}

func reverse(n int) int {
	r := 0
	for n > 0 {
		r = r*10 + n%10
		n /= 10
	}
	return r
}
```

## 2443 — Sum Of Number And Its Reverse

```go
package main

// LeetCode #2443: Sum of Number and Its Reverse
// https://leetcode.com/problems/sum-of-number-and-its-reverse/
// Difficulty: Medium
// Time: O(num) | Space: O(1)
// Iterate from 0 to num, check if i + reverse(i) == num.

import "fmt"

func main() {
	fmt.Println(sumOfNumberAndReverse(443)) // true (241+142=443)
	fmt.Println(sumOfNumberAndReverse(63))  // false
	fmt.Println(sumOfNumberAndReverse(181)) // true (90+9=99? wait) (140+41=181)
}

func sumOfNumberAndReverse(num int) bool {
	for i := 0; i <= num; i++ {
		if i+reverse(i) == num {
			return true
		}
	}
	return false
}

func reverse(n int) int {
	r := 0
	for n > 0 {
		r = r*10 + n%10
		n /= 10
	}
	return r
}
```

## 2445 — Number Of Nodes With Value One

```go
package main

// LeetCode #2445: Number of Nodes With Value One
// https://leetcode.com/problems/number-of-nodes-with-value-one/
// Difficulty: Medium
// Time: O(sqrt(n) + q * log n) | Space: O(log n)
// Complete binary tree with n nodes. Each query k toggles depths divisible by k.
// Count nodes at depths with odd toggle count.

import "fmt"

func main() {
	// Complete binary tree with 5 nodes: depths are 1(1), 2(2), 3(2 nodes since incomplete)
	// Queries: k=1 toggles all depths, k=2 toggles depths 2 only, k=3 toggles depth 3 only
	fmt.Println(numberOfNodes(5, []int{1, 2})) // 3 (depths 2,3 toggled)

	// n=6, depths: 1(1), 2(2), 3(3)
	fmt.Println(numberOfNodes(6, []int{3})) // 3
}

func numberOfNodes(n int, queries []int) int {
	// Compute depth range
	maxDepth := 0
	for i := n; i > 0; i >>= 1 {
		maxDepth++
	}
	if maxDepth == 0 {
		return 0
	}

	// Nodes at each depth (excluding last might be incomplete)
	nodesAtDepth := make([]int, maxDepth+1)
	remaining := n
	for d := 1; d <= maxDepth; d++ {
		level := 1 << (d - 1) // 2^(d-1)
		if remaining >= level {
			nodesAtDepth[d] = level
			remaining -= level
		} else {
			nodesAtDepth[d] = remaining
			remaining = 0
		}
	}

	// Toggle depths that are multiples of each query k
	toggle := make([]int, maxDepth+1)
	for _, k := range queries {
		if k <= maxDepth {
			for d := k; d <= maxDepth; d += k {
				toggle[d] ^= 1
			}
		}
	}

	ans := 0
	for d := 1; d <= maxDepth; d++ {
		if toggle[d] == 1 {
			ans += nodesAtDepth[d]
		}
	}
	return ans
}
```

## 2447 — Number Of Subarrays With Gcd Equal To K

```go
package main

// LeetCode #2447: Number of Subarrays With GCD Equal to K
// https://leetcode.com/problems/number-of-subarrays-with-gcd-equal-to-k/
// Difficulty: Medium
// Time: O(n^2) worst-case | Space: O(1)
// For each i, expand j and track GCD. Count when GCD == k.

import "fmt"

func main() {
	fmt.Println(subarrayGCD([]int{9, 3, 1, 2, 6, 3}, 3)) // 4
	fmt.Println(subarrayGCD([]int{4}, 7))                  // 0
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func subarrayGCD(nums []int, k int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		cur := 0
		for j := i; j < len(nums); j++ {
			cur = gcd(cur, nums[j])
			if cur == k {
				ans++
			}
			if cur < k {
				break // GCD only decreases, can't reach k again
			}
		}
	}
	return ans
}
```

## 2450 — Number Of Distinct Binary Strings After Applying Operations

```go
package main

// LeetCode #2450: Number of Distinct Binary Strings After Applying Operations
// https://leetcode.com/problems/number-of-distinct-binary-strings-after-applying-operations/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// Count distinct binary strings reachable by inverting any k-length substring.
// This is equivalent to 2^(number_of_free_variables).

import "fmt"

func main() {
	fmt.Println(distinctBinaryStrings("110", 2)) // 4
	fmt.Println(distinctBinaryStrings("10110", 5)) // 2
}

const MOD = 1000000007

func distinctBinaryStrings(s string, k int) int {
	n := len(s)
	// Number of reachable strings = 2^(n-k+1) if k > 0
	// Because each of the first n-k+1 bits can be independently flipped
	if k > n {
		return 1
	}
	pow := 1
	for i := 0; i < n-k+1; i++ {
		pow = (pow * 2) % MOD
	}
	return pow
}
```

## 2452 — Words Within Two Edits Of Dictionary

```go
package main

// LeetCode #2452: Words Within Two Edits of Dictionary
// https://leetcode.com/problems/words-within-two-edits-of-dictionary/
// Difficulty: Medium
// Time: O(n * m * L) | Space: O(1) where L = word length
// For each query word, check Hamming distance to each dictionary word.

import "fmt"

func main() {
	fmt.Println(twoEditWords([]string{"word", "note", "ants", "wood"}, []string{"wood", "joke", "moat"})) // ["word","note","wood"]
	fmt.Println(twoEditWords([]string{"yes"}, []string{"not"}))                                            // []
}

func twoEditWords(queries []string, dictionary []string) []string {
	ans := make([]string, 0)
	for _, q := range queries {
		for _, d := range dictionary {
			if hammingDist(q, d) <= 2 {
				ans = append(ans, q)
				break
			}
		}
	}
	return ans
}

func hammingDist(a, b string) int {
	dist := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			dist++
		}
	}
	return dist
}
```

## 2453 — Destroy Sequential Targets

```go
package main

// LeetCode #2453: Destroy Sequential Targets
// https://leetcode.com/problems/destroy-sequential-targets/
// Difficulty: Medium
// Time: O(n) | Space: O(n)
// Group nums by nums[i] % space. The group with max size gives max targets.
// Pick smallest nums[i] from that group.

import "fmt"

func main() {
	fmt.Println(destroyTargets([]int{3, 7, 8, 1, 1, 5}, 2)) // 1
	fmt.Println(destroyTargets([]int{1, 3, 5, 2, 4, 6}, 2)) // 1
}

func destroyTargets(nums []int, space int) int {
	freq := make(map[int]int)
	minVal := make(map[int]int)

	for _, v := range nums {
		rem := v % space
		freq[rem]++
		if _, ok := minVal[rem]; !ok || v < minVal[rem] {
			minVal[rem] = v
		}
	}

	maxFreq, ans := 0, 0
	for rem, f := range freq {
		if f > maxFreq || (f == maxFreq && minVal[rem] < ans) {
			maxFreq = f
			ans = minVal[rem]
		}
	}
	return ans
}
```

## 2456 — Most Popular Video Creator

```go
package main

// LeetCode #2456: Most Popular Video Creator
// https://leetcode.com/problems/most-popular-video-creator/
// Difficulty: Medium
// Time: O(n) | Space: O(n)
// Group by creator: track total views, best video (max views, smallest lexicographic).

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(mostPopularCreator([]string{"alice", "bob", "alice", "chris"}, []string{"one", "two", "three", "four"}, []int{5, 10, 5, 4}))
	// [[bob, two], [alice, one]]
}

type Creator struct {
	total    int
	bestID   string
	bestView int
}

func mostPopularCreator(creators []string, ids []string, views []int) [][]string {
	creatorsMap := make(map[string]*Creator)
	maxTotal := 0

	for i, name := range creators {
		c, ok := creatorsMap[name]
		if !ok {
			c = &Creator{bestView: math.MinInt32}
			creatorsMap[name] = c
		}
		c.total += views[i]
		if views[i] > c.bestView || (views[i] == c.bestView && ids[i] < c.bestID) {
			c.bestView = views[i]
			c.bestID = ids[i]
		}
		if c.total > maxTotal {
			maxTotal = c.total
		}
	}

	ans := make([][]string, 0)
	for name, c := range creatorsMap {
		if c.total == maxTotal {
			ans = append(ans, []string{name, c.bestID})
		}
	}
	return ans
}
```

## 2457 — Minimum Addition To Make Integer Beautiful

```go
package main

// LeetCode #2457: Minimum Addition to Make Integer Beautiful
// https://leetcode.com/problems/minimum-addition-to-make-integer-beautiful/
// Difficulty: Medium
// Time: O(log n * log n) | Space: O(log n)
// Try rounding up to higher digit positions until digit sum <= target.

import "fmt"

func main() {
	fmt.Println(makeIntegerBeautiful(16, 6))   // 4 (16+4=20, digit sum 2 <= 6)
	fmt.Println(makeIntegerBeautiful(467, 6))  // 33 (467+33=500, digit sum 5 <= 6)
	fmt.Println(makeIntegerBeautiful(1, 1))    // 0
}

func makeIntegerBeautiful(n int64, target int) int64 {
	if digitSum(n) <= target {
		return 0
	}

	pow10 := int64(10)
	for {
		next := ((n / pow10) + 1) * pow10
		if digitSum(next) <= target {
			return next - n
		}
		pow10 *= 10
	}
}

func digitSum(n int64) int {
	sum := 0
	for n > 0 {
		sum += int(n % 10)
		n /= 10
	}
	return sum
}
```

## 2461 — Maximum Sum Of Distinct Subarrays With Length K

```go
package main

// LeetCode #2461: Maximum Sum of Distinct Subarrays With Length K
// https://leetcode.com/problems/maximum-sum-of-distinct-subarrays-with-length-k/
// Difficulty: Medium
// Time: O(n) | Space: O(n)
// Sliding window with frequency map for distinct check.

import "fmt"

func main() {
	fmt.Println(maximumSubarraySum([]int{1, 5, 4, 2, 9, 9, 9}, 3)) // 15
	fmt.Println(maximumSubarraySum([]int{4, 4, 4}, 3))              // 0
}

func maximumSubarraySum(nums []int, k int) int64 {
	freq := make(map[int]int)
	var sum, ans int64
	dupCount := 0

	for i, v := range nums {
		sum += int64(v)
		freq[v]++
		if freq[v] == 2 {
			dupCount++
		}

		if i >= k {
			left := nums[i-k]
			sum -= int64(left)
			freq[left]--
			if freq[left] == 1 {
				dupCount--
			}
		}

		if i >= k-1 && dupCount == 0 && sum > ans {
			ans = sum
		}
	}
	return ans
}
```

## 2462 — Total Cost To Hire K Workers

```go
package main

// LeetCode #2462: Total Cost to Hire K Workers
// https://leetcode.com/problems/total-cost-to-hire-k-workers/
// Difficulty: Medium
// Time: O((candidates + k) log candidates) | Space: O(candidates)
// Two min-heaps: left and right. Each round pick cheaper candidate.

import (
	"container/heap"
	"fmt"
)

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

func main() {
	fmt.Println(totalCost([]int{17, 12, 10, 2, 7, 2, 11, 20, 8}, 3, 4)) // 11
	fmt.Println(totalCost([]int{1, 2, 4, 1}, 3, 3))                      // 4
}

func totalCost(costs []int, k int, candidates int) int64 {
	n := len(costs)
	leftH := &MinHeap{}
	rightH := &MinHeap{}
	heap.Init(leftH)
	heap.Init(rightH)

	left, right := 0, n-1
	for i := 0; i < candidates && left <= right; i++ {
		heap.Push(leftH, costs[left])
		left++
	}
	for i := 0; i < candidates && left <= right; i++ {
		heap.Push(rightH, costs[right])
		right--
	}

	var total int64
	for i := 0; i < k; i++ {
		if rightH.Len() == 0 || (leftH.Len() > 0 && (*leftH)[0] <= (*rightH)[0]) {
			total += int64(heap.Pop(leftH).(int))
			if left <= right {
				heap.Push(leftH, costs[left])
				left++
			}
		} else {
			total += int64(heap.Pop(rightH).(int))
			if left <= right {
				heap.Push(rightH, costs[right])
				right--
			}
		}
	}
	return total
}
```

## 2464 — Minimum Subarrays In A Valid Split

```go
package main

// LeetCode #2464: Minimum Subarrays in a Valid Split
// https://leetcode.com/problems/minimum-subarrays-in-a-valid-split/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// Greedy: start new subarray when GCD becomes 1.

import "fmt"

func main() {
	fmt.Println(validSplit([]int{2, 6, 3, 4, 3})) // 2
	fmt.Println(validSplit([]int{3, 5}))           // 2
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func validSplit(nums []int) int {
	ans := 1
	cur := 0
	for _, v := range nums {
		cur = gcd(cur, v)
		if cur == 1 {
			ans++
			cur = v
		}
	}
	if cur == 1 {
		// If last GCD is 1, we started but couldn't close
		// This shouldn't happen for valid input per problem constraints
	}
	return ans
}
```

## 2466 — Count Ways To Build Good Strings

```go
package main

// LeetCode #2466: Count Ways To Build Good Strings
// https://leetcode.com/problems/count-ways-to-build-good-strings/
// Difficulty: Medium
// Time: O(high) | Space: O(high)
// DP: dp[i] = ways to build string of length i.
// dp[i] = dp[i-zero] + dp[i-one] (if i >= zero/one).

import "fmt"

func main() {
	fmt.Println(countGoodStrings(3, 3, 1, 1)) // 8
	fmt.Println(countGoodStrings(2, 3, 1, 2)) // 5
}

const MOD = 1000000007

func countGoodStrings(low int, high int, zero int, one int) int {
	dp := make([]int, high+1)
	dp[0] = 1
	ans := 0

	for i := 1; i <= high; i++ {
		if i >= zero {
			dp[i] = (dp[i] + dp[i-zero]) % MOD
		}
		if i >= one {
			dp[i] = (dp[i] + dp[i-one]) % MOD
		}
		if i >= low {
			ans = (ans + dp[i]) % MOD
		}
	}
	return ans
}
```

## 2467 — Most Profitable Path In A Tree

```go
package main

// LeetCode #2467: Most Profitable Path in a Tree
// https://leetcode.com/problems/most-profitable-path-in-a-tree/
// Difficulty: Medium
// Time: O(n) | Space: O(n)
// Bob moves to root (fixed path). Alice moves from root to leaf,
// collecting max profit considering time-shared nodes.

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(mostProfitablePath([][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}}, 3, []int{-2, 4, 2, -4, 6}))
	// 6

	fmt.Println(mostProfitablePath([][]int{{0, 1}}, 1, []int{-7280, 2350}))
	// -7280
}

func mostProfitablePath(edges [][]int, bob int, amount []int) int {
	n := len(amount)
	graph := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	// Bob's path from bob to 0
	parent := make([]int, n)
	var dfsParent func(u, p int)
	dfsParent = func(u, p int) {
		parent[u] = p
		for _, v := range graph[u] {
			if v != p {
				dfsParent(v, u)
			}
		}
	}
	dfsParent(0, -1)

	// Bob's arrival time at each node
	bobTime := make([]int, n)
	for i := range bobTime {
		bobTime[i] = math.MaxInt32
	}
	t := 0
	for u := bob; u != -1; u = parent[u] {
		bobTime[u] = t
		t++
	}

	ans := math.MinInt32
	var dfsAlice func(u, p, t, profit int)
	dfsAlice = func(u, p, t, profit int) {
		if t < bobTime[u] {
			profit += amount[u]
		} else if t == bobTime[u] {
			profit += amount[u] / 2
		}
		isLeaf := true
		for _, v := range graph[u] {
			if v != p {
				isLeaf = false
				dfsAlice(v, u, t+1, profit)
			}
		}
		if isLeaf && profit > ans {
			ans = profit
		}
	}
	dfsAlice(0, -1, 0, 0)
	return ans
}
```

## 2470 — Number Of Subarrays With Lcm Equal To K

```go
package main

// LeetCode #2470: Number of Subarrays With LCM Equal to K
// https://leetcode.com/problems/number-of-subarrays-with-lcm-equal-to-k/
// Difficulty: Medium
// Time: O(n^2) worst-case | Space: O(1)
// For each start, expand and track LCM.

import "fmt"

func main() {
	fmt.Println(subarrayLCM([]int{3, 6, 2, 7, 1}, 6)) // 4
	fmt.Println(subarrayLCM([]int{3}, 2))              // 0
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

func subarrayLCM(nums []int, k int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		cur := 1
		for j := i; j < len(nums); j++ {
			cur = lcm(cur, nums[j])
			if cur == k {
				ans++
			}
			if cur > k {
				break
			}
		}
	}
	return ans
}
```

## 2471 — Minimum Number Of Operations To Sort A Binary Tree By Level

```go
package main

// LeetCode #2471: Minimum Number of Operations to Sort a Binary Tree by Level
// https://leetcode.com/problems/minimum-number-of-operations-to-sort-a-binary-tree-by-level/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)
// BFS level-order. For each level, count min swaps to sort (cycle decomposition).

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{1,
		&TreeNode{4,
			&TreeNode{7, nil, nil},
			&TreeNode{6, nil, nil},
		},
		&TreeNode{3,
			&TreeNode{8, nil, nil},
			&TreeNode{5, nil, nil},
		},
	}
	// Level 1: [1] sorted. Level 2: [4,3] -> swap, 1 op. Level 3: [7,6,8,5] -> 2 ops
	fmt.Println(minimumOperations(root)) // 3

	root2 := &TreeNode{1,
		&TreeNode{2, nil, nil},
		&TreeNode{3, nil, nil},
	}
	fmt.Println(minimumOperations(root2)) // 0
}

func minimumOperations(root *TreeNode) int {
	q := []*TreeNode{root}
	ans := 0
	for len(q) > 0 {
		n := len(q)
		vals := make([]int, n)
		for i := 0; i < n; i++ {
			vals[i] = q[i].Val
		}

		// Count min swaps to sort vals
		ans += minSwaps(vals)

		next := make([]*TreeNode, 0)
		for _, node := range q {
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		q = next
	}
	return ans
}

func minSwaps(arr []int) int {
	n := len(arr)
	sorted := make([]int, n)
	copy(sorted, arr)
	sort.Ints(sorted)

	pos := make(map[int]int)
	for i, v := range arr {
		pos[v] = i
	}

	visited := make([]bool, n)
	swaps := 0
	for i := 0; i < n; i++ {
		if visited[i] || arr[i] == sorted[i] {
			continue
		}
		cycle := 0
		j := i
		for !visited[j] {
			visited[j] = true
			j = pos[sorted[j]]
			cycle++
		}
		swaps += cycle - 1
	}
	return swaps
}
```

## 2473 — Minimum Cost To Buy Apples

```go
package main

// LeetCode #2473: Minimum Cost to Buy Apples
// https://leetcode.com/problems/minimum-cost-to-buy-apples/
// Difficulty: Medium
// Time: O(n * (n + m) log n) | Space: O(n + m)
// Run Dijkstra from each start node to find min round-trip cost for one apple.

import (
	"container/heap"
	"fmt"
	"math"
)

type Edge struct {
	to, cost int
}

type Item struct {
	node, dist int
}

type PriorityQueue []Item

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].dist < pq[j].dist }
func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x any) { *pq = append(*pq, x.(Item)) }
func (pq *PriorityQueue) Pop() any {
	old := *pq; n := len(old); x := old[n-1]; *pq = old[:n-1]; return x
}

func main() {
	fmt.Println(minCost(5, [][]int{{1,2,2},{2,3,4},{3,4,6}}, []int{3,5,8,12,20}, 2, 1))
	fmt.Println(minCost(4, [][]int{{1,2,5},{2,3,1},{3,4,2}}, []int{10,10,10,10}, 3, 2))
}

func minCost(n int, roads [][]int, appleCost []int, k int, start int) []int64 {
	graph := make([][]Edge, n)
	for _, r := range roads {
		u, v, c := r[0]-1, r[1]-1, r[2]
		graph[u] = append(graph[u], Edge{v, c})
		graph[v] = append(graph[v], Edge{u, c})
	}

	ans := make([]int64, n)
	for s := 0; s < n; s++ {
		dist := dijkstra(graph, s, n)
		minCost := int64(math.MaxInt64)
		for i := 0; i < n; i++ {
			if dist[i] == math.MaxInt64 {
				continue
			}
			total := int64(appleCost[i]) + int64(dist[i])*(1+int64(k))
			if total < minCost {
				minCost = total
			}
		}
		ans[s] = minCost
	}
	return ans
}

func dijkstra(graph [][]Edge, src, n int) []int {
	dist := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt64
	}
	dist[src] = 0
	pq := &PriorityQueue{{src, 0}}
	heap.Init(pq)

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(Item)
		if cur.dist > dist[cur.node] {
			continue
		}
		for _, e := range graph[cur.node] {
			if nd := cur.dist + e.cost; nd < dist[e.to] {
				dist[e.to] = nd
				heap.Push(pq, Item{e.to, nd})
			}
		}
	}
	return dist
}
```

## 2476 — Closest Nodes Queries In A Binary Search Tree

```go
package main

// LeetCode #2476: Closest Nodes Queries in a Binary Search Tree
// https://leetcode.com/problems/closest-nodes-queries-in-a-binary-search-tree/
// Difficulty: Medium
// Time: O(n + q log n) | Space: O(n)
// Inorder traversal to sorted array, then binary search each query.

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{6,
		&TreeNode{2,
			&TreeNode{1, nil, nil},
			&TreeNode{4, nil, nil},
		},
		&TreeNode{13,
			&TreeNode{9, nil, nil},
			&TreeNode{15, nil, nil},
		},
	}
	fmt.Println(closestNodes(root, []int{2, 5, 16})) // [[2,2],[2,4],[15,15]]

	root2 := &TreeNode{4, nil, nil}
	fmt.Println(closestNodes(root2, []int{1, 5})) // [[-1,4],[4,-1]]
}

func closestNodes(root *TreeNode, queries []int) [][]int {
	vals := make([]int, 0)
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		vals = append(vals, node.Val)
		inorder(node.Right)
	}
	inorder(root)

	ans := make([][]int, len(queries))
	for i, q := range queries {
		// Find smallest >= q
		idx := sort.SearchInts(vals, q)
		minVal := -1
		if idx < len(vals) {
			minVal = vals[idx]
		}
		maxVal := -1
		if idx > 0 {
			maxVal = vals[idx-1]
		}
		if idx < len(vals) && vals[idx] == q {
			// exact match
			minVal, maxVal = q, q
		}
		// Fix: if exact match, both are q
		if minVal == q && maxVal == -1 {
			maxVal = q
		}
		if minVal == q {
			maxVal = q
			minVal = q
		} else {
			// min is smallest >= q, max is largest < q
			if idx < len(vals) {
				minVal = vals[idx]
			}
			if idx > 0 {
				maxVal = vals[idx-1]
			}
		}
		ans[i] = []int{maxVal, minVal}
	}
	return ans
}
```

## 2477 — Minimum Fuel Cost To Report To The Capital

```go
package main

// LeetCode #2477: Minimum Fuel Cost to Report to the Capital
// https://leetcode.com/problems/minimum-fuel-cost-to-report-to-the-capital/
// Difficulty: Medium
// Time: O(n) | Space: O(n)
// Tree DP: Each person travels from leaf towards root (city 0).
// Accumulate people from children, compute fuel for each edge.

import "fmt"

func main() {
	fmt.Println(minimumFuelCost([][]int{{0, 1}, {0, 2}, {0, 3}}, 1)) // 3
	fmt.Println(minimumFuelCost([][]int{{3, 1}, {3, 2}, {1, 0}, {0, 4}, {0, 5}, {4, 6}}, 2)) // 7
}

func minimumFuelCost(roads [][]int, seats int) int64 {
	n := len(roads) + 1
	graph := make([][]int, n)
	for _, r := range roads {
		a, b := r[0], r[1]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	var ans int64
	var dfs func(u, parent int) int
	dfs = func(u, parent int) int {
		people := 1 // each node has 1 representative
		for _, v := range graph[u] {
			if v != parent {
				people += dfs(v, u)
			}
		}
		if u != 0 {
			// cars needed = ceil(people / seats)
			cars := (people + seats - 1) / seats
			ans += int64(cars)
		}
		return people
	}
	dfs(0, -1)
	return ans
}
```

## 2482 — Difference Between Ones And Zeros In Row And Column

```go
package main

// LeetCode #2482: Difference Between Ones and Zeros in Row and Column
// https://leetcode.com/problems/difference-between-ones-and-zeros-in-row-and-column/
// Difficulty: Medium
// Time: O(m * n) | Space: O(m + n)
// diff[i][j] = onesRow[i] + onesCol[j] - zerosRow[i] - zerosCol[j]
// = 2*onesRow[i] + 2*onesCol[j] - m - n

import "fmt"

func main() {
	fmt.Println(onesMinusZeros([][]int{{0, 1, 1}, {1, 0, 1}, {0, 0, 1}}))
	// [[0,0,4],[0,0,4],[-2,-2,2]]

	fmt.Println(onesMinusZeros([][]int{{1, 1, 1}, {1, 1, 1}}))
	// [[5,5,5],[5,5,5]]
}

func onesMinusZeros(grid [][]int) [][]int {
	m, n := len(grid), len(grid[0])
	rowOnes := make([]int, m)
	colOnes := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				rowOnes[i]++
				colOnes[j]++
			}
		}
	}

	diff := make([][]int, m)
	for i := 0; i < m; i++ {
		diff[i] = make([]int, n)
		for j := 0; j < n; j++ {
			diff[i][j] = 2*rowOnes[i] + 2*colOnes[j] - m - n
		}
	}
	return diff
}
```

## 2483 — Minimum Penalty For A Shop

```go
package main

// LeetCode #2483: Minimum Penalty for a Shop
// https://leetcode.com/problems/minimum-penalty-for-a-shop/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// Penalty at hour i = 'N' before i + 'Y' after/at i. Sweep left to right.

import "fmt"

func main() {
	fmt.Println(bestClosingTime("YYNY")) // 2
	fmt.Println(bestClosingTime("NNNNN")) // 0
	fmt.Println(bestClosingTime("YYYY"))  // 4
}

func bestClosingTime(customers string) int {
	// Start: close at hour 0
	penalty := 0
	for _, ch := range customers {
		if ch == 'Y' {
			penalty++
		}
	}
	minPenalty := penalty
	bestHour := 0

	// Try closing at hour 1..n
	for i, ch := range customers {
		if ch == 'Y' {
			penalty--
		} else {
			penalty++
		}
		if penalty < minPenalty {
			minPenalty = penalty
			bestHour = i + 1
		}
	}
	return bestHour
}
```

## 2486 — Append Characters To String To Make Subsequence

```go
package main

// LeetCode #2486: Append Characters to String to Make Subsequence
// https://leetcode.com/problems/append-characters-to-string-to-make-subsequence/
// Difficulty: Medium
// Time: O(n + m) | Space: O(1)
// Two pointers: match as much of t in s as possible.

import "fmt"

func main() {
	fmt.Println(appendCharacters("coaching", "coding")) // 4
	fmt.Println(appendCharacters("abcde", "a"))         // 0
	fmt.Println(appendCharacters("z", "abcde"))         // 5
}

func appendCharacters(s string, t string) int {
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			j++
		}
		i++
	}
	return len(t) - j
}
```

## 2487 — Remove Nodes From Linked List

```go
package main

// LeetCode #2487: Remove Nodes From Linked List
// https://leetcode.com/problems/remove-nodes-from-linked-list/
// Difficulty: Medium
// Time: O(n) | Space: O(n)
// Reverse list, track max so far, keep nodes >= max.

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// 5 -> 2 -> 13 -> 3 -> 8
	head := &ListNode{5, &ListNode{2, &ListNode{13, &ListNode{3, &ListNode{8, nil}}}}}
	res := removeNodes(head)
	for res != nil {
		fmt.Print(res.Val, " ") // 13 8
		res = res.Next
	}
	fmt.Println()
}

func removeNodes(head *ListNode) *ListNode {
	// Reverse
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	// Keep nodes >= max so far
	dummy := &ListNode{Next: prev}
	cur = dummy.Next
	maxSoFar := cur.Val
	for cur != nil && cur.Next != nil {
		if cur.Next.Val < maxSoFar {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
			maxSoFar = cur.Val
		}
	}

	// Reverse back
	var prev2 *ListNode
	cur = dummy.Next
	for cur != nil {
		next := cur.Next
		cur.Next = prev2
		prev2 = cur
		cur = next
	}
	return prev2
}
```

## 2489 — Number Of Substrings With Fixed Ratio

```go
package main

// LeetCode #2489: Number of Substrings With Fixed Ratio
// https://leetcode.com/problems/number-of-substrings-with-fixed-ratio/
// Difficulty: Medium
// Time: O(n) | Space: O(n)
// Count substrings where count('0') : count('1') = num1 : num2.
// Transform: track (num2 * cnt0 - num1 * cnt1), count equal values.

import "fmt"

func main() {
	fmt.Println(fixedRatio("01001", 2, 3)) // 2
	fmt.Println(fixedRatio("0000", 1, 1))  // 0
}

func fixedRatio(s string, num1 int, num2 int) int64 {
	prefix := make(map[int]int64)
	prefix[0] = 1
	var cnt0, cnt1 int64
	var ans int64

	for _, ch := range s {
		if ch == '0' {
			cnt0++
		} else {
			cnt1++
		}
		key := num2*int(cnt0) - num1*int(cnt1)
		ans += prefix[key]
		prefix[key]++
	}
	return ans
}
```

## 2491 — Divide Players Into Teams Of Equal Skill

```go
package main

// LeetCode #2491: Divide Players Into Teams of Equal Skill
// https://leetcode.com/problems/divide-players-into-teams-of-equal-skill/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)
// Sort, pair smallest with largest. Each pair sum must equal same target.

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(dividePlayers([]int{3, 2, 5, 1, 3, 4})) // 22
	fmt.Println(dividePlayers([]int{1, 1, 2, 3}))       // -1
}

func dividePlayers(skill []int) int64 {
	sort.Ints(skill)
	n := len(skill)
	target := skill[0] + skill[n-1]
	var sum int64
	for i := 0; i < n/2; i++ {
		if skill[i]+skill[n-1-i] != target {
			return -1
		}
		sum += int64(skill[i]) * int64(skill[n-1-i])
	}
	return sum
}
```

## 2492 — Minimum Score Of A Path Between Two Cities

```go
package main

// LeetCode #2492: Minimum Score of a Path Between Two Cities
// https://leetcode.com/problems/minimum-score-of-a-path-between-two-cities/
// Difficulty: Medium
// Time: O(n + m) | Space: O(n + m)
// DFS from city 1, find min edge in its connected component (must include city n).

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minScore(4, [][]int{{1, 2, 9}, {2, 3, 6}, {2, 4, 5}, {1, 4, 7}})) // 5
	fmt.Println(minScore(4, [][]int{{1, 2, 2}, {1, 3, 4}, {3, 4, 7}}))            // 2
}

func minScore(n int, roads [][]int) int {
	graph := make([][][2]int, n+1)
	for _, r := range roads {
		a, b, d := r[0], r[1], r[2]
		graph[a] = append(graph[a], [2]int{b, d})
		graph[b] = append(graph[b], [2]int{a, d})
	}

	visited := make([]bool, n+1)
	ans := math.MaxInt32

	var dfs func(u int)
	dfs = func(u int) {
		visited[u] = true
		for _, edge := range graph[u] {
			v, d := edge[0], edge[1]
			if d < ans {
				ans = d
			}
			if !visited[v] {
				dfs(v)
			}
		}
	}
	dfs(1)

	return ans
}
```

## 2495 — Number Of Subarrays Having Even Product

```go
package main

// LeetCode #2495: Number of Subarrays Having Even Product
// https://leetcode.com/problems/number-of-subarrays-having-even-product/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// Product is even if at least one even element.
// Count subarrays from last even position.

import "fmt"

func main() {
	fmt.Println(evenProduct([]int{1, 2, 3})) // 3 ([2], [1,2], [2,3], [1,2,3])
	fmt.Println(evenProduct([]int{1, 3, 5})) // 0
}

func evenProduct(nums []int) int64 {
	var ans int64
	lastEven := -1
	for i, v := range nums {
		if v%2 == 0 {
			lastEven = i
		}
		if lastEven != -1 {
			ans += int64(lastEven + 1)
		}
	}
	return ans
}
```

## 2497 — Maximum Star Sum Of A Graph

```go
package main

// LeetCode #2497: Maximum Star Sum of a Graph
// https://leetcode.com/problems/maximum-star-sum-of-a-graph/
// Difficulty: Medium
// Time: O(n + m log k) | Space: O(n + m)
// For each node, sort neighbor values descending, take top k positive.

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxStarSum([]int{1, 2, 3, 4, 10, -10, -20}, [][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}, {3, 5}, {3, 6}}, 2))
	// 16

	fmt.Println(maxStarSum([]int{-5}, [][]int{}, 0))
	// -5
}

func maxStarSum(vals []int, edges [][]int, k int) int {
	n := len(vals)
	neighbors := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		neighbors[u] = append(neighbors[u], vals[v])
		neighbors[v] = append(neighbors[v], vals[u])
	}

	ans := vals[0]
	for i := 0; i < n; i++ {
		sort.Sort(sort.Reverse(sort.IntSlice(neighbors[i])))
		sum := vals[i]
		for j := 0; j < k && j < len(neighbors[i]); j++ {
			if neighbors[i][j] > 0 {
				sum += neighbors[i][j]
			}
		}
		if sum > ans {
			ans = sum
		}
	}
	return ans
}
```

## 2498 — Frog Jump Ii

```go
package main

// LeetCode #2498: Frog Jump II
// https://leetcode.com/problems/frog-jump-ii/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// Frog can jump forward, max distance = max of (stones[i+2] - stones[i]) for i in 0..n-3
// and also stones[1] - stones[0] and stones[n-1] - stones[n-2].

import "fmt"

func main() {
	fmt.Println(maxJump([]int{0, 2, 5, 6, 7})) // 5
	fmt.Println(maxJump([]int{0, 3, 9}))        // 9
}

func maxJump(stones []int) int {
	n := len(stones)
	ans := stones[1] - stones[0]
	if n > 2 {
		ans = stones[n-1] - stones[n-2]
	}
	for i := 2; i < n; i++ {
		diff := stones[i] - stones[i-2]
		if diff > ans {
			ans = diff
		}
	}
	return ans
}
```

## 2501 — Longest Square Streak In An Array

```go
package main

// LeetCode #2501: Longest Square Streak in an Array
// https://leetcode.com/problems/longest-square-streak-in-an-array/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)
// Sort, use map to track longest streak ending at each value.

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(longestSquareStreak([]int{4, 3, 6, 16, 8, 2})) // 3 (2 -> 4 -> 16)
	fmt.Println(longestSquareStreak([]int{2, 3, 5, 6, 7}))     // -1
}

func longestSquareStreak(nums []int) int {
	sort.Ints(nums)
	dp := make(map[int]int)
	ans := -1

	for _, v := range nums {
		root := intSqrt(v)
		if root*root == v {
			if prev, ok := dp[root]; ok {
				dp[v] = prev + 1
			} else {
				dp[v] = 1
			}
		} else {
			dp[v] = 1
		}
		if dp[v] > ans {
			ans = dp[v]
		}
	}
	if ans < 2 {
		return -1
	}
	return ans
}

func intSqrt(n int) int {
	lo, hi := 1, n
	for lo <= hi {
		mid := (lo + hi) / 2
		if mid*mid == n {
			return mid
		} else if mid*mid < n {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return 0
}
```

## 2502 — Design Memory Allocator

```go
package main

// LeetCode #2502: Design Memory Allocator
// https://leetcode.com/problems/design-memory-allocator/
// Difficulty: Medium
// Time: O(n) per allocate | O(1) per free
// Array-based allocation: find first free block of size.

import "fmt"

type Allocator struct {
	mem []int
}

func main() {
	alloc := Constructor(10)
	fmt.Println(alloc.Allocate(1, 1))  // 0
	fmt.Println(alloc.Allocate(1, 2))  // 1
	fmt.Println(alloc.Allocate(1, 1))  // 2
	fmt.Println(alloc.Free(1))          // 2
	fmt.Println(alloc.Allocate(2, 1))  // 0 (freed 0 and 2)
	fmt.Println(alloc.Free(2))          // 1
	fmt.Println(alloc.Allocate(3, 1))  // -1 (need 3, only 1 free block of size 1)
}

func Constructor(n int) Allocator {
	return Allocator{mem: make([]int, n)}
}

func (a *Allocator) Allocate(size int, mID int) int {
	n := len(a.mem)
	for i := 0; i < n; {
		if a.mem[i] != 0 {
			i++
			continue
		}
		j := i
		for j < n && a.mem[j] == 0 {
			j++
		}
		if j-i >= size {
			for k := i; k < i+size; k++ {
				a.mem[k] = mID
			}
			return i
		}
		i = j
	}
	return -1
}

func (a *Allocator) Free(mID int) int {
	count := 0
	for i := range a.mem {
		if a.mem[i] == mID {
			a.mem[i] = 0
			count++
		}
	}
	return count
}
```

## 2505 — Bitwise Or Of All Subsequence Sums

```go
package main

// LeetCode #2505: Bitwise OR of All Subsequence Sums
// https://leetcode.com/problems/bitwise-or-of-all-subsequence-sums/
// Difficulty: Medium
// Time: O(n * 20) | Space: O(1)
// A bit is achievable if any number has that bit, or can be formed by combination.
// Result = OR of all prefix sums? Actually: any sum that can be formed = OR of
// all elements and their combinations. Answer = OR of all elements (since we can
// always form any single element sum via a subsequence of size 1).

import "fmt"

func main() {
	fmt.Println(subsequenceSumOr([]int{2, 1, 4})) // 7 (bits 0,1,2)
	fmt.Println(subsequenceSumOr([]int{2, 3}))    // 7 (sums: 0,2,3,5 -> OR = 7)
}

func subsequenceSumOr(nums []int) int {
	ans := 0
	for _, v := range nums {
		ans |= v
	}
	return ans
}
```

## 2507 — Smallest Value After Replacing With Sum Of Prime Factors

```go
package main

// LeetCode #2507: Smallest Value After Replacing With Sum of Prime Factors
// https://leetcode.com/problems/smallest-value-after-replacing-with-sum-of-prime-factors/
// Difficulty: Medium
// Time: O(sqrt(n) * iterations) | Space: O(1)
// Replace n with sum of its prime factors (with multiplicity) until stable.

import "fmt"

func main() {
	fmt.Println(smallestValue(15)) // 5 (15=3*5 -> 8=2*2*2 -> 6=2*3 -> 5=prime)
	fmt.Println(smallestValue(4))  // 4 (4=2*2 -> 4, stable)
}

func smallestValue(n int) int {
	for {
		sum := primeFactorSum(n)
		if sum == n {
			return n
		}
		n = sum
	}
}

func primeFactorSum(n int) int {
	sum := 0
	// Factor 2
	for n%2 == 0 {
		sum += 2
		n /= 2
	}
	// Odd factors
	for f := 3; f*f <= n; f += 2 {
		for n%f == 0 {
			sum += f
			n /= f
		}
	}
	if n > 1 {
		sum += n
	}
	return sum
}
```

## 2510 — Check If There Is A Path With Equal Number Of 0s And 1s

```go
package main

// LeetCode #2510: Check if There is a Path With Equal Number of 0's And 1's
// https://leetcode.com/problems/check-if-there-is-a-path-with-equal-number-of-0s-and-1s/
// Difficulty: Medium
// Time: O(m * n * (m+n)) | Space: O(m * n)
// DP: dp[i][j][diff] = reachable with given (ones - zeros) balance.
// Optimize: total cells must be even, and we need diff=0 at end.
// Since grid size is small, we can use 2D boolean DP with reachable sums.

import "fmt"

func main() {
	fmt.Println(isThereAPath([][]int{{0, 1, 0}, {1, 1, 0}, {0, 0, 1}})) // false
	fmt.Println(isThereAPath([][]int{{0, 1}, {1, 0}}))                   // true
}

func isThereAPath(grid [][]int) bool {
	r, c := len(grid), len(grid[0])
	total := r + c - 1
	if total%2 != 0 {
		return false
	}
	target := total / 2 // need this many ones

	// dp[i][j][k] = reachable with k ones
	dp := make([][][]bool, r)
	for i := range dp {
		dp[i] = make([][]bool, c)
		for j := range dp[i] {
			dp[i][j] = make([]bool, target+1)
		}
	}

	ones := 0
	if grid[0][0] == 1 {
		ones = 1
	}
	if ones <= target {
		dp[0][0][ones] = true
	}

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if i == 0 && j == 0 {
				continue
			}
			val := grid[i][j]
			for k := 0; k <= target; k++ {
				prev := false
				if i > 0 {
					prev = prev || dp[i-1][j][k]
				}
				if j > 0 {
					prev = prev || dp[i][j-1][k]
				}
				if prev && k+val <= target {
					dp[i][j][k+val] = true
				}
			}
		}
	}
	return dp[r-1][c-1][target]
}
```

## 2512 — Reward Top K Students

```go
package main

// LeetCode #2512: Reward Top K Students
// https://leetcode.com/problems/reward-top-k-students/
// Difficulty: Medium
// Time: O(n * L + n log n) | Space: O(n)
// Score each student by positive/negative keywords, sort, return top K IDs.

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(topStudents([]string{"smart", "brilliant", "studious"}, []string{"not"}, [][]string{
		{"this", "student", "is", "studious"},
		{"the", "student", "is", "smart"},
	}, []int{1, 2}, 2))
	// [2, 1]

	fmt.Println(topStudents([]string{"smart"}, []string{"boring"}, [][]string{
		{"this", "is", "smart"},
		{"this", "is", "boring"},
	}, []int{1, 2}, 1))
	// [1]
}

func topStudents(positiveFeedback []string, negativeFeedback []string, report [][]string, studentID []int, k int) []int {
	pos := make(map[string]bool)
	neg := make(map[string]bool)
	for _, w := range positiveFeedback {
		pos[w] = true
	}
	for _, w := range negativeFeedback {
		neg[w] = true
	}

	type student struct {
		id, score int
	}
	students := make([]student, len(studentID))

	for i, id := range studentID {
		score := 0
		for _, w := range report[i] {
			w = strings.ToLower(w)
			if pos[w] {
				score += 3
			} else if neg[w] {
				score -= 1
			}
		}
		students[i] = student{id, score}
	}

	sort.Slice(students, func(i, j int) bool {
		if students[i].score != students[j].score {
			return students[i].score > students[j].score
		}
		return students[i].id < students[j].id
	})

	ans := make([]int, k)
	for i := 0; i < k; i++ {
		ans[i] = students[i].id
	}
	return ans
}
```

## 2513 — Minimize The Maximum Of Two Arrays

```go
package main

// LeetCode #2513: Minimize the Maximum of Two Arrays
// https://leetcode.com/problems/minimize-the-maximum-of-two-arrays/
// Difficulty: Medium
// Time: O(log(max)) | Space: O(1)
// Binary search on answer. For value X:
//   count1 = X - X/divisor1 (numbers not divisible by divisor1)
//   count2 = X - X/divisor2 (numbers not divisible by divisor2)
//   common = X - X/lcm (numbers not divisible by either)
//   We need count1 >= uniqueCnt1, count2 >= uniqueCnt2, common >= uniqueCnt1 + uniqueCnt2

import "fmt"

func main() {
	fmt.Println(minimizeSet(2, 7, 1, 3)) // 4
	fmt.Println(minimizeSet(3, 5, 2, 1)) // 3
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

func minimizeSet(divisor1 int, divisor2 int, uniqueCnt1 int, uniqueCnt2 int) int {
	lo, hi := 1, 1<<31-1
	l := lcm(divisor1, divisor2)

	for lo < hi {
		mid := (lo + hi) / 2
		cnt1 := mid - mid/divisor1
		cnt2 := mid - mid/divisor2
		common := mid - mid/l

		if cnt1 >= uniqueCnt1 && cnt2 >= uniqueCnt2 && common >= uniqueCnt1+uniqueCnt2 {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}
```

## 2516 — Take K Of Each Character From Left And Right

```go
package main

// LeetCode #2516: Take K of Each Character From Left and Right
// https://leetcode.com/problems/take-k-of-each-character-from-left-and-right/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// Sliding window: find max middle substring that doesn't exceed total-k per char.
// Answer = n - len(max window).

import "fmt"

func main() {
	fmt.Println(takeCharacters("aabaaaacaabc", 2)) // 8
	fmt.Println(takeCharacters("a", 1))            // 1
}

func takeCharacters(s string, k int) int {
	n := len(s)
	if k == 0 {
		return 0
	}

	// Total counts
	total := make([]int, 3)
	for _, ch := range s {
		total[ch-'a']++
	}
	for _, c := range total {
		if c < k {
			return -1
		}
	}

	// Max middle window where each char <= total-chars - k
	need := []int{total[0] - k, total[1] - k, total[2] - k}
	cnt := make([]int, 3)
	left, maxWindow := 0, 0

	for right := 0; right < n; right++ {
		cnt[s[right]-'a']++
		for cnt[0] > need[0] || cnt[1] > need[1] || cnt[2] > need[2] {
			cnt[s[left]-'a']--
			left++
		}
		if right-left+1 > maxWindow {
			maxWindow = right - left + 1
		}
	}
	return n - maxWindow
}
```

## 2517 — Maximum Tastiness Of Candy Basket

```go
package main

// LeetCode #2517: Maximum Tastiness of Candy Basket
// https://leetcode.com/problems/maximum-tastiness-of-candy-basket/
// Difficulty: Medium
// Time: O(n log n + n log max) | Space: O(1)
// Binary search on answer. Check: can we pick k candies with min diff >= mid?

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximumTastiness([]int{13, 5, 1, 8, 21, 2}, 3)) // 8
	fmt.Println(maximumTastiness([]int{1, 3, 1}, 2))             // 2
}

func maximumTastiness(price []int, k int) int {
	sort.Ints(price)
	lo, hi := 0, price[len(price)-1]-price[0]

	for lo < hi {
		mid := (lo + hi + 1) / 2
		if canPick(price, k, mid) {
			lo = mid
		} else {
			hi = mid - 1
		}
	}
	return lo
}

func canPick(price []int, k, minDiff int) bool {
	count, last := 1, price[0]
	for i := 1; i < len(price); i++ {
		if price[i]-last >= minDiff {
			count++
			last = price[i]
		}
	}
	return count >= k
}
```

## 2521 — Distinct Prime Factors Of Product Of Array

```go
package main

// LeetCode #2521: Distinct Prime Factors of Product of Array
// https://leetcode.com/problems/distinct-prime-factors-of-product-of-array/
// Difficulty: Medium
// Time: O(n * sqrt(max)) | Space: O(number of primes)
// Find union of prime factors of all numbers.

import "fmt"

func main() {
	fmt.Println(distinctPrimeFactors([]int{2, 4, 3, 7, 10, 6})) // 4 (2, 3, 5, 7)
	fmt.Println(distinctPrimeFactors([]int{4, 8, 16}))           // 1 (2)
}

func distinctPrimeFactors(nums []int) int {
	primes := make(map[int]bool)
	for _, v := range nums {
		addPrimeFactors(v, primes)
	}
	return len(primes)
}

func addPrimeFactors(n int, set map[int]bool) {
	// Factor 2
	if n%2 == 0 {
		set[2] = true
		for n%2 == 0 {
			n /= 2
		}
	}
	// Odd factors
	for f := 3; f*f <= n; f += 2 {
		if n%f == 0 {
			set[f] = true
			for n%f == 0 {
				n /= f
			}
		}
	}
	if n > 1 {
		set[n] = true
	}
}
```

## 2522 — Partition String Into Substrings With Values At Most K

```go
package main

// LeetCode #2522: Partition String Into Substrings With Values at Most K
// https://leetcode.com/problems/partition-string-into-substrings-with-values-at-most-k/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// Greedy: extend substring while value <= k, then start new partition.

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minimumPartition("165462", 60)) // 4 (16|54|6|2)
	fmt.Println(minimumPartition("238182", 5))  // -1
}

func minimumPartition(s string, k int) int {
	ans := 1
	cur := 0
	for _, ch := range s {
		d := int(ch - '0')
		if d > k {
			return -1
		}
		if cur > math.MaxInt32/10 || cur*10+d > k {
			ans++
			cur = d
		} else {
			cur = cur*10 + d
		}
	}
	return ans
}
```

## 2523 — Closest Prime Numbers In Range

```go
package main

// LeetCode #2523: Closest Prime Numbers in Range
// https://leetcode.com/problems/closest-prime-numbers-in-range/
// Difficulty: Medium
// Time: O(right log log right) | Space: O(right)
// Sieve of Eratosthenes, find adjacent primes with min difference.

import "fmt"

func main() {
	fmt.Println(closestPrimes(10, 19)) // [11, 13]
	fmt.Println(closestPrimes(4, 6))   // [-1, -1]
}

func closestPrimes(left int, right int) []int {
	if right < 2 {
		return []int{-1, -1}
	}

	isPrime := make([]bool, right+1)
	for i := 2; i <= right; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i <= right; i++ {
		if isPrime[i] {
			for j := i * i; j <= right; j += i {
				isPrime[j] = false
			}
		}
	}

	prev := -1
	minDiff := right + 1
	ans := []int{-1, -1}

	for i := left; i <= right; i++ {
		if isPrime[i] {
			if prev != -1 && i-prev < minDiff {
				minDiff = i - prev
				ans = []int{prev, i}
			}
			prev = i
		}
	}
	return ans
}
```

## 2526 — Find Consecutive Integers From A Data Stream

```go
package main

// LeetCode #2526: Find Consecutive Integers from a Data Stream
// https://leetcode.com/problems/find-consecutive-integers-from-a-data-stream/
// Difficulty: Medium
// Time: O(1) per call | Space: O(1)
// Track count of consecutive `value` seen.

import "fmt"

type DataStream struct {
	value, k, count int
}

func main() {
	ds := Constructor(4, 3)
	fmt.Println(ds.Consec(4)) // false
	fmt.Println(ds.Consec(4)) // false
	fmt.Println(ds.Consec(4)) // true
	fmt.Println(ds.Consec(3)) // false
}

func Constructor(value int, k int) DataStream {
	return DataStream{value: value, k: k}
}

func (ds *DataStream) Consec(num int) bool {
	if num == ds.value {
		ds.count++
	} else {
		ds.count = 0
	}
	return ds.count >= ds.k
}
```

## 2527 — Find Xor Beauty Of Array

```go
package main

// LeetCode #2527: Find Xor-Beauty of Array
// https://leetcode.com/problems/find-xor-beauty-of-array/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// XOR of all (nums[i]|nums[j]) & nums[k] over all i,j,k = XOR of all nums[i].
// Because the expression simplifies: each bit appears in result iff it appears odd times in nums.

import "fmt"

func main() {
	fmt.Println(xorBeauty([]int{1, 4})) // 5
	fmt.Println(xorBeauty([]int{15, 45, 20, 2, 34, 35, 5, 44, 32, 30})) // 34
}

func xorBeauty(nums []int) int {
	ans := 0
	for _, v := range nums {
		ans ^= v
	}
	return ans
}
```

## 2530 — Maximal Score After Applying K Operations

```go
package main

// LeetCode #2530: Maximal Score After Applying K Operations
// https://leetcode.com/problems/maximal-score-after-applying-k-operations/
// Difficulty: Medium
// Time: O((n + k) log n) | Space: O(n)
// Max heap: each operation take max, add ceil(v/3), put back.

import (
	"container/heap"
	"fmt"
)

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() any {
	old := *h; n := len(old); x := old[n-1]; *h = old[:n-1]; return x
}

func main() {
	fmt.Println(maxKelements([]int{10, 10, 10, 10, 10}, 5)) // 50
	fmt.Println(maxKelements([]int{1, 10, 3, 3, 3}, 3))     // 17
}

func maxKelements(nums []int, k int) int64 {
	h := &MaxHeap{}
	heap.Init(h)
	for _, v := range nums {
		heap.Push(h, v)
	}
	var score int64
	for i := 0; i < k; i++ {
		v := heap.Pop(h).(int)
		score += int64(v)
		heap.Push(h, (v+2)/3) // ceil(v/3)
	}
	return score
}
```

## 2531 — Make Number Of Distinct Characters Equal

```go
package main

// LeetCode #2531: Make Number of Distinct Characters Equal
// https://leetcode.com/problems/make-number-of-distinct-characters-equal/
// Difficulty: Medium
// Time: O(26^2) | Space: O(26)
// Try swapping one char from word1 with one char from word2.
// Check if distinct counts become equal.

import "fmt"

func main() {
	fmt.Println(isItPossible("ac", "b"))    // false
	fmt.Println(isItPossible("abcc", "aab")) // true
}

func isItPossible(word1 string, word2 string) bool {
	c1, c2 := make([]int, 26), make([]int, 26)
	for _, ch := range word1 {
		c1[ch-'a']++
	}
	for _, ch := range word2 {
		c2[ch-'a']++
	}

	for i := 0; i < 26; i++ {
		if c1[i] == 0 {
			continue
		}
		for j := 0; j < 26; j++ {
			if c2[j] == 0 {
				continue
			}
			// swap char i (from word1) with char j (from word2)
			// decrement c1[i], c2[j]; increment c1[j], c2[i]
			c1[i]--
			c2[j]--
			c1[j]++
			c2[i]++

			d1, d2 := 0, 0
			for k := 0; k < 26; k++ {
				if c1[k] > 0 {
					d1++
				}
				if c2[k] > 0 {
					d2++
				}
			}

			if d1 == d2 {
				return true
			}

			// revert
			c1[i]++
			c2[j]++
			c1[j]--
			c2[i]--
		}
	}
	return false
}
```

## 2533 — Number Of Good Binary Strings

```go
package main

// LeetCode #2533: Number of Good Binary Strings
// https://leetcode.com/problems/number-of-good-binary-strings/
// Difficulty: Medium
// Time: O(maxLen) | Space: O(maxLen)
// DP: dp[i] = number of good strings of length i.
// Recurrence: dp[i] = dp[i-oneGroup] + dp[i-zeroGroup]
// Each maximal block of 1s must have length multiple of oneGroup.
// Each maximal block of 0s must have length multiple of zeroGroup.

import "fmt"

func main() {
	fmt.Println(goodBinaryStrings(2, 3, 1, 2)) // 5
	fmt.Println(goodBinaryStrings(3, 3, 1, 1)) // 8
}

const MOD = 1000000007

func goodBinaryStrings(minLength int, maxLength int, oneGroup int, zeroGroup int) int {
	dp := make([]int, maxLength+1)
	dp[0] = 1

	for i := 1; i <= maxLength; i++ {
		if i >= oneGroup {
			dp[i] = (dp[i] + dp[i-oneGroup]) % MOD
		}
		if i >= zeroGroup {
			dp[i] = (dp[i] + dp[i-zeroGroup]) % MOD
		}
	}

	ans := 0
	for i := minLength; i <= maxLength; i++ {
		ans = (ans + dp[i]) % MOD
	}
	return ans
}
```

## 2536 — Increment Submatrices By One

```go
package main

// LeetCode #2536: Increment Submatrices by One
// https://leetcode.com/problems/increment-submatrices-by-one/
// Difficulty: Medium
// Time: O(n^2 + q) | Space: O(n^2)

import "fmt"

func rangeAddQueries(n int, queries [][]int) [][]int {
	diff := make([][]int, n+1)
	for i := range diff {
		diff[i] = make([]int, n+1)
	}

	for _, q := range queries {
		r1, c1, r2, c2 := q[0], q[1], q[2], q[3]
		diff[r1][c1]++
		diff[r1][c2+1]--
		diff[r2+1][c1]--
		diff[r2+1][c2+1]++
	}

	mat := make([][]int, n)
	for i := range mat {
		mat[i] = make([]int, n)
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
			mat[i][j] = diff[i][j]
		}
	}
	return mat
}

func main() {
	// Test case 1: n=2, queries=[[0,0,0,0]]
	res1 := rangeAddQueries(2, [][]int{{0, 0, 0, 0}})
	fmt.Println("Test 1:", res1)
	// Expected: [[1,0],[0,0]]

	// Test case 2: n=2, queries=[[0,0,1,1]]
	res2 := rangeAddQueries(2, [][]int{{0, 0, 1, 1}})
	fmt.Println("Test 2:", res2)
	// Expected: [[1,1],[1,1]]

	// Test case 3: n=1
	res3 := rangeAddQueries(1, [][]int{{0, 0, 0, 0}})
	fmt.Println("Test 3:", res3)
	// Expected: [[1]]
}
```

## 2537 — Count The Number Of Good Subarrays

```go
package main

// LeetCode #2537: Count the Number of Good Subarrays
// https://leetcode.com/problems/count-the-number-of-good-subarrays/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func countGood(nums []int, k int) int64 {
	n := len(nums)
	freq := make(map[int]int)
	var pairs int64
	var ans int64
	left := 0

	for right := 0; right < n; right++ {
		val := nums[right]
		pairs += int64(freq[val])
		freq[val]++

		for pairs >= int64(k) {
			ans += int64(n - right)
			leftVal := nums[left]
			freq[leftVal]--
			pairs -= int64(freq[leftVal])
			left++
		}
	}
	return ans
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", countGood([]int{1, 1, 1, 1, 1}, 10))
	// Expected: 1

	// Test case 2
	fmt.Println("Test 2:", countGood([]int{3, 1, 4, 3, 2, 2, 4}, 2))
	// Expected: 4

	// Test case 3
	fmt.Println("Test 3:", countGood([]int{1, 2, 3}, 1))
	// Expected: 0
}
```

## 2539 — Count The Number Of Good Subsequences

```go
package main

// LeetCode #2539: Count the Number of Good Subsequences
// https://leetcode.com/problems/count-the-number-of-good-subsequences/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func countGoodSubsequences(s string) int {
	const mod = 1_000_000_007
	freq := make(map[rune]int)
	for _, ch := range s {
		freq[ch]++
	}

	maxFreq := 0
	for _, f := range freq {
		if f > maxFreq {
			maxFreq = f
		}
	}

	// Precompute factorials and inverse factorials
	maxN := maxFreq
	fact := make([]int, maxN+1)
	fact[0] = 1
	for i := 1; i <= maxN; i++ {
		fact[i] = fact[i-1] * i % mod
	}

	invFact := make([]int, maxN+1)
	invFact[maxN] = powMod(fact[maxN], mod-2, mod)
	for i := maxN; i > 0; i-- {
		invFact[i-1] = invFact[i] * i % mod
	}

	nCr := func(n, r int) int {
		if r < 0 || r > n {
			return 0
		}
		return fact[n] * invFact[r] % mod * invFact[n-r] % mod
	}

	var ans int
	for maxLen := 1; maxLen <= maxFreq; maxLen++ {
		ways := 1
		for _, f := range freq {
			if f >= maxLen {
				ways = ways * (nCr(f, maxLen) + 1) % mod
			}
		}
		ans = (ans + ways - 1 + mod) % mod
	}

	return ans
}

func powMod(a, b, mod int) int {
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
	// Test case 1
	fmt.Println("Test 1:", countGoodSubsequences("aabb"))
	// Expected: 11

	// Test case 2
	fmt.Println("Test 2:", countGoodSubsequences("leet"))
	// Expected: 12

	// Test case 3
	fmt.Println("Test 3:", countGoodSubsequences("abcd"))
	// Expected: 4
}
```

## 2541 — Minimum Operations To Make Array Equal Ii

```go
package main

// LeetCode #2541: Minimum Operations to Make Array Equal II
// https://leetcode.com/problems/minimum-operations-to-make-array-equal-ii/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func minOperations(nums1 []int, nums2 []int, k int) int64 {
	n := len(nums1)
	if k == 0 {
		for i := 0; i < n; i++ {
			if nums1[i] != nums2[i] {
				return -1
			}
		}
		return 0
	}

	var posDiff, negDiff int64
	for i := 0; i < n; i++ {
		diff := nums1[i] - nums2[i]
		if diff%k != 0 {
			return -1
		}
		if diff > 0 {
			posDiff += int64(diff)
		} else if diff < 0 {
			negDiff += int64(-diff)
		}
	}

	if posDiff != negDiff {
		return -1
	}
	return posDiff / int64(k)
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minOperations([]int{4, 3, 1, 4}, []int{1, 3, 7, 1}, 3))
	// Expected: 2

	// Test case 2: valid case
	fmt.Println("Test 2:", minOperations([]int{5, 10, 3}, []int{8, 6, 4}, 1))
	// Expected: 4

	// Test case 3: impossible (sums don't match or diff not divisible)
	fmt.Println("Test 3:", minOperations([]int{1, 2}, []int{2, 1}, 3))
	// Expected: -1
}
```

## 2542 — Maximum Subsequence Score

```go
package main

// LeetCode #2542: Maximum Subsequence Score
// https://leetcode.com/problems/maximum-subsequence-score/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"container/heap"
	"fmt"
	"sort"
)

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

func maxScore(nums1 []int, nums2 []int, k int) int64 {
	n := len(nums1)
	pairs := make([][2]int, n)
	for i := 0; i < n; i++ {
		pairs[i] = [2]int{nums2[i], nums1[i]}
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] > pairs[j][0]
	})

	h := &MinHeap{}
	heap.Init(h)
	var sum int64
	var ans int64

	for _, p := range pairs {
		val1, val2 := p[1], p[0]
		sum += int64(val1)
		heap.Push(h, val1)
		if h.Len() > k {
			sum -= int64(heap.Pop(h).(int))
		}
		if h.Len() == k {
			score := sum * int64(val2)
			if score > ans {
				ans = score
			}
		}
	}
	return ans
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", maxScore([]int{1, 3, 3, 2}, []int{2, 1, 3, 4}, 3))
	// Expected: 12

	// Test case 2
	fmt.Println("Test 2:", maxScore([]int{4, 2, 3, 1, 1}, []int{7, 5, 10, 9, 6}, 1))
	// Expected: 30

	// Test case 3
	fmt.Println("Test 3:", maxScore([]int{2, 1, 14, 12}, []int{11, 7, 13, 6}, 3))
	// Expected: 168
}
```

## 2545 — Sort The Students By Their Kth Score

```go
package main

// LeetCode #2545: Sort the Students by Their Kth Score
// https://leetcode.com/problems/sort-the-students-by-their-kth-score/
// Difficulty: Medium
// Time: O(m log m) | Space: O(1) (excluding output)

import (
	"fmt"
	"sort"
)

func sortTheStudents(score [][]int, k int) [][]int {
	sort.Slice(score, func(i, j int) bool {
		return score[i][k] > score[j][k]
	})
	return score
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", sortTheStudents([][]int{{10, 6, 9, 1}, {7, 5, 11, 2}, {4, 8, 3, 15}}, 2))
	// Expected: [[7,5,11,2],[10,6,9,1],[4,8,3,15]]

	// Test case 2
	fmt.Println("Test 2:", sortTheStudents([][]int{{3, 4}, {5, 6}}, 0))
	// Expected: [[5,6],[3,4]]

	// Test case 3: single student
	fmt.Println("Test 3:", sortTheStudents([][]int{{1, 2, 3}}, 1))
	// Expected: [[1,2,3]]
}
```

## 2546 — Apply Bitwise Operations To Make Strings Equal

```go
package main

// LeetCode #2546: Apply Bitwise Operations to Make Strings Equal
// https://leetcode.com/problems/apply-bitwise-operations-to-make-strings-equal/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func makeStringsEqual(s string, target string) bool {
	// Operation: choose i,j, set s[i]=s[i]|s[j], s[j]=s[i]^s[j]
	// We can change any position if there's at least one '1' in either string
	hasOneS := false
	hasOneT := false
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			hasOneS = true
		}
		if target[i] == '1' {
			hasOneT = true
		}
	}
	// If target has no '1', s must have no '1' too
	if !hasOneT {
		return !hasOneS
	}
	// If target has '1', s must have at least one '1' to make it work
	return hasOneS
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", makeStringsEqual("1010", "0101"))
	// Expected: true

	// Test case 2
	fmt.Println("Test 2:", makeStringsEqual("00", "11"))
	// Expected: false

	// Test case 3
	fmt.Println("Test 3:", makeStringsEqual("11", "00"))
	// Expected: false
}
```

## 2548 — Maximum Price To Fill A Bag

```go
package main

// LeetCode #2548: Maximum Price to Fill a Bag
// https://leetcode.com/problems/maximum-price-to-fill-a-bag/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func maxPrice(items [][]int, capacity int) float64 {
	// items[i] = [price, weight]
	// Sort by price/weight ratio descending
	sort.Slice(items, func(i, j int) bool {
		return float64(items[i][0])/float64(items[i][1]) > float64(items[j][0])/float64(items[j][1])
	})

	var totalPrice float64
	remaining := capacity

	for _, item := range items {
		if remaining <= 0 {
			break
		}
		price, weight := item[0], item[1]
		take := weight
		if take > remaining {
			take = remaining
		}
		totalPrice += float64(price) * float64(take) / float64(weight)
		remaining -= take
	}

	if remaining > 0 {
		return -1
	}
	return totalPrice
}

func main() {
	// Test case 1
	fmt.Printf("Test 1: %.5f\n", maxPrice([][]int{{50, 10}, {100, 20}, {120, 30}}, 50))
	// Expected: 240.00000

	// Test case 2: can't fill
	fmt.Printf("Test 2: %.5f\n", maxPrice([][]int{{50, 10}}, 20))
	// Expected: -1

	// Test case 3: exact fit
	fmt.Printf("Test 3: %.5f\n", maxPrice([][]int{{60, 10}, {100, 20}}, 30))
	// Expected: 160.00000
}
```

## 2550 — Count Collisions Of Monkeys On A Polygon

```go
package main

// LeetCode #2550: Count Collisions of Monkeys on a Polygon
// https://leetcode.com/problems/count-collisions-of-monkeys-on-a-polygon/
// Difficulty: Medium
// Time: O(log n) | Space: O(1)

import "fmt"

func monkeyMove(n int) int {
	// Total ways: 2^n. Only 2 ways avoid collisions: all clockwise or all anticlockwise
	// Answer: (2^n - 2 + mod) % mod
	const mod = 1_000_000_007

	pow2 := 1
	base := 2
	exp := n
	for exp > 0 {
		if exp&1 == 1 {
			pow2 = pow2 * base % mod
		}
		base = base * base % mod
		exp >>= 1
	}

	ans := (pow2 - 2 + mod) % mod
	return ans
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", monkeyMove(3))
	// Expected: 6

	// Test case 2
	fmt.Println("Test 2:", monkeyMove(4))
	// Expected: 14

	// Test case 3
	fmt.Println("Test 3:", monkeyMove(100))
}
```

## 2554 — Maximum Number Of Integers To Choose From A Range I

```go
package main

// LeetCode #2554: Maximum Number of Integers to Choose From a Range I
// https://leetcode.com/problems/maximum-number-of-integers-to-choose-from-a-range-i/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func maxCount(banned []int, n int, maxSum int) int {
	bannedSet := make(map[int]bool)
	for _, b := range banned {
		bannedSet[b] = true
	}

	count := 0
	sum := 0
	for i := 1; i <= n; i++ {
		if bannedSet[i] {
			continue
		}
		if sum+i > maxSum {
			break
		}
		sum += i
		count++
	}
	return count
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", maxCount([]int{1, 6, 5}, 5, 6))
	// Expected: 2 (choose 2,3)

	// Test case 2
	fmt.Println("Test 2:", maxCount([]int{1, 2, 3, 4, 5, 6, 7}, 8, 1))
	// Expected: 0

	// Test case 3
	fmt.Println("Test 3:", maxCount([]int{11}, 7, 50))
	// Expected: 7 (choose 1..7)
}
```

## 2555 — Maximize Win From Two Segments

```go
package main

// LeetCode #2555: Maximize Win From Two Segments
// https://leetcode.com/problems/maximize-win-from-two-segments/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func maximizeWin(prizePositions []int, k int) int {
	n := len(prizePositions)
	// dp[i] = max prizes we can win with one segment ending at or before position i
	dp := make([]int, n+1)
	ans := 0

	left := 0
	for right := 0; right < n; right++ {
		for prizePositions[right]-prizePositions[left] > k {
			left++
		}
		// Current segment [left, right] covers `right-left+1` prizes
		curr := right - left + 1
		// Best with one segment up to position before left
		dp[right+1] = max(dp[right], curr)
		// Combine with best segment before current segment
		ans = max(ans, curr+dp[left])
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", maximizeWin([]int{1, 1, 2, 2, 3, 3, 5}, 2))
	// Expected: 7

	// Test case 2
	fmt.Println("Test 2:", maximizeWin([]int{1, 2, 3, 4}, 0))
	// Expected: 2

	// Test case 3
	fmt.Println("Test 3:", maximizeWin([]int{1, 2, 3, 4, 5, 6}, 2))
	// Expected: 4
}
```

## 2556 — Disconnect Path In A Binary Matrix By At Most One Flip

```go
package main

// LeetCode #2556: Disconnect Path in a Binary Matrix by at Most One Flip
// https://leetcode.com/problems/disconnect-path-in-a-binary-matrix-by-at-most-one-flip/
// Difficulty: Medium
// Time: O(m*n) | Space: O(m*n)

import "fmt"

func isPossibleToCutPath(grid [][]int) bool {
	m, n := len(grid), len(grid[0])

	// First DFS from (0,0) to (m-1,n-1), mark visited cells
	var dfs1 func(r, c int) bool
	dfs1 = func(r, c int) bool {
		if r >= m || c >= n || grid[r][c] == 0 {
			return false
		}
		if r == m-1 && c == n-1 {
			return true
		}
		grid[r][c] = 0 // Mark as visited
		return dfs1(r+1, c) || dfs1(r, c+1)
	}

	if !dfs1(0, 0) {
		return true // Already disconnected
	}

	// Second DFS from (0,0) checking if there's still a path after first removal
	var dfs2 func(r, c int) bool
	dfs2 = func(r, c int) bool {
		if r >= m || c >= n || grid[r][c] == 0 {
			return false
		}
		if r == m-1 && c == n-1 {
			return true
		}
		grid[r][c] = 0
		return dfs2(r, c+1) || dfs2(r+1, c)
	}

	if !dfs2(0, 0) {
		return true
	}
	return false
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", isPossibleToCutPath([][]int{{1, 1, 1}, {1, 0, 0}, {1, 1, 1}}))
	// Expected: true

	// Test case 2
	fmt.Println("Test 2:", isPossibleToCutPath([][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}))
	// Expected: false

	// Test case 3: single cell
	fmt.Println("Test 3:", isPossibleToCutPath([][]int{{1}}))
	// Expected: false
}
```

## 2557 — Maximum Number Of Integers To Choose From A Range Ii

```go
package main

// LeetCode #2557: Maximum Number of Integers to Choose From a Range II
// https://leetcode.com/problems/maximum-number-of-integers-to-choose-from-a-range-ii/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func maxCount(banned []int, n int, maxSum int) int {
	sort.Ints(banned)
	bannedSet := make(map[int]bool)
	for _, b := range banned {
		bannedSet[b] = true
	}

	count := 0
	sum := 0
	for i := 1; i <= n; i++ {
		if bannedSet[i] {
			continue
		}
		if sum+i > maxSum {
			break
		}
		sum += i
		count++
	}
	return count
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", maxCount([]int{1, 6, 5}, 5, 6))
	// Expected: 2

	// Test case 2: large banned set
	fmt.Println("Test 2:", maxCount([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 15, 40))
	// Expected: 3 (pick 11,12,13)

	// Test case 3
	fmt.Println("Test 3:", maxCount([]int{}, 5, 15))
	// Expected: 5 (1+2+3+4+5=15)
}
```

## 2559 — Count Vowel Strings In Ranges

```go
package main

// LeetCode #2559: Count Vowel Strings in Ranges
// https://leetcode.com/problems/count-vowel-strings-in-ranges/
// Difficulty: Medium
// Time: O(n + q) | Space: O(n)

import "fmt"

func vowelStrings(words []string, queries [][]int) []int {
	isVowel := func(ch byte) bool {
		return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u'
	}

	n := len(words)
	prefix := make([]int, n+1)
	for i, w := range words {
		prefix[i+1] = prefix[i]
		if len(w) > 0 && isVowel(w[0]) && isVowel(w[len(w)-1]) {
			prefix[i+1]++
		}
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		ans[i] = prefix[q[1]+1] - prefix[q[0]]
	}
	return ans
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", vowelStrings([]string{"aba", "bcb", "ece", "aa", "e"}, [][]int{{0, 2}, {1, 4}, {1, 1}}))
	// Expected: [2,3,0]

	// Test case 2
	fmt.Println("Test 2:", vowelStrings([]string{"a", "e", "i"}, [][]int{{0, 2}, {0, 0}, {2, 2}}))
	// Expected: [3,1,1]

	// Test case 3: empty range
	fmt.Println("Test 3:", vowelStrings([]string{"a", "b", "c"}, [][]int{{1, 1}}))
	// Expected: [0]
}
```

## 2560 — House Robber Iv

```go
package main

// LeetCode #2560: House Robber IV
// https://leetcode.com/problems/house-robber-iv/
// Difficulty: Medium
// Time: O(n log max) | Space: O(1)

import "fmt"

func minCapability(nums []int, k int) int {
	canRob := func(cap int) bool {
		count := 0
		i := 0
		for i < len(nums) {
			if nums[i] <= cap {
				count++
				i += 2 // Skip adjacent house
			} else {
				i++
			}
		}
		return count >= k
	}

	left, right := nums[0], nums[0]
	for _, v := range nums {
		if v < left {
			left = v
		}
		if v > right {
			right = v
		}
	}

	for left < right {
		mid := left + (right-left)/2
		if canRob(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minCapability([]int{2, 3, 5, 9}, 2))
	// Expected: 5

	// Test case 2
	fmt.Println("Test 2:", minCapability([]int{2, 7, 9, 3, 1}, 2))
	// Expected: 2

	// Test case 3
	fmt.Println("Test 3:", minCapability([]int{1, 2, 3, 4, 5, 6, 7}, 3))
	// Expected: 5
}
```

## 2563 — Count The Number Of Fair Pairs

```go
package main

// LeetCode #2563: Count the Number of Fair Pairs
// https://leetcode.com/problems/count-the-number-of-fair-pairs/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func countFairPairs(nums []int, lower int, upper int) int64 {
	sort.Ints(nums)
	var ans int64
	n := len(nums)

	for i := 0; i < n; i++ {
		// Find lower bound for j > i such that nums[j] >= lower - nums[i]
		left := sort.Search(n, func(j int) bool {
			return j > i && nums[j] >= lower-nums[i]
		})
		// Find upper bound for j > i such that nums[j] <= upper - nums[i]
		right := sort.Search(n, func(j int) bool {
			return j > i && nums[j] > upper-nums[i]
		})
		if right > left {
			ans += int64(right - left)
		}
	}
	return ans
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", countFairPairs([]int{0, 1, 7, 4, 4, 5}, 3, 6))
	// Expected: 6

	// Test case 2
	fmt.Println("Test 2:", countFairPairs([]int{1, 7, 9, 2, 5}, 11, 11))
	// Expected: 1

	// Test case 3
	fmt.Println("Test 3:", countFairPairs([]int{0, 0, 0, 0, 0, 0}, 0, 0))
	// Expected: 15
}
```

## 2564 — Substring Xor Queries

```go
package main

// LeetCode #2564: Substring XOR Queries
// https://leetcode.com/problems/substring-xor-queries/
// Difficulty: Medium
// Time: O(n * 31 + q) | Space: O(n * 31)

import "fmt"

func substringXorQueries(s string, queries [][]int) [][]int {
	// For each possible value, store earliest [l, r]
	n := len(s)
	posMap := make(map[int][2]int)

	// For each starting position, compute values up to 31 bits (since val <= 10^9 < 2^30)
	for i := 0; i < n; i++ {
		if s[i] == '0' {
			if _, ok := posMap[0]; !ok {
				posMap[0] = [2]int{i, i}
			}
			continue
		}
		val := 0
		for j := i; j < n && j-i < 31; j++ {
			val = (val << 1) | int(s[j]-'0')
			if _, ok := posMap[val]; !ok {
				posMap[val] = [2]int{i, j}
			}
		}
	}

	ans := make([][]int, len(queries))
	for idx, q := range queries {
		first, second := q[0], q[1]
		target := first ^ second
		if pos, ok := posMap[target]; ok {
			ans[idx] = []int{pos[0], pos[1]}
		} else {
			ans[idx] = []int{-1, -1}
		}
	}
	return ans
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", substringXorQueries("101101", [][]int{{0, 5}, {1, 2}}))
	// Expected: [[0,2],[2,3]]

	// Test case 2
	fmt.Println("Test 2:", substringXorQueries("0101", [][]int{{12, 8}}))
	// value=12^8=4 (100), need substring "100"

	// Test case 3
	fmt.Println("Test 3:", substringXorQueries("1", [][]int{{0, 0}}))
	// value=0^0=0, need substring "0"
}
```

## 2567 — Minimum Score By Changing Two Elements

```go
package main

// LeetCode #2567: Minimum Score by Changing Two Elements
// https://leetcode.com/problems/minimum-score-by-changing-two-elements/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func minimizeSum(nums []int) int {
	n := len(nums)
	if n <= 3 {
		return 0
	}
	sort.Ints(nums)

	// After changing two elements, the min sum possible is:
	// Option 1: change two smallest to = third smallest, score = nums[n-1] - nums[2]
	// Option 2: change two largest to = third largest, score = nums[n-3] - nums[0]
	// Option 3: change one smallest and one largest, score = nums[n-2] - nums[1]

	diff1 := nums[n-1] - nums[2]
	diff2 := nums[n-3] - nums[0]
	diff3 := nums[n-2] - nums[1]

	result := diff1
	if diff2 < result {
		result = diff2
	}
	if diff3 < result {
		result = diff3
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minimizeSum([]int{1, 4, 3}))
	// Expected: 0

	// Test case 2
	fmt.Println("Test 2:", minimizeSum([]int{1, 4, 7, 8, 5}))
	// Expected: 3

	// Test case 3
	fmt.Println("Test 3:", minimizeSum([]int{1, 2, 3, 4, 5}))
	// Expected: 2
}
```

## 2568 — Minimum Impossible Or

```go
package main

// LeetCode #2568: Minimum Impossible OR
// https://leetcode.com/problems/minimum-impossible-or/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func minImpossibleOR(nums []int) int {
	seen := make(map[int]bool)
	for _, v := range nums {
		seen[v] = true
	}

	// Check powers of 2: 1, 2, 4, 8, ...
	// If any is missing, that's the answer (since it can't be formed by OR of smaller numbers)
	pow2 := 1
	for {
		if !seen[pow2] {
			return pow2
		}
		pow2 <<= 1
	}
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minImpossibleOR([]int{2, 1}))
	// Expected: 4

	// Test case 2
	fmt.Println("Test 2:", minImpossibleOR([]int{5, 3, 2}))
	// Expected: 1

	// Test case 3
	fmt.Println("Test 3:", minImpossibleOR([]int{1, 2, 4, 8}))
	// Expected: 16
}
```

## 2571 — Minimum Operations To Reduce An Integer To 0

```go
package main

// LeetCode #2571: Minimum Operations to Reduce an Integer to 0
// https://leetcode.com/problems/minimum-operations-to-reduce-an-integer-to-0/
// Difficulty: Medium
// Time: O(log n) | Space: O(1)

import (
	"fmt"
	"math"
)

func minOperations(n int) int {
	const INF = math.MaxInt32
	dp := [32][2]int{}
	for i := range dp {
		dp[i] = [2]int{INF, INF}
	}
	dp[0][0] = 0

	for i := 0; i < 31; i++ {
		bit := (n >> i) & 1
		for carry := 0; carry < 2; carry++ {
			if dp[i][carry] == INF {
				continue
			}
			val := bit + carry
			switch val {
			case 0:
				if dp[i][carry] < dp[i+1][0] {
					dp[i+1][0] = dp[i][carry]
				}
			case 1:
				if dp[i][carry]+1 < dp[i+1][0] {
					dp[i+1][0] = dp[i][carry] + 1
				}
				if dp[i][carry]+1 < dp[i+1][1] {
					dp[i+1][1] = dp[i][carry] + 1
				}
			case 2:
				if dp[i][carry] < dp[i+1][1] {
					dp[i+1][1] = dp[i][carry]
				}
			}
		}
	}

	ans := dp[31][0]
	if dp[31][1]+1 < ans {
		ans = dp[31][1] + 1
	}
	return ans
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minOperations(3))
	// Expected: 2

	// Test case 2
	fmt.Println("Test 2:", minOperations(6))
	// Expected: 2

	// Test case 3
	fmt.Println("Test 3:", minOperations(7))
	// Expected: 2
}
```

## 2572 — Count The Number Of Square Free Subsets

```go
package main

// LeetCode #2572: Count the Number of Square-Free Subsets
// https://leetcode.com/problems/count-the-number-of-square-free-subsets/
// Difficulty: Medium
// Time: O(n * 2^p) | Space: O(2^p)

import "fmt"

func squareFreeSubsets(nums []int) int {
	const mod = 1_000_000_007

	// Primes up to 30 (since nums[i] <= 30)
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	p := len(primes)

	// Map each number to its prime mask
	primeMask := make([]int, 31)
	for i := 1; i <= 30; i++ {
		mask := 0
		x := i
		for j, prime := range primes {
			cnt := 0
			for x%prime == 0 {
				x /= prime
				cnt++
			}
			if cnt > 1 {
				mask = -1 // Not square-free
				break
			}
			if cnt == 1 {
				mask |= (1 << j)
			}
		}
		primeMask[i] = mask
	}

	// Count frequency of each number
	freq := make([]int, 31)
	for _, v := range nums {
		freq[v]++
	}

	// DP: dp[mask] = number of ways to get this mask
	dp := make([]int, 1<<p)
	dp[0] = 1

	for val := 1; val <= 30; val++ {
		if freq[val] == 0 {
			continue
		}
		mask := primeMask[val]
		if mask < 0 {
			continue
		}

		// We have freq[val] copies of this value
		// Each copy can be either taken or not
		// Ways to take any subset of freq[val] copies = 2^freq[val] - 1
		pow := 1
		for i := 0; i < freq[val]; i++ {
			pow = (pow * 2) % mod
		}
		ways := (pow - 1 + mod) % mod

		// Update DP (knapsack style)
		newDP := make([]int, 1<<p)
		copy(newDP, dp)
		for m := 0; m < (1 << p); m++ {
			if dp[m] == 0 {
				continue
			}
			newMask := m | mask
			newDP[newMask] = (newDP[newMask] + dp[m]*ways) % mod
		}
		dp = newDP
	}

	// Sum all non-empty subsets: any mask (0 included for subsets only containing 1)
	ans := (dp[0] - 1 + mod) % mod // non-empty subsets of number 1 only
	for m := 1; m < (1 << p); m++ {
		ans = (ans + dp[m]) % mod
	}
	return ans
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", squareFreeSubsets([]int{3, 4, 4, 5}))
	// Expected: 3

	// Test case 2
	fmt.Println("Test 2:", squareFreeSubsets([]int{1}))
	// Expected: 1

	// Test case 3
	fmt.Println("Test 3:", squareFreeSubsets([]int{1, 2, 3, 4}))
	// Expected: 7
}
```

## 2575 — Find The Divisibility Array Of A String

```go
package main

// LeetCode #2575: Find the Divisibility Array of a String
// https://leetcode.com/problems/find-the-divisibility-array-of-a-string/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func divisibilityArray(word string, m int) []int {
	n := len(word)
	ans := make([]int, n)
	rem := 0
	for i := 0; i < n; i++ {
		rem = (rem*10 + int(word[i]-'0')) % m
		if rem == 0 {
			ans[i] = 1
		} else {
			ans[i] = 0
		}
	}
	return ans
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", divisibilityArray("998244353", 3))
	// Expected: [1,1,0,0,0,1,1,0,0]

	// Test case 2
	fmt.Println("Test 2:", divisibilityArray("1010", 10))
	// Expected: [0,1,0,1]

	// Test case 3
	fmt.Println("Test 3:", divisibilityArray("10", 5))
	// Expected: [0,0]
}
```

## 2576 — Find The Maximum Number Of Marked Indices

```go
package main

// LeetCode #2576: Find the Maximum Number of Marked Indices
// https://leetcode.com/problems/find-the-maximum-number-of-marked-indices/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func maxNumOfMarkedIndices(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	left, right := 0, n/2
	count := 0

	for left < n/2 && right < n {
		if 2*nums[left] <= nums[right] {
			count += 2
			left++
			right++
		} else {
			right++
		}
	}
	return count
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", maxNumOfMarkedIndices([]int{3, 5, 2, 4}))
	// Expected: 2

	// Test case 2
	fmt.Println("Test 2:", maxNumOfMarkedIndices([]int{9, 2, 5, 4}))
	// Expected: 4

	// Test case 3
	fmt.Println("Test 3:", maxNumOfMarkedIndices([]int{7, 6, 8}))
	// Expected: 0
}
```

## 2579 — Count Total Number Of Colored Cells

```go
package main

// LeetCode #2579: Count Total Number of Colored Cells
// https://leetcode.com/problems/count-total-number-of-colored-cells/
// Difficulty: Medium
// Time: O(1) | Space: O(1)

import "fmt"

func coloredCells(n int) int64 {
	// Formula: 1 + 2*n*(n-1) = 2n^2 - 2n + 1
	return int64(1 + 2*n*(n-1))
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", coloredCells(1))
	// Expected: 1

	// Test case 2
	fmt.Println("Test 2:", coloredCells(2))
	// Expected: 5

	// Test case 3
	fmt.Println("Test 3:", coloredCells(3))
	// Expected: 13
}
```

## 2580 — Count Ways To Group Overlapping Ranges

```go
package main

// LeetCode #2580: Count Ways to Group Overlapping Ranges
// https://leetcode.com/problems/count-ways-to-group-overlapping-ranges/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func countWays(ranges [][]int) int {
	const mod = 1_000_000_007

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	groups := 0
	end := -1
	for _, r := range ranges {
		if r[0] > end {
			groups++
		}
		if r[1] > end {
			end = r[1]
		}
	}

	// Each group can be in group 1 or group 2, so 2^groups ways
	ans := 1
	for i := 0; i < groups; i++ {
		ans = (ans * 2) % mod
	}
	return ans
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", countWays([][]int{{6, 10}, {5, 15}}))
	// Expected: 2

	// Test case 2
	fmt.Println("Test 2:", countWays([][]int{{1, 3}, {10, 20}, {2, 5}, {4, 8}}))
	// Expected: 4

	// Test case 3
	fmt.Println("Test 3:", countWays([][]int{{1, 2}, {3, 4}}))
	// Expected: 4
}
```

## 2583 — Kth Largest Sum In A Binary Tree

```go
package main

// LeetCode #2583: Kth Largest Sum in a Binary Tree
// https://leetcode.com/problems/kth-largest-sum-in-a-binary-tree/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"container/heap"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type MinHeap []int64

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(int64)) }
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	if root == nil {
		return -1
	}

	queue := []*TreeNode{root}
	h := &MinHeap{}
	heap.Init(h)

	for len(queue) > 0 {
		size := len(queue)
		var sum int64
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			sum += int64(node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		heap.Push(h, sum)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	if h.Len() < k {
		return -1
	}
	return (*h)[0]
}

func main() {
	// Test case 1: [5,8,9,2,1,3,7,4,6]
	root1 := &TreeNode{Val: 5}
	root1.Left = &TreeNode{Val: 8}
	root1.Right = &TreeNode{Val: 9}
	root1.Left.Left = &TreeNode{Val: 2}
	root1.Left.Right = &TreeNode{Val: 1}
	root1.Right.Left = &TreeNode{Val: 3}
	root1.Right.Right = &TreeNode{Val: 7}
	root1.Left.Left.Left = &TreeNode{Val: 4}
	root1.Left.Left.Right = &TreeNode{Val: 6}
	fmt.Println("Test 1:", kthLargestLevelSum(root1, 2))

	// Test case 2: single node
	root2 := &TreeNode{Val: 1}
	fmt.Println("Test 2:", kthLargestLevelSum(root2, 1))

	// Test case 3: k larger than levels
	root3 := &TreeNode{Val: 1}
	root3.Left = &TreeNode{Val: 2}
	fmt.Println("Test 3:", kthLargestLevelSum(root3, 3))
}
```

## 2587 — Rearrange Array To Maximize Prefix Score

```go
package main

// LeetCode #2587: Rearrange Array to Maximize Prefix Score
// https://leetcode.com/problems/rearrange-array-to-maximize-prefix-score/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func maxScore(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	var sum int64
	count := 0
	for _, v := range nums {
		sum += int64(v)
		if sum > 0 {
			count++
		} else {
			break
		}
	}
	return count
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", maxScore([]int{2, -1, 0, 1, -3, 3, -3}))
	// Expected: 6

	// Test case 2
	fmt.Println("Test 2:", maxScore([]int{-2, -3, 0}))
	// Expected: 0

	// Test case 3
	fmt.Println("Test 3:", maxScore([]int{1, 2, 3, 4, 5}))
	// Expected: 5
}
```

## 2588 — Count The Number Of Beautiful Subarrays

```go
package main

// LeetCode #2588: Count the Number of Beautiful Subarrays
// https://leetcode.com/problems/count-the-number-of-beautiful-subarrays/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func beautifulSubarrays(nums []int) int64 {
	prefixXor := make(map[int]int)
	prefixXor[0] = 1
	xor := 0
	var ans int64

	for _, v := range nums {
		xor ^= v
		ans += int64(prefixXor[xor])
		prefixXor[xor]++
	}
	return ans
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", beautifulSubarrays([]int{4, 3, 1, 2, 4}))
	// Expected: 2

	// Test case 2
	fmt.Println("Test 2:", beautifulSubarrays([]int{1, 10, 4}))
	// Expected: 0

	// Test case 3
	fmt.Println("Test 3:", beautifulSubarrays([]int{0, 0, 0}))
	// Expected: 6
}
```

## 2590 — Design A Todo List

```go
package main

// LeetCode #2590: Design a Todo List
// https://leetcode.com/problems/design-a-todo-list/
// Difficulty: Medium [Paid]
// Time: O(n log n) per getUserTasks | Space: O(n)

import (
	"fmt"
	"sort"
)

type Task struct {
	ID        int
	Desc      string
	DueDate   int
	Tags      []string
	Completed bool
	UserID    int
}

type TodoList struct {
	tasks   []*Task
	nextID  int
}

func Constructor() TodoList {
	return TodoList{
		tasks:  make([]*Task, 0),
		nextID: 1,
	}
}

func (this *TodoList) AddTask(userId int, taskDescription string, taskDueDate int, taskTags []string) int {
	id := this.nextID
	this.nextID++
	task := &Task{
		ID:        id,
		Desc:      taskDescription,
		DueDate:   taskDueDate,
		Tags:      taskTags,
		Completed: false,
		UserID:    userId,
	}
	this.tasks = append(this.tasks, task)
	return id
}

func (this *TodoList) GetAllTasks(userId int) []string {
	userTasks := []*Task{}
	for _, t := range this.tasks {
		if t.UserID == userId && !t.Completed {
			userTasks = append(userTasks, t)
		}
	}
	sort.Slice(userTasks, func(i, j int) bool {
		if userTasks[i].DueDate != userTasks[j].DueDate {
			return userTasks[i].DueDate < userTasks[j].DueDate
		}
		return userTasks[i].ID < userTasks[j].ID
	})
	res := make([]string, len(userTasks))
	for i, t := range userTasks {
		res[i] = t.Desc
	}
	return res
}

func (this *TodoList) GetTasksForTag(userId int, tag string) []string {
	userTasks := []*Task{}
	for _, t := range this.tasks {
		if t.UserID == userId && !t.Completed {
			for _, tg := range t.Tags {
				if tg == tag {
					userTasks = append(userTasks, t)
					break
				}
			}
		}
	}
	sort.Slice(userTasks, func(i, j int) bool {
		if userTasks[i].DueDate != userTasks[j].DueDate {
			return userTasks[i].DueDate < userTasks[j].DueDate
		}
		return userTasks[i].ID < userTasks[j].ID
	})
	res := make([]string, len(userTasks))
	for i, t := range userTasks {
		res[i] = t.Desc
	}
	return res
}

func (this *TodoList) CompleteTask(userId int, taskId int) {
	for _, t := range this.tasks {
		if t.ID == taskId && t.UserID == userId {
			t.Completed = true
			break
		}
	}
}

func main() {
	// Test case
	todo := Constructor()
	id1 := todo.AddTask(1, "Task1", 50, []string{})
	id2 := todo.AddTask(1, "Task2", 30, []string{"tag1"})
	todo.AddTask(2, "Task3", 10, []string{})
	fmt.Println("Test 1:", todo.GetAllTasks(1))
	// Expected: ["Task2", "Task1"]

	todo.CompleteTask(1, id2)
	fmt.Println("Test 2:", todo.GetAllTasks(1))
	// Expected: ["Task1"]

	todo.CompleteTask(1, id1)
	fmt.Println("Test 3:", todo.GetAllTasks(1))
	// Expected: []
}
```

## 2592 — Maximize Greatness Of An Array

```go
package main

// LeetCode #2592: Maximize Greatness of an Array
// https://leetcode.com/problems/maximize-greatness-of-an-array/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func maximizeGreatness(nums []int) int {
	sort.Ints(nums)
	j := 0
	for _, v := range nums {
		if v > nums[j] {
			j++
		}
	}
	return j
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", maximizeGreatness([]int{1, 3, 5, 2, 1, 3, 1}))
	// Expected: 4

	// Test case 2
	fmt.Println("Test 2:", maximizeGreatness([]int{1, 2, 3, 4}))
	// Expected: 3

	// Test case 3
	fmt.Println("Test 3:", maximizeGreatness([]int{1, 1, 1}))
	// Expected: 0
}
```

## 2593 — Find Score Of An Array After Marking All Elements

```go
package main

// LeetCode #2593: Find Score of an Array After Marking All Elements
// https://leetcode.com/problems/find-score-of-an-array-after-marking-all-elements/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"container/heap"
	"fmt"
)

type Item struct {
	val int
	idx int
}

type MinHeap []Item

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].val < h[j].val || (h[i].val == h[j].val && h[i].idx < h[j].idx) }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(Item)) }
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func findScore(nums []int) int64 {
	n := len(nums)
	marked := make([]bool, n)
	h := &MinHeap{}
	heap.Init(h)

	for i, v := range nums {
		heap.Push(h, Item{v, i})
	}

	var score int64
	for h.Len() > 0 {
		item := heap.Pop(h).(Item)
		val, idx := item.val, item.idx
		if marked[idx] {
			continue
		}
		score += int64(val)
		marked[idx] = true
		if idx-1 >= 0 {
			marked[idx-1] = true
		}
		if idx+1 < n {
			marked[idx+1] = true
		}
	}
	return score
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", findScore([]int{2, 1, 3, 4, 5, 2}))
	// Expected: 7

	// Test case 2
	fmt.Println("Test 2:", findScore([]int{2, 3, 5, 1, 3, 2}))
	// Expected: 5

	// Test case 3
	fmt.Println("Test 3:", findScore([]int{5, 4, 3, 2, 1}))
	// Expected: 9
}
```

## 2594 — Minimum Time To Repair Cars

```go
package main

// LeetCode #2594: Minimum Time to Repair Cars
// https://leetcode.com/problems/minimum-time-to-repair-cars/
// Difficulty: Medium
// Time: O(n log minTime) | Space: O(1)

import (
	"fmt"
	"math"
)

func repairCars(ranks []int, cars int) int64 {
	canRepair := func(t int64) bool {
		var total int64
		for _, r := range ranks {
			total += int64(math.Sqrt(float64(t / int64(r))))
			if total >= int64(cars) {
				return true
			}
		}
		return false
	}

	left := int64(1)
	right := int64(ranks[0]) * int64(cars) * int64(cars)
	for i := range ranks {
		candidate := int64(ranks[i]) * int64(cars) * int64(cars)
		if candidate < right {
			right = candidate
		}
	}

	for left < right {
		mid := left + (right-left)/2
		if canRepair(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", repairCars([]int{4, 2, 3, 1}, 10))
	// Expected: 16

	// Test case 2
	fmt.Println("Test 2:", repairCars([]int{5, 1, 8}, 6))
	// Expected: 16

	// Test case 3
	fmt.Println("Test 3:", repairCars([]int{1, 1, 1}, 5))
	// Expected: 3
}
```

## 2596 — Check Knight Tour Configuration

```go
package main

// LeetCode #2596: Check Knight Tour Configuration
// https://leetcode.com/problems/check-knight-tour-configuration/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n^2)

import "fmt"

func checkValidGrid(grid [][]int) bool {
	n := len(grid)
	if n == 0 {
		return false
	}
	if grid[0][0] != 0 {
		return false
	}

	// Find positions of each move
	pos := make([][2]int, n*n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			pos[grid[i][j]] = [2]int{i, j}
		}
	}

	// Knight moves
	dirs := [8][2]int{{-2, -1}, {-2, 1}, {-1, -2}, {-1, 2}, {1, -2}, {1, 2}, {2, -1}, {2, 1}}

	for k := 1; k < n*n; k++ {
		r, c := pos[k][0], pos[k][1]
		pr, pc := pos[k-1][0], pos[k-1][1]
		valid := false
		for _, d := range dirs {
			if pr+d[0] == r && pc+d[1] == c {
				valid = true
				break
			}
		}
		if !valid {
			return false
		}
	}
	return true
}

func main() {
	// Test case 1: valid
	fmt.Println("Test 1:", checkValidGrid([][]int{{0, 11, 16, 5, 20}, {17, 4, 19, 10, 15}, {12, 1, 8, 21, 6}, {3, 18, 23, 14, 9}, {24, 13, 2, 7, 22}}))
	// Expected: true

	// Test case 2: invalid (grid[0][0] != 0)
	fmt.Println("Test 2:", checkValidGrid([][]int{{1, 0}, {0, 1}}))
	// Expected: false

	// Test case 3: 1x1 grid
	fmt.Println("Test 3:", checkValidGrid([][]int{{0}}))
	// Expected: true
}
```

## 2597 — The Number Of Beautiful Subsets

```go
package main

// LeetCode #2597: The Number of Beautiful Subsets
// https://leetcode.com/problems/the-number-of-beautiful-subsets/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func beautifulSubsets(nums []int, k int) int {
	sort.Ints(nums)
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}

	// Group by residue modulo k
	grouped := make(map[int][]int)
	for v := range freq {
		grouped[v%k] = append(grouped[v%k], v)
	}

	var ans int = 1 // empty subset
	for _, vals := range grouped {
		sort.Ints(vals)
		// DP within each group: dp0 (ways not taking current), dp1 (ways taking current)
		dp0, dp1 := 1, 0 // dp0 = ways for "not taking previous value", dp1 = ways for "taking previous value"
		for i, v := range vals {
			ways0 := dp0 + dp1 // skip current value
			ways1 := dp0 * (1 << uint(freq[v]-1))
			if i > 0 && v-vals[i-1] == k {
				ways1 = dp0 * ((1 << uint(freq[v])) - 1)
			} else {
				ways1 = (dp0 + dp1) * ((1 << uint(freq[v])) - 1)
			}
			dp0, dp1 = ways0, ways1
		}
		ans *= (dp0 + dp1)
	}

	return ans - 1 // exclude empty subset
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", beautifulSubsets([]int{2, 4, 6}, 2))
	// Expected: 4

	// Test case 2
	fmt.Println("Test 2:", beautifulSubsets([]int{1}, 1))
	// Expected: 1

	// Test case 3
	fmt.Println("Test 3:", beautifulSubsets([]int{1, 2, 3, 4}, 1))
	// Expected: 7
}
```

## 2598 — Smallest Missing Non Negative Integer After Operations

```go
package main

// LeetCode #2598: Smallest Missing Non-negative Integer After Operations
// https://leetcode.com/problems/smallest-missing-non-negative-integer-after-operations/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func findSmallestInteger(nums []int, value int) int {
	freq := make([]int, value)
	for _, v := range nums {
		// Map to [0, value-1] range
		rem := ((v % value) + value) % value
		freq[rem]++
	}

	for i := 0; ; i++ {
		if freq[i%value] == 0 {
			return i
		}
		freq[i%value]--
	}
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", findSmallestInteger([]int{1, -10, 7, 13, 6, 8}, 5))
	// Expected: 4

	// Test case 2
	fmt.Println("Test 2:", findSmallestInteger([]int{1, 2, 3, 4, 5}, 2))
	// Expected: 4

	// Test case 3
	fmt.Println("Test 3:", findSmallestInteger([]int{0, 0, 0, 0}, 1))
	// Expected: 4
}
```

## 2599 — Make The Prefix Sum Non Negative

```go
package main

// LeetCode #2599: Make the Prefix Sum Non-negative
// https://leetcode.com/problems/make-the-prefix-sum-non-negative/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"container/heap"
	"fmt"
)

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

func makePrefSumNonNegative(nums []int) int {
	h := &MinHeap{}
	heap.Init(h)
	prefix := int64(0)
	moves := 0

	for _, v := range nums {
		prefix += int64(v)
		if v < 0 {
			heap.Push(h, v)
		}
		for prefix < 0 {
			// Move the most negative seen to the end
			smallest := heap.Pop(h).(int)
			prefix -= int64(smallest)
			moves++
		}
	}
	return moves
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", makePrefSumNonNegative([]int{2, 3, -5, 4}))
	// Expected: 0

	// Test case 2
	fmt.Println("Test 2:", makePrefSumNonNegative([]int{3, -5, -2, 6}))
	// Expected: 1

	// Test case 3
	fmt.Println("Test 3:", makePrefSumNonNegative([]int{1, -2, 3, -4, 5}))
	// Expected: 1
}
```

## 2601 — Prime Subtraction Operation

```go
package main

// LeetCode #2601: Prime Subtraction Operation
// https://leetcode.com/problems/prime-subtraction-operation/
// Difficulty: Medium
// Time: O(n * maxVal) | Space: O(maxVal)

import "fmt"

func primeSubOperation(nums []int) bool {
	// Sieve: find all primes up to 1000
	maxN := 1000
	isPrime := make([]bool, maxN+1)
	for i := 2; i <= maxN; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i <= maxN; i++ {
		if isPrime[i] {
			for j := i * i; j <= maxN; j += i {
				isPrime[j] = false
			}
		}
	}
	primes := []int{}
	for i := 2; i <= maxN; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}

	prev := 0
	for _, v := range nums {
		// Find largest prime p such that v-p > prev
		chosen := v
		for _, p := range primes {
			if p >= v {
				break
			}
			if v-p > prev {
				chosen = v - p
			} else {
				break
			}
		}
		if chosen <= prev {
			return false
		}
		prev = chosen
	}
	return true
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", primeSubOperation([]int{4, 9, 6, 10}))
	// Expected: true

	// Test case 2
	fmt.Println("Test 2:", primeSubOperation([]int{6, 8, 11, 12}))
	// Expected: true

	// Test case 3
	fmt.Println("Test 3:", primeSubOperation([]int{5, 8, 3}))
	// Expected: false
}
```

## 2602 — Minimum Operations To Make All Array Elements Equal

```go
package main

// LeetCode #2602: Minimum Operations to Make All Array Elements Equal
// https://leetcode.com/problems/minimum-operations-to-make-all-array-elements-equal/
// Difficulty: Medium
// Time: O((n+q) log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func minOperations(nums []int, queries []int) []int64 {
	n := len(nums)
	sort.Ints(nums)

	prefix := make([]int64, n+1)
	for i, v := range nums {
		prefix[i+1] = prefix[i] + int64(v)
	}

	ans := make([]int64, len(queries))
	for i, q := range queries {
		idx := sort.SearchInts(nums, q)
		leftCount := int64(idx)
		rightCount := int64(n - idx)
		leftSum := prefix[idx]
		rightSum := prefix[n] - prefix[idx]
		ops := int64(q)*leftCount - leftSum + rightSum - int64(q)*rightCount
		ans[i] = ops
	}
	return ans
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minOperations([]int{3, 1, 6, 8}, []int{1, 5}))
	// Expected: [8, 10]? Let me check...

	// Test case 2
	fmt.Println("Test 2:", minOperations([]int{2, 4, 6, 8}, []int{4, 5}))
	// Expected: [4, 4]

	// Test case 3
	fmt.Println("Test 3:", minOperations([]int{1, 2, 3, 4, 5}, []int{3}))
	// Expected: [6]
}
```

## 2606 — Find The Substring With Maximum Cost

```go
package main

// LeetCode #2606: Find the Substring With Maximum Cost
// https://leetcode.com/problems/find-the-substring-with-maximum-cost/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func maximumCostSubstring(s string, chars string, vals []int) int {
	// Build cost map
	cost := make([]int, 26)
	for i := 0; i < 26; i++ {
		cost[i] = i + 1 // 'a' = 1, 'b' = 2, ...
	}
	for i, c := range chars {
		cost[c-'a'] = vals[i]
	}

	// Kadane's algorithm
	maxEnd := 0
	maxSoFar := 0
	for _, ch := range s {
		val := cost[ch-'a']
		maxEnd = maxEnd + val
		if maxEnd < 0 {
			maxEnd = 0
		}
		if maxEnd > maxSoFar {
			maxSoFar = maxEnd
		}
	}
	return maxSoFar
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", maximumCostSubstring("adaa", "d", []int{-1000}))
	// Expected: 2 (choose "aa" for cost 2)

	// Test case 2
	fmt.Println("Test 2:", maximumCostSubstring("abc", "abc", []int{-1, -1, -1}))
	// Expected: 0 (empty substring)

	// Test case 3
	fmt.Println("Test 3:", maximumCostSubstring("aabc", "ab", []int{5, 5}))
	// Expected: 10 (choose "aa" with custom values)
}
```

## 2607 — Make K Subarray Sums Equal

```go
package main

// LeetCode #2607: Make K-Subarray Sums Equal
// https://leetcode.com/problems/make-k-subarray-sums-equal/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func makeSubKSumEqual(arr []int, k int) int64 {
	n := len(arr)
	g := gcd(n, k)
	var ans int64

	for i := 0; i < g; i++ {
		group := []int{}
		for j := i; j < n; j += g {
			group = append(group, arr[j])
		}
		sort.Ints(group)
		median := group[len(group)/2]
		for _, v := range group {
			diff := v - median
			if diff < 0 {
				diff = -diff
			}
			ans += int64(diff)
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

func main() {
	// Test case 1
	fmt.Println("Test 1:", makeSubKSumEqual([]int{1, 4, 1, 3}, 2))
	// Expected: 1

	// Test case 2
	fmt.Println("Test 2:", makeSubKSumEqual([]int{2, 5, 5, 7}, 3))
	// Expected: 5

	// Test case 3
	fmt.Println("Test 3:", makeSubKSumEqual([]int{1, 2, 3, 4, 5, 6}, 3))
	// Expected: 4
}
```

## 2610 — Convert An Array Into A 2d Array With Conditions

```go
package main

// LeetCode #2610: Convert an Array Into a 2D Array With Conditions
// https://leetcode.com/problems/convert-an-array-into-a-2d-array-with-conditions/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func findMatrix(nums []int) [][]int {
	freq := make(map[int]int)
	ans := [][]int{}

	for _, v := range nums {
		freq[v]++
		if freq[v] > len(ans) {
			ans = append(ans, []int{})
		}
		ans[freq[v]-1] = append(ans[freq[v]-1], v)
	}
	return ans
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", findMatrix([]int{1, 3, 4, 1, 2, 3, 1}))
	// Expected: [[1,3,4,2],[1,3],[1]]

	// Test case 2
	fmt.Println("Test 2:", findMatrix([]int{1, 2, 3, 4}))
	// Expected: [[1,2,3,4]]

	// Test case 3
	fmt.Println("Test 3:", findMatrix([]int{1, 1, 1, 1}))
	// Expected: [[1],[1],[1],[1]]
}
```

## 2611 — Mice And Cheese

```go
package main

// LeetCode #2611: Mice and Cheese
// https://leetcode.com/problems/mice-and-cheese/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func miceAndCheese(reward1 []int, reward2 []int, k int) int {
	n := len(reward1)
	diff := make([]int, n)
	for i := 0; i < n; i++ {
		diff[i] = reward1[i] - reward2[i]
	}

	// We need to pick k indices for mouse 1 to maximize total
	// total = sum(reward2) + sum of top k diffs
	total := 0
	for _, v := range reward2 {
		total += v
	}

	sort.Slice(diff, func(i, j int) bool {
		return diff[i] > diff[j]
	})

	for i := 0; i < k; i++ {
		total += diff[i]
	}
	return total
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", miceAndCheese([]int{1, 1, 3, 4}, []int{4, 4, 1, 1}, 2))
	// Expected: 15

	// Test case 2
	fmt.Println("Test 2:", miceAndCheese([]int{1, 1}, []int{1, 1}, 2))
	// Expected: 2

	// Test case 3
	fmt.Println("Test 3:", miceAndCheese([]int{2, 3, 5}, []int{5, 2, 1}, 1))
	// Expected: 12
}
```

## 2615 — Sum Of Distances

```go
package main

// LeetCode #2615: Sum of Distances
// https://leetcode.com/problems/sum-of-distances/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func distance(nums []int) []int64 {
	n := len(nums)
	// Group indices by value
	groups := make(map[int][]int)
	for i, v := range nums {
		groups[v] = append(groups[v], i)
	}

	ans := make([]int64, n)
	for _, indices := range groups {
		m := len(indices)
		prefix := make([]int64, m+1)
		for i, idx := range indices {
			prefix[i+1] = prefix[i] + int64(idx)
		}
		for i, idx := range indices {
			leftCount := int64(i)
			rightCount := int64(m - i - 1)
			leftSum := prefix[i]
			rightSum := prefix[m] - prefix[i+1]
			ans[idx] = int64(idx)*leftCount - leftSum + rightSum - int64(idx)*rightCount
		}
	}
	return ans
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", distance([]int{1, 3, 1, 1, 2}))
	// Expected: [5,0,3,4,0]

	// Test case 2
	fmt.Println("Test 2:", distance([]int{0, 5, 3}))
	// Expected: [0,0,0]

	// Test case 3
	fmt.Println("Test 3:", distance([]int{1, 1, 1, 1}))
	// Expected: [6,4,4,6]
}
```

## 2616 — Minimize The Maximum Difference Of Pairs

```go
package main

// LeetCode #2616: Minimize the Maximum Difference of Pairs
// https://leetcode.com/problems/minimize-the-maximum-difference-of-pairs/
// Difficulty: Medium
// Time: O(n log n + n log maxDiff) | Space: O(1)

import (
	"fmt"
	"sort"
)

func minimizeMax(nums []int, p int) int {
	if p == 0 {
		return 0
	}
	sort.Ints(nums)
	n := len(nums)

	canForm := func(mid int) bool {
		count := 0
		i := 0
		for i < n-1 {
			if nums[i+1]-nums[i] <= mid {
				count++
				i += 2
			} else {
				i++
			}
			if count >= p {
				return true
			}
		}
		return false
	}

	left, right := 0, nums[n-1]-nums[0]
	for left < right {
		mid := left + (right-left)/2
		if canForm(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minimizeMax([]int{10, 1, 2, 7, 1, 3}, 2))
	// Expected: 1

	// Test case 2
	fmt.Println("Test 2:", minimizeMax([]int{4, 2, 1, 2}, 1))
	// Expected: 0

	// Test case 3
	fmt.Println("Test 3:", minimizeMax([]int{3, 5, 2, 8, 1, 9}, 3))
	// Expected: 2
}
```

## 2618 — Check If Object Instance Of Class

```go
package main

// LeetCode #2618: Check if Object Instance of Class
// https://leetcode.com/problems/check-if-object-instance-of-class/
// Difficulty: Medium
// Time: O(d) where d is depth of type hierarchy | Space: O(1)

import (
	"fmt"
	"reflect"
)

func checkIfInstanceOf(obj any, classType reflect.Type) bool {
	if obj == nil || classType == nil {
		return false
	}

	t := reflect.TypeOf(obj)
	for t != nil {
		if t == classType {
			return true
		}
		if classType.Kind() == reflect.Interface && t.Implements(classType) {
			return true
		}
		// Move to pointer type and check
		ptrT := reflect.PointerTo(t)
		if ptrT == classType {
			return true
		}
		// Move to underlying type
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		} else {
			break
		}
	}
	return false
}

type MyClass struct{}

func main() {
	// Test case 1
	obj := &MyClass{}
	fmt.Println("Test 1:", checkIfInstanceOf(obj, reflect.TypeOf(MyClass{})))
	// Expected: true

	// Test case 2: int is instance of int
	fmt.Println("Test 2:", checkIfInstanceOf(5, reflect.TypeOf(0)))
	// Expected: true

	// Test case 3
	fmt.Println("Test 3:", checkIfInstanceOf("hello", reflect.TypeOf(0)))
	// Expected: false
}
```

## 2622 — Cache With Time Limit

```go
package main

// LeetCode #2622: Cache With Time Limit
// https://leetcode.com/problems/cache-with-time-limit/
// Difficulty: Medium
// Time: O(1) per operation average | Space: O(n)

import (
	"fmt"
	"time"
)

type TimedCache struct {
	cache    map[int]timedValue
}

type timedValue struct {
	value     int
	expiresAt time.Time
}

func NewTimedCache() TimedCache {
	return TimedCache{cache: make(map[int]timedValue)}
}

func (tc *TimedCache) Set(key int, value int, duration time.Duration) bool {
	now := time.Now()
	if existing, ok := tc.cache[key]; ok {
		if existing.expiresAt.After(now) {
			tc.cache[key] = timedValue{value, now.Add(duration)}
			return false // Key exists and hasn't expired
		}
	}
	tc.cache[key] = timedValue{value, now.Add(duration)}
	return true // Key was inserted (didn't exist or was expired)
}

func (tc *TimedCache) Get(key int) (int, bool) {
	now := time.Now()
	if existing, ok := tc.cache[key]; ok {
		if existing.expiresAt.After(now) {
			return existing.value, true
		}
		delete(tc.cache, key)
	}
	return 0, false
}

func (tc *TimedCache) Count() int {
	now := time.Now()
	count := 0
	for key, tv := range tc.cache {
		if tv.expiresAt.After(now) {
			count++
		} else {
			delete(tc.cache, key)
		}
	}
	return count
}

func main() {
	cache := NewTimedCache()

	// Test case 1: Set and Get
	cache.Set(1, 100, 500*time.Millisecond)
	val, ok := cache.Get(1)
	fmt.Println("Test 1:", val, ok)
	// Expected: 100 true

	// Test case 2: Count
	fmt.Println("Test 2:", cache.Count())
	// Expected: 1

	// Test case 3: Overwrite existing
	cache.Set(1, 200, 500*time.Millisecond)
	val, ok = cache.Get(1)
	fmt.Println("Test 3:", val, ok)
	// Expected: 200 true
}
```

## 2623 — Memoize

```go
package main

// LeetCode #2623: Memoize
// https://leetcode.com/problems/memoize/
// Difficulty: Medium
// Time: O(1) amortized per call | Space: O(n)

import (
	"fmt"
	"sync"
)

type MemoizedFn func(args ...int) int

func memoize(fn func(...int) int) MemoizedFn {
	var mu sync.Mutex
	cache := make(map[string]int)

	return func(args ...int) int {
		// Create a key from args
		key := ""
		for i, arg := range args {
			if i > 0 {
				key += ","
			}
			key += fmt.Sprintf("%d", arg)
		}

		mu.Lock()
		defer mu.Unlock()

		if val, ok := cache[key]; ok {
			return val
		}
		result := fn(args...)
		cache[key] = result
		return result
	}
}

var callCount int

func add(a, b int) int {
	callCount++
	return a + b
}

func main() {
	callCount = 0
	memoizedAdd := memoize(func(args ...int) int {
		return add(args[0], args[1])
	})

	// Test case 1
	fmt.Println("Test 1:", memoizedAdd(1, 2))
	// Expected: 3

	// Test case 2: same args, should use cache
	fmt.Println("Test 2:", memoizedAdd(1, 2))
	// Expected: 3

	// Test case 3: different args
	fmt.Println("Test 3:", memoizedAdd(2, 3))
	// Expected: 5

	fmt.Println("Calls:", callCount)
	// Expected: 2 (not 3)
}
```

## 2624 — Snail Traversal

```go
package main

// LeetCode #2624: Snail Traversal
// https://leetcode.com/problems/snail-traversal/
// Difficulty: Medium
// Time: O(n*m) | Space: O(n*m)

import "fmt"

func snailTraversal(arr []int, rowsCount int, colsCount int) [][]int {
	n := len(arr)
	if rowsCount*colsCount != n {
		return [][]int{}
	}

	result := make([][]int, rowsCount)
	for i := range result {
		result[i] = make([]int, colsCount)
	}

	idx := 0
	for col := 0; col < colsCount; col++ {
		if col%2 == 0 {
			// Top to bottom
			for row := 0; row < rowsCount; row++ {
				result[row][col] = arr[idx]
				idx++
			}
		} else {
			// Bottom to top
			for row := rowsCount - 1; row >= 0; row-- {
				result[row][col] = arr[idx]
				idx++
			}
		}
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", snailTraversal([]int{1, 2, 3, 4}, 2, 2))
	// Expected: [[1,4],[2,3]] (snail fill)

	// Test case 2
	fmt.Println("Test 2:", snailTraversal([]int{1, 2, 3, 4, 5, 6}, 2, 3))
	// Expected: [[1,4,5],[2,3,6]]

	// Test case 3: invalid dimensions
	fmt.Println("Test 3:", snailTraversal([]int{1, 2, 3}, 2, 2))
	// Expected: []
}
```

## 2625 — Flatten Deeply Nested Array

```go
package main

// LeetCode #2625: Flatten Deeply Nested Array
// https://leetcode.com/problems/flatten-deeply-nested-array/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func flatten(arr []any, depth int) []any {
	result := []any{}

	var dfs func(item any, currDepth int)
	dfs = func(item any, currDepth int) {
		switch v := item.(type) {
		case []any:
			if currDepth < depth {
				for _, sub := range v {
					dfs(sub, currDepth+1)
				}
			} else {
				result = append(result, v)
			}
		default:
			result = append(result, v)
		}
	}

	for _, item := range arr {
		dfs(item, 0)
	}
	return result
}

func main() {
	// Test case 1: flatten to depth 1
	arr1 := []any{1, []any{2, []any{3, 4}}, 5}
	fmt.Println("Test 1:", flatten(arr1, 1))
	// Expected: [1, 2, [3, 4], 5]

	// Test case 2: flatten to depth 2
	fmt.Println("Test 2:", flatten(arr1, 2))
	// Expected: [1, 2, 3, 4, 5]

	// Test case 3: depth 0 - no flattening
	fmt.Println("Test 3:", flatten(arr1, 0))
	// Expected: [1, [2, [3, 4]], 5]
}
```

## 2627 — Debounce

```go
package main

// LeetCode #2627: Debounce
// https://leetcode.com/problems/debounce/
// Difficulty: Medium
// Time: O(1) per call | Space: O(1)

import (
	"fmt"
	"sync"
	"time"
)

type DebouncedFn struct {
	mu       sync.Mutex
	timer    *time.Timer
	duration time.Duration
	fn       func(...int)
}

func (d *DebouncedFn) Call(args ...int) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.timer != nil {
		d.timer.Stop()
	}
	d.timer = time.AfterFunc(d.duration, func() {
		d.fn(args...)
	})
}

func (d *DebouncedFn) Cancel() {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.timer != nil {
		d.timer.Stop()
	}
}

func newDebounce(fn func(...int), duration time.Duration) *DebouncedFn {
	return &DebouncedFn{
		fn:       fn,
		duration: duration,
	}
}

func main() {
	calls := []int{}
	debounced := newDebounce(func(args ...int) {
		calls = append(calls, args[0])
	}, 50*time.Millisecond)

	debounced.Call(1)
	debounced.Call(2)
	debounced.Call(3)

	time.Sleep(100 * time.Millisecond)
	fmt.Println("Test 1:", calls)
	// Expected: [3] (only last call goes through)

	debounced.Call(4)
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Test 2:", calls)
	// Expected: [3, 4]

	debounced.Cancel()
	fmt.Println("Test 3: cancelled without call")
}
```

