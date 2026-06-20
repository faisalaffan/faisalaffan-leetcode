# Medium (Sedang) — Problem 0351–0522

## 0351 — Android Unlock Patterns

```go
package main

// LeetCode #351: Android Unlock Patterns
// https://leetcode.com/problems/android-unlock-patterns/
// Difficulty: Medium [Paid]
// Time: O(n!) | Space: O(n)

import "fmt"

func numberOfPatterns(m int, n int) int {
	// skip[i][j] = key that must be visited between i and j (0 if none)
	skip := [10][10]int{}
	skip[1][3] = 2
	skip[3][1] = 2
	skip[1][7] = 4
	skip[7][1] = 4
	skip[3][9] = 6
	skip[9][3] = 6
	skip[7][9] = 8
	skip[9][7] = 8
	skip[1][9] = 5
	skip[9][1] = 5
	skip[2][8] = 5
	skip[8][2] = 5
	skip[3][7] = 5
	skip[7][3] = 5
	skip[4][6] = 5
	skip[6][4] = 5

	visited := [10]bool{}
	var dfs func(cur int, remaining int) int
	dfs = func(cur int, remaining int) int {
		if remaining == 0 {
			return 1
		}
		visited[cur] = true
		count := 0
		for next := 1; next <= 9; next++ {
			if !visited[next] && (skip[cur][next] == 0 || visited[skip[cur][next]]) {
				count += dfs(next, remaining-1)
			}
		}
		visited[cur] = false
		return count
	}

	total := 0
	for length := m; length <= n; length++ {
		// Start from 1, 2, 5 (symmetry: 1,3,7,9 are same; 2,4,6,8 are same)
		total += dfs(1, length-1) * 4
		total += dfs(2, length-1) * 4
		total += dfs(5, length-1)
	}
	return total
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", numberOfPatterns(1, 1))
	// Expected: 9

	// Test case 2
	fmt.Println("Test 2:", numberOfPatterns(1, 2))
	// Expected: 65

	// Test case 3
	fmt.Println("Test 3:", numberOfPatterns(3, 3))
	// Expected: 320 (length exactly 3)
}
```

## 0353 — Design Snake Game

```go
package main

// LeetCode #353: Design Snake Game
// https://leetcode.com/problems/design-snake-game/
// Difficulty: Medium [Paid]
// Time O(1) per move | Space O(n)

import "fmt"

type SnakeGame struct {
	width, height int
	food          [][]int
	foodIdx       int
	snake         [][2]int // head is last element
	body          map[[2]int]struct{}
	score         int
}

func Constructor(width int, height int, food [][]int) SnakeGame {
	head := [2]int{0, 0}
	return SnakeGame{
		width: width, height: height,
		food:    food,
		foodIdx: 0,
		snake:   [][2]int{head},
		body:    map[[2]int]struct{}{head: {}},
		score:   0,
	}
}

func (sg *SnakeGame) Move(direction string) int {
	head := sg.snake[len(sg.snake)-1]
	var next [2]int
	switch direction {
	case "U":
		next = [2]int{head[0] - 1, head[1]}
	case "D":
		next = [2]int{head[0] + 1, head[1]}
	case "L":
		next = [2]int{head[0], head[1] - 1}
	case "R":
		next = [2]int{head[0], head[1] + 1}
	}

	// Check boundaries
	if next[0] < 0 || next[0] >= sg.height || next[1] < 0 || next[1] >= sg.width {
		return -1
	}

	// Check if eating food
	eat := sg.foodIdx < len(sg.food) && next[0] == sg.food[sg.foodIdx][0] && next[1] == sg.food[sg.foodIdx][1]

	if eat {
		sg.score++
		sg.foodIdx++
	} else {
		// Remove tail
		tail := sg.snake[0]
		delete(sg.body, tail)
		sg.snake = sg.snake[1:]
	}

	// Check collision with self
	if _, hit := sg.body[next]; hit {
		return -1
	}

	sg.snake = append(sg.snake, next)
	sg.body[next] = struct{}{}
	return sg.score
}

func main() {
	// Test case 1
	sg := Constructor(3, 2, [][]int{{1, 2}, {0, 1}})
	fmt.Println("Move R:", sg.Move("R")) // at (0,1)
	fmt.Println("Move D:", sg.Move("D")) // at (1,1)
	fmt.Println("Move R:", sg.Move("R")) // at (1,2), eat food → score 1
	fmt.Println("Move U:", sg.Move("U")) // at (0,2)
	fmt.Println("Move L:", sg.Move("L")) // at (0,1) - hits body
	// Expected: 0, 0, 1, 1, -1

	fmt.Println()

	// Test case 2: Hit wall
	sg2 := Constructor(1, 1, [][]int{})
	fmt.Println("Move R:", sg2.Move("R"))
	// Expected: -1
}
```

## 0355 — Design Twitter

```go
package main

// LeetCode #355: Design Twitter
// https://leetcode.com/problems/design-twitter/
// Difficulty: Medium
// Time: O(n log n) for newsFeed | Space: O(n)

import (
	"container/heap"
	"fmt"
)

type Tweet struct {
	id   int
	time int
	next *Tweet
}

type Twitter struct {
	users map[int]*User
	time  int
}

type User struct {
	id       int
	tweets   *Tweet
	followee map[int]bool
}

type tweetHeapItem struct {
	tweet *Tweet
}

type tweetMaxHeap []*tweetHeapItem

func (h tweetMaxHeap) Len() int           { return len(h) }
func (h tweetMaxHeap) Less(i, j int) bool { return h[i].tweet.time > h[j].tweet.time }
func (h tweetMaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *tweetMaxHeap) Push(x interface{}) {
	*h = append(*h, x.(*tweetHeapItem))
}

func (h *tweetMaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func Constructor() Twitter {
	return Twitter{users: make(map[int]*User)}
}

func (tw *Twitter) getUser(id int) *User {
	if _, ok := tw.users[id]; !ok {
		tw.users[id] = &User{id: id, followee: make(map[int]bool)}
		tw.users[id].followee[id] = true
	}
	return tw.users[id]
}

func (tw *Twitter) PostTweet(userId int, tweetId int) {
	user := tw.getUser(userId)
	tw.time++
	tweet := &Tweet{id: tweetId, time: tw.time, next: user.tweets}
	user.tweets = tweet
}

func (tw *Twitter) GetNewsFeed(userId int) []int {
	user := tw.getUser(userId)

	h := &tweetMaxHeap{}
	heap.Init(h)

	for followeeId := range user.followee {
		if followee, ok := tw.users[followeeId]; ok && followee.tweets != nil {
			heap.Push(h, &tweetHeapItem{tweet: followee.tweets})
		}
	}

	result := make([]int, 0, 10)
	for h.Len() > 0 && len(result) < 10 {
		item := heap.Pop(h).(*tweetHeapItem)
		result = append(result, item.tweet.id)
		if item.tweet.next != nil {
			heap.Push(h, &tweetHeapItem{tweet: item.tweet.next})
		}
	}
	return result
}

func (tw *Twitter) Follow(followerId int, followeeId int) {
	follower := tw.getUser(followerId)
	follower.followee[followeeId] = true
	tw.getUser(followeeId)
}

func (tw *Twitter) Unfollow(followerId int, followeeId int) {
	follower := tw.getUser(followerId)
	if followerId != followeeId {
		delete(follower.followee, followeeId)
	}
}

func main() {
	tw := Constructor()
	tw.PostTweet(1, 5)
	fmt.Println("NewsFeed 1:", tw.GetNewsFeed(1))
	// Expected: [5]

	tw.Follow(1, 2)
	tw.PostTweet(2, 6)
	fmt.Println("NewsFeed 1 after follow:", tw.GetNewsFeed(1))
	// Expected: [6, 5]

	tw.Unfollow(1, 2)
	fmt.Println("NewsFeed 1 after unfollow:", tw.GetNewsFeed(1))
	// Expected: [5]
}
```

## 0356 — Line Reflection

```go
package main

// LeetCode #356: Line Reflection
// https://leetcode.com/problems/line-reflection/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func isReflected(points [][]int) bool {
	if len(points) == 0 {
		return true
	}

	set := make(map[[2]int]bool)
	minX, maxX := points[0][0], points[0][0]

	for _, p := range points {
		set[[2]int{p[0], p[1]}] = true
		if p[0] < minX {
			minX = p[0]
		}
		if p[0] > maxX {
			maxX = p[0]
		}
	}

	sum := minX + maxX // 2 * mid
	for _, p := range points {
		reflect := [2]int{sum - p[0], p[1]}
		if !set[reflect] {
			return false
		}
	}
	return true
}

func main() {
	// Test case 1: Points symmetric across vertical line x=1
	fmt.Println("Test 1:", isReflected([][]int{{1, 1}, {-1, 1}}))
	// Expected: true

	// Test case 2: Not symmetric
	fmt.Println("Test 2:", isReflected([][]int{{1, 1}, {-1, -1}}))
	// Expected: false

	// Test case 3: Points on same line
	fmt.Println("Test 3:", isReflected([][]int{{0, 0}, {1, 0}, {2, 0}}))
	// Expected: true (midline at x=1, 0 reflects to 2)
}
```

## 0357 — Count Numbers With Unique Digits

```go
package main

// LeetCode #357: Count Numbers with Unique Digits
// https://leetcode.com/problems/count-numbers-with-unique-digits/
// Difficulty: Medium
// Time: O(1) | Space: O(1)

import "fmt"

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}

	count := 10 // n=1: 0-9
	product := 9
	available := 9

	for i := 2; i <= n && i <= 10; i++ {
		product *= available
		available--
		count += product
	}
	return count
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", countNumbersWithUniqueDigits(2))
	// Expected: 91

	// Test case 2
	fmt.Println("Test 2:", countNumbersWithUniqueDigits(0))
	// Expected: 1

	// Test case 3
	fmt.Println("Test 3:", countNumbersWithUniqueDigits(3))
	// Expected: 739
}
```

## 0360 — Sort Transformed Array

```go
package main

// LeetCode #360: Sort Transformed Array
// https://leetcode.com/problems/sort-transformed-array/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func sortTransformedArray(nums []int, a int, b int, c int) []int {
	n := len(nums)
	result := make([]int, n)
	left, right := 0, n-1

	// Parabola opens upward → fill from right; downward → fill from left
	idx := n - 1
	if a < 0 {
		idx = 0
	}

	f := func(x int) int {
		return a*x*x + b*x + c
	}

	for left <= right {
		lv, rv := f(nums[left]), f(nums[right])
		if a >= 0 {
			if lv > rv {
				result[idx] = lv
				left++
			} else {
				result[idx] = rv
				right--
			}
			idx--
		} else {
			if lv < rv {
				result[idx] = lv
				left++
			} else {
				result[idx] = rv
				right--
			}
			idx++
		}
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", sortTransformedArray([]int{-4, -2, 2, 4}, 1, 3, 5))
	// Expected: [3, 9, 15, 33]

	// Test case 2
	fmt.Println("Test 2:", sortTransformedArray([]int{-4, -2, 2, 4}, -1, 3, 5))
	// Expected: [-23, -5, 1, 7]

	// Test case 3: Single element
	fmt.Println("Test 3:", sortTransformedArray([]int{0}, 1, 0, 0))
	// Expected: [0]
}
```

## 0361 — Bomb Enemy

```go
package main

// LeetCode #361: Bomb Enemy
// https://leetcode.com/problems/bomb-enemy/
// Difficulty: Medium [Paid]
// Time: O(m*n) | Space: O(n)

import "fmt"

func maxKilledEnemies(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	rowHits := 0
	colHits := make([]int, n)
	maxKill := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// Reset row hit count at start of row or after wall
			if j == 0 || grid[i][j-1] == 'W' {
				rowHits = 0
				for k := j; k < n && grid[i][k] != 'W'; k++ {
					if grid[i][k] == 'E' {
						rowHits++
					}
				}
			}

			// Reset col hit count at start of col or after wall
			if i == 0 || grid[i-1][j] == 'W' {
				colHits[j] = 0
				for k := i; k < m && grid[k][j] != 'W'; k++ {
					if grid[k][j] == 'E' {
						colHits[j]++
					}
				}
			}

			// Place bomb at empty cell
			if grid[i][j] == '0' {
				if rowHits+colHits[j] > maxKill {
					maxKill = rowHits + colHits[j]
				}
			}
		}
	}
	return maxKill
}

func main() {
	// Test case 1
	grid1 := [][]byte{
		{'0', 'E', '0', '0'},
		{'E', '0', 'W', 'E'},
		{'0', 'E', '0', '0'},
	}
	fmt.Println("Test 1:", maxKilledEnemies(grid1))
	// Expected: 3

	// Test case 2
	grid2 := [][]byte{{'W', 'W', 'W'}, {'0', '0', '0'}, {'E', 'E', 'E'}}
	fmt.Println("Test 2:", maxKilledEnemies(grid2))
	// Expected: 1

	// Test case 3: Empty
	grid3 := [][]byte{{'0'}}
	fmt.Println("Test 3:", maxKilledEnemies(grid3))
	// Expected: 0
}
```

## 0362 — Design Hit Counter

```go
package main

// LeetCode #362: Design Hit Counter
// https://leetcode.com/problems/design-hit-counter/
// Difficulty: Medium [Paid]
// Time: O(1) per hit | O(s) per getHits | Space: O(1) (fixed 300 buckets)

import "fmt"

type HitCounter struct {
	timestamps [300]int
	hits       [300]int
}

func Constructor() HitCounter {
	return HitCounter{}
}

func (hc *HitCounter) Hit(timestamp int) {
	idx := timestamp % 300
	if hc.timestamps[idx] != timestamp {
		hc.timestamps[idx] = timestamp
		hc.hits[idx] = 1
	} else {
		hc.hits[idx]++
	}
}

func (hc *HitCounter) GetHits(timestamp int) int {
	total := 0
	for i := 0; i < 300; i++ {
		if timestamp-hc.timestamps[i] < 300 {
			total += hc.hits[i]
		}
	}
	return total
}

func main() {
	hc := Constructor()
	hc.Hit(1)
	hc.Hit(2)
	hc.Hit(3)
	fmt.Println("Hits at t=4:", hc.GetHits(4))
	// Expected: 3

	hc.Hit(300)
	fmt.Println("Hits at t=300:", hc.GetHits(300))
	// Expected: 4

	fmt.Println("Hits at t=301:", hc.GetHits(301))
	// Expected: 3

	fmt.Println("Hits at t=302:", hc.GetHits(302))
	// Expected: 2
}
```

## 0364 — Nested List Weight Sum Ii

```go
package main

// LeetCode #364: Nested List Weight Sum II
// https://leetcode.com/problems/nested-list-weight-sum-ii/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(d)

import "fmt"

type NestedInteger struct {
	integer int
	list    []*NestedInteger
	isInt   bool
}

func NewInt(val int) *NestedInteger {
	return &NestedInteger{integer: val, isInt: true}
}

func NewList(items ...*NestedInteger) *NestedInteger {
	return &NestedInteger{list: items, isInt: false}
}

func (n NestedInteger) IsInteger() bool          { return n.isInt }
func (n NestedInteger) GetInteger() int           { return n.integer }
func (n NestedInteger) GetList() []*NestedInteger { return n.list }

func depthSumInverse(nestedList []*NestedInteger) int {
	// First pass: find max depth
	var maxDepth func(list []*NestedInteger, depth int) int
	maxDepth = func(list []*NestedInteger, depth int) int {
		maxD := depth
		for _, ni := range list {
			if !ni.IsInteger() {
				if d := maxDepth(ni.GetList(), depth+1); d > maxD {
					maxD = d
				}
			}
		}
		return maxD
	}

	md := maxDepth(nestedList, 1)

	// Second pass: compute weighted sum
	var dfs func(list []*NestedInteger, depth int) int
	dfs = func(list []*NestedInteger, depth int) int {
		total := 0
		for _, ni := range list {
			if ni.IsInteger() {
				total += ni.GetInteger() * (md - depth + 1)
			} else {
				total += dfs(ni.GetList(), depth+1)
			}
		}
		return total
	}
	return dfs(nestedList, 1)
}

func main() {
	// Test case 1: [[1,1],2,[1,1]]
	n1 := NewList(NewInt(1), NewInt(1))
	n2 := NewInt(2)
	n3 := NewList(NewInt(1), NewInt(1))
	fmt.Println("Test 1:", depthSumInverse([]*NestedInteger{n1, n2, n3}))
	// Expected: 8 (deepest=2, 1*1+1*1 + 2*2 + 1*1+1*1)

	// Test case 2: [1,[4,[6]]]
	inner := NewList(NewInt(6))
	mid := NewList(NewInt(4), inner)
	fmt.Println("Test 2:", depthSumInverse([]*NestedInteger{NewInt(1), mid}))
	// Expected: 17 (1*3 + 4*2 + 6*1)

	// Test case 3: Single integer
	fmt.Println("Test 3:", depthSumInverse([]*NestedInteger{NewInt(5)}))
	// Expected: 5
}
```

## 0365 — Water And Jug Problem

```go
package main

// LeetCode #365: Water and Jug Problem
// https://leetcode.com/problems/water-and-jug-problem/
// Difficulty: Medium
// Time: O(log min(x,y)) | Space: O(1)

import "fmt"

func canMeasureWater(x int, y int, target int) bool {
	if target > x+y {
		return false
	}
	return target%gcd(x, y) == 0
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", canMeasureWater(3, 5, 4))
	// Expected: true

	// Test case 2
	fmt.Println("Test 2:", canMeasureWater(2, 6, 5))
	// Expected: false

	// Test case 3
	fmt.Println("Test 3:", canMeasureWater(1, 2, 3))
	// Expected: true
}
```

## 0366 — Find Leaves Of Binary Tree

```go
package main

// LeetCode #366: Find Leaves of Binary Tree
// https://leetcode.com/problems/find-leaves-of-binary-tree/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findLeaves(root *TreeNode) [][]int {
	result := [][]int{}
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return -1
		}
		leftHeight := dfs(node.Left)
		rightHeight := dfs(node.Right)
		height := max(leftHeight, rightHeight) + 1

		if height >= len(result) {
			result = append(result, []int{})
		}
		result[height] = append(result[height], node.Val)
		return height
	}
	dfs(root)
	return result
}

func main() {
	// Test case 1: [1,2,3,4,5]
	root1 := &TreeNode{Val: 1}
	root1.Left = &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}
	root1.Right = &TreeNode{Val: 3}
	fmt.Println("Test 1:", findLeaves(root1))
	// Expected: [[4,5,3],[2],[1]]

	// Test case 2: Single node
	root2 := &TreeNode{Val: 1}
	fmt.Println("Test 2:", findLeaves(root2))
	// Expected: [[1]]

	// Test case 3: Nil
	fmt.Println("Test 3:", findLeaves(nil))
	// Expected: []
}
```

## 0368 — Largest Divisible Subset

```go
package main

// LeetCode #368: Largest Divisible Subset
// https://leetcode.com/problems/largest-divisible-subset/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import (
	"fmt"
	"sort"
)

func largestDivisibleSubset(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	sort.Ints(nums)
	n := len(nums)
	dp := make([]int, n) // size of largest subset ending at i
	prev := make([]int, n)
	maxIdx := 0

	for i := 0; i < n; i++ {
		dp[i] = 1
		prev[i] = -1
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
				prev[i] = j
			}
		}
		if dp[i] > dp[maxIdx] {
			maxIdx = i
		}
	}

	// Reconstruct
	result := make([]int, 0, dp[maxIdx])
	for i := maxIdx; i >= 0; i = prev[i] {
		result = append(result, nums[i])
		// Reverse the traversal
	}
	// Reverse result
	for left, right := 0, len(result)-1; left < right; left, right = left+1, right-1 {
		result[left], result[right] = result[right], result[left]
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", largestDivisibleSubset([]int{1, 2, 3}))
	// Expected: [1, 2] or [1, 3]

	// Test case 2
	fmt.Println("Test 2:", largestDivisibleSubset([]int{1, 2, 4, 8}))
	// Expected: [1, 2, 4, 8]

	// Test case 3
	fmt.Println("Test 3:", largestDivisibleSubset([]int{3, 4, 8, 16}))
	// Expected: [4, 8, 16] or [3] (without 3)
}
```

## 0369 — Plus One Linked List

```go
package main

// LeetCode #369: Plus One Linked List
// https://leetcode.com/problems/plus-one-linked-list/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func plusOne(head *ListNode) *ListNode {
	// Sentinel node
	sentinel := &ListNode{Next: head}
	notNine := sentinel

	// Find rightmost node that is not 9
	for node := head; node != nil; node = node.Next {
		if node.Val != 9 {
			notNine = node
		}
	}

	// Increment rightmost non-9 node
	notNine.Val++
	// Set all following 9s to 0
	for node := notNine.Next; node != nil; node = node.Next {
		node.Val = 0
	}

	if sentinel.Val == 1 {
		return sentinel
	}
	return head
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print(" -> ")
		}
		head = head.Next
	}
	fmt.Println()
}

func main() {
	// Test case 1: 1->2->3
	head1 := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	fmt.Print("Test 1: ")
	printList(plusOne(head1))
	// Expected: 1->2->4

	// Test case 2: 9->9->9
	head2 := &ListNode{9, &ListNode{9, &ListNode{9, nil}}}
	fmt.Print("Test 2: ")
	printList(plusOne(head2))
	// Expected: 1->0->0->0

	// Test case 3: 9
	head3 := &ListNode{9, nil}
	fmt.Print("Test 3: ")
	printList(plusOne(head3))
	// Expected: 1->0
}
```

## 0370 — Range Addition

```go
package main

// LeetCode #370: Range Addition
// https://leetcode.com/problems/range-addition/
// Difficulty: Medium [Paid]
// Time: O(n + k) | Space: O(n)

import "fmt"

func getModifiedArray(length int, updates [][]int) []int {
	arr := make([]int, length+1)

	for _, upd := range updates {
		start, end, inc := upd[0], upd[1], upd[2]
		arr[start] += inc
		arr[end+1] -= inc
	}

	// Prefix sum
	result := make([]int, length)
	sum := 0
	for i := 0; i < length; i++ {
		sum += arr[i]
		result[i] = sum
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", getModifiedArray(5, [][]int{{1, 3, 2}, {2, 4, 3}, {0, 2, -2}}))
	// Expected: [-2, 0, 3, 5, 3]

	// Test case 2: Single update
	fmt.Println("Test 2:", getModifiedArray(3, [][]int{{0, 2, 5}}))
	// Expected: [5, 5, 5]

	// Test case 3: No updates
	fmt.Println("Test 3:", getModifiedArray(3, [][]int{}))
	// Expected: [0, 0, 0]
}
```

## 0371 — Sum Of Two Integers

```go
package main

// LeetCode #371: Sum of Two Integers
// https://leetcode.com/problems/sum-of-two-integers/
// Difficulty: Medium
// Time: O(1) | Space: O(1)

import "fmt"

func getSum(a int, b int) int {
	for b != 0 {
		carry := a & b
		a = a ^ b
		b = carry << 1
	}
	return a
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", getSum(1, 2))
	// Expected: 3

	// Test case 2
	fmt.Println("Test 2:", getSum(2, 3))
	// Expected: 5

	// Test case 3: Negative numbers
	fmt.Println("Test 3:", getSum(-1, 1))
	// Expected: 0
}
```

## 0372 — Super Pow

```go
package main

// LeetCode #372: Super Pow
// https://leetcode.com/problems/super-pow/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

const mod = 1337

func superPow(a int, b []int) int {
	a %= mod
	result := 1

	for _, digit := range b {
		result = (powMod(result, 10) * powMod(a, digit)) % mod
	}
	return result
}

func powMod(base, exp int) int {
	result := 1
	base %= mod
	for i := 0; i < exp; i++ {
		result = (result * base) % mod
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", superPow(2, []int{3}))
	// Expected: 8

	// Test case 2
	fmt.Println("Test 2:", superPow(2, []int{1, 0}))
	// Expected: 1024

	// Test case 3
	fmt.Println("Test 3:", superPow(1, []int{4, 3, 3, 8, 5, 2}))
	// Expected: 1
}
```

## 0373 — Find K Pairs With Smallest Sums

```go
package main

// LeetCode #373: Find K Pairs with Smallest Sums
// https://leetcode.com/problems/find-k-pairs-with-smallest-sums/
// Difficulty: Medium
// Time: O(k log min(k, n)) | Space: O(k)

import (
	"container/heap"
	"fmt"
)

type pair struct {
	i, j int
	sum  int
}

type minHeap []pair

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i].sum < h[j].sum }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(pair)) }
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	if len(nums1) == 0 || len(nums2) == 0 || k == 0 {
		return [][]int{}
	}

	h := &minHeap{}
	heap.Init(h)

	// Push first element of nums1 paired with each element of nums2
	for j := 0; j < len(nums2) && j < k; j++ {
		heap.Push(h, pair{0, j, nums1[0] + nums2[j]})
	}

	result := make([][]int, 0, k)
	for h.Len() > 0 && len(result) < k {
		p := heap.Pop(h).(pair)
		result = append(result, []int{nums1[p.i], nums2[p.j]})
		if p.i+1 < len(nums1) {
			heap.Push(h, pair{p.i + 1, p.j, nums1[p.i+1] + nums2[p.j]})
		}
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", kSmallestPairs([]int{1, 7, 11}, []int{2, 4, 6}, 3))
	// Expected: [[1,2],[1,4],[1,6]]

	// Test case 2
	fmt.Println("Test 2:", kSmallestPairs([]int{1, 1, 2}, []int{1, 2, 3}, 2))
	// Expected: [[1,1],[1,1]]

	// Test case 3
	fmt.Println("Test 3:", kSmallestPairs([]int{1, 2}, []int{3}, 3))
	// Expected: [[1,3],[2,3]]
}
```

## 0375 — Guess Number Higher Or Lower Ii

```go
package main

// LeetCode #375: Guess Number Higher or Lower II
// https://leetcode.com/problems/guess-number-higher-or-lower-ii/
// Difficulty: Medium
// Time: O(n^3) | Space: O(n^2)

import "fmt"

func getMoneyAmount(n int) int {
	dp := make([][]int, n+2)
	for i := range dp {
		dp[i] = make([]int, n+2)
	}

	for length := 2; length <= n; length++ {
		for start := 1; start <= n-length+1; start++ {
			end := start + length - 1
			dp[start][end] = 1<<31 - 1
			for pivot := start; pivot <= end; pivot++ {
				// Cost if pivot is wrong: pivot + max(cost of left, cost of right)
				cost := pivot + max(dp[start][pivot-1], dp[pivot+1][end])
				if cost < dp[start][end] {
					dp[start][end] = cost
				}
			}
		}
	}
	return dp[1][n]
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", getMoneyAmount(10))
	// Expected: 16

	// Test case 2
	fmt.Println("Test 2:", getMoneyAmount(1))
	// Expected: 0

	// Test case 3
	fmt.Println("Test 3:", getMoneyAmount(2))
	// Expected: 1
}
```

## 0376 — Wiggle Subsequence

```go
package main

// LeetCode #376: Wiggle Subsequence
// https://leetcode.com/problems/wiggle-subsequence/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func wiggleMaxLength(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	up, down := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			up = down + 1
		} else if nums[i] < nums[i-1] {
			down = up + 1
		}
	}
	if up > down {
		return up
	}
	return down
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", wiggleMaxLength([]int{1, 7, 4, 9, 2, 5}))
	// Expected: 6

	// Test case 2
	fmt.Println("Test 2:", wiggleMaxLength([]int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}))
	// Expected: 7

	// Test case 3
	fmt.Println("Test 3:", wiggleMaxLength([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
	// Expected: 2
}
```

## 0377 — Combination Sum Iv

```go
package main

// LeetCode #377: Combination Sum IV
// https://leetcode.com/problems/combination-sum-iv/
// Difficulty: Medium
// Time: O(target * n) | Space: O(target)

import "fmt"

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1

	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if num <= i {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", combinationSum4([]int{1, 2, 3}, 4))
	// Expected: 7

	// Test case 2
	fmt.Println("Test 2:", combinationSum4([]int{9}, 3))
	// Expected: 0

	// Test case 3
	fmt.Println("Test 3:", combinationSum4([]int{1, 2, 3}, 3))
	// Expected: 4
}
```

## 0378 — Kth Smallest Element In A Sorted Matrix

```go
package main

// LeetCode #378: Kth Smallest Element in a Sorted Matrix
// https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/
// Difficulty: Medium
// Time: O((m+n) * log(max-min)) | Space: O(1)

import "fmt"

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	lo, hi := matrix[0][0], matrix[n-1][n-1]

	for lo < hi {
		mid := lo + (hi-lo)/2
		count := 0
		row, col := n-1, 0
		for row >= 0 && col < n {
			if matrix[row][col] <= mid {
				count += row + 1
				col++
			} else {
				row--
			}
		}
		if count < k {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

func main() {
	// Test case 1
	matrix1 := [][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}
	fmt.Println("Test 1:", kthSmallest(matrix1, 8))
	// Expected: 13

	// Test case 2
	fmt.Println("Test 2:", kthSmallest([][]int{{-5}}, 1))
	// Expected: -5

	// Test case 3
	matrix3 := [][]int{{1, 2}, {1, 3}}
	fmt.Println("Test 3:", kthSmallest(matrix3, 2))
	// Expected: 1
}
```

## 0379 — Design Phone Directory

```go
package main

// LeetCode #379: Design Phone Directory
// https://leetcode.com/problems/design-phone-directory/
// Difficulty: Medium [Paid]
// Time: O(1) per operation | Space: O(n)

import "fmt"

type PhoneDirectory struct {
	available []int
	used      []bool
	idx       int
}

func Constructor(maxNumbers int) PhoneDirectory {
	available := make([]int, maxNumbers)
	for i := 0; i < maxNumbers; i++ {
		available[i] = i
	}
	return PhoneDirectory{
		available: available,
		used:      make([]bool, maxNumbers),
		idx:       0,
	}
}

func (pd *PhoneDirectory) Get() int {
	if pd.idx >= len(pd.available) {
		return -1
	}
	num := pd.available[pd.idx]
	pd.idx++
	pd.used[num] = true
	return num
}

func (pd *PhoneDirectory) Check(number int) bool {
	if number < 0 || number >= len(pd.used) {
		return false
	}
	return !pd.used[number]
}

func (pd *PhoneDirectory) Release(number int) {
	if number < 0 || number >= len(pd.used) || !pd.used[number] {
		return
	}
	pd.used[number] = false
	pd.idx--
	pd.available[pd.idx] = number
}

func main() {
	pd := Constructor(3)
	fmt.Println("Get:", pd.Get())      // 0
	fmt.Println("Get:", pd.Get())      // 1
	fmt.Println("Check 2:", pd.Check(2)) // true
	fmt.Println("Check 1:", pd.Check(1)) // false (in use)
	pd.Release(1)
	fmt.Println("Check 1 after release:", pd.Check(1)) // true
	fmt.Println("Get:", pd.Get())      // 1 (recycled)
	fmt.Println("Get:", pd.Get())      // 2
	fmt.Println("Get:", pd.Get())      // -1 (none left)
}
```

## 0380 — Insert Delete Getrandom O1

```go
package main

// LeetCode #380: Insert Delete GetRandom O(1)
// https://leetcode.com/problems/insert-delete-getrandom-o1/
// Difficulty: Medium
// Time: O(1) per operation | Space: O(n)

import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	nums  []int
	pos   map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{pos: make(map[int]int)}
}

func (rs *RandomizedSet) Insert(val int) bool {
	if _, ok := rs.pos[val]; ok {
		return false
	}
	rs.pos[val] = len(rs.nums)
	rs.nums = append(rs.nums, val)
	return true
}

func (rs *RandomizedSet) Remove(val int) bool {
	idx, ok := rs.pos[val]
	if !ok {
		return false
	}

	// Swap with last element
	last := len(rs.nums) - 1
	lastVal := rs.nums[last]
	rs.nums[idx] = lastVal
	rs.pos[lastVal] = idx
	rs.nums = rs.nums[:last]
	delete(rs.pos, val)
	return true
}

func (rs *RandomizedSet) GetRandom() int {
	return rs.nums[rand.Intn(len(rs.nums))]
}

func main() {
	rs := Constructor()
	fmt.Println("Insert 1:", rs.Insert(1)) // true
	fmt.Println("Remove 2:", rs.Remove(2)) // false
	fmt.Println("Insert 2:", rs.Insert(2)) // true
	fmt.Println("GetRandom:", rs.GetRandom()) // 1 or 2
	fmt.Println("Remove 1:", rs.Remove(1)) // true
	fmt.Println("Insert 2:", rs.Insert(2)) // false (already present)
	fmt.Println("GetRandom:", rs.GetRandom()) // 2
}
```

## 0382 — Linked List Random Node

```go
package main

// LeetCode #382: Linked List Random Node
// https://leetcode.com/problems/linked-list-random-node/
// Difficulty: Medium
// Time: O(1) for init, O(n) for getRandom | Space: O(1)

import (
	"fmt"
	"math/rand"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	head *ListNode
}

func Constructor(head *ListNode) Solution {
	return Solution{head: head}
}

// Reservoir sampling: O(n), uniform probability
func (s *Solution) GetRandom() int {
	result := s.head.Val
	node := s.head.Next
	i := 1
	for node != nil {
		i++
		if rand.Intn(i) == 0 {
			result = node.Val
		}
		node = node.Next
	}
	return result
}

func main() {
	// Test case: 1->2->3
	head := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	sol := Constructor(head)

	// Run multiple times to show randomness
	counts := map[int]int{}
	for i := 0; i < 30000; i++ {
		counts[sol.GetRandom()]++
	}
	fmt.Println("Counts:", counts)
	// Expected: roughly 10000 each
}
```

## 0384 — Shuffle An Array

```go
package main

// LeetCode #384: Shuffle an Array
// https://leetcode.com/problems/shuffle-an-array/
// Difficulty: Medium
// Time: O(n) per shuffle | Space: O(n)

import (
	"fmt"
	"math/rand"
)

type Solution struct {
	original []int
}

func Constructor(nums []int) Solution {
	orig := make([]int, len(nums))
	copy(orig, nums)
	return Solution{original: orig}
}

func (s *Solution) Reset() []int {
	result := make([]int, len(s.original))
	copy(result, s.original)
	return result
}

func (s *Solution) Shuffle() []int {
	result := make([]int, len(s.original))
	copy(result, s.original)
	// Fisher-Yates
	for i := len(result) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		result[i], result[j] = result[j], result[i]
	}
	return result
}

func main() {
	sol := Constructor([]int{1, 2, 3})
	fmt.Println("Reset:", sol.Reset())
	fmt.Println("Shuffle:", sol.Shuffle())
	fmt.Println("Shuffle:", sol.Shuffle())
	fmt.Println("Reset:", sol.Reset())
}
```

## 0385 — Mini Parser

```go
package main

// LeetCode #385: Mini Parser
// https://leetcode.com/problems/mini-parser/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"strconv"
	"strings"
)

type NestedInteger struct {
	integer int
	list    []*NestedInteger
	isInt   bool
}

func NewInt(val int) *NestedInteger {
	return &NestedInteger{integer: val, isInt: true}
}

func NewList(items ...*NestedInteger) *NestedInteger {
	return &NestedInteger{list: items, isInt: false}
}

func (n NestedInteger) IsInteger() bool          { return n.isInt }
func (n NestedInteger) GetInteger() int           { return n.integer }
func (n NestedInteger) GetList() []*NestedInteger { return n.list }
func (n *NestedInteger) Add(elem NestedInteger) {
	n.list = append(n.list, &elem)
}
func (n *NestedInteger) SetInteger(val int) {
	n.integer = val
	n.isInt = true
}

func deserialize(s string) *NestedInteger {
	if len(s) == 0 {
		return nil
	}
	if s[0] != '[' {
		val, _ := strconv.Atoi(s)
		return NewInt(val)
	}

	stack := []*NestedInteger{}
	var current *NestedInteger
	i := 0

	for i < len(s) {
		if s[i] == '[' {
			ni := &NestedInteger{isInt: false, list: []*NestedInteger{}}
			if current != nil {
				stack = append(stack, current)
			}
			current = ni
			i++
		} else if s[i] == ']' {
			if len(stack) > 0 {
				parent := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				parent.list = append(parent.list, current)
				current = parent
			}
			i++
		} else if s[i] == ',' {
			i++
		} else {
			// Parse integer
			j := i
			for j < len(s) && (s[j] == '-' || (s[j] >= '0' && s[j] <= '9')) {
				j++
			}
			val, _ := strconv.Atoi(s[i:j])
			nestedInt := NewInt(val)
			current.list = append(current.list, nestedInt)
			i = j
		}
	}
	return current
}

func (n *NestedInteger) String() string {
	if n.isInt {
		return strconv.Itoa(n.integer)
	}
	parts := make([]string, len(n.list))
	for i, child := range n.list {
		parts[i] = child.String()
	}
	return "[" + strings.Join(parts, ",") + "]"
}

func main() {
	// Test case 1
	n1 := deserialize("324")
	fmt.Println("Test 1:", n1.String())
	// Expected: 324

	// Test case 2
	n2 := deserialize("[123,[456,[789]]]")
	fmt.Println("Test 2:", n2.String())
	// Expected: [123,[456,[789]]]

	// Test case 3
	n3 := deserialize("[]")
	fmt.Println("Test 3:", n3.String())
	// Expected: []
}
```

## 0386 — Lexicographical Numbers

```go
package main

// LeetCode #386: Lexicographical Numbers
// https://leetcode.com/problems/lexicographical-numbers/
// Difficulty: Medium
// Time: O(n) | Space: O(1) (excluding output)

import "fmt"

func lexicalOrder(n int) []int {
	result := make([]int, 0, n)
	cur := 1

	for len(result) < n {
		result = append(result, cur)

		if cur*10 <= n {
			cur *= 10
		} else {
			for cur%10 == 9 || cur+1 > n {
				cur /= 10
			}
			cur++
		}
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", lexicalOrder(13))
	// Expected: [1,10,11,12,13,2,3,4,5,6,7,8,9]

	// Test case 2
	fmt.Println("Test 2:", lexicalOrder(2))
	// Expected: [1,2]

	// Test case 3
	fmt.Println("Test 3:", lexicalOrder(25))
	// Expected: [1,10,11,12,13,14,15,16,17,18,19,2,20,21,22,23,24,25,3,4,5,6,7,8,9]
}
```

## 0388 — Longest Absolute File Path

```go
package main

// LeetCode #388: Longest Absolute File Path
// https://leetcode.com/problems/longest-absolute-file-path/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"strings"
)

func lengthLongestPath(input string) int {
	lines := strings.Split(input, "\n")
	stack := make([]int, 0) // lengths at each depth
	maxLen := 0

	for _, line := range lines {
		// Count tabs to determine depth
		depth := 0
		for depth < len(line) && line[depth] == '\t' {
			depth++
		}

		// Current name without tabs
		name := line[depth:]

		// Pop stack to correct depth
		for len(stack) > depth {
			stack = stack[:len(stack)-1]
		}

		// Total length of current path
		total := len(name)
		if len(stack) > 0 {
			total += stack[len(stack)-1] + 1 // +1 for '/'
		}

		// Check if it's a file
		if strings.Contains(name, ".") {
			if total > maxLen {
				maxLen = total
			}
		}

		stack = append(stack, total)
	}
	return maxLen
}

func main() {
	// Test case 1
	input1 := "dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext"
	fmt.Println("Test 1:", lengthLongestPath(input1))
	// Expected: 20 ("dir/subdir2/file.ext")

	// Test case 2
	input2 := "dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext"
	fmt.Println("Test 2:", lengthLongestPath(input2))
	// Expected: 32 ("dir/subdir2/subsubdir2/file2.ext")

	// Test case 3: No files
	fmt.Println("Test 3:", lengthLongestPath("dir\n\tsubdir"))
	// Expected: 0
}
```

## 0390 — Elimination Game

```go
package main

// LeetCode #390: Elimination Game
// https://leetcode.com/problems/elimination-game/
// Difficulty: Medium
// Time: O(log n) | Space: O(1)

import "fmt"

func lastRemaining(n int) int {
	head := 1
	remaining := n
	step := 1
	leftToRight := true

	for remaining > 1 {
		if leftToRight || remaining%2 == 1 {
			head += step
		}
		remaining /= 2
		step *= 2
		leftToRight = !leftToRight
	}
	return head
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", lastRemaining(9))
	// Expected: 6

	// Test case 2
	fmt.Println("Test 2:", lastRemaining(1))
	// Expected: 1

	// Test case 3
	fmt.Println("Test 3:", lastRemaining(100))
	// Expected: 54
}
```

## 0393 — Utf 8 Validation

```go
package main

// LeetCode #393: UTF-8 Validation
// https://leetcode.com/problems/utf-8-validation/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func validUtf8(data []int) bool {
	remaining := 0

	for _, b := range data {
		if remaining == 0 {
			if b>>3 == 0b11110 {
				remaining = 3
			} else if b>>4 == 0b1110 {
				remaining = 2
			} else if b>>5 == 0b110 {
				remaining = 1
			} else if b>>7 != 0 {
				return false
			}
		} else {
			if b>>6 != 0b10 {
				return false
			}
			remaining--
		}
	}
	return remaining == 0
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", validUtf8([]int{197, 130, 1}))
	// Expected: true

	// Test case 2
	fmt.Println("Test 2:", validUtf8([]int{235, 140, 4}))
	// Expected: false

	// Test case 3
	fmt.Println("Test 3:", validUtf8([]int{255}))
	// Expected: false
}
```

## 0394 — Decode String

```go
package main

// LeetCode #394: Decode String
// https://leetcode.com/problems/decode-string/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"strings"
)

func decodeString(s string) string {
	numStack := make([]int, 0)
	strStack := make([]string, 0)
	curNum := 0
	curStr := ""

	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			curNum = curNum*10 + int(ch-'0')
		} else if ch == '[' {
			numStack = append(numStack, curNum)
			strStack = append(strStack, curStr)
			curNum = 0
			curStr = ""
		} else if ch == ']' {
			num := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]
			prevStr := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]
			curStr = prevStr + strings.Repeat(curStr, num)
		} else {
			curStr += string(ch)
		}
	}
	return curStr
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", decodeString("3[a]2[bc]"))
	// Expected: "aaabcbc"

	// Test case 2
	fmt.Println("Test 2:", decodeString("3[a2[c]]"))
	// Expected: "accaccacc"

	// Test case 3
	fmt.Println("Test 3:", decodeString("2[abc]3[cd]ef"))
	// Expected: "abcabccdcdcdef"
}
```

## 0395 — Longest Substring With At Least K Repeating Characters

```go
package main

// LeetCode #395: Longest Substring with At Least K Repeating Characters
// https://leetcode.com/problems/longest-substring-with-at-least-k-repeating-characters/
// Difficulty: Medium
// Time: O(n^2) worst case, O(n) average | Space: O(n)

import "fmt"

func longestSubstring(s string, k int) int {
	return longestSubstringHelper(s, 0, len(s), k)
}

func longestSubstringHelper(s string, start, end, k int) int {
	if end-start < k {
		return 0
	}

	// Count frequencies
	freq := [26]int{}
	for i := start; i < end; i++ {
		freq[s[i]-'a']++
	}

	// Find split point where char freq < k
	for i := start; i < end; i++ {
		if freq[s[i]-'a'] < k {
			left := longestSubstringHelper(s, start, i, k)
			right := longestSubstringHelper(s, i+1, end, k)
			if left > right {
				return left
			}
			return right
		}
	}

	return end - start
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", longestSubstring("aaabb", 3))
	// Expected: 3 ("aaa")

	// Test case 2
	fmt.Println("Test 2:", longestSubstring("ababbc", 2))
	// Expected: 5 ("ababb")

	// Test case 3
	fmt.Println("Test 3:", longestSubstring("aaabbb", 3))
	// Expected: 6
}
```

## 0396 — Rotate Function

```go
package main

// LeetCode #396: Rotate Function
// https://leetcode.com/problems/rotate-function/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func maxRotateFunction(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	sum := 0
	f0 := 0
	for i, num := range nums {
		sum += num
		f0 += i * num
	}

	maxVal := f0
	prev := f0
	for k := 1; k < n; k++ {
		curr := prev + sum - n*nums[n-k]
		if curr > maxVal {
			maxVal = curr
		}
		prev = curr
	}
	return maxVal
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", maxRotateFunction([]int{4, 3, 2, 6}))
	// Expected: 26

	// Test case 2
	fmt.Println("Test 2:", maxRotateFunction([]int{100}))
	// Expected: 0

	// Test case 3
	fmt.Println("Test 3:", maxRotateFunction([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	// Expected: 330
}
```

## 0397 — Integer Replacement

```go
package main

// LeetCode #397: Integer Replacement
// https://leetcode.com/problems/integer-replacement/
// Difficulty: Medium
// Time: O(log n) | Space: O(1)

import "fmt"

func integerReplacement(n int) int {
	count := 0
	for n > 1 {
		if n&1 == 0 {
			n >>= 1
		} else if n == 3 || n&3 == 1 {
			n--
		} else {
			n++
		}
		count++
	}
	return count
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", integerReplacement(8))
	// Expected: 3

	// Test case 2
	fmt.Println("Test 2:", integerReplacement(7))
	// Expected: 4

	// Test case 3
	fmt.Println("Test 3:", integerReplacement(4))
	// Expected: 2
}
```

## 0398 — Random Pick Index

```go
package main

// LeetCode #398: Random Pick Index
// https://leetcode.com/problems/random-pick-index/
// Difficulty: Medium
// Time: O(n) for init, O(1) for pick | Space: O(n)

import (
	"fmt"
	"math/rand"
)

type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	return Solution{nums: nums}
}

func (s *Solution) Pick(target int) int {
	// Reservoir sampling
	count := 0
	result := 0
	for i, num := range s.nums {
		if num == target {
			count++
			if rand.Intn(count) == 0 {
				result = i
			}
		}
	}
	return result
}

func main() {
	sol := Constructor([]int{1, 2, 3, 3, 3})
	counts := map[int]int{}
	for i := 0; i < 30000; i++ {
		counts[sol.Pick(3)]++
	}
	fmt.Println("Counts for target=3:", counts)
	// Expected: roughly 10000 each for indices 2,3,4
}
```

## 0399 — Evaluate Division

```go
package main

// LeetCode #399: Evaluate Division
// https://leetcode.com/problems/evaluate-division/
// Difficulty: Medium
// Time: O(n + q*n) | Space: O(n)

import "fmt"

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	// Build graph
	graph := make(map[string]map[string]float64)
	for i, eq := range equations {
		a, b := eq[0], eq[1]
		if graph[a] == nil {
			graph[a] = make(map[string]float64)
		}
		if graph[b] == nil {
			graph[b] = make(map[string]float64)
		}
		graph[a][b] = values[i]
		graph[b][a] = 1.0 / values[i]
	}

	var dfs func(src, dst string, visited map[string]bool) float64
	dfs = func(src, dst string, visited map[string]bool) float64 {
		if _, ok := graph[src]; !ok {
			return -1.0
		}
		if src == dst {
			return 1.0
		}
		visited[src] = true
		for neighbor, val := range graph[src] {
			if visited[neighbor] {
				continue
			}
			if neighbor == dst {
				return val
			}
			if result := dfs(neighbor, dst, visited); result != -1.0 {
				return val * result
			}
		}
		return -1.0
	}

	result := make([]float64, len(queries))
	for i, q := range queries {
		visited := make(map[string]bool)
		result[i] = dfs(q[0], q[1], visited)
	}
	return result
}

func main() {
	// Test case 1
	eq1 := [][]string{{"a", "b"}, {"b", "c"}}
	val1 := []float64{2.0, 3.0}
	q1 := [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}
	fmt.Println("Test 1:", calcEquation(eq1, val1, q1))
	// Expected: [6.0, 0.5, -1.0, 1.0, -1.0]

	// Test case 2
	eq2 := [][]string{{"a", "b"}, {"c", "d"}}
	val2 := []float64{1.0, 1.0}
	q2 := [][]string{{"a", "c"}, {"b", "d"}, {"b", "a"}}
	fmt.Println("Test 2:", calcEquation(eq2, val2, q2))
	// Expected: [-1, -1, 1]
}
```

## 0400 — Nth Digit

```go
package main

// LeetCode #400: Nth Digit
// https://leetcode.com/problems/nth-digit/
// Difficulty: Medium
// Time: O(log n) | Space: O(1)

import (
	"fmt"
	"strconv"
)

func findNthDigit(n int) int {
	// 1-9: 9 digits,  10-99: 90*2 digits,  100-999: 900*3 digits
	length := 1
	count := 9
	start := 1

	for n > length*count {
		n -= length * count
		length++
		count *= 10
		start *= 10
	}

	// Find the actual number
	num := start + (n-1)/length
	// Find the digit within the number
	digitIdx := (n - 1) % length
	digitStr := strconv.Itoa(num)
	return int(digitStr[digitIdx] - '0')
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", findNthDigit(3))
	// Expected: 3

	// Test case 2
	fmt.Println("Test 2:", findNthDigit(11))
	// Expected: 0 (from 10)

	// Test case 3
	fmt.Println("Test 3:", findNthDigit(190))
	// Expected: 1 (from 100)
}
```

## 0402 — Remove K Digits

```go
package main

// LeetCode #402: Remove K Digits
// https://leetcode.com/problems/remove-k-digits/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"strings"
)

func removeKdigits(num string, k int) string {
	if k >= len(num) {
		return "0"
	}

	stack := make([]byte, 0, len(num))
	for i := 0; i < len(num); i++ {
		for k > 0 && len(stack) > 0 && stack[len(stack)-1] > num[i] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, num[i])
	}

	// If still need to remove, remove from end
	stack = stack[:len(stack)-k]

	// Remove leading zeros
	result := strings.TrimLeft(string(stack), "0")
	if result == "" {
		return "0"
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", removeKdigits("1432219", 3))
	// Expected: "1219"

	// Test case 2
	fmt.Println("Test 2:", removeKdigits("10200", 1))
	// Expected: "200"

	// Test case 3
	fmt.Println("Test 3:", removeKdigits("10", 2))
	// Expected: "0"
}
```

## 0406 — Queue Reconstruction By Height

```go
package main

// LeetCode #406: Queue Reconstruction by Height
// https://leetcode.com/problems/queue-reconstruction-by-height/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import (
	"fmt"
	"sort"
)

func reconstructQueue(people [][]int) [][]int {
	// Sort by height descending, then by k ascending
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] != people[j][0] {
			return people[i][0] > people[j][0]
		}
		return people[i][1] < people[j][1]
	})

	result := make([][]int, 0, len(people))
	for _, p := range people {
		// Insert at index k
		k := p[1]
		result = append(result, nil)
		copy(result[k+1:], result[k:])
		result[k] = p
	}
	return result
}

func main() {
	// Test case 1
	p1 := [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}
	fmt.Println("Test 1:", reconstructQueue(p1))
	// Expected: [[5,0],[7,0],[5,2],[6,1],[4,4],[7,1]]

	// Test case 2
	p2 := [][]int{{6, 0}, {5, 0}, {4, 0}, {3, 2}, {2, 2}, {1, 4}}
	fmt.Println("Test 2:", reconstructQueue(p2))
	// Expected: [[4,0],[5,0],[2,2],[3,2],[1,4],[6,0]]
}
```

## 0413 — Arithmetic Slices

```go
package main

// LeetCode #413: Arithmetic Slices
// https://leetcode.com/problems/arithmetic-slices/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func numberOfArithmeticSlices(nums []int) int {
	if len(nums) < 3 {
		return 0
	}

	total := 0
	curr := 0

	for i := 2; i < len(nums); i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			curr++
			total += curr
		} else {
			curr = 0
		}
	}
	return total
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", numberOfArithmeticSlices([]int{1, 2, 3, 4}))
	// Expected: 3

	// Test case 2
	fmt.Println("Test 2:", numberOfArithmeticSlices([]int{1}))
	// Expected: 0

	// Test case 3
	fmt.Println("Test 3:", numberOfArithmeticSlices([]int{1, 2, 3, 8, 9, 10}))
	// Expected: 2
}
```

## 0416 — Partition Equal Subset Sum

```go
package main

// LeetCode #416: Partition Equal Subset Sum
// https://leetcode.com/problems/partition-equal-subset-sum/
// Difficulty: Medium
// Time: O(n * sum) | Space: O(sum)

import "fmt"

func canPartition(nums []int) bool {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum%2 != 0 {
		return false
	}

	target := sum / 2
	dp := make([]bool, target+1)
	dp[0] = true

	for _, num := range nums {
		for s := target; s >= num; s-- {
			if dp[s-num] {
				dp[s] = true
			}
		}
	}
	return dp[target]
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", canPartition([]int{1, 5, 11, 5}))
	// Expected: true

	// Test case 2
	fmt.Println("Test 2:", canPartition([]int{1, 2, 3, 5}))
	// Expected: false

	// Test case 3
	fmt.Println("Test 3:", canPartition([]int{2, 2, 2, 2}))
	// Expected: true
}
```

## 0417 — Pacific Atlantic Water Flow

```go
package main

// LeetCode #417: Pacific Atlantic Water Flow
// https://leetcode.com/problems/pacific-atlantic-water-flow/
// Difficulty: Medium
// Time: O(m*n) | Space: O(m*n)

import "fmt"

func pacificAtlantic(heights [][]int) [][]int {
	if len(heights) == 0 {
		return [][]int{}
	}
	m, n := len(heights), len(heights[0])
	pacific := make([][]bool, m)
	atlantic := make([][]bool, m)
	for i := range pacific {
		pacific[i] = make([]bool, n)
		atlantic[i] = make([]bool, n)
	}

	var dfs func(i, j int, visited [][]bool)
	dfs = func(i, j int, visited [][]bool) {
		visited[i][j] = true
		dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
		for _, d := range dirs {
			ni, nj := i+d[0], j+d[1]
			if ni >= 0 && ni < m && nj >= 0 && nj < n && !visited[ni][nj] && heights[ni][nj] >= heights[i][j] {
				dfs(ni, nj, visited)
			}
		}
	}

	// Pacific: top and left edges
	for i := 0; i < m; i++ {
		dfs(i, 0, pacific)
	}
	for j := 0; j < n; j++ {
		dfs(0, j, pacific)
	}

	// Atlantic: bottom and right edges
	for i := 0; i < m; i++ {
		dfs(i, n-1, atlantic)
	}
	for j := 0; j < n; j++ {
		dfs(m-1, j, atlantic)
	}

	result := [][]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if pacific[i][j] && atlantic[i][j] {
				result = append(result, []int{i, j})
			}
		}
	}
	return result
}

func main() {
	// Test case 1
	h1 := [][]int{
		{1, 2, 2, 3, 5},
		{3, 2, 3, 4, 4},
		{2, 4, 5, 3, 1},
		{6, 7, 1, 4, 5},
		{5, 1, 1, 2, 4},
	}
	fmt.Println("Test 1:", pacificAtlantic(h1))
	// Expected: [[0,4],[1,3],[1,4],[2,2],[3,0],[3,1],[4,0]] (in any order)

	// Test case 2
	h2 := [][]int{{1}}
	fmt.Println("Test 2:", pacificAtlantic(h2))
	// Expected: [[0,0]]
}
```

## 0418 — Sentence Screen Fitting

```go
package main

// LeetCode #418: Sentence Screen Fitting
// https://leetcode.com/problems/sentence-screen-fitting/
// Difficulty: Medium [Paid]
// Time: O(rows * avgWordLen) | Space: O(1)

import "fmt"

func wordsTyping(sentence []string, rows int, cols int) int {
	s := ""
	for _, w := range sentence {
		s += w + " "
	}

	n := len(s)
	start := 0

	for i := 0; i < rows; i++ {
		start += cols

		// If next char is a space, we can fit perfectly
		if s[start%n] == ' ' {
			start++
		} else {
			// Backtrack to nearest space
			for start > 0 && s[(start-1)%n] != ' ' {
				start--
			}
		}
	}
	return start / n
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", wordsTyping([]string{"hello", "world"}, 2, 8))
	// Expected: 1

	// Test case 2
	fmt.Println("Test 2:", wordsTyping([]string{"a", "bcd", "e"}, 3, 6))
	// Expected: 2

	// Test case 3
	fmt.Println("Test 3:", wordsTyping([]string{"i", "had", "apple", "pie"}, 4, 5))
	// Expected: 1
}
```

## 0419 — Battleships In A Board

```go
package main

// LeetCode #419: Battleships in a Board
// https://leetcode.com/problems/battleships-in-a-board/
// Difficulty: Medium
// Time: O(m*n) | Space: O(1)

import "fmt"

func countBattleships(board [][]byte) int {
	count := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'X' {
				// Count only if it's the start of a ship (no X above or to the left)
				if (i == 0 || board[i-1][j] != 'X') && (j == 0 || board[i][j-1] != 'X') {
					count++
				}
			}
		}
	}
	return count
}

func main() {
	// Test case 1
	b1 := [][]byte{
		{'X', '.', '.', 'X'},
		{'.', '.', '.', 'X'},
		{'.', '.', '.', 'X'},
	}
	fmt.Println("Test 1:", countBattleships(b1))
	// Expected: 2

	// Test case 2
	b2 := [][]byte{{'X'}}
	fmt.Println("Test 2:", countBattleships(b2))
	// Expected: 1

	// Test case 3: Empty
	b3 := [][]byte{{'.', '.'}, {'.', '.'}}
	fmt.Println("Test 3:", countBattleships(b3))
	// Expected: 0
}
```

## 0421 — Maximum Xor Of Two Numbers In An Array

```go
package main

// LeetCode #421: Maximum XOR of Two Numbers in an Array
// https://leetcode.com/problems/maximum-xor-of-two-numbers-in-an-array/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func findMaximumXOR(nums []int) int {
	maxXor := 0
	mask := 0

	// Try each bit from MSB to LSB
	for i := 31; i >= 0; i-- {
		mask |= 1 << i
		prefixSet := make(map[int]bool)
		for _, num := range nums {
			prefixSet[num&mask] = true
		}

		candidate := maxXor | (1 << i)
		for prefix := range prefixSet {
			if prefixSet[prefix^candidate] {
				maxXor = candidate
				break
			}
		}
	}
	return maxXor
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", findMaximumXOR([]int{3, 10, 5, 25, 2, 8}))
	// Expected: 28

	// Test case 2
	fmt.Println("Test 2:", findMaximumXOR([]int{14, 70, 53, 83, 49, 91, 36, 80, 92, 51, 66, 70}))
	// Expected: 127

	// Test case 3
	fmt.Println("Test 3:", findMaximumXOR([]int{0}))
	// Expected: 0
}
```

## 0423 — Reconstruct Original Digits From English

```go
package main

// LeetCode #423: Reconstruct Original Digits from English
// https://leetcode.com/problems/reconstruct-original-digits-from-english/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import (
	"fmt"
	"strings"
)

func originalDigits(s string) string {
	count := [26]int{}
	for _, ch := range s {
		count[ch-'a']++
	}

	// Unique identifying letters: z(0), w(2), u(4), x(6), g(8)
	digits := make([]int, 10)
	digits[0] = count['z'-'a']
	digits[2] = count['w'-'a']
	digits[4] = count['u'-'a']
	digits[6] = count['x'-'a']
	digits[8] = count['g'-'a']

	// Deduce remaining
	digits[1] = count['o'-'a'] - digits[0] - digits[2] - digits[4]
	digits[3] = count['h'-'a'] - digits[8]
	digits[5] = count['f'-'a'] - digits[4]
	digits[7] = count['s'-'a'] - digits[6]
	digits[9] = count['i'-'a'] - digits[5] - digits[6] - digits[8]

	var sb strings.Builder
	for d := 0; d <= 9; d++ {
		for i := 0; i < digits[d]; i++ {
			sb.WriteByte(byte('0' + d))
		}
	}
	return sb.String()
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", originalDigits("owoztneoer"))
	// Expected: "012"

	// Test case 2
	fmt.Println("Test 2:", originalDigits("fviefuro"))
	// Expected: "45"

	// Test case 3
	fmt.Println("Test 3:", originalDigits("zerozero"))
	// Expected: "00"
}
```

## 0424 — Longest Repeating Character Replacement

```go
package main

// LeetCode #424: Longest Repeating Character Replacement
// https://leetcode.com/problems/longest-repeating-character-replacement/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func characterReplacement(s string, k int) int {
	freq := [26]int{}
	left, maxFreq, maxLen := 0, 0, 0

	for right := 0; right < len(s); right++ {
		freq[s[right]-'A']++
		if freq[s[right]-'A'] > maxFreq {
			maxFreq = freq[s[right]-'A']
		}

		// Window size - maxFreq = chars to replace
		for right-left+1-maxFreq > k {
			freq[s[left]-'A']--
			left++
			// Recompute maxFreq (or keep old - it's safe)
		}

		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}
	return maxLen
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", characterReplacement("ABAB", 2))
	// Expected: 4

	// Test case 2
	fmt.Println("Test 2:", characterReplacement("AABABBA", 1))
	// Expected: 4

	// Test case 3
	fmt.Println("Test 3:", characterReplacement("AAAA", 2))
	// Expected: 4
}
```

## 0426 — Convert Binary Search Tree To Sorted Doubly Linked List

```go
package main

// LeetCode #426: Convert Binary Search Tree to Sorted Doubly Linked List
// https://leetcode.com/problems/convert-binary-search-tree-to-sorted-doubly-linked-list/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(h)

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func treeToDoublyList(root *Node) *Node {
	if root == nil {
		return nil
	}

	var first, last *Node

	var inorder func(node *Node)
	inorder = func(node *Node) {
		if node == nil {
			return
		}
		inorder(node.Left)

		if last != nil {
			last.Right = node
			node.Left = last
		} else {
			first = node
		}
		last = node

		inorder(node.Right)
	}

	inorder(root)

	// Close the circular doubly linked list
	last.Right = first
	first.Left = last
	return first
}

func printList(head *Node) {
	if head == nil {
		fmt.Println("nil")
		return
	}
	cur := head
	for {
		fmt.Print(cur.Val)
		cur = cur.Right
		if cur == head {
			break
		}
		fmt.Print(" <-> ")
	}
	fmt.Println(" (circular)")
}

func main() {
	// Test case 1: [4,2,5,1,3]
	root1 := &Node{Val: 4}
	root1.Left = &Node{Val: 2, Left: &Node{Val: 1}, Right: &Node{Val: 3}}
	root1.Right = &Node{Val: 5}
	fmt.Print("Test 1: ")
	printList(treeToDoublyList(root1))
	// Expected: 1 <-> 2 <-> 3 <-> 4 <-> 5 (circular)

	// Test case 2: Single node
	root2 := &Node{Val: 1}
	fmt.Print("Test 2: ")
	printList(treeToDoublyList(root2))
	// Expected: 1 (circular)

	// Test case 3: Nil
	fmt.Print("Test 3: ")
	printList(treeToDoublyList(nil))
	// Expected: nil
}
```

## 0427 — Construct Quad Tree

```go
package main

// LeetCode #427: Construct Quad Tree
// https://leetcode.com/problems/construct-quad-tree/
// Difficulty: Medium
// Time: O(n^2 log n) | Space: O(log n)

import "fmt"

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func construct(grid [][]int) *Node {
	return build(grid, 0, 0, len(grid))
}

func build(grid [][]int, r, c, size int) *Node {
	if size == 1 {
		return &Node{Val: grid[r][c] == 1, IsLeaf: true}
	}

	half := size / 2
	topLeft := build(grid, r, c, half)
	topRight := build(grid, r, c+half, half)
	bottomLeft := build(grid, r+half, c, half)
	bottomRight := build(grid, r+half, c+half, half)

	// Check if all children are leaves with same value
	if topLeft.IsLeaf && topRight.IsLeaf && bottomLeft.IsLeaf && bottomRight.IsLeaf &&
		topLeft.Val == topRight.Val && topRight.Val == bottomLeft.Val && bottomLeft.Val == bottomRight.Val {
		return &Node{Val: topLeft.Val, IsLeaf: true}
	}

	return &Node{
		IsLeaf:      false,
		Val:         true,
		TopLeft:     topLeft,
		TopRight:    topRight,
		BottomLeft:  bottomLeft,
		BottomRight: bottomRight,
	}
}

func printQuadTree(node *Node, indent string) {
	if node == nil {
		return
	}
	if node.IsLeaf {
		fmt.Printf("%sLeaf: %v\n", indent, node.Val)
	} else {
		fmt.Printf("%sInternal:\n", indent)
		printQuadTree(node.TopLeft, indent+"  TL: ")
		printQuadTree(node.TopRight, indent+"  TR: ")
		printQuadTree(node.BottomLeft, indent+"  BL: ")
		printQuadTree(node.BottomRight, indent+"  BR: ")
	}
}

func main() {
	// Test case 1
	grid1 := [][]int{
		{0, 1},
		{1, 0},
	}
	fmt.Println("Test 1:")
	n1 := construct(grid1)
	fmt.Println("  IsLeaf:", n1.IsLeaf) // false (quadtree)

	// Test case 2: All same value
	grid2 := [][]int{
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
	}
	n2 := construct(grid2)
	fmt.Println("Test 2: IsLeaf:", n2.IsLeaf, "Val:", n2.Val)
	// Expected: true, 1

	// Test case 3
	grid3 := [][]int{
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
	}
	n3 := construct(grid3)
	fmt.Println("Test 3: IsLeaf:", n3.IsLeaf)
	// Expected: false
}
```

## 0429 — N Ary Tree Level Order Traversal

```go
package main

// LeetCode #429: N-ary Tree Level Order Traversal
// https://leetcode.com/problems/n-ary-tree-level-order-traversal/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}

	queue := []*Node{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		level := make([]int, 0, levelSize)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			queue = append(queue, node.Children...)
		}
		result = append(result, level)
	}
	return result
}

func main() {
	// Test case 1: [1,null,3,2,4,null,5,6]
	root1 := &Node{Val: 1}
	root1.Children = []*Node{
		{Val: 3, Children: []*Node{{Val: 5}, {Val: 6}}},
		{Val: 2},
		{Val: 4},
	}
	fmt.Println("Test 1:", levelOrder(root1))
	// Expected: [[1],[3,2,4],[5,6]]

	// Test case 2: Single node
	root2 := &Node{Val: 1}
	fmt.Println("Test 2:", levelOrder(root2))
	// Expected: [[1]]

	// Test case 3: Nil
	fmt.Println("Test 3:", levelOrder(nil))
	// Expected: []
}
```

## 0430 — Flatten A Multilevel Doubly Linked List

```go
package main

// LeetCode #430: Flatten a Multilevel Doubly Linked List
// https://leetcode.com/problems/flatten-a-multilevel-doubly-linked-list/
// Difficulty: Medium
// Time: O(n) | Space: O(d)

import "fmt"

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	flattenDFS(root)
	return root
}

// Returns the tail of the flattened list
func flattenDFS(node *Node) *Node {
	current := node
	var last *Node

	for current != nil {
		next := current.Next

		if current.Child != nil {
			childTail := flattenDFS(current.Child)

			// Connect current to child
			current.Next = current.Child
			current.Child.Prev = current

			// Connect child tail to next
			if next != nil {
				childTail.Next = next
				next.Prev = childTail
			}

			current.Child = nil
			last = childTail
		} else {
			last = current
		}
		current = next
	}
	return last
}

func printList(head *Node) {
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print(" <-> ")
		}
		head = head.Next
	}
	fmt.Println()
}

func main() {
	// Build: 1 - 2 - 3 - 4 - 5 - 6
	//            |
	//            7 - 8 - 9 - 10
	//                |
	//               11 - 12
	n1 := &Node{Val: 1}
	n2 := &Node{Val: 2}
	n3 := &Node{Val: 3}
	n4 := &Node{Val: 4}
	n5 := &Node{Val: 5}
	n6 := &Node{Val: 6}
	n7 := &Node{Val: 7}
	n8 := &Node{Val: 8}
	n9 := &Node{Val: 9}
	n10 := &Node{Val: 10}
	n11 := &Node{Val: 11}
	n12 := &Node{Val: 12}

	n1.Next = n2
	n2.Prev = n1
	n2.Next = n3
	n3.Prev = n2
	n3.Next = n4
	n4.Prev = n3
	n4.Next = n5
	n5.Prev = n4
	n5.Next = n6
	n6.Prev = n5

	n3.Child = n7
	n7.Next = n8
	n8.Prev = n7
	n8.Next = n9
	n9.Prev = n8
	n9.Next = n10
	n10.Prev = n9

	n8.Child = n11
	n11.Next = n12
	n12.Prev = n11

	fmt.Print("Test 1: ")
	printList(flatten(n1))
	// Expected: 1 <-> 2 <-> 3 <-> 7 <-> 8 <-> 11 <-> 12 <-> 9 <-> 10 <-> 4 <-> 5 <-> 6

	// Test case 2: Nil
	fmt.Print("Test 2: ")
	printList(flatten(nil))

	// Test case 3: No children
	n21 := &Node{Val: 1}
	n22 := &Node{Val: 2}
	n21.Next = n22
	n22.Prev = n21
	fmt.Print("Test 3: ")
	printList(flatten(n21))
	// Expected: 1 <-> 2
}
```

## 0433 — Minimum Genetic Mutation

```go
package main

// LeetCode #433: Minimum Genetic Mutation
// https://leetcode.com/problems/minimum-genetic-mutation/
// Difficulty: Medium
// Time: O(4^10) worst case (BFS) | Space: O(n)

import "fmt"

func minMutation(startGene string, endGene string, bank []string) int {
	bankSet := make(map[string]bool)
	for _, b := range bank {
		bankSet[b] = true
	}

	if !bankSet[endGene] {
		return -1
	}

	genes := []byte{'A', 'C', 'G', 'T'}
	queue := []string{startGene}
	visited := map[string]bool{startGene: true}
	steps := 0

	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			curr := queue[0]
			queue = queue[1:]

			if curr == endGene {
				return steps
			}

			b := []byte(curr)
			for j := 0; j < 8; j++ {
				orig := b[j]
				for _, g := range genes {
					if g == orig {
						continue
					}
					b[j] = g
					next := string(b)
					if bankSet[next] && !visited[next] {
						visited[next] = true
						queue = append(queue, next)
					}
				}
				b[j] = orig
			}
		}
		steps++
	}
	return -1
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minMutation("AACCGGTT", "AACCGGTA", []string{"AACCGGTA"}))
	// Expected: 1

	// Test case 2
	fmt.Println("Test 2:", minMutation("AACCGGTT", "AAACGGTA", []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}))
	// Expected: 2

	// Test case 3: Not in bank
	fmt.Println("Test 3:", minMutation("AAAAACCC", "AACCCCCC", []string{"AAAACCCC", "AAACCCCC", "AACCCCCC"}))
	// Expected: 3
}
```

## 0435 — Non Overlapping Intervals

```go
package main

// LeetCode #435: Non-overlapping Intervals
// https://leetcode.com/problems/non-overlapping-intervals/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	// Sort by end
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	count := 0
	end := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < end {
			count++
		} else {
			end = intervals[i][1]
		}
	}
	return count
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", eraseOverlapIntervals([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}))
	// Expected: 1

	// Test case 2
	fmt.Println("Test 2:", eraseOverlapIntervals([][]int{{1, 2}, {1, 2}, {1, 2}}))
	// Expected: 2

	// Test case 3
	fmt.Println("Test 3:", eraseOverlapIntervals([][]int{{1, 2}, {2, 3}}))
	// Expected: 0
}
```

## 0436 — Find Right Interval

```go
package main

// LeetCode #436: Find Right Interval
// https://leetcode.com/problems/find-right-interval/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func findRightInterval(intervals [][]int) []int {
	n := len(intervals)
	// Create array of (start, index) pairs
	starts := make([][2]int, n)
	for i, iv := range intervals {
		starts[i] = [2]int{iv[0], i}
	}
	sort.Slice(starts, func(i, j int) bool {
		return starts[i][0] < starts[j][0]
	})

	result := make([]int, n)
	for i, iv := range intervals {
		target := iv[1]
		// Binary search
		idx := sort.Search(n, func(j int) bool {
			return starts[j][0] >= target
		})
		if idx < n {
			result[i] = starts[idx][1]
		} else {
			result[i] = -1
		}
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", findRightInterval([][]int{{1, 2}}))
	// Expected: [-1]

	// Test case 2
	fmt.Println("Test 2:", findRightInterval([][]int{{3, 4}, {2, 3}, {1, 2}}))
	// Expected: [-1, 0, 1]

	// Test case 3
	fmt.Println("Test 3:", findRightInterval([][]int{{1, 4}, {2, 3}, {3, 4}}))
	// Expected: [-1, 2, -1]
}
```

## 0437 — Path Sum Iii

```go
package main

// LeetCode #437: Path Sum III
// https://leetcode.com/problems/path-sum-iii/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	// prefix sum -> count
	prefixSum := map[int]int{0: 1}
	return dfs(root, targetSum, 0, prefixSum)
}

func dfs(node *TreeNode, targetSum, curSum int, prefixSum map[int]int) int {
	if node == nil {
		return 0
	}

	curSum += node.Val
	count := prefixSum[curSum-targetSum]

	prefixSum[curSum]++
	count += dfs(node.Left, targetSum, curSum, prefixSum)
	count += dfs(node.Right, targetSum, curSum, prefixSum)
	prefixSum[curSum]--

	return count
}

func main() {
	// Test case 1
	root1 := &TreeNode{Val: 10}
	root1.Left = &TreeNode{Val: 5, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: -2}}, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 1}}}
	root1.Right = &TreeNode{Val: -3, Right: &TreeNode{Val: 11}}
	fmt.Println("Test 1:", pathSum(root1, 8))
	// Expected: 3

	// Test case 2: Single node
	root2 := &TreeNode{Val: 1}
	fmt.Println("Test 2:", pathSum(root2, 1))
	// Expected: 1

	// Test case 3: Nil
	fmt.Println("Test 3:", pathSum(nil, 0))
	// Expected: 0
}
```

## 0438 — Find All Anagrams In A String

```go
package main

// LeetCode #438: Find All Anagrams in a String
// https://leetcode.com/problems/find-all-anagrams-in-a-string/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}

	pFreq := [26]int{}
	for _, ch := range p {
		pFreq[ch-'a']++
	}

	sFreq := [26]int{}
	result := []int{}

	// Initialize first window
	for i := 0; i < len(p); i++ {
		sFreq[s[i]-'a']++
	}
	if sFreq == pFreq {
		result = append(result, 0)
	}

	// Slide window
	for i := len(p); i < len(s); i++ {
		sFreq[s[i]-'a']++
		sFreq[s[i-len(p)]-'a']--
		if sFreq == pFreq {
			result = append(result, i-len(p)+1)
		}
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", findAnagrams("cbaebabacd", "abc"))
	// Expected: [0, 6]

	// Test case 2
	fmt.Println("Test 2:", findAnagrams("abab", "ab"))
	// Expected: [0, 1, 2]

	// Test case 3
	fmt.Println("Test 3:", findAnagrams("a", "ab"))
	// Expected: []
}
```

## 0439 — Ternary Expression Parser

```go
package main

// LeetCode #439: Ternary Expression Parser
// https://leetcode.com/problems/ternary-expression-parser/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func parseTernary(expression string) string {
	stack := make([]byte, 0)
	// Process from right to left
	for i := len(expression) - 1; i >= 0; i-- {
		ch := expression[i]
		if ch >= '0' && ch <= '9' || ch == 'T' || ch == 'F' {
			stack = append(stack, ch)
		} else if ch == '?' {
			// Evaluate: preceding char is condition
			trueVal := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			falseVal := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			cond := expression[i-1]
			i-- // skip the condition character (already consumed)
			if cond == 'T' {
				stack = append(stack, trueVal)
			} else {
				stack = append(stack, falseVal)
			}
		}
		// Skip ':'
	}
	return string(stack[0])
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", parseTernary("T?2:3"))
	// Expected: "2"

	// Test case 2
	fmt.Println("Test 2:", parseTernary("F?1:T?4:5"))
	// Expected: "4"

	// Test case 3
	fmt.Println("Test 3:", parseTernary("T?T?F:5:3"))
	// Expected: "F"
}
```

## 0442 — Find All Duplicates In An Array

```go
package main

// LeetCode #442: Find All Duplicates in an Array
// https://leetcode.com/problems/find-all-duplicates-in-an-array/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func findDuplicates(nums []int) []int {
	result := []int{}
	for _, num := range nums {
		idx := num
		if idx < 0 {
			idx = -idx
		}
		idx--
		if nums[idx] < 0 {
			result = append(result, idx+1)
		} else {
			nums[idx] = -nums[idx]
		}
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	// Expected: [2, 3]

	// Test case 2
	fmt.Println("Test 2:", findDuplicates([]int{1, 1, 2}))
	// Expected: [1]

	// Test case 3
	fmt.Println("Test 3:", findDuplicates([]int{1}))
	// Expected: []
}
```

## 0443 — String Compression

```go
package main

// LeetCode #443: String Compression
// https://leetcode.com/problems/string-compression/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {
	write := 0
	read := 0

	for read < len(chars) {
		ch := chars[read]
		count := 0

		// Count consecutive chars
		for read < len(chars) && chars[read] == ch {
			read++
			count++
		}

		// Write char
		chars[write] = ch
		write++

		// Write count
		if count > 1 {
			for _, d := range strconv.Itoa(count) {
				chars[write] = byte(d)
				write++
			}
		}
	}
	return write
}

func main() {
	// Test case 1
	chars1 := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	l1 := compress(chars1)
	fmt.Println("Test 1:", l1, string(chars1[:l1]))
	// Expected: 6, "a2b2c3"

	// Test case 2
	chars2 := []byte{'a'}
	l2 := compress(chars2)
	fmt.Println("Test 2:", l2, string(chars2[:l2]))
	// Expected: 1, "a"

	// Test case 3
	chars3 := []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}
	l3 := compress(chars3)
	fmt.Println("Test 3:", l3, string(chars3[:l3]))
	// Expected: 4, "ab12"
}
```

## 0444 — Sequence Reconstruction

```go
package main

// LeetCode #444: Sequence Reconstruction
// https://leetcode.com/problems/sequence-reconstruction/
// Difficulty: Medium [Paid]
// Time: O(n + m) | Space: O(n)

import "fmt"

func sequenceReconstruction(org []int, seqs [][]int) bool {
	n := len(org)
	// Build indegree map and edges
	indegree := make([]int, n+1)
	graph := make([][]int, n+1)
	exists := make([]bool, n+1)

	for _, seq := range seqs {
		for _, num := range seq {
			if num < 1 || num > n {
				return false
			}
			exists[num] = true
		}
		for i := 0; i < len(seq)-1; i++ {
			u, v := seq[i], seq[i+1]
			graph[u] = append(graph[u], v)
			indegree[v]++
		}
	}

	// Check all numbers exist
	for i := 1; i <= n; i++ {
		if !exists[i] {
			return false
		}
	}

	// BFS: only one node with indegree 0 at each step
	queue := []int{}
	for i := 1; i <= n; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	idx := 0
	for len(queue) == 1 {
		u := queue[0]
		queue = queue[1:]
		if u != org[idx] {
			return false
		}
		idx++
		for _, v := range graph[u] {
			indegree[v]--
			if indegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	return idx == n
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", sequenceReconstruction([]int{1, 2, 3}, [][]int{{1, 2}, {1, 3}, {2, 3}}))
	// Expected: true

	// Test case 2
	fmt.Println("Test 2:", sequenceReconstruction([]int{1, 2, 3}, [][]int{{1, 2}, {1, 3}}))
	// Expected: false

	// Test case 3
	fmt.Println("Test 3:", sequenceReconstruction([]int{1, 2, 3}, [][]int{{1, 2}, {2, 3}, {3, 1}}))
	// Expected: false (cycle)
}
```

## 0445 — Add Two Numbers Ii

```go
package main

// LeetCode #445: Add Two Numbers II
// https://leetcode.com/problems/add-two-numbers-ii/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s1 := []int{}
	s2 := []int{}

	for l1 != nil {
		s1 = append(s1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		s2 = append(s2, l2.Val)
		l2 = l2.Next
	}

	var head *ListNode
	carry := 0

	for len(s1) > 0 || len(s2) > 0 || carry > 0 {
		sum := carry
		if len(s1) > 0 {
			sum += s1[len(s1)-1]
			s1 = s1[:len(s1)-1]
		}
		if len(s2) > 0 {
			sum += s2[len(s2)-1]
			s2 = s2[:len(s2)-1]
		}

		node := &ListNode{Val: sum % 10, Next: head}
		head = node
		carry = sum / 10
	}

	return head
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print(" -> ")
		}
		head = head.Next
	}
	fmt.Println()
}

func main() {
	// Test case 1: 7->2->4->3 + 5->6->4 = 7->8->0->7
	l1 := &ListNode{7, &ListNode{2, &ListNode{4, &ListNode{3, nil}}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	fmt.Print("Test 1: ")
	printList(addTwoNumbers(l1, l2))
	// Expected: 7 -> 8 -> 0 -> 7

	// Test case 2: 5 + 5 = 1->0
	l3 := &ListNode{5, nil}
	l4 := &ListNode{5, nil}
	fmt.Print("Test 2: ")
	printList(addTwoNumbers(l3, l4))
	// Expected: 1 -> 0

	// Test case 3: 0 + 0 = 0
	l5 := &ListNode{0, nil}
	l6 := &ListNode{0, nil}
	fmt.Print("Test 3: ")
	printList(addTwoNumbers(l5, l6))
	// Expected: 0
}
```

## 0447 — Number Of Boomerangs

```go
package main

// LeetCode #447: Number of Boomerangs
// https://leetcode.com/problems/number-of-boomerangs/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import "fmt"

func numberOfBoomerangs(points [][]int) int {
	total := 0

	for i := 0; i < len(points); i++ {
		distCount := make(map[int]int)
		for j := 0; j < len(points); j++ {
			if i == j {
				continue
			}
			dx := points[i][0] - points[j][0]
			dy := points[i][1] - points[j][1]
			dist := dx*dx + dy*dy
			distCount[dist]++
		}
		for _, count := range distCount {
			total += count * (count - 1)
		}
	}
	return total
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", numberOfBoomerangs([][]int{{0, 0}, {1, 0}, {2, 0}}))
	// Expected: 2

	// Test case 2
	fmt.Println("Test 2:", numberOfBoomerangs([][]int{{1, 1}, {2, 2}, {3, 3}}))
	// Expected: 2

	// Test case 3
	fmt.Println("Test 3:", numberOfBoomerangs([][]int{{0, 0}}))
	// Expected: 0
}
```

## 0449 — Serialize And Deserialize Bst

```go
package main

// LeetCode #449: Serialize and Deserialize BST
// https://leetcode.com/problems/serialize-and-deserialize-bst/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct{}

func Constructor() Codec {
	return Codec{}
}

// Serialize using preorder traversal
func (c *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	result := []string{}
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		result = append(result, strconv.Itoa(node.Val))
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return strings.Join(result, ",")
}

// Deserialize using preorder + BST property
func (c *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}

	vals := strings.Split(data, ",")
	nums := make([]int, len(vals))
	for i, v := range vals {
		nums[i], _ = strconv.Atoi(v)
	}

	idx := 0
	var build func(lower, upper int) *TreeNode
	build = func(lower, upper int) *TreeNode {
		if idx >= len(nums) {
			return nil
		}
		val := nums[idx]
		if val < lower || val > upper {
			return nil
		}
		idx++
		return &TreeNode{
			Val:   val,
			Left:  build(lower, val),
			Right: build(val, upper),
		}
	}
	return build(-1<<31, 1<<31-1)
}

func inorderPrint(node *TreeNode) {
	if node == nil {
		return
	}
	inorderPrint(node.Left)
	fmt.Print(node.Val, " ")
	inorderPrint(node.Right)
}

func main() {
	c := Constructor()

	// Test case 1
	root1 := &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}
	data1 := c.serialize(root1)
	fmt.Println("Test 1 serialized:", data1)
	node1 := c.deserialize(data1)
	fmt.Print("Test 1 deserialized (inorder): ")
	inorderPrint(node1)
	fmt.Println()
	// Expected: "2,1,3" and inorder: 1 2 3

	// Test case 2: Nil
	data2 := c.serialize(nil)
	fmt.Println("Test 2 serialized:", data2)
	node2 := c.deserialize(data2)
	fmt.Print("Test 2 deserialized: ")
	inorderPrint(node2)
	fmt.Println()
	// Expected: "" and nothing

	// Test case 3
	root3 := &TreeNode{Val: 5, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 7, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 8}}}
	data3 := c.serialize(root3)
	fmt.Println("Test 3 serialized:", data3)
	node3 := c.deserialize(data3)
	fmt.Print("Test 3 deserialized (inorder): ")
	inorderPrint(node3)
	fmt.Println()
}
```

## 0450 — Delete Node In A Bst

```go
package main

// LeetCode #450: Delete Node in a BST
// https://leetcode.com/problems/delete-node-in-a-bst/
// Difficulty: Medium
// Time: O(h) | Space: O(h)

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else {
		// Node to delete found
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}

		// Find inorder successor (leftmost in right subtree)
		successor := root.Right
		for successor.Left != nil {
			successor = successor.Left
		}
		root.Val = successor.Val
		root.Right = deleteNode(root.Right, successor.Val)
	}
	return root
}

func inorderPrint(node *TreeNode) {
	if node == nil {
		return
	}
	inorderPrint(node.Left)
	fmt.Print(node.Val, " ")
	inorderPrint(node.Right)
}

func main() {
	// Test case 1: [5,3,6,2,4,null,7], key=3
	root1 := &TreeNode{Val: 5}
	root1.Left = &TreeNode{Val: 3, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 4}}
	root1.Right = &TreeNode{Val: 6, Right: &TreeNode{Val: 7}}
	fmt.Print("Test 1 before: ")
	inorderPrint(root1)
	fmt.Println()
	root1 = deleteNode(root1, 3)
	fmt.Print("Test 1 after: ")
	inorderPrint(root1)
	fmt.Println()
	// Expected inorder: 2 4 5 6 7

	// Test case 2: Key not found
	root2 := &TreeNode{Val: 1}
	root2 = deleteNode(root2, 2)
	fmt.Print("Test 2: ")
	inorderPrint(root2)
	fmt.Println()
	// Expected: 1

	// Test case 3: Delete leaf
	root3 := &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 2}}
	root3 = deleteNode(root3, 2)
	fmt.Print("Test 3: ")
	inorderPrint(root3)
	fmt.Println()
	// Expected: 0 1
}
```

## 0451 — Sort Characters By Frequency

```go
package main

// LeetCode #451: Sort Characters By Frequency
// https://leetcode.com/problems/sort-characters-by-frequency/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"strings"
)

func frequencySort(s string) string {
	freq := make([]int, 128)
	for _, ch := range s {
		freq[ch]++
	}

	// Bucket sort: index = frequency
	buckets := make([][]byte, len(s)+1)
	for ch := 0; ch < 128; ch++ {
		if freq[ch] > 0 {
			buckets[freq[ch]] = append(buckets[freq[ch]], byte(ch))
		}
	}

	var sb strings.Builder
	for count := len(buckets) - 1; count > 0; count-- {
		for _, ch := range buckets[count] {
			sb.WriteString(strings.Repeat(string(ch), count))
		}
	}
	return sb.String()
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", frequencySort("tree"))
	// Expected: "eert" or "eetr"

	// Test case 2
	fmt.Println("Test 2:", frequencySort("cccaaa"))
	// Expected: "aaaccc" or "cccaaa"

	// Test case 3
	fmt.Println("Test 3:", frequencySort("Aabb"))
	// Expected: "bbAa" or "bbaA"
}
```

## 0452 — Minimum Number Of Arrows To Burst Balloons

```go
package main

// LeetCode #452: Minimum Number of Arrows to Burst Balloons
// https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	arrows := 1
	end := points[0][1]

	for _, p := range points[1:] {
		if p[0] > end {
			arrows++
			end = p[1]
		}
	}
	return arrows
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", findMinArrowShots([][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}))
	// Expected: 2

	// Test case 2
	fmt.Println("Test 2:", findMinArrowShots([][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}))
	// Expected: 4

	// Test case 3
	fmt.Println("Test 3:", findMinArrowShots([][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}))
	// Expected: 2
}
```

## 0453 — Minimum Moves To Equal Array Elements

```go
package main

// LeetCode #453: Minimum Moves to Equal Array Elements
// https://leetcode.com/problems/minimum-moves-to-equal-array-elements/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import (
	"fmt"
	"math"
)

func minMoves(nums []int) int {
	minVal := math.MaxInt32
	sum := 0
	for _, n := range nums {
		sum += n
		if n < minVal {
			minVal = n
		}
	}
	return sum - minVal*len(nums)
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minMoves([]int{1, 2, 3}))
	// Expected: 3

	// Test case 2
	fmt.Println("Test 2:", minMoves([]int{1, 1, 1}))
	// Expected: 0

	// Test case 3
	fmt.Println("Test 3:", minMoves([]int{1, 1000000000}))
	// Expected: 999999999
}
```

## 0454 — 4Sum Ii

```go
package main

// LeetCode #454: 4Sum II
// https://leetcode.com/problems/4sum-ii/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n^2)

import "fmt"

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	sumMap := make(map[int]int)
	for _, a := range nums1 {
		for _, b := range nums2 {
			sumMap[a+b]++
		}
	}

	count := 0
	for _, c := range nums3 {
		for _, d := range nums4 {
			count += sumMap[-(c + d)]
		}
	}
	return count
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", fourSumCount([]int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}))
	// Expected: 2

	// Test case 2
	fmt.Println("Test 2:", fourSumCount([]int{0}, []int{0}, []int{0}, []int{0}))
	// Expected: 1

	// Test case 3
	fmt.Println("Test 3:", fourSumCount([]int{-1, -1}, []int{-1, 1}, []int{-1, 1}, []int{1, -1}))
	// Expected: 6
}
```

## 0456 — 132 Pattern

```go
package main

// LeetCode #456: 132 Pattern
// https://leetcode.com/problems/132-pattern/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func find132pattern(nums []int) bool {
	stack := []int{}
	s3 := -1 << 31 // nums[k] (the "2" in 132)

	// Iterate from right to left
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < s3 {
			return true
		}
		// nums[i] > stack top means nums[i] could be the "3"
		for len(stack) > 0 && nums[i] > stack[len(stack)-1] {
			s3 = stack[len(stack)-1] // s3 is the current largest "2" candidate
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums[i])
	}
	return false
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", find132pattern([]int{1, 2, 3, 4}))
	// Expected: false

	// Test case 2
	fmt.Println("Test 2:", find132pattern([]int{3, 1, 4, 2}))
	// Expected: true

	// Test case 3
	fmt.Println("Test 3:", find132pattern([]int{-1, 3, 2, 0}))
	// Expected: true
}
```

## 0457 — Circular Array Loop

```go
package main

// LeetCode #457: Circular Array Loop
// https://leetcode.com/problems/circular-array-loop/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func circularArrayLoop(nums []int) bool {
	n := len(nums)

	next := func(i int) int {
		return ((i+nums[i])%n + n) % n
	}

	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			continue
		}

		slow, fast := i, i
		for {
			slow = next(slow)
			fast = next(next(fast))
			if slow == fast {
				break
			}
		}

		// Found cycle, check direction and length
		if next(slow) != slow && nums[slow]*nums[next(slow)] > 0 {
			length := 1
			cur := next(slow)
			for cur != slow {
				length++
				cur = next(cur)
			}
			if length > 1 {
				return true
			}
		}

		// Mark visited
		cur := i
		for nums[cur]*nums[next(cur)] > 0 {
			tmp := next(cur)
			nums[cur] = 0
			cur = tmp
		}
	}
	return false
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", circularArrayLoop([]int{2, -1, 1, 2, 2}))
	// Expected: true

	// Test case 2
	fmt.Println("Test 2:", circularArrayLoop([]int{-1, -2, -3, -4, -5, 6}))
	// Expected: false

	// Test case 3
	fmt.Println("Test 3:", circularArrayLoop([]int{1, -1, 2}))
	// Expected: true
}
```

## 0462 — Minimum Moves To Equal Array Elements Ii

```go
package main

// LeetCode #462: Minimum Moves to Equal Array Elements II
// https://leetcode.com/problems/minimum-moves-to-equal-array-elements-ii/
// Difficulty: Medium
// Time: O(n log n) for sorting, O(n) for QuickSelect
// Space: O(log n) for sorting, O(1) for QuickSelect

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MinimumMovesToEqualArrayElementsIi([]int{1, 2, 3}))
	fmt.Println(MinimumMovesToEqualArrayElementsIi([]int{1, 10, 2, 9}))
	fmt.Println(MinimumMovesToEqualArrayElementsIi([]int{1, 0, 0, 8, 6}))
}

func MinimumMovesToEqualArrayElementsIi(nums []int) int {
	sort.Ints(nums)
	median := nums[len(nums)/2]
	moves := 0
	for _, num := range nums {
		diff := num - median
		if diff < 0 {
			diff = -diff
		}
		moves += diff
	}
	return moves
}
```

## 0464 — Can I Win

```go
package main

// LeetCode #464: Can I Win
// https://leetcode.com/problems/can-i-win/
// Difficulty: Medium
// Time: O(2^n) where n = maxChoosableInteger (max 20)
// Space: O(2^n)

import "fmt"

func main() {
	fmt.Println(CanIWin(10, 11))
	fmt.Println(CanIWin(10, 0))
	fmt.Println(CanIWin(10, 40))
}

func CanIWin(maxChoosableInteger int, desiredTotal int) bool {
	if desiredTotal <= 0 {
		return true
	}
	sum := maxChoosableInteger * (maxChoosableInteger + 1) / 2
	if sum < desiredTotal {
		return false
	}

	memo := make(map[int]bool)
	var dfs func(used int, currentTotal int) bool
	dfs = func(used int, currentTotal int) bool {
		if currentTotal >= desiredTotal {
			return false
		}
		if val, ok := memo[used]; ok {
			return val
		}
		for i := 1; i <= maxChoosableInteger; i++ {
			mask := 1 << (i - 1)
			if used&mask == 0 {
				if !dfs(used|mask, currentTotal+i) {
					memo[used] = true
					return true
				}
			}
		}
		memo[used] = false
		return false
	}

	return dfs(0, 0)
}
```

## 0467 — Unique Substrings In Wraparound String

```go
package main

// LeetCode #467: Unique Substrings in Wraparound String
// https://leetcode.com/problems/unique-substrings-in-wraparound-string/
// Difficulty: Medium
// Time: O(n)
// Space: O(1) (26 letters)

import "fmt"

func main() {
	fmt.Println(UniqueSubstringsInWraparoundString("a"))
	fmt.Println(UniqueSubstringsInWraparoundString("cac"))
	fmt.Println(UniqueSubstringsInWraparoundString("zab"))
}

func UniqueSubstringsInWraparoundString(s string) int {
	maxLen := make([]int, 26)
	curLen := 0

	for i := 0; i < len(s); i++ {
		if i > 0 && (s[i]-s[i-1] == 1 || (s[i-1] == 'z' && s[i] == 'a')) {
			curLen++
		} else {
			curLen = 1
		}
		idx := s[i] - 'a'
		if curLen > maxLen[idx] {
			maxLen[idx] = curLen
		}
	}

	total := 0
	for _, v := range maxLen {
		total += v
	}
	return total
}
```

## 0468 — Validate Ip Address

```go
package main

// LeetCode #468: Validate IP Address
// https://leetcode.com/problems/validate-ip-address/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(ValidateIpAddress("172.16.254.1"))
	fmt.Println(ValidateIpAddress("2001:0db8:85a3:0:0:8A2E:0370:7334"))
	fmt.Println(ValidateIpAddress("256.256.256.256"))
}

func ValidateIpAddress(queryIP string) string {
	if isIPv4(queryIP) {
		return "IPv4"
	}
	if isIPv6(queryIP) {
		return "IPv6"
	}
	return "Neither"
}

func isIPv4(s string) bool {
	parts := strings.Split(s, ".")
	if len(parts) != 4 {
		return false
	}
	for _, p := range parts {
		if len(p) == 0 || (len(p) > 1 && p[0] == '0') {
			return false
		}
		num, err := strconv.Atoi(p)
		if err != nil || num < 0 || num > 255 {
			return false
		}
	}
	return true
}

func isIPv6(s string) bool {
	parts := strings.Split(s, ":")
	if len(parts) != 8 {
		return false
	}
	for _, p := range parts {
		if len(p) == 0 || len(p) > 4 {
			return false
		}
		for _, c := range p {
			if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')) {
				return false
			}
		}
	}
	return true
}
```

## 0469 — Convex Polygon

```go
package main

// LeetCode #469: Convex Polygon
// https://leetcode.com/problems/convex-polygon/
// Difficulty: Medium [Paid]
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(ConvexPolygon([][]int{{0, 0}, {0, 1}, {1, 1}, {1, 0}}))
	fmt.Println(ConvexPolygon([][]int{{0, 0}, {0, 10}, {10, 10}, {10, 0}, {5, 5}}))
}

func ConvexPolygon(polygon [][]int) bool {
	n := len(polygon)
	if n < 3 {
		return false
	}

	var prevCross int
	first := true

	for i := 0; i < n; i++ {
		a, b, c := polygon[i], polygon[(i+1)%n], polygon[(i+2)%n]
		cross := (b[0]-a[0])*(c[1]-b[1]) - (b[1]-a[1])*(c[0]-b[0])
		if cross != 0 {
			if first {
				prevCross = cross
				first = false
			} else if (cross > 0 && prevCross < 0) || (cross < 0 && prevCross > 0) {
				return false
			}
		}
	}
	return true
}
```

## 0470 — Implement Rand10 Using Rand7

```go
package main

// LeetCode #470: Implement Rand10() Using Rand7()
// https://leetcode.com/problems/implement-rand10-using-rand7/
// Difficulty: Medium
// Time: O(1) expected
// Space: O(1)

import (
	"fmt"
	"math/rand"
)

func main() {
	// Test by generating a few random values
	for i := 0; i < 5; i++ {
		fmt.Println(ImplementRandOneZeroUsingRandSeven())
	}
}

func ImplementRandOneZeroUsingRandSeven() int {
	// Rejection sampling: (rand7()-1)*7 + rand7() gives 1-49 uniformly
	// Accept values 1-40 and map to 1-10
	for {
		val := (rand7()-1)*7 + rand7()
		if val <= 40 {
			return (val-1)%10 + 1
		}
	}
}

func rand7() int {
	return rand.Intn(7) + 1
}
```

## 0473 — Matchsticks To Square

```go
package main

// LeetCode #473: Matchsticks to Square
// https://leetcode.com/problems/matchsticks-to-square/
// Difficulty: Medium
// Time: O(4^n) worst case, but pruning makes it much faster
// Space: O(n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MatchsticksToSquare([]int{1, 1, 2, 2, 2}))
	fmt.Println(MatchsticksToSquare([]int{3, 3, 3, 3, 4}))
}

func MatchsticksToSquare(matchsticks []int) bool {
	sum := 0
	for _, m := range matchsticks {
		sum += m
	}
	if sum%4 != 0 {
		return false
	}
	target := sum / 4

	// Sort descending for better pruning
	sort.Sort(sort.Reverse(sort.IntSlice(matchsticks)))

	sides := make([]int, 4)
	var dfs func(idx int) bool
	dfs = func(idx int) bool {
		if idx == len(matchsticks) {
			return sides[0] == target && sides[1] == target && sides[2] == target
		}
		for i := 0; i < 4; i++ {
			if sides[i]+matchsticks[idx] > target {
				continue
			}
			// Optimization: skip duplicate side lengths
			if i > 0 && sides[i] == sides[i-1] {
				continue
			}
			sides[i] += matchsticks[idx]
			if dfs(idx + 1) {
				return true
			}
			sides[i] -= matchsticks[idx]
		}
		return false
	}

	return dfs(0)
}
```

## 0474 — Ones And Zeroes

```go
package main

// LeetCode #474: Ones and Zeroes
// https://leetcode.com/problems/ones-and-zeroes/
// Difficulty: Medium
// Time: O(m * n * len(strs))
// Space: O(m * n)

import "fmt"

func main() {
	fmt.Println(OnesAndZeroes([]string{"10", "0001", "111001", "1", "0"}, 5, 3))
	fmt.Println(OnesAndZeroes([]string{"10", "0", "1"}, 1, 1))
}

func OnesAndZeroes(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for _, s := range strs {
		zeros, ones := countBits(s)
		for i := m; i >= zeros; i-- {
			for j := n; j >= ones; j-- {
				if dp[i-zeros][j-ones]+1 > dp[i][j] {
					dp[i][j] = dp[i-zeros][j-ones] + 1
				}
			}
		}
	}

	return dp[m][n]
}

func countBits(s string) (int, int) {
	zeros, ones := 0, 0
	for _, c := range s {
		if c == '0' {
			zeros++
		} else {
			ones++
		}
	}
	return zeros, ones
}
```

## 0475 — Heaters

```go
package main

// LeetCode #475: Heaters
// https://leetcode.com/problems/heaters/
// Difficulty: Medium
// Time: O(n log n + m log n) where n = len(heaters), m = len(houses)
// Space: O(log n) for sorting

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(Heaters([]int{1, 2, 3}, []int{2}))
	fmt.Println(Heaters([]int{1, 2, 3, 4}, []int{1, 4}))
	fmt.Println(Heaters([]int{1, 5}, []int{2}))
}

func Heaters(houses []int, heaters []int) int {
	sort.Ints(heaters)
	maxRadius := 0

	for _, house := range houses {
		// Binary search to find nearest heater
		idx := sort.SearchInts(heaters, house)
		minDist := int(^uint(0) >> 1) // MaxInt

		if idx < len(heaters) {
			dist := heaters[idx] - house
			if dist < 0 {
				dist = -dist
			}
			if dist < minDist {
				minDist = dist
			}
		}
		if idx > 0 {
			dist := house - heaters[idx-1]
			if dist < 0 {
				dist = -dist
			}
			if dist < minDist {
				minDist = dist
			}
		}

		if minDist > maxRadius {
			maxRadius = minDist
		}
	}

	return maxRadius
}
```

## 0477 — Total Hamming Distance

```go
package main

// LeetCode #477: Total Hamming Distance
// https://leetcode.com/problems/total-hamming-distance/
// Difficulty: Medium
// Time: O(n * 32) = O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(TotalHammingDistance([]int{4, 14, 2}))
	fmt.Println(TotalHammingDistance([]int{4, 14, 4}))
}

func TotalHammingDistance(nums []int) int {
	total := 0
	n := len(nums)

	for bit := 0; bit < 32; bit++ {
		countOnes := 0
		for _, num := range nums {
			if num&(1<<bit) != 0 {
				countOnes++
			}
		}
		countZeros := n - countOnes
		total += countOnes * countZeros
	}

	return total
}
```

## 0478 — Generate Random Point In A Circle

```go
package main

// LeetCode #478: Generate Random Point in a Circle
// https://leetcode.com/problems/generate-random-point-in-a-circle/
// Difficulty: Medium
// Time: O(1) per call
// Space: O(1)

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	sol := Constructor([]float64{1, 0, 0})
	// Generate a few random points
	for i := 0; i < 3; i++ {
		p := sol.RandPoint()
		fmt.Printf("%.4f %.4f\n", p[0], p[1])
	}
}

type Solution struct {
	radius, xCenter, yCenter float64
}

func Constructor(radiusAndCenter []float64) Solution {
	return Solution{
		radius:  radiusAndCenter[0],
		xCenter: radiusAndCenter[1],
		yCenter: radiusAndCenter[2],
	}
}

func (s *Solution) RandPoint() []float64 {
	// Use random angle and random radius (sqrt for uniform distribution)
	angle := rand.Float64() * 2 * math.Pi
	r := s.radius * math.Sqrt(rand.Float64())
	x := s.xCenter + r*math.Cos(angle)
	y := s.yCenter + r*math.Sin(angle)
	return []float64{x, y}
}
```

## 0481 — Magical String

```go
package main

// LeetCode #481: Magical String
// https://leetcode.com/problems/magical-string/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(MagicalString(6))
	fmt.Println(MagicalString(1))
}

func MagicalString(n int) int {
	if n <= 0 {
		return 0
	}
	if n <= 3 {
		return 1
	}

	s := make([]int, n)
	s[0], s[1], s[2] = 1, 2, 2
	count := 1
	writeIdx := 3
	readIdx := 2

	for writeIdx < n {
		val := 3 - s[writeIdx-1] // toggle between 1 and 2
		countTimes := s[readIdx]
		for i := 0; i < countTimes && writeIdx < n; i++ {
			s[writeIdx] = val
			if val == 1 {
				count++
			}
			writeIdx++
		}
		readIdx++
	}

	return count
}
```

## 0484 — Find Permutation

```go
package main

// LeetCode #484: Find Permutation
// https://leetcode.com/problems/find-permutation/
// Difficulty: Medium [Paid]
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(FindPermutation("I"))
	fmt.Println(FindPermutation("DI"))
}

func FindPermutation(s string) []int {
	n := len(s) + 1
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = i + 1
	}

	// Reverse contiguous segments for each 'D'
	for i := 0; i < len(s); i++ {
		if s[i] == 'D' {
			j := i
			for j < len(s) && s[j] == 'D' {
				j++
			}
			// Reverse segment from i to j
			left, right := i, j
			for left < right {
				result[left], result[right] = result[right], result[left]
				left++
				right--
			}
			i = j
		}
	}

	return result
}
```

## 0486 — Predict The Winner

```go
package main

// LeetCode #486: Predict the Winner
// https://leetcode.com/problems/predict-the-winner/
// Difficulty: Medium
// Time: O(n^2)
// Space: O(n^2)

import "fmt"

func main() {
	fmt.Println(PredictTheWinner([]int{1, 5, 2}))
	fmt.Println(PredictTheWinner([]int{1, 5, 233, 7}))
}

func PredictTheWinner(nums []int) bool {
	n := len(nums)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = nums[i]
	}

	for length := 2; length <= n; length++ {
		for i := 0; i <= n-length; i++ {
			j := i + length - 1
			// Max of (pick left) or (pick right), minus opponent's optimal play
			left := nums[i] - dp[i+1][j]
			right := nums[j] - dp[i][j-1]
			if left > right {
				dp[i][j] = left
			} else {
				dp[i][j] = right
			}
		}
	}

	return dp[0][n-1] >= 0
}
```

## 0487 — Max Consecutive Ones Ii

```go
package main

// LeetCode #487: Max Consecutive Ones II
// https://leetcode.com/problems/max-consecutive-ones-ii/
// Difficulty: Medium [Paid]
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(MaxConsecutiveOnesIi([]int{1, 0, 1, 1, 0}))
	fmt.Println(MaxConsecutiveOnesIi([]int{1, 0, 1, 1, 0, 1}))
}

func MaxConsecutiveOnesIi(nums []int) int {
	maxLen := 0
	prevLen, curLen := 0, 0

	for _, num := range nums {
		if num == 1 {
			curLen++
		} else {
			prevLen = curLen
			curLen = 0
		}
		if prevLen+curLen+1 > maxLen {
			maxLen = prevLen + curLen + 1
		}
	}

	if maxLen > len(nums) {
		return len(nums)
	}
	return maxLen
}
```

## 0490 — The Maze

```go
package main

// LeetCode #490: The Maze
// https://leetcode.com/problems/the-maze/
// Difficulty: Medium [Paid]
// Time: O(m * n)
// Space: O(m * n)

import "fmt"

func main() {
	maze := [][]int{
		{0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0},
		{1, 1, 0, 1, 1},
		{0, 0, 0, 0, 0},
	}
	fmt.Println(TheMaze(maze, []int{0, 4}, []int{4, 4}))
	fmt.Println(TheMaze(maze, []int{0, 4}, []int{3, 2}))
}

func TheMaze(maze [][]int, start []int, destination []int) bool {
	m, n := len(maze), len(maze[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	queue := [][]int{start}
	visited[start[0]][start[1]] = true

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur[0] == destination[0] && cur[1] == destination[1] {
			return true
		}

		for _, d := range dirs {
			r, c := cur[0], cur[1]
			// Roll until hitting a wall
			for r+d[0] >= 0 && r+d[0] < m && c+d[1] >= 0 && c+d[1] < n && maze[r+d[0]][c+d[1]] == 0 {
				r += d[0]
				c += d[1]
			}
			if !visited[r][c] {
				visited[r][c] = true
				queue = append(queue, []int{r, c})
			}
		}
	}

	return false
}
```

## 0491 — Non Decreasing Subsequences

```go
package main

// LeetCode #491: Non-decreasing Subsequences
// https://leetcode.com/problems/non-decreasing-subsequences/
// Difficulty: Medium
// Time: O(2^n * n) worst case
// Space: O(2^n * n)

import "fmt"

func main() {
	fmt.Println(NonDecreasingSubsequences([]int{4, 6, 7, 7}))
	fmt.Println(NonDecreasingSubsequences([]int{4, 4, 3, 2, 1}))
}

func NonDecreasingSubsequences(nums []int) [][]int {
	result := [][]int{}
	var backtrack func(start int, path []int)
	backtrack = func(start int, path []int) {
		if len(path) >= 2 {
			cp := make([]int, len(path))
			copy(cp, path)
			result = append(result, cp)
		}
		used := make(map[int]bool)
		for i := start; i < len(nums); i++ {
			if used[nums[i]] {
				continue
			}
			if len(path) == 0 || nums[i] >= path[len(path)-1] {
				used[nums[i]] = true
				backtrack(i+1, append(path, nums[i]))
			}
		}
	}
	backtrack(0, []int{})
	return result
}
```

## 0494 — Target Sum

```go
package main

// LeetCode #494: Target Sum
// https://leetcode.com/problems/target-sum/
// Difficulty: Medium
// Time: O(n * sum)
// Space: O(sum)

import "fmt"

func main() {
	fmt.Println(TargetSum([]int{1, 1, 1, 1, 1}, 3))
	fmt.Println(TargetSum([]int{1}, 1))
}

func TargetSum(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum < target || (sum-target)%2 != 0 {
		return 0
	}
	negSum := (sum - target) / 2

	dp := make([]int, negSum+1)
	dp[0] = 1
	for _, num := range nums {
		for s := negSum; s >= num; s-- {
			dp[s] += dp[s-num]
		}
	}
	return dp[negSum]
}
```

## 0497 — Random Point In Non Overlapping Rectangles

```go
package main

// LeetCode #497: Random Point in Non-overlapping Rectangles
// https://leetcode.com/problems/random-point-in-non-overlapping-rectangles/
// Difficulty: Medium
// Time: O(n) for init, O(log n) per pick
// Space: O(n)

import (
	"fmt"
	"math/rand"
)

func main() {
	sol := Constructor([][]int{{-2, -2, 1, 1}, {2, 2, 4, 6}})
	for i := 0; i < 3; i++ {
		p := sol.Pick()
		fmt.Println(p)
	}
}

type Solution struct {
	rects      [][]int
	prefixSum  []int
	totalPts   int
}

func Constructor(rects [][]int) Solution {
	prefixSum := make([]int, len(rects))
	total := 0
	for i, r := range rects {
		pts := (r[2] - r[0] + 1) * (r[3] - r[1] + 1)
		total += pts
		prefixSum[i] = total
	}
	return Solution{rects: rects, prefixSum: prefixSum, totalPts: total}
}

func (s *Solution) Pick() []int {
	// Pick a random point index
	r := rand.Intn(s.totalPts) + 1
	// Binary search to find which rectangle
	idx := search(s.prefixSum, r)

	rect := s.rects[idx]
	prev := 0
	if idx > 0 {
		prev = s.prefixSum[idx-1]
	}
	offset := r - prev - 1
	width := rect[2] - rect[0] + 1
	x := rect[0] + offset%width
	y := rect[1] + offset/width
	return []int{x, y}
}

func search(prefixSum []int, target int) int {
	lo, hi := 0, len(prefixSum)-1
	for lo < hi {
		mid := lo + (hi-lo)/2
		if prefixSum[mid] >= target {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}
```

## 0498 — Diagonal Traverse

```go
package main

// LeetCode #498: Diagonal Traverse
// https://leetcode.com/problems/diagonal-traverse/
// Difficulty: Medium
// Time: O(m * n)
// Space: O(1) (excluding output)

import "fmt"

func main() {
	fmt.Println(DiagonalTraverse([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	fmt.Println(DiagonalTraverse([][]int{{1, 2}, {3, 4}}))
}

func DiagonalTraverse(mat [][]int) []int {
	if len(mat) == 0 || len(mat[0]) == 0 {
		return []int{}
	}
	m, n := len(mat), len(mat[0])
	result := make([]int, m*n)
	row, col := 0, 0
	dir := 1 // 1 = up-right, -1 = down-left

	for i := 0; i < m*n; i++ {
		result[i] = mat[row][col]
		if dir == 1 { // moving up-right
			if col == n-1 {
				row++
				dir = -1
			} else if row == 0 {
				col++
				dir = -1
			} else {
				row--
				col++
			}
		} else { // moving down-left
			if row == m-1 {
				col++
				dir = 1
			} else if col == 0 {
				row++
				dir = 1
			} else {
				row++
				col--
			}
		}
	}

	return result
}
```

## 0503 — Next Greater Element Ii

```go
package main

// LeetCode #503: Next Greater Element II
// https://leetcode.com/problems/next-greater-element-ii/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(NextGreaterElementIi([]int{1, 2, 1}))
	fmt.Println(NextGreaterElementIi([]int{1, 2, 3, 4, 3}))
}

func NextGreaterElementIi(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	for i := range result {
		result[i] = -1
	}

	stack := []int{}
	// Iterate twice to handle circular array
	for i := 0; i < 2*n; i++ {
		num := nums[i%n]
		for len(stack) > 0 && nums[stack[len(stack)-1]] < num {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[idx] = num
		}
		if i < n {
			stack = append(stack, i)
		}
	}

	return result
}
```

## 0505 — The Maze Ii

```go
package main

// LeetCode #505: The Maze II
// https://leetcode.com/problems/the-maze-ii/
// Difficulty: Medium [Paid]
// Time: O(m * n * log(m*n)) with Dijkstra, or O(m * n * max(m,n)) with BFS-like
// Space: O(m * n)

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	maze := [][]int{
		{0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0},
		{1, 1, 0, 1, 1},
		{0, 0, 0, 0, 0},
	}
	fmt.Println(TheMazeIi(maze, []int{0, 4}, []int{4, 4}))
	fmt.Println(TheMazeIi(maze, []int{0, 4}, []int{3, 2}))
}

type Item struct {
	r, c, dist int
	index      int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].dist < pq[j].dist }
func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i]; pq[i].index = i; pq[j].index = j }
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func TheMazeIi(maze [][]int, start []int, destination []int) int {
	m, n := len(maze), len(maze[0])
	dist := make([][]int, m)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = math.MaxInt32
		}
	}

	dist[start[0]][start[1]] = 0
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{r: start[0], c: start[1], dist: 0})

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(*Item)
		if cur.r == destination[0] && cur.c == destination[1] {
			return cur.dist
		}
		if cur.dist > dist[cur.r][cur.c] {
			continue
		}

		for _, d := range dirs {
			r, c, steps := cur.r, cur.c, 0
			for r+d[0] >= 0 && r+d[0] < m && c+d[1] >= 0 && c+d[1] < n && maze[r+d[0]][c+d[1]] == 0 {
				r += d[0]
				c += d[1]
				steps++
			}
			newDist := cur.dist + steps
			if newDist < dist[r][c] {
				dist[r][c] = newDist
				heap.Push(pq, &Item{r: r, c: c, dist: newDist})
			}
		}
	}

	return -1
}
```

## 0508 — Most Frequent Subtree Sum

```go
package main

// LeetCode #508: Most Frequent Subtree Sum
// https://leetcode.com/problems/most-frequent-subtree-sum/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Test case 1: root = [5,2,-3]
	root1 := &TreeNode{Val: 5}
	root1.Left = &TreeNode{Val: 2}
	root1.Right = &TreeNode{Val: -3}
	fmt.Println(FindFrequentTreeSum(root1))

	// Test case 2: root = [5,2,-5]
	root2 := &TreeNode{Val: 5}
	root2.Left = &TreeNode{Val: 2}
	root2.Right = &TreeNode{Val: -5}
	fmt.Println(FindFrequentTreeSum(root2))
}

func FindFrequentTreeSum(root *TreeNode) []int {
	sumFreq := make(map[int]int)
	maxFreq := 0

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		sum := node.Val + dfs(node.Left) + dfs(node.Right)
		sumFreq[sum]++
		if sumFreq[sum] > maxFreq {
			maxFreq = sumFreq[sum]
		}
		return sum
	}

	dfs(root)

	result := []int{}
	for sum, freq := range sumFreq {
		if freq == maxFreq {
			result = append(result, sum)
		}
	}
	return result
}
```

## 0510 — Inorder Successor In Bst Ii

```go
package main

// LeetCode #510: Inorder Successor in BST II
// https://leetcode.com/problems/inorder-successor-in-bst-ii/
// Difficulty: Medium [Paid]
// Time: O(h) where h is height of tree
// Space: O(1)

import "fmt"

type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	Parent *Node
}

func main() {
	// Build tree: [2,1,3]
	root := &Node{Val: 2}
	root.Left = &Node{Val: 1, Parent: root}
	root.Right = &Node{Val: 3, Parent: root}
	fmt.Println(InorderSuccessorInBstIi(root.Left).Val) // node 1 -> successor 2
	fmt.Println(InorderSuccessorInBstIi(root).Val)       // node 2 -> successor 3

	// For node 3, successor should be nil
	successor := InorderSuccessorInBstIi(root.Right)
	if successor == nil {
		fmt.Println("nil")
	} else {
		fmt.Println(successor.Val)
	}
}

func InorderSuccessorInBstIi(node *Node) *Node {
	if node == nil {
		return nil
	}

	// If right child exists, find leftmost in right subtree
	if node.Right != nil {
		cur := node.Right
		for cur.Left != nil {
			cur = cur.Left
		}
		return cur
	}

	// Otherwise, go up until we find a node that is a left child
	cur := node
	for cur.Parent != nil && cur.Parent.Right == cur {
		cur = cur.Parent
	}

	return cur.Parent
}
```

## 0513 — Find Bottom Left Tree Value

```go
package main

// LeetCode #513: Find Bottom Left Tree Value
// https://leetcode.com/problems/find-bottom-left-tree-value/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Test: root = [2,1,3]
	root1 := &TreeNode{Val: 2}
	root1.Left = &TreeNode{Val: 1}
	root1.Right = &TreeNode{Val: 3}
	fmt.Println(FindBottomLeftTreeValue(root1))

	// Test: root = [1,2,3,4,null,5,6,null,null,7]
	root2 := &TreeNode{Val: 1}
	root2.Left = &TreeNode{Val: 2, Left: &TreeNode{Val: 4}}
	root2.Right = &TreeNode{Val: 3, Left: &TreeNode{Val: 5, Left: &TreeNode{Val: 7}}, Right: &TreeNode{Val: 6}}
	fmt.Println(FindBottomLeftTreeValue(root2))
}

func FindBottomLeftTreeValue(root *TreeNode) int {
	queue := []*TreeNode{root}
	var leftmost int

	for len(queue) > 0 {
		levelSize := len(queue)
		leftmost = queue[0].Val
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return leftmost
}
```

## 0515 — Find Largest Value In Each Tree Row

```go
package main

// LeetCode #515: Find Largest Value in Each Tree Row
// https://leetcode.com/problems/find-largest-value-in-each-tree-row/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Test: root = [1,3,2,5,3,null,9]
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 3, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 3}}
	root.Right = &TreeNode{Val: 2, Right: &TreeNode{Val: 9}}
	fmt.Println(LargestValues(root))
}

func LargestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := []int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		maxVal := math.MinInt32
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Val > maxVal {
				maxVal = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, maxVal)
	}

	return result
}
```

## 0516 — Longest Palindromic Subsequence

```go
package main

// LeetCode #516: Longest Palindromic Subsequence
// https://leetcode.com/problems/longest-palindromic-subsequence/
// Difficulty: Medium
// Time: O(n^2)
// Space: O(n^2)

import "fmt"

func main() {
	fmt.Println(LongestPalindromicSubsequence("bbbab"))
	fmt.Println(LongestPalindromicSubsequence("cbbd"))
}

func LongestPalindromicSubsequence(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}

	for length := 2; length <= n; length++ {
		for i := 0; i <= n-length; i++ {
			j := i + length - 1
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				if dp[i+1][j] > dp[i][j-1] {
					dp[i][j] = dp[i+1][j]
				} else {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}

	return dp[0][n-1]
}
```

## 0518 — Coin Change Ii

```go
package main

// LeetCode #518: Coin Change II
// https://leetcode.com/problems/coin-change-ii/
// Difficulty: Medium
// Time: O(amount * n) where n = len(coins)
// Space: O(amount)

import "fmt"

func main() {
	fmt.Println(CoinChangeIi(5, []int{1, 2, 5}))
	fmt.Println(CoinChangeIi(3, []int{2}))
	fmt.Println(CoinChangeIi(10, []int{10}))
}

func CoinChangeIi(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}

	return dp[amount]
}
```

## 0519 — Random Flip Matrix

```go
package main

// LeetCode #519: Random Flip Matrix
// https://leetcode.com/problems/random-flip-matrix/
// Difficulty: Medium
// Time: O(1) per flip/reset amortized
// Space: O(k) where k = number of flips

import (
	"fmt"
	"math/rand"
)

func main() {
	sol := Constructor(3, 1)
	for i := 0; i < 3; i++ {
		fmt.Println(sol.Flip())
	}
	sol.Reset()
	fmt.Println("reset done")
}

type Solution struct {
	m, n, total int
	used        map[int]int
}

func Constructor(m int, n int) Solution {
	return Solution{m: m, n: n, total: m * n, used: make(map[int]int)}
}

func (s *Solution) Flip() []int {
	randIdx := rand.Intn(s.total)
	s.total--

	// Use Fisher-Yates style swap with map for sparse tracking
	actualIdx := randIdx
	if val, ok := s.used[randIdx]; ok {
		actualIdx = val
	}
	if val, ok := s.used[s.total]; ok {
		s.used[randIdx] = val
	} else {
		s.used[randIdx] = s.total
	}

	return []int{actualIdx / s.n, actualIdx % s.n}
}

func (s *Solution) Reset() {
	s.total = s.m * s.n
	s.used = make(map[int]int)
}
```

## 0522 — Longest Uncommon Subsequence Ii

```go
package main

// LeetCode #522: Longest Uncommon Subsequence II
// https://leetcode.com/problems/longest-uncommon-subsequence-ii/
// Difficulty: Medium
// Time: O(n^2 * L) where L is max length
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(FindLUSlength([]string{"aba", "cdc", "eae"}))
	fmt.Println(FindLUSlength([]string{"aaa", "aaa", "aa"}))
}

func FindLUSlength(strs []string) int {
	maxLen := -1

	for i := 0; i < len(strs); i++ {
		isUnique := true
		for j := 0; j < len(strs); j++ {
			if i != j && isSubseq(strs[i], strs[j]) {
				isUnique = false
				break
			}
		}
		if isUnique && len(strs[i]) > maxLen {
			maxLen = len(strs[i])
		}
	}

	return maxLen
}

func isSubseq(a, b string) bool {
	i := 0
	for j := 0; i < len(a) && j < len(b); j++ {
		if a[i] == b[j] {
			i++
		}
	}
	return i == len(a)
}
```

