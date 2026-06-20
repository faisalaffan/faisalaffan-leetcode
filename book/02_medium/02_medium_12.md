# Medium (Sedang) — Problem ��2238

## 2051 — The Category Of Each Member In The Store

```go
package main

// LeetCode #2051: The Category of Each Member in the Store
// https://leetcode.com/problems/the-category-of-each-member-in-the-store/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

type Member struct {
	ID       int
	VisitCnt int
	Spent    int
}

func categorizeMembers(members []Member) map[int]string {
	sort.Slice(members, func(i, j int) bool {
		return members[i].ID < members[j].ID
	})

	// Find max visits and max spent for normalization
	maxVisits := 0
	maxSpent := 0
	for _, m := range members {
		if m.VisitCnt > maxVisits {
			maxVisits = m.VisitCnt
		}
		if m.Spent > maxSpent {
			maxSpent = m.Spent
		}
	}

	result := make(map[int]string)
	for _, m := range members {
		// Determine category: premium, gold, silver, bronze, basic
		visitRatio := float64(m.VisitCnt) / float64(maxVisits)
		spentRatio := float64(m.Spent) / float64(maxSpent)
		score := visitRatio + spentRatio

		var category string
		if score >= 1.5 {
			category = "premium"
		} else if score >= 1.0 {
			category = "gold"
		} else if score >= 0.5 {
			category = "silver"
		} else if score > 0 {
			category = "bronze"
		} else {
			category = "basic"
		}
		result[m.ID] = category
	}
	return result
}

func main() {
	// Test case 1
	members1 := []Member{
		{1, 10, 1000},
		{2, 5, 500},
		{3, 1, 50},
	}
	result1 := categorizeMembers(members1)
	fmt.Println("Test 1:")
	for id := 1; id <= 3; id++ {
		fmt.Printf("  Member %d: %s\n", id, result1[id])
	}

	// Test case 2
	members2 := []Member{
		{1, 0, 0},
	}
	result2 := categorizeMembers(members2)
	fmt.Println("Test 2: Member 1:", result2[1])
	// Expected: basic
}
```

## 2052 — Minimum Cost To Separate Sentence Into Rows

```go
package main

// LeetCode #2052: Minimum Cost to Separate Sentence Into Rows
// https://leetcode.com/problems/minimum-cost-to-separate-sentence-into-rows/
// Difficulty: Medium [Paid]
// Time: O(n * k) | Space: O(n)

import "fmt"

func minimumCost(sentence string, k int) int {
	words := []string{}
	start := 0
	for i := 0; i <= len(sentence); i++ {
		if i == len(sentence) || sentence[i] == ' ' {
			words = append(words, sentence[start:i])
			start = i + 1
		}
	}

	n := len(words)
	if n == 0 {
		return 0
	}

	// Check for word longer than k
	for _, w := range words {
		if len(w) > k {
			return -1 // impossible
		}
	}

	// DP[i] = min cost to place words[i:]
	dp := make([]int, n+1)
	dp[n] = 0

	for i := n - 1; i >= 0; i-- {
		lineLen := len(words[i])
		best := int(1e9)
		j := i
		for j < n && lineLen <= k {
			if j == n-1 {
				// Last word, no extra cost
				if 0 < best {
					best = 0
				}
			} else {
				cost := (k - lineLen) * (k - lineLen)
				total := cost + dp[j+1]
				if total < best {
					best = total
				}
			}
			j++
			if j < n {
				lineLen += 1 + len(words[j]) // space + word
			}
		}
		dp[i] = best
	}

	return dp[0]
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minimumCost("hello world", 10))
	// Expected: 0 (both words fit on first line, last line no cost)

	// Test case 2
	fmt.Println("Test 2:", minimumCost("a b c d e", 2))
	// Expected: varies based on cost calculation

	// Test case 3
	fmt.Println("Test 3:", minimumCost("hello", 5))
	// Expected: 0
}
```

## 2054 — Two Best Non Overlapping Events

```go
package main

// LeetCode #2054: Two Best Non-Overlapping Events
// https://leetcode.com/problems/two-best-non-overlapping-events/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func maxTwoEvents(events [][]int) int {
	// Sort by end time
	sort.Slice(events, func(i, j int) bool {
		return events[i][1] < events[j][1]
	})

	n := len(events)
	// bestUpTo[i] = max value using events[0..i] (single event, non-overlapping)
	bestUpTo := make([]int, n)
	bestUpTo[0] = events[0][2]
	for i := 1; i < n; i++ {
		if events[i][2] > bestUpTo[i-1] {
			bestUpTo[i] = events[i][2]
		} else {
			bestUpTo[i] = bestUpTo[i-1]
		}
	}

	result := 0
	for i := 0; i < n; i++ {
		// Take event i
		result = max(result, events[i][2])
		// Find last event that ends before this event starts
		lo, hi := 0, i-1
		best := -1
		for lo <= hi {
			mid := lo + (hi-lo)/2
			if events[mid][1] < events[i][0] {
				best = mid
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
		if best != -1 {
			result = max(result, events[i][2]+bestUpTo[best])
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", maxTwoEvents([][]int{{1, 3, 2}, {4, 5, 2}, {2, 4, 3}}))
	// Expected: 4

	// Test case 2
	fmt.Println("Test 2:", maxTwoEvents([][]int{{1, 3, 2}, {4, 5, 2}, {1, 5, 5}}))
	// Expected: 5

	// Test case 3
	fmt.Println("Test 3:", maxTwoEvents([][]int{{1, 5, 3}, {1, 5, 1}, {6, 6, 5}}))
	// Expected: 8
}
```

## 2055 — Plates Between Candles

```go
package main

// LeetCode #2055: Plates Between Candles
// https://leetcode.com/problems/plates-between-candles/
// Difficulty: Medium
// Time: O(n + q) | Space: O(n)

import "fmt"

func platesBetweenCandles(s string, queries [][]int) []int {
	n := len(s)
	// Prefix sum of plates
	prefix := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i]
		if s[i] == '*' {
			prefix[i+1]++
		}
	}

	// Nearest candle to the left
	leftCandle := make([]int, n)
	last := -1
	for i := 0; i < n; i++ {
		if s[i] == '|' {
			last = i
		}
		leftCandle[i] = last
	}

	// Nearest candle to the right
	rightCandle := make([]int, n)
	last = -1
	for i := n - 1; i >= 0; i-- {
		if s[i] == '|' {
			last = i
		}
		rightCandle[i] = last
	}

	result := make([]int, len(queries))
	for i, q := range queries {
		left, right := q[0], q[1]
		l := rightCandle[left]
		r := leftCandle[right]
		if l == -1 || r == -1 || l >= r {
			result[i] = 0
		} else {
			result[i] = prefix[r] - prefix[l]
		}
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", platesBetweenCandles("**|**|***|", [][]int{{2, 5}, {5, 9}}))
	// Expected: [2, 3]

	// Test case 2
	fmt.Println("Test 2:", platesBetweenCandles("***|**|*****|**||**|*", [][]int{{1, 17}, {4, 5}, {14, 17}, {5, 11}, {15, 16}}))
	// Expected: [9, 0, 0, 0, 0]

	// Test case 3
	fmt.Println("Test 3:", platesBetweenCandles("|*|", [][]int{{0, 2}}))
	// Expected: [1]
}
```

## 2058 — Find The Minimum And Maximum Number Of Nodes Between Critical Points

```go
package main

// LeetCode #2058: Find the Minimum and Maximum Number of Nodes Between Critical Points
// https://leetcode.com/problems/find-the-minimum-and-maximum-number-of-nodes-between-critical-points/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func nodesBetweenCriticalPoints(head *ListNode) []int {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return []int{-1, -1}
	}

	prev := head
	curr := head.Next
	pos := 1
	firstCritical := -1
	lastCritical := -1
	minDist := int(1e9)

	for curr.Next != nil {
		if (curr.Val > prev.Val && curr.Val > curr.Next.Val) ||
			(curr.Val < prev.Val && curr.Val < curr.Next.Val) {
			if firstCritical == -1 {
				firstCritical = pos
			} else {
				dist := pos - lastCritical
				if dist < minDist {
					minDist = dist
				}
			}
			lastCritical = pos
		}
		prev = curr
		curr = curr.Next
		pos++
	}

	if firstCritical == lastCritical || firstCritical == -1 {
		return []int{-1, -1}
	}

	return []int{minDist, lastCritical - firstCritical}
}

func main() {
	// Test case 1
	head1 := &ListNode{3, &ListNode{1, nil}}
	fmt.Println("Test 1:", nodesBetweenCriticalPoints(head1))
	// Expected: [-1, -1] (not enough nodes)

	// Test case 2
	head2 := &ListNode{5, &ListNode{3, &ListNode{1, &ListNode{2, &ListNode{5, &ListNode{1, &ListNode{2, nil}}}}}}}
	fmt.Println("Test 2:", nodesBetweenCriticalPoints(head2))
	// Expected: [1, 3]

	// Test case 3
	head3 := &ListNode{1, &ListNode{3, &ListNode{2, &ListNode{2, &ListNode{3, &ListNode{2, &ListNode{2, &ListNode{2, &ListNode{7, nil}}}}}}}}}
	fmt.Println("Test 3:", nodesBetweenCriticalPoints(head3))
	// Expected: [3, 3]
}
```

## 2059 — Minimum Operations To Convert Number

```go
package main

// LeetCode #2059: Minimum Operations to Convert Number
// https://leetcode.com/problems/minimum-operations-to-convert-number/
// Difficulty: Medium
// Time: O(n * range) | Space: O(range)

import "fmt"

func minimumOperations(nums []int, start int, goal int) int {
	visited := make([]bool, 1001)
	queue := []int{start}
	steps := 0

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			curr := queue[i]
			if curr == goal {
				return steps
			}
			for _, v := range nums {
				for _, next := range []int{curr + v, curr - v, curr ^ v} {
					if next == goal {
						return steps + 1
					}
					if next >= 0 && next <= 1000 && !visited[next] {
						visited[next] = true
						queue = append(queue, next)
					}
				}
			}
		}
		queue = queue[size:]
		steps++
	}

	return -1
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minimumOperations([]int{1, 3}, 6, 4))
	// Expected: 2

	// Test case 2
	fmt.Println("Test 2:", minimumOperations([]int{2, 4, 8}, 3, 10))
	// Expected: 3

	// Test case 3
	fmt.Println("Test 3:", minimumOperations([]int{1}, 0, 1001))
	// Expected: -1 (goal outside range, unreachable)

	// Test case 4
	fmt.Println("Test 4:", minimumOperations([]int{2, 8, 16}, 0, 1))
	// Expected: -1
}
```

## 2061 — Number Of Spaces Cleaning Robot Cleaned

```go
package main

// LeetCode #2061: Number of Spaces Cleaning Robot Cleaned
// https://leetcode.com/problems/number-of-spaces-cleaning-robot-cleaned/
// Difficulty: Medium [Paid]
// Time: O(m*n) | Space: O(m*n)

import "fmt"

func numberOfCleanRooms(room [][]int) int {
	m, n := len(room), len(room[0])
	visited := make([][][4]bool, m)
	for i := range visited {
		visited[i] = make([][4]bool, n)
	}

	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // right, down, left, up
	cleaned := make(map[[2]int]bool)
	dir := 0
	r, c := 0, 0
	cleaned[[2]int{0, 0}] = true

	for {
		if visited[r][c][dir] {
			break
		}
		visited[r][c][dir] = true

		// Try to move in current direction
		nextR, nextC := r+dirs[dir][0], c+dirs[dir][1]

		if nextR >= 0 && nextR < m && nextC >= 0 && nextC < n && room[nextR][nextC] == 0 {
			r, c = nextR, nextC
			cleaned[[2]int{r, c}] = true
		} else {
			dir = (dir + 1) % 4
		}
	}

	return len(cleaned)
}

func main() {
	// Test case 1
	room1 := [][]int{{0, 0, 0}, {1, 1, 0}, {0, 0, 0}}
	fmt.Println("Test 1:", numberOfCleanRooms(room1))
	// Expected: 7

	// Test case 2
	room2 := [][]int{{0, 1, 0}, {1, 0, 0}, {0, 0, 0}}
	fmt.Println("Test 2:", numberOfCleanRooms(room2))
	// Expected: 1

	// Test case 3
	room3 := [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
	fmt.Println("Test 3:", numberOfCleanRooms(room3))
	// Expected: 9
}
```

## 2063 — Vowels Of All Substrings

```go
package main

// LeetCode #2063: Vowels of All Substrings
// https://leetcode.com/problems/vowels-of-all-substrings/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func countVowels(word string) int64 {
	n := len(word)
	vowels := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
	var result int64 = 0

	for i := 0; i < n; i++ {
		if vowels[word[i]] {
			// Number of substrings containing word[i]
			// = (i+1) * (n-i)
			result += int64(i+1) * int64(n-i)
		}
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", countVowels("aba"))
	// Expected: 6

	// Test case 2
	fmt.Println("Test 2:", countVowels("abc"))
	// Expected: 3

	// Test case 3
	fmt.Println("Test 3:", countVowels("no"))
	// Expected: 0
}
```

## 2064 — Minimized Maximum Of Products Distributed To Any Store

```go
package main

// LeetCode #2064: Minimized Maximum of Products Distributed to Any Store
// https://leetcode.com/problems/minimized-maximum-of-products-distributed-to-any-store/
// Difficulty: Medium
// Time: O(n log max(quantities)) | Space: O(1)

import "fmt"

func minimizedMaximum(n int, quantities []int) int {
	canDistribute := func(maxProducts int) bool {
		stores := 0
		for _, q := range quantities {
			stores += (q + maxProducts - 1) / maxProducts
			if stores > n {
				return false
			}
		}
		return stores <= n
	}

	left, right := 1, 0
	for _, q := range quantities {
		if q > right {
			right = q
		}
	}

	result := right
	for left <= right {
		mid := left + (right-left)/2
		if canDistribute(mid) {
			result = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minimizedMaximum(6, []int{11, 6}))
	// Expected: 3

	// Test case 2
	fmt.Println("Test 2:", minimizedMaximum(7, []int{15, 10, 10}))
	// Expected: 5

	// Test case 3
	fmt.Println("Test 3:", minimizedMaximum(1, []int{100000}))
	// Expected: 100000
}
```

## 2066 — Account Balance

```go
package main

// LeetCode #2066: Account Balance
// https://leetcode.com/problems/account-balance/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

type Transaction struct {
	AccountID int
	Amount    int
}

func accountBalance(transactions []Transaction) map[int]int {
	balances := make(map[int]int)
	for _, t := range transactions {
		balances[t.AccountID] += t.Amount
	}

	// Get sorted account IDs
	ids := make([]int, 0, len(balances))
	for id := range balances {
		ids = append(ids, id)
	}
	sort.Ints(ids)

	result := make(map[int]int)
	for _, id := range ids {
		result[id] = balances[id]
	}
	return result
}

func main() {
	// Test case 1
	trans1 := []Transaction{
		{1, 100}, {2, 200}, {1, -50}, {2, -100}, {3, 500},
	}
	result1 := accountBalance(trans1)
	fmt.Println("Test 1:")
	for _, id := range []int{1, 2, 3} {
		fmt.Printf("  Account %d: %d\n", id, result1[id])
	}
	// Expected: 1:50, 2:100, 3:500

	// Test case 2
	trans2 := []Transaction{{1, -100}}
	result2 := accountBalance(trans2)
	fmt.Println("Test 2: Account 1:", result2[1])
	// Expected: -100
}
```

## 2067 — Number Of Equal Count Substrings

```go
package main

// LeetCode #2067: Number of Equal Count Substrings
// https://leetcode.com/problems/number-of-equal-count-substrings/
// Difficulty: Medium [Paid]
// Time: O(n * alphabet) | Space: O(alphabet)

import "fmt"

func equalCountSubstrings(s string, count int) int {
	result := 0
	// Try different numbers of distinct characters
	for distinct := 1; distinct <= 26 && distinct*count <= len(s); distinct++ {
		freq := make([]int, 26)
		unique := 0
		exactCount := 0

		for i := 0; i < len(s); i++ {
			idx := int(s[i] - 'a')
			if freq[idx] == 0 {
				unique++
			}
			freq[idx]++
			if freq[idx] == count {
				exactCount++
			}

			// Remove leftmost when window too big
			if i >= distinct*count {
				left := int(s[i-distinct*count] - 'a')
				if freq[left] == count {
					exactCount--
				}
				freq[left]--
				if freq[left] == 0 {
					unique--
				}
			}

			if unique == distinct && exactCount == distinct {
				result++
			}
		}
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", equalCountSubstrings("aaabc", 3))
	// Expected: 3

	// Test case 2
	fmt.Println("Test 2:", equalCountSubstrings("abcd", 1))
	// Expected: 10

	// Test case 3
	fmt.Println("Test 3:", equalCountSubstrings("aabbcc", 2))
	// Expected: 6
}
```

## 2069 — Walking Robot Simulation Ii

```go
package main

// LeetCode #2069: Walking Robot Simulation II
// https://leetcode.com/problems/walking-robot-simulation-ii/
// Difficulty: Medium
// Time: O(1) per operation | Space: O(1)

import "fmt"

type Robot struct {
	w, h   int
	x, y   int
	dir    int
	dirs   [][2]int
	dirStr []string
	moved  bool
}

func Constructor(width int, height int) Robot {
	return Robot{
		w:      width,
		h:      height,
		x:      0,
		y:      0,
		dir:    0,
		dirs:   [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}},
		dirStr: []string{"East", "North", "West", "South"},
		moved:  false,
	}
}

func (r *Robot) Step(num int) {
	r.moved = true
	perimeter := 2 * (r.w + r.h - 2)
	if num >= perimeter {
		num %= perimeter
		if r.x == 0 && r.y == 0 {
			r.dir = 0 // Reset direction to East
		}
	}

	for i := 0; i < num; i++ {
		nx := r.x + r.dirs[r.dir][0]
		ny := r.y + r.dirs[r.dir][1]
		if nx < 0 || nx >= r.w || ny < 0 || ny >= r.h {
			r.dir = (r.dir + 1) % 4
			nx = r.x + r.dirs[r.dir][0]
			ny = r.y + r.dirs[r.dir][1]
		}
		r.x, r.y = nx, ny
	}
}

func (r *Robot) GetPos() []int {
	return []int{r.x, r.y}
}

func (r *Robot) GetDir() string {
	if !r.moved || (r.x == 0 && r.y == 0) {
		return "East"
	}
	return r.dirStr[r.dir]
}

func main() {
	robot := Constructor(6, 3)
	robot.Step(2)
	fmt.Println("Test 1 Pos:", robot.GetPos()) // [2, 0]
	fmt.Println("Test 1 Dir:", robot.GetDir()) // East
	robot.Step(2)
	fmt.Println("Test 2 Pos:", robot.GetPos()) // [4, 0]
	fmt.Println("Test 2 Dir:", robot.GetDir()) // East
	robot.Step(2)
	fmt.Println("Test 3 Pos:", robot.GetPos()) // [5, 1]
	fmt.Println("Test 3 Dir:", robot.GetDir()) // North

	robot2 := Constructor(3, 2)
	robot2.Step(2)
	robot2.Step(3) // goes around perimeter
	fmt.Println("Test 4 Pos:", robot2.GetPos())
	fmt.Println("Test 4 Dir:", robot2.GetDir())
}
```

## 2070 — Most Beautiful Item For Each Query

```go
package main

// LeetCode #2070: Most Beautiful Item for Each Query
// https://leetcode.com/problems/most-beautiful-item-for-each-query/
// Difficulty: Medium
// Time: O((n+q) log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func maximumBeauty(items [][]int, queries []int) []int {
	// Sort items by price
	sort.Slice(items, func(i, j int) bool {
		return items[i][0] < items[j][0]
	})

	// For each price, keep max beauty so far (monotonic)
	type item struct{ price, beauty int }
	filtered := []item{}
	maxBeauty := 0
	for _, it := range items {
		if it[1] > maxBeauty {
			maxBeauty = it[1]
		}
		// Only add if beauty increases (since sorted by price)
		if len(filtered) == 0 || it[1] > filtered[len(filtered)-1].beauty {
			filtered = append(filtered, item{it[0], maxBeauty})
		}
	}

	// Handle queries
	result := make([]int, len(queries))
	for i, q := range queries {
		// Binary search for last item with price <= q
		lo, hi := 0, len(filtered)-1
		best := 0
		for lo <= hi {
			mid := lo + (hi-lo)/2
			if filtered[mid].price <= q {
				best = filtered[mid].beauty
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
		result[i] = best
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", maximumBeauty([][]int{{1, 2}, {3, 2}, {2, 4}, {5, 6}, {3, 5}}, []int{1, 2, 3, 4, 5, 6}))
	// Expected: [2, 4, 5, 5, 6, 6]

	// Test case 2
	fmt.Println("Test 2:", maximumBeauty([][]int{{1, 2}, {1, 2}, {1, 3}, {1, 4}}, []int{1}))
	// Expected: [4]

	// Test case 3
	fmt.Println("Test 3:", maximumBeauty([][]int{{10, 100}}, []int{5, 10, 15}))
	// Expected: [0, 100, 100]
}
```

## 2074 — Reverse Nodes In Even Length Groups

```go
package main

// LeetCode #2074: Reverse Nodes in Even Length Groups
// https://leetcode.com/problems/reverse-nodes-in-even-length-groups/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseEvenLengthGroups(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{Next: head}
	prev := dummy
	groupSize := 1

	for prev.Next != nil {
		// Count nodes in this group
		count := 0
		curr := prev.Next
		for curr != nil && count < groupSize {
			curr = curr.Next
			count++
		}

		if count%2 == 0 {
			// Reverse this group
			first := prev.Next
			curr = first.Next
			for i := 1; i < count; i++ {
				next := curr.Next
				curr.Next = prev.Next
				prev.Next = curr
				first.Next = next
				curr = next
			}
			prev = first
		} else {
			// Move prev to last node of this group
			for i := 0; i < count; i++ {
				prev = prev.Next
			}
		}
		groupSize++
	}

	return dummy.Next
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}

func main() {
	// Test case 1
	head1 := &ListNode{5, &ListNode{2, &ListNode{6, &ListNode{3, &ListNode{9, &ListNode{1, &ListNode{7, &ListNode{3, &ListNode{8, &ListNode{4, nil}}}}}}}}}}
	fmt.Print("Test 1: ")
	printList(reverseEvenLengthGroups(head1))
	// Expected: [5,2,6,3,9,1,4,8,3,7]

	// Test case 2
	head2 := &ListNode{1, &ListNode{1, &ListNode{0, &ListNode{6, &ListNode{5, nil}}}}}
	fmt.Print("Test 2: ")
	printList(reverseEvenLengthGroups(head2))
	// Expected: [1,1,0,6,5]

	// Test case 3
	head3 := &ListNode{1, nil}
	fmt.Print("Test 3: ")
	printList(reverseEvenLengthGroups(head3))
	// Expected: [1]
}
```

## 2075 — Decode The Slanted Ciphertext

```go
package main

// LeetCode #2075: Decode the Slanted Ciphertext
// https://leetcode.com/problems/decode-the-slanted-ciphertext/
// Difficulty: Medium
// Time: O(r*c) | Space: O(r*c)

import (
	"fmt"
	"strings"
)

func decodeCiphertext(encodedText string, rows int) string {
	n := len(encodedText)
	cols := n / rows
	grid := make([][]byte, rows)
	for i := 0; i < rows; i++ {
		grid[i] = []byte(encodedText[i*cols : (i+1)*cols])
	}

	var result strings.Builder
	for startCol := 0; startCol < cols; startCol++ {
		r, c := 0, startCol
		for r < rows && c < cols {
			result.WriteByte(grid[r][c])
			r++
			c++
		}
	}

	// Trim trailing spaces
	s := result.String()
	s = strings.TrimRight(s, " ")
	return s
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", decodeCiphertext("ch   ie   pr", 3))
	// Expected: "cipher"

	// Test case 2
	fmt.Println("Test 2:", decodeCiphertext("iveo    eed   l te   olc", 4))
	// Expected: "i love leetcode"

	// Test case 3
	fmt.Println("Test 3:", decodeCiphertext("coding", 1))
	// Expected: "coding"
}
```

## 2077 — Paths In Maze That Lead To Same Room

```go
package main

// LeetCode #2077: Paths in Maze That Lead to Same Room
// https://leetcode.com/problems/paths-in-maze-that-lead-to-same-room/
// Difficulty: Medium [Paid]
// Time: O(n * deg^2) | Space: O(n + m)

import "fmt"

func numberOfPaths(corridors [][]int, n int) int {
	adj := make([][]int, n+1)
	for _, c := range corridors {
		u, v := c[0], c[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// For each pair of neighbors of a node, check if they are also connected
	// Use adjacency set for O(1) lookup
	adjSet := make([]map[int]bool, n+1)
	for i := 1; i <= n; i++ {
		adjSet[i] = make(map[int]bool)
		for _, v := range adj[i] {
			adjSet[i][v] = true
		}
	}

	count := 0
	for u := 1; u <= n; u++ {
		for _, v := range adj[u] {
			if v > u { // Count each pair once
				for _, w := range adj[v] {
					if w > v && adjSet[u][w] {
						count++
					}
				}
			}
		}
	}
	return count
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", numberOfPaths([][]int{{1, 2}, {5, 1}, {1, 3}, {2, 4}, {4, 5}, {2, 3}}, 5))
	// Expected: 2

	// Test case 2
	fmt.Println("Test 2:", numberOfPaths([][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}}, 4))
	// Expected: 0 (no triangle)

	// Test case 3
	fmt.Println("Test 3:", numberOfPaths([][]int{{1, 2}, {2, 3}, {3, 1}, {1, 4}}, 4))
	// Expected: 1
}
```

## 2079 — Watering Plants

```go
package main

// LeetCode #2079: Watering Plants
// https://leetcode.com/problems/watering-plants/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func wateringPlants(plants []int, capacity int) int {
	steps := 0
	curWater := capacity

	for i := 0; i < len(plants); i++ {
		if curWater < plants[i] {
			// Go back to river (i steps) and come back (i steps)
			steps += i*2 + 1
			curWater = capacity - plants[i]
		} else {
			curWater -= plants[i]
			steps++
		}
	}
	return steps
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", wateringPlants([]int{2, 2, 3, 3}, 5))
	// Expected: 14

	// Test case 2
	fmt.Println("Test 2:", wateringPlants([]int{1, 1, 1, 4, 2, 3}, 4))
	// Expected: 30

	// Test case 3
	fmt.Println("Test 3:", wateringPlants([]int{7, 7, 7, 7, 7, 7, 7}, 8))
	// Expected: 49
}
```

## 2080 — Range Frequency Queries

```go
package main

// LeetCode #2080: Range Frequency Queries
// https://leetcode.com/problems/range-frequency-queries/
// Difficulty: Medium
// Time: O(n) init, O(log n) per query | Space: O(n)

import (
	"fmt"
	"sort"
)

type RangeFreqQuery struct {
	pos map[int][]int
}

func Constructor(arr []int) RangeFreqQuery {
	pos := make(map[int][]int)
	for i, v := range arr {
		pos[v] = append(pos[v], i)
	}
	return RangeFreqQuery{pos: pos}
}

func (rfq *RangeFreqQuery) Query(left int, right int, value int) int {
	positions, ok := rfq.pos[value]
	if !ok {
		return 0
	}
	// First index >= left
	l := sort.Search(len(positions), func(i int) bool {
		return positions[i] >= left
	})
	// First index > right
	r := sort.Search(len(positions), func(i int) bool {
		return positions[i] > right
	})
	return r - l
}

func main() {
	rfq := Constructor([]int{12, 33, 4, 56, 22, 2, 34, 33, 22, 12, 34, 56})
	fmt.Println("Test 1 Query(1, 2, 4):", rfq.Query(1, 2, 4))   // 1
	fmt.Println("Test 2 Query(0, 11, 33):", rfq.Query(0, 11, 33)) // 2
	fmt.Println("Test 3 Query(0, 5, 99):", rfq.Query(0, 5, 99))   // 0

	rfq2 := Constructor([]int{1, 1, 1, 2, 2})
	fmt.Println("Test 4 Query(0, 2, 1):", rfq2.Query(0, 2, 1)) // 3
	fmt.Println("Test 5 Query(3, 4, 2):", rfq2.Query(3, 4, 2)) // 2
}
```

## 2083 — Substrings That Begin And End With The Same Letter

```go
package main

// LeetCode #2083: Substrings That Begin and End With the Same Letter
// https://leetcode.com/problems/substrings-that-begin-and-end-with-the-same-letter/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func numberOfSubstrings(s string) int64 {
	freq := make([]int64, 26)
	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
	}
	var result int64 = 0
	for _, f := range freq {
		result += f * (f + 1) / 2
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", numberOfSubstrings("abc"))
	// Expected: 3

	// Test case 2
	fmt.Println("Test 2:", numberOfSubstrings("abacaba"))
	// Expected: 12

	// Test case 3
	fmt.Println("Test 3:", numberOfSubstrings("aa"))
	// Expected: 3
}
```

## 2084 — Drop Type 1 Orders For Customers With Type 0 Orders

```go
package main

// LeetCode #2084: Drop Type 1 Orders for Customers With Type 0 Orders
// https://leetcode.com/problems/drop-type-1-orders-for-customers-with-type-0-orders/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

type Order struct {
	OrderID  int
	Customer int
	OrderType int // 0 or 1
}

func dropOrders(orders []Order) []Order {
	// Find customers with type 0 orders
	hasTypeZero := make(map[int]bool)
	for _, o := range orders {
		if o.OrderType == 0 {
			hasTypeZero[o.Customer] = true
		}
	}

	// Keep orders that don't need to be dropped
	result := []Order{}
	for _, o := range orders {
		if !(o.OrderType == 1 && hasTypeZero[o.Customer]) {
			result = append(result, o)
		}
	}
	return result
}

func main() {
	// Test case 1
	orders1 := []Order{
		{1, 1, 0},
		{2, 1, 1},
		{3, 2, 1},
		{4, 2, 0},
		{5, 3, 1},
	}
	result1 := dropOrders(orders1)
	fmt.Println("Test 1:")
	for _, o := range result1 {
		fmt.Printf("  Order %d (Customer %d, Type %d)\n", o.OrderID, o.Customer, o.OrderType)
	}
	// Expected: orders 1, 4, 5 (2 dropped because customer 1 has type 0)

	// Test case 2
	orders2 := []Order{
		{1, 1, 0},
		{2, 2, 0},
	}
	result2 := dropOrders(orders2)
	fmt.Println("Test 2:", len(result2))
	// Expected: 2 (no type 1 orders to drop)

	// Test case 3
	fmt.Println("Test 3:", len(dropOrders(nil)))
	// Expected: 0
}
```

## 2086 — Minimum Number Of Food Buckets To Feed The Hamsters

```go
package main

// LeetCode #2086: Minimum Number of Food Buckets to Feed the Hamsters
// https://leetcode.com/problems/minimum-number-of-food-buckets-to-feed-the-hamsters/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func minimumBuckets(hamsters string) int {
	n := len(hamsters)
	buckets := make([]bool, n)
	count := 0

	for i := 0; i < n; i++ {
		if hamsters[i] == 'H' {
			if i > 0 && buckets[i-1] {
				continue
			}
			if i+1 < n && hamsters[i+1] == '.' {
				buckets[i+1] = true
				count++
			} else if i > 0 && hamsters[i-1] == '.' {
				buckets[i-1] = true
				count++
			} else {
				return -1
			}
		}
	}
	return count
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minimumBuckets("H..H"))
	// Expected: 2

	// Test case 2
	fmt.Println("Test 2:", minimumBuckets(".H.H."))
	// Expected: 1

	// Test case 3
	fmt.Println("Test 3:", minimumBuckets("HH"))
	// Expected: -1
}
```

## 2087 — Minimum Cost Homecoming Of A Robot In A Grid

```go
package main

// LeetCode #2087: Minimum Cost Homecoming of a Robot in a Grid
// https://leetcode.com/problems/minimum-cost-homecoming-of-a-robot-in-a-grid/
// Difficulty: Medium
// Time: O(m + n) | Space: O(1)

import "fmt"

func minCost(startPos []int, homePos []int, rowCosts []int, colCosts []int) int {
	cost := 0
	r1, c1 := startPos[0], startPos[1]
	r2, c2 := homePos[0], homePos[1]

	// Move rows
	step := 1
	if r1 > r2 {
		step = -1
	}
	for r := r1 + step; r != r2+step; r += step {
		cost += rowCosts[r]
	}

	// Move columns
	step = 1
	if c1 > c2 {
		step = -1
	}
	for c := c1 + step; c != c2+step; c += step {
		cost += colCosts[c]
	}

	return cost
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minCost([]int{1, 0}, []int{2, 3}, []int{5, 4, 3}, []int{8, 2, 6, 7}))
	// Expected: 18

	// Test case 2
	fmt.Println("Test 2:", minCost([]int{0, 0}, []int{0, 0}, []int{5}, []int{5}))
	// Expected: 0

	// Test case 3
	fmt.Println("Test 3:", minCost([]int{2, 2}, []int{0, 0}, []int{1, 2, 3}, []int{4, 5, 6}))
	// Expected: 15 (row 1 + row 0 + col 1 + col 0 = 2+1+5+4 = 12... let me recalculate)
	// Moving from row 2 to 0: rowCosts[1] + rowCosts[0] = 2 + 1 = 3
	// Moving from col 2 to 0: colCosts[1] + colCosts[0] = 5 + 4 = 9
	// Total = 12
}
```

## 2090 — K Radius Subarray Averages

```go
package main

// LeetCode #2090: K Radius Subarray Averages
// https://leetcode.com/problems/k-radius-subarray-averages/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func getAverages(nums []int, k int) []int {
	n := len(nums)
	result := make([]int, n)
	for i := range result {
		result[i] = -1
	}

	if n < 2*k+1 {
		return result
	}

	// Prefix sum
	prefix := make([]int64, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + int64(nums[i])
	}

	for i := k; i < n-k; i++ {
		sum := prefix[i+k+1] - prefix[i-k]
		result[i] = int(sum / int64(2*k+1))
	}

	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", getAverages([]int{7, 4, 3, 9, 1, 8, 5, 2, 6}, 3))
	// Expected: [-1,-1,-1,5,4,4,-1,-1,-1]

	// Test case 2
	fmt.Println("Test 2:", getAverages([]int{100000}, 0))
	// Expected: [100000]

	// Test case 3
	fmt.Println("Test 3:", getAverages([]int{8}, 100000))
	// Expected: [-1]
}
```

## 2091 — Removing Minimum And Maximum From Array

```go
package main

// LeetCode #2091: Removing Minimum and Maximum From Array
// https://leetcode.com/problems/removing-minimum-and-maximum-from-array/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func minimumDeletions(nums []int) int {
	n := len(nums)
	minIdx, maxIdx := 0, 0
	minVal, maxVal := nums[0], nums[0]

	for i, v := range nums {
		if v < minVal {
			minVal = v
			minIdx = i
		}
		if v > maxVal {
			maxVal = v
			maxIdx = i
		}
	}

	// Possible strategies:
	// 1. Delete both from front: max(minIdx, maxIdx) + 1
	// 2. Delete both from back: n - min(minIdx, maxIdx)
	// 3. Delete one from front, one from back: min(minIdx, maxIdx) + 1 + n - max(minIdx, maxIdx)
	front := max(minIdx, maxIdx) + 1
	back := n - min(minIdx, maxIdx)
	mixed := min(minIdx, maxIdx) + 1 + n - max(minIdx, maxIdx)

	return min(front, min(back, mixed))
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

func main() {
	// Test case 1
	fmt.Println("Test 1:", minimumDeletions([]int{2, 10, 7, 5, 4, 1, 8, 6}))
	// Expected: 5

	// Test case 2
	fmt.Println("Test 2:", minimumDeletions([]int{0, -4, 19, 1, 8, -2, -3, 5}))
	// Expected: 3

	// Test case 3
	fmt.Println("Test 3:", minimumDeletions([]int{101}))
	// Expected: 1
}
```

## 2093 — Minimum Cost To Reach City With Discounts

```go
package main

// LeetCode #2093: Minimum Cost to Reach City With Discounts
// https://leetcode.com/problems/minimum-cost-to-reach-city-with-discounts/
// Difficulty: Medium [Paid]
// Time: O((n+m) * discounts * log(n*discounts)) | Space: O(n * discounts)

import (
	"container/heap"
	"fmt"
	"math"
)

type Edge struct {
	to, cost int
}

type State struct {
	city    int
	discounts int
	cost    int
}

type PriorityQueue []State

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].cost < pq[j].cost }
func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(State)) }
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[:n-1]
	return x
}

func minimumCost(n int, highways [][]int, discounts int) int {
	// Build adjacency list
	adj := make([][]Edge, n)
	for _, h := range highways {
		u, v, c := h[0], h[1], h[2]
		adj[u] = append(adj[u], Edge{v, c})
		adj[v] = append(adj[v], Edge{u, c})
	}

	// dist[city][discountsUsed] = min cost
	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, discounts+1)
		for j := range dist[i] {
			dist[i][j] = math.MaxInt32
		}
	}
	dist[0][0] = 0

	pq := &PriorityQueue{}
	heap.Push(pq, State{0, 0, 0})

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(State)
		if cur.cost > dist[cur.city][cur.discounts] {
			continue
		}

		for _, e := range adj[cur.city] {
			// Without discount
			nc := cur.cost + e.cost
			if nc < dist[e.to][cur.discounts] {
				dist[e.to][cur.discounts] = nc
				heap.Push(pq, State{e.to, cur.discounts, nc})
			}
			// With discount
			if cur.discounts < discounts {
				nc2 := cur.cost + e.cost/2
				if nc2 < dist[e.to][cur.discounts+1] {
					dist[e.to][cur.discounts+1] = nc2
					heap.Push(pq, State{e.to, cur.discounts + 1, nc2})
				}
			}
		}
	}

	result := math.MaxInt32
	for d := 0; d <= discounts; d++ {
		if dist[n-1][d] < result {
			result = dist[n-1][d]
		}
	}
	if result == math.MaxInt32 {
		return -1
	}
	return result
}

func main() {
	// Test case 1
	n1 := 5
	highways1 := [][]int{{0, 1, 4}, {0, 2, 2}, {2, 3, 3}, {1, 4, 1}, {3, 4, 5}}
	discounts1 := 1
	fmt.Println("Test 1:", minimumCost(n1, highways1, discounts1))
	// Expected: 6

	// Test case 2
	n2 := 4
	highways2 := [][]int{{1, 3, 17}, {1, 2, 7}, {3, 2, 5}, {0, 1, 6}, {3, 0, 20}}
	discounts2 := 20
	fmt.Println("Test 2:", minimumCost(n2, highways2, discounts2))
	// Expected: 8

	// Test case 3
	fmt.Println("Test 3:", minimumCost(2, [][]int{{0, 1, 10}}, 0))
	// Expected: 10
}
```

## 2095 — Delete The Middle Node Of A Linked List

```go
package main

// LeetCode #2095: Delete the Middle Node of a Linked List
// https://leetcode.com/problems/delete-the-middle-node-of-a-linked-list/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow, fast := head, head
	var prev *ListNode

	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Delete slow (middle) node
	prev.Next = slow.Next

	return head
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}

func main() {
	// Test case 1
	head1 := &ListNode{1, &ListNode{3, &ListNode{4, &ListNode{7, &ListNode{1, &ListNode{2, &ListNode{6, nil}}}}}}}
	fmt.Print("Test 1: ")
	printList(deleteMiddle(head1))
	// Expected: [1,3,4,1,2,6]

	// Test case 2
	head2 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	fmt.Print("Test 2: ")
	printList(deleteMiddle(head2))
	// Expected: [1,2,4]

	// Test case 3
	head3 := &ListNode{1, nil}
	fmt.Print("Test 3: ")
	printList(deleteMiddle(head3))
	// Expected: []
}
```

## 2096 — Step By Step Directions From A Binary Tree Node To Another

```go
package main

// LeetCode #2096: Step-By-Step Directions From a Binary Tree Node to Another
// https://leetcode.com/problems/step-by-step-directions-from-a-binary-tree-node-to-another/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getDirections(root *TreeNode, startValue int, destValue int) string {
	var findPath func(node *TreeNode, target int, path []byte) ([]byte, bool)
	findPath = func(node *TreeNode, target int, path []byte) ([]byte, bool) {
		if node == nil {
			return nil, false
		}
		if node.Val == target {
			return path, true
		}
		if p, ok := findPath(node.Left, target, append(path, 'L')); ok {
			return p, true
		}
		if p, ok := findPath(node.Right, target, append(path, 'R')); ok {
			return p, true
		}
		return nil, false
	}

	pathToStart, _ := findPath(root, startValue, []byte{})
	pathToDest, _ := findPath(root, destValue, []byte{})

	// Find common prefix length
	i := 0
	for i < len(pathToStart) && i < len(pathToDest) && pathToStart[i] == pathToDest[i] {
		i++
	}

	// Go up from start to LCA (U for each remaining step)
	result := make([]byte, len(pathToStart)-i)
	for j := range result {
		result[j] = 'U'
	}

	// Go down from LCA to dest
	result = append(result, pathToDest[i:]...)

	return string(result)
}

func main() {
	// Test case 1
	root1 := &TreeNode{5,
		&TreeNode{1, &TreeNode{3, nil, nil}, nil},
		&TreeNode{2, &TreeNode{6, nil, nil}, &TreeNode{4, nil, nil}},
	}
	fmt.Println("Test 1:", getDirections(root1, 3, 6))
	// Expected: "UURL"

	// Test case 2
	root2 := &TreeNode{2, &TreeNode{1, nil, nil}, nil}
	fmt.Println("Test 2:", getDirections(root2, 2, 1))
	// Expected: "L"

	// Test case 3
	root3 := &TreeNode{1, nil, nil}
	fmt.Println("Test 3:", getDirections(root3, 1, 1))
	// Expected: ""
}
```

## 2098 — Subsequence Of Size K With The Largest Even Sum

```go
package main

// LeetCode #2098: Subsequence of Size K With the Largest Even Sum
// https://leetcode.com/problems/subsequence-of-size-k-with-the-largest-even-sum/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func largestEvenSum(nums []int, k int) int64 {
	// Separate evens and odds, sort descending
	evens := []int{}
	odds := []int{}
	for _, v := range nums {
		if v%2 == 0 {
			evens = append(evens, v)
		} else {
			odds = append(odds, v)
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(evens)))
	sort.Sort(sort.Reverse(sort.IntSlice(odds)))

	// Prefix sums
	ePrefix := make([]int64, len(evens)+1)
	for i, v := range evens {
		ePrefix[i+1] = ePrefix[i] + int64(v)
	}
	oPrefix := make([]int64, len(odds)+1)
	for i, v := range odds {
		oPrefix[i+1] = oPrefix[i] + int64(v)
	}

	var result int64 = -1
	// Try picking i evens and (k-i) odds
	for i := 0; i <= k && i <= len(evens); i++ {
		j := k - i
		if j > len(odds) {
			continue
		}
		if i == 0 && j%2 != 0 {
			// Need at least one even number to make sum even
			if len(evens) == 0 {
				continue
			}
			// Try picking k odds with at least one even
			// Actually, problem requires subsequence of exactly size k
			// If all chosen numbers are odd and count is odd, sum is odd
			// We need at least 1 even
			_ = evens[0] // reference to avoid unused error
			continue
		}
		// Sum is even if we have even number of odd elements
		if j%2 == 0 {
			sum := ePrefix[i] + oPrefix[j]
			if sum > result {
				result = sum
			}
		}
	}

	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", largestEvenSum([]int{4, 1, 5, 3, 1}, 3))
	// Expected: 12

	// Test case 2
	fmt.Println("Test 2:", largestEvenSum([]int{4, 6, 2}, 3))
	// Expected: 12

	// Test case 3
	fmt.Println("Test 3:", largestEvenSum([]int{1, 3, 5}, 3))
	// Expected: -1 (sum would be 9, odd)
}
```

## 2100 — Find Good Days To Rob The Bank

```go
package main

// LeetCode #2100: Find Good Days to Rob the Bank
// https://leetcode.com/problems/find-good-days-to-rob-the-bank/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func goodDaysToRobBank(security []int, time int) []int {
	n := len(security)
	if n < 2*time+1 {
		return []int{}
	}

	// left[i] = number of consecutive non-increasing days ending at i
	left := make([]int, n)
	for i := 1; i < n; i++ {
		if security[i] <= security[i-1] {
			left[i] = left[i-1] + 1
		}
	}

	// right[i] = number of consecutive non-decreasing days starting at i
	right := make([]int, n)
	for i := n - 2; i >= 0; i-- {
		if security[i] <= security[i+1] {
			right[i] = right[i+1] + 1
		}
	}

	result := []int{}
	for i := time; i < n-time; i++ {
		if left[i] >= time && right[i] >= time {
			result = append(result, i)
		}
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", goodDaysToRobBank([]int{5, 3, 3, 3, 5, 6, 2}, 2))
	// Expected: [2, 3]

	// Test case 2
	fmt.Println("Test 2:", goodDaysToRobBank([]int{1, 1, 1, 1, 1}, 0))
	// Expected: [0, 1, 2, 3, 4]

	// Test case 3
	fmt.Println("Test 3:", goodDaysToRobBank([]int{1, 2, 3, 4, 5, 6}, 2))
	// Expected: []
}
```

## 2101 — Detonate The Maximum Bombs

```go
package main

// LeetCode #2101: Detonate the Maximum Bombs
// https://leetcode.com/problems/detonate-the-maximum-bombs/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n^2)

import "fmt"

func maximumDetonation(bombs [][]int) int {
	n := len(bombs)
	if n == 0 {
		return 0
	}

	// Build adjacency (directed from i to j if i can detonate j)
	adj := make([][]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j {
				dx := bombs[i][0] - bombs[j][0]
				dy := bombs[i][1] - bombs[j][1]
				dist := int64(dx)*int64(dx) + int64(dy)*int64(dy)
				radius := int64(bombs[i][2])
				if dist <= radius*radius {
					adj[i] = append(adj[i], j)
				}
			}
		}
	}

	maxDetonated := 0
	for i := 0; i < n; i++ {
		visited := make([]bool, n)
		queue := []int{i}
		visited[i] = true
		count := 0

		for len(queue) > 0 {
			curr := queue[0]
			queue = queue[1:]
			count++
			for _, next := range adj[curr] {
				if !visited[next] {
					visited[next] = true
					queue = append(queue, next)
				}
			}
		}
		if count > maxDetonated {
			maxDetonated = count
		}
	}
	return maxDetonated
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", maximumDetonation([][]int{{2, 1, 3}, {6, 1, 4}}))
	// Expected: 2

	// Test case 2
	fmt.Println("Test 2:", maximumDetonation([][]int{{1, 1, 5}, {10, 10, 5}}))
	// Expected: 1

	// Test case 3
	fmt.Println("Test 3:", maximumDetonation([][]int{{1, 2, 3}, {2, 3, 1}, {3, 4, 2}, {4, 5, 3}, {5, 6, 4}}))
	// Expected: 5
}
```

## 2104 — Sum Of Subarray Ranges

```go
package main

// LeetCode #2104: Sum of Subarray Ranges
// https://leetcode.com/problems/sum-of-subarray-ranges/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func subArrayRanges(nums []int) int64 {
	n := len(nums)
	var result int64 = 0

	// For each element, count subarrays where it's the max and where it's the min
	// Use monotonic stack to find prev/next greater/smaller

	// As max: prevGreater, nextGreater
	prevGreater := make([]int, n)
	nextGreater := make([]int, n)
	for i := range prevGreater {
		prevGreater[i] = -1
		nextGreater[i] = n
	}

	stack := []int{}
	for i := 0; i < n; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i] {
			nextGreater[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			prevGreater[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	// As min: prevSmaller, nextSmaller
	prevSmaller := make([]int, n)
	nextSmaller := make([]int, n)
	for i := range prevSmaller {
		prevSmaller[i] = -1
		nextSmaller[i] = n
	}

	stack = []int{}
	for i := 0; i < n; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] > nums[i] {
			nextSmaller[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			prevSmaller[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	for i := 0; i < n; i++ {
		// Contribution as max
		leftMax := i - prevGreater[i]
		rightMax := nextGreater[i] - i
		result += int64(nums[i]) * int64(leftMax) * int64(rightMax)

		// Contribution as min
		leftMin := i - prevSmaller[i]
		rightMin := nextSmaller[i] - i
		result -= int64(nums[i]) * int64(leftMin) * int64(rightMin)
	}

	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", subArrayRanges([]int{1, 2, 3}))
	// Expected: 4

	// Test case 2
	fmt.Println("Test 2:", subArrayRanges([]int{1, 3, 3}))
	// Expected: 4

	// Test case 3
	fmt.Println("Test 3:", subArrayRanges([]int{4, -2, -3, 4, 1}))
	// Expected: 59
}
```

## 2105 — Watering Plants Ii

```go
package main

// LeetCode #2105: Watering Plants II
// https://leetcode.com/problems/watering-plants-ii/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func minimumRefill(plants []int, capacityA int, capacityB int) int {
	n := len(plants)
	alice := 0
	bob := n - 1
	waterA := capacityA
	waterB := capacityB
	refills := 0

	for alice < bob {
		// Alice waters
		if waterA < plants[alice] {
			refills++
			waterA = capacityA
		}
		waterA -= plants[alice]
		alice++

		// Bob waters
		if waterB < plants[bob] {
			refills++
			waterB = capacityB
		}
		waterB -= plants[bob]
		bob--
	}

	// Same plant?
	if alice == bob {
		if waterA >= waterB {
			if waterA < plants[alice] {
				refills++
			}
		} else {
			if waterB < plants[bob] {
				refills++
			}
		}
	}

	return refills
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minimumRefill([]int{2, 2, 3, 3}, 5, 5))
	// Expected: 1

	// Test case 2
	fmt.Println("Test 2:", minimumRefill([]int{2, 2, 3, 3}, 3, 4))
	// Expected: 2

	// Test case 3
	fmt.Println("Test 3:", minimumRefill([]int{5}, 10, 8))
	// Expected: 0
}
```

## 2107 — Number Of Unique Flavors After Sharing K Candies

```go
package main

// LeetCode #2107: Number of Unique Flavors After Sharing K Candies
// https://leetcode.com/problems/number-of-unique-flavors-after-sharing-k-candies/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func shareCandies(candies []int, k int) int {
	if k >= len(candies) {
		return 0
	}

	freq := make(map[int]int)
	for i := k; i < len(candies); i++ {
		freq[candies[i]]++
	}

	maxUnique := len(freq)
	for i := k; i < len(candies); i++ {
		// Add candies[i-k] back (give it away)
		freq[candies[i-k]]++
		// Remove candies[i] from the kept set
		freq[candies[i]]--
		if freq[candies[i]] == 0 {
			delete(freq, candies[i])
		}
		if len(freq) > maxUnique {
			maxUnique = len(freq)
		}
	}

	return maxUnique
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", shareCandies([]int{1, 2, 2, 3, 4, 3}, 3))
	// Expected: 3

	// Test case 2
	fmt.Println("Test 2:", shareCandies([]int{1, 1, 2, 3}, 2))
	// Expected: 2

	// Test case 3
	fmt.Println("Test 3:", shareCandies([]int{1, 1, 1, 1}, 2))
	// Expected: 1
}
```

## 2109 — Adding Spaces To A String

```go
package main

// LeetCode #2109: Adding Spaces to a String
// https://leetcode.com/problems/adding-spaces-to-a-string/
// Difficulty: Medium
// Time: O(n + m) | Space: O(n + m)

import (
	"fmt"
	"strings"
)

func addSpaces(s string, spaces []int) string {
	var result strings.Builder
	spaceIdx := 0
	n := len(s)

	for i := 0; i < n; i++ {
		if spaceIdx < len(spaces) && i == spaces[spaceIdx] {
			result.WriteByte(' ')
			spaceIdx++
		}
		result.WriteByte(s[i])
	}

	return result.String()
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", addSpaces("LeetcodeHelpsMeLearn", []int{8, 13, 15}))
	// Expected: "Leetcode Helps Me Learn"

	// Test case 2
	fmt.Println("Test 2:", addSpaces("icodeinpython", []int{1, 5, 7, 9}))
	// Expected: "i code in py thon"

	// Test case 3
	fmt.Println("Test 3:", addSpaces("spacing", []int{0}))
	// Expected: " spacing"
}
```

## 2110 — Number Of Smooth Descent Periods Of A Stock

```go
package main

// LeetCode #2110: Number of Smooth Descent Periods of a Stock
// https://leetcode.com/problems/number-of-smooth-descent-periods-of-a-stock/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func getDescentPeriods(prices []int) int64 {
	n := len(prices)
	var result int64 = 1 // single element
	length := 1

	for i := 1; i < n; i++ {
		if prices[i] == prices[i-1]-1 {
			length++
		} else {
			length = 1
		}
		result += int64(length)
	}

	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", getDescentPeriods([]int{3, 2, 1, 4}))
	// Expected: 7

	// Test case 2
	fmt.Println("Test 2:", getDescentPeriods([]int{8, 6, 7, 7}))
	// Expected: 4

	// Test case 3
	fmt.Println("Test 3:", getDescentPeriods([]int{1}))
	// Expected: 1
}
```

## 2112 — The Airport With The Most Traffic

```go
package main

// LeetCode #2112: The Airport With the Most Traffic
// https://leetcode.com/problems/the-airport-with-the-most-traffic/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func airportWithMostTraffic(flights [][]string) string {
	traffic := make(map[string]int)
	for _, f := range flights {
		departure, arrival := f[0], f[1]
		traffic[departure]++
		traffic[arrival]++
	}

	maxTraffic := 0
	airport := ""
	for a, t := range traffic {
		if t > maxTraffic || (t == maxTraffic && (airport == "" || a < airport)) {
			maxTraffic = t
			airport = a
		}
	}
	return airport
}

func main() {
	// Test case 1
	flights1 := [][]string{{"JFK", "LGA"}, {"JFK", "LAX"}, {"LAX", "SFO"}, {"LGA", "ORD"}}
	fmt.Println("Test 1:", airportWithMostTraffic(flights1))
	// Expected: "JFK" (3 flights)

	// Test case 2
	flights2 := [][]string{{"A", "B"}, {"B", "C"}, {"C", "A"}}
	fmt.Println("Test 2:", airportWithMostTraffic(flights2))
	// Expected: "A" (tie, alphabetically first)

	// Test case 3
	fmt.Println("Test 3:", airportWithMostTraffic([][]string{{"X", "Y"}}))
	// Expected: "X" (tie, alphabetically: X < Y)
}
```

## 2113 — Elements In Array After Removing And Replacing Elements

```go
package main

// LeetCode #2113: Elements in Array After Removing and Replacing Elements
// https://leetcode.com/problems/elements-in-array-after-removing-and-replacing-elements/
// Difficulty: Medium [Paid]
// Time: O(n + q) | Space: O(1)

import "fmt"

func elementInArray(nums []int, queries [][]int) []int {
	n := len(nums)
	result := make([]int, len(queries))

	for i, q := range queries {
		time := q[0] % (2 * n)
		idx := q[1]
		result[i] = -1

		if time < n {
			// Removal phase: first 'time' elements removed
			// Remaining: nums[time:]
			if idx+time < n {
				result[i] = nums[idx+time]
			}
		} else if time > n {
			// Replacement phase: k = time-n (1..n-1)
			// Elements nums[n-k..n-1] placed back at front
			k := time - n
			if idx < k {
				result[i] = nums[n-k+idx]
			}
		}
		// time == n: array empty, result stays -1
	}
	return result
}

func main() {
	// Test case 1
	nums1 := []int{1, 2, 3}
	queries1 := [][]int{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}, {5, 0}, {5, 1}}
	fmt.Println("Test 1:", elementInArray(nums1, queries1))
	// t=0: [1,2,3] → 1
	// t=1: [2,3] → 2
	// t=2: [3] → 3
	// t=3: [] → -1
	// t=4: [3] → 3
	// t=5: [2,3] → 2, 3
	// Expected: [1, 2, 3, -1, 3, 2, 3]

	// Test case 2
	nums2 := []int{5}
	queries2 := [][]int{{0, 0}, {1, 0}}
	fmt.Println("Test 2:", elementInArray(nums2, queries2))
	// t=0: [5] → 5
	// t=1: [] → -1
	// Expected: [5, -1]

	// Test case 3
	nums3 := []int{10, 20}
	queries3 := [][]int{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {0, 1}}
	fmt.Println("Test 3:", elementInArray(nums3, queries3))
	// Expected: [10, 20, -1, 20, 20]
}
```

## 2115 — Find All Possible Recipes From Given Supplies

```go
package main

// LeetCode #2115: Find All Possible Recipes from Given Supplies
// https://leetcode.com/problems/find-all-possible-recipes-from-given-supplies/
// Difficulty: Medium
// Time: O(n + m + s) | Space: O(n + m + s)

import "fmt"

func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	supplySet := make(map[string]bool)
	for _, s := range supplies {
		supplySet[s] = true
	}

	recipeIdx := make(map[string]int)
	for i, r := range recipes {
		recipeIdx[r] = i
	}

	// indegree for recipes (how many ingredients still needed)
	indegree := make([]int, len(recipes))
	// For each recipe ingredient, which recipes need it
	graph := make(map[string][]int)
	for i, ing := range ingredients {
		for _, ig := range ing {
			if !supplySet[ig] {
				indegree[i]++
				graph[ig] = append(graph[ig], i)
			}
		}
	}

	queue := []int{}
	for i, d := range indegree {
		if d == 0 {
			queue = append(queue, i)
		}
	}

	result := []string{}
	for len(queue) > 0 {
		r := queue[0]
		queue = queue[1:]
		result = append(result, recipes[r])

		// This recipe is now a supply for other recipes
		for _, next := range graph[recipes[r]] {
			indegree[next]--
			if indegree[next] == 0 {
				queue = append(queue, next)
			}
		}
	}

	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", findAllRecipes(
		[]string{"bread"},
		[][]string{{"yeast", "flour"}},
		[]string{"yeast", "flour", "corn"},
	))
	// Expected: ["bread"]

	// Test case 2
	fmt.Println("Test 2:", findAllRecipes(
		[]string{"bread", "sandwich"},
		[][]string{{"yeast", "flour"}, {"bread", "meat"}},
		[]string{"yeast", "flour", "meat"},
	))
	// Expected: ["bread", "sandwich"]

	// Test case 3
	fmt.Println("Test 3:", findAllRecipes(
		[]string{"bread", "sandwich", "burger"},
		[][]string{{"yeast", "flour"}, {"bread", "meat"}, {"sandwich", "meat", "bread"}},
		[]string{"yeast", "flour", "meat"},
	))
	// Expected: ["bread", "sandwich", "burger"]
}
```

## 2116 — Check If A Parentheses String Can Be Valid

```go
package main

// LeetCode #2116: Check if a Parentheses String Can Be Valid
// https://leetcode.com/problems/check-if-a-parentheses-string-can-be-valid/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func canBeValid(s string, locked string) bool {
	n := len(s)
	if n%2 != 0 {
		return false
	}

	// Left to right: check for excess ')'
	balance := 0
	flexible := 0
	for i := 0; i < n; i++ {
		if locked[i] == '0' {
			flexible++
		} else if s[i] == '(' {
			balance++
		} else {
			balance--
		}
		if balance+flexible < 0 {
			return false
		}
	}

	// Right to left: check for excess '('
	balance = 0
	flexible = 0
	for i := n - 1; i >= 0; i-- {
		if locked[i] == '0' {
			flexible++
		} else if s[i] == ')' {
			balance++
		} else {
			balance--
		}
		if balance+flexible < 0 {
			return false
		}
	}

	return true
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", canBeValid("))()))", "010100"))
	// Expected: true

	// Test case 2
	fmt.Println("Test 2:", canBeValid("()()", "0000"))
	// Expected: true

	// Test case 3
	fmt.Println("Test 3:", canBeValid(")", "0"))
	// Expected: false (odd length)
}
```

## 2120 — Execution Of All Suffix Instructions Staying In A Grid

```go
package main

// LeetCode #2120: Execution of All Suffix Instructions Staying in a Grid
// https://leetcode.com/problems/execution-of-all-suffix-instructions-staying-in-a-grid/
// Difficulty: Medium
// Time: O(m * n) or O(n^2) | Space: O(n)

import "fmt"

func executeInstructions(n int, startPos []int, s string) []int {
	m := len(s)
	result := make([]int, m)

	dirs := map[byte][2]int{
		'L': {0, -1},
		'R': {0, 1},
		'U': {-1, 0},
		'D': {1, 0},
	}

	for i := 0; i < m; i++ {
		r, c := startPos[0], startPos[1]
		count := 0
		for j := i; j < m; j++ {
			dr, dc := dirs[s[j]][0], dirs[s[j]][1]
			r += dr
			c += dc
			if r < 0 || r >= n || c < 0 || c >= n {
				break
			}
			count++
		}
		result[i] = count
	}

	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", executeInstructions(3, []int{0, 1}, "RRDDLU"))
	// Expected: [1, 5, 4, 3, 1, 0]

	// Test case 2
	fmt.Println("Test 2:", executeInstructions(2, []int{1, 1}, "LURD"))
	// Expected: [4, 1, 0, 0]

	// Test case 3
	fmt.Println("Test 3:", executeInstructions(1, []int{0, 0}, "LRUD"))
	// Expected: [0, 0, 0, 0]
}
```

## 2121 — Intervals Between Identical Elements

```go
package main

// LeetCode #2121: Intervals Between Identical Elements
// https://leetcode.com/problems/intervals-between-identical-elements/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func getDistances(arr []int) []int64 {
	n := len(arr)
	// Group indices by value
	groups := make(map[int][]int)
	for i, v := range arr {
		groups[v] = append(groups[v], i)
	}

	result := make([]int64, n)
	for _, indices := range groups {
		m := len(indices)
		if m <= 1 {
			continue
		}
		// Prefix sum of distances
		prefix := make([]int64, m+1)
		for i := 0; i < m; i++ {
			prefix[i+1] = prefix[i] + int64(indices[i])
		}
		for i, pos := range indices {
			// Sum of distances to all other same-value elements
			// Left side: pos * i - prefix[i]
			// Right side: (prefix[m] - prefix[i+1]) - pos * (m-1-i)
			left := int64(pos)*int64(i) - prefix[i]
			right := (prefix[m] - prefix[i+1]) - int64(pos)*int64(m-1-i)
			result[pos] = left + right
		}
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", getDistances([]int{2, 1, 3, 1, 2, 3, 3}))
	// Expected: [4, 2, 7, 2, 4, 4, 5]

	// Test case 2
	fmt.Println("Test 2:", getDistances([]int{10, 5, 10, 10}))
	// Expected: [5, 0, 3, 4]

	// Test case 3
	fmt.Println("Test 3:", getDistances([]int{1, 2, 3}))
	// Expected: [0, 0, 0]
}
```

## 2125 — Number Of Laser Beams In A Bank

```go
package main

// LeetCode #2125: Number of Laser Beams in a Bank
// https://leetcode.com/problems/number-of-laser-beams-in-a-bank/
// Difficulty: Medium
// Time: O(m*n) | Space: O(1)

import "fmt"

func numberOfBeams(bank []string) int {
	prevCount := 0
	result := 0

	for _, row := range bank {
		count := 0
		for _, c := range row {
			if c == '1' {
				count++
			}
		}
		if count > 0 {
			result += prevCount * count
			prevCount = count
		}
	}

	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", numberOfBeams([]string{"011001", "000000", "010100", "001000"}))
	// Expected: 8

	// Test case 2
	fmt.Println("Test 2:", numberOfBeams([]string{"000", "111", "000"}))
	// Expected: 0

	// Test case 3
	fmt.Println("Test 3:", numberOfBeams([]string{"101", "010", "101"}))
	// Expected: 4
}
```

## 2126 — Destroying Asteroids

```go
package main

// LeetCode #2126: Destroying Asteroids
// https://leetcode.com/problems/destroying-asteroids/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func asteroidsDestroyed(mass int, asteroids []int) bool {
	sort.Ints(asteroids)
	current := int64(mass)

	for _, a := range asteroids {
		if current < int64(a) {
			return false
		}
		current += int64(a)
	}

	return true
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", asteroidsDestroyed(10, []int{3, 9, 19, 5, 21}))
	// Expected: true

	// Test case 2
	fmt.Println("Test 2:", asteroidsDestroyed(5, []int{4, 9, 23, 4}))
	// Expected: false

	// Test case 3
	fmt.Println("Test 3:", asteroidsDestroyed(1, []int{1, 1, 1, 1}))
	// Expected: true
}
```

## 2128 — Remove All Ones With Row And Column Flips

```go
package main

// LeetCode #2128: Remove All Ones With Row and Column Flips
// https://leetcode.com/problems/remove-all-ones-with-row-and-column-flips/
// Difficulty: Medium [Paid]
// Time: O(m*n) | Space: O(1)

import "fmt"

func removeOnes(grid [][]int) bool {
	m, n := len(grid), len(grid[0])

	for i := 0; i < m; i++ {
		// Compare each row with first row: either equal or complementary
		same := true
		comp := true
		for j := 0; j < n; j++ {
			if grid[i][j] != grid[0][j] {
				same = false
			}
			if grid[i][j] != 1-grid[0][j] {
				comp = false
			}
		}
		if !same && !comp {
			return false
		}
	}
	return true
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", removeOnes([][]int{{0, 1, 0}, {1, 0, 1}, {0, 1, 0}}))
	// Expected: true

	// Test case 2
	fmt.Println("Test 2:", removeOnes([][]int{{1, 1, 0}, {0, 0, 0}, {0, 0, 0}}))
	// Expected: false

	// Test case 3
	fmt.Println("Test 3:", removeOnes([][]int{{0, 0, 0}, {0, 0, 0}}))
	// Expected: true
}
```

## 2130 — Maximum Twin Sum Of A Linked List

```go
package main

// LeetCode #2130: Maximum Twin Sum of a Linked List
// https://leetcode.com/problems/maximum-twin-sum-of-a-linked-list/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	// Find middle
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Reverse second half
	var prev *ListNode
	for slow != nil {
		next := slow.Next
		slow.Next = prev
		prev = slow
		slow = next
	}

	// Sum pairs
	maxSum := 0
	first, second := head, prev
	for second != nil {
		sum := first.Val + second.Val
		if sum > maxSum {
			maxSum = sum
		}
		first = first.Next
		second = second.Next
	}

	return maxSum
}

func main() {
	// Test case 1
	head1 := &ListNode{5, &ListNode{4, &ListNode{2, &ListNode{1, nil}}}}
	fmt.Println("Test 1:", pairSum(head1))
	// Expected: 6

	// Test case 2
	head2 := &ListNode{4, &ListNode{2, &ListNode{2, &ListNode{3, nil}}}}
	fmt.Println("Test 2:", pairSum(head2))
	// Expected: 7

	// Test case 3
	head3 := &ListNode{1, &ListNode{100000, nil}}
	fmt.Println("Test 3:", pairSum(head3))
	// Expected: 100001
}
```

## 2131 — Longest Palindrome By Concatenating Two Letter Words

```go
package main

// LeetCode #2131: Longest Palindrome by Concatenating Two-Letter Words
// https://leetcode.com/problems/longest-palindrome-by-concatenating-two-letter-words/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func longestPalindrome(words []string) int {
	freq := make(map[string]int)
	for _, w := range words {
		freq[w]++
	}

	length := 0
	centerUsed := false

	for w, count := range freq {
		if count == 0 {
			continue
		}
		rev := string([]byte{w[1], w[0]})

		if w == rev {
			// Same letter pair like "aa"
			pairs := count / 2
			length += pairs * 4
			if count%2 == 1 && !centerUsed {
				length += 2
				centerUsed = true
			}
			freq[w] = 0
		} else if revCount, ok := freq[rev]; ok && revCount > 0 {
			pairs := count
			if revCount < pairs {
				pairs = revCount
			}
			length += pairs * 4
			freq[w] -= pairs
			freq[rev] -= pairs
		}
	}

	return length
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", longestPalindrome([]string{"lc", "cl", "gg"}))
	// Expected: 6

	// Test case 2
	fmt.Println("Test 2:", longestPalindrome([]string{"ab", "ty", "yt", "lc", "cl", "ab"}))
	// Expected: 8

	// Test case 3
	fmt.Println("Test 3:", longestPalindrome([]string{"cc", "ll", "xx"}))
	// Expected: 2
}
```

## 2134 — Minimum Swaps To Group All 1s Together Ii

```go
package main

// LeetCode #2134: Minimum Swaps to Group All 1's Together II
// https://leetcode.com/problems/minimum-swaps-to-group-all-1s-together-ii/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func minSwaps(nums []int) int {
	n := len(nums)
	totalOnes := 0
	for _, v := range nums {
		if v == 1 {
			totalOnes++
		}
	}
	if totalOnes <= 1 {
		return 0
	}

	// Count zeros in first window of size totalOnes
	zeros := 0
	for i := 0; i < totalOnes; i++ {
		if nums[i] == 0 {
			zeros++
		}
	}
	minZeros := zeros

	// Slide window
	for i := totalOnes; i < n+totalOnes; i++ {
		// Remove element leaving window
		if nums[(i-totalOnes)%n] == 0 {
			zeros--
		}
		// Add element entering window
		if nums[i%n] == 0 {
			zeros++
		}
		if zeros < minZeros {
			minZeros = zeros
		}
	}

	return minZeros
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minSwaps([]int{0, 1, 0, 1, 1, 0, 0}))
	// Expected: 1

	// Test case 2
	fmt.Println("Test 2:", minSwaps([]int{0, 1, 1, 1, 0, 0, 1, 1, 0}))
	// Expected: 2

	// Test case 3
	fmt.Println("Test 3:", minSwaps([]int{1, 1, 0, 0, 1}))
	// Expected: 0
}
```

## 2135 — Count Words Obtained After Adding A Letter

```go
package main

// LeetCode #2135: Count Words Obtained After Adding a Letter
// https://leetcode.com/problems/count-words-obtained-after-adding-a-letter/
// Difficulty: Medium
// Time: O(n * L) | Space: O(n)

import "fmt"

func wordCount(startWords []string, targetWords []string) int {
	// Convert start words to bitmasks
	startSet := make(map[int]bool)
	for _, w := range startWords {
		mask := 0
		for _, c := range w {
			mask |= 1 << (c - 'a')
		}
		startSet[mask] = true
	}

	count := 0
	for _, w := range targetWords {
		mask := 0
		for _, c := range w {
			mask |= 1 << (c - 'a')
		}
		// Try removing each character
		for _, c := range w {
			bit := 1 << (c - 'a')
			if startSet[mask^bit] {
				count++
				break
			}
		}
	}

	return count
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", wordCount([]string{"ant", "act", "tack"}, []string{"tack", "act", "acti"}))
	// Expected: 2

	// Test case 2
	fmt.Println("Test 2:", wordCount([]string{"ab", "a"}, []string{"abc", "abcd"}))
	// Expected: 1
}
```

## 2137 — Pour Water Between Buckets To Make Water Levels Equal

```go
package main

// LeetCode #2137: Pour Water Between Buckets to Make Water Levels Equal
// https://leetcode.com/problems/pour-water-between-buckets-to-make-water-levels-equal/
// Difficulty: Medium [Paid]
// Time: O(n log precision) | Space: O(1)

import "fmt"

func equalizeWater(buckets []int, loss int) float64 {
	canReach := func(target float64) bool {
		need := 0.0
		give := 0.0
		for _, b := range buckets {
			amount := float64(b)
			if amount < target {
				need += target - amount
			} else {
				give += (amount - target) * (1.0 - float64(loss)/100.0)
			}
		}
		return give >= need
	}

	low, high := 0.0, 0.0
	for _, b := range buckets {
		if float64(b) > high {
			high = float64(b)
		}
	}

	for i := 0; i < 100; i++ {
		mid := (low + high) / 2
		if canReach(mid) {
			low = mid
		} else {
			high = mid
		}
	}

	return low
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", equalizeWater([]int{1, 2, 7}, 20))
	// Expected: ~2.000

	// Test case 2
	fmt.Println("Test 2:", equalizeWater([]int{2, 4, 6}, 50))
	// Expected: ~3.500

	// Test case 3
	fmt.Println("Test 3:", equalizeWater([]int{3, 3, 3}, 0))
	// Expected: 3.000
}
```

## 2139 — Minimum Moves To Reach Target Score

```go
package main

// LeetCode #2139: Minimum Moves to Reach Target Score
// https://leetcode.com/problems/minimum-moves-to-reach-target-score/
// Difficulty: Medium
// Time: O(log target) | Space: O(1)

import "fmt"

func minMoves(target int, maxDoubles int) int {
	moves := 0
	for target > 1 {
		if maxDoubles == 0 {
			moves += target - 1
			break
		}
		if target%2 == 1 {
			target--
			moves++
		} else {
			target /= 2
			maxDoubles--
			moves++
		}
	}
	return moves
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minMoves(5, 0))
	// Expected: 4

	// Test case 2
	fmt.Println("Test 2:", minMoves(19, 2))
	// Expected: 7

	// Test case 3
	fmt.Println("Test 3:", minMoves(10, 4))
	// Expected: 4
}
```

## 2140 — Solving Questions With Brainpower

```go
package main

// LeetCode #2140: Solving Questions With Brainpower
// https://leetcode.com/problems/solving-questions-with-brainpower/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func mostPoints(questions [][]int) int64 {
	n := len(questions)
	dp := make([]int64, n+1)

	for i := n - 1; i >= 0; i-- {
		points := int64(questions[i][0])
		brainpower := questions[i][1]
		// Skip this question
		best := dp[i+1]
		// Take this question
		next := i + brainpower + 1
		take := points
		if next <= n {
			take += dp[next]
		}
		if take > best {
			best = take
		}
		dp[i] = best
	}

	return dp[0]
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", mostPoints([][]int{{3, 2}, {4, 3}, {4, 4}, {2, 5}}))
	// Expected: 5

	// Test case 2
	fmt.Println("Test 2:", mostPoints([][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}}))
	// Expected: 7

	// Test case 3
	fmt.Println("Test 3:", mostPoints([][]int{{21, 5}, {92, 3}, {74, 2}, {39, 4}, {58, 2}, {5, 5}, {49, 4}, {65, 3}}))
	// Expected: 157
}
```

## 2142 — The Number Of Passengers In Each Bus I

```go
package main

// LeetCode #2142: The Number of Passengers in Each Bus I
// https://leetcode.com/problems/the-number-of-passengers-in-each-bus-i/
// Difficulty: Medium [Paid]
// Time: O(n log n + m log m) | Space: O(n + m)

import (
	"fmt"
	"sort"
)

type Bus struct {
	ID      int
	Arrival int
}

func busPassengers(buses [][]int, passengers []int) []int {
	// Sort buses by arrival time
	sort.Slice(buses, func(i, j int) bool {
		return buses[i][1] < buses[j][1]
	})

	// Sort passengers
	sort.Ints(passengers)

	result := make([]int, len(buses))
	pIdx := 0

	for i, bus := range buses {
		capacity := bus[0]
		count := 0
		for pIdx < len(passengers) && count < capacity && passengers[pIdx] <= bus[1] {
			count++
			pIdx++
		}
		result[i] = count
	}

	return result
}

func main() {
	// Test case 1: buses[capacity, arrival_time]
	buses1 := [][]int{{2, 10}, {3, 20}}
	passengers1 := []int{3, 8, 15, 18, 25}
	fmt.Println("Test 1:", busPassengers(buses1, passengers1))
	// Expected: [1, 2]

	// Test case 2
	buses2 := [][]int{{1, 5}, {2, 10}}
	passengers2 := []int{3, 8, 12}
	fmt.Println("Test 2:", busPassengers(buses2, passengers2))
	// Expected: [1, 1]

	// Test case 3
	buses3 := [][]int{{5, 100}}
	passengers3 := []int{}
	fmt.Println("Test 3:", busPassengers(buses3, passengers3))
	// Expected: [0]
}
```

## 2145 — Count The Hidden Sequences

```go
package main

// LeetCode #2145: Count the Hidden Sequences
// https://leetcode.com/problems/count-the-hidden-sequences/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func numberOfArrays(differences []int, lower int, upper int) int {
	cur := int64(0)
	minVal := int64(0)
	maxVal := int64(0)

	for _, d := range differences {
		cur += int64(d)
		if cur < minVal {
			minVal = cur
		}
		if cur > maxVal {
			maxVal = cur
		}
	}

	// Our sequence starts at some value x in [lower, upper]
	// All elements: x + prefix[i] must be in [lower, upper]
	// So: lower <= x + minVal AND x + maxVal <= upper
	// => x >= lower - minVal AND x <= upper - maxVal
	// => valid x in [max(lower, lower-minVal), min(upper, upper-maxVal)]

	low := int64(lower)
	high := int64(upper)
	minStart := low - minVal
	maxStart := high - maxVal

	if minStart > maxStart {
		return 0
	}

	start := minStart
	if start < low {
		start = low
	}
	end := maxStart
	if end > high {
		end = high
	}

	if start > end {
		return 0
	}
	return int(end - start + 1)
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", numberOfArrays([]int{1, -3, 4}, 1, 6))
	// Expected: 2

	// Test case 2
	fmt.Println("Test 2:", numberOfArrays([]int{3, -4, 5, 1, -2}, -4, 5))
	// Expected: 4

	// Test case 3
	fmt.Println("Test 3:", numberOfArrays([]int{4, -7, 2}, 3, 6))
	// Expected: 0
}
```

## 2146 — K Highest Ranked Items Within A Price Range

```go
package main

// LeetCode #2146: K Highest Ranked Items Within a Price Range
// https://leetcode.com/problems/k-highest-ranked-items-within-a-price-range/
// Difficulty: Medium
// Time: O(m*n log(m*n)) | Space: O(m*n)

import (
	"fmt"
	"sort"
)

func highestRankedKItems(grid [][]int, pricing []int, start []int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	low, high := pricing[0], pricing[1]

	// BFS
	type Item struct {
		dist, price, row, col int
	}
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	queue := [][2]int{{start[0], start[1]}}
	visited[start[0]][start[1]] = true
	items := []Item{}
	dist := 0

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			r, c := queue[i][0], queue[i][1]
			if grid[r][c] >= low && grid[r][c] <= high {
				items = append(items, Item{dist, grid[r][c], r, c})
			}
			for _, d := range dirs {
				nr, nc := r+d[0], c+d[1]
				if nr >= 0 && nr < m && nc >= 0 && nc < n && !visited[nr][nc] && grid[nr][nc] != 0 {
					visited[nr][nc] = true
					queue = append(queue, [2]int{nr, nc})
				}
			}
		}
		queue = queue[size:]
		dist++
	}

	// Sort by rank: distance, price, row, col
	sort.Slice(items, func(i, j int) bool {
		if items[i].dist != items[j].dist {
			return items[i].dist < items[j].dist
		}
		if items[i].price != items[j].price {
			return items[i].price < items[j].price
		}
		if items[i].row != items[j].row {
			return items[i].row < items[j].row
		}
		return items[i].col < items[j].col
	})

	// Take first k
	result := make([][]int, 0, k)
	for i := 0; i < len(items) && i < k; i++ {
		result = append(result, []int{items[i].row, items[i].col})
	}
	return result
}

func main() {
	// Test case 1
	grid1 := [][]int{{1, 2, 0, 1}, {1, 3, 0, 1}, {0, 2, 5, 1}}
	fmt.Println("Test 1:", highestRankedKItems(grid1, []int{2, 5}, []int{0, 0}, 3))
	// Expected: [[0,1],[1,1],[2,1]]

	// Test case 2
	grid2 := [][]int{{1, 1, 1}, {0, 0, 1}, {2, 3, 4}}
	fmt.Println("Test 2:", highestRankedKItems(grid2, []int{2, 3}, []int{0, 0}, 3))
	// Expected: [[2,1],[2,0]]

	// Test case 3
	grid3 := [][]int{{1, 2}, {3, 4}}
	fmt.Println("Test 3:", highestRankedKItems(grid3, []int{1, 4}, []int{0, 0}, 10))
	// Expected: [[0,0],[0,1],[1,0],[1,1]]
}
```

## 2149 — Rearrange Array Elements By Sign

```go
package main

// LeetCode #2149: Rearrange Array Elements by Sign
// https://leetcode.com/problems/rearrange-array-elements-by-sign/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func rearrangeArray(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	posIdx, negIdx := 0, 1

	for _, v := range nums {
		if v > 0 {
			result[posIdx] = v
			posIdx += 2
		} else {
			result[negIdx] = v
			negIdx += 2
		}
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", rearrangeArray([]int{3, 1, -2, -5, 2, -4}))
	// Expected: [3,-2,1,-5,2,-4]

	// Test case 2
	fmt.Println("Test 2:", rearrangeArray([]int{-1, 1}))
	// Expected: [1,-1]

	// Test case 3
	fmt.Println("Test 3:", rearrangeArray([]int{28, -41, 22, -8, -37, 46, 35, -9}))
	// Expected: [28,-41,22,-8,46,-37,35,-9]
}
```

## 2150 — Find All Lonely Numbers In The Array

```go
package main

// LeetCode #2150: Find All Lonely Numbers in the Array
// https://leetcode.com/problems/find-all-lonely-numbers-in-the-array/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func findLonely(nums []int) []int {
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}

	result := []int{}
	for _, v := range nums {
		if freq[v] == 1 && freq[v-1] == 0 && freq[v+1] == 0 {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", findLonely([]int{10, 6, 5, 8}))
	// Expected: [10, 8]

	// Test case 2
	fmt.Println("Test 2:", findLonely([]int{1, 3, 5, 3}))
	// Expected: [1, 5]
}
```

## 2152 — Minimum Number Of Lines To Cover Points

```go
package main

// LeetCode #2152: Minimum Number of Lines to Cover Points
// https://leetcode.com/problems/minimum-number-of-lines-to-cover-points/
// Difficulty: Medium [Paid]
// Time: O(n^2 * 2^n) | Space: O(2^n)

import "fmt"

func minimumLines(points [][]int) int {
	n := len(points)
	if n <= 1 {
		return n
	}

	// Precompute line masks
	lineMasks := []int{}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			mask := 0
			for k := 0; k < n; k++ {
				if isCollinear(points[i], points[j], points[k]) {
					mask |= 1 << k
				}
			}
			lineMasks = append(lineMasks, mask)
		}
	}

	// DP over subsets
	dp := make([]int, 1<<n)
	for i := range dp {
		dp[i] = n // max n lines
	}
	dp[0] = 0

	for mask := 0; mask < (1 << n); mask++ {
		for _, lm := range lineMasks {
			dp[mask|lm] = min(dp[mask|lm], dp[mask]+1)
		}
	}

	return dp[(1<<n)-1]
}

func isCollinear(a, b, c []int) bool {
	return (b[1]-a[1])*(c[0]-a[0]) == (c[1]-a[1])*(b[0]-a[0])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minimumLines([][]int{{0, 1}, {2, 3}, {4, 5}, {4, 3}}))
	// Expected: 2

	// Test case 2
	fmt.Println("Test 2:", minimumLines([][]int{{0, 2}, {-2, -2}, {1, 4}}))
	// Expected: 1

	// Test case 3
	fmt.Println("Test 3:", minimumLines([][]int{{0, 0}, {1, 1}, {2, 2}, {3, 3}}))
	// Expected: 1
}
```

## 2155 — All Divisions With The Highest Score Of A Binary Array

```go
package main

// LeetCode #2155: All Divisions With the Highest Score of a Binary Array
// https://leetcode.com/problems/all-divisions-with-the-highest-score-of-a-binary-array/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func maxScoreIndices(nums []int) []int {
	n := len(nums)
	totalOnes := 0
	for _, v := range nums {
		totalOnes += v
	}

	maxScore := totalOnes // score at index 0 (0 zeros left + totalOnes ones right)
	result := []int{0}
	zerosLeft := 0

	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			zerosLeft++
		}
		onesRight := totalOnes - (i + 1 - zerosLeft) // total - ones so far
		score := zerosLeft + onesRight
		if score > maxScore {
			maxScore = score
			result = []int{i + 1}
		} else if score == maxScore {
			result = append(result, i+1)
		}
	}

	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", maxScoreIndices([]int{0, 0, 1, 0}))
	// Expected: [2, 4]

	// Test case 2
	fmt.Println("Test 2:", maxScoreIndices([]int{0, 0, 0}))
	// Expected: [3]

	// Test case 3
	fmt.Println("Test 3:", maxScoreIndices([]int{1, 1}))
	// Expected: [0]
}
```

## 2159 — Order Two Columns Independently

```go
package main

// LeetCode #2159: Order Two Columns Independently
// https://leetcode.com/problems/order-two-columns-independently/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

type Row struct {
	col1, col2 int
}

func orderColumns(rows [][]int) [][]int {
	// Sort by col1 ascending, col2 ascending
	sort.Slice(rows, func(i, j int) bool {
		if rows[i][0] != rows[j][0] {
			return rows[i][0] < rows[j][0]
		}
		return rows[i][1] < rows[j][1]
	})

	return rows
}

func main() {
	// Test case 1
	data1 := [][]int{{3, 1}, {1, 3}, {2, 2}}
	fmt.Println("Test 1:", orderColumns(data1))
	// Expected: [[1,3],[2,2],[3,1]]

	// Test case 2
	data2 := [][]int{{5, 5}, {1, 1}, {3, 3}}
	fmt.Println("Test 2:", orderColumns(data2))
	// Expected: [[1,1],[3,3],[5,5]]
}
```

## 2161 — Partition Array According To Given Pivot

```go
package main

// LeetCode #2161: Partition Array According to Given Pivot
// https://leetcode.com/problems/partition-array-according-to-given-pivot/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func pivotArray(nums []int, pivot int) []int {
	less, equal, greater := []int{}, []int{}, []int{}
	for _, v := range nums {
		if v < pivot {
			less = append(less, v)
		} else if v == pivot {
			equal = append(equal, v)
		} else {
			greater = append(greater, v)
		}
	}
	return append(append(less, equal...), greater...)
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", pivotArray([]int{9, 12, 5, 10, 14, 3, 10}, 10))
	// Expected: [9,5,3,10,10,12,14]

	// Test case 2
	fmt.Println("Test 2:", pivotArray([]int{-3, 4, 3, 2}, 2))
	// Expected: [-3,2,4,3]

	// Test case 3
	fmt.Println("Test 3:", pivotArray([]int{1, 2, 3, 4, 5}, 3))
	// Expected: [1,2,3,4,5]
}
```

## 2162 — Minimum Cost To Set Cooking Time

```go
package main

// LeetCode #2162: Minimum Cost to Set Cooking Time
// https://leetcode.com/problems/minimum-cost-to-set-cooking-time/
// Difficulty: Medium
// Time: O(1) | Space: O(1)

import (
	"fmt"
)

func minCostSetTime(startAt int, moveCost int, pushCost int, targetSeconds int) int {
	// Try all valid representations: minutes:seconds
	cost := int(1e9)

	// mm:ss where 0 <= mm <= 99, 0 <= ss <= 99
	for m := 0; m <= 99; m++ {
		for s := 0; s <= 99; s++ {
			if m*60+s == targetSeconds {
				digits := []int{}
				if m >= 10 {
					digits = append(digits, m/10)
				}
				digits = append(digits, m%10)
				if s < 10 && len(digits) > 0 {
					// If we have minutes, seconds always need 2 digits
					digits = append(digits, s/10)
				} else if s >= 10 {
					digits = append(digits, s/10)
				}
				digits = append(digits, s%10)

				cur := startAt
				total := 0
				for _, d := range digits {
					if cur != d {
						total += moveCost
						cur = d
					}
					total += pushCost
				}
				if total < cost {
					cost = total
				}
			}
		}
	}

	return cost
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minCostSetTime(1, 2, 1, 600))
	// Expected: 6

	// Test case 2
	fmt.Println("Test 2:", minCostSetTime(0, 1, 2, 76))
	// Expected: 6

	// Test case 3
	fmt.Println("Test 3:", minCostSetTime(9, 100, 1, 600))
	// Expected: 8
}
```

## 2165 — Smallest Value Of The Rearranged Number

```go
package main

// LeetCode #2165: Smallest Value of the Rearranged Number
// https://leetcode.com/problems/smallest-value-of-the-rearranged-number/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func smallestNumber(num int64) int64 {
	if num == 0 {
		return 0
	}

	negative := num < 0
	if negative {
		num = -num
	}

	// Extract digits
	digits := []int{}
	for num > 0 {
		digits = append(digits, int(num%10))
		num /= 10
	}

	sort.Ints(digits)

	if negative {
		// Largest number from digits (descending)
		result := int64(0)
		for i := len(digits) - 1; i >= 0; i-- {
			result = result*10 + int64(digits[i])
		}
		return -result
	}

	// Smallest number: place smallest non-zero digit first
	// Find first non-zero digit
	i := 0
	for i < len(digits) && digits[i] == 0 {
		i++
	}
	// Swap first non-zero with position 0
	if i < len(digits) {
		digits[0], digits[i] = digits[i], digits[0]
	}

	result := int64(0)
	for _, d := range digits {
		result = result*10 + int64(d)
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", smallestNumber(310))
	// Expected: 103

	// Test case 2
	fmt.Println("Test 2:", smallestNumber(-7605))
	// Expected: -7650

	// Test case 3
	fmt.Println("Test 3:", smallestNumber(0))
	// Expected: 0
}
```

## 2166 — Design Bitset

```go
package main

// LeetCode #2166: Design Bitset
// https://leetcode.com/problems/design-bitset/
// Difficulty: Medium
// Time: O(1) per operation | Space: O(n)

import (
	"fmt"
	"strings"
)

type Bitset struct {
	bits  []bool
	flip  bool
	ones  int
	size  int
}

func Constructor(size int) Bitset {
	return Bitset{
		bits: make([]bool, size),
		size: size,
	}
}

func (this *Bitset) Fix(idx int) {
	if this.flip {
		if this.bits[idx] {
			this.bits[idx] = false
			this.ones++
		}
	} else {
		if !this.bits[idx] {
			this.bits[idx] = true
			this.ones++
		}
	}
}

func (this *Bitset) Unfix(idx int) {
	if this.flip {
		if !this.bits[idx] {
			this.bits[idx] = true
			this.ones--
		}
	} else {
		if this.bits[idx] {
			this.bits[idx] = false
			this.ones--
		}
	}
}

func (this *Bitset) Flip() {
	this.flip = !this.flip
	this.ones = this.size - this.ones
}

func (this *Bitset) All() bool {
	return this.ones == this.size
}

func (this *Bitset) One() bool {
	return this.ones > 0
}

func (this *Bitset) Count() int {
	return this.ones
}

func (this *Bitset) ToString() string {
	var sb strings.Builder
	for i := 0; i < this.size; i++ {
		bit := this.bits[i]
		if this.flip {
			bit = !bit
		}
		if bit {
			sb.WriteByte('1')
		} else {
			sb.WriteByte('0')
		}
	}
	return sb.String()
}

func main() {
	// Test case 1
	bs := Constructor(5)
	bs.Fix(3)
	bs.Fix(1)
	fmt.Println("Test 1 All:", bs.All())
	fmt.Println("Test 1 toString:", bs.ToString())
	// Expected: false, "01010"

	// Test case 2
	bs2 := Constructor(5)
	bs2.Fix(0)
	bs2.Fix(1)
	bs2.Flip()
	fmt.Println("Test 2 All:", bs2.All())
	fmt.Println("Test 2 toString:", bs2.ToString())
	// Expected: false, "00111"

	// Test case 3
	bs3 := Constructor(2)
	bs3.Flip()
	fmt.Println("Test 3 All:", bs3.All())
	fmt.Println("Test 3 toString:", bs3.ToString())
	// Expected: true, "11"
}
```

## 2168 — Unique Substrings With Equal Digit Frequency

```go
package main

// LeetCode #2168: Unique Substrings With Equal Digit Frequency
// https://leetcode.com/problems/unique-substrings-with-equal-digit-frequency/
// Difficulty: Medium [Paid]
// Time: O(n^2) | Space: O(n^2)

import "fmt"

func equalDigitFrequency(s string) int {
	n := len(s)
	seen := make(map[string]bool)

	for i := 0; i < n; i++ {
		freq := make([]int, 10)
		distinct := 0
		maxFreq := 0
		for j := i; j < n; j++ {
			d := int(s[j] - '0')
			if freq[d] == 0 {
				distinct++
			}
			freq[d]++
			if freq[d] > maxFreq {
				maxFreq = freq[d]
			}
			// All digits appear same frequency iff distinct * maxFreq == totalLen
			if distinct*maxFreq == j-i+1 {
				seen[s[i:j+1]] = true
			}
		}
	}

	return len(seen)
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", equalDigitFrequency("1212"))
	// Expected: 5

	// Test case 2
	fmt.Println("Test 2:", equalDigitFrequency("12321"))
	// Expected: 9
}
```

## 2170 — Minimum Operations To Make The Array Alternating

```go
package main

// LeetCode #2170: Minimum Operations to Make the Array Alternating
// https://leetcode.com/problems/minimum-operations-to-make-the-array-alternating/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func minimumOperations(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	// Frequency maps for even and odd positions
	evenFreq := make(map[int]int)
	oddFreq := make(map[int]int)

	for i, v := range nums {
		if i%2 == 0 {
			evenFreq[v]++
		} else {
			oddFreq[v]++
		}
	}

	// Get top 2 most frequent values for even positions
	evenTop1, evenTop2 := getTopTwo(evenFreq)
	oddTop1, oddTop2 := getTopTwo(oddFreq)

	evenCount := (n + 1) / 2 // number of even positions
	oddCount := n / 2        // number of odd positions

	if evenTop1.val != oddTop1.val {
		return (evenCount - evenTop1.count) + (oddCount - oddTop1.count)
	}

	// Try both combinations
	candidate1 := (evenCount - evenTop1.count) + (oddCount - oddTop2.count)
	candidate2 := (evenCount - evenTop2.count) + (oddCount - oddTop1.count)
	if candidate1 < candidate2 {
		return candidate1
	}
	return candidate2
}

type freqPair struct {
	val   int
	count int
}

func getTopTwo(freq map[int]int) (freqPair, freqPair) {
	top1 := freqPair{val: -1, count: 0}
	top2 := freqPair{val: -1, count: 0}

	for val, cnt := range freq {
		if cnt > top1.count {
			top2 = top1
			top1 = freqPair{val, cnt}
		} else if cnt > top2.count {
			top2 = freqPair{val, cnt}
		}
	}

	return top1, top2
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minimumOperations([]int{3, 1, 3, 2, 4, 3}))
	// Expected: 3

	// Test case 2
	fmt.Println("Test 2:", minimumOperations([]int{1, 2, 2, 2, 2}))
	// Expected: 2

	// Test case 3
	fmt.Println("Test 3:", minimumOperations([]int{1, 1, 1, 1}))
	// Expected: 2
}
```

## 2171 — Removing Minimum Number Of Magic Beans

```go
package main

// LeetCode #2171: Removing Minimum Number of Magic Beans
// https://leetcode.com/problems/removing-minimum-number-of-magic-beans/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func minimumRemoval(beans []int) int64 {
	sort.Ints(beans)
	n := len(beans)
	total := int64(0)
	for _, b := range beans {
		total += int64(b)
	}

	minRemoved := total // removing all beans is worst case
	for i, b := range beans {
		// If we make all remaining bags have 'b' beans:
		// we keep (n-i) * b beans
		removed := total - int64(n-i)*int64(b)
		if removed < minRemoved {
			minRemoved = removed
		}
	}

	return minRemoved
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minimumRemoval([]int{4, 1, 6, 5}))
	// Expected: 4

	// Test case 2
	fmt.Println("Test 2:", minimumRemoval([]int{2, 10, 3, 2}))
	// Expected: 7
}
```

## 2174 — Remove All Ones With Row And Column Flips Ii

```go
package main

// LeetCode #2174: Remove All Ones With Row and Column Flips II
// https://leetcode.com/problems/remove-all-ones-with-row-and-column-flips-ii/
// Difficulty: Medium [Paid]
// Time: O(2^(m*n) * m * n) | Space: O(2^(m*n))

import "fmt"

func removeOnes(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	total := m * n

	// Compress grid to bitmask
	start := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				start |= 1 << (i*n + j)
			}
		}
	}

	if start == 0 {
		return 0
	}

	dist := make([]int, 1<<total)
	for i := range dist {
		dist[i] = -1
	}
	dist[start] = 0
	queue := []int{start}

	for len(queue) > 0 {
		mask := queue[0]
		queue = queue[1:]

		// Only try cells that were 1 in the ORIGINAL grid
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if grid[i][j] == 0 {
					continue
				}
				next := mask
				// Clear entire row i
				for c := 0; c < n; c++ {
					next &^= 1 << (i*n + c)
				}
				// Clear entire column j
				for r := 0; r < m; r++ {
					next &^= 1 << (r*n + j)
				}

				if dist[next] == -1 {
					dist[next] = dist[mask] + 1
					if next == 0 {
						return dist[next]
					}
					queue = append(queue, next)
				}
			}
		}
	}

	return -1
}

func main() {
	// Test case 1 (Example 2 from problem)
	fmt.Println("Test 1:", removeOnes([][]int{{0, 1, 0}, {1, 0, 1}, {0, 1, 0}}))
	// Expected: 2

	// Test case 2 (Example 1 from problem)
	fmt.Println("Test 2:", removeOnes([][]int{{1, 1, 1}, {1, 1, 1}, {0, 1, 0}}))
	// Expected: 2

	// Test case 3 (all zeros)
	fmt.Println("Test 3:", removeOnes([][]int{{0, 0}, {0, 0}}))
	// Expected: 0
}
```

## 2175 — The Change In Global Rankings

```go
package main

// LeetCode #2175: The Change in Global Rankings
// https://leetcode.com/problems/the-change-in-global-rankings/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func globalRankings(pointsBefore []int, pointsAfter []int) []int {
	n := len(pointsBefore)

	// Create sorted list of (points, originalIndex) for before
	type pair struct {
		points int
		idx    int
	}
	before := make([]pair, n)
	for i := 0; i < n; i++ {
		before[i] = pair{pointsBefore[i], i}
	}
	sort.Slice(before, func(i, j int) bool {
		if before[i].points != before[j].points {
			return before[i].points > before[j].points
		}
		return before[i].idx < before[j].idx
	})

	// Compute rank before: rank = position (1-indexed) when sorted descending
	rankBefore := make([]int, n)
	for pos, p := range before {
		rankBefore[p.idx] = pos + 1
	}

	// Create sorted list for after
	after := make([]pair, n)
	for i := 0; i < n; i++ {
		after[i] = pair{pointsAfter[i], i}
	}
	sort.Slice(after, func(i, j int) bool {
		if after[i].points != after[j].points {
			return after[i].points > after[j].points
		}
		return after[i].idx < after[j].idx
	})

	// Compute rank after
	rankAfter := make([]int, n)
	for pos, p := range after {
		rankAfter[p.idx] = pos + 1
	}

	// Difference
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = rankBefore[i] - rankAfter[i]
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", globalRankings([]int{10, 20, 30}, []int{15, 25, 35}))
	// Expected: [0, 0, 0] (same relative order)

	// Test case 2
	fmt.Println("Test 2:", globalRankings([]int{50, 40, 30, 20}, []int{45, 45, 35, 25}))
	// Expected: varying based on rank changes
}
```

## 2177 — Find Three Consecutive Integers That Sum To A Given Number

```go
package main

// LeetCode #2177: Find Three Consecutive Integers That Sum to a Given Number
// https://leetcode.com/problems/find-three-consecutive-integers-that-sum-to-a-given-number/
// Difficulty: Medium
// Time: O(1) | Space: O(1)

import "fmt"

func sumOfThree(num int64) []int64 {
	if num%3 != 0 {
		return []int64{}
	}
	x := num / 3
	return []int64{x - 1, x, x + 1}
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", sumOfThree(33))
	// Expected: [10, 11, 12]

	// Test case 2
	fmt.Println("Test 2:", sumOfThree(4))
	// Expected: []

	// Test case 3
	fmt.Println("Test 3:", sumOfThree(0))
	// Expected: [-1, 0, 1]
}
```

## 2178 — Maximum Split Of Positive Even Integers

```go
package main

// LeetCode #2178: Maximum Split of Positive Even Integers
// https://leetcode.com/problems/maximum-split-of-positive-even-integers/
// Difficulty: Medium
// Time: O(sqrt(finalSum)) | Space: O(sqrt(finalSum))

import "fmt"

func maximumEvenSplit(finalSum int64) []int64 {
	if finalSum%2 != 0 {
		return []int64{}
	}

	result := []int64{}
	cur := int64(2)

	for finalSum >= cur {
		result = append(result, cur)
		finalSum -= cur
		cur += 2
	}

	// Add remaining to last element
	if finalSum > 0 {
		result[len(result)-1] += finalSum
	}

	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", maximumEvenSplit(12))
	// Expected: [2, 4, 6] or [2, 10] — both valid, but greedy gives max length

	// Test case 2
	fmt.Println("Test 2:", maximumEvenSplit(7))
	// Expected: []

	// Test case 3
	fmt.Println("Test 3:", maximumEvenSplit(28))
	// Expected: [2, 4, 6, 16] (greedy fills smallest unique evens, remainder added to last)
}
```

## 2181 — Merge Nodes In Between Zeros

```go
package main

// LeetCode #2181: Merge Nodes in Between Zeros
// https://leetcode.com/problems/merge-nodes-in-between-zeros/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeNodes(head *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	sum := 0

	for head != nil {
		if head.Val == 0 && sum > 0 {
			curr.Next = &ListNode{Val: sum}
			curr = curr.Next
			sum = 0
		}
		sum += head.Val
		head = head.Next
	}

	return dummy.Next
}

func main() {
	// Test case 1: Example [0,3,1,0,4,5,2,0]
	head1 := &ListNode{0, &ListNode{3, &ListNode{1, &ListNode{0, &ListNode{4, &ListNode{5, &ListNode{2, &ListNode{0, nil}}}}}}}}
	for n := mergeNodes(head1); n != nil; n = n.Next {
		fmt.Printf("%d ", n.Val)
	}
	fmt.Println()
	// Expected: 4 11

	// Test case 2: [0,1,0,3,0,2,2,0]
	head2 := &ListNode{0, &ListNode{1, &ListNode{0, &ListNode{3, &ListNode{0, &ListNode{2, &ListNode{2, &ListNode{0, nil}}}}}}}}
	for n := mergeNodes(head2); n != nil; n = n.Next {
		fmt.Printf("%d ", n.Val)
	}
	fmt.Println()
	// Expected: 1 3 4
}
```

## 2182 — Construct String With Repeat Limit

```go
package main

// LeetCode #2182: Construct String With Repeat Limit
// https://leetcode.com/problems/construct-string-with-repeat-limit/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func repeatLimitedString(s string, repeatLimit int) string {
	count := make([]int, 26)
	for _, ch := range s {
		count[ch-'a']++
	}

	result := make([]byte, 0, len(s))
	for i := 25; i >= 0; {
		if count[i] == 0 {
			i--
			continue
		}

		use := min(count[i], repeatLimit)
		for k := 0; k < use; k++ {
			result = append(result, byte('a'+i))
		}
		count[i] -= use

		if count[i] > 0 {
			j := i - 1
			for j >= 0 && count[j] == 0 {
				j--
			}
			if j < 0 {
				break
			}
			result = append(result, byte('a'+j))
			count[j]--
		} else {
			i--
		}
	}

	return string(result)
}

func main() {
	// Test case 1
	fmt.Println(repeatLimitedString("cczazcc", 3))
	// Expected: "zzcccac"

	// Test case 2
	fmt.Println(repeatLimitedString("aababab", 2))
	// Expected: "bbabaa"

	// Test case 3
	fmt.Println(repeatLimitedString("robnsdvpuxbapuqgopqvxdrchivlifeepy", 2))
	// Expected: "yxxvvuvusrrqqppopponliihgfeeddcba"
}
```

## 2184 — Number Of Ways To Build Sturdy Brick Wall

```go
package main

// LeetCode #2184: Number of Ways to Build Sturdy Brick Wall
// https://leetcode.com/problems/number-of-ways-to-build-sturdy-brick-wall/
// Difficulty: Medium [Paid]
// Time: O(n * m) | Space: O(w) where w = max width

import "fmt"

func buildWall(height int, width int, bricks []int) int {
	const mod = 1_000_000_007

	// Generate all valid row patterns via DFS
	var rows [][]int
	var dfs func(curr []int, w int)
	dfs = func(curr []int, w int) {
		if w == width {
			row := make([]int, len(curr))
			copy(row, curr)
			rows = append(rows, row)
			return
		}
		for _, b := range bricks {
			if w+b <= width {
				dfs(append(curr, b), w+b)
			}
		}
	}
	dfs([]int{}, 0)

	// Precompute which rows are compatible (no shared seam)
	n := len(rows)
	compat := make([][]bool, n)
	for i := 0; i < n; i++ {
		compat[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			compat[i][j] = true
		}
	}

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			seami, seamj := 0, 0
			pi, pj := 0, 0
			for pi < len(rows[i])-1 && pj < len(rows[j])-1 {
				seami += rows[i][pi]
				seamj += rows[j][pj]
				if seami == seamj {
					compat[i][j] = false
					compat[j][i] = false
					break
				}
				if seami < seamj {
					pi++
				} else {
					pj++
				}
			}
		}
	}

	// DP: ways[h][r] = ways to build up to height h ending with row r
	dp := make([][]int, height)
	for h := 0; h < height; h++ {
		dp[h] = make([]int, n)
	}
	for r := 0; r < n; r++ {
		dp[0][r] = 1
	}

	for h := 1; h < height; h++ {
		for r := 0; r < n; r++ {
			for p := 0; p < n; p++ {
				if compat[p][r] {
					dp[h][r] = (dp[h][r] + dp[h-1][p]) % mod
				}
			}
		}
	}

	total := 0
	for r := 0; r < n; r++ {
		total = (total + dp[height-1][r]) % mod
	}
	return total
}

func main() {
	// Test case 1
	fmt.Println(buildWall(2, 3, []int{1, 2}))
	// Expected: 2

	// Test case 2
	fmt.Println(buildWall(1, 1, []int{1}))
	// Expected: 1
}
```

## 2186 — Minimum Number Of Steps To Make Two Strings Anagram Ii

```go
package main

// LeetCode #2186: Minimum Number of Steps to Make Two Strings Anagram II
// https://leetcode.com/problems/minimum-number-of-steps-to-make-two-strings-anagram-ii/
// Difficulty: Medium
// Time: O(n + m) | Space: O(1)

import "fmt"

func minSteps(s string, t string) int {
	count := make([]int, 26)
	for _, ch := range s {
		count[ch-'a']++
	}
	for _, ch := range t {
		count[ch-'a']--
	}

	steps := 0
	for _, c := range count {
		if c > 0 {
			steps += c
		} else {
			steps -= c
		}
	}
	return steps
}

func main() {
	// Test case 1
	fmt.Println(minSteps("leetcode", "coats"))
	// Expected: 7

	// Test case 2
	fmt.Println(minSteps("night", "thing"))
	// Expected: 0

	// Test case 3
	fmt.Println(minSteps("aba", "bab"))
	// Expected: 2
}
```

## 2187 — Minimum Time To Complete Trips

```go
package main

// LeetCode #2187: Minimum Time to Complete Trips
// https://leetcode.com/problems/minimum-time-to-complete-trips/
// Difficulty: Medium
// Time: O(n log m) | Space: O(1)

import "fmt"

func minimumTime(time []int, totalTrips int) int64 {
	lo, hi := int64(1), int64(1)
	for _, t := range time {
		if int64(t) > hi {
			hi = int64(t)
		}
	}
	hi *= int64(totalTrips)

	for lo < hi {
		mid := lo + (hi-lo)/2
		trips := int64(0)
		for _, t := range time {
			trips += mid / int64(t)
			if trips >= int64(totalTrips) {
				break
			}
		}
		if trips >= int64(totalTrips) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func main() {
	// Test case 1
	fmt.Println(minimumTime([]int{1, 2, 3}, 5))
	// Expected: 3

	// Test case 2
	fmt.Println(minimumTime([]int{2}, 1))
	// Expected: 2

	// Test case 3
	fmt.Println(minimumTime([]int{5, 10, 10}, 9))
	// Expected: 25
}
```

## 2189 — Number Of Ways To Build House Of Cards

```go
package main

// LeetCode #2189: Number of Ways to Build House of Cards
// https://leetcode.com/problems/number-of-ways-to-build-house-of-cards/
// Difficulty: Medium [Paid]
// Time: O(n^2) | Space: O(n)

import "fmt"

func houseOfCards(remainingCards int) int {
	// DP: dp[c] = number of ways to use c cards
	dp := make([]int, remainingCards+1)
	dp[0] = 1

	// Each row (level) uses 3k - 1 cards where k >= 2
	for k := 2; 3*k-1 <= remainingCards; k++ {
		cards := 3*k - 1
		for c := remainingCards; c >= cards; c-- {
			dp[c] += dp[c-cards]
		}
	}
	return dp[remainingCards]
}

func main() {
	// Test case 1
	fmt.Println(houseOfCards(16))
	// Expected: 2

	// Test case 2
	fmt.Println(houseOfCards(2))
	// Expected: 0

	// Test case 3
	fmt.Println(houseOfCards(6))
	// Expected: 1
}
```

## 2191 — Sort The Jumbled Numbers

```go
package main

// LeetCode #2191: Sort the Jumbled Numbers
// https://leetcode.com/problems/sort-the-jumbled-numbers/
// Difficulty: Medium
// Time: O(n log n + n * d) | Space: O(n)

import (
	"fmt"
	"sort"
)

func sortJumbled(mapping []int, nums []int) []int {
	type pair struct {
		val    int
		mapped int
		idx    int
	}

	pairs := make([]pair, len(nums))
	for i, num := range nums {
		mapped := 0
		if num == 0 {
			mapped = mapping[0]
		} else {
			digits := []int{}
			n := num
			for n > 0 {
				digits = append(digits, n%10)
				n /= 10
			}
			for j := len(digits) - 1; j >= 0; j-- {
				mapped = mapped*10 + mapping[digits[j]]
			}
		}
		pairs[i] = pair{num, mapped, i}
	}

	sort.SliceStable(pairs, func(i, j int) bool {
		if pairs[i].mapped != pairs[j].mapped {
			return pairs[i].mapped < pairs[j].mapped
		}
		return pairs[i].idx < pairs[j].idx
	})

	result := make([]int, len(nums))
	for i, p := range pairs {
		result[i] = p.val
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println(sortJumbled([]int{8, 9, 4, 0, 2, 1, 3, 5, 7, 6}, []int{991, 338, 38}))
	// Expected: [338, 38, 991]

	// Test case 2
	fmt.Println(sortJumbled([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{789, 456, 123}))
	// Expected: [123, 456, 789]
}
```

## 2192 — All Ancestors Of A Node In A Directed Acyclic Graph

```go
package main

// LeetCode #2192: All Ancestors of a Node in a Directed Acyclic Graph
// https://leetcode.com/problems/all-ancestors-of-a-node-in-a-directed-acyclic-graph/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n^2)

import (
	"fmt"
	"sort"
)

func getAncestors(n int, edges [][]int) [][]int {
	graph := make([][]int, n)
	for _, e := range edges {
		graph[e[1]] = append(graph[e[1]], e[0])
	}

	result := make([][]int, n)
	for i := 0; i < n; i++ {
		visited := make([]bool, n)
		queue := []int{i}
		visited[i] = true
		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]
			for _, parent := range graph[node] {
				if !visited[parent] {
					visited[parent] = true
					queue = append(queue, parent)
				}
			}
		}
		ancestors := []int{}
		for j := 0; j < n; j++ {
			if j != i && visited[j] {
				ancestors = append(ancestors, j)
			}
		}
		sort.Ints(ancestors)
		result[i] = ancestors
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println(getAncestors(8, [][]int{{0, 3}, {0, 4}, {1, 3}, {2, 4}, {2, 7}, {3, 5}, {3, 6}, {3, 7}, {4, 6}}))

	// Test case 2
	fmt.Println(getAncestors(5, [][]int{{0, 1}, {0, 2}, {0, 3}, {0, 4}, {1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}))
}
```

## 2195 — Append K Integers With Minimal Sum

```go
package main

// LeetCode #2195: Append K Integers With Minimal Sum
// https://leetcode.com/problems/append-k-integers-with-minimal-sum/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func minimalKSum(nums []int, k int) int64 {
	sort.Ints(nums)
	var sum int64 = 0
	prev := 0

	for _, num := range nums {
		if num == prev {
			continue
		}
		if num > prev+1 {
			gap := num - prev - 1
			if gap >= k {
				// arithmetic series: (prev+1) + (prev+2) + ... + (prev+k)
				first := int64(prev) + 1
				last := int64(prev) + int64(k)
				sum += (first + last) * int64(k) / 2
				k = 0
				break
			}
			first := int64(prev) + 1
			last := int64(num) - 1
			sum += (first + last) * int64(gap) / 2
			k -= gap
		}
		prev = num
		if k == 0 {
			break
		}
	}

	if k > 0 {
		first := int64(prev) + 1
		last := int64(prev) + int64(k)
		sum += (first + last) * int64(k) / 2
	}

	return sum
}

func main() {
	// Test case 1
	fmt.Println(minimalKSum([]int{1, 4, 25, 10, 25}, 2))
	// Expected: 5

	// Test case 2
	fmt.Println(minimalKSum([]int{5, 6}, 6))
	// Expected: 25
}
```

## 2196 — Create Binary Tree From Descriptions

```go
package main

// LeetCode #2196: Create Binary Tree From Descriptions
// https://leetcode.com/problems/create-binary-tree-from-descriptions/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createBinaryTree(descriptions [][]int) *TreeNode {
	children := make(map[int]bool)
	nodes := make(map[int]*TreeNode)

	for _, desc := range descriptions {
		parent, child, isLeft := desc[0], desc[1], desc[2]
		if nodes[parent] == nil {
			nodes[parent] = &TreeNode{Val: parent}
		}
		if nodes[child] == nil {
			nodes[child] = &TreeNode{Val: child}
		}
		if isLeft == 1 {
			nodes[parent].Left = nodes[child]
		} else {
			nodes[parent].Right = nodes[child]
		}
		children[child] = true
	}

	for _, desc := range descriptions {
		if !children[desc[0]] {
			return nodes[desc[0]]
		}
	}
	return nil
}

func printTree(root *TreeNode) {
	if root == nil {
		fmt.Print("null ")
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
	fmt.Println()
}

func main() {
	// Test case 1
	root1 := createBinaryTree([][]int{{20, 15, 1}, {20, 17, 0}, {50, 20, 1}, {50, 80, 0}})
	printTree(root1)

	// Test case 2
	root2 := createBinaryTree([][]int{{1, 2, 1}, {2, 3, 0}, {3, 4, 1}})
	printTree(root2)
}
```

## 2198 — Number Of Single Divisor Triplets

```go
package main

// LeetCode #2198: Number of Single Divisor Triplets
// https://leetcode.com/problems/number-of-single-divisor-triplets/
// Difficulty: Medium [Paid]
// Time: O(n + m^3) | Space: O(m) where m = max value

import "fmt"

func singleDivisorTriplet(nums []int) int64 {
	count := make([]int, 101)
	for _, v := range nums {
		count[v]++
	}

	var res int64 = 0

	for a := 1; a <= 100; a++ {
		if count[a] == 0 {
			continue
		}
		for b := a; b <= 100; b++ {
			if count[b] == 0 {
				continue
			}
			for c := b; c <= 100; c++ {
				if count[c] == 0 {
					continue
				}

				s := a + b + c
				div := 0
				if s%a == 0 {
					div++
				}
				if s%b == 0 {
					div++
				}
				if s%c == 0 {
					div++
				}
				if div != 1 {
					continue
				}

				if a == b && b == c {
					res += int64(count[a]) * int64(count[a]-1) * int64(count[a]-2) / 6
				} else if a == b {
					res += int64(count[a]) * int64(count[a]-1) / 2 * int64(count[c])
				} else if b == c {
					res += int64(count[a]) * int64(count[b]) * int64(count[b]-1) / 2
				} else {
					res += int64(count[a]) * int64(count[b]) * int64(count[c])
				}
			}
		}
	}
	return res
}

func main() {
	// Test case 1
	fmt.Println(singleDivisorTriplet([]int{4, 6, 7, 3, 2}))
	// Expected: 12

	// Test case 2
	fmt.Println(singleDivisorTriplet([]int{1, 2, 2}))
	// Expected: 0
}
```

## 2201 — Count Artifacts That Can Be Extracted

```go
package main

// LeetCode #2201: Count Artifacts That Can Be Extracted
// https://leetcode.com/problems/count-artifacts-that-can-be-extracted/
// Difficulty: Medium
// Time: O(n + m) | Space: O(n)

import "fmt"

func digArtifacts(n int, artifacts [][]int, dig [][]int) int {
	grid := make([][]bool, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]bool, n)
	}
	for _, d := range dig {
		grid[d[0]][d[1]] = true
	}

	count := 0
	for _, art := range artifacts {
		r1, c1, r2, c2 := art[0], art[1], art[2], art[3]
		extracted := true
		for r := r1; r <= r2 && extracted; r++ {
			for c := c1; c <= c2; c++ {
				if !grid[r][c] {
					extracted = false
					break
				}
			}
		}
		if extracted {
			count++
		}
	}
	return count
}

func main() {
	// Test case 1
	fmt.Println(digArtifacts(2, [][]int{{0, 0, 0, 0}, {0, 1, 1, 1}}, [][]int{{0, 0}, {0, 1}}))
	// Expected: 1

	// Test case 2
	fmt.Println(digArtifacts(2, [][]int{{0, 0, 0, 0}, {0, 1, 1, 1}}, [][]int{{0, 0}, {0, 1}, {1, 1}}))
	// Expected: 2
}
```

## 2202 — Maximize The Topmost Element After K Moves

```go
package main

// LeetCode #2202: Maximize the Topmost Element After K Moves
// https://leetcode.com/problems/maximize-the-topmost-element-after-k-moves/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func maximumTop(nums []int, k int) int {
	n := len(nums)
	if n == 1 {
		if k%2 == 1 {
			return -1
		}
		return nums[0]
	}

	if k == 0 {
		return nums[0]
	}
	if k == 1 {
		return nums[1]
	}

	maxVal := -1
	for i := 0; i < n && i < k-1; i++ {
		if nums[i] > maxVal {
			maxVal = nums[i]
		}
	}
	if k < n {
		if nums[k] > maxVal {
			maxVal = nums[k]
		}
	}
	return maxVal
}

func main() {
	// Test case 1
	fmt.Println(maximumTop([]int{5, 2, 4, 3, 1}, 3))
	// Expected: 5

	// Test case 2
	fmt.Println(maximumTop([]int{2}, 1))
	// Expected: -1

	// Test case 3
	fmt.Println(maximumTop([]int{99, 95, 68, 24, 18}, 69))
	// Expected: 99
}
```

## 2207 — Maximize Number Of Subsequences In A String

```go
package main

// LeetCode #2207: Maximize Number of Subsequences in a String
// https://leetcode.com/problems/maximize-number-of-subsequences-in-a-string/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func maximumSubsequenceCount(text string, pattern string) int64 {
	var count int64 = 0
	var first, second int64 = 0, 0

	for _, ch := range text {
		if byte(ch) == pattern[1] {
			count += first
			second++
		}
		if byte(ch) == pattern[0] {
			first++
		}
	}

	if first > second {
		count += first
	} else {
		count += second
	}
	return count
}

func main() {
	// Test case 1
	fmt.Println(maximumSubsequenceCount("abdcdbc", "ac"))
	// Expected: 4

	// Test case 2
	fmt.Println(maximumSubsequenceCount("aabb", "ab"))
	// Expected: 6

	// Test case 3
	fmt.Println(maximumSubsequenceCount("fwymvreuftzgrcrxczjacqovduqaiig", "yy"))
	// Expected: 2
}
```

## 2208 — Minimum Operations To Halve Array Sum

```go
package main

// LeetCode #2208: Minimum Operations to Halve Array Sum
// https://leetcode.com/problems/minimum-operations-to-halve-array-sum/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"container/heap"
	"fmt"
)

type maxHeap []float64

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x any)        { *h = append(*h, x.(float64)) }
func (h *maxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func halveArray(nums []int) int {
	h := &maxHeap{}
	heap.Init(h)
	total := 0.0
	for _, v := range nums {
		total += float64(v)
		heap.Push(h, float64(v))
	}

	target := total / 2.0
	ops := 0
	for total > target {
		largest := heap.Pop(h).(float64)
		half := largest / 2.0
		total -= half
		heap.Push(h, half)
		ops++
	}
	return ops
}

func main() {
	// Test case 1
	fmt.Println(halveArray([]int{5, 19, 8, 1}))
	// Expected: 3

	// Test case 2
	fmt.Println(halveArray([]int{3, 8, 20}))
	// Expected: 3

	// Test case 3
	fmt.Println(halveArray([]int{1}))
	// Expected: 1
}
```

## 2211 — Count Collisions On A Road

```go
package main

// LeetCode #2211: Count Collisions on a Road
// https://leetcode.com/problems/count-collisions-on-a-road/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func countCollisions(directions string) int {
	n := len(directions)
	left, right := 0, n-1

	for left < n && directions[left] == 'L' {
		left++
	}
	for right >= 0 && directions[right] == 'R' {
		right--
	}

	count := 0
	for i := left; i <= right; i++ {
		if directions[i] != 'S' {
			count++
		}
	}
	return count
}

func main() {
	// Test case 1
	fmt.Println(countCollisions("RLRSLL"))
	// Expected: 5

	// Test case 2
	fmt.Println(countCollisions("LLRR"))
	// Expected: 0

	// Test case 3
	fmt.Println(countCollisions("SSRSSRLLRSLLRSRSSRLRRRRRRS"))
	// Expected: 20
}
```

## 2212 — Maximum Points In An Archery Competition

```go
package main

// LeetCode #2212: Maximum Points in an Archery Competition
// https://leetcode.com/problems/maximum-points-in-an-archery-competition/
// Difficulty: Medium
// Time: O(n * 2^n) | Space: O(n)

import "fmt"

func maximumBobPoints(numArrows int, aliceArrows []int) []int {
	bestScore := 0
	var bestMask int

	for mask := 1; mask < (1 << 12); mask++ {
		arrows := 0
		score := 0
		for i := 0; i < 12; i++ {
			if mask&(1<<i) != 0 {
				arrows += aliceArrows[i] + 1
				score += i
			}
		}
		if arrows <= numArrows && score > bestScore {
			bestScore = score
			bestMask = mask
		}
	}

	result := make([]int, 12)
	used := 0
	for i := 0; i < 12; i++ {
		if bestMask&(1<<i) != 0 {
			result[i] = aliceArrows[i] + 1
			used += result[i]
		}
	}
	// Put remaining arrows in first section (index 0)
	if used < numArrows {
		result[0] += numArrows - used
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println(maximumBobPoints(9, []int{1, 1, 0, 1, 0, 0, 2, 1, 0, 1, 2, 0}))
	// Expected: [0,0,0,0,0,0,0,0,1,1,1,0] or similar valid

	// Test case 2
	fmt.Println(maximumBobPoints(3, []int{0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 2}))
	// Expected: [0,1,1,0,0,0,0,0,0,0,0,0] or similar valid
}
```

## 2214 — Minimum Health To Beat Game

```go
package main

// LeetCode #2214: Minimum Health to Beat Game
// https://leetcode.com/problems/minimum-health-to-beat-game/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func minimumHealth(damage []int, armor int) int64 {
	var total int64 = 0
	maxDmg := 0
	for _, d := range damage {
		total += int64(d)
		if d > maxDmg {
			maxDmg = d
		}
	}
	// We can use armor to reduce the largest damage
	saved := armor
	if maxDmg < saved {
		saved = maxDmg
	}
	return total - int64(saved) + 1
}

func main() {
	// Test case 1
	fmt.Println(minimumHealth([]int{2, 7, 4, 3}, 4))
	// Expected: 13

	// Test case 2
	fmt.Println(minimumHealth([]int{3, 3, 3}, 0))
	// Expected: 10

	// Test case 3
	fmt.Println(minimumHealth([]int{1, 2, 3, 4}, 5))
	// Expected: 7
}
```

## 2216 — Minimum Deletions To Make Array Beautiful

```go
package main

// LeetCode #2216: Minimum Deletions to Make Array Beautiful
// https://leetcode.com/problems/minimum-deletions-to-make-array-beautiful/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func minDeletion(nums []int) int {
	n := len(nums)
	deletions := 0
	i := 0

	for i < n-1 {
		idx := i - deletions
		if idx%2 == 0 && nums[i] == nums[i+1] {
			deletions++
			i++
		} else {
			i++
		}
	}

	// If after deletions the array length is odd, delete last element
	if (n-deletions)%2 == 1 {
		deletions++
	}
	return deletions
}

func main() {
	// Test case 1
	fmt.Println(minDeletion([]int{1, 1, 2, 3, 5}))
	// Expected: 1

	// Test case 2
	fmt.Println(minDeletion([]int{1, 1, 2, 2, 3, 3}))
	// Expected: 2
}
```

## 2217 — Find Palindrome With Fixed Length

```go
package main

// LeetCode #2217: Find Palindrome With Fixed Length
// https://leetcode.com/problems/find-palindrome-with-fixed-length/
// Difficulty: Medium
// Time: O(n * len) | Space: O(1)

import "fmt"

func kthPalindrome(queries []int, intLength int) []int64 {
	halfLen := (intLength + 1) / 2
	start := 1
	for i := 1; i < halfLen; i++ {
		start *= 10
	}

	result := make([]int64, len(queries))
	for i, q := range queries {
		val := int64(start + q - 1)
		if val > int64(start*10-1) {
			result[i] = -1
			continue
		}
		s := fmt.Sprintf("%d", val)
		runes := []rune(s)
		// Mirror: for odd length, skip last char
		mirrorLen := halfLen - (intLength % 2)
		for j := mirrorLen - 1; j >= 0; j-- {
			s += string(runes[j])
		}
		fmt.Sscanf(s, "%d", &val)
		result[i] = val
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println(kthPalindrome([]int{1, 2, 3, 4, 5, 90}, 3))
	// Expected: [101, 111, 121, 131, 141, 999]

	// Test case 2
	fmt.Println(kthPalindrome([]int{2, 4, 6}, 4))
	// Expected: [1111, 1331, 1551]
}
```

## 2219 — Maximum Sum Score Of Array

```go
package main

// LeetCode #2219: Maximum Sum Score of Array
// https://leetcode.com/problems/maximum-sum-score-of-array/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func maximumSumScore(nums []int) int64 {
	n := len(nums)
	total := int64(0)
	for _, v := range nums {
		total += int64(v)
	}

	maxScore := int64(-1 << 63)
	var prefix int64 = 0
	for i := 0; i < n; i++ {
		prefix += int64(nums[i])
		suffix := total - prefix + int64(nums[i])
		score := prefix
		if suffix > score {
			score = suffix
		}
		if score > maxScore {
			maxScore = score
		}
	}
	return maxScore
}

func main() {
	// Test case 1
	fmt.Println(maximumSumScore([]int{1, 2, 3, 4, 5}))
	// Expected: 15

	// Test case 2
	fmt.Println(maximumSumScore([]int{-5, 10, -3, 4}))
	// Expected: 11

	// Test case 3
	fmt.Println(maximumSumScore([]int{-1, -2, -3}))
	// Expected: -1
}
```

## 2221 — Find Triangular Sum Of An Array

```go
package main

// LeetCode #2221: Find Triangular Sum of an Array
// https://leetcode.com/problems/find-triangular-sum-of-an-array/
// Difficulty: Medium
// Time: O(n^2) | Space: O(1)

import "fmt"

func triangularSum(nums []int) int {
	n := len(nums)
	for n > 1 {
		for i := 0; i < n-1; i++ {
			nums[i] = (nums[i] + nums[i+1]) % 10
		}
		n--
	}
	return nums[0]
}

func main() {
	// Test case 1
	fmt.Println(triangularSum([]int{1, 2, 3, 4, 5}))
	// Expected: 8

	// Test case 2
	fmt.Println(triangularSum([]int{5}))
	// Expected: 5

	// Test case 3
	fmt.Println(triangularSum([]int{2, 6, 6, 6}))
	// Expected: 4
}
```

## 2222 — Number Of Ways To Select Buildings

```go
package main

// LeetCode #2222: Number of Ways to Select Buildings
// https://leetcode.com/problems/number-of-ways-to-select-buildings/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func numberOfWays(s string) int64 {
	var total0, total1 int64 = 0, 0
	for _, ch := range s {
		if ch == '0' {
			total0++
		} else {
			total1++
		}
	}

	var ways, prefix0, prefix1 int64 = 0, 0, 0
	for _, ch := range s {
		if ch == '0' {
			ways += prefix1 * (total1 - prefix1)
			prefix0++
		} else {
			ways += prefix0 * (total0 - prefix0)
			prefix1++
		}
	}
	return ways
}

func main() {
	// Test case 1
	fmt.Println(numberOfWays("001101"))
	// Expected: 6

	// Test case 2
	fmt.Println(numberOfWays("11100"))
	// Expected: 0

	// Test case 3
	fmt.Println(numberOfWays("0001100100"))
	// Expected: 12
}
```

## 2225 — Find Players With Zero Or One Losses

```go
package main

// LeetCode #2225: Find Players With Zero or One Losses
// https://leetcode.com/problems/find-players-with-zero-or-one-losses/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func findWinners(matches [][]int) [][]int {
	losses := make(map[int]int)
	players := make(map[int]bool)

	for _, m := range matches {
		winner, loser := m[0], m[1]
		players[winner] = true
		players[loser] = true
		losses[loser]++
	}

	winners := []int{}
	oneLoss := []int{}
	for p := range players {
		l := losses[p]
		if l == 0 {
			winners = append(winners, p)
		} else if l == 1 {
			oneLoss = append(oneLoss, p)
		}
	}

	sort.Ints(winners)
	sort.Ints(oneLoss)

	return [][]int{winners, oneLoss}
}

func main() {
	// Test case 1
	fmt.Println(findWinners([][]int{{1, 3}, {2, 3}, {3, 6}, {5, 6}, {5, 7}, {4, 5}, {4, 8}, {4, 9}, {10, 4}, {10, 9}}))
	// Expected: [[1,2,10],[4,5,7,8]]

	// Test case 2
	fmt.Println(findWinners([][]int{{2, 3}, {1, 3}, {5, 4}, {6, 4}}))
	// Expected: [[1,2,5,6],[]]
}
```

## 2226 — Maximum Candies Allocated To K Children

```go
package main

// LeetCode #2226: Maximum Candies Allocated to K Children
// https://leetcode.com/problems/maximum-candies-allocated-to-k-children/
// Difficulty: Medium
// Time: O(n log m) | Space: O(1)

import "fmt"

func maximumCandies(candies []int, k int64) int {
	lo, hi := 1, 0
	for _, c := range candies {
		if c > hi {
			hi = c
		}
	}

	result := 0
	for lo <= hi {
		mid := lo + (hi-lo)/2
		var count int64 = 0
		for _, c := range candies {
			count += int64(c / mid)
		}
		if count >= k {
			result = mid
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println(maximumCandies([]int{5, 8, 6}, 3))
	// Expected: 5

	// Test case 2
	fmt.Println(maximumCandies([]int{2, 5}, 11))
	// Expected: 0

	// Test case 3
	fmt.Println(maximumCandies([]int{4, 7, 5}, 16))
	// Expected: 1
}
```

## 2228 — Users With Two Purchases Within Seven Days

```go
package main

// LeetCode #2228: Users With Two Purchases Within Seven Days
// https://leetcode.com/problems/users-with-two-purchases-within-seven-days/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func findUsers(purchases [][]int) []int {
	// Group purchases by user
	userPurchases := make(map[int][]int)
	for _, p := range purchases {
		userID, date := p[0], p[1]
		userPurchases[userID] = append(userPurchases[userID], date)
	}

	result := []int{}
	for userID, dates := range userPurchases {
		if len(dates) < 2 {
			continue
		}
		sort.Ints(dates)
		for i := 1; i < len(dates); i++ {
			if dates[i]-dates[i-1] <= 7 {
				result = append(result, userID)
				break
			}
		}
	}
	sort.Ints(result)
	return result
}

func main() {
	// Test case 1
	fmt.Println(findUsers([][]int{{1, 1}, {2, 2}, {1, 7}, {1, 15}, {2, 3}}))
	// Expected: [1]

	// Test case 2
	fmt.Println(findUsers([][]int{{1, 1}, {2, 1}, {3, 1}}))
	// Expected: []
}
```

## 2232 — Minimize Result By Adding Parentheses To Expression

```go
package main

// LeetCode #2232: Minimize Result by Adding Parentheses to Expression
// https://leetcode.com/problems/minimize-result-by-adding-parentheses-to-expression/
// Difficulty: Medium
// Time: O(n^2) | Space: O(1)

import (
	"fmt"
	"strconv"
)

func minimizeResult(expression string) string {
	plusIdx := 0
	for i, ch := range expression {
		if ch == '+' {
			plusIdx = i
			break
		}
	}
	left := expression[:plusIdx]
	right := expression[plusIdx+1:]

	minVal := int64(1 << 60)
	var bestLeft, bestRight int

	for i := 0; i < len(left); i++ {
		for j := 1; j <= len(right); j++ {
			leftPart1 := int64(1)
			if i > 0 {
				v, _ := strconv.ParseInt(left[:i], 10, 64)
				leftPart1 = v
			}
			middleLeft, _ := strconv.ParseInt(left[i:], 10, 64)
			middleRight, _ := strconv.ParseInt(right[:j], 10, 64)
			rightPart2 := int64(1)
			if j < len(right) {
				v, _ := strconv.ParseInt(right[j:], 10, 64)
				rightPart2 = v
			}

			val := leftPart1 * (middleLeft + middleRight) * rightPart2
			if val < minVal {
				minVal = val
				bestLeft = i
				bestRight = j
			}
		}
	}

	result := left[:bestLeft] + "(" + left[bestLeft:] + "+" + right[:bestRight] + ")" + right[bestRight:]
	return result
}

func main() {
	// Test case 1
	fmt.Println(minimizeResult("247+38"))
	// Expected: "2(47+38)" or "247(+3)8"

	// Test case 2
	fmt.Println(minimizeResult("12+34"))
	// Expected: "1(2+3)4"

	// Test case 3
	fmt.Println(minimizeResult("999+999"))
	// Expected: "(999+999)"
}
```

## 2233 — Maximum Product After K Increments

```go
package main

// LeetCode #2233: Maximum Product After K Increments
// https://leetcode.com/problems/maximum-product-after-k-increments/
// Difficulty: Medium
// Time: O(n + k log n) | Space: O(n)

import (
	"container/heap"
	"fmt"
)

type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func maximumProduct(nums []int, k int) int {
	const mod = 1_000_000_007
	h := &minHeap{}
	heap.Init(h)
	for _, v := range nums {
		heap.Push(h, v)
	}
	for i := 0; i < k; i++ {
		v := heap.Pop(h).(int)
		heap.Push(h, v+1)
	}
	prod := 1
	for h.Len() > 0 {
		prod = (prod * heap.Pop(h).(int)) % mod
	}
	return prod
}

func main() {
	// Test case 1
	fmt.Println(maximumProduct([]int{0, 4}, 5))
	// Expected: 20

	// Test case 2
	fmt.Println(maximumProduct([]int{6, 3, 3, 2}, 2))
	// Expected: 216
}
```

## 2237 — Count Positions On Street With Required Brightness

```go
package main

// LeetCode #2237: Count Positions on Street With Required Brightness
// https://leetcode.com/problems/count-positions-on-street-with-required-brightness/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func meetRequirement(n int, lights [][]int, requirement []int) int {
	diff := make([]int, n+2)
	for _, l := range lights {
		pos, r := l[0], l[1]
		left := pos - r
		if left < 0 {
			left = 0
		}
		right := pos + r
		if right > n {
			right = n
		}
		diff[left]++
		diff[right+1]--
	}

	count := 0
	brightness := 0
	for i := 0; i < n; i++ {
		brightness += diff[i]
		if brightness >= requirement[i] {
			count++
		}
	}
	return count
}

func main() {
	// Test case 1
	fmt.Println(meetRequirement(5, [][]int{{0, 1}, {2, 1}, {3, 2}}, []int{0, 2, 1, 4, 1}))
	// Expected: 3

	// Test case 2
	fmt.Println(meetRequirement(3, [][]int{{1, 1}}, []int{0, 1, 0}))
	// Expected: 2
}
```

## 2238 — Number Of Times A Driver Was A Passenger

```go
package main

// LeetCode #2238: Number of Times a Driver Was a Passenger
// https://leetcode.com/problems/number-of-times-a-driver-was-a-passenger/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func countPassengers(rides [][]int) []int {
	// Each ride: [driver_id, passenger_id]
	passengerCount := make(map[int]int)
	drivers := make(map[int]bool)

	for _, ride := range rides {
		driver, passenger := ride[0], ride[1]
		drivers[driver] = true
		passengerCount[passenger]++
	}

	maxID := 0
	for id := range drivers {
		if id > maxID {
			maxID = id
		}
	}

	result := make([]int, maxID+1)
	for id := 1; id <= maxID; id++ {
		if drivers[id] {
			result[id] = passengerCount[id]
		}
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println(countPassengers([][]int{{1, 2}, {2, 3}, {3, 1}}))
	// Expected: [0, 1, 1, 1]

	// Test case 2
	fmt.Println(countPassengers([][]int{{1, 2}, {1, 3}, {1, 4}}))
	// Expected: [0, 0]
}
```

