# Hard (Sulit) — Problem 0004–0552

## 0004 — Median Of Two Sorted Arrays

```go
package main

// LeetCode #4: Median of Two Sorted Arrays
// https://leetcode.com/problems/median-of-two-sorted-arrays/
// Difficulty: Hard
//
// Binary search on the smaller array to find the correct partition such that
// all left elements <= all right elements. O(log(min(m, n))) time, O(1) space.

import (
	"fmt"
	"math"
)

func main() {
	// Example 1: nums1 = [1,3], nums2 = [2] => 2.0
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	// Example 2: nums1 = [1,2], nums2 = [3,4] => 2.5
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
	// Edge: all zeros
	fmt.Println(findMedianSortedArrays([]int{0, 0}, []int{0, 0}))
	// Edge: one empty
	fmt.Println(findMedianSortedArrays([]int{}, []int{1}))
	fmt.Println(findMedianSortedArrays([]int{2}, []int{}))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// Ensure nums1 is the smaller array for O(log(min(m,n)))
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	m, n := len(nums1), len(nums2)
	left, right := 0, m

	for left <= right {
		i := (left + right) / 2 // partition position in nums1
		j := (m+n+1)/2 - i      // partition position in nums2

		// Get bounding values with sentinels for edge partitions
		maxLeftA := math.MinInt64
		if i > 0 {
			maxLeftA = nums1[i-1]
		}

		minRightA := math.MaxInt64
		if i < m {
			minRightA = nums1[i]
		}

		maxLeftB := math.MinInt64
		if j > 0 {
			maxLeftB = nums2[j-1]
		}

		minRightB := math.MaxInt64
		if j < n {
			minRightB = nums2[j]
		}

		if maxLeftA <= minRightB && maxLeftB <= minRightA {
			// Found the correct partition
			if (m+n)%2 == 1 {
				return float64(max(maxLeftA, maxLeftB))
			}
			return (float64(max(maxLeftA, maxLeftB)) + float64(min(minRightA, minRightB))) / 2.0
		} else if maxLeftA > minRightB {
			// Too far right — move partition left
			right = i - 1
		} else {
			// Too far left — move partition right
			left = i + 1
		}
	}

	return 0.0 // unreachable for valid input
}
```

## 0010 — Regular Expression Matching

```go
package main

// LeetCode #10: Regular Expression Matching
// https://leetcode.com/problems/regular-expression-matching/
// Difficulty: Hard

import "fmt"

// isMatch checks if string s matches pattern p.
// '.' matches any single character.
// '*' matches zero or more of the preceding element.
//
// Complexity: O(m*n) time, O(m*n) space where m = len(s), n = len(p)
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}

	// empty string matches empty pattern
	dp[0][0] = true

	// handle patterns like a*, a*b*, a*b*c* matching empty string
	for j := 1; j <= n; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '.' || p[j-1] == s[i-1] {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				// zero occurrences: skip the pattern char and *
				dp[i][j] = dp[i][j-2]
				// one or more occurrences: if preceding char matches s[i-1]
				if p[j-2] == '.' || p[j-2] == s[i-1] {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			}
		}
	}

	return dp[m][n]
}

func main() {
	// Test cases from LeetCode
	fmt.Println("Test: aa, a ->", isMatch("aa", "a"))           // false
	fmt.Println("Test: aa, a* ->", isMatch("aa", "a*"))          // true
	fmt.Println("Test: ab, .* ->", isMatch("ab", ".*"))          // true
	fmt.Println("Test: aab, c*a*b ->", isMatch("aab", "c*a*b")) // true
	fmt.Println("Test: mississippi, mis*is*p*. ->", isMatch("mississippi", "mis*is*p*.")) // false
}
```

## 0023 — Merge K Sorted Lists

```go
package main

// LeetCode #23: Merge k Sorted Lists
// https://leetcode.com/problems/merge-k-sorted-lists/
// Difficulty: Hard

import (
	"container/heap"
	"fmt"
)

// ListNode represents a node in a singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// MinHeap implements heap.Interface for ListNode pointers.
type MinHeap []*ListNode

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].Val < h[j].Val }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(*ListNode)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// mergeKLists merges k sorted linked lists using a min-heap.
//
// Complexity: O(N log k) time, O(k) space where N = total nodes, k = number of lists
func mergeKLists(lists []*ListNode) *ListNode {
	h := &MinHeap{}
	heap.Init(h)

	// push the head of each non-empty list into the heap
	for _, list := range lists {
		if list != nil {
			heap.Push(h, list)
		}
	}

	dummy := &ListNode{}
	curr := dummy

	for h.Len() > 0 {
		node := heap.Pop(h).(*ListNode)
		curr.Next = node
		curr = curr.Next
		if node.Next != nil {
			heap.Push(h, node.Next)
		}
	}

	return dummy.Next
}

// Helper to build a linked list from a slice.
func buildList(vals []int) *ListNode {
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

// Helper to convert a linked list to a slice.
func listToSlice(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func main() {
	// Test case from LeetCode: [[1,4,5],[1,3,4],[2,6]] -> [1,1,2,3,4,4,5,6]
	lists := []*ListNode{
		buildList([]int{1, 4, 5}),
		buildList([]int{1, 3, 4}),
		buildList([]int{2, 6}),
	}
	result := mergeKLists(lists)
	fmt.Println("Merged list:", listToSlice(result))

	// Edge case: empty lists
	var emptyLists []*ListNode
	result2 := mergeKLists(emptyLists)
	fmt.Println("Empty input:", listToSlice(result2))

	// Edge case: single list
	single := []*ListNode{buildList([]int{1})}
	result3 := mergeKLists(single)
	fmt.Println("Single list:", listToSlice(result3))
}
```

## 0025 — Reverse Nodes In K Group

```go
package main

// LeetCode #25: Reverse Nodes in k-Group
// https://leetcode.com/problems/reverse-nodes-in-k-group/
// Difficulty: Hard

import "fmt"

// ListNode represents a node in a singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// reverseKGroup reverses the nodes of a linked list k at a time.
//
// Complexity: O(n) time, O(1) space
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k <= 1 {
		return head
	}

	// Count nodes to know where to stop
	count := 0
	curr := head
	for curr != nil {
		count++
		curr = curr.Next
	}

	dummy := &ListNode{Next: head}
	prev := dummy

	for count >= k {
		// Reverse k nodes starting from prev.Next
		start := prev.Next
		curr = start
		var prevNode *ListNode
		for i := 0; i < k; i++ {
			nextTemp := curr.Next
			curr.Next = prevNode
			prevNode = curr
			curr = nextTemp
		}
		// Connect the reversed segment
		start.Next = curr
		prev.Next = prevNode
		prev = start
		count -= k
	}

	return dummy.Next
}

// Helper to build a linked list from a slice.
func buildList(vals []int) *ListNode {
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

// Helper to convert a linked list to a slice.
func listToSlice(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func main() {
	// Test case from LeetCode: [1,2,3,4,5], k=2 -> [2,1,4,3,5]
	list1 := buildList([]int{1, 2, 3, 4, 5})
	result1 := reverseKGroup(list1, 2)
	fmt.Println("k=2:", listToSlice(result1))

	// Test case from LeetCode: [1,2,3,4,5], k=3 -> [3,2,1,4,5]
	list2 := buildList([]int{1, 2, 3, 4, 5})
	result2 := reverseKGroup(list2, 3)
	fmt.Println("k=3:", listToSlice(result2))

	// Edge case: k=1 (no change)
	list3 := buildList([]int{1, 2, 3})
	result3 := reverseKGroup(list3, 1)
	fmt.Println("k=1:", listToSlice(result3))
}
```

## 0030 — Substring With Concatenation Of All Words

```go
package main

// LeetCode #30: Substring with Concatenation of All Words
// https://leetcode.com/problems/substring-with-concatenation-of-all-words/
// Difficulty: Hard

import "fmt"

// findSubstring finds all starting indices where concatenation of all words matches.
// All words have the same length.
//
// Complexity: O(n * m) time, O(k) space where n = len(s), m = len(words), k = number of unique words
func findSubstring(s string, words []string) []int {
	if len(s) == 0 || len(words) == 0 {
		return []int{}
	}

	wordLen := len(words[0])
	totalLen := wordLen * len(words)

	if len(s) < totalLen {
		return []int{}
	}

	// Build frequency map for words
	wordFreq := make(map[string]int)
	for _, w := range words {
		wordFreq[w]++
	}

	var result []int

	// Slide over the string in wordLen-sized chunks
	for i := 0; i < wordLen; i++ {
		left := i
		right := i
		windowFreq := make(map[string]int)
		count := 0

		for right+wordLen <= len(s) {
			word := s[right : right+wordLen]
			right += wordLen

			if _, exists := wordFreq[word]; exists {
				windowFreq[word]++
				count++

				// If we have too many of this word, slide left
				for windowFreq[word] > wordFreq[word] {
					leftWord := s[left : left+wordLen]
					windowFreq[leftWord]--
					count--
					left += wordLen
				}

				// If we've matched all words, record the starting index
				if count == len(words) {
					result = append(result, left)
					// Slide left window by one word
					leftWord := s[left : left+wordLen]
					windowFreq[leftWord]--
					count--
					left += wordLen
				}
			} else {
				// Reset window
				windowFreq = make(map[string]int)
				count = 0
				left = right
			}
		}
	}

	return result
}

func main() {
	// Test case from LeetCode
	fmt.Println("Test 1:", findSubstring("barfoothefoobarman", []string{"foo", "bar"})) // [0, 9]

	// Additional test cases
	fmt.Println("Test 2:", findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"})) // []
	fmt.Println("Test 3:", findSubstring("barfoofoobarthefoobarman", []string{"bar", "foo", "the"}))            // [6, 9, 12]
}
```

## 0032 — Longest Valid Parentheses

```go
package main

// LeetCode #32: Longest Valid Parentheses
// https://leetcode.com/problems/longest-valid-parentheses/
// Difficulty: Hard

import "fmt"

// longestValidParentheses finds the length of the longest valid parentheses substring.
// Uses a stack-based approach.
//
// Complexity: O(n) time, O(n) space
func longestValidParentheses(s string) int {
	stack := []int{-1} // base index for valid substring calculation
	maxLen := 0

	for i, ch := range s {
		if ch == '(' {
			stack = append(stack, i)
		} else {
			// Pop the matching '('
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				// No matching '('; set new base index
				stack = append(stack, i)
			} else {
				// Calculate length of current valid substring
				length := i - stack[len(stack)-1]
				if length > maxLen {
					maxLen = length
				}
			}
		}
	}

	return maxLen
}

func main() {
	// Test cases from LeetCode
	fmt.Println("Test 1: (() ->", longestValidParentheses("(()"))       // 2
	fmt.Println("Test 2: )()()) ->", longestValidParentheses(")()())"))  // 4
	fmt.Println("Test 3: '' ->", longestValidParentheses(""))           // 0
	fmt.Println("Test 4: ()() ->", longestValidParentheses("()()"))     // 4
	fmt.Println("Test 5: (()()) ->", longestValidParentheses("(()())")) // 6
}
```

## 0037 — Sudoku Solver

```go
package main

// LeetCode #37: Sudoku Solver
// https://leetcode.com/problems/sudoku-solver/
// Difficulty: Hard

import "fmt"

// solveSudoku solves a 9x9 Sudoku board in-place using backtracking.
//
// Complexity: O(9^(n)) time where n is number of empty cells, O(n) space for recursion
func solveSudoku(board [][]byte) {
	// Pre-compute rows, cols, boxes for O(1) constraint checking
	var rows [9][9]bool
	var cols [9][9]bool
	var boxes [9][9]bool

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				num := board[i][j] - '1'
				boxIdx := (i/3)*3 + j/3
				rows[i][num] = true
				cols[j][num] = true
				boxes[boxIdx][num] = true
			}
		}
	}

	var backtrack func() bool
	backtrack = func() bool {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				if board[i][j] == '.' {
					boxIdx := (i/3)*3 + j/3
					for num := byte(0); num < 9; num++ {
						if !rows[i][num] && !cols[j][num] && !boxes[boxIdx][num] {
							board[i][j] = num + '1'
							rows[i][num] = true
							cols[j][num] = true
							boxes[boxIdx][num] = true

							if backtrack() {
								return true
							}

							board[i][j] = '.'
							rows[i][num] = false
							cols[j][num] = false
							boxes[boxIdx][num] = false
						}
					}
					return false
				}
			}
		}
		return true
	}

	backtrack()
}

func printBoard(board [][]byte) {
	for i := 0; i < 9; i++ {
		if i%3 == 0 && i > 0 {
			fmt.Println("------+-------+------")
		}
		for j := 0; j < 9; j++ {
			if j%3 == 0 && j > 0 {
				fmt.Print("| ")
			}
			fmt.Print(string(board[i][j]), " ")
		}
		fmt.Println()
	}
}

func main() {
	// Test case from LeetCode
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	fmt.Println("Input:")
	printBoard(board)

	solveSudoku(board)

	fmt.Println("\nSolution:")
	printBoard(board)
}
```

## 0041 — First Missing Positive

```go
package main

// LeetCode #41: First Missing Positive
// https://leetcode.com/problems/first-missing-positive/
// Difficulty: Hard

import "fmt"

// firstMissingPositive finds the smallest missing positive integer.
// Uses cyclic sort (O(1) extra space by modifying the input array).
//
// Complexity: O(n) time, O(1) space
func firstMissingPositive(nums []int) int {
	n := len(nums)

	// Cyclic sort: place each number at its correct index (num-1)
	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	// Find the first index where the value is not (index+1)
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	return n + 1
}

func main() {
	// Test cases from LeetCode
	fmt.Println("Test 1: [1,2,0] ->", firstMissingPositive([]int{1, 2, 0}))          // 3
	fmt.Println("Test 2: [3,4,-1,1] ->", firstMissingPositive([]int{3, 4, -1, 1}))   // 2
	fmt.Println("Test 3: [7,8,9,11,12] ->", firstMissingPositive([]int{7, 8, 9, 11, 12})) // 1
	fmt.Println("Test 4: [1,2,3] ->", firstMissingPositive([]int{1, 2, 3}))          // 4
	fmt.Println("Test 5: [] ->", firstMissingPositive([]int{}))                       // 1
}
```

## 0042 — Trapping Rain Water

```go
package main

// LeetCode #42: Trapping Rain Water
// https://leetcode.com/problems/trapping-rain-water/
// Difficulty: Hard

import "fmt"

// trap calculates the total amount of water that can be trapped between bars.
// Uses the two-pointer approach.
//
// Complexity: O(n) time, O(1) space
func trap(height []int) int {
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	total := 0

	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				total += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				total += rightMax - height[right]
			}
			right--
		}
	}

	return total
}

func main() {
	// Test case from LeetCode
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println("Test 1:", trap(height)) // 6

	// Additional test cases
	fmt.Println("Test 2: [4,2,0,3,2,5] ->", trap([]int{4, 2, 0, 3, 2, 5})) // 9
	fmt.Println("Test 3: [1,2,3] ->", trap([]int{1, 2, 3}))                // 0
	fmt.Println("Test 4: [] ->", trap([]int{}))                            // 0
}
```

## 0044 — Wildcard Matching

```go
package main

// LeetCode #44: Wildcard Matching
// https://leetcode.com/problems/wildcard-matching/
// Difficulty: Hard

import "fmt"

// isMatchWildcard checks if string s matches pattern p.
// '?' matches any single character.
// '*' matches any sequence of characters (including empty).
//
// Complexity: O(m*n) time, O(m*n) space where m = len(s), n = len(p)
func isMatchWildcard(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}

	// empty string matches empty pattern
	dp[0][0] = true

	// handle patterns starting with '*' matching empty string
	for j := 1; j <= n; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-1]
		}
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '?' || p[j-1] == s[i-1] {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				// '*' matches empty (dp[i][j-1]) or one/more chars (dp[i-1][j])
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			}
		}
	}

	return dp[m][n]
}

func main() {
	// Test cases from LeetCode
	fmt.Println("Test 1: aa, a ->", isMatchWildcard("aa", "a"))     // false
	fmt.Println("Test 2: aa, * ->", isMatchWildcard("aa", "*"))     // true
	fmt.Println("Test 3: cb, ?a ->", isMatchWildcard("cb", "?a"))   // false
	fmt.Println("Test 4: adceb, *a*b ->", isMatchWildcard("adceb", "*a*b")) // true
	fmt.Println("Test 5: acdcb, a*c?b ->", isMatchWildcard("acdcb", "a*c?b")) // false
}
```

## 0051 — N Queens

```go
package main

// LeetCode #51: N-Queens
// https://leetcode.com/problems/n-queens/
// Difficulty: Hard

import "fmt"

// solveNQueens returns all distinct solutions to the N-Queens puzzle.
// Each solution is represented as a board where 'Q' marks a queen and '.' marks empty.
//
// Complexity: O(n!) time, O(n) space for recursion/board (excluding output)
func solveNQueens(n int) [][]string {
	var result [][]string
	board := make([]string, n)
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		board[i] = string(row)
	}

	cols := make([]bool, n)
	d1 := make([]bool, 2*n-1) // diagonal: row - col + n - 1
	d2 := make([]bool, 2*n-1) // anti-diagonal: row + col

	var backtrack func(row int)
	backtrack = func(row int) {
		if row == n {
			solution := make([]string, n)
			copy(solution, board)
			result = append(result, solution)
			return
		}

		for col := 0; col < n; col++ {
			idx1 := row - col + n - 1
			idx2 := row + col
			if cols[col] || d1[idx1] || d2[idx2] {
				continue
			}

			// Place queen
			r := []byte(board[row])
			r[col] = 'Q'
			board[row] = string(r)
			cols[col] = true
			d1[idx1] = true
			d2[idx2] = true

			backtrack(row + 1)

			// Remove queen
			r[col] = '.'
			board[row] = string(r)
			cols[col] = false
			d1[idx1] = false
			d2[idx2] = false
		}
	}

	backtrack(0)
	return result
}

func main() {
	// Test case from LeetCode: n=4 -> 2 solutions
	solutions := solveNQueens(4)
	fmt.Printf("n=4 has %d solutions:\n", len(solutions))
	for i, sol := range solutions {
		fmt.Printf("Solution %d:\n", i+1)
		for _, row := range sol {
			fmt.Println(row)
		}
		fmt.Println()
	}

	// Test n=1
	solutions1 := solveNQueens(1)
	fmt.Printf("n=1 has %d solutions\n", len(solutions1))
}
```

## 0052 — N Queens Ii

```go
package main

// LeetCode #52: N-Queens II
// https://leetcode.com/problems/n-queens-ii/
// Difficulty: Hard

import "fmt"

func main() {
	fmt.Println("52. N-Queens II")
	fmt.Println("n=4:", totalNQueens(4), "(expected 2)")
	fmt.Println("n=1:", totalNQueens(1), "(expected 1)")
	fmt.Println("n=8:", totalNQueens(8), "(expected 92)")
}

func totalNQueens(n int) int {
	cols := make([]bool, n)
	diag1 := make([]bool, 2*n-1) // r+c
	diag2 := make([]bool, 2*n-1) // r-c+n-1
	count := 0
	backtrack(n, 0, cols, diag1, diag2, &count)
	return count
}

func backtrack(n, row int, cols, diag1, diag2 []bool, count *int) {
	if row == n {
		*count++
		return
	}
	for col := 0; col < n; col++ {
		d1 := row + col
		d2 := row - col + n - 1
		if cols[col] || diag1[d1] || diag2[d2] {
			continue
		}
		cols[col], diag1[d1], diag2[d2] = true, true, true
		backtrack(n, row+1, cols, diag1, diag2, count)
		cols[col], diag1[d1], diag2[d2] = false, false, false
	}
}
```

## 0060 — Permutation Sequence

```go
package main

// LeetCode #60: Permutation Sequence
// https://leetcode.com/problems/permutation-sequence/
// Difficulty: Hard

import "fmt"

func main() {
	fmt.Println("60. Permutation Sequence")
	fmt.Println("n=3, k=3:", getPermutation(3, 3), "(expected 213)")
	fmt.Println("n=4, k=9:", getPermutation(4, 9), "(expected 2314)")
	fmt.Println("n=3, k=1:", getPermutation(3, 1), "(expected 123)")
}

func getPermutation(n int, k int) string {
	fact := 1
	nums := make([]byte, 0, n)
	for i := 1; i <= n; i++ {
		fact *= i
		nums = append(nums, byte('0'+i))
	}

	k-- // convert to 0-indexed
	result := make([]byte, 0, n)

	for i := 0; i < n; i++ {
		fact /= (n - i)
		idx := k / fact
		result = append(result, nums[idx])
		nums = append(nums[:idx], nums[idx+1:]...)
		k %= fact
	}

	return string(result)
}
```

## 0065 — Valid Number

```go
package main

// LeetCode #65: Valid Number
// https://leetcode.com/problems/valid-number/
// Difficulty: Hard

import "fmt"

func main() {
	fmt.Println("65. Valid Number")
	testCases := []struct {
		s string
		r bool
	}{
		{"0", true},
		{"0.1", true},
		{"abc", false},
		{"2e10", true},
		{"2e", false},
		{"e3", false},
		{".", false},
		{"1.", true},
		{".1", true},
		{"-.1", true},
		{"+.", false},
		{"3.", true},
		{"3.e10", true},
		{"3.e", false},
		{".e1", false},
		{"+.8", true},
		{"46.e3", true},
		{" 1", true},
		{"1 ", true},
		{"1 1", false},
		{"", false},
	}
	for _, tc := range testCases {
		got := isNumber(tc.s)
		status := "OK"
		if got != tc.r {
			status = "FAIL"
		}
		fmt.Printf("  %q -> %5v (expected %5v) [%s]\n", tc.s, got, tc.r, status)
	}
}

// DFA states:
//
//	0 start (allow whitespace)
//	1 sign before digits
//	2 integer part digits
//	3 decimal point (no digits after yet)
//	4 fractional part digits
//	5 'e' or 'E'
//	6 sign after 'e'
//	7 exponent digits
//	8 trailing whitespace
//	-1 invalid
func isNumber(s string) bool {
	state := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch state {
		case 0: // leading whitespace
			if c == ' ' {
				continue
			} else if c == '+' || c == '-' {
				state = 1
			} else if c >= '0' && c <= '9' {
				state = 2
			} else if c == '.' {
				state = 3
			} else {
				return false
			}
		case 1: // sign before digits
			if c >= '0' && c <= '9' {
				state = 2
			} else if c == '.' {
				state = 3
			} else {
				return false
			}
		case 2: // integer part digits
			if c >= '0' && c <= '9' {
				// stay
			} else if c == '.' {
				state = 4
			} else if c == 'e' || c == 'E' {
				state = 5
			} else if c == ' ' {
				state = 8
			} else {
				return false
			}
		case 3: // decimal point without integer digits
			if c >= '0' && c <= '9' {
				state = 4
			} else {
				return false
			}
		case 4: // fractional part digits
			if c >= '0' && c <= '9' {
				// stay
			} else if c == 'e' || c == 'E' {
				state = 5
			} else if c == ' ' {
				state = 8
			} else {
				return false
			}
		case 5: // 'e'/'E'
			if c >= '0' && c <= '9' {
				state = 7
			} else if c == '+' || c == '-' {
				state = 6
			} else {
				return false
			}
		case 6: // sign after 'e'
			if c >= '0' && c <= '9' {
				state = 7
			} else {
				return false
			}
		case 7: // exponent digits
			if c >= '0' && c <= '9' {
				// stay
			} else if c == ' ' {
				state = 8
			} else {
				return false
			}
		case 8: // trailing whitespace
			if c == ' ' {
				// stay
			} else {
				return false
			}
		default:
			return false
		}
	}
	return state == 2 || state == 4 || state == 7 || state == 8
}
```

## 0068 — Text Justification

```go
package main

// LeetCode #68: Text Justification
// https://leetcode.com/problems/text-justification/
// Difficulty: Hard

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("68. Text Justification")
	words := []string{"This", "is", "an", "example", "of", "text", "justification."}
	result := fullJustify(words, 16)
	fmt.Printf("maxWidth=16 -> %q\n", result)
	for _, line := range result {
		fmt.Printf("  %q (len=%d)\n", line, len(line))
	}

	words2 := []string{"What", "must", "be", "acknowledgment", "shall", "be"}
	result2 := fullJustify(words2, 16)
	fmt.Printf("\nmaxWidth=16 -> %q\n", result2)
	for _, line := range result2 {
		fmt.Printf("  %q (len=%d)\n", line, len(line))
	}

	words3 := []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}
	result3 := fullJustify(words3, 20)
	fmt.Printf("\nmaxWidth=20 -> %q\n", result3)
	for _, line := range result3 {
		fmt.Printf("  %q (len=%d)\n", line, len(line))
	}
}

func fullJustify(words []string, maxWidth int) []string {
	var result []string
	i := 0
	for i < len(words) {
		j := i + 1
		lineLen := len(words[i])
		for j < len(words) && lineLen+1+len(words[j]) <= maxWidth {
			lineLen += 1 + len(words[j])
			j++
		}

		line := ""
		wordCount := j - i
		spaceSlots := wordCount - 1

		if j == len(words) || wordCount == 1 {
			// left-justified (last line or single word)
			line = strings.Join(words[i:j], " ")
			line += strings.Repeat(" ", maxWidth-len(line))
		} else {
			totalSpaces := maxWidth - (lineLen - spaceSlots) // total spaces needed
			baseSpaces := totalSpaces / spaceSlots
			extraSpaces := totalSpaces % spaceSlots

			var sb strings.Builder
			for k := i; k < j; k++ {
				sb.WriteString(words[k])
				if k < j-1 {
					spaces := baseSpaces
					if k-i < extraSpaces {
						spaces++
					}
					sb.WriteString(strings.Repeat(" ", spaces))
				}
			}
			line = sb.String()
		}
		result = append(result, line)
		i = j
	}
	return result
}
```

## 0076 — Minimum Window Substring

```go
package main

// LeetCode #76: Minimum Window Substring
// https://leetcode.com/problems/minimum-window-substring/
// Difficulty: Hard

import "fmt"

func main() {
	fmt.Println("76. Minimum Window Substring")
	fmt.Printf("s=ADOBECODEBANC, t=ABC -> %q (expected BANC)\n", minWindow("ADOBECODEBANC", "ABC"))
	fmt.Printf("s=a, t=a -> %q (expected a)\n", minWindow("a", "a"))
	fmt.Printf("s=a, t=aa -> %q (expected empty)\n", minWindow("a", "aa"))
}

func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	need := [128]int{}
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	have := [128]int{}
	haveCount := 0
	needCount := 0
	for _, v := range need {
		if v > 0 {
			needCount++
		}
	}

	left := 0
	minLen := len(s) + 1
	start := 0

	for right := 0; right < len(s); right++ {
		c := s[right]
		have[c]++
		if have[c] == need[c] {
			haveCount++
		}

		for haveCount == needCount {
			if right-left+1 < minLen {
				minLen = right - left + 1
				start = left
			}
			lc := s[left]
			if have[lc] == need[lc] {
				haveCount--
			}
			have[lc]--
			left++
		}
	}

	if minLen > len(s) {
		return ""
	}
	return s[start : start+minLen]
}
```

## 0084 — Largest Rectangle In Histogram

```go
package main

// LeetCode #84: Largest Rectangle in Histogram
// https://leetcode.com/problems/largest-rectangle-in-histogram/
// Difficulty: Hard

import "fmt"

func main() {
	fmt.Println("84. Largest Rectangle in Histogram")
	fmt.Println("[2,1,5,6,2,3] ->", largestRectangleArea([]int{2, 1, 5, 6, 2, 3}), "(expected 10)")
	fmt.Println("[2,4] ->", largestRectangleArea([]int{2, 4}), "(expected 4)")
	fmt.Println("[2,1,2] ->", largestRectangleArea([]int{2, 1, 2}), "(expected 3)")
}

func largestRectangleArea(heights []int) int {
	n := len(heights)
	stack := make([]int, 0, n)
	maxArea := 0

	for i := 0; i <= n; i++ {
		var h int
		if i == n {
			h = 0
		} else {
			h = heights[i]
		}

		for len(stack) > 0 && h < heights[stack[len(stack)-1]] {
			height := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]

			width := i
			if len(stack) > 0 {
				width = i - stack[len(stack)-1] - 1
			}

			area := height * width
			if area > maxArea {
				maxArea = area
			}
		}
		stack = append(stack, i)
	}

	return maxArea
}
```

## 0085 — Maximal Rectangle

```go
package main

// LeetCode #85: Maximal Rectangle
// https://leetcode.com/problems/maximal-rectangle/
// Difficulty: Hard

import "fmt"

func main() {
	fmt.Println("85. Maximal Rectangle")
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	fmt.Println("Example 1 ->", maximalRectangle(matrix), "(expected 6)")

	matrix2 := [][]byte{{'0'}}
	fmt.Println("Example 2 ->", maximalRectangle(matrix2), "(expected 0)")

	matrix3 := [][]byte{{'1'}}
	fmt.Println("Example 3 ->", maximalRectangle(matrix3), "(expected 1)")
}

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	rows, cols := len(matrix), len(matrix[0])
	heights := make([]int, cols)
	maxArea := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if matrix[r][c] == '1' {
				heights[c]++
			} else {
				heights[c] = 0
			}
		}
		area := largestRectangleArea(heights)
		if area > maxArea {
			maxArea = area
		}
	}

	return maxArea
}

func largestRectangleArea(heights []int) int {
	n := len(heights)
	stack := make([]int, 0, n)
	maxArea := 0

	for i := 0; i <= n; i++ {
		var h int
		if i == n {
			h = 0
		} else {
			h = heights[i]
		}

		for len(stack) > 0 && h < heights[stack[len(stack)-1]] {
			height := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]

			width := i
			if len(stack) > 0 {
				width = i - stack[len(stack)-1] - 1
			}

			area := height * width
			if area > maxArea {
				maxArea = area
			}
		}
		stack = append(stack, i)
	}

	return maxArea
}
```

## 0087 — Scramble String

```go
package main

// LeetCode #87: Scramble String
// https://leetcode.com/problems/scramble-string/
// Difficulty: Hard

import "fmt"

func main() {
	fmt.Println("87. Scramble String")
	fmt.Println("great, rgeat ->", isScramble("great", "rgeat"), "(expected true)")
	fmt.Println("abcde, caebd ->", isScramble("abcde", "caebd"), "(expected false)")
	fmt.Println("a, a ->", isScramble("a", "a"), "(expected true)")
	fmt.Println("ab, ba ->", isScramble("ab", "ba"), "(expected true)")
}

func isScramble(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	if len(s1) != len(s2) {
		return false
	}

	n := len(s1)

	// memo[key] where key encodes (s1_start, s1_end, s2_start, s2_end)
	// but we can use a simpler DP: dp[k][i][j] = isScramble(s1[i:i+k], s2[j:j+k])
	dp := make([][][]bool, n+1)
	for k := 0; k <= n; k++ {
		dp[k] = make([][]bool, n)
		for i := 0; i < n; i++ {
			dp[k][i] = make([]bool, n)
		}
	}

	// k=1 base case
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			dp[1][i][j] = s1[i] == s2[j]
		}
	}

	// k from 2 to n
	for k := 2; k <= n; k++ {
		for i := 0; i <= n-k; i++ {
			for j := 0; j <= n-k; j++ {
				for split := 1; split < k; split++ {
					// no swap: s1[i:i+split] ~ s2[j:j+split] && s1[i+split:i+k] ~ s2[j+split:j+k]
					if dp[split][i][j] && dp[k-split][i+split][j+split] {
						dp[k][i][j] = true
						break
					}
					// swap: s1[i:i+split] ~ s2[j+k-split:j+k] && s1[i+split:i+k] ~ s2[j:j+k-split]
					if dp[split][i][j+k-split] && dp[k-split][i+split][j] {
						dp[k][i][j] = true
						break
					}
				}
			}
		}
	}

	return dp[n][0][0]
}
```

## 0115 — Distinct Subsequences

```go
package main

// LeetCode #115: Distinct Subsequences
// https://leetcode.com/problems/distinct-subsequences/
// Difficulty: Hard

import "fmt"

func main() {
	fmt.Println("115. Distinct Subsequences")
	fmt.Println("rabbbit, rabbit ->", numDistinct("rabbbit", "rabbit"), "(expected 3)")
	fmt.Println("babgbag, bag ->", numDistinct("babgbag", "bag"), "(expected 5)")
	fmt.Println("r, r ->", numDistinct("r", "r"), "(expected 1)")
}

func numDistinct(s string, t string) int {
	m, n := len(s), len(t)
	if n == 0 {
		return 1
	}
	if m < n {
		return 0
	}

	// dp[j] = number of distinct subsequences of s[:i] that equal t[:j]
	dp := make([]int, n+1)
	dp[0] = 1

	for i := 1; i <= m; i++ {
		prev := dp[0]
		for j := 1; j <= n; j++ {
			curr := dp[j]
			if s[i-1] == t[j-1] {
				dp[j] = prev + dp[j]
			}
			prev = curr
		}
	}

	return dp[n]
}
```

## 0123 — Best Time To Buy And Sell Stock Iii

```go
package main

// LeetCode #123: Best Time to Buy and Sell Stock III
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/
// Difficulty: Hard

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("123. Best Time to Buy and Sell Stock III")
	fmt.Println("[3,3,5,0,0,3,1,4] ->", maxProfit([]int{3, 3, 5, 0, 0, 3, 1, 4}), "(expected 6)")
	fmt.Println("[1,2,3,4,5] ->", maxProfit([]int{1, 2, 3, 4, 5}), "(expected 4)")
	fmt.Println("[7,6,4,3,1] ->", maxProfit([]int{7, 6, 4, 3, 1}), "(expected 0)")
	fmt.Println("[1] ->", maxProfit([]int{1}), "(expected 0)")
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	// State machine with at most 2 transactions:
	// buy1 = min price seen for first buy
	// sell1 = max profit after first sell
	// buy2 = effective cost after first profit (min price2 - profit1)
	// sell2 = max profit after second sell
	buy1 := math.MaxInt32
	buy2 := math.MaxInt32
	sell1 := 0
	sell2 := 0

	for _, price := range prices {
		if price < buy1 {
			buy1 = price
		}
		if price-buy1 > sell1 {
			sell1 = price - buy1
		}
		if price-sell1 < buy2 {
			buy2 = price - sell1
		}
		if price-buy2 > sell2 {
			sell2 = price - buy2
		}
	}

	return sell2
}
```

## 0124 — Binary Tree Maximum Path Sum

```go
package main

// LeetCode #124: Binary Tree Maximum Path Sum
// https://leetcode.com/problems/binary-tree-maximum-path-sum/
// Difficulty: Hard

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	maxSum := root.Val

	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftGain := max(0, dfs(node.Left))
		rightGain := max(0, dfs(node.Right))

		currentPathSum := node.Val + leftGain + rightGain
		if currentPathSum > maxSum {
			maxSum = currentPathSum
		}

		return node.Val + max(leftGain, rightGain)
	}

	dfs(root)
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Example: [-10,9,20,null,null,15,7] -> 42
	root := &TreeNode{Val: -10}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}

	result := maxPathSum(root)
	expected := 42

	fmt.Printf("maxPathSum([-10,9,20,null,null,15,7]) = %d\n", result)
	if result == expected {
		fmt.Println("PASS")
	} else {
		fmt.Printf("FAIL: expected %d\n", expected)
	}
}
```

## 0126 — Word Ladder Ii

```go
package main

// LeetCode #126: Word Ladder II
// https://leetcode.com/problems/word-ladder-ii/
// Difficulty: Hard

import (
	"fmt"
)

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	wordSet := make(map[string]bool)
	for _, w := range wordList {
		wordSet[w] = true
	}
	if !wordSet[endWord] {
		return [][]string{}
	}

	// BFS to find shortest distances from beginWord
	dist := map[string]int{beginWord: 0}
	queue := []string{beginWord}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		currDist := dist[curr]
		if curr == endWord {
			break
		}
		neighbors := getNeighbors(curr, wordSet)
		for _, nb := range neighbors {
			if _, seen := dist[nb]; !seen {
				dist[nb] = currDist + 1
				queue = append(queue, nb)
			}
		}
	}

	if _, reached := dist[endWord]; !reached {
		return [][]string{}
	}

	// DFS to reconstruct all shortest paths
	result := [][]string{}
	path := []string{beginWord}

	var dfs func(string)
	dfs = func(word string) {
		if word == endWord {
			tmp := make([]string, len(path))
			copy(tmp, path)
			result = append(result, tmp)
			return
		}
		for _, nb := range getNeighbors(word, wordSet) {
			if d, ok := dist[nb]; ok && d == dist[word]+1 {
				path = append(path, nb)
				dfs(nb)
				path = path[:len(path)-1]
			}
		}
	}

	dfs(beginWord)
	return result
}

func getNeighbors(word string, wordSet map[string]bool) []string {
	neighbors := []string{}
	bytes := []byte(word)
	for i := 0; i < len(bytes); i++ {
		original := bytes[i]
		for c := 'a'; c <= 'z'; c++ {
			bytes[i] = byte(c)
			candidate := string(bytes)
			if candidate != word && wordSet[candidate] {
				neighbors = append(neighbors, candidate)
			}
		}
		bytes[i] = original
	}
	return neighbors
}

func main() {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}

	result := findLadders(beginWord, endWord, wordList)
	fmt.Printf("findLadders(%q, %q, %v) = %v\n", beginWord, endWord, wordList, result)

	// Expected: [["hit","hot","dot","dog","cog"],["hit","hot","lot","log","cog"]]
	expectedCount := 2
	if len(result) == expectedCount {
		fmt.Println("PASS")
	} else {
		fmt.Printf("FAIL: expected %d paths, got %d\n", expectedCount, len(result))
	}
}
```

## 0127 — Word Ladder

```go
package main

// LeetCode #127: Word Ladder
// https://leetcode.com/problems/word-ladder/
// Difficulty: Hard

import (
	"fmt"
)

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordSet := make(map[string]bool)
	for _, w := range wordList {
		wordSet[w] = true
	}
	if !wordSet[endWord] {
		return 0
	}

	// Bidirectional BFS
	beginSet := map[string]bool{beginWord: true}
	endSet := map[string]bool{endWord: true}
	visited := map[string]bool{beginWord: true, endWord: true}

	dist := 1

	for len(beginSet) > 0 && len(endSet) > 0 {
		// Always expand the smaller set
		if len(beginSet) > len(endSet) {
			beginSet, endSet = endSet, beginSet
		}

		nextSet := make(map[string]bool)
		for word := range beginSet {
			neighbors := getNeighbors(word, wordSet)
			for _, nb := range neighbors {
				if endSet[nb] {
					return dist + 1
				}
				if !visited[nb] {
					visited[nb] = true
					nextSet[nb] = true
				}
			}
		}

		beginSet = nextSet
		dist++
	}

	return 0
}

func getNeighbors(word string, wordSet map[string]bool) []string {
	neighbors := []string{}
	bytes := []byte(word)
	for i := 0; i < len(bytes); i++ {
		original := bytes[i]
		for c := 'a'; c <= 'z'; c++ {
			bytes[i] = byte(c)
			candidate := string(bytes)
			if candidate != word && wordSet[candidate] {
				neighbors = append(neighbors, candidate)
			}
		}
		bytes[i] = original
	}
	return neighbors
}

func main() {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}

	result := ladderLength(beginWord, endWord, wordList)
	expected := 5

	fmt.Printf("ladderLength(%q, %q, %v) = %d\n", beginWord, endWord, wordList, result)
	if result == expected {
		fmt.Println("PASS")
	} else {
		fmt.Printf("FAIL: expected %d\n", expected)
	}
}
```

## 0132 — Palindrome Partitioning Ii

```go
package main

// LeetCode #132: Palindrome Partitioning II
// https://leetcode.com/problems/palindrome-partitioning-ii/
// Difficulty: Hard

import (
	"fmt"
	"math"
)

func minCut(s string) int {
	n := len(s)
	if n <= 1 {
		return 0
	}

	// dp[i] = min cuts for s[0:i]
	dp := make([]int, n)
	for i := range dp {
		dp[i] = math.MaxInt32
	}

	// isPalindrome[i][j] = s[i:j+1] is palindrome
	isPalindrome := make([][]bool, n)
	for i := range isPalindrome {
		isPalindrome[i] = make([]bool, n)
	}

	for end := 0; end < n; end++ {
		for start := 0; start <= end; start++ {
			if s[start] == s[end] && (end-start <= 2 || isPalindrome[start+1][end-1]) {
				isPalindrome[start][end] = true
			}
		}
	}

	for i := 0; i < n; i++ {
		if isPalindrome[0][i] {
			dp[i] = 0
		} else {
			for j := 0; j < i; j++ {
				if isPalindrome[j+1][i] && dp[j]+1 < dp[i] {
					dp[i] = dp[j] + 1
				}
			}
		}
	}

	return dp[n-1]
}

func main() {
	s := "aab"
	result := minCut(s)
	expected := 1

	fmt.Printf("minCut(%q) = %d\n", s, result)
	if result == expected {
		fmt.Println("PASS")
	} else {
		fmt.Printf("FAIL: expected %d\n", expected)
	}
}
```

## 0135 — Candy

```go
package main

// LeetCode #135: Candy
// https://leetcode.com/problems/candy/
// Difficulty: Hard

import (
	"fmt"
)

func candy(ratings []int) int {
	n := len(ratings)
	if n == 0 {
		return 0
	}

	candies := make([]int, n)
	for i := range candies {
		candies[i] = 1
	}

	// Left to right
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	// Right to left
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && candies[i] <= candies[i+1] {
			candies[i] = candies[i+1] + 1
		}
	}

	total := 0
	for _, c := range candies {
		total += c
	}
	return total
}

func main() {
	ratings := []int{1, 0, 2}
	result := candy(ratings)
	expected := 5

	fmt.Printf("candy(%v) = %d\n", ratings, result)
	if result == expected {
		fmt.Println("PASS")
	} else {
		fmt.Printf("FAIL: expected %d\n", expected)
	}
}
```

## 0140 — Word Break Ii

```go
package main

// LeetCode #140: Word Break II
// https://leetcode.com/problems/word-break-ii/
// Difficulty: Hard

import (
	"fmt"
	"strings"
)

func wordBreak(s string, wordDict []string) []string {
	wordSet := make(map[string]bool)
	for _, w := range wordDict {
		wordSet[w] = true
	}

	memo := make(map[string][]string)

	var dfs func(string) []string
	dfs = func(remaining string) []string {
		if results, ok := memo[remaining]; ok {
			return results
		}

		results := []string{}
		if wordSet[remaining] {
			results = append(results, remaining)
		}

		for i := 1; i < len(remaining); i++ {
			prefix := remaining[:i]
			if wordSet[prefix] {
				suffixSentences := dfs(remaining[i:])
				for _, sentence := range suffixSentences {
					results = append(results, prefix+" "+sentence)
				}
			}
		}

		memo[remaining] = results
		return results
	}

	result := dfs(s)

	// The problem expects sentences that don't include the last word as a standalone sentence
	// when the full string is a word itself. We need to filter if the full string matches
	// but we already handle that. But the example shows:
	// "catsanddog" -> ["cats and dog","cat sand dog"]
	// The full string "catsanddog" is NOT in the dict, so no issue in this case.

	return result
}

func main() {
	s := "catsanddog"
	wordDict := []string{"cat", "cats", "and", "sand", "dog"}

	result := wordBreak(s, wordDict)
	fmt.Printf("wordBreak(%q, %v) = %v\n", s, wordDict, result)

	// Check that result contains expected sentences (order-independent)
	expected := map[string]bool{
		"cats and dog": true,
		"cat sand dog": true,
	}

	pass := true
	if len(result) != len(expected) {
		pass = false
	} else {
		for _, r := range result {
			if !expected[r] {
				pass = false
				break
			}
		}
	}

	if pass {
		fmt.Println("PASS")
	} else {
		fmt.Printf("FAIL: expected %v\n", strings.Join(func() []string {
			keys := make([]string, 0, len(expected))
			for k := range expected {
				keys = append(keys, k)
			}
			return keys
		}(), ", "))
	}
}
```

## 0149 — Max Points On A Line

```go
package main

// LeetCode #149: Max Points on a Line
// https://leetcode.com/problems/max-points-on-a-line/
// Difficulty: Hard

import (
	"fmt"
)

func maxPoints(points [][]int) int {
	n := len(points)
	if n <= 2 {
		return n
	}

	maxCount := 0

	for i := 0; i < n; i++ {
		slopes := make(map[[2]int]int)
		duplicate := 1

		for j := i + 1; j < n; j++ {
			dx := points[j][0] - points[i][0]
			dy := points[j][1] - points[i][1]

			if dx == 0 && dy == 0 {
				duplicate++
				continue
			}

			g := gcd(dx, dy)
			dx /= g
			dy /= g

			// Normalize slope sign: ensure dx is positive, or if dx==0, dy positive
			if dx < 0 || (dx == 0 && dy < 0) {
				dx = -dx
				dy = -dy
			}

			key := [2]int{dx, dy}
			slopes[key]++
		}

		localMax := 0
		for _, count := range slopes {
			if count > localMax {
				localMax = count
			}
		}

		if localMax+duplicate > maxCount {
			maxCount = localMax + duplicate
		}
	}

	return maxCount
}

func gcd(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	points := [][]int{{1, 1}, {2, 2}, {3, 3}}
	result := maxPoints(points)
	expected := 3

	fmt.Printf("maxPoints(%v) = %d\n", points, result)
	if result == expected {
		fmt.Println("PASS")
	} else {
		fmt.Printf("FAIL: expected %d\n", expected)
	}
}
```

## 0154 — Find Minimum In Rotated Sorted Array Ii

```go
package main

// LeetCode #154: Find Minimum in Rotated Sorted Array II
// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array-ii/
// Difficulty: Hard

import (
	"fmt"
)

func findMin(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] > nums[right] {
			// Minimum is in the right half
			left = mid + 1
		} else if nums[mid] < nums[right] {
			// Minimum is in the left half (including mid)
			right = mid
		} else {
			// nums[mid] == nums[right], cannot determine, shrink
			right--
		}
	}

	return nums[left]
}

func main() {
	nums := []int{2, 2, 2, 0, 1}
	result := findMin(nums)
	expected := 0

	fmt.Printf("findMin(%v) = %d\n", nums, result)
	if result == expected {
		fmt.Println("PASS")
	} else {
		fmt.Printf("FAIL: expected %d\n", expected)
	}
}
```

## 0158 — Read N Characters Given Read4 Ii Call Multiple Times

```go
package main

import "fmt"

// LeetCode #158: Read N Characters Given Read4 II - Call Multiple Times
// https://leetcode.com/problems/read-n-characters-given-read4-ii-call-multiple-times/
// Difficulty: Hard [Paid]
//
// The key challenge: read() can be called multiple times. read4() reads up to 4 chars
// from the underlying file. Any chars read by read4() but not consumed must be buffered
// for subsequent read() calls.

// --- Simulated API -----------------------------------------------------------

var fileContent string
var filePos int

// read4 reads up to 4 characters from the simulated file into buf.
// Returns the actual number of characters read (0 when EOF).
func read4(buf []byte) int {
	n := 4
	if filePos+n > len(fileContent) {
		n = len(fileContent) - filePos
	}
	for i := 0; i < n; i++ {
		buf[i] = fileContent[filePos+i]
	}
	filePos += n
	return n
}

// --- Solution ----------------------------------------------------------------

// Solution wraps read4 with an internal buffer so that characters left over
// from a previous read4() call are used first on the next read() call.
type Solution struct {
	buf    [4]byte // internal buffer from read4
	bufPtr int     // next unconsumed position in buf
	bufCnt int     // number of valid bytes in buf (from last read4)
}

// Read reads up to n characters into buf. Returns the number of chars read.
func (s *Solution) Read(buf []byte, n int) int {
	total := 0
	for total < n {
		// Refill internal buffer when exhausted.
		if s.bufPtr >= s.bufCnt {
			s.bufCnt = read4(s.buf[:])
			s.bufPtr = 0
			if s.bufCnt == 0 {
				break // EOF
			}
		}
		// Copy from internal buffer to caller buffer.
		toCopy := min(n-total, s.bufCnt-s.bufPtr)
		for i := 0; i < toCopy; i++ {
			buf[total+i] = s.buf[s.bufPtr+i]
		}
		s.bufPtr += toCopy
		total += toCopy
	}
	return total
}

// --- Helpers -----------------------------------------------------------------

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// --- Tests -------------------------------------------------------------------

func main() {
	fmt.Println("=== 0158 Read N Characters Given Read4 II - Call Multiple Times ===")

	// Test 1: Single read consuming all content.
	fileContent = "abcde"
	filePos = 0
	sol := &Solution{}
	buf := make([]byte, 5)
	n := sol.Read(buf, 5)
	fmt.Printf("Test 1 - Single read(%d) = %d, got %q (expected 5 abcde)\n", 5, n, string(buf[:n]))

	// Test 2: Multiple sequential reads.
	fileContent = "abcdefghij"
	filePos = 0
	sol = &Solution{}
	buf1 := make([]byte, 4)
	n1 := sol.Read(buf1, 4)
	buf2 := make([]byte, 4)
	n2 := sol.Read(buf2, 4)
	buf3 := make([]byte, 4)
	n3 := sol.Read(buf3, 4)
	fmt.Printf("Test 2a - Read(4) = %d, got %q (expected 4 abcd)\n", n1, string(buf1[:n1]))
	fmt.Printf("Test 2b - Read(4) = %d, got %q (expected 4 efgh)\n", n2, string(buf2[:n2]))
	fmt.Printf("Test 2c - Read(4) = %d, got %q (expected 2 ij)\n", n3, string(buf3[:n3]))

	// Test 3: Read fewer bytes than available, then more.
	fileContent = "abcdef"
	filePos = 0
	sol = &Solution{}
	buf4 := make([]byte, 2)
	n4 := sol.Read(buf4, 2)
	fmt.Printf("Test 3a - Read(2) = %d, got %q (expected 2 ab)\n", n4, string(buf4[:n4]))
	buf5 := make([]byte, 5)
	n5 := sol.Read(buf5, 5)
	fmt.Printf("Test 3b - Read(5) = %d, got %q (expected 4 cdef)\n", n5, string(buf5[:n5]))

	// Test 4: Read zero bytes.
	fileContent = "abc"
	filePos = 0
	sol = &Solution{}
	buf6 := make([]byte, 0)
	n6 := sol.Read(buf6, 0)
	fmt.Printf("Test 4 - Read(0) = %d (expected 0)\n", n6)

	// Test 5: Empty file.
	fileContent = ""
	filePos = 0
	sol = &Solution{}
	buf7 := make([]byte, 3)
	n7 := sol.Read(buf7, 3)
	fmt.Printf("Test 5 - Read(3) on empty = %d (expected 0)\n", n7)
}
```

## 0174 — Dungeon Game

```go
package main

// LeetCode #174: Dungeon Game
// https://leetcode.com/problems/dungeon-game/
// Difficulty: Hard

import (
	"fmt"
)

func calculateMinimumHP(dungeon [][]int) int {
	if len(dungeon) == 0 || len(dungeon[0]) == 0 {
		return 1
	}

	m, n := len(dungeon), len(dungeon[0])

	// dp[i][j] = minimum health needed to reach bottom-right from (i, j)
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				// Bottom-right cell
				dp[i][j] = max(1, 1-dungeon[i][j])
			} else if i == m-1 {
				// Last row, can only go right
				dp[i][j] = max(1, dp[i][j+1]-dungeon[i][j])
			} else if j == n-1 {
				// Last column, can only go down
				dp[i][j] = max(1, dp[i+1][j]-dungeon[i][j])
			} else {
				// Can go right or down, take minimum health path
				minNext := min(dp[i][j+1], dp[i+1][j])
				dp[i][j] = max(1, minNext-dungeon[i][j])
			}
		}
	}

	return dp[0][0]
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

func main() {
	dungeon := [][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}}
	result := calculateMinimumHP(dungeon)
	expected := 7

	fmt.Printf("calculateMinimumHP(%v) = %d\n", dungeon, result)
	if result == expected {
		fmt.Println("PASS")
	} else {
		fmt.Printf("FAIL: expected %d\n", expected)
	}
}
```

## 0185 — Department Top Three Salaries

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #185: Department Top Three Salaries
// https://leetcode.com/problems/department-top-three-salaries/
// Difficulty: Hard
//
// For each department, find employees who earn one of the top 3 distinct salaries.
// If there are ties, include all employees with those salaries.

// Employee represents a row in the Employee table.
type Employee struct {
	ID           int
	Name         string
	Salary       int
	DepartmentID int
}

// Department represents a row in the Department table.
type Department struct {
	ID   int
	Name string
}

// Result represents one output row.
type Result struct {
	Department string
	Employee   string
	Salary     int
}

// departmentTopThreeSalaries returns employees with top 3 distinct salaries per dept.
// Time: O(E log E + D log D) for sorting, Space: O(E + D)
func departmentTopThreeSalaries(employees []Employee, departments []Department) []Result {
	// Build department name lookup.
	deptName := make(map[int]string)
	for _, d := range departments {
		deptName[d.ID] = d.Name
	}

	// Group employees by department.
	byDept := make(map[int][]Employee)
	for _, e := range employees {
		byDept[e.DepartmentID] = append(byDept[e.DepartmentID], e)
	}

	var results []Result

	for deptID, emps := range byDept {
		// Sort descending by salary.
		sort.Slice(emps, func(i, j int) bool {
			return emps[i].Salary > emps[j].Salary
		})

		// Collect top 3 distinct salaries.
		distinctSalaries := make([]int, 0)
		for _, e := range emps {
			if len(distinctSalaries) == 0 || e.Salary != distinctSalaries[len(distinctSalaries)-1] {
				distinctSalaries = append(distinctSalaries, e.Salary)
				if len(distinctSalaries) == 3 {
					break
				}
			}
		}

		// Build a set of qualifying salaries.
		qualifying := make(map[int]bool)
		for _, s := range distinctSalaries {
			qualifying[s] = true
		}

		// Gather employees whose salary is in the qualifying set.
		// Sort by salary desc, then name asc for stable output.
		var matched []Employee
		for _, e := range emps {
			if qualifying[e.Salary] {
				matched = append(matched, e)
			}
		}
		sort.Slice(matched, func(i, j int) bool {
			if matched[i].Salary != matched[j].Salary {
				return matched[i].Salary > matched[j].Salary
			}
			return matched[i].Name < matched[j].Name
		})

		for _, e := range matched {
			results = append(results, Result{
				Department: deptName[deptID],
				Employee:   e.Name,
				Salary:     e.Salary,
			})
		}
	}

	// Sort by department name for deterministic output.
	sort.Slice(results, func(i, j int) bool {
		if results[i].Department != results[j].Department {
			return results[i].Department < results[j].Department
		}
		if results[i].Salary != results[j].Salary {
			return results[i].Salary > results[j].Salary
		}
		return results[i].Employee < results[j].Employee
	})

	return results
}

// --- Tests -------------------------------------------------------------------

func main() {
	fmt.Println("=== 0185 Department Top Three Salaries ===")

	departments := []Department{
		{ID: 1, Name: "IT"},
		{ID: 2, Name: "Sales"},
	}

	employees := []Employee{
		{ID: 1, Name: "Joe", Salary: 85000, DepartmentID: 1},
		{ID: 2, Name: "Henry", Salary: 80000, DepartmentID: 2},
		{ID: 3, Name: "Sam", Salary: 60000, DepartmentID: 2},
		{ID: 4, Name: "Max", Salary: 90000, DepartmentID: 1},
		{ID: 5, Name: "Janet", Salary: 69000, DepartmentID: 1},
		{ID: 6, Name: "Randy", Salary: 85000, DepartmentID: 1},
		{ID: 7, Name: "Will", Salary: 70000, DepartmentID: 1},
	}

	fmt.Println("Employees:")
	for _, e := range employees {
		fmt.Printf("  %s (Dept %d, Salary %d)\n", e.Name, e.DepartmentID, e.Salary)
	}

	fmt.Println("\nTop 3 Salaries Per Department:")
	results := departmentTopThreeSalaries(employees, departments)
	for _, r := range results {
		fmt.Printf("  %s | %s | %d\n", r.Department, r.Employee, r.Salary)
	}

	// --- Edge cases ---
	fmt.Println("\n--- Edge Cases ---")

	// Single employee department.
	deps2 := []Department{
		{ID: 10, Name: "Engineering"},
		{ID: 20, Name: "HR"},
	}
	emps2 := []Employee{
		{ID: 1, Name: "Alice", Salary: 100000, DepartmentID: 10},
		{ID: 2, Name: "Bob", Salary: 50000, DepartmentID: 20},
		{ID: 3, Name: "Carol", Salary: 45000, DepartmentID: 20},
	}
	fmt.Println("Test - single employee dept:")
	for _, r := range departmentTopThreeSalaries(emps2, deps2) {
		fmt.Printf("  %s | %s | %d\n", r.Department, r.Employee, r.Salary)
	}

	// All same salary (ties).
	deps3 := []Department{{ID: 1, Name: "Support"}}
	emps3 := []Employee{
		{ID: 10, Name: "A", Salary: 50000, DepartmentID: 1},
		{ID: 20, Name: "B", Salary: 50000, DepartmentID: 1},
		{ID: 30, Name: "C", Salary: 50000, DepartmentID: 1},
		{ID: 40, Name: "D", Salary: 50000, DepartmentID: 1},
	}
	fmt.Println("Test - all same salary (all qualify):")
	for _, r := range departmentTopThreeSalaries(emps3, deps3) {
		fmt.Printf("  %s | %s | %d\n", r.Department, r.Employee, r.Salary)
	}

	// Empty departments.
	fmt.Println("Test - empty employees:", len(departmentTopThreeSalaries(nil, departments)))
}
```

## 0188 — Best Time To Buy And Sell Stock Iv

```go
package main

// LeetCode #188: Best Time to Buy and Sell Stock IV
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/
// Difficulty: Hard

import (
	"fmt"
)

func maxProfit(k int, prices []int) int {
	n := len(prices)
	if n == 0 || k == 0 {
		return 0
	}

	// If k >= n/2, we can do unlimited transactions (greedy)
	if k >= n/2 {
		profit := 0
		for i := 1; i < n; i++ {
			if prices[i] > prices[i-1] {
				profit += prices[i] - prices[i-1]
			}
		}
		return profit
	}

	// dp[i][j] = max profit with at most i transactions up to day j
	dp := make([][]int, k+1)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 1; i <= k; i++ {
		maxDiff := -prices[0] // max(dp[i-1][t] - prices[t]) for t < j
		for j := 1; j < n; j++ {
			dp[i][j] = max(dp[i][j-1], prices[j]+maxDiff)
			maxDiff = max(maxDiff, dp[i-1][j]-prices[j])
		}
	}

	return dp[k][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	k := 2
	prices := []int{3, 2, 6, 5, 0, 3}
	result := maxProfit(k, prices)
	expected := 7

	fmt.Printf("maxProfit(%d, %v) = %d\n", k, prices, result)
	if result == expected {
		fmt.Println("PASS")
	} else {
		fmt.Printf("FAIL: expected %d\n", expected)
	}
}
```

## 0212 — Word Search Ii

```go
package main

// LeetCode #212: Word Search II
// https://leetcode.com/problems/word-search-ii/
// Difficulty: Hard

import "fmt"

type TrieNode struct {
	children [26]*TrieNode
	word     string
}

func findWords(board [][]byte, words []string) []string {
	root := &TrieNode{}
	for _, w := range words {
		node := root
		for _, ch := range w {
			idx := ch - 'a'
			if node.children[idx] == nil {
				node.children[idx] = &TrieNode{}
			}
			node = node.children[idx]
		}
		node.word = w
	}

	result := make([]string, 0)
	dirs := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	var dfs func(r, c int, node *TrieNode)
	dfs = func(r, c int, node *TrieNode) {
		ch := board[r][c]
		if ch == '#' {
			return
		}
		idx := ch - 'a'
		child := node.children[idx]
		if child == nil {
			return
		}
		if child.word != "" {
			result = append(result, child.word)
			child.word = ""
		}

		board[r][c] = '#'
		for _, d := range dirs {
			nr, nc := r+d[0], c+d[1]
			if nr >= 0 && nr < len(board) && nc >= 0 && nc < len(board[0]) {
				dfs(nr, nc, child)
			}
		}
		board[r][c] = ch
	}

	for i := range board {
		for j := range board[0] {
			dfs(i, j, root)
		}
	}
	return result
}

func main() {
	board := [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}
	words := []string{"oath", "pea", "eat", "rain"}
	fmt.Println(findWords(board, words))
}
```

## 0214 — Shortest Palindrome

```go
package main

// LeetCode #214: Shortest Palindrome
// https://leetcode.com/problems/shortest-palindrome/
// Difficulty: Hard

import "fmt"

func shortestPalindrome(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}

	rev := make([]byte, n)
	for i := 0; i < n; i++ {
		rev[i] = s[n-1-i]
	}

	combined := s + "#" + string(rev)
	lps := make([]int, len(combined))

	for i := 1; i < len(combined); i++ {
		j := lps[i-1]
		for j > 0 && combined[i] != combined[j] {
			j = lps[j-1]
		}
		if combined[i] == combined[j] {
			j++
		}
		lps[i] = j
	}

	palLen := lps[len(lps)-1]
	suffix := rev[:n-palLen]
	return string(suffix) + s
}

func main() {
	fmt.Println(shortestPalindrome("aacecaaa"))
	fmt.Println(shortestPalindrome("abcd"))
}
```

## 0218 — The Skyline Problem

```go
package main

// LeetCode #218: The Skyline Problem
// https://leetcode.com/problems/the-skyline-problem/
// Difficulty: Hard

import (
	"container/heap"
	"fmt"
	"sort"
)

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func getSkyline(buildings [][]int) [][]int {
	n := len(buildings)
	points := make([][2]int, 0, 2*n)
	for _, b := range buildings {
		points = append(points, [2]int{b[0], -b[2]})
		points = append(points, [2]int{b[1], b[2]})
	}

	sort.Slice(points, func(i, j int) bool {
		if points[i][0] != points[j][0] {
			return points[i][0] < points[j][0]
		}
		return points[i][1] < points[j][1]
	})

	result := make([][]int, 0)
	h := &MaxHeap{}
	heap.Push(h, 0)
	prev := 0

	for _, p := range points {
		x, y := p[0], p[1]
		if y < 0 {
			heap.Push(h, -y)
		} else {
			toRemove := make([]int, 0)
			temp := &MaxHeap{}
			for h.Len() > 0 {
				top := heap.Pop(h).(int)
				if top == y {
					break
				}
				toRemove = append(toRemove, top)
			}
			for _, v := range toRemove {
				heap.Push(h, v)
			}
			for temp.Len() > 0 {
				heap.Push(h, heap.Pop(temp).(int))
			}
		}

		// rebuild heap to remove stale heights
		cleaned := &MaxHeap{}
		for h.Len() > 0 {
			top := heap.Pop(h).(int)
			if top != y {
				heap.Push(cleaned, top)
			} else {
				break
			}
		}
		for cleaned.Len() > 0 {
			heap.Push(h, heap.Pop(cleaned).(int))
		}

		if h.Len() > 0 {
			cur := (*h)[0]
			if cur != prev {
				result = append(result, []int{x, cur})
				prev = cur
			}
		}
	}

	return result
}

// Alternative approach using lazy deletion (priority queue)
func getSkyline2(buildings [][]int) [][]int {
	n := len(buildings)
	points := make([][2]int, 0, 2*n)
	for _, b := range buildings {
		points = append(points, [2]int{b[0], -b[2]})
		points = append(points, [2]int{b[1], b[2]})
	}

	sort.Slice(points, func(i, j int) bool {
		if points[i][0] != points[j][0] {
			return points[i][0] < points[j][0]
		}
		return points[i][1] < points[j][1]
	})

	// simpler: use a map-based multi-set for heights
	heights := make(map[int]int)
	heights[0] = 1
	prev := 0
	result := make([][]int, 0)
	// max queue using slice
	maxHeight := func() int {
		max := 0
		for h := range heights {
			if h > max {
				max = h
			}
		}
		return max
	}

	for _, p := range points {
		x, y := p[0], p[1]
		if y < 0 {
			heights[-y]++
		} else {
			heights[y]--
			if heights[y] == 0 {
				delete(heights, y)
			}
		}
		cur := maxHeight()
		if cur != prev {
			result = append(result, []int{x, cur})
			prev = cur
		}
	}
	return result
}

func main() {
	buildings := [][]int{{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8}}
	fmt.Println(getSkyline2(buildings))
}
```

## 0220 — Contains Duplicate Iii

```go
package main

// LeetCode #220: Contains Duplicate III
// https://leetcode.com/problems/contains-duplicate-iii/
// Difficulty: Hard

import "fmt"

func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
	if valueDiff < 0 || indexDiff <= 0 {
		return false
	}

	buckets := make(map[int]int)

	for i, num := range nums {
		bucketID := num / (valueDiff + 1)
		if num < 0 {
			bucketID--
		}

		if _, exists := buckets[bucketID]; exists {
			return true
		}
		if val, exists := buckets[bucketID-1]; exists && num-val <= valueDiff {
			return true
		}
		if val, exists := buckets[bucketID+1]; exists && val-num <= valueDiff {
			return true
		}

		buckets[bucketID] = num

		if i >= indexDiff {
			oldNum := nums[i-indexDiff]
			oldBucket := oldNum / (valueDiff + 1)
			if oldNum < 0 {
				oldBucket--
			}
			delete(buckets, oldBucket)
		}
	}

	return false
}

func main() {
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 2, 3, 1}, 3, 0))
}
```

## 0224 — Basic Calculator

```go
package main

// LeetCode #224: Basic Calculator
// https://leetcode.com/problems/basic-calculator/
// Difficulty: Hard

import "fmt"

func calculate(s string) int {
	stack := make([]int, 0)
	result := 0
	sign := 1
	num := 0

	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch >= '0' && ch <= '9' {
			num = 0
			for i < len(s) && s[i] >= '0' && s[i] <= '9' {
				num = num*10 + int(s[i]-'0')
				i++
			}
			result += sign * num
			i--
		} else if ch == '+' {
			sign = 1
		} else if ch == '-' {
			sign = -1
		} else if ch == '(' {
			stack = append(stack, result, sign)
			result = 0
			sign = 1
		} else if ch == ')' {
			result = stack[len(stack)-2] + stack[len(stack)-1]*result
			stack = stack[:len(stack)-2]
		}
	}

	return result
}

func main() {
	fmt.Println(calculate("1 + 1"))
	fmt.Println(calculate("(1+(4+5+2)-3)+(6+8)"))
}
```

## 0233 — Number Of Digit One

```go
package main

// LeetCode #233: Number of Digit One
// https://leetcode.com/problems/number-of-digit-one/
// Difficulty: Hard

import "fmt"

func countDigitOne(n int) int {
	count := 0
	if n <= 0 {
		return 0
	}

	factor := 1
	for factor <= n {
		lower := n % factor
		cur := (n / factor) % 10
		higher := n / (factor * 10)

		switch cur {
		case 0:
			count += higher * factor
		case 1:
			count += higher*factor + lower + 1
		default:
			count += (higher + 1) * factor
		}

		// Check overflow
		if factor > n/10 {
			break
		}
		factor *= 10
	}

	return count
}

func main() {
	fmt.Println(countDigitOne(13))
	fmt.Println(countDigitOne(0))
}
```

## 0239 — Sliding Window Maximum

```go
package main

// LeetCode #239: Sliding Window Maximum
// https://leetcode.com/problems/sliding-window-maximum/
// Difficulty: Hard

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return nil
	}

	deque := make([]int, 0) // stores indices
	result := make([]int, 0, len(nums)-k+1)

	for i, num := range nums {
		// Remove indices outside the window (from front)
		if len(deque) > 0 && deque[0] < i-k+1 {
			deque = deque[1:]
		}

		// Remove from back while current num is larger (maintain decreasing order)
		for len(deque) > 0 && nums[deque[len(deque)-1]] < num {
			deque = deque[:len(deque)-1]
		}

		deque = append(deque, i)

		// First valid window starts at index k-1
		if i >= k-1 {
			result = append(result, nums[deque[0]])
		}
	}

	return result
}

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}
```

## 0248 — Strobogrammatic Number Iii

```go
package main

// LeetCode #248: Strobogrammatic Number III
// https://leetcode.com/problems/strobogrammatic-number-iii/
// Difficulty: Hard [Paid]

import "fmt"

var pairs = [][2]byte{
	{'0', '0'},
	{'1', '1'},
	{'6', '9'},
	{'8', '8'},
	{'9', '6'},
}

func strobogrammaticInRange(low string, high string) int {
	count := 0
	lowNum := len(low)
	highNum := len(high)

	for length := lowNum; length <= highNum; length++ {
		var dfs func(cur []byte, left, right int)
		dfs = func(cur []byte, left, right int) {
			if left > right {
				s := string(cur)
				if (len(s) == len(low) && s < low) || (len(s) == len(high) && s > high) {
					return
				}
				count++
				return
			}

			for _, p := range pairs {
				cur[left] = p[0]
				cur[right] = p[1]
				if len(cur) > 1 && cur[0] == '0' {
					continue
				}
				if left == right && p[0] != p[1] {
					continue
				}
				dfs(cur, left+1, right-1)
			}
		}

		cur := make([]byte, length)
		dfs(cur, 0, length-1)
	}

	// fix: check edge cases properly — compare as strings with same length
	result := 0
	if len(low) == len(high) {
		for _, s := range generate(len(low)) {
			if s >= low && s <= high {
				result++
			}
		}
		return result
	}

	for _, s := range generate(len(low)) {
		if s >= low {
			result++
		}
	}
	for _, s := range generate(len(high)) {
		if s <= high {
			result++
		}
	}
	for l := len(low) + 1; l < len(high); l++ {
		result += countStrobogrammatic(l)
	}
	return result
}

func generate(n int) []string {
	return helper(n, n)
}

func helper(n, m int) []string {
	if n == 0 {
		return []string{""}
	}
	if n == 1 {
		return []string{"0", "1", "8"}
	}

	inner := helper(n-2, m)
	result := make([]string, 0)
	for _, s := range inner {
		if n != m {
			result = append(result, "0"+s+"0")
		}
		result = append(result, "1"+s+"1")
		result = append(result, "6"+s+"9")
		result = append(result, "8"+s+"8")
		result = append(result, "9"+s+"6")
	}
	return result
}

func countStrobogrammatic(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 3 // 0, 1, 8
	}
	valid := 1
	if n%2 == 1 {
		valid = 3
	}
	count := 4 // 1, 6, 8, 9 (first digit)
	for i := 0; i < n/2-1; i++ {
		count *= 5
	}
	count *= valid
	// subtract 0-padded ones
	if n%2 == 0 {
		return count
	}
	// odd length, middle can be 0,1,8
	return count
}

func main() {
	fmt.Println(strobogrammaticInRange("50", "100"))
}
```

## 0262 — Trips And Users

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #262: Trips and Users
// https://leetcode.com/problems/trips-and-users/
// Difficulty: Hard
//
// Find the cancellation rate of unbanned users (both client and driver must not be banned)
// for each day between "2013-10-01" and "2013-10-03".
// Cancellation rate = number of cancelled trips / total trips (by unbanned users).

// User represents a row in the Users table.
type User struct {
	UsersID int
	Banned  string // "Yes" or "No"
	Role    string // "client", "driver", "partner"
}

// Trip represents a row in the Trips table.
type Trip struct {
	ID        int
	ClientID  int
	DriverID  int
	CityID    int
	Status    string // "completed", "cancelled_by_driver", "cancelled_by_client"
	RequestAt string // date "YYYY-MM-DD"
}

// DailyRate holds one output row.
type DailyRate struct {
	Day    string
	Rate   float64 // cancellation rate, rounded to 2 decimal places
}

// RoundTo2 rounds f to 2 decimal places.
func roundTo2(f float64) float64 {
	return float64(int(f*100+0.5)) / 100.0
}

// tripsAndUsers computes the daily cancellation rate for unbanned users.
// Time: O(T + U + D log D) where T=#trips, U=#users, D=#distinct dates
func tripsAndUsers(trips []Trip, users []User) []DailyRate {
	// Build banned user set.
	banned := make(map[int]bool)
	for _, u := range users {
		if u.Banned == "Yes" {
			banned[u.UsersID] = true
		}
	}

	// Group trips by date, only including trips where both client and driver are unbanned.
	type dateStats struct {
		total       int
		cancelled   int
	}
	byDay := make(map[string]*dateStats)

	for _, t := range trips {
		if banned[t.ClientID] || banned[t.DriverID] {
			continue
		}
		if _, ok := byDay[t.RequestAt]; !ok {
			byDay[t.RequestAt] = &dateStats{}
		}
		byDay[t.RequestAt].total++
		if t.Status != "completed" {
			byDay[t.RequestAt].cancelled++
		}
	}

	// Collect dates in sorted order.
	var dates []string
	for d := range byDay {
		dates = append(dates, d)
	}
	sort.Strings(dates)

	var result []DailyRate
	for _, d := range dates {
		s := byDay[d]
		rate := 0.0
		if s.total > 0 {
			rate = roundTo2(float64(s.cancelled) / float64(s.total))
		}
		result = append(result, DailyRate{Day: d, Rate: rate})
	}
	return result
}

// --- Tests -------------------------------------------------------------------

func main() {
	fmt.Println("=== 0262 Trips and Users ===")

	users := []User{
		{UsersID: 1, Banned: "No", Role: "client"},
		{UsersID: 2, Banned: "Yes", Role: "client"},
		{UsersID: 3, Banned: "No", Role: "client"},
		{UsersID: 4, Banned: "No", Role: "client"},
		{UsersID: 10, Banned: "No", Role: "driver"},
		{UsersID: 11, Banned: "No", Role: "driver"},
		{UsersID: 12, Banned: "No", Role: "driver"},
		{UsersID: 13, Banned: "No", Role: "driver"},
	}

	trips := []Trip{
		{ID: 1, ClientID: 1, DriverID: 10, CityID: 1, Status: "completed", RequestAt: "2013-10-01"},
		{ID: 2, ClientID: 2, DriverID: 11, CityID: 1, Status: "cancelled_by_driver", RequestAt: "2013-10-01"},
		{ID: 3, ClientID: 3, DriverID: 12, CityID: 6, Status: "completed", RequestAt: "2013-10-01"},
		{ID: 4, ClientID: 4, DriverID: 13, CityID: 6, Status: "cancelled_by_client", RequestAt: "2013-10-01"},
		{ID: 5, ClientID: 1, DriverID: 10, CityID: 1, Status: "completed", RequestAt: "2013-10-02"},
		{ID: 6, ClientID: 2, DriverID: 11, CityID: 6, Status: "completed", RequestAt: "2013-10-02"},
		{ID: 7, ClientID: 3, DriverID: 12, CityID: 6, Status: "completed", RequestAt: "2013-10-02"},
		{ID: 8, ClientID: 2, DriverID: 12, CityID: 12, Status: "completed", RequestAt: "2013-10-03"},
		{ID: 9, ClientID: 3, DriverID: 10, CityID: 1, Status: "completed", RequestAt: "2013-10-03"},
		{ID: 10, ClientID: 4, DriverID: 13, CityID: 1, Status: "cancelled_by_driver", RequestAt: "2013-10-03"},
	}

	fmt.Println("Trips:", len(trips), "Users:", len(users))
	results := tripsAndUsers(trips, users)
	fmt.Println("\nDaily Cancellation Rate (unbanned users only):")
	for _, r := range results {
		fmt.Printf("  %s  rate=%.2f\n", r.Day, r.Rate)
	}

	// --- Edge cases ---
	fmt.Println("\n--- Edge Cases ---")

	// No trips.
	empty := tripsAndUsers(nil, users)
	fmt.Println("Empty trips:", len(empty))

	// All trips cancelled.
	trips2 := []Trip{
		{ID: 1, ClientID: 1, DriverID: 10, CityID: 1, Status: "cancelled_by_client", RequestAt: "2013-10-01"},
		{ID: 2, ClientID: 3, DriverID: 11, CityID: 1, Status: "cancelled_by_driver", RequestAt: "2013-10-01"},
	}
	r2 := tripsAndUsers(trips2, users)
	for _, r := range r2 {
		fmt.Printf("  All cancelled: %s rate=%.2f (expected 1.00)\n", r.Day, r.Rate)
	}

	// All banned users (trips filtered out).
	usersBanned := []User{
		{UsersID: 1, Banned: "Yes", Role: "client"},
		{UsersID: 10, Banned: "Yes", Role: "driver"},
	}
	trips3 := []Trip{
		{ID: 1, ClientID: 1, DriverID: 10, CityID: 1, Status: "completed", RequestAt: "2013-10-01"},
	}
	r3 := tripsAndUsers(trips3, usersBanned)
	fmt.Println("All banned users:", len(r3), "(expected 0)")
}
```

## 0265 — Paint House Ii

```go
package main

// LeetCode #265: Paint House II
// https://leetcode.com/problems/paint-house-ii/
// Difficulty: Hard [Paid]

import (
	"fmt"
	"math"
)

func minCostII(costs [][]int) int {
	if len(costs) == 0 {
		return 0
	}

	n := len(costs)
	k := len(costs[0])

	// Find min and second min for the first house
	prevMin1, prevMin2 := math.MaxInt32, math.MaxInt32
	prevMinIdx := -1

	for j := 0; j < k; j++ {
		cost := costs[0][j]
		if cost < prevMin1 {
			prevMin2 = prevMin1
			prevMin1 = cost
			prevMinIdx = j
		} else if cost < prevMin2 {
			prevMin2 = cost
		}
	}

	for i := 1; i < n; i++ {
		curMin1, curMin2 := math.MaxInt32, math.MaxInt32
		curMinIdx := -1

		for j := 0; j < k; j++ {
			var cost int
			if j == prevMinIdx {
				cost = costs[i][j] + prevMin2
			} else {
				cost = costs[i][j] + prevMin1
			}

			if cost < curMin1 {
				curMin2 = curMin1
				curMin1 = cost
				curMinIdx = j
			} else if cost < curMin2 {
				curMin2 = cost
			}
		}

		prevMin1, prevMin2 = curMin1, curMin2
		prevMinIdx = curMinIdx
	}

	return prevMin1
}

func main() {
	fmt.Println(minCostII([][]int{{1, 5, 3}, {2, 9, 4}}))
}
```

## 0269 — Alien Dictionary

```go
package main

// LeetCode #269: Alien Dictionary
// https://leetcode.com/problems/alien-dictionary/
// Difficulty: Hard [Paid]

import (
	"fmt"
)

func alienOrder(words []string) string {
	graph := make(map[byte][]byte)
	inDegree := make(map[byte]int)

	// Initialize all characters
	for _, w := range words {
		for i := 0; i < len(w); i++ {
			ch := w[i]
			if _, exists := graph[ch]; !exists {
				graph[ch] = make([]byte, 0)
				inDegree[ch] = 0
			}
		}
	}

	// Build graph
	for i := 0; i < len(words)-1; i++ {
		w1, w2 := words[i], words[i+1]
		minLen := len(w1)
		if len(w2) < minLen {
			minLen = len(w2)
		}

		// Check for invalid prefix case: w1 is longer and w1 starts with w2
		if len(w1) > len(w2) && w1[:len(w2)] == w2 {
			return ""
		}

		for j := 0; j < minLen; j++ {
			c1, c2 := w1[j], w2[j]
			if c1 != c2 {
				graph[c1] = append(graph[c1], c2)
				inDegree[c2]++
				break
			}
		}
	}

	// Topological sort (BFS / Kahn's algorithm)
	queue := make([]byte, 0)
	for ch := range graph {
		if inDegree[ch] == 0 {
			queue = append(queue, ch)
		}
	}

	result := make([]byte, 0, len(graph))
	for len(queue) > 0 {
		ch := queue[0]
		queue = queue[1:]
		result = append(result, ch)

		for _, neighbor := range graph[ch] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	if len(result) != len(graph) {
		return ""
	}

	return string(result)
}

func main() {
	fmt.Println(alienOrder([]string{"wrt", "wrf", "er", "ett", "rftt"}))
}
```

## 0272 — Closest Binary Search Tree Value Ii

```go
package main

// LeetCode #272: Closest Binary Search Tree Value II
// https://leetcode.com/problems/closest-binary-search-tree-value-ii/
// Difficulty: Hard [Paid]
//
// Approach: Inorder traversal + two-pointer sliding window.
//  1. Inorder traversal of BST gives sorted values.
//  2. Use two-pointer to maintain a window of k elements closest to target.
//  3. Expand from the window edges, always removing the farther element.

import (
	"fmt"
	"math"
)

// TreeNode definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Example: root=[4,2,5,1,3], target=3.714, k=2 -> [3,4]
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{Val: 5},
	}

	result := closestKValues(root, 3.714, 2)
	fmt.Println("Closest K Values:", result)

	// Test: k=3
	result2 := closestKValues(root, 3.714, 3)
	fmt.Println("Closest K Values (k=3):", result2)
}

// closestKValues returns k closest values to target in BST root.
func closestKValues(root *TreeNode, target float64, k int) []int {
	// Collect inorder traversal values.
	var values []int
	inorder(root, &values)

	// Two-pointer: find the window of size k.
	// First, find the starting point using binary search to locate first element >= target.
	idx := lowerBound(values, target)

	// Expand window: left goes backwards, right goes forwards.
	left := idx - 1
	right := idx

	for right-left-1 < k {
		if left < 0 {
			// Only right side available.
			right++
		} else if right >= len(values) {
			// Only left side available.
			left--
		} else {
			// Compare distances.
			distLeft := target - float64(values[left])
			distRight := float64(values[right]) - target
			if distLeft < distRight {
				left--
			} else {
				right++
			}
		}
	}

	// Extract window (left+1 ... right-1).
	return values[left+1 : right]
}

// inorder performs inorder traversal of BST.
func inorder(node *TreeNode, values *[]int) {
	if node == nil {
		return
	}
	inorder(node.Left, values)
	*values = append(*values, node.Val)
	inorder(node.Right, values)
}

// lowerBound finds the first index where values[i] >= target.
func lowerBound(values []int, target float64) int {
	lo, hi := 0, len(values)
	for lo < hi {
		mid := lo + (hi-lo)/2
		if float64(values[mid]) < target {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

// --- Alternative: O(n) without extra space using predecessor/successor ---

// closestKValuesO1Space uses inorder + reverse inorder to fill k closest.
// Not implemented here for brevity; the two-pointer approach above is standard.

// --- Older solution stub compatibility ---
func ClosestBinarySearchTreeValueIi() any {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{Val: 5},
	}
	return closestKValues(root, 3.714, 2)
}

// --- Testing utilities ---

// Helper to verify floating point comparison.
func equalFloat(a, b float64) bool {
	const eps = 1e-9
	return math.Abs(a-b) < eps
}
```

## 0273 — Integer To English Words

```go
package main

// LeetCode #273: Integer to English Words
// https://leetcode.com/problems/integer-to-english-words/
// Difficulty: Hard
//
// Approach: Groups of Three (Chunking).
//  1. Handle zero separately.
//  2. Process number in groups of 3 digits: billions, millions, thousands, ones.
//  3. For each group, convert the 3-digit number to words using helper functions.
//  4. Append the scale word (Billion, Million, Thousand) for non-zero groups.

import "fmt"

func main() {
	// Example 1: 123 -> "One Hundred Twenty Three"
	fmt.Println("123 ->", numberToWords(123))

	// Example 2: 12345 -> "Twelve Thousand Three Hundred Forty Five"
	fmt.Println("12345 ->", numberToWords(12345))

	// Example 3: 1234567 -> "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"
	fmt.Println("1234567 ->", numberToWords(1234567))

	// Edge cases
	fmt.Println("0 ->", numberToWords(0))
	fmt.Println("1000 ->", numberToWords(1000))
	fmt.Println("1000000 ->", numberToWords(1000000))
	fmt.Println("1000000000 ->", numberToWords(1000000000))
	fmt.Println("2147483647 ->", numberToWords(2147483647))
}

var (
	belowTwenty = []string{
		"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine",
		"Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen",
		"Seventeen", "Eighteen", "Nineteen",
	}
	tens = []string{
		"", "", "Twenty", "Thirty", "Forty", "Fifty",
		"Sixty", "Seventy", "Eighty", "Ninety",
	}
	thousands = []string{"", "Thousand", "Million", "Billion"}
)

// numberToWords converts a non-negative integer to English words.
func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}

	result := ""
	unitIndex := 0

	for num > 0 {
		chunk := num % 1000
		if chunk > 0 {
			chunkWords := convertChunk(chunk)
			if result == "" {
				result = chunkWords + " " + thousands[unitIndex]
			} else {
				result = chunkWords + " " + thousands[unitIndex] + " " + result
			}
		}
		num /= 1000
		unitIndex++
	}

	// Clean up extra spaces.
	// We trim spaces at the end instead of complex logic.
	return trimSpace(result)
}

// convertChunk converts a 3-digit number (0-999) to English words.
// Note: chunk is guaranteed > 0 when called.
func convertChunk(num int) string {
	var result string

	hundreds := num / 100
	remainder := num % 100

	if hundreds > 0 {
		result = belowTwenty[hundreds] + " Hundred"
	}

	if remainder > 0 {
		if result != "" {
			result += " "
		}
		if remainder < 20 {
			result += belowTwenty[remainder]
		} else {
			ten := remainder / 10
			one := remainder % 10
			result += tens[ten]
			if one > 0 {
				result += " " + belowTwenty[one]
			}
		}
	}

	return result
}

// trimSpace removes trailing spaces from a string.
func trimSpace(s string) string {
	if len(s) == 0 {
		return s
	}
	// Find the last non-space character.
	end := len(s) - 1
	for end >= 0 && s[end] == ' ' {
		end--
	}
	return s[:end+1]
}

// Stub compatibility.
func IntegerToEnglishWords() any {
	return numberToWords(123)
}
```

## 0282 — Expression Add Operators

```go
package main

// LeetCode #282: Expression Add Operators
// https://leetcode.com/problems/expression-add-operators/
// Difficulty: Hard
//
// Approach: Backtracking.
//  1. Try all possible splits of the string into operands.
//  2. At each step, try '+', '-', '*' operators.
//  3. For '*', we need to track the last operand to handle precedence:
//     current result = (result - lastOperand) + (lastOperand * currentNum)
//  4. Avoid numbers with leading zeros (e.g., "05" is invalid).

import (
	"fmt"
	"strconv"
)

func main() {
	// Example 1: "123", 6 -> ["1+2+3","1*2*3"]
	fmt.Println("\"123\", 6 ->", addOperators("123", 6))

	// Example 2: "232", 8 -> ["2*3+2","2+3*2"]
	fmt.Println("\"232\", 8 ->", addOperators("232", 8))

	// Example 3: "3456237490", 9191 -> []
	fmt.Println("\"3456237490\", 9191 ->", addOperators("3456237490", 9191))

	// Edge: "105", 5 -> ["1*0+5","10-5"]
	fmt.Println("\"105\", 5 ->", addOperators("105", 5))

	// Edge: "00", 0 -> ["0+0","0-0","0*0"]
	fmt.Println("\"00\", 0 ->", addOperators("00", 0))
}

// addOperators returns all possible expressions that evaluate to target.
func addOperators(num string, target int) []string {
	var result []string
	if len(num) == 0 {
		return result
	}
	backtrack(num, target, 0, 0, 0, "", &result)
	return result
}

// backtrack explores all possible expression constructions.
//   - num: remaining string to process
//   - target: target value
//   - index: current position in num
//   - currentValue: value of the expression built so far
//   - lastOperand: last operand added to the expression (for multiplication precedence)
//   - expression: the expression string built so far
//   - result: collects valid expressions
func backtrack(num string, target int, index int, currentValue int, lastOperand int, expression string, result *[]string) {
	if index == len(num) {
		if currentValue == target {
			*result = append(*result, expression)
		}
		return
	}

	for i := index; i < len(num); i++ {
		// Avoid numbers with leading zeros.
		if i > index && num[index] == '0' {
			break
		}

		// Parse current number from num[index:i+1].
		currentNum, _ := strconv.Atoi(num[index : i+1])

		if index == 0 {
			// First operand: no operator needed.
			backtrack(num, target, i+1, currentNum, currentNum, strconv.Itoa(currentNum), result)
		} else {
			// Try addition.
			backtrack(num, target, i+1, currentValue+currentNum, currentNum, expression+"+"+strconv.Itoa(currentNum), result)

			// Try subtraction.
			backtrack(num, target, i+1, currentValue-currentNum, -currentNum, expression+"-"+strconv.Itoa(currentNum), result)

			// Try multiplication: undo last operation, then multiply.
			// newValue = currentValue - lastOperand + (lastOperand * currentNum)
			backtrack(num, target, i+1, currentValue-lastOperand+lastOperand*currentNum, lastOperand*currentNum, expression+"*"+strconv.Itoa(currentNum), result)
		}
	}
}

// Stub compatibility.
func ExpressionAddOperators() any {
	return addOperators("123", 6)
}
```

## 0295 — Find Median From Data Stream

```go
package main

// LeetCode #295: Find Median from Data Stream
// https://leetcode.com/problems/find-median-from-data-stream/
// Difficulty: Hard
//
// Approach: Two Heaps.
//   - Max-heap (lo) stores the smaller half of numbers.
//   - Min-heap (hi) stores the larger half of numbers.
//   - Maintain invariant: lo has at most one more element than hi.
//   - Median: if lo.Len() > hi.Len(), median = lo[0]; else median = (lo[0] + hi[0]) / 2.

import (
	"container/heap"
	"fmt"
)

func main() {
	// Example usage.
	mf := Constructor()
	mf.AddNum(1)
	mf.AddNum(2)
	fmt.Println("Median after [1,2]:", mf.FindMedian()) // 1.5
	mf.AddNum(3)
	fmt.Println("Median after [1,2,3]:", mf.FindMedian()) // 2.0

	// Larger test.
	mf2 := Constructor()
	for _, v := range []int{5, 2, 8, 1, 9, 3, 7} {
		mf2.AddNum(v)
	}
	fmt.Println("Median of [1,2,3,5,7,8,9]:", mf2.FindMedian()) // 5.0
}

// --- Max Heap (for the smaller half) ---

// MaxHeap implements heap.Interface (stores negative values to simulate max-heap).
type MaxHeap []int

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool   { return h[i] > h[j] } // Reverse for max-heap
func (h MaxHeap) Swap(i, j int)        { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// --- Min Heap (for the larger half) ---

// MinHeap implements heap.Interface.
type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool   { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)        { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// --- MedianFinder ---

// MedianFinder maintains a data stream median.
type MedianFinder struct {
	lo *MaxHeap // max-heap for smaller half
	hi *MinHeap // min-heap for larger half
}

// Constructor initializes the MedianFinder.
func Constructor() MedianFinder {
	lo := &MaxHeap{}
	hi := &MinHeap{}
	heap.Init(lo)
	heap.Init(hi)
	return MedianFinder{lo: lo, hi: hi}
}

// AddNum adds a number to the data stream.
func (mf *MedianFinder) AddNum(num int) {
	// Step 1: Add to lo (max-heap).
	heap.Push(mf.lo, num)

	// Step 2: Move the largest from lo to hi to maintain ordering.
	heap.Push(mf.hi, heap.Pop(mf.lo).(int))

	// Step 3: Balance: lo size should be >= hi size.
	if mf.lo.Len() < mf.hi.Len() {
		heap.Push(mf.lo, heap.Pop(mf.hi).(int))
	}
}

// FindMedian returns the median of all elements so far.
func (mf *MedianFinder) FindMedian() float64 {
	if mf.lo.Len() > mf.hi.Len() {
		return float64((*mf.lo)[0])
	}
	return float64((*mf.lo)[0]+(*mf.hi)[0]) / 2.0
}

// Stub compatibility.
func FindMedianFromDataStream() any {
	mf := Constructor()
	mf.AddNum(1)
	mf.AddNum(2)
	return mf.FindMedian() // 1.5
}

// Ensure heap.Interface is satisfied.
var _ heap.Interface = (*MaxHeap)(nil)
var _ heap.Interface = (*MinHeap)(nil)
```

## 0296 — Best Meeting Point

```go
package main

// LeetCode #296: Best Meeting Point
// https://leetcode.com/problems/best-meeting-point/
// Difficulty: Hard [Paid]
//
// Approach: Median of coordinates (Manhattan distance).
//   The total Manhattan distance is minimized at the median of row coordinates
//   and median of column coordinates independently.
//   1. Collect all row coordinates and column coordinates where grid[i][j] == 1.
//   2. Sort each list.
//   3. The optimal meeting point is at the median values.
//   4. Sum the absolute differences from the medians.

import (
	"fmt"
	"sort"
)

func main() {
	// Example: [[1,0,0,0,1],[0,0,0,0,0],[0,0,1,0,0]]
	grid := [][]int{
		{1, 0, 0, 0, 1},
		{0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0},
	}
	fmt.Println("Best meeting point distance:", minTotalDistance(grid))

	// Test: [[0,0],[2,0],[1,1]]
	grid2 := [][]int{
		{0, 0},
		{2, 0},
		{1, 1},
	}
	fmt.Println("Best meeting point distance (2):", minTotalDistance(grid2)) // 4

	// Test: [[1,1]]
	grid3 := [][]int{
		{1, 1},
	}
	fmt.Println("Single row with 2 points:", minTotalDistance(grid3))

	// Test: no friends (should return 0).
	grid4 := [][]int{
		{0, 0},
		{0, 0},
	}
	fmt.Println("No friends:", minTotalDistance(grid4))
}

// minTotalDistance returns the minimum total Manhattan distance for a meeting point.
func minTotalDistance(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])

	// Collect row and column coordinates of all 1s.
	var rows, cols []int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				rows = append(rows, i)
				cols = append(cols, j)
			}
		}
	}

	if len(rows) == 0 {
		return 0
	}

	// Sort both lists.
	sort.Ints(rows)
	sort.Ints(cols)

	// Find medians.
	medianRow := rows[len(rows)/2]
	medianCol := cols[len(cols)/2]

	// Sum distances from median.
	total := 0
	for _, r := range rows {
		total += abs(r - medianRow)
	}
	for _, c := range cols {
		total += abs(c - medianCol)
	}

	return total
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Stub compatibility.
func BestMeetingPoint() any {
	grid := [][]int{
		{0, 0},
		{2, 0},
		{1, 1},
	}
	return minTotalDistance(grid)
}
```

## 0297 — Serialize And Deserialize Binary Tree

```go
package main

// LeetCode #297: Serialize and Deserialize Binary Tree
// https://leetcode.com/problems/serialize-and-deserialize-binary-tree/
// Difficulty: Hard
//
// Approach: BFS Level-Order (Iterative).
//   - Serialize: Level-order traversal using a queue. Use "null" for nil nodes.
//   - Trailing nulls are stripped to keep the serialized string concise.
//   - Deserialize: Use a queue to rebuild the tree from the BFS order.
//
// The serialized format is a comma-separated string like "4,2,5,1,3,null,null".

import (
	"fmt"
	"strconv"
	"strings"
)

// TreeNode definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Example: root=[4,2,5,1,3]
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{Val: 5},
	}

	codec := Constructor297()

	// Serialize.
	data := codec.serialize(root)
	fmt.Println("Serialized:", data)

	// Deserialize.
	decoded := codec.deserialize(data)
	fmt.Println("Deserialized root value:", decoded.Val)
	fmt.Println("Serialized again:", codec.serialize(decoded))

	// Test empty tree.
	emptyData := codec.serialize(nil)
	fmt.Println("Empty tree serialize:", emptyData)
	decodedEmpty := codec.deserialize(emptyData)
	fmt.Println("Empty tree deserialize:", decodedEmpty)

	// Test single node.
	single := &TreeNode{Val: 1}
	singleData := codec.serialize(single)
	fmt.Println("Single node serialize:", singleData)
	decodedSingle := codec.deserialize(singleData)
	fmt.Println("Single node deserialize root:", decodedSingle.Val)
}

// Codec handles serialization/deserialization.
type Codec struct{}

// Constructor297 creates a new Codec.
func Constructor297() Codec {
	return Codec{}
}

// serialize serializes a binary tree to a string.
func (c *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	var parts []string
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			parts = append(parts, "null")
		} else {
			parts = append(parts, strconv.Itoa(node.Val))
			queue = append(queue, node.Left, node.Right)
		}
	}

	// Strip trailing "null"s.
	end := len(parts) - 1
	for end >= 0 && parts[end] == "null" {
		end--
	}
	parts = parts[:end+1]

	return strings.Join(parts, ",")
}

// deserialize deserializes a string back to a binary tree.
func (c *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}

	parts := strings.Split(data, ",")
	if len(parts) == 0 || parts[0] == "" {
		return nil
	}

	val, _ := strconv.Atoi(parts[0])
	root := &TreeNode{Val: val}
	queue := []*TreeNode{root}
	index := 1

	for len(queue) > 0 && index < len(parts) {
		node := queue[0]
		queue = queue[1:]

		// Process left child.
		if index < len(parts) && parts[index] != "null" {
			val, _ = strconv.Atoi(parts[index])
			node.Left = &TreeNode{Val: val}
			queue = append(queue, node.Left)
		}
		index++

		// Process right child.
		if index < len(parts) && parts[index] != "null" {
			val, _ = strconv.Atoi(parts[index])
			node.Right = &TreeNode{Val: val}
			queue = append(queue, node.Right)
		}
		index++
	}

	return root
}

// Stub compatibility.
func SerializeAndDeserializeBinaryTree() any {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{Val: 5},
	}
	codec := Constructor297()
	data := codec.serialize(root)
	decoded := codec.deserialize(data)
	return decoded.Val // should be 4
}
```

## 0301 — Remove Invalid Parentheses

```go
package main

// LeetCode #301: Remove Invalid Parentheses
// https://leetcode.com/problems/remove-invalid-parentheses/
// Difficulty: Hard
//
// Approach: BFS.
//  1. BFS over all possible strings generated by removing one parenthesis.
//  2. At each level, check if any string is valid (balanced parentheses).
//  3. Return all valid strings at the first level where valid strings exist.
//
// Complexity:
//   - BFS ensures we find strings with minimum removals.
//   - Use a visited set to avoid duplicates.
//   - Use a "seen" set at each level to avoid processing the same string twice.

import "fmt"

func main() {
	// Example 1: "()())()" -> ["(())()","()()()"]
	fmt.Println("\"()())()\" ->", removeInvalidParentheses("()())()"))

	// Example 2: "(a)())()" -> ["(a())()","(a)()()"]
	fmt.Println("\"(a)())()\" ->", removeInvalidParentheses("(a)())()"))

	// Example 3: ")(" -> [""]
	fmt.Println("\")(\" ->", removeInvalidParentheses(")("))

	// Example 4: "n" -> ["n"]
	fmt.Println("\"n\" ->", removeInvalidParentheses("n"))

	// Example 5: "(r(()()(" -> ["r(()())","(r())()","r()()()","(r)()()"]
	fmt.Println("\"(r(()()(\" ->", removeInvalidParentheses("(r(()()("))
}

// removeInvalidParentheses removes the minimum number of invalid parentheses
// to make the string valid, returning all possible results.
func removeInvalidParentheses(s string) []string {
	var result []string
	if len(s) == 0 {
		result = append(result, "")
		return result
	}

	visited := make(map[string]bool)
	queue := []string{s}
	visited[s] = true
	found := false

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]

			if isValid(curr) {
				result = append(result, curr)
				found = true
			}

			if found {
				continue
			}

			// Generate all strings by removing one parenthesis.
			for j := 0; j < len(curr); j++ {
				if curr[j] != '(' && curr[j] != ')' {
					continue
				}
				next := curr[:j] + curr[j+1:]
				if !visited[next] {
					visited[next] = true
					queue = append(queue, next)
				}
			}
		}
		if found {
			break
		}
	}

	if len(result) == 0 {
		result = append(result, "")
	}

	return result
}

// isValid checks if a string has balanced parentheses (ignoring non-parenthesis chars).
func isValid(s string) bool {
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			count++
		} else if s[i] == ')' {
			count--
			if count < 0 {
				return false
			}
		}
	}
	return count == 0
}

// Stub compatibility.
func RemoveInvalidParentheses() any {
	return removeInvalidParentheses("()())()")
}
```

## 0302 — Smallest Rectangle Enclosing Black Pixels

```go
package main

// LeetCode #302: Smallest Rectangle Enclosing Black Pixels
// https://leetcode.com/problems/smallest-rectangle-enclosing-black-pixels/
// Difficulty: Hard [Paid]
//
// Approach: Binary Search on Boundaries.
//   The image has exactly one black region (all '1's are contiguous).
//   Given a known black pixel (x, y), we can binary search for the
//   topmost, bottommost, leftmost, and rightmost '1's.
//   - Top: smallest row in [0, x] that has a '1' in any column.
//   - Bottom: largest row in [x, m-1] that has a '1' in any column.
//   - Left: smallest col in [0, y] that has a '1' in any row.
//   - Right: largest col in [y, n-1] that has a '1' in any row.
//   Area = (bottom - top + 1) * (right - left + 1).

import "fmt"

func main() {
	// Example:
	// [["0","0","1","0"],
	//  ["0","1","1","0"],
	//  ["0","1","0","0"]]
	// x=0, y=2 -> area=6
	grid := [][]byte{
		{'0', '0', '1', '0'},
		{'0', '1', '1', '0'},
		{'0', '1', '0', '0'},
	}
	fmt.Println("Area:", minArea(grid, 0, 2)) // 6

	// Single pixel.
	grid2 := [][]byte{{'1'}}
	fmt.Println("Area (single):", minArea(grid2, 0, 0)) // 1

	// Single column.
	grid3 := [][]byte{
		{'0', '1'},
		{'0', '1'},
	}
	fmt.Println("Area (col):", minArea(grid3, 0, 1)) // 2

	// All zeros (edge case, shouldn't happen per problem constraints).
	// The problem guarantees at least one black pixel at (x,y).
	grid4 := [][]byte{
		{'0', '0'},
		{'1', '0'},
	}
	fmt.Println("Area (one pixel):", minArea(grid4, 1, 0)) // 1
}

// minArea returns the area of the smallest rectangle that encloses all '1's.
func minArea(image [][]byte, x int, y int) int {
	if len(image) == 0 || len(image[0]) == 0 {
		return 0
	}

	m, n := len(image), len(image[0])

	// Binary search for topmost row with a '1'.
	top := searchTop(image, 0, x, n)
	// Binary search for bottommost row with a '1'.
	bottom := searchBottom(image, x, m-1, n)
	// Binary search for leftmost column with a '1'.
	left := searchLeft(image, 0, y, m)
	// Binary search for rightmost column with a '1'.
	right := searchRight(image, y, n-1, m)

	return (bottom - top + 1) * (right - left + 1)
}

// hasBlackInRow checks if any column in row r has a '1'.
func hasBlackInRow(image [][]byte, row int, n int) bool {
	for c := 0; c < n; c++ {
		if image[row][c] == '1' {
			return true
		}
	}
	return false
}

// hasBlackInColumn checks if any row in column c has a '1'.
func hasBlackInColumn(image [][]byte, col int, m int) bool {
	for r := 0; r < m; r++ {
		if image[r][col] == '1' {
			return true
		}
	}
	return false
}

// searchTop finds the topmost row with a '1' in range [lo, hi].
func searchTop(image [][]byte, lo, hi int, n int) int {
	for lo < hi {
		mid := lo + (hi-lo)/2
		if hasBlackInRow(image, mid, n) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

// searchBottom finds the bottommost row with a '1' in range [lo, hi].
func searchBottom(image [][]byte, lo, hi int, n int) int {
	for lo < hi {
		mid := lo + (hi-lo+1)/2 // ceiling mid
		if hasBlackInRow(image, mid, n) {
			lo = mid
		} else {
			hi = mid - 1
		}
	}
	return lo
}

// searchLeft finds the leftmost column with a '1' in range [lo, hi].
func searchLeft(image [][]byte, lo, hi int, m int) int {
	for lo < hi {
		mid := lo + (hi-lo)/2
		if hasBlackInColumn(image, mid, m) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

// searchRight finds the rightmost column with a '1' in range [lo, hi].
func searchRight(image [][]byte, lo, hi int, m int) int {
	for lo < hi {
		mid := lo + (hi-lo+1)/2 // ceiling mid
		if hasBlackInColumn(image, mid, m) {
			lo = mid
		} else {
			hi = mid - 1
		}
	}
	return lo
}

// Stub compatibility.
func SmallestRectangleEnclosingBlackPixels() any {
	grid := [][]byte{
		{'0', '0', '1', '0'},
		{'0', '1', '1', '0'},
		{'0', '1', '0', '0'},
	}
	return minArea(grid, 0, 2)
}
```

## 0305 — Number Of Islands Ii

```go
package main

// LeetCode #305: Number of Islands II
// https://leetcode.com/problems/number-of-islands-ii/
// Difficulty: Hard [Paid]
//
// Approach: Union-Find (Disjoint Set Union) with path compression and union by size.
//   - Start with an empty grid of size m x n.
//   - For each land addition (r, c), mark it as land and check 4-directional neighbors.
//   - If a neighbor is also land, union the two cells.
//   - Track the current number of islands after each operation.
//   - Map 2D coordinates to 1D index: id = r * n + c.

import "fmt"

func main() {
	// Example: m=3, n=3, positions=[[0,0],[0,1],[1,2],[2,1]]
	positions := [][]int{{0, 0}, {0, 1}, {1, 2}, {2, 1}}
	result := numIslands2(3, 3, positions)
	fmt.Println("Number of islands after each addition:", result) // [1, 1, 2, 3]

	// Example: all adjacent
	positions2 := [][]int{{0, 0}, {0, 1}, {1, 1}, {1, 0}}
	result2 := numIslands2(2, 2, positions2)
	fmt.Println("All adjacent:", result2) // [1, 1, 1, 1]

	// Example: isolated islands
	positions3 := [][]int{{0, 0}, {0, 2}, {2, 0}, {2, 2}}
	result3 := numIslands2(3, 3, positions3)
	fmt.Println("Isolated:", result3) // [1, 2, 3, 4]

	// Example: duplicate positions.
	positions4 := [][]int{{0, 0}, {0, 0}}
	result4 := numIslands2(1, 1, positions4)
	fmt.Println("Duplicate:", result4) // [1, 1]
}

// UnionFind implements DSU with path compression and union by size.
type UnionFind struct {
	parent []int
	size   []int
}

// NewUnionFind creates a new UnionFind for n elements.
func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	return &UnionFind{parent: parent, size: size}
}

// Find finds the root of x (with path compression).
func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

// Union unions two elements. Returns true if they were actually merged.
func (uf *UnionFind) Union(x, y int) bool {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	if rootX == rootY {
		return false
	}

	// Union by size.
	if uf.size[rootX] < uf.size[rootY] {
		rootX, rootY = rootY, rootX
	}

	uf.parent[rootY] = rootX
	uf.size[rootX] += uf.size[rootY]
	return true
}

// numIslands2 returns the number of islands after each land addition.
func numIslands2(m int, n int, positions [][]int) []int {
	if m <= 0 || n <= 0 || len(positions) == 0 {
		return []int{}
	}

	grid := make([][]bool, m)
	for i := 0; i < m; i++ {
		grid[i] = make([]bool, n)
	}

	uf := NewUnionFind(m * n)
	result := make([]int, 0, len(positions))
	islands := 0

	// 4-directional neighbors.
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for _, pos := range positions {
		r, c := pos[0], pos[1]

		if grid[r][c] {
			// Duplicate position.
			result = append(result, islands)
			continue
		}

		grid[r][c] = true
		islands++

		for _, d := range dirs {
			nr, nc := r+d[0], c+d[1]
			if nr >= 0 && nr < m && nc >= 0 && nc < n && grid[nr][nc] {
				id1 := r*n + c
				id2 := nr*n + nc
				if uf.Union(id1, id2) {
					islands--
				}
			}
		}

		result = append(result, islands)
	}

	return result
}

// Stub compatibility.
func NumberOfIslandsIi() any {
	positions := [][]int{{0, 0}, {0, 1}, {1, 2}, {2, 1}}
	return numIslands2(3, 3, positions)
}
```

## 0312 — Burst Balloons

```go
package main

// LeetCode #312: Burst Balloons
// https://leetcode.com/problems/burst-balloons/
// Difficulty: Hard
//
// Approach: DP Interval (Divide and Conquer).
//   - Add sentinel balloons with value 1 at both ends (index 0 and n+1).
//   - Define dp[i][j] = max coins from bursting all balloons in (i, j) exclusively.
//   - For each k in (i, j), consider k as the LAST balloon to burst in this interval.
//     When k bursts, its neighbors are i and j (since all balloons in between
//     have already been burst).
//   - dp[i][j] = max over k: dp[i][k] + nums[i] * nums[k] * nums[j] + dp[k][j]
//   - Answer: dp[0][n+1] where n is the original length.

import (
	"fmt"
)

func main() {
	// Example 1: [3,1,5,8] -> 167
	// Explanation: nums = [3,1,5,8] -> [3,5,8] -> [3,8] -> [8] -> []
	// coins =  3*1*5    +   3*5*8   +  1*3*8  + 1*8*1 = 167
	nums := []int{3, 1, 5, 8}
	fmt.Println("Burst Balloons:", maxCoins(nums)) // 167

	// Example 2: [1,5] -> 10
	fmt.Println("[1,5]:", maxCoins([]int{1, 5})) // 1*0*5? Wait let me recalc.
	// Actually with sentinel: [1,1,5,1]
	// Burst 1: 1*1*5 + burst 5: 1*5*1 = 5 + 5 = 10. Yes, 10.

	// Example 3: single balloon [5] -> 5
	fmt.Println("[5]:", maxCoins([]int{5})) // 5

	// Example 4: empty
	fmt.Println("[]:", maxCoins([]int{})) // 0
}

// maxCoins returns the maximum coins obtainable by bursting all balloons.
func maxCoins(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	// Add sentinel balloons with value 1.
	arr := make([]int, n+2)
	arr[0] = 1
	arr[n+1] = 1
	for i := 0; i < n; i++ {
		arr[i+1] = nums[i]
	}

	// dp[i][j] = max coins from bursting all balloons strictly between i and j.
	dp := make([][]int, n+2)
	for i := range dp {
		dp[i] = make([]int, n+2)
	}

	// Fill dp by interval length.
	for length := 2; length <= n+1; length++ {
		for i := 0; i+length <= n+1; i++ {
			j := i + length
			// Try each k as the LAST balloon to burst in (i, j).
			for k := i + 1; k < j; k++ {
				// arr[k] is the last to burst, so its neighbors are arr[i] and arr[j].
				coins := dp[i][k] + arr[i]*arr[k]*arr[j] + dp[k][j]
				if coins > dp[i][j] {
					dp[i][j] = coins
				}
			}
		}
	}

	return dp[0][n+1]
}

// max returns the larger of two ints.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Stub compatibility.
func BurstBalloons() any {
	nums := []int{3, 1, 5, 8}
	return maxCoins(nums)
}
```

## 0315 — Count Of Smaller Numbers After Self

```go
package main

// LeetCode #315: Count of Smaller Numbers After Self
// https://leetcode.com/problems/count-of-smaller-numbers-after-self/
// Difficulty: Hard

import "fmt"

func countSmaller(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}

	// Pair each number with its original index
	type pair struct {
		val int
		idx int
	}
	arr := make([]pair, n)
	for i, v := range nums {
		arr[i] = pair{val: v, idx: i}
	}

	result := make([]int, n)

	var mergeSort func([]pair) []pair
	mergeSort = func(a []pair) []pair {
		if len(a) <= 1 {
			return a
		}
		mid := len(a) / 2
		left := mergeSort(a[:mid])
		right := mergeSort(a[mid:])

		// Merge while counting
		merged := make([]pair, 0, len(a))
		i, j := 0, 0
		for i < len(left) && j < len(right) {
			if left[i].val <= right[j].val {
				// All elements already placed from right that are smaller
				result[left[i].idx] += j
				merged = append(merged, left[i])
				i++
			} else {
				merged = append(merged, right[j])
				j++
			}
		}
		for i < len(left) {
			result[left[i].idx] += j
			merged = append(merged, left[i])
			i++
		}
		for j < len(right) {
			merged = append(merged, right[j])
			j++
		}
		return merged
	}

	mergeSort(arr)
	return result
}

func main() {
	// Example 1
	fmt.Println(countSmaller([]int{5, 2, 6, 1}))
	// [2, 1, 1, 0]

	// Example 2
	fmt.Println(countSmaller([]int{-1}))
	// [0]

	// Example 3
	fmt.Println(countSmaller([]int{-1, -1}))
	// [0, 0]
}
```

## 0317 — Shortest Distance From All Buildings

```go
package main

// LeetCode #317: Shortest Distance from All Buildings
// https://leetcode.com/problems/shortest-distance-from-all-buildings/
// Difficulty: Hard [Paid]

import "fmt"

func shortestDistance(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return -1
	}
	rows, cols := len(grid), len(grid[0])

	// totalDist[r][c] = sum of distances from all reachable buildings
	totalDist := make([][]int, rows)
	// reachable[r][c] = count of buildings that can reach this cell
	reachable := make([][]int, rows)
	for r := 0; r < rows; r++ {
		totalDist[r] = make([]int, cols)
		reachable[r] = make([]int, cols)
	}

	totalBuildings := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 1 {
				totalBuildings++
			}
		}
	}

	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	// BFS from each building
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] != 1 {
				continue
			}

			visited := make([][]bool, rows)
			for i := 0; i < rows; i++ {
				visited[i] = make([]bool, cols)
			}

			type cell struct{ r, c int }
			queue := []cell{{r, c}}
			visited[r][c] = true
			dist := 0

			for len(queue) > 0 {
				dist++
				nextQ := []cell{}
				for _, cur := range queue {
					for _, d := range dirs {
						nr, nc := cur.r+d[0], cur.c+d[1]
						if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
							continue
						}
						if visited[nr][nc] || grid[nr][nc] != 0 {
							continue
						}
						visited[nr][nc] = true
						totalDist[nr][nc] += dist
						reachable[nr][nc]++
						nextQ = append(nextQ, cell{nr, nc})
					}
				}
				queue = nextQ
			}
		}
	}

	ans := -1
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 0 && reachable[r][c] == totalBuildings {
				if ans == -1 || totalDist[r][c] < ans {
					ans = totalDist[r][c]
				}
			}
		}
	}
	return ans
}

func main() {
	// Example 1
	grid1 := [][]int{
		{1, 0, 2, 0, 1},
		{0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0},
	}
	fmt.Println(shortestDistance(grid1))
	// 7

	// Example 2
	grid2 := [][]int{{1}}
	fmt.Println(shortestDistance(grid2))
	// -1

	// Example 3
	grid3 := [][]int{{1, 0}}
	fmt.Println(shortestDistance(grid3))
	// 1
}
```

## 0321 — Create Maximum Number

```go
package main

// LeetCode #321: Create Maximum Number
// https://leetcode.com/problems/create-maximum-number/
// Difficulty: Hard

import "fmt"

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	m, n := len(nums1), len(nums2)

	// Pick max subsequence of length length from nums
	maxSubseq := func(nums []int, length int) []int {
		if length == 0 {
			return []int{}
		}
		stack := make([]int, 0, length)
		drop := len(nums) - length
		for _, v := range nums {
			for drop > 0 && len(stack) > 0 && stack[len(stack)-1] < v {
				stack = stack[:len(stack)-1]
				drop--
			}
			stack = append(stack, v)
		}
		return stack[:length]
	}

	// Merge two subsequences into lexicographically largest
	greater := func(a, b []int, i, j int) bool {
		for i < len(a) && j < len(b) {
			if a[i] != b[j] {
				return a[i] > b[j]
			}
			i++
			j++
		}
		return i < len(a)
	}

	merge := func(a, b []int) []int {
		res := make([]int, 0, len(a)+len(b))
		i, j := 0, 0
		for i < len(a) || j < len(b) {
			if j >= len(b) || (i < len(a) && greater(a, b, i, j)) {
				res = append(res, a[i])
				i++
			} else {
				res = append(res, b[j])
				j++
			}
		}
		return res
	}

	best := []int{}
	// Try all valid splits: take i from nums1, k-i from nums2
	start := 0
	if k > n {
		start = k - n
	}
	end := k
	if k > m {
		end = m
	}
	for i := start; i <= end; i++ {
		sub1 := maxSubseq(nums1, i)
		sub2 := maxSubseq(nums2, k-i)
		candidate := merge(sub1, sub2)
		if best == nil || greater(candidate, best, 0, 0) {
			best = candidate
		}
	}
	return best
}

func main() {
	// Example 1
	fmt.Println(maxNumber([]int{3, 4, 6, 5}, []int{9, 1, 2, 5, 8, 3}, 5))
	// [9, 8, 6, 5, 3]

	// Example 2
	fmt.Println(maxNumber([]int{6, 7}, []int{6, 0, 4}, 5))
	// [6, 7, 6, 0, 4]

	// Example 3
	fmt.Println(maxNumber([]int{3, 9}, []int{8, 9}, 3))
	// [9, 8, 9]
}
```

## 0327 — Count Of Range Sum

```go
package main

// LeetCode #327: Count of Range Sum
// https://leetcode.com/problems/count-of-range-sum/
// Difficulty: Hard

import "fmt"

func countRangeSum(nums []int, lower int, upper int) int {
	n := len(nums)
	prefix := make([]int64, n+1)
	for i, v := range nums {
		prefix[i+1] = prefix[i] + int64(v)
	}

	count := 0
	// Temporary buffer for merge sort
	temp := make([]int64, n+1)

	var mergeSort func(left, right int)
	mergeSort = func(left, right int) {
		if left >= right {
			return
		}
		mid := left + (right-left)/2
		mergeSort(left, mid)
		mergeSort(mid+1, right)

		// Count pairs crossing left and right halves
		// For each i in [left, mid], count j in [mid+1, right]
		// where lower <= prefix[j] - prefix[i] <= upper
		// i.e. prefix[i] + lower <= prefix[j] <= prefix[i] + upper
		lo, hi := mid+1, mid+1
		for i := left; i <= mid; i++ {
			for lo <= right && prefix[lo]-prefix[i] < int64(lower) {
				lo++
			}
			for hi <= right && prefix[hi]-prefix[i] <= int64(upper) {
				hi++
			}
			count += hi - lo
		}

		// Merge
		i, j, k := left, mid+1, left
		for i <= mid && j <= right {
			if prefix[i] <= prefix[j] {
				temp[k] = prefix[i]
				i++
			} else {
				temp[k] = prefix[j]
				j++
			}
			k++
		}
		for i <= mid {
			temp[k] = prefix[i]
			i++
			k++
		}
		for j <= right {
			temp[k] = prefix[j]
			j++
			k++
		}
		for i := left; i <= right; i++ {
			prefix[i] = temp[i]
		}
	}

	mergeSort(0, n)
	return count
}

func main() {
	// Example 1
	fmt.Println(countRangeSum([]int{-2, 5, -1}, -2, 2))
	// 3

	// Example 2
	fmt.Println(countRangeSum([]int{0}, 0, 0))
	// 1

	// Example 3
	fmt.Println(countRangeSum([]int{2147483647, -2147483648, -1, 0}, -1, 0))
	// 4
}
```

## 0329 — Longest Increasing Path In A Matrix

```go
package main

// LeetCode #329: Longest Increasing Path in a Matrix
// https://leetcode.com/problems/longest-increasing-path-in-a-matrix/
// Difficulty: Hard

import "fmt"

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	rows, cols := len(matrix), len(matrix[0])
	memo := make([][]int, rows)
	for r := 0; r < rows; r++ {
		memo[r] = make([]int, cols)
	}

	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	var dfs func(r, c int) int
	dfs = func(r, c int) int {
		if memo[r][c] != 0 {
			return memo[r][c]
		}
		best := 1
		for _, d := range dirs {
			nr, nc := r+d[0], c+d[1]
			if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
				continue
			}
			if matrix[nr][nc] <= matrix[r][c] {
				continue
			}
			path := 1 + dfs(nr, nc)
			if path > best {
				best = path
			}
		}
		memo[r][c] = best
		return best
	}

	ans := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if path := dfs(r, c); path > ans {
				ans = path
			}
		}
	}
	return ans
}

func main() {
	// Example 1
	matrix1 := [][]int{
		{9, 9, 4},
		{6, 6, 8},
		{2, 1, 1},
	}
	fmt.Println(longestIncreasingPath(matrix1))
	// 4

	// Example 2
	matrix2 := [][]int{
		{3, 4, 5},
		{3, 2, 6},
		{2, 2, 1},
	}
	fmt.Println(longestIncreasingPath(matrix2))
	// 4

	// Example 3
	matrix3 := [][]int{{1}}
	fmt.Println(longestIncreasingPath(matrix3))
	// 1
}
```

## 0330 — Patching Array

```go
package main

// LeetCode #330: Patching Array
// https://leetcode.com/problems/patching-array/
// Difficulty: Hard

import "fmt"

func minPatches(nums []int, n int) int {
	patches := 0
	miss := int64(1) // smallest sum we cannot form
	i := 0

	for miss <= int64(n) {
		if i < len(nums) && int64(nums[i]) <= miss {
			miss += int64(nums[i])
			i++
		} else {
			// Patch with miss itself
			miss += miss
			patches++
		}
	}
	return patches
}

func main() {
	// Example 1
	fmt.Println(minPatches([]int{1, 3}, 6))
	// 1

	// Example 2
	fmt.Println(minPatches([]int{1, 5, 10}, 20))
	// 2

	// Example 3
	fmt.Println(minPatches([]int{1, 2, 2}, 5))
	// 0

	// Example 4
	fmt.Println(minPatches([]int{1, 2, 31, 33}, 2147483647))
	// 28
}
```

## 0332 — Reconstruct Itinerary

```go
package main

// LeetCode #332: Reconstruct Itinerary
// https://leetcode.com/problems/reconstruct-itinerary/
// Difficulty: Hard

import (
	"fmt"
	"sort"
)

func findItinerary(tickets [][]string) []string {
	// Build adjacency list with sorting
	graph := make(map[string][]string)
	for _, t := range tickets {
		from, to := t[0], t[1]
		graph[from] = append(graph[from], to)
	}
	for from := range graph {
		sort.Sort(sort.Reverse(sort.StringSlice(graph[from])))
	}

	// Hierholzer's algorithm: Eulerian path
	route := []string{}

	var dfs func(airport string)
	dfs = func(airport string) {
		for len(graph[airport]) > 0 {
			next := graph[airport][len(graph[airport])-1]
			graph[airport] = graph[airport][:len(graph[airport])-1]
			dfs(next)
		}
		route = append(route, airport)
	}

	dfs("JFK")

	// Reverse to get correct order
	for i, j := 0, len(route)-1; i < j; i, j = i+1, j-1 {
		route[i], route[j] = route[j], route[i]
	}
	return route
}

func main() {
	// Example 1
	tickets1 := [][]string{
		{"MUC", "LHR"},
		{"JFK", "MUC"},
		{"SFO", "SJC"},
		{"LHR", "SFO"},
		{"JFK", "ATL"},
		{"ATL", "JFK"},
	}
	fmt.Println(findItinerary(tickets1))
	// [JFK ATL JFK MUC LHR SFO SJC]

	// Example 2
	tickets2 := [][]string{
		{"JFK", "SFO"},
		{"JFK", "ATL"},
		{"SFO", "ATL"},
		{"ATL", "JFK"},
		{"ATL", "SFO"},
	}
	fmt.Println(findItinerary(tickets2))
	// [JFK ATL JFK SFO ATL SFO]

	// Example 3
	tickets3 := [][]string{
		{"JFK", "KUL"},
		{"JFK", "NRT"},
		{"NRT", "JFK"},
	}
	fmt.Println(findItinerary(tickets3))
	// [JFK NRT JFK KUL]
}
```

## 0335 — Self Crossing

```go
package main

// LeetCode #335: Self Crossing
// https://leetcode.com/problems/self-crossing/
// Difficulty: Hard

import "fmt"

func isSelfCrossing(distance []int) bool {
	n := len(distance)
	if n < 4 {
		return false
	}

	for i := 3; i < n; i++ {
		// Case 1: i-th line crosses (i-3)-th line
		// x[i-3] >= x[i-1] and x[i] >= x[i-2]
		if distance[i] >= distance[i-2] && distance[i-1] <= distance[i-3] {
			return true
		}

		// Case 2: i-th line crosses (i-4)-th line (overlap of 5 edges)
		if i >= 4 {
			if distance[i-1] == distance[i-3] &&
				distance[i] >= distance[i-2]-distance[i-4] {
				return true
			}
		}

		// Case 3: i-th line crosses (i-5)-th line (overlap of 6 edges)
		if i >= 5 {
			if distance[i-2] >= distance[i-4] &&
				distance[i-3]-distance[i-5] >= 0 &&
				distance[i-1] >= distance[i-3]-distance[i-5] &&
				distance[i-1] <= distance[i-3] &&
				distance[i] >= distance[i-2]-distance[i-4] {
				return true
			}
		}
	}
	return false
}

func main() {
	// Example 1: crossing
	fmt.Println(isSelfCrossing([]int{2, 1, 1, 2}))
	// true

	// Example 2: no crossing
	fmt.Println(isSelfCrossing([]int{1, 2, 3, 4}))
	// false

	// Example 3: crossing
	fmt.Println(isSelfCrossing([]int{1, 1, 1, 2, 1}))
	// true

	// Example 4: no crossing
	fmt.Println(isSelfCrossing([]int{1, 1, 2, 2, 3, 3, 4, 4, 10, 4, 4, 3, 3, 2, 2, 1, 1}))
	// false
}
```

## 0336 — Palindrome Pairs

```go
package main

// LeetCode #336: Palindrome Pairs
// https://leetcode.com/problems/palindrome-pairs/
// Difficulty: Hard

import "fmt"

func palindromePairs(words []string) [][]int {
	// Build trie of reversed words
	type trieNode struct {
		child [26]*trieNode
		idx   int // index of word ending here, -1 if none
	}

	root := &trieNode{idx: -1}

	// Insert reversed word into trie
	for i, w := range words {
		node := root
		for j := len(w) - 1; j >= 0; j-- {
			c := w[j] - 'a'
			if node.child[c] == nil {
				node.child[c] = &trieNode{idx: -1}
			}
			node = node.child[c]
		}
		node.idx = i
	}

	isPalindrome := func(s string, l, r int) bool {
		for l < r {
			if s[l] != s[r] {
				return false
			}
			l++
			r--
		}
		return true
	}

	result := [][]int{}

	for i, w := range words {
		node := root

		// Check if remaining part of current word is palindrome and trie has a word ending here
		for j := 0; j < len(w); j++ {
			if node.idx != -1 && node.idx != i {
				if isPalindrome(w, j, len(w)-1) {
					result = append(result, []int{i, node.idx})
				}
			}
			c := w[j] - 'a'
			if node.child[c] == nil {
				node = nil
				break
			}
			node = node.child[c]
		}

		if node == nil {
			continue
		}

		// Exact match: reversed word completely matches current word
		if node.idx != -1 && node.idx != i {
			result = append(result, []int{i, node.idx})
		}

		// Check remaining trie paths where remaining prefix is palindrome
		// (current word is shorter than the trie word)
		var dfs func(*trieNode, []byte)
		dfs = func(n *trieNode, prefix []byte) {
			for c := 0; c < 26; c++ {
				if n.child[c] != nil {
					prefix = append(prefix, byte('a'+c))
					if n.child[c].idx != -1 && n.child[c].idx != i {
						if isPalindrome(string(prefix), 0, len(prefix)-1) {
							result = append(result, []int{i, n.child[c].idx})
						}
					}
					dfs(n.child[c], prefix)
					prefix = prefix[:len(prefix)-1]
				}
			}
		}
		dfs(node, []byte{})
	}

	return result
}

func main() {
	// Example 1
	fmt.Println(palindromePairs([]string{"abcd", "dcba", "lls", "s", "sssll"}))
	// [[0 1] [1 0] [3 2] [2 4]]

	// Example 2
	fmt.Println(palindromePairs([]string{"bat", "tab", "cat"}))
	// [[0 1] [1 0]]

	// Example 3
	fmt.Println(palindromePairs([]string{"a", ""}))
	// [[0 1] [1 0]]
}
```

## 0352 — Data Stream As Disjoint Intervals

```go
package main

// LeetCode #352: Data Stream as Disjoint Intervals
// https://leetcode.com/problems/data-stream-as-disjoint-intervals/
// Difficulty: Hard

import (
	"fmt"
	"sort"
)

// SummaryRanges maintains a set of disjoint intervals from a stream of numbers.
type SummaryRanges struct {
	intervals [][]int // sorted by start, each is [start, end] inclusive
}

func Constructor() SummaryRanges {
	return SummaryRanges{}
}

func (sr *SummaryRanges) AddNum(val int) {
	// Find position using binary search
	idx := sort.Search(len(sr.intervals), func(i int) bool {
		return sr.intervals[i][0] > val
	})

	// Check for merge with left neighbor (intervals[idx-1])
	if idx > 0 && sr.intervals[idx-1][1] >= val-1 {
		// Merge with left
		if val > sr.intervals[idx-1][1] {
			sr.intervals[idx-1][1] = val
		}
		// Check if we also need to merge with right
		if idx < len(sr.intervals) && sr.intervals[idx][0] <= sr.intervals[idx-1][1]+1 {
			if sr.intervals[idx][1] > sr.intervals[idx-1][1] {
				sr.intervals[idx-1][1] = sr.intervals[idx][1]
			}
			sr.intervals = append(sr.intervals[:idx], sr.intervals[idx+1:]...)
		}
	} else if idx < len(sr.intervals) && sr.intervals[idx][0] <= val+1 {
		// Merge with right only
		if val < sr.intervals[idx][0] {
			sr.intervals[idx][0] = val
		}
	} else {
		// New isolated interval - insert at position idx
		newInterval := []int{val, val}
		// Insert at idx
		sr.intervals = append(sr.intervals, nil)
		copy(sr.intervals[idx+1:], sr.intervals[idx:])
		sr.intervals[idx] = newInterval
	}
}

func (sr *SummaryRanges) GetIntervals() [][]int {
	return sr.intervals
}

func main() {
	// Example 1
	sr := Constructor()
	sr.AddNum(1)
	sr.AddNum(3)
	sr.AddNum(7)
	sr.AddNum(2)
	sr.AddNum(6)
	fmt.Println(sr.GetIntervals())
	// [[1 3] [6 7]]

	// Example 2
	sr2 := Constructor()
	sr2.AddNum(1)
	sr2.AddNum(3)
	sr2.AddNum(2)
	fmt.Println(sr2.GetIntervals())
	// [[1 3]]

	// Example 3: empty
	sr3 := Constructor()
	fmt.Println(sr3.GetIntervals())
	// []

	// Example 4: single
	sr4 := Constructor()
	sr4.AddNum(5)
	fmt.Println(sr4.GetIntervals())
	// [[5 5]]
}
```

## 0354 — Russian Doll Envelopes

```go
package main

// LeetCode #354: Russian Doll Envelopes
// https://leetcode.com/problems/russian-doll-envelopes/
// Difficulty: Hard
//
// Sort envelopes by width ascending. When widths are equal, sort by height
// descending to prevent same-width nesting. Then find LIS (Longest Increasing
// Subsequence) on heights using patience sorting O(n log n).

import (
	"fmt"
	"sort"
)

func main() {
	// Example 1: [[5,4],[6,4],[6,7],[2,3]] -> 3 (2,3 -> 5,4 -> 6,7)
	fmt.Println(maxEnvelopes([][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}))
	// Example 2: [[1,1],[1,1],[1,1]] -> 1
	fmt.Println(maxEnvelopes([][]int{{1, 1}, {1, 1}, {1, 1}}))
	// Example 3: [[1,2],[2,3],[3,4],[4,5]] -> 4
	fmt.Println(maxEnvelopes([][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}))
	// Edge: single envelope
	fmt.Println(maxEnvelopes([][]int{{1, 1}}))
	// Edge: all same width
	fmt.Println(maxEnvelopes([][]int{{1, 2}, {1, 3}, {1, 4}}))
}

func maxEnvelopes(envelopes [][]int) int {
	// Sort by width ascending; if width ties, height descending
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] != envelopes[j][0] {
			return envelopes[i][0] < envelopes[j][0]
		}
		return envelopes[i][1] > envelopes[j][1]
	})

	// LIS on heights using patience sorting (binary search)
	tails := make([]int, 0, len(envelopes))
	for _, e := range envelopes {
		h := e[1]
		idx := sort.SearchInts(tails, h)
		if idx == len(tails) {
			tails = append(tails, h)
		} else {
			tails[idx] = h
		}
	}
	return len(tails)
}
```

## 0358 — Rearrange String K Distance Apart

```go
package main

// LeetCode #358: Rearrange String k Distance Apart
// https://leetcode.com/problems/rearrange-string-k-distance-apart/
// Difficulty: Hard [Paid]
//
// Use a max-heap (by frequency) to always pick the most frequent available
// character. A cooldown queue enforces that a character cannot be reused
// within k distance. If the heap empties before the string is built, it's
// impossible and we return "".

import (
	"container/heap"
	"fmt"
)

func main() {
	// Example 1: "aabbcc", k=3 -> "abcabc" (or any valid)
	fmt.Println(rearrangeString("aabbcc", 3))
	// Example 2: "aaabc", k=3 -> "" (impossible)
	fmt.Println(rearrangeString("aaabc", 3))
	// Example 3: "aaadbbcc", k=2 -> "abacabcd" (or valid)
	fmt.Println(rearrangeString("aaadbbcc", 2))
	// Edge: k=0
	fmt.Println(rearrangeString("aabb", 0))
	// Edge: single char
	fmt.Println(rearrangeString("a", 1))
}

type charFreq struct {
	ch    byte
	count int
}

type maxHeap []charFreq

func (h maxHeap) Len() int            { return len(h) }
func (h maxHeap) Less(i, j int) bool  { return h[i].count > h[j].count }
func (h maxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x interface{}) { *h = append(*h, x.(charFreq)) }
func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type cooldownItem struct {
	ch      byte
	count   int
	readyAt int // position when this char becomes available again
}

func rearrangeString(s string, k int) string {
	if k <= 1 {
		return s
	}

	// Count frequencies
	freq := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		freq[s[i]]++
	}

	// Build max-heap
	h := &maxHeap{}
	heap.Init(h)
	for ch, cnt := range freq {
		heap.Push(h, charFreq{ch, cnt})
	}

	result := make([]byte, 0, len(s))
	q := make([]cooldownItem, 0)

	for len(result) < len(s) {
		// Replenish cooldown items whose readyAt <= current position
		if len(q) > 0 && q[0].readyAt <= len(result) {
			item := q[0]
			q = q[1:]
			heap.Push(h, charFreq{item.ch, item.count})
		}

		if h.Len() == 0 {
			return "" // impossible
		}

		cf := heap.Pop(h).(charFreq)
		result = append(result, cf.ch)
		cf.count--
		if cf.count > 0 {
			q = append(q, cooldownItem{cf.ch, cf.count, len(result) - 1 + k})
		}
	}

	return string(result)
}
```

## 0363 — Max Sum Of Rectangle No Larger Than K

```go
package main

// LeetCode #363: Max Sum of Rectangle No Larger Than K
// https://leetcode.com/problems/max-sum-of-rectangle-no-larger-than-k/
// Difficulty: Hard
//
// For each pair of rows (top, bottom), compute column sums and apply a
// 1D Kadane-with-bound approach: maintain a sorted prefix-sum list and
// binary-search for the smallest prefix >= current - k.
// O(m^2 * n * log n) time, O(n) space.

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	// Example 1: [[1,0,1],[0,-2,3]], k=2 -> 2 ([[0,-2],[0,3]] sum=2)
	fmt.Println(maxSumSubmatrix([][]int{{1, 0, 1}, {0, -2, 3}}, 2))
	// Example 2: [[2,2,-1]], k=3 -> 3
	fmt.Println(maxSumSubmatrix([][]int{{2, 2, -1}}, 3))
	// Example 3: [[1]], k=1 -> 1
	fmt.Println(maxSumSubmatrix([][]int{{1}}, 0))
	// Edge: larger matrix
	fmt.Println(maxSumSubmatrix([][]int{{5, -4, -3, 4}, {-3, -4, 4, 5}, {5, 1, 5, -4}}, 10))
}

func maxSumSubmatrix(matrix [][]int, k int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	rows, cols := len(matrix), len(matrix[0])
	best := math.MinInt64

	for top := 0; top < rows; top++ {
		colSums := make([]int, cols)
		for bottom := top; bottom < rows; bottom++ {
			// Accumulate row bottom into column sums
			for c := 0; c < cols; c++ {
				colSums[c] += matrix[bottom][c]
			}

			// 1D max subarray sum no larger than k on colSums
			best = max(best, maxSumNoLargerThanK(colSums, k))
			if best == k {
				return k
			}
		}
	}
	return best
}

// maxSumNoLargerThanK finds max subarray sum <= k using sorted prefix sums.
func maxSumNoLargerThanK(arr []int, k int) int {
	prefix := 0
	best := math.MinInt64

	// Sorted prefix sums (we maintain a sorted list)
	sorted := []int{0}

	for _, v := range arr {
		prefix += v
		// Find smallest prefix in sorted >= prefix - k
		idx := sort.SearchInts(sorted, prefix-k)
		if idx < len(sorted) {
			best = max(best, prefix-sorted[idx])
		}
		// Insert prefix into sorted position
		ins := sort.SearchInts(sorted, prefix)
		sorted = append(sorted, 0)
		copy(sorted[ins+1:], sorted[ins:])
		sorted[ins] = prefix
	}

	return best
}
```

## 0381 — Insert Delete Getrandom O1 Duplicates Allowed

```go
package main

// LeetCode #381: Insert Delete GetRandom O(1) - Duplicates allowed
// https://leetcode.com/problems/insert-delete-getrandom-o1-duplicates-allowed/
// Difficulty: Hard
//
// RandomizedCollection supports duplicates with O(1) average time.
// Uses a slice to store values and a map from value to a set of indices.
// Remove swaps the target with the last element to maintain O(1).

import (
	"fmt"
	"math/rand"
)

func main() {
	rc := Constructor()
	fmt.Println(rc.Insert(1))  // true
	fmt.Println(rc.Insert(1))  // false (duplicate exists)
	fmt.Println(rc.Insert(2))  // true
	fmt.Println(rc.GetRandom()) // 1 or 2
	fmt.Println(rc.Remove(1))  // true (removes one occurrence of 1)
	fmt.Println(rc.GetRandom()) // 1 or 2
	fmt.Println(rc.Remove(1))  // true (removes the last 1)
	fmt.Println(rc.Remove(1))  // false (no more 1s)
}

type RandomizedCollection struct {
	vals []int
	idx  map[int]map[int]struct{} // val -> set of indices
}

func Constructor() RandomizedCollection {
	return RandomizedCollection{idx: make(map[int]map[int]struct{})}
}

func (rc *RandomizedCollection) Insert(val int) bool {
	indices := rc.idx[val]
	if indices == nil {
		indices = make(map[int]struct{})
		rc.idx[val] = indices
	}
	rc.vals = append(rc.vals, val)
	indices[len(rc.vals)-1] = struct{}{}
	return len(indices) == 1
}

func (rc *RandomizedCollection) Remove(val int) bool {
	indices := rc.idx[val]
	if len(indices) == 0 {
		return false
	}

	// Pick any index of val
	var removeIdx int
	for k := range indices {
		removeIdx = k
		break
	}

	lastIdx := len(rc.vals) - 1
	lastVal := rc.vals[lastIdx]

	if removeIdx != lastIdx {
		// Swap with last element
		rc.vals[removeIdx] = lastVal
		// Update lastVal's indices: remove lastIdx, add removeIdx
		delete(rc.idx[lastVal], lastIdx)
		rc.idx[lastVal][removeIdx] = struct{}{}
	}

	// Remove val's index and truncate
	delete(indices, removeIdx)
	rc.vals = rc.vals[:lastIdx]

	return true
}

func (rc *RandomizedCollection) GetRandom() int {
	return rc.vals[rand.Intn(len(rc.vals))]
}
```

## 0391 — Perfect Rectangle

```go
package main

// LeetCode #391: Perfect Rectangle
// https://leetcode.com/problems/perfect-rectangle/
// Difficulty: Hard
//
// A perfect rectangle exactly covers its bounding box with no gaps or
// overlaps. Approach: (1) sum of areas == bounding box area; (2) track
// corner parity — internal corners appear an even number of times, valid
// corners exactly once.

import (
	"fmt"
)

func main() {
	// Example 1: true
	fmt.Println(isRectangleCover([][]int{
		{1, 1, 3, 3}, {3, 1, 4, 2}, {3, 2, 4, 4}, {1, 3, 2, 4}, {2, 3, 3, 4},
	}))
	// Example 2: false (gap)
	fmt.Println(isRectangleCover([][]int{
		{1, 1, 2, 3}, {1, 3, 2, 4}, {3, 1, 4, 2}, {3, 2, 4, 4},
	}))
	// Example 3: false (overlap)
	fmt.Println(isRectangleCover([][]int{
		{1, 1, 3, 3}, {3, 1, 4, 2}, {1, 3, 2, 4}, {3, 2, 4, 4},
	}))
	// Edge: single rectangle
	fmt.Println(isRectangleCover([][]int{{0, 0, 1, 1}}))
	// Edge: two adjacent
	fmt.Println(isRectangleCover([][]int{{0, 0, 1, 1}, {0, 1, 1, 2}}))
}

func isRectangleCover(rectangles [][]int) bool {
	if len(rectangles) == 0 {
		return false
	}

	type point struct{ x, y int }
	corners := make(map[point]int)

	minX, minY, maxX, maxY := 1<<30, 1<<30, -1<<30, -1<<30
	totalArea := 0

	for _, r := range rectangles {
		x1, y1, x2, y2 := r[0], r[1], r[2], r[3]
		totalArea += (x2 - x1) * (y2 - y1)

		if x1 < minX {
			minX = x1
		}
		if y1 < minY {
			minY = y1
		}
		if x2 > maxX {
			maxX = x2
		}
		if y2 > maxY {
			maxY = y2
		}

		// Toggle corner parity
		corners[point{x1, y1}]++
		corners[point{x1, y2}]++
		corners[point{x2, y1}]++
		corners[point{x2, y2}]++
	}

	// Area check
	bBoxArea := (maxX - minX) * (maxY - minY)
	if totalArea != bBoxArea {
		return false
	}

	// After all toggles, exactly 4 corners should remain (odd count = 1)
	// All internal corners must have even count.
	expected := map[point]bool{
		{minX, minY}: true,
		{minX, maxY}: true,
		{maxX, minY}: true,
		{maxX, maxY}: true,
	}

	for p, cnt := range corners {
		if cnt%2 != 0 {
			if !expected[p] {
				return false
			}
		}
	}

	// Verify that bounding box corners appear exactly once
	for p := range expected {
		if corners[p]%2 != 1 {
			return false
		}
	}

	return true
}
```

## 0403 — Frog Jump

```go
package main

// LeetCode #403: Frog Jump
// https://leetcode.com/problems/frog-jump/
// Difficulty: Hard
//
// A frog starts on stone 0 and jumps units of k. From stone i with a jump of
// size k, the next jump can be k-1, k, or k+1. Determine if the frog can
// reach the last stone.
// DP: map[stone]->set of jump sizes that can reach that stone.

import (
	"fmt"
)

func main() {
	// Example 1: [0,1,3,5,6,8,12,17] -> true
	fmt.Println(canCross([]int{0, 1, 3, 5, 6, 8, 12, 17}))
	// Example 2: [0,1,2,3,4,8,9,11] -> false
	fmt.Println(canCross([]int{0, 1, 2, 3, 4, 8, 9, 11}))
	// Edge: two stones
	fmt.Println(canCross([]int{0, 1}))
	// Edge: three stones, possible
	fmt.Println(canCross([]int{0, 1, 3}))
	// Edge: three stones, impossible (can't make jump 2 from stone 1)
	fmt.Println(canCross([]int{0, 2}))
}

func canCross(stones []int) bool {
	if len(stones) == 0 {
		return false
	}

	// Map stone position -> set of jump sizes
	dp := make(map[int]map[int]bool, len(stones))
	for _, s := range stones {
		dp[s] = make(map[int]bool)
	}
	dp[stones[0]][0] = true // start with jump 0

	lastStone := stones[len(stones)-1]
	stoneSet := make(map[int]bool, len(stones))
	for _, s := range stones {
		stoneSet[s] = true
	}

	for _, pos := range stones {
		for jump := range dp[pos] {
			for k := jump - 1; k <= jump+1; k++ {
				if k <= 0 {
					continue
				}
				next := pos + k
				if next == lastStone {
					return true
				}
				if stoneSet[next] {
					dp[next][k] = true
				}
			}
		}
	}

	return false
}
```

## 0407 — Trapping Rain Water Ii

```go
package main

// LeetCode #407: Trapping Rain Water II
// https://leetcode.com/problems/trapping-rain-water-ii/
// Difficulty: Hard
//
// Use a min-heap (priority queue) to BFS from the border inward. Always
// process the lowest border cell. If a neighbor is lower, water accumulates
// (border height - neighbor height) and the neighbor is "raised" to the
// border height before being pushed back.

import (
	"container/heap"
	"fmt"
)

func main() {
	// Example 1: [[1,4,3,1,3,2],[3,2,1,3,2,4],[2,3,3,2,3,1]] -> 4
	fmt.Println(trapRainWater([][]int{
		{1, 4, 3, 1, 3, 2},
		{3, 2, 1, 3, 2, 4},
		{2, 3, 3, 2, 3, 1},
	}))
	// Example 2: [[3,3,3,3,3],[3,2,2,2,3],[3,2,1,2,3],[3,2,2,2,3],[3,3,3,3,3]] -> 10
	fmt.Println(trapRainWater([][]int{
		{3, 3, 3, 3, 3},
		{3, 2, 2, 2, 3},
		{3, 2, 1, 2, 3},
		{3, 2, 2, 2, 3},
		{3, 3, 3, 3, 3},
	}))
	// Edge: single row
	fmt.Println(trapRainWater([][]int{{1, 2, 1}}))
	// Edge: single column
	fmt.Println(trapRainWater([][]int{{1}, {2}, {1}}))
}

type cell struct {
	h, r, c int
}

type minHeap []cell

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i].h < h[j].h }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(cell)) }
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func trapRainWater(heightMap [][]int) int {
	if len(heightMap) < 3 || len(heightMap[0]) < 3 {
		return 0
	}
	rows, cols := len(heightMap), len(heightMap[0])

	visited := make([][]bool, rows)
	for r := range visited {
		visited[r] = make([]bool, cols)
	}

	h := &minHeap{}
	heap.Init(h)

	// Push all border cells
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if r == 0 || r == rows-1 || c == 0 || c == cols-1 {
				heap.Push(h, cell{heightMap[r][c], r, c})
				visited[r][c] = true
			}
		}
	}

	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	total := 0

	for h.Len() > 0 {
		cur := heap.Pop(h).(cell)
		for _, d := range dirs {
			nr, nc := cur.r+d[0], cur.c+d[1]
			if nr >= 0 && nr < rows && nc >= 0 && nc < cols && !visited[nr][nc] {
				visited[nr][nc] = true
				if heightMap[nr][nc] < cur.h {
					total += cur.h - heightMap[nr][nc]
					heap.Push(h, cell{cur.h, nr, nc})
				} else {
					heap.Push(h, cell{heightMap[nr][nc], nr, nc})
				}
			}
		}
	}

	return total
}
```

## 0410 — Split Array Largest Sum

```go
package main

// LeetCode #410: Split Array Largest Sum
// https://leetcode.com/problems/split-array-largest-sum/
// Difficulty: Hard
//
// Binary search on the answer. The minimum possible largest sum is max(nums)
// and the maximum is sum(nums). For each candidate mid, greedily check if
// we can split the array into k contiguous subarrays each with sum <= mid.

import (
	"fmt"
)

func main() {
	// Example 1: [7,2,5,10,8], k=2 -> 18 (split at 10: [7,2,5] sum=14, [10,8] sum=18)
	fmt.Println(splitArray([]int{7, 2, 5, 10, 8}, 2))
	// Example 2: [1,2,3,4,5], k=2 -> 9
	fmt.Println(splitArray([]int{1, 2, 3, 4, 5}, 2))
	// Example 3: [1,4,4], k=3 -> 4
	fmt.Println(splitArray([]int{1, 4, 4}, 3))
	// Edge: single element
	fmt.Println(splitArray([]int{10}, 1))
	// Edge: all same
	fmt.Println(splitArray([]int{1, 1, 1, 1, 1}, 3))
}

func splitArray(nums []int, k int) int {
	left, right := 0, 0
	for _, v := range nums {
		if v > left {
			left = v
		}
		right += v
	}

	for left < right {
		mid := left + (right-left)/2
		if feasible(nums, k, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func feasible(nums []int, k, target int) bool {
	count, sum := 1, 0
	for _, v := range nums {
		if sum+v > target {
			count++
			sum = v
			if count > k {
				return false
			} else {
				continue
			}
		}
		sum += v
	}
	return true
}
```

## 0411 — Minimum Unique Word Abbreviation

```go
package main

// LeetCode #411: Minimum Unique Word Abbreviation
// https://leetcode.com/problems/minimum-unique-word-abbreviation/
// Difficulty: Hard [Paid]
//
// Enumerate all 2^n abbreviation patterns via bitmask. A bit = 1 means the
// character is abbreviated (counted as part of a number); bit = 0 means the
// character is kept literally. For each mask, check if ANY dictionary word
// of the same length matches all literal positions. If none matches, the
// abbreviation is valid. Return the shortest (and lexicographically smallest
// among ties).

import (
	"fmt"
	"strconv"
)

func main() {
	// Example 1: target="apple", dictionary=["blade"] -> "a4"
	fmt.Println(minAbbreviation("apple", []string{"blade"}))
	// Example 2: target="apple", dictionary=["plain", "amber", "blade"] -> "1p3"
	fmt.Println(minAbbreviation("apple", []string{"plain", "amber", "blade"}))
	// Example 3: target="usa", dictionary=["usa"] -> "3" (no unique abbr shorter)
	fmt.Println(minAbbreviation("usaandchinaaregreat", []string{"usaandchinaaregreat"}))
	// Edge: no dictionary
	fmt.Println(minAbbreviation("hello", []string{}))
	// Edge: same length but different chars
	fmt.Println(minAbbreviation("abcde", []string{"fghij"}))
}

func minAbbreviation(target string, dictionary []string) string {
	n := len(target)

	// Filter to same-length words only
	sameLen := make([]string, 0, len(dictionary))
	for _, w := range dictionary {
		if len(w) == n {
			sameLen = append(sameLen, w)
		}
	}

	// If no same-length words, just abbreviate entire word
	if len(sameLen) == 0 {
		return strconv.Itoa(n)
	}

	best := target // worst case: no abbreviation
	masks := 1 << uint(n)

	for mask := 0; mask < masks; mask++ {
		// Check if any dict word matches at all literal positions
		conflict := false
		for _, w := range sameLen {
			match := true
			for i := 0; i < n; i++ {
				if mask>>uint(i)&1 == 0 { // literal position
					if target[i] != w[i] {
						match = false
						break
					}
				}
			}
			if match {
				conflict = true
				break
			}
		}
		if conflict {
			continue
		}

		ab := abbreviate(target, mask)
		// Shorter is better; ties: lexicographically smaller
		if len(ab) < len(best) || (len(ab) == len(best) && ab < best) {
			best = ab
		}
	}

	return best
}

func abbreviate(word string, mask int) string {
	n := len(word)
	var res []byte
	count := 0

	for i := 0; i < n; i++ {
		if mask>>uint(i)&1 == 1 {
			count++
		} else {
			if count > 0 {
				res = append(res, strconv.Itoa(count)...)
				count = 0
			}
			res = append(res, word[i])
		}
	}
	if count > 0 {
		res = append(res, strconv.Itoa(count)...)
	}
	return string(res)
}
```

## 0420 — Strong Password Checker

```go
package main

// LeetCode #420: Strong Password Checker
// https://leetcode.com/problems/strong-password-checker/
// Difficulty: Hard
//
// A password is strong if:
//   1. Length 6..20
//   2. Contains at least one lowercase, one uppercase, one digit
//   3. No three consecutive repeating characters
// Returns the minimum number of changes (insert, delete, replace).
// Handles three cases: too short, too long, or in-range with missing types.

import (
	"fmt"
)

func main() {
	// Example 1: "a" -> 5 (insert 5 chars: need length+missing types)
	fmt.Println(strongPasswordChecker("a"))
	// Example 2: "aA1" -> 3 (need 3 more chars + missing types)
	fmt.Println(strongPasswordChecker("aA1"))
	// Example 3: "1337C0d3" -> 0 (already strong)
	fmt.Println(strongPasswordChecker("1337C0d3"))
	// Example 4: "aaa123" -> 1 (replace one 'a')
	fmt.Println(strongPasswordChecker("aaa123"))
	// Example 5: "aaa" -> 3 (insert 3: need length + fix repeat + missing types)
	fmt.Println(strongPasswordChecker("aaa"))
	// Example 6: "abababababababababaaa" (20+ chars with repeating) -> 3
	fmt.Println(strongPasswordChecker("abababababababababaaa"))
}

func strongPasswordChecker(password string) int {
	n := len(password)

	// Count missing character types
	hasLower, hasUpper, hasDigit := false, false, false
	for i := 0; i < n; i++ {
		ch := password[i]
		if ch >= 'a' && ch <= 'z' {
			hasLower = true
		} else if ch >= 'A' && ch <= 'Z' {
			hasUpper = true
		} else if ch >= '0' && ch <= '9' {
			hasDigit = true
		}
	}

	missing := 0
	if !hasLower {
		missing++
	}
	if !hasUpper {
		missing++
	}
	if !hasDigit {
		missing++
	}

	if n < 6 {
		// Too short. We need to add (6-n) chars.
		// Each add can fix one missing type and one repeat.
		return max(6-n, missing)
	}

	if n <= 20 {
		// In range. Only need to fix repeats and missing types.
		// One replace fixes one repeat block (every 3rd char) or one missing type.
		replaces := 0
		i := 0
		for i < n {
			j := i
			for j < n && password[j] == password[i] {
				j++
			}
			repeats := j - i
			if repeats >= 3 {
				replaces += repeats / 3
			}
			i = j
		}
		return max(replaces, missing)
	}

	// Too long (n > 20). Need to delete (n-20) chars + fix repeats + fix missing.
	// Strategy: delete chars strategically to break up repeat sequences.
	// For each repeat block of length L:
	//   - Deleting 1 reduces replacements needed by 1 if L%3 == 0
	//   - Deleting 2 reduces replacements needed by 1 if L%3 == 1 (after first delete)
	//   - Deleting 3 reduces replacements needed by 1 otherwise
	// We prioritize deletes that give the biggest reduction in replaces.

	over := n - 20
	replaces := 0

	// Count repeats and group by L%3
	type repeatInfo struct{ length int }
	repeats := make([]repeatInfo, 0)

	i := 0
	for i < n {
		j := i
		for j < n && password[j] == password[i] {
			j++
		}
		l := j - i
		if l >= 3 {
			repeats = append(repeats, repeatInfo{l})
			replaces += l / 3
		}
		i = j
	}

	if over <= 0 {
		return max(replaces, missing)
	}

	// Use deletes to reduce replaces
	// Priority: blocks where L%3==0 -> delete 1 reduces replaces by 1
	// Then: blocks where L%3==1 -> delete 2 reduces replaces by 1
	// Then: any block -> delete 3 reduces replaces by 1
	for i := range repeats {
		if over <= 0 {
			break
		}
		if repeats[i].length%3 == 0 {
			need := 1
			if over >= need {
				replaces -= repeats[i].length / 3
				repeats[i].length -= need
				replaces += repeats[i].length / 3
				over -= need
			}
		}
	}
	for i := range repeats {
		if over <= 0 {
			break
		}
		if repeats[i].length%3 == 1 {
			need := 2
			if over >= need {
				replaces -= repeats[i].length / 3
				repeats[i].length -= need
				replaces += repeats[i].length / 3
				over -= need
			}
		}
	}
	for i := range repeats {
		if over <= 0 {
			break
		}
		// Each 3 removes reduces replaces by 1
		need := repeats[i].length - 2 // we save replaces once length drops below 3
		if need <= 0 {
			// Already below 3 after previous deletes; no replace needed
			continue
		}
		if over >= need {
			replaces -= repeats[i].length / 3
			repeats[i].length -= need
			replaces += repeats[i].length / 3
			over -= need
		} else {
			// Partial: each group of 3 deletes eliminates 1 replace
			// Actually more nuanced: for any block with length >= 3, each 3 deletes
			// that keep length still >= 3 reduces replaces by 1
			replaces -= over / 3
			break
		}
	}

	return over + max(replaces, missing)
}
```

## 0425 — Word Squares

```go
package main

import "fmt"

// LeetCode #425: Word Squares
// https://leetcode.com/problems/word-squares/
// Difficulty: Hard
//
// Trie + backtracking. Build a trie of all words, then backtrack to form squares
// row by row. At each step, the next word must have a prefix matching the
// characters already placed in the current column.

func main() {
	// Example 1: ["area","lead","wall","lady","ball"] => 2 squares
	result := wordSquares([]string{"area", "lead", "wall", "lady", "ball"})
	fmt.Println("Word squares:", result)
	// Example 2: ["abat","baba","atan","atal"] => 2 squares
	result = wordSquares([]string{"abat", "baba", "atan", "atal"})
	fmt.Println("Word squares:", result)
	// Single word
	result = wordSquares([]string{"abc"})
	fmt.Println("Single:", result)
	// Edge: no solution
	result = wordSquares([]string{"a", "b"})
	fmt.Println("No solution:", result)
}

type trieNode struct {
	children [26]*trieNode
	words    []string
}

func wordSquares(words []string) [][]string {
	if len(words) == 0 {
		return nil
	}
	n := len(words[0])
	if n == 0 {
		return nil
	}

	root := &trieNode{}
	for _, w := range words {
		node := root
		root.words = append(root.words, w)
		for i := 0; i < len(w); i++ {
			idx := w[i] - 'a'
			if node.children[idx] == nil {
				node.children[idx] = &trieNode{}
			}
			node = node.children[idx]
			node.words = append(node.words, w)
		}
	}

	var result [][]string
	var current []string

	var backtrack func(step int)
	backtrack = func(step int) {
		if step == n {
			square := make([]string, n)
			copy(square, current)
			result = append(result, square)
			return
		}

		prefix := make([]byte, step)
		for i := 0; i < step; i++ {
			prefix[i] = current[i][step]
		}

		node := root
		for _, ch := range prefix {
			idx := ch - 'a'
			if node.children[idx] == nil {
				return
			}
			node = node.children[idx]
		}

		for _, candidate := range node.words {
			current = append(current, candidate)
			backtrack(step + 1)
			current = current[:len(current)-1]
		}
	}

	backtrack(0)
	return result
}
```

## 0428 — Serialize And Deserialize N Ary Tree

```go
package main

import (
	"fmt"
	"strconv"
	"strings"
)

// LeetCode #428: Serialize and Deserialize N-ary Tree
// https://leetcode.com/problems/serialize-and-deserialize-n-ary-tree/
// Difficulty: Hard
//
// BFS with child count. Serialize format: "val childCount val childCount ..."
// Deserialize uses a queue to rebuild the tree.

func main() {
	// Build tree: root = [1,null,3,2,4,null,5,6]
	// 1 has children 3,2,4; 3 has children 5,6
	root := &Node{Val: 1}
	n3 := &Node{Val: 3}
	n2 := &Node{Val: 2}
	n4 := &Node{Val: 4}
	n5 := &Node{Val: 5}
	n6 := &Node{Val: 6}
	root.Children = []*Node{n3, n2, n4}
	n3.Children = []*Node{n5, n6}

	codec := &Codec{}
	serialized := codec.serialize(root)
	fmt.Println("Serialized:", serialized)
	deserialized := codec.deserialize(serialized)
	fmt.Println("Root val:", deserialized.Val)
	fmt.Println("Children count:", len(deserialized.Children))

	// Empty tree
	empty := codec.serialize(nil)
	fmt.Println("Empty:", empty)
	fmt.Println("Empty deserialized:", codec.deserialize(empty))
}

// Node is an N-ary tree node.
type Node struct {
	Val      int
	Children []*Node
}

// Codec handles serialization/deserialization.
type Codec struct{}

func (c *Codec) serialize(root *Node) string {
	if root == nil {
		return ""
	}
	var parts []string
	queue := []*Node{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		parts = append(parts, strconv.Itoa(node.Val))
		parts = append(parts, strconv.Itoa(len(node.Children)))
		queue = append(queue, node.Children...)
	}
	return strings.Join(parts, " ")
}

func (c *Codec) deserialize(data string) *Node {
	if data == "" {
		return nil
	}
	tokens := strings.Split(data, " ")
	if len(tokens) == 0 {
		return nil
	}
	val, _ := strconv.Atoi(tokens[0])
	childCount, _ := strconv.Atoi(tokens[1])
	root := &Node{Val: val, Children: make([]*Node, 0, childCount)}

	type frame struct {
		node *Node
		need int
	}
	queue := []*frame{{node: root, need: childCount}}
	idx := 2
	for len(queue) > 0 {
		f := queue[0]
		queue = queue[1:]
		for i := 0; i < f.need; i++ {
			val, _ = strconv.Atoi(tokens[idx])
			childCount, _ = strconv.Atoi(tokens[idx+1])
			idx += 2
			child := &Node{Val: val, Children: make([]*Node, 0, childCount)}
			f.node.Children = append(f.node.Children, child)
			if childCount > 0 {
				queue = append(queue, &frame{node: child, need: childCount})
			}
		}
	}
	return root
}
```

## 0431 — Encode N Ary Tree To Binary Tree

```go
package main

import "fmt"

// LeetCode #431: Encode N-ary Tree to Binary Tree
// https://leetcode.com/problems/encode-n-ary-tree-to-binary-tree/
// Difficulty: Hard
//
// Left pointer = first child, Right pointer = next sibling.
// Encode: first child goes to Left, subsequent children chain via Right.
// Decode: collect Left and its Right chain as children.

func main() {
	// Build N-ary tree: 1 -> [3,2,4]; 3 -> [5,6]
	root := &NNode{Val: 1}
	n3 := &NNode{Val: 3}
	n2 := &NNode{Val: 2}
	n4 := &NNode{Val: 4}
	n5 := &NNode{Val: 5}
	n6 := &NNode{Val: 6}
	root.Children = []*NNode{n3, n2, n4}
	n3.Children = []*NNode{n5, n6}

	codec := &NaryCodec{}
	bt := codec.encode(root)
	fmt.Println("Encode root:", bt.Val)

	decoded := codec.decode(bt)
	fmt.Println("Decode root:", decoded.Val)
	fmt.Println("Decode children:", len(decoded.Children))

	// Nil tree
	fmt.Println("Encode nil:", codec.encode(nil))
	fmt.Println("Decode nil:", codec.decode(nil))

	// Leaf node
	leaf := &NNode{Val: 7}
	btLeaf := codec.encode(leaf)
	decodedLeaf := codec.decode(btLeaf)
	fmt.Println("Leaf decoded val:", decodedLeaf.Val)
	fmt.Println("Leaf children:", len(decodedLeaf.Children))
}

// NNode is an N-ary tree node.
type NNode struct {
	Val      int
	Children []*NNode
}

// TreeNode is a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// NaryCodec encodes/decodes N-ary <-> Binary trees.
type NaryCodec struct{}

func (c *NaryCodec) encode(root *NNode) *TreeNode {
	if root == nil {
		return nil
	}
	node := &TreeNode{Val: root.Val}
	if len(root.Children) > 0 {
		node.Left = c.encode(root.Children[0])
	}
	cur := node.Left
	for i := 1; i < len(root.Children); i++ {
		cur.Right = c.encode(root.Children[i])
		cur = cur.Right
	}
	return node
}

func (c *NaryCodec) decode(root *TreeNode) *NNode {
	if root == nil {
		return nil
	}
	node := &NNode{Val: root.Val}
	cur := root.Left
	for cur != nil {
		node.Children = append(node.Children, c.decode(cur))
		cur = cur.Right
	}
	if node.Children == nil {
		node.Children = []*NNode{}
	}
	return node
}
```

## 0432 — All Oone Data Structure

```go
package main

import "fmt"

// LeetCode #432: All O(1) Data Structure
// https://leetcode.com/problems/all-oone-data-structure/
// Difficulty: Hard
//
// Doubly linked list of frequency buckets. Each bucket holds keys with that count.
// inc, dec, getMaxKey, getMinKey all O(1).

func main() {
	allOne := Constructor()
	allOne.Inc("hello")
	allOne.Inc("hello")
	fmt.Println("MaxKey after 2x hello:", allOne.GetMaxKey())
	fmt.Println("MinKey after 2x hello:", allOne.GetMinKey())

	allOne.Inc("world")
	allOne.Inc("world")
	allOne.Inc("world")
	fmt.Println("MaxKey after 3x world:", allOne.GetMaxKey())

	allOne.Dec("world")
	allOne.Dec("world")
	fmt.Println("MaxKey after dec world 2x:", allOne.GetMaxKey())
	fmt.Println("MinKey after dec world 2x:", allOne.GetMinKey())

	allOne.Dec("hello")
	allOne.Dec("hello")
	fmt.Println("After dec hello to 0:")
	fmt.Println("MaxKey:", allOne.GetMaxKey())
	fmt.Println("MinKey:", allOne.GetMinKey())

	// Empty
	empty := Constructor()
	fmt.Println("Empty MaxKey:", empty.GetMaxKey())
	fmt.Println("Empty MinKey:", empty.GetMinKey())
}

type bucket struct {
	count int
	keys  map[string]bool
	prev  *bucket
	next  *bucket
}

type AllOne struct {
	head    *bucket
	tail    *bucket
	key2node map[string]*bucket
}

func Constructor() AllOne {
	head := &bucket{count: 0, keys: make(map[string]bool)}
	tail := &bucket{count: 0, keys: make(map[string]bool)}
	head.next = tail
	tail.prev = head
	return AllOne{head: head, tail: tail, key2node: make(map[string]*bucket)}
}

func (a *AllOne) Inc(key string) {
	cur, ok := a.key2node[key]
	if !ok {
		cur = a.head
	}
	nc := cur.count + 1
	if cur.next.count != nc {
		a.insertAfter(cur, &bucket{count: nc, keys: make(map[string]bool)})
	}
	nb := cur.next
	nb.keys[key] = true
	a.key2node[key] = nb

	if cur != a.head {
		delete(cur.keys, key)
		if len(cur.keys) == 0 {
			a.remove(cur)
		}
	}
}

func (a *AllOne) Dec(key string) {
	cur, ok := a.key2node[key]
	if !ok {
		return
	}
	nc := cur.count - 1
	if nc == 0 {
		delete(cur.keys, key)
		delete(a.key2node, key)
		if len(cur.keys) == 0 {
			a.remove(cur)
		}
		return
	}
	if cur.prev.count != nc {
		a.insertAfter(cur.prev, &bucket{count: nc, keys: make(map[string]bool)})
	}
	pb := cur.prev
	pb.keys[key] = true
	a.key2node[key] = pb

	delete(cur.keys, key)
	if len(cur.keys) == 0 {
		a.remove(cur)
	}
}

func (a *AllOne) GetMaxKey() string {
	if a.tail.prev == a.head {
		return ""
	}
	for k := range a.tail.prev.keys {
		return k
	}
	return ""
}

func (a *AllOne) GetMinKey() string {
	if a.head.next == a.tail {
		return ""
	}
	for k := range a.head.next.keys {
		return k
	}
	return ""
}

func (a *AllOne) insertAfter(prev *bucket, b *bucket) {
	b.prev = prev
	b.next = prev.next
	prev.next.prev = b
	prev.next = b
}

func (a *AllOne) remove(b *bucket) {
	b.prev.next = b.next
	b.next.prev = b.prev
}
```

## 0440 — K Th Smallest In Lexicographical Order

```go
package main

import "fmt"

// LeetCode #440: K-th Smallest in Lexicographical Order
// https://leetcode.com/problems/k-th-smallest-in-lexicographical-order/
// Difficulty: Hard
//
// Walk the lexicographic prefix tree. Count how many numbers exist under a
// given prefix in [1,n], then skip or descend. O(log^2 n) time, O(1) space.
// n=13, k=2 => 10

func main() {
	// Example 1
	fmt.Println("n=13, k=2 =>", findKthNumber(13, 2)) // 10
	// Example 2
	fmt.Println("n=1, k=1 =>", findKthNumber(1, 1)) // 1
	// Example 3
	fmt.Println("n=100, k=10 =>", findKthNumber(100, 10)) // 17
	// Larger
	fmt.Println("n=1000, k=100 =>", findKthNumber(1000, 100))
	// Edge: n=10, k=3 => 11? Let's see: [1,10,11,12,2,3,4,5,6,7,8,9] k=3 => 11
	fmt.Println("n=12, k=5 =>", findKthNumber(12, 5))
}

func findKthNumber(n int, k int) int {
	cur := 1
	k-- // convert to 0-indexed
	for k > 0 {
		steps := countSteps(n, cur, cur+1)
		if steps <= k {
			cur++
			k -= steps
		} else {
			cur *= 10
			k--
		}
	}
	return cur
}

// countSteps returns how many numbers in [1,n] are between `prefix` and
// `nextPrefix` when arranged in lexicographical order.
func countSteps(n int, prefix int, nextPrefix int) int {
	var steps int
	for prefix <= n {
		if nextPrefix <= n {
			steps += nextPrefix - prefix
		} else {
			steps += n - prefix + 1
		}
		prefix *= 10
		nextPrefix *= 10
	}
	return steps
}
```

## 0446 — Arithmetic Slices Ii Subsequence

```go
package main

import "fmt"

// LeetCode #446: Arithmetic Slices II - Subsequence
// https://leetcode.com/problems/arithmetic-slices-ii-subsequence/
// Difficulty: Hard
//
// DP with per-index difference maps. dp[i][diff] = number of arithmetic
// subsequences ending at i with common difference diff. For each pair (j,i),
// diff = nums[i]-nums[j]; the subsequences ending at i with that diff are
// dp[j][diff] (extend existing) + 1 (new pair). Sum all valid subsequences
// with length >= 3 (i.e. when dp[j][diff] >= 1).

func main() {
	// Example 1: [2,4,6,8,10] => 7
	fmt.Println("n=7:", numberOfArithmeticSlices([]int{2, 4, 6, 8, 10}))
	// Example 2: [7,7,7,7,7] => 16
	fmt.Println("n=16:", numberOfArithmeticSlices([]int{7, 7, 7, 7, 7}))
	// Example 3: [0,2000000000,-294967296] => 0
	fmt.Println("n=0:", numberOfArithmeticSlices([]int{0, 2000000000, -294967296}))
	// Edge: short array
	fmt.Println("n=0:", numberOfArithmeticSlices([]int{1, 2}))
}

func numberOfArithmeticSlices(nums []int) int {
	n := len(nums)
	if n < 3 {
		return 0
	}

	// dp[i] is a map from difference -> count of subsequences ending at i
	dp := make([]map[int]int, n)
	for i := range dp {
		dp[i] = make(map[int]int)
	}

	var total int
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			diff := nums[i] - nums[j]
			// count of subsequences ending at j with this diff
			count := dp[j][diff]
			// We add count to dp[i][diff] (extending existing subsequences)
			dp[i][diff] += count
			// Also add the pair (j,i) as a potential start of a new subsequence
			dp[i][diff]++

			// If count >= 1, then extending gives a valid subsequence of length >= 3
			if count >= 1 {
				total += count
			}
		}
	}
	return total
}
```

## 0458 — Poor Pigs

```go
package main

import "fmt"

// LeetCode #458: Poor Pigs
// https://leetcode.com/problems/poor-pigs/
// Difficulty: Hard
//
// Each pig has (minutesToTest/minutesToDie + 1) possible states:
// die at round 1, die at round 2, ..., survive all rounds.
// With P pigs, we can encode (states)^P bucket combinations.
// Find the smallest P such that states^P >= buckets.
// Uses iterative multiplication to avoid floating-point precision issues.

func main() {
	// Example 1: buckets=4, minToDie=15, minToTest=15 => states=2, 2^2 >= 4 => 2
	fmt.Println("Pigs:", poorPigs(4, 15, 15)) // 2
	// Example 2: buckets=1 => 0 pigs needed
	fmt.Println("Pigs:", poorPigs(1, 15, 15)) // 0
	// Example 3: buckets=1000, minToDie=15, minToTest=60 => states=5
	fmt.Println("Pigs:", poorPigs(1000, 15, 60)) // ceil(log5(1000)) = 5
	// Large exact power: states=5, 5^3=125 => 3
	fmt.Println("Pigs:", poorPigs(125, 1, 4)) // 3
	// Edge: 2 buckets, 1 round
	fmt.Println("Pigs:", poorPigs(2, 1, 1)) // 1
}

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	if buckets <= 1 {
		return 0
	}
	states := minutesToTest/minutesToDie + 1
	// Iterative approach: find smallest pigs such that states^pigs >= buckets
	pigs := 0
	covered := 1
	for covered < buckets {
		pigs++
		covered *= states
	}
	return pigs
}
```

## 0460 — Lfu Cache

```go
package main

import "fmt"

// LeetCode #460: LFU Cache
// https://leetcode.com/problems/lfu-cache/
// Difficulty: Hard
//
// Each frequency has a doubly linked LRU list. On access, move node to next
// frequency's list. On eviction, remove from the lowest frequency's LRU tail.
// get and put in O(1) amortized.

func main() {
	lfu := ConstructorLFU(2)
	lfu.Put(1, 1)
	lfu.Put(2, 2)
	fmt.Println("Get 1:", lfu.Get(1)) // 1
	lfu.Put(3, 3)                     // evicts key 2
	fmt.Println("Get 2:", lfu.Get(2)) // -1
	fmt.Println("Get 3:", lfu.Get(3)) // 3
	lfu.Put(4, 4)                     // evicts key 1
	fmt.Println("Get 1:", lfu.Get(1)) // -1
	fmt.Println("Get 3:", lfu.Get(3)) // 3
	fmt.Println("Get 4:", lfu.Get(4)) // 4

	// Single capacity
	lfu2 := ConstructorLFU(1)
	lfu2.Put(0, 0)
	fmt.Println("Get 0:", lfu2.Get(0)) // 0
	lfu2.Put(1, 1)
	fmt.Println("Get 0:", lfu2.Get(0)) // -1
	fmt.Println("Get 1:", lfu2.Get(1)) // 1
}

// LFUNode is a node in the LFU cache.
type LFUNode struct {
	key, value, freq int
	prev, next       *LFUNode
}

// FreqList is a doubly linked list for a specific frequency.
type FreqList struct {
	head, tail *LFUNode
}

func newFreqList() *FreqList {
	h := &LFUNode{}
	t := &LFUNode{}
	h.next = t
	t.prev = h
	return &FreqList{head: h, tail: t}
}

func (fl *FreqList) isEmpty() bool {
	return fl.head.next == fl.tail
}

// pushFront adds node to the front (most recently used).
func (fl *FreqList) pushFront(node *LFUNode) {
	node.prev = fl.head
	node.next = fl.head.next
	fl.head.next.prev = node
	fl.head.next = node
}

// remove removes a node from its current list.
func remove(node *LFUNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// popBack removes and returns the least recently used node.
func (fl *FreqList) popBack() *LFUNode {
	node := fl.tail.prev
	remove(node)
	return node
}

// LFUCache is the LFU cache.
type LFUCache struct {
	capacity int
	minFreq  int
	nodes    map[int]*LFUNode
	freqs    map[int]*FreqList
}

func ConstructorLFU(capacity int) LFUCache {
	return LFUCache{
		capacity: capacity,
		nodes:    make(map[int]*LFUNode),
		freqs:    make(map[int]*FreqList),
	}
}

func (c *LFUCache) Get(key int) int {
	node, ok := c.nodes[key]
	if !ok {
		return -1
	}
	c.incrementFreq(node)
	return node.value
}

func (c *LFUCache) Put(key int, value int) {
	if c.capacity == 0 {
		return
	}
	if node, ok := c.nodes[key]; ok {
		node.value = value
		c.incrementFreq(node)
		return
	}
	if len(c.nodes) >= c.capacity {
		c.evict()
	}
	node := &LFUNode{key: key, value: value, freq: 1}
	c.nodes[key] = node
	if c.freqs[1] == nil {
		c.freqs[1] = newFreqList()
	}
	c.freqs[1].pushFront(node)
	c.minFreq = 1
}

func (c *LFUCache) incrementFreq(node *LFUNode) {
	f := node.freq
	remove(node)
	if c.freqs[f].isEmpty() && f == c.minFreq {
		c.minFreq++
	}

	node.freq++
	if c.freqs[node.freq] == nil {
		c.freqs[node.freq] = newFreqList()
	}
	c.freqs[node.freq].pushFront(node)
}

func (c *LFUCache) evict() {
	list := c.freqs[c.minFreq]
	if list == nil || list.isEmpty() {
		return
	}
	node := list.popBack()
	delete(c.nodes, node.key)
}
```

## 0465 — Optimal Account Balancing

```go
package main

import "fmt"

// LeetCode #465: Optimal Account Balancing
// https://leetcode.com/problems/optimal-account-balancing/
// Difficulty: Hard
//
// Compute net balances, then find minimum transactions to settle debts using
// backtracking. For each non-zero balance, pair it with an opposite sign balance
// to settle. Use memoization on the bitmask of remaining non-zero balances.

func main() {
	// Example 1: [[0,1,10],[2,0,5]] => 2 transactions
	// Person 0 owes 10 to 1, person 2 owes 5 to 0.
	// Net: 0->5, 1->10, 2->-15, or better: 0:5, 1:10, 2:-15 => min 2
	fmt.Println("Min transactions:", minTransfers([][]int{{0, 1, 10}, {2, 0, 5}}))

	// Example 2: [[0,1,10],[1,0,1],[1,2,5],[2,0,5]] => 1
	fmt.Println("Min transactions:", minTransfers([][]int{{0, 1, 10}, {1, 0, 1}, {1, 2, 5}, {2, 0, 5}}))

	// Single transaction
	fmt.Println("Min transactions:", minTransfers([][]int{{0, 1, 100}}))
}

func minTransfers(transactions [][]int) int {
	balance := make(map[int]int)
	for _, t := range transactions {
		balance[t[0]] -= t[2]
		balance[t[1]] += t[2]
	}

	var debts []int
	for _, b := range balance {
		if b != 0 {
			debts = append(debts, b)
		}
	}
	if len(debts) == 0 {
		return 0
	}

	return backtrack(debts, 0)
}

func backtrack(debts []int, start int) int {
	// Skip settled debts
	for start < len(debts) && debts[start] == 0 {
		start++
	}
	if start == len(debts) {
		return 0
	}

	best := len(debts) // upper bound

	for i := start + 1; i < len(debts); i++ {
		if debts[i]*debts[start] < 0 {
			// Opposite sign: we can settle
			debts[i] += debts[start]
			best = min(best, 1+backtrack(debts, start+1))
			debts[i] -= debts[start] // backtrack
		}
	}
	return best
}
```

## 0466 — Count The Repetitions

```go
package main

import "fmt"

// LeetCode #466: Count The Repetitions
// https://leetcode.com/problems/count-the-repetitions/
// Difficulty: Hard
//
// Find the maximum m such that [s2, m] is a subsequence of [s1, n1].
// Greedy matching through repeated s1, counting how many full s2 sequences
// are matched as a subsequence.

func main() {
	// Example: s1="acb", n1=4 => "acbacbacbacb"
	// s2="ab", n2=2 => "abab"
	// "abab" is a subsequence of "acbacbacbacb" => m=2
	fmt.Println("m:", getMaxRepetitions("acb", 4, "ab", 2)) // 2

	// Example 2
	fmt.Println("m:", getMaxRepetitions("abc", 4, "ab", 2)) // 2

	// Repeated char
	fmt.Println("m:", getMaxRepetitions("aaa", 3, "aa", 1)) // 4

	// No match
	fmt.Println("m:", getMaxRepetitions("a", 1, "b", 1)) // 0
}

func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	if n1 == 0 {
		return 0
	}

	len1, len2 := len(s1), len(s2)
	totalChars := int64(n1) * int64(len1)

	matchCount := 0
	s2idx := 0

	for i := int64(0); i < totalChars; i++ {
		if s1[i%int64(len1)] == s2[s2idx] {
			s2idx++
			if s2idx == len2 {
				matchCount++
				s2idx = 0
			}
		}
	}
	return matchCount / n2
}
```

## 0471 — Encode String With Shortest Length

```go
package main

// LeetCode #471: Encode String with Shortest Length
// https://leetcode.com/problems/encode-string-with-shortest-length/
// Difficulty: Hard [Paid]
// Approach: DP interval encoding. dp[i][j] = shortest encoded form of s[i:j+1].
// For each substring, try to compress it using KMP-like period detection,
// or split it into two parts and combine.

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("471 - Encode String with Shortest Length")

	// Test cases
	fmt.Printf("encode(\"aaa\") = %q (expected: \"aaa\")\n", encode("aaa"))
	fmt.Printf("encode(\"aaaaa\") = %q (expected: \"5[a]\")\n", encode("aaaaa"))
	fmt.Printf("encode(\"aabcaabcd\") = %q (expected: \"2[aabc]d\")\n", encode("aabcaabcd"))
	fmt.Printf("encode(\"abbbabbbcabbbabbbc\") = %q (expected: \"2[2[abbb]c]\")\n", encode("abbbabbbcabbbabbbc"))
	fmt.Printf("encode(\"aaaaaaaaaa\") = %q (expected: \"10[a]\")\n", encode("aaaaaaaaaa"))
	fmt.Printf("encode(\"\") = %q (expected: \"\")\n", encode(""))
	fmt.Printf("encode(\"a\") = %q (expected: \"a\")\n", encode("a"))
	fmt.Printf("encode(\"ab\") = %q (expected: \"ab\")\n", encode("ab"))
	fmt.Printf("encode(\"abcabcabc\") = %q (expected: \"3[abc]\")\n", encode("abcabcabc"))
	fmt.Printf("encode(\"abbbabbbc\") = %q (expected: \"abbbabbbc\" (no compression -> \"a2[bbb]c\" vs original, original shorter))\n", encode("abbbabbbc"))
}

func encode(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}

	// dp[i][j] = shortest encoded form of s[i:j+1]
	dp := make([][]string, n)
	for i := range dp {
		dp[i] = make([]string, n)
		dp[i][i] = string(s[i])
	}

	// length from 2 to n
	for length := 2; length <= n; length++ {
		for i := 0; i+length-1 < n; i++ {
			j := i + length - 1
			substr := s[i : j+1]

			// Option 1: try to compress the whole substring
			// Find the smallest repeating pattern
			encoded := substr
			period := findPeriod(substr)
			if period > 0 && len(substr) > period {
				encoded = fmt.Sprintf("%d[%s]", len(substr)/period, dp[i][i+period-1])
			}

			// Option 2: split into two parts
			for k := i; k < j; k++ {
				combined := dp[i][k] + dp[k+1][j]
				if len(combined) < len(encoded) {
					encoded = combined
				}
			}

			// If the original is shorter, keep the original
			if len(substr) < len(encoded) {
				encoded = substr
			}

			dp[i][j] = encoded
		}
	}

	return dp[0][n-1]
}

// findPeriod returns the smallest period of the string.
// If the string can be formed by repeating a prefix, returns the length of that prefix.
// Otherwise returns 0.
func findPeriod(s string) int {
	n := len(s)
	// Try all possible periods from 1 to n/2
	for p := 1; p <= n/2; p++ {
		if n%p != 0 {
			continue
		}
		pattern := s[:p]
		if strings.Repeat(pattern, n/p) == s {
			return p
		}
	}
	return 0
}

```

## 0472 — Concatenated Words

```go
package main

// LeetCode #472: Concatenated Words
// https://leetcode.com/problems/concatenated-words/
// Difficulty: Hard
// Approach: DP word break. Sort words by length, for each word check if it
// can be formed by concatenating shorter words (already processed).

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("472 - Concatenated Words")

	// Test cases
	words1 := []string{"cat", "cats", "catsdogcats", "dog", "dogcatsdog", "hippopotamuses", "rat", "ratcatdogcat"}
	fmt.Printf("findAllConcatenatedWordsInADict(%v) = %v (expected: [catsdogcats dogcatsdog ratcatdogcat])\n",
		words1, findAllConcatenatedWordsInADict(words1))

	words2 := []string{"cat", "dog", "catdog"}
	fmt.Printf("findAllConcatenatedWordsInADict(%v) = %v (expected: [catdog])\n",
		words2, findAllConcatenatedWordsInADict(words2))

	words3 := []string{}
	fmt.Printf("findAllConcatenatedWordsInADict(%v) = %v (expected: [])\n",
		words3, findAllConcatenatedWordsInADict(words3))

	words4 := []string{"a"}
	fmt.Printf("findAllConcatenatedWordsInADict(%v) = %v (expected: [])\n",
		words4, findAllConcatenatedWordsInADict(words4))

	words5 := []string{"a", "aa", "aaa", "aaaa"}
	fmt.Printf("findAllConcatenatedWordsInADict(%v) = %v (expected: [aaa aaaa] or similar, depends on order)\n",
		words5, findAllConcatenatedWordsInADict(words5))

	words6 := []string{"", "a", "ab", "abc", "ababc"}
	fmt.Printf("findAllConcatenatedWordsInADict(%v) = %v\n",
		words6, findAllConcatenatedWordsInADict(words6))

	// Empty string should be handled - it's not a concatenated word
	// unless formed by 2+ empty strings, which conceptually doesn't count
}

func findAllConcatenatedWordsInADict(words []string) []string {
	// Sort words by length
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	wordSet := make(map[string]bool)
	result := []string{}

	for _, word := range words {
		if word == "" {
			continue
		}
		if canForm(word, wordSet) {
			result = append(result, word)
		}
		wordSet[word] = true
	}

	return result
}

// canForm checks if word can be formed by concatenating 2+ words from wordSet
func canForm(word string, wordSet map[string]bool) bool {
	if len(word) == 0 {
		return false
	}

	// dp[i] = can form word[0:i]
	dp := make([]bool, len(word)+1)
	dp[0] = true

	for i := 1; i <= len(word); i++ {
		for j := 0; j < i; j++ {
			if !dp[j] {
				continue
			}
			// Skip the case where we use the entire word itself
			if j == 0 && i == len(word) {
				// We need at least 2 words, so this single word match doesn't count
				continue
			}
			if wordSet[word[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(word)]
}
```

## 0479 — Largest Palindrome Product

```go
package main

// LeetCode #479: Largest Palindrome Product
// https://leetcode.com/problems/largest-palindrome-product/
// Difficulty: Hard
// Approach: Enumerate palindromes in descending order by building the first
// half and mirroring it. Check if the palindrome has a factor in the n-digit range.

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("479 - Largest Palindrome Product")

	for n := 1; n <= 8; n++ {
		fmt.Printf("n=%d -> %d\n", n, largestPalindrome(n))
	}
	// n=1->9, n=2->987, n=3->123, n=4->597, n=5->677, n=6->1218, n=7->877, n=8->475
}

func largestPalindrome(n int) int {
	if n == 1 {
		return 9
	}

	maxNum := int(math.Pow10(n)) - 1
	minNum := int(math.Pow10(n - 1))

	// Enumerate palindromes by iterating over the first half
	// Build palindrome: half + reverse(half) = an even-digit palindrome
	for half := maxNum; half >= 1; half-- {
		// Build the palindrome
		pal := int64(half)
		for temp := half; temp > 0; temp /= 10 {
			pal = pal*10 + int64(temp%10)
		}

		// Check if this palindrome is a product of two n-digit numbers
		// Only need to check factors up to sqrt(pal)
		for factor := int64(maxNum); factor*factor >= pal; factor-- {
			if pal%factor == 0 {
				other := pal / factor
				if other >= int64(minNum) && other <= int64(maxNum) {
					return int(pal % 1337)
				}
			}
		}
	}

	return 0
}
```

## 0480 — Sliding Window Median

```go
package main

// LeetCode #480: Sliding Window Median
// https://leetcode.com/problems/sliding-window-median/
// Difficulty: Hard
// Approach: Two heaps with lazy deletion (max-heap for left, min-heap for right).
// Maintain balance so left heap has either same count or one more than right heap.

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println("480 - Sliding Window Median")

	// Test cases
	nums1 := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k1 := 3
	fmt.Printf("nums=%v, k=%d -> %v (expected: [1 -1 -1 3 5 6])\n",
		nums1, k1, medianSlidingWindow(nums1, k1))

	nums2 := []int{1, 2, 3, 4, 5}
	k2 := 2
	fmt.Printf("nums=%v, k=%d -> %v (expected: [1.5 2.5 3.5 4.5])\n",
		nums2, k2, medianSlidingWindow(nums2, k2))

	nums3 := []int{1, 2}
	k3 := 1
	fmt.Printf("nums=%v, k=%d -> %v (expected: [1 2])\n",
		nums3, k3, medianSlidingWindow(nums3, k3))

	nums4 := []int{1, 1, 1, 1}
	k4 := 2
	fmt.Printf("nums=%v, k=%d -> %v (expected: [1 1 1])\n",
		nums4, k4, medianSlidingWindow(nums4, k4))

	nums5 := []int{2147483647, 2147483647}
	k5 := 2
	fmt.Printf("nums=%v, k=%d -> %v (expected: [2147483647])\n",
		nums5, k5, medianSlidingWindow(nums5, k5))

	nums6 := []int{5, 2, 3, 1, 4}
	k6 := 3
	fmt.Printf("nums=%v, k=%d -> %v\n",
		nums6, k6, medianSlidingWindow(nums6, k6))
}

// IntHeap for min-heap of int64 (to avoid overflow)
type IntHeap []int64

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int64)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// ReverseIntHeap for max-heap (by nesting IntHeap)
type ReverseIntHeap struct{ IntHeap }

func (h ReverseIntHeap) Less(i, j int) bool { return h.IntHeap[i] > h.IntHeap[j] }

// lazyHeap wraps a heap with a lazy deletion map
type lazyHeap struct {
	h      heap.Interface
	lazy   map[int64]int // value -> count to delete
	size   int
}

func newLazyHeap(isMax bool) *lazyHeap {
	var h heap.Interface
	if isMax {
		h = &ReverseIntHeap{}
	} else {
		h = &IntHeap{}
	}
	heap.Init(h)
	return &lazyHeap{h: h, lazy: make(map[int64]int)}
}

func (lh *lazyHeap) push(val int64) {
	heap.Push(lh.h, val)
	lh.size++
}

func (lh *lazyHeap) pop() int64 {
	lh.clean()
	lh.size--
	return heap.Pop(lh.h).(int64)
}

func (lh *lazyHeap) top() int64 {
	lh.clean()
	// We need to peek - let's access the underlying slice
	switch h := lh.h.(type) {
	case *IntHeap:
		return (*h)[0]
	case *ReverseIntHeap:
		return h.IntHeap[0]
	}
	return 0
}

func (lh *lazyHeap) lazyDelete(val int64) {
	lh.lazy[val]++
	lh.size--
}

func (lh *lazyHeap) clean() {
	for {
		var top int64
		switch h := lh.h.(type) {
		case *IntHeap:
			if len(*h) == 0 {
				return
			}
			top = (*h)[0]
		case *ReverseIntHeap:
			if len(h.IntHeap) == 0 {
				return
			}
			top = h.IntHeap[0]
		}

		if count, ok := lh.lazy[top]; ok && count > 0 {
			if count == 1 {
				delete(lh.lazy, top)
			} else {
				lh.lazy[top] = count - 1
			}
			heap.Pop(lh.h)
		} else {
			break
		}
	}
}

func (lh *lazyHeap) len() int {
	return lh.size
}

func medianSlidingWindow(nums []int, k int) []float64 {
	n := len(nums)
	if k == 0 || n == 0 {
		return nil
	}

	result := make([]float64, 0, n-k+1)

	left := newLazyHeap(true)  // max-heap (smaller half)
	right := newLazyHeap(false) // min-heap (larger half)

	// Initialize with first k elements
	for i := 0; i < k; i++ {
		addNum(int64(nums[i]), left, right)
	}

	result = append(result, findMedian(left, right, k))

	for i := k; i < n; i++ {
		// Remove nums[i-k]
		removeNum(int64(nums[i-k]), left, right)
		// Add nums[i]
		addNum(int64(nums[i]), left, right)
		result = append(result, findMedian(left, right, k))
	}

	return result
}

func addNum(val int64, left, right *lazyHeap) {
	if left.len() == 0 || val <= left.top() {
		left.push(val)
	} else {
		right.push(val)
	}

	// Rebalance
	if left.len() > right.len()+1 {
		right.push(left.pop())
	} else if right.len() > left.len() {
		left.push(right.pop())
	}
}

func removeNum(val int64, left, right *lazyHeap) {
	// Determine which heap the value is in and mark for lazy deletion
	if left.len() > 0 && val <= left.top() {
		left.lazyDelete(val)
	} else {
		right.lazyDelete(val)
	}

	// Rebalance
	if left.len() > right.len()+1 {
		right.push(left.pop())
	} else if right.len() > left.len() {
		left.push(right.pop())
	}
}

func findMedian(left, right *lazyHeap, k int) float64 {
	if k%2 == 1 {
		return float64(left.top())
	}
	return float64(left.top()+right.top()) / 2.0
}
```

## 0483 — Smallest Good Base

```go
package main

// LeetCode #483: Smallest Good Base
// https://leetcode.com/problems/smallest-good-base/
// Difficulty: Hard
// Approach: Binary search per length. For a number n represented as string,
// find the smallest base k such that n = 1 + k + k^2 + ... + k^(m-1).
// The base k is smallest when m (number of digits) is largest.
// Max m is about log2(n)+1, min m is 2.

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	fmt.Println("483 - Smallest Good Base")

	// Test cases
	testCases := []string{"13", "4681", "1000000000000000000", "3", "15", "1", "2251799813685247"}
	for _, tc := range testCases {
		result := smallestGoodBase(tc)
		fmt.Printf("n=%s -> base=%s\n", tc, result)
	}
	// Expected: "13"->"3", "4681"->"8", "3"->"2", "15"->"2"
	// "1"->"" (or "1"? LeetCode says: For n=1, return "1" (since 1 in any base > 1 is just "1"))
}

func smallestGoodBase(n string) string {
	num, _ := new(big.Int).SetString(n, 10)

	// Compare with 1 (big.Int)
	one := big.NewInt(1)
	if num.Cmp(one) == 0 {
		return "1"
	}

	// n = 1 + k + k^2 + ... + k^(m-1)
	// The max possible m is when k=2: n >= 1 + 2 + ... + 2^(m-1) = 2^m - 1
	// so m <= log2(n+1) ~ 60 for n up to 10^18
	maxM := num.BitLen() + 1 // log2(n) + 1

	// Try each m from max down to 2 (larger m -> smaller base k)
	for m := maxM; m >= 2; m-- {
		// Binary search for k
		low := big.NewInt(2)
		high := new(big.Int).Sub(num, big.NewInt(1))

		for low.Cmp(high) <= 0 {
			mid := new(big.Int).Add(low, high)
			mid.Div(mid, big.NewInt(2))

			// Compute sum = 1 + mid + mid^2 + ... + mid^(m-1)
			sum := geometricSum(mid, m)

			cmp := sum.Cmp(num)
			if cmp == 0 {
				return mid.String()
			} else if cmp < 0 {
				low.Add(mid, big.NewInt(1))
			} else {
				high.Sub(mid, big.NewInt(1))
			}
		}
	}

	return strconv.FormatInt(num.Int64()-1, 10)
}

// geometricSum returns sum_{i=0}^{m-1} base^i
func geometricSum(base *big.Int, m int) *big.Int {
	// Use Horner's method: (((1*base + 1)*base + 1)*base + 1)
	// Actually: 1 + base*(1 + base*(1 + base*(...)))
	if m == 0 {
		return big.NewInt(0)
	}

	// Compute using the formula (base^m - 1) / (base - 1)
	pow := new(big.Int).Exp(base, big.NewInt(int64(m)), nil)
	pow.Sub(pow, big.NewInt(1))

	baseMinus1 := new(big.Int).Sub(base, big.NewInt(1))
	return pow.Div(pow, baseMinus1)
}
```

## 0488 — Zuma Game

```go
package main

// LeetCode #488: Zuma Game
// https://leetcode.com/problems/zuma-game/
// Difficulty: Hard
// Approach: DFS + memoization. Try each hand ball at each position in the board.
// After placing, remove consecutive 3+ balls and continue recursively.

import (
	"fmt"
)

func main() {
	fmt.Println("488 - Zuma Game")

	// Test cases
	fmt.Printf("findMinStep(\"WRRBBW\", \"RB\") = %d (expected: -1)\n",
		findMinStep("WRRBBW", "RB"))
	fmt.Printf("findMinStep(\"WWRRBBWW\", \"WRBRW\") = %d (expected: 2)\n",
		findMinStep("WWRRBBWW", "WRBRW"))
	fmt.Printf("findMinStep(\"G\", \"GGGGG\") = %d (expected: 2)\n",
		findMinStep("G", "GGGGG"))
	fmt.Printf("findMinStep(\"RBYYBBRRB\", \"YRBGB\") = %d (expected: 3)\n",
		findMinStep("RBYYBBRRB", "YRBGB"))
	fmt.Printf("findMinStep(\"\", \"\") = %d (expected: 0)\n",
		findMinStep("", ""))
	fmt.Printf("findMinStep(\"RRWWRRBBRR\", \"WB\") = %d\n",
		findMinStep("RRWWRRBBRR", "WB"))
}

func findMinStep(board string, hand string) int {
	// Count hand balls
	handCount := make([]int, 26)
	for _, ch := range hand {
		handCount[ch-'A']++
	}

	result := dfs(board, handCount)
	if result > len(hand) {
		return -1
	}
	return result
}

func dfs(board string, handCount []int) int {
	if board == "" {
		return 0
	}

	// Prune: if there are balls in hand that don't exist on board, skip
	// and also if count is insufficient

	// Try every possible placement
	minUsed := len(handCount)*5 + 1 // larger than any possible answer

	for i := 0; i < len(board); i++ {
		// Try to insert a ball from hand
		for color := 0; color < 26; color++ {
			if handCount[color] == 0 {
				continue
			}
			// Check if this placement makes sense: the ball color should match
			// a neighbor, or there should be at least 2 of the same color
			ball := byte('A' + color)
			if board[i] == ball {
				// Insert ball at position i
				handCount[color]--
				newBoard := removeConsecutive(board[:i] + string(ball) + board[i:])
				used := 1 + dfs(newBoard, handCount)
				if used < minUsed {
					minUsed = used
				}
				handCount[color]++
			} else if i > 0 && i < len(board)-1 && board[i-1] == board[i+1] && board[i-1] == ball {
				// Special case: inserting between two same balls to form 3
				handCount[color]--
				newBoard := removeConsecutive(board[:i] + string(ball) + board[i:])
				used := 1 + dfs(newBoard, handCount)
				if used < minUsed {
					minUsed = used
				}
				handCount[color]++
			}
		}

		// Skip consecutive same characters to avoid duplicate placements
		for i+1 < len(board) && board[i] == board[i+1] {
			i++
		}
	}

	return minUsed
}

// removeConsecutive removes all groups of 3+ consecutive same characters
func removeConsecutive(s string) string {
	// Keep removing until no more groups of 3+
	for {
		newS := removeOnce(s)
		if newS == s {
			break
		}
		s = newS
	}
	return s
}

func removeOnce(s string) string {
	if len(s) < 3 {
		return s
	}

	// Find groups of 3+ consecutive same characters
	result := make([]byte, 0, len(s))
	i := 0
	for i < len(s) {
		j := i
		for j < len(s) && s[j] == s[i] {
			j++
		}
		if j-i < 3 {
			// Keep this group
			result = append(result, s[i:j]...)
		}
		// If >=3, skip (remove them)
		i = j
	}
	return string(result)
}
```

## 0489 — Robot Room Cleaner

```go
package main

// LeetCode #489: Robot Room Cleaner
// https://leetcode.com/problems/robot-room-cleaner/
// Difficulty: Hard [Paid]
// Approach: DFS backtracking with simulated robot API.
// The robot has 4 methods: move(), turnLeft(), turnRight(), clean().
// We explore the room using DFS, keeping track of visited cells.
// Directions: 0=up, 1=right, 2=down, 3=left

import (
	"fmt"
)

func main() {
	fmt.Println("489 - Robot Room Cleaner")

	// Test with a simulated room
	room := [][]int{
		{1, 1, 1, 1, 1, 0, 1, 1},
		{1, 1, 1, 1, 1, 0, 1, 1},
		{1, 0, 1, 1, 1, 1, 1, 1},
		{0, 0, 0, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 1},
	}
	row, col := 1, 3

	robot := NewSimulatedRobot(room, row, col)
	cleanRoom(robot)

	cleanedCount := 0
	for _, r := range robot.sim.cleaned {
		for _, cell := range r {
			if cell {
				cleanedCount++
			}
		}
	}
	fmt.Println("Cleaning complete. Cleaned:", cleanedCount)
	fmt.Println("Total cleaned cells in room:")
	count := 0
	for _, r := range room {
		for _, cell := range r {
			if cell == 1 {
				count++
			}
		}
	}
	fmt.Printf("Expected to clean: %d cells\n", count)
}

// Robot interface as provided by LeetCode
type Robot struct {
	// This would be the LeetCode API
	// For our simulation, we expose the methods
	sim *simulatedRobot
}

func (r *Robot) Move() bool     { return r.sim.move() }
func (r *Robot) TurnLeft()      { r.sim.turnLeft() }
func (r *Robot) TurnRight()     { r.sim.turnRight() }
func (r *Robot) Clean()         { r.sim.clean() }

// Direction vectors: 0=up, 1=right, 2=down, 3=left
var dirs = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

// cleanRoomClient is the actual solution function
func cleanRoomClient(robot *Robot) {
	visited := make(map[[2]int]bool)
	backtrack(robot, 0, 0, 0, visited) // start at (0, 0), facing up
}

func backtrack(robot *Robot, r, c, dir int, visited map[[2]int]bool) {
	key := [2]int{r, c}
	if visited[key] {
		return
	}
	visited[key] = true
	robot.Clean()

	// Try all 4 directions
	for i := 0; i < 4; i++ {
		newDir := (dir + i) % 4
		nr := r + dirs[newDir][0]
		nc := c + dirs[newDir][1]

		if robot.Move() {
			backtrack(robot, nr, nc, newDir, visited)
			// Go back
			robot.TurnLeft()
			robot.TurnLeft()
			robot.Move()
			robot.TurnLeft()
			robot.TurnLeft()
		} else {
			// Can't move, turn right to try next direction
			robot.TurnRight()
		}
	}

	// Turn back to original direction before returning
	// Actually, since we return to original position, the direction
	// should be restored. Let's adjust: after exploring all 4 directions,
	// we should be back at original direction.
	// Each Move() returns us, and unsuccessful moves just TurnRight.
	// At the end, we've turned right 4 times = full circle = back to original.
}

// For LeetCode submission, the function is normally:
// func cleanRoom(robot *Robot) { ... }
// We'll use this as the main function
func cleanRoom(robot *Robot) {
	visited := make(map[[2]int]bool)
	dfsClean(robot, 0, 0, 0, visited)
}

func dfsClean(robot *Robot, r, c, dir int, visited map[[2]int]bool) {
	key := [2]int{r, c}
	if visited[key] {
		return
	}
	visited[key] = true
	robot.Clean()

	// Try all 4 directions in order
	for i := 0; i < 4; i++ {
		newDir := (dir + i) % 4
		nr := r + dirs[newDir][0]
		nc := c + dirs[newDir][1]

		if robot.Move() {
			dfsClean(robot, nr, nc, newDir, visited)
			// Backtrack: turn 180, move, turn 180
			robot.TurnLeft()
			robot.TurnLeft()
			robot.Move()
			// Restore direction: turn right and then left = 180 total, but
			// we need to go back to original dir. Let's see:
			// Before backtrack, robot is at new cell facing newDir.
			// After TurnLeft+TurnLeft, faces (newDir+2)%4.
			// After Move(), back at (r,c) facing (newDir+2)%4.
			// We need to restore to original dir.
			// Turn right twice to go from (newDir+2)%4 to dir.
			robot.TurnRight()
			robot.TurnRight()
		}
		// Turn right to try next direction
		robot.TurnRight()
	}

	// After loop: we have turned right 4 times = back to original dir
}

// ---- Simulation for testing ----

type simulatedRobot struct {
	room    [][]int
	r, c    int
	dir     int // 0=up, 1=right, 2=down, 3=left
	visited map[[2]int]bool
	cleaned [][]bool
}

func NewSimulatedRobot(room [][]int, r, c int) *Robot {
	cleaned := make([][]bool, len(room))
	for i := range cleaned {
		cleaned[i] = make([]bool, len(room[i]))
	}
	return &Robot{
		sim: &simulatedRobot{
			room:    room,
			r:       r,
			c:       c,
			dir:     0, // start facing up
			visited: make(map[[2]int]bool),
			cleaned: cleaned,
		},
	}
}

func (s *simulatedRobot) move() bool {
	nr := s.r + dirs[s.dir][0]
	nc := s.c + dirs[s.dir][1]

	// Check bounds and walls
	if nr < 0 || nr >= len(s.room) || nc < 0 || nc >= len(s.room[0]) {
		return false
	}
	if s.room[nr][nc] == 0 {
		return false
	}

	s.r = nr
	s.c = nc
	s.visited[[2]int{nr, nc}] = true
	return true
}

func (s *simulatedRobot) turnLeft() {
	s.dir = (s.dir + 3) % 4
}

func (s *simulatedRobot) turnRight() {
	s.dir = (s.dir + 1) % 4
}

func (s *simulatedRobot) clean() {
	if s.r >= 0 && s.r < len(s.room) && s.c >= 0 && s.c < len(s.room[0]) {
		s.cleaned[s.r][s.c] = true
		s.visited[[2]int{s.r, s.c}] = true
	}
}
```

## 0493 — Reverse Pairs

```go
package main

// LeetCode #493: Reverse Pairs
// https://leetcode.com/problems/reverse-pairs/
// Difficulty: Hard
// Approach: Merge sort counting. During merge, count pairs (i, j) where
// i < j and nums[i] > 2 * nums[j]. This is similar to counting inversions
// but with the 2x multiplier.

import (
	"fmt"
)

func main() {
	fmt.Println("493 - Reverse Pairs")

	// Test cases
	fmt.Printf("reversePairs([1,3,2,3,1]) = %d (expected: 2)\n",
		reversePairs([]int{1, 3, 2, 3, 1}))
	fmt.Printf("reversePairs([2,4,3,5,1]) = %d (expected: 3)\n",
		reversePairs([]int{2, 4, 3, 5, 1}))
	fmt.Printf("reversePairs([]) = %d (expected: 0)\n",
		reversePairs([]int{}))
	fmt.Printf("reversePairs([1]) = %d (expected: 0)\n",
		reversePairs([]int{1}))
	fmt.Printf("reversePairs([1,1,1,1]) = %d (expected: 0)\n",
		reversePairs([]int{1, 1, 1, 1}))
	fmt.Printf("reversePairs([5,4,3,2,1]) = %d (expected: 4)\n",
		reversePairs([]int{5, 4, 3, 2, 1}))
	fmt.Printf("reversePairs([2147483647, 2147483647, -2147483648, -2147483648]) = %d\n",
		reversePairs([]int{2147483647, 2147483647, -2147483648, -2147483648}))
	fmt.Printf("reversePairs([1,2,3,4,5]) = %d (expected: 0)\n",
		reversePairs([]int{1, 2, 3, 4, 5}))
}

func reversePairs(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	temp := make([]int, len(nums))
	count := mergeSort(nums, temp, 0, len(nums)-1)
	return count
}

func mergeSort(nums []int, temp []int, left, right int) int {
	if left >= right {
		return 0
	}

	mid := left + (right-left)/2
	count := mergeSort(nums, temp, left, mid)
	count += mergeSort(nums, temp, mid+1, right)

	// Count reverse pairs across two halves
	count += countPairs(nums, left, mid, right)

	// Merge the two sorted halves
	merge(nums, temp, left, mid, right)

	return count
}

// countPairs counts how many reverse pairs exist between [left..mid] and [mid+1..right]
// Both halves are sorted.
func countPairs(nums []int, left, mid, right int) int {
	count := 0
	j := mid + 1

	for i := left; i <= mid; i++ {
		// For each nums[i], find first nums[j] where nums[i] <= 2*nums[j]
		// Use int64 to avoid overflow
		val := int64(nums[i])
		for j <= right && val > 2*int64(nums[j]) {
			j++
		}
		count += j - (mid + 1)
	}

	return count
}

func merge(nums []int, temp []int, left, mid, right int) {
	i, j, k := left, mid+1, left

	for i <= mid && j <= right {
		if nums[i] <= nums[j] {
			temp[k] = nums[i]
			i++
		} else {
			temp[k] = nums[j]
			j++
		}
		k++
	}

	for i <= mid {
		temp[k] = nums[i]
		i++
		k++
	}

	for j <= right {
		temp[k] = nums[j]
		j++
		k++
	}

	for idx := left; idx <= right; idx++ {
		nums[idx] = temp[idx]
	}
}
```

## 0499 — The Maze Iii

```go
package main

// LeetCode #499: The Maze III
// https://leetcode.com/problems/the-maze-iii/
// Difficulty: Hard [Paid]
// Approach: Dijkstra with lexicographic path. The ball rolls until it hits a wall.
// We use a priority queue to find the shortest path. If multiple paths have the
// same length, choose the lexicographically smaller path string.

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	fmt.Println("499 - The Maze III")

	// Test cases
	maze1 := [][]int{
		{0, 0, 0, 0, 0},
		{1, 1, 0, 0, 1},
		{0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1},
		{0, 1, 0, 0, 0},
	}
	ball1 := []int{4, 3}
	hole1 := []int{0, 1}
	fmt.Printf("findShortestWay(maze1, ball=%v, hole=%v) = %q (expected: \"lul\")\n",
		ball1, hole1, findShortestWay(maze1, ball1, hole1))

	maze2 := [][]int{
		{0, 0, 0, 0, 0},
		{1, 1, 0, 0, 1},
		{0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1},
		{0, 1, 0, 0, 0},
	}
	ball2 := []int{4, 3}
	hole2 := []int{3, 0}
	fmt.Printf("findShortestWay(maze2, ball=%v, hole=%v) = %q (expected: \"impossible\")\n",
		ball2, hole2, findShortestWay(maze2, ball2, hole2))

	maze3 := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	ball3 := []int{0, 0}
	hole3 := []int{2, 2}
	fmt.Printf("findShortestWay(maze3, ball=%v, hole=%v) = %q\n",
		ball3, hole3, findShortestWay(maze3, ball3, hole3))

	maze4 := [][]int{
		{0, 0},
		{0, 0},
	}
	ball4 := []int{0, 0}
	hole4 := []int{0, 1}
	fmt.Printf("findShortestWay(maze4, ball=%v, hole=%v) = %q (expected: \"r\")\n",
		ball4, hole4, findShortestWay(maze4, ball4, hole4))
}

// Directions: down, up, right, left (to match expected lexicographic order)
// "d", "l", "r", "u" - but we need sorted lexicographically: "d" < "l" < "r" < "u"
// Let's use: d, l, r, u

type Dir struct {
	dr, dc int
	ch     byte
}

var dirs = []Dir{
	{1, 0, 'd'},  // down
	{-1, 0, 'u'}, // up
	{0, 1, 'r'},  // right
	{0, -1, 'l'}, // left
}

// State represents (row, col) with distance and path
type State struct {
	r, c int
	dist int
	path string
	idx   int // for heap
}

// Priority queue
type PQ []*State

func (pq PQ) Len() int { return len(pq) }

func (pq PQ) Less(i, j int) bool {
	if pq[i].dist != pq[j].dist {
		return pq[i].dist < pq[j].dist
	}
	return pq[i].path < pq[j].path
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].idx = i
	pq[j].idx = j
}

func (pq *PQ) Push(x interface{}) {
	n := len(*pq)
	state := x.(*State)
	state.idx = n
	*pq = append(*pq, state)
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	state := old[n-1]
	state.idx = -1
	*pq = old[0 : n-1]
	return state
}

func findShortestWay(maze [][]int, ball []int, hole []int) string {
	m, n := len(maze), len(maze[0])
	startR, startC := ball[0], ball[1]
	holeR, holeC := hole[0], hole[1]

	// dist[r][c] = minimum distance to reach (r,c)
	dist := make([][]int, m)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = math.MaxInt32
		}
	}

	// path[r][c] = lexicographically smallest path to reach (r,c)
	path := make([][]string, m)
	for i := range path {
		path[i] = make([]string, n)
	}

	pq := &PQ{}
	heap.Init(pq)

	dist[startR][startC] = 0
	path[startR][startC] = ""
	heap.Push(pq, &State{r: startR, c: startC, dist: 0, path: ""})

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(*State)

		// Skip if we already found a better path to this cell
		if cur.dist > dist[cur.r][cur.c] {
			continue
		}
		if cur.path != path[cur.r][cur.c] {
			continue
		}

		// Try all 4 directions
		for _, d := range dirs {
			// Roll the ball in this direction until it hits a wall or goes out of bounds
			nr, nc := cur.r, cur.c
			steps := 0

			for {
				// Check if we hit the hole during rolling
				if nr == holeR && nc == holeC {
					break
				}

				nextR := nr + d.dr
				nextC := nc + d.dc

				// Stop if next position is out of bounds or is a wall
				if nextR < 0 || nextR >= m || nextC < 0 || nextC >= n {
					break
				}
				if maze[nextR][nextC] == 1 {
					break
				}

				nr = nextR
				nc = nextC
				steps++
			}

			if steps == 0 {
				continue
			}

			newDist := cur.dist + steps
			newPath := cur.path + string(d.ch)

			// Check if this is a better path to (nr, nc)
			if newDist < dist[nr][nc] || (newDist == dist[nr][nc] && newPath < path[nr][nc]) {
				dist[nr][nc] = newDist
				path[nr][nc] = newPath
				heap.Push(pq, &State{r: nr, c: nc, dist: newDist, path: newPath})
			}
		}
	}

	if dist[holeR][holeC] == math.MaxInt32 {
		return "impossible"
	}
	return path[holeR][holeC]
}
```

## 0502 — Ipo

```go
package main

// LeetCode #502: IPO
// https://leetcode.com/problems/ipo/
// Difficulty: Hard
// Approach: Two heaps. Sort projects by capital, use a max-heap for profits
// of affordable projects. Iterate k times, each time add all projects whose
// capital <= current w to the max-heap, then pick the most profitable one.

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	fmt.Println("502 - IPO")

	// Test cases
	k1 := 2
	w1 := 0
	profits1 := []int{1, 2, 3}
	capital1 := []int{0, 1, 1}
	fmt.Printf("k=%d, w=%d, profits=%v, capital=%v -> %d (expected: 4)\n",
		k1, w1, profits1, capital1, findMaximizedCapital(k1, w1, profits1, capital1))

	k2 := 3
	w2 := 0
	profits2 := []int{1, 2, 3}
	capital2 := []int{0, 1, 2}
	fmt.Printf("k=%d, w=%d, profits=%v, capital=%v -> %d (expected: 6)\n",
		k2, w2, profits2, capital2, findMaximizedCapital(k2, w2, profits2, capital2))

	k3 := 1
	w3 := 2
	profits3 := []int{1, 2, 3}
	capital3 := []int{1, 1, 2}
	fmt.Printf("k=%d, w=%d, profits=%v, capital=%v -> %d (expected: 5)\n",
		k3, w3, profits3, capital3, findMaximizedCapital(k3, w3, profits3, capital3))

	k4 := 0
	w4 := 0
	profits4 := []int{1, 2, 3}
	capital4 := []int{0, 1, 1}
	fmt.Printf("k=%d, w=%d, profits=%v, capital=%v -> %d (expected: 0)\n",
		k4, w4, profits4, capital4, findMaximizedCapital(k4, w4, profits4, capital4))

	// With w=0, no project is affordable since min capital=1 > 0
	k5 := 10
	w5 := 0
	profits5 := []int{1, 2, 3, 4, 5}
	capital5 := []int{5, 4, 3, 2, 1}
	fmt.Printf("k=%d, w=%d, profits=%v, capital=%v -> %d (expected: 0, none affordable)\n",
		k5, w5, profits5, capital5, findMaximizedCapital(k5, w5, profits5, capital5))

	// LeetCode example
	k6 := 2
	w6 := 0
	profits6 := []int{1, 2, 3}
	capital6 := []int{0, 9, 10}
	fmt.Printf("k=%d, w=%d, profits=%v, capital=%v -> %d (expected: 1)\n",
		k6, w6, profits6, capital6, findMaximizedCapital(k6, w6, profits6, capital6))
}

// Max-heap for profits
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	n := len(profits)

	// Create project pairs (capital, profit) and sort by capital
	projects := make([][2]int, n)
	for i := 0; i < n; i++ {
		projects[i] = [2]int{capital[i], profits[i]}
	}
	sort.Slice(projects, func(i, j int) bool {
		return projects[i][0] < projects[j][0]
	})

	// Max-heap for profits of available projects
	profitHeap := &MaxHeap{}
	heap.Init(profitHeap)

	idx := 0
	for i := 0; i < k; i++ {
		// Add all projects we can afford now
		for idx < n && projects[idx][0] <= w {
			heap.Push(profitHeap, projects[idx][1])
			idx++
		}

		// If no projects are available, we're done
		if profitHeap.Len() == 0 {
			break
		}

		// Pick the most profitable project
		w += heap.Pop(profitHeap).(int)
	}

	return w
}
```

## 0514 — Freedom Trail

```go
package main

// LeetCode #514: Freedom Trail
// https://leetcode.com/problems/freedom-trail/
// Difficulty: Hard

import "fmt"

func main() {
	fmt.Println(findRotateSteps("godding", "gd")) // Expected: 4
}

func findRotateSteps(ring string, key string) int {
	m, n := len(ring), len(key)
	// pos[c] = list of indices in ring where character c appears
	pos := make([][]int, 26)
	for i := 0; i < m; i++ {
		c := ring[i] - 'a'
		pos[c] = append(pos[c], i)
	}

	// dp[j] = min steps to spell up to current key char ending at ring index j
	dp := make([]int, m)
	for j := 0; j < m; j++ {
		if ring[j] == key[0] {
			dp[j] = minDist(j, 0, m) + 1 // steps to rotate + press
		} else {
			dp[j] = 1 << 30 // large number
		}
	}

	for i := 1; i < n; i++ {
		next := make([]int, m)
		for j := 0; j < m; j++ {
			next[j] = 1 << 30
		}
		for _, j := range pos[key[i]-'a'] {
			// from any previous position where we could have been
			for _, k := range pos[key[i-1]-'a'] {
				cost := dp[k] + minDist(j, k, m) + 1
				if cost < next[j] {
					next[j] = cost
				}
			}
		}
		dp = next
	}

	ans := 1 << 30
	for j := 0; j < m; j++ {
		if dp[j] < ans {
			ans = dp[j]
		}
	}
	return ans
}

func minDist(i, j, m int) int {
	d := i - j
	if d < 0 {
		d = -d
	}
	if d > m-d {
		return m - d
	}
	return d
}
```

## 0517 — Super Washing Machines

```go
package main

// LeetCode #517: Super Washing Machines
// https://leetcode.com/problems/super-washing-machines/
// Difficulty: Hard

import "fmt"

func main() {
	fmt.Println(findMinMoves([]int{1, 0, 5})) // Expected: 3
}

func findMinMoves(machines []int) int {
	n := len(machines)
	sum := 0
	for _, v := range machines {
		sum += v
	}
	if sum%n != 0 {
		return -1
	}
	target := sum / n

	ans := 0
	balance := 0
	for _, v := range machines {
		balance += v - target
		if balance > ans {
			ans = balance
		}
		if balance < -ans {
			ans = -balance
		}
		// A machine may need to receive from both sides simultaneously
		if v-target > ans {
			ans = v - target
		}
	}
	return ans
}
```

## 0527 — Word Abbreviation

```go
package main

// LeetCode #527: Word Abbreviation
// https://leetcode.com/problems/word-abbreviation/
// Difficulty: Hard

import (
	"fmt"
	"strconv"
)

func main() {
	input := []string{"like", "god", "internal", "me", "internet", "interval", "intension", "face", "intrusion"}
	output := wordsAbbreviation(input)
	fmt.Println(output)
	// Expected: ["l2e","god","internal","me","i6t","interval","inte4n","f4e","intr4n"]
}

func wordsAbbreviation(words []string) []string {
	n := len(words)
	ans := make([]string, n)
	prefix := make([]int, n)

	for i := 0; i < n; i++ {
		prefix[i] = 1
		ans[i] = abbreviate(words[i], 1)
	}

	for i := 0; i < n; i++ {
		for {
			conflict := false
			for j := i + 1; j < n; j++ {
				if ans[i] == ans[j] {
					prefix[j]++
					ans[j] = abbreviate(words[j], prefix[j])
					conflict = true
				}
			}
			if conflict {
				prefix[i]++
				ans[i] = abbreviate(words[i], prefix[i])
			} else {
				break
			}
		}
	}
	return ans
}

func abbreviate(s string, k int) string {
	if k >= len(s)-2 {
		return s
	}
	abbr := s[:k] + strconv.Itoa(len(s)-k-1) + s[len(s)-1:]
	if len(abbr) >= len(s) {
		return s
	}
	return abbr
}
```

## 0546 — Remove Boxes

```go
package main

// LeetCode #546: Remove Boxes
// https://leetcode.com/problems/remove-boxes/
// Difficulty: Hard

import "fmt"

func main() {
	fmt.Println(removeBoxes([]int{1, 3, 2, 2, 2, 3, 4, 3, 1})) // Expected: 23
}

func removeBoxes(boxes []int) int {
	n := len(boxes)
	// dp[l][r][k] = max points for boxes[l..r] with k extra boxes of same color as boxes[l] attached to the left
	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, n)
		}
	}
	return dfs(boxes, 0, n-1, 0, dp)
}

func dfs(boxes []int, l, r, k int, dp [][][]int) int {
	if l > r {
		return 0
	}
	if dp[l][r][k] > 0 {
		return dp[l][r][k]
	}

	// compress consecutive same colors
	origL := l
	origK := k
	for l+1 <= r && boxes[l+1] == boxes[l] {
		l++
		k++
	}

	// option 1: remove boxes[l..l] plus the k attached ones
	res := (k+1)*(k+1) + dfs(boxes, l+1, r, 0, dp)

	// option 2: merge with later same-color boxes
	for m := l + 1; m <= r; m++ {
		if boxes[m] == boxes[l] {
			res = max(res, dfs(boxes, l+1, m-1, 0, dp)+dfs(boxes, m, r, k+1, dp))
		}
	}

	dp[origL][r][origK] = res
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

## 0548 — Split Array With Equal Sum

```go
package main

// LeetCode #548: Split Array with Equal Sum
// https://leetcode.com/problems/split-array-with-equal-sum/
// Difficulty: Hard

import "fmt"

func main() {
	fmt.Println(splitArray([]int{1, 2, 1, 2, 1, 2, 1})) // Expected: true
}

func splitArray(nums []int) bool {
	n := len(nums)
	if n < 7 {
		return false
	}

	prefix := make([]int, n)
	prefix[0] = nums[0]
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] + nums[i]
	}

	// Try position j (second split, 0-indexed)
	for j := 3; j <= n-4; j++ {
		set := make(map[int]bool)
		// Try position i (first split)
		for i := 1; i < j-1; i++ {
			sum1 := prefix[i-1]
			sum2 := prefix[j-1] - prefix[i]
			if sum1 == sum2 {
				set[sum1] = true
			}
		}
		// Try position k (third split)
		for k := j + 2; k < n-1; k++ {
			sum3 := prefix[k-1] - prefix[j]
			sum4 := prefix[n-1] - prefix[k]
			if sum3 == sum4 && set[sum3] {
				return true
			}
		}
	}
	return false
}
```

## 0552 — Student Attendance Record Ii

```go
package main

// LeetCode #552: Student Attendance Record II
// https://leetcode.com/problems/student-attendance-record-ii/
// Difficulty: Hard

import "fmt"

const mod = 1_000_000_007

func main() {
	fmt.Println(checkRecord(2))     // Expected: 8
	fmt.Println(checkRecord(10101)) // Expected: 183236316
}

func checkRecord(n int) int {
	// States:
	// dp[i][a][l] where:
	//   i = length
	//   a = number of absences (0 or 1)
	//   l = consecutive lates (0, 1, or 2)
	//
	// We use 3D DP. dp[a][l] = count for current length.

	prev := [2][3]int{}
	prev[0][0] = 1 // "" -> empty string

	for i := 0; i < n; i++ {
		cur := [2][3]int{}
		for a := 0; a <= 1; a++ {
			for l := 0; l <= 2; l++ {
				if prev[a][l] == 0 {
					continue
				}
				// Append 'P' (present) — resets lates
				cur[a][0] = (cur[a][0] + prev[a][l]) % mod
				// Append 'A' (absent)
				if a < 1 {
					cur[a+1][0] = (cur[a+1][0] + prev[a][l]) % mod
				}
				// Append 'L' (late)
				if l < 2 {
					cur[a][l+1] = (cur[a][l+1] + prev[a][l]) % mod
				}
			}
		}
		prev = cur
	}

	ans := 0
	for a := 0; a <= 1; a++ {
		for l := 0; l <= 2; l++ {
			ans = (ans + prev[a][l]) % mod
		}
	}
	return ans
}
```

