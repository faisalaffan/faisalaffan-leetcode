# Medium (Sedang) — Problem 3237–3462

## 3237 — Alt And Tab Simulation

```go
package main

// LeetCode #3237: Alt and Tab Simulation
// https://leetcode.com/problems/alt-and-tab-simulation/
// Difficulty: Medium [Paid]
// Time: O(n + q) | Space: O(n)

import (
	"container/list"
	"fmt"
)

func simulationResult(windows []int, queries []int) []int {
	order := list.New()
	pos := make(map[int]*list.Element)
	for _, w := range windows {
		e := order.PushBack(w)
		pos[w] = e
	}

	for _, q := range queries {
		if e, ok := pos[q]; ok {
			order.MoveToFront(e)
		}
	}

	ans := make([]int, 0, order.Len())
	for e := order.Front(); e != nil; e = e.Next() {
		ans = append(ans, e.Value.(int))
	}
	return ans
}

func main() {
	fmt.Println(simulationResult([]int{1, 2, 3, 4}, []int{3, 1})) // Expected: [1 3 2 4]
	fmt.Println(simulationResult([]int{1, 2, 3}, []int{2}))        // Expected: [2 1 3]
}
```

## 3239 — Minimum Number Of Flips To Make Binary Grid Palindromic I

```go
package main

// LeetCode #3239: Minimum Number of Flips to Make Binary Grid Palindromic I
// https://leetcode.com/problems/minimum-number-of-flips-to-make-binary-grid-palindromic-i/
// Difficulty: Medium
// Time: O(m * n) | Space: O(1)

import "fmt"

func minFlips(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	rowFlips := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n/2; j++ {
			if grid[i][j] != grid[i][n-1-j] {
				rowFlips++
			}
		}
	}

	colFlips := 0
	for j := 0; j < n; j++ {
		for i := 0; i < m/2; i++ {
			if grid[i][j] != grid[m-1-i][j] {
				colFlips++
			}
		}
	}

	return min(rowFlips, colFlips)
}

func main() {
	fmt.Println(minFlips([][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 1}})) // Expected: 2
	fmt.Println(minFlips([][]int{{0, 1}, {0, 1}, {0, 0}}))          // Expected: 1
}
```

## 3240 — Minimum Number Of Flips To Make Binary Grid Palindromic Ii

```go
package main

// LeetCode #3240: Minimum Number of Flips to Make Binary Grid Palindromic II
// https://leetcode.com/problems/minimum-number-of-flips-to-make-binary-grid-palindromic-ii/
// Difficulty: Medium
// Time: O(m * n) | Space: O(1)

import "fmt"

func minFlips(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	ans := 0

	for i := 0; i < m/2; i++ {
		for j := 0; j < n/2; j++ {
			ones := grid[i][j] + grid[i][n-1-j] + grid[m-1-i][j] + grid[m-1-i][n-1-j]
			ans += min(ones, 4-ones)
		}
	}

	mismatchPairs := 0
	onesInMiddle := 0

	if m%2 == 1 {
		mid := m / 2
		for j := 0; j < n/2; j++ {
			if grid[mid][j] != grid[mid][n-1-j] {
				mismatchPairs++
				ans++
			} else if grid[mid][j] == 1 {
				onesInMiddle += 2
			}
		}
	}

	if n%2 == 1 {
		mid := n / 2
		for i := 0; i < m/2; i++ {
			if grid[i][mid] != grid[m-1-i][mid] {
				mismatchPairs++
				ans++
			} else if grid[i][mid] == 1 {
				onesInMiddle += 2
			}
		}
	}

	if m%2 == 1 && n%2 == 1 {
		if grid[m/2][n/2] == 1 {
			ans++
		}
	} else if mismatchPairs == 0 && onesInMiddle%4 != 0 {
		ans += 2
	}

	return ans
}

func main() {
	fmt.Println(minFlips([][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}})) // Expected: 3
	fmt.Println(minFlips([][]int{{0, 1}, {0, 1}, {0, 0}}))          // Expected: 2
}
```

## 3243 — Shortest Distance After Road Addition Queries I

```go
package main

// LeetCode #3243: Shortest Distance After Road Addition Queries I
// https://leetcode.com/problems/shortest-distance-after-road-addition-queries-i/
// Difficulty: Medium
// Time: O(q * (n + q)) | Space: O(n + q)

import (
	"fmt"
	"math"
)

func shortestDistanceAfterQueries(n int, queries [][]int) []int {
	graph := make([][]int, n)
	for i := 0; i < n-1; i++ {
		graph[i] = append(graph[i], i+1)
	}

	ans := make([]int, len(queries))

	for qi, q := range queries {
		u, v := q[0], q[1]
		graph[u] = append(graph[u], v)

		dist := make([]int, n)
		for i := range dist {
			dist[i] = math.MaxInt32
		}
		dist[0] = 0
		queue := []int{0}

		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			for _, nb := range graph[cur] {
				if dist[cur]+1 < dist[nb] {
					dist[nb] = dist[cur] + 1
					queue = append(queue, nb)
				}
			}
		}
		ans[qi] = dist[n-1]
	}
	return ans
}

func main() {
	fmt.Println(shortestDistanceAfterQueries(5, [][]int{{2, 4}, {0, 2}, {0, 4}})) // Expected: [3, 2, 1]
	fmt.Println(shortestDistanceAfterQueries(4, [][]int{{0, 3}, {0, 2}}))          // Expected: [1, 1]
}
```

## 3247 — Number Of Subsequences With Odd Sum

```go
package main

// LeetCode #3247: Number of Subsequences with Odd Sum
// https://leetcode.com/problems/number-of-subsequences-with-odd-sum/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func subsequenceCount(nums []int) int {
	const mod = 1000000007
	n := len(nums)
	oddCount := 0
	for _, v := range nums {
		if v%2 != 0 {
			oddCount++
		}
	}

	if oddCount == 0 {
		return 0
	}

	pow := 1
	for i := 0; i < n-1; i++ {
		pow = (pow * 2) % mod
	}
	return pow
}

func main() {
	fmt.Println(subsequenceCount([]int{1, 2, 3})) // Expected: 4
	fmt.Println(subsequenceCount([]int{2, 4, 6})) // Expected: 0
	fmt.Println(subsequenceCount([]int{1}))        // Expected: 1
}
```

## 3249 — Count The Number Of Good Nodes

```go
package main

// LeetCode #3249: Count the Number of Good Nodes
// https://leetcode.com/problems/count-the-number-of-good-nodes/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func countGoodNodes(edges [][]int) int {
	n := len(edges) + 1
	graph := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	ans := 0

	var dfs func(u, parent int) int
	dfs = func(u, parent int) int {
		size := 1
		childSize := -1
		good := true

		for _, v := range graph[u] {
			if v == parent {
				continue
			}
			sz := dfs(v, u)
			if childSize == -1 {
				childSize = sz
			} else if sz != childSize {
				good = false
			}
			size += sz
		}

		if good {
			ans++
		}
		return size
	}

	dfs(0, -1)
	return ans
}

func main() {
	fmt.Println(countGoodNodes([][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {2, 5}, {2, 6}})) // Expected: 7
	fmt.Println(countGoodNodes([][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {4, 5}}))        // Expected: 6
}
```

## 3252 — Premier League Table Ranking Ii

```go
package main

// LeetCode #3252: Premier League Table Ranking II
// https://leetcode.com/problems/premier-league-table-ranking-ii/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

type team struct {
	id     int
	points int
	gDiff  int
}

func premierLeagueRanking(teams [][]int) []int {
	var list []team
	for _, t := range teams {
		list = append(list, team{t[0], t[1], t[2]})
	}

	sort.Slice(list, func(i, j int) bool {
		if list[i].points != list[j].points {
			return list[i].points > list[j].points
		}
		if list[i].gDiff != list[j].gDiff {
			return list[i].gDiff > list[j].gDiff
		}
		return list[i].id < list[j].id
	})

	ans := make([]int, len(list))
	for i, t := range list {
		ans[i] = t.id
	}
	return ans
}

func main() {
	fmt.Println(premierLeagueRanking([][]int{{1, 10, 5}, {2, 8, 8}, {3, 10, 3}})) // Expected: [1 3 2]
	fmt.Println(premierLeagueRanking([][]int{{1, 6, 2}, {2, 6, 2}}))              // Expected: [1 2]
}
```

## 3253 — Construct String With Minimum Cost Easy

```go
package main

// LeetCode #3253: Construct String with Minimum Cost (Easy)
// https://leetcode.com/problems/construct-string-with-minimum-cost-easy/
// Difficulty: Medium [Paid]
// Time: O(n * m * L) | Space: O(n)

import (
	"fmt"
	"math"
)

func minimumCost(target string, words []string, costs []int) int {
	n := len(target)
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt32
	}

	for i := 0; i < n; i++ {
		if dp[i] == math.MaxInt32 {
			continue
		}
		for j, w := range words {
			if i+len(w) <= n && target[i:i+len(w)] == w {
				if dp[i]+costs[j] < dp[i+len(w)] {
					dp[i+len(w)] = dp[i] + costs[j]
				}
			}
		}
	}

	if dp[n] == math.MaxInt32 {
		return -1
	}
	return dp[n]
}

func main() {
	fmt.Println(minimumCost("abc", []string{"a", "bc", "abc"}, []int{1, 2, 3})) // Expected: 3
	fmt.Println(minimumCost("xyz", []string{"ab", "cd"}, []int{1, 2}))          // Expected: -1
}
```

## 3254 — Find The Power Of K Size Subarrays I

```go
package main

// LeetCode #3254: Find the Power of K-Size Subarrays I
// https://leetcode.com/problems/find-the-power-of-k-size-subarrays-i/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func resultsArray(nums []int, k int) []int {
	n := len(nums)
	ans := make([]int, n-k+1)
	consec := 1

	for i := range n {
		if i > 0 && nums[i] == nums[i-1]+1 {
			consec++
		} else {
			consec = 1
		}
		if i >= k-1 {
			if consec >= k {
				ans[i-k+1] = nums[i]
			} else {
				ans[i-k+1] = -1
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(resultsArray([]int{1, 2, 3, 4, 3, 5}, 3)) // Expected: [3, 4, -1, -1]
	fmt.Println(resultsArray([]int{2, 2, 2, 2, 2}, 4))    // Expected: [-1, -1]
	fmt.Println(resultsArray([]int{3, 2, 3, 2, 3, 2}, 2)) // Expected: [-1, 3, -1, 3, -1]
}
```

## 3255 — Find The Power Of K Size Subarrays Ii

```go
package main

// LeetCode #3255: Find the Power of K-Size Subarrays II
// https://leetcode.com/problems/find-the-power-of-k-size-subarrays-ii/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func resultsArray(nums []int, k int) []int {
	n := len(nums)
	ans := make([]int, n-k+1)
	consec := 1

	for i := range n {
		if i > 0 && nums[i] == nums[i-1]+1 {
			consec++
		} else {
			consec = 1
		}
		if i >= k-1 {
			if consec >= k {
				ans[i-k+1] = nums[i]
			} else {
				ans[i-k+1] = -1
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(resultsArray([]int{1, 2, 3, 4, 3, 5}, 3)) // Expected: [3, 4, -1, -1]
	fmt.Println(resultsArray([]int{1, 4, 5, 2, 3}, 3))    // Expected: [-1, 5, -1]
	fmt.Println(resultsArray([]int{1}, 1))                 // Expected: [1]
}
```

## 3259 — Maximum Energy Boost From Two Drinks

```go
package main

// LeetCode #3259: Maximum Energy Boost From Two Drinks
// https://leetcode.com/problems/maximum-energy-boost-from-two-drinks/
// Difficulty: Medium
// Time: O(n) Space: O(1)

import "fmt"

func main() {
	fmt.Println(maxEnergyBoost([]int{1, 3, 1}, []int{3, 1, 1}))       // 5
	fmt.Println(maxEnergyBoost([]int{4, 1, 1}, []int{1, 1, 3}))       // 7
	fmt.Println(maxEnergyBoost([]int{2, 2, 2, 2}, []int{3, 3, 3, 3})) // 12
}

func maxEnergyBoost(energyDrinkA []int, energyDrinkB []int) int64 {
	n := len(energyDrinkA)
	dpA, dpB := int64(0), int64(0)

	for i := 0; i < n; i++ {
		newA := max(dpA+int64(energyDrinkA[i]), dpB)
		newB := max(dpB+int64(energyDrinkB[i]), dpA)
		dpA, dpB = newA, newB
	}

	if dpA > dpB {
		return dpA
	}
	return dpB
}
```

## 3262 — Find Overlapping Shifts

```go
package main

// LeetCode #3262: Find Overlapping Shifts
// https://leetcode.com/problems/find-overlapping-shifts/
// Difficulty: Medium
// Time: O(n log n) Space: O(n)

import (
	"fmt"
	"sort"
)

func main() {
	// Test case 1
	shifts1 := []Shift{
		{1, 8, 12}, {1, 11, 15}, {1, 14, 18},
		{2, 9, 17}, {2, 16, 20},
		{3, 10, 12}, {3, 13, 15}, {3, 16, 18},
		{4, 8, 10}, {4, 9, 11},
	}
	fmt.Println(countOverlappingShifts(shifts1)) // [[1 2] [2 1] [4 1]]

	// Test case 2
	shifts2 := []Shift{{1, 1, 3}, {1, 2, 4}}
	fmt.Println(countOverlappingShifts(shifts2)) // [[1 1]]

	// Test case 3
	shifts3 := []Shift{{1, 1, 2}, {1, 3, 4}}
	fmt.Println(countOverlappingShifts(shifts3)) // []
}

type Shift struct {
	EmployeeID int
	StartTime  int
	EndTime    int
}

func countOverlappingShifts(shifts []Shift) [][2]int {
	// Group shifts by employee
	empShifts := make(map[int][]Shift)
	for _, s := range shifts {
		empShifts[s.EmployeeID] = append(empShifts[s.EmployeeID], s)
	}

	type result struct {
		employeeID int
		count      int
	}
	var results []result

	for empID, s := range empShifts {
		// Sort shifts by start time
		sort.Slice(s, func(i, j int) bool {
			return s[i].StartTime < s[j].StartTime
		})

		count := 0
		// Sweep line: track maximum end time seen so far
		maxEnd := s[0].EndTime
		for i := 1; i < len(s); i++ {
			if s[i].StartTime < maxEnd {
				count++
			}
			if s[i].EndTime > maxEnd {
				maxEnd = s[i].EndTime
			}
		}

		if count > 0 {
			results = append(results, result{empID, count})
		}
	}

	// Sort by employee ID
	sort.Slice(results, func(i, j int) bool {
		return results[i].employeeID < results[j].employeeID
	})

	out := make([][2]int, len(results))
	for i, r := range results {
		out[i] = [2]int{r.employeeID, r.count}
	}
	return out
}
```

## 3265 — Count Almost Equal Pairs I

```go
package main

// LeetCode #3265: Count Almost Equal Pairs I
// https://leetcode.com/problems/count-almost-equal-pairs-i/
// Difficulty: Medium
// Time: O(n^2 * d) Space: O(d) where d = number of digits

import "fmt"

func main() {
	fmt.Println(countPairs([]int{1, 1, 1}))                                           // 3
	fmt.Println(countPairs([]int{3, 12, 30, 17, 21}))                                 // 2
	fmt.Println(countPairs([]int{1023, 3012, 1230, 2301, 0, 0}))                     // 4
}

func countPairs(nums []int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if isAlmostEqual(nums[i], nums[j]) {
				ans++
			}
		}
	}
	return ans
}

func isAlmostEqual(a, b int) bool {
	sa, sb := fmt.Sprintf("%07d", a), fmt.Sprintf("%07d", b)
	diff := 0
	ca, cb := make([]int, 10), make([]int, 10)
	for k := 0; k < 7; k++ {
		if sa[k] != sb[k] {
			diff++
			if diff > 2 {
				return false
			}
		}
		ca[sa[k]-'0']++
		cb[sb[k]-'0']++
	}
	for k := 0; k < 10; k++ {
		if ca[k] != cb[k] {
			return false
		}
	}
	return diff <= 2
}
```

## 3271 — Hash Divided String

```go
package main

// LeetCode #3271: Hash Divided String
// https://leetcode.com/problems/hash-divided-string/
// Difficulty: Medium
// Time: O(n) Space: O(n/k) for result

import "fmt"

func main() {
	fmt.Println(stringHash("abcd", 2))                                             // "bf"
	fmt.Println(stringHash("mxz", 3))                                              // "i"
	fmt.Println(stringHash("leetcode", 4))                                         // "ob"
}

func stringHash(s string, k int) string {
	res := make([]byte, 0, len(s)/k)
	sum := 0
	for i, ch := range s {
		sum += int(ch - 'a')
		if (i+1)%k == 0 {
			res = append(res, byte('a'+sum%26))
			sum = 0
		}
	}
	return string(res)
}
```

## 3275 — K Th Nearest Obstacle Queries

```go
package main

// LeetCode #3275: K-th Nearest Obstacle Queries
// https://leetcode.com/problems/k-th-nearest-obstacle-queries/
// Difficulty: Medium
// Time: O(n log k) Space: O(k)

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(resultsArray([][]int{{1, 2}, {3, 4}, {2, 3}, {-3, 0}}, 2)) // [-1,7,5,3]
	fmt.Println(resultsArray([][]int{{5, 5}, {4, 4}, {3, 3}}, 1))          // [10,8,6]
	fmt.Println(resultsArray([][]int{{1, 1}, {2, 2}, {3, 3}}, 3))          // [-1,-1,6]
}

type MaxHeap []int

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func resultsArray(queries [][]int, k int) []int {
	ans := make([]int, len(queries))
	h := &MaxHeap{}
	heap.Init(h)

	for i, q := range queries {
		dist := abs(q[0]) + abs(q[1])
		heap.Push(h, dist)
		if h.Len() > k {
			heap.Pop(h)
		}
		if h.Len() == k {
			ans[i] = (*h)[0]
		} else {
			ans[i] = -1
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

## 3278 — Find Candidates For Data Scientist Position Ii

```go
package main

// LeetCode #3278: Find Candidates for Data Scientist Position II
// https://leetcode.com/problems/find-candidates-for-data-scientist-position-ii/
// Difficulty: Medium
// Time: O(c * p) Space: O(c * p)

import (
	"fmt"
	"sort"
)

func main() {
	candidates := []Candidate{
		{101, "Python", 5}, {101, "Tableau", 3}, {101, "PostgreSQL", 4}, {101, "TensorFlow", 2},
		{102, "Python", 4}, {102, "Tableau", 5}, {102, "PostgreSQL", 4}, {102, "R", 4},
		{103, "Python", 3}, {103, "Tableau", 5}, {103, "PostgreSQL", 5}, {103, "Spark", 4},
	}
	projects := []Project{
		{501, "Python", 4}, {501, "Tableau", 3}, {501, "PostgreSQL", 5},
		{502, "Python", 3}, {502, "Tableau", 4}, {502, "R", 2},
	}
	fmt.Println(topCandidates(candidates, projects))
	// Expected: [{501 101 105} {502 102 130}]
}

type Candidate struct {
	ID         int
	Skill      string
	Proficiency int
}

type Project struct {
	ID         int
	Skill      string
	Importance int
}

type ProjectResult struct {
	ProjectID   int
	CandidateID int
	Score       int
}

func topCandidates(candidates []Candidate, projects []Project) []ProjectResult {
	// Build candidate skill map: candidateID -> skill -> proficiency
	candSkills := make(map[int]map[string]int)
	for _, c := range candidates {
		if candSkills[c.ID] == nil {
			candSkills[c.ID] = make(map[string]int)
		}
		candSkills[c.ID][c.Skill] = c.Proficiency
	}

	// Build project skill map: projectID -> []{skill, importance}
	projSkills := make(map[int][]struct {
		skill      string
		importance int
	})
	projReqCount := make(map[int]int)
	for _, p := range projects {
		projSkills[p.ID] = append(projSkills[p.ID], struct {
			skill      string
			importance int
		}{p.Skill, p.Importance})
		projReqCount[p.ID]++
	}

	type result struct {
		projectID   int
		candidateID int
		score       int
	}
	var allResults []result

	for projID, reqs := range projSkills {
		for candID, skills := range candSkills {
			score := 100
			matched := 0
			qualified := true
			for _, req := range reqs {
				prof, ok := skills[req.skill]
				if !ok {
					qualified = false
					break
				}
				if prof > req.importance {
					score += 10
				} else if prof < req.importance {
					score -= 5
				}
				matched++
			}
			if qualified && matched == projReqCount[projID] {
				allResults = append(allResults, result{projID, candID, score})
			}
		}
	}

	// Sort by project, then by score desc, then candidate id asc
	sort.Slice(allResults, func(i, j int) bool {
		if allResults[i].projectID != allResults[j].projectID {
			return allResults[i].projectID < allResults[j].projectID
		}
		if allResults[i].score != allResults[j].score {
			return allResults[i].score > allResults[j].score
		}
		return allResults[i].candidateID < allResults[j].candidateID
	})

	// Pick top candidate per project
	var out []ProjectResult
	seen := make(map[int]bool)
	for _, r := range allResults {
		if !seen[r.projectID] {
			seen[r.projectID] = true
			out = append(out, ProjectResult{r.projectID, r.candidateID, r.score})
		}
	}
	return out
}
```

## 3281 — Maximize Score Of Numbers In Ranges

```go
package main

// LeetCode #3281: Maximize Score of Numbers in Ranges
// https://leetcode.com/problems/maximize-score-of-numbers-in-ranges/
// Difficulty: Medium
// Time: O(n log n + n log D) Space: O(1)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxPossibleScore([]int{2, 6, 13, 13}, 5)) // 5
	fmt.Println(maxPossibleScore([]int{1, 2, 3, 4, 5}, 3)) // 1
	fmt.Println(maxPossibleScore([]int{6, 0, 3}, 2))       // 4
}

func maxPossibleScore(start []int, d int) int {
	sort.Ints(start)
	n := len(start)

	lo, hi := 0, start[n-1]+d-start[0]
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if possible(start, d, mid) {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return hi
}

func possible(start []int, d int, score int) bool {
	prev := start[0]
	for i := 1; i < len(start); i++ {
		if prev+score > start[i]+d {
			return false
		}
		if prev+score > start[i] {
			prev = prev + score
		} else {
			prev = start[i]
		}
	}
	return true
}
```

## 3282 — Reach End Of Array With Max Score

```go
package main

// LeetCode #3282: Reach End of Array With Max Score
// https://leetcode.com/problems/reach-end-of-array-with-max-score/
// Difficulty: Medium
// Time: O(n) Space: O(1)

import "fmt"

func main() {
	fmt.Println(findMaximumScore([]int{1, 3, 1, 5}))   // 7
	fmt.Println(findMaximumScore([]int{4, 3, 1, 3, 2})) // 16
	fmt.Println(findMaximumScore([]int{2, 2, 2, 2}))    // 6
}

func findMaximumScore(nums []int) int64 {
	var res int64 = 0
	ma := nums[0]
	for i := 1; i < len(nums); i++ {
		res += int64(ma)
		if nums[i] > ma {
			ma = nums[i]
		}
	}
	return res
}
```

## 3284 — Sum Of Consecutive Subarrays

```go
package main

// LeetCode #3284: Sum of Consecutive Subarrays
// https://leetcode.com/problems/sum-of-consecutive-subarrays/
// Difficulty: Medium
// Time: O(n) Space: O(1)

import "fmt"

func main() {
	fmt.Println(getSum([]int{1, 2, 3}))    // 20
	fmt.Println(getSum([]int{1, 2, 3, 5})) // 25
	fmt.Println(getSum([]int{1, 2, 3, 4})) // 50
}

func getSum(nums []int) int {
	const mod = 1_000_000_007
	n := len(nums)
	f, g := 1, 1
	s, t := nums[0], nums[0]
	ans := nums[0]

	for i := 1; i < n; i++ {
		diff := nums[i] - nums[i-1]

		if diff == 1 {
			f++
			s += f * nums[i]
			ans = (ans + s) % mod
		} else {
			f = 1
			s = nums[i]
		}

		if diff == -1 {
			g++
			t += g * nums[i]
			ans = (ans + t) % mod
		} else {
			g = 1
			t = nums[i]
		}

		if diff != 1 && diff != -1 {
			ans = (ans + nums[i]) % mod
		}
	}

	return ans
}
```

## 3286 — Find A Safe Walk Through A Grid

```go
package main

// LeetCode #3286: Find a Safe Walk Through a Grid
// https://leetcode.com/problems/find-a-safe-walk-through-a-grid/
// Difficulty: Medium
// Time: O(m * n) Space: O(m * n)

import "fmt"

func main() {
	fmt.Println(findSafeWalk([][]int{{0, 1, 0}, {0, 1, 0}, {0, 0, 0}}, 1)) // true
	fmt.Println(findSafeWalk([][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}, 1)) // true
	fmt.Println(findSafeWalk([][]int{{0, 1, 1}, {1, 0, 1}, {1, 1, 0}}, 3)) // true
}

func findSafeWalk(grid [][]int, health int) bool {
	m, n := len(grid), len(grid[0])
	dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	// minHealthLost[i][j] = minimum health lost to reach (i,j)
	minLost := make([][]int, m)
	for i := range minLost {
		minLost[i] = make([]int, n)
		for j := range minLost[i] {
			minLost[i][j] = 1 << 30
		}
	}
	minLost[0][0] = grid[0][0]

	queue := [][2]int{{0, 0}}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		i, j := cur[0], cur[1]
		for _, d := range dirs {
			x, y := i+d[0], j+d[1]
			if x >= 0 && x < m && y >= 0 && y < n {
				cost := minLost[i][j] + grid[x][y]
				if cost < minLost[x][y] {
					minLost[x][y] = cost
					queue = append(queue, [2]int{x, y})
				}
			}
		}
	}

	return minLost[m-1][n-1] < health
}
```

## 3290 — Maximum Multiplication Score

```go
package main

// LeetCode #3290: Maximum Multiplication Score
// https://leetcode.com/problems/maximum-multiplication-score/
// Difficulty: Medium
// Time: O(n) Space: O(1)

import (
	"fmt"
)

func main() {
	fmt.Println(maxScore([]int{1, 2, 3, 4}, []int{5, 6, 7, 8}))          // 70
	fmt.Println(maxScore([]int{-1, -2, -3, -4}, []int{1, 2, 3, 4}))      // -20
	fmt.Println(maxScore([]int{3, 2, 1, 4}, []int{2, 3, 4, 5, 6}))       // 52
}

func maxScore(a []int, b []int) int64 {
	const negInf int64 = -1e18
	dp := [4]int64{negInf, negInf, negInf, negInf}

	for _, bi := range b {
		for i := 3; i >= 0; i-- {
			var prev int64
			if i > 0 {
				prev = dp[i-1]
			}
			val := prev + int64(a[i])*int64(bi)
			if val > dp[i] {
				dp[i] = val
			}
		}
	}

	return dp[3]
}
```

## 3291 — Minimum Number Of Valid Strings To Form Target I

```go
package main

// LeetCode #3291: Minimum Number of Valid Strings to Form Target I
// https://leetcode.com/problems/minimum-number-of-valid-strings-to-form-target-i/
// Difficulty: Medium
// Time: O(n * L) Space: O(total_chars + n) where L = average prefix length

import (
	"fmt"
)

func main() {
	fmt.Println(minValidStrings([]string{"abc", "aaaaa", "bcdef"}, "aabcdabc")) // 3
	fmt.Println(minValidStrings([]string{"ab", "bc", "cd"}, "abc"))            // 2
	fmt.Println(minValidStrings([]string{"a", "b", "c"}, "xyz"))               // -1
}

type trieNode struct {
	children [26]*trieNode
}

func minValidStrings(words []string, target string) int {
	root := &trieNode{}
	for _, w := range words {
		node := root
		for _, ch := range w {
			idx := ch - 'a'
			if node.children[idx] == nil {
				node.children[idx] = &trieNode{}
			}
			node = node.children[idx]
		}
	}

	n := len(target)
	inf := int(1e9)
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = inf
	}
	dp[0] = 0

	for i := 0; i < n; i++ {
		if dp[i] == inf {
			continue
		}
		node := root
		for j := i; j < n; j++ {
			idx := target[j] - 'a'
			if node.children[idx] == nil {
				break
			}
			node = node.children[idx]
			if dp[i]+1 < dp[j+1] {
				dp[j+1] = dp[i] + 1
			}
		}
	}

	if dp[n] == inf {
		return -1
	}
	return dp[n]
}
```

## 3293 — Calculate Product Final Price

```go
package main

// LeetCode #3293: Calculate Product Final Price
// https://leetcode.com/problems/calculate-product-final-price/
// Difficulty: Medium
// Time: O(p + d) Space: O(p)

import (
	"fmt"
	"sort"
)

func main() {
	products := []Product{
		{1, "Electronics", 1000},
		{2, "Clothing", 50},
		{3, "Electronics", 1200},
		{4, "Home", 500},
	}
	discounts := []Discount{
		{"Electronics", 10},
		{"Clothing", 20},
	}
	fmt.Println(calculateFinalPrice(products, discounts))
	// Expected: [{1 900 Electronics} {2 40 Clothing} {3 1080 Electronics} {4 500 Home}]

	// Test 2: No discounts
	fmt.Println(calculateFinalPrice(
		[]Product{{1, "Food", 100}},
		[]Discount{{"Electronics", 10}},
	))
	// Expected: [{1 100 Food}]
}

type Product struct {
	ID       int
	Category string
	Price    float64
}

type Discount struct {
	Category string
	Percent  int
}

type ProductResult struct {
	ID         int
	FinalPrice float64
	Category   string
}

func calculateFinalPrice(products []Product, discounts []Discount) []ProductResult {
	discMap := make(map[string]int)
	for _, d := range discounts {
		discMap[d.Category] = d.Percent
	}

	res := make([]ProductResult, len(products))
	for i, p := range products {
		fp := p.Price
		if perc, ok := discMap[p.Category]; ok {
			fp = p.Price * (1.0 - float64(perc)/100.0)
		}
		res[i] = ProductResult{p.ID, fp, p.Category}
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i].ID < res[j].ID
	})
	return res
}
```

## 3294 — Convert Doubly Linked List To Array Ii

```go
package main

// LeetCode #3294: Convert Doubly Linked List to Array II
// https://leetcode.com/problems/convert-doubly-linked-list-to-array-ii/
// Difficulty: Medium
// Time: O(n) Space: O(n) for output

import "fmt"

func main() {
	// Build list: 1 <-> 2 <-> 3 <-> 4 <-> 5
	nodes := make([]*DListNode, 5)
	for i := 0; i < 5; i++ {
		nodes[i] = &DListNode{Val: i + 1}
	}
	for i := 0; i < 5; i++ {
		if i > 0 {
			nodes[i].Prev = nodes[i-1]
		}
		if i < 4 {
			nodes[i].Next = nodes[i+1]
		}
	}
	fmt.Println(toArray(nodes[2])) // [1 2 3 4 5]

	// Single node
	single := &DListNode{Val: 42}
	fmt.Println(toArray(single)) // [42]

	// Two nodes
	twoA := &DListNode{Val: 10}
	twoB := &DListNode{Val: 20}
	twoA.Next = twoB
	twoB.Prev = twoA
	fmt.Println(toArray(twoB)) // [10 20]
}

type DListNode struct {
	Val  int
	Next *DListNode
	Prev *DListNode
}

func toArray(node *DListNode) []int {
	if node == nil {
		return []int{}
	}

	// Find head
	head := node
	for head.Prev != nil {
		head = head.Prev
	}

	// Collect values
	var res []int
	for cur := head; cur != nil; cur = cur.Next {
		res = append(res, cur.Val)
	}
	return res
}
```

## 3295 — Report Spam Message

```go
package main

// LeetCode #3295: Report Spam Message
// https://leetcode.com/problems/report-spam-message/
// Difficulty: Medium
// Time: O(n + m) Space: O(m)

import "fmt"

func main() {
	fmt.Println(reportSpam([]string{"hello", "world", "leetcode"}, []string{"world", "hello"})) // true
	fmt.Println(reportSpam([]string{"hello", "programming", "fun"}, []string{"world", "hello"})) // false
	fmt.Println(reportSpam([]string{"a", "b", "c", "d"}, []string{"a", "b", "x"}))               // true
}

func reportSpam(message []string, bannedWords []string) bool {
	banned := make(map[string]struct{}, len(bannedWords))
	for _, w := range bannedWords {
		banned[w] = struct{}{}
	}
	count := 0
	for _, w := range message {
		if _, ok := banned[w]; ok {
			count++
			if count >= 2 {
				return true
			}
		}
	}
	return false
}
```

## 3296 — Minimum Number Of Seconds To Make Mountain Height Zero

```go
package main

// LeetCode #3296: Minimum Number of Seconds to Make Mountain Height Zero
// https://leetcode.com/problems/minimum-number-of-seconds-to-make-mountain-height-zero/
// Difficulty: Medium
// Time: O(n log T) Space: O(1)

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minNumberOfSeconds(4, []int{2, 1, 1}))  // 3
	fmt.Println(minNumberOfSeconds(10, []int{3, 2, 2, 4})) // 12
	fmt.Println(minNumberOfSeconds(5, []int{1}))            // 15
}

func minNumberOfSeconds(mountainHeight int, workerTimes []int) int64 {
	lo, hi := int64(0), int64(1e18)

	for lo < hi {
		mid := lo + (hi-lo)/2
		if canReduce(mountainHeight, workerTimes, mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func canReduce(height int, workerTimes []int, t int64) bool {
	var total int64
	for _, wt := range workerTimes {
		// Solve: wt * x * (x+1) / 2 <= t
		// x^2 + x - 2t/wt <= 0
		// x = floor((sqrt(1 + 8*t/wt) - 1) / 2)
		d := math.Sqrt(1.0 + 8.0*float64(t)/float64(wt))
		x := int64((d - 1.0) / 2.0)
		total += x
		if total >= int64(height) {
			return true
		}
	}
	return total >= int64(height)
}
```

## 3297 — Count Substrings That Can Be Rearranged To Contain A String I

```go
package main

// LeetCode #3297: Count Substrings That Can Be Rearranged to Contain a String I
// https://leetcode.com/problems/count-substrings-that-can-be-rearranged-to-contain-a-string-i/
// Difficulty: Medium
// Time: O(n + m) Space: O(1)

import "fmt"

func main() {
	fmt.Println(validSubstringCount("bcca", "abc")) // 1
	fmt.Println(validSubstringCount("abcabc", "abc")) // 10
	fmt.Println(validSubstringCount("a", "aa"))       // 0
}

func validSubstringCount(word1 string, word2 string) int64 {
	if len(word1) < len(word2) {
		return 0
	}

	var cnt [26]int
	need := 0
	for _, ch := range word2 {
		idx := ch - 'a'
		if cnt[idx] == 0 {
			need++
		}
		cnt[idx]++
	}

	var win [26]int
	var ans int64
	left := 0

	for _, ch := range word1 {
		idx := int(ch - 'a')
		win[idx]++
		if win[idx] == cnt[idx] {
			need--
		}

		for need == 0 {
			leftIdx := int(word1[left] - 'a')
			if win[leftIdx] == cnt[leftIdx] {
				need++
			}
			win[leftIdx]--
			left++
		}

		ans += int64(left)
	}

	return ans
}
```

## 3301 — Maximize The Total Height Of Unique Towers

```go
package main

// LeetCode #3301: Maximize the Total Height of Unique Towers
// https://leetcode.com/problems/maximize-the-total-height-of-unique-towers/
// Difficulty: Medium
// Time: O(n log n) Space: O(1)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximumTotalSum([]int{2, 3, 4, 3}))   // 10
	fmt.Println(maximumTotalSum([]int{2, 2, 1}))      // -1
	fmt.Println(maximumTotalSum([]int{5, 4, 3, 2, 1})) // 15
}

func maximumTotalSum(maximumHeight []int) int64 {
	sort.Slice(maximumHeight, func(i, j int) bool {
		return maximumHeight[i] > maximumHeight[j]
	})

	var total int64
	prev := maximumHeight[0]
	total += int64(prev)

	for i := 1; i < len(maximumHeight); i++ {
		if prev <= 1 {
			return -1
		}
		h := maximumHeight[i]
		if h >= prev {
			h = prev - 1
		}
		total += int64(h)
		prev = h
	}
	return total
}
```

## 3302 — Find The Lexicographically Smallest Valid Sequence

```go
package main

// LeetCode #3302: Find the Lexicographically Smallest Valid Sequence
// https://leetcode.com/problems/find-the-lexicographically-smallest-valid-sequence/
// Difficulty: Medium
// Time: O(n + m) Space: O(n)

import "fmt"

func main() {
	fmt.Println(validSequence("abc", "ab"))   // [0 1]
	fmt.Println(validSequence("abc", "ad"))   // [0 2]
	fmt.Println(validSequence("abbc", "abc")) // [0 1 3]
}

func validSequence(word1 string, word2 string) []int {
	n, m := len(word1), len(word2)
	suf := make([]int, n+1)
	suf[n] = m
	j := m - 1
	for i := n - 1; i >= 0; i-- {
		if j >= 0 && word1[i] == word2[j] {
			j--
		}
		suf[i] = j + 1
	}

	ans := []int{}
	changed := false
	j = 0
	for i := 0; i < n && j < m; i++ {
		if word1[i] == word2[j] {
			ans = append(ans, i)
			j++
		} else if !changed && suf[i+1] <= j+1 {
			changed = true
			ans = append(ans, i)
			j++
		}
	}

	if j < m {
		return []int{}
	}
	return ans
}
```

## 3305 — Count Of Substrings Containing Every Vowel And K Consonants I

```go
package main

// LeetCode #3305: Count of Substrings Containing Every Vowel and K Consonants I
// https://leetcode.com/problems/count-of-substrings-containing-every-vowel-and-k-consonants-i/
// Difficulty: Medium
// Time: O(n^2) Space: O(1)

import "fmt"

func main() {
	fmt.Println(countOfSubstrings("aeioqq", 1))         // 0
	fmt.Println(countOfSubstrings("ieaouqqieaouqq", 1)) // 3
	fmt.Println(countOfSubstrings("aeiou", 0))          // 1
}

func countOfSubstrings(word string, k int) int64 {
	n := len(word)
	isVowel := func(c byte) bool {
		return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
	}

	var ans int64
	for i := 0; i < n; i++ {
		vowelCnt := make(map[byte]int)
		cons := 0
		for j := i; j < n; j++ {
			if isVowel(word[j]) {
				vowelCnt[word[j]]++
			} else {
				cons++
				if cons > k {
					break
				}
			}
			if cons == k && len(vowelCnt) == 5 {
				ans++
			}
		}
	}
	return ans
}
```

## 3306 — Count Of Substrings Containing Every Vowel And K Consonants Ii

```go
package main

// LeetCode #3306: Count of Substrings Containing Every Vowel and K Consonants II
// https://leetcode.com/problems/count-of-substrings-containing-every-vowel-and-k-consonants-ii/
// Difficulty: Medium
// Time: O(n) Space: O(1)

import "fmt"

func main() {
	fmt.Println(countOfSubstringsII("aeioqq", 1))         // 0
	fmt.Println(countOfSubstringsII("ieaouqqieaouqq", 1)) // 3
	fmt.Println(countOfSubstringsII("aeiou", 0))          // 1
}

func countOfSubstringsII(word string, k int) int64 {
	return atLeastK(word, k) - atLeastK(word, k+1)
}

func atLeastK(word string, k int) int64 {
	n := len(word)
	isVowel := func(c byte) bool {
		return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
	}

	vowelCnt := make(map[byte]int)
	cons := 0
	var ans int64
	left := 0

	for right := 0; right < n; right++ {
		c := word[right]
		if isVowel(c) {
			vowelCnt[c]++
		} else {
			cons++
		}

		for len(vowelCnt) == 5 && cons >= k {
			out := word[left]
			if isVowel(out) {
				vowelCnt[out]--
				if vowelCnt[out] == 0 {
					delete(vowelCnt, out)
				}
			} else {
				cons--
			}
			left++
		}
		ans += int64(left)
	}
	return ans
}
```

## 3308 — Find Top Performing Driver

```go
package main

// LeetCode #3308: Find Top Performing Driver
// https://leetcode.com/problems/find-top-performing-driver/
// Difficulty: Medium
// Time: O(d + v + t) Space: O(d + v)

import (
	"fmt"
	"sort"
)

func main() {
	drivers := []Driver{{1, "Gasoline"}, {2, "Gasoline"}, {3, "Diesel"}}
	vehicles := []Vehicle{{1, 1}, {2, 2}, {3, 1}}
	trips := []Trip{{1, 1, 4.5, 100}, {2, 1, 4.0, 200}, {3, 2, 5.0, 150}, {4, 3, 3.5, 50}}
	fmt.Println(topPerformingDriver(drivers, vehicles, trips))

	drivers2 := []Driver{{1, "Electric"}, {2, "Electric"}}
	vehicles2 := []Vehicle{{1, 1}, {2, 2}}
	trips2 := []Trip{{1, 1, 5.0, 100}, {2, 2, 5.0, 80}}
	fmt.Println(topPerformingDriver(drivers2, vehicles2, trips2))
}

type Driver struct {
	ID       int
	FuelType string
}

type Vehicle struct {
	ID       int
	DriverID int
}

type Trip struct {
	ID       int
	VehicleID int
	Rating   float64
	Distance int
}

type FuelRank struct {
	FuelType   string
	DriverID   int
	Rating     float64
	Distance   int
	Accidents  int
}

func topPerformingDriver(drivers []Driver, vehicles []Vehicle, trips []Trip) []FuelRank {
	// Build driver -> fuel type map
	driverFuel := make(map[int]string)
	for _, d := range drivers {
		driverFuel[d.ID] = d.FuelType
	}

	// Build vehicle -> driver map
	vehicleDriver := make(map[int]int)
	driverVehicles := make(map[int][]int)
	for _, v := range vehicles {
		vehicleDriver[v.ID] = v.DriverID
		driverVehicles[v.DriverID] = append(driverVehicles[v.DriverID], v.ID)
	}

	// Group trips by vehicle, then by driver
	type stats struct {
		sumRating float64
		count     int
		distance  int
	}
	driverStats := make(map[int]*stats)
	for _, t := range trips {
		dID := vehicleDriver[t.VehicleID]
		if driverStats[dID] == nil {
			driverStats[dID] = &stats{}
		}
		driverStats[dID].sumRating += t.Rating
		driverStats[dID].count++
		driverStats[dID].distance += t.Distance
	}

	type candidate struct {
		fuelType  string
		driverID  int
		rating    float64
		distance  int
		accidents int
	}

	fuelCands := make(map[string][]candidate)
	for dID, s := range driverStats {
		avgRating := s.sumRating / float64(s.count)
		avgRating = float64(int(avgRating*100)) / 100 // round to 2 decimals
		ft := driverFuel[dID]
		fuelCands[ft] = append(fuelCands[ft], candidate{
			fuelType: ft, driverID: dID,
			rating: avgRating, distance: s.distance,
		})
	}

	var result []FuelRank
	for ft, cands := range fuelCands {
		sort.Slice(cands, func(i, j int) bool {
			if cands[i].rating != cands[j].rating {
				return cands[i].rating > cands[j].rating
			}
			if cands[i].distance != cands[j].distance {
				return cands[i].distance > cands[j].distance
			}
			return cands[i].driverID < cands[j].driverID
		})
		best := cands[0]
		result = append(result, FuelRank{
			FuelType: ft, DriverID: best.driverID,
			Rating: best.rating, Distance: best.distance,
		})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].FuelType < result[j].FuelType
	})
	return result
}
```

## 3309 — Maximum Possible Number By Binary Concatenation

```go
package main

// LeetCode #3309: Maximum Possible Number by Binary Concatenation
// https://leetcode.com/problems/maximum-possible-number-by-binary-concatenation/
// Difficulty: Medium
// Time: O(1) Space: O(1)

import "fmt"

func main() {
	fmt.Println(maxGoodNumber([]int{1, 2, 3}))    // 30
	fmt.Println(maxGoodNumber([]int{2, 8, 16}))   // 1296
	fmt.Println(maxGoodNumber([]int{1, 1, 1}))    // 7
}

func maxGoodNumber(nums []int) int {
	// Try all 6 permutations
	perms := [][]int{
		{nums[0], nums[1], nums[2]},
		{nums[0], nums[2], nums[1]},
		{nums[1], nums[0], nums[2]},
		{nums[1], nums[2], nums[0]},
		{nums[2], nums[0], nums[1]},
		{nums[2], nums[1], nums[0]},
	}

	maxVal := 0
	for _, p := range perms {
		val := 0
		for _, x := range p {
			bits := 0
			temp := x
			for temp > 0 {
				bits++
				temp >>= 1
			}
			if x == 0 {
				bits = 1
			}
			val = (val << bits) | x
		}
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}
```

## 3310 — Remove Methods From Project

```go
package main

// LeetCode #3310: Remove Methods From Project
// https://leetcode.com/problems/remove-methods-from-project/
// Difficulty: Medium
// Time: O(n + m) Space: O(n + m)

import "fmt"

func main() {
	fmt.Println(remainingMethods(4, 1, [][]int{{1, 2}, {0, 1}, {2, 3}})) // [0]
	fmt.Println(remainingMethods(5, 0, [][]int{{1, 2}, {0, 2}, {0, 1}, {3, 4}})) // [3 4]
	fmt.Println(remainingMethods(3, 2, [][]int{{0, 1}, {1, 2}, {2, 0}})) // []
}

func remainingMethods(n int, k int, invocations [][]int) []int {
	adj := make([][]int, n)
	for _, inv := range invocations {
		a, b := inv[0], inv[1]
		adj[a] = append(adj[a], b)
	}

	// DFS to find all suspicious methods
	suspicious := make([]bool, n)
	var dfs func(u int)
	dfs = func(u int) {
		if suspicious[u] {
			return
		}
		suspicious[u] = true
		for _, v := range adj[u] {
			dfs(v)
		}
	}
	dfs(k)

	// Check if any non-suspicious method calls a suspicious one
	for _, inv := range invocations {
		a, b := inv[0], inv[1]
		if !suspicious[a] && suspicious[b] {
			// Cannot remove - return all methods
			res := make([]int, n)
			for i := 0; i < n; i++ {
				res[i] = i
			}
			return res
		}
	}

	// Return non-suspicious methods
	var res []int
	for i := 0; i < n; i++ {
		if !suspicious[i] {
			res = append(res, i)
		}
	}
	return res
}
```

## 3315 — Construct The Minimum Bitwise Array Ii

```go
package main

// LeetCode #3315: Construct the Minimum Bitwise Array II
// https://leetcode.com/problems/construct-the-minimum-bitwise-array-ii/
// Difficulty: Medium
// Time: O(n log m) Space: O(1) (excluding output)

import "fmt"

func main() {
	fmt.Println(minBitwiseArray([]int{11, 13, 31})) // [9 12 15]
	fmt.Println(minBitwiseArray([]int{2, 3, 5}))    // [-1 1 4]
	fmt.Println(minBitwiseArray([]int{7}))           // [3]
}

func minBitwiseArray(nums []int) []int {
	ans := make([]int, len(nums))
	for i, num := range nums {
		if num == 2 {
			ans[i] = -1
			continue
		}
		// Find rightmost block of 1s in binary
		p := 0
		for (num>>p)&1 == 1 {
			p++
		}
		ans[i] = num ^ (1 << (p - 1))
	}
	return ans
}
```

## 3316 — Find Maximum Removals From Source String

```go
package main

// LeetCode #3316: Find Maximum Removals From Source String
// https://leetcode.com/problems/find-maximum-removals-from-source-string/
// Difficulty: Medium
// Time: O(n * m) Space: O(m)

import "fmt"

func main() {
	fmt.Println(maxRemovals("abc", "ab", []int{0, 1}))                          // 0
	fmt.Println(maxRemovals("abbaa", "aba", []int{0, 1, 2, 3, 4}))              // 2
	fmt.Println(maxRemovals("abcde", "ace", []int{1, 2, 3}))                    // 1
}

func maxRemovals(source string, pattern string, targetIndices []int) int {
	n, m := len(source), len(pattern)
	target := make([]bool, n)
	for _, idx := range targetIndices {
		target[idx] = true
	}

	dp := make([]int, m+1)
	for i := range dp {
		dp[i] = -1_000_000_000
	}
	dp[m] = 0

	for i := n - 1; i >= 0; i-- {
		for j := 0; j <= m; j++ {
			if target[i] {
				dp[j]++
			}
			if j < m && source[i] == pattern[j] {
				if dp[j]+1 > dp[j] && dp[j+1] > dp[j] {
					if dp[j+1] > dp[j] {
						dp[j] = dp[j+1]
					}
				} else if dp[j+1] > dp[j] {
					dp[j] = dp[j+1]
				}
			}
		}
	}

	return dp[0]
}
```

## 3319 — K Th Largest Perfect Subtree Size In Binary Tree

```go
package main

// LeetCode #3319: K-th Largest Perfect Subtree Size in Binary Tree
// https://leetcode.com/problems/k-th-largest-perfect-subtree-size-in-binary-tree/
// Difficulty: Medium
// Time: O(n) Space: O(n)

import (
	"fmt"
	"sort"
)

func main() {
	// Tree: [1,2,3,4,5,6,7]
	root := &TreeNode{Val: 1,
		Left: &TreeNode{Val: 2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 3,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 7}}}
	fmt.Println(kthLargestPerfectSubtree(root, 1)) // 7
	fmt.Println(kthLargestPerfectSubtree(root, 3)) // 3
	fmt.Println(kthLargestPerfectSubtree(root, 5)) // -1
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargestPerfectSubtree(root *TreeNode, k int) int {
	var sizes []int
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		if left < 0 || left != right {
			return -1
		}
		cur := left + right + 1
		sizes = append(sizes, cur)
		return cur
	}
	dfs(root)

	if len(sizes) < k {
		return -1
	}
	sort.Slice(sizes, func(i, j int) bool { return sizes[i] > sizes[j] })
	return sizes[k-1]
}
```

## 3322 — Premier League Table Ranking Iii

```go
package main

// LeetCode #3322: Premier League Table Ranking III
// https://leetcode.com/problems/premier-league-table-ranking-iii/
// Difficulty: Medium
// Time: O(t log t) Space: O(t)

import (
	"fmt"
	"sort"
)

func main() {
	stats := []SeasonStats{
		{1, 1, "City", 38, 28, 6, 4, 80, 20},
		{1, 2, "United", 38, 22, 8, 8, 60, 30},
		{2, 1, "City", 38, 26, 5, 7, 75, 25},
		{2, 2, "Arsenal", 38, 26, 5, 7, 70, 20},
	}
	fmt.Println(premierLeagueRanking(stats))
}

type SeasonStats struct {
	SeasonID       int
	TeamID         int
	TeamName       string
	MatchesPlayed  int
	Wins           int
	Draws          int
	Losses         int
	GoalsFor       int
	GoalsAgainst   int
}

type TeamRank struct {
	SeasonID       int
	TeamID         int
	TeamName       string
	Points         int
	GoalDifference int
}

func premierLeagueRanking(stats []SeasonStats) []TeamRank {
	type teamData struct {
		SeasonID int
		TeamID   int
		Name     string
		Points   int
		GD       int
	}
	var data []teamData
	for _, s := range stats {
		points := s.Wins*3 + s.Draws
		gd := s.GoalsFor - s.GoalsAgainst
		data = append(data, teamData{s.SeasonID, s.TeamID, s.TeamName, points, gd})
	}

	// Group by season and rank
	seasonTeams := make(map[int][]teamData)
	for _, d := range data {
		seasonTeams[d.SeasonID] = append(seasonTeams[d.SeasonID], d)
	}

	var result []TeamRank
	for _, teams := range seasonTeams {
		sort.Slice(teams, func(i, j int) bool {
			if teams[i].Points != teams[j].Points {
				return teams[i].Points > teams[j].Points
			}
			if teams[i].GD != teams[j].GD {
				return teams[i].GD > teams[j].GD
			}
			return teams[i].Name < teams[j].Name
		})
		for _, t := range teams {
			result = append(result, TeamRank{
				SeasonID: t.SeasonID, TeamID: t.TeamID, TeamName: t.Name,
				Points: t.Points, GoalDifference: t.GD,
			})
		}
	}

	sort.Slice(result, func(i, j int) bool {
		if result[i].SeasonID != result[j].SeasonID {
			return result[i].SeasonID < result[j].SeasonID
		}
		return result[i].TeamName < result[j].TeamName
	})
	return result
}
```

## 3323 — Minimize Connected Groups By Inserting Interval

```go
package main

// LeetCode #3323: Minimize Connected Groups by Inserting Interval
// https://leetcode.com/problems/minimize-connected-groups-by-inserting-interval/
// Difficulty: Medium
// Time: O(n log n) Space: O(n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minConnectedGroups([][]int{{1, 3}, {5, 6}, {8, 10}}, 3)) // 2
	fmt.Println(minConnectedGroups([][]int{{1, 2}, {3, 4}, {5, 6}}, 1)) // 2
	fmt.Println(minConnectedGroups([][]int{{1, 10}}, 5))               // 1
}

func minConnectedGroups(intervals [][]int, k int) int {
	if len(intervals) == 0 {
		return 0
	}

	// Sort by start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// Merge overlapping intervals
	merged := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		last := merged[len(merged)-1]
		if intervals[i][0] <= last[1] {
			if intervals[i][1] > last[1] {
				last[1] = intervals[i][1]
			}
		} else {
			merged = append(merged, intervals[i])
		}
	}

	n := len(merged)
	if n <= 1 {
		return 1
	}

	// For each group, try to bridge as many groups as possible
	minGroups := n
	j := 0
	for i := 0; i < n; i++ {
		for j < n && merged[j][0]-merged[i][1]-1 <= k {
			j++
		}
		bridged := j - i - 1
		remaining := n - bridged
		if remaining < minGroups {
			minGroups = remaining
		}
	}

	return minGroups
}
```

## 3324 — Find The Sequence Of Strings Appeared On The Screen

```go
package main

// LeetCode #3324: Find the Sequence of Strings Appeared on the Screen
// https://leetcode.com/problems/find-the-sequence-of-strings-appeared-on-the-screen/
// Difficulty: Medium
// Time: O(n * 26) Space: O(n * 26) for output

import "fmt"

func main() {
	fmt.Println(stringSequence("abc")) // [a aa ab aba abb abc]
	fmt.Println(stringSequence("ab"))  // [a aa ab]
	fmt.Println(stringSequence("z"))   // [a b c d e f g h i j k l m n o p q r s t u v w x y z]
}

func stringSequence(target string) []string {
	var result []string
	var cur []byte

	for _, ch := range target {
		cur = append(cur, 'a')
		result = append(result, string(cur))
		for cur[len(cur)-1] != byte(ch) {
			cur[len(cur)-1]++
			result = append(result, string(cur))
		}
	}

	return result
}
```

## 3325 — Count Substrings With K Frequency Characters I

```go
package main

// LeetCode #3325: Count Substrings With K-Frequency Characters I
// https://leetcode.com/problems/count-substrings-with-k-frequency-characters-i/
// Difficulty: Medium
// Time: O(n) Space: O(1)

import "fmt"

func main() {
	fmt.Println(numberOfSubstrings("abacb", 2)) // 4
	fmt.Println(numberOfSubstrings("abcde", 1)) // 15
	fmt.Println(numberOfSubstrings("aaaa", 2))  // 3
}

func numberOfSubstrings(s string, k int) int {
	n := len(s)
	cnt := [26]int{}
	ans, left := 0, 0

	for right := 0; right < n; right++ {
		idx := s[right] - 'a'
		cnt[idx]++

		for cnt[idx] >= k {
			ans += n - right
			cnt[s[left]-'a']--
			left++
		}
	}

	return ans
}
```

## 3326 — Minimum Division Operations To Make Array Non Decreasing

```go
package main

// LeetCode #3326: Minimum Division Operations to Make Array Non Decreasing
// https://leetcode.com/problems/minimum-division-operations-to-make-array-non-decreasing/
// Difficulty: Medium
// Time: O(M log log M + n) Space: O(M) where M = 1e6

import "fmt"

func main() {
	fmt.Println(minOperations([]int{25, 7}))       // 1
	fmt.Println(minOperations([]int{7, 7, 6}))     // -1
	fmt.Println(minOperations([]int{1, 1, 1, 1}))  // 0
}

const mx = 1000001

var lpf [mx]int

func init() {
	for i := 2; i < mx; i++ {
		if lpf[i] == 0 {
			for j := i; j < mx; j += i {
				if lpf[j] == 0 {
					lpf[j] = i
				}
			}
		}
	}
}

func minOperations(nums []int) int {
	ans := 0
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] > nums[i+1] {
			if lpf[nums[i]] > nums[i+1] {
				return -1
			}
			nums[i] = lpf[nums[i]]
			ans++
		}
	}
	return ans
}
```

## 3328 — Find Cities In Each State Ii

```go
package main

// LeetCode #3328: Find Cities in Each State II
// https://leetcode.com/problems/find-cities-in-each-state-ii/
// Difficulty: Medium
// Time: O(c log c) Space: O(c)

import (
	"fmt"
	"sort"
)

func main() {
	cities := []CityInfo{
		{1, "NY", 1000000},
		{1, "LA", 500000},
		{2, "Chicago", 800000},
		{2, "Houston", 600000},
		{2, "Phoenix", 400000},
	}
	fmt.Println(findCityRanking(cities))
}

type CityInfo struct {
	StateID    int
	CityName   string
	Population int
}

type CityRank struct {
	StateID int
	CityName string
}

func findCityRanking(cities []CityInfo) []CityRank {
	stateCities := make(map[int][]CityInfo)
	for _, c := range cities {
		stateCities[c.StateID] = append(stateCities[c.StateID], c)
	}

	var result []CityRank
	for sid, cs := range stateCities {
		sort.Slice(cs, func(i, j int) bool {
			return cs[i].Population > cs[j].Population
		})
		for _, c := range cs {
			result = append(result, CityRank{sid, c.CityName})
		}
	}

	sort.Slice(result, func(i, j int) bool {
		if result[i].StateID != result[j].StateID {
			return result[i].StateID < result[j].StateID
		}
		return result[i].CityName < result[j].CityName
	})
	return result
}
```

## 3331 — Find Subtree Sizes After Changes

```go
package main

// LeetCode #3331: Find Subtree Sizes After Changes
// https://leetcode.com/problems/find-subtree-sizes-after-changes/
// Difficulty: Medium
// Time: O(n) Space: O(n)

import "fmt"

func main() {
	fmt.Println(findSubtreeSizes([]int{-1, 0, 0, 1, 1, 2}, "abacbe")) // [6 3 2 1 1 1]
	fmt.Println(findSubtreeSizes([]int{-1, 0, 0}, "abc"))             // [3 1 1]
}

func findSubtreeSizes(parent []int, s string) []int {
	n := len(parent)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		g[parent[i]] = append(g[parent[i]], i)
	}

	ans := make([]int, n)
	last := make([]int, 26)
	for i := range last {
		last[i] = -1
	}

	var dfs func(u int)
	dfs = func(u int) {
		old := last[s[u]-'a']
		last[s[u]-'a'] = u
		ans[u] = 1

		for _, v := range g[u] {
			dfs(v)
			p := last[s[v]-'a']
			if p == -1 {
				ans[u] += ans[v]
			} else {
				ans[p] += ans[v]
			}
		}

		last[s[u]-'a'] = old
	}

	dfs(0)
	return ans
}
```

## 3332 — Maximum Points Tourist Can Earn

```go
package main

// LeetCode #3332: Maximum Points Tourist Can Earn
// https://leetcode.com/problems/maximum-points-tourist-can-earn/
// Difficulty: Medium
// Time: O(k * n^2) Space: O(n)

import "fmt"

func main() {
	fmt.Println(maxScore(2, 1, [][]int{{2, 3}}, [][]int{{0, 2}, {1, 0}})) // 3
	fmt.Println(maxScore(3, 2, [][]int{{3, 4, 2}, {2, 1, 3}}, [][]int{{0, 2, 1}, {2, 0, 2}, {1, 3, 0}})) // 8
}

func maxScore(n int, k int, stayScore [][]int, travelScore [][]int) int {
	dp := make([]int, n)
	for i := 0; i < k; i++ {
		ndp := make([]int, n)
		// Copy dp and add stay score
		for curr := 0; curr < n; curr++ {
			ndp[curr] = dp[curr] + stayScore[i][curr]
		}
		// Try travel from any city to any city
		for curr := 0; curr < n; curr++ {
			for dest := 0; dest < n; dest++ {
				val := dp[curr] + travelScore[curr][dest]
				if val > ndp[dest] {
					ndp[dest] = val
				}
			}
		}
		dp = ndp
	}

	maxVal := 0
	for _, v := range dp {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}
```

## 3334 — Find The Maximum Factor Score Of Array

```go
package main

// LeetCode #3334: Find the Maximum Factor Score of Array
// https://leetcode.com/problems/find-the-maximum-factor-score-of-array/
// Difficulty: Medium
// Time: O(n) Space: O(n)

import "fmt"

func main() {
	fmt.Println(maxFactorScore([]int{2, 4, 8, 16})) // 64
	fmt.Println(maxFactorScore([]int{1, 2, 3, 4, 5})) // 60
	fmt.Println(maxFactorScore([]int{3}))             // 9
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func maxFactorScore(nums []int) int64 {
	n := len(nums)
	if n == 0 {
		return 0
	}

	preGCD := make([]int, n)
	preLCM := make([]int, n)
	preGCD[0] = nums[0]
	preLCM[0] = nums[0]
	for i := 1; i < n; i++ {
		preGCD[i] = gcd(preGCD[i-1], nums[i])
		preLCM[i] = lcm(preLCM[i-1], nums[i])
	}

	sufGCD := make([]int, n+1)
	sufLCM := make([]int, n+1)
	sufLCM[n] = 1
	for i := n - 1; i >= 0; i-- {
		sufGCD[i] = gcd(sufGCD[i+1], nums[i])
		sufLCM[i] = lcm(sufLCM[i+1], nums[i])
	}

	ans := int64(preGCD[n-1]) * int64(preLCM[n-1])

	for i := 0; i < n; i++ {
		var g int
		if i == 0 {
			g = sufGCD[1]
		} else {
			g = gcd(preGCD[i-1], sufGCD[i+1])
		}
		var l int
		if i == 0 {
			l = sufLCM[1]
		} else if i == n-1 {
			l = preLCM[n-2]
		} else {
			l = lcm(preLCM[i-1], sufLCM[i+1])
		}
		val := int64(g) * int64(l)
		if val > ans {
			ans = val
		}
	}

	return ans
}
```

## 3335 — Total Characters In String After Transformations I

```go
package main

// LeetCode #3335: Total Characters in String After Transformations I
// https://leetcode.com/problems/total-characters-in-string-after-transformations-i/
// Difficulty: Medium
// Time: O(n + 26t) Space: O(26)

import "fmt"

func main() {
	fmt.Println(lengthAfterTransformations("ab", 1)) // 2
	fmt.Println(lengthAfterTransformations("z", 1))  // 2
	fmt.Println(lengthAfterTransformations("az", 2)) // 5
}

func lengthAfterTransformations(s string, t int) int {
	const mod = 1_000_000_007
	freq := [26]int64{}
	for _, c := range s {
		freq[c-'a']++
	}

	for ; t > 0; t-- {
		next := [26]int64{}
		for i, cnt := range freq {
			if cnt == 0 {
				continue
			}
			if i == 25 { // 'z' -> "ab"
				next[0] = (next[0] + cnt) % mod
				next[1] = (next[1] + cnt) % mod
			} else {
				next[i+1] = (next[i+1] + cnt) % mod
			}
		}
		freq = next
	}

	var ans int64
	for _, cnt := range freq {
		ans = (ans + cnt) % mod
	}
	return int(ans)
}
```

## 3338 — Second Highest Salary Ii

```go
package main

// LeetCode #3338: Second Highest Salary II
// https://leetcode.com/problems/second-highest-salary-ii/
// Difficulty: Medium
// Time: O(e log e) Space: O(e)

import (
	"fmt"
	"sort"
)

func main() {
	employees := []Employee{
		{1, "IT", 100000},
		{2, "IT", 80000},
		{3, "IT", 80000},
		{4, "HR", 90000},
		{5, "HR", 70000},
		{6, "HR", 60000},
	}
	fmt.Println(secondHighestSalary(employees))
}

type Employee struct {
	ID     int
	Dept   string
	Salary int
}

type DeptSalary struct {
	Dept  string
	Salary int
}

func secondHighestSalary(employees []Employee) []DeptSalary {
	deptSalaries := make(map[string][]int)
	for _, e := range employees {
		deptSalaries[e.Dept] = append(deptSalaries[e.Dept], e.Salary)
	}

	var result []DeptSalary
	for dept, salaries := range deptSalaries {
		sort.Slice(salaries, func(i, j int) bool {
			return salaries[i] > salaries[j]
		})
		// Find second highest distinct salary
		seen := 1
		second := -1
		for _, s := range salaries {
			if s < salaries[0] {
				if seen == 1 || s < second {
					second = s
					seen++
				}
				if second != s {
					continue
				}
			}
		}
		// Actually simpler: just find 2nd distinct
		distinct := []int{salaries[0]}
		for _, s := range salaries {
			if s != distinct[len(distinct)-1] {
				distinct = append(distinct, s)
			}
		}
		if len(distinct) >= 2 {
			result = append(result, DeptSalary{dept, distinct[1]})
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Dept < result[j].Dept
	})
	return result
}
```

## 3339 — Find The Number Of K Even Arrays

```go
package main

// LeetCode #3339: Find the Number of K-Even Arrays
// https://leetcode.com/problems/find-the-number-of-k-even-arrays/
// Difficulty: Medium
// Time: O(n * k) Space: O(n * k)

import (
	"fmt"
)

func main() {
	fmt.Println(countKEvenArrays(3, 4, 2)) // 8
	fmt.Println(countKEvenArrays(5, 1, 0)) // 1
	fmt.Println(countKEvenArrays(7, 7, 5)) // 5832
}

func countKEvenArrays(n int, m int, k int) int {
	const mod = 1_000_000_007

	evens := m / 2
	odds := m - evens

	// dp[pos][pairs][parity] where parity 0=even, 1=odd
	dp := make([][][]int, n+1)
	for i := range dp {
		dp[i] = make([][]int, k+2)
		for j := range dp[i] {
			dp[i][j] = make([]int, 2)
		}
	}

	// Base: first element (1-indexed)
	dp[1][0][0] = evens
	dp[1][0][1] = odds

	for i := 2; i <= n; i++ {
		for j := 0; j <= k; j++ {
			// Place even: adds pair if prev was even
			if evens > 0 {
				dp[i][j][0] = (dp[i-1][j][1] * evens) % mod
				if j > 0 {
					dp[i][j][0] = (dp[i][j][0] + dp[i-1][j-1][0]*evens) % mod
				}
			}
			// Place odd: never adds a pair
			if odds > 0 {
				dp[i][j][1] = (dp[i-1][j][0] + dp[i-1][j][1]) % mod
				dp[i][j][1] = (dp[i][j][1] * odds) % mod
			}
		}
	}

	return (dp[n][k][0] + dp[n][k][1]) % mod
}
```

## 3341 — Find Minimum Time To Reach Last Room I

```go
package main

// LeetCode #3341: Find Minimum Time to Reach Last Room I
// https://leetcode.com/problems/find-minimum-time-to-reach-last-room-i/
// Difficulty: Medium
// Time: O(m * n * log(m * n)) Space: O(m * n)

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(minTimeToReach([][]int{{0, 4}, {4, 4}})) // 6
	fmt.Println(minTimeToReach([][]int{{0, 0, 0}, {0, 0, 0}})) // 2
}

type State struct {
	time int
	r    int
	c    int
}

type MinHeap []State

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].time < h[j].time }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(State)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func minTimeToReach(moveTime [][]int) int {
	m, n := len(moveTime), len(moveTime[0])
	dist := make([][]int, m)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = 1 << 60
		}
	}
	dist[0][0] = 0

	h := &MinHeap{}
	heap.Init(h)
	heap.Push(h, State{0, 0, 0})
	dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for h.Len() > 0 {
		cur := heap.Pop(h).(State)
		if cur.time > dist[cur.r][cur.c] {
			continue
		}
		if cur.r == m-1 && cur.c == n-1 {
			return cur.time
		}
		for _, d := range dirs {
			nr, nc := cur.r+d[0], cur.c+d[1]
			if nr >= 0 && nr < m && nc >= 0 && nc < n {
				wait := 0
				if cur.time < moveTime[nr][nc] {
					wait = moveTime[nr][nc] - cur.time
				}
				nt := cur.time + 1 + wait
				if nt < dist[nr][nc] {
					dist[nr][nc] = nt
					heap.Push(h, State{nt, nr, nc})
				}
			}
		}
	}
	return -1
}
```

## 3342 — Find Minimum Time To Reach Last Room Ii

```go
package main

// LeetCode #3342: Find Minimum Time to Reach Last Room II
// https://leetcode.com/problems/find-minimum-time-to-reach-last-room-ii/
// Difficulty: Medium
// Time: O(m*n*log(m*n)) Space: O(m*n)

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(minTimeToReach([][]int{{0, 4}, {4, 4}})) // 8
}

type State2 struct {
	time int
	r    int
	c    int
}

type MinHeap2 []State2

func (h MinHeap2) Len() int           { return len(h) }
func (h MinHeap2) Less(i, j int) bool { return h[i].time < h[j].time }
func (h MinHeap2) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap2) Push(x interface{}) { *h = append(*h, x.(State2)) }
func (h *MinHeap2) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func minTimeToReach(moveTime [][]int) int {
	m, n := len(moveTime), len(moveTime[0])
	dist := make([][]int, m)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = 1 << 60
		}
	}
	dist[0][0] = 0

	h := &MinHeap2{}
	heap.Init(h)
	heap.Push(h, State2{0, 0, 0})
	dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for h.Len() > 0 {
		cur := heap.Pop(h).(State2)
		if cur.time > dist[cur.r][cur.c] {
			continue
		}
		if cur.r == m-1 && cur.c == n-1 {
			return cur.time
		}
		for _, d := range dirs {
			nr, nc := cur.r+d[0], cur.c+d[1]
			if nr >= 0 && nr < m && nc >= 0 && nc < n {
				wait := 0
				if cur.time < moveTime[nr][nc] {
					wait = moveTime[nr][nc] - cur.time
				}
				// Move cost: 2 if coming from even cell, 1 if from odd cell
				moveCost := (cur.r+cur.c)%2 + 1
				nt := cur.time + moveCost + wait
				if nt < dist[nr][nc] {
					dist[nr][nc] = nt
					heap.Push(h, State2{nt, nr, nc})
				}
			}
		}
	}
	return -1
}
```

## 3344 — Maximum Sized Array

```go
package main

// LeetCode #3344: Maximum Sized Array
// https://leetcode.com/problems/maximum-sized-array/
// Difficulty: Medium [Paid]
// Time: O(mx^2) Space: O(mx)

import "fmt"

func main() {
	fmt.Println(maxSizedArray(1))  // 1
	fmt.Println(maxSizedArray(10)) // 2
	fmt.Println(maxSizedArray(0))  // 1
}

func maxSizedArray(s int) int {
	// Sum_{i=0}^{n-1} i * Sum_{j=0}^{n-1} Sum_{k=0}^{n-1} (j|k)
	// = (n-1)*n/2 * orSum(n)
	// Upper bound: sum ~ (n-1)^5/4, s <= 10^15 => n <= ~1330
	mx := 1335

	// Precompute orSum[n] = sum of (j|k) for all 0 <= j,k < n
	orSum := make([]int64, mx+1)
	for n := 1; n <= mx; n++ {
		orSum[n] = orSum[n-1]
		j := n - 1
		for k := 0; k < j; k++ {
			val := int64(j | k)
			orSum[n] += 2 * val
		}
		orSum[n] += int64(j | j)
	}

	ans := 1
	for n := 1; n <= mx; n++ {
		sumI := int64(n-1) * int64(n) / 2
		total := orSum[n] * sumI
		if total <= int64(s) {
			ans = n
		} else {
			break
		}
	}
	return ans
}
```

## 3346 — Maximum Frequency Of An Element After Performing Operations I

```go
package main

// LeetCode #3346: Maximum Frequency of an Element After Performing Operations I
// https://leetcode.com/problems/maximum-frequency-of-an-element-after-performing-operations-i/
// Difficulty: Medium
// Time: O(n log n) Space: O(n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxFrequency([]int{1, 4, 5}, 1, 2)) // 2
	fmt.Println(maxFrequency([]int{5, 11, 20}, 5, 3)) // 2
}

func maxFrequency(nums []int, k int, numOperations int) int {
	sort.Ints(nums)

	// Count duplicates (existing frequency)
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}

	// Sliding window: for each value in sorted array,
	// find how many numbers can reach it within k operations
	n := len(nums)
	ans := 0
	left := 0
	right := 0

	// Collect unique values
	unique := make([]int, 0, len(freq))
	for v := range freq {
		unique = append(unique, v)
	}
	sort.Ints(unique)

	for _, v := range unique {
		// Expand right: all nums[right] <= v + k
		for right < n && nums[right] <= v+k {
			right++
		}
		// Expand left: all nums[left] < v - k
		for left < n && nums[left] < v-k {
			left++
		}
		total := right - left
		candidates := total - freq[v]
		ops := numOperations
		if candidates > ops {
			candidates = ops
		}
		cur := freq[v] + candidates
		if cur > ans {
			ans = cur
		}
	}

	// Also consider values not in nums (midpoints)
	// For any pair of values where the distance is <= 2*k,
	// operations can be split between them
	left = 0
	for right = 0; right < n; right++ {
		for nums[right]-nums[left] > 2*k {
			left++
		}
		total := right - left + 1
		if total > numOperations {
			total = numOperations
		}
		if total > ans {
			ans = total
		}
	}

	return ans
}
```

## 3350 — Adjacent Increasing Subarrays Detection Ii

```go
package main

// LeetCode #3350: Adjacent Increasing Subarrays Detection II
// https://leetcode.com/problems/adjacent-increasing-subarrays-detection-ii/
// Difficulty: Medium
// Time: O(n) Space: O(n)

import "fmt"

func main() {
	fmt.Println(maxIncreasingSubarrays([]int{2, 5, 7, 8, 9, 2, 3, 4, 3, 1})) // 3
	fmt.Println(maxIncreasingSubarrays([]int{1, 2, 3, 4, 4, 4, 4, 5, 6, 7})) // 2
}

func maxIncreasingSubarrays(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}

	// len[i] = length of strictly increasing subarray ending at i
	lenEnd := make([]int, n)
	lenEnd[0] = 1
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			lenEnd[i] = lenEnd[i-1] + 1
		} else {
			lenEnd[i] = 1
		}
	}

	// For each position i, find max k such that:
	// subarray [i-k+1, i] is increasing AND [i+1, i+k] is increasing
	ans := 0
	for i := 0; i < n-1; i++ {
		// Current increasing subarray ending at i has length lenEnd[i]
		// Next increasing subarray starting at i+1 has length...
		// We can compute the length of the next increasing subarray starting at i+1
		k := 1
		maxK := 0
		for k <= lenEnd[i] && i+k < n {
			if lenEnd[i+k] >= k {
				maxK = k
			}
			k++
		}
		if maxK > ans {
			ans = maxK
		}
	}

	return ans
}
```

## 3355 — Zero Array Transformation I

```go
package main

// LeetCode #3355: Zero Array Transformation I
// https://leetcode.com/problems/zero-array-transformation-i/
// Difficulty: Medium
// Time: O(n + q) Space: O(n)

import "fmt"

func main() {
	fmt.Println(isZeroArray([]int{1, 0, 1}, [][]int{{0, 2}, {0, 2}})) // true
	fmt.Println(isZeroArray([]int{2, 0, 2}, [][]int{{0, 2}, {0, 2}, {1, 1}})) // true
}

func isZeroArray(nums []int, queries [][]int) bool {
	n := len(nums)
	diff := make([]int, n+1)

	for _, q := range queries {
		l, r := q[0], q[1]
		diff[l]++
		diff[r+1]--
	}

	cur := 0
	for i := 0; i < n; i++ {
		cur += diff[i]
		if cur < nums[i] {
			return false
		}
	}
	return true
}
```

## 3356 — Zero Array Transformation Ii

```go
package main

// LeetCode #3356: Zero Array Transformation II
// https://leetcode.com/problems/zero-array-transformation-ii/
// Difficulty: Medium
// Time: O((n + q) log q) Space: O(n)

import "fmt"

func main() {
	fmt.Println(minZeroArray([]int{2, 0, 2}, [][]int{{0, 2, 1}, {0, 2, 1}, {1, 1, 3}})) // 2
	fmt.Println(minZeroArray([]int{4, 3, 2, 1}, [][]int{{1, 3, 2}, {0, 2, 1}}))         // -1
}

func minZeroArray(nums []int, queries [][]int) int {
	m := len(queries)

	// Check if already zero
	allZero := true
	for _, v := range nums {
		if v != 0 {
			allZero = false
			break
		}
	}
	if allZero {
		return 0
	}

	// Check if impossible
	if !canMakeZero(nums, queries, m) {
		return -1
	}

	lo, hi := 1, m
	for lo < hi {
		mid := (lo + hi) / 2
		if canMakeZero(nums, queries, mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func canMakeZero(nums []int, queries [][]int, k int) bool {
	n := len(nums)
	diff := make([]int, n+1)
	for i := 0; i < k; i++ {
		l, r, val := queries[i][0], queries[i][1], queries[i][2]
		diff[l] += val
		diff[r+1] -= val
	}
	cur := 0
	for i := 0; i < n; i++ {
		cur += diff[i]
		if cur < nums[i] {
			return false
		}
	}
	return true
}
```

## 3361 — Shift Distance Between Two Strings

```go
package main

// LeetCode #3361: Shift Distance Between Two Strings
// https://leetcode.com/problems/shift-distance-between-two-strings/
// Difficulty: Medium
// Time: O(n + 26) Space: O(26)

import "fmt"

func main() {
	fmt.Println(shiftDistance("ab", "cd", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26})) // 6
	fmt.Println(shiftDistance("leetcode", "leetcode", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26})) // 0
}

func shiftDistance(s string, t string, nextCost []int, previousCost []int) int64 {
	// Precompute prefix sums for forward (next) and backward (prev) shifts
	// forward[i][j] = cost to go from i to j going forward
	// backward[i][j] = cost to go from i to j going backward

	// Prefix sums for cyclic shifts
	nextPref := make([]int, 53) // double for wrap-around
	prevPref := make([]int, 53)
	for i := 0; i < 52; i++ {
		nextPref[i+1] = nextPref[i] + nextCost[i%26]
		prevPref[i+1] = prevPref[i] + previousCost[i%26]
	}

	var ans int64
	for i := 0; i < len(s); i++ {
		a := int(s[i] - 'a')
		b := int(t[i] - 'a')
		if a == b {
			continue
		}

		// Forward: a -> a+1 -> ... -> b (going forward, wrapping)
		forward := 0
		if b >= a {
			forward = nextPref[b] - nextPref[a]
		} else {
			forward = nextPref[a+26] - nextPref[a]
			forward -= nextPref[b+26] - nextPref[b]
			// Hmm, this is getting complex. Let me simplify.
		}

		// Actually, let me compute forward and backward costs directly.
		forward = 0
		cur := a
		for cur != b {
			forward += nextCost[cur]
			cur = (cur + 1) % 26
		}

		backward := 0
		cur = a
		for cur != b {
			backward += previousCost[cur]
			cur = (cur - 1 + 26) % 26
		}

		if forward < backward {
			ans += int64(forward)
		} else {
			ans += int64(backward)
		}
	}
	return ans
}
```

## 3362 — Zero Array Transformation Iii

```go
package main

// LeetCode #3362: Zero Array Transformation III
// https://leetcode.com/problems/zero-array-transformation-iii/
// Difficulty: Medium
// Time: O(n log n) Space: O(n)

import (
	"container/heap"
	"fmt"
	"sort"
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
	fmt.Println(maxRemoval([]int{2, 0, 2}, [][]int{{0, 2}, {0, 2}, {1, 1}})) // 1
}

func maxRemoval(nums []int, queries [][]int) int {
	n := len(nums)
	m := len(queries)

	// Sort queries by left endpoint
	sort.Slice(queries, func(i, j int) bool {
		return queries[i][0] < queries[j][0]
	})

	// For each position i, we need nums[i] decrements at i
	// Greedily use queries with farthest right endpoint
	h := &MaxHeap{}
	heap.Init(h)

	diff := make([]int, n+1)
	qi := 0
	cur := 0

	for i := 0; i < n; i++ {
		// Add all queries that start at i
		for qi < m && queries[qi][0] == i {
			heap.Push(h, queries[qi][1])
			qi++
		}

		cur += diff[i]
		need := nums[i] - cur

		for need > 0 && h.Len() > 0 {
			r := heap.Pop(h).(int)
			cur++
			diff[r+1]++
			need--
		}

		if need > 0 {
			return -1
		}
	}

	// Unused queries can be removed
	return h.Len()
}
```

## 3365 — Rearrange K Substrings To Form Target String

```go
package main

// LeetCode #3365: Rearrange K Substrings to Form Target String
// https://leetcode.com/problems/rearrange-k-substrings-to-form-target-string/
// Difficulty: Medium
// Time: O(n) Space: O(n)

import "fmt"

func main() {
	fmt.Println(isPossibleToRearrange("abcd", "cdab", 2)) // true
	fmt.Println(isPossibleToRearrange("aabb", "bbaa", 2)) // true
	fmt.Println(isPossibleToRearrange("abcd", "acbd", 2)) // false
}

func isPossibleToRearrange(s string, t string, k int) bool {
	n := len(s)
	m := n / k

	cnt := make(map[string]int)
	for i := 0; i < n; i += m {
		cnt[s[i:i+m]]++
		cnt[t[i:i+m]]--
	}

	for _, v := range cnt {
		if v != 0 {
			return false
		}
	}
	return true
}
```

## 3366 — Minimum Array Sum

```go
package main

// LeetCode #3366: Minimum Array Sum
// https://leetcode.com/problems/minimum-array-sum/
// Difficulty: Medium
// Time: O(n * op1 * op2) Space: O(op1 * op2)

import "fmt"

func main() {
	fmt.Println(minArraySum([]int{2, 8, 3, 19, 3}, 3, 1, 1)) // 23
	fmt.Println(minArraySum([]int{2, 4, 3}, 3, 2, 1))        // 3
}

func minArraySum(nums []int, k int, op1 int, op2 int) int {
	// dp[j1][j2] = min sum using j1 op1 and j2 op2
	dp := make([][]int, op1+1)
	for i := range dp {
		dp[i] = make([]int, op2+1)
		for j := range dp[i] {
			dp[i][j] = 1 << 60
		}
	}
	dp[0][0] = 0

	for _, x := range nums {
		ndp := make([][]int, op1+1)
		for i := range ndp {
			ndp[i] = make([]int, op2+1)
			for j := range ndp[i] {
				ndp[i][j] = 1 << 60
			}
		}

		for j1 := 0; j1 <= op1; j1++ {
			for j2 := 0; j2 <= op2; j2++ {
				if dp[j1][j2] >= 1<<60 {
					continue
				}
				cur := dp[j1][j2]

				// No operation
				if cur+x < ndp[j1][j2] {
					ndp[j1][j2] = cur + x
				}

				// Only op1 (halve, round up)
				if j1+1 <= op1 {
					val := cur + (x+1)/2
					if val < ndp[j1+1][j2] {
						ndp[j1+1][j2] = val
					}
				}

				// Only op2 (subtract k)
				if j2+1 <= op2 && x >= k {
					val := cur + x - k
					if val < ndp[j1][j2+1] {
						ndp[j1][j2+1] = val
					}
				}

				// Both ops
				if j1+1 <= op1 && j2+1 <= op2 && x >= k {
					// Order 1: subtract then halve
					v1 := (x - k + 1) / 2
					// Order 2: halve then subtract
					half := (x + 1) / 2
					v2 := half - k
					if v2 < 0 {
						v2 = 1 << 60
					}
					best := v1
					if v2 < best {
						best = v2
					}
					val := cur + best
					if val < ndp[j1+1][j2+1] {
						ndp[j1+1][j2+1] = val
					}
				}
			}
		}
		dp = ndp
	}

	ans := 1 << 60
	for j1 := 0; j1 <= op1; j1++ {
		for j2 := 0; j2 <= op2; j2++ {
			if dp[j1][j2] < ans {
				ans = dp[j1][j2]
			}
		}
	}
	return ans
}
```

## 3371 — Identify The Largest Outlier In An Array

```go
package main

// LeetCode #3371: Identify the Largest Outlier in an Array
// https://leetcode.com/problems/identify-the-largest-outlier-in-an-array/
// Difficulty: Medium
// Time: O(n) Space: O(n)

import "fmt"

func main() {
	fmt.Println(getLargestOutlier([]int{2, 3, 5, 10}))   // 10
	fmt.Println(getLargestOutlier([]int{-2, -1, -3, -6, 4})) // 4
	fmt.Println(getLargestOutlier([]int{1, 1, 1, 1, 1, 5, 5}))  // 5
}

func getLargestOutlier(nums []int) int {
	freq := make(map[int]int)
	total := 0
	for _, v := range nums {
		freq[v]++
		total += v
	}

	ans := -(1 << 60)
	for _, v := range nums {
		outlier := total - 2*v
		if outlier == v && freq[v] < 2 {
			continue
		}
		if _, ok := freq[outlier]; ok && outlier > ans {
			ans = outlier
		}
	}
	return ans
}
```

## 3372 — Maximize The Number Of Target Nodes After Connecting Trees I

```go
package main

// LeetCode #3372: Maximize the Number of Target Nodes After Connecting Trees I
// https://leetcode.com/problems/maximize-the-number-of-target-nodes-after-connecting-trees-i/
// Difficulty: Medium
// Time: O(n^2 + m^2) Space: O(n + m)

import "fmt"

func main() {
	fmt.Println(maxTargetNodes([][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}}, [][]int{{0, 1}, {0, 2}, {0, 3}, {2, 4}, {2, 5}}, 3))
	fmt.Println(maxTargetNodes([][]int{{0, 1}, {1, 2}, {1, 3}}, [][]int{{0, 1}, {1, 2}, {1, 3}}, 2))
}

func maxTargetNodes(edges1 [][]int, edges2 [][]int, k int) []int {
	n := len(edges1) + 1
	m := len(edges2) + 1

	g1 := buildGraph3372(edges1, n)
	g2 := buildGraph3372(edges2, m)

	maxFrom2 := 0
	for i := 0; i < m; i++ {
		cnt := countWithinDist3372(g2, i, k-1, m)
		if cnt > maxFrom2 {
			maxFrom2 = cnt
		}
	}

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = countWithinDist3372(g1, i, k, n) + maxFrom2
	}
	return ans
}

func buildGraph3372(edges [][]int, n int) [][]int {
	g := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}
	return g
}

func countWithinDist3372(g [][]int, start int, maxDist int, n int) int {
	if maxDist < 0 {
		return 0
	}
	visited := make([]bool, n)
	queue := []int{start}
	visited[start] = true
	dist := 0
	count := 0

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			count++
			for _, v := range g[queue[i]] {
				if !visited[v] {
					visited[v] = true
					queue = append(queue, v)
				}
			}
		}
		queue = queue[size:]
		dist++
		if dist > maxDist {
			break
		}
	}
	return count
}
```

## 3376 — Minimum Time To Break Locks I

```go
package main

// LeetCode #3376: Minimum Time to Break Locks I
// https://leetcode.com/problems/minimum-time-to-break-locks-i/
// Difficulty: Medium
// Time: O(n! * n) Space: O(n)

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findMinimumTime([]int{3, 4, 1}, 1)) // 4
	fmt.Println(findMinimumTime([]int{2, 5, 4}, 2)) // 5
}

func findMinimumTime(strength []int, K int) int {
	n := len(strength)
	used := make([]bool, n)
	ans := math.MaxInt32

	var dfs func(idx int, time int, X int)
	dfs = func(idx int, time int, X int) {
		if idx == n {
			if time < ans {
				ans = time
			}
			return
		}
		if time >= ans {
			return
		}
		for i := 0; i < n; i++ {
			if used[i] {
				continue
			}
			used[i] = true
			need := (strength[i] + X - 1) / X
			dfs(idx+1, time+need, X+K)
			used[i] = false
		}
	}

	dfs(0, 0, 1)
	return ans
}
```

## 3377 — Digit Operations To Make Two Integers Equal

```go
package main

// LeetCode #3377: Digit Operations to Make Two Integers Equal
// https://leetcode.com/problems/digit-operations-to-make-two-integers-equal/
// Difficulty: Medium
// Time: O(N log N) Space: O(N)

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(minOperations(10, 12)) // 6 (not 10, paths: 10+11+12=33, 10+12=22 if allowed)
}

type Item struct {
	cost int
	num  int
}

type ItemHeap []Item

func (h ItemHeap) Len() int           { return len(h) }
func (h ItemHeap) Less(i, j int) bool { return h[i].cost < h[j].cost }
func (h ItemHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *ItemHeap) Push(x interface{}) { *h = append(*h, x.(Item)) }
func (h *ItemHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

const MAX = 10000

var isPrime []bool

func init() {
	isPrime = make([]bool, MAX+1)
	for i := 2; i <= MAX; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i <= MAX; i++ {
		if isPrime[i] {
			for j := i * i; j <= MAX; j += i {
				isPrime[j] = false
			}
		}
	}
}

func minOperations(n int, m int) int {
	if n == m {
		return n
	}
	if isPrime[n] || isPrime[m] {
		return -1
	}

	dist := make([]int, MAX+1)
	for i := range dist {
		dist[i] = 1 << 60
	}
	dist[n] = n

	h := &ItemHeap{}
	heap.Init(h)
	heap.Push(h, Item{n, n})

	for h.Len() > 0 {
		item := heap.Pop(h).(Item)
		if item.cost > dist[item.num] {
			continue
		}
		if item.num == m {
			return item.cost
		}

		// Generate neighbors by changing each digit
		digits := getDigits(item.num)
		for pos := 0; pos < len(digits); pos++ {
			orig := digits[pos]

			// Increment digit
			if orig < 9 {
				digits[pos] = orig + 1
				next := fromDigits(digits)
				if !isPrime[next] {
					nc := item.cost + next
					if nc < dist[next] {
						dist[next] = nc
						heap.Push(h, Item{nc, next})
					}
				}
			}

			// Decrement digit
			if orig > 0 && !(pos == 0 && orig == 1) {
				// Can't decrement to leading zero
				if pos > 0 || orig > 1 {
					digits[pos] = orig - 1
					next := fromDigits(digits)
					if !isPrime[next] {
						nc := item.cost + next
						if nc < dist[next] {
							dist[next] = nc
							heap.Push(h, Item{nc, next})
						}
					}
				}
			}

			digits[pos] = orig
		}
	}
	return -1
}

func getDigits(num int) []int {
	var d []int
	for num > 0 {
		d = append([]int{num % 10}, d...)
		num /= 10
	}
	return d
}

func fromDigits(d []int) int {
	r := 0
	for _, v := range d {
		r = r*10 + v
	}
	return r
}
```

## 3380 — Maximum Area Rectangle With Point Constraints I

```go
package main

// LeetCode #3380: Maximum Area Rectangle With Point Constraints I
// https://leetcode.com/problems/maximum-area-rectangle-with-point-constraints-i/
// Difficulty: Medium
// Time: O(n^4) Space: O(1)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxRectangleArea([][]int{{1, 1}, {1, 3}, {3, 1}, {3, 3}}))          // 4
	fmt.Println(maxRectangleArea([][]int{{1, 1}, {1, 3}, {3, 1}, {3, 3}, {2, 2}})) // -1
}

func maxRectangleArea(points [][]int) int {
	n := len(points)
	sp := make([][2]int, n)
	for i := 0; i < n; i++ {
		sp[i] = [2]int{points[i][0], points[i][1]}
	}

	ans := -1

	for a := 0; a < n; a++ {
		for b := a + 1; b < n; b++ {
			for c := b + 1; c < n; c++ {
				for d := c + 1; d < n; d++ {
					if area, ok := isRect3380([]int{a, b, c, d}, sp); ok && area > ans {
						ans = area
					}
				}
			}
		}
	}
	return ans
}

func isRect3380(idx []int, pts [][2]int) (int, bool) {
	xSet := make(map[int]bool)
	ySet := make(map[int]bool)
	for _, i := range idx {
		xSet[pts[i][0]] = true
		ySet[pts[i][1]] = true
	}
	if len(xSet) != 2 || len(ySet) != 2 {
		return 0, false
	}

	var xs, ys []int
	for x := range xSet {
		xs = append(xs, x)
	}
	for y := range ySet {
		ys = append(ys, y)
	}
	sort.Ints(xs)
	sort.Ints(ys)
	x1, x2 := xs[0], xs[1]
	y1, y2 := ys[0], ys[1]

	cornerSet := make(map[[2]int]bool)
	for _, i := range idx {
		cornerSet[pts[i]] = true
	}
	expected := [][2]int{{x1, y1}, {x1, y2}, {x2, y1}, {x2, y2}}
	for _, p := range expected {
		if !cornerSet[p] {
			return 0, false
		}
	}

	for _, p := range pts {
		if cornerSet[p] {
			continue
		}
		if p[0] >= x1 && p[0] <= x2 && p[1] >= y1 && p[1] <= y2 {
			return 0, false
		}
	}

	return (x2 - x1) * (y2 - y1), true
}
```

## 3381 — Maximum Subarray Sum With Length Divisible By K

```go
package main

// LeetCode #3381: Maximum Subarray Sum With Length Divisible by K
// https://leetcode.com/problems/maximum-subarray-sum-with-length-divisible-by-k/
// Difficulty: Medium
// Time: O(n) Space: O(k)

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxSubarraySum([]int{1, 2, 3, 4, 5, 6}, 2)) // 21
	fmt.Println(maxSubarraySum([]int{-1, -2, -3, -4, -5}, 3)) // -6
}

func maxSubarraySum(nums []int, k int) int64 {
	n := len(nums)
	prefix := make([]int64, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + int64(nums[i])
	}

	minPref := make([]int64, k)
	for i := range minPref {
		minPref[i] = math.MaxInt64
	}
	minPref[0] = 0 // prefix[0] = 0

	var ans int64 = math.MinInt64

	for i := 1; i <= n; i++ {
		r := i % k
		if minPref[r] != math.MaxInt64 {
			val := prefix[i] - minPref[r]
			if val > ans {
				ans = val
			}
		}
		if prefix[i] < minPref[r] {
			minPref[r] = prefix[i]
		}
	}

	if ans == math.MinInt64 {
		return 0
	}
	return ans
}
```

## 3387 — Maximize Amount After Two Days Of Conversions

```go
package main

// LeetCode #3387: Maximize Amount After Two Days of Conversions
// https://leetcode.com/problems/maximize-amount-after-two-days-of-conversions/
// Difficulty: Medium
// Time: O(n1 + n2) Space: O(currencies)

import "fmt"

func main() {
	fmt.Println(maxAmount("EUR", [][]string{{"EUR", "USD"}}, []float64{2.0}, [][]string{{"USD", "EUR"}}, []float64{0.5}))
	// 1.0 EUR -> 2.0 USD day1 -> 1.0 EUR day2 = 1.0
}

func maxAmount(initialCurrency string, pairs1 [][]string, rates1 []float64, pairs2 [][]string, rates2 []float64) float64 {
	// Day 1: Bellman-Ford to find max amount of each currency
	amounts1 := make(map[string]float64)
	amounts1[initialCurrency] = 1.0

	// Run Bellman-Ford (or just process all pairs repeatedly)
	for i := 0; i < len(pairs1); i++ {
		updated := false
		for j, p := range pairs1 {
			from, to := p[0], p[1]
			r := rates1[j]
			if v, ok := amounts1[from]; ok {
				if v*r > amounts1[to] {
					amounts1[to] = v * r
					updated = true
				}
			}
			if v, ok := amounts1[to]; ok {
				if v/r > amounts1[from] {
					amounts1[from] = v / r
					updated = true
				}
			}
		}
		if !updated {
			break
		}
	}

	// Day 2: start with day1 amounts, find max back to initial
	amounts2 := make(map[string]float64)
	for k, v := range amounts1 {
		amounts2[k] = v
	}

	for i := 0; i < len(pairs2); i++ {
		updated := false
		for j, p := range pairs2 {
			from, to := p[0], p[1]
			r := rates2[j]
			if v, ok := amounts2[from]; ok {
				if v*r > amounts2[to] {
					amounts2[to] = v * r
					updated = true
				}
			}
			if v, ok := amounts2[to]; ok {
				if v/r > amounts2[from] {
					amounts2[from] = v / r
					updated = true
				}
			}
		}
		if !updated {
			break
		}
	}

	return amounts2[initialCurrency]
}
```

## 3388 — Count Beautiful Splits In An Array

```go
package main

// LeetCode #3388: Count Beautiful Splits in an Array
// https://leetcode.com/problems/count-beautiful-splits-in-an-array/
// Difficulty: Medium
// Time: O(n^2) Space: O(n^2)

import "fmt"

func main() {
	fmt.Println(beautifulSplits([]int{1, 1, 2, 1})) // 2
	fmt.Println(beautifulSplits([]int{1, 2, 3, 4})) // 0
}

func beautifulSplits(nums []int) int {
	n := len(nums)

	// lcp[i][j] = longest common prefix of nums[i:] and nums[j:]
	lcp := make([][]int, n+1)
	for i := range lcp {
		lcp[i] = make([]int, n+1)
	}
	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if nums[i] == nums[j] {
				lcp[i][j] = lcp[i+1][j+1] + 1
			}
		}
	}

	ans := 0
	for i := 1; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			// nums1 = [0,i), nums2 = [i,j), nums3 = [j,n)
			ok := false
			// Check if nums1 is prefix of nums2
			if lcp[0][i] >= i {
				ok = true
			}
			// Check if nums2 is prefix of nums3
			if lcp[i][j] >= j-i {
				ok = true
			}
			if ok {
				ans++
			}
		}
	}
	return ans
}
```

## 3391 — Design A 3D Binary Matrix With Efficient Layer Tracking

```go
package main

// LeetCode #3391: Design a 3D Binary Matrix with Efficient Layer Tracking
// https://leetcode.com/problems/design-a-3d-binary-matrix-with-efficient-layer-tracking/
// Difficulty: Medium [Paid]
// Time: O(1) set/get, O(n*m) init  Space: O(n*m*l)

import "fmt"

type Matrix3D struct {
	data [][][]bool
	n    int
	m    int
	l    int
}

func Constructor(n, m, l int) *Matrix3D {
	mat := &Matrix3D{
		data: make([][][]bool, n),
		n:    n,
		m:    m,
		l:    l,
	}
	for i := 0; i < n; i++ {
		mat.data[i] = make([][]bool, m)
		for j := 0; j < m; j++ {
			mat.data[i][j] = make([]bool, l)
		}
	}
	return mat
}

func (mat *Matrix3D) SetCell(x, y, z int) {
	mat.data[x][y][z] = true
}

func (mat *Matrix3D) UnsetCell(x, y, z int) {
	mat.data[x][y][z] = false
}

func (mat *Matrix3D) GetCell(x, y, z int) bool {
	return mat.data[x][y][z]
}

func (mat *Matrix3D) LargestLayer() int {
	maxCount := 0
	maxLayer := 0
	for k := 0; k < mat.l; k++ {
		count := 0
		for i := 0; i < mat.n; i++ {
			for j := 0; j < mat.m; j++ {
				if mat.data[i][j][k] {
					count++
				}
			}
		}
		if count > maxCount {
			maxCount = count
			maxLayer = k
		}
	}
	return maxLayer
}

func main() {
	mat := Constructor(2, 2, 2)
	mat.SetCell(0, 0, 0)
	mat.SetCell(1, 1, 1)
	fmt.Println(mat.GetCell(0, 0, 0)) // true
	fmt.Println(mat.GetCell(1, 1, 0)) // false
	fmt.Println(mat.LargestLayer())   // 1 (both layers have 1 set)
	mat.UnsetCell(1, 1, 1)
	fmt.Println(mat.LargestLayer()) // 0
}
```

## 3393 — Count Paths With The Given Xor Value

```go
package main

// LeetCode #3393: Count Paths With the Given XOR Value
// https://leetcode.com/problems/count-paths-with-the-given-xor-value/
// Difficulty: Medium
// Time: O(m*n*U) Space: O(m*n*U) where U = maxXOR value

import "fmt"

func countPathsWithXorValue(grid [][]int, k int) int {
	const mod = 1_000_000_007
	u := 1
	for _, row := range grid {
		for _, val := range row {
			for u <= val {
				u <<= 1
			}
		}
	}
	if k >= u {
		return 0
	}

	m, n := len(grid), len(grid[0])
	f := make([][][]int, m+1)
	for i := range f {
		f[i] = make([][]int, n+1)
		for j := range f[i] {
			f[i][j] = make([]int, u)
		}
	}

	f[0][1][0] = 1
	for i, row := range grid {
		for j, val := range row {
			for x := 0; x < u; x++ {
				f[i+1][j+1][x] = (f[i+1][j][x^val] + f[i][j+1][x^val]) % mod
			}
		}
	}
	return f[m][n][k]
}

func main() {
	fmt.Println(countPathsWithXorValue([][]int{{2, 1, 5}, {7, 10, 0}, {12, 6, 4}}, 11)) // 3
	fmt.Println(countPathsWithXorValue([][]int{{1, 3, 3, 3}, {0, 3, 3, 2}, {3, 0, 1, 1}}, 2)) // 4
}
```

## 3394 — Check If Grid Can Be Cut Into Sections

```go
package main

// LeetCode #3394: Check if Grid can be Cut into Sections
// https://leetcode.com/problems/check-if-grid-can-be-cut-into-sections/
// Difficulty: Medium
// Time: O(n log n) Space: O(n)

import (
	"fmt"
	"slices"
)

type pair struct{ l, r int }

func check(intervals []pair) bool {
	slices.SortFunc(intervals, func(a, b pair) int { return a.l - b.l })
	cnt, maxR := 0, 0
	for _, p := range intervals {
		if p.l >= maxR {
			cnt++
		}
		if p.r > maxR {
			maxR = p.r
		}
	}
	return cnt >= 3
}

func checkValidCuts(_ int, rectangles [][]int) bool {
	n := len(rectangles)
	a := make([]pair, n)
	b := make([]pair, n)
	for i, rect := range rectangles {
		a[i] = pair{rect[0], rect[2]}
		b[i] = pair{rect[1], rect[3]}
	}
	return check(a) || check(b)
}

func main() {
	fmt.Println(checkValidCuts(5, [][]int{{1, 0, 5, 2}, {0, 2, 2, 4}, {3, 2, 5, 3}, {0, 4, 4, 5}})) // true
	fmt.Println(checkValidCuts(4, [][]int{{0, 0, 1, 1}, {2, 0, 3, 4}, {0, 2, 2, 3}, {3, 0, 4, 3}})) // true
	fmt.Println(checkValidCuts(4, [][]int{{0, 2, 2, 4}, {1, 0, 3, 2}, {2, 2, 3, 4}, {3, 0, 4, 2}, {3, 2, 4, 4}})) // false
}
```

## 3397 — Maximum Number Of Distinct Elements After Operations

```go
package main

// LeetCode #3397: Maximum Number of Distinct Elements After Operations
// https://leetcode.com/problems/maximum-number-of-distinct-elements-after-operations/
// Difficulty: Medium
// Time: O(n log n) Space: O(1)

import (
	"fmt"
	"math"
	"slices"
)

func maxDistinctElements(nums []int, k int) int {
	n := len(nums)
	if k*2+1 >= n {
		return n
	}

	slices.Sort(nums)
	pre := math.MinInt
	ans := 0
	for _, x := range nums {
		candidate := max(x-k, pre+1)
		if candidate <= x+k {
			ans++
			pre = candidate
		}
	}
	return ans
}

func main() {
	fmt.Println(maxDistinctElements([]int{1, 2, 2, 3, 3, 4}, 2)) // 6
	fmt.Println(maxDistinctElements([]int{4, 4, 4, 4}, 1))        // 3
	fmt.Println(maxDistinctElements([]int{1, 1, 1, 1}, 0))        // 1
}
```

## 3400 — Maximum Number Of Matching Indices After Right Shifts

```go
package main

// LeetCode #3400: Maximum Number of Matching Indices After Right Shifts
// https://leetcode.com/problems/maximum-number-of-matching-indices-after-right-shifts/
// Difficulty: Medium [Paid]
// Time: O(n^2) Space: O(1)

import "fmt"

func maximumMatchingIndices(nums []int, x int) int {
	n := len(nums)
	ans := 0
	for shift := 0; shift < n; shift++ {
		cnt := 0
		for i := 0; i < n; i++ {
			j := (i - shift + n) % n
			if nums[i] == nums[j] {
				cnt++
			}
		}
		if cnt > ans {
			ans = cnt
		}
	}
	return ans
}

func main() {
	fmt.Println(maximumMatchingIndices([]int{1, 2, 3, 4}, 0)) // 4
	fmt.Println(maximumMatchingIndices([]int{1, 2, 3, 1}, 0)) // 3
	fmt.Println(maximumMatchingIndices([]int{1, 1, 1, 1}, 0)) // 4
}
```

## 3403 — Find The Lexicographically Largest String From The Box I

```go
package main

// LeetCode #3403: Find the Lexicographically Largest String From the Box I
// https://leetcode.com/problems/find-the-lexicographically-largest-string-from-the-box-i/
// Difficulty: Medium
// Time: O(n^2) Space: O(n)

import "fmt"

func answerString(word string, numFriends int) string {
	if numFriends == 1 {
		return word
	}
	n := len(word)
	maxLen := n - numFriends + 1
	ans := word[:maxLen]
	for i := 0; i < n; i++ {
		end := i + maxLen
		if end > n {
			end = n
		}
		sub := word[i:end]
		if sub > ans {
			ans = sub
		}
	}
	return ans
}

func main() {
	fmt.Println(answerString("dbca", 2)) // "dbc"
	fmt.Println(answerString("gggg", 2)) // "ggg"
	fmt.Println(answerString("abc", 3))  // "c"
}
```

## 3404 — Count Special Subsequences

```go
package main

// LeetCode #3404: Count Special Subsequences
// https://leetcode.com/problems/count-special-subsequences/
// Difficulty: Medium
// Time: O(n^2) Space: O(n)

import "fmt"

func numberOfSubsequences(nums []int) int64 {
	n := len(nums)
	var ans int64
	cnt := make(map[float64]int)

	// For each r, q = r-2. Accumulate (p,q) pairs as r increases.
	for r := 4; r < n-2; r++ {
		q := r - 2
		b := float64(nums[q])
		for _, aVal := range nums[:q-1] {
			ratio := float64(aVal) / b
			cnt[ratio]++
		}

		c := float64(nums[r])
		for _, dVal := range nums[r+2:] {
			ratio := float64(dVal) / c
			ans += int64(cnt[ratio])
		}
	}
	return ans
}

func main() {
	fmt.Println(numberOfSubsequences([]int{1, 2, 3, 4, 3, 6, 1})) // 1
	fmt.Println(numberOfSubsequences([]int{3, 4, 3, 4, 3, 4, 3, 4})) // 3
}
```

## 3408 — Design Task Manager

```go
package main

// LeetCode #3408: Design Task Manager
// https://leetcode.com/problems/design-task-manager/
// Difficulty: Medium
// Time: O(log n) for operations  Space: O(n)

import (
	"container/heap"
	"fmt"
)

type Task struct {
	userId   int
	taskId   int
	priority int
}

type MaxHeap []Task

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].priority > h[j].priority || (h[i].priority == h[j].priority && h[i].taskId > h[j].taskId) }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) { *h = append(*h, x.(Task)) }

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type TaskManager struct {
	tasks map[int]Task // taskId -> Task
	heap  *MaxHeap
}

func Constructor(tasks [][]int) *TaskManager {
	tm := &TaskManager{
		tasks: make(map[int]Task),
		heap:  &MaxHeap{},
	}
	heap.Init(tm.heap)
	for _, t := range tasks {
		task := Task{userId: t[0], taskId: t[1], priority: t[2]}
		tm.tasks[t[1]] = task
		heap.Push(tm.heap, task)
	}
	return tm
}

func (tm *TaskManager) Add(userId, taskId, priority int) {
	task := Task{userId: userId, taskId: taskId, priority: priority}
	tm.tasks[taskId] = task
	heap.Push(tm.heap, task)
}

func (tm *TaskManager) Edit(taskId, newPriority int) {
	if task, ok := tm.tasks[taskId]; ok {
		task.priority = newPriority
		tm.tasks[taskId] = task
		heap.Push(tm.heap, task)
	}
}

func (tm *TaskManager) Rmv(taskId int) {
	delete(tm.tasks, taskId)
}

func (tm *TaskManager) ExecTop() int {
	for tm.heap.Len() > 0 {
		task := heap.Pop(tm.heap).(Task)
		if saved, ok := tm.tasks[task.taskId]; ok && saved.userId == task.userId && saved.priority == task.priority {
			delete(tm.tasks, task.taskId)
			return task.userId
		}
	}
	return -1
}

func main() {
	tm := Constructor([][]int{{1, 101, 10}, {2, 102, 20}, {3, 103, 15}})
	fmt.Println(tm.ExecTop()) // 2
	tm.Add(4, 104, 30)
	fmt.Println(tm.ExecTop()) // 4
	tm.Edit(102, 40)
	fmt.Println(tm.ExecTop()) // 2
	fmt.Println(tm.ExecTop()) // -1
}
```

## 3409 — Longest Subsequence With Decreasing Adjacent Difference

```go
package main

// LeetCode #3409: Longest Subsequence With Decreasing Adjacent Difference
// https://leetcode.com/problems/longest-subsequence-with-decreasing-adjacent-difference/
// Difficulty: Medium
// Time: O(n * maxDiff) Space: O(maxVal * maxDiff)

import (
	"fmt"
	"slices"
)

func longestSubsequence(nums []int) int {
	mx := slices.Max(nums)
	maxD := mx - slices.Min(nums)

	f := make([][]int, mx+1)
	for i := range f {
		f[i] = make([]int, maxD+1)
	}

	ans := 0
	for _, x := range nums {
		fx := 1
		for j := maxD; j >= 0; j-- {
			if x-j >= 0 {
				if f[x-j][j]+1 > fx {
					fx = f[x-j][j] + 1
				}
			}
			if x+j <= mx {
				if f[x+j][j]+1 > fx {
					fx = f[x+j][j] + 1
				}
			}
			f[x][j] = fx
			if fx > ans {
				ans = fx
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(longestSubsequence([]int{16, 6, 3}))              // 3
	fmt.Println(longestSubsequence([]int{6, 5, 3, 4, 2, 1}))      // 4
	fmt.Println(longestSubsequence([]int{10, 20, 30, 40, 50}))    // 5
}
```

## 3412 — Find Mirror Score Of A String

```go
package main

// LeetCode #3412: Find Mirror Score of a String
// https://leetcode.com/problems/find-mirror-score-of-a-string/
// Difficulty: Medium
// Time: O(n) Space: O(n)

import "fmt"

func calculateScore(s string) int64 {
	stacks := make([][]int, 26)
	var ans int64
	for i := 0; i < len(s); i++ {
		ch := int(s[i] - 'a')
		mirror := 25 - ch
		if len(stacks[mirror]) > 0 {
			j := stacks[mirror][len(stacks[mirror])-1]
			stacks[mirror] = stacks[mirror][:len(stacks[mirror])-1]
			ans += int64(i - j)
		} else {
			stacks[ch] = append(stacks[ch], i)
		}
	}
	return ans
}

func main() {
	fmt.Println(calculateScore("aczzx")) // 5
	fmt.Println(calculateScore("abcdef")) // 0
}
```

## 3413 — Maximum Coins From K Consecutive Bags

```go
package main

// LeetCode #3413: Maximum Coins From K Consecutive Bags
// https://leetcode.com/problems/maximum-coins-from-k-consecutive-bags/
// Difficulty: Medium
// Time: O(n log n) Space: O(log n)

import (
	"fmt"
	"slices"
)

func maximumWhiteTiles(tiles [][]int, carpetLen int) int {
	cover, left := 0, 0
	ans := 0
	for _, tile := range tiles {
		tl, tr, c := tile[0], tile[1], tile[2] * (tile[1] - tile[0] + 1)
		_ = c
		cover += (tr - tl + 1) * tile[2]
		for tiles[left][1]+carpetLen-1 < tr {
			cover -= (tiles[left][1] - tiles[left][0] + 1) * tiles[left][2]
			left++
		}
		uncover := max((tr-carpetLen+1-tiles[left][0])*tiles[left][2], 0)
		if cover-uncover > ans {
			ans = cover - uncover
		}
	}
	return ans
}

func maximumCoins(coins [][]int, k int) int64 {
	slices.SortFunc(coins, func(a, b []int) int { return a[0] - b[0] })
	ans := maximumWhiteTiles(coins, k)

	// reverse and negate for right-to-left pass
	slices.Reverse(coins)
	for _, t := range coins {
		t[0], t[1] = -t[1], -t[0]
	}
	ans2 := maximumWhiteTiles(coins, k)
	if ans2 > ans {
		ans = ans2
	}
	return int64(ans)
}

func main() {
	fmt.Println(maximumCoins([][]int{{8, 10, 1}, {1, 3, 2}, {5, 6, 4}}, 4)) // 10
	fmt.Println(maximumCoins([][]int{{1, 4, 2}, {5, 8, 1}}, 3)) // 6
}
```

## 3418 — Maximum Amount Of Money Robot Can Earn

```go
package main

// LeetCode #3418: Maximum Amount of Money Robot Can Earn
// https://leetcode.com/problems/maximum-amount-of-money-robot-can-earn/
// Difficulty: Medium
// Time: O(m*n) Space: O(n)

import (
	"fmt"
	"math"
)

func maximumAmount(coins [][]int) int {
	n := len(coins[0])
	f := make([][3]int, n+1)
	for j := range f {
		f[j] = [3]int{math.MinInt / 2, math.MinInt / 2, math.MinInt / 2}
	}
	f[1] = [3]int{}

	for _, row := range coins {
		for j, x := range row {
			f[j+1][2] = max(f[j][2]+x, f[j+1][2]+x,
				max(f[j][1], f[j+1][1]))
			f[j+1][1] = max(f[j][1]+x, f[j+1][1]+x,
				max(f[j][0], f[j+1][0]))
			f[j+1][0] = max(f[j][0], f[j+1][0]) + x
		}
	}
	return f[n][2]
}

func main() {
	fmt.Println(maximumAmount([][]int{{0, 1, -1}, {1, -2, 3}, {2, -3, 4}})) // 8
	fmt.Println(maximumAmount([][]int{{10, 10, 10}, {10, 10, 10}}))        // 40
}
```

## 3419 — Minimize The Maximum Edge Weight Of Graph

```go
package main

// LeetCode #3419: Minimize the Maximum Edge Weight of Graph
// https://leetcode.com/problems/minimize-the-maximum-edge-weight-of-graph/
// Difficulty: Medium
// Time: O((n+m) log n) Space: O(n+m)

import (
	"container/heap"
	"fmt"
	"math"
	"slices"
)

type edge struct{ to, w int }
type pair struct{ dis, x int }
type hp []pair

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].dis < h[j].dis }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }

func minMaxWeight(n int, edges [][]int, _ int) int {
	if len(edges) < n-1 {
		return -1
	}

	g := make([][]edge, n)
	for _, e := range edges {
		x, y, w := e[0], e[1], e[2]
		g[y] = append(g[y], edge{x, w}) // reverse graph
	}

	dis := make([]int, n)
	for i := range dis {
		dis[i] = math.MaxInt
	}
	dis[0] = 0
	h := &hp{{}}
	for h.Len() > 0 {
		p := heap.Pop(h).(pair)
		x := p.x
		d := p.dis
		if d > dis[x] {
			continue
		}
		for _, e := range g[x] {
			nd := max(d, e.w)
			if nd < dis[e.to] {
				dis[e.to] = nd
				heap.Push(h, pair{nd, e.to})
			}
		}
	}

	ans := slices.Max(dis)
	if ans == math.MaxInt {
		return -1
	}
	return ans
}

func main() {
	fmt.Println(minMaxWeight(6, [][]int{{0, 1, 4}, {0, 2, 3}, {1, 3, 2}, {1, 4, 1}, {2, 5, 5}, {3, 0, 6}, {4, 0, 7}, {5, 0, 8}}, 3)) // 5
	fmt.Println(minMaxWeight(3, [][]int{{0, 1, 2}, {1, 2, 3}, {2, 0, 4}}, 2)) // 4
}
```

## 3421 — Find Students Who Improved

```go
package main

// LeetCode #3421: Find Students Who Improved
// https://leetcode.com/problems/find-students-who-improved/
// Difficulty: Medium (SQL problem — implemented in Go)
// Time: O(n log n) Space: O(n)

import (
	"fmt"
	"sort"
)

type Score struct {
	StudentID int
	Subject   string
	Score     int
	ExamDate  string
}

type ImprovedStudent struct {
	StudentID   int
	Subject     string
	FirstScore  int
	LatestScore int
}

func findStudentsWhoImproved(scores []Score) []ImprovedStudent {
	group := make(map[[2]interface{}][]Score)

	for _, s := range scores {
		key := [2]interface{}{s.StudentID, s.Subject}
		group[key] = append(group[key], s)
	}

	var result []ImprovedStudent
	for key, exams := range group {
		sort.Slice(exams, func(i, j int) bool { return exams[i].ExamDate < exams[j].ExamDate })
		first := exams[0].Score
		last := exams[len(exams)-1].Score
		if last > first && len(exams) >= 2 {
			result = append(result, ImprovedStudent{
				StudentID:   key[0].(int),
				Subject:     key[1].(string),
				FirstScore:  first,
				LatestScore: last,
			})
		}
	}

	sort.Slice(result, func(i, j int) bool {
		if result[i].StudentID != result[j].StudentID {
			return result[i].StudentID < result[j].StudentID
		}
		return result[i].Subject < result[j].Subject
	})
	return result
}

func main() {
	scores := []Score{
		{1, "Math", 70, "2023-01-15"},
		{1, "Math", 85, "2023-02-15"},
		{2, "Science", 60, "2023-01-15"},
		{2, "Science", 55, "2023-02-15"},
		{1, "Science", 80, "2023-01-15"},
		{1, "Science", 90, "2023-02-15"},
	}
	result := findStudentsWhoImproved(scores)
	for _, r := range result {
		fmt.Printf("Student %d improved in %s: %d -> %d\n", r.StudentID, r.Subject, r.FirstScore, r.LatestScore)
	}
}
```

## 3422 — Minimum Operations To Make Subarray Elements Equal

```go
package main

// LeetCode #3422: Minimum Operations to Make Subarray Elements Equal
// https://leetcode.com/problems/minimum-operations-to-make-subarray-elements-equal/
// Difficulty: Medium [Paid]
// Time: O(n log k) Space: O(k)

import (
	"fmt"
	"math"
	"sort"
)

func minOperations(nums []int, k int) int64 {
	n := len(nums)
	if k > n {
		return 0
	}

	var ans int64 = math.MaxInt64
	window := make([]int, k)

	for i := 0; i <= n-k; i++ {
		copy(window, nums[i:i+k])
		sort.Ints(window)
		median := window[k/2]

		var ops int64
		for _, v := range window {
			diff := v - median
			if diff < 0 {
				diff = -diff
			}
			ops += int64(diff)
		}
		if ops < ans {
			ans = ops
		}
	}
	return ans
}

func main() {
	fmt.Println(minOperations([]int{1, 4, 2, 6}, 3)) // 3
	fmt.Println(minOperations([]int{1, 2, 3, 4}, 2)) // 0
}
```

## 3424 — Minimum Cost To Make Arrays Identical

```go
package main

// LeetCode #3424: Minimum Cost to Make Arrays Identical
// https://leetcode.com/problems/minimum-cost-to-make-arrays-identical/
// Difficulty: Medium
// Time: O(n log n) Space: O(n)

import (
	"fmt"
	"sort"
)

func minCost(arr []int, brr []int, k int64) int64 {
	var cost1 int64
	for i := 0; i < len(arr); i++ {
		diff := arr[i] - brr[i]
		if diff < 0 {
			diff = -diff
		}
		cost1 += int64(diff)
	}

	sortedArr := make([]int, len(arr))
	sortedBrr := make([]int, len(brr))
	copy(sortedArr, arr)
	copy(sortedBrr, brr)
	sort.Ints(sortedArr)
	sort.Ints(sortedBrr)

	var cost2 int64
	for i := 0; i < len(sortedArr); i++ {
		diff := sortedArr[i] - sortedBrr[i]
		if diff < 0 {
			diff = -diff
		}
		cost2 += int64(diff)
	}
	cost2 += k

	if cost1 < cost2 {
		return cost1
	}
	return cost2
}

func main() {
	fmt.Println(minCost([]int{4, 2, 5}, []int{6, 3, 1}, 2)) // 7
	fmt.Println(minCost([]int{1, 2, 3}, []int{4, 5, 6}, 1)) // 9
}
```

## 3428 — Maximum And Minimum Sums Of At Most Size K Subsequences

```go
package main

// LeetCode #3428: Maximum and Minimum Sums of at Most Size K Subsequences
// https://leetcode.com/problems/maximum-and-minimum-sums-of-at-most-size-k-subsequences/
// Difficulty: Medium
// Time: O(n log n) Space: O(1)

import (
	"fmt"
	"slices"
)

const mod3428 = 1_000_000_007
const mx = 100000

var fac [mx]int
var invFac [mx]int

func init() {
	fac[0] = 1
	for i := 1; i < mx; i++ {
		fac[i] = fac[i-1] * i % mod3428
	}
	invFac[mx-1] = pow3428(fac[mx-1], mod3428-2)
	for i := mx - 1; i > 0; i-- {
		invFac[i-1] = invFac[i] * i % mod3428
	}
}

func pow3428(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod3428
		}
		x = x * x % mod3428
	}
	return res
}

func comb(n, k int) int {
	if k > n || k < 0 {
		return 0
	}
	return fac[n] * invFac[k] % mod3428 * invFac[n-k] % mod3428
}

func minMaxSums(nums []int, k int) int {
	slices.Sort(nums)
	ans := 0
	s := 1
	n := len(nums)
	for i, x := range nums {
		ans = (ans + s*(x+nums[n-1-i])) % mod3428
		s = (s*2 - comb(i, k-1) + mod3428) % mod3428
	}
	return ans
}

func main() {
	fmt.Println(minMaxSums([]int{5, 0, 6}, 1)) // 22
	fmt.Println(minMaxSums([]int{1, 2, 3}, 2)) // 24
}
```

## 3429 — Paint House Iv

```go
package main

// LeetCode #3429: Paint House IV
// https://leetcode.com/problems/paint-house-iv/
// Difficulty: Medium
// Time: O(n) Space: O(1)

import (
	"fmt"
	"math"
)

func minCost3429(n int, cost [][]int) int64 {
	half := n / 2
	dp := make([][][]int64, half+1)
	for i := range dp {
		dp[i] = make([][]int64, 3)
		for j := 0; j < 3; j++ {
			dp[i][j] = make([]int64, 3)
			for k := 0; k < 3; k++ {
				dp[0][j][k] = 0
			}
		}
	}

	for i := 0; i < half; i++ {
		left := i
		right := n - 1 - i
		for preJ := 0; preJ < 3; preJ++ {
			for preK := 0; preK < 3; preK++ {
				best := int64(math.MaxInt64)
				for j := 0; j < 3; j++ {
					if j == preJ {
						continue
					}
					for k := 0; k < 3; k++ {
						if k == preK || k == j {
							continue
						}
						val := dp[i][j][k] + int64(cost[left][j]) + int64(cost[right][k])
						if val < best {
							best = val
						}
					}
				}
				dp[i+1][preJ][preK] = best
			}
		}
	}

	ans := int64(math.MaxInt64)
	for j := 0; j < 3; j++ {
		for k := 0; k < 3; k++ {
			if dp[half][j][k] < ans {
				ans = dp[half][j][k]
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(minCost3429(4, [][]int{{3, 5, 7}, {6, 2, 9}, {4, 8, 1}, {7, 3, 5}})) // 9
	fmt.Println(minCost3429(2, [][]int{{1, 2, 3}, {4, 5, 6}})) // 7
}
```

## 3431 — Minimum Unlocked Indices To Sort Nums

```go
package main

// LeetCode #3431: Minimum Unlocked Indices to Sort Nums
// https://leetcode.com/problems/minimum-unlocked-indices-to-sort-nums/
// Difficulty: Medium [Paid]
// Time: O(n^2) Space: O(n)

import (
	"fmt"
	"sort"
)

func minUnlockedIndices(nums []int, locked []int) int {
	n := len(nums)
	sorted := make([]int, n)
	copy(sorted, nums)
	sort.Ints(sorted)

	unlockCost := 0
	for i := 0; i < n; i++ {
		if nums[i] != sorted[i] && locked[i] == 1 {
			unlockCost++
		}
	}
	return unlockCost
}

func main() {
	fmt.Println(minUnlockedIndices([]int{1, 3, 2, 4}, []int{1, 1, 0, 1})) // 1
	fmt.Println(minUnlockedIndices([]int{1, 2, 3}, []int{1, 1, 1}))       // 0
}
```

## 3433 — Count Mentions Per User

```go
package main

// LeetCode #3433: Count Mentions Per User
// https://leetcode.com/problems/count-mentions-per-user/
// Difficulty: Medium
// Time: O(n log n + n*u) Space: O(n + u)

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func countMentions(numberOfUsers int, events [][]string) []int {
	sort.Slice(events, func(i, j int) bool {
		x, _ := strconv.Atoi(events[i][1])
		y, _ := strconv.Atoi(events[j][1])
		if x == y {
			return events[i][0][2] < events[j][0][2]
		}
		return x < y
	})

	ans := make([]int, numberOfUsers)
	onlineT := make([]int, numberOfUsers)
	lazy := 0

	for _, e := range events {
		etype := e[0]
		cur, _ := strconv.Atoi(e[1])
		s := e[2]

		if etype[0] == 'O' {
			userID, _ := strconv.Atoi(s)
			onlineT[userID] = cur + 60
		} else if s[0] == 'A' {
			lazy++
		} else if s[0] == 'H' {
			for i := 0; i < numberOfUsers; i++ {
				if onlineT[i] <= cur {
					ans[i]++
				}
			}
		} else {
			mentions := strings.Split(s, " ")
			for _, m := range mentions {
				userID, _ := strconv.Atoi(m[2:])
				ans[userID]++
			}
		}
	}

	if lazy > 0 {
		for i := 0; i < numberOfUsers; i++ {
			ans[i] += lazy
		}
	}
	return ans
}

func main() {
	fmt.Println(countMentions(3, [][]string{{"MESSAGE", "1", "ALL"}, {"OFFLINE", "2", "1"}, {"MESSAGE", "3", "HERE"}})) // [2, 1, 2]
	fmt.Println(countMentions(2, [][]string{{"MESSAGE", "0", "id0"}, {"MESSAGE", "1", "id1"}})) // [1, 1]
}
```

## 3434 — Maximum Frequency After Subarray Operation

```go
package main

// LeetCode #3434: Maximum Frequency After Subarray Operation
// https://leetcode.com/problems/maximum-frequency-after-subarray-operation/
// Difficulty: Medium
// Time: O(50*n) Space: O(1)

import "fmt"

func maxFrequency(nums []int, k int) int {
	cntK := 0
	for _, v := range nums {
		if v == k {
			cntK++
		}
	}
	ans := cntK
	for target := 1; target <= 50; target++ {
		if target == k {
			continue
		}
		cur := 0
		for _, x := range nums {
			if x == target {
				cur++
			} else if x == k {
				cur--
			}
			if cur < 0 {
				cur = 0
			}
			if cntK+cur > ans {
				ans = cntK + cur
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(maxFrequency([]int{10, 2, 3, 4, 5, 5, 4, 3, 2, 2}, 10)) // 4
	fmt.Println(maxFrequency([]int{1, 2, 3, 4, 5}, 1)) // 2
}
```

## 3437 — Permutations Iii

```go
package main

// LeetCode #3437: Permutations III
// https://leetcode.com/problems/permutations-iii/
// Difficulty: Medium [Paid]
// Time: O(n!) Space: O(n)

import "fmt"

func permute(n int) [][]int {
	var ans [][]int
	used := make([]bool, n+1)
	cur := make([]int, 0, n)

	var dfs func()
	dfs = func() {
		if len(cur) == n {
			tmp := make([]int, n)
			copy(tmp, cur)
			ans = append(ans, tmp)
			return
		}
		start := len(cur)%2 + 1
		for i := start; i <= n; i += 2 {
			if !used[i] {
				used[i] = true
				cur = append(cur, i)
				dfs()
				cur = cur[:len(cur)-1]
				used[i] = false
			}
		}
	}
	dfs()
	return ans
}

func main() {
	fmt.Println(len(permute(3))) // 2
	fmt.Println(len(permute(4))) // 4
	for _, p := range permute(3) {
		fmt.Println(p)
	}
}
```

## 3439 — Reschedule Meetings For Maximum Free Time I

```go
package main

// LeetCode #3439: Reschedule Meetings for Maximum Free Time I
// https://leetcode.com/problems/reschedule-meetings-for-maximum-free-time-i/
// Difficulty: Medium
// Time: O(n) Space: O(n)

import "fmt"

func maxFreeTime(eventTime int, k int, startTime []int, endTime []int) int {
	n := len(startTime)
	gaps := make([]int, 0, n+1)
	gaps = append(gaps, startTime[0])
	for i := 1; i < n; i++ {
		gaps = append(gaps, startTime[i]-endTime[i-1])
	}
	gaps = append(gaps, eventTime-endTime[n-1])

	window := 0
	for i := 0; i < k+1 && i < len(gaps); i++ {
		window += gaps[i]
	}
	ans := window
	for i := k + 1; i < len(gaps); i++ {
		window += gaps[i] - gaps[i-(k+1)]
		if window > ans {
			ans = window
		}
	}
	return ans
}

func main() {
	fmt.Println(maxFreeTime(10, 1, []int{0, 3, 7, 9}, []int{1, 4, 8, 10})) // 3
	fmt.Println(maxFreeTime(5, 2, []int{1, 3}, []int{2, 4})) // 2
}
```

## 3440 — Reschedule Meetings For Maximum Free Time Ii

```go
package main

// LeetCode #3440: Reschedule Meetings for Maximum Free Time II
// https://leetcode.com/problems/reschedule-meetings-for-maximum-free-time-ii/
// Difficulty: Medium
// Time: O(n) Space: O(n)

import "fmt"

func maxFreeTime3440(eventTime int, startTime []int, endTime []int) int {
	n := len(startTime)
	v := make([]int, 0, n+1)
	last := 0
	for i := 0; i < n; i++ {
		v = append(v, startTime[i]-last)
		last = endTime[i]
	}
	if last <= eventTime {
		v = append(v, eventTime-last)
	}

	m := len(v)
	q := make([]int, m)
	q[m-1] = 0
	for i := m - 2; i >= 0; i-- {
		if v[i+1] > q[i+1] {
			q[i] = v[i+1]
		} else {
			q[i] = q[i+1]
		}
	}

	mx := 0
	ans := 0
	for i := 1; i < m; i++ {
		length := endTime[i-1] - startTime[i-1]
		t := v[i] + v[i-1]
		if length <= mx || length <= q[i] {
			if t+length > ans {
				ans = t + length
			}
		} else {
			if t > ans {
				ans = t
			}
		}
		if v[i-1] > mx {
			mx = v[i-1]
		}
	}
	return ans
}

func main() {
	fmt.Println(maxFreeTime3440(10, []int{0, 3, 7, 9}, []int{1, 4, 8, 10})) // 6
	fmt.Println(maxFreeTime3440(5, []int{1, 3}, []int{2, 4})) // 2
}
```

## 3443 — Maximum Manhattan Distance After K Changes

```go
package main

// LeetCode #3443: Maximum Manhattan Distance After K Changes
// https://leetcode.com/problems/maximum-manhattan-distance-after-k-changes/
// Difficulty: Medium
// Time: O(n) Space: O(1)

import "fmt"

func maxDistance(s string, k int) int {
	x, y := 0, 0
	ans := 0
	for i, ch := range s {
		switch ch {
		case 'N':
			y++
		case 'S':
			y--
		case 'E':
			x++
		case 'W':
			x--
		}
		dist := abs(x) + abs(y) + 2*k
		if i+1 < dist {
			dist = i + 1
		}
		if dist > ans {
			ans = dist
		}
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	fmt.Println(maxDistance("NWSE", 1))   // 3
	fmt.Println(maxDistance("NSWWEW", 3)) // 6
}
```

## 3446 — Sort Matrix By Diagonals

```go
package main

// LeetCode #3446: Sort Matrix by Diagonals
// https://leetcode.com/problems/sort-matrix-by-diagonals/
// Difficulty: Medium
// Time: O(n^2 log n) Space: O(n)

import (
	"fmt"
	"slices"
)

func sortMatrix(grid [][]int) [][]int {
	n := len(grid)
	if n <= 1 {
		return grid
	}

	// Bottom-left (including main diagonal) — descending
	for i := 0; i < n; i++ {
		r, c := i, 0
		length := n - i
		diag := make([]int, length)
		for j := 0; j < length; j++ {
			diag[j] = grid[r+j][c+j]
		}
		slices.SortFunc(diag, func(a, b int) int { return b - a })
		for j := 0; j < length; j++ {
			grid[r+j][c+j] = diag[j]
		}
	}

	// Top-right (excluding main diagonal) — ascending
	for j := 1; j < n; j++ {
		r, c := 0, j
		length := n - j
		diag := make([]int, length)
		for k := 0; k < length; k++ {
			diag[k] = grid[r+k][c+k]
		}
		slices.Sort(diag)
		for k := 0; k < length; k++ {
			grid[r+k][c+k] = diag[k]
		}
	}
	return grid
}

func main() {
	fmt.Println(sortMatrix([][]int{{1, 7, 3}, {9, 8, 2}, {4, 5, 6}}))
	// [[8 2 3] [9 6 7] [4 5 1]]
	fmt.Println(sortMatrix([][]int{{0, 1}, {2, 3}}))
	// [[2 1] [2 3]]
}
```

## 3447 — Assign Elements To Groups With Constraints

```go
package main

// LeetCode #3447: Assign Elements to Groups with Constraints
// https://leetcode.com/problems/assign-elements-to-groups-with-constraints/
// Difficulty: Medium
// Time: O(mx log mx + n + m) Space: O(mx)

import (
	"fmt"
	"slices"
)

func assignElements(groups []int, elements []int) []int {
	mx := slices.Max(groups)
	target := make([]int, mx+1)
	for i := range target {
		target[i] = -1
	}

	for i, x := range elements {
		if x > mx || target[x] >= 0 {
			continue
		}
		for y := x; y <= mx; y += x {
			if target[y] < 0 {
				target[y] = i
			}
		}
	}

	for i, x := range groups {
		groups[i] = target[x]
	}
	return groups
}

func main() {
	fmt.Println(assignElements([]int{8, 4, 3, 2, 4}, []int{4, 2})) // [0 0 -1 1 0]
	fmt.Println(assignElements([]int{10, 5, 7}, []int{2, 5})) // [1 1 -1]
}
```

## 3453 — Separate Squares I

```go
package main

// LeetCode #3453: Separate Squares I
// https://leetcode.com/problems/separate-squares-i/
// Difficulty: Medium
// Time: O(n log n) Space: O(n)

import (
	"fmt"
	"math"
	"sort"
)

func separateSquares(squares [][]int) float64 {
	type event struct {
		y   float64
		dy  float64
		x1  float64
		x2  float64
	}
	var events []event
	var totalArea float64

	for _, sq := range squares {
		x, y, l := float64(sq[0]), float64(sq[1]), float64(sq[2])
		events = append(events, event{y, l, x, x + l})
		totalArea += l * l
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].y < events[j].y
	})

	halfArea := totalArea / 2.0

	lo, hi := float64(squares[0][1]), float64(squares[0][1])
	for _, sq := range squares {
		y := float64(sq[1])
		ye := y + float64(sq[2])
		if y < lo {
			lo = y
		}
		if ye > hi {
			hi = ye
		}
	}

	areaBelow := func(y float64) float64 {
		var area float64
		for _, sq := range squares {
			sy, l := float64(sq[1]), float64(sq[2])
			sye := sy + l
			if y <= sy {
				continue
			}
			overlap := math.Min(y, sye) - sy
			if overlap > 0 {
				area += overlap * l
			}
		}
		return area
	}

	for i := 0; i < 60; i++ {
		mid := (lo + hi) / 2
		if areaBelow(mid) < halfArea {
			lo = mid
		} else {
			hi = mid
		}
	}
	return (lo + hi) / 2
}

func main() {
	fmt.Printf("%.5f\n", separateSquares([][]int{{0, 0, 2}, {1, 1, 1}})) // 1.00000
	fmt.Printf("%.5f\n", separateSquares([][]int{{0, 0, 1}, {2, 2, 1}})) // 1.00000
}
```

## 3457 — Eat Pizzas

```go
package main

// LeetCode #3457: Eat Pizzas!
// https://leetcode.com/problems/eat-pizzas/
// Difficulty: Medium
// Time: O(n log n) Space: O(log n)

import (
	"fmt"
	"sort"
)

func maxWeight(pizzas []int) int64 {
	sort.Slice(pizzas, func(i, j int) bool {
		return pizzas[i] > pizzas[j]
	})
	days := len(pizzas) / 4
	odd := (days + 1) / 2
	ans := int64(0)
	for i := 0; i < odd; i++ {
		ans += int64(pizzas[i])
	}
	for i := odd + 1; i < odd+days/2*2; i += 2 {
		ans += int64(pizzas[i])
	}
	return ans
}

func main() {
	fmt.Println(maxWeight([]int{1, 2, 3, 4, 5, 6, 7, 8})) // 14
	fmt.Println(maxWeight([]int{2, 2, 2, 2, 2, 2, 2, 2})) // 8
}
```

## 3458 — Select K Disjoint Special Substrings

```go
package main

// LeetCode #3458: Select K Disjoint Special Substrings
// https://leetcode.com/problems/select-k-disjoint-special-substrings/
// Difficulty: Medium
// Time: O(n + 26^3) Space: O(26)

import (
	"fmt"
	"slices"
)

func maxSubstringLength(s string, k int) bool {
	if k == 0 {
		return true
	}

	pos := [26][]int{}
	for i, b := range s {
		b -= 'a'
		pos[b] = append(pos[b], i)
	}

	g := [26][]int{}
	for i, p := range pos {
		if p == nil {
			continue
		}
		l, r := p[0], p[len(p)-1]
		for j, q := range pos {
			if j == i || q == nil {
				continue
			}
			idx := lowerBound(q, l)
			if idx < len(q) && q[idx] <= r {
				g[i] = append(g[i], j)
			}
		}
	}

	visited := make([]bool, 26)
	var intervals [][2]int

	for i, p := range pos {
		if p == nil {
			continue
		}
		for j := range visited {
			visited[j] = false
		}
		curL, curR := len(s), 0

		var dfs func(x int)
		dfs = func(x int) {
			visited[x] = true
			pp := pos[x]
			if pp[0] < curL {
				curL = pp[0]
			}
			if pp[len(pp)-1] > curR {
				curR = pp[len(pp)-1]
			}
			for _, y := range g[x] {
				if !visited[y] {
					dfs(y)
				}
			}
		}
		dfs(i)

		if curL > 0 || curR < len(s)-1 {
			intervals = append(intervals, [2]int{curL, curR})
		}
	}

	slices.SortFunc(intervals, func(a, b [2]int) int { return a[1] - b[1] })
	ans := 0
	preR := -1
	for _, p := range intervals {
		if p[0] > preR {
			ans++
			preR = p[1]
		}
	}
	return ans >= k
}

func lowerBound(arr []int, target int) int {
	l, r := 0, len(arr)
	for l < r {
		mid := (l + r) / 2
		if arr[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

func main() {
	fmt.Println(maxSubstringLength("abcdbaefab", 2)) // true
	fmt.Println(maxSubstringLength("cbc", 1)) // true
}
```

## 3460 — Longest Common Prefix After At Most One Removal

```go
package main

// LeetCode #3460: Longest Common Prefix After at Most One Removal
// https://leetcode.com/problems/longest-common-prefix-after-at-most-one-removal/
// Difficulty: Medium [Paid]
// Time: O(min(n,m)) Space: O(1)

import "fmt"

func longestCommonPrefix(s string, t string) int {
	n, m := len(s), len(t)
	maxLen := 0

	// without removal
	i, j := 0, 0
	for i < n && j < m && s[i] == t[j] {
		i++
		j++
	}
	maxLen = i

	// with one removal from s
	i, j = 0, 0
	removed := false
	for i < n && j < m {
		if s[i] == t[j] {
			i++
			j++
		} else if !removed {
			i++
			removed = true
		} else {
			break
		}
	}
	if j > maxLen {
		maxLen = j
	}

	return maxLen
}

func main() {
	fmt.Println(longestCommonPrefix("abcde", "abfde")) // 2 (ab)
	fmt.Println(longestCommonPrefix("abc", "abc"))     // 3
	fmt.Println(longestCommonPrefix("abcd", "abxd"))   // 3 (abx vs abc — remove c from s -> abxd vs abxd... actually ab)
}
```

## 3462 — Maximum Sum With At Most K Elements

```go
package main

// LeetCode #3462: Maximum Sum With at Most K Elements
// https://leetcode.com/problems/maximum-sum-with-at-most-k-elements/
// Difficulty: Medium
// Time: O(n*m*log(m) + total*log(k)) Space: O(total)

import (
	"fmt"
	"sort"
)

func maxSum(grid [][]int, limits []int, k int) int64 {
	var candidates []int
	for i, row := range grid {
		sort.Ints(row)
		lim := limits[i]
		m := len(row)
		for j := m - lim; j < m; j++ {
			candidates = append(candidates, row[j])
		}
	}
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] > candidates[j]
	})
	var sum int64
	for i := 0; i < k && i < len(candidates); i++ {
		sum += int64(candidates[i])
	}
	return sum
}

func main() {
	fmt.Println(maxSum([][]int{{5, 3, 7}, {8, 2, 6}}, []int{2, 2}, 3)) // 21
	fmt.Println(maxSum([][]int{{1, 2}, {3, 4}}, []int{1, 1}, 2)) // 7
}
```

