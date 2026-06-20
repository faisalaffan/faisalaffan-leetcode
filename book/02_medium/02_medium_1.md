# Medium (Sedang) — Problem ��0166

## 0002 — Add Two Numbers

```go
package main

// LeetCode #2: Add Two Numbers
// https://leetcode.com/problems/add-two-numbers/
// Difficulty: Medium

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10
		curr.Next = &ListNode{Val: sum % 10}
		curr = curr.Next
	}

	return dummy.Next
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d", head.Val)
		if head.Next != nil {
			fmt.Print(" -> ")
		}
		head = head.Next
	}
	fmt.Println()
}

func main() {
	// Test case 1: l1 = [2,4,3], l2 = [5,6,4] -> [7,0,8]
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	result := addTwoNumbers(l1, l2)
	printList(result)

	// Test case 2: l1 = [0], l2 = [0] -> [0]
	l1 = &ListNode{0, nil}
	l2 = &ListNode{0, nil}
	result = addTwoNumbers(l1, l2)
	printList(result)

	// Test case 3: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9] -> [8,9,9,9,0,0,0,1]
	l1 = &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}}}
	l2 = &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}
	result = addTwoNumbers(l1, l2)
	printList(result)
}

// Time: O(max(m,n)) | Space: O(max(m,n))
```

## 0003 — Longest Substring Without Repeating Characters

```go
package main

// LeetCode #3: Longest Substring Without Repeating Characters
// https://leetcode.com/problems/longest-substring-without-repeating-characters/
// Difficulty: Medium

import "fmt"

func lengthOfLongestSubstring(s string) int {
	lastSeen := make(map[byte]int)
	maxLen := 0
	left := 0

	for right := 0; right < len(s); right++ {
		if idx, ok := lastSeen[s[right]]; ok && idx >= left {
			left = idx + 1
		}
		lastSeen[s[right]] = right
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}

	return maxLen
}

func main() {
	// Test case 1
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) // 3

	// Test case 2
	fmt.Println(lengthOfLongestSubstring("bbbbb")) // 1

	// Test case 3
	fmt.Println(lengthOfLongestSubstring("pwwkew")) // 3
}

// Time: O(n) | Space: O(min(n, alphabet_size))
```

## 0005 — Longest Palindromic Substring

```go
package main

// LeetCode #5: Longest Palindromic Substring
// https://leetcode.com/problems/longest-palindromic-substring/
// Difficulty: Medium

import "fmt"

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	start, maxLen := 0, 1

	expandAroundCenter := func(left, right int) {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if right-left+1 > maxLen {
				start = left
				maxLen = right - left + 1
			}
			left--
			right++
		}
	}

	for i := 0; i < len(s)-1; i++ {
		expandAroundCenter(i, i)   // odd length
		expandAroundCenter(i, i+1) // even length
	}

	return s[start : start+maxLen]
}

func main() {
	// Test case 1
	fmt.Println(longestPalindrome("babad")) // "bab" or "aba"

	// Test case 2
	fmt.Println(longestPalindrome("cbbd")) // "bb"

	// Test case 3
	fmt.Println(longestPalindrome("a")) // "a"
}

// Time: O(n^2) | Space: O(1)
```

## 0006 — Zigzag Conversion

```go
package main

// LeetCode #6: Zigzag Conversion
// https://leetcode.com/problems/zigzag-conversion/
// Difficulty: Medium

import "fmt"

func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}

	rows := make([][]byte, numRows)
	curRow := 0
	goingDown := false

	for i := 0; i < len(s); i++ {
		rows[curRow] = append(rows[curRow], s[i])
		if curRow == 0 || curRow == numRows-1 {
			goingDown = !goingDown
		}
		if goingDown {
			curRow++
		} else {
			curRow--
		}
	}

	result := make([]byte, 0, len(s))
	for _, row := range rows {
		result = append(result, row...)
	}

	return string(result)
}

func main() {
	// Test case 1
	fmt.Println(convert("PAYPALISHIRING", 3)) // "PAHNAPLSIIGYIR"

	// Test case 2
	fmt.Println(convert("PAYPALISHIRING", 4)) // "PINALSIGYAHRPI"

	// Test case 3
	fmt.Println(convert("A", 1)) // "A"
}

// Time: O(n) | Space: O(n)
```

## 0007 — Reverse Integer

```go
package main

// LeetCode #7: Reverse Integer
// https://leetcode.com/problems/reverse-integer/
// Difficulty: Medium

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	result := 0

	for x != 0 {
		digit := x % 10
		x /= 10

		if result > math.MaxInt32/10 || (result == math.MaxInt32/10 && digit > 7) {
			return 0
		}
		if result < math.MinInt32/10 || (result == math.MinInt32/10 && digit < -8) {
			return 0
		}

		result = result*10 + digit
	}

	return result
}

func main() {
	// Test case 1
	fmt.Println(reverse(123)) // 321

	// Test case 2
	fmt.Println(reverse(-123)) // -321

	// Test case 3
	fmt.Println(reverse(1534236469)) // 0 (overflow)
}

// Time: O(log₁₀(n)) | Space: O(1)
```

## 0008 — String To Integer Atoi

```go
package main

// LeetCode #8: String to Integer (atoi)
// https://leetcode.com/problems/string-to-integer-atoi/
// Difficulty: Medium

import (
	"fmt"
	"math"
)

func myAtoi(s string) int {
	i, n := 0, len(s)

	// Skip leading whitespace
	for i < n && s[i] == ' ' {
		i++
	}

	if i == n {
		return 0
	}

	// Handle sign
	sign := 1
	if s[i] == '+' {
		i++
	} else if s[i] == '-' {
		sign = -1
		i++
	}

	result := 0
	for i < n && s[i] >= '0' && s[i] <= '9' {
		digit := int(s[i] - '0')

		if result > math.MaxInt32/10 || (result == math.MaxInt32/10 && digit > 7) {
			if sign == 1 {
				return math.MaxInt32
			}
			return math.MinInt32
		}

		result = result*10 + digit
		i++
	}

	return result * sign
}

func main() {
	// Test case 1
	fmt.Println(myAtoi("42")) // 42

	// Test case 2
	fmt.Println(myAtoi("   -042")) // -42

	// Test case 3
	fmt.Println(myAtoi("1337c0d3")) // 1337

	// Test case 4
	fmt.Println(myAtoi("0-1")) // 0

	// Test case 5
	fmt.Println(myAtoi("words and 987")) // 0
}

// Time: O(n) | Space: O(1)
```

## 0011 — Container With Most Water

```go
package main

// LeetCode #11: Container With Most Water
// https://leetcode.com/problems/container-with-most-water/
// Difficulty: Medium

import "fmt"

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxWater := 0

	for left < right {
		h := height[left]
		if height[right] < h {
			h = height[right]
		}
		area := h * (right - left)
		if area > maxWater {
			maxWater = area
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxWater
}

func main() {
	// Test case 1
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})) // 49

	// Test case 2
	fmt.Println(maxArea([]int{1, 1})) // 1

	// Test case 3
	fmt.Println(maxArea([]int{4, 3, 2, 1, 4})) // 16
}

// Time: O(n) | Space: O(1)
```

## 0012 — Integer To Roman

```go
package main

// LeetCode #12: Integer to Roman
// https://leetcode.com/problems/integer-to-roman/
// Difficulty: Medium

import "fmt"

func intToRoman(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	result := ""
	for i := 0; i < len(values); i++ {
		for num >= values[i] {
			result += symbols[i]
			num -= values[i]
		}
	}

	return result
}

func main() {
	// Test case 1
	fmt.Println(intToRoman(3749)) // "MMMDCCXLIX"

	// Test case 2
	fmt.Println(intToRoman(58)) // "LVIII"

	// Test case 3
	fmt.Println(intToRoman(1994)) // "MCMXCIV"
}

// Time: O(1) | Space: O(1)
```

## 0015 — 3sum

```go
package main

// LeetCode #15: 3Sum
// https://leetcode.com/problems/3sum/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return result
}

func main() {
	// Test case 1
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4})) // [[-1 -1 2] [-1 0 1]]

	// Test case 2
	fmt.Println(threeSum([]int{0, 1, 1})) // []

	// Test case 3
	fmt.Println(threeSum([]int{0, 0, 0})) // [[0 0 0]]
}

// Time: O(n^2) | Space: O(1) (excluding output)
```

## 0016 — 3sum Closest

```go
package main

// LeetCode #16: 3Sum Closest
// https://leetcode.com/problems/3sum-closest/
// Difficulty: Medium

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	closest := nums[0] + nums[1] + nums[2]

	for i := 0; i < len(nums)-2; i++ {
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if math.Abs(float64(sum-target)) < math.Abs(float64(closest-target)) {
				closest = sum
			}
			if sum < target {
				left++
			} else if sum > target {
				right--
			} else {
				return sum
			}
		}
	}

	return closest
}

func main() {
	// Test case 1
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1)) // 2

	// Test case 2
	fmt.Println(threeSumClosest([]int{0, 0, 0}, 1)) // 0

	// Test case 3
	fmt.Println(threeSumClosest([]int{1, 1, 1, 0}, -100)) // 2
}

// Time: O(n^2) | Space: O(1)
```

## 0017 — Letter Combinations Of A Phone Number

```go
package main

// LeetCode #17: Letter Combinations of a Phone Number
// https://leetcode.com/problems/letter-combinations-of-a-phone-number/
// Difficulty: Medium

import "fmt"

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	phone := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	result := []string{""}
	for i := 0; i < len(digits); i++ {
		letters := phone[digits[i]]
		var temp []string
		for _, prefix := range result {
			for j := 0; j < len(letters); j++ {
				temp = append(temp, prefix+string(letters[j]))
			}
		}
		result = temp
	}

	return result
}

func main() {
	// Test case 1
	fmt.Println(letterCombinations("23")) // ["ad","ae","af","bd","be","bf","cd","ce","cf"]

	// Test case 2
	fmt.Println(letterCombinations("")) // []

	// Test case 3
	fmt.Println(letterCombinations("2")) // ["a","b","c"]
}

// Time: O(4^n) | Space: O(4^n)
```

## 0018 — 4sum

```go
package main

// LeetCode #18: 4Sum
// https://leetcode.com/problems/4sum/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	n := len(nums)

	for i := 0; i < n-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			left, right := j+1, n-1
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					left++
					right--
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}

	return result
}

func main() {
	// Test case 1
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0)) // [[-2 -1 1 2] [-2 0 0 2] [-1 0 0 1]]

	// Test case 2
	fmt.Println(fourSum([]int{2, 2, 2, 2, 2}, 8)) // [[2 2 2 2]]
}

// Time: O(n^3) | Space: O(1)
```

## 0019 — Remove Nth Node From End Of List

```go
package main

// LeetCode #19: Remove Nth Node From End of List
// https://leetcode.com/problems/remove-nth-node-from-end-of-list/
// Difficulty: Medium

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	fast, slow := dummy, dummy

	for i := 0; i <= n; i++ {
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return dummy.Next
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d", head.Val)
		if head.Next != nil {
			fmt.Print(" -> ")
		}
		head = head.Next
	}
	fmt.Println()
}

func main() {
	// Test case 1: remove 2nd from end of [1,2,3,4,5] -> [1,2,3,5]
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	result := removeNthFromEnd(head, 2)
	printList(result)

	// Test case 2: remove 1st from end of [1] -> []
	head = &ListNode{1, nil}
	result = removeNthFromEnd(head, 1)
	printList(result)

	// Test case 3: remove 1st from end of [1,2] -> [1]
	head = &ListNode{1, &ListNode{2, nil}}
	result = removeNthFromEnd(head, 1)
	printList(result)
}

// Time: O(n) | Space: O(1)
```

## 0022 — Generate Parentheses

```go
package main

// LeetCode #22: Generate Parentheses
// https://leetcode.com/problems/generate-parentheses/
// Difficulty: Medium

import "fmt"

func generateParenthesis(n int) []string {
	result := []string{}
	var backtrack func(curr string, open, close int)
	backtrack = func(curr string, open, close int) {
		if len(curr) == 2*n {
			result = append(result, curr)
			return
		}
		if open < n {
			backtrack(curr+"(", open+1, close)
		}
		if close < open {
			backtrack(curr+")", open, close+1)
		}
	}
	backtrack("", 0, 0)
	return result
}

func main() {
	// Test case 1
	fmt.Println(generateParenthesis(3)) // ["((()))","(()())","(())()","()(())","()()()"]

	// Test case 2
	fmt.Println(generateParenthesis(1)) // ["()"]

	// Test case 3
	fmt.Println(generateParenthesis(2)) // ["(())","()()"]
}

// Time: O(4^n / sqrt(n)) | Space: O(n)
```

## 0024 — Swap Nodes In Pairs

```go
package main

// LeetCode #24: Swap Nodes in Pairs
// https://leetcode.com/problems/swap-nodes-in-pairs/
// Difficulty: Medium

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy

	for head != nil && head.Next != nil {
		first := head
		second := head.Next

		prev.Next = second
		first.Next = second.Next
		second.Next = first

		prev = first
		head = first.Next
	}

	return dummy.Next
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d", head.Val)
		if head.Next != nil {
			fmt.Print(" -> ")
		}
		head = head.Next
	}
	fmt.Println()
}

func main() {
	// Test case 1: [1,2,3,4] -> [2,1,4,3]
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	result := swapPairs(head)
	printList(result)

	// Test case 2: [] -> []
	result = swapPairs(nil)
	printList(result)

	// Test case 3: [1] -> [1]
	head = &ListNode{1, nil}
	result = swapPairs(head)
	printList(result)
}

// Time: O(n) | Space: O(1)
```

## 0029 — Divide Two Integers

```go
package main

// LeetCode #29: Divide Two Integers
// https://leetcode.com/problems/divide-two-integers/
// Difficulty: Medium

import (
	"fmt"
	"math"
)

func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	negative := (dividend < 0) != (divisor < 0)
	a := abs(dividend)
	b := abs(divisor)
	result := 0

	for a >= b {
		temp := b
		multiple := 1
		for a >= temp<<1 {
			temp <<= 1
			multiple <<= 1
		}
		a -= temp
		result += multiple
	}

	if negative {
		return -result
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	// Test case 1
	fmt.Println(divide(10, 3)) // 3

	// Test case 2
	fmt.Println(divide(7, -3)) // -2

	// Test case 3
	fmt.Println(divide(math.MinInt32, -1)) // 2147483647
}

// Time: O(log n) | Space: O(1)
```

## 0031 — Next Permutation

```go
package main

// LeetCode #31: Next Permutation
// https://leetcode.com/problems/next-permutation/
// Difficulty: Medium

import "fmt"

func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 2

	// Find first decreasing element from right
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i >= 0 {
		j := n - 1
		for nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}

	// Reverse suffix
	for left, right := i+1, n-1; left < right; left, right = left+1, right-1 {
		nums[left], nums[right] = nums[right], nums[left]
	}
}

func main() {
	// Test case 1
	nums := []int{1, 2, 3}
	nextPermutation(nums)
	fmt.Println(nums) // [1 3 2]

	// Test case 2
	nums = []int{3, 2, 1}
	nextPermutation(nums)
	fmt.Println(nums) // [1 2 3]

	// Test case 3
	nums = []int{1, 1, 5}
	nextPermutation(nums)
	fmt.Println(nums) // [1 5 1]
}

// Time: O(n) | Space: O(1)
```

## 0033 — Search In Rotated Sorted Array

```go
package main

// LeetCode #33: Search in Rotated Sorted Array
// https://leetcode.com/problems/search-in-rotated-sorted-array/
// Difficulty: Medium

import "fmt"

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}

		if nums[left] <= nums[mid] {
			// Left half is sorted
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// Right half is sorted
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}

func main() {
	// Test case 1
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0)) // 4

	// Test case 2
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3)) // -1

	// Test case 3
	fmt.Println(search([]int{1}, 0)) // -1
}

// Time: O(log n) | Space: O(1)
```

## 0034 — Find First And Last Position Of Element In Sorted Array

```go
package main

// LeetCode #34: Find First and Last Position of Element in Sorted Array
// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
// Difficulty: Medium

import "fmt"

func searchRange(nums []int, target int) []int {
	result := []int{-1, -1}
	if len(nums) == 0 {
		return result
	}

	// Find first position
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if left < len(nums) && nums[left] == target {
		result[0] = left
	} else {
		return result
	}

	// Find last position
	right = len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	result[1] = right

	return result
}

func main() {
	// Test case 1
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8)) // [3, 4]

	// Test case 2
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6)) // [-1, -1]

	// Test case 3
	fmt.Println(searchRange([]int{}, 0)) // [-1, -1]
}

// Time: O(log n) | Space: O(1)
```

## 0036 — Valid Sudoku

```go
package main

// LeetCode #36: Valid Sudoku
// https://leetcode.com/problems/valid-sudoku/
// Difficulty: Medium

import "fmt"

func isValidSudoku(board [][]byte) bool {
	var rows [9][9]bool
	var cols [9][9]bool
	var boxes [9][9]bool

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			num := board[i][j] - '1'
			boxIdx := (i/3)*3 + j/3

			if rows[i][num] || cols[j][num] || boxes[boxIdx][num] {
				return false
			}
			rows[i][num] = true
			cols[j][num] = true
			boxes[boxIdx][num] = true
		}
	}

	return true
}

func main() {
	// Test case 1: valid board
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
	fmt.Println(isValidSudoku(board)) // true

	// Test case 2: invalid board
	board = [][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Println(isValidSudoku(board)) // false
}

// Time: O(1) | Space: O(1) (board is always 9x9)
```

## 0038 — Count And Say

```go
package main

// LeetCode #38: Count and Say
// https://leetcode.com/problems/count-and-say/
// Difficulty: Medium

import "fmt"

func countAndSay(n int) string {
	curr := "1"

	for i := 2; i <= n; i++ {
		var next []byte
		count := 1
		for j := 1; j < len(curr); j++ {
			if curr[j] == curr[j-1] {
				count++
			} else {
				next = append(next, byte('0'+count), curr[j-1])
				count = 1
			}
		}
		next = append(next, byte('0'+count), curr[len(curr)-1])
		curr = string(next)
	}

	return curr
}

func main() {
	// Test case 1
	fmt.Println(countAndSay(4)) // "1211"
	fmt.Println(countAndSay(1)) // "1"
	fmt.Println(countAndSay(5)) // "111221"
}

// Time: O(2^n) | Space: O(2^n)
```

## 0039 — Combination Sum

```go
package main

// LeetCode #39: Combination Sum
// https://leetcode.com/problems/combination-sum/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result := [][]int{}
	var backtrack func(start int, target int, path []int)
	backtrack = func(start int, target int, path []int) {
		if target == 0 {
			comb := make([]int, len(path))
			copy(comb, path)
			result = append(result, comb)
			return
		}
		for i := start; i < len(candidates); i++ {
			if candidates[i] > target {
				break
			}
			path = append(path, candidates[i])
			backtrack(i, target-candidates[i], path)
			path = path[:len(path)-1]
		}
	}
	backtrack(0, target, []int{})
	return result
}

func main() {
	// Test case 1
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7)) // [[2 2 3] [7]]

	// Test case 2
	fmt.Println(combinationSum([]int{2, 3, 5}, 8)) // [[2 2 2 2] [2 3 3] [3 5]]

	// Test case 3
	fmt.Println(combinationSum([]int{2}, 1)) // []
}

// Time: O(n^(target/min)) | Space: O(target/min)
```

## 0040 — Combination Sum Ii

```go
package main

// LeetCode #40: Combination Sum II
// https://leetcode.com/problems/combination-sum-ii/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result := [][]int{}
	var backtrack func(start int, target int, path []int)
	backtrack = func(start int, target int, path []int) {
		if target == 0 {
			comb := make([]int, len(path))
			copy(comb, path)
			result = append(result, comb)
			return
		}
		for i := start; i < len(candidates); i++ {
			if candidates[i] > target {
				break
			}
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			path = append(path, candidates[i])
			backtrack(i+1, target-candidates[i], path)
			path = path[:len(path)-1]
		}
	}
	backtrack(0, target, []int{})
	return result
}

func main() {
	// Test case 1
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8)) // [[1 1 6] [1 2 5] [1 7] [2 6]]

	// Test case 2
	fmt.Println(combinationSum2([]int{2, 5, 2, 1, 2}, 5)) // [[1 2 2] [5]]
}

// Time: O(2^n) | Space: O(n)
```

## 0043 — Multiply Strings

```go
package main

// LeetCode #43: Multiply Strings
// https://leetcode.com/problems/multiply-strings/
// Difficulty: Medium

import "fmt"

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	m, n := len(num1), len(num2)
	result := make([]byte, m+n)

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			prod := (num1[i]-'0')*(num2[j]-'0') + result[i+j+1]
			result[i+j+1] = prod % 10
			result[i+j] += prod / 10
		}
	}

	if result[0] == 0 {
		result = result[1:]
	}

	for i := range result {
		result[i] += '0'
	}

	return string(result)
}

func main() {
	// Test case 1
	fmt.Println(multiply("2", "3")) // "6"

	// Test case 2
	fmt.Println(multiply("123", "456")) // "56088"

	// Test case 3
	fmt.Println(multiply("999", "999")) // "998001"
}

// Time: O(m*n) | Space: O(m+n)
```

## 0045 — Jump Game Ii

```go
package main

// LeetCode #45: Jump Game II
// https://leetcode.com/problems/jump-game-ii/
// Difficulty: Medium

import "fmt"

func jump(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}

	jumps := 0
	currentEnd := 0
	farthest := 0

	for i := 0; i < n-1; i++ {
		if i+nums[i] > farthest {
			farthest = i + nums[i]
		}
		if i == currentEnd {
			jumps++
			currentEnd = farthest
			if currentEnd >= n-1 {
				break
			}
		}
	}

	return jumps
}

func main() {
	// Test case 1
	fmt.Println(jump([]int{2, 3, 1, 1, 4})) // 2

	// Test case 2
	fmt.Println(jump([]int{2, 3, 0, 1, 4})) // 2

	// Test case 3
	fmt.Println(jump([]int{0})) // 0
}

// Time: O(n) | Space: O(1)
```

## 0046 — Permutations

```go
package main

// LeetCode #46: Permutations
// https://leetcode.com/problems/permutations/
// Difficulty: Medium

import "fmt"

func permute(nums []int) [][]int {
	result := [][]int{}
	n := len(nums)
	var backtrack func(path []int, used []bool)
	backtrack = func(path []int, used []bool) {
		if len(path) == n {
			perm := make([]int, n)
			copy(perm, path)
			result = append(result, perm)
			return
		}
		for i := 0; i < n; i++ {
			if used[i] {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			backtrack(path, used)
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	backtrack([]int{}, make([]bool, n))
	return result
}

func main() {
	// Test case 1
	fmt.Println(permute([]int{1, 2, 3})) // [[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 1 2] [3 2 1]]

	// Test case 2
	fmt.Println(permute([]int{0, 1})) // [[0 1] [1 0]]

	// Test case 3
	fmt.Println(permute([]int{1})) // [[1]]
}

// Time: O(n * n!) | Space: O(n)
```

## 0047 — Permutations Ii

```go
package main

// LeetCode #47: Permutations II
// https://leetcode.com/problems/permutations-ii/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	n := len(nums)
	var backtrack func(path []int, used []bool)
	backtrack = func(path []int, used []bool) {
		if len(path) == n {
			perm := make([]int, n)
			copy(perm, path)
			result = append(result, perm)
			return
		}
		for i := 0; i < n; i++ {
			if used[i] {
				continue
			}
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			backtrack(path, used)
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	backtrack([]int{}, make([]bool, n))
	return result
}

func main() {
	// Test case 1
	fmt.Println(permuteUnique([]int{1, 1, 2})) // [[1 1 2] [1 2 1] [2 1 1]]

	// Test case 2
	fmt.Println(permuteUnique([]int{1, 2, 3})) // [[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 1 2] [3 2 1]]
}

// Time: O(n * n!) | Space: O(n)
```

## 0048 — Rotate Image

```go
package main

// LeetCode #48: Rotate Image
// https://leetcode.com/problems/rotate-image/
// Difficulty: Medium

import "fmt"

func rotate(matrix [][]int) {
	n := len(matrix)

	// Transpose
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// Reverse each row
	for i := 0; i < n; i++ {
		for left, right := 0, n-1; left < right; left, right = left+1, right-1 {
			matrix[i][left], matrix[i][right] = matrix[i][right], matrix[i][left]
		}
	}
}

func main() {
	// Test case 1
	m1 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(m1)
	fmt.Println(m1) // [[7 4 1] [8 5 2] [9 6 3]]

	// Test case 2
	m2 := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	rotate(m2)
	fmt.Println(m2) // [[15 13 2 5] [14 3 4 1] [12 6 8 9] [16 7 10 11]]
}

// Time: O(n^2) | Space: O(1)
```

## 0049 — Group Anagrams

```go
package main

// LeetCode #49: Group Anagrams
// https://leetcode.com/problems/group-anagrams/
// Difficulty: Medium

import "fmt"

func groupAnagrams(strs []string) [][]string {
	groups := make(map[[26]byte][]string)

	for _, s := range strs {
		var key [26]byte
		for i := 0; i < len(s); i++ {
			key[s[i]-'a']++
		}
		groups[key] = append(groups[key], s)
	}

	result := make([][]string, 0, len(groups))
	for _, v := range groups {
		result = append(result, v)
	}

	return result
}

func main() {
	// Test case 1
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	// [["bat"],["nat","tan"],["ate","eat","tea"]]

	// Test case 2
	fmt.Println(groupAnagrams([]string{""})) // [[""]]

	// Test case 3
	fmt.Println(groupAnagrams([]string{"a"})) // [["a"]]
}

// Time: O(n * k) | Space: O(n * k)
```

## 0050 — Powx N

```go
package main

// LeetCode #50: Pow(x, n)
// https://leetcode.com/problems/powx-n/
// Difficulty: Medium

import "fmt"

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}

	result := 1.0
	for n > 0 {
		if n&1 == 1 {
			result *= x
		}
		x *= x
		n >>= 1
	}

	return result
}

func main() {
	// Test case 1
	fmt.Println(myPow(2.0, 10)) // 1024

	// Test case 2
	fmt.Println(myPow(2.1, 3)) // 9.261

	// Test case 3
	fmt.Println(myPow(2.0, -2)) // 0.25
}

// Time: O(log n) | Space: O(1)
```

## 0053 — Maximum Subarray

```go
package main

// LeetCode #53: Maximum Subarray
// https://leetcode.com/problems/maximum-subarray/
// Difficulty: Medium

import "fmt"

func maxSubArray(nums []int) int {
	maxSum := nums[0]
	currSum := nums[0]

	for i := 1; i < len(nums); i++ {
		if currSum+nums[i] > nums[i] {
			currSum = currSum + nums[i]
		} else {
			currSum = nums[i]
		}
		if currSum > maxSum {
			maxSum = currSum
		}
	}

	return maxSum
}

func main() {
	// Test case 1
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6

	// Test case 2
	fmt.Println(maxSubArray([]int{1})) // 1

	// Test case 3
	fmt.Println(maxSubArray([]int{5, 4, -1, 7, 8})) // 23
}

// Time: O(n) | Space: O(1)
```

## 0054 — Spiral Matrix

```go
package main

// LeetCode #54: Spiral Matrix
// https://leetcode.com/problems/spiral-matrix/
// Difficulty: Medium

import "fmt"

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	m, n := len(matrix), len(matrix[0])
	result := make([]int, 0, m*n)
	top, bottom, left, right := 0, m-1, 0, n-1

	for top <= bottom && left <= right {
		for j := left; j <= right; j++ {
			result = append(result, matrix[top][j])
		}
		top++
		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--
		if top <= bottom {
			for j := right; j >= left; j-- {
				result = append(result, matrix[bottom][j])
			}
			bottom--
		}
		if left <= right {
			for i := bottom; i >= top; i-- {
				result = append(result, matrix[i][left])
			}
			left++
		}
	}

	return result
}

func main() {
	// Test case 1
	fmt.Println(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	// [1 2 3 6 9 8 7 4 5]

	// Test case 2
	fmt.Println(spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}))
	// [1 2 3 4 8 12 11 10 9 5 6 7]
}

// Time: O(m*n) | Space: O(1) (excluding output)
```

## 0055 — Jump Game

```go
package main

// LeetCode #55: Jump Game
// https://leetcode.com/problems/jump-game/
// Difficulty: Medium

import "fmt"

func canJump(nums []int) bool {
	reachable := 0
	for i := 0; i < len(nums); i++ {
		if i > reachable {
			return false
		}
		if i+nums[i] > reachable {
			reachable = i + nums[i]
		}
	}
	return true
}

func main() {
	// Test case 1
	fmt.Println(canJump([]int{2, 3, 1, 1, 4})) // true

	// Test case 2
	fmt.Println(canJump([]int{3, 2, 1, 0, 4})) // false

	// Test case 3
	fmt.Println(canJump([]int{0})) // true
}

// Time: O(n) | Space: O(1)
```

## 0056 — Merge Intervals

```go
package main

// LeetCode #56: Merge Intervals
// https://leetcode.com/problems/merge-intervals/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		last := result[len(result)-1]
		if intervals[i][0] <= last[1] {
			if intervals[i][1] > last[1] {
				last[1] = intervals[i][1]
			}
		} else {
			result = append(result, intervals[i])
		}
	}

	return result
}

func main() {
	// Test case 1
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	// [[1 6] [8 10] [15 18]]

	// Test case 2
	fmt.Println(merge([][]int{{1, 4}, {4, 5}})) // [[1 5]]

	// Test case 3
	fmt.Println(merge([][]int{{1, 4}, {2, 3}})) // [[1 4]]
}

// Time: O(n log n) | Space: O(n)
```

## 0057 — Insert Interval

```go
package main

// LeetCode #57: Insert Interval
// https://leetcode.com/problems/insert-interval/
// Difficulty: Medium

import "fmt"

func insert(intervals [][]int, newInterval []int) [][]int {
	result := [][]int{}
	i, n := 0, len(intervals)

	// Add all intervals ending before new interval starts
	for i < n && intervals[i][1] < newInterval[0] {
		result = append(result, intervals[i])
		i++
	}

	// Merge overlapping intervals
	for i < n && intervals[i][0] <= newInterval[1] {
		if intervals[i][0] < newInterval[0] {
			newInterval[0] = intervals[i][0]
		}
		if intervals[i][1] > newInterval[1] {
			newInterval[1] = intervals[i][1]
		}
		i++
	}
	result = append(result, newInterval)

	// Add remaining intervals
	for i < n {
		result = append(result, intervals[i])
		i++
	}

	return result
}

func main() {
	// Test case 1
	fmt.Println(insert([][]int{{1, 3}, {6, 9}}, []int{2, 5}))
	// [[1 5] [6 9]]

	// Test case 2
	fmt.Println(insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}))
	// [[1 2] [3 10] [12 16]]

	// Test case 3
	fmt.Println(insert([][]int{}, []int{5, 7})) // [[5 7]]
}

// Time: O(n) | Space: O(n)
```

## 0059 — Spiral Matrix Ii

```go
package main

// LeetCode #59: Spiral Matrix II
// https://leetcode.com/problems/spiral-matrix-ii/
// Difficulty: Medium

import "fmt"

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	top, bottom, left, right := 0, n-1, 0, n-1
	num := 1

	for top <= bottom && left <= right {
		for j := left; j <= right; j++ {
			matrix[top][j] = num
			num++
		}
		top++
		for i := top; i <= bottom; i++ {
			matrix[i][right] = num
			num++
		}
		right--
		if top <= bottom {
			for j := right; j >= left; j-- {
				matrix[bottom][j] = num
				num++
			}
			bottom--
		}
		if left <= right {
			for i := bottom; i >= top; i-- {
				matrix[i][left] = num
				num++
			}
			left++
		}
	}

	return matrix
}

func main() {
	// Test case 1
	fmt.Println(generateMatrix(3)) // [[1 2 3] [8 9 4] [7 6 5]]

	// Test case 2
	fmt.Println(generateMatrix(1)) // [[1]]

	// Test case 3
	fmt.Println(generateMatrix(4))
}

// Time: O(n^2) | Space: O(n^2)
```

## 0061 — Rotate List

```go
package main

// LeetCode #61: Rotate List
// https://leetcode.com/problems/rotate-list/
// Difficulty: Medium

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	// Find length and tail
	length := 1
	tail := head
	for tail.Next != nil {
		tail = tail.Next
		length++
	}

	k = k % length
	if k == 0 {
		return head
	}

	// Find new head (length - k)th node
	curr := head
	for i := 0; i < length-k-1; i++ {
		curr = curr.Next
	}

	newHead := curr.Next
	curr.Next = nil
	tail.Next = head

	return newHead
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d", head.Val)
		if head.Next != nil {
			fmt.Print(" -> ")
		}
		head = head.Next
	}
	fmt.Println()
}

func main() {
	// Test case 1: [1,2,3,4,5], k=2 -> [4,5,1,2,3]
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	result := rotateRight(head, 2)
	printList(result)

	// Test case 2: [0,1,2], k=4 -> [2,0,1]
	head = &ListNode{0, &ListNode{1, &ListNode{2, nil}}}
	result = rotateRight(head, 4)
	printList(result)

	// Test case 3: [], k=0 -> []
	result = rotateRight(nil, 0)
	printList(result)
}

// Time: O(n) | Space: O(1)
```

## 0062 — Unique Paths

```go
package main

// LeetCode #62: Unique Paths
// https://leetcode.com/problems/unique-paths/
// Difficulty: Medium

import "fmt"

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}

func main() {
	// Test case 1
	fmt.Println(uniquePaths(3, 7)) // 28

	// Test case 2
	fmt.Println(uniquePaths(3, 2)) // 3

	// Test case 3
	fmt.Println(uniquePaths(7, 3)) // 28
}

// Time: O(m*n) | Space: O(m*n)
```

## 0063 — Unique Paths Ii

```go
package main

// LeetCode #63: Unique Paths II
// https://leetcode.com/problems/unique-paths-ii/
// Difficulty: Medium

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])

	if obstacleGrid[0][0] == 1 {
		return 0
	}

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	dp[0][0] = 1
	for j := 1; j < n; j++ {
		if obstacleGrid[0][j] == 0 {
			dp[0][j] = dp[0][j-1]
		}
	}
	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 0 {
			dp[i][0] = dp[i-1][0]
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[m-1][n-1]
}

func main() {
	// Test case 1
	fmt.Println(uniquePathsWithObstacles([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}})) // 2

	// Test case 2
	fmt.Println(uniquePathsWithObstacles([][]int{{0, 1}, {0, 0}})) // 1

	// Test case 3
	fmt.Println(uniquePathsWithObstacles([][]int{{1, 0}})) // 0
}

// Time: O(m*n) | Space: O(m*n)
```

## 0064 — Minimum Path Sum

```go
package main

// LeetCode #64: Minimum Path Sum
// https://leetcode.com/problems/minimum-path-sum/
// Difficulty: Medium

import "fmt"

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	for j := 1; j < n; j++ {
		grid[0][j] += grid[0][j-1]
	}
	for i := 1; i < m; i++ {
		grid[i][0] += grid[i-1][0]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if grid[i-1][j] < grid[i][j-1] {
				grid[i][j] += grid[i-1][j]
			} else {
				grid[i][j] += grid[i][j-1]
			}
		}
	}

	return grid[m-1][n-1]
}

func main() {
	// Test case 1
	fmt.Println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}})) // 7

	// Test case 2
	fmt.Println(minPathSum([][]int{{1, 2, 3}, {4, 5, 6}})) // 12

	// Test case 3
	fmt.Println(minPathSum([][]int{{1}})) // 1
}

// Time: O(m*n) | Space: O(1)
```

## 0071 — Simplify Path

```go
package main

// LeetCode #71: Simplify Path
// https://leetcode.com/problems/simplify-path/
// Difficulty: Medium

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	parts := strings.Split(path, "/")
	stack := []string{}

	for _, part := range parts {
		if part == "" || part == "." {
			continue
		}
		if part == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, part)
		}
	}

	return "/" + strings.Join(stack, "/")
}

func main() {
	// Test case 1
	fmt.Println(simplifyPath("/home/")) // "/home"

	// Test case 2
	fmt.Println(simplifyPath("/home//foo/")) // "/home/foo"

	// Test case 3
	fmt.Println(simplifyPath("/../")) // "/"

	// Test case 4
	fmt.Println(simplifyPath("/a/./b/../../c/")) // "/c"
}

// Time: O(n) | Space: O(n)
```

## 0072 — Edit Distance

```go
package main

// LeetCode #72: Edit Distance
// https://leetcode.com/problems/edit-distance/
// Difficulty: Medium

import "fmt"

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])
			}
		}
	}

	return dp[m][n]
}

func main() {
	// Test case 1
	fmt.Println(minDistance("horse", "ros")) // 3

	// Test case 2
	fmt.Println(minDistance("intention", "execution")) // 5

	// Test case 3
	fmt.Println(minDistance("", "a")) // 1
}

// Time: O(m*n) | Space: O(m*n)
```

## 0073 — Set Matrix Zeroes

```go
package main

// LeetCode #73: Set Matrix Zeroes
// https://leetcode.com/problems/set-matrix-zeroes/
// Difficulty: Medium

import "fmt"

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	firstRowZero := false
	firstColZero := false

	for j := 0; j < n; j++ {
		if matrix[0][j] == 0 {
			firstRowZero = true
			break
		}
	}
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			firstColZero = true
			break
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if firstRowZero {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}
	if firstColZero {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}

func main() {
	// Test case 1
	m1 := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(m1)
	fmt.Println(m1) // [[1 0 1] [0 0 0] [1 0 1]]

	// Test case 2
	m2 := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroes(m2)
	fmt.Println(m2) // [[0 0 0 0] [0 4 5 0] [0 3 1 0]]
}

// Time: O(m*n) | Space: O(1)
```

## 0074 — Search A 2d Matrix

```go
package main

// LeetCode #74: Search a 2D Matrix
// https://leetcode.com/problems/search-a-2d-matrix/
// Difficulty: Medium

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	left, right := 0, m*n-1

	for left <= right {
		mid := left + (right-left)/2
		midVal := matrix[mid/n][mid%n]
		if midVal == target {
			return true
		} else if midVal < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}

func main() {
	// Test case 1
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3)) // true

	// Test case 2
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13)) // false

	// Test case 3
	fmt.Println(searchMatrix([][]int{{1}}, 0)) // false
}

// Time: O(log(m*n)) | Space: O(1)
```

## 0075 — Sort Colors

```go
package main

// LeetCode #75: Sort Colors
// https://leetcode.com/problems/sort-colors/
// Difficulty: Medium

import "fmt"

func sortColors(nums []int) {
	left, right := 0, len(nums)-1
	curr := 0

	for curr <= right {
		if nums[curr] == 0 {
			nums[left], nums[curr] = nums[curr], nums[left]
			left++
			curr++
		} else if nums[curr] == 2 {
			nums[right], nums[curr] = nums[curr], nums[right]
			right--
		} else {
			curr++
		}
	}
}

func main() {
	// Test case 1
	nums1 := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums1)
	fmt.Println(nums1) // [0 0 1 1 2 2]

	// Test case 2
	nums2 := []int{2, 0, 1}
	sortColors(nums2)
	fmt.Println(nums2) // [0 1 2]

	// Test case 3
	nums3 := []int{0}
	sortColors(nums3)
	fmt.Println(nums3) // [0]
}

// Time: O(n) | Space: O(1)
```

## 0077 — Combinations

```go
package main

// LeetCode #77: Combinations
// https://leetcode.com/problems/combinations/
// Difficulty: Medium

import "fmt"

func combine(n int, k int) [][]int {
	result := [][]int{}
	var backtrack func(start int, path []int)
	backtrack = func(start int, path []int) {
		if len(path) == k {
			comb := make([]int, k)
			copy(comb, path)
			result = append(result, comb)
			return
		}
		// Prune: if remaining numbers are not enough, skip
		for i := start; i <= n-(k-len(path))+1; i++ {
			path = append(path, i)
			backtrack(i+1, path)
			path = path[:len(path)-1]
		}
	}
	backtrack(1, []int{})
	return result
}

func main() {
	// Test case 1
	fmt.Println(combine(4, 2)) // [[1 2] [1 3] [1 4] [2 3] [2 4] [3 4]]

	// Test case 2
	fmt.Println(combine(1, 1)) // [[1]]

	// Test case 3
	fmt.Println(combine(4, 4)) // [[1 2 3 4]]
}

// Time: O(C(n,k) * k) | Space: O(k)
```

## 0078 — Subsets

```go
package main

// LeetCode #78: Subsets
// https://leetcode.com/problems/subsets/
// Difficulty: Medium

import "fmt"

func subsets(nums []int) [][]int {
	result := [][]int{{}}
	for _, num := range nums {
		n := len(result)
		for i := 0; i < n; i++ {
			newSubset := make([]int, len(result[i])+1)
			copy(newSubset, result[i])
			newSubset[len(result[i])] = num
			result = append(result, newSubset)
		}
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println(subsets([]int{1, 2, 3}))
	// [[] [1] [2] [1 2] [3] [1 3] [2 3] [1 2 3]]

	// Test case 2
	fmt.Println(subsets([]int{0})) // [[] [0]]

	// Test case 3
	fmt.Println(subsets([]int{})) // [[]]
}

// Time: O(n * 2^n) | Space: O(n * 2^n)
```

## 0079 — Word Search

```go
package main

// LeetCode #79: Word Search
// https://leetcode.com/problems/word-search/
// Difficulty: Medium

import "fmt"

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	var dfs func(i, j, idx int) bool
	dfs = func(i, j, idx int) bool {
		if idx == len(word)-1 {
			return board[i][j] == word[idx]
		}
		if board[i][j] != word[idx] {
			return false
		}

		tmp := board[i][j]
		board[i][j] = '#'

		for _, d := range dirs {
			ni, nj := i+d[0], j+d[1]
			if ni >= 0 && ni < m && nj >= 0 && nj < n {
				if dfs(ni, nj, idx+1) {
					board[i][j] = tmp
					return true
				}
			}
		}

		board[i][j] = tmp
		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}

	return false
}

func main() {
	// Test case 1
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	fmt.Println(exist(board, "ABCCED")) // true

	// Test case 2
	board = [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	fmt.Println(exist(board, "SEE")) // true

	// Test case 3
	board = [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	fmt.Println(exist(board, "ABCB")) // false
}

// Time: O(m*n*4^L) | Space: O(L)
```

## 0080 — Remove Duplicates From Sorted Array Ii

```go
package main

// LeetCode #80: Remove Duplicates from Sorted Array II
// https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/
// Difficulty: Medium

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	write := 2
	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[write-2] {
			nums[write] = nums[i]
			write++
		}
	}

	return write
}

func main() {
	// Test case 1
	nums1 := []int{1, 1, 1, 2, 2, 3}
	k1 := removeDuplicates(nums1)
	fmt.Println(nums1[:k1]) // [1 1 2 2 3]

	// Test case 2
	nums2 := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	k2 := removeDuplicates(nums2)
	fmt.Println(nums2[:k2]) // [0 0 1 1 2 3 3]
}

// Time: O(n) | Space: O(1)
```

## 0081 — Search In Rotated Sorted Array Ii

```go
package main

// LeetCode #81: Search in Rotated Sorted Array II
// https://leetcode.com/problems/search-in-rotated-sorted-array-ii/
// Difficulty: Medium

import "fmt"

func search(nums []int, target int) bool {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return true
		}

		// Handle duplicates: shrink window
		if nums[left] == nums[mid] && nums[mid] == nums[right] {
			left++
			right--
		} else if nums[left] <= nums[mid] {
			// Left half is sorted
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// Right half is sorted
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return false
}

func main() {
	// Test case 1
	fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 0)) // true

	// Test case 2
	fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 3)) // false

	// Test case 3
	fmt.Println(search([]int{1, 0, 1, 1, 1}, 0)) // true
}

// Time: O(log n) average, O(n) worst | Space: O(1)
```

## 0082 — Remove Duplicates From Sorted List Ii

```go
package main

// LeetCode #82: Remove Duplicates from Sorted List II
// https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/
// Difficulty: Medium

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy

	for head != nil {
		if head.Next != nil && head.Val == head.Next.Val {
			for head.Next != nil && head.Val == head.Next.Val {
				head = head.Next
			}
			prev.Next = head.Next
		} else {
			prev = head
		}
		head = head.Next
	}

	return dummy.Next
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d", head.Val)
		if head.Next != nil {
			fmt.Print(" -> ")
		}
		head = head.Next
	}
	fmt.Println()
}

func main() {
	// Test case 1: [1,2,3,3,4,4,5] -> [1,2,5]
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, &ListNode{4, &ListNode{4, &ListNode{5, nil}}}}}}}
	result := deleteDuplicates(head)
	printList(result)

	// Test case 2: [1,1,1,2,3] -> [2,3]
	head = &ListNode{1, &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, nil}}}}}
	result = deleteDuplicates(head)
	printList(result)

	// Test case 3: [1,1] -> []
	head = &ListNode{1, &ListNode{1, nil}}
	result = deleteDuplicates(head)
	printList(result)
}

// Time: O(n) | Space: O(1)
```

## 0086 — Partition List

```go
package main

// LeetCode #86: Partition List
// https://leetcode.com/problems/partition-list/
// Difficulty: Medium

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	lessHead := &ListNode{}
	greaterHead := &ListNode{}
	less, greater := lessHead, greaterHead

	for head != nil {
		if head.Val < x {
			less.Next = head
			less = less.Next
		} else {
			greater.Next = head
			greater = greater.Next
		}
		head = head.Next
	}

	greater.Next = nil
	less.Next = greaterHead.Next

	return lessHead.Next
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d", head.Val)
		if head.Next != nil {
			fmt.Print(" -> ")
		}
		head = head.Next
	}
	fmt.Println()
}

func main() {
	// Test case 1: [1,4,3,2,5,2], x=3 -> [1,2,2,4,3,5]
	head := &ListNode{1, &ListNode{4, &ListNode{3, &ListNode{2, &ListNode{5, &ListNode{2, nil}}}}}}
	result := partition(head, 3)
	printList(result)

	// Test case 2: [2,1], x=2 -> [1,2]
	head = &ListNode{2, &ListNode{1, nil}}
	result = partition(head, 2)
	printList(result)
}

// Time: O(n) | Space: O(1)
```

## 0089 — Gray Code

```go
package main

// LeetCode #89: Gray Code
// https://leetcode.com/problems/gray-code/
// Difficulty: Medium

import "fmt"

func grayCode(n int) []int {
	result := make([]int, 1<<n)
	for i := 0; i < len(result); i++ {
		result[i] = i ^ (i >> 1)
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println(grayCode(2)) // [0 1 3 2]

	// Test case 2
	fmt.Println(grayCode(1)) // [0 1]

	// Test case 3
	fmt.Println(grayCode(3)) // [0 1 3 2 6 7 5 4]
}

// Time: O(2^n) | Space: O(1)
```

## 0090 — Subsets Ii

```go
package main

// LeetCode #90: Subsets II
// https://leetcode.com/problems/subsets-ii/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{{}}
	start := 0

	for i := 0; i < len(nums); i++ {
		n := len(result)
		begin := 0
		if i > 0 && nums[i] == nums[i-1] {
			begin = start
		}
		start = n
		for j := begin; j < n; j++ {
			newSubset := make([]int, len(result[j])+1)
			copy(newSubset, result[j])
			newSubset[len(result[j])] = nums[i]
			result = append(result, newSubset)
		}
	}

	return result
}

func main() {
	// Test case 1
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
	// [[] [1] [2] [1 2] [2 2] [1 2 2]]

	// Test case 2
	fmt.Println(subsetsWithDup([]int{0})) // [[] [0]]
}

// Time: O(n * 2^n) | Space: O(n * 2^n)
```

## 0091 — Decode Ways

```go
package main

// LeetCode #91: Decode Ways
// https://leetcode.com/problems/decode-ways/
// Difficulty: Medium

import "fmt"

func numDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}

	n := len(s)
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		oneDigit := int(s[i-1] - '0')
		if oneDigit >= 1 {
			dp[i] += dp[i-1]
		}

		twoDigits := int(s[i-2]-'0')*10 + oneDigit
		if twoDigits >= 10 && twoDigits <= 26 {
			dp[i] += dp[i-2]
		}
	}

	return dp[n]
}

func main() {
	// Test case 1
	fmt.Println(numDecodings("12")) // 2

	// Test case 2
	fmt.Println(numDecodings("226")) // 3

	// Test case 3
	fmt.Println(numDecodings("06")) // 0
}

// Time: O(n) | Space: O(n)
```

## 0092 — Reverse Linked List Ii

```go
package main

// LeetCode #92: Reverse Linked List II
// https://leetcode.com/problems/reverse-linked-list-ii/
// Difficulty: Medium

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy

	// Move to position left
	for i := 1; i < left; i++ {
		prev = prev.Next
	}

	curr := prev.Next
	var next *ListNode

	// Reverse between left and right
	for i := 0; i < right-left; i++ {
		next = curr.Next
		curr.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
	}

	return dummy.Next
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d", head.Val)
		if head.Next != nil {
			fmt.Print(" -> ")
		}
		head = head.Next
	}
	fmt.Println()
}

func main() {
	// Test case 1: [1,2,3,4,5], left=2, right=4 -> [1,4,3,2,5]
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	result := reverseBetween(head, 2, 4)
	printList(result)

	// Test case 2: [5], left=1, right=1 -> [5]
	head = &ListNode{5, nil}
	result = reverseBetween(head, 1, 1)
	printList(result)

	// Test case 3: [1,2,3,4,5], left=1, right=5 -> [5,4,3,2,1]
	head = &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	result = reverseBetween(head, 1, 5)
	printList(result)
}

// Time: O(n) | Space: O(1)
```

## 0093 — Restore Ip Addresses

```go
package main

// LeetCode #93: Restore IP Addresses
// https://leetcode.com/problems/restore-ip-addresses/
// Difficulty: Medium

import "fmt"

func restoreIpAddresses(s string) []string {
	result := []string{}
	if len(s) < 4 || len(s) > 12 {
		return result
	}

	var backtrack func(start int, parts []string)
	backtrack = func(start int, parts []string) {
		if len(parts) == 4 && start == len(s) {
			ip := parts[0] + "." + parts[1] + "." + parts[2] + "." + parts[3]
			result = append(result, ip)
			return
		}
		if len(parts) == 4 || start == len(s) {
			return
		}

		for i := 1; i <= 3 && start+i <= len(s); i++ {
			segment := s[start : start+i]
			if (len(segment) > 1 && segment[0] == '0') || (i == 3 && segment > "255") {
				continue
			}
			parts = append(parts, segment)
			backtrack(start+i, parts)
			parts = parts[:len(parts)-1]
		}
	}

	backtrack(0, []string{})
	return result
}

func main() {
	// Test case 1
	fmt.Println(restoreIpAddresses("25525511135"))
	// ["255.255.11.135","255.255.111.35"]

	// Test case 2
	fmt.Println(restoreIpAddresses("0000")) // ["0.0.0.0"]

	// Test case 3
	fmt.Println(restoreIpAddresses("101023"))
	// ["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
}

// Time: O(3^4) = O(1) | Space: O(1)
```

## 0095 — Unique Binary Search Trees Ii

```go
package main

// LeetCode #95: Unique Binary Search Trees II
// https://leetcode.com/problems/unique-binary-search-trees-ii/
// Difficulty: Medium

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return build(1, n)
}

func build(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}

	result := []*TreeNode{}
	for i := start; i <= end; i++ {
		leftTrees := build(start, i-1)
		rightTrees := build(i+1, end)
		for _, left := range leftTrees {
			for _, right := range rightTrees {
				result = append(result, &TreeNode{Val: i, Left: left, Right: right})
			}
		}
	}
	return result
}

func printTreePreorder(root *TreeNode) {
	if root == nil {
		fmt.Print("null ")
		return
	}
	fmt.Printf("%d ", root.Val)
	printTreePreorder(root.Left)
	printTreePreorder(root.Right)
}

func main() {
	// Test case 1
	trees := generateTrees(3)
	fmt.Println(len(trees)) // 5
	for _, t := range trees {
		printTreePreorder(t)
		fmt.Println()
	}

	// Test case 2
	trees = generateTrees(1)
	fmt.Println(len(trees)) // 1
	for _, t := range trees {
		printTreePreorder(t)
		fmt.Println()
	}
}

// Time: O(4^n / n^(3/2)) | Space: O(4^n / n^(3/2))
```

## 0096 — Unique Binary Search Trees

```go
package main

// LeetCode #96: Unique Binary Search Trees
// https://leetcode.com/problems/unique-binary-search-trees/
// Difficulty: Medium

import "fmt"

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			dp[i] += dp[j] * dp[i-j-1]
		}
	}

	return dp[n]
}

func main() {
	// Test case 1
	fmt.Println(numTrees(3)) // 5

	// Test case 2
	fmt.Println(numTrees(1)) // 1

	// Test case 3
	fmt.Println(numTrees(4)) // 14
}

// Time: O(n^2) | Space: O(n)
```

## 0097 — Interleaving String

```go
package main

// LeetCode #97: Interleaving String
// https://leetcode.com/problems/interleaving-string/
// Difficulty: Medium

import "fmt"

func isInterleave(s1 string, s2 string, s3 string) bool {
	m, n := len(s1), len(s2)
	if m+n != len(s3) {
		return false
	}

	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}

	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = true
			} else if i == 0 {
				dp[i][j] = dp[i][j-1] && s2[j-1] == s3[i+j-1]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] && s1[i-1] == s3[i+j-1]
			} else {
				dp[i][j] = (dp[i-1][j] && s1[i-1] == s3[i+j-1]) || (dp[i][j-1] && s2[j-1] == s3[i+j-1])
			}
		}
	}

	return dp[m][n]
}

func main() {
	// Test case 1
	fmt.Println(isInterleave("aabcc", "dbbca", "aadbbcbcac")) // true

	// Test case 2
	fmt.Println(isInterleave("aabcc", "dbbca", "aadbbbaccc")) // false

	// Test case 3
	fmt.Println(isInterleave("", "", "")) // true
}

// Time: O(m*n) | Space: O(m*n)
```

## 0098 — Validate Binary Search Tree

```go
package main

// LeetCode #98: Validate Binary Search Tree
// https://leetcode.com/problems/validate-binary-search-tree/
// Difficulty: Medium

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return validate(root, math.MinInt64, math.MaxInt64)
}

func validate(node *TreeNode, min, max int64) bool {
	if node == nil {
		return true
	}
	if int64(node.Val) <= min || int64(node.Val) >= max {
		return false
	}
	return validate(node.Left, min, int64(node.Val)) && validate(node.Right, int64(node.Val), max)
}

func main() {
	// Test case 1: [2,1,3] -> true
	root := &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}
	fmt.Println(isValidBST(root)) // true

	// Test case 2: [5,1,4,null,null,3,6] -> false
	root = &TreeNode{Val: 5, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 6}}}
	fmt.Println(isValidBST(root)) // false

	// Test case 3: [2,2,2] -> false
	root = &TreeNode{Val: 2, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 2}}
	fmt.Println(isValidBST(root)) // false
}

// Time: O(n) | Space: O(n)
```

## 0099 — Recover Binary Search Tree

```go
package main

// LeetCode #99: Recover Binary Search Tree
// https://leetcode.com/problems/recover-binary-search-tree/
// Difficulty: Medium

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	var first, second, prev *TreeNode

	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		if prev != nil && prev.Val > node.Val {
			if first == nil {
				first = prev
			}
			second = node
		}
		prev = node
		inorder(node.Right)
	}

	inorder(root)
	first.Val, second.Val = second.Val, first.Val
}

func printTree(root *TreeNode) {
	if root == nil {
		return
	}
	printTree(root.Left)
	fmt.Printf("%d ", root.Val)
	printTree(root.Right)
}

func main() {
	// Test case 1: [1,3,null,null,2] -> [3,1,null,null,2]
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 3, Right: &TreeNode{Val: 2}}}
	recoverTree(root)
	printTree(root) // 1 2 3
	fmt.Println()

	// Test case 2: [3,1,4,null,null,2] -> [2,1,4,null,null,3]
	root = &TreeNode{Val: 3, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 2}}}
	recoverTree(root)
	printTree(root) // 1 2 3 4
	fmt.Println()
}

// Time: O(n) | Space: O(n)
```

## 0102 — Binary Tree Level Order Traversal

```go
package main

// LeetCode #102: Binary Tree Level Order Traversal
// https://leetcode.com/problems/binary-tree-level-order-traversal/
// Difficulty: Medium

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		level := make([]int, levelSize)
		for i := 0; i < levelSize; i++ {
			node := queue[i]
			level[i] = node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, level)
		queue = queue[levelSize:]
	}

	return result
}

func main() {
	// Test case 1: [3,9,20,null,null,15,7] -> [[3],[9,20],[15,7]]
	root := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	fmt.Println(levelOrder(root))

	// Test case 2
	fmt.Println(levelOrder(nil)) // []

	// Test case 3: [1] -> [[1]]
	fmt.Println(levelOrder(&TreeNode{Val: 1})) // [[1]]
}

// Time: O(n) | Space: O(n)
```

## 0103 — Binary Tree Zigzag Level Order Traversal

```go
package main

// LeetCode #103: Binary Tree Zigzag Level Order Traversal
// https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/
// Difficulty: Medium

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}
	leftToRight := true

	for len(queue) > 0 {
		levelSize := len(queue)
		level := make([]int, levelSize)
		for i := 0; i < levelSize; i++ {
			node := queue[i]
			if leftToRight {
				level[i] = node.Val
			} else {
				level[levelSize-1-i] = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, level)
		queue = queue[levelSize:]
		leftToRight = !leftToRight
	}

	return result
}

func main() {
	// Test case 1: [3,9,20,null,null,15,7] -> [[3],[20,9],[15,7]]
	root := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	fmt.Println(zigzagLevelOrder(root))

	// Test case 2
	fmt.Println(zigzagLevelOrder(nil)) // []

	// Test case 3: [1] -> [[1]]
	fmt.Println(zigzagLevelOrder(&TreeNode{Val: 1})) // [[1]]
}

// Time: O(n) | Space: O(n)
```

## 0105 — Construct Binary Tree From Preorder And Inorder Traversal

```go
package main

// LeetCode #105: Construct Binary Tree from Preorder and Inorder Traversal
// https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
// Difficulty: Medium

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	inorderMap := make(map[int]int)
	for i, v := range inorder {
		inorderMap[v] = i
	}

	var build func(preStart, preEnd, inStart, inEnd int) *TreeNode
	build = func(preStart, preEnd, inStart, inEnd int) *TreeNode {
		if preStart > preEnd || inStart > inEnd {
			return nil
		}

		rootVal := preorder[preStart]
		root := &TreeNode{Val: rootVal}
		inIdx := inorderMap[rootVal]
		leftSize := inIdx - inStart

		root.Left = build(preStart+1, preStart+leftSize, inStart, inIdx-1)
		root.Right = build(preStart+leftSize+1, preEnd, inIdx+1, inEnd)
		return root
	}

	return build(0, len(preorder)-1, 0, len(inorder)-1)
}

func printInorder(root *TreeNode) {
	if root == nil {
		return
	}
	printInorder(root.Left)
	fmt.Printf("%d ", root.Val)
	printInorder(root.Right)
}

func main() {
	// Test case 1
	root := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	printInorder(root) // 9 3 15 20 7
	fmt.Println()

	// Test case 2
	root = buildTree([]int{-1}, []int{-1})
	printInorder(root) // -1
	fmt.Println()
}

// Time: O(n) | Space: O(n)
```

## 0106 — Construct Binary Tree From Inorder And Postorder Traversal

```go
package main

// LeetCode #106: Construct Binary Tree from Inorder and Postorder Traversal
// https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
// Difficulty: Medium

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	inorderMap := make(map[int]int)
	for i, v := range inorder {
		inorderMap[v] = i
	}

	var build func(inStart, inEnd, postStart, postEnd int) *TreeNode
	build = func(inStart, inEnd, postStart, postEnd int) *TreeNode {
		if inStart > inEnd || postStart > postEnd {
			return nil
		}

		rootVal := postorder[postEnd]
		root := &TreeNode{Val: rootVal}
		inIdx := inorderMap[rootVal]
		rightSize := inEnd - inIdx

		root.Left = build(inStart, inIdx-1, postStart, postEnd-rightSize-1)
		root.Right = build(inIdx+1, inEnd, postEnd-rightSize, postEnd-1)
		return root
	}

	return build(0, len(inorder)-1, 0, len(postorder)-1)
}

func printInorder(root *TreeNode) {
	if root == nil {
		return
	}
	printInorder(root.Left)
	fmt.Printf("%d ", root.Val)
	printInorder(root.Right)
}

func main() {
	// Test case 1
	root := buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})
	printInorder(root) // 9 3 15 20 7
	fmt.Println()

	// Test case 2
	root = buildTree([]int{-1}, []int{-1})
	printInorder(root) // -1
	fmt.Println()
}

// Time: O(n) | Space: O(n)
```

## 0107 — Binary Tree Level Order Traversal Ii

```go
package main

// LeetCode #107: Binary Tree Level Order Traversal II
// https://leetcode.com/problems/binary-tree-level-order-traversal-ii/
// Difficulty: Medium

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		level := make([]int, levelSize)
		for i := 0; i < levelSize; i++ {
			node := queue[i]
			level[i] = node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append([][]int{level}, result...)
		queue = queue[levelSize:]
	}

	return result
}

func main() {
	// Test case 1: [3,9,20,null,null,15,7] -> [[15,7],[9,20],[3]]
	root := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	fmt.Println(levelOrderBottom(root))

	// Test case 2
	fmt.Println(levelOrderBottom(nil)) // []

	// Test case 3: [1] -> [[1]]
	fmt.Println(levelOrderBottom(&TreeNode{Val: 1})) // [[1]]
}

// Time: O(n) | Space: O(n)
```

## 0109 — Convert Sorted List To Binary Search Tree

```go
package main

// LeetCode #109: Convert Sorted List to Binary Search Tree
// https://leetcode.com/problems/convert-sorted-list-to-binary-search-tree/
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

func sortedListToBST(head *ListNode) *TreeNode {
	nums := []int{}
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}

	var build func(left, right int) *TreeNode
	build = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}
		mid := left + (right-left)/2
		node := &TreeNode{Val: nums[mid]}
		node.Left = build(left, mid-1)
		node.Right = build(mid+1, right)
		return node
	}

	return build(0, len(nums)-1)
}

func printInorder(root *TreeNode) {
	if root == nil {
		return
	}
	printInorder(root.Left)
	fmt.Printf("%d ", root.Val)
	printInorder(root.Right)
}

func main() {
	// Test case 1: [-10,-3,0,5,9]
	head := &ListNode{-10, &ListNode{-3, &ListNode{0, &ListNode{5, &ListNode{9, nil}}}}}
	root := sortedListToBST(head)
	printInorder(root) // -10 -3 0 5 9
	fmt.Println()

	// Test case 2: [] -> nil
	root = sortedListToBST(nil)
	printInorder(root)
	fmt.Println()
}

// Time: O(n) | Space: O(n)
```

## 0113 — Path Sum Ii

```go
package main

// LeetCode #113: Path Sum II
// https://leetcode.com/problems/path-sum-ii/
// Difficulty: Medium

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	result := [][]int{}
	var dfs func(node *TreeNode, remaining int, path []int)
	dfs = func(node *TreeNode, remaining int, path []int) {
		if node == nil {
			return
		}
		remaining -= node.Val
		path = append(path, node.Val)

		if node.Left == nil && node.Right == nil && remaining == 0 {
			validPath := make([]int, len(path))
			copy(validPath, path)
			result = append(result, validPath)
		} else {
			dfs(node.Left, remaining, path)
			dfs(node.Right, remaining, path)
		}

		path = path[:len(path)-1]
	}
	dfs(root, targetSum, []int{})
	return result
}

func main() {
	// Test case 1
	root := &TreeNode{Val: 5,
		Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 11, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 2}}},
		Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 13}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 1}}}}
	fmt.Println(pathSum(root, 22)) // [[5 4 11 2] [5 8 4 5]]

	// Test case 2
	fmt.Println(pathSum(nil, 0)) // []

	// Test case 3: [1,2], target=1 -> []
	root = &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	fmt.Println(pathSum(root, 1)) // []
}

// Time: O(n^2) | Space: O(n)
```

## 0114 — Flatten Binary Tree To Linked List

```go
package main

// LeetCode #114: Flatten Binary Tree to Linked List
// https://leetcode.com/problems/flatten-binary-tree-to-linked-list/
// Difficulty: Medium

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	curr := root
	for curr != nil {
		if curr.Left != nil {
			// Find rightmost node in left subtree
			prev := curr.Left
			for prev.Right != nil {
				prev = prev.Right
			}
			// Rewire
			prev.Right = curr.Right
			curr.Right = curr.Left
			curr.Left = nil
		}
		curr = curr.Right
	}
}

func printPreorder(root *TreeNode) {
	for root != nil {
		fmt.Printf("%d ", root.Val)
		if root.Left != nil {
			fmt.Print("(has left) ")
		}
		root = root.Right
	}
	fmt.Println()
}

func main() {
	// Test case 1: [1,2,5,3,4,null,6] -> [1,null,2,null,3,null,4,null,5,null,6]
	root := &TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 6}}}
	flatten(root)
	printPreorder(root) // 1 2 3 4 5 6

	// Test case 2
	root = nil
	flatten(root)
	printPreorder(root)

	// Test case 3: [0]
	root = &TreeNode{Val: 0}
	flatten(root)
	printPreorder(root) // 0
}

// Time: O(n) | Space: O(1)
```

## 0116 — Populating Next Right Pointers In Each Node

```go
package main

// LeetCode #116: Populating Next Right Pointers in Each Node
// https://leetcode.com/problems/populating-next-right-pointers-in-each-node/
// Difficulty: Medium

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	leftmost := root
	for leftmost.Left != nil {
		head := leftmost
		for head != nil {
			head.Left.Next = head.Right
			if head.Next != nil {
				head.Right.Next = head.Next.Left
			}
			head = head.Next
		}
		leftmost = leftmost.Left
	}

	return root
}

func printLevels(root *Node) {
	curr := root
	for curr != nil {
		head := curr
		for head != nil {
			nextStr := "null"
			if head.Next != nil {
				nextStr = fmt.Sprintf("%d", head.Next.Val)
			}
			fmt.Printf("%d->%s ", head.Val, nextStr)
			head = head.Next
		}
		fmt.Println()
		curr = curr.Left
	}
}

func main() {
	// Test case 1: [1,2,3,4,5,6,7]
	root := &Node{Val: 1,
		Left:  &Node{Val: 2, Left: &Node{Val: 4}, Right: &Node{Val: 5}},
		Right: &Node{Val: 3, Left: &Node{Val: 6}, Right: &Node{Val: 7}}}
	connect(root)
	printLevels(root)

	// Test case 2
	root = nil
	connect(root)
	fmt.Println("nil")
}

// Time: O(n) | Space: O(1)
```

## 0117 — Populating Next Right Pointers In Each Node Ii

```go
package main

// LeetCode #117: Populating Next Right Pointers in Each Node II
// https://leetcode.com/problems/populating-next-right-pointers-in-each-node-ii/
// Difficulty: Medium

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	curr := root
	for curr != nil {
		dummy := &Node{}
		tail := dummy

		for curr != nil {
			if curr.Left != nil {
				tail.Next = curr.Left
				tail = tail.Next
			}
			if curr.Right != nil {
				tail.Next = curr.Right
				tail = tail.Next
			}
			curr = curr.Next
		}

		curr = dummy.Next
	}

	return root
}

func printLevels(root *Node) {
	curr := root
	for curr != nil {
		head := curr
		for head != nil {
			nextStr := "null"
			if head.Next != nil {
				nextStr = fmt.Sprintf("%d", head.Next.Val)
			}
			fmt.Printf("%d->%s ", head.Val, nextStr)
			head = head.Next
		}
		fmt.Println()
		curr = curr.Left
		if curr == nil {
			// Follow Next for non-perfect trees
			curr = root.Next
			for curr != nil && curr.Left == nil && curr.Right == nil {
				curr = curr.Next
			}
		}
	}
}

func main() {
	// Test case 1: [1,2,3,4,5,null,7]
	root := &Node{Val: 1,
		Left:  &Node{Val: 2, Left: &Node{Val: 4}, Right: &Node{Val: 5}},
		Right: &Node{Val: 3, Right: &Node{Val: 7}}}
	connect(root)
	printLevels(root)

	// Test case 2
	root = nil
	connect(root)
	fmt.Println("nil")
}

// Time: O(n) | Space: O(1)
```

## 0120 — Triangle

```go
package main

// LeetCode #120: Triangle
// https://leetcode.com/problems/triangle/
// Difficulty: Medium

import "fmt"

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := make([]int, n)
	copy(dp, triangle[n-1])

	for i := n - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			if dp[j] < dp[j+1] {
				dp[j] = triangle[i][j] + dp[j]
			} else {
				dp[j] = triangle[i][j] + dp[j+1]
			}
		}
	}

	return dp[0]
}

func main() {
	// Test case 1
	fmt.Println(minimumTotal([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}})) // 11

	// Test case 2
	fmt.Println(minimumTotal([][]int{{-10}})) // -10

	// Test case 3
	fmt.Println(minimumTotal([][]int{{1}, {2, 3}})) // 3
}

// Time: O(n^2) | Space: O(n)
```

## 0122 — Best Time To Buy And Sell Stock Ii

```go
package main

// LeetCode #122: Best Time to Buy and Sell Stock II
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/
// Difficulty: Medium

import "fmt"

func maxProfit(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}

func main() {
	// Test case 1
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4})) // 7

	// Test case 2
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5})) // 4

	// Test case 3
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1})) // 0
}

// Time: O(n) | Space: O(1)
```

## 0128 — Longest Consecutive Sequence

```go
package main

// LeetCode #128: Longest Consecutive Sequence
// https://leetcode.com/problems/longest-consecutive-sequence/
// Difficulty: Medium

import "fmt"

func longestConsecutive(nums []int) int {
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	maxLen := 0
	for num := range numSet {
		if !numSet[num-1] {
			curr := num
			length := 1
			for numSet[curr+1] {
				curr++
				length++
			}
			if length > maxLen {
				maxLen = length
			}
		}
	}

	return maxLen
}

func main() {
	// Test case 1
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2})) // 4

	// Test case 2
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1})) // 9

	// Test case 3
	fmt.Println(longestConsecutive([]int{})) // 0
}

// Time: O(n) | Space: O(n)
```

## 0129 — Sum Root To Leaf Numbers

```go
package main

// LeetCode #129: Sum Root to Leaf Numbers
// https://leetcode.com/problems/sum-root-to-leaf-numbers/
// Difficulty: Medium

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	var dfs func(node *TreeNode, curr int) int
	dfs = func(node *TreeNode, curr int) int {
		if node == nil {
			return 0
		}
		curr = curr*10 + node.Val
		if node.Left == nil && node.Right == nil {
			return curr
		}
		return dfs(node.Left, curr) + dfs(node.Right, curr)
	}
	return dfs(root, 0)
}

func main() {
	// Test case 1: [1,2,3] -> 12+13=25
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	fmt.Println(sumNumbers(root)) // 25

	// Test case 2: [4,9,0,5,1] -> 495+491+40=1026
	root = &TreeNode{Val: 4,
		Left:  &TreeNode{Val: 9, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 1}},
		Right: &TreeNode{Val: 0}}
	fmt.Println(sumNumbers(root)) // 1026

	// Test case 3
	fmt.Println(sumNumbers(nil)) // 0
}

// Time: O(n) | Space: O(n)
```

## 0130 — Surrounded Regions

```go
package main

// LeetCode #130: Surrounded Regions
// https://leetcode.com/problems/surrounded-regions/
// Difficulty: Medium

import "fmt"

func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	m, n := len(board), len(board[0])

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || board[i][j] != 'O' {
			return
		}
		board[i][j] = '#'
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}

	// Mark border 'O' cells
	for i := 0; i < m; i++ {
		if board[i][0] == 'O' {
			dfs(i, 0)
		}
		if board[i][n-1] == 'O' {
			dfs(i, n-1)
		}
	}
	for j := 0; j < n; j++ {
		if board[0][j] == 'O' {
			dfs(0, j)
		}
		if board[m-1][j] == 'O' {
			dfs(m-1, j)
		}
	}

	// Flip remaining 'O' to 'X' and '#' back to 'O'
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}
}

func main() {
	// Test case 1
	board := [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}}
	solve(board)
	fmt.Println(board) // [[X X X X] [X X X X] [X X X X] [X O X X]]

	// Test case 2
	board = [][]byte{{'X'}}
	solve(board)
	fmt.Println(board) // [[X]]
}

// Time: O(m*n) | Space: O(m*n)
```

## 0131 — Palindrome Partitioning

```go
package main

// LeetCode #131: Palindrome Partitioning
// https://leetcode.com/problems/palindrome-partitioning/
// Difficulty: Medium

import "fmt"

func partition(s string) [][]string {
	result := [][]string{}
	n := len(s)

	// Precompute palindrome table
	pal := make([][]bool, n)
	for i := range pal {
		pal[i] = make([]bool, n)
	}
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] == s[j] && (j-i <= 2 || pal[i+1][j-1]) {
				pal[i][j] = true
			}
		}
	}

	var backtrack func(start int, path []string)
	backtrack = func(start int, path []string) {
		if start == n {
			part := make([]string, len(path))
			copy(part, path)
			result = append(result, part)
			return
		}
		for end := start; end < n; end++ {
			if pal[start][end] {
				path = append(path, s[start:end+1])
				backtrack(end+1, path)
				path = path[:len(path)-1]
			}
		}
	}

	backtrack(0, []string{})
	return result
}

func main() {
	// Test case 1
	fmt.Println(partition("aab")) // [["a","a","b"],["aa","b"]]

	// Test case 2
	fmt.Println(partition("a")) // [["a"]]

	// Test case 3
	fmt.Println(partition("ab")) // [["a","b"]]
}

// Time: O(n * 2^n) | Space: O(n^2)
```

## 0133 — Clone Graph

```go
package main

// LeetCode #133: Clone Graph
// https://leetcode.com/problems/clone-graph/
// Difficulty: Medium

import "fmt"

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	visited := make(map[*Node]*Node)

	var dfs func(n *Node) *Node
	dfs = func(n *Node) *Node {
		if clone, ok := visited[n]; ok {
			return clone
		}

		clone := &Node{Val: n.Val, Neighbors: make([]*Node, len(n.Neighbors))}
		visited[n] = clone

		for i, neighbor := range n.Neighbors {
			clone.Neighbors[i] = dfs(neighbor)
		}

		return clone
	}

	return dfs(node)
}

func printGraph(node *Node, visited map[*Node]bool) {
	if node == nil || visited[node] {
		return
	}
	visited[node] = true
	fmt.Printf("Node %d: [", node.Val)
	for i, n := range node.Neighbors {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(n.Val)
	}
	fmt.Println("]")
	for _, n := range node.Neighbors {
		printGraph(n, visited)
	}
}

func main() {
	// Test case 1: [[2,4],[1,3],[2,4],[1,3]]
	n1 := &Node{Val: 1}
	n2 := &Node{Val: 2}
	n3 := &Node{Val: 3}
	n4 := &Node{Val: 4}
	n1.Neighbors = []*Node{n2, n4}
	n2.Neighbors = []*Node{n1, n3}
	n3.Neighbors = []*Node{n2, n4}
	n4.Neighbors = []*Node{n1, n3}

	clone := cloneGraph(n1)
	printGraph(clone, make(map[*Node]bool))

	// Test case 2
	fmt.Println(cloneGraph(nil)) // nil
}

// Time: O(V+E) | Space: O(V)
```

## 0134 — Gas Station

```go
package main

// LeetCode #134: Gas Station
// https://leetcode.com/problems/gas-station/
// Difficulty: Medium

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	totalGas, totalCost := 0, 0
	for i := 0; i < len(gas); i++ {
		totalGas += gas[i]
		totalCost += cost[i]
	}

	if totalGas < totalCost {
		return -1
	}

	start, tank := 0, 0
	for i := 0; i < len(gas); i++ {
		tank += gas[i] - cost[i]
		if tank < 0 {
			start = i + 1
			tank = 0
		}
	}

	return start
}

func main() {
	// Test case 1
	fmt.Println(canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2})) // 3

	// Test case 2
	fmt.Println(canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3})) // -1

	// Test case 3
	fmt.Println(canCompleteCircuit([]int{5, 1, 2, 3, 4}, []int{4, 4, 1, 5, 1})) // 4
}

// Time: O(n) | Space: O(1)
```

## 0137 — Single Number Ii

```go
package main

// LeetCode #137: Single Number II
// https://leetcode.com/problems/single-number-ii/
// Difficulty: Medium

import "fmt"

func singleNumber(nums []int) int {
	ones, twos := 0, 0
	for _, num := range nums {
		ones = (ones ^ num) & ^twos
		twos = (twos ^ num) & ^ones
	}
	return ones
}

func main() {
	// Test case 1
	fmt.Println(singleNumber([]int{2, 2, 3, 2})) // 3

	// Test case 2
	fmt.Println(singleNumber([]int{0, 1, 0, 1, 0, 1, 99})) // 99

	// Test case 3
	fmt.Println(singleNumber([]int{30000, 500, 100, 30000, 100, 30000, 100})) // 500
}

// Time: O(n) | Space: O(1)
```

## 0138 — Copy List With Random Pointer

```go
package main

// LeetCode #138: Copy List with Random Pointer
// https://leetcode.com/problems/copy-list-with-random-pointer/
// Difficulty: Medium

import "fmt"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	// Interleave original nodes with copies
	curr := head
	for curr != nil {
		copy := &Node{Val: curr.Val}
		copy.Next = curr.Next
		curr.Next = copy
		curr = copy.Next
	}

	// Set random pointers for copies
	curr = head
	for curr != nil {
		if curr.Random != nil {
			curr.Next.Random = curr.Random.Next
		}
		curr = curr.Next.Next
	}

	// Separate original and copy lists
	dummy := &Node{}
	copyCurr := dummy
	curr = head
	for curr != nil {
		copyCurr.Next = curr.Next
		copyCurr = copyCurr.Next
		curr.Next = curr.Next.Next
		curr = curr.Next
	}

	return dummy.Next
}

func printList(head *Node) {
	for head != nil {
		randomStr := "nil"
		if head.Random != nil {
			randomStr = fmt.Sprintf("%d", head.Random.Val)
		}
		fmt.Printf("[%d, %s] ", head.Val, randomStr)
		head = head.Next
	}
	fmt.Println()
}

func main() {
	// Test case 1: [[7,null],[13,0],[11,4],[10,2],[1,0]]
	n0 := &Node{Val: 7}
	n1 := &Node{Val: 13}
	n2 := &Node{Val: 11}
	n3 := &Node{Val: 10}
	n4 := &Node{Val: 1}
	n0.Next, n0.Random = n1, nil
	n1.Next, n1.Random = n2, n0
	n2.Next, n2.Random = n3, n4
	n3.Next, n3.Random = n4, n2
	n4.Next, n4.Random = nil, n0

	copy := copyRandomList(n0)
	printList(copy)

	// Test case 2
	fmt.Println(copyRandomList(nil)) // nil
}

// Time: O(n) | Space: O(1)
```

## 0139 — Word Break

```go
package main

// LeetCode #139: Word Break
// https://leetcode.com/problems/word-break/
// Difficulty: Medium

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	wordSet := make(map[string]bool)
	for _, w := range wordDict {
		wordSet[w] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}

func main() {
	// Test case 1
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"})) // true

	// Test case 2
	fmt.Println(wordBreak("applepenapple", []string{"apple", "pen"})) // true

	// Test case 3
	fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"})) // false
}

// Time: O(n^2) | Space: O(n)
```

## 0142 — Linked List Cycle Ii

```go
package main

// LeetCode #142: Linked List Cycle II
// https://leetcode.com/problems/linked-list-cycle-ii/
// Difficulty: Medium

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			slow = head
			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}
			return slow
		}
	}

	return nil
}

func main() {
	// Test case 1: [3,2,0,-4], pos=1
	n0 := &ListNode{Val: 3}
	n1 := &ListNode{Val: 2}
	n2 := &ListNode{Val: 0}
	n3 := &ListNode{Val: -4}
	n0.Next = n1
	n1.Next = n2
	n2.Next = n3
	n3.Next = n1 // cycle
	result := detectCycle(n0)
	if result != nil {
		fmt.Println(result.Val) // 2
	}

	// Test case 2: no cycle
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2}}
	result = detectCycle(head)
	fmt.Println(result) // nil

	// Test case 3: single node, no cycle
	result = detectCycle(&ListNode{Val: 1})
	fmt.Println(result) // nil
}

// Time: O(n) | Space: O(1)
```

## 0143 — Reorder List

```go
package main

// LeetCode #143: Reorder List
// https://leetcode.com/problems/reorder-list/
// Difficulty: Medium

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// Find middle
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Reverse second half
	var prev *ListNode
	curr := slow.Next
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	slow.Next = nil

	// Merge two halves
	first, second := head, prev
	for first != nil && second != nil {
		firstNext := first.Next
		secondNext := second.Next
		first.Next = second
		second.Next = firstNext
		first = firstNext
		second = secondNext
	}
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d", head.Val)
		if head.Next != nil {
			fmt.Print(" -> ")
		}
		head = head.Next
	}
	fmt.Println()
}

func main() {
	// Test case 1: [1,2,3,4] -> [1,4,2,3]
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	reorderList(head)
	printList(head)

	// Test case 2: [1,2,3,4,5] -> [1,5,2,4,3]
	head = &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	reorderList(head)
	printList(head)

	// Test case 3: [1] -> [1]
	head = &ListNode{1, nil}
	reorderList(head)
	printList(head)
}

// Time: O(n) | Space: O(1)
```

## 0146 — Lru Cache

```go
package main

// LeetCode #146: LRU Cache
// https://leetcode.com/problems/lru-cache/
// Difficulty: Medium

import "fmt"

type Node struct {
	key, value int
	prev, next *Node
}

type LRUCache struct {
	capacity int
	cache    map[int]*Node
	head     *Node // LRU
	tail     *Node // MRU
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		head:     head,
		tail:     tail,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.moveToFront(node)
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		node.value = value
		this.moveToFront(node)
		return
	}

	if len(this.cache) == this.capacity {
		this.removeLRU()
	}

	node := &Node{key: key, value: value}
	this.cache[key] = node
	this.addToFront(node)
}

func (this *LRUCache) moveToFront(node *Node) {
	this.removeNode(node)
	this.addToFront(node)
}

func (this *LRUCache) addToFront(node *Node) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) removeLRU() {
	lru := this.tail.prev
	this.removeNode(lru)
	delete(this.cache, lru.key)
}

func main() {
	// Test case 1
	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	fmt.Println(lru.Get(1)) // 1
	lru.Put(3, 3)           // evicts key 2
	fmt.Println(lru.Get(2)) // -1
	lru.Put(4, 4)           // evicts key 1
	fmt.Println(lru.Get(1)) // -1
	fmt.Println(lru.Get(3)) // 3
	fmt.Println(lru.Get(4)) // 4
}

// Time: O(1) per operation | Space: O(capacity)
```

## 0147 — Insertion Sort List

```go
package main

// LeetCode #147: Insertion Sort List
// https://leetcode.com/problems/insertion-sort-list/
// Difficulty: Medium

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	dummy := &ListNode{}

	for head != nil {
		prev := dummy
		for prev.Next != nil && prev.Next.Val < head.Val {
			prev = prev.Next
		}
		next := head.Next
		head.Next = prev.Next
		prev.Next = head
		head = next
	}

	return dummy.Next
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d", head.Val)
		if head.Next != nil {
			fmt.Print(" -> ")
		}
		head = head.Next
	}
	fmt.Println()
}

func main() {
	// Test case 1: [4,2,1,3] -> [1,2,3,4]
	head := &ListNode{4, &ListNode{2, &ListNode{1, &ListNode{3, nil}}}}
	result := insertionSortList(head)
	printList(result)

	// Test case 2: [-1,5,3,4,0] -> [-1,0,3,4,5]
	head = &ListNode{-1, &ListNode{5, &ListNode{3, &ListNode{4, &ListNode{0, nil}}}}}
	result = insertionSortList(head)
	printList(result)

	// Test case 3: [] -> []
	result = insertionSortList(nil)
	printList(result)
}

// Time: O(n^2) | Space: O(1)
```

## 0148 — Sort List

```go
package main

// LeetCode #148: Sort List
// https://leetcode.com/problems/sort-list/
// Difficulty: Medium

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// Find middle
	var prev *ListNode
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	prev.Next = nil

	left := sortList(head)
	right := sortList(slow)

	return merge(left, right)
}

func merge(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			curr.Next = l1
			l1 = l1.Next
		} else {
			curr.Next = l2
			l2 = l2.Next
		}
		curr = curr.Next
	}

	if l1 != nil {
		curr.Next = l1
	}
	if l2 != nil {
		curr.Next = l2
	}

	return dummy.Next
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d", head.Val)
		if head.Next != nil {
			fmt.Print(" -> ")
		}
		head = head.Next
	}
	fmt.Println()
}

func main() {
	// Test case 1: [4,2,1,3] -> [1,2,3,4]
	head := &ListNode{4, &ListNode{2, &ListNode{1, &ListNode{3, nil}}}}
	result := sortList(head)
	printList(result)

	// Test case 2: [-1,5,3,4,0] -> [-1,0,3,4,5]
	head = &ListNode{-1, &ListNode{5, &ListNode{3, &ListNode{4, &ListNode{0, nil}}}}}
	result = sortList(head)
	printList(result)

	// Test case 3: [] -> []
	result = sortList(nil)
	printList(result)
}

// Time: O(n log n) | Space: O(log n)
```

## 0150 — Evaluate Reverse Polish Notation

```go
package main

// LeetCode #150: Evaluate Reverse Polish Notation
// https://leetcode.com/problems/evaluate-reverse-polish-notation/
// Difficulty: Medium

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := make([]int, 0, len(tokens))

	for _, token := range tokens {
		switch token {
		case "+":
			a, b := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, a+b)
		case "-":
			a, b := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, a-b)
		case "*":
			a, b := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, a*b)
		case "/":
			a, b := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, a/b)
		default:
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}

	return stack[0]
}

func main() {
	// Test case 1
	fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"})) // 9

	// Test case 2
	fmt.Println(evalRPN([]string{"4", "13", "5", "/", "+"})) // 6

	// Test case 3
	fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"})) // 22
}

// Time: O(n) | Space: O(n)
```

## 0151 — Reverse Words In A String

```go
package main

// LeetCode #151: Reverse Words in a String
// https://leetcode.com/problems/reverse-words-in-a-string/
// Difficulty: Medium

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	fields := strings.Fields(s)
	for i, j := 0, len(fields)-1; i < j; i, j = i+1, j-1 {
		fields[i], fields[j] = fields[j], fields[i]
	}
	return strings.Join(fields, " ")
}

func main() {
	// Test case 1
	fmt.Println(reverseWords("the sky is blue")) // "blue is sky the"

	// Test case 2
	fmt.Println(reverseWords("  hello world  ")) // "world hello"

	// Test case 3
	fmt.Println(reverseWords("a good   example")) // "example good a"
}

// Time: O(n) | Space: O(n)
```

## 0152 — Maximum Product Subarray

```go
package main

// LeetCode #152: Maximum Product Subarray
// https://leetcode.com/problems/maximum-product-subarray/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxProd, minProd, result := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			maxProd, minProd = minProd, maxProd
		}

		maxProd = max(nums[i], maxProd*nums[i])
		minProd = min(nums[i], minProd*nums[i])

		result = max(result, maxProd)
	}

	return result
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
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
	fmt.Println(maxProduct([]int{-2, 0, -1}))
	fmt.Println(maxProduct([]int{-2, 3, -4}))
}
```

## 0153 — Find Minimum In Rotated Sorted Array

```go
package main

// LeetCode #153: Find Minimum in Rotated Sorted Array
// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/
// Difficulty: Medium
// Time: O(log n), Space: O(1)

import "fmt"

func findMin(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return nums[left]
}

func main() {
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}))
	fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2}))
	fmt.Println(findMin([]int{11, 13, 15, 17}))
}
```

## 0155 — Min Stack

```go
package main

// LeetCode #155: Min Stack
// https://leetcode.com/problems/min-stack/
// Difficulty: Medium
// Time: O(1) per operation, Space: O(n)

import "fmt"

type MinStack struct {
	stack []int
	min   []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if len(this.min) == 0 || val <= this.GetMin() {
		this.min = append(this.min, val)
	}
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	if this.Top() == this.GetMin() {
		this.min = this.min[:len(this.min)-1]
	}
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}

func main() {
	ms := Constructor()
	ms.Push(-2)
	ms.Push(0)
	ms.Push(-3)
	fmt.Println(ms.GetMin())
	ms.Pop()
	fmt.Println(ms.Top())
	fmt.Println(ms.GetMin())
}
```

## 0156 — Binary Tree Upside Down

```go
package main

// LeetCode #156: Binary Tree Upside Down
// https://leetcode.com/problems/binary-tree-upside-down/
// Difficulty: Medium [Paid]
// Time: O(n), Space: O(h) for recursion

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func upsideDownBinaryTree(root *TreeNode) *TreeNode {
	if root == nil || root.Left == nil {
		return root
	}

	newRoot := upsideDownBinaryTree(root.Left)
	root.Left.Left = root.Right
	root.Left.Right = root
	root.Left = nil
	root.Right = nil

	return newRoot
}

func main() {
	root := &TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}, &TreeNode{3, nil, nil}}
	newRoot := upsideDownBinaryTree(root)
	fmt.Println(newRoot.Val)

	root2 := &TreeNode{1, &TreeNode{2, nil, nil}, nil}
	newRoot2 := upsideDownBinaryTree(root2)
	fmt.Println(newRoot2.Val)

	var root3 *TreeNode
	newRoot3 := upsideDownBinaryTree(root3)
	fmt.Println(newRoot3)
}
```

## 0159 — Longest Substring With At Most Two Distinct Characters

```go
package main

// LeetCode #159: Longest Substring with At Most Two Distinct Characters
// https://leetcode.com/problems/longest-substring-with-at-most-two-distinct-characters/
// Difficulty: Medium [Paid]
// Time: O(n), Space: O(1)

import "fmt"

func lengthOfLongestSubstringTwoDistinct(s string) int {
	charCount := make(map[byte]int)
	left, maxLen := 0, 0

	for right := 0; right < len(s); right++ {
		charCount[s[right]]++

		for len(charCount) > 2 {
			charCount[s[left]]--
			if charCount[s[left]] == 0 {
				delete(charCount, s[left])
			}
			left++
		}

		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}

	return maxLen
}

func main() {
	fmt.Println(lengthOfLongestSubstringTwoDistinct("eceba"))
	fmt.Println(lengthOfLongestSubstringTwoDistinct("ccaabbb"))
	fmt.Println(lengthOfLongestSubstringTwoDistinct("abc"))
}
```

## 0161 — One Edit Distance

```go
package main

// LeetCode #161: One Edit Distance
// https://leetcode.com/problems/one-edit-distance/
// Difficulty: Medium [Paid]
// Time: O(n), Space: O(1)

import "fmt"

func isOneEditDistance(s string, t string) bool {
	ns, nt := len(s), len(t)
	if abs(ns-nt) > 1 {
		return false
	}

	if ns > nt {
		s, t = t, s
		ns, nt = nt, ns
	}

	for i := 0; i < ns; i++ {
		if s[i] != t[i] {
			if ns == nt {
				return s[i+1:] == t[i+1:]
			}
			return s[i:] == t[i+1:]
		}
	}

	return ns+1 == nt
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(isOneEditDistance("ab", "acb"))
	fmt.Println(isOneEditDistance("", ""))
	fmt.Println(isOneEditDistance("a", ""))
}
```

## 0162 — Find Peak Element

```go
package main

// LeetCode #162: Find Peak Element
// https://leetcode.com/problems/find-peak-element/
// Difficulty: Medium
// Time: O(log n), Space: O(1)

import "fmt"

func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] > nums[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func main() {
	fmt.Println(findPeakElement([]int{1, 2, 3, 1}))
	fmt.Println(findPeakElement([]int{1, 2, 1, 3, 5, 6, 4}))
	fmt.Println(findPeakElement([]int{1}))
}
```

## 0164 — Maximum Gap

```go
package main

// LeetCode #164: Maximum Gap
// https://leetcode.com/problems/maximum-gap/
// Difficulty: Medium
// Time: O(n), Space: O(n) using bucket sort (Pigeonhole Principle)

import "fmt"

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	minVal, maxVal := nums[0], nums[0]
	for _, num := range nums {
		if num < minVal {
			minVal = num
		}
		if num > maxVal {
			maxVal = num
		}
	}

	if minVal == maxVal {
		return 0
	}

	n := len(nums)
	bucketSize := max(1, (maxVal-minVal)/(n-1))
	bucketCount := (maxVal-minVal)/bucketSize + 1

	bucketMin := make([]int, bucketCount)
	bucketMax := make([]int, bucketCount)
	for i := range bucketMin {
		bucketMin[i] = 1<<31 - 1
		bucketMax[i] = -1 << 31
	}

	for _, num := range nums {
		idx := (num - minVal) / bucketSize
		if num < bucketMin[idx] {
			bucketMin[idx] = num
		}
		if num > bucketMax[idx] {
			bucketMax[idx] = num
		}
	}

	maxGap := 0
	prevMax := minVal
	for i := 0; i < bucketCount; i++ {
		if bucketMin[i] == 1<<31-1 {
			continue
		}
		maxGap = max(maxGap, bucketMin[i]-prevMax)
		prevMax = bucketMax[i]
	}

	return maxGap
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maximumGap([]int{3, 6, 9, 1}))
	fmt.Println(maximumGap([]int{10}))
	fmt.Println(maximumGap([]int{1, 10000000}))
}
```

## 0165 — Compare Version Numbers

```go
package main

// LeetCode #165: Compare Version Numbers
// https://leetcode.com/problems/compare-version-numbers/
// Difficulty: Medium
// Time: O(n+m), Space: O(n+m)

import (
	"fmt"
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")

	n := len(v1)
	if len(v2) > n {
		n = len(v2)
	}

	for i := 0; i < n; i++ {
		num1, num2 := 0, 0
		if i < len(v1) {
			num1, _ = strconv.Atoi(v1[i])
		}
		if i < len(v2) {
			num2, _ = strconv.Atoi(v2[i])
		}

		if num1 < num2 {
			return -1
		}
		if num1 > num2 {
			return 1
		}
	}

	return 0
}

func main() {
	fmt.Println(compareVersion("1.01", "1.001"))
	fmt.Println(compareVersion("1.0", "1.0.0"))
	fmt.Println(compareVersion("0.1", "1.1"))
}
```

## 0166 — Fraction To Recurring Decimal

```go
package main

// LeetCode #166: Fraction to Recurring Decimal
// https://leetcode.com/problems/fraction-to-recurring-decimal/
// Difficulty: Medium
// Time: O(n) where n is the length of the repeating cycle, Space: O(n)

import (
	"fmt"
	"strconv"
)

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}

	result := ""

	if (numerator < 0) != (denominator < 0) {
		result += "-"
	}

	num := abs(numerator)
	den := abs(denominator)

	result += strconv.Itoa(num / den)

	remainder := num % den
	if remainder == 0 {
		return result
	}

	result += "."

	remainderMap := make(map[int]int)
	remainderMap[remainder] = len(result)

	for remainder != 0 {
		remainder *= 10
		result += strconv.Itoa(remainder / den)
		remainder = remainder % den

		if pos, ok := remainderMap[remainder]; ok {
			result = result[:pos] + "(" + result[pos:] + ")"
			break
		}
		remainderMap[remainder] = len(result)
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(fractionToDecimal(1, 2))
	fmt.Println(fractionToDecimal(2, 1))
	fmt.Println(fractionToDecimal(4, 333))
}
```

