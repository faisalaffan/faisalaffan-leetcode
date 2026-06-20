# Medium (Sedang) — Problem ��1664

## 1465 — Maximum Area Of A Piece Of Cake After Horizontal And Vertical Cuts

```go
package main

// LeetCode #1465: Maximum Area of a Piece of Cake After Horizontal and Vertical Cuts
// https://leetcode.com/problems/maximum-area-of-a-piece-of-cake-after-horizontal-and-vertical-cuts/
// Difficulty: Medium

import "fmt"
import "sort"

func main() {
	// Test case 1
	fmt.Println(maxArea(5, 4, []int{1, 2, 4}, []int{1, 3})) // 4

	// Test case 2
	fmt.Println(maxArea(5, 4, []int{3, 1}, []int{1})) // 6

	// Test case 3
	fmt.Println(maxArea(5, 4, []int{3}, []int{3})) // 9
}

const mod = 1_000_000_007

// Time: O(n log n + m log m) for sorting
// Space: O(1)
func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	sort.Ints(horizontalCuts)
	sort.Ints(verticalCuts)

	// Find max gap in horizontal cuts (including edges)
	maxHDiff := max(horizontalCuts[0], h-horizontalCuts[len(horizontalCuts)-1])
	for i := 1; i < len(horizontalCuts); i++ {
		diff := horizontalCuts[i] - horizontalCuts[i-1]
		if diff > maxHDiff {
			maxHDiff = diff
		}
	}

	// Find max gap in vertical cuts (including edges)
	maxVDiff := max(verticalCuts[0], w-verticalCuts[len(verticalCuts)-1])
	for i := 1; i < len(verticalCuts); i++ {
		diff := verticalCuts[i] - verticalCuts[i-1]
		if diff > maxVDiff {
			maxVDiff = diff
		}
	}

	area := (maxHDiff % mod) * (maxVDiff % mod) % mod
	return area
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

## 1466 — Reorder Routes To Make All Paths Lead To The City Zero

```go
package main

// LeetCode #1466: Reorder Routes to Make All Paths Lead to the City Zero
// https://leetcode.com/problems/reorder-routes-to-make-all-paths-lead-to-the-city-zero/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(minReorder(6, [][]int{{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5}})) // 3

	// Test case 2
	fmt.Println(minReorder(5, [][]int{{1, 0}, {1, 2}, {3, 2}, {3, 4}})) // 2

	// Test case 3
	fmt.Println(minReorder(3, [][]int{{1, 0}, {2, 0}})) // 0
}

// Time: O(n) where n = number of nodes
// Space: O(n) for adjacency list
func minReorder(n int, connections [][]int) int {
	// Build adjacency with direction info
	// For each edge, store [neighbor, direction]
	// direction=1 means original direction is away from 0, needs reorder
	// direction=0 means original direction is towards 0, ok
	adj := make([][][2]int, n)
	for _, conn := range connections {
		adj[conn[0]] = append(adj[conn[0]], [2]int{conn[1], 1}) // outgoing
		adj[conn[1]] = append(adj[conn[1]], [2]int{conn[0], 0}) // incoming
	}

	visited := make([]bool, n)
	changes := 0

	var dfs func(int)
	dfs = func(city int) {
		visited[city] = true
		for _, neighbor := range adj[city] {
			nextCity, needsReorder := neighbor[0], neighbor[1]
			if !visited[nextCity] {
				if needsReorder == 1 {
					changes++
				}
				dfs(nextCity)
			}
		}
	}

	dfs(0)
	return changes
}
```

## 1468 — Calculate Salaries

```go
package main

// LeetCode #1468: Calculate Salaries
// https://leetcode.com/problems/calculate-salaries/
// Difficulty: Medium

import "fmt"
import "sort"

func main() {
	// SQL problem - simulating in Go
	result := calculateSalaries(
		[]struct {
			companyID    int
			employeeID   int
			employeeName string
			salary       int
		}{
			{1, 1, "Alice", 100000},
			{1, 2, "Bob", 90000},
			{1, 3, "Charlie", 80000},
			{1, 4, "David", 70000},
			{1, 5, "Eve", 60000},
			{2, 6, "Frank", 120000},
			{2, 7, "Grace", 110000},
			{2, 8, "Henry", 100000},
		},
		[]struct {
			companyID int
			name      string
		}{
			{1, "Acme"},
			{2, "Globex"},
		},
	)
	for _, r := range result {
		fmt.Printf("%d %s %s %d\n", r.companyID, r.companyName, r.employeeName, r.tax)
	}
}

type taxResult struct {
	companyID     int
	companyName   string
	employeeName  string
	tax           int
}

// Time: O(n log n) for sorting
// Space: O(n)
func calculateSalaries(salaries []struct {
	companyID    int
	employeeID   int
	employeeName string
	salary       int
}, companies []struct {
	companyID int
	name      string
}) []taxResult {
	// Group salaries by company
	companySalaries := make(map[int][]struct {
		employeeID int
		name       string
		salary     int
	})
	for _, s := range salaries {
		companySalaries[s.companyID] = append(companySalaries[s.companyID], struct {
			employeeID int
			name       string
			salary     int
		}{s.employeeID, s.employeeName, s.salary})
	}

	companyNames := make(map[int]string)
	for _, c := range companies {
		companyNames[c.companyID] = c.name
	}

	var results []taxResult
	for companyID, emps := range companySalaries {
		// Sort by salary descending
		sort.Slice(emps, func(i, j int) bool {
			return emps[i].salary > emps[j].salary
		})

		totalEmp := len(emps)

		for _, emp := range emps {
			tax := emp.salary
			if emp.salary > 100000 {
				tax = emp.salary
			} else if emp.salary < 1000 {
				tax = 0
			} else {
				// Count employees with higher salary
				rank := 1
				for i := 0; i < totalEmp; i++ {
					if emps[i].salary > emp.salary {
						rank++
					}
				}
				// Max tax rate: 100000, reduced by 10% for each higher-rank employee
				maxTax := 100000
				reduction := (rank - 1) * 10000
				if rank >= 10 {
					tax = 0
				} else {
					tax = emp.salary
					mx := maxTax - reduction
					if tax > mx {
						tax = mx
					}
				}
			}
			if tax < 0 {
				tax = 0
			}

			results = append(results, taxResult{
				companyID,
				companyNames[companyID],
				emp.name,
				tax,
			})
		}
	}

	return results
}
```

## 1471 — The K Strongest Values In An Array

```go
package main

// LeetCode #1471: The k Strongest Values in an Array
// https://leetcode.com/problems/the-k-strongest-values-in-an-array/
// Difficulty: Medium

import "fmt"
import "sort"

func main() {
	// Test case 1
	fmt.Println(getStrongest([]int{1, 2, 3, 4, 5}, 2)) // [5,1]

	// Test case 2
	fmt.Println(getStrongest([]int{1, 1, 3, 5, 5}, 2)) // [5,5]

	// Test case 3
	fmt.Println(getStrongest([]int{6, 7, 11, 7, 6, 8}, 5)) // [11,8,6,6,7]

	// Test case 4
	fmt.Println(getStrongest([]int{6, -3, 7, 2, 11}, 3)) // [-3,11,2]
}

// Time: O(n log n) for sorting
// Space: O(1) for in-place sorting
func getStrongest(arr []int, k int) []int {
	sort.Ints(arr)
	n := len(arr)
	median := arr[(n-1)/2]

	// Sort by strength (|val - median|, then val)
	sort.Slice(arr, func(i, j int) bool {
		diffI := abs(arr[i] - median)
		diffJ := abs(arr[j] - median)
		if diffI != diffJ {
			return diffI > diffJ
		}
		return arr[i] > arr[j]
	})

	return arr[:k]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

## 1472 — Design Browser History

```go
package main

// LeetCode #1472: Design Browser History
// https://leetcode.com/problems/design-browser-history/
// Difficulty: Medium

import "fmt"

type BrowserHistory struct {
	history []string
	current int
}

func main() {
	bh := NewBrowserHistory("leetcode.com")
	bh.Visit("google.com")
	bh.Visit("facebook.com")
	bh.Visit("youtube.com")
	fmt.Println(bh.Back(1)) // "facebook.com"
	fmt.Println(bh.Back(1)) // "google.com"
	fmt.Println(bh.Forward(1)) // "facebook.com"
	bh.Visit("linkedin.com")
	fmt.Println(bh.Forward(2)) // "linkedin.com"
	fmt.Println(bh.Back(2)) // "google.com"
	fmt.Println(bh.Back(7)) // "leetcode.com"

	bh2 := NewBrowserHistory("a.com")
	bh2.Visit("b.com")
	fmt.Println(bh2.Back(1)) // "a.com"
	fmt.Println(bh2.Forward(1)) // "b.com"
}

func NewBrowserHistory(homepage string) BrowserHistory {
	return BrowserHistory{
		history: []string{homepage},
		current: 0,
	}
}

// Time: O(1)
func (this *BrowserHistory) Visit(url string) {
	this.current++
	this.history = this.history[:this.current]
	this.history = append(this.history, url)
}

// Time: O(1)
func (this *BrowserHistory) Back(steps int) string {
	this.current -= steps
	if this.current < 0 {
		this.current = 0
	}
	return this.history[this.current]
}

// Time: O(1)
func (this *BrowserHistory) Forward(steps int) string {
	this.current += steps
	if this.current >= len(this.history) {
		this.current = len(this.history) - 1
	}
	return this.history[this.current]
}
```

## 1476 — Subrectangle Queries

```go
package main

// LeetCode #1476: Subrectangle Queries
// https://leetcode.com/problems/subrectangle-queries/
// Difficulty: Medium

import "fmt"

type SubrectangleQueries struct {
	rectangle   [][]int
	updates     [][5]int // row1, col1, row2, col2, newValue (lazy)
}

func main() {
	sq := NewSubrectangleQueries([][]int{{1, 2, 1}, {4, 3, 4}, {3, 2, 1}, {1, 1, 1}})
	fmt.Println(sq.GetValue(0, 2)) // 1
	sq.UpdateSubrectangle(0, 0, 3, 2, 5)
	fmt.Println(sq.GetValue(0, 2)) // 5
	fmt.Println(sq.GetValue(3, 1)) // 5
	sq.UpdateSubrectangle(3, 0, 3, 2, 10)
	fmt.Println(sq.GetValue(3, 1)) // 10
	fmt.Println(sq.GetValue(0, 2)) // 5

	sq2 := NewSubrectangleQueries([][]int{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}})
	fmt.Println(sq2.GetValue(0, 0)) // 1
	sq2.UpdateSubrectangle(0, 0, 2, 2, 100)
	fmt.Println(sq2.GetValue(0, 0)) // 100
	fmt.Println(sq2.GetValue(2, 2)) // 100
	sq2.UpdateSubrectangle(1, 1, 2, 2, 200)
	fmt.Println(sq2.GetValue(0, 0)) // 100
	fmt.Println(sq2.GetValue(1, 1)) // 200
}

func NewSubrectangleQueries(rectangle [][]int) SubrectangleQueries {
	return SubrectangleQueries{rectangle: rectangle}
}

// Time: O(1)
func (this *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int) {
	this.updates = append(this.updates, [5]int{row1, col1, row2, col2, newValue})
}

// Time: O(k) where k = number of updates
func (this *SubrectangleQueries) GetValue(row int, col int) int {
	// Check most recent update first
	for i := len(this.updates) - 1; i >= 0; i-- {
		u := this.updates[i]
		if row >= u[0] && row <= u[2] && col >= u[1] && col <= u[3] {
			return u[4]
		}
	}
	return this.rectangle[row][col]
}
```

## 1477 — Find Two Non Overlapping Sub Arrays Each With Target Sum

```go
package main

// LeetCode #1477: Find Two Non-overlapping Sub-arrays Each With Target Sum
// https://leetcode.com/problems/find-two-non-overlapping-sub-arrays-each-with-target-sum/
// Difficulty: Medium

import "fmt"
import "math"

func main() {
	// Test case 1
	fmt.Println(minSumOfLengths([]int{3, 2, 2, 4, 3}, 3)) // 2

	// Test case 2
	fmt.Println(minSumOfLengths([]int{7, 3, 4, 7}, 7)) // 2

	// Test case 3
	fmt.Println(minSumOfLengths([]int{4, 3, 2, 6, 2, 3, 4}, 6)) // -1

	// Test case 4
	fmt.Println(minSumOfLengths([]int{1, 1, 1, 2, 1, 1}, 3)) // 3
}

// Time: O(n) where n = len(arr)
// Space: O(n) for prefix minimum array
func minSumOfLengths(arr []int, target int) int {
	n := len(arr)
	// left[i] = minimum length of subarray with sum = target ending at or before i
	left := make([]int, n)
	for i := range left {
		left[i] = math.MaxInt32
	}

	prefixSum := 0
	sumMap := make(map[int]int)
	sumMap[0] = -1
	bestLeft := math.MaxInt32

	for i := 0; i < n; i++ {
		prefixSum += arr[i]
		if j, ok := sumMap[prefixSum-target]; ok {
			length := i - j
			if length < bestLeft {
				bestLeft = length
			}
			if i > 0 && left[i-1] < bestLeft {
				bestLeft = left[i-1]
			}
		}
		left[i] = bestLeft
		sumMap[prefixSum] = i
	}

	// right[i] = minimum length of subarray with sum = target starting at or after i
	right := make([]int, n)
	for i := range right {
		right[i] = math.MaxInt32
	}

	suffixSum := 0
	sumMap = make(map[int]int)
	sumMap[0] = n
	bestRight := math.MaxInt32

	for i := n - 1; i >= 0; i-- {
		suffixSum += arr[i]
		if j, ok := sumMap[suffixSum-target]; ok {
			length := j - i
			if length < bestRight {
				bestRight = length
			}
			if i < n-1 && right[i+1] < bestRight {
				bestRight = right[i+1]
			}
		}
		right[i] = bestRight
		sumMap[suffixSum] = i
	}

	result := math.MaxInt32
	for i := 0; i < n-1; i++ {
		if left[i] != math.MaxInt32 && right[i+1] != math.MaxInt32 {
			total := left[i] + right[i+1]
			if total < result {
				result = total
			}
		}
	}

	if result == math.MaxInt32 {
		return -1
	}
	return result
}
```

## 1481 — Least Number Of Unique Integers After K Removals

```go
package main

// LeetCode #1481: Least Number of Unique Integers after K Removals
// https://leetcode.com/problems/least-number-of-unique-integers-after-k-removals/
// Difficulty: Medium

import "fmt"
import "sort"

func main() {
	// Test case 1
	fmt.Println(findLeastNumOfUniqueInts([]int{5, 5, 4}, 1)) // 1

	// Test case 2
	fmt.Println(findLeastNumOfUniqueInts([]int{4, 3, 1, 1, 3, 3, 2}, 3)) // 2

	// Test case 3
	fmt.Println(findLeastNumOfUniqueInts([]int{1, 2, 3, 4, 5}, 5)) // 0
}

// Time: O(n log n) for sorting frequencies
// Space: O(n) for frequency map
func findLeastNumOfUniqueInts(arr []int, k int) int {
	freq := make(map[int]int)
	for _, v := range arr {
		freq[v]++
	}

	counts := make([]int, 0, len(freq))
	for _, c := range freq {
		counts = append(counts, c)
	}

	sort.Ints(counts)

	remaining := k
	uniqueCount := len(counts)

	for _, c := range counts {
		if remaining >= c {
			remaining -= c
			uniqueCount--
		} else {
			break
		}
	}

	return uniqueCount
}
```

## 1482 — Minimum Number Of Days To Make M Bouquets

```go
package main

// LeetCode #1482: Minimum Number of Days to Make m Bouquets
// https://leetcode.com/problems/minimum-number-of-days-to-make-m-bouquets/
// Difficulty: Medium

import "fmt"
import "math"

func main() {
	// Test case 1
	fmt.Println(minDays([]int{1, 10, 3, 10, 2}, 3, 1)) // 3

	// Test case 2
	fmt.Println(minDays([]int{1, 10, 3, 10, 2}, 3, 2)) // -1

	// Test case 3
	fmt.Println(minDays([]int{7, 7, 7, 7, 12, 7, 7}, 2, 3)) // 12

	// Test case 4
	fmt.Println(minDays([]int{1, 10, 2, 9, 3, 8, 4, 7, 5, 6}, 4, 2)) // 9
}

// Time: O(n * log max(bloomDay)) for binary search
// Space: O(1)
func minDays(bloomDay []int, m int, k int) int {
	if m*k > len(bloomDay) {
		return -1
	}

	minDay, maxDay := math.MaxInt32, 0
	for _, d := range bloomDay {
		if d < minDay {
			minDay = d
		}
		if d > maxDay {
			maxDay = d
		}
	}

	canMake := func(day int) bool {
		bouquets := 0
		consecutive := 0
		for _, d := range bloomDay {
			if d <= day {
				consecutive++
				if consecutive == k {
					bouquets++
					consecutive = 0
				}
			} else {
				consecutive = 0
			}
		}
		return bouquets >= m
	}

	// Binary search
	left, right := minDay, maxDay
	for left < right {
		mid := left + (right-left)/2
		if canMake(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}
```

## 1485 — Clone Binary Tree With Random Pointer

```go
package main

// LeetCode #1485: Clone Binary Tree With Random Pointer
// https://leetcode.com/problems/clone-binary-tree-with-random-pointer/
// Difficulty: Medium

import "fmt"

type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	Random *Node
}

type NodeCopy struct {
	Val    int
	Left   *NodeCopy
	Right  *NodeCopy
	Random *NodeCopy
}

func main() {
	// Test case 1
	root := &Node{Val: 1}
	root.Left = &Node{Val: 2}
	root.Right = &Node{Val: 3}
	root.Left.Random = root.Right
	root.Right.Random = root.Left

	copied := copyRandomBinaryTree(root)
	fmt.Println(copied.Val) // 1
	fmt.Println(copied.Left.Val) // 2
	fmt.Println(copied.Right.Val) // 3
	fmt.Println(copied.Left.Random.Val) // 3
	fmt.Println(copied.Right.Random.Val) // 2

	// Test case 2 - nil
	fmt.Println(copyRandomBinaryTree(nil)) // nil

	// Test case 3 - single node
	root3 := &Node{Val: 42}
	copied3 := copyRandomBinaryTree(root3)
	fmt.Println(copied3.Val) // 42
}

// Time: O(n) where n = number of nodes
// Space: O(n) for the map
func copyRandomBinaryTree(root *Node) *NodeCopy {
	if root == nil {
		return nil
	}

	// Map from original node to copy
	nodeMap := make(map[*Node]*NodeCopy)

	var dfs func(*Node) *NodeCopy
	dfs = func(node *Node) *NodeCopy {
		if node == nil {
			return nil
		}
		if copy, ok := nodeMap[node]; ok {
			return copy
		}
		copy := &NodeCopy{Val: node.Val}
		nodeMap[node] = copy
		copy.Left = dfs(node.Left)
		copy.Right = dfs(node.Right)
		return copy
	}

	rootCopy := dfs(root)

	// Set random pointers
	var setRandom func(*Node, *NodeCopy)
	setRandom = func(orig *Node, copy *NodeCopy) {
		if orig == nil || copy == nil {
			return
		}
		if orig.Random != nil {
			copy.Random = nodeMap[orig.Random]
		}
		setRandom(orig.Left, copy.Left)
		setRandom(orig.Right, copy.Right)
	}
	setRandom(root, rootCopy)

	return rootCopy
}
```

## 1487 — Making File Names Unique

```go
package main

// LeetCode #1487: Making File Names Unique
// https://leetcode.com/problems/making-file-names-unique/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(GetFolderNames([]string{"pes", "fifa", "gta", "pes(2019)"}))
	fmt.Println(GetFolderNames([]string{"gta", "gta(1)", "gta", "avalon"}))
	fmt.Println(GetFolderNames([]string{"onepiece", "onepiece(1)", "onepiece(2)", "onepiece", "onepiece(1)"}))
}

func GetFolderNames(names []string) []string {
	// Time: O(N) average, Space: O(N)
	used := make(map[string]int)
	result := make([]string, len(names))

	for i, name := range names {
		if _, exists := used[name]; !exists {
			used[name] = 1
			result[i] = name
			continue
		}

		k := used[name]
		candidate := name + "(" + itoa(k) + ")"
		for {
			if _, exists := used[candidate]; exists {
				k++
				candidate = name + "(" + itoa(k) + ")"
			} else {
				break
			}
		}
		used[name] = k + 1
		used[candidate] = 1
		result[i] = candidate
	}

	return result
}

// Simple int to string for positive ints
func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	var buf [12]byte
	pos := len(buf)
	for n > 0 {
		pos--
		buf[pos] = byte('0' + n%10)
		n /= 10
	}
	return string(buf[pos:])
}
```

## 1488 — Avoid Flood In The City

```go
package main

// LeetCode #1488: Avoid Flood in The City
// https://leetcode.com/problems/avoid-flood-in-the-city/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(AvoidFlood([]int{1, 2, 3, 4}))
	fmt.Println(AvoidFlood([]int{1, 2, 0, 0, 2, 1}))
	fmt.Println(AvoidFlood([]int{1, 2, 0, 1, 2}))
}

func AvoidFlood(rains []int) []int {
	// Time: O(N log N), Space: O(N)
	n := len(rains)
	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1 // default for rain days
	}

	lastRain := make(map[int]int) // lake -> last rain day
	dryDays := make([]int, 0)     // indices of dry days (0s)

	for i, lake := range rains {
		if lake == 0 {
			dryDays = append(dryDays, i)
			ans[i] = 1 // placeholder
			continue
		}

		ans[i] = -1 // rain day, no action

		if prev, exists := lastRain[lake]; exists {
			// Find a dry day after prev to dry this lake
			idx := sort.Search(len(dryDays), func(j int) bool {
				return dryDays[j] > prev
			})
			if idx == len(dryDays) {
				return nil // impossible to prevent flood
			}
			ans[dryDays[idx]] = lake
			// Remove used dry day
			dryDays = append(dryDays[:idx], dryDays[idx+1:]...)
		}
		lastRain[lake] = i
	}

	// Remaining dry days can be any positive number
	for _, idx := range dryDays {
		ans[idx] = 1
	}

	return ans
}
```

## 1490 — Clone N Ary Tree

```go
package main

// LeetCode #1490: Clone N-ary Tree
// https://leetcode.com/problems/clone-n-ary-tree/
// Difficulty: Medium [Paid]

import "fmt"

// Node is an N-ary tree node.
type Node struct {
	Val      int
	Children []*Node
}

func main() {
	// Build tree: root = 1 -> [3, 2, 4]; 3 -> [5, 6]
	root := &Node{Val: 1}
	child3 := &Node{Val: 3}
	child2 := &Node{Val: 2}
	child4 := &Node{Val: 4}
	root.Children = []*Node{child3, child2, child4}
	child3.Children = []*Node{{Val: 5}, {Val: 6}}

	cloned := CloneTree(root)
	fmt.Println("Root cloned:", cloned != nil && cloned != root)
	fmt.Println("Root val:", cloned.Val)
	fmt.Println("Children count:", len(cloned.Children))
	fmt.Println("Deep cloned:", cloned.Children[0].Children[0].Val == 5 && cloned.Children[0] != child3.Children[0])

	// Test nil
	fmt.Println(CloneTree(nil))
}

func CloneTree(root *Node) *Node {
	// Time: O(N), Space: O(N) (recursion stack)
	if root == nil {
		return nil
	}

	clone := &Node{Val: root.Val}
	clone.Children = make([]*Node, len(root.Children))
	for i, child := range root.Children {
		clone.Children[i] = CloneTree(child)
	}
	return clone
}
```

## 1492 — The Kth Factor Of N

```go
package main

// LeetCode #1492: The kth Factor of n
// https://leetcode.com/problems/the-kth-factor-of-n/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(KthFactor(12, 3))
	fmt.Println(KthFactor(7, 2))
	fmt.Println(KthFactor(4, 4))
}

func KthFactor(n int, k int) int {
	// Time: O(sqrt(N)), Space: O(1)
	// Count factors from 1 to sqrt(n)
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			k--
			if k == 0 {
				return i
			}
		}
	}

	// Count factors from sqrt(n) down to 1 (the paired factors)
	// Start from the largest paired factor
	for i := intSqrt(n); i >= 1; i-- {
		if n%i == 0 && i*i != n { // don't double count perfect square root
			k--
			if k == 0 {
				return n / i
			}
		}
	}

	return -1
}

func intSqrt(n int) int {
	for i := 1; i*i <= n; i++ {
		if i*i == n {
			return i
		}
	}
	// floor sqrt
	result := 0
	for result*result <= n {
		result++
	}
	return result - 1
}
```

## 1493 — Longest Subarray Of 1s After Deleting One Element

```go
package main

// LeetCode #1493: Longest Subarray of 1's After Deleting One Element
// https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(LongestSubarray([]int{1, 1, 0, 1}))
	fmt.Println(LongestSubarray([]int{0, 1, 1, 1, 0, 1, 1, 0, 1}))
	fmt.Println(LongestSubarray([]int{1, 1, 1}))
}

func LongestSubarray(nums []int) int {
	// Time: O(N), Space: O(1)
	// Sliding window with at most one zero
	left := 0
	zeroCount := 0
	maxLen := 0

	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zeroCount++
		}

		for zeroCount > 1 {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}

		// Window length minus the one element we must delete
		currLen := right - left
		if currLen > maxLen {
			maxLen = currLen
		}
	}

	return maxLen
}
```

## 1497 — Check If Array Pairs Are Divisible By K

```go
package main

// LeetCode #1497: Check If Array Pairs Are Divisible by k
// https://leetcode.com/problems/check-if-array-pairs-are-divisible-by-k/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(CanArrange([]int{1, 2, 3, 4, 5, 10, 6, 7, 8, 9}, 5))
	fmt.Println(CanArrange([]int{1, 2, 3, 4, 5, 6}, 7))
	fmt.Println(CanArrange([]int{1, 2, 3, 4, 5, 6}, 10))
}

func CanArrange(arr []int, k int) bool {
	// Time: O(N), Space: O(K)
	remainder := make([]int, k)
	for _, num := range arr {
		r := ((num % k) + k) % k
		remainder[r]++
	}

	// Numbers divisible by k must pair among themselves
	if remainder[0]%2 != 0 {
		return false
	}

	// For i and k-i, their counts must match
	for i := 1; i < k; i++ {
		if remainder[i] != remainder[k-i] {
			return false
		}
	}

	return true
}
```

## 1498 — Number Of Subsequences That Satisfy The Given Sum Condition

```go
package main

// LeetCode #1498: Number of Subsequences That Satisfy the Given Sum Condition
// https://leetcode.com/problems/number-of-subsequences-that-satisfy-the-given-sum-condition/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(NumSubseq([]int{3, 5, 6, 7}, 9))
	fmt.Println(NumSubseq([]int{3, 3, 6, 8}, 10))
	fmt.Println(NumSubseq([]int{2, 3, 3, 4, 6, 7}, 12))
}

func NumSubseq(nums []int, target int) int {
	// Time: O(N log N), Space: O(N)
	const mod = 1_000_000_007

	// Sort nums
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	quickSort(sorted, 0, len(sorted)-1)

	// Precompute powers of 2
	pow := make([]int, len(sorted))
	pow[0] = 1
	for i := 1; i < len(sorted); i++ {
		pow[i] = (pow[i-1] * 2) % mod
	}

	count := 0
	left, right := 0, len(sorted)-1

	for left <= right {
		if sorted[left]+sorted[right] <= target {
			// All subsequences with sorted[left] as min and any subset of elements between left+1..right
			count = (count + pow[right-left]) % mod
			left++
		} else {
			right--
		}
	}

	return count
}

func quickSort(arr []int, low, high int) {
	if low < high {
		p := partition(arr, low, high)
		quickSort(arr, low, p-1)
		quickSort(arr, p+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
```

## 1500 — Design A File Sharing System

```go
package main

// LeetCode #1500: Design a File Sharing System
// https://leetcode.com/problems/design-a-file-sharing-system/
// Difficulty: Medium [Paid]

import "fmt"

// FileSharing simulates a file sharing system.
type FileSharing struct {
	chunks   map[int][]int   // user -> list of chunk IDs
	nextID   int             // next user ID
	released []int           // reusable user IDs
}

func Constructor(m int) FileSharing {
	return FileSharing{
		chunks:   make(map[int][]int),
		nextID:   1,
		released: make([]int, 0),
	}
}

func (fs *FileSharing) Join(ownedChunks []int) int {
	var userID int
	if len(fs.released) > 0 {
		// Reuse the smallest released ID
		userID = fs.released[0]
		fs.released = fs.released[1:]
	} else {
		userID = fs.nextID
		fs.nextID++
	}

	// Store a copy of chunks
	chunks := make([]int, len(ownedChunks))
	copy(chunks, ownedChunks)
	fs.chunks[userID] = chunks
	return userID
}

func (fs *FileSharing) Leave(userID int) {
	delete(fs.chunks, userID)
	// Insert sorted
	idx := 0
	for idx < len(fs.released) && fs.released[idx] < userID {
		idx++
	}
	fs.released = append(fs.released, 0)
	copy(fs.released[idx+1:], fs.released[idx:])
	fs.released[idx] = userID
}

func (fs *FileSharing) Request(userID int, chunkID int) []int {
	owners := make([]int, 0)
	for uid, chunks := range fs.chunks {
		if uid != userID {
			for _, c := range chunks {
				if c == chunkID {
					owners = append(owners, uid)
					break
				}
			}
		}
	}

	// Sort owners
	sortInts(owners)

	// Requesting user gets the chunk
	fs.chunks[userID] = append(fs.chunks[userID], chunkID)

	return owners
}

func sortInts(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

func main() {
	fs := Constructor(4)
	user1 := fs.Join([]int{1, 2})
	fmt.Println("User1 ID:", user1)
	user2 := fs.Join([]int{2, 3})
	fmt.Println("User2 ID:", user2)
	user3 := fs.Join([]int{4})
	fmt.Println("User3 ID:", user3)

	owners := fs.Request(user2, 1)
	fmt.Println("Request chunk 1 owners:", owners)
	owners = fs.Request(user2, 4)
	fmt.Println("Request chunk 4 owners:", owners)

	fs.Leave(user1)
	user4 := fs.Join([]int{1})
	fmt.Println("User4 ID (reused):", user4)
}
```

## 1501 — Countries You Can Safely Invest In

```go
package main

// LeetCode #1501: Countries You Can Safely Invest In
// https://leetcode.com/problems/countries-you-can-safely-invest-in/
// Difficulty: Medium [Paid]

import "fmt"

func main() {
	// SQL problem translated to Go: find countries whose average call duration
	// exceeds the global average.
	// Tables: Person(id, name, phone_number), Country(code, name), Calls(caller_id, callee_id, duration)

	// person ID -> country code
	personCountry := map[int]string{1: "US", 2: "US", 3: "UK", 4: "UK", 5: "IN"}
	// country code -> name
	countryName := map[string]string{"US": "USA", "UK": "UK", "IN": "India"}
	// calls: {caller_id, callee_id, duration}
	calls := [][3]int{{1, 2, 10}, {2, 1, 20}, {3, 4, 5}, {4, 3, 15}, {5, 1, 30}}

	result := FindSafeCountries(personCountry, countryName, calls)
	fmt.Println("Safe countries:", result)
}

func FindSafeCountries(personCountry map[int]string, countryName map[string]string, calls [][3]int) []string {
	// Time: O(N), Space: O(K) where N = calls, K = countries
	countryDur := make(map[string]int)
	countryCount := make(map[string]int)
	globalDur := 0
	globalCount := 0

	for _, c := range calls {
		caller, callee, dur := c[0], c[1], c[2]
		globalDur += dur
		globalCount++

		callerCode := personCountry[caller]
		calleeCode := personCountry[callee]

		countryDur[callerCode] += dur
		countryCount[callerCode]++
		if callerCode != calleeCode {
			countryDur[calleeCode] += dur
			countryCount[calleeCode]++
		}
	}

	if globalCount == 0 {
		return nil
	}
	globalAvg := float64(globalDur) / float64(globalCount)

	result := make([]string, 0)
	for code, dur := range countryDur {
		avg := float64(dur) / float64(countryCount[code])
		if avg > globalAvg {
			if name, ok := countryName[code]; ok {
				result = append(result, name)
			} else {
				result = append(result, code)
			}
		}
	}
	return result
}
```

## 1503 — Last Moment Before All Ants Fall Out Of A Plank

```go
package main

// LeetCode #1503: Last Moment Before All Ants Fall Out of a Plank
// https://leetcode.com/problems/last-moment-before-all-ants-fall-out-of-a-plank/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(GetLastMoment(4, []int{4, 3}, []int{0, 1}))
	fmt.Println(GetLastMoment(7, []int{}, []int{0, 1, 2, 3, 4, 5, 6, 7}))
	fmt.Println(GetLastMoment(7, []int{0, 1, 2, 3, 4, 5, 6, 7}, []int{}))
}

func GetLastMoment(n int, left []int, right []int) int {
	// Time: O(N), Space: O(1)
	// When ants meet, they reverse direction. This is equivalent to
	// ants passing through each other (identity swap).
	// So the last moment is max of:
	//   - ants moving left: their starting position (time to reach 0)
	//   - ants moving right: n - their starting position (time to reach n)
	maxTime := 0

	for _, pos := range left {
		if pos > maxTime {
			maxTime = pos
		}
	}

	for _, pos := range right {
		t := n - pos
		if t > maxTime {
			maxTime = t
		}
	}

	return maxTime
}
```

## 1504 — Count Submatrices With All Ones

```go
package main

// LeetCode #1504: Count Submatrices With All Ones
// https://leetcode.com/problems/count-submatrices-with-all-ones/
// Difficulty: Medium

import "fmt"

func main() {
	mat1 := [][]int{{1, 0, 1}, {1, 1, 0}, {1, 1, 0}}
	fmt.Println(NumSubmat(mat1))

	mat2 := [][]int{{0, 1, 1, 0}, {0, 1, 1, 1}, {1, 1, 1, 0}}
	fmt.Println(NumSubmat(mat2))

	fmt.Println(NumSubmat([][]int{{1, 1}, {1, 1}}))
}

func NumSubmat(mat [][]int) int {
	// Time: O(R*C), Space: O(C)
	if len(mat) == 0 || len(mat[0]) == 0 {
		return 0
	}
	rows, cols := len(mat), len(mat[0])
	height := make([]int, cols)
	total := 0

	for r := 0; r < rows; r++ {
		// Update height of consecutive 1s in each column
		for c := 0; c < cols; c++ {
			if mat[r][c] == 1 {
				height[c]++
			} else {
				height[c] = 0
			}
		}

		// Count submatrices ending at row r using monotonic stack
		stack := make([]int, 0)
		sum := make([]int, cols)

		for c := 0; c < cols; c++ {
			// Pop from stack while height[c] <= height[stack top]
			for len(stack) > 0 && height[stack[len(stack)-1]] >= height[c] {
				stack = stack[:len(stack)-1]
			}

			if len(stack) > 0 {
				prev := stack[len(stack)-1]
				sum[c] = sum[prev] + height[c]*(c-prev)
			} else {
				sum[c] = height[c] * (c + 1)
			}

			stack = append(stack, c)
			total += sum[c]
		}
	}

	return total
}
```

## 1506 — Find Root Of N Ary Tree

```go
package main

// LeetCode #1506: Find Root of N-Ary Tree
// https://leetcode.com/problems/find-root-of-n-ary-tree/
// Difficulty: Medium [Paid]

import "fmt"

// Node is an N-ary tree node.
type Node struct {
	Val      int
	Children []*Node
}

func main() {
	// Build tree: root = 1 -> [2, 3, 4]; 3 -> [5, 6]
	child5 := &Node{Val: 5}
	child6 := &Node{Val: 6}
	child2 := &Node{Val: 2}
	child3 := &Node{Val: 3, Children: []*Node{child5, child6}}
	child4 := &Node{Val: 4}
	root := &Node{Val: 1, Children: []*Node{child2, child3, child4}}

	// All nodes in random order (without knowing root)
	allNodes := []*Node{child2, child5, child3, child4, root, child6}

	found := FindRoot(allNodes)
	fmt.Println("Found root val:", found.Val)
}

func FindRoot(tree []*Node) *Node {
	// Time: O(N), Space: O(1)
	// The root is the only node that is never a child.
	// XOR all node values + all child values. Root value remains.
	var xorSum int
	for _, node := range tree {
		xorSum ^= node.Val
		for _, child := range node.Children {
			xorSum ^= child.Val
		}
	}

	// Find node with matching value
	for _, node := range tree {
		if node.Val == xorSum {
			return node
		}
	}
	return nil
}
```

## 1508 — Range Sum Of Sorted Subarray Sums

```go
package main

// LeetCode #1508: Range Sum of Sorted Subarray Sums
// https://leetcode.com/problems/range-sum-of-sorted-subarray-sums/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(RangeSum([]int{1, 2, 3, 4}, 4, 1, 5))
	fmt.Println(RangeSum([]int{1, 2, 3, 4}, 4, 3, 4))
	fmt.Println(RangeSum([]int{1, 2, 3, 4}, 4, 1, 10))
}

func RangeSum(nums []int, n int, left int, right int) int {
	// Time: O(N^2 log N), Space: O(N^2)
	const mod = 1_000_000_007

	// Generate all subarray sums
	sums := make([]int, 0, n*(n+1)/2)
	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += nums[j]
			sums = append(sums, sum)
		}
	}

	// Sort
	sort.Ints(sums)

	// Sum from left-1 to right-1
	result := 0
	for i := left - 1; i < right; i++ {
		result = (result + sums[i]) % mod
	}

	return result
}
```

## 1509 — Minimum Difference Between Largest And Smallest Value In Three Moves

```go
package main

// LeetCode #1509: Minimum Difference Between Largest and Smallest Value in Three Moves
// https://leetcode.com/problems/minimum-difference-between-largest-and-smallest-value-in-three-moves/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MinDifference([]int{5, 3, 2, 4}))
	fmt.Println(MinDifference([]int{1, 5, 0, 10, 14}))
	fmt.Println(MinDifference([]int{3, 100, 20}))
}

func MinDifference(nums []int) int {
	// Time: O(N log N), Space: O(1) if ignoring sort space
	if len(nums) <= 4 {
		return 0
	}

	sort.Ints(nums)
	n := len(nums)

	// After 3 moves, we can change up to 3 values.
	// The minimum difference will be between some combination
	// of removing 0-3 from left and 3-0 from right.
	minDiff := nums[n-1] - nums[0]
	for i := 0; i <= 3; i++ {
		diff := nums[n-1-(3-i)] - nums[i]
		if diff < minDiff {
			minDiff = diff
		}
	}

	return minDiff
}
```

## 1513 — Number Of Substrings With Only 1s

```go
package main

// LeetCode #1513: Number of Substrings With Only 1s
// https://leetcode.com/problems/number-of-substrings-with-only-1s/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(NumSub("0110111"))
	fmt.Println(NumSub("101"))
	fmt.Println(NumSub("111111"))
}

func NumSub(s string) int {
	// Time: O(N), Space: O(1)
	const mod = 1_000_000_007

	count := 0
	consecutive := 0

	for _, ch := range s {
		if ch == '1' {
			consecutive++
			// Each new 1 adds 'consecutive' new substrings ending at this position
			count = (count + consecutive) % mod
		} else {
			consecutive = 0
		}
	}

	return count
}
```

## 1514 — Path With Maximum Probability

```go
package main

// LeetCode #1514: Path with Maximum Probability
// https://leetcode.com/problems/path-with-maximum-probability/
// Difficulty: Medium

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(MaxProbability(3, [][]int{{0, 1}, {1, 2}, {0, 2}}, []float64{0.5, 0.5, 0.2}, 0, 2))
	fmt.Println(MaxProbability(3, [][]int{{0, 1}, {1, 2}, {0, 2}}, []float64{0.5, 0.5, 0.3}, 0, 2))
	fmt.Println(MaxProbability(3, [][]int{{0, 1}}, []float64{0.5}, 0, 2))
}

type Edge struct {
	to   int
	prob float64
}

type Item struct {
	node   int
	prob   float64
	index  int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i].prob > pq[j].prob } // max-heap
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
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
	*pq = old[:n-1]
	return item
}

func MaxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	// Time: O(E log V), Space: O(E + V)
	// Build adjacency list
	graph := make([][]Edge, n)
	for i, e := range edges {
		u, v := e[0], e[1]
		p := succProb[i]
		graph[u] = append(graph[u], Edge{v, p})
		graph[v] = append(graph[v], Edge{u, p})
	}

	// Dijkstra-like (max probability)
	prob := make([]float64, n)
	prob[start] = 1.0

	pq := &PriorityQueue{}
	heap.Push(pq, &Item{node: start, prob: 1.0})

	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)
		node := item.node
		curProb := item.prob

		if node == end {
			return curProb
		}

		if curProb < prob[node] {
			continue
		}

		for _, edge := range graph[node] {
			newProb := curProb * edge.prob
			if newProb > prob[edge.to] {
				prob[edge.to] = newProb
				heap.Push(pq, &Item{node: edge.to, prob: newProb})
			}
		}
	}

	return 0.0
}
```

## 1519 — Number Of Nodes In The Sub Tree With The Same Label

```go
package main

// LeetCode #1519: Number of Nodes in the Sub-Tree With the Same Label
// https://leetcode.com/problems/number-of-nodes-in-the-sub-tree-with-the-same-label/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(CountSubTrees(7, [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}}, "abaedcd"))
	fmt.Println(CountSubTrees(4, [][]int{{0, 1}, {1, 2}, {0, 3}}, "bbbb"))
	fmt.Println(CountSubTrees(5, [][]int{{0, 1}, {0, 2}, {1, 3}, {0, 4}}, "aabab"))
}

func CountSubTrees(n int, edges [][]int, labels string) []int {
	// Time: O(N), Space: O(N)
	graph := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	result := make([]int, n)
	visited := make([]bool, n)

	var dfs func(node int) []int
	dfs = func(node int) []int {
		visited[node] = true
		// Count array for 26 lowercase letters
		count := make([]int, 26)
		count[labels[node]-'a'] = 1

		for _, nei := range graph[node] {
			if visited[nei] {
				continue
			}
			childCount := dfs(nei)
			for i := 0; i < 26; i++ {
				count[i] += childCount[i]
			}
		}

		result[node] = count[labels[node]-'a']
		return count
	}

	dfs(0)
	return result
}
```

## 1522 — Diameter Of N Ary Tree

```go
package main

// LeetCode #1522: Diameter of N-Ary Tree
// https://leetcode.com/problems/diameter-of-n-ary-tree/
// Difficulty: Medium [Paid]

import "fmt"

// Node is an N-ary tree node.
type Node struct {
	Val      int
	Children []*Node
}

func main() {
	// Tree: 1 -> [2, 3, 4]; 3 -> [5, 6]
	child5 := &Node{Val: 5}
	child6 := &Node{Val: 6}
	child2 := &Node{Val: 2}
	child3 := &Node{Val: 3, Children: []*Node{child5, child6}}
	child4 := &Node{Val: 4}
	root := &Node{Val: 1, Children: []*Node{child2, child3, child4}}
	fmt.Println(Diameter(root))

	// Single node
	fmt.Println(Diameter(&Node{Val: 1}))

	// Linear chain: 1 -> 2 -> 3
	n3 := &Node{Val: 3}
	n2 := &Node{Val: 2, Children: []*Node{n3}}
	n1 := &Node{Val: 1, Children: []*Node{n2}}
	fmt.Println(Diameter(n1))
}

func Diameter(root *Node) int {
	// Time: O(N), Space: O(H) where H = height
	maxDiameter := 0

	var dfs func(node *Node) int
	dfs = func(node *Node) int {
		if node == nil {
			return 0
		}
		// Track top two deepest paths from children
		firstMax, secondMax := 0, 0

		for _, child := range node.Children {
			depth := dfs(child)
			if depth > firstMax {
				secondMax = firstMax
				firstMax = depth
			} else if depth > secondMax {
				secondMax = depth
			}
		}

		// Diameter through this node = sum of two deepest child paths
		if firstMax+secondMax > maxDiameter {
			maxDiameter = firstMax + secondMax
		}

		// Return max depth from this node
		return firstMax + 1
	}

	dfs(root)
	return maxDiameter
}
```

## 1524 — Number Of Sub Arrays With Odd Sum

```go
package main

// LeetCode #1524: Number of Sub-arrays With Odd Sum
// https://leetcode.com/problems/number-of-sub-arrays-with-odd-sum/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(NumOfSubarrays([]int{1, 3, 5}))
	fmt.Println(NumOfSubarrays([]int{2, 4, 6}))
	fmt.Println(NumOfSubarrays([]int{1, 2, 3, 4, 5, 6, 7}))
}

func NumOfSubarrays(arr []int) int {
	// Time: O(N), Space: O(1)
	const mod = 1_000_000_007

	oddCount := 0
	evenCount := 1 // prefix sum = 0 is even
	prefixSum := 0
	result := 0

	for _, num := range arr {
		prefixSum += num

		if prefixSum%2 == 0 {
			// Current prefix is even
			result = (result + oddCount) % mod
			evenCount++
		} else {
			// Current prefix is odd
			result = (result + evenCount) % mod
			oddCount++
		}
	}

	return result
}
```

## 1525 — Number Of Good Ways To Split A String

```go
package main

// LeetCode #1525: Number of Good Ways to Split a String
// https://leetcode.com/problems/number-of-good-ways-to-split-a-string/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(NumSplits("aacaba"))
	fmt.Println(NumSplits("abcd"))
	fmt.Println(NumSplits("aaaaa"))
}

func NumSplits(s string) int {
	// Time: O(N), Space: O(1) (26 chars)
	n := len(s)
	leftCount := make([]int, 26)
	rightCount := make([]int, 26)
	leftUnique := 0
	rightUnique := 0

	// Initialize right side
	for i := 0; i < n; i++ {
		idx := s[i] - 'a'
		if rightCount[idx] == 0 {
			rightUnique++
		}
		rightCount[idx]++
	}

	result := 0
	for i := 0; i < n-1; i++ {
		idx := s[i] - 'a'
		// Move char from right to left
		if leftCount[idx] == 0 {
			leftUnique++
		}
		leftCount[idx]++

		rightCount[idx]--
		if rightCount[idx] == 0 {
			rightUnique--
		}

		if leftUnique == rightUnique {
			result++
		}
	}

	return result
}
```

## 1529 — Minimum Suffix Flips

```go
package main

// LeetCode #1529: Minimum Suffix Flips
// https://leetcode.com/problems/minimum-suffix-flips/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MinFlips("10111"))
	fmt.Println(MinFlips("101"))
	fmt.Println(MinFlips("00000"))
}

func MinFlips(target string) int {
	// Time: O(N), Space: O(1)
	// Count transitions from 0 to 1 or 1 to 0
	flips := 0
	curr := byte('0') // current state of flipped prefix

	for i := 0; i < len(target); i++ {
		if target[i] != curr {
			flips++
			curr = target[i]
		}
	}

	return flips
}
```

## 1530 — Number Of Good Leaf Nodes Pairs

```go
package main

// LeetCode #1530: Number of Good Leaf Nodes Pairs
// https://leetcode.com/problems/number-of-good-leaf-nodes-pairs/
// Difficulty: Medium

import "fmt"

// TreeNode is a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Tree: [1,2,3,null,4]
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2, Right: &TreeNode{Val: 4}}
	root.Right = &TreeNode{Val: 3}
	fmt.Println(CountPairs(root, 3))

	// Tree: [1,2,3,4,5,6,7]
	root2 := &TreeNode{Val: 1}
	root2.Left = &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}
	root2.Right = &TreeNode{Val: 3, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 7}}
	fmt.Println(CountPairs(root2, 3))

	// Single node
	fmt.Println(CountPairs(&TreeNode{Val: 1}, 1))
}

func CountPairs(root *TreeNode, distance int) int {
	// Time: O(N * distance^2), Space: O(N * distance)
	pairs := 0

	var dfs func(node *TreeNode) []int
	dfs = func(node *TreeNode) []int {
		if node == nil {
			return nil
		}
		if node.Left == nil && node.Right == nil {
			return []int{1} // leaf, distance 1 from here
		}

		left := dfs(node.Left)
		right := dfs(node.Right)

		if left == nil && right == nil {
			return nil
		}

		// Count pairs between left and right subtrees
		if left != nil && right != nil {
			for _, ld := range left {
				for _, rd := range right {
					if ld+rd <= distance {
						pairs++
					}
				}
			}
		}

		// Merge distances, incrementing by 1 (edge to parent)
		result := make([]int, 0)
		if left != nil {
			for _, d := range left {
				if d+1 < distance {
					result = append(result, d+1)
				}
			}
		}
		if right != nil {
			for _, d := range right {
				if d+1 < distance {
					result = append(result, d+1)
				}
			}
		}

		return result
	}

	dfs(root)
	return pairs
}
```

## 1532 — The Most Recent Three Orders

```go
package main

// LeetCode #1532: The Most Recent Three Orders
// https://leetcode.com/problems/the-most-recent-three-orders/
// Difficulty: Medium [Paid]

import (
	"fmt"
	"sort"
)

func main() {
	// SQL problem: For each customer, find their 3 most recent orders.
	// Translated to Go.
	// Tables: Customers(customer_id, name), Orders(order_id, order_date, customer_id, cost)

	orders := []struct{ id, customerID int; date string; cost float64 }{
		{1, 1, "2020-07-31", 30.0},
		{2, 1, "2020-07-30", 40.0},
		{3, 1, "2020-07-29", 20.0},
		{4, 1, "2020-07-28", 50.0},
		{5, 2, "2020-07-31", 10.0},
		{6, 2, "2020-07-30", 15.0},
		{7, 3, "2020-07-31", 25.0},
	}

	result := RecentThreeOrders(orders)
	fmt.Println("Recent 3 orders per customer:")
	for _, r := range result {
		fmt.Printf("  Customer %d: Order %d on %s ($%.2f)\n", r.customerID, r.orderID, r.date, r.cost)
	}
}

type orderRec struct {
	customerID int
	orderID    int
	date       string
	cost       float64
}

func RecentThreeOrders(orders []struct{ id, customerID int; date string; cost float64 }) []orderRec {
	// Group orders by customer
	customerOrders := make(map[int][]struct{ id int; date string; cost float64 })
	for _, o := range orders {
		customerOrders[o.customerID] = append(customerOrders[o.customerID], struct{ id int; date string; cost float64 }{o.id, o.date, o.cost})
	}

	result := make([]orderRec, 0)
	for cid, ords := range customerOrders {
		// Sort by date descending
		sort.Slice(ords, func(i, j int) bool {
			return ords[i].date > ords[j].date
		})
		// Take top 3
		for i := 0; i < 3 && i < len(ords); i++ {
			result = append(result, orderRec{cid, ords[i].id, ords[i].date, ords[i].cost})
		}
	}

	return result
}
```

## 1533 — Find The Index Of The Large Integer

```go
package main

// LeetCode #1533: Find the Index of the Large Integer
// https://leetcode.com/problems/find-the-index-of-the-large-integer/
// Difficulty: Medium [Paid]

import "fmt"

func main() {
	// Problem: There is an array where one element is larger than the rest
	// (which are all equal). Find the index of the larger element.
	// We have an ArrayReader API that returns:
	//   -1 if sum(arr[l..r]) < sum(arr[subL..subR])
	//   0 if equal
	//   1 if sum(arr[l..r]) > sum(arr[subL..subR])

	arr := []int{1, 1, 1, 1, 1, 10, 1, 1, 1}
	reader := &arrayReader{arr: arr}
	idx := getIndex(reader, len(arr))
	fmt.Println("Index of larger element:", idx, "value:", arr[idx])

	// All equal
	arr2 := []int{2, 2, 2, 2, 2}
	reader2 := &arrayReader{arr: arr2}
	idx2 := getIndex(reader2, len(arr2))
	fmt.Println("Index (all equal):", idx2)
}

type arrayReader struct {
	arr []int
}

func (ar *arrayReader) compareSub(l, r, subL, subR int) int {
	s1 := sumRange(ar.arr, l, r)
	s2 := sumRange(ar.arr, subL, subR)
	if s1 < s2 {
		return -1
	} else if s1 > s2 {
		return 1
	}
	return 0
}

func sumRange(arr []int, l, r int) int {
	sum := 0
	for i := l; i <= r; i++ {
		sum += arr[i]
	}
	return sum
}

// getIndex uses the compareSub API to find the index of the largest element.
func getIndex(reader interface{ compareSub(int, int, int, int) int }, n int) int {
	// Time: O(log N), Space: O(1)
	left, right := 0, n-1

	for left < right {
		mid := left + (right-left)/2
		len1 := mid - left + 1
		len2 := right - mid

		if len1 == len2 {
			result := reader.compareSub(left, mid, mid+1, right)
			if result == 1 {
				right = mid
			} else if result == -1 {
				left = mid + 1
			} else {
				return -1 // all equal
			}
		} else {
			// len1 > len2 (since mid is floor)
			result := reader.compareSub(left, mid-1, mid+1, right)
			if result == 0 {
				return mid // the extra element is the larger one
			} else if result == 1 {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}

	return left
}
```

## 1535 — Find The Winner Of An Array Game

```go
package main

// LeetCode #1535: Find the Winner of an Array Game
// https://leetcode.com/problems/find-the-winner-of-an-array-game/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(GetWinner([]int{2, 1, 3, 5, 4, 6, 7}, 2))
	fmt.Println(GetWinner([]int{3, 2, 1}, 10))
	fmt.Println(GetWinner([]int{1, 11, 22, 33, 44, 55, 66, 77, 88, 99}, 1000000000))
}

func GetWinner(arr []int, k int) int {
	// Time: O(N), Space: O(1)
	if k == 0 {
		return 0
	}

	// If k >= n, the maximum element wins
	n := len(arr)
	maxVal := arr[0]
	for _, v := range arr {
		if v > maxVal {
			maxVal = v
		}
	}
	if k >= n {
		return maxVal
	}

	current := arr[0]
	wins := 0

	for i := 1; i < n; i++ {
		if current > arr[i] {
			wins++
		} else {
			current = arr[i]
			wins = 1
		}

		if wins == k {
			return current
		}
	}

	// If we've gone through the whole array, the max element wins
	return maxVal
}
```

## 1536 — Minimum Swaps To Arrange A Binary Grid

```go
package main

// LeetCode #1536: Minimum Swaps to Arrange a Binary Grid
// https://leetcode.com/problems/minimum-swaps-to-arrange-a-binary-grid/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MinSwaps([][]int{{0, 0, 1}, {1, 1, 0}, {1, 0, 0}}))
	fmt.Println(MinSwaps([][]int{{0, 1, 1, 0}, {0, 1, 1, 0}, {0, 1, 1, 0}, {0, 1, 1, 0}}))
	fmt.Println(MinSwaps([][]int{{1, 0, 0}, {1, 1, 0}, {1, 1, 1}}))
}

func MinSwaps(grid [][]int) int {
	// Time: O(N^2), Space: O(N)
	n := len(grid)

	// trailingZeros[i] = number of trailing zeros in row i
	trailingZeros := make([]int, n)
	for i := 0; i < n; i++ {
		count := 0
		for j := n - 1; j >= 0 && grid[i][j] == 0; j-- {
			count++
		}
		trailingZeros[i] = count
	}

	swaps := 0

	for i := 0; i < n; i++ {
		// Row i needs at least n-i-1 trailing zeros
		needed := n - i - 1
		found := -1

		for j := i; j < n; j++ {
			if trailingZeros[j] >= needed {
				found = j
				break
			}
		}

		if found == -1 {
			return -1
		}

		// Bubble the found row up to position i
		for j := found; j > i; j-- {
			trailingZeros[j], trailingZeros[j-1] = trailingZeros[j-1], trailingZeros[j]
			swaps++
		}
	}

	return swaps
}
```

## 1538 — Guess The Majority In A Hidden Array

```go
package main

// LeetCode #1538: Guess the Majority in a Hidden Array
// https://leetcode.com/problems/guess-the-majority-in-a-hidden-array/
// Difficulty: Medium [Paid]

import "fmt"

func main() {
	// Problem: There is an array of 0s and 1s. Use a query API that returns
	// whether the majority of 4 indices are same (0 or 1) or not.
	// Find the index of an element that is different from the majority,
	// or return -1 if all elements are the same.
	//
	// We simulate with a known array.

	arr := []int{0, 0, 1, 0, 0}
	idx := getMajorityIndex(len(arr), func(a, b, c, d int) bool {
		count0, count1 := 0, 0
		for _, i := range []int{a, b, c, d} {
			if arr[i] == 0 {
				count0++
			} else {
				count1++
			}
		}
		return count0 > count1 // returns true if majority is 0
	})
	fmt.Println("Different element index:", idx) // should be 2

	arr2 := []int{1, 1, 1, 1, 1}
	idx2 := getMajorityIndex(len(arr2), func(a, b, c, d int) bool {
		count0, count1 := 0, 0
		for _, i := range []int{a, b, c, d} {
			if arr2[i] == 0 {
				count0++
			} else {
				count1++
			}
		}
		return count0 > count1
	})
	fmt.Println("All same:", idx2) // should be -1
}

func getMajorityIndex(n int, query func(int, int, int, int) bool) int {
	// Time: O(N), Space: O(1)
	if n < 4 {
		return -1
	}

	// Compare 0,1,2,3 with 0,1,2,4 to see if 3 and 4 differ
	// If query(0,1,2,3) == query(0,1,2,4), then 3 and 4 are same
	// Otherwise they differ

	// Find two indices that differ
	diff := -1
	for i := 1; i < n; i++ {
		if query(0, 1, 2, i) != query(0, 1, 2, 0) {
			diff = i
			break
		}
	}

	if diff == -1 {
		return -1 // all elements are the same
	}

	// The differing element is either at position 0 or at diff
	// Check if 0,1,2,3 all agree
	if query(0, 1, 2, 0) == query(1, 0, 2, 0) {
		// They agree, so the majority is the common value, and
		// the minority is at diff
		return diff
	}
	return 0
}
```

## 1540 — Can Convert String In K Moves

```go
package main

// LeetCode #1540: Can Convert String in K Moves
// https://leetcode.com/problems/can-convert-string-in-k-moves/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(CanConvertString("input", "ouput", 9))
	fmt.Println(CanConvertString("abc", "bcd", 10))
	fmt.Println(CanConvertString("aab", "bbb", 27))
}

func CanConvertString(s string, t string, k int) bool {
	// Time: O(N), Space: O(1)
	if len(s) != len(t) {
		return false
	}

	// Count how many times each shift value is needed
	shiftCount := make([]int, 26)

	for i := 0; i < len(s); i++ {
		if s[i] == t[i] {
			continue
		}
		// Compute needed shift (positive modulo 26)
		shift := (int(t[i]) - int(s[i]) + 26) % 26
		if shift == 0 {
			continue
		}

		// For each subsequent time we need this shift, add 26
		shiftCount[shift]++
		needed := shift + (shiftCount[shift]-1)*26
		if needed > k {
			return false
		}
	}

	return true
}
```

## 1541 — Minimum Insertions To Balance A Parentheses String

```go
package main

// LeetCode #1541: Minimum Insertions to Balance a Parentheses String
// https://leetcode.com/problems/minimum-insertions-to-balance-a-parentheses-string/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MinInsertions("(()))"))
	fmt.Println(MinInsertions("())"))
	fmt.Println(MinInsertions("))())("))
}

func MinInsertions(s string) int {
	// Time: O(N), Space: O(1)
	// Each '(' needs two ')' to balance.
	insertions := 0
	open := 0 // number of '(' that need closing

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			open++
		} else { // ')'
			if open > 0 {
				// Check if next char is also ')'
				if i+1 < len(s) && s[i+1] == ')' {
					// Both ')' found, consume both
					i++ // skip next ')'
				} else {
					// Need one more ')'
					insertions++
				}
				open--
			} else {
				// Need a '(' before this ')'
				if i+1 < len(s) && s[i+1] == ')' {
					// Insert '(' and consume both ')'
					insertions++ // for '('
					i++          // consume next ')'
				} else {
					// Insert '(' and one ')'
					insertions += 2
				}
			}
		}
	}

	// Each remaining '(' needs two ')'
	insertions += open * 2

	return insertions
}
```

## 1545 — Find Kth Bit In Nth Binary String

```go
package main

// LeetCode #1545: Find Kth Bit in Nth Binary String
// https://leetcode.com/problems/find-kth-bit-in-nth-binary-string/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(FindKthBit(3, 1))
	fmt.Println(FindKthBit(4, 11))
	fmt.Println(FindKthBit(1, 1))
}

func FindKthBit(n int, k int) byte {
	// Time: O(N), Space: O(N)
	// S1 = "0"
	// Si = Si-1 + "1" + reverse(invert(Si-1))
	// The length of Sn is 2^n - 1
	// We can recursively find the bit without constructing the string

	if n == 1 {
		return '0'
	}

	mid := 1 << (n - 1) // 2^(n-1), the middle position (1-indexed)

	if k == mid {
		return '1'
	} else if k < mid {
		return FindKthBit(n-1, k)
	} else {
		// k > mid: mirror position
		mirrorK := mid*2 - k
		bit := FindKthBit(n-1, mirrorK)
		// invert: '0' <-> '1'
		if bit == '0' {
			return '1'
		}
		return '0'
	}
}
```

## 1546 — Maximum Number Of Non Overlapping Subarrays With Sum Equals Target

```go
package main

// LeetCode #1546: Maximum Number of Non-Overlapping Subarrays With Sum Equals Target
// https://leetcode.com/problems/maximum-number-of-non-overlapping-subarrays-with-sum-equals-target/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MaxNonOverlapping([]int{1, 1, 1, 1, 1}, 2))
	fmt.Println(MaxNonOverlapping([]int{-1, 3, 5, 1, 4, 2, -9}, 6))
	fmt.Println(MaxNonOverlapping([]int{-2, 6, 6, 3, 5, 4, 1, 2, 8}, 10))
}

func MaxNonOverlapping(nums []int, target int) int {
	// Time: O(N), Space: O(N)
	// Greedy: use prefix sums map to find earliest non-overlapping subarray
	prefixSum := 0
	seen := make(map[int]int)
	seen[0] = -1 // prefix sum of empty array
	result := 0
	lastEnd := -1 // last used subarray end index

	for i, num := range nums {
		prefixSum += num
		if prevEnd, exists := seen[prefixSum-target]; exists && prevEnd >= lastEnd {
			result++
			lastEnd = i
		}
		seen[prefixSum] = i
	}

	return result
}
```

## 1549 — The Most Recent Orders For Each Product

```go
package main

// LeetCode #1549: The Most Recent Orders for Each Product
// https://leetcode.com/problems/the-most-recent-orders-for-each-product/
// Difficulty: Medium [Paid]

import (
	"fmt"
	"sort"
)

func main() {
	// SQL problem: for each product, find the most recent order(s).
	// Tables: Products(product_id, product_name), Orders(order_id, product_id, order_date)

	products := map[int]string{
		1: "Product A",
		2: "Product B",
		3: "Product C",
	}

	orders := []struct{ orderID, productID int; date string }{
		{1, 1, "2020-07-31"},
		{2, 1, "2020-07-30"},
		{3, 1, "2020-07-29"},
		{4, 2, "2020-07-31"},
		{5, 2, "2020-07-30"},
		{6, 3, "2020-07-31"},
	}

	result := MostRecentOrders(products, orders)
	fmt.Println("Most recent orders per product:")
	for _, r := range result {
		fmt.Printf("  %s: Order %d on %s\n", r.productName, r.orderID, r.date)
	}
}

type recentOrderInfo struct {
	productName string
	orderID     int
	date        string
}

func MostRecentOrders(products map[int]string, orders []struct{ orderID, productID int; date string }) []recentOrderInfo {
	// Group orders by product
	productOrders := make(map[int][]struct{ orderID int; date string })
	for _, o := range orders {
		productOrders[o.productID] = append(productOrders[o.productID], struct{ orderID int; date string }{o.orderID, o.date})
	}

	// Find most recent order date per product
	productRecent := make(map[int]string)
	for pid, ords := range productOrders {
		sort.Slice(ords, func(i, j int) bool {
			return ords[i].date > ords[j].date
		})
		productRecent[pid] = ords[0].date
	}

	// Collect orders that match the most recent date for their product
	result := make([]recentOrderInfo, 0)
	for pid, pname := range products {
		recentDate, ok := productRecent[pid]
		if !ok {
			continue
		}
		for _, o := range orders {
			if o.productID == pid && o.date == recentDate {
				result = append(result, recentOrderInfo{pname, o.orderID, o.date})
			}
		}
	}

	return result
}
```

## 1551 — Minimum Operations To Make Array Equal

```go
package main

// LeetCode #1551: Minimum Operations to Make Array Equal
// https://leetcode.com/problems/minimum-operations-to-make-array-equal/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MinOperations(3))
	fmt.Println(MinOperations(6))
	fmt.Println(MinOperations(1))
}

func MinOperations(n int) int {
	// Time: O(1), Space: O(1)
	// arr[i] = 2*i + 1 for i in [0..n-1]
	// Target value = n (average of first n odd numbers)
	// Total operations = sum of (target - arr[i]) for arr[i] < target
	// = sum of (n - (2*i+1)) for i where 2*i+1 < n
	// = sum of (n - 2*i - 1) for i < n/2
	//
	// For n even: n=2k, sum_{i=0}^{k-1} (2k-2i-1) = k^2
	// For n odd: n=2k+1, sum_{i=0}^{k-1} (2k+1-2i-1) = k*(k+1) = k^2 + k = k(k+1)

	return n * n / 4
}
```

## 1552 — Magnetic Force Between Two Balls

```go
package main

// LeetCode #1552: Magnetic Force Between Two Balls
// https://leetcode.com/problems/magnetic-force-between-two-balls/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MaxDistance([]int{1, 2, 3, 4, 7}, 3))
	fmt.Println(MaxDistance([]int{5, 4, 3, 2, 1, 1000000000}, 2))
	fmt.Println(MaxDistance([]int{1, 2, 3, 4, 5, 6}, 3))
}

func MaxDistance(position []int, m int) int {
	// Time: O(N log N + N log MAX), Space: O(1)
	sort.Ints(position)

	low, high := 1, position[len(position)-1]-position[0]
	result := 0

	for low <= high {
		mid := low + (high-low)/2
		if canPlace(position, m, mid) {
			result = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return result
}

func canPlace(position []int, m int, minDist int) bool {
	count := 1
	lastPos := position[0]

	for i := 1; i < len(position); i++ {
		if position[i]-lastPos >= minDist {
			count++
			lastPos = position[i]
			if count >= m {
				return true
			}
		}
	}

	return false
}
```

## 1554 — Strings Differ By One Character

```go
package main

// LeetCode #1554: Strings Differ by One Character
// https://leetcode.com/problems/strings-differ-by-one-character/
// Difficulty: Medium [Paid]

import "fmt"

func main() {
	fmt.Println(DifferByOne([]string{"abcd", "acbd", "aacd"}))
	fmt.Println(DifferByOne([]string{"ab", "cd", "yz"}))
	fmt.Println(DifferByOne([]string{"abcd", "cccc", "abxd", "abzd"}))
}

func DifferByOne(dict []string) bool {
	// Time: O(N*M^2) where N = len(dict), M = string length
	// Space: O(N*M)
	// Use rolling hash: for each position, check if any two strings
	// become identical when that position is skipped.

	n := len(dict)
	if n < 2 {
		return false
	}
	m := len(dict[0])

	for skipIdx := 0; skipIdx < m; skipIdx++ {
		seen := make(map[string]bool)
		for i := 0; i < n; i++ {
			// Create string without char at skipIdx
			key := dict[i][:skipIdx] + dict[i][skipIdx+1:]
			if seen[key] {
				return true
			}
			seen[key] = true
		}
	}

	return false
}
```

## 1555 — Bank Account Summary

```go
package main

// LeetCode #1555: Bank Account Summary
// https://leetcode.com/problems/bank-account-summary/
// Difficulty: Medium [Paid]

import "fmt"

func main() {
	// SQL problem: find users with balance > 10000 or overdrawn users.
	// Translated to Go.
	// Tables: Users(user_id, user_name, credit), Transactions(trans_id, user_id, amount, transacted_on)

	users := map[int]string{1: "Alice", 2: "Bob", 3: "Charlie"}
	credits := map[int]int{1: 5000, 2: 8000, 3: 3000}
	txns := []struct{ userID, amount int }{
		{1, 2000}, {2, -3000}, {3, 10000}, {1, -1000},
	}

	result := AccountSummary(users, credits, txns)
	fmt.Println("Account summary:")
	for _, r := range result {
		fmt.Printf("  %s: balance=%d, status=%s\n", r.name, r.balance, r.status)
	}
}

type accountInfo struct {
	userID  int
	name    string
	balance int
	status  string
}

func AccountSummary(users map[int]string, credits map[int]int, txns []struct{ userID, amount int }) []accountInfo {
	balance := make(map[int]int)
	for uid, credit := range credits {
		balance[uid] = credit
	}
	for _, t := range txns {
		balance[t.userID] += t.amount
	}

	result := make([]accountInfo, 0)
	for uid, name := range users {
		b := balance[uid]
		status := "Normal"
		if b < 0 {
			status = "Overdrawn"
		}
		result = append(result, accountInfo{uid, name, b, status})
	}

	return result
}
```

## 1557 — Minimum Number Of Vertices To Reach All Nodes

```go
package main

// LeetCode #1557: Minimum Number of Vertices to Reach All Nodes
// https://leetcode.com/problems/minimum-number-of-vertices-to-reach-all-nodes/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(FindSmallestSetOfVertices(6, [][]int{{0, 1}, {0, 2}, {2, 5}, {3, 4}, {4, 2}}))
	fmt.Println(FindSmallestSetOfVertices(3, [][]int{{0, 1}, {2, 1}}))
	fmt.Println(FindSmallestSetOfVertices(5, [][]int{{0, 1}, {2, 1}, {3, 1}, {4, 0}}))
}

func FindSmallestSetOfVertices(n int, edges [][]int) []int {
	// Time: O(N + E), Space: O(N)
	// Nodes with indegree 0 must be in the result since they can't be reached
	indegree := make([]int, n)
	for _, e := range edges {
		indegree[e[1]]++
	}

	result := make([]int, 0)
	for i := 0; i < n; i++ {
		if indegree[i] == 0 {
			result = append(result, i)
		}
	}

	return result
}
```

## 1558 — Minimum Numbers Of Function Calls To Make Target Array

```go
package main

// LeetCode #1558: Minimum Numbers of Function Calls to Make Target Array
// https://leetcode.com/problems/minimum-numbers-of-function-calls-to-make-target-array/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MinOperations([]int{1, 5}))
	fmt.Println(MinOperations([]int{2, 2}))
	fmt.Println(MinOperations([]int{4, 2, 5}))
}

func MinOperations(nums []int) int {
	// Time: O(N), Space: O(1)
	// Operations: (1) increment one element, (2) double all elements
	// Count total increments (set bits) + maximum number of doublings (highest bit position)
	increments := 0
	maxDoublings := 0

	for _, num := range nums {
		// Count bits (increment operations)
		bits := 0
		pos := 0
		for n := num; n > 0; n >>= 1 {
			if n&1 == 1 {
				bits++
			}
			pos++
		}
		increments += bits
		if pos-1 > maxDoublings {
			maxDoublings = pos - 1
		}
	}

	return increments + maxDoublings
}
```

## 1559 — Detect Cycles In 2d Grid

```go
package main

// LeetCode #1559: Detect Cycles in 2D Grid
// https://leetcode.com/problems/detect-cycles-in-2d-grid/
// Difficulty: Medium

import "fmt"

func main() {
	grid1 := [][]byte{{'a', 'a', 'a', 'a'}, {'a', 'b', 'b', 'a'}, {'a', 'b', 'b', 'a'}, {'a', 'a', 'a', 'a'}}
	fmt.Println(ContainsCycle(grid1))

	grid2 := [][]byte{{'c', 'c', 'c', 'a'}, {'c', 'd', 'c', 'c'}, {'c', 'c', 'e', 'c'}, {'f', 'c', 'c', 'c'}}
	fmt.Println(ContainsCycle(grid2))

	grid3 := [][]byte{{'a', 'b'}, {'b', 'a'}}
	fmt.Println(ContainsCycle(grid3))
}

func ContainsCycle(grid [][]byte) bool {
	// Time: O(R*C), Space: O(R*C)
	if len(grid) == 0 || len(grid[0]) == 0 {
		return false
	}

	rows, cols := len(grid), len(grid[0])
	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, cols)
	}

	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	var dfs func(r, c, pr, pc int) bool
	dfs = func(r, c, pr, pc int) bool {
		visited[r][c] = true

		for _, d := range dirs {
			nr, nc := r+d[0], c+d[1]
			if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
				continue
			}
			if nr == pr && nc == pc {
				continue
			}
			if grid[nr][nc] != grid[r][c] {
				continue
			}
			if visited[nr][nc] {
				return true // cycle found
			}
			if dfs(nr, nc, r, c) {
				return true
			}
		}

		return false
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if !visited[r][c] {
				if dfs(r, c, -1, -1) {
					return true
				}
			}
		}
	}

	return false
}
```

## 1561 — Maximum Number Of Coins You Can Get

```go
package main

// LeetCode #1561: Maximum Number of Coins You Can Get
// https://leetcode.com/problems/maximum-number-of-coins-you-can-get/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MaxCoins([]int{2, 4, 1, 2, 7, 8}))
	fmt.Println(MaxCoins([]int{2, 4, 5}))
	fmt.Println(MaxCoins([]int{9, 8, 7, 6, 5, 1, 2, 3, 4}))
}

func MaxCoins(piles []int) int {
	// Time: O(N log N), Space: O(1)
	sort.Ints(piles)

	coins := 0
	n := len(piles)
	// Alice gets the largest n/3, you get the next n/3, Bob gets the smallest n/3
	// You take the second largest in each group of 3
	start := n / 3

	for i := start; i < n; i += 2 {
		coins += piles[i]
	}

	return coins
}
```

## 1562 — Find Latest Group Of Size M

```go
package main

// LeetCode #1562: Find Latest Group of Size M
// https://leetcode.com/problems/find-latest-group-of-size-m/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(FindLatestStep([]int{3, 5, 1, 2, 4}, 1))
	fmt.Println(FindLatestStep([]int{3, 1, 5, 4, 2}, 2))
	fmt.Println(FindLatestStep([]int{1}, 1))
}

func FindLatestStep(arr []int, m int) int {
	// Time: O(N), Space: O(N)
	// Use union-find-like approach tracking group lengths
	n := len(arr)
	if n == m {
		return n
	}

	length := make([]int, n+2) // length of group at position i
	count := make([]int, n+2)  // count of groups of length i
	result := -1

	for step := 0; step < n; step++ {
		pos := arr[step]
		leftLen := length[pos-1]
		rightLen := length[pos+1]
		total := leftLen + rightLen + 1

		// Decrement counts for the merging groups
		count[leftLen]--
		count[rightLen]--
		// Increment count for the new merged group
		count[total]++

		// Update lengths at boundaries
		length[pos-leftLen] = total
		length[pos+rightLen] = total

		// Check if we have exactly m groups of size m
		if count[m] > 0 {
			// This step is valid (but since we want the latest step before m disappears,
			// we'll track the result after the step)
		}

		// The result is the latest step where count[m] > 0
		// But we want the latest step where a group of size m *exists*
		// We need to check AFTER the current operation
	}

	// Re-simulate to find latest step with count[m] > 0
	length = make([]int, n+2)
	count = make([]int, n+2)

	for step := 0; step < n; step++ {
		pos := arr[step]
		leftLen := length[pos-1]
		rightLen := length[pos+1]
		total := leftLen + rightLen + 1

		count[leftLen]--
		count[rightLen]--
		count[total]++

		length[pos-leftLen] = total
		length[pos+rightLen] = total

		if count[m] > 0 {
			result = step + 1 // 1-indexed step
		}
	}

	return result
}
```

## 1564 — Put Boxes Into The Warehouse I

```go
package main

// LeetCode #1564: Put Boxes Into the Warehouse I
// https://leetcode.com/problems/put-boxes-into-the-warehouse-i/
// Difficulty: Medium [Paid]

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MaxBoxesInWarehouse([]int{4, 3, 4, 1}, []int{5, 3, 3, 4, 1}))
	fmt.Println(MaxBoxesInWarehouse([]int{1, 2, 3}, []int{3, 2, 1}))
	fmt.Println(MaxBoxesInWarehouse([]int{1, 2, 2, 3, 4}, []int{3, 4, 1, 2}))
}

func MaxBoxesInWarehouse(boxes []int, warehouse []int) int {
	// Time: O(N log N + M), Space: O(1)
	sort.Ints(boxes)

	// Preprocess warehouse: each position's max usable height
	// is min of itself and all previous positions' heights
	for i := 1; i < len(warehouse); i++ {
		if warehouse[i] > warehouse[i-1] {
			warehouse[i] = warehouse[i-1]
		}
	}

	boxIdx := 0
	// Try to fit boxes from the largest to smallest, entering from right
	for i := len(warehouse) - 1; i >= 0 && boxIdx < len(boxes); i-- {
		if boxes[boxIdx] <= warehouse[i] {
			boxIdx++
		}
	}

	return boxIdx
}
```

## 1567 — Maximum Length Of Subarray With Positive Product

```go
package main

// LeetCode #1567: Maximum Length of Subarray With Positive Product
// https://leetcode.com/problems/maximum-length-of-subarray-with-positive-product/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(GetMaxLen([]int{1, -2, -3, 4}))
	fmt.Println(GetMaxLen([]int{0, 1, -2, -3, -4}))
	fmt.Println(GetMaxLen([]int{-1, -2, -3, 0, 1}))
}

func GetMaxLen(nums []int) int {
	// Time: O(N), Space: O(1)
	// Track first occurrence of positive and negative prefix products
	maxLen := 0
	firstPos := -1
	firstNeg := -1
	prefix := 1 // 1 = positive, -1 = negative

	for i, num := range nums {
		if num > 0 {
			prefix = prefix // sign unchanged
		} else if num < 0 {
			prefix = -prefix
		} else {
			// Reset at zero
			prefix = 1
			firstPos = -1
			firstNeg = -1
			continue
		}

		if prefix == 1 {
			if firstPos == -1 {
				firstPos = i
			}
			if i-firstPos+1 > maxLen {
				maxLen = i - firstPos + 1
			}
		} else { // prefix == -1
			if firstNeg == -1 {
				firstNeg = i
			}
			if i-firstNeg+1 > maxLen {
				maxLen = i - firstNeg + 1
			}
		}
	}

	return maxLen
}
```

## 1570 — Dot Product Of Two Sparse Vectors

```go
package main

// LeetCode #1570: Dot Product of Two Sparse Vectors
// https://leetcode.com/problems/dot-product-of-two-sparse-vectors/
// Difficulty: Medium [Paid]

import "fmt"

func main() {
	v1 := ConstructorSparse([]int{1, 0, 0, 2, 3})
	v2 := ConstructorSparse([]int{0, 3, 0, 4, 0})

	dotProduct := v1.dotProduct(v2)
	fmt.Println("Dot product:", dotProduct)

	// Additional test
	v3 := ConstructorSparse([]int{0, 1, 0, 0, 0})
	v4 := ConstructorSparse([]int{0, 0, 0, 0, 2})
	fmt.Println("Dot product (no overlap):", v3.dotProduct(v4))
}

// SparseVector stores non-zero elements with their indices.
type SparseVector struct {
	pairs [][2]int // [index, value]
}

func ConstructorSparse(nums []int) SparseVector {
	pairs := make([][2]int, 0)
	for i, num := range nums {
		if num != 0 {
			pairs = append(pairs, [2]int{i, num})
		}
	}
	return SparseVector{pairs: pairs}
}

func (sv *SparseVector) dotProduct(vec SparseVector) int {
	// Time: O(N+M), Space: O(1)
	// Two pointer approach on sorted pairs
	result := 0
	i, j := 0, 0

	for i < len(sv.pairs) && j < len(vec.pairs) {
		if sv.pairs[i][0] == vec.pairs[j][0] {
			result += sv.pairs[i][1] * vec.pairs[j][1]
			i++
			j++
		} else if sv.pairs[i][0] < vec.pairs[j][0] {
			i++
		} else {
			j++
		}
	}

	return result
}
```

## 1573 — Number Of Ways To Split A String

```go
package main

// LeetCode #1573: Number of Ways to Split a String
// https://leetcode.com/problems/number-of-ways-to-split-a-string/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(NumWays("10101"))
	fmt.Println(NumWays("1001"))
	fmt.Println(NumWays("0000"))
}

func NumWays(s string) int {
	// Time: O(N), Space: O(1)
	const mod = 1_000_000_007

	// Count total ones
	totalOnes := 0
	for _, ch := range s {
		if ch == '1' {
			totalOnes++
		}
	}

	if totalOnes == 0 {
		// All zeros: need to choose 2 cut positions out of n-1 gaps
		// C(n-1, 2) = (n-1)*(n-2)/2
		n := len(s)
		return ((n - 1) * (n - 2) / 2) % mod
	}

	if totalOnes%3 != 0 {
		return 0
	}

	onesPerPart := totalOnes / 3
	count := 0
	firstWays := 0
	secondWays := 0

	for _, ch := range s {
		if ch == '1' {
			count++
		}

		if count == onesPerPart {
			firstWays++
		} else if count == 2*onesPerPart {
			secondWays++
		}
	}

	return (firstWays * secondWays) % mod
}
```

## 1574 — Shortest Subarray To Be Removed To Make Array Sorted

```go
package main

// LeetCode #1574: Shortest Subarray to be Removed to Make Array Sorted
// https://leetcode.com/problems/shortest-subarray-to-be-removed-to-make-array-sorted/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(FindLengthOfShortestSubarray([]int{1, 2, 3, 10, 4, 2, 3, 5}))
	fmt.Println(FindLengthOfShortestSubarray([]int{5, 4, 3, 2, 1}))
	fmt.Println(FindLengthOfShortestSubarray([]int{1, 2, 3}))
}

func FindLengthOfShortestSubarray(arr []int) int {
	// Time: O(N), Space: O(1)
	n := len(arr)

	// Find longest non-decreasing prefix
	left := 0
	for left < n-1 && arr[left] <= arr[left+1] {
		left++
	}

	if left == n-1 {
		return 0 // already sorted
	}

	// Find longest non-decreasing suffix
	right := n - 1
	for right > 0 && arr[right-1] <= arr[right] {
		right--
	}

	// Remove everything between left and right (minimum so far)
	result := minInt(n-left-1, right)

	// Try to merge prefix and suffix
	i, j := 0, right
	for i <= left && j < n {
		if arr[i] <= arr[j] {
			result = minInt(result, j-i-1)
			i++
		} else {
			j++
		}
	}

	return result
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```

## 1577 — Number Of Ways Where Square Of Number Is Equal To Product Of Two Numbers

```go
package main

// LeetCode #1577: Number of Ways Where Square of Number Is Equal to Product of Two Numbers
// https://leetcode.com/problems/number-of-ways-where-square-of-number-is-equal-to-product-of-two-numbers/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(NumTriplets([]int{7, 4}, []int{5, 2, 8, 9}))
	fmt.Println(NumTriplets([]int{1, 1}, []int{1, 1, 1}))
	fmt.Println(NumTriplets([]int{7, 7, 8, 3}, []int{1, 2, 9, 7}))
}

func NumTriplets(nums1 []int, nums2 []int) int {
	// Time: O(N^2 + M^2), Space: O(N^2)
	// Count pairs in each array that multiply to a specific product
	return countSquareProducts(nums1, nums2) + countSquareProducts(nums2, nums1)
}

func countSquareProducts(nums1 []int, nums2 []int) int {
	// Count nums1[i]^2 == nums2[j] * nums2[k] for j < k
	productCount := make(map[int]int)
	for j := 0; j < len(nums2); j++ {
		for k := j + 1; k < len(nums2); k++ {
			product := nums2[j] * nums2[k]
			productCount[product]++
		}
	}

	count := 0
	for _, v := range nums1 {
		square := v * v
		count += productCount[square]
	}

	return count
}
```

## 1578 — Minimum Time To Make Rope Colorful

```go
package main

// LeetCode #1578: Minimum Time to Make Rope Colorful
// https://leetcode.com/problems/minimum-time-to-make-rope-colorful/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MinCost("abaac", []int{1, 2, 3, 4, 5}))
	fmt.Println(MinCost("abc", []int{1, 2, 3}))
	fmt.Println(MinCost("aabaa", []int{1, 2, 3, 4, 1}))
}

func MinCost(colors string, neededTime []int) int {
	// Time: O(N), Space: O(1)
	n := len(colors)
	totalTime := 0

	i := 0
	for i < n {
		j := i
		maxTime := 0
		sum := 0
		for j < n && colors[j] == colors[i] {
			maxTime = maxInt(maxTime, neededTime[j])
			sum += neededTime[j]
			j++
		}
		// Keep the max time balloon, remove the rest
		totalTime += sum - maxTime
		i = j
	}

	return totalTime
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

## 1580 — Put Boxes Into The Warehouse Ii

```go
package main

// LeetCode #1580: Put Boxes Into the Warehouse II
// https://leetcode.com/problems/put-boxes-into-the-warehouse-ii/
// Difficulty: Medium [Paid]

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MaxBoxesInWarehouseII([]int{1, 2, 3}, []int{3, 2, 1}))
	fmt.Println(MaxBoxesInWarehouseII([]int{1, 2, 2, 3, 4}, []int{3, 4, 1, 2}))
	fmt.Println(MaxBoxesInWarehouseII([]int{4, 3, 4, 1}, []int{5, 3, 3, 4, 1}))
}

func MaxBoxesInWarehouseII(boxes []int, warehouse []int) int {
	// Time: O(N log N + M), Space: O(1)
	// In warehouse II, boxes can enter from either left or right side.
	// We can think of it as: each position's max height is the min of
	// the prefix max from left and prefix max from right.
	sort.Ints(boxes)

	n := len(warehouse)
	// Preprocess: effective height at each position
	leftMax := make([]int, n)
	rightMax := make([]int, n)

	leftMax[0] = warehouse[0]
	for i := 1; i < n; i++ {
		if warehouse[i] < leftMax[i-1] {
			leftMax[i] = warehouse[i]
		} else {
			leftMax[i] = leftMax[i-1]
		}
	}

	rightMax[n-1] = warehouse[n-1]
	for i := n - 2; i >= 0; i-- {
		if warehouse[i] < rightMax[i+1] {
			rightMax[i] = warehouse[i]
		} else {
			rightMax[i] = rightMax[i+1]
		}
	}

	// Effective height = max(leftMax, rightMax) since we can enter from either side
	effective := make([]int, n)
	for i := 0; i < n; i++ {
		if leftMax[i] > rightMax[i] {
			effective[i] = leftMax[i]
		} else {
			effective[i] = rightMax[i]
		}
	}

	// Greedily fit boxes
	boxIdx := 0
	for i := 0; i < n && boxIdx < len(boxes); i++ {
		if boxes[boxIdx] <= effective[i] {
			boxIdx++
		}
	}

	return boxIdx
}
```

## 1583 — Count Unhappy Friends

```go
package main

// LeetCode #1583: Count Unhappy Friends
// https://leetcode.com/problems/count-unhappy-friends/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(UnhappyFriends(4, [][]int{{1, 2, 3}, {3, 2, 0}, {3, 1, 0}, {1, 2, 0}}, [][]int{{0, 1}, {2, 3}}))
	fmt.Println(UnhappyFriends(2, [][]int{{1}, {0}}, [][]int{{0, 1}}))
	fmt.Println(UnhappyFriends(4, [][]int{{1, 3, 2}, {2, 3, 0}, {1, 3, 0}, {0, 2, 1}}, [][]int{{0, 1}, {2, 3}}))
}

func UnhappyFriends(n int, preferences [][]int, pairs [][]int) int {
	// Time: O(N^2), Space: O(N^2)
	// Build preference rank matrix: rank[i][j] = how much i prefers j
	rank := make([][]int, n)
	for i := 0; i < n; i++ {
		rank[i] = make([]int, n)
		for pos, j := range preferences[i] {
			rank[i][j] = pos
		}
	}

	// Map partner
	partner := make(map[int]int)
	for _, p := range pairs {
		partner[p[0]] = p[1]
		partner[p[1]] = p[0]
	}

	unhappy := 0

	for x := 0; x < n; x++ {
		y := partner[x]
		for _, u := range preferences[x] {
			if u == y {
				break
			}
			// x prefers u over y
			v := partner[u]
			// Is u unhappy? Check if u prefers x over v
			if rank[u][x] < rank[u][v] {
				unhappy++
				break
			}
		}
	}

	return unhappy
}
```

## 1584 — Min Cost To Connect All Points

```go
package main

// LeetCode #1584: Min Cost to Connect All Points
// https://leetcode.com/problems/min-cost-to-connect-all-points/
// Difficulty: Medium

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(MinCostConnectPoints([][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}}))
	fmt.Println(MinCostConnectPoints([][]int{{3, 12}, {-2, 5}, {-4, 1}}))
	fmt.Println(MinCostConnectPoints([][]int{{0, 0}, {1, 1}, {1, 0}, {-1, 1}}))
}

type Edge struct {
	to   int
	dist int
}

type PQItem struct {
	dist  int
	node  int
	index int
}

type PQ []*PQItem

func (pq PQ) Len() int           { return len(pq) }
func (pq PQ) Less(i, j int) bool { return pq[i].dist < pq[j].dist }
func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *PQ) Push(x interface{}) {
	n := len(*pq)
	item := x.(*PQItem)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[:n-1]
	return item
}

func MinCostConnectPoints(points [][]int) int {
	// Time: O(N^2 log N), Space: O(N)
	// Prim's algorithm
	n := len(points)
	if n <= 1 {
		return 0
	}

	visited := make([]bool, n)
	pq := &PQ{}
	heap.Init(pq)
	heap.Push(pq, &PQItem{dist: 0, node: 0})

	totalCost := 0
	edgesUsed := 0

	for pq.Len() > 0 && edgesUsed < n {
		item := heap.Pop(pq).(*PQItem)
		if visited[item.node] {
			continue
		}
		visited[item.node] = true
		totalCost += item.dist
		edgesUsed++

		for i := 0; i < n; i++ {
			if !visited[i] {
				dist := abs(points[item.node][0]-points[i][0]) + abs(points[item.node][1]-points[i][1])
				heap.Push(pq, &PQItem{dist: dist, node: i})
			}
		}
	}

	return totalCost
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

## 1586 — Binary Search Tree Iterator Ii

```go
package main

// LeetCode #1586: Binary Search Tree Iterator II
// https://leetcode.com/problems/binary-search-tree-iterator-ii/
// Difficulty: Medium [Paid]

import "fmt"

// TreeNode is a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Tree: [7, 3, 15, null, null, 9, 20]
	root := &TreeNode{Val: 7}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 15, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20}}

	it := ConstructorBST(root)
	fmt.Println("Next:", it.Next())      // 3
	fmt.Println("Next:", it.Next())      // 7
	fmt.Println("HasPrev:", it.HasPrev()) // true
	fmt.Println("Prev:", it.Prev())      // 3
	fmt.Println("Next:", it.Next())      // 7
	fmt.Println("Next:", it.Next())      // 9
	fmt.Println("Next:", it.Next())      // 15
	fmt.Println("HasNext:", it.HasNext()) // true
	fmt.Println("Next:", it.Next())      // 20
	fmt.Println("HasNext:", it.HasNext()) // false
}

type BSTIterator struct {
	stack []*TreeNode
	pos   int
	order []int
}

func ConstructorBST(root *TreeNode) BSTIterator {
	return BSTIterator{order: inorder(root)}
}

func inorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	result = append(result, inorder(root.Left)...)
	result = append(result, root.Val)
	result = append(result, inorder(root.Right)...)
	return result
}

func (it *BSTIterator) HasNext() bool {
	return it.pos < len(it.order)
}

func (it *BSTIterator) Next() int {
	val := it.order[it.pos]
	it.pos++
	return val
}

func (it *BSTIterator) HasPrev() bool {
	return it.pos > 1
}

func (it *BSTIterator) Prev() int {
	it.pos--
	return it.order[it.pos-1]
}
```

## 1589 — Maximum Sum Obtained Of Any Permutation

```go
package main

// LeetCode #1589: Maximum Sum Obtained of Any Permutation
// https://leetcode.com/problems/maximum-sum-obtained-of-any-permutation/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MaxSumRangeQuery([]int{1, 2, 3, 4, 5}, [][]int{{1, 3}, {0, 1}}))
	fmt.Println(MaxSumRangeQuery([]int{1, 2, 3, 4, 5, 6}, [][]int{{0, 1}}))
	fmt.Println(MaxSumRangeQuery([]int{1, 2, 3, 4, 5, 6}, [][]int{{0, 3}, {1, 5}}))
}

func MaxSumRangeQuery(nums []int, requests [][]int) int {
	// Time: O(N log N + M), Space: O(N)
	const mod = 1_000_000_007

	n := len(nums)
	// Difference array to count frequency of each index
	freq := make([]int, n+1)
	for _, req := range requests {
		freq[req[0]]++
		freq[req[1]+1]--
	}

	// Convert to actual frequency
	for i := 1; i < n; i++ {
		freq[i] += freq[i-1]
	}

	// Sort both nums and frequencies
	sort.Ints(nums)
	freqCounts := freq[:n]
	sort.Ints(freqCounts)

	// Assign largest numbers to most frequent positions
	result := 0
	for i := 0; i < n; i++ {
		result = (result + nums[i]*freqCounts[i]) % mod
	}

	return result
}
```

## 1590 — Make Sum Divisible By P

```go
package main

// LeetCode #1590: Make Sum Divisible by P
// https://leetcode.com/problems/make-sum-divisible-by-p/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MinSubarray([]int{3, 1, 4, 2}, 6))
	fmt.Println(MinSubarray([]int{6, 3, 5, 2}, 9))
	fmt.Println(MinSubarray([]int{1, 2, 3}, 7))
}

func MinSubarray(nums []int, p int) int {
	// Time: O(N), Space: O(N)
	n := len(nums)

	// Total sum modulo p
	totalSum := 0
	for _, num := range nums {
		totalSum = (totalSum + num) % p
	}

	target := totalSum // the remainder we need to remove
	if target == 0 {
		return 0
	}

	// Map from prefix sum modulo p to index
	prefixMap := make(map[int]int)
	prefixMap[0] = -1
	prefixSum := 0
	minLen := n

	for i, num := range nums {
		prefixSum = (prefixSum + num) % p
		// We need prefixSum - prefixSum[j] ≡ target (mod p)
		// => prefixSum[j] ≡ prefixSum - target (mod p)
		needed := (prefixSum - target + p) % p
		if j, exists := prefixMap[needed]; exists {
			if i-j < minLen {
				minLen = i - j
			}
		}
		prefixMap[prefixSum] = i
	}

	if minLen == n {
		return -1
	}
	return minLen
}
```

## 1593 — Split A String Into The Max Number Of Unique Substrings

```go
package main

// LeetCode #1593: Split a String Into the Max Number of Unique Substrings
// https://leetcode.com/problems/split-a-string-into-the-max-number-of-unique-substrings/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MaxUniqueSplit("ababccc"))
	fmt.Println(MaxUniqueSplit("aba"))
	fmt.Println(MaxUniqueSplit("aa"))
}

func MaxUniqueSplit(s string) int {
	// Time: O(2^N), Space: O(N)
	used := make(map[string]bool)
	maxCount := 0

	var backtrack func(start int, count int)
	backtrack = func(start int, count int) {
		if start == len(s) {
			if count > maxCount {
				maxCount = count
			}
			return
		}

		// Pruning: remaining chars <= max possible new substrings
		if count+(len(s)-start) <= maxCount {
			return
		}

		for end := start + 1; end <= len(s); end++ {
			sub := s[start:end]
			if !used[sub] {
				used[sub] = true
				backtrack(end, count+1)
				used[sub] = false
			}
		}
	}

	backtrack(0, 0)
	return maxCount
}
```

## 1594 — Maximum Non Negative Product In A Matrix

```go
package main

// LeetCode #1594: Maximum Non Negative Product in a Matrix
// https://leetcode.com/problems/maximum-non-negative-product-in-a-matrix/
// Difficulty: Medium

import "fmt"

func main() {
	grid1 := [][]int{{-1, -2, -3}, {-2, -3, -3}, {-3, -3, -2}}
	fmt.Println(MaxProductPath(grid1))

	grid2 := [][]int{{1, -2, 1}, {1, -2, 1}, {3, -4, 1}}
	fmt.Println(MaxProductPath(grid2))

	grid3 := [][]int{{1, 3}, {0, -4}}
	fmt.Println(MaxProductPath(grid3))
}

func MaxProductPath(grid [][]int) int {
	// Time: O(R*C), Space: O(C)
	const mod = 1_000_000_007

	rows, cols := len(grid), len(grid[0])
	if rows == 0 || cols == 0 {
		return -1
	}

	// minDP[r][c] = minimum product to reach (r,c)
	// maxDP[r][c] = maximum product to reach (r,c)
	minDP := make([][]int64, rows)
	maxDP := make([][]int64, rows)
	for i := 0; i < rows; i++ {
		minDP[i] = make([]int64, cols)
		maxDP[i] = make([]int64, cols)
	}

	maxDP[0][0] = int64(grid[0][0])
	minDP[0][0] = int64(grid[0][0])

	// First row
	for c := 1; c < cols; c++ {
		val := int64(grid[0][c])
		maxDP[0][c] = maxDP[0][c-1] * val
		minDP[0][c] = maxDP[0][c-1] * val
	}

	// First column
	for r := 1; r < rows; r++ {
		val := int64(grid[r][0])
		maxDP[r][0] = maxDP[r-1][0] * val
		minDP[r][0] = maxDP[r-1][0] * val
	}

	for r := 1; r < rows; r++ {
		for c := 1; c < cols; c++ {
			val := int64(grid[r][c])

			options := []int64{
				maxDP[r-1][c] * val,
				maxDP[r][c-1] * val,
				minDP[r-1][c] * val,
				minDP[r][c-1] * val,
			}

			maxVal := options[0]
			minVal := options[0]
			for _, opt := range options {
				if opt > maxVal {
					maxVal = opt
				}
				if opt < minVal {
					minVal = opt
				}
			}

			maxDP[r][c] = maxVal
			minDP[r][c] = minVal
		}
	}

	if maxDP[rows-1][cols-1] < 0 {
		return -1
	}

	return int(maxDP[rows-1][cols-1] % mod)
}
```

## 1596 — The Most Frequently Ordered Products For Each Customer

```go
package main

// LeetCode #1596: The Most Frequently Ordered Products for Each Customer
// https://leetcode.com/problems/the-most-frequently-ordered-products-for-each-customer/
// Difficulty: Medium [Paid]

import "fmt"

func main() {
	// SQL problem: for each customer, find their most frequently ordered product(s).

	// Using maps directly instead of named structs
	orderData := []struct{ customerID, productID int }{
		{1, 10}, {1, 10}, {1, 20},
		{2, 20}, {2, 20}, {2, 30},
		{3, 10},
	}

	products := map[int]string{10: "Product A", 20: "Product B", 30: "Product C"}
	customers := map[int]string{1: "Alice", 2: "Bob", 3: "Charlie"}

	result := MostFrequentProducts(orderData, products, customers)
	fmt.Println("Most frequently ordered products per customer:")
	for _, r := range result {
		fmt.Printf("  %s -> %s (%d orders)\n", r.customerName, r.productName, r.count)
	}
}

type freqProduct struct {
	customerName string
	productName  string
	count        int
}

func MostFrequentProducts(orders []struct{ customerID, productID int }, products, customers map[int]string) []freqProduct {
	customerCounts := make(map[int]map[int]int)

	for _, o := range orders {
		if customerCounts[o.customerID] == nil {
			customerCounts[o.customerID] = make(map[int]int)
		}
		customerCounts[o.customerID][o.productID]++
	}

	result := make([]freqProduct, 0)
	for cid, prodCounts := range customerCounts {
		maxCount := 0
		for _, count := range prodCounts {
			if count > maxCount {
				maxCount = count
			}
		}
		for pid, count := range prodCounts {
			if count == maxCount {
				result = append(result, freqProduct{
					customerName: customers[cid],
					productName:  products[pid],
					count:        count,
				})
			}
		}
	}

	return result
}
```

## 1599 — Maximum Profit Of Operating A Centennial Wheel

```go
package main

// LeetCode #1599: Maximum Profit of Operating a Centennial Wheel
// https://leetcode.com/problems/maximum-profit-of-operating-a-centennial-wheel/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MinOperationsMaxProfit([]int{8, 3}, 5, 6))
	fmt.Println(MinOperationsMaxProfit([]int{10, 9, 6}, 6, 4))
	fmt.Println(MinOperationsMaxProfit([]int{3, 4, 0, 5, 1}, 1, 92))
}

func MinOperationsMaxProfit(customers []int, boardingCost int, runningCost int) int {
	// Time: O(N), Space: O(1)
	// Each rotation can board up to 4 customers, costs runningCost, earns boardingCost per customer
	maxProfit := -1
	maxRotation := -1
	waiting := 0
	profit := 0
	rotation := 0

	for i := 0; i < len(customers) || waiting > 0; i++ {
		if i < len(customers) {
			waiting += customers[i]
		}

		// Board up to 4 customers
		boarded := 4
		if waiting < 4 {
			boarded = waiting
		}
		waiting -= boarded

		profit += boarded*boardingCost - runningCost
		rotation++

		if profit > maxProfit {
			maxProfit = profit
			maxRotation = rotation
		}
	}

	return maxRotation
}
```

## 1600 — Throne Inheritance

```go
package main

// LeetCode #1600: Throne Inheritance
// https://leetcode.com/problems/throne-inheritance/
// Difficulty: Medium

import "fmt"

func main() {
	t := ConstructorThrone("king")
	t.Birth("king", "andy")
	t.Birth("king", "bob")
	t.Birth("king", "catherine")
	t.Birth("andy", "matthew")
	t.Birth("bob", "alex")
	t.Birth("bob", "asha")

	inheritance := t.GetInheritanceOrder()
	fmt.Println("Inheritance order:", inheritance)

	t.Death("bob")
	inheritance2 := t.GetInheritanceOrder()
	fmt.Println("After Bob's death:", inheritance2)
}

type ThroneInheritance struct {
	king     string
	children map[string][]string
	dead     map[string]bool
}

func ConstructorThrone(kingName string) ThroneInheritance {
	return ThroneInheritance{
		king:     kingName,
		children: make(map[string][]string),
		dead:     make(map[string]bool),
	}
}

func (t *ThroneInheritance) Birth(parentName string, childName string) {
	t.children[parentName] = append(t.children[parentName], childName)
}

func (t *ThroneInheritance) Death(name string) {
	t.dead[name] = true
}

func (t *ThroneInheritance) GetInheritanceOrder() []string {
	// Preorder traversal of the family tree
	result := make([]string, 0)
	t.dfs(t.king, &result)
	return result
}

func (t *ThroneInheritance) dfs(name string, result *[]string) {
	if !t.dead[name] {
		*result = append(*result, name)
	}
	for _, child := range t.children[name] {
		t.dfs(child, result)
	}
}
```

## 1602 — Find Nearest Right Node In Binary Tree

```go
package main

// LeetCode #1602: Find Nearest Right Node in Binary Tree
// https://leetcode.com/problems/find-nearest-right-node-in-binary-tree/
// Difficulty: Medium [Paid]

import "fmt"

// TreeNode is a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Tree: [1, 2, 3, null, 4, 5, 6]
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2, Right: &TreeNode{Val: 4}}
	root.Right = &TreeNode{Val: 3, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 6}}

	// Find nearest right node of node with value 4
	u := root.Left.Right // node 4
	result := FindNearestRightNode(root, u)
	if result != nil {
		fmt.Println("Nearest right of 4:", result.Val) // should be 5
	} else {
		fmt.Println("Nearest right of 4: nil")
	}

	// Tree: [3, 4, 2, null, null, null, 1]
	root2 := &TreeNode{Val: 3}
	root2.Left = &TreeNode{Val: 4}
	root2.Right = &TreeNode{Val: 2, Right: &TreeNode{Val: 1}}

	result2 := FindNearestRightNode(root2, root2.Left)
	if result2 != nil {
		fmt.Println("Nearest right of 4:", result2.Val)
	} else {
		fmt.Println("Nearest right of 4: nil")
	}
}

func FindNearestRightNode(root *TreeNode, u *TreeNode) *TreeNode {
	// Time: O(N), Space: O(N)
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if node == u {
				// Return the next node in the queue (right sibling)
				if i+1 < levelSize {
					return queue[0]
				}
				return nil
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return nil
}
```

## 1604 — Alert Using Same Key Card Three Or More Times In A One Hour Period

```go
package main

// LeetCode #1604: Alert Using Same Key-Card Three or More Times in a One Hour Period
// https://leetcode.com/problems/alert-using-same-key-card-three-or-more-times-in-a-one-hour-period/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(AlertNames([]string{"daniel", "daniel", "daniel", "luis", "luis", "luis", "luis"},
		[]string{"10:00", "10:40", "11:00", "09:00", "11:00", "13:00", "15:00"}))
	fmt.Println(AlertNames([]string{"alice", "alice", "alice", "bob", "bob", "bob", "bob"},
		[]string{"12:01", "12:00", "18:00", "21:00", "21:20", "21:30", "23:00"}))
	fmt.Println(AlertNames([]string{"a", "a", "a", "a", "b"},
		[]string{"00:00", "00:50", "01:00", "01:30", "00:10"}))
}

func AlertNames(keyName []string, keyTime []string) []string {
	// Time: O(N log N), Space: O(N)
	n := len(keyName)
	records := make(map[string][]int)

	for i := 0; i < n; i++ {
		minutes := parseTime(keyTime[i])
		records[keyName[i]] = append(records[keyName[i]], minutes)
	}

	alerted := make([]string, 0)

	for name, times := range records {
		sort.Ints(times)
		for i := 2; i < len(times); i++ {
			if times[i]-times[i-2] <= 60 {
				alerted = append(alerted, name)
				break
			}
		}
	}

	sort.Strings(alerted)
	return alerted
}

func parseTime(t string) int {
	hours := int(t[0]-'0')*10 + int(t[1]-'0')
	minutes := int(t[3]-'0')*10 + int(t[4]-'0')
	return hours*60 + minutes
}
```

## 1605 — Find Valid Matrix Given Row And Column Sums

```go
package main

// LeetCode #1605: Find Valid Matrix Given Row and Column Sums
// https://leetcode.com/problems/find-valid-matrix-given-row-and-column-sums/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(RestoreMatrix([]int{3, 8}, []int{4, 7}))
	fmt.Println(RestoreMatrix([]int{5, 7, 10}, []int{8, 6, 8}))
	fmt.Println(RestoreMatrix([]int{14, 9}, []int{6, 9, 8}))
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func RestoreMatrix(rowSum []int, colSum []int) [][]int {
	// Time: O(R*C), Space: O(R*C)
	rows, cols := len(rowSum), len(colSum)
	result := make([][]int, rows)
	for i := 0; i < rows; i++ {
		result[i] = make([]int, cols)
	}

	i, j := 0, 0
	for i < rows && j < cols {
		val := minInt(rowSum[i], colSum[j])
		result[i][j] = val
		rowSum[i] -= val
		colSum[j] -= val

		if rowSum[i] == 0 {
			i++
		}
		if colSum[j] == 0 {
			j++
		}
	}

	return result
}
```

## 1609 — Even Odd Tree

```go
package main

// LeetCode #1609: Even Odd Tree
// https://leetcode.com/problems/even-odd-tree/
// Difficulty: Medium

import "fmt"

// TreeNode is a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Tree: [1,10,4,3,null,7,9,12,8,6,null,null,2]
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 10, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 12}, Right: &TreeNode{Val: 8}}}
	root.Right = &TreeNode{Val: 4, Left: &TreeNode{Val: 7, Left: &TreeNode{Val: 6}}, Right: &TreeNode{Val: 9, Right: &TreeNode{Val: 2}}}
	fmt.Println(IsEvenOddTree(root))

	// Simple tree
	root2 := &TreeNode{Val: 5, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 2}}
	fmt.Println(IsEvenOddTree(root2))

	// Tree: [5,9,1,3,5,7]
	root3 := &TreeNode{Val: 5}
	root3.Left = &TreeNode{Val: 9, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 5}}
	root3.Right = &TreeNode{Val: 1, Left: &TreeNode{Val: 7}}
	fmt.Println(IsEvenOddTree(root3))
}

func IsEvenOddTree(root *TreeNode) bool {
	// Time: O(N), Space: O(N)
	if root == nil {
		return true
	}

	queue := []*TreeNode{root}
	level := 0

	for len(queue) > 0 {
		levelSize := len(queue)
		var prevVal int

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if level%2 == 0 {
				// Even level: odd values, strictly increasing
				if node.Val%2 == 0 {
					return false
				}
				if i > 0 && node.Val <= prevVal {
					return false
				}
			} else {
				// Odd level: even values, strictly decreasing
				if node.Val%2 != 0 {
					return false
				}
				if i > 0 && node.Val >= prevVal {
					return false
				}
			}

			prevVal = node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		level++
	}

	return true
}
```

## 1612 — Check If Two Expression Trees Are Equivalent

```go
package main

// LeetCode #1612: Check If Two Expression Trees are Equivalent
// https://leetcode.com/problems/check-if-two-expression-trees-are-equivalent/
// Difficulty: Medium [Paid]

import "fmt"

// Node is an expression tree node.
// Leaf nodes contain lowercase letters.
// Non-leaf nodes contain '+' operator.
type Node struct {
	Val    byte
	Left   *Node
	Right  *Node
}

func main() {
	// Expression tree 1: (a + (b + c))
	root1 := &Node{Val: '+'}
	root1.Left = &Node{Val: 'a'}
	root1.Right = &Node{Val: '+', Left: &Node{Val: 'b'}, Right: &Node{Val: 'c'}}

	// Expression tree 2: ((b + a) + c)
	root2 := &Node{Val: '+'}
	root2.Left = &Node{Val: '+', Left: &Node{Val: 'b'}, Right: &Node{Val: 'a'}}
	root2.Right = &Node{Val: 'c'}

	fmt.Println(CheckEquivalence(root1, root2)) // true

	// Different trees
	root3 := &Node{Val: '+'}
	root3.Left = &Node{Val: 'a'}
	root3.Right = &Node{Val: 'b'}

	fmt.Println(CheckEquivalence(root1, root3)) // false
}

func CheckEquivalence(root1 *Node, root2 *Node) bool {
	// Time: O(N), Space: O(1)
	// Expression trees only contain '+', so equivalence = same multiset of leaf values
	count1 := make(map[byte]int)
	count2 := make(map[byte]int)

	countLeaves(root1, count1)
	countLeaves(root2, count2)

	if len(count1) != len(count2) {
		return false
	}
	for k, v := range count1 {
		if count2[k] != v {
			return false
		}
	}
	return true
}

func countLeaves(node *Node, count map[byte]int) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil {
		count[node.Val]++
		return
	}
	countLeaves(node.Left, count)
	countLeaves(node.Right, count)
}
```

## 1613 — Find The Missing Ids

```go
package main

// LeetCode #1613: Find the Missing IDs
// https://leetcode.com/problems/find-the-missing-ids/
// Difficulty: Medium [Paid]

import (
	"fmt"
	"sort"
)

func main() {
	// SQL problem: find missing customer IDs within the range.
	// Translated to Go.
	// Table: Customers(customer_id)

	customerIDs := []int{1, 2, 4, 7, 8, 10}
	missing := FindMissingIDs(customerIDs)
	fmt.Println("Missing IDs:", missing)
}

func FindMissingIDs(customerIDs []int) []int {
	// Time: O(N log N), Space: O(1)
	if len(customerIDs) == 0 {
		return nil
	}

	sort.Ints(customerIDs)
	result := make([]int, 0)

	// IDs range from 1 to max(customerID)
	for i := 1; i < customerIDs[len(customerIDs)-1]; i++ {
		// Binary search
		idx := sort.SearchInts(customerIDs, i)
		if idx == len(customerIDs) || customerIDs[idx] != i {
			result = append(result, i)
		}
	}

	return result
}
```

## 1615 — Maximal Network Rank

```go
package main

// LeetCode #1615: Maximal Network Rank
// https://leetcode.com/problems/maximal-network-rank/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MaximalNetworkRank(4, [][]int{{0, 1}, {0, 3}, {1, 2}, {1, 3}}))
	fmt.Println(MaximalNetworkRank(5, [][]int{{0, 1}, {0, 3}, {1, 2}, {1, 3}, {2, 3}, {2, 4}}))
	fmt.Println(MaximalNetworkRank(2, [][]int{{0, 1}}))
}

func MaximalNetworkRank(n int, roads [][]int) int {
	// Time: O(N^2), Space: O(N^2)
	degree := make([]int, n)
	connected := make([][]bool, n)
	for i := 0; i < n; i++ {
		connected[i] = make([]bool, n)
	}

	for _, r := range roads {
		u, v := r[0], r[1]
		degree[u]++
		degree[v]++
		connected[u][v] = true
		connected[v][u] = true
	}

	maxRank := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			rank := degree[i] + degree[j]
			if connected[i][j] {
				rank-- // shared edge counted twice
			}
			if rank > maxRank {
				maxRank = rank
			}
		}
	}

	return maxRank
}
```

## 1616 — Split Two Strings To Make Palindrome

```go
package main

// LeetCode #1616: Split Two Strings to Make Palindrome
// https://leetcode.com/problems/split-two-strings-to-make-palindrome/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(CheckPalindromeFormation("x", "y"))
	fmt.Println(CheckPalindromeFormation("abdef", "fecab"))
	fmt.Println(CheckPalindromeFormation("ulacfd", "jizalu"))
}

func CheckPalindromeFormation(a string, b string) bool {
	// Time: O(N), Space: O(1)
	return canForm(a, b) || canForm(b, a)
}

func canForm(a, b string) bool {
	left := 0
	right := len(a) - 1

	// Find the first mismatch from the outside
	for left < right && a[left] == b[right] {
		left++
		right--
	}

	if left >= right {
		return true
	}

	// Check if a[left..right] is palindrome
	return isPalindrome(a, left, right) || isPalindrome(b, left, right)
}

func isPalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
```

## 1618 — Maximum Font To Fit A Sentence In A Screen

```go
package main

// LeetCode #1618: Maximum Font to Fit a Sentence in a Screen
// https://leetcode.com/problems/maximum-font-to-fit-a-sentence-in-a-screen/
// Difficulty: Medium [Paid]

import "fmt"

func main() {
	// Problem: Given a screen width and height, find the maximum font size
	// that can display all characters of a string within the screen.
	// We have a FontInfo API that gives width and height for a font size.

	sentence := "Hello World"
	width := 80
	height := 30

	// Available font sizes
	fonts := []int{6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 36, 48, 72}

	// Mock font info
	fontInfo := &mockFontInfo{}

	maxFont := MaxFont(sentence, width, height, fonts, fontInfo)
	fmt.Println("Maximum font size:", maxFont)
}

type FontInfo interface {
	getFontWidth(fontSize int) int
	getFontHeight(fontSize int) int
}

type mockFontInfo struct{}

func (m *mockFontInfo) getFontWidth(fontSize int) int {
	// Approximate: width ~ font_size * 0.6 per character
	return fontSize * 6 / 10
}

func (m *mockFontInfo) getFontHeight(fontSize int) int {
	return fontSize
}

func MaxFont(text string, w int, h int, fonts []int, fontInfo interface{ getFontWidth(int) int; getFontHeight(int) int }) int {
	// Time: O(log N * L), Space: O(1)
	// Binary search on font sizes
	left, right := 0, len(fonts)-1
	result := -1

	for left <= right {
		mid := left + (right-left)/2
		fontSize := fonts[mid]

		charWidth := fontInfo.getFontWidth(fontSize)
		charHeight := fontInfo.getFontHeight(fontSize)

		// Check if text fits
		charsPerLine := w / charWidth
		if charsPerLine == 0 {
			right = mid - 1
			continue
		}

		lines := (len(text) + charsPerLine - 1) / charsPerLine
		if lines*charHeight <= h {
			result = fontSize
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result
}
```

## 1620 — Coordinate With Maximum Network Quality

```go
package main

// LeetCode #1620: Coordinate With Maximum Network Quality
// https://leetcode.com/problems/coordinate-with-maximum-network-quality/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(BestCoordinate([][]int{{1, 2, 5}, {2, 1, 7}, {3, 1, 9}}, 2))
	fmt.Println(BestCoordinate([][]int{{23, 11, 21}}, 9))
	fmt.Println(BestCoordinate([][]int{{2, 1, 9}, {0, 1, 9}}, 2))
}

func BestCoordinate(towers [][]int, radius int) []int {
	// Time: O(N * R^2), Space: O(1)
	// Signal quality for a point (x,y) = sum of floor(tower_signal / (1 + d))
	// where d = Euclidean distance from tower to (x,y)
	// Range of coordinates is bounded by tower positions +/- radius

	minX, minY := 51, 51
	maxX, maxY := 0, 0

	for _, t := range towers {
		if t[0] < minX {
			minX = t[0]
		}
		if t[0] > maxX {
			maxX = t[0]
		}
		if t[1] < minY {
			minY = t[1]
		}
		if t[1] > maxY {
			maxY = t[1]
		}
	}

	bestX, bestY := 0, 0
	bestQuality := 0

	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			quality := 0
			for _, t := range towers {
				d := sqrtDist(x, y, t[0], t[1])
				if d > radius*radius {
					continue
				}
				// signal = floor(t[2] / (1 + sqrt(d)))
				// We avoid sqrt by computing: quality += t[2] / (1 + sqrt(d))
				// For integer comparisons we use the formula directly
				signal := float64(t[2]) / (1.0 + sqrt(float64(d)))
				quality += int(signal)
			}
			if quality > bestQuality || (quality == bestQuality && (x < bestX || (x == bestX && y < bestY))) {
				bestQuality = quality
				bestX, bestY = x, y
			}
		}
	}

	return []int{bestX, bestY}
}

func sqrtDist(x1, y1, x2, y2 int) int {
	dx := x1 - x2
	dy := y1 - y2
	return dx*dx + dy*dy
}

func sqrt(n float64) float64 {
	if n <= 0 {
		return 0
	}
	x := n
	y := (x + 1) / 2
	for y < x {
		x = y
		y = (x + n/x) / 2
	}
	return x
}
```

## 1621 — Number Of Sets Of K Non Overlapping Line Segments

```go
package main

// LeetCode #1621: Number of Sets of K Non-Overlapping Line Segments
// https://leetcode.com/problems/number-of-sets-of-k-non-overlapping-line-segments/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(NumberOfSets(4, 2))
	fmt.Println(NumberOfSets(3, 1))
	fmt.Println(NumberOfSets(30, 7))
}

func NumberOfSets(n int, k int) int {
	// Time: O(N*K), Space: O(N)
	const mod = 1_000_000_007

	// dp[i][j][0] = ways using i points with j segments, not ending at i
	// dp[i][j][1] = ways using i points with j segments, ending at i

	dp := make([][][2]int, n+1)
	for i := range dp {
		dp[i] = make([][2]int, k+1)
	}

	// Base: with 0 segments, there's 1 way
	for i := 1; i <= n; i++ {
		dp[i][0][0] = 1
	}

	for j := 1; j <= k; j++ {
		for i := 2; i <= n; i++ {
			// Not ending at i: just carry forward
			dp[i][j][0] = (dp[i-1][j][0] + dp[i-1][j][1]) % mod

			// Ending at i: either extend previous ending at i-1, or start new ending at i
			dp[i][j][1] = (dp[i-1][j-1][0] + dp[i-1][j][1]) % mod
		}
	}

	return (dp[n][k][0] + dp[n][k][1]) % mod
}
```

## 1625 — Lexicographically Smallest String After Applying Operations

```go
package main

// LeetCode #1625: Lexicographically Smallest String After Applying Operations
// https://leetcode.com/problems/lexicographically-smallest-string-after-applying-operations/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(FindLexSmallestString("5525", 9, 2))
	fmt.Println(FindLexSmallestString("74", 5, 1))
	fmt.Println(FindLexSmallestString("0011", 4, 2))
}

func FindLexSmallestString(s string, a int, b int) string {
	// Time: O(N^2), Space: O(N)
	// BFS over all possible states
	n := len(s)
	seen := make(map[string]bool)
	queue := []string{s}
	seen[s] = true
	smallest := s

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr < smallest {
			smallest = curr
		}

		// Operation 1: add a to odd-position digits
		bytes := []byte(curr)
		for i := 1; i < len(bytes); i += 2 {
			val := int(bytes[i] - '0')
			val = (val + a) % 10
			bytes[i] = byte('0' + val)
		}
		next1 := string(bytes)

		// Operation 2: rotate right by b
		next2 := curr[n-b:] + curr[:n-b]

		for _, next := range []string{next1, next2} {
			if !seen[next] {
				seen[next] = true
				queue = append(queue, next)
			}
		}
	}

	return smallest
}
```

## 1626 — Best Team With No Conflicts

```go
package main

// LeetCode #1626: Best Team With No Conflicts
// https://leetcode.com/problems/best-team-with-no-conflicts/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(BestTeamScore([]int{1, 3, 5, 10, 15}, []int{1, 2, 3, 4, 5}))
	fmt.Println(BestTeamScore([]int{4, 5, 6, 5}, []int{2, 1, 2, 1}))
	fmt.Println(BestTeamScore([]int{1, 2, 3, 5}, []int{8, 9, 10, 1}))
}

func BestTeamScore(scores []int, ages []int) int {
	// Time: O(N^2), Space: O(N)
	n := len(scores)
	players := make([][2]int, n)
	for i := 0; i < n; i++ {
		players[i] = [2]int{ages[i], scores[i]}
	}

	// Sort by age, then by score
	sort.Slice(players, func(i, j int) bool {
		if players[i][0] != players[j][0] {
			return players[i][0] < players[j][0]
		}
		return players[i][1] < players[j][1]
	})

	// LIS-like DP
	dp := make([]int, n)
	maxScore := 0

	for i := 0; i < n; i++ {
		dp[i] = players[i][1]
		for j := 0; j < i; j++ {
			if players[j][1] <= players[i][1] {
				if dp[j]+players[i][1] > dp[i] {
					dp[i] = dp[j] + players[i][1]
				}
			}
		}
		if dp[i] > maxScore {
			maxScore = dp[i]
		}
	}

	return maxScore
}
```

## 1628 — Design An Expression Tree With Evaluate Function

```go
package main

// LeetCode #1628: Design an Expression Tree With Evaluate Function
// https://leetcode.com/problems/design-an-expression-tree-with-evaluate-function/
// Difficulty: Medium [Paid]

import (
	"fmt"
	"strconv"
)

func main() {
	// Build expression: (3 + 4) * (5 - 2)
	// Postfix: ["3", "4", "+", "5", "2", "-", "*"]
	postfix := []string{"3", "4", "+", "5", "2", "-", "*"}
	tree := BuildExpressionTree(postfix)
	result := tree.Evaluate()
	fmt.Println("Result:", result) // should be 21

	// Simple: 2 + 3
	postfix2 := []string{"2", "3", "+"}
	tree2 := BuildExpressionTree(postfix2)
	fmt.Println("Result:", tree2.Evaluate()) // should be 5
}

// Node is an expression tree node.
type Node struct {
	val   string
	left  *Node
	right *Node
}

func (n *Node) Evaluate() int {
	if n.left == nil && n.right == nil {
		val, _ := strconv.Atoi(n.val)
		return val
	}

	leftVal := n.left.Evaluate()
	rightVal := n.right.Evaluate()

	switch n.val {
	case "+":
		return leftVal + rightVal
	case "-":
		return leftVal - rightVal
	case "*":
		return leftVal * rightVal
	case "/":
		return leftVal / rightVal
	}
	return 0
}

func BuildExpressionTree(postfix []string) *Node {
	// Time: O(N), Space: O(N)
	stack := make([]*Node, 0)

	for _, token := range postfix {
		node := &Node{val: token}
		if token == "+" || token == "-" || token == "*" || token == "/" {
			node.right = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			node.left = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, node)
	}

	return stack[0]
}
```

## 1630 — Arithmetic Subarrays

```go
package main

// LeetCode #1630: Arithmetic Subarrays
// https://leetcode.com/problems/arithmetic-subarrays/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(CheckArithmeticSubarrays([]int{4, 6, 5, 9, 3, 7}, []int{0, 0, 2}, []int{2, 3, 5}))
	fmt.Println(CheckArithmeticSubarrays([]int{-12, -9, -3, -12, -6, 15, 20, -25, -20, -15, -10}, []int{0, 1, 6, 4, 8, 7}, []int{4, 4, 9, 7, 9, 10}))
}

func CheckArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	// Time: O(M * N log N), Space: O(N)
	result := make([]bool, len(l))

	for i := 0; i < len(l); i++ {
		sub := make([]int, r[i]-l[i]+1)
		copy(sub, nums[l[i]:r[i]+1])
		result[i] = isArithmetic(sub)
	}

	return result
}

func isArithmetic(arr []int) bool {
	if len(arr) <= 2 {
		return true
	}

	sort.Ints(arr)
	diff := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != diff {
			return false
		}
	}
	return true
}
```

## 1631 — Path With Minimum Effort

```go
package main

// LeetCode #1631: Path With Minimum Effort
// https://leetcode.com/problems/path-with-minimum-effort/
// Difficulty: Medium

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	fmt.Println(MinimumEffortPath([][]int{{1, 2, 2}, {3, 8, 2}, {5, 3, 5}}))
	fmt.Println(MinimumEffortPath([][]int{{1, 2, 3}, {3, 8, 4}, {5, 3, 5}}))
	fmt.Println(MinimumEffortPath([][]int{{1, 10, 6, 7, 9, 10, 4, 9}}))
}

type Point struct {
	x, y, effort int
	index        int
}

type EffortPQ []*Point

func (pq EffortPQ) Len() int           { return len(pq) }
func (pq EffortPQ) Less(i, j int) bool { return pq[i].effort < pq[j].effort }
func (pq EffortPQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *EffortPQ) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Point)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *EffortPQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[:n-1]
	return item
}

func MinimumEffortPath(heights [][]int) int {
	// Time: O(R*C log(R*C)), Space: O(R*C)
	// Dijkstra-like: track minimum effort to reach each cell
	rows, cols := len(heights), len(heights[0])
	if rows == 0 || cols == 0 {
		return 0
	}

	effort := make([][]int, rows)
	for i := 0; i < rows; i++ {
		effort[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			effort[i][j] = math.MaxInt32
		}
	}
	effort[0][0] = 0

	pq := &EffortPQ{}
	heap.Push(pq, &Point{x: 0, y: 0, effort: 0})

	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	for pq.Len() > 0 {
		p := heap.Pop(pq).(*Point)
		if p.effort > effort[p.x][p.y] {
			continue
		}
		if p.x == rows-1 && p.y == cols-1 {
			return p.effort
		}

		for _, d := range dirs {
			nx, ny := p.x+d[0], p.y+d[1]
			if nx < 0 || nx >= rows || ny < 0 || ny >= cols {
				continue
			}
			diff := heights[p.x][p.y] - heights[nx][ny]
			if diff < 0 {
				diff = -diff
			}
			newEffort := p.effort
			if diff > newEffort {
				newEffort = diff
			}
			if newEffort < effort[nx][ny] {
				effort[nx][ny] = newEffort
				heap.Push(pq, &Point{x: nx, y: ny, effort: newEffort})
			}
		}
	}

	return 0
}
```

## 1634 — Add Two Polynomials Represented As Linked Lists

```go
package main

// LeetCode #1634: Add Two Polynomials Represented as Linked Lists
// https://leetcode.com/problems/add-two-polynomials-represented-as-linked-lists/
// Difficulty: Medium [Paid]

import "fmt"

// PolyNode represents a term in a polynomial.
type PolyNode struct {
	Coefficient int
	Power       int
	Next        *PolyNode
}

func main() {
	// Polynomial 1: 5x^3 + 4x^2 + (-1)x^0
	p1 := &PolyNode{Coefficient: 5, Power: 3}
	p1.Next = &PolyNode{Coefficient: 4, Power: 2}
	p1.Next.Next = &PolyNode{Coefficient: -1, Power: 0}

	// Polynomial 2: 3x^2 + 2x^1 + 1x^0
	p2 := &PolyNode{Coefficient: 3, Power: 2}
	p2.Next = &PolyNode{Coefficient: 2, Power: 1}
	p2.Next.Next = &PolyNode{Coefficient: 1, Power: 0}

	result := AddPoly(p1, p2)
	printPoly(result)

	// Polynomial: x^3 + (-1)x^0
	p3 := &PolyNode{Coefficient: 1, Power: 3, Next: &PolyNode{Coefficient: -1, Power: 0}}
	// Polynomial: x^2 + x^1
	p4 := &PolyNode{Coefficient: 1, Power: 2, Next: &PolyNode{Coefficient: 1, Power: 1}}
	result2 := AddPoly(p3, p4)
	printPoly(result2)
}

func printPoly(p *PolyNode) {
	first := true
	for p != nil {
		if p.Coefficient != 0 {
			if !first {
				fmt.Print(" + ")
			}
			fmt.Printf("%dx^%d", p.Coefficient, p.Power)
			first = false
		}
		p = p.Next
	}
	fmt.Println()
}

func AddPoly(poly1 *PolyNode, poly2 *PolyNode) *PolyNode {
	// Time: O(N+M), Space: O(1)
	dummy := &PolyNode{}
	curr := dummy

	for poly1 != nil && poly2 != nil {
		if poly1.Power > poly2.Power {
			if poly1.Coefficient != 0 {
				curr.Next = &PolyNode{Coefficient: poly1.Coefficient, Power: poly1.Power}
				curr = curr.Next
			}
			poly1 = poly1.Next
		} else if poly2.Power > poly1.Power {
			if poly2.Coefficient != 0 {
				curr.Next = &PolyNode{Coefficient: poly2.Coefficient, Power: poly2.Power}
				curr = curr.Next
			}
			poly2 = poly2.Next
		} else {
			sum := poly1.Coefficient + poly2.Coefficient
			if sum != 0 {
				curr.Next = &PolyNode{Coefficient: sum, Power: poly1.Power}
				curr = curr.Next
			}
			poly1 = poly1.Next
			poly2 = poly2.Next
		}
	}

	for poly1 != nil {
		if poly1.Coefficient != 0 {
			curr.Next = &PolyNode{Coefficient: poly1.Coefficient, Power: poly1.Power}
			curr = curr.Next
		}
		poly1 = poly1.Next
	}

	for poly2 != nil {
		if poly2.Coefficient != 0 {
			curr.Next = &PolyNode{Coefficient: poly2.Coefficient, Power: poly2.Power}
			curr = curr.Next
		}
		poly2 = poly2.Next
	}

	return dummy.Next
}
```

## 1638 — Count Substrings That Differ By One Character

```go
package main

// LeetCode #1638: Count Substrings That Differ by One Character
// https://leetcode.com/problems/count-substrings-that-differ-by-one-character/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(CountSubstrings("aba", "baba"))
	fmt.Println(CountSubstrings("ab", "bb"))
	fmt.Println(CountSubstrings("abe", "bbc"))
}

func CountSubstrings(s string, t string) int {
	// Time: O(N*M*min(N,M)), Space: O(1)
	count := 0

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(t); j++ {
			diff := 0
			k := 0
			for i+k < len(s) && j+k < len(t) && diff <= 1 {
				if s[i+k] != t[j+k] {
					diff++
				}
				if diff == 1 {
					count++
				}
				k++
			}
		}
	}

	return count
}
```

## 1641 — Count Sorted Vowel Strings

```go
package main

// LeetCode #1641: Count Sorted Vowel Strings
// https://leetcode.com/problems/count-sorted-vowel-strings/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(CountVowelStrings(1))
	fmt.Println(CountVowelStrings(2))
	fmt.Println(CountVowelStrings(33))
}

func CountVowelStrings(n int) int {
	// Time: O(N), Space: O(1)
	// Combinatorics: C(n+4, 4) = (n+4)*(n+3)*(n+2)*(n+1)/24
	return (n + 4) * (n + 3) * (n + 2) * (n + 1) / 24
}
```

## 1642 — Furthest Building You Can Reach

```go
package main

// LeetCode #1642: Furthest Building You Can Reach
// https://leetcode.com/problems/furthest-building-you-can-reach/
// Difficulty: Medium

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(FurthestBuilding([]int{4, 2, 7, 6, 9, 14, 12}, 5, 1))
	fmt.Println(FurthestBuilding([]int{4, 12, 2, 7, 3, 18, 20, 3, 19}, 10, 2))
	fmt.Println(FurthestBuilding([]int{14, 3, 19, 3}, 17, 0))
}

// MinHeap for int
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func FurthestBuilding(heights []int, bricks int, ladders int) int {
	// Time: O(N log K), Space: O(K) where K = ladders
	// Use a min-heap to store the largest climbs where ladders are used
	h := &IntHeap{}
	heap.Init(h)

	for i := 0; i < len(heights)-1; i++ {
		diff := heights[i+1] - heights[i]
		if diff <= 0 {
			continue
		}

		heap.Push(h, diff)

		// If we have more climbs than ladders, use bricks for the smallest climb
		if h.Len() > ladders {
			bricks -= heap.Pop(h).(int)
			if bricks < 0 {
				return i
			}
		}
	}

	return len(heights) - 1
}
```

## 1644 — Lowest Common Ancestor Of A Binary Tree Ii

```go
package main

// LeetCode #1644: Lowest Common Ancestor of a Binary Tree II
// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree-ii/
// Difficulty: Medium [Paid]

import "fmt"

// TreeNode is a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Tree: [3,5,1,6,2,0,8,null,null,7,4]
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 5, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 4}}}
	root.Right = &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 8}}

	// LCA of 5 and 1 is 3
	p := root.Left // 5
	q := root.Right // 1
	lca := LowestCommonAncestorII(root, p, q)
	if lca != nil {
		fmt.Println("LCA of 5 and 1:", lca.Val) // 3
	} else {
		fmt.Println("LCA: nil")
	}

	// LCA of 5 and 4 is 5
	p = root.Left // 5
	q = root.Left.Right.Right // 4
	lca = LowestCommonAncestorII(root, p, q)
	if lca != nil {
		fmt.Println("LCA of 5 and 4:", lca.Val) // 5
	}

	// Node not in tree
	notInTree := &TreeNode{Val: 10}
	lca = LowestCommonAncestorII(root, root.Left, notInTree)
	if lca == nil {
		fmt.Println("LCA of 5 and 10: nil (10 not in tree)")
	}
}

// LowestCommonAncestorII differs from LCA I in that p and q may not exist in the tree.
func LowestCommonAncestorII(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	// Time: O(N), Space: O(N)
	foundP := false
	foundQ := false

	var dfs func(node *TreeNode) *TreeNode
	dfs = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}

		left := dfs(node.Left)
		right := dfs(node.Right)

		if node == p {
			foundP = true
			return node
		}
		if node == q {
			foundQ = true
			return node
		}

		if left != nil && right != nil {
			return node
		}
		if left != nil {
			return left
		}
		return right
	}

	lca := dfs(root)

	if foundP && foundQ {
		return lca
	}
	return nil
}
```

## 1647 — Minimum Deletions To Make Character Frequencies Unique

```go
package main

// LeetCode #1647: Minimum Deletions to Make Character Frequencies Unique
// https://leetcode.com/problems/minimum-deletions-to-make-character-frequencies-unique/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MinDeletions("aab"))
	fmt.Println(MinDeletions("aaabbbcc"))
	fmt.Println(MinDeletions("ceabaacb"))
}

func MinDeletions(s string) int {
	// Time: O(N), Space: O(1)
	freq := make([]int, 26)
	for _, ch := range s {
		freq[ch-'a']++
	}

	used := make(map[int]bool)
	deletions := 0

	for _, f := range freq {
		for f > 0 && used[f] {
			f--
			deletions++
		}
		if f > 0 {
			used[f] = true
		}
	}

	return deletions
}
```

## 1648 — Sell Diminishing Valued Colored Balls

```go
package main

// LeetCode #1648: Sell Diminishing-Valued Colored Balls
// https://leetcode.com/problems/sell-diminishing-valued-colored-balls/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MaxProfit([]int{2, 5}, 4))
	fmt.Println(MaxProfit([]int{3, 5}, 6))
	fmt.Println(MaxProfit([]int{2, 8, 4, 10, 6}, 20))
}

func MaxProfit(inventory []int, orders int) int {
	// Time: O(N log N), Space: O(N)
	const mod = 1_000_000_007

	// Sort descending
	sort.Slice(inventory, func(i, j int) bool {
		return inventory[i] > inventory[j]
	})

	// Append 0 for convenience
	inventory = append(inventory, 0)
	n := len(inventory)

	profit := 0
	count := 0

	for i := 0; i < n-1; i++ {
		if inventory[i] == inventory[i+1] {
			continue
		}

		// Height difference between current and next level
		height := inventory[i] - inventory[i+1]
		width := i + 1 // number of types with this count
		total := height * width

		if count+total <= orders {
			// Take all balls at this level
			// Sum for this level: width * sum of (inventory[i], inventory[i]-1, ..., inventory[i+1]+1)
			top := inventory[i]
			bottom := inventory[i+1] + 1
			sumLevel := (top + bottom) * height / 2
			profit = (profit + sumLevel*width) % mod
			count += total
		} else {
			// Take only some
			remaining := orders - count
			fullRows := remaining / width
			extra := remaining % width

			top := inventory[i]
			bottom := top - fullRows + 1
			sumFull := (top + bottom) * fullRows / 2
			profit = (profit + sumFull*width) % mod
			profit = (profit + bottom*extra) % mod

			count = orders
			break
		}
	}

	return profit
}
```

## 1650 — Lowest Common Ancestor Of A Binary Tree Iii

```go
package main

// LeetCode #1650: Lowest Common Ancestor of a Binary Tree III
// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree-iii/
// Difficulty: Medium [Paid]

import "fmt"

// Node with parent pointer.
type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	Parent *Node
}

func main() {
	// Tree: [3,5,1,6,2,0,8,null,null,7,4]
	root := &Node{Val: 3}
	n5 := &Node{Val: 5, Parent: root}
	n1 := &Node{Val: 1, Parent: root}
	root.Left = n5
	root.Right = n1

	n6 := &Node{Val: 6, Parent: n5}
	n2 := &Node{Val: 2, Parent: n5}
	n5.Left = n6
	n5.Right = n2

	n0 := &Node{Val: 0, Parent: n1}
	n8 := &Node{Val: 8, Parent: n1}
	n1.Left = n0
	n1.Right = n8

	n7 := &Node{Val: 7, Parent: n2}
	n4 := &Node{Val: 4, Parent: n2}
	n2.Left = n7
	n2.Right = n4

	p := n5 // 5
	q := n1 // 1
	lca := LowestCommonAncestorIII(p, q)
	fmt.Println("LCA of 5 and 1:", lca.Val) // 3

	p = n5 // 5
	q = n4 // 4
	lca = LowestCommonAncestorIII(p, q)
	fmt.Println("LCA of 5 and 4:", lca.Val) // 5
}

func LowestCommonAncestorIII(p *Node, q *Node) *Node {
	// Time: O(H), Space: O(1)
	// Same as intersection of two linked lists
	a, b := p, q
	for a != b {
		if a == nil {
			a = q
		} else {
			a = a.Parent
		}
		if b == nil {
			b = p
		} else {
			b = b.Parent
		}
	}
	return a
}
```

## 1653 — Minimum Deletions To Make String Balanced

```go
package main

// LeetCode #1653: Minimum Deletions to Make String Balanced
// https://leetcode.com/problems/minimum-deletions-to-make-string-balanced/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MinimumDeletions("aababbab"))
	fmt.Println(MinimumDeletions("bbaaaaabb"))
	fmt.Println(MinimumDeletions("a"))
}

func MinimumDeletions(s string) int {
	// Time: O(N), Space: O(1)
	// Count 'a's on the right
	aCount := 0
	for _, ch := range s {
		if ch == 'a' {
			aCount++
		}
	}

	bCount := 0
	minDeletions := len(s)

	for _, ch := range s {
		if ch == 'a' {
			aCount--
		}

		// Deletions needed: remove all 'b's before this point + remove all 'a's after
		deletions := bCount + aCount
		if deletions < minDeletions {
			minDeletions = deletions
		}

		if ch == 'b' {
			bCount++
		}
	}

	return minDeletions
}
```

## 1654 — Minimum Jumps To Reach Home

```go
package main

// LeetCode #1654: Minimum Jumps to Reach Home
// https://leetcode.com/problems/minimum-jumps-to-reach-home/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MinimumJumps([]int{14, 4, 18, 1, 15}, 3, 15, 9))
	fmt.Println(MinimumJumps([]int{8, 3, 16, 6, 12, 20}, 15, 13, 11))
	fmt.Println(MinimumJumps([]int{1, 6, 2, 14, 5, 17, 4}, 16, 9, 7))
}

func MinimumJumps(forbidden []int, a int, b int, x int) int {
	// Time: O(limit), Space: O(limit)
	// BFS with state (position, direction) where direction 0=right, 1=left
	forbiddenSet := make(map[int]bool)
	for _, f := range forbidden {
		forbiddenSet[f] = true
	}

	// Upper bound: we shouldn't go beyond max(x, max(forbidden)) + a + b
	limit := x
	for _, f := range forbidden {
		if f > limit {
			limit = f
		}
	}
	limit += a + b

	type state struct {
		pos int
		dir int // 0 = came from left (can go either way), 1 = came from right (can only go right)
	}

	visited := make(map[struct{ pos, dir int }]bool)
	queue := []state{{pos: 0, dir: 0}}
	visited[struct{ pos, dir int }{0, 0}] = true
	jumps := 0

	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			curr := queue[0]
			queue = queue[1:]

			if curr.pos == x {
				return jumps
			}

			// Move forward (right)
			nextPos := curr.pos + a
			if nextPos <= limit && !forbiddenSet[nextPos] {
				key := struct{ pos, dir int }{nextPos, 0}
				if !visited[key] {
					visited[key] = true
					queue = append(queue, state{pos: nextPos, dir: 0})
				}
			}

			// Move backward (left) only if we didn't come from a backward move
			if curr.dir != 1 {
				nextPos = curr.pos - b
				if nextPos >= 0 && !forbiddenSet[nextPos] {
					key := struct{ pos, dir int }{nextPos, 1}
					if !visited[key] {
						visited[key] = true
						queue = append(queue, state{pos: nextPos, dir: 1})
					}
				}
			}
		}
		jumps++
	}

	return -1
}
```

## 1657 — Determine If Two Strings Are Close

```go
package main

// LeetCode #1657: Determine if Two Strings Are Close
// https://leetcode.com/problems/determine-if-two-strings-are-close/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(CloseStrings("abc", "bca"))
	fmt.Println(CloseStrings("a", "aa"))
	fmt.Println(CloseStrings("cabbba", "abbccc"))
}

func CloseStrings(word1 string, word2 string) bool {
	// Time: O(N + M + 26 log 26), Space: O(1)
	if len(word1) != len(word2) {
		return false
	}

	freq1 := make([]int, 26)
	freq2 := make([]int, 26)
	set1 := make([]bool, 26)
	set2 := make([]bool, 26)

	for _, ch := range word1 {
		freq1[ch-'a']++
		set1[ch-'a'] = true
	}
	for _, ch := range word2 {
		freq2[ch-'a']++
		set2[ch-'a'] = true
	}

	// Check same character set
	for i := 0; i < 26; i++ {
		if set1[i] != set2[i] {
			return false
		}
	}

	// Check same frequency multiset
	sort.Ints(freq1)
	sort.Ints(freq2)
	for i := 0; i < 26; i++ {
		if freq1[i] != freq2[i] {
			return false
		}
	}

	return true
}
```

## 1658 — Minimum Operations To Reduce X To Zero

```go
package main

// LeetCode #1658: Minimum Operations to Reduce X to Zero
// https://leetcode.com/problems/minimum-operations-to-reduce-x-to-zero/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MinOperations([]int{1, 1, 4, 2, 3}, 5))
	fmt.Println(MinOperations([]int{5, 6, 7, 8, 9}, 4))
	fmt.Println(MinOperations([]int{3, 2, 20, 1, 1, 3}, 10))
}

func MinOperations(nums []int, x int) int {
	// Time: O(N), Space: O(1)
	// Find longest subarray with sum = total - x
	total := 0
	for _, num := range nums {
		total += num
	}

	target := total - x
	if target < 0 {
		return -1
	}
	if target == 0 {
		return len(nums)
	}

	maxLen := -1
	left := 0
	currSum := 0

	for right := 0; right < len(nums); right++ {
		currSum += nums[right]

		for currSum > target {
			currSum -= nums[left]
			left++
		}

		if currSum == target {
			if right-left+1 > maxLen {
				maxLen = right - left + 1
			}
		}
	}

	if maxLen == -1 {
		return -1
	}
	return len(nums) - maxLen
}
```

## 1660 — Correct A Binary Tree

```go
package main

// LeetCode #1660: Correct a Binary Tree
// https://leetcode.com/problems/correct-a-binary-tree/
// Difficulty: Medium [Paid]

import "fmt"

// TreeNode is a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Tree where a node has an invalid right pointer pointing to a node at the same level or below.
	// Example: 1 -> 2, 3; 2 -> 4 (right points to 3)
	//          1
	//         / \
	//        2   3
	//         \
	//          4 (right -> 3, invalid)
	n3 := &TreeNode{Val: 3}
	n4 := &TreeNode{Val: 4, Right: n3} // invalid pointer
	n2 := &TreeNode{Val: 2, Right: n4}
	root := &TreeNode{Val: 1, Left: n2, Right: n3}

	corrected := CorrectBinaryTree(root)
	fmt.Println("Root val:", corrected.Val)
	fmt.Println("Left:", corrected.Left.Val)
	fmt.Println("Right:", corrected.Right.Val)
	// Left.Right should be 4 but 4's Right should be nil (corrected)
	if corrected.Left.Right != nil {
		fmt.Println("Left.Right.Right (should be nil):", corrected.Left.Right.Right)
	}
}

func CorrectBinaryTree(root *TreeNode) *TreeNode {
	// Time: O(N), Space: O(N)
	// BFS to find the invalid node, remove it
	visited := make(map[*TreeNode]bool)
	parent := make(map[*TreeNode]*TreeNode)
	queue := []*TreeNode{root}
	visited[root] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.Left != nil {
			if visited[node.Left] {
				// node.Left points to an already visited node (invalid)
				removeNode(parent, node, 'L')
				return root
			}
			visited[node.Left] = true
			parent[node.Left] = node
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			if visited[node.Right] {
				// node.Right points to an already visited node (invalid)
				removeNode(parent, node, 'R')
				return root
			}
			visited[node.Right] = true
			parent[node.Right] = node
			queue = append(queue, node.Right)
		}
	}

	return root
}

func removeNode(parent map[*TreeNode]*TreeNode, node *TreeNode, child byte) {
	if p, ok := parent[node]; ok {
		if p.Left == node {
			p.Left = nil
		} else {
			p.Right = nil
		}
	} else {
		// node is root, can't easily remove - but this case doesn't happen
		// since root can't have invalid pointers
	}
}
```

## 1663 — Smallest String With A Given Numeric Value

```go
package main

// LeetCode #1663: Smallest String With A Given Numeric Value
// https://leetcode.com/problems/smallest-string-with-a-given-numeric-value/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(GetSmallestString(3, 27))
	fmt.Println(GetSmallestString(5, 73))
	fmt.Println(GetSmallestString(1, 26))
}

func GetSmallestString(n int, k int) string {
	// Time: O(N), Space: O(N)
	// Greedy: fill from the end with 'z' as much as possible
	result := make([]byte, n)
	for i := range result {
		result[i] = 'a'
	}

	k -= n // All positions have at least 'a' (value 1)

	for i := n - 1; i >= 0 && k > 0; i-- {
		add := 25 // 'z' - 'a' = 25
		if k < add {
			add = k
		}
		result[i] = byte('a' + add)
		k -= add
	}

	return string(result)
}
```

## 1664 — Ways To Make A Fair Array

```go
package main

// LeetCode #1664: Ways to Make a Fair Array
// https://leetcode.com/problems/ways-to-make-a-fair-array/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(WaysToMakeFair([]int{2, 1, 6, 4}))
	fmt.Println(WaysToMakeFair([]int{1, 1, 1}))
	fmt.Println(WaysToMakeFair([]int{1, 2, 3, 4, 5}))
}

func WaysToMakeFair(nums []int) int {
	// Time: O(N), Space: O(1)

	// Calculate total sum at even and odd indices
	totalEven := 0
	totalOdd := 0
	for i, num := range nums {
		if i%2 == 0 {
			totalEven += num
		} else {
			totalOdd += num
		}
	}

	result := 0
	prefixEven := 0
	prefixOdd := 0

	for i, num := range nums {
		if i%2 == 0 {
			totalEven -= num
		} else {
			totalOdd -= num
		}

		// After removing nums[i], all indices shift:
		// Elements to the right of i swap parity
		// Even sum = prefixEven + totalOdd
		// Odd sum = prefixOdd + totalEven
		if prefixEven+totalOdd == prefixOdd+totalEven {
			result++
		}

		if i%2 == 0 {
			prefixEven += num
		} else {
			prefixOdd += num
		}
	}

	return result
}
```

