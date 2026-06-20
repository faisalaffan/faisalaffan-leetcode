# Easy (Mudah) — Problem ��3688

## 3174 — Clear Digits

```go
package main

// LeetCode #3174: Clear Digits
// https://leetcode.com/problems/clear-digits/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ClearDigits("abc"))
	fmt.Println(ClearDigits("cb34"))
	fmt.Println(ClearDigits("a1b2c3"))
}

// ClearDigits removes all digits and their nearest non-digit character to the left.
// Time: O(n). Space: O(n).
func ClearDigits(s string) string {
	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}
```

## 3178 — Find The Child Who Has The Ball After K Seconds

```go
package main

// LeetCode #3178: Find the Child Who Has the Ball After K Seconds
// https://leetcode.com/problems/find-the-child-who-has-the-ball-after-k-seconds/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindTheChildWhoHasTheBallAfterKSeconds(3, 5))
	fmt.Println(FindTheChildWhoHasTheBallAfterKSeconds(5, 6))
	fmt.Println(FindTheChildWhoHasTheBallAfterKSeconds(4, 2))
}

// FindTheChildWhoHasTheBallAfterKSeconds returns the child who has the ball after k seconds.
// Children pass the ball left-to-right, then right-to-left, repeatedly.
// Time: O(1). Space: O(1).
func FindTheChildWhoHasTheBallAfterKSeconds(n int, k int) int {
	cycleLen := 2 * (n - 1)
	k %= cycleLen
	if k < n {
		return k
	}
	return cycleLen - k
}
```

## 3184 — Count Pairs That Form A Complete Day I

```go
package main

// LeetCode #3184: Count Pairs That Form a Complete Day I
// https://leetcode.com/problems/count-pairs-that-form-a-complete-day-i/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountPairsThatFormACompleteDayI([]int{12, 12, 30, 24, 24}))
	fmt.Println(CountPairsThatFormACompleteDayI([]int{72, 48, 24, 3}))
}

// CountPairsThatFormACompleteDayI counts pairs (i, j) where i < j and hours[i] + hours[j] is divisible by 24.
// Time: O(n). Space: O(24).
func CountPairsThatFormACompleteDayI(hours []int) int {
	count := 0
	rem := make([]int, 24)
	for _, h := range hours {
		r := h % 24
		need := (24 - r) % 24
		count += rem[need]
		rem[r]++
	}
	return count
}
```

## 3190 — Find Minimum Operations To Make All Elements Divisible By Three

```go
package main

// LeetCode #3190: Find Minimum Operations to Make All Elements Divisible by Three
// https://leetcode.com/problems/find-minimum-operations-to-make-all-elements-divisible-by-three/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindMinimumOperationsToMakeAllElementsDivisibleByThree([]int{1, 2, 3, 4}))
	fmt.Println(FindMinimumOperationsToMakeAllElementsDivisibleByThree([]int{3, 6, 9}))
}

// FindMinimumOperationsToMakeAllElementsDivisibleByThree returns the minimum operations to make all elements divisible by 3.
// Each operation adds or subtracts 1 from an element.
// Time: O(n). Space: O(1).
func FindMinimumOperationsToMakeAllElementsDivisibleByThree(nums []int) int {
	ops := 0
	for _, num := range nums {
		r := num % 3
		if r == 1 || r == 2 {
			ops++
		}
	}
	return ops
}
```

## 3194 — Minimum Average Of Smallest And Largest Elements

```go
package main

// LeetCode #3194: Minimum Average of Smallest and Largest Elements
// https://leetcode.com/problems/minimum-average-of-smallest-and-largest-elements/
// Difficulty: Easy

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(MinimumAverageOfSmallestAndLargestElements([]int{1, 9, 8, 3, 10, 5}))
	fmt.Println(MinimumAverageOfSmallestAndLargestElements([]int{1, 2, 3, 7, 8, 9}))
}

// MinimumAverageOfSmallestAndLargestElements returns the minimum average of the smallest and largest elements.
// Time: O(n log n). Space: O(1) (or O(n) due to sorting).
func MinimumAverageOfSmallestAndLargestElements(nums []int) float64 {
	sort.Ints(nums)
	n := len(nums)
	minAvg := math.MaxFloat64
	for i := 0; i < n/2; i++ {
		avg := float64(nums[i]+nums[n-1-i]) / 2.0
		if avg < minAvg {
			minAvg = avg
		}
	}
	return minAvg
}
```

## 3198 — Find Cities In Each State

```go
package main

// LeetCode #3198: Find Cities in Each State
// https://leetcode.com/problems/find-cities-in-each-state/
// Difficulty: Easy [Paid]

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	cities := [][]string{
		{"New York", "Albany"},
		{"California", "Los Angeles"},
		{"New York", "Buffalo"},
		{"California", "San Francisco"},
		{"Texas", "Houston"},
	}
	result := FindCitiesInEachState(cities)
	for _, r := range result {
		fmt.Println(r)
	}
}

// FindCitiesInEachState groups cities by state and returns them as comma-separated strings.
// Time: O(n log n). Space: O(n).
func FindCitiesInEachState(data [][]string) [][]string {
	stateCities := make(map[string][]string)
	for _, row := range data {
		state, city := row[0], row[1]
		stateCities[state] = append(stateCities[state], city)
	}

	states := make([]string, 0, len(stateCities))
	for s := range stateCities {
		states = append(states, s)
	}
	sort.Strings(states)

	result := make([][]string, len(states))
	for i, s := range states {
		cities := stateCities[s]
		sort.Strings(cities)
		result[i] = []string{s, strings.Join(cities, ", ")}
	}
	return result
}
```

## 3199 — Count Triplets With Even Xor Set Bits I

```go
package main

// LeetCode #3199: Count Triplets with Even XOR Set Bits I
// https://leetcode.com/problems/count-triplets-with-even-xor-set-bits-i/
// Difficulty: Easy [Paid]

import "fmt"

func main() {
	fmt.Println(CountTripletsWithEvenXorSetBitsI([]int{1}, []int{2}, []int{3}))
	fmt.Println(CountTripletsWithEvenXorSetBitsI([]int{1, 2}, []int{3, 4}, []int{5, 6}))
}

// popcount returns the number of set bits in x.
func popcount(x int) int {
	count := 0
	for x > 0 {
		count += x & 1
		x >>= 1
	}
	return count
}

// CountTripletsWithEvenXorSetBitsI counts triplets (i, j, k) where a[i] XOR b[j] XOR c[k] has even number of set bits.
// Time: O(|a| * |b| * |c|). Space: O(1).
func CountTripletsWithEvenXorSetBitsI(a []int, b []int, c []int) int {
	count := 0
	for _, va := range a {
		for _, vb := range b {
			for _, vc := range c {
				if popcount(va^vb^vc)%2 == 0 {
					count++
				}
			}
		}
	}
	return count
}
```

## 3200 — Maximum Height Of A Triangle

```go
package main

// LeetCode #3200: Maximum Height of a Triangle
// https://leetcode.com/problems/maximum-height-of-a-triangle/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MaximumHeightOfATriangle(2, 4))
	fmt.Println(MaximumHeightOfATriangle(2, 1))
	fmt.Println(MaximumHeightOfATriangle(10, 10))
}

// maxHeight tries building a triangle starting with the given first color.
func maxHeight(red, blue int, firstRed bool) int {
	h := 0
	need := 1
	for {
		if firstRed {
			if red < need {
				break
			}
			red -= need
		} else {
			if blue < need {
				break
			}
			blue -= need
		}
		h++
		need++
		firstRed = !firstRed
	}
	return h
}

// MaximumHeightOfATriangle returns the maximum height of a triangle using red and blue balls.
// Time: O(sqrt(n)). Space: O(1).
func MaximumHeightOfATriangle(red int, blue int) int {
	h1 := maxHeight(red, blue, true)
	h2 := maxHeight(red, blue, false)
	if h1 > h2 {
		return h1
	}
	return h2
}
```

## 3206 — Alternating Groups I

```go
package main

// LeetCode #3206: Alternating Groups I
// https://leetcode.com/problems/alternating-groups-i/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(AlternatingGroupsI([]int{1, 1, 1}))
	fmt.Println(AlternatingGroupsI([]int{0, 1, 0, 0, 1}))
}

// AlternatingGroupsI counts the number of groups of 3 adjacent elements where all three are alternating.
// Time: O(n). Space: O(1).
func AlternatingGroupsI(colors []int) int {
	n := len(colors)
	count := 0
	for i := 0; i < n; i++ {
		if colors[i] == colors[(i+2)%n] && colors[i] != colors[(i+1)%n] {
			count++
		}
	}
	return count
}
```

## 3210 — Find The Encrypted String

```go
package main

// LeetCode #3210: Find the Encrypted String
// https://leetcode.com/problems/find-the-encrypted-string/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindTheEncryptedString("dart", 3))
	fmt.Println(FindTheEncryptedString("aaa", 1))
	fmt.Println(FindTheEncryptedString("abcd", 5))
}

// FindTheEncryptedString returns the encrypted string by rotating each character by k positions forward.
// Time: O(n). Space: O(n).
func FindTheEncryptedString(s string, k int) string {
	n := len(s)
	if n == 0 {
		return s
	}
	k %= n
	return s[k:] + s[:k]
}
```

## 3216 — Lexicographically Smallest String After A Swap

```go
package main

// LeetCode #3216: Lexicographically Smallest String After a Swap
// https://leetcode.com/problems/lexicographically-smallest-string-after-a-swap/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(LexicographicallySmallestStringAfterASwap("45320"))
	fmt.Println(LexicographicallySmallestStringAfterASwap("001"))
}

// LexicographicallySmallestStringAfterASwap makes the smallest string by swapping one pair of adjacent same-parity digits where left > right.
// Time: O(n). Space: O(n).
func LexicographicallySmallestStringAfterASwap(s string) string {
	b := []byte(s)
	for i := 0; i < len(b)-1; i++ {
		if b[i] > b[i+1] && (b[i]%2 == b[i+1]%2) {
			b[i], b[i+1] = b[i+1], b[i]
			break
		}
	}
	return string(b)
}
```

## 3222 — Find The Winning Player In Coin Game

```go
package main

// LeetCode #3222: Find the Winning Player in Coin Game
// https://leetcode.com/problems/find-the-winning-player-in-coin-game/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindTheWinningPlayerInCoinGame(2, 7))
	fmt.Println(FindTheWinningPlayerInCoinGame(4, 11))
}

// FindTheWinningPlayerInCoinGame returns the player who wins the coin game.
// Players alternately take 1 "75" coin and 4 "10" coins (value 115).
// Alice goes first. If a player cannot take exactly 115, they lose.
// Time: O(min(x, y/4)). Space: O(1).
func FindTheWinningPlayerInCoinGame(x int, y int) string {
	turns := 0
	for x >= 1 && y >= 4 {
		x -= 1
		y -= 4
		turns++
	}
	if turns%2 == 0 {
		return "Bob"
	}
	return "Alice"
}
```

## 3226 — Number Of Bit Changes To Make Two Integers Equal

```go
package main

// LeetCode #3226: Number of Bit Changes to Make Two Integers Equal
// https://leetcode.com/problems/number-of-bit-changes-to-make-two-integers-equal/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(NumberOfBitChangesToMakeTwoIntegersEqual(13, 4))
	fmt.Println(NumberOfBitChangesToMakeTwoIntegersEqual(21, 21))
	fmt.Println(NumberOfBitChangesToMakeTwoIntegersEqual(14, 13))
}

// NumberOfBitChangesToMakeTwoIntegersEqual returns the number of bit changes needed to make n equal to k, or -1 if impossible.
// Time: O(log n). Space: O(1).
func NumberOfBitChangesToMakeTwoIntegersEqual(n int, k int) int {
	// n must have all set bits that k has, since we can only change 1->0, not 0->1
	if n&k != k {
		return -1
	}
	// Count bits where n has 1 and k has 0
	diff := n ^ k
	count := 0
	for diff > 0 {
		count += diff & 1
		diff >>= 1
	}
	return count
}
```

## 3232 — Find If Digit Game Can Be Won

```go
package main

// LeetCode #3232: Find if Digit Game Can Be Won
// https://leetcode.com/problems/find-if-digit-game-can-be-won/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindIfDigitGameCanBeWon([]int{1, 2, 3, 4, 10}))
	fmt.Println(FindIfDigitGameCanBeWon([]int{1, 2, 3, 4, 5, 14}))
}

// FindIfDigitGameCanBeWon returns true if Alice can win the digit game.
// Alice takes all single-digit numbers, Bob takes all two-digit+ numbers.
// Alice wins if the sums are not equal.
// Time: O(n). Space: O(1).
func FindIfDigitGameCanBeWon(nums []int) bool {
	sumSingle := 0
	sumDouble := 0
	for _, num := range nums {
		if num < 10 {
			sumSingle += num
		} else {
			sumDouble += num
		}
	}
	return sumSingle != sumDouble
}
```

## 3238 — Find The Number Of Winning Players

```go
package main

// LeetCode #3238: Find the Number of Winning Players
// https://leetcode.com/problems/find-the-number-of-winning-players/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindTheNumberOfWinningPlayers(4, [][]int{{0, 0}, {1, 0}, {1, 0}, {2, 1}, {2, 1}, {2, 0}}))
	fmt.Println(FindTheNumberOfWinningPlayers(5, [][]int{{1, 1}, {1, 2}, {1, 3}, {1, 4}}))
}

// FindTheNumberOfWinningPlayers counts players who have picked at least i+1 balls of the same color (where i is player index).
// Time: O(n). Space: O(n).
func FindTheNumberOfWinningPlayers(n int, pick [][]int) int {
	// Count colors per player
	playerColors := make([]map[int]int, n)
	for i := range playerColors {
		playerColors[i] = make(map[int]int)
	}
	for _, p := range pick {
		player, color := p[0], p[1]
		playerColors[player][color]++
	}

	winners := 0
	for i := 0; i < n; i++ {
		for _, count := range playerColors[i] {
			if count > i {
				winners++
				break
			}
		}
	}
	return winners
}
```

## 3242 — Design Neighbor Sum Service

```go
package main

// LeetCode #3242: Design Neighbor Sum Service
// https://leetcode.com/problems/design-neighbor-sum-service/
// Difficulty: Easy

import "fmt"

func main() {
	grid := [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
	}
	ns := NewNeighborSum(grid)
	fmt.Println(ns.AdjacentSum(1))
	fmt.Println(ns.AdjacentSum(4))
	fmt.Println(ns.DiagonalSum(4))
	fmt.Println(ns.DiagonalSum(8))
}

// NeighborSum provides sums of adjacent/diagonal elements for a given value in a grid.
type NeighborSum struct {
	grid [][]int
	pos  map[int][2]int // value -> [row, col]
}

// NewNeighborSum initializes a NeighborSum with the given grid.
// Time: O(n^2). Space: O(n^2).
func NewNeighborSum(grid [][]int) NeighborSum {
	pos := make(map[int][2]int)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			pos[grid[i][j]] = [2]int{i, j}
		}
	}
	return NeighborSum{grid: grid, pos: pos}
}

// AdjacentSum returns the sum of orthogonal neighbors of the cell containing value.
// Time: O(1). Space: O(1).
func (ns *NeighborSum) AdjacentSum(value int) int {
	p, ok := ns.pos[value]
	if !ok {
		return 0
	}
	r, c := p[0], p[1]
	sum := 0
	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, d := range dirs {
		nr, nc := r+d[0], c+d[1]
		if nr >= 0 && nr < len(ns.grid) && nc >= 0 && nc < len(ns.grid[0]) {
			sum += ns.grid[nr][nc]
		}
	}
	return sum
}

// DiagonalSum returns the sum of diagonal neighbors of the cell containing value.
// Time: O(1). Space: O(1).
func (ns *NeighborSum) DiagonalSum(value int) int {
	p, ok := ns.pos[value]
	if !ok {
		return 0
	}
	r, c := p[0], p[1]
	sum := 0
	dirs := [][2]int{{-1, -1}, {-1, 1}, {1, -1}, {1, 1}}
	for _, d := range dirs {
		nr, nc := r+d[0], c+d[1]
		if nr >= 0 && nr < len(ns.grid) && nc >= 0 && nc < len(ns.grid[0]) {
			sum += ns.grid[nr][nc]
		}
	}
	return sum
}
```

## 3246 — Premier League Table Ranking

```go
package main

// LeetCode #3246: Premier League Table Ranking
// https://leetcode.com/problems/premier-league-table-ranking/
// Difficulty: Easy [Paid]

import (
	"fmt"
	"sort"
)

func main() {
	stats := []TeamStat{
		{TeamID: 1, TeamName: "City", Wins: 10, Draws: 3, Losses: 2},
		{TeamID: 2, TeamName: "United", Wins: 8, Draws: 5, Losses: 2},
		{TeamID: 3, TeamName: "Liverpool", Wins: 10, Draws: 3, Losses: 2},
	}
	result := PremierLeagueTableRanking(stats)
	for _, r := range result {
		fmt.Println(r)
	}
}

// TeamStat represents a team's statistics.
type TeamStat struct {
	TeamID   int
	TeamName string
	Wins     int
	Draws    int
	Losses   int
}

// TeamRank represents a team's final ranking.
type TeamRank struct {
	TeamID   int
	TeamName string
	Points   int
	Position int
}

// PremierLeagueTableRanking ranks teams by points (3 per win, 1 per draw), with ties sharing the same rank.
// Time: O(n log n). Space: O(n).
func PremierLeagueTableRanking(stats []TeamStat) []TeamRank {
	type team struct {
		id     int
		name   string
		points int
	}
	teams := make([]team, len(stats))
	for i, s := range stats {
		teams[i] = team{s.TeamID, s.TeamName, s.Wins*3 + s.Draws}
	}
	sort.Slice(teams, func(i, j int) bool {
		if teams[i].points != teams[j].points {
			return teams[i].points > teams[j].points
		}
		return teams[i].name < teams[j].name
	})

	result := make([]TeamRank, len(teams))
	for i, t := range teams {
		rank := i + 1
		if i > 0 && t.points == teams[i-1].points {
			rank = result[i-1].Position
		}
		result[i] = TeamRank{t.id, t.name, t.points, rank}
	}
	return result
}
```

## 3248 — Snake In Matrix

```go
package main

// LeetCode #3248: Snake in Matrix
// https://leetcode.com/problems/snake-in-matrix/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(SnakeInMatrix(3, []string{"RIGHT", "DOWN"}))
	fmt.Println(SnakeInMatrix(2, []string{"DOWN", "RIGHT", "UP"}))
}

// SnakeInMatrix returns the final position of the snake in an n x n matrix after following commands.
// Time: O(m). Space: O(1).
func SnakeInMatrix(n int, commands []string) int {
	r, c := 0, 0
	for _, cmd := range commands {
		switch cmd {
		case "UP":
			r--
		case "DOWN":
			r++
		case "LEFT":
			c--
		case "RIGHT":
			c++
		}
	}
	return r*n + c
}
```

## 3258 — Count Substrings That Satisfy K Constraint I

```go
package main

// LeetCode #3258: Count Substrings That Satisfy K-Constraint I
// https://leetcode.com/problems/count-substrings-that-satisfy-k-constraint-i/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountSubstringsThatSatisfyKConstraintI("10101", 1))
	fmt.Println(CountSubstringsThatSatisfyKConstraintI("1010101", 2))
	fmt.Println(CountSubstringsThatSatisfyKConstraintI("11111", 1))
}

// CountSubstringsThatSatisfyKConstraintI counts substrings where both number of 0s and 1s <= k.
// Time: O(n^2). Space: O(1).
func CountSubstringsThatSatisfyKConstraintI(s string, k int) int {
	n := len(s)
	count := 0
	for i := 0; i < n; i++ {
		zeros, ones := 0, 0
		for j := i; j < n; j++ {
			if s[j] == '0' {
				zeros++
			} else {
				ones++
			}
			if zeros <= k || ones <= k {
				count++
			} else {
				break
			}
		}
	}
	return count
}
```

## 3263 — Convert Doubly Linked List To Array I

```go
package main

// LeetCode #3263: Convert Doubly Linked List to Array I
// https://leetcode.com/problems/convert-doubly-linked-list-to-array-i/
// Difficulty: Easy [Paid]

import "fmt"

func main() {
	// 1 <-> 2 <-> 3
	head := &Node{Val: 1}
	head.Next = &Node{Val: 2, Prev: head}
	head.Next.Next = &Node{Val: 3, Prev: head.Next}
	fmt.Println(ConvertDoublyLinkedListToArrayI(head))
}

// Node represents a doubly-linked list node.
type Node struct {
	Val  int
	Next *Node
	Prev *Node
}

// ConvertDoublyLinkedListToArrayI converts a doubly linked list to an integer array.
// Time: O(n). Space: O(n).
func ConvertDoublyLinkedListToArrayI(head *Node) []int {
	result := []int{}
	for cur := head; cur != nil; cur = cur.Next {
		result = append(result, cur.Val)
	}
	return result
}
```

## 3264 — Final Array State After K Multiplication Operations I

```go
package main

// LeetCode #3264: Final Array State After K Multiplication Operations I
// https://leetcode.com/problems/final-array-state-after-k-multiplication-operations-i/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FinalArrayStateAfterKMultiplicationOperationsI([]int{2, 1, 3, 5, 6}, 5, 2))
	fmt.Println(FinalArrayStateAfterKMultiplicationOperationsI([]int{1, 2}, 3, 4))
}

// FinalArrayStateAfterKMultiplicationOperationsI finds the minimum element each time, multiplies it by multiplier, and repeats k times.
// Time: O(k * n). Space: O(1).
func FinalArrayStateAfterKMultiplicationOperationsI(nums []int, k int, multiplier int) []int {
	for t := 0; t < k; t++ {
		// Find index of minimum element
		minIdx := 0
		for i := 1; i < len(nums); i++ {
			if nums[i] < nums[minIdx] {
				minIdx = i
			}
		}
		nums[minIdx] *= multiplier
	}
	return nums
}
```

## 3270 — Find The Key Of The Numbers

```go
package main

// LeetCode #3270: Find the Key of the Numbers
// https://leetcode.com/problems/find-the-key-of-the-numbers/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindTheKeyOfTheNumbers(1, 10, 1000))
	fmt.Println(FindTheKeyOfTheNumbers(987, 879, 798))
	fmt.Println(FindTheKeyOfTheNumbers(123, 456, 789))
}

// minDigit returns the minimum digit in a number at a given decimal place (1, 10, 100, 1000).
func minDigit(num1, num2, num3, place int) int {
	d1 := (num1 / place) % 10
	d2 := (num2 / place) % 10
	d3 := (num3 / place) % 10
	minD := d1
	if d2 < minD {
		minD = d2
	}
	if d3 < minD {
		minD = d3
	}
	return minD
}

// FindTheKeyOfTheNumbers returns a 4-digit key by taking the minimum digit at each position across three numbers.
// Time: O(1). Space: O(1).
func FindTheKeyOfTheNumbers(num1 int, num2 int, num3 int) int {
	key := 0
	places := []int{1000, 100, 10, 1}
	for _, p := range places {
		key = key*10 + minDigit(num1, num2, num3, p)
	}
	return key
}
```

## 3274 — Check If Two Chessboard Squares Have The Same Color

```go
package main

// LeetCode #3274: Check if Two Chessboard Squares Have the Same Color
// https://leetcode.com/problems/check-if-two-chessboard-squares-have-the-same-color/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CheckIfTwoChessboardSquaresHaveTheSameColor("a1", "c3"))
	fmt.Println(CheckIfTwoChessboardSquaresHaveTheSameColor("a1", "h3"))
}

// CheckIfTwoChessboardSquaresHaveTheSameColor returns true if both squares are the same color on a chessboard.
// Time: O(1). Space: O(1).
func CheckIfTwoChessboardSquaresHaveTheSameColor(coordinate1 string, coordinate2 string) bool {
	c1 := (int(coordinate1[0]-'a') + int(coordinate1[1]-'1')) % 2
	c2 := (int(coordinate2[0]-'a') + int(coordinate2[1]-'1')) % 2
	return c1 == c2
}
```

## 3280 — Convert Date To Binary

```go
package main

// LeetCode #3280: Convert Date to Binary
// https://leetcode.com/problems/convert-date-to-binary/
// Difficulty: Easy

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(ConvertDateToBinary("2080-02-29"))
	fmt.Println(ConvertDateToBinary("1900-01-01"))
}

// toBinary converts an integer to its binary string representation without leading zeros.
func toBinary(n int) string {
	if n == 0 {
		return "0"
	}
	s := ""
	for n > 0 {
		s = string('0'+byte(n%2)) + s
		n /= 2
	}
	return s
}

// ConvertDateToBinary converts a date string to binary format.
// Time: O(1). Space: O(1).
func ConvertDateToBinary(date string) string {
	parts := strings.Split(date, "-")
	year, _ := strconv.Atoi(parts[0])
	month, _ := strconv.Atoi(parts[1])
	day, _ := strconv.Atoi(parts[2])
	return toBinary(year) + "-" + toBinary(month) + "-" + toBinary(day)
}
```

## 3285 — Find Indices Of Stable Mountains

```go
package main

// LeetCode #3285: Find Indices of Stable Mountains
// https://leetcode.com/problems/find-indices-of-stable-mountains/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindIndicesOfStableMountains([]int{1, 2, 3, 4, 5}, 2))
	fmt.Println(FindIndicesOfStableMountains([]int{10, 1, 10, 1, 10}, 3))
}

// FindIndicesOfStableMountains returns indices of stable mountains (where the previous mountain's height > threshold).
// Time: O(n). Space: O(n).
func FindIndicesOfStableMountains(height []int, threshold int) []int {
	result := []int{}
	for i := 1; i < len(height); i++ {
		if height[i-1] > threshold {
			result = append(result, i)
		}
	}
	return result
}
```

## 3289 — The Two Sneaky Numbers Of Digitville

```go
package main

// LeetCode #3289: The Two Sneaky Numbers of Digitville
// https://leetcode.com/problems/the-two-sneaky-numbers-of-digitville/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(TheTwoSneakyNumbersOfDigitville([]int{0, 1, 1, 0}))
	fmt.Println(TheTwoSneakyNumbersOfDigitville([]int{0, 3, 2, 1, 3, 2}))
	fmt.Println(TheTwoSneakyNumbersOfDigitville([]int{7, 1, 5, 4, 3, 4, 6, 0, 9, 5, 8, 2}))
}

// TheTwoSneakyNumbersOfDigitville returns the two numbers that appear twice in the array.
// Time: O(n). Space: O(n).
func TheTwoSneakyNumbersOfDigitville(nums []int) []int {
	seen := make(map[int]int)
	result := []int{}
	for _, num := range nums {
		seen[num]++
		if seen[num] == 2 {
			result = append(result, num)
		}
	}
	return result
}
```

## 3300 — Minimum Element After Replacement With Digit Sum

```go
package main

// LeetCode #3300: Minimum Element After Replacement With Digit Sum
// https://leetcode.com/problems/minimum-element-after-replacement-with-digit-sum/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MinimumElementAfterReplacementWithDigitSum([]int{10, 12, 13, 14}))
	fmt.Println(MinimumElementAfterReplacementWithDigitSum([]int{1, 2, 3, 4}))
	fmt.Println(MinimumElementAfterReplacementWithDigitSum([]int{999, 19, 199}))
}

// digitSum returns the sum of digits of n.
func digitSum(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

// MinimumElementAfterReplacementWithDigitSum returns the minimum element after replacing each element with its digit sum.
// Time: O(n * log n). Space: O(1).
func MinimumElementAfterReplacementWithDigitSum(nums []int) int {
	minVal := int(^uint(0) >> 1) // MaxInt
	for _, num := range nums {
		s := digitSum(num)
		if s < minVal {
			minVal = s
		}
	}
	return minVal
}
```

## 3304 — Find The K Th Character In String Game I

```go
package main

// LeetCode #3304: Find the K-th Character in String Game I
// https://leetcode.com/problems/find-the-k-th-character-in-string-game-i/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(string(FindTheKThCharacterInStringGameI(5)))
	fmt.Println(string(FindTheKThCharacterInStringGameI(10)))
	fmt.Println(string(FindTheKThCharacterInStringGameI(1)))
}

// FindTheKThCharacterInStringGameI returns the k-th character in the string game.
// Start with "a", each step append the incremented version of the current string.
// Time: O(log k). Space: O(1).
func FindTheKThCharacterInStringGameI(k int) byte {
	// The string doubles in length each step. k is 1-indexed.
	// The character at position k depends on how many times we wrap around.
	k-- // 0-indexed
	count := 0
	for k > 0 {
		// Find the largest power of 2 <= k
		msb := 1
		for msb*2 <= k {
			msb *= 2
		}
		k -= msb
		count++
	}
	return byte('a' + byte(count%26))
}
```

## 3314 — Construct The Minimum Bitwise Array I

```go
package main

// LeetCode #3314: Construct the Minimum Bitwise Array I
// https://leetcode.com/problems/construct-the-minimum-bitwise-array-i/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ConstructTheMinimumBitwiseArrayI([]int{2, 3, 5, 7}))
	fmt.Println(ConstructTheMinimumBitwiseArrayI([]int{11, 13, 31}))
}

// ConstructTheMinimumBitwiseArrayI returns an array where ans[i] is the smallest number such that ans[i] | (ans[i]+1) == nums[i].
// Time: O(n * min_val). Space: O(n).
func ConstructTheMinimumBitwiseArrayI(nums []int) []int {
	result := make([]int, len(nums))
	for i, num := range nums {
		found := false
		for candidate := 0; candidate < num; candidate++ {
			if candidate|(candidate+1) == num {
				result[i] = candidate
				found = true
				break
			}
		}
		if !found {
			result[i] = -1
		}
	}
	return result
}
```

## 3318 — Find X Sum Of All K Long Subarrays I

```go
package main

// LeetCode #3318: Find X-Sum of All K-Long Subarrays I
// https://leetcode.com/problems/find-x-sum-of-all-k-long-subarrays-i/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(FindXSumOfAllKLongSubarraysI([]int{1, 1, 2, 2, 3, 4, 2, 3}, 6, 2))
	fmt.Println(FindXSumOfAllKLongSubarraysI([]int{3, 8, 7, 8, 7, 5}, 2, 2))
}

// FindXSumOfAllKLongSubarraysI calculates x-sum for each k-length subarray.
// The x-sum is the sum of the top x most frequent elements (breaking ties by larger value).
// Time: O(n * k log k). Space: O(k).
func FindXSumOfAllKLongSubarraysI(nums []int, k int, x int) []int {
	n := len(nums)
	result := make([]int, n-k+1)
	for start := 0; start <= n-k; start++ {
		freq := make(map[int]int)
		for i := start; i < start+k; i++ {
			freq[nums[i]]++
		}

		type pair struct {
			val int
			cnt int
		}
		pairs := make([]pair, 0, len(freq))
		for val, cnt := range freq {
			pairs = append(pairs, pair{val, cnt})
		}
		sort.Slice(pairs, func(i, j int) bool {
			if pairs[i].cnt != pairs[j].cnt {
				return pairs[i].cnt > pairs[j].cnt
			}
			return pairs[i].val > pairs[j].val
		})

		sum := 0
		for i := 0; i < x && i < len(pairs); i++ {
			sum += pairs[i].val * pairs[i].cnt
		}
		result[start] = sum
	}
	return result
}
```

## 3330 — Find The Original Typed String I

```go
package main

// LeetCode #3330: Find the Original Typed String I
// https://leetcode.com/problems/find-the-original-typed-string-i/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindTheOriginalTypedStringI("aabbccdd"))
	fmt.Println(FindTheOriginalTypedStringI("aaaa"))
	fmt.Println(FindTheOriginalTypedStringI("abc"))
}

// FindTheOriginalTypedStringI counts possible original strings where adjacent equal characters could be merged.
// Time: O(n). Space: O(1).
func FindTheOriginalTypedStringI(word string) int {
	count := 1
	streak := 1
	for i := 1; i < len(word); i++ {
		if word[i] == word[i-1] {
			streak++
		} else {
			streak = 1
		}
		if streak >= 2 {
			// If we have at least 2 of the same char consecutively, we can type fewer
			count++
		}
	}
	return count
}
```

## 3340 — Check Balanced String

```go
package main

// LeetCode #3340: Check Balanced String
// https://leetcode.com/problems/check-balanced-string/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CheckBalancedString("1234"))
	fmt.Println(CheckBalancedString("24123"))
}

// CheckBalancedString returns true if sum of digits at even positions equals sum at odd positions.
// Time: O(n). Space: O(1).
func CheckBalancedString(num string) bool {
	evenSum, oddSum := 0, 0
	for i, ch := range num {
		digit := int(ch - '0')
		if i%2 == 0 {
			evenSum += digit
		} else {
			oddSum += digit
		}
	}
	return evenSum == oddSum
}
```

## 3345 — Smallest Divisible Digit Product I

```go
package main

// LeetCode #3345: Smallest Divisible Digit Product I
// https://leetcode.com/problems/smallest-divisible-digit-product-i/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(SmallestDivisibleDigitProductI(10, 2))
	fmt.Println(SmallestDivisibleDigitProductI(15, 3))
}

// digitProduct returns the product of digits of n.
func digitProduct(n int) int {
	product := 1
	for n > 0 {
		product *= n % 10
		n /= 10
	}
	return product
}

// SmallestDivisibleDigitProductI returns the smallest number >= n whose digit product is divisible by t.
// Time: O(answer * log n). Space: O(1).
func SmallestDivisibleDigitProductI(n int, t int) int {
	for {
		dp := digitProduct(n)
		if dp%t == 0 {
			return n
		}
		n++
	}
}
```

## 3349 — Adjacent Increasing Subarrays Detection I

```go
package main

// LeetCode #3349: Adjacent Increasing Subarrays Detection I
// https://leetcode.com/problems/adjacent-increasing-subarrays-detection-i/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(AdjacentIncreasingSubarraysDetectionI([]int{2, 5, 7, 8, 9, 2, 3, 4, 3, 1}, 3))
	fmt.Println(AdjacentIncreasingSubarraysDetectionI([]int{1, 2, 3, 4, 4, 4, 4, 5, 6, 7}, 5))
}

// AdjacentIncreasingSubarraysDetectionI returns true if there exist two adjacent k-length increasing subarrays.
// Time: O(n). Space: O(n).
func AdjacentIncreasingSubarraysDetectionI(nums []int, k int) bool {
	n := len(nums)
	if n < 2*k {
		return false
	}

	// inc[i] = true if subarray starting at i of length k is strictly increasing
	inc := make([]bool, n-k+1)
	for i := 0; i <= n-k; i++ {
		isInc := true
		for j := i; j < i+k-1; j++ {
			if nums[j] >= nums[j+1] {
				isInc = false
				break
			}
		}
		inc[i] = isInc
	}

	for i := 0; i <= n-2*k; i++ {
		if inc[i] && inc[i+k] {
			return true
		}
	}
	return false
}
```

## 3353 — Minimum Total Operations

```go
package main

// LeetCode #3353: Minimum Total Operations
// https://leetcode.com/problems/minimum-total-operations/
// Difficulty: Easy [Paid]

import "fmt"

func main() {
	fmt.Println(MinimumTotalOperations([]int{1, 2, 3, 4}))
	fmt.Println(MinimumTotalOperations([]int{1, 1, 1}))
}

// MinimumTotalOperations returns the minimum number of operations to make all elements zero.
// Each operation picks a subarray and subtracts the minimum value in it from all its elements.
// Time: O(n). Space: O(1).
func MinimumTotalOperations(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	ops := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			ops += nums[i] - nums[i-1]
		}
	}
	return ops
}
```

## 3354 — Make Array Elements Equal To Zero

```go
package main

// LeetCode #3354: Make Array Elements Equal to Zero
// https://leetcode.com/problems/make-array-elements-equal-to-zero/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MakeArrayElementsEqualToZero([]int{1, 0, 2, 0, 3}))
	fmt.Println(MakeArrayElementsEqualToZero([]int{2, 3, 4, 0, 4, 1, 0}))
}

// MakeArrayElementsEqualToZero counts valid starting positions to make all elements zero.
// From a starting position, move left/right and decrement each visited element by 1.
// Time: O(n^2). Space: O(1).
func MakeArrayElementsEqualToZero(nums []int) int {
	n := len(nums)
	count := 0

	for start := 0; start < n; start++ {
		arr := make([]int, n)
		copy(arr, nums)

		pos := start
		dir := -1 // start going left
		allZero := true
		for _, v := range arr {
			if v != 0 {
				allZero = false
				break
			}
		}
		if allZero {
			count++
			continue
		}

		moves := 0
		for moves < 100000 {
			if arr[pos] > 0 {
				arr[pos]--
			}
			// Check if all zero
			zero := true
			for _, v := range arr {
				if v != 0 {
					zero = false
					break
				}
			}
			if zero {
				count++
				break
			}
			// Move
			nextPos := pos + dir
			if nextPos < 0 || nextPos >= n {
				dir = -dir
			} else {
				pos = nextPos
			}
			moves++
		}
	}
	return count
}
```

## 3358 — Books With Null Ratings

```go
package main

// LeetCode #3358: Books with NULL Ratings
// https://leetcode.com/problems/books-with-null-ratings/
// Difficulty: Easy [Paid]

import (
	"fmt"
	"sort"
)

func main() {
	books := []Book{
		{BookID: 1, Title: "Book A", Author: "Author X", PublishedYear: 2020, Rating: nil},
		{BookID: 2, Title: "Book B", Author: "Author Y", PublishedYear: 2021, Rating: ptr(4)},
		{BookID: 3, Title: "Book C", Author: "Author Z", PublishedYear: 2019, Rating: nil},
	}
	result := BooksWithNullRatings(books)
	for _, b := range result {
		fmt.Println(b)
	}
}

func ptr(i int) *int { return &i }

// Book represents a book with optional rating.
type Book struct {
	BookID        int
	Title         string
	Author        string
	PublishedYear int
	Rating        *int
}

// BooksWithNullRatings returns books that have NULL ratings, sorted by book_id.
// Time: O(n log n). Space: O(n).
func BooksWithNullRatings(books []Book) []Book {
	result := []Book{}
	for _, b := range books {
		if b.Rating == nil {
			result = append(result, b)
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].BookID < result[j].BookID
	})
	return result
}
```

## 3360 — Stone Removal Game

```go
package main

// LeetCode #3360: Stone Removal Game
// https://leetcode.com/problems/stone-removal-game/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(StoneRemovalGame(10))
	fmt.Println(StoneRemovalGame(7))
}

// StoneRemovalGame returns true if Alice wins the stone removal game.
// Alice goes first; they remove stones starting from 1 and increase by 1 each turn.
// Alice wins if she can make the last move.
// Time: O(sqrt(n)). Space: O(1).
func StoneRemovalGame(n int) bool {
	turn := 0 // 0 for Alice, 1 for Bob
	remove := 1
	for n >= remove {
		n -= remove
		remove++
		turn = 1 - turn
	}
	// If Alice made the last move, Alice wins, else Bob wins
	return turn != 0
}
```

## 3364 — Minimum Positive Sum Subarray

```go
package main

// LeetCode #3364: Minimum Positive Sum Subarray
// https://leetcode.com/problems/minimum-positive-sum-subarray/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MinimumPositiveSumSubarray([]int{3, -2, 1, 4}, 2, 3))
	fmt.Println(MinimumPositiveSumSubarray([]int{-2, 2, -2, 2}, 1, 2))
	fmt.Println(MinimumPositiveSumSubarray([]int{1, 2, 3, 4}, 2, 4))
}

// MinimumPositiveSumSubarray returns the minimum positive sum of any subarray with length between l and r.
// Time: O(n * (r-l+1)). Space: O(1).
func MinimumPositiveSumSubarray(nums []int, l int, r int) int {
	n := len(nums)
	minPos := -1
	for length := l; length <= r; length++ {
		for start := 0; start <= n-length; start++ {
			sum := 0
			for i := start; i < start+length; i++ {
				sum += nums[i]
			}
			if sum > 0 && (minPos == -1 || sum < minPos) {
				minPos = sum
			}
		}
	}
	return minPos
}
```

## 3370 — Smallest Number With All Set Bits

```go
package main

// LeetCode #3370: Smallest Number With All Set Bits
// https://leetcode.com/problems/smallest-number-with-all-set-bits/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(SmallestNumberWithAllSetBits(5))
	fmt.Println(SmallestNumberWithAllSetBits(10))
	fmt.Println(SmallestNumberWithAllSetBits(3))
}

// SmallestNumberWithAllSetBits returns the smallest number >= n whose binary representation consists of all 1s.
// Time: O(log n). Space: O(1).
func SmallestNumberWithAllSetBits(n int) int {
	result := 1
	for result < n {
		result = (result << 1) | 1
	}
	return result
}
```

## 3375 — Minimum Operations To Make Array Values Equal To K

```go
package main

// LeetCode #3375: Minimum Operations to Make Array Values Equal to K
// https://leetcode.com/problems/minimum-operations-to-make-array-values-equal-to-k/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MinimumOperationsToMakeArrayValuesEqualToK([]int{5, 2, 5, 4, 5}, 2))
	fmt.Println(MinimumOperationsToMakeArrayValuesEqualToK([]int{2, 1, 2}, 2))
	fmt.Println(MinimumOperationsToMakeArrayValuesEqualToK([]int{9, 7, 5, 3}, 1))
}

// MinimumOperationsToMakeArrayValuesEqualToK returns the minimum operations to reduce all numbers to k.
// In one operation, you can change any number > x to x for some x.
// Time: O(n). Space: O(n).
func MinimumOperationsToMakeArrayValuesEqualToK(nums []int, k int) int {
	minVal := nums[0]
	for _, v := range nums {
		if v < minVal {
			minVal = v
		}
	}
	if minVal < k {
		return -1
	}

	seen := make(map[int]bool)
	for _, v := range nums {
		if v > k {
			seen[v] = true
		}
	}
	return len(seen)
}
```

## 3379 — Transformed Array

```go
package main

// LeetCode #3379: Transformed Array
// https://leetcode.com/problems/transformed-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(TransformedArray([]int{3, -2, 1, 1}))
	fmt.Println(TransformedArray([]int{-1, 4, -1}))
}

// TransformedArray constructs a new array where result[i] = nums[(i + nums[i]) mod n], handling negative wrap-around.
// Time: O(n). Space: O(n).
func TransformedArray(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	for i, val := range nums {
		idx := (i + val) % n
		if idx < 0 {
			idx += n
		}
		result[i] = nums[idx]
	}
	return result
}
```

## 3386 — Button With Longest Push Time

```go
package main

// LeetCode #3386: Button with Longest Push Time
// https://leetcode.com/problems/button-with-longest-push-time/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ButtonWithLongestPushTime([][]int{{1, 2}, {2, 5}, {3, 9}, {1, 15}}))
	fmt.Println(ButtonWithLongestPushTime([][]int{{10, 5}, {1, 7}}))
}

// ButtonWithLongestPushTime returns the button index with the longest duration between consecutive events.
// Each event is [button_index, timestamp].
// Time: O(n). Space: O(1).
func ButtonWithLongestPushTime(events [][]int) int {
	maxDuration := 0
	buttonIndex := events[0][0]
	prevTime := events[0][1]

	for i := 1; i < len(events); i++ {
		duration := events[i][1] - prevTime
		if duration > maxDuration || (duration == maxDuration && events[i][0] < buttonIndex) {
			maxDuration = duration
			buttonIndex = events[i][0]
		}
		prevTime = events[i][1]
	}
	return buttonIndex
}
```

## 3392 — Count Subarrays Of Length Three With A Condition

```go
package main

// LeetCode #3392: Count Subarrays of Length Three With a Condition
// https://leetcode.com/problems/count-subarrays-of-length-three-with-a-condition/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountSubarraysOfLengthThreeWithACondition([]int{1, 2, 1, 2, 1}))
	fmt.Println(CountSubarraysOfLengthThreeWithACondition([]int{1, 3, 5, 7, 9}))
}

// CountSubarraysOfLengthThreeWithACondition counts subarrays of length 3 where the sum of first and last equals the middle.
// Time: O(n). Space: O(1).
func CountSubarraysOfLengthThreeWithACondition(nums []int) int {
	count := 0
	for i := 0; i <= len(nums)-3; i++ {
		if nums[i]+nums[i+2] == nums[i+1] {
			count++
		}
	}
	return count
}
```

## 3396 — Minimum Number Of Operations To Make Elements In Array Distinct

```go
package main

// LeetCode #3396: Minimum Number of Operations to Make Elements in Array Distinct
// https://leetcode.com/problems/minimum-number-of-operations-to-make-elements-in-array-distinct/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MinimumNumberOfOperationsToMakeElementsInArrayDistinct([]int{1, 2, 3, 4, 2, 3, 3, 5, 7}))
	fmt.Println(MinimumNumberOfOperationsToMakeElementsInArrayDistinct([]int{4, 5, 6, 4, 4}))
	fmt.Println(MinimumNumberOfOperationsToMakeElementsInArrayDistinct([]int{6, 7, 8, 9}))
}

// MinimumNumberOfOperationsToMakeElementsInArrayDistinct returns minimum operations to make all elements distinct.
// Each operation removes the first 3 elements (or all remaining if < 3).
// Time: O(n). Space: O(n).
func MinimumNumberOfOperationsToMakeElementsInArrayDistinct(nums []int) int {
	ops := 0
	for {
		seen := make(map[int]bool)
		distinct := true
		for _, v := range nums {
			if seen[v] {
				distinct = false
				break
			}
			seen[v] = true
		}
		if distinct || len(nums) == 0 {
			return ops
		}
		// Remove first 3 elements
		if len(nums) <= 3 {
			nums = []int{}
		} else {
			nums = nums[3:]
		}
		ops++
	}
}
```

## 3402 — Minimum Operations To Make Columns Strictly Increasing

```go
package main

// LeetCode #3402: Minimum Operations to Make Columns Strictly Increasing
// https://leetcode.com/problems/minimum-operations-to-make-columns-strictly-increasing/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MinimumOperationsToMakeColumnsStrictlyIncreasing([][]int{{3, 2}, {1, 3}, {3, 4}, {0, 1}}))
	fmt.Println(MinimumOperationsToMakeColumnsStrictlyIncreasing([][]int{{3, 2, 1}, {2, 1, 0}, {1, 2, 3}}))
}

// MinimumOperationsToMakeColumnsStrictlyIncreasing returns minimum operations to make each column strictly increasing.
// Each operation increments an element by 1.
// Time: O(n * m). Space: O(1).
func MinimumOperationsToMakeColumnsStrictlyIncreasing(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	rows, cols := len(grid), len(grid[0])
	ops := 0
	for c := 0; c < cols; c++ {
		for r := 1; r < rows; r++ {
			if grid[r][c] <= grid[r-1][c] {
				diff := grid[r-1][c] - grid[r][c] + 1
				ops += diff
				grid[r][c] += diff
			}
		}
	}
	return ops
}
```

## 3407 — Substring Matching Pattern

```go
package main

// LeetCode #3407: Substring Matching Pattern
// https://leetcode.com/problems/substring-matching-pattern/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(SubstringMatchingPattern("leetcode", "ee*e"))
	fmt.Println(SubstringMatchingPattern("car", "c*r"))
	fmt.Println(SubstringMatchingPattern("test", "t*t"))
}

// SubstringMatchingPattern checks if s matches pattern p where '*' matches any sequence of characters.
// Time: O(n * m). Space: O(n).
func SubstringMatchingPattern(s string, p string) bool {
	starIdx := -1
	for i, ch := range p {
		if ch == '*' {
			starIdx = i
			break
		}
	}

	left := p[:starIdx]
	right := p[starIdx+1:]

	// Left part must be prefix of some substring, right part must be suffix
	return strings.Contains(s, left) && strings.Contains(s, right) &&
		strings.Index(s, left) <= len(s)-len(right) &&
		strings.Index(s, left)+len(left) <= strings.LastIndex(s, right)
}
```

## 3411 — Maximum Subarray With Equal Products

```go
package main

// LeetCode #3411: Maximum Subarray With Equal Products
// https://leetcode.com/problems/maximum-subarray-with-equal-products/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MaximumSubarrayWithEqualProducts([]int{1, 2, 1, 2, 1, 1, 1}))
	fmt.Println(MaximumSubarrayWithEqualProducts([]int{2, 3, 4, 5, 6}))
}

// gcd returns the greatest common divisor of a and b.
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// MaximumSubarrayWithEqualProducts returns the length of the longest subarray where product == lcm * gcd.
// Time: O(n^2). Space: O(1).
func MaximumSubarrayWithEqualProducts(nums []int) int {
	n := len(nums)
	maxLen := 0
	for i := 0; i < n; i++ {
		p := 1
		g := nums[i]
		l := nums[i]
		for j := i; j < n; j++ {
			p *= nums[j]
			g = gcd(g, nums[j])
			l = l * nums[j] / gcd(l, nums[j])
			if p == l*g {
				if j-i+1 > maxLen {
					maxLen = j - i + 1
				}
			}
		}
	}
	return maxLen
}
```

## 3415 — Find Products With Three Consecutive Digits

```go
package main

// LeetCode #3415: Find Products with Three Consecutive Digits
// https://leetcode.com/problems/find-products-with-three-consecutive-digits/
// Difficulty: Easy [Paid]

import (
	"fmt"
	"regexp"
)

func main() {
	products := []Product{
		{ProductID: 1, Name: "ABC123XYZ"},
		{ProductID: 2, Name: "Product456"},
		{ProductID: 3, Name: "Item12"},
	}
	result := FindProductsWithThreeConsecutiveDigits(products)
	for _, p := range result {
		fmt.Printf("%d: %s\n", p.ProductID, p.Name)
	}
}

// Product represents a product.
type Product struct {
	ProductID int
	Name      string
}

// FindProductsWithThreeConsecutiveDigits returns products whose name contains at least three consecutive digits.
// Time: O(n). Space: O(n).
func FindProductsWithThreeConsecutiveDigits(products []Product) []Product {
	re := regexp.MustCompile(`\d{3,}`)
	result := []Product{}
	for _, p := range products {
		if re.MatchString(p.Name) {
			result = append(result, p)
		}
	}
	return result
}
```

## 3417 — Zigzag Grid Traversal With Skip

```go
package main

// LeetCode #3417: Zigzag Grid Traversal With Skip
// https://leetcode.com/problems/zigzag-grid-traversal-with-skip/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ZigzagGridTraversalWithSkip([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	fmt.Println(ZigzagGridTraversalWithSkip([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}}))
}

// ZigzagGridTraversalWithSkip traverses the grid in zigzag order but skips every other element.
// Time: O(m * n). Space: O(m * n).
func ZigzagGridTraversalWithSkip(grid [][]int) []int {
	if len(grid) == 0 {
		return nil
	}
	m, n := len(grid), len(grid[0])
	result := []int{}
	skip := false
	for i := 0; i < m; i++ {
		if i%2 == 0 {
			for j := 0; j < n; j++ {
				if !skip {
					result = append(result, grid[i][j])
				}
				skip = !skip
			}
		} else {
			for j := n - 1; j >= 0; j-- {
				if !skip {
					result = append(result, grid[i][j])
				}
				skip = !skip
			}
		}
	}
	return result
}
```

## 3423 — Maximum Difference Between Adjacent Elements In A Circular Array

```go
package main

// LeetCode #3423: Maximum Difference Between Adjacent Elements in a Circular Array
// https://leetcode.com/problems/maximum-difference-between-adjacent-elements-in-a-circular-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MaximumDifferenceBetweenAdjacentElementsInACircularArray([]int{1, 2, 4}))
	fmt.Println(MaximumDifferenceBetweenAdjacentElementsInACircularArray([]int{-5, -1, -3}))
}

// MaximumDifferenceBetweenAdjacentElementsInACircularArray returns the max absolute diff between adjacent elements (circular).
// Time: O(n). Space: O(1).
func MaximumDifferenceBetweenAdjacentElementsInACircularArray(nums []int) int {
	n := len(nums)
	maxDiff := 0
	for i := 0; i < n; i++ {
		diff := nums[i] - nums[(i+1)%n]
		if diff < 0 {
			diff = -diff
		}
		if diff > maxDiff {
			maxDiff = diff
		}
	}
	return maxDiff
}
```

## 3427 — Sum Of Variable Length Subarrays

```go
package main

// LeetCode #3427: Sum of Variable Length Subarrays
// https://leetcode.com/problems/sum-of-variable-length-subarrays/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(SumOfVariableLengthSubarrays([]int{2, 3, 1}))
	fmt.Println(SumOfVariableLengthSubarrays([]int{3, 1, 1, 2}))
}

// SumOfVariableLengthSubarrays computes sum of subarrays where each subarray starts at i-(nums[i]%something) and ends at i.
// Time: O(n^2). Space: O(1).
func SumOfVariableLengthSubarrays(nums []int) int {
	n := len(nums)
	total := 0
	for i := 0; i < n; i++ {
		start := i - nums[i]
		if start < 0 {
			start = 0
		}
		for j := start; j <= i; j++ {
			total += nums[j]
		}
	}
	return total
}
```

## 3432 — Count Partitions With Even Sum Difference

```go
package main

// LeetCode #3432: Count Partitions with Even Sum Difference
// https://leetcode.com/problems/count-partitions-with-even-sum-difference/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountPartitionsWithEvenSumDifference([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(CountPartitionsWithEvenSumDifference([]int{10, 10, 10, 10, 10}))
}

// CountPartitionsWithEvenSumDifference counts partitions where the difference between left and right sums is even.
// Time: O(n). Space: O(1).
func CountPartitionsWithEvenSumDifference(nums []int) int {
	totalSum := 0
	for _, v := range nums {
		totalSum += v
	}
	leftSum := 0
	count := 0
	for i := 0; i < len(nums)-1; i++ {
		leftSum += nums[i]
		rightSum := totalSum - leftSum
		if (leftSum-rightSum)%2 == 0 {
			count++
		}
	}
	return count
}
```

## 3436 — Find Valid Emails

```go
package main

// LeetCode #3436: Find Valid Emails
// https://leetcode.com/problems/find-valid-emails/
// Difficulty: Easy

import (
	"fmt"
	"regexp"
)

func main() {
	users := []User{
		{UserID: 1, Email: "alice@example.com"},
		{UserID: 2, Email: "bob@example"},
		{UserID: 3, Email: "@example.com"},
		{UserID: 4, Email: "charlie@example.com"},
	}
	result := FindValidEmails(users)
	for _, r := range result {
		fmt.Printf("%d: %s\n", r.UserID, r.Email)
	}
}

// User represents a user with an email.
type User struct {
	UserID int
	Email  string
}

// FindValidEmails returns users with valid email addresses (alphanumeric prefix, letter-only domain, .com suffix).
// Time: O(n). Space: O(n).
func FindValidEmails(users []User) []User {
	re := regexp.MustCompile(`^[A-Za-z0-9_]+@[A-Za-z]+\.com$`)
	result := []User{}
	for _, u := range users {
		if re.MatchString(u.Email) {
			result = append(result, u)
		}
	}
	return result
}
```

## 3438 — Find Valid Pair Of Adjacent Digits In String

```go
package main

// LeetCode #3438: Find Valid Pair of Adjacent Digits in String
// https://leetcode.com/problems/find-valid-pair-of-adjacent-digits-in-string/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindValidPairOfAdjacentDigitsInString("2523533"))
	fmt.Println(FindValidPairOfAdjacentDigitsInString("111"))
}

// FindValidPairOfAdjacentDigitsInString finds the first pair of adjacent equal digits where the digit's frequency > the digit.
// Time: O(n). Space: O(n).
func FindValidPairOfAdjacentDigitsInString(s string) string {
	freq := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		freq[s[i]]++
	}
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			cnt := freq[s[i]]
			if cnt > int(s[i]-'0') {
				return s[i : i+2]
			}
		}
	}
	return ""
}
```

## 3442 — Maximum Difference Between Even And Odd Frequency I

```go
package main

// LeetCode #3442: Maximum Difference Between Even and Odd Frequency I
// https://leetcode.com/problems/maximum-difference-between-even-and-odd-frequency-i/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MaximumDifferenceBetweenEvenAndOddFrequencyI("aaaaabbc"))
	fmt.Println(MaximumDifferenceBetweenEvenAndOddFrequencyI("abcabcab"))
}

// MaximumDifferenceBetweenEvenAndOddFrequencyI returns the max difference between max even-frequency and min odd-frequency in s.
// Time: O(n). Space: O(1).
func MaximumDifferenceBetweenEvenAndOddFrequencyI(s string) int {
	freq := make([]int, 26)
	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
	}

	maxEven := 0
	minOdd := -1
	for _, f := range freq {
		if f == 0 {
			continue
		}
		if f%2 == 0 && f > maxEven {
			maxEven = f
		} else if f%2 == 1 && (minOdd == -1 || f < minOdd) {
			minOdd = f
		}
	}

	if maxEven == 0 || minOdd == -1 {
		return 0
	}
	return maxEven - minOdd
}
```

## 3450 — Maximum Students On A Single Bench

```go
package main

// LeetCode #3450: Maximum Students on a Single Bench
// https://leetcode.com/problems/maximum-students-on-a-single-bench/
// Difficulty: Easy [Paid]

import "fmt"

func main() {
	fmt.Println(MaximumStudentsOnASingleBench([][]int{{1, 2}, {2, 3}, {1, 3}, {1, 2}}))
	fmt.Println(MaximumStudentsOnASingleBench([][]int{{1, 1}, {2, 1}, {3, 1}}))
}

// MaximumStudentsOnASingleBench returns the maximum number of different students on any single bench.
// Each entry is [student_id, bench_id].
// Time: O(n). Space: O(n).
func MaximumStudentsOnASingleBench(students [][]int) int {
	benchStudents := make(map[int]map[int]bool)
	for _, s := range students {
		studentID, benchID := s[0], s[1]
		if benchStudents[benchID] == nil {
			benchStudents[benchID] = make(map[int]bool)
		}
		benchStudents[benchID][studentID] = true
	}
	maxCount := 0
	for _, students := range benchStudents {
		if len(students) > maxCount {
			maxCount = len(students)
		}
	}
	return maxCount
}
```

## 3452 — Sum Of Good Numbers

```go
package main

// LeetCode #3452: Sum of Good Numbers
// https://leetcode.com/problems/sum-of-good-numbers/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(SumOfGoodNumbers([]int{1, 3, 2, 1, 5, 4}, 2))
	fmt.Println(SumOfGoodNumbers([]int{2, 1}, 1))
}

// SumOfGoodNumbers returns sum of numbers that are greater than both nums[i-k] and nums[i+k] (or if out of bounds).
// Time: O(n). Space: O(1).
func SumOfGoodNumbers(nums []int, k int) int {
	n := len(nums)
	sum := 0
	for i := 0; i < n; i++ {
		good := true
		if i-k >= 0 && nums[i] <= nums[i-k] {
			good = false
		}
		if i+k < n && nums[i] <= nums[i+k] {
			good = false
		}
		if good {
			sum += nums[i]
		}
	}
	return sum
}
```

## 3456 — Find Special Substring Of Length K

```go
package main

// LeetCode #3456: Find Special Substring of Length K
// https://leetcode.com/problems/find-special-substring-of-length-k/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindSpecialSubstringOfLengthK("aaabaaa", 3))
	fmt.Println(FindSpecialSubstringOfLengthK("abc", 2))
}

// FindSpecialSubstringOfLengthK returns true if there is a substring of length k consisting of a single character, surrounded by different characters (or boundaries).
// Time: O(n). Space: O(1).
func FindSpecialSubstringOfLengthK(s string, k int) bool {
	n := len(s)
	for i := 0; i <= n-k; i++ {
		same := true
		for j := i; j < i+k-1; j++ {
			if s[j] != s[j+1] {
				same = false
				break
			}
		}
		if !same {
			continue
		}
		if i > 0 && s[i-1] == s[i] {
			continue
		}
		if i+k < n && s[i+k] == s[i] {
			continue
		}
		return true
	}
	return false
}
```

## 3461 — Check If Digits Are Equal In String After Operations I

```go
package main

// LeetCode #3461: Check If Digits Are Equal in String After Operations I
// https://leetcode.com/problems/check-if-digits-are-equal-in-string-after-operations-i/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CheckIfDigitsAreEqualInStringAfterOperationsI("1234"))
	fmt.Println(CheckIfDigitsAreEqualInStringAfterOperationsI("1111"))
}

// CheckIfDigitsAreEqualInStringAfterOperationsI repeatedly replaces adjacent digit pairs with (sum % 10) until 2 digits remain, then checks equality.
// Time: O(n^2). Space: O(n).
func CheckIfDigitsAreEqualInStringAfterOperationsI(s string) bool {
	digits := make([]int, len(s))
	for i, ch := range s {
		digits[i] = int(ch - '0')
	}

	for len(digits) > 2 {
		next := make([]int, len(digits)-1)
		for i := 0; i < len(digits)-1; i++ {
			next[i] = (digits[i] + digits[i+1]) % 10
		}
		digits = next
	}
	return digits[0] == digits[1]
}
```

## 3465 — Find Products With Valid Serial Numbers

```go
package main

// LeetCode #3465: Find Products with Valid Serial Numbers
// https://leetcode.com/problems/find-products-with-valid-serial-numbers/
// Difficulty: Easy

import (
	"fmt"
	"regexp"
)

func main() {
	products := []InvProduct{
		{ProductID: 1, Name: "Widget", SerialNumber: "SN-12345-ABC"},
		{ProductID: 2, Name: "Gadget", SerialNumber: "invalid"},
		{ProductID: 3, Name: "Doohickey", SerialNumber: "SN-67890-XYZ"},
	}
	result := FindProductsWithValidSerialNumbers(products)
	for _, p := range result {
		fmt.Printf("%d: %s (%s)\n", p.ProductID, p.Name, p.SerialNumber)
	}
}

// InvProduct represents a product with a serial number.
type InvProduct struct {
	ProductID    int
	Name         string
	SerialNumber string
}

// FindProductsWithValidSerialNumbers returns products whose serial number matches a valid pattern.
// Time: O(n). Space: O(n).
func FindProductsWithValidSerialNumbers(products []InvProduct) []InvProduct {
	re := regexp.MustCompile(`^SN-\d{5}-[A-Z]{3}$`)
	result := []InvProduct{}
	for _, p := range products {
		if re.MatchString(p.SerialNumber) {
			result = append(result, p)
		}
	}
	return result
}
```

## 3467 — Transform Array By Parity

```go
package main

// LeetCode #3467: Transform Array by Parity
// https://leetcode.com/problems/transform-array-by-parity/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(TransformArrayByParity([]int{4, 3, 2, 1}))
	fmt.Println(TransformArrayByParity([]int{1, 5, 2, 8, 3}))
}

// TransformArrayByParity transforms array: even numbers -> 0 (sorted first), odd numbers -> 1.
// Time: O(n log n). Space: O(1).
func TransformArrayByParity(nums []int) []int {
	// Count evens
	evenCount := 0
	for _, v := range nums {
		if v%2 == 0 {
			evenCount++
		}
	}
	result := make([]int, len(nums))
	for i := 0; i < evenCount; i++ {
		result[i] = 0
	}
	for i := evenCount; i < len(nums); i++ {
		result[i] = 1
	}
	return result
}
```

## 3471 — Find The Largest Almost Missing Integer

```go
package main

// LeetCode #3471: Find the Largest Almost Missing Integer
// https://leetcode.com/problems/find-the-largest-almost-missing-integer/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindTheLargestAlmostMissingInteger([]int{3, 9, 2, 3, 1, 6, 7, 8, 9}, 2))
	fmt.Println(FindTheLargestAlmostMissingInteger([]int{0, 0}, 1))
}

// FindTheLargestAlmostMissingInteger returns the largest integer that appears fewer than k times in nums.
// Time: O(n). Space: O(n).
func FindTheLargestAlmostMissingInteger(nums []int, k int) int {
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}
	largest := -1
	for val, count := range freq {
		if count < k && val > largest {
			largest = val
		}
	}
	return largest
}
```

## 3477 — Fruits Into Baskets Ii

```go
package main

// LeetCode #3477: Fruits Into Baskets II
// https://leetcode.com/problems/fruits-into-baskets-ii/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FruitsIntoBasketsIi([]int{4, 2, 5}, []int{3, 5, 4}))
	fmt.Println(FruitsIntoBasketsIi([]int{3, 6, 1}, []int{6, 4, 7}))
}

// FruitsIntoBasketsIi counts fruits that cannot be placed into baskets.
// Each fruit i can go into basket j if fruits[i] <= baskets[j].
// Time: O(n * m). Space: O(1).
func FruitsIntoBasketsIi(fruits []int, baskets []int) int {
	used := make([]bool, len(baskets))
	unplaced := 0
	for _, f := range fruits {
		placed := false
		for j, b := range baskets {
			if !used[j] && f <= b {
				used[j] = true
				placed = true
				break
			}
		}
		if !placed {
			unplaced++
		}
	}
	return unplaced
}
```

## 3483 — Unique 3 Digit Even Numbers

```go
package main

// LeetCode #3483: Unique 3-Digit Even Numbers
// https://leetcode.com/problems/unique-3-digit-even-numbers/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(UniqueThreeDigitEvenNumbers([]int{1, 2, 3, 4}))
	fmt.Println(UniqueThreeDigitEvenNumbers([]int{0, 2, 2}))
}

// UniqueThreeDigitEvenNumbers counts unique 3-digit even numbers that can be formed from digits (no leading zero).
// Time: O(n^3). Space: O(n).
func UniqueThreeDigitEvenNumbers(digits []int) int {
	used := make(map[int]bool)
	n := len(digits)
	for i := 0; i < n; i++ {
		if digits[i] == 0 {
			continue
		}
		for j := 0; j < n; j++ {
			if j == i {
				continue
			}
			for k := 0; k < n; k++ {
				if k == i || k == j {
					continue
				}
				num := digits[i]*100 + digits[j]*10 + digits[k]
				if num%2 == 0 {
					used[num] = true
				}
			}
		}
	}
	return len(used)
}
```

## 3487 — Maximum Unique Subarray Sum After Deletion

```go
package main

// LeetCode #3487: Maximum Unique Subarray Sum After Deletion
// https://leetcode.com/problems/maximum-unique-subarray-sum-after-deletion/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MaximumUniqueSubarraySumAfterDeletion([]int{1, 2, 3, 4, 5}))
	fmt.Println(MaximumUniqueSubarraySumAfterDeletion([]int{1, 1, 0, 1, 1}))
	fmt.Println(MaximumUniqueSubarraySumAfterDeletion([]int{1, 2, -1, -2, 1, 0, -1}))
}

// MaximumUniqueSubarraySumAfterDeletion returns max sum of a subarray with all unique elements, after optionally deleting one element.
// Time: O(n). Space: O(n).
func MaximumUniqueSubarraySumAfterDeletion(nums []int) int {
	// After deletion means we can delete any element, then find max sum of a subarray with distinct elements.
	// Equivalent to finding max sum of a subarray with at most one duplicate.
	// Simplified: find max sum subarray where all elements are distinct (allow deleting one element).
	// We'll use sliding window that allows one "skip" (a duplicate that we can delete).
	maxSum := 0
	for i := 0; i < len(nums); i++ {
		seen := make(map[int]int)
		sum := 0
		dupAllowed := true
		for j := i; j < len(nums); j++ {
			seen[nums[j]]++
			if seen[nums[j]] > 2 {
				break
			}
			if seen[nums[j]] == 2 {
				if dupAllowed {
					dupAllowed = false
				} else {
					break
				}
			}
			sum += nums[j]
			if sum > maxSum {
				maxSum = sum
			}
		}
	}
	return maxSum
}
```

## 3491 — Phone Number Prefix

```go
package main

// LeetCode #3491: Phone Number Prefix
// https://leetcode.com/problems/phone-number-prefix/
// Difficulty: Easy [Paid]

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(PhoneNumberPrefix([]string{"123", "1234", "567", "7890"}))
	fmt.Println(PhoneNumberPrefix([]string{"111", "222", "333"}))
}

// PhoneNumberPrefix returns true if no number is a prefix of another number.
// Time: O(n log n * m). Space: O(1).
func PhoneNumberPrefix(numbers []string) bool {
	sort.Strings(numbers)
	for i := 0; i < len(numbers)-1; i++ {
		if len(numbers[i]) <= len(numbers[i+1]) {
			isPrefix := true
			for j := 0; j < len(numbers[i]); j++ {
				if numbers[i][j] != numbers[i+1][j] {
					isPrefix = false
					break
				}
			}
			if isPrefix {
				return false
			}
		}
	}
	return true
}
```

## 3492 — Maximum Containers On A Ship

```go
package main

// LeetCode #3492: Maximum Containers on a Ship
// https://leetcode.com/problems/maximum-containers-on-a-ship/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MaximumContainersOnAShip(3, 2, 10))
	fmt.Println(MaximumContainersOnAShip(5, 3, 50))
}

// MaximumContainersOnAShip returns the maximum number of containers that can be loaded on an n x n ship, each weighing w, within maxWeight.
// Time: O(1). Space: O(1).
func MaximumContainersOnAShip(n int, w int, maxWeight int) int {
	maxBySpace := n * n
	maxByWeight := maxWeight / w
	if maxBySpace < maxByWeight {
		return maxBySpace
	}
	return maxByWeight
}
```

## 3498 — Reverse Degree Of A String

```go
package main

// LeetCode #3498: Reverse Degree of a String
// https://leetcode.com/problems/reverse-degree-of-a-string/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ReverseDegreeOfAString("abc"))
	fmt.Println(ReverseDegreeOfAString("zaba"))
}

// ReverseDegreeOfAString computes the sum of (position_in_reversed_alphabet * (i+1)) for each character.
// Reverse: a=26, b=25, ..., z=1.
// Time: O(n). Space: O(1).
func ReverseDegreeOfAString(s string) int {
	sum := 0
	for i, ch := range s {
		revPos := 26 - int(ch-'a')
		sum += revPos * (i + 1)
	}
	return sum
}
```

## 3502 — Minimum Cost To Reach Every Position

```go
package main

// LeetCode #3502: Minimum Cost to Reach Every Position
// https://leetcode.com/problems/minimum-cost-to-reach-every-position/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MinimumCostToReachEveryPosition([]int{5, 3, 4, 1, 3, 2}))
	fmt.Println(MinimumCostToReachEveryPosition([]int{1, 2, 3, 2, 1}))
}

// MinimumCostToReachEveryPosition returns an array where ans[i] is the minimum cost to reach position i.
// You can travel from j to i (j < i) at cost cost[i], or from i to j at cost cost[j].
// Time: O(n). Space: O(n).
func MinimumCostToReachEveryPosition(cost []int) []int {
	n := len(cost)
	result := make([]int, n)
	result[0] = cost[0]
	for i := 1; i < n; i++ {
		if cost[i] < result[i-1] {
			result[i] = cost[i]
		} else {
			result[i] = result[i-1]
		}
	}
	return result
}
```

## 3507 — Minimum Pair Removal To Sort Array I

```go
package main

// LeetCode #3507: Minimum Pair Removal to Sort Array I
// https://leetcode.com/problems/minimum-pair-removal-to-sort-array-i/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MinimumPairRemovalToSortArrayI([]int{5, 2, 3, 1}))
	fmt.Println(MinimumPairRemovalToSortArrayI([]int{1, 2, 3, 4}))
}

// MinimumPairRemovalToSortArrayI returns the minimum number of pairs to remove so the remaining array is sorted.
// A pair is two adjacent elements.
// Time: O(n^2). Space: O(n).
func MinimumPairRemovalToSortArrayI(nums []int) int {
	n := len(nums)
	// If already sorted
	sorted := true
	for i := 1; i < n; i++ {
		if nums[i] < nums[i-1] {
			sorted = false
			break
		}
	}
	if sorted {
		return 0
	}

	// Try removing pairs
	minOps := n
	var try func(arr []int, ops int)
	try = func(arr []int, ops int) {
		if ops >= minOps {
			return
		}
		// Check sorted
		ok := true
		for i := 1; i < len(arr); i++ {
			if arr[i] < arr[i-1] {
				ok = false
				break
			}
		}
		if ok {
			if ops < minOps {
				minOps = ops
			}
			return
		}
		if len(arr) < 2 {
			return
		}
		// Remove each possible adjacent pair
		for i := 0; i < len(arr)-1; i++ {
			next := make([]int, 0, len(arr)-2)
			next = append(next, arr[:i]...)
			next = append(next, arr[i+2:]...)
			try(next, ops+1)
		}
	}
	try(nums, 0)
	return minOps
}
```

## 3512 — Minimum Operations To Make Array Sum Divisible By K

```go
package main

// LeetCode #3512: Minimum Operations to Make Array Sum Divisible by K
// https://leetcode.com/problems/minimum-operations-to-make-array-sum-divisible-by-k/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MinimumOperationsToMakeArraySumDivisibleByK([]int{3, 9, 7}, 5))
	fmt.Println(MinimumOperationsToMakeArraySumDivisibleByK([]int{4, 1, 3}, 4))
}

// MinimumOperationsToMakeArraySumDivisibleByK returns the min operations (incrementing elements by 1) to make sum divisible by k.
// Time: O(n). Space: O(1).
func MinimumOperationsToMakeArraySumDivisibleByK(nums []int, k int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	rem := sum % k
	if rem == 0 {
		return 0
	}
	return rem
}
```

## 3516 — Find Closest Person

```go
package main

// LeetCode #3516: Find Closest Person
// https://leetcode.com/problems/find-closest-person/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindClosestPerson(1, 2, 3))
	fmt.Println(FindClosestPerson(1, 3, 2))
	fmt.Println(FindClosestPerson(2, 1, 3))
}

// FindClosestPerson finds who is closer to person z. Returns 0 for tie, 1 for person 1, 2 for person 2.
// Time: O(1). Space: O(1).
func FindClosestPerson(x int, y int, z int) int {
	dx := z - x
	if dx < 0 {
		dx = -dx
	}
	dy := z - y
	if dy < 0 {
		dy = -dy
	}
	if dx < dy {
		return 1
	} else if dy < dx {
		return 2
	}
	return 0
}
```

## 3536 — Maximum Product Of Two Digits

```go
package main

// LeetCode #3536: Maximum Product of Two Digits
// https://leetcode.com/problems/maximum-product-of-two-digits/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MaximumProductOfTwoDigits(34))
	fmt.Println(MaximumProductOfTwoDigits(10))
	fmt.Println(MaximumProductOfTwoDigits(99))
}

// MaximumProductOfTwoDigits returns the maximum product of any two digits in n.
// Time: O(log n). Space: O(1).
func MaximumProductOfTwoDigits(n int) int {
	digits := []int{}
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}
	maxProd := 0
	for i := 0; i < len(digits); i++ {
		for j := i + 1; j < len(digits); j++ {
			prod := digits[i] * digits[j]
			if prod > maxProd {
				maxProd = prod
			}
		}
	}
	return maxProd
}
```

## 3541 — Find Most Frequent Vowel And Consonant

```go
package main

// LeetCode #3541: Find Most Frequent Vowel and Consonant
// https://leetcode.com/problems/find-most-frequent-vowel-and-consonant/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindMostFrequentVowelAndConsonant("hello world"))
	fmt.Println(FindMostFrequentVowelAndConsonant("aabbccddee"))
}

// isVowel returns true if the byte is a lowercase vowel.
func isVowel(b byte) bool {
	return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u'
}

// FindMostFrequentVowelAndConsonant returns the sum of max vowel frequency and max consonant frequency.
// Time: O(n). Space: O(1).
func FindMostFrequentVowelAndConsonant(s string) int {
	vowelFreq := make([]int, 26)
	consonantFreq := make([]int, 26)
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch >= 'a' && ch <= 'z' {
			if isVowel(ch) {
				vowelFreq[ch-'a']++
			} else {
				consonantFreq[ch-'a']++
			}
		}
	}
	maxVowel := 0
	for _, v := range vowelFreq {
		if v > maxVowel {
			maxVowel = v
		}
	}
	maxConsonant := 0
	for _, v := range consonantFreq {
		if v > maxConsonant {
			maxConsonant = v
		}
	}
	return maxVowel + maxConsonant
}
```

## 3545 — Minimum Deletions For At Most K Distinct Characters

```go
package main

// LeetCode #3545: Minimum Deletions for At Most K Distinct Characters
// https://leetcode.com/problems/minimum-deletions-for-at-most-k-distinct-characters/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MinimumDeletionsForAtMostKDistinctCharacters("aabbbcc", 2))
	fmt.Println(MinimumDeletionsForAtMostKDistinctCharacters("abcde", 2))
}

// MinimumDeletionsForAtMostKDistinctCharacters returns min deletions so the string has at most k distinct characters.
// Time: O(n log n). Space: O(1).
func MinimumDeletionsForAtMostKDistinctCharacters(s string, k int) int {
	freq := make([]int, 26)
	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
	}
	sort.Slice(freq, func(i, j int) bool {
		return freq[i] > freq[j]
	})

	// Count distinct characters
	distinct := 0
	for _, f := range freq {
		if f > 0 {
			distinct++
		}
	}
	if distinct <= k {
		return 0
	}

	// Delete the least frequent characters (from the end of sorted freq)
	deletions := 0
	for i := len(freq) - 1; i >= 0 && distinct > k; i-- {
		if freq[i] > 0 {
			deletions += freq[i]
			distinct--
		}
	}
	return deletions
}
```

## 3550 — Smallest Index With Digit Sum Equal To Index

```go
package main

// LeetCode #3550: Smallest Index With Digit Sum Equal to Index
// https://leetcode.com/problems/smallest-index-with-digit-sum-equal-to-index/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(SmallestIndexWithDigitSumEqualToIndex([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}))
	fmt.Println(SmallestIndexWithDigitSumEqualToIndex([]int{10, 11, 12, 13, 14}))
}

// digitSum returns sum of digits.
func sumDigits(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

// SmallestIndexWithDigitSumEqualToIndex returns the smallest index i where digit sum of nums[i] equals i, else -1.
// Time: O(n * log max). Space: O(1).
func SmallestIndexWithDigitSumEqualToIndex(nums []int) int {
	for i, v := range nums {
		if sumDigits(v) == i {
			return i
		}
	}
	return -1
}
```

## 3560 — Find Minimum Log Transportation Cost

```go
package main

// LeetCode #3560: Find Minimum Log Transportation Cost
// https://leetcode.com/problems/find-minimum-log-transportation-cost/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindMinimumLogTransportationCost(10, 10, 3))
	fmt.Println(FindMinimumLogTransportationCost(5, 5, 2))
}

// FindMinimumLogTransportationCost returns the minimum transportation cost for logs of size n x m, cutting with factor k.
// Cost = max(0, max(n, m) - k) * k
// Time: O(1). Space: O(1).
func FindMinimumLogTransportationCost(n int, m int, k int) int {
	larger := n
	if m > larger {
		larger = m
	}
	if larger <= k {
		return 0
	}
	return (larger - k) * k
}
```

## 3570 — Find Books With No Available Copies

```go
package main

// LeetCode #3570: Find Books with No Available Copies
// https://leetcode.com/problems/find-books-with-no-available-copies/
// Difficulty: Easy [Paid]

import (
	"fmt"
	"sort"
)

func main() {
	books := []LibBook{
		{BookID: 1, Title: "Book A", AvailableCopies: 0},
		{BookID: 2, Title: "Book B", AvailableCopies: 3},
		{BookID: 3, Title: "Book C", AvailableCopies: 0},
	}
	result := FindBooksWithNoAvailableCopies(books)
	for _, b := range result {
		fmt.Printf("%d: %s\n", b.BookID, b.Title)
	}
}

// LibBook represents a library book.
type LibBook struct {
	BookID          int
	Title           string
	AvailableCopies int
}

// FindBooksWithNoAvailableCopies returns books with zero available copies, sorted by book_id.
// Time: O(n log n). Space: O(n).
func FindBooksWithNoAvailableCopies(books []LibBook) []LibBook {
	result := []LibBook{}
	for _, b := range books {
		if b.AvailableCopies == 0 {
			result = append(result, b)
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].BookID < result[j].BookID
	})
	return result
}
```

## 3571 — Find The Shortest Superstring Ii

```go
package main

// LeetCode #3571: Find the Shortest Superstring II
// https://leetcode.com/problems/find-the-shortest-superstring-ii/
// Difficulty: Easy [Paid]

import "fmt"

func main() {
	fmt.Println(FindTheShortestSuperstringIi("abc", "bcd"))
	fmt.Println(FindTheShortestSuperstringIi("abc", "xyz"))
}

// overlap returns the length of overlapping suffix of a that matches prefix of b.
func overlap(a, b string) int {
	maxOvlp := 0
	maxLen := len(a)
	if len(b) < maxLen {
		maxLen = len(b)
	}
	for i := 1; i <= maxLen; i++ {
		if a[len(a)-i:] == b[:i] {
			maxOvlp = i
		}
	}
	return maxOvlp
}

// FindTheShortestSuperstringIi returns the shortest string that contains both s1 and s2 as substrings.
// Time: O(n*m). Space: O(n+m).
func FindTheShortestSuperstringIi(s1 string, s2 string) string {
	ovlp12 := overlap(s1, s2)
	ovlp21 := overlap(s2, s1)

	s1s2 := s1 + s2[ovlp12:]
	s2s1 := s2 + s1[ovlp21:]

	if len(s1s2) < len(s2s1) {
		return s1s2
	} else if len(s2s1) < len(s1s2) {
		return s2s1
	}
	// Same length, return lexicographically smaller
	if s1s2 < s2s1 {
		return s1s2
	}
	return s2s1
}
```

## 3581 — Count Odd Letters From Number

```go
package main

// LeetCode #3581: Count Odd Letters from Number
// https://leetcode.com/problems/count-odd-letters-from-number/
// Difficulty: Easy [Paid]

import "fmt"

func main() {
	fmt.Println(CountOddLettersFromNumber(41))
	fmt.Println(CountOddLettersFromNumber(20))
	fmt.Println(CountOddLettersFromNumber(7))
}

// Time: O(log n) - number of digits of n
// Space: O(1)
func CountOddLettersFromNumber(n int) int {
	digitWords := [10]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	mask := 0
	for n > 0 {
		d := n % 10
		word := digitWords[d]
		for _, ch := range word {
			mask ^= 1 << (ch - 'a')
		}
		n /= 10
	}

	ans := 0
	for mask > 0 {
		ans += mask & 1
		mask >>= 1
	}
	return ans
}
```

## 3582 — Generate Tag For Video Caption

```go
package main

// LeetCode #3582: Generate Tag for Video Caption
// https://leetcode.com/problems/generate-tag-for-video-caption/
// Difficulty: Easy

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(GenerateTagForVideoCaption("Leetcode daily streak achieved"))
	fmt.Println(GenerateTagForVideoCaption("can I Go There"))
	fmt.Println(GenerateTagForVideoCaption("hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh"))
}

// Time: O(n)
// Space: O(n)
func GenerateTagForVideoCaption(caption string) string {
	words := strings.Fields(caption)
	for i, w := range words {
		if i == 0 {
			words[i] = strings.ToLower(w)
		} else {
			runes := []rune(w)
			for j, r := range runes {
				if j == 0 {
					runes[j] = unicode.ToUpper(r)
				} else {
					runes[j] = unicode.ToLower(r)
				}
			}
			words[i] = string(runes)
		}
	}

	result := "#" + strings.Join(words, "")
	if len(result) > 100 {
		result = result[:100]
	}
	return result
}
```

## 3591 — Check If Any Element Has Prime Frequency

```go
package main

// LeetCode #3591: Check if Any Element Has Prime Frequency
// https://leetcode.com/problems/check-if-any-element-has-prime-frequency/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CheckIfAnyElementHasPrimeFrequency([]int{1, 2, 3, 4, 5, 4}))
	fmt.Println(CheckIfAnyElementHasPrimeFrequency([]int{1, 2, 3, 4, 5}))
	fmt.Println(CheckIfAnyElementHasPrimeFrequency([]int{2, 2, 2, 4, 4}))
}

// Time: O(n + sqrt(m)) where m is max frequency
// Space: O(1)
func CheckIfAnyElementHasPrimeFrequency(nums []int) bool {
	freq := [101]int{}
	for _, v := range nums {
		freq[v]++
	}

	for _, f := range freq {
		if f > 1 && isPrime(f) {
			return true
		}
	}
	return false
}

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
```

## 3602 — Hexadecimal And Hexatrigesimal Conversion

```go
package main

// LeetCode #3602: Hexadecimal and Hexatrigesimal Conversion
// https://leetcode.com/problems/hexadecimal-and-hexatrigesimal-conversion/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

const digits = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	fmt.Println(HexadecimalAndHexatrigesimalConversion(4))
}

// Time: O(log n)
// Space: O(log n)
func HexadecimalAndHexatrigesimalConversion(n int) string {
	sq := n * n
	cube := n * n * n
	return toBase(sq, 16) + toBase(cube, 36)
}

func toBase(num, base int) string {
	if num == 0 {
		return "0"
	}
	var res strings.Builder
	for num > 0 {
		res.WriteByte(digits[num%base])
		num /= base
	}

	// reverse
	s := []byte(res.String())
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return string(s)
}
```

## 3606 — Coupon Code Validator

```go
package main

// LeetCode #3606: Coupon Code Validator
// https://leetcode.com/problems/coupon-code-validator/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(CouponCodeValidator(
		[]string{"SAVE20", "", "PHARMA5", "SAVE@20"},
		[]string{"restaurant", "grocery", "pharmacy", "restaurant"},
		[]bool{true, true, true, true},
	))
}

// Time: O(n log n)
// Space: O(n)
func CouponCodeValidator(code []string, businessLine []string, isActive []bool) []string {
	bizOrder := map[string]int{
		"electronics": 0,
		"grocery":     1,
		"pharmacy":    2,
		"restaurant":  3,
	}

	type entry struct {
		code string
		biz  string
	}
	var valid []entry

	for i, c := range code {
		if !isActive[i] {
			continue
		}
		if _, ok := bizOrder[businessLine[i]]; !ok {
			continue
		}
		if c == "" {
			continue
		}
		ok := true
		for _, ch := range c {
			if !isAlphanumeric(ch) && ch != '_' {
				ok = false
				break
			}
		}
		if ok {
			valid = append(valid, entry{c, businessLine[i]})
		}
	}

	sort.Slice(valid, func(i, j int) bool {
		if valid[i].biz != valid[j].biz {
			return bizOrder[valid[i].biz] < bizOrder[valid[j].biz]
		}
		return valid[i].code < valid[j].code
	})

	res := make([]string, len(valid))
	for i, e := range valid {
		res[i] = e.code
	}
	return res
}

func isAlphanumeric(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9')
}
```

## 3622 — Check Divisibility By Digit Sum And Product

```go
package main

// LeetCode #3622: Check Divisibility by Digit Sum and Product
// https://leetcode.com/problems/check-divisibility-by-digit-sum-and-product/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CheckDivisibilityByDigitSumAndProduct(99))
	fmt.Println(CheckDivisibilityByDigitSumAndProduct(23))
	fmt.Println(CheckDivisibilityByDigitSumAndProduct(10))
}

// Time: O(log n)
// Space: O(1)
func CheckDivisibilityByDigitSumAndProduct(n int) bool {
	x := n
	sum := 0
	prod := 1
	for x > 0 {
		d := x % 10
		sum += d
		prod *= d
		x /= 10
	}
	return n%(sum+prod) == 0
}
```

## 3633 — Earliest Finish Time For Land And Water Rides I

```go
package main

// LeetCode #3633: Earliest Finish Time for Land and Water Rides I
// https://leetcode.com/problems/earliest-finish-time-for-land-and-water-rides-i/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(EarliestFinishTimeForLandAndWaterRidesI([]int{2, 8}, []int{4, 1}, []int{6}, []int{3}))
	fmt.Println(EarliestFinishTimeForLandAndWaterRidesI([]int{5}, []int{3}, []int{1}, []int{10}))
}

// Time: O(n*m) where n = len(landStartTime), m = len(waterStartTime)
// Space: O(1)
func EarliestFinishTimeForLandAndWaterRidesI(landStartTime, landDuration, waterStartTime, waterDuration []int) int {
	ans := int(^uint(0) >> 1) // max int

	// Land -> Water
	for i := range landStartTime {
		finishLand := landStartTime[i] + landDuration[i]
		for j := range waterStartTime {
			startWater := max(finishLand, waterStartTime[j])
			finish := startWater + waterDuration[j]
			if finish < ans {
				ans = finish
			}
		}
	}

	// Water -> Land
	for i := range waterStartTime {
		finishWater := waterStartTime[i] + waterDuration[i]
		for j := range landStartTime {
			startLand := max(finishWater, landStartTime[j])
			finish := startLand + landDuration[j]
			if finish < ans {
				ans = finish
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
```

## 3637 — Trionic Array I

```go
package main

// LeetCode #3637: Trionic Array I
// https://leetcode.com/problems/trionic-array-i/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(TrionicArrayI([]int{1, 3, 5, 4, 2, 6}))
	fmt.Println(TrionicArrayI([]int{2, 1, 3}))
}

// Time: O(n)
// Space: O(1)
func TrionicArrayI(nums []int) bool {
	n := len(nums)
	if n < 4 || nums[0] >= nums[1] || nums[n-2] >= nums[n-1] {
		return false
	}

	i := 1
	// Phase 1: strictly increasing
	for i < n && nums[i] > nums[i-1] {
		i++
	}
	p := i - 1

	// Phase 2: strictly decreasing
	for i < n && nums[i] < nums[i-1] {
		i++
	}
	q := i - 1

	// Phase 3: strictly increasing
	for i < n && nums[i] > nums[i-1] {
		i++
	}

	return i == n && p > 0 && q > p && q < n-1
}
```

## 3643 — Flip Square Submatrix Vertically

```go
package main

// LeetCode #3643: Flip Square Submatrix Vertically
// https://leetcode.com/problems/flip-square-submatrix-vertically/
// Difficulty: Easy

import "fmt"

func main() {
	grid := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	fmt.Println(FlipSquareSubmatrixVertically(grid, 1, 1, 3))
}

// Time: O(k^2)
// Space: O(1)
func FlipSquareSubmatrixVertically(grid [][]int, x, y, k int) [][]int {
	for i := x; i < x+k/2; i++ {
		i2 := x + k - 1 - (i - x)
		for j := y; j < y+k; j++ {
			grid[i][j], grid[i2][j] = grid[i2][j], grid[i][j]
		}
	}
	return grid
}
```

## 3658 — Gcd Of Odd And Even Sums

```go
package main

// LeetCode #3658: GCD of Odd and Even Sums
// https://leetcode.com/problems/gcd-of-odd-and-even-sums/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(GcdOfOddAndEvenSums(4))
	fmt.Println(GcdOfOddAndEvenSums(1))
	fmt.Println(GcdOfOddAndEvenSums(10))
}

// Time: O(1)
// Space: O(1)
// The answer is always n because:
// sum of first n odd numbers = n^2
// sum of first n even numbers = n(n+1)
// gcd(n^2, n(n+1)) = n (since n and n+1 are coprime)
func GcdOfOddAndEvenSums(n int) int {
	return n
}
```

## 3662 — Filter Characters By Frequency

```go
package main

// LeetCode #3662: Filter Characters by Frequency
// https://leetcode.com/problems/filter-characters-by-frequency/
// Difficulty: Easy [Paid]

import "fmt"

func main() {
	fmt.Println(FilterCharactersByFrequency("aadbbcccca", 3))
	fmt.Println(FilterCharactersByFrequency("xyz", 2))
}

// Time: O(n)
// Space: O(1)
func FilterCharactersByFrequency(s string, k int) string {
	cnt := [26]int{}
	for _, ch := range s {
		cnt[ch-'a']++
	}

	res := make([]byte, 0, len(s))
	for _, ch := range s {
		if cnt[ch-'a'] < k {
			res = append(res, byte(ch))
		}
	}
	return string(res)
}
```

## 3663 — Find The Least Frequent Digit

```go
package main

// LeetCode #3663: Find The Least Frequent Digit
// https://leetcode.com/problems/find-the-least-frequent-digit/
// Difficulty: Easy

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(FindTheLeastFrequentDigit(1553322))
	fmt.Println(FindTheLeastFrequentDigit(723344511))
}

// Time: O(log n) - number of digits
// Space: O(1)
func FindTheLeastFrequentDigit(n int) int {
	cnt := [10]int{}
	for n > 0 {
		cnt[n%10]++
		n /= 10
	}

	minCnt := math.MaxInt
	ans := 0
	for d := 0; d <= 9; d++ {
		if cnt[d] > 0 && (cnt[d] < minCnt || (cnt[d] == minCnt && d < ans)) {
			minCnt = cnt[d]
			ans = d
		}
	}
	return ans
}
```

## 3667 — Sort Array By Absolute Value

```go
package main

// LeetCode #3667: Sort Array By Absolute Value
// https://leetcode.com/problems/sort-array-by-absolute-value/
// Difficulty: Easy [Paid]

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(SortArrayByAbsoluteValue([]int{3, -1, -4, 1, 5}))
}

// Time: O(n log n)
// Space: O(n)
func SortArrayByAbsoluteValue(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return abs(nums[i]) < abs(nums[j])
	})
	return nums
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

## 3668 — Restore Finishing Order

```go
package main

// LeetCode #3668: Restore Finishing Order
// https://leetcode.com/problems/restore-finishing-order/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(RestoreFinishingOrder([]int{3, 1, 2, 4, 5}, []int{2, 4}))
	fmt.Println(RestoreFinishingOrder([]int{1, 2, 3, 4, 5}, []int{1, 3, 5}))
}

// Time: O(n)
// Space: O(n)
func RestoreFinishingOrder(order []int, friends []int) []int {
	friendSet := make(map[int]bool)
	for _, f := range friends {
		friendSet[f] = true
	}

	res := make([]int, 0, len(friends))
	for _, id := range order {
		if friendSet[id] {
			res = append(res, id)
		}
	}
	return res
}

func init() {
	_ = sort.Ints
}
```

## 3674 — Minimum Operations To Equalize Array

```go
package main

// LeetCode #3674: Minimum Operations to Equalize Array
// https://leetcode.com/problems/minimum-operations-to-equalize-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MinimumOperationsToEqualizeArray([]int{1, 2, 3}))
	fmt.Println(MinimumOperationsToEqualizeArray([]int{5, 5, 5}))
}

// Time: O(n)
// Space: O(1)
func MinimumOperationsToEqualizeArray(nums []int) int {
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[0] {
			return 1
		}
	}
	return 0
}
```

## 3678 — Smallest Absent Positive Greater Than Average

```go
package main

// LeetCode #3678: Smallest Absent Positive Greater Than Average
// https://leetcode.com/problems/smallest-absent-positive-greater-than-average/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(SmallestAbsentPositiveGreaterThanAverage([]int{3, 5}))
	fmt.Println(SmallestAbsentPositiveGreaterThanAverage([]int{-1, 1, 2}))
	fmt.Println(SmallestAbsentPositiveGreaterThanAverage([]int{4, -1}))
}

// Time: O(n + m) where m is the range of candidate values
// Space: O(n)
func SmallestAbsentPositiveGreaterThanAverage(nums []int) int {
	has := make(map[int]bool)
	sum := 0
	for _, x := range nums {
		has[x] = true
		sum += x
	}

	avg := float64(sum) / float64(len(nums))
	ans := 1
	if int(avg)+1 > ans {
		ans = int(avg) + 1
	}

	for has[ans] {
		ans++
	}
	return ans
}
```

## 3683 — Earliest Time To Finish One Task

```go
package main

// LeetCode #3683: Earliest Time to Finish One Task
// https://leetcode.com/problems/earliest-time-to-finish-one-task/
// Difficulty: Easy

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(EarliestTimeToFinishOneTask([][]int{{1, 6}, {2, 3}}))
	fmt.Println(EarliestTimeToFinishOneTask([][]int{{100, 100}, {100, 100}, {100, 100}}))
}

// Time: O(n)
// Space: O(1)
func EarliestTimeToFinishOneTask(tasks [][]int) int {
	ans := math.MaxInt
	for _, t := range tasks {
		finish := t[0] + t[1]
		if finish < ans {
			ans = finish
		}
	}
	return ans
}
```

## 3684 — Maximize Sum Of At Most K Distinct Elements

```go
package main

// LeetCode #3684: Maximize Sum of At Most K Distinct Elements
// https://leetcode.com/problems/maximize-sum-of-at-most-k-distinct-elements/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MaximizeSumOfAtMostKDistinctElements([]int{84, 93, 100, 77, 90}, 3))
	fmt.Println(MaximizeSumOfAtMostKDistinctElements([]int{84, 93, 100, 77, 93}, 3))
	fmt.Println(MaximizeSumOfAtMostKDistinctElements([]int{1, 1, 1, 2, 2, 2}, 6))
}

// Time: O(n log n)
// Space: O(n)
func MaximizeSumOfAtMostKDistinctElements(nums []int, k int) []int {
	seen := make(map[int]bool)
	unique := make([]int, 0)
	for _, v := range nums {
		if !seen[v] {
			seen[v] = true
			unique = append(unique, v)
		}
	}

	sort.Slice(unique, func(i, j int) bool {
		return unique[i] > unique[j]
	})

	size := k
	if size > len(unique) {
		size = len(unique)
	}
	return unique[:size]
}
```

## 3687 — Library Late Fee Calculator

```go
package main

// LeetCode #3687: Library Late Fee Calculator
// https://leetcode.com/problems/library-late-fee-calculator/
// Difficulty: Easy [Paid]

import "fmt"

func main() {
	fmt.Println(LibraryLateFeeCalculator([]int{5, 1, 7}))
	fmt.Println(LibraryLateFeeCalculator([]int{1, 1}))
}

// Time: O(n)
// Space: O(1)
func LibraryLateFeeCalculator(daysLate []int) int {
	ans := 0
	for _, x := range daysLate {
		if x == 1 {
			ans += 1
		} else if x > 5 {
			ans += 3 * x
		} else {
			ans += 2 * x
		}
	}
	return ans
}
```

## 3688 — Bitwise Or Of Even Numbers In An Array

```go
package main

// LeetCode #3688: Bitwise OR of Even Numbers in an Array
// https://leetcode.com/problems/bitwise-or-of-even-numbers-in-an-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(BitwiseOrOfEvenNumbersInAnArray([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(BitwiseOrOfEvenNumbersInAnArray([]int{7, 9, 11}))
	fmt.Println(BitwiseOrOfEvenNumbersInAnArray([]int{1, 8, 16}))
}

// Time: O(n)
// Space: O(1)
func BitwiseOrOfEvenNumbersInAnArray(nums []int) int {
	res := 0
	for _, n := range nums {
		if n%2 == 0 {
			res |= n
		}
	}
	return res
}
```

