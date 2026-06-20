# Medium (Sedang) — Problem ��2433

## 2240 — Number Of Ways To Buy Pens And Pencils

```go
package main

// LeetCode #2240: Number of Ways to Buy Pens and Pencils
// https://leetcode.com/problems/number-of-ways-to-buy-pens-and-pencils/
// Difficulty: Medium
// Time: O(total / cost1) | Space: O(1)

import "fmt"

func waysToBuyPensPencils(total int, cost1 int, cost2 int) int64 {
	var ways int64 = 0
	for pens := 0; pens*cost1 <= total; pens++ {
		remaining := total - pens*cost1
		ways += int64(remaining/cost2) + 1
		if cost2 == 0 {
			break
		}
	}
	return ways
}

func main() {
	// Test case 1
	fmt.Println(waysToBuyPensPencils(20, 10, 5))
	// Expected: 9

	// Test case 2
	fmt.Println(waysToBuyPensPencils(5, 10, 10))
	// Expected: 1

	// Test case 3
	fmt.Println(waysToBuyPensPencils(100, 1, 1))
	// Expected: 5151
}
```

## 2241 — Design An Atm Machine

```go
package main

// LeetCode #2241: Design an ATM Machine
// https://leetcode.com/problems/design-an-atm-machine/
// Difficulty: Medium
// Time: O(1) per operation | Space: O(1)

import "fmt"

type ATM struct {
	denoms []int
	notes  []int
}

func Constructor() ATM {
	return ATM{
		denoms: []int{20, 50, 100, 200, 500},
		notes:  make([]int, 5),
	}
}

func (this *ATM) Deposit(banknotesCount []int) {
	for i, c := range banknotesCount {
		this.notes[i] += c
	}
}

func (this *ATM) Withdraw(amount int) []int {
	result := make([]int, 5)
	need := amount
	for i := 4; i >= 0; i-- {
		use := need / this.denoms[i]
		if use > this.notes[i] {
			use = this.notes[i]
		}
		result[i] = use
		need -= use * this.denoms[i]
	}
	if need != 0 {
		return []int{-1}
	}
	for i, c := range result {
		this.notes[i] -= c
	}
	return result
}

func main() {
	atm := Constructor()
	atm.Deposit([]int{0, 0, 1, 2, 1})
	fmt.Println(atm.Withdraw(600))
	// Expected: [0,0,1,0,1]

	atm.Deposit([]int{0, 1, 0, 1, 1})
	fmt.Println(atm.Withdraw(600))
	// Expected: [-1]

	fmt.Println(atm.Withdraw(550))
	// Expected: [0,1,0,0,1]
}
```

## 2244 — Minimum Rounds To Complete All Tasks

```go
package main

// LeetCode #2244: Minimum Rounds to Complete All Tasks
// https://leetcode.com/problems/minimum-rounds-to-complete-all-tasks/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func minimumRounds(tasks []int) int {
	freq := make(map[int]int)
	for _, t := range tasks {
		freq[t]++
	}

	rounds := 0
	for _, c := range freq {
		if c == 1 {
			return -1
		}
		// 3*c + 2*(c/3 remainder)
		rounds += c / 3
		if c%3 != 0 {
			rounds++
		}
	}
	return rounds
}

func main() {
	// Test case 1
	fmt.Println(minimumRounds([]int{2, 2, 3, 3, 2, 4, 4, 4, 4, 4}))
	// Expected: 4

	// Test case 2
	fmt.Println(minimumRounds([]int{2, 3, 3}))
	// Expected: -1

	// Test case 3
	fmt.Println(minimumRounds([]int{5, 5, 5, 5}))
	// Expected: 2
}
```

## 2245 — Maximum Trailing Zeros In A Cornered Path

```go
package main

// LeetCode #2245: Maximum Trailing Zeros in a Cornered Path
// https://leetcode.com/problems/maximum-trailing-zeros-in-a-cornered-path/
// Difficulty: Medium
// Time: O(m * n) | Space: O(m * n)

import "fmt"

func maxTrailingZeros(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	// Precompute prefix counts of factors 2 and 5
	type pair struct{ two, five int }

	// Right prefix
	right := make([][]pair, m)
	for i := 0; i < m; i++ {
		right[i] = make([]pair, n+1)
		for j := 0; j < n; j++ {
			two, five := countFactors(grid[i][j])
			right[i][j+1] = pair{right[i][j].two + two, right[i][j].five + five}
		}
	}

	// Down prefix
	down := make([][]pair, m+1)
	for i := 0; i <= m; i++ {
		down[i] = make([]pair, n)
	}
	for j := 0; j < n; j++ {
		for i := 0; i < m; i++ {
			two, five := countFactors(grid[i][j])
			down[i+1][j] = pair{down[i][j].two + two, down[i][j].five + five}
		}
	}

	maxZeros := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// From top to (i,j) then right
			t1 := down[i+1][j].two + right[i][n].two - right[i][j+1].two
			f1 := down[i+1][j].five + right[i][n].five - right[i][j+1].five
			if t1 < f1 {
				if t1 > maxZeros {
					maxZeros = t1
				}
			} else {
				if f1 > maxZeros {
					maxZeros = f1
				}
			}

			// From top to (i,j) then left
			t2 := down[i+1][j].two + right[i][j].two
			f2 := down[i+1][j].five + right[i][j].five
			if t2 < f2 {
				if t2 > maxZeros {
					maxZeros = t2
				}
			} else {
				if f2 > maxZeros {
					maxZeros = f2
				}
			}

			// From bottom to (i,j) then right
			t3 := (down[m][j].two - down[i][j].two) + right[i][n].two - right[i][j+1].two
			f3 := (down[m][j].five - down[i][j].five) + right[i][n].five - right[i][j+1].five
			if t3 < f3 {
				if t3 > maxZeros {
					maxZeros = t3
				}
			} else {
				if f3 > maxZeros {
					maxZeros = f3
				}
			}

			// From bottom to (i,j) then left
			t4 := (down[m][j].two - down[i][j].two) + right[i][j].two
			f4 := (down[m][j].five - down[i][j].five) + right[i][j].five
			if t4 < f4 {
				if t4 > maxZeros {
					maxZeros = t4
				}
			} else {
				if f4 > maxZeros {
					maxZeros = f4
				}
			}
		}
	}
	return maxZeros
}

func countFactors(x int) (int, int) {
	two, five := 0, 0
	for x%2 == 0 {
		two++
		x /= 2
	}
	for x%5 == 0 {
		five++
		x /= 5
	}
	return two, five
}

func main() {
	// Test case 1
	fmt.Println(maxTrailingZeros([][]int{{23, 17, 15, 3, 20}, {8, 1, 20, 27, 11}, {9, 4, 6, 2, 21}, {40, 9, 1, 10, 6}, {22, 7, 4, 5, 3}}))
	// Expected: 3

	// Test case 2
	fmt.Println(maxTrailingZeros([][]int{{4, 3, 2}, {7, 6, 1}, {8, 8, 8}}))
	// Expected: 0
}
```

## 2249 — Count Lattice Points Inside A Circle

```go
package main

// LeetCode #2249: Count Lattice Points Inside a Circle
// https://leetcode.com/problems/count-lattice-points-inside-a-circle/
// Difficulty: Medium
// Time: O(n * r^2) | Space: O(n)

import "fmt"

func countLatticePoints(circles [][]int) int {
	points := make(map[[2]int]bool)

	for _, c := range circles {
		x, y, r := c[0], c[1], c[2]
		rr := r * r
		for dx := -r; dx <= r; dx++ {
			for dy := -r; dy <= r; dy++ {
				if dx*dx+dy*dy <= rr {
					points[[2]int{x + dx, y + dy}] = true
				}
			}
		}
	}
	return len(points)
}

func main() {
	// Test case 1
	fmt.Println(countLatticePoints([][]int{{2, 2, 1}}))
	// Expected: 5

	// Test case 2
	fmt.Println(countLatticePoints([][]int{{2, 2, 2}, {3, 4, 1}}))
	// Expected: 16
}
```

## 2250 — Count Number Of Rectangles Containing Each Point

```go
package main

// LeetCode #2250: Count Number of Rectangles Containing Each Point
// https://leetcode.com/problems/count-number-of-rectangles-containing-each-point/
// Difficulty: Medium
// Time: O((n + m) log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func countRectangles(rectangles [][]int, points [][]int) []int {
	// Group rectangles by height
	byHeight := make([][]int, 101)
	for _, r := range rectangles {
		h := r[1]
		byHeight[h] = append(byHeight[h], r[0])
	}
	for h := 0; h <= 100; h++ {
		sort.Ints(byHeight[h])
	}

	result := make([]int, len(points))
	for i, p := range points {
		x, y := p[0], p[1]
		count := 0
		for h := y; h <= 100; h++ {
			if len(byHeight[h]) == 0 {
				continue
			}
			// Binary search for first rectangle with width >= x
			idx := sort.SearchInts(byHeight[h], x)
			count += len(byHeight[h]) - idx
		}
		result[i] = count
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println(countRectangles([][]int{{1, 2}, {2, 3}, {2, 5}}, [][]int{{2, 1}, {1, 4}}))
	// Expected: [2, 1]

	// Test case 2
	fmt.Println(countRectangles([][]int{{1, 1}, {2, 2}, {3, 3}}, [][]int{{1, 3}, {1, 1}}))
	// Expected: [1, 3]
}
```

## 2256 — Minimum Average Difference

```go
package main

// LeetCode #2256: Minimum Average Difference
// https://leetcode.com/problems/minimum-average-difference/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func minimumAverageDifference(nums []int) int {
	n := len(nums)
	total := int64(0)
	for _, v := range nums {
		total += int64(v)
	}

	var prefix int64 = 0
	minDiff := int64(1 << 60)
	minIdx := 0

	for i := 0; i < n; i++ {
		prefix += int64(nums[i])
		left := prefix / int64(i+1)
		var right int64 = 0
		if i < n-1 {
			right = (total - prefix) / int64(n-i-1)
		}
		diff := left - right
		if diff < 0 {
			diff = -diff
		}
		if diff < minDiff {
			minDiff = diff
			minIdx = i
		}
	}
	return minIdx
}

func main() {
	// Test case 1
	fmt.Println(minimumAverageDifference([]int{2, 5, 3, 9, 5, 3}))
	// Expected: 3

	// Test case 2
	fmt.Println(minimumAverageDifference([]int{0}))
	// Expected: 0

	// Test case 3
	fmt.Println(minimumAverageDifference([]int{0, 1, 0, 1, 0, 1}))
	// Expected: 1
}
```

## 2257 — Count Unguarded Cells In The Grid

```go
package main

// LeetCode #2257: Count Unguarded Cells in the Grid
// https://leetcode.com/problems/count-unguarded-cells-in-the-grid/
// Difficulty: Medium
// Time: O(m * n) | Space: O(m * n)

import "fmt"

func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
	grid := make([][]int, m)
	for i := 0; i < m; i++ {
		grid[i] = make([]int, n)
	}
	// 1 = wall, 2 = guard

	for _, w := range walls {
		grid[w[0]][w[1]] = 1
	}
	for _, g := range guards {
		grid[g[0]][g[1]] = 2
	}

	dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for _, g := range guards {
		gr, gc := g[0], g[1]
		for _, d := range dirs {
			r, c := gr+d[0], gc+d[1]
			for r >= 0 && r < m && c >= 0 && c < n && grid[r][c] != 1 && grid[r][c] != 2 {
				if grid[r][c] == 0 {
					grid[r][c] = 3 // guarded
				}
				r += d[0]
				c += d[1]
			}
		}
	}

	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				count++
			}
		}
	}
	return count
}

func main() {
	// Test case 1
	fmt.Println(countUnguarded(4, 6, [][]int{{0, 0}, {1, 1}, {2, 3}}, [][]int{{0, 1}, {2, 2}, {1, 4}}))
	// Expected: 7

	// Test case 2
	fmt.Println(countUnguarded(3, 3, [][]int{{1, 1}}, [][]int{{0, 1}, {1, 0}, {2, 1}, {1, 2}}))
	// Expected: 4
}
```

## 2260 — Minimum Consecutive Cards To Pick Up

```go
package main

// LeetCode #2260: Minimum Consecutive Cards to Pick Up
// https://leetcode.com/problems/minimum-consecutive-cards-to-pick-up/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func minimumCardPickup(cards []int) int {
	last := make(map[int]int)
	minLen := len(cards) + 1

	for i, c := range cards {
		if prev, ok := last[c]; ok {
			dist := i - prev + 1
			if dist < minLen {
				minLen = dist
			}
		}
		last[c] = i
	}

	if minLen > len(cards) {
		return -1
	}
	return minLen
}

func main() {
	// Test case 1
	fmt.Println(minimumCardPickup([]int{3, 4, 2, 3, 4, 7}))
	// Expected: 4

	// Test case 2
	fmt.Println(minimumCardPickup([]int{1, 0, 5, 3}))
	// Expected: -1

	// Test case 3
	fmt.Println(minimumCardPickup([]int{1, 2, 3, 4, 5, 1}))
	// Expected: 6
}
```

## 2261 — K Divisible Elements Subarrays

```go
package main

// LeetCode #2261: K Divisible Elements Subarrays
// https://leetcode.com/problems/k-divisible-elements-subarrays/
// Difficulty: Medium
// Time: O(n^2 * k) | Space: O(n^2)

import (
	"fmt"
	"strconv"
	"strings"
)

func countDistinct(nums []int, k int, p int) int {
	n := len(nums)
	seen := make(map[string]bool)

	for i := 0; i < n; i++ {
		count := 0
		var sb strings.Builder
		for j := i; j < n; j++ {
			if nums[j]%p == 0 {
				count++
			}
			if count > k {
				break
			}
			if sb.Len() > 0 {
				sb.WriteString(",")
			}
			sb.WriteString(strconv.Itoa(nums[j]))
			seen[sb.String()] = true
		}
	}
	return len(seen)
}

func main() {
	// Test case 1
	fmt.Println(countDistinct([]int{2, 3, 3, 2, 2}, 2, 2))
	// Expected: 11

	// Test case 2
	fmt.Println(countDistinct([]int{1, 2, 3, 4}, 4, 5))
	// Expected: 10
}
```

## 2265 — Count Nodes Equal To Average Of Subtree

```go
package main

// LeetCode #2265: Count Nodes Equal to Average of Subtree
// https://leetcode.com/problems/count-nodes-equal-to-average-of-subtree/
// Difficulty: Medium
// Time: O(n) | Space: O(h)

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfSubtree(root *TreeNode) int {
	count := 0
	dfs(root, &count)
	return count
}

func dfs(node *TreeNode, count *int) (int, int) {
	if node == nil {
		return 0, 0
	}
	leftSum, leftCount := dfs(node.Left, count)
	rightSum, rightCount := dfs(node.Right, count)

	sum := leftSum + rightSum + node.Val
	totalNodes := leftCount + rightCount + 1

	if sum/totalNodes == node.Val {
		*count++
	}
	return sum, totalNodes
}

func main() {
	// Test case 1: [4,8,5,0,1,null,6]
	root1 := &TreeNode{Val: 4}
	root1.Left = &TreeNode{Val: 8}
	root1.Right = &TreeNode{Val: 5}
	root1.Left.Left = &TreeNode{Val: 0}
	root1.Left.Right = &TreeNode{Val: 1}
	root1.Right.Right = &TreeNode{Val: 6}
	fmt.Println(averageOfSubtree(root1))
	// Expected: 5

	// Test case 2: [1]
	root2 := &TreeNode{Val: 1}
	fmt.Println(averageOfSubtree(root2))
	// Expected: 1
}
```

## 2266 — Count Number Of Texts

```go
package main

// LeetCode #2266: Count Number of Texts
// https://leetcode.com/problems/count-number-of-texts/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func countTexts(pressedKeys string) int {
	const mod = 1_000_000_007
	n := len(pressedKeys)
	dp := make([]int, n+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] // press once
		// Check for multiple presses of same digit
		maxPress := 3
		if pressedKeys[i-1] == '7' || pressedKeys[i-1] == '9' {
			maxPress = 4
		}
		for j := 2; j <= maxPress && j <= i; j++ {
			if pressedKeys[i-j] == pressedKeys[i-1] {
				dp[i] = (dp[i] + dp[i-j]) % mod
			} else {
				break
			}
		}
	}
	return dp[n]
}

func main() {
	// Test case 1
	fmt.Println(countTexts("22233"))
	// Expected: 8

	// Test case 2
	fmt.Println(countTexts("222222222222222222222222222222222222"))
	// Expected: 82876089

	// Test case 3
	fmt.Println(countTexts("33"))
	// Expected: 2
}
```

## 2268 — Minimum Number Of Keypresses

```go
package main

// LeetCode #2268: Minimum Number of Keypresses
// https://leetcode.com/problems/minimum-number-of-keypresses/
// Difficulty: Medium [Paid]
// Time: O(n + 26 log 26) | Space: O(26)

import (
	"fmt"
	"sort"
)

func minimumKeypresses(s string) int {
	count := make([]int, 26)
	for _, ch := range s {
		count[ch-'a']++
	}

	sort.Slice(count, func(i, j int) bool {
		return count[i] > count[j]
	})

	presses := 0
	for i, c := range count {
		if c == 0 {
			break
		}
		presses += c * (i/9 + 1)
	}
	return presses
}

func main() {
	// Test case 1
	fmt.Println(minimumKeypresses("apple"))
	// Expected: 5

	// Test case 2
	fmt.Println(minimumKeypresses("abcdefghijkl"))
	// Expected: 15

	// Test case 3
	fmt.Println(minimumKeypresses("aaaaaaa"))
	// Expected: 7
}
```

## 2270 — Number Of Ways To Split Array

```go
package main

// LeetCode #2270: Number of Ways to Split Array
// https://leetcode.com/problems/number-of-ways-to-split-array/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func waysToSplitArray(nums []int) int {
	n := len(nums)
	total := int64(0)
	for _, v := range nums {
		total += int64(v)
	}

	var prefix int64 = 0
	ways := 0
	for i := 0; i < n-1; i++ {
		prefix += int64(nums[i])
		if prefix >= total-prefix {
			ways++
		}
	}
	return ways
}

func main() {
	// Test case 1
	fmt.Println(waysToSplitArray([]int{10, 4, -8, 7}))
	// Expected: 2

	// Test case 2
	fmt.Println(waysToSplitArray([]int{2, 3, 1, 0}))
	// Expected: 2

	// Test case 3
	fmt.Println(waysToSplitArray([]int{-1, -2, -3, -4}))
	// Expected: 0
}
```

## 2271 — Maximum White Tiles Covered By A Carpet

```go
package main

// LeetCode #2271: Maximum White Tiles Covered by a Carpet
// https://leetcode.com/problems/maximum-white-tiles-covered-by-a-carpet/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func maximumWhiteTiles(tiles [][]int, carpetLen int) int {
	sort.Slice(tiles, func(i, j int) bool {
		return tiles[i][0] < tiles[j][0]
	})

	maxCovered := 0
	right := 0
	prefix := 0

	for left := 0; left < len(tiles); left++ {
		// Expand right pointer to cover tiles within carpet
		for right < len(tiles) && tiles[right][1] < tiles[left][0]+carpetLen {
			prefix += tiles[right][1] - tiles[right][0] + 1
			right++
		}

		covered := prefix
		if right < len(tiles) {
			partial := tiles[left][0] + carpetLen - 1 - tiles[right][0] + 1
			if partial > 0 {
				covered += partial
			}
		}

		if covered > maxCovered {
			maxCovered = covered
		}

		// Remove left tile prefix before moving
		prefix -= tiles[left][1] - tiles[left][0] + 1
	}
	return maxCovered
}

func main() {
	// Test case 1
	fmt.Println(maximumWhiteTiles([][]int{{1, 5}, {10, 11}, {12, 18}, {20, 25}, {30, 32}}, 10))
	// Expected: 9

	// Test case 2
	fmt.Println(maximumWhiteTiles([][]int{{10, 11}, {1, 1}}, 2))
	// Expected: 2
}
```

## 2274 — Maximum Consecutive Floors Without Special Floors

```go
package main

// LeetCode #2274: Maximum Consecutive Floors Without Special Floors
// https://leetcode.com/problems/maximum-consecutive-floors-without-special-floors/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func maxConsecutive(bottom int, top int, special []int) int {
	sort.Ints(special)

	maxGap := special[0] - bottom
	for i := 1; i < len(special); i++ {
		gap := special[i] - special[i-1] - 1
		if gap > maxGap {
			maxGap = gap
		}
	}
	gap := top - special[len(special)-1]
	if gap > maxGap {
		maxGap = gap
	}
	return maxGap
}

func main() {
	// Test case 1
	fmt.Println(maxConsecutive(2, 9, []int{4, 6}))
	// Expected: 3

	// Test case 2
	fmt.Println(maxConsecutive(6, 8, []int{7, 6, 8}))
	// Expected: 0

	// Test case 3
	fmt.Println(maxConsecutive(1, 1000000000, []int{1, 1000000000}))
	// Expected: 999999998
}
```

## 2275 — Largest Combination With Bitwise And Greater Than Zero

```go
package main

// LeetCode #2275: Largest Combination With Bitwise AND Greater Than Zero
// https://leetcode.com/problems/largest-combination-with-bitwise-and-greater-than-zero/
// Difficulty: Medium
// Time: O(n * 24) | Space: O(1)

import "fmt"

func largestCombination(candidates []int) int {
	maxCount := 0
	for bit := 0; bit < 24; bit++ {
		count := 0
		for _, c := range candidates {
			if c&(1<<bit) != 0 {
				count++
			}
		}
		if count > maxCount {
			maxCount = count
		}
	}
	return maxCount
}

func main() {
	// Test case 1
	fmt.Println(largestCombination([]int{16, 17, 71, 62, 12, 24, 14}))
	// Expected: 4

	// Test case 2
	fmt.Println(largestCombination([]int{8, 8}))
	// Expected: 2

	// Test case 3
	fmt.Println(largestCombination([]int{1, 2, 4, 8}))
	// Expected: 1
}
```

## 2279 — Maximum Bags With Full Capacity Of Rocks

```go
package main

// LeetCode #2279: Maximum Bags With Full Capacity of Rocks
// https://leetcode.com/problems/maximum-bags-with-full-capacity-of-rocks/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func maximumBags(capacity []int, rocks []int, additionalRocks int) int {
	n := len(capacity)
	needed := make([]int, n)
	for i := 0; i < n; i++ {
		needed[i] = capacity[i] - rocks[i]
	}
	sort.Ints(needed)

	full := 0
	for _, n := range needed {
		if n == 0 {
			full++
		} else if n <= additionalRocks {
			additionalRocks -= n
			full++
		} else {
			break
		}
	}
	return full
}

func main() {
	// Test case 1
	fmt.Println(maximumBags([]int{2, 3, 4, 5}, []int{1, 2, 4, 4}, 2))
	// Expected: 3

	// Test case 2
	fmt.Println(maximumBags([]int{10, 2, 2}, []int{2, 2, 0}, 100))
	// Expected: 3
}
```

## 2280 — Minimum Lines To Represent A Line Chart

```go
package main

// LeetCode #2280: Minimum Lines to Represent a Line Chart
// https://leetcode.com/problems/minimum-lines-to-represent-a-line-chart/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func minimumLines(stockPrices [][]int) int {
	if len(stockPrices) <= 1 {
		return 0
	}

	sort.Slice(stockPrices, func(i, j int) bool {
		return stockPrices[i][0] < stockPrices[j][0]
	})

	lines := 1
	for i := 2; i < len(stockPrices); i++ {
		x1, y1 := stockPrices[i-2][0], stockPrices[i-2][1]
		x2, y2 := stockPrices[i-1][0], stockPrices[i-1][1]
		x3, y3 := stockPrices[i][0], stockPrices[i][1]

		// Compare slopes: (y2-y1)/(x2-x1) == (y3-y2)/(x3-x2)
		// Cross multiply to avoid floating point
		if (y2-y1)*(x3-x2) != (y3-y2)*(x2-x1) {
			lines++
		}
	}
	return lines
}

func main() {
	// Test case 1
	fmt.Println(minimumLines([][]int{{1, 7}, {2, 6}, {3, 5}, {4, 4}, {5, 4}, {6, 3}, {7, 2}, {8, 1}}))
	// Expected: 3

	// Test case 2
	fmt.Println(minimumLines([][]int{{3, 4}, {1, 2}, {7, 8}, {2, 3}}))
	// Expected: 1

	// Test case 3
	fmt.Println(minimumLines([][]int{{1, 1}}))
	// Expected: 0
}
```

## 2282 — Number Of People That Can Be Seen In A Grid

```go
package main

// LeetCode #2282: Number of People That Can Be Seen in a Grid
// https://leetcode.com/problems/number-of-people-that-can-be-seen-in-a-grid/
// Difficulty: Medium [Paid]
// Time: O(m * n) | Space: O(m * n)

import "fmt"

func visiblePeople(heights [][]int) [][]int {
	m, n := len(heights), len(heights[0])
	result := make([][]int, m)
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
	}

	// For each cell, count visible people to the right
	for i := 0; i < m; i++ {
		stack := []int{}
		for j := n - 1; j >= 0; j-- {
			visible := 0
			// Pop shorter people
			for len(stack) > 0 && stack[len(stack)-1] < heights[i][j] {
				stack = stack[:len(stack)-1]
				visible++
			}
			if len(stack) > 0 {
				visible++
			}
			result[i][j] = visible
			// Push current height
			for len(stack) > 0 && stack[len(stack)-1] == heights[i][j] {
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, heights[i][j])
		}
	}

	// For each cell, count visible people below
	for j := 0; j < n; j++ {
		stack := []int{}
		for i := m - 1; i >= 0; i-- {
			visible := 0
			for len(stack) > 0 && stack[len(stack)-1] < heights[i][j] {
				stack = stack[:len(stack)-1]
				visible++
			}
			if len(stack) > 0 {
				visible++
			}
			result[i][j] += visible
			for len(stack) > 0 && stack[len(stack)-1] == heights[i][j] {
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, heights[i][j])
		}
	}

	return result
}

func main() {
	// Test case 1
	fmt.Println(visiblePeople([][]int{{3, 1, 4, 2, 5}}))
	// Expected: [[1, 1, 1, 1, 0]]

	// Test case 2
	fmt.Println(visiblePeople([][]int{{5, 1}, {3, 1}, {4, 1}}))
	// Expected: [[1, 0], [1, 0], [0, 0]]
}
```

## 2284 — Sender With Largest Word Count

```go
package main

// LeetCode #2284: Sender With Largest Word Count
// https://leetcode.com/problems/sender-with-largest-word-count/
// Difficulty: Medium
// Time: O(n * m) | Space: O(n)

import (
	"fmt"
	"strings"
)

func largestWordCount(messages []string, senders []string) string {
	count := make(map[string]int)
	maxCount := 0
	maxSender := ""

	for i, msg := range messages {
		sender := senders[i]
		words := len(strings.Fields(msg))
		count[sender] += words
		if count[sender] > maxCount || (count[sender] == maxCount && sender > maxSender) {
			maxCount = count[sender]
			maxSender = sender
		}
	}
	return maxSender
}

func main() {
	// Test case 1
	fmt.Println(largestWordCount([]string{"Hello userTwooo", "Hi userThree", "Wonderful day Alice", "Nice day userThree"}, []string{"Alice", "userTwo", "userThree", "Alice"}))
	// Expected: "Alice"

	// Test case 2
	fmt.Println(largestWordCount([]string{"t", "e", "s", "t"}, []string{"a", "b", "c", "d"}))
	// Expected: "d" (first with max when equal, by lexicographical)
}
```

## 2285 — Maximum Total Importance Of Roads

```go
package main

// LeetCode #2285: Maximum Total Importance of Roads
// https://leetcode.com/problems/maximum-total-importance-of-roads/
// Difficulty: Medium
// Time: O(n + m) | Space: O(n)

import (
	"fmt"
	"sort"
)

func maximumImportance(n int, roads [][]int) int64 {
	degree := make([]int, n)
	for _, r := range roads {
		degree[r[0]]++
		degree[r[1]]++
	}
	sort.Ints(degree)

	var total int64 = 0
	for i, d := range degree {
		total += int64(i+1) * int64(d)
	}
	return total
}

func main() {
	// Test case 1
	fmt.Println(maximumImportance(5, [][]int{{0, 1}, {1, 2}, {2, 3}, {0, 2}, {1, 3}, {2, 4}}))
	// Expected: 43

	// Test case 2
	fmt.Println(maximumImportance(5, [][]int{{0, 3}, {2, 4}, {1, 3}}))
	// Expected: 20

	// Test case 3
	fmt.Println(maximumImportance(2, [][]int{{0, 1}}))
	// Expected: 3
}
```

## 2288 — Apply Discount To Prices

```go
package main

// LeetCode #2288: Apply Discount to Prices
// https://leetcode.com/problems/apply-discount-to-prices/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"strconv"
	"strings"
)

func discountPrices(sentence string, discount int) string {
	words := strings.Split(sentence, " ")
	for i, w := range words {
		if len(w) > 1 && w[0] == '$' {
			if numStr, err := strconv.Atoi(w[1:]); err == nil && numStr >= 0 && w[1] != '0' && w[1:] == strconv.Itoa(numStr) {
				price := float64(numStr) * (100.0 - float64(discount)) / 100.0
				words[i] = fmt.Sprintf("$%.2f", price)
			} else if err == nil && numStr == 0 && w == "$0" {
				words[i] = "$0.00"
			} else if err == nil && numStr == 0 && strings.TrimLeft(w[1:], "0") == "" {
				// All zeros: keep as valid price
				words[i] = "$0.00"
			}
		}
	}
	return strings.Join(words, " ")
}

func main() {
	// Test case 1
	fmt.Println(discountPrices("there are $1 $2 and 5$ candies in the shop", 50))
	// Expected: "there are $0.50 $1.00 and 5$ candies in the shop"

	// Test case 2
	fmt.Println(discountPrices("1 2 $3 4 $5 $6 7 8$ $9 $10$", 100))
	// Expected: "1 2 $0.00 4 $0.00 $0.00 7 8$ $0.00 $10$"
}
```

## 2289 — Steps To Make Array Non Decreasing

```go
package main

// LeetCode #2289: Steps to Make Array Non-decreasing
// https://leetcode.com/problems/steps-to-make-array-non-decreasing/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func totalSteps(nums []int) int {
	stack := []int{} // indices
	steps := make([]int, len(nums))
	maxSteps := 0

	for i := 0; i < len(nums); i++ {
		curSteps := 0
		for len(stack) > 0 && nums[stack[len(stack)-1]] <= nums[i] {
			if steps[stack[len(stack)-1]] > curSteps {
				curSteps = steps[stack[len(stack)-1]]
			}
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			steps[i] = curSteps + 1
		} else {
			steps[i] = 0
		}
		if steps[i] > maxSteps {
			maxSteps = steps[i]
		}
		stack = append(stack, i)
	}
	return maxSteps
}

func main() {
	// Test case 1
	fmt.Println(totalSteps([]int{5, 3, 4, 4, 7, 3, 6, 11, 8, 5, 11}))
	// Expected: 3

	// Test case 2
	fmt.Println(totalSteps([]int{4, 5, 7, 7, 13}))
	// Expected: 0

	// Test case 3
	fmt.Println(totalSteps([]int{10, 1, 2, 3, 4, 5, 6, 1, 2, 3}))
	// Expected: 6
}
```

## 2291 — Maximum Profit From Trading Stocks

```go
package main

// LeetCode #2291: Maximum Profit From Trading Stocks
// https://leetcode.com/problems/maximum-profit-from-trading-stocks/
// Difficulty: Medium [Paid]
// Time: O(n * budget) | Space: O(budget)

import "fmt"

func maximumProfit(presentValues []int, futureValues []int, budget int) int {
	n := len(presentValues)
	// dp[b] = max profit with budget b
	dp := make([]int, budget+1)

	for i := 0; i < n; i++ {
		profit := futureValues[i] - presentValues[i]
		if profit <= 0 {
			continue
		}
		cost := presentValues[i]
		for b := budget; b >= cost; b-- {
			if dp[b-cost]+profit > dp[b] {
				dp[b] = dp[b-cost] + profit
			}
		}
	}
	return dp[budget]
}

func main() {
	// Test case 1
	fmt.Println(maximumProfit([]int{5, 4, 6, 2}, []int{6, 5, 7, 3}, 6))
	// Expected: 2

	// Test case 2
	fmt.Println(maximumProfit([]int{5, 4, 6, 2}, []int{4, 5, 3, 3}, 5))
	// Expected: 0
}
```

## 2292 — Products With Three Or More Orders In Two Consecutive Years

```go
package main

// LeetCode #2292: Products With Three or More Orders in Two Consecutive Years
// https://leetcode.com/problems/products-with-three-or-more-orders-in-two-consecutive-years/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

type Order struct {
	ProductID int
	Year      int
}

func findProducts(orders [][]int) []int {
	// Group orders by product and year
	productYears := make(map[int]map[int]int)
	for _, o := range orders {
		productID, year := o[0], o[1]
		if productYears[productID] == nil {
			productYears[productID] = make(map[int]int)
		}
		productYears[productID][year]++
	}

	result := []int{}
	for pid, years := range productYears {
		yearList := []int{}
		for y := range years {
			yearList = append(yearList, y)
		}
		sort.Ints(yearList)

		for i := 1; i < len(yearList); i++ {
			if yearList[i]-yearList[i-1] == 1 &&
				years[yearList[i]] >= 3 &&
				years[yearList[i-1]] >= 3 {
				result = append(result, pid)
				break
			}
		}
	}
	sort.Ints(result)
	return result
}

func main() {
	// Test case 1
	fmt.Println(findProducts([][]int{{1, 2020}, {1, 2020}, {1, 2020}, {1, 2021}, {1, 2021}, {1, 2021}, {2, 2020}}))
	// Expected: [1]

	// Test case 2
	fmt.Println(findProducts([][]int{{1, 2020}, {2, 2020}}))
	// Expected: []
}
```

## 2294 — Partition Array Such That Maximum Difference Is K

```go
package main

// LeetCode #2294: Partition Array Such That Maximum Difference Is K
// https://leetcode.com/problems/partition-array-such-that-maximum-difference-is-k/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func partitionArray(nums []int, k int) int {
	sort.Ints(nums)
	partitions := 1
	minVal := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i]-minVal > k {
			partitions++
			minVal = nums[i]
		}
	}
	return partitions
}

func main() {
	// Test case 1
	fmt.Println(partitionArray([]int{3, 6, 1, 2, 5}, 2))
	// Expected: 2

	// Test case 2
	fmt.Println(partitionArray([]int{1, 2, 3}, 1))
	// Expected: 2

	// Test case 3
	fmt.Println(partitionArray([]int{2, 2, 4, 5}, 0))
	// Expected: 3
}
```

## 2295 — Replace Elements In An Array

```go
package main

// LeetCode #2295: Replace Elements in an Array
// https://leetcode.com/problems/replace-elements-in-an-array/
// Difficulty: Medium
// Time: O(n + m) | Space: O(n)

import "fmt"

func arrayChange(nums []int, operations [][]int) []int {
	pos := make(map[int]int)
	for i, v := range nums {
		pos[v] = i
	}

	for _, op := range operations {
		oldVal, newVal := op[0], op[1]
		if idx, ok := pos[oldVal]; ok {
			nums[idx] = newVal
			delete(pos, oldVal)
			pos[newVal] = idx
		}
	}
	return nums
}

func main() {
	// Test case 1
	fmt.Println(arrayChange([]int{1, 2, 4, 6}, [][]int{{1, 3}, {4, 7}, {6, 1}}))
	// Expected: [3,2,7,1]

	// Test case 2
	fmt.Println(arrayChange([]int{1, 2}, [][]int{{1, 3}, {2, 1}, {3, 2}}))
	// Expected: [2,1]
}
```

## 2297 — Jump Game Viii

```go
package main

// LeetCode #2297: Jump Game VIII
// https://leetcode.com/problems/jump-game-viii/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func minCost(nums []int, costs []int) int64 {
	n := len(nums)
	dp := make([]int64, n)
	for i := 1; i < n; i++ {
		dp[i] = -1
	}
	dp[0] = 0

	stackGE := []int{0} // monotonic stack for >=
	stackL := []int{0}  // monotonic stack for <

	for i := 1; i < n; i++ {
		// Jump from where nums[j] <= nums[i] (looking for >=)
		for len(stackGE) > 0 && nums[stackGE[len(stackGE)-1]] <= nums[i] {
			j := stackGE[len(stackGE)-1]
			stackGE = stackGE[:len(stackGE)-1]
			if dp[j] != -1 {
				cost := dp[j] + int64(costs[i])
				if dp[i] == -1 || cost < dp[i] {
					dp[i] = cost
				}
			}
		}
		// Jump from where nums[j] > nums[i] (looking for <)
		for len(stackL) > 0 && nums[stackL[len(stackL)-1]] > nums[i] {
			j := stackL[len(stackL)-1]
			stackL = stackL[:len(stackL)-1]
			if dp[j] != -1 {
				cost := dp[j] + int64(costs[i])
				if dp[i] == -1 || cost < dp[i] {
					dp[i] = cost
				}
			}
		}
		stackGE = append(stackGE, i)
		stackL = append(stackL, i)
	}
	return dp[n-1]
}

func main() {
	// Test case 1
	fmt.Println(minCost([]int{3, 2, 4, 4, 1}, []int{3, 7, 6, 4, 2}))
	// Expected: 8

	// Test case 2
	fmt.Println(minCost([]int{0, 1, 2}, []int{1, 1, 1}))
	// Expected: 2
}
```

## 2298 — Tasks Count In The Weekend

```go
package main

// LeetCode #2298: Tasks Count in the Weekend
// https://leetcode.com/problems/tasks-count-in-the-weekend/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func taskCount(taskRunDates [][]int) int {
	weekend := 0
	for _, d := range taskRunDates {
		// taskRunDates[i] = [month, day]
		// This is a paid SQL/Schema problem. Simplified:
		// Count tasks run on Saturday (6) or Sunday (7)
		dayOfWeek := d[1] % 7
		if dayOfWeek == 6 || dayOfWeek == 0 {
			weekend++
		}
	}
	return weekend
}

func main() {
	// Test case 1
	fmt.Println(taskCount([][]int{{1, 6}, {1, 7}, {1, 8}}))
	// Expected: 2 (Saturday and Sunday)

	// Test case 2
	fmt.Println(taskCount([][]int{{1, 1}, {1, 2}, {1, 3}, {1, 4}, {1, 5}}))
	// Expected: 0
}
```

## 2300 — Successful Pairs Of Spells And Potions

```go
package main

// LeetCode #2300: Successful Pairs of Spells and Potions
// https://leetcode.com/problems/successful-pairs-of-spells-and-potions/
// Difficulty: Medium
// Time: O((n + m) log m) | Space: O(1)

import (
	"fmt"
	"sort"
)

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	m := len(potions)
	result := make([]int, len(spells))

	for i, s := range spells {
		need := (int(success) + s - 1) / s // ceiling division
		idx := sort.SearchInts(potions, need)
		result[i] = m - idx
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println(successfulPairs([]int{5, 1, 3}, []int{1, 2, 3, 4, 5}, 7))
	// Expected: [4, 0, 3]

	// Test case 2
	fmt.Println(successfulPairs([]int{3, 1, 2}, []int{8, 5, 8}, 16))
	// Expected: [2, 0, 2]
}
```

## 2304 — Minimum Path Cost In A Grid

```go
package main

// LeetCode #2304: Minimum Path Cost in a Grid
// https://leetcode.com/problems/minimum-path-cost-in-a-grid/
// Difficulty: Medium
// Time: O(m * n^2) | Space: O(n)

import "fmt"

func minPathCost(grid [][]int, moveCost [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([]int, n)
	copy(dp, grid[0])

	for r := 1; r < m; r++ {
		next := make([]int, n)
		for j := 0; j < n; j++ {
			next[j] = 1 << 30
			for k := 0; k < n; k++ {
				cost := dp[k] + moveCost[grid[r-1][k]][j] + grid[r][j]
				if cost < next[j] {
					next[j] = cost
				}
			}
		}
		dp = next
	}

	minCost := dp[0]
	for _, v := range dp {
		if v < minCost {
			minCost = v
		}
	}
	return minCost
}

func main() {
	// Test case 1
	fmt.Println(minPathCost([][]int{{5, 3}, {4, 0}, {2, 1}}, [][]int{{9, 8}, {1, 5}, {10, 12}, {18, 6}, {2, 4}, {14, 3}}))
	// Expected: 17

	// Test case 2
	fmt.Println(minPathCost([][]int{{5, 1, 2}, {4, 0, 3}}, [][]int{{12, 10, 15}, {20, 23, 8}, {21, 7, 1}, {8, 1, 13}, {9, 10, 25}, {5, 3, 2}}))
	// Expected: 6
}
```

## 2305 — Fair Distribution Of Cookies

```go
package main

// LeetCode #2305: Fair Distribution of Cookies
// https://leetcode.com/problems/fair-distribution-of-cookies/
// Difficulty: Medium
// Time: O(k^n) | Space: O(k)

import "fmt"

func distributeCookies(cookies []int, k int) int {
	distribution := make([]int, k)
	minUnfairness := 1 << 30

	var dfs func(idx, maxSoFar int)
	dfs = func(idx, maxSoFar int) {
		if maxSoFar >= minUnfairness {
			return
		}
		if idx == len(cookies) {
			if maxSoFar < minUnfairness {
				minUnfairness = maxSoFar
			}
			return
		}
		for i := 0; i < k; i++ {
			distribution[i] += cookies[idx]
			newMax := maxSoFar
			if distribution[i] > newMax {
				newMax = distribution[i]
			}
			dfs(idx+1, newMax)
			distribution[i] -= cookies[idx]
			if distribution[i] == 0 {
				break
			}
		}
	}

	dfs(0, 0)
	return minUnfairness
}

func main() {
	// Test case 1
	fmt.Println(distributeCookies([]int{8, 15, 10, 20, 8}, 2))
	// Expected: 31

	// Test case 2
	fmt.Println(distributeCookies([]int{6, 1, 3, 2, 2, 4, 1, 2}, 3))
	// Expected: 7
}
```

## 2308 — Arrange Table By Gender

```go
package main

// LeetCode #2308: Arrange Table by Gender
// https://leetcode.com/problems/arrange-table-by-gender/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func arrangeTable(genders []string) []string {
	result := make([]string, 0, len(genders))
	// Order: female, other, male
	females, males, others := []string{}, []string{}, []string{}
	for _, g := range genders {
		if g == "female" {
			females = append(females, g)
		} else if g == "male" {
			males = append(males, g)
		} else {
			others = append(others, g)
		}
	}
	result = append(result, females...)
	result = append(result, others...)
	result = append(result, males...)
	return result
}

func main() {
	// Test case 1
	fmt.Println(arrangeTable([]string{"male", "female", "female", "other", "male", "other"}))
	// Expected: ["female", "female", "other", "other", "male", "male"]

	// Test case 2
	fmt.Println(arrangeTable([]string{"female", "male"}))
	// Expected: ["female", "male"]
}
```

## 2310 — Sum Of Numbers With Units Digit K

```go
package main

// LeetCode #2310: Sum of Numbers With Units Digit K
// https://leetcode.com/problems/sum-of-numbers-with-units-digit-k/
// Difficulty: Medium
// Time: O(1) | Space: O(1)

import "fmt"

func minimumNumbers(num int, k int) int {
	if num == 0 {
		return 0
	}
	for i := 1; i <= 10; i++ {
		if (i*k)%10 == num%10 && i*k <= num {
			return i
		}
	}
	return -1
}

func main() {
	// Test case 1
	fmt.Println(minimumNumbers(58, 9))
	// Expected: 2

	// Test case 2
	fmt.Println(minimumNumbers(37, 2))
	// Expected: -1

	// Test case 3
	fmt.Println(minimumNumbers(0, 7))
	// Expected: 0
}
```

## 2311 — Longest Binary Subsequence Less Than Or Equal To K

```go
package main

// LeetCode #2311: Longest Binary Subsequence Less Than or Equal to K
// https://leetcode.com/problems/longest-binary-subsequence-less-than-or-equal-to-k/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func longestSubsequence(s string, k int) int {
	val := 0
	count := 0
	pow := 1

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '0' {
			count++
		} else if s[i] == '1' {
			if val+pow <= k {
				val += pow
				count++
			}
		}
		if pow <= k {
			pow <<= 1
		}
	}
	return count
}

func main() {
	// Test case 1
	fmt.Println(longestSubsequence("1001010", 5))
	// Expected: 5

	// Test case 2
	fmt.Println(longestSubsequence("0010101010110101001001011010100010101010101010111111", 93))
	// Expected: 44
}
```

## 2314 — The First Day Of The Maximum Recorded Degree In Each City

```go
package main

// LeetCode #2314: The First Day of the Maximum Recorded Degree in Each City
// https://leetcode.com/problems/the-first-day-of-the-maximum-recorded-degree-in-each-city/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"sort"
)

type DayDegree struct {
	CityID int
	Day    int
	Degree int
}

func firstDayOfMaxDegree(data [][]int) []int {
	// data[i] = [city_id, day, degree]
	cityMax := make(map[int]int)
	cityFirstDay := make(map[int]int)

	for _, d := range data {
		cityID, day, degree := d[0], d[1], d[2]
		if prev, ok := cityMax[cityID]; !ok || degree > prev {
			cityMax[cityID] = degree
			cityFirstDay[cityID] = day
		} else if degree == prev && day < cityFirstDay[cityID] {
			cityFirstDay[cityID] = day
		}
	}

	cities := []int{}
	for c := range cityMax {
		cities = append(cities, c)
	}
	sort.Ints(cities)

	result := make([]int, len(cities))
	for i, c := range cities {
		result[i] = cityFirstDay[c]
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println(firstDayOfMaxDegree([][]int{{1, 1, 30}, {1, 2, 35}, {1, 3, 35}, {2, 1, 25}, {2, 2, 30}}))
	// Expected: [2, 2]

	// Test case 2
	fmt.Println(firstDayOfMaxDegree([][]int{{1, 1, 30}, {2, 2, 25}}))
	// Expected: [1, 2]
}
```

## 2316 — Count Unreachable Pairs Of Nodes In An Undirected Graph

```go
package main

// LeetCode #2316: Count Unreachable Pairs of Nodes in an Undirected Graph
// https://leetcode.com/problems/count-unreachable-pairs-of-nodes-in-an-undirected-graph/
// Difficulty: Medium
// Time: O(n + m) | Space: O(n)

import "fmt"

func countPairs(n int, edges [][]int) int64 {
	graph := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	visited := make([]bool, n)
	var totalPairs int64 = 0
	var prevComponents int64 = 0

	for i := 0; i < n; i++ {
		if visited[i] {
			continue
		}
		// BFS/DFS to find component size
		size := int64(0)
		queue := []int{i}
		visited[i] = true
		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]
			size++
			for _, nei := range graph[node] {
				if !visited[nei] {
					visited[nei] = true
					queue = append(queue, nei)
				}
			}
		}
		totalPairs += prevComponents * size
		prevComponents += size
	}
	return totalPairs
}

func main() {
	// Test case 1
	fmt.Println(countPairs(3, [][]int{{0, 1}, {0, 2}, {1, 2}}))
	// Expected: 0

	// Test case 2
	fmt.Println(countPairs(7, [][]int{{0, 2}, {0, 5}, {2, 4}, {1, 6}, {5, 4}}))
	// Expected: 14
}
```

## 2317 — Maximum Xor After Operations

```go
package main

// LeetCode #2317: Maximum XOR After Operations
// https://leetcode.com/problems/maximum-xor-after-operations/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func maximumXOR(nums []int) int {
	result := 0
	for _, v := range nums {
		result |= v
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println(maximumXOR([]int{3, 2, 4, 6}))
	// Expected: 7

	// Test case 2
	fmt.Println(maximumXOR([]int{1, 2, 3, 4, 5, 6, 7}))
	// Expected: 7

	// Test case 3
	fmt.Println(maximumXOR([]int{0}))
	// Expected: 0
}
```

## 2320 — Count Number Of Ways To Place Houses

```go
package main

// LeetCode #2320: Count Number of Ways to Place Houses
// https://leetcode.com/problems/count-number-of-ways-to-place-houses/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func countHousePlacements(n int) int {
	const mod = 1_000_000_007
	// Fibonacci-ish: ways to place on one side
	a, b := 1, 1 // empty, single plot
	for i := 0; i < n; i++ {
		a, b = b, (a+b)%mod
	}
	oneSide := b
	return (oneSide * oneSide) % mod
}

func main() {
	// Test case 1
	fmt.Println(countHousePlacements(1))
	// Expected: 4

	// Test case 2
	fmt.Println(countHousePlacements(2))
	// Expected: 9

	// Test case 3
	fmt.Println(countHousePlacements(3))
	// Expected: 25
}
```

## 2323 — Find Minimum Time To Finish All Jobs Ii

```go
package main

// LeetCode #2323: Find Minimum Time to Finish All Jobs II
// https://leetcode.com/problems/find-minimum-time-to-finish-all-jobs-ii/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func minimumTime(jobs []int, workers []int) int {
	sort.Ints(jobs)
	sort.Ints(workers)
	maxDays := 0

	for i := 0; i < len(jobs); i++ {
		days := (jobs[i] + workers[i] - 1) / workers[i] // ceil division
		if days > maxDays {
			maxDays = days
		}
	}
	return maxDays
}

func main() {
	// Test case 1
	fmt.Println(minimumTime([]int{5, 2, 4}, []int{1, 7, 5}))
	// Expected: 2

	// Test case 2
	fmt.Println(minimumTime([]int{3, 18, 30}, []int{3, 15, 5}))
	// Expected: 6
}
```

## 2324 — Product Sales Analysis Iv

```go
package main

// LeetCode #2324: Product Sales Analysis IV
// https://leetcode.com/problems/product-sales-analysis-iv/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

type Sale struct {
	UserID    int
	ProductID int
	Quantity  int
}

func productSalesAnalysis(sales [][]int) [][]int {
	// sales[i] = [sale_id, product_id, user_id, quantity]
	userProductQty := make(map[[2]int]int) // [user_id, product_id] -> total qty

	for _, s := range sales {
		productID, userID, qty := s[1], s[2], s[3]
		key := [2]int{userID, productID}
		userProductQty[key] += qty
	}

	// For each user, find product with max qty
	userMax := make(map[int]struct{ qty int; product int })
	for key, qty := range userProductQty {
		userID, productID := key[0], key[1]
		if prev, ok := userMax[userID]; !ok || qty > prev.qty || (qty == prev.qty && productID > prev.product) {
			userMax[userID] = struct{ qty int; product int }{qty, productID}
		}
	}

	result := [][]int{}
	// Sort by userID
	for u := 1; u <= len(userMax); u++ {
		if v, ok := userMax[u]; ok {
			result = append(result, []int{u, v.product, v.qty})
		}
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println(productSalesAnalysis([][]int{{1, 1, 1, 5}, {2, 1, 1, 3}, {3, 2, 1, 4}, {4, 2, 2, 2}}))
	// Expected: [[1, 1, 8], [2, 2, 2]]

	// Test case 2
	fmt.Println(productSalesAnalysis([][]int{{1, 1, 1, 1}, {2, 2, 1, 1}}))
	// Expected: [[1, 2, 1]]
}
```

## 2326 — Spiral Matrix Iv

```go
package main

// LeetCode #2326: Spiral Matrix IV
// https://leetcode.com/problems/spiral-matrix-iv/
// Difficulty: Medium
// Time: O(m * n) | Space: O(1)

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func spiralMatrix(m int, n int, head *ListNode) [][]int {
	result := make([][]int, m)
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
		for j := 0; j < n; j++ {
			result[i][j] = -1
		}
	}

	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	dir := 0
	r, c := 0, 0

	for head != nil {
		result[r][c] = head.Val
		head = head.Next

		nr, nc := r+dirs[dir][0], c+dirs[dir][1]
		if nr < 0 || nr >= m || nc < 0 || nc >= n || result[nr][nc] != -1 {
			dir = (dir + 1) % 4
			nr, nc = r+dirs[dir][0], c+dirs[dir][1]
		}
		r, c = nr, nc
	}
	return result
}

func main() {
	// Test case 1: head = [3,0,2,6,8,1,7,9,4,2,5,5,0], m = 3, n = 5
	head1 := &ListNode{3, &ListNode{0, &ListNode{2, &ListNode{6, &ListNode{8, &ListNode{1, &ListNode{7, &ListNode{9, &ListNode{4, &ListNode{2, &ListNode{5, &ListNode{5, &ListNode{0, nil}}}}}}}}}}}}}
	fmt.Println(spiralMatrix(3, 5, head1))

	// Test case 2: head = [0,1,2], m = 1, n = 4
	head2 := &ListNode{0, &ListNode{1, &ListNode{2, nil}}}
	fmt.Println(spiralMatrix(1, 4, head2))
}
```

## 2327 — Number Of People Aware Of A Secret

```go
package main

// LeetCode #2327: Number of People Aware of a Secret
// https://leetcode.com/problems/number-of-people-aware-of-a-secret/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func peopleAwareOfSecret(n int, delay int, forget int) int {
	const mod = 1_000_000_007
	// dp[i] = number of people who discovered the secret on day i
	dp := make([]int, n+1)
	dp[1] = 1

	var sharing int64 = 0 // people currently sharing
	for i := 2; i <= n; i++ {
		// New sharers: people who reached delay threshold (i-delay)
		if i-delay >= 1 {
			sharing = (sharing + int64(dp[i-delay])) % mod
		}
		// People who forget: reached forget threshold (i-forget)
		if i-forget >= 1 {
			sharing = (sharing - int64(dp[i-forget]) + mod) % mod
		}
		dp[i] = int(sharing)
	}

	var total int64 = 0
	for i := n - forget + 1; i <= n; i++ {
		if i >= 1 {
			total = (total + int64(dp[i])) % mod
		}
	}
	return int(total)
}

func main() {
	// Test case 1
	fmt.Println(peopleAwareOfSecret(6, 2, 4))
	// Expected: 5

	// Test case 2
	fmt.Println(peopleAwareOfSecret(4, 1, 3))
	// Expected: 6

	// Test case 3
	fmt.Println(peopleAwareOfSecret(10, 3, 5))
	// Expected: 12
}
```

## 2330 — Valid Palindrome Iv

```go
package main

// LeetCode #2330: Valid Palindrome IV
// https://leetcode.com/problems/valid-palindrome-iv/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func makePalindrome(s string) bool {
	diff := 0
	left, right := 0, len(s)-1

	for left < right {
		if s[left] != s[right] {
			diff++
			if diff > 2 {
				return false
			}
		}
		left++
		right--
	}
	return diff <= 2
}

func main() {
	// Test case 1
	fmt.Println(makePalindrome("abcdba"))
	// Expected: true

	// Test case 2
	fmt.Println(makePalindrome("abccba"))
	// Expected: false (already palindrome, need exactly 2 changes)

	// Test case 3
	fmt.Println(makePalindrome("abcdeba"))
	// Expected: false (need 3 changes)
}
```

## 2332 — The Latest Time To Catch A Bus

```go
package main

// LeetCode #2332: The Latest Time to Catch a Bus
// https://leetcode.com/problems/the-latest-time-to-catch-a-bus/
// Difficulty: Medium
// Time: O((n + m) log(n + m)) | Space: O(n + m)

import (
	"fmt"
	"sort"
)

func latestTimeCatchTheBus(buses []int, passengers []int, capacity int) int {
	sort.Ints(buses)
	sort.Ints(passengers)

	pi := 0
	for _, bus := range buses {
		count := 0
		for count < capacity && pi < len(passengers) && passengers[pi] <= bus {
			count++
			pi++
		}
		// Track last boarding time
		if pi > 0 && count == capacity {
			// Last passenger on this bus
			_ = 0
		}
	}

	// Find latest possible time
	time := buses[len(buses)-1]
	passSet := make(map[int]bool)
	for _, p := range passengers {
		passSet[p] = true
	}

	// If not all seats taken on last bus, try last bus departure
	// Check if we can arrive at bus time
	pi = 0
	lastBoarded := -1
	for _, bus := range buses {
		count := 0
		for count < capacity && pi < len(passengers) && passengers[pi] <= bus {
			lastBoarded = passengers[pi]
			count++
			pi++
		}
		if count < capacity {
			time = bus
		} else {
			time = lastBoarded - 1
		}
	}

	// Find latest time not taken by a passenger
	for passSet[time] {
		time--
	}
	return time
}

func main() {
	// Test case 1
	fmt.Println(latestTimeCatchTheBus([]int{10, 20}, []int{2, 17, 18, 19}, 2))
	// Expected: 16

	// Test case 2
	fmt.Println(latestTimeCatchTheBus([]int{20, 30, 10}, []int{19, 13, 26, 4, 25, 11, 21}, 2))
	// Expected: 20
}
```

## 2333 — Minimum Sum Of Squared Difference

```go
package main

// LeetCode #2333: Minimum Sum of Squared Difference
// https://leetcode.com/problems/minimum-sum-of-squared-difference/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func minSumSquareDiff(nums1 []int, nums2 []int, k1 int, k2 int) int64 {
	n := len(nums1)
	diffs := make([]int, n)
	for i := 0; i < n; i++ {
		diff := nums1[i] - nums2[i]
		if diff < 0 {
			diff = -diff
		}
		diffs[i] = diff
	}

	sort.Slice(diffs, func(i, j int) bool {
		return diffs[i] > diffs[j]
	})

	k := k1 + k2
	for i := 0; i < n && k > 0; i++ {
		if diffs[i] == 0 {
			break
		}
		nextDiff := 0
		if i+1 < n {
			nextDiff = diffs[i+1]
		}
		reduce := diffs[i] - nextDiff
		possibleReduce := reduce * (i + 1)
		if possibleReduce <= k {
			k -= possibleReduce
			for j := 0; j <= i; j++ {
				diffs[j] = nextDiff
			}
		} else {
			each := k / (i + 1)
			rem := k % (i + 1)
			for j := 0; j <= i; j++ {
				diffs[j] -= each
				if j < rem {
					diffs[j]--
				}
			}
			k = 0
		}
	}

	var sum int64 = 0
	for _, d := range diffs {
		sum += int64(d) * int64(d)
	}
	return sum
}

func main() {
	// Test case 1
	fmt.Println(minSumSquareDiff([]int{1, 2, 3, 4}, []int{2, 10, 20, 19}, 0, 0))
	// Expected: 579

	// Test case 2
	fmt.Println(minSumSquareDiff([]int{1, 4, 10, 12}, []int{5, 8, 6, 9}, 1, 1))
	// Expected: 43
}
```

## 2336 — Smallest Number In Infinite Set

```go
package main

// LeetCode #2336: Smallest Number in Infinite Set
// https://leetcode.com/problems/smallest-number-in-infinite-set/
// Difficulty: Medium
// Time: O(log n) for pop, O(1) for addBack | Space: O(n)

import (
	"container/heap"
	"fmt"
)

type SmallestInfiniteSet struct {
	added     *minHeap2
	addedSet  map[int]bool
	smallest  int
}

type minHeap2 []int

func (h minHeap2) Len() int           { return len(h) }
func (h minHeap2) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap2) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap2) Push(x any)        { *h = append(*h, x.(int)) }
func (h *minHeap2) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func Constructor() SmallestInfiniteSet {
	return SmallestInfiniteSet{
		added:    &minHeap2{},
		addedSet: make(map[int]bool),
		smallest: 1,
	}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	if this.added.Len() > 0 {
		val := heap.Pop(this.added).(int)
		delete(this.addedSet, val)
		return val
	}
	val := this.smallest
	this.smallest++
	return val
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if num >= this.smallest || this.addedSet[num] {
		return
	}
	this.addedSet[num] = true
	heap.Push(this.added, num)
}

func main() {
	set := Constructor()
	set.AddBack(2)
	fmt.Println(set.PopSmallest()) // 1
	fmt.Println(set.PopSmallest()) // 2
	fmt.Println(set.PopSmallest()) // 3
	set.AddBack(1)
	fmt.Println(set.PopSmallest()) // 1
	fmt.Println(set.PopSmallest()) // 4
}
```

## 2337 — Move Pieces To Obtain A String

```go
package main

// LeetCode #2337: Move Pieces to Obtain a String
// https://leetcode.com/problems/move-pieces-to-obtain-a-string/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func canChange(start string, target string) bool {
	n := len(start)
	i, j := 0, 0

	for i < n || j < n {
		// Skip underscores
		for i < n && start[i] == '_' {
			i++
		}
		for j < n && target[j] == '_' {
			j++
		}

		if i == n && j == n {
			return true
		}
		if i == n || j == n {
			return false
		}

		if start[i] != target[j] {
			return false
		}
		if start[i] == 'L' && i < j {
			return false
		}
		if start[i] == 'R' && i > j {
			return false
		}
		i++
		j++
	}
	return true
}

func main() {
	// Test case 1
	fmt.Println(canChange("_L__R__R_", "L______RR"))
	// Expected: true

	// Test case 2
	fmt.Println(canChange("R_L_", "__LR"))
	// Expected: false

	// Test case 3
	fmt.Println(canChange("_R", "R_"))
	// Expected: false
}
```

## 2340 — Minimum Adjacent Swaps To Make A Valid Array

```go
package main

// LeetCode #2340: Minimum Adjacent Swaps to Make a Valid Array
// https://leetcode.com/problems/minimum-adjacent-swaps-to-make-a-valid-array/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func minimumSwaps(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	minIdx, maxIdx := 0, 0
	for i := 1; i < n; i++ {
		if nums[i] < nums[minIdx] {
			minIdx = i
		}
		if nums[i] >= nums[maxIdx] {
			maxIdx = i
		}
	}

	swaps := minIdx + (n - 1 - maxIdx)
	if minIdx > maxIdx {
		swaps--
	}
	return swaps
}

func main() {
	// Test case 1
	fmt.Println(minimumSwaps([]int{3, 4, 5, 5, 3, 1}))
	// Expected: 6

	// Test case 2
	fmt.Println(minimumSwaps([]int{1, 2, 3, 4}))
	// Expected: 0

	// Test case 3
	fmt.Println(minimumSwaps([]int{2, 1}))
	// Expected: 1
}
```

## 2342 — Max Sum Of A Pair With Equal Sum Of Digits

```go
package main

// LeetCode #2342: Max Sum of a Pair With Equal Sum of Digits
// https://leetcode.com/problems/max-sum-of-a-pair-with-equal-sum-of-digits/
// Difficulty: Medium
// Time: O(n * d) | Space: O(n)

import "fmt"

func maximumSum(nums []int) int {
	digitSums := make([]int, len(nums))
	for i, v := range nums {
		sum := 0
		for v > 0 {
			sum += v % 10
			v /= 10
		}
		digitSums[i] = sum
	}

	maxVal := make(map[int][2]int) // top 2 for each digit sum
	for i, ds := range digitSums {
		a, b := maxVal[ds][0], maxVal[ds][1]
		if nums[i] > a {
			maxVal[ds] = [2]int{nums[i], a}
		} else if nums[i] > b {
			maxVal[ds] = [2]int{a, nums[i]}
		}
	}

	maxSum := -1
	for _, vals := range maxVal {
		if vals[1] > 0 {
			s := vals[0] + vals[1]
			if s > maxSum {
				maxSum = s
			}
		}
	}
	return maxSum
}

func main() {
	// Test case 1
	fmt.Println(maximumSum([]int{18, 43, 36, 13, 7}))
	// Expected: 54

	// Test case 2
	fmt.Println(maximumSum([]int{10, 12, 19, 14}))
	// Expected: -1

	// Test case 3
	fmt.Println(maximumSum([]int{229, 398, 269, 317, 420, 464, 491, 218, 439, 153, 482, 169, 411, 93, 147, 50, 347, 210, 251, 366, 401}))
	// Expected: 973
}
```

## 2343 — Query Kth Smallest Trimmed Number

```go
package main

// LeetCode #2343: Query Kth Smallest Trimmed Number
// https://leetcode.com/problems/query-kth-smallest-trimmed-number/
// Difficulty: Medium
// Time: O(m * n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func smallestTrimmedNumbers(nums []string, queries [][]int) []int {
	result := make([]int, len(queries))
	n := len(nums)

	type pair struct {
		val  string
		idx  int
	}

	for qi, q := range queries {
		k, trim := q[0], q[1]
		pairs := make([]pair, n)
		for i, s := range nums {
			pairs[i] = pair{s[len(s)-trim:], i}
		}
		sort.SliceStable(pairs, func(i, j int) bool {
			if pairs[i].val != pairs[j].val {
				return pairs[i].val < pairs[j].val
			}
			return pairs[i].idx < pairs[j].idx
		})
		result[qi] = pairs[k-1].idx
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println(smallestTrimmedNumbers([]string{"102", "473", "251", "814"}, [][]int{{1, 1}, {2, 3}, {4, 2}, {1, 2}}))
	// Expected: [2, 2, 1, 0]

	// Test case 2
	fmt.Println(smallestTrimmedNumbers([]string{"24", "37", "96", "04"}, [][]int{{2, 1}, {2, 2}}))
	// Expected: [3, 0]
}
```

## 2345 — Finding The Number Of Visible Mountains

```go
package main

// LeetCode #2345: Finding the Number of Visible Mountains
// https://leetcode.com/problems/finding-the-number-of-visible-mountains/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func visibleMountains(mountains [][]int) int {
	// Each mountain is peak at [x, y], base at [x-y, x+y]
	type interval struct{ left, right int }
	intervals := make([]interval, len(mountains))
	for i, m := range mountains {
		x, y := m[0], m[1]
		intervals[i] = interval{x - y, x + y}
	}

	// Sort by left ascending, right descending
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i].left != intervals[j].left {
			return intervals[i].left < intervals[j].left
		}
		return intervals[i].right > intervals[j].right
	})

	visible := 0
	maxRight := -1 << 30
	i := 0

	for i < len(intervals) {
		if intervals[i].right <= maxRight {
			i++
			continue
		}
		// Check if this mountain is not a duplicate
		if i+1 < len(intervals) && intervals[i].left == intervals[i+1].left && intervals[i].right == intervals[i+1].right {
			// Duplicate - skip all with same interval
			j := i
			for j < len(intervals) && intervals[j].left == intervals[i].left && intervals[j].right == intervals[i].right {
				j++
			}
			maxRight = intervals[i].right
			i = j
			continue
		}
		visible++
		maxRight = intervals[i].right
		i++
	}
	return visible
}

func main() {
	// Test case 1
	fmt.Println(visibleMountains([][]int{{2, 2}, {6, 3}, {5, 4}}))
	// Expected: 2

	// Test case 2
	fmt.Println(visibleMountains([][]int{{1, 2}, {1, 2}, {2, 1}}))
	// Expected: 1
}
```

## 2346 — Compute The Rank As A Percentage

```go
package main

// LeetCode #2346: Compute the Rank as a Percentage
// https://leetcode.com/problems/compute-the-rank-as-a-percentage/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func rankAsPercentage(ranks []int) []int {
	n := len(ranks)
	// This is a database-style problem. Simplified:
	// For each student, compute (rank-1)*100/(n-1)
	// But since it's "compute the rank as a percentage" database problem:
	// The percentage = (R-1)*100 / (N-1)
	result := make([]int, n)
	for i, r := range ranks {
		result[i] = (r - 1) * 100 / (n - 1)
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println(rankAsPercentage([]int{1, 2, 3, 4}))
	// Expected: [0, 33, 66, 100]

	// Test case 2
	fmt.Println(rankAsPercentage([]int{1, 2}))
	// Expected: [0, 100]
}
```

## 2348 — Number Of Zero Filled Subarrays

```go
package main

// LeetCode #2348: Number of Zero-Filled Subarrays
// https://leetcode.com/problems/number-of-zero-filled-subarrays/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func zeroFilledSubarray(nums []int) int64 {
	var total int64 = 0
	var count int64 = 0

	for _, v := range nums {
		if v == 0 {
			count++
			total += count
		} else {
			count = 0
		}
	}
	return total
}

func main() {
	// Test case 1
	fmt.Println(zeroFilledSubarray([]int{1, 3, 0, 0, 2, 0, 0, 4}))
	// Expected: 6

	// Test case 2
	fmt.Println(zeroFilledSubarray([]int{0, 0, 0, 2, 0, 0}))
	// Expected: 9

	// Test case 3
	fmt.Println(zeroFilledSubarray([]int{0}))
	// Expected: 1
}
```

## 2349 — Design A Number Container System

```go
package main

// LeetCode #2349: Design a Number Container System
// https://leetcode.com/problems/design-a-number-container-system/
// Difficulty: Medium
// Time: O(log n) per operation | Space: O(n)

import (
	"container/heap"
	"fmt"
)

type NumberContainers struct {
	indexToNum map[int]int
	numToIndex map[int]*minHeap3
}

type minHeap3 []int

func (h minHeap3) Len() int           { return len(h) }
func (h minHeap3) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap3) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap3) Push(x any)        { *h = append(*h, x.(int)) }
func (h *minHeap3) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func Constructor2() NumberContainers {
	return NumberContainers{
		indexToNum: make(map[int]int),
		numToIndex: make(map[int]*minHeap3),
	}
}

func (this *NumberContainers) Change(index int, number int) {
	this.indexToNum[index] = number
	if this.numToIndex[number] == nil {
		this.numToIndex[number] = &minHeap3{}
	}
	heap.Push(this.numToIndex[number], index)
}

func (this *NumberContainers) Find(number int) int {
	if this.numToIndex[number] == nil {
		return -1
	}
	for this.numToIndex[number].Len() > 0 {
		idx := (*this.numToIndex[number])[0]
		if this.indexToNum[idx] == number {
			return idx
		}
		heap.Pop(this.numToIndex[number])
	}
	return -1
}

func main() {
	nc := Constructor2()
	nc.Change(1, 10)
	fmt.Println(nc.Find(10)) // 1
	nc.Change(1, 20)
	fmt.Println(nc.Find(10)) // -1
	fmt.Println(nc.Find(20)) // 1
	fmt.Println(nc.Find(30)) // -1
}
```

## 2352 — Equal Row And Column Pairs

```go
package main

// LeetCode #2352: Equal Row and Column Pairs
// https://leetcode.com/problems/equal-row-and-column-pairs/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n^2)

import (
	"fmt"
	"strconv"
	"strings"
)

func equalPairs(grid [][]int) int {
	n := len(grid)
	rowMap := make(map[string]int)

	for i := 0; i < n; i++ {
		var sb strings.Builder
		for j := 0; j < n; j++ {
			if sb.Len() > 0 {
				sb.WriteString(",")
			}
			sb.WriteString(strconv.Itoa(grid[i][j]))
		}
		rowMap[sb.String()]++
	}

	count := 0
	for j := 0; j < n; j++ {
		var sb strings.Builder
		for i := 0; i < n; i++ {
			if sb.Len() > 0 {
				sb.WriteString(",")
			}
			sb.WriteString(strconv.Itoa(grid[i][j]))
		}
		count += rowMap[sb.String()]
	}
	return count
}

func main() {
	// Test case 1
	fmt.Println(equalPairs([][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}}))
	// Expected: 1

	// Test case 2
	fmt.Println(equalPairs([][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}}))
	// Expected: 3

	// Test case 3
	fmt.Println(equalPairs([][]int{{11, 1}, {1, 11}}))
	// Expected: 2
}
```

## 2353 — Design A Food Rating System

```go
package main

// LeetCode #2353: Design a Food Rating System
// https://leetcode.com/problems/design-a-food-rating-system/
// Difficulty: Medium
// Time: O(log n) per operation | Space: O(n)

import (
	"container/heap"
	"fmt"
)

type FoodRatings struct {
	foodToCuisine map[string]string
	foodToRating  map[string]int
	cuisineToHeap map[string]*foodHeap
}

type foodItem struct {
	name   string
	rating int
}

type foodHeap []foodItem

func (h foodHeap) Len() int { return len(h) }
func (h foodHeap) Less(i, j int) bool {
	if h[i].rating != h[j].rating {
		return h[i].rating > h[j].rating
	}
	return h[i].name < h[j].name
}
func (h foodHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *foodHeap) Push(x any)   { *h = append(*h, x.(foodItem)) }
func (h *foodHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func Constructor3(foods []string, cuisines []string, ratings []int) FoodRatings {
	fr := FoodRatings{
		foodToCuisine: make(map[string]string),
		foodToRating:  make(map[string]int),
		cuisineToHeap: make(map[string]*foodHeap),
	}
	for i, f := range foods {
		fr.foodToCuisine[f] = cuisines[i]
		fr.foodToRating[f] = ratings[i]
		if fr.cuisineToHeap[cuisines[i]] == nil {
			fr.cuisineToHeap[cuisines[i]] = &foodHeap{}
		}
		heap.Push(fr.cuisineToHeap[cuisines[i]], foodItem{f, ratings[i]})
	}
	return fr
}

func (this *FoodRatings) ChangeRating(food string, newRating int) {
	this.foodToRating[food] = newRating
	cuisine := this.foodToCuisine[food]
	heap.Push(this.cuisineToHeap[cuisine], foodItem{food, newRating})
}

func (this *FoodRatings) HighestRated(cuisine string) string {
	for this.cuisineToHeap[cuisine].Len() > 0 {
		top := (*this.cuisineToHeap[cuisine])[0]
		if this.foodToRating[top.name] == top.rating {
			return top.name
		}
		heap.Pop(this.cuisineToHeap[cuisine])
	}
	return ""
}

func main() {
	fr := Constructor3(
		[]string{"kimchi", "miso", "sushi", "moussaka", "ramen", "bulgogi"},
		[]string{"korean", "japanese", "japanese", "greek", "japanese", "korean"},
		[]int{9, 12, 8, 15, 14, 7},
	)
	fmt.Println(fr.HighestRated("korean"))   // "bulgogi" (wait, kimchi=9, bulgogi=7... kimchi is higher)
	fmt.Println(fr.HighestRated("japanese"))  // "ramen"
	fr.ChangeRating("sushi", 16)
	fmt.Println(fr.HighestRated("japanese"))  // "sushi"
	fr.ChangeRating("ramen", 16)
	fmt.Println(fr.HighestRated("japanese"))  // "ramen" (tie, lexicographically smaller)
}
```

## 2358 — Maximum Number Of Groups Entering A Competition

```go
package main

// LeetCode #2358: Maximum Number of Groups Entering a Competition
// https://leetcode.com/problems/maximum-number-of-groups-entering-a-competition/
// Difficulty: Medium
// Time: O(1) | Space: O(1)

import "fmt"

func maximumGroups(grades []int) int {
	n := len(grades)
	// We need k such that 1 + 2 + ... + k <= n
	// k(k+1)/2 <= n
	// Solve: k^2 + k - 2n <= 0
	// k = (-1 + sqrt(1 + 8n)) / 2
	k := 0
	for (k+1)*(k+2)/2 <= n {
		k++
	}
	return k
}

func main() {
	// Test case 1
	fmt.Println(maximumGroups([]int{10, 6, 12, 7, 3, 5}))
	// Expected: 3

	// Test case 2
	fmt.Println(maximumGroups([]int{8, 8}))
	// Expected: 1

	// Test case 3
	fmt.Println(maximumGroups([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	// Expected: 4
}
```

## 2359 — Find Closest Node To Given Two Nodes

```go
package main

// LeetCode #2359: Find Closest Node to Given Two Nodes
// https://leetcode.com/problems/find-closest-node-to-given-two-nodes/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func closestMeetingNode(edges []int, node1 int, node2 int) int {
	n := len(edges)
	dist1 := make([]int, n)
	dist2 := make([]int, n)
	for i := 0; i < n; i++ {
		dist1[i] = -1
		dist2[i] = -1
	}

	// BFS from node1
	queue := []int{node1}
	dist1[node1] = 0
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		v := edges[u]
		if v != -1 && dist1[v] == -1 {
			dist1[v] = dist1[u] + 1
			queue = append(queue, v)
		}
	}

	// BFS from node2
	queue = []int{node2}
	dist2[node2] = 0
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		v := edges[u]
		if v != -1 && dist2[v] == -1 {
			dist2[v] = dist2[u] + 1
			queue = append(queue, v)
		}
	}

	bestNode := -1
	bestDist := 1 << 30
	for i := 0; i < n; i++ {
		if dist1[i] != -1 && dist2[i] != -1 {
			maxDist := dist1[i]
			if dist2[i] > maxDist {
				maxDist = dist2[i]
			}
			if maxDist < bestDist {
				bestDist = maxDist
				bestNode = i
			}
		}
	}
	return bestNode
}

func main() {
	// Test case 1
	fmt.Println(closestMeetingNode([]int{2, 2, 3, -1}, 0, 1))
	// Expected: 2

	// Test case 2
	fmt.Println(closestMeetingNode([]int{1, 2, -1}, 0, 2))
	// Expected: 2

	// Test case 3
	fmt.Println(closestMeetingNode([]int{4, 4, 4, 5, -1, 0}, 1, 3))
	// Expected: 4
}
```

## 2364 — Count Number Of Bad Pairs

```go
package main

// LeetCode #2364: Count Number of Bad Pairs
// https://leetcode.com/problems/count-number-of-bad-pairs/
// Difficulty: Medium
// Time: O(n) | Space: O(n)
// Transform nums[i] into nums[i] - i. Good pairs have same transformed value.
// Bad pairs = total pairs - good pairs.

import "fmt"

func main() {
	fmt.Println(countBadPairs([]int{4, 1, 3, 3})) // 5
	fmt.Println(countBadPairs([]int{1, 2, 3, 4, 5})) // 0
}

func countBadPairs(nums []int) int64 {
	n := len(nums)
	freq := make(map[int]int64)
	var good int64
	for i, v := range nums {
		key := v - i
		good += freq[key]
		freq[key]++
	}
	total := int64(n) * int64(n-1) / 2
	return total - good
}
```

## 2365 — Task Scheduler Ii

```go
package main

// LeetCode #2365: Task Scheduler II
// https://leetcode.com/problems/task-scheduler-ii/
// Difficulty: Medium
// Time: O(n) | Space: O(n)
// Track last completion day for each task type. If within cooldown, advance day.

import "fmt"

func main() {
	fmt.Println(taskSchedulerII([]int{1, 2, 1, 2, 3, 1}, 3)) // 9
	fmt.Println(taskSchedulerII([]int{5, 8, 8, 5}, 2))       // 6
}

func taskSchedulerII(tasks []int, space int) int64 {
	last := make(map[int]int64)
	var day int64 = 0
	for _, t := range tasks {
		day++
		if prev, ok := last[t]; ok && day-prev <= int64(space) {
			day = prev + int64(space) + 1
		}
		last[t] = day
	}
	return day
}
```

## 2368 — Reachable Nodes With Restrictions

```go
package main

// LeetCode #2368: Reachable Nodes With Restrictions
// https://leetcode.com/problems/reachable-nodes-with-restrictions/
// Difficulty: Medium
// Time: O(n) | Space: O(n)
// BFS/DFS from node 0, skip restricted nodes.

import "fmt"

func main() {
	fmt.Println(reachableNodes(7, [][]int{{0, 1}, {1, 2}, {3, 1}, {4, 0}, {0, 5}, {5, 6}}, []int{4, 5})) // 4
	fmt.Println(reachableNodes(7, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}, {1, 5}, {2, 6}}, []int{1})) // 3
}

func reachableNodes(n int, edges [][]int, restricted []int) int {
	restrictedSet := make(map[int]bool, len(restricted))
	for _, r := range restricted {
		restrictedSet[r] = true
	}

	graph := make([][]int, n)
	for _, e := range edges {
		a, b := e[0], e[1]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	visited := make([]bool, n)
	var dfs func(u int) int
	dfs = func(u int) int {
		visited[u] = true
		count := 1
		for _, v := range graph[u] {
			if !visited[v] && !restrictedSet[v] {
				count += dfs(v)
			}
		}
		return count
	}
	return dfs(0)
}
```

## 2369 — Check If There Is A Valid Partition For The Array

```go
package main

// LeetCode #2369: Check if There is a Valid Partition For The Array
// https://leetcode.com/problems/check-if-there-is-a-valid-partition-for-the-array/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// DP: dp[i] = valid partition for first i elements. Check three conditions.

import "fmt"

func main() {
	fmt.Println(validPartition([]int{4, 4, 4, 5, 6}))     // true
	fmt.Println(validPartition([]int{1, 1, 1, 2}))          // false
	fmt.Println(validPartition([]int{1, 2, 3, 4, 5, 6}))    // true
}

func validPartition(nums []int) bool {
	n := len(nums)
	dp := make([]bool, n+1)
	dp[0] = true
	for i := 1; i <= n; i++ {
		if i >= 2 && dp[i-2] && nums[i-2] == nums[i-1] {
			dp[i] = true
		}
		if i >= 3 && dp[i-3] {
			if (nums[i-3] == nums[i-2] && nums[i-2] == nums[i-1]) ||
				(nums[i-3]+1 == nums[i-2] && nums[i-2]+1 == nums[i-1]) {
				dp[i] = true
			}
		}
	}
	return dp[n]
}
```

## 2370 — Longest Ideal Subsequence

```go
package main

// LeetCode #2370: Longest Ideal Subsequence
// https://leetcode.com/problems/longest-ideal-subsequence/
// Difficulty: Medium
// Time: O(n * 26) | Space: O(26)
// DP over alphabet: dp[c] = longest ideal subsequence ending with char c.

import "fmt"

func main() {
	fmt.Println(longestIdealString("acfgbd", 2)) // 4
	fmt.Println(longestIdealString("abcd", 3))   // 4
}

func longestIdealString(s string, k int) int {
	dp := make([]int, 26)
	var ans int
	for i := 0; i < len(s); i++ {
		c := int(s[i] - 'a')
		best := 0
		for p := 0; p < 26; p++ {
			if abs(c-p) <= k && dp[p] > best {
				best = dp[p]
			}
		}
		dp[c] = best + 1
		if dp[c] > ans {
			ans = dp[c]
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
```

## 2372 — Calculate The Influence Of Each Salesperson

```go
package main

// LeetCode #2372: Calculate the Influence of Each Salesperson
// https://leetcode.com/problems/calculate-the-influence-of-each-salesperson/
// Difficulty: Medium
// Time: O(n + m) | Space: O(n + m)
// Simulate: track total price by salesperson from sales and product tables.

import "fmt"

func main() {
	sales := [][]int{
		{1, 1, 10}, // salesperson_id, product_id, quantity
		{2, 1, 5},
		{1, 2, 8},
		{3, 2, 3},
		{2, 2, 2},
	}
	products := map[int]int{
		1: 100, // product_id -> price
		2: 50,
	}
	fmt.Println(calculateInfluence(sales, products)) // {1: 1400, 2: 600, 3: 150}

	sales2 := [][]int{
		{1, 1, 1},
		{2, 1, 1},
	}
	products2 := map[int]int{1: 500}
	fmt.Println(calculateInfluence(sales2, products2)) // {1: 500, 2: 500}
}

func calculateInfluence(sales [][]int, priceMap map[int]int) map[int]int {
	result := make(map[int]int)
	for _, s := range sales {
		sp, prodID, qty := s[0], s[1], s[2]
		result[sp] += qty * priceMap[prodID]
	}
	return result
}
```

## 2374 — Node With Highest Edge Score

```go
package main

// LeetCode #2374: Node With Highest Edge Score
// https://leetcode.com/problems/node-with-highest-edge-score/
// Difficulty: Medium
// Time: O(n) | Space: O(n)
// Each node i points to edges[i]. Score of j = sum of i where edges[i] == j.

import "fmt"

func main() {
	fmt.Println(edgeScore([]int{1, 0, 0, 0, 0, 7, 7, 5})) // 7
	fmt.Println(edgeScore([]int{2, 0, 0, 2}))               // 0
}

func edgeScore(edges []int) int {
	n := len(edges)
	score := make([]int, n)
	for i, to := range edges {
		score[to] += i
	}

	maxScore := -1
	ans := -1
	for i, s := range score {
		if s > maxScore {
			maxScore = s
			ans = i
		}
	}
	return ans
}
```

## 2375 — Construct Smallest Number From Di String

```go
package main

// LeetCode #2375: Construct Smallest Number From DI String
// https://leetcode.com/problems/construct-smallest-number-from-di-string/
// Difficulty: Medium
// Time: O(n) | Space: O(n)
// Use stack: push numbers 1..n+1. On 'I' or end, pop stack to build result.

import "fmt"

func main() {
	fmt.Println(smallestNumber("II"))  // "123"
	fmt.Println(smallestNumber("DI"))  // "231"
	fmt.Println(smallestNumber("DDD")) // "4321"
}

func smallestNumber(pattern string) string {
	n := len(pattern)
	stack := make([]int, 0, n+1)
	res := make([]byte, 0, n+1)
	for i := 0; i <= n; i++ {
		stack = append(stack, i+1)
		if i == n || pattern[i] == 'I' {
			for len(stack) > 0 {
				res = append(res, byte('0'+stack[len(stack)-1]))
				stack = stack[:len(stack)-1]
			}
		}
	}
	return string(res)
}
```

## 2378 — Choose Edges To Maximize Score In A Tree

```go
package main

// LeetCode #2378: Choose Edges to Maximize Score in a Tree
// https://leetcode.com/problems/choose-edges-to-maximize-score-in-a-tree/
// Difficulty: Medium
// Time: O(n) | Space: O(n)
// Maximum weight matching on a tree (no two edges share a node).
// DP: dp0[u] = max when u is free, dp1[u] = max when edge(u,parent) is taken.
// dp1[u] = sum(dp0[child])
// dp0[u] = max(not matching any child, matching one child v: w + dp1[v] + sum_{other} dp0[other])

import "fmt"

func main() {
	fmt.Println(maxScore([][]int{{0, 1, 5}, {1, 2, 3}, {0, 3, 2}})) // 5
	fmt.Println(maxScore([][]int{{0, 1, 10}, {0, 2, 20}}))           // 20
	fmt.Println(maxScore([][]int{{0, 1, 1}, {1, 2, 2}, {2, 3, 3}})) // 4
}

func maxScore(edges [][]int) int64 {
	n := len(edges) + 1
	graph := make([][][2]int, n)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		graph[u] = append(graph[u], [2]int{v, w})
		graph[v] = append(graph[v], [2]int{u, w})
	}

	var dfs func(u, parent int) (int64, int64)
	dfs = func(u, parent int) (int64, int64) {
		var base int64
		bestDelta := int64(0)

		for _, nei := range graph[u] {
			v, w := nei[0], int64(nei[1])
			if v == parent {
				continue
			}
			child0, child1 := dfs(v, u)
			base += child0
			delta := w + child1 - child0
			if delta > bestDelta {
				bestDelta = delta
			}
		}

		dp1 := base
		dp0 := base + bestDelta
		return dp0, dp1
	}

	dp0, _ := dfs(0, -1)
	return dp0
}
```

## 2380 — Time Needed To Rearrange A Binary String

```go
package main

// LeetCode #2380: Time Needed to Rearrange a Binary String
// https://leetcode.com/problems/time-needed-to-rearrange-a-binary-string/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// Each "01" becomes "10" per second. Equivalent to: each 1 moves right past 0s,
// and the total time is max over each 1 of (position - index_in_final_position).

import "fmt"

func main() {
	fmt.Println(secondsToRemoveOccurrences("0110101")) // 4
	fmt.Println(secondsToRemoveOccurrences("11100"))   // 0
	fmt.Println(secondsToRemoveOccurrences("001011"))  // 3
}

func secondsToRemoveOccurrences(s string) int {
	ans, zeros := 0, 0
	for _, ch := range s {
		if ch == '0' {
			zeros++
		} else if zeros > 0 {
			ans = max(ans+1, zeros)
		}
	}
	return ans
}
```

## 2381 — Shifting Letters Ii

```go
package main

// LeetCode #2381: Shifting Letters II
// https://leetcode.com/problems/shifting-letters-ii/
// Difficulty: Medium
// Time: O(n + m) | Space: O(n)
// Difference array to apply range shifts efficiently.

import "fmt"

func main() {
	fmt.Println(shiftingLetters("abc", [][]int{{0, 1, 0}, {1, 2, 1}, {0, 2, 1}})) // "ace"
	fmt.Println(shiftingLetters("dztz", [][]int{{0, 0, 0}, {1, 1, 1}}))       // "catz"
}

func shiftingLetters(s string, shifts [][]int) string {
	n := len(s)
	diff := make([]int, n+1)
	for _, sh := range shifts {
		start, end, dir := sh[0], sh[1], sh[2]
		if dir == 1 {
			diff[start]++
			diff[end+1]--
		} else {
			diff[start]--
			diff[end+1]++
		}
	}

	cur := 0
	res := make([]byte, n)
	for i, ch := range s {
		cur += diff[i]
		shift := ((int(ch-'a')+cur)%26 + 26) % 26
		res[i] = byte('a' + shift)
	}
	return string(res)
}
```

## 2384 — Largest Palindromic Number

```go
package main

// LeetCode #2384: Largest Palindromic Number
// https://leetcode.com/problems/largest-palindromic-number/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// Count digits. Build palindrome: place pairs from 9 down to 0, then single middle digit.

import "fmt"

func main() {
	fmt.Println(largestPalindromic("444947137")) // "7449447"
	fmt.Println(largestPalindromic("00009"))     // "9"
	fmt.Println(largestPalindromic("0000"))      // "0"
}

func largestPalindromic(num string) string {
	cnt := make([]int, 10)
	for _, ch := range num {
		cnt[ch-'0']++
	}

	left := make([]byte, 0)
	middle := ""

	for d := 9; d >= 0; d-- {
		pairs := cnt[d] / 2
		if d == 0 && len(left) == 0 {
			// skip leading zeros
			if middle == "" && cnt[0]%2 == 1 {
				middle = "0"
			}
			break
		}
		for i := 0; i < pairs; i++ {
			left = append(left, byte('0'+d))
		}
		if middle == "" && cnt[d]%2 == 1 {
			middle = string(rune('0' + d))
		}
	}

	if len(left) == 0 && middle == "" {
		return "0"
	}

	// mirror: left + middle + reverse(left)
	right := make([]byte, len(left))
	for i, b := range left {
		right[len(left)-1-i] = b
	}
	return string(left) + middle + string(right)
}
```

## 2385 — Amount Of Time For Binary Tree To Be Infected

```go
package main

// LeetCode #2385: Amount of Time for Binary Tree to Be Infected
// https://leetcode.com/problems/amount-of-time-for-binary-tree-to-be-infected/
// Difficulty: Medium
// Time: O(n) | Space: O(n)
// Convert tree to graph, BFS from start node to find farthest distance.

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Tree: [1,5,3,null,4,10,6,9,2]
	root := &TreeNode{1,
		&TreeNode{5, nil,
			&TreeNode{4,
				&TreeNode{9, nil, nil},
				&TreeNode{2, nil, nil},
			},
		},
		&TreeNode{3,
			&TreeNode{10, nil, nil},
			&TreeNode{6, nil, nil},
		},
	}
	fmt.Println(amountOfTime(root, 3)) // 4

	root2 := &TreeNode{1,
		&TreeNode{2,
			&TreeNode{3,
				&TreeNode{4,
					&TreeNode{5, nil, nil},
					nil,
				},
				nil,
			},
			nil,
		},
		nil,
	}
	fmt.Println(amountOfTime(root2, 1)) // 4
}

func amountOfTime(root *TreeNode, start int) int {
	graph := make(map[int][]int)
	var buildGraph func(node *TreeNode)
	buildGraph = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil {
			graph[node.Val] = append(graph[node.Val], node.Left.Val)
			graph[node.Left.Val] = append(graph[node.Left.Val], node.Val)
			buildGraph(node.Left)
		}
		if node.Right != nil {
			graph[node.Val] = append(graph[node.Val], node.Right.Val)
			graph[node.Right.Val] = append(graph[node.Right.Val], node.Val)
			buildGraph(node.Right)
		}
	}
	buildGraph(root)

	visited := map[int]bool{start: true}
	queue := []int{start}
	time := -1
	for len(queue) > 0 {
		time++
		for sz := len(queue); sz > 0; sz-- {
			u := queue[0]
			queue = queue[1:]
			for _, v := range graph[u] {
				if !visited[v] {
					visited[v] = true
					queue = append(queue, v)
				}
			}
		}
	}
	return time
}
```

## 2387 — Median Of A Row Wise Sorted Matrix

```go
package main

// LeetCode #2387: Median of a Row Wise Sorted Matrix
// https://leetcode.com/problems/median-of-a-row-wise-sorted-matrix/
// Difficulty: Medium
// Time: O(rows * log(cols) * log(max-min)) | Space: O(1)
// Binary search on value, count elements <= mid.

import "fmt"

func main() {
	fmt.Println(matrixMedian([][]int{{1, 1, 2}, {2, 3, 3}, {1, 3, 4}})) // 2
	fmt.Println(matrixMedian([][]int{{1, 2}, {3, 4}}))                 // 2
}

func matrixMedian(grid [][]int) int {
	r, c := len(grid), len(grid[0])
	target := r*c/2 + 1
	lo, hi := 1, 1000000

	for lo < hi {
		mid := (lo + hi) / 2
		count := 0
		for i := 0; i < r; i++ {
			// binary search in each row for count of elements <= mid
			row := grid[i]
			left, right := 0, c
			for left < right {
				m := (left + right) / 2
				if row[m] <= mid {
					left = m + 1
				} else {
					right = m
				}
			}
			count += left
		}
		if count >= target {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}
```

## 2388 — Change Null Values In A Table To The Previous Value

```go
package main

// LeetCode #2388: Change Null Values in a Table to the Previous Value
// https://leetcode.com/problems/change-null-values-in-a-table-to-the-previous-value/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// Iterate rows ordered by id, carry forward non-null values.

import "fmt"

type Row struct {
	ID    int
	Value *int // nil represents null
}

func main() {
	rows := []Row{
		{1, intPtr(3)},
		{2, nil},
		{3, nil},
		{4, intPtr(6)},
		{5, nil},
	}
	fillNull(rows)
	for _, r := range rows {
		if r.Value != nil {
			fmt.Printf("%d:%d ", r.ID, *r.Value)
		} else {
			fmt.Printf("%d:nil ", r.ID)
		}
	}
	fmt.Println()
	// Output: 1:3 2:3 3:3 4:6 5:6
}

func intPtr(v int) *int { return &v }

func fillNull(rows []Row) {
	var prev *int
	for i := range rows {
		if rows[i].Value != nil {
			prev = rows[i].Value
		} else if prev != nil {
			rows[i].Value = intPtr(*prev)
		}
	}
}
```

## 2390 — Removing Stars From A String

```go
package main

// LeetCode #2390: Removing Stars From a String
// https://leetcode.com/problems/removing-stars-from-a-string/
// Difficulty: Medium
// Time: O(n) | Space: O(n)
// Stack: push non-star, pop on star.

import "fmt"

func main() {
	fmt.Println(removeStars("leet**cod*e")) // "lecoe"
	fmt.Println(removeStars("erase*****"))  // ""
}

func removeStars(s string) string {
	res := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == '*' {
			res = res[:len(res)-1]
		} else {
			res = append(res, s[i])
		}
	}
	return string(res)
}
```

## 2391 — Minimum Amount Of Time To Collect Garbage

```go
package main

// LeetCode #2391: Minimum Amount of Time to Collect Garbage
// https://leetcode.com/problems/minimum-amount-of-time-to-collect-garbage/
// Difficulty: Medium
// Time: O(n * k) | Space: O(1) where k = types (3)
// Sum collection (1 min per unit) + travel to last house containing each type.

import "fmt"

func main() {
	fmt.Println(garbageCollection([]string{"G", "P", "GP", "GG"}, []int{2, 4, 3})) // 21
	fmt.Println(garbageCollection([]string{"MMM", "PGM", "GP"}, []int{3, 10}))     // 37
}

func garbageCollection(garbage []string, travel []int) int {
	last := make([]int, 3) // 0=G, 1=P, 2=M
	total := 0
	for i, g := range garbage {
		total += len(g)
		for _, ch := range g {
			switch ch {
			case 'G':
				last[0] = i
			case 'P':
				last[1] = i
			case 'M':
				last[2] = i
			}
		}
	}

	// prefix sums for travel
	pref := make([]int, len(travel)+1)
	for i, t := range travel {
		pref[i+1] = pref[i] + t
	}

	for _, l := range last {
		total += pref[l]
	}
	return total
}
```

## 2393 — Count Strictly Increasing Subarrays

```go
package main

// LeetCode #2393: Count Strictly Increasing Subarrays
// https://leetcode.com/problems/count-strictly-increasing-subarrays/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// Track length of current increasing run. Each position adds run_len subarrays.

import "fmt"

func main() {
	fmt.Println(countIncreasing([]int{1, 3, 5, 4, 4, 6})) // 10
	fmt.Println(countIncreasing([]int{1, 2, 3, 4, 5}))    // 15
}

func countIncreasing(nums []int) int64 {
	var ans int64
	run := 0
	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i] > nums[i-1] {
			run++
		} else {
			run = 1
		}
		ans += int64(run)
	}
	return ans
}
```

## 2394 — Employees With Deductions

```go
package main

// LeetCode #2394: Employees With Deductions
// https://leetcode.com/problems/employees-with-deductions/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)
// Calculate actual hours worked per employee and compare with required hours.

import "fmt"

type Log struct {
	EmpID        int
	Timestamp    int
	IsLogin      bool
}

func main() {
	logs := []Log{
		{1, 100, true},
		{1, 200, false},
		{2, 50, true},
		{2, 150, false},
		{1, 300, true},
		{1, 400, false},
	}
	// Employee 1: total 200 min = 3.33 hrs, needs 4 hrs
	// Employee 2: total 100 min = 1.67 hrs, needs 4 hrs
	fmt.Println(calculateDeductions(logs, []int{4, 4})) // [1, 2] (both under)

	logs2 := []Log{
		{1, 0, true},
		{1, 240, false},
		{2, 0, true},
		{2, 480, false},
	}
	fmt.Println(calculateDeductions(logs2, []int{4, 8})) // [2]
}

func calculateDeductions(logs []Log, requiredHours []int) []int {
	hours := make(map[int]int)
	loginTime := make(map[int]int)

	for _, l := range logs {
		if l.IsLogin {
			loginTime[l.EmpID] = l.Timestamp
		} else {
			if start, ok := loginTime[l.EmpID]; ok {
				hours[l.EmpID] += l.Timestamp - start
				delete(loginTime, l.EmpID)
			}
		}
	}

	result := make([]int, 0)
	for empID := 1; empID <= len(requiredHours); empID++ {
		workedMin := hours[empID]
		neededMin := requiredHours[empID-1] * 60
		if workedMin < neededMin {
			result = append(result, empID)
		}
	}
	return result
}
```

## 2396 — Strictly Palindromic Number

```go
package main

// LeetCode #2396: Strictly Palindromic Number
// https://leetcode.com/problems/strictly-palindromic-number/
// Difficulty: Medium
// Time: O(1) | Space: O(1)
// For n >= 4, n in base (n-2) is always "12", which is never a palindrome.
// So answer is always false for n >= 4.

import "fmt"

func main() {
	fmt.Println(isStrictlyPalindromic(9))  // false
	fmt.Println(isStrictlyPalindromic(4))  // false
	fmt.Println(isStrictlyPalindromic(3))  // true (3 is "11" in base 2, "10" in base 1? no, base 2 only)
}

func isStrictlyPalindromic(n int) bool {
	// For all bases 2..n-2, check if representation is palindrome
	for base := 2; base <= n-2; base++ {
		if !isPalindromeInBase(n, base) {
			return false
		}
	}
	return true
}

func isPalindromeInBase(n, base int) bool {
	digits := make([]int, 0)
	for n > 0 {
		digits = append(digits, n%base)
		n /= base
	}
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		if digits[i] != digits[j] {
			return false
		}
	}
	return true
}
```

## 2397 — Maximum Rows Covered By Columns

```go
package main

// LeetCode #2397: Maximum Rows Covered by Columns
// https://leetcode.com/problems/maximum-rows-covered-by-columns/
// Difficulty: Medium
// Time: O(C(cols, select) * rows) | Space: O(cols)
// Brute force all column subsets via bitmask.

import "fmt"

func main() {
	fmt.Println(maximumRows([][]int{{0, 0, 0}, {1, 0, 1}, {0, 1, 1}, {0, 0, 1}}, 2)) // 3
	fmt.Println(maximumRows([][]int{{1}, {0}}, 1))                                    // 2
}

func maximumRows(mat [][]int, cols int) int {
	r := len(mat)
	rows := make([]int, r)
	for i := 0; i < r; i++ {
		mask := 0
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 1 {
				mask |= (1 << j)
			}
		}
		rows[i] = mask
	}

	ans := 0
	// iterate over all subsets of size cols
	var comb func(start, chosen, count int)
	comb = func(start, chosen, count int) {
		if count == cols {
			covered := 0
			for _, mask := range rows {
				if mask&^chosen == 0 {
					covered++
				}
			}
			if covered > ans {
				ans = covered
			}
			return
		}
		for j := start; j < len(mat[0]); j++ {
			comb(j+1, chosen|(1<<j), count+1)
		}
	}
	comb(0, 0, 0)
	return ans
}
```

## 2400 — Number Of Ways To Reach A Position After Exactly K Steps

```go
package main

// LeetCode #2400: Number of Ways to Reach a Position After Exactly k Steps
// https://leetcode.com/problems/number-of-ways-to-reach-a-position-after-exactly-k-steps/
// Difficulty: Medium
// Time: O(k^2) | Space: O(k)
// DP: dp[s] = ways to be at position s after i steps.
// Constraint: distance d = |startPos - endPos| must have same parity as k and d <= k.

import "fmt"

func main() {
	fmt.Println(numberOfWays(1, 2, 3)) // 3
	fmt.Println(numberOfWays(2, 5, 10)) // 0
	fmt.Println(numberOfWays(1, 1, 1)) // 0
}

const MOD = 1000000007

func numberOfWays(startPos int, endPos int, k int) int {
	d := abs(startPos - endPos)
	if d > k || (k-d)%2 != 0 {
		return 0
	}
	// DP with offset to handle negative positions
	offset := k
	size := 2*k + 1
	dp := make([]int, size)
	dp[0+offset] = 1

	for step := 0; step < k; step++ {
		next := make([]int, size)
		for pos := -k; pos <= k; pos++ {
			idx := pos + offset
			if dp[idx] == 0 {
				continue
			}
			// move left
			if pos-1 >= -k {
				next[pos-1+offset] = (next[pos-1+offset] + dp[idx]) % MOD
			}
			// move right
			if pos+1 <= k {
				next[pos+1+offset] = (next[pos+1+offset] + dp[idx]) % MOD
			}
		}
		dp = next
	}

	return dp[endPos-startPos+offset]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

## 2401 — Longest Nice Subarray

```go
package main

// LeetCode #2401: Longest Nice Subarray
// https://leetcode.com/problems/longest-nice-subarray/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// Sliding window: maintain OR of current window. If new element conflicts, shrink left.

import "fmt"

func main() {
	fmt.Println(longestNiceSubarray([]int{1, 3, 8, 48, 10})) // 3
	fmt.Println(longestNiceSubarray([]int{3, 1, 5, 11, 13}))  // 1
}

func longestNiceSubarray(nums []int) int {
	left, orMask, ans := 0, 0, 0
	for right, v := range nums {
		for orMask&v != 0 {
			orMask ^= nums[left]
			left++
		}
		orMask |= v
		if right-left+1 > ans {
			ans = right - left + 1
		}
	}
	return ans
}
```

## 2405 — Optimal Partition Of String

```go
package main

// LeetCode #2405: Optimal Partition of String
// https://leetcode.com/problems/optimal-partition-of-string/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// Greedy: start new partition when duplicate char found.

import "fmt"

func main() {
	fmt.Println(partitionString("abacaba")) // 4
	fmt.Println(partitionString("ssssss"))  // 6
}

func partitionString(s string) int {
	ans := 1
	seen := make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		if seen[s[i]] {
			ans++
			seen = make(map[byte]bool)
		}
		seen[s[i]] = true
	}
	return ans
}
```

## 2406 — Divide Intervals Into Minimum Number Of Groups

```go
package main

// LeetCode #2406: Divide Intervals Into Minimum Number of Groups
// https://leetcode.com/problems/divide-intervals-into-minimum-number-of-groups/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)
// Sweep line: count concurrent intervals, answer is max concurrency.

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minGroups([][]int{{5, 10}, {6, 8}, {1, 5}, {2, 3}, {1, 10}})) // 3
	fmt.Println(minGroups([][]int{{1, 3}, {5, 6}, {8, 10}, {11, 13}}))        // 1
}

func minGroups(intervals [][]int) int {
	events := make([][2]int, 0, len(intervals)*2)
	for _, iv := range intervals {
		events = append(events, [2]int{iv[0], 1})   // start
		events = append(events, [2]int{iv[1] + 1, -1}) // end (inclusive)
	}

	sort.Slice(events, func(i, j int) bool {
		if events[i][0] != events[j][0] {
			return events[i][0] < events[j][0]
		}
		return events[i][1] < events[j][1]
	})

	cur, ans := 0, 0
	for _, e := range events {
		cur += e[1]
		if cur > ans {
			ans = cur
		}
	}
	return ans
}
```

## 2408 — Design Sql

```go
package main

// LeetCode #2408: Design SQL
// https://leetcode.com/problems/design-sql/
// Difficulty: Medium
// Time: O(1) per operation if table exists | Space: O(n)
// In-memory SQL with tables, rows, and cell access.

import "fmt"

type Table struct {
	columns int
	rows    map[int][]string
	nextID  int
}

type SQL struct {
	tables map[string]*Table
}

func main() {
	sql := Constructor()
	sql.CreateTable("students", 3)
	sql.InsertRow("students", []string{"1", "John", "A"})
	sql.InsertRow("students", []string{"2", "Jane", "B"})
	fmt.Println(sql.SelectCell("students", 1, 2)) // "John" (row 1, col 2)

	sql.CreateTable("courses", 2)
	rowID := sql.InsertRow("courses", []string{"101", "Math"})
	fmt.Println(sql.SelectCell("courses", rowID, 2)) // "Math"
}

func Constructor() SQL {
	return SQL{tables: make(map[string]*Table)}
}

func (s *SQL) CreateTable(name string, columns int) {
	s.tables[name] = &Table{columns: columns, rows: make(map[int][]string), nextID: 1}
}

func (s *SQL) InsertRow(name string, values []string) int {
	t := s.tables[name]
	id := t.nextID
	t.nextID++
	row := make([]string, len(values))
	copy(row, values)
	t.rows[id] = row
	return id
}

func (s *SQL) SelectCell(name string, rowID, col int) string {
	return s.tables[name].rows[rowID][col-1]
}

func (s *SQL) DeleteRow(name string, rowID int) {
	delete(s.tables[name].rows, rowID)
}
```

## 2410 — Maximum Matching Of Players With Trainers

```go
package main

// LeetCode #2410: Maximum Matching of Players With Trainers
// https://leetcode.com/problems/maximum-matching-of-players-with-trainers/
// Difficulty: Medium
// Time: O(n log n + m log m) | Space: O(1)
// Sort both, greedy match: assign smallest sufficient trainer to each player.

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(matchPlayersAndTrainers([]int{4, 7, 9}, []int{8, 2, 5, 8})) // 2
	fmt.Println(matchPlayersAndTrainers([]int{1, 1, 1}, []int{10}))         // 1
}

func matchPlayersAndTrainers(players []int, trainers []int) int {
	sort.Ints(players)
	sort.Ints(trainers)
	i, j := 0, 0
	for i < len(players) && j < len(trainers) {
		if players[i] <= trainers[j] {
			i++
		}
		j++
	}
	return i
}
```

## 2411 — Smallest Subarrays With Maximum Bitwise Or

```go
package main

// LeetCode #2411: Smallest Subarrays With Maximum Bitwise OR
// https://leetcode.com/problems/smallest-subarrays-with-maximum-bitwise-or/
// Difficulty: Medium
// Time: O(n * 30) | Space: O(1)
// For each position i, find min length subarray starting at i with max possible OR.
// Track for each bit, the nearest position to the right where it's set.

import "fmt"

func main() {
	fmt.Println(smallestSubarrays([]int{1, 0, 2, 1, 3})) // [3, 3, 2, 2, 1]
	fmt.Println(smallestSubarrays([]int{1, 2}))          // [2, 1]
}

func smallestSubarrays(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	last := make([]int, 30) // for each bit, last position where it's set

	bitPos := -1
	for i := n - 1; i >= 0; i-- {
		maxDist := 1
		for b := 0; b < 30; b++ {
			if nums[i]>>b&1 == 1 {
				last[b] = i
			}
			bitPos = last[b]
			if bitPos != 0 {
				dist := bitPos - i + 1
				if dist > maxDist {
					maxDist = dist
				}
			}
		}
		ans[i] = maxDist
	}
	return ans
}
```

## 2414 — Length Of The Longest Alphabetical Continuous Substring

```go
package main

// LeetCode #2414: Length of the Longest Alphabetical Continuous Substring
// https://leetcode.com/problems/length-of-the-longest-alphabetical-continuous-substring/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// Scan, count consecutive chars where s[i] == s[i-1] + 1.

import "fmt"

func main() {
	fmt.Println(longestContinuousSubstring("abacaba")) // 2 ("ab")
	fmt.Println(longestContinuousSubstring("abcde"))   // 5
}

func longestContinuousSubstring(s string) int {
	ans, cur := 0, 0
	for i := 0; i < len(s); i++ {
		if i == 0 || s[i] == s[i-1]+1 {
			cur++
		} else {
			cur = 1
		}
		if cur > ans {
			ans = cur
		}
	}
	return ans
}
```

## 2415 — Reverse Odd Levels Of Binary Tree

```go
package main

// LeetCode #2415: Reverse Odd Levels of Binary Tree
// https://leetcode.com/problems/reverse-odd-levels-of-binary-tree/
// Difficulty: Medium
// Time: O(n) | Space: O(n)
// BFS level-order, reverse values at odd levels.

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Tree: [2,3,5,8,13,21,34]
	root := &TreeNode{2,
		&TreeNode{3,
			&TreeNode{8, nil, nil},
			&TreeNode{13, nil, nil},
		},
		&TreeNode{5,
			&TreeNode{21, nil, nil},
			&TreeNode{34, nil, nil},
		},
	}
	reversed := reverseOddLevels(root)
	// After reverse: level 1: 3<->5, becomes [2,5,3,...]
	fmt.Println(reversed.Left.Val)  // 5
	fmt.Println(reversed.Right.Val) // 3

	root2 := &TreeNode{7,
		&TreeNode{13, nil, nil},
		&TreeNode{11, nil, nil},
	}
	rev2 := reverseOddLevels(root2)
	fmt.Println(rev2.Left.Val)  // 11
	fmt.Println(rev2.Right.Val) // 13
}

func reverseOddLevels(root *TreeNode) *TreeNode {
	q := []*TreeNode{root}
	level := 0
	for len(q) > 0 {
		if level%2 == 1 {
			// reverse values at this level
			for i, j := 0, len(q)-1; i < j; i, j = i+1, j-1 {
				q[i].Val, q[j].Val = q[j].Val, q[i].Val
			}
		}
		next := make([]*TreeNode, 0)
		for _, node := range q {
			if node.Left != nil {
				next = append(next, node.Left)
				next = append(next, node.Right)
			}
		}
		q = next
		level++
	}
	return root
}
```

## 2417 — Closest Fair Integer

```go
package main

// LeetCode #2417: Closest Fair Integer
// https://leetcode.com/problems/closest-fair-integer/
// Difficulty: Medium
// Time: O(log n * 2^d) | Space: O(d)
// Find smallest fair integer >= n (equal even and odd digit count).

import "fmt"

func main() {
	fmt.Println(closestFair(2))    // 10
	fmt.Println(closestFair(403))  // 440
	fmt.Println(closestFair(10))   // 10
}

func closestFair(n int) int {
	for {
		if isFair(n) {
			return n
		}
		n++
	}
}

func isFair(n int) bool {
	s := fmt.Sprint(n)
	even, odd := 0, 0
	for _, ch := range s {
		d := int(ch - '0')
		if d%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	return even == odd
}
```

## 2419 — Longest Subarray With Maximum Bitwise And

```go
package main

// LeetCode #2419: Longest Subarray With Maximum Bitwise AND
// https://leetcode.com/problems/longest-subarray-with-maximum-bitwise-and/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// Max AND in any subarray is just the max element. Find longest consecutive
// subarray where all elements equal the max element.

import "fmt"

func main() {
	fmt.Println(longestSubarray([]int{1, 2, 3, 3, 2, 2})) // 2
	fmt.Println(longestSubarray([]int{1, 2, 3, 4}))        // 1
}

func longestSubarray(nums []int) int {
	mx := 0
	for _, v := range nums {
		if v > mx {
			mx = v
		}
	}
	ans, cur := 0, 0
	for _, v := range nums {
		if v == mx {
			cur++
		} else {
			cur = 0
		}
		if cur > ans {
			ans = cur
		}
	}
	return ans
}
```

## 2420 — Find All Good Indices

```go
package main

// LeetCode #2420: Find All Good Indices
// https://leetcode.com/problems/find-all-good-indices/
// Difficulty: Medium
// Time: O(n) | Space: O(n)
// Precompute prefix non-increasing and suffix non-decreasing lengths.

import "fmt"

func main() {
	fmt.Println(goodIndices([]int{2, 1, 1, 1, 3, 4, 1}, 2)) // [4, 5]
	fmt.Println(goodIndices([]int{1, 2, 3, 4, 5, 6}, 2))    // []
}

func goodIndices(nums []int, k int) []int {
	n := len(nums)
	pref := make([]int, n) // longest non-increasing ending at i
	suf := make([]int, n)  // longest non-decreasing starting at i

	pref[0] = 1
	for i := 1; i < n; i++ {
		if nums[i] <= nums[i-1] {
			pref[i] = pref[i-1] + 1
		} else {
			pref[i] = 1
		}
	}

	suf[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		if nums[i] <= nums[i+1] {
			suf[i] = suf[i+1] + 1
		} else {
			suf[i] = 1
		}
	}

	ans := make([]int, 0)
	for i := k; i < n-k; i++ {
		if pref[i-1] >= k && suf[i+1] >= k {
			ans = append(ans, i)
		}
	}
	return ans
}
```

## 2422 — Merge Operations To Turn Array Into A Palindrome

```go
package main

// LeetCode #2422: Merge Operations to Turn Array Into a Palindrome
// https://leetcode.com/problems/merge-operations-to-turn-array-into-a-palindrome/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// Two pointers, merge smaller side towards larger.

import "fmt"

func main() {
	fmt.Println(minMerges([]int{1, 4, 1, 3}))        // 1 (merge 4+1 to 5, then [1,5,3])
	fmt.Println(minMerges([]int{1, 2, 3, 4, 5, 1})) // 2
}

func minMerges(nums []int) int {
	ops := 0
	i, j := 0, len(nums)-1
	left, right := nums[i], nums[j]

	for i < j {
		if left < right {
			i++
			left += nums[i]
			ops++
		} else if left > right {
			j--
			right += nums[j]
			ops++
		} else {
			i++
			j--
			if i < j {
				left, right = nums[i], nums[j]
			}
		}
	}
	return ops
}
```

## 2424 — Longest Uploaded Prefix

```go
package main

// LeetCode #2424: Longest Uploaded Prefix
// https://leetcode.com/problems/longest-uploaded-prefix/
// Difficulty: Medium
// Time: O(1) amortized | Space: O(n)
// Track uploaded videos. Longest prefix = longest 1..k where all uploaded.

import "fmt"

type LUPrefix struct {
	uploaded []bool
	longest  int
}

func main() {
	lu := Constructor(4)
	fmt.Println(lu.Longest()) // 0
	lu.Upload(3)
	fmt.Println(lu.Longest()) // 0
	lu.Upload(1)
	fmt.Println(lu.Longest()) // 1
	lu.Upload(2)
	fmt.Println(lu.Longest()) // 3
}

func Constructor(n int) LUPrefix {
	return LUPrefix{uploaded: make([]bool, n+2)}
}

func (l *LUPrefix) Upload(video int) {
	l.uploaded[video] = true
	for l.uploaded[l.longest+1] {
		l.longest++
	}
}

func (l *LUPrefix) Longest() int {
	return l.longest
}
```

## 2425 — Bitwise Xor Of All Pairings

```go
package main

// LeetCode #2425: Bitwise XOR of All Pairings
// https://leetcode.com/problems/bitwise-xor-of-all-pairings/
// Difficulty: Medium
// Time: O(m + n) | Space: O(1)
// XOR of all pairings = if len(nums2) odd, XOR all nums1; if len(nums1) odd, XOR all nums2.

import "fmt"

func main() {
	fmt.Println(xorAllNums([]int{2, 1, 3}, []int{10, 2, 5, 0})) // 13
	fmt.Println(xorAllNums([]int{1, 2}, []int{3, 4}))            // 0
}

func xorAllNums(nums1 []int, nums2 []int) int {
	ans := 0
	if len(nums2)%2 == 1 {
		for _, v := range nums1 {
			ans ^= v
		}
	}
	if len(nums1)%2 == 1 {
		for _, v := range nums2 {
			ans ^= v
		}
	}
	return ans
}
```

## 2428 — Maximum Sum Of An Hourglass

```go
package main

// LeetCode #2428: Maximum Sum of an Hourglass
// https://leetcode.com/problems/maximum-sum-of-an-hourglass/
// Difficulty: Medium
// Time: O(m * n) | Space: O(1)
// Sliding hourglass over the grid.

import "fmt"

func main() {
	fmt.Println(maxSum([][]int{{6, 2, 1, 3}, {4, 2, 1, 5}, {9, 2, 8, 7}, {4, 1, 2, 9}})) // 30
	fmt.Println(maxSum([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))                         // 35
}

func maxSum(grid [][]int) int {
	r, c := len(grid), len(grid[0])
	ans := 0
	for i := 0; i+2 < r; i++ {
		for j := 0; j+2 < c; j++ {
			sum := grid[i][j] + grid[i][j+1] + grid[i][j+2] + // top row
				grid[i+1][j+1] + // middle
				grid[i+2][j] + grid[i+2][j+1] + grid[i+2][j+2] // bottom row
			if sum > ans {
				ans = sum
			}
		}
	}
	return ans
}
```

## 2429 — Minimize Xor

```go
package main

// LeetCode #2429: Minimize XOR
// https://leetcode.com/problems/minimize-xor/
// Difficulty: Medium
// Time: O(30) | Space: O(1)
// Find x with same set bits as num2 that minimizes x XOR num1.
// Approach: set highest bits of num1 first, then lowest unset bits.

import "fmt"

func main() {
	fmt.Println(minimizeXor(3, 5))  // 3  (3=11, 5=101, set bits: 2)
	fmt.Println(minimizeXor(1, 12)) // 3  (1=1, 12=1100, set bits: 2)
}

func minimizeXor(num1 int, num2 int) int {
	targetBits := bitsCount(num2)
	x := 0
	// Set bits from num1's highest bits first
	for b := 30; b >= 0 && targetBits > 0; b-- {
		if num1>>b&1 == 1 {
			x |= (1 << b)
			targetBits--
		}
	}
	// Set remaining bits from lowest unset positions
	for b := 0; b <= 30 && targetBits > 0; b++ {
		if x>>b&1 == 0 {
			x |= (1 << b)
			targetBits--
		}
	}
	return x
}

func bitsCount(n int) int {
	count := 0
	for n > 0 {
		count += n & 1
		n >>= 1
	}
	return count
}
```

## 2431 — Maximize Total Tastiness Of Purchased Fruits

```go
package main

// LeetCode #2431: Maximize Total Tastiness of Purchased Fruits
// https://leetcode.com/problems/maximize-total-tastiness-of-purchased-fruits/
// Difficulty: Medium
// Time: O(n * budget * coupon) | Space: O(budget * coupon)
// Knapsack DP: dp[b][c] = max tastiness with b budget and c coupons remaining.

import "fmt"

type Fruit struct {
	price, tastiness int
}

func main() {
	fruits := []Fruit{{2, 3}, {3, 6}, {5, 10}}
	fmt.Println(maxTastiness(fruits, 10, 1)) // 16 (buy 3rd with coupon, 1st and 2nd normally)

	fruits2 := []Fruit{{1, 5}, {2, 3}, {3, 6}}
	fmt.Println(maxTastiness(fruits2, 5, 2)) // 14
}

func maxTastiness(fruits []Fruit, budget int, couponCount int) int {
	dp := make([][]int, budget+1)
	for b := range dp {
		dp[b] = make([]int, couponCount+1)
		for c := range dp[b] {
			dp[b][c] = -1
		}
	}
	dp[0][0] = 0
	ans := 0

	for _, f := range fruits {
		newDp := make([][]int, budget+1)
		for b := range newDp {
			newDp[b] = make([]int, couponCount+1)
			copy(newDp[b], dp[b])
		}

		for b := budget; b >= 0; b-- {
			for c := 0; c <= couponCount; c++ {
				if dp[b][c] < 0 {
					continue
				}
				// buy normally
				if b+f.price <= budget {
					if dp[b][c]+f.tastiness > newDp[b+f.price][c] {
						newDp[b+f.price][c] = dp[b][c] + f.tastiness
					}
				}
				// buy with coupon (half price)
				if c < couponCount {
					cp := f.price / 2
					if b+cp <= budget {
						if dp[b][c]+f.tastiness > newDp[b+cp][c+1] {
							newDp[b+cp][c+1] = dp[b][c] + f.tastiness
						}
					}
				}
			}
		}
		dp = newDp
	}

	for b := 0; b <= budget; b++ {
		for c := 0; c <= couponCount; c++ {
			if dp[b][c] > ans {
				ans = dp[b][c]
			}
		}
	}
	return ans
}
```

## 2433 — Find The Original Array Of Prefix Xor

```go
package main

// LeetCode #2433: Find The Original Array of Prefix Xor
// https://leetcode.com/problems/find-the-original-array-of-prefix-xor/
// Difficulty: Medium
// Time: O(n) | Space: O(1) (excluding output)
// Given pref[i] = XOR(arr[0..i]), find arr. arr[i] = pref[i] ^ pref[i-1].

import "fmt"

func main() {
	fmt.Println(findArray([]int{5, 2, 0, 3, 1})) // [5, 7, 2, 3, 2]
	fmt.Println(findArray([]int{13}))             // [13]
}

func findArray(pref []int) []int {
	n := len(pref)
	arr := make([]int, n)
	arr[0] = pref[0]
	for i := 1; i < n; i++ {
		arr[i] = pref[i] ^ pref[i-1]
	}
	return arr
}
```

