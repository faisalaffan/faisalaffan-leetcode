# Medium (Sedang) — Problem 0901–1077

## 0901 — Online Stock Span

```go
package main

// LeetCode #901: Online Stock Span
// https://leetcode.com/problems/online-stock-span/
// Difficulty: Medium

import "fmt"

type StockSpanner struct {
	stk []pair
}

type pair struct {
	price, span int
}

func Constructor() StockSpanner {
	return StockSpanner{[]pair{}}
}

func (this *StockSpanner) Next(price int) int {
	span := 1
	for len(this.stk) > 0 && this.stk[len(this.stk)-1].price <= price {
		span += this.stk[len(this.stk)-1].span
		this.stk = this.stk[:len(this.stk)-1]
	}
	this.stk = append(this.stk, pair{price, span})
	return span
}

func main() {
	obj := Constructor()
	fmt.Println(obj.Next(100))
	fmt.Println(obj.Next(80))
	fmt.Println(obj.Next(60))
	fmt.Println(obj.Next(70))
	fmt.Println(obj.Next(60))
	fmt.Println(obj.Next(75))
	fmt.Println(obj.Next(85))
}
```

## 0904 — Fruit Into Baskets

```go
package main

// LeetCode #904: Fruit Into Baskets
// https://leetcode.com/problems/fruit-into-baskets/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(FruitIntoBaskets([]int{1, 2, 1}))
	fmt.Println(FruitIntoBaskets([]int{0, 1, 2, 2}))
	fmt.Println(FruitIntoBaskets([]int{1, 2, 3, 2, 2}))
}

// Time: O(n) | Space: O(1)
func FruitIntoBaskets(fruits []int) int {
	cnt := make(map[int]int)
	left, ans := 0, 0

	for right, fruit := range fruits {
		cnt[fruit]++
		for len(cnt) > 2 {
			cnt[fruits[left]]--
			if cnt[fruits[left]] == 0 {
				delete(cnt, fruits[left])
			}
			left++
		}
		if right-left+1 > ans {
			ans = right - left + 1
		}
	}

	return ans
}
```

## 0907 — Sum Of Subarray Minimums

```go
package main

// LeetCode #907: Sum of Subarray Minimums
// https://leetcode.com/problems/sum-of-subarray-minimums/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(SumOfSubarrayMinimums([]int{3, 1, 2, 4}))
	fmt.Println(SumOfSubarrayMinimums([]int{11, 81, 94, 43, 3}))
	fmt.Println(SumOfSubarrayMinimums([]int{71, 55, 82, 55}))
}

// Time: O(n) | Space: O(n)
func SumOfSubarrayMinimums(arr []int) int {
	const mod = 1_000_000_007
	n := len(arr)

	prevSmaller := make([]int, n)
	nextSmaller := make([]int, n)

	for i := 0; i < n; i++ {
		prevSmaller[i] = -1
		nextSmaller[i] = n
	}

	var stack []int
	for i := 0; i < n; i++ {
		for len(stack) > 0 && arr[stack[len(stack)-1]] > arr[i] {
			nextSmaller[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	stack = nil
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && arr[stack[len(stack)-1]] >= arr[i] {
			prevSmaller[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	ans := 0
	for i := 0; i < n; i++ {
		left := i - prevSmaller[i]
		right := nextSmaller[i] - i
		ans = (ans + arr[i]*left*right) % mod
	}

	return ans
}
```

## 0909 — Snakes And Ladders

```go
package main

// LeetCode #909: Snakes and Ladders
// https://leetcode.com/problems/snakes-and-ladders/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(SnakesAndLadders([][]int{
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, 35, -1, -1, 13, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, 15, -1, -1, -1, -1},
	}))
	fmt.Println(SnakesAndLadders([][]int{{-1, -1}, {-1, 3}}))
}

// Time: O(n^2) | Space: O(n^2)
func SnakesAndLadders(board [][]int) int {
	n := len(board)
	target := n * n

	// Convert board position to (row, col)
	posToCoord := func(pos int) (int, int) {
		row := (pos - 1) / n
		col := (pos - 1) % n
		if row%2 == 1 {
			col = n - 1 - col
		}
		return n - 1 - row, col
	}

	dist := make([]int, target+1)
	for i := range dist {
		dist[i] = -1
	}
	dist[1] = 0

	queue := []int{1}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr == target {
			return dist[curr]
		}

		for next := curr + 1; next <= curr+6 && next <= target; next++ {
			r, c := posToCoord(next)
			dest := next
			if board[r][c] != -1 {
				dest = board[r][c]
			}
			if dist[dest] == -1 {
				dist[dest] = dist[curr] + 1
				queue = append(queue, dest)
			}
		}
	}

	return -1
}
```

## 0910 — Smallest Range Ii

```go
package main

// LeetCode #910: Smallest Range II
// https://leetcode.com/problems/smallest-range-ii/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(SmallestRangeIi([]int{1}, 0))
	fmt.Println(SmallestRangeIi([]int{0, 10}, 2))
	fmt.Println(SmallestRangeIi([]int{1, 3, 6}, 3))
}

// Time: O(n log n) | Space: O(log n)
func SmallestRangeIi(nums []int, k int) int {
	sort.Ints(nums)
	n := len(nums)
	ans := nums[n-1] - nums[0]

	for i := 1; i < n; i++ {
		mx := max(nums[i-1]+k, nums[n-1]-k)
		mn := min(nums[0]+k, nums[i]-k)
		ans = min(ans, mx-mn)
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
```

## 0911 — Online Election

```go
package main

// LeetCode #911: Online Election
// https://leetcode.com/problems/online-election/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

type TopVotedCandidate struct {
	times []int
	wins  []int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	n := len(times)
	wins := make([]int, n)
	votes := make(map[int]int)
	leader := -1

	for i := 0; i < n; i++ {
		votes[persons[i]]++
		if leader == -1 || votes[persons[i]] >= votes[leader] {
			leader = persons[i]
		}
		wins[i] = leader
	}

	return TopVotedCandidate{times, wins}
}

func (this *TopVotedCandidate) Q(t int) int {
	idx := sort.SearchInts(this.times, t)
	if idx < len(this.times) && this.times[idx] == t {
		return this.wins[idx]
	}
	return this.wins[idx-1]
}

func main() {
	obj := Constructor([]int{0, 1, 1, 0, 0, 1, 0}, []int{0, 5, 10, 15, 20, 25, 30})
	fmt.Println(obj.Q(3))
	fmt.Println(obj.Q(12))
	fmt.Println(obj.Q(25))
	fmt.Println(obj.Q(15))
	fmt.Println(obj.Q(24))
	fmt.Println(obj.Q(8))
}
```

## 0912 — Sort An Array

```go
package main

// LeetCode #912: Sort an Array
// https://leetcode.com/problems/sort-an-array/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(SortAnArray([]int{5, 2, 3, 1}))
	fmt.Println(SortAnArray([]int{5, 1, 1, 2, 0, 0}))
	fmt.Println(SortAnArray([]int{3, -1}))
}

// Time: O(n log n) | Space: O(n)
func SortAnArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2
	left := SortAnArray(nums[:mid])
	right := SortAnArray(nums[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
```

## 0915 — Partition Array Into Disjoint Intervals

```go
package main

// LeetCode #915: Partition Array into Disjoint Intervals
// https://leetcode.com/problems/partition-array-into-disjoint-intervals/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(PartitionArrayIntoDisjointIntervals([]int{5, 0, 3, 8, 6}))
	fmt.Println(PartitionArrayIntoDisjointIntervals([]int{1, 1, 1, 0, 6, 12}))
	fmt.Println(PartitionArrayIntoDisjointIntervals([]int{1, 1}))
}

// Time: O(n) | Space: O(1)
func PartitionArrayIntoDisjointIntervals(nums []int) int {
	leftMax, curMax, idx := nums[0], nums[0], 0

	for i, v := range nums {
		if v > curMax {
			curMax = v
		}
		if v < leftMax {
			leftMax = curMax
			idx = i
		}
	}

	return idx + 1
}
```

## 0916 — Word Subsets

```go
package main

// LeetCode #916: Word Subsets
// https://leetcode.com/problems/word-subsets/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(WordSubsets([]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"e", "o"}))
	fmt.Println(WordSubsets([]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"l", "e"}))
	fmt.Println(WordSubsets([]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"e", "oo"}))
}

// Time: O((n + m) * L) | Space: O(1)
func WordSubsets(words1 []string, words2 []string) []string {
	maxCnt := [26]int{}
	for _, word := range words2 {
		var cnt [26]int
		for _, ch := range word {
			cnt[ch-'a']++
		}
		for i := 0; i < 26; i++ {
			if cnt[i] > maxCnt[i] {
				maxCnt[i] = cnt[i]
			}
		}
	}

	var ans []string
	for _, word := range words1 {
		var cnt [26]int
		for _, ch := range word {
			cnt[ch-'a']++
		}
		ok := true
		for i := 0; i < 26; i++ {
			if cnt[i] < maxCnt[i] {
				ok = false
				break
			}
		}
		if ok {
			ans = append(ans, word)
		}
	}

	return ans
}
```

## 0918 — Maximum Sum Circular Subarray

```go
package main

// LeetCode #918: Maximum Sum Circular Subarray
// https://leetcode.com/problems/maximum-sum-circular-subarray/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MaximumSumCircularSubarray([]int{1, -2, 3, -2}))
	fmt.Println(MaximumSumCircularSubarray([]int{5, -3, 5}))
	fmt.Println(MaximumSumCircularSubarray([]int{-3, -2, -3}))
}

// Time: O(n) | Space: O(1)
func MaximumSumCircularSubarray(nums []int) int {
	total := 0
	maxSum, curMax := nums[0], 0
	minSum, curMin := nums[0], 0

	for _, v := range nums {
		total += v
		curMax = max(curMax+v, v)
		maxSum = max(maxSum, curMax)
		curMin = min(curMin+v, v)
		minSum = min(minSum, curMin)
	}

	// If all numbers are negative, return the max (non-circular)
	if maxSum < 0 {
		return maxSum
	}

	return max(maxSum, total-minSum)
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
```

## 0919 — Complete Binary Tree Inserter

```go
package main

// LeetCode #919: Complete Binary Tree Inserter
// https://leetcode.com/problems/complete-binary-tree-inserter/
// Difficulty: Medium

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type CBTInserter struct {
	root *TreeNode
	q    []*TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	q := []*TreeNode{root}
	for {
		node := q[0]
		if node.Left != nil {
			q = append(q, node.Left)
		} else {
			break
		}
		if node.Right != nil {
			q = append(q, node.Right)
		} else {
			break
		}
		q = q[1:]
	}
	return CBTInserter{root, q}
}

func (this *CBTInserter) Insert(val int) int {
	parent := this.q[0]
	node := &TreeNode{Val: val}
	if parent.Left == nil {
		parent.Left = node
	} else {
		parent.Right = node
		this.q = this.q[1:]
	}
	this.q = append(this.q, node)
	return parent.Val
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
}

func main() {
	root := &TreeNode{Val: 1}
	obj := Constructor(root)
	fmt.Println(obj.Insert(2))
	fmt.Println(obj.Get_root().Val)
	fmt.Println(obj.Insert(3))
	fmt.Println(obj.Insert(4))
	fmt.Println(obj.Get_root().Val)
}
```

## 0921 — Minimum Add To Make Parentheses Valid

```go
package main

// LeetCode #921: Minimum Add to Make Parentheses Valid
// https://leetcode.com/problems/minimum-add-to-make-parentheses-valid/
// Difficulty: Medium

import "fmt"

// Time: O(n) | Space: O(1)
func minAddToMakeValid(s string) int {
	open, add := 0, 0
	for _, ch := range s {
		if ch == '(' {
			open++
		} else {
			if open > 0 {
				open--
			} else {
				add++
			}
		}
	}
	return add + open
}

func main() {
	fmt.Println(minAddToMakeValid("())"))
	fmt.Println(minAddToMakeValid("((("))
	fmt.Println(minAddToMakeValid("()"))
	fmt.Println(minAddToMakeValid("()))(("))
}
```

## 0923 — 3Sum With Multiplicity

```go
package main

// LeetCode #923: 3Sum With Multiplicity
// https://leetcode.com/problems/3sum-with-multiplicity/
// Difficulty: Medium

import "fmt"

const mod = 1_000_000_007

// Time: O(n^2) | Space: O(1) if sort in-place considered O(1), else O(n)
func threeSumMulti(arr []int, target int) int {
	var cnt [101]int
	for _, v := range arr {
		cnt[v]++
	}

	ans := 0
	// Case 1: all three same
	for i := 0; i <= 100; i++ {
		if cnt[i] >= 3 && i*3 == target {
			ans = (ans + cnt[i]*(cnt[i]-1)*(cnt[i]-2)/6) % mod
		}
		// Case 2: two same, one different
		if cnt[i] >= 2 {
			remain := target - 2*i
			if remain >= 0 && remain <= 100 && remain != i && cnt[remain] > 0 {
				ans = (ans + cnt[i]*(cnt[i]-1)/2*cnt[remain]) % mod
			}
		}
		// Case 3: all three different
		for j := i + 1; j <= 100; j++ {
			k := target - i - j
			if k > j && k <= 100 && cnt[k] > 0 {
				ans = (ans + cnt[i]*cnt[j]*cnt[k]) % mod
			}
		}
	}

	return ans
}

func main() {
	fmt.Println(threeSumMulti([]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}, 8))
	fmt.Println(threeSumMulti([]int{1, 1, 2, 2, 2, 2}, 5))
	fmt.Println(threeSumMulti([]int{2, 1, 3}, 6))
}
```

## 0926 — Flip String To Monotone Increasing

```go
package main

// LeetCode #926: Flip String to Monotone Increasing
// https://leetcode.com/problems/flip-string-to-monotone-increasing/
// Difficulty: Medium

import "fmt"

// Time: O(n) | Space: O(1)
func minFlipsMonoIncr(s string) int {
	ones, flips := 0, 0
	for _, ch := range s {
		if ch == '1' {
			ones++
		} else {
			flips = min(flips+1, ones)
		}
	}
	return flips
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minFlipsMonoIncr("00110"))
	fmt.Println(minFlipsMonoIncr("010110"))
	fmt.Println(minFlipsMonoIncr("00011000"))
}
```

## 0930 — Binary Subarrays With Sum

```go
package main

// LeetCode #930: Binary Subarrays With Sum
// https://leetcode.com/problems/binary-subarrays-with-sum/
// Difficulty: Medium

import "fmt"

// Time: O(n) | Space: O(1)
func numSubarraysWithSum(nums []int, goal int) int {
	// sliding window for sum <= goal, then subtract sum < goal
	atMost := func(g int) int {
		if g < 0 {
			return 0
		}
		sum, cnt, left := 0, 0, 0
		for right, v := range nums {
			sum += v
			for sum > g {
				sum -= nums[left]
				left++
			}
			cnt += right - left + 1
		}
		return cnt
	}
	return atMost(goal) - atMost(goal-1)
}

func main() {
	fmt.Println(numSubarraysWithSum([]int{1, 0, 1, 0, 1}, 2))
	fmt.Println(numSubarraysWithSum([]int{0, 0, 0, 0, 0}, 0))
}
```

## 0931 — Minimum Falling Path Sum

```go
package main

// LeetCode #931: Minimum Falling Path Sum
// https://leetcode.com/problems/minimum-falling-path-sum/
// Difficulty: Medium

import (
	"fmt"
	"math"
)

// Time: O(n^2) | Space: O(1) — modifies input in place
func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			best := matrix[i-1][j]
			if j > 0 {
				best = min(best, matrix[i-1][j-1])
			}
			if j+1 < n {
				best = min(best, matrix[i-1][j+1])
			}
			matrix[i][j] += best
		}
	}

	ans := math.MaxInt32
	for _, v := range matrix[n-1] {
		ans = min(ans, v)
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minFallingPathSum([][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}))
	fmt.Println(minFallingPathSum([][]int{{-19, 57}, {-40, -5}}))
}
```

## 0932 — Beautiful Array

```go
package main

// LeetCode #932: Beautiful Array
// https://leetcode.com/problems/beautiful-array/
// Difficulty: Medium

import "fmt"

// Time: O(n log n) | Space: O(n)
func beautifulArray(n int) []int {
	memo := make(map[int][]int)
	var dfs func(int) []int
	dfs = func(n int) []int {
		if v, ok := memo[n]; ok {
			return v
		}
		res := make([]int, n)
		if n == 1 {
			res[0] = 1
		} else {
			left := dfs((n + 1) / 2)
			right := dfs(n / 2)
			for i, v := range left {
				res[i] = 2*v - 1
			}
			for i, v := range right {
				res[(n+1)/2+i] = 2 * v
			}
		}
		memo[n] = res
		return res
	}
	return dfs(n)
}

func main() {
	fmt.Println(beautifulArray(4))
	fmt.Println(beautifulArray(5))
	fmt.Println(beautifulArray(1))
}
```

## 0934 — Shortest Bridge

```go
package main

// LeetCode #934: Shortest Bridge
// https://leetcode.com/problems/shortest-bridge/
// Difficulty: Medium

import "fmt"

// Time: O(n^2) | Space: O(n^2)
func shortestBridge(grid [][]int) int {
	n := len(grid)
	dirs := []int{1, 0, -1, 0, 1}
	q := make([][3]int, 0)

	var dfs func(x, y int)
	dfs = func(x, y int) {
		grid[x][y] = 2
		for d := 0; d < 4; d++ {
			nx, ny := x+dirs[d], y+dirs[d+1]
			if nx >= 0 && nx < n && ny >= 0 && ny < n {
				if grid[nx][ny] == 1 {
					dfs(nx, ny)
				} else if grid[nx][ny] == 0 {
					grid[nx][ny] = 2
					q = append(q, [3]int{nx, ny, 1})
				}
			}
		}
	}

	// DFS to mark first island
	found := false
	for i := 0; i < n && !found; i++ {
		for j := 0; j < n && !found; j++ {
			if grid[i][j] == 1 {
				grid[i][j] = 2
				dfs(i, j)
				found = true
			}
		}
	}

	// BFS to find shortest path to second island
	for len(q) > 0 {
		cell := q[0]
		q = q[1:]
		for d := 0; d < 4; d++ {
			nx, ny := cell[0]+dirs[d], cell[1]+dirs[d+1]
			if nx >= 0 && nx < n && ny >= 0 && ny < n {
				if grid[nx][ny] == 1 {
					return cell[2]
				}
				if grid[nx][ny] == 0 {
					grid[nx][ny] = 2
					q = append(q, [3]int{nx, ny, cell[2] + 1})
				}
			}
		}
	}

	return -1
}

func main() {
	fmt.Println(shortestBridge([][]int{{0, 1}, {1, 0}}))
	fmt.Println(shortestBridge([][]int{{0, 1, 0}, {0, 0, 0}, {0, 0, 1}}))
	fmt.Println(shortestBridge([][]int{{1, 1, 1, 1, 1}, {1, 0, 0, 0, 1}, {1, 0, 1, 0, 1}, {1, 0, 0, 0, 1}, {1, 1, 1, 1, 1}}))
}
```

## 0935 — Knight Dialer

```go
package main

// LeetCode #935: Knight Dialer
// https://leetcode.com/problems/knight-dialer/
// Difficulty: Medium

import "fmt"

const mod = 1_000_000_007

// Time: O(n) | Space: O(1)
func knightDialer(n int) int {
	dp := [10]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	for i := 1; i < n; i++ {
		dp = [10]int{
			(dp[4] + dp[6]) % mod,
			(dp[6] + dp[8]) % mod,
			(dp[7] + dp[9]) % mod,
			(dp[4] + dp[8]) % mod,
			(dp[0] + dp[3] + dp[9]) % mod,
			0,
			(dp[0] + dp[1] + dp[7]) % mod,
			(dp[2] + dp[6]) % mod,
			(dp[1] + dp[3]) % mod,
			(dp[2] + dp[4]) % mod,
		}
	}
	sum := 0
	for _, v := range dp {
		sum = (sum + v) % mod
	}
	return sum
}

func main() {
	fmt.Println(knightDialer(1))
	fmt.Println(knightDialer(2))
	fmt.Println(knightDialer(3131))
}
```

## 0937 — Reorder Data In Log Files

```go
package main

// LeetCode #937: Reorder Data in Log Files
// https://leetcode.com/problems/reorder-data-in-log-files/
// Difficulty: Medium

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

// Time: O(n log n) | Space: O(n)
func reorderLogFiles(logs []string) []string {
	sort.SliceStable(logs, func(i, j int) bool {
		ci := strings.SplitN(logs[i], " ", 2)
		cj := strings.SplitN(logs[j], " ", 2)
		idI, restI := ci[0], ci[1]
		idJ, restJ := cj[0], cj[1]

		iDig := unicode.IsDigit(rune(restI[0]))
		jDig := unicode.IsDigit(rune(restJ[0]))

		if iDig && jDig {
			return false
		}
		if !iDig && !jDig {
			if restI != restJ {
				return restI < restJ
			}
			return idI < idJ
		}
		return !iDig
	})
	return logs
}

func main() {
	fmt.Println(reorderLogFiles([]string{"dig1 8 1 5 1", "let1 art can", "dig2 3 6", "let2 own kit dig", "let3 art zero"}))
	fmt.Println(reorderLogFiles([]string{"a1 9 2 3 1", "g1 act car", "zo4 4 7", "ab1 off key dog", "a8 act zoo"}))
}
```

## 0939 — Minimum Area Rectangle

```go
package main

// LeetCode #939: Minimum Area Rectangle
// https://leetcode.com/problems/minimum-area-rectangle/
// Difficulty: Medium

import (
	"fmt"
	"math"
)

// Time: O(n^2) | Space: O(n)
func minAreaRect(points [][]int) int {
	set := make(map[[2]int]bool)
	for _, p := range points {
		set[[2]int{p[0], p[1]}] = true
	}

	ans := math.MaxInt32
	n := len(points)

	for i := 0; i < n; i++ {
		x1, y1 := points[i][0], points[i][1]
		for j := i + 1; j < n; j++ {
			x2, y2 := points[j][0], points[j][1]
			if x1 == x2 || y1 == y2 {
				continue
			}
			if set[[2]int{x1, y2}] && set[[2]int{x2, y1}] {
				area := abs(x1-x2) * abs(y1-y2)
				if area < ans {
					ans = area
				}
			}
		}
	}

	if ans == math.MaxInt32 {
		return 0
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
	fmt.Println(minAreaRect([][]int{{1, 1}, {1, 3}, {3, 1}, {3, 3}, {2, 2}}))
	fmt.Println(minAreaRect([][]int{{1, 1}, {1, 3}, {3, 1}, {3, 3}, {4, 1}, {4, 3}}))
}
```

## 0945 — Minimum Increment To Make Array Unique

```go
package main

// LeetCode #945: Minimum Increment to Make Array Unique
// https://leetcode.com/problems/minimum-increment-to-make-array-unique/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

// Time: O(n log n) | Space: O(1) ignoring sort space
func minIncrementForUnique(nums []int) int {
	sort.Ints(nums)
	moves := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			moves += nums[i-1] + 1 - nums[i]
			nums[i] = nums[i-1] + 1
		}
	}
	return moves
}

func main() {
	fmt.Println(minIncrementForUnique([]int{1, 2, 2}))
	fmt.Println(minIncrementForUnique([]int{3, 2, 1, 2, 1, 7}))
}
```

## 0946 — Validate Stack Sequences

```go
package main

// LeetCode #946: Validate Stack Sequences
// https://leetcode.com/problems/validate-stack-sequences/
// Difficulty: Medium

import "fmt"

// Time: O(n) | Space: O(n)
func validateStackSequences(pushed []int, popped []int) bool {
	stack := make([]int, 0)
	j := 0
	for _, x := range pushed {
		stack = append(stack, x)
		for len(stack) > 0 && j < len(popped) && stack[len(stack)-1] == popped[j] {
			stack = stack[:len(stack)-1]
			j++
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}))
	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}))
}
```

## 0947 — Most Stones Removed With Same Row Or Column

```go
package main

// LeetCode #947: Most Stones Removed with Same Row or Column
// https://leetcode.com/problems/most-stones-removed-with-same-row-or-column/
// Difficulty: Medium

import "fmt"

// Time: O(n * α(n)) | Space: O(n)
func removeStones(stones [][]int) int {
	n := len(stones)
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
		ra, rb := find(a), find(b)
		if ra != rb {
			parent[rb] = ra
		}
	}

	// Union stones sharing row or column
	rowMap := make(map[int]int)
	colMap := make(map[int]int)
	for i, s := range stones {
		if v, ok := rowMap[s[0]]; ok {
			union(i, v)
		} else {
			rowMap[s[0]] = i
		}
		if v, ok := colMap[s[1]]; ok {
			union(i, v)
		} else {
			colMap[s[1]] = i
		}
	}

	// Count connected components
	components := 0
	for i := 0; i < n; i++ {
		if parent[i] == i {
			components++
		}
	}

	return n - components
}

func main() {
	fmt.Println(removeStones([][]int{{0, 0}, {0, 1}, {1, 0}, {1, 2}, {2, 1}, {2, 2}}))
	fmt.Println(removeStones([][]int{{0, 0}, {0, 2}, {1, 1}, {2, 0}, {2, 2}}))
	fmt.Println(removeStones([][]int{{0, 0}}))
}
```

## 0948 — Bag Of Tokens

```go
package main

// LeetCode #948: Bag of Tokens
// https://leetcode.com/problems/bag-of-tokens/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

// Time: O(n log n) | Space: O(1)
func bagOfTokensScore(tokens []int, power int) int {
	sort.Ints(tokens)
	left, right := 0, len(tokens)-1
	score, maxScore := 0, 0

	for left <= right {
		if power >= tokens[left] {
			power -= tokens[left]
			left++
			score++
			if score > maxScore {
				maxScore = score
			}
		} else if score > 0 {
			power += tokens[right]
			right--
			score--
		} else {
			break
		}
	}
	return maxScore
}

func main() {
	fmt.Println(bagOfTokensScore([]int{100}, 50))
	fmt.Println(bagOfTokensScore([]int{200, 100}, 150))
	fmt.Println(bagOfTokensScore([]int{100, 200, 300, 400}, 200))
}
```

## 0949 — Largest Time For Given Digits

```go
package main

// LeetCode #949: Largest Time for Given Digits
// https://leetcode.com/problems/largest-time-for-given-digits/
// Difficulty: Medium

import "fmt"

// Time: O(1) | Space: O(1)
func largestTimeFromDigits(arr []int) string {
	ans := -1
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 4; k++ {
				if i != j && j != k && i != k {
					l := 6 - i - j - k
					h := arr[i]*10 + arr[j]
					m := arr[k]*10 + arr[l]
					if h < 24 && m < 60 {
						if t := h*60 + m; t > ans {
							ans = t
						}
					}
				}
			}
		}
	}
	if ans < 0 {
		return ""
	}
	return fmt.Sprintf("%02d:%02d", ans/60, ans%60)
}

func main() {
	fmt.Println(largestTimeFromDigits([]int{1, 2, 3, 4}))
	fmt.Println(largestTimeFromDigits([]int{5, 5, 5, 5}))
	fmt.Println(largestTimeFromDigits([]int{2, 0, 6, 6}))
}
```

## 0950 — Reveal Cards In Increasing Order

```go
package main

// LeetCode #950: Reveal Cards In Increasing Order
// https://leetcode.com/problems/reveal-cards-in-increasing-order/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

// Time: O(n log n) | Space: O(n)
func deckRevealedIncreasing(deck []int) []int {
	sort.Ints(deck)
	n := len(deck)
	q := make([]int, n)
	for i := range q {
		q[i] = i
	}

	ans := make([]int, n)
	for _, card := range deck {
		idx := q[0]
		q = q[1:]
		ans[idx] = card
		if len(q) > 0 {
			q = append(q, q[0])
			q = q[1:]
		}
	}
	return ans
}

func main() {
	fmt.Println(deckRevealedIncreasing([]int{17, 13, 11, 2, 3, 5, 7}))
	fmt.Println(deckRevealedIncreasing([]int{1, 1000}))
}
```

## 0951 — Flip Equivalent Binary Trees

```go
package main

// LeetCode #951: Flip Equivalent Binary Trees
// https://leetcode.com/problems/flip-equivalent-binary-trees/
// Difficulty: Medium

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n) | Space: O(n)
func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil || root1.Val != root2.Val {
		return false
	}
	return (flipEquiv(root1.Left, root2.Left) && flipEquiv(root1.Right, root2.Right)) ||
		(flipEquiv(root1.Left, root2.Right) && flipEquiv(root1.Right, root2.Left))
}

func main() {
	// [1,2,3,4,5,6,null,null,null,7,8]
	root1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{
				Val: 5,
				Left:  &TreeNode{Val: 7},
				Right: &TreeNode{Val: 8},
			},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 6},
		},
	}
	// [1,3,2,null,6,4,5,null,null,null,null,8,7]
	root2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   3,
			Right: &TreeNode{Val: 6},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Right: &TreeNode{Val: 8},
			},
			Right: &TreeNode{
				Val:   5,
				Left:  &TreeNode{Val: 7},
			},
		},
	}
	fmt.Println(flipEquiv(root1, root2))

	// Trivial: both nil
	fmt.Println(flipEquiv(nil, nil))

	// Trivial: one nil
	fmt.Println(flipEquiv(&TreeNode{Val: 1}, nil))
}
```

## 0954 — Array Of Doubled Pairs

```go
package main

// LeetCode #954: Array of Doubled Pairs
// https://leetcode.com/problems/array-of-doubled-pairs/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

// Time: O(n log n) | Space: O(n)
func canReorderDoubled(arr []int) bool {
	freq := make(map[int]int)
	for _, v := range arr {
		freq[v]++
	}

	keys := make([]int, 0, len(freq))
	for k := range freq {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return abs(keys[i]) < abs(keys[j])
	})

	for _, v := range keys {
		if freq[v] > freq[2*v] {
			return false
		}
		freq[2*v] -= freq[v]
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(canReorderDoubled([]int{3, 1, 3, 6}))
	fmt.Println(canReorderDoubled([]int{2, 1, 2, 6}))
	fmt.Println(canReorderDoubled([]int{4, -2, 2, -4}))
	fmt.Println(canReorderDoubled([]int{1, 2, 4, 16, 8, 4}))
}
```

## 0955 — Delete Columns To Make Sorted Ii

```go
package main

// LeetCode #955: Delete Columns to Make Sorted II
// https://leetcode.com/problems/delete-columns-to-make-sorted-ii/
// Difficulty: Medium

import "fmt"

// Time: O(n * m) | Space: O(n)
func minDeletionSize(strs []string) int {
	m := len(strs)
	n := len(strs[0])
	cut := make([]bool, m)
	ans := 0

	for col := 0; col < n; col++ {
		ok := true
		for row := 0; row+1 < m; row++ {
			if !cut[row] && strs[row][col] > strs[row+1][col] {
				ans++
				ok = false
				break
			}
		}
		if ok {
			for row := 0; row+1 < m; row++ {
				if strs[row][col] < strs[row+1][col] {
					cut[row] = true
				}
			}
		}
	}

	return ans
}

func main() {
	fmt.Println(minDeletionSize([]string{"ca", "bb", "ac"}))
	fmt.Println(minDeletionSize([]string{"xc", "yb", "za"}))
	fmt.Println(minDeletionSize([]string{"zyx", "wvu", "tsr"}))
}
```

## 0957 — Prison Cells After N Days

```go
package main

// LeetCode #957: Prison Cells After N Days
// https://leetcode.com/problems/prison-cells-after-n-days/
// Difficulty: Medium

import "fmt"

// Time: O(1) | Space: O(1)
func prisonAfterNDays(cells []int, n int) []int {
	seen := make(map[[8]int]int)
	cycle := false

	for n > 0 {
		key := toArray(cells)
		if day, ok := seen[key]; ok && !cycle {
			n %= day - n
			cycle = true
		}
		seen[key] = n

		if n > 0 {
			n--
			cells = nextDay(cells)
		}
	}

	return cells
}

func toArray(cells []int) [8]int {
	return [8]int{cells[0], cells[1], cells[2], cells[3], cells[4], cells[5], cells[6], cells[7]}
}

func nextDay(cells []int) []int {
	next := make([]int, 8)
	for i := 1; i < 7; i++ {
		if cells[i-1] == cells[i+1] {
			next[i] = 1
		}
	}
	return next
}

func main() {
	fmt.Println(prisonAfterNDays([]int{0, 1, 0, 1, 1, 0, 0, 1}, 7))
	fmt.Println(prisonAfterNDays([]int{1, 0, 0, 1, 0, 0, 1, 0}, 1000000000))
}
```

## 0958 — Check Completeness Of A Binary Tree

```go
package main

// LeetCode #958: Check Completeness of a Binary Tree
// https://leetcode.com/problems/check-completeness-of-a-binary-tree/
// Difficulty: Medium

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n) | Space: O(n)
func isCompleteTree(root *TreeNode) bool {
	q := []*TreeNode{root}
	seenNull := false
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node == nil {
			seenNull = true
		} else {
			if seenNull {
				return false
			}
			q = append(q, node.Left, node.Right)
		}
	}
	return true
}

func main() {
	// Complete tree: [1,2,3,4,5,6]
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 6},
		},
	}
	fmt.Println(isCompleteTree(root))

	// Not complete: [1,2,3,4,5,null,7]
	root2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val:   3,
			Right: &TreeNode{Val: 7},
		},
	}
	fmt.Println(isCompleteTree(root2))
}
```

## 0959 — Regions Cut By Slashes

```go
package main

// LeetCode #959: Regions Cut By Slashes
// https://leetcode.com/problems/regions-cut-by-slashes/
// Difficulty: Medium

import "fmt"

// Time: O(n^2 * α(n^2)) | Space: O(n^2)
func regionsBySlashes(grid []string) int {
	n := len(grid)
	size := n * n * 4
	parent := make([]int, size)
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
		ra, rb := find(a), find(b)
		if ra != rb {
			parent[rb] = ra
			size--
		}
	}

	for i, row := range grid {
		for j, ch := range row {
			base := (i*n + j) * 4
			// Connect internal triangles
			if ch == '/' {
				union(base, base+3)
				union(base+1, base+2)
			} else if ch == '\\' {
				union(base, base+1)
				union(base+2, base+3)
			} else {
				union(base, base+1)
				union(base+1, base+2)
				union(base+2, base+3)
			}
			// Connect with right neighbor
			if j+1 < n {
				union(base+1, (base+4)+3)
			}
			// Connect with bottom neighbor
			if i+1 < n {
				union(base+2, (base+4*n))
			}
		}
	}

	return size
}

func main() {
	fmt.Println(regionsBySlashes([]string{" /", "/ "}))
	fmt.Println(regionsBySlashes([]string{" /", "  "}))
	fmt.Println(regionsBySlashes([]string{"/\\", "\\/"}))
}
```

## 0962 — Maximum Width Ramp

```go
package main

// LeetCode #962: Maximum Width Ramp
// https://leetcode.com/problems/maximum-width-ramp/
// Difficulty: Medium

import "fmt"

// Time: O(n) | Space: O(n)
func maxWidthRamp(nums []int) int {
	n := len(nums)
	stack := make([]int, 0)

	// Build decreasing stack of indices
	for i := 0; i < n; i++ {
		if len(stack) == 0 || nums[stack[len(stack)-1]] > nums[i] {
			stack = append(stack, i)
		}
	}

	ans := 0
	for j := n - 1; j >= 0; j-- {
		for len(stack) > 0 && nums[stack[len(stack)-1]] <= nums[j] {
			width := j - stack[len(stack)-1]
			if width > ans {
				ans = width
			}
			stack = stack[:len(stack)-1]
		}
	}
	return ans
}

func main() {
	fmt.Println(maxWidthRamp([]int{6, 0, 8, 2, 1, 5}))
	fmt.Println(maxWidthRamp([]int{9, 8, 1, 0, 1, 9, 4, 0, 4, 1}))
}
```

## 0963 — Minimum Area Rectangle Ii

```go
package main

// LeetCode #963: Minimum Area Rectangle II
// https://leetcode.com/problems/minimum-area-rectangle-ii/
// Difficulty: Medium

import (
	"fmt"
	"math"
)

// Time: O(n^3) | Space: O(n)
func minAreaFreeRect(points [][]int) float64 {
	n := len(points)
	set := make(map[[2]int]bool)
	for _, p := range points {
		set[[2]int{p[0], p[1]}] = true
	}

	ans := math.MaxFloat64

	for i := 0; i < n; i++ {
		x1, y1 := points[i][0], points[i][1]
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			x2, y2 := points[j][0], points[j][1]
			dx1, dy1 := x2-x1, y2-y1
			for k := j + 1; k < n; k++ {
				if k == i {
					continue
				}
				x3, y3 := points[k][0], points[k][1]
				dx2, dy2 := x3-x1, y3-y1

				// Check perpendicular
				if dx1*dx2+dy1*dy2 != 0 {
					continue
				}

				// Fourth point: (x2 + dx2, y2 + dy2)
				x4, y4 := x2+dx2, y2+dy2
				if !set[[2]int{x4, y4}] {
					continue
				}

				area := math.Sqrt(float64(dx1*dx1+dy1*dy1)) * math.Sqrt(float64(dx2*dx2+dy2*dy2))
				if area < ans {
					ans = area
				}
			}
		}
	}

	if ans == math.MaxFloat64 {
		return 0
	}
	return ans
}

func main() {
	fmt.Println(minAreaFreeRect([][]int{{1, 2}, {2, 1}, {1, 0}, {0, 1}}))
	fmt.Println(minAreaFreeRect([][]int{{0, 1}, {1, 0}, {1, 2}, {2, 1}, {1, 1}}))
}
```

## 0966 — Vowel Spellchecker

```go
package main

// LeetCode #966: Vowel Spellchecker
// https://leetcode.com/problems/vowel-spellchecker/
// Difficulty: Medium

import (
	"fmt"
	"strings"
)

// Time: O(n * L) | Space: O(n * L)
func spellchecker(wordlist []string, queries []string) []string {
	exact := make(map[string]bool)
	lower := make(map[string]string)
	vowel := make(map[string]string)

	for _, w := range wordlist {
		exact[w] = true
		lo := strings.ToLower(w)
		if _, ok := lower[lo]; !ok {
			lower[lo] = w
		}
		vw := devowel(lo)
		if _, ok := vowel[vw]; !ok {
			vowel[vw] = w
		}
	}

	ans := make([]string, len(queries))
	for i, q := range queries {
		if exact[q] {
			ans[i] = q
		} else if w, ok := lower[strings.ToLower(q)]; ok {
			ans[i] = w
		} else if w, ok := vowel[devowel(strings.ToLower(q))]; ok {
			ans[i] = w
		} else {
			ans[i] = ""
		}
	}
	return ans
}

func devowel(s string) string {
	var sb strings.Builder
	for _, ch := range s {
		if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
			sb.WriteByte('*')
		} else {
			sb.WriteRune(ch)
		}
	}
	return sb.String()
}

func main() {
	fmt.Println(spellchecker([]string{"KiTe", "kite", "hare", "Hare"}, []string{"kite", "Kite", "KiTe", "Hare", "HARE", "Hear", "hear", "keti", "keet", "keto"}))
	fmt.Println(spellchecker([]string{"yellow", "wood"}, []string{"Yello", "wood", "yellow"}))
}
```

## 0967 — Numbers With Same Consecutive Differences

```go
package main

// LeetCode #967: Numbers With Same Consecutive Differences
// https://leetcode.com/problems/numbers-with-same-consecutive-differences/
// Difficulty: Medium

import "fmt"

// Time: O(n * 2^n) | Space: O(n * 2^n)
func numsSameConsecDiff(n int, k int) []int {
	if n == 1 {
		return []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	}

	ans := make([]int, 0)
	for d := 1; d <= 9; d++ {
		dfs(n, k, d, &ans)
	}
	return ans
}

func dfs(n, k, cur int, ans *[]int) {
	if n == 1 {
		*ans = append(*ans, cur)
		return
	}
	last := cur % 10
	if last+k <= 9 {
		dfs(n-1, k, cur*10+last+k, ans)
	}
	if k != 0 && last-k >= 0 {
		dfs(n-1, k, cur*10+last-k, ans)
	}
}

func main() {
	fmt.Println(numsSameConsecDiff(3, 7))
	fmt.Println(numsSameConsecDiff(2, 1))
	fmt.Println(numsSameConsecDiff(2, 0))
}
```

## 0969 — Pancake Sorting

```go
package main

// LeetCode #969: Pancake Sorting
// https://leetcode.com/problems/pancake-sorting/
// Difficulty: Medium

import "fmt"

// Time: O(n^2) | Space: O(n)
func pancakeSort(arr []int) []int {
	ans := make([]int, 0)
	n := len(arr)

	for i := n; i > 1; i-- {
		// Find index of max in arr[:i]
		maxIdx := 0
		for j := 1; j < i; j++ {
			if arr[j] > arr[maxIdx] {
				maxIdx = j
			}
		}
		if maxIdx == i-1 {
			continue
		}
		// Flip to bring max to front
		if maxIdx > 0 {
			reverse(arr, maxIdx)
			ans = append(ans, maxIdx+1)
		}
		// Flip to put max at correct position
		reverse(arr, i-1)
		ans = append(ans, i)
	}

	return ans
}

func reverse(arr []int, end int) {
	for i, j := 0, end; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func main() {
	fmt.Println(pancakeSort([]int{3, 2, 4, 1}))
	fmt.Println(pancakeSort([]int{1, 2, 3}))
}
```

## 0970 — Powerful Integers

```go
package main

// LeetCode #970: Powerful Integers
// https://leetcode.com/problems/powerful-integers/
// Difficulty: Medium

import "fmt"

// Time: O(log_x(bound) * log_y(bound)) | Space: O(log_x(bound) * log_y(bound))
func powerfulIntegers(x int, y int, bound int) []int {
	seen := make(map[int]bool)

	for a := 1; a <= bound; a *= x {
		for b := 1; a+b <= bound; b *= y {
			seen[a+b] = true
			if y == 1 {
				break
			}
		}
		if x == 1 {
			break
		}
	}

	ans := make([]int, 0, len(seen))
	for v := range seen {
		ans = append(ans, v)
	}
	return ans
}

func main() {
	fmt.Println(powerfulIntegers(2, 3, 10))
	fmt.Println(powerfulIntegers(3, 5, 15))
	fmt.Println(powerfulIntegers(2, 1, 10))
}
```

## 0971 — Flip Binary Tree To Match Preorder Traversal

```go
package main

// LeetCode #971: Flip Binary Tree To Match Preorder Traversal
// https://leetcode.com/problems/flip-binary-tree-to-match-preorder-traversal/
// Difficulty: Medium

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n) | Space: O(n)
func flipMatchVoyage(root *TreeNode, voyage []int) []int {
	i := 0
	ok := true
	ans := make([]int, 0)

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil || !ok {
			return
		}
		if node.Val != voyage[i] {
			ok = false
			return
		}
		i++
		if node.Left != nil && node.Left.Val != voyage[i] {
			ans = append(ans, node.Val)
			dfs(node.Right)
			dfs(node.Left)
		} else {
			dfs(node.Left)
			dfs(node.Right)
		}
	}

	dfs(root)
	if !ok {
		return []int{-1}
	}
	return ans
}

func main() {
	// [1,2] voyage=[2,1]
	root1 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	fmt.Println(flipMatchVoyage(root1, []int{2, 1}))

	// [1,2,3] voyage=[1,3,2]
	root2 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	fmt.Println(flipMatchVoyage(root2, []int{1, 3, 2}))

	// [1,2,3] voyage=[1,2,3]
	root3 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	fmt.Println(flipMatchVoyage(root3, []int{1, 2, 3}))
}
```

## 0973 — K Closest Points To Origin

```go
package main

// LeetCode #973: K Closest Points to Origin
// https://leetcode.com/problems/k-closest-points-to-origin/
// Difficulty: Medium
//
// Approach: QuickSelect (partition-based selection)
// Time: O(n) average, O(n^2) worst-case
// Space: O(1) excluding output

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(kClosest([][]int{{1, 3}, {-2, 2}}, 1))         // [[-2,2]]
	fmt.Println(kClosest([][]int{{3, 3}, {5, -1}, {-2, 4}}, 2)) // [[3,3],[-2,4]]
}

func kClosest(points [][]int, k int) [][]int {
	if len(points) <= k {
		return points
	}

	// QuickSelect
	left, right := 0, len(points)-1
	for left < right {
		pivot := partition(points, left, right)
		if pivot == k {
			break
		} else if pivot < k {
			left = pivot + 1
		} else {
			right = pivot - 1
		}
	}

	return points[:k]
}

func partition(points [][]int, left, right int) int {
	pivotIdx := left + rand.Intn(right-left+1)
	points[pivotIdx], points[right] = points[right], points[pivotIdx]
	pivotDist := dist(points[right])
	storeIdx := left

	for i := left; i < right; i++ {
		if dist(points[i]) <= pivotDist {
			points[storeIdx], points[i] = points[i], points[storeIdx]
			storeIdx++
		}
	}
	points[storeIdx], points[right] = points[right], points[storeIdx]
	return storeIdx
}

func dist(p []int) int {
	return p[0]*p[0] + p[1]*p[1]
}
```

## 0974 — Subarray Sums Divisible By K

```go
package main

// LeetCode #974: Subarray Sums Divisible by K
// https://leetcode.com/problems/subarray-sums-divisible-by-k/
// Difficulty: Medium
//
// Approach: Prefix sum + modulo counting
// Time: O(n)
// Space: O(k)

import "fmt"

func main() {
	fmt.Println(subarraysDivByK([]int{4, 5, 0, -2, -3, 1}, 5)) // 7
	fmt.Println(subarraysDivByK([]int{5}, 9))                  // 0
	fmt.Println(subarraysDivByK([]int{-1, 2, 9}, 2))           // 2
}

func subarraysDivByK(nums []int, k int) int {
	modCount := make(map[int]int)
	modCount[0] = 1
	prefixSum := 0
	result := 0

	for _, n := range nums {
		prefixSum += n
		mod := prefixSum % k
		if mod < 0 {
			mod += k
		}
		result += modCount[mod]
		modCount[mod]++
	}

	return result
}
```

## 0978 — Longest Turbulent Subarray

```go
package main

// LeetCode #978: Longest Turbulent Subarray
// https://leetcode.com/problems/longest-turbulent-subarray/
// Difficulty: Medium
//
// Approach: Sliding window
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(maxTurbulenceSize([]int{9, 4, 2, 10, 7, 8, 8, 1, 9})) // 5
	fmt.Println(maxTurbulenceSize([]int{4, 8, 12, 16}))                // 2
	fmt.Println(maxTurbulenceSize([]int{100}))                         // 1
}

func maxTurbulenceSize(arr []int) int {
	n := len(arr)
	if n < 2 {
		return n
	}

	left := 0
	result := 1

	for right := 1; right < n; right++ {
		cmp := sign(arr[right-1] - arr[right])
		if cmp == 0 {
			left = right
		} else if right == n-1 || cmp*sign(arr[right]-arr[right+1]) != -1 {
			if right-left+1 > result {
				result = right - left + 1
			}
			left = right
		}
	}

	return result
}

func sign(x int) int {
	if x > 0 {
		return 1
	} else if x < 0 {
		return -1
	}
	return 0
}
```

## 0979 — Distribute Coins In Binary Tree

```go
package main

// LeetCode #979: Distribute Coins in Binary Tree
// https://leetcode.com/problems/distribute-coins-in-binary-tree/
// Difficulty: Medium
//
// Approach: DFS - post-order traversal
// Time: O(n)
// Space: O(h) where h is tree height

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Example 1: [3,0,0]
	root1 := &TreeNode{3, &TreeNode{0, nil, nil}, &TreeNode{0, nil, nil}}
	fmt.Println(distributeCoins(root1)) // 2

	// Example 2: [0,3,0]
	root2 := &TreeNode{0, &TreeNode{3, nil, nil}, &TreeNode{0, nil, nil}}
	fmt.Println(distributeCoins(root2)) // 3
}

func distributeCoins(root *TreeNode) int {
	moves := 0
	dfs(root, &moves)
	return moves
}

func dfs(node *TreeNode, moves *int) int {
	if node == nil {
		return 0
	}
	left := dfs(node.Left, moves)
	right := dfs(node.Right, moves)
	*moves += abs(left) + abs(right)
	return node.Val + left + right - 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

## 0981 — Time Based Key Value Store

```go
package main

// LeetCode #981: Time Based Key-Value Store
// https://leetcode.com/problems/time-based-key-value-store/
// Difficulty: Medium
//
// Approach: Hash map of key -> sorted list of (timestamp, value) pairs + binary search
// Time: O(1) for set, O(log n) for get
// Space: O(n)

import (
	"fmt"
	"sort"
)

func main() {
	tv := Constructor()
	tv.Set("foo", "bar", 1)
	fmt.Println(tv.Get("foo", 1))  // "bar"
	fmt.Println(tv.Get("foo", 3))  // "bar"
	tv.Set("foo", "bar2", 4)
	fmt.Println(tv.Get("foo", 4))  // "bar2"
	fmt.Println(tv.Get("foo", 5))  // "bar2"
	fmt.Println(tv.Get("foo", 0))  // ""
}

type pair struct {
	timestamp int
	value     string
}

type TimeMap struct {
	store map[string][]pair
}

func Constructor() TimeMap {
	return TimeMap{store: make(map[string][]pair)}
}

func (tm *TimeMap) Set(key string, value string, timestamp int) {
	tm.store[key] = append(tm.store[key], pair{timestamp, value})
}

func (tm *TimeMap) Get(key string, timestamp int) string {
	pairs, ok := tm.store[key]
	if !ok || len(pairs) == 0 {
		return ""
	}

	idx := sort.Search(len(pairs), func(i int) bool {
		return pairs[i].timestamp > timestamp
	})

	if idx == 0 {
		return ""
	}
	return pairs[idx-1].value
}
```

## 0983 — Minimum Cost For Tickets

```go
package main

// LeetCode #983: Minimum Cost For Tickets
// https://leetcode.com/problems/minimum-cost-for-tickets/
// Difficulty: Medium
//
// Approach: DP (bottom-up) over travel days
// Time: O(n) where n is the range of days (last travel day)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(mincostTickets([]int{1, 4, 6, 7, 8, 20}, []int{2, 7, 15})) // 11
	fmt.Println(mincostTickets([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31}, []int{2, 7, 15})) // 17
	fmt.Println(mincostTickets([]int{1, 2, 3}, []int{2, 7, 15})) // 6
}

func mincostTickets(days []int, costs []int) int {
	lastDay := days[len(days)-1]
	dp := make([]int, lastDay+1)
	travelSet := make(map[int]bool)
	for _, d := range days {
		travelSet[d] = true
	}

	for i := 1; i <= lastDay; i++ {
		if !travelSet[i] {
			dp[i] = dp[i-1]
			continue
		}
		one := dp[i-1] + costs[0]
		seven := dp[max(0, i-7)] + costs[1]
		thirty := dp[max(0, i-30)] + costs[2]
		dp[i] = min(one, min(seven, thirty))
	}

	return dp[lastDay]
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
```

## 0984 — String Without Aaa Or Bbb

```go
package main

// LeetCode #984: String Without AAA or BBB
// https://leetcode.com/problems/string-without-aaa-or-bbb/
// Difficulty: Medium
//
// Approach: Greedy - always append the character with more remaining count,
//           but skip a triple by using the other character if we just wrote two.
// Time: O(a + b)
// Space: O(a + b) for output

import "fmt"

func main() {
	fmt.Println(strWithout3a3b(1, 2)) // "bba" or "bab"
	fmt.Println(strWithout3a3b(4, 1)) // "aabaa"
	fmt.Println(strWithout3a3b(3, 3)) // "ababa" or similar
}

func strWithout3a3b(a int, b int) string {
	result := make([]byte, 0, a+b)

	for a > 0 || b > 0 {
		writeA := false
		if a > b {
			writeA = true
		} else if a < b {
			writeA = false
		} else {
			// equal counts: prefer the one that won't create triple
			if len(result) >= 2 && result[len(result)-1] == 'a' && result[len(result)-2] == 'a' {
				writeA = false
			} else if len(result) >= 2 && result[len(result)-1] == 'b' && result[len(result)-2] == 'b' {
				writeA = true
			} else {
				writeA = true
			}
		}

		if writeA {
			result = append(result, 'a')
			a--
			if a > 0 && a >= b {
				result = append(result, 'a')
				a--
			}
		} else {
			result = append(result, 'b')
			b--
			if b > 0 && b >= a {
				result = append(result, 'b')
				b--
			}
		}
	}

	return string(result)
}
```

## 0985 — Sum Of Even Numbers After Queries

```go
package main

// LeetCode #985: Sum of Even Numbers After Queries
// https://leetcode.com/problems/sum-of-even-numbers-after-queries/
// Difficulty: Medium
//
// Approach: Maintain running sum of even numbers, update incrementally
// Time: O(n + q) where n = len(nums), q = len(queries)
// Space: O(1) excluding output

import "fmt"

func main() {
	fmt.Println(sumEvenAfterQueries([]int{1, 2, 3, 4}, [][]int{{1, 0}, {-3, 1}, {-4, 0}, {2, 3}})) // [8,6,2,4]
	fmt.Println(sumEvenAfterQueries([]int{1}, [][]int{{4, 0}}))                                  // [0]
}

func sumEvenAfterQueries(nums []int, queries [][]int) []int {
	result := make([]int, len(queries))
	sum := 0

	for _, n := range nums {
		if n%2 == 0 {
			sum += n
		}
	}

	for i, q := range queries {
		val, idx := q[0], q[1]
		if nums[idx]%2 == 0 {
			sum -= nums[idx]
		}
		nums[idx] += val
		if nums[idx]%2 == 0 {
			sum += nums[idx]
		}
		result[i] = sum
	}

	return result
}
```

## 0986 — Interval List Intersections

```go
package main

// LeetCode #986: Interval List Intersections
// https://leetcode.com/problems/interval-list-intersections/
// Difficulty: Medium
//
// Approach: Two pointers
// Time: O(m + n)
// Space: O(1) excluding output

import "fmt"

func main() {
	fmt.Println(intervalIntersection([][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}}, [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}}))
	// [[1,2],[5,5],[8,10],[15,23],[24,24],[25,25]]
	fmt.Println(intervalIntersection([][]int{{1, 3}, {5, 9}}, [][]int{}))
	// []
	fmt.Println(intervalIntersection([][]int{{1, 7}}, [][]int{{3, 10}}))
	// [[3,7]]
}

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	i, j := 0, 0
	result := [][]int{}

	for i < len(firstList) && j < len(secondList) {
		start := max(firstList[i][0], secondList[j][0])
		end := min(firstList[i][1], secondList[j][1])

		if start <= end {
			result = append(result, []int{start, end})
		}

		if firstList[i][1] < secondList[j][1] {
			i++
		} else {
			j++
		}
	}

	if len(result) == 0 {
		return [][]int{}
	}
	return result
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
```

## 0988 — Smallest String Starting From Leaf

```go
package main

// LeetCode #988: Smallest String Starting From Leaf
// https://leetcode.com/problems/smallest-string-starting-from-leaf/
// Difficulty: Medium
//
// Approach: DFS backtracking from root to leaf, compare strings lexicographically
// Time: O(n * h) worst-case where h is tree height
// Space: O(h)

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Example: [0,1,2,3,4,3,4]
	root := &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 3, Left: nil, Right: nil},
			Right: &TreeNode{Val: 4, Left: nil, Right: nil},
		},
		Right: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 3, Left: nil, Right: nil},
			Right: &TreeNode{Val: 4, Left: nil, Right: nil},
		},
	}
	fmt.Println(smallestFromLeaf(root)) // "dba"

	// Example: [25,1,null,0,0,1,null,null,null,0]
	root2 := &TreeNode{
		Val: 25,
		Left: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 0, Left: nil, Right: nil},
			Right: &TreeNode{Val: 0, Left: &TreeNode{Val: 1, Left: nil, Right: &TreeNode{Val: 0, Left: nil, Right: nil}}, Right: nil},
		},
		Right: nil,
	}
	fmt.Println(smallestFromLeaf(root2)) // "abz"
}

func smallestFromLeaf(root *TreeNode) string {
	result := ""
	dfs(root, []byte{}, &result)
	return result
}

func dfs(node *TreeNode, buf []byte, result *string) {
	if node == nil {
		return
	}

	buf = append(buf, byte('a'+node.Val))

	if node.Left == nil && node.Right == nil {
		s := reverse(string(buf))
		if *result == "" || s < *result {
			*result = s
		}
		return
	}

	if node.Left != nil {
		dfs(node.Left, buf, result)
	}
	if node.Right != nil {
		dfs(node.Right, buf, result)
	}
}

func reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}
```

## 0990 — Satisfiability Of Equality Equations

```go
package main

// LeetCode #990: Satisfiability of Equality Equations
// https://leetcode.com/problems/satisfiability-of-equality-equations/
// Difficulty: Medium
//
// Approach: Union-Find (DSU)
// Time: O(n * alpha(N)) where n = len(equations), alpha is inverse Ackermann
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(equationsPossible([]string{"a==b", "b!=a"}))                           // false
	fmt.Println(equationsPossible([]string{"b==a", "a==b"}))                           // true
	fmt.Println(equationsPossible([]string{"a==b", "b==c", "a==c"}))                   // true
	fmt.Println(equationsPossible([]string{"a==b", "b!=c", "c==a"}))                   // false
}

func equationsPossible(equations []string) bool {
	parent := make([]int, 26)
	for i := 0; i < 26; i++ {
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

	// Process all equalities first
	for _, eq := range equations {
		if eq[1] == '=' {
			union(int(eq[0]-'a'), int(eq[3]-'a'))
		}
	}

	// Then check all inequalities
	for _, eq := range equations {
		if eq[1] == '!' {
			if find(int(eq[0]-'a')) == find(int(eq[3]-'a')) {
				return false
			}
		}
	}

	return true
}
```

## 0991 — Broken Calculator

```go
package main

// LeetCode #991: Broken Calculator
// https://leetcode.com/problems/broken-calculator/
// Difficulty: Medium
//
// Approach: Work backwards from target to startValue
//   - If target is even, divide by 2 (reverse of multiply by 2)
//   - If target is odd, add 1 (reverse of subtract 1)
// Time: O(log target)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(brokenCalc(2, 3))  // 2
	fmt.Println(brokenCalc(5, 8))  // 2
	fmt.Println(brokenCalc(3, 10)) // 3
}

func brokenCalc(startValue int, target int) int {
	ops := 0
	for target > startValue {
		if target%2 == 0 {
			target /= 2
		} else {
			target++
		}
		ops++
	}
	return ops + (startValue - target)
}
```

## 0994 — Rotting Oranges

```go
package main

// LeetCode #994: Rotting Oranges
// https://leetcode.com/problems/rotting-oranges/
// Difficulty: Medium
//
// Approach: BFS from all rotten oranges simultaneously
// Time: O(m * n)
// Space: O(m * n)

import "fmt"

func main() {
	fmt.Println(orangesRotting([][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}})) // 4
	fmt.Println(orangesRotting([][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}})) // -1
	fmt.Println(orangesRotting([][]int{{0, 2}}))                          // 0
}

func orangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	queue := make([][2]int, 0)
	fresh := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, [2]int{i, j})
			} else if grid[i][j] == 1 {
				fresh++
			}
		}
	}

	if fresh == 0 {
		return 0
	}

	minutes := 0
	for len(queue) > 0 {
		minutes++
		size := len(queue)
		for k := 0; k < size; k++ {
			cur := queue[k]
			for _, d := range dirs {
				ni, nj := cur[0]+d[0], cur[1]+d[1]
				if ni >= 0 && ni < m && nj >= 0 && nj < n && grid[ni][nj] == 1 {
					grid[ni][nj] = 2
					fresh--
					queue = append(queue, [2]int{ni, nj})
				}
			}
		}
		queue = queue[size:]
	}

	if fresh > 0 {
		return -1
	}
	return minutes - 1
}
```

## 0998 — Maximum Binary Tree Ii

```go
package main

// LeetCode #998: Maximum Binary Tree II
// https://leetcode.com/problems/maximum-binary-tree-ii/
// Difficulty: Medium
//
// Approach: Since val is appended to the end of the original array,
//           if val > root.Val, it becomes the new root (with old root as left child).
//           Otherwise, recurse into the right subtree.
// Time: O(h) where h is tree height
// Space: O(h)

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Example: root = [4,1,3,null,null,2], val = 5
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{Val: 1, Left: nil, Right: nil},
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 2, Left: nil, Right: nil},
			Right: nil,
		},
	}
	result := insertIntoMaxTree(root, 5)
	printTree(result)
	fmt.Println()

	// val < root: root = [5,2,4,null,1], val = 3
	root2 := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{Val: 1, Left: nil, Right: nil},
		},
		Right: &TreeNode{Val: 4, Left: nil, Right: nil},
	}
	result2 := insertIntoMaxTree(root2, 3)
	printTree(result2)
	fmt.Println()
}

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if val > root.Val {
		return &TreeNode{Val: val, Left: root}
	}
	root.Right = insertIntoMaxTree(root.Right, val)
	return root
}

func printTree(root *TreeNode) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			fmt.Print("null ")
			continue
		}
		fmt.Printf("%d ", node.Val)
		queue = append(queue, node.Left, node.Right)
	}
}
```

## 1003 — Check If Word Is Valid After Substitutions

```go
package main

// LeetCode #1003: Check If Word Is Valid After Substitutions
// https://leetcode.com/problems/check-if-word-is-valid-after-substitutions/
// Difficulty: Medium
//
// Approach: Use a stack. When we see "c", check if top two are "a" and "b".
//           If so, they form "abc" and we pop them. Otherwise, push "c".
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(isValid("aabcbc"))  // true
	fmt.Println(isValid("abcabcababcc")) // true
	fmt.Println(isValid("abccba"))  // false
	fmt.Println(isValid("cababc"))  // false
}

func isValid(s string) bool {
	stack := make([]byte, 0, len(s))

	for i := 0; i < len(s); i++ {
		if s[i] == 'c' {
			n := len(stack)
			if n >= 2 && stack[n-1] == 'b' && stack[n-2] == 'a' {
				stack = stack[:n-2]
				continue
			}
		}
		stack = append(stack, s[i])
	}

	return len(stack) == 0
}
```

## 1004 — Max Consecutive Ones Iii

```go
package main

// LeetCode #1004: Max Consecutive Ones III
// https://leetcode.com/problems/max-consecutive-ones-iii/
// Difficulty: Medium
//
// Approach: Sliding window (expand right, shrink left when zeros > k)
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2)) // 6
	fmt.Println(longestOnes([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3)) // 10
	fmt.Println(longestOnes([]int{0, 0, 0, 0}, 0)) // 0
}

func longestOnes(nums []int, k int) int {
	left := 0
	zeros := 0
	result := 0

	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zeros++
		}

		for zeros > k {
			if nums[left] == 0 {
				zeros--
			}
			left++
		}

		if right-left+1 > result {
			result = right - left + 1
		}
	}

	return result
}
```

## 1006 — Clumsy Factorial

```go
package main

// LeetCode #1006: Clumsy Factorial
// https://leetcode.com/problems/clumsy-factorial/
// Difficulty: Medium
//
// Approach: Math - use stack to handle operator precedence (*, / before +, -)
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(clumsy(4))  // 7
	fmt.Println(clumsy(10)) // 12
	fmt.Println(clumsy(1))  // 1
}

func clumsy(n int) int {
	stack := []int{n}
	op := 0

	for i := n - 1; i >= 1; i-- {
		switch op % 4 {
		case 0: // *
			stack[len(stack)-1] *= i
		case 1: // /
			stack[len(stack)-1] /= i
		case 2: // +
			stack = append(stack, i)
		case 3: // -
			stack = append(stack, -i)
		}
		op++
	}

	result := 0
	for _, v := range stack {
		result += v
	}
	return result
}
```

## 1007 — Minimum Domino Rotations For Equal Row

```go
package main

// LeetCode #1007: Minimum Domino Rotations For Equal Row
// https://leetcode.com/problems/minimum-domino-rotations-for-equal-row/
// Difficulty: Medium
//
// Approach: Check if we can make all values equal to tops[0] or bottoms[0]
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(minDominoRotations([]int{2, 1, 2, 4, 2, 2}, []int{5, 2, 6, 2, 3, 2})) // 2
	fmt.Println(minDominoRotations([]int{3, 5, 1, 2, 3}, []int{3, 6, 3, 3, 4}))       // -1
	fmt.Println(minDominoRotations([]int{1, 2, 1, 1, 1, 2, 2, 2}, []int{2, 1, 2, 2, 2, 1, 1, 1})) // 2
}

func minDominoRotations(tops []int, bottoms []int) int {
	result := check(tops, bottoms, tops[0])
	if result != -1 {
		return result
	}
	return check(tops, bottoms, bottoms[0])
}

func check(tops, bottoms []int, target int) int {
	topRot := 0
	bottomRot := 0

	for i := 0; i < len(tops); i++ {
		if tops[i] != target && bottoms[i] != target {
			return -1
		}
		if tops[i] != target {
			topRot++
		}
		if bottoms[i] != target {
			bottomRot++
		}
	}

	if topRot < bottomRot {
		return topRot
	}
	return bottomRot
}
```

## 1008 — Construct Binary Search Tree From Preorder Traversal

```go
package main

// LeetCode #1008: Construct Binary Search Tree from Preorder Traversal
// https://leetcode.com/problems/construct-binary-search-tree-from-preorder-traversal/
// Difficulty: Medium
//
// Approach: Use upper bound recursion. First element is root. Recursively build
//           left subtree with upper bound = root.Val, then right subtree.
// Time: O(n)
// Space: O(h) where h is tree height

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	result := bstFromPreorder([]int{8, 5, 1, 7, 10, 12})
	printTree(result)
	fmt.Println()

	result2 := bstFromPreorder([]int{1, 3})
	printTree(result2)
	fmt.Println()
}

func bstFromPreorder(preorder []int) *TreeNode {
	idx := 0
	return build(preorder, &idx, 1<<31-1)
}

func build(preorder []int, idx *int, bound int) *TreeNode {
	if *idx >= len(preorder) || preorder[*idx] > bound {
		return nil
	}

	node := &TreeNode{Val: preorder[*idx]}
	*idx++

	node.Left = build(preorder, idx, node.Val)
	node.Right = build(preorder, idx, bound)

	return node
}

func printTree(root *TreeNode) {
	if root == nil {
		fmt.Print("[]")
		return
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			fmt.Print("null ")
			continue
		}
		fmt.Printf("%d ", node.Val)
		queue = append(queue, node.Left, node.Right)
	}
}
```

## 1010 — Pairs Of Songs With Total Durations Divisible By 60

```go
package main

// LeetCode #1010: Pairs of Songs With Total Durations Divisible by 60
// https://leetcode.com/problems/pairs-of-songs-with-total-durations-divisible-by-60/
// Difficulty: Medium
//
// Approach: Use modulo counting like Two Sum
// Time: O(n)
// Space: O(60) = O(1)

import "fmt"

func main() {
	fmt.Println(numPairsDivisibleBy60([]int{30, 20, 150, 100, 40})) // 3
	fmt.Println(numPairsDivisibleBy60([]int{60, 60, 60}))           // 3
}

func numPairsDivisibleBy60(time []int) int {
	count := make([]int, 60)
	result := 0

	for _, t := range time {
		mod := t % 60
		need := (60 - mod) % 60
		result += count[need]
		count[mod]++
	}

	return result
}
```

## 1011 — Capacity To Ship Packages Within D Days

```go
package main

// LeetCode #1011: Capacity To Ship Packages Within D Days
// https://leetcode.com/problems/capacity-to-ship-packages-within-d-days/
// Difficulty: Medium
//
// Approach: Binary search on capacity
// Time: O(n * log(sum(weights)))
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(shipWithinDays([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5)) // 15
	fmt.Println(shipWithinDays([]int{3, 2, 2, 4, 1, 4}, 3))             // 6
	fmt.Println(shipWithinDays([]int{1, 2, 3, 1, 1}, 4))                // 3
}

func shipWithinDays(weights []int, days int) int {
	left, right := 0, 0
	for _, w := range weights {
		if w > left {
			left = w
		}
		right += w
	}

	for left < right {
		mid := left + (right-left)/2
		if canShip(weights, days, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func canShip(weights []int, days int, capacity int) bool {
	dayCount := 1
	current := 0
	for _, w := range weights {
		if current+w > capacity {
			dayCount++
			current = w
			if dayCount > days {
				return false
			}
		} else {
			current += w
		}
	}
	return true
}
```

## 1014 — Best Sightseeing Pair

```go
package main

// LeetCode #1014: Best Sightseeing Pair
// https://leetcode.com/problems/best-sightseeing-pair/
// Difficulty: Medium
//
// Approach: Track max value of (values[i] + i) seen so far
// Score = values[i] + values[j] + i - j = (values[i] + i) + (values[j] - j)
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(maxScoreSightseeingPair([]int{8, 1, 5, 2, 6})) // 11
	fmt.Println(maxScoreSightseeingPair([]int{1, 2}))           // 2
}

func maxScoreSightseeingPair(values []int) int {
	maxI := values[0]
	result := 0

	for j := 1; j < len(values); j++ {
		if maxI+values[j]-j > result {
			result = maxI + values[j] - j
		}
		if values[j]+j > maxI {
			maxI = values[j] + j
		}
	}

	return result
}
```

## 1015 — Smallest Integer Divisible By K

```go
package main

// LeetCode #1015: Smallest Integer Divisible by K
// https://leetcode.com/problems/smallest-integer-divisible-by-k/
// Difficulty: Medium
//
// Approach: Build remainder incrementally. Keep dividing.
//   N = 1, 11, 111, ... = N*10 + 1 (mod K)
//   If N % K == 0, return length. If remainder repeats, return -1.
// Time: O(K)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(smallestRepunitDivByK(1))  // 1
	fmt.Println(smallestRepunitDivByK(2))  // -1
	fmt.Println(smallestRepunitDivByK(3))  // 3
}

func smallestRepunitDivByK(k int) int {
	if k%2 == 0 || k%5 == 0 {
		return -1
	}

	remainder := 0
	for length := 1; length <= k; length++ {
		remainder = (remainder*10 + 1) % k
		if remainder == 0 {
			return length
		}
	}

	return -1
}
```

## 1016 — Binary String With Substrings Representing 1 To N

```go
package main

// LeetCode #1016: Binary String With Substrings Representing 1 To N
// https://leetcode.com/problems/binary-string-with-substrings-representing-1-to-n/
// Difficulty: Medium
//
// Approach: Check if binary representation of each number from n down to 1
//           is a substring of s. Start from n and go down for efficiency.
// Time: O(n * len(s) * log n)
// Space: O(log n)

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(queryString("0110", 3))  // true
	fmt.Println(queryString("0110", 4))  // false
	fmt.Println(queryString("1", 1))     // true
}

func queryString(s string, n int) bool {
	for i := n; i >= 1; i-- {
		binary := fmt.Sprintf("%b", i)
		if !strings.Contains(s, binary) {
			return false
		}
	}
	return true
}
```

## 1017 — Convert To Base 2

```go
package main

// LeetCode #1017: Convert to Base -2
// https://leetcode.com/problems/convert-to-base-2/
// Difficulty: Medium
//
// Approach: Repeated division by -2. Handle negative remainder.
// Time: O(log n)
// Space: O(log n)

import "fmt"

func main() {
	fmt.Println(baseNeg2(2))  // "110"
	fmt.Println(baseNeg2(3))  // "111"
	fmt.Println(baseNeg2(4))  // "100"
}

func baseNeg2(n int) string {
	if n == 0 {
		return "0"
	}

	result := ""
	for n != 0 {
		remainder := n % -2
		n /= -2
		if remainder < 0 {
			remainder += 2
			n++
		}
		result = string(rune('0'+remainder)) + result
	}

	return result
}
```

## 1019 — Next Greater Node In Linked List

```go
package main

// LeetCode #1019: Next Greater Node In Linked List
// https://leetcode.com/problems/next-greater-node-in-linked-list/
// Difficulty: Medium
//
// Approach: Convert linked list to array, then use monotonic stack
// Time: O(n)
// Space: O(n)

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// [2,1,5]
	head := &ListNode{2, &ListNode{1, &ListNode{5, nil}}}
	fmt.Println(nextLargerNodes(head)) // [5,5,0]

	// [2,7,4,3,5]
	head2 := &ListNode{2, &ListNode{7, &ListNode{4, &ListNode{3, &ListNode{5, nil}}}}}
	fmt.Println(nextLargerNodes(head2)) // [7,0,5,5,0]
}

func nextLargerNodes(head *ListNode) []int {
	// Convert to array
	vals := make([]int, 0)
	for cur := head; cur != nil; cur = cur.Next {
		vals = append(vals, cur.Val)
	}

	result := make([]int, len(vals))
	stack := make([]int, 0) // indices

	for i, v := range vals {
		for len(stack) > 0 && vals[stack[len(stack)-1]] < v {
			result[stack[len(stack)-1]] = v
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	// Remaining in stack have no next greater (already zero-initialized)
	return result
}
```

## 1020 — Number Of Enclaves

```go
package main

// LeetCode #1020: Number of Enclaves
// https://leetcode.com/problems/number-of-enclaves/
// Difficulty: Medium
//
// Approach: DFS from border cells, mark reachable land, then count remaining
// Time: O(m * n)
// Space: O(m * n)

import "fmt"

func main() {
	fmt.Println(numEnclaves([][]int{{0, 0, 0, 0}, {1, 0, 1, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}})) // 3
	fmt.Println(numEnclaves([][]int{{0, 1, 1, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 0}})) // 0
}

func numEnclaves(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != 1 {
			return
		}
		grid[i][j] = 0
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}

	for i := 0; i < m; i++ {
		dfs(i, 0)
		dfs(i, n-1)
	}
	for j := 0; j < n; j++ {
		dfs(0, j)
		dfs(m-1, j)
	}

	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				count++
			}
		}
	}
	return count
}
```

## 1023 — Camelcase Matching

```go
package main

// LeetCode #1023: Camelcase Matching
// https://leetcode.com/problems/camelcase-matching/
// Difficulty: Medium
//
// Approach: For each query, use two pointers to match pattern characters
// Time: O(n * (len(query) + len(pattern)))
// Space: O(n) for output

import "fmt"

func main() {
	fmt.Println(camelMatch([]string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"}, "FB"))
	// [true,false,true,true,false]
	fmt.Println(camelMatch([]string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"}, "FoBa"))
	// [true,false,true,false,false]
}

func camelMatch(queries []string, pattern string) []bool {
	result := make([]bool, len(queries))

	for i, q := range queries {
		result[i] = matches(q, pattern)
	}

	return result
}

func matches(query, pattern string) bool {
	j := 0
	for i := 0; i < len(query); i++ {
		if j < len(pattern) && query[i] == pattern[j] {
			j++
		} else if query[i] >= 'A' && query[i] <= 'Z' {
			return false
		}
	}
	return j == len(pattern)
}
```

## 1024 — Video Stitching

```go
package main

// LeetCode #1024: Video Stitching
// https://leetcode.com/problems/video-stitching/
// Difficulty: Medium
//
// Approach: Sort clips by start, then greedy extend reach
// Time: O(n log n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(videoStitching([][]int{{0, 2}, {4, 6}, {8, 10}, {1, 9}, {1, 5}, {5, 9}}, 10)) // 3
	fmt.Println(videoStitching([][]int{{0, 1}, {6, 8}, {0, 2}, {5, 6}, {0, 4}, {0, 3}, {6, 7}, {1, 3}, {4, 7}, {1, 4}, {2, 5}, {2, 6}, {3, 4}, {4, 5}, {5, 7}, {6, 9}}, 9)) // 3
	fmt.Println(videoStitching([][]int{{0, 2}, {2, 4}}, 4)) // 2
}

func videoStitching(clips [][]int, time int) int {
	// Sort by start time, then by end time descending
	for i := 0; i < len(clips); i++ {
		for j := i + 1; j < len(clips); j++ {
			if clips[i][0] > clips[j][0] || (clips[i][0] == clips[j][0] && clips[i][1] < clips[j][1]) {
				clips[i], clips[j] = clips[j], clips[i]
			}
		}
	}

	count := 0
	curEnd := 0
	i := 0
	n := len(clips)

	for curEnd < time {
		farthest := curEnd
		for i < n && clips[i][0] <= curEnd {
			if clips[i][1] > farthest {
				farthest = clips[i][1]
			}
			i++
		}
		if farthest == curEnd {
			return -1
		}
		curEnd = farthest
		count++
	}

	return count
}
```

## 1026 — Maximum Difference Between Node And Ancestor

```go
package main

// LeetCode #1026: Maximum Difference Between Node and Ancestor
// https://leetcode.com/problems/maximum-difference-between-node-and-ancestor/
// Difficulty: Medium
//
// Approach: DFS tracking min and max values on path from root to leaf
// Time: O(n)
// Space: O(h) where h is tree height

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 8,
		Left: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 1, Left: nil, Right: nil},
			Right: &TreeNode{Val: 6, Left: &TreeNode{Val: 4, Left: nil, Right: nil}, Right: &TreeNode{Val: 7, Left: nil, Right: nil}},
		},
		Right: &TreeNode{
			Val:   10,
			Left:  nil,
			Right: &TreeNode{Val: 14, Left: &TreeNode{Val: 13, Left: nil, Right: nil}, Right: nil},
		},
	}
	fmt.Println(maxAncestorDiff(root)) // 7

	root2 := &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 0, Left: &TreeNode{Val: 3, Left: nil, Right: nil}, Right: nil}}}
	fmt.Println(maxAncestorDiff(root2)) // 3
}

func maxAncestorDiff(root *TreeNode) int {
	result := 0
	dfs(root, root.Val, root.Val, &result)
	return result
}

func dfs(node *TreeNode, minVal, maxVal int, result *int) {
	if node == nil {
		return
	}

	diff := node.Val - minVal
	if diff < 0 {
		diff = -diff
	}
	if diff > *result {
		*result = diff
	}
	diff = node.Val - maxVal
	if diff < 0 {
		diff = -diff
	}
	if diff > *result {
		*result = diff
	}

	if node.Val < minVal {
		minVal = node.Val
	}
	if node.Val > maxVal {
		maxVal = node.Val
	}

	dfs(node.Left, minVal, maxVal, result)
	dfs(node.Right, minVal, maxVal, result)
}
```

## 1027 — Longest Arithmetic Subsequence

```go
package main

// LeetCode #1027: Longest Arithmetic Subsequence
// https://leetcode.com/problems/longest-arithmetic-subsequence/
// Difficulty: Medium
//
// Approach: DP with hash map per index tracking difference -> length
// Time: O(n^2)
// Space: O(n^2)

import "fmt"

func main() {
	fmt.Println(longestArithSeqLength([]int{3, 6, 9, 12}))    // 4
	fmt.Println(longestArithSeqLength([]int{9, 4, 7, 2, 10})) // 3
	fmt.Println(longestArithSeqLength([]int{20, 1, 15, 3, 10, 5, 8})) // 4
}

func longestArithSeqLength(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}

	dp := make([]map[int]int, n)
	result := 2

	for i := 0; i < n; i++ {
		dp[i] = make(map[int]int)
		for j := 0; j < i; j++ {
			diff := nums[i] - nums[j]
			length := 2
			if prev, ok := dp[j][diff]; ok {
				length = prev + 1
			}
			dp[i][diff] = length
			if length > result {
				result = length
			}
		}
	}

	return result
}
```

## 1029 — Two City Scheduling

```go
package main

// LeetCode #1029: Two City Scheduling
// https://leetcode.com/problems/two-city-scheduling/
// Difficulty: Medium
//
// Approach: Sort by difference (costA - costB), send first half to A, rest to B
// Time: O(n log n)
// Space: O(1)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(twoCitySchedCost([][]int{{10, 20}, {30, 200}, {400, 50}, {30, 20}})) // 110
	fmt.Println(twoCitySchedCost([][]int{{259, 770}, {448, 54}, {926, 667}, {184, 139}, {840, 118}, {577, 469}})) // 1859
}

func twoCitySchedCost(costs [][]int) int {
	sort.Slice(costs, func(i, j int) bool {
		return (costs[i][0] - costs[i][1]) < (costs[j][0] - costs[j][1])
	})

	n := len(costs) / 2
	total := 0
	for i := 0; i < n; i++ {
		total += costs[i][0]
	}
	for i := n; i < 2*n; i++ {
		total += costs[i][1]
	}
	return total
}
```

## 1031 — Maximum Sum Of Two Non Overlapping Subarrays

```go
package main

// LeetCode #1031: Maximum Sum of Two Non-Overlapping Subarrays
// https://leetcode.com/problems/maximum-sum-of-two-non-overlapping-subarrays/
// Difficulty: Medium
//
// Approach: DP with prefix sums. Consider both orders (L then M, M then L)
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(maxSumTwoNoOverlap([]int{0, 6, 5, 2, 2, 5, 1, 9, 4}, 1, 2)) // 20
	fmt.Println(maxSumTwoNoOverlap([]int{3, 8, 1, 3, 2, 1, 8, 9, 0}, 3, 2)) // 29
	fmt.Println(maxSumTwoNoOverlap([]int{2, 1, 5, 6, 0, 9, 5, 0, 3, 8}, 4, 3)) // 31
}

func maxSumTwoNoOverlap(nums []int, firstLen int, secondLen int) int {
	n := len(nums)
	prefix := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + nums[i]
	}

	// Try (firstLen before secondLen) and (secondLen before firstLen)
	result := 0
	// Case 1: firstLen comes first
	leftMax := 0
	for i := firstLen; i <= n-secondLen; i++ {
		leftSum := prefix[i] - prefix[i-firstLen]
		if leftSum > leftMax {
			leftMax = leftSum
		}
		rightSum := prefix[i+secondLen] - prefix[i]
		if leftMax+rightSum > result {
			result = leftMax + rightSum
		}
	}

	// Case 2: secondLen comes first
	leftMax = 0
	for i := secondLen; i <= n-firstLen; i++ {
		leftSum := prefix[i] - prefix[i-secondLen]
		if leftSum > leftMax {
			leftMax = leftSum
		}
		rightSum := prefix[i+firstLen] - prefix[i]
		if leftMax+rightSum > result {
			result = leftMax + rightSum
		}
	}

	return result
}
```

## 1033 — Moving Stones Until Consecutive

```go
package main

// LeetCode #1033: Moving Stones Until Consecutive
// https://leetcode.com/problems/moving-stones-until-consecutive/
// Difficulty: Medium
//
// Approach: Sort the positions. Min moves = 0, 1, or 2 based on gaps.
//           Max moves = distance between extremes - 2.
// Time: O(1)
// Space: O(1)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(numMovesStones(1, 2, 5)) // [1,2]
	fmt.Println(numMovesStones(4, 3, 2)) // [0,0]
	fmt.Println(numMovesStones(3, 5, 1)) // [1,2]
}

func numMovesStones(a int, b int, c int) []int {
	stones := []int{a, b, c}
	sort.Ints(stones)
	x, y, z := stones[0], stones[1], stones[2]

	minMoves := 2
	if z-x == 2 {
		minMoves = 0
	} else if z-y <= 2 || y-x <= 2 {
		minMoves = 1
	}

	maxMoves := (z - x - 2)

	return []int{minMoves, maxMoves}
}
```

## 1034 — Coloring A Border

```go
package main

// LeetCode #1034: Coloring A Border
// https://leetcode.com/problems/coloring-a-border/
// Difficulty: Medium
//
// Approach: DFS to find connected component, then color border cells
// Time: O(m * n)
// Space: O(m * n)

import "fmt"

func main() {
	fmt.Println(colorBorder([][]int{{1, 1}, {1, 2}}, 0, 0, 3)) // [[3,3],[3,2]]
	fmt.Println(colorBorder([][]int{{1, 2, 2}, {2, 3, 2}}, 0, 1, 3)) // [[1,3,3],[2,3,3]]
}

func colorBorder(grid [][]int, row int, col int, color int) [][]int {
	m, n := len(grid), len(grid[0])
	origColor := grid[row][col]
	if origColor == color {
		return grid
	}

	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var dfs func(r, c int)
	dfs = func(r, c int) {
		visited[r][c] = true
		for _, d := range dirs {
			nr, nc := r+d[0], c+d[1]
			if nr < 0 || nr >= m || nc < 0 || nc >= n || grid[nr][nc] != origColor {
				grid[r][c] = color
			} else if !visited[nr][nc] {
				dfs(nr, nc)
			}
		}
	}

	dfs(row, col)
	return grid
}
```

## 1035 — Uncrossed Lines

```go
package main

// LeetCode #1035: Uncrossed Lines
// https://leetcode.com/problems/uncrossed-lines/
// Difficulty: Medium
//
// Approach: DP - Longest Common Subsequence (LCS)
// Time: O(m * n)
// Space: O(m * n)

import "fmt"

func main() {
	fmt.Println(maxUncrossedLines([]int{1, 4, 2}, []int{1, 2, 4}))    // 2
	fmt.Println(maxUncrossedLines([]int{2, 5, 1, 2, 5}, []int{10, 5, 2, 1, 5, 2})) // 3
}

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				if dp[i-1][j] > dp[i][j-1] {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}

	return dp[m][n]
}
```

## 1038 — Binary Search Tree To Greater Sum Tree

```go
package main

// LeetCode #1038: Binary Search Tree to Greater Sum Tree
// https://leetcode.com/problems/binary-search-tree-to-greater-sum-tree/
// Difficulty: Medium
//
// Approach: Reverse inorder (right -> root -> left) accumulating sum
// Time: O(n)
// Space: O(h) where h is tree height

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 0, Left: nil, Right: nil},
			Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3, Left: nil, Right: nil}},
		},
		Right: &TreeNode{
			Val:   6,
			Left:  &TreeNode{Val: 5, Left: nil, Right: nil},
			Right: &TreeNode{Val: 7, Right: &TreeNode{Val: 8, Left: nil, Right: nil}},
		},
	}
	result := bstToGst(root)
	printTree(result)
	fmt.Println()
}

func bstToGst(root *TreeNode) *TreeNode {
	sum := 0
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Right)
		sum += node.Val
		node.Val = sum
		dfs(node.Left)
	}
	dfs(root)
	return root
}

func printTree(root *TreeNode) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			fmt.Print("null ")
			continue
		}
		fmt.Printf("%d ", node.Val)
		queue = append(queue, node.Left, node.Right)
	}
}
```

## 1039 — Minimum Score Triangulation Of Polygon

```go
package main

// LeetCode #1039: Minimum Score Triangulation of Polygon
// https://leetcode.com/problems/minimum-score-triangulation-of-polygon/
// Difficulty: Medium
//
// Approach: Interval DP. dp[i][j] = min score triangulating polygon from i to j.
// Time: O(n^3)
// Space: O(n^2)

import "fmt"

func main() {
	fmt.Println(minScoreTriangulation([]int{1, 2, 3}))    // 6
	fmt.Println(minScoreTriangulation([]int{3, 7, 4, 5})) // 144
}

func minScoreTriangulation(values []int) int {
	n := len(values)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for length := 2; length < n; length++ {
		for i := 0; i+length < n; i++ {
			j := i + length
			dp[i][j] = 1<<31 - 1
			for k := i + 1; k < j; k++ {
				score := dp[i][k] + dp[k][j] + values[i]*values[j]*values[k]
				if score < dp[i][j] {
					dp[i][j] = score
				}
			}
		}
	}

	return dp[0][n-1]
}
```

## 1040 — Moving Stones Until Consecutive Ii

```go
package main

// LeetCode #1040: Moving Stones Until Consecutive II
// https://leetcode.com/problems/moving-stones-until-consecutive-ii/
// Difficulty: Medium
//
// Approach: Sort stones. Max moves: spread left or right. Min moves: sliding window.
// Time: O(n log n)
// Space: O(1)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(numMovesStonesII([]int{7, 4, 9}))    // [1,2]
	fmt.Println(numMovesStonesII([]int{6, 5, 4, 3, 10})) // [2,3]
}

func numMovesStonesII(stones []int) []int {
	sort.Ints(stones)
	n := len(stones)

	// Max moves: fill gaps from one end, leaving one stone at the other end
	maxMoves := max(stones[n-1]-stones[1]-(n-2), stones[n-2]-stones[0]-(n-2))

	// Min moves: sliding window of size n
	minMoves := n
	j := 0
	for i := 0; i < n; i++ {
		for j+1 < n && stones[j+1]-stones[i] < n {
			j++
		}
		already := j - i + 1
		moves := n - already
		// Special case: n-1 stones already consecutive, last one far away
		if moves == 1 && stones[j]-stones[i]+1 == n-1 && (j-i+1 == n-1) {
			moves = 2
		}
		if moves < minMoves {
			minMoves = moves
		}
	}

	return []int{minMoves, maxMoves}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

## 1041 — Robot Bounded In Circle

```go
package main

// LeetCode #1041: Robot Bounded In Circle
// https://leetcode.com/problems/robot-bounded-in-circle/
// Difficulty: Medium
//
// Approach: Simulate robot movement. Robot is bounded in circle iff
//           final position is (0,0) OR direction != North.
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(isRobotBounded("GGLLGG")) // true
	fmt.Println(isRobotBounded("GG"))     // false
	fmt.Println(isRobotBounded("GL"))     // true
}

func isRobotBounded(instructions string) bool {
	x, y := 0, 0
	dirX, dirY := 0, 1 // facing north

	for _, c := range instructions {
		switch c {
		case 'G':
			x += dirX
			y += dirY
		case 'L':
			dirX, dirY = -dirY, dirX
		case 'R':
			dirX, dirY = dirY, -dirX
		}
	}

	return (x == 0 && y == 0) || !(dirX == 0 && dirY == 1)
}
```

## 1042 — Flower Planting With No Adjacent

```go
package main

// LeetCode #1042: Flower Planting With No Adjacent
// https://leetcode.com/problems/flower-planting-with-no-adjacent/
// Difficulty: Medium
//
// Approach: Greedy coloring with a graph. Each garden gets a flower type
//           not used by its neighbors.
// Time: O(n + e) where e = len(paths)
// Space: O(n + e)

import "fmt"

func main() {
	fmt.Println(gardenNoAdj(3, [][]int{{1, 2}, {2, 3}, {3, 1}})) // [1,2,3]
	fmt.Println(gardenNoAdj(4, [][]int{{1, 2}, {3, 4}}))         // [1,2,1,2]
}

func gardenNoAdj(n int, paths [][]int) []int {
	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, 0)
	}

	for _, p := range paths {
		u, v := p[0]-1, p[1]-1
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	result := make([]int, n)
	for i := 0; i < n; i++ {
		used := make([]bool, 5)
		for _, neighbor := range graph[i] {
			if result[neighbor] != 0 {
				used[result[neighbor]] = true
			}
		}
		for flower := 1; flower <= 4; flower++ {
			if !used[flower] {
				result[i] = flower
				break
			}
		}
	}

	return result
}
```

## 1043 — Partition Array For Maximum Sum

```go
package main

// LeetCode #1043: Partition Array for Maximum Sum
// https://leetcode.com/problems/partition-array-for-maximum-sum/
// Difficulty: Medium
//
// Approach: DP. dp[i] = max sum for prefix ending at i.
//           For each i, try all partition lengths up to k.
// Time: O(n * k)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(maxSumAfterPartitioning([]int{1, 15, 7, 9, 2, 5, 10}, 3)) // 84
	fmt.Println(maxSumAfterPartitioning([]int{1, 4, 1, 5, 7, 3, 6, 1, 9, 9, 3}, 4)) // 83
}

func maxSumAfterPartitioning(arr []int, k int) int {
	n := len(arr)
	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		maxVal := 0
		for j := 1; j <= k && i-j >= 0; j++ {
			if arr[i-j] > maxVal {
				maxVal = arr[i-j]
			}
			sum := dp[i-j] + maxVal*j
			if sum > dp[i] {
				dp[i] = sum
			}
		}
	}

	return dp[n]
}
```

## 1045 — Customers Who Bought All Products

```go
package main

// LeetCode #1045: Customers Who Bought All Products
// https://leetcode.com/problems/customers-who-bought-all-products/
// Difficulty: Medium
//
// Approach: Count distinct products per customer, compare to total products.
// Time: O(n) where n = len(customer_product)
// Space: O(m) where m = number of customers

import "fmt"

func main() {
	// Simulated: customer_id, product_key
	customer := []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
	product := []int{1, 1, 1, 1, 1, 2, 2, 2, 2, 2}
	fmt.Println(customersWhoBoughtAllProducts(customer, product, 2)) // [1,2,3,4,5]

	customer2 := []int{1, 1, 2}
	product2 := []int{1, 2, 1}
	fmt.Println(customersWhoBoughtAllProducts(customer2, product2, 2)) // [1]
}

func customersWhoBoughtAllProducts(customer, product []int, totalProducts int) []int {
	bought := make(map[int]map[int]bool)
	customerSet := make(map[int]bool)

	for i := 0; i < len(customer); i++ {
		c, p := customer[i], product[i]
		customerSet[c] = true
		if bought[c] == nil {
			bought[c] = make(map[int]bool)
		}
		bought[c][p] = true
	}

	result := make([]int, 0)
	for c := range customerSet {
		if len(bought[c]) == totalProducts {
			result = append(result, c)
		}
	}

	return result
}
```

## 1048 — Longest String Chain

```go
package main

// LeetCode #1048: Longest String Chain
// https://leetcode.com/problems/longest-string-chain/
// Difficulty: Medium
//
// Approach: Sort by length, DP with hash map.
//           For each word, check all possible predecessors by removing one char.
// Time: O(n * L^2) where L is max word length
// Space: O(n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(longestStrChain([]string{"a", "b", "ba", "bca", "bda", "bdca"})) // 4
	fmt.Println(longestStrChain([]string{"xbc", "pcxbcf", "xb", "cxbc", "pcxbc"})) // 5
}

func longestStrChain(words []string) int {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	dp := make(map[string]int)
	result := 1

	for _, w := range words {
		best := 1
		for i := 0; i < len(w); i++ {
			pred := w[:i] + w[i+1:]
			if val, ok := dp[pred]; ok {
				if val+1 > best {
					best = val + 1
				}
			}
		}
		dp[w] = best
		if best > result {
			result = best
		}
	}

	return result
}
```

## 1049 — Last Stone Weight Ii

```go
package main

// LeetCode #1049: Last Stone Weight II
// https://leetcode.com/problems/last-stone-weight-ii/
// Difficulty: Medium
//
// Approach: DP - subset sum. Partition stones into two groups.
//           Minimize |sum - 2*subsetSum|.
// Time: O(n * sum)
// Space: O(sum)

import "fmt"

func main() {
	fmt.Println(lastStoneWeightII([]int{2, 7, 4, 1, 8, 1})) // 1
	fmt.Println(lastStoneWeightII([]int{31, 26, 33, 21, 40})) // 5
}

func lastStoneWeightII(stones []int) int {
	total := 0
	for _, s := range stones {
		total += s
	}

	target := total / 2
	dp := make([]bool, target+1)
	dp[0] = true

	for _, s := range stones {
		for j := target; j >= s; j-- {
			if dp[j-s] {
				dp[j] = true
			}
		}
	}

	for j := target; j >= 0; j-- {
		if dp[j] {
			return total - 2*j
		}
	}

	return total
}
```

## 1052 — Grumpy Bookstore Owner

```go
package main

// LeetCode #1052: Grumpy Bookstore Owner
// https://leetcode.com/problems/grumpy-bookstore-owner/
// Difficulty: Medium
//
// Approach: Sliding window - find best minutes to use secret technique
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(maxSatisfied([]int{1, 0, 1, 2, 1, 1, 7, 5}, []int{0, 1, 0, 1, 0, 1, 0, 1}, 3)) // 16
	fmt.Println(maxSatisfied([]int{1}, []int{0}, 1)) // 1
}

func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	n := len(customers)
	baseSatisfied := 0
	for i := 0; i < n; i++ {
		if grumpy[i] == 0 {
			baseSatisfied += customers[i]
		}
	}

	extraSatisfied := 0
	for i := 0; i < minutes; i++ {
		if grumpy[i] == 1 {
			extraSatisfied += customers[i]
		}
	}

	maxExtra := extraSatisfied
	for i := minutes; i < n; i++ {
		if grumpy[i-minutes] == 1 {
			extraSatisfied -= customers[i-minutes]
		}
		if grumpy[i] == 1 {
			extraSatisfied += customers[i]
		}
		if extraSatisfied > maxExtra {
			maxExtra = extraSatisfied
		}
	}

	return baseSatisfied + maxExtra
}
```

## 1053 — Previous Permutation With One Swap

```go
package main

// LeetCode #1053: Previous Permutation With One Swap
// https://leetcode.com/problems/previous-permutation-with-one-swap/
// Difficulty: Medium
//
// Approach: Find rightmost pair where arr[i] > arr[i+1].
//           Swap with the largest smaller element to the right.
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(prevPermOpt1([]int{3, 2, 1}))    // [3,1,2]
	fmt.Println(prevPermOpt1([]int{1, 1, 5}))    // [1,1,5]
	fmt.Println(prevPermOpt1([]int{1, 9, 4, 6, 7})) // [1,7,4,6,9]
}

func prevPermOpt1(arr []int) []int {
	i := len(arr) - 2
	for i >= 0 && arr[i] <= arr[i+1] {
		i--
	}

	if i < 0 {
		return arr
	}

	// Find rightmost smaller than arr[i]
	j := len(arr) - 1
	for arr[j] >= arr[i] {
		j--
	}
	// Skip duplicates
	for arr[j] == arr[j-1] {
		j--
	}

	arr[i], arr[j] = arr[j], arr[i]
	return arr
}
```

## 1054 — Distant Barcodes

```go
package main

// LeetCode #1054: Distant Barcodes
// https://leetcode.com/problems/distant-barcodes/
// Difficulty: Medium
//
// Approach: Count frequencies, place most frequent in even indices, then odd
// Time: O(n log n)
// Space: O(n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(rearrangeBarcodes([]int{1, 1, 1, 2, 2, 2})) // [1,2,1,2,1,2] or similar
	fmt.Println(rearrangeBarcodes([]int{1, 1, 1, 1, 2, 2, 3, 3})) // valid rearrangement
}

func rearrangeBarcodes(barcodes []int) []int {
	n := len(barcodes)
	freq := make(map[int]int)
	for _, b := range barcodes {
		freq[b]++
	}

	type pair struct {
		val   int
		count int
	}
	pairs := make([]pair, 0, len(freq))
	for val, count := range freq {
		pairs = append(pairs, pair{val, count})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].count > pairs[j].count
	})

	result := make([]int, n)
	idx := 0

	for _, p := range pairs {
		for k := 0; k < p.count; k++ {
			result[idx] = p.val
			idx += 2
			if idx >= n {
				idx = 1
			}
		}
	}

	return result
}
```

## 1055 — Shortest Way To Form String

```go
package main

// LeetCode #1055: Shortest Way to Form String
// https://leetcode.com/problems/shortest-way-to-form-string/
// Difficulty: Medium
//
// Approach: Greedy with two pointers. Count subsequence matches.
// Time: O(n * m) worst case, O(n + m) with precomputed indices
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(shortestWay("abc", "abcbc"))   // 2
	fmt.Println(shortestWay("abc", "acdbc"))   // -1
	fmt.Println(shortestWay("xyz", "xzyxz"))   // 2
}

func shortestWay(source string, target string) int {
	count := 0
	i := 0 // index in target

	// Pre-check: all chars in target must exist in source
	sourceSet := make(map[byte]bool)
	for k := 0; k < len(source); k++ {
		sourceSet[source[k]] = true
	}
	for k := 0; k < len(target); k++ {
		if !sourceSet[target[k]] {
			return -1
		}
	}

	for i < len(target) {
		count++
		j := 0 // index in source
		for j < len(source) && i < len(target) {
			if source[j] == target[i] {
				i++
			}
			j++
		}
	}

	return count
}
```

## 1057 — Campus Bikes

```go
package main

// LeetCode #1057: Campus Bikes
// https://leetcode.com/problems/campus-bikes/
// Difficulty: Medium
//
// Approach: Bucket sort by Manhattan distance. Assign closest pairs.
// Time: O(W * B) where W = workers, B = bikes
// Space: O(W * B)

import "fmt"

func main() {
	fmt.Println(assignBikes([][]int{{0, 0}, {2, 1}}, [][]int{{1, 2}, {3, 3}})) // [1,0]
	fmt.Println(assignBikes([][]int{{0, 0}, {1, 1}, {2, 0}}, [][]int{{1, 0}, {2, 2}, {2, 1}})) // [0,2,1]
}

func assignBikes(workers [][]int, bikes [][]int) []int {
	w, b := len(workers), len(bikes)
	// Buckets of distances: max distance is 2000 (0-1000, 0-1000)
	maxDist := 2000
	buckets := make([][][2]int, maxDist+1)

	for i := 0; i < w; i++ {
		for j := 0; j < b; j++ {
			dist := abs(workers[i][0]-bikes[j][0]) + abs(workers[i][1]-bikes[j][1])
			buckets[dist] = append(buckets[dist], [2]int{i, j})
		}
	}

	result := make([]int, w)
	for i := range result {
		result[i] = -1
	}
	bikeUsed := make([]bool, b)

	for d := 0; d <= maxDist; d++ {
		for _, pair := range buckets[d] {
			wi, bi := pair[0], pair[1]
			if result[wi] == -1 && !bikeUsed[bi] {
				result[wi] = bi
				bikeUsed[bi] = true
			}
		}
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

## 1058 — Minimize Rounding Error To Meet Target

```go
package main

// LeetCode #1058: Minimize Rounding Error to Meet Target
// https://leetcode.com/problems/minimize-rounding-error-to-meet-target/
// Difficulty: Medium
//
// Approach: For each price, floor and ceil. Compute min total error via DP.
// Time: O(n * target) effectively O(n) after sorting diffs
// Space: O(n)

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(minimizeRoundingErrorToMeetTarget([]string{"0.700", "2.800", "4.900"}, 8)) // "1.000"
	fmt.Println(minimizeRoundingErrorToMeetTarget([]string{"1.500", "2.500", "3.500"}, 10)) // "-1"
}

func minimizeRoundingErrorToMeetTarget(prices []string, target int) string {
	n := len(prices)
	floors := make([]int, n)
	diffs := make([]float64, n)
	floorSum := 0

	for i, p := range prices {
		f, _ := strconv.ParseFloat(p, 64)
		floor := int(math.Floor(f))
		floors[i] = floor
		floorSum += floor
		diffs[i] = f - float64(floor)
	}

	// We need to ceil (target - floorSum) items
	ceilCount := target - floorSum
	if ceilCount < 0 || ceilCount > n {
		return "-1"
	}

	// Sort diffs descending to ceil those with smallest rounding error
	sort.Slice(diffs, func(i, j int) bool {
		return diffs[i] > diffs[j]
	})

	totalErr := 0.0
	for i := 0; i < ceilCount; i++ {
		totalErr += 1.0 - diffs[i]
	}
	for i := ceilCount; i < n; i++ {
		totalErr += diffs[i]
	}

	return fmt.Sprintf("%.3f", math.Round(totalErr*1000)/1000)
}
```

## 1059 — All Paths From Source Lead To Destination

```go
package main

// LeetCode #1059: All Paths from Source Lead to Destination
// https://leetcode.com/problems/all-paths-from-source-lead-to-destination/
// Difficulty: Medium
//
// Approach: DFS with cycle detection. Every path from source must end at destination.
// Time: O(V + E)
// Space: O(V + E)

import "fmt"

func main() {
	fmt.Println(leadsToDestination(3, [][]int{{0, 1}, {0, 2}}, 0, 2)) // false
	fmt.Println(leadsToDestination(4, [][]int{{0, 1}, {0, 2}, {1, 3}, {2, 3}}, 0, 3)) // true
}

func leadsToDestination(n int, edges [][]int, source int, destination int) bool {
	graph := make([][]int, n)
	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
	}

	state := make([]int, n) // 0=unvisited, 1=visiting, 2=processed

	var dfs func(node int) bool
	dfs = func(node int) bool {
		if state[node] == 1 {
			return false // cycle
		}
		if state[node] == 2 {
			return true
		}

		if len(graph[node]) == 0 {
			return node == destination
		}

		state[node] = 1
		for _, next := range graph[node] {
			if !dfs(next) {
				return false
			}
		}
		state[node] = 2
		return true
	}

	return dfs(source)
}
```

## 1060 — Missing Element In Sorted Array

```go
package main

// LeetCode #1060: Missing Element in Sorted Array
// https://leetcode.com/problems/missing-element-in-sorted-array/
// Difficulty: Medium
//
// Approach: Binary search. Number of missing elements up to index i = nums[i] - nums[0] - i.
// Time: O(log n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(missingElement([]int{4, 7, 9, 10}, 1))  // 5
	fmt.Println(missingElement([]int{4, 7, 9, 10}, 3))  // 8
	fmt.Println(missingElement([]int{1, 2, 4}, 3))      // 6
}

func missingElement(nums []int, k int) int {
	n := len(nums)

	missingCount := func(idx int) int {
		return nums[idx] - nums[0] - idx
	}

	if missingCount(n-1) < k {
		return nums[n-1] + (k - missingCount(n-1))
	}

	left, right := 0, n-1
	for left < right {
		mid := left + (right-left)/2
		if missingCount(mid) < k {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return nums[left-1] + (k - missingCount(left-1))
}
```

## 1061 — Lexicographically Smallest Equivalent String

```go
package main

// LeetCode #1061: Lexicographically Smallest Equivalent String
// https://leetcode.com/problems/lexicographically-smallest-equivalent-string/
// Difficulty: Medium
//
// Approach: DSU (Union-Find) to group equivalent characters,
//           then replace each char in baseStr with its smallest root.
// Time: O((m + n) * alpha(26)) where m = len(s1), n = len(baseStr)
// Space: O(26) = O(1)

import "fmt"

func main() {
	fmt.Println(smallestEquivalentString("parker", "morris", "parser")) // "makkek"
	fmt.Println(smallestEquivalentString("hello", "world", "hold"))     // "hdld"
}

func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
	parent := make([]int, 26)
	for i := 0; i < 26; i++ {
		parent[i] = i
	}

	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func(a, b int) {
		pa, pb := find(a), find(b)
		if pa < pb {
			parent[pb] = pa
		} else {
			parent[pa] = pb
		}
	}

	for i := 0; i < len(s1); i++ {
		union(int(s1[i]-'a'), int(s2[i]-'a'))
	}

	result := make([]byte, len(baseStr))
	for i := 0; i < len(baseStr); i++ {
		result[i] = byte('a' + find(int(baseStr[i]-'a')))
	}

	return string(result)
}
```

## 1062 — Longest Repeating Substring

```go
package main

// LeetCode #1062: Longest Repeating Substring
// https://leetcode.com/problems/longest-repeating-substring/
// Difficulty: Medium
//
// Approach: DP - find longest common prefix between all pairs of suffixes
// Time: O(n^2)
// Space: O(n^2) can be O(n) with optimized DP

import "fmt"

func main() {
	fmt.Println(longestRepeatingSubstring("abcd"))    // 0
	fmt.Println(longestRepeatingSubstring("abbaba"))  // 2
	fmt.Println(longestRepeatingSubstring("aabcaabdaab")) // 3
}

func longestRepeatingSubstring(s string) int {
	n := len(s)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	result := 0
	for i := 1; i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			if s[i-1] == s[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > result {
					result = dp[i][j]
				}
			}
		}
	}

	return result
}
```

## 1066 — Campus Bikes Ii

```go
package main

// LeetCode #1066: Campus Bikes II
// https://leetcode.com/problems/campus-bikes-ii/
// Difficulty: Medium
//
// Approach: DP with bitmask (minimum assignment cost)
// Time: O(W * 2^B) where W = workers, B = bikes
// Space: O(2^B)

import "fmt"

func main() {
	fmt.Println(assignBikesII([][]int{{0, 0}, {2, 1}}, [][]int{{1, 2}, {3, 3}})) // 6
	fmt.Println(assignBikesII([][]int{{0, 0}, {1, 1}, {2, 0}}, [][]int{{1, 0}, {2, 2}, {2, 1}})) // 4
}

func assignBikesII(workers [][]int, bikes [][]int) int {
	w, b := len(workers), len(bikes)
	dp := make([]int, 1<<b)
	for i := range dp {
		dp[i] = -1
	}

	var dfs func(workerIdx int, mask int) int
	dfs = func(workerIdx int, mask int) int {
		if workerIdx == w {
			return 0
		}
		if dp[mask] != -1 {
			return dp[mask]
		}

		minDist := 1<<31 - 1
		for j := 0; j < b; j++ {
			if mask&(1<<j) == 0 {
				dist := abs(workers[workerIdx][0]-bikes[j][0]) + abs(workers[workerIdx][1]-bikes[j][1])
				total := dist + dfs(workerIdx+1, mask|(1<<j))
				if total < minDist {
					minDist = total
				}
			}
		}
		dp[mask] = minDist
		return minDist
	}

	return dfs(0, 0)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

## 1070 — Product Sales Analysis Iii

```go
package main

// LeetCode #1070: Product Sales Analysis III
// https://leetcode.com/problems/product-sales-analysis-iii/
// Difficulty: Medium
//
// Approach: Find first year of each product, get its sale data.
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	// Input: sales = [(product_id, year, quantity, price), ...]
	sales := [][]int{{1, 100, 2008, 10, 5000}, {2, 100, 2009, 12, 5000}, {3, 100, 2008, 15, 5000}}
	fmt.Println(productSalesAnalysisIII(sales)) // [[100,2008,10,5000]]
}

func productSalesAnalysisIII(sales [][]int) [][]int {
	// sales[i] = [sale_id, product_id, year, quantity, price]
	firstYear := make(map[int]int) // product_id -> min year
	saleByProd := make(map[int][]int)

	for _, s := range sales {
		pid, year := s[1], s[2]
		if _, ok := firstYear[pid]; !ok || year < firstYear[pid] {
			firstYear[pid] = year
		}
		saleByProd[pid] = s
	}

	result := make([][]int, 0)
	seen := make(map[int]bool)
	for _, s := range sales {
		pid, year := s[1], s[2]
		if !seen[pid] && year == firstYear[pid] {
			// Return [product_id, year, quantity, price]
			result = append(result, []int{pid, year, s[3], s[4]})
			seen[pid] = true
		}
	}

	return result
}
```

## 1072 — Flip Columns For Maximum Number Of Equal Rows

```go
package main

// LeetCode #1072: Flip Columns For Maximum Number of Equal Rows
// https://leetcode.com/problems/flip-columns-for-maximum-number-of-equal-rows/
// Difficulty: Medium
//
// Approach: Normalize each row to a pattern (starting with 0).
//           Rows with same or complementary pattern can be made equal.
// Time: O(m * n)
// Space: O(m * n)

import "fmt"

func main() {
	fmt.Println(maxEqualRowsAfterFlips([][]int{{0, 1}, {1, 1}})) // 1
	fmt.Println(maxEqualRowsAfterFlips([][]int{{0, 0, 0}, {0, 0, 1}, {1, 1, 0}})) // 2
}

func maxEqualRowsAfterFlips(matrix [][]int) int {
	patternCount := make(map[string]int)

	for _, row := range matrix {
		pattern := make([]byte, len(row))
		for j := 0; j < len(row); j++ {
			if row[0] == 0 {
				pattern[j] = byte('0' + row[j])
			} else {
				pattern[j] = byte('0' + 1 - row[j])
			}
		}
		patternCount[string(pattern)]++
	}

	result := 0
	for _, count := range patternCount {
		if count > result {
			result = count
		}
	}
	return result
}
```

## 1073 — Adding Two Negabinary Numbers

```go
package main

// LeetCode #1073: Adding Two Negabinary Numbers
// https://leetcode.com/problems/adding-two-negabinary-numbers/
// Difficulty: Medium
//
// Approach: Sum digits from right to left with carry in base -2
// Time: O(max(m, n))
// Space: O(max(m, n))

import "fmt"

func main() {
	fmt.Println(addNegabinary([]int{1, 1, 1, 1, 1}, []int{1, 0, 1})) // [1,0,0,0,0]
	fmt.Println(addNegabinary([]int{0}, []int{0}))                   // [0]
}

func addNegabinary(arr1 []int, arr2 []int) []int {
	i, j := len(arr1)-1, len(arr2)-1
	carry := 0
	result := make([]int, 0)

	for i >= 0 || j >= 0 || carry != 0 {
		if i >= 0 {
			carry += arr1[i]
			i--
		}
		if j >= 0 {
			carry += arr2[j]
			j--
		}

		result = append(result, carry&1)
		carry = -(carry >> 1)
	}

	// Reverse and remove leading zeros
	for left, right := 0, len(result)-1; left < right; left, right = left+1, right-1 {
		result[left], result[right] = result[right], result[left]
	}

	// Remove leading zeros
	start := 0
	for start < len(result)-1 && result[start] == 0 {
		start++
	}

	return result[start:]
}
```

## 1077 — Project Employees Iii

```go
package main

// LeetCode #1077: Project Employees III
// https://leetcode.com/problems/project-employees-iii/
// Difficulty: Medium
//
// Approach: Group employee experience by project, find max per project
// Time: O(n log n) where n = len(project)
// Space: O(n)

import "fmt"

func main() {
	// (project_id, employee_id, experience_years)
	project := [][]int{{1, 1}, {1, 2}, {2, 3}, {2, 4}}
	employee := [][]int{{1, 5}, {2, 3}, {3, 7}, {4, 2}}
	fmt.Println(projectEmployeesIII(project, employee))
}

func projectEmployeesIII(project [][]int, employee [][]int) [][]int {
	expMap := make(map[int]int)
	for _, e := range employee {
		expMap[e[0]] = e[1]
	}

	// For each project, find max experience and which employees have it
	type projInfo struct {
		maxExp int
		empID  int
	}
	projMax := make(map[int]projInfo)

	for _, p := range project {
		projID, empID := p[0], p[1]
		exp := expMap[empID]

		if info, ok := projMax[projID]; !ok || exp > info.maxExp {
			projMax[projID] = projInfo{exp, empID}
		}
	}

	result := make([][]int, 0)
	for _, p := range project {
		projID, empID := p[0], p[1]
		info := projMax[projID]
		if empID == info.empID {
			result = append(result, []int{projID, empID})
		}
	}

	return result
}
```

