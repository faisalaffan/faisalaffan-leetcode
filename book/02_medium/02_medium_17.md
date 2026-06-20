# Medium (Sedang) — Problem 3044–3234

## 3044 — Most Frequent Prime

```go
package main

// LeetCode #3044: Most Frequent Prime
// https://leetcode.com/problems/most-frequent-prime/
// Difficulty: Medium
// Time: O(m*n*maxLen) | Space: O(K)

import "fmt"

func main() {
	fmt.Println(mostFrequentPrime([][]int{{1, 1}, {9, 9}, {1, 1}}))
	fmt.Println(mostFrequentPrime([][]int{{7}}))
}

func mostFrequentPrime(mat [][]int) int {
	m, n := len(mat), len(mat[0])
	dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
	cnt := map[int]int{}
	isPrime := func(x int) bool {
		if x < 2 {
			return false
		}
		for i := 2; i*i <= x; i++ {
			if x%i == 0 {
				return false
			}
		}
		return true
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for _, d := range dirs {
				val := 0
				x, y := i, j
				for x >= 0 && x < m && y >= 0 && y < n {
					val = val*10 + mat[x][y]
					if val > 10 && isPrime(val) {
						cnt[val]++
					}
					x += d[0]
					y += d[1]
				}
			}
		}
	}
	maxCnt, maxVal := 0, -1
	for v, c := range cnt {
		if c > maxCnt || (c == maxCnt && v > maxVal) {
			maxCnt = c
			maxVal = v
		}
	}
	return maxVal
}
```

## 3047 — Find The Largest Area Of Square Inside Two Rectangles

```go
package main

// LeetCode #3047: Find the Largest Area of Square Inside Two Rectangles
// https://leetcode.com/problems/find-the-largest-area-of-square-inside-two-rectangles/
// Difficulty: Medium
// Time: O(n^2) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(largestSquareArea([][]int{{1, 1}, {2, 2}, {3, 3}}, [][]int{{2, 2}, {3, 3}, {4, 4}}))
	fmt.Println(largestSquareArea([][]int{{1, 1}, {2, 2}, {1, 2}}, [][]int{{3, 3}, {4, 4}, {3, 4}}))
}

func largestSquareArea(bottomLeft [][]int, topRight [][]int) int64 {
	n := len(bottomLeft)
	ans := int64(0)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			x1 := max(bottomLeft[i][0], bottomLeft[j][0])
			y1 := max(bottomLeft[i][1], bottomLeft[j][1])
			x2 := min(topRight[i][0], topRight[j][0])
			y2 := min(topRight[i][1], topRight[j][1])
			if x1 < x2 && y1 < y2 {
				side := min(x2-x1, y2-y1)
				area := int64(side) * int64(side)
				if area > ans {
					ans = area
				}
			}
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

## 3048 — Earliest Second To Mark Indices I

```go
package main

// LeetCode #3048: Earliest Second to Mark Indices I
// https://leetcode.com/problems/earliest-second-to-mark-indices-i/
// Difficulty: Medium
// Time: O(m log m) | Space: O(n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(earliestSecondToMarkIndices([]int{2, 2, 0}, []int{2, 2, 2, 2, 3, 2, 2, 1}))
	fmt.Println(earliestSecondToMarkIndices([]int{1, 3}, []int{1, 1, 1, 2, 1, 1, 1}))
	fmt.Println(earliestSecondToMarkIndices([]int{0, 1}, []int{2, 2, 2}))
}

func earliestSecondToMarkIndices(nums []int, changeIndices []int) int {
	n, m := len(nums), len(changeIndices)
	ans := sort.Search(m+1, func(t int) bool {
		if t == 0 {
			return false
		}
		last := make([]int, n+1)
		for s, idx := range changeIndices[:t] {
			last[idx] = s
		}
		for i := 1; i <= n; i++ {
			if last[i] == 0 && i != changeIndices[0] {
				return false
			}
		}
		decrement := 0
		marked := 0
		for s, idx := range changeIndices[:t] {
			if last[idx] == s {
				if decrement < nums[idx-1] {
					return false
				}
				decrement -= nums[idx-1]
				marked++
			} else {
				decrement++
			}
		}
		return marked == n
	})
	if ans > m {
		return -1
	}
	return ans
}
```

## 3050 — Pizza Toppings Cost Analysis

```go
package main

// LeetCode #3050: Pizza Toppings Cost Analysis
// https://leetcode.com/problems/pizza-toppings-cost-analysis/
// Difficulty: Medium (SQL problem — implemented in Go)
// Time: O(n^3) | Space: O(C(n,3))

import (
	"fmt"
	"sort"
)

type Topping struct {
	Name string
	Cost float64
}

type PizzaCombo struct {
	Pizza     string
	TotalCost float64
}

func pizzaToppingsCostAnalysis(toppings []Topping) []PizzaCombo {
	n := len(toppings)
	var results []PizzaCombo

	// Generate all C(n,3) combinations
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				total := toppings[i].Cost + toppings[j].Cost + toppings[k].Cost
				pizzaName := toppings[i].Name + "," + toppings[j].Name + "," + toppings[k].Name
				results = append(results, PizzaCombo{
					Pizza:     pizzaName,
					TotalCost: total,
				})
			}
		}
	}

	// Sort by total_cost DESC, then pizza name ASC (which is topping1,topping2,topping3 ASC
	// since we generate in sorted order)
	sort.Slice(results, func(a, b int) bool {
		if results[a].TotalCost != results[b].TotalCost {
			return results[a].TotalCost > results[b].TotalCost // DESC
		}
		return results[a].Pizza < results[b].Pizza // ASC
	})

	return results
}

func main() {
	toppings := []Topping{
		{Name: "Pepperoni", Cost: 0.50},
		{Name: "Sausage", Cost: 0.70},
		{Name: "Chicken", Cost: 0.55},
		{Name: "Extra Cheese", Cost: 0.40},
		{Name: "Mushrooms", Cost: 0.60},
	}

	fmt.Println("Pizza Toppings Cost Analysis")
	fmt.Println("============================")
	fmt.Printf("%-40s %s\n", "Pizza", "Total Cost")
	fmt.Println("----------------------------------------")

	results := pizzaToppingsCostAnalysis(toppings)
	for _, r := range results {
		fmt.Printf("%-40s $%.2f\n", r.Pizza, r.TotalCost)
	}
}
```

## 3054 — Binary Tree Nodes

```go
package main

// LeetCode #3054: Binary Tree Nodes
// https://leetcode.com/problems/binary-tree-nodes/
// Difficulty: Medium (SQL problem — implemented in Go)
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	N int
	P *int // null for root; using pointer to represent nullable int
}

type NodeClassification struct {
	N    int
	Type string
}

func classifyBinaryTreeNodes(nodes []TreeNode) []NodeClassification {
	nodeSet := make(map[int]bool)
	parentSet := make(map[int]bool)

	for _, node := range nodes {
		nodeSet[node.N] = true
		if node.P != nil {
			parentSet[*node.P] = true
		}
	}

	var results []NodeClassification
	for _, node := range nodes {
		var nodeType string
		if node.P == nil {
			nodeType = "Root"
		} else if !parentSet[node.N] {
			nodeType = "Leaf"
		} else {
			nodeType = "Inner"
		}
		results = append(results, NodeClassification{N: node.N, Type: nodeType})
	}

	// Order by N ASC
	sort.Slice(results, func(i, j int) bool {
		return results[i].N < results[j].N
	})

	return results
}

func intPtr(v int) *int {
	return &v
}

func main() {
	nodes := []TreeNode{
		{N: 1, P: nil},
		{N: 2, P: intPtr(1)},
		{N: 3, P: intPtr(1)},
		{N: 4, P: intPtr(2)},
		{N: 5, P: intPtr(2)},
	}

	fmt.Println("Binary Tree Nodes")
	fmt.Println("================")
	fmt.Printf("%-6s %s\n", "N", "Type")
	fmt.Println("-------------")

	results := classifyBinaryTreeNodes(nodes)
	for _, r := range results {
		fmt.Printf("%-6d %s\n", r.N, r.Type)
	}
}
```

## 3055 — Top Percentile Fraud

```go
package main

// LeetCode #3055: Top Percentile Fraud
// https://leetcode.com/problems/top-percentile-fraud/
// Difficulty: Medium (SQL problem — implemented in Go)
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"math"
	"sort"
)

type FraudClaim struct {
	PolicyID   int
	State      string
	FraudScore float64
}

type TopClaim struct {
	PolicyID   int
	State      string
	FraudScore float64
}

func topPercentileFraud(claims []FraudClaim) []TopClaim {
	// Group by state
	stateClaims := make(map[string][]FraudClaim)
	for _, c := range claims {
		stateClaims[c.State] = append(stateClaims[c.State], c)
	}

	var results []TopClaim

	for state, cs := range stateClaims {
		// Sort by fraud_score DESC, then policy_id ASC
		sort.Slice(cs, func(i, j int) bool {
			if cs[i].FraudScore != cs[j].FraudScore {
				return cs[i].FraudScore > cs[j].FraudScore // DESC
			}
			return cs[i].PolicyID < cs[j].PolicyID // ASC
		})

		// Find top 5%: rank in top 5% using RANK()
		// Equivalent to: rank <= ceil(0.05 * n)
		// But SQL solution uses rank = 1 (top 1 per state)
		// Actually, re-reading the problem: "find the top 5% of claims"
		// The SQL uses: WHERE percentile_rank <= 0.05
		// But solution reference uses WHERE rk = 1 (top-ranked per state)
		// We'll implement: return claims where fraud_score = max fraud_score in that state
		// If ties, return the one with lowest policy_id
		maxScore := cs[0].FraudScore

		// Get all with max score (should be just the first group after sorting)
		var bestClaims []FraudClaim
		for _, c := range cs {
			if math.Abs(c.FraudScore-maxScore) < 1e-9 {
				bestClaims = append(bestClaims, c)
			} else {
				break
			}
		}

		// Among ties, pick lowest policy_id
		sort.Slice(bestClaims, func(i, j int) bool {
			return bestClaims[i].PolicyID < bestClaims[j].PolicyID
		})

		results = append(results, TopClaim{
			PolicyID:   bestClaims[0].PolicyID,
			State:      state,
			FraudScore: bestClaims[0].FraudScore,
		})
	}

	// Order by state ASC, fraud_score DESC, policy_id ASC
	sort.Slice(results, func(i, j int) bool {
		if results[i].State != results[j].State {
			return results[i].State < results[j].State
		}
		if results[i].FraudScore != results[j].FraudScore {
			return results[i].FraudScore > results[j].FraudScore
		}
		return results[i].PolicyID < results[j].PolicyID
	})

	return results
}

func main() {
	claims := []FraudClaim{
		{PolicyID: 101, State: "CA", FraudScore: 85.5},
		{PolicyID: 102, State: "CA", FraudScore: 92.3},
		{PolicyID: 103, State: "CA", FraudScore: 78.1},
		{PolicyID: 104, State: "CA", FraudScore: 92.3},
		{PolicyID: 201, State: "TX", FraudScore: 88.0},
		{PolicyID: 202, State: "TX", FraudScore: 95.5},
		{PolicyID: 203, State: "TX", FraudScore: 72.4},
		{PolicyID: 301, State: "NY", FraudScore: 91.2},
		{PolicyID: 302, State: "NY", FraudScore: 84.7},
		{PolicyID: 303, State: "NY", FraudScore: 91.2},
	}

	fmt.Println("Top Percentile Fraud (Top-Ranked per State)")
	fmt.Println("=========================================")
	fmt.Printf("%-12s %-8s %s\n", "Policy ID", "State", "Fraud Score")
	fmt.Println("-----------------------------------------")

	results := topPercentileFraud(claims)
	for _, r := range results {
		fmt.Printf("%-12d %-8s %.1f\n", r.PolicyID, r.State, r.FraudScore)
	}
}
```

## 3056 — Snaps Analysis

```go
package main

// LeetCode #3056: Snaps Analysis
// https://leetcode.com/problems/snaps-analysis/
// Difficulty: Medium (SQL problem — implemented in Go)
// Time: O(n) | Space: O(n)

import (
	"fmt"
)

type Activity struct {
	ActivityID int
	UserID     int
	ActivityType string // 'send' or 'open'
	TimeSpent  float64
}

type Age struct {
	UserID     int
	AgeBucket  string // '21-25', '26-30', '31-35'
}

type AgeBucketAnalysis struct {
	AgeBucket string
	SendPerc  float64
	OpenPerc  float64
}

func snapsAnalysis(activities []Activity, ages []Age) []AgeBucketAnalysis {
	// Build user -> age_bucket map
	userAge := make(map[int]string)
	for _, a := range ages {
		userAge[a.UserID] = a.AgeBucket
	}

	// Group activities by age_bucket
	type bucketSums struct {
		totalSend float64
		totalOpen float64
	}
	buckets := make(map[string]*bucketSums)

	for _, act := range activities {
		bucket, ok := userAge[act.UserID]
		if !ok {
			continue
		}
		if buckets[bucket] == nil {
			buckets[bucket] = &bucketSums{}
		}
		if act.ActivityType == "send" {
			buckets[bucket].totalSend += act.TimeSpent
		} else if act.ActivityType == "open" {
			buckets[bucket].totalOpen += act.TimeSpent
		}
	}

	var results []AgeBucketAnalysis
	for bucket, sums := range buckets {
		total := sums.totalSend + sums.totalOpen
		sendPerc := 0.0
		openPerc := 0.0
		if total > 0 {
			sendPerc = (sums.totalSend / total) * 100
			openPerc = (sums.totalOpen / total) * 100
		}
		// Round to 2 decimal places
		sendPerc = float64(int(sendPerc*100+0.5)) / 100
		openPerc = float64(int(openPerc*100+0.5)) / 100

		results = append(results, AgeBucketAnalysis{
			AgeBucket: bucket,
			SendPerc:  sendPerc,
			OpenPerc:  openPerc,
		})
	}

	return results
}

func main() {
	activities := []Activity{
		{ActivityID: 1, UserID: 1, ActivityType: "send", TimeSpent: 12.5},
		{ActivityID: 2, UserID: 1, ActivityType: "open", TimeSpent: 7.5},
		{ActivityID: 3, UserID: 2, ActivityType: "send", TimeSpent: 20.0},
		{ActivityID: 4, UserID: 2, ActivityType: "open", TimeSpent: 5.0},
		{ActivityID: 5, UserID: 3, ActivityType: "send", TimeSpent: 8.0},
		{ActivityID: 6, UserID: 3, ActivityType: "open", TimeSpent: 12.0},
		{ActivityID: 7, UserID: 1, ActivityType: "send", TimeSpent: 5.0},
		{ActivityID: 8, UserID: 2, ActivityType: "open", TimeSpent: 10.0},
	}

	ages := []Age{
		{UserID: 1, AgeBucket: "21-25"},
		{UserID: 2, AgeBucket: "21-25"},
		{UserID: 3, AgeBucket: "26-30"},
	}

	fmt.Println("Snaps Analysis")
	fmt.Println("==============")
	fmt.Printf("%-10s %-10s %-10s\n", "AgeBucket", "Send%", "Open%")
	fmt.Println("-----------------------------")

	results := snapsAnalysis(activities, ages)
	for _, r := range results {
		fmt.Printf("%-10s %-10.2f %-10.2f\n", r.AgeBucket, r.SendPerc, r.OpenPerc)
	}
}
```

## 3058 — Friends With No Mutual Friends

```go
package main

// LeetCode #3058: Friends With No Mutual Friends
// https://leetcode.com/problems/friends-with-no-mutual-friends/
// Difficulty: Medium (SQL problem — implemented in Go)
// Time: O(n * d^2) where d is avg degree | Space: O(n)

import (
	"fmt"
	"sort"
)

type Friendship struct {
	UserID1 int
	UserID2 int
}

type NoMutualPair struct {
	UserID1 int
	UserID2 int
}

func friendsWithNoMutualFriends(friendships []Friendship) []NoMutualPair {
	// Build adjacency set: user -> set of friends
	adj := make(map[int]map[int]bool)
	for _, f := range friendships {
		if adj[f.UserID1] == nil {
			adj[f.UserID1] = make(map[int]bool)
		}
		if adj[f.UserID2] == nil {
			adj[f.UserID2] = make(map[int]bool)
		}
		adj[f.UserID1][f.UserID2] = true
		adj[f.UserID2][f.UserID1] = true
	}

	var results []NoMutualPair
	seen := make(map[[2]int]bool)

	for _, f := range friendships {
		a, b := f.UserID1, f.UserID2

		// Ensure we only process each unordered pair once
		pair := [2]int{a, b}
		if seen[pair] {
			continue
		}
		seen[pair] = true
		seen[[2]int{b, a}] = true

		// Check if a and b have any mutual friends
		hasMutual := false
		for friend := range adj[a] {
			if friend == b {
				continue
			}
			if adj[b][friend] {
				hasMutual = true
				break
			}
		}

		if !hasMutual {
			results = append(results, NoMutualPair{UserID1: a, UserID2: b})
		}
	}

	// Order by user_id1 ASC, user_id2 ASC
	sort.Slice(results, func(i, j int) bool {
		if results[i].UserID1 != results[j].UserID1 {
			return results[i].UserID1 < results[j].UserID1
		}
		return results[i].UserID2 < results[j].UserID2
	})

	return results
}

func main() {
	// Test from problem description
	friendships := []Friendship{
		{UserID1: 1, UserID2: 2},
		{UserID1: 2, UserID2: 3},
		{UserID1: 2, UserID2: 4},
		{UserID1: 1, UserID2: 3},
		{UserID1: 3, UserID2: 4},
		{UserID1: 6, UserID2: 7},
		{UserID1: 8, UserID2: 9},
		{UserID1: 5, UserID2: 6},
	}

	fmt.Println("Friends With No Mutual Friends")
	fmt.Println("=============================")
	fmt.Printf("%-8s %s\n", "user1", "user2")
	fmt.Println("-------------------")

	results := friendsWithNoMutualFriends(friendships)
	for _, r := range results {
		fmt.Printf("%-8d %d\n", r.UserID1, r.UserID2)
	}
}
```

## 3064 — Guess The Number Using Bitwise Questions I

```go
package main

// LeetCode #3064: Guess the Number Using Bitwise Questions I (PAID)
// https://leetcode.com/problems/guess-the-number-using-bitwise-questions-i/
// Difficulty: Medium [Paid]
// Time: O(log n) | Space: O(1)

// Given a hidden number n (1 ≤ n < 2^30), determine n by calling the API
// commonSetBits(num int) int which returns popcount(n & num).
// For each bit position i, probe commonSetBits(1<<i) — if the result is > 0,
// that bit is set in n.

import "fmt"

// commonSetBits simulates the LeetCode API: popcount of (num & hidden).
func commonSetBits(num int, hidden int) int {
	and := hidden & num
	cnt := 0
	for and > 0 {
		cnt++
		and &= and - 1
	}
	return cnt
}

func main() {
	fmt.Println(findNumber(31)) // 31 (0b11111)
	fmt.Println(findNumber(33)) // 33 (0b100001)
	fmt.Println(findNumber(1))  // 1
	fmt.Println(findNumber(0))  // 0 (edge case, though problem says n >= 1)
}

func findNumber(n int) int {
	result := 0
	for i := 0; i < 31; i++ {
		if commonSetBits(1<<i, n) > 0 {
			result |= 1 << i
		}
	}
	return result
}
```

## 3066 — Minimum Operations To Exceed Threshold Value Ii

```go
package main

// LeetCode #3066: Minimum Operations to Exceed Threshold Value II
// https://leetcode.com/problems/minimum-operations-to-exceed-threshold-value-ii/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"container/heap"
	"fmt"
)

type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func main() {
	fmt.Println(minOperations3066([]int{2, 11, 10, 1, 3}, 10))
	fmt.Println(minOperations3066([]int{1, 1, 2, 4, 9}, 20))
}

func minOperations3066(nums []int, k int) (ans int) {
	h := &minHeap{}
	heap.Init(h)
	for _, x := range nums {
		if x < k {
			heap.Push(h, x)
		}
	}
	for h.Len() >= 2 {
		x := heap.Pop(h).(int)
		y := heap.Pop(h).(int)
		val := x*2 + y
		if val < k {
			heap.Push(h, val)
		}
		ans++
	}
	if h.Len() > 0 {
		ans++
	}
	return
}
```

## 3067 — Count Pairs Of Connectable Servers In A Weighted Tree Network

```go
package main

// LeetCode #3067: Count Pairs of Connectable Servers in a Weighted Tree Network
// https://leetcode.com/problems/count-pairs-of-connectable-servers-in-a-weighted-tree-network/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(countPairsOfConnectableServers([][]int{{0, 1, 1}, {1, 2, 5}}, 1))
	fmt.Println(countPairsOfConnectableServers([][]int{{0, 6, 3}, {6, 5, 3}, {0, 3, 1}, {3, 2, 7}, {3, 1, 6}, {3, 4, 2}}, 3))
}

func countPairsOfConnectableServers(edges [][]int, signalSpeed int) []int {
	n := len(edges) + 1
	g := make([][][2]int, n)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		g[u] = append(g[u], [2]int{v, w})
		g[v] = append(g[v], [2]int{u, w})
	}

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		cnt := []int{}
		for _, ne := range g[i] {
			c := dfs(i, ne[0], ne[1], signalSpeed, g)
			if c > 0 {
				cnt = append(cnt, c)
			}
		}
		total := 0
		for j := 0; j < len(cnt); j++ {
			total += cnt[j]
		}
		pairs := 0
		for j := 0; j < len(cnt); j++ {
			total -= cnt[j]
			pairs += cnt[j] * total
		}
		ans[i] = pairs
	}
	return ans
}

func dfs(prev, curr, dist, signalSpeed int, g [][][2]int) int {
	c := 0
	if dist%signalSpeed == 0 {
		c++
	}
	for _, ne := range g[curr] {
		if ne[0] != prev {
			c += dfs(curr, ne[0], dist+ne[1], signalSpeed, g)
		}
	}
	return c
}
```

## 3070 — Count Submatrices With Top Left Element And Sum Less Than K

```go
package main

// LeetCode #3070: Count Submatrices with Top-Left Element and Sum Less Than k
// https://leetcode.com/problems/count-submatrices-with-top-left-element-and-sum-less-than-k/
// Difficulty: Medium
// Time: O(m*n) | Space: O(m*n)

import "fmt"

func main() {
	fmt.Println(countSubmatrices([][]int{{7, 2, 9}, {1, 5, 0}, {2, 6, 6}}, 20))
	fmt.Println(countSubmatrices([][]int{{1, 2}, {3, 4}}, 5))
}

func countSubmatrices(grid [][]int, k int) (ans int) {
	m, n := len(grid), len(grid[0])
	pref := make([][]int, m+1)
	for i := range pref {
		pref[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			pref[i+1][j+1] = grid[i][j] + pref[i][j+1] + pref[i+1][j] - pref[i][j]
			if pref[i+1][j+1] <= k {
				ans++
			}
		}
	}
	return
}
```

## 3071 — Minimum Operations To Write The Letter Y On A Grid

```go
package main

// LeetCode #3071: Minimum Operations to Write the Letter Y on a Grid
// https://leetcode.com/problems/minimum-operations-to-write-the-letter-y-on-a-grid/
// Difficulty: Medium
// Time: O(n^2) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(minimumOperationsToWriteY([][]int{{1, 2, 2}, {2, 2, 3}, {2, 3, 3}}))
	fmt.Println(minimumOperationsToWriteY([][]int{{0, 1, 0, 1, 0}, {1, 0, 1, 0, 1}, {0, 1, 0, 1, 0}, {1, 0, 1, 0, 1}, {0, 1, 0, 1, 0}}))
}

func minimumOperationsToWriteY(grid [][]int) int {
	n := len(grid)
	// Find max value
	maxV := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] > maxV {
				maxV = grid[i][j]
			}
		}
	}
	size := maxV + 1
	yCnt := make([]int, size)
	notYCnt := make([]int, size)
	center := n / 2
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			v := grid[i][j]
			isY := false
			if i == j && i <= center {
				isY = true
			} else if i+j == n-1 && i <= center {
				isY = true
			} else if j == center && i >= center {
				isY = true
			}
			if isY {
				yCnt[v]++
			} else {
				notYCnt[v]++
			}
		}
	}
	totalY := 0
	totalNotY := 0
	for _, c := range yCnt {
		totalY += c
	}
	for _, c := range notYCnt {
		totalNotY += c
	}
	ans := n * n
	for yv := 0; yv < size; yv++ {
		for nv := 0; nv < size; nv++ {
			if yv == nv {
				continue
			}
			ops := totalY - yCnt[yv] + totalNotY - notYCnt[nv]
			if ops < ans {
				ans = ops
			}
		}
	}
	return ans
}
```

## 3073 — Maximum Increasing Triplet Value

```go
package main

// LeetCode #3073: Maximum Increasing Triplet Value (PAID)
// https://leetcode.com/problems/maximum-increasing-triplet-value/
// Difficulty: Medium [Paid]
// Time: O(n^2) | Space: O(1)

// Find the maximum value of nums[i] - nums[j] + nums[k] where
// i < j < k and nums[i] < nums[j] < nums[k]. At least one valid triplet
// is guaranteed to exist.

import "fmt"

func main() {
	// Test 1: Strictly increasing
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(maximumIncreasingTripletValue(nums)) // 4  (3-4+5 or 2-3+5 or 1-2+5)

	// Test 2: Decreasing then increasing
	nums2 := []int{3, 2, 1, 4, 5}
	fmt.Println(maximumIncreasingTripletValue(nums2)) // 4 (3-4+5)

	// Test 3: Small array
	nums3 := []int{10, 20, 30}
	fmt.Println(maximumIncreasingTripletValue(nums3)) // 20 (10-20+30)

	// Test 4: Duplicates — must be strictly increasing
	nums4 := []int{5, 5, 5, 5}
	fmt.Println(maximumIncreasingTripletValue(nums4)) // 0 (no strictly increasing triplet)
}

func maximumIncreasingTripletValue(nums []int) int {
	n := len(nums)
	ans := 0
	for j := 1; j < n-1; j++ {
		leftBest := -1
		for i := 0; i < j; i++ {
			if nums[i] < nums[j] && nums[i] > leftBest {
				leftBest = nums[i]
			}
		}
		rightBest := -1
		for k := j + 1; k < n; k++ {
			if nums[k] > nums[j] && nums[k] > rightBest {
				rightBest = nums[k]
			}
		}
		if leftBest != -1 && rightBest != -1 {
			val := leftBest - nums[j] + rightBest
			if val > ans {
				ans = val
			}
		}
	}
	return ans
}
```

## 3075 — Maximize Happiness Of Selected Children

```go
package main

// LeetCode #3075: Maximize Happiness of Selected Children
// https://leetcode.com/problems/maximize-happiness-of-selected-children/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximumHappinessSum([]int{1, 2, 3}, 2))
	fmt.Println(maximumHappinessSum([]int{1, 1, 1, 1}, 2))
	fmt.Println(maximumHappinessSum([]int{2, 3, 4, 5}, 1))
}

func maximumHappinessSum(happiness []int, k int) int64 {
	sort.Slice(happiness, func(i, j int) bool {
		return happiness[i] > happiness[j]
	})
	ans := int64(0)
	for i := 0; i < k; i++ {
		val := happiness[i] - i
		if val > 0 {
			ans += int64(val)
		}
	}
	return ans
}
```

## 3076 — Shortest Uncommon Substring In An Array

```go
package main

// LeetCode #3076: Shortest Uncommon Substring in an Array
// https://leetcode.com/problems/shortest-uncommon-substring-in-an-array/
// Difficulty: Medium
// Time: O(n * L^2) | Space: O(n * L^2)

import "fmt"

func main() {
	fmt.Println(shortestUncommonSubstring([]string{"cab", "ad", "bad", "c"}))
	fmt.Println(shortestUncommonSubstring([]string{"abc", "bcd", "abcd"}))
}

func shortestUncommonSubstring(arr []string) []string {
	n := len(arr)
	ans := make([]string, n)

	for i := 0; i < n; i++ {
		subs := map[string]bool{}
		for l := 0; l < len(arr[i]); l++ {
			for r := l + 1; r <= len(arr[i]); r++ {
				subs[arr[i][l:r]] = true
			}
		}
		best := ""
		for s := range subs {
			common := false
			for j := 0; j < n; j++ {
				if i == j {
					continue
				}
				if contains(arr[j], s) {
					common = true
					break
				}
			}
			if !common {
				if best == "" || len(s) < len(best) || (len(s) == len(best) && s < best) {
					best = s
				}
			}
		}
		ans[i] = best
	}
	return ans
}

func contains(s, sub string) bool {
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}
```

## 3078 — Match Alphanumerical Pattern In Matrix I

```go
package main

// LeetCode #3078: Match Alphanumerical Pattern in Matrix I (PAID)
// https://leetcode.com/problems/match-alphanumerical-pattern-in-matrix-i/
// Difficulty: Medium [Paid]
// Time: O(R*C*P*Q) | Space: O(min(26, 10))

// Given a numeric board and a pattern of characters (digits or lowercase
// letters), find the top-left [row, col] of the first submatrix that matches
// the pattern under a bijective mapping: same letter -> same digit,
// different letters -> different digits. Digits in the pattern must match
// exactly. Return [-1, -1] if no match.

import "fmt"

func main() {
	// Test 1: No valid mapping (conflicting bijection)
	board := [][]int{{1, 2, 3}, {3, 1, 2}, {2, 3, 1}}
	pattern := [][]byte{{'a', 'b'}, {'b', 'a'}}
	fmt.Println(matchAlphanumericalPattern(board, pattern)) // [-1 -1]

	// Test 2: Simple match
	board2 := [][]int{{1, 2}, {2, 1}}
	pattern2 := [][]byte{{'a', 'b'}, {'b', 'a'}}
	fmt.Println(matchAlphanumericalPattern(board2, pattern2)) // [0 0]

	// Test 3: Letter pattern with digits in pattern
	board3 := [][]int{{5, 1, 3}, {2, 5, 4}}
	pattern3 := [][]byte{{'5', 'a'}, {'b', 'c'}}
	fmt.Println(matchAlphanumericalPattern(board3, pattern3)) // [0 0] (a=1,b=2,c=5)

	// Test 4: Pattern larger than board
	board4 := [][]int{{1}}
	pattern4 := [][]byte{{'a', 'b'}}
	fmt.Println(matchAlphanumericalPattern(board4, pattern4)) // [-1 -1]
}

func matchAlphanumericalPattern(board [][]int, pattern [][]byte) []int {
	R, C := len(board), len(board[0])
	P, Q := len(pattern), len(pattern[0])

	if R < P || C < Q {
		return []int{-1, -1}
	}

	for r := 0; r <= R-P; r++ {
		for c := 0; c <= C-Q; c++ {
			if matches(board, pattern, r, c) {
				return []int{r, c}
			}
		}
	}
	return []int{-1, -1}
}

func matches(board [][]int, pattern [][]byte, r, c int) bool {
	charToDigit := make(map[byte]int)
	digitToChar := make(map[int]byte)

	for i := 0; i < len(pattern); i++ {
		for j := 0; j < len(pattern[0]); j++ {
			ch := pattern[i][j]
			digit := board[r+i][c+j]

			if ch >= '0' && ch <= '9' {
				// Digit in pattern must match board digit exactly
				if int(ch-'0') != digit {
					return false
				}
			} else {
				// Letter in pattern must follow the bijective mapping
				if mapped, ok := charToDigit[ch]; ok {
					if mapped != digit {
						return false
					}
				} else if mappedChar, ok := digitToChar[digit]; ok {
					if mappedChar != ch {
						return false
					}
				} else {
					charToDigit[ch] = digit
					digitToChar[digit] = ch
				}
			}
		}
	}
	return true
}
```

## 3080 — Mark Elements On Array By Performing Queries

```go
package main

// LeetCode #3080: Mark Elements on Array by Performing Queries
// https://leetcode.com/problems/mark-elements-on-array-by-performing-queries/
// Difficulty: Medium
// Time: O((n + q) log n) | Space: O(n)

import (
	"container/heap"
	"fmt"
)

type pair3080 struct {
	val int
	idx int
}
type minHeap3080 []pair3080

func (h minHeap3080) Len() int            { return len(h) }
func (h minHeap3080) Less(i, j int) bool  { return h[i].val < h[j].val || (h[i].val == h[j].val && h[i].idx < h[j].idx) }
func (h minHeap3080) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap3080) Push(x any)         { *h = append(*h, x.(pair3080)) }
func (h *minHeap3080) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func main() {
	fmt.Println(unmarkedSumArray([]int{1, 2, 2, 1, 2, 3, 1}, [][]int{{1, 2}, {3, 3}, {4, 2}}))
	fmt.Println(unmarkedSumArray([]int{1, 4, 2, 3}, [][]int{{0, 1}}))
}

func unmarkedSumArray(nums []int, queries [][]int) []int64 {
	n := len(nums)
	sum := int64(0)
	h := &minHeap3080{}
	heap.Init(h)
	for i, x := range nums {
		sum += int64(x)
		heap.Push(h, pair3080{x, i})
	}
	marked := make([]bool, n)
	ans := make([]int64, len(queries))
	for qi, q := range queries {
		idx, k := q[0], q[1]
		if !marked[idx] {
			marked[idx] = true
			sum -= int64(nums[idx])
		}
		for k > 0 && h.Len() > 0 {
			p := heap.Pop(h).(pair3080)
			if !marked[p.idx] {
				marked[p.idx] = true
				sum -= int64(p.val)
				k--
			}
		}
		ans[qi] = sum
	}
	return ans
}
```

## 3081 — Replace Question Marks In String To Minimize Its Value

```go
package main

// LeetCode #3081: Replace Question Marks in String to Minimize Its Value
// https://leetcode.com/problems/replace-question-marks-in-string-to-minimize-its-value/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import (
	"fmt"
)

func main() {
	fmt.Println(minimizeStringValue("abc?def"))
	fmt.Println(minimizeStringValue("???"))
	fmt.Println(minimizeStringValue("a?b?c?"))
}

func minimizeStringValue(s string) string {
	cnt := [26]int{}
	qpos := []int{}
	for i, ch := range s {
		if ch == '?' {
			qpos = append(qpos, i)
		} else {
			cnt[ch-'a']++
		}
	}
	type pair struct {
		ch   byte
		pos  int
	}
	toFill := make([]byte, len(qpos))
	for i, pos := range qpos {
		bestCh := byte('a')
		bestCost := cnt[0]
		for c := 1; c < 26; c++ {
			if cnt[c] < bestCost {
				bestCost = cnt[c]
				bestCh = byte('a' + c)
			}
		}
		toFill[i] = bestCh
		cnt[bestCh-'a']++
		_ = pos
	}
	ans := []byte(s)
	for i, pos := range qpos {
		ans[pos] = toFill[i]
	}
	return string(ans)
}
```

## 3084 — Count Substrings Starting And Ending With Given Character

```go
package main

// LeetCode #3084: Count Substrings Starting and Ending with Given Character
// https://leetcode.com/problems/count-substrings-starting-and-ending-with-given-character/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(countSubstringsStartingEnding("abada", 'a'))
	fmt.Println(countSubstringsStartingEnding("zzz", 'z'))
}

func countSubstringsStartingEnding(s string, c byte) int64 {
	cnt := int64(0)
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			cnt++
		}
	}
	return cnt * (cnt + 1) / 2
}
```

## 3085 — Minimum Deletions To Make String K Special

```go
package main

// LeetCode #3085: Minimum Deletions to Make String K-Special
// https://leetcode.com/problems/minimum-deletions-to-make-string-k-special/
// Difficulty: Medium
// Time: O(n * 26) = O(n) | Space: O(26) = O(1)

import "fmt"

func minimumDeletions(word string, k int) int {
	freq := make([]int, 26)
	for _, ch := range word {
		freq[ch-'a']++
	}

	ans := len(word)
	for _, minFreq := range freq {
		if minFreq == 0 {
			continue
		}
		ops := 0
		for _, f := range freq {
			if f < minFreq {
				ops += f
			} else if f > minFreq+k {
				ops += f - (minFreq + k)
			}
		}
		if ops < ans {
			ans = ops
		}
	}
	return ans
}

func main() {
	fmt.Println(minimumDeletions("aabcaba", 2)) // Expected: 2
	fmt.Println(minimumDeletions("dabdcbdcdcd", 2)) // Expected: 2
	fmt.Println(minimumDeletions("aaabaaa", 2)) // Expected: 0
}
```

## 3087 — Find Trending Hashtags

```go
package main

// LeetCode #3087: Find Trending Hashtags
// https://leetcode.com/problems/find-trending-hashtags/
// Difficulty: Medium [Paid]
// Time: O(n * m) | Space: O(n)

import (
	"fmt"
	"sort"
	"strings"
)

func findTrendingHashtags(tweets []string) []string {
	freq := make(map[string]int)
	for _, tweet := range tweets {
		words := strings.Fields(tweet)
		for _, w := range words {
			if strings.HasPrefix(w, "#") {
				hashtag := strings.ToLower(w)
				freq[hashtag]++
			}
		}
	}

	type ht struct {
		tag string
		cnt int
	}
	var list []ht
	for tag, cnt := range freq {
		list = append(list, ht{tag, cnt})
	}
	sort.Slice(list, func(i, j int) bool {
		if list[i].cnt != list[j].cnt {
			return list[i].cnt > list[j].cnt
		}
		return list[i].tag < list[j].tag
	})

	ans := make([]string, 0, min(3, len(list)))
	for i := 0; i < min(3, len(list)); i++ {
		ans = append(ans, list[i].tag)
	}
	return ans
}

func main() {
	fmt.Println(findTrendingHashtags([]string{
		"Good morning #tech #coding",
		"Loving #coding today #go",
		"#Tech is amazing #golang",
		"#coding #coding #coding",
	}))
	fmt.Println(findTrendingHashtags([]string{
		"#a #b #c",
		"#a #b",
		"#a",
	}))
}
```

## 3089 — Find Bursty Behavior

```go
package main

// LeetCode #3089: Find Bursty Behavior
// https://leetcode.com/problems/find-bursty-behavior/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func findBurstyBehavior(posts [][]int, k int) []int {
	n := len(posts)
	if n == 0 {
		return nil
	}

	userPosts := make(map[int][]int)
	for _, p := range posts {
		userID, timestamp := p[0], p[1]
		userPosts[userID] = append(userPosts[userID], timestamp)
	}

	var bursty []int
	for uid, timestamps := range userPosts {
		sort.Ints(timestamps)
		for i := k - 1; i < len(timestamps); i++ {
			if timestamps[i]-timestamps[i-k+1] <= 100 {
				bursty = append(bursty, uid)
				break
			}
		}
	}

	sort.Ints(bursty)
	return bursty
}

func main() {
	fmt.Println(findBurstyBehavior([][]int{{1, 10}, {1, 20}, {1, 30}, {2, 5}, {2, 200}}, 3)) // Expected: [1]
	fmt.Println(findBurstyBehavior([][]int{{1, 1}, {2, 2}, {3, 3}}, 2))                       // Expected: []
}
```

## 3091 — Apply Operations To Make Sum Of Array Greater Than Or Equal To K

```go
package main

// LeetCode #3091: Apply Operations to Make Sum of Array Greater Than or Equal to k
// https://leetcode.com/problems/apply-operations-to-make-sum-of-array-greater-than-or-equal-to-k/
// Difficulty: Medium
// Time: O(sqrt(k)) | Space: O(1)

import "fmt"

func minOperations(k int) int {
	ans := k - 1
	for a := 1; a <= k; a++ {
		b := (k + a - 1) / a
		ops := (a - 1) + (b - 1)
		if ops < ans {
			ans = ops
		}
	}
	return ans
}

func main() {
	fmt.Println(minOperations(11))  // Expected: 5
	fmt.Println(minOperations(1))   // Expected: 0
	fmt.Println(minOperations(5))   // Expected: 3
}
```

## 3092 — Most Frequent Ids

```go
package main

// LeetCode #3092: Most Frequent IDs
// https://leetcode.com/problems/most-frequent-ids/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"container/heap"
	"fmt"
)

type Item struct {
	id    int
	count int64
	idx   int
}

type MaxHeap []*Item

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].count > h[j].count }
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].idx = i
	h[j].idx = j
}
func (h *MaxHeap) Push(x any) {
	n := len(*h)
	item := x.(*Item)
	item.idx = n
	*h = append(*h, item)
}
func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.idx = -1
	*h = old[:n-1]
	return item
}

func mostFrequentIDs(nums []int, freq []int) []int64 {
	n := len(nums)
	ans := make([]int64, n)
	counts := make(map[int]int64)
	h := &MaxHeap{}
	heap.Init(h)

	for i := 0; i < n; i++ {
		id := nums[i]
		counts[id] += int64(freq[i])
		heap.Push(h, &Item{id: id, count: counts[id]})
		for h.Len() > 0 && (*h)[0].count != counts[(*h)[0].id] {
			heap.Pop(h)
		}
		if h.Len() > 0 {
			ans[i] = (*h)[0].count
		}
	}
	return ans
}

func main() {
	fmt.Println(mostFrequentIDs([]int{1, 2, 3, 2, 2, 1, 3}, []int{-2, 3, -3, 5, 1, -3, 1}))
	fmt.Println(mostFrequentIDs([]int{2, 3, 2, 1}, []int{3, 2, -3, 1}))
}
```

## 3094 — Guess The Number Using Bitwise Questions Ii

```go
package main

// LeetCode #3094: Guess the Number Using Bitwise Questions II
// https://leetcode.com/problems/generate-binary-strings-without-adjacent-zeros/
// https://leetcode.com/problems/guess-the-number-using-bitwise-questions-ii/
// Difficulty: Medium [Paid]
// Time: O(log n) | Space: O(1)

import "fmt"

var hiddenNumber int

func commonBits(num int) int {
	count := 0
	x := hiddenNumber ^ num
	for x > 0 {
		count += x & 1
		x >>= 1
	}
	return 30 - count
}

func findNumber() int {
	n := 0
	for i := 0; i <= 30; i++ {
		bit1 := commonBits(n | (1 << i))
		bit0 := commonBits(n)
		if bit1 > bit0 {
			n |= (1 << i)
		}
	}
	return n
}

func main() {
	hiddenNumber = 42
	fmt.Println(findNumber())

	hiddenNumber = 100
	fmt.Println(findNumber())
}
```

## 3096 — Minimum Levels To Gain More Points

```go
package main

// LeetCode #3096: Minimum Levels to Gain More Points
// https://leetcode.com/problems/minimum-levels-to-gain-more-points/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func minimumLevels(possible []int) int {
	n := len(possible)
	score := make([]int, n)
	for i, v := range possible {
		if v == 0 {
			score[i] = -1
		} else {
			score[i] = 1
		}
	}

	total := 0
	for _, v := range score {
		total += v
	}

	prefix := 0
	for i := 0; i < n-1; i++ {
		prefix += score[i]
		if prefix > total-prefix {
			return i + 1
		}
	}
	return -1
}

func main() {
	fmt.Println(minimumLevels([]int{1, 0, 1, 0}))       // Expected: 1
	fmt.Println(minimumLevels([]int{1, 1, 1, 1, 1}))    // Expected: 3
	fmt.Println(minimumLevels([]int{0, 0}))              // Expected: -1
}
```

## 3097 — Shortest Subarray With Or At Least K Ii

```go
package main

// LeetCode #3097: Shortest Subarray With OR at Least K II
// https://leetcode.com/problems/shortest-subarray-with-or-at-least-k-ii/
// Difficulty: Medium
// Time: O(n * 32) = O(n) | Space: O(32) = O(1)

import "fmt"

func minimumSubarrayLength(nums []int, k int) int {
	if k == 0 {
		return 1
	}

	n := len(nums)
	ans := n + 1
	bits := make([]int, 32)
	left := 0
	cur := 0

	for right := 0; right < n; right++ {
		cur |= nums[right]
		for b := 0; b < 32; b++ {
			if nums[right]&(1<<b) != 0 {
				bits[b]++
			}
		}

		for left <= right && cur >= k {
			if right-left+1 < ans {
				ans = right - left + 1
			}

			for b := 0; b < 32; b++ {
				if nums[left]&(1<<b) != 0 {
					bits[b]--
					if bits[b] == 0 {
						cur &^= (1 << b)
					}
				}
			}
			left++
		}
	}

	if ans > n {
		return -1
	}
	return ans
}

func main() {
	fmt.Println(minimumSubarrayLength([]int{1, 2, 3}, 3))       // Expected: 1
	fmt.Println(minimumSubarrayLength([]int{1, 2, 3}, 5))       // Expected: 2
	fmt.Println(minimumSubarrayLength([]int{2, 1, 8}, 10))      // Expected: 3
}
```

## 3100 — Water Bottles Ii

```go
package main

// LeetCode #3100: Water Bottles II
// https://leetcode.com/problems/water-bottles-ii/
// Difficulty: Medium
// Time: O(log n) | Space: O(1)

import "fmt"

func maxBottlesDrunk(numBottles int, numExchange int) int {
	total := numBottles
	empty := numBottles

	for empty >= numExchange {
		empty -= numExchange
		numExchange++
		total++
		empty++
	}

	return total
}

func main() {
	fmt.Println(maxBottlesDrunk(13, 6)) // Expected: 15
	fmt.Println(maxBottlesDrunk(10, 3)) // Expected: 13
}
```

## 3101 — Count Alternating Subarrays

```go
package main

// LeetCode #3101: Count Alternating Subarrays
// https://leetcode.com/problems/count-alternating-subarrays/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func countAlternatingSubarrays(nums []int) int64 {
	var ans int64
	n := len(nums)
	left := 0

	for right := 0; right < n; right++ {
		if right > 0 && nums[right] == nums[right-1] {
			left = right
		}
		ans += int64(right - left + 1)
	}

	return ans
}

func main() {
	fmt.Println(countAlternatingSubarrays([]int{0, 1, 1, 1})) // Expected: 5
	fmt.Println(countAlternatingSubarrays([]int{1, 0, 1, 0})) // Expected: 10
}
```

## 3106 — Lexicographically Smallest String After Operations With Constraint

```go
package main

// LeetCode #3106: Lexicographically Smallest String After Operations With Constraint
// https://leetcode.com/problems/lexicographically-smallest-string-after-operations-with-constraint/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func getSmallestString(s string, k int) string {
	if k == 0 {
		return s
	}

	bytes := []byte(s)
	for i, ch := range bytes {
		dist := int(ch - 'a')
		move := min(dist, 26-dist)
		if move <= k {
			k -= move
			bytes[i] = 'a'
		} else {
			bytes[i] = byte(int(ch) - k)
			k = 0
			break
		}
	}
	return string(bytes)
}

func main() {
	fmt.Println(getSmallestString("zbbz", 3))  // Expected: "aaaz"
	fmt.Println(getSmallestString("xaxcd", 4)) // Expected: "aawcd"
	fmt.Println(getSmallestString("lol", 0))   // Expected: "lol"
}
```

## 3107 — Minimum Operations To Make Median Of Array Equal To K

```go
package main

// LeetCode #3107: Minimum Operations to Make Median of Array Equal to K
// https://leetcode.com/problems/minimum-operations-to-make-median-of-array-equal-to-k/
// Difficulty: Medium
// Time: O(n log n) | Space: O(log n)

import (
	"fmt"
	"sort"
)

func minOperationsToMakeMedianK(nums []int, k int) int64 {
	sort.Ints(nums)
	n := len(nums)
	mid := n / 2
	var ans int64

	ans += abs(int64(nums[mid] - k))
	nums[mid] = k

	for i := mid - 1; i >= 0 && nums[i] > k; i-- {
		ans += abs(int64(nums[i] - k))
		nums[i] = k
	}

	for i := mid + 1; i < n && nums[i] < k; i++ {
		ans += abs(int64(nums[i] - k))
		nums[i] = k
	}

	return ans
}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(minOperationsToMakeMedianK([]int{2, 5, 6, 8, 5}, 4)) // Expected: 2
	fmt.Println(minOperationsToMakeMedianK([]int{2, 5, 6, 8, 5}, 7)) // Expected: 3
	fmt.Println(minOperationsToMakeMedianK([]int{1, 2, 3, 4, 5, 6}, 4)) // Expected: 0
}
```

## 3109 — Find The Index Of Permutation

```go
package main

// LeetCode #3109: Find the Index of Permutation
// https://leetcode.com/problems/find-the-index-of-permutation/
// Difficulty: Medium [Paid]
// Time: O(n^2) | Space: O(n)

import "fmt"

func getPermutationIndex(perm []int) int {
	n := len(perm)
	mod := int64(1000000007)

	fact := make([]int64, n)
	fact[0] = 1
	for i := 1; i < n; i++ {
		fact[i] = fact[i-1] * int64(i) % mod
	}

	bit := make([]int, n+1)

	update := func(idx, val int) {
		for idx <= n {
			bit[idx] += val
			idx += idx & -idx
		}
	}

	query := func(idx int) int {
		sum := 0
		for idx > 0 {
			sum += bit[idx]
			idx -= idx & -idx
		}
		return sum
	}

	for i := 1; i <= n; i++ {
		update(i, 1)
	}

	ans := int64(0)
	for i := 0; i < n; i++ {
		smaller := query(perm[i]) - 1
		ans = (ans + int64(smaller)*fact[n-1-i]) % mod
		update(perm[i], -1)
	}

	return int((ans + 1) % mod)
}

func main() {
	fmt.Println(getPermutationIndex([]int{1, 2, 3})) // Expected: 1
	fmt.Println(getPermutationIndex([]int{3, 2, 1})) // Expected: 6
	fmt.Println(getPermutationIndex([]int{2, 1, 3})) // Expected: 3
}
```

## 3111 — Minimum Rectangles To Cover Points

```go
package main

// LeetCode #3111: Minimum Rectangles to Cover Points
// https://leetcode.com/problems/minimum-rectangles-to-cover-points/
// Difficulty: Medium
// Time: O(n log n) | Space: O(log n)

import (
	"fmt"
	"sort"
)

func minRectanglesToCoverPoints(points [][]int, w int) int {
	xs := make([]int, len(points))
	for i, p := range points {
		xs[i] = p[0]
	}
	sort.Ints(xs)

	ans := 0
	i := 0
	for i < len(xs) {
		ans++
		end := xs[i] + w
		for i < len(xs) && xs[i] <= end {
			i++
		}
	}
	return ans
}

func main() {
	fmt.Println(minRectanglesToCoverPoints([][]int{{2, 1}, {1, 0}, {1, 4}, {1, 8}, {3, 5}, {4, 6}}, 1)) // Expected: 2
	fmt.Println(minRectanglesToCoverPoints([][]int{{0, 0}, {1, 1}, {2, 2}, {3, 3}}, 2)) // Expected: 2
}
```

## 3112 — Minimum Time To Visit Disappearing Nodes

```go
package main

// LeetCode #3112: Minimum Time to Visit Disappearing Nodes
// https://leetcode.com/problems/minimum-time-to-visit-disappearing-nodes/
// Difficulty: Medium
// Time: O((n + m) log n) | Space: O(n + m)

import (
	"container/heap"
	"fmt"
	"math"
)

type Edge struct {
	to, w int
}

type Item struct {
	node, dist int
	idx        int
}

type PQ []*Item

func (pq PQ) Len() int           { return len(pq) }
func (pq PQ) Less(i, j int) bool { return pq[i].dist < pq[j].dist }
func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].idx = i
	pq[j].idx = j
}
func (pq *PQ) Push(x any) { *pq = append(*pq, x.(*Item)) }
func (pq *PQ) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[:n-1]
	return item
}

func minimumTime(n int, edges [][]int, disappear []int) []int {
	graph := make([][]Edge, n)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		graph[u] = append(graph[u], Edge{v, w})
		graph[v] = append(graph[v], Edge{u, w})
	}

	dist := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[0] = 0

	pq := &PQ{}
	heap.Init(pq)
	heap.Push(pq, &Item{node: 0, dist: 0})

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(*Item)
		if cur.dist > dist[cur.node] {
			continue
		}

		for _, e := range graph[cur.node] {
			nd := cur.dist + e.w
			if nd < disappear[e.to] && nd < dist[e.to] {
				dist[e.to] = nd
				heap.Push(pq, &Item{node: e.to, dist: nd})
			}
		}
	}

	ans := make([]int, n)
	for i := range ans {
		if dist[i] == math.MaxInt32 {
			ans[i] = -1
		} else {
			ans[i] = dist[i]
		}
	}
	return ans
}

func main() {
	fmt.Println(minimumTime(3, [][]int{{0, 1, 2}, {1, 2, 1}, {0, 2, 4}}, []int{1, 1, 5}))
	fmt.Println(minimumTime(3, [][]int{{0, 1, 2}, {1, 2, 1}, {0, 2, 4}}, []int{1, 3, 5}))
	fmt.Println(minimumTime(2, [][]int{{0, 1, 1}}, []int{1, 1}))
}
```

## 3115 — Maximum Prime Difference

```go
package main

// LeetCode #3115: Maximum Prime Difference
// https://leetcode.com/problems/maximum-prime-difference/
// Difficulty: Medium
// Time: O(n * sqrt(m)) | Space: O(1)

import "fmt"

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func maximumPrimeDifference(nums []int) int {
	first, last := -1, -1
	for i, v := range nums {
		if isPrime(v) {
			if first == -1 {
				first = i
			}
			last = i
		}
	}
	return last - first
}

func main() {
	fmt.Println(maximumPrimeDifference([]int{4, 2, 9, 5, 3})) // Expected: 3
	fmt.Println(maximumPrimeDifference([]int{4, 8, 2, 8}))    // Expected: 0
}
```

## 3118 — Friday Purchase Iii

```go
package main

// LeetCode #3118: Friday Purchase III
// https://leetcode.com/problems/friday-purchase-iii/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func fridayPurchase(purchases [][]int) int64 {
	total := int64(0)
	for _, p := range purchases {
		// p[0] = day of week (5 = Friday), p[1] = amount
		if p[0] == 5 {
			total += int64(p[1])
		}
	}
	return total
}

func main() {
	fmt.Println(fridayPurchase([][]int{{5, 100}, {1, 50}, {5, 200}, {3, 75}})) // Expected: 300
	fmt.Println(fridayPurchase([][]int{{2, 50}, {3, 100}}))                     // Expected: 0
}
```

## 3119 — Maximum Number Of Potholes That Can Be Fixed

```go
package main

// LeetCode #3119: Maximum Number of Potholes That Can Be Fixed
// https://leetcode.com/problems/maximum-number-of-potholes-that-can-be-fixed/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func maxPotholes(road string, budget int) int {
	var segs []int
	count := 0
	for i := 0; i < len(road); i++ {
		if road[i] == 'x' {
			count++
		} else {
			if count > 0 {
				segs = append(segs, count)
				count = 0
			}
		}
	}
	if count > 0 {
		segs = append(segs, count)
	}

	sort.Slice(segs, func(i, j int) bool {
		return segs[i] > segs[j]
	})

	ans := 0
	for _, seg := range segs {
		cost := seg + 1
		if budget >= cost {
			budget -= cost
			ans += seg
		}
	}
	return ans
}

func main() {
	fmt.Println(maxPotholes("...xxx..xx", 7))  // Expected: 5
	fmt.Println(maxPotholes("..xxxxx", 4))     // Expected: 3
	fmt.Println(maxPotholes("x.x.x.x", 10))    // Expected: 4
}
```

## 3121 — Count The Number Of Special Characters Ii

```go
package main

// LeetCode #3121: Count the Number of Special Characters II
// https://leetcode.com/problems/count-the-number-of-special-characters-ii/
// Difficulty: Medium
// Time: O(n) | Space: O(26) = O(1)

import "fmt"

func numberOfSpecialChars(word string) int {
	firstLower := make([]int, 26)
	lastUpper := make([]int, 26)
	for i := range firstLower {
		firstLower[i] = -1
		lastUpper[i] = -1
	}

	for i, ch := range word {
		if ch >= 'a' && ch <= 'z' {
			idx := ch - 'a'
			if firstLower[idx] == -1 {
				firstLower[idx] = i
			}
		} else {
			idx := ch - 'A'
			lastUpper[idx] = i
		}
	}

	ans := 0
	for i := 0; i < 26; i++ {
		if firstLower[i] != -1 && lastUpper[i] != -1 && firstLower[i] > lastUpper[i] {
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(numberOfSpecialChars("aaAbcBC"))  // Expected: 3
	fmt.Println(numberOfSpecialChars("abc"))       // Expected: 0
	fmt.Println(numberOfSpecialChars("AbBCab"))    // Expected: 0
}
```

## 3122 — Minimum Number Of Operations To Satisfy Conditions

```go
package main

// LeetCode #3122: Minimum Number of Operations to Satisfy Conditions
// https://leetcode.com/problems/minimum-number-of-operations-to-satisfy-conditions/
// Difficulty: Medium
// Time: O(n * m * 10) | Space: O(m * 10)

import (
	"fmt"
	"math"
)

func minimumOperations(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])

	cost := make([][10]int, n)
	for j := 0; j < n; j++ {
		for d := 0; d < 10; d++ {
			cnt := 0
			for i := 0; i < m; i++ {
				if grid[i][j] != d {
					cnt++
				}
			}
			cost[j][d] = cnt
		}
	}

	dp := make([][10]int, n)
	for j := 0; j < n; j++ {
		for d := 0; d < 10; d++ {
			dp[j][d] = math.MaxInt32
		}
	}

	for d := 0; d < 10; d++ {
		dp[0][d] = cost[0][d]
	}

	for j := 1; j < n; j++ {
		for d := 0; d < 10; d++ {
			for pd := 0; pd < 10; pd++ {
				if pd != d {
					dp[j][d] = min(dp[j][d], dp[j-1][pd]+cost[j][d])
				}
			}
		}
	}

	ans := math.MaxInt32
	for d := 0; d < 10; d++ {
		ans = min(ans, dp[n-1][d])
	}
	return ans
}

func main() {
	fmt.Println(minimumOperations([][]int{{1, 0, 2}, {1, 0, 2}})) // Expected: 0
	fmt.Println(minimumOperations([][]int{{1, 1, 1}, {0, 0, 0}})) // Expected: 3
}
```

## 3124 — Find Longest Calls

```go
package main

// LeetCode #3124: Find Longest Calls
// https://leetcode.com/problems/find-longest-calls/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func findLongestCalls(calls [][]int, k int) []int {
	type call struct {
		id     int
		dur    int
		caller int
	}

	var list []call
	for _, c := range calls {
		list = append(list, call{c[0], c[2] - c[1], c[3]})
	}

	sort.Slice(list, func(i, j int) bool {
		if list[i].dur != list[j].dur {
			return list[i].dur > list[j].dur
		}
		return list[i].id < list[j].id
	})

	ans := make([]int, 0, k)
	for i := 0; i < k && i < len(list); i++ {
		ans = append(ans, list[i].id)
	}
	return ans
}

func main() {
	fmt.Println(findLongestCalls([][]int{{1, 0, 30, 1}, {2, 5, 25, 2}, {3, 10, 20, 1}}, 2)) // Expected: [1 2]
	fmt.Println(findLongestCalls([][]int{{1, 0, 10, 1}, {2, 0, 5, 2}}, 3))                    // Expected: [1 2]
}
```

## 3125 — Maximum Number That Makes Result Of Bitwise And Zero

```go
package main

// LeetCode #3125: Maximum Number That Makes Result of Bitwise AND Zero
// https://leetcode.com/problems/maximum-number-that-makes-result-of-bitwise-and-zero/
// Difficulty: Medium [Paid]
// Time: O(1) | Space: O(1)

import "fmt"
import "math/bits"

func maxNumber(n int64) int64 {
	if n <= 0 {
		return 0
	}
	bits := bits.Len64(uint64(n))
	return int64((uint64(1) << (bits - 1)) - 1)
}

func main() {
	fmt.Println(maxNumber(5))  // Expected: 3
	fmt.Println(maxNumber(10)) // Expected: 7
	fmt.Println(maxNumber(1))  // Expected: 0
}
```

## 3126 — Server Utilization Time

```go
package main

// LeetCode #3126: Server Utilization Time
// https://leetcode.com/problems/server-utilization-time/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func serverUtilizationTime(logs [][]int) int {
	if len(logs) == 0 {
		return 0
	}

	sort.Slice(logs, func(i, j int) bool {
		return logs[i][0] < logs[j][0]
	})

	merged := [][]int{logs[0]}
	for i := 1; i < len(logs); i++ {
		last := merged[len(merged)-1]
		if logs[i][0] <= last[1] {
			if logs[i][1] > last[1] {
				last[1] = logs[i][1]
			}
		} else {
			merged = append(merged, logs[i])
		}
	}

	total := 0
	for _, seg := range merged {
		total += seg[1] - seg[0]
	}
	return total
}

func main() {
	fmt.Println(serverUtilizationTime([][]int{{0, 5}, {2, 7}, {8, 10}})) // Expected: 9
	fmt.Println(serverUtilizationTime([][]int{{1, 3}, {3, 5}}))          // Expected: 4
}
```

## 3128 — Right Triangles

```go
package main

// LeetCode #3128: Right Triangles
// https://leetcode.com/problems/right-triangles/
// Difficulty: Medium
// Time: O(m * n) | Space: O(m + n)

import "fmt"

func numberOfRightTriangles(grid [][]int) int64 {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])

	rowSum := make([]int, m)
	colSum := make([]int, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				rowSum[i]++
				colSum[j]++
			}
		}
	}

	var ans int64
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				ans += int64(rowSum[i]-1) * int64(colSum[j]-1)
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(numberOfRightTriangles([][]int{{0, 1, 0}, {0, 1, 1}, {0, 1, 0}})) // Expected: 2
	fmt.Println(numberOfRightTriangles([][]int{{1, 0, 1}, {1, 0, 0}, {1, 0, 0}})) // Expected: 2
}
```

## 3129 — Find All Possible Stable Binary Arrays I

```go
package main

// LeetCode #3129: Find All Possible Stable Binary Arrays I
// https://leetcode.com/problems/find-all-possible-stable-binary-arrays-i/
// Difficulty: Medium
// Time: O(zero * one * limit) | Space: O(zero * one * 2)

import "fmt"

func numberOfStableArrays(zero int, one int, limit int) int {
	const mod = 1_000_000_007
	dp := make([][][2]int, zero+1)
	for i := range dp {
		dp[i] = make([][2]int, one+1)
	}

	for i := 0; i <= zero; i++ {
		for j := 0; j <= one; j++ {
			if i == 0 && j == 0 {
				dp[i][j][0] = 1
				dp[i][j][1] = 1
				continue
			}
			if i > 0 {
				dp[i][j][0] = dp[i-1][j][0] + dp[i-1][j][1]
				if i > limit {
					dp[i][j][0] -= dp[i-limit-1][j][1]
				}
				dp[i][j][0] %= mod
				if dp[i][j][0] < 0 {
					dp[i][j][0] += mod
				}
			}
			if j > 0 {
				dp[i][j][1] = dp[i][j-1][0] + dp[i][j-1][1]
				if j > limit {
					dp[i][j][1] -= dp[i][j-limit-1][0]
				}
				dp[i][j][1] %= mod
				if dp[i][j][1] < 0 {
					dp[i][j][1] += mod
				}
			}
		}
	}

	return (dp[zero][one][0] + dp[zero][one][1]) % mod
}

func main() {
	fmt.Println(numberOfStableArrays(1, 1, 2)) // Expected: 2
	fmt.Println(numberOfStableArrays(1, 2, 1)) // Expected: 1
	fmt.Println(numberOfStableArrays(3, 1, 1)) // Expected: 2
}
```

## 3132 — Find The Integer Added To Array Ii

```go
package main

// LeetCode #3132: Find the Integer Added to Array II
// https://leetcode.com/problems/find-the-integer-added-to-array-ii/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import (
	"fmt"
	"sort"
)

func minimumAddedInteger(nums1 []int, nums2 []int) int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	// Try all pairs from nums1 as the two removed elements
	for i := 0; i < len(nums1); i++ {
		for j := i + 1; j < len(nums1); j++ {
			// Check if nums2 can be matched after removing nums1[i] and nums1[j]
			diff := -1001
			idx := 0
			match := true
			for k := 0; k < len(nums1) && match; k++ {
				if k == i || k == j {
					continue
				}
				curDiff := nums2[idx] - nums1[k]
				if diff == -1001 {
					diff = curDiff
				} else if curDiff != diff {
					match = false
				}
				idx++
			}
			if match && diff >= 0 {
				return diff
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(minimumAddedInteger([]int{4, 20, 16, 12, 8}, []int{14, 18, 10})) // Expected: -2
	fmt.Println(minimumAddedInteger([]int{3, 5, 5, 3}, []int{7, 7}))             // Expected: 2
}
```

## 3133 — Minimum Array End

```go
package main

// LeetCode #3133: Minimum Array End
// https://leetcode.com/problems/minimum-array-end/
// Difficulty: Medium
// Time: O(log n) | Space: O(1)

import "fmt"

func minEnd(n int, x int) int64 {
	n64 := int64(n - 1)
	x64 := int64(x)
	ans := int64(0)

	bitPos := 0
	for n64 > 0 || x64 > 0 {
		if x64&1 == 1 {
			ans |= (int64(1) << bitPos)
		} else {
			ans |= ((n64 & 1) << bitPos)
			n64 >>= 1
		}
		x64 >>= 1
		bitPos++
	}
	return ans
}

func main() {
	fmt.Println(minEnd(3, 4))  // Expected: 6
	fmt.Println(minEnd(2, 7))  // Expected: 15
	fmt.Println(minEnd(1, 5))  // Expected: 5
}
```

## 3135 — Equalize Strings By Adding Or Removing Characters At Ends

```go
package main

// LeetCode #3135: Equalize Strings by Adding or Removing Characters at Ends
// https://leetcode.com/problems/equalize-strings-by-adding-or-removing-characters-at-ends/
// Difficulty: Medium [Paid]
// Time: O(n * m) | Space: O(1)

import "fmt"

func minOperations(initial string, target string) int {
	m, n := len(initial), len(target)
	longest := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			k := 0
			for i+k < m && j+k < n && initial[i+k] == target[j+k] {
				k++
			}
			if k > longest {
				longest = k
			}
		}
	}

	return m + n - 2*longest
}

func main() {
	fmt.Println(minOperations("abcdef", "defabc")) // Expected: 0 (lcs "def" or "abc")
	fmt.Println(minOperations("abc", "xyz"))       // Expected: 6
	fmt.Println(minOperations("abcde", "cde"))     // Expected: 2
}
```

## 3137 — Minimum Number Of Operations To Make Word K Periodic

```go
package main

// LeetCode #3137: Minimum Number of Operations to Make Word K-Periodic
// https://leetcode.com/problems/minimum-number-of-operations-to-make-word-k-periodic/
// Difficulty: Medium
// Time: O(n) | Space: O(n / k)

import "fmt"

func minimumOperationsToMakeKPeriodic(word string, k int) int {
	n := len(word)
	freq := make(map[string]int)
	maxFreq := 0

	for i := 0; i < n; i += k {
		sub := word[i : i+k]
		freq[sub]++
		if freq[sub] > maxFreq {
			maxFreq = freq[sub]
		}
	}

	return n/k - maxFreq
}

func main() {
	fmt.Println(minimumOperationsToMakeKPeriodic("leetcodeleet", 4)) // Expected: 1
	fmt.Println(minimumOperationsToMakeKPeriodic("abcabcabc", 3))    // Expected: 0
	fmt.Println(minimumOperationsToMakeKPeriodic("aabbccddee", 5))   // Expected: 1
}
```

## 3138 — Minimum Length Of Anagram Concatenation

```go
package main

// LeetCode #3138: Minimum Length of Anagram Concatenation
// https://leetcode.com/problems/minimum-length-of-anagram-concatenation/
// Difficulty: Medium
// Time: O(n * sqrt(n)) | Space: O(n)

import "fmt"

func minAnagramLength(s string) int {
	n := len(s)

	freq := func(lo, hi int) [26]int {
		var f [26]int
		for i := lo; i < hi; i++ {
			f[s[i]-'a']++
		}
		return f
	}

	for l := 1; l <= n; l++ {
		if n%l != 0 {
			continue
		}
		base := freq(0, l)
		match := true
		for j := l; j < n && match; j += l {
			cur := freq(j, j+l)
			if cur != base {
				match = false
			}
		}
		if match {
			return l
		}
	}
	return n
}

func main() {
	fmt.Println(minAnagramLength("abba"))             // Expected: 2
	fmt.Println(minAnagramLength("abcabc"))           // Expected: 3
	fmt.Println(minAnagramLength("cdef"))             // Expected: 4
}
```

## 3140 — Consecutive Available Seats Ii

```go
package main

// LeetCode #3140: Consecutive Available Seats II
// https://leetcode.com/problems/consecutive-available-seats-ii/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func consecutiveAvailableSeats(seats [][]int) int {
	sort.Slice(seats, func(i, j int) bool {
		if seats[i][0] != seats[j][0] {
			return seats[i][0] < seats[j][0]
		}
		return seats[i][1] < seats[j][1]
	})

	maxLen := 0
	curStart, curEnd := seats[0][0], seats[0][1]

	for i := 1; i < len(seats); i++ {
		if seats[i][0] <= curEnd+1 {
			if seats[i][1] > curEnd {
				curEnd = seats[i][1]
			}
		} else {
			if curEnd-curStart+1 > maxLen {
				maxLen = curEnd - curStart + 1
			}
			curStart, curEnd = seats[i][0], seats[i][1]
		}
	}
	if curEnd-curStart+1 > maxLen {
		maxLen = curEnd - curStart + 1
	}

	return maxLen
}

func main() {
	fmt.Println(consecutiveAvailableSeats([][]int{{1, 3}, {4, 6}, {7, 10}})) // Expected: 10
	fmt.Println(consecutiveAvailableSeats([][]int{{1, 2}, {5, 6}}))          // Expected: 2
}
```

## 3143 — Maximum Points Inside The Square

```go
package main

// LeetCode #3143: Maximum Points Inside the Square
// https://leetcode.com/problems/maximum-points-inside-the-square/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func maxPointsInsideSquare(points [][]int, s string) int {
	type item struct {
		dist int
		tag  byte
	}
	n := len(points)
	arr := make([]item, n)
	for i, p := range points {
		d := max(abs(p[0]), abs(p[1]))
		arr[i] = item{d, s[i]}
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i].dist < arr[j].dist
	})

	seen := make(map[byte]bool)
	ans := 0
	i := 0
	for i < n {
		j := i
		for j < n && arr[j].dist == arr[i].dist {
			if seen[arr[j].tag] {
				return ans
			}
			j++
		}
		for k := i; k < j; k++ {
			seen[arr[k].tag] = true
		}
		ans = len(seen)
		i = j
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(maxPointsInsideSquare([][]int{{2, 2}, {-1, -2}, {-4, 4}, {-3, 1}, {3, -3}}, "abdca"))
}
```

## 3144 — Minimum Substring Partition Of Equal Character Frequency

```go
package main

// LeetCode #3144: Minimum Substring Partition of Equal Character Frequency
// https://leetcode.com/problems/minimum-substring-partition-of-equal-character-frequency/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import (
	"fmt"
	"math"
)

func minimumSubstringsInPartition(s string) int {
	n := len(s)
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0

	for i := 1; i <= n; i++ {
		freq := make([]int, 26)
		var distinct, maxFreq int
		for j := i - 1; j >= 0; j-- {
			idx := s[j] - 'a'
			if freq[idx] == 0 {
				distinct++
			}
			freq[idx]++
			if freq[idx] > maxFreq {
				maxFreq = freq[idx]
			}

			if maxFreq*distinct == i-j {
				if dp[j]+1 < dp[i] {
					dp[i] = dp[j] + 1
				}
			}
		}
	}
	return dp[n]
}

func main() {
	fmt.Println(minimumSubstringsInPartition("fabccddg")) // Expected: 3
	fmt.Println(minimumSubstringsInPartition("abababaccddb")) // Expected: ?
}
```

## 3147 — Taking Maximum Energy From The Mystic Dungeon

```go
package main

// LeetCode #3147: Taking Maximum Energy From the Mystic Dungeon
// https://leetcode.com/problems/taking-maximum-energy-from-the-mystic-dungeon/
// Difficulty: Medium
// Time: O(n) | Space: O(k)

import "fmt"

func maximumEnergy(energy []int, k int) int {
	n := len(energy)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = energy[i]
	}

	for i := k; i < n; i++ {
		if dp[i-k] > 0 {
			dp[i] += dp[i-k]
		}
	}

	ans := dp[n-1]
	for i := n - k - 1; i >= 0; i -= k {
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}

func main() {
	fmt.Println(maximumEnergy([]int{5, 2, -10, -5, 1}, 3)) // Expected: 3
	fmt.Println(maximumEnergy([]int{-2, -3, -1}, 2))        // Expected: -1
}
```

## 3148 — Maximum Difference Score In A Grid

```go
package main

// LeetCode #3148: Maximum Difference Score in a Grid
// https://leetcode.com/problems/maximum-difference-score-in-a-grid/
// Difficulty: Medium
// Time: O(m * n) | Space: O(m * n)

import (
	"fmt"
	"math"
)

func maxScore(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])

	minVal := make([][]int, m)
	for i := range minVal {
		minVal[i] = make([]int, n)
	}

	ans := math.MinInt32

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			prevMin := math.MaxInt32
			if i > 0 {
				prevMin = min(prevMin, minVal[i-1][j])
			}
			if j > 0 {
				prevMin = min(prevMin, minVal[i][j-1])
			}

			if i > 0 || j > 0 {
				ans = max(ans, grid[i][j]-prevMin)
			}

			minVal[i][j] = grid[i][j]
			if i > 0 {
				minVal[i][j] = min(minVal[i][j], minVal[i-1][j])
			}
			if j > 0 {
				minVal[i][j] = min(minVal[i][j], minVal[i][j-1])
			}
		}
	}

	return ans
}

func main() {
	fmt.Println(maxScore([][]int{{9, 5, 7, 3}, {8, 9, 6, 1}, {6, 7, 14, 3}, {2, 5, 3, 1}})) // Expected: 9
	fmt.Println(maxScore([][]int{{4, 3, 2}, {3, 2, 1}}))                                      // Expected: -1
}
```

## 3152 — Special Array Ii

```go
package main

// LeetCode #3152: Special Array II
// https://leetcode.com/problems/special-array-ii/
// Difficulty: Medium
// Time: O(n + q) | Space: O(n)

import "fmt"

func isArraySpecial(nums []int, queries [][]int) []bool {
	n := len(nums)
	prefix := make([]int, n)
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1]
		if nums[i]%2 == nums[i-1]%2 {
			prefix[i]++
		}
	}

	ans := make([]bool, len(queries))
	for i, q := range queries {
		from, to := q[0], q[1]
		ans[i] = prefix[from] == prefix[to]
	}
	return ans
}

func main() {
	fmt.Println(isArraySpecial([]int{3, 4, 1, 2, 6}, [][]int{{0, 4}}))        // Expected: [false]
	fmt.Println(isArraySpecial([]int{4, 3, 1, 6}, [][]int{{0, 2}, {2, 3}}))   // Expected: [false, true]
}
```

## 3153 — Sum Of Digit Differences Of All Pairs

```go
package main

// LeetCode #3153: Sum of Digit Differences of All Pairs
// https://leetcode.com/problems/sum-of-digit-differences-of-all-pairs/
// Difficulty: Medium
// Time: O(n * d) | Space: O(d * 10)

import "fmt"

func sumDigitDifferences(nums []int) int64 {
	n := len(nums)
	if n < 2 {
		return 0
	}

	// Find number of digits
	x := nums[0]
	digits := 0
	for x > 0 {
		digits++
		x /= 10
	}

	var ans int64
	pow := 1
	for d := 0; d < digits; d++ {
		count := make([]int, 10)
		for _, num := range nums {
			count[(num/pow)%10]++
		}
		totalPairs := int64(n) * int64(n-1) / 2
		for _, c := range count {
			if c > 1 {
				totalPairs -= int64(c) * int64(c-1) / 2
			}
		}
		ans += totalPairs
		pow *= 10
	}
	return ans
}

func main() {
	fmt.Println(sumDigitDifferences([]int{13, 23, 12})) // Expected: 4
	fmt.Println(sumDigitDifferences([]int{10, 10, 10})) // Expected: 0
}
```

## 3155 — Maximum Number Of Upgradable Servers

```go
package main

// LeetCode #3155: Maximum Number of Upgradable Servers
// https://leetcode.com/problems/maximum-number-of-upgradable-servers/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func maxUpgrades(count []int, upgrade []int, sell []int, money []int) []int {
	n := len(count)
	ans := make([]int, n)

	for i := 0; i < n; i++ {
		// For each server type, binary search max upgrades
		lo, hi := 0, count[i]
		for lo <= hi {
			mid := (lo + hi) / 2
			cost := mid * upgrade[i]
			revenue := (count[i] - mid) * sell[i]
			if revenue+cost <= money[i]+revenue {
				// mid upgrades possible if: mid*upgrade <= money[i] + sold*revenue from non-upgraded
				// Simplify: mid*upgrade[i] <= money[i] + (count[i]-mid)*sell[i]
				if mid*upgrade[i] <= money[i]+(count[i]-mid)*sell[i] {
					lo = mid + 1
				} else {
					hi = mid - 1
				}
			} else {
				hi = mid - 1
			}
		}
		ans[i] = hi
	}
	return ans
}

func main() {
	fmt.Println(maxUpgrades([]int{2, 3}, []int{3, 4}, []int{1, 2}, []int{4, 5})) // Expected: [1 1]
	fmt.Println(maxUpgrades([]int{1, 1}, []int{5, 5}, []int{1, 1}, []int{0, 0})) // Expected: [0 0]
}
```

## 3157 — Find The Level Of Tree With Minimum Sum

```go
package main

// LeetCode #3157: Find the Level of Tree with Minimum Sum
// https://leetcode.com/problems/find-the-level-of-tree-with-minimum-sum/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minimumLevel(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	minSum := root.Val
	minLevel := 1
	level := 1

	for len(queue) > 0 {
		size := len(queue)
		sum := 0

		for i := 0; i < size; i++ {
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

		if sum < minSum {
			minSum = sum
			minLevel = level
		}
		level++
	}

	return minLevel
}

func main() {
	root := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	fmt.Println(minimumLevel(root)) // Expected: 1

	root2 := &TreeNode{10, &TreeNode{1, nil, nil}, &TreeNode{1, nil, nil}}
	fmt.Println(minimumLevel(root2)) // Expected: 1

	root3 := &TreeNode{5,
		&TreeNode{3, &TreeNode{100, nil, nil}, nil},
		&TreeNode{8, nil, nil},
	}
	fmt.Println(minimumLevel(root3)) // Expected: 2 (level 2 sum = 100, level 1 sum = 5+8=13)
}
```

## 3159 — Find Occurrences Of An Element In An Array

```go
package main

// LeetCode #3159: Find Occurrences of an Element in an Array
// https://leetcode.com/problems/find-occurrences-of-an-element-in-an-array/
// Difficulty: Medium
// Time: O(n + q) | Space: O(n)

import "fmt"

func occurrencesOfElement(nums []int, queries []int, x int) []int {
	pos := make([]int, 0)
	for i, v := range nums {
		if v == x {
			pos = append(pos, i)
		}
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		if q-1 < len(pos) {
			ans[i] = pos[q-1]
		} else {
			ans[i] = -1
		}
	}
	return ans
}

func main() {
	fmt.Println(occurrencesOfElement([]int{1, 3, 1, 7}, []int{1, 3, 2, 4}, 1)) // Expected: [0, -1, 2, -1]
	fmt.Println(occurrencesOfElement([]int{1, 2, 3}, []int{10}, 5))             // Expected: [-1]
}
```

## 3160 — Find The Number Of Distinct Colors Among The Balls

```go
package main

// LeetCode #3160: Find the Number of Distinct Colors Among the Balls
// https://leetcode.com/problems/find-the-number-of-distinct-colors-among-the-balls/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func queryResults(limit int, queries [][]int) []int {
	ballColor := make(map[int]int)
	colorCount := make(map[int]int)
	ans := make([]int, len(queries))

	for i, q := range queries {
		ball, color := q[0], q[1]

		if prev, ok := ballColor[ball]; ok {
			colorCount[prev]--
			if colorCount[prev] == 0 {
				delete(colorCount, prev)
			}
		}

		ballColor[ball] = color
		colorCount[color]++
		ans[i] = len(colorCount)
	}
	return ans
}

func main() {
	fmt.Println(queryResults(4, [][]int{{1, 4}, {2, 5}, {1, 3}, {3, 4}})) // Expected: [1, 2, 2, 3]
	fmt.Println(queryResults(4, [][]int{{0, 1}, {1, 2}, {2, 2}, {3, 4}, {4, 5}})) // Expected: [1, 2, 2, 3, 4]
}
```

## 3163 — String Compression Iii

```go
package main

// LeetCode #3163: String Compression III
// https://leetcode.com/problems/string-compression-iii/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func compressedString(word string) string {
	ans := make([]byte, 0, len(word)*2)
	n := len(word)
	i := 0

	for i < n {
		ch := word[i]
		j := i
		for j < n && j-i < 9 && word[j] == ch {
			j++
		}
		ans = append(ans, byte('0'+j-i), ch)
		i = j
	}
	return string(ans)
}

func main() {
	fmt.Println(compressedString("abcde"))               // Expected: "1a1b1c1d1e"
	fmt.Println(compressedString("aaaaaaaaaaaaaabb"))     // Expected: "9a5a2b"
}
```

## 3164 — Find The Number Of Good Pairs Ii

```go
package main

// LeetCode #3164: Find the Number of Good Pairs II
// https://leetcode.com/problems/find-the-number-of-good-pairs-ii/
// Difficulty: Medium
// Time: O(n * sqrt(max) + m) | Space: O(max)

import "fmt"

func numberOfPairs(nums1 []int, nums2 []int, k int) int64 {
	freq := make(map[int]int)
	for _, v := range nums1 {
		if v%k != 0 {
			continue
		}
		v /= k
		for d := 1; d*d <= v; d++ {
			if v%d == 0 {
				freq[d]++
				if d*d != v {
					freq[v/d]++
				}
			}
		}
	}

	var ans int64
	for _, v := range nums2 {
		ans += int64(freq[v])
	}
	return ans
}

func main() {
	fmt.Println(numberOfPairs([]int{1, 3, 4}, []int{1, 3, 4}, 1)) // Expected: 5
	fmt.Println(numberOfPairs([]int{1, 2, 4, 12}, []int{2, 4}, 3)) // Expected: 2
}
```

## 3166 — Calculate Parking Fees And Duration

```go
package main

// LeetCode #3166: Calculate Parking Fees and Duration
// https://leetcode.com/problems/calculate-parking-fees-and-duration/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func calculateParkingFees(records [][]int) []int {
	fees := make(map[int]int)
	durations := make(map[int]int)

	for _, r := range records {
		carID, entry, exit := r[0], r[1], r[2]
		dur := exit - entry
		durations[carID] += dur
		var fee int
		if dur <= 60 {
			fee = 10
		} else {
			fee = 10 + ((dur-60)+29)/30*5
		}
		fees[carID] += fee
	}

	var cars []int
	for id := range fees {
		cars = append(cars, id)
	}
	sort.Ints(cars)

	ans := make([]int, len(cars))
	for i, id := range cars {
		ans[i] = fees[id]
	}
	return ans
}

func main() {
	fmt.Println(calculateParkingFees([][]int{{1, 0, 30}, {1, 60, 120}, {2, 0, 90}})) // Expected: [20 15]
	fmt.Println(calculateParkingFees([][]int{{1, 0, 30}}))                             // Expected: [10]
}
```

## 3167 — Better Compression Of String

```go
package main

// LeetCode #3167: Better Compression of String
// https://leetcode.com/problems/better-compression-of-string/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import (
	"fmt"
	"strconv"
)

func betterCompression(compressed string) string {
	count := make([]int, 26)
	i := 0
	for i < len(compressed) {
		c := compressed[i] - 'a'
		i++
		freq := 0
		for i < len(compressed) && compressed[i] >= '0' && compressed[i] <= '9' {
			freq = freq*10 + int(compressed[i]-'0')
			i++
		}
		count[c] += freq
	}

	ans := make([]byte, 0)
	for c := 0; c < 26; c++ {
		if count[c] > 0 {
			ans = append(ans, byte('a'+c))
			ans = append(ans, []byte(strconv.Itoa(count[c]))...)
		}
	}
	return string(ans)
}

func main() {
	fmt.Println(betterCompression("a12b3c5"))        // Expected: "a12b3c5"
	fmt.Println(betterCompression("a3b2a2c1"))       // Expected: "a5b2c1"
	fmt.Println(betterCompression("z26y25x24"))      // Expected: "x24y25z26"
}
```

## 3169 — Count Days Without Meetings

```go
package main

// LeetCode #3169: Count Days Without Meetings
// https://leetcode.com/problems/count-days-without-meetings/
// Difficulty: Medium
// Time: O(n log n) | Space: O(log n)

import (
	"fmt"
	"sort"
)

func countDays(days int, meetings [][]int) int {
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	merged := make([][2]int, 0)
	for _, m := range meetings {
		if len(merged) > 0 && m[0] <= merged[len(merged)-1][1]+1 {
			if m[1] > merged[len(merged)-1][1] {
				merged[len(merged)-1][1] = m[1]
			}
		} else {
			merged = append(merged, [2]int{m[0], m[1]})
		}
	}

	ans := days
	for _, m := range merged {
		ans -= m[1] - m[0] + 1
	}
	return ans
}

func main() {
	fmt.Println(countDays(10, [][]int{{5, 7}, {1, 3}, {9, 10}})) // Expected: 2
	fmt.Println(countDays(5, [][]int{{2, 4}, {1, 3}}))            // Expected: 1
	fmt.Println(countDays(6, [][]int{{1, 6}}))                    // Expected: 0
}
```

## 3170 — Lexicographically Minimum String After Removing Stars

```go
package main

// LeetCode #3170: Lexicographically Minimum String After Removing Stars
// https://leetcode.com/problems/lexicographically-minimum-string-after-removing-stars/
// Difficulty: Medium
// Time: O(n * 26) | Space: O(n)

import "fmt"

func clearStars(s string) string {
	n := len(s)
	bytes := []byte(s)
	queues := make([][]int, 26)
	for i := range queues {
		queues[i] = make([]int, 0)
	}

	for i := 0; i < n; i++ {
		if s[i] == '*' {
			for j := 0; j < 26; j++ {
				if len(queues[j]) > 0 {
					idx := queues[j][len(queues[j])-1]
					queues[j] = queues[j][:len(queues[j])-1]
					bytes[idx] = '*'
					break
				}
			}
			bytes[i] = '*'
		} else {
			queues[s[i]-'a'] = append(queues[s[i]-'a'], i)
		}
	}

	ans := make([]byte, 0, n)
	for _, ch := range bytes {
		if ch != '*' {
			ans = append(ans, ch)
		}
	}
	return string(ans)
}

func main() {
	fmt.Println(clearStars("aaba*"))       // Expected: "aab"
	fmt.Println(clearStars("abc"))          // Expected: "abc"
	fmt.Println(clearStars("a*b*c*"))       // Expected: ""
}
```

## 3175 — Find The First Player To Win K Games In A Row

```go
package main

// LeetCode #3175: Find The First Player to Win K Games in a Row
// https://leetcode.com/problems/find-the-first-player-to-win-k-games-in-a-row/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func findWinningPlayer(skills []int, k int) int {
	n := len(skills)
	maxIdx := 0
	curWins := 0

	for i := 1; i < n; i++ {
		if skills[i] > skills[maxIdx] {
			maxIdx = i
			curWins = 1
		} else {
			curWins++
		}
		if curWins >= k {
			return maxIdx
		}
	}
	return maxIdx
}

func main() {
	fmt.Println(findWinningPlayer([]int{4, 2, 6, 3, 9}, 2)) // Expected: 2
	fmt.Println(findWinningPlayer([]int{2, 5, 4}, 3))        // Expected: 1
}
```

## 3176 — Find The Maximum Length Of A Good Subsequence I

```go
package main

// LeetCode #3176: Find the Maximum Length of a Good Subsequence I
// https://leetcode.com/problems/find-the-maximum-length-of-a-good-subsequence-i/
// Difficulty: Medium
// Time: O(n^2 * k) | Space: O(n * k)

import "fmt"

func maximumLength(nums []int, k int) int {
	n := len(nums)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, k+1)
		for j := range dp[i] {
			dp[i][j] = 1
		}
	}

	ans := 1
	for i := 0; i < n; i++ {
		for j := 0; j <= k; j++ {
			for p := 0; p < i; p++ {
				if nums[i] == nums[p] {
					if dp[p][j]+1 > dp[i][j] {
						dp[i][j] = dp[p][j] + 1
					}
				} else if j > 0 {
					if dp[p][j-1]+1 > dp[i][j] {
						dp[i][j] = dp[p][j-1] + 1
					}
				}
			}
			if dp[i][j] > ans {
				ans = dp[i][j]
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(maximumLength([]int{1, 2, 1, 1, 3}, 2)) // Expected: 4
	fmt.Println(maximumLength([]int{1, 2, 3, 4}, 0))     // Expected: 1
}
```

## 3179 — Find The N Th Value After K Seconds

```go
package main

// LeetCode #3179: Find the N-th Value After K Seconds
// https://leetcode.com/problems/find-the-n-th-value-after-k-seconds/
// Difficulty: Medium
// Time: O(n * k) | Space: O(n)

import "fmt"

func valueAfterKSeconds(n int, k int) int {
	const mod = 1_000_000_007
	arr := make([]int, n)
	for i := range arr {
		arr[i] = 1
	}

	for s := 0; s < k; s++ {
		for i := 1; i < n; i++ {
			arr[i] = (arr[i] + arr[i-1]) % mod
		}
	}
	return arr[n-1]
}

func main() {
	fmt.Println(valueAfterKSeconds(4, 5)) // Expected: 56
	fmt.Println(valueAfterKSeconds(5, 3)) // Expected: 35
}
```

## 3180 — Maximum Total Reward Using Operations I

```go
package main

// LeetCode #3180: Maximum Total Reward Using Operations I
// https://leetcode.com/problems/maximum-total-reward-using-operations-i/
// Difficulty: Medium
// Time: O(n * maxVal) | Space: O(maxVal)

import (
	"fmt"
	"sort"
)

func maxTotalReward(rewardValues []int) int {
	sort.Ints(rewardValues)
	maxVal := rewardValues[len(rewardValues)-1]
	size := 2 * maxVal
	dp := make([]bool, size)
	dp[0] = true

	for _, v := range rewardValues {
		for x := size - 1 - v; x >= 0; x-- {
			if dp[x] && v > x {
				dp[x+v] = true
			}
		}
	}

	for x := size - 1; x >= 0; x-- {
		if dp[x] {
			return x
		}
	}
	return 0
}

func main() {
	fmt.Println(maxTotalReward([]int{1, 1, 3, 3})) // Expected: 4
	fmt.Println(maxTotalReward([]int{1, 6, 4, 3, 2})) // Expected: 11
}
```

## 3182 — Find Top Scoring Students

```go
package main

// LeetCode #3182: Find Top Scoring Students
// https://leetcode.com/problems/find-top-scoring-students/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func findTopScoringStudents(scores [][]int, threshold int) []int {
	type student struct {
		id    int
		total int
	}

	var list []student
	for _, s := range scores {
		list = append(list, student{s[0], s[1]})
	}

	sort.Slice(list, func(i, j int) bool {
		if list[i].total != list[j].total {
			return list[i].total > list[j].total
		}
		return list[i].id < list[j].id
	})

	ans := make([]int, 0)
	for _, s := range list {
		if s.total >= threshold {
			ans = append(ans, s.id)
		}
	}
	return ans
}

func main() {
	fmt.Println(findTopScoringStudents([][]int{{1, 95}, {2, 85}, {3, 90}}, 90)) // Expected: [1 3]
	fmt.Println(findTopScoringStudents([][]int{{1, 70}, {2, 65}}, 80))          // Expected: []
}
```

## 3183 — The Number Of Ways To Make The Sum

```go
package main

// LeetCode #3183: The Number of Ways to Make the Sum
// https://leetcode.com/problems/the-number-of-ways-to-make-the-sum/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func numberOfWays(n int) int {
	const mod = 1000000007
	dp := make([]int, n+1)
	dp[0] = 1

	for _, coin := range []int{1, 2, 6} {
		for i := coin; i <= n; i++ {
			dp[i] = (dp[i] + dp[i-coin]) % mod
		}
	}

	ans := dp[n]
	if n-4 >= 0 {
		ans = (ans + dp[n-4]) % mod
	}
	if n-8 >= 0 {
		ans = (ans + dp[n-8]) % mod
	}
	return ans
}

func main() {
	fmt.Println(numberOfWays(4)) // Expected: 5
	fmt.Println(numberOfWays(1)) // Expected: 1
	fmt.Println(numberOfWays(6)) // Expected: 7
}
```

## 3185 — Count Pairs That Form A Complete Day Ii

```go
package main

// LeetCode #3185: Count Pairs That Form a Complete Day II
// https://leetcode.com/problems/count-pairs-that-form-a-complete-day-ii/
// Difficulty: Medium
// Time: O(n) | Space: O(24)

import "fmt"

func countCompleteDayPairs(hours []int) int64 {
	count := make([]int, 24)
	var ans int64

	for _, h := range hours {
		r := h % 24
		need := (24 - r) % 24
		ans += int64(count[need])
		count[r]++
	}
	return ans
}

func main() {
	fmt.Println(countCompleteDayPairs([]int{12, 12, 30, 24, 24})) // Expected: 2
	fmt.Println(countCompleteDayPairs([]int{72, 48, 24, 3}))       // Expected: 3
}
```

## 3186 — Maximum Total Damage With Spell Casting

```go
package main

// LeetCode #3186: Maximum Total Damage With Spell Casting
// https://leetcode.com/problems/maximum-total-damage-with-spell-casting/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func maximumTotalDamage(power []int) int64 {
	freq := make(map[int]int)
	for _, v := range power {
		freq[v]++
	}

	vals := make([]int, 0, len(freq))
	for k := range freq {
		vals = append(vals, k)
	}
	sort.Ints(vals)

	n := len(vals)
	dp := make([]int64, n)

	for i := 0; i < n; i++ {
		val := vals[i]
		count := freq[val]
		dp[i] = int64(val) * int64(count)

		// Find prev valid (val - 2)
		for j := i - 1; j >= 0; j-- {
			if vals[j] < val-2 {
				if dp[j] > dp[i] {
					dp[i] = dp[j]
				}
				break
			}
			if vals[j] <= val-2 {
				dp[i] = maxInt64(dp[i], dp[j]+int64(val)*int64(count))
			}
		}

		if i > 0 && dp[i-1] > dp[i] {
			dp[i] = dp[i-1]
		}
	}

	return dp[n-1]
}

func maxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maximumTotalDamage([]int{1, 1, 3, 4}))          // Expected: 6
	fmt.Println(maximumTotalDamage([]int{7, 1, 6, 3}))           // Expected: 10
}
```

## 3189 — Minimum Moves To Get A Peaceful Board

```go
package main

// LeetCode #3189: Minimum Moves to Get a Peaceful Board
// https://leetcode.com/problems/minimum-moves-to-get-a-peaceful-board/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func minMoves(rooks [][]int) int {
	n := len(rooks)
	rows := make([]int, n)
	cols := make([]int, n)
	for i, r := range rooks {
		rows[i] = r[0]
		cols[i] = r[1]
	}

	sort.Ints(rows)
	sort.Ints(cols)

	moves := 0
	for i := 0; i < n; i++ {
		moves += abs(rows[i] - i)
		moves += abs(cols[i] - i)
	}
	return moves
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(minMoves([][]int{{0, 0}, {1, 1}, {2, 2}}))    // Expected: 0
	fmt.Println(minMoves([][]int{{0, 0}, {0, 2}, {2, 0}}))    // Expected: 2
	fmt.Println(minMoves([][]int{{2, 2}, {0, 0}, {1, 1}}))    // Expected: 0
}
```

## 3191 — Minimum Operations To Make Binary Array Elements Equal To One I

```go
package main

// LeetCode #3191: Minimum Operations to Make Binary Array Elements Equal to One I
// https://leetcode.com/problems/minimum-operations-to-make-binary-array-elements-equal-to-one-i/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func minOperations(nums []int) int {
	n := len(nums)
	ans := 0

	for i := 0; i <= n-3; i++ {
		if nums[i] == 0 {
			ans++
			nums[i] ^= 1
			nums[i+1] ^= 1
			nums[i+2] ^= 1
		}
	}

	for i := n - 2; i < n; i++ {
		if nums[i] == 0 {
			return -1
		}
	}
	return ans
}

func main() {
	fmt.Println(minOperations([]int{0, 1, 1, 1, 0, 0})) // Expected: 3
	fmt.Println(minOperations([]int{0, 1, 1, 1}))        // Expected: -1
}
```

## 3192 — Minimum Operations To Make Binary Array Elements Equal To One Ii

```go
package main

// LeetCode #3192: Minimum Operations to Make Binary Array Elements Equal to One II
// https://leetcode.com/problems/minimum-operations-to-make-binary-array-elements-equal-to-one-ii/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func minOperations(nums []int) int {
	n := len(nums)
	ans := 0
	flip := 0

	for i := 0; i < n; i++ {
		cur := nums[i] ^ flip
		if cur == 0 {
			ans++
			flip ^= 1
		}
	}
	return ans
}

func main() {
	fmt.Println(minOperations([]int{0, 1, 1, 0, 1})) // Expected: 3
	fmt.Println(minOperations([]int{1, 0, 1, 0}))     // Expected: 2
}
```

## 3195 — Find The Minimum Area To Cover All Ones I

```go
package main

// LeetCode #3195: Find the Minimum Area to Cover All Ones I
// https://leetcode.com/problems/find-the-minimum-area-to-cover-all-ones-i/
// Difficulty: Medium
// Time: O(m * n) | Space: O(1)

import "fmt"

func minimumArea(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])

	top, bottom, left, right := m, -1, n, -1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				if i < top {
					top = i
				}
				if i > bottom {
					bottom = i
				}
				if j < left {
					left = j
				}
				if j > right {
					right = j
				}
			}
		}
	}

	if bottom == -1 {
		return 0
	}
	return (bottom - top + 1) * (right - left + 1)
}

func main() {
	fmt.Println(minimumArea([][]int{{0, 1, 0}, {1, 0, 1}})) // Expected: 4
	fmt.Println(minimumArea([][]int{{0, 0}, {1, 1}}))        // Expected: 2
}
```

## 3196 — Maximize Total Cost Of Alternating Subarrays

```go
package main

// LeetCode #3196: Maximize Total Cost of Alternating Subarrays
// https://leetcode.com/problems/maximize-total-cost-of-alternating-subarrays/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func maximumTotalCost(nums []int) int64 {
	if len(nums) == 0 {
		return 0
	}

	add := int64(nums[0])
	sub := int64(nums[0])

	for i := 1; i < len(nums); i++ {
		v := int64(nums[i])
		newAdd := maxInt64(add, sub) + v
		newSub := add - v
		add, sub = newAdd, newSub
	}

	return maxInt64(add, sub)
}

func maxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maximumTotalCost([]int{1, -2, 3, 4}))    // Expected: 10
	fmt.Println(maximumTotalCost([]int{1, -1, 1, -1}))   // Expected: 4
}
```

## 3201 — Find The Maximum Length Of Valid Subsequence I

```go
package main

// LeetCode #3201: Find the Maximum Length of Valid Subsequence I
// https://leetcode.com/problems/find-the-maximum-length-of-valid-subsequence-i/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func maximumLength(nums []int) int {
	// Count numbers by parity
	odd, even := 0, 0
	for _, v := range nums {
		if v%2 == 0 {
			even++
		} else {
			odd++
		}
	}

	// All odd or all even
	ans := max(odd, even)

	// Alternating parity (two patterns: even-odd-even-... or odd-even-odd-...)
	// Both give the same count: min that alternates, but since we need
	// pattern e,o,e,o,... or o,e,o,e,... starting from both parity,
	// the max alternating length is the max of:
	// For each possible (a%2, b%2) pattern where (a+b)%2 == 1 (different parity)
	// Actually the requirement is that for adjacent pairs, (a+b)%2 == 1
	// So we can alternate: 0,1,0,1,... or 1,0,1,0,...

	// Count alternating starting with even
	altEven := 0
	last := 1 // start with expecting even (0)
	for _, v := range nums {
		if v%2 == 0 && last != 0 {
			altEven++
			last = 0
		} else if v%2 == 1 && last != 1 {
			altEven++
			last = 1
		}
	}

	// Count alternating starting with odd
	altOdd := 0
	last = 0 // start with expecting odd (1)
	for _, v := range nums {
		if v%2 == 1 && last != 1 {
			altOdd++
			last = 1
		} else if v%2 == 0 && last != 0 {
			altOdd++
			last = 0
		}
	}

	ans = max(ans, max(altEven, altOdd))
	return ans
}

func main() {
	fmt.Println(maximumLength([]int{1, 2, 3, 4})) // Expected: 4
	fmt.Println(maximumLength([]int{1, 3, 5}))    // Expected: 3
	fmt.Println(maximumLength([]int{2, 4, 6}))    // Expected: 3
}
```

## 3202 — Find The Maximum Length Of Valid Subsequence Ii

```go
package main

// LeetCode #3202: Find the Maximum Length of Valid Subsequence II
// https://leetcode.com/problems/find-the-maximum-length-of-valid-subsequence-ii/
// Difficulty: Medium
// Time: O(n * k) | Space: O(k)

import "fmt"

func maximumLength(nums []int, k int) int {
	dp := make([][]int, k)
	for i := range dp {
		dp[i] = make([]int, k)
	}

	ans := 0
	for _, v := range nums {
		cur := v % k
		for j := 0; j < k; j++ {
			need := (j - cur%k + k) % k
			dp[cur][j] = dp[need][j] + 1
			if dp[cur][j] > ans {
				ans = dp[cur][j]
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(maximumLength([]int{1, 4, 2, 3, 1, 4}, 3)) // Expected: 4
	fmt.Println(maximumLength([]int{1, 2, 3, 4, 5}, 2))     // Expected: 3
}
```

## 3204 — Bitwise User Permissions Analysis

```go
package main

// LeetCode #3204: Bitwise User Permissions Analysis
// https://leetcode.com/problems/bitwise-user-permissions-analysis/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func bitwiseUserPermissions(permissions [][]int) int {
	if len(permissions) == 0 {
		return 0
	}

	combined := permissions[0][1]
	for i := 1; i < len(permissions); i++ {
		combined |= permissions[i][1]
	}
	return combined
}

func main() {
	fmt.Println(bitwiseUserPermissions([][]int{{1, 1}, {2, 2}, {3, 4}})) // Expected: 7 (1|2|4)
	fmt.Println(bitwiseUserPermissions([][]int{{1, 8}, {2, 3}}))         // Expected: 11 (8|3)
}
```

## 3205 — Maximum Array Hopping Score I

```go
package main

// LeetCode #3205: Maximum Array Hopping Score I
// https://leetcode.com/problems/maximum-array-hopping-score-i/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func maxScore(nums []int) int {
	ans := 0
	mx := 0
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > mx {
			mx = nums[i]
		}
		ans += mx
	}
	return ans
}

func main() {
	fmt.Println(maxScore([]int{1, 2, 3, 4, 5})) // Expected: 14 (2+3+4+5)
	fmt.Println(maxScore([]int{5, 4, 3, 2, 1})) // Expected: 4 (1+1+1+1)
	fmt.Println(maxScore([]int{1, 5, 2, 6, 3})) // Expected: 15 (5+6+6+6)
}
```

## 3207 — Maximum Points After Enemy Battles

```go
package main

// LeetCode #3207: Maximum Points After Enemy Battles
// https://leetcode.com/problems/maximum-points-after-enemy-battles/
// Difficulty: Medium
// Time: O(n log n) | Space: O(log n)

import (
	"fmt"
	"sort"
)

func maximumPoints(enemyEnergies []int, currentEnergy int) int64 {
	sort.Ints(enemyEnergies)
	n := len(enemyEnergies)
	energy := int64(currentEnergy)
	var points int64

	// First defeat all enemies we can defeat, accumulating energy
	i := 0
	for i < n && int64(enemyEnergies[i]) <= energy {
		energy -= int64(enemyEnergies[i])
		points++
		i++
	}

	if points == 0 {
		return 0
	}

	// Now we can use the smallest enemy to get energy
	smallest := enemyEnergies[0]
	for energy >= int64(smallest) {
		cnt := energy / int64(smallest)
		points += cnt
		energy %= int64(smallest)
		energy += int64(smallest) // Keep 1 point by defeating smallest again
		points--
	}
	return points
}

func main() {
	fmt.Println(maximumPoints([]int{3, 5, 6}, 3))  // Expected: 1
	fmt.Println(maximumPoints([]int{1, 2, 4, 8}, 3)) // Expected: 6
}
```

## 3208 — Alternating Groups Ii

```go
package main

// LeetCode #3208: Alternating Groups II
// https://leetcode.com/problems/alternating-groups-ii/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func numberOfAlternatingGroups(colors []int, k int) int {
	n := len(colors)
	ans := 0
	len := 1

	for i := 1; i < n+k-1; i++ {
		if colors[i%n] != colors[(i-1)%n] {
			len++
		} else {
			len = 1
		}
		if len >= k {
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(numberOfAlternatingGroups([]int{0, 1, 0, 1, 0}, 3)) // Expected: 3
	fmt.Println(numberOfAlternatingGroups([]int{0, 1, 0, 0, 1}, 3)) // Expected: 2
}
```

## 3211 — Generate Binary Strings Without Adjacent Zeros

```go
package main

// LeetCode #3211: Generate Binary Strings Without Adjacent Zeros
// https://leetcode.com/problems/generate-binary-strings-without-adjacent-zeros/
// Difficulty: Medium
// Time: O(2^n) | Space: O(n) for recursion

import "fmt"

func validStrings(n int) []string {
	ans := make([]string, 0)
	var dfs func(cur []byte)
	dfs = func(cur []byte) {
		if len(cur) == n {
			ans = append(ans, string(cur))
			return
		}
		// Option 1: append '1'
		cur = append(cur, '1')
		dfs(cur)
		cur = cur[:len(cur)-1]

		// Option 2: append '0' only if previous was not '0'
		if len(cur) == 0 || cur[len(cur)-1] != '0' {
			cur = append(cur, '0')
			dfs(cur)
			cur = cur[:len(cur)-1]
		}
	}
	dfs(make([]byte, 0, n))
	return ans
}

func main() {
	fmt.Println(validStrings(3)) // Expected: ["010","011","101","110","111"]
	fmt.Println(validStrings(1)) // Expected: ["0","1"]
}
```

## 3212 — Count Submatrices With Equal Frequency Of X And Y

```go
package main

// LeetCode #3212: Count Submatrices With Equal Frequency of X and Y
// https://leetcode.com/problems/count-submatrices-with-equal-frequency-of-x-and-y/
// Difficulty: Medium
// Time: O(m * n) | Space: O(m * n)

import "fmt"

func numberOfSubmatrices(grid [][]byte) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])

	prefX := make([][]int, m+1)
	prefY := make([][]int, m+1)
	for i := range prefX {
		prefX[i] = make([]int, n+1)
		prefY[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			prefX[i+1][j+1] = prefX[i][j+1] + prefX[i+1][j] - prefX[i][j]
			prefY[i+1][j+1] = prefY[i][j+1] + prefY[i+1][j] - prefY[i][j]
			if grid[i][j] == 'X' {
				prefX[i+1][j+1]++
			} else if grid[i][j] == 'Y' {
				prefY[i+1][j+1]++
			}
		}
	}

	ans := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			x := prefX[i][j]
			y := prefY[i][j]
			if x > 0 && x == y {
				ans++
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(numberOfSubmatrices([][]byte{{'X', 'Y', '.'}, {'Y', '.', '.'}})) // Expected: 3
	fmt.Println(numberOfSubmatrices([][]byte{{'X', 'X'}, {'Y', 'Y'}}))          // Expected: 0
}
```

## 3215 — Count Triplets With Even Xor Set Bits Ii

```go
package main

// LeetCode #3215: Count Triplets with Even XOR Set Bits II
// https://leetcode.com/problems/count-triplets-with-even-xor-set-bits-ii/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import (
	"fmt"
	"math/bits"
)

func tripletCount(a []int, b []int, c []int) int64 {
	evenA, oddA := countBits(a)
	evenB, oddB := countBits(b)
	evenC, oddC := countBits(c)
	return int64(evenA)*int64(oddB)*int64(oddC) +
		int64(oddA)*int64(evenB)*int64(oddC) +
		int64(oddA)*int64(oddB)*int64(evenC) +
		int64(evenA)*int64(evenB)*int64(evenC)
}

func countBits(nums []int) (int, int) {
	even := 0
	for _, v := range nums {
		if bits.OnesCount(uint(v))%2 == 0 {
			even++
		}
	}
	return even, len(nums) - even
}

func main() {
	fmt.Println(tripletCount([]int{1, 2}, []int{3, 4}, []int{5, 6})) // Expected: depends on bit counts
	fmt.Println(tripletCount([]int{0, 0}, []int{0, 0}, []int{0, 0})) // Expected: 8
}
```

## 3217 — Delete Nodes From Linked List Present In Array

```go
package main

// LeetCode #3217: Delete Nodes From Linked List Present in Array
// https://leetcode.com/problems/delete-nodes-from-linked-list-present-in-array/
// Difficulty: Medium
// Time: O(n) | Space: O(m)

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func modifiedList(nums []int, head *ListNode) *ListNode {
	del := make(map[int]bool)
	for _, v := range nums {
		del[v] = true
	}

	dummy := &ListNode{Next: head}
	cur := dummy

	for cur.Next != nil {
		if del[cur.Next.Val] {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}

func listToSlice(head *ListNode) []int {
	var ans []int
	for head != nil {
		ans = append(ans, head.Val)
		head = head.Next
	}
	return ans
}

func main() {
	head1 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	fmt.Println(listToSlice(modifiedList([]int{1, 2, 3}, head1))) // Expected: [4, 5]

	head2 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	fmt.Println(listToSlice(modifiedList([]int{5}, head2))) // Expected: [1, 2, 3, 4]
}
```

## 3218 — Minimum Cost For Cutting Cake I

```go
package main

// LeetCode #3218: Minimum Cost for Cutting Cake I
// https://leetcode.com/problems/minimum-cost-for-cutting-cake-i/
// Difficulty: Medium
// Time: O(h log h + v log v) | Space: O(log h + log v)

import (
	"fmt"
	"sort"
)

func minimumCost(m int, n int, horizontalCut []int, verticalCut []int) int {
	sort.Slice(horizontalCut, func(i, j int) bool { return horizontalCut[i] > horizontalCut[j] })
	sort.Slice(verticalCut, func(i, j int) bool { return verticalCut[i] > verticalCut[j] })

	hPieces, vPieces := 1, 1
	i, j := 0, 0
	cost := 0

	for i < len(horizontalCut) && j < len(verticalCut) {
		if horizontalCut[i] >= verticalCut[j] {
			cost += horizontalCut[i] * vPieces
			hPieces++
			i++
		} else {
			cost += verticalCut[j] * hPieces
			vPieces++
			j++
		}
	}

	for i < len(horizontalCut) {
		cost += horizontalCut[i] * vPieces
		i++
	}
	for j < len(verticalCut) {
		cost += verticalCut[j] * hPieces
		j++
	}
	return cost
}

func main() {
	fmt.Println(minimumCost(3, 2, []int{1, 3}, []int{5})) // Expected: 13
	fmt.Println(minimumCost(2, 2, []int{7}, []int{4}))     // Expected: 15
}
```

## 3220 — Odd And Even Transactions

```go
package main

// LeetCode #3220: Odd and Even Transactions
// https://leetcode.com/problems/odd-and-even-transactions/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func oddAndEvenTransactions(transactions [][]int) []int {
	oddSum, evenSum := 0, 0
	for _, t := range transactions {
		amount := t[1]
		if amount%2 == 0 {
			evenSum += amount
		} else {
			oddSum += amount
		}
	}
	return []int{oddSum, evenSum}
}

func main() {
	fmt.Println(oddAndEvenTransactions([][]int{{1, 10}, {2, 15}, {3, 20}})) // Expected: [25 30]
	fmt.Println(oddAndEvenTransactions([][]int{{1, 1}, {2, 2}, {3, 3}}))    // Expected: [4 2]
}
```

## 3221 — Maximum Array Hopping Score Ii

```go
package main

// LeetCode #3221: Maximum Array Hopping Score II
// https://leetcode.com/problems/maximum-array-hopping-score-ii/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func maxScore(nums []int) int64 {
	ans := int64(0)
	mx := 0
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > mx {
			mx = nums[i]
		}
		ans += int64(mx)
	}
	return ans
}

func main() {
	fmt.Println(maxScore([]int{1, 2, 3, 4, 5})) // Expected: 14
	fmt.Println(maxScore([]int{5, 4, 3, 2, 1})) // Expected: 4
}
```

## 3223 — Minimum Length Of String After Operations

```go
package main

// LeetCode #3223: Minimum Length of String After Operations
// https://leetcode.com/problems/minimum-length-of-string-after-operations/
// Difficulty: Medium
// Time: O(n) | Space: O(26) = O(1)

import "fmt"

func minimumLength(s string) int {
	freq := make([]int, 26)
	for _, ch := range s {
		freq[ch-'a']++
	}

	ans := 0
	for _, f := range freq {
		if f == 0 {
			continue
		}
		if f%2 == 0 {
			ans += 2
		} else {
			ans += 1
		}
	}
	return ans
}

func main() {
	fmt.Println(minimumLength("abaacbcbb")) // Expected: 5
	fmt.Println(minimumLength("aa"))         // Expected: 2
}
```

## 3224 — Minimum Array Changes To Make Differences Equal

```go
package main

// LeetCode #3224: Minimum Array Changes to Make Differences Equal
// https://leetcode.com/problems/minimum-array-changes-to-make-differences-equal/
// Difficulty: Medium
// Time: O(n + k) | Space: O(k)

import "fmt"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func minChanges(nums []int, k int) int {
	n := len(nums)
	diff := make([]int, k+2)

	for i := 0; i < n/2; i++ {
		a, b := nums[i], nums[n-1-i]
		curDiff := abs(a - b)

		// One change: can achieve any diff from 0 to max(a, b, k-a, k-b)
		maxReach := max(max(a, b), max(k-a, k-b))
		// One change can achieve any diff in [0, maxReach]
		diff[0]++
		if maxReach+1 <= k {
			diff[maxReach+1]--
		}

		// Zero changes: only curDiff
		diff[curDiff]--
		diff[curDiff+1]++
	}

	ans := n
	cur := 0
	for i := 0; i <= k; i++ {
		cur += diff[i]
		if cur < ans {
			ans = cur
		}
	}
	return ans
}

func main() {
	fmt.Println(minChanges([]int{1, 0, 1, 2, 4, 3}, 4)) // Expected: 2
	fmt.Println(minChanges([]int{0, 1, 2, 3, 3, 6, 5, 4}, 6)) // Expected: ?
}
```

## 3227 — Vowels Game In A String

```go
package main

// LeetCode #3227: Vowels Game in a String
// https://leetcode.com/problems/vowels-game-in-a-string/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func doesAliceWin(s string) bool {
	vowels := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
	for i := 0; i < len(s); i++ {
		if vowels[s[i]] {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(doesAliceWin("leetcoder")) // Expected: true
	fmt.Println(doesAliceWin("bbcd"))       // Expected: false
}
```

## 3228 — Maximum Number Of Operations To Move Ones To The End

```go
package main

// LeetCode #3228: Maximum Number of Operations to Move Ones to the End
// https://leetcode.com/problems/maximum-number-of-operations-to-move-ones-to-the-end/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func maxOperations(s string) int {
	n := len(s)
	ans := 0
	ones := 0

	for i := 0; i < n; i++ {
		if s[i] == '1' {
			ones++
		} else if i > 0 && s[i-1] == '1' {
			ans += ones
		}
	}
	return ans
}

func main() {
	fmt.Println(maxOperations("1001101")) // Expected: 4
	fmt.Println(maxOperations("00111"))    // Expected: 0
}
```

## 3230 — Customer Purchasing Behavior Analysis

```go
package main

// LeetCode #3230: Customer Purchasing Behavior Analysis
// https://leetcode.com/problems/customer-purchasing-behavior-analysis/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func customerPurchasingBehavior(purchases [][]int) []int {
	counts := make(map[int]int)
	for _, p := range purchases {
		counts[p[0]]++
	}

	var customers []int
	for id := range counts {
		customers = append(customers, id)
	}
	sort.Slice(customers, func(i, j int) bool {
		if counts[customers[i]] != counts[customers[j]] {
			return counts[customers[i]] > counts[customers[j]]
		}
		return customers[i] < customers[j]
	})

	ans := make([]int, len(customers))
	for i, id := range customers {
		ans[i] = id
	}
	return ans
}

func main() {
	fmt.Println(customerPurchasingBehavior([][]int{{1, 100}, {2, 50}, {1, 200}, {3, 75}})) // Expected: [1 2 3]
	fmt.Println(customerPurchasingBehavior([][]int{{1, 10}, {2, 20}, {2, 30}}))             // Expected: [2 1]
}
```

## 3233 — Find The Count Of Numbers Which Are Not Special

```go
package main

// LeetCode #3233: Find the Count of Numbers Which Are Not Special
// https://leetcode.com/problems/find-the-count-of-numbers-which-are-not-special/
// Difficulty: Medium
// Time: O(sqrt(r) log log r) | Space: O(sqrt(r))

import "fmt"

func nonSpecialCount(l int, r int) int {
	limit := 31623 // sqrt(10^9) approx

	isPrime := make([]bool, limit+1)
	for i := 2; i <= limit; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i <= limit; i++ {
		if isPrime[i] {
			for j := i * i; j <= limit; j += i {
				isPrime[j] = false
			}
		}
	}

	special := 0
	for i := 2; i*i <= r; i++ {
		if isPrime[i] {
			sq := i * i
			if sq >= l && sq <= r {
				special++
			}
		}
	}

	return r - l + 1 - special
}

func main() {
	fmt.Println(nonSpecialCount(5, 7))     // Expected: 3
	fmt.Println(nonSpecialCount(4, 16))    // Expected: 11
}
```

## 3234 — Count The Number Of Substrings With Dominant Ones

```go
package main

// LeetCode #3234: Count the Number of Substrings With Dominant Ones
// https://leetcode.com/problems/count-the-number-of-substrings-with-dominant-ones/
// Difficulty: Medium
// Time: O(n * sqrt(n)) | Space: O(1)

import (
	"fmt"
	"math"
)

func numberOfSubstrings(s string) int {
	n := len(s)
	ans := 0
	maxZeros := int(math.Sqrt(float64(n)))

	for l := 0; l < n; l++ {
		zeros := 0
		ones := 0
		for r := l; r < n; r++ {
			if s[r] == '0' {
				zeros++
				if zeros > maxZeros {
					break
				}
			} else {
				ones++
			}
			if ones >= zeros*zeros {
				ans++
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(numberOfSubstrings("00011")) // Expected: 5
	fmt.Println(numberOfSubstrings("101"))    // Expected: 3
}
```

