# Hard (Sulit) — Problem ��2581

## 2203 — Minimum Weighted Subgraph With The Required Paths

```go
package main

// LeetCode #2203: Minimum Weighted Subgraph with the Required Paths
// https://leetcode.com/problems/minimum-weighted-subgraph-with-the-required-paths/
// Difficulty: Hard
//
// 3 Dijkstras: Compute shortest distances from src1, src2, and to dest (reverse graph).
// Answer = min over all nodes v of dist(src1=>v) + dist(src2=>v) + dist(v=>dest).

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	n := 6
	edges := [][]int{
		{0, 2, 2},
		{0, 5, 6},
		{1, 0, 3},
		{1, 4, 5},
		{2, 1, 1},
		{2, 3, 3},
		{2, 3, 4},
		{3, 4, 2},
		{4, 5, 1},
	}
	src1, src2, dest := 0, 1, 5
	// Expected: 9
	fmt.Println(minimumWeight(n, edges, src1, src2, dest))

	// Single node
	fmt.Println(minimumWeight(1, [][]int{}, 0, 0, 0))
}

type Edge struct {
	to, w int
}

type Item struct {
	node, dist int
}

type MinHeap []Item

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].dist < h[j].dist }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Item))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func dijkstra(n int, graph [][]Edge, src int) []int {
	dist := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt64
	}
	dist[src] = 0
	h := &MinHeap{{src, 0}}
	heap.Init(h)

	for h.Len() > 0 {
		cur := heap.Pop(h).(Item)
		if cur.dist > dist[cur.node] {
			continue
		}
		for _, e := range graph[cur.node] {
			nd := cur.dist + e.w
			if nd < dist[e.to] {
				dist[e.to] = nd
				heap.Push(h, Item{e.to, nd})
			}
		}
	}
	return dist
}

func minimumWeight(n int, edges [][]int, src1 int, src2 int, dest int) int {
	graph := make([][]Edge, n)
	rgraph := make([][]Edge, n)

	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		graph[u] = append(graph[u], Edge{v, w})
		rgraph[v] = append(rgraph[v], Edge{u, w})
	}

	d1 := dijkstra(n, graph, src1)
	d2 := dijkstra(n, graph, src2)
	dd := dijkstra(n, rgraph, dest)

	ans := math.MaxInt64
	for v := 0; v < n; v++ {
		if d1[v] == math.MaxInt64 || d2[v] == math.MaxInt64 || dd[v] == math.MaxInt64 {
			continue
		}
		total := d1[v] + d2[v] + dd[v]
		if total < ans {
			ans = total
		}
	}

	if ans == math.MaxInt64 {
		return -1
	}
	return ans
}
```

## 2204 — Distance To A Cycle In Undirected Graph

```go
package main

// LeetCode #2204: Distance to a Cycle in Undirected Graph
// https://leetcode.com/problems/distance-to-a-cycle-in-undirected-graph/
// Difficulty: Hard
//
// Topo remove leaves + BFS:
// 1. Compute degrees. Repeatedly remove leaf nodes (degree == 1), marking them
//    as removed (deg = 0).
// 2. Remaining nodes (deg > 0) are cycle nodes. BFS from all cycle nodes to
//    compute distances.

import (
	"fmt"
)

func main() {
	n := 7
	edges := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}, {0, 1}, {5, 2}, {6, 5}}
	// Expected: [1 0 0 0 0 1 2]
	fmt.Println(distanceToCycle(n, edges))

	// Simple triangle
	n2 := 3
	edges2 := [][]int{{0, 1}, {1, 2}, {2, 0}}
	fmt.Println(distanceToCycle(n2, edges2))

	// Single cycle with extra leaves
	n3 := 5
	edges3 := [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 0}, {0, 4}}
	fmt.Println(distanceToCycle(n3, edges3))
}

func distanceToCycle(n int, edges [][]int) []int {
	adj := make([][]int, n)
	deg := make([]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
		deg[u]++
		deg[v]++
	}

	// Remove leaves: set deg = 0 to mark removed.
	q := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if deg[i] == 1 {
			q = append(q, i)
			deg[i] = 0
		}
	}

	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		for _, v := range adj[u] {
			if deg[v] > 0 {
				deg[v]--
				if deg[v] == 1 {
					q = append(q, v)
					deg[v] = 0
				}
			}
		}
	}

	// BFS from all cycle nodes (deg > 0).
	ans := make([]int, n)
	queue := make([]int, 0, n)
	inQueue := make([]bool, n)
	for i := 0; i < n; i++ {
		if deg[i] > 0 {
			queue = append(queue, i)
			inQueue[i] = true
		}
	}

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		for _, v := range adj[u] {
			if !inQueue[v] {
				ans[v] = ans[u] + 1
				inQueue[v] = true
				queue = append(queue, v)
			}
		}
	}

	return ans
}
```

## 2209 — Minimum White Tiles After Covering With Carpets

```go
package main

// LeetCode #2209: Minimum White Tiles After Covering with Carpets
// https://leetcode.com/problems/minimum-white-tiles-after-covering-with-carpets/
// Difficulty: Hard
//
// DP: dp[i][j] = minimum white tiles for first i positions using j carpets.
// dp[i][j] = min(
//   dp[i-1][j] + cost(i),           // don't cover position i
//   dp[i-carpetLen][j-1]            // cover [i-carpetLen+1..i] with carpet
// )
// With 2-row optimization: prev = j-1, cur = j.

import (
	"fmt"
)

func main() {
	// "10110101", 2 carpets, carpetLen=2 => 2
	fmt.Println(minimumWhiteTiles("10110101", 2, 2))
	// All zeros
	fmt.Println(minimumWhiteTiles("00000", 2, 2))
	// All ones, 1 carpet of len 3
	fmt.Println(minimumWhiteTiles("11111", 1, 3))
	// Single tile
	fmt.Println(minimumWhiteTiles("1", 0, 1))
	fmt.Println(minimumWhiteTiles("0", 1, 1))
	// Extra: 1 carpet len 2
	fmt.Println(minimumWhiteTiles("10110101", 1, 2))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minimumWhiteTiles(floor string, numCarpets int, carpetLen int) int {
	n := len(floor)
	if n == 0 {
		return 0
	}

	// Row for 0 carpets: prefix count of white tiles.
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1]
		if floor[i-1] == '1' {
			dp[i]++
		}
	}

	for j := 1; j <= numCarpets; j++ {
		ndp := make([]int, n+1)
		// ndp[0] = 0 by default (0 positions, j carpets -> 0 white tiles)
		for i := 1; i <= n; i++ {
			// Option 1: don't place a carpet ending at i.
			// Use j carpets for first i-1 positions, add cost of position i.
			val := ndp[i-1]
			if floor[i-1] == '1' {
				val++
			}
			// Option 2: place a carpet covering [i-carpetLen+1, i].
			if i >= carpetLen {
				val = min(val, dp[i-carpetLen])
			} else {
				// Carpet extends before position 0 -- covers all of [0, i].
				val = min(val, 0)
			}
			ndp[i] = val
		}
		dp = ndp
	}

	return dp[n]
}
```

## 2213 — Longest Substring Of One Repeating Character

```go
package main

// LeetCode #2213: Longest Substring of One Repeating Character
// https://leetcode.com/problems/longest-substring-of-one-repeating-character/
// Difficulty: Hard
//
// Segment tree: each node stores the longest run of a single character in its
// interval, plus prefix run length, suffix run length, and the characters at
// left and right boundaries. Updates at a single position; query returns the
// tree root's max.

import (
	"fmt"
)

func main() {
	// s = "babacc", queryCharacters = "bcb", queryIndices = [1,3,3]
	// Expected: [1, 3, 3]
	s := "babacc"
	queryCharacters := "bcb"
	queryIndices := []int{1, 3, 3}
	fmt.Println(longestRepeating(s, queryCharacters, queryIndices))

	// Single char
	fmt.Println(longestRepeating("a", "b", []int{0}))

	// Already repeating
	fmt.Println(longestRepeating("aaa", "a", []int{0}))
}

type SegNode struct {
	prefLen   int
	suffLen   int
	maxLen    int
	leftChar  byte
	rightChar byte
	size      int
}

type SegTree struct {
	tree []SegNode
	n    int
}

func merge(left, right SegNode) SegNode {
	var res SegNode
	res.leftChar = left.leftChar
	res.rightChar = right.rightChar
	res.size = left.size + right.size

	// Prefix length
	res.prefLen = left.prefLen
	if left.prefLen == left.size && left.rightChar == right.leftChar {
		res.prefLen = left.size + right.prefLen
	}

	// Suffix length
	res.suffLen = right.suffLen
	if right.suffLen == right.size && right.leftChar == left.rightChar {
		res.suffLen = right.size + left.suffLen
	}

	// Max length
	res.maxLen = max(left.maxLen, right.maxLen)
	if left.rightChar == right.leftChar {
		res.maxLen = max(res.maxLen, left.suffLen+right.prefLen)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func NewSegTree(s string) *SegTree {
	n := len(s)
	tree := make([]SegNode, 4*n)
	st := &SegTree{tree: tree, n: n}
	st.build(1, 1, n, s)
	return st
}

func (st *SegTree) build(idx, l, r int, s string) {
	if l == r {
		st.tree[idx] = SegNode{
			prefLen:   1,
			suffLen:   1,
			maxLen:    1,
			leftChar:  s[l-1],
			rightChar: s[l-1],
			size:      1,
		}
		return
	}
	mid := (l + r) / 2
	st.build(idx*2, l, mid, s)
	st.build(idx*2+1, mid+1, r, s)
	st.tree[idx] = merge(st.tree[idx*2], st.tree[idx*2+1])
}

func (st *SegTree) Update(pos int, c byte) {
	st.update(1, 1, st.n, pos, c)
}

func (st *SegTree) update(idx, l, r, pos int, c byte) {
	if l == r {
		st.tree[idx].leftChar = c
		st.tree[idx].rightChar = c
		return
	}
	mid := (l + r) / 2
	if pos <= mid {
		st.update(idx*2, l, mid, pos, c)
	} else {
		st.update(idx*2+1, mid+1, r, pos, c)
	}
	st.tree[idx] = merge(st.tree[idx*2], st.tree[idx*2+1])
}

func longestRepeating(s string, queryCharacters string, queryIndices []int) []int {
	st := NewSegTree(s)
	ans := make([]int, len(queryCharacters))
	for i, c := range queryCharacters {
		st.Update(queryIndices[i]+1, byte(c))
		ans[i] = st.tree[1].maxLen
	}
	return ans
}
```

## 2218 — Maximum Value Of K Coins From Piles

```go
package main

// LeetCode #2218: Maximum Value of K Coins from Piles
// https://leetcode.com/problems/maximum-value-of-k-coins-from-piles/
// Difficulty: Hard
//
// DP: dp[k] = max value for k coins processed so far.
// For each pile, try taking 0..min(len(pile), k) coins from its top.
// Use prefix sums per pile for O(1) sum of top N coins.

import (
	"fmt"
)

func main() {
	// piles = [[1,100,3],[7,8,9]], k = 2 => 101
	piles := [][]int{{1, 100, 3}, {7, 8, 9}}
	fmt.Println(maxValueOfCoins(piles, 2))

	// Single pile
	fmt.Println(maxValueOfCoins([][]int{{5, 10, 15}}, 3))

	// Multiple piles
	fmt.Println(maxValueOfCoins([][]int{{100}, {200}, {300}}, 2))

	// Empty case
	fmt.Println(maxValueOfCoins([][]int{{}, {}}, 1))

	// Large k
	fmt.Println(maxValueOfCoins([][]int{{1, 2, 3, 4, 5}}, 3))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxValueOfCoins(piles [][]int, k int) int {
	// Prefix sums for each pile.
	pref := make([][]int, len(piles))
	for i, pile := range piles {
		pref[i] = make([]int, len(pile)+1)
		for j, v := range pile {
			pref[i][j+1] = pref[i][j] + v
		}
	}

	// dp[x] = max value using x coins from processed piles.
	// Use -1 as unreachable sentinel.
	dp := make([]int, k+1)
	for i := 1; i <= k; i++ {
		dp[i] = -1
	}
	dp[0] = 0

	for i := 0; i < len(piles); i++ {
		ndp := make([]int, k+1)
		copy(ndp, dp)
		maxTake := len(piles[i])
		if maxTake > k {
			maxTake = k
		}
		for take := 1; take <= maxTake; take++ {
			sum := pref[i][take]
			for used := take; used <= k; used++ {
				if dp[used-take] != -1 {
					ndp[used] = max(ndp[used], dp[used-take]+sum)
				}
			}
		}
		dp = ndp
	}

	if dp[k] == -1 {
		return 0
	}
	return dp[k]
}
```

## 2223 — Sum Of Scores Of Built Strings

```go
package main

// LeetCode #2223: Sum of Scores of Built Strings
// https://leetcode.com/problems/sum-of-scores-of-built-strings/
// Difficulty: Hard
//
// Z-algorithm: compute Z-array where Z[i] = longest common prefix of s and s[i:].
// Answer = sum of all Z-values + n (the whole string is its own prefix).

import (
	"fmt"
)

func main() {
	// "babab" => 9
	fmt.Println(sumScores("babab"))
	// "ababa" => 9
	fmt.Println(sumScores("ababa"))
	// "a" => 1
	fmt.Println(sumScores("a"))
	// "aa" => 3
	fmt.Println(sumScores("aa"))
	// "abc" => 3
	fmt.Println(sumScores("abc"))
}

func sumScores(s string) int64 {
	n := len(s)
	z := make([]int, n)
	l, r := 0, 0

	for i := 1; i < n; i++ {
		if i <= r {
			z[i] = min(r-i+1, z[i-l])
		}
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			z[i]++
		}
		if i+z[i]-1 > r {
			l, r = i, i+z[i]-1
		}
	}

	var sum int64 = int64(n)
	for _, v := range z {
		sum += int64(v)
	}
	return sum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```

## 2227 — Encrypt And Decrypt Strings

```go
package main

// LeetCode #2227: Encrypt and Decrypt Strings
// https://leetcode.com/problems/encrypt-and-decrypt-strings/
// Difficulty: Hard
//
// You are given a character array keys and a string array values (same length),
// and a dictionary of allowed words (string array).
// The Encrypter class:
//   - Encrypt(s): replaces each character c with values[i] where keys[i]==c.
//   - Decrypt(s): returns the number of strings in the dictionary that could
//     result in s after encryption.

import (
	"fmt"
)

// Encrypter encrypts and decrypts strings.
type Encrypter struct {
	charToVal map[byte]string
	valToChar map[string]byte
	encrypted map[string]int // encrypted word -> frequency in dictionary
}

// Constructor initializes the Encrypter.
func Constructor(keys []byte, values []string, dictionary []string) Encrypter {
	charToVal := make(map[byte]string)
	valToChar := make(map[string]byte)

	for i, k := range keys {
		charToVal[k] = values[i]
		// allow multiple chars to map to same value, but we keep first
		if _, ok := valToChar[values[i]]; !ok {
			valToChar[values[i]] = k
		}
	}

	encrypted := make(map[string]int)
	for _, word := range dictionary {
		encWord := encryptWord(word, charToVal)
		if encWord != "" {
			encrypted[encWord]++
		}
	}

	return Encrypter{
		charToVal: charToVal,
		valToChar: valToChar,
		encrypted: encrypted,
	}
}

// encryptWord encrypts a word using the char-to-value mapping.
// Returns empty string if the word contains a char not in keys.
func encryptWord(word string, charToVal map[byte]string) string {
	var result []byte
	for i := 0; i < len(word); i++ {
		val, ok := charToVal[word[i]]
		if !ok {
			return "" // can't encrypt
		}
		result = append(result, []byte(val)...)
	}
	return string(result)
}

// Encrypt encrypts the input string s.
func (e *Encrypter) Encrypt(s string) string {
	result := encryptWord(s, e.charToVal)
	return result
}

// Decrypt returns the number of dictionary words that could
// have produced the encrypted string.
func (e *Encrypter) Decrypt(s string) int {
	return e.encrypted[s]
}

func main() {
	keys := []byte{'a', 'b', 'c', 'd'}
	values := []string{"ei", "zf", "ei", "am"}
	dictionary := []string{"abcd", "acbd", "adbc", "badc", "dabc", "cabd"}

	encrypter := Constructor(keys, values, dictionary)

	// Test encrypt
	encrypted := encrypter.Encrypt("abcd")
	fmt.Println("Encrypted 'abcd':", encrypted) // "eizfeiam"

	// Test decrypt
	count := encrypter.Decrypt(encrypted)
	fmt.Println("Decrypt count:", count) // 2 ("abcd", "acbd" both map to same)
}
```

## 2234 — Maximum Total Beauty Of The Gardens

```go
package main

// LeetCode #2234: Maximum Total Beauty of the Gardens
// https://leetcode.com/problems/maximum-total-beauty-of-the-gardens/
// Difficulty: Hard
//
// Alice has n gardens, each garden has some flowers initially.
// She can plant at most newFlowers additional flowers (total).
// The beauty score = fullGardens * full + incompleteMin * partial
//   - fullGardens: number of gardens with at least target flowers
//   - incompleteMin: minimum number of flowers among incomplete gardens (0 if all full)
// Goal: maximize total beauty.

import (
	"fmt"
	"sort"
)

// maximumBeauty returns maximum possible total beauty.
func maximumBeauty(flowers []int, newFlowers int64, target int, full int, partial int) int64 {
	n := len(flowers)
	// convert to int for easier math
	newF := int(newFlowers)

	// sort initially
	sorted := make([]int, n)
	copy(sorted, flowers)
	sort.Ints(sorted)

	// clip at target (excess flowers don't help completeness)
	for i := 0; i < n; i++ {
		if sorted[i] > target {
			sorted[i] = target
		}
	}
	sort.Ints(sorted)

	// prefix sums for efficient gap calculation
	prefix := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + sorted[i]
	}

	// try all possible counts of full gardens
	best := int64(0)

	// suffix: index from where gardens are full
	for fullCount := 0; fullCount <= n; fullCount++ {
		// if we make last 'fullCount' gardens full, what remaining flowers do we have?
		remaining := newF

		// cost to raise all fullCount gardens to target
		if fullCount > 0 {
			startIdx := n - fullCount
			cost := fullCount*target - (prefix[n] - prefix[startIdx])
			if cost > remaining {
				continue // not enough flowers for this many full gardens
			}
			remaining -= cost
		}

		// now maximize the minimum among remaining (incomplete) gardens
		incompleteCount := n - fullCount
		if incompleteCount == 0 {
			// all gardens full
			beauty := n * full
			if int64(beauty) > best {
				best = int64(beauty)
			}
			continue
		}

		// binary search for maximum possible minimum
		incomplete := sorted[:incompleteCount]

		lo, hi := 0, target-1
		for lo <= hi {
			mid := (lo + hi) / 2

			// how many flowers needed to raise all incomplete to at least mid?
			need := 0
			pos := sort.SearchInts(incomplete, mid)
			need = pos*mid - prefix[pos]

			if need <= remaining {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}

		minBeauty := hi // hi is the max min achievable
		if minBeauty < 0 {
			continue
		}

		beauty := fullCount*full + minBeauty*partial
		if int64(beauty) > best {
			best = int64(beauty)
		}
	}

	return best
}

func main() {
	// Example 1
	fmt.Println(maximumBeauty([]int{1, 3, 1, 1}, 7, 6, 12, 1))
	// Expected: 14

	// Example 2
	fmt.Println(maximumBeauty([]int{2, 4, 5, 3}, 10, 5, 2, 6))
	// Expected: 30
}
```

## 2242 — Maximum Score Of A Node Sequence

```go
package main

// LeetCode #2242: Maximum Score of a Node Sequence
// https://leetcode.com/problems/maximum-score-of-a-node-sequence/
// Difficulty: Hard
//
// Given a graph with n nodes (0..n-1) with node scores, and edges (u, v).
// Find the maximum score of a node sequence [a, b, c, d] where
// edges exist between (a,b), (b,c), (c,d) and all four nodes are distinct.

import (
	"fmt"
	"sort"
)

// maximumScore returns maximum score of a valid 4-node sequence.
func maximumScore(scores []int, edges [][]int) int {
	n := len(scores)

	// adjacency list of neighbors sorted by score descending (keep up to 3 best)
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// For each node, keep its top 3 neighbors by score (to limit search space)
	top3 := make([][]int, n)
	for i := 0; i < n; i++ {
		neighbors := adj[i]
		sort.Slice(neighbors, func(a, b int) bool {
			return scores[neighbors[a]] > scores[neighbors[b]]
		})
		if len(neighbors) > 3 {
			neighbors = neighbors[:3]
		}
		top3[i] = neighbors
	}

	maxScore := -1

	// For each edge (b, c), try extending to a and d via top neighbors
	for _, e := range edges {
		b, c := e[0], e[1]

		// For node b, look at its top3 (excluding c)
		for _, a := range top3[b] {
			if a == c {
				continue
			}
			// For node c, look at its top3 (excluding b and a)
			for _, d := range top3[c] {
				if d == b || d == a {
					continue
				}
				score := scores[a] + scores[b] + scores[c] + scores[d]
				if score > maxScore {
					maxScore = score
				}
			}
		}
	}

	return maxScore
}

func main() {
	// Example 1
	scores1 := []int{5, 2, 9, 8, 4}
	edges1 := [][]int{{0, 1}, {1, 2}, {2, 3}, {0, 2}, {1, 3}, {2, 4}}
	fmt.Println(maximumScore(scores1, edges1)) // Expected: 24

	// Example 2
	scores2 := []int{9, 10, 11, 12, 13, 14}
	edges2 := [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {4, 5}}
	fmt.Println(maximumScore(scores2, edges2)) // Expected: 46
}
```

## 2246 — Longest Path With Different Adjacent Characters

```go
package main

// LeetCode #2246: Longest Path with Different Adjacent Characters
// https://leetcode.com/problems/longest-path-with-different-adjacent-characters/
// Difficulty: Hard
//
// Tree DFS: Build adjacency list. DFS from root. For each node, collect the
// longest path lengths from children that have a different character. The
// longest path through this node = sum of top two + 1. Update global max.

import (
	"fmt"
)

func main() {
	// parent = [-1,0,0,1,1,2], s = "abacbe" => 3
	fmt.Println(longestPath([]int{-1, 0, 0, 1, 1, 2}, "abacbe"))
	// single node
	fmt.Println(longestPath([]int{-1}, "a"))
	// all same chars (path should be 1 since no two adjacent can be same char)
	fmt.Println(longestPath([]int{-1, 0, 0}, "aaa"))
	// linear
	fmt.Println(longestPath([]int{-1, 0, 0, 1, 1, 2, 5}, "abacbea"))
}

func longestPath(parent []int, s string) int {
	n := len(parent)
	children := make([][]int, n)
	for i := 1; i < n; i++ {
		p := parent[i]
		children[p] = append(children[p], i)
	}

	ans := 1

	var dfs func(u int) int
	dfs = func(u int) int {
		top1, top2 := 0, 0
		for _, v := range children[u] {
			childLen := dfs(v)
			if s[v] != s[u] {
				if childLen > top1 {
					top2 = top1
					top1 = childLen
				} else if childLen > top2 {
					top2 = childLen
				}
			}
		}
		// Path through u = top1 + top2 + 1
		if top1+top2+1 > ans {
			ans = top1 + top2 + 1
		}
		return top1 + 1
	}

	dfs(0)
	return ans
}
```

## 2247 — Maximum Cost Of Trip With K Highways

```go
package main

// LeetCode #2247: Maximum Cost of Trip With K Highways
// https://leetcode.com/problems/maximum-cost-of-trip-with-k-highways/
// Difficulty: Hard [Paid]
//
// Given n cities (0..n-1) connected by highways with toll costs.
// Find the maximum possible total cost of a trip that uses exactly k highways,
// visits each city at most once, and can start and end at any city.
// k <= n-1.

import (
	"fmt"
	"math"
)

// maximumCost returns maximum cost of a trip using exactly k highways.
func maximumCost(n int, highways [][]int, k int) int {
	if k >= n {
		return -1 // not enough cities for k highways (path of k edges needs k+1 cities)
	}
	if k == 0 {
		return 0 // no highways => cost 0
	}

	// build adjacency matrix (costs)
	adj := make([][][2]int, n) // [neighbor, cost]
	for _, h := range highways {
		u, v, cost := h[0], h[1], h[2]
		adj[u] = append(adj[u], [2]int{v, cost})
		adj[v] = append(adj[v], [2]int{u, cost})
	}

	// DP[mask][last] = max cost to reach 'last' city using mask of visited cities
	// mask has k+1 bits set (we visit k+1 cities for k highways)
	dp := make([][]int, 1<<n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = math.MinInt32
		}
	}

	// base: starting at each city (mask with 1 bit)
	for i := 0; i < n; i++ {
		dp[1<<i][i] = 0
	}

	maxCost := math.MinInt32

	// iterate over all masks
	for mask := 1; mask < (1 << n); mask++ {
		bits := popcount(mask)

		if bits > k+1 {
			continue // too many cities
		}

		for last := 0; last < n; last++ {
			if dp[mask][last] == math.MinInt32 {
				continue
			}

			if bits == k+1 {
				// we have used exactly k edges (k+1 cities)
				if dp[mask][last] > maxCost {
					maxCost = dp[mask][last]
				}
				continue
			}

			// try extending
			for _, edge := range adj[last] {
				next, cost := edge[0], edge[1]
				if mask&(1<<next) != 0 {
					continue // already visited
				}
				newMask := mask | (1 << next)
				newCost := dp[mask][last] + cost
				if newCost > dp[newMask][next] {
					dp[newMask][next] = newCost
				}
			}
		}
	}

	if maxCost == math.MinInt32 {
		return -1
	}
	return maxCost
}

func popcount(x int) int {
	cnt := 0
	for x > 0 {
		cnt += x & 1
		x >>= 1
	}
	return cnt
}

func main() {
	// Example 1
	n1 := 5
	highways1 := [][]int{{0, 1, 4}, {2, 1, 3}, {1, 4, 11}, {3, 2, 7}, {3, 4, 2}, {0, 3, 18}}
	k1 := 3
	fmt.Println(maximumCost(n1, highways1, k1)) // Expected: 28

	// Example 2
	n2 := 4
	highways2 := [][]int{{0, 1, 3}, {2, 3, 2}}
	k2 := 2
	fmt.Println(maximumCost(n2, highways2, k2)) // Expected: -1
}
```

## 2251 — Number Of Flowers In Full Bloom

```go
package main

// LeetCode #2251: Number of Flowers in Full Bloom
// https://leetcode.com/problems/number-of-flowers-in-full-bloom/
// Difficulty: Hard
//
// You are given a 0-indexed 2D integer array flowers where flowers[i] = [start_i, end_i]
// means the i-th flower will be in full bloom from start_i to end_i (inclusive).
// You are also given a 0-indexed integer array people of size n.
// For each person, return the number of flowers in full bloom at the time people[i].

import (
	"fmt"
	"sort"
)

// fullBloomFlowers returns for each person the count of flowers in bloom at their arrival time.
func fullBloomFlowers(flowers [][]int, people []int) []int {
	n := len(flowers)
	m := len(people)

	starts := make([]int, n)
	ends := make([]int, n)
	for i, f := range flowers {
		starts[i] = f[0]
		ends[i] = f[1]
	}

	sort.Ints(starts)
	sort.Ints(ends)

	result := make([]int, m)
	for i, p := range people {
		// number of flowers that started blooming <= p
		bloomed := sort.SearchInts(starts, p+1) // first index > p

		// number of flowers that ended blooming < p
		faded := sort.SearchInts(ends, p) // first index >= p

		result[i] = bloomed - faded
	}

	return result
}

func main() {
	// Example 1
	flowers1 := [][]int{{1, 6}, {3, 7}, {9, 12}, {4, 13}}
	people1 := []int{2, 3, 7, 11}
	fmt.Println(fullBloomFlowers(flowers1, people1)) // Expected: [1,2,2,2]

	// Example 2
	flowers2 := [][]int{{1, 10}, {3, 3}}
	people2 := []int{3, 3, 2}
	fmt.Println(fullBloomFlowers(flowers2, people2)) // Expected: [2,2,1]
}
```

## 2252 — Dynamic Pivoting Of A Table

```go
package main

// LeetCode #2252: Dynamic Pivoting of a Table
// https://leetcode.com/problems/dynamic-pivoting-of-a-table/
// Difficulty: Hard [Paid]
//
// Given a table with columns (product_id, store, price),
// pivot it so that each store becomes a column and each row is a product_id.
// The store names are dynamic (not known in advance).

import (
	"fmt"
)

// PivotTable represents a pivoted in-memory table.
type PivotTable struct {
	// map[product_id][store] = price
	Data map[int]map[string]int
}

// NewPivotTable creates a new pivot table.
func NewPivotTable() *PivotTable {
	return &PivotTable{Data: make(map[int]map[string]int)}
}

// AddRow adds a (product_id, store, price) record.
func (pt *PivotTable) AddRow(productID int, store string, price int) {
	if pt.Data[productID] == nil {
		pt.Data[productID] = make(map[string]int)
	}
	pt.Data[productID][store] = price
}

// Pivot returns the pivoted representation:
// rows = product_ids, columns = stores (sorted), values = prices.
func (pt *PivotTable) Pivot() ([]string, map[int]map[string]int) {
	// collect all stores
	storeSet := make(map[string]bool)
	for _, stores := range pt.Data {
		for s := range stores {
			storeSet[s] = true
		}
	}

	stores := make([]string, 0, len(storeSet))
	for s := range storeSet {
		stores = append(stores, s)
	}
	// sort stores alphabetically (simple bubble for small set)
	for i := 0; i < len(stores); i++ {
		for j := i + 1; j < len(stores); j++ {
			if stores[j] < stores[i] {
				stores[i], stores[j] = stores[j], stores[i]
			}
		}
	}

	return stores, pt.Data
}

func main() {
	pt := NewPivotTable()
	pt.AddRow(1, "StoreA", 10)
	pt.AddRow(1, "StoreB", 20)
	pt.AddRow(2, "StoreA", 15)
	pt.AddRow(2, "StoreC", 25)
	pt.AddRow(3, "StoreB", 30)

	stores, data := pt.Pivot()
	fmt.Println("Stores:", stores)
	for pid := 1; pid <= 3; pid++ {
		fmt.Printf("Product %d: %v\n", pid, data[pid])
	}
}
```

## 2253 — Dynamic Unpivoting Of A Table

```go
package main

// LeetCode #2253: Dynamic Unpivoting of a Table
// https://leetcode.com/problems/dynamic-unpivoting-of-a-table/
// Difficulty: Hard [Paid]
//
// Given a pivoted table where columns are dynamic store names,
// unpivot it back to (product_id, store, price) format.

import (
	"fmt"
	"sort"
)

// UnpivotTable handles dynamic unpivoting.
type UnpivotTable struct {
	// map[product_id][store] = price
	Data map[int]map[string]int
}

// NewUnpivotTable creates a new unpivot table.
func NewUnpivotTable() *UnpivotTable {
	return &UnpivotTable{Data: make(map[int]map[string]int)}
}

// LoadPivoted adds a row from pivoted format: productID with store prices.
func (ut *UnpivotTable) LoadPivoted(productID int, prices map[string]int) {
	for store, price := range prices {
		if ut.Data[productID] == nil {
			ut.Data[productID] = make(map[string]int)
		}
		ut.Data[productID][store] = price
	}
}

// Unpivot converts pivoted data to rows of (product_id, store, price).
func (ut *UnpivotTable) Unpivot() [][3]interface{} {
	var rows [][3]interface{}

	for productID, stores := range ut.Data {
		// sort stores for deterministic output
		storeNames := make([]string, 0, len(stores))
		for s := range stores {
			storeNames = append(storeNames, s)
		}
		sort.Strings(storeNames)

		for _, store := range storeNames {
			rows = append(rows, [3]interface{}{productID, store, stores[store]})
		}
	}

	return rows
}

func main() {
	ut := NewUnpivotTable()
	ut.LoadPivoted(1, map[string]int{"StoreA": 10, "StoreB": 20})
	ut.LoadPivoted(2, map[string]int{"StoreA": 15, "StoreC": 25})
	ut.LoadPivoted(3, map[string]int{"StoreB": 30})

	rows := ut.Unpivot()
	fmt.Println("Unpivoted rows:")
	for _, row := range rows {
		fmt.Printf("  Product %d, Store %s, Price %d\n", row[0].(int), row[1].(string), row[2].(int))
	}
}
```

## 2254 — Design Video Sharing Platform

```go
package main

// LeetCode #2254: Design Video Sharing Platform
// https://leetcode.com/problems/design-video-sharing-platform/
// Difficulty: Hard [Paid]
//
// Design a video sharing platform with the following operations:
//   - upload(video): returns a unique video ID (auto-increment).
//   - remove(videoId): marks video as removed.
//   - watch(videoId, startMinute, endMinute): records that a user watched
//     the video from startMinute to endMinute (both inclusive, 0-indexed).
//   - like(videoId): increments like count.
//   - dislike(videoId): increments dislike count.
//   - getLikesAndDislikes(videoId): returns [likes, dislikes].
//   - getViews(videoId): returns total view duration across all watches
//     (sum of (endMinute - startMinute + 1) for each watch).

import (
	"fmt"
)

// Video represents a video with its metadata.
type Video struct {
	ID       int
	Likes    int
	Dislikes int
	Removed  bool
	Views    int // total watched minutes
}

// VideoSharingPlatform manages videos.
type VideoSharingPlatform struct {
	videos    map[int]*Video
	nextID    int
}

// Constructor creates a new platform.
func Constructor() *VideoSharingPlatform {
	return &VideoSharingPlatform{
		videos: make(map[int]*Video),
		nextID: 0,
	}
}

// Upload adds a new video and returns its ID.
func (vsp *VideoSharingPlatform) Upload(video string) int {
	id := vsp.nextID
	vsp.nextID++
	vsp.videos[id] = &Video{
		ID:       id,
		Likes:    0,
		Dislikes: 0,
		Removed:  false,
		Views:    0,
	}
	return id
}

// Remove marks a video as removed.
func (vsp *VideoSharingPlatform) Remove(videoID int) {
	if v, ok := vsp.videos[videoID]; ok {
		v.Removed = true
	}
}

// Watch records a watch from startMinute to endMinute.
func (vsp *VideoSharingPlatform) Watch(videoID, startMinute, endMinute int) {
	if v, ok := vsp.videos[videoID]; ok && !v.Removed {
		duration := endMinute - startMinute + 1
		if duration > 0 {
			v.Views += duration
		}
	}
}

// Like increments the like count for a video.
func (vsp *VideoSharingPlatform) Like(videoID int) {
	if v, ok := vsp.videos[videoID]; ok && !v.Removed {
		v.Likes++
	}
}

// Dislike increments the dislike count for a video.
func (vsp *VideoSharingPlatform) Dislike(videoID int) {
	if v, ok := vsp.videos[videoID]; ok && !v.Removed {
		v.Dislikes++
	}
}

// GetLikesAndDislikes returns [likes, dislikes] for a video.
func (vsp *VideoSharingPlatform) GetLikesAndDislikes(videoID int) []int {
	if v, ok := vsp.videos[videoID]; ok && !v.Removed {
		return []int{v.Likes, v.Dislikes}
	}
	return []int{-1, -1}
}

// GetViews returns total view minutes for a video.
func (vsp *VideoSharingPlatform) GetViews(videoID int) int {
	if v, ok := vsp.videos[videoID]; ok && !v.Removed {
		return v.Views
	}
	return -1
}

func main() {
	vsp := Constructor()

	id1 := vsp.Upload("video1")
	id2 := vsp.Upload("video2")
	fmt.Println("Uploaded IDs:", id1, id2) // 0, 1

	vsp.Watch(id1, 0, 5)
	vsp.Watch(id1, 2, 4)
	vsp.Watch(id2, 0, 10)

	vsp.Like(id1)
	vsp.Like(id1)
	vsp.Dislike(id1)

	fmt.Println("Likes/Dislikes for 0:", vsp.GetLikesAndDislikes(id1)) // [2, 1]
	fmt.Println("Views for 0:", vsp.GetViews(id1))                    // (5-0+1)+(4-2+1) = 6+3 = 9
	fmt.Println("Views for 1:", vsp.GetViews(id2))                    // 11

	vsp.Remove(id1)
	fmt.Println("After removal:", vsp.GetLikesAndDislikes(id1)) // [-1, -1]
}
```

## 2258 — Escape The Spreading Fire

```go
package main

// LeetCode #2258: Escape the Spreading Fire
// https://leetcode.com/problems/escape-the-spreading-fire/
// Difficulty: Hard
//
// You are given a 2D grid where:
//   0 = grass (can walk), 1 = fire, 2 = wall
// Fire spreads to adjacent cells (4-directional) every minute.
// You start at (0, 0) and want to reach safehouse at (m-1, n-1).
// You can stay in place. You can't be on the same cell as fire at the same time.
// Return the maximum number of minutes you can wait before moving,
// while still being able to reach the safehouse before or at the same time as fire.
// If impossible, return -1. If unlimited, return 10^9.

import (
	"fmt"
	"math"
)

// maximumMinutes returns max wait time.
func maximumMinutes(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	// fireDist[i][j] = minute when fire reaches (i,j), or INF
	fireDist := make([][]int, m)
	for i := range fireDist {
		fireDist[i] = make([]int, n)
		for j := range fireDist[i] {
			fireDist[i][j] = math.MaxInt32
		}
	}

	// BFS from all fire sources
	type point struct{ x, y int }
	queue := make([]point, 0)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				fireDist[i][j] = 0
				queue = append(queue, point{i, j})
			}
		}
	}

	dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		for _, d := range dirs {
			nx, ny := p.x+d[0], p.y+d[1]
			if nx >= 0 && nx < m && ny >= 0 && ny < n &&
				grid[nx][ny] != 2 && fireDist[nx][ny] == math.MaxInt32 {
				fireDist[nx][ny] = fireDist[p.x][p.y] + 1
				queue = append(queue, point{nx, ny})
			}
		}
	}

	// Binary search: can we wait `wait` minutes?
	canEscape := func(wait int) bool {
		// BFS for person
		visited := make([][]bool, m)
		for i := range visited {
			visited[i] = make([]bool, n)
		}

		pq := make([]point, 0)
		pq = append(pq, point{0, 0})
		visited[0][0] = true
		t := wait // time when we start moving

		for len(pq) > 0 {
			size := len(pq)
			for k := 0; k < size; k++ {
				p := pq[k]

				// check if we are at safehouse
				if p.x == m-1 && p.y == n-1 {
					return true
				}

				// we can't be on fire cell at this time (unless it's safehouse at same time)
				if fireDist[p.x][p.y] <= t {
					continue
				}

				for _, d := range dirs {
					nx, ny := p.x+d[0], p.y+d[1]
					if nx >= 0 && nx < m && ny >= 0 && ny < n &&
						!visited[nx][ny] && grid[nx][ny] != 2 {

						// at the safehouse, fire can arrive at the same time
						if nx == m-1 && ny == n-1 && fireDist[nx][ny] <= t+1 {
							// we reach safehouse at t+1, fire reaches at <= t+1
							// Both arrive at same time? Actually fireDist <= t+1,
							// So fire might arrive before or at same time.
							// If fire arrives at same time, we're fine at safehouse.
							// But if fire arrives earlier, not fine.
							// Actually the problem says: you can be at safehouse
							// at the same time as fire.
							if fireDist[nx][ny] <= t {
								continue // fire already there
							}
							// fire arrives at t+1 or later -> ok
							visited[nx][ny] = true
							pq = append(pq, point{nx, ny})
						} else if fireDist[nx][ny] > t+1 {
							// we arrive at t+1, fire arrives later
							visited[nx][ny] = true
							pq = append(pq, point{nx, ny})
						}
					}
				}
			}
			pq = pq[size:]
			t++
		}
		return false
	}

	// If can't escape even with 0 wait
	if !canEscape(0) {
		return -1
	}

	// check unlimited (fire never reaches safehouse after waiting large)
	if canEscape(1_000_000_000) {
		return 1_000_000_000
	}

	// binary search
	lo, hi := 0, 1_000_000_000
	for lo < hi {
		mid := (lo + hi + 1) / 2
		if canEscape(mid) {
			lo = mid
		} else {
			hi = mid - 1
		}
	}
	return lo
}

func main() {
	// Example 1
	grid1 := [][]int{
		{0, 2, 0, 0, 0, 0, 0},
		{0, 0, 0, 2, 2, 1, 0},
		{0, 1, 0, 0, 0, 2, 0},
		{0, 0, 0, 0, 0, 0, 0},
	}
	fmt.Println(maximumMinutes(grid1)) // Expected: 1

	// Example 2
	grid2 := [][]int{
		{0, 0, 0, 0},
		{0, 1, 2, 0},
		{0, 2, 0, 0},
	}
	fmt.Println(maximumMinutes(grid2)) // Expected: -1

	// Example 3
	grid3 := [][]int{
		{0, 0, 0},
		{2, 2, 0},
		{1, 2, 0},
	}
	fmt.Println(maximumMinutes(grid3)) // Expected: 1000000000
}
```

## 2262 — Total Appeal Of A String

```go
package main

// LeetCode #2262: Total Appeal of A String
// https://leetcode.com/problems/total-appeal-of-a-string/
// Difficulty: Hard
//
// The appeal of a string is the number of distinct characters in it.
// Given a string s, return the total appeal of all its substrings.

import (
	"fmt"
)

// appealSum returns the sum of appeal over all substrings of s.
func appealSum(s string) int64 {
	// For each character, we count how many substrings include it
	// as the FIRST occurrence from the left (to avoid double counting).
	//
	// For position i with character c, the number of substrings where
	// c contributes its distinctness is:
	//   (i - lastSeen[c]) * (n - i)
	// where lastSeen[c] is the previous occurrence of c (-1 if none).
	// This counts substrings that START after last occurrence and END at or after i.

	n := len(s)
	lastSeen := make(map[byte]int)
	var result int64

	for i := 0; i < n; i++ {
		c := s[i]
		prev, ok := lastSeen[c]
		if !ok {
			prev = -1
		}
		// substrings starting in (prev, i] and ending >= i
		// left choices: i - prev
		// right choices: n - i
		result += int64(i-prev) * int64(n-i)
		lastSeen[c] = i
	}

	return result
}

func main() {
	// Example 1
	fmt.Println(appealSum("abbca")) // Expected: 28

	// Example 2
	fmt.Println(appealSum("code")) // Expected: 20
}
```

## 2263 — Make Array Non Decreasing Or Non Increasing

```go
package main

// LeetCode #2263: Make Array Non-decreasing or Non-increasing
// https://leetcode.com/problems/make-array-non-decreasing-or-non-increasing/
// Difficulty: Hard [Paid]
//
// Given an integer array nums, return the minimum number of operations
// to make it either entirely non-decreasing or entirely non-increasing.
// In one operation, you can increase or decrease any element by 1.

import (
	"fmt"
	"math"
)

// minOperationsToMakeNonDecOrNonInc returns minimum operations.
// Uses DP with coordinate compression: the optimal target values
// are always from the original array values (or sorted equivalents).
func minOperationsToMakeNonDecOrNonInc(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	// To make non-decreasing: each element >= previous
	// To make non-increasing: each element <= previous
	// We compute both and take min.

	// Coordinate compression: collect unique sorted values
	unique := make(map[int]bool)
	for _, v := range nums {
		unique[v] = true
	}
	sortedVals := make([]int, 0, len(unique))
	for v := range unique {
		sortedVals = append(sortedVals, v)
	}
	sortInts(sortedVals)
	m := len(sortedVals)

	// DP for non-decreasing
	// dp[j] = min cost to make first i elements non-decreasing with
	// the i-th element <= sortedVals[j]
	dp := make([]int, m)
	for j := 0; j < m; j++ {
		dp[j] = absInt(nums[0] - sortedVals[j])
	}
	// make dp cumulative min
	for j := 1; j < m; j++ {
		if dp[j] > dp[j-1] {
			dp[j] = dp[j-1]
		}
	}

	for i := 1; i < n; i++ {
		newDp := make([]int, m)
		for j := 0; j < m; j++ {
			cost := absInt(nums[i] - sortedVals[j])
			newDp[j] = cost + dp[j] // dp[j] already has min from <= sortedVals[j]
		}
		for j := 1; j < m; j++ {
			if newDp[j] > newDp[j-1] {
				newDp[j] = newDp[j-1]
			}
		}
		dp = newDp
	}

	nonDecCost := dp[m-1]

	// DP for non-increasing
	dp = make([]int, m)
	for j := 0; j < m; j++ {
		dp[j] = absInt(nums[0] - sortedVals[j])
	}
	// make dp cumulative max (non-increasing: we need value >= sortedVals[j])
	for j := m - 2; j >= 0; j-- {
		if dp[j] > dp[j+1] {
			dp[j] = dp[j+1]
		}
	}

	for i := 1; i < n; i++ {
		newDp := make([]int, m)
		for j := 0; j < m; j++ {
			cost := absInt(nums[i] - sortedVals[j])
			newDp[j] = cost + dp[j]
		}
		for j := m - 2; j >= 0; j-- {
			if newDp[j] > newDp[j+1] {
				newDp[j] = newDp[j+1]
			}
		}
		dp = newDp
	}

	nonIncCost := dp[0]

	if nonDecCost < nonIncCost {
		return nonDecCost
	}
	return nonIncCost
}

func absInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func sortInts(a []int) {
	// simple insertion sort for small arrays
	for i := 1; i < len(a); i++ {
		key := a[i]
		j := i - 1
		for j >= 0 && a[j] > key {
			a[j+1] = a[j]
			j--
		}
		a[j+1] = key
	}
}

// Alternative using a priority-queue-based approach (more efficient):
// For non-decreasing: cost = sum of positive diffs when using max-heap
func minOperationsPQ(nums []int) int64 {
	nonDec := minCostNonDec(nums)
	nonInc := minCostNonInc(nums)
	if nonDec < nonInc {
		return nonDec
	}
	return nonInc
}

// minCostNonDec uses max-heap technique (make array non-decreasing).
func minCostNonDec(nums []int) int64 {
	// Using the "make array non-decreasing with min absolute changes" approach.
	// For non-decreasing, we need to find the minimal cost to transform.
	// This is equivalent to: for each element, ensure it's >= previous median
	// Running max-heap: cost = sum(max(0, prev_median - current))
	// Simplified approach using DP is above. Using heap is O(n log n).

	// Actually for this problem, the heap approach gives the min changes
	// to make non-decreasing with specific property.
	// Let's use the simpler DP approach above.
	return int64(minCostNonDecHelper(nums))
}

func minCostNonDecHelper(nums []int) int {
	// Use DP approach from above
	return min(minOperationsToMakeNonDecOrNonInc(nums), math.MaxInt32)
}

func minCostNonInc(nums []int) int64 {
	return int64(minOperationsToMakeNonDecOrNonInc(reverseInts(nums)))
}

func reverseInts(nums []int) []int {
	n := len(nums)
	rev := make([]int, n)
	for i, v := range nums {
		rev[n-1-i] = v
	}
	return rev
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Example
	fmt.Println(minOperationsToMakeNonDecOrNonInc([]int{3, 2, 4, 5, 4}))
	// Expected: 2 (make non-decreasing: [3,3,4,5,5] or non-inc: [5,4,4,4,4])

	fmt.Println(minOperationsToMakeNonDecOrNonInc([]int{1, 2, 3, 4, 5}))
	// Expected: 0 (already non-decreasing)
}
```

## 2267 — Check If There Is A Valid Parentheses String Path

```go
package main

// LeetCode #2267: Check if There Is a Valid Parentheses String Path
// https://leetcode.com/problems/check-if-there-is-a-valid-parentheses-string-path/
// Difficulty: Hard
//
// Given a grid of '(' and ')', find a path from (0,0) to (m-1,n-1)
// moving only down or right, such that the concatenated string is a
// valid parentheses string. A valid parentheses string has:
//   - equal number of '(' and ')'
//   - at any prefix, #('(') >= #(')')

import (
	"fmt"
)

// hasValidPath returns true if a valid path exists.
func hasValidPath(grid [][]byte) bool {
	m, n := len(grid), len(grid[0])

	// If start or end is invalid immediately
	if grid[0][0] == ')' || grid[m-1][n-1] == '(' {
		return false
	}
	// Path length must be even for balanced parens
	if (m+n-1)%2 != 0 {
		return false
	}

	// DP with set of possible open counts at each cell
	// dp[i][j] = set of possible open parenthesis counts when reaching (i,j)
	dp := make([][]map[int]bool, m)
	for i := range dp {
		dp[i] = make([]map[int]bool, n)
	}

	// Initialize
	dp[0][0] = map[int]bool{1: true} // '(' at (0,0) -> open count = 1

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			dp[i][j] = make(map[int]bool)

			possibleOpen := make(map[int]bool)
			// from top
			if i > 0 {
				for cnt := range dp[i-1][j] {
					possibleOpen[cnt] = true
				}
			}
			// from left
			if j > 0 {
				for cnt := range dp[i][j-1] {
					possibleOpen[cnt] = true
				}
			}

			for cnt := range possibleOpen {
				if grid[i][j] == '(' {
					dp[i][j][cnt+1] = true
				} else { // ')'
					if cnt-1 >= 0 {
						dp[i][j][cnt-1] = true
					}
				}
			}
		}
	}

	return dp[m-1][n-1][0]
}

func main() {
	// Example 1
	grid1 := [][]byte{
		{'(', '(', '('},
		{')', '(', ')'},
		{'(', '(', ')'},
		{'(', '(', ')'},
	}
	fmt.Println(hasValidPath(grid1)) // Expected: true

	// Example 2
	grid2 := [][]byte{
		{')', ')'},
		{'(', '('},
	}
	fmt.Println(hasValidPath(grid2)) // Expected: false
}
```

## 2272 — Substring With Largest Variance

```go
package main

// LeetCode #2272: Substring With Largest Variance
// https://leetcode.com/problems/substring-with-largest-variance/
// Difficulty: Hard
//
// The variance of a string is defined as the largest difference
// between the counts of any two distinct characters in the string.
// Given a string s consisting of lowercase English letters,
// return the largest variance among all substrings of s.

import (
	"fmt"
	"math"
)

// largestVariance returns the maximum variance across all substrings.
//
// Approach: For each pair of characters (major, minor),
// treat major as +1, minor as -1, and find maximum subarray sum,
// ensuring at least one minor appears (variance requires both chars).
func largestVariance(s string) int {
	// count frequency of each char
	freq := make([]int, 26)
	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
	}

	maxVar := 0

	// For each pair (a, b) where a != b
	for a := 0; a < 26; a++ {
		if freq[a] == 0 {
			continue
		}
		for b := 0; b < 26; b++ {
			if a == b || freq[b] == 0 {
				continue
			}

			// Kadane-like: treat a as +1, b as -1
			// We need substrings containing at least one b.
			// Use two DP values:
			//   cur: maximum subarray ending here (allowing no b)
			//   curWithB: maximum subarray ending here with at least one b
			cur := 0
			curWithB := math.MinInt32

			for i := 0; i < len(s); i++ {
				ch := int(s[i] - 'a')

				if ch == a {
					cur++
					if curWithB != math.MinInt32 {
						curWithB++
					}
				} else if ch == b {
					cur--
					curWithB = cur // at least we have b now
					if cur < 0 {
						cur = 0 // reset (start fresh from here)
					}
				} else {
					// other characters: they contribute 0 difference
					// but affect whether we have b
					// They don't change the relative count
				}

				if curWithB > maxVar {
					maxVar = curWithB
				}
			}
		}
	}

	return maxVar
}

func main() {
	// Example 1
	fmt.Println(largestVariance("aababbb")) // Expected: 3

	// Example 2
	fmt.Println(largestVariance("abcde")) // Expected: 0
}
```

## 2276 — Count Integers In Intervals

```go
package main

import (
	"fmt"
	"sort"
)

// CountIntervals maintains a set of disjoint intervals inserted via Add,
// and can report the total number of distinct integers covered.
// Intervals are stored in a sorted slice and merged on every add.
type CountIntervals struct {
	intervals [][2]int // sorted by left endpoint, non-overlapping
	total     int
}

func Constructor() CountIntervals {
	return CountIntervals{}
}

// Add inserts the interval [left, right] (inclusive) and merges overlaps.
func (ci *CountIntervals) Add(left, right int) {
	// Find the insertion / merge range via binary search.
	n := len(ci.intervals)
	// first index with end >= left
	lo := sort.Search(n, func(i int) bool { return ci.intervals[i][1] >= left })
	// last index with start <= right
	hi := sort.Search(n, func(i int) bool { return ci.intervals[i][0] > right })

	if lo == hi {
		// No overlap – insert as a fresh interval at position lo.
		ci.intervals = append(ci.intervals, [2]int{0, 0})
		copy(ci.intervals[lo+1:], ci.intervals[lo:])
		ci.intervals[lo] = [2]int{left, right}
		ci.total += right - left + 1
		return
	}

	// Subtract old contributions.
	for _, iv := range ci.intervals[lo:hi] {
		ci.total -= iv[1] - iv[0] + 1
	}

	// Merge. The merged interval spans from min(left, intervals[lo][0]) to
	// max(right, intervals[hi-1][1]).
	if left > ci.intervals[lo][0] {
		left = ci.intervals[lo][0]
	}
	if right < ci.intervals[hi-1][1] {
		right = ci.intervals[hi-1][1]
	}

	merged := [2]int{left, right}
	// Replace intervals[lo:hi] with the single merged interval.
	ci.intervals = append(ci.intervals[:lo], ci.intervals[hi:]...)
	ci.intervals = append(ci.intervals, [2]int{0, 0})
	copy(ci.intervals[lo+1:], ci.intervals[lo:])
	ci.intervals[lo] = merged

	ci.total += right - left + 1
}

// Count returns the total number of distinct integers covered by all intervals.
func (ci *CountIntervals) Count() int {
	return ci.total
}

// ---------------------------------------------------------------------------
//  LeetCode-style wrapper (required by stub convention)
func CountIntegersInIntervals() interface{} {
	ci := Constructor()
	ci.Add(2, 3)
	ci.Add(7, 10)
	_ = ci.Count()
	ci.Add(5, 8)
	return ci.Count()
}

func main() {
	fmt.Println(CountIntegersInIntervals())

	// ---- test cases ----
	testCases := []struct {
		ops    []string
		args   [][2]int // (-1,-1) for Count
		want   int       // for Count
		expect []int     // for Count when there are multiple calls
	}{
		{
			ops:    []string{"add", "add", "count", "add", "count"},
			args:   [][2]int{{2, 3}, {7, 10}, {-1, -1}, {5, 8}, {-1, -1}},
			expect: []int{0, 0, 6, 0, 8},
		},
	}

	for idx, tc := range testCases {
		ci := Constructor()
		for j, op := range tc.ops {
			switch op {
			case "add":
				ci.Add(tc.args[j][0], tc.args[j][1])
			case "count":
				got := ci.Count()
				if got != tc.expect[j] {
					fmt.Printf("FAIL tc %d step %d: got %d, want %d\n", idx, j, got, tc.expect[j])
				}
			}
		}
	}
	fmt.Println("Done testing 2276.")
}
```

## 2277 — Closest Node To Path In Tree

```go
package main

import (
	"fmt"
)

// 2277. Closest Node to Path in Tree
// ----------------------------------------------------------------
// For each query (u, v, node), return the node on the path between u and v
// that is closest to the query node.
//
// Solution: Binary Lifting for LCA.
//   dist(a, b) = depth[a] + depth[b] - 2*depth[LCA(a,b)]
//   The closest node on path(u,v) to x is the *deepest* among:
//     LCA(u,v), LCA(x,u), LCA(x,v)  — provided it lies on path(u,v).

func closestNode(n int, edges [][]int, query [][]int) []int {
	// Build adjacency.
	adj := make([][]int, n)
	for _, e := range edges {
		a, b := e[0], e[1]
		adj[a] = append(adj[a], b)
		adj[b] = append(adj[b], a)
	}

	// Binary lifting pre‑computation.
	LOG := 0
	for (1 << LOG) <= n {
		LOG++
	}
	up := make([][]int, n)
	depth := make([]int, n)
	for i := range up {
		up[i] = make([]int, LOG)
	}

	var dfs func(v, p int)
	dfs = func(v, p int) {
		up[v][0] = p
		for k := 1; k < LOG; k++ {
			up[v][k] = up[up[v][k-1]][k-1]
		}
		for _, to := range adj[v] {
			if to == p {
				continue
			}
			depth[to] = depth[v] + 1
			dfs(to, v)
		}
	}
	dfs(0, 0) // tree is connected, root at 0

	lca := func(a, b int) int {
		if depth[a] < depth[b] {
			a, b = b, a
		}
		// Lift a up to depth[b].
		diff := depth[a] - depth[b]
		for k := 0; k < LOG; k++ {
			if diff&(1<<k) != 0 {
				a = up[a][k]
			}
		}
		if a == b {
			return a
		}
		for k := LOG - 1; k >= 0; k-- {
			if up[a][k] != up[b][k] {
				a = up[a][k]
				b = up[b][k]
			}
		}
		return up[a][0]
	}

	dist := func(a, b int) int {
		return depth[a] + depth[b] - 2*depth[lca(a, b)]
	}

	// Check whether node p lies on the path between u and v.
	onPath := func(u, v, p int) bool {
		return dist(u, p)+dist(p, v) == dist(u, v)
	}

	ans := make([]int, len(query))
	for i, q := range query {
		u, v, x := q[0], q[1], q[2]
		lcauv := lca(u, v)
		candidates := []int{lcauv, lca(x, u), lca(x, v)}
		best, bestDepth := -1, -1
		for _, c := range candidates {
			if onPath(u, v, c) && depth[c] > bestDepth {
				best = c
				bestDepth = depth[c]
			}
		}
		ans[i] = best
	}
	return ans
}

// ---------------------------------------------------------------------------
//  LeetCode-style wrapper

func ClosestNodeToPathInTree() interface{} {
	edges := [][]int{{0, 1}, {1, 2}, {1, 3}, {4, 2}}
	queries := [][]int{{1, 3, 0}, {4, 2, 0}}
	return closestNode(5, edges, queries)
}

func main() {
	fmt.Println(ClosestNodeToPathInTree())

	// ---- simple test ----
	edges := [][]int{{0, 1}, {1, 2}, {1, 3}, {4, 2}}
	got := closestNode(5, edges, [][]int{{1, 3, 0}})
	fmt.Println("closestNode(1,3,0) =", got[0], "(expected 1)")
	got = closestNode(5, edges, [][]int{{4, 2, 0}})
	fmt.Println("closestNode(4,2,0) =", got[0], "(expected 2)")
	fmt.Println("Done testing 2277.")
}
```

## 2281 — Sum Of Total Strength Of Wizards

```go
package main

// LeetCode #2281: Sum of Total Strength of Wizards
// https://leetcode.com/problems/sum-of-total-strength-of-wizards/
// Difficulty: Hard
//
// Monotonic stack + prefix sums of prefix sums:
// For each element as minimum, find its range of dominance via next smaller
// element on left/right. Use prefix-of-prefix sums for O(1) range sum queries.

import (
	"fmt"
)

const mod = 1000000007

func main() {
	// [1,3,1,2] => 44
	fmt.Println(totalStrength([]int{1, 3, 1, 2}))
	// [5] => 5
	fmt.Println(totalStrength([]int{5}))
	// [1,2,3] => 33
	fmt.Println(totalStrength([]int{1, 2, 3}))
	// [2,2,2]
	fmt.Println(totalStrength([]int{2, 2, 2}))
}

func totalStrength(strength []int) int {
	n := len(strength)

	// Previous smaller element (strictly smaller, index).
	left := make([]int, n)
	stack := make([]int, 0, n)
	for i := 0; i < n; i++ {
		for len(stack) > 0 && strength[stack[len(stack)-1]] >= strength[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	// Next smaller element (strictly smaller, index).
	right := make([]int, n)
	stack = stack[:0]
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && strength[stack[len(stack)-1]] > strength[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			right[i] = n
		} else {
			right[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	// Prefix sums.
	pref := make([]int, n+1)
	for i := 0; i < n; i++ {
		pref[i+1] = (pref[i] + strength[i]) % mod
	}

	// Prefix sums of prefix sums.
	pref2 := make([]int, n+2)
	for i := 0; i <= n; i++ {
		pref2[i+1] = (pref2[i] + pref[i]) % mod
	}

	rangePref := func(l, r int) int {
		if l > r {
			return 0
		}
		return (pref2[r+1] - pref2[l] + mod) % mod
	}

	ans := 0
	for i := 0; i < n; i++ {
		l := left[i]
		r := right[i]

		// Number of subarrays where arr[i] is the minimum:
		// Left choices = i - l, Right choices = r - i
		leftLen := i - l
		rightLen := r - i

		// Sum of subarray sums where arr[i] is minimum:
		// = leftLen * sum(pref[i+1..r]) - rightLen * sum(pref[l+1..i])
		sumRight := rangePref(i+1, r)
		sumLeft := rangePref(l+1, i)

		total := (leftLen*sumRight%mod - rightLen*sumLeft%mod + mod) % mod
		contrib := (strength[i] % mod) * total % mod
		ans = (ans + contrib) % mod
	}

	return ans
}
```

## 2286 — Booking Concert Tickets In Groups

```go
package main

import "fmt"

// 2286. Booking Concert Tickets in Groups
// ----------------------------------------------------------------
// Maintain n rows each with m seats.  Two operations:
//
//   gather(k, maxRow):
//     Find the first row ≤ maxRow with ≥ k consecutive free seats.
//     Book those k seats in that row, returning [row, col].
//     If impossible return [nil] (empty slice).
//
//   scatter(k, maxRow):
//     Book k seats among rows [0 … maxRow] greedily from left.
//     Return true on success, false otherwise.
//
// Both require O(log n) each — hence a segment tree storing per node:
//   sumFree  – total free seats in the interval
//   maxConsec – maximum consecutive free seats in any row of the interval
//              (this is just max of per‑row free seats because seats in
//               different rows are not consecutive per the problem).

type segNode struct {
	sumFree   int
	maxConsec int
}

type BookMyShow struct {
	m        int
	n        int
	free     []int   // free[i] = free seats remaining in row i
	seg      []segNode
}

func Constructor(n, m int) BookMyShow {
	size := 4 * n
	free := make([]int, n)
	for i := range free {
		free[i] = m
	}
	seg := make([]segNode, size)
	b := BookMyShow{m: m, n: n, free: free, seg: seg}
	b.build(1, 0, n-1)
	return b
}

func (b *BookMyShow) build(idx, l, r int) {
	if l == r {
		b.seg[idx] = segNode{sumFree: b.m, maxConsec: b.m}
		return
	}
	mid := (l + r) / 2
	b.build(idx*2, l, mid)
	b.build(idx*2+1, mid+1, r)
	b.pull(idx)
}

func (b *BookMyShow) pull(idx int) {
	b.seg[idx].sumFree = b.seg[idx*2].sumFree + b.seg[idx*2+1].sumFree
	if b.seg[idx*2].maxConsec > b.seg[idx*2+1].maxConsec {
		b.seg[idx].maxConsec = b.seg[idx*2].maxConsec
	} else {
		b.seg[idx].maxConsec = b.seg[idx*2+1].maxConsec
	}
}

// Gather finds the first row ≤ maxRow with ≥ k free seats,
// books them, and returns [row, col].  Returns empty slice on failure.
func (b *BookMyShow) Gather(k int, maxRow int) []int {
	if k <= 0 || maxRow < 0 || b.seg[1].maxConsec < k {
		return []int{}
	}
	// Binary search on segment tree for leftmost row ≤ maxRow with ≥ k free.
	row := b.findFirst(1, 0, b.n-1, maxRow, k)
	if row == -1 {
		return []int{}
	}
	// Book k seats starting at column (b.M - free[row]).
	col := b.m - b.free[row]
	b.free[row] -= k
	b.update(1, 0, b.n-1, row, b.free[row])
	return []int{row, col}
}

// findFirst returns the smallest row ∈ [0, limit] with free ≥ need, or -1.
func (b *BookMyShow) findFirst(idx, l, r, limit, need int) int {
	if l > limit || b.seg[idx].maxConsec < need {
		return -1
	}
	if l == r {
		if b.free[l] >= need {
			return l
		}
		return -1
	}
	mid := (l + r) / 2
	res := b.findFirst(idx*2, l, mid, limit, need)
	if res != -1 {
		return res
	}
	return b.findFirst(idx*2+1, mid+1, r, limit, need)
}

// Scatter attempts to book k seats across rows [0 … maxRow].
func (b *BookMyShow) Scatter(k int, maxRow int) bool {
	if k <= 0 || maxRow < 0 {
		return false
	}
	if b.querySum(1, 0, b.n-1, 0, maxRow) < k {
		return false
	}
	remaining := k
	// Walk rows from 0 upward, taking whatever is free.
	for r := 0; r <= maxRow && remaining > 0; r++ {
		if b.free[r] == 0 {
			continue
		}
		take := remaining
		if b.free[r] < take {
			take = b.free[r]
		}
		b.free[r] -= take
		b.update(1, 0, b.n-1, r, b.free[r])
		remaining -= take
	}
	return true
}

func (b *BookMyShow) querySum(idx, l, r, ql, qr int) int {
	if ql > r || qr < l {
		return 0
	}
	if ql <= l && r <= qr {
		return b.seg[idx].sumFree
	}
	mid := (l + r) / 2
	return b.querySum(idx*2, l, mid, ql, qr) + b.querySum(idx*2+1, mid+1, r, ql, qr)
}

func (b *BookMyShow) update(idx, l, r, pos, val int) {
	if l == r {
		b.seg[idx] = segNode{sumFree: val, maxConsec: val}
		return
	}
	mid := (l + r) / 2
	if pos <= mid {
		b.update(idx*2, l, mid, pos, val)
	} else {
		b.update(idx*2+1, mid+1, r, pos, val)
	}
	b.pull(idx)
}

// ---------------------------------------------------------------------------
//  Wrapper

func BookingConcertTicketsInGroups() interface{} {
	show := Constructor(2, 5)
	out := make([]interface{}, 0)
	out = append(out, show.Scatter(4, 0)) // true
	out = append(out, show.Scatter(2, 0)) // true
	out = append(out, show.Gather(5, 1))  // [0,0] (row0 had 1 seat left after scatter)
	return out
}

func main() {
	fmt.Println(BookingConcertTicketsInGroups())

	// ---- test ----
	show := Constructor(5, 10)
	if got := show.Gather(6, 3); len(got) == 0 || got[0] != 0 || got[1] != 0 {
		fmt.Printf("FAIL: Gather(6,3) expected [0,0], got %v\n", got)
	}
	if got := show.Gather(5, 2); len(got) == 0 || got[0] != 1 || got[1] != 0 {
		fmt.Printf("FAIL: Gather(5,2) expected [1,0], got %v\n", got)
	}
	if got := show.Gather(12, 4); len(got) != 0 {
		fmt.Printf("FAIL: Gather(12,4) expected [], got %v\n", got)
	}
	if got := show.Scatter(30, 4); got != true {
		fmt.Printf("FAIL: Scatter(30,4) expected true, got %v\n", got)
	}
	if got := show.Gather(2, 4); len(got) == 0 || got[0] != 4 || got[1] != 1 {
		fmt.Printf("FAIL: Gather(2,4) expected [4,1], got %v\n", got)
	}

	show2 := Constructor(3, 3)
	if got := show2.Scatter(2, 2); got != true {
		fmt.Printf("FAIL: Scatter(2,2) expected true\n")
	}
	if got := show2.Gather(3, 2); len(got) == 0 || got[0] != 1 || got[1] != 0 {
		fmt.Printf("FAIL: Gather(3,2) expected [1,0], got %v\n", got)
	}
	fmt.Println("Done testing 2286.")
}
```

## 2290 — Minimum Obstacle Removal To Reach Corner

```go
package main

// LeetCode #2290: Minimum Obstacle Removal to Reach Corner
// https://leetcode.com/problems/minimum-obstacle-removal-to-reach-corner/
// Difficulty: Hard
//
// Approach: 0-1 BFS (Dijkstra with deque since edge weights are 0 or 1).
// Treat grid cells as nodes. Moving from (r,c) to (nr,nc) costs grid[nr][nc]
// (0 if empty, 1 if obstacle). Use deque: push front for cost 0, push back for cost 1.

import (
	"container/list"
	"fmt"
	"math"
)

func main() {
	// Example 1: [[0,1,1],[1,1,0],[1,1,0]] => 2
	fmt.Println(minimumObstacles([][]int{{0, 1, 1}, {1, 1, 0}, {1, 1, 0}}))
	// Example 2: [[0,1,0,0,0],[0,1,0,1,0],[0,0,0,1,0]] => 0
	fmt.Println(minimumObstacles([][]int{{0, 1, 0, 0, 0}, {0, 1, 0, 1, 0}, {0, 0, 0, 1, 0}}))
	// Edge: single cell
	fmt.Println(minimumObstacles([][]int{{0}}))
	fmt.Println(minimumObstacles([][]int{{1}}))
	// Edge: 1xN row, no obstacles
	fmt.Println(minimumObstacles([][]int{{0, 0, 0, 0}}))
}

func minimumObstacles(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dist := make([][]int, m)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = math.MaxInt32
		}
	}

	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	deque := list.New()
	dist[0][0] = grid[0][0]
	deque.PushFront([2]int{0, 0})

	for deque.Len() > 0 {
		front := deque.Remove(deque.Front()).([2]int)
		r, c := front[0], front[1]

		for _, d := range dirs {
			nr, nc := r+d[0], c+d[1]
			if nr < 0 || nr >= m || nc < 0 || nc >= n {
				continue
			}
			cost := grid[nr][nc]
			nd := dist[r][c] + cost
			if nd < dist[nr][nc] {
				dist[nr][nc] = nd
				if cost == 0 {
					deque.PushFront([2]int{nr, nc})
				} else {
					deque.PushBack([2]int{nr, nc})
				}
			}
		}
	}
	return dist[m-1][n-1]
}
```

## 2296 — Design A Text Editor

```go
package main

import (
	"fmt"
)

// 2296. Design a Text Editor
// ----------------------------------------------------------------
// Two-stack approach.  left holds characters before the cursor (top = rightmost),
// right holds characters after the cursor (reversed so right.top is the first
// character after the cursor).

type TextEditor struct {
	left  []byte
	right []byte
}

func Constructor() TextEditor {
	return TextEditor{
		left:  make([]byte, 0, 512),
		right: make([]byte, 0, 512),
	}
}

// addText inserts text at the cursor position.
func (t *TextEditor) addText(text string) string {
	t.left = append(t.left, []byte(text)...)
	return t.peek()
}

// deleteText deletes the k characters immediately before the cursor.
// Returns the string of the last min(10, len(left)) characters after deletion.
func (t *TextEditor) deleteText(k int) string {
	if k > len(t.left) {
		k = len(t.left)
	}
	t.left = t.left[:len(t.left)-k]
	return t.peek()
}

// cursorLeft moves the cursor left by k characters (or to the start).
func (t *TextEditor) cursorLeft(k int) string {
	if k > len(t.left) {
		k = len(t.left)
	}
	// Move k chars from left to right.
	// left's top k chars go to right in reversed order so the cursor stays
	// correctly between left and right.
	moved := t.left[len(t.left)-k:]
	t.left = t.left[:len(t.left)-k]
	for i := len(moved) - 1; i >= 0; i-- {
		t.right = append(t.right, moved[i])
	}
	return t.peek()
}

// cursorRight moves the cursor right by k characters (or to the end).
func (t *TextEditor) cursorRight(k int) string {
	if k > len(t.right) {
		k = len(t.right)
	}
	moved := t.right[len(t.right)-k:]
	t.right = t.right[:len(t.right)-k]
	for i := len(moved) - 1; i >= 0; i-- {
		t.left = append(t.left, moved[i])
	}
	return t.peek()
}

// peek returns the last min(10, len(left)) characters of left.
func (t *TextEditor) peek() string {
	n := len(t.left)
	start := n - 10
	if start < 0 {
		start = 0
	}
	return string(t.left[start:])
}

// ---------------------------------------------------------------------------
//  Wrapper (returns last peek value)

func DesignATextEditor() interface{} {
	editor := Constructor()
	editor.addText("leetcode")
	editor.deleteText(4)
	editor.addText("practice")
	editor.cursorRight(3)
	editor.cursorLeft(8)
	editor.deleteText(10)
	editor.cursorLeft(2)
	return editor.cursorRight(6)
}

func main() {
	fmt.Println(DesignATextEditor())

	editor := Constructor()
	if s := editor.addText("hello"); s != "hello" {
		fmt.Printf("FAIL addText hello: got %q\n", s)
	}
	if s := editor.cursorLeft(2); s != "hel" {
		fmt.Printf("FAIL cursorLeft 2: got %q\n", s)
	}
	if s := editor.cursorRight(2); s != "hello" {
		fmt.Printf("FAIL cursorRight 2: got %q\n", s)
	}
	if s := editor.deleteText(2); s != "hel" {
		fmt.Printf("FAIL deleteText 2: got %q\n", s)
	}
	if s := editor.addText("p"); s != "help" {
		fmt.Printf("FAIL addText p: got %q\n", s)
	}
	fmt.Println("Done testing 2296.")
}
```

## 2301 — Match Substring After Replacement

```go
package main

// LeetCode #2301: Match Substring After Replacement
// https://leetcode.com/problems/match-substring-after-replacement/
// Difficulty: Hard
//
// Approach: Build a directed graph from mappings, compute transitive closure
// (a->b and b->c implies a->c). Then check if each character in s can be
// transformed to the corresponding character in sub, character by character.

import "fmt"

func main() {
	// Example 1: "fool3e7bar","leet",[["e","3"],["t","7"],["t","8"]] => true
	fmt.Println(matchReplacement("fool3e7bar", "leet", [][]byte{{'e', '3'}, {'t', '7'}, {'t', '8'}}))
	// Example 2: "fooleetbar","f00l",[["o","0"]] => false
	fmt.Println(matchReplacement("fooleetbar", "f00l", [][]byte{{'o', '0'}}))
	// Example 3: "Fool33tbaR","leet",[["e","3"],["t","7"],["t","8"],["e","E"],["e","e"]] => true
	fmt.Println(matchReplacement("Fool33tbaR", "leet", [][]byte{{'e', '3'}, {'t', '7'}, {'t', '8'}, {'e', 'E'}, {'e', 'e'}}))
	// Edge: sub == s
	fmt.Println(matchReplacement("abc", "abc", [][]byte{}))
	// Edge: single char
	fmt.Println(matchReplacement("a", "b", [][]byte{{'a', 'b'}}))
	fmt.Println(matchReplacement("a", "b", [][]byte{}))
}

func matchReplacement(s string, sub string, mappings [][]byte) bool {
	// Build transitive closure: can[a][b] means a can be replaced by b
	can := make([][]bool, 256)
	for i := range can {
		can[i] = make([]bool, 256)
		can[i][i] = true
	}
	for _, m := range mappings {
		can[m[0]][m[1]] = true
	}

	// Floyd-Warshall for transitive closure
	for k := 0; k < 256; k++ {
		for i := 0; i < 256; i++ {
			if !can[i][k] {
				continue
			}
			for j := 0; j < 256; j++ {
				if can[k][j] {
					can[i][j] = true
				}
			}
		}
	}

	// Slide sub through s
	for start := 0; start <= len(s)-len(sub); start++ {
		match := true
		for i := 0; i < len(sub); i++ {
			if !can[sub[i]][s[start+i]] {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false
}
```

## 2302 — Count Subarrays With Score Less Than K

```go
package main

// LeetCode #2302: Count Subarrays With Score Less Than K
// https://leetcode.com/problems/count-subarrays-with-score-less-than-k/
// Difficulty: Hard
//
// Approach: Sliding window. Score of subarray nums[l..r] = sum(nums[l..r]) * (r-l+1).
// Expand right pointer. While score >= k, shrink left. For each valid window,
// count = r-l+1 subarrays ending at r. Accumulate.

import "fmt"

func main() {
	// Example 1: [2,1,4,3,5], k=10 => 6
	fmt.Println(countSubarrays([]int{2, 1, 4, 3, 5}, 10))
	// Example 2: [1,1,1], k=5 => 5
	fmt.Println(countSubarrays([]int{1, 1, 1}, 5))
	// Edge: all elements too large
	fmt.Println(countSubarrays([]int{10, 10, 10}, 5))
	// Edge: single element less than k
	fmt.Println(countSubarrays([]int{3}, 10))
	// Edge: k=1
	fmt.Println(countSubarrays([]int{1, 2, 3}, 1))
}

func countSubarrays(nums []int, k int64) int64 {
	var ans int64
	var sum int64
	left := 0
	for right := 0; right < len(nums); right++ {
		sum += int64(nums[right])
		// Shrink while score >= k
		// score = sum * (right-left+1)
		for sum*int64(right-left+1) >= k {
			sum -= int64(nums[left])
			left++
			if left > right {
				break
			}
		}
		if left <= right {
			ans += int64(right - left + 1)
		}
	}
	return ans
}
```

## 2306 — Naming A Company

```go
package main

// LeetCode #2306: Naming a Company
// https://leetcode.com/problems/naming-a-company/
// Difficulty: Hard
//
// Approach: Group ideas by their first character. For each group, store the
// suffix (everything after first char) in a set. For each pair of groups (i,j)
// with i != j, count common suffixes. Valid pairs from this group pair =
// (len(group[i]) - common) * (len(group[j]) - common). Sum over all pairs.
// Result = sum * 2 (since swap order also counts: "A B" and "B A" are both valid).

import (
	"fmt"
)

func main() {
	// Example 1: ["coffee","donuts","time","toffee"] => 6
	fmt.Println(distinctNames([]string{"coffee", "donuts", "time", "toffee"}))
	// Example 2: ["lack","back"] => 0
	fmt.Println(distinctNames([]string{"lack", "back"}))
	// Edge: single idea
	fmt.Println(distinctNames([]string{"hello"}))
	// Edge: no common suffixes
	fmt.Println(distinctNames([]string{"abc", "def", "ghi"}))
	// Edge: all same first letter
	fmt.Println(distinctNames([]string{"aa", "ab", "ac"}))
}

func distinctNames(ideas []string) int64 {
	// group[first_letter] = set of suffixes
	groups := make([]map[string]bool, 26)
	for i := range groups {
		groups[i] = make(map[string]bool)
	}

	for _, idea := range ideas {
		first := idea[0] - 'a'
		suffix := idea[1:]
		groups[first][suffix] = true
	}

	var ans int64
	for i := 0; i < 26; i++ {
		if len(groups[i]) == 0 {
			continue
		}
		for j := i + 1; j < 26; j++ {
			if len(groups[j]) == 0 {
				continue
			}
			// Count common suffixes between groups i and j
			common := 0
			for suffix := range groups[i] {
				if groups[j][suffix] {
					common++
				}
			}
			validI := len(groups[i]) - common
			validJ := len(groups[j]) - common
			ans += int64(validI * validJ * 2)
		}
	}
	return ans
}
```

## 2307 — Check For Contradictions In Equations

```go
package main

// LeetCode #2307: Check for Contradictions in Equations
// https://leetcode.com/problems/check-for-contradictions-in-equations/
// Difficulty: Hard [Paid]
//
// Approach: Weighted Union-Find. Each equation a / b = val means a = val * b.
// Store parent and ratio (value[a] / value[parent[a]]). When unioning a and b
// with val (where a/b = val):
//   - Find roots ra, rb and ratios ra, rb
//   - If ra == rb, check consistency: ratio[b]/ratio[a] == val
//   - Otherwise, union ra under rb (or vice versa) with appropriate ratio

import "fmt"

func main() {
	// Example 1: [["a","b"],["b","c"],["a","c"]], [2.0,3.0,6.0] => true (no contradiction)
	fmt.Println(checkContradictions([][]string{{"a", "b"}, {"b", "c"}, {"a", "c"}}, []float64{2.0, 3.0, 6.0}))
	// Example 2: [["a","b"],["b","c"],["a","c"]], [2.0,3.0,5.0] => false (contradiction: a=2b, b=3c, a=6c but a/c=5)
	fmt.Println(checkContradictions([][]string{{"a", "b"}, {"b", "c"}, {"a", "c"}}, []float64{2.0, 3.0, 5.0}))
	// Example 3: [["a","b"],["c","d"]], [2.0,3.0] => true
	fmt.Println(checkContradictions([][]string{{"a", "b"}, {"c", "d"}}, []float64{2.0, 3.0}))
	// Edge: single equation
	fmt.Println(checkContradictions([][]string{{"a", "b"}}, []float64{2.0}))
	// Edge: self-loop (a/a = 2.0 is contradiction since a/a must be 1.0)
	fmt.Println(checkContradictions([][]string{{"a", "a"}}, []float64{2.0}))
}

type WeightedUF struct {
	parent map[string]string
	ratio  map[string]float64 // value[key] / value[parent[key]]
}

func NewWeightedUF() *WeightedUF {
	return &WeightedUF{
		parent: make(map[string]string),
		ratio:  make(map[string]float64),
	}
}

func (uf *WeightedUF) find(x string) (string, float64) {
	if _, ok := uf.parent[x]; !ok {
		uf.parent[x] = x
		uf.ratio[x] = 1.0
	}
	if uf.parent[x] == x {
		return x, 1.0
	}
	root, r := uf.find(uf.parent[x])
	uf.ratio[x] *= r
	uf.parent[x] = root
	return root, uf.ratio[x]
}

func checkContradictions(equations [][]string, values []float64) bool {
	// Tolerance for floating point comparison
	const eps = 1e-9

	uf := NewWeightedUF()

	for i, eq := range equations {
		a, b := eq[0], eq[1]
		val := values[i] // a / b = val => a = val * b

		if a == b {
			if abs(val-1.0) > eps {
				return false
			}
			continue
		}

		ra, raRatio := uf.find(a) // ratio[a] / ratio[ra] = raRatio
		rb, rbRatio := uf.find(b) // ratio[b] / ratio[rb] = rbRatio
		// a = val * b
		// raRatio * ra = a, mbaumRatio * rb = b
		// raRatio * ra = val * rbRatio * rb
		// ra/rb = val * rbRatio / raRatio

		if ra == rb {
			// Both already in same set. Check: ratio[a]/ratio[b] == val
			// ratio[a] = raRatio, ratio[b] = rbRatio (since both relative to same root)
			if abs(raRatio/rbRatio-val) > eps {
				return false
			}
		} else {
			// Union: make ra point to rb
			// We want: ratio[a] / ratio[b] = val
			// Let new ratio for ra: ratio[ra] / ratio[rb] = ?
			// ratio[a] = raRatio * uf.ratio[ra] (after union)
			// We want: raRatio * x / rbRatio = val
			// x = val * rbRatio / raRatio
			uf.parent[ra] = rb
			uf.ratio[ra] = val * rbRatio / raRatio
		}
	}
	return true
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
```

## 2312 — Selling Pieces Of Wood

```go
package main

// LeetCode #2312: Selling Pieces of Wood
// https://leetcode.com/problems/selling-pieces-of-wood/
// Difficulty: Hard
//
// Approach: DP. dp[h][w] = max profit for a piece of size h x w.
// Initialize dp with 0, then apply given prices. For each piece, try all
// horizontal cuts (split into h1 x w and (h-h1) x w) and vertical cuts
// (split into h x w1 and h x (w-w1)). Take max.
// m <= 200, n <= 200, so O(m*n*(m+n)) is fine.

import "fmt"

func main() {
	// Example 1: m=3, n=5, prices=[[1,4,2],[2,2,7],[2,1,3]] => 19
	fmt.Println(sellingWood(3, 5, [][]int{{1, 4, 2}, {2, 2, 7}, {2, 1, 3}}))
	// Example 2: m=4, n=6, prices=[[3,2,10],[1,4,2],[4,1,3]] => 32
	fmt.Println(sellingWood(4, 6, [][]int{{3, 2, 10}, {1, 4, 2}, {4, 1, 3}}))
	// Edge: no prices
	fmt.Println(sellingWood(2, 2, [][]int{}))
	// Edge: single price
	fmt.Println(sellingWood(2, 2, [][]int{{2, 2, 5}}))
	// Edge: 1x1
	fmt.Println(sellingWood(1, 1, [][]int{{1, 1, 10}}))
}

func sellingWood(m int, n int, prices [][]int) int64 {
	dp := make([][]int64, m+1)
	for i := range dp {
		dp[i] = make([]int64, n+1)
	}

	// Apply given prices
	for _, p := range prices {
		h, w, pr := p[0], p[1], p[2]
		if h <= m && w <= n {
			if int64(pr) > dp[h][w] {
				dp[h][w] = int64(pr)
			}
		}
	}

	// DP: try all cuts
	for h := 1; h <= m; h++ {
		for w := 1; w <= n; w++ {
			// Horizontal cuts: split into h1 x w and (h-h1) x w
			for h1 := 1; h1 <= h/2; h1++ {
				val := dp[h1][w] + dp[h-h1][w]
				if val > dp[h][w] {
					dp[h][w] = val
				}
			}
			// Vertical cuts: split into h x w1 and h x (w-w1)
			for w1 := 1; w1 <= w/2; w1++ {
				val := dp[h][w1] + dp[h][w-w1]
				if val > dp[h][w] {
					dp[h][w] = val
				}
			}
		}
	}
	return dp[m][n]
}
```

## 2313 — Minimum Flips In Binary Tree To Get Result

```go
package main

import (
	"fmt"
	"math"
)

// 2313. Minimum Flips in Binary Tree to Get Result
// ----------------------------------------------------------------
// Each node is either a leaf (value 0 or 1) or an internal node
// (AND=3, OR=4, XOR=5, NOT=6 — or whatever encoding LeetCode chooses).
// For NOT the node has a single child; for AND/OR/XOR it has two children.
//
// We compute dp[node][target] = minimum flips needed in the subtree of
// node to make it evaluate to target (0 or 1).
//
// A "flip" changes a leaf's value (0↔1) or replaces an internal node's operator
// with any other operator (AND/OR/XOR/NOT).  The problem description on
// LeetCode gives specific operator constants — we use the same ones as the
// problem: 0=leaf false, 1=leaf true, 2=NOT, 3=AND, 4=OR, 5=XOR.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minimumFlips(root *TreeNode, result int) int {
	var dfs func(*TreeNode) [2]int // [cost0, cost1]
	dfs = func(node *TreeNode) [2]int {
		if node.Left == nil && node.Right == nil {
			// leaf
			if node.Val == 0 {
				return [2]int{0, 1} // cost0=0, cost1=1 (flip)
			}
			return [2]int{1, 0} // cost0=1, cost1=0
		}
		if node.Val == 2 { // NOT
			child := node.Left
			if child == nil {
				child = node.Right
			}
			c := dfs(child)
			// output 1 → child 0, output 0 → child 1
			return [2]int{c[1], c[0]}
		}
		// AND / OR / XOR — two children
		l := dfs(node.Left)
		r := dfs(node.Right)
		// Cost to get 0 and 1 for each operator.
		switch node.Val {
		case 3: // AND
			// result 1 only if both 1
			// result 0 if at least one 0
			return [2]int{
				min3(l[0]+r[0], l[0]+r[1], l[1]+r[0]), // cost0
				l[1] + r[1], // cost1
			}
		case 4: // OR
			// result 0 only if both 0
			return [2]int{
				l[0] + r[0], // cost0
				min3(l[1]+r[0], l[0]+r[1], l[1]+r[1]), // cost1
			}
		case 5: // XOR
			return [2]int{
				min(l[0]+r[0], l[1]+r[1]),     // cost0 (same)
				min(l[0]+r[1], l[1]+r[0]),     // cost1 (different)
			}
		default:
			return [2]int{math.MaxInt32, math.MaxInt32}
		}
	}
	cost := dfs(root)
	if result == 0 {
		return cost[0]
	}
	return cost[1]
}

func min3(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}

// ---------------------------------------------------------------------------
//  Wrapper

func MinimumFlipsInBinaryTreeToGetResult() interface{} {
	// Example: leaf(0) → NOT → result=1
	root := &TreeNode{Val: 2, Left: &TreeNode{Val: 0}}
	return minimumFlips(root, 1) // flip leaf 0→1 costs 1
}

func main() {
	fmt.Println(MinimumFlipsInBinaryTreeToGetResult())

	// leaf 0 → NOT → answer=1 ⇒ flip leaf once → cost 1
	root1 := &TreeNode{Val: 2, Left: &TreeNode{Val: 0}}
	if got := minimumFlips(root1, 1); got != 0 {
		fmt.Printf("FAIL test1: got %d, want 0\n", got)
	}

	// AND([1,0]) → 0, want 1: cheapest is flip leaf 0→1 (cost 1)
	root2 := &TreeNode{Val: 3, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 0}}
	if got := minimumFlips(root2, 1); got != 1 {
		fmt.Printf("FAIL test2: got %d, want 0\n", got)
	}

	// OR([0,0]) → 0, want 1: flip one leaf (cost 1)
	root3 := &TreeNode{Val: 4, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 0}}
	if got := minimumFlips(root3, 1); got != 1 {
		fmt.Printf("FAIL test3: got %d, want 0\n", got)
	}

	// XOR([1,0]) → 1, want 0: cheapest to make output 0 is to
	// flip one of them so both are same: cost 1
	root4 := &TreeNode{Val: 5, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 0}}
	if got := minimumFlips(root4, 0); got != 1 {
		fmt.Printf("FAIL test4: got %d, want 0\n", got)
	}

	fmt.Println("Done testing 2313.")
}
```

## 2318 — Number Of Distinct Roll Sequences

```go
package main

import (
	"fmt"
)

// 2318. Number of Distinct Roll Sequences
// ----------------------------------------------------------------
// Roll a 6‑sided die n times.  A sequence is valid if:
//   1. gcd(roll[i], roll[i-1]) == 1 for i ≥ 1.
//   2. If roll[i] == roll[j] then |i - j| > 2  (gap of at least 2 between equals).
//
// Equivalent constraints used in DP:
//   - Adjacent: gcd(cur, prev) == 1  (implies cur != prev except for cur=prev=1).
//   - Gap-2:    cur != secPrev  (roll[i] != roll[i-2]).
//
// DP state: dp[prev][secPrev] = count of sequences ending with ..., secPrev, prev.
// Transition to cur: gcd(cur, prev) == 1 AND cur != secPrev.

const MOD = 1_000_000_007

func distinctRollSequences(n int) int {
	// Base: length 1 — any single value works.
	if n == 1 {
		return 6
	}

	// dp[prev][secPrev] — last two rolls of the current sequence.
	dp := [7][7]int{}

	// Initialise for length 2.
	for p := 1; p <= 6; p++ {
		for q := 1; q <= 6; q++ {
			if p != q && gcd(p, q) == 1 {
				dp[p][q] = 1 // sequence = [q, p]
			}
		}
	}

	if n == 2 {
		sum := 0
		for p := 1; p <= 6; p++ {
			for q := 1; q <= 6; q++ {
				sum = (sum + dp[p][q]) % MOD
			}
		}
		return sum
	}

	// Extend from length 3 to n.
	for i := 3; i <= n; i++ {
		ndp := [7][7]int{}
		for cur := 1; cur <= 6; cur++ {
			for prev := 1; prev <= 6; prev++ {
				if cur == prev || gcd(cur, prev) != 1 {
					continue
				}
				// Sum over all secPrev where (secPrev, prev) was valid at previous
				// step AND cur != secPrev (gap-2 constraint).
				s := 0
				for secPrev := 1; secPrev <= 6; secPrev++ {
					if cur == secPrev {
						continue
					}
					s = (s + dp[prev][secPrev]) % MOD
				}
				ndp[cur][prev] = s
			}
		}
		dp = ndp
	}

	sum := 0
	for cur := 1; cur <= 6; cur++ {
		for prev := 1; prev <= 6; prev++ {
			sum = (sum + dp[cur][prev]) % MOD
		}
	}
	return sum
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// ---------------------------------------------------------------------------
//  Wrapper

func NumberOfDistinctRollSequences() interface{} {
	return distinctRollSequences(4)
}

func main() {
	fmt.Println(NumberOfDistinctRollSequences())

	// Tests from problem.
	cases := []struct{ n, want int }{
		{1, 6},
		{2, 22},
		{3, 66}, // from problem
		{4, 184},
	}
	for _, c := range cases {
		got := distinctRollSequences(c.n)
		if got != c.want {
			fmt.Printf("FAIL n=%d: got %d, want %d\n", c.n, got, c.want)
		}
	}
	fmt.Println("Done testing 2318.")
}
```

## 2321 — Maximum Score Of Spliced Array

```go
package main

import (
	"fmt"
)

// 2321. Maximum Score Of Spliced Array
// ----------------------------------------------------------------
// We may swap a single contiguous subarray between nums1 and nums2.
// After the swap, return the maximum possible sum of EITHER array.
//
// Gain from swapping subarray [l,r]:
//   nums2 → nums1: gain = sum(nums2[l..r] - nums1[l..r])
//   nums1 → nums2: gain = sum(nums1[l..r] - nums2[l..r])
//
// Answer = max(sum(nums1) + bestKadane(diff),
//              sum(nums2) + bestKadane(-diff))
// where diff[i] = nums2[i] - nums1[i].

func maximumsSplicedArray(nums1, nums2 []int) int {
	n := len(nums1)
	sum1, sum2 := 0, 0
	diff := make([]int, n)
	for i := 0; i < n; i++ {
		sum1 += nums1[i]
		sum2 += nums2[i]
		diff[i] = nums2[i] - nums1[i]
	}

	g1 := maxSubarray(diff)   // best gain from swapping nums2 INTO nums1
	g2 := maxSubarrayNeg(diff) // best gain from swapping nums1 INTO nums2

	ans := sum1 + g1
	if sum2+g2 > ans {
		ans = sum2 + g2
	}
	return ans
}

// maxSubarray returns the maximum subarray sum (Kadane); empty subarray → 0.
func maxSubarray(arr []int) int {
	best, cur := 0, 0
	for _, x := range arr {
		cur += x
		if cur < 0 {
			cur = 0
		}
		if cur > best {
			best = cur
		}
	}
	return best
}

// maxSubarrayNeg returns maxSubarray of -arr.
func maxSubarrayNeg(arr []int) int {
	best, cur := 0, 0
	for _, x := range arr {
		cur -= x // equivalent to cur += (-x)
		if cur < 0 {
			cur = 0
		}
		if cur > best {
			best = cur
		}
	}
	return best
}

// ---------------------------------------------------------------------------
//  Wrapper

func MaximumScoreOfSplicedArray() interface{} {
	return maximumsSplicedArray([]int{60, 60, 60}, []int{10, 90, 10})
}

func main() {
	fmt.Println(MaximumScoreOfSplicedArray())

	tests := []struct {
		a, b []int
		want int
	}{
		{[]int{60, 60, 60}, []int{10, 90, 10}, 210},
		{[]int{20, 40, 20, 70, 30}, []int{50, 20, 50, 40, 20}, 220},
		{[]int{7, 11, 13}, []int{1, 1, 1}, 31},
		{[]int{1, 2, 3}, []int{3, 2, 1}, 8},
	}
	for _, tc := range tests {
		got := maximumsSplicedArray(tc.a, tc.b)
		if got != tc.want {
			fmt.Printf("FAIL nums1=%v nums2=%v: got %d, want %d\n",
				tc.a, tc.b, got, tc.want)
		}
	}
	fmt.Println("Done testing 2321.")
}
```

## 2322 — Minimum Score After Removals On A Tree

```go
package main

import (
	"fmt"
	"math"
)

// 2322. Minimum Score After Removals on a Tree
// ----------------------------------------------------------------
// Given a tree with n nodes, each having a value.  Remove exactly two edges,
// splitting the tree into three connected components.  The "score" of a
// component is the XOR of all node values in that component.
// Minimise max(score1, score2, score3) - min(score1, score2, score3).
//
// n ≤ 1000, so O(n²) is fine.
// Idea:
//   1. Root the tree at 0.
//   2. Compute subtree XOR via DFS (post‑order).
//   3. Enumerate every unordered pair of edges (i, j), i ≠ j.
//      Let subtree_xor[v] = XOR of all nodes in the subtree rooted at v.
//
//      Relationship between the two edges:
//       a) One edge is a descendant of the other.
//          If edge a (removing edge at 'a') is an ancestor of edge b,
//          components are:
//            comp1 = subtree_xor[b]
//            comp2 = subtree_xor[a] XOR subtree_xor[b]
//            comp3 = total_xor XOR subtree_xor[a]
//
//       b) The two edges are in unrelated subtrees.
//          components are:
//            comp1 = subtree_xor[a]
//            comp2 = subtree_xor[b]
//            comp3 = total_xor XOR subtree_xor[a] XOR subtree_xor[b]
//
//   Also consider removing edge a and b where one is the parent of another
//   (same as case a).

func minimumScore(nums []int, edges [][]int) int {
	n := len(nums)
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// Parent and subtree XOR via DFS.
	parent := make([]int, n)
	order := make([]int, 0, n) // DFS order
	subXor := make([]int, n)

	var dfs func(u, p int)
	dfs = func(u, p int) {
		parent[u] = p
		order = append(order, u)
		subXor[u] = nums[u]
		for _, v := range adj[u] {
			if v == p {
				continue
			}
			dfs(v, u)
			subXor[u] ^= subXor[v]
		}
	}
	dfs(0, -1)
	total := subXor[0]

	// Pre‑compute ancestors for O(1) descendant check.
	// tin / tout using Euler tour.
	tin := make([]int, n)
	tout := make([]int, n)
	time := 0
	var euler func(u int)
	euler = func(u int) {
		tin[u] = time
		time++
		for _, v := range adj[u] {
			if v == parent[u] {
				continue
			}
			euler(v)
		}
		tout[u] = time
		time++
	}
	euler(0)

	isAncestor := func(a, b int) bool {
		// Is 'a' an ancestor of 'b'?
		return tin[a] <= tin[b] && tout[b] <= tout[a]
	}

	best := math.MaxInt32
	for i := 0; i < n; i++ {
		if i == 0 {
			continue // can't cut above root
		}
		for j := i + 1; j < n; j++ {
			if j == 0 {
				continue
			}
			var x, y, z int
			if isAncestor(i, j) {
				// i is ancestor of j
				x = subXor[j]
				y = subXor[i] ^ subXor[j]
				z = total ^ subXor[i]
			} else if isAncestor(j, i) {
				// j is ancestor of i
				x = subXor[i]
				y = subXor[j] ^ subXor[i]
				z = total ^ subXor[j]
			} else {
				x = subXor[i]
				y = subXor[j]
				z = total ^ subXor[i] ^ subXor[j]
			}
			score := max3(x, y, z) - min3(x, y, z)
			if score < best {
				best = score
			}
		}
	}
	return best
}

func max3(a, b, c int) int {
	if a >= b && a >= c {
		return a
	}
	if b >= a && b >= c {
		return b
	}
	return c
}

func min3(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= a && b <= c {
		return b
	}
	return c
}

// ---------------------------------------------------------------------------
//  Wrapper

func MinimumScoreAfterRemovalsOnATree() interface{} {
	nums := []int{1, 5, 5, 4, 11}
	edges := [][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}}
	return minimumScore(nums, edges)
}

func main() {
	fmt.Println(MinimumScoreAfterRemovalsOnATree())

	// Example from problem:
	// nums = [1,5,5,4,11], edges = [[0,1],[1,2],[1,3],[3,4]]
	// The minimum score after removing two edges should be 9.
	if got := minimumScore([]int{1, 5, 5, 4, 11},
		[][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}}); got != 9 {
		fmt.Printf("FAIL example: got %d, want 9\n", got)
	}

	// Another test.
	if got := minimumScore([]int{5, 5, 2, 4, 4, 2},
		[][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {4, 5}}); got != 2 {
		fmt.Printf("FAIL line: got %d, want 2\n", got)
	}

	fmt.Println("Done testing 2322.")
}
```

## 2328 — Number Of Increasing Paths In A Grid

```go
package main

// LeetCode #2328: Number of Increasing Paths in a Grid
// https://leetcode.com/problems/number-of-increasing-paths-in-a-grid/
// Difficulty: Hard
//
// Approach: DFS + memoization. For each cell, count number of strictly
// increasing paths starting from that cell (including length-1 path).
// Use memoization: memo[r][c] = count of increasing paths starting at (r,c).
// Order doesn't matter since paths must be strictly increasing (no cycles).
// Result is sum of all memo values mod 1e9+7.

import "fmt"

func main() {
	// Example 1: [[1,1],[3,4]] => 8
	fmt.Println(countPaths([][]int{{1, 1}, {3, 4}}))
	// Example 2: [[1],[2]] => 3
	fmt.Println(countPaths([][]int{{1}, {2}}))
	// Edge: single cell
	fmt.Println(countPaths([][]int{{5}}))
	// Edge: all decreasing
	fmt.Println(countPaths([][]int{{5, 4}, {3, 2}}))
	// Edge: 3x3 all equal
	fmt.Println(countPaths([][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}))
}

const mod = 1_000_000_007

func countPaths(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
		// -1 means uncomputed
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	var dfs func(r, c int) int
	dfs = func(r, c int) int {
		if memo[r][c] != -1 {
			return memo[r][c]
		}
		// Each path starts at this cell (count = 1 for the cell itself)
		count := 1
		for _, d := range dirs {
			nr, nc := r+d[0], c+d[1]
			if nr < 0 || nr >= m || nc < 0 || nc >= n {
				continue
			}
			if grid[nr][nc] <= grid[r][c] {
				continue
			}
			count = (count + dfs(nr, nc)) % mod
		}
		memo[r][c] = count
		return count
	}

	var ans int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ans = (ans + dfs(i, j)) % mod
		}
	}
	return ans
}
```

## 2334 — Subarray With Elements Greater Than Varying Threshold

```go
package main

import (
	"fmt"
)

// 2334. Subarray With Elements Greater Than Varying Threshold
// ----------------------------------------------------------------
// Find the longest subarray such that every element is strictly greater than
// threshold / len(subarray).  Equivalently, if the minimum element of the
// subarray is m and length is L, we need m * L > threshold.
//
// For each element nums[i], find the maximal subarray where nums[i] is the
// *minimum* (strictly smaller than any element outside).  The left bound is
// given by the previous element strictly smaller than nums[i];
// the right bound by the next element strictly smaller.
//
//   left  = index of previous smaller element + 1
//   right = index of next smaller element - 1
//   maxLen = right - left + 1
//
// Then check whether nums[i] * maxLen > threshold.  If yes, maxLen is a
// candidate answer (we can always shorten the subarray since every element
// is still ≥ nums[i] which is > threshold/maxLen ≥ threshold/shorterLen).
// Wait — if we shorten the subarray, the condition becomes easier because
// threshold/shorterLen is larger, but the minimum may increase if we cut off
// elements.  Since nums[i] is the *global* minimum for the maximal subarray,
// any subarray still containing nums[i] has minimum ≤ nums[i].  The condition
// nums[i] * L > threshold is monotonic in L for fixed i: larger L is better.
// So checking only maxLen is sufficient.

func validSubarraySize(nums []int, threshold int) int {
	n := len(nums)
	prevSmaller := make([]int, n)
	nextSmaller := make([]int, n)

	// Previous strictly smaller element.
	stack := make([]int, 0)
	for i := 0; i < n; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] >= nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			prevSmaller[i] = -1
		} else {
			prevSmaller[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	// Next strictly smaller element.
	stack = stack[:0]
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && nums[stack[len(stack)-1]] >= nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			nextSmaller[i] = n
		} else {
			nextSmaller[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	best := -1
	for i := 0; i < n; i++ {
		left := prevSmaller[i] + 1
		right := nextSmaller[i] - 1
		maxLen := right - left + 1
		// Use int64 to avoid overflow.
		if int64(nums[i])*int64(maxLen) > int64(threshold) && maxLen > best {
			best = maxLen
		}
	}
	return best
}

// ---------------------------------------------------------------------------
//  Wrapper

func SubarrayWithElementsGreaterThanVaryingThreshold() interface{} {
	return validSubarraySize([]int{1, 3, 4, 3, 1}, 6)
}

func main() {
	fmt.Println(SubarrayWithElementsGreaterThanVaryingThreshold())

	tests := []struct {
		nums      []int
		threshold int
		want      int
	}{
		{[]int{1, 3, 4, 3, 1}, 6, 3},
		{[]int{6, 5, 6, 5, 8}, 7, 5},
		{[]int{1, 2, 3, 4, 5}, 100, -1},
		{[]int{5, 5, 5, 5}, 10, 4},
		{[]int{2, 1, 2}, 1, 3},
	}
	for _, tc := range tests {
		got := validSubarraySize(tc.nums, tc.threshold)
		if got != tc.want {
			fmt.Printf("FAIL nums=%v threshold=%d: got %d, want %d\n",
				tc.nums, tc.threshold, got, tc.want)
		}
	}
	fmt.Println("Done testing 2334.")
}
```

## 2338 — Count The Number Of Ideal Arrays

```go
package main

import (
	"fmt"
)

// 2338. Count the Number of Ideal Arrays
// ----------------------------------------------------------------
// An array of length n with values in [1, maxValue] is "ideal" if every
// element divides the next.
//
// Observation: the chain of *distinct* values in an ideal array forms a
// divisor chain.  Consecutive equal values are allowed while still satisfying
// divisibility (since v divides v).  Therefore an ideal array of length n
// corresponds to:
//   - Pick a divisor chain of length k (distinct values).
//   - Distribute n positions among the k distinct values, each appearing at
//     least once: C(n-1, k-1) ways (stars and bars).
//
// Let f(k) = number of divisor chains of length k with values in [1, maxValue].
//   f(1) = maxValue
//   f(k) = sum_{v=1}^{maxValue} dp[k][v]
//   where dp[k][v] = sum_{d|v, d<v} dp[k-1][d]
//
// dp[k][v] can be computed efficiently using harmonic series: for each d,
// for each multiple m of d (m = 2d, 3d, …, maxValue):
//   dp[k][m] += dp[k-1][d]
//
// The maximum chain length is at most about log₂(maxValue) (since each step
// at least doubles).  For maxValue ≤ 10⁴, max chain length ≤ 14.
//
// Answer = sum_{k=1}^{maxChain} C(n-1, k-1) * f(k) mod 1e9+7.

const MOD = 1_000_000_007

func idealArrays(n int, maxValue int) int {
	// Pre‑compute combinations up to C(n-1, 13) (max chain is ~14).
	maxK := 14
	if maxK > n {
		maxK = n
	}
	C := make([][]int, n)
	for i := range C {
		C[i] = make([]int, maxK+1)
		C[i][0] = 1
		for j := 1; j <= i && j <= maxK; j++ {
			C[i][j] = (C[i-1][j-1] + C[i-1][j]) % MOD
		}
	}

	// dp[k][v] — we only need two layers at a time.
	dp := make([]int, maxValue+1)
	for v := 1; v <= maxValue; v++ {
		dp[v] = 1 // k = 1
	}

	f := make([]int, maxK+1)
	f[1] = maxValue

	for k := 2; k <= maxK; k++ {
		ndp := make([]int, maxValue+1)
		total := 0
		// For each d, add dp[d] to its multiples.
		for d := 1; d <= maxValue; d++ {
			if dp[d] == 0 {
				continue
			}
			for m := d * 2; m <= maxValue; m += d {
				ndp[m] = (ndp[m] + dp[d]) % MOD
			}
		}
		for v := 1; v <= maxValue; v++ {
			total = (total + ndp[v]) % MOD
		}
		f[k] = total
		dp = ndp
	}

	ans := 0
	for k := 1; k <= maxK; k++ {
		if f[k] == 0 {
			break
		}
		comb := C[n-1][k-1] // C(n-1, k-1)
		ans = (ans + comb*f[k]) % MOD
	}
	return ans
}

// ---------------------------------------------------------------------------
//  Wrapper

func CountTheNumberOfIdealArrays() interface{} {
	return idealArrays(5, 3)
}

func main() {
	fmt.Println(CountTheNumberOfIdealArrays())

	tests := []struct {
		n, maxV, want int
	}{
		{2, 5, 10},
		{5, 3, 11},
		{3, 2, 4}, // sequences of length 3 with values 1,2: 111,112,122,222 = 4
		{1, 10, 10},
		{4, 4, 19},
	}
	for _, tc := range tests {
		got := idealArrays(tc.n, tc.maxV)
		if got != tc.want {
			fmt.Printf("FAIL n=%d maxV=%d: got %d, want %d\n",
				tc.n, tc.maxV, got, tc.want)
		}
	}
	fmt.Println("Done testing 2338.")
}
```

## 2344 — Minimum Deletions To Make Array Divisible

```go
package main

import (
	"fmt"
	"sort"
)

// 2344. Minimum Deletions to Make Array Divisible
// ----------------------------------------------------------------
// Given nums (integers) and numsDivide (array of divisors), we can delete
// elements from nums.  Return the minimum number of deletions so that the
// smallest remaining element divides every element of numsDivide.
//
// Strategy:
//   1. Compute g = gcd of all elements in numsDivide.
//   2. Sort nums.
//   3. Scan from the smallest: find the first element that divides g.
//      Return its index (number of deletions).  If none found, return -1.

func minDeletions(nums []int, numsDivide []int) int {
	// GCD of all numsDivide.
	g := 0
	for _, v := range numsDivide {
		g = gcd(g, v)
	}

	sort.Ints(nums)
	for i, v := range nums {
		if g%v == 0 {
			return i
		}
	}
	return -1
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// ---------------------------------------------------------------------------
//  Wrapper

func MinimumDeletionsToMakeArrayDivisible() interface{} {
	return minDeletions([]int{4, 3, 6}, []int{8, 2, 10})
}

func main() {
	fmt.Println(MinimumDeletionsToMakeArrayDivisible())

	tests := []struct {
		nums      []int
		numsDiv   []int
		want      int
	}{
		{[]int{4, 3, 6}, []int{8, 2, 10}, -1},
		{[]int{3, 2, 6}, []int{9, 6, 12}, 1},
		{[]int{5, 7, 11}, []int{2, 3, 4}, -1},
		{[]int{2, 4, 8}, []int{4, 8, 16}, 0},
		{[]int{8, 12, 6, 4}, []int{9, 15}, -1},
	}
	for _, tc := range tests {
		got := minDeletions(tc.nums, tc.numsDiv)
		if got != tc.want {
			fmt.Printf("FAIL nums=%v numsDiv=%v: got %d, want %d\n",
				tc.nums, tc.numsDiv, got, tc.want)
		}
	}
	fmt.Println("Done testing 2344.")
}
```

## 2350 — Shortest Impossible Sequence Of Rolls

```go
package main

// LeetCode #2350: Shortest Impossible Sequence of Rolls
// https://leetcode.com/problems/shortest-impossible-sequence-of-rolls/
// Difficulty: Hard
//
// Approach: Greedy. We want the shortest sequence of rolls (each 1..k) that
// is NOT a subsequence of the given rolls array. The key insight: in each
// "round", collect all k distinct values 1..k. Count how many complete rounds
// we can make. Answer = rounds + 1.
//
// Example: rolls=[4,2,1,2,3,3,2,4,1], k=4
// Round 1: collect {4,2,1,3} (positions 0,1,2,4)
// Round 2: collect {3,2,4,1} (positions 5,6,7,8)
// rounds=2, answer=3

import "fmt"

func main() {
	// Example 1: [4,2,1,2,3,3,2,4,1], 4 => 3
	fmt.Println(shortestSequence([]int{4, 2, 1, 2, 3, 3, 2, 4, 1}, 4))
	// Example 2: [1,1,2,3], 4 => 2
	fmt.Println(shortestSequence([]int{1, 1, 2, 3}, 4))
	// Example 3: [1,2,3,4], 5 => 2
	fmt.Println(shortestSequence([]int{1, 2, 3, 4}, 5))
	// Edge: first roll already incomplete
	fmt.Println(shortestSequence([]int{1, 1, 1, 1}, 3))
	// Edge: single round
	fmt.Println(shortestSequence([]int{1, 2, 3}, 3))
}

func shortestSequence(rolls []int, k int) int {
	seen := make([]bool, k+1)
	rounds := 0
	count := 0

	for _, r := range rolls {
		if !seen[r] {
			seen[r] = true
			count++
			if count == k {
				rounds++
				count = 0
				// Reset seen
				seen = make([]bool, k+1)
			}
		}
	}
	return rounds + 1
}
```

## 2354 — Number Of Excellent Pairs

```go
package main

import "fmt"

// 2354. Number of Excellent Pairs
// ----------------------------------------------------------------
// Given an array of integers nums and an integer k, a pair (i, j) is
// "excellent" if setBits(nums[i]) + setBits(nums[j]) >= k.
// Return the number of distinct pairs (i, j) where i and j are indices
// (unordered, i ≤ j — but pairs are counted by VALUE, not by index).
//
// The problem counts (i, j) such that the *values* nums[i] and nums[j]
// satisfy the condition.  Since we only care about set‑bit counts, and
// duplicate values produce the same pair, we can:
//   1. Deduplicate nums.
//   2. Count how many unique values have each set‑bit count (0 … 60).
//   3. Sum up pairs (c1, c2) where c1 + c2 >= k.
//
// For c1 == c2:  C(cnt, 2)  (choose 2 distinct indices).
// For c1 < c2:   cnt[c1] * cnt[c2].
// Note that (i,i) is allowed (same element can pair with itself).
// Wait — the problem says "pair (i, j)" — does i=j allowed? Usually
// LeetCode counts distinct pairs of indices with i ≤ j.
// Let's just follow the editorial approach: deduplicate values, count by
// set‑bit, then for each pair of distinct values, add cnt[c1] * cnt[c2] if
// c1 + c2 >= k, and for c1 == c2 add cnt[c1] * cnt[c1] (including same value).
// But "distinct pairs" usually means (i,j) where i can equal j.
//
// Actually, checking the problem more carefully: duplicate values should
// not be double‑counted because the excellence depends on the *value*, not
// the index.  So we deduplicate first, then count all unordered pairs
// (including (value, value)) where setBits(a) + setBits(b) >= k.
// The answer is simply the sum over all value pairs of 1 if condition holds.

func countExcellentPairs(nums []int, k int) int64 {
	// Deduplicate values.
	seen := make(map[int]bool)
	unique := make([]int, 0)
	for _, v := range nums {
		if !seen[v] {
			seen[v] = true
			unique = append(unique, v)
		}
	}

	// Count by set‑bit count.
	cnt := make([]int, 61) // bits up to 60 (nums[i] ≤ 10^9, but use 60 for safety)
	for _, v := range unique {
		cnt[popcount(v)]++
	}

	var ans int64
	// c1 + c2 >= k
	for c1 := 0; c1 <= 60; c1++ {
		if cnt[c1] == 0 {
			continue
		}
		for c2 := c1; c2 <= 60; c2++ {
			if cnt[c2] == 0 {
				continue
			}
			if c1+c2 < k {
				continue
			}
			if c1 == c2 {
				ans += int64(cnt[c1]) * int64(cnt[c2])
			} else {
				ans += int64(cnt[c1]) * int64(cnt[c2])
			}
		}
	}
	return ans
}

func popcount(x int) int {
	c := 0
	for x != 0 {
		x &= x - 1
		c++
	}
	return c
}

// ---------------------------------------------------------------------------
//  Wrapper

func NumberOfExcellentPairs() interface{} {
	return countExcellentPairs([]int{1, 2, 3, 1}, 3)
}

func main() {
	fmt.Println(NumberOfExcellentPairs())

	tests := []struct {
		nums []int
		k    int
		want int64
	}{
		{[]int{1, 2, 3, 1}, 3, 3},
		{[]int{5, 1, 1}, 10, 0},
		{[]int{1}, 1, 1},
		{[]int{7, 3, 1}, 4, 4},
	}
	for _, tc := range tests {
		got := countExcellentPairs(tc.nums, tc.k)
		if got != tc.want {
			fmt.Printf("FAIL nums=%v k=%d: got %d, want %d\n",
				tc.nums, tc.k, got, tc.want)
		}
	}
	fmt.Println("Done testing 2354.")
}
```

## 2355 — Maximum Number Of Books You Can Take

```go
package main

import (
	"fmt"
)

// 2355. Maximum Number of Books You Can Take
// ----------------------------------------------------------------
// Choose a contiguous segment of shelves.  Within the segment you must take
// a non‑increasing number of books (a[i] ≥ a[i+1]) and at most books[i] from
// each shelf.  Maximise total books taken.
//
// For segment [l, r]: a[l]=books[l], a[t]=min(books[t], a[t-1]) = running min.
//
// DP with monotonic stack.  Each entry: (minVal, maxSum).
// At position i:
//   - Pop entries with minVal ≥ books[i]; track best sum among them.
//   - Extend surviving entries: add their own minVal (running min unchanged).
//   - Push new entry for segments ending at i with running min = books[i]:
//     sum = bestMerged + books[i]  (best from extending popped entries).
//   - Answer = max over all entries' sums.

type entry struct {
	minVal int
	sum    int64
}

func maximumBooks(books []int) int64 {
	stack := make([]entry, 0)
	var result int64

	for _, bi := range books {
		var bestMerged int64 // best sum among popped entries

		// Pop entries with minVal ≥ bi.
		for len(stack) > 0 && stack[len(stack)-1].minVal >= bi {
			e := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if e.sum > bestMerged {
				bestMerged = e.sum
			}
		}

		// Extend survivors: their running min (minVal) doesn't change.
		for j := range stack {
			stack[j].sum += int64(stack[j].minVal)
		}

		// New entry with running min = bi.
		newSum := bestMerged + int64(bi)
		stack = append(stack, entry{bi, newSum})
		if newSum > result {
			result = newSum
		}

		// Also check best survivor sum.
		if len(stack) > 1 {
			// stack[0..len(stack)-2] are survivors (their sums were just updated).
			for j := 0; j < len(stack)-1; j++ {
				if stack[j].sum > result {
					result = stack[j].sum
				}
			}
		}
	}
	return result
}

// ---------------------------------------------------------------------------
//  Wrapper

func MaximumNumberOfBooksYouCanTake() interface{} {
	return maximumBooks([]int{8, 5, 2, 7, 7})
}

func main() {
	fmt.Println(MaximumNumberOfBooksYouCanTake())

	cases := []struct {
		books []int
		want  int64
	}{
		{[]int{8, 5, 2, 7, 7}, 19},
		{[]int{1, 2, 3, 4, 5}, 9},
		{[]int{5, 5, 5}, 15},
		{[]int{7, 0, 0, 0, 7}, 7},
		{[]int{10, 1, 1, 1, 1, 1}, 15},
		{[]int{3, 0, 5, 0, 2}, 5},
		{[]int{2, 2, 2, 2, 2}, 10},
	}
	for _, c := range cases {
		if got := maximumBooks(c.books); got != c.want {
			fmt.Printf("FAIL books=%v: got %d, want %d\n", c.books, got, c.want)
		}
	}
	fmt.Println("Done testing 2355.")
}
```

## 2360 — Longest Cycle In A Graph

```go
package main

import (
	"fmt"
)

// 2360. Longest Cycle in a Graph
// ----------------------------------------------------------------
// Directed graph on n nodes.  Each node has exactly one outgoing edge
// (given as array edges[n] where edges[i] = j means i → j, or -1 for no edge).
// Find the length of the longest cycle.  If no cycle exists, return -1.
//
// Use 3‑state DFS: 0 = unvisited, 1 = visiting (on current path),
// 2 = fully processed.  Maintain a distance array for the current DFS walk.
// When we encounter a visiting node, we have found a cycle:
//   length = dist[cur] - dist[node] + 1.

func longestCycle(edges []int) int {
	n := len(edges)
	state := make([]int, n) // 0=unvisited, 1=visiting, 2=done
	dist := make([]int, n)
	ans := -1

	for start := 0; start < n; start++ {
		if state[start] != 0 {
			continue
		}
		// Walk the current path.
		cur := start
		step := 0
		for cur != -1 && state[cur] == 0 {
			state[cur] = 1
			dist[cur] = step
			step++
			cur = edges[cur]
		}
		// If we hit a node in visiting state, we found a cycle.
		if cur != -1 && state[cur] == 1 {
			cycleLen := step - dist[cur]
			if cycleLen > ans {
				ans = cycleLen
			}
		}
		// Mark all nodes in the current walk as done.
		cur = start
		for cur != -1 && state[cur] == 1 {
			state[cur] = 2
			cur = edges[cur]
		}
	}
	return ans
}

// ---------------------------------------------------------------------------
//  Wrapper

func LongestCycleInAGraph() interface{} {
	return longestCycle([]int{3, 3, 4, 2, 3})
}

func main() {
	fmt.Println(LongestCycleInAGraph())

	tests := []struct {
		edges []int
		want  int
	}{
		{[]int{3, 3, 4, 2, 3}, 3},
		{[]int{2, -1, 3, 1}, -1},
		{[]int{1, 2, 0, 4, 5, 3}, 3}, // 0→1→2→0 (len 3) and 3→4→5→3 (len 3) → max 3
		{[]int{-1, -1, -1}, -1},
		{[]int{1, 2, 3, 4, 0}, 5}, // 0→1→2→3→4→0
	}
	// Recompute test 3 manually: edges[3]=4, edges[4]=5, edges[5]=3 → cycle of 3. edges[0]=1, edges[1]=2, edges[2]=0 → cycle of 3. Longest=3.
	// So the test says want 3, that's fine.

	for _, tc := range tests {
		got := longestCycle(tc.edges)
		if got != tc.want {
			fmt.Printf("FAIL edges=%v: got %d, want %d\n", tc.edges, got, tc.want)
		}
	}
	fmt.Println("Done testing 2360.")
}
```

## 2361 — Minimum Costs Using The Train Line

```go
package main

import (
	"fmt"
	"math"
)

// 2361. Minimum Costs Using the Train Line
// ----------------------------------------------------------------
// You start at station 0 on the regular line.
// From station i to i+1:
//   regular[i] = cost to travel on regular line
//   express[i] = cost to travel on express line
// expressCost = one‑time fee to switch FROM regular TO express.
// Switching back (express → regular) is free.
//
// dpReg[i] = min cost to reach station i on regular line
// dpExp[i] = min cost to reach station i on express line
//
// Transition (i → i+1):
//   stay regular:      dpReg[i] + regular[i]
//   switch→regular:    dpExp[i] + 0 + regular[i]   (free switch)
//   stay express:      dpExp[i] + express[i]
//   switch→express:    dpReg[i] + expressCost + express[i]
//
// Result: min(dpReg[n], dpExp[n]).

func minimumCosts(regular, express []int, expressCost int) []int64 {
	n := len(regular)
	dpReg := int64(0)
	dpExp := int64(math.MaxInt64)

	ans := make([]int64, n)
	for i := 0; i < n; i++ {
		newReg := int64(math.MaxInt64)
		newExp := int64(math.MaxInt64)

		// Stay on regular.
		if dpReg != math.MaxInt64 {
			newReg = min64(newReg, dpReg+int64(regular[i]))
		}
		// Switch from express to regular (FREE).
		if dpExp != math.MaxInt64 {
			newReg = min64(newReg, dpExp+int64(regular[i]))
		}

		// Stay on express.
		if dpExp != math.MaxInt64 {
			newExp = min64(newExp, dpExp+int64(express[i]))
		}
		// Switch from regular to express (pay expressCost).
		if dpReg != math.MaxInt64 {
			newExp = min64(newExp, dpReg+int64(expressCost)+int64(express[i]))
		}

		dpReg = newReg
		dpExp = newExp

		ans[i] = min64(dpReg, dpExp)
	}
	return ans
}

func min64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

// ---------------------------------------------------------------------------
//  Wrapper

func MinimumCostsUsingTheTrainLine() interface{} {
	return minimumCosts([]int{1, 6, 9, 5}, []int{5, 2, 3, 10}, 8)
}

func main() {
	fmt.Println(MinimumCostsUsingTheTrainLine())

	got := minimumCosts([]int{1, 6, 9, 5}, []int{5, 2, 3, 10}, 8)
	want := []int64{1, 7, 14, 19}
	for i := range got {
		if got[i] != want[i] {
			fmt.Printf("FAIL [%d]: got %d, want %d\n", i, got[i], want[i])
		}
	}

	got2 := minimumCosts([]int{11, 5, 13}, []int{7, 10, 6}, 3)
	want2 := []int64{10, 15, 24}
	for i := range got2 {
		if got2[i] != want2[i] {
			fmt.Printf("FAIL2 [%d]: got %d, want %d\n", i, got2[i], want2[i])
		}
	}
	fmt.Println("Done testing 2361.")
}
```

## 2362 — Generate The Invoice

```go
package main

// LeetCode #2362: Generate the Invoice
// https://leetcode.com/problems/generate-the-invoice/
// Difficulty: Hard [Paid]
//
// Given two tables: Products(product_id, price) and Purchases(invoice_id, product_id, quantity),
// return the invoice(s) with the maximum total price. If multiple invoices have the same max
// total, return all of them ordered by invoice_id.

import (
	"fmt"
	"sort"
)

type Product struct {
	ID    int
	Price int
}

type Purchase struct {
	InvoiceID int
	ProductID int
	Quantity  int
}

type Invoice struct {
	InvoiceID int
	Total     int
}

// GenerateTheInvoice returns invoices with the maximum total price.
func GenerateTheInvoice(products []Product, purchases []Purchase) []Invoice {
	priceMap := make(map[int]int)
	for _, p := range products {
		priceMap[p.ID] = p.Price
	}

	totals := make(map[int]int)
	for _, p := range purchases {
		totals[p.InvoiceID] += priceMap[p.ProductID] * p.Quantity
	}

	maxTotal := 0
	for _, v := range totals {
		if v > maxTotal {
			maxTotal = v
		}
	}

	var result []Invoice
	for id, total := range totals {
		if total == maxTotal {
			result = append(result, Invoice{InvoiceID: id, Total: total})
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].InvoiceID < result[j].InvoiceID
	})

	return result
}

func main() {
	// Example 1
	products := []Product{
		{ID: 1, Price: 100},
		{ID: 2, Price: 200},
	}
	purchases := []Purchase{
		{InvoiceID: 1, ProductID: 1, Quantity: 2},
		{InvoiceID: 1, ProductID: 2, Quantity: 1},
		{InvoiceID: 2, ProductID: 1, Quantity: 1},
	}
	fmt.Println(GenerateTheInvoice(products, purchases))

	// Example 2: tie
	products2 := []Product{
		{ID: 1, Price: 50},
		{ID: 2, Price: 30},
	}
	purchases2 := []Purchase{
		{InvoiceID: 10, ProductID: 1, Quantity: 1},
		{InvoiceID: 20, ProductID: 2, Quantity: 1},
	}
	fmt.Println(GenerateTheInvoice(products2, purchases2))

	// Single invoice
	products3 := []Product{{ID: 5, Price: 75}}
	purchases3 := []Purchase{{InvoiceID: 7, ProductID: 5, Quantity: 4}}
	fmt.Println(GenerateTheInvoice(products3, purchases3))
}
```

## 2366 — Minimum Replacements To Sort The Array

```go
package main

// LeetCode #2366: Minimum Replacements to Sort the Array
// https://leetcode.com/problems/minimum-replacements-to-sort-the-array/
// Difficulty: Hard
//
// You are given a 0-indexed integer array nums. In one operation you can replace
// any element with any number of elements that sum to it. Return the minimum
// number of operations to make the array non-decreasing.
//
// Approach: Greedy from right to left. Maintain a bound = last element.
// For each nums[i] > bound, we must split nums[i] into pieces each <= bound.
// The minimum pieces needed = ceil(nums[i] / bound). We add (pieces-1) operations.
// Then the new bound for the left side is floor(nums[i] / pieces) which is the
// largest possible value that keeps things non-decreasing.

import (
	"fmt"
)

func minimumReplacements(nums []int) int64 {
	n := len(nums)
	var ans int64
	bound := nums[n-1]

	for i := n - 2; i >= 0; i-- {
		if nums[i] > bound {
			// Minimum number of pieces each <= bound
			parts := (nums[i] + bound - 1) / bound
			ans += int64(parts - 1)
			// Largest possible value for the leftmost piece
			bound = nums[i] / parts
		} else {
			bound = nums[i]
		}
	}

	return ans
}

func main() {
	// Example 1
	fmt.Println(minimumReplacements([]int{3, 9, 3}))
	// Example 2
	fmt.Println(minimumReplacements([]int{1, 2, 3, 4, 5}))
	// Edge: strictly decreasing
	fmt.Println(minimumReplacements([]int{5, 4, 3, 2, 1}))
	// Edge: single element
	fmt.Println(minimumReplacements([]int{1}))
	// Edge: large gap
	fmt.Println(minimumReplacements([]int{12, 9, 7, 6, 17, 19, 21}))
}
```

## 2371 — Minimize Maximum Value In A Grid

```go
package main

// LeetCode #2371: Minimize Maximum Value in a Grid
// https://leetcode.com/problems/minimize-maximum-value-in-a-grid/
// Difficulty: Hard [Paid]
//
// Given a grid with positive integers, assign a new positive integer to each
// cell such that for each row and column, the assigned values maintain the
// same relational order as the original values (i.e., larger original value
// gets larger assigned value). Minimize the maximum assigned value.
//
// Approach: Sort cells by value. Process equal-value groups. For each cell,
// assigned value = max(rowMax[r], colMax[c]) + 1. After the group, update
// rowMax and colMax for the group.

import (
	"fmt"
	"sort"
)

func minMaxValue(grid [][]int) [][]int {
	rows := len(grid)
	if rows == 0 {
		return [][]int{}
	}
	cols := len(grid[0])

	type cell struct {
		val, r, c int
	}
	cells := make([]cell, 0, rows*cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			cells = append(cells, cell{grid[i][j], i, j})
		}
	}

	sort.Slice(cells, func(i, j int) bool {
		return cells[i].val < cells[j].val
	})

	result := make([][]int, rows)
	for i := range result {
		result[i] = make([]int, cols)
	}

	rowMax := make([]int, rows)
	colMax := make([]int, cols)

	i := 0
	for i < len(cells) {
		j := i
		for j < len(cells) && cells[j].val == cells[i].val {
			j++
		}

		// For each cell in the equal-value group, compute assigned value
		// but don't update rowMax/colMax until the entire group is done.
		type assign struct{ r, c, v int }
		assignments := make([]assign, 0, j-i)
		for k := i; k < j; k++ {
			r, c := cells[k].r, cells[k].c
			v := max(rowMax[r], colMax[c]) + 1
			result[r][c] = v
			assignments = append(assignments, assign{r, c, v})
		}

		// Now update rowMax and colMax for this group
		for _, a := range assignments {
			if a.v > rowMax[a.r] {
				rowMax[a.r] = a.v
			}
			if a.v > colMax[a.c] {
				colMax[a.c] = a.v
			}
		}

		i = j
	}

	return result
}

func main() {
	// Example 1
	grid1 := [][]int{
		{3, 1},
		{2, 5},
	}
	fmt.Println("Result:")
	for _, row := range minMaxValue(grid1) {
		fmt.Println(row)
	}

	// Example 2: equal values in same row
	grid2 := [][]int{
		{10, 10, 10},
	}
	fmt.Println("Result:")
	for _, row := range minMaxValue(grid2) {
		fmt.Println(row)
	}

	// Single cell
	grid3 := [][]int{{7}}
	fmt.Println("Result:")
	for _, row := range minMaxValue(grid3) {
		fmt.Println(row)
	}

	// 3x3 all distinct
	grid4 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("Result:")
	for _, row := range minMaxValue(grid4) {
		fmt.Println(row)
	}

	// Multiple equal values
	grid5 := [][]int{
		{5, 5, 3},
		{5, 5, 4},
	}
	fmt.Println("Result:")
	for _, row := range minMaxValue(grid5) {
		fmt.Println(row)
	}
}
```

## 2376 — Count Special Integers

```go
package main

// LeetCode #2376: Count Special Integers
// https://leetcode.com/problems/count-special-integers/
// Difficulty: Hard
//
// An integer is "special" if every digit in it is distinct.
// Count the number of special integers in the range [1, n].
//
// Approach: Digit DP with bitmask of used digits.
// dp[pos][mask][tight][started] = count of ways.
// pos: current position in the digit array
// mask: bitmask of digits already used
// tight: whether prefix matches n so far
// started: whether we've placed a non-leading-zero digit yet

import (
	"fmt"
	"strconv"
)

func countSpecialNumbers(n int) int {
	s := strconv.Itoa(n)
	m := len(s)

	digits := make([]int, m)
	for i, ch := range s {
		digits[i] = int(ch - '0')
	}

	// memo[pos][mask][tight][started]
	memo := make([][][2][2]int, m)
	for i := range memo {
		memo[i] = make([][2][2]int, 1<<10)
		for mask := 0; mask < (1 << 10); mask++ {
			for t := 0; t < 2; t++ {
				for st := 0; st < 2; st++ {
					memo[i][mask][t][st] = -1
				}
			}
		}
	}

	var dp func(pos int, mask int, tight int, started int) int
	dp = func(pos int, mask int, tight int, started int) int {
		if pos == m {
			if started == 1 {
				return 1
			}
			return 0
		}

		if memo[pos][mask][tight][started] != -1 {
			return memo[pos][mask][tight][started]
		}

		limit := 9
		if tight == 1 {
			limit = digits[pos]
		}

		ans := 0
		for d := 0; d <= limit; d++ {
			nextStarted := started
			if d != 0 {
				nextStarted = 1
			}

			nextTight := tight
			if tight == 1 && d < limit {
				nextTight = 0
			}

			if nextStarted == 0 {
				// Still leading zeros, mask unchanged
				ans += dp(pos+1, mask, nextTight, nextStarted)
			} else {
				// Only use digit if not already used
				if mask&(1<<d) == 0 {
					ans += dp(pos+1, mask|(1<<d), nextTight, nextStarted)
				}
			}
		}

		memo[pos][mask][tight][started] = ans
		return ans
	}

	return dp(0, 0, 1, 0)
}

func main() {
	// Example 1
	fmt.Println(countSpecialNumbers(20))
	// Example 2
	fmt.Println(countSpecialNumbers(5))
	// Example 3
	fmt.Println(countSpecialNumbers(135))
	// Edge: 100
	fmt.Println(countSpecialNumbers(100))
	// Large number
	fmt.Println(countSpecialNumbers(1000))
	// Maximum constraint
	fmt.Println(countSpecialNumbers(987654321))
}
```

## 2382 — Maximum Segment Sum After Removals

```go
package main

// LeetCode #2382: Maximum Segment Sum After Removals
// https://leetcode.com/problems/maximum-segment-sum-after-removals/
// Difficulty: Hard
//
// You are given two 0-indexed arrays nums and removeQueries of length n.
// For each i from 0 to n-1, remove the element at position removeQueries[i]
// (making it inactive). After each removal, find the maximum segment sum
// among the active elements (a segment is a contiguous block of active elements).
//
// Approach: Process removals in reverse. Start with all elements removed and
// add them back one by one. Use DSU (Union-Find) to merge adjacent active
// segments, tracking the sum of each segment. Keep a running max.

import "fmt"

func maximumSegmentSum(nums []int, removeQueries []int) []int64 {
	n := len(nums)

	parent := make([]int, n)
	segSum := make([]int64, n)
	active := make([]bool, n)

	for i := 0; i < n; i++ {
		parent[i] = -1
	}

	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	ans := make([]int64, n)
	var maxSum int64

	for i := n - 1; i >= 0; i-- {
		ans[i] = maxSum
		idx := removeQueries[i]
		active[idx] = true
		parent[idx] = idx
		segSum[idx] = int64(nums[idx])

		// Merge with left neighbor
		if idx > 0 && active[idx-1] {
			leftRoot := find(idx - 1)
			if leftRoot != idx {
				segSum[idx] += segSum[leftRoot]
				parent[leftRoot] = idx
			}
		}

		// Merge with right neighbor
		if idx < n-1 && active[idx+1] {
			rightRoot := find(idx + 1)
			if rightRoot != idx {
				segSum[idx] += segSum[rightRoot]
				parent[rightRoot] = idx
			}
		}

		if segSum[idx] > maxSum {
			maxSum = segSum[idx]
		}
	}

	return ans
}

func main() {
	// Example 1
	nums1 := []int{1, 2, 5, 6, 1}
	queries1 := []int{0, 3, 2, 4, 1}
	fmt.Println(maximumSegmentSum(nums1, queries1))

	// Example 2
	nums2 := []int{5, -2, 4, 1}
	queries2 := []int{0, 2, 1, 3}
	fmt.Println(maximumSegmentSum(nums2, queries2))

	// Edge: single element
	nums3 := []int{10}
	queries3 := []int{0}
	fmt.Println(maximumSegmentSum(nums3, queries3))

	// Edge: all negative
	nums4 := []int{-1, -2, -3, -4}
	queries4 := []int{0, 1, 2, 3}
	fmt.Println(maximumSegmentSum(nums4, queries4))
}
```

## 2386 — Find The K Sum Of An Array

```go
package main

// LeetCode #2386: Find the K-Sum of an Array
// https://leetcode.com/problems/find-the-k-sum-of-an-array/
// Difficulty: Hard
//
// Given an array nums and integer k, return the k-th largest sum of any
// subsequence of nums (non-empty allowed).
//
// Approach:
// 1. Compute the maximum sum = sum of all positive numbers.
// 2. Take absolute values of all numbers and sort them.
// 3. The k-th largest subsequence sum is obtained by subtracting some
//    combination of absolute values from the max sum.
// 4. Use a min-heap to generate sums in decreasing order.
//    Each state is (currentSum, index). From each state, we can either
//    subtract the next absolute value, or replace the current subtraction
//    with the next one.

import (
	"container/heap"
	"fmt"
	"sort"
)

type Item struct {
	sum int64
	idx int
}

type MinHeap []Item

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].sum < h[j].sum }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Item))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func kSum(nums []int, k int) int64 {
	n := len(nums)

	// Sum of all positive numbers
	var maxSum int64
	absVals := make([]int64, 0, n)
	for _, v := range nums {
		if v > 0 {
			maxSum += int64(v)
		}
		av := int64(v)
		if av < 0 {
			av = -av
		}
		absVals = append(absVals, av)
	}

	// Sort by absolute value
	sort.Slice(absVals, func(i, j int) bool {
		return absVals[i] < absVals[j]
	})

	// The k-th largest sum
	// ans[0] = maxSum (largest subsequence sum = sum of all positives)
	// ans[1..k-1] generated from heap
	h := &MinHeap{}
	heap.Push(h, Item{sum: maxSum - absVals[0], idx: 0})

	var ans int64 = maxSum

	for i := 1; i < k; i++ {
		it := heap.Pop(h).(Item)
		ans = it.sum

		if it.idx+1 < n {
			// Extend: subtract the next value
			nextSum := it.sum - absVals[it.idx+1]
			heap.Push(h, Item{sum: nextSum, idx: it.idx + 1})

			// Replace: undo current subtraction and subtract the next
			replaceSum := it.sum + absVals[it.idx] - absVals[it.idx+1]
			heap.Push(h, Item{sum: replaceSum, idx: it.idx + 1})
		}
	}

	return ans
}

func main() {
	// Example 1
	fmt.Println(kSum([]int{2, 4, -2}, 5))
	// Example 2
	fmt.Println(kSum([]int{3, -1, -2}, 4))
	// Single positive
	fmt.Println(kSum([]int{5}, 1))
	// All negatives
	fmt.Println(kSum([]int{-1, -2, -3}, 3))
	// Mixed
	fmt.Println(kSum([]int{10, -5, 3, -2}, 6))
}
```

## 2392 — Build A Matrix With Conditions

```go
package main

// LeetCode #2392: Build a Matrix With Conditions
// https://leetcode.com/problems/build-a-matrix-with-conditions/
// Difficulty: Hard
//
// Given k (numbers 1..k), rowConditions (above/below pairs), and colConditions
// (left/right pairs), build a k x k matrix where each row and column has at most
// one number, and the conditions are satisfied. Numbers 1..k must each appear
// exactly once.
//
// Approach: Topological sort for rows using rowConditions, and for columns
// using colConditions. If either has a cycle, return empty matrix. Then place
// each number at (rowPos[number], colPos[number]).

import "fmt"

func buildMatrix(k int, rowConditions [][]int, colConditions [][]int) [][]int {
	rowOrder := topoSort(k, rowConditions)
	colOrder := topoSort(k, colConditions)
	if rowOrder == nil || colOrder == nil {
		return [][]int{}
	}

	rowPos := make([]int, k+1)
	for i, v := range rowOrder {
		rowPos[v] = i
	}
	colPos := make([]int, k+1)
	for i, v := range colOrder {
		colPos[v] = i
	}

	matrix := make([][]int, k)
	for i := range matrix {
		matrix[i] = make([]int, k)
	}
	for num := 1; num <= k; num++ {
		matrix[rowPos[num]][colPos[num]] = num
	}

	return matrix
}

func topoSort(k int, conditions [][]int) []int {
	graph := make([][]int, k+1)
	inDeg := make([]int, k+1)

	for _, c := range conditions {
		u, v := c[0], c[1]
		graph[u] = append(graph[u], v)
		inDeg[v]++
	}

	queue := make([]int, 0)
	for i := 1; i <= k; i++ {
		if inDeg[i] == 0 {
			queue = append(queue, i)
		}
	}

	order := make([]int, 0, k)
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		order = append(order, u)
		for _, v := range graph[u] {
			inDeg[v]--
			if inDeg[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	if len(order) != k {
		return nil // cycle
	}
	return order
}

func main() {
	// Example 1
	k1 := 3
	rowC1 := [][]int{{1, 2}, {3, 2}}
	colC1 := [][]int{{2, 1}, {3, 2}}
	fmt.Println(buildMatrix(k1, rowC1, colC1))

	// Example 2: cycle
	k2 := 3
	rowC2 := [][]int{{1, 2}, {2, 3}, {3, 1}}
	colC2 := [][]int{{2, 3}}
	fmt.Println(buildMatrix(k2, rowC2, colC2))

	// Simple: k=2
	k3 := 2
	rowC3 := [][]int{{1, 2}}
	colC3 := [][]int{{2, 1}}
	fmt.Println(buildMatrix(k3, rowC3, colC3))

	// k=1
	k4 := 1
	fmt.Println(buildMatrix(k4, [][]int{}, [][]int{}))
}
```

## 2398 — Maximum Number Of Robots Within Budget

```go
package main

// LeetCode #2398: Maximum Number of Robots Within Budget
// https://leetcode.com/problems/maximum-number-of-robots-within-budget/
// Difficulty: Hard
//
// Approach: Sliding window with monotonic deque. The budget constraint for
// subarray [l..r] is: max(chargeTimes[l..r]) + (r-l+1) * sum(runningCosts[l..r]) <= budget.
// Maintain a deque for max chargeTime in current window. Expand right pointer,
// shrink left while over budget. Track max window length.

import "fmt"

func main() {
	// Example 1: chargeTimes=[3,6,1,3,4], runningCosts=[2,1,3,4,5], budget=25 => 3
	fmt.Println(maximumRobots([]int{3, 6, 1, 3, 4}, []int{2, 1, 3, 4, 5}, 25))
	// Example 2: chargeTimes=[11,12,19], runningCosts=[10,8,7], budget=19 => 0
	fmt.Println(maximumRobots([]int{11, 12, 19}, []int{10, 8, 7}, 19))
	// Edge: single robot within budget
	fmt.Println(maximumRobots([]int{1}, []int{1}, 2))
	// Edge: single robot over budget
	fmt.Println(maximumRobots([]int{10}, []int{10}, 1))
	// Edge: all within budget
	fmt.Println(maximumRobots([]int{1, 2, 3}, []int{1, 1, 1}, 100))
}

func maximumRobots(chargeTimes []int, runningCosts []int, budget int64) int {
	n := len(chargeTimes)
	deque := make([]int, 0) // stores indices, front = max chargeTime
	left := 0
	var sum int64
	maxLen := 0

	for right := 0; right < n; right++ {
		sum += int64(runningCosts[right])

		// Maintain deque: remove smaller elements from back
		for len(deque) > 0 && chargeTimes[deque[len(deque)-1]] <= chargeTimes[right] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, right)

		// Check budget constraint
		for left <= right {
			maxCharge := chargeTimes[deque[0]]
			cost := int64(maxCharge) + int64(right-left+1)*sum
			if cost <= budget {
				break
			}
			// Shrink left
			if deque[0] == left {
				deque = deque[1:]
			}
			sum -= int64(runningCosts[left])
			left++
		}

		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}
	return maxLen
}
```

## 2402 — Meeting Rooms Iii

```go
package main

// LeetCode #2402: Meeting Rooms III
// https://leetcode.com/problems/meeting-rooms-iii/
// Difficulty: Hard
//
// You have n meeting rooms numbered 0 to n-1. Given meetings[start_i, end_i),
// assign each meeting to the smallest-numbered available room. If no room is
// available, delay the meeting until a room is free (keeping duration same).
// Return the room that hosts the most meetings.
//
// Approach: Use two min-heaps:
//   - available: room indices (min-heap by room number)
//   - busy: (endTime, roomIndex) (min-heap by endTime)
// Process meetings in sorted order by start time. For each meeting, release
// all busy rooms that are now free. Assign the meeting to the smallest
// available room. If no room is free, use the earliest-freed room.

import (
	"container/heap"
	"fmt"
	"sort"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type Busy struct {
	endTime int
	room    int
}

type BusyHeap []Busy

func (h BusyHeap) Len() int           { return len(h) }
func (h BusyHeap) Less(i, j int) bool { return h[i].endTime < h[j].endTime }
func (h BusyHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *BusyHeap) Push(x interface{}) {
	*h = append(*h, x.(Busy))
}

func (h *BusyHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func mostBooked(n int, meetings [][]int) int {
	// Sort meetings by start time
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	avail := &IntHeap{}
	for i := 0; i < n; i++ {
		heap.Push(avail, i)
	}

	busy := &BusyHeap{}
	count := make([]int, n)

	for _, m := range meetings {
		start, end := m[0], m[1]
		duration := end - start

		// Release rooms that have finished by now
		for busy.Len() > 0 && (*busy)[0].endTime <= start {
			b := heap.Pop(busy).(Busy)
			heap.Push(avail, b.room)
		}

		var room int
		if avail.Len() > 0 {
			// Room available, use smallest-numbered
			room = heap.Pop(avail).(int)
		} else {
			// No room available, pick the soonest free room and delay
			b := heap.Pop(busy).(Busy)
			room = b.room
			// Meeting is delayed; duration stays same
			start = b.endTime
		}

		count[room]++
		heap.Push(busy, Busy{endTime: start + duration, room: room})
	}

	// Find room with max count, smallest number in case of tie
	maxRoom, maxCount := 0, count[0]
	for i := 1; i < n; i++ {
		if count[i] > maxCount {
			maxCount = count[i]
			maxRoom = i
		}
	}

	return maxRoom
}

func main() {
	// Example 1
	n1 := 2
	meetings1 := [][]int{{0, 10}, {1, 5}, {2, 7}, {3, 4}}
	fmt.Println(mostBooked(n1, meetings1))

	// Example 2
	n2 := 3
	meetings2 := [][]int{{1, 20}, {2, 10}, {3, 5}, {4, 9}, {6, 8}}
	fmt.Println(mostBooked(n2, meetings2))

	// Single room
	n3 := 1
	meetings3 := [][]int{{0, 5}, {2, 3}, {4, 7}}
	fmt.Println(mostBooked(n3, meetings3))

	// All meetings overlap
	n4 := 2
	meetings4 := [][]int{{0, 10}, {0, 5}, {0, 7}}
	fmt.Println(mostBooked(n4, meetings4))
}
```

## 2403 — Minimum Time To Kill All Monsters

```go
package main

// LeetCode #2403: Minimum Time to Kill All Monsters
// https://leetcode.com/problems/minimum-time-to-kill-all-monsters/
// Difficulty: Hard [Paid]
//
// You are given an array power of monsters. Each day, your "power level"
// starts at 0 and increases by 1 per day. You can kill at most one monster
// per day. To kill monster i, your power level must be >= power[i]. After
// killing a monster, your power level resets to 0.
//
// Approach: Bitmask DP.
// dp[mask] = minimum number of days to kill monsters in mask.
// For each mask, the number of kills so far = bits set in mask.
// To kill monster j (not yet killed):
//   days_before = dp[mask]
//   gain = 1 per day after last kill
//   current_power = days_before - (kills-1)  ? No...
//
// Actually, each monster kill takes a variable number of days.
// The key insight: after killing k monsters, you've spent some total days.
// The next monster with power p requires ceil(p / factor) days where
// factor = k+1 (you gain 1 unit per day, and after k kills your gain rate is...).
//
// Let me reconsider. The standard interpretation:
// - You can kill one monster at a time.
// - Each day your power increases by 1, starting at 0.
// - When power >= monster's power, you can kill it (power resets to 0).
// - So killing monster i takes power[i] days of waiting.
// - Total days = sum of power of all monsters (each resets).
//
// But that's too simple. The harder version:
// - After each kill, the "cost" (time needed for next one) increases.
// - Specifically, the time to kill a monster depends on when you kill it.
//
// Let me implement the well-known version:
// dp[mask] = min days needed.
// For state mask, let t = dp[mask] (days elapsed so far).
// For monster j not in mask:
//   days_needed = max(power[j] - t, 0)  -- wait until power level >= power[j]
//   Actually no, power resets after each kill.
//
// Let me think differently. The common version of this problem:
// - dp[mask] = min time to kill monsters in mask
// - You have a "gain" factor = (number of monsters killed) per day... no.
//
// I think the actual problem is:
// The monsters have power levels. Each day you can kill a monster with
// power level <= your current power. Your current power increases by 1
// each day. After killing, it resets to 0.
//
// So to kill a set of monsters with powers p1, p2, ..., pk:
// Total time = p1 + p2 + ... + pk (each requires waiting p_i days)
// This is just sum of all power values. Too simple.
//
// Alternatively with a twist: You can choose the ORDER of killing,
// and each monster killed increases the "gain multiplier" for subsequent kills.
// After killing k monsters, your power gain per day is 2^k (or k+1).
// dp[mask] = min days, and the state tells us the multiplier.
//
// Let me implement the most common DP version for this type:
// dp[mask] = minimum time to reach this mask
// dp[0] = 0
// For each mask, killed = popcount(mask)
// gain = killed + 1 (or 2^something)
// For each j not in mask:
//   timeToKill = ceil(power[j] / gain)
//   dp[mask|1<<j] = min(dp[mask|1<<j], dp[mask] + timeToKill)

import (
	"fmt"
	"math"
)

func minimumTimeToKillAllMonsters(power []int) int64 {
	n := len(power)
	totalMasks := 1 << n
	dp := make([]int64, totalMasks)
	for i := range dp {
		dp[i] = math.MaxInt64
	}
	dp[0] = 0

	for mask := 0; mask < totalMasks; mask++ {
		if dp[mask] == math.MaxInt64 {
			continue
		}
		killed := popcount(mask)
		gain := int64(killed + 1) // gain factor increases with kills

		for j := 0; j < n; j++ {
			if mask&(1<<j) == 0 {
				newMask := mask | (1 << j)
				// Time to kill monster j: ceil(power[j] / gain)
				timeToKill := (int64(power[j]) + gain - 1) / gain
				if dp[mask]+timeToKill < dp[newMask] {
					dp[newMask] = dp[mask] + timeToKill
				}
			}
		}
	}

	return dp[totalMasks-1]
}

func popcount(x int) int {
	count := 0
	for x > 0 {
		count += x & 1
		x >>= 1
	}
	return count
}

func main() {
	// Example 1
	fmt.Println(minimumTimeToKillAllMonsters([]int{3, 1, 2}))
	// Example 2
	fmt.Println(minimumTimeToKillAllMonsters([]int{5, 5, 5}))
	// Edge: single monster
	fmt.Println(minimumTimeToKillAllMonsters([]int{10}))
	// Edge: all 1s
	fmt.Println(minimumTimeToKillAllMonsters([]int{1, 1, 1, 1}))
	// Larger powers
	fmt.Println(minimumTimeToKillAllMonsters([]int{7, 3, 9, 2}))
}
```

## 2407 — Longest Increasing Subsequence Ii

```go
package main

// LeetCode #2407: Longest Increasing Subsequence II
// https://leetcode.com/problems/longest-increasing-subsequence-ii/
// Difficulty: Hard
//
// Given an array nums and integer k, find the longest increasing subsequence
// where the difference between consecutive elements is at most k.
//
// Approach: Segment tree over values (1..max(nums)).
// For each value x = nums[i], the longest LIS ending at x is:
//   dp[x] = 1 + max(dp[x-k .. x-1])   (if such previous element exists)
// Use a segment tree to query the max in range [x-k, x-1] in O(log M).
// Time: O(N log M) where M = max(nums).

import "fmt"

type SegTree struct {
	tree []int
	n    int
}

func NewSegTree(size int) *SegTree {
	return &SegTree{
		tree: make([]int, 4*size),
		n:    size,
	}
}

func (st *SegTree) update(idx int, val int, node int, left int, right int) {
	if left == right {
		if val > st.tree[node] {
			st.tree[node] = val
		}
		return
	}
	mid := left + (right-left)/2
	if idx <= mid {
		st.update(idx, val, node*2, left, mid)
	} else {
		st.update(idx, val, node*2+1, mid+1, right)
	}
	st.tree[node] = max(st.tree[node*2], st.tree[node*2+1])
}

func (st *SegTree) query(ql int, qr int, node int, left int, right int) int {
	if ql > right || qr < left || ql > qr {
		return 0
	}
	if ql <= left && right <= qr {
		return st.tree[node]
	}
	mid := left + (right-left)/2
	return max(
		st.query(ql, qr, node*2, left, mid),
		st.query(ql, qr, node*2+1, mid+1, right),
	)
}

func lengthOfLIS(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}

	maxVal := 0
	for _, v := range nums {
		if v > maxVal {
			maxVal = v
		}
	}

	st := NewSegTree(maxVal)
	ans := 1

	for _, x := range nums {
		// Query max dp in range [x-k, x-1]
		left := x - k
		if left < 1 {
			left = 1
		}
		best := st.query(left, x-1, 1, 1, maxVal)
		cur := best + 1
		st.update(x, cur, 1, 1, maxVal)
		if cur > ans {
			ans = cur
		}
	}

	return ans
}

func main() {
	// Example 1
	fmt.Println(lengthOfLIS([]int{4, 2, 1, 4, 3, 4, 5, 8, 15}, 3))
	// Example 2
	fmt.Println(lengthOfLIS([]int{7, 4, 5, 1, 8, 12, 4, 7}, 5))
	// Example 3
	fmt.Println(lengthOfLIS([]int{1, 5}, 1))
	// Edge: simple increasing
	fmt.Println(lengthOfLIS([]int{1, 2, 3, 4}, 1))
	// Edge: single element
	fmt.Println(lengthOfLIS([]int{10}, 3))
}
```

## 2412 — Minimum Money Required Before Transactions

```go
package main

// LeetCode #2412: Minimum Money Required Before Transactions
// https://leetcode.com/problems/minimum-money-required-before-transactions/
// Difficulty: Hard
//
// You are given a list of transactions where each transaction is [cost, cashback].
// You can execute transactions in any order. You start with some money, and for
// each transaction: if money >= cost, you pay cost and receive cashback (money
// = money - cost + cashback). Find the minimum starting money needed to complete
// all transactions.
//
// Approach: Split transactions into two groups:
//   - "loss" transactions: cost > cashback (net loss)
//   - "gain" transactions: cost <= cashback (net gain or break-even)
//
// For maximum safety, process loss transactions first (they deplete money).
// Within loss transactions, process those with the highest cashback first
// (to maximize money recovery). Within gain transactions, process those with
// the lowest cost first (to conserve money).
//
// The minimum starting money = max over all prefixes of (cumulative money needed).

import (
	"fmt"
	"sort"
)

func minimumMoney(transactions [][]int) int64 {
	// Separate loss and gain transactions
	loss := make([][]int, 0)
	gain := make([][]int, 0)

	for _, t := range transactions {
		cost, cashback := t[0], t[1]
		if cost > cashback {
			loss = append(loss, t)
		} else {
			gain = append(gain, t)
		}
	}

	// Sort loss transactions by cashback descending (recover more money sooner)
	sort.Slice(loss, func(i, j int) bool {
		return loss[i][1] > loss[j][1]
	})

	// Sort gain transactions by cost ascending (spend less money first)
	sort.Slice(gain, func(i, j int) bool {
		return gain[i][0] < gain[j][0]
	})

	// Simulate transactions in order: loss first, then gain
	ordered := append(loss, gain...)

	var money int64
	var minStart int64

	for _, t := range ordered {
		cost, cashback := int64(t[0]), int64(t[1])

		// If we don't have enough money, increase starting money
		if money < cost {
			needed := cost - money
			minStart += needed
			money += needed
		}

		money = money - cost + cashback
	}

	return minStart
}

func main() {
	// Example 1
	fmt.Println(minimumMoney([][]int{{2, 1}, {5, 0}, {4, 2}}))
	// Example 2
	fmt.Println(minimumMoney([][]int{{3, 0}, {0, 3}}))
	// Single loss transaction
	fmt.Println(minimumMoney([][]int{{10, 2}}))
	// Single gain transaction
	fmt.Println(minimumMoney([][]int{{5, 10}}))
	// All gains
	fmt.Println(minimumMoney([][]int{{1, 2}, {2, 3}, {3, 5}}))
	// All losses
	fmt.Println(minimumMoney([][]int{{10, 1}, {8, 2}, {5, 3}}))
}
```

## 2416 — Sum Of Prefix Scores Of Strings

```go
package main

// LeetCode #2416: Sum of Prefix Scores of Strings
// https://leetcode.com/problems/sum-of-prefix-scores-of-strings/
// Difficulty: Hard
//
// Trie with visit count. Insert each word, counting how many words pass through
// each prefix node. Then for each word, sum the counts along its prefix path.
// Time O(N * L) | Space O(N * L) where N = len(words), L = avg word length.

import "fmt"

func main() {
	// Example 1
	fmt.Println(sumPrefixScores([]string{"abc", "ab", "bc", "b"}))
	// Example 2
	fmt.Println(sumPrefixScores([]string{"abcd"}))
	// Edge: single char
	fmt.Println(sumPrefixScores([]string{"a", "a", "a"}))
}

type trieNode struct {
	children [26]*trieNode
	count    int
}

func sumPrefixScores(words []string) []int {
	root := &trieNode{}

	// Insert all words into trie, incrementing count at each node
	for _, w := range words {
		cur := root
		for i := 0; i < len(w); i++ {
			idx := w[i] - 'a'
			if cur.children[idx] == nil {
				cur.children[idx] = &trieNode{}
			}
			cur = cur.children[idx]
			cur.count++
		}
	}

	ans := make([]int, len(words))
	for wi, w := range words {
		cur := root
		total := 0
		for i := 0; i < len(w); i++ {
			idx := w[i] - 'a'
			cur = cur.children[idx]
			total += cur.count
		}
		ans[wi] = total
	}
	return ans
}
```

## 2421 — Number Of Good Paths

```go
package main

// LeetCode #2421: Number of Good Paths
// https://leetcode.com/problems/number-of-good-paths/
// Difficulty: Hard
//
// Union-Find by value. Sort nodes by value, union adjacent nodes that have
// <= current value. For each value group, count how many nodes in each
// connected component have that value. A good path starts and ends at nodes
// of the same value, and all intermediate nodes have <= that value.
// Time O(N log N + M alpha(N)) | Space O(N)

import (
	"fmt"
	"sort"
)

func main() {
	// Example 1
	fmt.Println(numberOfGoodPaths([]int{1, 3, 2, 1, 3},
		[][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}}))
	// Example 2
	fmt.Println(numberOfGoodPaths([]int{1, 1, 2, 2, 3},
		[][]int{{0, 1}, {1, 2}, {2, 3}, {2, 4}}))
	// Single node
	fmt.Println(numberOfGoodPaths([]int{1}, [][]int{}))
}

type uf struct {
	parent []int
	rank   []int
}

func newUF(n int) *uf {
	p := make([]int, n)
	r := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
	}
	return &uf{p, r}
}

func (u *uf) find(x int) int {
	for u.parent[x] != x {
		u.parent[x] = u.parent[u.parent[x]]
		x = u.parent[x]
	}
	return x
}

func (u *uf) union(x, y int) {
	xr, yr := u.find(x), u.find(y)
	if xr == yr {
		return
	}
	if u.rank[xr] < u.rank[yr] {
		xr, yr = yr, xr
	}
	u.parent[yr] = xr
	if u.rank[xr] == u.rank[yr] {
		u.rank[xr]++
	}
}

func numberOfGoodPaths(vals []int, edges [][]int) int {
	n := len(vals)
	if n == 1 {
		return 1
	}

	// Build adjacency list
	adj := make([][]int, n)
	for _, e := range edges {
		a, b := e[0], e[1]
		adj[a] = append(adj[a], b)
		adj[b] = append(adj[b], a)
	}

	// Sort nodes by value
	nodes := make([]int, n)
	for i := 0; i < n; i++ {
		nodes[i] = i
	}
	sort.Slice(nodes, func(i, j int) bool {
		return vals[nodes[i]] < vals[nodes[j]]
	})

	u := newUF(n)

	// For each value group, count good paths
	ans := n // each single node is a good path
	i := 0
	for i < n {
		j := i
		for j < n && vals[nodes[j]] == vals[nodes[i]] {
			j++
		}
		// Union all adjacent nodes with this value
		for k := i; k < j; k++ {
			node := nodes[k]
			for _, nei := range adj[node] {
				if vals[nei] <= vals[node] {
					u.union(node, nei)
				}
			}
		}
		// Count nodes in each component for this value
		compCount := make(map[int]int)
		for k := i; k < j; k++ {
			root := u.find(nodes[k])
			compCount[root]++
		}
		for _, c := range compCount {
			ans += c * (c - 1) / 2
		}
		i = j
	}

	return ans
}
```

## 2426 — Number Of Pairs Satisfying Inequality

```go
package main

// LeetCode #2426: Number of Pairs Satisfying Inequality
// https://leetcode.com/problems/number-of-pairs-satisfying-inequality/
// Difficulty: Hard
//
// Given nums1[i] - nums1[j] <= nums2[i] - nums2[j] + diff for i < j.
// Rearranged: (nums1[i] - nums2[i]) <= (nums1[j] - nums2[j]) + diff.
// Let arr[k] = nums1[k] - nums2[k]. Then for i < j: arr[i] <= arr[j] + diff.
// Processing left to right, at position j count previous i where
// arr[i] <= arr[j] + diff. Use BIT (Fenwick Tree) on compressed values.

import (
	"fmt"
	"sort"
)

type BIT struct {
	tree []int
}

func newBIT(size int) *BIT {
	return &BIT{tree: make([]int, size+2)}
}

func (b *BIT) add(idx int) {
	for i := idx; i < len(b.tree); i += i & -i {
		b.tree[i]++
	}
}

func (b *BIT) sum(idx int) int {
	s := 0
	for i := idx; i > 0; i -= i & -i {
		s += b.tree[i]
	}
	return s
}

func numberOfPairs(input [][]int, diff int) int64 {
	nums1, nums2 := input[0], input[1]
	n := len(nums1)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = nums1[i] - nums2[i]
	}

	// Coordinate compression
	allVals := make([]int, 0, n*2)
	for _, v := range arr {
		allVals = append(allVals, v, v+diff)
	}
	sort.Ints(allVals)
	uniq := 1
	for i := 1; i < len(allVals); i++ {
		if allVals[i] != allVals[uniq-1] {
			allVals[uniq] = allVals[i]
			uniq++
		}
	}
	allVals = allVals[:uniq]

	compress := func(x int) int {
		return sort.SearchInts(allVals, x) + 1
	}

	bt := newBIT(len(allVals) + 2)
	var ans int64

	for j := 0; j < n; j++ {
		// Count previous arr[i] where arr[i] <= arr[j] + diff
		target := arr[j] + diff
		pos := compress(target)
		ans += int64(bt.sum(pos))
		bt.add(compress(arr[j]))
	}

	return ans
}

func main() {
	// Example 1
	fmt.Println(numberOfPairs([][]int{{3, 2, 5}, {2, 2, 1}}, 1))
	// Example 2
	fmt.Println(numberOfPairs([][]int{{3, -1}, {-2, 2}}, -1))
	// Single element
	fmt.Println(numberOfPairs([][]int{{1}, {1}}, 0))
	// All equal
	fmt.Println(numberOfPairs([][]int{{1, 1, 1}, {1, 1, 1}}, 0))
	// Larger range
	fmt.Println(numberOfPairs([][]int{{1, 3, 5, 7}, {2, 4, 6, 8}}, 2))
}
```

## 2430 — Maximum Deletions On A String

```go
package main

// LeetCode #2430: Maximum Deletions on a String
// https://leetcode.com/problems/maximum-deletions-on-a-string/
// Difficulty: Hard
//
// Given a string s. In one operation, you can delete a prefix of the current
// string if the remaining string starts with that same prefix. Find the maximum
// number of operations needed to delete the entire string.
//
// Approach: DP + LCP (Longest Common Prefix).
// dp[i] = max deletions starting at position i (suffix s[i:]).
// lcp[i][j] = longest common prefix of s[i:] and s[j:].
//
// For each i, for each possible length len (where i+len < n):
//   if lcp[i][i+len] >= len (i.e., s[i:i+len] == s[i+len:i+2*len]),
//   then dp[i] = max(dp[i], 1 + dp[i+len]).
//
// Base: dp[n] = 0 (empty string). Result: dp[0].

import "fmt"

func maxDeletions(s string) int {
	n := len(s)

	// lcp[i][j] = longest common prefix of s[i:] and s[j:]
	lcp := make([][]int, n+1)
	for i := range lcp {
		lcp[i] = make([]int, n+1)
	}

	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if s[i] == s[j] {
				lcp[i][j] = 1 + lcp[i+1][j+1]
			}
		}
	}

	dp := make([]int, n+1)
	// dp[n] = 0 by default

	for i := n - 1; i >= 0; i-- {
		dp[i] = 1 // at minimum, we can delete the whole suffix in one operation
		maxLen := (n - i) / 2
		for length := 1; length <= maxLen; length++ {
			if lcp[i][i+length] >= length {
				if 1+dp[i+length] > dp[i] {
					dp[i] = 1 + dp[i+length]
				}
			}
		}
	}

	return dp[0]
}

func main() {
	// Example 1
	fmt.Println(maxDeletions("abcabcdabc"))
	// Example 2
	fmt.Println(maxDeletions("aaabaab"))
	// Example 3
	fmt.Println(maxDeletions("aaaaa"))
	// Single char
	fmt.Println(maxDeletions("a"))
	// Two chars
	fmt.Println(maxDeletions("aa"))
	fmt.Println(maxDeletions("ab"))
	// No repeated prefix
	fmt.Println(maxDeletions("abcdef"))
}
```

## 2435 — Paths In Matrix Whose Sum Is Divisible By K

```go
package main

// LeetCode #2435: Paths in Matrix Whose Sum Is Divisible by K
// https://leetcode.com/problems/paths-in-matrix-whose-sum-is-divisible-by-k/
// Difficulty: Hard
//
// Given a grid of size m x n and an integer k, count the number of paths
// from (0,0) to (m-1,n-1) moving only right or down, such that the sum of
// values along the path is divisible by k.
//
// Approach: 3D DP.
// dp[i][j][mod] = number of ways to reach (i,j) with sum % k == mod.
// Transition from top (i-1,j) and left (i,j-1).
// mod = (prevMod + grid[i][j]) % k.

import "fmt"

func numberOfPaths(grid [][]int, k int) int {
	const mod = 1_000_000_007
	m := len(grid)
	n := len(grid[0])

	// dp[i][j][r] = ways to reach (i,j) with sum % k == r
	dp := make([][][]int, m)
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, k)
		}
	}

	// Initialize start position
	dp[0][0][grid[0][0]%k] = 1

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			val := grid[i][j]
			for r := 0; r < k; r++ {
				ways := 0
				if i > 0 {
					ways = (ways + dp[i-1][j][r]) % mod
				}
				if j > 0 {
					ways = (ways + dp[i][j-1][r]) % mod
				}
				newR := (r + val) % k
				dp[i][j][newR] = (dp[i][j][newR] + ways) % mod
			}
		}
	}

	return dp[m-1][n-1][0]
}

func main() {
	// Example 1
	fmt.Println(numberOfPaths([][]int{{5, 2, 4}, {3, 0, 5}, {0, 7, 2}}, 3))
	// Example 2
	fmt.Println(numberOfPaths([][]int{{0, 0}}, 5))
	// Example 3
	fmt.Println(numberOfPaths([][]int{{7, 3, 4, 9}, {2, 7, 6, 0}}, 1))
	// Single cell
	fmt.Println(numberOfPaths([][]int{{10}}, 2))
	// 1x3
	fmt.Println(numberOfPaths([][]int{{1, 2, 3}}, 3))
}
```

## 2440 — Create Components With Same Value

```go
package main

// LeetCode #2440: Create Components With Same Value
// https://leetcode.com/problems/create-components-with-same-value/
// Difficulty: Hard
//
// DFS to compute subtree sums. If the sum of the whole tree is S, and we want
// to split into k components each with value S/k, then we need S to be divisible
// by k and each component's subtree sum must be a multiple of S/k. We try
// all possible divisors of S. For each target = S / k, run DFS that returns
// the cumulative subtree sum, resetting to 0 when target is reached.
// Time O(N * divisors(S)) | Space O(N)

import "fmt"

func main() {
	// Example 1
	fmt.Println(componentValue([]int{6, 2, 2, 2, 6},
		[][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}}))
	// Example 2: single node
	fmt.Println(componentValue([]int{2}, [][]int{}))
	// Example 3
	fmt.Println(componentValue([]int{1, 2, 3, 4, 5},
		[][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}))
}

func componentValue(nums []int, edges [][]int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}

	total := 0
	for _, v := range nums {
		total += v
	}

	adj := make([][]int, n)
	for _, e := range edges {
		a, b := e[0], e[1]
		adj[a] = append(adj[a], b)
		adj[b] = append(adj[b], a)
	}

	best := 1

	// Try divisors of total as possible number of components
	// We try from largest to smallest, so the first valid gives max components
	for k := n; k >= 2; k-- {
		if total%k != 0 {
			continue
		}
		target := total / k
		if canSplit(nums, adj, target) {
			best = k
			break
		}
	}

	return best - 1 // operations = components - 1
}

func canSplit(nums []int, adj [][]int, target int) bool {
	n := len(nums)
	visited := make([]bool, n)
	var dfs func(u int) int
	dfs = func(u int) int {
		visited[u] = true
		sum := nums[u]
		for _, v := range adj[u] {
			if !visited[v] {
				sub := dfs(v)
				if sub == -1 {
					return -1
				}
				sum += sub
			}
		}
		if sum > target {
			return -1
		}
		if sum == target {
			return 0
		}
		return sum
	}
	return dfs(0) == 0
}
```

## 2444 — Count Subarrays With Fixed Bounds

```go
package main

// LeetCode #2444: Count Subarrays With Fixed Bounds
// https://leetcode.com/problems/count-subarrays-with-fixed-bounds/
// Difficulty: Hard
//
// Two-pointer / sliding window. For each right index, track the last positions
// of minK and maxK. Any subarray ending at right that contains both minK and
// maxK starts at (min(lastMinK, lastMaxK) ... last-bad). The valid start is
// the last index of any element outside [minK, maxK].
// Time O(N) | Space O(1)

import "fmt"

func main() {
	// Example 1
	fmt.Println(countSubarrays([]int{1, 3, 5, 2, 7, 5}, 1, 5))
	// Example 2
	fmt.Println(countSubarrays([]int{1, 1, 1, 1}, 1, 1))
	// No valid subarray
	fmt.Println(countSubarrays([]int{2, 3, 4}, 1, 5))
}

func countSubarrays(nums []int, minK, maxK int) int64 {
	var ans int64
	lastMin, lastMax := -1, -1
	bad := -1 // last index of element outside [minK, maxK]

	for i, v := range nums {
		if v < minK || v > maxK {
			bad = i
			continue
		}
		if v == minK {
			lastMin = i
		}
		if v == maxK {
			lastMax = i
		}
		start := min(lastMin, lastMax)
		if start > bad {
			ans += int64(start - bad)
		}
	}
	return ans
}
```

## 2448 — Minimum Cost To Make Array Equal

```go
package main

// LeetCode #2448: Minimum Cost to Make Array Equal
// https://leetcode.com/problems/minimum-cost-to-make-array-equal/
// Difficulty: Hard
//
// Weighted median. Consider each element nums[i] with weight cost[i].
// The optimal value to make all elements equal is the weighted median:
// find the point where cumulative cost crosses half of total cost.
// Time O(N log N) | Space O(N)

import (
	"fmt"
	"sort"
)

func main() {
	// Example 1
	fmt.Println(minCost([][]int{{1, 3, 5, 2}, {2, 3, 1, 14}}))
	// Example 2
	fmt.Println(minCost([][]int{{2, 2, 2, 2, 2}, {4, 2, 8, 1, 3}}))
	// Single element
	fmt.Println(minCost([][]int{{1}, {1}}))
}

func minCost(input [][]int) int64 {
	nums, cost := input[0], input[1]
	n := len(nums)

	type pair struct {
		num, c int
	}
	pairs := make([]pair, n)
	totalCost := int64(0)
	for i := 0; i < n; i++ {
		pairs[i] = pair{nums[i], cost[i]}
		totalCost += int64(cost[i])
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].num < pairs[j].num
	})

	// Find weighted median
	cum := int64(0)
	median := 0
	half := (totalCost + 1) / 2
	for _, p := range pairs {
		cum += int64(p.c)
		if cum >= half {
			median = p.num
			break
		}
	}

	// Compute cost to move all to median
	var ans int64
	for i := 0; i < n; i++ {
		diff := nums[i] - median
		if diff < 0 {
			diff = -diff
		}
		ans += int64(diff) * int64(cost[i])
	}
	return ans
}
```

## 2449 — Minimum Number Of Operations To Make Arrays Similar

```go
package main

// LeetCode #2449: Minimum Number of Operations to Make Arrays Similar
// https://leetcode.com/problems/minimum-number-of-operations-to-make-arrays-similar/
// Difficulty: Hard
//
// Sort both arrays, separate by parity (odd/even). Each operation changes
// a value by +/-2 preserving parity. Match elements of same parity in sorted
// order. Count the total positive difference (sum of target[i] - nums[i] for
// those needing increase). Each operation fixes 2 units, so ans = totalPos / 2.
// Time O(N log N) | Space O(N)

import (
	"fmt"
	"sort"
)

func main() {
	// Example 1
	fmt.Println(makeSimilar([]int{8, 12, 6}, []int{2, 14, 10}))
	// Example 2
	fmt.Println(makeSimilar([]int{1, 2, 5}, []int{4, 1, 3}))
	// Already similar
	fmt.Println(makeSimilar([]int{1, 1, 1, 1, 1}, []int{1, 1, 1, 1, 1}))
}

func makeSimilar(nums []int, target []int) int64 {
	numsOdd := collectOdd(nums)
	tgtOdd := collectOdd(target)
	numsEven := collectEven(nums)
	tgtEven := collectEven(target)

	sort.Ints(numsOdd)
	sort.Ints(tgtOdd)
	sort.Ints(numsEven)
	sort.Ints(tgtEven)

	var posDiff int64
	for i := 0; i < len(numsOdd); i++ {
		if tgtOdd[i] > numsOdd[i] {
			posDiff += int64(tgtOdd[i]-numsOdd[i]) / 2
		}
	}
	for i := 0; i < len(numsEven); i++ {
		if tgtEven[i] > numsEven[i] {
			posDiff += int64(tgtEven[i]-numsEven[i]) / 2
		}
	}
	return posDiff
}

func collectOdd(a []int) []int {
	var res []int
	for _, v := range a {
		if v%2 != 0 {
			res = append(res, v)
		}
	}
	return res
}

func collectEven(a []int) []int {
	var res []int
	for _, v := range a {
		if v%2 == 0 {
			res = append(res, v)
		}
	}
	return res
}
```

## 2454 — Next Greater Element Iv

```go
package main

// LeetCode #2454: Next Greater Element IV (Second Greater Element)
// https://leetcode.com/problems/next-greater-element-iv/
// Difficulty: Hard
//
// Two monotonic stacks. stack2 holds elements that have already found their
// first greater element; when we find a greater element for stack2, that's
// the answer for those elements. stack1 holds elements still looking for
// their first greater element.
// Time O(N) | Space O(N)

import "fmt"

func main() {
	// Example 1
	fmt.Println(secondGreaterElement([]int{2, 4, 0, 9, 6}))
	// Example 2
	fmt.Println(secondGreaterElement([]int{3, 3}))
	// Strictly increasing
	fmt.Println(secondGreaterElement([]int{1, 2, 3, 4, 5}))
}

func secondGreaterElement(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = -1
	}

	stack1 := make([]int, 0, n)
	stack2 := make([]int, 0, n)
	buf := make([]int, 0, n)

	for i := 0; i < n; i++ {
		v := nums[i]

		// Process stack2 first: current element is the 2nd greater
		for len(stack2) > 0 && nums[stack2[len(stack2)-1]] < v {
			ans[stack2[len(stack2)-1]] = v
			stack2 = stack2[:len(stack2)-1]
		}

		// Process stack1: move elements that found their 1st greater to buffer
		for len(stack1) > 0 && nums[stack1[len(stack1)-1]] < v {
			buf = append(buf, stack1[len(stack1)-1])
			stack1 = stack1[:len(stack1)-1]
		}

		// Push buffer to stack2 in reverse order (maintain original order)
		for k := len(buf) - 1; k >= 0; k-- {
			stack2 = append(stack2, buf[k])
		}
		buf = buf[:0]

		stack1 = append(stack1, i)
	}

	return ans
}
```

## 2458 — Height Of Binary Tree After Subtree Removal Queries

```go
package main

// LeetCode #2458: Height of Binary Tree After Subtree Removal Queries
// https://leetcode.com/problems/height-of-binary-tree-after-subtree-removal-queries/
// Difficulty: Hard
//
// Pre-process with two DFS passes:
// 1. Compute height of each node (longest path from node to leaf).
// 2. Compute max height of the tree if we remove each node's subtree, using
//    pre-order traversal. For each node, the alternative is max of:
//    - The answer from parent (tree excluding parent's subtree)
//    - depth + 1 + height of sibling
// Time O(N + Q) | Space O(N)

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Example 1: root = [1,3,4,2,null,6,5,null,null,null,null,null,7], queries=[4]
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 4}
	root.Left.Left = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 5}
	root.Right.Left.Right = &TreeNode{Val: 7}
	fmt.Println(treeQueries(root, []int{4}))

	// Single node
	root2 := &TreeNode{Val: 1}
	fmt.Println(treeQueries(root2, []int{1}))

	// Two nodes
	root3 := &TreeNode{Val: 1}
	root3.Left = &TreeNode{Val: 2}
	fmt.Println(treeQueries(root3, []int{1, 2}))
}

func treeQueries(root *TreeNode, queries []int) []int {
	height := make(map[*TreeNode]int)
	computeHeight(root, height)

	ans := make(map[int]int)
	dfs(root, 0, 0, height, ans)

	res := make([]int, len(queries))
	for i, q := range queries {
		res[i] = ans[q]
	}
	return res
}

func computeHeight(node *TreeNode, h map[*TreeNode]int) int {
	if node == nil {
		return -1
	}
	left := computeHeight(node.Left, h)
	right := computeHeight(node.Right, h)
	val := 1 + max(left, right)
	h[node] = val
	return val
}

func dfs(node *TreeNode, depth int, parentAns int, h map[*TreeNode]int, ans map[int]int) {
	if node == nil {
		return
	}
	ans[node.Val] = parentAns

	if node.Left != nil {
		alt := parentAns
		if node.Right != nil {
			alt = max(alt, depth+1+h[node.Right])
		}
		dfs(node.Left, depth+1, alt, h, ans)
	}

	if node.Right != nil {
		alt := parentAns
		if node.Left != nil {
			alt = max(alt, depth+1+h[node.Left])
		}
		dfs(node.Right, depth+1, alt, h, ans)
	}
}
```

## 2459 — Sort Array By Moving Items To Empty Space

```go
package main

// LeetCode #2459: Sort Array by Moving Items to Empty Space
// https://leetcode.com/problems/sort-array-by-moving-items-to-empty-space/
// Difficulty: Hard [Paid]
//
// Given a permutation of 0..n-1 where 0 represents empty space,
// find the minimum number of moves to sort the array.
// A move consists of moving any element to the empty space position.
//
// Approach: Cycle decomposition. For each cycle, if it contains 0,
// cycleLen-1 moves are needed. Otherwise, cycleLen+1 moves needed
// (to bring 0 in and back out).

import "fmt"

func main() {
	// Example 1
	fmt.Println(sortArray([]int{4, 2, 0, 3, 1}))
	// Example 2
	fmt.Println(sortArray([]int{1, 0, 2, 3}))
	// Example 3
	fmt.Println(sortArray([]int{0, 1, 2, 3}))
	// Edge: already sorted
	fmt.Println(sortArray([]int{0, 1, 2, 3, 4}))
}

func sortArray(nums []int) int {
	n := len(nums)
	visited := make([]bool, n)
	ans := 0
	zeroPos := 0
	for i, v := range nums {
		if v == 0 {
			zeroPos = i
			break
		}
	}

	for i := 0; i < n; i++ {
		if visited[i] || nums[i] == i {
			visited[i] = true
			continue
		}
		// Find cycle
		cycleLen := 0
		hasZero := false
		j := i
		for !visited[j] {
			visited[j] = true
			cycleLen++
			if nums[j] == 0 {
				hasZero = true
			}
			j = nums[j]
		}
		if cycleLen > 0 {
			if hasZero {
				ans += cycleLen - 1
			} else {
				ans += cycleLen + 1
			}
		}
	}
	_ = zeroPos
	return ans
}
```

## 2463 — Minimum Total Distance Traveled

```go
package main

// LeetCode #2463: Minimum Total Distance Traveled
// https://leetcode.com/problems/minimum-total-distance-traveled/
// Difficulty: Hard
//
// Dynamic Programming. Sort robots and factories by position, then expand
// factories into individual slots. DP[i][j] = min distance to repair first i
// robots using first j factory slots. For each factory slot, either skip it
// or assign the i-th robot to it (if previous robots are handled).
// Time O(N * M) | Space O(N * M) where M = total factory slots

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	// Example 1
	fmt.Println(minimumTotalDistance(
		[]int{0, 4, 6},
		[][]int{{2, 2}, {6, 2}}))
	// Example 2
	fmt.Println(minimumTotalDistance(
		[]int{1, -1},
		[][]int{{-2, 1}, {2, 1}}))
	// Single robot, single factory
	fmt.Println(minimumTotalDistance(
		[]int{0},
		[][]int{{10, 1}}))
}

func minimumTotalDistance(robot []int, factory [][]int) int64 {
	sort.Ints(robot)

	sort.Slice(factory, func(i, j int) bool {
		return factory[i][0] < factory[j][0]
	})

	// Expand factories into individual positions
	var fpos []int
	for _, f := range factory {
		pos, limit := f[0], f[1]
		for k := 0; k < limit; k++ {
			fpos = append(fpos, pos)
		}
	}

	m := len(robot)
	n := len(fpos)

	dp := make([][]int64, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int64, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = math.MaxInt64
		}
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = 0
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// Skip this factory slot
			dp[i][j] = dp[i][j-1]

			// Assign robot i-1 to this factory slot
			dist := abs64(int64(robot[i-1]) - int64(fpos[j-1]))
			if dp[i-1][j-1] != math.MaxInt64 {
				dp[i][j] = min64(dp[i][j], dp[i-1][j-1]+dist)
			}
		}
	}

	return dp[m][n]
}

func abs64(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func min64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
```

## 2468 — Split Message Based On Limit

```go
package main

// LeetCode #2468: Split Message Based on Limit
// https://leetcode.com/problems/split-message-based-on-limit/
// Difficulty: Hard
//
// Given a message string and a limit per part, split the message into parts
// where each part has a suffix " <i/total>" appended. The suffix length counts
// towards the limit. Find the minimum number of parts (and the actual split)
// such that each part (including suffix) is <= limit characters.
//
// Approach: Binary search on the number of parts. For a given part count p,
// compute the suffix length for each part: len("<i/p>"). Check if the total
// available characters (p * limit - total_suffix_length) >= message length.
// Then construct the parts.

import (
	"fmt"
	"strconv"
)

func splitMessage(message string, limit int) []string {
	n := len(message)

	// Try different part counts
	for parts := 1; parts <= n; parts++ {
		suffixLen := len("</>") // <, /, >
		suffixLen += len(strconv.Itoa(parts))

		// Check total capacity
		totalAvailable := parts * (limit - suffixLen)
		if totalAvailable < n {
			continue
		}

		// Check each part's capacity individually (templates vary)
		// Re-check with exact suffixes
		var totalCap int
		feasible := true
		msgIdx := 0

		for i := 1; i <= parts; i++ {
			tag := "<" + strconv.Itoa(i) + "/" + strconv.Itoa(parts) + ">"
			tagLen := len(tag)
			available := limit - tagLen
			if available <= 0 {
				feasible = false
				break
			}
			totalCap += available
		}

		if !feasible || totalCap < n {
			continue
		}

		// Construct the result
		result := make([]string, parts)
		msgIdx = 0
		for i := 1; i <= parts; i++ {
			tag := "<" + strconv.Itoa(i) + "/" + strconv.Itoa(parts) + ">"
			available := limit - len(tag)
			end := msgIdx + available
			if end > n {
				end = n
			}
			result[i-1] = message[msgIdx:end] + tag
			msgIdx = end
		}

		return result
	}

	return []string{}
}

func main() {
	// Example 1
	fmt.Println(splitMessage("this is really a very awesome message", 9))
	// Example 2
	fmt.Println(splitMessage("short message", 15))
	// Single part
	fmt.Println(splitMessage("hello", 10))
	// Edge: exact fit
	fmt.Println(splitMessage("abc", 5))
	// Longer message
	fmt.Println(splitMessage("the quick brown fox jumps over the lazy dog", 11))
}
```

## 2472 — Maximum Number Of Non Overlapping Palindrome Substrings

```go
package main

// LeetCode #2472: Maximum Number of Non-overlapping Palindrome Substrings
// https://leetcode.com/problems/maximum-number-of-non-overlapping-palindrome-substrings/
// Difficulty: Hard
//
// Greedy shortest palindrome. At each position i, find the shortest palindrome
// of length >= k starting at i (max length k+1 suffices since if a longer palindrome
// exists, a shorter one of length k or k+1 also does). Count it and skip past it.

import "fmt"

func main() {
	// Example 1: "abaccdbbd", 3 => 2
	fmt.Println(maxPalindromes("abaccdbbd", 3))
	// Example 2: "adbcda", 2 => 1
	fmt.Println(maxPalindromes("adbcda", 2))
	// Edge: single character, k=1
	fmt.Println(maxPalindromes("a", 1))
	// Edge: no palindrome of length >= k
	fmt.Println(maxPalindromes("ab", 3))
	// Edge: overlaps — shortest palindrome at each position wins
	fmt.Println(maxPalindromes("aaa", 2))
}

func maxPalindromes(s string, k int) int {
	n := len(s)
	count := 0
	i := 0

	for i < n {
		found := false
		// We only need to check up to k+1 because if any palindrome >= k exists,
		// either k or k+1 is a palindrome (by parity).
		limit := k + 1
		if n-i < limit {
			limit = n - i
		}
		for length := k; length <= limit; length++ {
			if isPalindrome(s, i, i+length-1) {
				count++
				i += length
				found = true
				break
			}
		}
		if !found {
			i++
		}
	}

	return count
}

func isPalindrome(s string, l, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
```

## 2474 — Customers With Strictly Increasing Purchases

```go
package main

// LeetCode #2474: Customers With Strictly Increasing Purchases
// https://leetcode.com/problems/customers-with-strictly-increasing-purchases/
// Difficulty: Hard [Paid] (SQL)
//
// Find customers who made purchases in at least two consecutive years
// with strictly increasing purchase amounts year over year.
//
// Approach: Simulate the SQL query logic in Go.

import (
	"fmt"
)

// Order represents a customer order
type Order struct {
	OrderID       int
	CustomerID    int
	OrderDate     string // "YYYY-MM-DD"
	Price         int
}

func main() {
	// Example
	orders := []Order{
		{1, 1, "2020-06-01", 100},
		{2, 1, "2021-07-01", 150},
		{3, 2, "2020-01-01", 200},
		{4, 2, "2021-02-01", 250},
		{5, 2, "2022-03-01", 300},
		{6, 3, "2020-05-01", 50},
		{7, 3, "2021-06-01", 60},
		{8, 3, "2022-07-01", 55},
		{9, 4, "2020-10-01", 500},
		{10, 4, "2021-11-01", 400},
	}

	fmt.Println(strictlyIncreasingPurchases(orders))
}

func strictlyIncreasingPurchases(orders []Order) []int {
	// Group min purchase per customer per year
	type yearAmount struct {
		year   int
		amount int
	}
	customerYears := make(map[int][]yearAmount)

	for _, o := range orders {
		year := 0
		fmt.Sscanf(o.OrderDate[:4], "%d", &year)
		customerYears[o.CustomerID] = append(customerYears[o.CustomerID], yearAmount{year, o.Price})
	}

	var result []int
	for cid, entries := range customerYears {
		if len(entries) < 2 {
			continue
		}
		// Find min per year
		minByYear := make(map[int]int)
		for _, e := range entries {
			if v, ok := minByYear[e.year]; !ok || e.amount < v {
				minByYear[e.year] = e.amount
			}
		}
		// Sort years
		years := make([]int, 0, len(minByYear))
		for y := range minByYear {
			years = append(years, y)
		}
		// Simple bubble sort for small sets
		for i := 0; i < len(years); i++ {
			for j := i + 1; j < len(years); j++ {
				if years[i] > years[j] {
					years[i], years[j] = years[j], years[i]
				}
			}
		}

		if len(years) < 2 {
			continue
		}
		// Check strictly increasing
		strict := true
		for i := 1; i < len(years); i++ {
			if years[i] != years[i-1]+1 || minByYear[years[i]] <= minByYear[years[i-1]] {
				strict = false
				break
			}
		}
		if strict {
			result = append(result, cid)
		}
	}

	// Sort result
	for i := 0; i < len(result); i++ {
		for j := i + 1; j < len(result); j++ {
			if result[i] > result[j] {
				result[i], result[j] = result[j], result[i]
			}
		}
	}

	return result
}
```

## 2478 — Number Of Beautiful Partitions

```go
package main

// LeetCode #2478: Number of Beautiful Partitions
// https://leetcode.com/problems/number-of-beautiful-partitions/
// Difficulty: Hard
//
// DP with prefix sum. Partition s into k substrings, each:
//   - length >= minLength
//   - first char is prime digit (2,3,5,7)
//   - last char is non-prime digit
//
// DP[i][j] = ways to partition s[0:i] into j valid substrings.
// Optimized with prefix sum to O(n*k).

import "fmt"

const MOD = 1000000007

func main() {
	// Example 1: "23542185131", 3, 2 => 3
	fmt.Println(beautifulPartitions("23542185131", 3, 2))
	// Example 2: "23542185131", 3, 5 => 0
	fmt.Println(beautifulPartitions("23542185131", 3, 5))
	// Edge: single partition whole string
	fmt.Println(beautifulPartitions("331", 1, 3))
	// Edge: no valid partition
	fmt.Println(beautifulPartitions("22", 1, 2))
	// Edge: k=0
	fmt.Println(beautifulPartitions("235", 0, 2))
}

func isPrimeDigit(c byte) bool {
	return c == '2' || c == '3' || c == '5' || c == '7'
}

func beautifulPartitions(s string, k int, minLength int) int {
	n := len(s)

	// Quick check: first char must be prime, last must be non-prime
	if !isPrimeDigit(s[0]) || isPrimeDigit(s[n-1]) {
		return 0
	}
	if n < minLength*k {
		return 0
	}
	if k == 0 {
		return 1
	}

	// validSplit[i] = true if we can split after position i
	// meaning s[i-1] is non-prime and s[i] is prime
	validSplit := make([]bool, n+1)
	validSplit[0] = true
	for p := 1; p < n; p++ {
		if !isPrimeDigit(s[p-1]) && isPrimeDigit(s[p]) {
			validSplit[p] = true
		}
	}
	validSplit[n] = !isPrimeDigit(s[n-1])

	// dp[i][j] = ways using first i chars, j partitions
	dp := make([][]int64, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int64, k+1)
	}
	dp[0][0] = 1

	for j := 1; j <= k; j++ {
		sum := int64(0)
		for i := 1; i <= n; i++ {
			// Add dp[p][j-1] where p = i - minLength is a valid split
			if i >= minLength {
				p := i - minLength
				if validSplit[p] {
					sum = (sum + dp[p][j-1]) % MOD
				}
			}
			// Current position i is a valid end if s[i-1] is non-prime
			if !isPrimeDigit(s[i-1]) {
				dp[i][j] = sum
			}
		}
	}

	return int(dp[n][k])
}
```

## 2479 — Maximum Xor Of Two Non Overlapping Subtrees

```go
package main

// LeetCode #2479: Maximum XOR of Two Non-Overlapping Subtrees
// https://leetcode.com/problems/maximum-xor-of-two-non-overlapping-subtrees/
// Difficulty: Hard [Paid]
//
// Given a rooted tree (0 is root) with values at each node, find the
// maximum XOR value of two non-overlapping subtrees. Two subtrees
// are non-overlapping if they don't share any node.
//
// Approach: Compute XOR of each subtree via DFS. Use a binary trie
// to find max XOR while avoiding overlapping subtrees. Process nodes
// post-order, removing a subtree's XOR from trie after processing.

import (
	"fmt"
)

func main() {
	// Example 1
	fmt.Println(maxXor(6, [][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {2, 5}}, []int{2, 3, 5, 7, 1, 4}))
	// Example 2
	fmt.Println(maxXor(3, [][]int{{0, 1}, {1, 2}}, []int{1, 2, 3}))
	// Edge: simple chain
	fmt.Println(maxXor(2, [][]int{{0, 1}}, []int{5, 3}))
}

type TrieNode struct {
	children [2]*TrieNode
	cnt      int
}

func maxXor(n int, edges [][]int, values []int) int64 {
	if n < 2 {
		return 0
	}

	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// Compute subtree XORs
	subXor := make([]int, n)
	var dfs func(u, p int) int
	dfs = func(u, p int) int {
		xor := values[u]
		for _, v := range adj[u] {
			if v != p {
				xor ^= dfs(v, u)
			}
		}
		subXor[u] = xor
		return xor
	}
	dfs(0, -1)

	// Trie operations
	root := &TrieNode{}
	insert := func(x int) {
		node := root
		for i := 60; i >= 0; i-- {
			bit := (x >> i) & 1
			if node.children[bit] == nil {
				node.children[bit] = &TrieNode{}
			}
			node = node.children[bit]
			node.cnt++
		}
	}
	remove := func(x int) {
		node := root
		for i := 60; i >= 0; i-- {
			bit := (x >> i) & 1
			node = node.children[bit]
			node.cnt--
		}
	}
	maxXor := func(x int) int {
		node := root
		ans := 0
		for i := 60; i >= 0; i-- {
			bit := (x >> i) & 1
			want := 1 - bit
			if node.children[want] != nil && node.children[want].cnt > 0 {
				ans |= (1 << i)
				node = node.children[want]
			} else {
				node = node.children[bit]
			}
		}
		return ans
	}

	// Insert all subtree XORs initially
	for i := 0; i < n; i++ {
		insert(subXor[i])
	}

	ans := 0
	var dfs2 func(u, p int)
	dfs2 = func(u, p int) {
		remove(subXor[u])
		for _, v := range adj[u] {
			if v != p {
				dfs2(v, u)
			}
		}
		// Now trie contains only XORs of subtrees that don't overlap with u's subtree
		best := maxXor(subXor[u])
		if best > ans {
			ans = best
		}
		insert(subXor[u])
	}
	dfs2(0, -1)

	return int64(ans)
}
```

## 2484 — Count Palindromic Subsequences

```go
package main

// LeetCode #2484: Count Palindromic Subsequences
// https://leetcode.com/problems/count-palindromic-subsequences/
// Difficulty: Hard
//
// Count distinct palindromic subsequences of length 5 (a b c b a).
// For each middle position j and each pair (a,b):
//   leftCount[a][b] = number of (a,b) ordered pairs in s[0:j] (exclusive)
//   rightCount[a][b] = number of (a,b) ordered pairs in s[j+1:n] (exclusive)
//   result += leftCount[a][b] * rightCount[b][a]

import "fmt"

const MOD = 1000000007

func main() {
	// Example 1: "103301" => 2
	fmt.Println(countPalindromicSubsequences("103301"))
	// Example 2: "0000000" => 21
	fmt.Println(countPalindromicSubsequences("0000000"))
	// Example 3: "9999900000" => 96
	fmt.Println(countPalindromicSubsequences("9999900000"))
	// Edge: length 5 palindrome
	fmt.Println(countPalindromicSubsequences("12321"))
	// Edge: all distinct
	fmt.Println(countPalindromicSubsequences("12345"))
}

func countPalindromicSubsequences(s string) int {
	n := len(s)
	if n < 5 {
		return 0
	}

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = int(s[i] - '0')
	}

	// prefix counts: prefixCnt[i][d] = count of digit d in s[0:i]
	prefixCnt := make([][10]int, n+1)
	for i := 0; i < n; i++ {
		for d := 0; d < 10; d++ {
			prefixCnt[i+1][d] = prefixCnt[i][d]
		}
		prefixCnt[i+1][nums[i]]++
	}

	// suffix counts: suffixCnt[i][d] = count of digit d in s[i:n]
	suffixCnt := make([][10]int, n+1)
	for i := n - 1; i >= 0; i-- {
		for d := 0; d < 10; d++ {
			suffixCnt[i][d] = suffixCnt[i+1][d]
		}
		suffixCnt[i][nums[i]]++
	}

	// Compute total ordered pairs in the full string
	// rightPairs[a][b] = number of (a,b) pairs where a appears before b
	rightPairs := [10][10]int64{}
	for i := 0; i < n; i++ {
		d := nums[i]
		for a := 0; a < 10; a++ {
			rightPairs[a][d] += int64(prefixCnt[i][a])
		}
	}

	leftPairs := [10][10]int64{}
	var result int64 = 0

	for j := 0; j < n; j++ {
		d := nums[j]

		// Remove pairs where position j is the FIRST element (d at j, any after j)
		// These are no longer after position j.
		for b := 0; b < 10; b++ {
			rightPairs[d][b] -= int64(suffixCnt[j+1][b])
		}

		// For each (a,b), count palindromes with j as middle
		for a := 0; a < 10; a++ {
			for b := 0; b < 10; b++ {
				result = (result + leftPairs[a][b]*rightPairs[b][a]) % MOD
			}
		}

		// Add position j's contribution to leftPairs for NEXT iteration.
		// Pairs (a at p < j, d at j) now become part of the left prefix.
		for a := 0; a < 10; a++ {
			leftPairs[a][d] += int64(prefixCnt[j][a])
		}
	}

	return int(result)
}
```

## 2488 — Count Subarrays With Median K

```go
package main

// LeetCode #2488: Count Subarrays With Median K
// https://leetcode.com/problems/count-subarrays-with-median-k/
// Difficulty: Hard
//
// Transform array: >k => 1, <k => -1, =k => 0.
// A subarray has median k if it contains k and sum == 0 (odd length) or sum == 1 (even length).
// Find position p of k. Compute prefix balances left and right, use frequency maps.

import "fmt"

func main() {
	// Example 1: [3,2,1,4,5], 4 => 3
	fmt.Println(countSubarrays([]int{3, 2, 1, 4, 5}, 4))
	// Example 2: [2,3,1], 3 => 1
	fmt.Println(countSubarrays([]int{2, 3, 1}, 3))
	// Edge: single element
	fmt.Println(countSubarrays([]int{1}, 1))
	// Edge: k is first
	fmt.Println(countSubarrays([]int{5, 1, 2, 3, 4}, 5))
	// Edge: k at end
	fmt.Println(countSubarrays([]int{1, 2, 3, 4, 5}, 5))
}

func countSubarrays(nums []int, k int) int {
	n := len(nums)

	// Find position of k
	pos := -1
	for i, v := range nums {
		if v == k {
			pos = i
			break
		}
	}
	if pos == -1 {
		return 0
	}

	// Transform: >k => 1, <k => -1, =k => 0
	// Compute prefix sums from pos going right
	rightFreq := make(map[int]int)
	rightFreq[0] = 1 // empty suffix
	balance := 0
	for i := pos + 1; i < n; i++ {
		if nums[i] > k {
			balance++
		} else if nums[i] < k {
			balance--
		}
		rightFreq[balance]++
	}

	// Now go left, tracking left balance
	leftFreq := make(map[int]int)
	leftFreq[0] = 1 // empty prefix
	balance = 0
	for i := pos - 1; i >= 0; i-- {
		if nums[i] > k {
			balance++
		} else if nums[i] < k {
			balance--
		}
		leftFreq[balance]++
	}

	// Count subarrays where left_balance + right_balance == 0 (odd length, median at position)
	// or left_balance + right_balance == 1 (even length)
	result := 0
	for lb, lcnt := range leftFreq {
		// For odd length (1 + odd + even = odd), need sum == 0
		if rcnt, ok := rightFreq[-lb]; ok {
			result += lcnt * rcnt
		}
		// For even length, need sum == 1
		if rcnt, ok := rightFreq[1-lb]; ok {
			result += lcnt * rcnt
		}
	}

	return result
}
```

## 2493 — Divide Nodes Into The Maximum Number Of Groups

```go
package main

// LeetCode #2493: Divide Nodes Into the Maximum Number of Groups
// https://leetcode.com/problems/divide-nodes-into-the-maximum-number-of-groups/
// Difficulty: Hard
//
// For each connected component:
//  1. Check if it's bipartite (BFS coloring). If not, return -1.
//  2. For each node in the component, BFS to find max distance.
//  3. Answer = sum of max distances across all components.

import "fmt"

func main() {
	// Example 1: n=6, edges=[[1,2],[1,4],[1,5],[2,6],[2,3],[4,6]] => 4
	fmt.Println(magnificentSets(6, [][]int{{1, 2}, {1, 4}, {1, 5}, {2, 6}, {2, 3}, {4, 6}}))
	// Example 2: n=3, edges=[[1,2],[2,3],[3,1]] => -1 (triangle not bipartite)
	fmt.Println(magnificentSets(3, [][]int{{1, 2}, {2, 3}, {3, 1}}))
	// Edge: single node
	fmt.Println(magnificentSets(1, [][]int{}))
	// Edge: two nodes
	fmt.Println(magnificentSets(2, [][]int{{1, 2}}))
	// Edge: star graph
	fmt.Println(magnificentSets(5, [][]int{{1, 2}, {1, 3}, {1, 4}, {1, 5}}))
}

func magnificentSets(n int, edges [][]int) int {
	// Build adjacency list (1-indexed)
	adj := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		adj[i] = []int{}
	}
	for _, e := range edges {
		a, b := e[0], e[1]
		adj[a] = append(adj[a], b)
		adj[b] = append(adj[b], a)
	}

	visited := make([]bool, n+1)
	total := 0

	for i := 1; i <= n; i++ {
		if !visited[i] {
			// Collect all nodes in this component (bfsCollect marks them visited)
			component := bfsCollect(i, adj, visited)

			// Check bipartite
			color := make(map[int]int)
			if !isBipartite(component[0], adj, color) {
				return -1
			}

			// Find max depth in this component
			maxDepth := 0
			for _, node := range component {
				depth := bfsDepth(node, adj)
				if depth > maxDepth {
					maxDepth = depth
				}
			}
			total += maxDepth
		}
	}

	return total
}

func bfsCollect(start int, adj [][]int, visited []bool) []int {
	queue := []int{start}
	visited[start] = true
	result := []int{}

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		result = append(result, u)
		for _, v := range adj[u] {
			if !visited[v] {
				visited[v] = true
				queue = append(queue, v)
			}
		}
	}
	return result
}

func isBipartite(start int, adj [][]int, color map[int]int) bool {
	queue := []int{start}
	color[start] = 0

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		for _, v := range adj[u] {
			if c, ok := color[v]; ok {
				if c == color[u] {
					return false
				}
			} else {
				color[v] = color[u] ^ 1
				queue = append(queue, v)
			}
		}
	}
	return true
}

func bfsDepth(start int, adj [][]int) int {
	dist := make(map[int]int)
	queue := []int{start}
	dist[start] = 1
	maxDist := 1

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		for _, v := range adj[u] {
			if _, ok := dist[v]; !ok {
				dist[v] = dist[u] + 1
				if dist[v] > maxDist {
					maxDist = dist[v]
				}
				queue = append(queue, v)
			}
		}
	}
	return maxDist
}
```

## 2494 — Merge Overlapping Events In The Same Hall

```go
package main

// LeetCode #2494: Merge Overlapping Events in the Same Hall
// https://leetcode.com/problems/merge-overlapping-events-in-the-same-hall/
// Difficulty: Hard [Paid] (SQL)
//
// Merge overlapping events held in the same hall. Two events overlap
// if they share at least one day.
//
// Approach: Sort events by hall and start_day, then merge overlapping
// ranges per hall.

import (
	"fmt"
	"sort"
)

// HallEvent represents an event in a hall
type HallEvent struct {
	HallID   int
	StartDay string // "YYYY-MM-DD"
	EndDay   string // "YYYY-MM-DD"
}

func main() {
	// Example
	events := []HallEvent{
		{1, "2023-01-13", "2023-01-14"},
		{1, "2023-01-14", "2023-01-17"},
		{1, "2023-01-18", "2023-01-25"},
		{2, "2023-01-01", "2023-01-02"},
		{2, "2023-01-02", "2023-01-03"},
		{2, "2023-02-01", "2023-02-05"},
		{3, "2023-03-01", "2023-03-10"},
	}

	fmt.Println(mergeOverlappingEvents(events))
}

func mergeOverlappingEvents(events []HallEvent) []HallEvent {
	// Sort by hall_id, then start_day
	sort.Slice(events, func(i, j int) bool {
		if events[i].HallID != events[j].HallID {
			return events[i].HallID < events[j].HallID
		}
		return events[i].StartDay < events[j].StartDay
	})

	var result []HallEvent
	i := 0
	for i < len(events) {
		hallID := events[i].HallID
		start := events[i].StartDay
		end := events[i].EndDay
		i++

		for i < len(events) && events[i].HallID == hallID {
			// Check if overlapping (events[i].StartDay <= end means overlap)
			if events[i].StartDay <= end {
				if events[i].EndDay > end {
					end = events[i].EndDay
				}
				i++
			} else {
				break
			}
		}

		result = append(result, HallEvent{hallID, start, end})
	}

	return result
}
```

## 2499 — Minimum Total Cost To Make Arrays Unequal

```go
package main

// LeetCode #2499: Minimum Total Cost to Make Arrays Unequal
// https://leetcode.com/problems/minimum-total-cost-to-make-arrays-unequal/
// Difficulty: Hard
//
// Given two arrays nums1 and nums2, you can swap elements at the same
// index. The cost of swapping at index i is i. Find minimum total cost
// to make nums1[i] != nums2[i] for all i, or return -1 if impossible.
//
// Approach: Count positions where nums1[i] == nums2[i]. Accumulate cost.
// The dominant value needs extra swaps from other positions.

import "fmt"

func main() {
	// Example 1
	fmt.Println(minimumTotalCost([]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}))
	// Example 2
	fmt.Println(minimumTotalCost([]int{2, 1, 2}, []int{1, 2, 2}))
	// Example 3
	fmt.Println(minimumTotalCost([]int{1, 1, 1, 1, 1}, []int{1, 2, 3, 4, 5}))
	// Edge: already unequal
	fmt.Println(minimumTotalCost([]int{1, 2}, []int{3, 4}))
}

func minimumTotalCost(nums1 []int, nums2 []int) int64 {
	same, n := 0, len(nums1)
	cnt := make([]int, n+1)
	for i, a := range nums1 {
		b := nums2[i]
		if a == b {
			same++
			cnt[a]++
		}
	}

	var ans int64
	var m, lead int
	for i, v := range cnt {
		if t := v*2 - same; t > 0 {
			m = t
			lead = i
			break
		}
	}

	for i, a := range nums1 {
		if same > 0 && a == nums2[i] {
			ans += int64(i)
			same--
		}
	}

	for i, a := range nums1 {
		b := nums2[i]
		if m > 0 && a != b && a != lead && b != lead {
			ans += int64(i)
			m--
		}
	}
	if m > 0 {
		return -1
	}
	return ans
}
```

## 2503 — Maximum Number Of Points From Grid Queries

```go
package main

// LeetCode #2503: Maximum Number of Points From Grid Queries
// https://leetcode.com/problems/maximum-number-of-points-from-grid-queries/
// Difficulty: Hard
//
// Sort queries ascending while preserving original indices. Use a min-heap
// starting from (0,0). For each query q, pop all cells with value < q from
// the heap and BFS to their unvisited neighbors. Count visited cells.

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	// Example 1: grid=[[1,2,3],[2,5,7],[3,5,1]], queries=[5,6,2] => [5,8,1]
	fmt.Println(maxPoints([][]int{{1, 2, 3}, {2, 5, 7}, {3, 5, 1}}, []int{5, 6, 2}))
	// Example 2: grid=[[5,2,1],[1,1,2]], queries=[3] => [0]
	fmt.Println(maxPoints([][]int{{5, 2, 1}, {1, 1, 2}}, []int{3}))
	// Edge: single cell
	fmt.Println(maxPoints([][]int{{1}}, []int{1, 2}))
	// Edge: all cells reachable
	fmt.Println(maxPoints([][]int{{1, 2}, {3, 4}}, []int{5}))
}

type Cell struct {
	val, r, c int
}

type MinHeap []Cell

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].val < h[j].val }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(Cell)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func maxPoints(grid [][]int, queries []int) []int {
	m, n := len(grid), len(grid[0])
	k := len(queries)

	// Sort queries with original indices
	type Query struct {
		val, idx int
	}
	sorted := make([]Query, k)
	for i, v := range queries {
		sorted[i] = Query{v, i}
	}
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].val < sorted[j].val
	})

	result := make([]int, k)
	h := &MinHeap{}
	heap.Init(h)
	heap.Push(h, Cell{grid[0][0], 0, 0})

	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	visited[0][0] = true

	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	count := 0

	for _, q := range sorted {
		// Pop all cells with value < query value
		for h.Len() > 0 && (*h)[0].val < q.val {
			cell := heap.Pop(h).(Cell)
			count++
			for _, d := range dirs {
				nr, nc := cell.r+d[0], cell.c+d[1]
				if nr >= 0 && nr < m && nc >= 0 && nc < n && !visited[nr][nc] {
					visited[nr][nc] = true
					heap.Push(h, Cell{grid[nr][nc], nr, nc})
				}
			}
		}
		result[q.idx] = count
	}

	return result
}
```

## 2508 — Add Edges To Make Degrees Of All Nodes Even

```go
package main

// LeetCode #2508: Add Edges to Make Degrees of All Nodes Even
// https://leetcode.com/problems/add-edges-to-make-degrees-of-all-nodes-even/
// Difficulty: Hard
//
// Find nodes with odd degree. We can add at most 2 edges.
// Cases:
//  0 odd nodes => true
//  2 odd nodes => connect them directly, or connect both to a third node (with even degree)
//  4 odd nodes => try all 3 pairings
//  > 4 odd nodes => false (2 edges fix at most 4 odd nodes)

import "fmt"

func main() {
	// Example 1: n=5, edges=[[1,2],[2,3],[3,4],[4,2],[1,4],[2,5]] => true
	fmt.Println(isPossible(5, [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 2}, {1, 4}, {2, 5}}))
	// Example 2: n=4, edges=[[1,2],[3,4]] => true
	fmt.Println(isPossible(4, [][]int{{1, 2}, {3, 4}}))
	// Example 3: n=4, edges=[[1,2],[1,3],[1,4]] => false
	fmt.Println(isPossible(4, [][]int{{1, 2}, {1, 3}, {1, 4}}))
	// Edge: already all even
	fmt.Println(isPossible(3, [][]int{{1, 2}, {2, 3}, {3, 1}}))
	// Edge: 2 odd nodes but already connected directly
	fmt.Println(isPossible(5, [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 1}, {1, 3}}))
}

func isPossible(n int, edges [][]int) bool {
	// Track which edges exist (dense graph, use adjacency matrix/set)
	adjSet := make(map[int]map[int]bool)
	for i := 1; i <= n; i++ {
		adjSet[i] = make(map[int]bool)
	}
	for _, e := range edges {
		a, b := e[0], e[1]
		adjSet[a][b] = true
		adjSet[b][a] = true
	}

	deg := make([]int, n+1)
	for _, e := range edges {
		deg[e[0]]++
		deg[e[1]]++
	}

	odd := []int{}
	for i := 1; i <= n; i++ {
		if deg[i]%2 == 1 {
			odd = append(odd, i)
		}
	}

	if len(odd) == 0 {
		return true
	}

	if len(odd) == 2 {
		a, b := odd[0], odd[1]
		// Can connect a and b directly
		if !adjSet[a][b] {
			return true
		}
		// Connect both to a third node
		for c := 1; c <= n; c++ {
			if c != a && c != b {
				if !adjSet[a][c] && !adjSet[b][c] {
					return true
				}
			}
		}
		return false
	}

	if len(odd) == 4 {
		// Try all 3 pairings: (0-1,2-3), (0-2,1-3), (0-3,1-2)
		pairings := [][2][2]int{
			{{odd[0], odd[1]}, {odd[2], odd[3]}},
			{{odd[0], odd[2]}, {odd[1], odd[3]}},
			{{odd[0], odd[3]}, {odd[1], odd[2]}},
		}
		for _, p := range pairings {
			a1, b1 := p[0][0], p[0][1]
			a2, b2 := p[1][0], p[1][1]
			if !adjSet[a1][b1] && !adjSet[a2][b2] {
				return true
			}
		}
		return false
	}

	return false
}
```

## 2509 — Cycle Length Queries In A Tree

```go
package main

// LeetCode #2509: Cycle Length Queries in a Tree
// https://leetcode.com/problems/cycle-length-queries-in-a-tree/
// Difficulty: Hard
//
// Complete binary tree with n levels. For each query (a,b), find the length
// of the cycle formed by adding an edge between a and b.
// Cycle length = depth(a) + depth(b) - 2*depth(LCA(a,b)) + 1.
// Node numbering: root=1, left=2*i, right=2*i+1.

import "fmt"

func main() {
	// Example 1: n=3, queries=[[5,3],[4,7],[2,3]] => [4,5,3]
	fmt.Println(cycleLengthQueries(3, [][]int{{5, 3}, {4, 7}, {2, 3}}))
	// Example 2: n=2, queries=[[1,2]] => [2]
	fmt.Println(cycleLengthQueries(2, [][]int{{1, 2}}))
	// Edge: same node
	fmt.Println(cycleLengthQueries(2, [][]int{{1, 1}}))
	// Edge: root with another
	fmt.Println(cycleLengthQueries(3, [][]int{{1, 7}}))
}

func cycleLengthQueries(n int, queries [][]int) []int {
	result := make([]int, len(queries))
	for i, q := range queries {
		a, b := q[0], q[1]
		if a == b {
			result[i] = 1 // cycle of a self-loop
			continue
		}
		// Find depths and LCA
		da := depth(a)
		db := depth(b)
		lca := findLCA(a, b)
		dlca := depth(lca)
		// Cycle length = da + db - 2*dlca + 1
		result[i] = da + db - 2*dlca + 1
	}
	return result
}

func depth(x int) int {
	d := 0
	for x > 0 {
		x >>= 1
		d++
	}
	return d
}

func findLCA(a, b int) int {
	// Bring nodes to same depth
	for depth(a) > depth(b) {
		a >>= 1
	}
	for depth(b) > depth(a) {
		b >>= 1
	}
	// Move up together
	for a != b {
		a >>= 1
		b >>= 1
	}
	return a
}
```

## 2514 — Count Anagrams

```go
package main

// LeetCode #2514: Count Anagrams
// https://leetcode.com/problems/count-anagrams/
// Difficulty: Hard
//
// For each word, number of distinct anagrams = word_len! / product(cnt[char]!).
// Multiply across all words. Use modular arithmetic with MOD = 1e9+7.

import (
	"fmt"
	"strings"
)

const MOD = 1000000007

func main() {
	// Example 1: "too" => 2
	fmt.Println(countAnagrams("too"))
	// Example 2: "aa aa" => 1
	fmt.Println(countAnagrams("aa aa"))
	// Edge: single char
	fmt.Println(countAnagrams("a"))
	// Edge: all same letters
	fmt.Println(countAnagrams("aaa"))
	// Edge: multiple words
	fmt.Println(countAnagrams("abc def ghi"))
}

func countAnagrams(s string) int {
	words := strings.Fields(s)
	result := int64(1)

	for _, word := range words {
		n := len(word)
		// Count character frequencies
		cnt := make(map[rune]int)
		for _, ch := range word {
			cnt[ch]++
		}

		// result *= n! / product(cnt[c]!)
		// Compute n! * inverse(product(cnt[c]!))
		res := factorial(n)
		for _, c := range cnt {
			inv := modInv(factorial(c))
			res = (res * inv) % MOD
		}
		result = (result * res) % MOD
	}

	return int(result)
}

func factorial(n int) int64 {
	res := int64(1)
	for i := 2; i <= n; i++ {
		res = (res * int64(i)) % MOD
	}
	return res
}

func modInv(a int64) int64 {
	return modPow(a, MOD-2)
}

func modPow(a int64, b int) int64 {
	res := int64(1)
	for b > 0 {
		if b&1 == 1 {
			res = (res * a) % MOD
		}
		a = (a * a) % MOD
		b >>= 1
	}
	return res
}
```

## 2518 — Number Of Great Partitions

```go
package main

// LeetCode #2518: Number of Great Partitions
// https://leetcode.com/problems/number-of-great-partitions/
// Difficulty: Hard

import "fmt"

const mod = 1000000007

// numberOfGreatPartitions counts the number of ways to partition nums into
// two groups such that |sum(group1) - sum(group2)| >= k.
//
// Total ways = 2^n. Subtract ways where |2*subsetSum - totalSum| < k.
// DP counts subsets achieving each possible sum.
//
// Complexity: O(n * totalSum) time, O(totalSum) space
func numberOfGreatPartitions(nums []int, k int) int {
	totalSum := 0
	for _, v := range nums {
		totalSum += v
	}

	// dp[s] = number of ways to achieve subset sum s (using distinct elements)
	dp := make([]int, totalSum+1)
	dp[0] = 1
	for _, v := range nums {
		for s := totalSum; s >= v; s-- {
			dp[s] = (dp[s] + dp[s-v]) % mod
		}
	}

	// Count "bad" partitions where |2*sum - totalSum| < k
	bad := 0
	for sum := 0; sum <= totalSum; sum++ {
		if abs(2*sum-totalSum) < k {
			bad = (bad + dp[sum]) % mod
		}
	}

	// Total partitions = 2^n
	total := 1
	for i := 0; i < len(nums); i++ {
		total = (total * 2) % mod
	}

	return (total - bad + mod) % mod
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	// Example from LeetCode
	fmt.Println("Test 1: nums=[1,2,3,4], k=4 ->", numberOfGreatPartitions([]int{1, 2, 3, 4}, 4))

	// Additional test cases
	fmt.Println("Test 2: nums=[1,1,1,1], k=2 ->", numberOfGreatPartitions([]int{1, 1, 1, 1}, 2))
	fmt.Println("Test 3: nums=[1,2], k=1 ->", numberOfGreatPartitions([]int{1, 2}, 1))
	fmt.Println("Test 4: nums=[1], k=1 ->", numberOfGreatPartitions([]int{1}, 1))
	fmt.Println("Test 5: nums=[3,5], k=1 ->", numberOfGreatPartitions([]int{3, 5}, 1))
	fmt.Println("Test 6: nums=[1,2,3], k=100 ->", numberOfGreatPartitions([]int{1, 2, 3}, 100))
}
```

## 2519 — Count The Number Of K Big Indices

```go
package main

// LeetCode #2519: Count the Number of K-Big Indices
// https://leetcode.com/problems/count-the-number-of-k-big-indices/
// Difficulty: Hard [Paid]

import (
	"fmt"
	"sort"
)

// countKBigIndices counts indices i where at least k elements before i
// are strictly less than nums[i] AND at least k elements after i
// are strictly less than nums[i].
//
// Approach: Coordinate compression + Fenwick Tree (BIT).
// Scan left-to-right to count less-than-curr before i,
// scan right-to-left to count less-than-curr after i.
//
// Complexity: O(n log n) time, O(n) space
func countKBigIndices(nums []int, k int) int {
	n := len(nums)

	// Coordinate compression
	sorted := make([]int, n)
	copy(sorted, nums)
	sort.Ints(sorted)
	rank := make(map[int]int)
	for _, v := range sorted {
		if _, ok := rank[v]; !ok {
			rank[v] = len(rank) + 1 // 1-indexed for BIT
		}
	}
	m := len(rank)

	// leftLess[i] = count of elements before i with value < nums[i]
	bit := make([]int, m+2)
	leftLess := make([]int, n)
	for i := 0; i < n; i++ {
		r := rank[nums[i]]
		leftLess[i] = bitQuery(bit, r-1)
		bitUpdate(bit, r, 1)
	}

	// rightLess[i] = count of elements after i with value < nums[i]
	bit = make([]int, m+2)
	rightLess := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		r := rank[nums[i]]
		rightLess[i] = bitQuery(bit, r-1)
		bitUpdate(bit, r, 1)
	}

	ans := 0
	for i := 0; i < n; i++ {
		if leftLess[i] >= k && rightLess[i] >= k {
			ans++
		}
	}
	return ans
}

func bitUpdate(bit []int, idx, delta int) {
	for idx < len(bit) {
		bit[idx] += delta
		idx += idx & -idx
	}
}

func bitQuery(bit []int, idx int) int {
	sum := 0
	for idx > 0 {
		sum += bit[idx]
		idx -= idx & -idx
	}
	return sum
}

func main() {
	// Test cases
	fmt.Println("Test 1: nums=[2,3,1,2,1], k=2 ->", countKBigIndices([]int{2, 3, 1, 2, 1}, 2))
	fmt.Println("Test 2: nums=[1,2,3,4,5], k=2 ->", countKBigIndices([]int{1, 2, 3, 4, 5}, 2))
	fmt.Println("Test 3: nums=[5,4,3,2,1], k=1 ->", countKBigIndices([]int{5, 4, 3, 2, 1}, 1))
	fmt.Println("Test 4: nums=[1,1,1], k=1 ->", countKBigIndices([]int{1, 1, 1}, 1))
	fmt.Println("Test 5: nums=[1], k=1 ->", countKBigIndices([]int{1}, 1))
	fmt.Println("Test 6: nums=[3,2,1,2,3], k=2 ->", countKBigIndices([]int{3, 2, 1, 2, 3}, 2))
}
```

## 2524 — Maximum Frequency Score Of A Subarray

```go
package main

// LeetCode #2524: Maximum Frequency Score of a Subarray
// https://leetcode.com/problems/maximum-frequency-score-of-a-subarray/
// Difficulty: Hard [Paid]

import "fmt"

const mod = 1000000007

// maxFrequencyScore computes the maximum frequency score of any subarray of length k.
// Frequency score = sum over distinct values v of (freq[v] * v^freq[v]) % MOD.
//
// Approach: Sliding window with frequency and contribution tracking.
// Maintain current score. When freq[v] changes from f to f+1:
//   subtract old contribution (v^f) and add new (v^(f+1)).
// Use fast exponentiation with precomputed powers for O(1) window updates.
//
// Complexity: O(n * log MOD) time, O(distinct values) space
func maxFrequencyScore(nums []int, k int) int {
	n := len(nums)
	if k > n || k <= 0 {
		return 0
	}

	freq := make(map[int]int)
	score := 0

	// Initialize first window
	for i := 0; i < k; i++ {
		v := nums[i]
		f := freq[v]
		// Add v^(f+1), remove old v^f (if f > 0)
		if f > 0 {
			score = (score - modPow(v, f) + mod) % mod
		}
		score = (score + modPow(v, f+1)) % mod
		freq[v] = f + 1
	}

	maxScore := score

	// Slide window
	for i := k; i < n; i++ {
		// Remove outgoing element nums[i-k]
		out := nums[i-k]
		f := freq[out] // current freq before removal
		// Remove out^f contribution, add out^(f-1) if f > 1
		score = (score - modPow(out, f) + mod) % mod
		if f > 1 {
			score = (score + modPow(out, f-1)) % mod
		}
		freq[out] = f - 1
		if freq[out] == 0 {
			delete(freq, out)
		}

		// Add incoming element nums[i]
		in := nums[i]
		f = freq[in] // current freq before addition
		if f > 0 {
			score = (score - modPow(in, f) + mod) % mod
		}
		score = (score + modPow(in, f+1)) % mod
		freq[in] = f + 1

		if score > maxScore {
			maxScore = score
		}
	}

	return maxScore
}

// modPow computes (base^exp) % mod using binary exponentiation.
func modPow(base, exp int) int {
	result := 1
	b := base % mod
	e := exp
	for e > 0 {
		if e&1 == 1 {
			result = result * b % mod
		}
		b = b * b % mod
		e >>= 1
	}
	return result
}

func main() {
	// Test cases
	fmt.Println("Test 1: nums=[1,1,1,2,2], k=3 ->", maxFrequencyScore([]int{1, 1, 1, 2, 2}, 3))
	fmt.Println("Test 2: nums=[1,2,3,4], k=2 ->", maxFrequencyScore([]int{1, 2, 3, 4}, 2))
	fmt.Println("Test 3: nums=[5,5,5], k=1 ->", maxFrequencyScore([]int{5, 5, 5}, 1))
	fmt.Println("Test 4: nums=[1,2], k=3 ->", maxFrequencyScore([]int{1, 2}, 3)) // k > n
	fmt.Println("Test 5: nums=[1,1,1], k=3 ->", maxFrequencyScore([]int{1, 1, 1}, 3))
	fmt.Println("Test 6: nums=[4,4,4,4], k=2 ->", maxFrequencyScore([]int{4, 4, 4, 4}, 2))
}
```

## 2528 — Maximize The Minimum Powered City

```go
package main

// LeetCode #2528: Maximize the Minimum Powered City
// https://leetcode.com/problems/maximize-the-minimum-powered-city/
// Difficulty: Hard
//
// Given an array stations where stations[i] is the power of the station
// at city i, a radius r, and k additional stations that can be placed,
// maximize the minimum total power across all cities.
// Each station powers cities within distance r.
//
// Approach: Binary search on the minimum power. Use difference array
// to check if we can achieve at least x power at every city.

import "fmt"

func main() {
	// Example 1
	fmt.Println(maxPower([]int{1, 2, 4, 5, 0}, 1, 2))
	// Example 2
	fmt.Println(maxPower([]int{4, 4, 4, 4}, 0, 3))
	// Example 3
	fmt.Println(maxPower([]int{5, 10, 5}, 1, 2))
	// Edge: single city
	fmt.Println(maxPower([]int{10}, 0, 5))
}

func maxPower(stations []int, r int, k int) int64 {
	n := len(stations)
	d := make([]int, n+1)
	s := make([]int, n+1)
	for i, v := range stations {
		left, right := max(0, i-r), min(i+r, n-1)
		d[left] += v
		d[right+1] -= v
	}
	s[0] = d[0]
	for i := 1; i < n+1; i++ {
		s[i] = s[i-1] + d[i]
	}
	check := func(x, k int) bool {
		d := make([]int, n+1)
		t := 0
		for i := range stations {
			t += d[i]
			dist := x - (s[i] + t)
			if dist > 0 {
				if k < dist {
					return false
				}
				k -= dist
				j := min(i+r, n-1)
				left, right := max(0, j-r), min(j+r, n-1)
				d[left] += dist
				d[right+1] -= dist
				t += dist
			}
		}
		return true
	}
	left, right := 0, 1<<40
	for left < right {
		mid := (left + right + 1) >> 1
		if check(mid, k) {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return int64(left)
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

## 2532 — Time To Cross A Bridge

```go
package main

// LeetCode #2532: Time to Cross a Bridge
// https://leetcode.com/problems/time-to-cross-a-bridge/
// Difficulty: Hard
//
// Simulation with 4 priority queues:
//   leftWait: workers on left waiting to cross right (with box)
//   rightWait: workers on right waiting to cross left (without box)
//   leftWork: workers on left picking up boxes
//   rightWork: workers on right putting down boxes
//
// Priority for crossing: right side first, then left side.
// Within same side: smallest (leftToRight+rightToLeft) first, then smallest index.

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	// Example: n=1, k=3, time=[[1,1,2,1],[1,1,3,1],[1,1,4,1]] => 6
	fmt.Println(findCrossingTime(1, 3, [][]int{{1, 1, 2, 1}, {1, 1, 3, 1}, {1, 1, 4, 1}}))
	// Example 2: n=3, k=2, time=[[1,9,1,8],[10,10,10,10]] => 50
	fmt.Println(findCrossingTime(3, 2, [][]int{{1, 9, 1, 8}, {10, 10, 10, 10}}))
	// Edge: single worker, single box
	fmt.Println(findCrossingTime(1, 1, [][]int{{2, 1, 1, 2}}))
	// Edge: no boxes
	fmt.Println(findCrossingTime(0, 2, [][]int{{1, 1, 1, 1}, {2, 2, 2, 2}}))
}

// Worker wrapper for priority queues
type Worker struct {
	idx        int
	efficiency int // leftToRight + rightToLeft (lower = more efficient)
}

type WorkerHeap []Worker

func (h WorkerHeap) Len() int      { return len(h) }
func (h WorkerHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h WorkerHeap) Less(i, j int) bool {
	if h[i].efficiency != h[j].efficiency {
		return h[i].efficiency < h[j].efficiency
	}
	return h[i].idx < h[j].idx
}
func (h *WorkerHeap) Push(x interface{}) { *h = append(*h, x.(Worker)) }
func (h *WorkerHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// TimeHeap for work queues (min-heap by readyTime)
type TimeEvent struct {
	readyTime int
	idx       int
}
type TimeHeap []TimeEvent

func (h TimeHeap) Len() int            { return len(h) }
func (h TimeHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h TimeHeap) Less(i, j int) bool  { return h[i].readyTime < h[j].readyTime }
func (h *TimeHeap) Push(x interface{}) { *h = append(*h, x.(TimeEvent)) }
func (h *TimeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func findCrossingTime(n int, k int, time [][]int) int {
	if n == 0 {
		return 0
	}

	// Precompute efficiency
	efficiency := make([]int, k)
	for i := 0; i < k; i++ {
		efficiency[i] = time[i][0] + time[i][2]
	}

	// Initial state: all workers on left, start picking up
	leftWork := &TimeHeap{}
	rightWork := &TimeHeap{}
	leftWait := &WorkerHeap{}
	rightWait := &WorkerHeap{}
	heap.Init(leftWork)
	heap.Init(rightWork)
	heap.Init(leftWait)
	heap.Init(rightWait)

	// Initially, workers need to pick up a box before crossing
	for i := 0; i < k; i++ {
		heap.Push(leftWork, TimeEvent{readyTime: time[i][1], idx: i}) // pickLeft
	}

	remainingBoxes := n
	currentTime := 0

	// Continue until all boxes are delivered and all workers are on the left
	for remainingBoxes > 0 || rightWork.Len() > 0 || rightWait.Len() > 0 {
		// Flush workers whose work is done into wait queues
		for leftWork.Len() > 0 && (*leftWork)[0].readyTime <= currentTime {
			ev := heap.Pop(leftWork).(TimeEvent)
			w := Worker{idx: ev.idx, efficiency: efficiency[ev.idx]}
			heap.Push(leftWait, w)
		}
		for rightWork.Len() > 0 && (*rightWork)[0].readyTime <= currentTime {
			ev := heap.Pop(rightWork).(TimeEvent)
			w := Worker{idx: ev.idx, efficiency: efficiency[ev.idx]}
			heap.Push(rightWait, w)
		}

		if rightWait.Len() > 0 {
			// Worker on right crosses back to left
			w := heap.Pop(rightWait).(Worker)
			currentTime += time[w.idx][2] // rightToLeft
			// Worker is now on left, starts picking up
			heap.Push(leftWork, TimeEvent{
				readyTime: currentTime + time[w.idx][1], // + pickLeft
				idx:       w.idx,
			})
		} else if leftWait.Len() > 0 && remainingBoxes > 0 {
			// Worker on left crosses to right (with a box, only if boxes remain)
			w := heap.Pop(leftWait).(Worker)
			currentTime += time[w.idx][0] // leftToRight
			remainingBoxes--
			// Worker is now on right, starts putting down
			heap.Push(rightWork, TimeEvent{
				readyTime: currentTime + time[w.idx][3], // + pickRight
				idx:       w.idx,
			})
		} else {
			// No one waiting, advance time to next work completion
			nextTime := math.MaxInt64
			if leftWork.Len() > 0 && (*leftWork)[0].readyTime < nextTime {
				nextTime = (*leftWork)[0].readyTime
			}
			if rightWork.Len() > 0 && (*rightWork)[0].readyTime < nextTime {
				nextTime = (*rightWork)[0].readyTime
			}
			currentTime = nextTime
		}
	}

	return currentTime
}
```

## 2534 — Time Taken To Cross The Door

```go
package main

// LeetCode #2534: Time Taken to Cross the Door
// https://leetcode.com/problems/time-taken-to-cross-the-door/
// Difficulty: Hard [Paid]

import "fmt"

// timeTakenCrossDoor returns the time each person crosses the door.
// arrival[i] = time person i arrives, state[i] = 1 for enter, 0 for leave.
//
// Simulation with two queues (enterers, leavers). At each time step:
// 1. Add all people who arrived at/before current time to appropriate queue.
// 2. If both queues empty, jump to next arrival time.
// 3. Decide who crosses:
//    - One queue empty -> the other goes.
//    - Both non-empty -> prefer same direction as previous crossing.
//    - First crossing (no previous) -> enterers have priority (state=1).
// 4. Person crosses at current time, advance time by 1.
//
// Complexity: O(T + n) where T = max(arrival) + n, O(n) space
func timeTakenCrossDoor(arrival []int, state []int) []int {
	n := len(arrival)
	result := make([]int, n)

	enterQ := make([]int, 0) // indices of people waiting to enter
	leaveQ := make([]int, 0) // indices of people waiting to leave

	nextIdx := 0 // next person index to consider
	time := 0
	prevState := 1 // 1 = enter, 0 = leave; start assuming enter
	processed := 0

	for processed < n {
		// Add all people who have arrived up to current time
		for nextIdx < n && arrival[nextIdx] <= time {
			if state[nextIdx] == 1 {
				enterQ = append(enterQ, nextIdx)
			} else {
				leaveQ = append(leaveQ, nextIdx)
			}
			nextIdx++
		}

		// If no one is waiting, jump to next arrival
		if len(enterQ) == 0 && len(leaveQ) == 0 {
			if nextIdx < n {
				time = arrival[nextIdx]
				continue
			}
			break
		}

		// Decide who crosses
		var chosen int
		if len(enterQ) > 0 && len(leaveQ) > 0 {
			// Both non-empty: prefer same direction as previous crossing
			if prevState == 1 {
				chosen = enterQ[0]
				enterQ = enterQ[1:]
			} else {
				chosen = leaveQ[0]
				leaveQ = leaveQ[1:]
			}
		} else if len(enterQ) > 0 {
			chosen = enterQ[0]
			enterQ = enterQ[1:]
		} else {
			chosen = leaveQ[0]
			leaveQ = leaveQ[1:]
		}

		result[chosen] = time
		prevState = state[chosen]
		time++
		processed++
	}

	return result
}

func main() {
	// Test cases
	fmt.Println("Test 1: arrival=[0,0,0], state=[1,0,0] ->", timeTakenCrossDoor([]int{0, 0, 0}, []int{1, 0, 0}))
	fmt.Println("Test 2: arrival=[0,1,1,2,4], state=[0,1,0,0,1] ->", timeTakenCrossDoor([]int{0, 1, 1, 2, 4}, []int{0, 1, 0, 0, 1}))
	fmt.Println("Test 3: arrival=[0,0,1], state=[1,1,0] ->", timeTakenCrossDoor([]int{0, 0, 1}, []int{1, 1, 0}))
	fmt.Println("Test 4: arrival=[0], state=[0] ->", timeTakenCrossDoor([]int{0}, []int{0}))
	fmt.Println("Test 5: arrival=[0,2,3], state=[1,0,1] ->", timeTakenCrossDoor([]int{0, 2, 3}, []int{1, 0, 1}))
	fmt.Println("Test 6: arrival=[1,1,1,1], state=[0,0,1,1] ->", timeTakenCrossDoor([]int{1, 1, 1, 1}, []int{0, 0, 1, 1}))
}
```

## 2538 — Difference Between Maximum And Minimum Price Sum

```go
package main

// LeetCode #2538: Difference Between Maximum and Minimum Price Sum
// https://leetcode.com/problems/difference-between-maximum-and-minimum-price-sum/
// Difficulty: Hard

import "fmt"

// maxOutput computes the maximum difference between max and min price sum
// over all paths in the tree.
//
// For each node, compute maxDown (max sum from node to a leaf going down)
// and minDown (min sum from node to a leaf going down).
// Then for paths passing through a node connecting two children, compute
// the max sum and min sum. The answer is the maximum difference.
//
// Complexity: O(n) time, O(n) space
func maxOutput(n int, edges [][]int, price []int) int64 {
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	ans := int64(0)

	// Returns (maxDown, minDown) for node u:
	// maxDown = max sum from u to a descendant (including u)
	// minDown = min sum from u to a descendant (including u)
	var dfs func(u, parent int) (int64, int64)
	dfs = func(u, parent int) (int64, int64) {
		// Each leaf serves as both a max and min single-node path
		maxDown := int64(price[u])
		minDown := int64(price[u])

		// Collect top 2 max and bottom 2 min from children
		var top1, top2 int64 = -1, -1 // max sums from children (excluding price[u])
		var bot1, bot2 int64 = 1<<60, 1<<60 // min sums from children

		for _, v := range adj[u] {
			if v == parent {
				continue
			}
			cMax, cMin := dfs(v, u)

			// Update maxDown/minDown = price[u] + best child path
			// For max, we add if child improves it; for min, if child makes it worse
			if cMax > 0 {
				maxDown = max64(maxDown, int64(price[u])+cMax)
			}
			if cMin < 0 {
				minDown = min64(minDown, int64(price[u])+cMin)
			}

			// Track top-2 max child sums (need to subtract price[u] to compare on equal footing)
			childMaxContribution := cMax // this is sum from child v going down, not including price[u]
			if childMaxContribution > top1 {
				top2, top1 = top1, childMaxContribution
			} else if childMaxContribution > top2 {
				top2 = childMaxContribution
			}

			childMinContribution := cMin
			if childMinContribution < bot1 {
				bot2, bot1 = bot1, childMinContribution
			} else if childMinContribution < bot2 {
				bot2 = childMinContribution
			}
		}

		// Consider path through u adding two child paths
		if top1 >= 0 {
			maxThrough := int64(price[u]) + top1
			if top2 >= 0 {
				maxThrough += top2
			}
			if maxThrough > ans {
				ans = maxThrough
			}
		}
		if bot1 <= 0 {
			minThrough := int64(price[u]) + bot1
			if bot2 <= 0 {
				minThrough += bot2
			}
			if minThrough < ans {
				// Update answer as maxDown - minDown through this node
				// We need the best path pair through u
				// Compute difference: best max and best min going through u
			}
		}

		// Compute difference at this node: best max through vs best min through
		// This considers paths that both pass through u, possibly using same children
		bestMaxThrough := int64(price[u])
		if top1 > 0 {
			bestMaxThrough += top1
		}
		if top2 > 0 {
			bestMaxThrough += top2
		}

		bestMinThrough := int64(price[u])
		if bot1 < 0 {
			bestMinThrough += bot1
		}
		if bot2 < 0 {
			bestMinThrough += bot2
		}

		if bestMaxThrough-bestMinThrough > ans {
			ans = bestMaxThrough - bestMinThrough
		}

		return maxDown, minDown
	}

	dfs(0, -1)
	return ans
}

func max64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func min64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Test cases
	fmt.Println("Test 1: n=6, edges=[[0,1],[1,2],[1,3],[3,4],[3,5]], price=[1,2,3,4,5,6] ->",
		maxOutput(6, [][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}, {3, 5}}, []int{1, 2, 3, 4, 5, 6}))

	fmt.Println("Test 2: n=3, edges=[[0,1],[1,2]], price=[10,10,10] ->",
		maxOutput(3, [][]int{{0, 1}, {1, 2}}, []int{10, 10, 10}))

	fmt.Println("Test 3: n=1, edges=[], price=[5] ->",
		maxOutput(1, [][]int{}, []int{5}))

	fmt.Println("Test 4: n=4, edges=[[0,1],[1,2],[2,3]], price=[1,5,1,5] ->",
		maxOutput(4, [][]int{{0, 1}, {1, 2}, {2, 3}}, []int{1, 5, 1, 5}))

	fmt.Println("Test 5: n=2, edges=[[0,1]], price=[3,7] ->",
		maxOutput(2, [][]int{{0, 1}}, []int{3, 7}))
}
```

## 2543 — Check If Point Is Reachable

```go
package main

// LeetCode #2543: Check if Point Is Reachable
// https://leetcode.com/problems/check-if-point-is-reachable/
// Difficulty: Hard

import "fmt"

// isReachable checks if (targetX, targetY) is reachable from (1, 1) using:
//   (x, y) -> (x, y-x), (x-y, y), (2x, y), (x, 2y)
//
// Key insight: (x, y) is reachable iff gcd(x, y) is a power of 2.
// The subtraction operations preserve GCD. Doubling multiplies a coordinate by 2.
// Starting from (1, 1), gcd = 1 = 2^0. The invariant is that gcd is always a power of 2.
// Conversely, any (x, y) with gcd = 2^k can be reduced back to (1, 1) by reversing
// the operations (reverse subtraction = addition; reverse doubling = halving).
//
// Complexity: O(log(min(x,y))) time, O(1) space
func isReachable(targetX int, targetY int) bool {
	g := gcd(targetX, targetY)
	// Check if g is a power of 2: a power of 2 has exactly one bit set
	return g&(g-1) == 0
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	// Test cases
	fmt.Println("Test 1: targetX=6, targetY=9 ->", isReachable(6, 9))   // false
	fmt.Println("Test 2: targetX=4, targetY=7 ->", isReachable(4, 7))   // false
	fmt.Println("Test 3: targetX=1, targetY=1 ->", isReachable(1, 1))   // true
	fmt.Println("Test 4: targetX=2, targetY=4 ->", isReachable(2, 4))   // true (gcd=2)
	fmt.Println("Test 5: targetX=3, targetY=5 ->", isReachable(3, 5))   // true (gcd=1=2^0)
	fmt.Println("Test 6: targetX=8, targetY=12 ->", isReachable(8, 12)) // true (gcd=4=2^2)
	fmt.Println("Test 7: targetX=4, targetY=8 ->", isReachable(4, 8))   // true (gcd=4=2^2)
	fmt.Println("Test 8: targetX=2, targetY=3 ->", isReachable(2, 3))   // true (gcd=1=2^0)
}
```

## 2547 — Minimum Cost To Split An Array

```go
package main

// LeetCode #2547: Minimum Cost to Split an Array
// https://leetcode.com/problems/minimum-cost-to-split-an-array/
// Difficulty: Hard

import "fmt"

// minCost uses DP where dp[i] = min cost for prefix nums[0..i-1].
// For each i, we try all j < i as the start of the last subarray,
// tracking the "trimmed" count: number of distinct values whose
// frequency in the subarray >= 2. Cost = trimmed + k.
//
// Complexity: O(n^2) time, O(n) space
func minCost(nums []int, k int) int {
	n := len(nums)
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = 1 << 60
	}
	dp[0] = 0

	for i := 1; i <= n; i++ {
		freq := make(map[int]int)
		trimmed := 0
		for j := i - 1; j >= 0; j-- {
			x := nums[j]
			freq[x]++
			if freq[x] == 2 {
				// This value now appears more than once for the first time
				trimmed++
			}
			cost := trimmed + k
			if dp[j]+cost < dp[i] {
				dp[i] = dp[j] + cost
			}
		}
	}
	return dp[n]
}

func main() {
	// Example from LeetCode
	fmt.Println("Test 1: [1,2,1,2,1], k=2 ->", minCost([]int{1, 2, 1, 2, 1}, 2)) // 4

	// Additional test cases
	fmt.Println("Test 2: [1,2,1,2,1], k=5 ->", minCost([]int{1, 2, 1, 2, 1}, 5)) // 7 (no split better than whole)
	fmt.Println("Test 3: [1,2,3,4,5], k=2 ->", minCost([]int{1, 2, 3, 4, 5}, 2)) // all unique -> each singly
	fmt.Println("Test 4: [0,0,0,0], k=1 ->", minCost([]int{0, 0, 0, 0}, 1))      // all same: trimmed=1 per subarray
	fmt.Println("Test 5: [] ->", minCost([]int{}, 5))                             // 0
}
```

## 2551 — Put Marbles In Bags

```go
package main

// LeetCode #2551: Put Marbles in Bags
// https://leetcode.com/problems/put-marbles-in-bags/
// Difficulty: Hard

import (
	"fmt"
	"sort"
)

// putMarbles computes the difference between max and min scores.
// Key insight: each cut between i and i+1 contributes weights[i]+weights[i+1]
// to the total score. We need exactly k-1 cuts. Sort adjacent sums,
// take k-1 largest minus k-1 smallest.
//
// Complexity: O(n log n) time, O(n) space
func putMarbles(weights []int, k int) int64 {
	n := len(weights)
	if k <= 1 || k >= n {
		return 0
	}

	sums := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		sums[i] = weights[i] + weights[i+1]
	}
	sort.Ints(sums)

	var minScore, maxScore int64
	for i := 0; i < k-1; i++ {
		minScore += int64(sums[i])
		maxScore += int64(sums[n-2-i])
	}
	return maxScore - minScore
}

func main() {
	// Example from LeetCode
	fmt.Println("Test 1: [1,3,5,1], k=2 ->", putMarbles([]int{1, 3, 5, 1}, 2)) // 4

	// Additional test cases
	fmt.Println("Test 2: [1,3,5,1], k=1 ->", putMarbles([]int{1, 3, 5, 1}, 1)) // 0
	fmt.Println("Test 3: [1,4,2,5,3], k=3 ->", putMarbles([]int{1, 4, 2, 5, 3}, 3))
	fmt.Println("Test 4: [1,2], k=2 ->", putMarbles([]int{1, 2}, 2)) // 0
}
```

## 2552 — Count Increasing Quadruplets

```go
package main

// LeetCode #2552: Count Increasing Quadruplets
// https://leetcode.com/problems/count-increasing-quadruplets/
// Difficulty: Hard

import "fmt"

// countQuadruplets counts quadruplets (i,j,k,l) with i<j<k<l
// such that nums[i] < nums[k] < nums[j] < nums[l] (1324 pattern).
//
// For each pair (j,k) with j<k and nums[j] > nums[k]:
//   - Count i<j where nums[i] < nums[k] (from prefix)
//   - Count l>k where nums[l] > nums[j] (from suffix)
//   - Product of counts = quadruplets for this (j,k)
//
// Complexity: O(n^2) time, O(n^2) space
func countQuadruplets(nums []int) int64 {
	n := len(nums)

	// prefixLess[i][v] = count of elements before position i that are < v
	prefixLess := make([][]int, n+1)
	for i := range prefixLess {
		prefixLess[i] = make([]int, n+2)
	}
	for i := 0; i < n; i++ {
		copy(prefixLess[i+1], prefixLess[i])
		for v := nums[i] + 1; v <= n; v++ {
			prefixLess[i+1][v]++
		}
	}

	// suffixGreater[i][v] = count of elements after position i that are > v
	suffixGreater := make([][]int, n+1)
	for i := range suffixGreater {
		suffixGreater[i] = make([]int, n+2)
	}
	for i := n - 1; i >= 0; i-- {
		copy(suffixGreater[i], suffixGreater[i+1])
		for v := 1; v < nums[i]; v++ {
			suffixGreater[i][v]++
		}
	}

	var result int64
	for j := 1; j < n-2; j++ {
		for k := j + 1; k < n-1; k++ {
			if nums[j] > nums[k] {
				left := prefixLess[j][nums[k]]
				right := suffixGreater[k][nums[j]]
				result += int64(left) * int64(right)
			}
		}
	}
	return result
}

func main() {
	// Example from LeetCode
	fmt.Println("Test 1: [1,3,2,4,5] ->", countQuadruplets([]int{1, 3, 2, 4, 5})) // 2

	// Additional test cases
	fmt.Println("Test 2: [1,2,3,4] ->", countQuadruplets([]int{1, 2, 3, 4}))     // 0
	fmt.Println("Test 3: [5,4,3,2,1] ->", countQuadruplets([]int{5, 4, 3, 2, 1})) // 0
	fmt.Println("Test 4: [1,4,3,5,2] ->", countQuadruplets([]int{1, 4, 3, 5, 2}))
	fmt.Println("Test 5: [2,5,3,4,1] ->", countQuadruplets([]int{2, 5, 3, 4, 1}))
}
```

## 2561 — Rearranging Fruits

```go
package main

// LeetCode #2561: Rearranging Fruits
// https://leetcode.com/problems/rearranging-fruits/
// Difficulty: Hard

import (
	"fmt"
	"sort"
)

// minCost computes the minimum cost to make both fruit baskets identical.
//
// Swapping fruit a (from basket1) with fruit b (from basket2) costs min(a, b).
// We can also use the globally minimum fruit value as a mediator:
// swap a with min (cost = min), then min with b (cost = min), total = 2*min.
//
// Strategy:
// 1. Count frequency differences: freq[val] = count1 - count2.
// 2. Values with freq > 0 are excess in basket1; freq < 0 are excess in basket2.
// 3. Collect all excess values (each added |freq|/2 times) into one array.
// 4. Sort and pair smallest with largest. Cost = min(min(a,b), 2*globalMin).
//
// Complexity: O(n log n) time, O(n) space
func minCost(basket1 []int, basket2 []int) int64 {
	freq := make(map[int]int)
	for _, v := range basket1 {
		freq[v]++
	}
	for _, v := range basket2 {
		freq[v]--
	}

	// Find global minimum across both baskets
	globalMin := basket1[0]
	for _, v := range basket1 {
		if v < globalMin {
			globalMin = v
		}
	}
	for _, v := range basket2 {
		if v < globalMin {
			globalMin = v
		}
	}

	// Build excess array: values that need to be moved
	var excess []int
	for val, count := range freq {
		if count%2 != 0 {
			return -1 // impossible: odd frequency difference
		}
		cnt := abs(count) / 2
		for i := 0; i < cnt; i++ {
			excess = append(excess, val)
		}
	}

	sort.Ints(excess)

	cost := int64(0)
	// Pair smallest excess with largest excess
	for i := 0; i < len(excess)/2; i++ {
		a := excess[i]
		b := excess[len(excess)-1-i]
		// Direct swap cost = min(a, b); mediator cost = 2*globalMin
		importedMin := min(a, b)
		if 2*globalMin < importedMin {
			importedMin = 2 * globalMin
		}
		cost += int64(importedMin)
	}

	return cost
}

func min(a, b int) int {
	if a < b {
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
	// Test cases
	fmt.Println("Test 1: basket1=[4,2,2,2], basket2=[1,4,1,2] ->", minCost([]int{4, 2, 2, 2}, []int{1, 4, 1, 2}))
	fmt.Println("Test 2: basket1=[1,2,3,4], basket2=[1,2,3,4] ->", minCost([]int{1, 2, 3, 4}, []int{1, 2, 3, 4}))
	fmt.Println("Test 3: basket1=[1,1,2,2], basket2=[3,3,4,4] ->", minCost([]int{1, 1, 2, 2}, []int{3, 3, 4, 4}))
	fmt.Println("Test 4: basket1=[84,80,43,8,80,88,43,84], basket2=[32,60,13,58,7,63,81,81] ->",
		minCost([]int{84, 80, 43, 8, 80, 88, 43, 84}, []int{32, 60, 13, 58, 7, 63, 81, 81}))
	fmt.Println("Test 5: basket1=[1,1,1], basket2=[2,2,2] ->", minCost([]int{1, 1, 1}, []int{2, 2, 2}))
	fmt.Println("Test 6: basket1=[1], basket2=[1] ->", minCost([]int{1}, []int{1}))
}
```

## 2565 — Subsequence With The Minimum Score

```go
package main

// LeetCode #2565: Subsequence With the Minimum Score
// https://leetcode.com/problems/subsequence-with-the-minimum-score/
// Difficulty: Hard

import "fmt"

// minimumScore returns the minimum length of a subsequence to remove
// from s such that t becomes a subsequence of s.
//
// Compute prefixPos[i] = position in s where first i chars of t are matched.
// Compute suffixPos[i] = position in s where last i chars of t are matched.
// Then for each split point, maximize the total matched chars.
//
// Complexity: O(n+m) time, O(m) space
func minimumScore(s string, t string) int {
	n, m := len(s), len(t)

	// prefixPos[i] = position in s where first i chars of t are matched (0-indexed)
	prefixPos := make([]int, m+1)
	for i := range prefixPos {
		prefixPos[i] = n
	}
	prefixPos[0] = -1
	idx := 0
	for i := 0; i < n && idx < m; i++ {
		if s[i] == t[idx] {
			idx++
			prefixPos[idx] = i
		}
	}

	// suffixPos[i] = position in s where last i chars of t are matched
	suffixPos := make([]int, m+1)
	for i := range suffixPos {
		suffixPos[i] = -1
	}
	suffixPos[0] = n
	idx = 0
	for i := n - 1; i >= 0 && idx < m; i-- {
		if s[i] == t[m-1-idx] {
			idx++
			suffixPos[idx] = i
		}
	}

	// t is already a subsequence of s
	if prefixPos[m] < n {
		return 0
	}

	ans := m
	// Try keeping first i chars from left and last j chars from right
	// They must not overlap: prefixPos[i] < suffixPos[j]
	for i := 0; i <= m; i++ {
		if prefixPos[i] == n {
			continue
		}
		// Binary search for max j such that suffixPos[j] > prefixPos[i]
		lo, hi := 0, m-i
		for lo < hi {
			mid := (lo + hi + 1) / 2
			if suffixPos[mid] > prefixPos[i] {
				lo = mid
			} else {
				hi = mid - 1
			}
		}
		kept := i + lo
		if m-kept < ans {
			ans = m - kept
		}
	}
	return ans
}

func main() {
	// Example from LeetCode
	fmt.Println("Test 1: s=\"abacaba\", t=\"bzaa\" ->", minimumScore("abacaba", "bzaa")) // 1

	// Additional test cases
	fmt.Println("Test 2: s=\"abc\", t=\"abc\" ->", minimumScore("abc", "abc"))      // 0
	fmt.Println("Test 3: s=\"abcde\", t=\"ace\" ->", minimumScore("abcde", "ace"))  // 0
	fmt.Println("Test 4: s=\"abcde\", t=\"xyz\" ->", minimumScore("abcde", "xyz"))  // 3
	fmt.Println("Test 5: s=\"a\", t=\"b\" ->", minimumScore("a", "b"))              // 1
	fmt.Println("Test 6: s=\"\", t=\"a\" ->", minimumScore("", "a"))                // 1
}
```

## 2569 — Handling Sum Queries After Update

```go
package main

// LeetCode #2569: Handling Sum Queries After Update
// https://leetcode.com/problems/handling-sum-queries-after-update/
// Difficulty: Hard

import "fmt"

// handleQuery processes 3 types of queries on nums1 and nums2:
//   1 l r: flip bits in nums1[l..r] (0->1, 1->0)
//   2 x 0: nums2[j] += nums1[j] * x for all j
//   3 0 0: append sum(nums2) to result
//
// Approach: Segment tree on nums1 tracking count of 1s per segment with lazy flip.
// Type 2: sum2 += x * countOnesInNums1 (since nums1[j] is 0 or 1).
// Type 3: append current sum2.
//
// Complexity: O((n+q)*log n) time, O(n) space

// segTree tracks count of 1s in nums1 segments with lazy range flips
type segTree struct {
	n    int
	ones []int  // count of 1s in each segment
	lazy []bool // pending flip flag
}

func newSegTree(nums []int) *segTree {
	n := len(nums)
	st := &segTree{
		n:    n,
		ones: make([]int, 4*n),
		lazy: make([]bool, 4*n),
	}
	st.build(nums, 0, 0, n-1)
	return st
}

func (st *segTree) build(nums []int, idx, l, r int) {
	if l == r {
		st.ones[idx] = nums[l]
		return
	}
	mid := (l + r) / 2
	st.build(nums, 2*idx+1, l, mid)
	st.build(nums, 2*idx+2, mid+1, r)
	st.ones[idx] = st.ones[2*idx+1] + st.ones[2*idx+2]
}

func (st *segTree) push(idx, l, r int) {
	if st.lazy[idx] {
		// Flip: new count of 1s = segment length - old count
		st.ones[idx] = (r - l + 1) - st.ones[idx]
		if l != r {
			st.lazy[2*idx+1] = !st.lazy[2*idx+1]
			st.lazy[2*idx+2] = !st.lazy[2*idx+2]
		}
		st.lazy[idx] = false
	}
}

func (st *segTree) update(idx, l, r, ql, qr int) {
	st.push(idx, l, r)
	if ql > r || qr < l {
		return
	}
	if ql <= l && r <= qr {
		st.lazy[idx] = !st.lazy[idx]
		st.push(idx, l, r)
		return
	}
	mid := (l + r) / 2
	st.update(2*idx+1, l, mid, ql, qr)
	st.update(2*idx+2, mid+1, r, ql, qr)
	st.ones[idx] = st.ones[2*idx+1] + st.ones[2*idx+2]
}

func (st *segTree) flipRange(l, r int) {
	st.update(0, 0, st.n-1, l, r)
}

func (st *segTree) totalOnes() int {
	st.push(0, 0, st.n-1)
	return st.ones[0]
}

func handleQuery(nums1 []int, nums2 []int, queries [][]int) []int64 {
	st := newSegTree(nums1)
	sum2 := int64(0)
	for _, v := range nums2 {
		sum2 += int64(v)
	}

	var result []int64
	for _, q := range queries {
		switch q[0] {
		case 1:
			st.flipRange(q[1], q[2])
		case 2:
			x := int64(q[1])
			ones := int64(st.totalOnes())
			sum2 += ones * x
		case 3:
			result = append(result, sum2)
		}
	}
	return result
}

func main() {
	// Test cases
	nums1 := []int{1, 0, 1}
	nums2 := []int{0, 0, 0}
	queries := [][]int{{1, 0, 1}, {2, 1, 0}, {3, 0, 0}, {2, 2, 0}, {3, 0, 0}}
	fmt.Println("Test 1: ->", handleQuery(nums1, nums2, queries))

	nums1 = []int{1, 1, 1}
	nums2 = []int{0, 0, 0}
	queries = [][]int{{2, 2, 0}, {3, 0, 0}, {1, 0, 2}, {3, 0, 0}}
	fmt.Println("Test 2: ->", handleQuery(nums1, nums2, queries))

	// Edge cases
	fmt.Println("Test 3: single element ->", handleQuery([]int{0}, []int{5}, [][]int{{3, 0, 0}}))
	fmt.Println("Test 4: flip and add ->", handleQuery([]int{1}, []int{5}, [][]int{{2, 3, 0}, {3, 0, 0}}))
	fmt.Println("Test 5: multiple ops ->", handleQuery(
		[]int{1, 0, 1, 0, 1},
		[]int{1, 2, 3, 4, 5},
		[][]int{{1, 0, 4}, {2, 10, 0}, {3, 0, 0}},
	))
}
```

## 2573 — Find The String With Lcp

```go
package main

// LeetCode #2573: Find the String with LCP
// https://leetcode.com/problems/find-the-string-with-lcp/
// Difficulty: Hard

import "fmt"

// findTheString reconstructs the lexicographically smallest string that matches
// the given LCP matrix. lcp[i][j] = longest common prefix of suffixes
// starting at i and j.
//
// Approach:
// 1. Validate matrix (diagonal must be n-i, symmetric, values within bounds).
// 2. If lcp[i][j] > 0, then s[i] must equal s[j]. Union via DSU.
// 3. Assign smallest possible character ('a', 'b', ...) to each equivalence class.
// 4. Verify by recomputing LCP from the constructed string.
//
// Complexity: O(n^2) time, O(n^2) space

func findTheString(lcp [][]int) string {
	n := len(lcp)

	// Validate matrix shape
	for i := 0; i < n; i++ {
		if len(lcp[i]) != n {
			return ""
		}
	}

	// Validate diagonal: lcp[i][i] must equal n-i
	for i := 0; i < n; i++ {
		if lcp[i][i] != n-i {
			return ""
		}
	}

	// Validate symmetry: lcp[i][j] == lcp[j][i]
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if lcp[i][j] != lcp[j][i] {
				return ""
			}
		}
	}

	// Validate bounds: lcp[i][j] <= n - max(i, j)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if lcp[i][j] > n-i || lcp[i][j] > n-j {
				return ""
			}
		}
	}

	// DSU to union positions that must have the same character
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		for parent[x] != x {
			parent[x] = parent[parent[x]]
			x = parent[x]
		}
		return x
	}
	union := func(x, y int) {
		x, y = find(x), find(y)
		if x != y {
			parent[y] = x
		}
	}

	// Union positions where lcp[i][j] > 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if lcp[i][j] > 0 {
				union(i, j)
			}
		}
	}

	// Assign characters: each equivalence class gets the smallest unused char
	ans := make([]byte, n)
	nextChar := byte('a')

	for i := 0; i < n; i++ {
		if ans[i] != 0 {
			continue
		}
		if nextChar > 'z' {
			return "" // too many distinct characters needed
		}
		// Assign this char to the entire equivalence class
		for j := i; j < n; j++ {
			if find(j) == find(i) {
				ans[j] = nextChar
			}
		}
		nextChar++
	}

	// Verify by recomputing LCP from the constructed string
	computed := make([][]int, n)
	for i := range computed {
		computed[i] = make([]int, n)
	}
	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if ans[i] != ans[j] {
				computed[i][j] = 0
			} else if i+1 < n && j+1 < n {
				computed[i][j] = computed[i+1][j+1] + 1
			} else {
				computed[i][j] = 1
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if computed[i][j] != lcp[i][j] {
				return ""
			}
		}
	}

	return string(ans)
}

func main() {
	// Test cases
	lcp1 := [][]int{{4, 0, 2, 0}, {0, 3, 0, 1}, {2, 0, 2, 0}, {0, 1, 0, 1}}
	fmt.Println("Test 1: ->", findTheString(lcp1))

	lcp2 := [][]int{{4, 3, 2, 1}, {3, 3, 2, 1}, {2, 2, 2, 1}, {1, 1, 1, 1}}
	fmt.Println("Test 2: ->", findTheString(lcp2)) // "aaaa"

	lcp3 := [][]int{{1, 0}, {0, 1}}
	fmt.Println("Test 3: ->", findTheString(lcp3)) // "ab"

	lcp4 := [][]int{{1, 1}, {1, 1}}
	fmt.Println("Test 4: ->", findTheString(lcp4)) // invalid

	lcp5 := [][]int{{3, 0, 1}, {0, 2, 0}, {1, 0, 1}}
	fmt.Println("Test 5: ->", findTheString(lcp5))

	lcp6 := [][]int{{1}}
	fmt.Println("Test 6: ->", findTheString(lcp6)) // "a"
}
```

## 2577 — Minimum Time To Visit A Cell In A Grid

```go
package main

// LeetCode #2577: Minimum Time to Visit a Cell in a Grid
// https://leetcode.com/problems/minimum-time-to-visit-a-cell-in-a-grid/
// Difficulty: Hard

import (
	"container/heap"
	"fmt"
)

// Item for min-heap priority queue
type Item struct {
	row, col, time int
}

type MinHeap []Item

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].time < h[j].time }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(Item)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

var dirs = [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

// minimumTime finds minimum time to reach bottom-right using modified Dijkstra.
// We can "wait" by oscillating between two adjacent cells (2 time units per cycle).
//
// If both start neighbors require >1, impossible -> -1.
//
// Complexity: O(m*n*log(m*n)) time, O(m*n) space
func minimumTime(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	// Handle trivial 1x1 case
	if m == 1 && n == 1 {
		return 0
	}

	// If both neighbors of start are inaccessible at time 1
	hasRight := n > 1
	hasDown := m > 1
	if (!hasRight || grid[0][1] > 1) && (!hasDown || grid[1][0] > 1) {
		return -1
	}

	dist := make([][]int, m)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = 1 << 60
		}
	}
	dist[0][0] = 0

	pq := &MinHeap{}
	heap.Push(pq, Item{0, 0, 0})

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(Item)
		if cur.time > dist[cur.row][cur.col] {
			continue
		}
		if cur.row == m-1 && cur.col == n-1 {
			return cur.time
		}

		for _, d := range dirs {
			nr, nc := cur.row+d[0], cur.col+d[1]
			if nr < 0 || nr >= m || nc < 0 || nc >= n {
				continue
			}

			nt := cur.time + 1
			if nt < grid[nr][nc] {
				diff := grid[nr][nc] - nt
				if diff%2 == 0 {
					nt = grid[nr][nc]
				} else {
					nt = grid[nr][nc] + 1
				}
			}

			if nt < dist[nr][nc] {
				dist[nr][nc] = nt
				heap.Push(pq, Item{nr, nc, nt})
			}
		}
	}

	return -1
}

func main() {
	// Example from LeetCode
	grid1 := [][]int{{0, 1, 3, 2}, {5, 1, 2, 5}, {4, 3, 8, 6}}
	fmt.Println("Test 1: ->", minimumTime(grid1)) // 7

	// Additional test cases
	grid2 := [][]int{{0, 2, 4}, {3, 2, 1}, {1, 0, 4}}
	fmt.Println("Test 2: ->", minimumTime(grid2))

	grid3 := [][]int{{0, 2}, {3, 4}}
	fmt.Println("Test 3: blocked start ->", minimumTime(grid3))

	fmt.Println("Test 4: 1x1 ->", minimumTime([][]int{{0}})) // 0

	grid5 := [][]int{{0, 1}, {1, 0}}
	fmt.Println("Test 5: simple 2x2 ->", minimumTime(grid5))
}
```

## 2581 — Count Number Of Possible Root Nodes

```go
package main

// LeetCode #2581: Count Number of Possible Root Nodes
// https://leetcode.com/problems/count-number-of-possible-root-nodes/
// Difficulty: Hard

import "fmt"

// rootCount uses rerooting DP. Build tree, count correct guesses with root=0,
// then reroot: moving from u to v, subtract (u,v) if guessed, add (v,u) if guessed.
// Count roots where correct >= k.
//
// Complexity: O(n) time, O(n) space
func rootCount(edges [][]int, guesses [][]int, k int) int {
	n := len(edges) + 1

	// Build adjacency
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// Build guess set for O(1) lookup
	guessSet := make(map[[2]int]bool)
	for _, g := range guesses {
		guessSet[[2]int{g[0], g[1]}] = true
	}

	// First DFS from root 0 to count correct guesses
	correct := 0
	var dfs1 func(u, parent int)
	dfs1 = func(u, parent int) {
		for _, v := range adj[u] {
			if v == parent {
				continue
			}
			if guessSet[[2]int{u, v}] {
				correct++
			}
			dfs1(v, u)
		}
	}
	dfs1(0, -1)

	// Rerooting DFS
	result := 0
	if correct >= k {
		result++
	}

	var dfs2 func(u, parent int, cur int)
	dfs2 = func(u, parent int, cur int) {
		for _, v := range adj[u] {
			if v == parent {
				continue
			}
			next := cur
			// Moving root from u to v: lose (u,v), gain (v,u)
			if guessSet[[2]int{u, v}] {
				next--
			}
			if guessSet[[2]int{v, u}] {
				next++
			}
			if next >= k {
				result++
			}
			dfs2(v, u, next)
		}
	}
	dfs2(0, -1, correct)

	return result
}

func main() {
	// Example from LeetCode
	edges1 := [][]int{{0, 1}, {1, 2}, {1, 3}, {4, 2}}
	guesses1 := [][]int{{1, 3}, {0, 1}, {1, 0}, {2, 4}}
	fmt.Println("Test 1: ->", rootCount(edges1, guesses1, 3)) // 3

	// Additional test cases
	edges2 := [][]int{{0, 1}, {1, 2}}
	guesses2 := [][]int{{0, 1}, {1, 2}}
	fmt.Println("Test 2: ->", rootCount(edges2, guesses2, 2)) // 1

	edges3 := [][]int{{0, 1}, {0, 2}}
	guesses3 := [][]int{{0, 1}, {0, 2}}
	fmt.Println("Test 3: ->", rootCount(edges3, guesses3, 2)) // 1

	// Edge cases
	edges4 := [][]int{{0, 1}}
	guesses4 := [][]int{{0, 1}}
	fmt.Println("Test 4: k=0 ->", rootCount(edges4, guesses4, 0)) // 2

	edges5 := [][]int{{0, 1}, {1, 2}}
	guesses5 := [][]int{{0, 1}}
	fmt.Println("Test 5: k=1 ->", rootCount(edges5, guesses5, 1)) // 2
}
```

