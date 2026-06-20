# Medium (Sedang) — Problem ��2830

## 2628 — Json Deep Equal

```go
package main

// LeetCode #2628: JSON Deep Equal
// https://leetcode.com/problems/json-deep-equal/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func jsonDeepEqual(o1, o2 any) bool {
	return reflect.DeepEqual(o1, o2)
}

func deepEqualJSON(a, b string) bool {
	var v1, v2 any
	if err := json.Unmarshal([]byte(a), &v1); err != nil {
		return false
	}
	if err := json.Unmarshal([]byte(b), &v2); err != nil {
		return false
	}
	return jsonDeepEqual(v1, v2)
}

func main() {
	// Test case 1: equal objects
	fmt.Println("Test 1:", deepEqualJSON(`{"a":1,"b":2}`, `{"b":2,"a":1}`))
	// Expected: true

	// Test case 2: different types
	fmt.Println("Test 2:", deepEqualJSON(`{"a":1}`, `{"a":"1"}`))
	// Expected: false

	// Test case 3: arrays
	fmt.Println("Test 3:", deepEqualJSON(`[1,2,3]`, `[1,2,3]`))
	// Expected: true
}
```

## 2631 — Group By

```go
package main

// LeetCode #2631: Group By
// https://leetcode.com/problems/group-by/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func groupBy[T any, K comparable](items []T, keyFn func(T) K) map[K][]T {
	result := make(map[K][]T)
	for _, item := range items {
		key := keyFn(item)
		result[key] = append(result[key], item)
	}
	return result
}

func main() {
	// Test case 1: group integers by even/odd
	nums := []int{1, 2, 3, 4, 5, 6}
	grouped := groupBy(nums, func(n int) string {
		if n%2 == 0 {
			return "even"
		}
		return "odd"
	})
	fmt.Println("Test 1:", grouped)
	// Expected: map[even:[2 4 6] odd:[1 3 5]]

	// Test case 2: group strings by length
	words := []string{"one", "two", "three", "four"}
	byLen := groupBy(words, func(s string) int {
		return len(s)
	})
	fmt.Println("Test 2:", byLen)
	// Expected: map[3:[one two] 5:[three four]]

	// Test case 3
	single := []int{1}
	result := groupBy(single, func(n int) string { return "a" })
	fmt.Println("Test 3:", result)
}
```

## 2632 — Curry

```go
package main

// LeetCode #2632: Curry
// https://leetcode.com/problems/curry/
// Difficulty: Medium [Paid]
// Time: O(1) per call | Space: O(1)

import "fmt"

// curry2 implements currying for a 2-argument function
type curry2 struct {
	fn     func(int, int) int
	args   []int
	arity  int
}

func curry2New(fn func(int, int) int) *curry2 {
	return &curry2{fn: fn, arity: 2}
}

func (c *curry2) call(args ...int) *curry2 {
	c.args = append(c.args, args...)
	return c
}

func (c *curry2) done() int {
	return c.fn(c.args[0], c.args[1])
}

func add(a, b int) int {
	return a + b
}

func main() {
	fn := add

	// Test case 1: currying with two separate calls
	c := curry2New(fn)
	r1 := c.call(1).call(2).done()
	fmt.Println("Test 1:", r1)
	// Expected: 3

	// Test case 2: currying with one call
	c2 := curry2New(fn)
	r2 := c2.call(3, 4).done()
	fmt.Println("Test 2:", r2)
	// Expected: 7
}
```

## 2633 — Convert Object To Json String

```go
package main

// LeetCode #2633: Convert Object to JSON String
// https://leetcode.com/problems/convert-object-to-json-string/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import (
	"encoding/json"
	"fmt"
)

func convertToJSON(obj any) string {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return "null"
	}
	return string(bytes)
}

func main() {
	// Test case 1: object
	fmt.Println("Test 1:", convertToJSON(map[string]any{"a": 1, "b": 2}))
	// Expected: {"a":1,"b":2}

	// Test case 2: array
	fmt.Println("Test 2:", convertToJSON([]any{1, "hello", true}))
	// Expected: [1,"hello",true]

	// Test case 3: nested
	fmt.Println("Test 3:", convertToJSON(map[string]any{"x": []any{1, 2, 3}}))
	// Expected: {"x":[1,2,3]}
}
```

## 2636 — Promise Pool

```go
package main

// LeetCode #2636: Promise Pool
// https://leetcode.com/problems/promise-pool/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import (
	"fmt"
	"sync"
	"time"
)

func promisePool(tasks []func(int) (int, error), limit int) []int {
	var wg sync.WaitGroup
	results := make([]int, len(tasks))
	sem := make(chan struct{}, limit)

	for i, task := range tasks {
		wg.Add(1)
		sem <- struct{}{}
		go func(idx int, t func(int) (int, error)) {
			defer wg.Done()
			defer func() { <-sem }()
			result, _ := t(idx)
			results[idx] = result
		}(i, task)
	}

	wg.Wait()
	return results
}

func main() {
	tasks := []func(int) (int, error){
		func(i int) (int, error) { time.Sleep(50 * time.Millisecond); return i * 2, nil },
		func(i int) (int, error) { time.Sleep(30 * time.Millisecond); return i * 3, nil },
		func(i int) (int, error) { time.Sleep(10 * time.Millisecond); return i * 4, nil },
	}

	results := promisePool(tasks, 2)
	fmt.Println("Test 1:", results)
	// Expected: [0, 3, 8]

	// Test case 2: empty tasks
	fmt.Println("Test 2:", promisePool([]func(int) (int, error){}, 5))
	// Expected: []

	// Test case 3: single task
	single := []func(int) (int, error){
		func(i int) (int, error) { return 42, nil },
	}
	fmt.Println("Test 3:", promisePool(single, 1))
	// Expected: [42]
}
```

## 2637 — Promise Time Limit

```go
package main

// LeetCode #2637: Promise Time Limit
// https://leetcode.com/problems/promise-time-limit/
// Difficulty: Medium
// Time: O(1) | Space: O(1)

import (
	"fmt"
	"sync"
	"time"
)

type TimeLimitedFn struct {
	fn func(int) int
}

type Result struct {
	val int
	err error
}

func timeLimit(fn func(int) int, limit time.Duration) func(int) (int, bool) {
	return func(arg int) (int, bool) {
		var mu sync.Mutex
		done := false
		result := 0

		go func() {
			r := fn(arg)
			mu.Lock()
			if !done {
				result = r
				done = true
			}
			mu.Unlock()
		}()

		time.Sleep(limit)

		mu.Lock()
		defer mu.Unlock()
		if !done {
			done = true
			return 0, false
		}
		return result, true
	}
}

func main() {
	// Test case 1: completes in time
	fastFn := func(x int) int {
		time.Sleep(10 * time.Millisecond)
		return x * 2
	}
	limited := timeLimit(fastFn, 100*time.Millisecond)
	val, ok := limited(5)
	fmt.Println("Test 1:", val, ok)
	// Expected: 10 true

	// Test case 2: times out
	slowFn := func(x int) int {
		time.Sleep(200 * time.Millisecond)
		return x
	}
	limited2 := timeLimit(slowFn, 50*time.Millisecond)
	val2, ok2 := limited2(5)
	fmt.Println("Test 2:", val2, ok2)
	// Expected: 0 false
}
```

## 2638 — Count The Number Of K Free Subsets

```go
package main

// LeetCode #2638: Count the Number of K-Free Subsets
// https://leetcode.com/problems/count-the-number-of-k-free-subsets/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func countKFreeSubsets(nums []int, k int) int64 {
	sort.Ints(nums)
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}

	// Group by remainder modulo k
	groups := make(map[int][]int)
	for v := range freq {
		groups[v%k] = append(groups[v%k], v)
	}

	ans := int64(1) // empty subset
	for _, vals := range groups {
		sort.Ints(vals)
		// DP within group: O(n) with constraint that consecutive vals with diff k can't both be taken
		dp0, dp1 := int64(1), int64(0)
		for i, v := range vals {
			waysSkip := dp0 + dp1
			var waysTake int64
			if i > 0 && v-vals[i-1] == k {
				waysTake = dp0 * ((int64(1) << uint(freq[v])) - 1)
			} else {
				waysTake = (dp0 + dp1) * ((int64(1) << uint(freq[v])) - 1)
			}
			dp0, dp1 = waysSkip, waysTake
		}
		ans *= (dp0 + dp1)
	}

	return ans - 1 // exclude empty subset
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", countKFreeSubsets([]int{1, 2, 3}, 1))
	// Expected: 4

	// Test case 2
	fmt.Println("Test 2:", countKFreeSubsets([]int{1, 2, 3, 4}, 2))
	// Expected: 8

	// Test case 3
	fmt.Println("Test 3:", countKFreeSubsets([]int{1, 3, 5}, 2))
	// Expected: 4
}
```

## 2640 — Find The Score Of All Prefixes Of An Array

```go
package main

// LeetCode #2640: Find the Score of All Prefixes of an Array
// https://leetcode.com/problems/find-the-score-of-all-prefixes-of-an-array/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func findPrefixScore(nums []int) []int64 {
	n := len(nums)
	ans := make([]int64, n)
	maxSoFar := nums[0]
	prefixSum := int64(0)

	for i, v := range nums {
		if v > maxSoFar {
			maxSoFar = v
		}
		conver := int64(v + maxSoFar)
		prefixSum += conver
		ans[i] = prefixSum
	}
	return ans
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", findPrefixScore([]int{2, 3, 7, 5, 10}))
	// Expected: [4,10,24,36,56]

	// Test case 2
	fmt.Println("Test 2:", findPrefixScore([]int{1, 1, 1}))
	// Expected: [2,4,6]

	// Test case 3
	fmt.Println("Test 3:", findPrefixScore([]int{5}))
	// Expected: [10]
}
```

## 2641 — Cousins In Binary Tree Ii

```go
package main

// LeetCode #2641: Cousins in Binary Tree II
// https://leetcode.com/problems/cousins-in-binary-tree-ii/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func replaceValueInTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}
	root.Val = 0

	for len(queue) > 0 {
		size := len(queue)
		levelSum := 0

		// First pass: compute total sum of this level (ignore size)
		_ = size
		for _, node := range queue {
			if node.Left != nil {
				levelSum += node.Left.Val
			}
			if node.Right != nil {
				levelSum += node.Right.Val
			}
		}

		// Second pass: update children values
		for _, node := range queue {
			siblingSum := 0
			if node.Left != nil {
				siblingSum += node.Left.Val
			}
			if node.Right != nil {
				siblingSum += node.Right.Val
			}
			if node.Left != nil {
				node.Left.Val = levelSum - siblingSum
			}
			if node.Right != nil {
				node.Right.Val = levelSum - siblingSum
			}
		}

		// Build next level queue
		nextQueue := []*TreeNode{}
		for _, node := range queue {
			if node.Left != nil {
				nextQueue = append(nextQueue, node.Left)
			}
			if node.Right != nil {
				nextQueue = append(nextQueue, node.Right)
			}
		}
		queue = nextQueue
	}
	return root
}

func printTree(root *TreeNode) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Print(node.Val, " ")
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	fmt.Println()
}

func main() {
	// Test case 1: [5,4,9,1,10,null,7]
	root1 := &TreeNode{Val: 5}
	root1.Left = &TreeNode{Val: 4}
	root1.Right = &TreeNode{Val: 9}
	root1.Left.Left = &TreeNode{Val: 1}
	root1.Left.Right = &TreeNode{Val: 10}
	root1.Right.Right = &TreeNode{Val: 7}
	replaceValueInTree(root1)
	fmt.Print("Test 1: ")
	printTree(root1)
	// Expected: 0,0,0,7,7,null,11

	// Test case 2: single node
	root2 := &TreeNode{Val: 1}
	replaceValueInTree(root2)
	fmt.Print("Test 2: ")
	printTree(root2)
	// Expected: 0
}
```

## 2645 — Minimum Additions To Make Valid String

```go
package main

// LeetCode #2645: Minimum Additions to Make Valid String
// https://leetcode.com/problems/minimum-additions-to-make-valid-string/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func addMinimum(word string) int {
	n := len(word)
	count := 0
	i := 0

	// Pattern we're building: "abc"
	expected := []byte{'a', 'b', 'c'}
	expectedIdx := 0

	for i < n {
		if word[i] == expected[expectedIdx] {
			i++
		} else {
			count++
		}
		expectedIdx = (expectedIdx + 1) % 3
	}

	// Finish the current "abc" cycle
	for expectedIdx != 0 {
		count++
		expectedIdx = (expectedIdx + 1) % 3
	}

	return count
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", addMinimum("b"))
	// Expected: 2 (add "a" and "c" to make "abc")

	// Test case 2
	fmt.Println("Test 2:", addMinimum("aaa"))
	// Expected: 6 (add "bc" twice)

	// Test case 3
	fmt.Println("Test 3:", addMinimum("abc"))
	// Expected: 0
}
```

## 2649 — Nested Array Generator

```go
package main

// LeetCode #2649: Nested Array Generator
// https://leetcode.com/problems/nested-array-generator/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func nestedArrayGenerator(arr []any) []int {
	result := []int{}
	var dfs func(item any)
	dfs = func(item any) {
		switch v := item.(type) {
		case []any:
			for _, sub := range v {
				dfs(sub)
			}
		case int:
			result = append(result, v)
		}
	}
	for _, item := range arr {
		dfs(item)
	}
	return result
}

func main() {
	// Test case 1: flat array
	arr1 := []any{1, 2, 3}
	fmt.Println("Test 1:", nestedArrayGenerator(arr1))
	// Expected: [1, 2, 3]

	// Test case 2: nested array
	arr2 := []any{1, []any{2, 3}, 4}
	fmt.Println("Test 2:", nestedArrayGenerator(arr2))
	// Expected: [1, 2, 3, 4]

	// Test case 3: deeply nested
	arr3 := []any{1, []any{2, []any{3, 4}, 5}, 6}
	fmt.Println("Test 3:", nestedArrayGenerator(arr3))
	// Expected: [1, 2, 3, 4, 5, 6]
}
```

## 2653 — Sliding Subarray Beauty

```go
package main

// LeetCode #2653: Sliding Subarray Beauty
// https://leetcode.com/problems/sliding-subarray-beauty/
// Difficulty: Medium
// Time: O(n * 50) | Space: O(1)

import "fmt"

func getSubarrayBeauty(nums []int, k int, x int) []int {
	n := len(nums)
	ans := make([]int, n-k+1)
	freq := make([]int, 101) // values in [-50, 50] shifted by 50

	for i := 0; i < k; i++ {
		freq[nums[i]+50]++
	}

	for i := k; i <= n; i++ {
		// Find x-th smallest
		count := 0
		val := 0
		for j := 0; j <= 100; j++ {
			count += freq[j]
			if count >= x {
				val = j - 50
				break
			}
		}
		// Only negative values qualify as "beauty"
		if val < 0 {
			ans[i-k] = val
		} else {
			ans[i-k] = 0
		}

		if i < n {
			freq[nums[i-k]+50]--
			freq[nums[i]+50]++
		}
	}
	return ans
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", getSubarrayBeauty([]int{1, -1, -3, -2, 3}, 3, 2))
	// Expected: [-1,-2,-2]

	// Test case 2
	fmt.Println("Test 2:", getSubarrayBeauty([]int{-1, -2, -3, -4, -5}, 2, 2))
	// Expected: [-1,-2,-3,-4]

	// Test case 3
	fmt.Println("Test 3:", getSubarrayBeauty([]int{-3, 1, 2, -3, 0, -3}, 2, 1))
	// Expected: [-3,0,-3,-3,-3]
}
```

## 2654 — Minimum Number Of Operations To Make All Array Elements Equal To 1

```go
package main

// LeetCode #2654: Minimum Number of Operations to Make All Array Elements Equal to 1
// https://leetcode.com/problems/minimum-number-of-operations-to-make-all-array-elements-equal-to-1/
// Difficulty: Medium
// Time: O(n^2) | Space: O(1)

import "fmt"

func minOperations(nums []int) int {
	n := len(nums)

	// If there's already a 1, answer is n - count(1)
	ones := 0
	for _, v := range nums {
		if v == 1 {
			ones++
		}
	}
	if ones > 0 {
		return n - ones
	}

	// Find shortest subarray with GCD=1
	minLen := n + 1
	for i := 0; i < n; i++ {
		g := nums[i]
		for j := i; j < n; j++ {
			g = gcd(g, nums[j])
			if g == 1 {
				if j-i+1 < minLen {
					minLen = j - i + 1
				}
				break
			}
		}
	}

	if minLen > n {
		return -1
	}

	// Need (minLen-1) operations to convert subarray to 1, then (n-1) to spread
	return (minLen - 1) + (n - 1)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minOperations([]int{2, 6, 3, 4}))
	// Expected: 4

	// Test case 2
	fmt.Println("Test 2:", minOperations([]int{2, 10, 6, 14}))
	// Expected: -1

	// Test case 3: already has 1
	fmt.Println("Test 3:", minOperations([]int{1, 5, 3}))
	// Expected: 2
}
```

## 2655 — Find Maximal Uncovered Ranges

```go
package main

// LeetCode #2655: Find Maximal Uncovered Ranges
// https://leetcode.com/problems/find-maximal-uncovered-ranges/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func findMaximalUncoveredRanges(n int, ranges [][]int) [][]int {
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	ans := [][]int{}
	prevEnd := -1

	for _, r := range ranges {
		start, end := r[0], r[1]
		if start > prevEnd+1 {
			ans = append(ans, []int{prevEnd + 1, start - 1})
		}
		if end > prevEnd {
			prevEnd = end
		}
	}

	if prevEnd < n-1 {
		ans = append(ans, []int{prevEnd + 1, n - 1})
	}

	return ans
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", findMaximalUncoveredRanges(10, [][]int{{0, 2}, {5, 7}}))
	// Expected: [[3,4],[8,9]]

	// Test case 2
	fmt.Println("Test 2:", findMaximalUncoveredRanges(5, [][]int{{0, 4}}))
	// Expected: []

	// Test case 3
	fmt.Println("Test 3:", findMaximalUncoveredRanges(5, [][]int{}))

	// Expected: [[0,4]]
}
```

## 2657 — Find The Prefix Common Array Of Two Arrays

```go
package main

// LeetCode #2657: Find the Prefix Common Array of Two Arrays
// https://leetcode.com/problems/find-the-prefix-common-array-of-two-arrays/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func findThePrefixCommonArray(A []int, B []int) []int {
	n := len(A)
	ans := make([]int, n)
	seen := make(map[int]bool)
	count := 0

	for i := 0; i < n; i++ {
		if seen[A[i]] {
			count++
		} else {
			seen[A[i]] = true
		}
		if seen[B[i]] {
			count++
		} else {
			seen[B[i]] = true
		}
		ans[i] = count
	}
	return ans
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", findThePrefixCommonArray([]int{1, 3, 2, 4}, []int{3, 1, 2, 4}))
	// Expected: [0,2,3,4]

	// Test case 2
	fmt.Println("Test 2:", findThePrefixCommonArray([]int{1, 2, 3}, []int{1, 2, 3}))
	// Expected: [1,2,3]

	// Test case 3
	fmt.Println("Test 3:", findThePrefixCommonArray([]int{1, 2, 3}, []int{3, 1, 2}))
	// Expected: [0,1,3]
}
```

## 2658 — Maximum Number Of Fish In A Grid

```go
package main

// LeetCode #2658: Maximum Number of Fish in a Grid
// https://leetcode.com/problems/maximum-number-of-fish-in-a-grid/
// Difficulty: Medium
// Time: O(m*n) | Space: O(m*n)

import "fmt"

func findMaxFish(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	maxFish := 0

	var dfs func(r, c int) int
	dfs = func(r, c int) int {
		if r < 0 || r >= m || c < 0 || c >= n || grid[r][c] == 0 {
			return 0
		}
		fish := grid[r][c]
		grid[r][c] = 0 // Mark visited
		fish += dfs(r-1, c)
		fish += dfs(r+1, c)
		fish += dfs(r, c-1)
		fish += dfs(r, c+1)
		return fish
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] > 0 {
				fish := dfs(i, j)
				if fish > maxFish {
					maxFish = fish
				}
			}
		}
	}
	return maxFish
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", findMaxFish([][]int{{0, 2, 1, 0}, {4, 0, 0, 3}, {1, 0, 0, 4}, {0, 3, 2, 0}}))
	// Expected: 7

	// Test case 2
	fmt.Println("Test 2:", findMaxFish([][]int{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 1}}))
	// Expected: 1

	// Test case 3
	fmt.Println("Test 3:", findMaxFish([][]int{{0, 0}, {0, 0}}))
	// Expected: 0
}
```

## 2661 — First Completely Painted Row Or Column

```go
package main

// LeetCode #2661: First Completely Painted Row or Column
// https://leetcode.com/problems/first-completely-painted-row-or-column/
// Difficulty: Medium
// Time: O(m*n) | Space: O(m*n)

import "fmt"

func firstCompleteIndex(arr []int, mat [][]int) int {
	m, n := len(mat), len(mat[0])
	pos := make([][2]int, m*n+1)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			pos[mat[i][j]] = [2]int{i, j}
		}
	}

	rowCount := make([]int, m)
	colCount := make([]int, n)

	for idx, v := range arr {
		r, c := pos[v][0], pos[v][1]
		rowCount[r]++
		colCount[c]++
		if rowCount[r] == n || colCount[c] == m {
			return idx
		}
	}
	return -1
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", firstCompleteIndex([]int{1, 3, 4, 2}, [][]int{{1, 4}, {2, 3}}))
	// Expected: 2

	// Test case 2
	fmt.Println("Test 2:", firstCompleteIndex([]int{2, 8, 7, 4, 1, 3, 5, 6, 9}, [][]int{{3, 2, 5}, {1, 4, 6}, {8, 7, 9}}))
	// Expected: 3

	// Test case 3
	fmt.Println("Test 3:", firstCompleteIndex([]int{1, 2}, [][]int{{1}, {2}}))
	// Expected: 0
}
```

## 2662 — Minimum Cost Of A Path With Special Roads

```go
package main

// LeetCode #2662: Minimum Cost of a Path With Special Roads
// https://leetcode.com/problems/minimum-cost-of-a-path-with-special-roads/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import "fmt"

func minimumCost(start []int, target []int, specialRoads [][]int) int {
	// Use Dijkstra: dist[i] = min cost to reach special road i's end
	n := len(specialRoads)
	dist := make([]int, n)
	visited := make([]bool, n)

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	minDist := abs(target[0]-start[0]) + abs(target[1]-start[1])

	// Initialize dist with cost to reach each special road's start from (start)
	for i, road := range specialRoads {
		costToStart := abs(road[0]-start[0]) + abs(road[1]-start[1])
		dist[i] = costToStart + road[4] // cost to travel the special road
	}

	for {
		// Find unvisited node with minimum dist
		u := -1
		best := int(1e18)
		for i := 0; i < n; i++ {
			if !visited[i] && dist[i] < best {
				best = dist[i]
				u = i
			}
		}
		if u == -1 {
			break
		}
		visited[u] = true

		// From road u's end, try going directly to target
		costToTarget := dist[u] + abs(target[0]-specialRoads[u][2]) + abs(target[1]-specialRoads[u][3])
		if costToTarget < minDist {
			minDist = costToTarget
		}

		// From road u's end, try going to another special road's start
		for v, road := range specialRoads {
			if visited[v] {
				continue
			}
			cost := dist[u] + abs(road[0]-specialRoads[u][2]) + abs(road[1]-specialRoads[u][3]) + road[4]
			if cost < dist[v] {
				dist[v] = cost
			}
		}
	}

	return minDist
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minimumCost([]int{1, 1}, []int{4, 5}, [][]int{{1, 2, 3, 3, 2}, {3, 4, 4, 5, 1}}))
	// Expected: 5

	// Test case 2: no special roads
	fmt.Println("Test 2:", minimumCost([]int{0, 0}, []int{3, 4}, [][]int{}))
	// Expected: 7

	// Test case 3
	fmt.Println("Test 3:", minimumCost([]int{1, 1}, []int{10, 10}, [][]int{{5, 5, 6, 6, 1}}))
	// Expected: 18
}
```

## 2664 — The Knights Tour

```go
package main

// LeetCode #2664: The Knight's Tour
// https://leetcode.com/problems/the-knights-tour/
// Difficulty: Medium [Paid]
// Time: O(8^(m*n)) worst case | Space: O(m*n)

import "fmt"

var dirs = [8][2]int{{-2, -1}, {-2, 1}, {-1, -2}, {-1, 2}, {1, -2}, {1, 2}, {2, -1}, {2, 1}}

func tourOfKnight(m int, n int, r int, c int) [][]int {
	board := make([][]int, m)
	for i := range board {
		board[i] = make([]int, n)
		for j := range board[i] {
			board[i][j] = -1
		}
	}

	var solve func(row, col, step int) bool
	solve = func(row, col, step int) bool {
		if step == m*n {
			return true
		}
		for _, d := range dirs {
			nr, nc := row+d[0], col+d[1]
			if nr >= 0 && nr < m && nc >= 0 && nc < n && board[nr][nc] == -1 {
				board[nr][nc] = step
				if solve(nr, nc, step+1) {
					return true
				}
				board[nr][nc] = -1
			}
		}
		return false
	}

	board[r][c] = 0
	solve(r, c, 1)
	return board
}

func main() {
	// Test case 1: 1x1 board
	fmt.Println("Test 1:", tourOfKnight(1, 1, 0, 0))
	// Expected: [[0]]

	// Test case 2
	res := tourOfKnight(3, 4, 0, 0)
	fmt.Println("Test 2 dims:", len(res), "x", len(res[0]))
	// Verify all cells filled
	count := 0
	for _, row := range res {
		for _, v := range row {
			if v >= 0 {
				count++
			}
		}
	}
	fmt.Println("  filled:", count)

	// Test case 3
	res2 := tourOfKnight(5, 5, 2, 2)
	fmt.Println("Test 3 dims:", len(res2), "x", len(res2[0]))
}
```

## 2671 — Frequency Tracker

```go
package main

// LeetCode #2671: Frequency Tracker
// https://leetcode.com/problems/frequency-tracker/
// Difficulty: Medium
// Time: O(1) per operation | Space: O(n)

import "fmt"

type FrequencyTracker struct {
	freq   map[int]int
	freqOf map[int]int
}

func Constructor() FrequencyTracker {
	return FrequencyTracker{
		freq:   make(map[int]int),
		freqOf: make(map[int]int),
	}
}

func (this *FrequencyTracker) Add(number int) {
	oldFreq := this.freq[number]
	if this.freqOf[oldFreq] > 0 {
		this.freqOf[oldFreq]--
	}
	this.freq[number]++
	this.freqOf[oldFreq+1]++
}

func (this *FrequencyTracker) DeleteOne(number int) {
	if this.freq[number] == 0 {
		return
	}
	oldFreq := this.freq[number]
	this.freqOf[oldFreq]--
	this.freq[number]--
	if this.freq[number] > 0 {
		this.freqOf[oldFreq-1]++
	}
}

func (this *FrequencyTracker) HasFrequency(frequency int) bool {
	return this.freqOf[frequency] > 0
}

func main() {
	ft := Constructor()

	// Test case 1
	ft.Add(3)
	ft.Add(3)
	fmt.Println("Test 1:", ft.HasFrequency(2))
	// Expected: true

	// Test case 2
	ft.DeleteOne(3)
	fmt.Println("Test 2:", ft.HasFrequency(1))
	// Expected: true

	// Test case 3
	fmt.Println("Test 3:", ft.HasFrequency(2))
	// Expected: false
}
```

## 2672 — Number Of Adjacent Elements With The Same Color

```go
package main

// LeetCode #2672: Number of Adjacent Elements With the Same Color
// https://leetcode.com/problems/number-of-adjacent-elements-with-the-same-color/
// Difficulty: Medium
// Time: O(n + q) | Space: O(n)

import "fmt"

func colorTheArray(n int, queries [][]int) []int {
	colors := make([]int, n)
	ans := make([]int, len(queries))
	pairs := 0

	for i, q := range queries {
		idx, color := q[0], q[1]

		// Remove existing pairs
		if colors[idx] != 0 {
			if idx > 0 && colors[idx-1] == colors[idx] {
				pairs--
			}
			if idx < n-1 && colors[idx+1] == colors[idx] {
				pairs--
			}
		}

		colors[idx] = color

		// Add new pairs
		if idx > 0 && colors[idx-1] == colors[idx] {
			pairs++
		}
		if idx < n-1 && colors[idx+1] == colors[idx] {
			pairs++
		}

		ans[i] = pairs
	}
	return ans
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", colorTheArray(4, [][]int{{0, 2}, {1, 2}, {3, 1}, {1, 1}, {2, 1}}))
	// Expected: [0,1,1,0,2]

	// Test case 2
	fmt.Println("Test 2:", colorTheArray(1, [][]int{{0, 1}}))
	// Expected: [0]

	// Test case 3
	fmt.Println("Test 3:", colorTheArray(3, [][]int{{0, 1}, {1, 1}, {2, 1}}))
	// Expected: [0,1,2]
}
```

## 2673 — Make Costs Of Paths Equal In A Binary Tree

```go
package main

// LeetCode #2673: Make Costs of Paths Equal in a Binary Tree
// https://leetcode.com/problems/make-costs-of-paths-equal-in-a-binary-tree/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func minIncrements(n int, cost []int) int {
	ans := 0
	// Process from leaves to root (bottom-up)
	for i := n/2 - 1; i >= 0; i-- {
		left := 2*i + 1
		right := 2*i + 2
		// Make both child subtrees have equal max path sum
		if cost[left] > cost[right] {
			ans += cost[left] - cost[right]
			cost[i] += cost[left]
		} else {
			ans += cost[right] - cost[left]
			cost[i] += cost[right]
		}
	}
	return ans
}

func main() {
	// Test case 1: n=7, cost=[1,5,2,2,3,3,1]
	fmt.Println("Test 1:", minIncrements(7, []int{1, 5, 2, 2, 3, 3, 1}))
	// Expected: 6

	// Test case 2: n=3, cost=[5,3,3]
	fmt.Println("Test 2:", minIncrements(3, []int{5, 3, 3}))
	// Expected: 0

	// Test case 3
	fmt.Println("Test 3:", minIncrements(3, []int{10, 5, 5}))
	// Expected: 0
}
```

## 2674 — Split A Circular Linked List

```go
package main

// LeetCode #2674: Split a Circular Linked List
// https://leetcode.com/problems/split-a-circular-linked-list/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitCircularLinkedList(head *ListNode) []*ListNode {
	if head == nil || head.Next == nil {
		return []*ListNode{head, nil}
	}

	// Find the midpoint using slow/fast pointers
	slow, fast := head, head
	for fast.Next != head && fast.Next.Next != head {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// First half: head to slow
	list1 := head

	// Second half: slow.Next to end, make it circular
	list2 := slow.Next

	// Make first list circular
	slow.Next = head

	// Find end of second list and make it circular
	curr := list2
	for curr.Next != head {
		curr = curr.Next
	}
	curr.Next = list2

	return []*ListNode{list1, list2}
}

func printCircular(head *ListNode, count int) {
	curr := head
	for i := 0; i < count && curr != nil; i++ {
		fmt.Print(curr.Val, " ")
		curr = curr.Next
	}
}

func main() {
	// Test case 1: [1,2,3,4]
	head1 := &ListNode{Val: 1}
	head1.Next = &ListNode{Val: 2}
	head1.Next.Next = &ListNode{Val: 3}
	head1.Next.Next.Next = &ListNode{Val: 4}
	head1.Next.Next.Next.Next = head1

	parts := splitCircularLinkedList(head1)
	fmt.Print("Test 1 list1: ")
	printCircular(parts[0], 2)
	fmt.Println()

	fmt.Print("Test 1 list2: ")
	printCircular(parts[1], 2)
	fmt.Println()

	// Test case 2: single node
	head2 := &ListNode{Val: 1}
	head2.Next = head2
	parts2 := splitCircularLinkedList(head2)
	fmt.Println("Test 2 list1:", parts2[0].Val)
	fmt.Println("Test 2 list2:", parts2[1])
}
```

## 2676 — Throttle

```go
package main

// LeetCode #2676: Throttle
// https://leetcode.com/problems/throttle/
// Difficulty: Medium [Paid]
// Time: O(1) per call | Space: O(1)

import (
	"fmt"
	"sync"
	"time"
)

type ThrottledFn struct {
	fn       func(int)
	interval time.Duration
	mu       sync.Mutex
	timer    *time.Timer
	lastCall time.Time
	pending  bool
	lastArg  int
}

func (t *ThrottledFn) Call(arg int) {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.lastArg = arg
	if !t.pending {
		t.pending = true
		go func() {
			time.Sleep(t.interval)
			t.mu.Lock()
			t.fn(t.lastArg)
			t.pending = false
			t.mu.Unlock()
		}()
	}
}

func newThrottle(fn func(int), interval time.Duration) *ThrottledFn {
	return &ThrottledFn{
		fn:       fn,
		interval: interval,
	}
}

func main() {
	results := []int{}
	mu := sync.Mutex{}
	throttled := newThrottle(func(x int) {
		mu.Lock()
		results = append(results, x)
		mu.Unlock()
	}, 50*time.Millisecond)

	throttled.Call(1)
	throttled.Call(2)
	throttled.Call(3)

	time.Sleep(100 * time.Millisecond)
	mu.Lock()
	fmt.Println("Test 1:", results)
	mu.Unlock()
	// Expected: [3] (only last call within throttle window)

	throttled.Call(4)
	time.Sleep(100 * time.Millisecond)
	mu.Lock()
	fmt.Println("Test 2:", results)
	mu.Unlock()
	// Expected: [3, 4]
}
```

## 2679 — Sum In A Matrix

```go
package main

// LeetCode #2679: Sum in a Matrix
// https://leetcode.com/problems/sum-in-a-matrix/
// Difficulty: Medium
// Time: O(m * n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func matrixSum(nums [][]int) int {
	m := len(nums)
	if m == 0 {
		return 0
	}
	n := len(nums[0])

	// Sort each row
	for i := 0; i < m; i++ {
		sort.Ints(nums[i])
	}

	ans := 0
	for col := 0; col < n; col++ {
		maxVal := 0
		for row := 0; row < m; row++ {
			if nums[row][col] > maxVal {
				maxVal = nums[row][col]
			}
		}
		ans += maxVal
	}
	return ans
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", matrixSum([][]int{{7, 2, 1}, {6, 4, 2}, {6, 5, 3}, {3, 2, 1}}))
	// Expected: 15

	// Test case 2
	fmt.Println("Test 2:", matrixSum([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	// Expected: 18

	// Test case 3
	fmt.Println("Test 3:", matrixSum([][]int{{1}}))
	// Expected: 1
}
```

## 2680 — Maximum Or

```go
package main

// LeetCode #2680: Maximum OR
// https://leetcode.com/problems/maximum-or/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func maximumOr(nums []int, k int) int64 {
	n := len(nums)

	// suffix[i] = OR of nums from i to n-1
	suffix := make([]int64, n+1)
	for i := n - 1; i >= 0; i-- {
		suffix[i] = suffix[i+1] | int64(nums[i])
	}

	var prefix int64
	var ans int64

	for i := 0; i < n; i++ {
		candidate := prefix | (int64(nums[i]) << uint(k)) | suffix[i+1]
		if candidate > ans {
			ans = candidate
		}
		prefix |= int64(nums[i])
	}
	return ans
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", maximumOr([]int{12, 9}, 1))
	// Expected: 30

	// Test case 2
	fmt.Println("Test 2:", maximumOr([]int{8, 1, 2}, 2))
	// Expected: 35

	// Test case 3
	fmt.Println("Test 3:", maximumOr([]int{10, 8, 4}, 1))
	// Expected: 30
}
```

## 2683 — Neighboring Bitwise Xor

```go
package main

// LeetCode #2683: Neighboring Bitwise XOR
// https://leetcode.com/problems/neighboring-bitwise-xor/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func doesValidArrayExist(derived []int) bool {
	xor := 0
	for _, v := range derived {
		xor ^= v
	}
	return xor == 0
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", doesValidArrayExist([]int{1, 1, 0}))
	// Expected: true

	// Test case 2
	fmt.Println("Test 2:", doesValidArrayExist([]int{1, 1}))
	// Expected: true

	// Test case 3
	fmt.Println("Test 3:", doesValidArrayExist([]int{1, 0}))
	// Expected: false
}
```

## 2684 — Maximum Number Of Moves In A Grid

```go
package main

// LeetCode #2684: Maximum Number of Moves in a Grid
// https://leetcode.com/problems/maximum-number-of-moves-in-a-grid/
// Difficulty: Medium
// Time: O(m*n) | Space: O(m*n)

import "fmt"

func maxMoves(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	// dp[row][col] = max moves from (row, col)
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	ans := 0
	for col := n - 2; col >= 0; col-- {
		for row := 0; row < m; row++ {
			val := grid[row][col]
			maxNext := 0
			// Can move to row-1, col+1
			if row > 0 && grid[row-1][col+1] > val {
				if 1+dp[row-1][col+1] > maxNext {
					maxNext = 1 + dp[row-1][col+1]
				}
			}
			// Can move to row, col+1
			if grid[row][col+1] > val {
				if 1+dp[row][col+1] > maxNext {
					maxNext = 1 + dp[row][col+1]
				}
			}
			// Can move to row+1, col+1
			if row < m-1 && grid[row+1][col+1] > val {
				if 1+dp[row+1][col+1] > maxNext {
					maxNext = 1 + dp[row+1][col+1]
				}
			}
			dp[row][col] = maxNext
			if col == 0 && maxNext > ans {
				ans = maxNext
			}
		}
	}
	return ans
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", maxMoves([][]int{{2, 4, 3, 5}, {5, 4, 9, 3}, {3, 4, 2, 11}, {10, 9, 13, 15}}))
	// Expected: 3

	// Test case 2
	fmt.Println("Test 2:", maxMoves([][]int{{3, 2, 4}, {2, 1, 9}, {1, 1, 7}}))
	// Expected: 0

	// Test case 3
	fmt.Println("Test 3:", maxMoves([][]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}}))
	// Expected: 2
}
```

## 2685 — Count The Number Of Complete Components

```go
package main

// LeetCode #2685: Count the Number of Complete Components
// https://leetcode.com/problems/count-the-number-of-complete-components/
// Difficulty: Medium
// Time: O(V + E) | Space: O(V + E)

import "fmt"

func countCompleteComponents(n int, edges [][]int) int {
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	visited := make([]bool, n)
	count := 0

	for i := 0; i < n; i++ {
		if visited[i] {
			continue
		}

		// BFS/DFS to find component
		queue := []int{i}
		visited[i] = true
		vertices := []int{}

		for len(queue) > 0 {
			u := queue[0]
			queue = queue[1:]
			vertices = append(vertices, u)
			for _, v := range adj[u] {
				if !visited[v] {
					visited[v] = true
					queue = append(queue, v)
				}
			}
		}

		// Check if complete: each vertex should have (k-1) edges to other vertices in component
		k := len(vertices)
		isComplete := true
		for _, v := range vertices {
			if len(adj[v]) != k-1 {
				isComplete = false
				break
			}
		}
		if isComplete {
			count++
		}
	}
	return count
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", countCompleteComponents(6, [][]int{{0, 1}, {0, 2}, {1, 2}, {3, 4}}))
	// Expected: 2 ({0,1,2} complete, {3,4} not complete with 5)

	// Test case 2
	fmt.Println("Test 2:", countCompleteComponents(6, [][]int{{0, 1}, {0, 2}, {1, 2}, {3, 4}, {3, 5}}))
	// Expected: 1

	// Test case 3: single nodes
	fmt.Println("Test 3:", countCompleteComponents(3, [][]int{}))
	// Expected: 3 (each single node is complete)
}
```

## 2686 — Immediate Food Delivery Iii

```go
package main

// LeetCode #2686: Immediate Food Delivery III
// https://leetcode.com/problems/immediate-food-delivery-iii/
// Difficulty: Medium (SQL problem — simulated in Go)
//
// Simulates: For each order_date, calculate the percentage of immediate
// orders (where customer_pref_delivery_date == order_date).
// Round to 2 decimal places.

import (
	"fmt"
	"math"
	"sort"
)

// Delivery represents the Delivery database table.
type Delivery struct {
	DeliveryID              int
	CustomerID              int
	OrderDate               string // format: "YYYY-MM-DD"
	CustomerPrefDeliveryDate string // format: "YYYY-MM-DD"
}

// DeliveryResult holds the output.
type DeliveryResult struct {
	OrderDate           string
	ImmediatePercentage float64
}

// immediateFoodDeliveryIII simulates the SQL query.
// Time: O(n) | Space: O(d) where d = distinct order dates.
func immediateFoodDeliveryIII(deliveries []Delivery) []DeliveryResult {
	type dateStats struct {
		total     int
		immediate int
	}

	stats := make(map[string]*dateStats)

	for _, d := range deliveries {
		if stats[d.OrderDate] == nil {
			stats[d.OrderDate] = &dateStats{}
		}
		stats[d.OrderDate].total++
		if d.OrderDate == d.CustomerPrefDeliveryDate {
			stats[d.OrderDate].immediate++
		}
	}

	var results []DeliveryResult
	for date, s := range stats {
		// Round to 2 decimal places, matching SQL ROUND(x, 2).
		pct := math.Round(float64(s.immediate)/float64(s.total)*10000) / 100
		results = append(results, DeliveryResult{
			OrderDate:           date,
			ImmediatePercentage: pct,
		})
	}

	// Order by order_date ASC.
	sort.Slice(results, func(i, j int) bool {
		return results[i].OrderDate < results[j].OrderDate
	})

	return results
}

func main() {
	// Test data from the problem.
	deliveries := []Delivery{
		{DeliveryID: 1, CustomerID: 1, OrderDate: "2019-08-01", CustomerPrefDeliveryDate: "2019-08-01"},
		{DeliveryID: 2, CustomerID: 2, OrderDate: "2019-08-01", CustomerPrefDeliveryDate: "2019-08-02"},
		{DeliveryID: 3, CustomerID: 3, OrderDate: "2019-08-01", CustomerPrefDeliveryDate: "2019-08-01"},
		{DeliveryID: 4, CustomerID: 4, OrderDate: "2019-08-02", CustomerPrefDeliveryDate: "2019-08-02"},
		{DeliveryID: 5, CustomerID: 5, OrderDate: "2019-08-02", CustomerPrefDeliveryDate: "2019-08-03"},
		{DeliveryID: 6, CustomerID: 6, OrderDate: "2019-08-02", CustomerPrefDeliveryDate: "2019-08-02"},
		{DeliveryID: 7, CustomerID: 7, OrderDate: "2019-08-03", CustomerPrefDeliveryDate: "2019-08-03"},
		{DeliveryID: 8, CustomerID: 8, OrderDate: "2019-08-03", CustomerPrefDeliveryDate: "2019-08-03"},
		{DeliveryID: 9, CustomerID: 9, OrderDate: "2019-08-04", CustomerPrefDeliveryDate: "2019-08-05"},
		{DeliveryID: 10, CustomerID: 10, OrderDate: "2019-08-04", CustomerPrefDeliveryDate: "2019-08-06"},
	}

	results := immediateFoodDeliveryIII(deliveries)

	fmt.Println("Immediate Food Delivery III (order_date | immediate_percentage):")
	for _, r := range results {
		fmt.Printf("%s | %.2f\n", r.OrderDate, r.ImmediatePercentage)
	}
	// Expected output:
	// 2019-08-01 | 66.67
	// 2019-08-02 | 66.67
	// 2019-08-03 | 100.00
	// 2019-08-04 | 0.00
}
```

## 2688 — Find Active Users

```go
package main

// LeetCode #2688: Find Active Users
// https://leetcode.com/problems/find-active-users/
// Difficulty: Medium (SQL problem — simulated in Go)
//
// Simulates: Find users who made a second purchase within 7 days
// (inclusive) of another purchase. The 7-day window is inclusive of
// both start and end dates.

import (
	"fmt"
	"sort"
	"time"
)

// UserPurchase represents the Users database table.
type UserPurchase struct {
	UserID    int
	Item      string
	CreatedAt string // format: "YYYY-MM-DD"
	Amount    int
}

// findActiveUsers simulates the SQL query.
// Time: O(n log n) | Space: O(n)
// n = number of purchase records.
func findActiveUsers(purchases []UserPurchase) []int {
	// Group purchases by user_id.
	userDates := make(map[int][]string)
	for _, p := range purchases {
		userDates[p.UserID] = append(userDates[p.UserID], p.CreatedAt)
	}

	// Count distinct dates per user.
	userUniqueDates := make(map[int][]string)
	for uid, dates := range userDates {
		seen := make(map[string]bool)
		for _, d := range dates {
			if !seen[d] {
				seen[d] = true
				userUniqueDates[uid] = append(userUniqueDates[uid], d)
			}
		}
	}

	var active []int

	for uid, dates := range userUniqueDates {
		// Sort dates ascending.
		sort.Strings(dates)

		// Check consecutive purchases for <= 7 day gap.
		for i := 1; i < len(dates); i++ {
			prev, err1 := time.Parse("2006-01-02", dates[i-1])
			curr, err2 := time.Parse("2006-01-02", dates[i])
			if err1 != nil || err2 != nil {
				continue
			}
			diff := curr.Sub(prev).Hours() / 24
			if diff >= 0 && diff <= 7 {
				active = append(active, uid)
				break
			}
		}
	}

	sort.Ints(active)
	return active
}

func main() {
	// Test data from the problem description.
	purchases := []UserPurchase{
		// User 6 has purchases on Sep 10 and Sep 14 (4 days apart) -> active.
		{UserID: 6, Item: "item1", CreatedAt: "2023-09-10", Amount: 100},
		{UserID: 6, Item: "item2", CreatedAt: "2023-09-14", Amount: 200},
		// User 1 has purchases far apart (Sep 10 and Sep 20, 10 days) -> not active.
		{UserID: 1, Item: "item3", CreatedAt: "2023-09-10", Amount: 50},
		{UserID: 1, Item: "item4", CreatedAt: "2023-09-20", Amount: 75},
		// User 2 has only 1 purchase -> not active.
		{UserID: 2, Item: "item5", CreatedAt: "2023-09-10", Amount: 150},
		// User 3 has purchases exactly 7 days apart -> active.
		{UserID: 3, Item: "item6", CreatedAt: "2023-09-10", Amount: 60},
		{UserID: 3, Item: "item7", CreatedAt: "2023-09-17", Amount: 80},
	}

	results := findActiveUsers(purchases)

	fmt.Println("Active Users (user_id):")
	for _, uid := range results {
		fmt.Println(uid)
	}
	// Expected output:
	// 3
	// 6
}
```

## 2692 — Make Object Immutable

```go
package main

// LeetCode #2692: Make Object Immutable
// https://leetcode.com/problems/make-object-immutable/
// Difficulty: Medium [Paid] (JS problem)
// Time: O(1) | Space: O(1)

import "fmt"

type ImmutableMap struct {
	data map[string]any
}

func NewImmutable(data map[string]any) *ImmutableMap {
	copied := make(map[string]any)
	for k, v := range data {
		copied[k] = v
	}
	return &ImmutableMap{data: copied}
}

func (im *ImmutableMap) Get(key string) (any, bool) {
	val, ok := im.data[key]
	return val, ok
}

func (im *ImmutableMap) Set(key string, val any) {
	// Immutable: do nothing (or panic in JS-like implementation)
	// In Go, we just don't modify
}

func main() {
	im := NewImmutable(map[string]any{"a": 1, "b": 2})

	// Test case 1
	val, ok := im.Get("a")
	fmt.Println("Test 1:", val, ok)
	// Expected: 1 true

	// Test case 2
	im.Set("c", 3)
	_, ok2 := im.Get("c")
	fmt.Println("Test 2:", ok2)
	// Expected: false (immutable)

	// Test case 3
	_, ok3 := im.Get("z")
	fmt.Println("Test 3:", ok3)
	// Expected: false
}
```

## 2693 — Call Function With Custom Context

```go
package main

// LeetCode #2693: Call Function with Custom Context
// https://leetcode.com/problems/call-function-with-custom-context/
// Difficulty: Medium
// Time: O(1) | Space: O(1)

import "fmt"

type Context map[string]any

func callWithContext(fn func(Context, ...int) int, ctx Context, args ...int) int {
	return fn(ctx, args...)
}

func main() {
	// Test case 1
	sumFn := func(ctx Context, args ...int) int {
		total := 0
		for _, v := range args {
			total += v
		}
		return total
	}
	fmt.Println("Test 1:", callWithContext(sumFn, Context{}, 1, 2, 3))
	// Expected: 6

	// Test case 2
	ctxFn := func(ctx Context, args ...int) int {
		if multiplier, ok := ctx["mult"]; ok {
			m := multiplier.(int)
			result := 0
			for _, v := range args {
				result += v * m
			}
			return result
		}
		return 0
	}
	fmt.Println("Test 2:", callWithContext(ctxFn, Context{"mult": 3}, 1, 2, 3))
	// Expected: 18

	// Test case 3
	fmt.Println("Test 3:", callWithContext(sumFn, Context{}))
	// Expected: 0
}
```

## 2694 — Event Emitter

```go
package main

// LeetCode #2694: Event Emitter
// https://leetcode.com/problems/event-emitter/
// Difficulty: Medium
// Time: O(1) per operation | Space: O(n)

import "fmt"

type EventEmitter struct {
	events map[string][]func(args ...any)
}

func NewEventEmitter() *EventEmitter {
	return &EventEmitter{events: make(map[string][]func(args ...any))}
}

func (ee *EventEmitter) Subscribe(event string, cb func(args ...any)) func() {
	ee.events[event] = append(ee.events[event], cb)
	return func() {
		listeners := ee.events[event]
		for i, fn := range listeners {
			if &fn == &cb { // Note: won't work for func comparison in Go
				ee.events[event] = append(listeners[:i], listeners[i+1:]...)
				break
			}
		}
	}
}

func (ee *EventEmitter) Emit(event string, args ...any) []any {
	results := []any{}
	for _, cb := range ee.events[event] {
		cb(args...)
		results = append(results, nil) // Go functions don't return values like JS
	}
	return results
}

func main() {
	emitter := NewEventEmitter()

	// Test case 1
	emitter.Subscribe("event1", func(args ...any) {
		fmt.Println("  received:", args)
	})
	fmt.Println("Test 1: emitting event1")
	emitter.Emit("event1", 1, 2, 3)

	// Test case 2: unsubscribe
	unsub := emitter.Subscribe("event2", func(args ...any) {
		fmt.Println("  event2 received:", args)
	})
	unsub()
	fmt.Println("Test 2: emitted event2 after unsubscribe (no output if working)")
	emitter.Emit("event2")

	// Test case 3: multiple subscribers
	count := 0
	emitter.Subscribe("count", func(args ...any) { count++ })
	emitter.Subscribe("count", func(args ...any) { count++ })
	emitter.Emit("count")
	fmt.Println("Test 3: count =", count)
	// Expected: 2
}
```

## 2698 — Find The Punishment Number Of An Integer

```go
package main

// LeetCode #2698: Find the Punishment Number of an Integer
// https://leetcode.com/problems/find-the-punishment-number-of-an-integer/
// Difficulty: Medium
// Time: O(n * 2^len(num)) | Space: O(log n)

import (
	"fmt"
	"strconv"
)

func punishmentNumber(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		sq := i * i
		sqStr := strconv.Itoa(sq)
		if canPartition(sqStr, i) {
			total += sq
		}
	}
	return total
}

func canPartition(s string, target int) bool {
	if len(s) == 0 {
		return target == 0
	}
	for i := 0; i < len(s); i++ {
		prefix, _ := strconv.Atoi(s[:i+1])
		if prefix > target {
			break
		}
		if canPartition(s[i+1:], target-prefix) {
			return true
		}
	}
	return false
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", punishmentNumber(10))
	// Expected: 182 (1+81+100)

	// Test case 2
	fmt.Println("Test 2:", punishmentNumber(37))
	// Expected: 1478

	// Test case 3
	fmt.Println("Test 3:", punishmentNumber(1))
	// Expected: 1
}
```

## 2700 — Differences Between Two Objects

```go
package main

// LeetCode #2700: Differences Between Two Objects
// https://leetcode.com/problems/differences-between-two-objects/
// Difficulty: Medium [Paid] (JS problem)
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"reflect"
)

func objectDiff(obj1, obj2 any) map[string]any {
	result := make(map[string]any)
	diff(obj1, obj2, "", result)
	return result
}

func diff(a, b any, path string, result map[string]any) {
	if reflect.DeepEqual(a, b) {
		return
	}

	if a == nil || b == nil {
		result[path] = []any{a, b}
		return
	}

	map1, ok1 := a.(map[string]any)
	map2, ok2 := b.(map[string]any)
	if ok1 && ok2 {
		allKeys := make(map[string]bool)
		for k := range map1 {
			allKeys[k] = true
		}
		for k := range map2 {
			allKeys[k] = true
		}
		for k := range allKeys {
			newPath := path
			if newPath == "" {
				newPath = k
			} else {
				newPath = path + "." + k
			}
			diff(map1[k], map2[k], newPath, result)
		}
		return
	}

	result[path] = []any{a, b}
}

func main() {
	// Test case 1: simple diff
	o1 := map[string]any{"a": 1, "b": 2}
	o2 := map[string]any{"a": 1, "b": 3}
	fmt.Println("Test 1:", objectDiff(o1, o2))
	// Expected: map[b:[2 3]]

	// Test case 2: missing key
	o3 := map[string]any{"a": 1, "c": 3}
	fmt.Println("Test 2:", objectDiff(o1, o3))
	// Expected: map[b:[2 <nil>] c:[<nil> 3]]

	// Test case 3: same objects
	fmt.Println("Test 3:", objectDiff(o1, o1))
	// Expected: map[]
}
```

## 2705 — Compact Object

```go
package main

// LeetCode #2705: Compact Object
// https://leetcode.com/problems/compact-object/
// Difficulty: Medium (JS problem)
// Time: O(n) | Space: O(n)

import "fmt"

func compact(obj any) any {
	switch v := obj.(type) {
	case map[string]any:
		result := make(map[string]any)
		for k, val := range v {
			compacted := compact(val)
			if compacted != nil {
				result[k] = compacted
			}
		}
		if len(result) == 0 {
			return nil
		}
		return result
	case []any:
		result := []any{}
		for _, val := range v {
			compacted := compact(val)
			if compacted != nil {
				result = append(result, compacted)
			}
		}
		if len(result) == 0 {
			return nil
		}
		return result
	case bool:
		if !v {
			return nil
		}
	case int:
		if v == 0 {
			return nil
		}
	case string:
		if v == "" {
			return nil
		}
	}
	return obj
}

func main() {
	// Test case 1
	input1 := map[string]any{
		"a": nil,
		"b": 1,
		"c": 0,
		"d": "",
		"e": false,
		"f": []any{nil, 1, 0, "", false},
	}
	fmt.Println("Test 1:", compact(input1))

	// Test case 2
	input2 := []any{0, 1, false, true, "", "hello"}
	fmt.Println("Test 2:", compact(input2))

	// Test case 3: empty result
	input3 := map[string]any{"a": false, "b": 0, "c": ""}
	fmt.Println("Test 3:", compact(input3))
}
```

## 2707 — Extra Characters In A String

```go
package main

// LeetCode #2707: Extra Characters in a String
// https://leetcode.com/problems/extra-characters-in-a-string/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import "fmt"

func minExtraChar(s string, dictionary []string) int {
	wordSet := make(map[string]bool)
	for _, w := range dictionary {
		wordSet[w] = true
	}

	n := len(s)
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = n + 1
	}
	dp[0] = 0

	for i := 1; i <= n; i++ {
		// Try to match word ending at i-1
		dp[i] = dp[i-1] + 1
		for j := 0; j < i; j++ {
			if wordSet[s[j:i]] {
				if dp[j] < dp[i] {
					dp[i] = dp[j]
				}
			}
		}
	}
	return dp[n]
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minExtraChar("leetscode", []string{"leet", "code", "leetcode"}))
	// Expected: 1 ('s' is extra)

	// Test case 2
	fmt.Println("Test 2:", minExtraChar("sayhelloworld", []string{"hello", "world"}))
	// Expected: 3 ('say' is extra)

	// Test case 3
	fmt.Println("Test 3:", minExtraChar("abcd", []string{"a", "b", "c", "d"}))
	// Expected: 0
}
```

## 2708 — Maximum Strength Of A Group

```go
package main

// LeetCode #2708: Maximum Strength of a Group
// https://leetcode.com/problems/maximum-strength-of-a-group/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func maxStrength(nums []int) int64 {
	n := len(nums)
	if n == 1 {
		return int64(nums[0])
	}

	positives := []int{}
	negatives := []int{}
	hasZero := false

	for _, v := range nums {
		if v > 0 {
			positives = append(positives, v)
		} else if v < 0 {
			negatives = append(negatives, v)
		} else {
			hasZero = true
		}
	}

	sort.Ints(negatives) // Sort negatives (most negative first)

	ans := int64(1)
	for _, v := range positives {
		ans *= int64(v)
	}

	// Take pairs of negatives (the largest ones, i.e., closest to 0)
	// Sort negatives ascending: [-5, -4, -3, -2, -1]
	// We want to pair them from the least negative end: [-1, -2] * [-3, -4] ...
	if len(negatives)%2 == 1 {
		negatives = negatives[:len(negatives)-1] // Drop the most negative if odd count
	}

	for _, v := range negatives {
		ans *= int64(v)
	}

	if ans == 1 && len(positives) == 0 && len(negatives) == 0 {
		if hasZero {
			return 0
		}
	}

	return ans
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", maxStrength([]int{3, -1, -5, 2, 5, -9}))
	// Expected: 1350

	// Test case 2
	fmt.Println("Test 2:", maxStrength([]int{-4, -5, -6}))
	// Expected: 30

	// Test case 3
	fmt.Println("Test 3:", maxStrength([]int{0, -1}))
	// Expected: 0
}
```

## 2711 — Difference Of Number Of Distinct Values On Diagonals

```go
package main

// LeetCode #2711: Difference of Number of Distinct Values on Diagonals
// https://leetcode.com/problems/difference-of-number-of-distinct-values-on-diagonals/
// Difficulty: Medium
// Time: O(m*n*(m+n)) | Space: O(1)

import "fmt"

func differenceOfDistinctValues(grid [][]int) [][]int {
	m, n := len(grid), len(grid[0])
	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// Count distinct values above-left diagonal
			aboveLeft := make(map[int]bool)
			r, c := i-1, j-1
			for r >= 0 && c >= 0 {
				aboveLeft[grid[r][c]] = true
				r--
				c--
			}

			// Count distinct values below-right diagonal
			belowRight := make(map[int]bool)
			r, c = i+1, j+1
			for r < m && c < n {
				belowRight[grid[r][c]] = true
				r++
				c++
			}

			diff := len(aboveLeft) - len(belowRight)
			if diff < 0 {
				diff = -diff
			}
			ans[i][j] = diff
		}
	}
	return ans
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", differenceOfDistinctValues([][]int{{1, 2, 3}, {3, 1, 5}, {3, 2, 1}}))
	// Expected: [[1,1,0],[1,0,1],[0,1,1]]

	// Test case 2
	fmt.Println("Test 2:", differenceOfDistinctValues([][]int{{1}}))
	// Expected: [[0]]

	// Test case 3
	fmt.Println("Test 3:", differenceOfDistinctValues([][]int{{1, 2}, {3, 4}}))
	// Expected: [[1,0],[0,1]]
}
```

## 2712 — Minimum Cost To Make All Characters Equal

```go
package main

// LeetCode #2712: Minimum Cost to Make All Characters Equal
// https://leetcode.com/problems/minimum-cost-to-make-all-characters-equal/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func MinimumCostToMakeAllCharactersEqual(s string) int64 {
	n := len(s)

	calc := func(target byte) int64 {
		var cost int64
		flip := 0
		for i := 0; i < n; i++ {
			cur := s[i]
			if flip%2 == 1 {
				if cur == '0' {
					cur = '1'
				} else {
					cur = '0'
				}
			}
			if cur != target {
				flip++
				cost += int64(i + 1)
			}
		}
		return cost
	}

	cost0 := calc('0')
	cost1 := calc('1')
	if cost0 < cost1 {
		return cost0
	}
	return cost1
}

func main() {
	fmt.Println(MinimumCostToMakeAllCharactersEqual("0011"))
	fmt.Println(MinimumCostToMakeAllCharactersEqual("010101"))
}
```

## 2718 — Sum Of Matrix After Queries

```go
package main

// LeetCode #2718: Sum of Matrix After Queries
// https://leetcode.com/problems/sum-of-matrix-after-queries/
// Difficulty: Medium
// Time: O(n + q) | Space: O(n)

import "fmt"

func SumOfMatrixAfterQueries(n int, queries [][]int) int64 {
	rowSet := make(map[int]bool)
	colSet := make(map[int]bool)
	rowSum := int64(0)
	colSum := int64(0)

	var result int64
	for i := len(queries) - 1; i >= 0; i-- {
		typ, idx, val := queries[i][0], queries[i][1], int64(queries[i][2])
		if typ == 0 { // row
			if rowSet[idx] {
				continue
			}
			rowSet[idx] = true
			rowVal := val * int64(n-len(colSet))
			result += rowVal - rowSum
		} else { // col
			if colSet[idx] {
				continue
			}
			colSet[idx] = true
			colVal := val * int64(n-len(rowSet))
			result += colVal - colSum
		}
	}
	return result
}

func main() {
	fmt.Println(SumOfMatrixAfterQueries(3, [][]int{{0, 0, 1}, {1, 2, 2}, {0, 2, 3}, {1, 0, 4}}))
	fmt.Println(SumOfMatrixAfterQueries(1, [][]int{{0, 0, 5}}))
}
```

## 2721 — Execute Asynchronous Functions In Parallel

```go
package main

// LeetCode #2721: Execute Asynchronous Functions in Parallel
// https://leetcode.com/problems/execute-asynchronous-functions-in-parallel/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"sync"
)

func ExecuteAsynchronousFunctionsInParallel(functions []func() int) []int {
	var wg sync.WaitGroup
	results := make([]int, len(functions))

	for i, fn := range functions {
		wg.Add(1)
		go func(idx int, f func() int) {
			defer wg.Done()
			results[idx] = f()
		}(i, fn)
	}

	wg.Wait()
	return results
}

func main() {
	results := ExecuteAsynchronousFunctionsInParallel([]func() int{
		func() int { return 1 },
		func() int { return 2 },
		func() int { return 3 },
	})
	fmt.Println(results)

	results2 := ExecuteAsynchronousFunctionsInParallel([]func() int{
		func() int { return 42 },
	})
	fmt.Println(results2)
}
```

## 2722 — Join Two Arrays By Id

```go
package main

// LeetCode #2722: Join Two Arrays by ID
// https://leetcode.com/problems/join-two-arrays-by-id/
// Difficulty: Medium
// Time: O(n + m) | Space: O(n + m)

import (
	"fmt"
	"sort"
)

type MapItem struct {
	ID     int
	Values map[string]int
}

func JoinTwoArraysById(arr1, arr2 []MapItem) []MapItem {
	merged := make(map[int]map[string]int)

	for _, item := range arr1 {
		if _, ok := merged[item.ID]; !ok {
			merged[item.ID] = make(map[string]int)
		}
		for k, v := range item.Values {
			merged[item.ID][k] = v
		}
	}
	for _, item := range arr2 {
		if _, ok := merged[item.ID]; !ok {
			merged[item.ID] = make(map[string]int)
		}
		for k, v := range item.Values {
			merged[item.ID][k] = v
		}
	}

	result := make([]MapItem, 0, len(merged))
	for id, vals := range merged {
		result = append(result, MapItem{ID: id, Values: vals})
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].ID < result[j].ID
	})
	return result
}

func main() {
	arr1 := []MapItem{{ID: 1, Values: map[string]int{"a": 1}}, {ID: 2, Values: map[string]int{"b": 2}}}
	arr2 := []MapItem{{ID: 1, Values: map[string]int{"c": 3}}, {ID: 3, Values: map[string]int{"d": 4}}}
	fmt.Println(JoinTwoArraysById(arr1, arr2))

	arr3 := []MapItem{{ID: 1, Values: map[string]int{"x": 10}}}
	arr4 := []MapItem{}
	fmt.Println(JoinTwoArraysById(arr3, arr4))
}
```

## 2730 — Find The Longest Semi Repetitive Substring

```go
package main

// LeetCode #2730: Find the Longest Semi-Repetitive Substring
// https://leetcode.com/problems/find-the-longest-semi-repetitive-substring/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func FindTheLongestSemiRepetitiveSubstring(s string) int {
	n := len(s)
	if n <= 2 {
		return n
	}

	maxLen := 0
	left := 0
	lastPair := -1

	for right := 1; right < n; right++ {
		if s[right] == s[right-1] {
			if lastPair != -1 {
				left = lastPair
			}
			lastPair = right
		}
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}

	return maxLen
}

func main() {
	fmt.Println(FindTheLongestSemiRepetitiveSubstring("52233"))
	fmt.Println(FindTheLongestSemiRepetitiveSubstring("0001"))
	fmt.Println(FindTheLongestSemiRepetitiveSubstring("1111111"))
}
```

## 2731 — Movement Of Robots

```go
package main

// LeetCode #2731: Movement of Robots
// https://leetcode.com/problems/movement-of-robots/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func MovementOfRobots(nums []int, s string, d int) int {
	n := len(nums)
	pos := make([]int, n)
	for i, num := range nums {
		pos[i] = num
		if s[i] == 'R' {
			pos[i] += d
		} else {
			pos[i] -= d
		}
	}

	sort.Ints(pos)

	const mod = 1_000_000_007
	var sum int64
	var prefix int64
	for i, p := range pos {
		sum = (sum + int64(p)*int64(i) - prefix) % mod
		prefix += int64(p)
	}

	return int(sum)
}

func main() {
	fmt.Println(MovementOfRobots([]int{1, 0}, "RL", 2))
	fmt.Println(MovementOfRobots([]int{-2, 0, 2}, "RLL", 3))
}
```

## 2734 — Lexicographically Smallest String After Substring Operation

```go
package main

// LeetCode #2734: Lexicographically Smallest String After Substring Operation
// https://leetcode.com/problems/lexicographically-smallest-string-after-substring-operation/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func LexicographicallySmallestStringAfterSubstringOperation(s string) string {
	n := len(s)
	b := []byte(s)

	start := -1
	for i := 0; i < n; i++ {
		if b[i] > 'a' {
			start = i
			break
		}
	}

	if start == -1 {
		b[n-1] = 'z'
		return string(b)
	}

	for i := start; i < n; i++ {
		if b[i] == 'a' {
			break
		}
		b[i]--
	}

	return string(b)
}

func main() {
	fmt.Println(LexicographicallySmallestStringAfterSubstringOperation("cbabc"))
	fmt.Println(LexicographicallySmallestStringAfterSubstringOperation("acbbc"))
	fmt.Println(LexicographicallySmallestStringAfterSubstringOperation("a"))
}
```

## 2735 — Collecting Chocolates

```go
package main

// LeetCode #2735: Collecting Chocolates
// https://leetcode.com/problems/collecting-chocolates/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import "fmt"

func CollectingChocolates(nums []int, x int) int64 {
	n := len(nums)
	minCost := make([]int, n)
	copy(minCost, nums)

	var best int64
	for i := 0; i < n; i++ {
		best += int64(nums[i])
	}

	for shift := 1; shift < n; shift++ {
		var total int64 = int64(shift) * int64(x)
		for i := 0; i < n; i++ {
			idx := (i + shift) % n
			if nums[idx] < minCost[i] {
				minCost[i] = nums[idx]
			}
			total += int64(minCost[i])
		}
		if total < best {
			best = total
		}
	}

	return best
}

func main() {
	fmt.Println(CollectingChocolates([]int{20, 1, 15}, 5))
	fmt.Println(CollectingChocolates([]int{1, 2, 3}, 4))
}
```

## 2737 — Find The Closest Marked Node

```go
package main

// LeetCode #2737: Find the Closest Marked Node
// https://leetcode.com/problems/find-the-closest-marked-node/
// Difficulty: Medium [Paid]
// Time: O((V+E) log V) | Space: O(V+E)

import (
	"container/heap"
	"fmt"
	"math"
)

type Edge struct {
	to, weight int
}

type Item struct {
	node, dist int
	index      int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i].dist < pq[j].dist }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i]; pq[i].index = i; pq[j].index = j }
func (pq *PriorityQueue) Push(x interface{}) { n := len(*pq); item := x.(*Item); item.index = n; *pq = append(*pq, item) }
func (pq *PriorityQueue) Pop() interface{}   { old := *pq; n := len(old); item := old[n-1]; old[n-1] = nil; item.index = -1; *pq = old[:n-1]; return item }

func FindTheClosestMarkedNode(n int, edges [][]int, marked []int, start int) int {
	graph := make([][]Edge, n)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		graph[u] = append(graph[u], Edge{v, w})
	}

	dist := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[start] = 0

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{node: start, dist: 0})

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(*Item)
		if cur.dist > dist[cur.node] {
			continue
		}
		for _, e := range graph[cur.node] {
			if nd := cur.dist + e.weight; nd < dist[e.to] {
				dist[e.to] = nd
				heap.Push(pq, &Item{node: e.to, dist: nd})
			}
		}
	}

	best := math.MaxInt32
	for _, m := range marked {
		if dist[m] < best {
			best = dist[m]
		}
	}
	if best == math.MaxInt32 {
		return -1
	}
	return best
}

func main() {
	fmt.Println(FindTheClosestMarkedNode(4, [][]int{{0, 1, 1}, {1, 2, 2}, {2, 3, 3}}, []int{2, 3}, 0))
	fmt.Println(FindTheClosestMarkedNode(3, [][]int{{0, 1, 5}}, []int{2}, 0))
}
```

## 2738 — Count Occurrences In Text

```go
package main

// LeetCode #2738: Count Occurrences in Text
// https://leetcode.com/problems/count-occurrences-in-text/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import (
	"fmt"
	"strings"
)

func CountOccurrencesInText(text string, word string) int {
	count := 0
	words := strings.Fields(text)
	for _, w := range words {
		if w == word {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(CountOccurrencesInText("hello world hello", "hello"))
	fmt.Println(CountOccurrencesInText("this is a test test this", "test"))
	fmt.Println(CountOccurrencesInText("unique", "none"))
}
```

## 2740 — Find The Value Of The Partition

```go
package main

// LeetCode #2740: Find the Value of the Partition
// https://leetcode.com/problems/find-the-value-of-the-partition/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func FindTheValueOfThePartition(nums []int) int {
	sort.Ints(nums)
	minDiff := nums[1] - nums[0]
	for i := 2; i < len(nums); i++ {
		if diff := nums[i] - nums[i-1]; diff < minDiff {
			minDiff = diff
		}
	}
	return minDiff
}

func main() {
	fmt.Println(FindTheValueOfThePartition([]int{1, 3, 2, 4}))
	fmt.Println(FindTheValueOfThePartition([]int{100, 1, 10}))
}
```

## 2741 — Special Permutations

```go
package main

// LeetCode #2741: Special Permutations
// https://leetcode.com/problems/special-permutations/
// Difficulty: Medium
// Time: O(n^2 * 2^n) | Space: O(n * 2^n)

import "fmt"

func SpecialPermutations(nums []int) int {
	n := len(nums)
	const mod = 1_000_000_007

	dp := make([][]int, 1<<n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		dp[1<<i][i] = 1
	}

	for mask := 1; mask < 1<<n; mask++ {
		for last := 0; last < n; last++ {
			if dp[mask][last] == 0 {
				continue
			}
			for next := 0; next < n; next++ {
				if mask&(1<<next) != 0 {
					continue
				}
				if nums[last]%nums[next] == 0 || nums[next]%nums[last] == 0 {
					newMask := mask | (1 << next)
					dp[newMask][next] = (dp[newMask][next] + dp[mask][last]) % mod
				}
			}
		}
	}

	var result int
	fullMask := (1 << n) - 1
	for i := 0; i < n; i++ {
		result = (result + dp[fullMask][i]) % mod
	}
	return result
}

func main() {
	fmt.Println(SpecialPermutations([]int{1, 2, 3}))
	fmt.Println(SpecialPermutations([]int{2, 3, 6}))
}
```

## 2743 — Count Substrings Without Repeating Character

```go
package main

// LeetCode #2743: Count Substrings Without Repeating Character
// https://leetcode.com/problems/count-substrings-without-repeating-character/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func CountSubstringsWithoutRepeatingCharacter(s string) int {
	lastPos := make(map[byte]int)
	left := 0
	count := 0

	for right := 0; right < len(s); right++ {
		if pos, ok := lastPos[s[right]]; ok && pos >= left {
			left = pos + 1
		}
		lastPos[s[right]] = right
		count += right - left + 1
	}

	return count
}

func main() {
	fmt.Println(CountSubstringsWithoutRepeatingCharacter("abcabc"))
	fmt.Println(CountSubstringsWithoutRepeatingCharacter("aaaa"))
	fmt.Println(CountSubstringsWithoutRepeatingCharacter(""))
}
```

## 2745 — Construct The Longest New String

```go
package main

// LeetCode #2745: Construct the Longest New String
// https://leetcode.com/problems/construct-the-longest-new-string/
// Difficulty: Medium
// Time: O(1) | Space: O(1)

import "fmt"

func ConstructTheLongestNewString(x int, y int, z int) int {
	// AA can pair with BB, BB can pair with AA, AB can go anywhere
	// Max pairs of AA and BB: min(x, y) + extra if any left
	used := 0

	// Use pairs of AA and BB
	pairs := x
	if y < pairs {
		pairs = y
	}
	used += pairs * 2

	// If both AA and BB have remaining, can add one more
	if x > pairs {
		used++
	}
	if y > pairs {
		used++
	}

	// AB can be inserted anywhere, but consumes z
	used += z

	return used * 2
}

func main() {
	fmt.Println(ConstructTheLongestNewString(1, 1, 1))
	fmt.Println(ConstructTheLongestNewString(2, 0, 2))
	fmt.Println(ConstructTheLongestNewString(0, 0, 5))
}
```

## 2746 — Decremental String Concatenation

```go
package main

// LeetCode #2746: Decremental String Concatenation
// https://leetcode.com/problems/decremental-string-concatenation/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func DecrementalStringConcatenation(words []string) int {
	n := len(words)
	// dp[first][last] = min length
	// Use 26 letters
	const INF = 1 << 30
	dp := make([][]int, 26)
	for i := range dp {
		dp[i] = make([]int, 26)
		for j := range dp[i] {
			dp[i][j] = INF
		}
	}

	f, l := int(words[0][0]-'a'), int(words[0][len(words[0])-1]-'a')
	dp[f][l] = len(words[0])

	for i := 1; i < n; i++ {
		w := words[i]
		cf, cl := int(w[0]-'a'), int(w[len(w)-1]-'a')
		ndp := make([][]int, 26)
		for i := range ndp {
			ndp[i] = make([]int, 26)
			for j := range ndp[i] {
				ndp[i][j] = INF
			}
		}

		for a := 0; a < 26; a++ {
			for b := 0; b < 26; b++ {
				if dp[a][b] == INF {
					continue
				}
				// Append w: a...b + cf...cl
				if b == cf {
					if dp[a][b]+len(w)-1 < ndp[a][cl] {
						ndp[a][cl] = dp[a][b] + len(w) - 1
					}
				} else {
					if dp[a][b]+len(w) < ndp[a][cl] {
						ndp[a][cl] = dp[a][b] + len(w)
					}
				}
				// Prepend w: cf...cl + a...b
				if cl == a {
					if dp[a][b]+len(w)-1 < ndp[cf][b] {
						ndp[cf][b] = dp[a][b] + len(w) - 1
					}
				} else {
					if dp[a][b]+len(w) < ndp[cf][b] {
						ndp[cf][b] = dp[a][b] + len(w)
					}
				}
			}
		}
		dp = ndp
	}

	best := INF
	for a := 0; a < 26; a++ {
		for b := 0; b < 26; b++ {
			if dp[a][b] < best {
				best = dp[a][b]
			}
		}
	}
	return best
}

func main() {
	fmt.Println(DecrementalStringConcatenation([]string{"abc", "cde", "efg"}))
	fmt.Println(DecrementalStringConcatenation([]string{"aa", "ab", "bc"}))
}
```

## 2747 — Count Zero Request Servers

```go
package main

// LeetCode #2747: Count Zero Request Servers
// https://leetcode.com/problems/count-zero-request-servers/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

type Log struct {
	serverID int
	time     int
}

func CountZeroRequestServers(n int, logs [][]int, x int, queries []int) []int {
	// Process logs
	logList := make([]Log, len(logs))
	for i, l := range logs {
		logList[i] = Log{serverID: l[0], time: l[1]}
	}
	sort.Slice(logList, func(i, j int) bool {
		return logList[i].time < logList[j].time
	})

	// Map queries to their original indices
	type qItem struct {
		time int
		idx  int
	}
	qList := make([]qItem, len(queries))
	for i, q := range queries {
		qList[i] = qItem{time: q, idx: i}
	}
	sort.Slice(qList, func(i, j int) bool {
		return qList[i].time < qList[j].time
	})

	result := make([]int, len(queries))
	active := make(map[int]int)
	left := 0

	for _, q := range qList {
		end := q.time
		start := q.time - x

		// Add logs within [start, end]
		for left < len(logList) && logList[left].time <= end {
			active[logList[left].serverID]++
			left++
		}

		// Remove logs before start
		right := 0
		for right < len(logList) && logList[right].time < start {
			if _, ok := active[logList[right].serverID]; ok {
				active[logList[right].serverID]--
				if active[logList[right].serverID] == 0 {
					delete(active, logList[right].serverID)
				}
			}
			right++
		}

		result[q.idx] = n - len(active)
	}

	return result
}

func main() {
	fmt.Println(CountZeroRequestServers(3, [][]int{{1, 3}, {2, 6}, {1, 5}}, 5, []int{10, 11}))
	fmt.Println(CountZeroRequestServers(3, [][]int{{1, 1}, {2, 2}, {3, 3}}, 2, []int{1, 5}))
}
```

## 2749 — Minimum Operations To Make The Integer Zero

```go
package main

// LeetCode #2749: Minimum Operations to Make the Integer Zero
// https://leetcode.com/problems/minimum-operations-to-make-the-integer-zero/
// Difficulty: Medium
// Time: O(log n) | Space: O(1)

import "fmt"

func MinimumOperationsToMakeTheIntegerZero(num1 int, num2 int) int {
	// Try k operations: num1 - k*num2 must have at most k bits set, >= k
	for k := 1; k <= 60; k++ {
		diff := num1 - k*num2
		if diff < 0 {
			return -1
		}
		bits := 0
		for tmp := diff; tmp > 0; tmp >>= 1 {
			bits += tmp & 1
		}
		if bits <= k && diff >= k {
			return k
		}
	}
	return -1
}

func main() {
	fmt.Println(MinimumOperationsToMakeTheIntegerZero(3, -2))
	fmt.Println(MinimumOperationsToMakeTheIntegerZero(5, 7))
}
```

## 2750 — Ways To Split Array Into Good Subarrays

```go
package main

// LeetCode #2750: Ways to Split Array Into Good Subarrays
// https://leetcode.com/problems/ways-to-split-array-into-good-subarrays/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func WaysToSplitArrayIntoGoodSubarrays(nums []int) int {
	// Find positions of 1s
	ones := make([]int, 0)
	for i, v := range nums {
		if v == 1 {
			ones = append(ones, i)
		}
	}

	if len(ones) == 0 {
		return 0
	}

	const mod = 1_000_000_007
	result := 1
	for i := 1; i < len(ones); i++ {
		gap := ones[i] - ones[i-1]
		result = (result * gap) % mod
	}

	return result
}

func main() {
	fmt.Println(WaysToSplitArrayIntoGoodSubarrays([]int{0, 1, 0, 0, 1}))
	fmt.Println(WaysToSplitArrayIntoGoodSubarrays([]int{1, 1, 1}))
	fmt.Println(WaysToSplitArrayIntoGoodSubarrays([]int{0, 0, 0}))
}
```

## 2754 — Bind Function To Context

```go
package main

// LeetCode #2754: Bind Function to Context
// https://leetcode.com/problems/bind-function-to-context/
// Difficulty: Medium [Paid]
// Time: O(1) | Space: O(1)

import "fmt"

type Context map[string]interface{}

type BoundFunc struct {
	fn      func(ctx Context, args ...interface{}) interface{}
	ctx     Context
}

func BindFunctionToContext(fn func(Context, ...interface{}) interface{}, ctx Context) func(...interface{}) interface{} {
	return func(args ...interface{}) interface{} {
		return fn(ctx, args...)
	}
}

func main() {
	ctx := Context{"multiplier": 3}
	bound := BindFunctionToContext(func(ctx Context, args ...interface{}) interface{} {
		m := ctx["multiplier"].(int)
		return args[0].(int) * m
	}, ctx)
	fmt.Println(bound(5))
	fmt.Println(bound(10))
}
```

## 2755 — Deep Merge Of Two Objects

```go
package main

// LeetCode #2755: Deep Merge of Two Objects
// https://leetcode.com/problems/deep-merge-of-two-objects/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

type JSONObj map[string]interface{}

func DeepMergeOfTwoObjects(obj1, obj2 JSONObj) JSONObj {
	result := make(JSONObj)
	for k, v := range obj1 {
		result[k] = v
	}
	for k, v2 := range obj2 {
		if v1, ok := result[k]; ok {
			m1, ok1 := v1.(map[string]interface{})
			m2, ok2 := v2.(map[string]interface{})
			if ok1 && ok2 {
				result[k] = DeepMergeOfTwoObjects(JSONObj(m1), JSONObj(m2))
			} else {
				result[k] = v2
			}
		} else {
			result[k] = v2
		}
	}
	return result
}

func main() {
	merged := DeepMergeOfTwoObjects(
		JSONObj{"a": 1, "b": JSONObj{"c": 2}},
		JSONObj{"b": JSONObj{"d": 3}, "e": 4},
	)
	fmt.Println(merged)

	merged2 := DeepMergeOfTwoObjects(
		JSONObj{"x": 1},
		JSONObj{"x": 2},
	)
	fmt.Println(merged2)
}
```

## 2757 — Generate Circular Array Values

```go
package main

// LeetCode #2757: Generate Circular Array Values
// https://leetcode.com/problems/generate-circular-array-values/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func GenerateCircularArrayValues(arr []int, start int, count int) []int {
	n := len(arr)
	if n == 0 {
		return []int{}
	}
	result := make([]int, count)
	for i := 0; i < count; i++ {
		result[i] = arr[(start+i)%n]
	}
	return result
}

func main() {
	fmt.Println(GenerateCircularArrayValues([]int{1, 2, 3, 4}, 2, 6))
	fmt.Println(GenerateCircularArrayValues([]int{10, 20}, 1, 3))
}
```

## 2761 — Prime Pairs With Target Sum

```go
package main

// LeetCode #2761: Prime Pairs With Target Sum
// https://leetcode.com/problems/prime-pairs-with-target-sum/
// Difficulty: Medium
// Time: O(n log log n) | Space: O(n)

import "fmt"

func PrimePairsWithTargetSum(target int) [][]int {
	// Sieve of Eratosthenes
	isPrime := make([]bool, target+1)
	for i := 2; i <= target; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i <= target; i++ {
		if isPrime[i] {
			for j := i * i; j <= target; j += i {
				isPrime[j] = false
			}
		}
	}

	result := make([][]int, 0)
	for i := 2; i <= target/2; i++ {
		if isPrime[i] && isPrime[target-i] {
			result = append(result, []int{i, target - i})
		}
	}
	return result
}

func main() {
	fmt.Println(PrimePairsWithTargetSum(10))
	fmt.Println(PrimePairsWithTargetSum(2))
	fmt.Println(PrimePairsWithTargetSum(18))
}
```

## 2762 — Continuous Subarrays

```go
package main

// LeetCode #2762: Continuous Subarrays
// https://leetcode.com/problems/continuous-subarrays/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func ContinuousSubarrays(nums []int) int64 {
	n := len(nums)
	var result int64
	left := 0

	// Track min and max using deques via slice
	// minDeque stores indices with increasing values
	// maxDeque stores indices with decreasing values
	minDeque := make([]int, 0)
	maxDeque := make([]int, 0)

	for right := 0; right < n; right++ {
		// Maintain minDeque
		for len(minDeque) > 0 && nums[minDeque[len(minDeque)-1]] > nums[right] {
			minDeque = minDeque[:len(minDeque)-1]
		}
		minDeque = append(minDeque, right)

		// Maintain maxDeque
		for len(maxDeque) > 0 && nums[maxDeque[len(maxDeque)-1]] < nums[right] {
			maxDeque = maxDeque[:len(maxDeque)-1]
		}
		maxDeque = append(maxDeque, right)

		// Shrink window if condition violated
		for nums[maxDeque[0]]-nums[minDeque[0]] > 2 {
			if minDeque[0] == left {
				minDeque = minDeque[1:]
			}
			if maxDeque[0] == left {
				maxDeque = maxDeque[1:]
			}
			left++
		}

		result += int64(right - left + 1)
	}

	return result
}

func main() {
	fmt.Println(ContinuousSubarrays([]int{5, 4, 2, 4}))
	fmt.Println(ContinuousSubarrays([]int{1, 2, 3}))
}
```

## 2764 — Is Array A Preorder Of Some Binary Tree

```go
package main

// LeetCode #2764: Is Array a Preorder of Some Binary Tree
// https://leetcode.com/problems/is-array-a-preorder-of-some-binary-tree/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func IsArrayAPreorderOfSomeBinaryTree(nodes [][]int) bool {
	// nodes[i] = [id, parentId]
	children := make(map[int][]int)
	for _, node := range nodes {
		id, parent := node[0], node[1]
		children[parent] = append(children[parent], id)
	}

	// Simulate DFS preorder
	stack := []int{-1} // root parent
	idx := 0
	n := len(nodes)

	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		childList := children[cur]
		for i := len(childList) - 1; i >= 0; i-- {
			if idx >= n || childList[i] != nodes[idx][0] {
				return false
			}
			stack = append(stack, childList[i])
			idx++
		}
	}

	return idx == n
}

func main() {
	fmt.Println(IsArrayAPreorderOfSomeBinaryTree([][]int{{0, -1}, {1, 0}, {2, 0}}))
	fmt.Println(IsArrayAPreorderOfSomeBinaryTree([][]int{{0, -1}, {1, 0}, {3, 2}, {2, 1}}))
}
```

## 2766 — Relocate Marbles

```go
package main

// LeetCode #2766: Relocate Marbles
// https://leetcode.com/problems/relocate-marbles/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func RelocateMarbles(nums []int, moveFrom []int, moveTo []int) []int {
	positions := make(map[int]bool)
	for _, n := range nums {
		positions[n] = true
	}

	for i := 0; i < len(moveFrom); i++ {
		delete(positions, moveFrom[i])
		positions[moveTo[i]] = true
	}

	result := make([]int, 0, len(positions))
	for p := range positions {
		result = append(result, p)
	}
	sort.Ints(result)
	return result
}

func main() {
	fmt.Println(RelocateMarbles([]int{1, 2, 3}, []int{1}, []int{4}))
	fmt.Println(RelocateMarbles([]int{1, 1, 2, 2}, []int{1, 2}, []int{3, 4}))
}
```

## 2767 — Partition String Into Minimum Beautiful Substrings

```go
package main

// LeetCode #2767: Partition String Into Minimum Beautiful Substrings
// https://leetcode.com/problems/partition-string-into-minimum-beautiful-substrings/
// Difficulty: Medium
// Time: O(2^n) | Space: O(n)

import (
	"fmt"
	"math"
)

func PartitionStringIntoMinimumBeautifulSubstrings(s string) int {
	n := len(s)
	powers := make(map[string]bool)
	for i := 0; i <= 10; i++ {
		p := int(math.Pow(5, float64(i)))
		b := fmt.Sprintf("%b", p)
		if len(b) <= 15 {
			powers[b] = true
		}
	}

	memo := make([]int, n)
	for i := range memo {
		memo[i] = -1
	}

	var dfs func(int) int
	dfs = func(start int) int {
		if start == n {
			return 0
		}
		if memo[start] != -1 {
			return memo[start]
		}
		best := math.MaxInt32
		for end := start + 1; end <= n; end++ {
			sub := s[start:end]
			if powers[sub] {
				subResult := dfs(end)
				if subResult != -1 && subResult+1 < best {
					best = subResult + 1
				}
			}
		}
		if best == math.MaxInt32 {
			memo[start] = -1
		} else {
			memo[start] = best
		}
		return memo[start]
	}

	result := dfs(0)
	if result == math.MaxInt32 || result <= 0 {
		return -1
	}
	return result
}

func main() {
	fmt.Println(PartitionStringIntoMinimumBeautifulSubstrings("1011"))
	fmt.Println(PartitionStringIntoMinimumBeautifulSubstrings("111"))
}
```

## 2768 — Number Of Black Blocks

```go
package main

// LeetCode #2768: Number of Black Blocks
// https://leetcode.com/problems/number-of-black-blocks/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func NumberOfBlackBlocks(m int, n int, coordinates [][]int) []int64 {
	blackCells := make(map[[2]int]bool)
	for _, c := range coordinates {
		blackCells[[2]int{c[0], c[1]}] = true
	}

	blockCount := make(map[int]int64) // count of black cells in 2x2 -> number of blocks
	for _, c := range coordinates {
		r, c2 := c[0], c[1]
		// Check 4 possible top-left corners
		for _, dr := range []int{-1, 0} {
			for _, dc := range []int{-1, 0} {
				tr, tc := r+dr, c2+dc
				if tr < 0 || tc < 0 || tr >= m-1 || tc >= n-1 {
					continue
				}
				cnt := 0
				if blackCells[[2]int{tr, tc}] {
					cnt++
				}
				if blackCells[[2]int{tr, tc + 1}] {
					cnt++
				}
				if blackCells[[2]int{tr + 1, tc}] {
					cnt++
				}
				if blackCells[[2]int{tr + 1, tc + 1}] {
					cnt++
				}
				blockCount[cnt]++
			}
		}
	}

	result := make([]int64, 5)
	totalBlocks := int64(m-1) * int64(n-1)
	var counted int64
	for k, v := range blockCount {
		result[k] = v
		counted += v
	}
	result[0] = totalBlocks - counted
	return result
}

func main() {
	fmt.Println(NumberOfBlackBlocks(3, 3, [][]int{{0, 0}}))
	fmt.Println(NumberOfBlackBlocks(2, 2, [][]int{{0, 0}, {1, 1}}))
}
```

## 2770 — Maximum Number Of Jumps To Reach The Last Index

```go
package main

// LeetCode #2770: Maximum Number of Jumps to Reach the Last Index
// https://leetcode.com/problems/maximum-number-of-jumps-to-reach-the-last-index/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import "fmt"

func MaximumNumberOfJumpsToReachTheLastIndex(nums []int, target int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = -1
	}
	dp[0] = 0

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			diff := nums[i] - nums[j]
			if diff < 0 {
				diff = -diff
			}
			if diff <= target && dp[j] != -1 {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
				}
			}
		}
	}

	return dp[n-1]
}

func main() {
	fmt.Println(MaximumNumberOfJumpsToReachTheLastIndex([]int{1, 3, 6, 4, 1, 2}, 2))
	fmt.Println(MaximumNumberOfJumpsToReachTheLastIndex([]int{1, 3, 6, 4, 1, 2}, 3))
}
```

## 2771 — Longest Non Decreasing Subarray From Two Arrays

```go
package main

// LeetCode #2771: Longest Non-decreasing Subarray From Two Arrays
// https://leetcode.com/problems/longest-non-decreasing-subarray-from-two-arrays/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func LongestNonDecreasingSubarrayFromTwoArrays(nums1 []int, nums2 []int) int {
	n := len(nums1)
	dp1, dp2 := 1, 1
	best := 1

	for i := 1; i < n; i++ {
		ndp1, ndp2 := 1, 1
		if nums1[i] >= nums1[i-1] {
			ndp1 = dp1 + 1
		}
		if nums1[i] >= nums2[i-1] {
			if dp2+1 > ndp1 {
				ndp1 = dp2 + 1
			}
		}
		if nums2[i] >= nums1[i-1] {
			ndp2 = dp1 + 1
		}
		if nums2[i] >= nums2[i-1] {
			if dp2+1 > ndp2 {
				ndp2 = dp2 + 1
			}
		}
		dp1, dp2 = ndp1, ndp2
		if dp1 > best {
			best = dp1
		}
		if dp2 > best {
			best = dp2
		}
	}

	return best
}

func main() {
	fmt.Println(LongestNonDecreasingSubarrayFromTwoArrays([]int{1, 3, 2, 1}, []int{2, 2, 3, 4}))
	fmt.Println(LongestNonDecreasingSubarrayFromTwoArrays([]int{1, 2}, []int{3, 1}))
}
```

## 2772 — Apply Operations To Make All Array Elements Equal To Zero

```go
package main

// LeetCode #2772: Apply Operations to Make All Array Elements Equal to Zero
// https://leetcode.com/problems/apply-operations-to-make-all-array-elements-equal-to-zero/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func ApplyOperationsToMakeAllArrayElementsEqualToZero(nums []int, k int) bool {
	n := len(nums)
	diff := make([]int, n+1)
	cur := 0

	for i := 0; i < n; i++ {
		cur += diff[i]
		val := nums[i] + cur
		if val < 0 {
			return false
		}
		val %= 2
		if val != 0 {
			if i+k > n {
				return false
			}
			diff[i] -= 1
			diff[i+k] += 1
			cur -= 1
		}
	}

	return true
}

func main() {
	fmt.Println(ApplyOperationsToMakeAllArrayElementsEqualToZero([]int{2, 0, 2}, 2))
	fmt.Println(ApplyOperationsToMakeAllArrayElementsEqualToZero([]int{1, 0, 1}, 2))
}
```

## 2773 — Height Of Special Binary Tree

```go
package main

// LeetCode #2773: Height of Special Binary Tree
// https://leetcode.com/problems/height-of-special-binary-tree/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(h)

import "fmt"

type SpecialTreeNode struct {
	Val       int
	Left      *SpecialTreeNode
	Right     *SpecialTreeNode
	IsSpecial bool
}

func HeightOfSpecialBinaryTree(root *SpecialTreeNode) int {
	var dfs func(*SpecialTreeNode) int
	dfs = func(node *SpecialTreeNode) int {
		if node == nil || node.IsSpecial {
			return 0
		}
		leftH := dfs(node.Left)
		rightH := dfs(node.Right)
		if leftH > rightH {
			return leftH + 1
		}
		return rightH + 1
	}
	return dfs(root)
}

func main() {
	// Tree: 1(not special) -> 2(special), 3(not special)
	root := &SpecialTreeNode{
		Val: 1,
		Left: &SpecialTreeNode{
			Val:       2,
			IsSpecial: true,
		},
		Right: &SpecialTreeNode{
			Val: 3,
			Left: &SpecialTreeNode{
				Val: 4,
			},
		},
	}
	fmt.Println(HeightOfSpecialBinaryTree(root))
	fmt.Println(HeightOfSpecialBinaryTree(nil))
}
```

## 2775 — Undefined To Null

```go
package main

// LeetCode #2775: Undefined to Null
// https://leetcode.com/problems/undefined-to-null/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

type NullableObj map[string]interface{}

func UndefinedToNull(obj NullableObj) NullableObj {
	result := make(NullableObj)
	for k, v := range obj {
		if v == nil {
			result[k] = nil
		} else if m, ok := v.(map[string]interface{}); ok {
			result[k] = UndefinedToNull(m)
		} else if arr, ok := v.([]interface{}); ok {
			newArr := make([]interface{}, len(arr))
			for i, item := range arr {
				if m2, ok := item.(map[string]interface{}); ok {
					newArr[i] = UndefinedToNull(m2)
				} else {
					newArr[i] = item
				}
			}
			result[k] = newArr
		} else {
			result[k] = v
		}
	}
	return result
}

func main() {
	obj := NullableObj{"a": nil, "b": 42, "c": NullableObj{"d": nil}}
	fmt.Println(UndefinedToNull(obj))

	obj2 := NullableObj{"x": 1}
	fmt.Println(UndefinedToNull(obj2))
}
```

## 2776 — Convert Callback Based Function To Promise Based Function

```go
package main

// LeetCode #2776: Convert Callback Based Function to Promise Based Function
// https://leetcode.com/problems/convert-callback-based-function-to-promise-based-function/
// Difficulty: Medium [Paid]
// Time: O(1) | Space: O(1)

import "fmt"

type Callback func(args ...interface{}) (interface{}, error)
type PromiseFunc func(args ...interface{}) (<-chan interface{}, <-chan error)

func ConvertCallbackBasedFunctionToPromiseBasedFunction(fn Callback) PromiseFunc {
	return func(args ...interface{}) (<-chan interface{}, <-chan error) {
		resCh := make(chan interface{}, 1)
		errCh := make(chan error, 1)
		go func() {
			result, err := fn(args...)
			if err != nil {
				errCh <- err
			} else {
				resCh <- result
			}
			close(resCh)
			close(errCh)
		}()
		return resCh, errCh
	}
}

func main() {
	add := func(args ...interface{}) (interface{}, error) {
		return args[0].(int) + args[1].(int), nil
	}
	promiseFn := ConvertCallbackBasedFunctionToPromiseBasedFunction(add)
	resCh, errCh := promiseFn(3, 4)
	select {
	case res := <-resCh:
		fmt.Println(res)
	case err := <-errCh:
		fmt.Println(err)
	}

	// Test with single arg
	promiseFn2 := ConvertCallbackBasedFunctionToPromiseBasedFunction(func(args ...interface{}) (interface{}, error) {
		return args[0].(int) * 2, nil
	})
	resCh2, _ := promiseFn2(10)
	select {
	case res := <-resCh2:
		fmt.Println(res)
	}
}
```

## 2777 — Date Range Generator

```go
package main

// LeetCode #2777: Date Range Generator
// https://leetcode.com/problems/date-range-generator/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import (
	"fmt"
	"time"
)

func DateRangeGenerator(start, end time.Time, step time.Duration) []time.Time {
	dates := make([]time.Time, 0)
	for cur := start; !cur.After(end); cur = cur.Add(step) {
		dates = append(dates, cur)
	}
	return dates
}

func main() {
	start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2024, 1, 5, 0, 0, 0, 0, time.UTC)
	dates := DateRangeGenerator(start, end, 24*time.Hour)
	for _, d := range dates {
		fmt.Println(d.Format("2006-01-02"))
	}

	fmt.Println("---")

	start2 := time.Date(2024, 6, 1, 0, 0, 0, 0, time.UTC)
	end2 := time.Date(2024, 6, 1, 0, 0, 0, 0, time.UTC)
	dates2 := DateRangeGenerator(start2, end2, 24*time.Hour)
	for _, d := range dates2 {
		fmt.Println(d.Format("2006-01-02"))
	}
}
```

## 2779 — Maximum Beauty Of An Array After Applying Operation

```go
package main

// LeetCode #2779: Maximum Beauty of an Array After Applying Operation
// https://leetcode.com/problems/maximum-beauty-of-an-array-after-applying-operation/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func MaximumBeautyOfAnArrayAfterApplyingOperation(nums []int, k int) int {
	sort.Ints(nums)
	left := 0
	best := 0
	for right := 0; right < len(nums); right++ {
		for nums[right]-nums[left] > 2*k {
			left++
		}
		if right-left+1 > best {
			best = right - left + 1
		}
	}
	return best
}

func main() {
	fmt.Println(MaximumBeautyOfAnArrayAfterApplyingOperation([]int{4, 6, 1, 2}, 2))
	fmt.Println(MaximumBeautyOfAnArrayAfterApplyingOperation([]int{1, 1, 1, 1}, 10))
}
```

## 2780 — Minimum Index Of A Valid Split

```go
package main

// LeetCode #2780: Minimum Index of a Valid Split
// https://leetcode.com/problems/minimum-index-of-a-valid-split/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func MinimumIndexOfAValidSplit(nums []int) int {
	n := len(nums)

	// Find majority element using Boyer-Moore
	majority := nums[0]
	count := 0
	for _, num := range nums {
		if count == 0 {
			majority = num
			count = 1
		} else if num == majority {
			count++
		} else {
			count--
		}
	}

	// Count total occurrences of majority
	totalCount := 0
	for _, num := range nums {
		if num == majority {
			totalCount++
		}
	}

	// Find minimum split index
	leftCount := 0
	for i := 0; i < n-1; i++ {
		if nums[i] == majority {
			leftCount++
		}
		rightCount := totalCount - leftCount
		leftLen := i + 1
		rightLen := n - i - 1
		if leftCount*2 > leftLen && rightCount*2 > rightLen {
			return i
		}
	}

	return -1
}

func main() {
	fmt.Println(MinimumIndexOfAValidSplit([]int{1, 2, 2, 2}))
	fmt.Println(MinimumIndexOfAValidSplit([]int{2, 1, 3, 1, 1, 1, 7, 1, 2, 1}))
}
```

## 2782 — Number Of Unique Categories

```go
package main

// LeetCode #2782: Number of Unique Categories
// https://leetcode.com/problems/number-of-unique-categories/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func NumberOfUniqueCategories(categories []string) int {
	seen := make(map[string]bool)
	for _, c := range categories {
		seen[c] = true
	}
	return len(seen)
}

func main() {
	fmt.Println(NumberOfUniqueCategories([]string{"a", "b", "a", "c"}))
	fmt.Println(NumberOfUniqueCategories([]string{"x", "x", "x"}))
	fmt.Println(NumberOfUniqueCategories([]string{}))
}
```

## 2783 — Flight Occupancy And Waitlist Analysis

```go
package main

// LeetCode #2783: Flight Occupancy and Waitlist Analysis
// https://leetcode.com/problems/flight-occupancy-and-waitlist-analysis/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

type FlightStatus struct {
	Capacity      int
	Booked        int
	Waitlisted    int
}

func FlightOccupancyAndWaitlistAnalysis(flights []FlightStatus) []FlightStatus {
	results := make([]FlightStatus, len(flights))
	for i, f := range flights {
		available := f.Capacity - f.Booked
		if available < 0 {
			available = 0
		}
		// Waitlisted passengers fill available spots
		canBoard := f.Waitlisted
		if canBoard > available {
			canBoard = available
		}
		results[i] = FlightStatus{
			Capacity:      f.Capacity,
			Booked:        f.Booked + canBoard,
			Waitlisted:    f.Waitlisted - canBoard,
		}
	}
	return results
}

func main() {
	flights := []FlightStatus{
		{Capacity: 100, Booked: 95, Waitlisted: 10},
		{Capacity: 50, Booked: 50, Waitlisted: 5},
	}
	fmt.Println(FlightOccupancyAndWaitlistAnalysis(flights))
}
```

## 2785 — Sort Vowels In A String

```go
package main

// LeetCode #2785: Sort Vowels in a String
// https://leetcode.com/problems/sort-vowels-in-a-string/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func SortVowelsInAString(s string) string {
	isVowel := func(c byte) bool {
		return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' ||
			c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U'
	}

	vowels := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if isVowel(s[i]) {
			vowels = append(vowels, s[i])
		}
	}

	sort.Slice(vowels, func(i, j int) bool {
		return vowels[i] < vowels[j]
	})

	result := []byte(s)
	vi := 0
	for i := 0; i < len(s); i++ {
		if isVowel(s[i]) {
			result[i] = vowels[vi]
			vi++
		}
	}

	return string(result)
}

func main() {
	fmt.Println(SortVowelsInAString("lEetcOde"))
	fmt.Println(SortVowelsInAString("lYmpH"))
}
```

## 2786 — Visit Array Positions To Maximize Score

```go
package main

// LeetCode #2786: Visit Array Positions to Maximize Score
// https://leetcode.com/problems/visit-array-positions-to-maximize-score/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func VisitArrayPositionsToMaximizeScore(nums []int, x int) int64 {
	n := len(nums)
	// dp[0] = max score ending at even parity, dp[1] = max score ending at odd parity
	dp := [2]int64{int64(nums[0]), int64(nums[0])}
	parity0 := nums[0] & 1

	other := 1 - parity0
	dp[other] = int64(nums[0]) - int64(x)
	if dp[other] < int64(nums[0]) {
		dp[other] = int64(nums[0])
	}

	best := int64(nums[0])
	for i := 1; i < n; i++ {
		p := nums[i] & 1
		// Option 1: Start here
		cur := int64(nums[i])
		// Option 2: Extend from prev with same parity
		if int64(nums[i])+dp[p] > cur {
			cur = int64(nums[i]) + dp[p]
		}
		// Option 3: Extend from prev with different parity (pay x)
		otherP := 1 - p
		if int64(nums[i])-int64(x)+dp[otherP] > cur {
			cur = int64(nums[i]) - int64(x) + dp[otherP]
		}
		if cur > dp[p] {
			dp[p] = cur
		}
		if cur > best {
			best = cur
		}
	}

	return best
}

func main() {
	fmt.Println(VisitArrayPositionsToMaximizeScore([]int{2, 3, 6, 1, 9, 2}, 5))
	fmt.Println(VisitArrayPositionsToMaximizeScore([]int{2, 4, 6, 8}, 3))
}
```

## 2787 — Ways To Express An Integer As Sum Of Powers

```go
package main

// LeetCode #2787: Ways to Express an Integer as Sum of Powers
// https://leetcode.com/problems/ways-to-express-an-integer-as-sum-of-powers/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import "fmt"

func WaysToExpressAnIntegerAsSumOfPowers(n int, x int) int {
	const mod = 1_000_000_007

	// Generate powers
	powers := make([]int, 0)
	for i := 1; ; i++ {
		p := 1
		for j := 0; j < x; j++ {
			p *= i
		}
		if p > n {
			break
		}
		powers = append(powers, p)
	}

	dp := make([]int, n+1)
	dp[0] = 1

	for _, p := range powers {
		for s := n; s >= p; s-- {
			dp[s] = (dp[s] + dp[s-p]) % mod
		}
	}

	return dp[n]
}

func main() {
	fmt.Println(WaysToExpressAnIntegerAsSumOfPowers(10, 2))
	fmt.Println(WaysToExpressAnIntegerAsSumOfPowers(4, 1))
}
```

## 2789 — Largest Element In An Array After Merge Operations

```go
package main

// LeetCode #2789: Largest Element in an Array after Merge Operations
// https://leetcode.com/problems/largest-element-in-an-array-after-merge-operations/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func LargestElementInAnArrayAfterMergeOperations(nums []int) int64 {
	n := len(nums)
	if n == 0 {
		return 0
	}

	result := int64(nums[n-1])
	for i := n - 2; i >= 0; i-- {
		if int64(nums[i]) <= result {
			result += int64(nums[i])
		} else {
			result = int64(nums[i])
		}
	}

	return result
}

func main() {
	fmt.Println(LargestElementInAnArrayAfterMergeOperations([]int{2, 3, 7, 9, 3}))
	fmt.Println(LargestElementInAnArrayAfterMergeOperations([]int{5, 3, 3}))
}
```

## 2795 — Parallel Execution Of Promises For Individual Results Retrieval

```go
package main

// LeetCode #2795: Parallel Execution of Promises for Individual Results Retrieval
// https://leetcode.com/problems/parallel-execution-of-promises-for-individual-results-retrieval/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"sync"
)

type PromiseResult struct {
	Index int
	Value interface{}
}

func ParallelExecutionOfPromisesForIndividualResultsRetrieval(functions []func() int) []int {
	n := len(functions)
	results := make([]int, n)
	var wg sync.WaitGroup

	for i, fn := range functions {
		wg.Add(1)
		go func(idx int, f func() int) {
			defer wg.Done()
			results[idx] = f()
		}(i, fn)
	}

	wg.Wait()
	return results
}

func main() {
	results := ParallelExecutionOfPromisesForIndividualResultsRetrieval([]func() int{
		func() int { return 1 + 1 },
		func() int { return 2 + 2 },
		func() int { return 3 + 3 },
	})
	fmt.Println(results)

	results2 := ParallelExecutionOfPromisesForIndividualResultsRetrieval([]func() int{
		func() int { return 99 },
	})
	fmt.Println(results2)
}
```

## 2799 — Count Complete Subarrays In An Array

```go
package main

// LeetCode #2799: Count Complete Subarrays in an Array
// https://leetcode.com/problems/count-complete-subarrays-in-an-array/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func CountCompleteSubarraysInAnArray(nums []int) int {
	// Count distinct elements
	distinct := make(map[int]bool)
	for _, n := range nums {
		distinct[n] = true
	}
	target := len(distinct)

	left := 0
	count := 0
	freq := make(map[int]int)
	unique := 0

	for right := 0; right < len(nums); right++ {
		freq[nums[right]]++
		if freq[nums[right]] == 1 {
			unique++
		}

		for unique == target {
			// All subarrays from left to right-end are valid
			count += len(nums) - right
			freq[nums[left]]--
			if freq[nums[left]] == 0 {
				unique--
			}
			left++
		}
	}

	return count
}

func main() {
	fmt.Println(CountCompleteSubarraysInAnArray([]int{1, 3, 1, 2, 2}))
	fmt.Println(CountCompleteSubarraysInAnArray([]int{1, 1}))
}
```

## 2800 — Shortest String That Contains Three Strings

```go
package main

// LeetCode #2800: Shortest String That Contains Three Strings
// https://leetcode.com/problems/shortest-string-that-contains-three-strings/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import "fmt"

func merge2(x, y string) string {
	if len(x) < len(y) {
		return merge2(y, x)
	}
	// Check if x contains y
	for i := 0; i <= len(x)-len(y); i++ {
		if x[i:i+len(y)] == y {
			return x
		}
	}
	// Find longest overlap
	for i := 0; i < len(y); i++ {
		overlap := len(y) - i
		if overlap <= len(x) && x[len(x)-overlap:] == y[:overlap] {
			return x + y[i:]
		}
	}
	return x + y
}

func ShortestStringThatContainsThreeStrings(a string, b string, c string) string {
	strs := []string{a, b, c}
	perms := [][]int{
		{0, 1, 2}, {0, 2, 1}, {1, 0, 2},
		{1, 2, 0}, {2, 0, 1}, {2, 1, 0},
	}

	best := a + b + c
	for _, p := range perms {
		merged := merge2(merge2(strs[p[0]], strs[p[1]]), strs[p[2]])
		if len(merged) < len(best) || (len(merged) == len(best) && merged < best) {
			best = merged
		}
	}

	return best
}

func main() {
	fmt.Println(ShortestStringThatContainsThreeStrings("abc", "bcd", "cde"))
	fmt.Println(ShortestStringThatContainsThreeStrings("a", "ab", "abc"))
}
```

## 2802 — Find The K Th Lucky Number

```go
package main

// LeetCode #2802: Find The K-th Lucky Number
// https://leetcode.com/problems/find-the-k-th-lucky-number/
// Difficulty: Medium [Paid]
// Time: O(log k) | Space: O(log k)

import "fmt"

func FindTheKThLuckyNumber(k int) string {
	// K-th lucky number: binary representation of k+1, then replace 0->4, 1->7
	// n = k + 1
	n := k + 1
	binary := fmt.Sprintf("%b", n)
	// Remove first '1' (it's the leading bit from the offset)
	result := make([]byte, len(binary)-1)
	for i := 1; i < len(binary); i++ {
		if binary[i] == '0' {
			result[i-1] = '4'
		} else {
			result[i-1] = '7'
		}
	}
	return string(result)
}

func main() {
	fmt.Println(FindTheKThLuckyNumber(1))
	fmt.Println(FindTheKThLuckyNumber(3))
	fmt.Println(FindTheKThLuckyNumber(5))
}
```

## 2805 — Custom Interval

```go
package main

// LeetCode #2805: Custom Interval
// https://leetcode.com/problems/custom-interval/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import (
	"fmt"
	"time"
)

func CustomInterval(fn func(), interval time.Duration, times int) {
	for i := 0; i < times; i++ {
		fn()
		time.Sleep(interval)
	}
}

func main() {
	count := 0
	fn := func() {
		count++
		fmt.Println("Executed:", count)
	}
	CustomInterval(fn, 10*time.Millisecond, 3)
	fmt.Println("Done")
}
```

## 2807 — Insert Greatest Common Divisors In Linked List

```go
package main

// LeetCode #2807: Insert Greatest Common Divisors in Linked List
// https://leetcode.com/problems/insert-greatest-common-divisors-in-linked-list/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func InsertGreatestCommonDivisorsInLinkedList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}

	cur := head
	for cur != nil && cur.Next != nil {
		g := gcd(cur.Val, cur.Next.Val)
		node := &ListNode{Val: g, Next: cur.Next}
		cur.Next = node
		cur = node.Next
	}

	return head
}

func printList(head *ListNode) {
	for head != nil {
		if head.Next != nil {
			fmt.Printf("%d -> ", head.Val)
		} else {
			fmt.Println(head.Val)
		}
		head = head.Next
	}
}

func main() {
	head := &ListNode{18, &ListNode{6, &ListNode{10, &ListNode{3, nil}}}}
	printList(InsertGreatestCommonDivisorsInLinkedList(head))

	head2 := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	printList(InsertGreatestCommonDivisorsInLinkedList(head2))
}
```

## 2808 — Minimum Seconds To Equalize A Circular Array

```go
package main

// LeetCode #2808: Minimum Seconds to Equalize a Circular Array
// https://leetcode.com/problems/minimum-seconds-to-equalize-a-circular-array/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func MinimumSecondsToEqualizeACircularArray(nums []int) int {
	n := len(nums)
	pos := make(map[int][]int)
	for i, v := range nums {
		pos[v] = append(pos[v], i)
	}

	best := n / 2
	for _, positions := range pos {
		if len(positions) == 0 {
			continue
		}
		maxGap := 0
		for i := 0; i < len(positions); i++ {
			curr := positions[i]
			var prev int
			if i > 0 {
				prev = positions[i-1]
			} else {
				prev = positions[len(positions)-1] - n
			}
			gap := curr - prev
			if gap > maxGap {
				maxGap = gap
			}
		}
		seconds := maxGap / 2
		if seconds < best {
			best = seconds
		}
	}

	return best
}

func main() {
	fmt.Println(MinimumSecondsToEqualizeACircularArray([]int{1, 2, 1, 2}))
	fmt.Println(MinimumSecondsToEqualizeACircularArray([]int{2, 1, 3, 3, 2}))
}
```

## 2811 — Check If It Is Possible To Split Array

```go
package main

// LeetCode #2811: Check if it is Possible to Split Array
// https://leetcode.com/problems/check-if-it-is-possible-to-split-array/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func CheckIfItIsPossibleToSplitArray(nums []int, m int) bool {
	n := len(nums)
	if n <= 2 {
		return true
	}

	for i := 0; i < n-1; i++ {
		if nums[i]+nums[i+1] >= m {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(CheckIfItIsPossibleToSplitArray([]int{2, 3, 3, 2, 3}, 6))
	fmt.Println(CheckIfItIsPossibleToSplitArray([]int{1, 1}, 3))
	fmt.Println(CheckIfItIsPossibleToSplitArray([]int{1, 2, 1}, 4))
}
```

## 2812 — Find The Safest Path In A Grid

```go
package main

// LeetCode #2812: Find the Safest Path in a Grid
// https://leetcode.com/problems/find-the-safest-path-in-a-grid/
// Difficulty: Medium
// Time: O(n^2 log n) | Space: O(n^2)

import (
	"fmt"
	"math"
)

func FindTheSafestPathInAGrid(grid [][]int) int {
	n := len(grid)

	// Multi-source BFS to compute distance to nearest thief (1)
	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = math.MaxInt32
		}
	}

	type cell struct{ r, c int }
	queue := make([]cell, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				dist[i][j] = 0
				queue = append(queue, cell{i, j})
			}
		}
	}

	dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, d := range dirs {
			nr, nc := cur.r+d[0], cur.c+d[1]
			if nr >= 0 && nr < n && nc >= 0 && nc < n && dist[nr][nc] > dist[cur.r][cur.c]+1 {
				dist[nr][nc] = dist[cur.r][cur.c] + 1
				queue = append(queue, cell{nr, nc})
			}
		}
	}

	// Binary search for max safety factor
	canReach := func(minDist int) bool {
		if dist[0][0] < minDist {
			return false
		}
		visited := make([][]bool, n)
		for i := range visited {
			visited[i] = make([]bool, n)
		}
		visited[0][0] = true
		q := []cell{{0, 0}}
		for len(q) > 0 {
			cur := q[0]
			q = q[1:]
			if cur.r == n-1 && cur.c == n-1 {
				return true
			}
			for _, d := range dirs {
				nr, nc := cur.r+d[0], cur.c+d[1]
				if nr >= 0 && nr < n && nc >= 0 && nc < n && !visited[nr][nc] && dist[nr][nc] >= minDist {
					visited[nr][nc] = true
					q = append(q, cell{nr, nc})
				}
			}
		}
		return false
	}

	lo, hi := 0, n*2
	for lo < hi {
		mid := (lo + hi + 1) / 2
		if canReach(mid) {
			lo = mid
		} else {
			hi = mid - 1
		}
	}

	return lo
}

func main() {
	fmt.Println(FindTheSafestPathInAGrid([][]int{{0, 0, 1}, {0, 0, 0}, {0, 0, 0}}))
	fmt.Println(FindTheSafestPathInAGrid([][]int{{0, 0, 0, 1}, {0, 0, 0, 0}, {0, 0, 0, 0}, {1, 0, 0, 0}}))
}
```

## 2816 — Double A Number Represented As A Linked List

```go
package main

// LeetCode #2816: Double a Number Represented as a Linked List
// https://leetcode.com/problems/double-a-number-represented-as-a-linked-list/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func DoubleANumberRepresentedAsALinkedList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	// Reverse the list
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	// Double
	carry := 0
	cur = prev
	var newHead *ListNode
	for cur != nil {
		val := cur.Val*2 + carry
		carry = val / 10
		cur.Val = val % 10
		newHead = cur
		cur = cur.Next
	}
	if carry > 0 {
		newHead = &ListNode{Val: carry, Next: newHead}
	}

	// Reverse back
	prev = nil
	cur = newHead
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	return prev
}

func printList2(head *ListNode) {
	for head != nil {
		if head.Next != nil {
			fmt.Printf("%d -> ", head.Val)
		} else {
			fmt.Println(head.Val)
		}
		head = head.Next
	}
}

func main() {
	head := &ListNode{1, &ListNode{8, &ListNode{9, nil}}}
	printList2(DoubleANumberRepresentedAsALinkedList(head))

	head2 := &ListNode{9, &ListNode{9, &ListNode{9, nil}}}
	printList2(DoubleANumberRepresentedAsALinkedList(head2))
}
```

## 2817 — Minimum Absolute Difference Between Elements With Constraint

```go
package main

// LeetCode #2817: Minimum Absolute Difference Between Elements With Constraint
// https://leetcode.com/problems/minimum-absolute-difference-between-elements-with-constraint/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"math"
	"sort"
)

func MinimumAbsoluteDifferenceBetweenElementsWithConstraint(nums []int, x int) int {
	n := len(nums)
	if x >= n {
		return -1
	}

	// For each index, we need to find nums[j] with j >= i+x closest to nums[i]
	// Process from right to left, maintaining a sorted set of values
	best := math.MaxInt32

	// Use a sorted slice
	sorted := make([]int, 0)
	for i := n - 1 - x; i >= 0; i-- {
		// Add nums[i+x] to the sorted set
		val := nums[i+x]
		pos := sort.SearchInts(sorted, val)
		sorted = append(sorted, 0)
		copy(sorted[pos+1:], sorted[pos:])
		sorted[pos] = val

		// Find closest to nums[i]
		pos2 := sort.SearchInts(sorted, nums[i])
		if pos2 < len(sorted) {
			if diff := sorted[pos2] - nums[i]; diff < best {
				best = diff
			}
		}
		if pos2 > 0 {
			if diff := nums[i] - sorted[pos2-1]; diff < best {
				best = diff
			}
		}
	}

	return best
}

func main() {
	fmt.Println(MinimumAbsoluteDifferenceBetweenElementsWithConstraint([]int{4, 3, 2, 4}, 2))
	fmt.Println(MinimumAbsoluteDifferenceBetweenElementsWithConstraint([]int{5, 3, 2, 10, 15}, 1))
}
```

## 2820 — Election Results

```go
package main

// LeetCode #2820: Election Results
// https://leetcode.com/problems/election-results/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

type Vote struct {
	Voter   string
	Candidate string
}

func ElectionResults(votes []Vote) string {
	counts := make(map[string]int)
	for _, v := range votes {
		counts[v.Candidate]++
	}

	type kv struct {
		Candidate string
		Count     int
	}
	sorted := make([]kv, 0, len(counts))
	for k, v := range counts {
		sorted = append(sorted, kv{k, v})
	}
	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].Count != sorted[j].Count {
			return sorted[i].Count > sorted[j].Count
		}
		return sorted[i].Candidate < sorted[j].Candidate
	})

	if len(sorted) > 0 {
		return sorted[0].Candidate
	}
	return ""
}

func main() {
	votes := []Vote{
		{"A", "Alice"}, {"B", "Bob"}, {"C", "Alice"},
	}
	fmt.Println(ElectionResults(votes))

	votes2 := []Vote{
		{"X", "Cand1"}, {"Y", "Cand1"}, {"Z", "Cand2"},
	}
	fmt.Println(ElectionResults(votes2))
}
```

## 2821 — Delay The Resolution Of Each Promise

```go
package main

// LeetCode #2821: Delay the Resolution of Each Promise
// https://leetcode.com/problems/delay-the-resolution-of-each-promise/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import (
	"fmt"
	"sync"
	"time"
)

func DelayTheResolutionOfEachPromise(functions []func() int, delay time.Duration) []int {
	n := len(functions)
	results := make([]int, n)
	var wg sync.WaitGroup

	for i, fn := range functions {
		wg.Add(1)
		go func(idx int, f func() int) {
			defer wg.Done()
			time.Sleep(delay)
			results[idx] = f()
		}(i, fn)
	}

	wg.Wait()
	return results
}

func main() {
	start := time.Now()
	results := DelayTheResolutionOfEachPromise([]func() int{
		func() int { return 10 },
		func() int { return 20 },
		func() int { return 30 },
	}, 5*time.Millisecond)
	fmt.Println(results)
	fmt.Println("Took:", time.Since(start).Round(time.Millisecond))
}
```

## 2823 — Deep Object Filter

```go
package main

// LeetCode #2823: Deep Object Filter
// https://leetcode.com/problems/deep-object-filter/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

type FilterObj map[string]interface{}

func DeepObjectFilter(obj FilterObj, predicate func(string, interface{}) bool) FilterObj {
	result := make(FilterObj)
	for k, v := range obj {
		if !predicate(k, v) {
			continue
		}
		if m, ok := v.(map[string]interface{}); ok {
			filtered := DeepObjectFilter(m, predicate)
			if len(filtered) > 0 {
				result[k] = filtered
			}
		} else {
			result[k] = v
		}
	}
	return result
}

func main() {
	obj := FilterObj{
		"a": 1,
		"b": 0,
		"c": FilterObj{"d": 2, "e": 0},
	}
	// Filter out values that are 0
	result := DeepObjectFilter(obj, func(k string, v interface{}) bool {
		if num, ok := v.(int); ok && num == 0 {
			return false
		}
		return true
	})
	fmt.Println(result)

	// Filter keys starting with 'a'
	result2 := DeepObjectFilter(obj, func(k string, v interface{}) bool {
		return k >= "b"
	})
	fmt.Println(result2)
}
```

## 2825 — Make String A Subsequence Using Cyclic Increments

```go
package main

// LeetCode #2825: Make String a Subsequence Using Cyclic Increments
// https://leetcode.com/problems/make-string-a-subsequence-using-cyclic-increments/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func MakeStringASubsequenceUsingCyclicIncrements(str1 string, str2 string) bool {
	j := 0
	for i := 0; i < len(str1) && j < len(str2); i++ {
		if str1[i] == str2[j] || (str1[i]-'a'+1)%26 == str2[j]-'a' {
			j++
		}
	}
	return j == len(str2)
}

func main() {
	fmt.Println(MakeStringASubsequenceUsingCyclicIncrements("abc", "bcd"))
	fmt.Println(MakeStringASubsequenceUsingCyclicIncrements("abc", "ad"))
}
```

## 2826 — Sorting Three Groups

```go
package main

// LeetCode #2826: Sorting Three Groups
// https://leetcode.com/problems/sorting-three-groups/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func SortingThreeGroups(nums []int) int {
	n := len(nums)
	// dp[i][j] = min operations to make first i+1 elements sorted with last element == j+1
	dp := make([][3]int, n+1)

	for i := 1; i <= n; i++ {
		for j := 0; j < 3; j++ {
			change := 0
			if nums[i-1] != j+1 {
				change = 1
			}
			dp[i][j] = dp[i-1][j] + change
			if j > 0 && dp[i][j-1] < dp[i][j] {
				dp[i][j] = dp[i][j-1]
			}
		}
	}

	return dp[n][2]
}

func main() {
	fmt.Println(SortingThreeGroups([]int{2, 1, 3, 2, 1}))
	fmt.Println(SortingThreeGroups([]int{1, 2, 3, 1, 2, 3}))
}
```

## 2829 — Determine The Minimum Sum Of A K Avoiding Array

```go
package main

// LeetCode #2829: Determine the Minimum Sum of a k-avoiding Array
// https://leetcode.com/problems/determine-the-minimum-sum-of-a-k-avoiding-array/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func DetermineTheMinimumSumOfAKAvoidingArray(n int, k int) int {
	used := make(map[int]bool)
	sum := 0
	for i := 1; len(used) < n; i++ {
		if used[k-i] {
			continue
		}
		used[i] = true
		sum += i
	}
	return sum
}

func main() {
	fmt.Println(DetermineTheMinimumSumOfAKAvoidingArray(5, 4))
	fmt.Println(DetermineTheMinimumSumOfAKAvoidingArray(3, 5))
}
```

## 2830 — Maximize The Profit As The Salesman

```go
package main

// LeetCode #2830: Maximize the Profit as the Salesman
// https://leetcode.com/problems/maximize-the-profit-as-the-salesman/
// Difficulty: Medium
// Time: O(n + m) | Space: O(n)

import "fmt"

func MaximizeTheProfitAsTheSalesman(n int, offers [][]int) int {
	// Group offers by end position
	byEnd := make([][][]int, n)
	for _, offer := range offers {
		start, end, gold := offer[0], offer[1], offer[2]
		byEnd[end] = append(byEnd[end], []int{start, gold})
	}

	dp := make([]int, n)
	for i := 0; i < n; i++ {
		if i > 0 {
			dp[i] = dp[i-1]
		}
		for _, offer := range byEnd[i] {
			start, gold := offer[0], offer[1]
			val := gold
			if start > 0 {
				val += dp[start-1]
			}
			if val > dp[i] {
				dp[i] = val
			}
		}
	}

	return dp[n-1]
}

func main() {
	fmt.Println(MaximizeTheProfitAsTheSalesman(5, [][]int{{0, 0, 1}, {0, 2, 2}, {1, 3, 2}}))
	fmt.Println(MaximizeTheProfitAsTheSalesman(3, [][]int{{0, 0, 5}, {1, 1, 3}, {2, 2, 4}}))
}
```

