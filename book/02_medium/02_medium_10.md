# Medium (Sedang) — Problem ��1864

## 1666 — Change The Root Of A Binary Tree

```go
package main

// LeetCode #1666: Change the Root of a Binary Tree
// https://leetcode.com/problems/change-the-root-of-a-binary-tree/
// Difficulty: Medium [Paid]
// Time: O(n), Space: O(1) where n = path length from leaf to root

import "fmt"

type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	Parent *Node
}

func flipBinaryTree(root *Node, leaf *Node) *Node {
	if leaf == nil {
		return nil
	}

	cur := leaf
	var newParent *Node

	for cur != nil {
		oldParent := cur.Parent
		cur.Parent = newParent

		// Clear the link if newParent was one of cur's children
		if cur.Left == newParent {
			cur.Left = nil
		}
		if cur.Right == newParent {
			cur.Right = nil
		}

		// Make the old parent a child of current node
		if oldParent != nil {
			if cur.Right == nil {
				cur.Right = oldParent
			} else if cur.Left == nil {
				cur.Left = oldParent
			} else {
				// Both children occupied: move left to right, put oldParent in left
				cur.Right = cur.Left
				cur.Left = oldParent
			}
		}

		newParent = cur
		cur = oldParent
	}

	return leaf
}

func printTree(node *Node, indent string) {
	if node == nil {
		return
	}
	fmt.Printf("%sNode(%d)", indent, node.Val)
	if node.Parent != nil {
		fmt.Printf(" parent=%d", node.Parent.Val)
	} else {
		fmt.Printf(" parent=nil")
	}
	fmt.Println()
	printTree(node.Left, indent+"  L:")
	printTree(node.Right, indent+"  R:")
}

func main() {
	// Test case 1: leaf = 5 (a true leaf)
	// Tree:
	//     1
	//    / \
	//   2   3
	//  / \
	// 4   5
	root := &Node{Val: 1}
	root.Left = &Node{Val: 2, Parent: root}
	root.Right = &Node{Val: 3, Parent: root}
	root.Left.Left = &Node{Val: 4, Parent: root.Left}
	root.Left.Right = &Node{Val: 5, Parent: root.Left}

	newRoot := flipBinaryTree(root, root.Left.Right)
	fmt.Println("Test 1: Flip with leaf=5")
	fmt.Println("New root is leaf:", newRoot.Val == 5)
	fmt.Println("New root parent nil:", newRoot.Parent == nil)
	fmt.Println("5.Right should be 2:", newRoot.Right.Val == 2)
	fmt.Println("2.Parent should be 5:", newRoot.Right.Parent.Val == 5)
	fmt.Println()

	// Test case 2: leaf = root (no change needed)
	root2 := &Node{Val: 10}
	root2.Left = &Node{Val: 20, Parent: root2}
	newRoot2 := flipBinaryTree(root2, root2)
	fmt.Println("Test 2: Leaf is root (no change)")
	fmt.Println("Root unchanged:", newRoot2.Val == 10)
	fmt.Println("Root parent nil:", newRoot2.Parent == nil)
	fmt.Println()

	// Test case 3: Simple chain
	// 1 -> 2 -> 3, leaf = 3
	root3 := &Node{Val: 1}
	root3.Right = &Node{Val: 2, Parent: root3}
	root3.Right.Right = &Node{Val: 3, Parent: root3.Right}

	newRoot3 := flipBinaryTree(root3, root3.Right.Right)
	fmt.Println("Test 3: Chain flip with leaf=3")
	fmt.Println("New root is leaf:", newRoot3.Val == 3)
	fmt.Println("3.Right should be 2:", newRoot3.Right.Val == 2)
	fmt.Println("2.Right should be 1:", newRoot3.Right.Right.Val == 1)
	fmt.Println("1.Parent should be 2:", newRoot3.Right.Right.Parent.Val == 2)
	fmt.Println()
}
```

## 1669 — Merge In Between Linked Lists

```go
package main

// LeetCode #1669: Merge In Between Linked Lists
// https://leetcode.com/problems/merge-in-between-linked-lists/
// Difficulty: Medium
// Time: O(n + m), Space: O(1)

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	dummy := &ListNode{Next: list1}
	prev := dummy

	// Move to node just before position a
	for i := 0; i < a; i++ {
		prev = prev.Next
	}

	// Find end of segment to remove (node at position b)
	end := prev
	for i := a; i <= b; i++ {
		end = end.Next
	}

	// Connect prev node to list2
	prev.Next = list2

	// Find tail of list2
	tail := list2
	for tail != nil && tail.Next != nil {
		tail = tail.Next
	}

	// Connect tail of list2 to node after position b
	if tail != nil {
		tail.Next = end.Next
	}

	return dummy.Next
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func main() {
	// Test case 1
	l1 := &ListNode{0, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}}
	l2 := &ListNode{100, &ListNode{101, &ListNode{102, nil}}}
	result := mergeInBetween(l1, 3, 4, l2)
	fmt.Print("Test 1: ")
	printList(result) // Expected: 0 1 2 100 101 102 5

	// Test case 2
	l1 = &ListNode{0, &ListNode{1, &ListNode{2, &ListNode{3, nil}}}}
	l2 = &ListNode{100, nil}
	result = mergeInBetween(l1, 1, 2, l2)
	fmt.Print("Test 2: ")
	printList(result) // Expected: 0 100 3

	// Test case 3
	l1 = &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	l2 = &ListNode{10, &ListNode{20, nil}}
	result = mergeInBetween(l1, 0, 1, l2)
	fmt.Print("Test 3: ")
	printList(result) // Expected: 10 20 3
}
```

## 1670 — Design Front Middle Back Queue

```go
package main

// LeetCode #1670: Design Front Middle Back Queue
// https://leetcode.com/problems/design-front-middle-back-queue/
// Difficulty: Medium
// Time: O(1) per operation (amortized), Space: O(n)

import "fmt"

type FrontMiddleBackQueue struct {
	left  []int
	right []int
}

func Constructor() FrontMiddleBackQueue {
	return FrontMiddleBackQueue{}
}

func (q *FrontMiddleBackQueue) balance() {
	// Keep left size >= right size, and left size - right size <= 1
	if len(q.left) > len(q.right)+1 {
		// Move last of left to front of right
		v := q.left[len(q.left)-1]
		q.left = q.left[:len(q.left)-1]
		q.right = append([]int{v}, q.right...)
	} else if len(q.right) > len(q.left) {
		// Move first of right to back of left
		v := q.right[0]
		q.right = q.right[1:]
		q.left = append(q.left, v)
	}
}

func (q *FrontMiddleBackQueue) PushFront(val int) {
	q.left = append([]int{val}, q.left...)
	q.balance()
}

func (q *FrontMiddleBackQueue) PushMiddle(val int) {
	q.left = append(q.left, val)
	q.balance()
}

func (q *FrontMiddleBackQueue) PushBack(val int) {
	q.right = append(q.right, val)
	q.balance()
}

func (q *FrontMiddleBackQueue) PopFront() int {
	if len(q.left) == 0 {
		return -1
	}
	val := q.left[0]
	q.left = q.left[1:]
	q.balance()
	return val
}

func (q *FrontMiddleBackQueue) PopMiddle() int {
	if len(q.left) == 0 {
		return -1
	}
	val := q.left[len(q.left)-1]
	q.left = q.left[:len(q.left)-1]
	q.balance()
	return val
}

func (q *FrontMiddleBackQueue) PopBack() int {
	if len(q.right) == 0 {
		if len(q.left) == 0 {
			return -1
		}
		val := q.left[len(q.left)-1]
		q.left = q.left[:len(q.left)-1]
		return val
	}
	val := q.right[len(q.right)-1]
	q.right = q.right[:len(q.right)-1]
	q.balance()
	return val
}

func main() {
	q := Constructor()
	q.PushFront(1)
	q.PushBack(2)
	q.PushMiddle(3)
	q.PushMiddle(4)

	fmt.Println("PopFront:", q.PopFront()) // 1
	fmt.Println("PopMiddle:", q.PopMiddle()) // 3
	fmt.Println("PopMiddle:", q.PopMiddle()) // 4
	fmt.Println("PopBack:", q.PopBack()) // 2
	fmt.Println("PopFront:", q.PopFront()) // -1 (empty)

	// Test case 2
	q2 := Constructor()
	q2.PushFront(1)
	q2.PushFront(2)
	q2.PushBack(3)
	q2.PushBack(4)
	fmt.Println("PopFront:", q2.PopFront()) // 2
	fmt.Println("PopBack:", q2.PopBack())   // 4
	fmt.Println("PopFront:", q2.PopFront()) // 1
}
```

## 1673 — Find The Most Competitive Subsequence

```go
package main

// LeetCode #1673: Find the Most Competitive Subsequence
// https://leetcode.com/problems/find-the-most-competitive-subsequence/
// Difficulty: Medium
// Time: O(n), Space: O(k)

import "fmt"

func mostCompetitive(nums []int, k int) []int {
	n := len(nums)
	stack := make([]int, 0, k)
	toRemove := n - k

	for _, num := range nums {
		for len(stack) > 0 && toRemove > 0 && stack[len(stack)-1] > num {
			stack = stack[:len(stack)-1]
			toRemove--
		}
		stack = append(stack, num)
	}

	// If we haven't removed enough, trim from end
	return stack[:k]
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", mostCompetitive([]int{3, 5, 2, 6}, 2)) // Expected: [2, 6]

	// Test case 2
	fmt.Println("Test 2:", mostCompetitive([]int{2, 4, 3, 3, 5, 4, 9, 6}, 4)) // Expected: [2, 3, 3, 4]

	// Test case 3
	fmt.Println("Test 3:", mostCompetitive([]int{1, 2, 3, 4, 5}, 3)) // Expected: [1, 2, 3]

	// Test case 4
	fmt.Println("Test 4:", mostCompetitive([]int{5, 4, 3, 2, 1}, 2)) // Expected: [1, 2]
}
```

## 1674 — Minimum Moves To Make Array Complementary

```go
package main

// LeetCode #1674: Minimum Moves to Make Array Complementary
// https://leetcode.com/problems/minimum-moves-to-make-array-complementary/
// Difficulty: Medium
// Time: O(n + limit), Space: O(limit)

import "fmt"

func minMoves(nums []int, limit int) int {
	n := len(nums)
	delta := make([]int, 2*limit+2) // difference array

	for i := 0; i < n/2; i++ {
		a, b := nums[i], nums[n-1-i]
		// Pair sum
		pairSum := a + b
		// min achievable sum: min(a,b) + 1
		// max achievable sum: max(a,b) + limit

		// 0 changes: [pairSum, pairSum] is already achievable
		// 1 change: [min(a,b)+1, max(a,b)+limit]
		// 2 changes: everything else [2, 2*limit]

		delta[2] += 2
		delta[min(a, b)+1]-- // Reduce to 1 change at this range start
		delta[pairSum]--     // Make it 0 changes at pairSum
		delta[pairSum+1]++   // Back to 1 change after pairSum
		delta[max(a, b)+limit+1]++ // Back to 2 changes
	}

	ans := n
	cur := 0
	for target := 2; target <= 2*limit; target++ {
		cur += delta[target]
		if cur < ans {
			ans = cur
		}
	}
	return ans
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
	fmt.Println("Test 1:", minMoves([]int{1, 2, 4, 3}, 4)) // Expected: 1

	// Test case 2
	fmt.Println("Test 2:", minMoves([]int{1, 2, 2, 1}, 2)) // Expected: 2

	// Test case 3
	fmt.Println("Test 3:", minMoves([]int{1, 1, 1, 1}, 3)) // Expected: 0 (all pairs already sum to 2)
}
```

## 1676 — Lowest Common Ancestor Of A Binary Tree Iv

```go
package main

// LeetCode #1676: Lowest Common Ancestor of a Binary Tree IV
// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree-iv/
// Difficulty: Medium [Paid]
// Time: O(n), Space: O(h) where h is height

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root *TreeNode, nodes []*TreeNode) *TreeNode {
	nodeSet := make(map[*TreeNode]bool)
	for _, n := range nodes {
		nodeSet[n] = true
	}
	return dfs(root, nodeSet)
}

func dfs(node *TreeNode, nodeSet map[*TreeNode]bool) *TreeNode {
	if node == nil {
		return nil
	}
	if nodeSet[node] {
		return node
	}
	left := dfs(node.Left, nodeSet)
	right := dfs(node.Right, nodeSet)
	if left != nil && right != nil {
		return node
	}
	if left != nil {
		return left
	}
	return right
}

func main() {
	// Test case 1
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 1}
	root.Left.Left = &TreeNode{Val: 6}
	root.Left.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 0}
	root.Right.Right = &TreeNode{Val: 8}
	root.Left.Right.Left = &TreeNode{Val: 7}
	root.Left.Right.Right = &TreeNode{Val: 4}

	result := lowestCommonAncestor(root, []*TreeNode{root.Left, root.Right})
	fmt.Println("Test 1 (LCA of 5 and 1):", result.Val) // Expected: 3

	// Test case 2: LCA of 6, 7, 4 → 5
	result = lowestCommonAncestor(root, []*TreeNode{root.Left.Left, root.Left.Right.Left, root.Left.Right.Right})
	fmt.Println("Test 2 (LCA of 6, 7, 4):", result.Val) // Expected: 5

	// Test case 3: Single node
	result = lowestCommonAncestor(root, []*TreeNode{root.Left.Right})
	fmt.Println("Test 3 (single node):", result.Val) // Expected: 2
}
```

## 1679 — Max Number Of K Sum Pairs

```go
package main

// LeetCode #1679: Max Number of K-Sum Pairs
// https://leetcode.com/problems/max-number-of-k-sum-pairs/
// Difficulty: Medium
// Time: O(n), Space: O(n)

import "fmt"

func maxOperations(nums []int, k int) int {
	counts := make(map[int]int)
	ops := 0

	for _, num := range nums {
		complement := k - num
		if counts[complement] > 0 {
			ops++
			counts[complement]--
		} else {
			counts[num]++
		}
	}
	return ops
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", maxOperations([]int{1, 2, 3, 4}, 5)) // Expected: 2

	// Test case 2
	fmt.Println("Test 2:", maxOperations([]int{3, 1, 3, 4, 3}, 6)) // Expected: 1

	// Test case 3
	fmt.Println("Test 3:", maxOperations([]int{1, 2, 3, 4, 5, 6}, 7)) // Expected: 3
}
```

## 1680 — Concatenation Of Consecutive Binary Numbers

```go
package main

// LeetCode #1680: Concatenation of Consecutive Binary Numbers
// https://leetcode.com/problems/concatenation-of-consecutive-binary-numbers/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

const mod = 1_000_000_007

func concatenatedBinary(n int) int {
	result := 0
	lenBits := 0

	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			lenBits++
		}
		result = (result<<lenBits | i) % mod
	}
	return result
}

func main() {
	fmt.Println(concatenatedBinary(1))   // Expected: 1
	fmt.Println(concatenatedBinary(3))   // Expected: 27
	fmt.Println(concatenatedBinary(12))  // Expected: 505379714
}
```

## 1682 — Longest Palindromic Subsequence Ii

```go
package main

// LeetCode #1682: Longest Palindromic Subsequence II
// https://leetcode.com/problems/longest-palindromic-subsequence-ii/
// Difficulty: Medium [Paid]
// Time: O(n^2), Space: O(n^2)

import "fmt"

func longestPalindromeSubseq(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	// dp[i][j] = longest LPS length in s[i..j] with no equal adjacent chars
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}

	for length := 2; length <= n; length++ {
		for i := 0; i+length-1 < n; i++ {
			j := i + length - 1
			if s[i] == s[j] {
				// Check adjacent condition: s[i] != the char that comes before/after
				if length == 2 {
					dp[i][j] = 2
				} else if s[i] != s[i+1] && s[j] != s[j-1] {
					dp[i][j] = max(dp[i][j], dp[i+1][j-1]+2)
				} else {
					dp[i][j] = max(dp[i][j], dp[i+1][j-1])
				}
			}
			dp[i][j] = max(dp[i][j], max(dp[i+1][j], dp[i][j-1]))
		}
	}
	return dp[0][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(longestPalindromeSubseq("bbabab"))    // Expected: 4 ("baba" or "abab")
	fmt.Println(longestPalindromeSubseq("dcbccacdb")) // Expected: 4
	fmt.Println(longestPalindromeSubseq("a"))         // Expected: 1
}
```

## 1685 — Sum Of Absolute Differences In A Sorted Array

```go
package main

// LeetCode #1685: Sum of Absolute Differences in a Sorted Array
// https://leetcode.com/problems/sum-of-absolute-differences-in-a-sorted-array/
// Difficulty: Medium
// Time: O(n), Space: O(1) (excluding output)

import "fmt"

func getSumAbsoluteDifferences(nums []int) []int {
	n := len(nums)
	total := 0
	for _, v := range nums {
		total += v
	}

	result := make([]int, n)
	prefix := 0
	for i, v := range nums {
		// Left side: v * i - prefix_sum_left
		// Right side: (total - prefix_sum_left - v) - v * (n-1-i)
		left := v*i - prefix
		right := (total - prefix - v) - v*(n-1-i)
		result[i] = left + right
		prefix += v
	}
	return result
}

func main() {
	fmt.Println(getSumAbsoluteDifferences([]int{2, 3, 5}))    // Expected: [4, 3, 5]
	fmt.Println(getSumAbsoluteDifferences([]int{1, 4, 6, 8, 10})) // Expected: [24, 15, 13, 15, 21]
	fmt.Println(getSumAbsoluteDifferences([]int{1, 2}))       // Expected: [1, 1]
}
```

## 1686 — Stone Game Vi

```go
package main

// LeetCode #1686: Stone Game VI
// https://leetcode.com/problems/stone-game-vi/
// Difficulty: Medium
// Time: O(n log n), Space: O(n)

import (
	"fmt"
	"sort"
)

func stoneGameVI(aliceValues []int, bobValues []int) int {
	n := len(aliceValues)
	pairs := make([][2]int, n)
	for i := 0; i < n; i++ {
		pairs[i] = [2]int{aliceValues[i] + bobValues[i], i}
	}

	// Sort by sum descending
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] > pairs[j][0]
	})

	aliceScore := 0
	bobScore := 0
	for i := 0; i < n; i++ {
		idx := pairs[i][1]
		if i%2 == 0 {
			aliceScore += aliceValues[idx]
		} else {
			bobScore += bobValues[idx]
		}
	}

	if aliceScore > bobScore {
		return 1
	} else if bobScore > aliceScore {
		return -1
	}
	return 0
}

func main() {
	fmt.Println(stoneGameVI([]int{1, 3}, []int{2, 1}))      // Expected: 1 (Alice wins)
	fmt.Println(stoneGameVI([]int{1, 2}, []int{3, 1}))      // Expected: 0 (tie)
	fmt.Println(stoneGameVI([]int{2, 4, 3}, []int{1, 6, 7})) // Expected: -1 (Bob wins)
}
```

## 1689 — Partitioning Into Minimum Number Of Deci Binary Numbers

```go
package main

// LeetCode #1689: Partitioning Into Minimum Number Of Deci-Binary Numbers
// https://leetcode.com/problems/partitioning-into-minimum-number-of-deci-binary-numbers/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

func minPartitions(n string) int {
	maxDigit := 0
	for _, ch := range n {
		digit := int(ch - '0')
		if digit > maxDigit {
			maxDigit = digit
		}
		if maxDigit == 9 {
			break // can't get higher than 9
		}
	}
	return maxDigit
}

func main() {
	fmt.Println(minPartitions("32"))     // Expected: 3
	fmt.Println(minPartitions("82734"))  // Expected: 8
	fmt.Println(minPartitions("27346209830709182346")) // Expected: 9
}
```

## 1690 — Stone Game Vii

```go
package main

// LeetCode #1690: Stone Game VII
// https://leetcode.com/problems/stone-game-vii/
// Difficulty: Medium
// Time: O(n^2), Space: O(n)

import "fmt"

func stoneGameVII(stones []int) int {
	n := len(stones)
	prefix := make([]int, n+1)
	for i, v := range stones {
		prefix[i+1] = prefix[i] + v
	}

	// dp[i][j] = max score difference (current player - other player) for subarray i..j
	dp := make([]int, n)

	for length := 2; length <= n; length++ {
		for i := 0; i+length-1 < n; i++ {
			j := i + length - 1
			// Remove left: score = sum of rest, then subtract opponent's optimal
			removeLeft := (prefix[j+1] - prefix[i+1]) - dp[i+1]
			// Remove right: score = sum of rest, then subtract opponent's optimal
			removeRight := (prefix[j] - prefix[i]) - dp[i]
			dp[i] = max(removeLeft, removeRight)
		}
	}
	return dp[0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(stoneGameVII([]int{5, 3, 1, 4, 2})) // Expected: 6
	fmt.Println(stoneGameVII([]int{7, 90, 5, 1, 100, 10, 10, 2})) // Expected: 122
	fmt.Println(stoneGameVII([]int{1, 2})) // Expected: 2
}
```

## 1695 — Maximum Erasure Value

```go
package main

// LeetCode #1695: Maximum Erasure Value
// https://leetcode.com/problems/maximum-erasure-value/
// Difficulty: Medium
// Time: O(n), Space: O(n)

import "fmt"

func maximumUniqueSubarray(nums []int) int {
	lastPos := make(map[int]int)
	maxSum := 0
	currentSum := 0
	left := 0

	for right, num := range nums {
		if pos, ok := lastPos[num]; ok && pos >= left {
			// Remove elements from left to pos
			for left <= pos {
				currentSum -= nums[left]
				left++
			}
		}
		currentSum += num
		lastPos[num] = right
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}
	return maxSum
}

func main() {
	fmt.Println(maximumUniqueSubarray([]int{4, 2, 4, 5, 6}))   // Expected: 17
	fmt.Println(maximumUniqueSubarray([]int{5, 2, 1, 2, 5, 2, 1, 2, 5})) // Expected: 8
	fmt.Println(maximumUniqueSubarray([]int{1})) // Expected: 1
}
```

## 1696 — Jump Game Vi

```go
package main

// LeetCode #1696: Jump Game VI
// https://leetcode.com/problems/jump-game-vi/
// Difficulty: Medium
// Time: O(n), Space: O(k)

import "fmt"

func maxResult(nums []int, k int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]

	// Monotonic deque storing indices with decreasing dp values
	deque := make([]int, 0, n)
	deque = append(deque, 0)

	for i := 1; i < n; i++ {
		// Remove out-of-range elements
		for len(deque) > 0 && deque[0] < i-k {
			deque = deque[1:]
		}

		// dp[i] = nums[i] + max dp from i-k to i-1
		dp[i] = nums[i] + dp[deque[0]]

		// Maintain decreasing order
		for len(deque) > 0 && dp[deque[len(deque)-1]] <= dp[i] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
	}
	return dp[n-1]
}

func main() {
	fmt.Println(maxResult([]int{1, -1, -2, 4, -7, 3}, 2)) // Expected: 7
	fmt.Println(maxResult([]int{10, -5, -2, 4, 0, 3}, 3)) // Expected: 17
	fmt.Println(maxResult([]int{1, -5, -20, 4, -1, 3, -6, -3}, 2)) // Expected: 0
}
```

## 1698 — Number Of Distinct Substrings In A String

```go
package main

// LeetCode #1698: Number of Distinct Substrings in a String
// https://leetcode.com/problems/number-of-distinct-substrings-in-a-string/
// Difficulty: Medium [Paid]
// Time: O(n^2), Space: O(n^2) using trie

import "fmt"

type TrieNode struct {
	children [26]*TrieNode
}

func countDistinct(s string) int {
	root := &TrieNode{}
	count := 0

	for i := 0; i < len(s); i++ {
		node := root
		for j := i; j < len(s); j++ {
			idx := s[j] - 'a'
			if node.children[idx] == nil {
				node.children[idx] = &TrieNode{}
				count++
			}
			node = node.children[idx]
		}
	}
	return count
}

func main() {
	fmt.Println(countDistinct("aabbaba")) // Expected: 21
	fmt.Println(countDistinct("abcdef"))  // Expected: 21
	fmt.Println(countDistinct("a"))       // Expected: 1
}
```

## 1699 — Number Of Calls Between Two Persons

```go
package main

// LeetCode #1699: Number of Calls Between Two Persons
// https://leetcode.com/problems/number-of-calls-between-two-persons/
// Difficulty: Medium [Paid] (SQL)
// This is a SQL problem. In Go we simulate the logic.
// Time: O(n), Space: O(n)

import "fmt"

type CallRecord struct {
	FromID int
	ToID   int
	Dur    int
}

func numberOfCalls(records []CallRecord) [][3]int {
	callMap := make(map[[2]int]int) // [min,max] -> total duration

	for _, r := range records {
		a, b := r.FromID, r.ToID
		if a > b {
			a, b = b, a
		}
		key := [2]int{a, b}
		callMap[key] += r.Dur
	}

	result := make([][3]int, 0, len(callMap))
	for key, dur := range callMap {
		result = append(result, [3]int{key[0], key[1], dur})
	}
	return result
}

func main() {
	records := []CallRecord{
		{1, 2, 10},
		{2, 1, 20},
		{1, 3, 30},
	}
	result := numberOfCalls(records)
	for _, r := range result {
		fmt.Printf("(%d,%d): %d\n", r[0], r[1], r[2])
	}
}
```

## 1701 — Average Waiting Time

```go
package main

// LeetCode #1701: Average Waiting Time
// https://leetcode.com/problems/average-waiting-time/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

func averageWaitingTime(customers [][]int) float64 {
	currentTime := 0
	totalWait := 0

	for _, c := range customers {
		arrival, prepTime := c[0], c[1]
		if currentTime < arrival {
			currentTime = arrival
		}
		currentTime += prepTime
		totalWait += currentTime - arrival
	}

	return float64(totalWait) / float64(len(customers))
}

func main() {
	fmt.Println(averageWaitingTime([][]int{{1, 2}, {2, 5}, {4, 3}})) // Expected: 5.0
	fmt.Println(averageWaitingTime([][]int{{5, 2}, {5, 4}, {10, 3}, {20, 1}})) // Expected: 3.25
	fmt.Println(averageWaitingTime([][]int{{2, 3}, {6, 3}, {7, 5}, {11, 3}, {15, 2}, {18, 1}})) // Expected: 4.16667
}
```

## 1702 — Maximum Binary String After Change

```go
package main

// LeetCode #1702: Maximum Binary String After Change
// https://leetcode.com/problems/maximum-binary-string-after-change/
// Difficulty: Medium
// Time: O(n), Space: O(n)

import "fmt"

func maximumBinaryString(binary string) string {
	n := len(binary)
	zeros := 0
	firstZero := -1

	for i, ch := range binary {
		if ch == '0' {
			zeros++
			if firstZero == -1 {
				firstZero = i
			}
		}
	}

	if zeros <= 1 {
		return binary
	}

	// Result: all 1s except position (firstZero + zeros - 1)
	result := make([]byte, n)
	for i := range result {
		if i == firstZero+zeros-1 {
			result[i] = '0'
		} else {
			result[i] = '1'
		}
	}
	return string(result)
}

func main() {
	fmt.Println(maximumBinaryString("000110")) // Expected: "111011"
	fmt.Println(maximumBinaryString("01"))     // Expected: "01"
	fmt.Println(maximumBinaryString("10"))     // Expected: "10"
}
```

## 1705 — Maximum Number Of Eaten Apples

```go
package main

// LeetCode #1705: Maximum Number of Eaten Apples
// https://leetcode.com/problems/maximum-number-of-eaten-apples/
// Difficulty: Medium
// Time: O(n log n), Space: O(n)

import (
	"container/heap"
	"fmt"
)

type Apple struct {
	rottenDay int
	count     int
}

type MinHeap []Apple

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].rottenDay < h[j].rottenDay }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(Apple)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func eatenApples(apples []int, days []int) int {
	h := &MinHeap{}
	heap.Init(h)
	eaten := 0
	day := 0

	for day < len(apples) || h.Len() > 0 {
		// New apples grow
		if day < len(apples) && apples[day] > 0 {
			heap.Push(h, Apple{rottenDay: day + days[day], count: apples[day]})
		}

		// Remove rotten apples
		for h.Len() > 0 && h.Len() > 0 && (*h)[0].rottenDay <= day {
			heap.Pop(h)
		}

		// Eat one apple
		if h.Len() > 0 {
			top := &(*h)[0]
			top.count--
			if top.count == 0 {
				heap.Pop(h)
			}
			eaten++
		}

		day++
	}
	return eaten
}

func main() {
	fmt.Println(eatenApples([]int{1, 2, 3, 5, 2}, []int{3, 2, 1, 4, 2})) // Expected: 7
	fmt.Println(eatenApples([]int{3, 0, 0, 0, 0, 2}, []int{3, 0, 0, 0, 0, 2})) // Expected: 5
	fmt.Println(eatenApples([]int{1}, []int{2})) // Expected: 1
}
```

## 1706 — Where Will The Ball Fall

```go
package main

// LeetCode #1706: Where Will the Ball Fall
// https://leetcode.com/problems/where-will-the-ball-fall/
// Difficulty: Medium
// Time: O(m * n), Space: O(1) (excluding output)

import "fmt"

func findBall(grid [][]int) []int {
	m, n := len(grid), len(grid[0])
	result := make([]int, n)

	for col := 0; col < n; col++ {
		c := col
		for r := 0; r < m; r++ {
			// If the cell is 1 (sloping right), check the cell to the right
			if grid[r][c] == 1 {
				if c+1 >= n || grid[r][c+1] == -1 {
					c = -1
					break
				}
				c++
			} else {
				// Cell is -1 (sloping left), check the cell to the left
				if c-1 < 0 || grid[r][c-1] == 1 {
					c = -1
					break
				}
				c--
			}
		}
		result[col] = c
	}
	return result
}

func main() {
	fmt.Println(findBall([][]int{{1, 1, 1, -1, -1}, {1, 1, 1, -1, -1}, {-1, -1, -1, 1, 1}, {1, 1, 1, 1, -1}, {-1, -1, -1, -1, -1}}))
	// Expected: [1, -1, -1, -1, -1]

	fmt.Println(findBall([][]int{{-1}})) // Expected: [-1]

	fmt.Println(findBall([][]int{{1, 1, 1, 1, 1, 1}, {-1, -1, -1, -1, -1, -1}, {1, 1, 1, 1, 1, 1}, {-1, -1, -1, -1, -1, -1}}))
	// Expected: [0, 1, 2, 3, 4, -1]
}
```

## 1709 — Biggest Window Between Visits

```go
package main

// LeetCode #1709: Biggest Window Between Visits
// https://leetcode.com/problems/biggest-window-between-visits/
// Difficulty: Medium [Paid] (SQL)
// This is a SQL problem. In Go we simulate the logic.
// Time: O(n log n), Space: O(n)

import (
	"fmt"
	"sort"
)

func biggestWindow(visits []int) int {
	if len(visits) == 0 {
		return 0
	}
	sort.Ints(visits)
	maxGap := 0
	for i := 1; i < len(visits); i++ {
		gap := visits[i] - visits[i-1]
		if gap > maxGap {
			maxGap = gap
		}
	}
	return maxGap
}

func main() {
	fmt.Println(biggestWindow([]int{1, 3, 7, 10})) // Expected: 4 (between 3 and 7)
	fmt.Println(biggestWindow([]int{1, 2, 3, 4}))  // Expected: 1
	fmt.Println(biggestWindow([]int{5}))            // Expected: 0
}
```

## 1711 — Count Good Meals

```go
package main

// LeetCode #1711: Count Good Meals
// https://leetcode.com/problems/count-good-meals/
// Difficulty: Medium
// Time: O(n * log(maxVal)), Space: O(n)

import "fmt"

const mod = 1_000_000_007

func countPairs(deliciousness []int) int {
	count := make(map[int]int)
	result := 0

	for _, d := range deliciousness {
		// Check each power of 2
		for sum := 1; sum <= (1 << 21); sum <<= 1 {
			complement := sum - d
			if cnt, ok := count[complement]; ok {
				result = (result + cnt) % mod
			}
		}
		count[d]++
	}
	return result
}

func main() {
	fmt.Println(countPairs([]int{1, 3, 5, 7, 9}))    // Expected: 4
	fmt.Println(countPairs([]int{1, 1, 1, 3, 3, 3, 7})) // Expected: 15
	fmt.Println(countPairs([]int{2, 4, 6, 8})) // Expected: 3
}
```

## 1712 — Ways To Split Array Into Three Subarrays

```go
package main

// LeetCode #1712: Ways to Split Array Into Three Subarrays
// https://leetcode.com/problems/ways-to-split-array-into-three-subarrays/
// Difficulty: Medium
// Time: O(n), Space: O(n)

import "fmt"

const mod = 1_000_000_007

func waysToSplit(nums []int) int {
	n := len(nums)
	prefix := make([]int, n)
	prefix[0] = nums[0]
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] + nums[i]
	}

	total := prefix[n-1]
	result := 0
	left := 0

	// For each possible left split point i (first part ends at i)
	for i := 0; i < n-2; i++ {
		// Left sum = prefix[i]
		left += nums[i]

		// Find min j where mid sum >= left sum
		// mid sum = prefix[j] - prefix[i], need prefix[j] >= 2*prefix[i]
		minMid := lowerBound(prefix, 2*left, i+1)

		// Find max j where mid sum <= right sum
		// right sum = total - prefix[j], need prefix[j] <= (total + left) / 2
		maxMid := upperBound(prefix, (total+left)/2, i+1, n-2)

		if minMid <= maxMid {
			result = (result + (maxMid - minMid + 1)) % mod
		}
	}
	return result
}

func lowerBound(arr []int, target int, start int) int {
	lo, hi := start, len(arr)-2
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if arr[mid] >= target {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func upperBound(arr []int, target int, start int, end int) int {
	lo, hi := start, end
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if arr[mid] <= target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return hi
}

func main() {
	fmt.Println(waysToSplit([]int{1, 1, 1}))             // Expected: 1
	fmt.Println(waysToSplit([]int{1, 2, 2, 2, 5, 0}))   // Expected: 3
	fmt.Println(waysToSplit([]int{3, 2, 1}))             // Expected: 0
}
```

## 1715 — Count Apples And Oranges

```go
package main

// LeetCode #1715: Count Apples and Oranges
// https://leetcode.com/problems/count-apples-and-oranges/
// Difficulty: Medium [Paid] (SQL)
// This is a SQL problem. In Go we simulate the logic.
// Time: O(n), Space: O(n)

import "fmt"

type Box struct {
	BoxID      int
	AppleCnt   int
	OrangeCnt  int
}

type Chest struct {
	ChestID    int
	AppleCnt   int
	OrangeCnt  int
}

func countFruits(boxes []Box, chestMap map[int]Chest) (int, int) {
	totalApples := 0
	totalOranges := 0
	for _, b := range boxes {
		totalApples += b.AppleCnt
		totalOranges += b.OrangeCnt
		if c, ok := chestMap[b.BoxID]; ok {
			totalApples += c.AppleCnt
			totalOranges += c.OrangeCnt
		}
	}
	return totalApples, totalOranges
}

func main() {
	boxes := []Box{
		{BoxID: 1, AppleCnt: 5, OrangeCnt: 3},
		{BoxID: 2, AppleCnt: 2, OrangeCnt: 8},
	}
	chests := map[int]Chest{
		1: {AppleCnt: 10, OrangeCnt: 4},
	}
	apples, oranges := countFruits(boxes, chests)
	fmt.Printf("Apples: %d, Oranges: %d\n", apples, oranges) // Expected: 17, 15
}
```

## 1717 — Maximum Score From Removing Substrings

```go
package main

// LeetCode #1717: Maximum Score From Removing Substrings
// https://leetcode.com/problems/maximum-score-from-removing-substrings/
// Difficulty: Medium
// Time: O(n), Space: O(n)

import "fmt"

func maximumGain(s string, x int, y int) int {
	// Ensure we always process the higher-scoring pair first
	if y > x {
		s = reverse(s)
		x, y = y, x
	}

	ans := 0

	// First pass: remove "ab" for x points
	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 && stack[len(stack)-1] == 'a' && s[i] == 'b' {
			stack = stack[:len(stack)-1]
			ans += x
		} else {
			stack = append(stack, s[i])
		}
	}

	// Second pass: remove "ba" for y points from remaining
	stack2 := make([]byte, 0, len(stack))
	for i := 0; i < len(stack); i++ {
		if len(stack2) > 0 && stack2[len(stack2)-1] == 'b' && stack[i] == 'a' {
			stack2 = stack2[:len(stack2)-1]
			ans += y
		} else {
			stack2 = append(stack2, stack[i])
		}
	}

	return ans
}

func reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func main() {
	fmt.Println(maximumGain("cdbcbbaaabab", 4, 5)) // Expected: 19
	fmt.Println(maximumGain("aabbaaxybbaabb", 5, 4)) // Expected: 20
	fmt.Println(maximumGain("ab", 1, 2)) // Expected: 1
}
```

## 1718 — Construct The Lexicographically Largest Valid Sequence

```go
package main

// LeetCode #1718: Construct the Lexicographically Largest Valid Sequence
// https://leetcode.com/problems/construct-the-lexicographically-largest-valid-sequence/
// Difficulty: Medium
// Time: O(n!), Space: O(n) for backtracking

import "fmt"

func constructDistancedSequence(n int) []int {
	length := 2*n - 1
	result := make([]int, length)
	used := make([]bool, n+1)

	var backtrack func(pos int) bool
	backtrack = func(pos int) bool {
		if pos == length {
			return true
		}
		if result[pos] != 0 {
			return backtrack(pos + 1)
		}

		// Try largest number first for lexicographically largest
		for num := n; num >= 1; num-- {
			if used[num] {
				continue
			}
			if num == 1 {
				result[pos] = 1
				used[1] = true
				if backtrack(pos + 1) {
					return true
				}
				result[pos] = 0
				used[1] = false
			} else {
				nextPos := pos + num
				if nextPos < length && result[nextPos] == 0 {
					result[pos] = num
					result[nextPos] = num
					used[num] = true
					if backtrack(pos + 1) {
						return true
					}
					result[pos] = 0
					result[nextPos] = 0
					used[num] = false
				}
			}
		}
		return false
	}

	backtrack(0)
	return result
}

func main() {
	fmt.Println(constructDistancedSequence(3)) // Expected: [3, 1, 2, 3, 2]
	fmt.Println(constructDistancedSequence(5)) // Expected: [5, 3, 1, 4, 3, 5, 2, 4, 2]
	fmt.Println(constructDistancedSequence(1)) // Expected: [1]
}
```

## 1721 — Swapping Nodes In A Linked List

```go
package main

// LeetCode #1721: Swapping Nodes in a Linked List
// https://leetcode.com/problems/swapping-nodes-in-a-linked-list/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapNodes(head *ListNode, k int) *ListNode {
	// First pass: find kth from beginning
	first := head
	for i := 1; i < k; i++ {
		first = first.Next
	}

	// Two pointer approach for kth from end
	slow := head
	fast := first
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	// Swap values
	first.Val, slow.Val = slow.Val, first.Val
	return head
}

func makeList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	curr := head
	for i := 1; i < len(vals); i++ {
		curr.Next = &ListNode{Val: vals[i]}
		curr = curr.Next
	}
	return head
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func main() {
	l1 := makeList([]int{1, 2, 3, 4, 5})
	printList(swapNodes(l1, 2)) // Expected: 1 4 3 2 5

	l2 := makeList([]int{7, 9, 6, 6, 7, 8, 3, 0, 9, 5})
	printList(swapNodes(l2, 5)) // Expected: 7 9 6 6 8 7 3 0 9 5

	l3 := makeList([]int{1, 2})
	printList(swapNodes(l3, 1)) // Expected: 2 1
}
```

## 1722 — Minimize Hamming Distance After Swap Operations

```go
package main

// LeetCode #1722: Minimize Hamming Distance After Swap Operations
// https://leetcode.com/problems/minimize-hamming-distance-after-swap-operations/
// Difficulty: Medium
// Time: O(n + swaps), Space: O(n)

import "fmt"

type UnionFind struct {
	parent []int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &UnionFind{parent}
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) {
	px, py := uf.Find(x), uf.Find(y)
	if px != py {
		uf.parent[px] = py
	}
}

func minimumHammingDistance(source []int, target []int, allowedSwaps [][]int) int {
	n := len(source)
	uf := NewUnionFind(n)

	for _, swap := range allowedSwaps {
		uf.Union(swap[0], swap[1])
	}

	// Group indices by component
	groups := make(map[int][]int)
	for i := 0; i < n; i++ {
		root := uf.Find(i)
		groups[root] = append(groups[root], i)
	}

	hamming := 0
	for _, indices := range groups {
		counts := make(map[int]int)
		for _, idx := range indices {
			counts[source[idx]]++
		}
		for _, idx := range indices {
			if counts[target[idx]] > 0 {
				counts[target[idx]]--
			} else {
				hamming++
			}
		}
	}
	return hamming
}

func main() {
	fmt.Println(minimumHammingDistance([]int{1, 2, 3, 4}, []int{2, 1, 4, 5}, [][]int{{0, 1}, {2, 3}})) // Expected: 1
	fmt.Println(minimumHammingDistance([]int{1, 2, 3, 4}, []int{1, 3, 2, 4}, [][]int{})) // Expected: 2
	fmt.Println(minimumHammingDistance([]int{5, 1, 2, 4, 3}, []int{1, 5, 4, 2, 3}, [][]int{{0, 4}, {4, 2}, {1, 3}, {1, 4}})) // Expected: 0
}
```

## 1726 — Tuple With Same Product

```go
package main

// LeetCode #1726: Tuple with Same Product
// https://leetcode.com/problems/tuple-with-same-product/
// Difficulty: Medium
// Time: O(n^2), Space: O(n^2)

import "fmt"

func tupleSameProduct(nums []int) int {
	n := len(nums)
	productCount := make(map[int]int)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			product := nums[i] * nums[j]
			productCount[product]++
		}
	}

	result := 0
	for _, count := range productCount {
		if count > 1 {
			// Each pair of pairs = 8 tuples (4! / 3 = 8)
			result += count * (count - 1) / 2 * 8
		}
	}
	return result
}

func main() {
	fmt.Println(tupleSameProduct([]int{2, 3, 4, 6}))     // Expected: 8
	fmt.Println(tupleSameProduct([]int{1, 2, 4, 5, 10})) // Expected: 16
	fmt.Println(tupleSameProduct([]int{1, 2, 3, 4, 6, 12})) // Expected: 40
}
```

## 1727 — Largest Submatrix With Rearrangements

```go
package main

// LeetCode #1727: Largest Submatrix With Rearrangements
// https://leetcode.com/problems/largest-submatrix-with-rearrangements/
// Difficulty: Medium
// Time: O(m * n log n), Space: O(n)

import (
	"fmt"
	"sort"
)

func largestSubmatrix(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	maxArea := 0

	heights := make([]int, n)

	for r := 0; r < m; r++ {
		// Update heights
		for c := 0; c < n; c++ {
			if matrix[r][c] == 1 {
				heights[c]++
			} else {
				heights[c] = 0
			}
		}

		// Sort heights for this row (to find max rectangle that can be formed
		// by rearranging columns)
		sorted := make([]int, n)
		copy(sorted, heights)
		sort.Slice(sorted, func(i, j int) bool {
			return sorted[i] > sorted[j]
		})

		// For each column, area = height * (col index) because it's sorted
		for c := 0; c < n; c++ {
			area := sorted[c] * (c + 1)
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

func main() {
	fmt.Println(largestSubmatrix([][]int{{0, 0, 1}, {1, 1, 1}, {1, 0, 1}})) // Expected: 4
	fmt.Println(largestSubmatrix([][]int{{1, 0, 1, 0, 1}})) // Expected: 3
	fmt.Println(largestSubmatrix([][]int{{1, 1, 0}, {1, 0, 1}})) // Expected: 2
}
```

## 1730 — Shortest Path To Get Food

```go
package main

// LeetCode #1730: Shortest Path to Get Food
// https://leetcode.com/problems/shortest-path-to-get-food/
// Difficulty: Medium [Paid]
// Time: O(m * n), Space: O(m * n)

import "fmt"

func getFood(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	startR, startC := 0, 0

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if grid[r][c] == '*' {
				startR, startC = r, c
			}
		}
	}

	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	queue := [][2]int{{startR, startC}}
	visited[startR][startC] = true
	steps := 0

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			r, c := queue[i][0], queue[i][1]
			if grid[r][c] == '#' {
				return steps
			}
			for _, d := range dirs {
				nr, nc := r+d[0], c+d[1]
				if nr >= 0 && nr < m && nc >= 0 && nc < n && !visited[nr][nc] && grid[nr][nc] != 'X' {
					visited[nr][nc] = true
					queue = append(queue, [2]int{nr, nc})
				}
			}
		}
		queue = queue[size:]
		steps++
	}
	return -1
}

func main() {
	grid1 := [][]byte{
		{'X', 'X', 'X', 'X', 'X', 'X'},
		{'X', '*', 'O', 'O', 'O', 'X'},
		{'X', 'O', 'O', '#', 'O', 'X'},
		{'X', 'X', 'X', 'X', 'X', 'X'},
	}
	fmt.Println(getFood(grid1)) // Expected: 3

	grid2 := [][]byte{
		{'X', 'X', 'X', 'X', 'X'},
		{'X', '*', 'X', 'O', 'X'},
		{'X', 'O', 'X', '#', 'X'},
		{'X', 'X', 'X', 'X', 'X'},
	}
	fmt.Println(getFood(grid2)) // Expected: -1

	grid3 := [][]byte{
		{'*', 'O', '#'},
	}
	fmt.Println(getFood(grid3)) // Expected: 2
}
```

## 1733 — Minimum Number Of People To Teach

```go
package main

// LeetCode #1733: Minimum Number of People to Teach
// https://leetcode.com/problems/minimum-number-of-people-to-teach/
// Difficulty: Medium
// Time: O(n * m + f) where n = user count, m = avg languages per user, f = friend pairs

import "fmt"

func minimumTeachings(n int, languages [][]int, friendships [][]int) int {
	langSet := make([]map[int]bool, len(languages))
	for i, langs := range languages {
		langSet[i] = make(map[int]bool)
		for _, l := range langs {
			langSet[i][l] = true
		}
	}

	// Find users who cannot communicate
	cannotCommunicate := make([]bool, len(languages))
	for _, f := range friendships {
		u, v := f[0]-1, f[1]-1
		canCommunicate := false
		for l := range langSet[u] {
			if langSet[v][l] {
				canCommunicate = true
				break
			}
		}
		if !canCommunicate {
			cannotCommunicate[u] = true
			cannotCommunicate[v] = true
		}
	}

	// Find the most common language among users who cannot communicate
	langCount := make(map[int]int)
	for i, cn := range cannotCommunicate {
		if cn {
			for l := range langSet[i] {
				langCount[l]++
			}
		}
	}

	// Count users who don't know the most common language
	maxLang := 0
	for _, c := range langCount {
		if c > maxLang {
			maxLang = c
		}
	}

	totalCannot := 0
	for _, cn := range cannotCommunicate {
		if cn {
			totalCannot++
		}
	}

	return totalCannot - maxLang
}

func main() {
	fmt.Println(minimumTeachings(2, [][]int{{1}, {2}, {1, 2}}, [][]int{{1, 2}, {1, 3}, {2, 3}})) // Expected: 1
	fmt.Println(minimumTeachings(3, [][]int{{2}, {1, 3}, {1, 2}, {3}}, [][]int{{1, 4}, {1, 2}, {3, 4}, {2, 3}})) // Expected: 2
	fmt.Println(minimumTeachings(2, [][]int{{1}, {2}}, [][]int{{1, 2}})) // Expected: 1
}
```

## 1734 — Decode Xored Permutation

```go
package main

// LeetCode #1734: Decode XORed Permutation
// https://leetcode.com/problems/decode-xored-permutation/
// Difficulty: Medium
// Time: O(n), Space: O(n)

import "fmt"

func decode(encoded []int) []int {
	n := len(encoded) + 1

	// XOR of all numbers from 1 to n
	totalXor := 0
	for i := 1; i <= n; i++ {
		totalXor ^= i
	}

	// XOR of encoded[1], encoded[3], encoded[5], ...
	xorOdd := 0
	for i := 1; i < len(encoded); i += 2 {
		xorOdd ^= encoded[i]
	}

	// First element = totalXor ^ xorOdd
	perm := make([]int, n)
	perm[0] = totalXor ^ xorOdd

	// Decode the rest
	for i := 1; i < n; i++ {
		perm[i] = perm[i-1] ^ encoded[i-1]
	}

	return perm
}

func main() {
	fmt.Println(decode([]int{3, 1}))          // Expected: [1, 2, 3]
	fmt.Println(decode([]int{6, 5, 4, 6}))    // Expected: [2, 4, 1, 5, 3]
	fmt.Println(decode([]int{5, 6, 1, 6, 2, 1})) // Expected: [1 4 2 3 5 7 6]
}
```

## 1737 — Change Minimum Characters To Satisfy One Of Three Conditions

```go
package main

// LeetCode #1737: Change Minimum Characters to Satisfy One of Three Conditions
// https://leetcode.com/problems/change-minimum-characters-to-satisfy-one-of-three-conditions/
// Difficulty: Medium
// Time: O(n + m), Space: O(26)

import "fmt"

func minCharacters(a string, b string) int {
	countA := make([]int, 26)
	countB := make([]int, 26)

	for _, ch := range a {
		countA[ch-'a']++
	}
	for _, ch := range b {
		countB[ch-'a']++
	}

	m, n := len(a), len(b)

	// Condition 3: make all characters in both strings the same
	ans := m + n
	for i := 0; i < 26; i++ {
		changes := (m - countA[i]) + (n - countB[i])
		if changes < ans {
			ans = changes
		}
	}

	// Condition 1: a < b lexicographically (every char in a < every char in b)
	// Condition 2: b < a lexicographically (every char in b < every char in a)
	for condition := 0; condition < 2; condition++ {
		prefixA := 0
		prefixB := 0
		for i := 0; i < 25; i++ {
			prefixA += countA[i]
			prefixB += countB[i]
			if condition == 0 {
				// a < b: change all a[i..25] to some smaller char, change all b[0..i] to some larger char
				changes := (m - prefixA) + prefixB
				if changes < ans {
					ans = changes
				}
			} else {
				// b < a: change all b[i..25] to some smaller char, change all a[0..i] to some larger char
				changes := (n - prefixB) + prefixA
				if changes < ans {
					ans = changes
				}
			}
		}
	}

	return ans
}

func main() {
	fmt.Println(minCharacters("aba", "caa")) // Expected: 2
	fmt.Println(minCharacters("dabadd", "cda")) // Expected: 3
	fmt.Println(minCharacters("a", "a")) // Expected: 2
}
```

## 1738 — Find Kth Largest Xor Coordinate Value

```go
package main

// LeetCode #1738: Find Kth Largest XOR Coordinate Value
// https://leetcode.com/problems/find-kth-largest-xor-coordinate-value/
// Difficulty: Medium
// Time: O(m * n * log(m*n)), Space: O(m * n)

import (
	"fmt"
	"sort"
)

func kthLargestValue(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])
	prefix := make([][]int, m)
	for i := 0; i < m; i++ {
		prefix[i] = make([]int, n)
	}

	values := make([]int, 0, m*n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			val := matrix[i][j]
			if i > 0 {
				val ^= prefix[i-1][j]
			}
			if j > 0 {
				val ^= prefix[i][j-1]
			}
			if i > 0 && j > 0 {
				val ^= prefix[i-1][j-1]
			}
			prefix[i][j] = val
			values = append(values, val)
		}
	}

	sort.Slice(values, func(i, j int) bool {
		return values[i] > values[j]
	})
	return values[k-1]
}

func main() {
	fmt.Println(kthLargestValue([][]int{{5, 2}, {1, 6}}, 1)) // Expected: 7
	fmt.Println(kthLargestValue([][]int{{5, 2}, {1, 6}}, 2)) // Expected: 5
	fmt.Println(kthLargestValue([][]int{{5, 2}, {1, 6}}, 3)) // Expected: 4
}
```

## 1740 — Find Distance In A Binary Tree

```go
package main

// LeetCode #1740: Find Distance in a Binary Tree
// https://leetcode.com/problems/find-distance-in-a-binary-tree/
// Difficulty: Medium [Paid]
// Time: O(n), Space: O(h)

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findDistance(root *TreeNode, p int, q int) int {
	if p == q {
		return 0
	}
	lca := findLCA(root, p, q)
	return distFrom(lca, p, 0) + distFrom(lca, q, 0)
}

func findLCA(node *TreeNode, p, q int) *TreeNode {
	if node == nil || node.Val == p || node.Val == q {
		return node
	}
	left := findLCA(node.Left, p, q)
	right := findLCA(node.Right, p, q)
	if left != nil && right != nil {
		return node
	}
	if left != nil {
		return left
	}
	return right
}

func distFrom(node *TreeNode, target, dist int) int {
	if node == nil {
		return -1
	}
	if node.Val == target {
		return dist
	}
	if d := distFrom(node.Left, target, dist+1); d != -1 {
		return d
	}
	return distFrom(node.Right, target, dist+1)
}

func main() {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 1}
	root.Left.Left = &TreeNode{Val: 6}
	root.Left.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 0}
	root.Right.Right = &TreeNode{Val: 8}
	root.Left.Right.Left = &TreeNode{Val: 7}
	root.Left.Right.Right = &TreeNode{Val: 4}

	fmt.Println(findDistance(root, 5, 0)) // Expected: 3
	fmt.Println(findDistance(root, 5, 7)) // Expected: 2
	fmt.Println(findDistance(root, 6, 4)) // Expected: 3
}
```

## 1743 — Restore The Array From Adjacent Pairs

```go
package main

// LeetCode #1743: Restore the Array From Adjacent Pairs
// https://leetcode.com/problems/restore-the-array-from-adjacent-pairs/
// Difficulty: Medium
// Time: O(n), Space: O(n)

import "fmt"

func restoreArray(adjacentPairs [][]int) []int {
	graph := make(map[int][]int)
	for _, pair := range adjacentPairs {
		u, v := pair[0], pair[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	// Find the first element (has only 1 neighbor)
	start := 0
	for node, neighbors := range graph {
		if len(neighbors) == 1 {
			start = node
			break
		}
	}

	n := len(adjacentPairs) + 1
	result := make([]int, n)
	result[0] = start
	result[1] = graph[start][0]

	for i := 2; i < n; i++ {
		neighbors := graph[result[i-1]]
		if neighbors[0] == result[i-2] {
			result[i] = neighbors[1]
		} else {
			result[i] = neighbors[0]
		}
	}
	return result
}

func main() {
	fmt.Println(restoreArray([][]int{{2, 1}, {3, 4}, {3, 2}})) // Expected: [1, 2, 3, 4]
	fmt.Println(restoreArray([][]int{{4, -2}, {1, 4}, {-3, 1}})) // Expected: [-2, 4, 1, -3]
	fmt.Println(restoreArray([][]int{{100, -100}})) // Expected: [100, -100]
}
```

## 1744 — Can You Eat Your Favorite Candy On Your Favorite Day

```go
package main

// LeetCode #1744: Can You Eat Your Favorite Candy on Your Favorite Day?
// https://leetcode.com/problems/can-you-eat-your-favorite-candy-on-your-favorite-day/
// Difficulty: Medium
// Time: O(n + q), Space: O(n)

import "fmt"

func canEat(candiesCount []int, queries [][]int) []bool {
	n := len(candiesCount)
	prefix := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + candiesCount[i]
	}

	result := make([]bool, len(queries))
	for i, q := range queries {
		favType, favDay, dailyCap := q[0], q[1], q[2]

		// Earliest day we can eat favType candy (at dailyCap per day)
		minDay := prefix[favType] / dailyCap
		// Latest day we can eat favType candy (at 1 per day)
		maxDay := prefix[favType+1] - 1

		result[i] = favDay >= minDay && favDay <= maxDay
	}
	return result
}

func main() {
	fmt.Println(canEat([]int{7, 4, 5, 3, 8}, [][]int{{0, 2, 2}, {4, 2, 4}, {2, 13, 100}}))
	// Expected: [true, false, true]

	fmt.Println(canEat([]int{5, 2, 6, 4, 1}, [][]int{{3, 1, 2}, {4, 10, 3}, {3, 10, 100}, {0, 5, 1}}))
	// Expected: [false, true, true, false]

	fmt.Println(canEat([]int{16, 38, 8, 41, 30, 31, 14, 45, 3, 2, 24, 23, 38, 30, 4, 43}, [][]int{{8, 19, 38}}))
}
```

## 1746 — Maximum Subarray Sum After One Operation

```go
package main

// LeetCode #1746: Maximum Subarray Sum After One Operation
// https://leetcode.com/problems/maximum-subarray-sum-after-one-operation/
// Difficulty: Medium [Paid]
// Time: O(n), Space: O(1)

import "fmt"

func maxSumAfterOperation(nums []int) int {
	// dp0: max subarray sum without using operation
	// dp1: max subarray sum with exactly one operation used
	dp0 := 0
	dp1 := 0
	maxSum := nums[0] * nums[0]

	for _, v := range nums {
		// Either start new or extend
		newDp0 := max(v, dp0+v)
		newDp1 := max(v*v, dp0+v*v, dp1+v)

		dp0 = newDp0
		dp1 = newDp1
		maxSum = max(maxSum, dp1)
	}
	return maxSum
}

func max(nums ...int) int {
	result := nums[0]
	for _, v := range nums[1:] {
		if v > result {
			result = v
		}
	}
	return result
}

func main() {
	fmt.Println(maxSumAfterOperation([]int{2, -1, -4, -3})) // Expected: 17
	fmt.Println(maxSumAfterOperation([]int{1, -2, 3, 4}))   // Expected: 19 ([3,4] with 4^2=16: 3+16=19)

	fmt.Println(maxSumAfterOperation([]int{-1, -1, -1})) // Expected: 1 (replace -1 with 1)
}
```

## 1747 — Leetflex Banned Accounts

```go
package main

// LeetCode #1747: Leetflex Banned Accounts
// https://leetcode.com/problems/leetflex-banned-accounts/
// Difficulty: Medium [Paid] (SQL)
// This is a SQL problem. In Go we simulate the logic.
// Time: O(n log n), Space: O(n)

import (
	"fmt"
	"sort"
)

type Login struct {
	AccountID int
	IPAddress string
	LoginTime int
}

func findBanned(logins []Login) []int {
	// Group by account
	groups := make(map[int][]Login)
	for _, l := range logins {
		groups[l.AccountID] = append(groups[l.AccountID], l)
	}

	banned := make(map[int]bool)
	for accID, records := range groups {
		sort.Slice(records, func(i, j int) bool {
			return records[i].LoginTime < records[j].LoginTime
		})
		// Track latest login time per IP for this account
		lastTime := make(map[string]int)
		for _, r := range records {
			if prevTime, ok := lastTime[r.IPAddress]; ok {
				// Same IP, update last time
				_ = prevTime
			}
			lastTime[r.IPAddress] = r.LoginTime
		}

		// Check if any IP has concurrent sessions
		active := make(map[string]int)
		for _, r := range records {
			if _, ok := active[r.IPAddress]; ok {
				// Check if there's a different IP with an active session
				for ip, t := range active {
					if ip != r.IPAddress && t < r.LoginTime {
						banned[accID] = true
					}
				}
				// End current session for this IP and start new one
				delete(active, r.IPAddress)
			}
			active[r.IPAddress] = r.LoginTime
		}
	}

	result := make([]int, 0, len(banned))
	for id := range banned {
		result = append(result, id)
	}
	sort.Ints(result)
	return result
}

func main() {
	logins := []Login{
		{1, "IP1", 1},
		{1, "IP2", 2},
		{1, "IP1", 3},
	}
	fmt.Println(findBanned(logins))
}
```

## 1749 — Maximum Absolute Sum Of Any Subarray

```go
package main

// LeetCode #1749: Maximum Absolute Sum of Any Subarray
// https://leetcode.com/problems/maximum-absolute-sum-of-any-subarray/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

func maxAbsoluteSum(nums []int) int {
	maxEnding := 0
	minEnding := 0
	maxSoFar := 0

	for _, v := range nums {
		maxEnding = max(0, maxEnding+v)
		minEnding = min(0, minEnding+v)
		maxSoFar = max(maxSoFar, maxEnding, -minEnding)
	}
	return maxSoFar
}

func max(nums ...int) int {
	r := nums[0]
	for _, v := range nums[1:] {
		if v > r {
			r = v
		}
	}
	return r
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxAbsoluteSum([]int{1, -3, 2, 3, -4})) // Expected: 5
	fmt.Println(maxAbsoluteSum([]int{2, -5, 1, -4, 3, -2})) // Expected: 8
	fmt.Println(maxAbsoluteSum([]int{-1})) // Expected: 1
}
```

## 1750 — Minimum Length Of String After Deleting Similar Ends

```go
package main

// LeetCode #1750: Minimum Length of String After Deleting Similar Ends
// https://leetcode.com/problems/minimum-length-of-string-after-deleting-similar-ends/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

func minimumLength(s string) int {
	left, right := 0, len(s)-1

	for left < right && s[left] == s[right] {
		ch := s[left]
		// Delete from left
		for left <= right && s[left] == ch {
			left++
		}
		// Delete from right
		for left <= right && s[right] == ch {
			right--
		}
	}
	return right - left + 1
}

func main() {
	fmt.Println(minimumLength("ca"))            // Expected: 2
	fmt.Println(minimumLength("cabaabac"))      // Expected: 0
	fmt.Println(minimumLength("aabccabba"))     // Expected: 3
}
```

## 1753 — Maximum Score From Removing Stones

```go
package main

// LeetCode #1753: Maximum Score From Removing Stones
// https://leetcode.com/problems/maximum-score-from-removing-stones/
// Difficulty: Medium
// Time: O(1), Space: O(1)

import (
	"fmt"
	"sort"
)

func maximumScore(a int, b int, c int) int {
	piles := []int{a, b, c}
	sort.Ints(piles)
	// If the largest pile is >= sum of other two, we can only take sum of those two
	if piles[2] >= piles[0]+piles[1] {
		return piles[0] + piles[1]
	}
	// Otherwise, we can take (a+b+c)/2 stones
	return (a + b + c) / 2
}

func main() {
	fmt.Println(maximumScore(2, 4, 6)) // Expected: 6
	fmt.Println(maximumScore(4, 4, 6)) // Expected: 7
	fmt.Println(maximumScore(1, 8, 8)) // Expected: 8
}
```

## 1754 — Largest Merge Of Two Strings

```go
package main

// LeetCode #1754: Largest Merge Of Two Strings
// https://leetcode.com/problems/largest-merge-of-two-strings/
// Difficulty: Medium
// Time: O(n^2), Space: O(n)

import "fmt"

func largestMerge(word1 string, word2 string) string {
	result := make([]byte, 0, len(word1)+len(word2))
	i, j := 0, 0

	for i < len(word1) && j < len(word2) {
		// Pick the character that leads to a larger overall string
		if word1[i:] > word2[j:] {
			result = append(result, word1[i])
			i++
		} else {
			result = append(result, word2[j])
			j++
		}
	}

	// Append remaining characters
	result = append(result, word1[i:]...)
	result = append(result, word2[j:]...)
	return string(result)
}

func main() {
	fmt.Println(largestMerge("cabaa", "bcaaa")) // Expected: "cbcabaaaaa"
	fmt.Println(largestMerge("abcabc", "abdcaba")) // Expected: "abdcabcabcaba"
	fmt.Println(largestMerge("a", "b")) // Expected: "ba"
}
```

## 1756 — Design Most Recently Used Queue

```go
package main

// LeetCode #1756: Design Most Recently Used Queue
// https://leetcode.com/problems/design-most-recently-used-queue/
// Difficulty: Medium [Paid]
// Time: O(n) per operation, Space: O(n)

import "fmt"

type MRUQueue struct {
	data []int
}

func Constructor(n int) MRUQueue {
	data := make([]int, n)
	for i := 0; i < n; i++ {
		data[i] = i + 1
	}
	return MRUQueue{data}
}

func (q *MRUQueue) Fetch(k int) int {
	// 1-indexed, fetch kth element and move to end
	val := q.data[k-1]
	// Remove
	q.data = append(q.data[:k-1], q.data[k:]...)
	// Move to end
	q.data = append(q.data, val)
	return val
}

func main() {
	q := Constructor(8)
	fmt.Println(q.Fetch(3)) // Expected: 3
	fmt.Println(q.Fetch(5)) // Expected: 6
	fmt.Println(q.Fetch(2)) // Expected: 2
	fmt.Println(q.Fetch(8)) // Expected: 3
}
```

## 1759 — Count Number Of Homogenous Substrings

```go
package main

// LeetCode #1759: Count Number of Homogenous Substrings
// https://leetcode.com/problems/count-number-of-homogenous-substrings/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

const mod = 1_000_000_007

func countHomogenous(s string) int {
	result := 0
	count := 0

	for i := 0; i < len(s); i++ {
		if i > 0 && s[i] == s[i-1] {
			count++
		} else {
			count = 1
		}
		result = (result + count) % mod
	}
	return result
}

func main() {
	fmt.Println(countHomogenous("abbcccaa")) // Expected: 13
	fmt.Println(countHomogenous("xy"))        // Expected: 2
	fmt.Println(countHomogenous("zzzzz"))     // Expected: 15
}
```

## 1760 — Minimum Limit Of Balls In A Bag

```go
package main

// LeetCode #1760: Minimum Limit of Balls in a Bag
// https://leetcode.com/problems/minimum-limit-of-balls-in-a-bag/
// Difficulty: Medium
// Time: O(n log M) where M = max(nums), Space: O(1)

import "fmt"

func minimumSize(nums []int, maxOperations int) int {
	left, right := 1, 0
	for _, v := range nums {
		if v > right {
			right = v
		}
	}

	for left < right {
		mid := left + (right-left)/2
		if canDivide(nums, maxOperations, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func canDivide(nums []int, maxOps, limit int) bool {
	ops := 0
	for _, v := range nums {
		if v > limit {
			ops += (v - 1) / limit
			if ops > maxOps {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(minimumSize([]int{9}, 2))               // Expected: 3
	fmt.Println(minimumSize([]int{2, 4, 8, 2}, 4))      // Expected: 2
	fmt.Println(minimumSize([]int{7, 17}, 2))            // Expected: 7
}
```

## 1762 — Buildings With An Ocean View

```go
package main

// LeetCode #1762: Buildings With an Ocean View
// https://leetcode.com/problems/buildings-with-an-ocean-view/
// Difficulty: Medium [Paid]
// Time: O(n), Space: O(1) excluding output

import "fmt"

func findBuildings(heights []int) []int {
	result := make([]int, 0)
	maxHeight := -1

	for i := len(heights) - 1; i >= 0; i-- {
		if heights[i] > maxHeight {
			result = append([]int{i}, result...)
			maxHeight = heights[i]
		}
	}
	return result
}

func main() {
	fmt.Println(findBuildings([]int{4, 2, 3, 1}))    // Expected: [0, 2, 3]
	fmt.Println(findBuildings([]int{4, 3, 2, 1}))    // Expected: [0, 1, 2, 3]
	fmt.Println(findBuildings([]int{1, 3, 2, 4}))    // Expected: [3]
}
```

## 1764 — Form Array By Concatenating Subarrays Of Another Array

```go
package main

// LeetCode #1764: Form Array by Concatenating Subarrays of Another Array
// https://leetcode.com/problems/form-array-by-concatenating-subarrays-of-another-array/
// Difficulty: Medium
// Time: O(n * m), Space: O(1)

import "fmt"

func canChoose(groups [][]int, nums []int) bool {
	idx := 0
	for _, group := range groups {
		found := false
		for idx <= len(nums)-len(group) {
			if matches(nums, group, idx) {
				idx += len(group)
				found = true
				break
			}
			idx++
		}
		if !found {
			return false
		}
	}
	return true
}

func matches(nums []int, group []int, start int) bool {
	for i, v := range group {
		if nums[start+i] != v {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(canChoose([][]int{{1, -1, -1}, {3, -2, 0}}, []int{1, -1, 0, 1, -1, -1, 3, -2, 0})) // Expected: true
	fmt.Println(canChoose([][]int{{10, -2}, {1, 2, 3, 4}}, []int{1, 2, 3, 4, 10, -2})) // Expected: false
	fmt.Println(canChoose([][]int{{1, 2, 3}, {3, 4}}, []int{7, 7, 1, 2, 3, 4, 7, 7})) // Expected: false
}
```

## 1765 — Map Of Highest Peak

```go
package main

// LeetCode #1765: Map of Highest Peak
// https://leetcode.com/problems/map-of-highest-peak/
// Difficulty: Medium
// Time: O(m * n), Space: O(m * n)

import "fmt"

func highestPeak(isWater [][]int) [][]int {
	m, n := len(isWater), len(isWater[0])
	result := make([][]int, m)
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
		for j := 0; j < n; j++ {
			result[i][j] = -1
		}
	}

	queue := make([][2]int, 0, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if isWater[i][j] == 1 {
				result[i][j] = 0
				queue = append(queue, [2]int{i, j})
			}
		}
	}

	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for len(queue) > 0 {
		size := len(queue)
		for k := 0; k < size; k++ {
			r, c := queue[k][0], queue[k][1]
			for _, d := range dirs {
				nr, nc := r+d[0], c+d[1]
				if nr >= 0 && nr < m && nc >= 0 && nc < n && result[nr][nc] == -1 {
					result[nr][nc] = result[r][c] + 1
					queue = append(queue, [2]int{nr, nc})
				}
			}
		}
		queue = queue[size:]
	}
	return result
}

func main() {
	fmt.Println(highestPeak([][]int{{0, 1}, {0, 0}})) // Expected: [[1,0],[2,1]]
	fmt.Println(highestPeak([][]int{{0, 0, 1}, {1, 0, 0}, {0, 0, 0}})) // Expected: [[1,1,0],[0,1,1],[1,2,2]]
}
```

## 1769 — Minimum Number Of Operations To Move All Balls To Each Box

```go
package main

// LeetCode #1769: Minimum Number of Operations to Move All Balls to Each Box
// https://leetcode.com/problems/minimum-number-of-operations-to-move-all-balls-to-each-box/
// Difficulty: Medium
// Time: O(n), Space: O(1) excluding output

import "fmt"

func minOperations(boxes string) []int {
	n := len(boxes)
	result := make([]int, n)

	// Left to right: count balls and accumulate moves
	balls := 0
	moves := 0
	for i := 0; i < n; i++ {
		result[i] += moves
		if boxes[i] == '1' {
			balls++
		}
		moves += balls
	}

	// Right to left
	balls = 0
	moves = 0
	for i := n - 1; i >= 0; i-- {
		result[i] += moves
		if boxes[i] == '1' {
			balls++
		}
		moves += balls
	}

	return result
}

func main() {
	fmt.Println(minOperations("110"))    // Expected: [1, 1, 3]
	fmt.Println(minOperations("001011")) // Expected: [11, 8, 5, 4, 3, 4]
	fmt.Println(minOperations("0"))      // Expected: [0]
}
```

## 1772 — Sort Features By Popularity

```go
package main

// LeetCode #1772: Sort Features by Popularity
// https://leetcode.com/problems/sort-features-by-popularity/
// Difficulty: Medium [Paid]
// Time: O(f * r), Space: O(f)

import (
	"fmt"
	"sort"
)

func sortFeatures(features []string, responses []string) []string {
	featureRank := make(map[string]int)
	for i, f := range features {
		featureRank[f] = i
	}

	freq := make(map[string]int)
	for _, resp := range responses {
		seen := make(map[string]bool)
		word := ""
		for _, ch := range resp + " " {
			if ch == ' ' {
				if word != "" && !seen[word] {
					freq[word]++
					seen[word] = true
				}
				word = ""
			} else {
				word += string(ch)
			}
		}
	}

	sorted := make([]string, len(features))
	copy(sorted, features)
	sort.SliceStable(sorted, func(i, j int) bool {
		fi, fj := freq[sorted[i]], freq[sorted[j]]
		if fi != fj {
			return fi > fj
		}
		return featureRank[sorted[i]] < featureRank[sorted[j]]
	})
	return sorted
}

func main() {
	fmt.Println(sortFeatures(
		[]string{"cooler", "lock", "touch"},
		[]string{"i like cooler cooler", "lock touch cool", "locker like touch"},
	)) // Expected: ["touch", "cooler", "lock"] or ["touch", "lock", "cooler"] depending on frequency

	fmt.Println(sortFeatures(
		[]string{"a", "b", "c"},
		[]string{"a b", "b c", "c a"},
	)) // Expected: ["a", "b", "c"] (all appear in 2 responses, stable sort by original order)
}
```

## 1774 — Closest Dessert Cost

```go
package main

// LeetCode #1774: Closest Dessert Cost
// https://leetcode.com/problems/closest-dessert-cost/
// Difficulty: Medium
// Time: O(b * 3^t), Space: O(t)

import "fmt"

func closestCost(baseCosts []int, toppingCosts []int, target int) int {
	best := baseCosts[0]
	for _, b := range baseCosts {
		best = minDiff(best, b, target)
		dfs(toppingCosts, 0, b, target, &best)
	}
	return best
}

func dfs(toppings []int, idx int, current int, target int, best *int) {
	if idx == len(toppings) {
		*best = minDiff(*best, current, target)
		return
	}
	// Try 0, 1, or 2 of each topping
	for count := 0; count <= 2; count++ {
		dfs(toppings, idx+1, current+count*toppings[idx], target, best)
	}
}

func minDiff(a, b, target int) int {
	diffA := abs(a - target)
	diffB := abs(b - target)
	if diffA < diffB || (diffA == diffB && a < b) {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(closestCost([]int{1, 7}, []int{3, 4}, 10)) // Expected: 10
	fmt.Println(closestCost([]int{2, 3}, []int{4, 5, 100}, 18)) // Expected: 17
	fmt.Println(closestCost([]int{10}, []int{1}, 1)) // Expected: 10
}
```

## 1775 — Equal Sum Arrays With Minimum Number Of Operations

```go
package main

// LeetCode #1775: Equal Sum Arrays With Minimum Number of Operations
// https://leetcode.com/problems/equal-sum-arrays-with-minimum-number-of-operations/
// Difficulty: Medium
// Time: O(n + m), Space: O(1)

import "fmt"

func minOperations(nums1 []int, nums2 []int) int {
	sum1, sum2 := 0, 0
	for _, v := range nums1 {
		sum1 += v
	}
	for _, v := range nums2 {
		sum2 += v
	}

	if sum1 == sum2 {
		return 0
	}

	// Make nums1 the one with smaller sum
	if sum1 > sum2 {
		nums1, nums2 = nums2, nums1
		sum1, sum2 = sum2, sum1
	}

	diff := sum2 - sum1
	count := make([]int, 7) // possible increments/decrements

	// nums1: smaller sum, we want to increase values (change to max 6)
	for _, v := range nums1 {
		count[6-v]++ // max possible increase
	}
	// nums2: larger sum, we want to decrease values (change to min 1)
	for _, v := range nums2 {
		count[v-1]++ // max possible decrease
	}

	ops := 0
	for i := 6; i >= 1; i-- {
		for count[i] > 0 && diff > 0 {
			diff -= i
			count[i]--
			ops++
		}
	}

	if diff > 0 {
		return -1
	}
	return ops
}

func main() {
	fmt.Println(minOperations([]int{1, 2, 3, 4, 5, 6}, []int{1, 1, 2, 2, 2, 2})) // Expected: 3
	fmt.Println(minOperations([]int{1, 1, 1, 1}, []int{6, 6, 6, 6})) // Expected: 4
	fmt.Println(minOperations([]int{6, 6}, []int{1})) // Expected: 3
}
```

## 1778 — Shortest Path In A Hidden Grid

```go
package main

// LeetCode #1778: Shortest Path in a Hidden Grid
// https://leetcode.com/problems/shortest-path-in-a-hidden-grid/
// Difficulty: Medium [Paid]
// This is an interactive problem. We implement the solving algorithm.
// Time: O(m * n), Space: O(m * n)

import "fmt"

// GridMaster is the interface provided by LeetCode for the hidden grid problem.
// Real implementation: GridMaster.canMove(dir), GridMaster.move(dir), GridMaster.isTarget()
type GridMaster interface {
	CanMove(dir byte) bool
	Move(dir byte) bool // returns true if moved successfully
	IsTarget() bool
}

func findShortestPath(master GridMaster) int {
	// Discover the grid via DFS
	grid := make(map[[2]int]int) // 0=unvisited, 1=empty, 2=target, -1=blocked
	targetPos := [2]int{-1, -1}

	dirs := []byte{'U', 'D', 'L', 'R'}
	dr := map[byte]int{'U': -1, 'D': 1, 'L': 0, 'R': 0}
	dc := map[byte]int{'U': 0, 'D': 0, 'L': -1, 'R': 1}
	rev := map[byte]byte{'U': 'D', 'D': 'U', 'L': 'R', 'R': 'L'}

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if master.IsTarget() {
			targetPos = [2]int{r, c}
			grid[[2]int{r, c}] = 2
		}
		for _, d := range dirs {
			nr, nc := r+dr[d], c+dc[d]
			key := [2]int{nr, nc}
			if _, visited := grid[key]; !visited && master.CanMove(d) {
				master.Move(d)
				grid[key] = 1
				dfs(nr, nc)
				// Move back
				master.Move(rev[d])
			} else if !master.CanMove(d) {
				grid[key] = -1 // blocked (wall)
			}
		}
	}

	grid[[2]int{0, 0}] = 1
	dfs(0, 0)

	if targetPos == [2]int{-1, -1} {
		return -1
	}

	// BFS for shortest path
	queue := [][2]int{{0, 0}}
	visited := make(map[[2]int]bool)
	visited[[2]int{0, 0}] = true
	steps := 0

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			r, c := queue[i][0], queue[i][1]
			if r == targetPos[0] && c == targetPos[1] {
				return steps
			}
			for _, d := range dirs {
				nr, nc := r+dr[d], c+dc[d]
				key := [2]int{nr, nc}
				if v, ok := grid[key]; ok && v != -1 && !visited[key] {
					visited[key] = true
					queue = append(queue, [2]int{nr, nc})
				}
			}
		}
		queue = queue[size:]
		steps++
	}
	return -1
}

func main() {
	fmt.Println("Interactive problem - test via LeetCode platform")
	fmt.Println("Implementation ready for GridMaster interface")
}
```

## 1780 — Check If Number Is A Sum Of Powers Of Three

```go
package main

// LeetCode #1780: Check if Number is a Sum of Powers of Three
// https://leetcode.com/problems/check-if-number-is-a-sum-of-powers-of-three/
// Difficulty: Medium
// Time: O(log_3 n), Space: O(1)

import "fmt"

func checkPowersOfThree(n int) bool {
	for n > 0 {
		if n%3 == 2 {
			return false
		}
		n /= 3
	}
	return true
}

func main() {
	fmt.Println(checkPowersOfThree(12)) // Expected: true (3^2 + 3^1 = 9 + 3)
	fmt.Println(checkPowersOfThree(91)) // Expected: true (3^4 + 3^2 + 3^0 = 81 + 9 + 1)
	fmt.Println(checkPowersOfThree(21)) // Expected: false (21 = 2*9 + 3, cannot use 2)
}
```

## 1781 — Sum Of Beauty Of All Substrings

```go
package main

// LeetCode #1781: Sum of Beauty of All Substrings
// https://leetcode.com/problems/sum-of-beauty-of-all-substrings/
// Difficulty: Medium
// Time: O(n^2), Space: O(26)

import "fmt"

func beautySum(s string) int {
	n := len(s)
	result := 0

	for i := 0; i < n; i++ {
		count := make([]int, 26)
		for j := i; j < n; j++ {
			count[s[j]-'a']++
			minFreq, maxFreq := n, 0
			for _, f := range count {
				if f > 0 {
					if f < minFreq {
						minFreq = f
					}
					if f > maxFreq {
						maxFreq = f
					}
				}
			}
			result += maxFreq - minFreq
		}
	}
	return result
}

func main() {
	fmt.Println(beautySum("aabcb")) // Expected: 5
	fmt.Println(beautySum("aabcbaa")) // Expected: 17
	fmt.Println(beautySum("x")) // Expected: 0
}
```

## 1783 — Grand Slam Titles

```go
package main

// LeetCode #1783: Grand Slam Titles
// https://leetcode.com/problems/grand-slam-titles/
// Difficulty: Medium [Paid] (SQL)
// This is a SQL problem. In Go we simulate the logic.
// Time: O(n), Space: O(n)

import "fmt"

type Player struct {
	ID   int
	Name string
}

type Championship struct {
	Year  int
	WonBy int // player ID
	Tournament string
}

func countTitles(players []Player, championships []Championship) map[string]int {
	titles := make(map[string]int)
	for _, c := range championships {
		for _, p := range players {
			if p.ID == c.WonBy {
				titles[p.Name]++
				break
			}
		}
	}
	return titles
}

func main() {
	players := []Player{
		{1, "Federer"},
		{2, "Nadal"},
		{3, "Djokovic"},
	}
	champs := []Championship{
		{2023, 3, "Wimbledon"},
		{2023, 3, "US Open"},
		{2023, 2, "French Open"},
	}
	titles := countTitles(players, champs)
	for name, count := range titles {
		fmt.Printf("%s: %d\n", name, count)
	}
}
```

## 1785 — Minimum Elements To Add To Form A Given Sum

```go
package main

// LeetCode #1785: Minimum Elements to Add to Form a Given Sum
// https://leetcode.com/problems/minimum-elements-to-add-to-form-a-given-sum/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

func minElements(nums []int, limit int, goal int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	diff := goal - sum
	if diff < 0 {
		diff = -diff
	}

	// Minimum elements = ceil(diff / limit)
	return (diff + limit - 1) / limit
}

func main() {
	fmt.Println(minElements([]int{1, -1, 1}, 3, -4)) // Expected: 2
	fmt.Println(minElements([]int{1, -10, 9, 1}, 100, 0)) // Expected: 1
	fmt.Println(minElements([]int{0}, 1, 1000000)) // Expected: 1000000
}
```

## 1786 — Number Of Restricted Paths From First To Last Node

```go
package main

// LeetCode #1786: Number of Restricted Paths From First to Last Node
// https://leetcode.com/problems/number-of-restricted-paths-from-first-to-last-node/
// Difficulty: Medium
// Time: O(E log V), Space: O(V + E)

import (
	"container/heap"
	"fmt"
)

const mod = 1_000_000_007

type Edge struct {
	to, weight int
}

type Item struct {
	node, dist int
}

type PriorityQueue []Item

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i].dist < pq[j].dist }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(Item)) }
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[:n-1]
	return x
}

func countRestrictedPaths(n int, edges [][]int) int {
	graph := make([][]Edge, n+1)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		graph[u] = append(graph[u], Edge{v, w})
		graph[v] = append(graph[v], Edge{u, w})
	}

	// Dijkstra from node n to all nodes
	dist := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dist[i] = 1 << 60
	}
	dist[n] = 0

	pq := &PriorityQueue{}
	heap.Push(pq, Item{n, 0})

	for pq.Len() > 0 {
		item := heap.Pop(pq).(Item)
		u := item.node
		if item.dist > dist[u] {
			continue
		}
		for _, e := range graph[u] {
			if nd := dist[u] + e.weight; nd < dist[e.to] {
				dist[e.to] = nd
				heap.Push(pq, Item{e.to, nd})
			}
		}
	}

	// DP: count restricted paths
	memo := make([]int, n+1)
	for i := 1; i <= n; i++ {
		memo[i] = -1
	}

	var dfs func(u int) int
	dfs = func(u int) int {
		if u == n {
			return 1
		}
		if memo[u] != -1 {
			return memo[u]
		}
		total := 0
		for _, e := range graph[u] {
			if dist[e.to] < dist[u] {
				total = (total + dfs(e.to)) % mod
			}
		}
		memo[u] = total
		return total
	}

	return dfs(1)
}

func main() {
	fmt.Println(countRestrictedPaths(5, [][]int{{1, 2, 3}, {1, 3, 3}, {2, 3, 1}, {1, 4, 2}, {5, 2, 2}, {3, 5, 1}, {5, 4, 10}}))
	// Expected: 3

	fmt.Println(countRestrictedPaths(7, [][]int{{1, 3, 1}, {4, 1, 2}, {7, 3, 4}, {2, 5, 3}, {5, 6, 1}, {6, 7, 2}, {7, 5, 3}, {2, 6, 4}}))
	// Expected: 1

	fmt.Println(countRestrictedPaths(2, [][]int{{1, 2, 5}}))
	// Expected: 1
}
```

## 1792 — Maximum Average Pass Ratio

```go
package main

// LeetCode #1792: Maximum Average Pass Ratio
// https://leetcode.com/problems/maximum-average-pass-ratio/
// Difficulty: Medium
// Time: O((n+k) log n), Space: O(n)

import (
	"container/heap"
	"fmt"
)

type Class struct {
	pass, total int
	gain        float64
}

type MaxHeap []Class

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].gain > h[j].gain }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(Class)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	h := &MaxHeap{}
	heap.Init(h)

	for _, c := range classes {
		pass, total := c[0], c[1]
		gain := float64(pass+1)/float64(total+1) - float64(pass)/float64(total)
		heap.Push(h, Class{pass, total, gain})
	}

	for i := 0; i < extraStudents; i++ {
		c := heap.Pop(h).(Class)
		c.pass++
		c.total++
		c.gain = float64(c.pass+1)/float64(c.total+1) - float64(c.pass)/float64(c.total)
		heap.Push(h, c)
	}

	sum := 0.0
	for h.Len() > 0 {
		c := heap.Pop(h).(Class)
		sum += float64(c.pass) / float64(c.total)
	}
	return sum / float64(len(classes))
}

func main() {
	fmt.Printf("%.5f\n", maxAverageRatio([][]int{{1, 2}, {3, 5}, {2, 2}}, 2)) // Expected: 0.78333
	fmt.Printf("%.5f\n", maxAverageRatio([][]int{{2, 4}, {3, 9}, {4, 5}, {2, 10}}, 4)) // Expected: 0.53485
	fmt.Printf("%.5f\n", maxAverageRatio([][]int{{1, 2}}, 1)) // Expected: 0.66667
}
```

## 1794 — Count Pairs Of Equal Substrings With Minimum Difference

```go
package main

// LeetCode #1794: Count Pairs of Equal Substrings With Minimum Difference
// https://leetcode.com/problems/count-pairs-of-equal-substrings-with-minimum-difference/
// Difficulty: Medium [Paid]
// Time: O(n), Space: O(1)

import "fmt"

func countQuadruples(firstString string, secondString string) int {
	firstPos := make([]int, 26)
	lastPos := make([]int, 26)
	for i := range firstPos {
		firstPos[i] = -1
		lastPos[i] = -1
	}

	for i, ch := range firstString {
		idx := ch - 'a'
		if firstPos[idx] == -1 {
			firstPos[idx] = i
		}
	}
	for i, ch := range secondString {
		idx := ch - 'a'
		lastPos[idx] = i
	}

	minDiff := 1 << 30
	count := 0

	for i := 0; i < 26; i++ {
		if firstPos[i] != -1 && lastPos[i] != -1 {
			diff := firstPos[i] - lastPos[i]
			if diff < minDiff {
				minDiff = diff
				count = 1
			} else if diff == minDiff {
				count++
			}
		}
	}
	return count
}

func main() {
	fmt.Println(countQuadruples("abcd", "bcd")) // test 1
	fmt.Println(countQuadruples("abc", "abc")) // test 2
	fmt.Println(countQuadruples("abb", "b")) // test 3
}
```

## 1797 — Design Authentication Manager

```go
package main

// LeetCode #1797: Design Authentication Manager
// https://leetcode.com/problems/design-authentication-manager/
// Difficulty: Medium
// Time: O(n) per operation, Space: O(n)

import (
	"fmt"
)

type AuthenticationManager struct {
	ttl     int
	tokens  map[string]int
}

func Constructor(timeToLive int) AuthenticationManager {
	return AuthenticationManager{
		ttl:    timeToLive,
		tokens: make(map[string]int),
	}
}

func (am *AuthenticationManager) Generate(tokenId string, currentTime int) {
	am.tokens[tokenId] = currentTime + am.ttl
}

func (am *AuthenticationManager) Renew(tokenId string, currentTime int) {
	if expiry, ok := am.tokens[tokenId]; ok && expiry > currentTime {
		am.tokens[tokenId] = currentTime + am.ttl
	}
}

func (am *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
	count := 0
	for _, expiry := range am.tokens {
		if expiry > currentTime {
			count++
		}
	}
	return count
}

func main() {
	am := Constructor(5)
	am.Generate("aaa", 1)
	fmt.Println(am.CountUnexpiredTokens(2)) // Expected: 1
	am.Renew("aaa", 3)
	fmt.Println(am.CountUnexpiredTokens(6)) // Expected: 0
	am.Generate("bbb", 7)
	am.Generate("ccc", 8)
	fmt.Println(am.CountUnexpiredTokens(10)) // Expected: 2
	am.Renew("bbb", 11)
	fmt.Println(am.CountUnexpiredTokens(12)) // Expected: 1
}
```

## 1798 — Maximum Number Of Consecutive Values You Can Make

```go
package main

// LeetCode #1798: Maximum Number of Consecutive Values You Can Make
// https://leetcode.com/problems/maximum-number-of-consecutive-values-you-can-make/
// Difficulty: Medium
// Time: O(n log n), Space: O(1)

import (
	"fmt"
	"sort"
)

func getMaximumConsecutive(coins []int) int {
	sort.Ints(coins)
	maxReach := 0
	for _, c := range coins {
		if c > maxReach+1 {
			break
		}
		maxReach += c
	}
	return maxReach + 1
}

func main() {
	fmt.Println(getMaximumConsecutive([]int{1, 3}))          // Expected: 2
	fmt.Println(getMaximumConsecutive([]int{1, 1, 1, 4}))   // Expected: 8
	fmt.Println(getMaximumConsecutive([]int{1, 4, 10, 3, 1})) // Expected: 20
}
```

## 1801 — Number Of Orders In The Backlog

```go
package main

// LeetCode #1801: Number of Orders in the Backlog
// https://leetcode.com/problems/number-of-orders-in-the-backlog/
// Difficulty: Medium
// Time: O(n log n), Space: O(n)

import (
	"container/heap"
	"fmt"
)

type Order struct {
	price  int
	amount int
}

type MaxHeap []Order

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].price > h[j].price }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(Order)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type MinHeap []Order

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].price < h[j].price }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(Order)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

const mod = 1_000_000_007

func getNumberOfBacklogOrders(orders [][]int) int {
	buys := &MaxHeap{}
	sells := &MinHeap{}
	heap.Init(buys)
	heap.Init(sells)

	for _, o := range orders {
		price, amount, orderType := o[0], o[1], o[2]
		if orderType == 0 { // buy
			for amount > 0 && sells.Len() > 0 && (*sells)[0].price <= price {
				top := &(*sells)[0]
				if top.amount > amount {
					top.amount -= amount
					amount = 0
				} else {
					amount -= top.amount
					heap.Pop(sells)
				}
			}
			if amount > 0 {
				heap.Push(buys, Order{price, amount})
			}
		} else { // sell
			for amount > 0 && buys.Len() > 0 && (*buys)[0].price >= price {
				top := &(*buys)[0]
				if top.amount > amount {
					top.amount -= amount
					amount = 0
				} else {
					amount -= top.amount
					heap.Pop(buys)
				}
			}
			if amount > 0 {
				heap.Push(sells, Order{price, amount})
			}
		}
	}

	total := 0
	for buys.Len() > 0 {
		total = (total + heap.Pop(buys).(Order).amount) % mod
	}
	for sells.Len() > 0 {
		total = (total + heap.Pop(sells).(Order).amount) % mod
	}
	return total
}

func main() {
	fmt.Println(getNumberOfBacklogOrders([][]int{{10, 5, 0}, {15, 2, 1}, {25, 1, 1}, {30, 4, 0}})) // Expected: 6
	fmt.Println(getNumberOfBacklogOrders([][]int{{7, 1000000000, 1}, {15, 3, 0}, {5, 999999995, 0}, {5, 1, 1}})) // Expected: 999999984
}
```

## 1802 — Maximum Value At A Given Index In A Bounded Array

```go
package main

// LeetCode #1802: Maximum Value at a Given Index in a Bounded Array
// https://leetcode.com/problems/maximum-value-at-a-given-index-in-a-bounded-array/
// Difficulty: Medium
// Time: O(log maxVal), Space: O(1)

import "fmt"

func maxValue(n int, index int, maxSum int) int {
	left, right := 1, maxSum

	for left < right {
		mid := left + (right-left+1)/2
		if canMake(n, index, maxSum, mid) {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return left
}

func canMake(n, index, maxSum, val int) bool {
	// Sum of left side (decreasing from val to some minimum)
	leftLen := index
	rightLen := n - index - 1

	total := val // the peak
	total += sumTriangle(val-1, leftLen)
	total += sumTriangle(val-1, rightLen)

	return total <= maxSum
}

// sum of sequence starting from max down to some minimum, limited by count
func sumTriangle(max, count int) int {
	if count <= 0 {
		return 0
	}
	if max >= count {
		// Enough height: use arithmetic series
		// max, max-1, ..., max-count+1
		return (max + max - count + 1) * count / 2
	}
	// Not enough height: max + (max-1) + ... + 1 + 1 + ... + 1 (count - max times)
	return max*(max+1)/2 + (count - max)
}

func main() {
	fmt.Println(maxValue(4, 2, 6))  // Expected: 2
	fmt.Println(maxValue(6, 1, 10)) // Expected: 3
	fmt.Println(maxValue(8, 7, 14)) // Expected: 4
}
```

## 1804 — Implement Trie Ii Prefix Tree

```go
package main

// LeetCode #1804: Implement Trie II (Prefix Tree)
// https://leetcode.com/problems/implement-trie-ii-prefix-tree/
// Difficulty: Medium [Paid]
// Time: O(L) per operation, Space: O(total characters)

import "fmt"

type TrieNode struct {
	children [26]*TrieNode
	wordCnt  int
	prefixCnt int
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{&TrieNode{}}
}

func (t *Trie) Insert(word string) {
	node := t.root
	for _, ch := range word {
		idx := ch - 'a'
		if node.children[idx] == nil {
			node.children[idx] = &TrieNode{}
		}
		node = node.children[idx]
		node.prefixCnt++
	}
	node.wordCnt++
}

func (t *Trie) CountWordsEqualTo(word string) int {
	node := t.root
	for _, ch := range word {
		idx := ch - 'a'
		if node.children[idx] == nil {
			return 0
		}
		node = node.children[idx]
	}
	return node.wordCnt
}

func (t *Trie) CountWordsStartingWith(prefix string) int {
	node := t.root
	for _, ch := range prefix {
		idx := ch - 'a'
		if node.children[idx] == nil {
			return 0
		}
		node = node.children[idx]
	}
	return node.prefixCnt
}

func (t *Trie) Erase(word string) {
	node := t.root
	for _, ch := range word {
		idx := ch - 'a'
		node = node.children[idx]
		node.prefixCnt--
	}
	node.wordCnt--
}

func main() {
	trie := Constructor()
	trie.Insert("apple")
	trie.Insert("apple")
	fmt.Println(trie.CountWordsEqualTo("apple"))  // Expected: 2
	fmt.Println(trie.CountWordsStartingWith("app")) // Expected: 2
	trie.Erase("apple")
	fmt.Println(trie.CountWordsEqualTo("apple"))  // Expected: 1
	fmt.Println(trie.CountWordsStartingWith("app")) // Expected: 1
	trie.Erase("apple")
	fmt.Println(trie.CountWordsEqualTo("apple"))  // Expected: 0
}
```

## 1806 — Minimum Number Of Operations To Reinitialize A Permutation

```go
package main

// LeetCode #1806: Minimum Number of Operations to Reinitialize a Permutation
// https://leetcode.com/problems/minimum-number-of-operations-to-reinitialize-a-permutation/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

func reinitializePermutation(n int) int {
	ops := 0
	i := 1

	for {
		ops++
		if i%2 == 0 {
			i /= 2
		} else {
			i = n/2 + (i-1)/2
		}
		if i == 1 {
			break
		}
	}
	return ops
}

func main() {
	fmt.Println(reinitializePermutation(2))  // Expected: 1
	fmt.Println(reinitializePermutation(4))  // Expected: 2
	fmt.Println(reinitializePermutation(6))  // Expected: 4
}
```

## 1807 — Evaluate The Bracket Pairs Of A String

```go
package main

// LeetCode #1807: Evaluate the Bracket Pairs of a String
// https://leetcode.com/problems/evaluate-the-bracket-pairs-of-a-string/
// Difficulty: Medium
// Time: O(n), Space: O(k) where k = number of knowledge pairs

import (
	"fmt"
	"strings"
)

func evaluate(s string, knowledge [][]string) string {
	dict := make(map[string]string)
	for _, kv := range knowledge {
		dict[kv[0]] = kv[1]
	}

	var result strings.Builder
	i := 0
	for i < len(s) {
		if s[i] == '(' {
			j := i + 1
			for s[j] != ')' {
				j++
			}
			key := s[i+1 : j]
			if val, ok := dict[key]; ok {
				result.WriteString(val)
			} else {
				result.WriteByte('?')
			}
			i = j + 1
		} else {
			result.WriteByte(s[i])
			i++
		}
	}
	return result.String()
}

func main() {
	fmt.Println(evaluate("(name)is(age)yearsold", [][]string{{"name", "bob"}, {"age", "two"}})) // Expected: "bobistwoyearsold"
	fmt.Println(evaluate("hi(name)", [][]string{{"a", "b"}})) // Expected: "hi?"
	fmt.Println(evaluate("(a)(a)(a)aaa", [][]string{{"a", "yes"}})) // Expected: "yesyesyesaaa"
}
```

## 1810 — Minimum Path Cost In A Hidden Grid

```go
package main

// LeetCode #1810: Minimum Path Cost in a Hidden Grid
// https://leetcode.com/problems/minimum-path-cost-in-a-hidden-grid/
// Difficulty: Medium [Paid]
// This is an interactive problem. We implement the solving algorithm.
// Time: O(m * n log(m*n)), Space: O(m * n)

import "fmt"

// GridMaster is the interface provided by LeetCode for the hidden grid problem.
// Real implementation: GridMaster.canMove(dir), GridMaster.move(dir), GridMaster.isTarget()
// GridMaster.move(dir) returns int cost to move
type GridMaster interface {
	CanMove(dir byte) bool
	Move(dir byte) int
	IsTarget() bool
}

func findShortestPath(master GridMaster) int {
	// Discover grid: visited[r][c] = true means the cell was reached
	visited := make(map[[2]int]bool)
	costs := make(map[[2]int]int)
	targetPos := [2]int{-1, -1}

	dirs := []byte{'U', 'D', 'L', 'R'}
	dr := map[byte]int{'U': -1, 'D': 1, 'L': 0, 'R': 0}
	dc := map[byte]int{'U': 0, 'D': 0, 'L': -1, 'R': 1}
	rev := map[byte]byte{'U': 'D', 'D': 'U', 'L': 'R', 'R': 'L'}

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if master.IsTarget() {
			targetPos = [2]int{r, c}
		}
		visited[[2]int{r, c}] = true
		for _, d := range dirs {
			nr, nc := r+dr[d], c+dc[d]
			key := [2]int{nr, nc}
			if visited[key] {
				continue
			}
			if master.CanMove(d) {
				cost := master.Move(d)
				costs[key] = cost
				dfs(nr, nc)
				master.Move(rev[d])
			}
		}
	}

	dfs(0, 0)

	if targetPos == [2]int{-1, -1} {
		return -1
	}

	// Dijkstra for shortest weighted path
	dist := make(map[[2]int]int)
	dist[[2]int{0, 0}] = 0
	pq := [][3]int{{0, 0, 0}} // dist, r, c

	for len(pq) > 0 {
		// Extract min
		minIdx := 0
		for i := 1; i < len(pq); i++ {
			if pq[i][0] < pq[minIdx][0] {
				minIdx = i
			}
		}
		cur := pq[minIdx]
		pq = append(pq[:minIdx], pq[minIdx+1:]...)
		d, r, c := cur[0], cur[1], cur[2]

		if [2]int{r, c} == targetPos {
			return d
		}

		if d > dist[[2]int{r, c}] {
			continue
		}

		for _, dir := range dirs {
			nr, nc := r+dr[dir], c+dc[dir]
			nk := [2]int{nr, nc}
			if !visited[nk] {
				continue
			}
			nd := d + costs[nk]
			if old, ok := dist[nk]; !ok || nd < old {
				dist[nk] = nd
				pq = append(pq, [3]int{nd, nr, nc})
			}
		}
	}
	return -1
}

func main() {
	fmt.Println("Interactive problem - test via LeetCode platform")
	fmt.Println("Implementation ready for GridMaster interface with costs")
}
```

## 1811 — Find Interview Candidates

```go
package main

// LeetCode #1811: Find Interview Candidates
// https://leetcode.com/problems/find-interview-candidates/
// Difficulty: Medium [Paid] (SQL)
// This is a SQL problem. In Go we simulate the logic.
// Time: O(n), Space: O(n)

import "fmt"

type ContestScore struct {
	ContestID int
	GoldID    int // user ID of gold medalist
	SilverID  int
	BronzeID  int
}

type UserContest struct {
	UserID    int
	ContestID int
	Score     int
}

func findCandidates(contests []ContestScore, submissions []UserContest) []int {
	medalCount := make(map[int]int)
	for _, c := range contests {
		medalCount[c.GoldID]++
		medalCount[c.SilverID]++
		medalCount[c.BronzeID]++
	}

	candidateSet := make(map[int]bool)

	// Any user with 3+ medals
	for user, count := range medalCount {
		if count >= 3 {
			candidateSet[user] = true
		}
	}

	// Users who won gold in consecutive contests (contests can be consecutive by contest_id)
	userContests := make(map[int][]int)
	for _, c := range contests {
		userContests[c.GoldID] = append(userContests[c.GoldID], c.ContestID)
	}

	for user, ids := range userContests {
		// Sort contest IDs
		for i := 1; i < len(ids); i++ {
			if ids[i] == ids[i-1]+1 {
				candidateSet[user] = true
				break
			}
		}
	}

	result := make([]int, 0, len(candidateSet))
	for u := range candidateSet {
		result = append(result, u)
	}
	return result
}

func main() {
	contests := []ContestScore{
		{1, 1, 2, 3},
		{2, 1, 4, 5},
		{3, 1, 6, 7},
	}
	fmt.Println(findCandidates(contests, nil)) // Expected: [1] (gold in 3 consecutive contests)
}
```

## 1813 — Sentence Similarity Iii

```go
package main

// LeetCode #1813: Sentence Similarity III
// https://leetcode.com/problems/sentence-similarity-iii/
// Difficulty: Medium
// Time: O(n + m), Space: O(n + m)

import (
	"fmt"
	"strings"
)

func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	words1 := strings.Split(sentence1, " ")
	words2 := strings.Split(sentence2, " ")

	// Make words1 the shorter one
	if len(words1) > len(words2) {
		words1, words2 = words2, words1
	}

	i, j := 0, len(words1)-1
	k, l := 0, len(words2)-1

	// Match from beginning
	for i < len(words1) && k <= l && words1[i] == words2[k] {
		i++
		k++
	}

	// Match from end
	for j >= i && l >= k && words1[j] == words2[l] {
		j--
		l--
	}

	return i > j // All words in shorter sentence matched
}

func main() {
	fmt.Println(areSentencesSimilar("My name is Haley", "My Haley")) // Expected: true
	fmt.Println(areSentencesSimilar("of", "A lot of words")) // Expected: false
	fmt.Println(areSentencesSimilar("Eating right now", "Eating")) // Expected: true
}
```

## 1814 — Count Nice Pairs In An Array

```go
package main

// LeetCode #1814: Count Nice Pairs in an Array
// https://leetcode.com/problems/count-nice-pairs-in-an-array/
// Difficulty: Medium
// Time: O(n log M) where M = max digit length, Space: O(n)

import "fmt"

const mod = 1_000_000_007

func countNicePairs(nums []int) int {
	count := make(map[int]int)
	result := 0

	for _, v := range nums {
		key := v - rev(v)
		result = (result + count[key]) % mod
		count[key]++
	}
	return result
}

func rev(x int) int {
	r := 0
	for x > 0 {
		r = r*10 + x%10
		x /= 10
	}
	return r
}

func main() {
	fmt.Println(countNicePairs([]int{42, 11, 1, 97})) // Expected: 2
	fmt.Println(countNicePairs([]int{13, 10, 35, 24, 76})) // Expected: 4
	fmt.Println(countNicePairs([]int{1, 1, 1, 1})) // Expected: 6
}
```

## 1817 — Finding The Users Active Minutes

```go
package main

// LeetCode #1817: Finding the Users Active Minutes
// https://leetcode.com/problems/finding-the-users-active-minutes/
// Difficulty: Medium
// Time: O(n), Space: O(n)

import "fmt"

func findingUsersActiveMinutes(logs [][]int, k int) []int {
	userMinutes := make(map[int]map[int]bool)
	for _, log := range logs {
		id, min := log[0], log[1]
		if userMinutes[id] == nil {
			userMinutes[id] = make(map[int]bool)
		}
		userMinutes[id][min] = true
	}

	result := make([]int, k)
	for _, minutes := range userMinutes {
		uam := len(minutes)
		if uam <= k {
			result[uam-1]++
		}
	}
	return result
}

func main() {
	fmt.Println(findingUsersActiveMinutes([][]int{{0, 5}, {1, 2}, {0, 2}, {0, 5}, {1, 3}}, 5)) // Expected: [0, 2, 0, 0, 0]
	fmt.Println(findingUsersActiveMinutes([][]int{{1, 1}, {2, 2}, {2, 3}}, 4)) // Expected: [1, 1, 0, 0]
}
```

## 1818 — Minimum Absolute Sum Difference

```go
package main

// LeetCode #1818: Minimum Absolute Sum Difference
// https://leetcode.com/problems/minimum-absolute-sum-difference/
// Difficulty: Medium
// Time: O(n log n), Space: O(n)

import (
	"fmt"
	"sort"
)

const mod = 1_000_000_007

func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	n := len(nums1)
	sorted := make([]int, n)
	copy(sorted, nums1)
	sort.Ints(sorted)

	total := 0
	maxReduction := 0

	for i := 0; i < n; i++ {
		origDiff := abs(nums1[i] - nums2[i])
		total = (total + origDiff) % mod

		// Find closest value to nums2[i] in sorted nums1
		idx := sort.SearchInts(sorted, nums2[i])
		if idx < n {
			maxReduction = max(maxReduction, origDiff-abs(sorted[idx]-nums2[i]))
		}
		if idx > 0 {
			maxReduction = max(maxReduction, origDiff-abs(sorted[idx-1]-nums2[i]))
		}
	}

	return (total - maxReduction + mod) % mod
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minAbsoluteSumDiff([]int{1, 7, 5}, []int{2, 3, 5})) // Expected: 3
	fmt.Println(minAbsoluteSumDiff([]int{2, 4, 6, 8, 10}, []int{2, 4, 6, 8, 10})) // Expected: 0
	fmt.Println(minAbsoluteSumDiff([]int{1, 10, 4, 4, 2, 7}, []int{9, 3, 5, 1, 7, 4})) // Expected: 20
}
```

## 1820 — Maximum Number Of Accepted Invitations

```go
package main

// LeetCode #1820: Maximum Number of Accepted Invitations
// https://leetcode.com/problems/maximum-number-of-accepted-invitations/
// Difficulty: Medium [Paid]
// Time: O(m * n^2) using DFS for bipartite matching

import "fmt"

func maximumInvitations(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	match := make([]int, n)
	for i := range match {
		match[i] = -1
	}

	var dfs func(u int, seen []bool) bool
	dfs = func(u int, seen []bool) bool {
		for v := 0; v < n; v++ {
			if grid[u][v] == 1 && !seen[v] {
				seen[v] = true
				if match[v] == -1 || dfs(match[v], seen) {
					match[v] = u
					return true
				}
			}
		}
		return false
	}

	result := 0
	for u := 0; u < m; u++ {
		seen := make([]bool, n)
		if dfs(u, seen) {
			result++
		}
	}
	return result
}

func main() {
	fmt.Println(maximumInvitations([][]int{{1, 1, 1}, {1, 0, 1}, {0, 0, 1}})) // Expected: 3
	fmt.Println(maximumInvitations([][]int{{1, 0, 1, 0}, {1, 0, 0, 0}, {0, 0, 1, 0}, {1, 1, 1, 0}})) // Expected: 3
}
```

## 1823 — Find The Winner Of The Circular Game

```go
package main

// LeetCode #1823: Find the Winner of the Circular Game
// https://leetcode.com/problems/find-the-winner-of-the-circular-game/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

func findTheWinner(n int, k int) int {
	winner := 0 // 0-indexed position for 1 person
	for i := 2; i <= n; i++ {
		winner = (winner + k) % i
	}
	return winner + 1 // convert to 1-indexed
}

func main() {
	fmt.Println(findTheWinner(5, 2)) // Expected: 3
	fmt.Println(findTheWinner(6, 5)) // Expected: 1
	fmt.Println(findTheWinner(1, 1)) // Expected: 1
}
```

## 1824 — Minimum Sideway Jumps

```go
package main

// LeetCode #1824: Minimum Sideway Jumps
// https://leetcode.com/problems/minimum-sideway-jumps/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

func minSideJumps(obstacles []int) int {
	// dp[i] = min jumps to reach lane i (0-indexed: 0, 1, 2 for lanes 1, 2, 3)
	dp := []int{1, 0, 1} // start at lane 2 (index 1)

	for _, obs := range obstacles {
		if obs > 0 {
			lane := obs - 1
			dp[lane] = 1 << 30 // blocked
		}
		// Try jumping sideways from other lanes
		for i := 0; i < 3; i++ {
			if i != obs-1 {
				for j := 0; j < 3; j++ {
					if j != i && j != obs-1 {
						dp[i] = min(dp[i], dp[j]+1)
					}
				}
			}
		}
	}
	return min(dp[0], min(dp[1], dp[2]))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minSideJumps([]int{0, 1, 2, 3, 0})) // Expected: 2
	fmt.Println(minSideJumps([]int{0, 1, 1, 3, 3, 0})) // Expected: 0
	fmt.Println(minSideJumps([]int{0, 2, 1, 0, 3, 0})) // Expected: 2
}
```

## 1828 — Queries On Number Of Points Inside A Circle

```go
package main

// LeetCode #1828: Queries on Number of Points Inside a Circle
// https://leetcode.com/problems/queries-on-number-of-points-inside-a-circle/
// Difficulty: Medium
// Time: O(n * q), Space: O(q)

import "fmt"

func countPoints(points [][]int, queries [][]int) []int {
	result := make([]int, len(queries))

	for i, q := range queries {
		cx, cy, r := q[0], q[1], q[2]
		r2 := r * r
		count := 0
		for _, p := range points {
			dx := p[0] - cx
			dy := p[1] - cy
			if dx*dx+dy*dy <= r2 {
				count++
			}
		}
		result[i] = count
	}
	return result
}

func main() {
	fmt.Println(countPoints([][]int{{1, 3}, {3, 3}, {5, 3}, {2, 2}}, [][]int{{2, 3, 1}, {4, 3, 1}, {1, 1, 2}}))
	// Expected: [3, 2, 2]

	fmt.Println(countPoints([][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}}, [][]int{{1, 2, 2}, {2, 2, 2}, {4, 3, 2}, {4, 3, 3}}))
	// Expected: [2, 3, 2, 3]
}
```

## 1829 — Maximum Xor For Each Query

```go
package main

// LeetCode #1829: Maximum XOR for Each Query
// https://leetcode.com/problems/maximum-xor-for-each-query/
// Difficulty: Medium
// Time: O(n), Space: O(1) excluding output

import "fmt"

func getMaximumXor(nums []int, maximumBit int) []int {
	n := len(nums)
	xor := 0
	for _, v := range nums {
		xor ^= v
	}

	maxVal := (1 << maximumBit) - 1
	result := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		// Best k is the one that maximizes xor ^ k, i.e., xor ^ maxVal
		result[n-1-i] = xor ^ maxVal
		// Remove last element for next query
		xor ^= nums[i]
	}
	return result
}

func main() {
	fmt.Println(getMaximumXor([]int{0, 1, 1, 3}, 2)) // Expected: [0, 3, 2, 3]
	fmt.Println(getMaximumXor([]int{2, 3, 4, 7}, 3)) // Expected: [5, 2, 6, 5]
	fmt.Println(getMaximumXor([]int{0, 1, 2, 2, 5, 7}, 3)) // Expected: [4, 3, 6, 4, 6, 7]
}
```

## 1831 — Maximum Transaction Each Day

```go
package main

// LeetCode #1831: Maximum Transaction Each Day
// https://leetcode.com/problems/maximum-transaction-each-day/
// Difficulty: Medium [Paid] (SQL)
// This is a SQL problem. In Go we simulate the logic.
// Time: O(n log n), Space: O(n)

import (
	"fmt"
	"sort"
)

type Transaction struct {
	ID     int
	Day    int
	Amount int
}

func maxTransactionPerDay(transactions []Transaction) []int {
	// Group by day, find max amount
	dayMax := make(map[int]int)
	for _, t := range transactions {
		if t.Amount > dayMax[t.Day] {
			dayMax[t.Day] = t.Amount
		}
	}

	// Find transaction IDs that have max amount for their day
	result := make([]int, 0)
	for _, t := range transactions {
		if t.Amount == dayMax[t.Day] {
			result = append(result, t.ID)
		}
	}
	sort.Ints(result)
	return result
}

func main() {
	transactions := []Transaction{
		{1, 1, 100},
		{2, 1, 200},
		{3, 2, 150},
		{4, 2, 100},
	}
	fmt.Println(maxTransactionPerDay(transactions)) // Expected: [2, 3]
}
```

## 1833 — Maximum Ice Cream Bars

```go
package main

// LeetCode #1833: Maximum Ice Cream Bars
// https://leetcode.com/problems/maximum-ice-cream-bars/
// Difficulty: Medium
// Time: O(n log n), Space: O(1)

import (
	"fmt"
	"sort"
)

func maxIceCream(costs []int, coins int) int {
	sort.Ints(costs)
	count := 0
	for _, c := range costs {
		if coins >= c {
			coins -= c
			count++
		} else {
			break
		}
	}
	return count
}

func main() {
	fmt.Println(maxIceCream([]int{1, 3, 2, 4, 1}, 7)) // Expected: 4
	fmt.Println(maxIceCream([]int{10, 6, 8, 7, 7, 8}, 5)) // Expected: 0
	fmt.Println(maxIceCream([]int{1, 6, 3, 1, 2, 5}, 20)) // Expected: 6
}
```

## 1834 — Single Threaded Cpu

```go
package main

// LeetCode #1834: Single-Threaded CPU
// https://leetcode.com/problems/single-threaded-cpu/
// Difficulty: Medium
// Time: O(n log n), Space: O(n)

import (
	"container/heap"
	"fmt"
	"sort"
)

type Task struct {
	index       int
	enqueueTime int
	processTime int
}

type MinHeap []Task

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool {
	if h[i].processTime != h[j].processTime {
		return h[i].processTime < h[j].processTime
	}
	return h[i].index < h[j].index
}
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(Task)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func getOrder(tasks [][]int) []int {
	n := len(tasks)
	taskList := make([]Task, n)
	for i, t := range tasks {
		taskList[i] = Task{index: i, enqueueTime: t[0], processTime: t[1]}
	}

	sort.Slice(taskList, func(i, j int) bool {
		return taskList[i].enqueueTime < taskList[j].enqueueTime
	})

	result := make([]int, 0, n)
	pq := &MinHeap{}
	heap.Init(pq)
	time := 0
	i := 0

	for i < n || pq.Len() > 0 {
		// Add all available tasks
		for i < n && taskList[i].enqueueTime <= time {
			heap.Push(pq, taskList[i])
			i++
		}
		if pq.Len() == 0 {
			time = taskList[i].enqueueTime
			continue
		}
		t := heap.Pop(pq).(Task)
		result = append(result, t.index)
		time += t.processTime
	}
	return result
}

func main() {
	fmt.Println(getOrder([][]int{{1, 2}, {2, 4}, {3, 2}, {4, 1}})) // Expected: [0, 2, 3, 1]
	fmt.Println(getOrder([][]int{{7, 10}, {7, 12}, {7, 5}, {7, 4}, {7, 2}})) // Expected: [4, 3, 2, 0, 1]
	fmt.Println(getOrder([][]int{{5, 2}, {7, 2}, {9, 4}, {6, 3}, {5, 10}, {1, 1}})) // Expected: [5, 0, 1, 3, 2, 4]
}
```

## 1836 — Remove Duplicates From An Unsorted Linked List

```go
package main

// LeetCode #1836: Remove Duplicates From an Unsorted Linked List
// https://leetcode.com/problems/remove-duplicates-from-an-unsorted-linked-list/
// Difficulty: Medium [Paid]
// Time: O(n), Space: O(n)

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicatesUnsorted(head *ListNode) *ListNode {
	// Count frequencies
	count := make(map[int]int)
	curr := head
	for curr != nil {
		count[curr.Val]++
		curr = curr.Next
	}

	// Remove duplicates
	dummy := &ListNode{Next: head}
	prev := dummy
	curr = head
	for curr != nil {
		if count[curr.Val] > 1 {
			prev.Next = curr.Next
		} else {
			prev = curr
		}
		curr = curr.Next
	}
	return dummy.Next
}

func makeList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	curr := head
	for i := 1; i < len(vals); i++ {
		curr.Next = &ListNode{Val: vals[i]}
		curr = curr.Next
	}
	return head
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func main() {
	l1 := makeList([]int{1, 2, 3, 2})
	printList(deleteDuplicatesUnsorted(l1)) // Expected: 1 3

	l2 := makeList([]int{2, 1, 1, 2})
	printList(deleteDuplicatesUnsorted(l2)) // Expected: (empty)

	l3 := makeList([]int{3, 2, 2, 1, 3, 2, 4})
	printList(deleteDuplicatesUnsorted(l3)) // Expected: 1 4
}
```

## 1838 — Frequency Of The Most Frequent Element

```go
package main

// LeetCode #1838: Frequency of the Most Frequent Element
// https://leetcode.com/problems/frequency-of-the-most-frequent-element/
// Difficulty: Medium
// Time: O(n log n), Space: O(1)

import (
	"fmt"
	"sort"
)

func maxFrequency(nums []int, k int) int {
	sort.Ints(nums)
	left := 0
	total := 0
	maxFreq := 0

	for right := 0; right < len(nums); right++ {
		total += nums[right]

		// Shrink window if we can't make all elements in window equal
		for nums[right]*(right-left+1)-total > k {
			total -= nums[left]
			left++
		}

		if right-left+1 > maxFreq {
			maxFreq = right - left + 1
		}
	}
	return maxFreq
}

func main() {
	fmt.Println(maxFrequency([]int{1, 2, 4}, 5))       // Expected: 3
	fmt.Println(maxFrequency([]int{1, 4, 8, 13}, 5))   // Expected: 2
	fmt.Println(maxFrequency([]int{3, 9, 6}, 2))       // Expected: 1
}
```

## 1839 — Longest Substring Of All Vowels In Order

```go
package main

// LeetCode #1839: Longest Substring Of All Vowels in Order
// https://leetcode.com/problems/longest-substring-of-all-vowels-in-order/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

func longestBeautifulSubstring(word string) int {
	vowels := "aeiou"
	maxLen := 0
	i := 0
	n := len(word)

	for i < n {
		// Start of a new substring
		vowelIdx := 0
		start := i

		// Check if starts with 'a'
		if word[i] != 'a' {
			i++
			continue
		}

		for i < n && vowelIdx < 5 {
			if word[i] == vowels[vowelIdx] {
				i++
			} else if vowelIdx+1 < 5 && word[i] == vowels[vowelIdx+1] {
				vowelIdx++
				i++
			} else {
				break
			}
		}

		if vowelIdx == 4 {
			length := i - start
			if length > maxLen {
				maxLen = length
			}
		}
	}
	return maxLen
}

func main() {
	fmt.Println(longestBeautifulSubstring("aeiaaioaaaaeiiiiouuuooaauuaeiu")) // Expected: 13
	fmt.Println(longestBeautifulSubstring("aeeeiiiioooauuuaeiou")) // Expected: 5
	fmt.Println(longestBeautifulSubstring("aaaa")) // Expected: 0 (no 'e')
}
```

## 1841 — League Statistics

```go
package main

// LeetCode #1841: League Statistics
// https://leetcode.com/problems/league-statistics/
// Difficulty: Medium [Paid] (SQL)
// This is a SQL problem. In Go we simulate the logic.
// Time: O(n log n), Space: O(n)

import (
	"fmt"
	"sort"
)

type Match struct {
	HomeTeam int
	AwayTeam int
	HomeGoals int
	AwayGoals int
}

type TeamStats struct {
	TeamID   int
	Played   int
	Won      int
	Drawn    int
	Lost     int
	GoalsFor int
	GoalsAgainst int
}

func (ts TeamStats) Points() int {
	return ts.Won*3 + ts.Drawn
}

func leagueStandings(matches []Match) []TeamStats {
	stats := make(map[int]*TeamStats)

	for _, m := range matches {
		if stats[m.HomeTeam] == nil {
			stats[m.HomeTeam] = &TeamStats{TeamID: m.HomeTeam}
		}
		if stats[m.AwayTeam] == nil {
			stats[m.AwayTeam] = &TeamStats{TeamID: m.AwayTeam}
		}

		home := stats[m.HomeTeam]
		away := stats[m.AwayTeam]
		home.Played++
		away.Played++
		home.GoalsFor += m.HomeGoals
		home.GoalsAgainst += m.AwayGoals
		away.GoalsFor += m.AwayGoals
		away.GoalsAgainst += m.HomeGoals

		if m.HomeGoals > m.AwayGoals {
			home.Won++
			away.Lost++
		} else if m.HomeGoals < m.AwayGoals {
			away.Won++
			home.Lost++
		} else {
			home.Drawn++
			away.Drawn++
		}
	}

	result := make([]TeamStats, 0, len(stats))
	for _, ts := range stats {
		result = append(result, *ts)
	}
	sort.Slice(result, func(i, j int) bool {
		pi, pj := result[i].Points(), result[j].Points()
		if pi != pj {
			return pi > pj
		}
		// Goal difference
		gdi := result[i].GoalsFor - result[i].GoalsAgainst
		gdj := result[j].GoalsFor - result[j].GoalsAgainst
		if gdi != gdj {
			return gdi > gdj
		}
		return result[i].GoalsFor > result[j].GoalsFor
	})
	return result
}

func main() {
	matches := []Match{
		{1, 2, 2, 1},
		{1, 3, 1, 1},
		{2, 3, 3, 0},
	}
	standings := leagueStandings(matches)
	for _, s := range standings {
		fmt.Printf("Team %d: %d pts (%dW %dD %dL, %d:%d)\n",
			s.TeamID, s.Points(), s.Won, s.Drawn, s.Lost, s.GoalsFor, s.GoalsAgainst)
	}
}
```

## 1843 — Suspicious Bank Accounts

```go
package main

// LeetCode #1843: Suspicious Bank Accounts
// https://leetcode.com/problems/suspicious-bank-accounts/
// Difficulty: Medium [Paid] (SQL)
// This is a SQL problem.
// In Go we provide a placeholder implementation.
// Time: O(n), Space: O(n)

import "fmt"

func findSuspicious() []int {
	return []int{}
}

func main() {
	fmt.Println("SQL problem - Suspicious Bank Accounts")
	fmt.Println("Implementation requires database queries with JOINs")
}
```

## 1845 — Seat Reservation Manager

```go
package main

// LeetCode #1845: Seat Reservation Manager
// https://leetcode.com/problems/seat-reservation-manager/
// Difficulty: Medium
// Time: O(log n) per operation, Space: O(n)

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type SeatManager struct {
	h        *MinHeap
	nextSeat int
	n        int
}

func Constructor(n int) SeatManager {
	return SeatManager{
		h:        &MinHeap{},
		nextSeat: 1,
		n:        n,
	}
}

func (sm *SeatManager) Reserve() int {
	if sm.h.Len() > 0 {
		return heap.Pop(sm.h).(int)
	}
	seat := sm.nextSeat
	sm.nextSeat++
	return seat
}

func (sm *SeatManager) Unreserve(seatNumber int) {
	heap.Push(sm.h, seatNumber)
}

func main() {
	sm := Constructor(5)
	fmt.Println(sm.Reserve())    // Expected: 1
	fmt.Println(sm.Reserve())    // Expected: 2
	sm.Unreserve(1)
	fmt.Println(sm.Reserve())    // Expected: 1
	fmt.Println(sm.Reserve())    // Expected: 3
	fmt.Println(sm.Reserve())    // Expected: 4
	sm.Unreserve(2)
	sm.Unreserve(3)
	fmt.Println(sm.Reserve())    // Expected: 2 (smallest unreserved)
}
```

## 1846 — Maximum Element After Decreasing And Rearranging

```go
package main

// LeetCode #1846: Maximum Element After Decreasing and Rearranging
// https://leetcode.com/problems/maximum-element-after-decreasing-and-rearranging/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MaximumElementAfterDecreasingAndRearranging([]int{2, 2, 1, 2, 1}))
	fmt.Println(MaximumElementAfterDecreasingAndRearranging([]int{100, 1, 1000}))
	fmt.Println(MaximumElementAfterDecreasingAndRearranging([]int{1, 2, 3, 4, 5}))
}

// Time: O(n log n), Space: O(1) (ignoring sort space)
func MaximumElementAfterDecreasingAndRearranging(arr []int) int {
	sort.Ints(arr)
	arr[0] = 1
	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] > 1 {
			arr[i] = arr[i-1] + 1
		}
	}
	return arr[len(arr)-1]
}
```

## 1849 — Splitting A String Into Descending Consecutive Values

```go
package main

// LeetCode #1849: Splitting a String Into Descending Consecutive Values
// https://leetcode.com/problems/splitting-a-string-into-descending-consecutive-values/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(SplitString("1234"))
	fmt.Println(SplitString("050043"))
	fmt.Println(SplitString("9080701"))
}

// Time: O(n^2), Space: O(n) for recursion
func SplitString(s string) bool {
	var dfs func(idx int, prev int64, count int) bool
	dfs = func(idx int, prev int64, count int) bool {
		if idx == len(s) {
			return count >= 2
		}
		num := int64(0)
		for i := idx; i < len(s); i++ {
			num = num*10 + int64(s[i]-'0')
			if num > 1<<62 {
				break
			}
			if count == 0 || prev-num == 1 {
				if dfs(i+1, num, count+1) {
					return true
				}
			}
			if num == 0 {
				break
			}
		}
		return false
	}
	return dfs(0, 0, 0)
}
```

## 1850 — Minimum Adjacent Swaps To Reach The Kth Smallest Number

```go
package main

// LeetCode #1850: Minimum Adjacent Swaps to Reach the Kth Smallest Number
// https://leetcode.com/problems/minimum-adjacent-swaps-to-reach-the-kth-smallest-number/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(GetMinSwaps("5489355142", 4))
	fmt.Println(GetMinSwaps("11112", 4))
	fmt.Println(GetMinSwaps("00123", 1))
}

// Time: O(n*k + n^2), Space: O(n)
func GetMinSwaps(num string, k int) int {
	nums := []byte(num)
	n := len(nums)

	// Generate k-th next permutation
	var nextPerm func()
	nextPerm = func() {
		// Find longest non-increasing suffix
		i := n - 2
		for i >= 0 && nums[i] >= nums[i+1] {
			i--
		}
		if i < 0 {
			return
		}
		// Find rightmost element > nums[i]
		j := n - 1
		for nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
		// Reverse suffix
		left, right := i+1, n-1
		for left < right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}

	for i := 0; i < k; i++ {
		nextPerm()
	}
	target := string(nums)

	// Count minimum adjacent swaps from original to target
	original := []byte(num)
	swaps := 0
	for i := 0; i < n; i++ {
		if original[i] != target[i] {
			j := i
			for j < n && original[j] != target[i] {
				j++
			}
			for j > i {
				original[j], original[j-1] = original[j-1], original[j]
				swaps++
				j--
			}
		}
	}
	return swaps
}
```

## 1852 — Distinct Numbers In Each Subarray

```go
package main

// LeetCode #1852: Distinct Numbers in Each Subarray
// https://leetcode.com/problems/distinct-numbers-in-each-subarray/
// Difficulty: Medium [Paid]

import "fmt"

func main() {
	fmt.Println(DistinctNumbers([]int{1, 2, 3, 2, 2, 1, 3}, 3))
	fmt.Println(DistinctNumbers([]int{1, 1, 1, 1, 1}, 2))
	fmt.Println(DistinctNumbers([]int{1, 2, 3, 4}, 1))
}

// Time: O(n), Space: O(k) where k = distinct elements in window
func DistinctNumbers(nums []int, k int) []int {
	n := len(nums)
	if k > n {
		return nil
	}
	result := make([]int, n-k+1)
	freq := make(map[int]int)
	distinct := 0

	for i := 0; i < n; i++ {
		// Add right element
		freq[nums[i]]++
		if freq[nums[i]] == 1 {
			distinct++
		}

		// Remove left element when window exceeds k
		if i >= k {
			freq[nums[i-k]]--
			if freq[nums[i-k]] == 0 {
				distinct--
			}
		}

		// Record result for completed windows
		if i >= k-1 {
			result[i-k+1] = distinct
		}
	}
	return result
}
```

## 1855 — Maximum Distance Between A Pair Of Values

```go
package main

// LeetCode #1855: Maximum Distance Between a Pair of Values
// https://leetcode.com/problems/maximum-distance-between-a-pair-of-values/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MaxDistance([]int{55, 30, 5, 4, 2}, []int{100, 20, 10, 10, 5}))
	fmt.Println(MaxDistance([]int{2, 2, 2}, []int{10, 10, 1}))
	fmt.Println(MaxDistance([]int{30, 29, 19, 5}, []int{25, 25, 25, 25, 25}))
}

// Time: O(m+n), Space: O(1)
func MaxDistance(nums1 []int, nums2 []int) int {
	i, j := 0, 0
	maxDist := 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] <= nums2[j] {
			if j-i > maxDist {
				maxDist = j - i
			}
			j++
		} else {
			i++
		}
	}
	return maxDist
}
```

## 1856 — Maximum Subarray Min Product

```go
package main

// LeetCode #1856: Maximum Subarray Min-Product
// https://leetcode.com/problems/maximum-subarray-min-product/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MaxSumMinProduct([]int{1, 2, 3, 2}))
	fmt.Println(MaxSumMinProduct([]int{2, 3, 3, 1, 2}))
	fmt.Println(MaxSumMinProduct([]int{3, 1, 5, 6, 4, 2}))
}

const mod = 1000000007

// Time: O(n), Space: O(n)
func MaxSumMinProduct(nums []int) int {
	n := len(nums)
	prefix := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + nums[i]
	}

	// Monotonic stack to find previous smaller and next smaller elements
	left := make([]int, n)  // left[i] = index of previous smaller element
	right := make([]int, n) // right[i] = index of next smaller element

	stack := make([]int, 0)
	for i := 0; i < n; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] >= nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	stack = make([]int, 0)
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && nums[stack[len(stack)-1]] >= nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			right[i] = n
		} else {
			right[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	maxProd := int64(0)
	for i := 0; i < n; i++ {
		sum := int64(prefix[right[i]] - prefix[left[i]+1])
		prod := sum * int64(nums[i])
		if prod > maxProd {
			maxProd = prod
		}
	}
	return int(maxProd % mod)
}
```

## 1858 — Longest Word With All Prefixes

```go
package main

// LeetCode #1858: Longest Word With All Prefixes
// https://leetcode.com/problems/longest-word-with-all-prefixes/
// Difficulty: Medium [Paid]

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(LongestWord([]string{"k", "ki", "kir", "kira", "kiran"}))
	fmt.Println(LongestWord([]string{"a", "banana", "app", "appl", "ap", "apply", "apple"}))
	fmt.Println(LongestWord([]string{"abc", "ab", "a"}))
}

// Time: O(n log n + total chars), Space: O(total unique prefixes)
func LongestWord(words []string) string {
	prefixSet := make(map[string]bool)
	for _, w := range words {
		prefixSet[w] = true
	}

	sort.Slice(words, func(i, j int) bool {
		if len(words[i]) != len(words[j]) {
			return len(words[i]) > len(words[j])
		}
		return words[i] < words[j]
	})

	for _, w := range words {
		valid := true
		for i := 1; i <= len(w); i++ {
			if !prefixSet[w[:i]] {
				valid = false
				break
			}
		}
		if valid {
			return w
		}
	}
	return ""
}
```

## 1860 — Incremental Memory Leak

```go
package main

// LeetCode #1860: Incremental Memory Leak
// https://leetcode.com/problems/incremental-memory-leak/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MemLeak(2, 2))
	fmt.Println(MemLeak(8, 11))
	fmt.Println(MemLeak(1, 1))
}

// Time: O(sqrt(memory1+memory2)), Space: O(1)
func MemLeak(memory1 int, memory2 int) []int {
	t := 1
	for memory1 >= t || memory2 >= t {
		if memory1 >= memory2 {
			memory1 -= t
		} else {
			memory2 -= t
		}
		t++
	}
	return []int{t, memory1, memory2}
}
```

## 1861 — Rotating The Box

```go
package main

// LeetCode #1861: Rotating the Box
// https://leetcode.com/problems/rotating-the-box/
// Difficulty: Medium

import "fmt"

func main() {
	box1 := [][]byte{{'#', '.', '#'}}
	fmt.Println(RotateTheBox(box1))

	box2 := [][]byte{{'#', '.', '*', '.'}, {'#', '#', '*', '.'}}
	fmt.Println(RotateTheBox(box2))

	box3 := [][]byte{{'#', '#', '*', '.', '*', '.'},
		{'#', '#', '#', '*', '.', '.'},
		{'#', '#', '#', '.', '#', '.'}}
	fmt.Println(RotateTheBox(box3))
}

// Time: O(m*n), Space: O(m*n) for result
func RotateTheBox(box [][]byte) [][]byte {
	m, n := len(box), len(box[0])

	// Apply gravity to each row (stones fall to the right)
	for i := 0; i < m; i++ {
		empty := n - 1
		for j := n - 1; j >= 0; j-- {
			if box[i][j] == '*' {
				empty = j - 1
			} else if box[i][j] == '#' {
				box[i][j] = '.'
				box[i][empty] = '#'
				empty--
			}
		}
	}

	// Rotate 90 degrees clockwise
	result := make([][]byte, n)
	for i := range result {
		result[i] = make([]byte, m)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			result[j][m-1-i] = box[i][j]
		}
	}
	return result
}
```

## 1864 — Minimum Number Of Swaps To Make The Binary String Alternating

```go
package main

// LeetCode #1864: Minimum Number of Swaps to Make the Binary String Alternating
// https://leetcode.com/problems/minimum-number-of-swaps-to-make-the-binary-string-alternating/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MinSwaps("111000"))
	fmt.Println(MinSwaps("010"))
	fmt.Println(MinSwaps("1110"))
}

// Time: O(n), Space: O(1)
func MinSwaps(s string) int {
	n := len(s)
	ones := 0
	zeros := 0
	for _, c := range s {
		if c == '1' {
			ones++
		} else {
			zeros++
		}
	}
	if abs(ones-zeros) > 1 {
		return -1
	}

	// Count mismatches when starting with '0' and starting with '1'
	mismatch0 := 0 // pattern: 010101...
	mismatch1 := 0 // pattern: 101010...
	for i, c := range s {
		if i%2 == 0 {
			if c == '1' {
				mismatch0++
			} else {
				mismatch1++
			}
		} else {
			if c == '0' {
				mismatch0++
			} else {
				mismatch1++
			}
		}
	}

	// Each swap fixes 2 mismatches
	if n%2 == 0 {
		return min(mismatch0/2, mismatch1/2)
	}
	// For odd length, only one pattern is valid
	if zeros > ones {
		return mismatch0 / 2
	}
	return mismatch1 / 2
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```

