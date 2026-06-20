# Medium (Sedang) — Problem 1079–1262

## 1079 — Letter Tile Possibilities

```go
package main

// LeetCode #1079: Letter Tile Possibilities
// https://leetcode.com/problems/letter-tile-possibilities/
// Difficulty: Medium
//
// Approach: Backtracking with frequency count
// Time: O(n!) where n = len(tiles) worst case
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(numTilePossibilities("AAB")) // 8
	fmt.Println(numTilePossibilities("AAABBC")) // 188
}

func numTilePossibilities(tiles string) int {
	freq := make([]int, 26)
	for _, c := range tiles {
		freq[c-'A']++
	}

	var dfs func() int
	dfs = func() int {
		count := 0
		for i := 0; i < 26; i++ {
			if freq[i] == 0 {
				continue
			}
			count++
			freq[i]--
			count += dfs()
			freq[i]++
		}
		return count
	}

	return dfs()
}
```

## 1080 — Insufficient Nodes In Root To Leaf Paths

```go
package main

// LeetCode #1080: Insufficient Nodes in Root to Leaf Paths
// https://leetcode.com/problems/insufficient-nodes-in-root-to-leaf-paths/
// Difficulty: Medium
//
// Approach: DFS post-order. Delete node if sum from root to leaf < limit.
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
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4, Left: &TreeNode{Val: 8, Left: nil, Right: nil}, Right: nil},
			Right: &TreeNode{Val: 5, Left: nil, Right: nil},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 6, Left: nil, Right: nil},
			Right: &TreeNode{Val: 7, Left: nil, Right: nil},
		},
	}
	result := sufficientSubset(root, 10)
	printTree(result)
	fmt.Println()
}

func sufficientSubset(root *TreeNode, limit int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		if root.Val < limit {
			return nil
		}
		return root
	}

	root.Left = sufficientSubset(root.Left, limit-root.Val)
	root.Right = sufficientSubset(root.Right, limit-root.Val)

	if root.Left == nil && root.Right == nil {
		return nil
	}
	return root
}

func printTree(root *TreeNode) {
	if root == nil {
		fmt.Print("null")
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

## 1081 — Smallest Subsequence Of Distinct Characters

```go
package main

// LeetCode #1081: Smallest Subsequence of Distinct Characters
// https://leetcode.com/problems/smallest-subsequence-of-distinct-characters/
// Difficulty: Medium
//
// Approach: Monotonic stack (greedy). Track last occurrence and used set.
// Time: O(n)
// Space: O(26) = O(1)

import "fmt"

func main() {
	fmt.Println(smallestSubsequence("bcabc"))  // "abc"
	fmt.Println(smallestSubsequence("cbacdcbc")) // "acdb"
}

func smallestSubsequence(s string) string {
	lastOccur := make([]int, 26)
	for i := 0; i < len(s); i++ {
		lastOccur[s[i]-'a'] = i
	}

	used := make([]bool, 26)
	stack := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		ch := s[i]
		if used[ch-'a'] {
			continue
		}

		for len(stack) > 0 && stack[len(stack)-1] > ch && lastOccur[stack[len(stack)-1]-'a'] > i {
			used[stack[len(stack)-1]-'a'] = false
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, ch)
		used[ch-'a'] = true
	}

	return string(stack)
}
```

## 1087 — Brace Expansion

```go
package main

// LeetCode #1087: Brace Expansion
// https://leetcode.com/problems/brace-expansion/
// Difficulty: Medium
//
// Approach: Backtracking - parse braces and generate all expansions
// Time: O(n * k) where k is number of expansions
// Space: O(n * k)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(expand("{a,b}c{d,e}f")) // ["acdf","acef","bcdf","bcef"]
	fmt.Println(expand("abcd"))          // ["abcd"]
}

func expand(s string) []string {
	result := make([]string, 0)
	backtrack(s, 0, "", &result)
	sort.Strings(result)
	return result
}

func backtrack(s string, idx int, cur string, result *[]string) {
	if idx == len(s) {
		*result = append(*result, cur)
		return
	}

	if s[idx] == '{' {
		// Find the closing brace
		end := idx + 1
		for s[end] != '}' {
			end++
		}
		// Parse options
		options := make([]byte, 0)
		for k := idx + 1; k < end; k++ {
			if s[k] != ',' {
				options = append(options, s[k])
			}
		}
		sort.Slice(options, func(i, j int) bool {
			return options[i] < options[j]
		})
		for _, opt := range options {
			backtrack(s, end+1, cur+string(opt), result)
		}
	} else {
		backtrack(s, idx+1, cur+string(s[idx]), result)
	}
}
```

## 1090 — Largest Values From Labels

```go
package main

// LeetCode #1090: Largest Values From Labels
// https://leetcode.com/problems/largest-values-from-labels/
// Difficulty: Medium
//
// Approach: Sort by value descending, pick items with label limits
// Time: O(n log n)
// Space: O(n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(largestValsFromLabels([]int{5, 4, 3, 2, 1}, []int{1, 1, 2, 2, 3}, 3, 1)) // 9
	fmt.Println(largestValsFromLabels([]int{5, 4, 3, 2, 1}, []int{1, 3, 3, 3, 2}, 3, 2)) // 12
}

func largestValsFromLabels(values []int, labels []int, numWanted int, useLimit int) int {
	n := len(values)
	type pair struct {
		val   int
		label int
	}
	items := make([]pair, n)
	for i := 0; i < n; i++ {
		items[i] = pair{values[i], labels[i]}
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].val > items[j].val
	})

	labelCount := make(map[int]int)
	result := 0
	selected := 0

	for _, item := range items {
		if selected >= numWanted {
			break
		}
		if labelCount[item.label] < useLimit {
			result += item.val
			labelCount[item.label]++
			selected++
		}
	}

	return result
}
```

## 1091 — Shortest Path In Binary Matrix

```go
package main

// LeetCode #1091: Shortest Path in Binary Matrix
// https://leetcode.com/problems/shortest-path-in-binary-matrix/
// Difficulty: Medium
//
// Approach: BFS from (0,0) to (n-1,n-1) with 8-directional moves
// Time: O(n^2)
// Space: O(n^2)

import "fmt"

func main() {
	fmt.Println(shortestPathBinaryMatrix([][]int{{0, 1}, {1, 0}})) // 2
	fmt.Println(shortestPathBinaryMatrix([][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}})) // 4
}

func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return -1
	}
	if n == 1 {
		return 1
	}

	dirs := [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	queue := make([][2]int, 0, n*n)
	queue = append(queue, [2]int{0, 0})
	grid[0][0] = 1
	distance := 1

	for len(queue) > 0 {
		size := len(queue)
		distance++
		for k := 0; k < size; k++ {
			cur := queue[k]
			for _, d := range dirs {
				ni, nj := cur[0]+d[0], cur[1]+d[1]
				if ni >= 0 && ni < n && nj >= 0 && nj < n && grid[ni][nj] == 0 {
					if ni == n-1 && nj == n-1 {
						return distance
					}
					grid[ni][nj] = 1
					queue = append(queue, [2]int{ni, nj})
				}
			}
		}
		queue = queue[size:]
	}

	return -1
}
```

## 1093 — Statistics From A Large Sample

```go
package main

// LeetCode #1093: Statistics from a Large Sample
// https://leetcode.com/problems/statistics-from-a-large-sample/
// Difficulty: Medium
//
// Approach: Single pass to compute min, max, sum, mode.
//           Two-pointer for median.
// Time: O(n) where n = len(count) = 256
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(sampleStats([]int{0, 1, 3, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}))
	// Expected: [1.00000,3.00000,2.37500,2.50000,3.00000]
}

func sampleStats(count []int) []float64 {
	n := 0
	sum := 0
	minVal := -1
	maxVal := 0
	modeVal := 0
	modeCount := 0

	for i, c := range count {
		if c > 0 {
			n += c
			sum += i * c
			if minVal == -1 {
				minVal = i
			}
			maxVal = i
			if c > modeCount {
				modeCount = c
				modeVal = i
			}
		}
	}

	mean := float64(sum) / float64(n)

	// Median
	median := 0.0
	left := n / 2
	right := left + 1
	if n%2 == 1 {
		right = left
	}

	leftVal, rightVal := 0, 0
	acc := 0
	for i, c := range count {
		if c > 0 {
			if acc < left && acc+c >= left {
				leftVal = i
			}
			if acc < right && acc+c >= right {
				rightVal = i
			}
			acc += c
		}
	}
	median = float64(leftVal+rightVal) / 2.0

	return []float64{float64(minVal), float64(maxVal), mean, median, float64(modeVal)}
}
```

## 1094 — Car Pooling

```go
package main

// LeetCode #1094: Car Pooling
// https://leetcode.com/problems/car-pooling/
// Difficulty: Medium
//
// Approach: Difference array (prefix sum)
// Time: O(n + maxLocation)
// Space: O(maxLocation)

import "fmt"

func main() {
	fmt.Println(carPooling([][]int{{2, 1, 5}, {3, 3, 7}}, 4)) // false
	fmt.Println(carPooling([][]int{{2, 1, 5}, {3, 3, 7}}, 5)) // true
}

func carPooling(trips [][]int, capacity int) bool {
	maxLoc := 0
	for _, t := range trips {
		if t[2] > maxLoc {
			maxLoc = t[2]
		}
	}

	diff := make([]int, maxLoc+2)
	for _, t := range trips {
		diff[t[1]] += t[0]
		diff[t[2]] -= t[0]
	}

	current := 0
	for _, d := range diff {
		current += d
		if current > capacity {
			return false
		}
	}

	return true
}
```

## 1098 — Unpopular Books

```go
package main

// LeetCode #1098: Unpopular Books
// https://leetcode.com/problems/unpopular-books/
// Difficulty: Medium
//
// Approach: Filter books ordered less than 10 times in the last year
// Time: O(n + m) where n = books, m = orders
// Space: O(n)

import "fmt"

func main() {
	// Books: (book_id, name)
	books := []string{"Book A", "Book B", "Book C", "Book D"}
	// Orders: (book_name, quantity, days_ago)
	orders := []struct {
		name     string
		quantity int
		daysAgo  int
	}{
		{"Book A", 5, 30},
		{"Book B", 15, 10},
		{"Book C", 8, 20},
	}
	fmt.Println(unpopularBooks(books, orders))
}

func unpopularBooks(books []string, orders []struct {
	name     string
	quantity int
	daysAgo  int
}) []string {
	orderCount := make(map[string]int)
	for _, o := range orders {
		if o.daysAgo <= 365 {
			orderCount[o.name] += o.quantity
		}
	}

	result := make([]string, 0)
	for _, b := range books {
		count, exists := orderCount[b]
		if !exists || count < 10 {
			result = append(result, b)
		}
	}

	return result
}
```

## 1100 — Find K Length Substrings With No Repeated Characters

```go
package main

// LeetCode #1100: Find K-Length Substrings With No Repeated Characters
// https://leetcode.com/problems/find-k-length-substrings-with-no-repeated-characters/
// Difficulty: Medium
//
// Approach: Sliding window with frequency array
// Time: O(n)
// Space: O(26) = O(1)

import "fmt"

func main() {
	fmt.Println(numKLenSubstrNoRepeats("havefunonleetcode", 5)) // 6
	fmt.Println(numKLenSubstrNoRepeats("home", 5))              // 0
}

func numKLenSubstrNoRepeats(s string, k int) int {
	if k > len(s) || k > 26 {
		return 0
	}

	freq := make([]int, 26)
	duplicates := 0
	result := 0

	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
		if freq[s[i]-'a'] == 2 {
			duplicates++
		}

		if i >= k {
			freq[s[i-k]-'a']--
			if freq[s[i-k]-'a'] == 1 {
				duplicates--
			}
		}

		if i >= k-1 && duplicates == 0 {
			result++
		}
	}

	return result
}
```

## 1101 — The Earliest Moment When Everyone Become Friends

```go
package main

// LeetCode #1101: The Earliest Moment When Everyone Become Friends
// https://leetcode.com/problems/the-earliest-moment-when-everyone-become-friends/
// Difficulty: Medium
//
// Approach: Sort logs by timestamp, Union-Find
// Time: O(n log n + m * alpha(n)) where n = logs, m = N
// Space: O(N)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(earliestAcq([][]int{{20190101, 0, 1}, {20190104, 3, 4}, {20190107, 2, 3}, {20190211, 1, 5}, {20190224, 2, 4}, {20190301, 0, 3}, {20190312, 1, 2}, {20190322, 4, 5}}, 6)) // 20190301
	fmt.Println(earliestAcq([][]int{{0, 0, 1}, {1, 1, 2}, {2, 0, 2}}, 3)) // 2
}

func earliestAcq(logs [][]int, n int) int {
	sort.Slice(logs, func(i, j int) bool {
		return logs[i][0] < logs[j][0]
	})

	parent := make([]int, n)
	for i := 0; i < n; i++ {
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

	groups := n
	for _, log := range logs {
		ts, a, b := log[0], log[1], log[2]
		if find(a) != find(b) {
			union(a, b)
			groups--
			if groups == 1 {
				return ts
			}
		}
	}

	return -1
}
```

## 1102 — Path With Maximum Minimum Value

```go
package main

// LeetCode #1102: Path With Maximum Minimum Value
// https://leetcode.com/problems/path-with-maximum-minimum-value/
// Difficulty: Medium
//
// Approach: BFS with max-heap (priority queue). Always visit cell with largest value.
// Time: O(m * n * log(m * n))
// Space: O(m * n)

import "fmt"

func main() {
	fmt.Println(maximumMinimumPath([][]int{{5, 4, 5}, {1, 2, 6}, {7, 4, 6}})) // 4
	fmt.Println(maximumMinimumPath([][]int{{2, 0, 1}, {4, 3, 2}, {5, 4, 5}})) // 3
}

type cell struct {
	i, j, val int
}

func maximumMinimumPath(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	// Max-heap using slice
	heap := []cell{{0, 0, grid[0][0]}}
	visited[0][0] = true
	result := grid[0][0]

	for len(heap) > 0 {
		// Find max value in heap (simple linear scan for small heap)
		maxIdx := 0
		for k := 1; k < len(heap); k++ {
			if heap[k].val > heap[maxIdx].val {
				maxIdx = k
			}
		}
		cur := heap[maxIdx]
		heap[maxIdx] = heap[len(heap)-1]
		heap = heap[:len(heap)-1]

		if cur.val < result {
			result = cur.val
		}

		if cur.i == m-1 && cur.j == n-1 {
			return result
		}

		for _, d := range dirs {
			ni, nj := cur.i+d[0], cur.j+d[1]
			if ni >= 0 && ni < m && nj >= 0 && nj < n && !visited[ni][nj] {
				visited[ni][nj] = true
				heap = append(heap, cell{ni, nj, grid[ni][nj]})
			}
		}
	}

	return result
}
```

## 1104 — Path In Zigzag Labelled Binary Tree

```go
package main

// LeetCode #1104: Path In Zigzag Labelled Binary Tree
// https://leetcode.com/problems/path-in-zigzag-labelled-binary-tree/
// Difficulty: Medium
//
// Approach: Find level, compute reverse label, traverse to root
// Time: O(log n)
// Space: O(log n)

import "fmt"

func main() {
	fmt.Println(pathInZigZagTree(14)) // [1,3,4,14]
	fmt.Println(pathInZigZagTree(26)) // [1,2,6,10,26]
}

func pathInZigZagTree(label int) []int {
	result := make([]int, 0)

	for label > 0 {
		result = append(result, label)
		level := 0
		for (1 << level) <= label {
			level++
		}
		level--

		// In zigzag levels, the "position" is reversed
		// Min and max of this level
		minVal := 1 << level
		maxVal := (1 << (level + 1)) - 1
		// The parent of label (in zigzag) needs to find the "mirror" position
		parent := minVal + maxVal - label
		label = parent / 2
	}

	// Reverse to get root-to-leaf
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}
```

## 1105 — Filling Bookcase Shelves

```go
package main

// LeetCode #1105: Filling Bookcase Shelves
// https://leetcode.com/problems/filling-bookcase-shelves/
// Difficulty: Medium
//
// Approach: DP. dp[i] = min height to place first i books.
//           Try placing books i-1..j on the same shelf.
// Time: O(n^2)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(minHeightShelves([][]int{{1, 1}, {2, 3}, {2, 3}, {1, 1}, {1, 1}, {1, 1}, {1, 2}}, 4)) // 6
	fmt.Println(minHeightShelves([][]int{{1, 3}, {2, 4}, {3, 2}}, 6))                               // 4
}

func minHeightShelves(books [][]int, shelfWidth int) int {
	n := len(books)
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = 1<<31 - 1
	}

	for i := 1; i <= n; i++ {
		width := 0
		height := 0
		for j := i; j > 0; j-- {
			width += books[j-1][0]
			if width > shelfWidth {
				break
			}
			if books[j-1][1] > height {
				height = books[j-1][1]
			}
			if dp[j-1]+height < dp[i] {
				dp[i] = dp[j-1] + height
			}
		}
	}

	return dp[n]
}
```

## 1107 — New Users Daily Count

```go
package main

// LeetCode #1107: New Users Daily Count
// https://leetcode.com/problems/new-users-daily-count/
// Difficulty: Medium
//
// Approach: Track first login date per user, count by date
// Time: O(n) where n = len(traffic)
// Space: O(m) where m = unique users

import "fmt"

func main() {
	// traffic: (user_id, activity date, is_login)
	traffic := [][3]int{
		{1, 1, 1},  // user 1 login on day 1
		{2, 1, 1},  // user 2 login on day 1
		{3, 2, 1},  // user 3 login on day 2
		{1, 3, 0},  // user 1 logout on day 3
		{2, 3, 0},  // user 2 logout on day 3
		{4, 3, 1},  // user 4 login on day 3
	}
	fmt.Println(newUsersDailyCount(traffic))
}

func newUsersDailyCount(traffic [][3]int) map[int]int {
	firstLogin := make(map[int]int) // userID -> first login date
	for _, t := range traffic {
		userID, date, isLogin := t[0], t[1], t[2]
		if isLogin == 1 {
			if _, exists := firstLogin[userID]; !exists || date < firstLogin[userID] {
				firstLogin[userID] = date
			}
		}
	}

	result := make(map[int]int)
	for _, date := range firstLogin {
		result[date]++
	}

	return result
}
```

## 1109 — Corporate Flight Bookings

```go
package main

// LeetCode #1109: Corporate Flight Bookings
// https://leetcode.com/problems/corporate-flight-bookings/
// Difficulty: Medium
//
// Approach: Difference array (prefix sum)
// Time: O(n + len(bookings))
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(corpFlightBookings([][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}}, 5)) // [10,55,45,25,25]
	fmt.Println(corpFlightBookings([][]int{{1, 1, 5}, {2, 2, 10}}, 2))             // [5,10]
}

func corpFlightBookings(bookings [][]int, n int) []int {
	result := make([]int, n+2)
	for _, b := range bookings {
		first, last, seats := b[0], b[1], b[2]
		result[first] += seats
		result[last+1] -= seats
	}

	for i := 1; i <= n; i++ {
		result[i] += result[i-1]
	}

	return result[1 : n+1]
}
```

## 1110 — Delete Nodes And Return Forest

```go
package main

// LeetCode #1110: Delete Nodes And Return Forest
// https://leetcode.com/problems/delete-nodes-and-return-forest/
// Difficulty: Medium
//
// Approach: DFS post-order. If node should be deleted, add children to forest.
// Time: O(n)
// Space: O(n + h) where h is tree height

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4, Left: nil, Right: nil},
			Right: &TreeNode{Val: 5, Left: nil, Right: nil},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 6, Left: nil, Right: nil},
			Right: &TreeNode{Val: 7, Left: nil, Right: nil},
		},
	}
	result := delNodes(root, []int{3, 5})
	for _, tree := range result {
		if tree != nil {
			fmt.Printf("%d ", tree.Val)
		}
	}
	fmt.Println()
}

func delNodes(root *TreeNode, toDelete []int) []*TreeNode {
	deleteSet := make(map[int]bool)
	for _, v := range toDelete {
		deleteSet[v] = true
	}

	result := make([]*TreeNode, 0)
	if !deleteSet[root.Val] {
		result = append(result, root)
	}

	var dfs func(node *TreeNode) *TreeNode
	dfs = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}

		node.Left = dfs(node.Left)
		node.Right = dfs(node.Right)

		if deleteSet[node.Val] {
			if node.Left != nil {
				result = append(result, node.Left)
			}
			if node.Right != nil {
				result = append(result, node.Right)
			}
			return nil
		}
		return node
	}

	dfs(root)
	return result
}
```

## 1111 — Maximum Nesting Depth Of Two Valid Parentheses Strings

```go
package main

// LeetCode #1111: Maximum Nesting Depth of Two Valid Parentheses Strings
// https://leetcode.com/problems/maximum-nesting-depth-of-two-valid-parentheses-strings/
// Difficulty: Medium
//
// Approach: Assign '(' to group A or B based on even/odd depth.
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(maxDepthAfterSplit("(()())")) // [0,1,1,1,1,0] or similar
	fmt.Println(maxDepthAfterSplit("()(())()")) // [0,0,0,1,1,0,0,0]
}

func maxDepthAfterSplit(seq string) []int {
	result := make([]int, len(seq))
	depth := 0

	for i, c := range seq {
		if c == '(' {
			depth++
			result[i] = depth % 2
		} else {
			result[i] = depth % 2
			depth--
		}
	}

	return result
}
```

## 1112 — Highest Grade For Each Student

```go
package main

// LeetCode #1112: Highest Grade For Each Student
// https://leetcode.com/problems/highest-grade-for-each-student/
// Difficulty: Medium
//
// Approach: Track best grade (highest, then earliest course_id) per student
// Time: O(n)
// Space: O(m) where m = unique students

import "fmt"

func main() {
	// (student_id, course_id, grade)
	enrollments := [][]int{{1, 1, 90}, {1, 2, 95}, {2, 1, 85}, {2, 2, 85}, {3, 1, 70}}
	fmt.Println(highestGradeForEachStudent(enrollments))
}

func highestGradeForEachStudent(enrollments [][]int) [][]int {
	type best struct {
		grade    int
		courseID int
	}
	bestMap := make(map[int]best)

	for _, e := range enrollments {
		sid, cid, grade := e[0], e[1], e[2]
		if existing, ok := bestMap[sid]; !ok || grade > existing.grade || (grade == existing.grade && cid < existing.courseID) {
			bestMap[sid] = best{grade, cid}
		}
	}

	result := make([][]int, 0, len(bestMap))
	for sid, b := range bestMap {
		result = append(result, []int{sid, b.courseID, b.grade})
	}

	return result
}
```

## 1115 — Print Foobar Alternately

```go
package main

// LeetCode #1115: Print FooBar Alternately
// https://leetcode.com/problems/print-foobar-alternately/
// Difficulty: Medium
//
// Approach: Two goroutines synchronized with channels
// Time: O(n)
// Space: O(1)

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	fooCh := make(chan bool, 1)
	barCh := make(chan bool, 1)
	fooCh <- true

	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			<-fooCh
			fmt.Print("foo")
			barCh <- true
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			<-barCh
			fmt.Print("bar")
			fooCh <- true
		}
	}()

	wg.Wait()
	fmt.Println()
}
```

## 1116 — Print Zero Even Odd

```go
package main

// LeetCode #1116: Print Zero Even Odd
// https://leetcode.com/problems/print-zero-even-odd/
// Difficulty: Medium
//
// Approach: Three goroutines with channels for ordering
// Time: O(n)
// Space: O(1)

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	zeroCh := make(chan bool, 1)
	evenCh := make(chan bool, 1)
	oddCh := make(chan bool, 1)
	zeroCh <- true

	n := 5
	go func() {
		defer wg.Done()
		for i := 1; i <= n; i++ {
			<-zeroCh
			fmt.Print(0)
			if i%2 == 0 {
				evenCh <- true
			} else {
				oddCh <- true
			}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 2; i <= n; i += 2 {
			<-evenCh
			fmt.Print(i)
			zeroCh <- true
		}
	}()

	go func() {
		defer wg.Done()
		for i := 1; i <= n; i += 2 {
			<-oddCh
			fmt.Print(i)
			zeroCh <- true
		}
	}()

	wg.Wait()
	fmt.Println()
}
```

## 1117 — Building H2O

```go
package main

// LeetCode #1117: Building H2O
// https://leetcode.com/problems/building-h2o/
// Difficulty: Medium
//
// Approach: Use mutex and condition variable to synchronize H and O threads.
//           Barrier ensures exactly 2 H + 1 O per molecule.
// Time: O(n) where n = number of water molecules
// Space: O(1)

import (
	"fmt"
	"sync"
)

type H2O struct {
	mu      sync.Mutex
	hCount  int
	oCount  int
	hNeed   int
	oNeed   int
	cond    *sync.Cond
}

func NewH2O() *H2O {
	h2o := &H2O{hNeed: 2, oNeed: 1}
	h2o.cond = sync.NewCond(&h2o.mu)
	return h2o
}

func (h2o *H2O) Hydrogen(releaseHydrogen func()) {
	h2o.mu.Lock()
	for h2o.hCount >= h2o.hNeed {
		h2o.cond.Wait()
	}
	h2o.hCount++
	releaseHydrogen()
	if h2o.hCount == h2o.hNeed && h2o.oCount == h2o.oNeed {
		h2o.hCount = 0
		h2o.oCount = 0
	}
	h2o.cond.Broadcast()
	h2o.mu.Unlock()
}

func (h2o *H2O) Oxygen(releaseOxygen func()) {
	h2o.mu.Lock()
	for h2o.oCount >= h2o.oNeed {
		h2o.cond.Wait()
	}
	h2o.oCount++
	releaseOxygen()
	if h2o.hCount == h2o.hNeed && h2o.oCount == h2o.oNeed {
		h2o.hCount = 0
		h2o.oCount = 0
	}
	h2o.cond.Broadcast()
	h2o.mu.Unlock()
}

func main() {
	var wg sync.WaitGroup
	h2o := NewH2O()
	output := make([]byte, 0, 12)

	// Form 2 water molecules: need 4 H and 2 O
	atoms := []byte{'H', 'H', 'O', 'H', 'H', 'O'}
	var mu sync.Mutex

	for _, atom := range atoms {
		wg.Add(1)
		go func(a byte) {
			defer wg.Done()
			if a == 'H' {
				h2o.Hydrogen(func() {
					mu.Lock()
					output = append(output, 'H')
					mu.Unlock()
				})
			} else {
				h2o.Oxygen(func() {
					mu.Lock()
					output = append(output, 'O')
					mu.Unlock()
				})
			}
		}(atom)
	}

	wg.Wait()
	fmt.Println(string(output))
}
```

## 1120 — Maximum Average Subtree

```go
package main

// LeetCode #1120: Maximum Average Subtree
// https://leetcode.com/problems/maximum-average-subtree/
// Difficulty: Medium
//
// Approach: DFS post-order. Return sum and count for each subtree.
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
		Val: 5,
		Left: &TreeNode{
			Val:   6,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Println(maximumAverageSubtree(root)) // 6.0

	root2 := &TreeNode{
		Val: 2,
		Right: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Println(maximumAverageSubtree(root2)) // 1.5
}

func maximumAverageSubtree(root *TreeNode) float64 {
	maxAvg := 0.0
	dfs(root, &maxAvg)
	return maxAvg
}

func dfs(node *TreeNode, maxAvg *float64) (int, int) {
	if node == nil {
		return 0, 0
	}

	leftSum, leftCount := dfs(node.Left, maxAvg)
	rightSum, rightCount := dfs(node.Right, maxAvg)

	sum := node.Val + leftSum + rightSum
	count := 1 + leftCount + rightCount

	avg := float64(sum) / float64(count)
	if avg > *maxAvg {
		*maxAvg = avg
	}

	return sum, count
}
```

## 1123 — Lowest Common Ancestor Of Deepest Leaves

```go
package main

// LeetCode #1123: Lowest Common Ancestor of Deepest Leaves
// https://leetcode.com/problems/lowest-common-ancestor-of-deepest-leaves/
// Difficulty: Medium
//
// Approach: DFS. Return depth and LCA candidate for each subtree.
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
		Val: 3,
		Left: &TreeNode{
			Val:   5,
			Left:  &TreeNode{Val: 6, Left: nil, Right: nil},
			Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 7, Left: nil, Right: nil}, Right: &TreeNode{Val: 4, Left: nil, Right: nil}},
		},
		Right: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 0, Left: nil, Right: nil},
			Right: &TreeNode{Val: 8, Left: nil, Right: nil},
		},
	}
	result := lcaDeepestLeaves(root)
	fmt.Println(result.Val) // 2
}

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	_, lca := dfs(root)
	return lca
}

func dfs(node *TreeNode) (int, *TreeNode) {
	if node == nil {
		return 0, nil
	}

	leftDepth, leftLCA := dfs(node.Left)
	rightDepth, rightLCA := dfs(node.Right)

	if leftDepth > rightDepth {
		return leftDepth + 1, leftLCA
	} else if rightDepth > leftDepth {
		return rightDepth + 1, rightLCA
	}
	return leftDepth + 1, node
}
```

## 1124 — Longest Well Performing Interval

```go
package main

// LeetCode #1124: Longest Well-Performing Interval
// https://leetcode.com/problems/longest-well-performing-interval/
// Difficulty: Medium
//
// Approach: Prefix sum. Map first occurrence of each prefix sum.
//           hours[i] > 8 -> +1, else -1. Find longest subarray with sum > 0.
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(longestWPI([]int{9, 9, 6, 0, 6, 6, 9})) // 3
	fmt.Println(longestWPI([]int{6, 6, 6}))              // 0
}

func longestWPI(hours []int) int {
	prefix := 0
	firstSeen := make(map[int]int)
	result := 0

	for i, h := range hours {
		if h > 8 {
			prefix++
		} else {
			prefix--
		}

		if prefix > 0 {
			result = i + 1
		} else {
			if _, ok := firstSeen[prefix]; !ok {
				firstSeen[prefix] = i
			}
			if j, ok := firstSeen[prefix-1]; ok {
				if i-j > result {
					result = i - j
				}
			}
		}
	}

	return result
}
```

## 1126 — Active Businesses

```go
package main

// LeetCode #1126: Active Businesses
// https://leetcode.com/problems/active-businesses/
// Difficulty: Medium
//
// Approach: Count occurrences per (event_type, occurences), find avg per event_type.
//           Filter businesses where > 1 event type has > avg occurences.
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	// (business_id, event_type, occurences)
	// Biz 1: reviews=10, ads=7 -> reviews(10 > avg 6.67) ✓, ads(7 > avg 5) ✓ -> 2 above, active
	// Biz 2: reviews=3, ads=2  -> reviews(3 < 6.67) ✗, ads(2 < 5) ✗
	// Biz 3: reviews=7, ads=6  -> reviews(7 > 6.67) ✓, ads(6 > 5) ✓ -> 2 above, active
	events := [][]int{
		{1, 1, 10}, {1, 2, 7},
		{2, 1, 3}, {2, 2, 2},
		{3, 1, 7}, {3, 2, 6},
	}
	fmt.Println(activeBusinesses(events)) // [1,3]
}

func activeBusinesses(events [][]int) []int {
	// events[i] = [business_id, event_type, occurences]
	typeTotals := make(map[int]int)
	typeCount := make(map[int]int)
	bizEvents := make(map[int]map[int]int)

	for _, e := range events {
		bizID, eventType, occ := e[0], e[1], e[2]
		typeTotals[eventType] += occ
		typeCount[eventType]++
		if bizEvents[bizID] == nil {
			bizEvents[bizID] = make(map[int]int)
		}
		bizEvents[bizID][eventType] = occ
	}

	// Compute average per event type
	typeAvg := make(map[int]float64)
	for t := range typeTotals {
		typeAvg[t] = float64(typeTotals[t]) / float64(typeCount[t])
	}

	result := make([]int, 0)
	for bizID, events := range bizEvents {
		aboveAvgCount := 0
		for eventType, occ := range events {
			if float64(occ) > typeAvg[eventType] {
				aboveAvgCount++
			}
		}
		if aboveAvgCount > 1 {
			result = append(result, bizID)
		}
	}

	return result
}
```

## 1129 — Shortest Path With Alternating Colors

```go
package main

// LeetCode #1129: Shortest Path with Alternating Colors
// https://leetcode.com/problems/shortest-path-with-alternating-colors/
// Difficulty: Medium
//
// Approach: BFS with 2 states per node (reached by red edge / blue edge)
// Time: O(n + e)
// Space: O(n + e)

import "fmt"

func main() {
	fmt.Println(shortestAlternatingPaths(3, [][]int{{0, 1}}, [][]int{{2, 1}}))          // [0,1,-1]
	fmt.Println(shortestAlternatingPaths(3, [][]int{{0, 1}}, [][]int{{1, 2}}))          // [0,1,2]
}

func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	redGraph := make([][]int, n)
	blueGraph := make([][]int, n)

	for _, e := range redEdges {
		redGraph[e[0]] = append(redGraph[e[0]], e[1])
	}
	for _, e := range blueEdges {
		blueGraph[e[0]] = append(blueGraph[e[0]], e[1])
	}

	// dist[node][0] = distance reaching node via red edge
	// dist[node][1] = distance reaching node via blue edge
	dist := make([][2]int, n)
	for i := 1; i < n; i++ {
		dist[i] = [2]int{-1, -1}
	}

	queue := make([][2]int, 0)
	queue = append(queue, [2]int{0, 0}) // reached 0 via red (start counts as either)
	queue = append(queue, [2]int{0, 1}) // reached 0 via blue

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		node, color := cur[0], cur[1]
		nextColor := 1 - color

		if nextColor == 0 { // next edge should be red
			for _, next := range redGraph[node] {
				if dist[next][nextColor] == -1 {
					dist[next][nextColor] = dist[node][color] + 1
					queue = append(queue, [2]int{next, nextColor})
				}
			}
		} else { // next edge should be blue
			for _, next := range blueGraph[node] {
				if dist[next][nextColor] == -1 {
					dist[next][nextColor] = dist[node][color] + 1
					queue = append(queue, [2]int{next, nextColor})
				}
			}
		}
	}

	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = minDist(dist[i])
	}
	return result
}

func minDist(d [2]int) int {
	if d[0] == -1 {
		return d[1]
	}
	if d[1] == -1 {
		return d[0]
	}
	if d[0] < d[1] {
		return d[0]
	}
	return d[1]
}
```

## 1130 — Minimum Cost Tree From Leaf Values

```go
package main

// LeetCode #1130: Minimum Cost Tree From Leaf Values
// https://leetcode.com/problems/minimum-cost-tree-from-leaf-values/
// Difficulty: Medium
//
// Approach: Monotonic decreasing stack (greedy)
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(mctFromLeafValues([]int{6, 2, 4}))    // 32
	fmt.Println(mctFromLeafValues([]int{4, 11}))      // 44
}

func mctFromLeafValues(arr []int) int {
	stack := make([]int, 0)
	result := 0

	for _, v := range arr {
		for len(stack) > 0 && stack[len(stack)-1] <= v {
			mid := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				result += mid * v
			} else {
				if stack[len(stack)-1] < v {
					result += mid * stack[len(stack)-1]
				} else {
					result += mid * v
				}
			}
		}
		stack = append(stack, v)
	}

	for len(stack) > 1 {
		last := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result += last * stack[len(stack)-1]
	}

	return result
}
```

## 1131 — Maximum Of Absolute Value Expression

```go
package main

// LeetCode #1131: Maximum of Absolute Value Expression
// https://leetcode.com/problems/maximum-of-absolute-value-expression/
// Difficulty: Medium
//
// Approach: The expression |arr1[i] - arr1[j]| + |arr2[i] - arr2[j]| + |i - j|
//           can be expanded by considering all sign combinations.
//           Compute max of (arr1[i] + arr2[i] + i) - min of same, etc.
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(maxAbsValExpr([]int{1, -2, -5, 0, 10}, []int{0, -2, -1, -7, -4})) // 20
	fmt.Println(maxAbsValExpr([]int{1, 2, 3, 4}, []int{-1, 4, 5, 6}))             // 13
}

func maxAbsValExpr(arr1 []int, arr2 []int) int {
	n := len(arr1)
	// Try all sign combinations for the three terms
	// Each term can be +/-: a1 + a2 + i, a1 + a2 - i, a1 - a2 + i, etc.
	// Actually: |x| = max(x, -x), so we need 8 combinations = 2^3
	// But many are redundant. The key combinations are:
	// (arr1[i] + arr2[i] + i), (arr1[i] + arr2[i] - i),
	// (arr1[i] - arr2[i] + i), (arr1[i] - arr2[i] - i)
	// and their negations (which are just -arr1[i] - arr2[i] - i, etc.)

	combos := [4]int{}
	for k := 0; k < 4; k++ {
		combos[k] = arr1[0] + arr2[0] + 0
		if k&1 != 0 {
			combos[k] = arr1[0] + arr2[0] - 0
		}
		if k&2 != 0 {
			combos[k] = arr1[0] - arr2[0] + 0
			if k&1 != 0 {
				combos[k] = arr1[0] - arr2[0] - 0
			}
		}
	}

	result := 0
	for i := 0; i < n; i++ {
		for _, sign1 := range []int{-1, 1} {
			for _, sign2 := range []int{-1, 1} {
				for _, sign3 := range []int{-1, 1} {
					val := sign1*arr1[i] + sign2*arr2[i] + sign3*i
					for j := i + 1; j < n; j++ {
						val2 := sign1*arr1[j] + sign2*arr2[j] + sign3*j
						diff := val - val2
						if diff < 0 {
							diff = -diff
						}
						if diff > result {
							result = diff
						}
					}
				}
			}
		}
	}

	return result
}
```

## 1132 — Reported Posts Ii

```go
package main

// LeetCode #1132: Reported Posts II
// https://leetcode.com/problems/reported-posts-ii/
// Difficulty: Medium
//
// Approach: Calculate average daily spam removal percentage
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	// (date, post_id, action_type) where action_type: 0=report, 1=remove
	// Day 1: post 1 reported, post 2 reported
	// Day 2: post 1 removed (counts for day 1)
	// Day 3: post 3 reported
	// Day 1: remove rate = 1/2 = 50%. Day 3: remove rate = 0/1 = 0%. Avg = 25%
	actions := [][3]int{
		{1, 1, 0},
		{1, 2, 0},
		{2, 1, 1},
		{3, 3, 0},
	}
	fmt.Println(reportedPostsIi(actions))
}

func reportedPostsIi(actions [][3]int) float64 {
	// First pass: identify reported posts per date
	reported := make(map[int]map[int]bool)
	removedPosts := make(map[int]bool) // posts that were eventually removed

	for _, a := range actions {
		date, postID, actionType := a[0], a[1], a[2]
		if actionType == 0 { // report
			if reported[date] == nil {
				reported[date] = make(map[int]bool)
			}
			reported[date][postID] = true
		} else { // remove
			removedPosts[postID] = true
		}
	}

	totalPct := 0.0
	dayCount := 0

	for _, posts := range reported {
		reportCount := len(posts)
		if reportCount > 0 {
			removeCount := 0
			for postID := range posts {
				if removedPosts[postID] {
					removeCount++
				}
			}
			totalPct += float64(removeCount) / float64(reportCount) * 100
			dayCount++
		}
	}

	if dayCount == 0 {
		return 0
	}
	return totalPct / float64(dayCount)
}
```

## 1135 — Connecting Cities With Minimum Cost

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #1135: Connecting Cities With Minimum Cost
// https://leetcode.com/problems/connecting-cities-with-minimum-cost/

// Minimum cost to connect all cities (Kruskal's MST algorithm).

// Solution uses Union-Find and Kruskal's minimum spanning tree algorithm:
// 1. Sort edges by cost ascending
// 2. Connect cities using union-find, accumulating cost
// 3. If all cities are connected (single component), return cost; else -1

// Time complexity: O(E log E) = O(E log V) for sorting edges
// Space complexity: O(V) for the union-find data structure

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

func (uf *unionFind) union(x, y int) bool {
	px, py := uf.find(x), uf.find(y)
	if px == py {
		return false
	}
	if uf.rank[px] < uf.rank[py] {
		px, py = py, px
	}
	uf.parent[py] = px
	if uf.rank[px] == uf.rank[py] {
		uf.rank[px]++
	}
	return true
}

func (uf *unionFind) connected() bool {
	root := uf.find(0)
	for i := 1; i < len(uf.parent); i++ {
		if uf.find(i) != root {
			return false
		}
	}
	return true
}

func minimumCost(n int, connections [][]int) int {
	if n <= 1 {
		return 0
	}

	// Sort by cost ascending
	sort.Slice(connections, func(i, j int) bool {
		return connections[i][2] < connections[j][2]
	})

	uf := newUnionFind(n)
	totalCost := 0
	edgesUsed := 0

	for _, conn := range connections {
		u, v, cost := conn[0]-1, conn[1]-1, conn[2]
		if uf.union(u, v) {
			totalCost += cost
			edgesUsed++
			if edgesUsed == n-1 {
				return totalCost
			}
		}
	}

	return -1
}

func main() {
	// Test case 1: Example from LeetCode
	n := 3
	connections := [][]int{{1, 2, 5}, {1, 3, 6}, {2, 3, 1}}
	fmt.Printf("minimumCost(%v, %v) = %d (expected: 6)\n", n, connections, minimumCost(n, connections))

	// Test case 2: Not all cities connected
	n2 := 4
	connections2 := [][]int{{1, 2, 3}, {3, 4, 4}}
	fmt.Printf("minimumCost(%v, %v) = %d (expected: -1)\n", n2, connections2, minimumCost(n2, connections2))

	// Test case 3: Single city
	n3 := 1
	connections3 := [][]int{}
	fmt.Printf("minimumCost(%v, %v) = %d (expected: 0)\n", n3, connections3, minimumCost(n3, connections3))
}
```

## 1136 — Parallel Courses

```go
package main

import (
	"fmt"
)

// LeetCode #1136: Parallel Courses
// https://leetcode.com/problems/parallel-courses/
// Difficulty: Medium [Paid]

// topological sort (Kahn's algorithm) to find minimum semesters.

// Time: O(n + len(relations))
// Space: O(n + len(relations))

func minimumSemesters(n int, relations [][]int) int {
	adj := make([][]int, n+1)
	indeg := make([]int, n+1)

	for _, r := range relations {
		adj[r[0]] = append(adj[r[0]], r[1])
		indeg[r[1]]++
	}

	queue := make([]int, 0)
	for i := 1; i <= n; i++ {
		if indeg[i] == 0 {
			queue = append(queue, i)
		}
	}

	semesters := 0
	taken := 0

	for len(queue) > 0 {
		semesters++
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			taken++
			for _, next := range adj[cur] {
				indeg[next]--
				if indeg[next] == 0 {
					queue = append(queue, next)
				}
			}
		}
	}

	if taken != n {
		return -1
	}
	return semesters
}

func main() {
	fmt.Printf("%d (expected: 2)\n", minimumSemesters(3, [][]int{{1, 3}, {2, 3}}))
	fmt.Printf("%d (expected: -1)\n", minimumSemesters(3, [][]int{{1, 2}, {2, 3}, {3, 1}}))
	fmt.Printf("%d (expected: 1)\n", minimumSemesters(3, [][]int{}))
}
```

## 1138 — Alphabet Board Path

```go
package main

import (
	"fmt"
	"strings"
)

// LeetCode #1138: Alphabet Board Path
// https://leetcode.com/problems/alphabet-board-path/

// On an alphabet board (5x6, 'a' at (0,0), 'z' at (5,0)),
// we start at position 'a'. We can move U/D/L/R and append '!' to pick.
// Return the shortest sequence to spell the target string.

// The only tricky part is 'z' which is alone in the last row.
// When moving from or to 'z', we must avoid out-of-bounds moves.
// Strategy: always move U first, then L/R, then D (to handle 'z' adjacency).

// Time complexity: O(n * 5) = O(n) where n = len(target)
// Space complexity: O(n) for the result

func alphabetBoardPath(target string) string {
	var result strings.Builder
	curR, curC := 0, 0

	for _, ch := range target {
		idx := int(ch - 'a')
		targetR := idx / 5
		targetC := idx % 5

		// For 'z', we need to move up first then left/right
		// because there's no cell below 'z' and no cell to the right.
		// General case: move U first (never out of bounds), then L/R, then D
		for curR > targetR {
			result.WriteByte('U')
			curR--
		}
		for curC > targetC {
			result.WriteByte('L')
			curC--
		}
		for curC < targetC {
			result.WriteByte('R')
			curC++
		}
		for curR < targetR {
			result.WriteByte('D')
			curR++
		}

		result.WriteByte('!')
	}

	return result.String()
}

func main() {
	// Test case 1
	fmt.Printf("alphabetBoardPath(%q) = %q (expected: %q)\n",
		"leet", alphabetBoardPath("leet"), "DDR!UURRR!!DDD!")

	// Test case 2
	fmt.Printf("alphabetBoardPath(%q) = %q (expected: %q)\n",
		"code", alphabetBoardPath("code"), "RR!DDR!UUL!R!")

	// Test case 3: Single character
	fmt.Printf("alphabetBoardPath(%q) = %q (expected: %q)\n",
		"a", alphabetBoardPath("a"), "!")
}
```

## 1139 — Largest 1 Bordered Square

```go
package main

import (
	"fmt"
)

// LeetCode #1139: Largest 1-Bordered Square
// https://leetcode.com/problems/largest-1-bordered-square/
// Difficulty: Medium

// Precompute horizontal/vertical consecutive 1s, then check each cell as
// bottom-right corner.

// Time: O(m * n * min(m, n))
// Space: O(m * n)

func largest1BorderedSquare(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	hor := make([][]int, m)
	ver := make([][]int, m)
	for i := 0; i < m; i++ {
		hor[i] = make([]int, n)
		ver[i] = make([]int, n)
	}
	maxSide := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				if j == 0 {
					hor[i][j] = 1
				} else {
					hor[i][j] = hor[i][j-1] + 1
				}
				if i == 0 {
					ver[i][j] = 1
				} else {
					ver[i][j] = ver[i-1][j] + 1
				}
				minSide := hor[i][j]
				if ver[i][j] < minSide {
					minSide = ver[i][j]
				}
				for s := minSide; s > maxSide; s-- {
					if hor[i-s+1][j] >= s && ver[i][j-s+1] >= s {
						maxSide = s
						break
					}
				}
			}
		}
	}
	return maxSide * maxSide
}

func main() {
	fmt.Printf("%d (expected: 9)\n", largest1BorderedSquare([][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}))
	fmt.Printf("%d (expected: 1)\n", largest1BorderedSquare([][]int{{1, 1, 0, 0}}))
	fmt.Printf("%d (expected: 0)\n", largest1BorderedSquare([][]int{{0}}))
}
```

## 1140 — Stone Game Ii

```go
package main

import (
	"fmt"
)

// LeetCode #1140: Stone Game II
// https://leetcode.com/problems/stone-game-ii/
// Difficulty: Medium

// DP with memoization. dp[i][m] = max stones current player can get
// from piles[i:] with current M = m.

// Time: O(n^3) but typically O(n^2)
// Space: O(n^2)

func stoneGameII(piles []int) int {
	n := len(piles)
	suffix := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		suffix[i] = suffix[i+1] + piles[i]
	}

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := n - 1; i >= 0; i-- {
		for m := 1; m <= n; m++ {
			best := 0
			for x := 1; x <= 2*m && i+x <= n; x++ {
				opponent := dp[i+x][max(m, x)]
				player := suffix[i] - opponent
				if player > best {
					best = player
				}
			}
			dp[i][m] = best
		}
	}
	return dp[0][1]
}

func main() {
	fmt.Printf("%d (expected: 10)\n", stoneGameII([]int{2, 7, 9, 4, 4}))
	fmt.Printf("%d (expected: 104)\n", stoneGameII([]int{1, 2, 3, 4, 5, 100}))
	fmt.Printf("%d (expected: 1)\n", stoneGameII([]int{1}))
}
```

## 1143 — Longest Common Subsequence

```go
package main

import (
	"fmt"
)

// LeetCode #1143: Longest Common Subsequence
// https://leetcode.com/problems/longest-common-subsequence/

// Given two strings text1 and text2, return the length of their
// longest common subsequence. If there is no common subsequence, return 0.

// Solution uses classic 2D DP:
// dp[i][j] = LCS of text1[:i] and text2[:j]
// If text1[i-1] == text2[j-1]: dp[i][j] = dp[i-1][j-1] + 1
// Else: dp[i][j] = max(dp[i-1][j], dp[i][j-1])

// Time complexity: O(m*n)
// Space complexity: O(m*n), can be optimized to O(min(m,n))

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
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

func main() {
	// Test case 1
	fmt.Printf("longestCommonSubsequence(%q, %q) = %d (expected: 3)\n",
		"abcde", "ace", longestCommonSubsequence("abcde", "ace"))

	// Test case 2
	fmt.Printf("longestCommonSubsequence(%q, %q) = %d (expected: 3)\n",
		"abc", "abc", longestCommonSubsequence("abc", "abc"))

	// Test case 3: No common subsequence
	fmt.Printf("longestCommonSubsequence(%q, %q) = %d (expected: 0)\n",
		"abc", "def", longestCommonSubsequence("abc", "def"))
}
```

## 1144 — Decrease Elements To Make Array Zigzag

```go
package main

import (
	"fmt"
)

// LeetCode #1144: Decrease Elements To Make Array Zigzag
// https://leetcode.com/problems/decrease-elements-to-make-array-zigzag/
// Difficulty: Medium

// We can make array zigzag in 2 patterns:
// A[0] < A[1] > A[2] < A[3] > ...  (even-indexed are valleys)
// A[0] > A[1] < A[2] > A[3] < ...  (odd-indexed are valleys)
// Compute min total decreases for each pattern.

// Time: O(n)
// Space: O(1)

func movesToMakeZigzag(nums []int) int {
	n := len(nums)
	movesEven := 0 // even indices are valleys (less than neighbors)
	movesOdd := 0  // odd indices are valleys

	for i, v := range nums {
		left := 1001
		if i > 0 {
			left = nums[i-1]
		}
		right := 1001
		if i < n-1 {
			right = nums[i+1]
		}
		neighbor := left
		if right < neighbor {
			neighbor = right
		}
		if i%2 == 0 {
			// even index: should be smaller than neighbors (valley)
			if v >= neighbor {
				movesEven += v - (neighbor - 1)
			}
		} else {
			// odd index: should be smaller than neighbors (valley)
			if v >= neighbor {
				movesOdd += v - (neighbor - 1)
			}
		}
	}

	// For pattern 2: even indices should be peaks, odd indices valleys
	movesEven2 := 0
	movesOdd2 := 0
	for i, v := range nums {
		left := 1001
		if i > 0 {
			left = nums[i-1]
		}
		right := 1001
		if i < n-1 {
			right = nums[i+1]
		}
		neighbor := left
		if right < neighbor {
			neighbor = right
		}
		if i%2 == 0 {
			// even index: should be larger than neighbors (peak)
			if v >= neighbor {
				movesOdd2 += v - (neighbor - 1)
			}
		} else {
			// odd index: should be larger than neighbors (peak)
			if v >= neighbor {
				movesEven2 += v - (neighbor - 1)
			}
		}
	}

	result := movesEven + movesOdd
	if movesEven2+movesOdd2 < result {
		result = movesEven2 + movesOdd2
	}
	return result
}

func main() {
	fmt.Printf("%d (expected: 2)\n", movesToMakeZigzag([]int{1, 2, 3}))
	fmt.Printf("%d (expected: 4)\n", movesToMakeZigzag([]int{9, 6, 1, 6, 2}))
	fmt.Printf("%d (expected: 0)\n", movesToMakeZigzag([]int{1, 3, 2}))
}
```

## 1145 — Binary Tree Coloring Game

```go
package main

import (
	"fmt"
)

// LeetCode #1145: Binary Tree Coloring Game
// https://leetcode.com/problems/binary-tree-coloring-game/
// Difficulty: Medium

// Two players color a binary tree. Player 1 colors node x blue.
// Player 2 can choose any other node as red. Then they alternate.
// A move = color an uncolored neighbor of your colored node.
// Player 2 wins if they can color more nodes.

// Player 1's initial node x splits tree into 3 components:
// left subtree, right subtree, and rest of tree.
// Player 2 should choose the root of the largest component.
// Player 2 wins if max(leftSize, rightSize, restSize) > n/2.

// Time: O(n)
// Space: O(h) where h is tree height (recursion)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
	var xNode *TreeNode
	var findX func(node *TreeNode)
	findX = func(node *TreeNode) {
		if node == nil || xNode != nil {
			return
		}
		if node.Val == x {
			xNode = node
			return
		}
		findX(node.Left)
		findX(node.Right)
	}
	findX(root)

	var count func(node *TreeNode) int
	count = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		return 1 + count(node.Left) + count(node.Right)
	}

	leftSize := count(xNode.Left)
	rightSize := count(xNode.Right)
	restSize := n - leftSize - rightSize - 1

	maxSize := leftSize
	if rightSize > maxSize {
		maxSize = rightSize
	}
	if restSize > maxSize {
		maxSize = restSize
	}

	return maxSize > n/2
}

func main() {
	// Test: root = [1,2,3,4,5,6,7,8,9,10,11], n = 11, x = 3
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{Val: 4,
				Left:  &TreeNode{Val: 8},
				Right: &TreeNode{Val: 9}},
			Right: &TreeNode{Val: 5,
				Left: &TreeNode{Val: 10},
				Right: &TreeNode{Val: 11}}},
		Right: &TreeNode{
			Val: 3,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 7}},
	}
	fmt.Printf("%t (expected: true)\n", btreeGameWinningMove(root, 11, 3))

	// Simple tree: [1,2,3], n=3, x=1 -> left(1), right(1), rest(0) -> max=1 <= 1.5 -> false
	root2 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	fmt.Printf("%t (expected: false)\n", btreeGameWinningMove(root2, 3, 1))

	// Single node
	root3 := &TreeNode{Val: 1}
	fmt.Printf("%t (expected: false)\n", btreeGameWinningMove(root3, 1, 1))
}
```

## 1146 — Snapshot Array

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #1146: Snapshot Array
// https://leetcode.com/problems/snapshot-array/
// Difficulty: Medium

// SnapshotArray supports set(index, val), snap(), and get(index, snap_id).
// Instead of copying full array on each snap, store (snap_id, value) history per index.

// Time: set O(1) amortized, snap O(1), get O(log k) where k = history length
// Space: O(n + total_sets)

type SnapshotArray struct {
	data   [][]pair
	snapID int
}

type pair struct {
	snapID int
	val    int
}

func Constructor(length int) SnapshotArray {
	data := make([][]pair, length)
	for i := 0; i < length; i++ {
		data[i] = make([]pair, 0)
		// Initialize with snapID -1, value 0
		data[i] = append(data[i], pair{-1, 0})
	}
	return SnapshotArray{data: data, snapID: 0}
}

func (this *SnapshotArray) Set(index int, val int) {
	idx := &this.data[index]
	// If last entry has same snapID, just update value
	if len(*idx) > 0 && (*idx)[len(*idx)-1].snapID == this.snapID {
		(*idx)[len(*idx)-1].val = val
		return
	}
	*idx = append(*idx, pair{this.snapID, val})
}

func (this *SnapshotArray) Snap() int {
	id := this.snapID
	this.snapID++
	return id
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
	history := this.data[index]
	// Binary search for the largest snapID <= snap_id
	i := sort.Search(len(history), func(i int) bool {
		return history[i].snapID > snap_id
	})
	return history[i-1].val
}

func main() {
	sa := Constructor(3)
	sa.Set(0, 5)
	fmt.Printf("snap: %d (expected: 0)\n", sa.Snap())
	sa.Set(0, 6)
	fmt.Printf("get(0, 0) = %d (expected: 5)\n", sa.Get(0, 0))
	fmt.Printf("get(0, 1) = %d (expected: 6)\n", sa.Get(0, 1))

	sa2 := Constructor(1)
	sa2.Set(0, 1)
	fmt.Printf("snap: %d (expected: 0)\n", sa2.Snap())
	fmt.Printf("get(0, 0) = %d (expected: 1)\n", sa2.Get(0, 0))
	sa2.Set(0, 2)
	fmt.Printf("snap: %d (expected: 1)\n", sa2.Snap())
	fmt.Printf("get(0, 0) = %d (expected: 1)\n", sa2.Get(0, 0))
	fmt.Printf("get(0, 1) = %d (expected: 2)\n", sa2.Get(0, 1))
}
```

## 1149 — Article Views Ii

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #1149: Article Views II
// https://leetcode.com/problems/article-views-ii/
// Difficulty: Medium [Paid]

// Given a table with viewer_id, article_id, and view_date, find all viewers
// who viewed at least two distinct articles on the same day.
// Return result sorted by viewer_id.

// Time: O(n log n)
// Space: O(n)

func articleViewsII(views [][]int) []int {
	// views[i] = [viewer_id, article_id, view_date]
	type view struct {
		viewerID int
		date     int
		article  int
	}

	type key struct {
		viewerID int
		date     int
	}

	seen := make(map[key]map[int]bool)

	for _, v := range views {
		k := key{v[0], v[2]}
		if seen[k] == nil {
			seen[k] = make(map[int]bool)
		}
		seen[k][v[1]] = true
	}

	resultSet := make(map[int]bool)
	for k, articles := range seen {
		if len(articles) >= 2 {
			resultSet[k.viewerID] = true
		}
	}

	result := make([]int, 0, len(resultSet))
	for id := range resultSet {
		result = append(result, id)
	}
	sort.Ints(result)
	return result
}

func main() {
	views := [][]int{
		{1, 1, 20200101},
		{1, 2, 20200101},
		{2, 3, 20200101},
		{1, 3, 20200102},
	}
	fmt.Printf("%v (expected: [1])\n", articleViewsII(views))

	views2 := [][]int{
		{1, 1, 20200101},
		{1, 2, 20200101},
		{2, 1, 20200101},
		{2, 3, 20200101},
	}
	fmt.Printf("%v (expected: [1 2])\n", articleViewsII(views2))

	views3 := [][]int{
		{1, 1, 20200101},
		{1, 1, 20200101},
	}
	fmt.Printf("%v (expected: [])\n", articleViewsII(views3))
}
```

## 1151 — Minimum Swaps To Group All 1S Together

```go
package main

import (
	"fmt"
)

// LeetCode #1151: Minimum Swaps to Group All 1's Together
// https://leetcode.com/problems/minimum-swaps-to-group-all-1s-together/
// Difficulty: Medium [Paid]

// Use sliding window of size = total number of 1s.
// Minimum swaps = window with maximum 1s (least 0s to swap out).

// Time: O(n)
// Space: O(1)

func minSwaps(data []int) int {
	totalOnes := 0
	for _, v := range data {
		if v == 1 {
			totalOnes++
		}
	}
	if totalOnes <= 1 {
		return 0
	}

	currOnes := 0
	for i := 0; i < totalOnes; i++ {
		if data[i] == 1 {
			currOnes++
		}
	}

	maxOnes := currOnes
	for i := totalOnes; i < len(data); i++ {
		if data[i] == 1 {
			currOnes++
		}
		if data[i-totalOnes] == 1 {
			currOnes--
		}
		if currOnes > maxOnes {
			maxOnes = currOnes
		}
	}

	return totalOnes - maxOnes
}

func main() {
	fmt.Printf("%d (expected: 1)\n", minSwaps([]int{1, 0, 1, 0, 1}))
	fmt.Printf("%d (expected: 0)\n", minSwaps([]int{0, 0, 0, 1, 0}))
	fmt.Printf("%d (expected: 2)\n", minSwaps([]int{1, 0, 1, 0, 1, 0, 0, 1, 1, 0, 1}))
}
```

## 1152 — Analyze User Website Visit Pattern

```go
package main

import (
	"fmt"
	"sort"
	"strings"
)

// LeetCode #1152: Analyze User Website Visit Pattern
// https://leetcode.com/problems/analyze-user-website-visit-pattern/
// Difficulty: Medium [Paid]

// Given usernames, timestamps, and websites, find the most visited
// 3-sequence pattern (3 websites in order). Ties broken by lexicographic order.

// Time: O(n^3) worst, but constrained by problem input size
// Space: O(n)

func mostVisitedPattern(username []string, timestamp []int, website []string) []string {
	n := len(username)

	// Group visits by user
	userVisits := make(map[string][]visit)
	for i := 0; i < n; i++ {
		userVisits[username[i]] = append(userVisits[username[i]], visit{timestamp[i], website[i]})
	}

	// Sort each user's visits by timestamp
	for user := range userVisits {
		sort.Slice(userVisits[user], func(i, j int) bool {
			return userVisits[user][i].time < userVisits[user][j].time
		})
	}

	// Count patterns across users
	patternCount := make(map[string]int)

	for _, visits := range userVisits {
		if len(visits) < 3 {
			continue
		}
		userPatterns := make(map[string]bool)
		// Generate all 3-sequences for this user
		for i := 0; i < len(visits); i++ {
			for j := i + 1; j < len(visits); j++ {
				for k := j + 1; k < len(visits); k++ {
					pat := visits[i].site + "," + visits[j].site + "," + visits[k].site
					userPatterns[pat] = true
				}
			}
		}
		for pat := range userPatterns {
			patternCount[pat]++
		}
	}

	// Find best pattern
	bestPat := ""
	bestCount := 0
	for pat, count := range patternCount {
		if count > bestCount || (count == bestCount && (bestPat == "" || pat < bestPat)) {
			bestCount = count
			bestPat = pat
		}
	}

	return strings.Split(bestPat, ",")
}

type visit struct {
	time int
	site string
}

func main() {
	username := []string{"joe", "joe", "joe", "james", "james", "james", "james", "mary", "mary", "mary"}
	timestamp := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	website := []string{"home", "about", "career", "home", "cart", "maps", "home", "home", "about", "career"}
	fmt.Printf("%v (expected: [\"home\" \"about\" \"career\"])\n", mostVisitedPattern(username, timestamp, website))

	username2 := []string{"u1", "u1", "u1", "u2", "u2", "u2"}
	timestamp2 := []int{1, 2, 3, 4, 5, 6}
	website2 := []string{"a", "b", "c", "a", "b", "c"}
	fmt.Printf("%v (expected: [\"a\" \"b\" \"c\"])\n", mostVisitedPattern(username2, timestamp2, website2))
}
```

## 1155 — Number Of Dice Rolls With Target Sum

```go
package main

import (
	"fmt"
)

// LeetCode #1155: Number of Dice Rolls With Target Sum
// https://leetcode.com/problems/number-of-dice-rolls-with-target-sum/
// Difficulty: Medium

// dp[d][t] = number of ways to get sum t using d dice with k faces.
// dp[d][t] = sum(dp[d-1][t-f] for f in 1..k if t-f >= 0)

// Time: O(n * target * k)
// Space: O(target) with 1D DP optimization

func numRollsToTarget(n int, k int, target int) int {
	const mod = 1_000_000_007

	dp := make([]int, target+1)
	dp[0] = 1

	for dice := 0; dice < n; dice++ {
		next := make([]int, target+1)
		for sum := 0; sum <= target; sum++ {
			if dp[sum] == 0 {
				continue
			}
			for face := 1; face <= k && sum+face <= target; face++ {
				next[sum+face] = (next[sum+face] + dp[sum]) % mod
			}
		}
		dp = next
	}
	return dp[target]
}

func main() {
	fmt.Printf("%d (expected: 1)\n", numRollsToTarget(1, 6, 3))
	fmt.Printf("%d (expected: 6)\n", numRollsToTarget(2, 6, 7))
	fmt.Printf("%d (expected: 222616187)\n", numRollsToTarget(30, 30, 500))
}
```

## 1156 — Swap For Longest Repeated Character Substring

```go
package main

import (
	"fmt"
)

// LeetCode #1156: Swap For Longest Repeated Character Substring
// https://leetcode.com/problems/swap-for-longest-repeated-character-substring/
// Difficulty: Medium

// For each character, find the longest group. We can swap one character
// from elsewhere. Consider groups separated by exactly 1 char.

// Time: O(n)
// Space: O(n)

func maxRepOpt1(text string) int {
	n := len(text)

	// Count total frequency per character
	freq := make(map[byte]int)
	for i := 0; i < n; i++ {
		freq[text[i]]++
	}

	// Compute groups: (char, length)
	type group struct {
		char byte
		len  int
	}
	groups := make([]group, 0)

	i := 0
	for i < n {
		j := i
		for j < n && text[j] == text[i] {
			j++
		}
		groups = append(groups, group{text[i], j - i})
		i = j
	}

	maxLen := 0

	for idx, g := range groups {
		// Case 1: extend group by 1 if there's a spare of this char
		if freq[g.char] > g.len {
			if g.len+1 > maxLen {
				maxLen = g.len + 1
			}
		} else {
			if g.len > maxLen {
				maxLen = g.len
			}
		}

		// Case 2: merge two groups separated by exactly 1 char
		if idx > 0 && idx < len(groups)-1 &&
			groups[idx-1].char == groups[idx+1].char &&
			g.len == 1 {
			total := groups[idx-1].len + groups[idx+1].len
			if freq[groups[idx-1].char] > total {
				total++
			}
			if total > maxLen {
				maxLen = total
			}
		}
	}

	return maxLen
}

func main() {
	fmt.Printf("%d (expected: 3)\n", maxRepOpt1("ababa"))
	fmt.Printf("%d (expected: 6)\n", maxRepOpt1("aaabaaa"))
	fmt.Printf("%d (expected: 4)\n", maxRepOpt1("aaaaa"))
}
```

## 1158 — Market Analysis I

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #1158: Market Analysis I
// https://leetcode.com/problems/market-analysis-i/
// Difficulty: Medium

// For each user, find total orders placed in 2019 and their join date.
// Sort by buyer_id.

// Time: O(n log n)
// Space: O(n)

type order struct {
	orderID    int
	orderDate  string
	productID  int
	buyerID    int
	sellerID   int
}

type user struct {
	userID    int
	joinDate  string
}

func marketAnalysisI(users []user, orders []order) [][]int {
	orderCount := make(map[int]int)
	for _, o := range orders {
		if len(o.orderDate) >= 4 && o.orderDate[:4] == "2019" {
			orderCount[o.buyerID]++
		}
	}

	result := make([][]int, 0, len(users))
	for _, u := range users {
		count := orderCount[u.userID]
		result = append(result, []int{u.userID, count})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i][0] < result[j][0]
	})

	return result
}

func main() {
	users := []user{
		{1, "2018-01-01"},
		{2, "2018-01-01"},
	}
	orders := []order{
		{1, "2019-01-01", 1, 1, 2},
		{2, "2020-01-01", 2, 1, 2},
		{3, "2019-05-01", 3, 2, 1},
	}
	fmt.Printf("%v (expected: [[1 1] [2 1]])\n", marketAnalysisI(users, orders))

	users2 := []user{
		{1, "2019-01-01"},
		{2, "2019-01-01"},
	}
	orders2 := []order{
		{1, "2020-01-01", 1, 1, 2},
		{2, "2020-05-01", 2, 2, 1},
	}
	fmt.Printf("%v (expected: [[1 0] [2 0]])\n", marketAnalysisI(users2, orders2))
}
```

## 1161 — Maximum Level Sum Of A Binary Tree

```go
package main

import (
	"fmt"
)

// LeetCode #1161: Maximum Level Sum of a Binary Tree
// https://leetcode.com/problems/maximum-level-sum-of-a-binary-tree/
// Difficulty: Medium

// BFS level order traversal, track sum per level.

// Time: O(n)
// Space: O(n)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	maxSum := root.Val
	maxLevel := 1
	currentLevel := 0

	for len(queue) > 0 {
		currentLevel++
		size := len(queue)
		sum := 0

		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			sum += node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		if sum > maxSum {
			maxSum = sum
			maxLevel = currentLevel
		}
	}

	return maxLevel
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   7,
			Left:  &TreeNode{Val: 7},
			Right: &TreeNode{Val: -8},
		},
		Right: &TreeNode{Val: 0},
	}
	fmt.Printf("%d (expected: 2)\n", maxLevelSum(root))

	root2 := &TreeNode{
		Val:  -100,
		Left: &TreeNode{Val: -200},
		Right: &TreeNode{
			Val:   -300,
			Left:  &TreeNode{Val: -20},
			Right: &TreeNode{Val: -5},
		},
	}
	fmt.Printf("%d (expected: 3)\n", maxLevelSum(root2))
}
```

## 1162 — As Far From Land As Possible

```go
package main

import (
	"fmt"
)

// LeetCode #1162: As Far from Land as Possible
// https://leetcode.com/problems/as-far-from-land-as-possible/
// Difficulty: Medium

// Multi-source BFS from all land cells simultaneously.
// The last water cell to be reached has the max distance.

// Time: O(m*n)
// Space: O(m*n)

func maxDistance(grid [][]int) int {
	n := len(grid)
	if n == 0 {
		return -1
	}

	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	queue := make([][2]int, 0)

	// Add all land cells to queue
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				queue = append(queue, [2]int{i, j})
			}
		}
	}

	// If all water or all land, return -1
	if len(queue) == 0 || len(queue) == n*n {
		return -1
	}

	distance := -1
	for len(queue) > 0 {
		distance++
		size := len(queue)
		for i := 0; i < size; i++ {
			r, c := queue[0][0], queue[0][1]
			queue = queue[1:]
			for _, d := range dirs {
				nr, nc := r+d[0], c+d[1]
				if nr >= 0 && nr < n && nc >= 0 && nc < n && grid[nr][nc] == 0 {
					grid[nr][nc] = 1 // mark visited
					queue = append(queue, [2]int{nr, nc})
				}
			}
		}
	}

	return distance
}

func main() {
	grid1 := [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}
	fmt.Printf("%d (expected: 2)\n", maxDistance(grid1))

	grid2 := [][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 0}}
	fmt.Printf("%d (expected: 4)\n", maxDistance(grid2))

	grid3 := [][]int{{0, 0}, {0, 0}}
	fmt.Printf("%d (expected: -1)\n", maxDistance(grid3))
}
```

## 1164 — Product Price At A Given Date

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #1164: Product Price at a Given Date
// https://leetcode.com/problems/product-price-at-a-given-date/
// Difficulty: Medium

// For each product, find its price on 2019-08-16.
// If no change after date, price is 10 (default).
// Find the most recent price change on or before the date.

// Time: O(n log n)
// Space: O(n)

type productPrice struct {
	productID int
	price     int
}

func productPriceAtDate(products [][]int) []productPrice {
	// products[i] = [product_id, new_price, change_date]
	// Find price of each product on 2019-08-16

	// Group by product
	changes := make(map[int][][2]int) // productID -> [(date, price)]
	for _, p := range products {
		id, price, date := p[0], p[1], p[2]
		changes[id] = append(changes[id], [2]int{date, price})
	}

	// Sort each product's changes by date
	for id := range changes {
		sort.Slice(changes[id], func(i, j int) bool {
			return changes[id][i][0] < changes[id][j][0]
		})
	}

	targetDate := 20190816
	result := make([]productPrice, 0, len(changes))

	for id, vals := range changes {
		price := 10 // default price
		for _, v := range vals {
			if v[0] <= targetDate {
				price = v[1]
			} else {
				break
			}
		}
		result = append(result, productPrice{id, price})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].productID < result[j].productID
	})

	return result
}

func main() {
	products := [][]int{
		{1, 20, 20190801},
		{2, 50, 20190801},
		{1, 30, 20190815},
		{1, 40, 20190817},
		{2, 80, 20190814},
	}
	fmt.Printf("%v (expected: [{1 30} {2 80}])\n", productPriceAtDate(products))

	products2 := [][]int{
		{1, 20, 20190817},
		{2, 50, 20190817},
	}
	fmt.Printf("%v (expected: [{1 10} {2 10}])\n", productPriceAtDate(products2))
}
```

## 1166 — Design File System

```go
package main

import (
	"fmt"
	"strings"
)

// LeetCode #1166: Design File System
// https://leetcode.com/problems/design-file-system/
// Difficulty: Medium [Paid]

// FileSystem supports createPath(path, value) and get(path).
// Paths are absolute, use "/" as delimiter.

// Time: createPath O(L) where L = path depth, get O(L)
// Space: O(P) where P = number of paths

type FileSystem struct {
	paths map[string]int
}

func Constructor() FileSystem {
	return FileSystem{paths: make(map[string]int)}
}

func (this *FileSystem) CreatePath(path string, value int) bool {
	if path == "" || path == "/" {
		return false
	}
	if _, exists := this.paths[path]; exists {
		return false
	}

	// Check parent exists (unless parent is root)
	lastSlash := strings.LastIndex(path, "/")
	if lastSlash > 0 {
		parent := path[:lastSlash]
		if _, exists := this.paths[parent]; !exists {
			return false
		}
	}

	this.paths[path] = value
	return true
}

func (this *FileSystem) Get(path string) int {
	if val, exists := this.paths[path]; exists {
		return val
	}
	return -1
}

func main() {
	fs := Constructor()
	fmt.Printf("%t (expected: true)\n", fs.CreatePath("/a", 1))
	fmt.Printf("%d (expected: 1)\n", fs.Get("/a"))
	fmt.Printf("%t (expected: true)\n", fs.CreatePath("/a/b", 2))
	fmt.Printf("%d (expected: 2)\n", fs.Get("/a/b"))
	fmt.Printf("%t (expected: false)\n", fs.CreatePath("/a/b", 3)) // already exists
	fmt.Printf("%t (expected: false)\n", fs.CreatePath("/c/d", 4)) // parent doesn't exist
	fmt.Printf("%d (expected: -1)\n", fs.Get("/c"))
}
```

## 1167 — Minimum Cost To Connect Sticks

```go
package main

import (
	"container/heap"
	"fmt"
)

// LeetCode #1167: Minimum Cost to Connect Sticks
// https://leetcode.com/problems/minimum-cost-to-connect-sticks/
// Difficulty: Medium [Paid]

// Always combine the two smallest sticks (greedy + min-heap).
// Cost of each combination = sum of the two sticks.

// Time: O(n log n)
// Space: O(n)

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func connectSticks(sticks []int) int {
	if len(sticks) <= 1 {
		return 0
	}

	h := &MinHeap{}
	heap.Init(h)
	for _, s := range sticks {
		heap.Push(h, s)
	}

	total := 0
	for h.Len() > 1 {
		a := heap.Pop(h).(int)
		b := heap.Pop(h).(int)
		cost := a + b
		total += cost
		heap.Push(h, cost)
	}

	return total
}

func main() {
	fmt.Printf("%d (expected: 14)\n", connectSticks([]int{2, 4, 3}))
	fmt.Printf("%d (expected: 30)\n", connectSticks([]int{1, 8, 3, 5}))
	fmt.Printf("%d (expected: 0)\n", connectSticks([]int{5}))
}
```

## 1169 — Invalid Transactions

```go
package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// LeetCode #1169: Invalid Transactions
// https://leetcode.com/problems/invalid-transactions/
// Difficulty: Medium

// A transaction is invalid if: amount > 1000, or same name within
// 60 min of another transaction at different city.

// Time: O(n^2) due to pairwise comparison
// Space: O(n)

type Transaction struct {
	name   string
	time   int
	amount int
	city   string
	idx    int
}

func invalidTransactions(transactions []string) []string {
	n := len(transactions)
	txs := make([]Transaction, n)

	for i, t := range transactions {
		parts := strings.Split(t, ",")
		tm, _ := strconv.Atoi(parts[1])
		amt, _ := strconv.Atoi(parts[2])
		txs[i] = Transaction{parts[0], tm, amt, parts[3], i}
	}

	invalid := make([]bool, n)

	for i := 0; i < n; i++ {
		if txs[i].amount > 1000 {
			invalid[txs[i].idx] = true
		}
		for j := i + 1; j < n; j++ {
			if txs[i].name == txs[j].name &&
				abs(txs[i].time-txs[j].time) <= 60 &&
				txs[i].city != txs[j].city {
				invalid[txs[i].idx] = true
				invalid[txs[j].idx] = true
			}
		}
	}

	result := make([]string, 0)
	for i, v := range invalid {
		if v {
			result = append(result, transactions[i])
		}
	}
	sort.Strings(result)
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	t1 := []string{"alice,20,800,mtv", "alice,50,100,beijing"}
	fmt.Printf("%v (expected: [alice,20,800,mtv alice,50,100,beijing])\n", invalidTransactions(t1))

	t2 := []string{"alice,20,800,mtv", "bob,50,1200,mtv"}
	fmt.Printf("%v (expected: [bob,50,1200,mtv])\n", invalidTransactions(t2))

	t3 := []string{"alice,20,800,mtv", "alice,50,100,mtv"}
	fmt.Printf("%v (expected: [])\n", invalidTransactions(t3))
}
```

## 1170 — Compare Strings By Frequency Of The Smallest Character

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #1170: Compare Strings by Frequency of the Smallest Character
// https://leetcode.com/problems/compare-strings-by-frequency-of-the-smallest-character/
// Difficulty: Medium

// f(s) = frequency of smallest character in s.
// For each query word, count words in words[] with f(w) > f(query).

// Time: O((n + m) * L) where L = average string length
// Space: O(n)

func f(s string) int {
	minChar := s[0]
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] < minChar {
			minChar = s[i]
			count = 1
		} else if s[i] == minChar {
			count++
		}
	}
	return count
}

func numSmallerByFrequency(queries []string, words []string) []int {
	wordFreqs := make([]int, len(words))
	for i, w := range words {
		wordFreqs[i] = f(w)
	}
	sort.Ints(wordFreqs)

	result := make([]int, len(queries))
	for i, q := range queries {
		qf := f(q)
		// Binary search for first wordFreq > qf
		idx := sort.Search(len(wordFreqs), func(j int) bool {
			return wordFreqs[j] > qf
		})
		result[i] = len(wordFreqs) - idx
	}
	return result
}

func main() {
	fmt.Printf("%v (expected: [1])\n", numSmallerByFrequency([]string{"cbd"}, []string{"zaaaz"}))
	fmt.Printf("%v (expected: [1 2])\n", numSmallerByFrequency([]string{"bbb", "cc"}, []string{"a", "aa", "aaa", "aaaa"}))
}
```

## 1171 — Remove Zero Sum Consecutive Nodes From Linked List

```go
package main

import (
	"fmt"
)

// LeetCode #1171: Remove Zero Sum Consecutive Nodes from Linked List
// https://leetcode.com/problems/remove-zero-sum-consecutive-nodes-from-linked-list/
// Difficulty: Medium

// Use prefix sum with a hashmap. If same prefix sum repeats,
// the segment between has sum 0, so remove it.

// Time: O(n)
// Space: O(n)

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeZeroSumSublists(head *ListNode) *ListNode {
	dummy := &ListNode{0, head}
	prefix := 0
	seen := make(map[int]*ListNode)
	seen[0] = dummy

	for cur := dummy; cur != nil; cur = cur.Next {
		prefix += cur.Val
		seen[prefix] = cur
	}

	prefix = 0
	for cur := dummy; cur != nil; cur = cur.Next {
		prefix += cur.Val
		cur.Next = seen[prefix].Next
	}

	return dummy.Next
}

func listToSlice(head *ListNode) []int {
	result := make([]int, 0)
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func sliceToList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{nums[0], nil}
	cur := head
	for i := 1; i < len(nums); i++ {
		cur.Next = &ListNode{nums[i], nil}
		cur = cur.Next
	}
	return head
}

func main() {
	l1 := sliceToList([]int{1, 2, -3, 3, 1})
	fmt.Printf("%v (expected: [3 1] or [1 2 1])\n", listToSlice(removeZeroSumSublists(l1)))

	l2 := sliceToList([]int{1, 2, 3, -3, 4})
	fmt.Printf("%v (expected: [1 2 4])\n", listToSlice(removeZeroSumSublists(l2)))

	l3 := sliceToList([]int{1, -1})
	fmt.Printf("%v (expected: [])\n", listToSlice(removeZeroSumSublists(l3)))
}
```

## 1174 — Immediate Food Delivery Ii

```go
package main

import (
	"fmt"
	"math"
)

// LeetCode #1174: Immediate Food Delivery II
// https://leetcode.com/problems/immediate-food-delivery-ii/
// Difficulty: Medium

// For each customer's first order, find percentage that is "immediate"
// (order_date == customer_pref_delivery_date).

// Time: O(n)
// Space: O(n)

type delivery struct {
	ID         int
	customerID int
	orderDate  int
	prefDate   int
}

func immediateFoodDeliveryII(deliveries []delivery) float64 {
	// For each customer, track first order (earliest orderDate, tie-break by smallest ID)
	type firstOrder struct {
		date     int
		prefDate int
	}
	custFirst := make(map[int]firstOrder)

	for _, d := range deliveries {
		fo, exists := custFirst[d.customerID]
		if !exists || d.orderDate < fo.date {
			custFirst[d.customerID] = firstOrder{d.orderDate, d.prefDate}
		}
	}

	immediate := 0
	total := 0
	for _, fo := range custFirst {
		total++
		if fo.date == fo.prefDate {
			immediate++
		}
	}

	if total == 0 {
		return 0
	}
	return math.Round(float64(immediate)/float64(total)*100) / 100
}

func main() {
	d1 := []delivery{
		{1, 1, 20190801, 20190801},
		{2, 2, 20190802, 20190803},
		{3, 1, 20190803, 20190804},
	}
	fmt.Printf("%.2f (expected: 0.50)\n", immediateFoodDeliveryII(d1))

	d2 := []delivery{
		{1, 1, 20190801, 20190801},
		{2, 1, 20190802, 20190802},
	}
	fmt.Printf("%.2f (expected: 1.00)\n", immediateFoodDeliveryII(d2))
}
```

## 1177 — Can Make Palindrome From Substring

```go
package main

import (
	"fmt"
)

// LeetCode #1177: Can Make Palindrome from Substring
// https://leetcode.com/problems/can-make-palindrome-from-substring/
// Difficulty: Medium

// For each query [left, right, k], check if we can rearrange
// substring s[left..right] into a palindrome with at most k replacements.
// A palindrome can have at most 1 odd count character.
// We need floor(oddCount/2) <= k replacements.

// Time: O(n * 26 + m) where m = len(queries)
// Space: O(n * 26) — prefix sums per character

func canMakePaliQueries(s string, queries [][]int) []bool {
	n := len(s)
	// prefix[i][c] = count of char c in s[0:i]
	prefix := make([][26]int, n+1)

	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i]
		prefix[i+1][s[i]-'a']++
	}

	result := make([]bool, len(queries))
	for idx, q := range queries {
		l, r, k := q[0], q[1]+1, q[2]
		oddCount := 0
		for c := 0; c < 26; c++ {
			count := prefix[r][c] - prefix[l][c]
			if count%2 == 1 {
				oddCount++
			}
		}
		result[idx] = oddCount/2 <= k
	}
	return result
}

func main() {
	fmt.Printf("%v (expected: [true false])\n",
		canMakePaliQueries("abcda", [][]int{{3, 3, 0}, {1, 2, 0}}))

	fmt.Printf("%v (expected: [true])\n",
		canMakePaliQueries("abcda", [][]int{{0, 3, 1}}))

	fmt.Printf("%v (expected: [true true false])\n",
		canMakePaliQueries("abcdd", [][]int{{0, 4, 1}, {0, 2, 1}, {1, 2, 0}}))
}
```

## 1181 — Before And After Puzzle

```go
package main

import (
	"fmt"
	"sort"
	"strings"
)

// LeetCode #1181: Before and After Puzzle
// https://leetcode.com/problems/before-and-after-puzzle/
// Difficulty: Medium [Paid]

// Given list of phrases. Merge phrase i and j if last word of i
// equals first word of j. Result: i + j[firstWordLen:].
// Return sorted unique results.

// Time: O(n^2 * L)
// Space: O(n^2)

func beforeAndAfterPuzzles(phrases []string) []string {
	n := len(phrases)
	firstWords := make([]string, n)
	lastWords := make([]string, n)
	words := make([][]string, n)

	for i, p := range phrases {
		words[i] = strings.Fields(p)
		firstWords[i] = words[i][0]
		lastWords[i] = words[i][len(words[i])-1]
	}

	seen := make(map[string]bool)
	result := make([]string, 0)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			if lastWords[i] == firstWords[j] {
				merged := phrases[i] + phrases[j][len(firstWords[j]):]
				if !seen[merged] {
					seen[merged] = true
					result = append(result, merged)
				}
			}
		}
	}

	sort.Strings(result)
	return result
}

func main() {
	fmt.Printf("%v (expected: [writing code rocks])\n",
		beforeAndAfterPuzzles([]string{"writing code", "code rocks"}))

	fmt.Printf("%v (expected: [a d a b c d])\n",
		beforeAndAfterPuzzles([]string{"a b", "b c", "c d"}))
}
```

## 1182 — Shortest Distance To Target Color

```go
package main

import (
	"fmt"
	"math"
)

// LeetCode #1182: Shortest Distance to Target Color
// https://leetcode.com/problems/shortest-distance-to-target-color/
// Difficulty: Medium [Paid]

// Precompute nearest distance to each color from both directions.

// Time: O(n + m) where m = len(queries)
// Space: O(n)

func shortestDistanceColor(colors []int, queries [][]int) []int {
	n := len(colors)

	// left[i][c] = nearest distance to color c from left side up to i
	left := make([][3]int, n)
	for i := 0; i < n; i++ {
		for c := 0; c < 3; c++ {
			left[i][c] = math.MaxInt32
		}
	}

	for i := 0; i < n; i++ {
		if i > 0 {
			left[i] = left[i-1]
		}
		left[i][colors[i]-1] = 0
		// Update distances: all non-zero entries get +1
		for c := 0; c < 3; c++ {
			if left[i][c] != 0 && left[i][c] != math.MaxInt32 {
				// This gets complicated. Let's take a simpler approach.
			}
		}
	}

	// Simpler approach: for each color, store sorted positions
	positions := make([][]int, 4) // 1-indexed colors
	for i, c := range colors {
		positions[c] = append(positions[c], i)
	}

	binarySearch := func(pos []int, target int) int {
		lo, hi := 0, len(pos)-1
		if target <= pos[lo] {
			return pos[lo]
		}
		if target >= pos[hi] {
			return pos[hi]
		}
		for lo <= hi {
			mid := (lo + hi) / 2
			if pos[mid] == target {
				return target
			}
			if pos[mid] < target {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
		// lo is insertion point
		dist := abs(target - pos[lo])
		if lo > 0 && abs(target-pos[lo-1]) < dist {
			dist = abs(target - pos[lo-1])
			return pos[lo-1]
		}
		return pos[lo]
	}

	result := make([]int, len(queries))
	for i, q := range queries {
		idx, color := q[0], q[1]
		if len(positions[color]) == 0 {
			result[i] = -1
		} else {
			nearest := binarySearch(positions[color], idx)
			result[i] = abs(idx - nearest)
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

func main() {
	fmt.Printf("%v (expected: [3 0 3])\n",
		shortestDistanceColor([]int{1, 1, 2, 1, 3, 2, 2, 3, 3},
			[][]int{{1, 3}, {2, 2}, {6, 1}}))

	fmt.Printf("%v (expected: [-1])\n",
		shortestDistanceColor([]int{1, 2},
			[][]int{{0, 3}}))
}
```

## 1186 — Maximum Subarray Sum With One Deletion

```go
package main

import (
	"fmt"
)

// LeetCode #1186: Maximum Subarray Sum with One Deletion
// https://leetcode.com/problems/maximum-subarray-sum-with-one-deletion/
// Difficulty: Medium

// Kadane's algorithm variant with one deletion allowed.
// dp_no_del[i] = max subarray sum ending at i without deletion
// dp_del[i] = max subarray sum ending at i with one deletion

// Time: O(n)
// Space: O(1)

func maximumSum(arr []int) int {
	n := len(arr)
	if n == 0 {
		return 0
	}

	noDel := arr[0]
	withDel := arr[0]
	result := arr[0]

	for i := 1; i < n; i++ {
		withDel = max(withDel+arr[i], noDel)
		noDel = max(noDel+arr[i], arr[i])
		result = max(result, max(noDel, withDel))
	}

	return result
}

func main() {
	fmt.Printf("%d (expected: 4)\n", maximumSum([]int{1, -2, 0, 3}))
	fmt.Printf("%d (expected: -1)\n", maximumSum([]int{-1, -1, -1, -1}))
	fmt.Printf("%d (expected: 7)\n", maximumSum([]int{1, -2, -2, 3, -1, 4}))
}
```

## 1188 — Design Bounded Blocking Queue

```go
package main

import (
	"fmt"
	"sync"
)

// LeetCode #1188: Design Bounded Blocking Queue
// https://leetcode.com/problems/design-bounded-blocking-queue/
// Difficulty: Medium [Paid]

// Thread-safe bounded queue using mutex and condition variable.

// Time: O(1) per operation
// Space: O(capacity)

type BoundedBlockingQueue struct {
	capacity int
	queue    []int
	mu       sync.Mutex
	notFull  *sync.Cond
	notEmpty *sync.Cond
}

func NewBoundedBlockingQueue(capacity int) *BoundedBlockingQueue {
	bq := &BoundedBlockingQueue{
		capacity: capacity,
		queue:    make([]int, 0),
	}
	bq.notFull = sync.NewCond(&bq.mu)
	bq.notEmpty = sync.NewCond(&bq.mu)
	return bq
}

func (bq *BoundedBlockingQueue) Enqueue(element int) {
	bq.mu.Lock()
	defer bq.mu.Unlock()

	for len(bq.queue) >= bq.capacity {
		bq.notFull.Wait()
	}
	bq.queue = append(bq.queue, element)
	bq.notEmpty.Signal()
}

func (bq *BoundedBlockingQueue) Dequeue() int {
	bq.mu.Lock()
	defer bq.mu.Unlock()

	for len(bq.queue) == 0 {
		bq.notEmpty.Wait()
	}
	val := bq.queue[0]
	bq.queue = bq.queue[1:]
	bq.notFull.Signal()
	return val
}

func (bq *BoundedBlockingQueue) Size() int {
	bq.mu.Lock()
	defer bq.mu.Unlock()
	return len(bq.queue)
}

func main() {
	bq := NewBoundedBlockingQueue(2)
	bq.Enqueue(1)
	bq.Enqueue(2)
	fmt.Printf("size: %d (expected: 2)\n", bq.Size())
	fmt.Printf("dequeue: %d (expected: 1)\n", bq.Dequeue())
	fmt.Printf("size: %d (expected: 1)\n", bq.Size())
	bq.Enqueue(3)
	fmt.Printf("dequeue: %d (expected: 2)\n", bq.Dequeue())
	fmt.Printf("dequeue: %d (expected: 3)\n", bq.Dequeue())
}
```

## 1190 — Reverse Substrings Between Each Pair Of Parentheses

```go
package main

import (
	"fmt"
)

// LeetCode #1190: Reverse Substrings Between Each Pair of Parentheses
// https://leetcode.com/problems/reverse-substrings-between-each-pair-of-parentheses/
// Difficulty: Medium

// Find matching parentheses, then process from outside-in.
// Use wormhole approach: when hitting '(', jump to matching ')' and reverse direction.

// Time: O(n)
// Space: O(n)

func reverseParentheses(s string) string {
	n := len(s)
	pair := make([]int, n)
	stack := make([]int, 0)

	for i, ch := range s {
		if ch == '(' {
			stack = append(stack, i)
		} else if ch == ')' {
			open := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			pair[open] = i
			pair[i] = open
		}
	}

	result := make([]byte, 0, n)
	dir := 1
	for i := 0; i < n; i += dir {
		if s[i] == '(' || s[i] == ')' {
			i = pair[i]
			dir = -dir
		} else {
			result = append(result, s[i])
		}
	}

	return string(result)
}

func main() {
	fmt.Printf("%q (expected: %q)\n", reverseParentheses("(abcd)"), "dcba")
	fmt.Printf("%q (expected: %q)\n", reverseParentheses("(u(love)i)"), "iloveu")
	fmt.Printf("%q (expected: %q)\n", reverseParentheses("(ed(et(oc))el)"), "leetcode")
}
```

## 1191 — K Concatenation Maximum Sum

```go
package main

import (
	"fmt"
)

// LeetCode #1191: K-Concatenation Maximum Sum
// https://leetcode.com/problems/k-concatenation-maximum-sum/
// Difficulty: Medium

// If k == 1, just Kadane. If k >= 2, compute max suffix + max prefix +
// (k-2)*total if total > 0.

// Time: O(n)
// Space: O(1)

const mod = 1_000_000_007

func kConcatenationMaxSum(arr []int, k int) int {
	n := len(arr)
	if n == 0 {
		return 0
	}

	// Kadane on single array
	kadane := func(nums []int) int64 {
		var maxEnd, maxSoFar int64
		for _, v := range nums {
			maxEnd = max64(maxEnd+int64(v), int64(v))
			if maxEnd > maxSoFar {
				maxSoFar = maxEnd
			}
		}
		return maxSoFar
	}

	singleMax := kadane(arr)

	if k == 1 {
		return int(singleMax % mod)
	}

	// Total sum
	var total int64
	for _, v := range arr {
		total += int64(v)
	}

	// Max prefix sum
	var prefixSum, maxPrefix int64
	for _, v := range arr {
		prefixSum += int64(v)
		if prefixSum > maxPrefix {
			maxPrefix = prefixSum
		}
	}

	// Max suffix sum
	var suffixSum, maxSuffix int64
	for i := n - 1; i >= 0; i-- {
		suffixSum += int64(arr[i])
		if suffixSum > maxSuffix {
			maxSuffix = suffixSum
		}
	}

	doubleMax := kadane(append(arr, arr...))
	if k == 2 {
		if doubleMax > singleMax {
			return int(doubleMax % mod)
		}
		return int(singleMax % mod)
	}

	// k >= 3
	var result int64
	result = max64(singleMax, max64(doubleMax, maxPrefix+maxSuffix+total*int64(k-2)))
	// Also consider just maxPrefix + maxSuffix + total * (k-2), but may overflow
	result = max64(result, maxPrefix+maxSuffix)

	if result < 0 {
		return 0
	}
	return int(result % mod)
}

func max64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Printf("%d (expected: 9)\n", kConcatenationMaxSum([]int{1, -2, 1}, 5))
	fmt.Printf("%d (expected: 2)\n", kConcatenationMaxSum([]int{-1, -2}, 7))
	fmt.Printf("%d (expected: 4)\n", kConcatenationMaxSum([]int{1, 2}, 1))
}
```

## 1193 — Monthly Transactions I

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #1193: Monthly Transactions I
// https://leetcode.com/problems/monthly-transactions-i/
// Difficulty: Medium

// Group transactions by country and month, count approved and total.

// Time: O(n log n)
// Space: O(n)

type trans struct {
	id      int
	country string
	state   string
	amount  int
	date    string
}

type monthlyStat struct {
	month         string
	country       string
	approvedCount int
	approvedTotal int
	totalCount    int
	totalAmount   int
}

func monthlyTransactionsI(transactions []trans) []monthlyStat {
	type key struct {
		month   string
		country string
	}
	stats := make(map[key]monthlyStat)

	for _, t := range transactions {
		month := t.date[:7]
		k := key{month, t.country}
		s := stats[k]
		s.month = month
		s.country = t.country
		s.totalCount++
		s.totalAmount += t.amount
		if t.state == "approved" {
			s.approvedCount++
			s.approvedTotal += t.amount
		}
		stats[k] = s
	}

	result := make([]monthlyStat, 0, len(stats))
	for _, s := range stats {
		result = append(result, s)
	}

	sort.Slice(result, func(i, j int) bool {
		if result[i].month != result[j].month {
			return result[i].month < result[j].month
		}
		return result[i].country < result[j].country
	})

	return result
}

func main() {
	transactions := []trans{
		{1, "US", "approved", 1000, "2018-12-01"},
		{2, "US", "declined", 2000, "2018-12-02"},
		{3, "US", "approved", 3000, "2019-01-01"},
		{4, "DE", "approved", 4000, "2019-01-02"},
	}
	result := monthlyTransactionsI(transactions)
	for _, r := range result {
		fmt.Printf("%s/%s: approved=%d/%d total=%d/%d\n",
			r.month, r.country, r.approvedCount, r.approvedTotal, r.totalCount, r.totalAmount)
	}
}
```

## 1195 — Fizz Buzz Multithreaded

```go
package main

import (
	"fmt"
	"sync"
)

// LeetCode #1195: Fizz Buzz Multithreaded
// https://leetcode.com/problems/fizz-buzz-multithreaded/
// Difficulty: Medium

// Multithreaded FizzBuzz using goroutines and WaitGroup.
// Each goroutine prints one line at a time.

// Time: O(n)
// Space: O(n)

type FizzBuzz struct {
	n    int
	done chan struct{}
}

func NewFizzBuzz(n int) *FizzBuzz {
	return &FizzBuzz{n: n, done: make(chan struct{})}
}

func (fb *FizzBuzz) Start() []string {
	result := make([]string, fb.n)
	var wg sync.WaitGroup
	wg.Add(4)

	// fizzbuzz goroutine
	go func() {
		defer wg.Done()
		for i := 1; i <= fb.n; i++ {
			if i%15 == 0 {
				result[i-1] = "FizzBuzz"
			}
		}
	}()

	// fizz goroutine
	go func() {
		defer wg.Done()
		for i := 1; i <= fb.n; i++ {
			if i%3 == 0 && i%5 != 0 {
				result[i-1] = "Fizz"
			}
		}
	}()

	// buzz goroutine
	go func() {
		defer wg.Done()
		for i := 1; i <= fb.n; i++ {
			if i%5 == 0 && i%3 != 0 {
				result[i-1] = "Buzz"
			}
		}
	}()

	// number goroutine
	go func() {
		defer wg.Done()
		for i := 1; i <= fb.n; i++ {
			if i%3 != 0 && i%5 != 0 {
				result[i-1] = fmt.Sprintf("%d", i)
			}
		}
	}()

	wg.Wait()
	return result
}

func main() {
	fb := NewFizzBuzz(15)
	result := fb.Start()
	fmt.Printf("%v\n", result)
}
```

## 1197 — Minimum Knight Moves

```go
package main

import (
	"fmt"
)

// LeetCode #1197: Minimum Knight Moves
// https://leetcode.com/problems/minimum-knight-moves/
// Difficulty: Medium [Paid]

// BFS from (0,0) to (x,y). Use symmetry: 8 symmetric quadrants.
// Constrain BFS to non-negative coordinates for efficiency.

// Time: O(|x|*|y|) worst case
// Space: O(|x|*|y|)

func minKnightMoves(x int, y int) int {
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	if x < y {
		x, y = y, x
	}

	// BFS
	dirs := [][]int{{2, 1}, {1, 2}, {-1, 2}, {-2, 1},
		{-2, -1}, {-1, -2}, {1, -2}, {2, -1}}

	visited := make(map[[2]int]bool)
	queue := [][2]int{{0, 0}}
	visited[[2]int{0, 0}] = true
	steps := 0

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			if cur[0] == x && cur[1] == y {
				return steps
			}
			for _, d := range dirs {
				nx, ny := cur[0]+d[0], cur[1]+d[1]
				// Prune: only explore positive but not too far
				if nx >= -2 && ny >= -2 && nx <= x+2 && ny <= y+2 {
					key := [2]int{nx, ny}
					if !visited[key] {
						visited[key] = true
						queue = append(queue, key)
					}
				}
			}
		}
		steps++
	}

	return -1
}

func main() {
	fmt.Printf("%d (expected: 1)\n", minKnightMoves(2, 1))
	fmt.Printf("%d (expected: 2)\n", minKnightMoves(5, 5))
	fmt.Printf("%d (expected: 0)\n", minKnightMoves(0, 0))
}
```

## 1198 — Find Smallest Common Element In All Rows

```go
package main

import (
	"fmt"
)

// LeetCode #1198: Find Smallest Common Element in All Rows
// https://leetcode.com/problems/find-smallest-common-element-in-all-rows/
// Difficulty: Medium [Paid]

// Find smallest integer that appears in every row.

// Time: O(m * n)
// Space: O(max value) = O(10000) since values are 1..10000

func smallestCommonElement(mat [][]int) int {
	if len(mat) == 0 {
		return -1
	}

	count := make([]int, 10001)
	for _, v := range mat[0] {
		count[v] = 1
	}

	for i := 1; i < len(mat); i++ {
		for _, v := range mat[i] {
			if count[v] == i {
				count[v]++
			}
		}
	}

	for v := 1; v <= 10000; v++ {
		if count[v] == len(mat) {
			return v
		}
	}
	return -1
}

func main() {
	fmt.Printf("%d (expected: 5)\n",
		smallestCommonElement([][]int{{1, 2, 3, 4, 5}, {2, 4, 5, 8, 10}, {3, 5, 7, 9, 11}, {1, 3, 5, 7, 9}}))

	fmt.Printf("%d (expected: -1)\n",
		smallestCommonElement([][]int{{1, 2, 3}, {4, 5, 6}}))

	fmt.Printf("%d (expected: 2)\n",
		smallestCommonElement([][]int{{2, 3}, {2, 5}}))
}
```

## 1201 — Ugly Number Iii

```go
package main

import (
	"fmt"
)

// LeetCode #1201: Ugly Number III
// https://leetcode.com/problems/ugly-number-iii/
// Difficulty: Medium

// Find the nth number divisible by a, b, or c.
// Binary search on answer + inclusion-exclusion principle.

// Time: O(log(maxVal)) ~ O(log(2*10^9))
// Space: O(1)

func nthUglyNumber(n int, a int, b int, c int) int {
	// LCM helpers
	gcd := func(x, y int) int {
		for y != 0 {
			x, y = y, x%y
		}
		return x
	}
	lcm := func(x, y int) int {
		return x / gcd(x, y) * y
	}

	ab := lcm(a, b)
	bc := lcm(b, c)
	ac := lcm(a, c)
	abc := lcm(a, lcm(b, c))

	countUgly := func(num int) int {
		return num/a + num/b + num/c - num/ab - num/ac - num/bc + num/abc
	}

	lo, hi := 1, 2000000000
	for lo < hi {
		mid := lo + (hi-lo)/2
		if countUgly(mid) >= n {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func main() {
	fmt.Printf("%d (expected: 4)\n", nthUglyNumber(3, 2, 3, 5))
	fmt.Printf("%d (expected: 6)\n", nthUglyNumber(4, 2, 3, 4))
	fmt.Printf("%d (expected: 10)\n", nthUglyNumber(5, 2, 11, 13))
}
```

## 1202 — Smallest String With Swaps

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #1202: Smallest String With Swaps
// https://leetcode.com/problems/smallest-string-with-swaps/
// Difficulty: Medium

// Union-Find on indices. Indices in same connected component can be
// rearranged arbitrarily. Sort chars within each component.

// Time: O(n log n)
// Space: O(n)

type uf struct {
	parent []int
	rank   []int
}

func newUF(n int) *uf {
	p := make([]int, n)
	r := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
	}
	return &uf{p, r}
}

func (u *uf) find(x int) int {
	if u.parent[x] != x {
		u.parent[x] = u.find(u.parent[x])
	}
	return u.parent[x]
}

func (u *uf) union(x, y int) {
	x, y = u.find(x), u.find(y)
	if x == y {
		return
	}
	if u.rank[x] < u.rank[y] {
		x, y = y, x
	}
	u.parent[y] = x
	if u.rank[x] == u.rank[y] {
		u.rank[x]++
	}
}

func smallestStringWithSwaps(s string, pairs [][]int) string {
	n := len(s)
	u := newUF(n)
	for _, p := range pairs {
		u.union(p[0], p[1])
	}

	// Group indices by root
	groups := make(map[int][]int)
	for i := 0; i < n; i++ {
		root := u.find(i)
		groups[root] = append(groups[root], i)
	}

	result := make([]byte, n)
	for _, indices := range groups {
		chars := make([]byte, len(indices))
		for i, idx := range indices {
			chars[i] = s[idx]
		}
		sort.Slice(chars, func(i, j int) bool { return chars[i] < chars[j] })

		sort.Ints(indices)
		for i, idx := range indices {
			result[idx] = chars[i]
		}
	}

	return string(result)
}

func main() {
	fmt.Printf("%q (expected: %q)\n",
		smallestStringWithSwaps("dcab", [][]int{{0, 3}, {1, 2}}), "bacd")

	fmt.Printf("%q (expected: %q)\n",
		smallestStringWithSwaps("dcab", [][]int{{0, 3}, {1, 2}, {0, 2}}), "abcd")
}
```

## 1204 — Last Person To Fit In The Bus

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #1204: Last Person to Fit in the Bus
// https://leetcode.com/problems/last-person-to-fit-in-the-bus/
// Difficulty: Medium

// People queue for bus with weight limit 1000.
// Find the last person name that can board without exceeding limit.

// Time: O(n log n)
// Space: O(n)

type person struct {
	name   string
	weight int
	turn   int
}

func lastToFit(people []person) string {
	sort.Slice(people, func(i, j int) bool {
		return people[i].turn < people[j].turn
	})

	total := 0
	lastName := ""
	for _, p := range people {
		if total+p.weight <= 1000 {
			total += p.weight
			lastName = p.name
		} else {
			break
		}
	}
	return lastName
}

func main() {
	people := []person{
		{"Alice", 200, 1},
		{"Bob", 300, 2},
		{"Charlie", 400, 3},
		{"Dave", 200, 4},
	}
	fmt.Printf("%q (expected: \"Charlie\")\n", lastToFit(people))

	people2 := []person{
		{"John", 500, 1},
		{"Jane", 600, 2},
	}
	fmt.Printf("%q (expected: \"John\")\n", lastToFit(people2))
}
```

## 1205 — Monthly Transactions Ii

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #1205: Monthly Transactions II
// https://leetcode.com/problems/monthly-transactions-ii/
// Difficulty: Medium [Paid]

// For each month and country: approved_count, approved_amount,
// chargeback_count, chargeback_amount.

// Time: O(n log n)
// Space: O(n)

type tx struct {
	id          int
	country     string
	state       string
	amount      int
	transDate   string
	chargebackDate string
}

type monthlyStat struct {
	month            string
	country          string
	approvedCount    int
	approvedAmount   int
	chargebackCount  int
	chargebackAmount int
}

func monthlyTransactionsII(transactions []tx) []monthlyStat {
	type key struct {
		month   string
		country string
	}
	stats := make(map[key]monthlyStat)

	for _, t := range transactions {
		if t.state == "approved" || t.state == "declined" {
			month := t.transDate[:7]
			k := key{month, t.country}
			s := stats[k]
			s.month = month
			s.country = t.country
			if t.state == "approved" {
				s.approvedCount++
				s.approvedAmount += t.amount
			}
			stats[k] = s
		}
		if t.chargebackDate != "" {
			month := t.chargebackDate[:7]
			k := key{month, t.country}
			s := stats[k]
			s.month = month
			s.country = t.country
			s.chargebackCount++
			s.chargebackAmount += t.amount
			stats[k] = s
		}
	}

	result := make([]monthlyStat, 0, len(stats))
	for _, s := range stats {
		result = append(result, s)
	}
	sort.Slice(result, func(i, j int) bool {
		if result[i].month != result[j].month {
			return result[i].month < result[j].month
		}
		return result[i].country < result[j].country
	})

	return result
}

func main() {
	txs := []tx{
		{1, "US", "approved", 1000, "2018-12-01", ""},
		{2, "US", "declined", 2000, "2018-12-02", ""},
		{3, "US", "approved", 2000, "2019-01-01", "2019-01-05"},
	}
	result := monthlyTransactionsII(txs)
	for _, r := range result {
		fmt.Printf("month=%s country=%s approved=%d/%d chargeback=%d/%d\n",
			r.month, r.country, r.approvedCount, r.approvedAmount, r.chargebackCount, r.chargebackAmount)
	}
}
```

## 1208 — Get Equal Substrings Within Budget

```go
package main

import (
	"fmt"
)

// LeetCode #1208: Get Equal Substrings Within Budget
// https://leetcode.com/problems/get-equal-substrings-within-budget/
// Difficulty: Medium

// Sliding window: maxCost - cost[i] = difference between s[i] and t[i].
// Find longest substring with total cost <= maxCost.

// Time: O(n)
// Space: O(1)

func equalSubstring(s string, t string, maxCost int) int {
	n := len(s)
	cost := make([]int, n)
	for i := 0; i < n; i++ {
		c := int(s[i]) - int(t[i])
		if c < 0 {
			c = -c
		}
		cost[i] = c
	}

	left := 0
	currCost := 0
	maxLen := 0

	for right := 0; right < n; right++ {
		currCost += cost[right]
		for currCost > maxCost {
			currCost -= cost[left]
			left++
		}
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}

	return maxLen
}

func main() {
	fmt.Printf("%d (expected: 3)\n", equalSubstring("abcd", "bcdf", 3))
	fmt.Printf("%d (expected: 1)\n", equalSubstring("abcd", "cdef", 3))
	fmt.Printf("%d (expected: 1)\n", equalSubstring("abcd", "acde", 0))
}
```

## 1209 — Remove All Adjacent Duplicates In String Ii

```go
package main

import (
	"fmt"
)

// LeetCode #1209: Remove All Adjacent Duplicates in String II
// https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string-ii/
// Difficulty: Medium

// Use stack of (char, count). When count reaches k, pop.

// Time: O(n)
// Space: O(n)

func removeDuplicates(s string, k int) string {
	type pair struct {
		char  byte
		count int
	}
	stack := make([]pair, 0)

	for i := 0; i < len(s); i++ {
		if len(stack) > 0 && stack[len(stack)-1].char == s[i] {
			stack[len(stack)-1].count++
			if stack[len(stack)-1].count == k {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, pair{s[i], 1})
		}
	}

	result := make([]byte, 0)
	for _, p := range stack {
		for j := 0; j < p.count; j++ {
			result = append(result, p.char)
		}
	}
	return string(result)
}

func main() {
	fmt.Printf("%q (expected: %q)\n", removeDuplicates("abcd", 2), "abcd")
	fmt.Printf("%q (expected: %q)\n", removeDuplicates("deeedbbcccbdaa", 3), "aa")
	fmt.Printf("%q (expected: %q)\n", removeDuplicates("pbbcggttciiippooaais", 2), "ps")
}
```

## 1212 — Team Scores In Football Tournament

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #1212: Team Scores in Football Tournament
// https://leetcode.com/problems/team-scores-in-football-tournament/
// Difficulty: Medium [Paid]

// Teams get 3 pts for win, 1 for draw, 0 for loss.
// Calculate total points per team, sorted by points desc then name asc.

// Time: O(n log n)
// Space: O(n)

type matchResult struct {
	host       string
	guest      string
	hostScore  int
	guestScore int
}

type teamScore struct {
	name   string
	points int
}

func calculateTeamScores(teams []string, results []matchResult) []teamScore {
	points := make(map[string]int)

	for _, r := range results {
		if r.hostScore > r.guestScore {
			points[r.host] += 3
		} else if r.hostScore < r.guestScore {
			points[r.guest] += 3
		} else {
			points[r.host] += 1
			points[r.guest] += 1
		}
	}

	scores := make([]teamScore, 0, len(teams))
	for _, t := range teams {
		scores = append(scores, teamScore{t, points[t]})
	}

	sort.Slice(scores, func(i, j int) bool {
		if scores[i].points != scores[j].points {
			return scores[i].points > scores[j].points
		}
		return scores[i].name < scores[j].name
	})

	return scores
}

func main() {
	results := []matchResult{
		{"A", "B", 2, 1},
		{"C", "A", 1, 1},
		{"B", "C", 0, 3},
	}
	scores := calculateTeamScores([]string{"A", "B", "C"}, results)
	for _, s := range scores {
		fmt.Printf("%s: %d\n", s.name, s.points)
	}
}
```

## 1214 — Two Sum Bsts

```go
package main

import (
	"fmt"
)

// LeetCode #1214: Two Sum BSTs
// https://leetcode.com/problems/two-sum-bsts/
// Difficulty: Medium [Paid]

// Given two BSTs and a target, return true if there exists
// a node from each tree whose values sum to target.

// Time: O(n + m)
// Space: O(n + m)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func twoSumBSTs(root1 *TreeNode, root2 *TreeNode, target int) bool {
	vals := make(map[int]bool)

	var collect func(node *TreeNode)
	collect = func(node *TreeNode) {
		if node == nil {
			return
		}
		vals[node.Val] = true
		collect(node.Left)
		collect(node.Right)
	}
	collect(root1)

	var find func(node *TreeNode) bool
	find = func(node *TreeNode) bool {
		if node == nil {
			return false
		}
		if vals[target-node.Val] {
			return true
		}
		return find(node.Left) || find(node.Right)
	}
	return find(root2)
}

func main() {
	// Tree1: [2,1,4], Tree2: [1,0,3], target=5
	r1 := &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 4}}
	r2 := &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 3}}
	fmt.Printf("%t (expected: true)\n", twoSumBSTs(r1, r2, 5))

	// target=10 -> false
	fmt.Printf("%t (expected: false)\n", twoSumBSTs(r1, r2, 10))
}
```

## 1215 — Stepping Numbers

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #1215: Stepping Numbers
// https://leetcode.com/problems/stepping-numbers/
// Difficulty: Medium [Paid]

// Stepping numbers are numbers where adjacent digits differ by 1.
// Return all stepping numbers in range [low, high] sorted.

// Time: O(2^n) where n = number of digits in high
// Space: O(2^n)

func countSteppingNumbers(low int, high int) []int {
	result := make([]int, 0)

	var dfs func(num int)
	dfs = func(num int) {
		if num > high {
			return
		}
		if num >= low {
			result = append(result, num)
		}
		lastDigit := num % 10
		if lastDigit > 0 {
			dfs(num*10 + (lastDigit - 1))
		}
		if lastDigit < 9 {
			dfs(num*10 + (lastDigit + 1))
		}
	}

	for i := 0; i <= 9; i++ {
		dfs(i)
	}

	sort.Ints(result)
	return result
}

func main() {
	fmt.Printf("%v (expected: [0 1 2 3 4 5 6 7 8 9 10 12])\n", countSteppingNumbers(0, 12))
	fmt.Printf("%v (expected: [10 12])\n", countSteppingNumbers(10, 14))
}
```

## 1218 — Longest Arithmetic Subsequence Of Given Difference

```go
package main

import (
	"fmt"
)

// LeetCode #1218: Longest Arithmetic Subsequence of Given Difference
// https://leetcode.com/problems/longest-arithmetic-subsequence-of-given-difference/
// Difficulty: Medium

// dp[x] = length of longest arithmetic subsequence ending with value x.
// dp[x] = dp[x-difference] + 1

// Time: O(n)
// Space: O(n)

func longestSubsequence(arr []int, difference int) int {
	dp := make(map[int]int)
	maxLen := 0

	for _, v := range arr {
		prev := v - difference
		if count, exists := dp[prev]; exists {
			dp[v] = count + 1
		} else {
			dp[v] = 1
		}
		if dp[v] > maxLen {
			maxLen = dp[v]
		}
	}

	return maxLen
}

func main() {
	fmt.Printf("%d (expected: 4)\n", longestSubsequence([]int{1, 2, 3, 4}, 1))
	fmt.Printf("%d (expected: 1)\n", longestSubsequence([]int{1, 3, 5, 7}, 1))
	fmt.Printf("%d (expected: 4)\n", longestSubsequence([]int{1, 5, 7, 8, 5, 3, 4, 2, 1}, -2))
}
```

## 1219 — Path With Maximum Gold

```go
package main

import (
	"fmt"
)

// LeetCode #1219: Path with Maximum Gold
// https://leetcode.com/problems/path-with-maximum-gold/
// Difficulty: Medium

// DFS from each cell with gold, backtracking. Max gold collected.

// Time: O(m*n * 4^(k)) where k = max cells with gold
// Space: O(k) for recursion

func getMaximumGold(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	maxGold := 0

	var dfs func(r, c, gold int)
	dfs = func(r, c, gold int) {
		val := grid[r][c]
		gold += val
		if gold > maxGold {
			maxGold = gold
		}

		grid[r][c] = 0 // mark visited
		for _, d := range dirs {
			nr, nc := r+d[0], c+d[1]
			if nr >= 0 && nr < m && nc >= 0 && nc < n && grid[nr][nc] > 0 {
				dfs(nr, nc, gold)
			}
		}
		grid[r][c] = val // restore
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] > 0 {
				dfs(i, j, 0)
			}
		}
	}

	return maxGold
}

func main() {
	fmt.Printf("%d (expected: 24)\n",
		getMaximumGold([][]int{{0, 6, 0}, {5, 8, 7}, {0, 9, 0}}))

	fmt.Printf("%d (expected: 28)\n",
		getMaximumGold([][]int{{1, 0, 7}, {2, 0, 6}, {3, 4, 5}, {0, 3, 0}, {9, 0, 20}}))
}
```

## 1222 — Queens That Can Attack The King

```go
package main

import (
	"fmt"
)

// LeetCode #1222: Queens That Can Attack the King
// https://leetcode.com/problems/queens-that-can-attack-the-king/
// Difficulty: Medium

// Queens can attack the king if no other piece blocks.
// From king's position, check 8 directions for nearest queen.

// Time: O(n) where n = number of queens
// Space: O(1)

func queensAttacktheKing(queens [][]int, king []int) [][]int {
	queenSet := make(map[[2]int]bool)
	for _, q := range queens {
		queenSet[[2]int{q[0], q[1]}] = true
	}

	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0},
		{1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

	result := make([][]int, 0)

	for _, d := range dirs {
		r, c := king[0]+d[0], king[1]+d[1]
		for r >= 0 && r < 8 && c >= 0 && c < 8 {
			if queenSet[[2]int{r, c}] {
				result = append(result, []int{r, c})
				break
			}
			r += d[0]
			c += d[1]
		}
	}

	return result
}

func main() {
	fmt.Printf("%v (expected: [[0 1] [1 0] [3 3]])\n",
		queensAttacktheKing([][]int{{0, 1}, {1, 0}, {4, 0}, {0, 4}, {3, 3}, {2, 4}},
			[]int{0, 0}))

	fmt.Printf("%v (expected: [[2 2] [4 2]])\n",
		queensAttacktheKing([][]int{{0, 0}, {2, 2}, {4, 2}, {5, 5}},
			[]int{3, 2}))
}
```

## 1226 — The Dining Philosophers

```go
package main

import (
	"fmt"
	"sync"
)

// LeetCode #1226: The Dining Philosophers
// https://leetcode.com/problems/the-dining-philosophers/
// Difficulty: Medium

// 5 philosophers, 5 forks. Each needs 2 forks to eat.
// Deadlock avoidance: odd philosophers pick left then right,
// even philosophers pick right then left.

// Time: O(1) per eat call
// Space: O(1)

type DiningPhilosophers struct {
	mu     sync.Mutex
	forks  [5]sync.Mutex
}

func NewDiningPhilosophers() *DiningPhilosophers {
	return &DiningPhilosophers{}
}

func (dp *DiningPhilosophers) WantsToEat(philosopher int,
	eat func(),
	pickLeftFork func(),
	pickRightFork func(),
	putLeftFork func(),
	putRightFork func()) {

	left := philosopher
	right := (philosopher + 1) % 5

	// To avoid deadlock: always pick lower-numbered fork first
	if left < right {
		dp.forks[left].Lock()
		dp.forks[right].Lock()
	} else {
		dp.forks[right].Lock()
		dp.forks[left].Lock()
	}

	pickLeftFork()
	pickRightFork()
	eat()
	putLeftFork()
	putRightFork()

	dp.forks[left].Unlock()
	dp.forks[right].Unlock()
}

func main() {
	dp := NewDiningPhilosophers()
	var wg sync.WaitGroup

	eatCount := 0
	var mu sync.Mutex

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(p int) {
			defer wg.Done()
			for j := 0; j < 2; j++ {
				dp.WantsToEat(p,
					func() {
						mu.Lock()
						eatCount++
						mu.Unlock()
					},
					func() { fmt.Printf("P%d picks left\n", p) },
					func() { fmt.Printf("P%d picks right\n", p) },
					func() { fmt.Printf("P%d puts left\n", p) },
					func() { fmt.Printf("P%d puts right\n", p) },
				)
			}
		}(i)
	}
	wg.Wait()
	fmt.Printf("Total eats: %d (expected: 10)\n", eatCount)
}
```

## 1227 — Airplane Seat Assignment Probability

```go
package main

import (
	"fmt"
)

// LeetCode #1227: Airplane Seat Assignment Probability
// https://leetcode.com/problems/airplane-seat-assignment-probability/
// Difficulty: Medium

// n passengers board a plane with n seats. First passenger picks randomly.
// Others: if their seat is free, take it. Otherwise, pick random empty seat.
// Find probability that nth passenger gets their own seat.

// For n=1: 1.0
// For n>=2: 0.5

// Time: O(1)
// Space: O(1)

func nthPersonGetsNthSeat(n int) float64 {
	if n == 1 {
		return 1.0
	}
	return 0.5
}

func main() {
	fmt.Printf("%.2f (expected: 1.00)\n", nthPersonGetsNthSeat(1))
	fmt.Printf("%.2f (expected: 0.50)\n", nthPersonGetsNthSeat(2))
	fmt.Printf("%.2f (expected: 0.50)\n", nthPersonGetsNthSeat(100))
}
```

## 1229 — Meeting Scheduler

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #1229: Meeting Scheduler
// https://leetcode.com/problems/meeting-scheduler/
// Difficulty: Medium [Paid]

// Find earliest time slot of given duration that works for both.
// Two pointers approach after sorting slots by start time.

// Time: O(n log n + m log m)
// Space: O(1)

func minAvailableDuration(slots1 [][]int, slots2 [][]int, duration int) []int {
	sort.Slice(slots1, func(i, j int) bool { return slots1[i][0] < slots1[j][0] })
	sort.Slice(slots2, func(i, j int) bool { return slots2[i][0] < slots2[j][0] })

	i, j := 0, 0
	for i < len(slots1) && j < len(slots2) {
		start := max(slots1[i][0], slots2[j][0])
		end := min(slots1[i][1], slots2[j][1])
		if end-start >= duration {
			return []int{start, start + duration}
		}
		if slots1[i][1] < slots2[j][1] {
			i++
		} else {
			j++
		}
	}
	return []int{}
}

func main() {
	fmt.Printf("%v (expected: [60 68])\n",
		minAvailableDuration([][]int{{10, 50}, {60, 120}, {140, 210}},
			[][]int{{0, 15}, {60, 70}}, 8))

	fmt.Printf("%v (expected: [])\n",
		minAvailableDuration([][]int{{10, 50}, {60, 120}},
			[][]int{{0, 15}, {55, 58}}, 5))
}
```

## 1230 — Toss Strange Coins

```go
package main

import (
	"fmt"
)

// LeetCode #1230: Toss Strange Coins
// https://leetcode.com/problems/toss-strange-coins/
// Difficulty: Medium [Paid]

// Probability that exactly target coins land heads.
// dp[j] = probability of j heads after processing i coins.

// Time: O(n * target)
// Space: O(target)

func probabilityOfHeads(prob []float64, target int) float64 {
	n := len(prob)
	dp := make([]float64, target+1)
	dp[0] = 1.0

	for i := 0; i < n; i++ {
		for j := min(target, i+1); j >= 0; j-- {
			if j > 0 {
				dp[j] = dp[j-1]*prob[i] + dp[j]*(1-prob[i])
			} else {
				dp[0] = dp[0] * (1 - prob[i])
			}
		}
	}

	return dp[target]
}

func main() {
	fmt.Printf("%.5f (expected: 0.40000)\n",
		probabilityOfHeads([]float64{0.4}, 1))

	fmt.Printf("%.5f (expected: 0.40000)\n",
		probabilityOfHeads([]float64{0.5, 0.5, 0.5, 0.5, 0.5}, 0))
}
```

## 1233 — Remove Sub Folders From The Filesystem

```go
package main

import (
	"fmt"
	"sort"
	"strings"
)

// LeetCode #1233: Remove Sub-Folders from the Filesystem
// https://leetcode.com/problems/remove-sub-folders-from-the-filesystem/
// Difficulty: Medium

// Sort folders lexicographically. If a folder is a prefix of next,
// the next is a sub-folder. Add "/" to check exact prefix match.

// Time: O(n log n * L) where L = average path length
// Space: O(n)

func removeSubfolders(folder []string) []string {
	sort.Strings(folder)
	result := make([]string, 0)
	result = append(result, folder[0])

	for i := 1; i < len(folder); i++ {
		last := result[len(result)-1]
		if !strings.HasPrefix(folder[i], last+"/") {
			result = append(result, folder[i])
		}
	}

	return result
}

func main() {
	fmt.Printf("%v (expected: [/a /c/d])\n",
		removeSubfolders([]string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"}))

	fmt.Printf("%v (expected: [/a])\n",
		removeSubfolders([]string{"/a", "/a/b/c", "/a/b"}))

	fmt.Printf("%v (expected: [/a/b /c /d])\n",
		removeSubfolders([]string{"/a/b", "/c", "/d"}))
}
```

## 1234 — Replace The Substring For Balanced String

```go
package main

import (
	"fmt"
)

// LeetCode #1234: Replace the Substring for Balanced String
// https://leetcode.com/problems/replace-the-substring-for-balanced-string/
// Difficulty: Medium

// Sliding window. Find smallest substring such that outside it,
// each of Q,W,E,R appears at most n/4 times.

// Time: O(n)
// Space: O(1)

func balancedString(s string) int {
	n := len(s)
	target := n / 4
	count := make(map[byte]int)
	for i := 0; i < n; i++ {
		count[s[i]]++
	}

	// Check if already balanced
	balanced := true
	for _, c := range []byte{'Q', 'W', 'E', 'R'} {
		if count[c] > target {
			balanced = false
			break
		}
	}
	if balanced {
		return 0
	}

	left := 0
	minLen := n

	for right := 0; right < n; right++ {
		count[s[right]]--

		for left <= right {
			ok := true
			for _, c := range []byte{'Q', 'W', 'E', 'R'} {
				if count[c] > target {
					ok = false
					break
				}
			}
			if !ok {
				break
			}
			if right-left+1 < minLen {
				minLen = right - left + 1
			}
			count[s[left]]++
			left++
		}
	}

	return minLen
}

func main() {
	fmt.Printf("%d (expected: 0)\n", balancedString("QWER"))
	fmt.Printf("%d (expected: 1)\n", balancedString("QQWE"))
	fmt.Printf("%d (expected: 2)\n", balancedString("QQQW"))
}
```

## 1236 — Web Crawler

```go
package main

import (
	"fmt"
)

// LeetCode #1236: Web Crawler
// https://leetcode.com/problems/web-crawler/
// Difficulty: Medium [Paid]

// Given a startUrl and an HtmlParser, crawl all reachable URLs
// under the same hostname (same domain).

// Time: O(V + E) where V = #urls, E = #links
// Space: O(V)

type HtmlParser interface {
	GetUrls(url string) []string
}

type mockParser struct {
	urls map[string][]string
}

func (m *mockParser) GetUrls(url string) []string {
	return m.urls[url]
}

func crawl(startUrl string, parser HtmlParser) []string {
	// Extract hostname
	getHost := func(url string) string {
		// Skip protocol
		host := ""
		for i := 0; i < len(url)-7; i++ {
			if url[i:i+7] == "http://" {
				url = url[7:]
				break
			}
		}
		for _, c := range url {
			if c == '/' || c == ':' {
				break
			}
			host += string(c)
		}
		return host
	}

	hostname := getHost(startUrl)
	visited := make(map[string]bool)
	queue := []string{startUrl}
	visited[startUrl] = true

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		for _, next := range parser.GetUrls(cur) {
			if !visited[next] && getHost(next) == hostname {
				visited[next] = true
				queue = append(queue, next)
			}
		}
	}

	result := make([]string, 0, len(visited))
	for url := range visited {
		result = append(result, url)
	}
	return result
}

func main() {
	parser := &mockParser{
		urls: map[string][]string{
			"http://example.com":     {"http://example.com/about", "http://other.com"},
			"http://example.com/about": {"http://example.com/contact"},
			"http://example.com/contact": {},
		},
	}
	result := crawl("http://example.com", parser)
	fmt.Printf("%v\n", result)
}
```

## 1237 — Find Positive Integer Solution For A Given Equation

```go
package main

import (
	"fmt"
)

// LeetCode #1237: Find Positive Integer Solution for a Given Equation
// https://leetcode.com/problems/find-positive-integer-solution-for-a-given-equation/
// Difficulty: Medium

// Given custom function f(x,y) that's monotonically increasing in both x and y,
// find all [x,y] where f(x,y) = z.
// Start from x=1, y=1000 and move inwards.

// Time: O(x + y) where x,y in [1, 1000]
// Space: O(n) for result

type customFunction func(int, int) int

func findSolution(customfunction customFunction, z int) [][]int {
	result := make([][]int, 0)
	x, y := 1, 1000

	for x <= 1000 && y >= 1 {
		val := customfunction(x, y)
		if val == z {
			result = append(result, []int{x, y})
			x++
			y--
		} else if val < z {
			x++
		} else {
			y--
		}
	}

	return result
}

func main() {
	// f(x,y) = x + y
	f1 := func(x, y int) int { return x + y }
	fmt.Printf("%v (expected: [[1 4] [2 3] [3 2] [4 1]])\n", findSolution(f1, 5))

	// f(x,y) = x * y
	f2 := func(x, y int) int { return x * y }
	fmt.Printf("%v (expected: [[1 12] [2 6] [3 4] [4 3] [6 2] [12 1]])\n", findSolution(f2, 12))
}
```

## 1238 — Circular Permutation In Binary Representation

```go
package main

import (
	"fmt"
)

// LeetCode #1238: Circular Permutation in Binary Representation
// https://leetcode.com/problems/circular-permutation-in-binary-representation/
// Difficulty: Medium

// Generate Gray code sequence starting with start.
// Gray code: consecutive numbers differ by exactly 1 bit.

// Time: O(2^n)
// Space: O(2^n)

func circularPermutation(n int, start int) []int {
	size := 1 << n
	result := make([]int, size)

	for i := 0; i < size; i++ {
		result[i] = start ^ i ^ (i >> 1)
	}

	return result
}

func main() {
	fmt.Printf("%v (expected: [3 2 0 1] or similar)\n", circularPermutation(2, 3))
	fmt.Printf("%v\n", circularPermutation(3, 2))
}
```

## 1239 — Maximum Length Of A Concatenated String With Unique Characters

```go
package main

import (
	"fmt"
)

// LeetCode #1239: Maximum Length of a Concatenated String with Unique Characters
// https://leetcode.com/problems/maximum-length-of-a-concatenated-string-with-unique-characters/
// Difficulty: Medium

// Backtracking with bitmask. Each string is represented as a bitmask.
// Only concatenate if no character conflict.

// Time: O(2^n) worst case
// Space: O(n)

func maxLength(arr []string) int {
	masks := make([]int, 0)

	for _, s := range arr {
		mask := 0
		valid := true
		for _, ch := range s {
			bit := 1 << (ch - 'a')
			if mask&bit != 0 {
				valid = false
				break
			}
			mask |= bit
		}
		if valid {
			masks = append(masks, mask)
		}
	}

	maxLen := 0
	var backtrack func(idx, mask, length int)
	backtrack = func(idx, mask, length int) {
		if length > maxLen {
			maxLen = length
		}
		for i := idx; i < len(masks); i++ {
			if mask&masks[i] == 0 {
				backtrack(i+1, mask|masks[i], length+countBits(masks[i]))
			}
		}
	}

	backtrack(0, 0, 0)
	return maxLen
}

func countBits(mask int) int {
	count := 0
	for mask > 0 {
		count += mask & 1
		mask >>= 1
	}
	return count
}

func main() {
	fmt.Printf("%d (expected: 4)\n", maxLength([]string{"un", "iq", "ue"}))
	fmt.Printf("%d (expected: 6)\n", maxLength([]string{"cha", "r", "act", "ers"}))
	fmt.Printf("%d (expected: 26)\n", maxLength([]string{"abcdefghijklmnopqrstuvwxyz"}))
}
```

## 1242 — Web Crawler Multithreaded

```go
package main

import (
	"fmt"
	"sync"
)

// LeetCode #1242: Web Crawler Multithreaded
// https://leetcode.com/problems/web-crawler-multithreaded/
// Difficulty: Medium [Paid]

// Multithreaded web crawler using goroutines and channels.
// Same hostname constraint.

// Time: O(V + E)
// Space: O(V)

type HtmlParser interface {
	GetUrls(url string) []string
}

type mockParser struct {
	urls map[string][]string
}

func (m *mockParser) GetUrls(url string) []string {
	return m.urls[url]
}

func crawlParallel(startUrl string, parser HtmlParser) []string {
	getHost := func(url string) string {
		host := ""
		for i := 0; i < len(url)-7; i++ {
			if url[i:i+7] == "http://" {
				url = url[7:]
				break
			}
		}
		for _, c := range url {
			if c == '/' || c == ':' {
				break
			}
			host += string(c)
		}
		return host
	}

	hostname := getHost(startUrl)
	visited := make(map[string]bool)
	var mu sync.Mutex
	var wg sync.WaitGroup

	var crawl func(url string)
	crawl = func(url string) {
		defer wg.Done()
		urls := parser.GetUrls(url)
		for _, next := range urls {
			mu.Lock()
			if visited[next] || getHost(next) != hostname {
				mu.Unlock()
				continue
			}
			visited[next] = true
			mu.Unlock()

			wg.Add(1)
			go crawl(next)
		}
	}

	visited[startUrl] = true
	wg.Add(1)
	go crawl(startUrl)
	wg.Wait()

	result := make([]string, 0, len(visited))
	for url := range visited {
		result = append(result, url)
	}
	return result
}

func main() {
	parser := &mockParser{
		urls: map[string][]string{
			"http://example.com":       {"http://example.com/about", "http://other.com"},
			"http://example.com/about": {"http://example.com/contact"},
			"http://example.com/contact": {},
		},
	}
	result := crawlParallel("http://example.com", parser)
	fmt.Printf("%v\n", result)
}
```

## 1244 — Design A Leaderboard

```go
package main

import (
	"container/heap"
	"fmt"
)

// LeetCode #1244: Design A Leaderboard
// https://leetcode.com/problems/design-a-leaderboard/
// Difficulty: Medium [Paid]

// Leaderboard supports addScore, top(K), and reset.

// Time: addScore O(1), top O(K log n), reset O(1)
// Space: O(n)

type Leaderboard struct {
	scores map[int]int
}

func Constructor() Leaderboard {
	return Leaderboard{scores: make(map[int]int)}
}

func (lb *Leaderboard) AddScore(playerId int, score int) {
	lb.scores[playerId] += score
}

func (lb *Leaderboard) Top(K int) int {
	// Use min-heap of size K
	h := &minHeap{}
	heap.Init(h)
	for _, s := range lb.scores {
		heap.Push(h, s)
		if h.Len() > K {
			heap.Pop(h)
		}
	}
	sum := 0
	for h.Len() > 0 {
		sum += heap.Pop(h).(int)
	}
	return sum
}

func (lb *Leaderboard) Reset(playerId int) {
	delete(lb.scores, playerId)
}

type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func main() {
	lb := Constructor()
	lb.AddScore(1, 73)
	lb.AddScore(2, 56)
	lb.AddScore(3, 39)
	lb.AddScore(4, 51)
	lb.AddScore(5, 4)
	fmt.Printf("top(1) = %d (expected: 73)\n", lb.Top(1))
	lb.Reset(1)
	lb.Reset(2)
	lb.AddScore(2, 51)
	fmt.Printf("top(3) = %d\n", lb.Top(3))
}
```

## 1245 — Tree Diameter

```go
package main

import (
	"fmt"
)

// LeetCode #1245: Tree Diameter
// https://leetcode.com/problems/tree-diameter/
// Difficulty: Medium [Paid]

// Two BFS: find farthest node from any node, then farthest from that node.
// Distance = diameter.

// Time: O(n)
// Space: O(n)

func treeDiameter(edges [][]int) int {
	if len(edges) == 0 {
		return 0
	}

	n := len(edges) + 1
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	bfs := func(start int) (int, int) {
		dist := make([]int, n)
		for i := range dist {
			dist[i] = -1
		}
		queue := []int{start}
		dist[start] = 0
		farthestNode, maxDist := start, 0

		for len(queue) > 0 {
			u := queue[0]
			queue = queue[1:]
			for _, v := range adj[u] {
				if dist[v] == -1 {
					dist[v] = dist[u] + 1
					queue = append(queue, v)
					if dist[v] > maxDist {
						maxDist = dist[v]
						farthestNode = v
					}
				}
			}
		}
		return farthestNode, maxDist
	}

	farthest, _ := bfs(0)
	_, diameter := bfs(farthest)
	return diameter
}

func main() {
	fmt.Printf("%d (expected: 2)\n", treeDiameter([][]int{{0, 1}, {1, 2}, {2, 3}}))
	fmt.Printf("%d (expected: 3)\n", treeDiameter([][]int{{0, 1}, {0, 2}, {1, 3}}))
	fmt.Printf("%d (expected: 0)\n", treeDiameter([][]int{}))
}
```

## 1247 — Minimum Swaps To Make Strings Equal

```go
package main

import (
	"fmt"
)

// LeetCode #1247: Minimum Swaps to Make Strings Equal
// https://leetcode.com/problems/minimum-swaps-to-make-strings-equal/
// Difficulty: Medium

// Count positions where s1[i] != s2[i].
// If odd count -> not possible.
// Count patterns: x_y (s1=x, s2=y) and y_x.
// Answer = x_y/2 + y_x/2 + (x_y%2)*2

// Time: O(n)
// Space: O(1)

func minimumSwap(s1 string, s2 string) int {
	n := len(s1)
	if n != len(s2) {
		return -1
	}

	xy, yx := 0, 0
	for i := 0; i < n; i++ {
		if s1[i] == 'x' && s2[i] == 'y' {
			xy++
		} else if s1[i] == 'y' && s2[i] == 'x' {
			yx++
		}
	}

	if (xy+yx)%2 == 1 {
		return -1
	}

	return xy/2 + yx/2 + (xy%2)*2
}

func main() {
	fmt.Printf("%d (expected: 1)\n", minimumSwap("xy", "yx"))
	fmt.Printf("%d (expected: 2)\n", minimumSwap("xx", "yy"))
	fmt.Printf("%d (expected: -1)\n", minimumSwap("xy", "xx"))
}
```

## 1248 — Count Number Of Nice Subarrays

```go
package main

import (
	"fmt"
)

// LeetCode #1248: Count Number of Nice Subarrays
// https://leetcode.com/problems/count-number-of-nice-subarrays/
// Difficulty: Medium

// Count subarrays with exactly k odd numbers.
// Convert to atMost(k) - atMost(k-1).

// Time: O(n)
// Space: O(1)

func numberOfSubarrays(nums []int, k int) int {
	atMost := func(target int) int {
		if target < 0 {
			return 0
		}
		left, count, result := 0, 0, 0
		for right := 0; right < len(nums); right++ {
			if nums[right]%2 == 1 {
				count++
			}
			for count > target {
				if nums[left]%2 == 1 {
					count--
				}
				left++
			}
			result += right - left + 1
		}
		return result
	}

	return atMost(k) - atMost(k-1)
}

func main() {
	fmt.Printf("%d (expected: 2)\n", numberOfSubarrays([]int{1, 1, 2, 1, 1}, 3))
	fmt.Printf("%d (expected: 0)\n", numberOfSubarrays([]int{2, 4, 6}, 1))
	fmt.Printf("%d (expected: 16)\n", numberOfSubarrays([]int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2}, 2))
}
```

## 1249 — Minimum Remove To Make Valid Parentheses

```go
package main

import (
	"fmt"
)

// LeetCode #1249: Minimum Remove to Make Valid Parentheses
// https://leetcode.com/problems/minimum-remove-to-make-valid-parentheses/
// Difficulty: Medium

// First pass: remove unmatched ')'. Second pass: remove unmatched '('.

// Time: O(n)
// Space: O(n)

func minRemoveToMakeValid(s string) string {
	n := len(s)
	stack := make([]int, 0)
	remove := make([]bool, n)

	for i := 0; i < n; i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else if s[i] == ')' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			} else {
				remove[i] = true
			}
		}
	}

	for _, idx := range stack {
		remove[idx] = true
	}

	result := make([]byte, 0, n)
	for i := 0; i < n; i++ {
		if !remove[i] {
			result = append(result, s[i])
		}
	}
	return string(result)
}

func main() {
	fmt.Printf("%q (expected: %q)\n", minRemoveToMakeValid("lee(t(c)o)de)"), "lee(t(c)o)de")
	fmt.Printf("%q (expected: %q)\n", minRemoveToMakeValid("a)b(c)d"), "ab(c)d")
	fmt.Printf("%q (expected: %q)\n", minRemoveToMakeValid("))(("), "")
}
```

## 1253 — Reconstruct A 2 Row Binary Matrix

```go
package main

import (
	"fmt"
)

// LeetCode #1253: Reconstruct a 2-Row Binary Matrix
// https://leetcode.com/problems/reconstruct-a-2-row-binary-matrix/
// Difficulty: Medium

// colsum[i] = sum of col i (top + bottom).
// Fill col with 2 first (top=1, bottom=1), then 1s.
// Greedy: use top row capacity first.

// Time: O(n)
// Space: O(n)

func reconstructMatrix(upper int, lower int, colsum []int) [][]int {
	n := len(colsum)
	top := make([]int, n)
	bottom := make([]int, n)

	for i, s := range colsum {
		if s == 2 {
			top[i] = 1
			bottom[i] = 1
			upper--
			lower--
		}
	}

	if upper < 0 || lower < 0 {
		return [][]int{}
	}

	for i, s := range colsum {
		if s == 1 {
			if upper > 0 {
				top[i] = 1
				upper--
			} else if lower > 0 {
				bottom[i] = 1
				lower--
			} else {
				return [][]int{}
			}
		}
	}

	if upper != 0 || lower != 0 {
		return [][]int{}
	}

	return [][]int{top, bottom}
}

func main() {
	fmt.Printf("%v (expected: [[1 1 0 0] [0 0 1 1]])\n",
		reconstructMatrix(2, 2, []int{1, 1, 1, 1}))

	fmt.Printf("%v (expected: [[]])\n",
		reconstructMatrix(2, 1, []int{1, 1, 1}))

	fmt.Printf("%v\n",
		reconstructMatrix(5, 5, []int{2, 1, 2, 0, 1, 2}))
}
```

## 1254 — Number Of Closed Islands

```go
package main

import (
	"fmt"
)

// LeetCode #1254: Number of Closed Islands
// https://leetcode.com/problems/number-of-closed-islands/
// Difficulty: Medium

// A closed island is a group of 0s not connected to the border.
// First mark all border-connected 0s, then count remaining islands.

// Time: O(m*n)
// Space: O(m*n) recursion depth

func closedIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= m || c < 0 || c >= n || grid[r][c] != 0 {
			return
		}
		grid[r][c] = 1 // mark as water/visited
		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}

	// Mark border-connected 0s
	for i := 0; i < m; i++ {
		dfs(i, 0)
		dfs(i, n-1)
	}
	for j := 0; j < n; j++ {
		dfs(0, j)
		dfs(m-1, j)
	}

	count := 0
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if grid[i][j] == 0 {
				count++
				dfs(i, j)
			}
		}
	}

	return count
}

func main() {
	fmt.Printf("%d (expected: 2)\n",
		closedIsland([][]int{
			{1, 1, 1, 1, 1, 1, 1, 0},
			{1, 0, 0, 0, 0, 1, 1, 0},
			{1, 0, 1, 0, 1, 1, 1, 0},
			{1, 0, 0, 0, 0, 1, 0, 1},
			{1, 1, 1, 1, 1, 1, 1, 0},
		}))

	fmt.Printf("%d (expected: 0)\n",
		closedIsland([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}))
}
```

## 1256 — Encode Number

```go
package main

import (
	"fmt"
)

// LeetCode #1256: Encode Number
// https://leetcode.com/problems/encode-number/
// Difficulty: Medium [Paid]

// Encode n as binary of n+1, then remove first bit.
// n=0 -> "0" (binary of 1 -> "1", remove first -> "")
// Wait: n=0 -> "". Let me check the pattern.
// 0: "" (1->"1", drop first->"")
// 1: "0" (2->"10", drop first->"0")
// 2: "1" (3->"11", drop first->"1")
// 3: "00" (4->"100", drop first->"00")
// 4: "01" (5->"101", drop first->"01")

// Time: O(log n)
// Space: O(log n)

func encode(num int) string {
	if num == 0 {
		return ""
	}

	// num+1 in binary, then drop first bit
	n := num + 1
	result := ""

	for n > 1 {
		if n%2 == 0 {
			result = "0" + result
		} else {
			result = "1" + result
		}
		n /= 2
	}

	return result
}

func main() {
	fmt.Printf("%q (expected: %q)\n", encode(0), "")
	fmt.Printf("%q (expected: %q)\n", encode(1), "0")
	fmt.Printf("%q (expected: %q)\n", encode(2), "1")
	fmt.Printf("%q (expected: %q)\n", encode(3), "00")
}
```

## 1257 — Smallest Common Region

```go
package main

import (
	"fmt"
)

// LeetCode #1257: Smallest Common Region
// https://leetcode.com/problems/smallest-common-region/
// Difficulty: Medium [Paid]

// Given region hierarchy, find smallest common region.
// Build parent map, then find common ancestor.

// Time: O(n) where n = total regions
// Space: O(n)

func findSmallestRegion(regions [][]string, region1 string, region2 string) string {
	parent := make(map[string]string)

	for _, list := range regions {
		for i := 1; i < len(list); i++ {
			parent[list[i]] = list[0]
		}
	}

	// Find path from region1 to root
	path := make(map[string]bool)
	r := region1
	path[r] = true
	for {
		p, exists := parent[r]
		if !exists {
			break
		}
		path[p] = true
		r = p
	}

	// Find common ancestor
	r = region2
	for {
		if path[r] {
			return r
		}
		r = parent[r]
	}
}

func main() {
	regions := [][]string{
		{"Earth", "North America", "South America"},
		{"North America", "USA", "Canada"},
		{"USA", "California", "Texas"},
		{"Canada", "Ontario", "Quebec"},
	}
	fmt.Printf("%q (expected: %q)\n",
		findSmallestRegion(regions, "California", "Ontario"), "North America")

	fmt.Printf("%q (expected: %q)\n",
		findSmallestRegion(regions, "California", "Texas"), "USA")
}
```

## 1258 — Synonymous Sentences

```go
package main

import (
	"fmt"
	"sort"
	"strings"
)

// LeetCode #1258: Synonymous Sentences
// https://leetcode.com/problems/synonymous-sentences/
// Difficulty: Medium [Paid]

// Union-Find to group synonyms, then generate all sentences
// by replacing words with all synonyms in their group.

// Time: O(2^k) where k = number of synonym groups per sentence
// Space: O(n)

type uf struct {
	parent map[string]string
}

func newUF() *uf {
	return &uf{parent: make(map[string]string)}
}

func (u *uf) find(x string) string {
	if _, exists := u.parent[x]; !exists {
		u.parent[x] = x
	}
	if u.parent[x] != x {
		u.parent[x] = u.find(u.parent[x])
	}
	return u.parent[x]
}

func (u *uf) union(x, y string) {
	u.parent[u.find(x)] = u.find(y)
}

func generateSentences(synonyms [][]string, text string) []string {
	u := newUF()
	for _, s := range synonyms {
		u.union(s[0], s[1])
	}

	// Group words by root
	groups := make(map[string][]string)
	for _, s := range synonyms {
		for _, w := range s {
			root := u.find(w)
			groups[root] = append(groups[root], w)
		}
	}

	// Sort and dedupe each group
	for root := range groups {
		wordSet := make(map[string]bool)
		for _, w := range groups[root] {
			wordSet[w] = true
		}
		groups[root] = make([]string, 0, len(wordSet))
		for w := range wordSet {
			groups[root] = append(groups[root], w)
		}
		sort.Strings(groups[root])
	}

	// Map word -> group root
	wordToRoot := make(map[string]string)
	for root, words := range groups {
		for _, w := range words {
			wordToRoot[w] = root
		}
	}

	words := strings.Fields(text)
	result := make([]string, 0)

	var backtrack func(idx int, current []string)
	backtrack = func(idx int, current []string) {
		if idx == len(words) {
			result = append(result, strings.Join(current, " "))
			return
		}
		root, hasSynonyms := wordToRoot[words[idx]]
		if !hasSynonyms {
			backtrack(idx+1, append(current, words[idx]))
		} else {
			for _, syn := range groups[root] {
				backtrack(idx+1, append(current, syn))
			}
		}
	}

	backtrack(0, []string{})
	return result
}

func main() {
	result := generateSentences([][]string{{"happy", "joy"}, {"sad", "sorrow"}}, "I am happy today")
	fmt.Printf("%v\n", result)
}
```

## 1261 — Find Elements In A Contaminated Binary Tree

```go
package main

import (
	"fmt"
)

// LeetCode #1261: Find Elements in a Contaminated Binary Tree
// https://leetcode.com/problems/find-elements-in-a-contaminated-binary-tree/
// Difficulty: Medium

// Tree is contaminated (all vals = -1). Recover with rules:
// root.val = 0, left.val = 2*parent+1, right.val = 2*parent+2.
// Then support Find(target) operation.

// Time: O(n) init, O(1) find
// Space: O(n)

type FindElements struct {
	vals map[int]bool
}

func Constructor(root *TreeNode) FindElements {
	fe := FindElements{vals: make(map[int]bool)}
	fe.recover(root, 0)
	return fe
}

func (fe *FindElements) recover(node *TreeNode, val int) {
	if node == nil {
		return
	}
	node.Val = val
	fe.vals[val] = true
	fe.recover(node.Left, 2*val+1)
	fe.recover(node.Right, 2*val+2)
}

func (fe *FindElements) Find(target int) bool {
	return fe.vals[target]
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Tree: [-1,null,-1]
	root := &TreeNode{Val: -1, Right: &TreeNode{Val: -1}}
	fe := Constructor(root)
	fmt.Printf("%t (expected: false)\n", fe.Find(1))
	fmt.Printf("%t (expected: true)\n", fe.Find(2))

	// Tree: [-1,-1,-1,-1,-1]
	root2 := &TreeNode{
		Val: -1,
		Left: &TreeNode{
			Val:   -1,
			Left:  &TreeNode{Val: -1},
			Right: &TreeNode{Val: -1},
		},
		Right: &TreeNode{Val: -1},
	}
	fe2 := Constructor(root2)
	fmt.Printf("%t (expected: true)\n", fe2.Find(1))
	fmt.Printf("%t (expected: true)\n", fe2.Find(3))
	fmt.Printf("%t (expected: false)\n", fe2.Find(5))
}
```

## 1262 — Greatest Sum Divisible By Three

```go
package main

import (
	"fmt"
)

// LeetCode #1262: Greatest Sum Divisible by Three
// https://leetcode.com/problems/greatest-sum-divisible-by-three/
// Difficulty: Medium

// DP with 3 states: max sum with remainder 0, 1, 2.

// Time: O(n)
// Space: O(1)

func maxSumDivThree(nums []int) int {
	dp := [3]int{0, -1, -1}

	for _, v := range nums {
		next := dp
		for r := 0; r < 3; r++ {
			if dp[r] != -1 {
				nr := (r + v) % 3
				if dp[r]+v > next[nr] {
					next[nr] = dp[r] + v
				}
			}
		}
		dp = next
	}

	return dp[0]
}

func main() {
	fmt.Printf("%d (expected: 18)\n", maxSumDivThree([]int{3, 6, 5, 1, 8}))
	fmt.Printf("%d (expected: 0)\n", maxSumDivThree([]int{4}))
	fmt.Printf("%d (expected: 12)\n", maxSumDivThree([]int{1, 2, 3, 4, 4}))
}
```

