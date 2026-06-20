# Medium (Sedang) — Problem 0167–0348

## 0167 — Two Sum Ii Input Array Is Sorted

```go
package main

// LeetCode #167: Two Sum II - Input Array Is Sorted
// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}

	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{2, 3, 4}, 6))
	fmt.Println(twoSum([]int{-1, 0}, -1))
}
```

## 0172 — Factorial Trailing Zeroes

```go
package main

// LeetCode #172: Factorial Trailing Zeroes
// https://leetcode.com/problems/factorial-trailing-zeroes/
// Difficulty: Medium
// Time: O(log n), Space: O(1)

import "fmt"

func trailingZeroes(n int) int {
	count := 0
	for n >= 5 {
		n /= 5
		count += n
	}
	return count
}

func main() {
	fmt.Println(trailingZeroes(3))
	fmt.Println(trailingZeroes(5))
	fmt.Println(trailingZeroes(0))
}
```

## 0173 — Binary Search Tree Iterator

```go
package main

// LeetCode #173: Binary Search Tree Iterator
// https://leetcode.com/problems/binary-search-tree-iterator/
// Difficulty: Medium
// Time: O(1) average per next/hasNext, Space: O(h)

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	iter := BSTIterator{}
	iter.pushLeft(root)
	return iter
}

func (this *BSTIterator) pushLeft(node *TreeNode) {
	for node != nil {
		this.stack = append(this.stack, node)
		node = node.Left
	}
}

func (this *BSTIterator) Next() int {
	node := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	this.pushLeft(node.Right)
	return node.Val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0
}

func main() {
	root := &TreeNode{7, &TreeNode{3, nil, nil}, &TreeNode{15, &TreeNode{9, nil, nil}, &TreeNode{20, nil, nil}}}
	iter := Constructor(root)
	fmt.Println(iter.Next())
	fmt.Println(iter.Next())
	fmt.Println(iter.HasNext())
	fmt.Println(iter.Next())
	fmt.Println(iter.HasNext())
	fmt.Println(iter.Next())
	fmt.Println(iter.HasNext())
}
```

## 0176 — Second Highest Salary

```go
package main

// LeetCode #176: Second Highest Salary
// https://leetcode.com/problems/second-highest-salary/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

import "math"

func secondHighestSalary(salaries []int) int {
	first, second := math.MinInt32, math.MinInt32

	for _, s := range salaries {
		if s > first {
			second = first
			first = s
		} else if s > second && s < first {
			second = s
		}
	}

	if second == math.MinInt32 {
		return 0
	}
	return second
}

func main() {
	fmt.Println(secondHighestSalary([]int{100, 200, 300}))
	fmt.Println(secondHighestSalary([]int{100, 100}))
	fmt.Println(secondHighestSalary([]int{50, 100, 100, 75}))
}
```

## 0177 — Nth Highest Salary

```go
package main

// LeetCode #177: Nth Highest Salary
// https://leetcode.com/problems/nth-highest-salary/
// Difficulty: Medium
// Time: O(n log n) average, Space: O(n) for quickselect

import (
	"fmt"
	"sort"
)

func nthHighestSalary(salaries []int, n int) int {
	if n <= 0 || n > len(salaries) {
		return 0
	}

	// Use sort + deduplicate
	seen := make(map[int]bool)
	unique := []int{}
	for _, s := range salaries {
		if !seen[s] {
			seen[s] = true
			unique = append(unique, s)
		}
	}

	if n > len(unique) {
		return 0
	}

	sort.Slice(unique, func(i, j int) bool {
		return unique[i] > unique[j]
	})

	return unique[n-1]
}

func main() {
	fmt.Println(nthHighestSalary([]int{100, 200, 300, 200}, 2))
	fmt.Println(nthHighestSalary([]int{100, 100}, 2))
	fmt.Println(nthHighestSalary([]int{60, 70, 80, 90, 100}, 3))
}
```

## 0178 — Rank Scores

```go
package main

// LeetCode #178: Rank Scores
// https://leetcode.com/problems/rank-scores/
// Difficulty: Medium
// Time: O(n log n), Space: O(n)

import (
	"fmt"
	"sort"
)

func rankScores(scores []int) []int {
	if len(scores) == 0 {
		return nil
	}

	type pair struct {
		score int
		idx   int
	}

	pairs := make([]pair, len(scores))
	for i, s := range scores {
		pairs[i] = pair{s, i}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].score > pairs[j].score
	})

	ranks := make([]int, len(scores))
	rank := 1
	for i, p := range pairs {
		if i > 0 && p.score < pairs[i-1].score {
			rank = i + 1
		}
		ranks[p.idx] = rank
	}

	return ranks
}

func main() {
	fmt.Println(rankScores([]int{100, 90, 90, 80}))
	fmt.Println(rankScores([]int{50, 60, 70}))
	fmt.Println(rankScores([]int{90, 80, 80, 70, 60, 60}))
}
```

## 0179 — Largest Number

```go
package main

// LeetCode #179: Largest Number
// https://leetcode.com/problems/largest-number/
// Difficulty: Medium
// Time: O(n log n), Space: O(n)

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	strs := make([]string, len(nums))
	for i, num := range nums {
		strs[i] = strconv.Itoa(num)
	}

	sort.Slice(strs, func(i, j int) bool {
		return strs[i]+strs[j] > strs[j]+strs[i]
	})

	if strs[0] == "0" {
		return "0"
	}

	return strings.Join(strs, "")
}

func main() {
	fmt.Println(largestNumber([]int{10, 2}))
	fmt.Println(largestNumber([]int{3, 30, 34, 5, 9}))
	fmt.Println(largestNumber([]int{0, 0}))
}
```

## 0180 — Consecutive Numbers

```go
package main

// LeetCode #180: Consecutive Numbers
// https://leetcode.com/problems/consecutive-numbers/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

func consecutiveNumbers(nums []int) []int {
	if len(nums) < 3 {
		return nil
	}

	result := []int{}
	seen := make(map[int]bool)

	for i := 0; i < len(nums)-2; i++ {
		if nums[i] == nums[i+1] && nums[i] == nums[i+2] && !seen[nums[i]] {
			result = append(result, nums[i])
			seen[nums[i]] = true
		}
	}

	return result
}

func main() {
	fmt.Println(consecutiveNumbers([]int{1, 1, 1, 2, 2, 3, 3, 3}))
	fmt.Println(consecutiveNumbers([]int{1, 2, 3, 4}))
	fmt.Println(consecutiveNumbers([]int{1, 1, 1, 1, 2, 2, 2}))
}
```

## 0184 — Department Highest Salary

```go
package main

// LeetCode #184: Department Highest Salary
// https://leetcode.com/problems/department-highest-salary/
// Difficulty: Medium
// Time: O(n), Space: O(d) where d is number of departments

import "fmt"

type Employee struct {
	Name       string
	Salary     int
	Department string
}

type DeptSalary struct {
	Department string
	Employee   string
	Salary     int
}

func departmentHighestSalary(employees []Employee) []DeptSalary {
	deptMax := make(map[string]int)
	for _, e := range employees {
		if e.Salary > deptMax[e.Department] {
			deptMax[e.Department] = e.Salary
		}
	}

	result := []DeptSalary{}
	for _, e := range employees {
		if e.Salary == deptMax[e.Department] {
			result = append(result, DeptSalary{e.Department, e.Name, e.Salary})
		}
	}

	return result
}

func main() {
	emps := []Employee{
		{"Joe", 85000, "IT"},
		{"Jim", 90000, "IT"},
		{"Henry", 80000, "Sales"},
		{"Sam", 60000, "Sales"},
		{"Max", 90000, "IT"},
	}
	fmt.Println(departmentHighestSalary(emps))

	emps2 := []Employee{
		{"A", 50000, "Eng"},
		{"B", 60000, "Eng"},
	}
	fmt.Println(departmentHighestSalary(emps2))

	fmt.Println(departmentHighestSalary(nil))
}
```

## 0186 — Reverse Words In A String Ii

```go
package main

// LeetCode #186: Reverse Words in a String II
// https://leetcode.com/problems/reverse-words-in-a-string-ii/
// Difficulty: Medium [Paid]
// Time: O(n), Space: O(1) in-place

import "fmt"

func reverseWords(s []byte) {
	reverse := func(arr []byte, l, r int) {
		for l < r {
			arr[l], arr[r] = arr[r], arr[l]
			l++
			r--
		}
	}

	reverse(s, 0, len(s)-1)

	start := 0
	for i := 0; i <= len(s); i++ {
		if i == len(s) || s[i] == ' ' {
			reverse(s, start, i-1)
			start = i + 1
		}
	}
}

func main() {
	s1 := []byte("the sky is blue")
	reverseWords(s1)
	fmt.Println(string(s1))

	s2 := []byte("hello world")
	reverseWords(s2)
	fmt.Println(string(s2))

	s3 := []byte("a")
	reverseWords(s3)
	fmt.Println(string(s3))
}
```

## 0187 — Repeated Dna Sequences

```go
package main

// LeetCode #187: Repeated DNA Sequences
// https://leetcode.com/problems/repeated-dna-sequences/
// Difficulty: Medium
// Time: O(n), Space: O(n)

import "fmt"

func findRepeatedDnaSequences(s string) []string {
	if len(s) < 10 {
		return nil
	}

	seen := make(map[string]int)
	result := []string{}

	for i := 0; i <= len(s)-10; i++ {
		sub := s[i : i+10]
		seen[sub]++
		if seen[sub] == 2 {
			result = append(result, sub)
		}
	}

	return result
}

func main() {
	fmt.Println(findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"))
	fmt.Println(findRepeatedDnaSequences("AAAAAAAAAAAAA"))
	fmt.Println(findRepeatedDnaSequences("AAAAAAAAAAA"))
}
```

## 0189 — Rotate Array

```go
package main

// LeetCode #189: Rotate Array
// https://leetcode.com/problems/rotate-array/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

func rotate(nums []int, k int) {
	n := len(nums)
	if n == 0 {
		return
	}
	k = k % n
	if k == 0 {
		return
	}

	reverse := func(arr []int, l, r int) {
		for l < r {
			arr[l], arr[r] = arr[r], arr[l]
			l++
			r--
		}
	}

	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func main() {
	nums1 := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums1, 3)
	fmt.Println(nums1)

	nums2 := []int{-1, -100, 3, 99}
	rotate(nums2, 2)
	fmt.Println(nums2)

	nums3 := []int{1}
	rotate(nums3, 0)
	fmt.Println(nums3)
}
```

## 0192 — Word Frequency

```go
package main

// LeetCode #192: Word Frequency
// https://leetcode.com/problems/word-frequency/
// Difficulty: Medium
// Time: O(n log n), Space: O(n)

import (
	"fmt"
	"sort"
	"strings"
)

func wordFrequency(text string) []string {
	words := strings.Fields(text)
	freq := make(map[string]int)

	for _, w := range words {
		freq[w]++
	}

	type kv struct {
		word  string
		count int
	}

	var sorted []kv
	for w, c := range freq {
		sorted = append(sorted, kv{w, c})
	}

	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].count != sorted[j].count {
			return sorted[i].count > sorted[j].count
		}
		return sorted[i].word < sorted[j].word
	})

	result := make([]string, len(sorted))
	for i, kv := range sorted {
		result[i] = fmt.Sprintf("%s %d", kv.word, kv.count)
	}
	return result
}

func main() {
	for _, line := range wordFrequency("the day is sunny the the the sunny is is") {
		fmt.Println(line)
	}
	fmt.Println("---")
	for _, line := range wordFrequency("hello world hello") {
		fmt.Println(line)
	}
	fmt.Println("---")
	for _, line := range wordFrequency("") {
		fmt.Println(line)
	}
}
```

## 0194 — Transpose File

```go
package main

// LeetCode #194: Transpose File
// https://leetcode.com/problems/transpose-file/
// Difficulty: Medium
// Time: O(m*n), Space: O(m*n)

import (
	"fmt"
	"strings"
)

func transposeFile(content string) []string {
	lines := strings.Split(strings.TrimSpace(content), "\n")
	if len(lines) == 0 || lines[0] == "" {
		return nil
	}

	matrix := make([][]string, len(lines))
	for i, line := range lines {
		matrix[i] = strings.Fields(line)
	}

	if len(matrix) == 0 {
		return nil
	}

	cols := 0
	for _, row := range matrix {
		if len(row) > cols {
			cols = len(row)
		}
	}

	result := make([]string, cols)
	for c := 0; c < cols; c++ {
		var row []string
		for r := 0; r < len(matrix); r++ {
			if c < len(matrix[r]) {
				row = append(row, matrix[r][c])
			}
		}
		result[c] = strings.Join(row, " ")
	}

	return result
}

func main() {
	content := "name age\nalice 21\nryan 30"
	for _, line := range transposeFile(content) {
		fmt.Println(line)
	}
	fmt.Println("---")
	content2 := "a b c\nd e f"
	for _, line := range transposeFile(content2) {
		fmt.Println(line)
	}
}
```

## 0198 — House Robber

```go
package main

// LeetCode #198: House Robber
// https://leetcode.com/problems/house-robber/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

func rob(nums []int) int {
	prev, curr := 0, 0

	for _, num := range nums {
		prev, curr = curr, max(curr, prev+num)
	}

	return curr
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
	fmt.Println(rob([]int{0}))
}
```

## 0199 — Binary Tree Right Side View

```go
package main

// LeetCode #199: Binary Tree Right Side View
// https://leetcode.com/problems/binary-tree-right-side-view/
// Difficulty: Medium
// Time: O(n), Space: O(h) for queue

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := []int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if i == levelSize-1 {
				result = append(result, node.Val)
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return result
}

func main() {
	root := &TreeNode{1, &TreeNode{2, nil, &TreeNode{5, nil, nil}}, &TreeNode{3, nil, &TreeNode{4, nil, nil}}}
	fmt.Println(rightSideView(root))

	root2 := &TreeNode{1, nil, &TreeNode{3, nil, nil}}
	fmt.Println(rightSideView(root2))

	fmt.Println(rightSideView(nil))
}
```

## 0200 — Number Of Islands

```go
package main

// LeetCode #200: Number of Islands
// https://leetcode.com/problems/number-of-islands/
// Difficulty: Medium
// Time: O(m*n), Space: O(m*n) worst case for recursion

import "fmt"

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])
	count := 0

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] == '0' {
			return
		}
		grid[r][c] = '0'
		dfs(r-1, c)
		dfs(r+1, c)
		dfs(r, c-1)
		dfs(r, c+1)
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '1' {
				count++
				dfs(r, c)
			}
		}
	}

	return count
}

func main() {
	grid1 := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	fmt.Println(numIslands(grid1))

	grid2 := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	fmt.Println(numIslands(grid2))
}
```

## 0201 — Bitwise And Of Numbers Range

```go
package main

// LeetCode #201: Bitwise AND of Numbers Range
// https://leetcode.com/problems/bitwise-and-of-numbers-range/
// Difficulty: Medium
// Time: O(log max(left, right)), Space: O(1)

import "fmt"

func rangeBitwiseAnd(left int, right int) int {
	shift := 0
	for left < right {
		left >>= 1
		right >>= 1
		shift++
	}
	return left << shift
}

func main() {
	fmt.Println(rangeBitwiseAnd(5, 7))
	fmt.Println(rangeBitwiseAnd(0, 0))
	fmt.Println(rangeBitwiseAnd(1, 2147483647))
}
```

## 0204 — Count Primes

```go
package main

// LeetCode #204: Count Primes
// https://leetcode.com/problems/count-primes/
// Difficulty: Medium
// Time: O(n log log n), Space: O(n)

import "fmt"

func countPrimes(n int) int {
	if n < 2 {
		return 0
	}

	isPrime := make([]bool, n)
	for i := 2; i < n; i++ {
		isPrime[i] = true
	}

	for i := 2; i*i < n; i++ {
		if isPrime[i] {
			for j := i * i; j < n; j += i {
				isPrime[j] = false
			}
		}
	}

	count := 0
	for i := 2; i < n; i++ {
		if isPrime[i] {
			count++
		}
	}

	return count
}

func main() {
	fmt.Println(countPrimes(10))
	fmt.Println(countPrimes(0))
	fmt.Println(countPrimes(1))
}
```

## 0207 — Course Schedule

```go
package main

// LeetCode #207: Course Schedule
// https://leetcode.com/problems/course-schedule/
// Difficulty: Medium
// Time: O(V+E), Space: O(V+E)

import "fmt"

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	inDegree := make([]int, numCourses)

	for _, pre := range prerequisites {
		course, prereq := pre[0], pre[1]
		graph[prereq] = append(graph[prereq], course)
		inDegree[course]++
	}

	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	count := 0
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		count++

		for _, neighbor := range graph[node] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	return count == numCourses
}

func main() {
	fmt.Println(canFinish(2, [][]int{{1, 0}}))
	fmt.Println(canFinish(2, [][]int{{1, 0}, {0, 1}}))
	fmt.Println(canFinish(5, [][]int{{1, 0}, {2, 1}, {3, 2}}))
}
```

## 0208 — Implement Trie Prefix Tree

```go
package main

// LeetCode #208: Implement Trie (Prefix Tree)
// https://leetcode.com/problems/implement-trie-prefix-tree/
// Difficulty: Medium
// Time: O(n) per operation, Space: O(total characters)

import "fmt"

type TrieNode struct {
	children [26]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{&TrieNode{}}
}

func (this *Trie) Insert(word string) {
	node := this.root
	for i := 0; i < len(word); i++ {
		idx := word[i] - 'a'
		if node.children[idx] == nil {
			node.children[idx] = &TrieNode{}
		}
		node = node.children[idx]
	}
	node.isEnd = true
}

func (this *Trie) Search(word string) bool {
	node := this.root
	for i := 0; i < len(word); i++ {
		idx := word[i] - 'a'
		if node.children[idx] == nil {
			return false
		}
		node = node.children[idx]
	}
	return node.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this.root
	for i := 0; i < len(prefix); i++ {
		idx := prefix[i] - 'a'
		if node.children[idx] == nil {
			return false
		}
		node = node.children[idx]
	}
	return true
}

func main() {
	trie := Constructor()
	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))
	fmt.Println(trie.Search("app"))
	fmt.Println(trie.StartsWith("app"))
	trie.Insert("app")
	fmt.Println(trie.Search("app"))
}
```

## 0209 — Minimum Size Subarray Sum

```go
package main

// LeetCode #209: Minimum Size Subarray Sum
// https://leetcode.com/problems/minimum-size-subarray-sum/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import (
	"fmt"
	"math"
)

func minSubArrayLen(target int, nums []int) int {
	left, sum := 0, 0
	minLen := math.MaxInt32

	for right := 0; right < len(nums); right++ {
		sum += nums[right]

		for sum >= target {
			if right-left+1 < minLen {
				minLen = right - left + 1
			}
			sum -= nums[left]
			left++
		}
	}

	if minLen == math.MaxInt32 {
		return 0
	}
	return minLen
}

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	fmt.Println(minSubArrayLen(4, []int{1, 4, 4}))
	fmt.Println(minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1}))
}
```

## 0210 — Course Schedule Ii

```go
package main

// LeetCode #210: Course Schedule II
// https://leetcode.com/problems/course-schedule-ii/
// Difficulty: Medium
// Time: O(V+E), Space: O(V+E)

import "fmt"

func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := make([][]int, numCourses)
	inDegree := make([]int, numCourses)

	for _, pre := range prerequisites {
		course, prereq := pre[0], pre[1]
		graph[prereq] = append(graph[prereq], course)
		inDegree[course]++
	}

	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	result := []int{}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node)

		for _, neighbor := range graph[node] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	if len(result) != numCourses {
		return nil
	}
	return result
}

func main() {
	fmt.Println(findOrder(2, [][]int{{1, 0}}))
	fmt.Println(findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))
	fmt.Println(findOrder(1, [][]int{}))
}
```

## 0211 — Design Add And Search Words Data Structure

```go
package main

// LeetCode #211: Design Add and Search Words Data Structure
// https://leetcode.com/problems/design-add-and-search-words-data-structure/
// Difficulty: Medium
// Time: O(n) for add, O(26^m) worst case for search with wildcards, Space: O(total chars)

import "fmt"

type WordDictionaryNode struct {
	children [26]*WordDictionaryNode
	isEnd    bool
}

type WordDictionary struct {
	root *WordDictionaryNode
}

func Constructor() WordDictionary {
	return WordDictionary{&WordDictionaryNode{}}
}

func (this *WordDictionary) AddWord(word string) {
	node := this.root
	for i := 0; i < len(word); i++ {
		idx := word[i] - 'a'
		if node.children[idx] == nil {
			node.children[idx] = &WordDictionaryNode{}
		}
		node = node.children[idx]
	}
	node.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	return this.search(this.root, word, 0)
}

func (this *WordDictionary) search(node *WordDictionaryNode, word string, idx int) bool {
	if node == nil {
		return false
	}
	if idx == len(word) {
		return node.isEnd
	}

	if word[idx] == '.' {
		for i := 0; i < 26; i++ {
			if this.search(node.children[i], word, idx+1) {
				return true
			}
		}
		return false
	}

	child := node.children[word[idx]-'a']
	return this.search(child, word, idx+1)
}

func main() {
	wd := Constructor()
	wd.AddWord("bad")
	wd.AddWord("dad")
	wd.AddWord("mad")
	fmt.Println(wd.Search("pad"))
	fmt.Println(wd.Search("bad"))
	fmt.Println(wd.Search(".ad"))
	fmt.Println(wd.Search("b.."))
}
```

## 0213 — House Robber Ii

```go
package main

// LeetCode #213: House Robber II
// https://leetcode.com/problems/house-robber-ii/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	robLinear := func(arr []int) int {
		prev, curr := 0, 0
		for _, num := range arr {
			prev, curr = curr, max(curr, prev+num)
		}
		return curr
	}

	return max(robLinear(nums[1:]), robLinear(nums[:len(nums)-1]))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(rob([]int{2, 3, 2}))
	fmt.Println(rob([]int{1, 2, 3, 1}))
	fmt.Println(rob([]int{1, 2, 3}))
}
```

## 0215 — Kth Largest Element In An Array

```go
package main

// LeetCode #215: Kth Largest Element in an Array
// https://leetcode.com/problems/kth-largest-element-in-an-array/
// Difficulty: Medium
// Time: O(n) average with quickselect, O(n log n) worst case, Space: O(log n)

import "fmt"

func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(nums []int, left, right, kSmallest int) int {
	if left == right {
		return nums[left]
	}

	pivotIdx := partition(nums, left, right)

	if kSmallest == pivotIdx {
		return nums[kSmallest]
	} else if kSmallest < pivotIdx {
		return quickSelect(nums, left, pivotIdx-1, kSmallest)
	}
	return quickSelect(nums, pivotIdx+1, right, kSmallest)
}

func partition(nums []int, left, right int) int {
	pivot := nums[right]
	storeIdx := left

	for i := left; i < right; i++ {
		if nums[i] < pivot {
			nums[storeIdx], nums[i] = nums[i], nums[storeIdx]
			storeIdx++
		}
	}

	nums[storeIdx], nums[right] = nums[right], nums[storeIdx]
	return storeIdx
}

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
	fmt.Println(findKthLargest([]int{1}, 1))
}
```

## 0216 — Combination Sum Iii

```go
package main

// LeetCode #216: Combination Sum III
// https://leetcode.com/problems/combination-sum-iii/
// Difficulty: Medium
// Time: O(C(9,k)), Space: O(k)

import "fmt"

func combinationSum3(k int, n int) [][]int {
	result := [][]int{}
	var backtrack func(start, remaining int, combo []int)
	backtrack = func(start, remaining int, combo []int) {
		if len(combo) == k && remaining == 0 {
			comboCopy := make([]int, len(combo))
			copy(comboCopy, combo)
			result = append(result, comboCopy)
			return
		}
		if len(combo) > k || remaining < 0 {
			return
		}

		for i := start; i <= 9; i++ {
			combo = append(combo, i)
			backtrack(i+1, remaining-i, combo)
			combo = combo[:len(combo)-1]
		}
	}

	backtrack(1, n, []int{})
	return result
}

func main() {
	fmt.Println(combinationSum3(3, 7))
	fmt.Println(combinationSum3(3, 9))
	fmt.Println(combinationSum3(4, 1))
}
```

## 0221 — Maximal Square

```go
package main

// LeetCode #221: Maximal Square
// https://leetcode.com/problems/maximal-square/
// Difficulty: Medium
// Time: O(m*n), Space: O(n)

import "fmt"

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	rows, cols := len(matrix), len(matrix[0])
	dp := make([]int, cols+1)
	maxLen, prev := 0, 0

	for i := 0; i < rows; i++ {
		for j := 1; j <= cols; j++ {
			temp := dp[j]
			if matrix[i][j-1] == '1' {
				dp[j] = min(dp[j], min(dp[j-1], prev)) + 1
				if dp[j] > maxLen {
					maxLen = dp[j]
				}
			} else {
				dp[j] = 0
			}
			prev = temp
		}
	}

	return maxLen * maxLen
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	matrix1 := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	fmt.Println(maximalSquare(matrix1))

	matrix2 := [][]byte{{'0', '1'}, {'1', '0'}}
	fmt.Println(maximalSquare(matrix2))

	matrix3 := [][]byte{{'0'}}
	fmt.Println(maximalSquare(matrix3))
}
```

## 0223 — Rectangle Area

```go
package main

// LeetCode #223: Rectangle Area
// https://leetcode.com/problems/rectangle-area/
// Difficulty: Medium
// Time: O(1), Space: O(1)

import "fmt"

func computeArea(ax1, ay1, ax2, ay2, bx1, by1, bx2, by2 int) int {
	areaA := (ax2 - ax1) * (ay2 - ay1)
	areaB := (bx2 - bx1) * (by2 - by1)

	overlapX := max(0, min(ax2, bx2)-max(ax1, bx1))
	overlapY := max(0, min(ay2, by2)-max(ay1, by1))

	return areaA + areaB - overlapX*overlapY
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
	fmt.Println(computeArea(-3, 0, 3, 4, 0, -1, 9, 2))
	fmt.Println(computeArea(-2, -2, 2, 2, -2, -2, 2, 2))
	fmt.Println(computeArea(0, 0, 0, 0, -1, -1, 1, 1))
}
```

## 0227 — Basic Calculator Ii

```go
package main

// LeetCode #227: Basic Calculator II
// https://leetcode.com/problems/basic-calculator-ii/
// Difficulty: Medium
// Time: O(n), Space: O(n)

import "fmt"

func calculate(s string) int {
	stack := []int{}
	num := 0
	op := '+'

	for i, ch := range s {
		if ch >= '0' && ch <= '9' {
			num = num*10 + int(ch-'0')
		}

		if ch == '+' || ch == '-' || ch == '*' || ch == '/' || i == len(s)-1 {
			if s[i] == ' ' && i != len(s)-1 {
				continue
			}
			switch op {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num
			case '/':
				stack[len(stack)-1] /= num
			}
			op = ch
			num = 0
		}
	}

	result := 0
	for _, v := range stack {
		result += v
	}
	return result
}

func main() {
	fmt.Println(calculate("3+2*2"))
	fmt.Println(calculate(" 3/2 "))
	fmt.Println(calculate(" 3+5 / 2 "))
}
```

## 0229 — Majority Element Ii

```go
package main

// LeetCode #229: Majority Element II
// https://leetcode.com/problems/majority-element-ii/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

func majorityElement(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}

	candidate1, candidate2 := 0, 0
	count1, count2 := 0, 0

	for _, num := range nums {
		if num == candidate1 {
			count1++
		} else if num == candidate2 {
			count2++
		} else if count1 == 0 {
			candidate1 = num
			count1 = 1
		} else if count2 == 0 {
			candidate2 = num
			count2 = 1
		} else {
			count1--
			count2--
		}
	}

	count1, count2 = 0, 0
	for _, num := range nums {
		if num == candidate1 {
			count1++
		} else if num == candidate2 {
			count2++
		}
	}

	result := []int{}
	threshold := len(nums) / 3
	if count1 > threshold {
		result = append(result, candidate1)
	}
	if count2 > threshold {
		result = append(result, candidate2)
	}

	return result
}

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))
	fmt.Println(majorityElement([]int{1}))
	fmt.Println(majorityElement([]int{1, 2}))
}
```

## 0230 — Kth Smallest Element In A Bst

```go
package main

// LeetCode #230: Kth Smallest Element in a BST
// https://leetcode.com/problems/kth-smallest-element-in-a-bst/
// Difficulty: Medium
// Time: O(h+k), Space: O(h)

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	stack := []*TreeNode{}
	node := root

	for {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}

		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--

		if k == 0 {
			return node.Val
		}

		node = node.Right
	}
}

func main() {
	root := &TreeNode{3, &TreeNode{1, nil, &TreeNode{2, nil, nil}}, &TreeNode{4, nil, nil}}
	fmt.Println(kthSmallest(root, 1))

	root2 := &TreeNode{5, &TreeNode{3, &TreeNode{2, &TreeNode{1, nil, nil}, nil}, &TreeNode{4, nil, nil}}, &TreeNode{6, nil, nil}}
	fmt.Println(kthSmallest(root2, 3))
}
```

## 0235 — Lowest Common Ancestor Of A Binary Search Tree

```go
package main

// LeetCode #235: Lowest Common Ancestor of a Binary Search Tree
// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/
// Difficulty: Medium
// Time: O(h), Space: O(1)

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	for root != nil {
		if p.Val < root.Val && q.Val < root.Val {
			root = root.Left
		} else if p.Val > root.Val && q.Val > root.Val {
			root = root.Right
		} else {
			return root
		}
	}
	return nil
}

func main() {
	root := &TreeNode{6, &TreeNode{2, &TreeNode{0, nil, nil}, &TreeNode{4, &TreeNode{3, nil, nil}, &TreeNode{5, nil, nil}}}, &TreeNode{8, &TreeNode{7, nil, nil}, &TreeNode{9, nil, nil}}}
	fmt.Println(lowestCommonAncestor(root, root.Left, root.Right).Val)

	p := root.Left
	q := root.Left.Right
	fmt.Println(lowestCommonAncestor(root, p, q).Val)

	p2 := root.Left
	q2 := root.Left.Left
	fmt.Println(lowestCommonAncestor(root, p2, q2).Val)
}
```

## 0236 — Lowest Common Ancestor Of A Binary Tree

```go
package main

// LeetCode #236: Lowest Common Ancestor of a Binary Tree
// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/
// Difficulty: Medium
// Time: O(n), Space: O(h)

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}
	return right
}

func main() {
	root := &TreeNode{3, &TreeNode{5, &TreeNode{6, nil, nil}, &TreeNode{2, &TreeNode{7, nil, nil}, &TreeNode{4, nil, nil}}}, &TreeNode{1, &TreeNode{0, nil, nil}, &TreeNode{8, nil, nil}}}
	p, q := root.Left, root.Right
	fmt.Println(lowestCommonAncestor(root, p, q).Val)

	p, q = root.Left, root.Left.Right.Right
	fmt.Println(lowestCommonAncestor(root, p, q).Val)

	p, q = root.Left.Left, root.Left.Right
	fmt.Println(lowestCommonAncestor(root, p, q).Val)
}
```

## 0237 — Delete Node In A Linked List

```go
package main

// LeetCode #237: Delete Node in a Linked List
// https://leetcode.com/problems/delete-node-in-a-linked-list/
// Difficulty: Medium
// Time: O(1), Space: O(1)

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
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
	head := &ListNode{4, &ListNode{5, &ListNode{1, &ListNode{9, nil}}}}
	deleteNode(head.Next)
	printList(head)

	head2 := &ListNode{4, &ListNode{5, &ListNode{1, &ListNode{9, nil}}}}
	deleteNode(head2.Next.Next)
	printList(head2)

	head3 := &ListNode{1, &ListNode{2, nil}}
	deleteNode(head3)
	printList(head3)
}
```

## 0238 — Product Of Array Except Self

```go
package main

// LeetCode #238: Product of Array Except Self
// https://leetcode.com/problems/product-of-array-except-self/
// Difficulty: Medium
// Time: O(n), Space: O(1) excluding output array

import "fmt"

func productExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	result[0] = 1
	for i := 1; i < n; i++ {
		result[i] = result[i-1] * nums[i-1]
	}

	right := 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= right
		right *= nums[i]
	}

	return result
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
	fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3}))
	fmt.Println(productExceptSelf([]int{0, 0}))
}
```

## 0240 — Search A 2D Matrix Ii

```go
package main

// LeetCode #240: Search a 2D Matrix II
// https://leetcode.com/problems/search-a-2d-matrix-ii/
// Difficulty: Medium
// Time: O(m+n), Space: O(1)

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	row, col := 0, len(matrix[0])-1

	for row < len(matrix) && col >= 0 {
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] > target {
			col--
		} else {
			row++
		}
	}

	return false
}

func main() {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	fmt.Println(searchMatrix(matrix, 5))
	fmt.Println(searchMatrix(matrix, 20))
	fmt.Println(searchMatrix(matrix, 30))
}
```

## 0241 — Different Ways To Add Parentheses

```go
package main

// LeetCode #241: Different Ways to Add Parentheses
// https://leetcode.com/problems/different-ways-to-add-parentheses/
// Difficulty: Medium
// Time: O(2^n), Space: O(2^n) for result storage

import (
	"fmt"
	"strconv"
)

func diffWaysToCompute(expression string) []int {
	if isNumber(expression) {
		num, _ := strconv.Atoi(expression)
		return []int{num}
	}

	result := []int{}
	for i, ch := range expression {
		if ch == '+' || ch == '-' || ch == '*' {
			left := diffWaysToCompute(expression[:i])
			right := diffWaysToCompute(expression[i+1:])

			for _, l := range left {
				for _, r := range right {
					switch ch {
					case '+':
						result = append(result, l+r)
					case '-':
						result = append(result, l-r)
					case '*':
						result = append(result, l*r)
					}
				}
			}
		}
	}

	return result
}

func isNumber(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(diffWaysToCompute("2-1-1"))
	fmt.Println(diffWaysToCompute("2*3-4*5"))
	fmt.Println(diffWaysToCompute("3"))
}
```

## 0244 — Shortest Word Distance Ii

```go
package main

// LeetCode #244: Shortest Word Distance II
// https://leetcode.com/problems/shortest-word-distance-ii/
// Difficulty: Medium [Paid]
// Time: O(n) for init, O(m+n) for shortest, Space: O(n)

import (
	"fmt"
	"math"
)

type WordDistance struct {
	indices map[string][]int
}

func Constructor(wordsDict []string) WordDistance {
	indices := make(map[string][]int)
	for i, w := range wordsDict {
		indices[w] = append(indices[w], i)
	}
	return WordDistance{indices}
}

func (this *WordDistance) Shortest(word1, word2 string) int {
	list1 := this.indices[word1]
	list2 := this.indices[word2]

	minDist := math.MaxInt32
	i, j := 0, 0

	for i < len(list1) && j < len(list2) {
		dist := list1[i] - list2[j]
		if dist < 0 {
			minDist = min(minDist, -dist)
			i++
		} else {
			minDist = min(minDist, dist)
			j++
		}
	}

	return minDist
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	wd := Constructor([]string{"practice", "makes", "perfect", "coding", "makes"})
	fmt.Println(wd.Shortest("coding", "practice"))
	fmt.Println(wd.Shortest("makes", "coding"))
	fmt.Println(wd.Shortest("makes", "practice"))
}
```

## 0245 — Shortest Word Distance Iii

```go
package main

// LeetCode #245: Shortest Word Distance III
// https://leetcode.com/problems/shortest-word-distance-iii/
// Difficulty: Medium [Paid]
// Time: O(n), Space: O(1)

import (
	"fmt"
	"math"
)

func shortestWordDistance(wordsDict []string, word1 string, word2 string) int {
	idx1, idx2 := -1, -1
	minDist := math.MaxInt32

	for i, word := range wordsDict {
		if word == word1 {
			idx1 = i
		}
		if word == word2 {
			if word1 == word2 {
				idx1 = idx2
			}
			idx2 = i
		}
		if idx1 != -1 && idx2 != -1 {
			dist := idx1 - idx2
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
	fmt.Println(shortestWordDistance([]string{"practice", "makes", "perfect", "coding", "makes"}, "makes", "coding"))
	fmt.Println(shortestWordDistance([]string{"practice", "makes", "perfect", "coding", "makes"}, "makes", "makes"))
	fmt.Println(shortestWordDistance([]string{"a", "a"}, "a", "a"))
}
```

## 0247 — Strobogrammatic Number Ii

```go
package main

// LeetCode #247: Strobogrammatic Number II
// https://leetcode.com/problems/strobogrammatic-number-ii/
// Difficulty: Medium [Paid]
// Time: O(5^(n/2)), Space: O(n)

import "fmt"

func findStrobogrammatic(n int) []string {
	return build(n, n)
}

func build(n, final int) []string {
	if n == 0 {
		return []string{""}
	}
	if n == 1 {
		return []string{"0", "1", "8"}
	}

	inner := build(n-2, final)
	result := []string{}

	for _, s := range inner {
		if n != final {
			result = append(result, "0"+s+"0")
		}
		result = append(result, "1"+s+"1")
		result = append(result, "6"+s+"9")
		result = append(result, "8"+s+"8")
		result = append(result, "9"+s+"6")
	}

	return result
}

func main() {
	fmt.Println(findStrobogrammatic(2))
	fmt.Println(findStrobogrammatic(1))
	fmt.Println(findStrobogrammatic(3))
}
```

## 0249 — Group Shifted Strings

```go
package main

// LeetCode #249: Group Shifted Strings
// https://leetcode.com/problems/group-shifted-strings/
// Difficulty: Medium [Paid]
// Time: O(n * m), Space: O(n * m)

import (
	"fmt"
	"strings"
)

func groupStrings(strs []string) [][]string {
	groups := make(map[string][]string)

	for _, s := range strs {
		key := getKey(s)
		groups[key] = append(groups[key], s)
	}

	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}
	return result
}

func getKey(s string) string {
	if len(s) == 0 {
		return ""
	}

	shift := s[0] - 'a'
	var sb strings.Builder

	for i := 0; i < len(s); i++ {
		diff := (int(s[i]-'a') - int(shift) + 26) % 26
		sb.WriteByte(byte(diff + 'a'))
	}

	return sb.String()
}

func main() {
	fmt.Println(groupStrings([]string{"abc", "bcd", "acef", "xyz", "az", "ba", "a", "z"}))
	fmt.Println(groupStrings([]string{"a"}))
	fmt.Println(groupStrings([]string{"ab", "bc", "cd"}))
}
```

## 0250 — Count Univalue Subtrees

```go
package main

// LeetCode #250: Count Univalue Subtrees
// https://leetcode.com/problems/count-univalue-subtrees/
// Difficulty: Medium [Paid]
// Time: O(n), Space: O(h)

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countUnivalSubtrees(root *TreeNode) int {
	count := 0
	isUnival(root, &count)
	return count
}

func isUnival(node *TreeNode, count *int) bool {
	if node == nil {
		return true
	}

	leftUni := isUnival(node.Left, count)
	rightUni := isUnival(node.Right, count)

	if !leftUni || !rightUni {
		return false
	}

	if node.Left != nil && node.Left.Val != node.Val {
		return false
	}
	if node.Right != nil && node.Right.Val != node.Val {
		return false
	}

	*count++
	return true
}

func main() {
	root := &TreeNode{5, &TreeNode{1, &TreeNode{5, nil, nil}, &TreeNode{5, nil, nil}}, &TreeNode{5, nil, &TreeNode{5, nil, nil}}}
	fmt.Println(countUnivalSubtrees(root))

	root2 := &TreeNode{5, &TreeNode{5, &TreeNode{5, nil, nil}, &TreeNode{5, nil, nil}}, &TreeNode{5, nil, nil}}
	fmt.Println(countUnivalSubtrees(root2))

	fmt.Println(countUnivalSubtrees(nil))
}
```

## 0251 — Flatten 2D Vector

```go
package main

// LeetCode #251: Flatten 2D Vector
// https://leetcode.com/problems/flatten-2d-vector/
// Difficulty: Medium [Paid]
// Time: O(1) amortized per next/hasNext, Space: O(1) excluding input

import "fmt"

type Vector2D struct {
	vec    [][]int
	row    int
	col    int
}

func Constructor(vec [][]int) Vector2D {
	return Vector2D{vec, 0, 0}
}

func (this *Vector2D) advance() {
	for this.row < len(this.vec) && this.col >= len(this.vec[this.row]) {
		this.row++
		this.col = 0
	}
}

func (this *Vector2D) Next() int {
	this.advance()
	val := this.vec[this.row][this.col]
	this.col++
	return val
}

func (this *Vector2D) HasNext() bool {
	this.advance()
	return this.row < len(this.vec)
}

func main() {
	iter := Constructor([][]int{{1, 2}, {3}, {4, 5, 6}})
	for iter.HasNext() {
		fmt.Print(iter.Next(), " ")
	}
	fmt.Println()

	iter2 := Constructor([][]int{{}, {1}, {}})
	for iter2.HasNext() {
		fmt.Print(iter2.Next(), " ")
	}
	fmt.Println()

	iter3 := Constructor([][]int{{}})
	fmt.Println(iter3.HasNext())
}
```

## 0253 — Meeting Rooms Ii

```go
package main

// LeetCode #253: Meeting Rooms II
// https://leetcode.com/problems/meeting-rooms-ii/
// Difficulty: Medium [Paid]
// Time: O(n log n), Space: O(n)

import (
	"fmt"
	"sort"
)

func minMeetingRooms(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	starts := make([]int, len(intervals))
	ends := make([]int, len(intervals))

	for i, interval := range intervals {
		starts[i] = interval[0]
		ends[i] = interval[1]
	}

	sort.Ints(starts)
	sort.Ints(ends)

	rooms, endIdx := 0, 0

	for i := 0; i < len(starts); i++ {
		if starts[i] < ends[endIdx] {
			rooms++
		} else {
			endIdx++
		}
	}

	return rooms
}

func main() {
	fmt.Println(minMeetingRooms([][]int{{0, 30}, {5, 10}, {15, 20}}))
	fmt.Println(minMeetingRooms([][]int{{7, 10}, {2, 4}}))
	fmt.Println(minMeetingRooms([][]int{{0, 5}, {5, 10}, {10, 15}}))
}
```

## 0254 — Factor Combinations

```go
package main

// LeetCode #254: Factor Combinations
// https://leetcode.com/problems/factor-combinations/
// Difficulty: Medium [Paid]
// Time: O(n log n), Space: O(log n)

import "fmt"

func getFactors(n int) [][]int {
	result := [][]int{}
	var backtrack func(start, remaining int, path []int)
	backtrack = func(start, remaining int, path []int) {
		if len(path) > 0 {
			combo := make([]int, len(path)+1)
			copy(combo, path)
			combo[len(combo)-1] = remaining
			result = append(result, combo)
		}

		for i := start; i*i <= remaining; i++ {
			if remaining%i == 0 {
				path = append(path, i)
				backtrack(i, remaining/i, path)
				path = path[:len(path)-1]
			}
		}
	}

	backtrack(2, n, []int{})
	return result
}

func main() {
	fmt.Println(getFactors(12))
	fmt.Println(getFactors(37))
	fmt.Println(getFactors(32))
}
```

## 0255 — Verify Preorder Sequence In Binary Search Tree

```go
package main

// LeetCode #255: Verify Preorder Sequence in Binary Search Tree
// https://leetcode.com/problems/verify-preorder-sequence-in-binary-search-tree/
// Difficulty: Medium [Paid]
// Time: O(n), Space: O(n)

import "fmt"

func verifyPreorder(preorder []int) bool {
	stack := []int{}
	lower := ^int(^uint(0) >> 1) // math.MinInt

	for _, val := range preorder {
		if val < lower {
			return false
		}
		for len(stack) > 0 && val > stack[len(stack)-1] {
			lower = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, val)
	}

	return true
}

func main() {
	fmt.Println(verifyPreorder([]int{5, 2, 1, 3, 6}))
	fmt.Println(verifyPreorder([]int{5, 2, 6, 1, 3}))
	fmt.Println(verifyPreorder([]int{1, 2, 3}))
}
```

## 0256 — Paint House

```go
package main

// LeetCode #256: Paint House
// https://leetcode.com/problems/paint-house/
// Difficulty: Medium [Paid]
// Time: O(n), Space: O(1)

import "fmt"

func minCost(costs [][]int) int {
	if len(costs) == 0 {
		return 0
	}

	for i := 1; i < len(costs); i++ {
		costs[i][0] += min(costs[i-1][1], costs[i-1][2])
		costs[i][1] += min(costs[i-1][0], costs[i-1][2])
		costs[i][2] += min(costs[i-1][0], costs[i-1][1])
	}

	last := costs[len(costs)-1]
	return min(last[0], min(last[1], last[2]))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minCost([][]int{{17, 2, 17}, {16, 16, 5}, {14, 3, 19}}))
	fmt.Println(minCost([][]int{{7, 6, 2}}))
	fmt.Println(minCost([][]int{}))
}
```

## 0259 — 3Sum Smaller

```go
package main

// LeetCode #259: 3Sum Smaller
// https://leetcode.com/problems/3sum-smaller/
// Difficulty: Medium [Paid]
// Time: O(n^2), Space: O(1)

import (
	"fmt"
	"sort"
)

func threeSumSmaller(nums []int, target int) int {
	sort.Ints(nums)
	count := 0

	for i := 0; i < len(nums)-2; i++ {
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum < target {
				count += right - left
				left++
			} else {
				right--
			}
		}
	}

	return count
}

func main() {
	fmt.Println(threeSumSmaller([]int{-2, 0, 1, 3}, 2))
	fmt.Println(threeSumSmaller([]int{1, 1, -2}, 1))
	fmt.Println(threeSumSmaller([]int{0, 0, 0}, 0))
}
```

## 0260 — Single Number Iii

```go
package main

// LeetCode #260: Single Number III
// https://leetcode.com/problems/single-number-iii/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

func singleNumber(nums []int) []int {
	xor := 0
	for _, num := range nums {
		xor ^= num
	}

	diff := xor & -xor

	num1, num2 := 0, 0
	for _, num := range nums {
		if num&diff == 0 {
			num1 ^= num
		} else {
			num2 ^= num
		}
	}

	return []int{num1, num2}
}

func main() {
	fmt.Println(singleNumber([]int{1, 2, 1, 3, 2, 5}))
	fmt.Println(singleNumber([]int{-1, 0}))
	fmt.Println(singleNumber([]int{0, 1}))
}
```

## 0261 — Graph Valid Tree

```go
package main

// LeetCode #261: Graph Valid Tree
// https://leetcode.com/problems/graph-valid-tree/
// Difficulty: Medium [Paid]
// Time: O(V+E), Space: O(V+E)

import "fmt"

func validTree(n int, edges [][]int) bool {
	if len(edges) != n-1 {
		return false
	}

	parent := make([]int, n)
	for i := 0; i < n; i++ {
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
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(validTree(5, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}}))
	fmt.Println(validTree(5, [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4}}))
	fmt.Println(validTree(4, [][]int{{0, 1}, {2, 3}}))
}
```

## 0264 — Ugly Number Ii

```go
package main

// LeetCode #264: Ugly Number II
// https://leetcode.com/problems/ugly-number-ii/
// Difficulty: Medium
// Time: O(n), Space: O(n)

import "fmt"

func nthUglyNumber(n int) int {
	ugly := make([]int, n)
	ugly[0] = 1

	p2, p3, p5 := 0, 0, 0

	for i := 1; i < n; i++ {
		next := min(ugly[p2]*2, min(ugly[p3]*3, ugly[p5]*5))
		ugly[i] = next

		if next == ugly[p2]*2 {
			p2++
		}
		if next == ugly[p3]*3 {
			p3++
		}
		if next == ugly[p5]*5 {
			p5++
		}
	}

	return ugly[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(nthUglyNumber(10))
	fmt.Println(nthUglyNumber(1))
	fmt.Println(nthUglyNumber(7))
}
```

## 0267 — Palindrome Permutation Ii

```go
package main

// LeetCode #267: Palindrome Permutation II
// https://leetcode.com/problems/palindrome-permutation-ii/
// Difficulty: Medium [Paid]
// Time: O((n/2)!), Space: O(n)

import "fmt"

func generatePalindromes(s string) []string {
	charCount := make([]byte, 128)
	for i := 0; i < len(s); i++ {
		charCount[s[i]]++
	}

	oddChar := byte(0)
	oddCount := 0
	for i := range charCount {
		if charCount[i]%2 == 1 {
			oddChar = byte(i)
			oddCount++
		}
	}

	if oddCount > 1 {
		return nil
	}

	half := []byte{}
	for i := range charCount {
		for j := 0; j < int(charCount[i])/2; j++ {
			half = append(half, byte(i))
		}
	}

	result := []string{}
	used := make([]bool, len(half))
	var backtrack func(path []byte)
	backtrack = func(path []byte) {
		if len(path) == len(half) {
			pal := string(path)
			rev := ""
			for i := len(path) - 1; i >= 0; i-- {
				rev += string(path[i])
			}
			if oddChar != 0 {
				pal += string(oddChar)
			}
			result = append(result, pal+rev)
			return
		}

		for i := 0; i < len(half); i++ {
			if used[i] {
				continue
			}
			if i > 0 && half[i] == half[i-1] && !used[i-1] {
				continue
			}
			used[i] = true
			path = append(path, half[i])
			backtrack(path)
			path = path[:len(path)-1]
			used[i] = false
		}
	}

	backtrack([]byte{})
	return result
}

func main() {
	fmt.Println(generatePalindromes("aabb"))
	fmt.Println(generatePalindromes("abc"))
	fmt.Println(generatePalindromes("a"))
}
```

## 0271 — Encode And Decode Strings

```go
package main

// LeetCode #271: Encode and Decode Strings
// https://leetcode.com/problems/encode-and-decode-strings/
// Difficulty: Medium [Paid]
// Time: O(n), Space: O(n)

import (
	"fmt"
	"strconv"
	"strings"
)

type Codec struct{}

func (codec *Codec) Encode(strs []string) string {
	var sb strings.Builder
	for _, s := range strs {
		sb.WriteString(strconv.Itoa(len(s)))
		sb.WriteByte('#')
		sb.WriteString(s)
	}
	return sb.String()
}

func (codec *Codec) Decode(s string) []string {
	result := []string{}
	i := 0
	for i < len(s) {
		j := i
		for s[j] != '#' {
			j++
		}
		length, _ := strconv.Atoi(s[i:j])
		start := j + 1
		result = append(result, s[start:start+length])
		i = start + length
	}
	return result
}

func main() {
	codec := &Codec{}
	encoded := codec.Encode([]string{"hello", "world", "leet", "code"})
	fmt.Println(encoded)
	fmt.Println(codec.Decode(encoded))

	encoded2 := codec.Encode([]string{""})
	fmt.Println(codec.Decode(encoded2))

	encoded3 := codec.Encode([]string{"a", "", "b"})
	fmt.Println(codec.Decode(encoded3))
}
```

## 0274 — H Index

```go
package main

// LeetCode #274: H-Index
// https://leetcode.com/problems/h-index/
// Difficulty: Medium
// Time: O(n), Space: O(n)

import "fmt"

func hIndex(citations []int) int {
	n := len(citations)
	buckets := make([]int, n+1)

	for _, c := range citations {
		if c >= n {
			buckets[n]++
		} else {
			buckets[c]++
		}
	}

	count := 0
	for i := n; i >= 0; i-- {
		count += buckets[i]
		if count >= i {
			return i
		}
	}

	return 0
}

func main() {
	fmt.Println(hIndex([]int{3, 0, 6, 1, 5}))
	fmt.Println(hIndex([]int{1, 3, 1}))
	fmt.Println(hIndex([]int{0}))
}
```

## 0275 — H Index Ii

```go
package main

// LeetCode #275: H-Index II
// https://leetcode.com/problems/h-index-ii/
// Difficulty: Medium
// Time: O(log n), Space: O(1)

import "fmt"

func hIndex(citations []int) int {
	n := len(citations)
	left, right := 0, n-1

	for left <= right {
		mid := left + (right-left)/2
		if citations[mid] >= n-mid {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return n - left
}

func main() {
	fmt.Println(hIndex([]int{0, 1, 3, 5, 6}))
	fmt.Println(hIndex([]int{1, 2, 100}))
	fmt.Println(hIndex([]int{0}))
}
```

## 0276 — Paint Fence

```go
package main

// LeetCode #276: Paint Fence
// https://leetcode.com/problems/paint-fence/
// Difficulty: Medium [Paid]
// Time: O(n), Space: O(1)

import "fmt"

func numWays(n int, k int) int {
	if n == 0 || k == 0 {
		return 0
	}
	if n == 1 {
		return k
	}

	same := k
	diff := k * (k - 1)

	for i := 3; i <= n; i++ {
		same, diff = diff, (same+diff)*(k-1)
	}

	return same + diff
}

func main() {
	fmt.Println(numWays(3, 2))
	fmt.Println(numWays(1, 1))
	fmt.Println(numWays(7, 2))
}
```

## 0277 — Find The Celebrity

```go
package main

// LeetCode #277: Find the Celebrity
// https://leetcode.com/problems/find-the-celebrity/
// Difficulty: Medium [Paid]
// Time: O(n), Space: O(1)

import "fmt"

var knowsMatrix [][]int

func knows(a, b int) bool {
	return knowsMatrix[a][b] == 1
}

func findCelebrity(n int) int {
	candidate := 0

	for i := 1; i < n; i++ {
		if knows(candidate, i) {
			candidate = i
		}
	}

	for i := 0; i < n; i++ {
		if i == candidate {
			continue
		}
		if knows(candidate, i) || !knows(i, candidate) {
			return -1
		}
	}

	return candidate
}

func main() {
	knowsMatrix = [][]int{
		{1, 1, 0},
		{0, 1, 0},
		{1, 1, 1},
	}
	fmt.Println(findCelebrity(3))

	knowsMatrix = [][]int{
		{1, 0, 1},
		{1, 1, 0},
		{0, 1, 1},
	}
	fmt.Println(findCelebrity(3))

	knowsMatrix = [][]int{{1, 1}, {0, 1}}
	fmt.Println(findCelebrity(2))
}
```

## 0279 — Perfect Squares

```go
package main

// LeetCode #279: Perfect Squares
// https://leetcode.com/problems/perfect-squares/
// Difficulty: Medium
// Time: O(n * sqrt(n)), Space: O(n)

import (
	"fmt"
	"math"
)

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt32
	}

	for i := 1; i <= n; i++ {
		for j := 1; j*j <= i; j++ {
			sq := j * j
			if 1+dp[i-sq] < dp[i] {
				dp[i] = 1 + dp[i-sq]
			}
		}
	}

	return dp[n]
}

func main() {
	fmt.Println(numSquares(12))
	fmt.Println(numSquares(13))
	fmt.Println(numSquares(1))
}
```

## 0280 — Wiggle Sort

```go
package main

// LeetCode #280: Wiggle Sort
// https://leetcode.com/problems/wiggle-sort/
// Difficulty: Medium [Paid]
// Time: O(n), Space: O(1)

import "fmt"

func wiggleSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		if i%2 == 1 {
			if nums[i] < nums[i-1] {
				nums[i], nums[i-1] = nums[i-1], nums[i]
			}
		} else {
			if nums[i] > nums[i-1] {
				nums[i], nums[i-1] = nums[i-1], nums[i]
			}
		}
	}
}

func main() {
	nums1 := []int{3, 5, 2, 1, 6, 4}
	wiggleSort(nums1)
	fmt.Println(nums1)

	nums2 := []int{1, 2, 3, 4}
	wiggleSort(nums2)
	fmt.Println(nums2)

	nums3 := []int{1}
	wiggleSort(nums3)
	fmt.Println(nums3)
}
```

## 0281 — Zigzag Iterator

```go
package main

// LeetCode #281: Zigzag Iterator
// https://leetcode.com/problems/zigzag-iterator/
// Difficulty: Medium [Paid]
// Time: O(1) amortized per next/hasNext, Space: O(n)

import "fmt"

type ZigzagIterator struct {
	vectors [][]int
	indices []int
	curr    int
	total   int
}

func Constructor(v1, v2 []int) *ZigzagIterator {
	vectors := [][]int{v1, v2}
	indices := make([]int, 2)
	total := 0
	for _, v := range vectors {
		total += len(v)
	}
	return &ZigzagIterator{vectors, indices, 0, total}
}

func (this *ZigzagIterator) next() int {
	for this.indices[this.curr] >= len(this.vectors[this.curr]) {
		this.curr = (this.curr + 1) % 2
	}
	val := this.vectors[this.curr][this.indices[this.curr]]
	this.indices[this.curr]++
	this.curr = (this.curr + 1) % 2
	this.total--
	return val
}

func (this *ZigzagIterator) hasNext() bool {
	return this.total > 0
}

func main() {
	iter := Constructor([]int{1, 2, 3}, []int{4, 5, 6, 7})
	for iter.hasNext() {
		fmt.Print(iter.next(), " ")
	}
	fmt.Println()

	iter2 := Constructor([]int{1}, []int{})
	for iter2.hasNext() {
		fmt.Print(iter2.next(), " ")
	}
	fmt.Println()
}
```

## 0284 — Peeking Iterator

```go
package main

// LeetCode #284: Peeking Iterator
// https://leetcode.com/problems/peeking-iterator/
// Difficulty: Medium
// Time: O(1) per operation, Space: O(1)

import "fmt"

type Iterator struct {
	data []int
	pos  int
}

func (this *Iterator) hasNext() bool {
	return this.pos < len(this.data)
}

func (this *Iterator) next() int {
	val := this.data[this.pos]
	this.pos++
	return val
}

type PeekingIterator struct {
	iter    *Iterator
	hasPeek bool
	peekVal int
}

func Constructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{iter, false, 0}
}

func (this *PeekingIterator) hasNext() bool {
	return this.hasPeek || this.iter.hasNext()
}

func (this *PeekingIterator) next() int {
	if this.hasPeek {
		this.hasPeek = false
		return this.peekVal
	}
	return this.iter.next()
}

func (this *PeekingIterator) peek() int {
	if !this.hasPeek {
		this.peekVal = this.iter.next()
		this.hasPeek = true
	}
	return this.peekVal
}

func main() {
	iter := &Iterator{[]int{1, 2, 3}, 0}
	pIter := Constructor(iter)
	fmt.Println(pIter.peek())
	fmt.Println(pIter.next())
	fmt.Println(pIter.peek())
	fmt.Println(pIter.hasNext())
	fmt.Println(pIter.next())
	fmt.Println(pIter.hasNext())
	fmt.Println(pIter.next())
	fmt.Println(pIter.hasNext())
}
```

## 0285 — Inorder Successor In Bst

```go
package main

// LeetCode #285: Inorder Successor in BST
// https://leetcode.com/problems/inorder-successor-in-bst/
// Difficulty: Medium [Paid]
// Time: O(h), Space: O(1)

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var successor *TreeNode

	for root != nil {
		if p.Val < root.Val {
			successor = root
			root = root.Left
		} else {
			root = root.Right
		}
	}

	return successor
}

func main() {
	root := &TreeNode{5, &TreeNode{3, &TreeNode{2, nil, nil}, &TreeNode{4, nil, nil}}, &TreeNode{6, nil, &TreeNode{7, nil, nil}}}
	p := root.Left
	fmt.Println(inorderSuccessor(root, p).Val)

	p2 := root.Right
	fmt.Println(inorderSuccessor(root, p2))

	p3 := root.Left.Right
	fmt.Println(inorderSuccessor(root, p3).Val)
}
```

## 0286 — Walls And Gates

```go
package main

// LeetCode #286: Walls and Gates
// https://leetcode.com/problems/walls-and-gates/
// Difficulty: Medium [Paid]
// Time: O(m*n), Space: O(m*n) for queue

import "fmt"

func wallsAndGates(rooms [][]int) {
	if len(rooms) == 0 || len(rooms[0]) == 0 {
		return
	}

	rows, cols := len(rooms), len(rooms[0])
	queue := [][2]int{}
	INF := 2147483647

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if rooms[r][c] == 0 {
				queue = append(queue, [2]int{r, c})
			}
		}
	}

	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]

		for _, dir := range dirs {
			r, c := cell[0]+dir[0], cell[1]+dir[1]
			if r >= 0 && r < rows && c >= 0 && c < cols && rooms[r][c] == INF {
				rooms[r][c] = rooms[cell[0]][cell[1]] + 1
				queue = append(queue, [2]int{r, c})
			}
		}
	}
}

func main() {
	INF := 2147483647
	rooms1 := [][]int{
		{INF, -1, 0, INF},
		{INF, INF, INF, -1},
		{INF, -1, INF, -1},
		{0, -1, INF, INF},
	}
	wallsAndGates(rooms1)
	fmt.Println(rooms1)

	rooms2 := [][]int{{INF}}
	wallsAndGates(rooms2)
	fmt.Println(rooms2)
}
```

## 0287 — Find The Duplicate Number

```go
package main

// LeetCode #287: Find the Duplicate Number
// https://leetcode.com/problems/find-the-duplicate-number/
// Difficulty: Medium
// Time: O(n), Space: O(1) using Floyd's Cycle Detection

import "fmt"

func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[0]
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	slow = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}

func main() {
	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 2}))
	fmt.Println(findDuplicate([]int{3, 1, 3, 4, 2}))
	fmt.Println(findDuplicate([]int{3, 3, 3, 3, 3}))
}
```

## 0288 — Unique Word Abbreviation

```go
package main

// LeetCode #288: Unique Word Abbreviation
// https://leetcode.com/problems/unique-word-abbreviation/
// Difficulty: Medium [Paid]
// Time: O(n) for init, O(1) for isUnique, Space: O(n)

import (
	"fmt"
	"strconv"
)

type ValidWordAbbr struct {
	abbrMap map[string]string
}

func Constructor(dictionary []string) ValidWordAbbr {
	abbrMap := make(map[string]string)
	for _, word := range dictionary {
		abbr := getAbbr(word)
		if existing, ok := abbrMap[abbr]; ok {
			if existing != word {
				abbrMap[abbr] = ""
			}
		} else {
			abbrMap[abbr] = word
		}
	}
	return ValidWordAbbr{abbrMap}
}

func (this *ValidWordAbbr) IsUnique(word string) bool {
	abbr := getAbbr(word)
	val, ok := this.abbrMap[abbr]
	return !ok || val == word
}

func getAbbr(s string) string {
	if len(s) <= 2 {
		return s
	}
	return string(s[0]) + strconv.Itoa(len(s)-2) + string(s[len(s)-1])
}

func main() {
	vwa := Constructor([]string{"deer", "door", "cake", "card"})
	fmt.Println(vwa.IsUnique("dear"))
	fmt.Println(vwa.IsUnique("cart"))
	fmt.Println(vwa.IsUnique("cane"))
	fmt.Println(vwa.IsUnique("make"))

	vwa2 := Constructor([]string{"a", "a"})
	fmt.Println(vwa2.IsUnique("a"))
}
```

## 0289 — Game Of Life

```go
package main

// LeetCode #289: Game of Life
// https://leetcode.com/problems/game-of-life/
// Difficulty: Medium
// Time: O(m*n), Space: O(1)

import "fmt"

func gameOfLife(board [][]int) {
	if len(board) == 0 {
		return
	}

	rows, cols := len(board), len(board[0])

	countLive := func(r, c int) int {
		count := 0
		dirs := [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
		for _, d := range dirs {
			nr, nc := r+d[0], c+d[1]
			if nr >= 0 && nr < rows && nc >= 0 && nc < cols && (board[nr][nc] == 1 || board[nr][nc] == 2) {
				count++
			}
		}
		return count
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			live := countLive(r, c)
			if board[r][c] == 1 && (live < 2 || live > 3) {
				board[r][c] = 2
			}
			if board[r][c] == 0 && live == 3 {
				board[r][c] = -1
			}
		}
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if board[r][c] == 2 {
				board[r][c] = 0
			} else if board[r][c] == -1 {
				board[r][c] = 1
			}
		}
	}
}

func main() {
	board1 := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	gameOfLife(board1)
	fmt.Println(board1)

	board2 := [][]int{{1, 1}, {1, 0}}
	gameOfLife(board2)
	fmt.Println(board2)
}
```

## 0291 — Word Pattern Ii

```go
package main

// LeetCode #291: Word Pattern II
// https://leetcode.com/problems/word-pattern-ii/
// Difficulty: Medium [Paid]
// Time: O(2^n) worst case, Space: O(n)

import "fmt"

func wordPatternMatch(pattern string, s string) bool {
	pMap := make(map[byte]string)
	sMap := make(map[string]byte)

	var backtrack func(patIdx, strIdx int) bool
	backtrack = func(patIdx, strIdx int) bool {
		if patIdx == len(pattern) && strIdx == len(s) {
			return true
		}
		if patIdx >= len(pattern) || strIdx >= len(s) {
			return false
		}

		ch := pattern[patIdx]
		if mapped, ok := pMap[ch]; ok {
			if strIdx+len(mapped) > len(s) || s[strIdx:strIdx+len(mapped)] != mapped {
				return false
			}
			return backtrack(patIdx+1, strIdx+len(mapped))
		}

		for end := strIdx + 1; end <= len(s); end++ {
			candidate := s[strIdx:end]
			if existing, ok := sMap[candidate]; ok && existing != ch {
				continue
			}

			pMap[ch] = candidate
			sMap[candidate] = ch
			if backtrack(patIdx+1, end) {
				return true
			}
			delete(pMap, ch)
			delete(sMap, candidate)
		}

		return false
	}

	return backtrack(0, 0)
}

func main() {
	fmt.Println(wordPatternMatch("abab", "redblueredblue"))
	fmt.Println(wordPatternMatch("aaaa", "asdasdasdasd"))
	fmt.Println(wordPatternMatch("ab", "aa"))
}
```

## 0294 — Flip Game Ii

```go
package main

// LeetCode #294: Flip Game II
// https://leetcode.com/problems/flip-game-ii/
// Difficulty: Medium [Paid]
// Time: O(n!!) worst case with memo, Space: O(n!)

import "fmt"

func canWin(currentState string) bool {
	memo := make(map[string]bool)
	return canWinHelper(currentState, memo)
}

func canWinHelper(state string, memo map[string]bool) bool {
	if res, ok := memo[state]; ok {
		return res
	}

	bytes := []byte(state)
	for i := 0; i < len(state)-1; i++ {
		if bytes[i] == '+' && bytes[i+1] == '+' {
			bytes[i], bytes[i+1] = '-', '-'
			opponentWins := canWinHelper(string(bytes), memo)
			bytes[i], bytes[i+1] = '+', '+'

			if !opponentWins {
				memo[state] = true
				return true
			}
		}
	}

	memo[state] = false
	return false
}

func main() {
	fmt.Println(canWin("++++"))
	fmt.Println(canWin("+"))
	fmt.Println(canWin("+++"))
}
```

## 0298 — Binary Tree Longest Consecutive Sequence

```go
package main

// LeetCode #298: Binary Tree Longest Consecutive Sequence
// https://leetcode.com/problems/binary-tree-longest-consecutive-sequence/
// Difficulty: Medium [Paid]
// Time: O(n), Space: O(h)

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestConsecutive(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxLen := 1
	var dfs func(node *TreeNode, parentVal int, length int)
	dfs = func(node *TreeNode, parentVal int, length int) {
		if node == nil {
			return
		}

		if node.Val == parentVal+1 {
			length++
		} else {
			length = 1
		}

		if length > maxLen {
			maxLen = length
		}

		dfs(node.Left, node.Val, length)
		dfs(node.Right, node.Val, length)
	}

	dfs(root, root.Val, 1)
	return maxLen
}

func main() {
	root := &TreeNode{1, nil, &TreeNode{3, &TreeNode{2, nil, nil}, &TreeNode{4, nil, &TreeNode{5, nil, nil}}}}
	fmt.Println(longestConsecutive(root))

	root2 := &TreeNode{2, nil, &TreeNode{3, &TreeNode{2, &TreeNode{1, nil, nil}, nil}, nil}}
	fmt.Println(longestConsecutive(root2))

	fmt.Println(longestConsecutive(nil))
}
```

## 0299 — Bulls And Cows

```go
package main

// LeetCode #299: Bulls and Cows
// https://leetcode.com/problems/bulls-and-cows/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import (
	"fmt"
	"strconv"
)

func getHint(secret string, guess string) string {
	bulls, cows := 0, 0
	secretCount := [10]int{}
	guessCount := [10]int{}

	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			bulls++
		} else {
			secretCount[secret[i]-'0']++
			guessCount[guess[i]-'0']++
		}
	}

	for i := 0; i < 10; i++ {
		if secretCount[i] < guessCount[i] {
			cows += secretCount[i]
		} else {
			cows += guessCount[i]
		}
	}

	return strconv.Itoa(bulls) + "A" + strconv.Itoa(cows) + "B"
}

func main() {
	fmt.Println(getHint("1807", "7810"))
	fmt.Println(getHint("1123", "0111"))
	fmt.Println(getHint("1", "0"))
}
```

## 0300 — Longest Increasing Subsequence

```go
package main

// LeetCode #300: Longest Increasing Subsequence
// https://leetcode.com/problems/longest-increasing-subsequence/
// Difficulty: Medium
// Time: O(n log n), Space: O(n)

import "fmt"

func lengthOfLIS(nums []int) int {
	tails := []int{}

	for _, num := range nums {
		left, right := 0, len(tails)
		for left < right {
			mid := left + (right-left)/2
			if tails[mid] < num {
				left = mid + 1
			} else {
				right = mid
			}
		}

		if left == len(tails) {
			tails = append(tails, num)
		} else {
			tails[left] = num
		}
	}

	return len(tails)
}

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	fmt.Println(lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))
	fmt.Println(lengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7}))
}
```

## 0304 — Range Sum Query 2D Immutable

```go
package main

// LeetCode #304: Range Sum Query 2D - Immutable
// https://leetcode.com/problems/range-sum-query-2d-immutable/
// Difficulty: Medium
// Time: O(m*n) for init, O(1) for sum, Space: O(m*n)

import "fmt"

type NumMatrix struct {
	prefix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return NumMatrix{}
	}

	rows, cols := len(matrix), len(matrix[0])
	prefix := make([][]int, rows+1)
	for i := range prefix {
		prefix[i] = make([]int, cols+1)
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			prefix[r+1][c+1] = matrix[r][c] + prefix[r][c+1] + prefix[r+1][c] - prefix[r][c]
		}
	}

	return NumMatrix{prefix}
}

func (this *NumMatrix) SumRegion(row1, col1, row2, col2 int) int {
	return this.prefix[row2+1][col2+1] - this.prefix[row1][col2+1] - this.prefix[row2+1][col1] + this.prefix[row1][col1]
}

func main() {
	matrix := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}
	nm := Constructor(matrix)
	fmt.Println(nm.SumRegion(2, 1, 4, 3))
	fmt.Println(nm.SumRegion(1, 1, 2, 2))
	fmt.Println(nm.SumRegion(1, 2, 2, 4))
}
```

## 0306 — Additive Number

```go
package main

// LeetCode #306: Additive Number
// https://leetcode.com/problems/additive-number/
// Difficulty: Medium
// Time: O(n^2), Space: O(n)

import (
	"fmt"
	"strconv"
	"strings"
)

func isAdditiveNumber(num string) bool {
	n := len(num)

	for firstEnd := 1; firstEnd <= n/2; firstEnd++ {
		if num[0] == '0' && firstEnd > 1 {
			break
		}
		num1, _ := strconv.ParseInt(num[:firstEnd], 10, 64)

		for secondEnd := firstEnd + 1; max(firstEnd, secondEnd-firstEnd) <= n-secondEnd; secondEnd++ {
			if num[firstEnd] == '0' && secondEnd-firstEnd > 1 {
				break
			}
			num2, _ := strconv.ParseInt(num[firstEnd:secondEnd], 10, 64)

			if isValid(num, num1, num2, secondEnd) {
				return true
			}
		}
	}

	return false
}

func isValid(num string, num1, num2 int64, start int) bool {
	if start == len(num) {
		return false
	}

	for start < len(num) {
		sum := num1 + num2
		sumStr := strconv.FormatInt(sum, 10)

		if !strings.HasPrefix(num[start:], sumStr) {
			return false
		}

		start += len(sumStr)
		num1, num2 = num2, sum
	}

	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(isAdditiveNumber("112358"))
	fmt.Println(isAdditiveNumber("199100199"))
	fmt.Println(isAdditiveNumber("12"))
}
```

## 0307 — Range Sum Query Mutable

```go
package main

// LeetCode #307: Range Sum Query - Mutable
// https://leetcode.com/problems/range-sum-query-mutable/
// Difficulty: Medium
// Time: O(log n) per operation, Space: O(n)

import "fmt"

type NumArray struct {
	nums []int
	tree []int
	n    int
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	tree := make([]int, n+1)
	na := NumArray{nums, tree, n}
	for i, val := range nums {
		na.add(i+1, val)
	}
	return na
}

func (this *NumArray) add(idx, val int) {
	for idx <= this.n {
		this.tree[idx] += val
		idx += idx & -idx
	}
}

func (this *NumArray) sum(idx int) int {
	res := 0
	for idx > 0 {
		res += this.tree[idx]
		idx -= idx & -idx
	}
	return res
}

func (this *NumArray) Update(index int, val int) {
	diff := val - this.nums[index]
	this.nums[index] = val
	this.add(index+1, diff)
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.sum(right+1) - this.sum(left)
}

func main() {
	na := Constructor([]int{1, 3, 5})
	fmt.Println(na.SumRange(0, 2))
	na.Update(1, 2)
	fmt.Println(na.SumRange(0, 2))

	na2 := Constructor([]int{1})
	fmt.Println(na2.SumRange(0, 0))
	na2.Update(0, 5)
	fmt.Println(na2.SumRange(0, 0))
}
```

## 0308 — Range Sum Query 2D Mutable

```go
package main

// LeetCode #308: Range Sum Query 2D - Mutable
// https://leetcode.com/problems/range-sum-query-2d-mutable/
// Difficulty: Medium [Paid]
// Time: O(log m * log n) for update/sum, Space: O(m*n)

import "fmt"

type NumMatrix struct {
	matrix [][]int
	bit    [][]int
	rows   int
	cols   int
}

func Constructor(matrix [][]int) NumMatrix {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return NumMatrix{}
	}

	rows, cols := len(matrix), len(matrix[0])
	bit := make([][]int, rows+1)
	for i := range bit {
		bit[i] = make([]int, cols+1)
	}

	nm := NumMatrix{matrix, bit, rows, cols}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			nm.add(r+1, c+1, matrix[r][c])
		}
	}
	return nm
}

func (this *NumMatrix) add(row, col, val int) {
	for r := row; r <= this.rows; r += r & -r {
		for c := col; c <= this.cols; c += c & -c {
			this.bit[r][c] += val
		}
	}
}

func (this *NumMatrix) sum(row, col int) int {
	res := 0
	for r := row; r > 0; r -= r & -r {
		for c := col; c > 0; c -= c & -c {
			res += this.bit[r][c]
		}
	}
	return res
}

func (this *NumMatrix) Update(row int, col int, val int) {
	diff := val - this.matrix[row][col]
	this.matrix[row][col] = val
	this.add(row+1, col+1, diff)
}

func (this *NumMatrix) SumRegion(row1, col1, row2, col2 int) int {
	return this.sum(row2+1, col2+1) - this.sum(row1, col2+1) - this.sum(row2+1, col1) + this.sum(row1, col1)
}

func main() {
	matrix := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}
	nm := Constructor(matrix)
	fmt.Println(nm.SumRegion(2, 1, 4, 3))
	nm.Update(3, 2, 2)
	fmt.Println(nm.SumRegion(2, 1, 4, 3))
}
```

## 0309 — Best Time To Buy And Sell Stock With Cooldown

```go
package main

// LeetCode #309: Best Time to Buy and Sell Stock with Cooldown
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	hold := -prices[0]
	cool := 0
	sell := 0

	for i := 1; i < len(prices); i++ {
		prevSell := sell
		sell = max(sell, hold+prices[i])
		hold = max(hold, cool-prices[i])
		cool = max(cool, prevSell)
	}

	return max(sell, cool)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxProfit([]int{1, 2, 3, 0, 2}))
	fmt.Println(maxProfit([]int{1}))
	fmt.Println(maxProfit([]int{1, 2, 4}))
}
```

## 0310 — Minimum Height Trees

```go
package main

// LeetCode #310: Minimum Height Trees
// https://leetcode.com/problems/minimum-height-trees/
// Difficulty: Medium
// Time: O(V), Space: O(V+E)

import "fmt"

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	adj := make([]map[int]bool, n)
	for i := 0; i < n; i++ {
		adj[i] = make(map[int]bool)
	}

	for _, edge := range edges {
		adj[edge[0]][edge[1]] = true
		adj[edge[1]][edge[0]] = true
	}

	leaves := []int{}
	for i := 0; i < n; i++ {
		if len(adj[i]) == 1 {
			leaves = append(leaves, i)
		}
	}

	remaining := n
	for remaining > 2 {
		remaining -= len(leaves)
		newLeaves := []int{}

		for _, leaf := range leaves {
			for neighbor := range adj[leaf] {
				delete(adj[neighbor], leaf)
				if len(adj[neighbor]) == 1 {
					newLeaves = append(newLeaves, neighbor)
				}
			}
		}

		leaves = newLeaves
	}

	return leaves
}

func main() {
	fmt.Println(findMinHeightTrees(4, [][]int{{1, 0}, {1, 2}, {1, 3}}))
	fmt.Println(findMinHeightTrees(6, [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}}))
	fmt.Println(findMinHeightTrees(1, [][]int{}))
}
```

## 0311 — Sparse Matrix Multiplication

```go
package main

// LeetCode #311: Sparse Matrix Multiplication
// https://leetcode.com/problems/sparse-matrix-multiplication/
// Difficulty: Medium [Paid]
// Time: O(m*n*k), Space: O(m*k) optimized for sparse matrices

import "fmt"

func multiply(mat1 [][]int, mat2 [][]int) [][]int {
	m, k, n := len(mat1), len(mat1[0]), len(mat2[0])
	result := make([][]int, m)
	for i := range result {
		result[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for kk := 0; kk < k; kk++ {
			if mat1[i][kk] != 0 {
				for j := 0; j < n; j++ {
					if mat2[kk][j] != 0 {
						result[i][j] += mat1[i][kk] * mat2[kk][j]
					}
				}
			}
		}
	}

	return result
}

func main() {
	mat1 := [][]int{{1, 0, 0}, {-1, 0, 3}}
	mat2 := [][]int{{7, 0, 0}, {0, 0, 0}, {0, 0, 1}}
	fmt.Println(multiply(mat1, mat2))

	mat1 = [][]int{{0}}
	mat2 = [][]int{{0}}
	fmt.Println(multiply(mat1, mat2))

	mat1 = [][]int{{1, -5}}
	mat2 = [][]int{{12}, {-1}}
	fmt.Println(multiply(mat1, mat2))
}
```

## 0313 — Super Ugly Number

```go
package main

// LeetCode #313: Super Ugly Number
// https://leetcode.com/problems/super-ugly-number/
// Difficulty: Medium
// Time: O(n * len(primes)), Space: O(n + len(primes))

import "fmt"

func nthSuperUglyNumber(n int, primes []int) int {
	ugly := make([]int, n)
	ugly[0] = 1

	pointers := make([]int, len(primes))
	values := make([]int, len(primes))
	for i, p := range primes {
		values[i] = p
	}

	for i := 1; i < n; i++ {
		minVal := values[0]
		for _, v := range values {
			if v < minVal {
				minVal = v
			}
		}
		ugly[i] = minVal

		for j := range values {
			if values[j] == minVal {
				pointers[j]++
				values[j] = ugly[pointers[j]] * primes[j]
			}
		}
	}

	return ugly[n-1]
}

func main() {
	fmt.Println(nthSuperUglyNumber(12, []int{2, 7, 13, 19}))
	fmt.Println(nthSuperUglyNumber(1, []int{2, 3, 5}))
	fmt.Println(nthSuperUglyNumber(6, []int{2, 3, 5}))
}
```

## 0314 — Binary Tree Vertical Order Traversal

```go
package main

// LeetCode #314: Binary Tree Vertical Order Traversal
// https://leetcode.com/problems/binary-tree-vertical-order-traversal/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type queueItem struct {
	node *TreeNode
	col  int
}

func verticalOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	colMap := make(map[int][]int)
	minCol, maxCol := 0, 0
	queue := []queueItem{{root, 0}}

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]
		node, col := item.node, item.col

		colMap[col] = append(colMap[col], node.Val)

		if col < minCol {
			minCol = col
		}
		if col > maxCol {
			maxCol = col
		}

		if node.Left != nil {
			queue = append(queue, queueItem{node.Left, col - 1})
		}
		if node.Right != nil {
			queue = append(queue, queueItem{node.Right, col + 1})
		}
	}

	result := make([][]int, 0, maxCol-minCol+1)
	for c := minCol; c <= maxCol; c++ {
		result = append(result, colMap[c])
	}
	return result
}

func main() {
	// Test case 1: Example tree [3,9,20,null,null,15,7]
	root1 := &TreeNode{Val: 3}
	root1.Left = &TreeNode{Val: 9}
	root1.Right = &TreeNode{Val: 20}
	root1.Right.Left = &TreeNode{Val: 15}
	root1.Right.Right = &TreeNode{Val: 7}
	fmt.Println("Test 1:", verticalOrder(root1))
	// Expected: [[9],[3,15],[20],[7]]

	// Test case 2: Single node
	root2 := &TreeNode{Val: 1}
	fmt.Println("Test 2:", verticalOrder(root2))
	// Expected: [[1]]

	// Test case 3: Nil
	fmt.Println("Test 3:", verticalOrder(nil))
	// Expected: []
}
```

## 0316 — Remove Duplicate Letters

```go
package main

// LeetCode #316: Remove Duplicate Letters
// https://leetcode.com/problems/remove-duplicate-letters/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func removeDuplicateLetters(s string) string {
	// Count last occurrence of each character
	lastOccur := [26]int{}
	for i := range s {
		lastOccur[s[i]-'a'] = i
	}

	stack := make([]byte, 0, len(s))
	seen := [26]bool{}

	for i := 0; i < len(s); i++ {
		ch := s[i]
		if seen[ch-'a'] {
			continue
		}

		// Pop while stack top is greater and appears later
		for len(stack) > 0 && ch < stack[len(stack)-1] && i < lastOccur[stack[len(stack)-1]-'a'] {
			seen[stack[len(stack)-1]-'a'] = false
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, ch)
		seen[ch-'a'] = true
	}

	return string(stack)
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", removeDuplicateLetters("bcabc"))
	// Expected: "abc"

	// Test case 2
	fmt.Println("Test 2:", removeDuplicateLetters("cbacdcbc"))
	// Expected: "acdb"

	// Test case 3: Single character
	fmt.Println("Test 3:", removeDuplicateLetters("aaaa"))
	// Expected: "a"
}
```

## 0318 — Maximum Product Of Word Lengths

```go
package main

// LeetCode #318: Maximum Product of Word Lengths
// https://leetcode.com/problems/maximum-product-of-word-lengths/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import "fmt"

func maxProduct(words []string) int {
	bits := make([]int, len(words))
	for i, w := range words {
		mask := 0
		for _, ch := range w {
			mask |= 1 << (ch - 'a')
		}
		bits[i] = mask
	}

	maxProd := 0
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if bits[i]&bits[j] == 0 {
				prod := len(words[i]) * len(words[j])
				if prod > maxProd {
					maxProd = prod
				}
			}
		}
	}
	return maxProd
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", maxProduct([]string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}))
	// Expected: 16 (abcw * xtfn)

	// Test case 2
	fmt.Println("Test 2:", maxProduct([]string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}))
	// Expected: 4 (ab * cd)

	// Test case 3
	fmt.Println("Test 3:", maxProduct([]string{"a", "aa", "aaa", "aaaa"}))
	// Expected: 0
}
```

## 0319 — Bulb Switcher

```go
package main

// LeetCode #319: Bulb Switcher
// https://leetcode.com/problems/bulb-switcher/
// Difficulty: Medium
// Time: O(1) | Space: O(1)

import (
	"fmt"
	"math"
)

func bulbSwitch(n int) int {
	// Only bulbs at perfect square positions are toggled an odd number of times
	return int(math.Sqrt(float64(n)))
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", bulbSwitch(3))
	// Expected: 1

	// Test case 2
	fmt.Println("Test 2:", bulbSwitch(0))
	// Expected: 0

	// Test case 3
	fmt.Println("Test 3:", bulbSwitch(25))
	// Expected: 5
}
```

## 0320 — Generalized Abbreviation

```go
package main

// LeetCode #320: Generalized Abbreviation
// https://leetcode.com/problems/generalized-abbreviation/
// Difficulty: Medium [Paid]
// Time: O(n * 2^n) | Space: O(n * 2^n)

import (
	"fmt"
	"strconv"
)

func generateAbbreviations(word string) []string {
	result := []string{}
	backtrack(word, 0, 0, "", &result)
	return result
}

func backtrack(word string, index int, count int, cur string, result *[]string) {
	if index == len(word) {
		if count > 0 {
			cur += strconv.Itoa(count)
		}
		*result = append(*result, cur)
		return
	}

	// Abbreviate current character (increase count)
	backtrack(word, index+1, count+1, cur, result)

	// Keep current character
	if count > 0 {
		cur += strconv.Itoa(count)
	}
	cur += string(word[index])
	backtrack(word, index+1, 0, cur, result)
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", generateAbbreviations("word"))
	// Expected: ["word", "1ord", "w1rd", "2rd", "wo1d", "1o1d", "w2d", "3d", "wor1", "1or1", "w1r1", "2r1", "wo2", "1o2", "w3", "4"]

	// Test case 2: Empty string
	fmt.Println("Test 2:", generateAbbreviations(""))
	// Expected: [""]

	// Test case 3: Single character
	fmt.Println("Test 3:", generateAbbreviations("a"))
	// Expected: ["a", "1"]
}
```

## 0322 — Coin Change

```go
package main

// LeetCode #322: Coin Change
// https://leetcode.com/problems/coin-change/
// Difficulty: Medium
// Time: O(n * amount) | Space: O(amount)

import "fmt"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
		for _, c := range coins {
			if c <= i && dp[i-c]+1 < dp[i] {
				dp[i] = dp[i-c] + 1
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", coinChange([]int{1, 2, 5}, 11))
	// Expected: 3 (5+5+1)

	// Test case 2
	fmt.Println("Test 2:", coinChange([]int{2}, 3))
	// Expected: -1

	// Test case 3
	fmt.Println("Test 3:", coinChange([]int{1}, 0))
	// Expected: 0
}
```

## 0323 — Number Of Connected Components In An Undirected Graph

```go
package main

// LeetCode #323: Number of Connected Components in an Undirected Graph
// https://leetcode.com/problems/number-of-connected-components-in-an-undirected-graph/
// Difficulty: Medium [Paid]
// Time: O(V + E) | Space: O(V)

import "fmt"

func countComponents(n int, edges [][]int) int {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}

	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func(a, b int) {
		ra, rb := find(a), find(b)
		if ra != rb {
			parent[ra] = rb
			n--
		}
	}

	for _, e := range edges {
		union(e[0], e[1])
	}
	return n
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", countComponents(5, [][]int{{0, 1}, {1, 2}, {3, 4}}))
	// Expected: 2

	// Test case 2
	fmt.Println("Test 2:", countComponents(5, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}))
	// Expected: 1

	// Test case 3: No edges
	fmt.Println("Test 3:", countComponents(3, [][]int{}))
	// Expected: 3
}
```

## 0324 — Wiggle Sort Ii

```go
package main

// LeetCode #324: Wiggle Sort II
// https://leetcode.com/problems/wiggle-sort-ii/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func wiggleSort(nums []int) {
	n := len(nums)
	sorted := make([]int, n)
	copy(sorted, nums)
	sort.Ints(sorted)

	// Fill from end of sorted array into odd positions first, then even
	mid := (n + 1) / 2
	j, k := mid-1, n-1
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			nums[i] = sorted[j]
			j--
		} else {
			nums[i] = sorted[k]
			k--
		}
	}
}

func main() {
	// Test case 1
	nums1 := []int{1, 5, 1, 1, 6, 4}
	wiggleSort(nums1)
	fmt.Println("Test 1:", nums1)

	// Test case 2
	nums2 := []int{1, 3, 2, 2, 3, 1}
	wiggleSort(nums2)
	fmt.Println("Test 2:", nums2)

	// Test case 3
	nums3 := []int{1, 2, 3}
	wiggleSort(nums3)
	fmt.Println("Test 3:", nums3)
}
```

## 0325 — Maximum Size Subarray Sum Equals K

```go
package main

// LeetCode #325: Maximum Size Subarray Sum Equals k
// https://leetcode.com/problems/maximum-size-subarray-sum-equals-k/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func maxSubArrayLen(nums []int, k int) int {
	prefixSum := 0
	maxLen := 0
	// Map prefix sum -> earliest index
	sumMap := map[int]int{0: -1}

	for i, num := range nums {
		prefixSum += num

		if idx, ok := sumMap[prefixSum-k]; ok {
			if i-idx > maxLen {
				maxLen = i - idx
			}
		}

		// Only store first occurrence (earliest index)
		if _, ok := sumMap[prefixSum]; !ok {
			sumMap[prefixSum] = i
		}
	}
	return maxLen
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", maxSubArrayLen([]int{1, -1, 5, -2, 3}, 3))
	// Expected: 4 ([1,-1,5,-2])

	// Test case 2
	fmt.Println("Test 2:", maxSubArrayLen([]int{-2, -1, 2, 1}, 1))
	// Expected: 2 ([-1,2])

	// Test case 3
	fmt.Println("Test 3:", maxSubArrayLen([]int{1, 2, 3}, 6))
	// Expected: 3 ([1,2,3])
}
```

## 0328 — Odd Even Linked List

```go
package main

// LeetCode #328: Odd Even Linked List
// https://leetcode.com/problems/odd-even-linked-list/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	odd, even := head, head.Next
	evenHead := even

	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}

	odd.Next = evenHead
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
	// Test case 1: 1->2->3->4->5
	head1 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	fmt.Print("Test 1: ")
	printList(oddEvenList(head1))
	// Expected: 1->3->5->2->4

	// Test case 2: 2->1->3->5->6->4->7
	head2 := &ListNode{2, &ListNode{1, &ListNode{3, &ListNode{5, &ListNode{6, &ListNode{4, &ListNode{7, nil}}}}}}}
	fmt.Print("Test 2: ")
	printList(oddEvenList(head2))
	// Expected: 2->3->6->7->1->5->4

	// Test case 3: Single node
	head3 := &ListNode{1, nil}
	fmt.Print("Test 3: ")
	printList(oddEvenList(head3))
	// Expected: 1
}
```

## 0331 — Verify Preorder Serialization Of A Binary Tree

```go
package main

// LeetCode #331: Verify Preorder Serialization of a Binary Tree
// https://leetcode.com/problems/verify-preorder-serialization-of-a-binary-tree/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"strings"
)

func isValidSerialization(preorder string) bool {
	nodes := strings.Split(preorder, ",")
	// Each non-null node consumes 1 slot and creates 2 new slots
	// Each null node consumes 1 slot
	slots := 1
	for _, node := range nodes {
		slots--
		if slots < 0 {
			return false
		}
		if node != "#" {
			slots += 2
		}
	}
	return slots == 0
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", isValidSerialization("9,3,4,#,#,1,#,#,2,#,6,#,#"))
	// Expected: true

	// Test case 2
	fmt.Println("Test 2:", isValidSerialization("1,#"))
	// Expected: false

	// Test case 3
	fmt.Println("Test 3:", isValidSerialization("9,#,#,1"))
	// Expected: false
}
```

## 0333 — Largest Bst Subtree

```go
package main

// LeetCode #333: Largest BST Subtree
// https://leetcode.com/problems/largest-bst-subtree/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(h)

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type bstInfo struct {
	min, max int
	size     int
	isBST    bool
}

func largestBSTSubtree(root *TreeNode) int {
	maxSize := 0
	postOrder(root, &maxSize)
	return maxSize
}

func postOrder(node *TreeNode, maxSize *int) bstInfo {
	if node == nil {
		return bstInfo{1<<31 - 1, -1 << 31, 0, true}
	}

	left := postOrder(node.Left, maxSize)
	right := postOrder(node.Right, maxSize)

	info := bstInfo{}
	if left.isBST && right.isBST && node.Val > left.max && node.Val < right.min {
		info.min = min(node.Val, left.min)
		if left.size == 0 {
			info.min = node.Val
		}
		info.max = max(node.Val, right.max)
		if right.size == 0 {
			info.max = node.Val
		}
		info.size = left.size + right.size + 1
		info.isBST = true
		if info.size > *maxSize {
			*maxSize = info.size
		}
	} else {
		info.isBST = false
		info.size = max(left.size, right.size)
	}
	return info
}

func main() {
	// Test case 1: [10,5,15,1,8,null,7]
	root1 := &TreeNode{Val: 10}
	root1.Left = &TreeNode{Val: 5, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 8}}
	root1.Right = &TreeNode{Val: 15, Right: &TreeNode{Val: 7}}
	fmt.Println("Test 1:", largestBSTSubtree(root1))
	// Expected: 3

	// Test case 2: Single node
	root2 := &TreeNode{Val: 1}
	fmt.Println("Test 2:", largestBSTSubtree(root2))
	// Expected: 1

	// Test case 3: Nil
	fmt.Println("Test 3:", largestBSTSubtree(nil))
	// Expected: 0
}
```

## 0334 — Increasing Triplet Subsequence

```go
package main

// LeetCode #334: Increasing Triplet Subsequence
// https://leetcode.com/problems/increasing-triplet-subsequence/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func increasingTriplet(nums []int) bool {
	first, second := 1<<31-1, 1<<31-1
	for _, num := range nums {
		if num <= first {
			first = num
		} else if num <= second {
			second = num
		} else {
			return true
		}
	}
	return false
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", increasingTriplet([]int{1, 2, 3, 4, 5}))
	// Expected: true

	// Test case 2
	fmt.Println("Test 2:", increasingTriplet([]int{5, 4, 3, 2, 1}))
	// Expected: false

	// Test case 3
	fmt.Println("Test 3:", increasingTriplet([]int{2, 1, 5, 0, 4, 6}))
	// Expected: true
}
```

## 0337 — House Robber Iii

```go
package main

// LeetCode #337: House Robber III
// https://leetcode.com/problems/house-robber-iii/
// Difficulty: Medium
// Time: O(n) | Space: O(h)

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	robRoot, skipRoot := dfs(root)
	if robRoot > skipRoot {
		return robRoot
	}
	return skipRoot
}

// Returns (robThis, skipThis)
func dfs(node *TreeNode) (int, int) {
	if node == nil {
		return 0, 0
	}

	leftRob, leftSkip := dfs(node.Left)
	rightRob, rightSkip := dfs(node.Right)

	// If we rob this node, we must skip children
	robThis := node.Val + leftSkip + rightSkip
	// If we skip this node, we can take best of each child
	skipThis := max(leftRob, leftSkip) + max(rightRob, rightSkip)

	return robThis, skipThis
}

func main() {
	// Test case 1: [3,2,3,null,3,null,1]
	root1 := &TreeNode{Val: 3}
	root1.Left = &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}
	root1.Right = &TreeNode{Val: 3, Right: &TreeNode{Val: 1}}
	fmt.Println("Test 1:", rob(root1))
	// Expected: 7 (3+3+1)

	// Test case 2: [3,4,5,1,3,null,1]
	root2 := &TreeNode{Val: 3}
	root2.Left = &TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}
	root2.Right = &TreeNode{Val: 5, Right: &TreeNode{Val: 1}}
	fmt.Println("Test 2:", rob(root2))
	// Expected: 9

	// Test case 3: Single node
	root3 := &TreeNode{Val: 10}
	fmt.Println("Test 3:", rob(root3))
	// Expected: 10
}
```

## 0339 — Nested List Weight Sum

```go
package main

// LeetCode #339: Nested List Weight Sum
// https://leetcode.com/problems/nested-list-weight-sum/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(d)

import "fmt"

// NestedInteger represents a nested integer structure (LeetCode interface)
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

func (n NestedInteger) IsInteger() bool           { return n.isInt }
func (n NestedInteger) GetInteger() int            { return n.integer }
func (n NestedInteger) GetList() []*NestedInteger  { return n.list }

func depthSum(nestedList []*NestedInteger) int {
	var dfs func(list []*NestedInteger, depth int) int
	dfs = func(list []*NestedInteger, depth int) int {
		total := 0
		for _, ni := range list {
			if ni.IsInteger() {
				total += ni.GetInteger() * depth
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
	fmt.Println("Test 1:", depthSum([]*NestedInteger{n1, n2, n3}))
	// Expected: 10 (2*1 + 1*2 + 1*2 + 1*2 + 1*2)

	// Test case 2: [1,[4,[6]]]
	inner := NewList(NewInt(6))
	mid := NewList(NewInt(4), inner)
	fmt.Println("Test 2:", depthSum([]*NestedInteger{NewInt(1), mid}))
	// Expected: 27 (1*1 + 4*2 + 6*3)

	// Test case 3: Single integer
	fmt.Println("Test 3:", depthSum([]*NestedInteger{NewInt(5)}))
	// Expected: 5
}
```

## 0340 — Longest Substring With At Most K Distinct Characters

```go
package main

// LeetCode #340: Longest Substring with At Most K Distinct Characters
// https://leetcode.com/problems/longest-substring-with-at-most-k-distinct-characters/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(k)

import "fmt"

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	if k == 0 {
		return 0
	}

	freq := make(map[byte]int)
	left, maxLen := 0, 0

	for right := 0; right < len(s); right++ {
		freq[s[right]]++

		for len(freq) > k {
			freq[s[left]]--
			if freq[s[left]] == 0 {
				delete(freq, s[left])
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
	// Test case 1
	fmt.Println("Test 1:", lengthOfLongestSubstringKDistinct("eceba", 2))
	// Expected: 3 ("ece")

	// Test case 2
	fmt.Println("Test 2:", lengthOfLongestSubstringKDistinct("aa", 1))
	// Expected: 2 ("aa")

	// Test case 3
	fmt.Println("Test 3:", lengthOfLongestSubstringKDistinct("a", 0))
	// Expected: 0
}
```

## 0341 — Flatten Nested List Iterator

```go
package main

// LeetCode #341: Flatten Nested List Iterator
// https://leetcode.com/problems/flatten-nested-list-iterator/
// Difficulty: Medium
// Time: O(n) initialization, O(1) next/hasNext | Space: O(d)

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

type NestedIterator struct {
	stack []*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	stack := make([]*NestedInteger, 0)
	// Push in reverse order
	for i := len(nestedList) - 1; i >= 0; i-- {
		stack = append(stack, nestedList[i])
	}
	return &NestedIterator{stack: stack}
}

func (it *NestedIterator) Next() int {
	top := it.stack[len(it.stack)-1]
	it.stack = it.stack[:len(it.stack)-1]
	return top.GetInteger()
}

func (it *NestedIterator) HasNext() bool {
	for len(it.stack) > 0 {
		top := it.stack[len(it.stack)-1]
		if top.IsInteger() {
			return true
		}
		it.stack = it.stack[:len(it.stack)-1]
		list := top.GetList()
		for i := len(list) - 1; i >= 0; i-- {
			it.stack = append(it.stack, list[i])
		}
	}
	return false
}

func main() {
	// Test case 1: [[1,1],2,[1,1]]
	n1 := NewList(NewInt(1), NewInt(1))
	n2 := NewInt(2)
	n3 := NewList(NewInt(1), NewInt(1))
	it1 := Constructor([]*NestedInteger{n1, n2, n3})
	fmt.Print("Test 1: ")
	for it1.HasNext() {
		fmt.Print(it1.Next(), " ")
	}
	fmt.Println()
	// Expected: 1 1 2 1 1

	// Test case 2: [1,[4,[6]]]
	inner := NewList(NewInt(6))
	mid := NewList(NewInt(4), inner)
	it2 := Constructor([]*NestedInteger{NewInt(1), mid})
	fmt.Print("Test 2: ")
	for it2.HasNext() {
		fmt.Print(it2.Next(), " ")
	}
	fmt.Println()
	// Expected: 1 4 6

	// Test case 3: Empty
	it3 := Constructor([]*NestedInteger{})
	fmt.Print("Test 3: ")
	for it3.HasNext() {
		fmt.Print(it3.Next(), " ")
	}
	fmt.Println()
	// Expected: (nothing)
}
```

## 0343 — Integer Break

```go
package main

// LeetCode #343: Integer Break
// https://leetcode.com/problems/integer-break/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func integerBreak(n int) int {
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}

	product := 1
	for n > 4 {
		product *= 3
		n -= 3
	}
	product *= n
	return product
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", integerBreak(2))
	// Expected: 1

	// Test case 2
	fmt.Println("Test 2:", integerBreak(10))
	// Expected: 36

	// Test case 3
	fmt.Println("Test 3:", integerBreak(5))
	// Expected: 6
}
```

## 0347 — Top K Frequent Elements

```go
package main

// LeetCode #347: Top K Frequent Elements
// https://leetcode.com/problems/top-k-frequent-elements/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func topKFrequent(nums []int, k int) []int {
	// Count frequencies
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	// Bucket sort by frequency (index = frequency)
	buckets := make([][]int, len(nums)+1)
	for num, count := range freq {
		buckets[count] = append(buckets[count], num)
	}

	result := make([]int, 0, k)
	for i := len(buckets) - 1; i >= 0 && len(result) < k; i-- {
		for _, num := range buckets[i] {
			result = append(result, num)
			if len(result) == k {
				break
			}
		}
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
	// Expected: [1, 2]

	// Test case 2
	fmt.Println("Test 2:", topKFrequent([]int{1}, 1))
	// Expected: [1]

	// Test case 3
	fmt.Println("Test 3:", topKFrequent([]int{1, 2, 3, 1, 2, 1}, 2))
	// Expected: [1, 2]
}
```

## 0348 — Design Tic Tac Toe

```go
package main

// LeetCode #348: Design Tic-Tac-Toe
// https://leetcode.com/problems/design-tic-tac-toe/
// Difficulty: Medium [Paid]
// Time: O(1) per move | Space: O(n)

import "fmt"

type TicTacToe struct {
	rows     []int
	cols     []int
	diag     int
	antiDiag int
	n        int
}

func Constructor(n int) TicTacToe {
	return TicTacToe{
		rows: make([]int, n),
		cols: make([]int, n),
		diag: 0, antiDiag: 0, n: n,
	}
}

func (t *TicTacToe) Move(row int, col int, player int) int {
	// Player 1: +1, Player 2: -1
	val := 1
	if player == 2 {
		val = -1
	}

	t.rows[row] += val
	t.cols[col] += val

	if row == col {
		t.diag += val
	}
	if row+col == t.n-1 {
		t.antiDiag += val
	}

	// Check win
	if abs(t.rows[row]) == t.n || abs(t.cols[col]) == t.n ||
		abs(t.diag) == t.n || abs(t.antiDiag) == t.n {
		return player
	}
	return 0
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	// Test case 1
	toe := Constructor(3)
	fmt.Println("Move 1:", toe.Move(0, 0, 1)) // Player 1
	fmt.Println("Move 2:", toe.Move(0, 2, 2)) // Player 2
	fmt.Println("Move 3:", toe.Move(2, 2, 1)) // Player 1
	fmt.Println("Move 4:", toe.Move(1, 1, 2)) // Player 2
	fmt.Println("Move 5:", toe.Move(2, 0, 1)) // Player 1
	fmt.Println("Move 6:", toe.Move(1, 0, 2)) // Player 2
	fmt.Println("Move 7:", toe.Move(2, 1, 1)) // Player 1 - wins
	// Expected: 0, 0, 0, 0, 0, 0, 1

	fmt.Println()

	// Test case 2: Player 2 wins
	toe2 := Constructor(3)
	fmt.Println("Move 1:", toe2.Move(0, 0, 1))
	fmt.Println("Move 2:", toe2.Move(1, 0, 2))
	fmt.Println("Move 3:", toe2.Move(0, 2, 1))
	fmt.Println("Move 4:", toe2.Move(1, 1, 2))
	fmt.Println("Move 5:", toe2.Move(0, 1, 1))
	fmt.Println("Move 6:", toe2.Move(1, 2, 2)) // Player 2 wins
	// Expected: 0, 0, 0, 0, 0, 2
}
```

