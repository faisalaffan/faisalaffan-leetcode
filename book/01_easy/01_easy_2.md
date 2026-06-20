# Easy (Mudah) — Problem 0414–0824

## 0414 — Third Maximum Number

```go
package main

// LeetCode #414: Third Maximum Number
// https://leetcode.com/problems/third-maximum-number/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func ThirdMaximumNumber(nums []int) int {
	var max1, max2, max3 *int
	for _, v := range nums {
		val := v
		if max1 != nil && val == *max1 {
			continue
		}
		if max1 == nil || val > *max1 {
			max3 = max2
			max2 = max1
			max1 = &val
		} else if max2 == nil || val > *max2 {
			max3 = max2
			max2 = &val
		} else if max3 == nil || val > *max3 {
			max3 = &val
		}
	}
	if max3 != nil {
		return *max3
	}
	return *max1
}

func main() {
	fmt.Println(ThirdMaximumNumber([]int{3, 2, 1}))
	fmt.Println(ThirdMaximumNumber([]int{1, 2}))
	fmt.Println(ThirdMaximumNumber([]int{2, 2, 3, 1}))
}
```

## 0415 — Add Strings

```go
package main

// LeetCode #415: Add Strings
// https://leetcode.com/problems/add-strings/
// Difficulty: Easy

import "fmt"

// Time: O(n+m), Space: O(max(n,m))
func AddStrings(num1, num2 string) string {
	i, j := len(num1)-1, len(num2)-1
	carry := 0
	var result []byte
	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry
		if i >= 0 {
			sum += int(num1[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(num2[j] - '0')
			j--
		}
		carry = sum / 10
		result = append([]byte{byte(sum%10 + '0')}, result...)
	}
	return string(result)
}

func main() {
	fmt.Println(AddStrings("11", "123"))
	fmt.Println(AddStrings("456", "77"))
	fmt.Println(AddStrings("0", "0"))
}
```

## 0422 — Valid Word Square

```go
package main

// LeetCode #422: Valid Word Square
// https://leetcode.com/problems/valid-word-square/
// Difficulty: Easy [Paid]

import "fmt"

// Time: O(n*m), Space: O(1)
func ValidWordSquare(words []string) bool {
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ {
			if j >= len(words) || i >= len(words[j]) || words[i][j] != words[j][i] {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(ValidWordSquare([]string{"abcd", "bnrt", "crmy", "dtyx"}))
	fmt.Println(ValidWordSquare([]string{"abcd", "bnrt", "crm", "dt"}))
	fmt.Println(ValidWordSquare([]string{"ball", "area", "lead", "lady"}))
}
```

## 0434 — Number Of Segments In A String

```go
package main

// LeetCode #434: Number of Segments in a String
// https://leetcode.com/problems/number-of-segments-in-a-string/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func NumberOfSegmentsInAString(s string) int {
	count := 0
	inSegment := false
	for _, c := range s {
		if c != ' ' && !inSegment {
			count++
			inSegment = true
		} else if c == ' ' {
			inSegment = false
		}
	}
	return count
}

func main() {
	fmt.Println(NumberOfSegmentsInAString("Hello, my name is John"))
	fmt.Println(NumberOfSegmentsInAString("Hello"))
	fmt.Println(NumberOfSegmentsInAString(""))
}
```

## 0441 — Arranging Coins

```go
package main

// LeetCode #441: Arranging Coins
// https://leetcode.com/problems/arranging-coins/
// Difficulty: Easy

import "fmt"

// Time: O(log n), Space: O(1)
func ArrangingCoins(n int) int {
	lo, hi := 1, n
	for lo <= hi {
		mid := lo + (hi-lo)/2
		sum := mid * (mid + 1) / 2
		if sum == n {
			return mid
		} else if sum < n {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return hi
}

func main() {
	fmt.Println(ArrangingCoins(5))
	fmt.Println(ArrangingCoins(8))
	fmt.Println(ArrangingCoins(1))
}
```

## 0448 — Find All Numbers Disappeared In An Array

```go
package main

// LeetCode #448: Find All Numbers Disappeared in an Array
// https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func FindAllNumbersDisappearedInAnArray(nums []int) []int {
	for _, v := range nums {
		idx := v
		if idx < 0 {
			idx = -idx
		}
		idx--
		if nums[idx] > 0 {
			nums[idx] = -nums[idx]
		}
	}
	var result []int
	for i, v := range nums {
		if v > 0 {
			result = append(result, i+1)
		}
	}
	return result
}

func main() {
	fmt.Println(FindAllNumbersDisappearedInAnArray([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	fmt.Println(FindAllNumbersDisappearedInAnArray([]int{1, 1}))
}
```

## 0455 — Assign Cookies

```go
package main

// LeetCode #455: Assign Cookies
// https://leetcode.com/problems/assign-cookies/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

// Time: O(n log n + m log m), Space: O(1)
func AssignCookies(g, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	i, j := 0, 0
	for i < len(g) && j < len(s) {
		if s[j] >= g[i] {
			i++
		}
		j++
	}
	return i
}

func main() {
	fmt.Println(AssignCookies([]int{1, 2, 3}, []int{1, 1}))
	fmt.Println(AssignCookies([]int{1, 2}, []int{1, 2, 3}))
}
```

## 0459 — Repeated Substring Pattern

```go
package main

// LeetCode #459: Repeated Substring Pattern
// https://leetcode.com/problems/repeated-substring-pattern/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

// Time: O(n), Space: O(n)
func RepeatedSubstringPattern(s string) bool {
	t := s + s
	return strings.Contains(t[1:len(t)-1], s)
}

func main() {
	fmt.Println(RepeatedSubstringPattern("abab"))
	fmt.Println(RepeatedSubstringPattern("aba"))
	fmt.Println(RepeatedSubstringPattern("abcabcabcabc"))
}
```

## 0461 — Hamming Distance

```go
package main

// LeetCode #461: Hamming Distance
// https://leetcode.com/problems/hamming-distance/
// Difficulty: Easy

import "fmt"

// Time: O(1), Space: O(1)
func HammingDistance(x, y int) int {
	xor := x ^ y
	count := 0
	for xor > 0 {
		xor &= xor - 1
		count++
	}
	return count
}

func main() {
	fmt.Println(HammingDistance(1, 4))
	fmt.Println(HammingDistance(3, 1))
}
```

## 0463 — Island Perimeter

```go
package main

// LeetCode #463: Island Perimeter
// https://leetcode.com/problems/island-perimeter/
// Difficulty: Easy

import "fmt"

// Time: O(m*n), Space: O(1)
func IslandPerimeter(grid [][]int) int {
	perimeter := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				perimeter += 4
				if i > 0 && grid[i-1][j] == 1 {
					perimeter -= 2
				}
				if j > 0 && grid[i][j-1] == 1 {
					perimeter -= 2
				}
			}
		}
	}
	return perimeter
}

func main() {
	fmt.Println(IslandPerimeter([][]int{
		{0, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 1, 0, 0},
		{1, 1, 0, 0},
	}))
	fmt.Println(IslandPerimeter([][]int{{1}}))
}
```

## 0476 — Number Complement

```go
package main

// LeetCode #476: Number Complement
// https://leetcode.com/problems/number-complement/
// Difficulty: Easy

import "fmt"

// Time: O(1), Space: O(1)
func NumberComplement(num int) int {
	mask := ^0
	for num&mask != 0 {
		mask <<= 1
	}
	return ^num & ^mask
}

func main() {
	fmt.Println(NumberComplement(5))
	fmt.Println(NumberComplement(1))
	fmt.Println(NumberComplement(2))
}
```

## 0482 — License Key Formatting

```go
package main

// LeetCode #482: License Key Formatting
// https://leetcode.com/problems/license-key-formatting/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(n)
func LicenseKeyFormatting(s string, k int) string {
	var result []byte
	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '-' {
			continue
		}
		if count == k {
			result = append([]byte{'-'}, result...)
			count = 0
		}
		c := s[i]
		if c >= 'a' && c <= 'z' {
			c -= 32
		}
		result = append([]byte{c}, result...)
		count++
	}
	return string(result)
}

func main() {
	fmt.Println(LicenseKeyFormatting("5F3Z-2e-9-w", 4))
	fmt.Println(LicenseKeyFormatting("2-5g-3-J", 2))
}
```

## 0485 — Max Consecutive Ones

```go
package main

// LeetCode #485: Max Consecutive Ones
// https://leetcode.com/problems/max-consecutive-ones/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func MaxConsecutiveOnes(nums []int) int {
	maxCount, count := 0, 0
	for _, v := range nums {
		if v == 1 {
			count++
			if count > maxCount {
				maxCount = count
			}
		} else {
			count = 0
		}
	}
	return maxCount
}

func main() {
	fmt.Println(MaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1}))
	fmt.Println(MaxConsecutiveOnes([]int{1, 0, 1, 1, 0, 1}))
}
```

## 0492 — Construct The Rectangle

```go
package main

// LeetCode #492: Construct the Rectangle
// https://leetcode.com/problems/construct-the-rectangle/
// Difficulty: Easy

import "fmt"

// Time: O(sqrt(n)), Space: O(1)
func ConstructTheRectangle(area int) []int {
	w := 1
	for i := 1; i*i <= area; i++ {
		if area%i == 0 {
			w = i
		}
	}
	return []int{area / w, w}
}

func main() {
	fmt.Println(ConstructTheRectangle(4))
	fmt.Println(ConstructTheRectangle(37))
	fmt.Println(ConstructTheRectangle(122122))
}
```

## 0495 — Teemo Attacking

```go
package main

// LeetCode #495: Teemo Attacking
// https://leetcode.com/problems/teemo-attacking/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func TeemoAttacking(timeSeries []int, duration int) int {
	total := 0
	for i := 0; i < len(timeSeries)-1; i++ {
		total += min(duration, timeSeries[i+1]-timeSeries[i])
	}
	if len(timeSeries) > 0 {
		total += duration
	}
	return total
}

func main() {
	fmt.Println(TeemoAttacking([]int{1, 4}, 2))
	fmt.Println(TeemoAttacking([]int{1, 2}, 2))
}
```

## 0496 — Next Greater Element I

```go
package main

// LeetCode #496: Next Greater Element I
// https://leetcode.com/problems/next-greater-element-i/
// Difficulty: Easy

import "fmt"

// Time: O(n+m), Space: O(m)
func NextGreaterElementI(nums1, nums2 []int) []int {
	nextGreater := make(map[int]int)
	var stack []int
	for _, v := range nums2 {
		for len(stack) > 0 && stack[len(stack)-1] < v {
			nextGreater[stack[len(stack)-1]] = v
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, v)
	}
	result := make([]int, len(nums1))
	for i, v := range nums1 {
		if val, ok := nextGreater[v]; ok {
			result[i] = val
		} else {
			result[i] = -1
		}
	}
	return result
}

func main() {
	fmt.Println(NextGreaterElementI([]int{4, 1, 2}, []int{1, 3, 4, 2}))
	fmt.Println(NextGreaterElementI([]int{2, 4}, []int{1, 2, 3, 4}))
}
```

## 0500 — Keyboard Row

```go
package main

// LeetCode #500: Keyboard Row
// https://leetcode.com/problems/keyboard-row/
// Difficulty: Easy

import "fmt"

// Time: O(n*k), Space: O(n)
func KeyboardRow(words []string) []string {
	rows := []string{"qwertyuiop", "asdfghjkl", "zxcvbnm"}
	rowMap := make(map[byte]int)
	for i, row := range rows {
		for j := 0; j < len(row); j++ {
			rowMap[row[j]] = i
		}
	}
	var result []string
	for _, word := range words {
		if len(word) == 0 {
			continue
		}
		targetRow := rowMap[toLower(word[0])]
		sameRow := true
		for i := 1; i < len(word); i++ {
			if rowMap[toLower(word[i])] != targetRow {
				sameRow = false
				break
			}
		}
		if sameRow {
			result = append(result, word)
		}
	}
	return result
}

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}

func main() {
	fmt.Println(KeyboardRow([]string{"Hello", "Alaska", "Dad", "Peace"}))
	fmt.Println(KeyboardRow([]string{"omk"}))
	fmt.Println(KeyboardRow([]string{"adsdf", "sfd"}))
}
```

## 0501 — Find Mode In Binary Search Tree

```go
package main

// LeetCode #501: Find Mode in Binary Search Tree
// https://leetcode.com/problems/find-mode-in-binary-search-tree/
// Difficulty: Easy

import "fmt"

// TreeNode represents a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n), Space: O(n)
func FindModeInBinarySearchTree(root *TreeNode) []int {
	var result []int
	maxCount, currentCount := 0, 0
	var prev *int

	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		if prev != nil && *prev == node.Val {
			currentCount++
		} else {
			currentCount = 1
		}
		prev = &node.Val
		if currentCount > maxCount {
			maxCount = currentCount
			result = []int{node.Val}
		} else if currentCount == maxCount {
			result = append(result, node.Val)
		}
		inorder(node.Right)
	}
	inorder(root)
	return result
}

func main() {
	// Test: [1,null,2,2]
	root1 := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 2},
		},
	}
	fmt.Println(FindModeInBinarySearchTree(root1))

	// Test: [0]
	root2 := &TreeNode{Val: 0}
	fmt.Println(FindModeInBinarySearchTree(root2))
}
```

## 0504 — Base 7

```go
package main

// LeetCode #504: Base 7
// https://leetcode.com/problems/base-7/
// Difficulty: Easy

import "fmt"

// Time: O(log n), Space: O(log n)
func BaseSeven(num int) string {
	if num == 0 {
		return "0"
	}
	negative := num < 0
	if negative {
		num = -num
	}
	var result []byte
	for num > 0 {
		result = append([]byte{byte('0' + num%7)}, result...)
		num /= 7
	}
	if negative {
		result = append([]byte{'-'}, result...)
	}
	return string(result)
}

func main() {
	fmt.Println(BaseSeven(100))
	fmt.Println(BaseSeven(-7))
	fmt.Println(BaseSeven(0))
}
```

## 0506 — Relative Ranks

```go
package main

// LeetCode #506: Relative Ranks
// https://leetcode.com/problems/relative-ranks/
// Difficulty: Easy

import (
	"fmt"
	"sort"
	"strconv"
)

// Time: O(n log n), Space: O(n)
func RelativeRanks(score []int) []string {
	sorted := make([]int, len(score))
	copy(sorted, score)
	sort.Sort(sort.Reverse(sort.IntSlice(sorted)))
	rank := make(map[int]string)
	for i, s := range sorted {
		switch i {
		case 0:
			rank[s] = "Gold Medal"
		case 1:
			rank[s] = "Silver Medal"
		case 2:
			rank[s] = "Bronze Medal"
		default:
			rank[s] = strconv.Itoa(i + 1)
		}
	}
	result := make([]string, len(score))
	for i, s := range score {
		result[i] = rank[s]
	}
	return result
}

func main() {
	fmt.Println(RelativeRanks([]int{5, 4, 3, 2, 1}))
	fmt.Println(RelativeRanks([]int{10, 3, 8, 9, 4}))
}
```

## 0507 — Perfect Number

```go
package main

// LeetCode #507: Perfect Number
// https://leetcode.com/problems/perfect-number/
// Difficulty: Easy

import "fmt"

// Time: O(sqrt(n)), Space: O(1)
func PerfectNumber(num int) bool {
	if num <= 1 {
		return false
	}
	sum := 1
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			sum += i
			if i != num/i {
				sum += num / i
			}
		}
	}
	return sum == num
}

func main() {
	fmt.Println(PerfectNumber(28))
	fmt.Println(PerfectNumber(7))
	fmt.Println(PerfectNumber(6))
}
```

## 0509 — Fibonacci Number

```go
package main

// LeetCode #509: Fibonacci Number
// https://leetcode.com/problems/fibonacci-number/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func FibonacciNumber(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func main() {
	fmt.Println(FibonacciNumber(2))
	fmt.Println(FibonacciNumber(3))
	fmt.Println(FibonacciNumber(4))
}
```

## 0511 — Game Play Analysis I

```go
package main

// LeetCode #511: Game Play Analysis I
// https://leetcode.com/problems/game-play-analysis-i/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(n)
func GamePlayAnalysisI() string {
	return "SELECT player_id, MIN(event_date) AS first_login FROM Activity GROUP BY player_id"
}

func main() {
	fmt.Println(GamePlayAnalysisI())
}
```

## 0512 — Game Play Analysis Ii

```go
package main

// LeetCode #512: Game Play Analysis II
// https://leetcode.com/problems/game-play-analysis-ii/
// Difficulty: Easy [Paid]

import "fmt"

// Time: O(n), Space: O(n)
func GamePlayAnalysisIi() string {
	return "SELECT player_id, device_id FROM Activity WHERE (player_id, event_date) IN (SELECT player_id, MIN(event_date) FROM Activity GROUP BY player_id)"
}

func main() {
	fmt.Println(GamePlayAnalysisIi())
}
```

## 0520 — Detect Capital

```go
package main

// LeetCode #520: Detect Capital
// https://leetcode.com/problems/detect-capital/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func DetectCapital(word string) bool {
	upperCount := 0
	for _, c := range word {
		if c >= 'A' && c <= 'Z' {
			upperCount++
		}
	}
	return upperCount == len(word) || upperCount == 0 || (upperCount == 1 && word[0] >= 'A' && word[0] <= 'Z')
}

func main() {
	fmt.Println(DetectCapital("USA"))
	fmt.Println(DetectCapital("FlaG"))
	fmt.Println(DetectCapital("Google"))
}
```

## 0521 — Longest Uncommon Subsequence I

```go
package main

// LeetCode #521: Longest Uncommon Subsequence I
// https://leetcode.com/problems/longest-uncommon-subsequence-i/
// Difficulty: Easy

import "fmt"

// Time: O(1), Space: O(1)
func LongestUncommonSubsequenceI(a, b string) int {
	if a == b {
		return -1
	}
	if len(a) > len(b) {
		return len(a)
	}
	return len(b)
}

func main() {
	fmt.Println(LongestUncommonSubsequenceI("aba", "cdc"))
	fmt.Println(LongestUncommonSubsequenceI("aaa", "bbb"))
	fmt.Println(LongestUncommonSubsequenceI("aaa", "aaa"))
}
```

## 0530 — Minimum Absolute Difference In Bst

```go
package main

// LeetCode #530: Minimum Absolute Difference in BST
// https://leetcode.com/problems/minimum-absolute-difference-in-bst/
// Difficulty: Easy

import (
	"fmt"
	"math"
)

// TreeNode represents a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n), Space: O(h)
func MinimumAbsoluteDifferenceInBst(root *TreeNode) int {
	minDiff := math.MaxInt32
	var prev *int
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		if prev != nil {
			diff := node.Val - *prev
			if diff < minDiff {
				minDiff = diff
			}
		}
		prev = &node.Val
		inorder(node.Right)
	}
	inorder(root)
	return minDiff
}

func main() {
	// Test: [4,2,6,1,3]
	root1 := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{Val: 6},
	}
	fmt.Println(MinimumAbsoluteDifferenceInBst(root1))

	// Test: [1,0,48,null,null,12,49]
	root2 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 0},
		Right: &TreeNode{
			Val:   48,
			Left:  &TreeNode{Val: 12},
			Right: &TreeNode{Val: 49},
		},
	}
	fmt.Println(MinimumAbsoluteDifferenceInBst(root2))
}
```

## 0541 — Reverse String Ii

```go
package main

// LeetCode #541: Reverse String II
// https://leetcode.com/problems/reverse-string-ii/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(n)
func ReverseStringIi(s string, k int) string {
	b := []byte(s)
	for i := 0; i < len(b); i += 2 * k {
		lo, hi := i, i+k-1
		if hi >= len(b) {
			hi = len(b) - 1
		}
		for lo < hi {
			b[lo], b[hi] = b[hi], b[lo]
			lo++
			hi--
		}
	}
	return string(b)
}

func main() {
	fmt.Println(ReverseStringIi("abcdefg", 2))
	fmt.Println(ReverseStringIi("abcd", 2))
}
```

## 0543 — Diameter Of Binary Tree

```go
package main

// LeetCode #543: Diameter of Binary Tree
// https://leetcode.com/problems/diameter-of-binary-tree/
// Difficulty: Easy

import "fmt"

// TreeNode represents a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n), Space: O(h)
func DiameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		if left+right > maxDiameter {
			maxDiameter = left + right
		}
		if left > right {
			return left + 1
		}
		return right + 1
	}
	dfs(root)
	return maxDiameter
}

func main() {
	// Test: [1,2,3,4,5]
	root1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{Val: 3},
	}
	fmt.Println(DiameterOfBinaryTree(root1))

	// Test: [1,2]
	root2 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	fmt.Println(DiameterOfBinaryTree(root2))
}
```

## 0551 — Student Attendance Record I

```go
package main

// LeetCode #551: Student Attendance Record I
// https://leetcode.com/problems/student-attendance-record-i/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func StudentAttendanceRecordI(s string) bool {
	absences, lateStreak := 0, 0
	for _, c := range s {
		if c == 'A' {
			absences++
			lateStreak = 0
			if absences >= 2 {
				return false
			}
		} else if c == 'L' {
			lateStreak++
			if lateStreak >= 3 {
				return false
			}
		} else {
			lateStreak = 0
		}
	}
	return true
}

func main() {
	fmt.Println(StudentAttendanceRecordI("PPALLP"))
	fmt.Println(StudentAttendanceRecordI("PPALLL"))
}
```

## 0557 — Reverse Words In A String Iii

```go
package main

// LeetCode #557: Reverse Words in a String III
// https://leetcode.com/problems/reverse-words-in-a-string-iii/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(n)
func ReverseWordsInAStringIii(s string) string {
	b := []byte(s)
	start := 0
	for i := 0; i <= len(b); i++ {
		if i == len(b) || b[i] == ' ' {
			for lo, hi := start, i-1; lo < hi; lo, hi = lo+1, hi-1 {
				b[lo], b[hi] = b[hi], b[lo]
			}
			start = i + 1
		}
	}
	return string(b)
}

func main() {
	fmt.Println(ReverseWordsInAStringIii("Let's take LeetCode contest"))
	fmt.Println(ReverseWordsInAStringIii("Mr Ding"))
}
```

## 0559 — Maximum Depth Of N Ary Tree

```go
package main

// LeetCode #559: Maximum Depth of N-ary Tree
// https://leetcode.com/problems/maximum-depth-of-n-ary-tree/
// Difficulty: Easy

import "fmt"

// Node represents an N-ary tree node.
type Node struct {
	Val      int
	Children []*Node
}

// Time: O(n), Space: O(h)
func MaximumDepthOfNAryTree(root *Node) int {
	if root == nil {
		return 0
	}
	maxDepth := 0
	for _, child := range root.Children {
		if depth := MaximumDepthOfNAryTree(child); depth > maxDepth {
			maxDepth = depth
		}
	}
	return maxDepth + 1
}

func main() {
	// Test: [1,null,3,2,4,null,5,6]
	root1 := &Node{
		Val: 1,
		Children: []*Node{
			{Val: 3, Children: []*Node{
				{Val: 5},
				{Val: 6},
			}},
			{Val: 2},
			{Val: 4},
		},
	}
	fmt.Println(MaximumDepthOfNAryTree(root1))

	// Test: single node
	root2 := &Node{Val: 1}
	fmt.Println(MaximumDepthOfNAryTree(root2))
}
```

## 0561 — Array Partition

```go
package main

// LeetCode #561: Array Partition
// https://leetcode.com/problems/array-partition/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

// Time: O(n log n), Space: O(1)
func ArrayPartition(nums []int) int {
	sort.Ints(nums)
	sum := 0
	for i := 0; i < len(nums); i += 2 {
		sum += nums[i]
	}
	return sum
}

func main() {
	fmt.Println(ArrayPartition([]int{1, 4, 3, 2}))
	fmt.Println(ArrayPartition([]int{6, 2, 6, 5, 1, 2}))
}
```

## 0563 — Binary Tree Tilt

```go
package main

// LeetCode #563: Binary Tree Tilt
// https://leetcode.com/problems/binary-tree-tilt/
// Difficulty: Easy

import "fmt"

// TreeNode represents a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n), Space: O(h)
func BinaryTreeTilt(root *TreeNode) int {
	totalTilt := 0
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		tilt := left - right
		if tilt < 0 {
			tilt = -tilt
		}
		totalTilt += tilt
		return left + right + node.Val
	}
	dfs(root)
	return totalTilt
}

func main() {
	// Test: [1,2,3]
	root1 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}
	fmt.Println(BinaryTreeTilt(root1))

	// Test: [4,2,9,3,5,null,7]
	root2 := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val:   9,
			Right: &TreeNode{Val: 7},
		},
	}
	fmt.Println(BinaryTreeTilt(root2))
}
```

## 0566 — Reshape The Matrix

```go
package main

// LeetCode #566: Reshape the Matrix
// https://leetcode.com/problems/reshape-the-matrix/
// Difficulty: Easy

import "fmt"

// Time: O(m*n), Space: O(m*n)
func ReshapeTheMatrix(mat [][]int, r, c int) [][]int {
	m, n := len(mat), len(mat[0])
	if m*n != r*c {
		return mat
	}
	result := make([][]int, r)
	for i := range result {
		result[i] = make([]int, c)
	}
	for i := 0; i < m*n; i++ {
		result[i/c][i%c] = mat[i/n][i%n]
	}
	return result
}

func main() {
	fmt.Println(ReshapeTheMatrix([][]int{{1, 2}, {3, 4}}, 1, 4))
	fmt.Println(ReshapeTheMatrix([][]int{{1, 2}, {3, 4}}, 2, 4))
}
```

## 0572 — Subtree Of Another Tree

```go
package main

// LeetCode #572: Subtree of Another Tree
// https://leetcode.com/problems/subtree-of-another-tree/
// Difficulty: Easy

import "fmt"

// TreeNode represents a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// Time: O(m*n), Space: O(h)
func SubtreeOfAnotherTree(root, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	if isSameTree(root, subRoot) {
		return true
	}
	return SubtreeOfAnotherTree(root.Left, subRoot) || SubtreeOfAnotherTree(root.Right, subRoot)
}

func main() {
	// Test: root=[3,4,5,1,2], subRoot=[4,1,2]
	root1 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   4,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 2},
		},
		Right: &TreeNode{Val: 5},
	}
	sub1 := &TreeNode{
		Val:   4,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 2},
	}
	fmt.Println(SubtreeOfAnotherTree(root1, sub1))

	// Test: root=[3,4,5,1,2,null,null,null,null,0], subRoot=[4,1,2]
	root2 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:  4,
			Left: &TreeNode{Val: 1},
			Right: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 0},
			},
		},
		Right: &TreeNode{Val: 5},
	}
	sub2 := &TreeNode{
		Val:   4,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 2},
	}
	fmt.Println(SubtreeOfAnotherTree(root2, sub2))
}
```

## 0575 — Distribute Candies

```go
package main

// LeetCode #575: Distribute Candies
// https://leetcode.com/problems/distribute-candies/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(n)
func DistributeCandies(candyType []int) int {
	types := make(map[int]bool)
	for _, c := range candyType {
		types[c] = true
	}
	maxAllowed := len(candyType) / 2
	if len(types) < maxAllowed {
		return len(types)
	}
	return maxAllowed
}

func main() {
	fmt.Println(DistributeCandies([]int{1, 1, 2, 2, 3, 3}))
	fmt.Println(DistributeCandies([]int{1, 1, 2, 3}))
	fmt.Println(DistributeCandies([]int{6, 6, 6, 6}))
}
```

## 0577 — Employee Bonus

```go
package main

// LeetCode #577: Employee Bonus
// https://leetcode.com/problems/employee-bonus/
// Difficulty: Easy

import "fmt"

func EmployeeBonus() string {
	return "SELECT e.name, b.bonus FROM Employee e LEFT JOIN Bonus b ON e.empId = b.empId WHERE b.bonus < 1000 OR b.bonus IS NULL"
}

func main() {
	fmt.Println(EmployeeBonus())
}
```

## 0584 — Find Customer Referee

```go
package main

// LeetCode #584: Find Customer Referee
// https://leetcode.com/problems/find-customer-referee/
// Difficulty: Easy

import "fmt"

func FindCustomerReferee() string {
	return "SELECT name FROM Customer WHERE referee_id != 2 OR referee_id IS NULL"
}

func main() {
	fmt.Println(FindCustomerReferee())
}
```

## 0586 — Customer Placing The Largest Number Of Orders

```go
package main

// LeetCode #586: Customer Placing the Largest Number of Orders
// https://leetcode.com/problems/customer-placing-the-largest-number-of-orders/
// Difficulty: Easy

import "fmt"

func CustomerPlacingTheLargestNumberOfOrders() string {
	return "SELECT customer_number FROM Orders GROUP BY customer_number ORDER BY COUNT(*) DESC LIMIT 1"
}

func main() {
	fmt.Println(CustomerPlacingTheLargestNumberOfOrders())
}
```

## 0589 — N Ary Tree Preorder Traversal

```go
package main

// LeetCode #589: N-ary Tree Preorder Traversal
// https://leetcode.com/problems/n-ary-tree-preorder-traversal/
// Difficulty: Easy

import "fmt"

// Node represents an N-ary tree node.
type Node struct {
	Val      int
	Children []*Node
}

// Time: O(n), Space: O(n)
func NAryTreePreorderTraversal(root *Node) []int {
	var result []int
	var dfs func(node *Node)
	dfs = func(node *Node) {
		if node == nil {
			return
		}
		result = append(result, node.Val)
		for _, child := range node.Children {
			dfs(child)
		}
	}
	dfs(root)
	return result
}

func main() {
	// Test: [1,null,3,2,4,null,5,6]
	root1 := &Node{
		Val: 1,
		Children: []*Node{
			{Val: 3, Children: []*Node{
				{Val: 5},
				{Val: 6},
			}},
			{Val: 2},
			{Val: 4},
		},
	}
	fmt.Println(NAryTreePreorderTraversal(root1))

	// Test: single node
	root2 := &Node{Val: 1}
	fmt.Println(NAryTreePreorderTraversal(root2))
}
```

## 0590 — N Ary Tree Postorder Traversal

```go
package main

// LeetCode #590: N-ary Tree Postorder Traversal
// https://leetcode.com/problems/n-ary-tree-postorder-traversal/
// Difficulty: Easy

import "fmt"

// Node represents an N-ary tree node.
type Node struct {
	Val      int
	Children []*Node
}

// Time: O(n), Space: O(n)
func NAryTreePostorderTraversal(root *Node) []int {
	var result []int
	var dfs func(node *Node)
	dfs = func(node *Node) {
		if node == nil {
			return
		}
		for _, child := range node.Children {
			dfs(child)
		}
		result = append(result, node.Val)
	}
	dfs(root)
	return result
}

func main() {
	// Test: [1,null,3,2,4,null,5,6]
	root1 := &Node{
		Val: 1,
		Children: []*Node{
			{Val: 3, Children: []*Node{
				{Val: 5},
				{Val: 6},
			}},
			{Val: 2},
			{Val: 4},
		},
	}
	fmt.Println(NAryTreePostorderTraversal(root1))

	// Test: single node
	root2 := &Node{Val: 1}
	fmt.Println(NAryTreePostorderTraversal(root2))
}
```

## 0594 — Longest Harmonious Subsequence

```go
package main

// LeetCode #594: Longest Harmonious Subsequence
// https://leetcode.com/problems/longest-harmonious-subsequence/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(n)
func LongestHarmoniousSubsequence(nums []int) int {
	count := make(map[int]int)
	for _, v := range nums {
		count[v]++
	}
	maxLen := 0
	for v, c := range count {
		if c2, ok := count[v+1]; ok {
			if c+c2 > maxLen {
				maxLen = c + c2
			}
		}
	}
	return maxLen
}

func main() {
	fmt.Println(LongestHarmoniousSubsequence([]int{1, 3, 2, 2, 5, 2, 3, 7}))
	fmt.Println(LongestHarmoniousSubsequence([]int{1, 2, 3, 4}))
	fmt.Println(LongestHarmoniousSubsequence([]int{1, 1, 1, 1}))
}
```

## 0595 — Big Countries

```go
package main

// LeetCode #595: Big Countries
// https://leetcode.com/problems/big-countries/
// Difficulty: Easy

import "fmt"

func BigCountries() string {
	return "SELECT name, population, area FROM World WHERE area >= 3000000 OR population >= 25000000"
}

func main() {
	fmt.Println(BigCountries())
}
```

## 0596 — Classes With At Least 5 Students

```go
package main

// LeetCode #596: Classes With at Least 5 Students
// https://leetcode.com/problems/classes-with-at-least-5-students/
// Difficulty: Easy

import "fmt"

func ClassesWithAtLeastFiveStudents() string {
	return "SELECT class FROM Courses GROUP BY class HAVING COUNT(student) >= 5"
}

func main() {
	fmt.Println(ClassesWithAtLeastFiveStudents())
}
```

## 0597 — Friend Requests I Overall Acceptance Rate

```go
package main

// LeetCode #597: Friend Requests I: Overall Acceptance Rate
// https://leetcode.com/problems/friend-requests-i-overall-acceptance-rate/
// Difficulty: Easy [Paid]

import "fmt"

func FriendRequestsIOverallAcceptanceRate() string {
	return "SELECT ROUND(IFNULL((SELECT COUNT(DISTINCT requester_id, accepter_id) FROM RequestAccepted) / (SELECT COUNT(DISTINCT sender_id, send_to_id) FROM FriendRequest), 0), 2) AS accept_rate"
}

func main() {
	fmt.Println(FriendRequestsIOverallAcceptanceRate())
}
```

## 0598 — Range Addition Ii

```go
package main

// LeetCode #598: Range Addition II
// https://leetcode.com/problems/range-addition-ii/
// Difficulty: Easy

import "fmt"

// Time: O(k), Space: O(1)
func RangeAdditionIi(m, n int, ops [][]int) int {
	minA, minB := m, n
	for _, op := range ops {
		if op[0] < minA {
			minA = op[0]
		}
		if op[1] < minB {
			minB = op[1]
		}
	}
	return minA * minB
}

func main() {
	fmt.Println(RangeAdditionIi(3, 3, [][]int{{2, 2}, {3, 3}}))
	fmt.Println(RangeAdditionIi(3, 3, [][]int{}))
}
```

## 0599 — Minimum Index Sum Of Two Lists

```go
package main

// LeetCode #599: Minimum Index Sum of Two Lists
// https://leetcode.com/problems/minimum-index-sum-of-two-lists/
// Difficulty: Easy

import (
	"fmt"
	"math"
)

// Time: O(n+m), Space: O(n)
func MinimumIndexSumOfTwoLists(list1, list2 []string) []string {
	index := make(map[string]int)
	for i, s := range list1 {
		index[s] = i
	}
	minSum := math.MaxInt32
	var result []string
	for j, s := range list2 {
		if i, ok := index[s]; ok {
			sum := i + j
			if sum < minSum {
				minSum = sum
				result = []string{s}
			} else if sum == minSum {
				result = append(result, s)
			}
		}
	}
	return result
}

func main() {
	fmt.Println(MinimumIndexSumOfTwoLists(
		[]string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
		[]string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"},
	))
	fmt.Println(MinimumIndexSumOfTwoLists(
		[]string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
		[]string{"KFC", "Shogun", "Burger King"},
	))
}
```

## 0603 — Consecutive Available Seats

```go
package main

// LeetCode #603: Consecutive Available Seats
// https://leetcode.com/problems/consecutive-available-seats/
// Difficulty: Easy [Paid]

import "fmt"

func ConsecutiveAvailableSeats() string {
	return "SELECT DISTINCT c1.seat_id FROM Cinema c1 JOIN Cinema c2 ON ABS(c1.seat_id - c2.seat_id) = 1 AND c1.free = 1 AND c2.free = 1 ORDER BY c1.seat_id"
}

func main() {
	fmt.Println(ConsecutiveAvailableSeats())
}
```

## 0604 — Design Compressed String Iterator

```go
package main

// LeetCode #604: Design Compressed String Iterator
// https://leetcode.com/problems/design-compressed-string-iterator/
// Difficulty: Easy [Paid]

import (
	"fmt"
	"strconv"
)

// StringIterator iterates over a compressed string.
type StringIterator struct {
	chars []byte
	counts []int
	index  int
}

// Constructor creates a new StringIterator from compressed string.
func Constructor(compressedString string) StringIterator {
	var chars []byte
	var counts []int
	i := 0
	for i < len(compressedString) {
		c := compressedString[i]
		i++
		numStart := i
		for i < len(compressedString) && compressedString[i] >= '0' && compressedString[i] <= '9' {
			i++
		}
		count, _ := strconv.Atoi(compressedString[numStart:i])
		chars = append(chars, c)
		counts = append(counts, count)
	}
	return StringIterator{chars: chars, counts: counts, index: 0}
}

// Next returns the next character or ' ' if exhausted.
// Time: O(1), Space: O(1)
func (it *StringIterator) Next() byte {
	if !it.HasNext() {
		return ' '
	}
	c := it.chars[it.index]
	it.counts[it.index]--
	if it.counts[it.index] == 0 {
		it.index++
	}
	return c
}

// HasNext returns true if there are more characters.
// Time: O(1), Space: O(1)
func (it *StringIterator) HasNext() bool {
	return it.index < len(it.chars)
}

func main() {
	obj := Constructor("L1e2t1C1o1d1e1")
	fmt.Printf("%c ", obj.Next())
	fmt.Printf("%c ", obj.Next())
	fmt.Printf("%c ", obj.Next())
	fmt.Printf("%c ", obj.Next())
	fmt.Printf("%c ", obj.Next())
	fmt.Printf("%c ", obj.Next())
	fmt.Printf("%c ", obj.Next())
	fmt.Printf("%c ", obj.Next())
	fmt.Printf("%c\n", obj.Next())
	fmt.Println(obj.HasNext())
}
```

## 0605 — Can Place Flowers

```go
package main

// LeetCode #605: Can Place Flowers
// https://leetcode.com/problems/can-place-flowers/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func CanPlaceFlowers(flowerbed []int, n int) bool {
	count := 0
	for i := 0; i < len(flowerbed) && count < n; i++ {
		if flowerbed[i] == 0 &&
			(i == 0 || flowerbed[i-1] == 0) &&
			(i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
			flowerbed[i] = 1
			count++
		}
	}
	return count >= n
}

func main() {
	fmt.Println(CanPlaceFlowers([]int{1, 0, 0, 0, 1}, 1))
	fmt.Println(CanPlaceFlowers([]int{1, 0, 0, 0, 1}, 2))
	fmt.Println(CanPlaceFlowers([]int{0, 0, 1, 0, 0}, 1))
}
```

## 0607 — Sales Person

```go
package main

// LeetCode #607: Sales Person
// https://leetcode.com/problems/sales-person/
// Difficulty: Easy

import "fmt"

func SalesPerson() string {
	return "SELECT s.name FROM SalesPerson s WHERE s.sales_id NOT IN (SELECT o.sales_id FROM Orders o JOIN Company c ON o.com_id = c.com_id WHERE c.name = 'RED')"
}

func main() {
	fmt.Println(SalesPerson())
}
```

## 0610 — Triangle Judgement

```go
package main

// LeetCode #610: Triangle Judgement
// https://leetcode.com/problems/triangle-judgement/
// Difficulty: Easy

import "fmt"

func TriangleJudgement() string {
	return "SELECT x, y, z, CASE WHEN x + y > z AND x + z > y AND y + z > x THEN 'Yes' ELSE 'No' END AS triangle FROM Triangle"
}

func main() {
	fmt.Println(TriangleJudgement())
}
```

## 0613 — Shortest Distance In A Line

```go
package main

// LeetCode #613: Shortest Distance in a Line
// https://leetcode.com/problems/shortest-distance-in-a-line/
// Difficulty: Easy [Paid]

import "fmt"

func ShortestDistanceInALine() string {
	return "SELECT MIN(ABS(p1.x - p2.x)) AS shortest FROM Point p1 JOIN Point p2 ON p1.x != p2.x"
}

func main() {
	fmt.Println(ShortestDistanceInALine())
}
```

## 0617 — Merge Two Binary Trees

```go
package main

// LeetCode #617: Merge Two Binary Trees
// https://leetcode.com/problems/merge-two-binary-trees/
// Difficulty: Easy

import "fmt"

// TreeNode represents a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n+m), Space: O(h)
func MergeTwoBinaryTrees(root1, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	return &TreeNode{
		Val:   root1.Val + root2.Val,
		Left:  MergeTwoBinaryTrees(root1.Left, root2.Left),
		Right: MergeTwoBinaryTrees(root1.Right, root2.Right),
	}
}

func printTree(root *TreeNode) {
	if root == nil {
		fmt.Print("null ")
		return
	}
	fmt.Printf("%d ", root.Val)
	printTree(root.Left)
	printTree(root.Right)
}

func main() {
	// Test: root1=[1,3,2,5], root2=[2,1,3,null,4,null,7]
	root1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 5},
		},
		Right: &TreeNode{Val: 2},
	}
	root2 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:   1,
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:   3,
			Right: &TreeNode{Val: 7},
		},
	}
	printTree(MergeTwoBinaryTrees(root1, root2))
	fmt.Println()

	// Test: root1=[1], root2=[1,2]
	r1 := &TreeNode{Val: 1}
	r2 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	printTree(MergeTwoBinaryTrees(r1, r2))
	fmt.Println()
}
```

## 0619 — Biggest Single Number

```go
package main

// LeetCode #619: Biggest Single Number
// https://leetcode.com/problems/biggest-single-number/
// Difficulty: Easy

import "fmt"

func BiggestSingleNumber() string {
	return "SELECT MAX(num) AS num FROM (SELECT num FROM MyNumbers GROUP BY num HAVING COUNT(num) = 1) AS single_numbers"
}

func main() {
	fmt.Println(BiggestSingleNumber())
}
```

## 0620 — Not Boring Movies

```go
package main

// LeetCode #620: Not Boring Movies
// https://leetcode.com/problems/not-boring-movies/
// Difficulty: Easy

import "fmt"

func NotBoringMovies() string {
	return "SELECT * FROM Cinema WHERE id % 2 = 1 AND description != 'boring' ORDER BY rating DESC"
}

func main() {
	fmt.Println(NotBoringMovies())
}
```

## 0627 — Swap Sex Of Employees

```go
package main

// LeetCode #627: Swap Sex of Employees
// https://leetcode.com/problems/swap-sex-of-employees/
// Difficulty: Easy

import "fmt"

func SwapSexOfEmployees() string {
	return "UPDATE Salary SET sex = CASE WHEN sex = 'm' THEN 'f' ELSE 'm' END"
}

func main() {
	fmt.Println(SwapSexOfEmployees())
}
```

## 0628 — Maximum Product Of Three Numbers

```go
package main

// LeetCode #628: Maximum Product of Three Numbers
// https://leetcode.com/problems/maximum-product-of-three-numbers/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

// Time: O(n log n), Space: O(1)
func MaximumProductOfThreeNumbers(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	// The max product is either the three largest numbers
	// or two smallest (negative) numbers times the largest.
	p1 := nums[n-1] * nums[n-2] * nums[n-3]
	p2 := nums[0] * nums[1] * nums[n-1]
	if p1 > p2 {
		return p1
	}
	return p2
}

func main() {
	fmt.Println(MaximumProductOfThreeNumbers([]int{1, 2, 3}))
	fmt.Println(MaximumProductOfThreeNumbers([]int{1, 2, 3, 4}))
	fmt.Println(MaximumProductOfThreeNumbers([]int{-1, -2, -3}))
}
```

## 0637 — Average Of Levels In Binary Tree

```go
package main

// LeetCode #637: Average of Levels in Binary Tree
// https://leetcode.com/problems/average-of-levels-in-binary-tree/
// Difficulty: Easy

import "fmt"

// TreeNode represents a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n), Space: O(n)
func AverageOfLevelsInBinaryTree(root *TreeNode) []float64 {
	var result []float64
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		sum := 0
		for i := 0; i < levelSize; i++ {
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
		result = append(result, float64(sum)/float64(levelSize))
	}
	return result
}

func main() {
	// Test: [3,9,20,null,null,15,7]
	root1 := &TreeNode{
		Val: 3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}
	fmt.Println(AverageOfLevelsInBinaryTree(root1))

	// Test: [3,9,20,15,7]
	root2 := &TreeNode{
		Val:   3,
		Left:  &TreeNode{Val: 9, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
		Right: &TreeNode{Val: 20},
	}
	fmt.Println(AverageOfLevelsInBinaryTree(root2))
}
```

## 0643 — Maximum Average Subarray I

```go
package main

// LeetCode #643: Maximum Average Subarray I
// https://leetcode.com/problems/maximum-average-subarray-i/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4)) // 12.75
	fmt.Println(findMaxAverage([]int{5}, 1))                     // 5.0
	fmt.Println(findMaxAverage([]int{-1}, 1))                    // -1.0
}

// findMaxAverage finds a contiguous subarray of length k that has the maximum average value.
// Time: O(n). Space: O(1).
func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	maxSum := sum
	for i := k; i < len(nums); i++ {
		sum += nums[i] - nums[i-k]
		if sum > maxSum {
			maxSum = sum
		}
	}
	return float64(maxSum) / float64(k)
}
```

## 0645 — Set Mismatch

```go
package main

// LeetCode #645: Set Mismatch
// https://leetcode.com/problems/set-mismatch/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(findErrorNums([]int{1, 2, 2, 4})) // [2, 3]
	fmt.Println(findErrorNums([]int{1, 1}))        // [1, 2]
	fmt.Println(findErrorNums([]int{2, 2}))        // [2, 1]
}

// findErrorNums finds the duplicated and missing number in the set.
// Time: O(n). Space: O(1).
func findErrorNums(nums []int) []int {
	n := len(nums)
	sum := 0
	sumSq := 0
	expectedSum := n * (n + 1) / 2
	expectedSumSq := n * (n + 1) * (2*n + 1) / 6

	for _, v := range nums {
		sum += v
		sumSq += v * v
	}

	// diff = duplicate - missing
	diff := sum - expectedSum
	// sqDiff = duplicate^2 - missing^2
	sqDiff := sumSq - expectedSumSq
	// duplicate + missing = sqDiff / diff
	plus := sqDiff / diff

	dup := (diff + plus) / 2
	miss := (plus - diff) / 2
	return []int{dup, miss}
}
```

## 0653 — Two Sum Iv Input Is A Bst

```go
package main

// LeetCode #653: Two Sum IV - Input is a BST
// https://leetcode.com/problems/two-sum-iv-input-is-a-bst/
// Difficulty: Easy

import "fmt"

// TreeNode defines a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Example: [5,3,6,2,4,null,7], k=9 => true
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 2},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:   6,
			Right: &TreeNode{Val: 7},
		},
	}
	fmt.Println(findTarget(root, 9))  // true
	fmt.Println(findTarget(root, 28)) // false
}

// findTarget returns true if there exist two elements in the BST that sum to k.
// Time: O(n). Space: O(n).
func findTarget(root *TreeNode, k int) bool {
	seen := make(map[int]bool)
	return find(root, k, seen)
}

func find(node *TreeNode, k int, seen map[int]bool) bool {
	if node == nil {
		return false
	}
	if seen[k-node.Val] {
		return true
	}
	seen[node.Val] = true
	return find(node.Left, k, seen) || find(node.Right, k, seen)
}
```

## 0657 — Robot Return To Origin

```go
package main

// LeetCode #657: Robot Return to Origin
// https://leetcode.com/problems/robot-return-to-origin/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(judgeCircle("UD"))   // true
	fmt.Println(judgeCircle("LL"))   // false
	fmt.Println(judgeCircle(""))     // true
}

// judgeCircle returns true if the robot returns to origin after executing all moves.
// Time: O(n). Space: O(1).
func judgeCircle(moves string) bool {
	x, y := 0, 0
	for _, c := range moves {
		switch c {
		case 'U':
			y++
		case 'D':
			y--
		case 'L':
			x--
		case 'R':
			x++
		}
	}
	return x == 0 && y == 0
}
```

## 0661 — Image Smoother

```go
package main

// LeetCode #661: Image Smoother
// https://leetcode.com/problems/image-smoother/
// Difficulty: Easy

import "fmt"

func main() {
	img := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	fmt.Println(imageSmoother(img))
	// [[0,0,0],[0,0,0],[0,0,0]]

	img2 := [][]int{{100, 200, 100}, {200, 50, 200}, {100, 200, 100}}
	fmt.Println(imageSmoother(img2))
}

// imageSmoother applies a 3x3 smoother to each cell of the image.
// Time: O(m*n). Space: O(m*n).
func imageSmoother(img [][]int) [][]int {
	m, n := len(img), len(img[0])
	result := make([][]int, m)
	for i := range result {
		result[i] = make([]int, n)
	}

	directions := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 0}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sum, count := 0, 0
			for _, d := range directions {
				ni, nj := i+d[0], j+d[1]
				if ni >= 0 && ni < m && nj >= 0 && nj < n {
					sum += img[ni][nj]
					count++
				}
			}
			result[i][j] = sum / count
		}
	}
	return result
}
```

## 0671 — Second Minimum Node In A Binary Tree

```go
package main

// LeetCode #671: Second Minimum Node In a Binary Tree
// https://leetcode.com/problems/second-minimum-node-in-a-binary-tree/
// Difficulty: Easy

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
	// Example: [2,2,5,null,null,5,7] => 5
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{Val: 2},
		Right: &TreeNode{
			Val:   5,
			Left:  &TreeNode{Val: 5},
			Right: &TreeNode{Val: 7},
		},
	}
	fmt.Println(findSecondMinimumValue(root)) // 5

	root2 := &TreeNode{
		Val:  2,
		Left: &TreeNode{Val: 2},
		Right: &TreeNode{Val: 2},
	}
	fmt.Println(findSecondMinimumValue(root2)) // -1
}

// findSecondMinimumValue finds the second minimum value in a special binary tree
// where each node's value is the minimum of its children.
// Time: O(n). Space: O(n).
func findSecondMinimumValue(root *TreeNode) int {
	result := math.MaxInt64
	dfs(root, root.Val, &result)
	if result == math.MaxInt64 {
		return -1
	}
	return result
}

func dfs(node *TreeNode, rootVal int, second *int) {
	if node == nil {
		return
	}
	if node.Val > rootVal && node.Val < *second {
		*second = node.Val
	}
	dfs(node.Left, rootVal, second)
	dfs(node.Right, rootVal, second)
}
```

## 0674 — Longest Continuous Increasing Subsequence

```go
package main

// LeetCode #674: Longest Continuous Increasing Subsequence
// https://leetcode.com/problems/longest-continuous-increasing-subsequence/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(findLengthOfLCIS([]int{1, 3, 5, 4, 7}))    // 3
	fmt.Println(findLengthOfLCIS([]int{2, 2, 2, 2, 2}))    // 1
	fmt.Println(findLengthOfLCIS([]int{1, 3, 5, 7}))       // 4
}

// findLengthOfLCIS finds the length of the longest continuous increasing subsequence.
// Time: O(n). Space: O(1).
func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxLen, curr := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			curr++
			if curr > maxLen {
				maxLen = curr
			}
		} else {
			curr = 1
		}
	}
	return maxLen
}
```

## 0680 — Valid Palindrome Ii

```go
package main

// LeetCode #680: Valid Palindrome II
// https://leetcode.com/problems/valid-palindrome-ii/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(validPalindrome("aba"))    // true
	fmt.Println(validPalindrome("abca"))   // true
	fmt.Println(validPalindrome("abc"))    // false
	fmt.Println(validPalindrome("deeee"))  // true
}

// validPalindrome checks if the string can be a palindrome after deleting at most one character.
// Time: O(n). Space: O(1).
func validPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return isPal(s, l+1, r) || isPal(s, l, r-1)
		}
		l++
		r--
	}
	return true
}

func isPal(s string, l, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
```

## 0682 — Baseball Game

```go
package main

// LeetCode #682: Baseball Game
// https://leetcode.com/problems/baseball-game/
// Difficulty: Easy

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(calPoints([]string{"5", "2", "C", "D", "+"})) // 30
	fmt.Println(calPoints([]string{"5", "-2", "4", "C", "D", "9", "+", "+"})) // 27
	fmt.Println(calPoints([]string{"1", "C"})) // 0
}

// calPoints calculates the total score for a baseball game based on operations.
// Time: O(n). Space: O(n).
func calPoints(operations []string) int {
	stack := make([]int, 0, len(operations))
	for _, op := range operations {
		switch op {
		case "C":
			stack = stack[:len(stack)-1]
		case "D":
			stack = append(stack, 2*stack[len(stack)-1])
		case "+":
			stack = append(stack, stack[len(stack)-1]+stack[len(stack)-2])
		default:
			n, _ := strconv.Atoi(op)
			stack = append(stack, n)
		}
	}
	sum := 0
	for _, v := range stack {
		sum += v
	}
	return sum
}
```

## 0693 — Binary Number With Alternating Bits

```go
package main

// LeetCode #693: Binary Number with Alternating Bits
// https://leetcode.com/problems/binary-number-with-alternating-bits/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(hasAlternatingBits(5))  // true (101)
	fmt.Println(hasAlternatingBits(7))  // false (111)
	fmt.Println(hasAlternatingBits(11)) // false (1011)
	fmt.Println(hasAlternatingBits(10)) // true (1010)
}

// hasAlternatingBits checks if the binary representation of n has alternating bits.
// Time: O(log n). Space: O(1).
func hasAlternatingBits(n int) bool {
	// XOR with n>>1 gives all 1s if alternating
	x := n ^ (n >> 1)
	// Check if x is all 1s (i.e., x & (x+1) == 0)
	return x&(x+1) == 0
}
```

## 0696 — Count Binary Substrings

```go
package main

// LeetCode #696: Count Binary Substrings
// https://leetcode.com/problems/count-binary-substrings/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(countBinarySubstrings("00110011")) // 6
	fmt.Println(countBinarySubstrings("10101"))    // 4
	fmt.Println(countBinarySubstrings("00110"))    // 3
}

// countBinarySubstrings counts substrings that have equal numbers of 0s and 1s.
// Time: O(n). Space: O(1).
func countBinarySubstrings(s string) int {
	prev, curr, result := 0, 1, 0
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			curr++
		} else {
			prev = curr
			curr = 1
		}
		if prev >= curr {
			result++
		}
	}
	return result
}
```

## 0697 — Degree Of An Array

```go
package main

// LeetCode #697: Degree of an Array
// https://leetcode.com/problems/degree-of-an-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(findShortestSubArray([]int{1, 2, 2, 3, 1}))       // 2
	fmt.Println(findShortestSubArray([]int{1, 2, 2, 3, 1, 4, 2})) // 6
	fmt.Println(findShortestSubArray([]int{1}))                    // 1
}

// findShortestSubArray finds the smallest subarray length with the same degree as the array.
// Time: O(n). Space: O(n).
func findShortestSubArray(nums []int) int {
	first := make(map[int]int)
	count := make(map[int]int)
	maxCount := 0
	minLen := len(nums)

	for i, v := range nums {
		if _, ok := first[v]; !ok {
			first[v] = i
		}
		count[v]++
		if count[v] > maxCount {
			maxCount = count[v]
		}
	}

	for v, c := range count {
		if c == maxCount {
			// find last occurrence
			last := 0
			for i := len(nums) - 1; i >= 0; i-- {
				if nums[i] == v {
					last = i
					break
				}
			}
			length := last - first[v] + 1
			if length < minLen {
				minLen = length
			}
		}
	}
	return minLen
}
```

## 0700 — Search In A Binary Search Tree

```go
package main

// LeetCode #700: Search in a Binary Search Tree
// https://leetcode.com/problems/search-in-a-binary-search-tree/
// Difficulty: Easy

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Example: [4,2,7,1,3], val=2
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{Val: 7},
	}
	fmt.Println(searchBST(root, 2))  // node with value 2
	fmt.Println(searchBST(root, 5))  // nil
}

// searchBST searches for a node with the given value in a BST.
// Time: O(log n) average, O(n) worst. Space: O(1).
func searchBST(root *TreeNode, val int) *TreeNode {
	curr := root
	for curr != nil {
		if val == curr.Val {
			return curr
		} else if val < curr.Val {
			curr = curr.Left
		} else {
			curr = curr.Right
		}
	}
	return nil
}
```

## 0703 — Kth Largest Element In A Stream

```go
package main

// LeetCode #703: Kth Largest Element in a Stream
// https://leetcode.com/problems/kth-largest-element-in-a-stream/
// Difficulty: Easy

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// KthLargest maintains the kth largest element in a stream.
type KthLargest struct {
	k    int
	heap *MinHeap
}

// Constructor creates a KthLargest instance.
func Constructor(k int, nums []int) KthLargest {
	h := &MinHeap{}
	heap.Init(h)
	kl := KthLargest{k: k, heap: h}
	for _, v := range nums {
		kl.Add(v)
	}
	return kl
}

// Add adds a new value and returns the kth largest.
// Time: O(log k). Space: O(k).
func (kl *KthLargest) Add(val int) int {
	heap.Push(kl.heap, val)
	if kl.heap.Len() > kl.k {
		heap.Pop(kl.heap)
	}
	return (*kl.heap)[0]
}

func main() {
	kl := Constructor(3, []int{4, 5, 8, 2})
	fmt.Println(kl.Add(3))  // 4
	fmt.Println(kl.Add(5))  // 5
	fmt.Println(kl.Add(10)) // 5
	fmt.Println(kl.Add(9))  // 8
	fmt.Println(kl.Add(4))  // 8
}
```

## 0704 — Binary Search

```go
package main

// LeetCode #704: Binary Search
// https://leetcode.com/problems/binary-search/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9))    // 4
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 2))    // -1
	fmt.Println(search([]int{5}, 5))                      // 0
}

// search performs binary search on a sorted array.
// Time: O(log n). Space: O(1).
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}
```

## 0705 — Design Hashset

```go
package main

// LeetCode #705: Design HashSet
// https://leetcode.com/problems/design-hashset/
// Difficulty: Easy

import "fmt"

const hashSetSize = 1000

type hashNode struct {
	key  int
	next *hashNode
}

// MyHashSet implements a hash set using separate chaining.
type MyHashSet struct {
	buckets []*hashNode
}

// Constructor creates a MyHashSet.
func Constructor() MyHashSet {
	return MyHashSet{buckets: make([]*hashNode, hashSetSize)}
}

// Add inserts a key into the set.
// Time: O(1) average. Space: O(n).
func (s *MyHashSet) Add(key int) {
	if s.Contains(key) {
		return
	}
	idx := key % hashSetSize
	s.buckets[idx] = &hashNode{key: key, next: s.buckets[idx]}
}

// Remove deletes a key from the set.
func (s *MyHashSet) Remove(key int) {
	idx := key % hashSetSize
	curr := s.buckets[idx]
	var prev *hashNode
	for curr != nil {
		if curr.key == key {
			if prev == nil {
				s.buckets[idx] = curr.next
			} else {
				prev.next = curr.next
			}
			return
		}
		prev = curr
		curr = curr.next
	}
}

// Contains checks if a key exists in the set.
func (s *MyHashSet) Contains(key int) bool {
	idx := key % hashSetSize
	curr := s.buckets[idx]
	for curr != nil {
		if curr.key == key {
			return true
		}
		curr = curr.next
	}
	return false
}

func main() {
	hs := Constructor()
	hs.Add(1)
	hs.Add(2)
	fmt.Println(hs.Contains(1)) // true
	fmt.Println(hs.Contains(3)) // false
	hs.Add(2)
	fmt.Println(hs.Contains(2)) // true
	hs.Remove(2)
	fmt.Println(hs.Contains(2)) // false
}
```

## 0706 — Design Hashmap

```go
package main

// LeetCode #706: Design HashMap
// https://leetcode.com/problems/design-hashmap/
// Difficulty: Easy

import "fmt"

const hashMapSize = 1000

type kvNode struct {
	key   int
	value int
	next  *kvNode
}

// MyHashMap implements a hash map using separate chaining.
type MyHashMap struct {
	buckets []*kvNode
}

// Constructor creates a MyHashMap.
func Constructor() MyHashMap {
	return MyHashMap{buckets: make([]*kvNode, hashMapSize)}
}

// Put inserts a key-value pair into the map.
// Time: O(1) average. Space: O(n).
func (m *MyHashMap) Put(key int, value int) {
	idx := key % hashMapSize
	curr := m.buckets[idx]
	for curr != nil {
		if curr.key == key {
			curr.value = value
			return
		}
		curr = curr.next
	}
	m.buckets[idx] = &kvNode{key: key, value: value, next: m.buckets[idx]}
}

// Get returns the value for a key, or -1 if not found.
func (m *MyHashMap) Get(key int) int {
	idx := key % hashMapSize
	curr := m.buckets[idx]
	for curr != nil {
		if curr.key == key {
			return curr.value
		}
		curr = curr.next
	}
	return -1
}

// Remove deletes a key from the map.
func (m *MyHashMap) Remove(key int) {
	idx := key % hashMapSize
	curr := m.buckets[idx]
	var prev *kvNode
	for curr != nil {
		if curr.key == key {
			if prev == nil {
				m.buckets[idx] = curr.next
			} else {
				prev.next = curr.next
			}
			return
		}
		prev = curr
		curr = curr.next
	}
}

func main() {
	hm := Constructor()
	hm.Put(1, 1)
	hm.Put(2, 2)
	fmt.Println(hm.Get(1)) // 1
	fmt.Println(hm.Get(3)) // -1
	hm.Put(2, 1)
	fmt.Println(hm.Get(2)) // 1
	hm.Remove(2)
	fmt.Println(hm.Get(2)) // -1
}
```

## 0709 — To Lower Case

```go
package main

// LeetCode #709: To Lower Case
// https://leetcode.com/problems/to-lower-case/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(toLowerCase("Hello"))  // "hello"
	fmt.Println(toLowerCase("LOVELY")) // "lovely"
	fmt.Println(toLowerCase("here"))   // "here"
}

// toLowerCase converts a string to lowercase.
// Time: O(n). Space: O(n).
func toLowerCase(s string) string {
	b := []byte(s)
	for i, c := range b {
		if c >= 'A' && c <= 'Z' {
			b[i] = c + 32
		}
	}
	return string(b)
}
```

## 0717 — 1 Bit And 2 Bit Characters

```go
package main

// LeetCode #717: 1-bit and 2-bit Characters
// https://leetcode.com/problems/1-bit-and-2-bit-characters/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(isOneBitCharacter([]int{1, 0, 0}))          // true
	fmt.Println(isOneBitCharacter([]int{1, 1, 1, 0}))      // false
	fmt.Println(isOneBitCharacter([]int{0}))                // true
}

// isOneBitCharacter checks if the last character must be a one-bit character.
// Time: O(n). Space: O(1).
func isOneBitCharacter(bits []int) bool {
	i := 0
	for i < len(bits)-1 {
		if bits[i] == 1 {
			i += 2
		} else {
			i++
		}
	}
	return i == len(bits)-1
}
```

## 0724 — Find Pivot Index

```go
package main

// LeetCode #724: Find Pivot Index
// https://leetcode.com/problems/find-pivot-index/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))   // 3
	fmt.Println(pivotIndex([]int{1, 2, 3}))             // -1
	fmt.Println(pivotIndex([]int{2, 1, -1}))            // 0
}

// pivotIndex finds the index where sum of left elements equals sum of right elements.
// Time: O(n). Space: O(1).
func pivotIndex(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	leftSum := 0
	for i, v := range nums {
		if leftSum == total-leftSum-v {
			return i
		}
		leftSum += v
	}
	return -1
}
```

## 0728 — Self Dividing Numbers

```go
package main

// LeetCode #728: Self Dividing Numbers
// https://leetcode.com/problems/self-dividing-numbers/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(selfDividingNumbers(1, 22)) // [1,2,3,4,5,6,7,8,9,11,12,15,22]
	fmt.Println(selfDividingNumbers(47, 85)) // [48,55,66,77]
}

// selfDividingNumbers returns all self-dividing numbers in range [left, right].
// Time: O(n * d) where n = count, d = digits per number. Space: O(1) excluding output.
func selfDividingNumbers(left int, right int) []int {
	result := make([]int, 0)
	for n := left; n <= right; n++ {
		if isSelfDividing(n) {
			result = append(result, n)
		}
	}
	return result
}

func isSelfDividing(n int) bool {
	original := n
	for n > 0 {
		digit := n % 10
		if digit == 0 || original%digit != 0 {
			return false
		}
		n /= 10
	}
	return true
}
```

## 0733 — Flood Fill

```go
package main

// LeetCode #733: Flood Fill
// https://leetcode.com/problems/flood-fill/
// Difficulty: Easy

import "fmt"

func main() {
	image := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	fmt.Println(floodFill(image, 1, 1, 2))
	// [[2,2,2],[2,2,0],[2,0,1]]

	image2 := [][]int{{0, 0, 0}, {0, 0, 0}}
	fmt.Println(floodFill(image2, 0, 0, 0))
	// [[0,0,0],[0,0,0]]
}

// floodFill performs a flood fill on the image starting from (sr, sc).
// Time: O(m*n). Space: O(m*n) for recursion stack.
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	originalColor := image[sr][sc]
	if originalColor == color {
		return image
	}
	dfs(image, sr, sc, originalColor, color)
	return image
}

func dfs(image [][]int, r, c, originalColor, newColor int) {
	if r < 0 || r >= len(image) || c < 0 || c >= len(image[0]) {
		return
	}
	if image[r][c] != originalColor {
		return
	}
	image[r][c] = newColor
	dfs(image, r+1, c, originalColor, newColor)
	dfs(image, r-1, c, originalColor, newColor)
	dfs(image, r, c+1, originalColor, newColor)
	dfs(image, r, c-1, originalColor, newColor)
}
```

## 0734 — Sentence Similarity

```go
package main

// LeetCode #734: Sentence Similarity
// https://leetcode.com/problems/sentence-similarity/
// Difficulty: Easy [Paid]
// Note: This is a premium problem. Implementation based on public description.

import "fmt"

func main() {
	s1 := []string{"great", "acting", "skills"}
	s2 := []string{"fine", "drama", "talent"}
	pairs := [][]string{{"great", "fine"}, {"acting", "drama"}, {"skills", "talent"}}
	fmt.Println(areSentencesSimilar(s1, s2, pairs)) // true

	s1 = []string{"great"}
	s2 = []string{"great"}
	fmt.Println(areSentencesSimilar(s1, s2, [][]string{})) // true

	s1 = []string{"great"}
	s2 = []string{"doubleplus", "good"}
	fmt.Println(areSentencesSimilar(s1, s2, pairs)) // false
}

// areSentencesSimilar checks if two sentences are similar according to given similarity pairs.
// Time: O(n + p) where n = len(sentence), p = len(pairs). Space: O(p).
func areSentencesSimilar(sentence1 []string, sentence2 []string, similarPairs [][]string) bool {
	if len(sentence1) != len(sentence2) {
		return false
	}
	// Build bidirectional map
	pairMap := make(map[string]map[string]bool)
	for _, p := range similarPairs {
		a, b := p[0], p[1]
		if pairMap[a] == nil {
			pairMap[a] = make(map[string]bool)
		}
		if pairMap[b] == nil {
			pairMap[b] = make(map[string]bool)
		}
		pairMap[a][b] = true
		pairMap[b][a] = true
	}

	for i := range sentence1 {
		w1, w2 := sentence1[i], sentence2[i]
		if w1 == w2 {
			continue
		}
		if !pairMap[w1][w2] {
			return false
		}
	}
	return true
}
```

## 0744 — Find Smallest Letter Greater Than Target

```go
package main

// LeetCode #744: Find Smallest Letter Greater Than Target
// https://leetcode.com/problems/find-smallest-letter-greater-than-target/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(string(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'a'))) // 'c'
	fmt.Println(string(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'c'))) // 'f'
	fmt.Println(string(nextGreatestLetter([]byte{'x', 'x', 'y', 'y'}, 'z'))) // 'x'
}

// nextGreatestLetter finds the smallest letter in letters that is greater than target.
// Time: O(log n). Space: O(1).
func nextGreatestLetter(letters []byte, target byte) byte {
	l, r := 0, len(letters)
	for l < r {
		mid := l + (r-l)/2
		if letters[mid] <= target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return letters[l%len(letters)]
}
```

## 0746 — Min Cost Climbing Stairs

```go
package main

// LeetCode #746: Min Cost Climbing Stairs
// https://leetcode.com/problems/min-cost-climbing-stairs/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(minCostClimbingStairs([]int{10, 15, 20}))                    // 15
	fmt.Println(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1})) // 6
}

// minCostClimbingStairs returns the minimum cost to reach the top of the stairs.
// Time: O(n). Space: O(1).
func minCostClimbingStairs(cost []int) int {
	a, b := cost[0], cost[1]
	for i := 2; i < len(cost); i++ {
		a, b = b, min(a, b)+cost[i]
	}
	return min(a, b)
}
```

## 0747 — Largest Number At Least Twice Of Others

```go
package main

// LeetCode #747: Largest Number At Least Twice of Others
// https://leetcode.com/problems/largest-number-at-least-twice-of-others/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(dominantIndex([]int{3, 6, 1, 0}))   // 1
	fmt.Println(dominantIndex([]int{1, 2, 3, 4}))    // -1
	fmt.Println(dominantIndex([]int{1}))              // 0
}

// dominantIndex returns the index of the largest element if it is at least twice as large as all others.
// Time: O(n). Space: O(1).
func dominantIndex(nums []int) int {
	maxIdx := 0
	for i, v := range nums {
		if v > nums[maxIdx] {
			maxIdx = i
		}
	}
	for i, v := range nums {
		if i != maxIdx && v*2 > nums[maxIdx] {
			return -1
		}
	}
	return maxIdx
}
```

## 0748 — Shortest Completing Word

```go
package main

// LeetCode #748: Shortest Completing Word
// https://leetcode.com/problems/shortest-completing-word/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(shortestCompletingWord("1s3 PSt", []string{"step", "steps", "stripe", "stepple"})) // "steps"
	fmt.Println(shortestCompletingWord("1s3 456", []string{"looks", "pest", "stew", "show"}))     // "pest"
}

// shortestCompletingWord finds the shortest word that contains all letters in licensePlate.
// Time: O(n * m) where n = len(words), m = avg word length. Space: O(1).
func shortestCompletingWord(licensePlate string, words []string) string {
	targetCount := [26]int{}
	for _, c := range strings.ToLower(licensePlate) {
		if c >= 'a' && c <= 'z' {
			targetCount[c-'a']++
		}
	}

	result := ""
	for _, word := range words {
		if result != "" && len(word) >= len(result) {
			continue
		}
		wordCount := [26]int{}
		for _, c := range word {
			wordCount[c-'a']++
		}
		ok := true
		for i := 0; i < 26; i++ {
			if wordCount[i] < targetCount[i] {
				ok = false
				break
			}
		}
		if ok {
			result = word
		}
	}
	return result
}
```

## 0760 — Find Anagram Mappings

```go
package main

// LeetCode #760: Find Anagram Mappings
// https://leetcode.com/problems/find-anagram-mappings/
// Difficulty: Easy [Paid]
// Note: This is a premium problem. Implementation based on public description.

import "fmt"

func main() {
	fmt.Println(anagramMappings([]int{12, 28, 46, 32, 50}, []int{50, 12, 32, 46, 28})) // [1,4,3,2,0]
	fmt.Println(anagramMappings([]int{1, 2}, []int{2, 1}))                              // [1,0]
}

// anagramMappings returns a mapping array P where P[i] is the index of A[i] in B.
// Time: O(n). Space: O(n).
func anagramMappings(nums1 []int, nums2 []int) []int {
	pos := make(map[int]int)
	for i, v := range nums2 {
		pos[v] = i
	}
	result := make([]int, len(nums1))
	for i, v := range nums1 {
		result[i] = pos[v]
	}
	return result
}
```

## 0762 — Prime Number Of Set Bits In Binary Representation

```go
package main

// LeetCode #762: Prime Number of Set Bits in Binary Representation
// https://leetcode.com/problems/prime-number-of-set-bits-in-binary-representation/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(countPrimeSetBits(6, 10))   // 4
	fmt.Println(countPrimeSetBits(10, 15))  // 5
	fmt.Println(countPrimeSetBits(1, 2))    // 1
}

// countPrimeSetBits counts numbers in [left, right] whose binary representation has a prime number of set bits.
// Time: O((right-left+1) * log n). Space: O(1).
func countPrimeSetBits(left int, right int) int {
	// Primes up to 20 (since max int is 10^6, < 2^20)
	primes := map[int]bool{2: true, 3: true, 5: true, 7: true, 11: true, 13: true, 17: true, 19: true}
	count := 0
	for n := left; n <= right; n++ {
		bits := 0
		for x := n; x > 0; x >>= 1 {
			bits += x & 1
		}
		if primes[bits] {
			count++
		}
	}
	return count
}
```

## 0766 — Toeplitz Matrix

```go
package main

// LeetCode #766: Toeplitz Matrix
// https://leetcode.com/problems/toeplitz-matrix/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(isToeplitzMatrix([][]int{{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2}})) // true
	fmt.Println(isToeplitzMatrix([][]int{{1, 2}, {2, 2}}))                            // false
}

// isToeplitzMatrix checks if every diagonal from top-left to bottom-right has the same element.
// Time: O(m*n). Space: O(1).
func isToeplitzMatrix(matrix [][]int) bool {
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] != matrix[i-1][j-1] {
				return false
			}
		}
	}
	return true
}
```

## 0771 — Jewels And Stones

```go
package main

// LeetCode #771: Jewels and Stones
// https://leetcode.com/problems/jewels-and-stones/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(numJewelsInStones("aA", "aAAbbbb")) // 3
	fmt.Println(numJewelsInStones("z", "ZZ"))       // 0
	fmt.Println(numJewelsInStones("", "abc"))       // 0
}

// numJewelsInStones counts how many stones are also jewels.
// Time: O(j + s). Space: O(j).
func numJewelsInStones(jewels string, stones string) int {
	jSet := make(map[rune]bool)
	for _, c := range jewels {
		jSet[c] = true
	}
	count := 0
	for _, c := range stones {
		if jSet[c] {
			count++
		}
	}
	return count
}
```

## 0783 — Minimum Distance Between Bst Nodes

```go
package main

// LeetCode #783: Minimum Distance Between BST Nodes
// https://leetcode.com/problems/minimum-distance-between-bst-nodes/
// Difficulty: Easy

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
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{Val: 6},
	}
	fmt.Println(minDiffInBST(root)) // 1
}

// minDiffInBST finds the minimum difference between values of any two nodes in a BST.
// Time: O(n). Space: O(n).
func minDiffInBST(root *TreeNode) int {
	minDiff := math.MaxInt32
	var prev *int
	inorder(root, &prev, &minDiff)
	return minDiff
}

func inorder(node *TreeNode, prev **int, minDiff *int) {
	if node == nil {
		return
	}
	inorder(node.Left, prev, minDiff)
	if *prev != nil {
		diff := node.Val - **prev
		if diff < *minDiff {
			*minDiff = diff
		}
	}
	*prev = &node.Val
	inorder(node.Right, prev, minDiff)
}
```

## 0796 — Rotate String

```go
package main

// LeetCode #796: Rotate String
// https://leetcode.com/problems/rotate-string/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(rotateString("abcde", "cdeab")) // true
	fmt.Println(rotateString("abcde", "abced")) // false
}

// rotateString checks if goal can be obtained by rotating s.
// Time: O(n). Space: O(n).
func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	return strings.Contains(s+s, goal)
}
```

## 0800 — Similar Rgb Color

```go
package main

// LeetCode #800: Similar RGB Color
// https://leetcode.com/problems/similar-rgb-color/
// Difficulty: Easy [Paid]
// Note: This is a premium problem. Implementation based on public description.

import "fmt"

func main() {
	fmt.Println(similarRGB("#09f166")) // "#11ee66"
	fmt.Println(similarRGB("#4e3fe1")) // "#44ee11" (note: this is an approximation)
}

// similarRGB finds the most similar RGB shorthand for the given color.
// Time: O(1). Space: O(1).
func similarRGB(color string) string {
	// For each component, find the nearest shorthand (00, 11, ..., ff)
	result := "#"
	for i := 1; i < len(color); i += 2 {
		// Extract the 2-digit hex value
		val := 0
		for j := 0; j < 2; j++ {
			c := color[i+j]
			if c >= '0' && c <= '9' {
				val = val*16 + int(c-'0')
			} else {
				val = val*16 + int(c-'a'+10)
			}
		}
		// Find the nearest shorthand: 0x00, 0x11, 0x22, ..., 0xff
		nearest := (val + 8) / 17 * 17
		if nearest > 255 {
			nearest = 255
		}
		hex := fmt.Sprintf("%02x", nearest)
		result += hex
	}
	return result
}
```

## 0804 — Unique Morse Code Words

```go
package main

// LeetCode #804: Unique Morse Code Words
// https://leetcode.com/problems/unique-morse-code-words/
// Difficulty: Easy

import "fmt"

var morse = []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

func main() {
	fmt.Println(uniqueMorseRepresentations([]string{"gin", "zen", "gig", "msg"})) // 2
	fmt.Println(uniqueMorseRepresentations([]string{"a"}))                        // 1
}

// uniqueMorseRepresentations counts unique Morse code transformations of words.
// Time: O(n * m) where n = len(words), m = avg len. Space: O(n).
func uniqueMorseRepresentations(words []string) int {
	set := make(map[string]bool)
	for _, word := range words {
		var code string
		for _, c := range word {
			code += morse[c-'a']
		}
		set[code] = true
	}
	return len(set)
}
```

## 0806 — Number Of Lines To Write String

```go
package main

// LeetCode #806: Number of Lines To Write String
// https://leetcode.com/problems/number-of-lines-to-write-string/
// Difficulty: Easy

import "fmt"

func main() {
	widths := []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}
	fmt.Println(numberOfLines(widths, "abcdefghijklmnopqrstuvwxyz")) // [3, 60]

	widths2 := []int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}
	fmt.Println(numberOfLines(widths2, "bbbcccdddaaa")) // [2, 4]
}

// numberOfLines returns the lines and last line width needed to write the string.
// Time: O(n). Space: O(1).
func numberOfLines(widths []int, s string) []int {
	lines, currentWidth := 1, 0
	for _, c := range s {
		w := widths[c-'a']
		if currentWidth+w > 100 {
			lines++
			currentWidth = w
		} else {
			currentWidth += w
		}
	}
	return []int{lines, currentWidth}
}
```

## 0812 — Largest Triangle Area

```go
package main

// LeetCode #812: Largest Triangle Area
// https://leetcode.com/problems/largest-triangle-area/
// Difficulty: Easy

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(largestTriangleArea([][]int{{0, 0}, {0, 1}, {1, 0}, {0, 2}, {2, 0}})) // 2.0
	fmt.Println(largestTriangleArea([][]int{{1, 0}, {0, 0}, {0, 1}}))                 // 0.5
}

// largestTriangleArea finds the largest area of any triangle formed by any 3 points.
// Time: O(n^3). Space: O(1).
func largestTriangleArea(points [][]int) float64 {
	maxArea := 0.0
	n := len(points)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				area := areaTriangle(points[i], points[j], points[k])
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}
	return maxArea
}

func areaTriangle(a, b, c []int) float64 {
	return math.Abs(float64(a[0]*(b[1]-c[1])+b[0]*(c[1]-a[1])+c[0]*(a[1]-b[1]))) / 2.0
}
```

## 0819 — Most Common Word

```go
package main

// LeetCode #819: Most Common Word
// https://leetcode.com/problems/most-common-word/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(mostCommonWord("Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"})) // "ball"
	fmt.Println(mostCommonWord("a.", []string{}))                                                          // "a"
}

// mostCommonWord finds the most frequent word in paragraph that is not banned.
// Time: O(n + m). Space: O(n + m).
func mostCommonWord(paragraph string, banned []string) string {
	bannedSet := make(map[string]bool)
	for _, w := range banned {
		bannedSet[w] = true
	}

	// Normalize: lowercase and split by non-letter characters
	normalized := strings.ToLower(paragraph)
	cleaned := strings.NewReplacer("!", " ", "?", " ", "'", " ", ",", " ", ";", " ", ".", " ").Replace(normalized)

	words := strings.Fields(cleaned)
	counts := make(map[string]int)
	maxCount := 0
	result := ""
	for _, w := range words {
		if bannedSet[w] {
			continue
		}
		counts[w]++
		if counts[w] > maxCount {
			maxCount = counts[w]
			result = w
		}
	}
	return result
}
```

## 0821 — Shortest Distance To A Character

```go
package main

// LeetCode #821: Shortest Distance to a Character
// https://leetcode.com/problems/shortest-distance-to-a-character/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(shortestToChar("loveleetcode", 'e')) // [3,2,1,0,1,0,0,1,2,2,1,0]
	fmt.Println(shortestToChar("aaab", 'b'))          // [3,2,1,0]
}

// shortestToChar returns the shortest distance from each character to the target character c.
// Time: O(n). Space: O(1) excluding output.
func shortestToChar(s string, c byte) []int {
	n := len(s)
	result := make([]int, n)
	// Initialize with large value
	for i := range result {
		result[i] = n
	}

	// Left to right
	last := -n
	for i := 0; i < n; i++ {
		if s[i] == c {
			last = i
		}
		result[i] = min(result[i], i-last)
	}

	// Right to left
	last = 2 * n
	for i := n - 1; i >= 0; i-- {
		if s[i] == c {
			last = i
		}
		result[i] = min(result[i], last-i)
	}
	return result
}
```

## 0824 — Goat Latin

```go
package main

// LeetCode #824: Goat Latin
// https://leetcode.com/problems/goat-latin/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(toGoatLatin("I speak Goat Latin"))  // "Imaa peaksmaaa oatGmaaaa atinLmaaaaa"
	fmt.Println(toGoatLatin("The quick brown fox")) // "heTmaa uickqmaaa rainbowmaaaa oxfmaaaaa"
}

// toGoatLatin converts a sentence to Goat Latin.
// Time: O(n). Space: O(n).
func toGoatLatin(sentence string) string {
	vowels := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'A': true, 'E': true, 'I': true, 'O': true, 'U': true}
	words := strings.Fields(sentence)
	var result []string
	for i, word := range words {
		var transformed string
		if vowels[word[0]] {
			transformed = word + "ma"
		} else {
			transformed = word[1:] + string(word[0]) + "ma"
		}
		transformed += strings.Repeat("a", i+1)
		result = append(result, transformed)
	}
	return strings.Join(result, " ")
}
```

