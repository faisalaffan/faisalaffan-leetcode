# Medium (Sedang) — Problem ��3043

## 2831 — Find The Longest Equal Subarray

```go
package main

// LeetCode #2831: Find the Longest Equal Subarray
// https://leetcode.com/problems/find-the-longest-equal-subarray/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func FindTheLongestEqualSubarray(nums []int, k int) int {
	pos := make(map[int][]int)
	for i, v := range nums {
		pos[v] = append(pos[v], i)
	}

	best := 0
	for _, indices := range pos {
		left := 0
		for right := 0; right < len(indices); right++ {
			// Elements between indices[left] and indices[right] that need to be removed
			for indices[right]-indices[left]-(right-left) > k {
				left++
			}
			if right-left+1 > best {
				best = right - left + 1
			}
		}
	}

	return best
}

func main() {
	fmt.Println(FindTheLongestEqualSubarray([]int{1, 3, 2, 3, 1, 3}, 3))
	fmt.Println(FindTheLongestEqualSubarray([]int{1, 1, 2, 2, 1, 1}, 2))
}
```

## 2832 — Maximal Range That Each Element Is Maximum In It

```go
package main

// LeetCode #2832: Maximal Range That Each Element Is Maximum in It
// https://leetcode.com/problems/maximal-range-that-each-element-is-maximum-in-it/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func MaximalRangeThatEachElementIsMaximumInIt(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	// Previous greater element
	prev := make([]int, n)
	stack := make([]int, 0)
	for i := 0; i < n; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			prev[i] = stack[len(stack)-1]
		} else {
			prev[i] = -1
		}
		stack = append(stack, i)
	}

	// Next greater element
	next := make([]int, n)
	stack = make([]int, 0)
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			next[i] = stack[len(stack)-1]
		} else {
			next[i] = n
		}
		stack = append(stack, i)
	}

	for i := 0; i < n; i++ {
		result[i] = next[i] - prev[i] - 1
	}

	return result
}

func main() {
	fmt.Println(MaximalRangeThatEachElementIsMaximumInIt([]int{1, 5, 4, 3, 6}))
	fmt.Println(MaximalRangeThatEachElementIsMaximumInIt([]int{1, 2, 1}))
}
```

## 2834 — Find The Minimum Possible Sum Of A Beautiful Array

```go
package main

// LeetCode #2834: Find the Minimum Possible Sum of a Beautiful Array
// https://leetcode.com/problems/find-the-minimum-possible-sum-of-a-beautiful-array/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func FindTheMinimumPossibleSumOfABeautifulArray(n int, target int) int {
	used := make(map[int]bool)
	sum := 0
	for i := 1; len(used) < n; i++ {
		if used[target-i] {
			continue
		}
		used[i] = true
		sum += i
	}
	return sum
}

func main() {
	fmt.Println(FindTheMinimumPossibleSumOfABeautifulArray(2, 3))
	fmt.Println(FindTheMinimumPossibleSumOfABeautifulArray(3, 3))
	fmt.Println(FindTheMinimumPossibleSumOfABeautifulArray(5, 5))
}
```

## 2838 — Maximum Coins Heroes Can Collect

```go
package main

// LeetCode #2838: Maximum Coins Heroes Can Collect
// https://leetcode.com/problems/maximum-coins-heroes-can-collect/
// Difficulty: Medium [Paid]
// Time: O(n log n + m log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func MaximumCoinsHeroesCanCollect(heroes []int, monsters []int, coins []int) []int64 {
	n := len(monsters)
	type monster struct {
		power int
		coin  int
	}
	monsterList := make([]monster, n)
	for i := 0; i < n; i++ {
		monsterList[i] = monster{monsters[i], coins[i]}
	}
	sort.Slice(monsterList, func(i, j int) bool {
		return monsterList[i].power < monsterList[j].power
	})

	prefix := make([]int64, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + int64(monsterList[i].coin)
	}

	heroSorted := make([]int, len(heroes))
	copy(heroSorted, heroes)
	sort.Ints(heroSorted)

	heroMap := make(map[int]int64)
	for _, h := range heroSorted {
		if _, ok := heroMap[h]; ok {
			continue
		}
		idx := sort.Search(n, func(i int) bool {
			return monsterList[i].power > h
		})
		heroMap[h] = prefix[idx]
	}

	result := make([]int64, len(heroes))
	for i, h := range heroes {
		result[i] = heroMap[h]
	}
	return result
}

func main() {
	fmt.Println(MaximumCoinsHeroesCanCollect([]int{1, 4, 2}, []int{1, 1, 3, 5}, []int{2, 3, 4, 5}))
	fmt.Println(MaximumCoinsHeroesCanCollect([]int{5}, []int{1, 2, 3}, []int{10, 20, 30}))
}
```

## 2840 — Check If Strings Can Be Made Equal With Operations Ii

```go
package main

// LeetCode #2840: Check if Strings Can be Made Equal With Operations II
// https://leetcode.com/problems/check-if-strings-can-be-made-equal-with-operations-ii/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func CheckIfStringsCanBeMadeEqualWithOperationsIi(s1 string, s2 string) bool {
	n := len(s1)
	// We can swap characters at positions with same parity
	// Count character frequencies at even and odd positions
	even := make([]int, 26)
	odd := make([]int, 26)

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			even[s1[i]-'a']++
			even[s2[i]-'a']--
		} else {
			odd[s1[i]-'a']++
			odd[s2[i]-'a']--
		}
	}

	for i := 0; i < 26; i++ {
		if even[i] != 0 || odd[i] != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(CheckIfStringsCanBeMadeEqualWithOperationsIi("abcd", "cdab"))
	fmt.Println(CheckIfStringsCanBeMadeEqualWithOperationsIi("abcd", "dcba"))
}
```

## 2841 — Maximum Sum Of Almost Unique Subarray

```go
package main

// LeetCode #2841: Maximum Sum of Almost Unique Subarray
// https://leetcode.com/problems/maximum-sum-of-almost-unique-subarray/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func MaximumSumOfAlmostUniqueSubarray(nums []int, m int, k int) int64 {
	n := len(nums)
	if n < k {
		return 0
	}

	freq := make(map[int]int)
	var sum int64
	var best int64

	for i := 0; i < k; i++ {
		freq[nums[i]]++
		sum += int64(nums[i])
	}

	if len(freq) >= m {
		best = sum
	}

	for i := k; i < n; i++ {
		// Remove leftmost
		left := nums[i-k]
		freq[left]--
		if freq[left] == 0 {
			delete(freq, left)
		}
		sum -= int64(left)

		// Add rightmost
		right := nums[i]
		freq[right]++
		sum += int64(right)

		if len(freq) >= m && sum > best {
			best = sum
		}
	}

	return best
}

func main() {
	fmt.Println(MaximumSumOfAlmostUniqueSubarray([]int{2, 6, 7, 3, 1, 7}, 3, 4))
	fmt.Println(MaximumSumOfAlmostUniqueSubarray([]int{1, 1, 1, 3}, 2, 2))
}
```

## 2844 — Minimum Operations To Make A Special Number

```go
package main

// LeetCode #2844: Minimum Operations to Make a Special Number
// https://leetcode.com/problems/minimum-operations-to-make-a-special-number/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import (
	"fmt"
	"math"
)

func MinimumOperationsToMakeASpecialNumber(num string) int {
	n := len(num)
	best := n // Remove all

	// Find "00"
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if num[i] == '0' && num[j] == '0' {
				ops := (n - 1 - j) + (j - i - 1)
				if ops < best {
					best = ops
				}
			}
			if num[i] == '2' && num[j] == '5' {
				ops := (n - 1 - j) + (j - i - 1)
				if ops < best {
					best = ops
				}
			}
			if num[i] == '5' && num[j] == '0' {
				ops := (n - 1 - j) + (j - i - 1)
				if ops < best {
					best = ops
				}
			}
			if num[i] == '7' && num[j] == '5' {
				ops := (n - 1 - j) + (j - i - 1)
				if ops < best {
					best = ops
				}
			}
		}
	}

	// Also check for single "0"
	for i := 0; i < n; i++ {
		if num[i] == '0' {
			ops := n - 1
			if ops < best {
				best = ops
			}
		}
	}

	if best == math.MaxInt32 {
		return n
	}
	return best
}

func main() {
	fmt.Println(MinimumOperationsToMakeASpecialNumber("2245047"))
	fmt.Println(MinimumOperationsToMakeASpecialNumber("2908305"))
	fmt.Println(MinimumOperationsToMakeASpecialNumber("10"))
}
```

## 2845 — Count Of Interesting Subarrays

```go
package main

// LeetCode #2845: Count of Interesting Subarrays
// https://leetcode.com/problems/count-of-interesting-subarrays/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func CountOfInterestingSubarrays(nums []int, modulo int, k int) int64 {
	prefix := make(map[int]int64)
	prefix[0] = 1
	var count int64
	var sum int

	for _, num := range nums {
		if num%modulo == k {
			sum++
		}
		need := (sum%modulo - k%modulo + modulo) % modulo
		count += prefix[need]
		prefix[sum%modulo]++
	}

	return count
}

func main() {
	fmt.Println(CountOfInterestingSubarrays([]int{3, 2, 4}, 2, 1))
	fmt.Println(CountOfInterestingSubarrays([]int{1, 2, 3, 4}, 3, 1))
}
```

## 2847 — Smallest Number With Given Digit Product

```go
package main

// LeetCode #2847: Smallest Number With Given Digit Product
// https://leetcode.com/problems/smallest-number-with-given-digit-product/
// Difficulty: Medium [Paid]
// Time: O(log n) | Space: O(log n)

import "fmt"

func SmallestNumberWithGivenDigitProduct(n int64) string {
	if n == 0 {
		return "0"
	}
	if n == 1 {
		return "1"
	}

	digits := make([]byte, 0)
	for i := 9; i >= 2; i-- {
		for n%int64(i) == 0 {
			digits = append([]byte{byte('0' + i)}, digits...)
			n /= int64(i)
		}
	}

	if n > 1 {
		return "-1"
	}

	return string(digits)
}

func main() {
	fmt.Println(SmallestNumberWithGivenDigitProduct(36))
	fmt.Println(SmallestNumberWithGivenDigitProduct(17))
	fmt.Println(SmallestNumberWithGivenDigitProduct(1))
}
```

## 2849 — Determine If A Cell Is Reachable At A Given Time

```go
package main

// LeetCode #2849: Determine if a Cell Is Reachable at a Given Time
// https://leetcode.com/problems/determine-if-a-cell-is-reachable-at-a-given-time/
// Difficulty: Medium
// Time: O(1) | Space: O(1)

import "fmt"

func DetermineIfACellIsReachableAtAGivenTime(sx int, sy int, fx int, fy int, t int) bool {
	dx := sx - fx
	if dx < 0 {
		dx = -dx
	}
	dy := sy - fy
	if dy < 0 {
		dy = -dy
	}
	minDist := dx
	if dy > minDist {
		minDist = dy
	}
	if minDist == 0 {
		return t != 1
	}
	return t >= minDist
}

func main() {
	fmt.Println(DetermineIfACellIsReachableAtAGivenTime(1, 2, 3, 4, 3))
	fmt.Println(DetermineIfACellIsReachableAtAGivenTime(1, 2, 1, 2, 1))
	fmt.Println(DetermineIfACellIsReachableAtAGivenTime(1, 1, 1, 1, 0))
}
```

## 2850 — Minimum Moves To Spread Stones Over Grid

```go
package main

// LeetCode #2850: Minimum Moves to Spread Stones Over Grid
// https://leetcode.com/problems/minimum-moves-to-spread-stones-over-grid/
// Difficulty: Medium
// Time: O(1) | Space: O(1)

import (
	"fmt"
	"math"
)

func MinimumMovesToSpreadStonesOverGrid(grid [][]int) int {
	// Grid is always 3x3
	// Collect positions of empty cells (0) and cells with extra stones (>1)
	zeros := make([][2]int, 0)
	extras := make([][2]int, 0)

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[i][j] == 0 {
				zeros = append(zeros, [2]int{i, j})
			}
			for k := 1; k < grid[i][j]; k++ {
				extras = append(extras, [2]int{i, j})
			}
		}
	}

	if len(zeros) == 0 {
		return 0
	}

	// Use permutation to find minimum total moves
	n := len(zeros)
	perm := make([]int, n)
	for i := range perm {
		perm[i] = i
	}

	best := math.MaxInt32
	for {
		total := 0
		for i, p := range perm {
			dx := extras[i][0] - zeros[p][0]
			if dx < 0 {
				dx = -dx
			}
			dy := extras[i][1] - zeros[p][1]
			if dy < 0 {
				dy = -dy
			}
			total += dx + dy
		}
		if total < best {
			best = total
		}

		// Next permutation
		k := -1
		for i := n - 2; i >= 0; i-- {
			if perm[i] < perm[i+1] {
				k = i
				break
			}
		}
		if k == -1 {
			break
		}
		l := n - 1
		for perm[l] <= perm[k] {
			l--
		}
		perm[k], perm[l] = perm[l], perm[k]
		for i, j := k+1, n-1; i < j; i, j = i+1, j-1 {
			perm[i], perm[j] = perm[j], perm[i]
		}
	}

	return best
}

func main() {
	fmt.Println(MinimumMovesToSpreadStonesOverGrid([][]int{{1, 1, 0}, {1, 1, 1}, {1, 2, 1}}))
	fmt.Println(MinimumMovesToSpreadStonesOverGrid([][]int{{1, 3, 0}, {1, 0, 0}, {1, 0, 3}}))
}
```

## 2852 — Sum Of Remoteness Of All Cells

```go
package main

// LeetCode #2852: Sum of Remoteness of All Cells
// https://leetcode.com/problems/sum-of-remoteness-of-all-cells/
// Difficulty: Medium [Paid]
// Time: O(n^2) | Space: O(n^2)

import "fmt"

func SumOfRemotenessOfAllCells(grid [][]int) []int64 {
	n := len(grid)
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	var totalSum int64
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] > 0 {
				totalSum += int64(grid[i][j])
			}
		}
	}

	result := make([]int64, n*n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] > 0 {
				result[i*n+j] = totalSum - int64(grid[i][j])
			}
		}
	}

	return result
}

func main() {
	fmt.Println(SumOfRemotenessOfAllCells([][]int{{1, 2}, {3, 4}}))
	fmt.Println(SumOfRemotenessOfAllCells([][]int{{5, 0}, {0, 5}}))
}
```

## 2854 — Rolling Average Steps

```go
package main

// LeetCode #2854: Rolling Average Steps
// https://leetcode.com/problems/rolling-average-steps/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(k)

import "fmt"

func RollingAverageSteps(steps []int, k int) []float64 {
	n := len(steps)
	if n < k {
		return []float64{}
	}

	result := make([]float64, n-k+1)
	var sum int
	for i := 0; i < k; i++ {
		sum += steps[i]
	}
	result[0] = float64(sum) / float64(k)

	for i := k; i < n; i++ {
		sum += steps[i] - steps[i-k]
		result[i-k+1] = float64(sum) / float64(k)
	}

	return result
}

func main() {
	fmt.Println(RollingAverageSteps([]int{1, 2, 3, 4, 5}, 3))
	fmt.Println(RollingAverageSteps([]int{10, 20}, 2))
}
```

## 2856 — Minimum Array Length After Pair Removals

```go
package main

// LeetCode #2856: Minimum Array Length After Pair Removals
// https://leetcode.com/problems/minimum-array-length-after-pair-removals/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func MinimumArrayLengthAfterPairRemovals(nums []int) int {
	n := len(nums)
	// Find max frequency
	maxFreq := 0
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
		if freq[v] > maxFreq {
			maxFreq = freq[v]
		}
	}

	remaining := n - maxFreq
	if maxFreq > remaining {
		return maxFreq - remaining
	}
	if n%2 == 0 {
		return 0
	}
	return 1
}

func main() {
	fmt.Println(MinimumArrayLengthAfterPairRemovals([]int{1, 1, 2, 2, 3, 3}))
	fmt.Println(MinimumArrayLengthAfterPairRemovals([]int{1, 1, 2, 3}))
}
```

## 2857 — Count Pairs Of Points With Distance K

```go
package main

// LeetCode #2857: Count Pairs of Points With Distance k
// https://leetcode.com/problems/count-pairs-of-points-with-distance-k/
// Difficulty: Medium
// Time: O(n * k) | Space: O(n)

import "fmt"

func CountPairsOfPointsWithDistanceK(coordinates [][]int, k int) int {
	cache := make(map[[2]int]int)
	count := 0

	for _, coord := range coordinates {
		x, y := coord[0], coord[1]
		// XOR property: (x1 ^ x2) + (y1 ^ y2) = k
		// For each possible a, b such that a + b = k:
		for a := 0; a <= k; a++ {
			b := k - a
			px := x ^ a
			py := y ^ b
			count += cache[[2]int{px, py}]
		}
		cache[[2]int{x, y}]++
	}

	return count
}

func main() {
	fmt.Println(CountPairsOfPointsWithDistanceK([][]int{{1, 2}, {4, 2}, {1, 3}, {5, 2}}, 2))
	fmt.Println(CountPairsOfPointsWithDistanceK([][]int{{0, 0}, {1, 1}, {2, 2}}, 2))
}
```

## 2860 — Happy Students

```go
package main

// LeetCode #2860: Happy Students
// https://leetcode.com/problems/happy-students/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func HappyStudents(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	count := 0

	// Check for empty selection (all students are happy)
	if nums[0] != 0 {
		count++
	}

	for i := 0; i < n; i++ {
		selected := i + 1
		// nums[i] < selected: the student at position i needs fewer than selected
		if nums[i] < selected {
			// Check if next student (if exists) needs more than selected
			if i+1 >= n || nums[i+1] > selected {
				count++
			}
		}
	}

	return count
}

func main() {
	fmt.Println(HappyStudents([]int{1, 1}))
	fmt.Println(HappyStudents([]int{6, 0, 3, 3, 6, 7, 2, 7}))
}
```

## 2861 — Maximum Number Of Alloys

```go
package main

// LeetCode #2861: Maximum Number of Alloys
// https://leetcode.com/problems/maximum-number-of-alloys/
// Difficulty: Medium
// Time: O(n * log m) | Space: O(1)

import "fmt"

func MaximumNumberOfAlloys(n int, k int, budget int, composition [][]int, stock []int, cost []int) int {
	canMake := func(count int, comp []int) bool {
		var totalCost int64
		for i := 0; i < n; i++ {
			needed := int64(comp[i]) * int64(count)
			if needed > int64(stock[i]) {
				totalCost += (needed - int64(stock[i])) * int64(cost[i])
				if totalCost > int64(budget) {
					return false
				}
			}
		}
		return totalCost <= int64(budget)
	}

	best := 0
	for _, comp := range composition {
		lo, hi := 0, 1_000_000_000
		for lo < hi {
			mid := (lo + hi + 1) / 2
			if canMake(mid, comp) {
				lo = mid
			} else {
				hi = mid - 1
			}
		}
		if lo > best {
			best = lo
		}
	}

	return best
}

func main() {
	fmt.Println(MaximumNumberOfAlloys(3, 2, 15, [][]int{{1, 1, 1}, {2, 2, 2}}, []int{0, 0, 0}, []int{1, 2, 3}))
	fmt.Println(MaximumNumberOfAlloys(2, 2, 10, [][]int{{1, 2}, {2, 1}}, []int{5, 5}, []int{1, 1}))
}
```

## 2863 — Maximum Length Of Semi Decreasing Subarrays

```go
package main

// LeetCode #2863: Maximum Length of Semi-Decreasing Subarrays
// https://leetcode.com/problems/maximum-length-of-semi-decreasing-subarrays/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func MaximumLengthOfSemiDecreasingSubarrays(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}

	// Store indices where nums[i] > nums[i+1] (start of decreasing)
	starts := make([]int, 0)
	for i := 0; i < n-1; i++ {
		if nums[i] > nums[i+1] {
			starts = append(starts, i)
		}
	}

	if len(starts) == 0 {
		return 1
	}

	best := 1
	for i := 0; i < len(starts); i++ {
		end := starts[i] + 1
		// If not the last start, limit by next start
		limit := n - 1
		if i+1 < len(starts) {
			limit = starts[i+1]
		}
		// Extend to the right while keeping non-increasing property
		for end < limit && nums[end] >= nums[end+1] {
			end++
		}
		length := end - starts[i] + 1
		if length > best {
			best = length
		}
	}

	return best
}

func main() {
	fmt.Println(MaximumLengthOfSemiDecreasingSubarrays([]int{1, 2, 3, 4}))
	fmt.Println(MaximumLengthOfSemiDecreasingSubarrays([]int{7, 6, 5, 4, 3, 2, 1, 6, 10, 11}))
}
```

## 2865 — Beautiful Towers I

```go
package main

// LeetCode #2865: Beautiful Towers I
// https://leetcode.com/problems/beautiful-towers-i/
// Difficulty: Medium
// Time: O(n^2) | Space: O(1)

import "fmt"

func BeautifulTowersI(maxHeights []int) int64 {
	n := len(maxHeights)
	var best int64

	for peak := 0; peak < n; peak++ {
		var total int64 = int64(maxHeights[peak])
		prev := maxHeights[peak]

		// Left side
		for i := peak - 1; i >= 0; i-- {
			h := maxHeights[i]
			if h > prev {
				h = prev
			}
			total += int64(h)
			prev = h
		}

		prev = maxHeights[peak]
		// Right side
		for i := peak + 1; i < n; i++ {
			h := maxHeights[i]
			if h > prev {
				h = prev
			}
			total += int64(h)
			prev = h
		}

		if total > best {
			best = total
		}
	}

	return best
}

func main() {
	fmt.Println(BeautifulTowersI([]int{5, 3, 4, 1, 1}))
	fmt.Println(BeautifulTowersI([]int{6, 5, 3, 9, 2, 7}))
}
```

## 2866 — Beautiful Towers Ii

```go
package main

// LeetCode #2866: Beautiful Towers II
// https://leetcode.com/problems/beautiful-towers-ii/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func BeautifulTowersIi(maxHeights []int) int64 {
	n := len(maxHeights)

	// Left to right: sum of heights ending at i with non-decreasing left side
	left := make([]int64, n)
	stack := make([]int, 0)
	for i := 0; i < n; i++ {
		for len(stack) > 0 && maxHeights[stack[len(stack)-1]] > maxHeights[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			left[i] = int64(maxHeights[i]) * int64(i+1)
		} else {
			prev := stack[len(stack)-1]
			left[i] = left[prev] + int64(maxHeights[i])*int64(i-prev)
		}
		stack = append(stack, i)
	}

	// Right to left
	right := make([]int64, n)
	stack = make([]int, 0)
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && maxHeights[stack[len(stack)-1]] > maxHeights[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			right[i] = int64(maxHeights[i]) * int64(n-i)
		} else {
			prev := stack[len(stack)-1]
			right[i] = right[prev] + int64(maxHeights[i])*int64(prev-i)
		}
		stack = append(stack, i)
	}

	var best int64
	for i := 0; i < n; i++ {
		total := left[i] + right[i] - int64(maxHeights[i])
		if total > best {
			best = total
		}
	}

	return best
}

func main() {
	fmt.Println(BeautifulTowersIi([]int{5, 3, 4, 1, 1}))
	fmt.Println(BeautifulTowersIi([]int{6, 5, 3, 9, 2, 7}))
}
```

## 2870 — Minimum Number Of Operations To Make Array Empty

```go
package main

// LeetCode #2870: Minimum Number of Operations to Make Array Empty
// https://leetcode.com/problems/minimum-number-of-operations-to-make-array-empty/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func MinimumNumberOfOperationsToMakeArrayEmpty(nums []int) int {
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}

	ops := 0
	for _, c := range freq {
		if c == 1 {
			return -1
		}
		// Use as many 3s as possible
		ops += c / 3
		if c%3 != 0 {
			ops++
		}
	}

	return ops
}

func main() {
	fmt.Println(MinimumNumberOfOperationsToMakeArrayEmpty([]int{2, 3, 3, 3, 3, 2}))
	fmt.Println(MinimumNumberOfOperationsToMakeArrayEmpty([]int{1, 1, 1, 1}))
	fmt.Println(MinimumNumberOfOperationsToMakeArrayEmpty([]int{1, 2, 3}))
}
```

## 2871 — Split Array Into Maximum Number Of Subarrays

```go
package main

// LeetCode #2871: Split Array Into Maximum Number of Subarrays
// https://leetcode.com/problems/split-array-into-maximum-number-of-subarrays/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func SplitArrayIntoMaximumNumberOfSubarrays(nums []int) int {
	// We need to split such that AND of each subarray's bitwise AND is minimum
	// Minimum possible AND of any subarray is the AND of entire array
	minAnd := nums[0]
	for _, v := range nums[1:] {
		minAnd &= v
	}

	if minAnd != 0 {
		return 1
	}

	// Count subarrays whose AND equals 0
	count := 0
	curAnd := nums[0]
	for _, v := range nums[1:] {
		if curAnd == 0 {
			count++
			curAnd = v
		} else {
			curAnd &= v
		}
	}
	if curAnd == 0 {
		count++
	}

	return count
}

func main() {
	fmt.Println(SplitArrayIntoMaximumNumberOfSubarrays([]int{1, 0, 2, 0, 1, 0}))
	fmt.Println(SplitArrayIntoMaximumNumberOfSubarrays([]int{1, 2, 3}))
}
```

## 2874 — Maximum Value Of An Ordered Triplet Ii

```go
package main

// LeetCode #2874: Maximum Value of an Ordered Triplet II
// https://leetcode.com/problems/maximum-value-of-an-ordered-triplet-ii/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func MaximumValueOfAnOrderedTripletIi(nums []int) int64 {
	n := len(nums)
	if n < 3 {
		return 0
	}

	var best int64
	maxDiff := int64(0) // max(nums[i] - nums[j]) for i < j
	maxVal := int64(nums[0])

	for k := 1; k < n; k++ {
		// Try nums[k] as the third element
		val := maxDiff * int64(nums[k])
		if val > best {
			best = val
		}

		// Update maxDiff with nums[k] as nums[j]
		diff := maxVal - int64(nums[k])
		if diff > maxDiff {
			maxDiff = diff
		}

		// Update maxVal
		if int64(nums[k]) > maxVal {
			maxVal = int64(nums[k])
		}
	}

	return best
}

func main() {
	fmt.Println(MaximumValueOfAnOrderedTripletIi([]int{12, 6, 1, 2, 7}))
	fmt.Println(MaximumValueOfAnOrderedTripletIi([]int{1, 10, 3, 4, 19}))
}
```

## 2875 — Minimum Size Subarray In Infinite Array

```go
package main

// LeetCode #2875: Minimum Size Subarray in Infinite Array
// https://leetcode.com/problems/minimum-size-subarray-in-infinite-array/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import (
	"fmt"
	"math"
)

func MinimumSizeSubarrayInInfiniteArray(nums []int, target int) int {
	n := len(nums)
	var totalSum int
	for _, v := range nums {
		totalSum += v
	}

	// If target is 0, we need empty subarray
	if target == 0 {
		return 0
	}

	repeats := target / totalSum
	remainder := target % totalSum

	if remainder == 0 {
		return repeats * n
	}

	// Find minimum subarray with sum == remainder in doubled array
	extended := make([]int, n*2)
	copy(extended, nums)
	copy(extended[n:], nums)

	best := math.MaxInt32
	left := 0
	sum := 0

	for right := 0; right < len(extended); right++ {
		sum += extended[right]
		for sum > remainder && left <= right {
			sum -= extended[left]
			left++
		}
		if sum == remainder {
			length := right - left + 1
			if length < best {
				best = length
			}
		}
	}

	if best == math.MaxInt32 {
		return -1
	}

	return repeats*n + best
}

func main() {
	fmt.Println(MinimumSizeSubarrayInInfiniteArray([]int{1, 2, 3}, 5))
	fmt.Println(MinimumSizeSubarrayInInfiniteArray([]int{1, 1, 1}, 4))
}
```

## 2892 — Minimizing Array After Replacing Pairs With Their Product

```go
package main

// LeetCode #2892: Minimizing Array After Replacing Pairs With Their Product
// https://leetcode.com/problems/minimizing-array-after-replacing-pairs-with-their-product/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func MinimizingArrayAfterReplacingPairsWithTheirProduct(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}

	// Greedy: merge adjacent elements whose product is <= their sum
	// Actually, we merge if product > maxElement (since product will grow)
	// For positive integers > 1, product grows quickly
	// If any element is 1, it can always be merged (product = other * 1, sum = other + 1)

	result := make([]int, 0, n)
	for _, v := range nums {
		result = append(result, v)
		for len(result) >= 2 {
			last := len(result) - 1
			// Merge if product <= sum
			p := result[last] * result[last-1]
			s := result[last] + result[last-1]
			if p <= s {
				merged := result[last] * result[last-1]
				result = result[:last-1]
				result = append(result, merged)
			} else {
				break
			}
		}
	}

	return len(result)
}

func main() {
	fmt.Println(MinimizingArrayAfterReplacingPairsWithTheirProduct([]int{2, 3, 3}))
	fmt.Println(MinimizingArrayAfterReplacingPairsWithTheirProduct([]int{1, 2, 1, 2}))
}
```

## 2893 — Calculate Orders Within Each Interval

```go
package main

// LeetCode #2893: Calculate Orders Within Each Interval
// https://leetcode.com/problems/calculate-orders-within-each-interval/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

type Order struct {
	Time   int
	Amount int
}

func CalculateOrdersWithinEachInterval(orders []Order, interval int) []int {
	if len(orders) == 0 {
		return []int{}
	}

	// Group orders by interval
	maxTime := orders[len(orders)-1].Time
	bucketCount := maxTime/interval + 1
	buckets := make([]int, bucketCount)

	for _, o := range orders {
		idx := o.Time / interval
		buckets[idx] += o.Amount
	}

	return buckets
}

func main() {
	orders := []Order{
		{0, 10}, {1, 20}, {4, 30}, {6, 40},
	}
	fmt.Println(CalculateOrdersWithinEachInterval(orders, 3))

	orders2 := []Order{
		{0, 5}, {2, 10},
	}
	fmt.Println(CalculateOrdersWithinEachInterval(orders2, 5))
}
```

## 2895 — Minimum Processing Time

```go
package main

// LeetCode #2895: Minimum Processing Time
// https://leetcode.com/problems/minimum-processing-time/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func MinimumProcessingTime(processorTime []int, taskTime []int) int {
	sort.Ints(processorTime)
	sort.Ints(taskTime)

	// Each processor handles 4 tasks
	n := len(processorTime)
	maxTime := 0

	for i := 0; i < n; i++ {
		// Assign 4 largest remaining tasks to the slowest processor
		// Processor runs concurrently, so time = processorTime[i] + max of its 4 tasks
		taskIdx := len(taskTime) - 1 - i*4
		time := processorTime[i] + taskTime[taskIdx]
		if time > maxTime {
			maxTime = time
		}
	}

	return maxTime
}

func main() {
	fmt.Println(MinimumProcessingTime([]int{8, 10}, []int{2, 2, 3, 1, 8, 7, 4, 5}))
	fmt.Println(MinimumProcessingTime([]int{10, 20}, []int{2, 3, 1, 2, 5, 8, 4, 3}))
}
```

## 2896 — Apply Operations To Make Two Strings Equal

```go
package main

// LeetCode #2896: Apply Operations to Make Two Strings Equal
// https://leetcode.com/problems/apply-operations-to-make-two-strings-equal/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func ApplyOperationsToMakeTwoStringsEqual(s1 string, s2 string, x int) int {
	// Find positions where characters differ
	diff := make([]int, 0)
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			diff = append(diff, i)
		}
	}

	if len(diff)%2 != 0 {
		return -1
	}
	if len(diff) == 0 {
		return 0
	}

	m := len(diff)
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, m)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	var solve func(l, r int) int
	solve = func(l, r int) int {
		if l > r {
			return 0
		}
		if dp[l][r] != -1 {
			return dp[l][r]
		}

		// Option 1: flip s[l] and s[l+1] (cost = diff[l+1] - diff[l])
		best := solve(l+2, r) + diff[l+1] - diff[l]

		// Option 2: use operation with cost x
		cost := solve(l+1, r-1) + x
		if cost < best {
			best = cost
		}

		dp[l][r] = best
		return best
	}

	return solve(0, m-1)
}

func main() {
	fmt.Println(ApplyOperationsToMakeTwoStringsEqual("1100011000", "0101001010", 2))
	fmt.Println(ApplyOperationsToMakeTwoStringsEqual("10110", "00011", 4))
}
```

## 2898 — Maximum Linear Stock Score

```go
package main

// LeetCode #2898: Maximum Linear Stock Score
// https://leetcode.com/problems/maximum-linear-stock-score/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func MaximumLinearStockScore(prices []int) int64 {
	// For each stock, score = sum of prices where prices[i] - i is same
	score := make(map[int]int64)
	var best int64

	for i, p := range prices {
		key := p - i
		score[key] += int64(p)
		if score[key] > best {
			best = score[key]
		}
	}

	return best
}

func main() {
	fmt.Println(MaximumLinearStockScore([]int{1, 2, 3, 4}))
	fmt.Println(MaximumLinearStockScore([]int{2, 1, 3}))
}
```

## 2901 — Longest Unequal Adjacent Groups Subsequence Ii

```go
package main

// LeetCode #2901: Longest Unequal Adjacent Groups Subsequence II
// https://leetcode.com/problems/longest-unequal-adjacent-groups-subsequence-ii/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import "fmt"

func LongestUnequalAdjacentGroupsSubsequenceIi(words []string, groups []int) []string {
	n := len(words)
	dp := make([]int, n)
	prev := make([]int, n)
	for i := range prev {
		prev[i] = -1
	}

	hamming := func(a, b string) int {
		if len(a) != len(b) {
			return -1
		}
		diff := 0
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				diff++
			}
		}
		return diff
	}

	bestLen := 0
	bestIdx := 0

	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if groups[j] != groups[i] && hamming(words[j], words[i]) == 1 {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					prev[i] = j
				}
			}
		}
		if dp[i] > bestLen {
			bestLen = dp[i]
			bestIdx = i
		}
	}

	result := make([]string, bestLen)
	for i := bestLen - 1; i >= 0; i-- {
		result[i] = words[bestIdx]
		bestIdx = prev[bestIdx]
	}

	return result
}

func main() {
	fmt.Println(LongestUnequalAdjacentGroupsSubsequenceIi([]string{"bab", "dab", "cab"}, []int{1, 2, 2}))
	fmt.Println(LongestUnequalAdjacentGroupsSubsequenceIi([]string{"a", "b", "c", "d"}, []int{1, 2, 3, 4}))
}
```

## 2904 — Shortest And Lexicographically Smallest Beautiful String

```go
package main

// LeetCode #2904: Shortest and Lexicographically Smallest Beautiful String
// https://leetcode.com/problems/shortest-and-lexicographically-smallest-beautiful-string/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(shortestBeautifulSubstring("100011001", 3))
	fmt.Println(shortestBeautifulSubstring("1011", 2))
	fmt.Println(shortestBeautifulSubstring("000", 1))
}

func shortestBeautifulSubstring(s string, k int) string {
	n := len(s)
	ans := ""
	for i := 0; i < n; i++ {
		for j := i + k; j <= n; j++ {
			t := s[i:j]
			cnt := 0
			for _, c := range t {
				if c == '1' {
					cnt++
				}
			}
			if cnt == k && (ans == "" || j-i < len(ans) || (j-i == len(ans) && t < ans)) {
				ans = t
			}
		}
	}
	return ans
}
```

## 2905 — Find Indices With Index And Value Difference Ii

```go
package main

// LeetCode #2905: Find Indices With Index and Value Difference II
// https://leetcode.com/problems/find-indices-with-index-and-value-difference-ii/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(findIndices([]int{5, 1, 4, 1}, 2, 4))
	fmt.Println(findIndices([]int{2, 1}, 0, 0))
	fmt.Println(findIndices([]int{1, 2, 3}, 2, 4))
}

func findIndices(nums []int, indexDifference int, valueDifference int) []int {
	mi, mx := 0, 0
	for i := indexDifference; i < len(nums); i++ {
		j := i - indexDifference
		if nums[j] < nums[mi] {
			mi = j
		}
		if nums[j] > nums[mx] {
			mx = j
		}
		if nums[i]-nums[mi] >= valueDifference {
			return []int{mi, i}
		}
		if nums[mx]-nums[i] >= valueDifference {
			return []int{mx, i}
		}
	}
	return []int{-1, -1}
}
```

## 2906 — Construct Product Matrix

```go
package main

// LeetCode #2906: Construct Product Matrix
// https://leetcode.com/problems/construct-product-matrix/
// Difficulty: Medium
// Time: O(m*n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(constructProductMatrix([][]int{{1, 2}, {3, 4}}))
	fmt.Println(constructProductMatrix([][]int{{2, 3, 4}, {5, 6, 7}}))
}

func constructProductMatrix(grid [][]int) [][]int {
	const mod int = 12345
	n, m := len(grid), len(grid[0])
	p := make([][]int, n)
	for i := range p {
		p[i] = make([]int, m)
	}
	suf := 1
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			p[i][j] = suf
			suf = suf * grid[i][j] % mod
		}
	}
	pre := 1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			p[i][j] = p[i][j] * pre % mod
			pre = pre * grid[i][j] % mod
		}
	}
	return p
}
```

## 2907 — Maximum Profitable Triplets With Increasing Prices I

```go
package main

// LeetCode #2907: Maximum Profitable Triplets With Increasing Prices I
// https://leetcode.com/problems/maximum-profitable-triplets-with-increasing-prices-i/
// Difficulty: Medium [Paid]
// Time: O(n^2) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{10, 20, 30}, []int{1, 2, 3}))
	fmt.Println(maxProfit([]int{1, 2, 3, 4}, []int{5, 4, 3, 2}))
	fmt.Println(maxProfit([]int{3, 2, 1}, []int{10, 20, 30}))
}

func maxProfit(prices []int, profits []int) int {
	n := len(prices)
	ans := -1
	for j, x := range profits {
		left, right := 0, 0
		for i := 0; i < j; i++ {
			if prices[i] < prices[j] {
				if profits[i] > left {
					left = profits[i]
				}
			}
		}
		for k := j + 1; k < n; k++ {
			if prices[j] < prices[k] {
				if profits[k] > right {
					right = profits[k]
				}
			}
		}
		if left > 0 && right > 0 {
			if left+x+right > ans {
				ans = left + x + right
			}
		}
	}
	return ans
}
```

## 2909 — Minimum Sum Of Mountain Triplets Ii

```go
package main

// LeetCode #2909: Minimum Sum of Mountain Triplets II
// https://leetcode.com/problems/minimum-sum-of-mountain-triplets-ii/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(minimumSum([]int{8, 6, 1, 5, 3}))
	fmt.Println(minimumSum([]int{5, 4, 8, 7, 10, 2}))
	fmt.Println(minimumSum([]int{6, 5, 4, 3, 4, 5}))
}

func minimumSum(nums []int) int {
	n := len(nums)
	const inf = 1 << 30
	right := make([]int, n+1)
	right[n] = inf
	for i := n - 1; i >= 0; i-- {
		if right[i+1] < nums[i] {
			right[i] = right[i+1]
		} else {
			right[i] = nums[i]
		}
	}
	ans, left := inf, inf
	for i, x := range nums {
		if left < x && right[i+1] < x {
			sum := left + x + right[i+1]
			if sum < ans {
				ans = sum
			}
		}
		if x < left {
			left = x
		}
	}
	if ans == inf {
		return -1
	}
	return ans
}
```

## 2910 — Minimum Number Of Groups To Create A Valid Assignment

```go
package main

// LeetCode #2910: Minimum Number of Groups to Create a Valid Assignment
// https://leetcode.com/problems/minimum-number-of-groups-to-create-a-valid-assignment/
// Difficulty: Medium
// Time: O(n + m*minFreq) | Space: O(m)

import "fmt"

func main() {
	fmt.Println(minGroupsForValidAssignment([]int{3, 3, 3, 3, 3, 1, 1}))
	fmt.Println(minGroupsForValidAssignment([]int{10, 10, 10, 10, 10}))
	fmt.Println(minGroupsForValidAssignment([]int{1, 1, 1, 2, 2, 2}))
}

func minGroupsForValidAssignment(nums []int) int {
	cnt := map[int]int{}
	for _, x := range nums {
		cnt[x]++
	}
	k := len(nums)
	for _, v := range cnt {
		if v < k {
			k = v
		}
	}
	for ; ; k-- {
		ans := 0
		ok := true
		for _, v := range cnt {
			if v/k < v%k {
				ok = false
				break
			}
			ans += (v + k) / (k + 1)
		}
		if ok {
			return ans
		}
	}
}
```

## 2914 — Minimum Number Of Changes To Make Binary String Beautiful

```go
package main

// LeetCode #2914: Minimum Number of Changes to Make Binary String Beautiful
// https://leetcode.com/problems/minimum-number-of-changes-to-make-binary-string-beautiful/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(minChanges("1001"))
	fmt.Println(minChanges("10"))
	fmt.Println(minChanges("0000"))
}

func minChanges(s string) (ans int) {
	for i := 1; i < len(s); i += 2 {
		if s[i] != s[i-1] {
			ans++
		}
	}
	return
}
```

## 2915 — Length Of The Longest Subsequence That Sums To Target

```go
package main

// LeetCode #2915: Length of the Longest Subsequence That Sums to Target
// https://leetcode.com/problems/length-of-the-longest-subsequence-that-sums-to-target/
// Difficulty: Medium
// Time: O(n*target) | Space: O(n*target)

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubsequence([]int{1, 2, 3, 4, 5}, 9))
	fmt.Println(lengthOfLongestSubsequence([]int{4, 1, 3, 2, 1, 5}, 7))
	fmt.Println(lengthOfLongestSubsequence([]int{1, 1, 5, 4, 5}, 3))
}

func lengthOfLongestSubsequence(nums []int, target int) int {
	n := len(nums)
	const negInf = -(1 << 30)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, target+1)
		for j := range f[i] {
			f[i][j] = negInf
		}
	}
	f[0][0] = 0
	for i := 1; i <= n; i++ {
		x := nums[i-1]
		for j := 0; j <= target; j++ {
			f[i][j] = f[i-1][j]
			if j >= x && f[i-1][j-x]+1 > f[i][j] {
				f[i][j] = f[i-1][j-x] + 1
			}
		}
	}
	if f[n][target] <= 0 {
		return -1
	}
	return f[n][target]
}
```

## 2918 — Minimum Equal Sum Of Two Arrays After Replacing Zeros

```go
package main

// LeetCode #2918: Minimum Equal Sum of Two Arrays After Replacing Zeros
// https://leetcode.com/problems/minimum-equal-sum-of-two-arrays-after-replacing-zeros/
// Difficulty: Medium
// Time: O(n+m) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(minSum([]int{3, 2, 0, 1, 0}, []int{6, 5, 0}))
	fmt.Println(minSum([]int{2, 0, 2, 0}, []int{1, 4}))
	fmt.Println(minSum([]int{1, 2}, []int{3, 4}))
}

func minSum(nums1 []int, nums2 []int) int64 {
	s1, s2 := int64(0), int64(0)
	z1, z2 := 0, 0
	for _, x := range nums1 {
		if x == 0 {
			z1++
		} else {
			s1 += int64(x)
		}
	}
	for _, x := range nums2 {
		if x == 0 {
			z2++
		} else {
			s2 += int64(x)
		}
	}
	min1, min2 := s1+int64(z1), s2+int64(z2)
	if min1 < min2 && z1 == 0 {
		return -1
	}
	if min2 < min1 && z2 == 0 {
		return -1
	}
	if min1 > min2 {
		return min1
	}
	return min2
}
```

## 2919 — Minimum Increment Operations To Make Array Beautiful

```go
package main

// LeetCode #2919: Minimum Increment Operations to Make Array Beautiful
// https://leetcode.com/problems/minimum-increment-operations-to-make-array-beautiful/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(minIncrementOperations([]int{2, 3, 0, 0, 2}, 4))
	fmt.Println(minIncrementOperations([]int{0, 1, 3, 3}, 5))
	fmt.Println(minIncrementOperations([]int{1, 1, 2}, 1))
}

func minIncrementOperations(nums []int, k int) int64 {
	f, g, h := int64(0), int64(0), int64(0)
	for _, x := range nums {
		add := int64(0)
		if x < k {
			add = int64(k - x)
		}
		f, g, h = g, h, min64(f, min64(g, h))+add
	}
	return min64(f, min64(g, h))
}

func min64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
```

## 2922 — Market Analysis Iii

```go
package main

// LeetCode #2922: Market Analysis III
// https://leetcode.com/problems/market-analysis-iii/
// Difficulty: Medium (SQL problem — implemented in Go)
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"sort"
)

type User struct {
	SellerID      int
	JoinDate      string
	FavoriteBrand string
}

type Item struct {
	ItemID    int
	ItemBrand string
}

type Order struct {
	OrderID   int
	OrderDate string
	ItemID    int
	SellerID  int
}

type SellerResult struct {
	SellerID int
	NumItems int
}

func marketAnalysisIII(users []User, items []Item, orders []Order) []SellerResult {
	// Build item_id -> item_brand map
	itemBrand := make(map[int]string)
	for _, item := range items {
		itemBrand[item.ItemID] = item.ItemBrand
	}

	// Build seller_id -> favorite_brand map
	sellerBrand := make(map[int]string)
	for _, u := range users {
		sellerBrand[u.SellerID] = u.FavoriteBrand
	}

	// For each seller, count distinct items where item_brand != favorite_brand
	sellerCounts := make(map[int]map[int]bool) // seller_id -> set of item_ids
	for _, o := range orders {
		brand, ok := itemBrand[o.ItemID]
		if !ok {
			continue
		}
		fav, ok := sellerBrand[o.SellerID]
		if !ok {
			continue
		}
		if brand != fav {
			if sellerCounts[o.SellerID] == nil {
				sellerCounts[o.SellerID] = make(map[int]bool)
			}
			sellerCounts[o.SellerID][o.ItemID] = true
		}
	}

	// Find max count
	maxCount := 0
	for _, itemsSet := range sellerCounts {
		if len(itemsSet) > maxCount {
			maxCount = len(itemsSet)
		}
	}

	// Collect sellers with max count
	var results []SellerResult
	for sid, itemsSet := range sellerCounts {
		if len(itemsSet) == maxCount {
			results = append(results, SellerResult{SellerID: sid, NumItems: len(itemsSet)})
		}
	}

	// Order by seller_id ASC
	sort.Slice(results, func(i, j int) bool {
		return results[i].SellerID < results[j].SellerID
	})

	return results
}

func main() {
	users := []User{
		{SellerID: 1, JoinDate: "2024-01-01", FavoriteBrand: "Apple"},
		{SellerID: 2, JoinDate: "2024-01-02", FavoriteBrand: "Samsung"},
		{SellerID: 3, JoinDate: "2024-01-03", FavoriteBrand: "Google"},
	}

	items := []Item{
		{ItemID: 1, ItemBrand: "Apple"},
		{ItemID: 2, ItemBrand: "Samsung"},
		{ItemID: 3, ItemBrand: "Google"},
		{ItemID: 4, ItemBrand: "Apple"},
	}

	orders := []Order{
		{OrderID: 1, OrderDate: "2024-02-01", ItemID: 2, SellerID: 1},  // seller1 buys Samsung != Apple
		{OrderID: 2, OrderDate: "2024-02-02", ItemID: 3, SellerID: 1},  // seller1 buys Google != Apple
		{OrderID: 3, OrderDate: "2024-02-03", ItemID: 1, SellerID: 2},  // seller2 buys Apple != Samsung
		{OrderID: 4, OrderDate: "2024-02-04", ItemID: 4, SellerID: 2},  // seller2 buys Apple != Samsung
		{OrderID: 5, OrderDate: "2024-02-05", ItemID: 3, SellerID: 2},  // seller2 buys Google != Samsung
		{OrderID: 6, OrderDate: "2024-02-06", ItemID: 1, SellerID: 3},  // seller3 buys Apple != Google
		{OrderID: 7, OrderDate: "2024-02-07", ItemID: 2, SellerID: 3},  // seller3 buys Samsung != Google
	}

	fmt.Println("Market Analysis III")
	fmt.Println("===================")
	fmt.Printf("%-12s %s\n", "Seller ID", "Num Items")
	fmt.Println("---------------------")

	results := marketAnalysisIII(users, items, orders)
	for _, r := range results {
		fmt.Printf("%-12d %d\n", r.SellerID, r.NumItems)
	}
}
```

## 2924 — Find Champion Ii

```go
package main

// LeetCode #2924: Find Champion II
// https://leetcode.com/problems/find-champion-ii/
// Difficulty: Medium
// Time: O(n+m) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(findChampionII(3, [][]int{{0, 1}, {1, 2}}))
	fmt.Println(findChampionII(4, [][]int{{0, 2}, {1, 3}, {1, 2}}))
	fmt.Println(findChampionII(2, [][]int{{0, 1}}))
}

func findChampionII(n int, edges [][]int) int {
	indeg := make([]int, n)
	for _, e := range edges {
		indeg[e[1]]++
	}
	ans, cnt := -1, 0
	for i, x := range indeg {
		if x == 0 {
			cnt++
			ans = i
		}
	}
	if cnt == 1 {
		return ans
	}
	return -1
}
```

## 2925 — Maximum Score After Applying Operations On A Tree

```go
package main

// LeetCode #2925: Maximum Score After Applying Operations on a Tree
// https://leetcode.com/problems/maximum-score-after-applying-operations-on-a-tree/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(maximumScoreAfterOperations([][]int{{0, 1}, {0, 2}, {0, 3}}, []int{1, 2, 3, 4}))
	fmt.Println(maximumScoreAfterOperations([][]int{{0, 1}}, []int{10, 20}))
}

func maximumScoreAfterOperations(edges [][]int, values []int) int64 {
	g := make([][]int, len(values))
	for _, e := range edges {
		a, b := e[0], e[1]
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}
	var dfs func(int, int) (int64, int64)
	dfs = func(i, fa int) (int64, int64) {
		a, b := int64(0), int64(0)
		leaf := true
		for _, j := range g[i] {
			if j != fa {
				leaf = false
				aa, bb := dfs(j, i)
				a += aa
				b += bb
			}
		}
		if leaf {
			return int64(values[i]), 0
		}
		va := int64(values[i]) + a
		vb := b + int64(values[i])
		if a > vb {
			vb = a
		}
		return va, vb
	}
	_, b := dfs(0, -1)
	return b
}
```

## 2929 — Distribute Candies Among Children Ii

```go
package main

// LeetCode #2929: Distribute Candies Among Children II
// https://leetcode.com/problems/distribute-candies-among-children-ii/
// Difficulty: Medium
// Time: O(1) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(distributeCandies(5, 2))
	fmt.Println(distributeCandies(3, 3))
	fmt.Println(distributeCandies(10, 3))
}

func distributeCandies(n int, limit int) int64 {
	if n > 3*limit {
		return 0
	}
	ans := int64(n+2) * int64(n+1) / 2
	if n > limit {
		ans -= 3 * int64(n-limit+1) * int64(n-limit) / 2
	}
	if n-2 >= 2*limit {
		ans += 3 * int64(n-2*limit) * int64(n-2*limit+1) / 2
	}
	return ans
}
```

## 2930 — Number Of Strings Which Can Be Rearranged To Contain Substring

```go
package main

// LeetCode #2930: Number of Strings Which Can Be Rearranged to Contain Substring
// https://leetcode.com/problems/number-of-strings-which-can-be-rearranged-to-contain-substring/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(stringCount(4))
	fmt.Println(stringCount(10))
}

func stringCount(n int) int {
	const mod int = 1e9 + 7
	f := make([][2][3][2]int, n+1)
	for i := range f {
		for j := range f[i] {
			for k := range f[i][j] {
				for l := range f[i][j][k] {
					f[i][j][k][l] = -1
				}
			}
		}
	}
	var dfs func(i, l, e, t int) int
	dfs = func(i, l, e, t int) int {
		if i == 0 {
			if l == 1 && e == 2 && t == 1 {
				return 1
			}
			return 0
		}
		if f[i][l][e][t] == -1 {
			a := dfs(i-1, l, e, t) * 23 % mod
			b := dfs(i-1, min2(1, l+1), e, t)
			c := dfs(i-1, l, min2(2, e+1), t)
			d := dfs(i-1, l, e, min2(1, t+1))
			f[i][l][e][t] = (a + b + c + d) % mod
		}
		return f[i][l][e][t]
	}
	return dfs(n, 0, 0, 0)
}

func min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```

## 2933 — High Access Employees

```go
package main

// LeetCode #2933: High-Access Employees
// https://leetcode.com/problems/high-access-employees/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(findHighAccessEmployees([][]string{{"a","0549"},{"b","0456"},{"a","1234"},{"a","1440"}}))
	fmt.Println(findHighAccessEmployees([][]string{{"d","0002"},{"d","0003"},{"d","0004"},{"d","0005"}}))
}

func findHighAccessEmployees(accessTimes [][]string) (ans []string) {
	d := map[string][]int{}
	for _, e := range accessTimes {
		name, s := e[0], e[1]
		h, _ := strconv.Atoi(s[:2])
		m, _ := strconv.Atoi(s[2:])
		t := h*60 + m
		d[name] = append(d[name], t)
	}
	for name, ts := range d {
		sort.Ints(ts)
		for i := 2; i < len(ts); i++ {
			if ts[i]-ts[i-2] < 60 {
				ans = append(ans, name)
				break
			}
		}
	}
	return
}
```

## 2934 — Minimum Operations To Maximize Last Elements In Arrays

```go
package main

// LeetCode #2934: Minimum Operations to Maximize Last Elements in Arrays
// https://leetcode.com/problems/minimum-operations-to-maximize-last-elements-in-arrays/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(minOperationsArrays([]int{1, 2, 7}, []int{4, 5, 3}))
	fmt.Println(minOperationsArrays([]int{2, 3, 4, 5, 9}, []int{8, 7, 6, 1, 2}))
	fmt.Println(minOperationsArrays([]int{1, 5, 4}, []int{2, 5, 3}))
}

func minOperationsArrays(nums1 []int, nums2 []int) int {
	n := len(nums1)
	f := func(x, y int) (cnt int) {
		for i, a := range nums1[:n-1] {
			b := nums2[i]
			if a <= x && b <= y {
				continue
			}
			if !(a <= y && b <= x) {
				return -1
			}
			cnt++
		}
		return
	}
	a, b := f(nums1[n-1], nums2[n-1]), f(nums2[n-1], nums1[n-1])
	if a+b == -2 {
		return -1
	}
	if a < 0 {
		a = 1 << 30
	}
	if b < 0 {
		b = 1 << 30
	}
	if a < b+1 {
		return a
	}
	return b + 1
}
```

## 2936 — Number Of Equal Numbers Blocks

```go
package main

// LeetCode #2936: Number of Equal Numbers Blocks
// https://leetcode.com/problems/number-of-equal-numbers-blocks/
// Difficulty: Medium

import "fmt"

type BigArray struct {
	data []int
}

func (b *BigArray) at(index int) int {
	return b.data[index]
}

func (b *BigArray) size() int {
	return len(b.data)
}

func countBlocks(arr *BigArray) int {
	n := arr.size()
	if n == 0 {
		return 0
	}

	blocks := 0
	i := 0

	for i < n {
		blocks++
		val := arr.at(i)

		// Binary search to find the rightmost position where arr.at(pos) == val
		left, right := i, n-1
		for left < right {
			mid := (left + right + 1) / 2
			if arr.at(mid) == val {
				left = mid
			} else {
				right = mid - 1
			}
		}

		i = left + 1
	}

	return blocks
}

func main() {
	// Test case 1: [1,1,1,2,2,3,3,3,3] -> 3 blocks
	arr1 := &BigArray{data: []int{1, 1, 1, 2, 2, 3, 3, 3, 3}}
	fmt.Println(countBlocks(arr1)) // 3

	// Test case 2: [5,5,5,5,5] -> 1 block
	arr2 := &BigArray{data: []int{5, 5, 5, 5, 5}}
	fmt.Println(countBlocks(arr2)) // 1

	// Test case 3: [1,2,3,4,5] -> 5 blocks
	arr3 := &BigArray{data: []int{1, 2, 3, 4, 5}}
	fmt.Println(countBlocks(arr3)) // 5

	// Test case 4: empty array -> 0 blocks
	arr4 := &BigArray{data: []int{}}
	fmt.Println(countBlocks(arr4)) // 0

	// Test case 5: single element -> 1 block
	arr5 := &BigArray{data: []int{42}}
	fmt.Println(countBlocks(arr5)) // 1
}

// Time: O(k log n) | Space: O(1) where k = number of blocks
```

## 2938 — Separate Black And White Balls

```go
package main

// LeetCode #2938: Separate Black and White Balls
// https://leetcode.com/problems/separate-black-and-white-balls/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(minimumSteps("101"))
	fmt.Println(minimumSteps("100"))
	fmt.Println(minimumSteps("0111"))
}

func minimumSteps(s string) (ans int64) {
	n := len(s)
	cnt := 0
	for i := n - 1; i >= 0; i-- {
		if s[i] == '1' {
			cnt++
			ans += int64(n - i - cnt)
		}
	}
	return
}
```

## 2939 — Maximum Xor Product

```go
package main

// LeetCode #2939: Maximum Xor Product
// https://leetcode.com/problems/maximum-xor-product/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(maximumXorProduct(12, 5, 4))
	fmt.Println(maximumXorProduct(6, 7, 5))
	fmt.Println(maximumXorProduct(1, 6, 3))
}

func maximumXorProduct(a int64, b int64, n int) int {
	const mod int64 = 1e9 + 7
	ax := (a >> n) << n
	bx := (b >> n) << n
	for i := n - 1; i >= 0; i-- {
		x, y := (a>>i)&1, (b>>i)&1
		if x == y {
			ax |= 1 << i
			bx |= 1 << i
		} else if ax < bx {
			ax |= 1 << i
		} else {
			bx |= 1 << i
		}
	}
	ax %= mod
	bx %= mod
	return int(ax * bx % mod)
}
```

## 2943 — Maximize Area Of Square Hole In Grid

```go
package main

// LeetCode #2943: Maximize Area of Square Hole in Grid
// https://leetcode.com/problems/maximize-area-of-square-hole-in-grid/
// Difficulty: Medium
// Time: O(h log h + v log v) | Space: O(1)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximizeSquareArea(3, 4, []int{2}, []int{2}))
	fmt.Println(maximizeSquareArea(2, 2, []int{1}, []int{1}))
}

func maximizeSquareArea(m int, n int, hBars []int, vBars []int) int {
	calc := func(nums []int, limit int) int {
		nums = append(nums, 1)
		nums = append(nums, limit)
		sort.Ints(nums)
		ans, cnt := 1, 1
		for i := 1; i < len(nums); i++ {
			if nums[i] == nums[i-1]+1 {
				cnt++
				if cnt > ans {
					ans = cnt
				}
			} else {
				cnt = 1
			}
		}
		return ans
	}
	x := calc(hBars, m)
	y := calc(vBars, n)
	if x > y {
		x = y
	}
	return x * x
}
```

## 2944 — Minimum Number Of Coins For Fruits

```go
package main

// LeetCode #2944: Minimum Number of Coins for Fruits
// https://leetcode.com/problems/minimum-number-of-coins-for-fruits/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(minimumCoins([]int{3, 1, 2}))
	fmt.Println(minimumCoins([]int{1, 10, 1, 1}))
	fmt.Println(minimumCoins([]int{26, 53, 10, 24, 25, 20, 63, 51}))
}

func minimumCoins(prices []int) int {
	n := len(prices)
	f := make([]int, n+1)
	var dfs func(int) int
	dfs = func(i int) int {
		if i*2 >= n {
			return prices[i-1]
		}
		if f[i] == 0 {
			f[i] = 1 << 30
			for j := i + 1; j <= i*2+1; j++ {
				cost := dfs(j) + prices[i-1]
				if cost < f[i] {
					f[i] = cost
				}
			}
		}
		return f[i]
	}
	return dfs(1)
}
```

## 2947 — Count Beautiful Substrings I

```go
package main

// LeetCode #2947: Count Beautiful Substrings I
// https://leetcode.com/problems/count-beautiful-substrings-i/
// Difficulty: Medium
// Time: O(n^2) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(beautifulSubstrings("baeyh", 2))
	fmt.Println(beautifulSubstrings("ab", 1))
	fmt.Println(beautifulSubstrings("a", 1))
}

func beautifulSubstrings(s string, k int) (ans int) {
	n := len(s)
	vowels := [26]bool{}
	for _, c := range "aeiou" {
		vowels[c-'a'] = true
	}
	for i := 0; i < n; i++ {
		v := 0
		for j := i; j < n; j++ {
			if vowels[s[j]-'a'] {
				v++
			}
			c := j - i + 1 - v
			if v == c && v*c%k == 0 {
				ans++
			}
		}
	}
	return
}
```

## 2948 — Make Lexicographically Smallest Array By Swapping Elements

```go
package main

// LeetCode #2948: Make Lexicographically Smallest Array by Swapping Elements
// https://leetcode.com/problems/make-lexicographically-smallest-array-by-swapping-elements/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(lexicographicallySmallestArray([]int{1, 5, 3, 9, 8}, 2))
	fmt.Println(lexicographicallySmallestArray([]int{1, 7, 6, 18, 2, 1}, 3))
	fmt.Println(lexicographicallySmallestArray([]int{1, 2, 3}, 0))
}

func lexicographicallySmallestArray(nums []int, limit int) []int {
	n := len(nums)
	idx := make([]int, n)
	for i := range idx {
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool {
		return nums[idx[i]] < nums[idx[j]]
	})
	ans := make([]int, n)
	for i := 0; i < n; {
		j := i + 1
		for j < n && nums[idx[j]]-nums[idx[j-1]] <= limit {
			j++
		}
		t := make([]int, j-i)
		copy(t, idx[i:j])
		sort.Ints(t)
		for k := i; k < j; k++ {
			ans[t[k-i]] = nums[idx[k]]
		}
		i = j
	}
	return ans
}
```

## 2950 — Number Of Divisible Substrings

```go
package main

// LeetCode #2950: Number of Divisible Substrings
// https://leetcode.com/problems/number-of-divisible-substrings/
// Difficulty: Medium

import "fmt"

func numberOfDivisibleSubstrings(s string) int {
	n := len(s)

	// Build mapping: old phone keypad, each char maps to a digit
	// a,b,c->1; d,e,f->2; g,h,i->3; j,k,l->4; m,n,o->5;
	// p,q,r->6; s,t,u->7; v,w,x->8; y,z->9
	mapping := make([]int, 26)
	for i := 0; i < 26; i++ {
		if i <= 2 { // a,b,c -> 1
			mapping[i] = 1
		} else if i <= 5 { // d,e,f -> 2
			mapping[i] = 2
		} else if i <= 8 { // g,h,i -> 3
			mapping[i] = 3
		} else if i <= 11 { // j,k,l -> 4
			mapping[i] = 4
		} else if i <= 14 { // m,n,o -> 5
			mapping[i] = 5
		} else if i <= 17 { // p,q,r -> 6
			mapping[i] = 6
		} else if i <= 20 { // s,t,u -> 7
			mapping[i] = 7
		} else if i <= 23 { // v,w,x -> 8
			mapping[i] = 8
		} else { // y,z -> 9
			mapping[i] = 9
		}
	}

	count := 0
	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += mapping[s[j]-'a']
			length := j - i + 1
			if sum%length == 0 {
				count++
			}
		}
	}

	return count
}

func main() {
	// Test case 1: "abc"
	// a=1, b=1, c=2
	// Substrings: "a"(1%1=0), "b"(1%1=0), "c"(2%1=0),
	// "ab"(2%2=0), "bc"(3%2!=0), "abc"(4%3!=0) -> 4
	fmt.Println(numberOfDivisibleSubstrings("abc")) // 4

	// Test case 2: single character
	fmt.Println(numberOfDivisibleSubstrings("a")) // 1

	// Test case 3: "abcd"
	// a=1,b=1,c=2,d=2
	// "a"(1), "b"(1), "c"(2), "d"(2),
	// "ab"(2%2=0), "bc"(3%2!=0), "cd"(4%2=0),
	// "abc"(4%3!=0), "bcd"(5%3!=0),
	// "abcd"(6%4!=0)
	fmt.Println(numberOfDivisibleSubstrings("abcd")) // 6

	// Test case 4: empty-ish (single char repeated)
	fmt.Println(numberOfDivisibleSubstrings("aaa")) // 6
	// "a"(1), "a"(1), "a"(1), "aa"(2%2=0), "aa"(2%2=0), "aaa"(3%3=0) -> 6
}

// Time: O(n^2) | Space: O(1)
```

## 2952 — Minimum Number Of Coins To Be Added

```go
package main

// LeetCode #2952: Minimum Number of Coins to be Added
// https://leetcode.com/problems/minimum-number-of-coins-to-be-added/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minimumAddedCoins([]int{1, 4, 10}, 19))
	fmt.Println(minimumAddedCoins([]int{1, 4, 10, 5, 7, 19}, 19))
	fmt.Println(minimumAddedCoins([]int{1, 1, 1}, 20))
}

func minimumAddedCoins(coins []int, target int) (ans int) {
	sort.Ints(coins)
	for i, s := 0, 1; s <= target; {
		if i < len(coins) && coins[i] <= s {
			s += coins[i]
			i++
		} else {
			s <<= 1
			ans++
		}
	}
	return
}
```

## 2955 — Number Of Same End Substrings

```go
package main

// LeetCode #2955: Number of Same-End Substrings
// https://leetcode.com/problems/number-of-same-end-substrings/
// Difficulty: Medium

import "fmt"

func numberOfSameEndSubstrings(s string, queries [][]int) []int {
	n := len(s)

	// prefix[c][i+1] = count of character c in s[0..i]
	prefix := make([][]int, 26)
	for c := 0; c < 26; c++ {
		prefix[c] = make([]int, n+1)
	}
	for i := 0; i < n; i++ {
		for c := 0; c < 26; c++ {
			prefix[c][i+1] = prefix[c][i]
		}
		prefix[s[i]-'a'][i+1]++
	}

	ans := make([]int, len(queries))
	for qi, q := range queries {
		l, r := q[0], q[1]
		total := 0
		for c := 0; c < 26; c++ {
			x := prefix[c][r+1] - prefix[c][l]
			if x > 0 {
				total += x * (x + 1) / 2
			}
		}
		ans[qi] = total
	}

	return ans
}

func main() {
	// Test case 1: s = "abca", queries = [[0,3]]
	// 'a' count=2: 2*3/2=3, 'b' count=1: 1, 'c' count=1: 1 -> total=5
	fmt.Println(numberOfSameEndSubstrings("abca", [][]int{{0, 3}})) // [5]

	// Test case 2: single character
	fmt.Println(numberOfSameEndSubstrings("a", [][]int{{0, 0}})) // [1]

	// Test case 3: all same characters
	fmt.Println(numberOfSameEndSubstrings("aaaa", [][]int{{0, 3}})) // [10]
	// 'a' count=4: 4*5/2=10

	// Test case 4: multiple queries
	// "abc": [0,0]->a=1->1, [0,1]->a=1,b=1->2, [0,2]->a=1,b=1,c=1->3
	fmt.Println(numberOfSameEndSubstrings("abc", [][]int{{0, 0}, {0, 1}, {0, 2}}))
	// [1, 2, 3]
	fmt.Println(numberOfSameEndSubstrings("aba", [][]int{{0, 2}})) // [4]
	// 'a' at 0 and 2: count=2 -> 2*3/2=3, 'b' at 1: count=1 -> 1. total=4
	// Substrings: "a"(0), "b"(1), "a"(2), "aba"(0..2) -> 4
}

// Time: O(n * 26 + q * 26) = O(n + q) | Space: O(26 * n)
```

## 2957 — Remove Adjacent Almost Equal Characters

```go
package main

// LeetCode #2957: Remove Adjacent Almost-Equal Characters
// https://leetcode.com/problems/remove-adjacent-almost-equal-characters/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(removeAlmostEqualCharacters("aaaaa"))
	fmt.Println(removeAlmostEqualCharacters("abdcd"))
	fmt.Println(removeAlmostEqualCharacters("acb"))
}

func removeAlmostEqualCharacters(word string) (ans int) {
	for i := 1; i < len(word); i++ {
		d := int(word[i]) - int(word[i-1])
		if d < 0 {
			d = -d
		}
		if d < 2 {
			ans++
			i++
		}
	}
	return
}
```

## 2958 — Length Of Longest Subarray With At Most K Frequency

```go
package main

// LeetCode #2958: Length of Longest Subarray With at Most K Frequency
// https://leetcode.com/problems/length-of-longest-subarray-with-at-most-k-frequency/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(maxSubarrayLength([]int{1, 2, 3, 1, 2, 3, 1, 2}, 2))
	fmt.Println(maxSubarrayLength([]int{1, 2, 1, 2, 1, 2, 1, 2}, 1))
	fmt.Println(maxSubarrayLength([]int{5, 5, 5, 5, 5, 5, 5}, 4))
}

func maxSubarrayLength(nums []int, k int) (ans int) {
	cnt := map[int]int{}
	for i, j, n := 0, 0, len(nums); i < n; i++ {
		cnt[nums[i]]++
		for ; cnt[nums[i]] > k; j++ {
			cnt[nums[j]]--
		}
		if i-j+1 > ans {
			ans = i - j + 1
		}
	}
	return
}
```

## 2961 — Double Modular Exponentiation

```go
package main

// LeetCode #2961: Double Modular Exponentiation
// https://leetcode.com/problems/double-modular-exponentiation/
// Difficulty: Medium
// Time: O(n log m) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(getGoodIndices([][]int{{2, 3, 3, 10}, {3, 3, 3, 1}, {6, 1, 1, 4}}, 2))
	fmt.Println(getGoodIndices([][]int{{39, 3, 1000, 1000}}, 17))
}

func getGoodIndices(variables [][]int, target int) (ans []int) {
	qpow := func(a, n, mod int) int {
		ans := 1
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				ans = ans * a % mod
			}
			a = a * a % mod
		}
		return ans
	}
	for i, e := range variables {
		a, b, c, m := e[0], e[1], e[2], e[3]
		if qpow(qpow(a, b, 10), c, m) == target {
			ans = append(ans, i)
		}
	}
	return
}
```

## 2962 — Count Subarrays Where Max Element Appears At Least K Times

```go
package main

// LeetCode #2962: Count Subarrays Where Max Element Appears at Least K Times
// https://leetcode.com/problems/count-subarrays-where-max-element-appears-at-least-k-times/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(countSubarrays2962([]int{1, 3, 2, 3, 3}, 2))
	fmt.Println(countSubarrays2962([]int{1, 4, 2, 1}, 3))
}

func countSubarrays2962(nums []int, k int) (ans int64) {
	mx := 0
	for _, x := range nums {
		if x > mx {
			mx = x
		}
	}
	n := len(nums)
	cnt, j := 0, 0
	for _, x := range nums {
		for ; j < n && cnt < k; j++ {
			if nums[j] == mx {
				cnt++
			}
		}
		if cnt < k {
			break
		}
		ans += int64(n - j + 1)
		if x == mx {
			cnt--
		}
	}
	return
}
```

## 2964 — Number Of Divisible Triplet Sums

```go
package main

// LeetCode #2964: Number of Divisible Triplet Sums
// https://leetcode.com/problems/number-of-divisible-triplet-sums/
// Difficulty: Medium

import "fmt"

func numberOfDivisibleTripletSums(nums []int, d int) int {
	n := len(nums)
	count := 0

	// pre[rem] = count of elements before current j with remainder rem
	pre := make(map[int]int)

	for j := 0; j < n; j++ {
		// For each k > j, find nums[i] (i < j) such that
		// (nums[i] + nums[j] + nums[k]) % d == 0
		for k := j + 1; k < n; k++ {
			needed := (d - (nums[j]+nums[k])%d) % d
			count += pre[needed]
		}
		// Add current element's remainder to pre for future j
		pre[nums[j]%d]++
	}

	return count
}

func main() {
	// Test case 1: nums = [3,3,4,7,8], d = 5 -> 3
	fmt.Println(numberOfDivisibleTripletSums([]int{3, 3, 4, 7, 8}, 5)) // 3

	// Test case 2: nums = [3,3,3,3], d = 3 -> 4
	fmt.Println(numberOfDivisibleTripletSums([]int{3, 3, 3, 3}, 3)) // 4

	// Test case 3: nums = [1,2,3], d = 3 -> 1 (1+2+3=6, 6%3==0)
	fmt.Println(numberOfDivisibleTripletSums([]int{1, 2, 3}, 3)) // 1

	// Test case 4: nums = [1,1,1], d = 1 -> 1 (all triplets sum to 3, 3%1==0)
	fmt.Println(numberOfDivisibleTripletSums([]int{1, 1, 1}, 1)) // 1

	// Test case 5: nums = [1,2,3,4], d = 2
	// Triplets: (0,1,2)=6%2=0, (0,1,3)=7%2=1, (0,2,3)=8%2=0, (1,2,3)=9%2=1 -> 2
	fmt.Println(numberOfDivisibleTripletSums([]int{1, 2, 3, 4}, 2)) // 2
}

// Time: O(n^2) | Space: O(d)
```

## 2966 — Divide Array Into Arrays With Max Difference

```go
package main

// LeetCode #2966: Divide Array Into Arrays With Max Difference
// https://leetcode.com/problems/divide-array-into-arrays-with-max-difference/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(divideArray2966([]int{1, 3, 4, 8, 7, 9, 3, 5, 1}, 2))
	fmt.Println(divideArray2966([]int{1, 3, 3, 2, 7, 3}, 3))
	fmt.Println(divideArray2966([]int{1, 2, 3}, 0))
}

func divideArray2966(nums []int, k int) [][]int {
	sort.Ints(nums)
	ans := [][]int{}
	for i := 0; i < len(nums); i += 3 {
		t := make([]int, 3)
		copy(t, nums[i:i+3])
		if t[2]-t[0] > k {
			return [][]int{}
		}
		ans = append(ans, t)
	}
	return ans
}
```

## 2967 — Minimum Cost To Make Array Equalindromic

```go
package main

// LeetCode #2967: Minimum Cost to Make Array Equalindromic
// https://leetcode.com/problems/minimum-cost-to-make-array-equalindromic/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minimumCostEqualindromic([]int{1, 2, 3, 4, 5}))
	fmt.Println(minimumCostEqualindromic([]int{10, 12, 13, 14, 15}))
}

func minimumCostEqualindromic(nums []int) int64 {
	sort.Ints(nums)
	median := nums[len(nums)/2]

	isPal := func(x int) bool {
		s := fmt.Sprintf("%d", x)
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			if s[i] != s[j] {
				return false
			}
		}
		return true
	}

	nextPal := median
	for !isPal(nextPal) {
		nextPal++
	}
	prevPal := median
	for !isPal(prevPal) {
		prevPal--
	}

	cost1, cost2 := int64(0), int64(0)
	for _, num := range nums {
		cost1 += int64(abs(num - prevPal))
		cost2 += int64(abs(num - nextPal))
	}
	if cost1 < cost2 {
		return cost1
	}
	return cost2
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

## 2971 — Find Polygon With The Largest Perimeter

```go
package main

// LeetCode #2971: Find Polygon With the Largest Perimeter
// https://leetcode.com/problems/find-polygon-with-the-largest-perimeter/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(largestPerimeter([]int{2, 1, 2}))
	fmt.Println(largestPerimeter([]int{1, 2, 1, 10}))
	fmt.Println(largestPerimeter([]int{1, 2, 3}))
}

func largestPerimeter(nums []int) int64 {
	sort.Ints(nums)
	s := int64(0)
	for _, x := range nums {
		s += int64(x)
	}
	for i := len(nums) - 1; i >= 2; i-- {
		if int64(nums[i]) < s-int64(nums[i]) {
			return s
		}
		s -= int64(nums[i])
	}
	return -1
}
```

## 2975 — Maximum Square Area By Removing Fences From A Field

```go
package main

// LeetCode #2975: Maximum Square Area by Removing Fences From a Field
// https://leetcode.com/problems/maximum-square-area-by-removing-fences-from-a-field/
// Difficulty: Medium
// Time: O(h^2 + v^2) | Space: O(h^2 + v^2)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximizeSquareAreaField(4, 3, []int{2, 3}, []int{2}))
	fmt.Println(maximizeSquareAreaField(6, 4, []int{3, 5}, []int{2, 3}))
}

func maximizeSquareAreaField(m int, n int, hFences []int, vFences []int) int {
	const mod = 1_000_000_007

	getGaps := func(fences []int, limit int) map[int]bool {
		arr := append(fences, 1, limit)
		sort.Ints(arr)
		gaps := map[int]bool{}
		for i := 0; i < len(arr); i++ {
			for j := i + 1; j < len(arr); j++ {
				gaps[arr[j]-arr[i]] = true
			}
		}
		return gaps
	}

	hGaps := getGaps(hFences, m)
	vGaps := getGaps(vFences, n)

	maxLen := 0
	for d := range hGaps {
		if vGaps[d] && d > maxLen {
			maxLen = d
		}
	}
	if maxLen == 0 {
		return -1
	}
	return (maxLen * maxLen) % mod
}
```

## 2976 — Minimum Cost To Convert String I

```go
package main

// LeetCode #2976: Minimum Cost to Convert String I
// https://leetcode.com/problems/minimum-cost-to-convert-string-i/
// Difficulty: Medium
// Time: O(26^3 + n) | Space: O(26^2)

import "fmt"

func main() {
	fmt.Println(minimumCostConvert("abcd", "acbe", []byte("ab"), []byte("ce"), []int{5, 3}))
	fmt.Println(minimumCostConvert("aaaa", "bbbb", []byte("a"), []byte("b"), []int{2}))
}

func minimumCostConvert(source string, target string, original []byte, changed []byte, cost []int) (ans int64) {
	const inf = 1 << 29
	g := make([][]int, 26)
	for i := range g {
		g[i] = make([]int, 26)
		for j := range g[i] {
			if i == j {
				g[i][j] = 0
			} else {
				g[i][j] = inf
			}
		}
	}
	for i := 0; i < len(original); i++ {
		x := int(original[i] - 'a')
		y := int(changed[i] - 'a')
		z := cost[i]
		if z < g[x][y] {
			g[x][y] = z
		}
	}
	for k := 0; k < 26; k++ {
		for i := 0; i < 26; i++ {
			for j := 0; j < 26; j++ {
				if g[i][k]+g[k][j] < g[i][j] {
					g[i][j] = g[i][k] + g[k][j]
				}
			}
		}
	}
	for i := 0; i < len(source); i++ {
		x := int(source[i] - 'a')
		y := int(target[i] - 'a')
		if x != y {
			if g[x][y] >= inf {
				return -1
			}
			ans += int64(g[x][y])
		}
	}
	return
}
```

## 2978 — Symmetric Coordinates

```go
package main

// LeetCode #2978: Symmetric Coordinates
// https://leetcode.com/problems/symmetric-coordinates/
// Difficulty: Medium (SQL problem — implemented in Go)
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"sort"
)

type Coordinate struct {
	X int
	Y int
}

func symmetricCoordinates(coords []Coordinate) []Coordinate {
	// Build set of all coordinates
	coordSet := make(map[[2]int]bool)
	for _, c := range coords {
		coordSet[[2]int{c.X, c.Y}] = true
	}

	seen := make(map[[2]int]bool)
	var results []Coordinate

	for _, c := range coords {
		x, y := c.X, c.Y
		if x > y {
			continue
		}

		if seen[[2]int{x, y}] {
			continue
		}
		seen[[2]int{x, y}] = true

		// Check if symmetric pair (y,x) also exists
		if coordSet[[2]int{y, x}] {
			results = append(results, Coordinate{X: x, Y: y})
		}
	}

	// Order by X ASC, Y ASC
	sort.Slice(results, func(i, j int) bool {
		if results[i].X != results[j].X {
			return results[i].X < results[j].X
		}
		return results[i].Y < results[j].Y
	})

	return results
}

func main() {
	coords := []Coordinate{
		{X: 1, Y: 2},
		{X: 3, Y: 4},
		{X: 2, Y: 1},
		{X: 5, Y: 6},
		{X: 4, Y: 3},
		{X: 7, Y: 7},
		{X: 8, Y: 9},
		{X: 1, Y: 3},
	}

	fmt.Println("Symmetric Coordinates")
	fmt.Println("====================")
	fmt.Printf("%-6s %s\n", "X", "Y")
	fmt.Println("-----------")

	results := symmetricCoordinates(coords)
	for _, r := range results {
		fmt.Printf("%-6d %d\n", r.X, r.Y)
	}
}
```

## 2979 — Most Expensive Item That Can Not Be Bought

```go
package main

// LeetCode #2979: Most Expensive Item That Can Not Be Bought
// https://leetcode.com/problems/most-expensive-item-that-can-not-be-bought/
// Difficulty: Medium [Paid]
// Time: O(1) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(mostExpensiveItem(2, 5))
	fmt.Println(mostExpensiveItem(3, 7))
}

func mostExpensiveItem(primeOne int, primeTwo int) int {
	return primeOne*primeTwo - primeOne - primeTwo
}
```

## 2981 — Find Longest Special Substring That Occurs Thrice I

```go
package main

// LeetCode #2981: Find Longest Special Substring That Occurs Thrice I
// https://leetcode.com/problems/find-longest-special-substring-that-occurs-thrice-i/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(maximumLength2981("aaaa"))
	fmt.Println(maximumLength2981("abcdef"))
	fmt.Println(maximumLength2981("abcaba"))
}

func maximumLength2981(s string) int {
	n := len(s)
	l, r := 0, n
	check := func(x int) bool {
		cnt := [26]int{}
		for i := 0; i < n; {
			j := i + 1
			for j < n && s[j] == s[i] {
				j++
			}
			k := s[i] - 'a'
			add := j - i - x + 1
			if add > 0 {
				cnt[k] += add
			}
			if cnt[k] >= 3 {
				return true
			}
			i = j
		}
		return false
	}
	for l < r {
		mid := (l + r + 1) >> 1
		if check(mid) {
			l = mid
		} else {
			r = mid - 1
		}
	}
	if l == 0 {
		return -1
	}
	return l
}
```

## 2982 — Find Longest Special Substring That Occurs Thrice Ii

```go
package main

// LeetCode #2982: Find Longest Special Substring That Occurs Thrice II
// https://leetcode.com/problems/find-longest-special-substring-that-occurs-thrice-ii/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(maximumLength2982("aaaa"))
	fmt.Println(maximumLength2982("abcdef"))
	fmt.Println(maximumLength2982("abcaba"))
}

func maximumLength2982(s string) int {
	n := len(s)
	l, r := 0, n
	check := func(x int) bool {
		cnt := [26]int{}
		for i := 0; i < n; {
			j := i + 1
			for j < n && s[j] == s[i] {
				j++
			}
			k := s[i] - 'a'
			add := j - i - x + 1
			if add > 0 {
				cnt[k] += add
			}
			if cnt[k] >= 3 {
				return true
			}
			i = j
		}
		return false
	}
	for l < r {
		mid := (l + r + 1) >> 1
		if check(mid) {
			l = mid
		} else {
			r = mid - 1
		}
	}
	if l == 0 {
		return -1
	}
	return l
}
```

## 2984 — Find Peak Calling Hours For Each City

```go
package main

// LeetCode #2984: Find Peak Calling Hours for Each City
// https://leetcode.com/problems/find-peak-calling-hours-for-each-city/
// Difficulty: Medium (SQL problem — implemented in Go)
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"sort"
)

type Call struct {
	CallerID    int
	RecipientID int
	CallTime    string // datetime string, e.g. "2024-01-01 14:30:00"
	City        string
}

type PeakHour struct {
	City      string
	PeakHour  int
	CallCount int
}

func extractHour(datetime string) int {
	// Format: "2024-01-01 14:30:00"
	// Extract the hour (two digits after the first space and before the first colon after space)
	var year, month, day, hour, min, sec int
	_, err := fmt.Sscanf(datetime, "%d-%d-%d %d:%d:%d", &year, &month, &day, &hour, &min, &sec)
	if err != nil {
		return 0
	}
	return hour
}

func findPeakCallingHours(calls []Call) []PeakHour {
	// Count calls per city per hour: map[city]map[hour]count
	cityHourCounts := make(map[string]map[int]int)
	for _, c := range calls {
		hour := extractHour(c.CallTime)
		if cityHourCounts[c.City] == nil {
			cityHourCounts[c.City] = make(map[int]int)
		}
		cityHourCounts[c.City][hour]++
	}

	var results []PeakHour

	for city, hourCounts := range cityHourCounts {
		// Find max count for this city
		maxCount := 0
		for _, count := range hourCounts {
			if count > maxCount {
				maxCount = count
			}
		}

		// Collect all hours with max count
		for hour, count := range hourCounts {
			if count == maxCount {
				results = append(results, PeakHour{
					City:      city,
					PeakHour:  hour,
					CallCount: count,
				})
			}
		}
	}

	// Order by peak_hour DESC, city DESC
	sort.Slice(results, func(i, j int) bool {
		if results[i].PeakHour != results[j].PeakHour {
			return results[i].PeakHour > results[j].PeakHour // DESC
		}
		return results[i].City > results[j].City // DESC
	})

	return results
}

func main() {
	calls := []Call{
		{CallerID: 1, RecipientID: 2, CallTime: "2024-01-01 09:15:00", City: "New York"},
		{CallerID: 3, RecipientID: 4, CallTime: "2024-01-01 09:30:00", City: "New York"},
		{CallerID: 5, RecipientID: 6, CallTime: "2024-01-01 14:00:00", City: "New York"},
		{CallerID: 7, RecipientID: 8, CallTime: "2024-01-01 09:45:00", City: "New York"},
		{CallerID: 9, RecipientID: 10, CallTime: "2024-01-01 10:00:00", City: "Los Angeles"},
		{CallerID: 11, RecipientID: 12, CallTime: "2024-01-01 10:15:00", City: "Los Angeles"},
		{CallerID: 13, RecipientID: 14, CallTime: "2024-01-01 14:00:00", City: "Los Angeles"},
		{CallerID: 15, RecipientID: 16, CallTime: "2024-01-01 10:30:00", City: "Los Angeles"},
		{CallerID: 17, RecipientID: 18, CallTime: "2024-01-01 10:00:00", City: "Chicago"},
		{CallerID: 19, RecipientID: 20, CallTime: "2024-01-01 11:00:00", City: "Chicago"},
		{CallerID: 21, RecipientID: 22, CallTime: "2024-01-01 10:00:00", City: "Chicago"},
	}

	fmt.Println("Find Peak Calling Hours for Each City")
	fmt.Println("====================================")
	fmt.Printf("%-15s %-12s %s\n", "City", "Peak Hour", "Call Count")
	fmt.Println("--------------------------------------")

	results := findPeakCallingHours(calls)
	for _, r := range results {
		fmt.Printf("%-15s %-12d %d\n", r.City, r.PeakHour, r.CallCount)
	}
}
```

## 2986 — Find Third Transaction

```go
package main

// LeetCode #2986: Find Third Transaction
// https://leetcode.com/problems/find-third-transaction/
// Difficulty: Medium (SQL problem — simulated in Go)
//
// Simulates: For each user with >=3 transactions, find the 3rd transaction
// (by date ASC). Include only if the 3rd transaction's spend > BOTH the
// 1st and 2nd transaction's spend.

import (
	"fmt"
	"sort"
)

// Transaction represents the Transactions database table.
type Transaction struct {
	UserID          int
	Spend           float64
	TransactionDate string
}

// ThirdTransactionResult holds the output for each qualifying user.
type ThirdTransactionResult struct {
	UserID                int
	ThirdTransactionSpend float64
	ThirdTransactionDate  string
}

// findThirdTransaction simulates the SQL query.
// Time: O(n log n) | Space: O(n)
// n = number of transactions per user, overall O(n log n) due to sorting.
func findThirdTransaction(transactions []Transaction) []ThirdTransactionResult {
	// Group transactions by user_id.
	userTxns := make(map[int][]Transaction)
	for _, t := range transactions {
		userTxns[t.UserID] = append(userTxns[t.UserID], t)
	}

	var results []ThirdTransactionResult

	for _, txns := range userTxns {
		if len(txns) < 3 {
			continue
		}
		// Sort by transaction_date ASC.
		sort.Slice(txns, func(i, j int) bool {
			return txns[i].TransactionDate < txns[j].TransactionDate
		})

		first := txns[0].Spend
		second := txns[1].Spend
		third := txns[2].Spend

		// Check if 3rd transaction's spend > both 1st and 2nd.
		if third > first && third > second {
			results = append(results, ThirdTransactionResult{
				UserID:                txns[0].UserID,
				ThirdTransactionSpend: third,
				ThirdTransactionDate:  txns[2].TransactionDate,
			})
		}
	}

	// Order by user_id ASC.
	sort.Slice(results, func(i, j int) bool {
		return results[i].UserID < results[j].UserID
	})

	return results
}

func main() {
	// Test data from the problem.
	transactions := []Transaction{
		{UserID: 1, Spend: 7.44, TransactionDate: "2022-07-11"},
		{UserID: 1, Spend: 49.78, TransactionDate: "2022-07-12"},
		{UserID: 1, Spend: 65.56, TransactionDate: "2022-07-13"},
		{UserID: 1, Spend: 30.00, TransactionDate: "2022-07-14"},
		// User 2 has 3 transactions, but 3rd (25.00) is not > 2nd (30.00).
		{UserID: 2, Spend: 10.00, TransactionDate: "2022-07-11"},
		{UserID: 2, Spend: 30.00, TransactionDate: "2022-07-12"},
		{UserID: 2, Spend: 25.00, TransactionDate: "2022-07-13"},
		// User 3 has only 2 transactions, does not qualify.
		{UserID: 3, Spend: 5.00, TransactionDate: "2022-07-11"},
		{UserID: 3, Spend: 8.00, TransactionDate: "2022-07-12"},
	}

	results := findThirdTransaction(transactions)

	fmt.Println("Third Transaction Results (user_id | third_transaction_spend | third_transaction_date):")
	for _, r := range results {
		fmt.Printf("%d | %.2f | %s\n", r.UserID, r.ThirdTransactionSpend, r.ThirdTransactionDate)
	}
	// Expected output:
	// 1 | 65.56 | 2022-07-13
}
```

## 2988 — Manager Of The Largest Department

```go
package main

// LeetCode #2988: Manager of the Largest Department
// https://leetcode.com/problems/manager-of-the-largest-department/
// Difficulty: Medium (SQL problem — simulated in Go)
//
// Simulates: Find managers (position = 'Manager') of the department(s)
// with the most employees. If multiple departments tie, include all
// managers from those departments.

import (
	"fmt"
	"sort"
)

// Employee represents the Employees database table.
type Employee struct {
	EmpID    int
	EmpName  string
	DepID    int
	Position string
}

// ManagerResult holds the output.
type ManagerResult struct {
	EmpName string
	DepID   int
}

// findManagersOfLargestDepartment simulates the SQL query.
// Time: O(n) | Space: O(n)
// n = number of employees.
func findManagersOfLargestDepartment(employees []Employee) []ManagerResult {
	// Count employees per department.
	depCount := make(map[int]int)
	for _, e := range employees {
		depCount[e.DepID]++
	}

	// Find the maximum employee count.
	maxCount := 0
	for _, count := range depCount {
		if count > maxCount {
			maxCount = count
		}
	}

	// Collect managers from departments with maxCount employees.
	var results []ManagerResult
	for _, e := range employees {
		if e.Position == "Manager" && depCount[e.DepID] == maxCount {
			results = append(results, ManagerResult{EmpName: e.EmpName, DepID: e.DepID})
		}
	}

	// Order by dep_id ASC.
	sort.Slice(results, func(i, j int) bool {
		return results[i].DepID < results[j].DepID
	})

	return results
}

func main() {
	// Test data.
	employees := []Employee{
		// Department 1 has 3 employees, managers: Alice.
		{EmpID: 1, EmpName: "Alice", DepID: 1, Position: "Manager"},
		{EmpID: 2, EmpName: "Bob", DepID: 1, Position: "Employee"},
		{EmpID: 3, EmpName: "Charlie", DepID: 1, Position: "Employee"},
		// Department 2 has 2 employees, manager: Diana.
		{EmpID: 4, EmpName: "Diana", DepID: 2, Position: "Manager"},
		{EmpID: 5, EmpName: "Eve", DepID: 2, Position: "Employee"},
		// Department 3 has 3 employees (tie with dep 1), managers: Frank.
		{EmpID: 6, EmpName: "Frank", DepID: 3, Position: "Manager"},
		{EmpID: 7, EmpName: "Grace", DepID: 3, Position: "Employee"},
		{EmpID: 8, EmpName: "Henry", DepID: 3, Position: "Employee"},
	}

	results := findManagersOfLargestDepartment(employees)

	fmt.Println("Manager(s) of Largest Department(s) (emp_name | dep_id):")
	for _, r := range results {
		fmt.Printf("%s | %d\n", r.EmpName, r.DepID)
	}
	// Expected output (departments 1 and 3 tie with 3 employees each):
	// Alice | 1
	// Frank | 3
}
```

## 2989 — Class Performance

```go
package main

// LeetCode #2989: Class Performance
// https://leetcode.com/problems/class-performance/
// Difficulty: Medium (SQL problem — simulated in Go)
//
// Simulates: Find the difference between the highest total score
// (sum of 3 assignments) and the lowest total score across all students.

import (
	"fmt"
)

// Score represents the Scores database table.
type Score struct {
	StudentID   int
	StudentName string
	Assignment1 int
	Assignment2 int
	Assignment3 int
}

// classPerformance computes max(total) - min(total) across all students.
// Time: O(n) | Space: O(1)
// n = number of students.
func classPerformance(scores []Score) int {
	if len(scores) == 0 {
		return 0
	}

	minTotal := 0
	maxTotal := 0

	for i, s := range scores {
		total := s.Assignment1 + s.Assignment2 + s.Assignment3
		if i == 0 {
			minTotal = total
			maxTotal = total
		} else {
			if total < minTotal {
				minTotal = total
			}
			if total > maxTotal {
				maxTotal = total
			}
		}
	}

	return maxTotal - minTotal
}

func main() {
	// Test data from the problem.
	scores := []Score{
		{StudentID: 309, StudentName: "Owen", Assignment1: 88, Assignment2: 47, Assignment3: 87},
		{StudentID: 321, StudentName: "Claire", Assignment1: 98, Assignment2: 95, Assignment3: 37},
		{StudentID: 338, StudentName: "Julian", Assignment1: 100, Assignment2: 64, Assignment3: 43},
		{StudentID: 423, StudentName: "Peyton", Assignment1: 60, Assignment2: 44, Assignment3: 47},
		{StudentID: 896, StudentName: "David", Assignment1: 32, Assignment2: 37, Assignment3: 50},
		{StudentID: 235, StudentName: "Camila", Assignment1: 31, Assignment2: 53, Assignment3: 69},
	}

	result := classPerformance(scores)

	fmt.Printf("Class Performance — difference in score: %d\n", result)
	// Expected: 111 (Claire total=230, David total=119)
}
```

## 2992 — Number Of Self Divisible Permutations

```go
package main

// LeetCode #2992: Number of Self-Divisible Permutations
// https://leetcode.com/problems/number-of-self-divisible-permutations/
// Difficulty: Medium

import "fmt"

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func numberOfSelfDivisiblePermutations(n int) int {
	maskCount := 1 << n
	dp := make([]int, maskCount)
	dp[0] = 1

	// position (1-indexed) = number of set bits in the mask
	// Precompute popcount for each mask
	popcount := make([]int, maskCount)
	for mask := 1; mask < maskCount; mask++ {
		popcount[mask] = popcount[mask>>1] + (mask & 1)
	}

	for mask := 1; mask < maskCount; mask++ {
		pos := popcount[mask] // 1-indexed position
		for j := 0; j < n; j++ {
			if mask&(1<<j) != 0 && gcd(pos, j+1) == 1 {
				dp[mask] += dp[mask^(1<<j)]
			}
		}
	}

	return dp[maskCount-1]
}

func main() {
	// Test case 1: n = 1 -> 1 (only [1])
	fmt.Println(numberOfSelfDivisiblePermutations(1)) // 1

	// Test case 2: n = 2 -> 1 (only [2,1])
	fmt.Println(numberOfSelfDivisiblePermutations(2)) // 1

	// Test case 3: n = 3 -> 3 ([1,3,2], [3,1,2], [2,3,1])
	fmt.Println(numberOfSelfDivisiblePermutations(3)) // 3

	// Test case 4: n = 4
	fmt.Println(numberOfSelfDivisiblePermutations(4)) // ?

	// Test case 5: n = 5
	fmt.Println(numberOfSelfDivisiblePermutations(5)) // ?
}

// Time: O(n * 2^n) | Space: O(2^n)
```

## 2993 — Friday Purchases I

```go
package main

// LeetCode #2993: Friday Purchases I
// https://leetcode.com/problems/friday-purchases-i/
// Difficulty: Medium (SQL problem — simulated in Go)
//
// Simulates: For Fridays in November 2023, calculate total spending per
// week of the month. Week number = CEIL(day_of_month / 7). Only include
// weeks with at least one purchase.

import (
	"fmt"
	"sort"
	"time"
)

// Purchase represents the Purchases database table.
type Purchase struct {
	UserID       int
	PurchaseDate string // format: "YYYY-MM-DD"
	AmountSpend  int
}

// WeeklyResult holds the output.
type WeeklyResult struct {
	WeekOfMonth int
	TotalAmount int
}

// findFridayPurchasesI simulates the SQL query.
// Time: O(n) | Space: O(d) where d = distinct weeks
// n = number of purchases.
func findFridayPurchasesI(purchases []Purchase) []WeeklyResult {
	weeklyAmount := make(map[int]int)

	for _, p := range purchases {
		// Parse the date.
		t, err := time.Parse("2006-01-02", p.PurchaseDate)
		if err != nil {
			continue
		}

		// Filter: November 2023 only.
		if t.Year() != 2023 || t.Month() != time.November {
			continue
		}

		// Filter: Friday only (weekday = 5 = time.Friday).
		if t.Weekday() != time.Friday {
			continue
		}

		// Compute week_of_month: CEIL(day / 7).
		day := t.Day()
		weekOfMonth := (day-1)/7 + 1

		weeklyAmount[weekOfMonth] += p.AmountSpend
	}

	var results []WeeklyResult
	for w, amt := range weeklyAmount {
		results = append(results, WeeklyResult{WeekOfMonth: w, TotalAmount: amt})
	}

	// Order by week_of_month ASC.
	sort.Slice(results, func(i, j int) bool {
		return results[i].WeekOfMonth < results[j].WeekOfMonth
	})

	return results
}

func main() {
	// Test data: various purchases in November 2023.
	purchases := []Purchase{
		// Friday Nov 3 (week 1)
		{UserID: 1, PurchaseDate: "2023-11-03", AmountSpend: 50},
		{UserID: 2, PurchaseDate: "2023-11-03", AmountSpend: 30},
		// Non-Friday should be excluded
		{UserID: 3, PurchaseDate: "2023-11-04", AmountSpend: 20},
		// Friday Nov 10 (week 2)
		{UserID: 4, PurchaseDate: "2023-11-10", AmountSpend: 100},
		// Friday Nov 17 (week 3) — no purchases, should not appear
		// Friday Nov 24 (week 4)
		{UserID: 5, PurchaseDate: "2023-11-24", AmountSpend: 200},
		{UserID: 6, PurchaseDate: "2023-11-24", AmountSpend: 75},
		// Outside November
		{UserID: 7, PurchaseDate: "2023-10-06", AmountSpend: 40},
		{UserID: 8, PurchaseDate: "2023-12-01", AmountSpend: 60},
	}

	results := findFridayPurchasesI(purchases)

	fmt.Println("Friday Purchases I (week_of_month | total_amount):")
	for _, r := range results {
		fmt.Printf("%d | %d\n", r.WeekOfMonth, r.TotalAmount)
	}
	// Expected output:
	// 1 | 80
	// 2 | 100
	// 4 | 275
}
```

## 2997 — Minimum Number Of Operations To Make Array Xor Equal To K

```go
package main

// LeetCode #2997: Minimum Number of Operations to Make Array XOR Equal to K
// https://leetcode.com/problems/minimum-number-of-operations-to-make-array-xor-equal-to-k/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(minOperationsXOR([]int{2, 1, 3, 4}, 1))
	fmt.Println(minOperationsXOR([]int{2, 0, 2, 0}, 0))
}

func minOperationsXOR(nums []int, k int) (ans int) {
	xor := 0
	for _, x := range nums {
		xor ^= x
	}
	return bits.OnesCount(uint(xor ^ k))
}
```

## 2998 — Minimum Number Of Operations To Make X And Y Equal

```go
package main

// LeetCode #2998: Minimum Number of Operations to Make X and Y Equal
// https://leetcode.com/problems/minimum-number-of-operations-to-make-x-and-y-equal/
// Difficulty: Medium
// Time: O(log n) | Space: O(log n)

import "fmt"

func main() {
	fmt.Println(minimumOperationsToMakeEqual(26, 1))
	fmt.Println(minimumOperationsToMakeEqual(3, 10))
	fmt.Println(minimumOperationsToMakeEqual(54, 2))
}

func minimumOperationsToMakeEqual(x int, y int) int {
	f := map[int]int{}
	var dfs func(int) int
	dfs = func(x int) int {
		if y >= x {
			return y - x
		}
		if v, ok := f[x]; ok {
			return v
		}
		a := x%5 + 1 + dfs(x/5)
		b := 5 - x%5 + 1 + dfs(x/5+1)
		c := x%11 + 1 + dfs(x/11)
		d := 11 - x%11 + 1 + dfs(x/11+1)
		res := x - y
		if a < res {
			res = a
		}
		if b < res {
			res = b
		}
		if c < res {
			res = c
		}
		if d < res {
			res = d
		}
		f[x] = res
		return res
	}
	return dfs(x)
}
```

## 3001 — Minimum Moves To Capture The Queen

```go
package main

// LeetCode #3001: Minimum Moves to Capture The Queen
// https://leetcode.com/problems/minimum-moves-to-capture-the-queen/
// Difficulty: Medium
// Time: O(1) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(minMovesToCaptureTheQueen(1, 1, 8, 8, 2, 3))
	fmt.Println(minMovesToCaptureTheQueen(5, 3, 3, 4, 5, 2))
}

func minMovesToCaptureTheQueen(a, b, c, d, e, f int) int {
	// Rook and queen on same row, bishop not in between
	if a == e && !(a == c && (d-b)*(d-f) < 0) {
		return 1
	}
	// Rook and queen on same column, bishop not in between
	if b == f && !(b == d && (c-a)*(c-e) < 0) {
		return 1
	}
	// Bishop and queen on same diagonal, rook not in between
	if c+d == e+f && !(a+b == e+f && (a-c)*(a-e) < 0) {
		return 1
	}
	if c-d == e-f && !(a-b == e-f && (a-c)*(a-e) < 0) {
		return 1
	}
	return 2
}
```

## 3002 — Maximum Size Of A Set After Removals

```go
package main

// LeetCode #3002: Maximum Size of a Set After Removals
// https://leetcode.com/problems/maximum-size-of-a-set-after-removals/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(maximumSetSize([]int{1, 2, 1, 2}, []int{1, 2, 1, 2}))
	fmt.Println(maximumSetSize([]int{1, 2, 3, 4}, []int{4, 3, 2, 1}))
	fmt.Println(maximumSetSize([]int{1, 2, 1, 2}, []int{3, 4, 5, 6}))
}

func maximumSetSize(nums1 []int, nums2 []int) int {
	s1 := map[int]bool{}
	s2 := map[int]bool{}
	for _, x := range nums1 {
		s1[x] = true
	}
	for _, x := range nums2 {
		s2[x] = true
	}
	a, b, c := 0, 0, 0
	for x := range s1 {
		if !s2[x] {
			a++
		}
	}
	for x := range s2 {
		if !s1[x] {
			b++
		} else {
			c++
		}
	}
	n := len(nums1)
	if a > n/2 {
		a = n / 2
	}
	if b > n/2 {
		b = n / 2
	}
	res := a + b + c
	if res > n {
		res = n
	}
	return res
}
```

## 3004 — Maximum Subtree Of The Same Color

```go
package main

// LeetCode #3004: Maximum Subtree of the Same Color (PAID)
// https://leetcode.com/problems/maximum-subtree-of-the-same-color/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

// Given an undirected tree rooted at 0, find the size of the largest subtree
// where all nodes have the same color. A subtree is rooted at some node v
// and includes all of v's descendants.

import "fmt"

func main() {
	// Test 1: Entire tree is same color
	edges := [][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {2, 5}, {2, 6}}
	colors := []int{1, 1, 1, 1, 1, 1, 1}
	fmt.Println(maximumSubtreeOfSameColor(edges, colors)) // 7

	// Test 2: One subtree branch differs
	colors2 := []int{1, 2, 1, 1, 1, 1, 1}
	fmt.Println(maximumSubtreeOfSameColor(edges, colors2)) // 3 (subtree [2,5,6] all color 1)

	// Test 3: All nodes different
	colors3 := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(maximumSubtreeOfSameColor(edges, colors3)) // 1 (each node alone)
}

func maximumSubtreeOfSameColor(edges [][]int, colors []int) int {
	n := len(colors)
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	ans := 0
	var dfs func(u, p int) (size int, same bool)
	dfs = func(u, p int) (size int, same bool) {
		size = 1
		same = true
		for _, v := range adj[u] {
			if v == p {
				continue
			}
			childSize, childSame := dfs(v, u)
			if childSame && colors[v] == colors[u] {
				size += childSize
			} else {
				same = false
			}
		}
		if same && size > ans {
			ans = size
		}
		return
	}

	dfs(0, -1)
	return ans
}
```

## 3006 — Find Beautiful Indices In The Given Array I

```go
package main

// LeetCode #3006: Find Beautiful Indices in the Given Array I
// https://leetcode.com/problems/find-beautiful-indices-in-the-given-array-i/
// Difficulty: Medium
// Time: O(n+m) | Space: O(n+m)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(beautifulIndices("abcd", "a", "da", 1))
	fmt.Println(beautifulIndices("isawsquirrelnearmysquirrelhouseohmy", "my", "squirrel", 15))
	fmt.Println(beautifulIndices("bcd", "a", "a", 1))
}

func beautifulIndices(s string, a string, b string, k int) (ans []int) {
	posA := kmpSearch(s, a)
	posB := kmpSearch(s, b)
	for _, i := range posA {
		idx := sort.SearchInts(posB, i)
		if idx < len(posB) && posB[idx]-i <= k {
			ans = append(ans, i)
		} else if idx > 0 && i-posB[idx-1] <= k {
			ans = append(ans, i)
		}
	}
	return
}

func kmpSearch(text, pattern string) (pos []int) {
	if pattern == "" {
		return
	}
	n, m := len(text), len(pattern)
	pi := make([]int, m)
	for i := 1; i < m; i++ {
		j := pi[i-1]
		for j > 0 && pattern[i] != pattern[j] {
			j = pi[j-1]
		}
		if pattern[i] == pattern[j] {
			j++
		}
		pi[i] = j
	}
	j := 0
	for i := 0; i < n; i++ {
		for j > 0 && text[i] != pattern[j] {
			j = pi[j-1]
		}
		if text[i] == pattern[j] {
			j++
		}
		if j == m {
			pos = append(pos, i-m+1)
			j = pi[j-1]
		}
	}
	return
}
```

## 3007 — Maximum Number That Sum Of The Prices Is Less Than Or Equal To K

```go
package main

// LeetCode #3007: Maximum Number That Sum of the Prices Is Less Than or Equal to K
// https://leetcode.com/problems/maximum-number-that-sum-of-the-prices-is-less-than-or-equal-to-k/
// Difficulty: Medium
// Time: O(log^2 n) | Space: O(log n)

import "fmt"

func main() {
	fmt.Println(findMaximumNumber(9, 1))
	fmt.Println(findMaximumNumber(7, 2))
}

func findMaximumNumber(k int64, x int) int64 {
	var lo, hi int64 = 0, 1e18
	for lo < hi {
		mid := (lo + hi + 1) / 2
		if countPrice(mid, x) <= k {
			lo = mid
		} else {
			hi = mid - 1
		}
	}
	return lo
}

func countPrice(num int64, x int) int64 {
	var ans int64
	n := num
	for bit := x - 1; int64(1<<bit) <= n; bit += x {
		cycle := int64(1 << (bit + 1))
		full := n / cycle
		ans += full * int64(1<<bit)
		rem := n % cycle
		ans += max(0, rem-int64(1<<bit)+1)
	}
	return ans
}
```

## 3011 — Find If Array Can Be Sorted

```go
package main

// LeetCode #3011: Find if Array Can Be Sorted
// https://leetcode.com/problems/find-if-array-can-be-sorted/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(canSortArray([]int{8, 4, 2, 30, 15}))
	fmt.Println(canSortArray([]int{1, 2, 3, 4, 5}))
	fmt.Println(canSortArray([]int{3, 16, 8, 4, 2}))
}

func canSortArray(nums []int) bool {
	prevMax, curMax, curMin := 0, 0, 0
	prevBits := -1
	for _, x := range nums {
		b := bitsCount(x)
		if b != prevBits {
			if prevBits != -1 {
				prevMax = curMax
			}
			curMax, curMin = x, x
			prevBits = b
		} else {
			if x < curMin {
				curMin = x
			}
			if x > curMax {
				curMax = x
			}
		}
		if curMin < prevMax {
			return false
		}
	}
	return true
}

func bitsCount(x int) int {
	cnt := 0
	for x > 0 {
		cnt += x & 1
		x >>= 1
	}
	return cnt
}
```

## 3012 — Minimize Length Of Array Using Operations

```go
package main

// LeetCode #3012: Minimize Length of Array Using Operations
// https://leetcode.com/problems/minimize-length-of-array-using-operations/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(minimumArrayLength([]int{1, 4, 3, 1}))
	fmt.Println(minimumArrayLength([]int{5, 5, 5, 10, 5}))
	fmt.Println(minimumArrayLength([]int{3, 5}))
}

func minimumArrayLength(nums []int) int {
	minVal := nums[0]
	for _, x := range nums[1:] {
		if x < minVal {
			minVal = x
		}
	}
	cnt := 0
	for _, x := range nums {
		if x%minVal != 0 {
			return 1
		}
		if x == minVal {
			cnt++
		}
	}
	return (cnt + 1) / 2
}
```

## 3015 — Count The Number Of Houses At A Certain Distance I

```go
package main

// LeetCode #3015: Count the Number of Houses at a Certain Distance I
// https://leetcode.com/problems/count-the-number-of-houses-at-a-certain-distance-i/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(countOfPairs(3, 1, 3))
	fmt.Println(countOfPairs(5, 2, 4))
	fmt.Println(countOfPairs(4, 1, 1))
}

func countOfPairs(n int, x int, y int) []int {
	ans := make([]int, n)
	if x > y {
		x, y = y, x
	}
	for i := 1; i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			dist := j - i
			via := abs(i-x) + 1 + abs(y-j)
			if via < dist {
				dist = via
			}
			ans[dist-1] += 2
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

## 3016 — Minimum Number Of Pushes To Type Word Ii

```go
package main

// LeetCode #3016: Minimum Number of Pushes to Type Word II
// https://leetcode.com/problems/minimum-number-of-pushes-to-type-word-ii/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minimumPushes("abcde"))
	fmt.Println(minimumPushes("xyzxyzxyzxyz"))
	fmt.Println(minimumPushes("aabbccddeeffgghhiiiiii"))
}

func minimumPushes(word string) int {
	cnt := make([]int, 26)
	for _, ch := range word {
		cnt[ch-'a']++
	}
	sort.Slice(cnt, func(i, j int) bool {
		return cnt[i] > cnt[j]
	})
	ans := 0
	for i, c := range cnt {
		ans += c * (i/8 + 1)
	}
	return ans
}
```

## 3020 — Find The Maximum Number Of Elements In Subset

```go
package main

// LeetCode #3020: Find the Maximum Number of Elements in Subset
// https://leetcode.com/problems/find-the-maximum-number-of-elements-in-subset/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(maximumLength([]int{5, 4, 1, 2, 2}))
	fmt.Println(maximumLength([]int{1, 3, 2, 4}))
	fmt.Println(maximumLength([]int{1, 1}))
}

func maximumLength(nums []int) int {
	cnt := map[int]int{}
	for _, x := range nums {
		cnt[x]++
	}

	ans := 0
	seen := map[int]bool{}

	// x = 1 special: 1^2 = 1, all 1's chain together
	if c := cnt[1]; c > 0 {
		seen[1] = true
		if c > ans {
			ans = c
		}
	}

	for x := range cnt {
		if seen[x] || x == 1 {
			continue
		}
		chain := []int{}
		for y := x; cnt[y] > 0 && y <= 1e9; y = y * y {
			chain = append(chain, y)
			seen[y] = true
		}
		if len(chain) == 0 {
			continue
		}
		result := 0
		for i := len(chain) - 1; i >= 0; i-- {
			v := chain[i]
			c := cnt[v]
			if result == 0 {
				result = 1
				c--
			}
			take := c
			if take > 2 {
				take = 2
			}
			result += take
		}
		if result > ans {
			ans = result
		}
	}
	if ans == 0 {
		ans = 1
	}
	return ans
}
```

## 3021 — Alice And Bob Playing Flower Game

```go
package main

// LeetCode #3021: Alice and Bob Playing Flower Game
// https://leetcode.com/problems/alice-and-bob-playing-flower-game/
// Difficulty: Medium
// Time: O(1) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(flowerGame(3, 2))
	fmt.Println(flowerGame(1, 1))
	fmt.Println(flowerGame(4, 4))
}

func flowerGame(n int, m int) int64 {
	oddN := int64((n + 1) / 2)
	evenN := int64(n / 2)
	oddM := int64((m + 1) / 2)
	evenM := int64(m / 2)
	return oddN*evenM + evenN*oddM
}
```

## 3023 — Find Pattern In Infinite Stream I

```go
package main

// LeetCode #3023: Find Pattern in Infinite Stream I (PAID)
// https://leetcode.com/problems/find-pattern-in-infinite-stream-i/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(m) where m = len(pattern) ≤ 100

// Given a binary array pattern and an InfiniteStream that yields bits via Next(),
// find the first index (0-indexed) where pattern finishes matching as a
// contiguous subsequence in the stream.

import "fmt"

// InfiniteStream simulates the LeetCode API. Internally backed by a repeating
// slice so we can test without an actual infinite source.
type InfiniteStream struct {
	data []int
	pos  int
}

func NewInfiniteStream(data []int) *InfiniteStream {
	return &InfiniteStream{data: data, pos: 0}
}

// Next returns the next bit. The internal data cycles through the given
// slice to simulate an infinite source. With pattern length ≤ 100 and
// typical data ≤ 10⁴, the pattern will repeat quickly.
func (s *InfiniteStream) Next() int {
	val := s.data[s.pos%len(s.data)]
	s.pos++
	return val
}

func main() {
	// Test: pattern appears starting at index 2
	stream := NewInfiniteStream([]int{0, 1, 0, 1, 0, 1})
	fmt.Println(findPattern(stream, []int{0, 1, 0})) // 2

	// Test: pattern at a later index
	stream2 := NewInfiniteStream([]int{1, 1, 1, 0, 1, 0, 0})
	fmt.Println(findPattern(stream2, []int{1, 0})) // 3

	// Test: single-bit pattern
	stream3 := NewInfiniteStream([]int{1, 0, 0, 0})
	fmt.Println(findPattern(stream3, []int{1})) // 0
}

func findPattern(stream *InfiniteStream, pattern []int) int {
	m := len(pattern)
	if m == 0 {
		return 0
	}

	// Build pattern bitmask (pattern[0] is shifted into the most
	// significant position of the mask).
	patternMask := 0
	for _, bit := range pattern {
		patternMask = (patternMask << 1) | bit
	}

	mask := (1 << m) - 1
	window := 0
	idx := 0
	// Safety upper bound: the pattern must appear at some point in an
	// underlying periodic sequence; we scan up to 10⁶ positions.
	maxIter := 1_000_000

	for idx < maxIter {
		bit := stream.Next()
		window = ((window << 1) | bit) & mask
		if idx >= m-1 && window == patternMask {
			return idx
		}
		idx++
	}
	return -1
}
```

## 3025 — Find The Number Of Ways To Place People I

```go
package main

// LeetCode #3025: Find the Number of Ways to Place People I
// https://leetcode.com/problems/find-the-number-of-ways-to-place-people-i/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(numberOfPairs([][]int{{1, 1}, {2, 2}, {3, 3}}))
	fmt.Println(numberOfPairs([][]int{{6, 2}, {4, 4}, {2, 6}}))
	fmt.Println(numberOfPairs([][]int{{3, 1}, {1, 3}, {1, 1}}))
}

func numberOfPairs(points [][]int) (ans int) {
	n := len(points)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			x1, y1 := points[i][0], points[i][1]
			x2, y2 := points[j][0], points[j][1]
			if x1 > x2 || y1 < y2 {
				continue
			}
			ok := true
			for k := 0; k < n; k++ {
				if k == i || k == j {
					continue
				}
				x3, y3 := points[k][0], points[k][1]
				if x3 >= x2 && x3 <= x1 && y3 >= y2 && y3 <= y1 {
					if x3 == x2 && y3 == y2 || x3 == x1 && y3 == y1 {
						continue
					}
					if x3 >= x2 && x3 <= x1 && y3 >= y2 && y3 <= y1 {
						ok = false
						break
					}
				}
			}
			if ok {
				ans++
			}
		}
	}
	return
}
```

## 3026 — Maximum Good Subarray Sum

```go
package main

// LeetCode #3026: Maximum Good Subarray Sum
// https://leetcode.com/problems/maximum-good-subarray-sum/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(maximumSum([]int{1, 2, 3, 4, 5, 6}, 1))
	fmt.Println(maximumSum([]int{-1, 3, 2, 4, 5}, 3))
	fmt.Println(maximumSum([]int{-1, -2, -3, -4}, 2))
}

func maximumSum(nums []int, k int) int64 {
	const negInf int64 = -(1 << 60)
	ans := negInf
	pref := int64(0)
	first := map[int]int64{}
	for _, x := range nums {
		if v, ok := first[x]; ok {
			if pref+int64(x)-v > ans {
				ans = pref + int64(x) - v
			}
			if pref < v {
				first[x] = pref
			}
		} else {
			first[x] = pref
		}
		pref += int64(x)
		need := x - k
		if v, ok := first[need]; ok {
			if pref-v > ans {
				ans = pref - v
			}
		}
		need2 := x + k
		if v, ok := first[need2]; ok {
			if pref-v > ans {
				ans = pref - v
			}
		}
	}
	if ans == negInf {
		return 0
	}
	return ans
}
```

## 3029 — Minimum Time To Revert Word To Initial State I

```go
package main

// LeetCode #3029: Minimum Time to Revert Word to Initial State I
// https://leetcode.com/problems/minimum-time-to-revert-word-to-initial-state-i/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(minimumTimeToInitialState("abacaba", 3))
	fmt.Println(minimumTimeToInitialState("abacaba", 2))
	fmt.Println(minimumTimeToInitialState("abcbabcd", 2))
}

func minimumTimeToInitialState(word string, k int) int {
	n := len(word)
	pi := make([]int, n)
	for i := 1; i < n; i++ {
		j := pi[i-1]
		for j > 0 && word[i] != word[j] {
			j = pi[j-1]
		}
		if word[i] == word[j] {
			j++
		}
		pi[i] = j
	}
	j := n
	for j > 0 && j > n-k {
		j = pi[j-1]
	}
	for t := 1; ; t++ {
		if t*k >= n {
			return t
		}
		if n-t*k <= j && (n-t*k)%k == 0 {
			return t
		}
	}
}
```

## 3030 — Find The Grid Of Region Average

```go
package main

// LeetCode #3030: Find the Grid of Region Average
// https://leetcode.com/problems/find-the-grid-of-region-average/
// Difficulty: Medium
// Time: O(m*n) | Space: O(m*n)

import "fmt"

func main() {
	fmt.Println(resultGrid([][]int{{5, 6, 7}, {8, 9, 10}, {11, 12, 13}}, 2))
	fmt.Println(resultGrid([][]int{{10, 20, 30}, {15, 25, 30}, {20, 30, 40}}, 5))
	fmt.Println(resultGrid([][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}, 0))
}

func resultGrid(image [][]int, threshold int) [][]int {
	m, n := len(image), len(image[0])
	sum := make([][]int, m)
	cnt := make([][]int, m)
	for i := range sum {
		sum[i] = make([]int, n)
		cnt[i] = make([]int, n)
	}

	for i := 0; i+2 < m; i++ {
		for j := 0; j+2 < n; j++ {
			ok := true
			total := 0
		check:
			for x := i; x < i+3; x++ {
				for y := j; y < j+3; y++ {
					total += image[x][y]
					if x > i && abs(image[x][y]-image[x-1][y]) > threshold {
						ok = false
						break check
					}
					if y > j && abs(image[x][y]-image[x][y-1]) > threshold {
						ok = false
						break check
					}
				}
			}
			if !ok {
				continue
			}
			avg := total / 9
			for x := i; x < i+3; x++ {
				for y := j; y < j+3; y++ {
					sum[x][y] += avg
					cnt[x][y]++
				}
			}
		}
	}

	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
		for j := range ans[i] {
			if cnt[i][j] == 0 {
				ans[i][j] = image[i][j]
			} else {
				ans[i][j] = sum[i][j] / cnt[i][j]
			}
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

## 3034 — Number Of Subarrays That Match A Pattern I

```go
package main

// LeetCode #3034: Number of Subarrays That Match a Pattern I
// https://leetcode.com/problems/number-of-subarrays-that-match-a-pattern-i/
// Difficulty: Medium
// Time: O(n*m) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(countMatchingSubarrays([]int{1, 2, 3, 4, 5, 6}, []int{1, 1}))
	fmt.Println(countMatchingSubarrays([]int{1, 4, 4, 1, 3, 5, 5, 3}, []int{1, 0, -1}))
}

func countMatchingSubarrays(nums []int, pattern []int) (ans int) {
	n, m := len(nums), len(pattern)
outer:
	for i := 0; i+m < n; i++ {
		for k := 0; k < m; k++ {
			diff := 0
			if nums[i+k+1] > nums[i+k] {
				diff = 1
			} else if nums[i+k+1] < nums[i+k] {
				diff = -1
			}
			if diff != pattern[k] {
				continue outer
			}
		}
		ans++
	}
	return
}
```

## 3035 — Maximum Palindromes After Operations

```go
package main

// LeetCode #3035: Maximum Palindromes After Operations
// https://leetcode.com/problems/maximum-palindromes-after-operations/
// Difficulty: Medium
// Time: O(n * L + A log A) | Space: O(A)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxPalindromesAfterOperations([]string{"abbb", "ba", "aa"}))
	fmt.Println(maxPalindromesAfterOperations([]string{"abc", "ab"}))
	fmt.Println(maxPalindromesAfterOperations([]string{"cd", "ef", "a"}))
}

func maxPalindromesAfterOperations(words []string) int {
	freq := [26]int{}
	lens := make([]int, len(words))
	for i, w := range words {
		lens[i] = len(w)
		for _, ch := range w {
			freq[ch-'a']++
		}
	}
	pairs := 0
	for _, c := range freq {
		pairs += c / 2
	}
	sort.Slice(lens, func(i, j int) bool {
		return lens[i] < lens[j]
	})
	ans := 0
	for _, l := range lens {
		need := l / 2
		if pairs >= need {
			pairs -= need
			ans++
		}
	}
	return ans
}
```

## 3039 — Apply Operations To Make String Empty

```go
package main

// LeetCode #3039: Apply Operations to Make String Empty
// https://leetcode.com/problems/apply-operations-to-make-string-empty/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(lastNonEmptyString("aabcbbca"))
	fmt.Println(lastNonEmptyString("abcd"))
	fmt.Println(lastNonEmptyString("aaa"))
}

func lastNonEmptyString(s string) string {
	cnt := [26]int{}
	last := [26]int{}
	for i, ch := range s {
		idx := ch - 'a'
		cnt[idx]++
		last[idx] = i
	}
	maxFreq := 0
	for _, c := range cnt {
		if c > maxFreq {
			maxFreq = c
		}
	}
	type pair struct {
		pos int
		ch  byte
	}
	cands := []pair{}
	for i := 0; i < 26; i++ {
		if cnt[i] == maxFreq {
			cands = append(cands, pair{last[i], byte('a' + i)})
		}
	}
	// Sort by last occurrence position
	for i := 0; i < len(cands); i++ {
		for j := i + 1; j < len(cands); j++ {
			if cands[i].pos > cands[j].pos {
				cands[i], cands[j] = cands[j], cands[i]
			}
		}
	}
	ans := make([]byte, len(cands))
	for i, c := range cands {
		ans[i] = c.ch
	}
	return string(ans)
}
```

## 3040 — Maximum Number Of Operations With The Same Score Ii

```go
package main

// LeetCode #3040: Maximum Number of Operations With the Same Score II
// https://leetcode.com/problems/maximum-number-of-operations-with-the-same-score-ii/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n^2)

import "fmt"

func main() {
	fmt.Println(maxOperations3040([]int{3, 2, 1, 2, 3, 4}))
	fmt.Println(maxOperations3040([]int{3, 2, 6, 1, 4}))
}

func maxOperations3040(nums []int) int {
	n := len(nums)
	memo := map[[2]int]int{}

	var dfs func(l, r, target int) int
	dfs = func(l, r, target int) int {
		if l >= r {
			return 0
		}
		key := [2]int{l, r}
		if v, ok := memo[key]; ok {
			return v
		}
		best := 0
		if nums[l]+nums[l+1] == target {
			if res := 1 + dfs(l+2, r, target); res > best {
				best = res
			}
		}
		if nums[r-1]+nums[r] == target {
			if res := 1 + dfs(l, r-2, target); res > best {
				best = res
			}
		}
		if nums[l]+nums[r] == target {
			if res := 1 + dfs(l+1, r-1, target); res > best {
				best = res
			}
		}
		memo[key] = best
		return best
	}

	ans := 0
	// Try all 3 possible targets from first operation
	targets := []int{nums[0] + nums[1], nums[n-2] + nums[n-1], nums[0] + nums[n-1]}
	for _, t := range targets {
		memo = map[[2]int]int{}
		if res := dfs(0, n-1, t); res > ans {
			ans = res
		}
	}
	return ans
}
```

## 3043 — Find The Length Of The Longest Common Prefix

```go
package main

// LeetCode #3043: Find the Length of the Longest Common Prefix
// https://leetcode.com/problems/find-the-length-of-the-longest-common-prefix/
// Difficulty: Medium
// Time: O(n*logM + m*logM) | Space: O(n*logM)

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(longestCommonPrefix([]int{1, 10, 100}, []int{1000}))
	fmt.Println(longestCommonPrefix([]int{1, 2, 3}, []int{4, 4, 4}))
}

func longestCommonPrefix(arr1 []int, arr2 []int) int {
	prefixes := map[int]bool{}
	for _, x := range arr1 {
		for x > 0 {
			prefixes[x] = true
			x /= 10
		}
	}
	ans := 0
	for _, x := range arr2 {
		for x > 0 {
			if prefixes[x] {
				if len(strconv.Itoa(x)) > ans {
					ans = len(strconv.Itoa(x))
				}
				break
			}
			x /= 10
		}
	}
	return ans
}
```

