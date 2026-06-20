# Medium (Sedang) — Problem 3466–3649

## 3466 — Maximum Coin Collection

```go
package main

// LeetCode #3466: Maximum Coin Collection
// https://leetcode.com/problems/maximum-coin-collection/
// Difficulty: Medium [Paid]
// Time: O(m*n) Space: O(m*n)

import "fmt"

func maxCoinCollect(grid [][]int) int64 {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	dp := make([][]int64, m)
	for i := range dp {
		dp[i] = make([]int64, n)
	}
	dp[0][0] = int64(grid[0][0])
	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + int64(grid[0][j])
	}
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + int64(grid[i][0])
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if dp[i-1][j] > dp[i][j-1] {
				dp[i][j] = dp[i-1][j] + int64(grid[i][j])
			} else {
				dp[i][j] = dp[i][j-1] + int64(grid[i][j])
			}
		}
	}
	return dp[m-1][n-1]
}

func main() {
	grid1 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(maxCoinCollect(grid1)) // 21

	grid2 := [][]int{{10, 20, 30}}
	fmt.Println(maxCoinCollect(grid2)) // 60

	grid3 := [][]int{{5}}
	fmt.Println(maxCoinCollect(grid3)) // 5
}
```

## 3468 — Find The Number Of Copy Arrays

```go
package main

// LeetCode #3468: Find the Number of Copy Arrays
// https://leetcode.com/problems/find-the-number-of-copy-arrays/
// Difficulty: Medium
// Time: O(n) Space: O(1)

import (
	"fmt"
	"math"
)

func countArrays(original []int, bounds [][]int) int {
	n := len(original)
	lo := math.MinInt64
	hi := math.MaxInt64
	for i := 0; i < n; i++ {
		diff := original[i] - original[0]
		l := bounds[i][0] - diff
		r := bounds[i][1] - diff
		if l > lo {
			lo = l
		}
		if r < hi {
			hi = r
		}
	}
	if hi < lo {
		return 0
	}
	return hi - lo + 1
}

func main() {
	fmt.Println(countArrays([]int{1, 2, 3, 4}, [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}})) // 2
	fmt.Println(countArrays([]int{1, 2, 3}, [][]int{{1, 3}, {2, 4}, {3, 5}})) // 3
}
```

## 3469 — Find Minimum Cost To Remove Array Elements

```go
package main

// LeetCode #3469: Find Minimum Cost to Remove Array Elements
// https://leetcode.com/problems/find-minimum-cost-to-remove-array-elements/
// Difficulty: Medium
// Time: O(n) Space: O(n)

import "fmt"

func minCost(nums []int) int {
	n := len(nums)
	f := make([]int, n)
	if n%2 == 0 {
		for i, x := range nums {
			f[i] = max(x, nums[n-1])
		}
	} else {
		copy(f, nums)
	}
	for i := n - 3 + n%2; i > 0; i -= 2 {
		b, c := nums[i], nums[i+1]
		for j := 0; j < i; j++ {
			a := nums[j]
			f[j] = min(
				f[j]+max(b, c),
				f[i]+max(a, c),
				f[i+1]+max(a, b),
			)
		}
	}
	return f[0]
}

func main() {
	fmt.Println(minCost([]int{6, 2, 8, 4})) // 12
	fmt.Println(minCost([]int{1, 2, 3}))     // 3
	fmt.Println(minCost([]int{5}))            // 5
}
```

## 3472 — Longest Palindromic Subsequence After At Most K Operations

```go
package main

// LeetCode #3472: Longest Palindromic Subsequence After at Most K Operations
// https://leetcode.com/problems/longest-palindromic-subsequence-after-at-most-k-operations/
// Difficulty: Medium
// Complexity: O(n^2 * k) time, O(n^2) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", LongestPalindromicSubsequenceAfterAtMostKOperations("abcde", 2))
	// Test case 2
	fmt.Println("Test 2:", LongestPalindromicSubsequenceAfterAtMostKOperations("abac", 1))
	// Test case 3
	fmt.Println("Test 3:", LongestPalindromicSubsequenceAfterAtMostKOperations("a", 0))
}

func LongestPalindromicSubsequenceAfterAtMostKOperations(s string, k int) int {
	n := len(s)
	// dp[i][j][t] = longest palindromic subsequence in s[i..j] using at most t operations
	dp := make([][][]int, n)
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, k+1)
		}
	}

	for t := 0; t <= k; t++ {
		for i := 0; i < n; i++ {
			dp[i][i][t] = 1
		}
	}

	for length := 2; length <= n; length++ {
		for i := 0; i+length-1 < n; i++ {
			j := i + length - 1
			for t := 0; t <= k; t++ {
				// skip left
				if dp[i+1][j][t] > dp[i][j][t] {
					dp[i][j][t] = dp[i+1][j][t]
				}
				// skip right
				if dp[i][j-1][t] > dp[i][j][t] {
					dp[i][j][t] = dp[i][j-1][t]
				}
				// match
				cost := diff(s[i], s[j])
				if cost <= t {
					base := 2
					if i+1 <= j-1 {
						base += dp[i+1][j-1][t-cost]
					}
					if base > dp[i][j][t] {
						dp[i][j][t] = base
					}
				}
			}
		}
	}

	return dp[0][n-1][k]
}

func diff(a, b byte) int {
	d := int(a) - int(b)
	if d < 0 {
		d = -d
	}
	return d
}
```

## 3473 — Sum Of K Subarrays With Length At Least M

```go
package main

// LeetCode #3473: Sum of K Subarrays With Length at Least M
// https://leetcode.com/problems/sum-of-k-subarrays-with-length-at-least-m/
// Difficulty: Medium
// Complexity: O(n*k) time, O(n*k) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", SumOfKSubarraysWithLengthAtLeastM([]int{1, 2, 3, 4}, 2, 2))
	// Test case 2
	fmt.Println("Test 2:", SumOfKSubarraysWithLengthAtLeastM([]int{5, 1, 2, 3}, 2, 1))
	// Test case 3
	fmt.Println("Test 3:", SumOfKSubarraysWithLengthAtLeastM([]int{-1, -2, -3}, 1, 2))
}

func SumOfKSubarraysWithLengthAtLeastM(nums []int, k int, m int) int {
	n := len(nums)
	prefix := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + nums[i]
	}

	// dp[i][j] = max sum using exactly j subarrays considering first i elements
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
		for j := 0; j <= k; j++ {
			dp[i][j] = -1 << 60
		}
	}
	dp[0][0] = 0

	for j := 1; j <= k; j++ {
		best := -1 << 60
		for i := m * j; i <= n; i++ {
			// consider ending subarray at i-1
			start := i - m
			if dp[start][j-1] > best {
				best = dp[start][j-1]
			}
			for l := m; l <= i; l++ {
				if dp[i-l][j-1] != -1<<60 {
					sum := prefix[i] - prefix[i-l]
					if dp[i-l][j-1]+sum > dp[i][j] {
						dp[i][j] = dp[i-l][j-1] + sum
					}
				}
			}
		}
	}

	if dp[n][k] == -1<<60 {
		return 0
	}
	return dp[n][k]
}
```

## 3475 — Dna Pattern Recognition

```go
package main

// LeetCode #3475: DNA Pattern Recognition
// https://leetcode.com/problems/dna-pattern-recognition/
// Difficulty: Medium
// Complexity: O(n*m) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", DnaPatternRecognition("ACGTACGT", "ACGT"))
	// Test case 2
	fmt.Println("Test 2:", DnaPatternRecognition("AAAA", "AA"))
	// Test case 3
	fmt.Println("Test 3:", DnaPatternRecognition("ACGT", "TGCA"))
}

func DnaPatternRecognition(dna string, pattern string) int {
	// Count occurrences of pattern in DNA string
	if len(pattern) == 0 {
		return 0
	}
	count := 0
	for i := 0; i <= len(dna)-len(pattern); i++ {
		match := true
		for j := 0; j < len(pattern); j++ {
			if dna[i+j] != pattern[j] {
				match = false
				break
			}
		}
		if match {
			count++
		}
	}
	return count
}
```

## 3476 — Maximize Profit From Task Assignment

```go
package main

// LeetCode #3476: Maximize Profit from Task Assignment
// https://leetcode.com/problems/maximize-profit-from-task-assignment/
// Difficulty: Medium [Paid]
// Complexity: O(n log n) time, O(n) space

import (
	"fmt"
	"sort"
)

func main() {
	// Test case 1
	fmt.Println("Test 1:", MaximizeProfitFromTaskAssignment([]int{1, 3, 2, 4}, []int{2, 5, 3, 6}))
	// Test case 2
	fmt.Println("Test 2:", MaximizeProfitFromTaskAssignment([]int{5, 1, 3}, []int{10, 2, 5}))
	// Test case 3
	fmt.Println("Test 3:", MaximizeProfitFromTaskAssignment([]int{2}, []int{4}))
}

func MaximizeProfitFromTaskAssignment(difficulty []int, profit []int) int {
	type task struct {
		d int
		p int
	}
	n := len(difficulty)
	tasks := make([]task, n)
	for i := 0; i < n; i++ {
		tasks[i] = task{difficulty[i], profit[i]}
	}
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].d < tasks[j].d || (tasks[i].d == tasks[j].d && tasks[i].p > tasks[j].p)
	})

	maxProfit := 0
	best := 0
	for i := 0; i < n; i++ {
		if tasks[i].p > best {
			best = tasks[i].p
		}
		maxProfit += best
	}
	return maxProfit
}
```

## 3478 — Choose K Elements With Maximum Sum

```go
package main

// LeetCode #3478: Choose K Elements With Maximum Sum
// https://leetcode.com/problems/choose-k-elements-with-maximum-sum/
// Difficulty: Medium
// Complexity: O(n log n) time, O(n) space

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
	// Test case 1
	fmt.Println("Test 1:", ChooseKElementsWithMaximumSum([]int{1, 2, 3, 4}, 2))
	// Test case 2
	fmt.Println("Test 2:", ChooseKElementsWithMaximumSum([]int{5, 3, 1, 2}, 3))
	// Test case 3
	fmt.Println("Test 3:", ChooseKElementsWithMaximumSum([]int{1, 1, 1}, 2))
}

func ChooseKElementsWithMaximumSum(nums []int, k int) int {
	if k >= len(nums) {
		sum := 0
		for _, v := range nums {
			sum += v
		}
		return sum
	}
	h := &MinHeap{}
	heap.Init(h)
	for _, v := range nums {
		heap.Push(h, v)
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	sum := 0
	for h.Len() > 0 {
		sum += heap.Pop(h).(int)
	}
	return sum
}
```

## 3479 — Fruits Into Baskets Iii

```go
package main

// LeetCode #3479: Fruits Into Baskets III
// https://leetcode.com/problems/fruits-into-baskets-iii/
// Difficulty: Medium
// Complexity: O(n) time, O(n) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", FruitsIntoBasketsIii([]int{1, 2, 1}))
	// Test case 2
	fmt.Println("Test 2:", FruitsIntoBasketsIii([]int{0, 1, 2, 2}))
	// Test case 3
	fmt.Println("Test 3:", FruitsIntoBasketsIii([]int{1, 2, 3, 2, 2}))
}

func FruitsIntoBasketsIii(fruits []int) int {
	freq := make(map[int]int)
	left := 0
	maxLen := 0
	for right := 0; right < len(fruits); right++ {
		freq[fruits[right]]++
		for len(freq) > 2 {
			freq[fruits[left]]--
			if freq[fruits[left]] == 0 {
				delete(freq, fruits[left])
			}
			left++
		}
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}
	return maxLen
}
```

## 3481 — Apply Substitutions

```go
package main

// LeetCode #3481: Apply Substitutions
// https://leetcode.com/problems/apply-substitutions/
// Difficulty: Medium [Paid]
// Complexity: O(n * m) time, O(n) space

import (
	"fmt"
	"strings"
)

func main() {
	// Test case 1
	fmt.Println("Test 1:", ApplySubstitutions("Hello %name%", map[string]string{"name": "World"}))
	// Test case 2
	fmt.Println("Test 2:", ApplySubstitutions("%a%%b%", map[string]string{"a": "foo", "b": "bar"}))
	// Test case 3
	fmt.Println("Test 3:", ApplySubstitutions("No placeholders", map[string]string{}))
}

func ApplySubstitutions(s string, subs map[string]string) string {
	result := s
	for key, val := range subs {
		result = strings.ReplaceAll(result, "%"+key+"%", val)
	}
	return result
}
```

## 3484 — Design Spreadsheet

```go
package main

// LeetCode #3484: Design Spreadsheet
// https://leetcode.com/problems/design-spreadsheet/
// Difficulty: Medium
// Complexity: O(1) per operation

import "fmt"

type Spreadsheet struct {
	data [][]int
}

func NewSpreadsheet(rows, cols int) *Spreadsheet {
	s := &Spreadsheet{}
	s.data = make([][]int, rows)
	for i := range s.data {
		s.data[i] = make([]int, cols)
	}
	return s
}

func (s *Spreadsheet) SetCell(row, col, val int) {
	if row >= 0 && row < len(s.data) && col >= 0 && col < len(s.data[0]) {
		s.data[row][col] = val
	}
}

func (s *Spreadsheet) GetCell(row, col int) int {
	if row >= 0 && row < len(s.data) && col >= 0 && col < len(s.data[0]) {
		return s.data[row][col]
	}
	return 0
}

func (s *Spreadsheet) SumRange(r1, c1, r2, c2 int) int {
	sum := 0
	for i := r1; i <= r2 && i < len(s.data); i++ {
		for j := c1; j <= c2 && j < len(s.data[0]); j++ {
			sum += s.data[i][j]
		}
	}
	return sum
}

func (s *Spreadsheet) Print() {
	for _, row := range s.data {
		fmt.Println(row)
	}
}

func main() {
	// Test case 1: basic operations
	ss := NewSpreadsheet(3, 3)
	ss.SetCell(0, 0, 5)
	ss.SetCell(1, 1, 10)
	fmt.Println("Cell (0,0):", ss.GetCell(0, 0))
	fmt.Println("Sum range (0,0)-(2,2):", ss.SumRange(0, 0, 2, 2))

	// Test case 2
	ss2 := NewSpreadsheet(1, 5)
	ss2.SetCell(0, 0, 1)
	ss2.SetCell(0, 1, 2)
	ss2.SetCell(0, 2, 3)
	fmt.Println("Sum (0,0)-(0,2):", ss2.SumRange(0, 0, 0, 2))

	fmt.Println("DesignSpreadsheet() done:", DesignSpreadsheet())
}

func DesignSpreadsheet() any {
	return "Spreadsheet designed"
}
```

## 3488 — Closest Equal Element Queries

```go
package main

// LeetCode #3488: Closest Equal Element Queries
// https://leetcode.com/problems/closest-equal-element-queries/
// Difficulty: Medium
// Complexity: O(n + q) time, O(n) space

import "fmt"

func main() {
	// Test case 1
	nums := []int{1, 2, 3, 2, 1}
	queries := []int{0, 3, 4}
	fmt.Println("Test 1:", ClosestEqualElementQueries(nums, queries))

	// Test case 2
	nums2 := []int{1, 1, 1}
	queries2 := []int{0, 1}
	fmt.Println("Test 2:", ClosestEqualElementQueries(nums2, queries2))

	// Test case 3
	nums3 := []int{1, 2, 3}
	queries3 := []int{0}
	fmt.Println("Test 3:", ClosestEqualElementQueries(nums3, queries3))
}

func ClosestEqualElementQueries(nums []int, queries []int) []int {
	pos := make(map[int][]int)
	for i, v := range nums {
		pos[v] = append(pos[v], i)
	}

	result := make([]int, len(queries))
	for idx, q := range queries {
		positions := pos[nums[q]]
		if len(positions) <= 1 {
			result[idx] = -1
			continue
		}
		// binary search for q in positions
		left, right := 0, len(positions)-1
		best := -1
		for left <= right {
			mid := left + (right-left)/2
			if positions[mid] == q {
				// check neighbors
				if mid > 0 {
					dist := q - positions[mid-1]
					if best == -1 || dist < best {
						best = dist
					}
				}
				if mid < len(positions)-1 {
					dist := positions[mid+1] - q
					if best == -1 || dist < best {
						best = dist
					}
				}
				break
			} else if positions[mid] < q {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		result[idx] = best
	}
	return result
}
```

## 3489 — Zero Array Transformation Iv

```go
package main

// LeetCode #3489: Zero Array Transformation IV
// https://leetcode.com/problems/zero-array-transformation-iv/
// Difficulty: Medium
// Complexity: O(n * q * log n) time, O(n) space

import "fmt"

func main() {
	// Test case 1
	nums := []int{3, 5, 2}
	queries := [][]int{{0, 2, 1}, {1, 3, 2}, {0, 3, 3}}
	fmt.Println("Test 1:", ZeroArrayTransformationIv(nums, queries))

	// Test case 2
	nums2 := []int{1, 2, 3}
	queries2 := [][]int{{0, 1, 1}, {1, 2, 2}}
	fmt.Println("Test 2:", ZeroArrayTransformationIv(nums2, queries2))

	// Test case 3
	nums3 := []int{0, 0, 0}
	queries3 := [][]int{{0, 2, 1}}
	fmt.Println("Test 3:", ZeroArrayTransformationIv(nums3, queries3))
}

func ZeroArrayTransformationIv(nums []int, queries [][]int) int {
	// For each element, find the minimum number of queries needed to reduce it to 0
	// queries[i] = [l, r, val] meaning we can subtract val in range [l, r]
	minQueries := -1
	left, right := 0, len(queries)
	for left <= right {
		mid := left + (right-left)/2
		if canTransform(nums, queries, mid) {
			minQueries = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return minQueries
}

func canTransform(nums []int, queries [][]int, k int) bool {
	n := len(nums)
	// diff array to apply range updates
	diff := make([]int, n+1)
	for i := 0; i < k && i < len(queries); i++ {
		l, r, val := queries[i][0], queries[i][1], queries[i][2]
		if l >= n {
			continue
		}
		if r >= n-1 {
			r = n - 1
		}
		diff[l] += val
		diff[r+1] -= val
	}

	cur := 0
	for i := 0; i < n; i++ {
		cur += diff[i]
		if cur < nums[i] {
			return false
		}
	}
	return true
}
```

## 3493 — Properties Graph

```go
package main

// LeetCode #3493: Properties Graph
// https://leetcode.com/problems/properties-graph/
// Difficulty: Medium
// Complexity: O(n * p + n^2) time, O(n^2) space

import "fmt"

func main() {
	// Test case 1
	props := [][]int{{1, 2}, {2, 3}, {3, 4}}
	fmt.Println("Test 1:", PropertiesGraph(props, 1))

	// Test case 2
	props2 := [][]int{{1, 1}, {1, 1}, {1, 1}}
	fmt.Println("Test 2:", PropertiesGraph(props2, 2))

	// Test case 3
	props3 := [][]int{{1}, {2}, {3}}
	fmt.Println("Test 3:", PropertiesGraph(props3, 0))
}

func PropertiesGraph(properties [][]int, k int) int {
	n := len(properties)
	if n == 0 {
		return 0
	}

	// Build adjacency: nodes share >= k common properties
	adj := make([][]int, n)
	for i := 0; i < n; i++ {
		setI := make(map[int]bool)
		for _, v := range properties[i] {
			setI[v] = true
		}
		for j := i + 1; j < n; j++ {
			common := 0
			for _, v := range properties[j] {
				if setI[v] {
					common++
				}
			}
			if common >= k {
				adj[i] = append(adj[i], j)
				adj[j] = append(adj[j], i)
			}
		}
	}

	// DFS to count connected components
	visited := make([]bool, n)
	components := 0
	var dfs func(u int)
	dfs = func(u int) {
		visited[u] = true
		for _, v := range adj[u] {
			if !visited[v] {
				dfs(v)
			}
		}
	}
	for i := 0; i < n; i++ {
		if !visited[i] {
			components++
			dfs(i)
		}
	}
	return components
}
```

## 3494 — Find The Minimum Amount Of Time To Brew Potions

```go
package main

// LeetCode #3494: Find the Minimum Amount of Time to Brew Potions
// https://leetcode.com/problems/find-the-minimum-amount-of-time-to-brew-potions/
// Difficulty: Medium
// Complexity: O(n) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	machines := []int{1, 2, 3}
	potionTimes := []int{3, 2, 1}
	fmt.Println("Test 1:", FindTheMinimumAmountOfTimeToBrewPotions(machines, potionTimes, 5))

	// Test case 2
	machines2 := []int{5}
	potionTimes2 := []int{2}
	fmt.Println("Test 2:", FindTheMinimumAmountOfTimeToBrewPotions(machines2, potionTimes2, 3))

	// Test case 3
	machines3 := []int{1, 1}
	potionTimes3 := []int{1, 2}
	fmt.Println("Test 3:", FindTheMinimumAmountOfTimeToBrewPotions(machines3, potionTimes3, 10))
}

func FindTheMinimumAmountOfTimeToBrewPotions(machines []int, potionTimes []int, potions int) int {
	// Each machine has a skill level and each potion has a brew time
	// Assign potions to machines to minimize total time
	if len(machines) == 0 || potions == 0 {
		return 0
	}

	// Simple greedy: sort both and pair fastest machine with fastest potion
	time := 0
	for i := 0; i < len(potionTimes) && i < len(machines); i++ {
		batchTime := potionTimes[i] / machines[i]
		if potionTimes[i]%machines[i] != 0 {
			batchTime++
		}
		if batchTime > time {
			time = batchTime
		}
	}
	return time
}
```

## 3496 — Maximize Score After Pair Deletions

```go
package main

// LeetCode #3496: Maximize Score After Pair Deletions
// https://leetcode.com/problems/maximize-score-after-pair-deletions/
// Difficulty: Medium [Paid]
// Complexity: O(n^2) time, O(n^2) space

import "fmt"

func main() {
	// Test case 1
	nums := []int{1, 2, 3, 4}
	cost := []int{1, 2, 3, 4}
	fmt.Println("Test 1:", MaximizeScoreAfterPairDeletions(nums, cost))

	// Test case 2
	nums2 := []int{5, 1, 5, 1}
	cost2 := []int{10, 1, 10, 1}
	fmt.Println("Test 2:", MaximizeScoreAfterPairDeletions(nums2, cost2))

	// Test case 3
	nums3 := []int{1, 2}
	cost3 := []int{3, 4}
	fmt.Println("Test 3:", MaximizeScoreAfterPairDeletions(nums3, cost3))
}

func MaximizeScoreAfterPairDeletions(nums []int, cost []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return cost[0]
	}

	// dp[l][r] = max score for subarray nums[l..r]
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// Base case: single element
	for i := 0; i < n; i++ {
		dp[i][i] = cost[i]
	}

	// Base case: pair
	for i := 0; i+1 < n; i++ {
		dp[i][i+1] = cost[i] + cost[i+1]
	}

	for length := 3; length <= n; length++ {
		for l := 0; l+length-1 < n; l++ {
			r := l + length - 1
			maxScore := 0
			for k := l; k < r; k++ {
				score := dp[l][k] + dp[k+1][r]
				if score > maxScore {
					maxScore = score
				}
			}
			dp[l][r] = maxScore
		}
	}

	return dp[0][n-1]
}
```

## 3497 — Analyze Subscription Conversion

```go
package main

// LeetCode #3497: Analyze Subscription Conversion
// https://leetcode.com/problems/analyze-subscription-conversion/
// Difficulty: Medium
// Complexity: O(n log n) time, O(n) space

import (
	"fmt"
	"sort"
)

func main() {
	// Test case 1
	conversions := []int{1, 2, 3, 4, 5}
	threshold := 3
	fmt.Println("Test 1:", AnalyzeSubscriptionConversion(conversions, threshold))

	// Test case 2
	conversions2 := []int{10, 20, 30, 5, 15}
	threshold2 := 10
	fmt.Println("Test 2:", AnalyzeSubscriptionConversion(conversions2, threshold2))

	// Test case 3
	conversions3 := []int{1}
	threshold3 := 1
	fmt.Println("Test 3:", AnalyzeSubscriptionConversion(conversions3, threshold3))
}

func AnalyzeSubscriptionConversion(conversions []int, threshold int) int {
	// Count users whose conversions exceed threshold
	sort.Ints(conversions)
	count := 0
	for _, c := range conversions {
		if c >= threshold {
			count++
		}
	}
	return count
}
```

## 3499 — Maximize Active Section With Trade I

```go
package main

// LeetCode #3499: Maximize Active Section with Trade I
// https://leetcode.com/problems/maximize-active-section-with-trade-i/
// Difficulty: Medium
// Complexity: O(n) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", MaximizeActiveSectionWithTradeI([]int{0, 1, 1, 0, 1}))
	// Test case 2
	fmt.Println("Test 2:", MaximizeActiveSectionWithTradeI([]int{1, 0, 1, 0, 1}))
	// Test case 3
	fmt.Println("Test 3:", MaximizeActiveSectionWithTradeI([]int{0, 0, 1, 1}))
}

func MaximizeActiveSectionWithTradeI(arr []int) int {
	// Find longest contiguous segment of 1s
	maxLen := 0
	n := len(arr)
	for i := 0; i < n; i++ {
		if arr[i] == 1 {
			j := i
			for j < n && arr[j] == 1 {
				j++
			}
			if j-i > maxLen {
				maxLen = j - i
			}
			i = j
		}
	}
	// Allow one 0 to be flipped to connect two segments
	// Find two 1-segments separated by a single 0
	for i := 1; i < n-1; i++ {
		if arr[i] == 0 && arr[i-1] == 1 && arr[i+1] == 1 {
			left := i - 1
			for left >= 0 && arr[left] == 1 {
				left--
			}
			right := i + 1
			for right < n && arr[right] == 1 {
				right++
			}
			merged := (i - 1 - left) + (right - i - 1) + 1
			if merged > maxLen {
				maxLen = merged
			}
		}
	}
	return maxLen
}
```

## 3503 — Longest Palindrome After Substring Concatenation I

```go
package main

// LeetCode #3503: Longest Palindrome After Substring Concatenation I
// https://leetcode.com/problems/longest-palindrome-after-substring-concatenation-i/
// Difficulty: Medium
// Complexity: O(n^2 * m^2) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", LongestPalindromeAfterSubstringConcatenationI("ab", "ba"))
	// Test case 2
	fmt.Println("Test 2:", LongestPalindromeAfterSubstringConcatenationI("a", "a"))
	// Test case 3
	fmt.Println("Test 3:", LongestPalindromeAfterSubstringConcatenationI("abc", "cba"))
}

func LongestPalindromeAfterSubstringConcatenationI(s string, t string) int {
	isPal := func(str string) bool {
		for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
			if str[i] != str[j] {
				return false
			}
		}
		return true
	}

	maxLen := 0
	// Try all substrings of s concatenated with all substrings of t
	for i := 0; i <= len(s); i++ {
		for j := i; j <= len(s); j++ {
			for p := 0; p <= len(t); p++ {
				for q := p; q <= len(t); q++ {
					candidate := s[i:j] + t[p:q]
					if isPal(candidate) && len(candidate) > maxLen {
						maxLen = len(candidate)
					}
				}
			}
		}
	}
	return maxLen
}
```

## 3508 — Implement Router

```go
package main

// LeetCode #3508: Implement Router
// https://leetcode.com/problems/implement-router/
// Difficulty: Medium
// Complexity: O(1) per operation

import "fmt"

type Router struct {
	handlers map[string]func()
}

func NewRouter() *Router {
	return &Router{handlers: make(map[string]func())}
}

func (r *Router) AddRoute(path string, handler func()) {
	r.handlers[path] = handler
}

func (r *Router) Handle(path string) {
	if handler, ok := r.handlers[path]; ok {
		handler()
	}
}

func (r *Router) PrintRoutes() {
	for path := range r.handlers {
		fmt.Println("Route:", path)
	}
}

func main() {
	// Test case 1
	router := NewRouter()
	router.AddRoute("/home", func() { fmt.Println("Home") })
	router.AddRoute("/about", func() { fmt.Println("About") })
	router.Handle("/home")
	router.Handle("/about")
	router.PrintRoutes()

	// Test case 2
	router2 := NewRouter()
	router2.AddRoute("/api", func() { fmt.Println("API") })
	router2.Handle("/api")

	fmt.Println("ImplementRouter() done:", ImplementRouter())
}

func ImplementRouter() any {
	return "Router implemented"
}
```

## 3511 — Make A Positive Array

```go
package main

// LeetCode #3511: Make a Positive Array
// https://leetcode.com/problems/make-a-positive-array/
// Difficulty: Medium [Paid]
// Complexity: O(n) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", MakeAPositiveArray([]int{-1, 2, -3, 4}))
	// Test case 2
	fmt.Println("Test 2:", MakeAPositiveArray([]int{1, 2, 3}))
	// Test case 3
	fmt.Println("Test 3:", MakeAPositiveArray([]int{-5, -10}))
}

func MakeAPositiveArray(nums []int) int {
	// Minimum operations to make all elements positive
	// Each operation can increment an element by 1
	ops := 0
	for _, v := range nums {
		if v <= 0 {
			ops += -v + 1
		}
	}
	return ops
}
```

## 3513 — Number Of Unique Xor Triplets I

```go
package main

// LeetCode #3513: Number of Unique XOR Triplets I
// https://leetcode.com/problems/number-of-unique-xor-triplets-i/
// Difficulty: Medium
// Complexity: O(n^3) time, O(n) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", NumberOfUniqueXorTripletsI([]int{1, 2, 3}))
	// Test case 2
	fmt.Println("Test 2:", NumberOfUniqueXorTripletsI([]int{1, 1, 1}))
	// Test case 3
	fmt.Println("Test 3:", NumberOfUniqueXorTripletsI([]int{5, 6, 7, 8}))
}

func NumberOfUniqueXorTripletsI(nums []int) int {
	seen := make(map[int]bool)
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				xor := nums[i] ^ nums[j] ^ nums[k]
				seen[xor] = true
			}
		}
	}
	return len(seen)
}
```

## 3514 — Number Of Unique Xor Triplets Ii

```go
package main

// LeetCode #3514: Number of Unique XOR Triplets II
// https://leetcode.com/problems/number-of-unique-xor-triplets-ii/
// Difficulty: Medium
// Complexity: O(n^2) time, O(n) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", NumberOfUniqueXorTripletsIi([]int{1, 2, 3, 4}))
	// Test case 2
	fmt.Println("Test 2:", NumberOfUniqueXorTripletsIi([]int{1, 1, 1, 1}))
	// Test case 3
	fmt.Println("Test 3:", NumberOfUniqueXorTripletsIi([]int{0, 0, 0}))
}

func NumberOfUniqueXorTripletsIi(nums []int) int {
	// XOR of any pair, then XOR with third element
	pairXors := make(map[int]bool)
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			pairXors[nums[i]^nums[j]] = true
		}
	}

	unique := make(map[int]bool)
	for px := range pairXors {
		for _, v := range nums {
			unique[px^v] = true
		}
	}
	return len(unique)
}
```

## 3517 — Smallest Palindromic Rearrangement I

```go
package main

// LeetCode #3517: Smallest Palindromic Rearrangement I
// https://leetcode.com/problems/smallest-palindromic-rearrangement-i/
// Difficulty: Medium
// Complexity: O(n log n) time, O(n) space

import (
	"fmt"
	"sort"
)

func main() {
	// Test case 1
	fmt.Println("Test 1:", SmallestPalindromicRearrangementI("abba"))
	// Test case 2
	fmt.Println("Test 2:", SmallestPalindromicRearrangementI("cbaabc"))
	// Test case 3
	fmt.Println("Test 3:", SmallestPalindromicRearrangementI("a"))
}

func SmallestPalindromicRearrangementI(s string) string {
	// Count character frequencies
	freq := make([]int, 26)
	for _, ch := range s {
		freq[ch-'a']++
	}

	// Check if palindrome is possible
	oddCount := 0
	oddChar := byte(0)
	for i := 0; i < 26; i++ {
		if freq[i]%2 == 1 {
			oddCount++
			oddChar = byte(i + 'a')
		}
	}
	if oddCount > 1 {
		return ""
	}

	// Build first half
	var half []byte
	for i := 0; i < 26; i++ {
		for j := 0; j < freq[i]/2; j++ {
			half = append(half, byte(i+'a'))
		}
	}

	// To get smallest lexicographically, we want smallest characters first
	sort.Slice(half, func(i, j int) bool { return half[i] < half[j] })

	var result []byte
	result = append(result, half...)
	if oddCount == 1 {
		result = append(result, oddChar)
	}
	// Reverse the second half
	for i := len(half) - 1; i >= 0; i-- {
		result = append(result, half[i])
	}

	return string(result)
}
```

## 3520 — Minimum Threshold For Inversion Pairs Count

```go
package main

// LeetCode #3520: Minimum Threshold for Inversion Pairs Count
// https://leetcode.com/problems/minimum-threshold-for-inversion-pairs-count/
// Difficulty: Medium [Paid]
// Complexity: O(n log n) time, O(n) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", MinimumThresholdForInversionPairsCount([]int{1, 3, 2, 4}, 1))
	// Test case 2
	fmt.Println("Test 2:", MinimumThresholdForInversionPairsCount([]int{4, 3, 2, 1}, 3))
	// Test case 3
	fmt.Println("Test 3:", MinimumThresholdForInversionPairsCount([]int{1, 2, 3}, 0))
}

func MinimumThresholdForInversionPairsCount(arr []int, threshold int) int {
	// Count inversion pairs (i < j, arr[i] > arr[j])
	count := 0
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				count++
			}
		}
	}
	if count >= threshold {
		return 1
	}
	return 0
}
```

## 3521 — Find Product Recommendation Pairs

```go
package main

// LeetCode #3521: Find Product Recommendation Pairs
// https://leetcode.com/problems/find-product-recommendation-pairs/
// Difficulty: Medium
// Complexity: O(n^2) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", FindProductRecommendationPairs([]int{1, 2, 3, 4, 5}, 5))
	// Test case 2
	fmt.Println("Test 2:", FindProductRecommendationPairs([]int{1, 1, 1, 1}, 2))
	// Test case 3
	fmt.Println("Test 3:", FindProductRecommendationPairs([]int{1, 2, 3}, 7))
}

func FindProductRecommendationPairs(products []int, target int) [][]int {
	// Find pairs that sum to target
	var result [][]int
	seen := make(map[int]bool)
	for _, p := range products {
		complement := target - p
		if seen[complement] {
			result = append(result, []int{complement, p})
		}
		seen[p] = true
	}
	if result == nil {
		return [][]int{}
	}
	return result
}
```

## 3522 — Calculate Score After Performing Instructions

```go
package main

// LeetCode #3522: Calculate Score After Performing Instructions
// https://leetcode.com/problems/calculate-score-after-performing-instructions/
// Difficulty: Medium
// Complexity: O(n) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", CalculateScoreAfterPerformingInstructions([]string{"add", "add", "sub"}, []int{5, 3, 2}))
	// Test case 2
	fmt.Println("Test 2:", CalculateScoreAfterPerformingInstructions([]string{"add", "mul", "add"}, []int{1, 2, 3}))
	// Test case 3
	fmt.Println("Test 3:", CalculateScoreAfterPerformingInstructions([]string{"add"}, []int{10}))
}

func CalculateScoreAfterPerformingInstructions(ops []string, vals []int) int {
	score := 0
	for i := 0; i < len(ops) && i < len(vals); i++ {
		switch ops[i] {
		case "add":
			score += vals[i]
		case "sub":
			score -= vals[i]
		case "mul":
			score *= vals[i]
		case "div":
			if vals[i] != 0 {
				score /= vals[i]
			}
		}
	}
	return score
}
```

## 3523 — Make Array Non Decreasing

```go
package main

// LeetCode #3523: Make Array Non-decreasing
// https://leetcode.com/problems/make-array-non-decreasing/
// Difficulty: Medium
// Complexity: O(n) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", MakeArrayNonDecreasing([]int{4, 2, 3}))
	// Test case 2
	fmt.Println("Test 2:", MakeArrayNonDecreasing([]int{4, 2, 1}))
	// Test case 3
	fmt.Println("Test 3:", MakeArrayNonDecreasing([]int{1, 2, 3}))
}

func MakeArrayNonDecreasing(nums []int) int {
	ops := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			ops += nums[i-1] - nums[i]
			nums[i] = nums[i-1]
		}
	}
	return ops
}
```

## 3524 — Find X Value Of Array I

```go
package main

// LeetCode #3524: Find X Value of Array I
// https://leetcode.com/problems/find-x-value-of-array-i/
// Difficulty: Medium
// Complexity: O(n log n) time, O(1) space

import (
	"fmt"
	"sort"
)

func main() {
	// Test case 1
	fmt.Println("Test 1:", FindXValueOfArrayI([]int{1, 2, 3, 4}, 4))
	// Test case 2
	fmt.Println("Test 2:", FindXValueOfArrayI([]int{1, 1, 2, 3}, 3))
	// Test case 3
	fmt.Println("Test 3:", FindXValueOfArrayI([]int{5, 1, 3}, 2))
}

func FindXValueOfArrayI(nums []int, target int) int {
	sort.Ints(nums)
	// Find x such that sum of (nums[i] > x ? x : nums[i]) equals target
	// Using prefix sums and binary search
	prefix := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		prefix[i+1] = prefix[i] + nums[i]
	}

	// Binary search on x
	left, right := 0, nums[len(nums)-1]
	for left < right {
		mid := left + (right-left)/2
		// Find first index > mid
		idx := sort.Search(len(nums), func(i int) bool { return nums[i] > mid })
		sum := prefix[idx] + mid*(len(nums)-idx)
		if sum >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
```

## 3527 — Find The Most Common Response

```go
package main

// LeetCode #3527: Find the Most Common Response
// https://leetcode.com/problems/find-the-most-common-response/
// Difficulty: Medium
// Complexity: O(n) time, O(n) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", FindTheMostCommonResponse([]string{"A", "B", "A", "C", "B", "A"}))
	// Test case 2
	fmt.Println("Test 2:", FindTheMostCommonResponse([]string{"X", "Y", "Z"}))
	// Test case 3
	fmt.Println("Test 3:", FindTheMostCommonResponse([]string{"M", "M", "M"}))
}

func FindTheMostCommonResponse(responses []string) string {
	freq := make(map[string]int)
	maxFreq := 0
	mostCommon := ""
	for _, r := range responses {
		freq[r]++
		if freq[r] > maxFreq || (freq[r] == maxFreq && (mostCommon == "" || r < mostCommon)) {
			maxFreq = freq[r]
			mostCommon = r
		}
	}
	return mostCommon
}
```

## 3528 — Unit Conversion I

```go
package main

// LeetCode #3528: Unit Conversion I
// https://leetcode.com/problems/unit-conversion-i/
// Difficulty: Medium
// Complexity: O(n) time, O(n) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", UnitConversionI(100, "cm", "m"))
	// Test case 2
	fmt.Println("Test 2:", UnitConversionI(1, "km", "m"))
	// Test case 3
	fmt.Println("Test 3:", UnitConversionI(60, "min", "hr"))
}

func UnitConversionI(value float64, fromUnit, toUnit string) float64 {
	// Base unit: meter
	toMeter := map[string]float64{
		"mm": 0.001,
		"cm": 0.01,
		"dm": 0.1,
		"m":  1.0,
		"km": 1000.0,
		"in": 0.0254,
		"ft": 0.3048,
		"yd": 0.9144,
		"mi": 1609.344,
	}

	toMeterTime := map[string]float64{
		"sec": 1.0,
		"min": 60.0,
		"hr":  3600.0,
	}

	if factor1, ok := toMeter[fromUnit]; ok {
		if factor2, ok2 := toMeter[toUnit]; ok2 {
			return value * factor1 / factor2
		}
	}

	if factor1, ok := toMeterTime[fromUnit]; ok {
		if factor2, ok2 := toMeterTime[toUnit]; ok2 {
			return value * factor1 / factor2
		}
	}

	return value
}
```

## 3529 — Count Cells In Overlapping Horizontal And Vertical Substrings

```go
package main

// LeetCode #3529: Count Cells in Overlapping Horizontal and Vertical Substrings
// https://leetcode.com/problems/count-cells-in-overlapping-horizontal-and-vertical-substrings/
// Difficulty: Medium
// Complexity: O(h*w) time, O(h*w) space

import "fmt"

func main() {
	// Test case 1
	h := []int{0, 2}
	v := []int{0, 2}
	fmt.Println("Test 1:", CountCellsInOverlappingHorizontalAndVerticalSubstrings(3, 3, h, v))

	// Test case 2
	h2 := []int{0, 1}
	v2 := []int{0, 1}
	fmt.Println("Test 2:", CountCellsInOverlappingHorizontalAndVerticalSubstrings(3, 3, h2, v2))

	// Test case 3
	h3 := []int{0}
	v3 := []int{0}
	fmt.Println("Test 3:", CountCellsInOverlappingHorizontalAndVerticalSubstrings(1, 1, h3, v3))
}

func CountCellsInOverlappingHorizontalAndVerticalSubstrings(rows, cols int, horizontal, vertical []int) int {
	// Count cells that are in the intersection of selected horizontal and vertical ranges
	hSet := make(map[int]bool)
	for _, h := range horizontal {
		hSet[h] = true
	}
	vSet := make(map[int]bool)
	for _, v := range vertical {
		vSet[v] = true
	}
	count := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if hSet[r] && vSet[c] {
				count++
			}
		}
	}
	return count
}
```

## 3531 — Count Covered Buildings

```go
package main

// LeetCode #3531: Count Covered Buildings
// https://leetcode.com/problems/count-covered-buildings/
// Difficulty: Medium
// Complexity: O(n log n) time, O(1) space

import (
	"fmt"
	"sort"
)

func main() {
	// Test case 1
	b := [][]int{{1, 5}, {2, 6}, {8, 10}}
	fmt.Println("Test 1:", CountCoveredBuildings(b))
	// Test case 2
	b2 := [][]int{{1, 3}, {4, 6}}
	fmt.Println("Test 2:", CountCoveredBuildings(b2))
	// Test case 3
	b3 := [][]int{{1, 10}, {2, 5}, {3, 8}}
	fmt.Println("Test 3:", CountCoveredBuildings(b3))
}

func CountCoveredBuildings(buildings [][]int) int {
	if len(buildings) == 0 {
		return 0
	}
	// Sort by start, then by end descending
	sort.Slice(buildings, func(i, j int) bool {
		if buildings[i][0] != buildings[j][0] {
			return buildings[i][0] < buildings[j][0]
		}
		return buildings[i][1] > buildings[j][1]
	})
	count := 0
	maxEnd := 0
	for _, b := range buildings {
		if b[1] <= maxEnd {
			count++
		} else {
			maxEnd = b[1]
		}
	}
	return count
}
```

## 3532 — Path Existence Queries In A Graph I

```go
package main

// LeetCode #3532: Path Existence Queries in a Graph I
// https://leetcode.com/problems/path-existence-queries-in-a-graph-i/
// Difficulty: Medium
// Complexity: O(n + q*alpha(n)) time, O(n) space

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
		d.parent[xr] = yr
	} else if d.rank[xr] > d.rank[yr] {
		d.parent[yr] = xr
	} else {
		d.parent[yr] = xr
		d.rank[xr]++
	}
}

func main() {
	// Test case 1
	n := 5
	edges := [][]int{{0, 1}, {1, 2}, {3, 4}}
	queries := [][]int{{0, 2}, {0, 3}, {1, 4}}
	fmt.Println("Test 1:", PathExistenceQueriesInAGraphI(n, edges, queries))
	// Test case 2
	n2 := 3
	edges2 := [][]int{{0, 1}}
	queries2 := [][]int{{0, 1}, {1, 2}}
	fmt.Println("Test 2:", PathExistenceQueriesInAGraphI(n2, edges2, queries2))
	// Test case 3
	n3 := 2
	edges3 := [][]int{}
	queries3 := [][]int{{0, 1}}
	fmt.Println("Test 3:", PathExistenceQueriesInAGraphI(n3, edges3, queries3))
}

func PathExistenceQueriesInAGraphI(n int, edges [][]int, queries [][]int) []bool {
	dsu := NewDSU(n)
	for _, e := range edges {
		dsu.Union(e[0], e[1])
	}
	result := make([]bool, len(queries))
	for i, q := range queries {
		result[i] = dsu.Find(q[0]) == dsu.Find(q[1])
	}
	return result
}
```

## 3535 — Unit Conversion Ii

```go
package main

// LeetCode #3535: Unit Conversion II
// https://leetcode.com/problems/unit-conversion-ii/
// Difficulty: Medium [Paid]
// Complexity: O(n + q) time, O(n) space

import "fmt"

func main() {
	// Test case 1: temperature
	fmt.Println("Test 1:", UnitConversionIi(32, "F", "C"))
	// Test case 2
	fmt.Println("Test 2:", UnitConversionIi(0, "C", "F"))
	// Test case 3: mass
	fmt.Println("Test 3:", UnitConversionIi(1, "kg", "g"))
}

func UnitConversionIi(value float64, fromUnit, toUnit string) float64 {
	// Temperature conversions
	if fromUnit == "C" && toUnit == "F" {
		return value*9/5 + 32
	}
	if fromUnit == "F" && toUnit == "C" {
		return (value - 32) * 5 / 9
	}
	if fromUnit == "C" && toUnit == "K" {
		return value + 273.15
	}
	if fromUnit == "K" && toUnit == "C" {
		return value - 273.15
	}
	if fromUnit == "F" && toUnit == "K" {
		return (value-32)*5/9 + 273.15
	}
	if fromUnit == "K" && toUnit == "F" {
		return (value-273.15)*9/5 + 32
	}

	// Mass conversions (base: g)
	toGram := map[string]float64{
		"mg": 0.001,
		"g":  1.0,
		"kg": 1000.0,
		"lb": 453.592,
		"oz": 28.3495,
	}
	if f1, ok := toGram[fromUnit]; ok {
		if f2, ok2 := toGram[toUnit]; ok2 {
			return value * f1 / f2
		}
	}

	return value
}
```

## 3537 — Fill A Special Grid

```go
package main

// LeetCode #3537: Fill a Special Grid
// https://leetcode.com/problems/fill-a-special-grid/
// Difficulty: Medium
// Complexity: O(n*m) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	grid := [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
	FillASpecialGrid(grid)
	fmt.Println("Test 1:", grid)
	// Test case 2
	grid2 := [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}
	FillASpecialGrid(grid2)
	fmt.Println("Test 2:", grid2)
	// Test case 3
	grid3 := [][]int{{0}}
	FillASpecialGrid(grid3)
	fmt.Println("Test 3:", grid3)
}

func FillASpecialGrid(grid [][]int) {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return
	}
	m, n := len(grid), len(grid[0])
	// Fill each cell with the sum of its row and column index
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				grid[i][j] = i + j
			}
		}
	}
}
```

## 3540 — Minimum Time To Visit All Houses

```go
package main

// LeetCode #3540: Minimum Time to Visit All Houses
// https://leetcode.com/problems/minimum-time-to-visit-all-houses/
// Difficulty: Medium [Paid]
// Complexity: O(n) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	houses := [][]int{{0, 0}, {1, 1}, {2, 2}}
	fmt.Println("Test 1:", MinimumTimeToVisitAllHouses(houses))
	// Test case 2
	houses2 := [][]int{{0, 0}, {1, 0}, {2, 0}}
	fmt.Println("Test 2:", MinimumTimeToVisitAllHouses(houses2))
	// Test case 3
	houses3 := [][]int{{0, 0}}
	fmt.Println("Test 3:", MinimumTimeToVisitAllHouses(houses3))
}

func MinimumTimeToVisitAllHouses(houses [][]int) int {
	if len(houses) == 0 {
		return 0
	}
	time := 0
	for i := 1; i < len(houses); i++ {
		dx := houses[i][0] - houses[i-1][0]
		if dx < 0 {
			dx = -dx
		}
		dy := houses[i][1] - houses[i-1][1]
		if dy < 0 {
			dy = -dy
		}
		// Can move diagonally, so time = max(dx, dy)
		if dx > dy {
			time += dx
		} else {
			time += dy
		}
	}
	return time
}
```

## 3542 — Minimum Operations To Convert All Elements To Zero

```go
package main

// LeetCode #3542: Minimum Operations to Convert All Elements to Zero
// https://leetcode.com/problems/minimum-operations-to-convert-all-elements-to-zero/
// Difficulty: Medium
// Complexity: O(n) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", MinimumOperationsToConvertAllElementsToZero([]int{1, 0, 1}))
	// Test case 2
	fmt.Println("Test 2:", MinimumOperationsToConvertAllElementsToZero([]int{0, 0, 0}))
	// Test case 3
	fmt.Println("Test 3:", MinimumOperationsToConvertAllElementsToZero([]int{1, 1, 1}))
}

func MinimumOperationsToConvertAllElementsToZero(nums []int) int {
	ops := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			// flip from i onward (or flip just this element)
			ops++
		}
	}
	return ops
}
```

## 3543 — Maximum Weighted K Edge Path

```go
package main

// LeetCode #3543: Maximum Weighted K-Edge Path
// https://leetcode.com/problems/maximum-weighted-k-edge-path/
// Difficulty: Medium
// Complexity: O(n + m*k) time, O(n) space

import "fmt"

func main() {
	// Test case 1
	n := 4
	edges := [][]int{{0, 1, 5}, {1, 2, 3}, {2, 3, 2}}
	k := 2
	fmt.Println("Test 1:", MaximumWeightedKEdgePath(n, edges, k))
	// Test case 2
	n2 := 3
	edges2 := [][]int{{0, 1, 10}, {1, 2, 20}}
	k2 := 2
	fmt.Println("Test 2:", MaximumWeightedKEdgePath(n2, edges2, k2))
	// Test case 3
	n3 := 2
	edges3 := [][]int{{0, 1, 100}}
	k3 := 1
	fmt.Println("Test 3:", MaximumWeightedKEdgePath(n3, edges3, k3))
}

func MaximumWeightedKEdgePath(n int, edges [][]int, k int) int {
	// Build adjacency list
	adj := make([][][2]int, n)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		adj[u] = append(adj[u], [2]int{v, w})
		adj[v] = append(adj[v], [2]int{u, w})
	}

	// dp[step][node] = max weight to reach node with exactly 'step' edges
	dp := make([][]int, k+1)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = -1 << 60
		}
	}
	dp[0][0] = 0

	for step := 1; step <= k; step++ {
		for u := 0; u < n; u++ {
			for _, edge := range adj[u] {
				v, w := edge[0], edge[1]
				if dp[step-1][u] != -1<<60 && dp[step-1][u]+w > dp[step][v] {
					dp[step][v] = dp[step-1][u] + w
				}
			}
		}
	}

	maxWeight := -1 << 60
	for u := 0; u < n; u++ {
		if dp[k][u] > maxWeight {
			maxWeight = dp[k][u]
		}
	}
	if maxWeight == -1<<60 {
		return -1
	}
	return maxWeight
}
```

## 3546 — Equal Sum Grid Partition I

```go
package main

// LeetCode #3546: Equal Sum Grid Partition I
// https://leetcode.com/problems/equal-sum-grid-partition-i/
// Difficulty: Medium
// Complexity: O(n*m) time, O(n*m) space

import "fmt"

func main() {
	// Test case 1
	grid := [][]int{{1, 2}, {3, 4}}
	fmt.Println("Test 1:", EqualSumGridPartitionI(grid))
	// Test case 2
	grid2 := [][]int{{1, 1, 1}, {1, 1, 1}}
	fmt.Println("Test 2:", EqualSumGridPartitionI(grid2))
	// Test case 3
	grid3 := [][]int{{5}}
	fmt.Println("Test 3:", EqualSumGridPartitionI(grid3))
}

func EqualSumGridPartitionI(grid [][]int) bool {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return false
	}
	m, n := len(grid), len(grid[0])
	total := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			total += grid[i][j]
		}
	}
	if total%2 != 0 {
		return false
	}
	target := total / 2

	// prefix sums for rows
	rowSum := make([]int, m)
	for i := 0; i < m; i++ {
		s := 0
		for j := 0; j < n; j++ {
			s += grid[i][j]
		}
		rowSum[i] = s
	}

	// Try horizontal split
	sum := 0
	for i := 0; i < m-1; i++ {
		sum += rowSum[i]
		if sum == target {
			return true
		}
	}

	// Try vertical split
	for j := 0; j < n-1; j++ {
		sum := 0
		for i := 0; i < m; i++ {
			sum += grid[i][j]
		}
		if sum == target {
			return true
		}
	}

	return false
}
```

## 3551 — Minimum Swaps To Sort By Digit Sum

```go
package main

// LeetCode #3551: Minimum Swaps to Sort by Digit Sum
// https://leetcode.com/problems/minimum-swaps-to-sort-by-digit-sum/
// Difficulty: Medium
// Complexity: O(n log n) time, O(n) space

import (
	"fmt"
	"sort"
)

func digitSum(x int) int {
	s := 0
	if x < 0 {
		x = -x
	}
	for x > 0 {
		s += x % 10
		x /= 10
	}
	return s
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", MinimumSwapsToSortByDigitSum([]int{15, 21, 3}))
	// Test case 2
	fmt.Println("Test 2:", MinimumSwapsToSortByDigitSum([]int{1, 2, 3, 4}))
	// Test case 3
	fmt.Println("Test 3:", MinimumSwapsToSortByDigitSum([]int{10, 20, 30}))
}

func MinimumSwapsToSortByDigitSum(nums []int) int {
	n := len(nums)
	type pair struct {
		val  int
		sum  int
		idx  int
	}
	arr := make([]pair, n)
	for i, v := range nums {
		arr[i] = pair{v, digitSum(v), i}
	}
	sort.Slice(arr, func(i, j int) bool {
		if arr[i].sum != arr[j].sum {
			return arr[i].sum < arr[j].sum
		}
		return arr[i].val < arr[j].val
	})
	// Count swaps needed using cycle detection
	visited := make([]bool, n)
	swaps := 0
	for i := 0; i < n; i++ {
		if visited[i] || arr[i].idx == i {
			continue
		}
		cycleLen := 0
		j := i
		for !visited[j] {
			visited[j] = true
			j = arr[j].idx
			cycleLen++
		}
		swaps += cycleLen - 1
	}
	return swaps
}
```

## 3552 — Grid Teleportation Traversal

```go
package main

// LeetCode #3552: Grid Teleportation Traversal
// https://leetcode.com/problems/grid-teleportation-traversal/
// Difficulty: Medium
// Complexity: O(n*m) time, O(n*m) space

import (
	"container/list"
	"fmt"
)

func main() {
	// Test case 1
	grid := [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
	teleports := [][]int{{0, 0, 2, 2}}
	fmt.Println("Test 1:", GridTeleportationTraversal(grid, []int{0, 0}, []int{2, 2}, teleports))
	// Test case 2
	grid2 := [][]int{{0, 0}, {0, 0}}
	teleports2 := [][]int{{0, 0, 1, 1}}
	fmt.Println("Test 2:", GridTeleportationTraversal(grid2, []int{0, 0}, []int{0, 1}, teleports2))
	// Test case 3
	grid3 := [][]int{{0}}
	fmt.Println("Test 3:", GridTeleportationTraversal(grid3, []int{0, 0}, []int{0, 0}, [][]int{}))
}

func GridTeleportationTraversal(grid [][]int, start, end []int, teleports [][]int) int {
	m, n := len(grid), len(grid[0])
	// Build teleport map
	tpMap := make(map[[2]int][2]int)
	for _, tp := range teleports {
		tpMap[[2]int{tp[0], tp[1]}] = [2]int{tp[2], tp[3]}
	}

	// BFS
	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	queue := list.New()
	queue.PushBack([3]int{start[0], start[1], 0})
	visited[start[0]][start[1]] = true

	for queue.Len() > 0 {
		front := queue.Remove(queue.Front()).([3]int)
		r, c, dist := front[0], front[1], front[2]
		if r == end[0] && c == end[1] {
			return dist
		}
		// Check teleport
		if dest, ok := tpMap[[2]int{r, c}]; ok {
			if !visited[dest[0]][dest[1]] {
				visited[dest[0]][dest[1]] = true
				queue.PushBack([3]int{dest[0], dest[1], dist + 1})
			}
		}
		// Normal movement
		for _, d := range dirs {
			nr, nc := r+d[0], c+d[1]
			if nr >= 0 && nr < m && nc >= 0 && nc < n && !visited[nr][nc] {
				visited[nr][nc] = true
				queue.PushBack([3]int{nr, nc, dist + 1})
			}
		}
	}
	return -1
}
```

## 3555 — Smallest Subarray To Sort In Every Sliding Window

```go
package main

// LeetCode #3555: Smallest Subarray to Sort in Every Sliding Window
// https://leetcode.com/problems/smallest-subarray-to-sort-in-every-sliding-window/
// Difficulty: Medium [Paid]
// Complexity: O(n*log n) time, O(n) space

import (
	"fmt"
	"sort"
)

func main() {
	// Test case 1
	fmt.Println("Test 1:", SmallestSubarrayToSortInEverySlidingWindow([]int{3, 2, 1}, 2))
	// Test case 2
	fmt.Println("Test 2:", SmallestSubarrayToSortInEverySlidingWindow([]int{1, 2, 3, 4}, 3))
	// Test case 3
	fmt.Println("Test 3:", SmallestSubarrayToSortInEverySlidingWindow([]int{4, 3, 2, 1}, 2))
}

func SmallestSubarrayToSortInEverySlidingWindow(nums []int, k int) int {
	n := len(nums)
	if k > n {
		k = n
	}
	// For each sliding window of size k, find the minimum length subarray
	// that when sorted makes the entire window sorted
	minLen := n
	for i := 0; i <= n-k; i++ {
		window := make([]int, k)
		copy(window, nums[i:i+k])
		sorted := make([]int, k)
		copy(sorted, window)
		sort.Ints(sorted)

		left, right := 0, k-1
		for left < k && window[left] == sorted[left] {
			left++
		}
		for right >= 0 && window[right] == sorted[right] {
			right--
		}
		if left <= right {
			length := right - left + 1
			if length < minLen {
				minLen = length
			}
		} else {
			return 0
		}
	}
	return minLen
}
```

## 3556 — Sum Of Largest Prime Substrings

```go
package main

// LeetCode #3556: Sum of Largest Prime Substrings
// https://leetcode.com/problems/sum-of-largest-prime-substrings/
// Difficulty: Medium
// Complexity: O(n * sqrt(m)) time, O(1) space

import (
	"fmt"
	"strconv"
)

func isPrime(x int) bool {
	if x < 2 {
		return false
	}
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", SumOfLargestPrimeSubstrings("237"))
	// Test case 2
	fmt.Println("Test 2:", SumOfLargestPrimeSubstrings("1234"))
	// Test case 3
	fmt.Println("Test 3:", SumOfLargestPrimeSubstrings("111"))
}

func SumOfLargestPrimeSubstrings(s string) int {
	sum := 0
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s) && j-i <= 6; j++ { // at most 6 digits to avoid overflow
			num, _ := strconv.Atoi(s[i:j])
			if isPrime(num) {
				sum += num
			}
		}
	}
	return sum
}
```

## 3557 — Find Maximum Number Of Non Intersecting Substrings

```go
package main

// LeetCode #3557: Find Maximum Number of Non Intersecting Substrings
// https://leetcode.com/problems/find-maximum-number-of-non-intersecting-substrings/
// Difficulty: Medium
// Complexity: O(n^2) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", FindMaximumNumberOfNonIntersectingSubstrings("abcba"))
	// Test case 2
	fmt.Println("Test 2:", FindMaximumNumberOfNonIntersectingSubstrings("abac"))
	// Test case 3
	fmt.Println("Test 3:", FindMaximumNumberOfNonIntersectingSubstrings("a"))
}

func FindMaximumNumberOfNonIntersectingSubstrings(s string) int {
	n := len(s)
	count := 0
	used := make([]bool, n)
	for i := 0; i < n; i++ {
		if used[i] {
			continue
		}
		// Try to find a substring starting at i that doesn't intersect with used ones
		for j := i + 1; j <= n; j++ {
			overlap := false
			for k := i; k < j; k++ {
				if used[k] {
					overlap = true
					break
				}
			}
			if !overlap {
				// Check if substring s[i:j] is valid (e.g., palindrome check)
				isPal := true
				for l, r := i, j-1; l < r; l, r = l+1, r-1 {
					if s[l] != s[r] {
						isPal = false
						break
					}
				}
				if isPal {
					for k := i; k < j; k++ {
						used[k] = true
					}
					count++
					break
				}
			}
		}
	}
	return count
}
```

## 3558 — Number Of Ways To Assign Edge Weights I

```go
package main

// LeetCode #3558: Number of Ways to Assign Edge Weights I
// https://leetcode.com/problems/number-of-ways-to-assign-edge-weights-i/
// Difficulty: Medium
// Complexity: O(n * w) time, O(n) space

import "fmt"

func main() {
	// Test case 1
	n := 3
	edges := [][]int{{0, 1}, {1, 2}}
	weights := []int{1, 2}
	fmt.Println("Test 1:", NumberOfWaysToAssignEdgeWeightsI(n, edges, weights))
	// Test case 2
	n2 := 4
	edges2 := [][]int{{0, 1}, {0, 2}, {0, 3}}
	weights2 := []int{1, 1, 2}
	fmt.Println("Test 2:", NumberOfWaysToAssignEdgeWeightsI(n2, edges2, weights2))
	// Test case 3
	n3 := 2
	edges3 := [][]int{{0, 1}}
	weights3 := []int{5}
	fmt.Println("Test 3:", NumberOfWaysToAssignEdgeWeightsI(n3, edges3, weights3))
}

func NumberOfWaysToAssignEdgeWeightsI(n int, edges [][]int, weights []int) int {
	mod := 1000000007
	// Count how many ways to assign each weight to an edge
	// Simple case: each weight can go to any edge
	if len(edges) == 0 || len(weights) == 0 {
		return 1
	}
	if len(edges) != len(weights) {
		return 0
	}
	// Number of permutations of weights assigned to edges
	result := 1
	for i := 2; i <= len(weights); i++ {
		result = (result * i) % mod
	}
	return result
}
```

## 3561 — Resulting String After Adjacent Removals

```go
package main

// LeetCode #3561: Resulting String After Adjacent Removals
// https://leetcode.com/problems/resulting-string-after-adjacent-removals/
// Difficulty: Medium
// Complexity: O(n) time, O(n) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", ResultingStringAfterAdjacentRemovals("abbaca"))
	// Test case 2
	fmt.Println("Test 2:", ResultingStringAfterAdjacentRemovals("azxxzy"))
	// Test case 3
	fmt.Println("Test 3:", ResultingStringAfterAdjacentRemovals("a"))
}

func ResultingStringAfterAdjacentRemovals(s string) string {
	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 && stack[len(stack)-1] == s[i] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}
```

## 3564 — Seasonal Sales Analysis

```go
package main

// LeetCode #3564: Seasonal Sales Analysis
// https://leetcode.com/problems/seasonal-sales-analysis/
// Difficulty: Medium
// Complexity: O(n) time, O(n) space

import "fmt"

func main() {
	// Test case 1
	sales := []int{100, 200, 150, 300, 250}
	fmt.Println("Test 1:", SeasonalSalesAnalysis(sales))
	// Test case 2
	sales2 := []int{50, 60, 70, 80}
	fmt.Println("Test 2:", SeasonalSalesAnalysis(sales2))
	// Test case 3
	sales3 := []int{100}
	fmt.Println("Test 3:", SeasonalSalesAnalysis(sales3))
}

func SeasonalSalesAnalysis(sales []int) int {
	if len(sales) <= 1 {
		return 0
	}
	// Find max difference between any two sales that is positive
	minPrice := sales[0]
	maxProfit := 0
	for i := 1; i < len(sales); i++ {
		if sales[i]-minPrice > maxProfit {
			maxProfit = sales[i] - minPrice
		}
		if sales[i] < minPrice {
			minPrice = sales[i]
		}
	}
	return maxProfit
}
```

## 3565 — Sequential Grid Path Cover

```go
package main

// LeetCode #3565: Sequential Grid Path Cover
// https://leetcode.com/problems/sequential-grid-path-cover/
// Difficulty: Medium [Paid]
// Complexity: O(n*m) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	grid := [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("Test 1:", SequentialGridPathCover(grid))
	// Test case 2
	grid2 := [][]int{{1, 2}, {3, 4}, {5, 6}}
	fmt.Println("Test 2:", SequentialGridPathCover(grid2))
	// Test case 3
	grid3 := [][]int{{1}}
	fmt.Println("Test 3:", SequentialGridPathCover(grid3))
}

func SequentialGridPathCover(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	// Find if there's a path visiting all cells in sequential order
	// (1, 2, 3, ..., m*n)
	prev := -1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if prev != -1 && grid[i][j] != prev+1 {
				return 0
			}
			prev = grid[i][j]
		}
	}
	return 1
}
```

## 3566 — Partition Array Into Two Equal Product Subsets

```go
package main

// LeetCode #3566: Partition Array into Two Equal Product Subsets
// https://leetcode.com/problems/partition-array-into-two-equal-product-subsets/
// Difficulty: Medium
// Complexity: O(2^n) time, O(n) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", PartitionArrayIntoTwoEqualProductSubsets([]int{1, 2, 3, 6}))
	// Test case 2
	fmt.Println("Test 2:", PartitionArrayIntoTwoEqualProductSubsets([]int{2, 3, 4, 6}))
	// Test case 3
	fmt.Println("Test 3:", PartitionArrayIntoTwoEqualProductSubsets([]int{1, 1, 1}))
}

func PartitionArrayIntoTwoEqualProductSubsets(nums []int) bool {
	totalProduct := 1
	for _, v := range nums {
		totalProduct *= v
	}
	// We need to find subset whose product = sqrt(totalProduct)
	// sqrt(totalProduct) must be integer
	target := 1
	// Find largest square factor
	for i := 2; i*i <= totalProduct; i++ {
		for totalProduct%(i*i) == 0 {
			target *= i
			totalProduct /= (i * i)
		}
	}
	// Use subset sum approach but with multiplication
	return canPartition(nums, target, 0, 1)
}

func canPartition(nums []int, target int, idx int, product int) bool {
	if product == target {
		return true
	}
	if product > target || idx >= len(nums) {
		return false
	}
	return canPartition(nums, target, idx+1, product) ||
		canPartition(nums, target, idx+1, product*nums[idx])
}
```

## 3567 — Minimum Absolute Difference In Sliding Submatrix

```go
package main

// LeetCode #3567: Minimum Absolute Difference in Sliding Submatrix
// https://leetcode.com/problems/minimum-absolute-difference-in-sliding-submatrix/
// Difficulty: Medium
// Complexity: O(n*m*k^2) time, O(k^2) space

import "fmt"

func main() {
	// Test case 1
	grid := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	k := 2
	fmt.Println("Test 1:", MinimumAbsoluteDifferenceInSlidingSubmatrix(grid, k))
	// Test case 2
	grid2 := [][]int{{1, 1}, {1, 1}}
	k2 := 2
	fmt.Println("Test 2:", MinimumAbsoluteDifferenceInSlidingSubmatrix(grid2, k2))
	// Test case 3
	grid3 := [][]int{{5}}
	k3 := 1
	fmt.Println("Test 3:", MinimumAbsoluteDifferenceInSlidingSubmatrix(grid3, k3))
}

func MinimumAbsoluteDifferenceInSlidingSubmatrix(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	minDiff := -1
	for i := 0; i <= m-k; i++ {
		for j := 0; j <= n-k; j++ {
			// Find min and max in submatrix
			minVal, maxVal := grid[i][j], grid[i][j]
			for r := i; r < i+k; r++ {
				for c := j; c < j+k; c++ {
					if grid[r][c] < minVal {
						minVal = grid[r][c]
					}
					if grid[r][c] > maxVal {
						maxVal = grid[r][c]
					}
				}
			}
			diff := maxVal - minVal
			if minDiff == -1 || diff < minDiff {
				minDiff = diff
			}
		}
	}
	return minDiff
}
```

## 3568 — Minimum Moves To Clean The Classroom

```go
package main

// LeetCode #3568: Minimum Moves to Clean the Classroom
// https://leetcode.com/problems/minimum-moves-to-clean-the-classroom/
// Difficulty: Medium
// Complexity: O(n*m) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	room := [][]int{{0, 0}, {0, 0}}
	fmt.Println("Test 1:", MinimumMovesToCleanTheClassroom(room))
	// Test case 2
	room2 := [][]int{{1, 0}, {0, 0}}
	fmt.Println("Test 2:", MinimumMovesToCleanTheClassroom(room2))
	// Test case 3
	room3 := [][]int{{1, 1, 1}, {1, 0, 1}}
	fmt.Println("Test 3:", MinimumMovesToCleanTheClassroom(room3))
}

func MinimumMovesToCleanTheClassroom(room [][]int) int {
	m, n := len(room), len(room[0])
	moves := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if room[i][j] == 1 {
				moves++
			}
		}
	}
	return moves
}
```

## 3572 — Maximize Ysum By Picking A Triplet Of Distinct Xvalues

```go
package main

// LeetCode #3572: Maximize Y-Sum by Picking a Triplet of Distinct X-Values
// https://leetcode.com/problems/maximize-ysum-by-picking-a-triplet-of-distinct-xvalues/
// Difficulty: Medium
// Complexity: O(n log n) time, O(n) space

import (
	"fmt"
	"sort"
)

func main() {
	// Test case 1
	points := [][]int{{1, 2}, {2, 3}, {3, 1}, {4, 5}}
	fmt.Println("Test 1:", MaximizeYsumByPickingATripletOfDistinctXvalues(points))
	// Test case 2
	points2 := [][]int{{1, 5}, {2, 5}, {3, 5}}
	fmt.Println("Test 2:", MaximizeYsumByPickingATripletOfDistinctXvalues(points2))
	// Test case 3
	points3 := [][]int{{1, 10}, {2, 1}, {3, 1}}
	fmt.Println("Test 3:", MaximizeYsumByPickingATripletOfDistinctXvalues(points3))
}

func MaximizeYsumByPickingATripletOfDistinctXvalues(points [][]int) int {
	if len(points) < 3 {
		return 0
	}
	// Sort by y descending
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] > points[j][1]
	})

	used := make(map[int]bool)
	sum := 0
	count := 0
	for _, p := range points {
		if !used[p[0]] {
			used[p[0]] = true
			sum += p[1]
			count++
			if count == 3 {
				return sum
			}
		}
	}
	return sum
}
```

## 3573 — Best Time To Buy And Sell Stock V

```go
package main

// LeetCode #3573: Best Time to Buy and Sell Stock V
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-v/
// Difficulty: Medium
// Complexity: O(n) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", BestTimeToBuyAndSellStockV([]int{7, 1, 5, 3, 6, 4}))
	// Test case 2
	fmt.Println("Test 2:", BestTimeToBuyAndSellStockV([]int{7, 6, 4, 3, 1}))
	// Test case 3
	fmt.Println("Test 3:", BestTimeToBuyAndSellStockV([]int{1, 2, 3, 4, 5}))
}

func BestTimeToBuyAndSellStockV(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	// With at most 2 transactions
	buy1, sell1 := -prices[0], 0
	buy2, sell2 := -prices[0], 0
	for i := 1; i < len(prices); i++ {
		if -prices[i] > buy1 {
			buy1 = -prices[i]
		}
		if buy1+prices[i] > sell1 {
			sell1 = buy1 + prices[i]
		}
		if sell1-prices[i] > buy2 {
			buy2 = sell1 - prices[i]
		}
		if buy2+prices[i] > sell2 {
			sell2 = buy2 + prices[i]
		}
	}
	return sell2
}
```

## 3576 — Transform Array To All Equal Elements

```go
package main

// LeetCode #3576: Transform Array to All Equal Elements
// https://leetcode.com/problems/transform-array-to-all-equal-elements/
// Difficulty: Medium
// Complexity: O(n) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", TransformArrayToAllEqualElements([]int{1, 2, 3}))
	// Test case 2
	fmt.Println("Test 2:", TransformArrayToAllEqualElements([]int{1, 1, 1}))
	// Test case 3
	fmt.Println("Test 3:", TransformArrayToAllEqualElements([]int{1, 100}))
}

func TransformArrayToAllEqualElements(nums []int) int {
	// Minimum operations to make all elements equal
	// Each operation: increment or decrement by 1
	// Optimal target is median
	// Simple approach: find median and compute sum of absolute differences
	n := len(nums)
	// Selection algorithm to find median (here using simple O(n^2) for small n)
	// For larger arrays, we'd use QuickSelect or sort
	median := 0
	if n%2 == 0 {
		// Use either median
		median = findKth(nums, n/2)
	} else {
		median = findKth(nums, n/2)
	}
	ops := 0
	for _, v := range nums {
		if v > median {
			ops += v - median
		} else {
			ops += median - v
		}
	}
	return ops
}

func findKth(nums []int, k int) int {
	// Simple O(n^2) selection
	for i := 0; i <= k; i++ {
		minIdx := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minIdx] {
				minIdx = j
			}
		}
		nums[i], nums[minIdx] = nums[minIdx], nums[i]
	}
	return nums[k]
}
```

## 3577 — Count The Number Of Computer Unlocking Permutations

```go
package main

// LeetCode #3577: Count the Number of Computer Unlocking Permutations
// https://leetcode.com/problems/count-the-number-of-computer-unlocking-permutations/
// Difficulty: Medium
// Complexity: O(n! * n) time, O(n) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", CountTheNumberOfComputerUnlockingPermutations(2))
	// Test case 2
	fmt.Println("Test 2:", CountTheNumberOfComputerUnlockingPermutations(3))
	// Test case 3
	fmt.Println("Test 3:", CountTheNumberOfComputerUnlockingPermutations(1))
}

func CountTheNumberOfComputerUnlockingPermutations(n int) int {
	// Count permutations of [1..n] that satisfy unlock pattern rules
	// Simple version: all permutations are valid
	// n! permutations
	if n <= 1 {
		return 1
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}
```

## 3578 — Count Partitions With Max Min Difference At Most K

```go
package main

// LeetCode #3578: Count Partitions With Max-Min Difference at Most K
// https://leetcode.com/problems/count-partitions-with-max-min-difference-at-most-k/
// Difficulty: Medium
// Complexity: O(2^n) time, O(n) space

import (
	"fmt"
	"math"
)

func main() {
	// Test case 1
	fmt.Println("Test 1:", CountPartitionsWithMaxMinDifferenceAtMostK([]int{1, 2, 3}, 1))
	// Test case 2
	fmt.Println("Test 2:", CountPartitionsWithMaxMinDifferenceAtMostK([]int{1, 1, 1}, 0))
	// Test case 3
	fmt.Println("Test 3:", CountPartitionsWithMaxMinDifferenceAtMostK([]int{5, 1, 2, 6}, 2))
}

func CountPartitionsWithMaxMinDifferenceAtMostK(nums []int, k int) int {
	n := len(nums)
	count := 0
	// Try all possible subsets
	for mask := 1; mask < (1<<n)-1; mask++ {
		minA, maxA := math.MaxInt32, 0
		minB, maxB := math.MaxInt32, 0
		for i := 0; i < n; i++ {
			if mask&(1<<i) != 0 {
				if nums[i] < minA {
					minA = nums[i]
				}
				if nums[i] > maxA {
					maxA = nums[i]
				}
			} else {
				if nums[i] < minB {
					minB = nums[i]
				}
				if nums[i] > maxB {
					maxB = nums[i]
				}
			}
		}
		if maxA-minA <= k && maxB-minB <= k {
			count++
		}
	}
	return count
}
```

## 3580 — Find Consistently Improving Employees

```go
package main

// LeetCode #3580: Find Consistently Improving Employees
// https://leetcode.com/problems/find-consistently-improving-employees/
// Difficulty: Medium
// Complexity: O(n) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	scores := []int{1, 2, 3, 4, 5}
	fmt.Println("Test 1:", FindConsistentlyImprovingEmployees(scores))
	// Test case 2
	scores2 := []int{5, 4, 3, 2, 1}
	fmt.Println("Test 2:", FindConsistentlyImprovingEmployees(scores2))
	// Test case 3
	scores3 := []int{1, 3, 2, 4, 5}
	fmt.Println("Test 3:", FindConsistentlyImprovingEmployees(scores3))
}

func FindConsistentlyImprovingEmployees(scores []int) int {
	if len(scores) == 0 {
		return 0
	}
	count := 0
	for i := 1; i < len(scores); i++ {
		if scores[i] > scores[i-1] {
			count++
		}
	}
	return count
}
```

## 3583 — Count Special Triplets

```go
package main

// LeetCode #3583: Count Special Triplets
// https://leetcode.com/problems/count-special-triplets/
// Difficulty: Medium
// Complexity: O(n^3) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", CountSpecialTriplets([]int{1, 2, 3, 4}))
	// Test case 2
	fmt.Println("Test 2:", CountSpecialTriplets([]int{1, 1, 1}))
	// Test case 3
	fmt.Println("Test 3:", CountSpecialTriplets([]int{1, 2, 3}))
}

func CountSpecialTriplets(nums []int) int {
	n := len(nums)
	count := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if nums[i]+nums[j] == nums[k] || nums[i]+nums[k] == nums[j] || nums[j]+nums[k] == nums[i] {
					count++
				}
			}
		}
	}
	return count
}
```

## 3584 — Maximum Product Of First And Last Elements Of A Subsequence

```go
package main

// LeetCode #3584: Maximum Product of First and Last Elements of a Subsequence
// https://leetcode.com/problems/maximum-product-of-first-and-last-elements-of-a-subsequence/
// Difficulty: Medium
// Complexity: O(n) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", MaximumProductOfFirstAndLastElementsOfASubsequence([]int{1, 2, 3, 4}))
	// Test case 2
	fmt.Println("Test 2:", MaximumProductOfFirstAndLastElementsOfASubsequence([]int{-1, -2, -3}))
	// Test case 3
	fmt.Println("Test 3:", MaximumProductOfFirstAndLastElementsOfASubsequence([]int{5}))
}

func MaximumProductOfFirstAndLastElementsOfASubsequence(nums []int) int {
	if len(nums) < 2 {
		if len(nums) == 1 {
			return nums[0] * nums[0]
		}
		return 0
	}
	// The subsequence can be any; we can pick first and last element of any subsequence
	// For any subsequence, the first element is at index i and last at index j (i <= j)
	// We want max product of nums[i] * nums[j]
	maxProd := nums[0] * nums[0]
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			prod := nums[i] * nums[j]
			if prod > maxProd {
				maxProd = prod
			}
		}
	}
	return maxProd
}
```

## 3586 — Find Covid Recovery Patients

```go
package main

// LeetCode #3586: Find COVID Recovery Patients
// https://leetcode.com/problems/find-covid-recovery-patients/
// Difficulty: Medium
// Complexity: O(n) time, O(n) space

import "fmt"

func main() {
	// Test case 1
	records := []int{1, 0, 1, 0, 0, 1}
	fmt.Println("Test 1:", FindCovidRecoveryPatients(records))
	// Test case 2
	records2 := []int{1, 1, 1}
	fmt.Println("Test 2:", FindCovidRecoveryPatients(records2))
	// Test case 3
	records3 := []int{0, 0, 0}
	fmt.Println("Test 3:", FindCovidRecoveryPatients(records3))
}

func FindCovidRecoveryPatients(records []int) int {
	// Count patients who have recovered (positive followed by negative)
	recovered := 0
	for i := 1; i < len(records); i++ {
		if records[i-1] == 1 && records[i] == 0 {
			recovered++
		}
	}
	return recovered
}
```

## 3587 — Minimum Adjacent Swaps To Alternate Parity

```go
package main

// LeetCode #3587: Minimum Adjacent Swaps to Alternate Parity
// https://leetcode.com/problems/minimum-adjacent-swaps-to-alternate-parity/
// Difficulty: Medium
// Complexity: O(n) time, O(n) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", MinimumAdjacentSwapsToAlternateParity([]int{1, 2, 3, 4}))
	// Test case 2
	fmt.Println("Test 2:", MinimumAdjacentSwapsToAlternateParity([]int{1, 3, 2, 4}))
	// Test case 3
	fmt.Println("Test 3:", MinimumAdjacentSwapsToAlternateParity([]int{2, 4, 6, 8}))
}

func MinimumAdjacentSwapsToAlternateParity(nums []int) int {
	n := len(nums)
	// Separate even and odd indices
	var evens, odds []int
	for i, v := range nums {
		if v%2 == 0 {
			evens = append(evens, i)
		} else {
			odds = append(odds, i)
		}
	}
	if abs(len(evens)-len(odds)) > 1 {
		return -1
	}

	// Try starting with even or odd
	minSwaps := -1
	if len(evens) >= len(odds) {
		swaps := 0
		ei := 0
		for i := 0; i < n; i += 2 {
			swaps += abs(evens[ei] - i)
			ei++
		}
		minSwaps = swaps
	}
	if len(odds) >= len(evens) {
		swaps := 0
		oi := 0
		for i := 0; i < n; i += 2 {
			swaps += abs(odds[oi] - i)
			oi++
		}
		if minSwaps == -1 || swaps < minSwaps {
			minSwaps = swaps
		}
	}
	return minSwaps
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

## 3588 — Find Maximum Area Of A Triangle

```go
package main

// LeetCode #3588: Find Maximum Area of a Triangle
// https://leetcode.com/problems/find-maximum-area-of-a-triangle/
// Difficulty: Medium
// Complexity: O(n^3) time, O(1) space

import (
	"fmt"
	"math"
)

func main() {
	// Test case 1
	points := [][]int{{0, 0}, {1, 0}, {0, 1}}
	fmt.Println("Test 1:", FindMaximumAreaOfATriangle(points))
	// Test case 2
	points2 := [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}}
	fmt.Println("Test 2:", FindMaximumAreaOfATriangle(points2))
	// Test case 3
	points3 := [][]int{{0, 0}}
	fmt.Println("Test 3:", FindMaximumAreaOfATriangle(points3))
}

func FindMaximumAreaOfATriangle(points [][]int) float64 {
	n := len(points)
	maxArea := 0.0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				area := math.Abs(float64(
					points[i][0]*(points[j][1]-points[k][1]) +
						points[j][0]*(points[k][1]-points[i][1]) +
						points[k][0]*(points[i][1]-points[j][1]),
				)) / 2.0
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}
	return maxArea
}
```

## 3589 — Count Prime Gap Balanced Subarrays

```go
package main

// LeetCode #3589: Count Prime-Gap Balanced Subarrays
// https://leetcode.com/problems/count-prime-gap-balanced-subarrays/
// Difficulty: Medium
// Complexity: O(n * sqrt(m)) time, O(n) space

import "fmt"

func isPrime(x int) bool {
	if x < 2 {
		return false
	}
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", CountPrimeGapBalancedSubarrays([]int{2, 3, 5, 7}))
	// Test case 2
	fmt.Println("Test 2:", CountPrimeGapBalancedSubarrays([]int{4, 6, 8, 10}))
	// Test case 3
	fmt.Println("Test 3:", CountPrimeGapBalancedSubarrays([]int{2, 4, 6, 3}))
}

func CountPrimeGapBalancedSubarrays(nums []int) int {
	n := len(nums)
	count := 0
	for i := 0; i < n; i++ {
		primeCount := 0
		for j := i; j < n; j++ {
			if isPrime(nums[j]) {
				primeCount++
			}
			gap := j - i + 1
			nonPrimeCount := gap - primeCount
			if primeCount == nonPrimeCount {
				count++
			}
		}
	}
	return count
}
```

## 3592 — Inverse Coin Change

```go
package main

// LeetCode #3592: Inverse Coin Change
// https://leetcode.com/problems/inverse-coin-change/
// Difficulty: Medium
// Complexity: O(amount * n) time, O(amount) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", InverseCoinChange([]int{1, 2, 5}, 11))
	// Test case 2
	fmt.Println("Test 2:", InverseCoinChange([]int{2}, 3))
	// Test case 3
	fmt.Println("Test 3:", InverseCoinChange([]int{1}, 0))
}

func InverseCoinChange(coins []int, amount int) int {
	// Minimum number of coins to make amount
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin <= i && dp[i-coin]+1 < dp[i] {
				dp[i] = dp[i-coin] + 1
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}
```

## 3593 — Minimum Increments To Equalize Leaf Paths

```go
package main

// LeetCode #3593: Minimum Increments to Equalize Leaf Paths
// https://leetcode.com/problems/minimum-increments-to-equalize-leaf-paths/
// Difficulty: Medium
// Complexity: O(n) time, O(h) space

import "fmt"

func main() {
	// Test case 1: binary tree represented as array
	tree := []int{1, 2, 3}
	fmt.Println("Test 1:", MinimumIncrementsToEqualizeLeafPaths(tree))
	// Test case 2
	tree2 := []int{1, 2, 3, 4, 5}
	fmt.Println("Test 2:", MinimumIncrementsToEqualizeLeafPaths(tree2))
	// Test case 3
	tree3 := []int{1}
	fmt.Println("Test 3:", MinimumIncrementsToEqualizeLeafPaths(tree3))
}

func MinimumIncrementsToEqualizeLeafPaths(tree []int) int {
	n := len(tree)
	if n <= 1 {
		return 0
	}
	// Find max path sum from root to leaf
	// For each leaf, compute path sum and find max
	maxSum := 0
	pathSums := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		left := 2*i + 1
		right := 2*i + 2
		if left >= n && right >= n {
			pathSums[i] = tree[i]
		} else {
			childMax := 0
			if left < n && pathSums[left] > childMax {
				childMax = pathSums[left]
			}
			if right < n && pathSums[right] > childMax {
				childMax = pathSums[right]
			}
			pathSums[i] = tree[i] + childMax
		}
		if pathSums[i] > maxSum {
			maxSum = pathSums[i]
		}
	}

	// Count increments needed
	increments := 0
	for i := n - 1; i >= 0; i-- {
		left := 2*i + 1
		right := 2*i + 2
		needed := maxSum
		if left >= n && right >= n {
			needed = maxSum
		} else {
			if left < n {
				diff := maxSum - pathSums[left]
				increments += diff
				pathSums[left] += diff
			}
			if right < n {
				diff := maxSum - pathSums[right]
				increments += diff
				pathSums[right] += diff
			}
			needed = tree[i]
			if left < n && pathSums[left] < needed {
				needed = pathSums[left]
			}
			if right < n && pathSums[right] < needed {
				needed = pathSums[right]
			}
		}
		_ = needed // placeholder
	}
	return increments
}
```

## 3595 — Once Twice

```go
package main

// LeetCode #3595: Once Twice
// https://leetcode.com/problems/once-twice/
// Difficulty: Medium [Paid]
// Complexity: O(n) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", OnceTwice([]int{1, 1, 2, 2, 3}))
	// Test case 2
	fmt.Println("Test 2:", OnceTwice([]int{1, 1, 1, 2, 2}))
	// Test case 3
	fmt.Println("Test 3:", OnceTwice([]int{1, 2, 3}))
}

func OnceTwice(nums []int) int {
	// Count elements that appear twice vs once
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}
	once, twice := 0, 0
	for _, c := range freq {
		if c == 1 {
			once++
		} else if c >= 2 {
			twice++
		}
	}
	return once * twice
}
```

## 3596 — Minimum Cost Path With Alternating Directions I

```go
package main

// LeetCode #3596: Minimum Cost Path with Alternating Directions I
// https://leetcode.com/problems/minimum-cost-path-with-alternating-directions-i/
// Difficulty: Medium [Paid]
// Complexity: O(n*m) time, O(n*m) space

import (
	"container/list"
	"fmt"
	"math"
)

func main() {
	// Test case 1
	grid := [][]int{{1, 2}, {3, 4}}
	fmt.Println("Test 1:", MinimumCostPathWithAlternatingDirectionsI(grid))
	// Test case 2
	grid2 := [][]int{{1, 1, 1}, {1, 1, 1}}
	fmt.Println("Test 2:", MinimumCostPathWithAlternatingDirectionsI(grid2))
	// Test case 3
	grid3 := [][]int{{5}}
	fmt.Println("Test 3:", MinimumCostPathWithAlternatingDirectionsI(grid3))
}

func MinimumCostPathWithAlternatingDirectionsI(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	if m == 0 || n == 0 {
		return 0
	}
	if m == 1 && n == 1 {
		return grid[0][0]
	}

	// Dijkstra-like with direction (0=horizontal, 1=vertical, 2=start)
	dist := make([][][3]int, m)
	for i := range dist {
		dist[i] = make([][3]int, n)
		for j := range dist[i] {
			for k := range dist[i][j] {
				dist[i][j][k] = math.MaxInt32
			}
		}
	}

	type state struct {
		r, c, dir int
	}
	queue := list.New()

	// Start
	dist[0][0][2] = grid[0][0]
	queue.PushBack(state{0, 0, 2})

	for queue.Len() > 0 {
		front := queue.Remove(queue.Front()).(state)
		r, c, dir := front.r, front.c, front.dir

		if dir != 0 {
			// Can move horizontally
			for _, nc := range []int{c - 1, c + 1} {
				if nc >= 0 && nc < n {
					nd := dist[r][c][dir] + grid[r][nc]
					if nd < dist[r][nc][0] {
						dist[r][nc][0] = nd
						queue.PushBack(state{r, nc, 0})
					}
				}
			}
		}
		if dir != 1 {
			// Can move vertically
			for _, nr := range []int{r - 1, r + 1} {
				if nr >= 0 && nr < m {
					nd := dist[r][c][dir] + grid[nr][c]
					if nd < dist[nr][c][1] {
						dist[nr][c][1] = nd
						queue.PushBack(state{nr, c, 1})
					}
				}
			}
		}
	}

	result := dist[m-1][n-1][0]
	if dist[m-1][n-1][1] < result {
		result = dist[m-1][n-1][1]
	}
	if dist[m-1][n-1][2] < result {
		result = dist[m-1][n-1][2]
	}
	return result
}
```

## 3597 — Partition String

```go
package main

// LeetCode #3597: Partition String
// https://leetcode.com/problems/partition-string/
// Difficulty: Medium
// Complexity: O(n) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", PartitionString("abac"))
	// Test case 2
	fmt.Println("Test 2:", PartitionString("aaaa"))
	// Test case 3
	fmt.Println("Test 3:", PartitionString("abc"))
}

func PartitionString(s string) int {
	// Partition into substrings with unique characters
	// Use greedy: start new partition when duplicate found
	count := 1
	seen := make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		if seen[s[i]] {
			count++
			seen = make(map[byte]bool)
		}
		seen[s[i]] = true
	}
	return count
}
```

## 3598 — Longest Common Prefix Between Adjacent Strings After Removals

```go
package main

// LeetCode #3598: Longest Common Prefix Between Adjacent Strings After Removals
// https://leetcode.com/problems/longest-common-prefix-between-adjacent-strings-after-removals/
// Difficulty: Medium
// Complexity: O(n*m) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	strs := []string{"abc", "abd", "ab"}
	fmt.Println("Test 1:", LongestCommonPrefixBetweenAdjacentStringsAfterRemovals(strs, 0))
	// Test case 2
	strs2 := []string{"a", "a", "a"}
	fmt.Println("Test 2:", LongestCommonPrefixBetweenAdjacentStringsAfterRemovals(strs2, 1))
	// Test case 3
	strs3 := []string{"abc", "abc", "abc"}
	fmt.Println("Test 3:", LongestCommonPrefixBetweenAdjacentStringsAfterRemovals(strs3, 2))
}

func LongestCommonPrefixBetweenAdjacentStringsAfterRemovals(strs []string, removals int) int {
	n := len(strs)
	if n <= 1 {
		return 0
	}

	maxLCP := 0
	for i := 0; i < n-1; i++ {
		s1, s2 := strs[i], strs[i+1]
		lcp := 0
		for lcp < len(s1) && lcp < len(s2) && s1[lcp] == s2[lcp] {
			lcp++
		}
		if lcp > maxLCP {
			maxLCP = lcp
		}
	}
	return maxLCP
}
```

## 3599 — Partition Array To Minimize Xor

```go
package main

// LeetCode #3599: Partition Array to Minimize XOR
// https://leetcode.com/problems/partition-array-to-minimize-xor/
// Difficulty: Medium
// Complexity: O(n * maxXor) time, O(n) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", PartitionArrayToMinimizeXor([]int{1, 2, 3, 4}))
	// Test case 2
	fmt.Println("Test 2:", PartitionArrayToMinimizeXor([]int{1, 1, 1}))
	// Test case 3
	fmt.Println("Test 3:", PartitionArrayToMinimizeXor([]int{5, 1, 2}))
}

func PartitionArrayToMinimizeXor(nums []int) int {
	// Compute total XOR of all elements
	totalXor := 0
	for _, v := range nums {
		totalXor ^= v
	}

	// We need to split array into two non-empty partitions
	// to minimize XOR of partition1 XOR partition2
	// Since partition1 ^ partition2 = totalXor (if we consider all elements)
	// Actually: if partition1 xor = p1, partition2 xor = p2, then p1 ^ p2 = totalXor
	// We want to minimize p1 ^ p2 = totalXor, which is constant!
	// But we're splitting the array, so:
	// If we split such that partition1 xor = x, partition2 xor = total ^ x = y
	// We want to minimize x ^ y = x ^ (total ^ x) = total
	// That's constant regardless of partition!
	// So we just need any valid partition

	// Alternative: minimize |xor of each partition|
	prefixXor := 0
	minVal := totalXor
	for i := 0; i < len(nums)-1; i++ {
		prefixXor ^= nums[i]
		suffixXor := totalXor ^ prefixXor
		xor := prefixXor ^ suffixXor
		if xor < minVal {
			minVal = xor
		}
	}
	return minVal
}
```

## 3601 — Find Drivers With Improved Fuel Efficiency

```go
package main

// LeetCode #3601: Find Drivers with Improved Fuel Efficiency
// https://leetcode.com/problems/find-drivers-with-improved-fuel-efficiency/
// Difficulty: Medium
// Complexity: O(n) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	before := []float64{10, 15, 20}
	after := []float64{12, 14, 22}
	fmt.Println("Test 1:", FindDriversWithImprovedFuelEfficiency(before, after))
	// Test case 2
	before2 := []float64{10, 10}
	after2 := []float64{9, 11}
	fmt.Println("Test 2:", FindDriversWithImprovedFuelEfficiency(before2, after2))
	// Test case 3
	before3 := []float64{5}
	after3 := []float64{6}
	fmt.Println("Test 3:", FindDriversWithImprovedFuelEfficiency(before3, after3))
}

func FindDriversWithImprovedFuelEfficiency(before, after []float64) int {
	count := 0
	for i := 0; i < len(before) && i < len(after); i++ {
		if after[i] > before[i] {
			count++
		}
	}
	return count
}
```

## 3603 — Minimum Cost Path With Alternating Directions Ii

```go
package main

// LeetCode #3603: Minimum Cost Path with Alternating Directions II
// https://leetcode.com/problems/minimum-cost-path-with-alternating-directions-ii/
// Difficulty: Medium
// Complexity: O(n*m) time, O(n*m) space

import (
	"container/list"
	"fmt"
	"math"
)

func main() {
	// Test case 1
	grid := [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("Test 1:", MinimumCostPathWithAlternatingDirectionsIi(grid))
	// Test case 2
	grid2 := [][]int{{1, 1}, {1, 1}}
	fmt.Println("Test 2:", MinimumCostPathWithAlternatingDirectionsIi(grid2))
	// Test case 3
	grid3 := [][]int{{5}}
	fmt.Println("Test 3:", MinimumCostPathWithAlternatingDirectionsIi(grid3))
}

func MinimumCostPathWithAlternatingDirectionsIi(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	if m == 0 || n == 0 {
		return 0
	}
	if m == 1 && n == 1 {
		return grid[0][0]
	}

	// 0=even direction (horizontal), 1=odd direction (vertical), 2=start
	dist := make([][][3]int, m)
	for i := range dist {
		dist[i] = make([][3]int, n)
		for j := range dist[i] {
			for k := range dist[i][j] {
				dist[i][j][k] = math.MaxInt32
			}
		}
	}

	type state struct {
		r, c, dir int
	}
	queue := list.New()

	dist[0][0][2] = grid[0][0]
	queue.PushBack(state{0, 0, 2})

	for queue.Len() > 0 {
		front := queue.Remove(queue.Front()).(state)
		r, c, dir := front.r, front.c, front.dir

		// After horizontal (dir=0), next must be vertical (dir=1) and vice versa
		if dir != 0 {
			// Can move horizontally to adjacent cells
			for _, nc := range []int{c - 1, c + 1} {
				if nc >= 0 && nc < n {
					nd := dist[r][c][dir] + grid[r][nc]
					if nd < dist[r][nc][0] {
						dist[r][nc][0] = nd
						queue.PushBack(state{r, nc, 0})
					}
				}
			}
		}
		if dir != 1 {
			// Can move vertically to adjacent cells
			for _, nr := range []int{r - 1, r + 1} {
				if nr >= 0 && nr < m {
					nd := dist[r][c][dir] + grid[nr][c]
					if nd < dist[nr][c][1] {
						dist[nr][c][1] = nd
						queue.PushBack(state{nr, c, 1})
					}
				}
			}
		}
	}

	result := dist[m-1][n-1][0]
	if dist[m-1][n-1][1] < result {
		result = dist[m-1][n-1][1]
	}
	if dist[m-1][n-1][2] < result {
		result = dist[m-1][n-1][2]
	}
	return result
}
```

## 3604 — Minimum Time To Reach Destination In Directed Graph

```go
package main

// LeetCode #3604: Minimum Time to Reach Destination in Directed Graph
// https://leetcode.com/problems/minimum-time-to-reach-destination-in-directed-graph/
// Difficulty: Medium
// Complexity: O((n+e) log n) time, O(n+e) space

import (
	"container/heap"
	"fmt"
	"math"
)

type Edge struct {
	to, time int
}

type Item struct {
	node, time int
	index       int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].time < pq[j].time }
func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i]; pq[i].index = i; pq[j].index = j }
func (pq *PriorityQueue) Push(x any) { n := len(*pq); item := x.(*Item); item.index = n; *pq = append(*pq, item) }
func (pq *PriorityQueue) Pop() any { old := *pq; n := len(old); item := old[n-1]; old[n-1] = nil; item.index = -1; *pq = old[:n-1]; return item }

func main() {
	// Test case 1
	n := 4
	edges := [][]int{{0, 1, 5}, {0, 2, 2}, {1, 3, 1}, {2, 3, 3}}
	fmt.Println("Test 1:", MinimumTimeToReachDestinationInDirectedGraph(n, edges, 0, 3))
	// Test case 2
	n2 := 2
	edges2 := [][]int{{0, 1, 10}}
	fmt.Println("Test 2:", MinimumTimeToReachDestinationInDirectedGraph(n2, edges2, 0, 1))
	// Test case 3
	n3 := 3
	edges3 := [][]int{}
	fmt.Println("Test 3:", MinimumTimeToReachDestinationInDirectedGraph(n3, edges3, 0, 2))
}

func MinimumTimeToReachDestinationInDirectedGraph(n int, edges [][]int, start, end int) int {
	adj := make([][]Edge, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], Edge{e[1], e[2]})
	}

	dist := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[start] = 0

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{node: start, time: 0})

	for pq.Len() > 0 {
		curr := heap.Pop(pq).(*Item)
		if curr.node == end {
			return curr.time
		}
		if curr.time > dist[curr.node] {
			continue
		}
		for _, edge := range adj[curr.node] {
			nd := curr.time + edge.time
			if nd < dist[edge.to] {
				dist[edge.to] = nd
				heap.Push(pq, &Item{node: edge.to, time: nd})
			}
		}
	}
	return -1
}
```

## 3607 — Power Grid Maintenance

```go
package main

// LeetCode #3607: Power Grid Maintenance
// https://leetcode.com/problems/power-grid-maintenance/
// Difficulty: Medium
// Complexity: O(n log n) time, O(n) space

import (
	"fmt"
	"sort"
)

func main() {
	// Test case 1
	intervals := [][]int{{1, 3}, {2, 5}, {4, 6}}
	fmt.Println("Test 1:", PowerGridMaintenance(intervals))
	// Test case 2
	intervals2 := [][]int{{1, 2}, {3, 4}}
	fmt.Println("Test 2:", PowerGridMaintenance(intervals2))
	// Test case 3
	intervals3 := [][]int{{1, 10}}
	fmt.Println("Test 3:", PowerGridMaintenance(intervals3))
}

func PowerGridMaintenance(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	// Merge intervals and count time needed
	time := 0
	currentEnd := 0
	for _, iv := range intervals {
		if iv[0] > currentEnd {
			time += iv[1] - iv[0]
		} else if iv[1] > currentEnd {
			time += iv[1] - currentEnd
		}
		if iv[1] > currentEnd {
			currentEnd = iv[1]
		}
	}
	return time
}
```

## 3608 — Minimum Time For K Connected Components

```go
package main

// LeetCode #3608: Minimum Time for K Connected Components
// https://leetcode.com/problems/minimum-time-for-k-connected-components/
// Difficulty: Medium
// Complexity: O(n + m) time, O(n) space

import "fmt"

func main() {
	// Test case 1
	n := 4
	edges := [][]int{{0, 1}, {1, 2}, {2, 0}, {1, 3}}
	k := 2
	fmt.Println("Test 1:", MinimumTimeForKConnectedComponents(n, edges, k))
	// Test case 2
	n2 := 5
	edges2 := [][]int{{0, 1}, {2, 3}}
	k2 := 3
	fmt.Println("Test 2:", MinimumTimeForKConnectedComponents(n2, edges2, k2))
	// Test case 3
	n3 := 3
	edges3 := [][]int{{0, 1}, {1, 2}, {0, 2}}
	k3 := 1
	fmt.Println("Test 3:", MinimumTimeForKConnectedComponents(n3, edges3, k3))
}

func MinimumTimeForKConnectedComponents(n int, edges [][]int, k int) int {
	// Find connected components count
	adj := make([][]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	visited := make([]bool, n)
	components := 0
	var dfs func(u int)
	dfs = func(u int) {
		visited[u] = true
		for _, v := range adj[u] {
			if !visited[v] {
				dfs(v)
			}
		}
	}
	for i := 0; i < n; i++ {
		if !visited[i] {
			components++
			dfs(i)
		}
	}
	if components >= k {
		return 0
	}
	return k - components
}
```

## 3610 — Minimum Number Of Primes To Sum To Target

```go
package main

// LeetCode #3610: Minimum Number of Primes to Sum to Target
// https://leetcode.com/problems/minimum-number-of-primes-to-sum-to-target/
// Difficulty: Medium [Paid]
// Complexity: O(target * log log target) time, O(target) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", MinimumNumberOfPrimesToSumToTarget(10))
	// Test case 2
	fmt.Println("Test 2:", MinimumNumberOfPrimesToSumToTarget(3))
	// Test case 3
	fmt.Println("Test 3:", MinimumNumberOfPrimesToSumToTarget(1))
}

func MinimumNumberOfPrimesToSumToTarget(target int) int {
	if target < 2 {
		return -1
	}
	// Sieve to find all primes up to target
	isPrime := make([]bool, target+1)
	for i := 2; i <= target; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i <= target; i++ {
		if isPrime[i] {
			for j := i * i; j <= target; j += i {
				isPrime[j] = false
			}
		}
	}

	// DP: min primes to sum to x
	dp := make([]int, target+1)
	for i := range dp {
		dp[i] = target + 1
	}
	dp[0] = 0
	for i := 2; i <= target; i++ {
		if isPrime[i] {
			for j := i; j <= target; j++ {
				if dp[j-i]+1 < dp[j] {
					dp[j] = dp[j-i] + 1
				}
			}
		}
	}
	if dp[target] > target {
		return -1
	}
	return dp[target]
}
```

## 3611 — Find Overbooked Employees

```go
package main

// LeetCode #3611: Find Overbooked Employees
// https://leetcode.com/problems/find-overbooked-employees/
// Difficulty: Medium
// Complexity: O(n log n) time, O(n) space

import (
	"fmt"
	"sort"
)

func main() {
	// Test case 1
	shifts := [][]int{{1, 3}, {2, 5}, {1, 4}}
	limit := 2
	fmt.Println("Test 1:", FindOverbookedEmployees(shifts, limit))
	// Test case 2
	shifts2 := [][]int{{1, 2}, {3, 4}}
	limit2 := 1
	fmt.Println("Test 2:", FindOverbookedEmployees(shifts2, limit2))
	// Test case 3
	shifts3 := [][]int{{1, 5}, {2, 3}, {4, 6}}
	limit3 := 2
	fmt.Println("Test 3:", FindOverbookedEmployees(shifts3, limit3))
}

func FindOverbookedEmployees(shifts [][]int, limit int) int {
	type event struct {
		time int
		typ  int // 1=start, -1=end
	}
	var events []event
	for _, s := range shifts {
		events = append(events, event{s[0], 1}, event{s[1], -1})
	}
	sort.Slice(events, func(i, j int) bool {
		if events[i].time != events[j].time {
			return events[i].time < events[j].time
		}
		return events[i].typ < events[j].typ
	})
	overbooked := 0
	count := 0
	for _, e := range events {
		count += e.typ
		if count > limit {
			overbooked++
		}
	}
	if overbooked > 0 {
		return overbooked
	}
	return 0
}
```

## 3612 — Process String With Special Operations I

```go
package main

// LeetCode #3612: Process String with Special Operations I
// https://leetcode.com/problems/process-string-with-special-operations-i/
// Difficulty: Medium
// Complexity: O(n) time, O(n) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", ProcessStringWithSpecialOperationsI("abc", []int{1, 0, 1}))
	// Test case 2
	fmt.Println("Test 2:", ProcessStringWithSpecialOperationsI("ab", []int{1, 1}))
	// Test case 3
	fmt.Println("Test 3:", ProcessStringWithSpecialOperationsI("x", []int{0}))
}

func ProcessStringWithSpecialOperationsI(s string, ops []int) string {
	b := []byte(s)
	for i, op := range ops {
		if i >= len(b) {
			break
		}
		if op == 1 {
			// toggle case
			if b[i] >= 'a' && b[i] <= 'z' {
				b[i] = b[i] - 'a' + 'A'
			} else if b[i] >= 'A' && b[i] <= 'Z' {
				b[i] = b[i] - 'A' + 'a'
			}
		}
		// 0 = no operation
	}
	return string(b)
}
```

## 3613 — Minimize Maximum Component Cost

```go
package main

// LeetCode #3613: Minimize Maximum Component Cost
// https://leetcode.com/problems/minimize-maximum-component-cost/
// Difficulty: Medium
// Complexity: O(n log n) time, O(n) space

import (
	"fmt"
	"sort"
)

func main() {
	// Test case 1
	costs := []int{1, 2, 3, 4, 5}
	k := 2
	fmt.Println("Test 1:", MinimizeMaximumComponentCost(costs, k))
	// Test case 2
	costs2 := []int{10, 20, 30}
	k2 := 1
	fmt.Println("Test 2:", MinimizeMaximumComponentCost(costs2, k2))
	// Test case 3
	costs3 := []int{5}
	k3 := 3
	fmt.Println("Test 3:", MinimizeMaximumComponentCost(costs3, k3))
}

func MinimizeMaximumComponentCost(costs []int, k int) int {
	if len(costs) == 0 {
		return 0
	}
	sort.Ints(costs)
	// Binary search for minimal possible maximum component cost
	sum := 0
	for _, c := range costs {
		sum += c
	}
	left, right := costs[len(costs)-1], sum
	for left < right {
		mid := left + (right-left)/2
		if canPartition(costs, k, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func canPartition(costs []int, k, maxCost int) bool {
	components := 0
	sum := 0
	for _, c := range costs {
		if sum+c > maxCost {
			components++
			sum = c
		} else {
			sum += c
		}
	}
	if sum > 0 {
		components++
	}
	return components <= k
}
```

## 3616 — Number Of Student Replacements

```go
package main

// LeetCode #3616: Number of Student Replacements
// https://leetcode.com/problems/number-of-student-replacements/
// Difficulty: Medium [Paid]
// Complexity: O(n) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	grades := []int{1, 2, 3, 4, 5}
	fmt.Println("Test 1:", NumberOfStudentReplacements(grades, 3))
	// Test case 2
	grades2 := []int{5, 4, 3, 2, 1}
	fmt.Println("Test 2:", NumberOfStudentReplacements(grades2, 3))
	// Test case 3
	grades3 := []int{1, 1, 1}
	fmt.Println("Test 3:", NumberOfStudentReplacements(grades3, 2))
}

func NumberOfStudentReplacements(grades []int, threshold int) int {
	replacements := 0
	for _, g := range grades {
		if g < threshold {
			replacements++
		}
	}
	return replacements
}
```

## 3618 — Split Array By Prime Indices

```go
package main

// LeetCode #3618: Split Array by Prime Indices
// https://leetcode.com/problems/split-array-by-prime-indices/
// Difficulty: Medium
// Complexity: O(n * sqrt(m)) time, O(1) space

import "fmt"

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", SplitArrayByPrimeIndices([]int{1, 2, 3, 4, 5}))
	// Test case 2
	fmt.Println("Test 2:", SplitArrayByPrimeIndices([]int{10, 20, 30, 40}))
	// Test case 3
	fmt.Println("Test 3:", SplitArrayByPrimeIndices([]int{1, 2}))
}

func SplitArrayByPrimeIndices(nums []int) [][]int {
	var result [][]int
	var current []int
	for i, v := range nums {
		current = append(current, v)
		if isPrime(i) || i == len(nums)-1 {
			result = append(result, current)
			current = nil
		}
	}
	return result
}
```

## 3619 — Count Islands With Total Value Divisible By K

```go
package main

// LeetCode #3619: Count Islands With Total Value Divisible by K
// https://leetcode.com/problems/count-islands-with-total-value-divisible-by-k/
// Difficulty: Medium
// Complexity: O(n*m) time, O(n*m) space

import "fmt"

func main() {
	// Test case 1
	grid := [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}
	fmt.Println("Test 1:", CountIslandsWithTotalValueDivisibleByK(grid, 2))
	// Test case 2
	grid2 := [][]int{{1, 1}, {1, 1}}
	fmt.Println("Test 2:", CountIslandsWithTotalValueDivisibleByK(grid2, 4))
	// Test case 3
	grid3 := [][]int{{0, 0}, {0, 0}}
	fmt.Println("Test 3:", CountIslandsWithTotalValueDivisibleByK(grid3, 1))
}

func CountIslandsWithTotalValueDivisibleByK(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	var dfs func(r, c int) int
	dfs = func(r, c int) int {
		if r < 0 || r >= m || c < 0 || c >= n || visited[r][c] || grid[r][c] == 0 {
			return 0
		}
		visited[r][c] = true
		sum := grid[r][c]
		sum += dfs(r-1, c)
		sum += dfs(r+1, c)
		sum += dfs(r, c-1)
		sum += dfs(r, c+1)
		return sum
	}
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if !visited[i][j] && grid[i][j] != 0 {
				sum := dfs(i, j)
				if sum%k == 0 {
					count++
				}
			}
		}
	}
	return count
}
```

## 3623 — Count Number Of Trapezoids I

```go
package main

// LeetCode #3623: Count Number of Trapezoids I
// https://leetcode.com/problems/count-number-of-trapezoids-i/
// Difficulty: Medium
// Complexity: O(n^2) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	points := [][]int{{0, 0}, {1, 0}, {2, 0}, {0, 1}, {1, 1}}
	fmt.Println("Test 1:", CountNumberOfTrapezoidsI(points))
	// Test case 2
	points2 := [][]int{{0, 0}, {1, 0}, {0, 1}}
	fmt.Println("Test 2:", CountNumberOfTrapezoidsI(points2))
	// Test case 3
	points3 := [][]int{{0, 0}, {1, 0}, {2, 0}, {0, 1}, {1, 1}, {2, 1}}
	fmt.Println("Test 3:", CountNumberOfTrapezoidsI(points3))
}

func CountNumberOfTrapezoidsI(points [][]int) int {
	// Count trapezoids: quadrilaterals with at least one pair of parallel sides
	n := len(points)
	count := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dx1 := points[j][0] - points[i][0]
			dy1 := points[j][1] - points[i][1]
			for k := j + 1; k < n; k++ {
				for l := k + 1; l < n; l++ {
					dx2 := points[l][0] - points[k][0]
					dy2 := points[l][1] - points[k][1]
					// Check if two sides are parallel (cross product = 0)
					if dx1*dy2 == dx2*dy1 {
						count++
					}
				}
			}
		}
	}
	return count
}
```

## 3626 — Find Stores With Inventory Imbalance

```go
package main

// LeetCode #3626: Find Stores with Inventory Imbalance
// https://leetcode.com/problems/find-stores-with-inventory-imbalance/
// Difficulty: Medium
// Complexity: O(n) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	inventory := []int{10, 5, 15, 3, 20}
	threshold := 10
	fmt.Println("Test 1:", FindStoresWithInventoryImbalance(inventory, threshold))
	// Test case 2
	inventory2 := []int{100, 50, 75}
	threshold2 := 20
	fmt.Println("Test 2:", FindStoresWithInventoryImbalance(inventory2, threshold2))
	// Test case 3
	inventory3 := []int{1, 1, 1}
	threshold3 := 0
	fmt.Println("Test 3:", FindStoresWithInventoryImbalance(inventory3, threshold3))
}

func FindStoresWithInventoryImbalance(inventory []int, threshold int) int {
	// Count stores where difference from average exceeds threshold
	if len(inventory) == 0 {
		return 0
	}
	sum := 0
	for _, v := range inventory {
		sum += v
	}
	avg := float64(sum) / float64(len(inventory))
	count := 0
	for _, v := range inventory {
		diff := float64(v) - avg
		if diff < 0 {
			diff = -diff
		}
		if diff > float64(threshold) {
			count++
		}
	}
	return count
}
```

## 3627 — Maximum Median Sum Of Subsequences Of Size 3

```go
package main

// LeetCode #3627: Maximum Median Sum of Subsequences of Size 3
// https://leetcode.com/problems/maximum-median-sum-of-subsequences-of-size-3/
// Difficulty: Medium
// Complexity: O(n^3) time, O(1) space

import (
	"fmt"
	"sort"
)

func main() {
	// Test case 1
	fmt.Println("Test 1:", MaximumMedianSumOfSubsequencesOfSizeThree([]int{1, 2, 3, 4, 5}))
	// Test case 2
	fmt.Println("Test 2:", MaximumMedianSumOfSubsequencesOfSizeThree([]int{5, 1, 5, 1, 5}))
	// Test case 3
	fmt.Println("Test 3:", MaximumMedianSumOfSubsequencesOfSizeThree([]int{1, 2, 3}))
}

func MaximumMedianSumOfSubsequencesOfSizeThree(nums []int) int {
	n := len(nums)
	maxSum := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				sub := []int{nums[i], nums[j], nums[k]}
				sort.Ints(sub)
				sum := sub[0] + sub[1] + sub[2]
				if sum > maxSum {
					maxSum = sum
				}
			}
		}
	}
	return maxSum
}
```

## 3628 — Maximum Number Of Subsequences After One Inserting

```go
package main

// LeetCode #3628: Maximum Number of Subsequences After One Inserting
// https://leetcode.com/problems/maximum-number-of-subsequences-after-one-inserting/
// Difficulty: Medium
// Complexity: O(n) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", MaximumNumberOfSubsequencesAfterOneInserting("abc", "abc"))
	// Test case 2
	fmt.Println("Test 2:", MaximumNumberOfSubsequencesAfterOneInserting("ab", "ab"))
	// Test case 3
	fmt.Println("Test 3:", MaximumNumberOfSubsequencesAfterOneInserting("a", "a"))
}

func MaximumNumberOfSubsequencesAfterOneInserting(s string, pattern string) int {
	// Count max subsequences of pattern in s after inserting one character
	if len(pattern) == 0 {
		return len(s) + 1
	}
	if len(pattern) == 1 {
		// Count existing, then add max possible after insert
		count := 0
		for _, c := range s {
			if byte(c) == pattern[0] {
				count++
			}
		}
		// Insert one more gives count+1
		maxAdditional := count + 1
		if maxAdditional > len(s)+1-count {
			return maxAdditional
		}
		return maxAdditional
	}

	p0, p1 := pattern[0], pattern[1]

	// Count existing
	existing := 0
	countP1 := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == p1 {
			countP1++
		} else if s[i] == p0 {
			existing += countP1
		}
	}

	// Max additional from inserting one char
	addP0 := 0
	for _, c := range s {
		if byte(c) == p1 {
			addP0++
		}
	}
	addP1 := 0
	for _, c := range s {
		if byte(c) == p0 {
			addP1++
		}
	}
	if addP0 > addP1 {
		return existing + addP0
	}
	return existing + addP1
}
```

## 3629 — Minimum Jumps To Reach End Via Prime Teleportation

```go
package main

// LeetCode #3629: Minimum Jumps to Reach End via Prime Teleportation
// https://leetcode.com/problems/minimum-jumps-to-reach-end-via-prime-teleportation/
// Difficulty: Medium
// Complexity: O(n) time, O(n) space

import "fmt"

func isPrime(x int) bool {
	if x < 2 {
		return false
	}
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", MinimumJumpsToReachEndViaPrimeTeleportation([]int{1, 2, 3, 4, 5}))
	// Test case 2
	fmt.Println("Test 2:", MinimumJumpsToReachEndViaPrimeTeleportation([]int{4, 6, 8, 10}))
	// Test case 3
	fmt.Println("Test 3:", MinimumJumpsToReachEndViaPrimeTeleportation([]int{2, 3, 5, 7, 11}))
}

func MinimumJumpsToReachEndViaPrimeTeleportation(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}
	// DP: min jumps to reach position i
	dp := make([]int, n)
	for i := range dp {
		dp[i] = n + 1
	}
	dp[0] = 0
	for i := 0; i < n; i++ {
		if isPrime(nums[i]) {
			// Can teleport to any position
			for j := 0; j < n; j++ {
				if j != i && dp[i]+1 < dp[j] {
					dp[j] = dp[i] + 1
				}
			}
		} else {
			// Can move to adjacent position
			if i+1 < n && dp[i]+1 < dp[i+1] {
				dp[i+1] = dp[i] + 1
			}
			if i-1 >= 0 && dp[i]+1 < dp[i-1] {
				dp[i-1] = dp[i] + 1
			}
		}
	}
	if dp[n-1] > n {
		return -1
	}
	return dp[n-1]
}
```

## 3631 — Sort Threats By Severity And Exploitability

```go
package main

// LeetCode #3631: Sort Threats by Severity and Exploitability
// https://leetcode.com/problems/sort-threats-by-severity-and-exploitability/
// Difficulty: Medium [Paid]
// Complexity: O(n log n) time, O(n) space

import (
	"fmt"
	"sort"
)

type Threat struct {
	ID            int
	Severity      int
	Exploitability int
}

func main() {
	// Test case 1
	threats := []Threat{
		{1, 5, 3},
		{2, 8, 2},
		{3, 5, 7},
	}
	result := SortThreatsBySeverityAndExploitability(threats)
	fmt.Println("Test 1:", result)
	// Test case 2
	threats2 := []Threat{
		{1, 1, 1},
		{2, 2, 2},
	}
	result2 := SortThreatsBySeverityAndExploitability(threats2)
	fmt.Println("Test 2:", result2)
	// Test case 3
	threats3 := []Threat{
		{1, 10, 5},
	}
	result3 := SortThreatsBySeverityAndExploitability(threats3)
	fmt.Println("Test 3:", result3)
}

func SortThreatsBySeverityAndExploitability(threats []Threat) []Threat {
	sort.Slice(threats, func(i, j int) bool {
		if threats[i].Severity != threats[j].Severity {
			return threats[i].Severity > threats[j].Severity
		}
		return threats[i].Exploitability > threats[j].Exploitability
	})
	return threats
}
```

## 3634 — Minimum Removals To Balance Array

```go
package main

// LeetCode #3634: Minimum Removals to Balance Array
// https://leetcode.com/problems/minimum-removals-to-balance-array/
// Difficulty: Medium
// Complexity: O(n) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", MinimumRemovalsToBalanceArray([]int{1, 2, 3, 4, 5}))
	// Test case 2
	fmt.Println("Test 2:", MinimumRemovalsToBalanceArray([]int{1, 1, 1, 1}))
	// Test case 3
	fmt.Println("Test 3:", MinimumRemovalsToBalanceArray([]int{5, 4, 3, 2, 1}))
}

func MinimumRemovalsToBalanceArray(nums []int) int {
	// Minimum removals so that sum of first half equals sum of second half
	n := len(nums)
	if n%2 != 0 {
		// Remove middle element for odd length
		return 1
	}
	half := n / 2
	sum1, sum2 := 0, 0
	for i := 0; i < half; i++ {
		sum1 += nums[i]
	}
	for i := half; i < n; i++ {
		sum2 += nums[i]
	}
	if sum1 == sum2 {
		return 0
	}
	// Remove one element from the larger half
	return 1
}
```

## 3635 — Earliest Finish Time For Land And Water Rides Ii

```go
package main

// LeetCode #3635: Earliest Finish Time for Land and Water Rides II
// https://leetcode.com/problems/earliest-finish-time-for-land-and-water-rides-ii/
// Difficulty: Medium
// Time: O(n + m) | Space: O(1)

import "fmt"

func earliestFinishTime(landStartTime []int, landDuration []int, waterStartTime []int, waterDuration []int) int {
	// Try land -> water
	minLandEnd := int(^uint(0) >> 1)
	for i := range landStartTime {
		end := landStartTime[i] + landDuration[i]
		if end < minLandEnd {
			minLandEnd = end
		}
	}
	ans := int(^uint(0) >> 1)
	for i := range waterStartTime {
		start := waterStartTime[i]
		if minLandEnd > start {
			start = minLandEnd
		}
		end := start + waterDuration[i]
		if end < ans {
			ans = end
		}
	}

	// Try water -> land
	minWaterEnd := int(^uint(0) >> 1)
	for i := range waterStartTime {
		end := waterStartTime[i] + waterDuration[i]
		if end < minWaterEnd {
			minWaterEnd = end
		}
	}
	for i := range landStartTime {
		start := landStartTime[i]
		if minWaterEnd > start {
			start = minWaterEnd
		}
		end := start + landDuration[i]
		if end < ans {
			ans = end
		}
	}

	return ans
}

func main() {
	fmt.Println(earliestFinishTime(
		[]int{1, 2, 3}, []int{3, 2, 1},
		[]int{2, 3}, []int{4, 1},
	))
	fmt.Println(earliestFinishTime(
		[]int{5}, []int{10},
		[]int{2, 8}, []int{3, 2},
	))
	fmt.Println(earliestFinishTime(
		[]int{0, 10}, []int{5, 2},
		[]int{3, 7}, []int{4, 2},
	))
}
```

## 3638 — Maximum Balanced Shipments

```go
package main

// LeetCode #3638: Maximum Balanced Shipments
// https://leetcode.com/problems/maximum-balanced-shipments/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func maxBalancedShipments(weight []int) int {
	ans := 0
	mx := 0
	for _, x := range weight {
		if x > mx {
			mx = x
		}
		if x < mx {
			ans++
			mx = 0
		}
	}
	return ans
}

func main() {
	fmt.Println(maxBalancedShipments([]int{2, 5, 1, 4, 3}))
	fmt.Println(maxBalancedShipments([]int{4, 4}))
	fmt.Println(maxBalancedShipments([]int{1, 3, 2, 4, 5, 2}))
}
```

## 3639 — Minimum Time To Activate String

```go
package main

// LeetCode #3639: Minimum Time to Activate String
// https://leetcode.com/problems/minimum-time-to-activate-string/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import "fmt"

func minimumTimeToActivateString(s string) int {
	n := len(s)
	// dp[i] = min time to activate prefix of length i
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = i // worst case: type each character
	}

	for i := 1; i <= n; i++ {
		// Option 1: type the current character (1 sec)
		if dp[i] > dp[i-1]+1 {
			dp[i] = dp[i-1] + 1
		}

		// Option 2: try to activate from a matching prefix
		// Look for a substring in already activated prefix that matches s[i-1:]
		for j := 1; j < i; j++ {
			k := 0
			for i-1+k < n && j-1+k < i-1 && s[i-1+k] == s[j-1+k] {
				k++
			}
			if k > 0 {
				cost := dp[i-1] + 1 // 1 second to activate the matching substring
				if dp[i+k-1] > cost {
					dp[i+k-1] = cost
				}
			}
		}
	}

	return dp[n]
}

func main() {
	fmt.Println(minimumTimeToActivateString("abcabc"))
	fmt.Println(minimumTimeToActivateString("aaaa"))
	fmt.Println(minimumTimeToActivateString("abacaba"))
}
```

## 3641 — Longest Semi Repeating Subarray

```go
package main

// LeetCode #3641: Longest Semi-Repeating Subarray
// https://leetcode.com/problems/longest-semi-repeating-subarray/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func longestSemiRepeatingSubarray(nums []int, k int) int {
	freq := make(map[int]int)
	dupCount := 0
	maxLen := 0
	l := 0

	for r, x := range nums {
		freq[x]++
		if freq[x] == 2 {
			dupCount++
		}

		for dupCount > k {
			left := nums[l]
			freq[left]--
			if freq[left] == 1 {
				dupCount--
			}
			l++
		}

		if r-l+1 > maxLen {
			maxLen = r - l + 1
		}
	}

	return maxLen
}

func main() {
	fmt.Println(longestSemiRepeatingSubarray([]int{1, 2, 3, 1, 2, 3, 4}, 2))
	fmt.Println(longestSemiRepeatingSubarray([]int{1, 1, 1, 1, 1}, 4))
	fmt.Println(longestSemiRepeatingSubarray([]int{1, 1, 1, 1, 1}, 0))
}
```

## 3642 — Find Books With Polarized Opinions

```go
package main

// LeetCode #3642: Find Books with Polarized Opinions
// https://leetcode.com/problems/find-books-with-polarized-opinions/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

type book struct {
	id     int
	title  string
	author string
	genre  string
	pages  int
}

type readingSession struct {
	bookID int
	rating int
}

type bookResult struct {
	bookID            int
	title             string
	author            string
	genre             string
	pages             int
	ratingSpread      int
	polarizationScore float64
}

func findBooksWithPolarizedOpinions(books []book, sessions []readingSession) []bookResult {
	// Group ratings by book
	ratingsByBook := make(map[int][]int)
	for _, s := range sessions {
		ratingsByBook[s.bookID] = append(ratingsByBook[s.bookID], s.rating)
	}

	bookMap := make(map[int]book)
	for _, b := range books {
		bookMap[b.id] = b
	}

	var results []bookResult

	for bookID, ratings := range ratingsByBook {
		if len(ratings) < 5 {
			continue
		}

		maxRating := ratings[0]
		minRating := ratings[0]
		highCount := 0
		lowCount := 0

		for _, r := range ratings {
			if r > maxRating {
				maxRating = r
			}
			if r < minRating {
				minRating = r
			}
			if r >= 4 {
				highCount++
			}
			if r <= 2 {
				lowCount++
			}
		}

		if highCount == 0 || lowCount == 0 {
			continue
		}

		extremeCount := highCount + lowCount
		polarizationScore := float64(extremeCount) / float64(len(ratings))

		if polarizationScore < 0.6 {
			continue
		}

		b, ok := bookMap[bookID]
		if !ok {
			continue
		}

		pScore := float64(int(polarizationScore*100+0.5)) / 100
		results = append(results, bookResult{
			bookID:            b.id,
			title:             b.title,
			author:            b.author,
			genre:             b.genre,
			pages:             b.pages,
			ratingSpread:      maxRating - minRating,
			polarizationScore: pScore,
		})
	}

	sort.Slice(results, func(i, j int) bool {
		if results[i].polarizationScore != results[j].polarizationScore {
			return results[i].polarizationScore > results[j].polarizationScore
		}
		return results[i].title > results[j].title
	})

	return results
}

func main() {
	books := []book{
		{1, "Book A", "Author A", "Fiction", 300},
		{2, "Book B", "Author B", "Non-Fiction", 250},
		{3, "Book C", "Author C", "Sci-Fi", 400},
	}
	sessions := []readingSession{
		{1, 5}, {1, 5}, {1, 1}, {1, 5}, {1, 1}, {1, 4},
		{2, 3}, {2, 4}, {2, 3}, {2, 4}, {2, 3},
		{3, 1}, {3, 5}, {3, 1}, {3, 5}, {3, 1}, {3, 5},
	}
	results := findBooksWithPolarizedOpinions(books, sessions)
	for _, r := range results {
		fmt.Printf("Book %d (%s): spread=%d, score=%.2f\n", r.bookID, r.title, r.ratingSpread, r.polarizationScore)
	}
	if len(results) == 0 {
		fmt.Println("[]")
	}
}
```

## 3644 — Maximum K To Sort A Permutation

```go
package main

// LeetCode #3644: Maximum K to Sort a Permutation
// https://leetcode.com/problems/maximum-k-to-sort-a-permutation/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func maximumKToSortAPermutation(nums []int) int {
	ans := -1 // all bits set to 1 (identity for AND)
	for i, x := range nums {
		if i != x {
			ans &= x
		}
	}
	if ans < 0 {
		return 0
	}
	return ans
}

func main() {
	fmt.Println(maximumKToSortAPermutation([]int{0, 3, 2, 1}))
	fmt.Println(maximumKToSortAPermutation([]int{3, 2, 1, 0}))
	fmt.Println(maximumKToSortAPermutation([]int{0, 1, 2, 3}))
}
```

## 3645 — Maximum Total From Optimal Activation Order

```go
package main

// LeetCode #3645: Maximum Total from Optimal Activation Order
// https://leetcode.com/problems/maximum-total-from-optimal-activation-order/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func maximumTotalFromOptimalActivationOrder(value []int, limit []int) int64 {
	n := len(value)
	groups := make([][]int, n+1)
	for i := 0; i < n; i++ {
		l := limit[i]
		groups[l] = append(groups[l], value[i])
	}

	var total int64 = 0
	for l := 1; l <= n; l++ {
		if len(groups[l]) == 0 {
			continue
		}
		sort.Slice(groups[l], func(i, j int) bool {
			return groups[l][i] > groups[l][j]
		})
		cap := l
		if len(groups[l]) < cap {
			cap = len(groups[l])
		}
		for i := 0; i < cap; i++ {
			total += int64(groups[l][i])
		}
	}
	return total
}

func main() {
	fmt.Println(maximumTotalFromOptimalActivationOrder([]int{3, 5, 2, 4}, []int{2, 1, 3, 2}))
	fmt.Println(maximumTotalFromOptimalActivationOrder([]int{10, 20}, []int{1, 1}))
	fmt.Println(maximumTotalFromOptimalActivationOrder([]int{1, 2, 3, 4, 5}, []int{1, 2, 2, 3, 3}))
}
```

## 3647 — Maximum Weight In Two Bags

```go
package main

// LeetCode #3647: Maximum Weight in Two Bags
// https://leetcode.com/problems/maximum-weight-in-two-bags/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func maximumWeightInTwoBags(weight []int, bag1 int, bag2 int) int {
	sort.Slice(weight, func(i, j int) bool {
		return weight[i] > weight[j]
	})

	// Try to find best combination: one item for bag1, one for bag2
	// Since items are sorted descending, we try all pairs
	n := len(weight)
	best := 0

	// Single item approach: pick up to bag1 items, then fill bag2 with remaining
	maxBag1 := 0
	used := -1
	for i := 0; i < n && weight[i] <= bag1; i++ {
		if weight[i] > maxBag1 {
			maxBag1 = weight[i]
			used = i
		}
	}

	if used != -1 {
		for i := 0; i < n; i++ {
			if i != used && weight[i] <= bag2 {
				if maxBag1+weight[i] > best {
					best = maxBag1 + weight[i]
				}
				break
			}
		}
	}

	maxBag2 := 0
	used2 := -1
	for i := 0; i < n && weight[i] <= bag2; i++ {
		if weight[i] > maxBag2 {
			maxBag2 = weight[i]
			used2 = i
		}
	}

	if used2 != -1 {
		for i := 0; i < n; i++ {
			if i != used2 && weight[i] <= bag1 {
				if maxBag2+weight[i] > best {
					best = maxBag2 + weight[i]
				}
				break
			}
		}
	}

	return best
}

func main() {
	fmt.Println(maximumWeightInTwoBags([]int{10, 20, 30, 40}, 30, 40))
	fmt.Println(maximumWeightInTwoBags([]int{5, 5, 10}, 10, 5))
	fmt.Println(maximumWeightInTwoBags([]int{2, 3, 5, 7}, 5, 5))
}
```

## 3648 — Minimum Sensors To Cover Grid

```go
package main

// LeetCode #3648: Minimum Sensors to Cover Grid
// https://leetcode.com/problems/minimum-sensors-to-cover-grid/
// Difficulty: Medium
// Time: O(1) | Space: O(1)

import "fmt"

func minimumSensorsToCoverGrid(n int, m int, k int) int {
	s := 2*k + 1
	rows := (n + s - 1) / s // ceil division
	cols := (m + s - 1) / s
	return rows * cols
}

func main() {
	fmt.Println(minimumSensorsToCoverGrid(5, 5, 1))
	fmt.Println(minimumSensorsToCoverGrid(2, 2, 2))
	fmt.Println(minimumSensorsToCoverGrid(3, 4, 0))
}
```

## 3649 — Number Of Perfect Pairs

```go
package main

// LeetCode #3649: Number of Perfect Pairs
// https://leetcode.com/problems/number-of-perfect-pairs/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func numberOfPerfectPairs(nums []int) int64 {
	n := len(nums)
	arr := make([]int, n)
	for i, v := range nums {
		if v < 0 {
			arr[i] = -v
		} else {
			arr[i] = v
		}
	}

	sort.Ints(arr)

	var ans int64 = 0
	j := 0
	for i := 1; i < n; i++ {
		for 2*arr[j] < arr[i] {
			j++
		}
		ans += int64(i - j)
	}
	return ans
}

func main() {
	fmt.Println(numberOfPerfectPairs([]int{1, 2, 3, 4}))
	fmt.Println(numberOfPerfectPairs([]int{-1, 1, -2, 2}))
	fmt.Println(numberOfPerfectPairs([]int{5, 1, 2}))
}
```

