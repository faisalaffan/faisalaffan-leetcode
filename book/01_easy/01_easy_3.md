# Easy (Mudah) — Problem ��1221

## 0830 — Positions Of Large Groups

```go
package main

// LeetCode #830: Positions of Large Groups
// https://leetcode.com/problems/positions-of-large-groups/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(largeGroupPositions("abbxxxxzzy"))    // [[3,6]]
	fmt.Println(largeGroupPositions("abc"))           // []
	fmt.Println(largeGroupPositions("abcdddeeeeaabbbcd")) // [[3,5],[6,9],[12,14]]
}

// largeGroupPositions finds all large groups (consecutive identical characters of length >= 3).
// Time: O(n). Space: O(1) excluding output.
func largeGroupPositions(s string) [][]int {
	result := make([][]int, 0)
	start := 0
	for i := 1; i <= len(s); i++ {
		if i == len(s) || s[i] != s[start] {
			if i-start >= 3 {
				result = append(result, []int{start, i - 1})
			}
			start = i
		}
	}
	return result
}
```

## 0832 — Flipping An Image

```go
package main

// LeetCode #832: Flipping an Image
// https://leetcode.com/problems/flipping-an-image/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(flipAndInvertImage([][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}})) // [[1,0,0],[0,1,0],[1,1,1]]
	fmt.Println(flipAndInvertImage([][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}}))
}

// flipAndInvertImage flips the image horizontally then inverts it.
// Time: O(m*n). Space: O(1) in-place.
func flipAndInvertImage(image [][]int) [][]int {
	for _, row := range image {
		l, r := 0, len(row)-1
		for l <= r {
			row[l], row[r] = 1-row[r], 1-row[l]
			l++
			r--
		}
	}
	return image
}
```

## 0836 — Rectangle Overlap

```go
package main

// LeetCode #836: Rectangle Overlap
// https://leetcode.com/problems/rectangle-overlap/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(isRectangleOverlap([]int{0, 0, 2, 2}, []int{1, 1, 3, 3})) // true
	fmt.Println(isRectangleOverlap([]int{0, 0, 1, 1}, []int{1, 0, 2, 1})) // false
	fmt.Println(isRectangleOverlap([]int{0, 0, 1, 1}, []int{2, 2, 3, 3})) // false
}

// isRectangleOverlap checks if two rectangles overlap (positive area).
// Time: O(1). Space: O(1).
func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	// Check if one rectangle is to the left of the other
	if rec1[2] <= rec2[0] || rec2[2] <= rec1[0] {
		return false
	}
	// Check if one rectangle is above the other
	if rec1[3] <= rec2[1] || rec2[3] <= rec1[1] {
		return false
	}
	return true
}
```

## 0844 — Backspace String Compare

```go
package main

// LeetCode #844: Backspace String Compare
// https://leetcode.com/problems/backspace-string-compare/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(backspaceCompare("ab#c", "ad#c")) // true
	fmt.Println(backspaceCompare("ab##", "c#d#")) // true
	fmt.Println(backspaceCompare("a#c", "b"))     // false
	fmt.Println(backspaceCompare("a##c", "#a#c")) // true
}

// backspaceCompare checks if two strings are equal after applying backspace.
// Time: O(n + m). Space: O(1).
func backspaceCompare(s string, t string) bool {
	i, j := len(s)-1, len(t)-1
	skipS, skipT := 0, 0
	for i >= 0 || j >= 0 {
		// Find next valid char in s
		for i >= 0 {
			if s[i] == '#' {
				skipS++
				i--
			} else if skipS > 0 {
				skipS--
				i--
			} else {
				break
			}
		}
		// Find next valid char in t
		for j >= 0 {
			if t[j] == '#' {
				skipT++
				j--
			} else if skipT > 0 {
				skipT--
				j--
			} else {
				break
			}
		}
		if i >= 0 && j >= 0 {
			if s[i] != t[j] {
				return false
			}
		} else if i >= 0 || j >= 0 {
			return false
		}
		i--
		j--
	}
	return true
}
```

## 0859 — Buddy Strings

```go
package main

// LeetCode #859: Buddy Strings
// https://leetcode.com/problems/buddy-strings/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(buddyStrings("ab", "ba"))   // true
	fmt.Println(buddyStrings("ab", "ab"))   // false
	fmt.Println(buddyStrings("aa", "aa"))   // true
	fmt.Println(buddyStrings("abcd", "badc")) // false
}

// buddyStrings checks if swapping two letters in s makes it equal to goal.
// Time: O(n). Space: O(1).
func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	if s == goal {
		// Need at least one duplicate character to swap
		seen := make(map[byte]bool)
		for i := range s {
			if seen[s[i]] {
				return true
			}
			seen[s[i]] = true
		}
		return false
	}
	diff := make([]int, 0)
	for i := range s {
		if s[i] != goal[i] {
			diff = append(diff, i)
			if len(diff) > 2 {
				return false
			}
		}
	}
	return len(diff) == 2 && s[diff[0]] == goal[diff[1]] && s[diff[1]] == goal[diff[0]]
}
```

## 0860 — Lemonade Change

```go
package main

// LeetCode #860: Lemonade Change
// https://leetcode.com/problems/lemonade-change/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(lemonadeChange([]int{5, 5, 5, 10, 20})) // true
	fmt.Println(lemonadeChange([]int{5, 5, 10, 10, 20})) // false
	fmt.Println(lemonadeChange([]int{5, 5, 5, 10, 5, 20, 5, 10, 5, 20})) // true
}

// lemonadeChange checks if we can provide correct change for each customer.
// Time: O(n). Space: O(1).
func lemonadeChange(bills []int) bool {
	fives, tens := 0, 0
	for _, bill := range bills {
		switch bill {
		case 5:
			fives++
		case 10:
			if fives == 0 {
				return false
			}
			fives--
			tens++
		case 20:
			if tens > 0 && fives > 0 {
				tens--
				fives--
			} else if fives >= 3 {
				fives -= 3
			} else {
				return false
			}
		}
	}
	return true
}
```

## 0867 — Transpose Matrix

```go
package main

// LeetCode #867: Transpose Matrix
// https://leetcode.com/problems/transpose-matrix/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(transpose([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})) // [[1,4,7],[2,5,8],[3,6,9]]
	fmt.Println(transpose([][]int{{1, 2, 3}, {4, 5, 6}}))             // [[1,4],[2,5],[3,6]]
}

// transpose returns the transpose of a matrix.
// Time: O(m*n). Space: O(m*n).
func transpose(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, m)
		for j := 0; j < m; j++ {
			result[i][j] = matrix[j][i]
		}
	}
	return result
}
```

## 0868 — Binary Gap

```go
package main

// LeetCode #868: Binary Gap
// https://leetcode.com/problems/binary-gap/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(binaryGap(22))  // 2 (10110)
	fmt.Println(binaryGap(8))   // 0 (1000)
	fmt.Println(binaryGap(5))   // 2 (101)
	fmt.Println(binaryGap(6))   // 1 (110)
}

// binaryGap finds the longest distance between two consecutive 1s in binary representation.
// Time: O(log n). Space: O(1).
func binaryGap(n int) int {
	last := -1
	maxDist := 0
	for i := 0; n > 0; i++ {
		if n&1 == 1 {
			if last != -1 {
				if i-last > maxDist {
					maxDist = i - last
				}
			}
			last = i
		}
		n >>= 1
	}
	return maxDist
}
```

## 0872 — Leaf Similar Trees

```go
package main

// LeetCode #872: Leaf-Similar Trees
// https://leetcode.com/problems/leaf-similar-trees/
// Difficulty: Easy

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// [3,5,1,6,2,9,8,null,null,7,4]
	root1 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   5,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 4}},
		},
		Right: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 9},
			Right: &TreeNode{Val: 8},
		},
	}
	root2 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   5,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 7},
		},
		Right: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 8}},
		},
	}
	fmt.Println(leafSimilar(root1, root2)) // true

	root3 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	root4 := &TreeNode{Val: 2, Left: &TreeNode{Val: 2}}
	fmt.Println(leafSimilar(root3, root4)) // false
}

// leafSimilar checks if two trees have the same leaf value sequence.
// Time: O(n + m). Space: O(n + m).
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leaves1 := make([]int, 0)
	collectLeaves(root1, &leaves1)
	leaves2 := make([]int, 0)
	collectLeaves(root2, &leaves2)
	if len(leaves1) != len(leaves2) {
		return false
	}
	for i, v := range leaves1 {
		if v != leaves2[i] {
			return false
		}
	}
	return true
}

func collectLeaves(node *TreeNode, leaves *[]int) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil {
		*leaves = append(*leaves, node.Val)
		return
	}
	collectLeaves(node.Left, leaves)
	collectLeaves(node.Right, leaves)
}
```

## 0876 — Middle Of The Linked List

```go
package main

// LeetCode #876: Middle of the Linked List
// https://leetcode.com/problems/middle-of-the-linked-list/
// Difficulty: Easy

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// [1,2,3,4,5] => 3
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	fmt.Println(middleNode(head).Val) // 3

	// [1,2,3,4,5,6] => 4
	head2 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}
	fmt.Println(middleNode(head2).Val) // 4
}

// middleNode returns the middle node of a linked list.
// Time: O(n). Space: O(1).
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
```

## 0883 — Projection Area Of 3d Shapes

```go
package main

// LeetCode #883: Projection Area of 3D Shapes
// https://leetcode.com/problems/projection-area-of-3d-shapes/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(projectionArea([][]int{{1, 2}, {3, 4}})) // 17
	fmt.Println(projectionArea([][]int{{2}}))             // 5
	fmt.Println(projectionArea([][]int{{1, 0}, {0, 2}})) // 8
}

// projectionArea returns the total area of the 3D shape projection.
// Time: O(n^2). Space: O(1).
func projectionArea(grid [][]int) int {
	n := len(grid)
	xy := 0
	xz := 0
	yz := 0
	for i := 0; i < n; i++ {
		maxRow, maxCol := 0, 0
		for j := 0; j < n; j++ {
			if grid[i][j] > 0 {
				xy++
			}
			if grid[i][j] > maxRow {
				maxRow = grid[i][j]
			}
			if grid[j][i] > maxCol {
				maxCol = grid[j][i]
			}
		}
		xz += maxRow
		yz += maxCol
	}
	return xy + xz + yz
}
```

## 0884 — Uncommon Words From Two Sentences

```go
package main

// LeetCode #884: Uncommon Words from Two Sentences
// https://leetcode.com/problems/uncommon-words-from-two-sentences/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(uncommonFromSentences("this apple is sweet", "this apple is sour")) // [sweet sour]
	fmt.Println(uncommonFromSentences("apple apple", "banana"))                     // [banana]
}

// uncommonFromSentences returns all uncommon words across two sentences.
// Time: O(n + m). Space: O(n + m).
func uncommonFromSentences(s1 string, s2 string) []string {
	count := make(map[string]int)
	for _, w := range strings.Fields(s1) {
		count[w]++
	}
	for _, w := range strings.Fields(s2) {
		count[w]++
	}
	result := make([]string, 0)
	for w, c := range count {
		if c == 1 {
			result = append(result, w)
		}
	}
	return result
}
```

## 0888 — Fair Candy Swap

```go
package main

// LeetCode #888: Fair Candy Swap
// https://leetcode.com/problems/fair-candy-swap/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(fairCandySwap([]int{1, 1}, []int{2, 2}))       // [1,2]
	fmt.Println(fairCandySwap([]int{1, 2}, []int{2, 3}))       // [1,2]
	fmt.Println(fairCandySwap([]int{2}, []int{1, 3}))          // [2,3]
}

// fairCandySwap finds a candy swap that makes both Alice and Bob have equal candy.
// Time: O(n + m). Space: O(m).
func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	sumA, sumB := 0, 0
	bSet := make(map[int]bool)
	for _, v := range aliceSizes {
		sumA += v
	}
	for _, v := range bobSizes {
		sumB += v
		bSet[v] = true
	}
	diff := (sumB - sumA) / 2
	for _, a := range aliceSizes {
		if bSet[a+diff] {
			return []int{a, a + diff}
		}
	}
	return nil
}
```

## 0892 — Surface Area Of 3d Shapes

```go
package main

// LeetCode #892: Surface Area of 3D Shapes
// https://leetcode.com/problems/surface-area-of-3d-shapes/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(surfaceArea([][]int{{2}}))                                   // 10
	fmt.Println(surfaceArea([][]int{{1, 2}, {3, 4}}))                       // 34
}

// surfaceArea calculates the surface area of a 3D shape on a grid.
// Time: O(n^2). Space: O(1).
func surfaceArea(grid [][]int) int {
	n := len(grid)
	area := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] > 0 {
				// Top + bottom
				area += 2
				// Four sides - subtract adjacent overlaps
				area += grid[i][j] * 4
				if i > 0 {
					area -= min(grid[i][j], grid[i-1][j]) * 2
				}
				if j > 0 {
					area -= min(grid[i][j], grid[i][j-1]) * 2
				}
			}
		}
	}
	return area
}
```

## 0896 — Monotonic Array

```go
package main

// LeetCode #896: Monotonic Array
// https://leetcode.com/problems/monotonic-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(isMonotonic([]int{1, 2, 2, 3}))   // true
	fmt.Println(isMonotonic([]int{6, 5, 4, 4}))   // true
	fmt.Println(isMonotonic([]int{1, 3, 2}))      // false
	fmt.Println(isMonotonic([]int{1, 1, 1}))      // true
}

// isMonotonic checks if the array is monotonic (either non-decreasing or non-increasing).
// Time: O(n). Space: O(1).
func isMonotonic(nums []int) bool {
	inc, dec := true, true
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			dec = false
		}
		if nums[i] < nums[i-1] {
			inc = false
		}
	}
	return inc || dec
}
```

## 0897 — Increasing Order Search Tree

```go
package main

// LeetCode #897: Increasing Order Search Tree
// https://leetcode.com/problems/increasing-order-search-tree/
// Difficulty: Easy

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
			Val:   3,
			Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 1}},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:   6,
			Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 9}},
		},
	}
	result := increasingBST(root)
	// Print inorder to verify
	printTree(result) // 1 2 3 4 5 6 7 8 9
}

func printTree(root *TreeNode) {
	for root != nil {
		fmt.Print(root.Val, " ")
		root = root.Right
	}
	fmt.Println()
}

// increasingBST rearranges the BST into increasing order (only right children).
// Time: O(n). Space: O(n).
func increasingBST(root *TreeNode) *TreeNode {
	var newRoot, prev *TreeNode
	inorderTree(root, &newRoot, &prev)
	return newRoot
}

func inorderTree(node *TreeNode, newRoot **TreeNode, prev **TreeNode) {
	if node == nil {
		return
	}
	inorderTree(node.Left, newRoot, prev)
	if *newRoot == nil {
		*newRoot = node
	}
	if *prev != nil {
		(*prev).Right = node
	}
	node.Left = nil
	*prev = node
	inorderTree(node.Right, newRoot, prev)
}
```

## 0905 — Sort Array By Parity

```go
package main

// LeetCode #905: Sort Array By Parity
// https://leetcode.com/problems/sort-array-by-parity/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(sortArrayByParity([]int{3, 1, 2, 4})) // [2,4,3,1] or [4,2,1,3] etc.
	fmt.Println(sortArrayByParity([]int{0}))           // [0]
	fmt.Println(sortArrayByParity([]int{1, 3, 5}))     // [1,3,5]
}

// sortArrayByParity moves all even numbers to the front, odd to the back.
// Time: O(n). Space: O(1).
func sortArrayByParity(nums []int) []int {
	l, r := 0, len(nums)-1
	for l < r {
		if nums[l]%2 == 0 {
			l++
		} else {
			nums[l], nums[r] = nums[r], nums[l]
			r--
		}
	}
	return nums
}
```

## 0908 — Smallest Range I

```go
package main

// LeetCode #908: Smallest Range I
// https://leetcode.com/problems/smallest-range-i/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(smallestRangeI([]int{1}, 0))          // 0
	fmt.Println(smallestRangeI([]int{0, 10}, 2))      // 6
	fmt.Println(smallestRangeI([]int{1, 3, 6}, 3))    // 0
}

// smallestRangeI returns the smallest possible range after modifying each element by at most k.
// Time: O(n). Space: O(1).
func smallestRangeI(nums []int, k int) int {
	minVal, maxVal := nums[0], nums[0]
	for _, v := range nums[1:] {
		if v < minVal {
			minVal = v
		}
		if v > maxVal {
			maxVal = v
		}
	}
	diff := (maxVal - k) - (minVal + k)
	if diff < 0 {
		return 0
	}
	return diff
}
```

## 0914 — X Of A Kind In A Deck Of Cards

```go
package main

// LeetCode #914: X of a Kind in a Deck of Cards
// https://leetcode.com/problems/x-of-a-kind-in-a-deck-of-cards/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(hasGroupsSizeX([]int{1, 2, 3, 4, 4, 3, 2, 1})) // true
	fmt.Println(hasGroupsSizeX([]int{1, 1, 1, 2, 2, 2, 3, 3})) // false
	fmt.Println(hasGroupsSizeX([]int{1}))                        // false
}

// hasGroupsSizeX checks if the deck can be split into groups of equal size with same values.
// Time: O(n log m) where m is max count. Space: O(n).
func hasGroupsSizeX(deck []int) bool {
	counts := make(map[int]int)
	for _, v := range deck {
		counts[v]++
	}
	g := -1
	for _, c := range counts {
		if g == -1 {
			g = c
		} else {
			g = gcd(g, c)
		}
	}
	return g >= 2
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
```

## 0917 — Reverse Only Letters

```go
package main

// LeetCode #917: Reverse Only Letters
// https://leetcode.com/problems/reverse-only-letters/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(reverseOnlyLetters("ab-cd"))      // "dc-ba"
	fmt.Println(reverseOnlyLetters("a-bC-dEf-ghIj")) // "j-Ih-gfE-dCba"
	fmt.Println(reverseOnlyLetters("Test1ng-Leet=code-Q!")) // "Qedo1ct-eeLg=ntse-T!"
}

// reverseOnlyLetters reverses only the letters in the string, keeping non-letters in place.
// Time: O(n). Space: O(n).
func reverseOnlyLetters(s string) string {
	b := []byte(s)
	l, r := 0, len(b)-1
	for l < r {
		if !isLetter(b[l]) {
			l++
			continue
		}
		if !isLetter(b[r]) {
			r--
			continue
		}
		b[l], b[r] = b[r], b[l]
		l++
		r--
	}
	return string(b)
}

func isLetter(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}
```

## 0922 — Sort Array By Parity Ii

```go
package main

// LeetCode #922: Sort Array By Parity II
// https://leetcode.com/problems/sort-array-by-parity-ii/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(sortArrayByParityII([]int{4, 2, 5, 7})) // [4,5,2,7] or [4,7,2,5]
	fmt.Println(sortArrayByParityII([]int{2, 3}))        // [2,3]
}

// sortArrayByParityII puts even numbers at even indices, odd numbers at odd indices.
// Time: O(n). Space: O(1).
func sortArrayByParityII(nums []int) []int {
	j := 1 // odd pointer
	for i := 0; i < len(nums); i += 2 {
		if nums[i]%2 == 1 {
			for nums[j]%2 == 1 {
				j += 2
			}
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	return nums
}
```

## 0925 — Long Pressed Name

```go
package main

// LeetCode #925: Long Pressed Name
// https://leetcode.com/problems/long-pressed-name/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(isLongPressedName("alex", "aaleex"))    // true
	fmt.Println(isLongPressedName("saeed", "ssaaedd"))  // false
	fmt.Println(isLongPressedName("leelee", "lleeelee")) // true
	fmt.Println(isLongPressedName("alex", "aaleexa"))   // false
}

// isLongPressedName checks if typed is a long-pressed version of name.
// Time: O(n + m). Space: O(1).
func isLongPressedName(name string, typed string) bool {
	if len(typed) < len(name) {
		return false
	}
	i, j := 0, 0
	for i < len(name) && j < len(typed) {
		if name[i] != typed[j] {
			return false
		}
		// Count occurrences in name
		c1 := 0
		ch := name[i]
		for i < len(name) && name[i] == ch {
			i++
			c1++
		}
		// Count occurrences in typed
		c2 := 0
		for j < len(typed) && typed[j] == ch {
			j++
			c2++
		}
		if c2 < c1 {
			return false
		}
	}
	return i == len(name) && j == len(typed)
}
```

## 0929 — Unique Email Addresses

```go
package main

// LeetCode #929: Unique Email Addresses
// https://leetcode.com/problems/unique-email-addresses/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(numUniqueEmails([]string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"})) // 2
	fmt.Println(numUniqueEmails([]string{"a@leetcode.com", "b@leetcode.com", "c@leetcode.com"}))                                                    // 3
}

// numUniqueEmails counts unique email addresses after applying normalization rules.
// Time: O(n * m). Space: O(n).
func numUniqueEmails(emails []string) int {
	set := make(map[string]bool)
	for _, email := range emails {
		parts := strings.Split(email, "@")
		local := parts[0]
		domain := parts[1]

		// Remove dots and everything after '+'
		cleaned := strings.ReplaceAll(local, ".", "")
		if plusIdx := strings.Index(cleaned, "+"); plusIdx != -1 {
			cleaned = cleaned[:plusIdx]
		}
		set[cleaned+"@"+domain] = true
	}
	return len(set)
}
```

## 0933 — Number Of Recent Calls

```go
package main

// LeetCode #933: Number of Recent Calls
// https://leetcode.com/problems/number-of-recent-calls/
// Difficulty: Easy

import "fmt"

// RecentCounter counts recent requests within a 3000ms window.
type RecentCounter struct {
	queue []int
}

// Constructor creates a RecentCounter.
func Constructor() RecentCounter {
	return RecentCounter{queue: make([]int, 0)}
}

// Ping adds a new request and returns the count of requests in the last 3000ms.
// Time: O(1) amortized. Space: O(n).
func (rc *RecentCounter) Ping(t int) int {
	rc.queue = append(rc.queue, t)
	for rc.queue[0] < t-3000 {
		rc.queue = rc.queue[1:]
	}
	return len(rc.queue)
}

func main() {
	rc := Constructor()
	fmt.Println(rc.Ping(1))    // 1
	fmt.Println(rc.Ping(100))  // 2
	fmt.Println(rc.Ping(3001)) // 3
	fmt.Println(rc.Ping(3002)) // 3
}
```

## 0938 — Range Sum Of Bst

```go
package main

// LeetCode #938: Range Sum of BST
// https://leetcode.com/problems/range-sum-of-bst/
// Difficulty: Easy

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val:   5,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 7},
		},
		Right: &TreeNode{
			Val:   15,
			Right: &TreeNode{Val: 18},
		},
	}
	fmt.Println(rangeSumBST(root, 7, 15)) // 32
}

// rangeSumBST returns the sum of all node values in the range [low, high].
// Time: O(n). Space: O(n).
func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	if root.Val > high {
		return rangeSumBST(root.Left, low, high)
	}
	if root.Val < low {
		return rangeSumBST(root.Right, low, high)
	}
	return root.Val + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
}
```

## 0941 — Valid Mountain Array

```go
package main

// LeetCode #941: Valid Mountain Array
// https://leetcode.com/problems/valid-mountain-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(validMountainArray([]int{2, 1}))          // false
	fmt.Println(validMountainArray([]int{3, 5, 5}))       // false
	fmt.Println(validMountainArray([]int{0, 3, 2, 1}))    // true
	fmt.Println(validMountainArray([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})) // false
}

// validMountainArray checks if the array is a valid mountain array.
// Time: O(n). Space: O(1).
func validMountainArray(arr []int) bool {
	n := len(arr)
	if n < 3 {
		return false
	}
	i := 0
	// Climb up
	for i+1 < n && arr[i] < arr[i+1] {
		i++
	}
	// Peak cannot be first or last
	if i == 0 || i == n-1 {
		return false
	}
	// Climb down
	for i+1 < n && arr[i] > arr[i+1] {
		i++
	}
	return i == n-1
}
```

## 0942 — Di String Match

```go
package main

// LeetCode #942: DI String Match
// https://leetcode.com/problems/di-string-match/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(diStringMatch("IDID")) // [0,4,1,3,2]
	fmt.Println(diStringMatch("III"))  // [0,1,2,3]
	fmt.Println(diStringMatch("DDI"))  // [3,2,0,1]
}

// diStringMatch returns a permutation that matches the DI pattern.
// Time: O(n). Space: O(n).
func diStringMatch(s string) []int {
	n := len(s)
	result := make([]int, n+1)
	low, high := 0, n
	for i, c := range s {
		if c == 'I' {
			result[i] = low
			low++
		} else {
			result[i] = high
			high--
		}
	}
	result[n] = low // or high (they are equal)
	return result
}
```

## 0944 — Delete Columns To Make Sorted

```go
package main

// LeetCode #944: Delete Columns to Make Sorted
// https://leetcode.com/problems/delete-columns-to-make-sorted/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(minDeletionSize([]string{"cba", "daf", "ghi"})) // 1
	fmt.Println(minDeletionSize([]string{"a", "b"}))             // 0
	fmt.Println(minDeletionSize([]string{"zyx", "wvu", "tsr"}))  // 3
}

// minDeletionSize counts columns to delete so the remaining columns are sorted.
// Time: O(n * m). Space: O(1).
func minDeletionSize(strs []string) int {
	if len(strs) == 0 {
		return 0
	}
	n, m := len(strs), len(strs[0])
	count := 0
	for col := 0; col < m; col++ {
		for row := 1; row < n; row++ {
			if strs[row][col] < strs[row-1][col] {
				count++
				break
			}
		}
	}
	return count
}
```

## 0953 — Verifying An Alien Dictionary

```go
package main

// LeetCode #953: Verifying an Alien Dictionary
// https://leetcode.com/problems/verifying-an-alien-dictionary/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(isAlienSorted([]string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz")) // true
	fmt.Println(isAlienSorted([]string{"word", "world", "row"}, "worldabcefghijkmnpqstuvxyz")) // false
	fmt.Println(isAlienSorted([]string{"apple", "app"}, "abcdefghijklmnopqrstuvwxyz")) // false
}

// isAlienSorted checks if words are sorted in the alien language order.
// Time: O(n * m). Space: O(1).
func isAlienSorted(words []string, order string) bool {
	orderMap := make(map[byte]int)
	for i := 0; i < len(order); i++ {
		orderMap[order[i]] = i
	}
	for i := 1; i < len(words); i++ {
		if !isLess(words[i-1], words[i], orderMap) {
			return false
		}
	}
	return true
}

func isLess(a, b string, order map[byte]int) bool {
	minLen := len(a)
	if len(b) < minLen {
		minLen = len(b)
	}
	for i := 0; i < minLen; i++ {
		if order[a[i]] < order[b[i]] {
			return true
		}
		if order[a[i]] > order[b[i]] {
			return false
		}
	}
	return len(a) <= len(b)
}
```

## 0961 — N Repeated Element In Size 2n Array

```go
package main

// LeetCode #961: N-Repeated Element in Size 2N Array
// https://leetcode.com/problems/n-repeated-element-in-size-2n-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(repeatedNTimes([]int{1, 2, 3, 3})) // 3
	fmt.Println(repeatedNTimes([]int{2, 1, 2, 5, 3, 2})) // 2
	fmt.Println(repeatedNTimes([]int{5, 1, 5, 2, 5, 3, 5, 4})) // 5
}

// repeatedNTimes finds the element repeated n times in a 2n size array.
// Time: O(n). Space: O(1).
func repeatedNTimes(nums []int) int {
	// Since the element appears n times in 2n, any two consecutive elements
	// must contain the repeated element (in most cases).
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] == nums[i+1] || nums[i] == nums[i+2] {
			return nums[i]
		}
	}
	// If not found yet, the repeated element is in the last 3 positions
	return nums[len(nums)-1]
}
```

## 0965 — Univalued Binary Tree

```go
package main

// LeetCode #965: Univalued Binary Tree
// https://leetcode.com/problems/univalued-binary-tree/
// Difficulty: Easy

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 1, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 1}},
		Right: &TreeNode{Val: 1, Right: &TreeNode{Val: 1}},
	}
	fmt.Println(isUnivalTree(root)) // true

	root2 := &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 2},
	}
	fmt.Println(isUnivalTree(root2)) // false
}

// isUnivalTree checks if all nodes in the tree have the same value.
// Time: O(n). Space: O(n).
func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return dfsUni(root, root.Val)
}

func dfsUni(node *TreeNode, val int) bool {
	if node == nil {
		return true
	}
	if node.Val != val {
		return false
	}
	return dfsUni(node.Left, val) && dfsUni(node.Right, val)
}
```

## 0976 — Largest Perimeter Triangle

```go
package main

// LeetCode #976: Largest Perimeter Triangle
// https://leetcode.com/problems/largest-perimeter-triangle/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(largestPerimeter([]int{2, 1, 2}))          // 5
	fmt.Println(largestPerimeter([]int{1, 2, 1}))          // 0
	fmt.Println(largestPerimeter([]int{3, 6, 2, 3}))       // 8
}

// largestPerimeter finds the largest perimeter of a triangle from the given side lengths.
// Time: O(n log n). Space: O(log n).
func largestPerimeter(nums []int) int {
	sort.Ints(nums)
	for i := len(nums) - 1; i >= 2; i-- {
		if nums[i-2]+nums[i-1] > nums[i] {
			return nums[i-2] + nums[i-1] + nums[i]
		}
	}
	return 0
}
```

## 0977 — Squares Of A Sorted Array

```go
package main

// LeetCode #977: Squares of a Sorted Array
// https://leetcode.com/problems/squares-of-a-sorted-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10})) // [0,1,9,16,100]
	fmt.Println(sortedSquares([]int{-7, -3, 2, 3, 11})) // [4,9,9,49,121]
}

// sortedSquares returns squares of each number sorted in non-decreasing order.
// Time: O(n). Space: O(n).
func sortedSquares(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	l, r := 0, n-1
	pos := n - 1
	for l <= r {
		leftSq := nums[l] * nums[l]
		rightSq := nums[r] * nums[r]
		if leftSq > rightSq {
			result[pos] = leftSq
			l++
		} else {
			result[pos] = rightSq
			r--
		}
		pos--
	}
	return result
}
```

## 0989 — Add To Array Form Of Integer

```go
package main

// LeetCode #989: Add to Array-Form of Integer
// https://leetcode.com/problems/add-to-array-form-of-integer/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(addToArrayForm([]int{1, 2, 0, 0}, 34)) // [1,2,3,4]
	fmt.Println(addToArrayForm([]int{2, 7, 4}, 181))   // [4,5,5]
	fmt.Println(addToArrayForm([]int{2, 1, 5}, 806))   // [1,0,2,1]
}

// addToArrayForm adds an integer to the array-form of a number.
// Time: O(max(n, log k)). Space: O(max(n, log k)).
func addToArrayForm(num []int, k int) []int {
	i := len(num) - 1
	result := make([]int, 0)
	carry := 0
	for i >= 0 || k > 0 || carry > 0 {
		digit := carry
		if i >= 0 {
			digit += num[i]
			i--
		}
		if k > 0 {
			digit += k % 10
			k /= 10
		}
		result = append(result, digit%10)
		carry = digit / 10
	}
	// Reverse
	l, r := 0, len(result)-1
	for l < r {
		result[l], result[r] = result[r], result[l]
		l++
		r--
	}
	return result
}
```

## 0993 — Cousins In Binary Tree

```go
package main

// LeetCode #993: Cousins in Binary Tree
// https://leetcode.com/problems/cousins-in-binary-tree/
// Difficulty: Easy

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
			Left:  &TreeNode{Val: 4},
		},
		Right: &TreeNode{Val: 3},
	}
	fmt.Println(isCousins(root, 4, 3)) // false

	root2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:   3,
			Right: &TreeNode{Val: 5},
		},
	}
	fmt.Println(isCousins(root2, 5, 4)) // true
}

// isCousins checks if two nodes are cousins (same depth, different parent).
// Time: O(n). Space: O(n).
func isCousins(root *TreeNode, x int, y int) bool {
	var xDepth, yDepth int
	var xParent, yParent *TreeNode
	dfsCousins(root, nil, 0, x, y, &xDepth, &yDepth, &xParent, &yParent)
	return xDepth == yDepth && xParent != nil && yParent != nil && xParent != yParent
}

func dfsCousins(node, parent *TreeNode, depth int, x, y int, xDepth, yDepth *int, xParent, yParent **TreeNode) {
	if node == nil {
		return
	}
	if node.Val == x {
		*xDepth = depth
		*xParent = parent
	}
	if node.Val == y {
		*yDepth = depth
		*yParent = parent
	}
	dfsCousins(node.Left, node, depth+1, x, y, xDepth, yDepth, xParent, yParent)
	dfsCousins(node.Right, node, depth+1, x, y, xDepth, yDepth, xParent, yParent)
}
```

## 0997 — Find The Town Judge

```go
package main

// LeetCode #997: Find the Town Judge
// https://leetcode.com/problems/find-the-town-judge/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(findJudge(2, [][]int{{1, 2}}))             // 2
	fmt.Println(findJudge(3, [][]int{{1, 3}, {2, 3}}))    // 3
	fmt.Println(findJudge(3, [][]int{{1, 3}, {2, 3}, {3, 1}})) // -1
	fmt.Println(findJudge(3, [][]int{{1, 2}, {2, 3}}))    // -1
}

// findJudge finds the town judge (trusted by everyone, trusts no one).
// Time: O(E). Space: O(V).
func findJudge(n int, trust [][]int) int {
	inDeg := make([]int, n+1)
	outDeg := make([]int, n+1)
	for _, t := range trust {
		outDeg[t[0]]++
		inDeg[t[1]]++
	}
	for i := 1; i <= n; i++ {
		if inDeg[i] == n-1 && outDeg[i] == 0 {
			return i
		}
	}
	return -1
}
```

## 0999 — Available Captures For Rook

```go
package main

// LeetCode #999: Available Captures for Rook
// https://leetcode.com/problems/available-captures-for-rook/
// Difficulty: Easy

import "fmt"

func main() {
	board := [][]byte{
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', 'p', '.', '.', '.', '.'},
		{'.', '.', '.', 'R', '.', '.', '.', 'p'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', 'p', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
	}
	fmt.Println(numRookCaptures(board)) // 3

	board2 := [][]byte{
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', 'p', 'p', 'p', 'p', 'p', '.', '.'},
		{'.', 'p', 'p', 'B', 'p', 'p', '.', '.'},
		{'.', 'p', 'B', 'R', 'B', 'p', '.', '.'},
		{'.', 'p', 'p', 'B', 'p', 'p', '.', '.'},
		{'.', 'p', 'p', 'p', 'p', 'p', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
	}
	fmt.Println(numRookCaptures(board2)) // 0
}

// numRookCaptures counts how many pawns the rook can capture.
// Time: O(1) (fixed 8x8 board). Space: O(1).
func numRookCaptures(board [][]byte) int {
	rRow, rCol := -1, -1
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board[i][j] == 'R' {
				rRow, rCol = i, j
				break
			}
		}
		if rRow != -1 {
			break
		}
	}

	dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	count := 0
	for _, d := range dirs {
		r, c := rRow+d[0], rCol+d[1]
		for r >= 0 && r < 8 && c >= 0 && c < 8 {
			if board[r][c] == 'B' {
				break
			}
			if board[r][c] == 'p' {
				count++
				break
			}
			r += d[0]
			c += d[1]
		}
	}
	return count
}
```

## 1002 — Find Common Characters

```go
package main

// LeetCode #1002: Find Common Characters
// https://leetcode.com/problems/find-common-characters/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(commonChars([]string{"bella", "label", "roller"})) // [e l l]
	fmt.Println(commonChars([]string{"cool", "lock", "cook"}))     // [c o]
}

// commonChars returns the common characters (with multiplicity) across all words.
// Time: O(n * m) where n = len(words), m = avg word length. Space: O(1).
func commonChars(words []string) []string {
	if len(words) == 0 {
		return []string{}
	}

	// Initialize with counts from first word
	freq := [26]int{}
	for _, c := range words[0] {
		freq[c-'a']++
	}

	for i := 1; i < len(words); i++ {
		currFreq := [26]int{}
		for _, c := range words[i] {
			currFreq[c-'a']++
		}
		for i := 0; i < 26; i++ {
			if currFreq[i] < freq[i] {
				freq[i] = currFreq[i]
			}
		}
	}

	result := make([]string, 0)
	for i := 0; i < 26; i++ {
		for j := 0; j < freq[i]; j++ {
			result = append(result, string(rune('a'+i)))
		}
	}
	return result
}
```

## 1005 — Maximize Sum Of Array After K Negations

```go
package main

// LeetCode #1005: Maximize Sum Of Array After K Negations
// https://leetcode.com/problems/maximize-sum-of-array-after-k-negations/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(largestSumAfterKNegations([]int{4, 2, 3}, 1))      // 5
	fmt.Println(largestSumAfterKNegations([]int{3, -1, 0, 2}, 3))  // 6
	fmt.Println(largestSumAfterKNegations([]int{2, -3, -1, 5, -4}, 2)) // 13
}

// largestSumAfterKNegations maximizes the sum by negating k elements.
// Time: O(n log n). Space: O(1).
func largestSumAfterKNegations(nums []int, k int) int {
	sort.Ints(nums)
	// Flip negatives
	for i := 0; i < len(nums) && nums[i] < 0 && k > 0; i++ {
		nums[i] = -nums[i]
		k--
	}
	// If k is odd, flip the smallest absolute value
	if k%2 == 1 {
		minIdx := 0
		for i := 1; i < len(nums); i++ {
			if nums[i] < nums[minIdx] {
				minIdx = i
			}
		}
		nums[minIdx] = -nums[minIdx]
	}
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}
```

## 1009 — Complement Of Base 10 Integer

```go
package main

// LeetCode #1009: Complement of Base 10 Integer
// https://leetcode.com/problems/complement-of-base-10-integer/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(bitwiseComplement(5))  // 2
	fmt.Println(bitwiseComplement(7))  // 0
	fmt.Println(bitwiseComplement(10)) // 5
	fmt.Println(bitwiseComplement(0))  // 1
}

// bitwiseComplement returns the complement of a base-10 integer's binary representation.
// Time: O(log n). Space: O(1).
func bitwiseComplement(n int) int {
	if n == 0 {
		return 1
	}
	mask := n
	// Set all bits to the right of the MSB
	mask |= mask >> 1
	mask |= mask >> 2
	mask |= mask >> 4
	mask |= mask >> 8
	mask |= mask >> 16
	return ^n & mask
}
```

## 1013 — Partition Array Into Three Parts With Equal Sum

```go
package main

// LeetCode #1013: Partition Array Into Three Parts With Equal Sum
// https://leetcode.com/problems/partition-array-into-three-parts-with-equal-sum/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(partitionArrayIntoThreePartsWithEqualSum([]int{0, 2, 1, -6, 6, -7, 9, 1, 2, 0, 1})) // true
	fmt.Println(partitionArrayIntoThreePartsWithEqualSum([]int{0, 2, 1, -6, 6, 7, 9, -1, 2, 0, 1})) // false
	fmt.Println(partitionArrayIntoThreePartsWithEqualSum([]int{3, 3, 6, 5, -2, 2, 5, 1, -9, 4}))    // true
}

// LeetCode submission: canThreePartsEqualSum
func canThreePartsEqualSum(arr []int) bool {
	total := 0
	for _, v := range arr {
		total += v
	}
	if total%3 != 0 {
		return false
	}
	target := total / 3
	sum, count := 0, 0
	for _, v := range arr {
		sum += v
		if sum == target {
			count++
			sum = 0
		}
	}
	return count >= 3
}

func partitionArrayIntoThreePartsWithEqualSum(arr []int) bool {
	return canThreePartsEqualSum(arr)
}
```

## 1018 — Binary Prefix Divisible By 5

```go
package main

// LeetCode #1018: Binary Prefix Divisible By 5
// https://leetcode.com/problems/binary-prefix-divisible-by-5/
// Difficulty: Easy
// Time: O(n) | Space: O(1) (excluding output)

import "fmt"

func main() {
	fmt.Println(prefixesDivBy5([]int{0, 1, 1}))          // [true,false,false]
	fmt.Println(prefixesDivBy5([]int{1, 1, 1}))          // [false,false,false]
	fmt.Println(prefixesDivBy5([]int{0, 1, 1, 1, 1, 1})) // [true,false,false,false,true,false]
}

// LeetCode submission: prefixesDivBy5
func prefixesDivBy5(nums []int) []bool {
	ans := make([]bool, len(nums))
	val := 0
	for i, b := range nums {
		val = (val*2 + b) % 5
		ans[i] = val == 0
	}
	return ans
}
```

## 1021 — Remove Outermost Parentheses

```go
package main

// LeetCode #1021: Remove Outermost Parentheses
// https://leetcode.com/problems/remove-outermost-parentheses/
// Difficulty: Easy
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(removeOuterParentheses("(()())(())"))          // "()()()"
	fmt.Println(removeOuterParentheses("(()())(())(()(()))"))  // "()()()()(())"
	fmt.Println(removeOuterParentheses("()()"))                // ""
}

// LeetCode submission: removeOuterParentheses
func removeOuterParentheses(s string) string {
	ans := make([]byte, 0, len(s))
	depth := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			if depth > 0 {
				ans = append(ans, '(')
			}
			depth++
		} else {
			depth--
			if depth > 0 {
				ans = append(ans, ')')
			}
		}
	}
	return string(ans)
}
```

## 1022 — Sum Of Root To Leaf Binary Numbers

```go
package main

// LeetCode #1022: Sum of Root To Leaf Binary Numbers
// https://leetcode.com/problems/sum-of-root-to-leaf-binary-numbers/
// Difficulty: Easy
// Time: O(n) | Space: O(h) where h is tree height

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	//     1
	//    / \
	//   0   1
	//  / \ / \
	// 0  1 0  1
	root := &TreeNode{1,
		&TreeNode{0, &TreeNode{0, nil, nil}, &TreeNode{1, nil, nil}},
		&TreeNode{1, &TreeNode{0, nil, nil}, &TreeNode{1, nil, nil}},
	}
	fmt.Println(sumRootToLeaf(root)) // 22

	// Single node
	fmt.Println(sumRootToLeaf(&TreeNode{1, nil, nil})) // 1
}

// LeetCode submission: sumRootToLeaf
func sumRootToLeaf(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(node *TreeNode, val int) int {
	if node == nil {
		return 0
	}
	val = val*2 + node.Val
	if node.Left == nil && node.Right == nil {
		return val
	}
	return dfs(node.Left, val) + dfs(node.Right, val)
}
```

## 1025 — Divisor Game

```go
package main

// LeetCode #1025: Divisor Game
// https://leetcode.com/problems/divisor-game/
// Difficulty: Easy
// Time: O(1) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(divisorGame(2)) // true
	fmt.Println(divisorGame(3)) // false
	fmt.Println(divisorGame(4)) // true
}

// LeetCode submission: divisorGame
func divisorGame(n int) bool {
	return n%2 == 0
}
```

## 1030 — Matrix Cells In Distance Order

```go
package main

// LeetCode #1030: Matrix Cells in Distance Order
// https://leetcode.com/problems/matrix-cells-in-distance-order/
// Difficulty: Easy
// Time: O(R*C) | Space: O(R*C)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(allCellsDistOrder(1, 2, 0, 0)) // [[0,0],[0,1]]
	fmt.Println(allCellsDistOrder(2, 2, 0, 1)) // [[0,1],[0,0],[1,1],[1,0]]
	fmt.Println(allCellsDistOrder(2, 3, 1, 2)) // [[1,2],[0,2],[1,1],[0,1],[1,0],[0,0]]
}

// LeetCode submission: allCellsDistOrder
func allCellsDistOrder(rows, cols, rCenter, cCenter int) [][]int {
	ans := make([][]int, 0, rows*cols)
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			ans = append(ans, []int{r, c})
		}
	}
	sort.Slice(ans, func(i, j int) bool {
		di := abs(ans[i][0]-rCenter) + abs(ans[i][1]-cCenter)
		dj := abs(ans[j][0]-rCenter) + abs(ans[j][1]-cCenter)
		return di < dj
	})
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

## 1037 — Valid Boomerang

```go
package main

// LeetCode #1037: Valid Boomerang
// https://leetcode.com/problems/valid-boomerang/
// Difficulty: Easy
// Time: O(1) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(isBoomerang([][]int{{1, 1}, {2, 3}, {3, 2}})) // true
	fmt.Println(isBoomerang([][]int{{1, 1}, {2, 2}, {3, 3}})) // false
	fmt.Println(isBoomerang([][]int{{0, 0}, {0, 2}, {0, 1}})) // false
}

// LeetCode submission: isBoomerang
func isBoomerang(points [][]int) bool {
	x1, y1 := points[0][0], points[0][1]
	x2, y2 := points[1][0], points[1][1]
	x3, y3 := points[2][0], points[2][1]
	// Check area of triangle: (x2-x1)*(y3-y1) != (x3-x1)*(y2-y1)
	return (x2-x1)*(y3-y1) != (x3-x1)*(y2-y1)
}
```

## 1046 — Last Stone Weight

```go
package main

// LeetCode #1046: Last Stone Weight
// https://leetcode.com/problems/last-stone-weight/
// Difficulty: Easy
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(lastStoneWeight([]int{2, 7, 4, 1, 8, 1})) // 1
	fmt.Println(lastStoneWeight([]int{1}))                 // 1
	fmt.Println(lastStoneWeight([]int{2, 2}))              // 0
}

// LeetCode submission: lastStoneWeight
func lastStoneWeight(stones []int) int {
	for len(stones) > 1 {
		sort.Ints(stones)
		n := len(stones)
		if stones[n-1] == stones[n-2] {
			stones = stones[:n-2]
		} else {
			stones[n-2] = stones[n-1] - stones[n-2]
			stones = stones[:n-1]
		}
	}
	if len(stones) == 1 {
		return stones[0]
	}
	return 0
}
```

## 1047 — Remove All Adjacent Duplicates In String

```go
package main

// LeetCode #1047: Remove All Adjacent Duplicates In String
// https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/
// Difficulty: Easy
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(removeDuplicates("abbaca")) // "ca"
	fmt.Println(removeDuplicates("azxxzy")) // "ay"
	fmt.Println(removeDuplicates("a"))      // "a"
}

// LeetCode submission: removeDuplicates
func removeDuplicates(s string) string {
	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 && stack[len(stack)-1] == s[i] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}
```

## 1050 — Actors And Directors Who Cooperated At Least Three Times

```go
package main

// LeetCode #1050: Actors and Directors Who Cooperated At Least Three Times
// https://leetcode.com/problems/actors-and-directors-who-cooperated-at-least-three-times/
// Difficulty: Easy (SQL)
// Time: O(n) | Space: O(n)
// Note: This is a SQL problem. Go implementation simulates the query logic.

import "fmt"

type actorDirector struct {
	actorID   int
	directorID int
	timestamp int
}

func main() {
	pairs := []actorDirector{
		{1, 1, 0}, {1, 1, 1}, {1, 1, 2},
		{1, 2, 3}, {2, 1, 4}, {2, 1, 5},
	}
	fmt.Println(actorsAndDirectorsWhoCooperatedAtLeastThreeTimes(pairs)) // [[1 1]]

	pairs2 := []actorDirector{{1, 1, 0}, {1, 1, 1}}
	fmt.Println(actorsAndDirectorsWhoCooperatedAtLeastThreeTimes(pairs2)) // []
}

// LeetCode submission: actorsAndDirectorsWhoCooperatedAtLeastThreeTimes (SQL equivalent)
func actorsAndDirectorsWhoCooperatedAtLeastThreeTimes(pairs []actorDirector) [][]int {
	count := make(map[[2]int]int)
	for _, p := range pairs {
		key := [2]int{p.actorID, p.directorID}
		count[key]++
	}
	var ans [][]int
	for k, v := range count {
		if v >= 3 {
			ans = append(ans, []int{k[0], k[1]})
		}
	}
	return ans
}
```

## 1051 — Height Checker

```go
package main

// LeetCode #1051: Height Checker
// https://leetcode.com/problems/height-checker/
// Difficulty: Easy
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(heightChecker([]int{1, 1, 4, 2, 1, 3})) // 3
	fmt.Println(heightChecker([]int{5, 1, 2, 3, 4}))    // 5
	fmt.Println(heightChecker([]int{1, 2, 3, 4, 5}))    // 0
}

// LeetCode submission: heightChecker
func heightChecker(heights []int) int {
	expected := make([]int, len(heights))
	copy(expected, heights)
	sort.Ints(expected)
	count := 0
	for i := range heights {
		if heights[i] != expected[i] {
			count++
		}
	}
	return count
}
```

## 1056 — Confusing Number

```go
package main

// LeetCode #1056: Confusing Number
// https://leetcode.com/problems/confusing-number/
// Difficulty: Easy [Paid]
// Time: O(log n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(confusingNumber(6))   // true (6 -> 9)
	fmt.Println(confusingNumber(89))  // true (89 -> 68)
	fmt.Println(confusingNumber(11))  // false (11 -> 11, same)
	fmt.Println(confusingNumber(25))  // false (invalid digit)
}

// LeetCode submission: confusingNumber
func confusingNumber(n int) bool {
	d := []int{0, 1, -1, -1, -1, -1, 9, -1, 8, 6}
	x, y := n, 0
	for x > 0 {
		v := x % 10
		if d[v] < 0 {
			return false
		}
		y = y*10 + d[v]
		x /= 10
	}
	return y != n
}
```

## 1064 — Fixed Point

```go
package main

// LeetCode #1064: Fixed Point
// https://leetcode.com/problems/fixed-point/
// Difficulty: Easy [Paid]
// Time: O(log n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(fixedPoint([]int{-10, -5, 0, 3, 7}))  // 3
	fmt.Println(fixedPoint([]int{0, 2, 5, 8, 17}))    // 0
	fmt.Println(fixedPoint([]int{-10, -5, 3, 4, 7, 9})) // -1
}

// LeetCode submission: fixedPoint
func fixedPoint(arr []int) int {
	lo, hi := 0, len(arr)-1
	for lo < hi {
		mid := (lo + hi) >> 1
		if arr[mid] >= mid {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	if arr[lo] == lo {
		return lo
	}
	return -1
}
```

## 1065 — Index Pairs Of A String

```go
package main

// LeetCode #1065: Index Pairs of a String
// https://leetcode.com/problems/index-pairs-of-a-string/
// Difficulty: Easy [Paid]
// Time: O(n^2 + m*k) | Space: O(m*k) for trie

import "fmt"

func main() {
	fmt.Println(indexPairs("thestoryofleetcodeandme", []string{"story", "fleet", "leetcode"}))
	// [[1,5],[3,7],[10,13],[10,18]]
	fmt.Println(indexPairs("ababa", []string{"aba", "ab"}))
	// [[0,1],[0,2],[2,3],[2,4]]
}

// LeetCode submission: indexPairs
func indexPairs(text string, words []string) [][]int {
	wordSet := make(map[string]bool)
	for _, w := range words {
		wordSet[w] = true
	}
	var ans [][]int
	for i := 0; i < len(text); i++ {
		for j := i; j < len(text); j++ {
			if wordSet[text[i:j+1]] {
				ans = append(ans, []int{i, j})
			}
		}
	}
	return ans
}
```

## 1068 — Product Sales Analysis I

```go
package main

// LeetCode #1068: Product Sales Analysis I
// https://leetcode.com/problems/product-sales-analysis-i/
// Difficulty: Easy (SQL)
// Time: O(n) | Space: O(n)
// Note: This is a SQL problem. Go implementation simulates the query logic.

import "fmt"

type Sale struct {
	saleID    int
	productID int
	year      int
	quantity  int
	price     int
}

type Product struct {
	productID   int
	productName string
}

func main() {
	sales := []Sale{
		{1, 100, 2008, 10, 5000},
		{2, 100, 2009, 12, 5000},
		{7, 200, 2011, 15, 9000},
	}
	products := []Product{
		{100, "Nokia"},
		{200, "Apple"},
		{300, "Samsung"},
	}
	// SQL equivalent: SELECT p.product_name, s.year, s.price
	// FROM Sales s JOIN Product p ON s.product_id = p.product_id
	fmt.Println(productSalesAnalysisI(sales, products))
}

// LeetCode submission: productSalesAnalysisI (SQL equivalent)
func productSalesAnalysisI(sales []Sale, products []Product) []map[string]int {
	prodMap := make(map[int]string)
	for _, p := range products {
		prodMap[p.productID] = p.productName
	}
	type result struct {
		name  string
		year  int
		price int
	}
	var ans []result
	for _, s := range sales {
		if name, ok := prodMap[s.productID]; ok {
			ans = append(ans, result{name, s.year, s.price})
		}
	}
	return nil // placeholder - SQL query is the real answer
}
```

## 1069 — Product Sales Analysis Ii

```go
package main

// LeetCode #1069: Product Sales Analysis II
// https://leetcode.com/problems/product-sales-analysis-ii/
// Difficulty: Easy [Paid] (SQL)
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println("SQL: SELECT product_id, SUM(quantity) AS total_quantity FROM Sales GROUP BY product_id")
}

// This is a SQL problem. The answer is the SQL query above.
```

## 1071 — Greatest Common Divisor Of Strings

```go
package main

// LeetCode #1071: Greatest Common Divisor of Strings
// https://leetcode.com/problems/greatest-common-divisor-of-strings/
// Difficulty: Easy
// Time: O(n + m) | Space: O(n + m)

import "fmt"

func main() {
	fmt.Println(gcdOfStrings("ABCABC", "ABC")) // "ABC"
	fmt.Println(gcdOfStrings("ABABAB", "ABAB")) // "AB"
	fmt.Println(gcdOfStrings("LEET", "CODE"))   // ""
}

// LeetCode submission: gcdOfStrings
func gcdOfStrings(str1, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}
	return str1[:gcd(len(str1), len(str2))]
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
```

## 1075 — Project Employees I

```go
package main

// LeetCode #1075: Project Employees I
// https://leetcode.com/problems/project-employees-i/
// Difficulty: Easy (SQL)
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println("SQL: SELECT p.project_id, ROUND(AVG(e.experience_years), 2) AS average_years FROM Project p JOIN Employee e ON p.employee_id = e.employee_id GROUP BY p.project_id")
}

// This is a SQL problem. The answer is the SQL query above.
```

## 1076 — Project Employees Ii

```go
package main

// LeetCode #1076: Project Employees II
// https://leetcode.com/problems/project-employees-ii/
// Difficulty: Easy [Paid] (SQL)
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println("SQL: SELECT project_id FROM Project GROUP BY project_id HAVING COUNT(employee_id) = (SELECT COUNT(employee_id) FROM Project GROUP BY project_id ORDER BY COUNT(employee_id) DESC LIMIT 1)")
}

// This is a SQL problem. The answer is the SQL query above.
```

## 1078 — Occurrences After Bigram

```go
package main

// LeetCode #1078: Occurrences After Bigram
// https://leetcode.com/problems/occurrences-after-bigram/
// Difficulty: Easy
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(findOcurrences("alice is a good girl she is a good student", "a", "good"))
	// ["girl","student"]
	fmt.Println(findOcurrences("we will we will rock you", "we", "will"))
	// ["we","rock"]
}

// LeetCode submission: findOcurrences
func findOcurrences(text, first, second string) []string {
	words := strings.Fields(text)
	var ans []string
	for i := 2; i < len(words); i++ {
		if words[i-2] == first && words[i-1] == second {
			ans = append(ans, words[i])
		}
	}
	return ans
}
```

## 1082 — Sales Analysis I

```go
package main

// LeetCode #1082: Sales Analysis I
// https://leetcode.com/problems/sales-analysis-i/
// Difficulty: Easy [Paid] (SQL)
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println("SQL: SELECT seller_id FROM Sales GROUP BY seller_id HAVING SUM(price) = (SELECT SUM(price) FROM Sales GROUP BY seller_id ORDER BY SUM(price) DESC LIMIT 1)")
}

// This is a SQL problem. The answer is the SQL query above.
```

## 1083 — Sales Analysis Ii

```go
package main

// LeetCode #1083: Sales Analysis II
// https://leetcode.com/problems/sales-analysis-ii/
// Difficulty: Easy [Paid] (SQL)
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println("SQL: SELECT DISTINCT s.buyer_id FROM Sales s JOIN Product p ON s.product_id = p.product_id WHERE p.product_name = 'S8' AND s.buyer_id NOT IN (SELECT s2.buyer_id FROM Sales s2 JOIN Product p2 ON s2.product_id = p2.product_id WHERE p2.product_name = 'iPhone')")
}

// This is a SQL problem. The answer is the SQL query above.
```

## 1084 — Sales Analysis Iii

```go
package main

// LeetCode #1084: Sales Analysis III
// https://leetcode.com/problems/sales-analysis-iii/
// Difficulty: Easy (SQL)
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println("SQL: SELECT p.product_id, p.product_name FROM Product p JOIN Sales s ON p.product_id = s.product_id GROUP BY p.product_id, p.product_name HAVING MIN(s.sale_date) >= '2019-01-01' AND MAX(s.sale_date) <= '2019-03-31'")
}

// This is a SQL problem. The answer is the SQL query above.
```

## 1085 — Sum Of Digits In The Minimum Number

```go
package main

// LeetCode #1085: Sum of Digits in the Minimum Number
// https://leetcode.com/problems/sum-of-digits-in-the-minimum-number/
// Difficulty: Easy [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(sumOfDigits([]int{34, 23, 1, 24, 75, 33, 54, 8})) // 0 (min=1, sum=1, odd)
	fmt.Println(sumOfDigits([]int{99, 77, 33, 66, 55}))           // 1 (min=33, sum=6, even)
}

// LeetCode submission: sumOfDigits
func sumOfDigits(nums []int) int {
	x := nums[0]
	for _, v := range nums {
		if v < x {
			x = v
		}
	}
	s := 0
	for x > 0 {
		s += x % 10
		x /= 10
	}
	if s%2 == 0 {
		return 1
	}
	return 0
}
```

## 1086 — High Five

```go
package main

// LeetCode #1086: High Five
// https://leetcode.com/problems/high-five/
// Difficulty: Easy [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func main() {
	items := [][]int{
		{1, 91}, {1, 92}, {2, 93}, {2, 97}, {1, 60},
		{2, 77}, {1, 65}, {1, 87}, {1, 100}, {2, 100},
		{2, 76},
	}
	fmt.Println(highFive(items)) // [[1,87],[2,88]]
}

// LeetCode submission: highFive
func highFive(items [][]int) [][]int {
	scores := make(map[int][]int)
	for _, item := range items {
		id, score := item[0], item[1]
		scores[id] = append(scores[id], score)
	}
	var ids []int
	for id := range scores {
		ids = append(ids, id)
	}
	sort.Ints(ids)
	var ans [][]int
	for _, id := range ids {
		s := scores[id]
		sort.Sort(sort.Reverse(sort.IntSlice(s)))
		sum := 0
		for i := 0; i < 5; i++ {
			sum += s[i]
		}
		ans = append(ans, []int{id, sum / 5})
	}
	return ans
}
```

## 1089 — Duplicate Zeros

```go
package main

// LeetCode #1089: Duplicate Zeros
// https://leetcode.com/problems/duplicate-zeros/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	arr1 := []int{1, 0, 2, 3, 0, 4, 5, 0}
	duplicateZeros(arr1)
	fmt.Println(arr1) // [1,0,0,2,3,0,0,4]

	arr2 := []int{1, 2, 3}
	duplicateZeros(arr2)
	fmt.Println(arr2) // [1,2,3]
}

// LeetCode submission: duplicateZeros
func duplicateZeros(arr []int) {
	n := len(arr)
	possibleDups := 0
	for i := 0; i+possibleDups < n; i++ {
		if arr[i] == 0 {
			possibleDups++
		}
	}
	last := n - 1 - possibleDups
	for i := last; i >= 0; i-- {
		if i+possibleDups < n {
			arr[i+possibleDups] = arr[i]
		}
		if arr[i] == 0 {
			possibleDups--
			if i+possibleDups < n {
				arr[i+possibleDups] = 0
			}
		}
	}
}
```

## 1099 — Two Sum Less Than K

```go
package main

// LeetCode #1099: Two Sum Less Than K
// https://leetcode.com/problems/two-sum-less-than-k/
// Difficulty: Easy [Paid]
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(twoSumLessThanK([]int{34, 23, 1, 24, 75, 33, 54, 8}, 60)) // 58
	fmt.Println(twoSumLessThanK([]int{10, 20, 30}, 15))                   // -1
}

// LeetCode submission: twoSumLessThanK
func twoSumLessThanK(nums []int, k int) int {
	sort.Ints(nums)
	ans := -1
	i, j := 0, len(nums)-1
	for i < j {
		sum := nums[i] + nums[j]
		if sum < k {
			if sum > ans {
				ans = sum
			}
			i++
		} else {
			j--
		}
	}
	return ans
}
```

## 1103 — Distribute Candies To People

```go
package main

// LeetCode #1103: Distribute Candies to People
// https://leetcode.com/problems/distribute-candies-to-people/
// Difficulty: Easy
// Time: O(sqrt(candies)) | Space: O(numPeople)

import "fmt"

func main() {
	fmt.Println(distributeCandies(7, 4))  // [1,2,3,1]
	fmt.Println(distributeCandies(10, 3)) // [5,2,3]
}

// LeetCode submission: distributeCandies
func distributeCandies(candies int, numPeople int) []int {
	ans := make([]int, numPeople)
	give := 1
	for candies > 0 {
		for i := 0; i < numPeople && candies > 0; i++ {
			if give <= candies {
				ans[i] += give
				candies -= give
			} else {
				ans[i] += candies
				candies = 0
			}
			give++
		}
	}
	return ans
}
```

## 1108 — Defanging An Ip Address

```go
package main

// LeetCode #1108: Defanging an IP Address
// https://leetcode.com/problems/defanging-an-ip-address/
// Difficulty: Easy
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(defangIPaddr("1.1.1.1"))       // "1[.]1[.]1[.]1"
	fmt.Println(defangIPaddr("255.100.50.0"))  // "255[.]100[.]50[.]0"
}

// LeetCode submission: defangIPaddr
func defangIPaddr(address string) string {
	return strings.ReplaceAll(address, ".", "[.]")
}
```

## 1113 — Reported Posts

```go
package main

// LeetCode #1113: Reported Posts
// https://leetcode.com/problems/reported-posts/
// Difficulty: Easy [Paid] (SQL)
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println("SQL: SELECT extra AS report_reason, COUNT(DISTINCT post_id) AS report_count FROM Actions WHERE action = 'report' AND action_date = '2019-07-04' GROUP BY extra")
}

// This is a SQL problem. The answer is the SQL query above.
```

## 1114 — Print In Order

```go
package main

// LeetCode #1114: Print in Order
// https://leetcode.com/problems/print-in-order/
// Difficulty: Easy (Concurrency)
// Time: O(1) | Space: O(1)

import (
	"fmt"
	"sync"
)

type Foo struct {
	wg1 sync.WaitGroup
	wg2 sync.WaitGroup
}

func NewFoo() *Foo {
	f := &Foo{}
	f.wg1.Add(1)
	f.wg2.Add(1)
	return f
}

func (f *Foo) first() {
	fmt.Print("first")
	f.wg1.Done()
}

func (f *Foo) second() {
	f.wg1.Wait()
	fmt.Print("second")
	f.wg2.Done()
}

func (f *Foo) third() {
	f.wg2.Wait()
	fmt.Print("third")
}

func main() {
	// Test: run in order 1,2,3
	f := NewFoo()
	var wg sync.WaitGroup
	wg.Add(3)
	go func() { f.first(); wg.Done() }()
	go func() { f.second(); wg.Done() }()
	go func() { f.third(); wg.Done() }()
	wg.Wait()
	fmt.Println()
}
```

## 1118 — Number Of Days In A Month

```go
package main

// LeetCode #1118: Number of Days in a Month
// https://leetcode.com/problems/number-of-days-in-a-month/
// Difficulty: Easy [Paid]
// Time: O(1) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(numberOfDays(1992, 7))  // 31
	fmt.Println(numberOfDays(2000, 2))  // 29
	fmt.Println(numberOfDays(1900, 2))  // 28
}

// LeetCode submission: numberOfDays
func numberOfDays(year, month int) int {
	leap := (year%4 == 0 && year%100 != 0) || (year%400 == 0)
	days := []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if leap {
		days[2] = 29
	}
	return days[month]
}
```

## 1119 — Remove Vowels From A String

```go
package main

// LeetCode #1119: Remove Vowels from a String
// https://leetcode.com/problems/remove-vowels-from-a-string/
// Difficulty: Easy [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(removeVowels("leetcodeisacommunityforcoders")) // "ltcdscmmntyfrcdrs"
	fmt.Println(removeVowels("aeiou"))                         // ""
}

// LeetCode submission: removeVowels
func removeVowels(s string) string {
	ans := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c != 'a' && c != 'e' && c != 'i' && c != 'o' && c != 'u' {
			ans = append(ans, c)
		}
	}
	return string(ans)
}
```

## 1122 — Relative Sort Array

```go
package main

// LeetCode #1122: Relative Sort Array
// https://leetcode.com/problems/relative-sort-array/
// Difficulty: Easy
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(relativeSortArray([]int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}, []int{2, 1, 4, 3, 9, 6}))
	// [2,2,2,1,4,3,3,9,6,7,19]
	fmt.Println(relativeSortArray([]int{28, 6, 22, 8, 44, 17}, []int{22, 28, 8, 6}))
	// [22,28,8,6,17,44]
}

// LeetCode submission: relativeSortArray
func relativeSortArray(arr1, arr2 []int) []int {
	rank := make(map[int]int, len(arr2))
	for i, v := range arr2 {
		rank[v] = i
	}
	sort.Slice(arr1, func(i, j int) bool {
		ri, okI := rank[arr1[i]]
		rj, okJ := rank[arr1[j]]
		if okI && okJ {
			return ri < rj
		}
		if okI {
			return true
		}
		if okJ {
			return false
		}
		return arr1[i] < arr1[j]
	})
	return arr1
}
```

## 1128 — Number Of Equivalent Domino Pairs

```go
package main

// LeetCode #1128: Number of Equivalent Domino Pairs
// https://leetcode.com/problems/number-of-equivalent-domino-pairs/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(numEquivDominoPairs([][]int{{1, 2}, {2, 1}, {3, 4}, {5, 6}})) // 1
	fmt.Println(numEquivDominoPairs([][]int{{1, 2}, {1, 2}, {1, 1}, {1, 2}}))  // 3
}

// LeetCode submission: numEquivDominoPairs
func numEquivDominoPairs(dominoes [][]int) int {
	count := make(map[[2]int]int)
	ans := 0
	for _, d := range dominoes {
		a, b := d[0], d[1]
		if a > b {
			a, b = b, a
		}
		key := [2]int{a, b}
		ans += count[key]
		count[key]++
	}
	return ans
}
```

## 1133 — Largest Unique Number

```go
package main

// LeetCode #1133: Largest Unique Number
// https://leetcode.com/problems/largest-unique-number/
// Difficulty: Easy [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(largestUniqueNumber([]int{5, 7, 3, 9, 4, 9, 8, 3, 1})) // 8
	fmt.Println(largestUniqueNumber([]int{9, 9, 8, 8}))                // -1
}

// LeetCode submission: largestUniqueNumber
func largestUniqueNumber(nums []int) int {
	count := make(map[int]int)
	for _, v := range nums {
		count[v]++
	}
	ans := -1
	for k, v := range count {
		if v == 1 && k > ans {
			ans = k
		}
	}
	return ans
}
```

## 1134 — Armstrong Number

```go
package main

// LeetCode #1134: Armstrong Number
// https://leetcode.com/problems/armstrong-number/
// Difficulty: Easy [Paid]
// Time: O(log n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(isArmstrong(153))  // true
	fmt.Println(isArmstrong(123))  // false
	fmt.Println(isArmstrong(1))    // true
}

// LeetCode submission: isArmstrong
func isArmstrong(n int) bool {
	digits := 0
	for x := n; x > 0; x /= 10 {
		digits++
	}
	sum := 0
	for x := n; x > 0; x /= 10 {
		d := x % 10
		p := 1
		for i := 0; i < digits; i++ {
			p *= d
		}
		sum += p
	}
	return sum == n
}
```

## 1137 — N Th Tribonacci Number

```go
package main

// LeetCode #1137: N-th Tribonacci Number
// https://leetcode.com/problems/n-th-tribonacci-number/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(tribonacci(4))  // 4
	fmt.Println(tribonacci(25)) // 1389537
	fmt.Println(tribonacci(0))  // 0
}

// LeetCode submission: tribonacci
func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	a, b, c := 0, 1, 1
	for i := 3; i <= n; i++ {
		a, b, c = b, c, a+b+c
	}
	return c
}
```

## 1141 — User Activity For The Past 30 Days I

```go
package main

// LeetCode #1141: User Activity for the Past 30 Days I
// https://leetcode.com/problems/user-activity-for-the-past-30-days-i/
// Difficulty: Easy (SQL)
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println("SQL: SELECT activity_date AS day, COUNT(DISTINCT user_id) AS active_users FROM Activity WHERE activity_date BETWEEN '2019-06-28' AND '2019-07-27' GROUP BY activity_date")
}

// This is a SQL problem. The answer is the SQL query above.
```

## 1142 — User Activity For The Past 30 Days Ii

```go
package main

// LeetCode #1142: User Activity for the Past 30 Days II
// https://leetcode.com/problems/user-activity-for-the-past-30-days-ii/
// Difficulty: Easy [Paid] (SQL)
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println("SQL: SELECT IFNULL(ROUND(COUNT(DISTINCT session_id) / COUNT(DISTINCT user_id), 2), 0) AS average_sessions_per_user FROM Activity WHERE activity_date BETWEEN '2019-06-28' AND '2019-07-27'")
}

// This is a SQL problem. The answer is the SQL query above.
```

## 1148 — Article Views I

```go
package main

// LeetCode #1148: Article Views I
// https://leetcode.com/problems/article-views-i/
// Difficulty: Easy (SQL)
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println("SQL: SELECT DISTINCT author_id AS id FROM Views WHERE author_id = viewer_id ORDER BY author_id")
}

// This is a SQL problem. The answer is the SQL query above.
```

## 1150 — Check If A Number Is Majority Element In A Sorted Array

```go
package main

// LeetCode #1150: Check If a Number Is Majority Element in a Sorted Array
// https://leetcode.com/problems/check-if-a-number-is-majority-element-in-a-sorted-array/
// Difficulty: Easy [Paid]
// Time: O(log n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(isMajorityElement([]int{2, 4, 5, 5, 5, 5, 5, 6, 6}, 5)) // true
	fmt.Println(isMajorityElement([]int{10, 100, 101, 101}, 101))        // false
}

// LeetCode submission: isMajorityElement
func isMajorityElement(nums []int, target int) bool {
	n := len(nums)
	// Binary search for first occurrence
	lo, hi := 0, n-1
	for lo < hi {
		mid := (lo + hi) >> 1
		if nums[mid] >= target {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	first := lo
	last := first + n/2
	return last < n && nums[last] == target
}
```

## 1154 — Day Of The Year

```go
package main

// LeetCode #1154: Day of the Year
// https://leetcode.com/problems/day-of-the-year/
// Difficulty: Easy
// Time: O(1) | Space: O(1)

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(dayOfYear("2019-01-09")) // 9
	fmt.Println(dayOfYear("2019-02-10")) // 41
	fmt.Println(dayOfYear("2000-03-01")) // 61
}

// LeetCode submission: dayOfYear
func dayOfYear(date string) int {
	parts := strings.Split(date, "-")
	year, _ := strconv.Atoi(parts[0])
	month, _ := strconv.Atoi(parts[1])
	day, _ := strconv.Atoi(parts[2])

	leap := (year%4 == 0 && year%100 != 0) || (year%400 == 0)
	days := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if leap {
		days[1] = 29
	}

	ans := 0
	for i := 0; i < month-1; i++ {
		ans += days[i]
	}
	return ans + day
}
```

## 1160 — Find Words That Can Be Formed By Characters

```go
package main

// LeetCode #1160: Find Words That Can Be Formed by Characters
// https://leetcode.com/problems/find-words-that-can-be-formed-by-characters/
// Difficulty: Easy
// Time: O(n * k) where k is max word length | Space: O(1)

import "fmt"

func main() {
	fmt.Println(countCharacters([]string{"cat", "bt", "hat", "tree"}, "atach")) // 6
	fmt.Println(countCharacters([]string{"hello", "world", "leetcode"}, "welldonehoneyr")) // 10
}

// LeetCode submission: countCharacters
func countCharacters(words []string, chars string) int {
	ch := make([]int, 26)
	for i := 0; i < len(chars); i++ {
		ch[chars[i]-'a']++
	}
	ans := 0
	for _, w := range words {
		need := make([]int, 26)
		ok := true
		for i := 0; i < len(w); i++ {
			idx := w[i] - 'a'
			need[idx]++
			if need[idx] > ch[idx] {
				ok = false
				break
			}
		}
		if ok {
			ans += len(w)
		}
	}
	return ans
}
```

## 1165 — Single Row Keyboard

```go
package main

// LeetCode #1165: Single-Row Keyboard
// https://leetcode.com/problems/single-row-keyboard/
// Difficulty: Easy [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(calculateTime("abcdefghijklmnopqrstuvwxyz", "cba")) // 4
	fmt.Println(calculateTime("pqrstuvwxyzabcdefghijklmno", "leetcode")) // 73
}

// LeetCode submission: calculateTime
func calculateTime(keyboard, word string) int {
	pos := [26]int{}
	for i := 0; i < len(keyboard); i++ {
		pos[keyboard[i]-'a'] = i
	}
	ans, cur := 0, 0
	for i := 0; i < len(word); i++ {
		next := pos[word[i]-'a']
		if next > cur {
			ans += next - cur
		} else {
			ans += cur - next
		}
		cur = next
	}
	return ans
}
```

## 1173 — Immediate Food Delivery I

```go
package main

// LeetCode #1173: Immediate Food Delivery I
// https://leetcode.com/problems/immediate-food-delivery-i/
// Difficulty: Easy [Paid] (SQL)
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println("SQL: SELECT ROUND(100.0 * SUM(CASE WHEN order_date = customer_pref_delivery_date THEN 1 ELSE 0 END) / COUNT(*), 2) AS immediate_percentage FROM Delivery")
}

// This is a SQL problem. The answer is the SQL query above.
```

## 1175 — Prime Arrangements

```go
package main

// LeetCode #1175: Prime Arrangements
// https://leetcode.com/problems/prime-arrangements/
// Difficulty: Easy
// Time: O(n log log n) | Space: O(n)

import "fmt"

const mod = 1_000_000_007

func main() {
	fmt.Println(numPrimeArrangements(5))  // 12
	fmt.Println(numPrimeArrangements(100)) // 682289015
}

// LeetCode submission: numPrimeArrangements
func numPrimeArrangements(n int) int {
	primeCount := countPrimes(n)
	nonPrimeCount := n - primeCount
	ans := 1
	for i := 1; i <= primeCount; i++ {
		ans = (ans * i) % mod
	}
	for i := 1; i <= nonPrimeCount; i++ {
		ans = (ans * i) % mod
	}
	return ans
}

func countPrimes(n int) int {
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i <= n; i++ {
		if isPrime[i] {
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}
	count := 0
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			count++
		}
	}
	return count
}
```

## 1176 — Diet Plan Performance

```go
package main

// LeetCode #1176: Diet Plan Performance
// https://leetcode.com/problems/diet-plan-performance/
// Difficulty: Easy [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(dietPlanPerformance([]int{1, 2, 3, 4, 5}, 1, 3, 3)) // 0
	fmt.Println(dietPlanPerformance([]int{3, 2}, 2, 0, 1))          // 1
}

// LeetCode submission: dietPlanPerformance
func dietPlanPerformance(calories []int, k int, lower int, upper int) int {
	n := len(calories)
	if n < k {
		return 0
	}
	sum := 0
	for i := 0; i < k; i++ {
		sum += calories[i]
	}
	ans := score(sum, lower, upper)
	for i := k; i < n; i++ {
		sum += calories[i] - calories[i-k]
		ans += score(sum, lower, upper)
	}
	return ans
}

func score(sum, lower, upper int) int {
	if sum < lower {
		return -1
	}
	if sum > upper {
		return 1
	}
	return 0
}
```

## 1179 — Reformat Department Table

```go
package main

// LeetCode #1179: Reformat Department Table
// https://leetcode.com/problems/reformat-department-table/
// Difficulty: Easy (SQL)
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println("SQL: SELECT id, SUM(CASE WHEN month='Jan' THEN revenue ELSE NULL END) AS Jan_Revenue, SUM(CASE WHEN month='Feb' THEN revenue ELSE NULL END) AS Feb_Revenue, SUM(CASE WHEN month='Mar' THEN revenue ELSE NULL END) AS Mar_Revenue, SUM(CASE WHEN month='Apr' THEN revenue ELSE NULL END) AS Apr_Revenue, SUM(CASE WHEN month='May' THEN revenue ELSE NULL END) AS May_Revenue, SUM(CASE WHEN month='Jun' THEN revenue ELSE NULL END) AS Jun_Revenue, SUM(CASE WHEN month='Jul' THEN revenue ELSE NULL END) AS Jul_Revenue, SUM(CASE WHEN month='Aug' THEN revenue ELSE NULL END) AS Aug_Revenue, SUM(CASE WHEN month='Sep' THEN revenue ELSE NULL END) AS Sep_Revenue, SUM(CASE WHEN month='Oct' THEN revenue ELSE NULL END) AS Oct_Revenue, SUM(CASE WHEN month='Nov' THEN revenue ELSE NULL END) AS Nov_Revenue, SUM(CASE WHEN month='Dec' THEN revenue ELSE NULL END) AS Dec_Revenue FROM Department GROUP BY id")
}

// This is a SQL problem. The answer is the SQL query above.
```

## 1180 — Count Substrings With Only One Distinct Letter

```go
package main

// LeetCode #1180: Count Substrings with Only One Distinct Letter
// https://leetcode.com/problems/count-substrings-with-only-one-distinct-letter/
// Difficulty: Easy [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(countLetters("aaaba")) // 8
	fmt.Println(countLetters("aaaaaaaaaa")) // 55
}

// LeetCode submission: countLetters
func countLetters(s string) int {
	ans := 0
	for i, n := 0, len(s); i < n; {
		j := i
		for j < n && s[j] == s[i] {
			j++
		}
		L := j - i
		ans += L * (L + 1) / 2
		i = j
	}
	return ans
}
```

## 1184 — Distance Between Bus Stops

```go
package main

// LeetCode #1184: Distance Between Bus Stops
// https://leetcode.com/problems/distance-between-bus-stops/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(distanceBetweenBusStops([]int{1, 2, 3, 4}, 0, 1)) // 1
	fmt.Println(distanceBetweenBusStops([]int{1, 2, 3, 4}, 0, 2)) // 3
	fmt.Println(distanceBetweenBusStops([]int{1, 2, 3, 4}, 0, 3)) // 4
}

// LeetCode submission: distanceBetweenBusStops
func distanceBetweenBusStops(distance []int, start, destination int) int {
	if start > destination {
		start, destination = destination, start
	}
	forward := 0
	for i := start; i < destination; i++ {
		forward += distance[i]
	}
	total := 0
	for _, d := range distance {
		total += d
	}
	backward := total - forward
	if backward < forward {
		return backward
	}
	return forward
}
```

## 1185 — Day Of The Week

```go
package main

// LeetCode #1185: Day of the Week
// https://leetcode.com/problems/day-of-the-week/
// Difficulty: Easy
// Time: O(1) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(dayOfTheWeek(31, 8, 2019)) // "Saturday"
	fmt.Println(dayOfTheWeek(18, 7, 1999)) // "Sunday"
	fmt.Println(dayOfTheWeek(15, 8, 1993)) // "Sunday"
}

// LeetCode submission: dayOfTheWeek
func dayOfTheWeek(day, month, year int) string {
	// Tomohiko Sakamoto's algorithm
	t := []int{0, 3, 2, 5, 0, 3, 5, 1, 4, 6, 2, 4}
	if month < 3 {
		year--
	}
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	idx := (year + year/4 - year/100 + year/400 + t[month-1] + day) % 7
	return days[idx]
}
```

## 1189 — Maximum Number Of Balloons

```go
package main

// LeetCode #1189: Maximum Number of Balloons
// https://leetcode.com/problems/maximum-number-of-balloons/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(maxNumberOfBalloons("nlaebolko"))           // 1
	fmt.Println(maxNumberOfBalloons("loonbalxballpoon"))    // 2
	fmt.Println(maxNumberOfBalloons("leetcode"))            // 0
}

// LeetCode submission: maxNumberOfBalloons
func maxNumberOfBalloons(text string) int {
	count := [26]int{}
	for i := 0; i < len(text); i++ {
		count[text[i]-'a']++
	}
	ans := count[1]           // b
	ans = min(ans, count[0])  // a
	ans = min(ans, count[11]/2) // l (needs 2)
	ans = min(ans, count[14]/2) // o (needs 2)
	ans = min(ans, count[13]) // n
	return ans
}
```

## 1196 — How Many Apples Can You Put Into The Basket

```go
package main

// LeetCode #1196: How Many Apples Can You Put into the Basket
// https://leetcode.com/problems/how-many-apples-can-you-put-into-the-basket/
// Difficulty: Easy [Paid]
// Time: O(n log n) | Space: O(log n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxNumberOfApples([]int{100, 200, 150, 1000}))          // 4
	fmt.Println(maxNumberOfApples([]int{900, 950, 800, 1000, 700, 800})) // 5
}

// LeetCode submission: maxNumberOfApples
func maxNumberOfApples(weight []int) int {
	sort.Ints(weight)
	sum := 0
	for i, w := range weight {
		sum += w
		if sum > 5000 {
			return i
		}
	}
	return len(weight)
}
```

## 1200 — Minimum Absolute Difference

```go
package main

// LeetCode #1200: Minimum Absolute Difference
// https://leetcode.com/problems/minimum-absolute-difference/
// Difficulty: Easy
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minimumAbsDifference([]int{4, 2, 1, 3}))       // [[1,2],[2,3],[3,4]]
	fmt.Println(minimumAbsDifference([]int{1, 3, 6, 10, 15}))  // [[1,3]]
}

// LeetCode submission: minimumAbsDifference
func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	minDiff := 1 << 31
	for i := 1; i < len(arr); i++ {
		diff := arr[i] - arr[i-1]
		if diff < minDiff {
			minDiff = diff
		}
	}
	var ans [][]int
	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] == minDiff {
			ans = append(ans, []int{arr[i-1], arr[i]})
		}
	}
	return ans
}
```

## 1207 — Unique Number Of Occurrences

```go
package main

// LeetCode #1207: Unique Number of Occurrences
// https://leetcode.com/problems/unique-number-of-occurrences/
// Difficulty: Easy
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(uniqueOccurrences([]int{1, 2, 2, 1, 1, 3})) // true
	fmt.Println(uniqueOccurrences([]int{1, 2}))             // false
	fmt.Println(uniqueOccurrences([]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0})) // true
}

// LeetCode submission: uniqueOccurrences
func uniqueOccurrences(arr []int) bool {
	freq := make(map[int]int)
	for _, v := range arr {
		freq[v]++
	}
	seen := make(map[int]bool)
	for _, f := range freq {
		if seen[f] {
			return false
		}
		seen[f] = true
	}
	return true
}
```

## 1211 — Queries Quality And Percentage

```go
package main

// LeetCode #1211: Queries Quality and Percentage
// https://leetcode.com/problems/queries-quality-and-percentage/
// Difficulty: Easy (SQL)
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println("SQL: SELECT query_name, ROUND(AVG(rating/position), 2) AS quality, ROUND(100.0 * SUM(CASE WHEN rating < 3 THEN 1 ELSE 0 END) / COUNT(*), 2) AS poor_query_percentage FROM Queries GROUP BY query_name")
}

// This is a SQL problem. The answer is the SQL query above.
```

## 1213 — Intersection Of Three Sorted Arrays

```go
package main

// LeetCode #1213: Intersection of Three Sorted Arrays
// https://leetcode.com/problems/intersection-of-three-sorted-arrays/
// Difficulty: Easy [Paid]
// Time: O(n) | Space: O(1) excluding output

import "fmt"

func main() {
	fmt.Println(arraysIntersection([]int{1, 2, 3, 4, 5}, []int{1, 2, 5, 7, 9}, []int{1, 3, 4, 5, 8}))
	// [1,5]
	fmt.Println(arraysIntersection([]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}))
	// []
}

// LeetCode submission: arraysIntersection
func arraysIntersection(arr1, arr2, arr3 []int) []int {
	var ans []int
	i, j, k := 0, 0, 0
	for i < len(arr1) && j < len(arr2) && k < len(arr3) {
		if arr1[i] == arr2[j] && arr2[j] == arr3[k] {
			ans = append(ans, arr1[i])
			i++; j++; k++
		} else if arr1[i] < arr2[j] {
			i++
		} else if arr2[j] < arr3[k] {
			j++
		} else {
			k++
		}
	}
	return ans
}
```

## 1217 — Minimum Cost To Move Chips To The Same Position

```go
package main

// LeetCode #1217: Minimum Cost to Move Chips to The Same Position
// https://leetcode.com/problems/minimum-cost-to-move-chips-to-the-same-position/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(minCostToMoveChips([]int{1, 2, 3}))    // 1
	fmt.Println(minCostToMoveChips([]int{2, 2, 2, 3, 3})) // 2
}

// LeetCode submission: minCostToMoveChips
func minCostToMoveChips(position []int) int {
	even, odd := 0, 0
	for _, p := range position {
		if p%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	if even < odd {
		return even
	}
	return odd
}
```

## 1221 — Split A String In Balanced Strings

```go
package main

// LeetCode #1221: Split a String in Balanced Strings
// https://leetcode.com/problems/split-a-string-in-balanced-strings/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(balancedStringSplit("RLRRLLRLRL")) // 4
	fmt.Println(balancedStringSplit("RLLLLRRRLR")) // 3
	fmt.Println(balancedStringSplit("LLLLRRRR"))   // 1
}

// LeetCode submission: balancedStringSplit
func balancedStringSplit(s string) int {
	count, ans := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'R' {
			count++
		} else {
			count--
		}
		if count == 0 {
			ans++
		}
	}
	return ans
}
```

