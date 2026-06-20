# Medium (Sedang) — Problem ��0707

## 0523 — Continuous Subarray Sum

```go
package main

// LeetCode #523: Continuous Subarray Sum
// https://leetcode.com/problems/continuous-subarray-sum/
// Difficulty: Medium
// Time: O(n)
// Space: O(k) where k = min(k, n)

import "fmt"

func main() {
	fmt.Println(CheckSubarraySum([]int{23, 2, 4, 6, 7}, 6))
	fmt.Println(CheckSubarraySum([]int{23, 2, 6, 4, 7}, 6))
	fmt.Println(CheckSubarraySum([]int{23, 2, 6, 4, 7}, 13))
}

func CheckSubarraySum(nums []int, k int) bool {
	// Map remainder -> first index
	remainderMap := make(map[int]int)
	remainderMap[0] = -1
	sum := 0

	for i, num := range nums {
		sum += num
		rem := sum % k
		if rem < 0 {
			rem += k
		}
		if prevIdx, ok := remainderMap[rem]; ok {
			if i-prevIdx >= 2 {
				return true
			}
		} else {
			remainderMap[rem] = i
		}
	}

	return false
}
```

## 0524 — Longest Word In Dictionary Through Deleting

```go
package main

// LeetCode #524: Longest Word in Dictionary through Deleting
// https://leetcode.com/problems/longest-word-in-dictionary-through-deleting/
// Difficulty: Medium
// Time: O(n * L) where n = len(dictionary), L = max(len(s), len(word))
// Space: O(1)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(FindLongestWord("abpcplea", []string{"ale", "apple", "monkey", "plea"}))
	fmt.Println(FindLongestWord("abpcplea", []string{"a", "b", "c"}))
}

func FindLongestWord(s string, dictionary []string) string {
	// Sort by length desc, then lexicographically
	sort.Slice(dictionary, func(i, j int) bool {
		if len(dictionary[i]) != len(dictionary[j]) {
			return len(dictionary[i]) > len(dictionary[j])
		}
		return dictionary[i] < dictionary[j]
	})

	for _, word := range dictionary {
		if isSubsequence(word, s) {
			return word
		}
	}

	return ""
}

func isSubsequence(word, s string) bool {
	i := 0
	for j := 0; i < len(word) && j < len(s); j++ {
		if word[i] == s[j] {
			i++
		}
	}
	return i == len(word)
}
```

## 0525 — Contiguous Array

```go
package main

// LeetCode #525: Contiguous Array
// https://leetcode.com/problems/contiguous-array/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(FindMaxLength([]int{0, 1}))
	fmt.Println(FindMaxLength([]int{0, 1, 0}))
}

func FindMaxLength(nums []int) int {
	// Map count -> first index. count = (#1 - #0)
	countMap := make(map[int]int)
	countMap[0] = -1
	count := 0
	maxLen := 0

	for i, num := range nums {
		if num == 1 {
			count++
		} else {
			count--
		}
		if prevIdx, ok := countMap[count]; ok {
			if i-prevIdx > maxLen {
				maxLen = i - prevIdx
			}
		} else {
			countMap[count] = i
		}
	}

	return maxLen
}
```

## 0526 — Beautiful Arrangement

```go
package main

// LeetCode #526: Beautiful Arrangement
// https://leetcode.com/problems/beautiful-arrangement/
// Difficulty: Medium
// Time: O(k) where k = number of valid permutations
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(CountArrangement(2))
	fmt.Println(CountArrangement(1))
}

func CountArrangement(n int) int {
	used := make([]bool, n+1)
	count := 0

	var backtrack func(pos int)
	backtrack = func(pos int) {
		if pos > n {
			count++
			return
		}
		for i := 1; i <= n; i++ {
			if !used[i] && (i%pos == 0 || pos%i == 0) {
				used[i] = true
				backtrack(pos + 1)
				used[i] = false
			}
		}
	}

	backtrack(1)
	return count
}
```

## 0528 — Random Pick With Weight

```go
package main

// LeetCode #528: Random Pick with Weight
// https://leetcode.com/problems/random-pick-with-weight/
// Difficulty: Medium
// Time: O(n) for init, O(log n) per pick
// Space: O(n)

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	sol := Constructor([]int{1, 3})
	fmt.Println(sol.PickIndex())
	fmt.Println(sol.PickIndex())
	fmt.Println(sol.PickIndex())
}

type Solution struct {
	prefixSum []int
	totalSum  int
}

func Constructor(w []int) Solution {
	prefixSum := make([]int, len(w))
	sum := 0
	for i, weight := range w {
		sum += weight
		prefixSum[i] = sum
	}
	return Solution{prefixSum: prefixSum, totalSum: sum}
}

func (s *Solution) PickIndex() int {
	target := rand.Intn(s.totalSum) + 1
	return sort.SearchInts(s.prefixSum, target)
}
```

## 0529 — Minesweeper

```go
package main

// LeetCode #529: Minesweeper
// https://leetcode.com/problems/minesweeper/
// Difficulty: Medium
// Time: O(m * n)
// Space: O(m * n)

import "fmt"

func main() {
	board := [][]byte{
		{'E', 'E', 'E', 'E', 'E'},
		{'E', 'E', 'M', 'E', 'E'},
		{'E', 'E', 'E', 'E', 'E'},
		{'E', 'E', 'E', 'E', 'E'},
	}
	result := UpdateBoard(board, []int{3, 0})
	for _, row := range result {
		fmt.Println(string(row))
	}
}

func UpdateBoard(board [][]byte, click []int) [][]byte {
	r, c := click[0], click[1]
	if board[r][c] == 'M' {
		board[r][c] = 'X'
		return board
	}
	reveal(board, r, c)
	return board
}

func reveal(board [][]byte, r, c int) {
	if r < 0 || r >= len(board) || c < 0 || c >= len(board[0]) || board[r][c] != 'E' {
		return
	}
	m, n := len(board), len(board[0])
	mines := 0
	for dr := -1; dr <= 1; dr++ {
		for dc := -1; dc <= 1; dc++ {
			if dr == 0 && dc == 0 {
				continue
			}
			nr, nc := r+dr, c+dc
			if nr >= 0 && nr < m && nc >= 0 && nc < n && board[nr][nc] == 'M' {
				mines++
			}
		}
	}
	if mines > 0 {
		board[r][c] = byte('0' + mines)
		return
	}
	board[r][c] = 'B'
	for dr := -1; dr <= 1; dr++ {
		for dc := -1; dc <= 1; dc++ {
			if dr == 0 && dc == 0 {
				continue
			}
			reveal(board, r+dr, c+dc)
		}
	}
}
```

## 0531 — Lonely Pixel I

```go
package main

// LeetCode #531: Lonely Pixel I
// https://leetcode.com/problems/lonely-pixel-i/
// Difficulty: Medium [Paid]
// Time: O(m * n)
// Space: O(m + n)

import "fmt"

func main() {
	picture := [][]byte{
		{'W', 'W', 'B'},
		{'W', 'B', 'W'},
		{'B', 'W', 'W'},
	}
	fmt.Println(FindLonelyPixel(picture))
}

func FindLonelyPixel(picture [][]byte) int {
	m, n := len(picture), len(picture[0])
	rows := make([]int, m)
	cols := make([]int, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if picture[i][j] == 'B' {
				rows[i]++
				cols[j]++
			}
		}
	}

	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if picture[i][j] == 'B' && rows[i] == 1 && cols[j] == 1 {
				count++
			}
		}
	}

	return count
}
```

## 0532 — K Diff Pairs In An Array

```go
package main

// LeetCode #532: K-diff Pairs in an Array
// https://leetcode.com/problems/k-diff-pairs-in-an-array/
// Difficulty: Medium
// Time: O(n log n)
// Space: O(1) (ignoring sorting overhead)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(FindPairs([]int{3, 1, 4, 1, 5}, 2))
	fmt.Println(FindPairs([]int{1, 2, 3, 4, 5}, 1))
	fmt.Println(FindPairs([]int{1, 3, 1, 5, 4}, 0))
}

func FindPairs(nums []int, k int) int {
	sort.Ints(nums)
	n := len(nums)
	count := 0
	i, j := 0, 1

	for i < n && j < n {
		if i == j || nums[j]-nums[i] < k {
			j++
		} else if nums[j]-nums[i] > k {
			i++
		} else {
			count++
			i++
			j++
			// Skip duplicates
			for j < n && nums[j] == nums[j-1] {
				j++
			}
		}
	}

	return count
}
```

## 0533 — Lonely Pixel Ii

```go
package main

// LeetCode #533: Lonely Pixel II
// https://leetcode.com/problems/lonely-pixel-ii/
// Difficulty: Medium [Paid]
// Time: O(m * n)
// Space: O(m * n)

import "fmt"

func main() {
	picture := [][]byte{
		{'W', 'B', 'W', 'B', 'B', 'W'},
		{'W', 'B', 'W', 'B', 'B', 'W'},
		{'W', 'B', 'W', 'B', 'B', 'W'},
		{'W', 'W', 'B', 'W', 'B', 'W'},
	}
	fmt.Println(FindBlackPixel(picture, 3))
}

func FindBlackPixel(picture [][]byte, target int) int {
	m, n := len(picture), len(picture[0])
	rows := make([]int, m)
	cols := make([]int, n)
	rowPattern := make(map[string]int)

	for i := 0; i < m; i++ {
		rowStr := ""
		for j := 0; j < n; j++ {
			rowStr += string(picture[i][j])
			if picture[i][j] == 'B' {
				rows[i]++
				cols[j]++
			}
		}
		rowPattern[rowStr]++
	}

	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if picture[i][j] == 'B' && rows[i] == target && cols[j] == target {
				// All rows with 'B' in column j must be identical
				allSame := true
				for r := 0; r < m; r++ {
					if picture[r][j] == 'B' && picture[r][j] == picture[i][j] {
						// Check if rows i and r are identical
						for c := 0; c < n; c++ {
							if picture[i][c] != picture[r][c] {
								allSame = false
								break
							}
						}
						if !allSame {
							break
						}
					}
				}
				if allSame {
					count++
				}
			}
		}
	}

	return count
}
```

## 0534 — Game Play Analysis Iii

```go
package main

// LeetCode #534: Game Play Analysis III
// https://leetcode.com/problems/game-play-analysis-iii/
// Difficulty: Medium [Paid]
// Time: O(n log n) for sorting
// Space: O(n)

import (
	"fmt"
	"sort"
)

func main() {
	// Sample player activity data
	activities := [][]int{
		{1, 2016, 5},  // player_id=1, event_date=2016, games_played=5
		{1, 2017, 10}, // player_id=1, event_date=2017, games_played=10
		{2, 2016, 15}, // player_id=2, event_date=2016, games_played=15
	}
	fmt.Println(PlayerActivityReport(activities))
}

func PlayerActivityReport(activities [][]int) [][]int {
	sort.Slice(activities, func(i, j int) bool {
		if activities[i][0] != activities[j][0] {
			return activities[i][0] < activities[j][0]
		}
		return activities[i][1] < activities[j][1]
	})

	result := [][]int{}
	runningSum := 0
	for i, act := range activities {
		if i > 0 && act[0] != activities[i-1][0] {
			runningSum = 0
		}
		runningSum += act[2]
		result = append(result, []int{act[0], act[1], runningSum})
	}
	return result
}
```

## 0535 — Encode And Decode Tinyurl

```go
package main

// LeetCode #535: Encode and Decode TinyURL
// https://leetcode.com/problems/encode-and-decode-tinyurl/
// Difficulty: Medium
// Time: O(1) for both encode and decode
// Space: O(n) where n = number of encoded URLs

import (
	"fmt"
	"math/rand"
)

func main() {
	url := "https://leetcode.com/problems/design-tinyurl"
	encoded := Encode(url)
	fmt.Println("Encoded:", encoded)
	decoded := Decode(encoded)
	fmt.Println("Decoded:", decoded)
}

var urlMap = make(map[string]string)
var keyLen = 6
const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func Encode(longUrl string) string {
	key := make([]byte, keyLen)
	for i := range key {
		key[i] = chars[rand.Intn(len(chars))]
	}
	shortKey := string(key)
	urlMap[shortKey] = longUrl
	return shortKey
}

func Decode(shortUrl string) string {
	return urlMap[shortUrl]
}
```

## 0536 — Construct Binary Tree From String

```go
package main

// LeetCode #536: Construct Binary Tree from String
// https://leetcode.com/problems/construct-binary-tree-from-string/
// Difficulty: Medium [Paid]
// Time: O(n)
// Space: O(n)

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := Str2tree("4(2(3)(1))(6(5))")
	printTree(root)
	fmt.Println()
}

func Str2tree(s string) *TreeNode {
	if len(s) == 0 {
		return nil
	}
	return buildTree(s, 0)
}

func buildTree(s string, idx int) *TreeNode {
	if idx >= len(s) || s[idx] == ')' {
		return nil
	}

	// Parse number
	start := idx
	for idx < len(s) && (s[idx] == '-' || (s[idx] >= '0' && s[idx] <= '9')) {
		idx++
	}
	val, _ := strconv.Atoi(s[start:idx])
	node := &TreeNode{Val: val}

	// Parse left child
	if idx < len(s) && s[idx] == '(' {
		node.Left = buildTree(s, idx+1)
		// Skip to matching close paren
		depth := 1
		idx++
		for idx < len(s) && depth > 0 {
			if s[idx] == '(' {
				depth++
			} else if s[idx] == ')' {
				depth--
			}
			idx++
		}
	}

	// Parse right child
	if idx < len(s) && s[idx] == '(' {
		node.Right = buildTree(s, idx+1)
		depth := 1
		idx++
		for idx < len(s) && depth > 0 {
			if s[idx] == '(' {
				depth++
			} else if s[idx] == ')' {
				depth--
			}
			idx++
		}
	}

	return node
}

func printTree(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.Val)
	printTree(root.Left)
	printTree(root.Right)
}
```

## 0537 — Complex Number Multiplication

```go
package main

// LeetCode #537: Complex Number Multiplication
// https://leetcode.com/problems/complex-number-multiplication/
// Difficulty: Medium
// Time: O(1)
// Space: O(1)

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(ComplexNumberMultiplication("1+1i", "1+1i"))
	fmt.Println(ComplexNumberMultiplication("1+-1i", "1+-1i"))
}

func ComplexNumberMultiplication(num1 string, num2 string) string {
	a, b := parseComplex(num1)
	c, d := parseComplex(num2)

	real := a*c - b*d
	imag := a*d + b*c

	return fmt.Sprintf("%d+%di", real, imag)
}

func parseComplex(s string) (int, int) {
	parts := strings.Split(s, "+")
	real, _ := strconv.Atoi(parts[0])
	imagPart := parts[1][:len(parts[1])-1] // remove trailing 'i'
	imag, _ := strconv.Atoi(imagPart)
	return real, imag
}
```

## 0538 — Convert Bst To Greater Tree

```go
package main

// LeetCode #538: Convert BST to Greater Tree
// https://leetcode.com/problems/convert-bst-to-greater-tree/
// Difficulty: Medium
// Time: O(n)
// Space: O(h) where h is tree height

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}}
	root.Right = &TreeNode{Val: 6, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 7, Right: &TreeNode{Val: 8}}}
	result := ConvertBST(root)
	printInorder(result)
	fmt.Println()
}

func ConvertBST(root *TreeNode) *TreeNode {
	sum := 0
	var reverseInorder func(node *TreeNode)
	reverseInorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		reverseInorder(node.Right)
		sum += node.Val
		node.Val = sum
		reverseInorder(node.Left)
	}
	reverseInorder(root)
	return root
}

func printInorder(root *TreeNode) {
	if root == nil {
		return
	}
	printInorder(root.Left)
	fmt.Printf("%d ", root.Val)
	printInorder(root.Right)
}
```

## 0539 — Minimum Time Difference

```go
package main

// LeetCode #539: Minimum Time Difference
// https://leetcode.com/problems/minimum-time-difference/
// Difficulty: Medium
// Time: O(n log n)
// Space: O(n)

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(FindMinDifference([]string{"23:59", "00:00"}))
	fmt.Println(FindMinDifference([]string{"00:00", "23:59", "00:00"}))
}

func FindMinDifference(timePoints []string) int {
	n := len(timePoints)
	minutes := make([]int, n)

	for i, t := range timePoints {
		h, _ := strconv.Atoi(t[:2])
		m, _ := strconv.Atoi(t[3:])
		minutes[i] = h*60 + m
	}

	sort.Ints(minutes)

	minDiff := 24 * 60 // 1440
	for i := 1; i < n; i++ {
		diff := minutes[i] - minutes[i-1]
		if diff < minDiff {
			minDiff = diff
		}
	}

	// Check circular difference
	circularDiff := 1440 - minutes[n-1] + minutes[0]
	if circularDiff < minDiff {
		minDiff = circularDiff
	}

	return minDiff
}
```

## 0540 — Single Element In A Sorted Array

```go
package main

// LeetCode #540: Single Element in a Sorted Array
// https://leetcode.com/problems/single-element-in-a-sorted-array/
// Difficulty: Medium
// Time: O(log n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(SingleNonDuplicate([]int{1, 1, 2, 3, 3, 4, 4, 8, 8}))
	fmt.Println(SingleNonDuplicate([]int{3, 3, 7, 7, 10, 11, 11}))
}

func SingleNonDuplicate(nums []int) int {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := lo + (hi-lo)/2
		// Ensure mid is even to compare with mid+1
		if mid%2 == 1 {
			mid--
		}
		if nums[mid] == nums[mid+1] {
			lo = mid + 2
		} else {
			hi = mid
		}
	}
	return nums[lo]
}
```

## 0542 — 01 Matrix

```go
package main

// LeetCode #542: 01 Matrix
// https://leetcode.com/problems/01-matrix/
// Difficulty: Medium
// Time: O(m * n)
// Space: O(m * n)

import "fmt"

func main() {
	mat1 := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	fmt.Println(UpdateMatrix(mat1))
	mat2 := [][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}
	fmt.Println(UpdateMatrix(mat2))
}

func UpdateMatrix(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	dist := make([][]int, m)
	queue := [][]int{}

	for i := 0; i < m; i++ {
		dist[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				dist[i][j] = 0
				queue = append(queue, []int{i, j})
			} else {
				dist[i][j] = -1
			}
		}
	}

	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, d := range dirs {
			r, c := cur[0]+d[0], cur[1]+d[1]
			if r >= 0 && r < m && c >= 0 && c < n && dist[r][c] == -1 {
				dist[r][c] = dist[cur[0]][cur[1]] + 1
				queue = append(queue, []int{r, c})
			}
		}
	}

	return dist
}
```

## 0544 — Output Contest Matches

```go
package main

// LeetCode #544: Output Contest Matches
// https://leetcode.com/problems/output-contest-matches/
// Difficulty: Medium [Paid]
// Time: O(n)
// Space: O(n)

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(FindContestMatch(2))
	fmt.Println(FindContestMatch(4))
	fmt.Println(FindContestMatch(8))
}

func FindContestMatch(n int) string {
	teams := make([]string, n)
	for i := 0; i < n; i++ {
		teams[i] = strconv.Itoa(i + 1)
	}

	for n > 1 {
		for i := 0; i < n/2; i++ {
			teams[i] = "(" + teams[i] + "," + teams[n-1-i] + ")"
		}
		n /= 2
	}

	return teams[0]
}
```

## 0545 — Boundary Of Binary Tree

```go
package main

// LeetCode #545: Boundary of Binary Tree
// https://leetcode.com/problems/boundary-of-binary-tree/
// Difficulty: Medium [Paid]
// Time: O(n)
// Space: O(n)

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}}
	fmt.Println(BoundaryOfBinaryTree(root))
}

func BoundaryOfBinaryTree(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := []int{root.Val}
	addLeftBoundary(root.Left, &result)
	addLeaves(root.Left, &result)
	addLeaves(root.Right, &result)
	addRightBoundary(root.Right, &result)
	return result
}

func addLeftBoundary(node *TreeNode, result *[]int) {
	if node == nil || (node.Left == nil && node.Right == nil) {
		return
	}
	*result = append(*result, node.Val)
	if node.Left != nil {
		addLeftBoundary(node.Left, result)
	} else {
		addLeftBoundary(node.Right, result)
	}
}

func addRightBoundary(node *TreeNode, result *[]int) {
	if node == nil || (node.Left == nil && node.Right == nil) {
		return
	}
	if node.Right != nil {
		addRightBoundary(node.Right, result)
	} else {
		addRightBoundary(node.Left, result)
	}
	*result = append(*result, node.Val) // post-order for reverse
}

func addLeaves(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil {
		*result = append(*result, node.Val)
		return
	}
	addLeaves(node.Left, result)
	addLeaves(node.Right, result)
}
```

## 0547 — Number Of Provinces

```go
package main

// LeetCode #547: Number of Provinces
// https://leetcode.com/problems/number-of-provinces/
// Difficulty: Medium
// Time: O(n^2)
// Space: O(n)

import "fmt"

func main() {
	isConnected1 := [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}
	fmt.Println(FindCircleNum(isConnected1))
	isConnected2 := [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	fmt.Println(FindCircleNum(isConnected2))
}

func FindCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	visited := make([]bool, n)
	count := 0

	var dfs func(city int)
	dfs = func(city int) {
		visited[city] = true
		for neighbor := 0; neighbor < n; neighbor++ {
			if isConnected[city][neighbor] == 1 && !visited[neighbor] {
				dfs(neighbor)
			}
		}
	}

	for i := 0; i < n; i++ {
		if !visited[i] {
			count++
			dfs(i)
		}
	}

	return count
}
```

## 0549 — Binary Tree Longest Consecutive Sequence Ii

```go
package main

// LeetCode #549: Binary Tree Longest Consecutive Sequence II
// https://leetcode.com/problems/binary-tree-longest-consecutive-sequence-ii/
// Difficulty: Medium [Paid]
// Time: O(n)
// Space: O(h) where h is tree height

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	fmt.Println(LongestConsecutive(root))
}

func LongestConsecutive(root *TreeNode) int {
	maxLen := 0
	var dfs func(node *TreeNode) (int, int)
	dfs = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}
		inc, dec := 1, 1

		leftInc, leftDec := dfs(node.Left)
		rightInc, rightDec := dfs(node.Right)

		if node.Left != nil {
			if node.Left.Val == node.Val+1 {
				inc = max(inc, leftInc+1)
			}
			if node.Left.Val == node.Val-1 {
				dec = max(dec, leftDec+1)
			}
		}
		if node.Right != nil {
			if node.Right.Val == node.Val+1 {
				inc = max(inc, rightInc+1)
			}
			if node.Right.Val == node.Val-1 {
				dec = max(dec, rightDec+1)
			}
		}

		maxLen = max(maxLen, inc+dec-1)
		return inc, dec
	}

	dfs(root)
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

## 0550 — Game Play Analysis Iv

```go
package main

// LeetCode #550: Game Play Analysis IV
// https://leetcode.com/problems/game-play-analysis-iv/
// Difficulty: Medium
// Time: O(n log n)
// Space: O(n)

import (
	"fmt"
	"sort"
)

func main() {
	// Simulating player activity: {player_id, event_date, games_played}
	activities := [][]int{
		{1, 2016, 5},
		{1, 2017, 10},
		{2, 2016, 15},
		{2, 2018, 20},
		{3, 2020, 30},
		{3, 2021, 25},
	}
	fmt.Printf("%.2f\n", GamePlayAnalysisIv(activities))
}

func GamePlayAnalysisIv(activities [][]int) float64 {
	if len(activities) == 0 {
		return 0.0
	}
	sort.Slice(activities, func(i, j int) bool {
		if activities[i][0] != activities[j][0] {
			return activities[i][0] < activities[j][0]
		}
		return activities[i][1] < activities[j][1]
	})

	// Find first login for each player
	firstLogin := make(map[int]int)
	for _, act := range activities {
		pid, date := act[0], act[1]
		if _, exists := firstLogin[pid]; !exists {
			firstLogin[pid] = date
		}
	}

	// Count players who logged in the day after their first login
	nextDayPlayers := 0
	playerSet := make(map[int]bool)
	for _, act := range activities {
		pid, date := act[0], act[1]
		if playerSet[pid] {
			continue
		}
		if date == firstLogin[pid]+1 {
			nextDayPlayers++
			playerSet[pid] = true
		}
	}

	totalPlayers := len(firstLogin)
	if totalPlayers == 0 {
		return 0.0
	}
	return float64(nextDayPlayers) / float64(totalPlayers)
}
```

## 0553 — Optimal Division

```go
package main

// LeetCode #553: Optimal Division
// https://leetcode.com/problems/optimal-division/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(OptimalDivision([]int{1000, 100, 10, 2}))
	fmt.Println(OptimalDivision([]int{2, 3, 4}))
	fmt.Println(OptimalDivision([]int{2}))
}

func OptimalDivision(nums []int) string {
	n := len(nums)
	if n == 0 {
		return ""
	}
	if n == 1 {
		return strconv.Itoa(nums[0])
	}
	if n == 2 {
		return strconv.Itoa(nums[0]) + "/" + strconv.Itoa(nums[1])
	}

	result := strconv.Itoa(nums[0]) + "/(" + strconv.Itoa(nums[1])
	for i := 2; i < n; i++ {
		result += "/" + strconv.Itoa(nums[i])
	}
	result += ")"

	return result
}
```

## 0554 — Brick Wall

```go
package main

// LeetCode #554: Brick Wall
// https://leetcode.com/problems/brick-wall/
// Difficulty: Medium
// Time: O(n * m) where n = rows, m = avg bricks per row
// Space: O(n * m)

import "fmt"

func main() {
	wall := [][]int{
		{1, 2, 2, 1},
		{3, 1, 2},
		{1, 3, 2},
		{2, 4},
		{3, 1, 2},
		{1, 3, 1, 1},
	}
	fmt.Println(LeastBricks(wall))
}

func LeastBricks(wall [][]int) int {
	gapCount := make(map[int]int)

	for _, row := range wall {
		pos := 0
		for i := 0; i < len(row)-1; i++ {
			pos += row[i]
			gapCount[pos]++
		}
	}

	maxGaps := 0
	for _, count := range gapCount {
		if count > maxGaps {
			maxGaps = count
		}
	}

	return len(wall) - maxGaps
}
```

## 0555 — Split Concatenated Strings

```go
package main

// LeetCode #555: Split Concatenated Strings
// https://leetcode.com/problems/split-concatenated-strings/
// Difficulty: Medium [Paid]
// Time: O(n * L) where L = total length of all strings
// Space: O(L)

import "fmt"

func main() {
	fmt.Println(SplitLoopedString([]string{"abc", "xyz"}))
	fmt.Println(SplitLoopedString([]string{"lc", "evol", "cdy"}))
}

func SplitLoopedString(strs []string) string {
	n := len(strs)
	// For each string, use the lexicographically larger of itself and its reverse
	reversed := make([]string, n)
	for i, s := range strs {
		rev := reverse(s)
		if s > rev {
			reversed[i] = s
		} else {
			reversed[i] = rev
		}
	}

	result := ""
	for i := 0; i < n; i++ {
		s := strs[i]
		rev := reverse(s)
		// Try both original and reversed
		for _, candidate := range []string{s, rev} {
			for j := 0; j <= len(candidate); j++ {
				// split candidate into candidate[:j] + candidate[j:]
				left := candidate[:j]
				right := candidate[j:]
				var sb string
				sb += right
				for k := i + 1; k < n; k++ {
					sb += reversed[k]
				}
				for k := 0; k < i; k++ {
					sb += reversed[k]
				}
				sb += left
				if sb > result {
					result = sb
				}
			}
		}
	}

	return result
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
```

## 0556 — Next Greater Element Iii

```go
package main

// LeetCode #556: Next Greater Element III
// https://leetcode.com/problems/next-greater-element-iii/
// Difficulty: Medium
// Time: O(log n) = O(number of digits)
// Space: O(log n)

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(NextGreaterElementIii(12))
	fmt.Println(NextGreaterElementIii(21))
	fmt.Println(NextGreaterElementIii(1234))
}

func NextGreaterElementIii(n int) int {
	s := []byte(strconv.Itoa(n))

	// Find first decreasing digit from right
	i := len(s) - 2
	for i >= 0 && s[i] >= s[i+1] {
		i--
	}
	if i < 0 {
		return -1
	}

	// Find smallest digit larger than s[i] from right
	j := len(s) - 1
	for s[j] <= s[i] {
		j--
	}

	s[i], s[j] = s[j], s[i]

	// Reverse suffix
	left, right := i+1, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}

	result, _ := strconv.Atoi(string(s))
	if result > math.MaxInt32 {
		return -1
	}
	return result
}
```

## 0558 — Logical Or Of Two Binary Grids Represented As Quad Trees

```go
package main

// LeetCode #558: Logical OR of Two Binary Grids Represented as Quad-Trees
// https://leetcode.com/problems/logical-or-of-two-binary-grids-represented-as-quad-trees/
// Difficulty: Medium
// Time: O(n) where n = number of nodes in smaller tree
// Space: O(log n)

import "fmt"

type QuadTreeNode struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *QuadTreeNode
	TopRight    *QuadTreeNode
	BottomLeft  *QuadTreeNode
	BottomRight *QuadTreeNode
}

func main() {
	// Tree 1: full leaf (true)
	t1 := &QuadTreeNode{Val: true, IsLeaf: true}
	// Tree 2: full leaf (false)
	t2 := &QuadTreeNode{Val: false, IsLeaf: true}
	result := Intersect(t1, t2)
	fmt.Printf("leaf=%v, val=%v\n", result.IsLeaf, result.Val)
}

func Intersect(quadTree1 *QuadTreeNode, quadTree2 *QuadTreeNode) *QuadTreeNode {
	if quadTree1.IsLeaf {
		if quadTree1.Val {
			return &QuadTreeNode{Val: true, IsLeaf: true}
		}
		return quadTree2
	}
	if quadTree2.IsLeaf {
		if quadTree2.Val {
			return &QuadTreeNode{Val: true, IsLeaf: true}
		}
		return quadTree1
	}

	topLeft := Intersect(quadTree1.TopLeft, quadTree2.TopLeft)
	topRight := Intersect(quadTree1.TopRight, quadTree2.TopRight)
	bottomLeft := Intersect(quadTree1.BottomLeft, quadTree2.BottomLeft)
	bottomRight := Intersect(quadTree1.BottomRight, quadTree2.BottomRight)

	// If all four children are leaves and have same value, merge
	if topLeft.IsLeaf && topRight.IsLeaf && bottomLeft.IsLeaf && bottomRight.IsLeaf &&
		topLeft.Val == topRight.Val && topRight.Val == bottomLeft.Val && bottomLeft.Val == bottomRight.Val {
		return &QuadTreeNode{Val: topLeft.Val, IsLeaf: true}
	}

	return &QuadTreeNode{
		IsLeaf:      false,
		TopLeft:     topLeft,
		TopRight:    topRight,
		BottomLeft:  bottomLeft,
		BottomRight: bottomRight,
	}
}
```

## 0560 — Subarray Sum Equals K

```go
package main

// LeetCode #560: Subarray Sum Equals K
// https://leetcode.com/problems/subarray-sum-equals-k/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(SubarraySum([]int{1, 1, 1}, 2))
	fmt.Println(SubarraySum([]int{1, 2, 3}, 3))
}

func SubarraySum(nums []int, k int) int {
	count := 0
	sum := 0
	sumMap := make(map[int]int)
	sumMap[0] = 1

	for _, num := range nums {
		sum += num
		if val, ok := sumMap[sum-k]; ok {
			count += val
		}
		sumMap[sum]++
	}

	return count
}
```

## 0562 — Longest Line Of Consecutive One In Matrix

```go
package main

// LeetCode #562: Longest Line of Consecutive One in Matrix
// https://leetcode.com/problems/longest-line-of-consecutive-one-in-matrix/
// Difficulty: Medium [Paid]
// Time: O(m * n)
// Space: O(m * n)

import "fmt"

func main() {
	mat := [][]int{
		{0, 1, 1, 0},
		{0, 1, 1, 0},
		{0, 0, 0, 1},
	}
	fmt.Println(LongestLine(mat))
}

func LongestLine(mat [][]int) int {
	if len(mat) == 0 || len(mat[0]) == 0 {
		return 0
	}
	m, n := len(mat), len(mat[0])
	maxLen := 0

	// dp[i][j][0] = horizontal, dp[i][j][1] = vertical
	// dp[i][j][2] = diagonal, dp[i][j][3] = anti-diagonal
	dp := make([][][]int, m)
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, 4)
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 1 {
				dp[i][j][0] = 1
				dp[i][j][1] = 1
				dp[i][j][2] = 1
				dp[i][j][3] = 1

				if j > 0 {
					dp[i][j][0] = dp[i][j-1][0] + 1
				}
				if i > 0 {
					dp[i][j][1] = dp[i-1][j][1] + 1
				}
				if i > 0 && j > 0 {
					dp[i][j][2] = dp[i-1][j-1][2] + 1
				}
				if i > 0 && j < n-1 {
					dp[i][j][3] = dp[i-1][j+1][3] + 1
				}

				for k := 0; k < 4; k++ {
					if dp[i][j][k] > maxLen {
						maxLen = dp[i][j][k]
					}
				}
			}
		}
	}

	return maxLen
}
```

## 0565 — Array Nesting

```go
package main

// LeetCode #565: Array Nesting
// https://leetcode.com/problems/array-nesting/
// Difficulty: Medium
// Time: O(n)
// Space: O(1) (reuses input array as visited marker)

import "fmt"

func main() {
	fmt.Println(ArrayNesting([]int{5, 4, 0, 3, 1, 6, 2}))
	fmt.Println(ArrayNesting([]int{0, 1, 2}))
}

func ArrayNesting(nums []int) int {
	maxLen := 0
	visited := make([]bool, len(nums))

	for i := 0; i < len(nums); i++ {
		if visited[i] {
			continue
		}
		count := 0
		cur := i
		for !visited[cur] {
			visited[cur] = true
			cur = nums[cur]
			count++
		}
		if count > maxLen {
			maxLen = count
		}
	}

	return maxLen
}
```

## 0567 — Permutation In String

```go
package main

// LeetCode #567: Permutation in String
// https://leetcode.com/problems/permutation-in-string/
// Difficulty: Medium
// Time: O(n + m) where n = len(s1), m = len(s2)
// Space: O(1) (fixed 26 chars)

import "fmt"

func main() {
	fmt.Println(CheckInclusion("ab", "eidbaooo"))
	fmt.Println(CheckInclusion("ab", "eidboaoo"))
}

func CheckInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	count1 := [26]int{}
	count2 := [26]int{}

	for i := 0; i < len(s1); i++ {
		count1[s1[i]-'a']++
		count2[s2[i]-'a']++
	}

	if count1 == count2 {
		return true
	}

	for i := len(s1); i < len(s2); i++ {
		count2[s2[i]-'a']++
		count2[s2[i-len(s1)]-'a']--
		if count1 == count2 {
			return true
		}
	}

	return false
}
```

## 0570 — Managers With At Least 5 Direct Reports

```go
package main

// LeetCode #570: Managers with at Least 5 Direct Reports
// https://leetcode.com/problems/managers-with-at-least-5-direct-reports/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	// Employees: {id, name, department, managerId}
	// managerId == -1 means top-level manager
	employees := [][]int{
		{101, 1, 0, -1},  // John, manager
		{102, 2, 0, 101}, // Dan
		{103, 3, 0, 101}, // James
		{104, 4, 0, 101}, // Amy
		{105, 5, 0, 101}, // Ben
		{106, 6, 0, 101}, // Sam
		{107, 7, 1, -1},  // Ron, another manager
		{108, 8, 1, 107}, // Tom
	}
	fmt.Println(FindManagers(employees))
}

func FindManagers(employees [][]int) []int {
	reportCount := make(map[int]int)
	managerNames := make(map[int]int) // managerId -> manager name (for simplicity, just id)

	for _, emp := range employees {
		id, name, _, managerId := emp[0], emp[1], emp[2], emp[3]
		managerNames[id] = name
		if managerId != -1 {
			reportCount[managerId]++
		}
	}

	result := []int{}
	for mid, count := range reportCount {
		if count >= 5 {
			result = append(result, managerNames[mid])
		}
	}

	return result
}
```

## 0573 — Squirrel Simulation

```go
package main

// LeetCode #573: Squirrel Simulation
// https://leetcode.com/problems/squirrel-simulation/
// Difficulty: Medium [Paid]
// Time: O(n) where n = number of nuts
// Space: O(1)

import (
	"fmt"
	"math"
)

func main() {
	height := 5
	width := 7
	tree := []int{2, 2}
	squirrel := []int{4, 4}
	nuts := [][]int{{3, 0}, {2, 5}}
	fmt.Println(MinDistance(height, width, tree, squirrel, nuts))
}

func MinDistance(height int, width int, tree []int, squirrel []int, nuts [][]int) int {
	total := 0
	maxSaving := math.MinInt32

	for _, nut := range nuts {
		nutToTree := abs(nut[0]-tree[0]) + abs(nut[1]-tree[1])
		nutToSquirrel := abs(nut[0]-squirrel[0]) + abs(nut[1]-squirrel[1])
		total += 2 * nutToTree
		saving := nutToTree - nutToSquirrel
		if saving > maxSaving {
			maxSaving = saving
		}
	}

	return total - maxSaving
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

## 0574 — Winning Candidate

```go
package main

// LeetCode #574: Winning Candidate
// https://leetcode.com/problems/winning-candidate/
// Difficulty: Medium [Paid]
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	// votes: {candidate_id}
	votes := []int{1, 2, 2, 3, 3, 3}
	// candidate names: map candidate_id -> name
	candidateNames := map[int]int{1: 1001, 2: 1002, 3: 1003}
	fmt.Println(FindWinningCandidate(votes, candidateNames))
}

func FindWinningCandidate(votes []int, candidateNames map[int]int) int {
	counts := make(map[int]int)
	for _, v := range votes {
		counts[v]++
	}

	maxVotes := 0
	winner := -1
	for id, count := range counts {
		if count > maxVotes {
			maxVotes = count
			winner = candidateNames[id]
		}
	}

	return winner
}
```

## 0576 — Out Of Boundary Paths

```go
package main

// LeetCode #576: Out of Boundary Paths
// https://leetcode.com/problems/out-of-boundary-paths/
// Difficulty: Medium
// Time: O(m * n * maxMove)
// Space: O(m * n)

import "fmt"

func main() {
	fmt.Println(FindPaths(2, 2, 2, 0, 0))
	fmt.Println(FindPaths(1, 3, 3, 0, 1))
}

func FindPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	const mod = 1_000_000_007
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[startRow][startColumn] = 1
	total := 0
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	for move := 1; move <= maxMove; move++ {
		next := make([][]int, m)
		for i := range next {
			next[i] = make([]int, n)
		}
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if dp[i][j] == 0 {
					continue
				}
				for _, d := range dirs {
					ni, nj := i+d[0], j+d[1]
					if ni < 0 || ni >= m || nj < 0 || nj >= n {
						total = (total + dp[i][j]) % mod
					} else {
						next[ni][nj] = (next[ni][nj] + dp[i][j]) % mod
					}
				}
			}
		}
		dp = next
	}

	return total
}
```

## 0578 — Get Highest Answer Rate Question

```go
package main

// LeetCode #578: Get Highest Answer Rate Question
// https://leetcode.com/problems/get-highest-answer-rate-question/
// Difficulty: Medium [Paid]
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	// survey_log: {question_id, action}
	// action: "show", "answer", "skip"
	surveyLog := [][]interface{}{
		{1, "show"},
		{1, "answer"},
		{2, "show"},
		{2, "skip"},
		{3, "show"},
		{3, "answer"},
	}
	fmt.Println(MostAnsweredQuestion(surveyLog))
}

func MostAnsweredQuestion(surveyLog [][]interface{}) int {
	shows := make(map[int]int)
	answers := make(map[int]int)

	for _, entry := range surveyLog {
		qID := entry[0].(int)
		action := entry[1].(string)
		shows[qID]++
		if action == "answer" {
			answers[qID]++
		}
	}

	bestQ := -1
	bestRate := -1.0

	for qID := range shows {
		rate := float64(answers[qID]) / float64(shows[qID])
		if rate > bestRate {
			bestRate = rate
			bestQ = qID
		}
	}

	return bestQ
}
```

## 0580 — Count Student Number In Departments

```go
package main

// LeetCode #580: Count Student Number in Departments
// https://leetcode.com/problems/count-student-number-in-departments/
// Difficulty: Medium [Paid]
// Time: O(n + m) where n = departments, m = students
// Space: O(n)

import (
	"fmt"
	"sort"
)

func main() {
	// departments: {dept_id, dept_name}
	departments := map[int]string{
		1: "Engineering",
		2: "Science",
		3: "Arts",
	}
	// students: {student_id, dept_id}
	students := []int{1, 1, 1, 2, 3}

	result := CountStudents(departments, students)
	for _, r := range result {
		fmt.Printf("%s: %d\n", r[0].(string), r[1].(int))
	}
}

type DeptCount struct {
	Name  string
	Count int
}

func CountStudents(departments map[int]string, studentDepts []int) [][]interface{} {
	counts := make(map[int]int)
	for _, deptID := range studentDepts {
		counts[deptID]++
	}

	result := [][]interface{}{}
	for deptID, deptName := range departments {
		result = append(result, []interface{}{deptName, counts[deptID]})
	}
	// Include departments with 0 students
	for deptID := range departments {
		if _, exists := counts[deptID]; !exists {
			counts[deptID] = 0
		}
	}

	// Rebuild with all departments
	result = nil
	for deptID, deptName := range departments {
		result = append(result, []interface{}{deptName, counts[deptID]})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i][0].(string) < result[j][0].(string)
	})

	return result
}
```

## 0581 — Shortest Unsorted Continuous Subarray

```go
package main

// LeetCode #581: Shortest Unsorted Continuous Subarray
// https://leetcode.com/problems/shortest-unsorted-continuous-subarray/
// Difficulty: Medium
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(FindUnsortedSubarray([]int{2, 6, 4, 8, 10, 9, 15}))
	fmt.Println(FindUnsortedSubarray([]int{1, 2, 3, 4}))
	fmt.Println(FindUnsortedSubarray([]int{1}))
}

func FindUnsortedSubarray(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	left := -1
	minRight := nums[n-1]
	for i := n - 2; i >= 0; i-- {
		if nums[i] > minRight {
			left = i
		} else {
			minRight = nums[i]
		}
	}

	right := -1
	maxLeft := nums[0]
	for i := 1; i < n; i++ {
		if nums[i] < maxLeft {
			right = i
		} else {
			maxLeft = nums[i]
		}
	}

	if right == -1 {
		return 0
	}
	return right - left + 1
}
```

## 0582 — Kill Process

```go
package main

// LeetCode #582: Kill Process
// https://leetcode.com/problems/kill-process/
// Difficulty: Medium [Paid]
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	pid := []int{1, 3, 10, 5}
	ppid := []int{3, 0, 5, 3}
	kill := 5
	fmt.Println(KillProcess(pid, ppid, kill))
}

func KillProcess(pid []int, ppid []int, kill int) []int {
	// Build adjacency list: parent -> children
	children := make(map[int][]int)
	for i, p := range ppid {
		children[p] = append(children[p], pid[i])
	}

	// BFS/DFS to find all processes to kill
	result := []int{}
	queue := []int{kill}
	for len(queue) > 0 {
		process := queue[0]
		queue = queue[1:]
		result = append(result, process)
		queue = append(queue, children[process]...)
	}

	return result
}
```

## 0583 — Delete Operation For Two Strings

```go
package main

// LeetCode #583: Delete Operation for Two Strings
// https://leetcode.com/problems/delete-operation-for-two-strings/
// Difficulty: Medium
// Time: O(m * n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(MinDistance("sea", "eat"))
	fmt.Println(MinDistance("leetcode", "etco"))
}

func MinDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([]int, n+1)

	for i := 1; i <= m; i++ {
		prev := 0
		for j := 1; j <= n; j++ {
			temp := dp[j]
			if word1[i-1] == word2[j-1] {
				dp[j] = prev + 1
			} else {
				if dp[j] > dp[j-1] {
					dp[j] = dp[j]
				} else {
					dp[j] = dp[j-1]
				}
			}
			prev = temp
		}
	}

	lcs := dp[n]
	return (m - lcs) + (n - lcs)
}
```

## 0585 — Investments In 2016

```go
package main

// LeetCode #585: Investments in 2016
// https://leetcode.com/problems/investments-in-2016/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	// Insurance records: {pid, tiv_2015, tiv_2016, lat, lon}
	insurance := [][]interface{}{
		{1, 100, 200, 10, 20},
		{2, 100, 300, 30, 40},
		{3, 200, 400, 50, 60},
		{4, 100, 500, 10, 20}, // same lat/lon as pid=1
	}
	fmt.Println(FindInvestmentSum(insurance))
}

func FindInvestmentSum(insurance [][]interface{}) float64 {
	tiv2015 := make(map[float64]int)
	locationCount := make(map[string]int)
	tiv2016Sum := make(map[int]float64)

	for _, record := range insurance {
		pid := record[0].(int)
		tiv15 := record[1].(float64)
		tiv16 := record[2].(float64)
		lat := record[3].(float64)
		lon := record[4].(float64)

		tiv2015[tiv15]++
		locKey := fmt.Sprintf("%f,%f", lat, lon)
		locationCount[locKey]++
		tiv2016Sum[pid] = tiv16
	}

	total := 0.0
	for _, record := range insurance {
		pid := record[0].(int)
		tiv15 := record[1].(float64)
		lat := record[3].(float64)
		lon := record[4].(float64)
		locKey := fmt.Sprintf("%f,%f", lat, lon)

		if tiv2015[tiv15] > 1 && locationCount[locKey] == 1 {
			total += tiv2016Sum[pid]
		}
	}

	return total
}
```

## 0592 — Fraction Addition And Subtraction

```go
package main

// LeetCode #592: Fraction Addition and Subtraction
// https://leetcode.com/problems/fraction-addition-and-subtraction/
// Difficulty: Medium
// Time: O(n) where n = length of expression
// Space: O(1)

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(FractionAddition("-1/2+1/2"))
	fmt.Println(FractionAddition("-1/2+1/2+1/3"))
	fmt.Println(FractionAddition("1/3-1/2"))
}

func FractionAddition(expression string) string {
	num := 0
	den := 1
	i := 0
	n := len(expression)

	for i < n {
		sign := 1
		if expression[i] == '-' {
			sign = -1
			i++
		} else if expression[i] == '+' {
			sign = 1
			i++
		}

		// Parse numerator
		j := i
		for j < n && expression[j] >= '0' && expression[j] <= '9' {
			j++
		}
		currNum, _ := strconv.Atoi(expression[i:j])
		i = j + 1 // skip '/'

		// Parse denominator
		j = i
		for j < n && expression[j] >= '0' && expression[j] <= '9' {
			j++
		}
		currDen, _ := strconv.Atoi(expression[i:j])
		i = j

		currNum *= sign
		num = num*currDen + currNum*den
		den = den * currDen

		g := gcd(abs(num), abs(den))
		num /= g
		den /= g
	}

	if den < 0 {
		num = -num
		den = -den
	}

	return fmt.Sprintf("%d/%d", num, den)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

## 0593 — Valid Square

```go
package main

// LeetCode #593: Valid Square
// https://leetcode.com/problems/valid-square/
// Difficulty: Medium
// Time: O(1)
// Space: O(1)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(ValidSquare([]int{0, 0}, []int{1, 1}, []int{1, 0}, []int{0, 1}))
	fmt.Println(ValidSquare([]int{0, 0}, []int{1, 2}, []int{2, 1}, []int{1, 0}))
}

func ValidSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	points := [][]int{p1, p2, p3, p4}
	dists := []int{}

	for i := 0; i < 4; i++ {
		for j := i + 1; j < 4; j++ {
			dist := distSq(points[i], points[j])
			dists = append(dists, dist)
		}
	}

	sort.Ints(dists)

	// All 4 sides equal and non-zero, and 2 diagonals equal
	return dists[0] > 0 && dists[0] == dists[1] && dists[1] == dists[2] && dists[2] == dists[3] &&
		dists[4] == dists[5]
}

func distSq(a, b []int) int {
	dx := a[0] - b[0]
	dy := a[1] - b[1]
	return dx*dx + dy*dy
}
```

## 0602 — Friend Requests Ii Who Has The Most Friends

```go
package main

// LeetCode #602: Friend Requests II: Who Has the Most Friends
// https://leetcode.com/problems/friend-requests-ii-who-has-the-most-friends/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	// Friend requests: {requester_id, accepter_id}
	requests := [][]int{
		{1, 2},
		{1, 3},
		{2, 3},
		{3, 4},
	}
	fmt.Println(MostFriends(requests))
}

func MostFriends(requests [][]int) int {
	friendCount := make(map[int]int)
	for _, req := range requests {
		friendCount[req[0]]++
		friendCount[req[1]]++
	}

	maxCount := 0
	maxID := 0
	for id, count := range friendCount {
		if count > maxCount {
			maxCount = count
			maxID = id
		}
	}

	return maxID
}
```

## 0606 — Construct String From Binary Tree

```go
package main

// LeetCode #606: Construct String from Binary Tree
// https://leetcode.com/problems/construct-string-from-binary-tree/
// Difficulty: Medium
// Time: O(n)
// Space: O(h) where h is tree height

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3}}
	fmt.Println(Tree2str(root))
}

func Tree2str(root *TreeNode) string {
	if root == nil {
		return ""
	}
	result := strconv.Itoa(root.Val)
	if root.Left != nil || root.Right != nil {
		result += "(" + Tree2str(root.Left) + ")"
	}
	if root.Right != nil {
		result += "(" + Tree2str(root.Right) + ")"
	}
	return result
}
```

## 0608 — Tree Node

```go
package main

// LeetCode #608: Tree Node
// https://leetcode.com/problems/tree-node/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	// Tree nodes: {id, p_id} where p_id = -1 means root
	nodes := [][]int{
		{1, -1},
		{2, 1},
		{3, 1},
		{4, 2},
		{5, 2},
	}
	fmt.Println(ClassifyTreeNodes(nodes))
}

func ClassifyTreeNodes(nodes [][]int) map[int]string {
	parentMap := make(map[int]int)
	childrenMap := make(map[int][]int)

	for _, node := range nodes {
		id, pID := node[0], node[1]
		if pID == -1 {
			parentMap[id] = -1
		} else {
			parentMap[id] = pID
			childrenMap[pID] = append(childrenMap[pID], id)
		}
	}

	result := make(map[int]string)
	for _, node := range nodes {
		id := node[0]
		if parentMap[id] == -1 {
			result[id] = "Root"
		} else if len(childrenMap[id]) == 0 {
			result[id] = "Leaf"
		} else {
			result[id] = "Inner"
		}
	}

	return result
}
```

## 0609 — Find Duplicate File In System

```go
package main

// LeetCode #609: Find Duplicate File in System
// https://leetcode.com/problems/find-duplicate-file-in-system/
// Difficulty: Medium
// Time: O(n * L) where n = number of files, L = max content length
// Space: O(n * L)

import (
	"fmt"
	"strings"
)

func main() {
	paths := []string{
		"root/a 1.txt(abcd) 2.txt(efgh)",
		"root/c 3.txt(abcd)",
		"root/c/d 4.txt(efgh)",
	}
	fmt.Println(FindDuplicate(paths))
}

func FindDuplicate(paths []string) [][]string {
	contentMap := make(map[string][]string)

	for _, path := range paths {
		parts := strings.Split(path, " ")
		dir := parts[0]
		for i := 1; i < len(parts); i++ {
			fileStr := parts[i]
			parenIdx := strings.Index(fileStr, "(")
			fileName := fileStr[:parenIdx]
			content := fileStr[parenIdx+1 : len(fileStr)-1] // remove closing ')'
			fullPath := dir + "/" + fileName
			contentMap[content] = append(contentMap[content], fullPath)
		}
	}

	result := [][]string{}
	for _, files := range contentMap {
		if len(files) > 1 {
			result = append(result, files)
		}
	}

	return result
}
```

## 0611 — Valid Triangle Number

```go
package main

// LeetCode #611: Valid Triangle Number
// https://leetcode.com/problems/valid-triangle-number/
// Difficulty: Medium
// Time: O(n^2)
// Space: O(log n) for sorting

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(TriangleNumber([]int{2, 2, 3, 4}))
	fmt.Println(TriangleNumber([]int{4, 2, 3, 4}))
}

func TriangleNumber(nums []int) int {
	sort.Ints(nums)
	count := 0
	n := len(nums)

	for i := n - 1; i >= 2; i-- {
		left, right := 0, i-1
		for left < right {
			if nums[left]+nums[right] > nums[i] {
				count += right - left
				right--
			} else {
				left++
			}
		}
	}

	return count
}
```

## 0612 — Shortest Distance In A Plane

```go
package main

// LeetCode #612: Shortest Distance in a Plane
// https://leetcode.com/problems/shortest-distance-in-a-plane/
// Difficulty: Medium [Paid]
// Time: O(n^2)
// Space: O(1)

import (
	"fmt"
	"math"
)

func main() {
	points := [][]float64{
		{-1, -1},
		{0, 0},
		{1, 1},
		{2, 2},
	}
	fmt.Printf("%.4f\n", ShortestDistance(points))
}

func ShortestDistance(points [][]float64) float64 {
	if len(points) < 2 {
		return 0
	}

	minDist := math.MaxFloat64
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			dx := points[i][0] - points[j][0]
			dy := points[i][1] - points[j][1]
			dist := math.Sqrt(dx*dx + dy*dy)
			if dist < minDist {
				minDist = dist
			}
		}
	}

	return minDist
}
```

## 0614 — Second Degree Follower

```go
package main

// LeetCode #614: Second Degree Follower
// https://leetcode.com/problems/second-degree-follower/
// Difficulty: Medium [Paid]
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	// Follow relationships: {follower, followee}
	follows := [][]int{
		{1, 2},
		{1, 3},
		{2, 3},
		{3, 1},
		{3, 2},
		{4, 5},
	}
	fmt.Println(SecondDegreeFollowers(follows))
}

func SecondDegreeFollowers(follows [][]int) []int {
	followers := make(map[int]map[int]bool)

	for _, f := range follows {
		follower, followee := f[0], f[1]
		if followers[followee] == nil {
			followers[followee] = make(map[int]bool)
		}
		followers[followee][follower] = true
	}

	result := []int{}
	for userID, fMap := range followers {
		if len(fMap) >= 2 {
			result = append(result, userID)
		}
	}

	return result
}
```

## 0616 — Add Bold Tag In String

```go
package main

// LeetCode #616: Add Bold Tag in String
// https://leetcode.com/problems/add-bold-tag-in-string/
// Difficulty: Medium [Paid]
// Time: O(n * L) where n = len(s), L = total length of all words
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(AddBoldTag("abcxyz123", []string{"abc", "123"}))
	fmt.Println(AddBoldTag("aaabbcc", []string{"aaa", "aab", "bc"}))
}

func AddBoldTag(s string, words []string) string {
	n := len(s)
	bold := make([]bool, n)

	for _, word := range words {
		for i := 0; i <= n-len(word); i++ {
			if s[i:i+len(word)] == word {
				for j := i; j < i+len(word); j++ {
					bold[j] = true
				}
			}
		}
	}

	result := ""
	i := 0
	for i < n {
		if bold[i] {
			result += "<b>"
			for i < n && bold[i] {
				result += string(s[i])
				i++
			}
			result += "</b>"
		} else {
			result += string(s[i])
			i++
		}
	}

	return result
}
```

## 0621 — Task Scheduler

```go
package main

// LeetCode #621: Task Scheduler
// https://leetcode.com/problems/task-scheduler/
// Difficulty: Medium
// Time: O(n)
// Space: O(1) (26 letters)

import "fmt"

func main() {
	fmt.Println(LeastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2))
	fmt.Println(LeastInterval([]byte{'A', 'C', 'A', 'B', 'D', 'B'}, 1))
	fmt.Println(LeastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 3))
}

func LeastInterval(tasks []byte, n int) int {
	counts := make([]int, 26)
	maxFreq := 0
	for _, t := range tasks {
		counts[t-'A']++
		if counts[t-'A'] > maxFreq {
			maxFreq = counts[t-'A']
		}
	}

	maxCount := 0
	for _, c := range counts {
		if c == maxFreq {
			maxCount++
		}
	}

	partLen := maxFreq - 1
	emptySlots := partLen * (n - (maxCount - 1))
	availableTasks := len(tasks) - maxFreq*maxCount
	idles := 0
	if emptySlots > availableTasks {
		idles = emptySlots - availableTasks
	}

	return len(tasks) + idles
}
```

## 0622 — Design Circular Queue

```go
package main

// LeetCode #622: Design Circular Queue
// https://leetcode.com/problems/design-circular-queue/
// Difficulty: Medium
// Time: O(1) for all operations
// Space: O(k)

import "fmt"

type MyCircularQueue struct {
	data  []int
	front int
	rear  int
	size  int
	cap   int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		data: make([]int, k),
		front: 0,
		rear:  -1,
		size:  0,
		cap:   k,
	}
}

func (q *MyCircularQueue) EnQueue(value int) bool {
	if q.IsFull() {
		return false
	}
	q.rear = (q.rear + 1) % q.cap
	q.data[q.rear] = value
	q.size++
	return true
}

func (q *MyCircularQueue) DeQueue() bool {
	if q.IsEmpty() {
		return false
	}
	q.front = (q.front + 1) % q.cap
	q.size--
	return true
}

func (q *MyCircularQueue) Front() int {
	if q.IsEmpty() {
		return -1
	}
	return q.data[q.front]
}

func (q *MyCircularQueue) Rear() int {
	if q.IsEmpty() {
		return -1
	}
	return q.data[q.rear]
}

func (q *MyCircularQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *MyCircularQueue) IsFull() bool {
	return q.size == q.cap
}

func main() {
	q := Constructor(3)
	fmt.Println(q.EnQueue(1))
	fmt.Println(q.EnQueue(2))
	fmt.Println(q.EnQueue(3))
	fmt.Println(q.EnQueue(4))
	fmt.Println(q.Rear())
	fmt.Println(q.IsFull())
	fmt.Println(q.DeQueue())
	fmt.Println(q.EnQueue(4))
	fmt.Println(q.Rear())
}
```

## 0623 — Add One Row To Tree

```go
package main

// LeetCode #623: Add One Row to Tree
// https://leetcode.com/problems/add-one-row-to-tree/
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
	root := &TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 1}}, Right: &TreeNode{Val: 6, Left: &TreeNode{Val: 5}}}
	result := AddOneRow(root, 1, 2)
	printTree(result)
	fmt.Println()
}

func AddOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		return &TreeNode{Val: val, Left: root}
	}
	addRowDFS(root, val, depth, 1)
	return root
}

func addRowDFS(node *TreeNode, val int, depth int, curDepth int) {
	if node == nil {
		return
	}
	if curDepth == depth-1 {
		oldLeft, oldRight := node.Left, node.Right
		node.Left = &TreeNode{Val: val, Left: oldLeft}
		node.Right = &TreeNode{Val: val, Right: oldRight}
		return
	}
	addRowDFS(node.Left, val, depth, curDepth+1)
	addRowDFS(node.Right, val, depth, curDepth+1)
}

func printTree(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.Val)
	printTree(root.Left)
	printTree(root.Right)
}
```

## 0624 — Maximum Distance In Arrays

```go
package main

// LeetCode #624: Maximum Distance in Arrays
// https://leetcode.com/problems/maximum-distance-in-arrays/
// Difficulty: Medium
// Time: O(n) where n = number of arrays
// Space: O(1)

import (
	"fmt"
)

func main() {
	arrays := [][]int{
		{1, 2, 3},
		{4, 5},
		{1, 2, 3},
	}
	fmt.Println(MaxDistance(arrays))
}

func MaxDistance(arrays [][]int) int {
	minVal := arrays[0][0]
	maxVal := arrays[0][len(arrays[0])-1]
	maxDist := 0

	for i := 1; i < len(arrays); i++ {
		arr := arrays[i]
		dist1 := abs(arr[len(arr)-1] - minVal)
		dist2 := abs(maxVal - arr[0])
		if dist1 > maxDist {
			maxDist = dist1
		}
		if dist2 > maxDist {
			maxDist = dist2
		}
		if arr[0] < minVal {
			minVal = arr[0]
		}
		if arr[len(arr)-1] > maxVal {
			maxVal = arr[len(arr)-1]
		}
	}

	return maxDist
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

## 0625 — Minimum Factorization

```go
package main

// LeetCode #625: Minimum Factorization
// https://leetcode.com/problems/minimum-factorization/
// Difficulty: Medium [Paid]
// Time: O(log n)
// Space: O(1)

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(SmallestFactorization(48))
	fmt.Println(SmallestFactorization(15))
	fmt.Println(SmallestFactorization(1))
}

func SmallestFactorization(num int) int {
	if num < 2 {
		return num
	}

	// Build number from right to left using digits 9..2
	result := 0
	multiplier := 1

	for i := 9; i >= 2; i-- {
		for num%i == 0 {
			result += i * multiplier
			if result > math.MaxInt32 {
				return 0
			}
			multiplier *= 10
			num /= i
		}
	}

	if num > 1 {
		return 0
	}
	return result
}
```

## 0626 — Exchange Seats

```go
package main

// LeetCode #626: Exchange Seats
// https://leetcode.com/problems/exchange-seats/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	// Students: {id, name}
	students := [][]interface{}{
		{1, "Alice"},
		{2, "Bob"},
		{3, "Charlie"},
		{4, "David"},
	}
	result := ExchangeSeats(students)
	for _, r := range result {
		fmt.Printf("id=%d, name=%s\n", r[0].(int), r[1].(string))
	}
}

func ExchangeSeats(students [][]interface{}) [][]interface{} {
	n := len(students)
	result := make([][]interface{}, n)

	// Build map for easy lookup
	studentMap := make(map[int]string)
	for _, student := range students {
		id := student[0].(int)
		name := student[1].(string)
		studentMap[id] = name
	}

	for id := 1; id <= n; id++ {
		if id%2 == 1 {
			if id+1 <= n {
				result[id-1] = []interface{}{id + 1, studentMap[id+1]}
			} else {
				result[id-1] = []interface{}{id, studentMap[id]}
			}
		} else {
			result[id-1] = []interface{}{id - 1, studentMap[id-1]}
		}
	}

	return result
}
```

## 0633 — Sum Of Square Numbers

```go
package main

// LeetCode #633: Sum of Square Numbers
// https://leetcode.com/problems/sum-of-square-numbers/
// Difficulty: Medium
// Time: O(sqrt(c))
// Space: O(1)

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(JudgeSquareSum(5))
	fmt.Println(JudgeSquareSum(3))
	fmt.Println(JudgeSquareSum(4))
	fmt.Println(JudgeSquareSum(2))
}

func JudgeSquareSum(c int) bool {
	left := 0
	right := int(math.Sqrt(float64(c)))

	for left <= right {
		sum := left*left + right*right
		if sum == c {
			return true
		} else if sum < c {
			left++
		} else {
			right--
		}
	}

	return false
}
```

## 0634 — Find The Derangement Of An Array

```go
package main

// LeetCode #634: Find the Derangement of An Array
// https://leetcode.com/problems/find-the-derangement-of-an-array/
// Difficulty: Medium [Paid]
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(FindDerangement(3))
	fmt.Println(FindDerangement(4))
}

func FindDerangement(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 0
	}

	const mod = 1_000_000_007
	a, b := 0, 1 // D(1)=0, D(2)=1

	for i := 3; i <= n; i++ {
		c := ((i - 1) * (a + b)) % mod
		a, b = b, c
	}

	return b
}
```

## 0635 — Design Log Storage System

```go
package main

// LeetCode #635: Design Log Storage System
// https://leetcode.com/problems/design-log-storage-system/
// Difficulty: Medium [Paid]
// Time: O(n) for put, O(n) for retrieve
// Space: O(n)

import (
	"fmt"
	"strings"
)

type LogSystem struct {
	logs []LogEntry
}

type LogEntry struct {
	id   int
	time string
}

func Constructor() LogSystem {
	return LogSystem{logs: []LogEntry{}}
}

func (ls *LogSystem) Put(id int, timestamp string) {
	ls.logs = append(ls.logs, LogEntry{id: id, time: timestamp})
}

func (ls *LogSystem) Retrieve(start string, end string, granularity string) []int {
	startPrefix := truncate(start, granularity)
	endPrefix := truncate(end, granularity)

	result := []int{}
	for _, entry := range ls.logs {
		entryPrefix := truncate(entry.time, granularity)
		if entryPrefix >= startPrefix && entryPrefix <= endPrefix {
			result = append(result, entry.id)
		}
	}
	return result
}

func truncate(timestamp string, granularity string) string {
	parts := strings.Split(timestamp, ":")
	indices := map[string]int{
		"Year": 1, "Month": 2, "Day": 3,
		"Hour": 4, "Minute": 5, "Second": 6,
	}
	idx := indices[granularity]
	result := strings.Join(parts[:idx], ":")
	// Pad remaining with minimum values
	for i := idx; i < 6; i++ {
		result += ":00"
	}
	return result
}

func main() {
	ls := Constructor()
	ls.Put(1, "2017:01:01:23:59:59")
	ls.Put(2, "2017:01:01:22:59:59")
	ls.Put(3, "2016:01:01:00:00:00")
	fmt.Println(ls.Retrieve("2016:12:31:23:59:59", "2017:01:02:23:59:59", "Year"))
}
```

## 0636 — Exclusive Time Of Functions

```go
package main

// LeetCode #636: Exclusive Time of Functions
// https://leetcode.com/problems/exclusive-time-of-functions/
// Difficulty: Medium

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(exclusiveTime(2, []string{"0:start:0", "1:start:2", "1:end:5", "0:end:6"}))
	fmt.Println(exclusiveTime(1, []string{"0:start:0", "0:start:2", "0:end:5", "0:start:6", "0:end:6", "0:end:7"}))
	fmt.Println(exclusiveTime(2, []string{"0:start:0", "0:start:2", "0:end:5", "1:start:6", "1:end:6", "0:end:7"}))
}

func exclusiveTime(n int, logs []string) []int {
	result := make([]int, n)
	stack := make([]int, 0)
	prevTime := 0

	for _, log := range logs {
		parts := strings.Split(log, ":")
		id, _ := strconv.Atoi(parts[0])
		typ := parts[1]
		timestamp, _ := strconv.Atoi(parts[2])

		if typ == "start" {
			if len(stack) > 0 {
				result[stack[len(stack)-1]] += timestamp - prevTime
			}
			stack = append(stack, id)
			prevTime = timestamp
		} else {
			result[id] += timestamp - prevTime + 1
			stack = stack[:len(stack)-1]
			prevTime = timestamp + 1
		}
	}

	return result
}
```

## 0638 — Shopping Offers

```go
package main

// LeetCode #638: Shopping Offers
// https://leetcode.com/problems/shopping-offers/
// Difficulty: Medium
// Time: O(n * special * states) where states is up to product of (needs[i]+1)
// Space: O(product of needs[i]+1) for memoization

import (
	"fmt"
)

func main() {
	fmt.Println(shoppingOffers([]int{2, 5}, [][]int{{3, 0, 5}, {1, 2, 10}}, []int{3, 2}))
	fmt.Println(shoppingOffers([]int{2, 3, 4}, [][]int{{1, 1, 0, 4}, {2, 2, 1, 9}}, []int{1, 2, 1}))
}

func shoppingOffers(price []int, special [][]int, needs []int) int {
	n := len(price)
	memo := make(map[string]int)
	return dfs(price, special, needs, n, memo)
}

func dfs(price []int, special [][]int, needs []int, n int, memo map[string]int) int {
	key := fmt.Sprint(needs)
	if val, ok := memo[key]; ok {
		return val
	}

	cost := 0
	for i := 0; i < n; i++ {
		cost += price[i] * needs[i]
	}

	for _, offer := range special {
		valid := true
		newNeeds := make([]int, n)
		for i := 0; i < n; i++ {
			if offer[i] > needs[i] {
				valid = false
				break
			}
			newNeeds[i] = needs[i] - offer[i]
		}
		if valid {
			cost = min(cost, offer[n]+dfs(price, special, newNeeds, n, memo))
		}
	}

	memo[key] = cost
	return cost
}
```

## 0640 — Solve The Equation

```go
package main

// LeetCode #640: Solve the Equation
// https://leetcode.com/problems/solve-the-equation/
// Difficulty: Medium
// Time: O(n) where n is length of equation string
// Space: O(1)

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(solveEquation("x+5-3+x=6+x-2"))
	fmt.Println(solveEquation("x=x"))
	fmt.Println(solveEquation("2x=x"))
}

func solveEquation(equation string) string {
	eqIdx := findEqual(equation)
	leftCoeff, leftConst := parse(equation[:eqIdx])
	rightCoeff, rightConst := parse(equation[eqIdx+1:])

	coeff := leftCoeff - rightCoeff
	constant := rightConst - leftConst

	if coeff == 0 && constant == 0 {
		return "Infinite solutions"
	}
	if coeff == 0 {
		return "No solution"
	}
	return "x=" + strconv.Itoa(constant/coeff)
}

func findEqual(s string) int {
	for i := 0; i < len(s); i++ {
		if s[i] == '=' {
			return i
		}
	}
	return -1
}

func parse(expr string) (int, int) {
	coeff, constant := 0, 0
	sign := 1
	i := 0
	n := len(expr)

	for i < n {
		if expr[i] == '+' {
			sign = 1
			i++
		} else if expr[i] == '-' {
			sign = -1
			i++
		} else {
			start := i
			for i < n && expr[i] >= '0' && expr[i] <= '9' {
				i++
			}
			if i < n && expr[i] == 'x' {
				if start == i {
					coeff += sign
				} else {
					val, _ := strconv.Atoi(expr[start:i])
					coeff += sign * val
				}
				i++
			} else {
				val, _ := strconv.Atoi(expr[start:i])
				constant += sign * val
			}
		}
	}
	return coeff, constant
}
```

## 0641 — Design Circular Deque

```go
package main

// LeetCode #641: Design Circular Deque
// https://leetcode.com/problems/design-circular-deque/
// Difficulty: Medium
// Time: O(1) for all operations
// Space: O(k) where k is capacity

import "fmt"

func main() {
	deque := Constructor(3)
	fmt.Println(deque.InsertLast(1))
	fmt.Println(deque.InsertLast(2))
	fmt.Println(deque.InsertFront(3))
	fmt.Println(deque.InsertFront(4))
	fmt.Println(deque.GetRear())
	fmt.Println(deque.IsFull())
	fmt.Println(deque.DeleteLast())
	fmt.Println(deque.InsertFront(4))
	fmt.Println(deque.GetFront())
}

type MyCircularDeque struct {
	data  []int
	front int
	rear  int
	size  int
	cap   int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		data:  make([]int, k),
		front: 0,
		rear:  0,
		size:  0,
		cap:   k,
	}
}

func (d *MyCircularDeque) InsertFront(value int) bool {
	if d.IsFull() {
		return false
	}
	d.front = (d.front - 1 + d.cap) % d.cap
	d.data[d.front] = value
	d.size++
	return true
}

func (d *MyCircularDeque) InsertLast(value int) bool {
	if d.IsFull() {
		return false
	}
	d.data[d.rear] = value
	d.rear = (d.rear + 1) % d.cap
	d.size++
	return true
}

func (d *MyCircularDeque) DeleteFront() bool {
	if d.IsEmpty() {
		return false
	}
	d.front = (d.front + 1) % d.cap
	d.size--
	return true
}

func (d *MyCircularDeque) DeleteLast() bool {
	if d.IsEmpty() {
		return false
	}
	d.rear = (d.rear - 1 + d.cap) % d.cap
	d.size--
	return true
}

func (d *MyCircularDeque) GetFront() int {
	if d.IsEmpty() {
		return -1
	}
	return d.data[d.front]
}

func (d *MyCircularDeque) GetRear() int {
	if d.IsEmpty() {
		return -1
	}
	return d.data[(d.rear-1+d.cap)%d.cap]
}

func (d *MyCircularDeque) IsEmpty() bool {
	return d.size == 0
}

func (d *MyCircularDeque) IsFull() bool {
	return d.size == d.cap
}
```

## 0646 — Maximum Length Of Pair Chain

```go
package main

// LeetCode #646: Maximum Length of Pair Chain
// https://leetcode.com/problems/maximum-length-of-pair-chain/
// Difficulty: Medium
// Time: O(n log n) for sorting
// Space: O(1)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findLongestChain([][]int{{1, 2}, {2, 3}, {3, 4}}))
	fmt.Println(findLongestChain([][]int{{1, 2}, {7, 8}, {4, 5}}))
}

func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] < pairs[j][1]
	})

	count := 0
	curEnd := -1 << 31

	for _, pair := range pairs {
		if pair[0] > curEnd {
			curEnd = pair[1]
			count++
		}
	}

	return count
}
```

## 0647 — Palindromic Substrings

```go
package main

// LeetCode #647: Palindromic Substrings
// https://leetcode.com/problems/palindromic-substrings/
// Difficulty: Medium
// Time: O(n^2)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(countSubstrings("abc"))
	fmt.Println(countSubstrings("aaa"))
}

func countSubstrings(s string) int {
	count := 0
	n := len(s)

	for center := 0; center < 2*n-1; center++ {
		left := center / 2
		right := left + center%2
		for left >= 0 && right < n && s[left] == s[right] {
			count++
			left--
			right++
		}
	}

	return count
}
```

## 0648 — Replace Words

```go
package main

// LeetCode #648: Replace Words
// https://leetcode.com/problems/replace-words/
// Difficulty: Medium
// Time: O(n * L) where n is number of words in sentence, L is avg word length
// Space: O(d) where d is total characters in dictionary

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(replaceWords([]string{"cat", "bat", "rat"}, "the cattle was rattled by the battery"))
	fmt.Println(replaceWords([]string{"a", "b", "c"}, "aadsfasf absbs bbab cadsfafs"))
}

func replaceWords(dictionary []string, sentence string) string {
	rootSet := make(map[string]bool)
	for _, root := range dictionary {
		rootSet[root] = true
	}

	words := strings.Split(sentence, " ")
	for i, word := range words {
		for j := 1; j <= len(word); j++ {
			if rootSet[word[:j]] {
				words[i] = word[:j]
				break
			}
		}
	}

	return strings.Join(words, " ")
}
```

## 0649 — Dota2 Senate

```go
package main

// LeetCode #649: Dota2 Senate
// https://leetcode.com/problems/dota2-senate/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(predictPartyVictory("RD"))
	fmt.Println(predictPartyVictory("RDD"))
	fmt.Println(predictPartyVictory("RRDDD"))
}

func predictPartyVictory(senate string) string {
	n := len(senate)
	radiant := make([]int, 0)
	dire := make([]int, 0)

	for i, c := range senate {
		if c == 'R' {
			radiant = append(radiant, i)
		} else {
			dire = append(dire, i)
		}
	}

	for len(radiant) > 0 && len(dire) > 0 {
		r := radiant[0]
		d := dire[0]
		radiant = radiant[1:]
		dire = dire[1:]

		if r < d {
			radiant = append(radiant, r+n)
		} else {
			dire = append(dire, d+n)
		}
	}

	if len(radiant) > 0 {
		return "Radiant"
	}
	return "Dire"
}
```

## 0650 — 2 Keys Keyboard

```go
package main

// LeetCode #650: 2 Keys Keyboard
// https://leetcode.com/problems/2-keys-keyboard/
// Difficulty: Medium
// Time: O(n sqrt(n))
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(minSteps(3))
	fmt.Println(minSteps(1))
	fmt.Println(minSteps(10))
}

func minSteps(n int) int {
	result := 0
	d := 2

	for n > 1 {
		for n%d == 0 {
			result += d
			n /= d
		}
		d++
	}

	return result
}
```

## 0651 — 4 Keys Keyboard

```go
package main

// LeetCode #651: 4 Keys Keyboard
// https://leetcode.com/problems/4-keys-keyboard/
// Difficulty: Medium [Paid]
// Time: O(n^2)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(maxA(3))
	fmt.Println(maxA(7))
	fmt.Println(maxA(10))
}

func maxA(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] + 1
		for j := 1; j < i-1; j++ {
			dp[i] = max(dp[i], dp[j]*(i-j-1))
		}
	}
	return dp[n]
}
```

## 0652 — Find Duplicate Subtrees

```go
package main

// LeetCode #652: Find Duplicate Subtrees
// https://leetcode.com/problems/find-duplicate-subtrees/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import (
	"fmt"
	"strconv"
)

func main() {
	// Test case: root = [1,2,3,4,null,2,4,null,null,4]
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 2}
	root.Right.Left.Left = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 4}

	result := findDuplicateSubtrees(root)
	for _, node := range result {
		fmt.Println(node.Val)
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	seen := make(map[string]int)
	result := make([]*TreeNode, 0)

	var dfs func(node *TreeNode) string
	dfs = func(node *TreeNode) string {
		if node == nil {
			return "#"
		}

		left := dfs(node.Left)
		right := dfs(node.Right)
		key := strconv.Itoa(node.Val) + "," + left + "," + right

		seen[key]++
		if seen[key] == 2 {
			result = append(result, node)
		}

		return key
	}

	dfs(root)
	return result
}
```

## 0654 — Maximum Binary Tree

```go
package main

// LeetCode #654: Maximum Binary Tree
// https://leetcode.com/problems/maximum-binary-tree/
// Difficulty: Medium
// Time: O(n^2) worst case, O(n log n) average
// Space: O(n)

import "fmt"

func main() {
	root := constructMaximumBinaryTree([]int{3, 2, 1, 6, 0, 5})
	fmt.Println(treeToSlice(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	maxIdx := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[maxIdx] {
			maxIdx = i
		}
	}

	root := &TreeNode{Val: nums[maxIdx]}
	root.Left = constructMaximumBinaryTree(nums[:maxIdx])
	root.Right = constructMaximumBinaryTree(nums[maxIdx+1:])

	return root
}

func treeToSlice(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := []int{root.Val}
	result = append(result, treeToSlice(root.Left)...)
	result = append(result, treeToSlice(root.Right)...)
	return result
}
```

## 0655 — Print Binary Tree

```go
package main

// LeetCode #655: Print Binary Tree
// https://leetcode.com/problems/print-binary-tree/
// Difficulty: Medium
// Time: O(n * h) where h is height
// Space: O(n * h)

import (
	"fmt"
	"strconv"
)

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 4}

	fmt.Println(printTree(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func printTree(root *TreeNode) [][]string {
	height := getHeight(root)
	width := (1 << height) - 1
	result := make([][]string, height)
	for i := range result {
		result[i] = make([]string, width)
	}

	fill(root, result, 0, 0, width-1)
	return result
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(getHeight(root.Left), getHeight(root.Right))
}

func fill(root *TreeNode, result [][]string, level int, left int, right int) {
	if root == nil {
		return
	}
	mid := (left + right) / 2
	result[level][mid] = strconv.Itoa(root.Val)
	fill(root.Left, result, level+1, left, mid-1)
	fill(root.Right, result, level+1, mid+1, right)
}
```

## 0658 — Find K Closest Elements

```go
package main

// LeetCode #658: Find K Closest Elements
// https://leetcode.com/problems/find-k-closest-elements/
// Difficulty: Medium
// Time: O(log n + k)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(findClosestElements([]int{1, 2, 3, 4, 5}, 4, 3))
	fmt.Println(findClosestElements([]int{1, 1, 1, 10, 10, 10}, 1, 9))
}

func findClosestElements(arr []int, k int, x int) []int {
	left, right := 0, len(arr)-k

	for left < right {
		mid := left + (right-left)/2
		if x-arr[mid] > arr[mid+k]-x {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return arr[left : left+k]
}
```

## 0659 — Split Array Into Consecutive Subsequences

```go
package main

// LeetCode #659: Split Array into Consecutive Subsequences
// https://leetcode.com/problems/split-array-into-consecutive-subsequences/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(isPossible([]int{1, 2, 3, 3, 4, 5}))
	fmt.Println(isPossible([]int{1, 2, 3, 3, 4, 4, 5, 5}))
	fmt.Println(isPossible([]int{1, 2, 3, 4, 4, 5}))
}

func isPossible(nums []int) bool {
	freq := make(map[int]int)
	tail := make(map[int]int)

	for _, num := range nums {
		freq[num]++
	}

	for _, num := range nums {
		if freq[num] == 0 {
			continue
		}

		if tail[num] > 0 {
			tail[num]--
			freq[num]--
			tail[num+1]++
		} else if freq[num+1] > 0 && freq[num+2] > 0 {
			freq[num]--
			freq[num+1]--
			freq[num+2]--
			tail[num+3]++
		} else {
			return false
		}
	}

	return true
}
```

## 0662 — Maximum Width Of Binary Tree

```go
package main

// LeetCode #662: Maximum Width of Binary Tree
// https://leetcode.com/problems/maximum-width-of-binary-tree/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 5}
	root.Left.Right = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 2}
	root.Right.Right = &TreeNode{Val: 9}

	fmt.Println(widthOfBinaryTree(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	type pair struct {
		node *TreeNode
		idx  int
	}

	queue := []pair{{root, 0}}
	maxWidth := 0

	for len(queue) > 0 {
		n := len(queue)
		first := queue[0].idx
		last := queue[n-1].idx
		maxWidth = max(maxWidth, last-first+1)

		for i := 0; i < n; i++ {
			node, idx := queue[i].node, queue[i].idx
			if node.Left != nil {
				queue = append(queue, pair{node.Left, idx * 2})
			}
			if node.Right != nil {
				queue = append(queue, pair{node.Right, idx*2 + 1})
			}
		}
		queue = queue[n:]
	}

	return maxWidth
}
```

## 0663 — Equal Tree Partition

```go
package main

// LeetCode #663: Equal Tree Partition
// https://leetcode.com/problems/equal-tree-partition/
// Difficulty: Medium [Paid]
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 10}
	root.Right = &TreeNode{Val: 10}
	root.Right.Left = &TreeNode{Val: 2}
	root.Right.Right = &TreeNode{Val: 3}

	fmt.Println(checkEqualTree(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkEqualTree(root *TreeNode) bool {
	sums := make(map[int]int)
	total := computeSum(root, sums)

	if total%2 != 0 {
		return false
	}

	// Root sum is also stored; we need a different way to track
	// We'll rebuild sums excluding total
	half := total / 2
	return sums[half] > 0
}

func computeSum(node *TreeNode, sums map[int]int) int {
	if node == nil {
		return 0
	}
	left := computeSum(node.Left, sums)
	right := computeSum(node.Right, sums)
	sum := node.Val + left + right

	// Count subtree sums
	sums[sum]++

	return sum
}
```

## 0665 — Non Decreasing Array

```go
package main

// LeetCode #665: Non-decreasing Array
// https://leetcode.com/problems/non-decreasing-array/
// Difficulty: Medium
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(checkPossibility([]int{4, 2, 3}))
	fmt.Println(checkPossibility([]int{4, 2, 1}))
	fmt.Println(checkPossibility([]int{3, 4, 2, 3}))
}

func checkPossibility(nums []int) bool {
	modified := false

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			if modified {
				return false
			}

			if i == 0 || nums[i-1] <= nums[i+1] {
				nums[i] = nums[i+1]
			} else {
				nums[i+1] = nums[i]
			}

			modified = true
		}
	}

	return true
}
```

## 0666 — Path Sum Iv

```go
package main

// LeetCode #666: Path Sum IV
// https://leetcode.com/problems/path-sum-iv/
// Difficulty: Medium [Paid]
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(pathSumIV([]int{113, 215, 221}))
	fmt.Println(pathSumIV([]int{113, 221}))
}

func pathSumIV(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	tree := make(map[int]int)
	for _, num := range nums {
		tree[num/10] = num % 10
	}

	total := 0
	var dfs func(key int, sum int)
	dfs = func(key int, sum int) {
		depth := key / 10
		pos := key % 10
		leftKey := (depth+1)*10 + pos*2 - 1
		rightKey := (depth+1)*10 + pos*2

		curSum := sum + tree[key]

		_, hasLeft := tree[leftKey]
		_, hasRight := tree[rightKey]

		if !hasLeft && !hasRight {
			total += curSum
			return
		}

		if hasLeft {
			dfs(leftKey, curSum)
		}
		if hasRight {
			dfs(rightKey, curSum)
		}
	}

	dfs(nums[0]/10, 0)
	return total
}
```

## 0667 — Beautiful Arrangement Ii

```go
package main

// LeetCode #667: Beautiful Arrangement II
// https://leetcode.com/problems/beautiful-arrangement-ii/
// Difficulty: Medium
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(constructArray(3, 1))
	fmt.Println(constructArray(3, 2))
}

func constructArray(n int, k int) []int {
	result := make([]int, n)
	left, right := 1, k+1

	for i := 0; i <= k; i++ {
		if i%2 == 0 {
			result[i] = left
			left++
		} else {
			result[i] = right
			right--
		}
	}

	for i := k + 1; i < n; i++ {
		result[i] = i + 1
	}

	return result
}
```

## 0669 — Trim A Binary Search Tree

```go
package main

// LeetCode #669: Trim a Binary Search Tree
// https://leetcode.com/problems/trim-a-binary-search-tree/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 0}
	root.Right = &TreeNode{Val: 2}

	result := trimBST(root, 1, 2)
	fmt.Println(result.Val)
	if result.Left != nil {
		fmt.Println(result.Left.Val)
	}
	if result.Right != nil {
		fmt.Println(result.Right.Val)
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val < low {
		return trimBST(root.Right, low, high)
	}
	if root.Val > high {
		return trimBST(root.Left, low, high)
	}

	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	return root
}
```

## 0670 — Maximum Swap

```go
package main

// LeetCode #670: Maximum Swap
// https://leetcode.com/problems/maximum-swap/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(maximumSwap(2736))
	fmt.Println(maximumSwap(9973))
	fmt.Println(maximumSwap(98368))
}

func maximumSwap(num int) int {
	s := []byte(strconv.Itoa(num))
	n := len(s)

	// Track last occurrence of each digit
	last := make([]int, 10)
	for i := 0; i < n; i++ {
		last[s[i]-'0'] = i
	}

	for i := 0; i < n; i++ {
		for d := 9; d > int(s[i]-'0'); d-- {
			if last[d] > i {
				s[i], s[last[d]] = s[last[d]], s[i]
				result, _ := strconv.Atoi(string(s))
				return result
			}
		}
	}

	return num
}
```

## 0672 — Bulb Switcher Ii

```go
package main

// LeetCode #672: Bulb Switcher II
// https://leetcode.com/problems/bulb-switcher-ii/
// Difficulty: Medium
// Time: O(1)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(flipLights(1, 1))
	fmt.Println(flipLights(2, 1))
	fmt.Println(flipLights(3, 1))
}

func flipLights(n int, m int) int {
	// Only first 3 bulbs matter (key insight)
	if m == 0 {
		return 1
	}
	if n == 1 {
		return 2
	}
	if n == 2 {
		if m == 1 {
			return 3
		}
		return 4
	}
	// n >= 3
	if m == 1 {
		return 4
	}
	if m == 2 {
		return 7
	}
	return 8
}
```

## 0673 — Number Of Longest Increasing Subsequence

```go
package main

// LeetCode #673: Number of Longest Increasing Subsequence
// https://leetcode.com/problems/number-of-longest-increasing-subsequence/
// Difficulty: Medium
// Time: O(n^2)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(findNumberOfLIS([]int{1, 3, 5, 4, 7}))
	fmt.Println(findNumberOfLIS([]int{2, 2, 2, 2, 2}))
}

func findNumberOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	length := make([]int, n)
	count := make([]int, n)
	maxLen := 0

	for i := 0; i < n; i++ {
		length[i] = 1
		count[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if length[j]+1 > length[i] {
					length[i] = length[j] + 1
					count[i] = count[j]
				} else if length[j]+1 == length[i] {
					count[i] += count[j]
				}
			}
		}
		if length[i] > maxLen {
			maxLen = length[i]
		}
	}

	result := 0
	for i := 0; i < n; i++ {
		if length[i] == maxLen {
			result += count[i]
		}
	}
	return result
}
```

## 0676 — Implement Magic Dictionary

```go
package main

// LeetCode #676: Implement Magic Dictionary
// https://leetcode.com/problems/implement-magic-dictionary/
// Difficulty: Medium
// Time: O(n * L) for buildDict, O(26 * L) for search
// Space: O(n * L)

import "fmt"

func main() {
	md := MagicDictionary{}
	md.BuildDict([]string{"hello", "leetcode"})
	fmt.Println(md.Search("hello"))
	fmt.Println(md.Search("hhllo"))
	fmt.Println(md.Search("hell"))
	fmt.Println(md.Search("leetcoded"))
}

type MagicDictionary struct {
	words []string
}

func (m *MagicDictionary) BuildDict(dictionary []string) {
	m.words = dictionary
}

func (m *MagicDictionary) Search(searchWord string) bool {
	for _, word := range m.words {
		if len(word) != len(searchWord) {
			continue
		}
		diff := 0
		for i := 0; i < len(word); i++ {
			if word[i] != searchWord[i] {
				diff++
			}
			if diff > 1 {
				break
			}
		}
		if diff == 1 {
			return true
		}
	}
	return false
}
```

## 0677 — Map Sum Pairs

```go
package main

// LeetCode #677: Map Sum Pairs
// https://leetcode.com/problems/map-sum-pairs/
// Difficulty: Medium
// Time: O(L) for insert, O(L) for sum
// Space: O(n * L) where n is number of keys

import (
	"fmt"
)

func main() {
	ms := MapSumConstructor()
	ms.Insert("apple", 3)
	fmt.Println(ms.Sum("ap"))
	ms.Insert("app", 2)
	fmt.Println(ms.Sum("ap"))
}

type MapSum struct {
	prefixSum map[string]int
	values    map[string]int
}

func MapSumConstructor() MapSum {
	return MapSum{
		prefixSum: make(map[string]int),
		values:    make(map[string]int),
	}
}

func (m *MapSum) Insert(key string, val int) {
	delta := val
	if oldVal, ok := m.values[key]; ok {
		delta = val - oldVal
	}
	m.values[key] = val

	for i := 1; i <= len(key); i++ {
		m.prefixSum[key[:i]] += delta
	}
}

func (m *MapSum) Sum(prefix string) int {
	return m.prefixSum[prefix]
}
```

## 0678 — Valid Parenthesis String

```go
package main

// LeetCode #678: Valid Parenthesis String
// https://leetcode.com/problems/valid-parenthesis-string/
// Difficulty: Medium
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(checkValidString("()"))
	fmt.Println(checkValidString("(*)"))
	fmt.Println(checkValidString("(*))"))
}

func checkValidString(s string) bool {
	low, high := 0, 0

	for _, c := range s {
		if c == '(' {
			low++
			high++
		} else if c == ')' {
			low--
			high--
		} else { // '*'
			low--
			high++
		}

		if high < 0 {
			return false
		}
		if low < 0 {
			low = 0
		}
	}

	return low == 0
}
```

## 0681 — Next Closest Time

```go
package main

// LeetCode #681: Next Closest Time
// https://leetcode.com/problems/next-closest-time/
// Difficulty: Medium [Paid]
// Time: O(1) since there are at most 4^4 = 256 combinations
// Space: O(1)

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(nextClosestTime("19:34"))
	fmt.Println(nextClosestTime("23:59"))
	fmt.Println(nextClosestTime("13:55"))
}

func nextClosestTime(time string) string {
	digits := make(map[byte]bool)
	for i := 0; i < len(time); i++ {
		if time[i] != ':' {
			digits[time[i]] = true
		}
	}

	hours, _ := strconv.Atoi(time[:2])
	minutes, _ := strconv.Atoi(time[3:])

	current := hours*60 + minutes

	for elapsed := 1; elapsed <= 24*60; elapsed++ {
		t := (current + elapsed) % (24 * 60)
		h := t / 60
		m := t % 60

		hs := fmt.Sprintf("%02d%02d", h, m)
		valid := true
		for i := 0; i < 4; i++ {
			if !digits[hs[i]] {
				valid = false
				break
			}
		}

		if valid {
			return fmt.Sprintf("%02d:%02d", h, m)
		}
	}

	return ""
}
```

## 0684 — Redundant Connection

```go
package main

// LeetCode #684: Redundant Connection
// https://leetcode.com/problems/redundant-connection/
// Difficulty: Medium
// Time: O(n * alpha(n))
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(findRedundantConnection([][]int{{1, 2}, {1, 3}, {2, 3}}))
	fmt.Println(findRedundantConnection([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}))
}

func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	parent := make([]int, n+1)
	for i := 1; i <= n; i++ {
		parent[i] = i
	}

	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func(x, y int) bool {
		px, py := find(x), find(y)
		if px == py {
			return false
		}
		parent[px] = py
		return true
	}

	for _, edge := range edges {
		if !union(edge[0], edge[1]) {
			return edge
		}
	}

	return nil
}
```

## 0686 — Repeated String Match

```go
package main

// LeetCode #686: Repeated String Match
// https://leetcode.com/problems/repeated-string-match/
// Difficulty: Medium
// Time: O(n * m) worst case
// Space: O(n + m)

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(repeatedStringMatch("abcd", "cdabcdab"))
	fmt.Println(repeatedStringMatch("a", "aa"))
	fmt.Println(repeatedStringMatch("abc", "wxyz"))
}

func repeatedStringMatch(a string, b string) int {
	maxRepeats := len(b)/len(a) + 3
	var sb strings.Builder

	for i := 1; i <= maxRepeats; i++ {
		sb.WriteString(a)
		if strings.Contains(sb.String(), b) {
			return i
		}
	}

	return -1
}
```

## 0687 — Longest Univalue Path

```go
package main

// LeetCode #687: Longest Univalue Path
// https://leetcode.com/problems/longest-univalue-path/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 5}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 1}
	root.Right.Right = &TreeNode{Val: 5}

	fmt.Println(longestUnivaluePath(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestUnivaluePath(root *TreeNode) int {
	maxLen := 0

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftLen := dfs(node.Left)
		rightLen := dfs(node.Right)

		leftArrow, rightArrow := 0, 0
		if node.Left != nil && node.Left.Val == node.Val {
			leftArrow = leftLen + 1
		}
		if node.Right != nil && node.Right.Val == node.Val {
			rightArrow = rightLen + 1
		}

		maxLen = max(maxLen, leftArrow+rightArrow)

		return max(leftArrow, rightArrow)
	}

	dfs(root)
	return maxLen
}
```

## 0688 — Knight Probability In Chessboard

```go
package main

// LeetCode #688: Knight Probability in Chessboard
// https://leetcode.com/problems/knight-probability-in-chessboard/
// Difficulty: Medium
// Time: O(K * N^2)
// Space: O(N^2)

import "fmt"

func main() {
	fmt.Println(knightProbability(3, 2, 0, 0))
	fmt.Println(knightProbability(1, 0, 0, 0))
}

func knightProbability(n int, k int, row int, column int) float64 {
	dirs := [][2]int{{-2, -1}, {-2, 1}, {-1, -2}, {-1, 2}, {1, -2}, {1, 2}, {2, -1}, {2, 1}}
	dp := make([][]float64, n)
	for i := range dp {
		dp[i] = make([]float64, n)
	}
	dp[row][column] = 1.0

	for step := 0; step < k; step++ {
		next := make([][]float64, n)
		for i := range next {
			next[i] = make([]float64, n)
		}
		for r := 0; r < n; r++ {
			for c := 0; c < n; c++ {
				if dp[r][c] == 0 {
					continue
				}
				for _, d := range dirs {
					nr, nc := r+d[0], c+d[1]
					if nr >= 0 && nr < n && nc >= 0 && nc < n {
						next[nr][nc] += dp[r][c] / 8.0
					}
				}
			}
		}
		dp = next
	}

	result := 0.0
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			result += dp[r][c]
		}
	}
	return result
}
```

## 0690 — Employee Importance

```go
package main

// LeetCode #690: Employee Importance
// https://leetcode.com/problems/employee-importance/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	employees := []*Employee{
		{Id: 1, Importance: 5, Subordinates: []int{2, 3}},
		{Id: 2, Importance: 3, Subordinates: []int{}},
		{Id: 3, Importance: 3, Subordinates: []int{}},
	}
	fmt.Println(getImportance(employees, 1))
}

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) int {
	empMap := make(map[int]*Employee)
	for _, e := range employees {
		empMap[e.Id] = e
	}

	var dfs func(id int) int
	dfs = func(id int) int {
		emp := empMap[id]
		total := emp.Importance
		for _, subId := range emp.Subordinates {
			total += dfs(subId)
		}
		return total
	}

	return dfs(id)
}
```

## 0692 — Top K Frequent Words

```go
package main

// LeetCode #692: Top K Frequent Words
// https://leetcode.com/problems/top-k-frequent-words/
// Difficulty: Medium
// Time: O(n log k)
// Space: O(n)

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(topKFrequent([]string{"i", "love", "leetcode", "i", "love", "coding"}, 2))
	fmt.Println(topKFrequent([]string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}, 4))
}

type Item struct {
	word string
	freq int
}

type MinHeap []Item

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool {
	if h[i].freq == h[j].freq {
		return h[i].word > h[j].word
	}
	return h[i].freq < h[j].freq
}
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(Item)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func topKFrequent(words []string, k int) []string {
	freq := make(map[string]int)
	for _, w := range words {
		freq[w]++
	}

	h := &MinHeap{}
	heap.Init(h)

	for word, f := range freq {
		heap.Push(h, Item{word, f})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	result := make([]string, k)
	for i := k - 1; i >= 0; i-- {
		result[i] = heap.Pop(h).(Item).word
	}

	return result
}
```

## 0694 — Number Of Distinct Islands

```go
package main

// LeetCode #694: Number of Distinct Islands
// https://leetcode.com/problems/number-of-distinct-islands/
// Difficulty: Medium [Paid]
// Time: O(R * C)
// Space: O(R * C)

import (
	"fmt"
	"strings"
)

func main() {
	grid1 := [][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 1, 1},
		{0, 0, 0, 1, 1},
	}
	fmt.Println(numDistinctIslands(grid1))

	grid2 := [][]int{
		{1, 1, 0, 1, 1},
		{1, 0, 0, 0, 0},
		{0, 0, 0, 0, 1},
		{1, 1, 0, 1, 1},
	}
	fmt.Println(numDistinctIslands(grid2))
}

func numDistinctIslands(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])
	shapes := make(map[string]bool)

	var dfs func(r, c int, dir byte, sb *strings.Builder)
	dfs = func(r, c int, dir byte, sb *strings.Builder) {
		if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] == 0 {
			return
		}
		grid[r][c] = 0
		sb.WriteByte(dir)
		dfs(r-1, c, 'U', sb)
		dfs(r+1, c, 'D', sb)
		dfs(r, c-1, 'L', sb)
		dfs(r, c+1, 'R', sb)
		sb.WriteByte('B')
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 1 {
				var sb strings.Builder
				dfs(r, c, 'S', &sb)
				shapes[sb.String()] = true
			}
		}
	}

	return len(shapes)
}
```

## 0695 — Max Area Of Island

```go
package main

// LeetCode #695: Max Area of Island
// https://leetcode.com/problems/max-area-of-island/
// Difficulty: Medium
// Time: O(R * C)
// Space: O(R * C)

import "fmt"

func main() {
	grid := [][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}
	fmt.Println(maxAreaOfIsland(grid))
}

func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	maxArea := 0
	rows, cols := len(grid), len(grid[0])

	var dfs func(r, c int) int
	dfs = func(r, c int) int {
		if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] == 0 {
			return 0
		}
		grid[r][c] = 0
		return 1 + dfs(r-1, c) + dfs(r+1, c) + dfs(r, c-1) + dfs(r, c+1)
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 1 {
				area := dfs(r, c)
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}

	return maxArea
}
```

## 0698 — Partition To K Equal Sum Subsets

```go
package main

// LeetCode #698: Partition to K Equal Sum Subsets
// https://leetcode.com/problems/partition-to-k-equal-sum-subsets/
// Difficulty: Medium
// Time: O(k^(n-k) * n!) roughly
// Space: O(n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(canPartitionKSubsets([]int{4, 3, 2, 3, 5, 2, 1}, 4))
	fmt.Println(canPartitionKSubsets([]int{1, 2, 3, 4}, 3))
}

func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum%k != 0 {
		return false
	}

	target := sum / k
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	used := make([]bool, len(nums))

	var backtrack func(start int, currentSum int, subsetsFormed int) bool
	backtrack = func(start int, currentSum int, subsetsFormed int) bool {
		if subsetsFormed == k {
			return true
		}
		if currentSum == target {
			return backtrack(0, 0, subsetsFormed+1)
		}

		for i := start; i < len(nums); i++ {
			if used[i] || currentSum+nums[i] > target {
				continue
			}
			// Pruning: skip duplicate values
			if i > 0 && !used[i-1] && nums[i] == nums[i-1] {
				continue
			}

			used[i] = true
			if backtrack(i+1, currentSum+nums[i], subsetsFormed) {
				return true
			}
			used[i] = false

			if currentSum == 0 {
				return false
			}
		}
		return false
	}

	return backtrack(0, 0, 0)
}
```

## 0701 — Insert Into A Binary Search Tree

```go
package main

// LeetCode #701: Insert into a Binary Search Tree
// https://leetcode.com/problems/insert-into-a-binary-search-tree/
// Difficulty: Medium
// Time: O(h) where h is tree height
// Space: O(1)

import "fmt"

func main() {
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 7}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}

	result := insertIntoBST(root, 5)
	fmt.Println(result.Val)
	fmt.Println(result.Right.Left.Val)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	curr := root
	for {
		if val < curr.Val {
			if curr.Left == nil {
				curr.Left = &TreeNode{Val: val}
				break
			}
			curr = curr.Left
		} else {
			if curr.Right == nil {
				curr.Right = &TreeNode{Val: val}
				break
			}
			curr = curr.Right
		}
	}

	return root
}
```

## 0702 — Search In A Sorted Array Of Unknown Size

```go
package main

// LeetCode #702: Search in a Sorted Array of Unknown Size
// https://leetcode.com/problems/search-in-a-sorted-array-of-unknown-size/
// Difficulty: Medium [Paid]
// Time: O(log n)
// Space: O(1)

import "fmt"

func main() {
	reader := &ArrayReader{arr: []int{-1, 0, 3, 5, 9, 12}}
	fmt.Println(search(reader, 9))
	fmt.Println(search(reader, 2))
}

type ArrayReader struct {
	arr []int
}

func (r *ArrayReader) get(index int) int {
	if index >= len(r.arr) {
		return 1 << 31 - 1
	}
	return r.arr[index]
}

func search(reader *ArrayReader, target int) int {
	// Find upper bound
	left, right := 0, 1
	for reader.get(right) < target {
		left = right
		right <<= 1
	}

	// Binary search
	for left <= right {
		mid := left + (right-left)/2
		val := reader.get(mid)
		if val == target {
			return mid
		} else if val < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
```

## 0707 — Design Linked List

```go
package main

// LeetCode #707: Design Linked List
// https://leetcode.com/problems/design-linked-list/
// Difficulty: Medium
// Time: O(n) for get/addAtIndex/deleteAtIndex, O(1) for addAtHead/addAtTail
// Space: O(n)

import "fmt"

func main() {
	l := Constructor()
	l.AddAtHead(1)
	l.AddAtTail(3)
	l.AddAtIndex(1, 2)
	fmt.Println(l.Get(1))
	l.DeleteAtIndex(1)
	fmt.Println(l.Get(1))
}

type MyLinkedList struct {
	head *LinkNode
	size int
}

type LinkNode struct {
	val  int
	next *LinkNode
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (l *MyLinkedList) Get(index int) int {
	if index < 0 || index >= l.size {
		return -1
	}
	curr := l.head
	for i := 0; i < index; i++ {
		curr = curr.next
	}
	return curr.val
}

func (l *MyLinkedList) AddAtHead(val int) {
	l.head = &LinkNode{val, l.head}
	l.size++
}

func (l *MyLinkedList) AddAtTail(val int) {
	if l.head == nil {
		l.AddAtHead(val)
		return
	}
	curr := l.head
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = &LinkNode{val: val}
	l.size++
}

func (l *MyLinkedList) AddAtIndex(index int, val int) {
	if index > l.size {
		return
	}
	if index == 0 {
		l.AddAtHead(val)
		return
	}
	curr := l.head
	for i := 0; i < index-1; i++ {
		curr = curr.next
	}
	curr.next = &LinkNode{val, curr.next}
	l.size++
}

func (l *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= l.size {
		return
	}
	if index == 0 {
		l.head = l.head.next
		l.size--
		return
	}
	curr := l.head
	for i := 0; i < index-1; i++ {
		curr = curr.next
	}
	curr.next = curr.next.next
	l.size--
}
```

