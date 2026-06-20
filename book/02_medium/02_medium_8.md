# Medium (Sedang) — Problem 1264–1462

## 1264 — Page Recommendations

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #1264: Page Recommendations
// https://leetcode.com/problems/page-recommendations/
// Difficulty: Medium [Paid]

// Recommend pages liked by friends of user1 but not liked by user1.

// Time: O(n log n)
// Space: O(n)

type like struct {
	userID int
	pageID int
}

type friend struct {
	user1 int
	user2 int
}

func pageRecommendations(userID int, likes []like, friendships []friend) []int {
	likedByUser := make(map[int]bool)
	for _, l := range likes {
		if l.userID == userID {
			likedByUser[l.pageID] = true
		}
	}

	friends := make(map[int]bool)
	for _, f := range friendships {
		if f.user1 == userID {
			friends[f.user2] = true
		}
		if f.user2 == userID {
			friends[f.user1] = true
		}
	}

	recommend := make(map[int]bool)
	for _, l := range likes {
		if friends[l.userID] && !likedByUser[l.pageID] {
			recommend[l.pageID] = true
		}
	}

	result := make([]int, 0, len(recommend))
	for p := range recommend {
		result = append(result, p)
	}
	sort.Ints(result)
	return result
}

func main() {
	likes := []like{
		{1, 101}, {1, 102}, {2, 101}, {2, 103}, {3, 102},
	}
	friendships := []friend{
		{1, 2}, {1, 3},
	}
	fmt.Printf("%v (expected: [103] or [102 103])\n",
		pageRecommendations(1, likes, friendships))

	fmt.Printf("%v (expected: [])\n",
		pageRecommendations(2, likes, friendships))
}
```

## 1265 — Print Immutable Linked List In Reverse

```go
package main

import (
	"fmt"
)

// LeetCode #1265: Print Immutable Linked List in Reverse
// https://leetcode.com/problems/print-immutable-linked-list-in-reverse/
// Difficulty: Medium [Paid]

// Print linked list in reverse using recursion (or stack).
// Immutable means we can't modify the list.

// Time: O(n)
// Space: O(n) for recursive stack

type ImmutableListNode struct {
	val  int
	next *ImmutableListNode
}

func (node *ImmutableListNode) getValue() int {
	return node.val
}

func (node *ImmutableListNode) getNext() *ImmutableListNode {
	return node.next
}

func printLinkedListInReverse(head *ImmutableListNode) {
	if head == nil {
		return
	}
	printLinkedListInReverse(head.getNext())
	fmt.Printf("%d ", head.getValue())
}

func main() {
	head := &ImmutableListNode{1, &ImmutableListNode{2, &ImmutableListNode{3, nil}}}
	fmt.Printf("Reversed: ")
	printLinkedListInReverse(head)
	fmt.Println()
}
```

## 1267 — Count Servers That Communicate

```go
package main

import (
	"fmt"
)

// LeetCode #1267: Count Servers that Communicate
// https://leetcode.com/problems/count-servers-that-communicate/
// Difficulty: Medium

// Count servers that can communicate with at least one other server
// in the same row or column.

// Time: O(m*n)
// Space: O(m+n)

func countServers(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	rowCount := make([]int, m)
	colCount := make([]int, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				rowCount[i]++
				colCount[j]++
			}
		}
	}

	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 && (rowCount[i] > 1 || colCount[j] > 1) {
				count++
			}
		}
	}

	return count
}

func main() {
	fmt.Printf("%d (expected: 3)\n",
		countServers([][]int{{1, 0}, {0, 1}}))

	fmt.Printf("%d (expected: 4)\n",
		countServers([][]int{{1, 0}, {1, 1}}))

	fmt.Printf("%d (expected: 0)\n",
		countServers([][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 1}}))
}
```

## 1268 — Search Suggestions System

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #1268: Search Suggestions System
// https://leetcode.com/problems/search-suggestions-system/
// Difficulty: Medium

// For each prefix of searchWord, return top 3 lexicographically
// smallest products that match the prefix.

// Time: O(n log n + m * n) where m = len(searchWord)
// Space: O(n)

func suggestedProducts(products []string, searchWord string) [][]string {
	sort.Strings(products)

	result := make([][]string, len(searchWord))

	for i := 1; i <= len(searchWord); i++ {
		prefix := searchWord[:i]
		suggestions := make([]string, 0)

		for _, p := range products {
			if len(p) >= i && p[:i] == prefix {
				suggestions = append(suggestions, p)
				if len(suggestions) == 3 {
					break
				}
			}
		}

		result[i-1] = suggestions
	}

	return result
}

func main() {
	products := []string{"mobile", "mouse", "moneypot", "monitor", "mousepad"}
	searchWord := "mouse"
	result := suggestedProducts(products, searchWord)
	for i, r := range result {
		fmt.Printf("prefix %q: %v\n", searchWord[:i+1], r)
	}
}
```

## 1270 — All People Report To The Given Manager

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #1270: All People Report to the Given Manager
// https://leetcode.com/problems/all-people-report-to-the-given-manager/
// Difficulty: Medium [Paid]

// Find all employees who directly or indirectly report to the head
// (manager_id = 1 or manager_id is in the reporting chain).

// Time: O(n)
// Space: O(n)

func allPeopleReportTo(employees [][]int) []int {
	// employees[i] = [employee_id, manager_id]
	// Find all employees who report to employee_id=1 (directly or indirectly)

	adj := make(map[int][]int)
	for _, e := range employees {
		empID, mgrID := e[0], e[1]
		if mgrID != 0 { // 0 means no manager (head)
			adj[mgrID] = append(adj[mgrID], empID)
		}
	}

	result := make([]int, 0)
	queue := []int{1}

	for len(queue) > 0 {
		mgr := queue[0]
		queue = queue[1:]
		for _, emp := range adj[mgr] {
			result = append(result, emp)
			queue = append(queue, emp)
		}
	}

	sort.Ints(result)
	return result
}

func main() {
	employees := [][]int{
		{1, 0},
		{2, 1},
		{3, 2},
		{4, 1},
		{5, 3},
	}
	fmt.Printf("%v (expected: [2 3 4 5])\n", allPeopleReportTo(employees))

	employees2 := [][]int{
		{1, 0},
		{2, 1},
		{3, 1},
	}
	fmt.Printf("%v (expected: [2 3])\n", allPeopleReportTo(employees2))
}
```

## 1272 — Remove Interval

```go
package main

import (
	"fmt"
)

// LeetCode #1272: Remove Interval
// https://leetcode.com/problems/remove-interval/
// Difficulty: Medium [Paid]

// Remove an interval from a set of non-overlapping intervals.
// Return the resulting intervals.

// Time: O(n)
// Space: O(n)

func removeInterval(intervals [][]int, toBeRemoved []int) [][]int {
	result := make([][]int, 0)
	rs, re := toBeRemoved[0], toBeRemoved[1]

	for _, interval := range intervals {
		s, e := interval[0], interval[1]
		// No overlap
		if e <= rs || s >= re {
			result = append(result, interval)
		} else {
			// Left part (if any)
			if s < rs {
				result = append(result, []int{s, rs})
			}
			// Right part (if any)
			if e > re {
				result = append(result, []int{re, e})
			}
		}
	}

	return result
}

func main() {
	fmt.Printf("%v (expected: [[0 1] [6 7]])\n",
		removeInterval([][]int{{0, 2}, {3, 4}, {5, 7}}, []int{1, 6}))

	fmt.Printf("%v (expected: [[3 4]])\n",
		removeInterval([][]int{{0, 5}}, []int{0, 3}))
}
```

## 1273 — Delete Tree Nodes

```go
package main

import (
	"fmt"
)

// LeetCode #1273: Delete Tree Nodes
// https://leetcode.com/problems/delete-tree-nodes/
// Difficulty: Medium [Paid]

// Delete subtrees where sum of node values = 0.
// Return number of remaining nodes.

// Time: O(n)
// Space: O(n)

func deleteTreeNodes(nodes int, parent []int, value []int) int {
	children := make([][]int, nodes)
	for i := 1; i < nodes; i++ {
		p := parent[i]
		children[p] = append(children[p], i)
	}

	var dfs func(node int) (int, int) // returns (sum, count)
	dfs = func(node int) (int, int) {
		sum := value[node]
		count := 1
		for _, child := range children[node] {
			childSum, childCount := dfs(child)
			sum += childSum
			count += childCount
		}
		if sum == 0 {
			return 0, 0
		}
		return sum, count
	}

	_, cnt := dfs(0)
	return cnt
}

func main() {
	fmt.Printf("%d (expected: 2)\n",
		deleteTreeNodes(7, []int{-1, 0, 0, 1, 2, 2, 2}, []int{1, -2, 4, 0, -2, -1, 1}))

	fmt.Printf("%d (expected: 4)\n",
		deleteTreeNodes(4, []int{-1, 0, 0, 1}, []int{1, 2, -1, -2}))
}
```

## 1276 — Number Of Burgers With No Waste Of Ingredients

```go
package main

import (
	"fmt"
)

// LeetCode #1276: Number of Burgers with No Waste of Ingredients
// https://leetcode.com/problems/number-of-burgers-with-no-waste-of-ingredients/
// Difficulty: Medium

// Jumbo burger = 4 slices of cheese, 1 tomato. Small = 2 cheese, 1 tomato.
// Given total tomato and cheese slices, find valid (jumbo, small) combination.
// Solve: 4j+2s=cheese, j+s=tomato => j = (cheese - 2*tomato)/2

// Time: O(1)
// Space: O(1)

func numOfBurgers(tomatoSlices int, cheeseSlices int) []int {
	// j + s = tomatoSlices
	// 4j + 2s = cheeseSlices
	// => 2j + (2j+2s) = cheeseSlices => 2j + 2*tomatoSlices = cheeseSlices
	// => j = (cheeseSlices - 2*tomatoSlices) / 2

	if cheeseSlices < 0 || tomatoSlices%2 != 0 {
		return []int{}
	}

	jumbo := (cheeseSlices - 2*tomatoSlices) / 2
	small := tomatoSlices - jumbo

	if jumbo < 0 || small < 0 || 4*jumbo+2*small != cheeseSlices {
		return []int{}
	}

	return []int{jumbo, small}
}

func main() {
	fmt.Printf("%v (expected: [1 6])\n", numOfBurgers(16, 7))
	fmt.Printf("%v (expected: [])\n", numOfBurgers(17, 4))
	fmt.Printf("%v (expected: [])\n", numOfBurgers(4, 17))
}
```

## 1277 — Count Square Submatrices With All Ones

```go
package main

import (
	"fmt"
)

// LeetCode #1277: Count Square Submatrices with All Ones
// https://leetcode.com/problems/count-square-submatrices-with-all-ones/
// Difficulty: Medium

// dp[i][j] = side length of largest square ending at (i,j).
// If grid[i][j]==1: dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])
// Sum all dp[i][j] = total squares.

// Time: O(m*n)
// Space: O(m*n)

func countSquares(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	total := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 1 {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = 1 + min(dp[i-1][j], min(dp[i][j-1], dp[i-1][j-1]))
				}
				total += dp[i][j]
			}
		}
	}

	return total
}

func main() {
	fmt.Printf("%d (expected: 15)\n",
		countSquares([][]int{
			{0, 1, 1, 1},
			{1, 1, 1, 1},
			{0, 1, 1, 1},
		}))

	fmt.Printf("%d (expected: 7)\n",
		countSquares([][]int{
			{1, 0, 1},
			{1, 1, 0},
			{1, 1, 0},
		}))
}
```

## 1282 — Group The People Given The Group Size They Belong To

```go
package main

import (
	"fmt"
)

// LeetCode #1282: Group the People Given the Group Size They Belong To
// https://leetcode.com/problems/group-the-people-given-the-group-size-they-belong-to/
// Difficulty: Medium

// Group people by desired group size, then partition each group.

// Time: O(n)
// Space: O(n)

func groupThePeople(groupSizes []int) [][]int {
	groups := make(map[int][]int)
	result := make([][]int, 0)

	for person, size := range groupSizes {
		groups[size] = append(groups[size], person)
		if len(groups[size]) == size {
			result = append(result, groups[size])
			delete(groups, size)
		}
	}

	return result
}

func main() {
	fmt.Printf("%v\n", groupThePeople([]int{3, 3, 3, 3, 3, 1, 3}))
	fmt.Printf("%v\n", groupThePeople([]int{2, 1, 3, 3, 3, 2}))
}
```

## 1283 — Find The Smallest Divisor Given A Threshold

```go
package main

import (
	"fmt"
	"math"
)

// LeetCode #1283: Find the Smallest Divisor Given a Threshold
// https://leetcode.com/problems/find-the-smallest-divisor-given-a-threshold/
// Difficulty: Medium

// Binary search the smallest divisor such that sum(ceil(nums[i]/div)) <= threshold.

// Time: O(n log max(nums))
// Space: O(1)

func smallestDivisor(nums []int, threshold int) int {
	maxVal := 0
	for _, v := range nums {
		if v > maxVal {
			maxVal = v
		}
	}

	lo, hi := 1, maxVal
	for lo < hi {
		mid := lo + (hi-lo)/2
		sum := 0
		for _, v := range nums {
			sum += int(math.Ceil(float64(v) / float64(mid)))
		}
		if sum <= threshold {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func main() {
	fmt.Printf("%d (expected: 3)\n", smallestDivisor([]int{1, 2, 5, 9}, 6))
	fmt.Printf("%d (expected: 44)\n", smallestDivisor([]int{44, 22, 33, 11, 1}, 5))
}
```

## 1285 — Find The Start And End Number Of Continuous Ranges

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #1285: Find the Start and End Number of Continuous Ranges
// https://leetcode.com/problems/find-the-start-and-end-number-of-continuous-ranges/
// Difficulty: Medium [Paid]

// Given a sorted list of unique integers, find continuous ranges.

// Time: O(n)
// Space: O(n)

func findContinuousRanges(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	result := make([][]int, 0)
	start := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1]+1 {
			result = append(result, []int{start, nums[i-1]})
			start = nums[i]
		}
	}
	result = append(result, []int{start, nums[len(nums)-1]})

	sort.Slice(result, func(i, j int) bool {
		return result[i][0] < result[j][0]
	})

	return result
}

func main() {
	fmt.Printf("%v (expected: [[1 3] [6 7] [9 9]])\n",
		findContinuousRanges([]int{1, 2, 3, 6, 7, 9}))

	fmt.Printf("%v (expected: [[1 1]])\n",
		findContinuousRanges([]int{1}))

	fmt.Printf("%v (expected: [])\n",
		findContinuousRanges([]int{}))
}
```

## 1286 — Iterator For Combination

```go
package main

import (
	"fmt"
)

// LeetCode #1286: Iterator for Combination
// https://leetcode.com/problems/iterator-for-combination/
// Difficulty: Medium

// Generate all combinations of length combinationLength from characters
// in lexicographic order using next() and hasNext().

// Time: O(C(n,k)) init, O(1) next/hasNext
// Space: O(C(n,k))

type CombinationIterator struct {
	combinations []string
	idx          int
}

func Constructor(characters string, combinationLength int) CombinationIterator {
	combs := make([]string, 0)
	n := len(characters)

	var backtrack func(start int, curr []byte)
	backtrack = func(start int, curr []byte) {
		if len(curr) == combinationLength {
			combs = append(combs, string(curr))
			return
		}
		for i := start; i < n; i++ {
			curr = append(curr, characters[i])
			backtrack(i+1, curr)
			curr = curr[:len(curr)-1]
		}
	}

	backtrack(0, []byte{})
	return CombinationIterator{combinations: combs, idx: 0}
}

func (this *CombinationIterator) Next() string {
	res := this.combinations[this.idx]
	this.idx++
	return res
}

func (this *CombinationIterator) HasNext() bool {
	return this.idx < len(this.combinations)
}

func main() {
	iter := Constructor("abc", 2)
	fmt.Printf("next: %q (expected: %q)\n", iter.Next(), "ab")
	fmt.Printf("hasNext: %t (expected: true)\n", iter.HasNext())
	fmt.Printf("next: %q (expected: %q)\n", iter.Next(), "ac")
	fmt.Printf("hasNext: %t (expected: true)\n", iter.HasNext())
	fmt.Printf("next: %q (expected: %q)\n", iter.Next(), "bc")
	fmt.Printf("hasNext: %t (expected: false)\n", iter.HasNext())
}
```

## 1288 — Remove Covered Intervals

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #1288: Remove Covered Intervals
// https://leetcode.com/problems/remove-covered-intervals/
// Difficulty: Medium

// Sort by start ascending, end descending. Track current max end.
// If end <= maxEnd, interval is covered.

// Time: O(n log n)
// Space: O(1)

func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] > intervals[j][1]
	})

	count := 0
	maxEnd := 0

	for _, inv := range intervals {
		if inv[1] > maxEnd {
			count++
			maxEnd = inv[1]
		}
	}

	return count
}

func main() {
	fmt.Printf("%d (expected: 2)\n",
		removeCoveredIntervals([][]int{{1, 4}, {3, 6}, {2, 8}}))

	fmt.Printf("%d (expected: 1)\n",
		removeCoveredIntervals([][]int{{1, 4}, {2, 3}}))

	fmt.Printf("%d (expected: 2)\n",
		removeCoveredIntervals([][]int{{0, 10}, {5, 12}}))
}
```

## 1291 — Sequential Digits

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #1291: Sequential Digits
// https://leetcode.com/problems/sequential-digits/
// Difficulty: Medium

// Generate all sequential digits numbers in range [low, high].
// Sequential = each digit is 1 more than previous.

// Time: O(1) — bounded by number of possible sequential digits
// Space: O(1)

func sequentialDigits(low int, high int) []int {
	result := make([]int, 0)
	digits := "123456789"

	for length := len(fmt.Sprintf("%d", low)); length <= len(fmt.Sprintf("%d", high)); length++ {
		for start := 0; start+length <= 9; start++ {
			num := 0
			for i := 0; i < length; i++ {
				num = num*10 + int(digits[start+i]-'0')
			}
			if num >= low && num <= high {
				result = append(result, num)
			}
		}
	}

	sort.Ints(result)
	return result
}

func main() {
	fmt.Printf("%v (expected: [123 234])\n", sequentialDigits(100, 300))
	fmt.Printf("%v (expected: [1234 2345 3456 4567 5678 6789 12345])\n", sequentialDigits(1000, 13000))
}
```

## 1292 — Maximum Side Length Of A Square With Sum Less Than Or Equal To Threshold

```go
package main

import (
	"fmt"
)

// LeetCode #1292: Maximum Side Length of a Square with Sum Less than or Equal to Threshold
// https://leetcode.com/problems/maximum-side-length-of-a-square-with-sum-less-than-or-equal-to-threshold/
// Difficulty: Medium

// Prefix sum matrix + binary search on side length.
// For each square, sum = prefix[r+s][c+s] - prefix[r][c+s] - prefix[r+s][c] + prefix[r][c].

// Time: O(m * n * log(min(m,n)))
// Space: O(m*n)

func maxSideLength(mat [][]int, threshold int) int {
	m, n := len(mat), len(mat[0])
	prefix := make([][]int, m+1)
	for i := range prefix {
		prefix[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			prefix[i][j] = mat[i-1][j-1] + prefix[i-1][j] + prefix[i][j-1] - prefix[i-1][j-1]
		}
	}

	maxSide := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			for s := maxSide + 1; s <= m-i+1 && s <= n-j+1; s++ {
				sum := prefix[i+s-1][j+s-1] - prefix[i-1][j+s-1] - prefix[i+s-1][j-1] + prefix[i-1][j-1]
				if sum <= threshold {
					if s > maxSide {
						maxSide = s
					}
				} else {
					break
				}
			}
		}
	}

	return maxSide
}

func main() {
	fmt.Printf("%d (expected: 2)\n",
		maxSideLength([][]int{{1, 1, 3, 2, 4, 3, 2}, {1, 1, 3, 2, 4, 3, 2}, {1, 1, 3, 2, 4, 3, 2}}, 4))

	fmt.Printf("%d (expected: 3)\n",
		maxSideLength([][]int{{2, 2, 2, 2, 2}, {2, 2, 2, 2, 2}, {2, 2, 2, 2, 2}}, 12))
}
```

## 1296 — Divide Array In Sets Of K Consecutive Numbers

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #1296: Divide Array in Sets of K Consecutive Numbers
// https://leetcode.com/problems/divide-array-in-sets-of-k-consecutive-numbers/
// Difficulty: Medium

// Sort array, greedily form groups of size k.
// Use frequency map to track available numbers.

// Time: O(n log n)
// Space: O(n)

func isPossibleDivide(nums []int, k int) bool {
	if len(nums)%k != 0 {
		return false
	}

	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}

	sort.Ints(nums)
	for _, v := range nums {
		if freq[v] == 0 {
			continue
		}
		for i := 0; i < k; i++ {
			if freq[v+i] == 0 {
				return false
			}
			freq[v+i]--
		}
	}

	return true
}

func main() {
	fmt.Printf("%t (expected: true)\n", isPossibleDivide([]int{1, 2, 3, 3, 4, 4, 5, 6}, 4))
	fmt.Printf("%t (expected: false)\n", isPossibleDivide([]int{3, 2, 1, 2, 3, 4, 3, 4, 5, 9, 10, 11}, 3))
	fmt.Printf("%t (expected: true)\n", isPossibleDivide([]int{1, 2, 3, 4}, 2))
}
```

## 1297 — Maximum Number Of Occurrences Of A Substring

```go
package main

import (
	"fmt"
)

// LeetCode #1297: Maximum Number of Occurrences of a Substring
// https://leetcode.com/problems/maximum-number-of-occurrences-of-a-substring/
// Difficulty: Medium

// Find max occurrences of any substring meeting constraints:
// - size between 1 and maxLetters distinct characters
// - substring size = minSize (longer substrings are less frequent)

// Time: O(n * minSize)
// Space: O(n)

func maxFreq(s string, maxLetters int, minSize int, maxSize int) int {
	count := make(map[string]int)
	maxOccur := 0

	for i := 0; i+minSize <= len(s); i++ {
		sub := s[i : i+minSize]
		// Count distinct chars
		chars := make(map[byte]bool)
		for j := 0; j < len(sub); j++ {
			chars[sub[j]] = true
		}
		if len(chars) <= maxLetters {
			count[sub]++
			if count[sub] > maxOccur {
				maxOccur = count[sub]
			}
		}
	}

	return maxOccur
}

func main() {
	fmt.Printf("%d (expected: 2)\n", maxFreq("aababcaab", 2, 3, 4))
	fmt.Printf("%d (expected: 2)\n", maxFreq("aaaa", 1, 3, 3))
	fmt.Printf("%d (expected: 3)\n", maxFreq("aabcabcab", 2, 3, 3))
}
```

## 1300 — Sum Of Mutated Array Closest To Target

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #1300: Sum of Mutated Array Closest to Target
// https://leetcode.com/problems/sum-of-mutated-array-closest-to-target/
// Difficulty: Medium

// Find integer value such that sum(arr[i] if arr[i] < value else value)
// is as close to target as possible. If tie, return smaller value.

// Time: O(n log n + n log max(arr))
// Space: O(1)

func findBestValue(arr []int, target int) int {
	sort.Ints(arr)
	n := len(arr)

	prefix := make([]int, n+1)
	for i := 1; i <= n; i++ {
		prefix[i] = prefix[i-1] + arr[i-1]
	}

	lo, hi := 0, arr[n-1]
	bestVal, minDiff := 0, target

	for lo <= hi {
		mid := lo + (hi-lo)/2

		// Find first index > mid
		idx := sort.Search(n, func(i int) bool {
			return arr[i] > mid
		})

		sum := prefix[idx] + mid*(n-idx)
		diff := abs(sum - target)

		if diff < minDiff || (diff == minDiff && mid < bestVal) {
			bestVal = mid
			minDiff = diff
		}

		if sum < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	return bestVal
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Printf("%d (expected: 3)\n", findBestValue([]int{4, 9, 3}, 10))
	fmt.Printf("%d (expected: 5)\n", findBestValue([]int{2, 3, 5}, 10))
	fmt.Printf("%d (expected: 11361)\n", findBestValue([]int{60864, 25176, 27249, 21296, 20204}, 56803))
}
```

## 1302 — Deepest Leaves Sum

```go
package main

import (
	"fmt"
)

// LeetCode #1302: Deepest Leaves Sum
// https://leetcode.com/problems/deepest-leaves-sum/
// Difficulty: Medium

// Sum values of the deepest leaves in a binary tree.

// Time: O(n)
// Space: O(h) where h = tree height

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSum := 0
		size := len(queue)

		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			levelSum += node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// If this is the last level, return its sum
		if len(queue) == 0 {
			return levelSum
		}
	}

	return 0
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4, Left: &TreeNode{Val: 7}},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val:   3,
			Right: &TreeNode{Val: 6, Right: &TreeNode{Val: 8}},
		},
	}
	fmt.Printf("%d (expected: 15)\n", deepestLeavesSum(root))

	root2 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	fmt.Printf("%d (expected: 5)\n", deepestLeavesSum(root2))
}
```

## 1305 — All Elements In Two Binary Search Trees

```go
package main

// LeetCode #1305: All Elements in Two Binary Search Trees
// https://leetcode.com/problems/all-elements-in-two-binary-search-trees/
// Difficulty: Medium

import "fmt"
import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Test case 1
	root1 := &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 4}}
	root2 := &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 3}}
	fmt.Println(getAllElements(root1, root2)) // [0,1,1,2,3,4]

	// Test case 2
	root3 := &TreeNode{Val: 1, Right: &TreeNode{Val: 8}}
	root4 := &TreeNode{Val: 8, Left: &TreeNode{Val: 1}}
	fmt.Println(getAllElements(root3, root4)) // [1,1,8,8]

	// Test case 3
	fmt.Println(getAllElements(nil, nil)) // []
}

// Time: O(m+n) for tree traversal + O(m+n) for merge = O(m+n)
// Space: O(m+n) for storing values
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	var vals1, vals2 []int
	inorder(root1, &vals1)
	inorder(root2, &vals2)
	return merge(vals1, vals2)
}

func inorder(node *TreeNode, vals *[]int) {
	if node == nil {
		return
	}
	inorder(node.Left, vals)
	*vals = append(*vals, node.Val)
	inorder(node.Right, vals)
}

func merge(a, b []int) []int {
	if len(a) == 0 {
		return b
	}
	if len(b) == 0 {
		return a
	}
	result := make([]int, 0, len(a)+len(b))
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}
	result = append(result, a[i:]...)
	result = append(result, b[j:]...)
	return result
}

// Below ensures the signature matches expected problem name
func getAllElementsSort(root1 *TreeNode, root2 *TreeNode) []int {
	var vals []int
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		vals = append(vals, node.Val)
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root1)
	dfs(root2)
	sort.Ints(vals)
	return vals
}
```

## 1306 — Jump Game Iii

```go
package main

// LeetCode #1306: Jump Game III
// https://leetcode.com/problems/jump-game-iii/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(canReach([]int{4, 2, 3, 0, 3, 1, 2}, 5)) // true

	// Test case 2
	fmt.Println(canReach([]int{4, 2, 3, 0, 3, 1, 2}, 0)) // true

	// Test case 3
	fmt.Println(canReach([]int{3, 0, 2, 1, 2}, 2)) // false
}

// Time: O(n) where n is length of arr
// Space: O(n) for visited array and queue
func canReach(arr []int, start int) bool {
	n := len(arr)
	visited := make([]bool, n)
	queue := []int{start}
	visited[start] = true

	for len(queue) > 0 {
		i := queue[0]
		queue = queue[1:]

		if arr[i] == 0 {
			return true
		}

		for _, next := range []int{i + arr[i], i - arr[i]} {
			if next >= 0 && next < n && !visited[next] {
				visited[next] = true
				queue = append(queue, next)
			}
		}
	}

	return false
}
```

## 1308 — Running Total For Different Genders

```go
package main

// LeetCode #1308: Running Total for Different Genders
// https://leetcode.com/problems/running-total-for-different-genders/
// Difficulty: Medium

import "fmt"

func main() {
	scores := []struct {
		playerName  string
		gender      string
		day         string
		scorePoints int
	}{
		{"Aron", "F", "2020-01-01", 17},
		{"Alice", "F", "2020-01-07", 23},
		{"Bajrang", "M", "2020-01-07", 7},
		{"Khali", "M", "2019-12-25", 11},
		{"Slaman", "M", "2019-12-30", 13},
		{"Joe", "M", "2019-12-31", 3},
		{"Jose", "M", "2019-12-18", 2},
		{"Priya", "F", "2019-12-31", 15},
		{"Priyanka", "F", "2019-11-23", 17},
	}

	result := runningTotal(scores)
	for _, r := range result {
		fmt.Printf("%s %s %s %d\n", r.gender, r.day, r.playerName, r.total)
	}
}

type result struct {
	gender     string
	day        string
	playerName string
	total      int
}

// Time: O(n log n) due to sorting
// Space: O(n) for storing results
func runningTotal(scores []struct {
	playerName string
	gender     string
	day        string
	scorePoints int
}) []result {
	// Group by gender, sort each group by day, then gender, then player_name
	// For simplicity, we process in Go using map
	type playerScore struct {
		player, day string
		score       int
	}

	groups := make(map[string][]playerScore)
	for _, s := range scores {
		groups[s.gender] = append(groups[s.gender], playerScore{s.playerName, s.day, s.scorePoints})
	}

	// Sort each group by day, then player_name
	for gender := range groups {
		// Bubble sort for simplicity - in real SQL this is ORDER BY
		group := groups[gender]
		for i := 0; i < len(group); i++ {
			for j := i + 1; j < len(group); j++ {
				if group[j].day < group[i].day ||
					(group[j].day == group[i].day && group[j].player < group[i].player) {
					group[i], group[j] = group[j], group[i]
				}
			}
		}
		groups[gender] = group
	}

	var results []result
	for _, gender := range []string{"F", "M"} {
		running := 0
		for _, ps := range groups[gender] {
			running += ps.score
			results = append(results, result{gender, ps.day, ps.player, running})
		}
	}
	return results
}
```

## 1310 — Xor Queries Of A Subarray

```go
package main

// LeetCode #1310: XOR Queries of a Subarray
// https://leetcode.com/problems/xor-queries-of-a-subarray/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(xorQueries([]int{1, 3, 4, 8}, [][]int{{0, 1}, {1, 2}, {0, 3}, {3, 3}}))
	// [2,7,14,8]

	// Test case 2
	fmt.Println(xorQueries([]int{4, 8, 2, 10}, [][]int{{2, 3}, {1, 3}, {0, 0}, {0, 3}}))
	// [8,0,4,4]

	// Test case 3
	fmt.Println(xorQueries([]int{2}, [][]int{{0, 0}}))
	// [2]
}

// Time: O(n + m) where n = len(arr), m = len(queries)
// Space: O(n) for prefix XOR array
func xorQueries(arr []int, queries [][]int) []int {
	n := len(arr)
	prefix := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] ^ arr[i]
	}

	result := make([]int, len(queries))
	for i, q := range queries {
		result[i] = prefix[q[1]+1] ^ prefix[q[0]]
	}
	return result
}
```

## 1311 — Get Watched Videos By Your Friends

```go
package main

// LeetCode #1311: Get Watched Videos by Your Friends
// https://leetcode.com/problems/get-watched-videos-by-your-friends/
// Difficulty: Medium

import "fmt"
import "sort"

func main() {
	// Test case 1
	fmt.Println(watchedVideosByFriends(
		[][]string{{"A", "B"}, {"C"}, {"B", "C"}, {"D"}},
		[][]int{{1, 2}, {0, 3}, {0, 3}, {1, 2}},
		0, 1,
	))
	// ["B","C"]

	// Test case 2
	fmt.Println(watchedVideosByFriends(
		[][]string{{"A", "B"}, {"C"}, {"B", "C"}, {"D"}},
		[][]int{{1, 2}, {0, 3}, {0, 3}, {1, 2}},
		0, 2,
	))
	// ["D"]

	// Test case 3
	fmt.Println(watchedVideosByFriends(
		[][]string{{"a"}, {"a"}, {"a"}},
		[][]int{{1}, {0}, {0}},
		0, 1,
	))
	// ["a"]
}

// Time: O(V + E + F log F) where V = friends count, F = videos count
// Space: O(V + F)
func watchedVideosByFriends(watchedVideos [][]string, friends [][]int, id int, level int) []string {
	n := len(friends)
	visited := make([]bool, n)
	queue := []int{id}
	visited[id] = true
	depth := 0

	for len(queue) > 0 && depth < level {
		size := len(queue)
		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]
			for _, f := range friends[curr] {
				if !visited[f] {
					visited[f] = true
					queue = append(queue, f)
				}
			}
		}
		depth++
	}

	if depth < level {
		return []string{}
	}

	freq := make(map[string]int)
	for _, person := range queue {
		for _, video := range watchedVideos[person] {
			freq[video]++
		}
	}

	videos := make([]string, 0, len(freq))
	for v := range freq {
		videos = append(videos, v)
	}

	sort.Slice(videos, func(i, j int) bool {
		if freq[videos[i]] != freq[videos[j]] {
			return freq[videos[i]] < freq[videos[j]]
		}
		return videos[i] < videos[j]
	})

	return videos
}
```

## 1314 — Matrix Block Sum

```go
package main

// LeetCode #1314: Matrix Block Sum
// https://leetcode.com/problems/matrix-block-sum/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(matrixBlockSum([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 1))
	// [[12,21,16],[27,45,33],[24,39,28]]

	// Test case 2
	fmt.Println(matrixBlockSum([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 2))
	// [[45,45,45],[45,45,45],[45,45,45]]

	// Test case 3
	fmt.Println(matrixBlockSum([][]int{{1}}, 1))
	// [[1]]
}

// Time: O(m*n) where m,n are matrix dimensions
// Space: O(m*n) for prefix sum matrix
func matrixBlockSum(mat [][]int, k int) [][]int {
	m, n := len(mat), len(mat[0])

	// Build 2D prefix sum (1-indexed)
	prefix := make([][]int, m+1)
	for i := range prefix {
		prefix[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			prefix[i+1][j+1] = mat[i][j] + prefix[i][j+1] + prefix[i+1][j] - prefix[i][j]
		}
	}

	// Calculate block sums
	result := make([][]int, m)
	for i := range result {
		result[i] = make([]int, n)
		for j := range result[i] {
			r1, c1 := max(0, i-k), max(0, j-k)
			r2, c2 := min(m-1, i+k), min(n-1, j+k)
			result[i][j] = prefix[r2+1][c2+1] - prefix[r1][c2+1] - prefix[r2+1][c1] + prefix[r1][c1]
		}
	}
	return result
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
```

## 1315 — Sum Of Nodes With Even Valued Grandparent

```go
package main

// LeetCode #1315: Sum of Nodes with Even-Valued Grandparent
// https://leetcode.com/problems/sum-of-nodes-with-even-valued-grandparent/
// Difficulty: Medium

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Test case 1
	root := &TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val: 7,
			Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 9}},
			Right: &TreeNode{Val: 7, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 4}},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 5}},
		},
	}
	fmt.Println(sumEvenGrandparent(root)) // 18

	// Test case 2: single node
	fmt.Println(sumEvenGrandparent(&TreeNode{Val: 1})) // 0

	// Test case 3
	root2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}},
		Right: &TreeNode{Val: 4, Right: &TreeNode{Val: 5}},
	}
	fmt.Println(sumEvenGrandparent(root2)) // 0 (no even-valued grandparent)
}

// Time: O(n) where n is number of nodes
// Space: O(h) where h is tree height (recursion stack)
func sumEvenGrandparent(root *TreeNode) int {
	return dfs(root, 1, 1) // parent and grandparent start as odd (1)
}

func dfs(node *TreeNode, parent, grandparent int) int {
	if node == nil {
		return 0
	}

	sum := 0
	if grandparent%2 == 0 {
		sum += node.Val
	}

	sum += dfs(node.Left, node.Val, parent)
	sum += dfs(node.Right, node.Val, parent)

	return sum
}
```

## 1318 — Minimum Flips To Make A Or B Equal To C

```go
package main

// LeetCode #1318: Minimum Flips to Make a OR b Equal to c
// https://leetcode.com/problems/minimum-flips-to-make-a-or-b-equal-to-c/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(minFlips(2, 6, 5)) // 3

	// Test case 2
	fmt.Println(minFlips(4, 2, 7)) // 1

	// Test case 3
	fmt.Println(minFlips(1, 2, 3)) // 0
}

// Time: O(bit length) = O(1) since 32 bits
// Space: O(1)
func minFlips(a int, b int, c int) int {
	flips := 0
	for i := 0; i < 32; i++ {
		bitC := (c >> i) & 1
		bitA := (a >> i) & 1
		bitB := (b >> i) & 1

		if bitC == 1 {
			if bitA == 0 && bitB == 0 {
				flips++ // need to flip one of them to 1
			}
		} else {
			if bitA == 1 {
				flips++
			}
			if bitB == 1 {
				flips++
			}
		}
	}
	return flips
}
```

## 1319 — Number Of Operations To Make Network Connected

```go
package main

// LeetCode #1319: Number of Operations to Make Network Connected
// https://leetcode.com/problems/number-of-operations-to-make-network-connected/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(makeConnected(4, [][]int{{0, 1}, {0, 2}, {1, 2}})) // 1

	// Test case 2
	fmt.Println(makeConnected(6, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}})) // 2

	// Test case 3
	fmt.Println(makeConnected(6, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}})) // -1
}

// Time: O(n + connections) for union-find operations
// Space: O(n) for parent and rank arrays
func makeConnected(n int, connections [][]int) int {
	if len(connections) < n-1 {
		return -1 // not enough cables
	}

	parent := make([]int, n)
	rank := make([]int, n)
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
		ra, rb := find(a), find(b)
		if ra == rb {
			return
		}
		if rank[ra] < rank[rb] {
			parent[ra] = rb
		} else if rank[ra] > rank[rb] {
			parent[rb] = ra
		} else {
			parent[rb] = ra
			rank[ra]++
		}
	}

	for _, conn := range connections {
		union(conn[0], conn[1])
	}

	components := 0
	for i := 0; i < n; i++ {
		if find(i) == i {
			components++
		}
	}

	return components - 1
}
```

## 1321 — Restaurant Growth

```go
package main

// LeetCode #1321: Restaurant Growth
// https://leetcode.com/problems/restaurant-growth/
// Difficulty: Medium

import "fmt"
import "sort"

func main() {
	customers := []struct {
		visitedOn string
		amount    int
	}{
		{"2020-01-01", 10},
		{"2020-01-02", 10},
		{"2020-01-03", 10},
		{"2020-01-04", 10},
		{"2020-01-05", 10},
		{"2020-01-06", 10},
		{"2020-01-07", 10},
		{"2020-01-08", 20},
		{"2020-01-09", 20},
		{"2020-01-10", 20},
	}

	result := restaurantGrowth(customers)
	for _, r := range result {
		fmt.Printf("%s %.2f\n", r.date, r.avg)
	}
}

type avgResult struct {
	date string
	avg  float64
}

// Time: O(n log n) due to sorting
// Space: O(n)
func restaurantGrowth(customers []struct {
	visitedOn string
	amount    int
}) []avgResult {
	// Group by date and sum amounts
	type daySum struct {
		date string
		total int
		count int
	}

	dateMap := make(map[string]*daySum)
	for _, c := range customers {
		if _, ok := dateMap[c.visitedOn]; !ok {
			dateMap[c.visitedOn] = &daySum{date: c.visitedOn}
		}
		dateMap[c.visitedOn].total += c.amount
		dateMap[c.visitedOn].count++
	}

	dates := make([]string, 0, len(dateMap))
	for d := range dateMap {
		dates = append(dates, d)
	}
	sort.Strings(dates)

	// Sliding window of 7 days
	var result []avgResult
	window := make([]int, 0, 7)

	for _, d := range dates {
		ds := dateMap[d]
		window = append(window, ds.total)
		if len(window) > 7 {
			window = window[1:]
		}
		if len(window) == 7 {
			sum := 0
			for _, v := range window {
				sum += v
			}
			result = append(result, avgResult{d, float64(sum) / 7.0})
		}
	}

	return result
}
```

## 1324 — Print Words Vertically

```go
package main

// LeetCode #1324: Print Words Vertically
// https://leetcode.com/problems/print-words-vertically/
// Difficulty: Medium

import "fmt"
import "strings"

func main() {
	// Test case 1
	fmt.Println(printVertically("HOW ARE YOU"))
	// ["HAY","ORO","WEU"]

	// Test case 2
	fmt.Println(printVertically("TO BE OR NOT TO BE"))
	// ["TBONTB","OEROOE","   T"]

	// Test case 3
	fmt.Println(printVertically("CONTEST IS COMING"))
	// ["CIC","OSO","N M","T I","E N","S G","T"]
}

// Time: O(m*n) where m = max word length, n = number of words
// Space: O(m*n) for the result
func printVertically(s string) []string {
	words := strings.Fields(s)
	maxLen := 0
	for _, w := range words {
		if len(w) > maxLen {
			maxLen = len(w)
		}
	}

	result := make([]string, maxLen)
	for i := 0; i < maxLen; i++ {
		var sb strings.Builder
		for _, w := range words {
			if i < len(w) {
				sb.WriteByte(w[i])
			} else {
				sb.WriteByte(' ')
			}
		}
		// Trim trailing spaces
		result[i] = strings.TrimRight(sb.String(), " ")
	}

	return result
}
```

## 1325 — Delete Leaves With A Given Value

```go
package main

// LeetCode #1325: Delete Leaves With a Given Value
// https://leetcode.com/problems/delete-leaves-with-a-given-value/
// Difficulty: Medium

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Test case 1
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 2}},
		Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 4}},
	}
	result := removeLeafNodes(root, 2)
	fmt.Println(result) // [1,null,3,null,4]

	// Test case 2
	root2 := &TreeNode{Val: 1, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 1}}
	result2 := removeLeafNodes(root2, 1)
	fmt.Println(result2) // nil

	// Test case 3
	root3 := &TreeNode{Val: 1, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 2}}}
	result3 := removeLeafNodes(root3, 3)
	fmt.Println(result3) // [1,3,null,null,2]
}

// Time: O(n) where n is number of nodes
// Space: O(h) for recursion stack, h = tree height
func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left = removeLeafNodes(root.Left, target)
	root.Right = removeLeafNodes(root.Right, target)

	// Post-order: if current node is a leaf and its value equals target, delete it
	if root.Left == nil && root.Right == nil && root.Val == target {
		return nil
	}

	return root
}

// For printing in main
func (n *TreeNode) String() string {
	if n == nil {
		return "nil"
	}
	return fmt.Sprintf("%d %s %s", n.Val, n.Left, n.Right)
}
```

## 1328 — Break A Palindrome

```go
package main

// LeetCode #1328: Break a Palindrome
// https://leetcode.com/problems/break-a-palindrome/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(breakPalindrome("abccba")) // "aaccba"

	// Test case 2
	fmt.Println(breakPalindrome("a")) // ""

	// Test case 3
	fmt.Println(breakPalindrome("aa")) // "ab"

	// Test case 4 - all 'a's
	fmt.Println(breakPalindrome("aaa")) // "aab"
}

// Time: O(n) where n = length of palindrome string
// Space: O(n) for the byte array
func breakPalindrome(palindrome string) string {
	n := len(palindrome)
	if n <= 1 {
		return ""
	}

	bytes := []byte(palindrome)
	// Try to change first non-'a' to 'a' (only in first half to maintain smallest lexicographically)
	for i := 0; i < n/2; i++ {
		if bytes[i] != 'a' {
			bytes[i] = 'a'
			return string(bytes)
		}
	}

	// All characters in first half are 'a', change last character to 'b'
	bytes[n-1] = 'b'
	return string(bytes)
}
```

## 1329 — Sort The Matrix Diagonally

```go
package main

// LeetCode #1329: Sort the Matrix Diagonally
// https://leetcode.com/problems/sort-the-matrix-diagonally/
// Difficulty: Medium

import "fmt"
import "sort"

func main() {
	// Test case 1
	fmt.Println(diagonalSort([][]int{{3, 3, 1, 1}, {2, 2, 1, 2}, {1, 1, 1, 2}}))
	// [[1,1,1,1],[1,2,2,2],[1,2,3,3]]

	// Test case 2
	fmt.Println(diagonalSort([][]int{{11, 25, 66, 1, 69, 7}, {23, 55, 17, 45, 15, 52}, {75, 31, 36, 44, 58, 8}, {22, 27, 33, 25, 68, 4}, {84, 28, 14, 11, 5, 50}}))

	// Test case 3
	fmt.Println(diagonalSort([][]int{{1}}))
	// [[1]]
}

// Time: O(m*n*log(min(m,n))) - sorting each diagonal
// Space: O(m*n) for storing diagonal elements
func diagonalSort(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])

	// Each diagonal starting from (i, 0) and (0, j)
	// Key insight: elements on same diagonal have same (i-j)

	// Group diagonals by (row - col) offset
	diagonals := make(map[int][]int)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			diagonals[i-j] = append(diagonals[i-j], mat[i][j])
		}
	}

	// Sort each diagonal
	for _, d := range diagonals {
		sort.Ints(d)
	}

	// Place sorted values back
	// We need to track where we are in each diagonal
	counters := make(map[int]int)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			offset := i - j
			mat[i][j] = diagonals[offset][counters[offset]]
			counters[offset]++
		}
	}

	return mat
}
```

## 1333 — Filter Restaurants By Vegan Friendly Price And Distance

```go
package main

// LeetCode #1333: Filter Restaurants by Vegan-Friendly, Price and Distance
// https://leetcode.com/problems/filter-restaurants-by-vegan-friendly-price-and-distance/
// Difficulty: Medium

import "fmt"
import "sort"

func main() {
	// Test case 1
	fmt.Println(filterRestaurants([][]int{
		{1, 4, 1, 40, 10},
		{2, 8, 0, 50, 5},
		{3, 8, 1, 30, 4},
		{4, 10, 0, 10, 3},
		{5, 1, 1, 15, 1},
	}, 1, 50, 10))
	// [3,1,5]

	// Test case 2
	fmt.Println(filterRestaurants([][]int{
		{1, 4, 1, 40, 10},
		{2, 8, 0, 50, 5},
		{3, 8, 1, 30, 4},
		{4, 10, 0, 10, 3},
		{5, 1, 1, 15, 1},
	}, 0, 50, 10))
	// [4,3,2,1,5]

	// Test case 3
	fmt.Println(filterRestaurants([][]int{
		{1, 4, 1, 40, 10},
	}, 1, 30, 10))
	// []
}

type restaurant struct {
	id, rating, veganFriendly, price, distance int
}

// Time: O(n log n) for sorting
// Space: O(n) for storing filtered results
func filterRestaurants(restaurants [][]int, veganFriendly int, maxPrice int, maxDistance int) []int {
	var filtered []restaurant
	for _, r := range restaurants {
		if veganFriendly == 1 && r[2] != 1 {
			continue
		}
		if r[3] > maxPrice {
			continue
		}
		if r[4] > maxDistance {
			continue
		}
		filtered = append(filtered, restaurant{r[0], r[1], r[2], r[3], r[4]})
	}

	sort.Slice(filtered, func(i, j int) bool {
		if filtered[i].rating != filtered[j].rating {
			return filtered[i].rating > filtered[j].rating
		}
		return filtered[i].id > filtered[j].id
	})

	result := make([]int, len(filtered))
	for i, r := range filtered {
		result[i] = r.id
	}
	return result
}
```

## 1334 — Find The City With The Smallest Number Of Neighbors At A Threshold Distance

```go
package main

// LeetCode #1334: Find the City With the Smallest Number of Neighbors at a Threshold Distance
// https://leetcode.com/problems/find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance/
// Difficulty: Medium

import "fmt"
import "math"

func main() {
	// Test case 1
	fmt.Println(findTheCity(4, [][]int{{0, 1, 3}, {1, 2, 1}, {1, 3, 4}, {2, 3, 1}}, 4)) // 3

	// Test case 2
	fmt.Println(findTheCity(5, [][]int{
		{0, 1, 2}, {0, 4, 8}, {1, 2, 3}, {1, 4, 2}, {2, 3, 1}, {3, 4, 1},
	}, 2)) // 0

	// Test case 3
	fmt.Println(findTheCity(6, [][]int{
		{0, 3, 5}, {2, 3, 7}, {0, 5, 2}, {0, 2, 4}, {1, 4, 5}, {3, 4, 6},
	}, 10)) // 4
}

// Time: O(n^3) where n = number of cities (Floyd-Warshall)
// Space: O(n^2) for distance matrix
func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	// Initialize distance matrix
	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			if i != j {
				dist[i][j] = math.MaxInt32
			}
		}
	}

	for _, e := range edges {
		dist[e[0]][e[1]] = e[2]
		dist[e[1]][e[0]] = e[2]
	}

	// Floyd-Warshall
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dist[i][k]+dist[k][j] < dist[i][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}

	minReachable := n
	result := -1
	for i := 0; i < n; i++ {
		count := 0
		for j := 0; j < n; j++ {
			if i != j && dist[i][j] <= distanceThreshold {
				count++
			}
		}
		if count <= minReachable {
			minReachable = count
			result = i
		}
	}

	return result
}
```

## 1338 — Reduce Array Size To The Half

```go
package main

// LeetCode #1338: Reduce Array Size to The Half
// https://leetcode.com/problems/reduce-array-size-to-the-half/
// Difficulty: Medium

import "fmt"
import "sort"

func main() {
	// Test case 1
	fmt.Println(minSetSize([]int{3, 3, 3, 3, 5, 5, 5, 2, 2, 7})) // 2

	// Test case 2
	fmt.Println(minSetSize([]int{7, 7, 7, 7, 7, 7})) // 1

	// Test case 3
	fmt.Println(minSetSize([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})) // 5
}

// Time: O(n log n) for sorting frequencies
// Space: O(n) for frequency map
func minSetSize(arr []int) int {
	n := len(arr)
	freq := make(map[int]int)
	for _, v := range arr {
		freq[v]++
	}

	counts := make([]int, 0, len(freq))
	for _, c := range freq {
		counts = append(counts, c)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(counts)))

	removed := 0
	half := n / 2
	for i, c := range counts {
		removed += c
		if removed >= half {
			return i + 1
		}
	}

	return len(counts)
}
```

## 1339 — Maximum Product Of Splitted Binary Tree

```go
package main

// LeetCode #1339: Maximum Product of Splitted Binary Tree
// https://leetcode.com/problems/maximum-product-of-splitted-binary-tree/
// Difficulty: Medium

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Test case 1
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}},
	}
	fmt.Println(maxProduct(root)) // 110

	// Test case 2
	root2 := &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}}}
	fmt.Println(maxProduct(root2)) // 90

	// Test case 3
	root3 := &TreeNode{Val: 1, Left: &TreeNode{Val: 1}}
	fmt.Println(maxProduct(root3)) // 1
}

const mod = 1_000_000_007

// Time: O(n) where n is number of nodes
// Space: O(h) for recursion stack
func maxProduct(root *TreeNode) int {
	var totalSum int
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		sum := node.Val + dfs(node.Left) + dfs(node.Right)
		return sum
	}
	totalSum = dfs(root)

	maxProd := 0
	var findMax func(*TreeNode) int
	findMax = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftSum := findMax(node.Left)
		rightSum := findMax(node.Right)
		subtreeSum := node.Val + leftSum + rightSum

		if subtreeSum != totalSum {
			prod := subtreeSum * (totalSum - subtreeSum)
			if prod > maxProd {
				maxProd = prod
			}
		}

		return subtreeSum
	}
	findMax(root)

	return maxProd % mod
}
```

## 1341 — Movie Rating

```go
package main

// LeetCode #1341: Movie Rating
// https://leetcode.com/problems/movie-rating/
// Difficulty: Medium

import "fmt"

func main() {
	movies := []struct {
		id    int
		title string
	}{
		{1, "Avengers"},
		{2, "Frozen 2"},
		{3, "Joker"},
	}
	users := []struct {
		id   int
		name string
	}{
		{1, "Daniel"},
		{2, "Monica"},
		{3, "Maria"},
	}
	ratings := []struct {
		userID    int
		movieID   int
		rating    int
		createdAt string
	}{
		{1, 1, 3, "2020-01-12"},
		{1, 2, 4, "2020-02-11"},
		{1, 3, 2, "2020-02-12"},
		{2, 1, 5, "2020-02-17"},
		{2, 2, 2, "2020-02-01"},
		{2, 3, 5, "2020-03-01"},
		{3, 1, 3, "2020-02-22"},
		{3, 2, 4, "2020-02-25"},
	}

	result := movieRating(ratings, users, movies)
	fmt.Println(result)
}

// Time: O(n) for counting
// Space: O(n) for maps
func movieRating(ratings []struct {
	userID    int
	movieID   int
	rating    int
	createdAt string
}, users []struct {
	id   int
	name string
}, movies []struct {
	id    int
	title string
}) string {
	// Count ratings per user
	userRatings := make(map[int]int)
	for _, r := range ratings {
		userRatings[r.userID]++
	}

	maxRatings := 0
	bestUser := ""
	for _, u := range users {
		cnt := userRatings[u.id]
		if cnt > maxRatings || (cnt == maxRatings && (bestUser == "" || u.name < bestUser)) {
			maxRatings = cnt
			bestUser = u.name
		}
	}

	// Average rating in Feb 2020
	movieScores := make(map[int]struct{ sum, count int })
	for _, r := range ratings {
		if r.createdAt >= "2020-02-01" && r.createdAt <= "2020-02-29" {
			s := movieScores[r.movieID]
			s.sum += r.rating
			s.count++
			movieScores[r.movieID] = s
		}
	}

	bestAvg := 0.0
	bestMovie := ""
	for _, m := range movies {
		if s, ok := movieScores[m.id]; ok && s.count > 0 {
			avg := float64(s.sum) / float64(s.count)
			if avg > bestAvg || (avg == bestAvg && (bestMovie == "" || m.title < bestMovie)) {
				bestAvg = avg
				bestMovie = m.title
			}
		}
	}

	return bestUser + " " + bestMovie
}
```

## 1343 — Number Of Sub Arrays Of Size K And Average Greater Than Or Equal To Threshold

```go
package main

// LeetCode #1343: Number of Sub-arrays of Size K and Average Greater than or Equal to Threshold
// https://leetcode.com/problems/number-of-sub-arrays-of-size-k-and-average-greater-than-or-equal-to-threshold/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(numOfSubarrays([]int{2, 2, 2, 2, 5, 5, 5, 8}, 3, 4)) // 3

	// Test case 2
	fmt.Println(numOfSubarrays([]int{11, 13, 17, 23, 29, 31, 7, 5, 2, 3}, 3, 5)) // 6

	// Test case 3
	fmt.Println(numOfSubarrays([]int{1, 1, 1, 1, 1}, 1, 0)) // 5
}

// Time: O(n) where n = len(arr)
// Space: O(1)
func numOfSubarrays(arr []int, k int, threshold int) int {
	targetSum := k * threshold
	windowSum := 0
	for i := 0; i < k; i++ {
		windowSum += arr[i]
	}

	count := 0
	if windowSum >= targetSum {
		count++
	}

	for i := k; i < len(arr); i++ {
		windowSum += arr[i] - arr[i-k]
		if windowSum >= targetSum {
			count++
		}
	}

	return count
}
```

## 1344 — Angle Between Hands Of A Clock

```go
package main

// LeetCode #1344: Angle Between Hands of a Clock
// https://leetcode.com/problems/angle-between-hands-of-a-clock/
// Difficulty: Medium

import "fmt"
import "math"

func main() {
	// Test case 1
	fmt.Println(angleClock(12, 30)) // 165

	// Test case 2
	fmt.Println(angleClock(3, 30)) // 75

	// Test case 3
	fmt.Println(angleClock(3, 15)) // 7.5

	// Test case 4
	fmt.Println(angleClock(4, 50)) // 155
}

// Time: O(1)
// Space: O(1)
func angleClock(hour int, minutes int) float64 {
	// Minute hand: 360 degrees in 60 minutes = 6 degrees per minute
	minAngle := float64(minutes) * 6.0

	// Hour hand: 360 degrees in 12 hours = 30 degrees per hour
	// Plus 0.5 degrees per minute (30 degrees / 60 minutes)
	hourAngle := float64(hour%12)*30.0 + float64(minutes)*0.5

	diff := math.Abs(hourAngle - minAngle)
	if diff > 180 {
		diff = 360 - diff
	}
	return diff
}
```

## 1347 — Minimum Number Of Steps To Make Two Strings Anagram

```go
package main

// LeetCode #1347: Minimum Number of Steps to Make Two Strings Anagram
// https://leetcode.com/problems/minimum-number-of-steps-to-make-two-strings-anagram/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(minSteps("bab", "aba")) // 1

	// Test case 2
	fmt.Println(minSteps("leetcode", "practice")) // 5

	// Test case 3
	fmt.Println(minSteps("anagram", "mangaar")) // 0
}

// Time: O(n) where n = length of strings
// Space: O(1) - fixed size array of 26
func minSteps(s string, t string) int {
	freq := make([]int, 26)
	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
		freq[t[i]-'a']--
	}

	steps := 0
	for _, f := range freq {
		if f > 0 {
			steps += f
		}
	}
	return steps
}
```

## 1348 — Tweet Counts Per Frequency

```go
package main

// LeetCode #1348: Tweet Counts Per Frequency
// https://leetcode.com/problems/tweet-counts-per-frequency/
// Difficulty: Medium

import "fmt"
import "sort"

type TweetCounts struct {
	tweets map[string][]int
}

func main() {
	tc := Constructor()
	tc.RecordTweet("tweet3", 0)
	tc.RecordTweet("tweet3", 60)
	tc.RecordTweet("tweet3", 10)
	fmt.Println(tc.GetTweetCountsPerFrequency("minute", "tweet3", 0, 59)) // [2]
	fmt.Println(tc.GetTweetCountsPerFrequency("minute", "tweet3", 0, 60)) // [2,1]

	tc2 := Constructor()
	tc2.RecordTweet("tweet1", 0)
	tc2.RecordTweet("tweet1", 10)
	tc2.RecordTweet("tweet2", 5)
	fmt.Println(tc2.GetTweetCountsPerFrequency("minute", "tweet2", 0, 100)) // [1]

	// Test case 3 - hour frequency
	tc3 := Constructor()
	tc3.RecordTweet("tweet", 0)
	tc3.RecordTweet("tweet", 3599)
	tc3.RecordTweet("tweet", 3600)
	fmt.Println(tc3.GetTweetCountsPerFrequency("hour", "tweet", 0, 7200)) // [2,1]
}

func Constructor() TweetCounts {
	return TweetCounts{tweets: make(map[string][]int)}
}

func (this *TweetCounts) RecordTweet(tweetName string, time int) {
	this.tweets[tweetName] = append(this.tweets[tweetName], time)
}

// Time: O(n log n + q) where n = tweets count, q = query interval count
// Space: O(n) for storing tweets
func (this *TweetCounts) GetTweetCountsPerFrequency(freq string, tweetName string, startTime int, endTime int) []int {
	f := 60 // default minute
	switch freq {
	case "hour":
		f = 3600
	case "day":
		f = 86400
	}

	times := this.tweets[tweetName]
	sort.Ints(times)

	size := (endTime-startTime)/f + 1
	result := make([]int, size)

	for _, t := range times {
		if t < startTime || t > endTime {
			continue
		}
		idx := (t - startTime) / f
		result[idx]++
	}

	return result
}
```

## 1352 — Product Of The Last K Numbers

```go
package main

// LeetCode #1352: Product of the Last K Numbers
// https://leetcode.com/problems/product-of-the-last-k-numbers/
// Difficulty: Medium

import "fmt"

type ProductOfNumbers struct {
	prefix []int // prefix product, reset after 0
}

func main() {
	pn := Constructor()
	pn.Add(3)
	pn.Add(0)
	pn.Add(2)
	pn.Add(5)
	pn.Add(4)
	fmt.Println(pn.GetProduct(2)) // 20
	fmt.Println(pn.GetProduct(3)) // 40
	fmt.Println(pn.GetProduct(4)) // 0

	pn2 := Constructor()
	pn2.Add(1)
	pn2.Add(2)
	pn2.Add(3)
	pn2.Add(4)
	fmt.Println(pn2.GetProduct(1)) // 4
	fmt.Println(pn2.GetProduct(4)) // 24
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{prefix: []int{1}}
}

// Time: O(1) for adding
// Space: O(n) for prefix array
func (this *ProductOfNumbers) Add(num int) {
	if num == 0 {
		this.prefix = []int{1}
		return
	}
	this.prefix = append(this.prefix, this.prefix[len(this.prefix)-1]*num)
}

// Time: O(1) for query
func (this *ProductOfNumbers) GetProduct(k int) int {
	if k >= len(this.prefix) {
		return 0
	}
	return this.prefix[len(this.prefix)-1] / this.prefix[len(this.prefix)-1-k]
}
```

## 1353 — Maximum Number Of Events That Can Be Attended

```go
package main

// LeetCode #1353: Maximum Number of Events That Can Be Attended
// https://leetcode.com/problems/maximum-number-of-events-that-can-be-attended/
// Difficulty: Medium

import "fmt"
import "sort"
import "container/heap"

type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func main() {
	// Test case 1
	fmt.Println(maxEvents([][]int{{1, 2}, {2, 3}, {3, 4}})) // 3

	// Test case 2
	fmt.Println(maxEvents([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 2}})) // 4

	// Test case 3
	fmt.Println(maxEvents([][]int{{1, 1}, {1, 2}, {1, 3}, {1, 4}, {2, 2}})) // 4

	// Test case 4 - single day
	fmt.Println(maxEvents([][]int{{1, 5}, {1, 5}, {1, 5}, {2, 3}, {2, 3}})) // 5
}

// Time: O(n log n) for sorting and heap operations
// Space: O(n) for heap
func maxEvents(events [][]int) int {
	if len(events) == 0 {
		return 0
	}

	// Sort by start day
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})

	h := &minHeap{}
	heap.Init(h)

	i := 0
	day := events[0][0]
	count := 0
	n := len(events)

	for i < n || h.Len() > 0 {
		// Add all events starting today
		for i < n && events[i][0] == day {
			heap.Push(h, events[i][1])
			i++
		}

		// Remove expired events (end day < today)
		for h.Len() > 0 && (*h)[0] < day {
			heap.Pop(h)
		}

		// Attend one event today (earliest ending)
		if h.Len() > 0 {
			heap.Pop(h)
			count++
			day++
		} else if i < n {
			day = events[i][0]
		}
	}

	return count
}
```

## 1355 — Activity Participants

```go
package main

// LeetCode #1355: Activity Participants
// https://leetcode.com/problems/activity-participants/
// Difficulty: Medium

import "fmt"

func main() {
	activities := []struct {
		id   int
		name string
	}{
		{1, "Eating"},
		{2, "Singing"},
		{3, "Horse Riding"},
	}
	friends := []struct {
		id         int
		name       string
		activityID int
	}{
		{1, "Jonathan D.", 1},
		{2, "Jade W.", 1},
		{3, "Victor J.", 1},
		{4, "Elvis O.", 2},
		{5, "Daniel A.", 2},
		{6, "Bob B.", 3},
	}

	result := activityParticipants(activities, friends)
	fmt.Println(result) // ["Singing"]
}

// Time: O(n) where n = number of friends
// Space: O(m) where m = number of activities
func activityParticipants(activities []struct {
	id   int
	name string
}, friends []struct {
	id         int
	name       string
	activityID int
}) []string {
	counts := make(map[int]int)
	for _, f := range friends {
		counts[f.activityID]++
	}

	if len(counts) == 0 {
		return nil
	}

	// Find min and max counts
	minCount, maxCount := len(friends), 0
	for _, c := range counts {
		if c < minCount {
			minCount = c
		}
		if c > maxCount {
			maxCount = c
		}
	}

	// Find activities with count between min and max (non-inclusive)
	var result []string
	for _, a := range activities {
		c := counts[a.id]
		if c > minCount && c < maxCount {
			result = append(result, a.name)
		}
	}
	return result
}
```

## 1357 — Apply Discount Every N Orders

```go
package main

// LeetCode #1357: Apply Discount Every n Orders
// https://leetcode.com/problems/apply-discount-every-n-orders/
// Difficulty: Medium

import "fmt"

type Cashier struct {
	n          int
	discount   int
	counter    int
	prices     map[int]int
}

func main() {
	cashier := NewCashier(3, 50, []int{1, 2, 3, 4, 5, 6, 7}, []int{100, 200, 300, 400, 300, 200, 100})
	fmt.Println(cashier.GetBill([]int{1, 2}, []int{1, 2}))    // 500.0
	fmt.Println(cashier.GetBill([]int{3, 7}, []int{10, 10}))  // 4000.0
	fmt.Println(cashier.GetBill([]int{1, 2, 3, 4, 5, 6, 7}, []int{1, 1, 1, 1, 1, 1, 1})) // 800.0
	fmt.Println(cashier.GetBill([]int{4}, []int{10}))         // 4000.0

	c2 := NewCashier(1, 10, []int{1, 2}, []int{100, 200})
	fmt.Println(c2.GetBill([]int{1}, []int{5})) // 450.0
}

func NewCashier(n int, discount int, products []int, prices []int) Cashier {
	priceMap := make(map[int]int)
	for i, p := range products {
		priceMap[p] = prices[i]
	}
	return Cashier{n, discount, 0, priceMap}
}

// Time: O(m) per bill where m = number of product types in the bill
// Space: O(p) where p = number of unique products
func (this *Cashier) GetBill(product []int, amount []int) float64 {
	this.counter++
	total := 0
	for i, p := range product {
		total += this.prices[p] * amount[i]
	}

	result := float64(total)
	if this.counter%this.n == 0 {
		result = result * float64(100-this.discount) / 100.0
	}
	return result
}
```

## 1358 — Number Of Substrings Containing All Three Characters

```go
package main

// LeetCode #1358: Number of Substrings Containing All Three Characters
// https://leetcode.com/problems/number-of-substrings-containing-all-three-characters/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(numberOfSubstrings("abcabc")) // 10

	// Test case 2
	fmt.Println(numberOfSubstrings("aaacb")) // 3

	// Test case 3
	fmt.Println(numberOfSubstrings("abc")) // 1
}

// Time: O(n) where n = length of string
// Space: O(1) - fixed array of 3
func numberOfSubstrings(s string) int {
	count := 0
	lastPos := [3]int{-1, -1, -1}

	for i := 0; i < len(s); i++ {
		lastPos[s[i]-'a'] = i
		// Find the minimum of the last positions of 'a', 'b', 'c'
		// This is the leftmost boundary of the substring containing all three
		minPos := lastPos[0]
		if lastPos[1] < minPos {
			minPos = lastPos[1]
		}
		if lastPos[2] < minPos {
			minPos = lastPos[2]
		}
		// If all three have been seen, count substrings ending at i
		if minPos >= 0 {
			count += minPos + 1
		}
	}

	return count
}
```

## 1361 — Validate Binary Tree Nodes

```go
package main

// LeetCode #1361: Validate Binary Tree Nodes
// https://leetcode.com/problems/validate-binary-tree-nodes/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(validateBinaryTreeNodes(4, []int{1, -1, 3, -1}, []int{2, -1, -1, -1})) // true

	// Test case 2
	fmt.Println(validateBinaryTreeNodes(4, []int{1, -1, 3, -1}, []int{2, 3, -1, -1})) // false

	// Test case 3
	fmt.Println(validateBinaryTreeNodes(2, []int{1, 0}, []int{-1, -1})) // false

	// Test case 4
	fmt.Println(validateBinaryTreeNodes(6, []int{1, -1, -1, 4, -1, -1}, []int{2, -1, -1, 5, -1, -1})) // false
}

// Time: O(n) where n = number of nodes
// Space: O(n) for in-degree and visited arrays
func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	// Track in-degree of each node (how many parents)
	inDegree := make([]int, n)
	for i := 0; i < n; i++ {
		if leftChild[i] != -1 {
			inDegree[leftChild[i]]++
			if inDegree[leftChild[i]] > 1 {
				return false
			}
		}
		if rightChild[i] != -1 {
			inDegree[rightChild[i]]++
			if inDegree[rightChild[i]] > 1 {
				return false
			}
		}
	}

	// Find root (node with in-degree 0)
	root := -1
	for i := 0; i < n; i++ {
		if inDegree[i] == 0 {
			if root != -1 {
				return false // more than one root
			}
			root = i
		}
	}
	if root == -1 {
		return false // cycle (no root)
	}

	// BFS/DFS from root to verify all nodes reachable
	visited := make([]bool, n)
	var dfs func(int)
	dfs = func(node int) {
		if node == -1 || visited[node] {
			return
		}
		visited[node] = true
		dfs(leftChild[node])
		dfs(rightChild[node])
	}
	dfs(root)

	// All nodes must be visited
	for _, v := range visited {
		if !v {
			return false
		}
	}

	return true
}
```

## 1362 — Closest Divisors

```go
package main

// LeetCode #1362: Closest Divisors
// https://leetcode.com/problems/closest-divisors/
// Difficulty: Medium

import "fmt"
import "math"

func main() {
	// Test case 1
	fmt.Println(closestDivisors(8)) // [3,3]

	// Test case 2
	fmt.Println(closestDivisors(123)) // [5,25]

	// Test case 3
	fmt.Println(closestDivisors(999)) // [25,40]

	// Test case 4
	fmt.Println(closestDivisors(1)) // [1,2]
}

// Time: O(sqrt(num)) for finding divisors
// Space: O(1)
func closestDivisors(num int) []int {
	// Check num+1 and num+2 for closest divisor pair
	result := []int{0, 0}
	minDiff := math.MaxInt32

	for n := num + 1; n <= num+2; n++ {
		for i := int(math.Sqrt(float64(n))); i >= 1; i-- {
			if n%i == 0 {
				j := n / i
				diff := j - i
				if diff < minDiff {
					minDiff = diff
					result[0] = i
					result[1] = j
				}
				break
			}
		}
	}

	return result
}
```

## 1364 — Number Of Trusted Contacts Of A Customer

```go
package main

// LeetCode #1364: Number of Trusted Contacts of a Customer
// https://leetcode.com/problems/number-of-trusted-contacts-of-a-customer/
// Difficulty: Medium

import "fmt"
import "sort"

func main() {
	invoices := []struct {
		invoiceID  int
		customerID int
		price      int
	}{
		{44, 2, 100},
		{55, 8, 120},
		{66, 9, 150},
	}
	customers := []struct {
		customerID   int
		customerName string
		email        string
	}{
		{1, "Alice", "alice@leetcode.com"},
		{2, "Bob", "bob@leetcode.com"},
		{13, "John", "john@leetcode.com"},
		{6, "Alex", "alex@leetcode.com"},
	}
	contacts := []struct {
		userID       int
		contactEmail string
		trusted      bool
	}{
		{1, "bob@leetcode.com", true},
		{1, "john@leetcode.com", true},
		{1, "jane@leetcode.com", false},
		{2, "alice@leetcode.com", true},
		{2, "john@leetcode.com", false},
	}

	result := trustedContacts(invoices, customers, contacts)
	for _, r := range result {
		fmt.Printf("%d %s %d %d\n", r.invoiceID, r.customerName, r.price, r.trustedCount)
	}
}

type invoiceResult struct {
	invoiceID    int
	customerName string
	price        int
	trustedCount int
}

// Time: O(n log n) for sorting
// Space: O(n)
func trustedContacts(invoices []struct {
	invoiceID  int
	customerID int
	price      int
}, customers []struct {
	customerID   int
	customerName string
	email        string
}, contacts []struct {
	userID       int
	contactEmail string
	trusted      bool
}) []invoiceResult {
	// Map customer emails to IDs for trusted contact lookup
	emailToCustomerID := make(map[string]int)
	for _, c := range customers {
		emailToCustomerID[c.email] = c.customerID
	}

	// Count trusted contacts per customer
	trustedCounts := make(map[int]int)
	for _, c := range contacts {
		if c.trusted {
			if custID, ok := emailToCustomerID[c.contactEmail]; ok {
				trustedCounts[custID]++
			}
		}
	}

	// Map customer IDs to names
	customerNames := make(map[int]string)
	for _, c := range customers {
		customerNames[c.customerID] = c.customerName
	}

	// Sort invoices by invoice ID
	sort.Slice(invoices, func(i, j int) bool {
		return invoices[i].invoiceID < invoices[j].invoiceID
	})

	var result []invoiceResult
	for _, inv := range invoices {
		name := customerNames[inv.customerID]
		if name == "" {
			name = ""
		}
		result = append(result, invoiceResult{
			inv.invoiceID,
			name,
			inv.price,
			trustedCounts[inv.customerID],
		})
	}

	return result
}
```

## 1366 — Rank Teams By Votes

```go
package main

// LeetCode #1366: Rank Teams by Votes
// https://leetcode.com/problems/rank-teams-by-votes/
// Difficulty: Medium

import "fmt"
import "sort"

func main() {
	// Test case 1
	fmt.Println(rankTeams([]string{"ABC", "ACB", "ABC", "ACB", "ACB"})) // "ACB"

	// Test case 2
	fmt.Println(rankTeams([]string{"WXYZ", "XYZW"})) // "XWYZ"

	// Test case 3
	fmt.Println(rankTeams([]string{"ZMNAGUEDSJYLBOPHRQICWFXTVK"})) // "ZMNAGUEDSJYLBOPHRQICWFXTVK"

	// Test case 4 - single vote
	fmt.Println(rankTeams([]string{"ABC"})) // "ABC"
}

// Time: O(n*m + t^2*log(t)) where n = votes, m = teams per vote, t = unique teams
// Space: O(t^2) for storing vote positions
func rankTeams(votes []string) string {
	if len(votes) == 0 {
		return ""
	}

	teams := votes[0]
	n := len(teams)

	// Count votes for each position for each team
	// score[team][position] = count
	score := make(map[byte][]int)
	for _, t := range []byte(teams) {
		score[t] = make([]int, n)
	}

	for _, vote := range votes {
		for pos, team := range []byte(vote) {
			score[team][pos]++
		}
	}

	// Sort teams
	teamList := []byte(teams)
	sort.Slice(teamList, func(i, j int) bool {
		a, b := teamList[i], teamList[j]
		for pos := 0; pos < n; pos++ {
			if score[a][pos] != score[b][pos] {
				return score[a][pos] > score[b][pos]
			}
		}
		return a < b
	})

	return string(teamList)
}
```

## 1367 — Linked List In Binary Tree

```go
package main

// LeetCode #1367: Linked List in Binary Tree
// https://leetcode.com/problems/linked-list-in-binary-tree/
// Difficulty: Medium

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Test case 1
	head := &ListNode{Val: 4, Next: &ListNode{Val: 2, Next: &ListNode{Val: 8}}}
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 4,
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{Val: 1},
			},
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{Val: 6},
				Right: &TreeNode{
					Val: 8,
					Left: &TreeNode{Val: 1},
					Right: &TreeNode{Val: 3},
				},
			},
		},
	}
	fmt.Println(isSubPath(head, root)) // true

	// Test case 2
	head2 := &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 2, Next: &ListNode{Val: 6}}}}
	fmt.Println(isSubPath(head2, root)) // true

	// Test case 3
	root3 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	head3 := &ListNode{Val: 4}
	fmt.Println(isSubPath(head3, root3)) // false
}

// Time: O(N * L) where N = tree nodes, L = list length
// Space: O(N) for recursion stack
func isSubPath(head *ListNode, root *TreeNode) bool {
	if root == nil {
		return false
	}
	if dfs(head, root) {
		return true
	}
	return isSubPath(head, root.Left) || isSubPath(head, root.Right)
}

func dfs(head *ListNode, node *TreeNode) bool {
	if head == nil {
		return true
	}
	if node == nil {
		return false
	}
	if head.Val != node.Val {
		return false
	}
	return dfs(head.Next, node.Left) || dfs(head.Next, node.Right)
}
```

## 1371 — Find The Longest Substring Containing Vowels In Even Counts

```go
package main

// LeetCode #1371: Find the Longest Substring Containing Vowels in Even Counts
// https://leetcode.com/problems/find-the-longest-substring-containing-vowels-in-even-counts/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(findTheLongestSubstring("eleetminicoworoep")) // 13

	// Test case 2
	fmt.Println(findTheLongestSubstring("leetcodeisgreat")) // 5

	// Test case 3
	fmt.Println(findTheLongestSubstring("bcbcbc")) // 6
}

// Time: O(n) where n = length of string
// Space: O(1) - fixed array of 32 states (2^5)
func findTheLongestSubstring(s string) int {
	// Map bitmask (5 bits for a,e,i,o,u) to first occurrence index
	// State 0 (all vowels even) at position 0
	firstSeen := make([]int, 32)
	for i := range firstSeen {
		firstSeen[i] = -1
	}
	firstSeen[0] = 0

	mask := 0
	maxLen := 0

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'a':
			mask ^= 1 << 0
		case 'e':
			mask ^= 1 << 1
		case 'i':
			mask ^= 1 << 2
		case 'o':
			mask ^= 1 << 3
		case 'u':
			mask ^= 1 << 4
		}

		if firstSeen[mask] == -1 {
			firstSeen[mask] = i + 1
		} else {
			length := i + 1 - firstSeen[mask]
			if length > maxLen {
				maxLen = length
			}
		}
	}

	return maxLen
}
```

## 1372 — Longest Zigzag Path In A Binary Tree

```go
package main

// LeetCode #1372: Longest ZigZag Path in a Binary Tree
// https://leetcode.com/problems/longest-zigzag-path-in-a-binary-tree/
// Difficulty: Medium

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Test case 1
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{Val: 1},
			Right: &TreeNode{
				Val: 1,
				Left: &TreeNode{Val: 1, Right: &TreeNode{Val: 1}},
				Right: &TreeNode{Val: 1},
			},
		},
	}
	fmt.Println(longestZigZag(root)) // 3

	// Test case 2: single node
	fmt.Println(longestZigZag(&TreeNode{Val: 1})) // 0

	// Test case 3
	root2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4}}},
	}
	fmt.Println(longestZigZag(root2)) // 2
}

// Time: O(n) where n = number of nodes
// Space: O(h) for recursion stack
func longestZigZag(root *TreeNode) int {
	maxLen := 0
	var dfs func(*TreeNode, int, bool) // node, current length, isLeft (true = coming from left)
	dfs = func(node *TreeNode, length int, fromLeft bool) {
		if node == nil {
			return
		}
		if length > maxLen {
			maxLen = length
		}
		// Going to left child
		if fromLeft {
			// Continuing zigzag: was going left, now going left = reset
			dfs(node.Left, 1, true)
			// Change direction: was going left, now going right = continue zigzag
			dfs(node.Right, length+1, false)
		} else {
			// Change direction: was going right, now going left = continue zigzag
			dfs(node.Left, length+1, true)
			// Continuing same direction: was going right, now going right = reset
			dfs(node.Right, 1, false)
		}
	}

	dfs(root.Left, 1, true)
	dfs(root.Right, 1, false)

	return maxLen
}
```

## 1375 — Number Of Times Binary String Is Prefix Aligned

```go
package main

// LeetCode #1375: Number of Times Binary String Is Prefix-Aligned
// https://leetcode.com/problems/number-of-times-binary-string-is-prefix-aligned/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(numTimesAllBlue([]int{3, 2, 4, 1, 5})) // 2

	// Test case 2
	fmt.Println(numTimesAllBlue([]int{4, 1, 2, 3})) // 1

	// Test case 3
	fmt.Println(numTimesAllBlue([]int{2, 1, 3})) // 1
}

// Time: O(n) where n = length of flips
// Space: O(1)
func numTimesAllBlue(flips []int) int {
	count := 0
	maxFlip := 0

	for i, f := range flips {
		if f > maxFlip {
			maxFlip = f
		}
		if maxFlip == i+1 {
			count++
		}
	}

	return count
}
```

## 1376 — Time Needed To Inform All Employees

```go
package main

// LeetCode #1376: Time Needed to Inform All Employees
// https://leetcode.com/problems/time-needed-to-inform-all-employees/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(numOfMinutes(6, 2, []int{2, 2, -1, 2, 2, 2}, []int{0, 0, 1, 0, 0, 0})) // 1

	// Test case 2
	fmt.Println(numOfMinutes(7, 6, []int{1, 2, 3, 4, 5, 6, -1}, []int{0, 6, 5, 4, 3, 2, 1})) // 21

	// Test case 3
	fmt.Println(numOfMinutes(1, 0, []int{-1}, []int{0})) // 0
}

// Time: O(n) where n = number of employees
// Space: O(n) for memoization and adjacency list
func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	// Build adjacency list (subordinates)
	subordinates := make([][]int, n)
	for i := 0; i < n; i++ {
		if manager[i] != -1 {
			subordinates[manager[i]] = append(subordinates[manager[i]], i)
		}
	}

	// DFS with memoization
	memo := make([]int, n)
	var dfs func(int) int
	dfs = func(id int) int {
		if memo[id] != 0 {
			return memo[id]
		}
		maxTime := 0
		for _, sub := range subordinates[id] {
			time := dfs(sub)
			if time > maxTime {
				maxTime = time
			}
		}
		memo[id] = informTime[id] + maxTime
		return memo[id]
	}

	return dfs(headID)
}
```

## 1381 — Design A Stack With Increment Operation

```go
package main

// LeetCode #1381: Design a Stack With Increment Operation
// https://leetcode.com/problems/design-a-stack-with-increment-operation/
// Difficulty: Medium

import "fmt"

type CustomStack struct {
	stack []int
	inc   []int // lazy increment array
}

func main() {
	cs := Constructor(3)
	cs.Push(1)
	cs.Push(2)
	fmt.Println(cs.Pop()) // 2
	cs.Push(2)
	cs.Push(3)
	cs.Push(4)
	cs.Increment(5, 100)
	cs.Increment(2, 100)
	fmt.Println(cs.Pop()) // 103
	fmt.Println(cs.Pop()) // 202
	fmt.Println(cs.Pop()) // 201
	fmt.Println(cs.Pop()) // -1

	cs2 := Constructor(2)
	cs2.Push(1)
	cs2.Increment(1, 100)
	cs2.Increment(1, 100)
	fmt.Println(cs2.Pop()) // 201
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{
		stack: make([]int, 0, maxSize),
		inc:   make([]int, 0, maxSize),
	}
}

// Time: O(1)
func (this *CustomStack) Push(x int) {
	if len(this.stack) < cap(this.stack) {
		this.stack = append(this.stack, x)
		this.inc = append(this.inc, 0)
	}
}

// Time: O(1)
func (this *CustomStack) Pop() int {
	if len(this.stack) == 0 {
		return -1
	}
	n := len(this.stack)
	top := this.stack[n-1] + this.inc[n-1]
	if n > 1 {
		this.inc[n-2] += this.inc[n-1]
	}
	this.stack = this.stack[:n-1]
	this.inc = this.inc[:n-1]
	return top
}

// Time: O(1) using lazy increment
func (this *CustomStack) Increment(k int, val int) {
	if len(this.stack) == 0 {
		return
	}
	idx := k - 1
	if idx >= len(this.stack) {
		idx = len(this.stack) - 1
	}
	this.inc[idx] += val
}
```

## 1382 — Balance A Binary Search Tree

```go
package main

// LeetCode #1382: Balance a Binary Search Tree
// https://leetcode.com/problems/balance-a-binary-search-tree/
// Difficulty: Medium

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Test case 1
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{Val: 4},
			},
		},
	}
	result := balanceBST(root)
	fmt.Println(result.Val) // 2

	// Test case 2 - empty
	fmt.Println(balanceBST(nil)) // nil

	// Test case 3
	root2 := &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}
	result2 := balanceBST(root2)
	fmt.Println(result2.Val) // 2
}

// Time: O(n) where n = number of nodes
// Space: O(n) for sorted values array
func balanceBST(root *TreeNode) *TreeNode {
	// Inorder traversal to get sorted values
	values := []int{}
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		values = append(values, node.Val)
		inorder(node.Right)
	}
	inorder(root)

	// Build balanced BST from sorted array
	var build func(int, int) *TreeNode
	build = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}
		mid := left + (right-left)/2
		node := &TreeNode{Val: values[mid]}
		node.Left = build(left, mid-1)
		node.Right = build(mid+1, right)
		return node
	}

	return build(0, len(values)-1)
}
```

## 1386 — Cinema Seat Allocation

```go
package main

// LeetCode #1386: Cinema Seat Allocation
// https://leetcode.com/problems/cinema-seat-allocation/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(maxNumberOfFamilies(3, [][]int{{1, 2}, {1, 3}, {1, 8}, {2, 6}, {3, 1}, {3, 10}})) // 4

	// Test case 2
	fmt.Println(maxNumberOfFamilies(2, [][]int{{2, 1}, {1, 8}, {2, 6}})) // 2

	// Test case 3
	fmt.Println(maxNumberOfFamilies(4, [][]int{{4, 3}, {1, 4}, {4, 6}, {1, 7}})) // 4
}

// Time: O(n) where n = number of reserved seats
// Space: O(k) where k = number of rows with reserved seats
func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
	// Map row to bitmask of reserved seats (columns 2-9, bits 0-7)
	rowMap := make(map[int]int)
	for _, seat := range reservedSeats {
		col := seat[1]
		if col >= 2 && col <= 9 {
			rowMap[seat[0]] |= 1 << (col - 2)
		}
	}

	total := (n - len(rowMap)) * 2 // rows with no reserved seats get 2 families each

	// Patterns for 4-person family groupings
	// Columns: 2-5 (left), 4-7 (middle), 6-9 (right)
	left := 0b11110000   // columns 2,3,4,5 (bits 0-3)
	middle := 0b00111100 // columns 4,5,6,7 (bits 2-5)
	right := 0b00001111  // columns 6,7,8,9 (bits 4-7)

	for _, mask := range rowMap {
		if mask&left == 0 && mask&right == 0 {
			total += 2
		} else if mask&left == 0 || mask&middle == 0 || mask&right == 0 {
			total += 1
		}
	}

	return total
}
```

## 1387 — Sort Integers By The Power Value

```go
package main

// LeetCode #1387: Sort Integers by The Power Value
// https://leetcode.com/problems/sort-integers-by-the-power-value/
// Difficulty: Medium

import "fmt"
import "sort"

func main() {
	// Test case 1
	fmt.Println(getKth(12, 15, 2)) // 13

	// Test case 2
	fmt.Println(getKth(1, 1, 1)) // 1

	// Test case 3
	fmt.Println(getKth(7, 11, 4)) // 7

	// Test case 4
	fmt.Println(getKth(10, 20, 5)) // 13
}

// Time: O(n log n) for sorting
// Space: O(n) for memoization and sorted array
func getKth(lo int, hi int, k int) int {
	memo := make(map[int]int)
	memo[1] = 0

	var power func(int) int
	power = func(x int) int {
		if val, ok := memo[x]; ok {
			return val
		}
		if x%2 == 0 {
			memo[x] = 1 + power(x/2)
		} else {
			memo[x] = 1 + power(3*x+1)
		}
		return memo[x]
	}

	type pair struct {
		val, power int
	}
	pairs := make([]pair, 0, hi-lo+1)
	for i := lo; i <= hi; i++ {
		pairs = append(pairs, pair{i, power(i)})
	}

	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].power != pairs[j].power {
			return pairs[i].power < pairs[j].power
		}
		return pairs[i].val < pairs[j].val
	})

	return pairs[k-1].val
}
```

## 1390 — Four Divisors

```go
package main

// LeetCode #1390: Four Divisors
// https://leetcode.com/problems/four-divisors/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(sumFourDivisors([]int{21, 4, 7})) // 32

	// Test case 2
	fmt.Println(sumFourDivisors([]int{21, 21})) // 64

	// Test case 3
	fmt.Println(sumFourDivisors([]int{1, 2, 3, 4, 5})) // 0
}

// Time: O(n * sqrt(m)) where n = len(nums), m = max value in nums
// Space: O(1)
func sumFourDivisors(nums []int) int {
	total := 0

	for _, num := range nums {
		divCount := 0
		divSum := 0

		for i := 1; i*i <= num; i++ {
			if num%i == 0 {
				divCount++
				divSum += i

				if i*i != num {
					divCount++
					divSum += num / i
				}
			}
			if divCount > 4 {
				break
			}
		}

		if divCount == 4 {
			total += divSum
		}
	}

	return total
}
```

## 1391 — Check If There Is A Valid Path In A Grid

```go
package main

// LeetCode #1391: Check if There is a Valid Path in a Grid
// https://leetcode.com/problems/check-if-there-is-a-valid-path-in-a-grid/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(hasValidPath([][]int{{2, 4, 3}, {6, 5, 2}})) // true

	// Test case 2
	fmt.Println(hasValidPath([][]int{{1, 2, 3}, {4, 5, 6}})) // false

	// Test case 3
	fmt.Println(hasValidPath([][]int{{4, 1}, {6, 1}})) // true

	// Test case 4
	fmt.Println(hasValidPath([][]int{{2}, {2}, {2}, {2}, {2}})) // false
}

// Directions: 0=right, 1=down, 2=left, 3=up
// Each street type defines which directions it connects to
// Street 1: left(2)-right(0)
// Street 2: up(3)-down(1)
// Street 3: left(2)-down(1)
// Street 4: right(0)-down(1)
// Street 5: left(2)-up(3)
// Street 6: right(0)-up(3)

var dirs = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // right, down, left, up

// street[type][direction] = list of compatible outgoing directions
var street = [7][][]int{
	{},
	{{}, {0, 2}, {}, {0, 2}, {}},    // type 1
	{{}, {1, 3}, {1, 3}, {}, {}},    // type 2
	{{}, {}, {1, 2}, {2, 1}, {}},    // type 3: left(2)->down(1), down(1)->left(2)
	{{}, {}, {0, 1}, {1, 0}, {}},    // type 4: right(0)->down(1), down(1)->right(0)
	{{}, {}, {2, 3}, {3, 2}, {}},    // type 5: left(2)->up(3), up(3)->left(2)
	{{}, {}, {0, 3}, {3, 0}, {}},    // type 6: right(0)->up(3), up(3)->right(0)
}

// Time: O(m*n) where m,n = grid dimensions
// Space: O(m*n) for visited array
func hasValidPath(grid [][]int) bool {
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	var dfs func(int, int) bool
	dfs = func(r, c int) bool {
		if r == m-1 && c == n-1 {
			return true
		}
		visited[r][c] = true

		st := grid[r][c]
		// Try each direction the current street allows
		for _, dir := range []int{0, 1, 2, 3} {
			nr, nc := r+dirs[dir][0], c+dirs[dir][1]
			if nr < 0 || nr >= m || nc < 0 || nc >= n || visited[nr][nc] {
				continue
			}
			// Check if current street allows this direction
			if !connects(st, dir) {
				continue
			}
			// Check if next street can connect back
			opp := (dir + 2) % 4
			if !connects(grid[nr][nc], opp) {
				continue
			}
			if dfs(nr, nc) {
				return true
			}
		}
		return false
	}

	return dfs(0, 0)
}

func connects(streetType, dir int) bool {
	switch streetType {
	case 1:
		return dir == 0 || dir == 2
	case 2:
		return dir == 1 || dir == 3
	case 3:
		return dir == 1 || dir == 2
	case 4:
		return dir == 0 || dir == 1
	case 5:
		return dir == 2 || dir == 3
	case 6:
		return dir == 0 || dir == 3
	}
	return false
}
```

## 1393 — Capital Gainloss

```go
package main

// LeetCode #1393: Capital Gain/Loss
// https://leetcode.com/problems/capital-gainloss/
// Difficulty: Medium

import "fmt"

func main() {
	result := capitalGainLoss(
		[]struct {
			stockName    string
			operation    string
			operationDay int
			price        int
		}{
			{"Leetcode", "Buy", 1, 1000},
			{"Corona Masks", "Buy", 2, 10},
			{"Leetcode", "Sell", 5, 9000},
			{"Handbags", "Buy", 17, 30000},
			{"Corona Masks", "Sell", 3, 1010},
			{"Corona Masks", "Buy", 4, 1000},
			{"Corona Masks", "Sell", 5, 500},
			{"Corona Masks", "Buy", 6, 1000},
			{"Corona Masks", "Sell", 7, 500},
			{"Handbags", "Sell", 29, 7000},
		},
	)
	for _, r := range result {
		fmt.Printf("%s %d\n", r.name, r.gain)
	}
}

type gainLoss struct {
	name string
	gain int
}

// Time: O(n) where n = number of transactions
// Space: O(k) where k = number of unique stock names
func capitalGainLoss(stocks []struct {
	stockName    string
	operation    string
	operationDay int
	price        int
}) []gainLoss {
	holdings := make(map[string]int) // net cost of buys
	for _, s := range stocks {
		if s.operation == "Buy" {
			holdings[s.stockName] -= s.price
		} else {
			holdings[s.stockName] += s.price
		}
	}

	var result []gainLoss
	for name, gain := range holdings {
		result = append(result, gainLoss{name, gain})
	}
	return result
}
```

## 1395 — Count Number Of Teams

```go
package main

// LeetCode #1395: Count Number of Teams
// https://leetcode.com/problems/count-number-of-teams/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(numTeams([]int{2, 5, 3, 4, 1})) // 3

	// Test case 2
	fmt.Println(numTeams([]int{2, 1, 3})) // 0

	// Test case 3
	fmt.Println(numTeams([]int{1, 2, 3, 4})) // 4
}

// Time: O(n^2) using middle element approach
// Space: O(1)
func numTeams(rating []int) int {
	n := len(rating)
	count := 0

	for j := 1; j < n-1; j++ {
		// Count elements smaller/larger on left
		leftSmaller, leftLarger := 0, 0
		for i := 0; i < j; i++ {
			if rating[i] < rating[j] {
				leftSmaller++
			} else if rating[i] > rating[j] {
				leftLarger++
			}
		}

		// Count elements smaller/larger on right
		rightSmaller, rightLarger := 0, 0
		for k := j + 1; k < n; k++ {
			if rating[k] < rating[j] {
				rightSmaller++
			} else if rating[k] > rating[j] {
				rightLarger++
			}
		}

		// Increasing: leftSmaller * rightLarger
		// Decreasing: leftLarger * rightSmaller
		count += leftSmaller*rightLarger + leftLarger*rightSmaller
	}

	return count
}
```

## 1396 — Design Underground System

```go
package main

// LeetCode #1396: Design Underground System
// https://leetcode.com/problems/design-underground-system/
// Difficulty: Medium

import "fmt"

type UndergroundSystem struct {
	checkins map[int]checkin
	travels  map[string]travel
}

type checkin struct {
	stationName string
	t           int
}

type travel struct {
	totalTime int
	count     int
}

func main() {
	us := Constructor()

	us.CheckIn(45, "Leyton", 3)
	us.CheckIn(32, "Paradise", 8)
	us.CheckIn(27, "Leyton", 10)
	us.CheckOut(45, "Waterloo", 15)
	us.CheckOut(27, "Waterloo", 20)
	us.CheckOut(32, "Cambridge", 22)
	fmt.Println(us.GetAverageTime("Paradise", "Cambridge")) // 14.0
	fmt.Println(us.GetAverageTime("Leyton", "Waterloo"))    // 11.0
	us.CheckIn(10, "Leyton", 24)
	fmt.Println(us.GetAverageTime("Leyton", "Waterloo"))    // 11.0
	us.CheckOut(10, "Waterloo", 38)
	fmt.Println(us.GetAverageTime("Leyton", "Waterloo"))    // 12.0

	us2 := Constructor()
	us2.CheckIn(1, "A", 1)
	us2.CheckIn(2, "A", 2)
	us2.CheckOut(1, "B", 10)
	us2.CheckOut(2, "B", 20)
	fmt.Println(us2.GetAverageTime("A", "B")) // 13.5
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{
		checkins: make(map[int]checkin),
		travels:  make(map[string]travel),
	}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	this.checkins[id] = checkin{stationName, t}
}

// Time: O(1)
func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	ci := this.checkins[id]
	delete(this.checkins, id)

	key := ci.stationName + "->" + stationName
	tr := this.travels[key]
	tr.totalTime += t - ci.t
	tr.count++
	this.travels[key] = tr
}

// Time: O(1)
func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	key := startStation + "->" + endStation
	tr := this.travels[key]
	return float64(tr.totalTime) / float64(tr.count)
}
```

## 1398 — Customers Who Bought Products A And B But Not C

```go
package main

// LeetCode #1398: Customers Who Bought Products A and B but Not C
// https://leetcode.com/problems/customers-who-bought-products-a-and-b-but-not-c/
// Difficulty: Medium

import "fmt"
import "sort"

func main() {
	// SQL problem - simulating in Go
	result := customersABnotC(
		[]struct {
			customerID   int
			customerName string
		}{
			{1, "Daniel"},
			{2, "Diana"},
			{3, "Elizabeth"},
			{4, "John"},
		},
		[]struct {
			orderID     int
			customerID  int
			productName string
		}{
			{1, 1, "A"},
			{2, 1, "B"},
			{3, 1, "C"},
			{4, 2, "A"},
			{5, 2, "B"},
			{6, 3, "A"},
		},
	)
	for _, r := range result {
		fmt.Printf("%d %s\n", r.id, r.name)
	}
	// Should output: 2 Diana (bought A and B but not C)
	// Note: 1 Daniel bought A and B but also C, excluded
	// Note: 3 Elizabeth bought A but not B, excluded
}

type customerResult struct {
	id   int
	name string
}

// Time: O(n) where n = number of orders
// Space: O(k) where k = number of customers
func customersABnotC(customers []struct {
	customerID   int
	customerName string
}, orders []struct {
	orderID     int
	customerID  int
	productName string
}) []customerResult {
	bought := make(map[int]map[string]bool)
	customerNames := make(map[int]string)

	for _, c := range customers {
		customerNames[c.customerID] = c.customerName
	}

	for _, o := range orders {
		if bought[o.customerID] == nil {
			bought[o.customerID] = make(map[string]bool)
		}
		bought[o.customerID][o.productName] = true
	}

	var result []customerResult
	for _, c := range customers {
		products := bought[c.customerID]
		if products["A"] && products["B"] && !products["C"] {
			result = append(result, customerResult{c.customerID, c.customerName})
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].id < result[j].id
	})

	return result
}
```

## 1400 — Construct K Palindrome Strings

```go
package main

// LeetCode #1400: Construct K Palindrome Strings
// https://leetcode.com/problems/construct-k-palindrome-strings/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(canConstruct("annabelle", 2)) // true

	// Test case 2
	fmt.Println(canConstruct("leetcode", 3)) // false

	// Test case 3
	fmt.Println(canConstruct("true", 4)) // true

	// Test case 4
	fmt.Println(canConstruct("yzyzyzyzyzyzyzy", 2)) // true
}

// Time: O(n) where n = length of string
// Space: O(1) - fixed array of 26
func canConstruct(s string, k int) bool {
	if len(s) < k {
		return false
	}

	freq := make([]int, 26)
	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
	}

	// Count characters with odd frequency
	oddCount := 0
	for _, f := range freq {
		if f%2 == 1 {
			oddCount++
		}
	}

	// Each palindrome can have at most 1 odd-count character
	return oddCount <= k
}
```

## 1401 — Circle And Rectangle Overlapping

```go
package main

// LeetCode #1401: Circle and Rectangle Overlapping
// https://leetcode.com/problems/circle-and-rectangle-overlapping/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(checkOverlap(1, 0, 0, 1, -1, 3, 1)) // true

	// Test case 2
	fmt.Println(checkOverlap(1, 1, 1, -3, -3, 3, 3)) // true

	// Test case 3
	fmt.Println(checkOverlap(1, 0, 0, -1, 0, 0, 1)) // false

	// Test case 4
	fmt.Println(checkOverlap(1, 1, 1, -3, -3, 3, -1)) // true
}

// Time: O(1)
// Space: O(1)
func checkOverlap(radius int, xCenter int, yCenter int, x1 int, y1 int, x2 int, y2 int) bool {
	// Find the closest point on the rectangle to the circle center
	closestX := max(x1, min(x2, xCenter))
	closestY := max(y1, min(y2, yCenter))

	// Calculate distance from circle center to closest point
	dx := xCenter - closestX
	dy := yCenter - closestY

	return dx*dx+dy*dy <= radius*radius
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
```

## 1404 — Number Of Steps To Reduce A Number In Binary Representation To One

```go
package main

// LeetCode #1404: Number of Steps to Reduce a Number in Binary Representation to One
// https://leetcode.com/problems/number-of-steps-to-reduce-a-number-in-binary-representation-to-one/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(numSteps("1101")) // 6

	// Test case 2
	fmt.Println(numSteps("10")) // 1

	// Test case 3
	fmt.Println(numSteps("1")) // 0

	// Test case 4
	fmt.Println(numSteps("1111011110000011100000110001011011110010111001010111110001"))
}

// Time: O(n) where n = length of binary string
// Space: O(1)
func numSteps(s string) int {
	steps := 0
	carry := 0

	for i := len(s) - 1; i > 0; i-- {
		digit := int(s[i]-'0') + carry
		if digit%2 == 1 {
			// Odd: add 1 (which makes it even, two operations: +1 and /2)
			steps += 2
			carry = 1
		} else {
			// Even: divide by 2 (one operation)
			steps++
			// carry stays (if we had carry, 1+0=1, but we divide by 2)
		}
	}

	return steps + carry
}
```

## 1405 — Longest Happy String

```go
package main

// LeetCode #1405: Longest Happy String
// https://leetcode.com/problems/longest-happy-string/
// Difficulty: Medium

import "fmt"
import "sort"

func main() {
	// Test case 1
	fmt.Println(longestDiverseString(1, 1, 7)) // "ccaccbcc" or "ccbccacc"

	// Test case 2
	fmt.Println(longestDiverseString(7, 1, 0)) // "aabaa"

	// Test case 3
	fmt.Println(longestDiverseString(0, 8, 11)) // "ccbccbbccbbccbbccbc"
}

type charCount struct {
	count int
	char  byte
}

// Time: O(a+b+c) - building the result string
// Space: O(1) - constant extra space
func longestDiverseString(a int, b int, c int) string {
	pairs := []charCount{{a, 'a'}, {b, 'b'}, {c, 'c'}}
	result := make([]byte, 0, a+b+c)

	for {
		// Sort by remaining count descending
		sort.Slice(pairs, func(i, j int) bool {
			return pairs[i].count > pairs[j].count
		})

		placed := false
		for i := 0; i < 3; i++ {
			if pairs[i].count == 0 {
				break
			}
			n := len(result)
			// Check if we can place this character
			if n >= 2 && result[n-1] == pairs[i].char && result[n-2] == pairs[i].char {
				continue
			}
			result = append(result, pairs[i].char)
			pairs[i].count--
			placed = true
			break
		}

		if !placed {
			break
		}
	}

	return string(result)
}
```

## 1409 — Queries On A Permutation With Key

```go
package main

// LeetCode #1409: Queries on a Permutation With Key
// https://leetcode.com/problems/queries-on-a-permutation-with-key/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(processQueries([]int{3, 1, 2, 1}, 5)) // [2,1,2,1]

	// Test case 2
	fmt.Println(processQueries([]int{4, 1, 2, 2}, 4)) // [3,1,2,0]

	// Test case 3
	fmt.Println(processQueries([]int{7, 5, 5, 8, 3}, 8)) // [6,5,0,7,5]
}

// Time: O(m*n) where m = len(queries), n = m (since P has m elements)
// Space: O(n) for the permutation
func processQueries(queries []int, m int) []int {
	// Build permutation P = [1, 2, ..., m]
	p := make([]int, m)
	for i := 0; i < m; i++ {
		p[i] = i + 1
	}

	result := make([]int, len(queries))

	for idx, q := range queries {
		// Find position of q in P
		pos := 0
		for p[pos] != q {
			pos++
		}
		result[idx] = pos

		// Move q to front by shifting elements before it
		for i := pos; i > 0; i-- {
			p[i] = p[i-1]
		}
		p[0] = q
	}

	return result
}
```

## 1410 — Html Entity Parser

```go
package main

// LeetCode #1410: HTML Entity Parser
// https://leetcode.com/problems/html-entity-parser/
// Difficulty: Medium

import "fmt"
import "strings"

func main() {
	// Test case 1
	fmt.Println(entityParser("&amp; is an HTML entity but &ambassador; is not."))
	// "& is an HTML entity but &ambassador; is not."

	// Test case 2
	fmt.Println(entityParser("and I quote: &quot;...&quot;"))
	// "and I quote: \"...\""

	// Test case 3
	fmt.Println(entityParser("x &gt; y &amp;&amp; x &lt; y"))
	// "x > y && x < y"

	// Test case 4
	fmt.Println(entityParser("leetcode.com&frasl;problemset&frasl;all"))
	// "leetcode.com/problemset/all"
}

// Time: O(n) where n = length of text
// Space: O(n) for the result
func entityParser(text string) string {
	entities := map[string]string{
		"&quot;":  "\"",
		"&apos;":  "'",
		"&amp;":   "&",
		"&gt;":    ">",
		"&lt;":    "<",
		"&frasl;": "/",
	}

	var sb strings.Builder
	for i := 0; i < len(text); i++ {
		if text[i] == '&' {
			matched := false
			for entity, replacement := range entities {
				if i+len(entity) <= len(text) && text[i:i+len(entity)] == entity {
					sb.WriteString(replacement)
					i += len(entity) - 1
					matched = true
					break
				}
			}
			if !matched {
				sb.WriteByte('&')
			}
		} else {
			sb.WriteByte(text[i])
		}
	}

	return sb.String()
}
```

## 1414 — Find The Minimum Number Of Fibonacci Numbers Whose Sum Is K

```go
package main

// LeetCode #1414: Find the Minimum Number of Fibonacci Numbers Whose Sum Is K
// https://leetcode.com/problems/find-the-minimum-number-of-fibonacci-numbers-whose-sum-is-k/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(findMinFibonacciNumbers(7)) // 2

	// Test case 2
	fmt.Println(findMinFibonacciNumbers(10)) // 2

	// Test case 3
	fmt.Println(findMinFibonacciNumbers(19)) // 3

	// Test case 4
	fmt.Println(findMinFibonacciNumbers(1)) // 1
}

// Time: O(log k) since Fibonacci numbers grow exponentially
// Space: O(1)
func findMinFibonacciNumbers(k int) int {
	// Generate all Fibonacci numbers <= k
	fib := []int{1, 1}
	for fib[len(fib)-1] <= k {
		next := fib[len(fib)-1] + fib[len(fib)-2]
		fib = append(fib, next)
	}

	count := 0
	remaining := k
	for i := len(fib) - 1; i >= 0; i-- {
		if fib[i] <= remaining {
			remaining -= fib[i]
			count++
		}
		if remaining == 0 {
			break
		}
	}

	return count
}
```

## 1415 — The K Th Lexicographical String Of All Happy Strings Of Length N

```go
package main

// LeetCode #1415: The k-th Lexicographical String of All Happy Strings of Length n
// https://leetcode.com/problems/the-k-th-lexicographical-string-of-all-happy-strings-of-length-n/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(getHappyString(1, 3)) // "c"

	// Test case 2
	fmt.Println(getHappyString(1, 4)) // ""

	// Test case 3
	fmt.Println(getHappyString(3, 9)) // "cab"

	// Test case 4
	fmt.Println(getHappyString(2, 7)) // ""
}

// Time: O(n) where n = length of string
// Space: O(n) for recursion
func getHappyString(n int, k int) string {
	// Total happy strings = 3 * 2^(n-1)
	total := 1
	for i := 1; i < n; i++ {
		total *= 2
	}
	total *= 3

	if k > total {
		return ""
	}

	result := make([]byte, n)
	choices := []byte{'a', 'b', 'c'}

	// Each position's choices depend on previous
	prev := byte(0)
	for i := 0; i < n; i++ {
		for _, c := range choices {
			if c == prev {
				continue
			}
			// Count remaining strings if we pick c here
			remaining := 1
			for j := i + 1; j < n; j++ {
				remaining *= 2
			}
			if k > remaining {
				k -= remaining
			} else {
				result[i] = c
				prev = c
				break
			}
		}
	}

	return string(result)
}
```

## 1418 — Display Table Of Food Orders In A Restaurant

```go
package main

// LeetCode #1418: Display Table of Food Orders in a Restaurant
// https://leetcode.com/problems/display-table-of-food-orders-in-a-restaurant/
// Difficulty: Medium

import "fmt"
import "sort"
import "strconv"

func main() {
	// Test case 1
	fmt.Println(displayTable([][]string{
		{"David", "3", "Ceviche"},
		{"Corina", "10", "Beef Burrito"},
		{"David", "3", "Fried Chicken"},
		{"Carla", "5", "Water"},
		{"Carla", "5", "Ceviche"},
		{"Rous", "3", "Ceviche"},
	}))

	// Test case 2
	fmt.Println(displayTable([][]string{
		{"James", "12", "Fried Chicken"},
		{"Ratesh", "12", "Fried Chicken"},
		{"Amadeus", "12", "Fried Chicken"},
		{"Adam", "1", "Canadian Waffles"},
		{"Brianna", "1", "Canadian Waffles"},
	}))
}

// Time: O(n log n) where n = number of orders
// Space: O(n) for maps
func displayTable(orders [][]string) [][]string {
	foodItems := make(map[string]bool)
	tableOrders := make(map[int]map[string]int)

	for _, o := range orders {
		table, _ := strconv.Atoi(o[1])
		food := o[2]

		foodItems[food] = true
		if tableOrders[table] == nil {
			tableOrders[table] = make(map[string]int)
		}
		tableOrders[table][food]++
	}

	// Sort food items (excluding "Table" header)
	foods := make([]string, 0, len(foodItems))
	for f := range foodItems {
		foods = append(foods, f)
	}
	sort.Strings(foods)

	// Sort table numbers
	tables := make([]int, 0, len(tableOrders))
	for t := range tableOrders {
		tables = append(tables, t)
	}
	sort.Ints(tables)

	// Build result
	result := make([][]string, 0, len(tables)+1)
	header := make([]string, 0, len(foods)+1)
	header = append(header, "Table")
	header = append(header, foods...)
	result = append(result, header)

	for _, t := range tables {
		row := make([]string, 0, len(foods)+1)
		row = append(row, strconv.Itoa(t))
		for _, f := range foods {
			row = append(row, strconv.Itoa(tableOrders[t][f]))
		}
		result = append(result, row)
	}

	return result
}
```

## 1419 — Minimum Number Of Frogs Croaking

```go
package main

// LeetCode #1419: Minimum Number of Frogs Croaking
// https://leetcode.com/problems/minimum-number-of-frogs-croaking/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(minNumberOfFrogs("croakcroak")) // 1

	// Test case 2
	fmt.Println(minNumberOfFrogs("crcoakroak")) // 2

	// Test case 3
	fmt.Println(minNumberOfFrogs("croakcrook")) // -1

	// Test case 4
	fmt.Println(minNumberOfFrogs("croakcroa")) // -1
}

// Time: O(n) where n = length of croakOfFrogs
// Space: O(1)
func minNumberOfFrogs(croakOfFrogs string) int {
	// Track count of each character in the sequence c->r->o->a->k
	count := make([]int, 5)
	active := 0
	maxFrogs := 0

	for _, ch := range croakOfFrogs {
		switch ch {
		case 'c':
			count[0]++
			active++
			if active > maxFrogs {
				maxFrogs = active
			}
		case 'r':
			if count[0] <= count[1] {
				return -1
			}
			count[1]++
		case 'o':
			if count[1] <= count[2] {
				return -1
			}
			count[2]++
		case 'a':
			if count[2] <= count[3] {
				return -1
			}
			count[3]++
		case 'k':
			if count[3] <= count[4] {
				return -1
			}
			count[4]++
			active--
		default:
			return -1
		}
	}

	// All counts must be equal (complete "croak" sequences)
	if count[0] != count[1] || count[1] != count[2] || count[2] != count[3] || count[3] != count[4] {
		return -1
	}

	return maxFrogs
}
```

## 1423 — Maximum Points You Can Obtain From Cards

```go
package main

// LeetCode #1423: Maximum Points You Can Obtain from Cards
// https://leetcode.com/problems/maximum-points-you-can-obtain-from-cards/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(maxScore([]int{1, 2, 3, 4, 5, 6, 1}, 3)) // 12

	// Test case 2
	fmt.Println(maxScore([]int{2, 2, 2}, 2)) // 4

	// Test case 3
	fmt.Println(maxScore([]int{9, 7, 7, 9, 7, 7, 9}, 7)) // 55

	// Test case 4
	fmt.Println(maxScore([]int{1, 1000, 1}, 1)) // 1
}

// Time: O(k) where k = number of cards to take
// Space: O(1)
func maxScore(cardPoints []int, k int) int {
	n := len(cardPoints)

	// Sum first k elements (taking from left)
	sum := 0
	for i := 0; i < k; i++ {
		sum += cardPoints[i]
	}

	maxSum := sum
	// Try taking from right instead
	for i := 0; i < k; i++ {
		// Remove one from left, add one from right
		sum = sum - cardPoints[k-1-i] + cardPoints[n-1-i]
		if sum > maxSum {
			maxSum = sum
		}
	}

	return maxSum
}
```

## 1424 — Diagonal Traverse Ii

```go
package main

// LeetCode #1424: Diagonal Traverse II
// https://leetcode.com/problems/diagonal-traverse-ii/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(findDiagonalOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	// [1,4,2,7,5,3,8,6,9]

	// Test case 2
	fmt.Println(findDiagonalOrder([][]int{
		{1, 2, 3, 4, 5},
		{6, 7},
		{8},
		{9, 10, 11},
		{12, 13, 14, 15, 16},
	}))
	// [1,6,2,8,7,3,9,4,12,10,5,13,11,14,15,16]

	// Test case 3
	fmt.Println(findDiagonalOrder([][]int{{1, 2, 3}}))
	// [1,2,3]
}

// Time: O(m*n) where m = number of rows, n = avg columns
// Space: O(m*n) for result
func findDiagonalOrder(nums [][]int) []int {
	// Group elements by (i+j) diagonal index
	// For each diagonal, elements appear in reverse row order
	diagonals := make([][]int, 0)

	for i, row := range nums {
		for j := range row {
			idx := i + j
			if idx >= len(diagonals) {
				diagonals = append(diagonals, []int{})
			}
			// Prepend to maintain reverse row order within diagonal
			diagonals[idx] = append(diagonals[idx], nums[i][j])
		}
	}

	result := make([]int, 0)
	for _, d := range diagonals {
		// Reverse the diagonal (since we prepended, it's in reverse)
		for i := len(d) - 1; i >= 0; i-- {
			result = append(result, d[i])
		}
	}

	return result
}
```

## 1428 — Leftmost Column With At Least A One

```go
package main

// LeetCode #1428: Leftmost Column with at Least a One
// https://leetcode.com/problems/leftmost-column-with-at-least-a-one/
// Difficulty: Medium

import "fmt"

type BinaryMatrix struct {
	grid [][]int
}

func (bm BinaryMatrix) Get(row, col int) int {
	return bm.grid[row][col]
}

func (bm BinaryMatrix) Dimensions() []int {
	if len(bm.grid) == 0 {
		return []int{0, 0}
	}
	return []int{len(bm.grid), len(bm.grid[0])}
}

func main() {
	// Test case 1
	bm := BinaryMatrix{[][]int{{0, 0}, {1, 1}}}
	fmt.Println(leftMostColumnWithOne(bm)) // 0

	// Test case 2
	bm2 := BinaryMatrix{[][]int{{0, 0}, {0, 1}}}
	fmt.Println(leftMostColumnWithOne(bm2)) // 1

	// Test case 3
	bm3 := BinaryMatrix{[][]int{{0, 0}, {0, 0}}}
	fmt.Println(leftMostColumnWithOne(bm3)) // -1
}

// Time: O(m + n) where m = rows, n = cols
// Space: O(1)
func leftMostColumnWithOne(binaryMatrix BinaryMatrix) int {
	dim := binaryMatrix.Dimensions()
	m, n := dim[0], dim[1]

	row, col := 0, n-1
	leftmost := -1

	for row < m && col >= 0 {
		if binaryMatrix.Get(row, col) == 1 {
			leftmost = col
			col--
		} else {
			row++
		}
	}

	return leftmost
}
```

## 1429 — First Unique Number

```go
package main

// LeetCode #1429: First Unique Number
// https://leetcode.com/problems/first-unique-number/
// Difficulty: Medium

import "fmt"

type FirstUnique struct {
	queue []int
	count map[int]int
}

func main() {
	fu := NewFirstUnique([]int{2, 3, 5})
	fmt.Println(fu.ShowFirstUnique()) // 2
	fu.Add(5)
	fmt.Println(fu.ShowFirstUnique()) // 2
	fu.Add(2)
	fmt.Println(fu.ShowFirstUnique()) // 3
	fu.Add(3)
	fmt.Println(fu.ShowFirstUnique()) // -1

	fu2 := NewFirstUnique([]int{7, 7, 7, 7, 7, 7})
	fmt.Println(fu2.ShowFirstUnique()) // -1
	fu2.Add(7)
	fu2.Add(3)
	fmt.Println(fu2.ShowFirstUnique()) // 3
}

func NewFirstUnique(nums []int) FirstUnique {
	fu := FirstUnique{count: make(map[int]int)}
	for _, n := range nums {
		fu.Add(n)
	}
	return fu
}

func (this *FirstUnique) Add(value int) {
	this.count[value]++
	if this.count[value] == 1 {
		this.queue = append(this.queue, value)
	}
}

// Time: O(1) amortized
func (this *FirstUnique) ShowFirstUnique() int {
	for len(this.queue) > 0 && this.count[this.queue[0]] > 1 {
		this.queue = this.queue[1:]
	}
	if len(this.queue) == 0 {
		return -1
	}
	return this.queue[0]
}
```

## 1430 — Check If A String Is A Valid Sequence From Root To Leaves Path In A Binary Tree

```go
package main

// LeetCode #1430: Check If a String Is a Valid Sequence from Root to Leaves Path in a Binary Tree
// https://leetcode.com/problems/check-if-a-string-is-a-valid-sequence-from-root-to-leaves-path-in-a-binary-tree/
// Difficulty: Medium

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Test case 1
	root := &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{Val: 0, Right: &TreeNode{Val: 1}},
			Right: &TreeNode{Val: 1, Left: &TreeNode{Val: 0}},
		},
		Right: &TreeNode{
			Val: 0,
			Left: &TreeNode{Val: 0},
		},
	}
	fmt.Println(isValidSequence(root, []int{0, 1, 0, 1})) // true
	fmt.Println(isValidSequence(root, []int{0, 0, 1}))    // false
	fmt.Println(isValidSequence(root, []int{0, 1, 1}))    // false

	// Test case 2 - empty tree
	fmt.Println(isValidSequence(nil, []int{1})) // false
}

// Time: O(n) where n = depth of the path
// Space: O(h) for recursion
func isValidSequence(root *TreeNode, arr []int) bool {
	return dfs(root, arr, 0)
}

func dfs(node *TreeNode, arr []int, idx int) bool {
	if node == nil || idx >= len(arr) {
		return false
	}
	if node.Val != arr[idx] {
		return false
	}
	if idx == len(arr)-1 {
		return node.Left == nil && node.Right == nil // must be a leaf
	}
	return dfs(node.Left, arr, idx+1) || dfs(node.Right, arr, idx+1)
}
```

## 1432 — Max Difference You Can Get From Changing An Integer

```go
package main

// LeetCode #1432: Max Difference You Can Get From Changing an Integer
// https://leetcode.com/problems/max-difference-you-can-get-from-changing-an-integer/
// Difficulty: Medium

import "fmt"
import "strconv"

func main() {
	// Test case 1
	fmt.Println(maxDiff(555)) // 888

	// Test case 2
	fmt.Println(maxDiff(9)) // 8

	// Test case 3
	fmt.Println(maxDiff(123456)) // 820000

	// Test case 4
	fmt.Println(maxDiff(10000)) // 20000

	// Test case 5
	fmt.Println(maxDiff(9288)) // 8700
}

// Time: O(n) where n = number of digits
// Space: O(n) for string conversion
func maxDiff(num int) int {
	s := strconv.Itoa(num)
	digits := []byte(s)

	// Find max: replace first non-'9' with '9'
	maxDigits := make([]byte, len(digits))
	copy(maxDigits, digits)
	targetMax := byte(0)
	for i := 0; i < len(maxDigits); i++ {
		if maxDigits[i] != '9' {
			targetMax = maxDigits[i]
			break
		}
	}
	if targetMax != 0 {
		for i := 0; i < len(maxDigits); i++ {
			if maxDigits[i] == targetMax {
				maxDigits[i] = '9'
			}
		}
	}
	maxVal, _ := strconv.Atoi(string(maxDigits))

	// Find min: replace first non-'0'/'1' appropriately
	minDigits := make([]byte, len(digits))
	copy(minDigits, digits)

	// If first digit is not '1', replace it with '1'
	// Otherwise find first digit > '1' to replace with '0' (but not leading)
	targetMin := byte(0)
	replacementMin := byte(0)

	if minDigits[0] != '1' {
		targetMin = minDigits[0]
		replacementMin = '1'
	} else {
		for i := 1; i < len(minDigits); i++ {
			if minDigits[i] != '0' && minDigits[i] != '1' {
				targetMin = minDigits[i]
				replacementMin = '0'
				break
			}
		}
	}

	if targetMin != 0 {
		for i := 0; i < len(minDigits); i++ {
			if minDigits[i] == targetMin {
				minDigits[i] = replacementMin
			}
		}
	}
	minVal, _ := strconv.Atoi(string(minDigits))

	return maxVal - minVal
}
```

## 1433 — Check If A String Can Break Another String

```go
package main

// LeetCode #1433: Check If a String Can Break Another String
// https://leetcode.com/problems/check-if-a-string-can-break-another-string/
// Difficulty: Medium

import "fmt"
import "sort"

func main() {
	// Test case 1
	fmt.Println(checkIfCanBreak("abc", "xya")) // true

	// Test case 2
	fmt.Println(checkIfCanBreak("abe", "acd")) // false

	// Test case 3
	fmt.Println(checkIfCanBreak("leetcodee", "interview")) // true

	// Test case 4
	fmt.Println(checkIfCanBreak("a", "b")) // true
}

// Time: O(n log n) for sorting
// Space: O(n) for byte slices
func checkIfCanBreak(s1 string, s2 string) bool {
	b1 := []byte(s1)
	b2 := []byte(s2)
	sort.Slice(b1, func(i, j int) bool { return b1[i] < b1[j] })
	sort.Slice(b2, func(i, j int) bool { return b2[i] < b2[j] })

	// Check if s1 can break s2
	canBreak1 := true
	for i := 0; i < len(b1); i++ {
		if b1[i] < b2[i] {
			canBreak1 = false
			break
		}
	}

	// Check if s2 can break s1
	canBreak2 := true
	for i := 0; i < len(b1); i++ {
		if b2[i] < b1[i] {
			canBreak2 = false
			break
		}
	}

	return canBreak1 || canBreak2
}
```

## 1438 — Longest Continuous Subarray With Absolute Diff Less Than Or Equal To Limit

```go
package main

// LeetCode #1438: Longest Continuous Subarray With Absolute Diff Less Than or Equal to Limit
// https://leetcode.com/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(longestSubarray([]int{8, 2, 4, 7}, 4)) // 2

	// Test case 2
	fmt.Println(longestSubarray([]int{10, 1, 2, 4, 7, 2}, 5)) // 4

	// Test case 3
	fmt.Println(longestSubarray([]int{4, 2, 2, 2, 4, 4, 2, 2}, 0)) // 3

	// Test case 4
	fmt.Println(longestSubarray([]int{1, 5, 6, 7, 8, 10, 6, 5, 6}, 4)) // 5
}

// Time: O(n) where n = len(nums)
// Space: O(n) for deques
func longestSubarray(nums []int, limit int) int {
	// Monotonic deques for tracking min and max in current window
	minDeque := make([]int, 0) // increasing
	maxDeque := make([]int, 0) // decreasing

	left := 0
	maxLen := 0

	for right := 0; right < len(nums); right++ {
		// Maintain min deque (increasing)
		for len(minDeque) > 0 && minDeque[len(minDeque)-1] > nums[right] {
			minDeque = minDeque[:len(minDeque)-1]
		}
		minDeque = append(minDeque, nums[right])

		// Maintain max deque (decreasing)
		for len(maxDeque) > 0 && maxDeque[len(maxDeque)-1] < nums[right] {
			maxDeque = maxDeque[:len(maxDeque)-1]
		}
		maxDeque = append(maxDeque, nums[right])

		// Shrink window if diff > limit
		for maxDeque[0]-minDeque[0] > limit {
			if nums[left] == minDeque[0] {
				minDeque = minDeque[1:]
			}
			if nums[left] == maxDeque[0] {
				maxDeque = maxDeque[1:]
			}
			left++
		}

		length := right - left + 1
		if length > maxLen {
			maxLen = length
		}
	}

	return maxLen
}
```

## 1440 — Evaluate Boolean Expression

```go
package main

// LeetCode #1440: Evaluate Boolean Expression
// https://leetcode.com/problems/evaluate-boolean-expression/
// Difficulty: Medium

import "fmt"

func main() {
	// SQL problem - simulating in Go
	result := evaluateBoolean(
		[]struct {
			leftOperand  string
			operator     string
			rightOperand string
		}{
			{"x", ">", "y"},
			{"x", "<", "y"},
			{"x", "=", "y"},
			{"y", ">", "x"},
			{"y", "<", "x"},
			{"x", "=", "x"},
		},
		[]struct {
			name  string
			value int
		}{
			{"x", 66},
			{"y", 77},
			{"z", 88},
		},
	)
	for _, r := range result {
		fmt.Printf("%s %s %s %t\n", r.left, r.op, r.right, r.value)
	}
}

type evalResult struct {
	left, op, right string
	value           bool
}

// Time: O(n) where n = number of expressions
// Space: O(k) where k = number of variables
func evaluateBoolean(expressions []struct {
	leftOperand  string
	operator     string
	rightOperand string
}, variables []struct {
	name  string
	value int
}) []evalResult {
	varMap := make(map[string]int)
	for _, v := range variables {
		varMap[v.name] = v.value
	}

	var results []evalResult
	for _, e := range expressions {
		leftVal := varMap[e.leftOperand]
		rightVal := varMap[e.rightOperand]
		var val bool

		switch e.operator {
		case ">":
			val = leftVal > rightVal
		case "<":
			val = leftVal < rightVal
		case "=":
			val = leftVal == rightVal
		case "!=":
			val = leftVal != rightVal
		case ">=":
			val = leftVal >= rightVal
		case "<=":
			val = leftVal <= rightVal
		}

		results = append(results, evalResult{e.leftOperand, e.operator, e.rightOperand, val})
	}

	return results
}
```

## 1441 — Build An Array With Stack Operations

```go
package main

// LeetCode #1441: Build an Array With Stack Operations
// https://leetcode.com/problems/build-an-array-with-stack-operations/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(buildArray([]int{1, 3}, 3)) // ["Push","Push","Pop","Push"]

	// Test case 2
	fmt.Println(buildArray([]int{1, 2, 3}, 3)) // ["Push","Push","Push"]

	// Test case 3
	fmt.Println(buildArray([]int{1, 2}, 4)) // ["Push","Push"]
}

// Time: O(n) where n = max number in target
// Space: O(n) for result
func buildArray(target []int, n int) []string {
	result := make([]string, 0, n*2)
	current := 1

	for _, t := range target {
		for current < t {
			result = append(result, "Push", "Pop")
			current++
		}
		result = append(result, "Push")
		current++
	}

	return result
}
```

## 1442 — Count Triplets That Can Form Two Arrays Of Equal Xor

```go
package main

// LeetCode #1442: Count Triplets That Can Form Two Arrays of Equal XOR
// https://leetcode.com/problems/count-triplets-that-can-form-two-arrays-of-equal-xor/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(countTriplets([]int{2, 3, 1, 6, 7})) // 4

	// Test case 2
	fmt.Println(countTriplets([]int{1, 1, 1, 1, 1})) // 10

	// Test case 3
	fmt.Println(countTriplets([]int{1, 2, 3})) // 2

	// Test case 4
	fmt.Println(countTriplets([]int{1})) // 0
}

// Time: O(n^2) where n = len(arr)
// Space: O(1)
func countTriplets(arr []int) int {
	n := len(arr)
	count := 0

	// For pairs (i, k) where arr[i]^...^arr[k] == 0,
	// any j between i+1 and k works, giving (k-i) triplets
	for i := 0; i < n; i++ {
		xor := arr[i]
		for k := i + 1; k < n; k++ {
			xor ^= arr[k]
			if xor == 0 {
				count += k - i
			}
		}
	}

	return count
}
```

## 1443 — Minimum Time To Collect All Apples In A Tree

```go
package main

// LeetCode #1443: Minimum Time to Collect All Apples in a Tree
// https://leetcode.com/problems/minimum-time-to-collect-all-apples-in-a-tree/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(minTime(7, [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}},
		[]bool{false, false, true, false, true, true, false})) // 8

	// Test case 2
	fmt.Println(minTime(7, [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}},
		[]bool{false, false, true, false, false, true, false})) // 6

	// Test case 3
	fmt.Println(minTime(4, [][]int{{0, 2}, {0, 3}, {1, 2}},
		[]bool{false, true, false, false})) // 4
}

// Time: O(n) where n = number of nodes
// Space: O(n) for adjacency list and recursion stack
func minTime(n int, edges [][]int, hasApple []bool) int {
	adj := make([][]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}

	visited := make([]bool, n)
	var dfs func(int) int
	dfs = func(node int) int {
		visited[node] = true
		time := 0

		for _, child := range adj[node] {
			if !visited[child] {
				childTime := dfs(child)
				if childTime > 0 || hasApple[child] {
					time += childTime + 2 // 2 for going down and back up
				}
			}
		}

		return time
	}

	return dfs(0)
}
```

## 1445 — Apples Oranges

```go
package main

// LeetCode #1445: Apples & Oranges
// https://leetcode.com/problems/apples-oranges/
// Difficulty: Medium

import "fmt"

func main() {
	// SQL problem - simulating in Go
	result := applesOranges([]struct {
		saleDate string
		fruit    string
		soldNum  int
	}{
		{"2020-05-01", "apples", 10},
		{"2020-05-01", "oranges", 8},
		{"2020-05-02", "apples", 15},
		{"2020-05-02", "oranges", 15},
		{"2020-05-03", "apples", 20},
		{"2020-05-03", "oranges", 0},
		{"2020-05-04", "apples", 15},
		{"2020-05-04", "oranges", 16},
	})
	for _, r := range result {
		fmt.Printf("%s %d\n", r.date, r.diff)
	}
}

type diffResult struct {
	date string
	diff int
}

// Time: O(n log n) for sorting
// Space: O(n)
func applesOranges(sales []struct {
	saleDate string
	fruit    string
	soldNum  int
}) []diffResult {
	// Group by date
	apples := make(map[string]int)
	oranges := make(map[string]int)

	dateSet := make(map[string]bool)
	for _, s := range sales {
		dateSet[s.saleDate] = true
		if s.fruit == "apples" {
			apples[s.saleDate] += s.soldNum
		} else {
			oranges[s.saleDate] += s.soldNum
		}
	}

	// Sort dates
	dates := make([]string, 0, len(dateSet))
	for d := range dateSet {
		dates = append(dates, d)
	}

	var result []diffResult
	for _, d := range dates {
		diff := apples[d] - oranges[d]
		result = append(result, diffResult{d, diff})
	}

	return result
}
```

## 1447 — Simplified Fractions

```go
package main

// LeetCode #1447: Simplified Fractions
// https://leetcode.com/problems/simplified-fractions/
// Difficulty: Medium

import "fmt"
import "strconv"

func main() {
	// Test case 1
	fmt.Println(simplifiedFractions(2)) // ["1/2"]

	// Test case 2
	fmt.Println(simplifiedFractions(3)) // ["1/2","1/3","2/3"]

	// Test case 3
	fmt.Println(simplifiedFractions(4)) // ["1/2","1/3","1/4","2/3","3/4"]

	// Test case 4
	fmt.Println(simplifiedFractions(1)) // []
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Time: O(n^2 * log(min(i,j))) for generating all fractions
// Space: O(n^2) for result
func simplifiedFractions(n int) []string {
	result := make([]string, 0)

	for denominator := 2; denominator <= n; denominator++ {
		for numerator := 1; numerator < denominator; numerator++ {
			if gcd(numerator, denominator) == 1 {
				result = append(result, strconv.Itoa(numerator)+"/"+strconv.Itoa(denominator))
			}
		}
	}

	return result
}
```

## 1448 — Count Good Nodes In Binary Tree

```go
package main

// LeetCode #1448: Count Good Nodes in Binary Tree
// https://leetcode.com/problems/count-good-nodes-in-binary-tree/
// Difficulty: Medium

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Test case 1
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 3}},
		Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 5}},
	}
	fmt.Println(goodNodes(root)) // 4

	// Test case 2
	root2 := &TreeNode{Val: 3, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 2}}}
	fmt.Println(goodNodes(root2)) // 3

	// Test case 3
	fmt.Println(goodNodes(&TreeNode{Val: 1})) // 1
}

// Time: O(n) where n = number of nodes
// Space: O(h) where h = tree height (recursion stack)
func goodNodes(root *TreeNode) int {
	return countGood(root, root.Val)
}

func countGood(node *TreeNode, maxSoFar int) int {
	if node == nil {
		return 0
	}

	count := 0
	if node.Val >= maxSoFar {
		count = 1
		maxSoFar = node.Val
	}

	count += countGood(node.Left, maxSoFar)
	count += countGood(node.Right, maxSoFar)

	return count
}
```

## 1451 — Rearrange Words In A Sentence

```go
package main

// LeetCode #1451: Rearrange Words in a Sentence
// https://leetcode.com/problems/rearrange-words-in-a-sentence/
// Difficulty: Medium

import "fmt"
import "sort"
import "strings"
import "unicode"

func main() {
	// Test case 1
	fmt.Println(arrangeWords("Leetcode is cool")) // "Is cool leetcode"

	// Test case 2
	fmt.Println(arrangeWords("Keep calm and code on")) // "On and keep calm code"

	// Test case 3
	fmt.Println(arrangeWords("To be or not to be")) // "To be or to be not"
}

type wordInfo struct {
	word   string
	index  int
	length int
}

// Time: O(n log n) where n = number of words
// Space: O(n) for storing words
func arrangeWords(text string) string {
	words := strings.Fields(text)
	if len(words) == 0 {
		return ""
	}

	// Lowercase first word
	runes := []rune(words[0])
	runes[0] = unicode.ToLower(runes[0])
	words[0] = string(runes)

	infos := make([]wordInfo, len(words))
	for i, w := range words {
		infos[i] = wordInfo{w, i, len(w)}
	}

	sort.Slice(infos, func(i, j int) bool {
		if infos[i].length != infos[j].length {
			return infos[i].length < infos[j].length
		}
		return infos[i].index < infos[j].index
	})

	result := make([]string, len(infos))
	for i, info := range infos {
		result[i] = info.word
	}

	// Uppercase first letter
	firstWord := result[0]
	runes = []rune(firstWord)
	runes[0] = unicode.ToUpper(runes[0])
	result[0] = string(runes)

	return strings.Join(result, " ")
}
```

## 1452 — People Whose List Of Favorite Companies Is Not A Subset Of Another List

```go
package main

// LeetCode #1452: People Whose List of Favorite Companies Is Not a Subset of Another List
// https://leetcode.com/problems/people-whose-list-of-favorite-companies-is-not-a-subset-of-another-list/
// Difficulty: Medium

import "fmt"
import "sort"

func main() {
	// Test case 1
	fmt.Println(peopleIndexes([][]string{
		{"leetcode", "google", "facebook"},
		{"google", "microsoft"},
		{"google", "facebook"},
		{"google"},
		{"amazon"},
	}))
	// [0,1,4]

	// Test case 2
	fmt.Println(peopleIndexes([][]string{
		{"leetcode", "google", "facebook"},
		{"leetcode", "amazon"},
		{"facebook", "google"},
	}))
	// [0,1]

	// Test case 3
	fmt.Println(peopleIndexes([][]string{
		{"nxaqhyoprhlhvxojucbwdjggdhc", "jzobpva", "wcrrbiqiminbhyvbzrqffsxznfpy", "tvc", "inm", "qwe", "xdw"},
		{"nxaqhyoprhlhvxojucbwdjggdhc", "jzobpva"},
		{"wcrrbiqiminbhyvbzrqffsxznfpy", "jzobpva"},
		{"wcrrbiqiminbhyvbzrqffsxznfpy"},
		{"jzobpva"},
	}))
	// [0,2]
}

// Time: O(n^2 * m) where n = number of people, m = avg companies per person
// Space: O(n * m) for storing company sets
func peopleIndexes(favoriteCompanies [][]string) []int {
	// Convert each person's companies to a set
	sets := make([]map[string]bool, len(favoriteCompanies))
	for i, companies := range favoriteCompanies {
		sets[i] = make(map[string]bool)
		for _, c := range companies {
			sets[i][c] = true
		}
	}

	// Sort by company list size descending to check larger lists first
	indices := make([]int, len(favoriteCompanies))
	for i := range indices {
		indices[i] = i
	}
	sort.Slice(indices, func(i, j int) bool {
		return len(favoriteCompanies[indices[i]]) > len(favoriteCompanies[indices[j]])
	})

	result := []int{}
	for _, i := range indices {
		isSubset := false
		for _, j := range indices {
			if i == j {
				continue
			}
			if isSubsetOf(sets[i], sets[j]) {
				isSubset = true
				break
			}
		}
		if !isSubset {
			result = append(result, i)
		}
	}

	sort.Ints(result)
	return result
}

func isSubsetOf(a, b map[string]bool) bool {
	if len(a) > len(b) {
		return false
	}
	for k := range a {
		if !b[k] {
			return false
		}
	}
	return true
}
```

## 1454 — Active Users

```go
package main

// LeetCode #1454: Active Users
// https://leetcode.com/problems/active-users/
// Difficulty: Medium

import "fmt"
import "sort"

func main() {
	accounts := []struct {
		id   int
		name string
	}{
		{1, "Daniel"},
		{2, "Bob"},
		{3, "John"},
		{4, "Alice"},
	}
	logins := []struct {
		userID    int
		loginDate string
	}{
		{1, "2020-01-01"}, {1, "2020-01-02"}, {1, "2020-01-03"},
		{1, "2020-01-06"}, {1, "2020-01-07"},
		{2, "2020-01-01"}, {2, "2020-01-02"}, {2, "2020-01-03"},
		{2, "2020-01-05"}, {2, "2020-01-06"}, {2, "2020-01-07"},
		{3, "2020-01-05"}, {3, "2020-01-06"}, {3, "2020-01-07"},
		{3, "2020-01-08"},
		{4, "2020-01-01"},
	}

	result := activeUsers(accounts, logins)
	for _, r := range result {
		fmt.Printf("%d %s\n", r.id, r.name)
	}
}

type activeResult struct {
	id   int
	name string
}

// Time: O(n log n) for sorting logins
// Space: O(n) for maps
func activeUsers(accounts []struct {
	id   int
	name string
}, logins []struct {
	userID    int
	loginDate string
}) []activeResult {
	// Group logins by user
	userLogins := make(map[int][]string)
	for _, l := range logins {
		userLogins[l.userID] = append(userLogins[l.userID], l.loginDate)
	}

	// Sort logins for each user and check for 5+ consecutive days
	activeIDs := make(map[int]bool)
	for userID, dates := range userLogins {
		sort.Strings(dates)

		// Remove duplicates
		unique := make([]string, 0, len(dates))
		seen := make(map[string]bool)
		for _, d := range dates {
			if !seen[d] {
				seen[d] = true
				unique = append(unique, d)
			}
		}

		// Check for 5 consecutive days
		consecutive := 1
		for i := 1; i < len(unique); i++ {
			if isNextDay(unique[i-1], unique[i]) {
				consecutive++
				if consecutive >= 5 {
					activeIDs[userID] = true
					break
				}
			} else {
				consecutive = 1
			}
		}
	}

	var result []activeResult
	for _, a := range accounts {
		if activeIDs[a.id] {
			result = append(result, activeResult{a.id, a.name})
		}
	}

	return result
}

func isNextDay(day1, day2 string) bool {
	// Simple date comparison (YYYY-MM-DD format)
	y1, m1, d1 := parseDate(day1)
	y2, m2, d2 := parseDate(day2)

	// Check if day2 is the next day after day1
	if y1 == y2 && m1 == m2 && d2-d1 == 1 {
		return true
	}
	// Handle month/year boundaries (simplified)
	daysInMonth := map[int]int{1: 31, 2: 28, 3: 31, 4: 30, 5: 31, 6: 30,
		7: 31, 8: 31, 9: 30, 10: 31, 11: 30, 12: 31}
	if y1 == y2 && m2 == m1+1 && d1 == daysInMonth[m1] && d2 == 1 {
		return true
	}
	if y2 == y1+1 && m1 == 12 && m2 == 1 && d1 == 31 && d2 == 1 {
		return true
	}
	return false
}

func parseDate(date string) (int, int, int) {
	var y, m, d int
	fmt.Sscanf(date, "%d-%d-%d", &y, &m, &d)
	return y, m, d
}
```

## 1456 — Maximum Number Of Vowels In A Substring Of Given Length

```go
package main

// LeetCode #1456: Maximum Number of Vowels in a Substring of Given Length
// https://leetcode.com/problems/maximum-number-of-vowels-in-a-substring-of-given-length/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(maxVowels("abciiidef", 3)) // 3

	// Test case 2
	fmt.Println(maxVowels("aeiou", 2)) // 2

	// Test case 3
	fmt.Println(maxVowels("leetcode", 3)) // 2

	// Test case 4
	fmt.Println(maxVowels("rhythms", 4)) // 0
}

func isVowel(ch byte) bool {
	return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u'
}

// Time: O(n) where n = length of string
// Space: O(1)
func maxVowels(s string, k int) int {
	// Count vowels in first window
	count := 0
	for i := 0; i < k; i++ {
		if isVowel(s[i]) {
			count++
		}
	}

	maxCount := count

	// Sliding window
	for i := k; i < len(s); i++ {
		if isVowel(s[i-k]) {
			count--
		}
		if isVowel(s[i]) {
			count++
		}
		if count > maxCount {
			maxCount = count
		}
	}

	return maxCount
}
```

## 1457 — Pseudo Palindromic Paths In A Binary Tree

```go
package main

// LeetCode #1457: Pseudo-Palindromic Paths in a Binary Tree
// https://leetcode.com/problems/pseudo-palindromic-paths-in-a-binary-tree/
// Difficulty: Medium

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Test case 1
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{Val: 3},
			Right: &TreeNode{Val: 1},
		},
		Right: &TreeNode{
			Val: 1,
			Right: &TreeNode{Val: 1},
		},
	}
	fmt.Println(pseudoPalindromicPaths(root)) // 2

	// Test case 2
	root2 := &TreeNode{
		Val: 2,
		Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 1}}},
		Right: &TreeNode{Val: 1},
	}
	fmt.Println(pseudoPalindromicPaths(root2)) // 1

	// Test case 3
	fmt.Println(pseudoPalindromicPaths(&TreeNode{Val: 9})) // 1
}

// Time: O(n) where n = number of nodes
// Space: O(h) for recursion stack
func pseudoPalindromicPaths(root *TreeNode) int {
	count := 0
	freq := make([]int, 10) // node values are 1-9

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		freq[node.Val]++
		if node.Left == nil && node.Right == nil {
			// Check if path is pseudo-palindromic
			oddCount := 0
			for _, f := range freq {
				if f%2 == 1 {
					oddCount++
				}
			}
			if oddCount <= 1 {
				count++
			}
		} else {
			dfs(node.Left)
			dfs(node.Right)
		}
		freq[node.Val]--
	}

	dfs(root)
	return count
}
```

## 1459 — Rectangles Area

```go
package main

// LeetCode #1459: Rectangles Area
// https://leetcode.com/problems/rectangles-area/
// Difficulty: Medium

import "fmt"

func main() {
	// SQL problem - simulating in Go
	result := rectanglesArea([]struct {
		id   int
		x, y int
	}{
		{1, 0, 0},
		{2, 2, 0},
		{3, 0, 2},
		{4, 2, 2},
		{5, 1, 1},
		{6, 3, 1},
		{7, 3, 3},
		{8, 1, 3},
	})
	for _, r := range result {
		fmt.Printf("%d,%d,%d,%d %d\n", r.p1, r.p2, r.p3, r.p4, r.area)
	}
}

type rectResult struct {
	p1, p2, p3, p4, area int
}

// Time: O(n^2) for finding all vertical pairs
// Space: O(n^2) for map
func rectanglesArea(points []struct {
	id   int
	x, y int
}) []rectResult {
	// Map from (y1,y2) -> list of x values that have both these y coords
	type yPair struct{ y1, y2 int }

	// Group points by x
	pointsByX := make(map[int][]int) // x -> [ids]
	coords := make(map[int]struct{ x, y int })

	for _, p := range points {
		pointsByX[p.x] = append(pointsByX[p.x], p.id)
		coords[p.id] = struct{ x, y int }{p.x, p.y}
	}

	verticals := make(map[yPair][]int)

	for _, ids := range pointsByX {
		for i := 0; i < len(ids); i++ {
			for j := i + 1; j < len(ids); j++ {
				yi := coords[ids[i]].y
				yj := coords[ids[j]].y
				var y1, y2 int
				if yi < yj {
					y1, y2 = yi, yj
				} else {
					y1, y2 = yj, yi
				}
				x := coords[ids[i]].x
				key := yPair{y1, y2}
				verticals[key] = append(verticals[key], x)
			}
		}
	}

	var results []rectResult
	for yp, xs := range verticals {
		// Sort xs
		for i := 0; i < len(xs); i++ {
			for j := i + 1; j < len(xs); j++ {
				if xs[j] < xs[i] {
					xs[i], xs[j] = xs[j], xs[i]
				}
			}
		}
		for i := 0; i < len(xs); i++ {
			for j := i + 1; j < len(xs); j++ {
				x1, x2 := xs[i], xs[j]
				area := (x2 - x1) * (yp.y2 - yp.y1)
				if area > 0 {
					var p1, p2, p3, p4 int
					for _, pt := range points {
						if pt.x == x1 && pt.y == yp.y1 {
							p1 = pt.id
						} else if pt.x == x2 && pt.y == yp.y1 {
							p2 = pt.id
						} else if pt.x == x1 && pt.y == yp.y2 {
							p3 = pt.id
						} else if pt.x == x2 && pt.y == yp.y2 {
							p4 = pt.id
						}
					}
					results = append(results, rectResult{p1, p2, p3, p4, area})
				}
			}
		}
	}

	return results
}
```

## 1461 — Check If A String Contains All Binary Codes Of Size K

```go
package main

// LeetCode #1461: Check If a String Contains All Binary Codes of Size K
// https://leetcode.com/problems/check-if-a-string-contains-all-binary-codes-of-size-k/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(hasAllCodes("00110110", 2)) // true

	// Test case 2
	fmt.Println(hasAllCodes("0110", 1)) // true

	// Test case 3
	fmt.Println(hasAllCodes("0110", 2)) // false

	// Test case 4
	fmt.Println(hasAllCodes("000000000010101010111010101001110100110100110010100101001000101010101011", 4)) // true
}

// Time: O(n*k) where n = len(s), or O(n) with bit manipulation
// Space: O(2^k) for storing seen codes
func hasAllCodes(s string, k int) bool {
	if len(s) < k {
		return false
	}

	needed := 1 << k
	seen := make([]bool, needed)
	mask := needed - 1
	hash := 0

	// Compute hash for first k bits
	for i := 0; i < k; i++ {
		hash = (hash << 1) | int(s[i]-'0')
	}
	seen[hash] = true
	count := 1

	// Sliding window with rolling hash
	for i := k; i < len(s); i++ {
		hash = ((hash << 1) & mask) | int(s[i]-'0')
		if !seen[hash] {
			seen[hash] = true
			count++
			if count == needed {
				return true
			}
		}
	}

	return count == needed
}
```

## 1462 — Course Schedule Iv

```go
package main

// LeetCode #1462: Course Schedule IV
// https://leetcode.com/problems/course-schedule-iv/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(checkIfPrerequisite(2, [][]int{{1, 0}}, [][]int{{0, 1}, {1, 0}}))
	// [false, true]

	// Test case 2
	fmt.Println(checkIfPrerequisite(2, [][]int{}, [][]int{{1, 0}, {0, 1}}))
	// [false, false]

	// Test case 3
	fmt.Println(checkIfPrerequisite(5, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}, [][]int{{0, 4}, {4, 0}, {1, 3}, {3, 0}}))
	// [true, false, true, false]
}

// Time: O(n^3) for Floyd-Warshall
// Space: O(n^2) for reachability matrix
func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	// Build adjacency list
	adj := make([][]int, numCourses)
	for _, p := range prerequisites {
		adj[p[0]] = append(adj[p[0]], p[1])
	}

	// Floyd-Warshall for reachability
	reachable := make([][]bool, numCourses)
	for i := range reachable {
		reachable[i] = make([]bool, numCourses)
	}

	for _, p := range prerequisites {
		reachable[p[0]][p[1]] = true
	}

	for k := 0; k < numCourses; k++ {
		for i := 0; i < numCourses; i++ {
			for j := 0; j < numCourses; j++ {
				if reachable[i][k] && reachable[k][j] {
					reachable[i][j] = true
				}
			}
		}
	}

	result := make([]bool, len(queries))
	for i, q := range queries {
		result[i] = reachable[q[0]][q[1]]
	}

	return result
}
```

