# Easy (Mudah) — Problem ��0412

## 0001 — Two Sum

```go
package main

// LeetCode #1: Two Sum
// https://leetcode.com/problems/two-sum/
// Difficulty: Easy
//
// Approaches (paste ONE twoSum function to LeetCode):
//   v1 HashMap:   Time O(n)   | Space O(n) | Mem 5-6MB | Runtime 0-4ms
//   v2 BruteForce: Time O(n²) | Space O(1) | Mem 4-5MB | Runtime 20-50ms
//   v3 UltraLow:  Time O(n²)  | Space O(1) | Mem 3-4MB | Runtime 15-30ms
//
// Catatan: Go runtime baseline ~2-3MB. Gak mungkin di bawah 2MB.
// LeetCode ukur RSS (resident set size) bukan heap aja.

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9)) // [0,1]
	fmt.Println(twoSum([]int{3, 2, 4}, 6))      // [1,2]
	fmt.Println(twoSum([]int{3, 3}, 6))         // [0,1]
}

// --- v1 HashMap: O(n) time, O(n) space ---
func twoSumHashMap(nums []int, target int) []int {
	seen := make(map[int]int, len(nums))
	for i, n := range nums {
		if j, ok := seen[target-n]; ok {
			return []int{j, i}
		}
		seen[n] = i
	}
	return nil
}

// --- v2 Brute Force: O(n²) time, O(1) space ---
func twoSumBruteForce(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// --- v3 Ultra Low Memory ---
// Hindari return nil (extra branch), minimalkan live vars, gunakan len(nums)
// dan akses indeks langsung tanpa variabel lokal tambahan.
func twoSum(nums []int, target int) []int {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		vi := nums[i]
		want := target - vi
		for j := i + 1; j < n; j++ {
			if nums[j] == want {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
```

## 0009 — Palindrome Number

```go
package main

// LeetCode #9: Palindrome Number
// https://leetcode.com/problems/palindrome-number/
// Difficulty: Easy

import "fmt"

// Time: O(log n) | Space: O(1)
func IsPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	reverted := 0
	for x > reverted {
		reverted = reverted*10 + x%10
		x /= 10
	}
	return x == reverted || x == reverted/10
}

func main() {
	fmt.Println(IsPalindrome(121))  // true
	fmt.Println(IsPalindrome(-121)) // false
	fmt.Println(IsPalindrome(10))   // false
}
```

## 0013 — Roman To Integer

```go
package main

// LeetCode #13: Roman to Integer
// https://leetcode.com/problems/roman-to-integer/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(1)
func RomanToInt(s string) int {
	vals := map[byte]int{
		'I': 1, 'V': 5, 'X': 10, 'L': 50,
		'C': 100, 'D': 500, 'M': 1000,
	}
	sum, prev := 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		cur := vals[s[i]]
		if cur < prev {
			sum -= cur
		} else {
			sum += cur
		}
		prev = cur
	}
	return sum
}

func main() {
	fmt.Println(RomanToInt("III"))
	fmt.Println(RomanToInt("LVIII"))
	fmt.Println(RomanToInt("MCMXCIV"))
}
```

## 0014 — Longest Common Prefix

```go
package main

// LeetCode #14: Longest Common Prefix
// https://leetcode.com/problems/longest-common-prefix/
// Difficulty: Easy

import "fmt"

// Time: O(n*m) | Space: O(1) where n=len(strs), m=len(shortest string)
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		c := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != c {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

func main() {
	fmt.Println(LongestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(LongestCommonPrefix([]string{"dog", "racecar", "car"}))
	fmt.Println(LongestCommonPrefix([]string{"a"}))
}
```

## 0020 — Valid Parentheses

```go
package main

// LeetCode #20: Valid Parentheses
// https://leetcode.com/problems/valid-parentheses/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(n)
func IsValid(s string) bool {
	pairs := map[byte]byte{')': '(', ']': '[', '}': '{'}
	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(IsValid("()"))
	fmt.Println(IsValid("()[]{}"))
	fmt.Println(IsValid("(]"))
}
```

## 0021 — Merge Two Sorted Lists

```go
package main

// LeetCode #21: Merge Two Sorted Lists
// https://leetcode.com/problems/merge-two-sorted-lists/
// Difficulty: Easy

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time: O(n+m) | Space: O(1)
func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}
		curr = curr.Next
	}
	if list1 != nil {
		curr.Next = list1
	} else {
		curr.Next = list2
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
	l1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	l2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	printList(MergeTwoLists(l1, l2))

	l3 := &ListNode{}
	l4 := &ListNode{}
	printList(MergeTwoLists(l3, l4))
}
```

## 0026 — Remove Duplicates From Sorted Array

```go
package main

// LeetCode #26: Remove Duplicates from Sorted Array
// https://leetcode.com/problems/remove-duplicates-from-sorted-array/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(1)
func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	k := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[k-1] {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}

func main() {
	n1 := []int{1, 1, 2}
	fmt.Println(RemoveDuplicates(n1), n1)
	n2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(RemoveDuplicates(n2), n2)
}
```

## 0027 — Remove Element

```go
package main

// LeetCode #27: Remove Element
// https://leetcode.com/problems/remove-element/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(1)
func RemoveElement(nums []int, val int) int {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}

func main() {
	n1 := []int{3, 2, 2, 3}
	fmt.Println(RemoveElement(n1, 3), n1)
	n2 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(RemoveElement(n2, 2), n2)
}
```

## 0028 — Find The Index Of The First Occurrence In A String

```go
package main

// LeetCode #28: Find the Index of the First Occurrence in a String
// https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/
// Difficulty: Easy

import "fmt"

// Time: O(n*m) | Space: O(1)
func StrStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(StrStr("sadbutsad", "sad"))
	fmt.Println(StrStr("leetcode", "leeto"))
	fmt.Println(StrStr("hello", "ll"))
}
```

## 0035 — Search Insert Position

```go
package main

// LeetCode #35: Search Insert Position
// https://leetcode.com/problems/search-insert-position/
// Difficulty: Easy

import "fmt"

// Time: O(log n) | Space: O(1)
func SearchInsert(nums []int, target int) int {
	lo, hi := 0, len(nums)
	for lo < hi {
		mid := lo + (hi-lo)/2
		if nums[mid] >= target {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func main() {
	fmt.Println(SearchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Println(SearchInsert([]int{1, 3, 5, 6}, 2))
	fmt.Println(SearchInsert([]int{1, 3, 5, 6}, 7))
}
```

## 0058 — Length Of Last Word

```go
package main

// LeetCode #58: Length of Last Word
// https://leetcode.com/problems/length-of-last-word/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(1)
func LengthOfLastWord(s string) int {
	length := 0
	i := len(s) - 1
	for i >= 0 && s[i] == ' ' {
		i--
	}
	for i >= 0 && s[i] != ' ' {
		length++
		i--
	}
	return length
}

func main() {
	fmt.Println(LengthOfLastWord("Hello World"))
	fmt.Println(LengthOfLastWord("   fly me   to   the moon  "))
	fmt.Println(LengthOfLastWord("luffy is still joyboy"))
}
```

## 0066 — Plus One

```go
package main

// LeetCode #66: Plus One
// https://leetcode.com/problems/plus-one/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(1) (excluding output)
func PlusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	return append([]int{1}, digits...)
}

func main() {
	fmt.Println(PlusOne([]int{1, 2, 3}))
	fmt.Println(PlusOne([]int{4, 3, 2, 1}))
	fmt.Println(PlusOne([]int{9}))
}
```

## 0067 — Add Binary

```go
package main

// LeetCode #67: Add Binary
// https://leetcode.com/problems/add-binary/
// Difficulty: Easy

import "fmt"

// Time: O(max(n,m)) | Space: O(max(n,m))
func AddBinary(a string, b string) string {
	i, j, carry := len(a)-1, len(b)-1, 0
	res := make([]byte, 0, max(len(a), len(b))+1)
	for i >= 0 || j >= 0 || carry > 0 {
		if i >= 0 {
			carry += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			carry += int(b[j] - '0')
			j--
		}
		res = append(res, byte('0'+carry%2))
		carry /= 2
	}
	// reverse
	for l, r := 0, len(res)-1; l < r; l, r = l+1, r-1 {
		res[l], res[r] = res[r], res[l]
	}
	return string(res)
}

func main() {
	fmt.Println(AddBinary("11", "1"))
	fmt.Println(AddBinary("1010", "1011"))
}
```

## 0069 — Sqrtx

```go
package main

// LeetCode #69: Sqrt(x)
// https://leetcode.com/problems/sqrtx/
// Difficulty: Easy

import "fmt"

// Time: O(log n) | Space: O(1)
func MySqrt(x int) int {
	if x < 2 {
		return x
	}
	lo, hi := 1, x/2
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if mid*mid == x {
			return mid
		} else if mid*mid < x {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return hi
}

func main() {
	fmt.Println(MySqrt(4))
	fmt.Println(MySqrt(8))
	fmt.Println(MySqrt(0))
}
```

## 0070 — Climbing Stairs

```go
package main

// LeetCode #70: Climbing Stairs
// https://leetcode.com/problems/climbing-stairs/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(1)
func ClimbStairs(n int) int {
	if n <= 2 {
		return n
	}
	a, b := 1, 2
	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func main() {
	fmt.Println(ClimbStairs(2))
	fmt.Println(ClimbStairs(3))
	fmt.Println(ClimbStairs(4))
}
```

## 0083 — Remove Duplicates From Sorted List

```go
package main

// LeetCode #83: Remove Duplicates from Sorted List
// https://leetcode.com/problems/remove-duplicates-from-sorted-list/
// Difficulty: Easy

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time: O(n) | Space: O(1)
func DeleteDuplicates(head *ListNode) *ListNode {
	curr := head
	for curr != nil && curr.Next != nil {
		if curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
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
	l1 := &ListNode{1, &ListNode{1, &ListNode{2, nil}}}
	printList(DeleteDuplicates(l1))
	l2 := &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, nil}}}}}
	printList(DeleteDuplicates(l2))
}
```

## 0088 — Merge Sorted Array

```go
package main

// LeetCode #88: Merge Sorted Array
// https://leetcode.com/problems/merge-sorted-array/
// Difficulty: Easy

import "fmt"

// Time: O(m+n) | Space: O(1)
func Merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1
	for j >= 0 {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
}

func main() {
	n1 := []int{1, 2, 3, 0, 0, 0}
	Merge(n1, 3, []int{2, 5, 6}, 3)
	fmt.Println(n1)
	n2 := []int{1}
	Merge(n2, 1, []int{}, 0)
	fmt.Println(n2)
}
```

## 0094 — Binary Tree Inorder Traversal

```go
package main

// LeetCode #94: Binary Tree Inorder Traversal
// https://leetcode.com/problems/binary-tree-inorder-traversal/
// Difficulty: Easy

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n) | Space: O(n)
func InorderTraversal(root *TreeNode) []int {
	var res []int
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		res = append(res, node.Val)
		dfs(node.Right)
	}
	dfs(root)
	return res
}

func main() {
	root := &TreeNode{1, nil, &TreeNode{2, &TreeNode{3, nil, nil}, nil}}
	fmt.Println(InorderTraversal(root))
	fmt.Println(InorderTraversal(nil))
}
```

## 0100 — Same Tree

```go
package main

// LeetCode #100: Same Tree
// https://leetcode.com/problems/same-tree/
// Difficulty: Easy

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n) | Space: O(h) where h is tree height
func IsSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}
	return IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
}

func main() {
	t1 := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	t2 := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	fmt.Println(IsSameTree(t1, t2))
	t3 := &TreeNode{1, &TreeNode{2, nil, nil}, nil}
	fmt.Println(IsSameTree(t1, t3))
}
```

## 0101 — Symmetric Tree

```go
package main

// LeetCode #101: Symmetric Tree
// https://leetcode.com/problems/symmetric-tree/
// Difficulty: Easy

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n) | Space: O(h)
func IsSymmetric(root *TreeNode) bool {
	var check func(*TreeNode, *TreeNode) bool
	check = func(a, b *TreeNode) bool {
		if a == nil && b == nil {
			return true
		}
		if a == nil || b == nil || a.Val != b.Val {
			return false
		}
		return check(a.Left, b.Right) && check(a.Right, b.Left)
	}
	return check(root, root)
}

func main() {
	root := &TreeNode{1, &TreeNode{2, &TreeNode{3, nil, nil}, &TreeNode{4, nil, nil}}, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{3, nil, nil}}}
	fmt.Println(IsSymmetric(root))
	root2 := &TreeNode{1, &TreeNode{2, nil, &TreeNode{3, nil, nil}}, &TreeNode{2, nil, &TreeNode{3, nil, nil}}}
	fmt.Println(IsSymmetric(root2))
}
```

## 0104 — Maximum Depth Of Binary Tree

```go
package main

// LeetCode #104: Maximum Depth of Binary Tree
// https://leetcode.com/problems/maximum-depth-of-binary-tree/
// Difficulty: Easy

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n) | Space: O(h)
func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := MaxDepth(root.Left)
	right := MaxDepth(root.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}

func main() {
	root := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	fmt.Println(MaxDepth(root))
	fmt.Println(MaxDepth(nil))
}
```

## 0108 — Convert Sorted Array To Binary Search Tree

```go
package main

// LeetCode #108: Convert Sorted Array to Binary Search Tree
// https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/
// Difficulty: Easy

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n) | Space: O(log n) (recursion stack)
func SortedArrayToBST(nums []int) *TreeNode {
	var build func(int, int) *TreeNode
	build = func(lo, hi int) *TreeNode {
		if lo > hi {
			return nil
		}
		mid := lo + (hi-lo)/2
		return &TreeNode{
			Val:   nums[mid],
			Left:  build(lo, mid-1),
			Right: build(mid+1, hi),
		}
	}
	return build(0, len(nums)-1)
}

func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left)
	fmt.Print(root.Val, " ")
	inorder(root.Right)
}

func main() {
	inorder(SortedArrayToBST([]int{-10, -3, 0, 5, 9}))
	fmt.Println()
	inorder(SortedArrayToBST([]int{1, 3}))
	fmt.Println()
}
```

## 0110 — Balanced Binary Tree

```go
package main

// LeetCode #110: Balanced Binary Tree
// https://leetcode.com/problems/balanced-binary-tree/
// Difficulty: Easy

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n) | Space: O(h)
func IsBalanced(root *TreeNode) bool {
	var height func(*TreeNode) int
	height = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l := height(node.Left)
		r := height(node.Right)
		if l == -1 || r == -1 || l-r > 1 || r-l > 1 {
			return -1
		}
		if l > r {
			return l + 1
		}
		return r + 1
	}
	return height(root) != -1
}

func main() {
	root := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	fmt.Println(IsBalanced(root))
	root2 := &TreeNode{1, &TreeNode{2, &TreeNode{3, &TreeNode{4, nil, nil}, nil}, nil}, &TreeNode{2, nil, nil}}
	fmt.Println(IsBalanced(root2))
}
```

## 0111 — Minimum Depth Of Binary Tree

```go
package main

// LeetCode #111: Minimum Depth of Binary Tree
// https://leetcode.com/problems/minimum-depth-of-binary-tree/
// Difficulty: Easy

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n) | Space: O(h)
func MinDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil {
		return MinDepth(root.Right) + 1
	}
	if root.Right == nil {
		return MinDepth(root.Left) + 1
	}
	l, r := MinDepth(root.Left), MinDepth(root.Right)
	if l < r {
		return l + 1
	}
	return r + 1
}

func main() {
	root := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	fmt.Println(MinDepth(root))
	root2 := &TreeNode{2, nil, &TreeNode{3, nil, &TreeNode{4, nil, &TreeNode{5, nil, &TreeNode{6, nil, nil}}}}}
	fmt.Println(MinDepth(root2))
}
```

## 0112 — Path Sum

```go
package main

// LeetCode #112: Path Sum
// https://leetcode.com/problems/path-sum/
// Difficulty: Easy

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n) | Space: O(h)
func HasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	targetSum -= root.Val
	if root.Left == nil && root.Right == nil {
		return targetSum == 0
	}
	return HasPathSum(root.Left, targetSum) || HasPathSum(root.Right, targetSum)
}

func main() {
	root := &TreeNode{5, &TreeNode{4, &TreeNode{11, &TreeNode{7, nil, nil}, &TreeNode{2, nil, nil}}, nil}, &TreeNode{8, &TreeNode{13, nil, nil}, &TreeNode{4, nil, &TreeNode{1, nil, nil}}}}
	fmt.Println(HasPathSum(root, 22))
	fmt.Println(HasPathSum(nil, 0))
}
```

## 0118 — Pascals Triangle

```go
package main

// LeetCode #118: Pascal's Triangle
// https://leetcode.com/problems/pascals-triangle/
// Difficulty: Easy

import "fmt"

// Time: O(numRows^2) | Space: O(numRows^2)
func Generate(numRows int) [][]int {
	res := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		res[i] = make([]int, i+1)
		res[i][0], res[i][i] = 1, 1
		for j := 1; j < i; j++ {
			res[i][j] = res[i-1][j-1] + res[i-1][j]
		}
	}
	return res
}

func main() {
	fmt.Println(Generate(5))
	fmt.Println(Generate(1))
}
```

## 0119 — Pascals Triangle Ii

```go
package main

// LeetCode #119: Pascal's Triangle II
// https://leetcode.com/problems/pascals-triangle-ii/
// Difficulty: Easy

import "fmt"

// Time: O(rowIndex^2) | Space: O(rowIndex)
func GetRow(rowIndex int) []int {
	res := make([]int, rowIndex+1)
	res[0] = 1
	for i := 1; i <= rowIndex; i++ {
		for j := i; j > 0; j-- {
			res[j] += res[j-1]
		}
	}
	return res
}

func main() {
	fmt.Println(GetRow(3))
	fmt.Println(GetRow(0))
	fmt.Println(GetRow(4))
}
```

## 0121 — Best Time To Buy And Sell Stock

```go
package main

// LeetCode #121: Best Time to Buy and Sell Stock
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(1)
func MaxProfit(prices []int) int {
	minPrice, maxProfit := prices[0], 0
	for _, p := range prices[1:] {
		if p < minPrice {
			minPrice = p
		} else if p-minPrice > maxProfit {
			maxProfit = p - minPrice
		}
	}
	return maxProfit
}

func main() {
	fmt.Println(MaxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(MaxProfit([]int{7, 6, 4, 3, 1}))
}
```

## 0125 — Valid Palindrome

```go
package main

// LeetCode #125: Valid Palindrome
// https://leetcode.com/problems/valid-palindrome/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(1)
func IsPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		for i < j && !isAlphaNum(s[i]) {
			i++
		}
		for i < j && !isAlphaNum(s[j]) {
			j--
		}
		if toLower(s[i]) != toLower(s[j]) {
			return false
		}
		i++
		j--
	}
	return true
}

func isAlphaNum(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}

func main() {
	fmt.Println(IsPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(IsPalindrome("race a car"))
	fmt.Println(IsPalindrome(" "))
}
```

## 0136 — Single Number

```go
package main

// LeetCode #136: Single Number
// https://leetcode.com/problems/single-number/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(1)
func SingleNumber(nums []int) int {
	result := 0
	for _, n := range nums {
		result ^= n
	}
	return result
}

func main() {
	fmt.Println(SingleNumber([]int{2, 2, 1}))
	fmt.Println(SingleNumber([]int{4, 1, 2, 1, 2}))
	fmt.Println(SingleNumber([]int{1}))
}
```

## 0141 — Linked List Cycle

```go
package main

// LeetCode #141: Linked List Cycle
// https://leetcode.com/problems/linked-list-cycle/
// Difficulty: Easy

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time: O(n) | Space: O(1)
func HasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func main() {
	n1 := &ListNode{3, nil}
	n2 := &ListNode{2, nil}
	n3 := &ListNode{0, nil}
	n4 := &ListNode{-4, nil}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n2
	fmt.Println(HasCycle(n1))

	single := &ListNode{1, nil}
	fmt.Println(HasCycle(single))
}
```

## 0144 — Binary Tree Preorder Traversal

```go
package main

// LeetCode #144: Binary Tree Preorder Traversal
// https://leetcode.com/problems/binary-tree-preorder-traversal/
// Difficulty: Easy

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n) | Space: O(n)
func PreorderTraversal(root *TreeNode) []int {
	var res []int
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return res
}

func main() {
	root := &TreeNode{1, nil, &TreeNode{2, &TreeNode{3, nil, nil}, nil}}
	fmt.Println(PreorderTraversal(root))
	fmt.Println(PreorderTraversal(nil))
}
```

## 0145 — Binary Tree Postorder Traversal

```go
package main

// LeetCode #145: Binary Tree Postorder Traversal
// https://leetcode.com/problems/binary-tree-postorder-traversal/
// Difficulty: Easy

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n) | Space: O(n)
func PostorderTraversal(root *TreeNode) []int {
	var res []int
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		dfs(node.Right)
		res = append(res, node.Val)
	}
	dfs(root)
	return res
}

func main() {
	root := &TreeNode{1, nil, &TreeNode{2, &TreeNode{3, nil, nil}, nil}}
	fmt.Println(PostorderTraversal(root))
	fmt.Println(PostorderTraversal(nil))
}
```

## 0157 — Read N Characters Given Read4

```go
package main

// LeetCode #157: Read N Characters Given Read4
// https://leetcode.com/problems/read-n-characters-given-read4/
// Difficulty: Easy [Paid]

import "fmt"

var _buf []byte
var _pos int

func read4(buf4 []byte) int {
	n := 0
	for i := 0; i < 4 && _pos < len(_buf); i++ {
		buf4[i] = _buf[_pos]
		_pos++
		n++
	}
	return n
}

// Time: O(n) | Space: O(1)
func Read(buf []byte, n int) int {
	buf4 := make([]byte, 4)
	total := 0
	for total < n {
		count := read4(buf4)
		if count == 0 {
			break
		}
		for i := 0; i < count && total < n; i++ {
			buf[total] = buf4[i]
			total++
		}
	}
	return total
}

func main() {
	_buf = []byte("abc")
	_pos = 0
	buf := make([]byte, 10)
	fmt.Println(Read(buf, 4), string(buf[:3]))

	_buf = []byte("abcde")
	_pos = 0
	buf2 := make([]byte, 10)
	fmt.Println(Read(buf2, 5), string(buf2[:5]))
}
```

## 0160 — Intersection Of Two Linked Lists

```go
package main

// LeetCode #160: Intersection of Two Linked Lists
// https://leetcode.com/problems/intersection-of-two-linked-lists/
// Difficulty: Easy

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time: O(n+m) | Space: O(1)
func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	a, b := headA, headB
	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}
		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}
	return a
}

func main() {
	common := &ListNode{8, &ListNode{4, &ListNode{5, nil}}}
	a := &ListNode{4, &ListNode{1, common}}
	b := &ListNode{5, &ListNode{6, &ListNode{1, common}}}
	fmt.Println(GetIntersectionNode(a, b).Val)
}
```

## 0163 — Missing Ranges

```go
package main

// LeetCode #163: Missing Ranges
// https://leetcode.com/problems/missing-ranges/
// Difficulty: Easy [Paid]

import (
	"fmt"
	"strconv"
)

// Time: O(n) | Space: O(1) excluding output
func FindMissingRanges(nums []int, lower int, upper int) []string {
	var res []string
	addRange := func(lo, hi int) {
		if lo > hi {
			return
		}
		if lo == hi {
			res = append(res, strconv.Itoa(lo))
		} else {
			res = append(res, strconv.Itoa(lo)+"->"+strconv.Itoa(hi))
		}
	}
	prev := lower - 1
	for i := 0; i <= len(nums); i++ {
		var curr int
		if i < len(nums) {
			curr = nums[i]
		} else {
			curr = upper + 1
		}
		if curr-prev > 1 {
			addRange(prev+1, curr-1)
		}
		prev = curr
	}
	return res
}

func main() {
	fmt.Println(FindMissingRanges([]int{0, 1, 3, 50, 75}, 0, 99))
	fmt.Println(FindMissingRanges([]int{-1}, -1, -1))
}
```

## 0168 — Excel Sheet Column Title

```go
package main

// LeetCode #168: Excel Sheet Column Title
// https://leetcode.com/problems/excel-sheet-column-title/
// Difficulty: Easy

import "fmt"

// Time: O(log n) | Space: O(log n)
func ConvertToTitle(columnNumber int) string {
	res := make([]byte, 0, 8)
	for columnNumber > 0 {
		columnNumber--
		res = append(res, byte('A'+columnNumber%26))
		columnNumber /= 26
	}
	// reverse
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}

func main() {
	fmt.Println(ConvertToTitle(1))
	fmt.Println(ConvertToTitle(28))
	fmt.Println(ConvertToTitle(701))
}
```

## 0169 — Majority Element

```go
package main

// LeetCode #169: Majority Element
// https://leetcode.com/problems/majority-element/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(1)
func MajorityElement(nums []int) int {
	candidate, count := 0, 0
	for _, n := range nums {
		if count == 0 {
			candidate = n
		}
		if n == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}

func main() {
	fmt.Println(MajorityElement([]int{3, 2, 3}))
	fmt.Println(MajorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}
```

## 0170 — Two Sum Iii Data Structure Design

```go
package main

// LeetCode #170: Two Sum III - Data structure design
// https://leetcode.com/problems/two-sum-iii-data-structure-design/
// Difficulty: Easy [Paid]

import "fmt"

type TwoSum struct {
	nums map[int]int
}

func Constructor() TwoSum {
	return TwoSum{nums: make(map[int]int)}
}

func (t *TwoSum) Add(number int) {
	t.nums[number]++
}

// Time: O(n) | Space: O(n)
func (t *TwoSum) Find(value int) bool {
	for num := range t.nums {
		want := value - num
		if want == num && t.nums[num] > 1 {
			return true
		}
		if want != num && t.nums[want] > 0 {
			return true
		}
	}
	return false
}

func main() {
	t := Constructor()
	t.Add(1)
	t.Add(3)
	t.Add(5)
	fmt.Println(t.Find(4))
	fmt.Println(t.Find(7))
}
```

## 0171 — Excel Sheet Column Number

```go
package main

// LeetCode #171: Excel Sheet Column Number
// https://leetcode.com/problems/excel-sheet-column-number/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(1)
func TitleToNumber(columnTitle string) int {
	result := 0
	for i := 0; i < len(columnTitle); i++ {
		result = result*26 + int(columnTitle[i]-'A'+1)
	}
	return result
}

func main() {
	fmt.Println(TitleToNumber("A"))
	fmt.Println(TitleToNumber("AB"))
	fmt.Println(TitleToNumber("ZY"))
}
```

## 0175 — Combine Two Tables

```go
package main

// LeetCode #175: Combine Two Tables
// https://leetcode.com/problems/combine-two-tables/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(1)
// Write your MySQL query statement below
func CombineTwoTables() string {
	return "SELECT Person.firstName, Person.lastName, Address.city, Address.state FROM Person LEFT JOIN Address ON Person.personId = Address.personId"
}

func main() {
	fmt.Println(CombineTwoTables())
}
```

## 0181 — Employees Earning More Than Their Managers

```go
package main

// LeetCode #181: Employees Earning More Than Their Managers
// https://leetcode.com/problems/employees-earning-more-than-their-managers/
// Difficulty: Easy

import "fmt"

func EmployeesEarningMoreThanTheirManagers() string {
	return "SELECT e.name AS Employee FROM Employee e JOIN Employee m ON e.managerId = m.id WHERE e.salary > m.salary"
}

func main() {
	fmt.Println(EmployeesEarningMoreThanTheirManagers())
}
```

## 0182 — Duplicate Emails

```go
package main

// LeetCode #182: Duplicate Emails
// https://leetcode.com/problems/duplicate-emails/
// Difficulty: Easy

import "fmt"

func DuplicateEmails() string {
	return "SELECT email FROM Person GROUP BY email HAVING COUNT(email) > 1"
}

func main() {
	fmt.Println(DuplicateEmails())
}
```

## 0183 — Customers Who Never Order

```go
package main

// LeetCode #183: Customers Who Never Order
// https://leetcode.com/problems/customers-who-never-order/
// Difficulty: Easy

import "fmt"

func CustomersWhoNeverOrder() string {
	return "SELECT c.name AS Customers FROM Customers c LEFT JOIN Orders o ON c.id = o.customerId WHERE o.customerId IS NULL"
}

func main() {
	fmt.Println(CustomersWhoNeverOrder())
}
```

## 0190 — Reverse Bits

```go
package main

// LeetCode #190: Reverse Bits
// https://leetcode.com/problems/reverse-bits/
// Difficulty: Easy

import "fmt"

// Time: O(1) | Space: O(1)
func ReverseBits(num uint32) uint32 {
	var result uint32
	for i := 0; i < 32; i++ {
		result <<= 1
		result |= num & 1
		num >>= 1
	}
	return result
}

func main() {
	fmt.Println(ReverseBits(43261596))    // 964176192
	fmt.Println(ReverseBits(4294967293))  // 3221225471
}
```

## 0191 — Number Of 1 Bits

```go
package main

// LeetCode #191: Number of 1 Bits
// https://leetcode.com/problems/number-of-1-bits/
// Difficulty: Easy

import "fmt"

// Time: O(1) | Space: O(1)
func HammingWeight(num uint32) int {
	count := 0
	for num != 0 {
		num &= num - 1
		count++
	}
	return count
}

func main() {
	fmt.Println(HammingWeight(11))    // 3
	fmt.Println(HammingWeight(128))   // 1
	fmt.Println(HammingWeight(4294967293)) // 31
}
```

## 0193 — Valid Phone Numbers

```go
package main

// LeetCode #193: Valid Phone Numbers
// https://leetcode.com/problems/valid-phone-numbers/
// Difficulty: Easy

import "fmt"

func ValidPhoneNumbers() string {
	return `grep -E '^(\([0-9]{3}\) [0-9]{3}-[0-9]{4}|[0-9]{3}-[0-9]{3}-[0-9]{4})$' file.txt`
}

func main() {
	fmt.Println(ValidPhoneNumbers())
}
```

## 0195 — Tenth Line

```go
package main

// LeetCode #195: Tenth Line
// https://leetcode.com/problems/tenth-line/
// Difficulty: Easy

import "fmt"

func TenthLine() string {
	return "sed -n '10p' file.txt"
}

func main() {
	fmt.Println(TenthLine())
}
```

## 0196 — Delete Duplicate Emails

```go
package main

// LeetCode #196: Delete Duplicate Emails
// https://leetcode.com/problems/delete-duplicate-emails/
// Difficulty: Easy

import "fmt"

func DeleteDuplicateEmails() string {
	return "DELETE p1 FROM Person p1, Person p2 WHERE p1.email = p2.email AND p1.id > p2.id"
}

func main() {
	fmt.Println(DeleteDuplicateEmails())
}
```

## 0197 — Rising Temperature

```go
package main

// LeetCode #197: Rising Temperature
// https://leetcode.com/problems/rising-temperature/
// Difficulty: Easy

import "fmt"

func RisingTemperature() string {
	return "SELECT w1.id FROM Weather w1 JOIN Weather w2 ON DATEDIFF(w1.recordDate, w2.recordDate) = 1 WHERE w1.temperature > w2.temperature"
}

func main() {
	fmt.Println(RisingTemperature())
}
```

## 0202 — Happy Number

```go
package main

// LeetCode #202: Happy Number
// https://leetcode.com/problems/happy-number/
// Difficulty: Easy

import "fmt"

// Time: O(log n) | Space: O(1)
func IsHappy(n int) bool {
	next := func(x int) int {
		sum := 0
		for x > 0 {
			d := x % 10
			sum += d * d
			x /= 10
		}
		return sum
	}
	slow, fast := n, next(n)
	for fast != 1 && slow != fast {
		slow = next(slow)
		fast = next(next(fast))
	}
	return fast == 1
}

func main() {
	fmt.Println(IsHappy(19))
	fmt.Println(IsHappy(2))
}
```

## 0203 — Remove Linked List Elements

```go
package main

// LeetCode #203: Remove Linked List Elements
// https://leetcode.com/problems/remove-linked-list-elements/
// Difficulty: Easy

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time: O(n) | Space: O(1)
func RemoveElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	curr := dummy
	for curr.Next != nil {
		if curr.Next.Val == val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
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
	l1 := &ListNode{1, &ListNode{2, &ListNode{6, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}}
	printList(RemoveElements(l1, 6))
	l2 := &ListNode{7, &ListNode{7, &ListNode{7, &ListNode{7, nil}}}}
	printList(RemoveElements(l2, 7))
}
```

## 0205 — Isomorphic Strings

```go
package main

// LeetCode #205: Isomorphic Strings
// https://leetcode.com/problems/isomorphic-strings/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(1) (fixed ASCII chars)
func IsIsomorphic(s string, t string) bool {
	m1 := make([]int, 256)
	m2 := make([]int, 256)
	for i := 0; i < len(s); i++ {
		if m1[s[i]] != m2[t[i]] {
			return false
		}
		m1[s[i]] = i + 1
		m2[t[i]] = i + 1
	}
	return true
}

func main() {
	fmt.Println(IsIsomorphic("egg", "add"))
	fmt.Println(IsIsomorphic("foo", "bar"))
	fmt.Println(IsIsomorphic("paper", "title"))
}
```

## 0206 — Reverse Linked List

```go
package main

// LeetCode #206: Reverse Linked List
// https://leetcode.com/problems/reverse-linked-list/
// Difficulty: Easy

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time: O(n) | Space: O(1)
func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}

func main() {
	l1 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	printList(ReverseList(l1))
	l2 := &ListNode{1, &ListNode{2, nil}}
	printList(ReverseList(l2))
}
```

## 0217 — Contains Duplicate

```go
package main

// LeetCode #217: Contains Duplicate
// https://leetcode.com/problems/contains-duplicate/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(n)
func ContainsDuplicate(nums []int) bool {
	seen := make(map[int]struct{}, len(nums))
	for _, n := range nums {
		if _, ok := seen[n]; ok {
			return true
		}
		seen[n] = struct{}{}
	}
	return false
}

func main() {
	fmt.Println(ContainsDuplicate([]int{1, 2, 3, 1}))
	fmt.Println(ContainsDuplicate([]int{1, 2, 3, 4}))
	fmt.Println(ContainsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
}
```

## 0219 — Contains Duplicate Ii

```go
package main

// LeetCode #219: Contains Duplicate II
// https://leetcode.com/problems/contains-duplicate-ii/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(n)
func ContainsNearbyDuplicate(nums []int, k int) bool {
	seen := make(map[int]int, len(nums))
	for i, n := range nums {
		if j, ok := seen[n]; ok && i-j <= k {
			return true
		}
		seen[n] = i
	}
	return false
}

func main() {
	fmt.Println(ContainsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
	fmt.Println(ContainsNearbyDuplicate([]int{1, 0, 1, 1}, 1))
	fmt.Println(ContainsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
}
```

## 0222 — Count Complete Tree Nodes

```go
package main

// LeetCode #222: Count Complete Tree Nodes
// https://leetcode.com/problems/count-complete-tree-nodes/
// Difficulty: Easy

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(log^2 n) | Space: O(log n)
func CountNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := func(node *TreeNode) int {
		d := 0
		for node != nil {
			d++
			node = node.Left
		}
		return d
	}
	rightDepth := func(node *TreeNode) int {
		d := 0
		for node != nil {
			d++
			node = node.Right
		}
		return d
	}
	l, r := leftDepth(root), rightDepth(root)
	if l == r {
		return (1 << l) - 1
	}
	return 1 + CountNodes(root.Left) + CountNodes(root.Right)
}

func main() {
	root := &TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}, &TreeNode{3, &TreeNode{6, nil, nil}, nil}}
	fmt.Println(CountNodes(root))
	fmt.Println(CountNodes(nil))
}
```

## 0225 — Implement Stack Using Queues

```go
package main

// LeetCode #225: Implement Stack using Queues
// https://leetcode.com/problems/implement-stack-using-queues/
// Difficulty: Easy

import "fmt"

type MyStack struct {
	q []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (s *MyStack) Push(x int) {
	s.q = append(s.q, x)
	for i := 0; i < len(s.q)-1; i++ {
		s.q = append(s.q, s.q[0])
		s.q = s.q[1:]
	}
}

func (s *MyStack) Pop() int {
	x := s.q[0]
	s.q = s.q[1:]
	return x
}

func (s *MyStack) Top() int {
	return s.q[0]
}

func (s *MyStack) Empty() bool {
	return len(s.q) == 0
}

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	fmt.Println(obj.Top())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Empty())
}
```

## 0226 — Invert Binary Tree

```go
package main

// LeetCode #226: Invert Binary Tree
// https://leetcode.com/problems/invert-binary-tree/
// Difficulty: Easy

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n) | Space: O(h)
func InvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = InvertTree(root.Right), InvertTree(root.Left)
	return root
}

func preorder(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Print(root.Val, " ")
	preorder(root.Left)
	preorder(root.Right)
}

func main() {
	root := &TreeNode{4, &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}}, &TreeNode{7, &TreeNode{6, nil, nil}, &TreeNode{9, nil, nil}}}
	preorder(InvertTree(root))
	fmt.Println()
}
```

## 0228 — Summary Ranges

```go
package main

// LeetCode #228: Summary Ranges
// https://leetcode.com/problems/summary-ranges/
// Difficulty: Easy

import (
	"fmt"
	"strconv"
)

// Time: O(n) | Space: O(1) excluding output
func SummaryRanges(nums []int) []string {
	var res []string
	i := 0
	for i < len(nums) {
		start := nums[i]
		for i+1 < len(nums) && nums[i+1] == nums[i]+1 {
			i++
		}
		if start == nums[i] {
			res = append(res, strconv.Itoa(start))
		} else {
			res = append(res, strconv.Itoa(start)+"->"+strconv.Itoa(nums[i]))
		}
		i++
	}
	return res
}

func main() {
	fmt.Println(SummaryRanges([]int{0, 1, 2, 4, 5, 7}))
	fmt.Println(SummaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
	fmt.Println(SummaryRanges([]int{}))
}
```

## 0231 — Power Of Two

```go
package main

// LeetCode #231: Power of Two
// https://leetcode.com/problems/power-of-two/
// Difficulty: Easy

import "fmt"

// Time: O(1) | Space: O(1)
func IsPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}

func main() {
	fmt.Println(IsPowerOfTwo(1))
	fmt.Println(IsPowerOfTwo(16))
	fmt.Println(IsPowerOfTwo(3))
}
```

## 0232 — Implement Queue Using Stacks

```go
package main

// LeetCode #232: Implement Queue using Stacks
// https://leetcode.com/problems/implement-queue-using-stacks/
// Difficulty: Easy

import "fmt"

type MyQueue struct {
	in  []int
	out []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (q *MyQueue) Push(x int) {
	q.in = append(q.in, x)
}

func (q *MyQueue) transfer() {
	if len(q.out) == 0 {
		for len(q.in) > 0 {
			q.out = append(q.out, q.in[len(q.in)-1])
			q.in = q.in[:len(q.in)-1]
		}
	}
}

func (q *MyQueue) Pop() int {
	q.transfer()
	x := q.out[len(q.out)-1]
	q.out = q.out[:len(q.out)-1]
	return x
}

func (q *MyQueue) Peek() int {
	q.transfer()
	return q.out[len(q.out)-1]
}

func (q *MyQueue) Empty() bool {
	return len(q.in) == 0 && len(q.out) == 0
}

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	fmt.Println(obj.Peek())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Empty())
}
```

## 0234 — Palindrome Linked List

```go
package main

// LeetCode #234: Palindrome Linked List
// https://leetcode.com/problems/palindrome-linked-list/
// Difficulty: Easy

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time: O(n) | Space: O(1)
func IsPalindrome(head *ListNode) bool {
	reverse := func(head *ListNode) *ListNode {
		var prev *ListNode
		for head != nil {
			next := head.Next
			head.Next = prev
			prev = head
			head = next
		}
		return prev
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	second := reverse(slow)
	first := head
	for second != nil {
		if first.Val != second.Val {
			return false
		}
		first = first.Next
		second = second.Next
	}
	return true
}

func main() {
	l1 := &ListNode{1, &ListNode{2, &ListNode{2, &ListNode{1, nil}}}}
	fmt.Println(IsPalindrome(l1))
	l2 := &ListNode{1, &ListNode{2, nil}}
	fmt.Println(IsPalindrome(l2))
}
```

## 0242 — Valid Anagram

```go
package main

// LeetCode #242: Valid Anagram
// https://leetcode.com/problems/valid-anagram/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(1) (fixed 26 chars)
func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	count := [26]int{}
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
		count[t[i]-'a']--
	}
	return count == [26]int{}
}

func main() {
	fmt.Println(IsAnagram("anagram", "nagaram"))
	fmt.Println(IsAnagram("rat", "car"))
}
```

## 0243 — Shortest Word Distance

```go
package main

// LeetCode #243: Shortest Word Distance
// https://leetcode.com/problems/shortest-word-distance/
// Difficulty: Easy [Paid]

import (
	"fmt"
	"math"
)

// Time: O(n) | Space: O(1)
func ShortestDistance(wordsDict []string, word1 string, word2 string) int {
	i1, i2 := -1, -1
	minDist := math.MaxInt32
	for i, w := range wordsDict {
		if w == word1 {
			i1 = i
		}
		if w == word2 {
			i2 = i
		}
		if i1 != -1 && i2 != -1 {
			dist := i1 - i2
			if dist < 0 {
				dist = -dist
			}
			if dist < minDist {
				minDist = dist
			}
		}
	}
	return minDist
}

func main() {
	fmt.Println(ShortestDistance([]string{"practice", "makes", "perfect", "coding", "makes"}, "coding", "practice"))
	fmt.Println(ShortestDistance([]string{"practice", "makes", "perfect", "coding", "makes"}, "makes", "coding"))
}
```

## 0246 — Strobogrammatic Number

```go
package main

// LeetCode #246: Strobogrammatic Number
// https://leetcode.com/problems/strobogrammatic-number/
// Difficulty: Easy [Paid]

import "fmt"

// Time: O(n) | Space: O(1)
func IsStrobogrammatic(num string) bool {
	pairs := map[byte]byte{'0': '0', '1': '1', '6': '9', '8': '8', '9': '6'}
	i, j := 0, len(num)-1
	for i <= j {
		if v, ok := pairs[num[i]]; !ok || v != num[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	fmt.Println(IsStrobogrammatic("69"))
	fmt.Println(IsStrobogrammatic("88"))
	fmt.Println(IsStrobogrammatic("962"))
}
```

## 0252 — Meeting Rooms

```go
package main

// LeetCode #252: Meeting Rooms
// https://leetcode.com/problems/meeting-rooms/
// Difficulty: Easy [Paid]

import (
	"fmt"
	"sort"
)

// Time: O(n log n) | Space: O(1)
func CanAttendMeetings(intervals [][]int) bool {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][1] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(CanAttendMeetings([][]int{{0, 30}, {5, 10}, {15, 20}}))
	fmt.Println(CanAttendMeetings([][]int{{7, 10}, {2, 4}}))
}
```

## 0257 — Binary Tree Paths

```go
package main

// LeetCode #257: Binary Tree Paths
// https://leetcode.com/problems/binary-tree-paths/
// Difficulty: Easy

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n) | Space: O(h)
func BinaryTreePaths(root *TreeNode) []string {
	var res []string
	var dfs func(*TreeNode, string)
	dfs = func(node *TreeNode, path string) {
		if node == nil {
			return
		}
		path += strconv.Itoa(node.Val)
		if node.Left == nil && node.Right == nil {
			res = append(res, path)
			return
		}
		path += "->"
		dfs(node.Left, path)
		dfs(node.Right, path)
	}
	dfs(root, "")
	return res
}

func main() {
	root := &TreeNode{1, &TreeNode{2, nil, &TreeNode{5, nil, nil}}, &TreeNode{3, nil, nil}}
	fmt.Println(BinaryTreePaths(root))
	fmt.Println(BinaryTreePaths(nil))
}
```

## 0258 — Add Digits

```go
package main

// LeetCode #258: Add Digits
// https://leetcode.com/problems/add-digits/
// Difficulty: Easy

import "fmt"

// Time: O(1) | Space: O(1)
func AddDigits(num int) int {
	if num == 0 {
		return 0
	}
	if num%9 == 0 {
		return 9
	}
	return num % 9
}

func main() {
	fmt.Println(AddDigits(38))
	fmt.Println(AddDigits(0))
	fmt.Println(AddDigits(9))
}
```

## 0263 — Ugly Number

```go
package main

// LeetCode #263: Ugly Number
// https://leetcode.com/problems/ugly-number/
// Difficulty: Easy

import "fmt"

// Time: O(log n) | Space: O(1)
func IsUgly(n int) bool {
	if n <= 0 {
		return false
	}
	for _, f := range []int{2, 3, 5} {
		for n%f == 0 {
			n /= f
		}
	}
	return n == 1
}

func main() {
	fmt.Println(IsUgly(6))
	fmt.Println(IsUgly(1))
	fmt.Println(IsUgly(14))
}
```

## 0266 — Palindrome Permutation

```go
package main

// LeetCode #266: Palindrome Permutation
// https://leetcode.com/problems/palindrome-permutation/
// Difficulty: Easy [Paid]

import "fmt"

// Time: O(n) | Space: O(1) (fixed 256 chars)
func CanPermutePalindrome(s string) bool {
	count := make(map[rune]int)
	for _, c := range s {
		count[c]++
	}
	oddCount := 0
	for _, v := range count {
		if v%2 == 1 {
			oddCount++
		}
	}
	return oddCount <= 1
}

func main() {
	fmt.Println(CanPermutePalindrome("code"))
	fmt.Println(CanPermutePalindrome("aab"))
	fmt.Println(CanPermutePalindrome("carerac"))
}
```

## 0268 — Missing Number

```go
package main

// LeetCode #268: Missing Number
// https://leetcode.com/problems/missing-number/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(1)
func MissingNumber(nums []int) int {
	n := len(nums)
	result := n
	for i, v := range nums {
		result ^= i ^ v
	}
	return result
}

func main() {
	fmt.Println(MissingNumber([]int{3, 0, 1}))
	fmt.Println(MissingNumber([]int{0, 1}))
	fmt.Println(MissingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
}
```

## 0270 — Closest Binary Search Tree Value

```go
package main

// LeetCode #270: Closest Binary Search Tree Value
// https://leetcode.com/problems/closest-binary-search-tree-value/
// Difficulty: Easy [Paid]

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(h) | Space: O(1)
func ClosestValue(root *TreeNode, target float64) int {
	closest := root.Val
	for root != nil {
		if math.Abs(float64(root.Val)-target) < math.Abs(float64(closest)-target) {
			closest = root.Val
		}
		if target < float64(root.Val) {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return closest
}

func main() {
	root := &TreeNode{4, &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}}, &TreeNode{5, nil, nil}}
	fmt.Println(ClosestValue(root, 3.714286))
	fmt.Println(ClosestValue(&TreeNode{1, nil, nil}, 4.428571))
}
```

## 0278 — First Bad Version

```go
package main

// LeetCode #278: First Bad Version
// https://leetcode.com/problems/first-bad-version/
// Difficulty: Easy

import "fmt"

var firstBad int

func isBadVersion(version int) bool {
	return version >= firstBad
}

// Time: O(log n) | Space: O(1)
func FirstBadVersion(n int) int {
	lo, hi := 1, n
	for lo < hi {
		mid := lo + (hi-lo)/2
		if isBadVersion(mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func main() {
	firstBad = 4
	fmt.Println(FirstBadVersion(5))
	firstBad = 1
	fmt.Println(FirstBadVersion(1))
}
```

## 0283 — Move Zeroes

```go
package main

// LeetCode #283: Move Zeroes
// https://leetcode.com/problems/move-zeroes/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(1)
func MoveZeroes(nums []int) {
	lastNonZero := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[lastNonZero] = nums[lastNonZero], nums[i]
			lastNonZero++
		}
	}
}

func main() {
	n1 := []int{0, 1, 0, 3, 12}
	MoveZeroes(n1)
	fmt.Println(n1)
	n2 := []int{0}
	MoveZeroes(n2)
	fmt.Println(n2)
}
```

## 0290 — Word Pattern

```go
package main

// LeetCode #290: Word Pattern
// https://leetcode.com/problems/word-pattern/
// Difficulty: Easy

import "strings"
import "fmt"

// Time: O(n) | Space: O(n)
func WordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}
	p2w := make(map[byte]string)
	w2p := make(map[string]byte)
	for i := 0; i < len(pattern); i++ {
		p := pattern[i]
		w := words[i]
		if mappedW, ok := p2w[p]; ok && mappedW != w {
			return false
		}
		if mappedP, ok := w2p[w]; ok && mappedP != p {
			return false
		}
		p2w[p] = w
		w2p[w] = p
	}
	return true
}

func main() {
	fmt.Println(WordPattern("abba", "dog cat cat dog"))
	fmt.Println(WordPattern("abba", "dog cat cat fish"))
	fmt.Println(WordPattern("aaaa", "dog cat cat dog"))
}
```

## 0292 — Nim Game

```go
package main

// LeetCode #292: Nim Game
// https://leetcode.com/problems/nim-game/
// Difficulty: Easy

import "fmt"

// Time: O(1) | Space: O(1)
func CanWinNim(n int) bool {
	return n%4 != 0
}

func main() {
	fmt.Println(CanWinNim(4))
	fmt.Println(CanWinNim(1))
	fmt.Println(CanWinNim(2))
}
```

## 0293 — Flip Game

```go
package main

// LeetCode #293: Flip Game
// https://leetcode.com/problems/flip-game/
// Difficulty: Easy [Paid]

import "fmt"

// Time: O(n) | Space: O(n) for output
func GeneratePossibleNextMoves(currentState string) []string {
	var res []string
	for i := 0; i < len(currentState)-1; i++ {
		if currentState[i] == '+' && currentState[i+1] == '+' {
			flipped := currentState[:i] + "--" + currentState[i+2:]
			res = append(res, flipped)
		}
	}
	return res
}

func main() {
	fmt.Println(GeneratePossibleNextMoves("++++"))
	fmt.Println(GeneratePossibleNextMoves("+"))
}
```

## 0303 — Range Sum Query Immutable

```go
package main

// LeetCode #303: Range Sum Query - Immutable
// https://leetcode.com/problems/range-sum-query-immutable/
// Difficulty: Easy

import "fmt"

type NumArray struct {
	prefix []int
}

// Time: O(n) for init, O(1) per query | Space: O(n)
func Constructor(nums []int) NumArray {
	prefix := make([]int, len(nums)+1)
	for i, v := range nums {
		prefix[i+1] = prefix[i] + v
	}
	return NumArray{prefix: prefix}
}

func (na *NumArray) SumRange(left int, right int) int {
	return na.prefix[right+1] - na.prefix[left]
}

func main() {
	na := Constructor([]int{-2, 0, 3, -5, 2, -1})
	fmt.Println(na.SumRange(0, 2))
	fmt.Println(na.SumRange(2, 5))
	fmt.Println(na.SumRange(0, 5))
}
```

## 0326 — Power Of Three

```go
package main

// LeetCode #326: Power of Three
// https://leetcode.com/problems/power-of-three/
// Difficulty: Easy

import "fmt"

// Time: O(log n) | Space: O(1)
func IsPowerOfThree(n int) bool {
	if n < 1 {
		return false
	}
	for n%3 == 0 {
		n /= 3
	}
	return n == 1
}

func main() {
	fmt.Println(IsPowerOfThree(27))
	fmt.Println(IsPowerOfThree(0))
	fmt.Println(IsPowerOfThree(-1))
}
```

## 0338 — Counting Bits

```go
package main

// LeetCode #338: Counting Bits
// https://leetcode.com/problems/counting-bits/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(n)
func CountingBits(n int) []int {
	ans := make([]int, n+1)
	for i := 1; i <= n; i++ {
		ans[i] = ans[i>>1] + (i & 1)
	}
	return ans
}

func main() {
	fmt.Println(CountingBits(2))
	fmt.Println(CountingBits(5))
	fmt.Println(CountingBits(0))
}
```

## 0342 — Power Of Four

```go
package main

// LeetCode #342: Power of Four
// https://leetcode.com/problems/power-of-four/
// Difficulty: Easy

import "fmt"

// Time: O(1), Space: O(1)
func PowerOfFour(n int) bool {
	return n > 0 && n&(n-1) == 0 && (n-1)%3 == 0
}

func main() {
	fmt.Println(PowerOfFour(16))
	fmt.Println(PowerOfFour(5))
	fmt.Println(PowerOfFour(1))
}
```

## 0344 — Reverse String

```go
package main

// LeetCode #344: Reverse String
// https://leetcode.com/problems/reverse-string/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func ReverseString(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	s1 := []byte("hello")
	ReverseString(s1)
	fmt.Println(string(s1))

	s2 := []byte("Hannah")
	ReverseString(s2)
	fmt.Println(string(s2))
}
```

## 0345 — Reverse Vowels Of A String

```go
package main

// LeetCode #345: Reverse Vowels of a String
// https://leetcode.com/problems/reverse-vowels-of-a-string/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(n)
func ReverseVowelsOfAString(s string) string {
	b := []byte(s)
	isVowel := func(c byte) bool {
		return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' ||
			c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U'
	}
	for i, j := 0, len(b)-1; i < j; {
		if !isVowel(b[i]) {
			i++
			continue
		}
		if !isVowel(b[j]) {
			j--
			continue
		}
		b[i], b[j] = b[j], b[i]
		i++
		j--
	}
	return string(b)
}

func main() {
	fmt.Println(ReverseVowelsOfAString("hello"))
	fmt.Println(ReverseVowelsOfAString("leetcode"))
	fmt.Println(ReverseVowelsOfAString("aA"))
}
```

## 0346 — Moving Average From Data Stream

```go
package main

// LeetCode #346: Moving Average from Data Stream
// https://leetcode.com/problems/moving-average-from-data-stream/
// Difficulty: Easy [Paid]

import "fmt"

// MovingAverage maintains a sliding window average of the last size values.
type MovingAverage struct {
	size  int
	queue []int
	sum   int
}

// Constructor creates a new MovingAverage with the given window size.
func Constructor(size int) MovingAverage {
	return MovingAverage{size: size}
}

// Next adds a value and returns the moving average.
// Time: O(1), Space: O(n)
func (m *MovingAverage) Next(val int) float64 {
	m.queue = append(m.queue, val)
	m.sum += val
	if len(m.queue) > m.size {
		m.sum -= m.queue[0]
		m.queue = m.queue[1:]
	}
	return float64(m.sum) / float64(len(m.queue))
}

func main() {
	obj := Constructor(3)
	fmt.Println(obj.Next(1))
	fmt.Println(obj.Next(10))
	fmt.Println(obj.Next(3))
	fmt.Println(obj.Next(5))
}
```

## 0349 — Intersection Of Two Arrays

```go
package main

// LeetCode #349: Intersection of Two Arrays
// https://leetcode.com/problems/intersection-of-two-arrays/
// Difficulty: Easy

import "fmt"

// Time: O(n+m), Space: O(n)
func IntersectionOfTwoArrays(nums1, nums2 []int) []int {
	set := make(map[int]bool)
	for _, v := range nums1 {
		set[v] = true
	}
	var result []int
	for _, v := range nums2 {
		if set[v] {
			result = append(result, v)
			delete(set, v)
		}
	}
	return result
}

func main() {
	fmt.Println(IntersectionOfTwoArrays([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(IntersectionOfTwoArrays([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}
```

## 0350 — Intersection Of Two Arrays Ii

```go
package main

// LeetCode #350: Intersection of Two Arrays II
// https://leetcode.com/problems/intersection-of-two-arrays-ii/
// Difficulty: Easy

import "fmt"

// Time: O(n+m), Space: O(min(n,m))
func IntersectionOfTwoArraysIi(nums1, nums2 []int) []int {
	count := make(map[int]int)
	for _, v := range nums1 {
		count[v]++
	}
	var result []int
	for _, v := range nums2 {
		if count[v] > 0 {
			result = append(result, v)
			count[v]--
		}
	}
	return result
}

func main() {
	fmt.Println(IntersectionOfTwoArraysIi([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(IntersectionOfTwoArraysIi([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}
```

## 0359 — Logger Rate Limiter

```go
package main

// LeetCode #359: Logger Rate Limiter
// https://leetcode.com/problems/logger-rate-limiter/
// Difficulty: Easy [Paid]

import "fmt"

// Logger tracks messages and their last printed timestamp.
type Logger struct {
	lastPrinted map[string]int
}

// Constructor creates a new Logger.
func Constructor() Logger {
	return Logger{lastPrinted: make(map[string]int)}
}

// ShouldPrintMessage returns true if the message should be printed at the given timestamp.
// A message can be printed if it hasn't been printed in the last 10 seconds.
// Time: O(1), Space: O(n)
func (l *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	if lastTs, ok := l.lastPrinted[message]; ok && timestamp-lastTs < 10 {
		return false
	}
	l.lastPrinted[message] = timestamp
	return true
}

func main() {
	obj := Constructor()
	fmt.Println(obj.ShouldPrintMessage(1, "foo"))
	fmt.Println(obj.ShouldPrintMessage(2, "bar"))
	fmt.Println(obj.ShouldPrintMessage(3, "foo"))
	fmt.Println(obj.ShouldPrintMessage(8, "bar"))
	fmt.Println(obj.ShouldPrintMessage(10, "foo"))
	fmt.Println(obj.ShouldPrintMessage(11, "foo"))
}
```

## 0367 — Valid Perfect Square

```go
package main

// LeetCode #367: Valid Perfect Square
// https://leetcode.com/problems/valid-perfect-square/
// Difficulty: Easy

import "fmt"

// Time: O(log n), Space: O(1)
func ValidPerfectSquare(num int) bool {
	lo, hi := 1, num
	for lo <= hi {
		mid := lo + (hi-lo)/2
		sq := mid * mid
		if sq == num {
			return true
		} else if sq < num {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return false
}

func main() {
	fmt.Println(ValidPerfectSquare(16))
	fmt.Println(ValidPerfectSquare(14))
	fmt.Println(ValidPerfectSquare(1))
}
```

## 0374 — Guess Number Higher Or Lower

```go
package main

// LeetCode #374: Guess Number Higher or Lower
// https://leetcode.com/problems/guess-number-higher-or-lower/
// Difficulty: Easy

import "fmt"

var pick int

func guess(num int) int {
	if num == pick {
		return 0
	} else if num < pick {
		return 1
	}
	return -1
}

// Time: O(log n), Space: O(1)
func GuessNumberHigherOrLower(n int) int {
	lo, hi := 1, n
	for lo <= hi {
		mid := lo + (hi-lo)/2
		switch guess(mid) {
		case 0:
			return mid
		case 1:
			lo = mid + 1
		case -1:
			hi = mid - 1
		}
	}
	return -1
}

func main() {
	pick = 6
	fmt.Println(GuessNumberHigherOrLower(10))

	pick = 1
	fmt.Println(GuessNumberHigherOrLower(1))

	pick = 1
	fmt.Println(GuessNumberHigherOrLower(2))
}
```

## 0383 — Ransom Note

```go
package main

// LeetCode #383: Ransom Note
// https://leetcode.com/problems/ransom-note/
// Difficulty: Easy

import "fmt"

// Time: O(n+m), Space: O(1)
func RansomNote(ransomNote, magazine string) bool {
	count := [26]int{}
	for _, c := range magazine {
		count[c-'a']++
	}
	for _, c := range ransomNote {
		count[c-'a']--
		if count[c-'a'] < 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(RansomNote("a", "b"))
	fmt.Println(RansomNote("aa", "ab"))
	fmt.Println(RansomNote("aa", "aab"))
}
```

## 0387 — First Unique Character In A String

```go
package main

// LeetCode #387: First Unique Character in a String
// https://leetcode.com/problems/first-unique-character-in-a-string/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func FirstUniqueCharacterInAString(s string) int {
	count := [26]int{}
	for _, c := range s {
		count[c-'a']++
	}
	for i, c := range s {
		if count[c-'a'] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(FirstUniqueCharacterInAString("leetcode"))
	fmt.Println(FirstUniqueCharacterInAString("loveleetcode"))
	fmt.Println(FirstUniqueCharacterInAString("aabb"))
}
```

## 0389 — Find The Difference

```go
package main

// LeetCode #389: Find the Difference
// https://leetcode.com/problems/find-the-difference/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func FindTheDifference(s, t string) byte {
	var diff byte
	for i := 0; i < len(s); i++ {
		diff ^= s[i]
	}
	for i := 0; i < len(t); i++ {
		diff ^= t[i]
	}
	return diff
}

func main() {
	fmt.Printf("%c\n", FindTheDifference("abcd", "abcde"))
	fmt.Printf("%c\n", FindTheDifference("", "y"))
	fmt.Printf("%c\n", FindTheDifference("a", "aa"))
}
```

## 0392 — Is Subsequence

```go
package main

// LeetCode #392: Is Subsequence
// https://leetcode.com/problems/is-subsequence/
// Difficulty: Easy

import "fmt"

// Time: O(n+m), Space: O(1)
func IsSubsequence(s, t string) bool {
	i := 0
	for j := 0; i < len(s) && j < len(t); j++ {
		if s[i] == t[j] {
			i++
		}
	}
	return i == len(s)
}

func main() {
	fmt.Println(IsSubsequence("abc", "ahbgdc"))
	fmt.Println(IsSubsequence("axc", "ahbgdc"))
	fmt.Println(IsSubsequence("", "ahbgdc"))
}
```

## 0401 — Binary Watch

```go
package main

// LeetCode #401: Binary Watch
// https://leetcode.com/problems/binary-watch/
// Difficulty: Easy

import (
	"fmt"
)

func countBits(n int) int {
	count := 0
	for n > 0 {
		n &= n - 1
		count++
	}
	return count
}

// Time: O(1), Space: O(1)
func BinaryWatch(turnedOn int) []string {
	var result []string
	for h := 0; h < 12; h++ {
		for m := 0; m < 60; m++ {
			if countBits(h)+countBits(m) == turnedOn {
				result = append(result, fmt.Sprintf("%d:%02d", h, m))
			}
		}
	}
	return result
}

func main() {
	fmt.Println(BinaryWatch(1))
	fmt.Println(BinaryWatch(9))
}
```

## 0404 — Sum Of Left Leaves

```go
package main

// LeetCode #404: Sum of Left Leaves
// https://leetcode.com/problems/sum-of-left-leaves/
// Difficulty: Easy

import "fmt"

// TreeNode represents a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n), Space: O(h)
func SumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := 0
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		sum += root.Left.Val
	}
	sum += SumOfLeftLeaves(root.Left)
	sum += SumOfLeftLeaves(root.Right)
	return sum
}

func main() {
	// Test case 1: [3,9,20,null,null,15,7]
	root1 := &TreeNode{
		Val: 3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}
	fmt.Println(SumOfLeftLeaves(root1))

	// Test case 2: [1]
	root2 := &TreeNode{Val: 1}
	fmt.Println(SumOfLeftLeaves(root2))
}
```

## 0405 — Convert A Number To Hexadecimal

```go
package main

// LeetCode #405: Convert a Number to Hexadecimal
// https://leetcode.com/problems/convert-a-number-to-hexadecimal/
// Difficulty: Easy

import "fmt"

// Time: O(1), Space: O(1)
func ConvertANumberToHexadecimal(num int) string {
	if num == 0 {
		return "0"
	}
	hex := "0123456789abcdef"
	var result []byte
	// Use uint32 to handle negative numbers via two's complement.
	n := uint32(num)
	for n > 0 {
		result = append([]byte{hex[n&0xf]}, result...)
		n >>= 4
	}
	return string(result)
}

func main() {
	fmt.Println(ConvertANumberToHexadecimal(26))
	fmt.Println(ConvertANumberToHexadecimal(-1))
	fmt.Println(ConvertANumberToHexadecimal(0))
}
```

## 0408 — Valid Word Abbreviation

```go
package main

// LeetCode #408: Valid Word Abbreviation
// https://leetcode.com/problems/valid-word-abbreviation/
// Difficulty: Easy [Paid]

import "fmt"

// Time: O(n), Space: O(1)
func ValidWordAbbreviation(word, abbr string) bool {
	i, j := 0, 0
	for i < len(word) && j < len(abbr) {
		if abbr[j] >= 'a' && abbr[j] <= 'z' {
			if word[i] != abbr[j] {
				return false
			}
			i++
			j++
			continue
		}
		if abbr[j] == '0' {
			return false
		}
		num := 0
		for j < len(abbr) && abbr[j] >= '0' && abbr[j] <= '9' {
			num = num*10 + int(abbr[j]-'0')
			j++
		}
		i += num
	}
	return i == len(word) && j == len(abbr)
}

func main() {
	fmt.Println(ValidWordAbbreviation("internationalization", "i12iz4n"))
	fmt.Println(ValidWordAbbreviation("apple", "a2e"))
	fmt.Println(ValidWordAbbreviation("hi", "1"))
}
```

## 0409 — Longest Palindrome

```go
package main

// LeetCode #409: Longest Palindrome
// https://leetcode.com/problems/longest-palindrome/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func LongestPalindrome(s string) int {
	count := [128]int{}
	for _, c := range s {
		count[c]++
	}
	length, odd := 0, 0
	for _, c := range count {
		length += (c / 2) * 2
		if c%2 == 1 {
			odd = 1
		}
	}
	return length + odd
}

func main() {
	fmt.Println(LongestPalindrome("abccccdd"))
	fmt.Println(LongestPalindrome("a"))
	fmt.Println(LongestPalindrome("bb"))
}
```

## 0412 — Fizz Buzz

```go
package main

// LeetCode #412: Fizz Buzz
// https://leetcode.com/problems/fizz-buzz/
// Difficulty: Easy

import (
	"fmt"
	"strconv"
)

// Time: O(n), Space: O(n)
func FizzBuzz(n int) []string {
	result := make([]string, n)
	for i := 1; i <= n; i++ {
		switch {
		case i%15 == 0:
			result[i-1] = "FizzBuzz"
		case i%3 == 0:
			result[i-1] = "Fizz"
		case i%5 == 0:
			result[i-1] = "Buzz"
		default:
			result[i-1] = strconv.Itoa(i)
		}
	}
	return result
}

func main() {
	fmt.Println(FizzBuzz(3))
	fmt.Println(FizzBuzz(5))
	fmt.Println(FizzBuzz(15))
}
```

