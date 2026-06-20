# Medium (Sedang) — Problem ��2049

## 1865 — Finding Pairs With A Certain Sum

```go
package main

// LeetCode #1865: Finding Pairs With a Certain Sum
// https://leetcode.com/problems/finding-pairs-with-a-certain-sum/
// Difficulty: Medium

import "fmt"

type FindSumPairs struct {
	nums1 []int
	nums2 []int
	freq  map[int]int
}

func Constructor(nums1 []int, nums2 []int) FindSumPairs {
	freq := make(map[int]int)
	for _, v := range nums2 {
		freq[v]++
	}
	return FindSumPairs{nums1: nums1, nums2: nums2, freq: freq}
}

func (this *FindSumPairs) Add(index int, val int) {
	old := this.nums2[index]
	this.freq[old]--
	this.nums2[index] += val
	this.freq[this.nums2[index]]++
}

func (this *FindSumPairs) Count(tot int) int {
	ans := 0
	for _, v := range this.nums1 {
		ans += this.freq[tot-v]
	}
	return ans
}

func main() {
	obj := Constructor([]int{1, 1, 2, 2, 2, 3}, []int{1, 4, 5, 2, 5, 4})
	fmt.Println(obj.Count(7))
	obj.Add(3, 2)
	fmt.Println(obj.Count(8))
	fmt.Println(obj.Count(4))
	obj.Add(0, 1)
	obj.Add(1, 1)
	fmt.Println(obj.Count(7))
}
```

## 1867 — Orders With Maximum Quantity Above Average

```go
package main

// LeetCode #1867: Orders With Maximum Quantity Above Average
// https://leetcode.com/problems/orders-with-maximum-quantity-above-average/
// Difficulty: Medium [Paid]

import "fmt"

func main() {
	// Sample: order_id, quantity
	orders := [][]int{{1, 10}, {2, 5}, {3, 8}, {4, 3}, {5, 12}}
	fmt.Println(OrdersAboveAverage(orders))
}

// Time: O(n), Space: O(1)
func OrdersAboveAverage(orders [][]int) int {
	if len(orders) == 0 {
		return 0
	}
	sum := 0
	for _, o := range orders {
		sum += o[1]
	}
	avg := float64(sum) / float64(len(orders))
	count := 0
	for _, o := range orders {
		if float64(o[1]) > avg {
			count++
		}
	}
	// Find max quantity among orders above average
	maxQty := 0
	for _, o := range orders {
		if float64(o[1]) > avg && o[1] > maxQty {
			maxQty = o[1]
		}
	}
	return maxQty
}
```

## 1868 — Product Of Two Run Length Encoded Arrays

```go
package main

// LeetCode #1868: Product of Two Run-Length Encoded Arrays
// https://leetcode.com/problems/product-of-two-run-length-encoded-arrays/
// Difficulty: Medium [Paid]

import "fmt"

func main() {
	fmt.Println(FindRLEArray([][]int{{1, 3}, {2, 3}}, [][]int{{6, 3}, {3, 3}}))
	fmt.Println(FindRLEArray([][]int{{1, 2}, {2, 2}}, [][]int{{5, 1}, {2, 1}}))
}

// Time: O(n1 + n2), Space: O(n1 + n2) for result
func FindRLEArray(encoded1 [][]int, encoded2 [][]int) [][]int {
	result := make([][]int, 0)
	i, j := 0, 0

	for i < len(encoded1) && j < len(encoded2) {
		val := encoded1[i][0] * encoded2[j][0]
		minFreq := min(encoded1[i][1], encoded2[j][1])

		// Merge with previous if same value
		if len(result) > 0 && result[len(result)-1][0] == val {
			result[len(result)-1][1] += minFreq
		} else {
			result = append(result, []int{val, minFreq})
		}

		encoded1[i][1] -= minFreq
		encoded2[j][1] -= minFreq

		if encoded1[i][1] == 0 {
			i++
		}
		if encoded2[j][1] == 0 {
			j++
		}
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```

## 1870 — Minimum Speed To Arrive On Time

```go
package main

// LeetCode #1870: Minimum Speed to Arrive on Time
// https://leetcode.com/problems/minimum-speed-to-arrive-on-time/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MinSpeedOnTime([]int{1, 3, 2}, 6.0))
	fmt.Println(MinSpeedOnTime([]int{1, 3, 2}, 2.7))
	fmt.Println(MinSpeedOnTime([]int{1, 3, 2}, 1.9))
}

// Time: O(n log maxDist), Space: O(1)
func MinSpeedOnTime(dist []int, hour float64) int {
	n := len(dist)
	if hour <= float64(n-1) {
		return -1
	}

	lo, hi := 1, 10000000
	ans := -1

	for lo <= hi {
		mid := lo + (hi-lo)/2
		if canReach(dist, mid, hour) {
			ans = mid
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return ans
}

func canReach(dist []int, speed int, hour float64) bool {
	time := 0.0
	for i := 0; i < len(dist)-1; i++ {
		time += float64((dist[i] + speed - 1) / speed) // ceil division
	}
	time += float64(dist[len(dist)-1]) / float64(speed)
	return time <= hour
}
```

## 1871 — Jump Game Vii

```go
package main

// LeetCode #1871: Jump Game VII
// https://leetcode.com/problems/jump-game-vii/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(CanReach("011010", 2, 3))
	fmt.Println(CanReach("01101110", 2, 3))
	fmt.Println(CanReach("00", 1, 1))
}

// Time: O(n), Space: O(n)
func CanReach(s string, minJump int, maxJump int) bool {
	n := len(s)
	if s[n-1] != '0' {
		return false
	}

	dp := make([]bool, n)
	prefix := make([]int, n+1)
	dp[0] = true
	prefix[1] = 1

	for i := 1; i < n; i++ {
		if s[i] == '1' {
			prefix[i+1] = prefix[i]
			continue
		}
		left := max(0, i-maxJump)
		right := i - minJump
		if right >= left {
			reachable := prefix[right+1] - prefix[left] > 0
			if reachable {
				dp[i] = true
			}
		}
		prefix[i+1] = prefix[i]
		if dp[i] {
			prefix[i+1]++
		}
	}
	return dp[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

## 1874 — Minimize Product Sum Of Two Arrays

```go
package main

// LeetCode #1874: Minimize Product Sum of Two Arrays
// https://leetcode.com/problems/minimize-product-sum-of-two-arrays/
// Difficulty: Medium [Paid]

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MinProductSum([]int{5, 3, 4, 2}, []int{4, 2, 2, 5}))
	fmt.Println(MinProductSum([]int{2, 1, 4, 5, 7}, []int{3, 2, 4, 8, 6}))
}

// Time: O(n log n), Space: O(1)
func MinProductSum(nums1 []int, nums2 []int) int {
	sort.Ints(nums1)
	sort.Sort(sort.Reverse(sort.IntSlice(nums2)))
	sum := 0
	for i := 0; i < len(nums1); i++ {
		sum += nums1[i] * nums2[i]
	}
	return sum
}
```

## 1875 — Group Employees Of The Same Salary

```go
package main

// LeetCode #1875: Group Employees of the Same Salary
// https://leetcode.com/problems/group-employees-of-the-same-salary/
// Difficulty: Medium [Paid]

import (
	"fmt"
	"sort"
)

func main() {
	// employee: [id, salary]
	employees := [][]int{{1, 50000}, {2, 60000}, {3, 50000}, {4, 70000}, {5, 60000}}
	fmt.Println(GroupEmployees(employees))
}

// Time: O(n log n), Space: O(n)
func GroupEmployees(employees [][]int) [][]int {
	salaryMap := make(map[int][]int)
	for _, emp := range employees {
		id, salary := emp[0], emp[1]
		salaryMap[salary] = append(salaryMap[salary], id)
	}

	result := make([][]int, 0)
	for _, ids := range salaryMap {
		if len(ids) >= 2 {
			sort.Ints(ids)
			result = append(result, ids)
		}
	}

	// Sort by first employee ID for deterministic output
	sort.Slice(result, func(i, j int) bool {
		return result[i][0] < result[j][0]
	})
	return result
}
```

## 1877 — Minimize Maximum Pair Sum In Array

```go
package main

// LeetCode #1877: Minimize Maximum Pair Sum in Array
// https://leetcode.com/problems/minimize-maximum-pair-sum-in-array/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MinPairSum([]int{3, 5, 2, 3}))
	fmt.Println(MinPairSum([]int{3, 5, 4, 2, 4, 6}))
}

// Time: O(n log n), Space: O(1)
func MinPairSum(nums []int) int {
	sort.Ints(nums)
	maxSum := 0
	n := len(nums)
	for i := 0; i < n/2; i++ {
		pairSum := nums[i] + nums[n-1-i]
		if pairSum > maxSum {
			maxSum = pairSum
		}
	}
	return maxSum
}
```

## 1878 — Get Biggest Three Rhombus Sums In A Grid

```go
package main

// LeetCode #1878: Get Biggest Three Rhombus Sums in a Grid
// https://leetcode.com/problems/get-biggest-three-rhombus-sums-in-a-grid/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(GetBiggestThree([][]int{{3, 4, 5, 1, 3}, {3, 3, 4, 2, 3}, {20, 30, 200, 40, 10}, {1, 5, 5, 4, 1}, {4, 3, 2, 2, 5}}))
	fmt.Println(GetBiggestThree([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	fmt.Println(GetBiggestThree([][]int{{7, 7, 7}}))
}

// Time: O(m*n*min(m,n)), Space: O(1)
func GetBiggestThree(grid [][]int) []int {
	m, n := len(grid), len(grid[0])
	top3 := make([]int, 0)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// Single cell rhombus (size 0)
			addToTop3(&top3, grid[i][j])
			// Expand rhombus size
			maxSize := min(min(i, m-1-i), min(j, n-1-j))
			for s := 1; s <= maxSize; s++ {
				sum := 0
				// Traverse 4 edges of rhombus
				// Top to right: (i-k, j+k) for k=0..s
				// Right to bottom: (i+s-k, j+s-k) for k=0..s
				for k := 0; k < s; k++ {
					sum += grid[i-k][j+k]     // top-left to top-right
					sum += grid[i+k][j+s-k]   // top-right to bottom
				}
				for k := 0; k < s; k++ {
					sum += grid[i+s-k][j-k]   // bottom to bottom-left
					sum += grid[i-k][j-(s-k)] // bottom-left to top
				}
				// Corners counted twice, subtract once
				// Actually the 4 loops above need careful boundary handling
				// Let's fix: iterate each edge separately
				sum = 0
				// Top-right edge: (i-t, j+t) for t=0..s
				for t := 0; t <= s; t++ {
					sum += grid[i-t][j+t]
				}
				// Right-bottom edge: (i+s-t, j+s-t) for t=1..s
				for t := 1; t <= s; t++ {
					sum += grid[i+s-t][j+s-t]
				}
				// Bottom-left edge: (i+t, j-t) for t=1..s-1
				for t := 1; t < s; t++ {
					sum += grid[i-t][j-t]
				}
				// Left-top edge: (i-t, j-t) for t=1..s-1
				for t := 1; t < s; t++ {
					sum += grid[i+t][j-t]
				}
				addToTop3(&top3, sum)
			}
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(top3)))
	if len(top3) > 3 {
		top3 = top3[:3]
	}
	return top3
}

func addToTop3(top3 *[]int, val int) {
	for _, v := range *top3 {
		if v == val {
			return
		}
	}
	*top3 = append(*top3, val)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```

## 1881 — Maximum Value After Insertion

```go
package main

// LeetCode #1881: Maximum Value After Insertion
// https://leetcode.com/problems/maximum-value-after-insertion/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MaxValue("99", 9))
	fmt.Println(MaxValue("-13", 2))
	fmt.Println(MaxValue("73", 6))
}

// Time: O(n), Space: O(n)
func MaxValue(n string, x int) string {
	result := make([]byte, 0, len(n)+1)
	negative := n[0] == '-'

	if negative {
		result = append(result, '-')
		inserted := false
		for i := 1; i < len(n); i++ {
			digit := int(n[i] - '0')
			if !inserted && x < digit {
				result = append(result, byte(x+'0'))
				inserted = true
			}
			result = append(result, n[i])
		}
		if !inserted {
			result = append(result, byte(x+'0'))
		}
	} else {
		inserted := false
		for i := 0; i < len(n); i++ {
			digit := int(n[i] - '0')
			if !inserted && x > digit {
				result = append(result, byte(x+'0'))
				inserted = true
			}
			result = append(result, n[i])
		}
		if !inserted {
			result = append(result, byte(x+'0'))
		}
	}
	return string(result)
}
```

## 1882 — Process Tasks Using Servers

```go
package main

// LeetCode #1882: Process Tasks Using Servers
// https://leetcode.com/problems/process-tasks-using-servers/
// Difficulty: Medium

import (
	"container/heap"
	"fmt"
)

type Server struct {
	index int
	weight int
	freeTime int
}

type MinHeapAvailable []Server
func (h MinHeapAvailable) Len() int { return len(h) }
func (h MinHeapAvailable) Less(i, j int) bool {
	if h[i].weight != h[j].weight {
		return h[i].weight < h[j].weight
	}
	return h[i].index < h[j].index
}
func (h MinHeapAvailable) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *MinHeapAvailable) Push(x interface{}) { *h = append(*h, x.(Server)) }
func (h *MinHeapAvailable) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type MinHeapBusy []Server
func (h MinHeapBusy) Len() int { return len(h) }
func (h MinHeapBusy) Less(i, j int) bool {
	if h[i].freeTime != h[j].freeTime {
		return h[i].freeTime < h[j].freeTime
	}
	if h[i].weight != h[j].weight {
		return h[i].weight < h[j].weight
	}
	return h[i].index < h[j].index
}
func (h MinHeapBusy) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *MinHeapBusy) Push(x interface{}) { *h = append(*h, x.(Server)) }
func (h *MinHeapBusy) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func main() {
	fmt.Println(AssignTasks([]int{3, 3, 2}, []int{1, 2, 3, 2, 1, 2}))
	fmt.Println(AssignTasks([]int{5, 1, 4, 3, 2}, []int{2, 1, 2, 4, 5, 2, 1}))
}

// Time: O((m + n) log n) where m = len(tasks), n = len(servers)
// Space: O(n)
func AssignTasks(servers []int, tasks []int) []int {
	m := len(tasks)

	available := &MinHeapAvailable{}
	heap.Init(available)
	for i, w := range servers {
		heap.Push(available, Server{index: i, weight: w, freeTime: 0})
	}

	busy := &MinHeapBusy{}
	heap.Init(busy)

	result := make([]int, m)
	time := 0

	for j := 0; j < m; j++ {
		time = max(time, j)
		// Release completed servers
		for busy.Len() > 0 && (*busy)[0].freeTime <= time {
			s := heap.Pop(busy).(Server)
			heap.Push(available, s)
		}

		if available.Len() == 0 {
			// Jump to next available server's free time
			time = (*busy)[0].freeTime
			for busy.Len() > 0 && (*busy)[0].freeTime <= time {
				s := heap.Pop(busy).(Server)
				heap.Push(available, s)
			}
		}

		s := heap.Pop(available).(Server)
		result[j] = s.index
		s.freeTime = time + tasks[j]
		heap.Push(busy, s)
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

## 1884 — Egg Drop With 2 Eggs And N Floors

```go
package main

// LeetCode #1884: Egg Drop With 2 Eggs and N Floors
// https://leetcode.com/problems/egg-drop-with-2-eggs-and-n-floors/
// Difficulty: Medium

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(TwoEggDrop(2))
	fmt.Println(TwoEggDrop(100))
	fmt.Println(TwoEggDrop(10))
}

// Time: O(1), Space: O(1)
func TwoEggDrop(n int) int {
	// Solve x(x+1)/2 >= n
	// x = ceil((-1 + sqrt(1 + 8n)) / 2)
	return int(math.Ceil((-1 + math.Sqrt(1+8*float64(n))) / 2))
}
```

## 1885 — Count Pairs In Two Arrays

```go
package main

// LeetCode #1885: Count Pairs in Two Arrays
// https://leetcode.com/problems/count-pairs-in-two-arrays/
// Difficulty: Medium [Paid]

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(CountPairs([]int{1, 3, 4}, []int{1, 3, 4}))
	fmt.Println(CountPairs([]int{1, 2, 3}, []int{2, 3, 4}))
	fmt.Println(CountPairs([]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}))
}

// Time: O(n log n), Space: O(n)
func CountPairs(nums1 []int, nums2 []int) int {
	n := len(nums1)
	diff := make([]int, n)
	for i := 0; i < n; i++ {
		diff[i] = nums1[i] - nums2[i]
	}
	sort.Ints(diff)

	left, right := 0, n-1
	count := 0
	for left < right {
		if diff[left]+diff[right] > 0 {
			count += right - left
			right--
		} else {
			left++
		}
	}
	return count
}
```

## 1887 — Reduction Operations To Make The Array Elements Equal

```go
package main

// LeetCode #1887: Reduction Operations to Make the Array Elements Equal
// https://leetcode.com/problems/reduction-operations-to-make-the-array-elements-equal/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(ReductionOperations([]int{5, 1, 3}))
	fmt.Println(ReductionOperations([]int{1, 1, 1}))
	fmt.Println(ReductionOperations([]int{1, 1, 2, 2, 3}))
}

// Time: O(n log n), Space: O(1)
func ReductionOperations(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	ops := 0
	for i := 1; i < n; i++ {
		if nums[i] != nums[i-1] {
			ops += n - i
		}
	}
	return ops
}
```

## 1888 — Minimum Number Of Flips To Make The Binary String Alternating

```go
package main

// LeetCode #1888: Minimum Number of Flips to Make the Binary String Alternating
// https://leetcode.com/problems/minimum-number-of-flips-to-make-the-binary-string-alternating/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MinFlips("111000"))
	fmt.Println(MinFlips("010"))
	fmt.Println(MinFlips("1110"))
}

// Time: O(n), Space: O(n)
func MinFlips(s string) int {
	n := len(s)
	double := s + s
	ans := n

	// Compare against "01" pattern
	mismatch0, mismatch1 := 0, 0
	for i := 0; i < len(double); i++ {
		expected0 := byte('0')
		if i%2 == 1 {
			expected0 = '1'
		}
		expected1 := byte('1')
		if i%2 == 1 {
			expected1 = '0'
		}

		if double[i] != expected0 {
			mismatch0++
		}
		if double[i] != expected1 {
			mismatch1++
		}

		if i >= n {
			left := i - n
			if double[left] != byte('0')+byte((left)%2) {
				mismatch0--
			}
			if double[left] != byte('1')-byte((left)%2) {
				mismatch1--
			}
		}

		if i >= n-1 {
			if mismatch0 < ans {
				ans = mismatch0
			}
			if mismatch1 < ans {
				ans = mismatch1
			}
		}
	}
	return ans
}
```

## 1891 — Cutting Ribbons

```go
package main

// LeetCode #1891: Cutting Ribbons
// https://leetcode.com/problems/cutting-ribbons/
// Difficulty: Medium [Paid]

import "fmt"

func main() {
	fmt.Println(MaxLength([]int{9, 7, 5}, 3))
	fmt.Println(MaxLength([]int{7, 5, 9}, 4))
	fmt.Println(MaxLength([]int{5, 7, 9}, 22))
}

// Time: O(n log maxLen), Space: O(1)
func MaxLength(ribbons []int, k int) int {
	left, right := 1, 0
	for _, r := range ribbons {
		if r > right {
			right = r
		}
	}

	ans := 0
	for left <= right {
		mid := left + (right-left)/2
		count := 0
		for _, r := range ribbons {
			count += r / mid
		}
		if count >= k {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}
```

## 1894 — Find The Student That Will Replace The Chalk

```go
package main

// LeetCode #1894: Find the Student that Will Replace the Chalk
// https://leetcode.com/problems/find-the-student-that-will-replace-the-chalk/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(ChalkReplacer([]int{5, 1, 5}, 22))
	fmt.Println(ChalkReplacer([]int{3, 4, 1, 2}, 25))
	fmt.Println(ChalkReplacer([]int{5, 2, 3}, 9))
}

// Time: O(n), Space: O(1)
func ChalkReplacer(chalk []int, k int) int {
	sum := 0
	for _, c := range chalk {
		sum += c
	}
	k %= sum

	for i, c := range chalk {
		if k < c {
			return i
		}
		k -= c
	}
	return 0
}
```

## 1895 — Largest Magic Square

```go
package main

// LeetCode #1895: Largest Magic Square
// https://leetcode.com/problems/largest-magic-square/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(LargestMagicSquare([][]int{{7, 1, 4, 5, 6}, {2, 5, 1, 6, 4}, {1, 5, 4, 3, 2}, {1, 2, 7, 3, 4}}))
	fmt.Println(LargestMagicSquare([][]int{{5, 1, 3, 1}, {9, 3, 3, 1}, {1, 3, 3, 8}}))
}

// Time: O(m*n*min(m,n)^2), Space: O(m*n)
func LargestMagicSquare(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	// Build prefix sums for rows and columns
	rowSum := make([][]int, m+1)
	colSum := make([][]int, m+1)
	for i := range rowSum {
		rowSum[i] = make([]int, n+1)
		colSum[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			rowSum[i+1][j+1] = rowSum[i+1][j] + grid[i][j]
			colSum[i+1][j+1] = colSum[i][j+1] + grid[i][j]
		}
	}

	maxK := min(m, n)
	for k := maxK; k >= 1; k-- {
		for i := 0; i+k <= m; i++ {
			for j := 0; j+k <= n; j++ {
				if isMagic(grid, i, j, k, rowSum, colSum) {
					return k
				}
			}
		}
	}
	return 1
}

func isMagic(grid [][]int, r, c, k int, rowSum, colSum [][]int) bool {
	target := rowSum[r+1][c+k] - rowSum[r+1][c]

	// Check rows
	for i := 0; i < k; i++ {
		sum := rowSum[r+i+1][c+k] - rowSum[r+i+1][c]
		if sum != target {
			return false
		}
	}
	// Check columns
	for j := 0; j < k; j++ {
		sum := colSum[r+k][c+j+1] - colSum[r][c+j+1]
		if sum != target {
			return false
		}
	}
	// Check diagonals
	diag1, diag2 := 0, 0
	for i := 0; i < k; i++ {
		diag1 += grid[r+i][c+i]
		diag2 += grid[r+i][c+k-1-i]
	}
	return diag1 == target && diag2 == target
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```

## 1898 — Maximum Number Of Removable Characters

```go
package main

// LeetCode #1898: Maximum Number of Removable Characters
// https://leetcode.com/problems/maximum-number-of-removable-characters/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MaxRemovals("abcacb", "ab", []int{3, 1, 0}))
	fmt.Println(MaxRemovals("abcbddddd", "abcd", []int{3, 2, 1, 4, 5, 6}))
	fmt.Println(MaxRemovals("abcab", "abc", []int{0, 1, 2, 3, 4}))
}

// Time: O((n+m) log k) where n = len(s), m = len(p), k = len(removable)
// Space: O(n)
func MaxRemovals(s string, p string, removable []int) int {
	left, right := 0, len(removable)
	ans := 0

	for left <= right {
		mid := left + (right-left)/2
		if canForm(s, p, removable, mid) {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}

func canForm(s string, p string, removable []int, k int) bool {
	removed := make([]bool, len(s))
	for i := 0; i < k; i++ {
		removed[removable[i]] = true
	}

	j := 0
	for i := 0; i < len(s) && j < len(p); i++ {
		if !removed[i] && s[i] == p[j] {
			j++
		}
	}
	return j == len(p)
}
```

## 1899 — Merge Triplets To Form Target Triplet

```go
package main

// LeetCode #1899: Merge Triplets to Form Target Triplet
// https://leetcode.com/problems/merge-triplets-to-form-target-triplet/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MergeTriplets([][]int{{2, 5, 3}, {1, 8, 4}, {1, 7, 5}}, []int{2, 7, 5}))
	fmt.Println(MergeTriplets([][]int{{3, 4, 5}, {4, 5, 6}}, []int{3, 2, 5}))
	fmt.Println(MergeTriplets([][]int{{2, 5, 3}, {2, 3, 4}, {1, 2, 5}, {5, 2, 3}}, []int{5, 5, 5}))
}

// Time: O(n), Space: O(1)
func MergeTriplets(triplets [][]int, target []int) bool {
	found := [3]bool{}
	for _, t := range triplets {
		// Skip any triplet that exceeds target
		if t[0] > target[0] || t[1] > target[1] || t[2] > target[2] {
			continue
		}
		if t[0] == target[0] {
			found[0] = true
		}
		if t[1] == target[1] {
			found[1] = true
		}
		if t[2] == target[2] {
			found[2] = true
		}
	}
	return found[0] && found[1] && found[2]
}
```

## 1901 — Find A Peak Element Ii

```go
package main

// LeetCode #1901: Find a Peak Element II
// https://leetcode.com/problems/find-a-peak-element-ii/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(FindPeakGrid([][]int{{1, 4, 3, 2}, {2, 3, 4, 5}, {5, 4, 3, 6}}))
	fmt.Println(FindPeakGrid([][]int{{10, 20, 15}, {21, 30, 14}, {7, 16, 32}}))
}

// Time: O(m log n), Space: O(1)
func FindPeakGrid(mat [][]int) []int {
	m, n := len(mat), len(mat[0])
	left, right := 0, n-1

	for left <= right {
		mid := left + (right-left)/2
		maxRow := 0
		for i := 0; i < m; i++ {
			if mat[i][mid] > mat[maxRow][mid] {
				maxRow = i
			}
		}

		leftVal := -1
		if mid > 0 {
			leftVal = mat[maxRow][mid-1]
		}
		rightVal := -1
		if mid < n-1 {
			rightVal = mat[maxRow][mid+1]
		}

		if mat[maxRow][mid] > leftVal && mat[maxRow][mid] > rightVal {
			return []int{maxRow, mid}
		} else if mat[maxRow][mid] < leftVal {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return []int{-1, -1}
}
```

## 1902 — Depth Of Bst Given Insertion Order

```go
package main

// LeetCode #1902: Depth of BST Given Insertion Order
// https://leetcode.com/problems/depth-of-bst-given-insertion-order/
// Difficulty: Medium [Paid]

import "fmt"

func main() {
	fmt.Println(MaxDepthBST([]int{2, 1, 4, 3}))
	fmt.Println(MaxDepthBST([]int{2, 1, 3, 4}))
	fmt.Println(MaxDepthBST([]int{1, 2, 3, 4}))
}

// Time: O(n log n), Space: O(n)
func MaxDepthBST(order []int) int {
	// Use map to store depth of each value
	// Use ordered map simulation via lower/higher logic
	depth := make(map[int]int)
	depth[order[0]] = 1
	ans := 1

	// We need to find the nearest smaller and larger values already inserted
	// Since Go doesn't have an ordered map, we'll maintain a sorted slice
	sorted := []int{order[0]}

	for i := 1; i < len(order); i++ {
		v := order[i]
		// Find lower (floor) and higher (ceiling)
		lower, higher := -1, -1
		lo, hi := 0, len(sorted)-1
		for lo <= hi {
			mid := lo + (hi-lo)/2
			if sorted[mid] < v {
				lower = sorted[mid]
				lo = mid + 1
			} else {
				higher = sorted[mid]
				hi = mid - 1
			}
		}

		lowerDepth, higherDepth := 0, 0
		if lower != -1 {
			lowerDepth = depth[lower]
		}
		if higher != -1 {
			higherDepth = depth[higher]
		}

		curDepth := 1 + max(lowerDepth, higherDepth)
		depth[v] = curDepth
		if curDepth > ans {
			ans = curDepth
		}

		// Insert into sorted slice (maintain sorted order)
		pos := lo
		sorted = append(sorted[:pos], append([]int{v}, sorted[pos:]...)...)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

## 1904 — The Number Of Full Rounds You Have Played

```go
package main

// LeetCode #1904: The Number of Full Rounds You Have Played
// https://leetcode.com/problems/the-number-of-full-rounds-you-have-played/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(NumberOfRounds("09:31", "10:14"))
	fmt.Println(NumberOfRounds("21:30", "03:00"))
}

// Time: O(1), Space: O(1)
func NumberOfRounds(loginTime string, logoutTime string) int {
	t1 := toMinutes(loginTime)
	t2 := toMinutes(logoutTime)

	if t1 > t2 {
		t2 += 24 * 60
	}

	// First full round starts at ceil(t1/15)*15
	start := ((t1 + 14) / 15) * 15
	// Last full round ends at floor(t2/15)*15
	end := (t2 / 15) * 15

	if end < start {
		return 0
	}
	return (end - start) / 15
}

func toMinutes(time string) int {
	hours := int(time[0]-'0')*10 + int(time[1]-'0')
	mins := int(time[3]-'0')*10 + int(time[4]-'0')
	return hours*60 + mins
}
```

## 1905 — Count Sub Islands

```go
package main

// LeetCode #1905: Count Sub Islands
// https://leetcode.com/problems/count-sub-islands/
// Difficulty: Medium

import "fmt"

func main() {
	grid1 := [][]int{{1, 1, 1, 0, 0}, {0, 1, 1, 1, 1}, {0, 0, 0, 0, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 1, 1}}
	grid2 := [][]int{{1, 1, 1, 0, 0}, {0, 0, 1, 1, 1}, {0, 1, 0, 0, 0}, {1, 0, 1, 1, 0}, {0, 1, 0, 1, 0}}
	fmt.Println(CountSubIslands(grid1, grid2))

	grid1b := [][]int{{1, 0, 1, 0, 1}, {1, 1, 1, 1, 1}, {0, 0, 0, 0, 0}, {1, 1, 1, 1, 1}, {1, 0, 1, 0, 1}}
	grid2b := [][]int{{0, 0, 0, 0, 0}, {1, 1, 1, 1, 1}, {0, 1, 0, 1, 0}, {0, 1, 0, 1, 0}, {1, 0, 0, 0, 1}}
	fmt.Println(CountSubIslands(grid1b, grid2b))
}

// Time: O(m*n), Space: O(m*n) worst-case recursion
func CountSubIslands(grid1 [][]int, grid2 [][]int) int {
	m, n := len(grid2), len(grid2[0])
	count := 0

	var dfs func(i, j int) bool
	dfs = func(i, j int) bool {
		if i < 0 || i >= m || j < 0 || j >= n || grid2[i][j] == 0 {
			return true
		}
		grid2[i][j] = 0
		result := grid1[i][j] == 1
		result = dfs(i-1, j) && result
		result = dfs(i+1, j) && result
		result = dfs(i, j-1) && result
		result = dfs(i, j+1) && result
		return result
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 1 && dfs(i, j) {
				count++
			}
		}
	}
	return count
}
```

## 1906 — Minimum Absolute Difference Queries

```go
package main

// LeetCode #1906: Minimum Absolute Difference Queries
// https://leetcode.com/problems/minimum-absolute-difference-queries/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MinDifference([]int{1, 3, 4, 8}, [][]int{{0, 1}, {1, 2}, {2, 3}, {0, 3}}))
	fmt.Println(MinDifference([]int{4, 5, 2, 2, 7, 10}, [][]int{{2, 3}, {0, 2}, {0, 5}, {3, 5}}))
}

const maxVal = 100

// Time: O((n+q)*maxVal), Space: O(n*maxVal)
func MinDifference(nums []int, queries [][]int) []int {
	n := len(nums)
	// prefixCount[i][v] = count of value v in nums[0..i-1]
	prefixCount := make([][maxVal + 1]int, n+1)
	for i := 0; i < n; i++ {
		prefixCount[i+1] = prefixCount[i]
		prefixCount[i+1][nums[i]]++
	}

	result := make([]int, len(queries))
	for qIdx, q := range queries {
		l, r := q[0], q[1]
		prev := -1
		minDiff := -1
		for v := 1; v <= maxVal; v++ {
			if prefixCount[r+1][v]-prefixCount[l][v] > 0 {
				if prev != -1 {
					diff := v - prev
					if minDiff == -1 || diff < minDiff {
						minDiff = diff
					}
				}
				prev = v
			}
		}
		result[qIdx] = minDiff
	}
	return result
}
```

## 1907 — Count Salary Categories

```go
package main

// LeetCode #1907: Count Salary Categories
// https://leetcode.com/problems/count-salary-categories/
// Difficulty: Medium [Paid]

import "fmt"

func main() {
	// accounts: [account_id, income]
	accounts := [][]int{{1, 20000}, {2, 50000}, {3, 100000}, {4, 80000}, {5, 30000}}
	fmt.Println(CountSalaryCategories(accounts))
}

// Time: O(n), Space: O(1)
func CountSalaryCategories(accounts [][]int) []int {
	low := 0
	mid := 0
	high := 0

	for _, a := range accounts {
		income := a[1]
		if income < 20000 {
			low++
		} else if income <= 50000 {
			mid++
		} else {
			high++
		}
	}
	return []int{low, mid, high}
}
```

## 1908 — Game Of Nim

```go
package main

// LeetCode #1908: Game of Nim
// https://leetcode.com/problems/game-of-nim/
// Difficulty: Medium [Paid]

import "fmt"

func main() {
	fmt.Println(NimGame([]int{1, 2, 3}))
	fmt.Println(NimGame([]int{1, 1, 1}))
	fmt.Println(NimGame([]int{1, 2}))
}

// Time: O(n), Space: O(1)
func NimGame(piles []int) bool {
	xor := 0
	for _, p := range piles {
		xor ^= p
	}
	return xor != 0
}
```

## 1910 — Remove All Occurrences Of A Substring

```go
package main

// LeetCode #1910: Remove All Occurrences of a Substring
// https://leetcode.com/problems/remove-all-occurrences-of-a-substring/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(RemoveOccurrences("daabcbaabcbc", "abc"))
	fmt.Println(RemoveOccurrences("axxxxyyyyb", "xy"))
	fmt.Println(RemoveOccurrences("aabababa", "aba"))
}

// Time: O(n*m) where n = len(s), m = len(part), Space: O(n)
func RemoveOccurrences(s string, part string) string {
	stack := make([]byte, 0)
	m := len(part)
	for i := 0; i < len(s); i++ {
		stack = append(stack, s[i])
		if len(stack) >= m && string(stack[len(stack)-m:]) == part {
			stack = stack[:len(stack)-m]
		}
	}
	return string(stack)
}
```

## 1911 — Maximum Alternating Subsequence Sum

```go
package main

// LeetCode #1911: Maximum Alternating Subsequence Sum
// https://leetcode.com/problems/maximum-alternating-subsequence-sum/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MaxAlternatingSum([]int{4, 2, 5, 3}))
	fmt.Println(MaxAlternatingSum([]int{5, 6, 7, 8}))
	fmt.Println(MaxAlternatingSum([]int{6, 2, 1, 2, 4, 5}))
}

// Time: O(n), Space: O(1)
func MaxAlternatingSum(nums []int) int64 {
	even := int64(nums[0]) // max alternating sum ending with even index (added)
	odd := int64(0)         // max alternating sum ending with odd index (subtracted)

	for i := 1; i < len(nums); i++ {
		newEven := max64(even, max64(odd+int64(nums[i]), int64(nums[i])))
		odd = max64(odd, even-int64(nums[i]))
		even = newEven
	}
	return even
}

func max64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
```

## 1914 — Cyclically Rotating A Grid

```go
package main

// LeetCode #1914: Cyclically Rotating a Grid
// https://leetcode.com/problems/cyclically-rotating-a-grid/
// Difficulty: Medium

import "fmt"

func main() {
	grid := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	fmt.Println(RotateGrid(grid, 2))

	grid2 := [][]int{{10, 20}, {30, 40}}
	fmt.Println(RotateGrid(grid2, 1))
}

// Time: O(m*n), Space: O(m+n)
func RotateGrid(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	top, bottom := 0, m-1
	left, right := 0, n-1

	for top < bottom && left < right {
		layerLen := 2*(right-left+1) + 2*(bottom-top+1) - 4
		kMod := k % layerLen

		// Extract layer
		arr := make([]int, layerLen)
		idx := 0
		for j := left; j <= right; j++ {
			arr[idx] = grid[top][j]
			idx++
		}
		for i := top + 1; i <= bottom; i++ {
			arr[idx] = grid[i][right]
			idx++
		}
		for j := right - 1; j >= left; j-- {
			arr[idx] = grid[bottom][j]
			idx++
		}
		for i := bottom - 1; i > top; i-- {
			arr[idx] = grid[i][left]
			idx++
		}

		// Place rotated
		idx = 0
		for j := left; j <= right; j++ {
			grid[top][j] = arr[(idx+kMod)%layerLen]
			idx++
		}
		for i := top + 1; i <= bottom; i++ {
			grid[i][right] = arr[(idx+kMod)%layerLen]
			idx++
		}
		for j := right - 1; j >= left; j-- {
			grid[bottom][j] = arr[(idx+kMod)%layerLen]
			idx++
		}
		for i := bottom - 1; i > top; i-- {
			grid[i][left] = arr[(idx+kMod)%layerLen]
			idx++
		}

		top++
		bottom--
		left++
		right--
	}
	return grid
}
```

## 1915 — Number Of Wonderful Substrings

```go
package main

// LeetCode #1915: Number of Wonderful Substrings
// https://leetcode.com/problems/number-of-wonderful-substrings/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(WonderfulSubstrings("aba"))
	fmt.Println(WonderfulSubstrings("aabb"))
	fmt.Println(WonderfulSubstrings("he"))
}

// Time: O(n * 10) = O(n), Space: O(2^10) = O(1)
func WonderfulSubstrings(word string) int64 {
	// mask represents parity of each of 10 letters
	count := make([]int64, 1024) // 2^10 possible masks
	count[0] = 1
	mask := 0
	var result int64 = 0

	for _, c := range word {
		mask ^= 1 << (c - 'a')

		// Count substrings where all letters have even count
		result += count[mask]

		// Count substrings where exactly one letter has odd count
		for i := 0; i < 10; i++ {
			result += count[mask^(1<<i)]
		}

		count[mask]++
	}
	return result
}
```

## 1918 — Kth Smallest Subarray Sum

```go
package main

// LeetCode #1918: Kth Smallest Subarray Sum
// https://leetcode.com/problems/kth-smallest-subarray-sum/
// Difficulty: Medium [Paid]

import (
	"fmt"
)

func main() {
	fmt.Println(KthSmallestSubarraySum([]int{2, 1, 3}, 4))
	fmt.Println(KthSmallestSubarraySum([]int{3, 3, 3}, 4))
}

// Time: O(n log sum), Space: O(1)
func KthSmallestSubarraySum(nums []int, k int) int {
	n := len(nums)
	prefix := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + nums[i]
	}
	totalSum := prefix[n]

	// Binary search on sum value
	left, right := 0, totalSum
	for left < right {
		mid := left + (right-left)/2
		if countSubarraysLE(nums, prefix, mid) >= k {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func countSubarraysLE(nums []int, prefix []int, target int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		// Find first j where prefix[j+1]-prefix[i] > target
		lo, hi := i, len(nums)-1
		for lo <= hi {
			mid := lo + (hi-lo)/2
			sum := prefix[mid+1] - prefix[i]
			if sum <= target {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
		count += lo - i
	}
	return count
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```

## 1921 — Eliminate Maximum Number Of Monsters

```go
package main

// LeetCode #1921: Eliminate Maximum Number of Monsters
// https://leetcode.com/problems/eliminate-maximum-number-of-monsters/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(EliminateMaximum([]int{1, 3, 4}, []int{1, 1, 1}))
	fmt.Println(EliminateMaximum([]int{1, 1, 2, 3}, []int{1, 1, 1, 1}))
	fmt.Println(EliminateMaximum([]int{3, 2, 4}, []int{5, 3, 2}))
}

// Time: O(n log n), Space: O(n)
func EliminateMaximum(dist []int, speed []int) int {
	n := len(dist)
	time := make([]int, n)
	for i := 0; i < n; i++ {
		time[i] = (dist[i] + speed[i] - 1) / speed[i] // ceil division
	}
	sort.Ints(time)

	for i := 0; i < n; i++ {
		if time[i] <= i {
			return i
		}
	}
	return n
}
```

## 1922 — Count Good Numbers

```go
package main

// LeetCode #1922: Count Good Numbers
// https://leetcode.com/problems/count-good-numbers/
// Difficulty: Medium

import "fmt"

const mod1922 = 1000000007

func main() {
	fmt.Println(CountGoodNumbers(1))
	fmt.Println(CountGoodNumbers(4))
	fmt.Println(CountGoodNumbers(50))
}

// Time: O(log n), Space: O(1)
func CountGoodNumbers(n int64) int {
	evenPositions := (n + 1) / 2 // positions 0, 2, 4, ... (0-indexed)
	oddPositions := n / 2        // positions 1, 3, 5, ...

	// Even positions: 5 choices (0,2,4,6,8)
	// Odd positions: 4 choices (2,3,5,7)
	return int(powMod(5, evenPositions) * powMod(4, oddPositions) % mod1922)
}

func powMod(base int64, exp int64) int64 {
	result := int64(1)
	b := base % mod1922
	e := exp
	for e > 0 {
		if e&1 == 1 {
			result = (result * b) % mod1922
		}
		b = (b * b) % mod1922
		e >>= 1
	}
	return result
}
```

## 1926 — Nearest Exit From Entrance In Maze

```go
package main

// LeetCode #1926: Nearest Exit from Entrance in Maze
// https://leetcode.com/problems/nearest-exit-from-entrance-in-maze/
// Difficulty: Medium

import "fmt"

func main() {
	maze := [][]byte{
		{'+', '+', '.', '+'},
		{'.', '.', '.', '+'},
		{'+', '+', '+', '.'}}
	fmt.Println(NearestExit(maze, []int{1, 2}))

	maze2 := [][]byte{
		{'+', '+', '+'},
		{'.', '.', '.'},
		{'+', '+', '+'}}
	fmt.Println(NearestExit(maze2, []int{1, 0}))
}

// Time: O(m*n), Space: O(m*n)
func NearestExit(maze [][]byte, entrance []int) int {
	m, n := len(maze), len(maze[0])
	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	queue := [][2]int{{entrance[0], entrance[1]}}
	maze[entrance[0]][entrance[1]] = '+' // mark as visited
	steps := 0

	for len(queue) > 0 {
		size := len(queue)
		steps++
		for i := 0; i < size; i++ {
			r, c := queue[i][0], queue[i][1]
			for _, d := range dirs {
				nr, nc := r+d[0], c+d[1]
				if nr >= 0 && nr < m && nc >= 0 && nc < n && maze[nr][nc] == '.' {
					if nr == 0 || nr == m-1 || nc == 0 || nc == n-1 {
						return steps
					}
					maze[nr][nc] = '+'
					queue = append(queue, [2]int{nr, nc})
				}
			}
		}
		queue = queue[size:]
	}
	return -1
}
```

## 1927 — Sum Game

```go
package main

// LeetCode #1927: Sum Game
// https://leetcode.com/problems/sum-game/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(SumGame("5023"))
	fmt.Println(SumGame("25??"))
	fmt.Println(SumGame("?3295???"))
}

// Time: O(n), Space: O(1)
func SumGame(num string) bool {
	n := len(num)
	leftSum, rightSum := 0, 0
	leftQ, rightQ := 0, 0

	for i := 0; i < n/2; i++ {
		if num[i] == '?' {
			leftQ++
		} else {
			leftSum += int(num[i] - '0')
		}
	}
	for i := n / 2; i < n; i++ {
		if num[i] == '?' {
			rightQ++
		} else {
			rightSum += int(num[i] - '0')
		}
	}

	// Alice wants to avoid tie, Bob wants tie
	// "?" on left side favors Alice when she puts 9, etc.
	// The optimal strategy:
	// Bob will try to minimize the difference,
	// Alice will try to maximize it.

	// If total number of ? is odd, Alice can always win
	if (leftQ+rightQ)%2 == 1 {
		return true
	}

	// Each pair of '?' on opposite sides can cancel out (one puts 9, other puts 0)
	diff := leftSum - rightSum
	diff += (leftQ - rightQ) * 9 / 2

	return diff != 0
}
```

## 1930 — Unique Length 3 Palindromic Subsequences

```go
package main

// LeetCode #1930: Unique Length-3 Palindromic Subsequences
// https://leetcode.com/problems/unique-length-3-palindromic-subsequences/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(CountPalindromicSubsequence("aabca"))
	fmt.Println(CountPalindromicSubsequence("adc"))
	fmt.Println(CountPalindromicSubsequence("bbcbaba"))
}

// Time: O(n * 26) = O(n), Space: O(1)
func CountPalindromicSubsequence(s string) int {
	// For each character, find first and last occurrence
	first := make([]int, 26)
	last := make([]int, 26)
	for i := 0; i < 26; i++ {
		first[i] = -1
		last[i] = -1
	}
	for i := 0; i < len(s); i++ {
		idx := int(s[i] - 'a')
		if first[idx] == -1 {
			first[idx] = i
		}
		last[idx] = i
	}

	count := 0
	for c := 0; c < 26; c++ {
		if first[c] != -1 && last[c]-first[c] > 1 {
			// Count unique characters between first and last occurrence
			seen := make([]bool, 26)
			for i := first[c] + 1; i < last[c]; i++ {
				seen[s[i]-'a'] = true
			}
			for _, v := range seen {
				if v {
					count++
				}
			}
		}
	}
	return count
}
```

## 1934 — Confirmation Rate

```go
package main

// LeetCode #1934: Confirmation Rate
// https://leetcode.com/problems/confirmation-rate/
// Difficulty: Medium [Paid]

import "fmt"

func main() {
	// signups: [user_id, signup_date]
	// confirmations: [user_id, action, confirmation_date]
	signups := [][]int{{1, 1}, {2, 1}, {3, 1}}
	confirmations := [][]int{{1, 1, 1}, {1, 1, 2}, {2, 1, 1}, {2, 0, 2}, {3, 0, 1}}
	fmt.Println(ConfirmationRate(signups, confirmations))
}

// Time: O(n + m), Space: O(n)
func ConfirmationRate(signups [][]int, confirmations [][]int) []float64 {
	userMap := make(map[int][]int) // user_id -> [confirmed, total]
	for _, s := range signups {
		userMap[s[0]] = []int{0, 0}
	}
	for _, c := range confirmations {
		uid := c[0]
		if _, ok := userMap[uid]; ok {
			userMap[uid][1]++
			if c[1] == 1 { // confirmed
				userMap[uid][0]++
			}
		}
	}

	result := make([]float64, 0, len(signups))
	for _, s := range signups {
		uid := s[0]
		data := userMap[uid]
		if data[1] == 0 {
			result = append(result, 0.0)
		} else {
			result = append(result, float64(data[0])/float64(data[1]))
		}
	}
	return result
}
```

## 1936 — Add Minimum Number Of Rungs

```go
package main

// LeetCode #1936: Add Minimum Number of Rungs
// https://leetcode.com/problems/add-minimum-number-of-rungs/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(AddRungs([]int{1, 3, 5, 10}, 2))
	fmt.Println(AddRungs([]int{3, 6, 8, 10}, 3))
	fmt.Println(AddRungs([]int{3, 4, 6, 7}, 2))
}

// Time: O(n), Space: O(1)
func AddRungs(rungs []int, dist int) int {
	count := 0
	prev := 0
	for _, r := range rungs {
		gap := r - prev
		if gap > dist {
			count += (gap - 1) / dist
		}
		prev = r
	}
	return count
}
```

## 1937 — Maximum Number Of Points With Cost

```go
package main

// LeetCode #1937: Maximum Number of Points with Cost
// https://leetcode.com/problems/maximum-number-of-points-with-cost/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MaxPoints([][]int{{1, 2, 3}, {1, 5, 1}, {3, 1, 1}}))
	fmt.Println(MaxPoints([][]int{{1, 5}, {2, 3}, {4, 2}}))
}

// Time: O(m*n), Space: O(n)
func MaxPoints(points [][]int) int64 {
	m, n := len(points), len(points[0])
	dp := make([]int64, n)
	for j := 0; j < n; j++ {
		dp[j] = int64(points[0][j])
	}

	for i := 1; i < m; i++ {
		left := make([]int64, n)
		right := make([]int64, n)

		// Left to right: max of dp[k] + k for k <= j
		left[0] = dp[0]
		for j := 1; j < n; j++ {
			if dp[j]+int64(j) > left[j-1] {
				left[j] = dp[j] + int64(j)
			} else {
				left[j] = left[j-1]
			}
		}

		// Right to left: max of dp[k] - k for k >= j
		right[n-1] = dp[n-1] - int64(n-1)
		for j := n - 2; j >= 0; j-- {
			if dp[j]-int64(j) > right[j+1] {
				right[j] = dp[j] - int64(j)
			} else {
				right[j] = right[j+1]
			}
		}

		newDp := make([]int64, n)
		for j := 0; j < n; j++ {
			newDp[j] = int64(points[i][j]) + max64(left[j]-int64(j), right[j]+int64(j))
		}
		dp = newDp
	}

	ans := int64(0)
	for _, v := range dp {
		if v > ans {
			ans = v
		}
	}
	return ans
}

func max64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
```

## 1940 — Longest Common Subsequence Between Sorted Arrays

```go
package main

// LeetCode #1940: Longest Common Subsequence Between Sorted Arrays
// https://leetcode.com/problems/longest-common-subsequence-between-sorted-arrays/
// Difficulty: Medium [Paid]

import "fmt"

func main() {
	fmt.Println(LongestCommonSubsequenceBetweenSortedArrays([][]int{{1, 3, 4}, {1, 4, 7, 9}}))
	fmt.Println(LongestCommonSubsequenceBetweenSortedArrays([][]int{{2, 3, 6, 8}, {1, 2, 3, 5, 6, 7, 10}, {2, 3, 4, 6, 9}}))
}

// Time: O(total elements), Space: O(unique elements)
func LongestCommonSubsequenceBetweenSortedArrays(arrs [][]int) []int {
	// Since arrays are sorted, use frequency counting
	// Numbers appearing in ALL arrays are the answer
	freq := make(map[int]int)
	for _, arr := range arrs {
		for _, v := range arr {
			freq[v]++
		}
	}

	n := len(arrs)
	result := make([]int, 0)
	for _, v := range arrs[0] {
		if freq[v] == n {
			result = append(result, v)
		}
	}
	return result
}
```

## 1942 — The Number Of The Smallest Unoccupied Chair

```go
package main

// LeetCode #1942: The Number of the Smallest Unoccupied Chair
// https://leetcode.com/problems/the-number-of-the-smallest-unoccupied-chair/
// Difficulty: Medium

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(SmallestChair([][]int{{1, 4}, {2, 3}, {4, 6}}, 1))
	fmt.Println(SmallestChair([][]int{{3, 10}, {1, 5}, {2, 6}}, 0))
}

type ChairEvent struct {
	time     int
	isLeave  bool
	friend   int
	chairIdx int
}

type MinHeapInt []int

func (h MinHeapInt) Len() int           { return len(h) }
func (h MinHeapInt) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeapInt) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeapInt) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeapInt) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// Time: O(n log n), Space: O(n)
func SmallestChair(times [][]int, targetFriend int) int {
	n := len(times)

	// Pair friend index with their times
	type friend struct {
		arrival    int
		leaving    int
		friendIdx  int
	}
	friends := make([]friend, n)
	for i, t := range times {
		friends[i] = friend{t[0], t[1], i}
	}
	sort.Slice(friends, func(i, j int) bool {
		return friends[i].arrival < friends[j].arrival
	})

	// Min heap of available chairs
	available := &MinHeapInt{}
	heap.Init(available)
	for i := 0; i < n; i++ {
		heap.Push(available, i)
	}

	// Min heap of occupied chairs (sorted by leave time)
	type occupied struct {
		leaveTime int
		chair     int
	}
	occupiedHeap := make([]occupied, 0)

	chairOf := make(map[int]int) // friend -> chair

	for _, f := range friends {
		// Release chairs of friends who have left
		for len(occupiedHeap) > 0 && occupiedHeap[0].leaveTime <= f.arrival {
			o := occupiedHeap[0]
			occupiedHeap = occupiedHeap[1:]
			heap.Push(available, o.chair)
		}

		// Assign smallest available chair
		chair := heap.Pop(available).(int)
		chairOf[f.friendIdx] = chair

		// Insert into occupied (sorted by leave time)
		occupiedHeap = append(occupiedHeap, occupied{f.leaving, chair})
		sort.Slice(occupiedHeap, func(i, j int) bool {
			return occupiedHeap[i].leaveTime < occupiedHeap[j].leaveTime
		})
	}

	return chairOf[targetFriend]
}
```

## 1943 — Describe The Painting

```go
package main

// LeetCode #1943: Describe the Painting
// https://leetcode.com/problems/describe-the-painting/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(SplitPainting([][]int{{1, 4, 5}, {4, 7, 7}, {1, 7, 9}}))
	fmt.Println(SplitPainting([][]int{{1, 7, 9}, {6, 8, 15}, {8, 10, 7}}))
	fmt.Println(SplitPainting([][]int{{1, 4, 5}, {1, 4, 7}, {4, 7, 1}, {4, 7, 11}}))
}

// Time: O(n log n), Space: O(n)
func SplitPainting(segments [][]int) [][]int64 {
	diff := make(map[int]int64)
	endpoints := make(map[int]bool)

	for _, seg := range segments {
		start, end, color := seg[0], seg[1], seg[2]
		diff[start] += int64(color)
		diff[end] -= int64(color)
		endpoints[start] = true
		endpoints[end] = true
	}

	// Sort unique endpoints
	points := make([]int, 0, len(endpoints))
	for p := range endpoints {
		points = append(points, p)
	}
	sort.Ints(points)

	result := make([][]int64, 0)
	var sum int64 = 0
	for i := 0; i < len(points)-1; i++ {
		sum += diff[points[i]]
		if sum != 0 {
			result = append(result, []int64{int64(points[i]), int64(points[i+1]), sum})
		}
	}
	return result
}
```

## 1946 — Largest Number After Mutating Substring

```go
package main

// LeetCode #1946: Largest Number After Mutating Substring
// https://leetcode.com/problems/largest-number-after-mutating-substring/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MaximumNumber("132", []int{9, 8, 5, 0, 3, 6, 4, 2, 6, 8}))
	fmt.Println(MaximumNumber("021", []int{9, 4, 3, 5, 7, 2, 1, 9, 0, 6}))
	fmt.Println(MaximumNumber("5", []int{1, 4, 7, 5, 3, 2, 5, 6, 9, 4}))
}

// Time: O(n), Space: O(n)
func MaximumNumber(num string, change []int) string {
	n := len(num)
	result := make([]byte, n)
	i := 0

	// Skip prefix where change doesn't increase value
	for i < n && change[num[i]-'0'] <= int(num[i]-'0') {
		result[i] = num[i]
		i++
	}

	// Mutate while change increases or keeps same value
	for i < n && change[num[i]-'0'] >= int(num[i]-'0') {
		result[i] = byte(change[num[i]-'0'] + '0')
		i++
	}

	// Copy remaining
	for i < n {
		result[i] = num[i]
		i++
	}
	return string(result)
}
```

## 1947 — Maximum Compatibility Score Sum

```go
package main

// LeetCode #1947: Maximum Compatibility Score Sum
// https://leetcode.com/problems/maximum-compatibility-score-sum/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MaxCompatibilitySum([][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 1}}, [][]int{{1, 0, 0}, {0, 0, 1}, {1, 1, 0}}))
	fmt.Println(MaxCompatibilitySum([][]int{{0, 0}, {0, 0}, {0, 0}}, [][]int{{1, 1}, {1, 1}, {1, 1}}))
}

// Time: O(m! * n * k) where m = len(students) <= 8, Space: O(m)
func MaxCompatibilitySum(students [][]int, mentors [][]int) int {
	m := len(students)
	score := make([][]int, m)
	for i := 0; i < m; i++ {
		score[i] = make([]int, m)
		for j := 0; j < m; j++ {
			s := 0
			for k := 0; k < len(students[i]); k++ {
				if students[i][k] == mentors[j][k] {
					s++
				}
			}
			score[i][j] = s
		}
	}

	used := make([]bool, m)
	return backtrackMax(0, m, score, used)
}

func backtrackMax(student int, m int, score [][]int, used []bool) int {
	if student == m {
		return 0
	}
	maxScore := 0
	for mentor := 0; mentor < m; mentor++ {
		if !used[mentor] {
			used[mentor] = true
			cur := score[student][mentor] + backtrackMax(student+1, m, score, used)
			if cur > maxScore {
				maxScore = cur
			}
			used[mentor] = false
		}
	}
	return maxScore
}
```

## 1949 — Strong Friendship

```go
package main

// LeetCode #1949: Strong Friendship
// https://leetcode.com/problems/strong-friendship/
// Difficulty: Medium [Paid]

import "fmt"

func main() {
	// friendships: [user1_id, user2_id]
	friendships := [][]int{{1, 2}, {1, 3}, {2, 3}, {1, 4}, {2, 4}, {1, 5}}
	fmt.Println(StrongFriendship(friendships))
}

// Time: O(n^2), Space: O(n^2)
func StrongFriendship(friendships [][]int) int {
	friendSet := make(map[int]map[int]bool)
	for _, f := range friendships {
		a, b := f[0], f[1]
		if friendSet[a] == nil {
			friendSet[a] = make(map[int]bool)
		}
		if friendSet[b] == nil {
			friendSet[b] = make(map[int]bool)
		}
		friendSet[a][b] = true
		friendSet[b][a] = true
	}

	count := 0
	// For each pair of users, check if they have at least 3 common friends
	users := make([]int, 0, len(friendSet))
	for u := range friendSet {
		users = append(users, u)
	}

	for i := 0; i < len(users); i++ {
		for j := i + 1; j < len(users); j++ {
			a, b := users[i], users[j]
			if friendSet[a][b] { // they are friends
				continue
			}
			common := 0
			for f := range friendSet[a] {
				if friendSet[b][f] {
					common++
				}
			}
			if common >= 3 {
				count++
			}
		}
	}
	return count
}
```

## 1950 — Maximum Of Minimum Values In All Subarrays

```go
package main

// LeetCode #1950: Maximum of Minimum Values in All Subarrays
// https://leetcode.com/problems/maximum-of-minimum-values-in-all-subarrays/
// Difficulty: Medium [Paid]

import "fmt"

func main() {
	fmt.Println(MaxOfMinValues([]int{10, 20, 30, 50, 10, 70, 30}))
	fmt.Println(MaxOfMinValues([]int{1, 2, 3, 4, 5}))
}

// Time: O(n), Space: O(n)
func MaxOfMinValues(nums []int) []int {
	n := len(nums)
	result := make([]int, n+1) // result[k] for k-length subarrays
	for i := range result {
		result[i] = 0
	}

	// Find previous smaller and next smaller elements
	left := make([]int, n)
	right := make([]int, n)
	for i := 0; i < n; i++ {
		left[i] = -1
		right[i] = n
	}

	stack := make([]int, 0)
	for i := 0; i < n; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] >= nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	stack = make([]int, 0)
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && nums[stack[len(stack)-1]] >= nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			right[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	for i := 0; i < n; i++ {
		length := right[i] - left[i] - 1
		if nums[i] > result[length] {
			result[length] = nums[i]
		}
	}

	// Fill missing values: max of min values for larger windows
	for i := n - 1; i >= 1; i-- {
		if result[i] < result[i+1] {
			result[i] = result[i+1]
		}
	}
	return result[1:]
}
```

## 1951 — All The Pairs With The Maximum Number Of Common Followers

```go
package main

// LeetCode #1951: All the Pairs With the Maximum Number of Common Followers
// https://leetcode.com/problems/all-the-pairs-with-the-maximum-number-of-common-followers/
// Difficulty: Medium [Paid]

import (
	"fmt"
	"sort"
)

func main() {
	// relations: [user_id, follower_id]
	relations := [][]int{{1, 3}, {2, 3}, {3, 4}, {1, 4}, {2, 4}, {1, 5}}
	fmt.Println(MaxCommonFollowers(relations))
}

// Time: O(n^2 * m) roughly, Space: O(n*m)
func MaxCommonFollowers(relations [][]int) [][]int {
	followers := make(map[int]map[int]bool)
	for _, r := range relations {
		user, follower := r[0], r[1]
		if followers[user] == nil {
			followers[user] = make(map[int]bool)
		}
		followers[user][follower] = true
	}

	users := make([]int, 0, len(followers))
	for u := range followers {
		users = append(users, u)
	}
	sort.Ints(users)

	maxCommon := 0
	result := make([][]int, 0)

	for i := 0; i < len(users); i++ {
		for j := i + 1; j < len(users); j++ {
			a, b := users[i], users[j]
			common := 0
			for f := range followers[a] {
				if followers[b][f] {
					common++
				}
			}
			if common > maxCommon {
				maxCommon = common
				result = [][]int{{a, b}}
			} else if common == maxCommon && common > 0 {
				result = append(result, []int{a, b})
			}
		}
	}
	return result
}
```

## 1953 — Maximum Number Of Weeks For Which You Can Work

```go
package main

// LeetCode #1953: Maximum Number of Weeks for Which You Can Work
// https://leetcode.com/problems/maximum-number-of-weeks-for-which-you-can-work/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(NumberOfWeeks([]int{1, 2, 3}))
	fmt.Println(NumberOfWeeks([]int{5, 2, 1}))
	fmt.Println(NumberOfWeeks([]int{9, 3, 6, 8, 2, 1}))
}

// Time: O(n), Space: O(1)
func NumberOfWeeks(milestones []int) int64 {
	var sum int64 = 0
	maxVal := 0
	for _, m := range milestones {
		sum += int64(m)
		if m > maxVal {
			maxVal = m
		}
	}

	rest := sum - int64(maxVal)
	if int64(maxVal) > rest+1 {
		return 2*rest + 1
	}
	return sum
}
```

## 1954 — Minimum Garden Perimeter To Collect Enough Apples

```go
package main

// LeetCode #1954: Minimum Garden Perimeter to Collect Enough Apples
// https://leetcode.com/problems/minimum-garden-perimeter-to-collect-enough-apples/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MinimumPerimeter(1))
	fmt.Println(MinimumPerimeter(13))
	fmt.Println(MinimumPerimeter(1000000000))
}

// Time: O(cuberoot(n)), Space: O(1)
func MinimumPerimeter(neededApples int64) int64 {
	// For a garden with side length 2n (total apples = 2n(n+1)(2n+1))
	// Apples = 2 * n * (n+1) * (2n+1)
	// Perimeter = 8 * n

	n := int64(1)
	for {
		apples := 2 * n * (n + 1) * (2*n + 1)
		if apples >= neededApples {
			return 8 * n
		}
		n++
	}
}
```

## 1958 — Check If Move Is Legal

```go
package main

// LeetCode #1958: Check if Move is Legal
// https://leetcode.com/problems/check-if-move-is-legal/
// Difficulty: Medium

import "fmt"

var dirs = [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

func main() {
	board := [][]byte{
		{'.', '.', '.', 'B', '.', '.', '.', '.'},
		{'.', '.', '.', 'W', '.', '.', '.', '.'},
		{'.', '.', '.', 'W', '.', '.', '.', '.'},
		{'.', '.', '.', 'W', '.', '.', '.', '.'},
		{'W', 'B', 'B', '.', 'W', 'W', 'W', 'B'},
		{'.', '.', '.', 'B', '.', '.', '.', '.'},
		{'.', '.', '.', 'B', '.', '.', '.', '.'},
		{'.', '.', '.', 'W', '.', '.', '.', '.'}}
	fmt.Println(CheckMove(board, 4, 3, 'B'))

	board2 := [][]byte{
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', 'B', '.', '.', 'W', '.', '.', '.'},
		{'.', '.', 'W', '.', '.', '.', '.', '.'},
		{'.', '.', '.', 'W', '.', '.', '.', '.'},
		{'.', 'W', '.', '.', 'W', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'}}
	fmt.Println(CheckMove(board2, 4, 1, 'B'))
}

// Time: O(1) since board is always 8x8, Space: O(1)
func CheckMove(board [][]byte, rMove int, cMove int, color byte) bool {
	opponent := byte('B')
	if color == 'B' {
		opponent = 'W'
	}

	for _, d := range dirs {
		dr, dc := d[0], d[1]
		r, c := rMove+dr, cMove+dc
		steps := 0

		for r >= 0 && r < 8 && c >= 0 && c < 8 && board[r][c] == opponent {
			r += dr
			c += dc
			steps++
		}

		if steps > 0 && r >= 0 && r < 8 && c >= 0 && c < 8 && board[r][c] == color {
			return true
		}
	}
	return false
}
```

## 1959 — Minimum Total Space Wasted With K Resizing Operations

```go
package main

// LeetCode #1959: Minimum Total Space Wasted With K Resizing Operations
// https://leetcode.com/problems/minimum-total-space-wasted-with-k-resizing-operations/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MinSpaceWastedKResizing([]int{10, 20, 30}, 1))
	fmt.Println(MinSpaceWastedKResizing([]int{10, 20, 15, 30, 20}, 2))
}

const INF = 1 << 30

// Time: O(k * n^2), Space: O(k * n)
func MinSpaceWastedKResizing(nums []int, k int) int {
	n := len(nums)
	k++ // k resizes = k+1 segments

	// g[i][j] = wasted space for segment nums[i..j]
	g := make([][]int, n)
	for i := 0; i < n; i++ {
		g[i] = make([]int, n)
		s, mx := 0, 0
		for j := i; j < n; j++ {
			s += nums[j]
			if nums[j] > mx {
				mx = nums[j]
			}
			g[i][j] = mx*(j-i+1) - s
		}
	}

	// dp[i][j] = min waste for first i elements with j segments
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
		for j := range dp[i] {
			dp[i][j] = INF
		}
	}
	dp[0][0] = 0

	for i := 1; i <= n; i++ {
		for j := 1; j <= k; j++ {
			for h := 0; h < i; h++ {
				val := dp[h][j-1] + g[h][i-1]
				if val < dp[i][j] {
					dp[i][j] = val
				}
			}
		}
	}
	return dp[n][k]
}
```

## 1962 — Remove Stones To Minimize The Total

```go
package main

// LeetCode #1962: Remove Stones to Minimize the Total
// https://leetcode.com/problems/remove-stones-to-minimize-the-total/
// Difficulty: Medium

import (
	"container/heap"
	"fmt"
)

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func main() {
	fmt.Println(MinStoneSum([]int{5, 4, 9}, 2))
	fmt.Println(MinStoneSum([]int{4, 3, 6, 7}, 3))
}

// Time: O(n + k log n), Space: O(n)
func MinStoneSum(piles []int, k int) int {
	h := &MaxHeap{}
	heap.Init(h)
	sum := 0
	for _, p := range piles {
		sum += p
		heap.Push(h, p)
	}

	for i := 0; i < k; i++ {
		cur := heap.Pop(h).(int)
		removed := cur / 2
		sum -= removed
		heap.Push(h, cur-removed)
	}
	return sum
}
```

## 1963 — Minimum Number Of Swaps To Make The String Balanced

```go
package main

// LeetCode #1963: Minimum Number of Swaps to Make the String Balanced
// https://leetcode.com/problems/minimum-number-of-swaps-to-make-the-string-balanced/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MinSwapsBalanced("][]["))
	fmt.Println(MinSwapsBalanced("]]][[["))
	fmt.Println(MinSwapsBalanced("[]"))
}

// Time: O(n), Space: O(1)
func MinSwapsBalanced(s string) int {
	unmatched := 0
	maxUnmatched := 0
	for _, c := range s {
		if c == '[' {
			unmatched--
		} else {
			unmatched++
		}
		if unmatched > maxUnmatched {
			maxUnmatched = unmatched
		}
	}
	return (maxUnmatched + 1) / 2
}
```

## 1966 — Binary Searchable Numbers In An Unsorted Array

```go
package main

// LeetCode #1966: Binary Searchable Numbers in an Unsorted Array
// https://leetcode.com/problems/binary-searchable-numbers-in-an-unsorted-array/
// Difficulty: Medium [Paid]

import "fmt"

func main() {
	fmt.Println(BinarySearchableNumbers([]int{2, 1, 3, 5, 4, 6}))
	fmt.Println(BinarySearchableNumbers([]int{1, 3, 2}))
	fmt.Println(BinarySearchableNumbers([]int{2, 3, 1}))
}

// Time: O(n), Space: O(n)
func BinarySearchableNumbers(nums []int) int {
	n := len(nums)
	prefixMax := make([]int, n)
	suffixMin := make([]int, n)

	prefixMax[0] = nums[0]
	for i := 1; i < n; i++ {
		if nums[i] > prefixMax[i-1] {
			prefixMax[i] = nums[i]
		} else {
			prefixMax[i] = prefixMax[i-1]
		}
	}

	suffixMin[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		if nums[i] < suffixMin[i+1] {
			suffixMin[i] = nums[i]
		} else {
			suffixMin[i] = suffixMin[i+1]
		}
	}

	count := 0
	for i := 0; i < n; i++ {
		if prefixMax[i] == suffixMin[i] {
			count++
		}
	}
	return count
}
```

## 1968 — Array With Elements Not Equal To Average Of Neighbors

```go
package main

// LeetCode #1968: Array With Elements Not Equal to Average of Neighbors
// https://leetcode.com/problems/array-with-elements-not-equal-to-average-of-neighbors/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(RearrangeArray([]int{1, 2, 3, 4, 5}))
	fmt.Println(RearrangeArray([]int{6, 2, 0, 9, 7}))
}

// Time: O(n log n), Space: O(n)
func RearrangeArray(nums []int) []int {
	sort.Ints(nums)
	n := len(nums)
	result := make([]int, n)
	left, right := 0, n-1
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			result[i] = nums[left]
			left++
		} else {
			result[i] = nums[right]
			right--
		}
	}
	return result
}
```

## 1969 — Minimum Non Zero Product Of The Array Elements

```go
package main

// LeetCode #1969: Minimum Non-Zero Product of the Array Elements
// https://leetcode.com/problems/minimum-non-zero-product-of-the-array-elements/
// Difficulty: Medium

import "fmt"

const mod1969 = 1000000007

func main() {
	fmt.Println(MinNonZeroProduct(1))
	fmt.Println(MinNonZeroProduct(2))
	fmt.Println(MinNonZeroProduct(3))
}

// Time: O(p), Space: O(1)
func MinNonZeroProduct(p int) int {
	maxVal := (int64(1) << uint(p)) - 1
	base := maxVal - 1
	exp := (int64(1) << uint(p-1)) - 1
	result := int(maxVal % mod1969)
	result = int(int64(result) * powMod1969(base%int64(mod1969), exp) % mod1969)
	return result
}

func powMod1969(base int64, exp int64) int64 {
	result := int64(1)
	b := base % int64(mod1969)
	e := exp
	for e > 0 {
		if e&1 == 1 {
			result = (result * b) % int64(mod1969)
		}
		b = (b * b) % int64(mod1969)
		e >>= 1
	}
	return result
}
```

## 1973 — Count Nodes Equal To Sum Of Descendants

```go
package main

// LeetCode #1973: Count Nodes Equal to Sum of Descendants
// https://leetcode.com/problems/count-nodes-equal-to-sum-of-descendants/
// Difficulty: Medium [Paid]

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 10,
		Left:  &TreeNode{Val: 3,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 1},
		},
		Right: &TreeNode{Val: 5}}
	fmt.Println(EqualToDescendants(root))

	root2 := &TreeNode{Val: 2,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 1}}
	fmt.Println(EqualToDescendants(root2))
}

// Time: O(n), Space: O(h) where h is tree height
func EqualToDescendants(root *TreeNode) int {
	count := 0
	dfs(root, &count)
	return count
}

func dfs(node *TreeNode, count *int) int64 {
	if node == nil {
		return 0
	}
	leftSum := dfs(node.Left, count)
	rightSum := dfs(node.Right, count)
	if leftSum+rightSum == int64(node.Val) {
		*count++
	}
	return leftSum + rightSum + int64(node.Val)
}
```

## 1975 — Maximum Matrix Sum

```go
package main

// LeetCode #1975: Maximum Matrix Sum
// https://leetcode.com/problems/maximum-matrix-sum/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MaxMatrixSum([][]int{{1, -1}, {-1, 1}}))
	fmt.Println(MaxMatrixSum([][]int{{1, 2, 3}, {-1, -2, -3}, {1, 2, 3}}))
}

// Time: O(m*n), Space: O(1)
func MaxMatrixSum(matrix [][]int) int64 {
	total := int64(0)
	negCount := 0
	minAbs := int64(1 << 31)

	for _, row := range matrix {
		for _, val := range row {
			if val < 0 {
				negCount++
			}
			abs := int64(val)
			if abs < 0 {
				abs = -abs
			}
			total += abs
			if abs < minAbs {
				minAbs = abs
			}
		}
	}

	if negCount%2 == 1 {
		total -= 2 * minAbs
	}
	return total
}
```

## 1976 — Number Of Ways To Arrive At Destination

```go
package main

// LeetCode #1976: Number of Ways to Arrive at Destination
// https://leetcode.com/problems/number-of-ways-to-arrive-at-destination/
// Difficulty: Medium

import (
	"container/heap"
	"fmt"
)

type Edge struct {
	node int
	time int
}
type Item struct {
	node int
	dist int64
}
type PriorityQueue []Item

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].dist < pq[j].dist }
func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(Item)) }
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[:n-1]
	return x
}

const mod1976 = 1000000007

func main() {
	fmt.Println(CountPaths(7, [][]int{{0, 6, 7}, {0, 1, 2}, {1, 2, 3}, {1, 3, 3}, {6, 3, 3}, {3, 5, 1}, {6, 5, 1}, {2, 5, 1}, {0, 4, 5}, {4, 6, 2}}))
	fmt.Println(CountPaths(2, [][]int{{1, 0, 10}}))
}

// Time: O(E log V), Space: O(V + E)
func CountPaths(n int, roads [][]int) int {
	graph := make([][]Edge, n)
	for _, r := range roads {
		u, v, t := r[0], r[1], r[2]
		graph[u] = append(graph[u], Edge{v, t})
		graph[v] = append(graph[v], Edge{u, t})
	}

	dist := make([]int64, n)
	ways := make([]int, n)
	for i := range dist {
		dist[i] = 1 << 62
	}
	dist[0] = 0
	ways[0] = 1

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, Item{0, 0})

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(Item)
		if cur.dist > dist[cur.node] {
			continue
		}
		for _, e := range graph[cur.node] {
			newDist := cur.dist + int64(e.time)
			if newDist < dist[e.node] {
				dist[e.node] = newDist
				ways[e.node] = ways[cur.node]
				heap.Push(pq, Item{e.node, newDist})
			} else if newDist == dist[e.node] {
				ways[e.node] = (ways[e.node] + ways[cur.node]) % mod1976
			}
		}
	}
	return ways[n-1]
}
```

## 1980 — Find Unique Binary String

```go
package main

// LeetCode #1980: Find Unique Binary String
// https://leetcode.com/problems/find-unique-binary-string/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(FindDifferentBinaryString([]string{"01", "10"}))
	fmt.Println(FindDifferentBinaryString([]string{"00", "01"}))
	fmt.Println(FindDifferentBinaryString([]string{"111", "011", "001"}))
}

// Time: O(n^2), Space: O(n)
func FindDifferentBinaryString(nums []string) string {
	n := len(nums)
	set := make(map[string]bool)
	for _, s := range nums {
		set[s] = true
	}

	// Generate candidates using Cantor diagonal argument
	result := make([]byte, n)
	for i := 0; i < n; i++ {
		result[i] = '0'
	}
	for {
		s := string(result)
		if !set[s] {
			return s
		}
		// Increment binary string
		j := n - 1
		for j >= 0 && result[j] == '1' {
			result[j] = '0'
			j--
		}
		if j < 0 {
			break
		}
		result[j] = '1'
	}
	return ""
}
```

## 1981 — Minimize The Difference Between Target And Chosen Elements

```go
package main

// LeetCode #1981: Minimize the Difference Between Target and Chosen Elements
// https://leetcode.com/problems/minimize-the-difference-between-target-and-chosen-elements/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MinimizeTheDifference([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 13))
	fmt.Println(MinimizeTheDifference([][]int{{1}, {2}, {3}}, 100))
	fmt.Println(MinimizeTheDifference([][]int{{1, 2, 9, 8, 7}}, 6))
}

// Time: O(m * n * maxSum), Space: O(maxSum) where maxSum = 70*70 = 4900
func MinimizeTheDifference(mat [][]int, target int) int {
	m, n := len(mat), len(mat[0])
	possible := make([]bool, 4901)
	possible[0] = true

	for i := 0; i < m; i++ {
		next := make([]bool, 4901)
		for s := 0; s < len(possible); s++ {
			if possible[s] {
				for j := 0; j < n; j++ {
					if s+mat[i][j] < len(next) {
						next[s+mat[i][j]] = true
					}
				}
			}
		}
		possible = next
	}

	ans := 1 << 30
	for s := 0; s < len(possible); s++ {
		if possible[s] {
			diff := s - target
			if diff < 0 {
				diff = -diff
			}
			if diff < ans {
				ans = diff
			}
		}
	}
	return ans
}
```

## 1983 — Widest Pair Of Indices With Equal Range Sum

```go
package main

// LeetCode #1983: Widest Pair of Indices With Equal Range Sum
// https://leetcode.com/problems/widest-pair-of-indices-with-equal-range-sum/
// Difficulty: Medium [Paid]

import "fmt"

func main() {
	fmt.Println(WidestPairOfIndicesWithEqualRangeSum([]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 2, 3}))
	fmt.Println(WidestPairOfIndicesWithEqualRangeSum([]int{1, 1, 1}, []int{1, 1, 1}))
	fmt.Println(WidestPairOfIndicesWithEqualRangeSum([]int{0, 1}, []int{1, 0}))
}

// Time: O(n), Space: O(n)
func WidestPairOfIndicesWithEqualRangeSum(nums1 []int, nums2 []int) int {
	first := make(map[int]int)
	first[0] = -1
	maxWidth := 0
	prefix1, prefix2 := 0, 0

	for i := 0; i < len(nums1); i++ {
		prefix1 += nums1[i]
		prefix2 += nums2[i]
		diff := prefix1 - prefix2

		if idx, ok := first[diff]; ok {
			if i-idx > maxWidth {
				maxWidth = i - idx
			}
		} else {
			first[diff] = i
		}
	}

	return maxWidth
}
```

## 1985 — Find The Kth Largest Integer In The Array

```go
package main

// LeetCode #1985: Find the Kth Largest Integer in the Array
// https://leetcode.com/problems/find-the-kth-largest-integer-in-the-array/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(FindTheKthLargestIntegerInTheArray([]string{"3", "6", "7", "10"}, 4))
	fmt.Println(FindTheKthLargestIntegerInTheArray([]string{"2", "21", "12", "1"}, 3))
	fmt.Println(FindTheKthLargestIntegerInTheArray([]string{"0", "0"}, 2))
}

// Time: O(n log n), Space: O(1) (ignoring sort space)
func FindTheKthLargestIntegerInTheArray(nums []string, k int) string {
	sort.Slice(nums, func(i, j int) bool {
		if len(nums[i]) != len(nums[j]) {
			return len(nums[i]) > len(nums[j])
		}
		return nums[i] > nums[j]
	})
	return nums[k-1]
}
```

## 1986 — Minimum Number Of Work Sessions To Finish The Tasks

```go
package main

// LeetCode #1986: Minimum Number of Work Sessions to Finish the Tasks
// https://leetcode.com/problems/minimum-number-of-work-sessions-to-finish-the-tasks/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MinimumNumberOfWorkSessionsToFinishTheTasks([]int{1, 2, 3}, 3))
	fmt.Println(MinimumNumberOfWorkSessionsToFinishTheTasks([]int{3, 1, 3, 1, 1}, 8))
	fmt.Println(MinimumNumberOfWorkSessionsToFinishTheTasks([]int{1, 2, 3, 4, 5}, 15))
}

// Time: O(2^n * n), Space: O(2^n)
func MinimumNumberOfWorkSessionsToFinishTheTasks(tasks []int, sessionTime int) int {
	n := len(tasks)
	m := 1 << n
	dp := make([]int, m)
	sessions := make([]int, m)

	for i := range dp {
		dp[i] = n + 1
		sessions[i] = sessionTime + 1
	}
	dp[0] = 0
	sessions[0] = 0

	for mask := 0; mask < m; mask++ {
		for i := 0; i < n; i++ {
			if mask&(1<<i) == 0 {
				newMask := mask | (1 << i)
				if sessions[mask]+tasks[i] <= sessionTime {
					if dp[newMask] > dp[mask] || (dp[newMask] == dp[mask] && sessions[newMask] > sessions[mask]+tasks[i]) {
						dp[newMask] = dp[mask]
						sessions[newMask] = sessions[mask] + tasks[i]
					}
				} else {
					if dp[newMask] > dp[mask]+1 || (dp[newMask] == dp[mask]+1 && sessions[newMask] > tasks[i]) {
						dp[newMask] = dp[mask] + 1
						sessions[newMask] = tasks[i]
					}
				}
			}
		}
	}

	return dp[m-1] + 1
}
```

## 1988 — Find Cutoff Score For Each School

```go
package main

// LeetCode #1988: Find Cutoff Score for Each School
// https://leetcode.com/problems/find-cutoff-score-for-each-school/
// Difficulty: Medium (SQL problem — simulated in Go)
//
// Simulates: For each school, find the minimum exam score such that
// the school's capacity >= student_count (school can accept all students
// from that score group). If no such score exists, output -1.
// SQL equivalent: LEFT JOIN Exam ON Schools.capacity >= Exam.student_count
// then MIN(score) per school.

import (
	"fmt"
	"sort"
)

// School represents the Schools database table.
type School struct {
	SchoolID int
	Capacity int
}

// Exam represents the Exam database table.
type Exam struct {
	Score        int
	StudentCount int
}

// SchoolResult holds the output.
type SchoolResult struct {
	SchoolID int
	Score    int
}

// findCutoffScore simulates the SQL query.
// Time: O(s * e) | Space: O(s)
// s = number of schools, e = number of exam rows.
// SQL equivalent:
//   SELECT school_id, MIN(IFNULL(score, -1)) AS score
//   FROM Schools LEFT JOIN Exam ON Schools.capacity >= Exam.student_count
//   GROUP BY school_id
func findCutoffScore(schools []School, exams []Exam) []SchoolResult {
	var results []SchoolResult

	for _, school := range schools {
		bestScore := -1
		for _, exam := range exams {
			if school.Capacity >= exam.StudentCount {
				if bestScore == -1 || exam.Score < bestScore {
					bestScore = exam.Score
				}
			}
		}
		results = append(results, SchoolResult{SchoolID: school.SchoolID, Score: bestScore})
	}

	// Order by school_id.
	sort.Slice(results, func(i, j int) bool {
		return results[i].SchoolID < results[j].SchoolID
	})

	return results
}

func main() {
	// Test data from the problem.
	schools := []School{
		{SchoolID: 5, Capacity: 48},
		{SchoolID: 9, Capacity: 9},
		{SchoolID: 10, Capacity: 99},
		{SchoolID: 11, Capacity: 151},
	}

	exams := []Exam{
		{Score: 975, StudentCount: 10},
		{Score: 966, StudentCount: 60},
		{Score: 844, StudentCount: 76},
		{Score: 749, StudentCount: 76},
		{Score: 744, StudentCount: 100},
	}

	results := findCutoffScore(schools, exams)

	fmt.Println("Cutoff Score for Each School (school_id | score):")
	for _, r := range results {
		fmt.Printf("%d | %d\n", r.SchoolID, r.Score)
	}
	// Expected output:
	// 5 | 975  (capacity 48, only 975(10) qualifies; 966 has 60 > 48)
	// 9 | -1   (capacity 9, minimum student_count is 10 > 9)
	// 10 | 749 (capacity 99, 749(76) qualifies, 744(100) > 99)
	// 11 | 744 (capacity 151, all scores qualify; minimum is 744)
}
```

## 1989 — Maximum Number Of People That Can Be Caught In Tag

```go
package main

// LeetCode #1989: Maximum Number of People That Can Be Caught in Tag
// https://leetcode.com/problems/maximum-number-of-people-that-can-be-caught-in-tag/
// Difficulty: Medium [Paid]

import "fmt"

func main() {
	fmt.Println(MaximumNumberOfPeopleThatCanBeCaughtInTag([]int{0, 1, 1, 0}, 2))
	fmt.Println(MaximumNumberOfPeopleThatCanBeCaughtInTag([]int{1, 1, 0, 0}, 1))
	fmt.Println(MaximumNumberOfPeopleThatCanBeCaughtInTag([]int{0, 0, 1, 0, 1, 0, 0, 1, 1, 0}, 5))
}

// Time: O(n), Space: O(1)
func MaximumNumberOfPeopleThatCanBeCaughtInTag(team []int, dist int) int {
	ans := 0
	n := len(team)
	j := 0
	for i := 0; i < n; i++ {
		if team[i] == 1 {
			for j < n && (team[j] == 1 || i-j > dist) {
				j++
			}
			if j < n && abs1989(i-j) <= dist {
				ans++
				j++
			}
		}
	}
	return ans
}

func abs1989(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

## 1990 — Count The Number Of Experiments

```go
package main

// LeetCode #1990: Count the Number of Experiments
// https://leetcode.com/problems/count-the-number-of-experiments/
// Difficulty: Medium (SQL problem — simulated in Go)
//
// Simulates: Count experiments for ALL 9 (platform, experiment_name)
// combinations (3 platforms x 3 names). Include zeros for missing
// combinations.

import (
	"fmt"
)

// Experiment represents the Experiments database table.
type Experiment struct {
	ExperimentID   int
	Platform       string // 'Android', 'IOS', 'Web'
	ExperimentName string // 'Reading', 'Sports', 'Programming'
}

// ExperimentResult holds the output.
type ExperimentResult struct {
	Platform       string
	ExperimentName string
	NumExperiments int
}

// countExperiments simulates the SQL query.
// Time: O(n + 9) = O(n) | Space: O(p * n) where p = platforms, n = names
// n = number of experiment rows.
func countExperiments(experiments []Experiment) []ExperimentResult {
	platforms := []string{"Android", "IOS", "Web"}
	names := []string{"Reading", "Sports", "Programming"}

	// Count existing experiments.
	counts := make(map[string]map[string]int)
	for _, e := range experiments {
		if counts[e.Platform] == nil {
			counts[e.Platform] = make(map[string]int)
		}
		counts[e.Platform][e.ExperimentName]++
	}

	// Build the 9-combination result.
	var results []ExperimentResult
	for _, p := range platforms {
		for _, n := range names {
			cnt := 0
			if counts[p] != nil {
				cnt = counts[p][n]
			}
			results = append(results, ExperimentResult{
				Platform:       p,
				ExperimentName: n,
				NumExperiments: cnt,
			})
		}
	}

	return results
}

func main() {
	// Test data from the problem.
	experiments := []Experiment{
		{ExperimentID: 1, Platform: "Android", ExperimentName: "Reading"},
		{ExperimentID: 2, Platform: "Android", ExperimentName: "Sports"},
		{ExperimentID: 3, Platform: "Android", ExperimentName: "Programming"},
		{ExperimentID: 4, Platform: "IOS", ExperimentName: "Reading"},
		{ExperimentID: 5, Platform: "IOS", ExperimentName: "Sports"},
		{ExperimentID: 6, Platform: "Web", ExperimentName: "Reading"},
	}

	results := countExperiments(experiments)

	fmt.Println("Experiment Counts (platform | experiment_name | num_experiments):")
	for _, r := range results {
		fmt.Printf("%s | %s | %d\n", r.Platform, r.ExperimentName, r.NumExperiments)
	}
	// Expected output (all 9 combinations):
	// Android | Reading | 1
	// Android | Sports | 1
	// Android | Programming | 1
	// IOS | Reading | 1
	// IOS | Sports | 1
	// IOS | Programming | 0
	// Web | Reading | 1
	// Web | Sports | 0
	// Web | Programming | 0
}
```

## 1992 — Find All Groups Of Farmland

```go
package main

// LeetCode #1992: Find All Groups of Farmland
// https://leetcode.com/problems/find-all-groups-of-farmland/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(FindAllGroupsOfFarmland([][]int{{1, 0, 0}, {0, 1, 1}, {0, 1, 1}}))
	fmt.Println(FindAllGroupsOfFarmland([][]int{{1, 1}, {1, 1}}))
	fmt.Println(FindAllGroupsOfFarmland([][]int{{0}}))
}

// Time: O(m*n), Space: O(1) (excluding output)
func FindAllGroupsOfFarmland(land [][]int) [][]int {
	m, n := len(land), len(land[0])
	result := make([][]int, 0)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if land[i][j] == 1 {
				r, c := i, j
				for r+1 < m && land[r+1][j] == 1 {
					r++
				}
				for c+1 < n && land[i][c+1] == 1 {
					c++
				}
				result = append(result, []int{i, j, r, c})
				for x := i; x <= r; x++ {
					for y := j; y <= c; y++ {
						land[x][y] = 0
					}
				}
			}
		}
	}

	return result
}
```

## 1993 — Operations On Tree

```go
package main

// LeetCode #1993: Operations on Tree
// https://leetcode.com/problems/operations-on-tree/
// Difficulty: Medium

import "fmt"

func main() {
	obj := Constructor1993([]int{-1, 0, 0, 1, 1, 2, 2})
	fmt.Println(obj.Lock(2, 2))
	fmt.Println(obj.Unlock(2, 3))
	fmt.Println(obj.Unlock(2, 2))
	fmt.Println(obj.Lock(4, 5))
	fmt.Println(obj.Upgrade(0, 1))
	fmt.Println(obj.Lock(0, 1))
}

// LockingTree struct represents the tree with lock states
type LockingTree struct {
	locked   []int
	parent   []int
	children [][]int
}

// Constructor1993 initializes the LockingTree
func Constructor1993(parent []int) LockingTree {
	n := len(parent)
	locked := make([]int, n)
	for i := range locked {
		locked[i] = -1
	}
	children := make([][]int, n)
	for i := 1; i < n; i++ {
		children[parent[i]] = append(children[parent[i]], i)
	}
	return LockingTree{locked, parent, children}
}

// Lock locks the node for the given user if it is unlocked
func (this *LockingTree) Lock(num int, user int) bool {
	if this.locked[num] == -1 {
		this.locked[num] = user
		return true
	}
	return false
}

// Unlock unlocks the node if it is locked by the given user
func (this *LockingTree) Unlock(num int, user int) bool {
	if this.locked[num] == user {
		this.locked[num] = -1
		return true
	}
	return false
}

// Upgrade locks the node for the user if:
// 1. node is unlocked
// 2. at least one locked descendant exists
// 3. no locked ancestors exist
func (this *LockingTree) Upgrade(num int, user int) bool {
	if this.locked[num] != -1 {
		return false
	}
	// Check ancestors
	x := num
	for ; x != -1; x = this.parent[x] {
		if this.locked[x] != -1 {
			return false
		}
	}
	// Find and unlock locked descendants
	find := false
	var dfs func(int)
	dfs = func(x int) {
		for _, y := range this.children[x] {
			if this.locked[y] != -1 {
				find = true
				this.locked[y] = -1
			}
			dfs(y)
		}
	}
	dfs(num)
	if !find {
		return false
	}
	this.locked[num] = user
	return true
}
```

## 1996 — The Number Of Weak Characters In The Game

```go
package main

// LeetCode #1996: The Number of Weak Characters in the Game
// https://leetcode.com/problems/the-number-of-weak-characters-in-the-game/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(TheNumberOfWeakCharactersInTheGame([][]int{{5, 5}, {6, 3}, {3, 6}}))
	fmt.Println(TheNumberOfWeakCharactersInTheGame([][]int{{2, 2}, {3, 3}}))
	fmt.Println(TheNumberOfWeakCharactersInTheGame([][]int{{1, 5}, {10, 4}, {4, 3}}))
}

// Time: O(n log n), Space: O(1) (ignoring sort space)
func TheNumberOfWeakCharactersInTheGame(properties [][]int) int {
	sort.Slice(properties, func(i, j int) bool {
		if properties[i][0] == properties[j][0] {
			return properties[i][1] < properties[j][1]
		}
		return properties[i][0] > properties[j][0]
	})

	ans := 0
	maxDef := 0
	for _, p := range properties {
		if p[1] < maxDef {
			ans++
		} else {
			maxDef = p[1]
		}
	}
	return ans
}
```

## 1997 — First Day Where You Have Been In All The Rooms

```go
package main

// LeetCode #1997: First Day Where You Have Been in All the Rooms
// https://leetcode.com/problems/first-day-where-you-have-been-in-all-the-rooms/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(FirstDayWhereYouHaveBeenInAllTheRooms([]int{0, 0}))
	fmt.Println(FirstDayWhereYouHaveBeenInAllTheRooms([]int{0, 1, 2, 0}))
	fmt.Println(FirstDayWhereYouHaveBeenInAllTheRooms([]int{0, 0, 2}))
}

// Time: O(n), Space: O(n)
func FirstDayWhereYouHaveBeenInAllTheRooms(nextVisit []int) int {
	const mod = 1_000_000_007
	n := len(nextVisit)
	s := make([]int, n)

	for i := 0; i < n-1; i++ {
		j := nextVisit[i]
		s[i+1] = (s[i]*2 - s[j] + 2) % mod
		if s[i+1] < 0 {
			s[i+1] += mod
		}
	}

	return s[n-1]
}
```

## 1999 — Smallest Greater Multiple Made Of Two Digits

```go
package main

// LeetCode #1999: Smallest Greater Multiple Made of Two Digits
// https://leetcode.com/problems/smallest-greater-multiple-made-of-two-digits/
// Difficulty: Medium [Paid]

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(SmallestGreaterMultipleMadeOfTwoDigits(2, 0, 2))
	fmt.Println(SmallestGreaterMultipleMadeOfTwoDigits(8, 2, 4))
	fmt.Println(SmallestGreaterMultipleMadeOfTwoDigits(2, 3, 9))
}

// Time: O(2^k * log n) where k is number of digits, Space: O(2^k)
func SmallestGreaterMultipleMadeOfTwoDigits(k int, digit1 int, digit2 int) int {
	if digit1 > digit2 {
		digit1, digit2 = digit2, digit1
	}

	nums := make([]int, 0)

	// BFS to generate all numbers using only digit1 and digit2
	if digit1 != 0 {
		nums = append(nums, digit1)
	}
	if digit2 != 0 && digit2 != digit1 {
		nums = append(nums, digit2)
	}

	queue := nums[:]
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr > math.MaxInt32/10 {
			continue
		}
		n1 := curr*10 + digit1
		if n1 <= math.MaxInt32 {
			nums = append(nums, n1)
			queue = append(queue, n1)
		}
		if digit2 != digit1 {
			n2 := curr*10 + digit2
			if n2 <= math.MaxInt32 {
				nums = append(nums, n2)
				queue = append(queue, n2)
			}
		}
	}

	sort.Ints(nums)

	for _, num := range nums {
		if num >= k && num%k == 0 {
			return num
		}
	}
	return -1
}
```

## 2001 — Number Of Pairs Of Interchangeable Rectangles

```go
package main

// LeetCode #2001: Number of Pairs of Interchangeable Rectangles
// https://leetcode.com/problems/number-of-pairs-of-interchangeable-rectangles/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(NumberOfPairsOfInterchangeableRectangles([][]int{{4, 8}, {3, 6}, {10, 20}, {15, 30}}))
	fmt.Println(NumberOfPairsOfInterchangeableRectangles([][]int{{4, 5}, {7, 8}}))
}

// Time: O(n log max(w,h)), Space: O(n)
func NumberOfPairsOfInterchangeableRectangles(rectangles [][]int) int64 {
	cnt := make(map[[2]int]int64)
	for _, r := range rectangles {
		w, h := r[0], r[1]
		g := gcd2001(w, h)
		key := [2]int{w / g, h / g}
		cnt[key]++
	}

	var ans int64
	for _, m := range cnt {
		ans += m * (m - 1) / 2
	}
	return ans
}

func gcd2001(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
```

## 2002 — Maximum Product Of The Length Of Two Palindromic Subsequences

```go
package main

// LeetCode #2002: Maximum Product of the Length of Two Palindromic Subsequences
// https://leetcode.com/problems/maximum-product-of-the-length-of-two-palindromic-subsequences/
// Difficulty: Medium

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(MaximumProductOfTheLengthOfTwoPalindromicSubsequences("leetcodecom"))
	fmt.Println(MaximumProductOfTheLengthOfTwoPalindromicSubsequences("bb"))
	fmt.Println(MaximumProductOfTheLengthOfTwoPalindromicSubsequences("accbcaxxcxx"))
}

// Time: O(3^n), Space: O(2^n)
func MaximumProductOfTheLengthOfTwoPalindromicSubsequences(s string) int {
	n := len(s)
	p := make([]bool, 1<<n)

	for mask := 1; mask < 1<<n; mask++ {
		p[mask] = true
		i, j := 0, n-1
		for i < j {
			for i < j && (mask>>i&1) == 0 {
				i++
			}
			for i < j && (mask>>j&1) == 0 {
				j--
			}
			if i < j && s[i] != s[j] {
				p[mask] = false
				break
			}
			i++
			j--
		}
	}

	ans := 0
	all := 1<<n - 1
	for mask := 1; mask < 1<<n; mask++ {
		if !p[mask] {
			continue
		}
		a := bits.OnesCount(uint(mask))
		rest := all ^ mask
		for sub := rest; sub > 0; sub = (sub - 1) & rest {
			if p[sub] {
				b := bits.OnesCount(uint(sub))
				if a*b > ans {
					ans = a * b
				}
			}
		}
	}

	return ans
}
```

## 2007 — Find Original Array From Doubled Array

```go
package main

// LeetCode #2007: Find Original Array From Doubled Array
// https://leetcode.com/problems/find-original-array-from-doubled-array/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(FindOriginalArrayFromDoubledArray([]int{1, 3, 4, 2, 6, 8}))
	fmt.Println(FindOriginalArrayFromDoubledArray([]int{6, 3, 0, 1}))
	fmt.Println(FindOriginalArrayFromDoubledArray([]int{1}))
}

// Time: O(n log n), Space: O(n)
func FindOriginalArrayFromDoubledArray(changed []int) []int {
	n := len(changed)
	if n%2 == 1 {
		return []int{}
	}

	sort.Ints(changed)
	maxVal := changed[n-1]
	cnt := make([]int, maxVal+1)
	for _, x := range changed {
		cnt[x]++
	}

	ans := make([]int, 0, n/2)
	for _, x := range changed {
		if cnt[x] == 0 {
			continue
		}
		if x*2 > maxVal || cnt[x*2] == 0 {
			return []int{}
		}
		ans = append(ans, x)
		cnt[x]--
		cnt[x*2]--
	}

	if len(ans) != n/2 {
		return []int{}
	}
	return ans
}
```

## 2008 — Maximum Earnings From Taxi

```go
package main

// LeetCode #2008: Maximum Earnings From Taxi
// https://leetcode.com/problems/maximum-earnings-from-taxi/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MaximumEarningsFromTaxi(5, [][]int{{2, 5, 4}, {1, 5, 1}}))
	fmt.Println(MaximumEarningsFromTaxi(20, [][]int{{1, 6, 1}, {3, 10, 2}, {10, 12, 3}, {11, 12, 2}, {12, 15, 2}, {13, 18, 1}}))
}

// Time: O(m log m), Space: O(m)
func MaximumEarningsFromTaxi(n int, rides [][]int) int64 {
	m := len(rides)
	sort.Slice(rides, func(i, j int) bool {
		return rides[i][1] < rides[j][1]
	})

	dp := make([]int64, m+1)
	endTimes := make([]int, m)
	for i := 0; i < m; i++ {
		endTimes[i] = rides[i][1]
	}

	for i := 1; i <= m; i++ {
		start, end, tip := rides[i-1][0], rides[i-1][1], rides[i-1][2]
		earn := int64(end - start + tip)
		dp[i] = dp[i-1]

		j := sort.Search(m, func(k int) bool {
			return rides[k][1] > start
		})
		if dp[j]+earn > dp[i] {
			dp[i] = dp[j] + earn
		}
	}

	return dp[m]
}
```

## 2012 — Sum Of Beauty In The Array

```go
package main

// LeetCode #2012: Sum of Beauty in the Array
// https://leetcode.com/problems/sum-of-beauty-in-the-array/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(SumOfBeautyInTheArray([]int{1, 2, 3}))
	fmt.Println(SumOfBeautyInTheArray([]int{2, 4, 6, 4}))
	fmt.Println(SumOfBeautyInTheArray([]int{3, 2, 1}))
}

// Time: O(n), Space: O(n)
func SumOfBeautyInTheArray(nums []int) int {
	n := len(nums)
	prefixMax := make([]int, n)
	suffixMin := make([]int, n)

	prefixMax[0] = nums[0]
	for i := 1; i < n; i++ {
		if nums[i] > prefixMax[i-1] {
			prefixMax[i] = nums[i]
		} else {
			prefixMax[i] = prefixMax[i-1]
		}
	}

	suffixMin[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		if nums[i] < suffixMin[i+1] {
			suffixMin[i] = nums[i]
		} else {
			suffixMin[i] = suffixMin[i+1]
		}
	}

	ans := 0
	for i := 1; i < n-1; i++ {
		if nums[i] > prefixMax[i-1] && nums[i] < suffixMin[i+1] {
			ans += 2
		} else if nums[i] > nums[i-1] && nums[i] < nums[i+1] {
			ans++
		}
	}

	return ans
}
```

## 2013 — Detect Squares

```go
package main

// LeetCode #2013: Detect Squares
// https://leetcode.com/problems/detect-squares/
// Difficulty: Medium

import "fmt"

func main() {
	obj := Constructor2013()
	obj.Add([]int{3, 10})
	obj.Add([]int{11, 2})
	obj.Add([]int{3, 2})
	fmt.Println(obj.Count([]int{11, 10}))
	fmt.Println(obj.Count([]int{14, 8}))
	obj.Add([]int{11, 2})
	fmt.Println(obj.Count([]int{11, 10}))
}

// DetectSquares struct
type DetectSquares struct {
	points map[[2]int]int
	xMap   map[int]map[int]int
}

// Constructor2013 initializes DetectSquares
func Constructor2013() DetectSquares {
	return DetectSquares{
		points: make(map[[2]int]int),
		xMap:   make(map[int]map[int]int),
	}
}

// Add adds a point
func (this *DetectSquares) Add(point []int) {
	key := [2]int{point[0], point[1]}
	this.points[key]++
	if this.xMap[point[0]] == nil {
		this.xMap[point[0]] = make(map[int]int)
	}
	this.xMap[point[0]][point[1]]++
}

// Count counts number of ways to form a square
func (this *DetectSquares) Count(point []int) int {
	x1, y1 := point[0], point[1]
	ans := 0

	for y2 := range this.xMap[x1] {
		if y2 == y1 {
			continue
		}
		side := y2 - y1
		if side < 0 {
			side = -side
		}

		// Two possible squares: x1 +/- side
		for _, x2 := range []int{x1 - side, x1 + side} {
			if this.xMap[x2] != nil {
				ans += this.xMap[x1][y2] * this.xMap[x2][y1] * this.xMap[x2][y2]
			}
		}
	}

	return ans
}
```

## 2015 — Average Height Of Buildings In Each Segment

```go
package main

// LeetCode #2015: Average Height of Buildings in Each Segment
// https://leetcode.com/problems/average-height-of-buildings-in-each-segment/
// Difficulty: Medium [Paid]

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(AverageHeightOfBuildingsInEachSegment([][]int{{1, 4, 2}, {3, 9, 4}}))
	fmt.Println(AverageHeightOfBuildingsInEachSegment([][]int{{1, 3, 2}, {2, 5, 3}, {2, 8, 3}}))
	fmt.Println(AverageHeightOfBuildingsInEachSegment([][]int{{1, 2, 1}, {5, 6, 1}}))
}

// Time: O(n log n), Space: O(n)
func AverageHeightOfBuildingsInEachSegment(buildings [][]int) [][]int {
	type event struct {
		pos    int
		height int
		change int
	}

	events := make([]event, 0)
	for _, b := range buildings {
		start, end, height := b[0], b[1], b[2]
		events = append(events, event{start, height, 1})
		events = append(events, event{end, height, -1})
	}

	sort.Slice(events, func(i, j int) bool {
		if events[i].pos != events[j].pos {
			return events[i].pos < events[j].pos
		}
		return events[i].height < events[j].height
	})

	result := make([][]int, 0)
	totalHeight := 0
	buildingCount := 0
	prevPos := -1

	for _, e := range events {
		if prevPos != -1 && prevPos < e.pos && buildingCount > 0 {
			avg := totalHeight / buildingCount
			n := len(result)
			if n > 0 && result[n-1][2] == avg && result[n-1][1] == prevPos {
				result[n-1][1] = e.pos
			} else {
				result = append(result, []int{prevPos, e.pos, avg})
			}
		}

		totalHeight += e.height * e.change
		buildingCount += e.change
		prevPos = e.pos
	}

	return result
}
```

## 2017 — Grid Game

```go
package main

// LeetCode #2017: Grid Game
// https://leetcode.com/problems/grid-game/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func gridGame(grid [][]int) int64 {
	n := len(grid[0])

	topPrefix := make([]int64, n+1)
	bottomPrefix := make([]int64, n+1)
	for i := 0; i < n; i++ {
		topPrefix[i+1] = topPrefix[i] + int64(grid[0][i])
		bottomPrefix[i+1] = bottomPrefix[i] + int64(grid[1][i])
	}

	result := int64(1<<63 - 1)

	for k := 0; k < n; k++ {
		// Robot 1 goes down at column k
		topRemaining := topPrefix[n] - topPrefix[k+1]
		bottomRemaining := bottomPrefix[k]
		score := topRemaining
		if bottomRemaining > score {
			score = bottomRemaining
		}
		if score < result {
			result = score
		}
	}

	return result
}

func main() {
	// Test case 1
	grid1 := [][]int{{2, 5, 4}, {1, 5, 1}}
	fmt.Println("Test 1:", gridGame(grid1))
	// Expected: 4

	// Test case 2
	grid2 := [][]int{{3, 3, 1}, {8, 5, 2}}
	fmt.Println("Test 2:", gridGame(grid2))
	// Expected: 4

	// Test case 3
	grid3 := [][]int{{1, 3, 1, 15}, {1, 3, 3, 1}}
	fmt.Println("Test 3:", gridGame(grid3))
	// Expected: 7
}
```

## 2018 — Check If Word Can Be Placed In Crossword

```go
package main

// LeetCode #2018: Check if Word Can Be Placed In Crossword
// https://leetcode.com/problems/check-if-word-can-be-placed-in-crossword/
// Difficulty: Medium
// Time: O(m*n*len(word)) | Space: O(1)

import "fmt"

func placeWordInCrossword(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	w := len(word)
	dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == '#' {
				continue
			}
			for _, d := range dirs {
				// Check if this is a valid starting position:
				// either at edge or previous cell is blocked
				pi, pj := i-d[0], j-d[1]
				if pi >= 0 && pi < m && pj >= 0 && pj < n && board[pi][pj] != '#' {
					continue
				}

				// Try to place the word
				ok := true
				for k := 0; k < w; k++ {
					ci, cj := i+k*d[0], j+k*d[1]
					if ci < 0 || ci >= m || cj < 0 || cj >= n {
						ok = false
						break
					}
					if board[ci][cj] == '#' || (board[ci][cj] != ' ' && board[ci][cj] != word[k]) {
						ok = false
						break
					}
				}
				if !ok {
					continue
				}

				// Check that cell after word is blocked or out of bounds
				ni, nj := i+w*d[0], j+w*d[1]
				if ni >= 0 && ni < m && nj >= 0 && nj < n && board[ni][nj] != '#' {
					continue
				}

				return true
			}
		}
	}
	return false
}

func main() {
	// Test case 1
	board1 := [][]byte{
		{'#', ' ', '#'},
		{' ', ' ', '#'},
		{'#', 'c', ' '},
	}
	fmt.Println("Test 1:", placeWordInCrossword(board1, "abc"))
	// Expected: true

	// Test case 2
	board2 := [][]byte{
		{' ', '#', 'a'},
		{' ', '#', 'c'},
		{' ', '#', 'a'},
	}
	fmt.Println("Test 2:", placeWordInCrossword(board2, "ac"))
	// Expected: false

	// Test case 3
	board3 := [][]byte{
		{' ', ' ', ' '},
		{' ', ' ', ' '},
		{' ', ' ', ' '},
	}
	fmt.Println("Test 3:", placeWordInCrossword(board3, "hello"))
	// Expected: false (board too small)
}
```

## 2020 — Number Of Accounts That Did Not Stream

```go
package main

// LeetCode #2020: Number of Accounts That Did Not Stream
// https://leetcode.com/problems/number-of-accounts-that-did-not-stream/
// Difficulty: Medium [Paid]
// Time: O(n + m) | Space: O(n)

import "fmt"

func numberOfAccountsThatDidNotStream(subscriptions [][]int, streams [][]int) int {
	if len(subscriptions) == 0 {
		return 0
	}

	// subscriptions[i] = [start_i, end_i] (account i's subscription period)
	// streams[j] = [date_j, account_j] (stream event)
	hasStreamed := make([]bool, len(subscriptions))

	for _, s := range streams {
		accID := s[1]
		date := s[0]
		if accID >= 0 && accID < len(subscriptions) {
			if date >= subscriptions[accID][0] && date <= subscriptions[accID][1] {
				hasStreamed[accID] = true
			}
		}
	}

	count := 0
	for _, v := range hasStreamed {
		if !v {
			count++
		}
	}
	return count
}

func main() {
	// Test case 1: Account 1 never streamed
	subs1 := [][]int{{0, 180}, {60, 365}, {180, 365}}
	streams1 := [][]int{{15, 0}, {150, 0}, {300, 2}}
	fmt.Println("Test 1:", numberOfAccountsThatDidNotStream(subs1, streams1))
	// Expected: 1

	// Test case 2: All streamed
	subs2 := [][]int{{0, 365}}
	streams2 := [][]int{{100, 0}}
	fmt.Println("Test 2:", numberOfAccountsThatDidNotStream(subs2, streams2))
	// Expected: 0

	// Test case 3: None streamed
	subs3 := [][]int{{0, 30}, {31, 60}}
	fmt.Println("Test 3:", numberOfAccountsThatDidNotStream(subs3, nil))
	// Expected: 2
}
```

## 2021 — Brightest Position On Street

```go
package main

// LeetCode #2021: Brightest Position on Street
// https://leetcode.com/problems/brightest-position-on-street/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func brightestPosition(lights [][]int) int {
	events := make([][2]int, 0, len(lights)*2)

	for _, l := range lights {
		pos, rng := l[0], l[1]
		events = append(events, [2]int{pos - rng, 1})
		events = append(events, [2]int{pos + rng + 1, -1})
	}

	sort.Slice(events, func(i, j int) bool {
		if events[i][0] != events[j][0] {
			return events[i][0] < events[j][0]
		}
		return events[i][1] < events[j][1]
	})

	maxBrightness := 0
	currBrightness := 0
	bestPos := events[0][0]

	for _, e := range events {
		currBrightness += e[1]
		if currBrightness > maxBrightness {
			maxBrightness = currBrightness
			bestPos = e[0]
		}
	}

	return bestPos
}

func main() {
	// Test case 1
	lights1 := [][]int{{-3, 2}, {1, 2}, {3, 3}}
	fmt.Println("Test 1:", brightestPosition(lights1))
	// Expected: -1

	// Test case 2
	lights2 := [][]int{{1, 0}, {0, 1}}
	fmt.Println("Test 2:", brightestPosition(lights2))
	// Expected: 1

	// Test case 3
	lights3 := [][]int{{1, 2}}
	fmt.Println("Test 3:", brightestPosition(lights3))
	// Expected: -1
}
```

## 2023 — Number Of Pairs Of Strings With Concatenation Equal To Target

```go
package main

// LeetCode #2023: Number of Pairs of Strings With Concatenation Equal to Target
// https://leetcode.com/problems/number-of-pairs-of-strings-with-concatenation-equal-to-target/
// Difficulty: Medium
// Time: O(n * L) | Space: O(n)

import "fmt"

func numOfPairs(nums []string, target string) int {
	freq := make(map[string]int)
	count := 0

	for _, num := range nums {
		// Check if any prefix/suffix of target matches
		for i := 1; i < len(target); i++ {
			prefix := target[:i]
			suffix := target[i:]
			if num == prefix {
				count += freq[suffix]
			}
			if num == suffix {
				count += freq[prefix]
			}
		}
		freq[num]++
	}

	return count
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", numOfPairs([]string{"777", "7", "77", "77"}, "777"))
	// Expected: 4

	// Test case 2
	fmt.Println("Test 2:", numOfPairs([]string{"123", "4", "12", "34"}, "1234"))
	// Expected: 2

	// Test case 3
	fmt.Println("Test 3:", numOfPairs([]string{"1", "1", "1"}, "11"))
	// Expected: 6
}
```

## 2024 — Maximize The Confusion Of An Exam

```go
package main

// LeetCode #2024: Maximize the Confusion of an Exam
// https://leetcode.com/problems/maximize-the-confusion-of-an-exam/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func maxConsecutiveAnswers(answerKey string, k int) int {
	return max(
		longestChar(answerKey, k, 'T'),
		longestChar(answerKey, k, 'F'),
	)
}

func longestChar(s string, k int, target byte) int {
	left := 0
	flips := 0
	maxLen := 0

	for right := 0; right < len(s); right++ {
		if s[right] != target {
			flips++
		}
		for flips > k {
			if s[left] != target {
				flips--
			}
			left++
		}
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", maxConsecutiveAnswers("TTFF", 2))
	// Expected: 4

	// Test case 2
	fmt.Println("Test 2:", maxConsecutiveAnswers("TFFT", 1))
	// Expected: 3

	// Test case 3
	fmt.Println("Test 3:", maxConsecutiveAnswers("TTFTTFTT", 1))
	// Expected: 5
}
```

## 2028 — Find Missing Observations

```go
package main

// LeetCode #2028: Find Missing Observations
// https://leetcode.com/problems/find-missing-observations/
// Difficulty: Medium
// Time: O(n + m) | Space: O(m)

import "fmt"

func missingRolls(rolls []int, mean int, n int) []int {
	m := len(rolls)
	totalSum := mean * (m + n)
	knownSum := 0
	for _, v := range rolls {
		knownSum += v
	}
	missingSum := totalSum - knownSum

	if missingSum < n || missingSum > 6*n {
		return []int{}
	}

	result := make([]int, n)
	for i := range result {
		val := missingSum / (n - i)
		if val > 6 {
			val = 6
		}
		result[i] = val
		missingSum -= val
	}

	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", missingRolls([]int{3, 2, 4, 3}, 4, 2))
	// Expected: [6,6]

	// Test case 2
	fmt.Println("Test 2:", missingRolls([]int{1, 5, 6}, 3, 4))
	// Expected: [2,3,2,2]

	// Test case 3
	fmt.Println("Test 3:", missingRolls([]int{1, 2, 3, 4}, 6, 4))
	// Expected: [] (impossible)
}
```

## 2029 — Stone Game Ix

```go
package main

// LeetCode #2029: Stone Game IX
// https://leetcode.com/problems/stone-game-ix/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func stoneGameIX(stones []int) bool {
	count := [3]int{}
	for _, v := range stones {
		count[v%3]++
	}

	// If no stone with value % 3 == 0 and the other two counts differ by at most 1
	if count[0]%2 == 0 {
		return count[1] > 0 && count[2] > 0
	}
	// If count[0] is odd
	return abs(count[1]-count[2]) > 2
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", stoneGameIX([]int{2, 1}))
	// Expected: true

	// Test case 2
	fmt.Println("Test 2:", stoneGameIX([]int{2}))
	// Expected: false

	// Test case 3
	fmt.Println("Test 3:", stoneGameIX([]int{5, 1, 2, 4, 3}))
	// Expected: false
}
```

## 2031 — Count Subarrays With More Ones Than Zeros

```go
package main

// LeetCode #2031: Count Subarrays With More Ones Than Zeros
// https://leetcode.com/problems/count-subarrays-with-more-ones-than-zeros/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func subarraysWithMoreOnesThanZeros(nums []int) int {
	// Treat 0 as -1, 1 as +1
	// prefix[j] - prefix[i] > 0 means subarray (i,j] has more ones
	// prefix[j] > prefix[i] -> count how many previous prefixes are smaller
	prefix := 0
	count := 0
	// Fenwick tree or segment tree for range counts, but simpler:
	// offset by n since prefix ranges from -n to n
	n := len(nums)
	tree := make([]int, 2*n+2)
	mod := int(1e9 + 7)

	add := func(idx int) {
		idx += n + 1
		for idx < len(tree) {
			tree[idx]++
			idx += idx & -idx
		}
	}

	sum := func(idx int) int {
		idx += n + 1
		res := 0
		for idx > 0 {
			res += tree[idx]
			idx -= idx & -idx
		}
		return res
	}

	add(0) // empty prefix
	for _, v := range nums {
		if v == 1 {
			prefix++
		} else {
			prefix--
		}
		// Count how many previous prefixes are less than current prefix
		// sum(prefix-1) = count of prefixes <= prefix-1
		count = (count + sum(prefix-1)) % mod
		add(prefix)
	}

	return count
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", subarraysWithMoreOnesThanZeros([]int{0, 1, 1, 0, 1}))
	// Expected: 9

	// Test case 2
	fmt.Println("Test 2:", subarraysWithMoreOnesThanZeros([]int{1, 0, 1}))
	// Expected: 3

	// Test case 3
	fmt.Println("Test 3:", subarraysWithMoreOnesThanZeros([]int{1, 1, 1}))
	// Expected: 6
}
```

## 2033 — Minimum Operations To Make A Uni Value Grid

```go
package main

// LeetCode #2033: Minimum Operations to Make a Uni-Value Grid
// https://leetcode.com/problems/minimum-operations-to-make-a-uni-value-grid/
// Difficulty: Medium
// Time: O(m*n log(m*n)) | Space: O(m*n)

import (
	"fmt"
	"sort"
)

func minOperations(grid [][]int, x int) int {
	m, n := len(grid), len(grid[0])
	vals := make([]int, 0, m*n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			vals = append(vals, grid[i][j])
		}
	}

	// Check all values have same remainder mod x
	rem := vals[0] % x
	for _, v := range vals {
		if v%x != rem {
			return -1
		}
	}

	sort.Ints(vals)
	median := vals[len(vals)/2]

	ops := 0
	for _, v := range vals {
		diff := v - median
		if diff < 0 {
			diff = -diff
		}
		ops += diff / x
	}
	return ops
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minOperations([][]int{{2, 4}, {6, 8}}, 2))
	// Expected: 4

	// Test case 2
	fmt.Println("Test 2:", minOperations([][]int{{1, 5}, {2, 3}}, 1))
	// Expected: 5

	// Test case 3
	fmt.Println("Test 3:", minOperations([][]int{{1, 2}, {3, 4}}, 2))
	// Expected: -1
}
```

## 2034 — Stock Price Fluctuation

```go
package main

// LeetCode #2034: Stock Price Fluctuation
// https://leetcode.com/problems/stock-price-fluctuation/
// Difficulty: Medium
// Time: O(log n) per operation | Space: O(n)

import (
	"container/heap"
	"fmt"
)

type StockPrice struct {
	prices   map[int]int
	latestTs int
	latestP  int
	minHeap  *MinHeap
	maxHeap  *MaxHeap
}

type PriceEntry struct {
	price int
	ts    int
}

type MinHeap []PriceEntry
func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].price < h[j].price }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(PriceEntry)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type MaxHeap []PriceEntry
func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i].price > h[j].price }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(PriceEntry)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func Constructor() StockPrice {
	return StockPrice{
		prices:  make(map[int]int),
		minHeap: &MinHeap{},
		maxHeap: &MaxHeap{},
	}
}

func (sp *StockPrice) Update(timestamp int, price int) {
	sp.prices[timestamp] = price
	if timestamp >= sp.latestTs {
		sp.latestTs = timestamp
		sp.latestP = price
	}
	heap.Push(sp.minHeap, PriceEntry{price, timestamp})
	heap.Push(sp.maxHeap, PriceEntry{price, timestamp})
}

func (sp *StockPrice) Current() int {
	return sp.latestP
}

func (sp *StockPrice) Maximum() int {
	for sp.maxHeap.Len() > 0 {
		top := (*sp.maxHeap)[0]
		if sp.prices[top.ts] == top.price {
			return top.price
		}
		heap.Pop(sp.maxHeap)
	}
	return 0
}

func (sp *StockPrice) Minimum() int {
	for sp.minHeap.Len() > 0 {
		top := (*sp.minHeap)[0]
		if sp.prices[top.ts] == top.price {
			return top.price
		}
		heap.Pop(sp.minHeap)
	}
	return 0
}

func main() {
	sp := Constructor()
	sp.Update(1, 10)
	sp.Update(2, 5)
	fmt.Println("Test 1 Current:", sp.Current())  // 5
	fmt.Println("Test 1 Maximum:", sp.Maximum())  // 10
	fmt.Println("Test 1 Minimum:", sp.Minimum())  // 5

	sp2 := Constructor()
	sp2.Update(1, 1)
	sp2.Update(2, 2)
	sp2.Update(3, 3)
	sp2.Update(1, 4)
	fmt.Println("Test 2 Maximum:", sp2.Maximum()) // 4
	fmt.Println("Test 2 Minimum:", sp2.Minimum()) // 2
}
```

## 2036 — Maximum Alternating Subarray Sum

```go
package main

// LeetCode #2036: Maximum Alternating Subarray Sum
// https://leetcode.com/problems/maximum-alternating-subarray-sum/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func maximumAlternatingSubarraySum(nums []int) int64 {
	// dp0: max sum ending at i with the current element as positive (even index in subarray)
	// dp1: max sum ending at i with the current element as negative (odd index in subarray)
	dp0 := int64(nums[0])
	dp1 := int64(-1 << 62) // very negative
	result := dp0

	for i := 1; i < len(nums); i++ {
		// nums[i] as positive: either start new, or continue from dp1
		newDp0 := max64(int64(nums[i]), dp1+int64(nums[i]))
		// nums[i] as negative: must continue from dp0
		newDp1 := dp0 - int64(nums[i])

		dp0, dp1 = newDp0, newDp1
		result = max64(result, dp0)
	}

	return result
}

func max64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", maximumAlternatingSubarraySum([]int{4, 2, 5, 3}))
	// Expected: 7

	// Test case 2
	fmt.Println("Test 2:", maximumAlternatingSubarraySum([]int{5, 6, 7, 8}))
	// Expected: 8

	// Test case 3
	fmt.Println("Test 3:", maximumAlternatingSubarraySum([]int{-1, -2, -3}))
	// Expected: -1
}
```

## 2038 — Remove Colored Pieces If Both Neighbors Are The Same Color

```go
package main

// LeetCode #2038: Remove Colored Pieces if Both Neighbors are the Same Color
// https://leetcode.com/problems/remove-colored-pieces-if-both-neighbors-are-the-same-color/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func winnerOfGame(colors string) bool {
	aMoves := 0
	bMoves := 0
	count := 1

	for i := 1; i < len(colors); i++ {
		if colors[i] == colors[i-1] {
			count++
		} else {
			if colors[i-1] == 'A' && count >= 3 {
				aMoves += count - 2
			} else if colors[i-1] == 'B' && count >= 3 {
				bMoves += count - 2
			}
			count = 1
		}
	}
	if colors[len(colors)-1] == 'A' && count >= 3 {
		aMoves += count - 2
	} else if colors[len(colors)-1] == 'B' && count >= 3 {
		bMoves += count - 2
	}

	return aMoves > bMoves
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", winnerOfGame("AAABABB"))
	// Expected: true

	// Test case 2
	fmt.Println("Test 2:", winnerOfGame("AA"))
	// Expected: false

	// Test case 3
	fmt.Println("Test 3:", winnerOfGame("ABBBBBBBAAA"))
	// Expected: false
}
```

## 2039 — The Time When The Network Becomes Idle

```go
package main

// LeetCode #2039: The Time When the Network Becomes Idle
// https://leetcode.com/problems/the-time-when-the-network-becomes-idle/
// Difficulty: Medium
// Time: O(n + m) | Space: O(n + m)

import "fmt"

func networkBecomesIdle(edges [][]int, patience []int) int {
	n := len(patience)
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// BFS to find shortest distance from node 0
	dist := make([]int, n)
	for i := range dist {
		dist[i] = -1
	}
	dist[0] = 0
	queue := []int{0}
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		for _, v := range adj[u] {
			if dist[v] == -1 {
				dist[v] = dist[u] + 1
				queue = append(queue, v)
			}
		}
	}

	maxTime := 0
	for i := 1; i < n; i++ {
		// Distance to master server and back
		roundTrip := dist[i] * 2
		p := patience[i]

		// Last message sent at: floor((roundTrip-1)/p) * p
		// Message arrives at: last_send + roundTrip
		lastSend := ((roundTrip - 1) / p) * p
		lastArrival := lastSend + roundTrip

		if lastArrival > maxTime {
			maxTime = lastArrival
		}
	}

	// Network becomes idle 1 ms after last message
	return maxTime + 1
}

func main() {
	// Test case 1
	edges1 := [][]int{{0, 1}, {1, 2}}
	patience1 := []int{0, 2, 1}
	fmt.Println("Test 1:", networkBecomesIdle(edges1, patience1))
	// Expected: 8

	// Test case 2
	edges2 := [][]int{{0, 1}, {0, 2}, {1, 2}}
	patience2 := []int{0, 10, 10}
	fmt.Println("Test 2:", networkBecomesIdle(edges2, patience2))
	// Expected: 3
}
```

## 2041 — Accepted Candidates From The Interviews

```go
package main

// LeetCode #2041: Accepted Candidates From the Interviews
// https://leetcode.com/problems/accepted-candidates-from-the-interviews/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

type Candidate struct {
	ID       int
	Scores   []int
	Accepted bool
}

func acceptedCandidates(candidates []Candidate, minScore int) []int {
	result := []int{}
	for _, c := range candidates {
		count := 0
		for _, s := range c.Scores {
			if s >= minScore {
				count++
			}
		}
		if count >= 2 {
			// Calculate total
			total := 0
			sorted := make([]int, len(c.Scores))
			copy(sorted, c.Scores)
			sort.Ints(sorted)
			for _, s := range sorted[1:] {
				total += s
			}
			result = append(result, c.ID)
			_ = total
		}
	}
	return result
}

func main() {
	// Test case 1
	candidates1 := []Candidate{
		{1, []int{5, 6, 7}, false},
		{2, []int{8, 9, 10}, false},
	}
	fmt.Println("Test 1:", acceptedCandidates(candidates1, 5))
	// Expected: [1, 2]

	// Test case 2
	candidates2 := []Candidate{
		{1, []int{4, 5, 6}, false},
		{2, []int{7, 3, 8}, false},
		{3, []int{2, 3, 4}, false},
	}
	fmt.Println("Test 2:", acceptedCandidates(candidates2, 5))
	// Expected: [1, 2]

	// Test case 3
	fmt.Println("Test 3:", acceptedCandidates(nil, 5))
	// Expected: []
}
```

## 2043 — Simple Bank System

```go
package main

// LeetCode #2043: Simple Bank System
// https://leetcode.com/problems/simple-bank-system/
// Difficulty: Medium
// Time: O(1) per operation | Space: O(n)

import "fmt"

type Bank struct {
	balance []int64
}

func Constructor(balance []int64) Bank {
	return Bank{balance: balance}
}

func (b *Bank) Transfer(account1 int, account2 int, money int64) bool {
	if account1 < 1 || account1 > len(b.balance) || account2 < 1 || account2 > len(b.balance) {
		return false
	}
	if b.balance[account1-1] < money {
		return false
	}
	b.balance[account1-1] -= money
	b.balance[account2-1] += money
	return true
}

func (b *Bank) Deposit(account int, money int64) bool {
	if account < 1 || account > len(b.balance) {
		return false
	}
	b.balance[account-1] += money
	return true
}

func (b *Bank) Withdraw(account int, money int64) bool {
	if account < 1 || account > len(b.balance) {
		return false
	}
	if b.balance[account-1] < money {
		return false
	}
	b.balance[account-1] -= money
	return true
}

func main() {
	bank := Constructor([]int64{10, 100, 20, 50, 30})
	fmt.Println("Test 1 Deposit(3, 10):", bank.Deposit(3, 10))   // true
	fmt.Println("Test 2 Transfer(5, 1, 20):", bank.Transfer(5, 1, 20)) // true
	fmt.Println("Test 3 Withdraw(5, 20):", bank.Withdraw(5, 20)) // true
	fmt.Println("Test 4 Transfer(3, 4, 15):", bank.Transfer(3, 4, 15)) // false (insufficient)
	fmt.Println("Test 5 Deposit(9, 10):", bank.Deposit(9, 10))   // false
}
```

## 2044 — Count Number Of Maximum Bitwise Or Subsets

```go
package main

// LeetCode #2044: Count Number of Maximum Bitwise-OR Subsets
// https://leetcode.com/problems/count-number-of-maximum-bitwise-or-subsets/
// Difficulty: Medium
// Time: O(2^n) | Space: O(n)

import "fmt"

func countMaxOrSubsets(nums []int) int {
	maxOr := 0
	for _, v := range nums {
		maxOr |= v
	}

	count := 0
	var backtrack func(idx int, curOr int)
	backtrack = func(idx int, curOr int) {
		if idx == len(nums) {
			if curOr == maxOr {
				count++
			}
			return
		}
		// Skip current
		backtrack(idx+1, curOr)
		// Take current
		backtrack(idx+1, curOr|nums[idx])
	}

	backtrack(0, 0)
	return count
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", countMaxOrSubsets([]int{3, 1}))
	// Expected: 2

	// Test case 2
	fmt.Println("Test 2:", countMaxOrSubsets([]int{2, 2, 2}))
	// Expected: 7

	// Test case 3
	fmt.Println("Test 3:", countMaxOrSubsets([]int{3, 2, 1, 5}))
	// Expected: 6
}
```

## 2046 — Sort Linked List Already Sorted Using Absolute Values

```go
package main

// LeetCode #2046: Sort Linked List Already Sorted Using Absolute Values
// https://leetcode.com/problems/sort-linked-list-already-sorted-using-absolute-values/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortLinkedList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// Separate into negative and non-negative lists
	var negHead, negTail *ListNode
	var posHead, posTail *ListNode

	curr := head
	for curr != nil {
		if curr.Val < 0 {
			// For absolute-sorted list, negatives are in reverse order
			// Prepend to maintain reverse
			next := curr.Next
			if negHead == nil {
				negHead = curr
				negTail = curr
				curr.Next = nil
			} else {
				curr.Next = negHead
				negHead = curr
			}
			curr = next
		} else {
			if posHead == nil {
				posHead = curr
				posTail = curr
			} else {
				posTail.Next = curr
				posTail = curr
			}
			curr = curr.Next
		}
	}

	if posTail != nil {
		posTail.Next = nil
	}

	// Combine: negatives (reversed order) followed by positives
	if negHead != nil {
		if posHead != nil {
			negTail.Next = posHead
		}
		return negHead
	}
	return posHead
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}

func main() {
	// Test case 1
	head1 := &ListNode{1, &ListNode{-2, &ListNode{-3, &ListNode{4, &ListNode{-5, nil}}}}}
	fmt.Print("Test 1: ")
	printList(sortLinkedList(head1))
	// Expected: -5 -3 -2 1 4

	// Test case 2
	head2 := &ListNode{-1, &ListNode{-2, &ListNode{-3, nil}}}
	fmt.Print("Test 2: ")
	printList(sortLinkedList(head2))
	// Expected: -3 -2 -1

	// Test case 3
	head3 := &ListNode{0, &ListNode{1, &ListNode{2, nil}}}
	fmt.Print("Test 3: ")
	printList(sortLinkedList(head3))
	// Expected: 0 1 2
}
```

## 2048 — Next Greater Numerically Balanced Number

```go
package main

// LeetCode #2048: Next Greater Numerically Balanced Number
// https://leetcode.com/problems/next-greater-numerically-balanced-number/
// Difficulty: Medium
// Time: O(n) where n is the number until we find it | Space: O(1)

import "fmt"

func nextBeautifulNumber(n int) int {
	for i := n + 1; ; i++ {
		if isBalanced(i) {
			return i
		}
	}
}

func isBalanced(n int) bool {
	digits := make([]int, 10)
	for n > 0 {
		d := n % 10
		digits[d]++
		n /= 10
	}
	for d := 1; d <= 9; d++ {
		if digits[d] > 0 && digits[d] != d {
			return false
		}
	}
	return digits[0] == 0
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", nextBeautifulNumber(1))
	// Expected: 22

	// Test case 2
	fmt.Println("Test 2:", nextBeautifulNumber(1000))
	// Expected: 1333

	// Test case 3
	fmt.Println("Test 3:", nextBeautifulNumber(3000))
	// Expected: 3133
}
```

## 2049 — Count Nodes With The Highest Score

```go
package main

// LeetCode #2049: Count Nodes With the Highest Score
// https://leetcode.com/problems/count-nodes-with-the-highest-score/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func countHighestScoreNodes(parents []int) int {
	n := len(parents)
	children := make([][]int, n)
	for i := 1; i < n; i++ {
		p := parents[i]
		children[p] = append(children[p], i)
	}

	subtreeSize := make([]int, n)
	var dfs func(u int) int
	dfs = func(u int) int {
		size := 1
		for _, v := range children[u] {
			size += dfs(v)
		}
		subtreeSize[u] = size
		return size
	}
	dfs(0)

	maxScore := 0
	count := 0
	for i := 0; i < n; i++ {
		score := 1
		remaining := n
		for _, v := range children[i] {
			score *= subtreeSize[v]
			remaining -= subtreeSize[v]
		}
		if i != 0 {
			remaining = n - subtreeSize[i]
		} else {
			remaining = 0
		}
		if remaining > 0 {
			score *= remaining
		}

		if score > maxScore {
			maxScore = score
			count = 1
		} else if score == maxScore {
			count++
		}
	}
	return count
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", countHighestScoreNodes([]int{-1, 2, 0, 2, 0}))
	// Expected: 3

	// Test case 2
	fmt.Println("Test 2:", countHighestScoreNodes([]int{-1, 2, 0}))
	// Expected: 2

	// Test case 3
	fmt.Println("Test 3:", countHighestScoreNodes([]int{-1, 0, 1, 2}))
	// Expected: 1
}
```

