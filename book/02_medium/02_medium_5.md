# Medium (Sedang) — Problem 0708–0900

## 0708 — Insert Into A Sorted Circular Linked List

```go
package main

// LeetCode #708: Insert into a Sorted Circular Linked List
// https://leetcode.com/problems/insert-into-a-sorted-circular-linked-list/
// Difficulty: Medium [Paid]
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	head := &CNode{Val: 3}
	head.Next = &CNode{Val: 4}
	head.Next.Next = &CNode{Val: 1}
	head.Next.Next.Next = head

	result := insert(head, 2)
	fmt.Println(result.Val)
}

type CNode struct {
	Val  int
	Next *CNode
}

func insert(head *CNode, insertVal int) *CNode {
	newNode := &CNode{Val: insertVal}
	if head == nil {
		newNode.Next = newNode
		return newNode
	}

	curr := head
	for curr.Next != head {
		if curr.Val <= insertVal && insertVal <= curr.Next.Val {
			break
		}
		if curr.Val > curr.Next.Val {
			if insertVal >= curr.Val || insertVal <= curr.Next.Val {
				break
			}
		}
		curr = curr.Next
	}

	newNode.Next = curr.Next
	curr.Next = newNode
	return head
}
```

## 0712 — Minimum Ascii Delete Sum For Two Strings

```go
package main

// LeetCode #712: Minimum ASCII Delete Sum for Two Strings
// https://leetcode.com/problems/minimum-ascii-delete-sum-for-two-strings/
// Difficulty: Medium
// Time: O(m * n)
// Space: O(m * n)

import "fmt"

func main() {
	fmt.Println(minimumDeleteSum("sea", "eat"))
	fmt.Println(minimumDeleteSum("delete", "leet"))
}

func minimumDeleteSum(s1 string, s2 string) int {
	m, n := len(s1), len(s2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		dp[i][0] = dp[i-1][0] + int(s1[i-1])
	}
	for j := 1; j <= n; j++ {
		dp[0][j] = dp[0][j-1] + int(s2[j-1])
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j]+int(s1[i-1]), dp[i][j-1]+int(s2[j-1]))
			}
		}
	}

	return dp[m][n]
}
```

## 0713 — Subarray Product Less Than K

```go
package main

// LeetCode #713: Subarray Product Less Than K
// https://leetcode.com/problems/subarray-product-less-than-k/
// Difficulty: Medium
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100))
	fmt.Println(numSubarrayProductLessThanK([]int{1, 2, 3}, 0))
}

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}

	count := 0
	product := 1
	left := 0

	for right := 0; right < len(nums); right++ {
		product *= nums[right]

		for product >= k {
			product /= nums[left]
			left++
		}

		count += right - left + 1
	}

	return count
}
```

## 0714 — Best Time To Buy And Sell Stock With Transaction Fee

```go
package main

// LeetCode #714: Best Time to Buy and Sell Stock with Transaction Fee
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/
// Difficulty: Medium
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{1, 3, 2, 8, 4, 9}, 2))
	fmt.Println(maxProfit([]int{1, 3, 7, 5, 10, 3}, 3))
}

func maxProfit(prices []int, fee int) int {
	cash := 0
	hold := -prices[0]

	for i := 1; i < len(prices); i++ {
		cash = max(cash, hold+prices[i]-fee)
		hold = max(hold, cash-prices[i])
	}

	return cash
}
```

## 0718 — Maximum Length Of Repeated Subarray

```go
package main

// LeetCode #718: Maximum Length of Repeated Subarray
// https://leetcode.com/problems/maximum-length-of-repeated-subarray/
// Difficulty: Medium
// Time: O(m * n)
// Space: O(m * n)

import "fmt"

func main() {
	fmt.Println(findLength([]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7}))
	fmt.Println(findLength([]int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}))
}

func findLength(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	maxLen := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > maxLen {
					maxLen = dp[i][j]
				}
			}
		}
	}

	return maxLen
}
```

## 0720 — Longest Word In Dictionary

```go
package main

// LeetCode #720: Longest Word in Dictionary
// https://leetcode.com/problems/longest-word-in-dictionary/
// Difficulty: Medium
// Time: O(n * L + n log n)
// Space: O(n * L)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(longestWord([]string{"w", "wo", "wor", "worl", "world"}))
	fmt.Println(longestWord([]string{"a", "banana", "app", "appl", "ap", "apply", "apple"}))
}

func longestWord(words []string) string {
	wordSet := make(map[string]bool)
	for _, w := range words {
		wordSet[w] = true
	}

	sort.Strings(words)

	result := ""
	for _, w := range words {
		if len(w) <= len(result) {
			continue
		}
		valid := true
		for i := 1; i < len(w); i++ {
			if !wordSet[w[:i]] {
				valid = false
				break
			}
		}
		if valid {
			result = w
		}
	}

	return result
}
```

## 0721 — Accounts Merge

```go
package main

// LeetCode #721: Accounts Merge
// https://leetcode.com/problems/accounts-merge/
// Difficulty: Medium
// Time: O(nk * alpha(nk))
// Space: O(nk)

import (
	"fmt"
	"sort"
)

func main() {
	accounts := [][]string{
		{"John", "johnsmith@mail.com", "john00@mail.com"},
		{"John", "johnnybravo@mail.com"},
		{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
		{"Mary", "mary@mail.com"},
	}
	fmt.Println(accountsMerge(accounts))
}

func accountsMerge(accounts [][]string) [][]string {
	parent := make(map[string]string)
	owner := make(map[string]string)

	var find func(x string) string
	find = func(x string) string {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func(x, y string) {
		px, py := find(x), find(y)
		if px != py {
			parent[px] = py
		}
	}

	for _, acc := range accounts {
		name := acc[0]
		firstEmail := acc[1]
		for _, email := range acc[1:] {
			parent[email] = email
			owner[email] = name
			union(firstEmail, email)
		}
	}

	groups := make(map[string][]string)
	for email := range parent {
		root := find(email)
		groups[root] = append(groups[root], email)
	}

	result := make([][]string, 0, len(groups))
	for root, emails := range groups {
		sort.Strings(emails)
		merged := append([]string{owner[root]}, emails...)
		result = append(result, merged)
	}

	return result
}
```

## 0722 — Remove Comments

```go
package main

// LeetCode #722: Remove Comments
// https://leetcode.com/problems/remove-comments/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import (
	"fmt"
	"strings"
)

func main() {
	source := []string{
		"/*Test program */",
		"int main()",
		"{ ",
		"  // variable declaration ",
		"int a, b, c;",
		"/* This is a test",
		"   multiline  ",
		"   comment for ",
		"   testing */",
		"a = b + c;",
		"}",
	}
	result := removeComments(source)
	for _, line := range result {
		fmt.Println(line)
	}
}

func removeComments(source []string) []string {
	result := make([]string, 0)
	inBlock := false
	var current strings.Builder

	for _, line := range source {
		i := 0
		n := len(line)

		if !inBlock {
			current.Reset()
		}

		for i < n {
			if !inBlock && i+1 < n && line[i] == '/' && line[i+1] == '*' {
				inBlock = true
				i += 2
			} else if inBlock && i+1 < n && line[i] == '*' && line[i+1] == '/' {
				inBlock = false
				i += 2
			} else if !inBlock && i+1 < n && line[i] == '/' && line[i+1] == '/' {
				break
			} else if !inBlock {
				current.WriteByte(line[i])
				i++
			} else {
				i++
			}
		}

		if !inBlock && current.Len() > 0 {
			result = append(result, current.String())
		}
	}

	return result
}
```

## 0723 — Candy Crush

```go
package main

// LeetCode #723: Candy Crush
// https://leetcode.com/problems/candy-crush/
// Difficulty: Medium [Paid]
// Time: O(R * C * max(R, C))
// Space: O(1)

import "fmt"

func main() {
	board := [][]int{
		{110, 5, 112, 113, 114},
		{210, 211, 5, 213, 214},
		{310, 311, 3, 313, 314},
		{410, 411, 412, 5, 414},
		{5, 1, 512, 3, 3},
		{610, 4, 1, 613, 614},
		{710, 1, 2, 713, 714},
		{810, 1, 2, 1, 1},
		{1, 1, 2, 2, 2},
		{4, 1, 4, 4, 1014},
	}
	result := candyCrush(board)
	for _, row := range result {
		fmt.Println(row)
	}
}

func candyCrush(board [][]int) [][]int {
	rows, cols := len(board), len(board[0])

	for {
		crushed := false

		// Mark horizontal crushes
		for r := 0; r < rows; r++ {
			for c := 0; c < cols-2; c++ {
				val := abs(board[r][c])
				if val != 0 && abs(board[r][c+1]) == val && abs(board[r][c+2]) == val {
					board[r][c] = -val
					board[r][c+1] = -val
					board[r][c+2] = -val
					crushed = true
				}
			}
		}

		// Mark vertical crushes
		for r := 0; r < rows-2; r++ {
			for c := 0; c < cols; c++ {
				val := abs(board[r][c])
				if val != 0 && abs(board[r+1][c]) == val && abs(board[r+2][c]) == val {
					board[r][c] = -val
					board[r+1][c] = -val
					board[r+2][c] = -val
					crushed = true
				}
			}
		}

		if !crushed {
			break
		}

		// Gravity: drop candies
		for c := 0; c < cols; c++ {
			writeRow := rows - 1
			for r := rows - 1; r >= 0; r-- {
				if board[r][c] > 0 {
					board[writeRow][c] = board[r][c]
					writeRow--
				}
			}
			for r := writeRow; r >= 0; r-- {
				board[r][c] = 0
			}
		}
	}

	return board
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

## 0725 — Split Linked List In Parts

```go
package main

// LeetCode #725: Split Linked List in Parts
// https://leetcode.com/problems/split-linked-list-in-parts/
// Difficulty: Medium
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 6}
	head.Next.Next.Next.Next.Next.Next = &ListNode{Val: 7}

	parts := splitListToParts(head, 3)
	for _, p := range parts {
		if p == nil {
			fmt.Println("nil")
		} else {
			fmt.Println(p.Val)
		}
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	length := 0
	for curr := head; curr != nil; curr = curr.Next {
		length++
	}

	partSize := length / k
	extra := length % k

	result := make([]*ListNode, k)
	curr := head

	for i := 0; i < k && curr != nil; i++ {
		result[i] = curr
		size := partSize
		if i < extra {
			size++
		}

		for j := 0; j < size-1; j++ {
			curr = curr.Next
		}

		next := curr.Next
		curr.Next = nil
		curr = next
	}

	return result
}
```

## 0729 — My Calendar I

```go
package main

// LeetCode #729: My Calendar I
// https://leetcode.com/problems/my-calendar-i/
// Difficulty: Medium
// Time: O(log n) per booking
// Space: O(n)

import "fmt"

func main() {
	cal := ConstructorCalendar()
	fmt.Println(cal.Book(10, 20))
	fmt.Println(cal.Book(15, 25))
	fmt.Println(cal.Book(20, 30))
}

type MyCalendar struct {
	books [][2]int
}

func ConstructorCalendar() MyCalendar {
	return MyCalendar{}
}

func (c *MyCalendar) Book(start int, end int) bool {
	for _, b := range c.books {
		if max(b[0], start) < min(b[1], end) {
			return false
		}
	}
	c.books = append(c.books, [2]int{start, end})
	return true
}
```

## 0731 — My Calendar Ii

```go
package main

// LeetCode #731: My Calendar II
// https://leetcode.com/problems/my-calendar-ii/
// Difficulty: Medium
// Time: O(n^2) per booking
// Space: O(n)

import "fmt"

func main() {
	cal := ConstructorCalendar2()
	fmt.Println(cal.Book(10, 20))
	fmt.Println(cal.Book(50, 60))
	fmt.Println(cal.Book(10, 40))
	fmt.Println(cal.Book(5, 15))
	fmt.Println(cal.Book(5, 10))
	fmt.Println(cal.Book(25, 55))
}

type MyCalendarTwo struct {
	books    [][2]int
	overlaps [][2]int
}

func ConstructorCalendar2() MyCalendarTwo {
	return MyCalendarTwo{}
}

func (c *MyCalendarTwo) Book(start int, end int) bool {
	for _, o := range c.overlaps {
		if max(o[0], start) < min(o[1], end) {
			return false
		}
	}
	for _, b := range c.books {
		if max(b[0], start) < min(b[1], end) {
			c.overlaps = append(c.overlaps, [2]int{max(b[0], start), min(b[1], end)})
		}
	}
	c.books = append(c.books, [2]int{start, end})
	return true
}
```

## 0735 — Asteroid Collision

```go
package main

// LeetCode #735: Asteroid Collision
// https://leetcode.com/problems/asteroid-collision/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(asteroidCollision([]int{5, 10, -5}))
	fmt.Println(asteroidCollision([]int{8, -8}))
	fmt.Println(asteroidCollision([]int{10, 2, -5}))
}

func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0)

	for _, a := range asteroids {
		for len(stack) > 0 && a < 0 && stack[len(stack)-1] > 0 {
			top := stack[len(stack)-1]
			if top+ a < 0 {
				stack = stack[:len(stack)-1]
			} else if top+ a == 0 {
				stack = stack[:len(stack)-1]
				a = 0
				break
			} else {
				a = 0
				break
			}
		}
		if a != 0 {
			stack = append(stack, a)
		}
	}

	return stack
}
```

## 0737 — Sentence Similarity Ii

```go
package main

// LeetCode #737: Sentence Similarity II
// https://leetcode.com/problems/sentence-similarity-ii/
// Difficulty: Medium [Paid]
// Time: O(n * alpha(n))
// Space: O(n)

import "fmt"

func main() {
	pairs := [][]string{
		{"great", "fine"},
		{"drama", "acting"},
		{"fine", "good"},
	}
	fmt.Println(areSentencesSimilarTwo([]string{"great", "acting", "skills"}, []string{"fine", "drama", "talent"}, pairs))
}

func areSentencesSimilarTwo(words1 []string, words2 []string, pairs [][]string) bool {
	if len(words1) != len(words2) {
		return false
	}

	parent := make(map[string]string)

	var find func(x string) string
	find = func(x string) string {
		if _, ok := parent[x]; !ok {
			parent[x] = x
		}
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func(x, y string) {
		px, py := find(x), find(y)
		if px != py {
			parent[px] = py
		}
	}

	for _, p := range pairs {
		union(p[0], p[1])
	}

	for i := 0; i < len(words1); i++ {
		if words1[i] == words2[i] {
			continue
		}
		if find(words1[i]) != find(words2[i]) {
			return false
		}
	}

	return true
}
```

## 0738 — Monotone Increasing Digits

```go
package main

// LeetCode #738: Monotone Increasing Digits
// https://leetcode.com/problems/monotone-increasing-digits/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(monotoneIncreasingDigits(10))
	fmt.Println(monotoneIncreasingDigits(1234))
	fmt.Println(monotoneIncreasingDigits(332))
}

func monotoneIncreasingDigits(n int) int {
	s := []byte(strconv.Itoa(n))
	i := 1

	for i < len(s) && s[i] >= s[i-1] {
		i++
	}

	if i == len(s) {
		return n
	}

	for i > 0 && s[i] < s[i-1] {
		s[i-1]--
		i--
	}

	for j := i + 1; j < len(s); j++ {
		s[j] = '9'
	}

	result, _ := strconv.Atoi(string(s))
	return result
}
```

## 0739 — Daily Temperatures

```go
package main

// LeetCode #739: Daily Temperatures
// https://leetcode.com/problems/daily-temperatures/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
	fmt.Println(dailyTemperatures([]int{30, 40, 50, 60}))
	fmt.Println(dailyTemperatures([]int{30, 60, 90}))
}

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	result := make([]int, n)
	stack := make([]int, 0)

	for i := 0; i < n; i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[idx] = i - idx
		}
		stack = append(stack, i)
	}

	return result
}
```

## 0740 — Delete And Earn

```go
package main

// LeetCode #740: Delete and Earn
// https://leetcode.com/problems/delete-and-earn/
// Difficulty: Medium
// Time: O(n + k) where k is max value
// Space: O(k)

import "fmt"

func main() {
	fmt.Println(deleteAndEarn([]int{3, 4, 2}))
	fmt.Println(deleteAndEarn([]int{2, 2, 3, 3, 3, 4}))
}

func deleteAndEarn(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxVal := 0
	for _, n := range nums {
		if n > maxVal {
			maxVal = n
		}
	}

	values := make([]int, maxVal+1)
	for _, n := range nums {
		values[n] += n
	}

	prev2, prev1 := 0, values[1]
	for i := 2; i <= maxVal; i++ {
		curr := max(prev1, prev2+values[i])
		prev2, prev1 = prev1, curr
	}

	return prev1
}
```

## 0742 — Closest Leaf In A Binary Tree

```go
package main

// LeetCode #742: Closest Leaf in a Binary Tree
// https://leetcode.com/problems/closest-leaf-in-a-binary-tree/
// Difficulty: Medium [Paid]
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	root := &BTNode{Val: 1}
	root.Left = &BTNode{Val: 3}
	root.Right = &BTNode{Val: 2}
	fmt.Println(findClosestLeaf(root, 1))
}

type BTNode struct {
	Val   int
	Left  *BTNode
	Right *BTNode
}

func findClosestLeaf(root *BTNode, k int) int {
	graph := make(map[int][]int)
	leaves := make(map[int]bool)
	visited := make(map[int]bool)

	var buildGraph func(node *BTNode)
	buildGraph = func(node *BTNode) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			leaves[node.Val] = true
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

	queue := []int{k}
	visited[k] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if leaves[node] {
			return node
		}
		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}

	return -1
}
```

## 0743 — Network Delay Time

```go
package main

// LeetCode #743: Network Delay Time
// https://leetcode.com/problems/network-delay-time/
// Difficulty: Medium
// Time: O(n + E log V)
// Space: O(n + E)

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	fmt.Println(networkDelayTime([][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, 4, 2))
	fmt.Println(networkDelayTime([][]int{{1, 2, 1}}, 2, 1))
}

type Edge struct {
	node int
	time int
}

type MinHeap2 []Edge

func (h MinHeap2) Len() int           { return len(h) }
func (h MinHeap2) Less(i, j int) bool { return h[i].time < h[j].time }
func (h MinHeap2) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap2) Push(x interface{}) { *h = append(*h, x.(Edge)) }
func (h *MinHeap2) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func networkDelayTime(times [][]int, n int, k int) int {
	graph := make([][]Edge, n+1)
	for _, t := range times {
		graph[t[0]] = append(graph[t[0]], Edge{t[1], t[2]})
	}

	dist := make([]int, n+1)
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[k] = 0

	h := &MinHeap2{}
	heap.Init(h)
	heap.Push(h, Edge{k, 0})

	for h.Len() > 0 {
		cur := heap.Pop(h).(Edge)
		if cur.time > dist[cur.node] {
			continue
		}
		for _, e := range graph[cur.node] {
			if nd := cur.time + e.time; nd < dist[e.node] {
				dist[e.node] = nd
				heap.Push(h, Edge{e.node, nd})
			}
		}
	}

	maxTime := 0
	for i := 1; i <= n; i++ {
		if dist[i] == math.MaxInt32 {
			return -1
		}
		if dist[i] > maxTime {
			maxTime = dist[i]
		}
	}
	return maxTime
}
```

## 0750 — Number Of Corner Rectangles

```go
package main

// LeetCode #750: Number of Corner Rectangles
// https://leetcode.com/problems/number-of-corner-rectangles/
// Difficulty: Medium [Paid]
// Time: O(R * C^2)
// Space: O(1)

import "fmt"

func main() {
	grid := [][]int{
		{1, 0, 0, 1, 0},
		{0, 0, 1, 0, 1},
		{0, 0, 0, 1, 0},
		{1, 0, 1, 0, 1},
	}
	fmt.Println(countCornerRectangles(grid))
}

func countCornerRectangles(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	count := 0

	for c1 := 0; c1 < cols; c1++ {
		for c2 := c1 + 1; c2 < cols; c2++ {
			pairs := 0
			for r := 0; r < rows; r++ {
				if grid[r][c1] == 1 && grid[r][c2] == 1 {
					pairs++
				}
			}
			count += pairs * (pairs - 1) / 2
		}
	}

	return count
}
```

## 0751 — Ip To Cidr

```go
package main

// LeetCode #751: IP to CIDR
// https://leetcode.com/problems/ip-to-cidr/
// Difficulty: Medium [Paid]
// Time: O(log n) for each block
// Space: O(1)

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	result := ipToCIDR("255.0.0.7", 10)
	for _, s := range result {
		fmt.Println(s)
	}
}

func ipToCIDR(ip string, n int) []string {
	start := ipToInt(ip)
	result := make([]string, 0)

	for n > 0 {
		mask := max(0, start&-start)
		for mask > n {
			mask >>= 1
		}

		result = append(result, intToIP(start)+"/"+strconv.Itoa(32-trailingZeros(mask)))
		start += mask
		n -= mask
	}

	return result
}

func ipToInt(ip string) int {
	parts := strings.Split(ip, ".")
	result := 0
	for _, p := range parts {
		val, _ := strconv.Atoi(p)
		result = result*256 + val
	}
	return result
}

func intToIP(n int) string {
	return fmt.Sprintf("%d.%d.%d.%d", n>>24, (n>>16)&255, (n>>8)&255, n&255)
}

func trailingZeros(x int) int {
	if x == 0 {
		return 32
	}
	count := 0
	for x&1 == 0 {
		x >>= 1
		count++
	}
	return count
}
```

## 0752 — Open The Lock

```go
package main

// LeetCode #752: Open the Lock
// https://leetcode.com/problems/open-the-lock/
// Difficulty: Medium
// Time: O(10^4 * 8) ~ O(1)
// Space: O(10^4)

import "fmt"

func main() {
	fmt.Println(openLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202"))
	fmt.Println(openLock([]string{"8888"}, "0009"))
}

func openLock(deadends []string, target string) int {
	dead := make(map[string]bool)
	for _, d := range deadends {
		dead[d] = true
	}

	if dead["0000"] {
		return -1
	}

	visited := make(map[string]bool)
	queue := []string{"0000"}
	visited["0000"] = true
	steps := 0

	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			curr := queue[i]
			if curr == target {
				return steps
			}

			for j := 0; j < 4; j++ {
				for _, d := range []int{-1, 1} {
					next := []byte(curr)
					next[j] = byte('0' + (int(next[j]-'0')+d+10)%10)
					s := string(next)
					if !visited[s] && !dead[s] {
						visited[s] = true
						queue = append(queue, s)
					}
				}
			}
		}
		queue = queue[n:]
		steps++
	}

	return -1
}
```

## 0754 — Reach A Number

```go
package main

// LeetCode #754: Reach a Number
// https://leetcode.com/problems/reach-a-number/
// Difficulty: Medium
// Time: O(sqrt(target))
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(reachNumber(3))
	fmt.Println(reachNumber(2))
}

func reachNumber(target int) int {
	if target < 0 {
		target = -target
	}

	sum := 0
	steps := 0

	for sum < target || (sum-target)%2 != 0 {
		steps++
		sum += steps
	}

	return steps
}
```

## 0755 — Pour Water

```go
package main

// LeetCode #755: Pour Water
// https://leetcode.com/problems/pour-water/
// Difficulty: Medium [Paid]
// Time: O(V * N)
// Space: O(1)

import "fmt"

func main() {
	heights := []int{2, 1, 1, 2, 1, 2, 2}
	result := pourWater(heights, 4, 3)
	fmt.Println(result)
}

func pourWater(heights []int, volume int, k int) []int {
	n := len(heights)

	for v := 0; v < volume; v++ {
		pos := k

		// Try left
		left := k
		for left > 0 && heights[left] >= heights[left-1] {
			left--
		}
		for left < k && heights[left] == heights[left+1] {
			left++
		}
		if heights[left] < heights[pos] {
			pos = left
		}

		if pos == k {
			// Try right
			right := k
			for right < n-1 && heights[right] >= heights[right+1] {
				right++
			}
			for right > k && heights[right] == heights[right-1] {
				right--
			}
			if heights[right] < heights[pos] {
				pos = right
			}
		}

		heights[pos]++
	}

	return heights
}
```

## 0756 — Pyramid Transition Matrix

```go
package main

// LeetCode #756: Pyramid Transition Matrix
// https://leetcode.com/problems/pyramid-transition-matrix/
// Difficulty: Medium
// Time: O(7^b) worst case where b is number of blocks
// Space: O(7^b)

import "fmt"

func main() {
	fmt.Println(pyramidTransition("BCD", []string{"BCG", "CDE", "GEA", "FFF"}))
	fmt.Println(pyramidTransition("AAAA", []string{"AAB", "AAC", "BCD", "BBE", "DEF"}))
}

func pyramidTransition(bottom string, allowed []string) bool {
	memo := make(map[string]bool)
	patterns := make(map[string][]byte)

	for _, a := range allowed {
		key := a[:2]
		patterns[key] = append(patterns[key], a[2])
	}

	var dfs func(row string, next string, idx int) bool
	dfs = func(row string, next string, idx int) bool {
		if len(row) == 1 {
			return true
		}

		key := row + "#" + next
		if val, ok := memo[key]; ok {
			return val
		}

		if idx == len(row)-1 {
			if dfs(next, "", 0) {
				memo[key] = true
				return true
			}
			memo[key] = false
			return false
		}

		chars := patterns[row[idx:idx+2]]
		for _, c := range chars {
			if dfs(row, next+string(c), idx+1) {
				memo[key] = true
				return true
			}
		}

		memo[key] = false
		return false
	}

	return dfs(bottom, "", 0)
}
```

## 0758 — Bold Words In String

```go
package main

// LeetCode #758: Bold Words in String
// https://leetcode.com/problems/bold-words-in-string/
// Difficulty: Medium [Paid]
// Time: O(n * L)
// Space: O(n)

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(boldWords([]string{"ab", "bc"}, "aabcd"))
	fmt.Println(boldWords([]string{"abc", "123"}, "abcxyz123"))
}

func boldWords(words []string, s string) string {
	n := len(s)
	bold := make([]bool, n)

	for _, word := range words {
		start := 0
		for {
			idx := strings.Index(s[start:], word)
			if idx == -1 {
				break
			}
			pos := start + idx
			for i := pos; i < pos+len(word); i++ {
				bold[i] = true
			}
			start = pos + 1
		}
	}

	var result strings.Builder
	i := 0
	for i < n {
		if bold[i] {
			result.WriteString("<b>")
			for i < n && bold[i] {
				result.WriteByte(s[i])
				i++
			}
			result.WriteString("</b>")
		} else {
			result.WriteByte(s[i])
			i++
		}
	}

	return result.String()
}
```

## 0763 — Partition Labels

```go
package main

// LeetCode #763: Partition Labels
// https://leetcode.com/problems/partition-labels/
// Difficulty: Medium
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
	fmt.Println(partitionLabels("eccbbbbdec"))
}

func partitionLabels(s string) []int {
	last := make([]int, 26)
	for i := 0; i < len(s); i++ {
		last[s[i]-'a'] = i
	}

	result := make([]int, 0)
	start, end := 0, 0

	for i := 0; i < len(s); i++ {
		if last[s[i]-'a'] > end {
			end = last[s[i]-'a']
		}
		if i == end {
			result = append(result, end-start+1)
			start = i + 1
		}
	}

	return result
}
```

## 0764 — Largest Plus Sign

```go
package main

// LeetCode #764: Largest Plus Sign
// https://leetcode.com/problems/largest-plus-sign/
// Difficulty: Medium
// Time: O(n^2)
// Space: O(n^2)

import "fmt"

func main() {
	fmt.Println(orderOfLargestPlusSign(5, [][]int{{4, 2}}))
	fmt.Println(orderOfLargestPlusSign(1, [][]int{{0, 0}}))
}

func orderOfLargestPlusSign(n int, mines [][]int) int {
	mineSet := make(map[int]bool)
	for _, m := range mines {
		mineSet[m[0]*n+m[1]] = true
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	best := 0

	for r := 0; r < n; r++ {
		count := 0
		for c := 0; c < n; c++ {
			if mineSet[r*n+c] {
				count = 0
			} else {
				count++
			}
			dp[r][c] = count
		}

		count = 0
		for c := n - 1; c >= 0; c-- {
			if mineSet[r*n+c] {
				count = 0
			} else {
				count++
			}
			dp[r][c] = min(dp[r][c], count)
		}
	}

	for c := 0; c < n; c++ {
		count := 0
		for r := 0; r < n; r++ {
			if mineSet[r*n+c] {
				count = 0
			} else {
				count++
			}
			dp[r][c] = min(dp[r][c], count)
		}

		count = 0
		for r := n - 1; r >= 0; r-- {
			if mineSet[r*n+c] {
				count = 0
			} else {
				count++
			}
			dp[r][c] = min(dp[r][c], count)
			if dp[r][c] > best {
				best = dp[r][c]
			}
		}
	}

	return best
}
```

## 0767 — Reorganize String

```go
package main

// LeetCode #767: Reorganize String
// https://leetcode.com/problems/reorganize-string/
// Difficulty: Medium
// Time: O(n log k) where k is alphabet size
// Space: O(n)

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(reorganizeString("aab"))
	fmt.Println(reorganizeString("aaab"))
}

type CharCount struct {
	char byte
	cnt  int
}

type CharHeap []CharCount

func (h CharHeap) Len() int            { return len(h) }
func (h CharHeap) Less(i, j int) bool  { return h[i].cnt > h[j].cnt }
func (h CharHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *CharHeap) Push(x interface{}) { *h = append(*h, x.(CharCount)) }
func (h *CharHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func reorganizeString(s string) string {
	freq := make([]int, 26)
	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
	}

	h := &CharHeap{}
	heap.Init(h)
	for i := 0; i < 26; i++ {
		if freq[i] > 0 {
			heap.Push(h, CharCount{byte(i + 'a'), freq[i]})
		}
	}

	result := make([]byte, 0, len(s))

	for h.Len() >= 2 {
		c1 := heap.Pop(h).(CharCount)
		c2 := heap.Pop(h).(CharCount)

		result = append(result, c1.char, c2.char)

		c1.cnt--
		c2.cnt--
		if c1.cnt > 0 {
			heap.Push(h, c1)
		}
		if c2.cnt > 0 {
			heap.Push(h, c2)
		}
	}

	if h.Len() == 1 {
		c := heap.Pop(h).(CharCount)
		if c.cnt > 1 {
			return ""
		}
		result = append(result, c.char)
	}

	return string(result)
}
```

## 0769 — Max Chunks To Make Sorted

```go
package main

// LeetCode #769: Max Chunks To Make Sorted
// https://leetcode.com/problems/max-chunks-to-make-sorted/
// Difficulty: Medium
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(maxChunksToSorted([]int{4, 3, 2, 1, 0}))
	fmt.Println(maxChunksToSorted([]int{1, 0, 2, 3, 4}))
}

func maxChunksToSorted(arr []int) int {
	count := 0
	maxVal := 0

	for i, val := range arr {
		if val > maxVal {
			maxVal = val
		}
		if maxVal == i {
			count++
		}
	}

	return count
}
```

## 0775 — Global And Local Inversions

```go
package main

// LeetCode #775: Global and Local Inversions
// https://leetcode.com/problems/global-and-local-inversions/
// Difficulty: Medium
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(isIdealPermutation([]int{1, 0, 2}))
	fmt.Println(isIdealPermutation([]int{1, 2, 0}))
}

func isIdealPermutation(nums []int) bool {
	maxVal := -1
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > maxVal {
			maxVal = nums[i]
		}
		if maxVal > nums[i+2] {
			return false
		}
	}
	return true
}
```

## 0776 — Split Bst

```go
package main

// LeetCode #776: Split BST
// https://leetcode.com/problems/split-bst/
// Difficulty: Medium [Paid]
// Time: O(h) where h is tree height
// Space: O(h)

import "fmt"

func main() {
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 6}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 7}

	left, right := splitBST(root, 4)
	fmt.Println(left.Val)
	fmt.Println(right.Val)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func splitBST(root *TreeNode, target int) (*TreeNode, *TreeNode) {
	if root == nil {
		return nil, nil
	}

	if root.Val <= target {
		left, right := splitBST(root.Right, target)
		root.Right = left
		return root, right
	}

	left, right := splitBST(root.Left, target)
	root.Left = right
	return left, root
}
```

## 0777 — Swap Adjacent In Lr String

```go
package main

// LeetCode #777: Swap Adjacent in LR String
// https://leetcode.com/problems/swap-adjacent-in-lr-string/
// Difficulty: Medium
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(canTransform("RXXLRXRXL", "XRLXXRRLX"))
	fmt.Println(canTransform("X", "L"))
}

func canTransform(start string, end string) bool {
	if len(start) != len(end) {
		return false
	}

	n := len(start)
	i, j := 0, 0

	for i < n && j < n {
		for i < n && start[i] == 'X' {
			i++
		}
		for j < n && end[j] == 'X' {
			j++
		}

		if i == n && j == n {
			return true
		}
		if i == n || j == n {
			return false
		}
		if start[i] != end[j] {
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
```

## 0779 — K Th Symbol In Grammar

```go
package main

// LeetCode #779: K-th Symbol in Grammar
// https://leetcode.com/problems/k-th-symbol-in-grammar/
// Difficulty: Medium
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(kthGrammar(1, 1))
	fmt.Println(kthGrammar(2, 1))
	fmt.Println(kthGrammar(2, 2))
}

func kthGrammar(n int, k int) int {
	if n == 1 {
		return 0
	}

	parent := kthGrammar(n-1, (k+1)/2)
	if k%2 == 0 {
		return 1 - parent
	}
	return parent
}
```

## 0781 — Rabbits In Forest

```go
package main

// LeetCode #781: Rabbits in Forest
// https://leetcode.com/problems/rabbits-in-forest/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(numRabbits([]int{1, 1, 2}))
	fmt.Println(numRabbits([]int{10, 10, 10}))
}

func numRabbits(answers []int) int {
	count := make(map[int]int)
	for _, a := range answers {
		count[a]++
	}

	result := 0
	for k, v := range count {
		groupSize := k + 1
		groups := (v + groupSize - 1) / groupSize
		result += groups * groupSize
	}

	return result
}
```

## 0784 — Letter Case Permutation

```go
package main

// LeetCode #784: Letter Case Permutation
// https://leetcode.com/problems/letter-case-permutation/
// Difficulty: Medium
// Time: O(2^n * n)
// Space: O(2^n * n)

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(letterCasePermutation("a1b2"))
	fmt.Println(letterCasePermutation("3z4"))
}

func letterCasePermutation(s string) []string {
	result := make([]string, 0)
	curr := make([]byte, len(s))

	var backtrack func(idx int)
	backtrack = func(idx int) {
		if idx == len(s) {
			result = append(result, string(curr))
			return
		}

		curr[idx] = s[idx]
		if unicode.IsLetter(rune(s[idx])) {
			curr[idx] = byte(unicode.ToLower(rune(s[idx])))
			backtrack(idx + 1)
			curr[idx] = byte(unicode.ToUpper(rune(s[idx])))
			backtrack(idx + 1)
		} else {
			backtrack(idx + 1)
		}
	}

	backtrack(0)
	return result
}
```

## 0785 — Is Graph Bipartite

```go
package main

// LeetCode #785: Is Graph Bipartite?
// https://leetcode.com/problems/is-graph-bipartite/
// Difficulty: Medium
// Time: O(V + E)
// Space: O(V)

import "fmt"

func main() {
	fmt.Println(isBipartite([][]int{{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2}}))
	fmt.Println(isBipartite([][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}}))
}

func isBipartite(graph [][]int) bool {
	n := len(graph)
	color := make([]int, n)

	var dfs func(node int, c int) bool
	dfs = func(node int, c int) bool {
		if color[node] != 0 {
			return color[node] == c
		}
		color[node] = c

		for _, neighbor := range graph[node] {
			if !dfs(neighbor, -c) {
				return false
			}
		}
		return true
	}

	for i := 0; i < n; i++ {
		if color[i] == 0 && !dfs(i, 1) {
			return false
		}
	}

	return true
}
```

## 0786 — K Th Smallest Prime Fraction

```go
package main

// LeetCode #786: K-th Smallest Prime Fraction
// https://leetcode.com/problems/k-th-smallest-prime-fraction/
// Difficulty: Medium
// Time: O(n log max)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(kthSmallestPrimeFraction([]int{1, 2, 3, 5}, 3))
	fmt.Println(kthSmallestPrimeFraction([]int{1, 7}, 1))
}

func kthSmallestPrimeFraction(arr []int, k int) []int {
	n := len(arr)
	left, right := 0.0, 1.0

	for left < right {
		mid := (left + right) / 2.0
		count := 0
		maxFraction := 0.0
		p, q := 0, 1

		j := 1
		for i := 0; i < n; i++ {
			for j < n && float64(arr[i])/float64(arr[j]) > mid {
				j++
			}
			if j == n {
				break
			}
			count += n - j

			fraction := float64(arr[i]) / float64(arr[j])
			if fraction > maxFraction {
				maxFraction = fraction
				p, q = arr[i], arr[j]
			}
		}

		if count == k {
			return []int{p, q}
		} else if count < k {
			left = mid
		} else {
			right = mid
		}
	}

	return nil
}
```

## 0787 — Cheapest Flights Within K Stops

```go
package main

// LeetCode #787: Cheapest Flights Within K Stops
// https://leetcode.com/problems/cheapest-flights-within-k-stops/
// Difficulty: Medium
// Time: O(K * E)
// Space: O(n)

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findCheapestPrice(3, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 0, 2, 1))
	fmt.Println(findCheapestPrice(3, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 0, 2, 0))
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	prices := make([]int, n)
	for i := range prices {
		prices[i] = math.MaxInt32
	}
	prices[src] = 0

	for i := 0; i <= k; i++ {
		temp := make([]int, n)
		copy(temp, prices)
		for _, f := range flights {
			from, to, price := f[0], f[1], f[2]
			if prices[from] != math.MaxInt32 && prices[from]+price < temp[to] {
				temp[to] = prices[from] + price
			}
		}
		prices = temp
	}

	if prices[dst] == math.MaxInt32 {
		return -1
	}
	return prices[dst]
}
```

## 0788 — Rotated Digits

```go
package main

// LeetCode #788: Rotated Digits
// https://leetcode.com/problems/rotated-digits/
// Difficulty: Medium
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(rotatedDigits(10))
	fmt.Println(rotatedDigits(1))
}

func rotatedDigits(n int) int {
	count := 0

	for i := 1; i <= n; i++ {
		if isGood(i) {
			count++
		}
	}

	return count
}

func isGood(n int) bool {
	valid := false
	for n > 0 {
		d := n % 10
		if d == 3 || d == 4 || d == 7 {
			return false
		}
		if d == 2 || d == 5 || d == 6 || d == 9 {
			valid = true
		}
		n /= 10
	}
	return valid
}
```

## 0789 — Escape The Ghosts

```go
package main

// LeetCode #789: Escape The Ghosts
// https://leetcode.com/problems/escape-the-ghosts/
// Difficulty: Medium
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(escapeGhosts([][]int{{1, 0}, {0, 3}}, []int{0, 1}))
	fmt.Println(escapeGhosts([][]int{{1, 0}}, []int{2, 0}))
}

func escapeGhosts(ghosts [][]int, target []int) bool {
	myDist := abs(target[0]) + abs(target[1])

	for _, g := range ghosts {
		ghostDist := abs(g[0]-target[0]) + abs(g[1]-target[1])
		if ghostDist <= myDist {
			return false
		}
	}

	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

## 0790 — Domino And Tromino Tiling

```go
package main

// LeetCode #790: Domino and Tromino Tiling
// https://leetcode.com/problems/domino-and-tromino-tiling/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(numTilings(3))
	fmt.Println(numTilings(1))
	fmt.Println(numTilings(5))
}

func numTilings(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	mod := 1000000007
	dp := make([]int, n+1)
	dp2 := make([]int, n+1)

	dp[0] = 1
	dp[1] = 1
	dp[2] = 2
	dp2[2] = 1

	for i := 3; i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-2] + 2*dp2[i-1]) % mod
		dp2[i] = (dp[i-2] + dp2[i-1]) % mod
	}

	return dp[n]
}
```

## 0791 — Custom Sort String

```go
package main

// LeetCode #791: Custom Sort String
// https://leetcode.com/problems/custom-sort-string/
// Difficulty: Medium
// Time: O(n + m)
// Space: O(1)

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(customSortString("cba", "abcd"))
	fmt.Println(customSortString("bcafg", "abcd"))
}

func customSortString(order string, s string) string {
	freq := make([]int, 26)
	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
	}

	var result strings.Builder
	for i := 0; i < len(order); i++ {
		c := order[i]
		for freq[c-'a'] > 0 {
			result.WriteByte(c)
			freq[c-'a']--
		}
	}

	for i := 0; i < 26; i++ {
		for freq[i] > 0 {
			result.WriteByte(byte(i + 'a'))
			freq[i]--
		}
	}

	return result.String()
}
```

## 0792 — Number Of Matching Subsequences

```go
package main

// LeetCode #792: Number of Matching Subsequences
// https://leetcode.com/problems/number-of-matching-subsequences/
// Difficulty: Medium
// Time: O(n + m * L) where n = len(s), m = len(words)
// Space: O(m)

import "fmt"

func main() {
	fmt.Println(numMatchingSubseq("abcde", []string{"a", "bb", "acd", "ace"}))
	fmt.Println(numMatchingSubseq("dsahjpjauf", []string{"ahjpjau", "ja", "ahbwzgqnuk", "tnmlanowax"}))
}

func numMatchingSubseq(s string, words []string) int {
	buckets := make([][]string, 26)
	for i := range buckets {
		buckets[i] = make([]string, 0)
	}

	for _, w := range words {
		buckets[w[0]-'a'] = append(buckets[w[0]-'a'], w)
	}

	count := 0

	for _, c := range s {
		idx := c - 'a'
		curr := buckets[idx]
		buckets[idx] = make([]string, 0)

		for _, w := range curr {
			if len(w) == 1 {
				count++
			} else {
				buckets[w[1]-'a'] = append(buckets[w[1]-'a'], w[1:])
			}
		}
	}

	return count
}
```

## 0794 — Valid Tic Tac Toe State

```go
package main

// LeetCode #794: Valid Tic-Tac-Toe State
// https://leetcode.com/problems/valid-tic-tac-toe-state/
// Difficulty: Medium
// Time: O(1)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(validTicTacToe([]string{"O  ", "   ", "   "}))
	fmt.Println(validTicTacToe([]string{"XOX", " X ", "   "}))
	fmt.Println(validTicTacToe([]string{"XOX", "O O", "XOX"}))
}

func validTicTacToe(board []string) bool {
	xCount, oCount := 0, 0
	for _, row := range board {
		for _, c := range row {
			if c == 'X' {
				xCount++
			} else if c == 'O' {
				oCount++
			}
		}
	}

	if xCount != oCount && xCount != oCount+1 {
		return false
	}

	xWin := isWinner(board, 'X')
	oWin := isWinner(board, 'O')

	if xWin && oWin {
		return false
	}
	if xWin && xCount != oCount+1 {
		return false
	}
	if oWin && xCount != oCount {
		return false
	}

	return true
}

func isWinner(board []string, player byte) bool {
	// Rows and columns
	for i := 0; i < 3; i++ {
		if board[i][0] == player && board[i][1] == player && board[i][2] == player {
			return true
		}
		if board[0][i] == player && board[1][i] == player && board[2][i] == player {
			return true
		}
	}
	// Diagonals
	if board[0][0] == player && board[1][1] == player && board[2][2] == player {
		return true
	}
	if board[0][2] == player && board[1][1] == player && board[2][0] == player {
		return true
	}
	return false
}
```

## 0795 — Number Of Subarrays With Bounded Maximum

```go
package main

// LeetCode #795: Number of Subarrays with Bounded Maximum
// https://leetcode.com/problems/number-of-subarrays-with-bounded-maximum/
// Difficulty: Medium
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(numSubarrayBoundedMax([]int{2, 1, 4, 3}, 2, 3))
	fmt.Println(numSubarrayBoundedMax([]int{2, 9, 2, 5, 6}, 2, 8))
}

func numSubarrayBoundedMax(nums []int, left int, right int) int {
	count := 0
	prevCount := 0
	prevLessIdx := -1

	for i, val := range nums {
		if val >= left && val <= right {
			prevCount = i - prevLessIdx
			count += prevCount
		} else if val < left {
			count += prevCount
		} else {
			prevCount = 0
			prevLessIdx = i
		}
	}

	return count
}
```

## 0797 — All Paths From Source To Target

```go
package main

// LeetCode #797: All Paths From Source to Target
// https://leetcode.com/problems/all-paths-from-source-to-target/
// Difficulty: Medium
// Time: O(2^n * n)
// Space: O(2^n * n)

import "fmt"

func main() {
	fmt.Println(allPathsSourceTarget([][]int{{1, 2}, {3}, {3}, {}}))
	fmt.Println(allPathsSourceTarget([][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}}))
}

func allPathsSourceTarget(graph [][]int) [][]int {
	result := make([][]int, 0)
	path := make([]int, 0)
	path = append(path, 0)

	var dfs func(node int)
	dfs = func(node int) {
		if node == len(graph)-1 {
			pathCopy := make([]int, len(path))
			copy(pathCopy, path)
			result = append(result, pathCopy)
			return
		}

		for _, neighbor := range graph[node] {
			path = append(path, neighbor)
			dfs(neighbor)
			path = path[:len(path)-1]
		}
	}

	dfs(0)
	return result
}
```

## 0799 — Champagne Tower

```go
package main

// LeetCode #799: Champagne Tower
// https://leetcode.com/problems/champagne-tower/
// Difficulty: Medium
// Time: O(query_row^2)
// Space: O(query_row)

import "fmt"

func main() {
	fmt.Println(champagneTower(1, 1, 1))
	fmt.Println(champagneTower(2, 1, 1))
	fmt.Println(champagneTower(100000009, 33, 17))
}

func champagneTower(poured int, queryRow int, queryGlass int) float64 {
	dp := make([]float64, queryRow+1)
	dp[0] = float64(poured)

	for row := 0; row < queryRow; row++ {
		next := make([]float64, queryRow+2)
		for col := 0; col <= row; col++ {
			if dp[col] > 1.0 {
				excess := (dp[col] - 1.0) / 2.0
				next[col] += excess
				next[col+1] += excess
			}
		}
		dp = next
	}

	if dp[queryGlass] > 1.0 {
		return 1.0
	}
	return dp[queryGlass]
}
```

## 0802 — Find Eventual Safe States

```go
package main

// LeetCode #802: Find Eventual Safe States
// https://leetcode.com/problems/find-eventual-safe-states/
// Difficulty: Medium
// Time: O(V + E)
// Space: O(V)

import "fmt"

func main() {
	fmt.Println(eventualSafeNodes([][]int{{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {}}))
	fmt.Println(eventualSafeNodes([][]int{{1, 2, 3, 4}, {1, 2}, {3, 4}, {0, 4}, {}}))
}

func eventualSafeNodes(graph [][]int) []int {
	n := len(graph)
	state := make([]int, n) // 0=unvisited, 1=visiting, 2=safe

	var dfs func(node int) bool
	dfs = func(node int) bool {
		if state[node] > 0 {
			return state[node] == 2
		}

		state[node] = 1
		for _, neighbor := range graph[node] {
			if !dfs(neighbor) {
				return false
			}
		}
		state[node] = 2
		return true
	}

	result := make([]int, 0)
	for i := 0; i < n; i++ {
		if dfs(i) {
			result = append(result, i)
		}
	}

	return result
}
```

## 0807 — Max Increase To Keep City Skyline

```go
package main

// LeetCode #807: Max Increase to Keep City Skyline
// https://leetcode.com/problems/max-increase-to-keep-city-skyline/
// Difficulty: Medium
// Time: O(n^2)
// Space: O(n)

import "fmt"

func main() {
	grid := [][]int{
		{3, 0, 8, 4},
		{2, 4, 5, 7},
		{9, 2, 6, 3},
		{0, 3, 1, 0},
	}
	fmt.Println(maxIncreaseKeepingSkyline(grid))
}

func maxIncreaseKeepingSkyline(grid [][]int) int {
	n := len(grid)
	rowMax := make([]int, n)
	colMax := make([]int, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] > rowMax[i] {
				rowMax[i] = grid[i][j]
			}
			if grid[i][j] > colMax[j] {
				colMax[j] = grid[i][j]
			}
		}
	}

	total := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			limit := min(rowMax[i], colMax[j])
			total += limit - grid[i][j]
		}
	}

	return total
}
```

## 0808 — Soup Servings

```go
package main

// LeetCode #808: Soup Servings
// https://leetcode.com/problems/soup-servings/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(SoupServings(50))
	fmt.Println(SoupServings(100))
	fmt.Println(SoupServings(800))
}

func SoupServings(n int) float64 {
	if n > 4800 {
		return 1.0
	}
	n = (n + 24) / 25

	memo := make([][]float64, n+1)
	for i := range memo {
		memo[i] = make([]float64, n+1)
	}
	var dfs func(int, int) float64
	dfs = func(a, b int) float64 {
		if a <= 0 && b <= 0 {
			return 0.5
		}
		if a <= 0 {
			return 1.0
		}
		if b <= 0 {
			return 0.0
		}
		if memo[a][b] > 0 {
			return memo[a][b]
		}
		memo[a][b] = 0.25 * (dfs(a-4, b) + dfs(a-3, b-1) + dfs(a-2, b-2) + dfs(a-1, b-3))
		return memo[a][b]
	}

	return dfs(n, n)
}
```

## 0809 — Expressive Words

```go
package main

// LeetCode #809: Expressive Words
// https://leetcode.com/problems/expressive-words/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(ExpressiveWords("heeellooo", []string{"hello", "hi", "helo"}))
	fmt.Println(ExpressiveWords("zzzzzyyyyy", []string{"zzyy", "zy", "zyy"}))
	fmt.Println(ExpressiveWords("abcd", []string{"abc"}))
}

func ExpressiveWords(s string, words []string) int {
	type group struct {
		char byte
		cnt  int
	}

	var encode func(string) []group
	encode = func(str string) []group {
		var groups []group
		for i := 0; i < len(str); {
			j := i
			for j < len(str) && str[j] == str[i] {
				j++
			}
			groups = append(groups, group{str[i], j - i})
			i = j
		}
		return groups
	}

	sGroups := encode(s)
	count := 0

	for _, word := range words {
		wGroups := encode(word)
		if len(wGroups) != len(sGroups) {
			continue
		}
		ok := true
		for i := range sGroups {
			if sGroups[i].char != wGroups[i].char {
				ok = false
				break
			}
			if sGroups[i].cnt < 3 && sGroups[i].cnt != wGroups[i].cnt {
				ok = false
				break
			}
			if sGroups[i].cnt >= 3 && wGroups[i].cnt > sGroups[i].cnt {
				ok = false
				break
			}
		}
		if ok {
			count++
		}
	}

	return count
}
```

## 0811 — Subdomain Visit Count

```go
package main

// LeetCode #811: Subdomain Visit Count
// https://leetcode.com/problems/subdomain-visit-count/
// Difficulty: Medium

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(SubdomainVisitCount([]string{"9001 discuss.leetcode.com"}))
	fmt.Println(SubdomainVisitCount([]string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}))
}

func SubdomainVisitCount(cpdomains []string) []string {
	counts := make(map[string]int)

	for _, cpdomain := range cpdomains {
		parts := strings.SplitN(cpdomain, " ", 2)
		count, _ := strconv.Atoi(parts[0])
		domain := parts[1]

		subdomains := strings.Split(domain, ".")
		for i := range subdomains {
			sub := strings.Join(subdomains[i:], ".")
			counts[sub] += count
		}
	}

	result := make([]string, 0, len(counts))
	for sub, cnt := range counts {
		result = append(result, strconv.Itoa(cnt)+" "+sub)
	}

	return result
}
```

## 0813 — Largest Sum Of Averages

```go
package main

// LeetCode #813: Largest Sum of Averages
// https://leetcode.com/problems/largest-sum-of-averages/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(LargestSumOfAverages([]int{9, 1, 2, 3, 9}, 3))
	fmt.Println(LargestSumOfAverages([]int{1, 2, 3, 4, 5, 6, 7}, 4))
	fmt.Println(LargestSumOfAverages([]int{4, 1, 7, 5, 6, 2, 3}, 4))
}

// Time: O(k * n^2) | Space: O(k * n)
func LargestSumOfAverages(nums []int, k int) float64 {
	n := len(nums)
	prefix := make([]float64, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + float64(nums[i])
	}

	dp := make([][]float64, n+1)
	for i := range dp {
		dp[i] = make([]float64, k+1)
	}

	for i := 1; i <= n; i++ {
		dp[i][1] = prefix[i] / float64(i)
	}

	for j := 2; j <= k; j++ {
		for i := j; i <= n; i++ {
			var best float64
			for x := j - 1; x < i; x++ {
				avg := (prefix[i] - prefix[x]) / float64(i-x)
				val := dp[x][j-1] + avg
				if val > best {
					best = val
				}
			}
			dp[i][j] = best
		}
	}

	return dp[n][k]
}
```

## 0814 — Binary Tree Pruning

```go
package main

// LeetCode #814: Binary Tree Pruning
// https://leetcode.com/problems/binary-tree-pruning/
// Difficulty: Medium

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Test case 1: [1,null,0,0,1] -> [1,null,0,null,1]
	root1 := &TreeNode{1, nil, &TreeNode{0, &TreeNode{0, nil, nil}, &TreeNode{1, nil, nil}}}
	r1 := BinaryTreePruning(root1)
	printTree(r1)
	fmt.Println()

	// Test case 2: [1,0,1,0,0,0,1] -> [1,null,1,null,1]
	root2 := &TreeNode{1,
		&TreeNode{0, &TreeNode{0, nil, nil}, &TreeNode{0, nil, nil}},
		&TreeNode{1, &TreeNode{0, nil, nil}, &TreeNode{1, nil, nil}},
	}
	r2 := BinaryTreePruning(root2)
	printTree(r2)
	fmt.Println()

	// Test case 3: nil tree
	r3 := BinaryTreePruning(nil)
	printTree(r3)
	fmt.Println()
}

func printTree(root *TreeNode) {
	if root == nil {
		fmt.Print("null")
		return
	}
	fmt.Print(root.Val)
	if root.Left != nil || root.Right != nil {
		fmt.Print(" ")
		printTree(root.Left)
		fmt.Print(" ")
		printTree(root.Right)
	}
}

// Time: O(n) | Space: O(h) where h is tree height
func BinaryTreePruning(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left = BinaryTreePruning(root.Left)
	root.Right = BinaryTreePruning(root.Right)

	if root.Left == nil && root.Right == nil && root.Val == 0 {
		return nil
	}

	return root
}
```

## 0816 — Ambiguous Coordinates

```go
package main

// LeetCode #816: Ambiguous Coordinates
// https://leetcode.com/problems/ambiguous-coordinates/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(AmbiguousCoordinates("(123)"))
	fmt.Println(AmbiguousCoordinates("(00011)"))
	fmt.Println(AmbiguousCoordinates("(0123)"))
}

// Time: O(n^3) | Space: O(n^2)
func AmbiguousCoordinates(s string) []string {
	s = s[1 : len(s)-1]
	n := len(s)

	var result []string

	for i := 1; i < n; i++ {
		left := validNums(s[:i])
		right := validNums(s[i:])
		for _, a := range left {
			for _, b := range right {
				result = append(result, "("+a+", "+b+")")
			}
		}
	}

	return result
}

func validNums(s string) []string {
	var res []string

	if len(s) == 0 {
		return res
	}

	if s[0] == '0' && s[len(s)-1] == '0' {
		if len(s) == 1 {
			res = append(res, s)
		}
		return res
	}

	if s[0] == '0' {
		res = append(res, "0."+s[1:])
		return res
	}

	res = append(res, s)
	if s[len(s)-1] == '0' {
		return res
	}

	for i := 1; i < len(s); i++ {
		res = append(res, s[:i]+"."+s[i:])
	}

	return res
}
```

## 0817 — Linked List Components

```go
package main

// LeetCode #817: Linked List Components
// https://leetcode.com/problems/linked-list-components/
// Difficulty: Medium

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// Test case 1: [0,1,2,3], nums=[0,1,3] -> 2
	head1 := &ListNode{0, &ListNode{1, &ListNode{2, &ListNode{3, nil}}}}
	fmt.Println(LinkedListComponents(head1, []int{0, 1, 3}))

	// Test case 2: [0,1,2,3,4], nums=[0,3,1,4] -> 2
	head2 := &ListNode{0, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}}
	fmt.Println(LinkedListComponents(head2, []int{0, 3, 1, 4}))

	// Test case 3: nil list
	fmt.Println(LinkedListComponents(nil, []int{1}))
}

// Time: O(n) | Space: O(m) where m = len(nums)
func LinkedListComponents(head *ListNode, nums []int) int {
	set := make(map[int]bool)
	for _, v := range nums {
		set[v] = true
	}

	count := 0
	inComponent := false
	for cur := head; cur != nil; cur = cur.Next {
		if set[cur.Val] {
			if !inComponent {
				count++
				inComponent = true
			}
		} else {
			inComponent = false
		}
	}

	return count
}
```

## 0820 — Short Encoding Of Words

```go
package main

// LeetCode #820: Short Encoding of Words
// https://leetcode.com/problems/short-encoding-of-words/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(ShortEncodingOfWords([]string{"time", "me", "bell"}))
	fmt.Println(ShortEncodingOfWords([]string{"t"}))
	fmt.Println(ShortEncodingOfWords([]string{"me", "time"}))
}

// Time: O(n * L^2) | Space: O(n * L) where L is average word length
func ShortEncodingOfWords(words []string) int {
	set := make(map[string]bool)
	for _, word := range words {
		set[word] = true
	}

	for _, word := range words {
		for i := 1; i < len(word); i++ {
			delete(set, word[i:])
		}
	}

	ans := 0
	for word := range set {
		ans += len(word) + 1
	}
	return ans
}
```

## 0822 — Card Flipping Game

```go
package main

// LeetCode #822: Card Flipping Game
// https://leetcode.com/problems/card-flipping-game/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(CardFlippingGame([]int{1, 2, 4, 4, 7}, []int{1, 3, 4, 1, 3}))
	fmt.Println(CardFlippingGame([]int{1, 1}, []int{1, 2}))
	fmt.Println(CardFlippingGame([]int{1, 1}, []int{2, 2}))
}

// Time: O(n) | Space: O(n)
func CardFlippingGame(fronts []int, backs []int) int {
	blocked := make(map[int]bool)
	for i := range fronts {
		if fronts[i] == backs[i] {
			blocked[fronts[i]] = true
		}
	}

	ans := 2001
	for _, v := range fronts {
		if !blocked[v] && v < ans {
			ans = v
		}
	}
	for _, v := range backs {
		if !blocked[v] && v < ans {
			ans = v
		}
	}

	if ans == 2001 {
		return 0
	}
	return ans
}
```

## 0823 — Binary Trees With Factors

```go
package main

// LeetCode #823: Binary Trees With Factors
// https://leetcode.com/problems/binary-trees-with-factors/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(BinaryTreesWithFactors([]int{2, 4}))
	fmt.Println(BinaryTreesWithFactors([]int{2, 4, 5, 10}))
	fmt.Println(BinaryTreesWithFactors([]int{2, 3, 4, 6, 8, 12, 24}))
}

// Time: O(n^2) | Space: O(n)
func BinaryTreesWithFactors(arr []int) int {
	const mod = 1_000_000_007
	sort.Ints(arr)

	dp := make(map[int]int)
	for _, x := range arr {
		dp[x] = 1
	}

	for i, x := range arr {
		for j := 0; j < i; j++ {
			if x%arr[j] == 0 {
				if val, ok := dp[x/arr[j]]; ok {
					dp[x] = (dp[x] + dp[arr[j]]*val) % mod
				}
			}
		}
	}

	ans := 0
	for _, v := range dp {
		ans = (ans + v) % mod
	}
	return ans
}
```

## 0825 — Friends Of Appropriate Ages

```go
package main

// LeetCode #825: Friends Of Appropriate Ages
// https://leetcode.com/problems/friends-of-appropriate-ages/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(FriendsOfAppropriateAges([]int{16, 16}))
	fmt.Println(FriendsOfAppropriateAges([]int{16, 17, 18}))
	fmt.Println(FriendsOfAppropriateAges([]int{20, 30, 100, 110, 120}))
}

// Time: O(n + R) where R = 120 | Space: O(R)
func FriendsOfAppropriateAges(ages []int) int {
	cnt := make([]int, 121)
	for _, age := range ages {
		cnt[age]++
	}

	prefix := make([]int, 121)
	for i := 1; i <= 120; i++ {
		prefix[i] = prefix[i-1] + cnt[i]
	}

	ans := 0
	for age := 1; age <= 120; age++ {
		if cnt[age] == 0 {
			continue
		}
		left := age/2 + 8
		if left > age {
			continue
		}
		total := prefix[age] - prefix[left-1] - 1 // exclude self
		ans += cnt[age] * total
	}

	return ans
}
```

## 0826 — Most Profit Assigning Work

```go
package main

// LeetCode #826: Most Profit Assigning Work
// https://leetcode.com/problems/most-profit-assigning-work/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MostProfitAssigningWork([]int{2, 4, 6, 8, 10}, []int{10, 20, 30, 40, 50}, []int{4, 5, 6, 7}))
	fmt.Println(MostProfitAssigningWork([]int{85, 47, 57}, []int{24, 66, 99}, []int{40, 25, 25}))
	fmt.Println(MostProfitAssigningWork([]int{68, 35, 52, 47, 86}, []int{67, 17, 1, 81, 3}, []int{92, 10, 85, 84, 82}))
}

// Time: O(n log n + m log m) | Space: O(n)
func MostProfitAssigningWork(difficulty []int, profit []int, worker []int) int {
	n := len(difficulty)
	jobs := make([][2]int, n)
	for i := range difficulty {
		jobs[i] = [2]int{difficulty[i], profit[i]}
	}

	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i][0] < jobs[j][0]
	})
	sort.Ints(worker)

	ans := 0
	idx := 0
	maxProfit := 0

	for _, w := range worker {
		for idx < n && jobs[idx][0] <= w {
			if jobs[idx][1] > maxProfit {
				maxProfit = jobs[idx][1]
			}
			idx++
		}
		ans += maxProfit
	}

	return ans
}
```

## 0831 — Masking Personal Information

```go
package main

// LeetCode #831: Masking Personal Information
// https://leetcode.com/problems/masking-personal-information/
// Difficulty: Medium

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(MaskingPersonalInformation("LeetCode@LeetCode.com"))
	fmt.Println(MaskingPersonalInformation("AB@qq.com"))
	fmt.Println(MaskingPersonalInformation("1(234)567-890"))
	fmt.Println(MaskingPersonalInformation("86-(10)12345678"))
}

// Time: O(n) | Space: O(n)
func MaskingPersonalInformation(s string) string {
	at := strings.IndexByte(s, '@')
	if at != -1 {
		s = strings.ToLower(s)
		return string(s[0]) + "*****" + string(s[at-1]) + s[at:]
	}

	var digits strings.Builder
	for _, c := range s {
		if unicode.IsDigit(c) {
			digits.WriteRune(c)
		}
	}
	d := digits.String()
	local := d[len(d)-4:]
	masked := "***-***-" + local

	if len(d) == 10 {
		return masked
	}
	return "+" + strings.Repeat("*", len(d)-10) + "-" + masked
}
```

## 0833 — Find And Replace In String

```go
package main

// LeetCode #833: Find And Replace in String
// https://leetcode.com/problems/find-and-replace-in-string/
// Difficulty: Medium

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(FindAndReplaceInString("abcd", []int{0, 2}, []string{"a", "cd"}, []string{"eee", "ffff"}))
	fmt.Println(FindAndReplaceInString("abcd", []int{0, 2}, []string{"ab", "ec"}, []string{"eee", "ffff"}))
	fmt.Println(FindAndReplaceInString("jjievdtjfb", []int{4, 6}, []string{"md", "tjfb"}, []string{"foe", "oov"}))
}

// Time: O(n + m) where n = len(s), m = len(indices) | Space: O(n + m)
func FindAndReplaceInString(s string, indices []int, sources []string, targets []string) string {
	n := len(s)
	replace := make([]int, n)
	for i := range replace {
		replace[i] = -1
	}

	for k, idx := range indices {
		if idx+len(sources[k]) <= n && s[idx:idx+len(sources[k])] == sources[k] {
			replace[idx] = k
		}
	}

	var sb strings.Builder
	for i := 0; i < n; {
		if replace[i] >= 0 {
			sb.WriteString(targets[replace[i]])
			i += len(sources[replace[i]])
		} else {
			sb.WriteByte(s[i])
			i++
		}
	}

	return sb.String()
}
```

## 0835 — Image Overlap

```go
package main

// LeetCode #835: Image Overlap
// https://leetcode.com/problems/image-overlap/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(ImageOverlap([][]int{{1, 1, 0}, {0, 1, 0}, {0, 1, 0}}, [][]int{{0, 0, 0}, {0, 1, 1}, {0, 0, 1}}))
	fmt.Println(ImageOverlap([][]int{{1}}, [][]int{{1}}))
	fmt.Println(ImageOverlap([][]int{{0}}, [][]int{{0}}))
}

// Time: O(n^4) | Space: O(1)
func ImageOverlap(img1 [][]int, img2 [][]int) int {
	n := len(img1)
	ans := 0

	for a := 1 - n; a < n; a++ {
		for b := 1 - n; b < n; b++ {
			count := 0
			for i := max(a, 0); i < min(n, n+a); i++ {
				for j := max(b, 0); j < min(n, n+b); j++ {
					if img2[i][j] == 1 && img1[i-a][j-b] == 1 {
						count++
					}
				}
			}
			if count > ans {
				ans = count
			}
		}
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

## 0837 — New 21 Game

```go
package main

// LeetCode #837: New 21 Game
// https://leetcode.com/problems/new-21-game/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(NewTwoOneGame(10, 1, 10))
	fmt.Println(NewTwoOneGame(6, 1, 10))
	fmt.Println(NewTwoOneGame(21, 17, 10))
}

// Time: O(n) | Space: O(n)
func NewTwoOneGame(n int, k int, maxPts int) float64 {
	if k == 0 || n >= k-1+maxPts {
		return 1.0
	}

	dp := make([]float64, n+1)
	dp[0] = 1.0
	windowSum := 1.0
	var ans float64

	for i := 1; i <= n; i++ {
		dp[i] = windowSum / float64(maxPts)
		if i < k {
			windowSum += dp[i]
		} else {
			ans += dp[i]
		}
		if i >= maxPts {
			windowSum -= dp[i-maxPts]
		}
	}

	return ans
}
```

## 0838 — Push Dominoes

```go
package main

// LeetCode #838: Push Dominoes
// https://leetcode.com/problems/push-dominoes/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(PushDominoes("RR.L"))
	fmt.Println(PushDominoes(".L.R...LR..L.."))
	fmt.Println(PushDominoes("L.R"))
}

// Time: O(n) | Space: O(n)
func PushDominoes(dominoes string) string {
	n := len(dominoes)
	res := []byte(dominoes)

	for i := 0; i < n; i++ {
		if res[i] == 'R' {
			// Find the next non-dot character
			j := i + 1
			for j < n && res[j] == '.' {
				j++
			}
			if j == n || res[j] == 'R' {
				// All dots between i and j become 'R'
				for k := i + 1; k < j; k++ {
					res[k] = 'R'
				}
			} else if res[j] == 'L' {
				// Collision: left and right meet in the middle
				left, right := i+1, j-1
				for left < right {
					res[left] = 'R'
					res[right] = 'L'
					left++
					right--
				}
			}
			i = j
		} else if res[i] == 'L' {
			// Propagate L leftwards
			j := i - 1
			for j >= 0 && res[j] == '.' {
				res[j] = 'L'
				j--
			}
		}
	}

	// Handle initial dots before first 'L'
	for i := 0; i < n && res[i] == '.'; i++ {
		if i+1 < n && res[i+1] == 'L' {
			res[i] = 'L'
		}
	}

	return string(res)
}
```

## 0840 — Magic Squares In Grid

```go
package main

// LeetCode #840: Magic Squares In Grid
// https://leetcode.com/problems/magic-squares-in-grid/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MagicSquaresInGrid([][]int{{4, 3, 8, 4}, {9, 5, 1, 9}, {2, 7, 6, 2}}))
	fmt.Println(MagicSquaresInGrid([][]int{{8}}))
	fmt.Println(MagicSquaresInGrid([][]int{{5, 5, 5}, {5, 5, 5}, {5, 5, 5}}))
}

// Time: O(m * n) | Space: O(1)
func MagicSquaresInGrid(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	if m < 3 || n < 3 {
		return 0
	}

	ans := 0
	for i := 0; i <= m-3; i++ {
		for j := 0; j <= n-3; j++ {
			if isMagic(grid, i, j) {
				ans++
			}
		}
	}
	return ans
}

func isMagic(grid [][]int, r, c int) bool {
	seen := [16]bool{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			v := grid[r+i][c+j]
			if v < 1 || v > 9 || seen[v] {
				return false
			}
			seen[v] = true
		}
	}

	sum := grid[r][c] + grid[r][c+1] + grid[r][c+2]
	for i := 0; i < 3; i++ {
		if grid[r+i][c]+grid[r+i][c+1]+grid[r+i][c+2] != sum {
			return false
		}
		if grid[r][c+i]+grid[r+1][c+i]+grid[r+2][c+i] != sum {
			return false
		}
	}
	if grid[r][c]+grid[r+1][c+1]+grid[r+2][c+2] != sum {
		return false
	}
	if grid[r][c+2]+grid[r+1][c+1]+grid[r+2][c] != sum {
		return false
	}

	return true
}
```

## 0841 — Keys And Rooms

```go
package main

// LeetCode #841: Keys and Rooms
// https://leetcode.com/problems/keys-and-rooms/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(KeysAndRooms([][]int{{1}, {2}, {3}, {}}))
	fmt.Println(KeysAndRooms([][]int{{1, 3}, {3, 0, 1}, {2}, {0}}))
	fmt.Println(KeysAndRooms([][]int{{2}, {}, {1}}))
}

// Time: O(n + k) where k = total keys | Space: O(n)
func KeysAndRooms(rooms [][]int) bool {
	n := len(rooms)
	visited := make([]bool, n)
	stack := []int{0}
	visited[0] = true
	count := 1

	for len(stack) > 0 {
		room := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for _, key := range rooms[room] {
			if !visited[key] {
				visited[key] = true
				count++
				stack = append(stack, key)
			}
		}
	}

	return count == n
}
```

## 0842 — Split Array Into Fibonacci Sequence

```go
package main

// LeetCode #842: Split Array into Fibonacci Sequence
// https://leetcode.com/problems/split-array-into-fibonacci-sequence/
// Difficulty: Medium

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(SplitArrayIntoFibonacciSequence("123456579"))
	fmt.Println(SplitArrayIntoFibonacciSequence("11235813"))
	fmt.Println(SplitArrayIntoFibonacciSequence("112358130"))
}

// Time: O(n^2) | Space: O(n)
func SplitArrayIntoFibonacciSequence(num string) []int {
	n := len(num)
	var ans []int

	var dfs func(int) bool
	dfs = func(pos int) bool {
		if pos == n {
			return len(ans) > 2
		}

		var x int
		for i := pos; i < n; i++ {
			if i > pos && num[pos] == '0' {
				break
			}
			x = x*10 + int(num[i]-'0')
			if x > math.MaxInt32 {
				break
			}
			if len(ans) > 1 && x > ans[len(ans)-1]+ans[len(ans)-2] {
				break
			}
			if len(ans) < 2 || x == ans[len(ans)-1]+ans[len(ans)-2] {
				ans = append(ans, x)
				if dfs(i + 1) {
					return true
				}
				ans = ans[:len(ans)-1]
			}
		}
		return false
	}

	dfs(0)
	return ans
}
```

## 0845 — Longest Mountain In Array

```go
package main

// LeetCode #845: Longest Mountain in Array
// https://leetcode.com/problems/longest-mountain-in-array/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(LongestMountainInArray([]int{2, 1, 4, 7, 3, 2, 5}))
	fmt.Println(LongestMountainInArray([]int{2, 2, 2}))
	fmt.Println(LongestMountainInArray([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}))
}

// Time: O(n) | Space: O(1)
func LongestMountainInArray(arr []int) int {
	n := len(arr)
	ans := 0
	i := 1

	for i < n {
		// Skip non-increasing start
		for i < n && arr[i] == arr[i-1] {
			i++
		}

		// Climb up
		up := 0
		for i < n && arr[i] > arr[i-1] {
			up++
			i++
		}

		// Climb down
		down := 0
		for i < n && arr[i] < arr[i-1] {
			down++
			i++
		}

		// Valid mountain needs both up and down segments
		if up > 0 && down > 0 {
			if up+down+1 > ans {
				ans = up + down + 1
			}
		}
	}

	return ans
}
```

## 0846 — Hand Of Straights

```go
package main

// LeetCode #846: Hand of Straights
// https://leetcode.com/problems/hand-of-straights/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(HandOfStraights([]int{1, 2, 3, 6, 2, 3, 4, 7, 8}, 3))
	fmt.Println(HandOfStraights([]int{1, 2, 3, 4, 5}, 4))
	fmt.Println(HandOfStraights([]int{2, 1}, 2))
}

// Time: O(n log n) | Space: O(n)
func HandOfStraights(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}

	count := make(map[int]int)
	for _, card := range hand {
		count[card]++
	}

	unique := make([]int, 0, len(count))
	for card := range count {
		unique = append(unique, card)
	}
	sort.Ints(unique)

	for _, card := range unique {
		if count[card] > 0 {
			freq := count[card]
			for i := 0; i < groupSize; i++ {
				if count[card+i] < freq {
					return false
				}
				count[card+i] -= freq
			}
		}
	}

	return true
}
```

## 0848 — Shifting Letters

```go
package main

// LeetCode #848: Shifting Letters
// https://leetcode.com/problems/shifting-letters/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(ShiftingLetters("abc", []int{3, 5, 9}))
	fmt.Println(ShiftingLetters("aaa", []int{1, 2, 3}))
	fmt.Println(ShiftingLetters("z", []int{52}))
}

// Time: O(n) | Space: O(n)
func ShiftingLetters(s string, shifts []int) string {
	n := len(s)
	// Calculate suffix sum of shifts
	for i := n - 2; i >= 0; i-- {
		shifts[i] = (shifts[i] + shifts[i+1]) % 26
	}

	res := []byte(s)
	for i := 0; i < n; i++ {
		res[i] = byte((int(res[i]-'a')+shifts[i])%26 + 'a')
	}

	return string(res)
}
```

## 0849 — Maximize Distance To Closest Person

```go
package main

// LeetCode #849: Maximize Distance to Closest Person
// https://leetcode.com/problems/maximize-distance-to-closest-person/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MaximizeDistanceToClosestPerson([]int{1, 0, 0, 0, 1, 0, 1}))
	fmt.Println(MaximizeDistanceToClosestPerson([]int{1, 0, 0, 0}))
	fmt.Println(MaximizeDistanceToClosestPerson([]int{0, 1}))
}

// Time: O(n) | Space: O(1)
func MaximizeDistanceToClosestPerson(seats []int) int {
	n := len(seats)
	ans := 0
	lastPerson := -1

	for i := 0; i < n; i++ {
		if seats[i] == 1 {
			if lastPerson == -1 {
				ans = i
			} else {
				dist := (i - lastPerson) / 2
				if dist > ans {
					ans = dist
				}
			}
			lastPerson = i
		}
	}

	// Check distance from last person to the end
	if seats[n-1] == 0 {
		dist := n - 1 - lastPerson
		if dist > ans {
			ans = dist
		}
	}

	return ans
}
```

## 0851 — Loud And Rich

```go
package main

// LeetCode #851: Loud and Rich
// https://leetcode.com/problems/loud-and-rich/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(LoudAndRich([][]int{{1, 0}, {2, 1}, {3, 1}, {3, 7}, {4, 3}, {5, 3}, {6, 3}}, []int{3, 2, 5, 4, 6, 1, 7, 0}))
	fmt.Println(LoudAndRich([][]int{{0, 1}, {1, 2}}, []int{0, 1, 2}))
	fmt.Println(LoudAndRich([][]int{}, []int{0}))
}

// Time: O(n + m) | Space: O(n + m) where m = len(richer)
func LoudAndRich(richer [][]int, quiet []int) []int {
	n := len(quiet)
	graph := make([][]int, n)
	indeg := make([]int, n)

	for _, r := range richer {
		a, b := r[0], r[1]
		graph[a] = append(graph[a], b)
		indeg[b]++
	}

	ans := make([]int, n)
	for i := range ans {
		ans[i] = i
	}

	var queue []int
	for i := 0; i < n; i++ {
		if indeg[i] == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		for _, v := range graph[u] {
			if quiet[ans[v]] > quiet[ans[u]] {
				ans[v] = ans[u]
			}
			indeg[v]--
			if indeg[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	return ans
}
```

## 0852 — Peak Index In A Mountain Array

```go
package main

// LeetCode #852: Peak Index in a Mountain Array
// https://leetcode.com/problems/peak-index-in-a-mountain-array/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(PeakIndexInAMountainArray([]int{0, 1, 0}))
	fmt.Println(PeakIndexInAMountainArray([]int{0, 2, 1, 0}))
	fmt.Println(PeakIndexInAMountainArray([]int{0, 10, 5, 2}))
}

// Time: O(log n) | Space: O(1)
func PeakIndexInAMountainArray(arr []int) int {
	left, right := 1, len(arr)-2
	for left < right {
		mid := (left + right) / 2
		if arr[mid] > arr[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
```

## 0853 — Car Fleet

```go
package main

// LeetCode #853: Car Fleet
// https://leetcode.com/problems/car-fleet/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(CarFleet(12, []int{10, 8, 0, 5, 3}, []int{2, 4, 1, 1, 3}))
	fmt.Println(CarFleet(10, []int{3}, []int{3}))
	fmt.Println(CarFleet(100, []int{0, 2, 4}, []int{4, 2, 1}))
}

// Time: O(n log n) | Space: O(n)
func CarFleet(target int, position []int, speed []int) int {
	n := len(position)
	if n == 0 {
		return 0
	}

	type car struct {
		pos  int
		time float64
	}
	cars := make([]car, n)
	for i := range position {
		cars[i] = car{position[i], float64(target-position[i]) / float64(speed[i])}
	}

	sort.Slice(cars, func(i, j int) bool {
		return cars[i].pos > cars[j].pos
	})

	fleets := 1
	maxTime := cars[0].time
	for i := 1; i < n; i++ {
		if cars[i].time > maxTime {
			fleets++
			maxTime = cars[i].time
		}
	}

	return fleets
}
```

## 0855 — Exam Room

```go
package main

// LeetCode #855: Exam Room
// https://leetcode.com/problems/exam-room/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

type ExamRoom struct {
	seats []int
	n     int
}

func Constructor(n int) ExamRoom {
	return ExamRoom{seats: []int{}, n: n}
}

func (this *ExamRoom) Seat() int {
	if len(this.seats) == 0 {
		this.seats = append(this.seats, 0)
		return 0
	}

	maxDist := this.seats[0]
	pos := 0

	for i := 0; i < len(this.seats)-1; i++ {
		dist := (this.seats[i+1] - this.seats[i]) / 2
		if dist > maxDist {
			maxDist = dist
			pos = this.seats[i] + dist
		}
	}

	dist := this.n - 1 - this.seats[len(this.seats)-1]
	if dist > maxDist {
		pos = this.n - 1
	}

	idx := sort.SearchInts(this.seats, pos)
	this.seats = append(this.seats, 0)
	copy(this.seats[idx+1:], this.seats[idx:])
	this.seats[idx] = pos

	return pos
}

func (this *ExamRoom) Leave(p int) {
	idx := sort.SearchInts(this.seats, p)
	this.seats = append(this.seats[:idx], this.seats[idx+1:]...)
}

func main() {
	// Test case 1
	obj1 := Constructor(10)
	fmt.Println(obj1.Seat())
	fmt.Println(obj1.Seat())
	fmt.Println(obj1.Seat())
	fmt.Println(obj1.Seat())
	obj1.Leave(4)
	fmt.Println(obj1.Seat())

	fmt.Println("---")

	// Test case 2
	obj2 := Constructor(5)
	fmt.Println(obj2.Seat())
	fmt.Println(obj2.Seat())
	fmt.Println(obj2.Seat())
	fmt.Println(obj2.Seat())
	fmt.Println(obj2.Seat())
}
```

## 0856 — Score Of Parentheses

```go
package main

// LeetCode #856: Score of Parentheses
// https://leetcode.com/problems/score-of-parentheses/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(ScoreOfParentheses("()"))
	fmt.Println(ScoreOfParentheses("(())"))
	fmt.Println(ScoreOfParentheses("()()"))
	fmt.Println(ScoreOfParentheses("(()(()))"))
}

// Time: O(n) | Space: O(n)
func ScoreOfParentheses(s string) int {
	stack := []int{0}
	for _, ch := range s {
		if ch == '(' {
			stack = append(stack, 0)
		} else {
			x := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if x != 0 {
				x *= 2
			} else {
				x = 1
			}
			stack[len(stack)-1] += x
		}
	}
	return stack[0]
}
```

## 0858 — Mirror Reflection

```go
package main

// LeetCode #858: Mirror Reflection
// https://leetcode.com/problems/mirror-reflection/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MirrorReflection(2, 1))
	fmt.Println(MirrorReflection(3, 1))
	fmt.Println(MirrorReflection(4, 3))
}

// Time: O(log min(p,q)) | Space: O(1)
func MirrorReflection(p int, q int) int {
	g := gcd(p, q)
	p /= g
	q /= g

	if p%2 == 0 {
		return 2
	}
	if q%2 == 0 {
		return 0
	}
	return 1
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
```

## 0861 — Score After Flipping Matrix

```go
package main

// LeetCode #861: Score After Flipping Matrix
// https://leetcode.com/problems/score-after-flipping-matrix/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(ScoreAfterFlippingMatrix([][]int{{0, 0, 1, 1}, {1, 0, 1, 0}, {1, 1, 0, 0}}))
	fmt.Println(ScoreAfterFlippingMatrix([][]int{{0}}))
	fmt.Println(ScoreAfterFlippingMatrix([][]int{{1, 1}, {1, 1}}))
}

// Time: O(m * n) | Space: O(1)
func ScoreAfterFlippingMatrix(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	// Ensure first column is all 1s
	for i := 0; i < m; i++ {
		if grid[i][0] == 0 {
			for j := 0; j < n; j++ {
				grid[i][j] ^= 1
			}
		}
	}

	ans := 0
	for j := 0; j < n; j++ {
		ones := 0
		for i := 0; i < m; i++ {
			ones += grid[i][j]
		}
		if ones < m-ones {
			ones = m - ones
		}
		ans += ones * (1 << (n - 1 - j))
	}

	return ans
}
```

## 0863 — All Nodes Distance K In Binary Tree

```go
package main

// LeetCode #863: All Nodes Distance K in Binary Tree
// https://leetcode.com/problems/all-nodes-distance-k-in-binary-tree/
// Difficulty: Medium

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Test case 1: [3,5,1,6,2,0,8,null,null,7,4], target=5, k=2 -> [7,4,1]
	root := &TreeNode{3,
		&TreeNode{5, &TreeNode{6, nil, nil}, &TreeNode{2, &TreeNode{7, nil, nil}, &TreeNode{4, nil, nil}}},
		&TreeNode{1, &TreeNode{0, nil, nil}, &TreeNode{8, nil, nil}},
	}
	fmt.Println(AllNodesDistanceKInBinaryTree(root, root.Left, 2))

	// Test case 2: target=5, k=1 -> [6,2,3]
	fmt.Println(AllNodesDistanceKInBinaryTree(root, root.Left, 1))

	// Test case 3: k=0
	fmt.Println(AllNodesDistanceKInBinaryTree(root, root.Left, 0))
}

// Time: O(n) | Space: O(n)
func AllNodesDistanceKInBinaryTree(root *TreeNode, target *TreeNode, k int) []int {
	parent := make(map[*TreeNode]*TreeNode)
	var buildParent func(*TreeNode, *TreeNode)
	buildParent = func(node, par *TreeNode) {
		if node == nil {
			return
		}
		parent[node] = par
		buildParent(node.Left, node)
		buildParent(node.Right, node)
	}
	buildParent(root, nil)

	visited := make(map[*TreeNode]bool)
	var ans []int
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, dist int) {
		if node == nil || visited[node] {
			return
		}
		visited[node] = true
		if dist == k {
			ans = append(ans, node.Val)
			return
		}
		dfs(node.Left, dist+1)
		dfs(node.Right, dist+1)
		dfs(parent[node], dist+1)
	}
	dfs(target, 0)

	return ans
}
```

## 0865 — Smallest Subtree With All The Deepest Nodes

```go
package main

// LeetCode #865: Smallest Subtree with all the Deepest Nodes
// https://leetcode.com/problems/smallest-subtree-with-all-the-deepest-nodes/
// Difficulty: Medium

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Test case 1: [3,5,1,6,2,0,8,null,null,7,4] -> [2,7,4]
	root := &TreeNode{3,
		&TreeNode{5, &TreeNode{6, nil, nil}, &TreeNode{2, &TreeNode{7, nil, nil}, &TreeNode{4, nil, nil}}},
		&TreeNode{1, &TreeNode{0, nil, nil}, &TreeNode{8, nil, nil}},
	}
	r1 := SmallestSubtreeWithAllTheDeepestNodes(root)
	fmt.Println(r1.Val)

	// Test case 2: [1] -> [1]
	root2 := &TreeNode{1, nil, nil}
	r2 := SmallestSubtreeWithAllTheDeepestNodes(root2)
	fmt.Println(r2.Val)

	// Test case 3: [0,1,3,null,2] -> [2]
	root3 := &TreeNode{0,
		&TreeNode{1, nil, &TreeNode{2, nil, nil}},
		&TreeNode{3, nil, nil},
	}
	r3 := SmallestSubtreeWithAllTheDeepestNodes(root3)
	fmt.Println(r3.Val)
}

// Time: O(n) | Space: O(h)
func SmallestSubtreeWithAllTheDeepestNodes(root *TreeNode) *TreeNode {
	var dfs func(*TreeNode) (*TreeNode, int)
	dfs = func(node *TreeNode) (*TreeNode, int) {
		if node == nil {
			return nil, 0
		}
		leftNode, leftDepth := dfs(node.Left)
		rightNode, rightDepth := dfs(node.Right)

		if leftDepth > rightDepth {
			return leftNode, leftDepth + 1
		} else if rightDepth > leftDepth {
			return rightNode, rightDepth + 1
		}
		return node, leftDepth + 1
	}

	node, _ := dfs(root)
	return node
}
```

## 0866 — Prime Palindrome

```go
package main

// LeetCode #866: Prime Palindrome
// https://leetcode.com/problems/prime-palindrome/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(PrimePalindrome(6))
	fmt.Println(PrimePalindrome(8))
	fmt.Println(PrimePalindrome(13))
}

// Time: O(N * sqrt(N)) | Space: O(1)
func PrimePalindrome(N int) int {
	if N <= 2 {
		return 2
	}

	for {
		// Even-length palindromes > 11 are divisible by 11, skip them
		if N >= 10000000 && N < 100000000 {
			N = 100000000
		}

		if isPalindrome(N) && isPrime(N) {
			return N
		}
		N++
	}
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	orig, rev := x, 0
	for x > 0 {
		rev = rev*10 + x%10
		x /= 10
	}
	return orig == rev
}

func isPrime(x int) bool {
	if x < 2 {
		return false
	}
	if x == 2 {
		return true
	}
	if x%2 == 0 {
		return false
	}
	for i := 3; i*i <= x; i += 2 {
		if x%i == 0 {
			return false
		}
	}
	return true
}
```

## 0869 — Reordered Power Of 2

```go
package main

// LeetCode #869: Reordered Power of 2
// https://leetcode.com/problems/reordered-power-of-2/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(ReorderedPowerOfTwo(1))
	fmt.Println(ReorderedPowerOfTwo(10))
	fmt.Println(ReorderedPowerOfTwo(46))
}

// Time: O(log n) | Space: O(1)
func ReorderedPowerOfTwo(n int) bool {
	sig := signature(n)
	for i := 1; i <= 1_000_000_000; i <<= 1 {
		if signature(i) == sig {
			return true
		}
	}
	return false
}

func signature(x int) [10]int {
	var cnt [10]int
	for x > 0 {
		cnt[x%10]++
		x /= 10
	}
	return cnt
}
```

## 0870 — Advantage Shuffle

```go
package main

// LeetCode #870: Advantage Shuffle
// https://leetcode.com/problems/advantage-shuffle/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(AdvantageShuffle([]int{2, 7, 11, 15}, []int{1, 10, 4, 11}))
	fmt.Println(AdvantageShuffle([]int{12, 24, 8, 32}, []int{13, 25, 32, 11}))
	fmt.Println(AdvantageShuffle([]int{2, 0, 4, 1, 2}, []int{1, 3, 0, 0, 2}))
}

// Time: O(n log n) | Space: O(n)
func AdvantageShuffle(nums1 []int, nums2 []int) []int {
	n := len(nums1)
	sort.Ints(nums1)

	idx := make([]int, n)
	for i := range idx {
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool {
		return nums2[idx[i]] < nums2[idx[j]]
	})

	ans := make([]int, n)
	left, right := 0, n-1
	for _, x := range nums1 {
		if x > nums2[idx[left]] {
			ans[idx[left]] = x
			left++
		} else {
			ans[idx[right]] = x
			right--
		}
	}

	return ans
}
```

## 0873 — Length Of Longest Fibonacci Subsequence

```go
package main

// LeetCode #873: Length of Longest Fibonacci Subsequence
// https://leetcode.com/problems/length-of-longest-fibonacci-subsequence/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(LengthOfLongestFibonacciSubsequence([]int{1, 2, 3, 4, 5, 6, 7, 8}))
	fmt.Println(LengthOfLongestFibonacciSubsequence([]int{1, 3, 7, 11, 12, 14, 18}))
	fmt.Println(LengthOfLongestFibonacciSubsequence([]int{1, 3, 5}))
}

// Time: O(n^2) | Space: O(n^2)
func LengthOfLongestFibonacciSubsequence(arr []int) int {
	n := len(arr)
	index := make(map[int]int, n)
	for i, v := range arr {
		index[v] = i
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	ans := 0
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			prev := arr[i] - arr[j]
			if k, ok := index[prev]; ok && k < j {
				dp[i][j] = dp[j][k] + 1
				if dp[i][j] > ans {
					ans = dp[i][j]
				}
			} else {
				dp[i][j] = 2
			}
		}
	}

	if ans >= 3 {
		return ans
	}
	return 0
}
```

## 0874 — Walking Robot Simulation

```go
package main

// LeetCode #874: Walking Robot Simulation
// https://leetcode.com/problems/walking-robot-simulation/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(WalkingRobotSimulation([]int{4, -1, 3}, [][]int{}))
	fmt.Println(WalkingRobotSimulation([]int{4, -1, 4, -2, 4}, [][]int{{2, 4}}))
	fmt.Println(WalkingRobotSimulation([]int{6, -1, -1, 6}, [][]int{}))
}

// Time: O(n + m) where n = len(commands), m = len(obstacles) | Space: O(m)
func WalkingRobotSimulation(commands []int, obstacles [][]int) int {
	obsSet := make(map[[2]int]bool)
	for _, o := range obstacles {
		obsSet[[2]int{o[0], o[1]}] = true
	}

	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	x, y, dir, maxDist := 0, 0, 0, 0

	for _, cmd := range commands {
		if cmd == -1 {
			dir = (dir + 1) % 4
		} else if cmd == -2 {
			dir = (dir + 3) % 4
		} else {
			for step := 0; step < cmd; step++ {
				nx, ny := x+dirs[dir][0], y+dirs[dir][1]
				if obsSet[[2]int{nx, ny}] {
					break
				}
				x, y = nx, ny
				dist := x*x + y*y
				if dist > maxDist {
					maxDist = dist
				}
			}
		}
	}

	return maxDist
}
```

## 0875 — Koko Eating Bananas

```go
package main

// LeetCode #875: Koko Eating Bananas
// https://leetcode.com/problems/koko-eating-bananas/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(KokoEatingBananas([]int{3, 6, 7, 11}, 8))
	fmt.Println(KokoEatingBananas([]int{30, 11, 23, 4, 20}, 5))
	fmt.Println(KokoEatingBananas([]int{30, 11, 23, 4, 20}, 6))
}

// Time: O(n log m) where m = max pile | Space: O(1)
func KokoEatingBananas(piles []int, h int) int {
	maxPile := 0
	for _, p := range piles {
		if p > maxPile {
			maxPile = p
		}
	}

	left, right := 1, maxPile
	for left < right {
		mid := left + (right-left)/2
		hours := 0
		for _, p := range piles {
			hours += (p + mid - 1) / mid
		}
		if hours <= h {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}
```

## 0877 — Stone Game

```go
package main

// LeetCode #877: Stone Game
// https://leetcode.com/problems/stone-game/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(StoneGame([]int{5, 3, 4, 5}))
	fmt.Println(StoneGame([]int{3, 7, 2, 5}))
	fmt.Println(StoneGame([]int{1, 100, 3, 2}))
}

// Time: O(1) | Space: O(1)
// Alex always wins because there are an even number of piles
// and total stones is odd (no ties), with Alex going first.
func StoneGame(piles []int) bool {
	return true
}
```

## 0880 — Decoded String At Index

```go
package main

// LeetCode #880: Decoded String at Index
// https://leetcode.com/problems/decoded-string-at-index/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(DecodedStringAtIndex("leet2code3", 10))
	fmt.Println(DecodedStringAtIndex("ha22", 5))
	fmt.Println(DecodedStringAtIndex("a2345678999999999999999", 1))
}

// Time: O(n) | Space: O(1)
func DecodedStringAtIndex(s string, k int) string {
	var size int64 = 0
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			size++
		} else {
			size *= int64(s[i] - '0')
		}
	}

	target := int64(k)
	for i := len(s) - 1; i >= 0; i-- {
		target %= size
		if target == 0 && s[i] >= 'a' && s[i] <= 'z' {
			return string(s[i])
		}
		if s[i] >= 'a' && s[i] <= 'z' {
			size--
		} else {
			size /= int64(s[i] - '0')
		}
	}

	return ""
}
```

## 0881 — Boats To Save People

```go
package main

// LeetCode #881: Boats to Save People
// https://leetcode.com/problems/boats-to-save-people/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(BoatsToSavePeople([]int{1, 2}, 3))
	fmt.Println(BoatsToSavePeople([]int{3, 2, 2, 1}, 3))
	fmt.Println(BoatsToSavePeople([]int{3, 5, 3, 4}, 5))
}

// Time: O(n log n) | Space: O(log n)
func BoatsToSavePeople(people []int, limit int) int {
	sort.Ints(people)
	left, right := 0, len(people)-1
	ans := 0

	for left <= right {
		if people[left]+people[right] <= limit {
			left++
		}
		right--
		ans++
	}

	return ans
}
```

## 0885 — Spiral Matrix Iii

```go
package main

// LeetCode #885: Spiral Matrix III
// https://leetcode.com/problems/spiral-matrix-iii/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(SpiralMatrixIii(1, 4, 0, 0))
	fmt.Println(SpiralMatrixIii(5, 6, 1, 4))
}

// Time: O(rows * cols) | Space: O(rows * cols)
func SpiralMatrixIii(rows int, cols int, rStart int, cStart int) [][]int {
	total := rows * cols
	ans := make([][]int, 0, total)
	ans = append(ans, []int{rStart, cStart})
	if total == 1 {
		return ans
	}

	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	r, c := rStart, cStart
	step := 1
	dir := 0

	for len(ans) < total {
		for i := 0; i < 2; i++ {
			for j := 0; j < step; j++ {
				r += dirs[dir][0]
				c += dirs[dir][1]
				if r >= 0 && r < rows && c >= 0 && c < cols {
					ans = append(ans, []int{r, c})
				}
			}
			dir = (dir + 1) % 4
		}
		step++
	}

	return ans
}
```

## 0886 — Possible Bipartition

```go
package main

// LeetCode #886: Possible Bipartition
// https://leetcode.com/problems/possible-bipartition/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(PossibleBipartition(4, [][]int{{1, 2}, {1, 3}, {2, 4}}))
	fmt.Println(PossibleBipartition(3, [][]int{{1, 2}, {1, 3}, {2, 3}}))
	fmt.Println(PossibleBipartition(5, [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {1, 5}}))
}

// Time: O(n + d) where d = len(dislikes) | Space: O(n + d)
func PossibleBipartition(n int, dislikes [][]int) bool {
	graph := make([][]int, n+1)
	for _, d := range dislikes {
		a, b := d[0], d[1]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	color := make([]int, n+1) // 0 = uncolored, 1 = group A, -1 = group B

	var dfs func(node, c int) bool
	dfs = func(node, c int) bool {
		if color[node] != 0 {
			return color[node] == c
		}
		color[node] = c
		for _, nei := range graph[node] {
			if !dfs(nei, -c) {
				return false
			}
		}
		return true
	}

	for i := 1; i <= n; i++ {
		if color[i] == 0 && !dfs(i, 1) {
			return false
		}
	}

	return true
}
```

## 0889 — Construct Binary Tree From Preorder And Postorder Traversal

```go
package main

// LeetCode #889: Construct Binary Tree from Preorder and Postorder Traversal
// https://leetcode.com/problems/construct-binary-tree-from-preorder-and-postorder-traversal/
// Difficulty: Medium

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Test case 1: pre = [1,2,4,5,3,6,7], post = [4,5,2,6,7,3,1]
	r1 := ConstructBinaryTreeFromPreorderAndPostorderTraversal([]int{1, 2, 4, 5, 3, 6, 7}, []int{4, 5, 2, 6, 7, 3, 1})
	fmt.Println(preorder(r1))

	// Test case 2: pre = [1], post = [1]
	r2 := ConstructBinaryTreeFromPreorderAndPostorderTraversal([]int{1}, []int{1})
	fmt.Println(preorder(r2))

	// Test case 3: pre = [2,1,3], post = [3,1,2]
	r3 := ConstructBinaryTreeFromPreorderAndPostorderTraversal([]int{2, 1, 3}, []int{3, 1, 2})
	fmt.Println(preorder(r3))
}

func preorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := []int{root.Val}
	res = append(res, preorder(root.Left)...)
	res = append(res, preorder(root.Right)...)
	return res
}

// Time: O(n) | Space: O(n)
func ConstructBinaryTreeFromPreorderAndPostorderTraversal(preorder []int, postorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	if len(preorder) == 1 {
		return root
	}

	// The second element in preorder is the root of the left subtree
	// Find it in postorder to determine left subtree size
	leftSize := 0
	for i := 0; i < len(postorder); i++ {
		if postorder[i] == preorder[1] {
			leftSize = i + 1
			break
		}
	}

	root.Left = ConstructBinaryTreeFromPreorderAndPostorderTraversal(preorder[1:1+leftSize], postorder[:leftSize])
	root.Right = ConstructBinaryTreeFromPreorderAndPostorderTraversal(preorder[1+leftSize:], postorder[leftSize:len(postorder)-1])

	return root
}
```

## 0890 — Find And Replace Pattern

```go
package main

// LeetCode #890: Find and Replace Pattern
// https://leetcode.com/problems/find-and-replace-pattern/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(FindAndReplacePattern([]string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}, "abb"))
	fmt.Println(FindAndReplacePattern([]string{"a", "b", "c"}, "a"))
	fmt.Println(FindAndReplacePattern([]string{"aa", "ab"}, "aa"))
}

// Time: O(n * m) where n = len(words), m = avg word length | Space: O(n)
func FindAndReplacePattern(words []string, pattern string) []string {
	var ans []string
	for _, word := range words {
		if isMatch(word, pattern) {
			ans = append(ans, word)
		}
	}
	return ans
}

func isMatch(word, pattern string) bool {
	if len(word) != len(pattern) {
		return false
	}
	w2p := make(map[byte]byte)
	p2w := make(map[byte]byte)

	for i := 0; i < len(word); i++ {
		wc, pc := word[i], pattern[i]
		if v, ok := w2p[wc]; ok && v != pc {
			return false
		}
		if v, ok := p2w[pc]; ok && v != wc {
			return false
		}
		w2p[wc] = pc
		p2w[pc] = wc
	}

	return true
}
```

## 0893 — Groups Of Special Equivalent Strings

```go
package main

// LeetCode #893: Groups of Special-Equivalent Strings
// https://leetcode.com/problems/groups-of-special-equivalent-strings/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(GroupsOfSpecialEquivalentStrings([]string{"abcd", "cdab", "cbad", "xyzz", "zzxy", "zzyx"}))
	fmt.Println(GroupsOfSpecialEquivalentStrings([]string{"abc", "acb", "bac", "bca", "cab", "cba"}))
	fmt.Println(GroupsOfSpecialEquivalentStrings([]string{"a"}))
}

// Time: O(n * m) where n = len(words), m = avg word length | Space: O(n)
func GroupsOfSpecialEquivalentStrings(words []string) int {
	groups := make(map[[52]int]bool)

	for _, word := range words {
		var key [52]int
		for i, c := range word {
			// Even indices: 0-25, Odd indices: 26-51
			key[int(c-'a')+26*(i%2)]++
		}
		groups[key] = true
	}

	return len(groups)
}
```

## 0894 — All Possible Full Binary Trees

```go
package main

// LeetCode #894: All Possible Full Binary Trees
// https://leetcode.com/problems/all-possible-full-binary-trees/
// Difficulty: Medium

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println(len(AllPossibleFullBinaryTrees(7)))
	fmt.Println(len(AllPossibleFullBinaryTrees(3)))
	fmt.Println(len(AllPossibleFullBinaryTrees(1)))
}

// Time: O(2^n) | Space: O(2^n)
func AllPossibleFullBinaryTrees(n int) []*TreeNode {
	if n%2 == 0 {
		return []*TreeNode{}
	}

	memo := make(map[int][]*TreeNode)

	var dfs func(int) []*TreeNode
	dfs = func(count int) []*TreeNode {
		if trees, ok := memo[count]; ok {
			return trees
		}

		if count == 1 {
			return []*TreeNode{{Val: 0}}
		}

		var res []*TreeNode
		for left := 1; left < count; left += 2 {
			right := count - 1 - left
			for _, l := range dfs(left) {
				for _, r := range dfs(right) {
					res = append(res, &TreeNode{Val: 0, Left: l, Right: r})
				}
			}
		}

		memo[count] = res
		return res
	}

	return dfs(n)
}
```

## 0898 — Bitwise Ors Of Subarrays

```go
package main

// LeetCode #898: Bitwise ORs of Subarrays
// https://leetcode.com/problems/bitwise-ors-of-subarrays/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(BitwiseOrsOfSubarrays([]int{0}))
	fmt.Println(BitwiseOrsOfSubarrays([]int{1, 1, 2}))
	fmt.Println(BitwiseOrsOfSubarrays([]int{1, 2, 4}))
}

// Time: O(n * log(max)) | Space: O(n)
func BitwiseOrsOfSubarrays(arr []int) int {
	set := make(map[int]bool)

	for i := 0; i < len(arr); i++ {
		set[arr[i]] = true
		for j := i - 1; j >= 0; j-- {
			if arr[i]|arr[j] == arr[j] {
				break
			}
			arr[j] |= arr[i]
			set[arr[j]] = true
		}
	}

	return len(set)
}
```

## 0900 — Rle Iterator

```go
package main

// LeetCode #900: RLE Iterator
// https://leetcode.com/problems/rle-iterator/
// Difficulty: Medium

import "fmt"

type RLEIterator struct {
	encoding []int
	idx      int
}

func Constructor(encoding []int) RLEIterator {
	return RLEIterator{encoding: encoding, idx: 0}
}

func (this *RLEIterator) Next(n int) int {
	for this.idx < len(this.encoding) {
		if this.encoding[this.idx] >= n {
			this.encoding[this.idx] -= n
			return this.encoding[this.idx+1]
		}
		n -= this.encoding[this.idx]
		this.idx += 2
	}
	return -1
}

func main() {
	// Test: encoding = [3,8,0,9,2,5]
	obj := Constructor([]int{3, 8, 0, 9, 2, 5})
	fmt.Println(obj.Next(2))  // 8
	fmt.Println(obj.Next(1))  // 8
	fmt.Println(obj.Next(1))  // 5
	fmt.Println(obj.Next(2))  // 5

	fmt.Println("---")

	// Test: encoding = [2,1,3,2]
	obj2 := Constructor([]int{2, 1, 3, 2})
	fmt.Println(obj2.Next(3))  // 2
	fmt.Println(obj2.Next(2))  // -1
}
```

