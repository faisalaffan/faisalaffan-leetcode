# Medium (Sedang) — Problem ��3961

## 3829 — Design Ride Sharing System

```go
package main

// LeetCode #3829: Design Ride Sharing System
// https://leetcode.com/problems/design-ride-sharing-system/
// Difficulty: Medium
// Time: O(1) per operation | Space: O(N)
// Approach: Use two FIFO queues for riders and drivers with timestamp ordering.

import "fmt"

type RideSharingSystem struct {
	t       int
	riders  [][2]int // (timestamp, riderId)
	drivers [][2]int // (timestamp, driverId)
	// rider timestamp lookup for cancel
	riderTS map[int]int
}

func ConstructorRideSharing() RideSharingSystem {
	return RideSharingSystem{
		riderTS: make(map[int]int),
	}
}

func (rs *RideSharingSystem) AddRider(riderId int) {
	rs.riderTS[riderId] = rs.t
	rs.riders = append(rs.riders, [2]int{rs.t, riderId})
	rs.t++
}

func (rs *RideSharingSystem) AddDriver(driverId int) {
	rs.drivers = append(rs.drivers, [2]int{rs.t, driverId})
	rs.t++
}

func (rs *RideSharingSystem) MatchDriverWithRider() [2]int {
	if len(rs.riders) == 0 || len(rs.drivers) == 0 {
		return [2]int{-1, -1}
	}
	driver := rs.drivers[0]
	rs.drivers = rs.drivers[1:]
	rider := rs.riders[0]
	rs.riders = rs.riders[1:]
	return [2]int{driver[1], rider[1]}
}

func (rs *RideSharingSystem) CancelRider(riderId int) {
	ts, ok := rs.riderTS[riderId]
	if !ok {
		return
	}
	for i, r := range rs.riders {
		if r[0] == ts && r[1] == riderId {
			rs.riders = append(rs.riders[:i], rs.riders[i+1:]...)
			break
		}
	}
	delete(rs.riderTS, riderId)
}

func main() {
	rs := ConstructorRideSharing()
	rs.AddRider(1)
	rs.AddDriver(10)
	rs.AddRider(2)
	rs.AddDriver(20)
	rs.AddRider(3)
	fmt.Println(rs.MatchDriverWithRider()) // Expected: [10 1]
	fmt.Println(rs.MatchDriverWithRider()) // Expected: [20 2]
	rs.CancelRider(3)
	fmt.Println(rs.MatchDriverWithRider()) // Expected: [-1 -1]
}
```

## 3831 — Median Of A Binary Search Tree Level

```go
package main

// LeetCode #3831: Median of a Binary Search Tree Level
// https://leetcode.com/problems/median-of-a-binary-search-tree-level/
// Difficulty: Medium [Paid]
// Time: O(N log W) | Space: O(W)
// Approach: DFS to collect all values at target level, sort, compute upper median.

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MedianOfABinarySearchTreeLevel(root *TreeNode, level int) int {
	values := []int{}

	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if depth == level {
			values = append(values, node.Val)
		}
		dfs(node.Left, depth+1)
		dfs(node.Right, depth+1)
	}

	dfs(root, 0)

	if len(values) == 0 {
		return -1
	}

	sort.Ints(values)
	return values[len(values)/2] // upper median
}

func main() {
	// Example 1: root = [4,null,5,null,7], level = 2
	root1 := &TreeNode{4, nil, &TreeNode{5, nil, &TreeNode{7, nil, nil}}}
	fmt.Println(MedianOfABinarySearchTreeLevel(root1, 2)) // Expected: 7

	// Example 2: root = [6,3,8], level = 1
	root2 := &TreeNode{6, &TreeNode{3, nil, nil}, &TreeNode{8, nil, nil}}
	fmt.Println(MedianOfABinarySearchTreeLevel(root2, 1)) // Expected: 8

	// Example 3: root = [2,1], level = 2
	root3 := &TreeNode{2, &TreeNode{1, nil, nil}, nil}
	fmt.Println(MedianOfABinarySearchTreeLevel(root3, 2)) // Expected: -1
}
```

## 3834 — Merge Adjacent Equal Elements

```go
package main

// LeetCode #3834: Merge Adjacent Equal Elements
// https://leetcode.com/problems/merge-adjacent-equal-elements/
// Difficulty: Medium
// Time: O(N) | Space: O(N)
// Approach: Use a stack to repeatedly merge leftmost adjacent equal pairs.

import "fmt"

func MergeAdjacentEqualElements(nums []int) []int64 {
	stack := make([]int64, 0)

	for _, v := range nums {
		cur := int64(v)
		// While stack top equals cur, merge (pop + double)
		for len(stack) > 0 && stack[len(stack)-1] == cur {
			cur += stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, cur)
	}

	return stack
}

func main() {
	// Example 1
	fmt.Println(MergeAdjacentEqualElements([]int{3, 1, 1, 2})) // Expected: [3 4]

	// Example 2
	fmt.Println(MergeAdjacentEqualElements([]int{2, 2, 4})) // Expected: [8]

	// Example 3
	fmt.Println(MergeAdjacentEqualElements([]int{3, 7, 5})) // Expected: [3 7 5]
}
```

## 3835 — Count Subarrays With Cost Less Than Or Equal To K

```go
package main

// LeetCode #3835: Count Subarrays With Cost Less Than or Equal to K
// https://leetcode.com/problems/count-subarrays-with-cost-less-than-or-equal-to-k/
// Difficulty: Medium
// Time: O(N) | Space: O(N)
// Approach: Sliding window with two monotonic deques for max and min.
// cost = (max - min) * length, expand right, shrink left when cost > k.

import "fmt"

func CountSubarraysWithCostLessThanOrEqualToK(nums []int, k int) int {
	n := len(nums)
	ans := 0
	left := 0

	// Monotonic deques for max (decreasing) and min (increasing)
	maxQ := make([]int, 0) // indices, values decreasing
	minQ := make([]int, 0) // indices, values increasing

	for right := 0; right < n; right++ {
		// Add nums[right] to max deque
		for len(maxQ) > 0 && nums[maxQ[len(maxQ)-1]] <= nums[right] {
			maxQ = maxQ[:len(maxQ)-1]
		}
		maxQ = append(maxQ, right)

		// Add nums[right] to min deque
		for len(minQ) > 0 && nums[minQ[len(minQ)-1]] >= nums[right] {
			minQ = minQ[:len(minQ)-1]
		}
		minQ = append(minQ, right)

		// Shrink window while cost > k
		for left <= right {
			curMin := nums[minQ[0]]
			curMax := nums[maxQ[0]]
			cost := (curMax - curMin) * (right - left + 1)
			if cost <= k {
				break
			}
			// Remove left from deques if at front
			if len(maxQ) > 0 && maxQ[0] == left {
				maxQ = maxQ[1:]
			}
			if len(minQ) > 0 && minQ[0] == left {
				minQ = minQ[1:]
			}
			left++
		}

		ans += right - left + 1
	}

	return ans
}

func main() {
	// Example 1
	fmt.Println(CountSubarraysWithCostLessThanOrEqualToK([]int{1, 3, 2}, 4)) // Expected: 5

	// Example 2
	fmt.Println(CountSubarraysWithCostLessThanOrEqualToK([]int{5, 5, 5, 5}, 0)) // Expected: 10

	// Example 3
	fmt.Println(CountSubarraysWithCostLessThanOrEqualToK([]int{1, 2, 3}, 0)) // Expected: 3
}
```

## 3837 — Delayed Count Of Equal Elements

```go
package main

// LeetCode #3837: Delayed Count of Equal Elements
// https://leetcode.com/problems/delayed-count-of-equal-elements/
// Difficulty: Medium [Paid]
// Time: O(N) | Space: O(N)
// Approach: Track frequency of each value and count pairs where values are equal.

import "fmt"

func DelayedCountOfEqualElements(nums []int) int64 {
	freq := make(map[int]int64)
	var ans int64

	for _, v := range nums {
		ans += freq[v]
		freq[v]++
	}

	return ans
}

func main() {
	// Example 1
	fmt.Println(DelayedCountOfEqualElements([]int{1, 2, 1, 2, 1})) // Expected: 4

	// Example 2
	fmt.Println(DelayedCountOfEqualElements([]int{1, 1, 1, 1})) // Expected: 6

	// Example 3
	fmt.Println(DelayedCountOfEqualElements([]int{1, 2, 3})) // Expected: 0
}
```

## 3839 — Number Of Prefix Connected Groups

```go
package main

// LeetCode #3839: Number of Prefix Connected Groups
// https://leetcode.com/problems/number-of-prefix-connected-groups/
// Difficulty: Medium
// Time: O(N * K) | Space: O(N)
// Approach: Group words by first k characters (prefix). Count groups with >= 2 words.

import "fmt"

func NumberOfPrefixConnectedGroups(words []string, k int) int {
	prefixCount := make(map[string]int)

	for _, w := range words {
		if len(w) < k {
			continue
		}
		prefixCount[w[:k]]++
	}

	ans := 0
	for _, cnt := range prefixCount {
		if cnt >= 2 {
			ans++
		}
	}
	return ans
}

func main() {
	// Example 1
	fmt.Println(NumberOfPrefixConnectedGroups([]string{"apple", "apply", "banana", "bandit"}, 2)) // Expected: 2

	// Example 2
	fmt.Println(NumberOfPrefixConnectedGroups([]string{"car", "cat", "cartoon"}, 3)) // Expected: 1

	// Example 3
	fmt.Println(NumberOfPrefixConnectedGroups([]string{"bat", "dog", "dog", "doggy", "bat"}, 3)) // Expected: 2
}
```

## 3840 — House Robber V

```go
package main

// LeetCode #3840: House Robber V
// https://leetcode.com/problems/house-robber-v/
// Difficulty: Medium
// Time: O(N) | Space: O(1)
// Approach: DP with two states (rob/notRob). If adjacent houses have same color,
// cannot rob both. If different colors, can rob both.

import "fmt"

func HouseRobberV(nums []int, colors []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	notRob, rob := 0, nums[0]

	for i := 1; i < n; i++ {
		newNotRob := max(notRob, rob)
		var newRob int
		if colors[i] != colors[i-1] {
			newRob = max(notRob, rob) + nums[i]
		} else {
			newRob = notRob + nums[i]
		}
		notRob, rob = newNotRob, newRob
	}

	return max(notRob, rob)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Example 1
	fmt.Println(HouseRobberV([]int{1, 4, 3, 5}, []int{1, 1, 2, 2})) // Expected: 9

	// Example 2
	fmt.Println(HouseRobberV([]int{3, 1, 2, 4}, []int{2, 3, 2, 2})) // Expected: 8

	// Example 3
	fmt.Println(HouseRobberV([]int{10, 1, 3, 9}, []int{1, 1, 1, 2})) // Expected: 22
}
```

## 3843 — First Element With Unique Frequency

```go
package main

// LeetCode #3843: First Element with Unique Frequency
// https://leetcode.com/problems/first-element-with-unique-frequency/
// Difficulty: Medium
// Time: O(N) | Space: O(N)
// Approach: Count frequencies, then count frequency-of-frequency, find first with freq=1.

import "fmt"

func FirstElementWithUniqueFrequency(nums []int) int {
	cnt := make(map[int]int)
	for _, v := range nums {
		cnt[v]++
	}

	freqCnt := make(map[int]int)
	for _, c := range cnt {
		freqCnt[c]++
	}

	for _, v := range nums {
		if freqCnt[cnt[v]] == 1 {
			return v
		}
	}

	return -1
}

func main() {
	// Example 1
	fmt.Println(FirstElementWithUniqueFrequency([]int{20, 10, 30, 30})) // Expected: 30

	// Example 2
	fmt.Println(FirstElementWithUniqueFrequency([]int{20, 20, 10, 30, 30, 30})) // Expected: 20

	// Example 3
	fmt.Println(FirstElementWithUniqueFrequency([]int{10, 10, 20, 20})) // Expected: -1
}
```

## 3844 — Longest Almost Palindromic Substring

```go
package main

// LeetCode #3844: Longest Almost-Palindromic Substring
// https://leetcode.com/problems/longest-almost-palindromic-substring/
// Difficulty: Medium
// Time: O(N^2) | Space: O(1)
// Approach: Expand from each center. Track longest palindrome with at most
// one character removal (either via skipping a mismatch or by extending past
// the palindrome boundary by one).

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func LongestAlmostPalindromicSubstring(s string) int {
	n := len(s)
	ans := 1

	// expand from center (l, r), allowing one skip
	check := func(l, r int) {
		// Phase 1: expand matching chars
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
		}

		pureLen := r - l - 1
		ans = max(ans, pureLen)

		// If we hit a boundary (l < 0 or r >= n), try including one more
		// char on the available side, then "removing" it.
		if l >= 0 && r >= n {
			// both sides have boundary: can't extend on either without breaking palindrome
			// but we can include the char at l and remove it
			ans = max(ans, pureLen+1)
		} else if l >= 0 && r < n {
			// left has char, right has char but they don't match (mismatch)
			// Try skipping left OR right, then continue expanding
			// skip left
			nl, nr := l-1, r
			for nl >= 0 && nr < n && s[nl] == s[nr] {
				nl--
				nr++
			}
			ans = max(ans, nr-nl-1)
			// skip right
			nl, nr = l, r+1
			for nl >= 0 && nr < n && s[nl] == s[nr] {
				nl--
				nr++
			}
			ans = max(ans, nr-nl-1)
		} else if l < 0 && r < n {
			// left boundary, right has char
			// include r and "remove" it
			ans = max(ans, pureLen+1)
		}
		// l < 0 && r >= n: both boundaries, nothing to add
	}

	for i := 0; i < n; i++ {
		check(i, i)   // odd length center
		check(i, i+1) // even length center
	}

	return ans
}

func main() {
	// Example 1
	fmt.Println(LongestAlmostPalindromicSubstring("abca")) // Expected: 4

	// Example 2
	fmt.Println(LongestAlmostPalindromicSubstring("abba")) // Expected: 4

	// Example 3
	fmt.Println(LongestAlmostPalindromicSubstring("zzabba")) // Expected: 5
}
```

## 3846 — Total Distance To Type A String Using One Finger

```go
package main

// LeetCode #3846: Total Distance to Type a String Using One Finger
// https://leetcode.com/problems/total-distance-to-type-a-string-using-one-finger/
// Difficulty: Medium [Paid]
// Time: O(N) | Space: O(1)
// Approach: Precompute keyboard positions, simulate typing from 'a'.

import "fmt"

func TotalDistanceToTypeAStringUsingOneFinger(s string) int {
	// Keyboard layout (row, col)
	keyboard := []string{
		"qwertyuiop",
		"asdfghjkl",
		"zxcvbnm",
	}

	pos := make(map[byte][2]int)
	for r, row := range keyboard {
		for c, ch := range row {
			pos[byte(ch)] = [2]int{r, c}
		}
	}

	total := 0
	cur := pos['a']
	for i := 0; i < len(s); i++ {
		next := pos[s[i]]
		dist := abs(cur[0]-next[0]) + abs(cur[1]-next[1])
		total += dist
		cur = next
	}

	return total
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	// Example 1
	fmt.Println(TotalDistanceToTypeAStringUsingOneFinger("hello")) // Expected: 17

	// Example 2
	fmt.Println(TotalDistanceToTypeAStringUsingOneFinger("a")) // Expected: 0

	// Example 3
	fmt.Println(TotalDistanceToTypeAStringUsingOneFinger("qaz")) // q: (0,0), a: (1,0), z: (2,0) = |0-1|+|0-0| + |1-2|+|0-0| = 1+1 = 2
}
```

## 3847 — Find The Score Difference In A Game

```go
package main

// LeetCode #3847: Find the Score Difference in a Game
// https://leetcode.com/problems/find-the-score-difference-in-a-game/
// Difficulty: Medium
// Time: O(N) | Space: O(1)
// Approach: Track active player and swap on odd or every 6th game.

import "fmt"

func FindTheScoreDifferenceInAGame(nums []int) int {
	first, second := 0, 0
	activeIsFirst := true

	for i, v := range nums {
		// Swap if odd
		if v%2 == 1 {
			activeIsFirst = !activeIsFirst
		}
		// Swap every 6th game (0-indexed, so i%6 == 5)
		if i%6 == 5 {
			activeIsFirst = !activeIsFirst
		}
		if activeIsFirst {
			first += v
		} else {
			second += v
		}
	}

	return first - second
}

func main() {
	// Example 1
	fmt.Println(FindTheScoreDifferenceInAGame([]int{1, 2, 3})) // Expected: 0

	// Example 2
	fmt.Println(FindTheScoreDifferenceInAGame([]int{2, 4, 2, 1, 2, 1})) // Expected: 4

	// Example 3
	fmt.Println(FindTheScoreDifferenceInAGame([]int{1})) // Expected: -1
}
```

## 3848 — Check Digitorial Permutation

```go
package main

// LeetCode #3848: Check Digitorial Permutation
// https://leetcode.com/problems/check-digitorial-permutation/
// Difficulty: Medium
// Time: O(log N) | Space: O(1)
// Approach: Compute sum of factorials of digits, check if any permutation
// of n equals that sum (i.e., they have same digit frequency).

import "fmt"

func CheckDigitorialPermutation(n int) bool {
	// Precompute factorials for digits 0-9
	fact := []int{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880}

	// Compute sum of factorials of digits
	temp := n
	sum := 0
	for temp > 0 {
		sum += fact[temp%10]
		temp /= 10
	}

	// Check if sum has same digits as n
	// If they have the same digits, one is a permutation of the other
	digitsN := make([]int, 10)
	digitsSum := make([]int, 10)

	temp = n
	for temp > 0 {
		digitsN[temp%10]++
		temp /= 10
	}

	temp = sum
	for temp > 0 {
		digitsSum[temp%10]++
		temp /= 10
	}

	// Handle sum = 0 (if n = 0, but n >= 1 per constraints)
	if sum == 0 {
		return false
	}

	for i := 0; i < 10; i++ {
		if digitsN[i] != digitsSum[i] {
			return false
		}
	}

	return true
}

func main() {
	// Example 1
	fmt.Println(CheckDigitorialPermutation(145)) // Expected: true

	// Example 2
	fmt.Println(CheckDigitorialPermutation(10)) // Expected: false

	// Example 3
	fmt.Println(CheckDigitorialPermutation(40585)) // 4!+0!+5!+8!+5! = 24+1+120+40320+120 = 40585
}
```

## 3849 — Maximum Bitwise Xor After Rearrangement

```go
package main

// LeetCode #3849: Maximum Bitwise XOR After Rearrangement
// https://leetcode.com/problems/maximum-bitwise-xor-after-rearrangement/
// Difficulty: Medium
// Time: O(N) | Space: O(1)
// Approach: Count 0s and 1s in t. Greedily match opposite bits for max XOR.

import "fmt"

func MaximumBitwiseXorAfterRearrangement(s string, t string) string {
	ones, zeros := 0, 0
	for _, ch := range t {
		if ch == '1' {
			ones++
		} else {
			zeros++
		}
	}

	ans := make([]byte, len(s))
	for i, ch := range s {
		if ch == '1' {
			if zeros > 0 {
				ans[i] = '1'
				zeros--
			} else {
				ans[i] = '0'
				ones--
			}
		} else {
			if ones > 0 {
				ans[i] = '1'
				ones--
			} else {
				ans[i] = '0'
				zeros--
			}
		}
	}

	return string(ans)
}

func main() {
	// Example 1
	fmt.Println(MaximumBitwiseXorAfterRearrangement("101", "011")) // Expected: "110"

	// Example 2
	fmt.Println(MaximumBitwiseXorAfterRearrangement("0110", "1110")) // Expected: "1101"

	// Example 3
	fmt.Println(MaximumBitwiseXorAfterRearrangement("0101", "1001")) // Expected: "1111"
}
```

## 3851 — Maximum Requests Without Violating The Limit

```go
package main

// LeetCode #3851: Maximum Requests Without Violating the Limit
// https://leetcode.com/problems/maximum-requests-without-violating-the-limit/
// Difficulty: Medium [Paid]
// Time: O(N log N) | Space: O(N)
// Approach: For each user, group requests by time. Use sliding window to
// find max requests that can be kept per user without exceeding k in any
// window of size `window`.

import (
	"fmt"
	"sort"
)

func MaximumRequestsWithoutViolatingTheLimit(requests [][]int, k int, window int) int {
	// Group requests by user
	userRequests := make(map[int][]int)
	for _, r := range requests {
		user, time := r[0], r[1]
		userRequests[user] = append(userRequests[user], time)
	}

	total := 0

	for _, times := range userRequests {
		sort.Ints(times)
		// Use DP to find max requests we can keep
		// For each request, we can either keep it or drop it
		n := len(times)
		dp := make([]int, n+1)
		for i := 0; i < n; i++ {
			// Drop this request
			dp[i+1] = max(dp[i+1], dp[i])
			// Keep this request, find how many we can keep that end before times[i]-window
			// We want the first j where times[j] > times[i] - window - 1
			j := sort.Search(n, func(x int) bool { return times[x] > times[i]-window-1 })
			count := i - j + 1
			if count <= k {
				dp[i+1] = max(dp[i+1], dp[j]+count)
			}
		}
		total += dp[n]
	}

	return total
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Example 1
	req1 := [][]int{{1, 1}, {2, 1}, {1, 7}, {2, 8}}
	fmt.Println(MaximumRequestsWithoutViolatingTheLimit(req1, 1, 4)) // Expected: 4

	// Example 2
	req2 := [][]int{{1, 2}, {1, 5}, {1, 2}, {1, 6}}
	fmt.Println(MaximumRequestsWithoutViolatingTheLimit(req2, 2, 5)) // Expected: 2

	// Example 3
	req3 := [][]int{{1, 1}, {2, 5}, {1, 2}, {3, 9}}
	fmt.Println(MaximumRequestsWithoutViolatingTheLimit(req3, 1, 1)) // Expected: 3
}
```

## 3853 — Merge Close Characters

```go
package main

// LeetCode #3853: Merge Close Characters
// https://leetcode.com/problems/merge-close-characters/
// Difficulty: Medium
// Time: O(N^2) | Space: O(N)
// Approach: Simulate merging. Track last position of each character.
// When a char appears within distance k of its last occurrence, skip it.

import "fmt"

func MergeCloseCharacters(s string, k int) string {
	result := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		ch := s[i]
		// Check if same char exists in result within distance k
		merge := false
		for j := len(result) - 1; j >= 0 && len(result)-1-j < k; j-- {
			if result[j] == ch {
				merge = true
				break
			}
		}
		if !merge {
			result = append(result, ch)
		}
	}

	return string(result)
}

func main() {
	// Example 1
	fmt.Println(MergeCloseCharacters("abca", 3)) // Expected: "abc"

	// Example 2
	fmt.Println(MergeCloseCharacters("aabca", 2)) // Expected: "abca"

	// Example 3
	fmt.Println(MergeCloseCharacters("yybyzybz", 2)) // Expected: "ybzybz"
}
```

## 3854 — Minimum Operations To Make Array Parity Alternating

```go
package main

// LeetCode #3854: Minimum Operations to Make Array Parity Alternating
// https://leetcode.com/problems/minimum-operations-to-make-array-parity-alternating/
// Difficulty: Medium
// Time: O(N) | Space: O(N)
// Approach: Compute min operations for two patterns (even-start, odd-start).
// Then compute min max-min range using sliding window on candidate values.

import (
	"fmt"
	"math"
	"sort"
)

func MinimumOperationsToMakeArrayParityAlternating(nums []int) []int {
	n := len(nums)

	// Helper to compute for a given starting parity
	solve := func(startParity int) (ops int, candidates [][2]int) {
		curParity := startParity
		for i := 0; i < n; i++ {
			if nums[i]%2 != curParity {
				ops++
				candidates = append(candidates, [2]int{nums[i] - 1, i})
				candidates = append(candidates, [2]int{nums[i] + 1, i})
			} else {
				candidates = append(candidates, [2]int{nums[i], i})
			}
			curParity ^= 1
		}
		return
	}

	ops1, cand1 := solve(0) // even at index 0
	ops2, cand2 := solve(1) // odd at index 0

	// Choose pattern with fewer ops
	var ops int
	var candidates [][2]int
	if ops1 < ops2 || (ops1 == ops2) {
		ops = ops1
		candidates = cand1
	} else {
		ops = ops2
		candidates = cand2
	}

	// Sliding window to find min range
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i][0] < candidates[j][0]
	})

	minRange := math.MaxInt32
	cnt := make([]int, n)
	empty := n
	left := 0
	for right := 0; right < len(candidates); right++ {
		idx := candidates[right][1]
		if cnt[idx] == 0 {
			empty--
		}
		cnt[idx]++
		for empty == 0 {
			r := candidates[right][0] - candidates[left][0]
			if r < minRange {
				minRange = r
			}
			if cnt[candidates[left][1]] == 1 {
				empty++
			}
			cnt[candidates[left][1]]--
			left++
		}
	}

	return []int{ops, minRange}
}

func main() {
	// Example 1
	fmt.Println(MinimumOperationsToMakeArrayParityAlternating([]int{-2, -3, 1, 4})) // Expected: [2 6]

	// Example 2
	fmt.Println(MinimumOperationsToMakeArrayParityAlternating([]int{0, 2, -2})) // Expected: [1 3]

	// Example 3
	fmt.Println(MinimumOperationsToMakeArrayParityAlternating([]int{7})) // Expected: [0 0]
}
```

## 3857 — Minimum Cost To Split Into Ones

```go
package main

// LeetCode #3857: Minimum Cost to Split into Ones
// https://leetcode.com/problems/minimum-cost-to-split-into-ones/
// Difficulty: Medium
// Time: O(1) | Space: O(1)
// Approach: Minimum cost = n*(n-1)/2. Equivalent to total edges in complete graph.

import "fmt"

func MinimumCostToSplitIntoOnes(n int) int {
	return n * (n - 1) / 2
}

func main() {
	// Example 1
	fmt.Println(MinimumCostToSplitIntoOnes(3)) // Expected: 3

	// Example 2
	fmt.Println(MinimumCostToSplitIntoOnes(4)) // Expected: 6

	// Example 3
	fmt.Println(MinimumCostToSplitIntoOnes(10))
}
```

## 3858 — Minimum Bitwise Or From Grid

```go
package main

// LeetCode #3858: Minimum Bitwise OR From Grid
// https://leetcode.com/problems/minimum-bitwise-or-from-grid/
// Difficulty: Medium
// Time: O(32 * M * N) | Space: O(1)
// Approach: Greedy bit-by-bit from MSB. Try to keep each bit as 0 if
// every row has at least one number with that bit (and all higher
// kept-zero bits) set to 0.

import "fmt"

func MinimumBitwiseOrFromGrid(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	ans := 0
	forbidden := 0

	for bit := 31; bit >= 0; bit-- {
		testForbidden := forbidden | (1 << bit)
		ok := true
		for i := 0; i < m; i++ {
			hasValid := false
			for j := 0; j < n; j++ {
				if grid[i][j]&testForbidden == 0 {
					hasValid = true
					break
				}
			}
			if !hasValid {
				ok = false
				break
			}
		}
		if ok {
			forbidden = testForbidden
		} else {
			ans |= (1 << bit)
		}
	}

	return ans
}

func main() {
	// Example 1
	fmt.Println(MinimumBitwiseOrFromGrid([][]int{{1, 5}, {2, 4}})) // Expected: 3

	// Example 2
	fmt.Println(MinimumBitwiseOrFromGrid([][]int{{3, 5}, {6, 4}})) // Expected: 5

	// Example 3
	fmt.Println(MinimumBitwiseOrFromGrid([][]int{{7, 9, 8}})) // Expected: 7
}
```

## 3860 — Unique Email Groups

```go
package main

// LeetCode #3860: Unique Email Groups
// https://leetcode.com/problems/unique-email-groups/
// Difficulty: Medium [Paid]
// Time: O(N * M) | Space: O(N * M)
// Approach: Normalize each email by processing local part (ignore dots,
// ignore after +) and lowercase everything. Count unique normalized forms.

import (
	"fmt"
	"strings"
)

func UniqueEmailGroups(emails []string) int {
	seen := make(map[string]bool)

	for _, email := range emails {
		parts := strings.SplitN(email, "@", 2)
		local, domain := parts[0], parts[1]

		var normLocal strings.Builder
		for _, ch := range local {
			if ch == '+' {
				break
			}
			if ch != '.' {
				normLocal.WriteRune(ch)
			}
		}

		normalized := strings.ToLower(normLocal.String()) + "@" + strings.ToLower(domain)
		seen[normalized] = true
	}

	return len(seen)
}

func main() {
	// Example 1
	emails1 := []string{
		"test.email+alex@leetcode.com",
		"test.e.mail+bob.cathy@leetcode.com",
		"testemail+david@lee.tcode.com",
	}
	fmt.Println(UniqueEmailGroups(emails1)) // Expected: 2

	// Example 2
	emails2 := []string{"A@B.com", "a@b.com", "ab+xy@b.com", "a.b@b.com"}
	fmt.Println(UniqueEmailGroups(emails2)) // Expected: 2

	// Example 3
	emails3 := []string{
		"a.b+c.d+e@DoMain.com",
		"ab+xyz@domain.com",
		"ab@domain.com",
	}
	fmt.Println(UniqueEmailGroups(emails3)) // Expected: 1
}
```

## 3862 — Find The Smallest Balanced Index

```go
package main

// LeetCode #3862: Find the Smallest Balanced Index
// https://leetcode.com/problems/find-the-smallest-balanced-index/
// Difficulty: Medium
// Time: O(N) | Space: O(N)
// Approach: Compute prefix sums and suffix products, check equality at each index.

import "fmt"

func FindTheSmallestBalancedIndex(nums []int) int {
	n := len(nums)

	prefixSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefixSum[i+1] = prefixSum[i] + nums[i]
	}

	suffixProd := make([]int, n+1)
	suffixProd[n] = 1
	for i := n - 1; i >= 0; i-- {
		suffixProd[i] = suffixProd[i+1] * nums[i]
	}

	for i := 0; i < n; i++ {
		leftSum := prefixSum[i]
		rightProd := suffixProd[i+1]
		if leftSum == rightProd {
			return i
		}
	}

	return -1
}

func main() {
	// Example 1
	fmt.Println(FindTheSmallestBalancedIndex([]int{2, 1, 2})) // Expected: 1

	// Example 2
	fmt.Println(FindTheSmallestBalancedIndex([]int{2, 8, 2, 2, 5})) // Expected: 2

	// Example 3
	fmt.Println(FindTheSmallestBalancedIndex([]int{1})) // Expected: -1
}
```

## 3863 — Minimum Operations To Sort A String

```go
package main

// LeetCode #3863: Minimum Operations to Sort a String
// https://leetcode.com/problems/minimum-operations-to-sort-a-string/
// Difficulty: Medium
// Time: O(N) | Space: O(1)
// Approach: Count characters and determine min swaps to sort.

import "fmt"

func MinimumOperationsToSortAString(s string) int {
	// Count character frequencies
	cnt := make([]int, 26)
	for _, ch := range s {
		cnt[ch-'a']++
	}

	// Build sorted string
	sorted := make([]byte, len(s))
	idx := 0
	for c := 0; c < 26; c++ {
		for j := 0; j < cnt[c]; j++ {
			sorted[idx] = byte('a' + c)
			idx++
		}
	}

	// Count mismatching positions (where s[i] != sorted[i])
	mismatch := 0
	for i := 0; i < len(s); i++ {
		if s[i] != sorted[i] {
			mismatch++
		}
	}

	// Each swap fixes two mismatches, but there might be cycles
	// Simple: (mismatch + 1) / 2 gives swaps needed
	return (mismatch + 1) / 2
}

func main() {
	// Example
	fmt.Println(MinimumOperationsToSortAString("cba")) // Expected: 1

	// Example
	fmt.Println(MinimumOperationsToSortAString("aab")) // Expected: 0

	// Example
	fmt.Println(MinimumOperationsToSortAString("acb")) // Expected: 1
}
```

## 3865 — Reverse K Subarrays

```go
package main

// LeetCode #3865: Reverse K Subarrays
// https://leetcode.com/problems/reverse-k-subarrays/
// Difficulty: Medium [Paid]
// Time: O(N) | Space: O(1)
// Approach: Count positions that don't match expected value at their index.
// Each reversal of a subarray of length k can fix at most 2 positions.

import "fmt"

func ReverseKSubarrays(nums []int, k int) int {
	n := len(nums)
	swaps := 0
	for i := 0; i < n; i++ {
		if nums[i] != i { // value should equal index for sorted array [0,1,2,...]
			// Find where i is
			j := i
			for j < n && nums[j] != i {
				j++
			}
			if j-i+1 >= k && j < n {
				// Reverse subarray i..j
				for l, r := i, j; l < r; l, r = l+1, r-1 {
					nums[l], nums[r] = nums[r], nums[l]
				}
				swaps++
			}
		}
	}
	return swaps
}

func main() {
	// Example
	fmt.Println(ReverseKSubarrays([]int{1, 0, 3, 2}, 2)) // Expected: ?

	// Example
	fmt.Println(ReverseKSubarrays([]int{2, 1, 0}, 3)) // Expected: ?

	// Example
	fmt.Println(ReverseKSubarrays([]int{0, 1, 2}, 1)) // Expected: 0
}
```

## 3867 — Sum Of Gcd Of Formed Pairs

```go
package main

// LeetCode #3867: Sum of GCD of Formed Pairs
// https://leetcode.com/problems/sum-of-gcd-of-formed-pairs/
// Difficulty: Medium
// Time: O(N log M) | Space: O(N)
// Approach: Build prefixGcd array where prefixGcd[i] = gcd(nums[i], max(nums[0..i])).
// Sort, pair smallest with largest, sum gcd of each pair.

import (
	"fmt"
	"sort"
)

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func SumOfGcdOfFormedPairs(nums []int) int {
	n := len(nums)
	prefixGcd := make([]int, n)
	mx := 0
	for i, v := range nums {
		if v > mx {
			mx = v
		}
		prefixGcd[i] = gcd(v, mx)
	}

	sort.Ints(prefixGcd)

	ans := 0
	for i := 0; i < n/2; i++ {
		ans += gcd(prefixGcd[i], prefixGcd[n-1-i])
	}
	return ans
}

func main() {
	// Example 1
	fmt.Println(SumOfGcdOfFormedPairs([]int{2, 6, 4})) // Expected: 2

	// Example 2
	fmt.Println(SumOfGcdOfFormedPairs([]int{3, 6, 2, 8})) // Expected: 5

	// Extra
	fmt.Println(SumOfGcdOfFormedPairs([]int{1})) // Expected: 0
}
```

## 3868 — Minimum Cost To Equalize Arrays Using Swaps

```go
package main

// LeetCode #3868: Minimum Cost to Equalize Arrays Using Swaps
// https://leetcode.com/problems/minimum-cost-to-equalize-arrays-using-swaps/
// Difficulty: Medium
// Time: O(N) | Space: O(N)
// Approach: Count frequencies in both arrays. If any value's total count is odd,
// impossible (return -1). Min cost = sum of positive differences / 2.

import "fmt"

func MinimumCostToEqualizeArraysUsingSwaps(nums1 []int, nums2 []int) int {
	cnt := make(map[int]int)
	for _, v := range nums1 {
		cnt[v]++
	}
	for _, v := range nums2 {
		cnt[v]--
	}

	posDiff := 0
	for _, c := range cnt {
		if c%2 != 0 {
			return -1
		}
		if c > 0 {
			posDiff += c
		}
	}
	return posDiff / 2
}

func main() {
	// Example 1
	fmt.Println(MinimumCostToEqualizeArraysUsingSwaps([]int{10, 20}, []int{20, 10})) // Expected: 0

	// Example 2
	fmt.Println(MinimumCostToEqualizeArraysUsingSwaps([]int{10, 10}, []int{20, 20})) // Expected: 1

	// Example 3
	fmt.Println(MinimumCostToEqualizeArraysUsingSwaps([]int{10, 20}, []int{30, 40})) // Expected: -1
}
```

## 3871 — Count Commas In Range Ii

```go
package main

// LeetCode #3871: Count Commas in Range II
// https://leetcode.com/problems/count-commas-in-range-ii/
// Difficulty: Medium
// Time: O(log N) | Space: O(1)
// Approach: Iterate powers of 1000 starting from 1000. For each power x <= n,
// add n - x + 1 (numbers that gain a comma at this magnitude).

import "fmt"

func CountCommasInRangeIi(n int64) int64 {
	var ans int64 = 0
	for x := int64(1000); x <= n; x *= 1000 {
		ans += n - x + 1
	}
	return ans
}

func main() {
	// Example 1
	fmt.Println(CountCommasInRangeIi(1002)) // Expected: 3

	// Example 2
	fmt.Println(CountCommasInRangeIi(998)) // Expected: 0

	// Extra
	fmt.Println(CountCommasInRangeIi(1000000)) // Expected: 999001 + 1 = 999002
}
```

## 3872 — Longest Arithmetic Sequence After Changing At Most One Element

```go
package main

// LeetCode #3872: Longest Arithmetic Sequence After Changing At Most One Element
// https://leetcode.com/problems/longest-arithmetic-sequence-after-changing-at-most-one-element/
// Difficulty: Medium
// Time: O(N) | Space: O(N)
// Approach: Compute prefix (arithmetic ending at i) and suffix (arithmetic starting at i).
// For each position, try changing it to connect left and right arithmetic sequences.

import "fmt"

func LongestArithmeticSequenceAfterChangingAtMostOneElement(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}

	pref := make([]int, n)
	pref[0] = 1
	pref[1] = 2
	for i := 2; i < n; i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			pref[i] = pref[i-1] + 1
		} else {
			pref[i] = 2
		}
	}

	suff := make([]int, n)
	suff[n-1] = 1
	suff[n-2] = 2
	for i := n - 3; i >= 0; i-- {
		if nums[i+2]-nums[i+1] == nums[i+1]-nums[i] {
			suff[i] = suff[i+1] + 1
		} else {
			suff[i] = 2
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		if pref[i] > ans {
			ans = pref[i]
		}
	}

	if ans < n {
		ans = max(ans, 1+suff[1])
	}
	if ans < n {
		ans = max(ans, 1+pref[n-2])
	}

	for i := 1; i < n-1; i++ {
		if (nums[i+1]-nums[i-1])%2 != 0 {
			continue
		}
		d := (nums[i+1] - nums[i-1]) / 2

		leftLen := 1
		if i >= 2 && nums[i-1]-nums[i-2] == d {
			leftLen = pref[i-1]
		}

		rightLen := 1
		if i <= n-3 && nums[i+2]-nums[i+1] == d {
			rightLen = suff[i+1]
		}

		total := leftLen + 1 + rightLen
		if total > ans {
			ans = total
		}
	}

	return ans
}

func main() {
	// Example 1
	fmt.Println(LongestArithmeticSequenceAfterChangingAtMostOneElement([]int{9, 7, 5, 10, 1})) // Expected: 5

	// Example 2
	fmt.Println(LongestArithmeticSequenceAfterChangingAtMostOneElement([]int{1, 2, 6, 7})) // Expected: 3

	// Extra
	fmt.Println(LongestArithmeticSequenceAfterChangingAtMostOneElement([]int{1, 2, 3, 4})) // Expected: 4
	fmt.Println(LongestArithmeticSequenceAfterChangingAtMostOneElement([]int{1}))          // Expected: 1
}
```

## 3874 — Valid Subarrays With Exactly One Peak

```go
package main

// LeetCode #3874: Valid Subarrays With Exactly One Peak
// https://leetcode.com/problems/valid-subarrays-with-exactly-one-peak/
// Difficulty: Medium [Paid]
// Time: O(N) | Space: O(N)
// Approach: Find all peaks. For each peak, count valid subarrays that contain
// exactly this peak, bounded by adjacent peaks and distance k from the peak.

import "fmt"

func ValidSubarraysWithExactlyOnePeak(nums []int, k int) int {
	n := len(nums)

	// Find peak indices
	peaks := []int{}
	for i := 1; i < n-1; i++ {
		if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
			peaks = append(peaks, i)
		}
	}

	if len(peaks) == 0 {
		return 0
	}

	ans := 0
	for idx, p := range peaks {
		// Left bound: can't go beyond k steps from peak, and can't include previous peak
		leftBound := p - k
		if idx > 0 && peaks[idx-1] >= leftBound {
			leftBound = peaks[idx-1] + 1
		}
		if leftBound < 0 {
			leftBound = 0
		}

		// Right bound: can't go beyond k steps from peak, and can't include next peak
		rightBound := p + k
		if idx < len(peaks)-1 && peaks[idx+1] <= rightBound {
			rightBound = peaks[idx+1] - 1
		}
		if rightBound >= n {
			rightBound = n - 1
		}

		leftOptions := p - leftBound + 1
		rightOptions := rightBound - p + 1
		ans += leftOptions * rightOptions
	}

	return ans
}

func main() {
	// Example 1
	fmt.Println(ValidSubarraysWithExactlyOnePeak([]int{1, 3, 2}, 1)) // Expected: 4

	// Example 2
	fmt.Println(ValidSubarraysWithExactlyOnePeak([]int{7, 8, 9}, 2)) // Expected: 0

	// Example 3
	fmt.Println(ValidSubarraysWithExactlyOnePeak([]int{4, 3, 5, 1}, 2)) // Expected: 6
}
```

## 3876 — Construct Uniform Parity Array Ii

```go
package main

// LeetCode #3876: Construct Uniform Parity Array II
// https://leetcode.com/problems/construct-uniform-parity-array-ii/
// Difficulty: Medium
// Time: O(N) | Space: O(1)
// Approach: If all same parity, return true. Else find smallest odd number.
// If any even number smaller than smallest odd, impossible (can't flip parity).

import "fmt"

func ConstructUniformParityArrayIi(nums1 []int) bool {
	hasEven, hasOdd := false, false
	minOdd := int(1e9 + 1)
	for _, v := range nums1 {
		if v%2 == 0 {
			hasEven = true
		} else {
			hasOdd = true
			if v < minOdd {
				minOdd = v
			}
		}
	}
	if !hasEven || !hasOdd {
		return true
	}
	// Both parities present. Check if any even < smallest odd.
	for _, v := range nums1 {
		if v%2 == 0 && v < minOdd {
			return false
		}
	}
	return true
}

func main() {
	// Example 1
	fmt.Println(ConstructUniformParityArrayIi([]int{1, 4, 7})) // Expected: true

	// Example 2
	fmt.Println(ConstructUniformParityArrayIi([]int{2, 3})) // Expected: false

	// Example 3
	fmt.Println(ConstructUniformParityArrayIi([]int{4, 6})) // Expected: true
}
```

## 3877 — Minimum Removals To Achieve Target Xor

```go
package main

// LeetCode #3877: Minimum Removals to Achieve Target XOR
// https://leetcode.com/problems/minimum-removals-to-achieve-target-xor/
// Difficulty: Medium
// Time: O(N * 2^M) | Space: O(2^M) where M = max bit length (14)
// Approach: DP tracking max selectable elements to achieve each XOR value.
// Answer = len(nums) - maxElementsForTarget (or -1 if unreachable).

import "fmt"

func MinimumRemovalsToAchieveTargetXor(nums []int, target int) int {
	maxXor := 1
	for _, v := range nums {
		for maxXor <= v {
			maxXor <<= 1
		}
	}
	if maxXor <= target {
		for maxXor <= target {
			maxXor <<= 1
		}
	}

	// dp[x] = max elements selectable to achieve XOR x
	dp := make([]int, maxXor)
	for i := range dp {
		dp[i] = -1
	}
	dp[0] = 0

	for _, v := range nums {
		ndp := make([]int, maxXor)
		copy(ndp, dp)
		for x := 0; x < maxXor; x++ {
			if dp[x] >= 0 {
				nx := x ^ v
				if dp[x]+1 > ndp[nx] {
					ndp[nx] = dp[x] + 1
				}
			}
		}
		dp = ndp
	}

	if dp[target] < 0 {
		return -1
	}
	return len(nums) - dp[target]
}

func main() {
	// Example 1
	fmt.Println(MinimumRemovalsToAchieveTargetXor([]int{1, 2, 3}, 2)) // Expected: 1

	// Example 2
	fmt.Println(MinimumRemovalsToAchieveTargetXor([]int{2, 4}, 1)) // Expected: -1

	// Example 3
	fmt.Println(MinimumRemovalsToAchieveTargetXor([]int{7}, 7)) // Expected: 0
}
```

## 3879 — Maximum Distinct Path Sum In A Binary Tree

```go
package main

// LeetCode #3879: Maximum Distinct Path Sum in a Binary Tree
// https://leetcode.com/problems/maximum-distinct-path-sum-in-a-binary-tree/
// Difficulty: Medium [Paid]
// Time: O(N^2) | Space: O(N)
// Approach: DFS from each node as start, exploring all paths with distinct values.

import (
	"fmt"
	"math"
)

// TreeNode definition for binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDistinctPathSum(root *TreeNode, visited map[int]bool) int {
	if root == nil || visited[root.Val] {
		return 0
	}

	visited[root.Val] = true
	defer func() { delete(visited, root.Val) }()

	leftSum := maxDistinctPathSum(root.Left, visited)
	rightSum := maxDistinctPathSum(root.Right, visited)

	return root.Val + max(leftSum, rightSum)
}

func MaximumDistinctPathSumInABinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	ans := math.MinInt32

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		visited := make(map[int]bool)
		sum := maxDistinctPathSum(node, visited)
		if sum > ans {
			ans = sum
		}
		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)
	return ans
}

func main() {
	// Example 1: root = [2,2,1]
	root1 := &TreeNode{Val: 2,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 1},
	}
	fmt.Println(MaximumDistinctPathSumInABinaryTree(root1)) // Expected: 3

	// Example 2: root = [1,-2,5,null,null,3,5]
	root2 := &TreeNode{Val: 1,
		Left:  &TreeNode{Val: -2},
		Right: &TreeNode{Val: 5,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 5},
		},
	}
	fmt.Println(MaximumDistinctPathSumInABinaryTree(root2)) // Expected: 9

	// Example 3: root = [4,6,6,null,null,null,9]
	root3 := &TreeNode{Val: 4,
		Left:  &TreeNode{Val: 6},
		Right: &TreeNode{Val: 6,
			Right: &TreeNode{Val: 9},
		},
	}
	fmt.Println(MaximumDistinctPathSumInABinaryTree(root3)) // Expected: 19
}
```

## 3881 — Direction Assignments With Exactly K Visible People

```go
package main

// LeetCode #3881: Direction Assignments with Exactly K Visible People
// https://leetcode.com/problems/direction-assignments-with-exactly-k-visible-people/
// Difficulty: Medium
// Time: O(N) | Space: O(N)
// Approach: Combinatorics. For each possible a (visible from left),
// b = k - a must be visible from right. Sum 2 * C(pos, a) * C(n-pos-1, b).

import "fmt"

const MOD int64 = 1000000007

func modPow(a int64, b int, mod int64) int64 {
	var res int64 = 1
	for b > 0 {
		if b&1 == 1 {
			res = res * a % mod
		}
		a = a * a % mod
		b >>= 1
	}
	return res
}

func DirectionAssignmentsWithExactlyKVisiblePeople(n int, pos int, k int) int {
	left := pos
	right := n - pos - 1

	fact := make([]int64, n+1)
	fact[0] = 1
	for i := 1; i <= n; i++ {
		fact[i] = fact[i-1] * int64(i) % MOD
	}

	invFact := make([]int64, n+1)
	invFact[n] = modPow(fact[n], int(MOD-2), MOD)
	for i := n - 1; i >= 0; i-- {
		invFact[i] = invFact[i+1] * int64(i+1) % MOD
	}

	nCr := func(nn, rr int) int64 {
		if rr < 0 || rr > nn {
			return 0
		}
		return fact[nn] * invFact[rr] % MOD * invFact[nn-rr] % MOD
	}

	var ans int64 = 0
	for a := 0; a <= left; a++ {
		b := k - a
		if b < 0 || b > right {
			continue
		}
		ans = (ans + nCr(left, a)*nCr(right, b)%MOD) % MOD
	}

	// Person at pos can be L or R
	ans = ans * 2 % MOD
	return int(ans)
}

func main() {
	// Example 1
	fmt.Println(DirectionAssignmentsWithExactlyKVisiblePeople(3, 1, 0)) // Expected: 2

	// Example 2
	fmt.Println(DirectionAssignmentsWithExactlyKVisiblePeople(3, 2, 1)) // Expected: 4

	// Example 3
	fmt.Println(DirectionAssignmentsWithExactlyKVisiblePeople(1, 0, 0)) // Expected: 2
}
```

## 3882 — Minimum Xor Path In A Grid

```go
package main

// LeetCode #3882: Minimum XOR Path in a Grid
// https://leetcode.com/problems/minimum-xor-path-in-a-grid/
// Difficulty: Medium
// Time: O(M * N * 2^B) | Space: O(N * 2^B) where B = 11 (grid values < 1024)
// Approach: DP tracking reachable XOR values at each cell. Only right/down moves.

import (
	"fmt"
	"math"
)

func MinimumXorPathInAGrid(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return -1
	}
	n := len(grid[0])
	if n == 0 {
		return -1
	}

	maxXor := 2048 // 2^11 since grid[i][j] <= 1023

	// Use bitset (boolean array) for each cell
	dp := make([][]bool, n)
	for j := 0; j < n; j++ {
		dp[j] = make([]bool, maxXor)
	}

	dp[0][grid[0][0]] = true

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			cur := make([]bool, maxXor)
			if i > 0 {
				for x := 0; x < maxXor; x++ {
					if dp[j][x] {
						cur[x^grid[i][j]] = true
					}
				}
			}
			if j > 0 {
				for x := 0; x < maxXor; x++ {
					if dp[j-1][x] {
						cur[x^grid[i][j]] = true
					}
				}
			}
			dp[j] = cur
		}
	}

	ans := math.MaxInt32
	for x := 0; x < maxXor; x++ {
		if dp[n-1][x] && x < ans {
			ans = x
		}
	}
	return ans
}

func main() {
	// Example 1
	fmt.Println(MinimumXorPathInAGrid([][]int{{1, 2}, {3, 4}})) // Expected: 6

	// Example 2
	fmt.Println(MinimumXorPathInAGrid([][]int{{6, 7}, {5, 8}})) // Expected: 9

	// Example 3
	fmt.Println(MinimumXorPathInAGrid([][]int{{2, 7, 5}})) // Expected: 0
}
```

## 3885 — Design Event Manager

```go
package main

// LeetCode #3885: Design Event Manager
// https://leetcode.com/problems/design-event-manager/
// Difficulty: Medium
// Time: O(N log N) init, O(log N) per operation | Space: O(N)
// Approach: Use max-heap (priority queue) on (-priority, eventId) + map for event priorities.

import (
	"container/heap"
	"fmt"
)

type Event struct {
	priority int
	eventId  int
}

type MaxHeap []Event

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool {
	if h[i].priority != h[j].priority {
		return h[i].priority > h[j].priority
	}
	return h[i].eventId < h[j].eventId
}
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(Event)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type EventManager struct {
	pq       *MaxHeap
	priorities map[int]int // eventId -> priority
}

func Constructor(events [][]int) EventManager {
	pq := &MaxHeap{}
	heap.Init(pq)
	pm := make(map[int]int)
	for _, e := range events {
		id, pri := e[0], e[1]
		pm[id] = pri
		heap.Push(pq, Event{priority: pri, eventId: id})
	}
	return EventManager{pq: pq, priorities: pm}
}

func (em *EventManager) UpdatePriority(eventId int, newPriority int) {
	em.priorities[eventId] = newPriority
	heap.Push(em.pq, Event{priority: newPriority, eventId: eventId})
}

func (em *EventManager) PollHighest() int {
	for em.pq.Len() > 0 {
		top := (*em.pq)[0]
		if pri, ok := em.priorities[top.eventId]; ok && pri == top.priority {
			heap.Pop(em.pq)
			delete(em.priorities, top.eventId)
			return top.eventId
		}
		heap.Pop(em.pq) // stale entry
	}
	return -1
}

func main() {
	// Example 1
	em := Constructor([][]int{{5, 7}, {2, 7}, {9, 4}})
	fmt.Println(em.PollHighest()) // Expected: 2
	em.UpdatePriority(9, 7)
	fmt.Println(em.PollHighest()) // Expected: 5
	fmt.Println(em.PollHighest()) // Expected: 9

	// Example 2
	em2 := Constructor([][]int{{4, 1}, {7, 2}})
	fmt.Println(em2.PollHighest()) // Expected: 7
	fmt.Println(em2.PollHighest()) // Expected: 4
	fmt.Println(em2.PollHighest()) // Expected: -1
}
```

## 3889 — Mirror Frequency Distance

```go
package main

// LeetCode #3889: Mirror Frequency Distance
// https://leetcode.com/problems/mirror-frequency-distance/
// Difficulty: Medium
// Time: O(N) | Space: O(1)
// Approach: Count character frequencies. For each unique char, compute mirror
// char and sum absolute frequency differences, counting each pair once.

import "fmt"

func MirrorFrequencyDistance(s string) int {
	freq := make([]int, 36) // 0-25: letters, 26-35: digits
	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			freq[ch-'a']++
		} else {
			freq[26+int(ch-'0')]++
		}
	}

	visited := make([]bool, 36)
	ans := 0
	for i := 0; i < 36; i++ {
		if visited[i] || freq[i] == 0 {
			continue
		}
		// Compute mirror
		var mirror int
		if i < 26 {
			mirror = 25 - i // 'a'->'z', 'b'->'y', etc.
		} else {
			mirror = 26 + (9 - (i - 26)) // '0'->'9', '1'->'8', etc.
		}
		visited[i] = true
		if mirror != i {
			visited[mirror] = true
		}
		diff := freq[i] - freq[mirror]
		if diff < 0 {
			diff = -diff
		}
		ans += diff
	}
	return ans
}

func main() {
	// Example 1
	fmt.Println(MirrorFrequencyDistance("ab1z9")) // Expected: 3

	// Example 2
	fmt.Println(MirrorFrequencyDistance("4m7n")) // Expected: 2

	// Example 3
	fmt.Println(MirrorFrequencyDistance("byby")) // Expected: 0
}
```

## 3890 — Integers With Multiple Sum Of Two Cubes

```go
package main

// LeetCode #3890: Integers With Multiple Sum of Two Cubes
// https://leetcode.com/problems/integers-with-multiple-sum-of-two-cubes/
// Difficulty: Medium
// Time: O(C^2) precompute + O(N log N) sort | Space: O(C^2)
// Approach: Enumerate a,b up to 1000 (since 1000^3 = 1e9). Count frequency
// of each sum. Return sums with >= 2 representations, sorted.

import (
	"fmt"
	"sort"
)

func IntegersWithMultipleSumOfTwoCubes(n int) []int {
	cubeCount := make(map[int]int)
	limit := 1000
	for a := 1; a <= limit; a++ {
		a3 := a * a * a
		if a3 > n {
			break
		}
		for b := a; b <= limit; b++ {
			b3 := b * b * b
			sum := a3 + b3
			if sum > n {
				break
			}
			cubeCount[sum]++
		}
	}

	ans := []int{}
	for sum, cnt := range cubeCount {
		if cnt >= 2 {
			ans = append(ans, sum)
		}
	}
	sort.Ints(ans)
	return ans
}

func main() {
	// Example 1
	fmt.Println(IntegersWithMultipleSumOfTwoCubes(4104)) // Expected: [1729 4104]

	// Example 2
	fmt.Println(IntegersWithMultipleSumOfTwoCubes(578)) // Expected: []

	// Extra
	fmt.Println(IntegersWithMultipleSumOfTwoCubes(1729)) // Expected: [1729]
}
```

## 3891 — Minimum Increase To Maximize Special Indices

```go
package main

// LeetCode #3891: Minimum Increase to Maximize Special Indices
// https://leetcode.com/problems/minimum-increase-to-maximize-special-indices/
// Difficulty: Medium
// Time: O(N) | Space: O(N)
// Approach: DP with memoization. Max special indices = ceil((n-2)/2) = (n-1)/2.
// Find min cost by considering each position as peak or not.

import (
	"fmt"
	"math"
)

func MinimumIncreaseToMaximizeSpecialIndices(nums []int) int {
	n := len(nums)
	if n < 3 {
		return 0
	}

	// cost to make each position a peak
	cost := make([]int, n)
	for i := 1; i < n-1; i++ {
		need := max(nums[i-1], nums[i+1]) + 1
		if nums[i] < need {
			cost[i] = need - nums[i]
		}
	}

	maxPeaks := (n - 1) / 2
	memo := make([][]int, n+1)
	for i := range memo {
		memo[i] = make([]int, maxPeaks+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	// dp(pos, peaksSelected) = min cost from pos to n-2 (pos starts at 1)
	var dp func(pos int, left int) int
	dp = func(pos int, left int) int {
		if left == 0 {
			return 0
		}
		if pos >= n-1 {
			if left > 0 {
				return math.MaxInt32
			}
			return 0
		}
		if memo[pos][left] != -1 {
			return memo[pos][left]
		}

		// Skip position pos
		best := dp(pos+1, left)

		// Make pos a peak (skip pos+1 since adjacent can't be peak)
		best = min(best, cost[pos]+dp(pos+2, left-1))

		memo[pos][left] = best
		return best
	}

	ans := dp(1, maxPeaks)
	return ans
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
	// Example 1
	fmt.Println(MinimumIncreaseToMaximizeSpecialIndices([]int{1, 2, 2})) // Expected: 1

	// Example 2
	fmt.Println(MinimumIncreaseToMaximizeSpecialIndices([]int{2, 1, 1, 3})) // Expected: 2

	// Example 3
	fmt.Println(MinimumIncreaseToMaximizeSpecialIndices([]int{5, 2, 1, 4, 3})) // Expected: 4
}
```

## 3893 — Maximum Team Size With Overlapping Intervals

```go
package main

// LeetCode #3893: Maximum Team Size with Overlapping Intervals
// https://leetcode.com/problems/maximum-team-size-with-overlapping-intervals/
// Difficulty: Medium [Paid]
// Time: O(N log N) | Space: O(N)
// Approach: For each employee, count overlapping intervals using binary search
// on sorted start and end times.

import (
	"fmt"
	"sort"
)

func MaximumTeamSizeWithOverlappingIntervals(startTime []int, endTime []int) int {
	n := len(startTime)
	st := make([]int, n)
	et := make([]int, n)
	copy(st, startTime)
	copy(et, endTime)
	sort.Ints(st)
	sort.Ints(et)

	ans := 0
	for i := 0; i < n; i++ {
		start := startTime[i]
		end := endTime[i]

		// Count employees whose start <= end
		startsBeforeEnd := sort.SearchInts(st, end+1)
		// Count employees whose end < start
		endsBeforeStart := sort.SearchInts(et, start)

		overlap := startsBeforeEnd - endsBeforeStart
		if overlap > ans {
			ans = overlap
		}
	}
	return ans
}

func main() {
	// Example 1
	fmt.Println(MaximumTeamSizeWithOverlappingIntervals([]int{1, 2, 3}, []int{4, 5, 6})) // Expected: 3

	// Example 2
	fmt.Println(MaximumTeamSizeWithOverlappingIntervals([]int{2, 5, 8}, []int{3, 7, 9})) // Expected: 1

	// Example 3
	fmt.Println(MaximumTeamSizeWithOverlappingIntervals([]int{3, 4, 6}, []int{8, 5, 7})) // Expected: 3
}
```

## 3895 — Count Digit Appearances

```go
package main

// LeetCode #3895: Count Digit Appearances
// https://leetcode.com/problems/count-digit-appearances/
// Difficulty: Medium
// Time: O(N * log M) | Space: O(1) where M = max value in nums
// Approach: For each number, extract digits and count matches to target digit.

import "fmt"

func CountDigitAppearances(nums []int, digit int) int {
	ans := 0
	for _, v := range nums {
		for v > 0 {
			if v%10 == digit {
				ans++
			}
			v /= 10
		}
	}
	return ans
}

func main() {
	// Example 1
	fmt.Println(CountDigitAppearances([]int{12, 54, 32, 22}, 2)) // Expected: 4

	// Example 2
	fmt.Println(CountDigitAppearances([]int{1, 34, 7}, 9)) // Expected: 0
}
```

## 3896 — Minimum Operations To Transform Array Into Alternating Prime

```go
package main

// LeetCode #3896: Minimum Operations to Transform Array into Alternating Prime
// https://leetcode.com/problems/minimum-operations-to-transform-array-into-alternating-prime/
// Difficulty: Medium
// Time: O(N log log M + N) | Space: O(M) where M = 200000
// Approach: Sieve primes. Even indices need next prime >= val (binary search).
// Odd indices need non-prime: if prime, inc to next non-prime (2->4 cost 2, rest cost 1).

import "fmt"

const MAX_VAL = 200000

func MinimumOperationsToTransformArrayIntoAlternatingPrime(nums []int) int64 {
	isPrime := make([]bool, MAX_VAL+1)
	for i := 2; i <= MAX_VAL; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i <= MAX_VAL; i++ {
		if isPrime[i] {
			for j := i * i; j <= MAX_VAL; j += i {
				isPrime[j] = false
			}
		}
	}

	// nextPrime[v] = smallest prime >= v
	nextPrime := make([]int, MAX_VAL+1)
	np := -1
	for i := MAX_VAL; i >= 2; i-- {
		if isPrime[i] {
			np = i
		}
		nextPrime[i] = np
	}

	var ans int64 = 0
	for i, v := range nums {
		if i%2 == 0 {
			// Even index: must be prime
			need := nextPrime[v]
			if need == -1 {
				// No prime >= v within MAX_VAL (shouldn't happen per constraints)
				ans += int64(MAX_VAL + 1 - v)
			} else {
				ans += int64(need - v)
			}
		} else {
			// Odd index: must be non-prime
			if isPrime[v] {
				if v == 2 {
					ans += 2 // 2 -> 4
				} else {
					ans += 1 // odd prime +1 -> even non-prime
				}
			}
		}
	}
	return ans
}

func main() {
	// Example 1
	fmt.Println(MinimumOperationsToTransformArrayIntoAlternatingPrime([]int{1, 2, 3, 4})) // Expected: 3

	// Example 2
	fmt.Println(MinimumOperationsToTransformArrayIntoAlternatingPrime([]int{5, 6, 7, 8})) // Expected: 0

	// Example 3
	fmt.Println(MinimumOperationsToTransformArrayIntoAlternatingPrime([]int{4, 4})) // Expected: 1
}
```

## 3899 — Angles Of A Triangle

```go
package main

// LeetCode #3899: Angles of a Triangle
// https://leetcode.com/problems/angles-of-a-triangle/
// Difficulty: Medium
// Time: O(1) | Space: O(1)
// Approach: Use law of cosines. Check triangle inequality first.

import (
	"fmt"
	"math"
	"sort"
)

func AnglesOfATriangle(sides []int) []float64 {
	a := float64(sides[0])
	b := float64(sides[1])
	c := float64(sides[2])
	s := []float64{a, b, c}
	sort.Float64s(s)
	a, b, c = s[0], s[1], s[2]

	// Triangle inequality
	if a+b <= c {
		return []float64{}
	}

	// Law of cosines: cos(A) = (b^2 + c^2 - a^2) / (2*b*c)
	angleA := math.Acos((b*b + c*c - a*a) / (2 * b * c)) * 180 / math.Pi
	angleB := math.Acos((a*a + c*c - b*b) / (2 * a * c)) * 180 / math.Pi
	angleC := 180 - angleA - angleB

	angles := []float64{math.Round(angleA*1e5) / 1e5, math.Round(angleB*1e5) / 1e5, math.Round(angleC*1e5) / 1e5}
	sort.Float64s(angles)
	return angles
}

func main() {
	// Example 1
	fmt.Println(AnglesOfATriangle([]int{3, 4, 5})) // Expected: [36.86990 53.13010 90.00000]

	// Example 2
	fmt.Println(AnglesOfATriangle([]int{2, 4, 2})) // Expected: []

	// Extra
	fmt.Println(AnglesOfATriangle([]int{1, 1, 1})) // Expected: ~[60 60 60]
}
```

## 3900 — Longest Balanced Substring After One Swap

```go
package main

// LeetCode #3900: Longest Balanced Substring After One Swap
// https://leetcode.com/problems/longest-balanced-substring-after-one-swap/
// Difficulty: Medium
// Time: O(N) | Space: O(N)
// Approach: Prefix sum (0->-1, 1->+1). Track first occurrence of each prefix sum.
// Check offsets 0 (no swap) and +/-2 (one swap possible).

import "fmt"

func LongestBalancedSubstringAfterOneSwap(s string) int {
	n := len(s)

	pref := make([]int, n+1)
	onesPref := make([]int, n+1) // prefix count of '1's
	for i := 0; i < n; i++ {
		onesPref[i+1] = onesPref[i]
		if s[i] == '1' {
			pref[i+1] = pref[i] + 1
			onesPref[i+1]++
		} else {
			pref[i+1] = pref[i] - 1
		}
	}

	totalOnes := onesPref[n]
	totalZeros := n - totalOnes

	firstPos := make(map[int]int)
	firstPos[0] = 0

	ans := 0
	for i := 1; i <= n; i++ {
		// No swap needed (offset 0)
		if pos, ok := firstPos[pref[i]]; ok {
			if i-pos > ans {
				ans = i - pos
			}
		} else {
			firstPos[pref[i]] = i
		}

		// Swap '1' in with '0' out (offset +2, sum too many 1s)
		if pos, ok := firstPos[pref[i]-2]; ok {
			onesInside := onesPref[i] - onesPref[pos]
			zerosInside := (i - pos) - onesInside
			zerosOutside := totalZeros - zerosInside
			if onesInside > 0 && zerosOutside > 0 {
				if i-pos > ans {
					ans = i - pos
				}
			}
		}

		// Swap '0' in with '1' out (offset -2, sum too many 0s)
		if pos, ok := firstPos[pref[i]+2]; ok {
			onesInside := onesPref[i] - onesPref[pos]
			zerosInside := (i - pos) - onesInside
			onesOutside := totalOnes - onesInside
			if zerosInside > 0 && onesOutside > 0 {
				if i-pos > ans {
					ans = i - pos
				}
			}
		}
	}

	return ans
}

func main() {
	// Example 1
	fmt.Println(LongestBalancedSubstringAfterOneSwap("100001")) // Expected: 4

	// Example 2
	fmt.Println(LongestBalancedSubstringAfterOneSwap("111")) // Expected: 0

	// Extra
	fmt.Println(LongestBalancedSubstringAfterOneSwap("01"))   // Expected: 2
	fmt.Println(LongestBalancedSubstringAfterOneSwap("1100")) // Expected: 4
}
```

## 3902 — Zigzag Level Sum Of Binary Tree

```go
package main

// LeetCode #3902: Zigzag Level Sum of Binary Tree
// https://leetcode.com/problems/zigzag-level-sum-of-binary-tree/
// Difficulty: Medium [Paid]
// Time: O(N) | Space: O(N)
// Approach: BFS level-order. At odd levels (left-to-right), stop at first
// node without left child. At even levels (right-to-left), stop at first
// node without right child. Collect children from all nodes.

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func ZigzagLevelSumOfBinaryTree(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ans := []int{}
	queue := []*TreeNode{root}
	level := 1

	for len(queue) > 0 {
		nextQ := []*TreeNode{}
		sum := 0

		if level%2 == 1 {
			// Odd: left to right, stop at first node without LEFT child
			for _, node := range queue {
				if node.Left == nil {
					break
				}
				sum += node.Val
			}
		} else {
			// Even: right to left, stop at first node without RIGHT child
			for i := len(queue) - 1; i >= 0; i-- {
				if queue[i].Right == nil {
					break
				}
				sum += queue[i].Val
			}
		}

		// Collect children from ALL nodes at this level
		for _, node := range queue {
			if node.Left != nil {
				nextQ = append(nextQ, node.Left)
			}
			if node.Right != nil {
				nextQ = append(nextQ, node.Right)
			}
		}

		ans = append(ans, sum)
		queue = nextQ
		level++
	}

	return ans
}

func makeTree(vals []interface{}) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}
	root := &TreeNode{Val: vals[0].(int)}
	queue := []*TreeNode{root}
	idx := 1
	for len(queue) > 0 && idx < len(vals) {
		node := queue[0]
		queue = queue[1:]
		if idx < len(vals) && vals[idx] != nil {
			node.Left = &TreeNode{Val: vals[idx].(int)}
			queue = append(queue, node.Left)
		}
		idx++
		if idx < len(vals) && vals[idx] != nil {
			node.Right = &TreeNode{Val: vals[idx].(int)}
			queue = append(queue, node.Right)
		}
		idx++
	}
	return root
}

func main() {
	// Example 1: root = [5,2,8,1,null,9,6]
	root1 := makeTree([]interface{}{5, 2, 8, 1, nil, 9, 6})
	fmt.Println(ZigzagLevelSumOfBinaryTree(root1)) // Expected: [5 8 0]

	// Example 2: root = [1,2,3,4,5,null,7]
	root2 := makeTree([]interface{}{1, 2, 3, 4, 5, nil, 7})
	fmt.Println(ZigzagLevelSumOfBinaryTree(root2)) // Expected: [1 5 0]
}
```

## 3904 — Smallest Stable Index Ii

```go
package main

// LeetCode #3904: Smallest Stable Index II
// https://leetcode.com/problems/smallest-stable-index-ii/
// Difficulty: Medium
// Time: O(N) | Space: O(N)
// Approach: Precompute suffix minimum, iterate prefix maximum.
// Check if max(nums[0..i]) - min(nums[i..n-1]) <= k.

import "fmt"

func SmallestStableIndexIi(nums []int, k int) int {
	n := len(nums)
	suffixMin := make([]int, n)
	suffixMin[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		suffixMin[i] = min(nums[i], suffixMin[i+1])
	}

	prefixMax := 0
	for i := 0; i < n; i++ {
		if nums[i] > prefixMax {
			prefixMax = nums[i]
		}
		if prefixMax-suffixMin[i] <= k {
			return i
		}
	}
	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Example 1
	fmt.Println(SmallestStableIndexIi([]int{5, 0, 1, 4}, 3)) // Expected: 3

	// Example 2
	fmt.Println(SmallestStableIndexIi([]int{3, 2, 1}, 1)) // Expected: -1

	// Example 3
	fmt.Println(SmallestStableIndexIi([]int{0}, 0)) // Expected: 0
}
```

## 3905 — Multi Source Flood Fill

```go
package main

// LeetCode #3905: Multi Source Flood Fill
// https://leetcode.com/problems/multi-source-flood-fill/
// Difficulty: Medium
// Time: O(N*M) | Space: O(N*M)
// Approach: Multi-source BFS. Track time and color for each cell.
// Use max color when multiple colors reach same cell at same time.

import "fmt"

func MultiSourceFloodFill(n int, m int, sources [][]int) [][]int {
	grid := make([][]int, n)
	time := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, m)
		time[i] = make([]int, m)
		for j := 0; j < m; j++ {
			time[i][j] = -1
		}
	}

	type Cell struct{ r, c, t, color int }
	queue := make([]Cell, 0)
	for _, src := range sources {
		r, c, color := src[0], src[1], src[2]
		grid[r][c] = color
		time[r][c] = 1
		queue = append(queue, Cell{r, c, 1, color})
	}

	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	head := 0
	for head < len(queue) {
		cur := queue[head]
		head++
		for _, d := range dirs {
			nr, nc := cur.r+d[0], cur.c+d[1]
			if nr < 0 || nr >= n || nc < 0 || nc >= m {
				continue
			}
			nt := cur.t + 1
			if time[nr][nc] == -1 {
				// Unvisited
				time[nr][nc] = nt
				grid[nr][nc] = cur.color
				queue = append(queue, Cell{nr, nc, nt, cur.color})
			} else if time[nr][nc] == nt && grid[nr][nc] < cur.color {
				// Same time, pick max color
				grid[nr][nc] = cur.color
			}
		}
	}

	return grid
}

func main() {
	// Example 1
	fmt.Println(MultiSourceFloodFill(3, 3, [][]int{{0, 0, 1}, {2, 2, 2}}))
	// Expected: [[1 1 2] [1 2 2] [2 2 2]]

	// Example 2
	fmt.Println(MultiSourceFloodFill(3, 3, [][]int{{0, 1, 3}, {1, 1, 5}}))
	// Expected: [[3 3 3] [5 5 5] [5 5 5]]

	// Example 3
	fmt.Println(MultiSourceFloodFill(2, 2, [][]int{{1, 1, 5}}))
	// Expected: [[5 5] [5 5]]
}
```

## 3907 — Count Smaller Elements With Opposite Parity

```go
package main

// LeetCode #3907: Count Smaller Elements With Opposite Parity
// https://leetcode.com/problems/count-smaller-elements-with-opposite-parity/
// Difficulty: Medium [Paid]
// Time: O(N log M) | Space: O(M) where M = max value
// Approach: Process right to left. Use two BITs (even, odd) to count smaller
// elements with opposite parity.

import (
	"fmt"
	"sort"
)

type BIT struct {
	tree []int
}

func NewBIT(size int) *BIT {
	return &BIT{tree: make([]int, size+2)}
}

func (b *BIT) Update(idx int, val int) {
	idx++
	for idx < len(b.tree) {
		b.tree[idx] += val
		idx += idx & -idx
	}
}

func (b *BIT) Query(idx int) int {
	idx++
	sum := 0
	for idx > 0 {
		sum += b.tree[idx]
		idx -= idx & -idx
	}
	return sum
}

func CountSmallerElementsWithOppositeParity(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)

	// Coordinate compress
	sorted := make([]int, n)
	copy(sorted, nums)
	sort.Ints(sorted)
	rank := make(map[int]int)
	for i, v := range sorted {
		rank[v] = i
	}

	evenBit := NewBIT(n)
	oddBit := NewBIT(n)

	for i := n - 1; i >= 0; i-- {
		r := rank[nums[i]]
		if nums[i]%2 == 0 {
			// Count odd elements smaller than nums[i]
			ans[i] = oddBit.Query(r - 1)
			evenBit.Update(r, 1)
		} else {
			// Count even elements smaller than nums[i]
			ans[i] = evenBit.Query(r - 1)
			oddBit.Update(r, 1)
		}
	}

	return ans
}

func main() {
	// Example 1
	fmt.Println(CountSmallerElementsWithOppositeParity([]int{5, 2, 4, 1, 3})) // Expected: [2 1 2 0 0]

	// Example 2
	fmt.Println(CountSmallerElementsWithOppositeParity([]int{4, 4, 1})) // Expected: [1 1 0]

	// Example 3
	fmt.Println(CountSmallerElementsWithOppositeParity([]int{7})) // Expected: [0]
}
```

## 3909 — Compare Sums Of Bitonic Parts

```go
package main

// LeetCode #3909: Compare Sums of Bitonic Parts
// https://leetcode.com/problems/compare-sums-of-bitonic-parts/
// Difficulty: Medium
// Time: O(N) | Space: O(1)
// Approach: Find peak where nums[i] > nums[i+1]. Sum left part from 0 to peak,
// sum right part from peak to n-1. Return -1/0/1.

import "fmt"

func CompareSumsOfBitonicParts(nums []int) int {
	n := len(nums)
	peak := 0
	for i := 1; i < n-1; i++ {
		if nums[i] > nums[i+1] {
			peak = i
			break
		}
	}

	leftSum := 0
	for i := 0; i <= peak; i++ {
		leftSum += nums[i]
	}
	rightSum := 0
	for i := peak; i < n; i++ {
		rightSum += nums[i]
	}

	if leftSum > rightSum {
		return 0
	} else if rightSum > leftSum {
		return 1
	}
	return -1
}

func main() {
	// Example 1
	fmt.Println(CompareSumsOfBitonicParts([]int{1, 3, 2, 1})) // Expected: 1

	// Example 2
	fmt.Println(CompareSumsOfBitonicParts([]int{2, 4, 5, 2})) // Expected: 0

	// Example 3
	fmt.Println(CompareSumsOfBitonicParts([]int{1, 2, 4, 3})) // Expected: -1
}
```

## 3913 — Sort Vowels By Frequency

```go
package main

// LeetCode #3913: Sort Vowels by Frequency
// https://leetcode.com/problems/sort-vowels-by-frequency/
// Difficulty: Medium
// Time: O(N) | Space: O(1)
// Approach: Count vowel frequencies, sort vowels by freq desc (tie: alpha),
// rebuild string by placing sorted vowels in vowel positions.

import (
	"fmt"
	"sort"
)

func isVowel(ch byte) bool {
	return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u'
}

func SortVowelsByFrequency(s string) string {
	n := len(s)
	freq := make(map[byte]int)
	for i := 0; i < n; i++ {
		if isVowel(s[i]) {
			freq[s[i]]++
		}
	}

	vowels := make([]byte, 0, len(freq))
	for v := range freq {
		vowels = append(vowels, v)
	}
	sort.Slice(vowels, func(i, j int) bool {
		if freq[vowels[i]] != freq[vowels[j]] {
			return freq[vowels[i]] > freq[vowels[j]]
		}
		return vowels[i] < vowels[j]
	})

	ans := make([]byte, n)
	vi := 0
	for i := 0; i < n; i++ {
		if isVowel(s[i]) {
			ans[i] = vowels[vi]
			freq[vowels[vi]]--
			if freq[vowels[vi]] == 0 {
				vi++
			}
		} else {
			ans[i] = s[i]
		}
	}

	return string(ans)
}

func main() {
	// Example
	fmt.Println(SortVowelsByFrequency("hello")) // "holle" or "hollo"? vowels: e,o. Both freq 1. alpha: e < o.
	// Position 1: e, position 4: o. Result: "holle"
	fmt.Println(SortVowelsByFrequency("leetcode"))
}
```

## 3914 — Minimum Operations To Make Array Non Decreasing

```go
package main

// LeetCode #3914: Minimum Operations to Make Array Non Decreasing
// https://leetcode.com/problems/minimum-operations-to-make-array-non-decreasing/
// Difficulty: Medium
// Time: O(N) | Space: O(1)
// Approach: Greedy. When nums[i] < nums[i-1], need to increase a suffix.
// Track operations with running max.

import "fmt"

func MinimumOperationsToMakeArrayNonDecreasing(nums []int) int64 {
	var ans int64 = 0
	mx := 0
	for _, v := range nums {
		if v >= mx {
			mx = v
		} else {
			ans += int64(mx - v)
		}
	}
	return ans
}

func main() {
	// Example
	fmt.Println(MinimumOperationsToMakeArrayNonDecreasing([]int{1, 2, 3}))        // Expected: 0
	fmt.Println(MinimumOperationsToMakeArrayNonDecreasing([]int{3, 2, 1}))        // Expected: 3
	fmt.Println(MinimumOperationsToMakeArrayNonDecreasing([]int{1, 2, 1, 2, 1})) // Expected: 2
}
```

## 3918 — Sum Of Primes Between Number And Its Reverse

```go
package main

// LeetCode #3918: Sum of Primes Between Number and Its Reverse
// https://leetcode.com/problems/sum-of-primes-between-number-and-its-reverse/
// Difficulty: Medium
// Time: O(N log log N) | Space: O(N) where N = max(n, rev(n)) <= 1000
// Approach: Sieve primes up to hi = max(n, rev(n)), sum primes in [lo, hi].

import "fmt"

func SumOfPrimesBetweenNumberAndItsReverse(n int) int64 {
	rev := 0
	for tmp := n; tmp > 0; tmp /= 10 {
		rev = rev*10 + tmp%10
	}

	lo, hi := n, rev
	if lo > hi {
		lo, hi = hi, lo
	}

	isPrime := make([]bool, hi+1)
	for i := 2; i <= hi; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i <= hi; i++ {
		if isPrime[i] {
			for j := i * i; j <= hi; j += i {
				isPrime[j] = false
			}
		}
	}

	var sum int64 = 0
	for i := lo; i <= hi; i++ {
		if isPrime[i] {
			sum += int64(i)
		}
	}
	return sum
}

func main() {
	// Example 1
	fmt.Println(SumOfPrimesBetweenNumberAndItsReverse(13)) // Expected: 132

	// Example 2
	fmt.Println(SumOfPrimesBetweenNumberAndItsReverse(10)) // Expected: 17

	// Example 3
	fmt.Println(SumOfPrimesBetweenNumberAndItsReverse(8)) // Expected: 0
}
```

## 3919 — Minimum Cost To Move Between Indices

```go
package main

// LeetCode #3919: Minimum Cost to Move Between Indices
// https://leetcode.com/problems/minimum-cost-to-move-between-indices/
// Difficulty: Medium
// Time: O(N + Q) | Space: O(N)
// Approach: Build directed cost graph. From x to closest(x) costs 1, else
// abs(nums[x]-nums[y]). Since nums sorted, cheapest path is step-by-step
// using cost-1 edges when available. Precompute prefix sums for O(1) queries.

import "fmt"

func MinimumCostToMoveBetweenIndices(nums []int, queries [][]int) []int64 {
	n := len(nums)
	if n <= 1 {
		ans := make([]int64, len(queries))
		return ans
	}

	// costLR[i] = min cost from i to i+1
	// costRL[i] = min cost from i+1 to i
	costLR := make([]int64, n-1)
	costRL := make([]int64, n-1)

	for i := 0; i < n; i++ {
		if i == 0 {
			// closest(0) = 1
			costLR[0] = 1
		} else if i == n-1 {
			// closest(n-1) = n-2
			costRL[n-2] = 1
		} else {
			leftDiff := nums[i] - nums[i-1]
			rightDiff := nums[i+1] - nums[i]
			if leftDiff <= rightDiff {
				// closest(i) = i-1
				costRL[i-1] = 1
				costLR[i] = int64(rightDiff)
			} else {
				// closest(i) = i+1
				costLR[i] = 1
				costRL[i-1] = int64(leftDiff)
			}
		}
	}

	// prefLR[k] = cost from 0 to k (going right)
	prefLR := make([]int64, n)
	for i := 0; i < n-1; i++ {
		prefLR[i+1] = prefLR[i] + costLR[i]
	}

	// prefRL[k] = cost from k to 0 (going left)
	prefRL := make([]int64, n)
	for i := 0; i < n-1; i++ {
		prefRL[i+1] = prefRL[i] + costRL[i]
	}

	ans := make([]int64, len(queries))
	for i, q := range queries {
		l, r := q[0], q[1]
		if l < r {
			ans[i] = prefLR[r] - prefLR[l]
		} else {
			ans[i] = prefRL[l] - prefRL[r]
		}
	}
	return ans
}

func main() {
	// Example 1
	fmt.Println(MinimumCostToMoveBetweenIndices(
		[]int{-5, -2, 3},
		[][]int{{0, 2}, {2, 0}, {1, 2}},
	)) // Expected: [6 2 5]

	// Example 2
	fmt.Println(MinimumCostToMoveBetweenIndices(
		[]int{0, 2, 3, 9},
		[][]int{{3, 0}, {1, 2}, {2, 0}},
	)) // Expected: [4 1 3]

	// Example 3: single element (edge case)
	fmt.Println(MinimumCostToMoveBetweenIndices(
		[]int{5},
		[][]int{{0, 0}},
	)) // Expected: [0]
}
```

## 3922 — Minimum Flips To Make Binary String Coherent

```go
package main

// LeetCode #3922: Minimum Flips to Make Binary String Coherent
// https://leetcode.com/problems/minimum-flips-to-make-binary-string-coherent/
// Difficulty: Medium
// Time: O(N) | Space: O(N)
// Approach: Valid strings avoid "011" and "110" subsequences. Valid patterns:
// all zeros, all ones, exactly one 1 (0*10*), or exactly two 1s with zeros
// between (10+1). Compute min flips to each pattern.

import "fmt"

func MinimumFlipsToMakeBinaryStringCoherent(s string) int {
	n := len(s)

	pref := make([]int, n+1)
	for i := 0; i < n; i++ {
		pref[i+1] = pref[i]
		if s[i] == '1' {
			pref[i+1]++
		}
	}
	totalOnes := pref[n]

	// Cat 1: all zeros
	ans := totalOnes

	// Cat 2: all ones
	if n-totalOnes < ans {
		ans = n - totalOnes
	}

	// Cat 3: exactly one 1 (0*10*)
	for i := 0; i < n; i++ {
		onesBefore := pref[i]
		onesAfter := pref[n] - pref[i+1]
		cost := onesBefore + onesAfter
		if s[i] == '0' {
			cost++
		}
		if cost < ans {
			ans = cost
		}
	}

	// Cat 4: 10+1 (two 1s with zeros in between)
	if n >= 2 {
		cost := 0
		if s[0] == '0' {
			cost++
		}
		if s[n-1] == '0' {
			cost++
		}
		if n > 2 {
			// Middle positions 1..n-2 must be 0
			cost += pref[n-1] - pref[1]
		}
		if cost < ans {
			ans = cost
		}
	}

	return ans
}

func main() {
	// Example 1
	fmt.Println(MinimumFlipsToMakeBinaryStringCoherent("1010")) // Expected: 1

	// Example 2
	fmt.Println(MinimumFlipsToMakeBinaryStringCoherent("0110")) // Expected: 1

	// Example 3
	fmt.Println(MinimumFlipsToMakeBinaryStringCoherent("1000")) // Expected: 0
}
```

## 3923 — Minimum Generations To Target Point

```go
package main

// LeetCode #3923: Minimum Generations to Target Point
// https://leetcode.com/problems/minimum-generations-to-target-point/
// Difficulty: Medium
// Time: O(K * N^2) | Space: O(7^3) where N = seen points ≤ 343, K = max generations
// Approach: BFS simulation. Generate new points by pairing all distinct seen points,
// computing floor midpoint. Coordinate space is [0,6]^3 (max 343 points).
// Track seen points in 3D boolean grid. Return generation when target appears.

import "fmt"

func MinimumGenerationsToTargetPoint(points [][]int, target []int) int {
	// 7x7x7 grid for coordinate space [0,6]
	var seen [7][7][7]bool
	var curGen [][3]int

	// Initialize with generation 0
	for _, p := range points {
		x, y, z := p[0], p[1], p[2]
		if !seen[x][y][z] {
			seen[x][y][z] = true
			curGen = append(curGen, [3]int{x, y, z})
		}
	}

	tx, ty, tz := target[0], target[1], target[2]
	if seen[tx][ty][tz] {
		return 0
	}

	// BFS generation by generation
	for gen := 1; gen <= 7; gen++ {
		// Collect all points seen so far
		allPoints := make([][3]int, 0, 343)
		for x := 0; x <= 6; x++ {
			for y := 0; y <= 6; y++ {
				for z := 0; z <= 6; z++ {
					if seen[x][y][z] {
						allPoints = append(allPoints, [3]int{x, y, z})
					}
				}
			}
		}

		if len(allPoints) <= 1 {
			break // need at least 2 distinct points to generate
		}

		var nextGen [][3]int
		for i := 0; i < len(allPoints); i++ {
			for j := i + 1; j < len(allPoints); j++ {
				// Compute floor midpoint
				nx := (allPoints[i][0] + allPoints[j][0]) / 2
				ny := (allPoints[i][1] + allPoints[j][1]) / 2
				nz := (allPoints[i][2] + allPoints[j][2]) / 2

				if !seen[nx][ny][nz] {
					seen[nx][ny][nz] = true
					nextGen = append(nextGen, [3]int{nx, ny, nz})
				}
			}
		}

		if seen[tx][ty][tz] {
			return gen
		}

		if len(nextGen) == 0 {
			break
		}

		curGen = nextGen
	}

	return -1
}

func main() {
	// Example 1
	fmt.Println(MinimumGenerationsToTargetPoint([][]int{{0, 0, 0}, {6, 6, 6}}, []int{3, 3, 3})) // Expected: 1

	// Example 2
	fmt.Println(MinimumGenerationsToTargetPoint([][]int{{0, 0, 0}, {5, 5, 5}}, []int{1, 1, 1})) // Expected: 2

	// Example 3
	fmt.Println(MinimumGenerationsToTargetPoint([][]int{{0, 0, 0}, {2, 2, 2}, {3, 3, 3}}, []int{2, 2, 2})) // Expected: 0

	// Example 4
	fmt.Println(MinimumGenerationsToTargetPoint([][]int{{1, 2, 3}}, []int{5, 5, 5})) // Expected: -1
}
```

## 3926 — Count Valid Word Occurrences

```go
package main

// LeetCode #3926: Count Valid Word Occurrences
// https://leetcode.com/problems/count-valid-word-occurrences/
// Difficulty: Medium
// Time: O(N + Q) | Space: O(N)
// Approach: Concatenate chunks, extract words (lowercase letters +
// joiner hyphens), count with hash map, look up queries.

import (
	"fmt"
	"strings"
)

func CountValidWordOccurrences(chunks []string, queries []string) []int {
	s := strings.Join(chunks, "")

	wordCount := make(map[string]int)
	n := len(s)
	i := 0
	for i < n {
		// Skip non-word chars (spaces, leading hyphens, isolated hyphens)
		if s[i] == ' ' {
			i++
			continue
		}
		if s[i] == '-' {
			i++
			continue
		}

		// Start of a word
		j := i
		for j < n {
			ch := s[j]
			if ch >= 'a' && ch <= 'z' {
				j++
				continue
			}
			if ch == '-' && j+1 < n && s[j+1] >= 'a' && s[j+1] <= 'z' {
				// Joiner hyphen: part of current word
				j += 2
				continue
			}
			break
		}

		if j > i {
			wordCount[s[i:j]]++
		}
		i = j
	}

	ans := make([]int, len(queries))
	for idx, q := range queries {
		ans[idx] = wordCount[q]
	}
	return ans
}

func main() {
	// Example 1
	fmt.Println(CountValidWordOccurrences(
		[]string{"hello wor", "ld hello"},
		[]string{"hello", "world", "wor"},
	)) // Expected: [2 1 0]

	// Example 2
	fmt.Println(CountValidWordOccurrences(
		[]string{"a-b a--b ", "a-", "b"},
		[]string{"a-b", "a", "b"},
	)) // Expected: [2 1 1]

	// Example 3
	fmt.Println(CountValidWordOccurrences(
		[]string{"-cat dog- mouse"},
		[]string{"cat", "dog", "mouse", "cat-dog"},
	)) // Expected: [1 1 1 0]
}
```

## 3927 — Minimize Array Sum Using Divisible Replacements

```go
package main

// LeetCode #3927: Minimize Array Sum Using Divisible Replacements
// https://leetcode.com/problems/minimize-array-sum-using-divisible-replacements/
// Difficulty: Medium
// Time: O(N sqrt(M)) | Space: O(N) where M = max value
// Approach: For each element, find its minimum divisor present in the array.
// Each number can be replaced with any present divisor (via chaining).
// Sum the minimal reachable value for each element.

import (
	"fmt"
	"math"
)

func MinimizeArraySumUsingDivisibleReplacements(nums []int) int64 {
	// Track which values exist
	exists := make(map[int]bool)
	minVal := math.MaxInt32
	for _, v := range nums {
		exists[v] = true
		if v < minVal {
			minVal = v
		}
	}

	var ans int64 = 0
	for _, v := range nums {
		reduced := false
		// Try divisors from 1 to sqrt(v)
		for d := 1; d*d <= v; d++ {
			if v%d == 0 {
				if exists[d] {
					ans += int64(d)
					reduced = true
					break
				}
				other := v / d
				if other != d && exists[other] && other < v {
					// Can't directly use if not minimum, but may be useful
				}
			}
		}
		if !reduced {
			ans += int64(v)
		}
	}
	return ans
}

func main() {
	// Example 1
	fmt.Println(MinimizeArraySumUsingDivisibleReplacements([]int{3, 6, 2})) // Expected: 7

	// Example 2
	fmt.Println(MinimizeArraySumUsingDivisibleReplacements([]int{4, 2, 8, 3})) // Expected: 9

	// Example 3
	fmt.Println(MinimizeArraySumUsingDivisibleReplacements([]int{7, 5, 9})) // Expected: 21
}
```

## 3932 — Count K Th Roots In A Range

```go
package main

// LeetCode #3932: Count K-th Roots in a Range
// https://leetcode.com/problems/count-k-th-roots-in-a-range/
// Difficulty: Medium
// Time: O(log r) | Space: O(1)
// Approach: Find smallest x s.t. x^k >= l, largest x s.t. x^k <= r.
// Count = hi - lo + 1. Use binary search to avoid overflow.

import (
	"fmt"
)

func powWithLimit(base int, k int, limit int) int {
	// Returns base^k, capped at limit+1 to avoid overflow
	result := 1
	for i := 0; i < k; i++ {
		if result > limit/base {
			return limit + 1
		}
		result *= base
	}
	return result
}

func CountKThRootsInARange(l int, r int, k int) int {
	// Find first x s.t. x^k >= l
	hiX := 1
	for powWithLimit(hiX, k, r) <= r {
		hiX *= 2
	}

	left := 1
	right := hiX
	for left < right {
		mid := left + (right-left)/2
		if powWithLimit(mid, k, r) >= l {
			right = mid
		} else {
			left = mid + 1
		}
	}
	first := left

	if powWithLimit(first, k, r) > r {
		return 0
	}

	// Find last x s.t. x^k <= r
	left, right = first, hiX
	for left < right {
		mid := (left + right + 1) / 2
		if powWithLimit(mid, k, r) <= r {
			left = mid
		} else {
			right = mid - 1
		}
	}
	last := left

	return last - first + 1
}

func main() {
	// Example 1
	fmt.Println(CountKThRootsInARange(1, 9, 3)) // Expected: 2

	// Example 2
	fmt.Println(CountKThRootsInARange(8, 30, 2)) // Expected: 3
}
```

## 3933 — Largest Local Values In A Matrix Ii

```go
package main

// LeetCode #3933: Largest Local Values in a Matrix II
// https://leetcode.com/problems/largest-local-values-in-a-matrix-ii/
// Difficulty: Medium
// Time: O(N*M*V) | Space: O(N*M*V) where V = max value <= 200
// Approach: Precompute 2D prefix sums for each value threshold.
// For each cell (i,j) with value v > 0, check if any cell in the
// (2v+1)x(2v+1) region (minus corners) has value > v.

import "fmt"

func LargestLocalValuesInAMatrixIi(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])

	// Find max value in grid
	maxVal := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] > maxVal {
				maxVal = grid[i][j]
			}
		}
	}

	// Precompute 2D prefix sums for each threshold t: count of cells >= t
	// pref[t][i+1][j+1] = count of cells with value >= t in rectangle [0..i]x[0..j]
	pref := make([][][]int, maxVal+2)
	for t := 1; t <= maxVal+1; t++ {
		p := make([][]int, m+1)
		for i := 0; i <= m; i++ {
			p[i] = make([]int, n+1)
		}
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				val := 0
				if grid[i][j] >= t {
					val = 1
				}
				p[i+1][j+1] = p[i][j+1] + p[i+1][j] - p[i][j] + val
			}
		}
		pref[t] = p
	}

	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			v := grid[i][j]
			if v == 0 {
				continue
			}

			// Define the check region: (2v+1)x(2v+1) minus corners
			r1 := max(0, i-v)
			c1 := max(0, j-v)
			r2 := min(m-1, i+v)
			c2 := min(n-1, j+v)

			// Count cells with value > v in the full rectangle
			p := pref[v+1]
			total := p[r2+1][c2+1] - p[r1][c2+1] - p[r2+1][c1] + p[r1][c1]

			// Subtract corners if they're within bounds and distinct from edges
			corners := [][2]int{
				{i - v, j - v},
				{i - v, j + v},
				{i + v, j - v},
				{i + v, j + v},
			}
			for _, c := range corners {
				cr, cc := c[0], c[1]
				if cr >= 0 && cr < m && cc >= 0 && cc < n {
					if grid[cr][cc] > v {
						total--
					}
				}
			}

			if total == 0 {
				ans++
			}
		}
	}

	return ans
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
	// Example 1
	grid1 := make([][]int, 7)
	for i := 0; i < 7; i++ {
		grid1[i] = make([]int, 7)
	}
	grid1[3][3] = 2
	fmt.Println(LargestLocalValuesInAMatrixIi(grid1)) // Expected: 1

	// Example 2
	fmt.Println(LargestLocalValuesInAMatrixIi([][]int{{1, 2}, {3, 4}})) // Expected: 1

	// Example 3
	fmt.Println(LargestLocalValuesInAMatrixIi([][]int{{1, 0, 1}, {0, 1, 0}, {1, 0, 1}})) // Expected: 5

	// Example 4
	fmt.Println(LargestLocalValuesInAMatrixIi([][]int{{1, 1}, {1, 1}})) // Expected: 4
}
```

## 3935 — Power Update After K Th Largest Insertion I

```go
package main

// LeetCode #3935: Power Update After K-th Largest Insertion I
// https://leetcode.com/problems/power-update-after-k-th-largest-insertion-i/
// Difficulty: Medium [Paid]
// Time: O(N log M + Q log M) | Space: O(M) where M = max value
// Approach: Maintain sorted multiset via Fenwick tree. For each query,
// insert val, find k-th largest via binary search on BIT prefix sums,
// update p = p ^ kth % MOD.

import (
	"fmt"
	"sort"
)

const MOD = 1000000007

type BIT struct {
	tree []int
	size int
}

func NewBIT(size int) *BIT {
	return &BIT{tree: make([]int, size+2), size: size}
}

func (b *BIT) Update(idx int, delta int) {
	idx++
	for idx <= b.size+1 {
		b.tree[idx] += delta
		idx += idx & -idx
	}
}

func (b *BIT) Query(idx int) int {
	idx++
	sum := 0
	for idx > 0 {
		sum += b.tree[idx]
		idx -= idx & -idx
	}
	return sum
}

func (b *BIT) KthLargest(k int) int {
	// Find smallest idx such that suffix sum (total - prefix) >= k
	// i.e., total - Query(idx-1) >= k
	// i.e., Query(idx-1) <= total - k
	total := b.Query(b.size - 1)
	target := total - k // number of elements strictly less than the kth largest

	// Binary search for first index with prefix sum > target
	lo, hi := 0, b.size-1
	for lo < hi {
		mid := (lo + hi) / 2
		if b.Query(mid) > target {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func PowerUpdateAfterKThLargestInsertionI(nums []int, p int, queries [][]int) []int {
	// Coordinate compress all values
	allVals := make([]int, 0, len(nums)+len(queries))
	allVals = append(allVals, nums...)
	for _, q := range queries {
		allVals = append(allVals, q[0])
	}
	sort.Ints(allVals)
	uniq := []int{allVals[0]}
	for i := 1; i < len(allVals); i++ {
		if allVals[i] != allVals[i-1] {
			uniq = append(uniq, allVals[i])
		}
	}

	rank := make(map[int]int)
	for i, v := range uniq {
		rank[v] = i
	}
	m := len(uniq)

	bit := NewBIT(m)
	for _, v := range nums {
		bit.Update(rank[v], 1)
	}

	cur := int64(p)
	ans := make([]int, len(queries))

	for idx, q := range queries {
		val, k := q[0], q[1]
		bit.Update(rank[val], 1)

		kthVal := uniq[bit.KthLargest(k)]

		cur = (cur ^ int64(kthVal)) % MOD
		ans[idx] = int(cur)
	}

	return ans
}

func main() {
	// Example 1
	fmt.Println(PowerUpdateAfterKThLargestInsertionI(
		[]int{2}, 4, [][]int{{3, 1}, {1, 2}},
	)) // Expected: [64, 4096]

	// Example 2
	fmt.Println(PowerUpdateAfterKThLargestInsertionI(
		[]int{7, 5}, 6, [][]int{{4, 3}, {7, 2}},
	)) // Expected: [1296, 220296870]
}
```

## 3937 — Minimum Operations To Make Array Modulo Alternating I

```go
package main

// LeetCode #3937: Minimum Operations to Make Array Modulo Alternating I
// https://leetcode.com/problems/minimum-operations-to-make-array-modulo-alternating-i/
// Difficulty: Medium
// Time: O(N * K^2) | Space: O(1)
// Approach: Enumerate all (x,y) pairs with x != y, 0 <= x,y < k.
// Even indices need modulo x, odd need modulo y.
// Cost per element = min(|curr - target|, k - |curr - target|).

import (
	"fmt"
)

func MinimumOperationsToMakeArrayModuloAlternatingI(nums []int, k int) int {
	n := len(nums)

	// Precompute remainders
	rem := make([]int, n)
	for i := 0; i < n; i++ {
		rem[i] = nums[i] % k
	}

	ans := -1
	for x := 0; x < k; x++ {
		for y := 0; y < k; y++ {
			if x == y {
				continue
			}
			cost := 0
			for i := 0; i < n; i++ {
				target := x
				if i%2 == 1 {
					target = y
				}
				diff := rem[i] - target
				if diff < 0 {
					diff = -diff
				}
				if diff > k-diff {
					diff = k - diff
				}
				cost += diff
			}
			if ans == -1 || cost < ans {
				ans = cost
			}
		}
	}

	return ans
}

func main() {
	// Example 1
	fmt.Println(MinimumOperationsToMakeArrayModuloAlternatingI([]int{1, 4, 2, 8}, 3)) // Expected: 2

	// Example 2
	fmt.Println(MinimumOperationsToMakeArrayModuloAlternatingI([]int{1, 1, 1}, 3)) // Expected: 1
}
```

## 3938 — Maximum Path Intersection Sum In A Grid

```go
package main

// LeetCode #3938: Maximum Path Intersection Sum in a Grid
// https://leetcode.com/problems/maximum-path-intersection-sum-in-a-grid/
// Difficulty: Medium
// Time: O(C * R^2) | Space: O(R^2) where R = min(m,n), C = max(m,n)
// Approach: Column-by-column DP with 2D prefix max optimization.
// dp[r1][r2] = max shared sum with P1 at exit row r1, P2 at exit row r2.
// For each column, compute shared interval overlap between the two paths
// within the column. Transition via 4 overlap cases, each O(1) using
// precomputed 2D prefix max arrays. Transpose grid so DP dim = min(m,n).

import (
	"fmt"
	"math"
)

const negInf = math.MinInt32 / 2

func maxPathIntersectionSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	// Transpose if m > n to make DP dimension the smaller one
	if m > n {
		t := make([][]int, n)
		for i := 0; i < n; i++ {
			t[i] = make([]int, m)
			for j := 0; j < m; j++ {
				t[i][j] = grid[j][i]
			}
		}
		grid = t
		m, n = n, m
	}
	// Now m <= n. DP dim = m.
	// P1 starts at (0,0), ends at (m-1,n-1), moves right/down.
	// P2 starts at (m-1,0), ends at (0,n-1), moves right/up.

	// Column prefix sums
	pref := make([][]int, n)
	for c := 0; c < n; c++ {
		pref[c] = make([]int, m+1)
		for r := 0; r < m; r++ {
			pref[c][r+1] = pref[c][r] + grid[r][c]
		}
	}

	// dp[r1][r2] after column 0
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, m)
		for j := 0; j < m; j++ {
			dp[i][j] = negInf
		}
	}

	// Column 0: P1 enters at row 0, P2 enters at row m-1
	// P1 visits [0, r1], P2 visits [r2, m-1]
	// Overlap = [max(0, r2), min(r1, m-1)] = [r2, r1] if r2 <= r1
	for r1 := 0; r1 < m; r1++ {
		for r2 := 0; r2 <= r1; r2++ {
			dp[r1][r2] = pref[0][r1+1] - pref[0][r2]
		}
	}

	// Process columns 1..n-1
	for c := 1; c < n; c++ {
		newdp := make([][]int, m)
		for i := 0; i < m; i++ {
			newdp[i] = make([]int, m)
			for j := 0; j < m; j++ {
				newdp[i][j] = negInf
			}
		}

		// 2D prefix max of dp (upper-left quadrant: p1 <= i, p2 >= j)
		pmax := make([][]int, m)
		for i := 0; i < m; i++ {
			pmax[i] = make([]int, m+1)
			for j := 0; j <= m; j++ {
				pmax[i][j] = negInf
			}
		}
		// Row-wise suffix max
		for i := 0; i < m; i++ {
			best := negInf
			for j := m - 1; j >= 0; j-- {
				if dp[i][j] > best {
					best = dp[i][j]
				}
				pmax[i][j] = best
			}
		}
		// Column-wise prefix max
		for j := 0; j < m; j++ {
			for i := 1; i < m; i++ {
				if pmax[i-1][j] > pmax[i][j] {
					pmax[i][j] = pmax[i-1][j]
				}
			}
		}

		// Global best (no overlap in this column)
		globalBest := negInf
		for i := 0; i < m; i++ {
			for j := 0; j < m; j++ {
				if dp[i][j] > globalBest {
					globalBest = dp[i][j]
				}
			}
		}

		// For each (r1, r2), compute best transition
		for r1 := 0; r1 < m; r1++ {
			for r2 := 0; r2 < m; r2++ {
				best := globalBest

				// Case 1: p1 <= r2 AND r1 <= p2
				// overlap = [r2, r1], sum = pref[r1+1] - pref[r2]
				if r2 < m && r1 < m {
					val := pmax[r2][r1]
					if val > negInf/2 {
						cand := val + pref[c][r1+1] - pref[c][r2]
						if cand > best {
							best = cand
						}
					}
				}

				// Case 2 & 4: scan p1 from r2+1 to r1
				// Need max_{p2 >= r1} dp[p1][p2] for each p1
				if r1 > r2 {
					// Precompute suffix max per p1 (computed lazily)
					for p1 := r2 + 1; p1 <= r1; p1++ {
						// suffix max over p2 >= r1
						f2 := negInf
						for p2 := r1; p2 < m; p2++ {
							if dp[p1][p2] > f2 {
								f2 = dp[p1][p2]
							}
						}
						if f2 > negInf/2 {
							// Case 2: p1 > r2 AND r1 <= p2 (p2 >= r1 always)
							// overlap = [p1, r1], sum = pref[r1+1] - pref[p1]
							cand2 := f2 + pref[c][r1+1] - pref[c][p1]
							if cand2 > best {
								best = cand2
							}

							// Case 3: p1 <= r2 (already covered by Case 1)
							// Case 4: p1 > r2 AND r1 > p2 (need p2 < r1)
							// overlap = [p1, p2] where p1 <= p2 < r1
							// max over p2 in [p1, r1-1]
							for p2 := p1; p2 < r1 && p2 < m; p2++ {
								if dp[p1][p2] > negInf/2 {
									cand4 := dp[p1][p2] + pref[c][p2+1] - pref[c][p1]
									if cand4 > best {
										best = cand4
									}
								}
							}

							// Case 3 also: p1 <= r2 AND r1 > p2
							// Already covered by Case 1 for the overlap sum
							// But Case 3 = [r2, p2], which differs from Case 1 = [r2, r1]
							// Need: max over p1 <= r2, p2 < r1
							// This is separate from the p1 > r2 loop
						}
					}

					// Case 3: p1 <= r2 AND r1 > p2 (p2 < r1)
					// overlap = [r2, p2], sum = pref[p2+1] - pref[r2]
					// Need max over p1 <= r2, p2 < r1
					if r1 > 0 {
						for p1 := 0; p1 <= r2 && p1 < m; p1++ {
							for p2 := r2; p2 < r1 && p2 < m; p2++ {
								if dp[p1][p2] > negInf/2 {
									cand3 := dp[p1][p2] + pref[c][p2+1] - pref[c][r2]
									if cand3 > best {
										best = cand3
									}
								}
							}
						}
					}
				}

				newdp[r1][r2] = best
			}
		}

		dp = newdp
	}

	// Answer: max over all exit pairs
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			if dp[i][j] > ans {
				ans = dp[i][j]
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(maxPathIntersectionSum([][]int{{1, 2, 0, -3}, {1, -2, 1, 0}, {-4, 2, -1, 3}, {3, -3, 3, -2}, {-1, -5, 0, 1}}))
	fmt.Println(maxPathIntersectionSum([][]int{{4, -2, -3}, {-1, -3, -1}, {-4, 2, -1}}))
}
```

## 3941 — Password Strength

```go
package main

// LeetCode #3941: Password Strength
// https://leetcode.com/problems/password-strength/
// Difficulty: Medium
// Time: O(N) | Space: O(1)
// Approach: Track distinct characters per category. Sum points:
// lowercase=1, uppercase=2, digit=3, special=5.

import "fmt"

func PasswordStrength(password string) int {
	lower := make(map[byte]bool)
	upper := make(map[byte]bool)
	digit := make(map[byte]bool)
	special := make(map[byte]bool)

	for i := 0; i < len(password); i++ {
		ch := password[i]
		if ch >= 'a' && ch <= 'z' {
			lower[ch] = true
		} else if ch >= 'A' && ch <= 'Z' {
			upper[ch] = true
		} else if ch >= '0' && ch <= '9' {
			digit[ch] = true
		} else {
			// Special: ! @ # $
			special[ch] = true
		}
	}

	return len(lower)*1 + len(upper)*2 + len(digit)*3 + len(special)*5
}

func main() {
	// Example 1
	fmt.Println(PasswordStrength("aA1!")) // Expected: 11

	// Example 2
	fmt.Println(PasswordStrength("bbB11#")) // Expected: 11
}
```

## 3942 — Minimum Operations To Sort A Permutation

```go
package main

// LeetCode #3942: Minimum Operations to Sort a Permutation
// https://leetcode.com/problems/minimum-operations-to-sort-a-permutation/
// Difficulty: Medium
// Time: O(N) | Space: O(N)
// Approach: Available ops are reverse and rotate left. Sorted array must be
// [0,1,...,n-1] or a rotation of it. Check all rotations of sorted array vs
// all rotations of array and its reverse. Return min ops.

import "fmt"

func MinimumOperationsToSortAPermutation(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	// Expected sorted array
	sorted := make([]int, n)
	for i := 0; i < n; i++ {
		sorted[i] = i
	}

	// Find position of 0 in nums. In sorted, 0 is at index 0.
	// After rotation, 0 moves. Sorted rotated by k: 0 at index (n-k)%n.
	// Actually: rotate left by k means element at index k moves to 0.
	// sorted left-rotated by k: [k, k+1, ..., n-1, 0, 1, ..., k-1]
	// In this rotation, 0 is at index n-k.

	// Check all rotations of sorted vs all rotations of nums
	// Ops to get from nums to a target rotation of sorted:
	// 1. Maybe reverse first, then rotate
	// 2. Maybe rotate first, then reverse

	// Strategy: try all possible rotations.

	// First, find pos of 0 in nums
	pos0 := -1
	for i, v := range nums {
		if v == 0 {
			pos0 = i
			break
		}
	}

	ans := -1

	// Try: rotate nums to match sorted (no reverse)
	// Need to rotate so that sorted[0]=0 aligns with nums[pos0]=0
	// Rotate left by pos0: nums becomes [nums[pos0], nums[pos0+1], ..., nums[(pos0-1+n)%n]]
	// This puts '0' at front, which matches sorted[0]=0
	// Then check if everything matches
	matches := true
	for i := 0; i < n; i++ {
		if nums[(pos0+i)%n] != i {
			matches = false
			break
		}
	}
	if matches {
		ans = pos0 // just rotate, no reverse
	}

	// Try: reverse then rotate to match sorted
	// Reverse: [nums[n-1], nums[n-2], ..., nums[0]]
	reversed := make([]int, n)
	for i := 0; i < n; i++ {
		reversed[i] = nums[n-1-i]
	}

	// Find pos of 0 in reversed
	pos0Rev := -1
	for i, v := range reversed {
		if v == 0 {
			pos0Rev = i
			break
		}
	}

	matches = true
	for i := 0; i < n; i++ {
		if reversed[(pos0Rev+i)%n] != i {
			matches = false
			break
		}
	}
	if matches {
		ops := 1 + pos0Rev // 1 reverse + pos0Rev rotations
		if ans == -1 || ops < ans {
			ans = ops
		}
	}

	// Try: rotate then reverse
	// But which rotation? We need sorted after all ops.
	// sorted rotated by some amount k.
	// We can also rotate nums (by r), then reverse.
	// sorted left-rotated by k: element k, k+1, ..., n-1, 0, ..., k-1
	// After reversal of (nums rotated by r): we get reversed_rotated
	// reversed_rotated[i] = nums[(r-1-i+n)%n] (since rotate left by r, then reverse)

	for k := 0; k < n; k++ {
		// Check if (nums rotated by r) reversed equals sorted rotated by k
		// for some r

		for r := 0; r < n; r++ {
			matches = true
			for i := 0; i < n; i++ {
				// rot = nums[(r+i)%n]
				// rev_rot = rot[n-1-i] = nums[(r+(n-1-i))%n]
				// should equal sorted[(k+i)%n]
				expected := (k + i) % n
				actual := nums[(r + n - 1 - i) % n]
				if actual != expected {
					matches = false
					break
				}
			}
			if matches {
				ops := r + 1 + k // r rotations, 1 reverse, k rotations
				if ans == -1 || ops < ans {
					ans = ops
				}
			}
		}
	}

	return ans
}

func main() {
	// Example 1
	fmt.Println(MinimumOperationsToSortAPermutation([]int{0, 2, 1})) // Expected: 2

	// Example 2
	fmt.Println(MinimumOperationsToSortAPermutation([]int{1, 0, 2})) // Expected: 2

	// Example 3
	fmt.Println(MinimumOperationsToSortAPermutation([]int{2, 0, 1, 3})) // Expected: -1
}
```

## 3946 — Maximum Number Of Items From Sale I

```go
package main

// LeetCode #3946: Maximum Number of Items From Sale I
// https://leetcode.com/problems/maximum-number-of-items-from-sale-i/
// Difficulty: Medium
// Time: O(N * budget + N^2) | Space: O(budget) where N = len(items)
// Approach: 0-1 knapsack. Each item's first copy gives (1 + out_degree)
// copies at cost price_i. Additional copies of cheapest item give 1 each.
// out_degree[i] = count of j where factor_i divides factor_j, j != i.

import (
	"fmt"
	"math"
)

func MaximumNumberOfItemsFromSaleI(items [][]int, budget int) int {
	m := len(items)

	// Compute out_degree: how many other items this item's factor divides
	outDeg := make([]int, m)
	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			if i != j && items[j][0]%items[i][0] == 0 {
				outDeg[i]++
			}
		}
	}

	// Find min price (for additional copies after activation)
	minPrice := math.MaxInt32
	for _, it := range items {
		if it[1] < minPrice {
			minPrice = it[1]
		}
	}

	// 0-1 knapsack: dp[b] = max copies from first copies with budget b
	dp := make([]int, budget+1)
	for b := 1; b <= budget; b++ {
		dp[b] = math.MinInt32
	}
	dp[0] = 0

	for i := 0; i < m; i++ {
		price := items[i][1]
		val := 1 + outDeg[i] // first copy value
		for b := budget; b >= price; b-- {
			if dp[b-price] > math.MinInt32 {
				cand := dp[b-price] + val
				if cand > dp[b] {
					dp[b] = cand
				}
			}
		}
	}

	// Best result: dp[b] + floor((budget-b)/minPrice) cheap copies
	best := 0
	for b := 0; b <= budget; b++ {
		if dp[b] > math.MinInt32 {
			total := dp[b] + (budget-b)/minPrice
			if total > best {
				best = total
			}
		}
	}

	return best
}

func main() {
	// Example 1
	fmt.Println(MaximumNumberOfItemsFromSaleI([][]int{{6, 2}, {2, 6}, {3, 4}}, 9)) // Expected: 4

	// Example 2
	fmt.Println(MaximumNumberOfItemsFromSaleI([][]int{{2, 4}, {3, 2}, {4, 1}, {6, 4}, {12, 4}}, 8)) // Expected: 10
}
```

## 3947 — Maximum Number Of Items From Sale Ii

```go
package main

// LeetCode #3947: Maximum Number of Items From Sale II
// https://leetcode.com/problems/maximum-number-of-items-from-sale-ii/
// Difficulty: Medium
// Time: O(N log N) | Space: O(N) where N = len(items)
// Approach: Each purchased copy of item i gives at most 1 free copy of a
// different item j where factor_i | factor_j (at most once per ordered pair).
// So first out_degree[i] copies of item i each give 2 items (1 purchased
// + 1 free). Subsequent copies give 1 each. Sort bonus copies by price,
// greedily buy cheapest bonus copies, then buy cheapest regular copies.

import (
	"fmt"
	"sort"
)

func MaximumNumberOfItemsFromSaleIi(items [][]int, budget int) int {
	m := len(items)

	// Compute out_degree per item (how many j where i|j, j != i)
	outDeg := make([]int, m)
	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			if i != j && items[j][0]%items[i][0] == 0 {
				outDeg[i]++
			}
		}
	}

	// Collect all bonus opportunities: each gives 2 copies at price price_i
	type bonus struct {
		price int
	}
	var bonuses []bonus
	for i := 0; i < m; i++ {
		for k := 0; k < outDeg[i]; k++ {
			bonuses = append(bonuses, bonus{price: items[i][1]})
		}
	}

	// Sort bonuses by price
	sort.Slice(bonuses, func(i, j int) bool {
		return bonuses[i].price < bonuses[j].price
	})

	// Find min price for regular copies
	minPrice := int(1e9 + 1)
	for _, it := range items {
		if it[1] < minPrice {
			minPrice = it[1]
		}
	}

	remaining := budget
	total := 0

	// Buy cheapest bonuses first
	for _, b := range bonuses {
		if remaining < b.price {
			break
		}
		// Bonus gives 2 copies for price_i
		// Only buy if better than 2 regular copies at minPrice
		if b.price < 2*minPrice {
			remaining -= b.price
			total += 2
		}
	}

	// Buy regular copies with remaining budget
	total += remaining / minPrice

	return total
}

func main() {
	// Example 1
	fmt.Println(MaximumNumberOfItemsFromSaleIi([][]int{{1, 6}, {2, 4}, {3, 5}}, 19)) // Expected: 5

	// Example 2
	fmt.Println(MaximumNumberOfItemsFromSaleIi([][]int{{2, 8}, {1, 10}, {6, 6}, {4, 12}, {5, 20}, {5, 17}}, 35)) // Expected: 7
}
```

## 3951 — Minimum Energy To Maintain Brightness

```go
package main

// LeetCode #3951: Minimum Energy to Maintain Brightness
// https://leetcode.com/problems/minimum-energy-to-maintain-brightness/
// Difficulty: Medium
// Time: O(N log N) | Space: O(N) where N = len(intervals)
// Approach: Each bulb illuminates 3 positions (self + adjacents). Need
// ceil(brightness/3) bulbs. Merge overlapping intervals. For each merged
// interval, energy += bulbs * intervalLength. Return total energy.

import (
	"fmt"
	"sort"
)

func MinimumEnergyToMaintainBrightness(n int, brightness int, intervals [][]int) int64 {
	bulbs := int64((brightness + 2) / 3)

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var total int64
	i := 0
	for i < len(intervals) {
		start := intervals[i][0]
		end := intervals[i][1]
		j := i + 1
		for j < len(intervals) && intervals[j][0] <= end+1 {
			if intervals[j][1] > end {
				end = intervals[j][1]
			}
			j++
		}
		total += int64(end - start + 1)
		i = j
	}

	return bulbs * total
}

func main() {
	// Example 1
	fmt.Println(MinimumEnergyToMaintainBrightness(5, 5, [][]int{{6, 12}})) // Expected: 14

	// Example 2
	fmt.Println(MinimumEnergyToMaintainBrightness(2, 1, [][]int{{0, 0}, {2, 2}})) // Expected: 2

	// Example 3
	fmt.Println(MinimumEnergyToMaintainBrightness(4, 2, [][]int{{1, 3}, {2, 4}})) // Expected: 4
}
```

## 3952 — Maximum Total Value Of Covered Indices

```go
package main

// LeetCode #3952: Maximum Total Value of Covered Indices
// https://leetcode.com/problems/maximum-total-value-of-covered-indices/
// Difficulty: Medium
// Time: O(N) | Space: O(N)
// Approach: DP over positions. Token at i can stay (cover i) or move left
// (cover i-1). dp[i][0] = token at i moves left, dp[i][1] = token at i stays.
// No two tokens can cover same index.

import "fmt"

func MaximumTotalValueOfCoveredIndices(nums []int, s string) int64 {
	n := len(nums)
	const negInf int64 = -1 << 60

	dp0, dp1 := int64(0), negInf

	for i := 0; i < n; i++ {
		if s[i] == '0' {
			// No token at i, carry forward previous best
			newDp0 := maxInt64(dp0, dp1)
			dp0, dp1 = newDp0, negInf
		} else {
			// Token at i
			prevDp0, prevDp1 := dp0, dp1

			// Stay at i: covers i, no conflict with any previous decision
			stay := maxInt64(prevDp0, prevDp1) + int64(nums[i])

			// Move left to i-1: covers i-1
			// Requires prev token didn't stay at i-1
			move := negInf
			if i >= 1 {
				// Need prevDp0 (prev token moved left or no prev token)
				move = prevDp0 + int64(nums[i-1])
			}

			dp0 = move
			dp1 = stay
		}
	}

	return maxInt64(dp0, dp1)
}

func maxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Example 1
	fmt.Println(MaximumTotalValueOfCoveredIndices([]int{9, 2, 6, 1}, "0101")) // Expected: 15

	// Example 2
	fmt.Println(MaximumTotalValueOfCoveredIndices([]int{5, 1, 4}, "001")) // Expected: 4

	// Example 3
	fmt.Println(MaximumTotalValueOfCoveredIndices([]int{9, 3, 5}, "011")) // Expected: 14
}
```

## 3955 — Valid Binary Strings With Cost Limit

```go
package main

// LeetCode #3955: Valid Binary Strings With Cost Limit
// https://leetcode.com/problems/valid-binary-strings-with-cost-limit/
// Difficulty: Medium
// Time: O(2^N) | Space: O(N * 2^N)
// Approach: Backtracking. Generate all strings without consecutive 1s,
// filter by cost (sum of indices where s[i]=='1') <= k.

import "fmt"

func ValidBinaryStringsWithCostLimit(n int, k int) []string {
	var ans []string
	var dfs func(pos int, prev byte, sum int, buf []byte)
	dfs = func(pos int, prev byte, sum int, buf []byte) {
		if sum > k {
			return
		}
		if pos == n {
			ans = append(ans, string(buf))
			return
		}

		// Place '0'
		buf[pos] = '0'
		dfs(pos+1, '0', sum, buf)

		// Place '1' only if no consecutive 1s
		if prev != '1' {
			buf[pos] = '1'
			dfs(pos+1, '1', sum+pos, buf)
		}
	}
	buf := make([]byte, n)
	dfs(0, '0', 0, buf)
	return ans
}

func main() {
	// Example 1
	fmt.Println(ValidBinaryStringsWithCostLimit(3, 1)) // Expected: ["000","010","100"]

	// Example 2
	fmt.Println(ValidBinaryStringsWithCostLimit(1, 0)) // Expected: ["0"]
}
```

## 3958 — Minimum Cost To Split Into Ones Ii

```go
package main

// LeetCode #3958: Minimum Cost to Split into Ones II
// https://leetcode.com/problems/minimum-cost-to-split-into-ones-ii/
// Difficulty: Medium [Paid]
// Time: O(1) | Space: O(1)
// Approach: Optimal strategy splits off one 1 at a time.
// Total cost = 1 + 2 + ... + (n-1) = n*(n-1)/2.

import "fmt"

func MinimumCostToSplitIntoOnesIi(n int) int64 {
	return int64(n) * int64(n-1) / 2
}

func main() {
	// Example 1
	fmt.Println(MinimumCostToSplitIntoOnesIi(3)) // Expected: 3

	// Example 2
	fmt.Println(MinimumCostToSplitIntoOnesIi(4)) // Expected: 6

	// Example 3
	fmt.Println(MinimumCostToSplitIntoOnesIi(1)) // Expected: 0
}
```

## 3960 — Frequency Balance Subarray

```go
package main

// LeetCode #3960: Frequency Balance Subarray
// https://leetcode.com/problems/frequency-balance-subarray/
// Difficulty: Medium
// Time: O(N^2) | Space: O(N)
// Approach: Enumerate all subarrays. Track element frequencies and
// frequency-of-frequencies. Valid if: 1 distinct value, or exactly 2
// distinct freq values where one is double the other.

import "fmt"

func FrequencyBalanceSubarray(nums []int) int {
	n := len(nums)
	ans := 0

	for l := 0; l < n; l++ {
		cnt := make(map[int]int)
		freq := make(map[int]int) // frequency-of-frequencies

		for r := l; r < n; r++ {
			val := nums[r]
			oldF := cnt[val]
			if oldF > 0 {
				freq[oldF]--
				if freq[oldF] == 0 {
					delete(freq, oldF)
				}
			}
			newF := oldF + 1
			cnt[val] = newF
			freq[newF]++

			// Check validity
			valid := false
			if len(cnt) == 1 {
				valid = true
			} else if len(freq) == 2 {
				// Exactly two distinct frequency values
				vals := make([]int, 0, 2)
				for f := range freq {
					vals = append(vals, f)
				}
				a, b := vals[0], vals[1]
				if a > b {
					a, b = b, a
				}
				// One must be double the other
				if b == 2*a {
					valid = true
				}
			}

			if valid && (r-l+1) > ans {
				ans = r - l + 1
			}
		}
	}

	return ans
}

func main() {
	// Example 1
	fmt.Println(FrequencyBalanceSubarray([]int{1, 2, 2, 1, 2, 3, 3, 3})) // Expected: 5

	// Example 2
	fmt.Println(FrequencyBalanceSubarray([]int{5, 5, 5, 5})) // Expected: 4

	// Example 3
	fmt.Println(FrequencyBalanceSubarray([]int{1, 2, 3, 4})) // Expected: 1
}
```

## 3961 — Maximize Sum Of Device Ratings

```go
package main

// LeetCode #3961: Maximize Sum of Device Ratings
// https://leetcode.com/problems/maximize-sum-of-device-ratings/
// Difficulty: Medium
// Time: O(M * N log N) | Space: O(M) where M = devices, N = units per device
// Approach: Sort each device's units. Rating after optimal transfers =
// second smallest value per device. Sum across devices, then adjust:
// subtract (minSecond - globalMin) since moving global min to the device
// with smallest second-min lowers that device's rating to globalMin.

import (
	"fmt"
	"math"
	"sort"
)

func MaximizeSumOfDeviceRatings(units [][]int) int64 {
	m := len(units)
	if m == 0 {
		return 0
	}
	n := len(units[0])

	if n == 1 {
		var sum int64
		for _, dev := range units {
			sum += int64(dev[0])
		}
		return sum
	}

	globalMin := math.MaxInt32
	minSecond := math.MaxInt32
	var sumSecond int64

	for _, dev := range units {
		sorted := make([]int, n)
		copy(sorted, dev)
		sort.Ints(sorted)

		if sorted[0] < globalMin {
			globalMin = sorted[0]
		}
		second := sorted[1]
		sumSecond += int64(second)
		if second < minSecond {
			minSecond = second
		}
	}

	return sumSecond - int64(minSecond-globalMin)
}

func main() {
	// Example 1
	fmt.Println(MaximizeSumOfDeviceRatings([][]int{{1, 3}, {2, 2}})) // Expected: 4

	// Example 2
	fmt.Println(MaximizeSumOfDeviceRatings([][]int{{1, 2, 3}, {4, 5, 6}})) // Expected: 6

	// Example 3
	fmt.Println(MaximizeSumOfDeviceRatings([][]int{{5, 5, 5}, {1, 1, 1}})) // Expected: 6
}
```

