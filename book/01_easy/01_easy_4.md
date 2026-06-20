# Easy (Mudah) — Problem ��1598

## 1228 — Missing Number In Arithmetic Progression

```go
package main

// LeetCode #1228: Missing Number In Arithmetic Progression
// https://leetcode.com/problems/missing-number-in-arithmetic-progression/
// Difficulty: Easy [Paid]
// Time: O(log n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(missingNumber([]int{5, 7, 11, 13}))  // 9
	fmt.Println(missingNumber([]int{15, 13, 12}))    // 14
}

// LeetCode submission: missingNumber
func missingNumber(arr []int) int {
	n := len(arr)
	diff := (arr[n-1] - arr[0]) / n
	lo, hi := 0, n-1
	for lo < hi {
		mid := (lo + hi) >> 1
		if arr[mid] == arr[0]+mid*diff {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return arr[0] + diff*lo
}
```

## 1232 — Check If It Is A Straight Line

```go
package main

// LeetCode #1232: Check If It Is a Straight Line
// https://leetcode.com/problems/check-if-it-is-a-straight-line/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(checkStraightLine([][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}})) // true
	fmt.Println(checkStraightLine([][]int{{1, 1}, {2, 2}, {3, 4}, {4, 5}, {5, 6}, {7, 7}})) // false
}

// LeetCode submission: checkStraightLine
func checkStraightLine(coordinates [][]int) bool {
	x0, y0 := coordinates[0][0], coordinates[0][1]
	x1, y1 := coordinates[1][0], coordinates[1][1]
	dx, dy := x1-x0, y1-y0
	for i := 2; i < len(coordinates); i++ {
		xi, yi := coordinates[i][0], coordinates[i][1]
		if (xi-x0)*dy != (yi-y0)*dx {
			return false
		}
	}
	return true
}
```

## 1241 — Number Of Comments Per Post

```go
package main

// LeetCode #1241: Number of Comments per Post
// https://leetcode.com/problems/number-of-comments-per-post/
// Difficulty: Easy [Paid] (SQL)
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println("SQL: SELECT DISTINCT p.sub_id AS post_id, (SELECT COUNT(DISTINCT c.sub_id) FROM Submissions c WHERE c.parent_id = p.sub_id) AS number_of_comments FROM Submissions p WHERE p.parent_id IS NULL ORDER BY p.sub_id")
}

// This is a SQL problem. The answer is the SQL query above.
```

## 1243 — Array Transformation

```go
package main

// LeetCode #1243: Array Transformation
// https://leetcode.com/problems/array-transformation/
// Difficulty: Easy [Paid]
// Time: O(n^2) worst case | Space: O(n)

import "fmt"

func main() {
	fmt.Println(transformArray([]int{6, 2, 3, 4})) // [6,3,3,4]
	fmt.Println(transformArray([]int{1, 6, 3, 4, 3, 5})) // [1,4,4,4,4,5]
}

// LeetCode submission: transformArray
func transformArray(arr []int) []int {
	if len(arr) <= 2 {
		return append([]int{}, arr...)
	}
	for {
		changed := false
		next := make([]int, len(arr))
		copy(next, arr)
		for i := 1; i < len(arr)-1; i++ {
			if arr[i] < arr[i-1] && arr[i] < arr[i+1] {
				next[i]++
				changed = true
			} else if arr[i] > arr[i-1] && arr[i] > arr[i+1] {
				next[i]--
				changed = true
			}
		}
		if !changed {
			break
		}
		arr = next
	}
	return arr
}
```

## 1251 — Average Selling Price

```go
package main

// LeetCode #1251: Average Selling Price
// https://leetcode.com/problems/average-selling-price/
// Difficulty: Easy (SQL)
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println("SQL: SELECT p.product_id, IFNULL(ROUND(SUM(u.units * p.price) / SUM(u.units), 2), 0) AS average_price FROM Prices p LEFT JOIN UnitsSold u ON p.product_id = u.product_id AND u.purchase_date BETWEEN p.start_date AND p.end_date GROUP BY p.product_id")
}

// This is a SQL problem. The answer is the SQL query above.
```

## 1252 — Cells With Odd Values In A Matrix

```go
package main

// LeetCode #1252: Cells with Odd Values in a Matrix
// https://leetcode.com/problems/cells-with-odd-values-in-a-matrix/
// Difficulty: Easy
// Time: O(n + m + k) | Space: O(n + m)

import "fmt"

func main() {
	fmt.Println(oddCells(2, 3, [][]int{{0, 1}, {1, 1}})) // 6
	fmt.Println(oddCells(2, 2, [][]int{{1, 1}, {0, 0}})) // 0
}

// LeetCode submission: oddCells
func oddCells(m, n int, indices [][]int) int {
	rows := make([]int, m)
	cols := make([]int, n)
	for _, idx := range indices {
		rows[idx[0]]++
		cols[idx[1]]++
	}
	oddRows, oddCols := 0, 0
	for _, v := range rows {
		if v%2 == 1 {
			oddRows++
		}
	}
	for _, v := range cols {
		if v%2 == 1 {
			oddCols++
		}
	}
	return oddRows*(n-oddCols) + (m-oddRows)*oddCols
}
```

## 1260 — Shift 2d Grid

```go
package main

// LeetCode #1260: Shift 2D Grid
// https://leetcode.com/problems/shift-2d-grid/
// Difficulty: Easy
// Time: O(m*n) | Space: O(m*n)

import "fmt"

func main() {
	fmt.Println(shiftGrid([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 1))
	// [[9,1,2],[3,4,5],[6,7,8]]
	fmt.Println(shiftGrid([][]int{{3, 8, 1, 9}, {19, 7, 2, 5}, {4, 6, 11, 10}, {12, 0, 21, 13}}, 4))
	// [[12,0,21,13],[3,8,1,9],[19,7,2,5],[4,6,11,10]]
}

// LeetCode submission: shiftGrid
func shiftGrid(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			idx := (i*n + j + k) % (m * n)
			ni, nj := idx/n, idx%n
			ans[ni][nj] = grid[i][j]
		}
	}
	return ans
}
```

## 1266 — Minimum Time Visiting All Points

```go
package main

// LeetCode #1266: Minimum Time Visiting All Points
// https://leetcode.com/problems/minimum-time-visiting-all-points/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(minTimeToVisitAllPoints([][]int{{1, 1}, {3, 4}, {-1, 0}})) // 7
	fmt.Println(minTimeToVisitAllPoints([][]int{{3, 2}, {-2, 2}}))         // 5
}

// LeetCode submission: minTimeToVisitAllPoints
func minTimeToVisitAllPoints(points [][]int) int {
	ans := 0
	for i := 1; i < len(points); i++ {
		dx := points[i][0] - points[i-1][0]
		dy := points[i][1] - points[i-1][1]
		if dx < 0 {
			dx = -dx
		}
		if dy < 0 {
			dy = -dy
		}
		if dx > dy {
			ans += dx
		} else {
			ans += dy
		}
	}
	return ans
}
```

## 1271 — Hexspeak

```go
package main

// LeetCode #1271: Hexspeak
// https://leetcode.com/problems/hexspeak/
// Difficulty: Easy [Paid]
// Time: O(log n) | Space: O(log n)

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(toHexspeak("257"))  // "IOI"
	fmt.Println(toHexspeak("3"))    // "ERROR"
	fmt.Println(toHexspeak("619"))  // "ERROR" (619=26B, B not allowed)
}

// LeetCode submission: toHexspeak
func toHexspeak(num string) string {
	n, _ := strconv.Atoi(num)
	hex := strconv.FormatInt(int64(n), 16)
	replacer := map[byte]byte{
		'0': 'O',
		'1': 'I',
	}
	ans := make([]byte, len(hex))
	for i := range hex {
		if r, ok := replacer[hex[i]]; ok {
			ans[i] = r
		} else if hex[i] >= 'a' && hex[i] <= 'f' {
			ans[i] = hex[i] - 'a' + 'A'
		} else {
			return "ERROR"
		}
	}
	return string(ans)
}
```

## 1275 — Find Winner On A Tic Tac Toe Game

```go
package main

// LeetCode #1275: Find Winner on a Tic Tac Toe Game
// https://leetcode.com/problems/find-winner-on-a-tic-tac-toe-game/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(tictactoe([][]int{{0, 0}, {2, 0}, {1, 1}, {2, 1}, {2, 2}})) // "A"
	fmt.Println(tictactoe([][]int{{0, 0}, {1, 1}, {0, 1}, {0, 2}, {1, 0}, {2, 0}})) // "B"
	fmt.Println(tictactoe([][]int{{0, 0}, {1, 1}, {2, 0}, {1, 0}, {1, 2}, {2, 1}, {0, 1}, {0, 2}, {2, 2}})) // "Draw"
}

// LeetCode submission: tictactoe
func tictactoe(moves [][]int) string {
	board := make([][]byte, 3)
	for i := range board {
		board[i] = make([]byte, 3)
	}
	for i, m := range moves {
		player := byte('A')
		if i%2 == 1 {
			player = 'B'
		}
		board[m[0]][m[1]] = player
	}
	// Check rows and cols
	for i := 0; i < 3; i++ {
		if board[i][0] != 0 && board[i][0] == board[i][1] && board[i][1] == board[i][2] {
			return string(board[i][0])
		}
		if board[0][i] != 0 && board[0][i] == board[1][i] && board[1][i] == board[2][i] {
			return string(board[0][i])
		}
	}
	// Check diagonals
	if board[0][0] != 0 && board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		return string(board[0][0])
	}
	if board[0][2] != 0 && board[0][2] == board[1][1] && board[1][1] == board[2][0] {
		return string(board[0][2])
	}
	if len(moves) == 9 {
		return "Draw"
	}
	return "Pending"
}
```

## 1279 — Traffic Light Controlled Intersection

```go
package main

// LeetCode #1279: Traffic Light Controlled Intersection
// https://leetcode.com/problems/traffic-light-controlled-intersection/
// Difficulty: Easy [Paid] (Concurrency)
// Time: O(1) | Space: O(1)

import (
	"fmt"
	"sync"
)

type TrafficLight struct {
	mu      sync.Mutex
	greenOn int // 1 = road A, 2 = road B
}

func NewTrafficLight() *TrafficLight {
	return &TrafficLight{greenOn: 1}
}

func (t *TrafficLight) CarArrived(carId, roadId, direction int, turnGreen, crossCar func()) {
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.greenOn != roadId {
		t.greenOn = roadId
		turnGreen()
	}
	crossCar()
}

func main() {
	light := NewTrafficLight()
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		light.CarArrived(1, 1, 1, func() {
			fmt.Println("Turn green for Road A")
		}, func() {
			fmt.Println("Car 1 crossing on Road A")
		})
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		light.CarArrived(2, 1, 2, func() {
			fmt.Println("Turn green for Road A")
		}, func() {
			fmt.Println("Car 2 crossing on Road A")
		})
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		light.CarArrived(3, 2, 4, func() {
			fmt.Println("Turn green for Road B")
		}, func() {
			fmt.Println("Car 3 crossing on Road B")
		})
	}()

	wg.Wait()
}
```

## 1280 — Students And Examinations

```go
package main

// LeetCode #1280: Students and Examinations
// https://leetcode.com/problems/students-and-examinations/
// Difficulty: Easy (SQL)
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println("SQL: SELECT s.student_id, s.student_name, sub.subject_name, COUNT(e.subject_name) AS attended_exams FROM Students s CROSS JOIN Subjects sub LEFT JOIN Examinations e ON s.student_id = e.student_id AND sub.subject_name = e.subject_name GROUP BY s.student_id, s.student_name, sub.subject_name ORDER BY s.student_id, sub.subject_name")
}

// This is a SQL problem. The answer is the SQL query above.
```

## 1281 — Subtract The Product And Sum Of Digits Of An Integer

```go
package main

// LeetCode #1281: Subtract the Product and Sum of Digits of an Integer
// https://leetcode.com/problems/subtract-the-product-and-sum-of-digits-of-an-integer/
// Difficulty: Easy
// Time: O(log n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(subtractProductAndSum(234))  // 15
	fmt.Println(subtractProductAndSum(4421)) // 21
}

// LeetCode submission: subtractProductAndSum
func subtractProductAndSum(n int) int {
	product := 1
	sum := 0
	for x := n; x > 0; x /= 10 {
		d := x % 10
		product *= d
		sum += d
	}
	return product - sum
}
```

## 1287 — Element Appearing More Than 25 In Sorted Array

```go
package main

// LeetCode #1287: Element Appearing More Than 25% In Sorted Array
// https://leetcode.com/problems/element-appearing-more-than-25-in-sorted-array/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(findSpecialInteger([]int{1, 2, 2, 6, 6, 6, 6, 7, 10})) // 6
	fmt.Println(findSpecialInteger([]int{1, 1}))                       // 1
}

// LeetCode submission: findSpecialInteger
func findSpecialInteger(arr []int) int {
	target := len(arr) / 4
	for i := 0; i < len(arr)-target; i++ {
		if arr[i] == arr[i+target] {
			return arr[i]
		}
	}
	return arr[0]
}
```

## 1290 — Convert Binary Number In A Linked List To Integer

```go
package main

// LeetCode #1290: Convert Binary Number in a Linked List to Integer
// https://leetcode.com/problems/convert-binary-number-in-a-linked-list-to-integer/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// 1 -> 0 -> 1 = 5
	head := &ListNode{1, &ListNode{0, &ListNode{1, nil}}}
	fmt.Println(getDecimalValue(head)) // 5

	// 0 -> 0 = 0
	fmt.Println(getDecimalValue(&ListNode{0, &ListNode{0, nil}})) // 0
}

// LeetCode submission: getDecimalValue
func getDecimalValue(head *ListNode) int {
	ans := 0
	for head != nil {
		ans = ans*2 + head.Val
		head = head.Next
	}
	return ans
}
```

## 1294 — Weather Type In Each Country

```go
package main

// LeetCode #1294: Weather Type in Each Country
// https://leetcode.com/problems/weather-type-in-each-country/
// Difficulty: Easy [Paid] (SQL)
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println("SQL: SELECT c.country_name, CASE WHEN AVG(w.weather_state * 1.0) <= 15 THEN 'Cold' WHEN AVG(w.weather_state * 1.0) >= 25 THEN 'Hot' ELSE 'Warm' END AS weather_type FROM Countries c JOIN Weather w ON c.country_id = w.country_id WHERE w.day BETWEEN '2019-11-01' AND '2019-11-30' GROUP BY c.country_id, c.country_name")
}

// This is a SQL problem. The answer is the SQL query above.
```

## 1295 — Find Numbers With Even Number Of Digits

```go
package main

// LeetCode #1295: Find Numbers with Even Number of Digits
// https://leetcode.com/problems/find-numbers-with-even-number-of-digits/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(findNumbers([]int{12, 345, 2, 6, 7896})) // 2
	fmt.Println(findNumbers([]int{555, 901, 482, 1771})) // 1
}

// LeetCode submission: findNumbers
func findNumbers(nums []int) int {
	count := 0
	for _, v := range nums {
		digits := 0
		for x := v; x > 0; x /= 10 {
			digits++
		}
		if digits%2 == 0 {
			count++
		}
	}
	return count
}
```

## 1299 — Replace Elements With Greatest Element On Right Side

```go
package main

// LeetCode #1299: Replace Elements with Greatest Element on Right Side
// https://leetcode.com/problems/replace-elements-with-greatest-element-on-right-side/
// Difficulty: Easy
// Time: O(n) | Space: O(1) excluding output

import "fmt"

func main() {
	fmt.Println(replaceElements([]int{17, 18, 5, 4, 6, 1})) // [18,6,6,6,1,-1]
	fmt.Println(replaceElements([]int{400}))                 // [-1]
}

// LeetCode submission: replaceElements
func replaceElements(arr []int) []int {
	ans := make([]int, len(arr))
	maxRight := -1
	for i := len(arr) - 1; i >= 0; i-- {
		ans[i] = maxRight
		if arr[i] > maxRight {
			maxRight = arr[i]
		}
	}
	return ans
}
```

## 1303 — Find The Team Size

```go
package main

// LeetCode #1303: Find the Team Size
// https://leetcode.com/problems/find-the-team-size/
// Difficulty: Easy [Paid] (SQL)
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println("SQL: SELECT employee_id, COUNT(team_id) OVER (PARTITION BY team_id) AS team_size FROM Employee")
}

// This is a SQL problem. The answer is the SQL query above.
```

## 1304 — Find N Unique Integers Sum Up To Zero

```go
package main

// LeetCode #1304: Find N Unique Integers Sum up to Zero
// https://leetcode.com/problems/find-n-unique-integers-sum-up-to-zero/
// Difficulty: Easy
// Time: O(n) | Space: O(1) excluding output

import "fmt"

func main() {
	fmt.Println(sumZero(5)) // [0,1,2,3,-6] or similar
	fmt.Println(sumZero(3)) // [0,1,-1]
	fmt.Println(sumZero(1)) // [0]
}

// LeetCode submission: sumZero
func sumZero(n int) []int {
	ans := make([]int, n)
	if n == 1 {
		return ans // [0]
	}
	half := n / 2
	for i := 0; i < half; i++ {
		ans[i] = i + 1
		ans[i+half] = -(i + 1)
	}
	if n%2 == 1 {
		ans[n-1] = 0
	}
	return ans
}
```

## 1309 — Decrypt String From Alphabet To Integer Mapping

```go
package main

// LeetCode #1309: Decrypt String from Alphabet to Integer Mapping
// https://leetcode.com/problems/decrypt-string-from-alphabet-to-integer-mapping/
// Difficulty: Easy
//
// LeetCode submission: func freqAlphabets(s string) string

import "fmt"

func main() {
	fmt.Println(DecryptStringFromAlphabetToIntegerMapping("10#11#12"))       // "jkab"
	fmt.Println(DecryptStringFromAlphabetToIntegerMapping("1326#"))           // "acz"
	fmt.Println(DecryptStringFromAlphabetToIntegerMapping("12345678910#11#12#13#14#15#16#17#18#19#20#21#22#23#24#25#26#")) // "abcdefghijklmnopqrstuvwxyz"
}

// Time: O(n), Space: O(n)
func DecryptStringFromAlphabetToIntegerMapping(s string) string {
	res := make([]byte, 0, len(s))
	i := 0
	for i < len(s) {
		if i+2 < len(s) && s[i+2] == '#' {
			num := (s[i]-'0')*10 + (s[i+1] - '0')
			res = append(res, byte('a'+num-1))
			i += 3
		} else {
			num := s[i] - '0'
			res = append(res, byte('a'+num-1))
			i++
		}
	}
	return string(res)
}
```

## 1313 — Decompress Run Length Encoded List

```go
package main

// LeetCode #1313: Decompress Run-Length Encoded List
// https://leetcode.com/problems/decompress-run-length-encoded-list/
// Difficulty: Easy
//
// LeetCode submission: func decompressRLElist(nums []int) []int

import "fmt"

func main() {
	fmt.Println(DecompressRunLengthEncodedList([]int{1, 2, 3, 4}))       // [2 4 4 4]
	fmt.Println(DecompressRunLengthEncodedList([]int{1, 1, 2, 3}))       // [1 3 3]
	fmt.Println(DecompressRunLengthEncodedList([]int{2, 5, 1, 7, 3, 9})) // [5 5 7 9 9 9]
}

// Time: O(n + totalLen), Space: O(totalLen)
func DecompressRunLengthEncodedList(nums []int) []int {
	totalLen := 0
	for i := 0; i < len(nums); i += 2 {
		totalLen += nums[i]
	}
	res := make([]int, 0, totalLen)
	for i := 0; i < len(nums); i += 2 {
		freq, val := nums[i], nums[i+1]
		for j := 0; j < freq; j++ {
			res = append(res, val)
		}
	}
	return res
}
```

## 1317 — Convert Integer To The Sum Of Two No Zero Integers

```go
package main

// LeetCode #1317: Convert Integer to the Sum of Two No-Zero Integers
// https://leetcode.com/problems/convert-integer-to-the-sum-of-two-no-zero-integers/
// Difficulty: Easy
//
// LeetCode submission: func getNoZeroIntegers(n int) []int

import "fmt"

func main() {
	fmt.Println(ConvertIntegerToTheSumOfTwoNoZeroIntegers(2))    // [1 1]
	fmt.Println(ConvertIntegerToTheSumOfTwoNoZeroIntegers(11))   // [2 9]
	fmt.Println(ConvertIntegerToTheSumOfTwoNoZeroIntegers(1010)) // [122 888]
}

// Time: O(n log n), Space: O(1)
func ConvertIntegerToTheSumOfTwoNoZeroIntegers(n int) []int {
	for a := 1; a < n; a++ {
		b := n - a
		if !hasZero(a) && !hasZero(b) {
			return []int{a, b}
		}
	}
	return []int{}
}

func hasZero(x int) bool {
	for x > 0 {
		if x%10 == 0 {
			return true
		}
		x /= 10
	}
	return false
}
```

## 1322 — Ads Performance

```go
package main

// LeetCode #1322: Ads Performance
// https://leetcode.com/problems/ads-performance/
// Difficulty: Easy [Paid]
//
// This is a SQL problem. The solution is the SQL query below.
// Table: Ads (ad_id, user_id, action) where action is enum('Clicked','Viewed','Ignored')

import "fmt"

func main() {
	fmt.Println(AdsPerformance())
}

// Time: N/A (SQL query), Space: N/A
func AdsPerformance() string {
	return `SELECT
  ad_id,
  ROUND(IFNULL(SUM(action = 'Clicked') / SUM(action IN ('Clicked', 'Viewed')) * 100, 0), 2) AS ctr
FROM Ads
GROUP BY ad_id
ORDER BY ctr DESC, ad_id ASC;`
}
```

## 1323 — Maximum 69 Number

```go
package main

// LeetCode #1323: Maximum 69 Number
// https://leetcode.com/problems/maximum-69-number/
// Difficulty: Easy
//
// LeetCode submission: func maximum69Number(num int) int

import "fmt"

func main() {
	fmt.Println(MaximumSixNineNumber(9669)) // 9969
	fmt.Println(MaximumSixNineNumber(9996)) // 9999
	fmt.Println(MaximumSixNineNumber(9999)) // 9999
}

// Time: O(log n), Space: O(1)
func MaximumSixNineNumber(num int) int {
	maxBase := 0
	base := 1
	x := num
	for x > 0 {
		if x%10 == 6 {
			maxBase = base
		}
		x /= 10
		base *= 10
	}
	return num + maxBase*3
}
```

## 1327 — List The Products Ordered In A Period

```go
package main

// LeetCode #1327: List the Products Ordered in a Period
// https://leetcode.com/problems/list-the-products-ordered-in-a-period/
// Difficulty: Easy
//
// This is a SQL problem. The solution is the SQL query below.
// Tables: Products (product_id, product_name, product_category), Orders (product_id, order_date, unit)

import "fmt"

func main() {
	fmt.Println(ListTheProductsOrderedInAPeriod())
}

// Time: N/A (SQL query), Space: N/A
func ListTheProductsOrderedInAPeriod() string {
	return `SELECT p.product_name, SUM(o.unit) AS unit
FROM Products p
JOIN Orders o ON p.product_id = o.product_id
WHERE o.order_date BETWEEN '2020-02-01' AND '2020-02-29'
GROUP BY p.product_id, p.product_name
HAVING SUM(o.unit) >= 100;`
}
```

## 1331 — Rank Transform Of An Array

```go
package main

// LeetCode #1331: Rank Transform of an Array
// https://leetcode.com/problems/rank-transform-of-an-array/
// Difficulty: Easy
//
// LeetCode submission: func arrayRankTransform(arr []int) []int

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(RankTransformOfAnArray([]int{40, 10, 20, 30}))       // [4 1 2 3]
	fmt.Println(RankTransformOfAnArray([]int{100, 100, 100}))        // [1 1 1]
	fmt.Println(RankTransformOfAnArray([]int{37, 12, 28, 9, 100, 56})) // [5 3 4 1 6 2]
}

// Time: O(n log n), Space: O(n)
func RankTransformOfAnArray(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	sorted := make([]int, len(arr))
	copy(sorted, arr)
	sort.Ints(sorted)

	rank := make(map[int]int, len(arr))
	cur := 1
	for _, v := range sorted {
		if _, seen := rank[v]; !seen {
			rank[v] = cur
			cur++
		}
	}

	res := make([]int, len(arr))
	for i, v := range arr {
		res[i] = rank[v]
	}
	return res
}
```

## 1332 — Remove Palindromic Subsequences

```go
package main

// LeetCode #1332: Remove Palindromic Subsequences
// https://leetcode.com/problems/remove-palindromic-subsequences/
// Difficulty: Easy
//
// LeetCode submission: func removePalindromeSub(s string) int

import "fmt"

func main() {
	fmt.Println(RemovePalindromicSubsequences("ababa"))  // 1 (already palindrome)
	fmt.Println(RemovePalindromicSubsequences("abb"))    // 2
	fmt.Println(RemovePalindromicSubsequences("baabb"))  // 2
}

// Time: O(n), Space: O(1)
func RemovePalindromicSubsequences(s string) int {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return 2
		}
		i++
		j--
	}
	return 1
}
```

## 1337 — The K Weakest Rows In A Matrix

```go
package main

// LeetCode #1337: The K Weakest Rows in a Matrix
// https://leetcode.com/problems/the-k-weakest-rows-in-a-matrix/
// Difficulty: Easy
//
// LeetCode submission: func kWeakestRows(mat [][]int, k int) []int

import (
	"fmt"
	"sort"
)

func main() {
	mat1 := [][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 1},
	}
	fmt.Println(TheKWeakestRowsInAMatrix(mat1, 3)) // [2 0 3]

	mat2 := [][]int{
		{1, 0, 0, 0},
		{1, 1, 1, 1},
		{1, 0, 0, 0},
		{1, 1, 0, 0},
	}
	fmt.Println(TheKWeakestRowsInAMatrix(mat2, 2)) // [0 2]
}

// Time: O(m * n + m log m), Space: O(m)
func TheKWeakestRowsInAMatrix(mat [][]int, k int) []int {
	strength := make([][2]int, len(mat))
	for i, row := range mat {
		s := 0
		for _, v := range row {
			if v == 0 {
				break
			}
			s++
		}
		strength[i] = [2]int{s, i}
	}

	sort.Slice(strength, func(i, j int) bool {
		if strength[i][0] != strength[j][0] {
			return strength[i][0] < strength[j][0]
		}
		return strength[i][1] < strength[j][1]
	})

	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = strength[i][1]
	}
	return res
}
```

## 1342 — Number Of Steps To Reduce A Number To Zero

```go
package main

// LeetCode #1342: Number of Steps to Reduce a Number to Zero
// https://leetcode.com/problems/number-of-steps-to-reduce-a-number-to-zero/
// Difficulty: Easy
//
// LeetCode submission: func numberOfSteps(num int) int

import "fmt"

func main() {
	fmt.Println(NumberOfStepsToReduceANumberToZero(14)) // 6
	fmt.Println(NumberOfStepsToReduceANumberToZero(8))  // 4
	fmt.Println(NumberOfStepsToReduceANumberToZero(0))  // 0
}

// Time: O(log n), Space: O(1)
func NumberOfStepsToReduceANumberToZero(num int) int {
	steps := 0
	for num > 0 {
		if num%2 == 0 {
			num /= 2
		} else {
			num--
		}
		steps++
	}
	return steps
}
```

## 1346 — Check If N And Its Double Exist

```go
package main

// LeetCode #1346: Check If N and Its Double Exist
// https://leetcode.com/problems/check-if-n-and-its-double-exist/
// Difficulty: Easy
//
// LeetCode submission: func checkIfExist(arr []int) bool

import "fmt"

func main() {
	fmt.Println(CheckIfNAndItsDoubleExist([]int{10, 2, 5, 3}))  // true (10 = 2*5)
	fmt.Println(CheckIfNAndItsDoubleExist([]int{3, 1, 7, 11}))   // false
	fmt.Println(CheckIfNAndItsDoubleExist([]int{7, 1, 14, 11}))  // true (14 = 2*7)
}

// Time: O(n), Space: O(n)
func CheckIfNAndItsDoubleExist(arr []int) bool {
	seen := make(map[int]bool, len(arr))
	for _, v := range arr {
		if seen[v*2] || (v%2 == 0 && seen[v/2]) {
			return true
		}
		seen[v] = true
	}
	return false
}
```

## 1350 — Students With Invalid Departments

```go
package main

// LeetCode #1350: Students With Invalid Departments
// https://leetcode.com/problems/students-with-invalid-departments/
// Difficulty: Easy [Paid]
//
// This is a SQL problem. The solution is the SQL query below.
// Tables: Students (id, name, department_id), Departments (id, name)

import "fmt"

func main() {
	fmt.Println(StudentsWithInvalidDepartments())
}

// Time: N/A (SQL query), Space: N/A
func StudentsWithInvalidDepartments() string {
	return `SELECT s.id, s.name
FROM Students s
LEFT JOIN Departments d ON s.department_id = d.id
WHERE d.id IS NULL;`
}
```

## 1351 — Count Negative Numbers In A Sorted Matrix

```go
package main

// LeetCode #1351: Count Negative Numbers in a Sorted Matrix
// https://leetcode.com/problems/count-negative-numbers-in-a-sorted-matrix/
// Difficulty: Easy
//
// LeetCode submission: func countNegatives(grid [][]int) int

import "fmt"

func main() {
	grid1 := [][]int{
		{4, 3, 2, -1},
		{3, 2, 1, -1},
		{1, 1, -1, -2},
		{-1, -1, -2, -3},
	}
	fmt.Println(CountNegativeNumbersInASortedMatrix(grid1)) // 8

	grid2 := [][]int{
		{3, 2},
		{1, 0},
	}
	fmt.Println(CountNegativeNumbersInASortedMatrix(grid2)) // 0
}

// Time: O(m + n), Space: O(1)
func CountNegativeNumbersInASortedMatrix(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	count := 0
	row, col := 0, n-1
	for row < m && col >= 0 {
		if grid[row][col] < 0 {
			count += m - row
			col--
		} else {
			row++
		}
	}
	return count
}
```

## 1356 — Sort Integers By The Number Of 1 Bits

```go
package main

// LeetCode #1356: Sort Integers by The Number of 1 Bits
// https://leetcode.com/problems/sort-integers-by-the-number-of-1-bits/
// Difficulty: Easy
//
// LeetCode submission: func sortByBits(arr []int) []int

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(SortIntegersByTheNumberOfOneBits([]int{0, 1, 2, 3, 4, 5, 6, 7, 8})) // [0 1 2 4 8 3 5 6 7]
	fmt.Println(SortIntegersByTheNumberOfOneBits([]int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1})) // [1 2 4 8 16 32 64 128 256 512 1024]
}

// Time: O(n log n), Space: O(1)
func SortIntegersByTheNumberOfOneBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		bi, bj := popcount(arr[i]), popcount(arr[j])
		if bi != bj {
			return bi < bj
		}
		return arr[i] < arr[j]
	})
	return arr
}

func popcount(x int) int {
	count := 0
	for x > 0 {
		count += x & 1
		x >>= 1
	}
	return count
}
```

## 1360 — Number Of Days Between Two Dates

```go
package main

// LeetCode #1360: Number of Days Between Two Dates
// https://leetcode.com/problems/number-of-days-between-two-dates/
// Difficulty: Easy
//
// LeetCode submission: func daysBetweenDates(date1 string, date2 string) int

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(NumberOfDaysBetweenTwoDates("2019-06-29", "2019-06-30")) // 1
	fmt.Println(NumberOfDaysBetweenTwoDates("2020-01-15", "2019-12-31")) // 15
}

// Time: O(1), Space: O(1)
func NumberOfDaysBetweenTwoDates(date1 string, date2 string) int {
	format := "2006-01-02"
	d1, _ := time.Parse(format, date1)
	d2, _ := time.Parse(format, date2)
	diff := d2.Sub(d1)
	if diff < 0 {
		diff = -diff
	}
	return int(diff.Hours() / 24)
}
```

## 1365 — How Many Numbers Are Smaller Than The Current Number

```go
package main

// LeetCode #1365: How Many Numbers Are Smaller Than the Current Number
// https://leetcode.com/problems/how-many-numbers-are-smaller-than-the-current-number/
// Difficulty: Easy
//
// LeetCode submission: func smallerNumbersThanCurrent(nums []int) []int

import "fmt"

func main() {
	fmt.Println(HowManyNumbersAreSmallerThanTheCurrentNumber([]int{8, 1, 2, 2, 3})) // [4 0 1 1 3]
	fmt.Println(HowManyNumbersAreSmallerThanTheCurrentNumber([]int{6, 5, 4, 8}))    // [2 1 0 3]
	fmt.Println(HowManyNumbersAreSmallerThanTheCurrentNumber([]int{7, 7, 7, 7}))    // [0 0 0 0]
}

// Time: O(n), Space: O(1) — since count array is fixed size 101
func HowManyNumbersAreSmallerThanTheCurrentNumber(nums []int) []int {
	count := make([]int, 101)
	for _, v := range nums {
		count[v]++
	}
	for i := 1; i < 101; i++ {
		count[i] += count[i-1]
	}
	res := make([]int, len(nums))
	for i, v := range nums {
		if v > 0 {
			res[i] = count[v-1]
		}
	}
	return res
}
```

## 1370 — Increasing Decreasing String

```go
package main

// LeetCode #1370: Increasing Decreasing String
// https://leetcode.com/problems/increasing-decreasing-string/
// Difficulty: Easy
//
// LeetCode submission: func sortString(s string) string

import "fmt"

func main() {
	fmt.Println(IncreasingDecreasingString("aaaabbbbcccc")) // "abccbaabccba"
	fmt.Println(IncreasingDecreasingString("rat"))          // "art"
	fmt.Println(IncreasingDecreasingString("leetcode"))     // "cdelotee"
}

// Time: O(n), Space: O(n)
func IncreasingDecreasingString(s string) string {
	freq := make([]int, 26)
	for _, ch := range s {
		freq[ch-'a']++
	}

	res := make([]byte, 0, len(s))
	for len(res) < len(s) {
		for i := 0; i < 26; i++ {
			if freq[i] > 0 {
				res = append(res, byte('a'+i))
				freq[i]--
			}
		}
		for i := 25; i >= 0; i-- {
			if freq[i] > 0 {
				res = append(res, byte('a'+i))
				freq[i]--
			}
		}
	}
	return string(res)
}
```

## 1374 — Generate A String With Characters That Have Odd Counts

```go
package main

// LeetCode #1374: Generate a String With Characters That Have Odd Counts
// https://leetcode.com/problems/generate-a-string-with-characters-that-have-odd-counts/
// Difficulty: Easy
//
// LeetCode submission: func generateTheString(n int) string

import "fmt"

func main() {
	fmt.Println(GenerateAStringWithCharactersThatHaveOddCounts(4)) // "aaab"
	fmt.Println(GenerateAStringWithCharactersThatHaveOddCounts(2)) // "ab"
	fmt.Println(GenerateAStringWithCharactersThatHaveOddCounts(7)) // "aaaaaaa"
}

// Time: O(n), Space: O(n)
func GenerateAStringWithCharactersThatHaveOddCounts(n int) string {
	if n%2 == 1 {
		return string(makeN('a', n))
	}
	return string(makeN('a', n-1)) + "b"
}

func makeN(ch byte, n int) []byte {
	res := make([]byte, n)
	for i := range res {
		res[i] = ch
	}
	return res
}
```

## 1378 — Replace Employee Id With The Unique Identifier

```go
package main

// LeetCode #1378: Replace Employee ID With The Unique Identifier
// https://leetcode.com/problems/replace-employee-id-with-the-unique-identifier/
// Difficulty: Easy
//
// This is a SQL problem. The solution is the SQL query below.
// Tables: Employees (id, name), EmployeeUNI (id, unique_id)

import "fmt"

func main() {
	fmt.Println(ReplaceEmployeeIdWithTheUniqueIdentifier())
}

// Time: N/A (SQL query), Space: N/A
func ReplaceEmployeeIdWithTheUniqueIdentifier() string {
	return `SELECT eu.unique_id, e.name
FROM Employees e
LEFT JOIN EmployeeUNI eu ON e.id = eu.id;`
}
```

## 1379 — Find A Corresponding Node Of A Binary Tree In A Clone Of That Tree

```go
package main

// LeetCode #1379: Find a Corresponding Node of a Binary Tree in a Clone of That Tree
// https://leetcode.com/problems/find-a-corresponding-node-of-a-binary-tree-in-a-clone-of-that-tree/
// Difficulty: Easy
//
// LeetCode submission: func getTargetCopy(original, cloned *TreeNode, target *TreeNode) *TreeNode

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Original: [7,4,3,null,null,6,19]
	original := &TreeNode{Val: 7,
		Left:  &TreeNode{Val: 4},
		Right: &TreeNode{Val: 3,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 19},
		},
	}
	cloned := &TreeNode{Val: 7,
		Left:  &TreeNode{Val: 4},
		Right: &TreeNode{Val: 3,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 19},
		},
	}
	target := original.Right // node with value 3
	result := FindACorrespondingNodeOfABinaryTreeInACloneOfThatTree(original, cloned, target)
	fmt.Println(result.Val) // 3
}

// Time: O(n), Space: O(h) where h is tree height
func FindACorrespondingNodeOfABinaryTreeInACloneOfThatTree(original, cloned *TreeNode, target *TreeNode) *TreeNode {
	if original == nil {
		return nil
	}
	if original == target {
		return cloned
	}
	left := FindACorrespondingNodeOfABinaryTreeInACloneOfThatTree(original.Left, cloned.Left, target)
	if left != nil {
		return left
	}
	return FindACorrespondingNodeOfABinaryTreeInACloneOfThatTree(original.Right, cloned.Right, target)
}
```

## 1380 — Lucky Numbers In A Matrix

```go
package main

// LeetCode #1380: Lucky Numbers in a Matrix
// https://leetcode.com/problems/lucky-numbers-in-a-matrix/
// Difficulty: Easy
//
// LeetCode submission: func luckyNumbers(matrix [][]int) []int

import "fmt"

func main() {
	mat1 := [][]int{
		{3, 7, 8},
		{9, 11, 13},
		{15, 16, 17},
	}
	fmt.Println(LuckyNumbersInAMatrix(mat1)) // [15]

	mat2 := [][]int{
		{1, 10, 4, 2},
		{9, 3, 8, 7},
		{15, 16, 17, 12},
	}
	fmt.Println(LuckyNumbersInAMatrix(mat2)) // [12]
}

// Time: O(m * n), Space: O(m + n)
func LuckyNumbersInAMatrix(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	rowMin := make([]int, m)
	for i := range rowMin {
		rowMin[i] = 1<<31 - 1
	}
	colMax := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			v := matrix[i][j]
			if v < rowMin[i] {
				rowMin[i] = v
			}
			if v > colMax[j] {
				colMax[j] = v
			}
		}
	}
	res := make([]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == rowMin[i] && matrix[i][j] == colMax[j] {
				res = append(res, matrix[i][j])
			}
		}
	}
	return res
}
```

## 1385 — Find The Distance Value Between Two Arrays

```go
package main

// LeetCode #1385: Find the Distance Value Between Two Arrays
// https://leetcode.com/problems/find-the-distance-value-between-two-arrays/
// Difficulty: Easy
//
// LeetCode submission: func findTheDistanceValue(arr1 []int, arr2 []int, d int) int

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(FindTheDistanceValueBetweenTwoArrays([]int{4, 5, 8}, []int{10, 9, 1, 8}, 2)) // 2
	fmt.Println(FindTheDistanceValueBetweenTwoArrays([]int{1, 4, 2, 3}, []int{-4, -3, 6, 10, 20, 30}, 3)) // 2
}

// Time: O(n log m + m log m), Space: O(1)
func FindTheDistanceValueBetweenTwoArrays(arr1 []int, arr2 []int, d int) int {
	sort.Ints(arr2)
	count := 0
	for _, v := range arr1 {
		if isFar(v, arr2, d) {
			count++
		}
	}
	return count
}

func isFar(v int, arr []int, d int) bool {
	idx := sort.SearchInts(arr, v)
	if idx < len(arr) && abs(arr[idx]-v) <= d {
		return false
	}
	if idx > 0 && abs(arr[idx-1]-v) <= d {
		return false
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

## 1389 — Create Target Array In The Given Order

```go
package main

// LeetCode #1389: Create Target Array in the Given Order
// https://leetcode.com/problems/create-target-array-in-the-given-order/
// Difficulty: Easy
//
// LeetCode submission: func createTargetArray(nums []int, index []int) []int

import "fmt"

func main() {
	fmt.Println(CreateTargetArrayInTheGivenOrder([]int{0, 1, 2, 3, 4}, []int{0, 1, 2, 2, 1})) // [0 4 1 3 2]
	fmt.Println(CreateTargetArrayInTheGivenOrder([]int{1, 2, 3, 4, 0}, []int{0, 1, 2, 3, 0})) // [0 1 2 3 4]
}

// Time: O(n^2), Space: O(n)
func CreateTargetArrayInTheGivenOrder(nums []int, index []int) []int {
	res := make([]int, 0, len(nums))
	for i, idx := range index {
		res = append(res[:idx], append([]int{nums[i]}, res[idx:]...)...)
	}
	return res
}
```

## 1394 — Find Lucky Integer In An Array

```go
package main

// LeetCode #1394: Find Lucky Integer in an Array
// https://leetcode.com/problems/find-lucky-integer-in-an-array/
// Difficulty: Easy
//
// LeetCode submission: func findLucky(arr []int) int

import "fmt"

func main() {
	fmt.Println(FindLuckyIntegerInAnArray([]int{2, 2, 3, 4}))       // 2
	fmt.Println(FindLuckyIntegerInAnArray([]int{1, 2, 2, 3, 3, 3})) // 3
	fmt.Println(FindLuckyIntegerInAnArray([]int{2, 2, 2, 3, 3}))    // -1
}

// Time: O(n), Space: O(n)
func FindLuckyIntegerInAnArray(arr []int) int {
	freq := make(map[int]int, len(arr))
	for _, v := range arr {
		freq[v]++
	}
	ans := -1
	for k, v := range freq {
		if k == v && k > ans {
			ans = k
		}
	}
	return ans
}
```

## 1399 — Count Largest Group

```go
package main

// LeetCode #1399: Count Largest Group
// https://leetcode.com/problems/count-largest-group/
// Difficulty: Easy
//
// LeetCode submission: func countLargestGroup(n int) int

import "fmt"

func main() {
	fmt.Println(CountLargestGroup(13)) // 4
	fmt.Println(CountLargestGroup(2))  // 2
	fmt.Println(CountLargestGroup(15)) // 6
}

// Time: O(n log n), Space: O(n)
func CountLargestGroup(n int) int {
	groups := make(map[int]int)
	maxSize := 0
	for i := 1; i <= n; i++ {
		s := digitSum(i)
		groups[s]++
		if groups[s] > maxSize {
			maxSize = groups[s]
		}
	}
	count := 0
	for _, v := range groups {
		if v == maxSize {
			count++
		}
	}
	return count
}

func digitSum(x int) int {
	sum := 0
	for x > 0 {
		sum += x % 10
		x /= 10
	}
	return sum
}
```

## 1403 — Minimum Subsequence In Non Increasing Order

```go
package main

// LeetCode #1403: Minimum Subsequence in Non-Increasing Order
// https://leetcode.com/problems/minimum-subsequence-in-non-increasing-order/
// Difficulty: Easy
//
// LeetCode submission: func minSubsequence(nums []int) []int

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MinimumSubsequenceInNonIncreasingOrder([]int{4, 3, 10, 9, 8})) // [10 9]
	fmt.Println(MinimumSubsequenceInNonIncreasingOrder([]int{4, 4, 7, 6, 7}))  // [7 7 6]
}

// Time: O(n log n), Space: O(1) excluding output
func MinimumSubsequenceInNonIncreasingOrder(nums []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	total := 0
	for _, v := range nums {
		total += v
	}
	sum := 0
	for i, v := range nums {
		sum += v
		if sum > total-sum {
			return nums[:i+1]
		}
	}
	return nums
}
```

## 1407 — Top Travellers

```go
package main

// LeetCode #1407: Top Travellers
// https://leetcode.com/problems/top-travellers/
// Difficulty: Easy
//
// This is a SQL problem. The solution is the SQL query below.
// Tables: Users (id, name), Rides (id, user_id, distance)

import "fmt"

func main() {
	fmt.Println(TopTravellers())
}

// Time: N/A (SQL query), Space: N/A
func TopTravellers() string {
	return `SELECT u.name, IFNULL(SUM(r.distance), 0) AS travelled_distance
FROM Users u
LEFT JOIN Rides r ON u.id = r.user_id
GROUP BY u.id, u.name
ORDER BY travelled_distance DESC, u.name ASC;`
}
```

## 1408 — String Matching In An Array

```go
package main

// LeetCode #1408: String Matching in an Array
// https://leetcode.com/problems/string-matching-in-an-array/
// Difficulty: Easy
//
// LeetCode submission: func stringMatching(words []string) []string

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(StringMatchingInAnArray([]string{"mass", "as", "hero", "superhero"})) // [as hero]
	fmt.Println(StringMatchingInAnArray([]string{"leetcode", "et", "code"}))          // [et code]
}

// Time: O(n^2 * L) where L is average word length, Space: O(n)
func StringMatchingInAnArray(words []string) []string {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	res := make([]string, 0)
	for i, w := range words {
		for j := i + 1; j < len(words); j++ {
			if strings.Contains(words[j], w) {
				res = append(res, w)
				break
			}
		}
	}
	return res
}
```

## 1413 — Minimum Value To Get Positive Step By Step Sum

```go
package main

// LeetCode #1413: Minimum Value to Get Positive Step by Step Sum
// https://leetcode.com/problems/minimum-value-to-get-positive-step-by-step-sum/
// Difficulty: Easy
//
// LeetCode submission: func minStartValue(nums []int) int

import "fmt"

func main() {
	fmt.Println(MinimumValueToGetPositiveStepByStepSum([]int{-3, 2, -3, 4, 2})) // 5
	fmt.Println(MinimumValueToGetPositiveStepByStepSum([]int{1, 2}))             // 1
	fmt.Println(MinimumValueToGetPositiveStepByStepSum([]int{1, -2, -3}))        // 5
}

// Time: O(n), Space: O(1)
func MinimumValueToGetPositiveStepByStepSum(nums []int) int {
	minSum, sum := 0, 0
	for _, v := range nums {
		sum += v
		if sum < minSum {
			minSum = sum
		}
	}
	return -minSum + 1
}
```

## 1417 — Reformat The String

```go
package main

// LeetCode #1417: Reformat The String
// https://leetcode.com/problems/reformat-the-string/
// Difficulty: Easy
//
// LeetCode submission: func reformat(s string) string

import "fmt"

func main() {
	fmt.Println(ReformatTheString("a0b1c2")) // "a0b1c2"
	fmt.Println(ReformatTheString("leetcode")) // ""
	fmt.Println(ReformatTheString("1229857369")) // ""
}

// Time: O(n), Space: O(n)
func ReformatTheString(s string) string {
	letters := make([]byte, 0, len(s))
	digits := make([]byte, 0, len(s))
	for i := range s {
		if s[i] >= 'a' && s[i] <= 'z' {
			letters = append(letters, s[i])
		} else {
			digits = append(digits, s[i])
		}
	}
	if abs(len(letters)-len(digits)) > 1 {
		return ""
	}
	res := make([]byte, len(s))
	var first, second []byte
	if len(letters) >= len(digits) {
		first, second = letters, digits
	} else {
		first, second = digits, letters
	}
	idx := 0
	for i := 0; i < len(first); i++ {
		res[idx] = first[i]
		idx++
		if i < len(second) {
			res[idx] = second[i]
			idx++
		}
	}
	return string(res)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

## 1421 — Npv Queries

```go
package main

// LeetCode #1421: NPV Queries
// https://leetcode.com/problems/npv-queries/
// Difficulty: Easy [Paid]
//
// This is a SQL problem. The solution is the SQL query below.
// Tables: NPV (id, year, npv), Queries (id, year)

import "fmt"

func main() {
	fmt.Println(NpvQueries())
}

// Time: N/A (SQL query), Space: N/A
func NpvQueries() string {
	return `SELECT q.id, q.year, IFNULL(n.npv, 0) AS npv
FROM Queries q
LEFT JOIN NPV n ON q.id = n.id AND q.year = n.year
ORDER BY q.id, q.year;`
}
```

## 1422 — Maximum Score After Splitting A String

```go
package main

// LeetCode #1422: Maximum Score After Splitting a String
// https://leetcode.com/problems/maximum-score-after-splitting-a-string/
// Difficulty: Easy
//
// LeetCode submission: func maxScore(s string) int

import "fmt"

func main() {
	fmt.Println(MaximumScoreAfterSplittingAString("011101")) // 5
	fmt.Println(MaximumScoreAfterSplittingAString("00111"))  // 5
	fmt.Println(MaximumScoreAfterSplittingAString("1111"))   // 3
}

// Time: O(n), Space: O(1)
func MaximumScoreAfterSplittingAString(s string) int {
	ones := 0
	for _, ch := range s {
		if ch == '1' {
			ones++
		}
	}
	zeros, maxScore := 0, 0
	for i := 0; i < len(s)-1; i++ {
		if s[i] == '0' {
			zeros++
		} else {
			ones--
		}
		if zeros+ones > maxScore {
			maxScore = zeros + ones
		}
	}
	return maxScore
}
```

## 1426 — Counting Elements

```go
package main

// LeetCode #1426: Counting Elements
// https://leetcode.com/problems/counting-elements/
// Difficulty: Easy [Paid]
//
// LeetCode submission: func countElements(arr []int) int

import "fmt"

func main() {
	fmt.Println(CountingElements([]int{1, 2, 3}))       // 2
	fmt.Println(CountingElements([]int{1, 1, 3, 3, 5, 5, 7, 7})) // 0
	fmt.Println(CountingElements([]int{1, 1, 2, 2}))    // 2
}

// Time: O(n), Space: O(n)
func CountingElements(arr []int) int {
	seen := make(map[int]bool, len(arr))
	for _, v := range arr {
		seen[v] = true
	}
	count := 0
	for _, v := range arr {
		if seen[v+1] {
			count++
		}
	}
	return count
}
```

## 1427 — Perform String Shifts

```go
package main

// LeetCode #1427: Perform String Shifts
// https://leetcode.com/problems/perform-string-shifts/
// Difficulty: Easy [Paid]
//
// LeetCode submission: func stringShift(s string, shift [][]int) string

import "fmt"

func main() {
	fmt.Println(PerformStringShifts("abc", [][]int{{0, 1}, {1, 2}}))             // "cab"
	fmt.Println(PerformStringShifts("abcdefg", [][]int{{1, 1}, {1, 1}, {0, 2}, {1, 3}})) // "efgabcd"
}

// Time: O(n + m), Space: O(n)
func PerformStringShifts(s string, shift [][]int) string {
	total := 0
	for _, sh := range shift {
		if sh[0] == 0 {
			total -= sh[1]
		} else {
			total += sh[1]
		}
	}
	n := len(s)
	total %= n
	if total < 0 {
		total += n
	}
	return s[n-total:] + s[:n-total]
}
```

## 1431 — Kids With The Greatest Number Of Candies

```go
package main

// LeetCode #1431: Kids With the Greatest Number of Candies
// https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/
// Difficulty: Easy
//
// LeetCode submission: func kidsWithCandies(candies []int, extraCandies int) []bool

import "fmt"

func main() {
	fmt.Println(KidsWithTheGreatestNumberOfCandies([]int{2, 3, 5, 1, 3}, 3)) // [true true true false true]
	fmt.Println(KidsWithTheGreatestNumberOfCandies([]int{4, 2, 1, 1, 2}, 1)) // [true false false false false]
}

// Time: O(n), Space: O(n)
func KidsWithTheGreatestNumberOfCandies(candies []int, extraCandies int) []bool {
	maxCandy := 0
	for _, c := range candies {
		if c > maxCandy {
			maxCandy = c
		}
	}
	res := make([]bool, len(candies))
	for i, c := range candies {
		res[i] = c+extraCandies >= maxCandy
	}
	return res
}
```

## 1435 — Create A Session Bar Chart

```go
package main

// LeetCode #1435: Create a Session Bar Chart
// https://leetcode.com/problems/create-a-session-bar-chart/
// Difficulty: Easy [Paid]
//
// This is a SQL problem. The solution is the SQL query below.
// Table: Sessions (session_id, duration)

import "fmt"

func main() {
	fmt.Println(CreateASessionBarChart())
}

// Time: N/A (SQL query), Space: N/A
func CreateASessionBarChart() string {
	return `SELECT
  CASE
    WHEN duration / 60 BETWEEN 0 AND 4 THEN '[0-5>'
    WHEN duration / 60 BETWEEN 5 AND 9 THEN '[5-10>'
    WHEN duration / 60 BETWEEN 10 AND 14 THEN '[10-15>'
    ELSE '15 or more'
  END AS bin,
  COUNT(*) AS total
FROM Sessions
GROUP BY bin;`
}
```

## 1436 — Destination City

```go
package main

// LeetCode #1436: Destination City
// https://leetcode.com/problems/destination-city/
// Difficulty: Easy
//
// LeetCode submission: func destCity(paths [][]string) string

import "fmt"

func main() {
	paths1 := [][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}}
	fmt.Println(DestinationCity(paths1)) // "Sao Paulo"

	paths2 := [][]string{{"B", "C"}, {"D", "B"}, {"C", "A"}}
	fmt.Println(DestinationCity(paths2)) // "A"
}

// Time: O(n), Space: O(n)
func DestinationCity(paths [][]string) string {
	outgoing := make(map[string]bool, len(paths))
	for _, p := range paths {
		outgoing[p[0]] = true
	}
	for _, p := range paths {
		if !outgoing[p[1]] {
			return p[1]
		}
	}
	return ""
}
```

## 1437 — Check If All 1s Are At Least Length K Places Away

```go
package main

// LeetCode #1437: Check If All 1's Are at Least Length K Places Away
// https://leetcode.com/problems/check-if-all-1s-are-at-least-length-k-places-away/
// Difficulty: Easy
//
// LeetCode submission: func kLengthApart(nums []int, k int) bool

import "fmt"

func main() {
	fmt.Println(CheckIfAllOneSAreAtLeastLengthKPlacesAway([]int{1, 0, 0, 0, 1, 0, 0, 1}, 2)) // true
	fmt.Println(CheckIfAllOneSAreAtLeastLengthKPlacesAway([]int{1, 0, 0, 1, 0, 1}, 2))       // false
}

// Time: O(n), Space: O(1)
func CheckIfAllOneSAreAtLeastLengthKPlacesAway(nums []int, k int) bool {
	prev := -k - 1
	for i, v := range nums {
		if v == 1 {
			if i-prev-1 < k {
				return false
			}
			prev = i
		}
	}
	return true
}
```

## 1446 — Consecutive Characters

```go
package main

// LeetCode #1446: Consecutive Characters
// https://leetcode.com/problems/consecutive-characters/
// Difficulty: Easy
//
// LeetCode submission: func maxPower(s string) int

import "fmt"

func main() {
	fmt.Println(ConsecutiveCharacters("leetcode")) // 2
	fmt.Println(ConsecutiveCharacters("abbcccddddeeeeedcba")) // 5
	fmt.Println(ConsecutiveCharacters("triplepillooooow")) // 5
}

// Time: O(n), Space: O(1)
func ConsecutiveCharacters(s string) int {
	ans, cur := 1, 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			cur++
			if cur > ans {
				ans = cur
			}
		} else {
			cur = 1
		}
	}
	return ans
}
```

## 1450 — Number Of Students Doing Homework At A Given Time

```go
package main

// LeetCode #1450: Number of Students Doing Homework at a Given Time
// https://leetcode.com/problems/number-of-students-doing-homework-at-a-given-time/
// Difficulty: Easy
//
// LeetCode submission: func busyStudent(startTime []int, endTime []int, queryTime int) int

import "fmt"

func main() {
	fmt.Println(NumberOfStudentsDoingHomeworkAtAGivenTime([]int{1, 2, 3}, []int{3, 2, 7}, 4)) // 1
	fmt.Println(NumberOfStudentsDoingHomeworkAtAGivenTime([]int{4}, []int{4}, 4))             // 1
}

// Time: O(n), Space: O(1)
func NumberOfStudentsDoingHomeworkAtAGivenTime(startTime []int, endTime []int, queryTime int) int {
	count := 0
	for i := range startTime {
		if startTime[i] <= queryTime && queryTime <= endTime[i] {
			count++
		}
	}
	return count
}
```

## 1455 — Check If A Word Occurs As A Prefix Of Any Word In A Sentence

```go
package main

// LeetCode #1455: Check If a Word Occurs As a Prefix of Any Word in a Sentence
// https://leetcode.com/problems/check-if-a-word-occurs-as-a-prefix-of-any-word-in-a-sentence/
// Difficulty: Easy
//
// LeetCode submission: func isPrefixOfWord(sentence string, searchWord string) int

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(CheckIfAWordOccursAsAPrefixOfAnyWordInASentence("i love eating burger", "burg")) // 4
	fmt.Println(CheckIfAWordOccursAsAPrefixOfAnyWordInASentence("this problem is an easy problem", "pro")) // 2
	fmt.Println(CheckIfAWordOccursAsAPrefixOfAnyWordInASentence("i am tired", "you")) // -1
}

// Time: O(n), Space: O(n)
func CheckIfAWordOccursAsAPrefixOfAnyWordInASentence(sentence string, searchWord string) int {
	words := strings.Split(sentence, " ")
	for i, w := range words {
		if strings.HasPrefix(w, searchWord) {
			return i + 1
		}
	}
	return -1
}
```

## 1460 — Make Two Arrays Equal By Reversing Subarrays

```go
package main

// LeetCode #1460: Make Two Arrays Equal by Reversing Subarrays
// https://leetcode.com/problems/make-two-arrays-equal-by-reversing-subarrays/
// Difficulty: Easy
//
// LeetCode submission: func canBeEqual(target []int, arr []int) bool

import "fmt"

func main() {
	fmt.Println(MakeTwoArraysEqualByReversingSubarrays([]int{1, 2, 3, 4}, []int{2, 4, 1, 3})) // true
	fmt.Println(MakeTwoArraysEqualByReversingSubarrays([]int{7}, []int{7}))                    // true
	fmt.Println(MakeTwoArraysEqualByReversingSubarrays([]int{3, 7, 9}, []int{3, 7, 11}))       // false
}

// Time: O(n), Space: O(n)
func MakeTwoArraysEqualByReversingSubarrays(target []int, arr []int) bool {
	if len(target) != len(arr) {
		return false
	}
	freq := make(map[int]int, len(target))
	for _, v := range target {
		freq[v]++
	}
	for _, v := range arr {
		freq[v]--
		if freq[v] < 0 {
			return false
		}
	}
	return true
}
```

## 1464 — Maximum Product Of Two Elements In An Array

```go
package main

// LeetCode #1464: Maximum Product of Two Elements in an Array
// https://leetcode.com/problems/maximum-product-of-two-elements-in-an-array/
// Difficulty: Easy
//
// LeetCode submission: func maxProduct(nums []int) int

import "fmt"

func main() {
	fmt.Println(MaximumProductOfTwoElementsInAnArray([]int{3, 4, 5, 2})) // 12
	fmt.Println(MaximumProductOfTwoElementsInAnArray([]int{1, 5, 4, 5})) // 16
	fmt.Println(MaximumProductOfTwoElementsInAnArray([]int{3, 7}))       // 12
}

// Time: O(n), Space: O(1)
func MaximumProductOfTwoElementsInAnArray(nums []int) int {
	first, second := 0, 0
	for _, v := range nums {
		if v > first {
			second = first
			first = v
		} else if v > second {
			second = v
		}
	}
	return (first - 1) * (second - 1)
}
```

## 1469 — Find All The Lonely Nodes

```go
package main

// LeetCode #1469: Find All The Lonely Nodes
// https://leetcode.com/problems/find-all-the-lonely-nodes/
// Difficulty: Easy [Paid]
//
// LeetCode submission: func getLonelyNodes(root *TreeNode) []int

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Tree: [1,2,3,null,4]
	root := &TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2, Right: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 3},
	}
	fmt.Println(FindAllTheLonelyNodes(root)) // [4]

	// Tree: [7,1,4,6,null,5,3,null,null,null,null,null,2]
	root2 := &TreeNode{Val: 7,
		Left:  &TreeNode{Val: 1, Left: &TreeNode{Val: 6}},
		Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 2}}},
	}
	fmt.Println(FindAllTheLonelyNodes(root2)) // [6 5 2]
}

// Time: O(n), Space: O(h) recursive + O(n) output
func FindAllTheLonelyNodes(root *TreeNode) []int {
	res := make([]int, 0)
	collectLonely(root, &res)
	return res
}

func collectLonely(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	if node.Left != nil && node.Right == nil {
		*res = append(*res, node.Left.Val)
	}
	if node.Left == nil && node.Right != nil {
		*res = append(*res, node.Right.Val)
	}
	collectLonely(node.Left, res)
	collectLonely(node.Right, res)
}
```

## 1470 — Shuffle The Array

```go
package main

// LeetCode #1470: Shuffle the Array
// https://leetcode.com/problems/shuffle-the-array/
// Difficulty: Easy
//
// LeetCode submission: func shuffle(nums []int, n int) []int

import "fmt"

func main() {
	fmt.Println(ShuffleTheArray([]int{2, 5, 1, 3, 4, 7}, 3)) // [2 3 5 4 1 7]
	fmt.Println(ShuffleTheArray([]int{1, 2, 3, 4, 4, 3, 2, 1}, 4)) // [1 4 2 3 3 2 4 1]
}

// Time: O(n), Space: O(n)
func ShuffleTheArray(nums []int, n int) []int {
	res := make([]int, 2*n)
	for i := 0; i < n; i++ {
		res[2*i] = nums[i]
		res[2*i+1] = nums[i+n]
	}
	return res
}
```

## 1474 — Delete N Nodes After M Nodes Of A Linked List

```go
package main

// LeetCode #1474: Delete N Nodes After M Nodes of a Linked List
// https://leetcode.com/problems/delete-n-nodes-after-m-nodes-of-a-linked-list/
// Difficulty: Easy [Paid]
//
// LeetCode submission: func deleteNodes(head *ListNode, m int, n int) *ListNode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// List: 1->2->3->4->5->6->7->8->9->10->11, m=2, n=3
	// Expected: 1->2->6->7->11
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 7, Next: &ListNode{Val: 8, Next: &ListNode{Val: 9, Next: &ListNode{Val: 10, Next: &ListNode{Val: 11}}}}}}}}}}}
	result := DeleteNNodesAfterMNodesOfALinkedList(head, 2, 3)
	for result != nil {
		fmt.Print(result.Val, " ")
		result = result.Next
	}
	fmt.Println() // 1 2 6 7 11
}

// Time: O(n), Space: O(1)
func DeleteNNodesAfterMNodesOfALinkedList(head *ListNode, m int, n int) *ListNode {
	cur := head
	for cur != nil {
		// Skip m nodes
		for i := 1; i < m && cur != nil; i++ {
			cur = cur.Next
		}
		if cur == nil {
			break
		}
		// Delete next n nodes
		temp := cur.Next
		for i := 0; i < n && temp != nil; i++ {
			temp = temp.Next
		}
		cur.Next = temp
		cur = temp
	}
	return head
}
```

## 1475 — Final Prices With A Special Discount In A Shop

```go
package main

// LeetCode #1475: Final Prices With a Special Discount in a Shop
// https://leetcode.com/problems/final-prices-with-a-special-discount-in-a-shop/
// Difficulty: Easy
//
// LeetCode submission: func finalPrices(prices []int) []int

import "fmt"

func main() {
	fmt.Println(FinalPricesWithASpecialDiscountInAShop([]int{8, 4, 6, 2, 3})) // [4 2 4 2 3]
	fmt.Println(FinalPricesWithASpecialDiscountInAShop([]int{1, 2, 3, 4, 5})) // [1 2 3 4 5]
}

// Time: O(n^2), Space: O(1) excluding output
func FinalPricesWithASpecialDiscountInAShop(prices []int) []int {
	res := make([]int, len(prices))
	for i := 0; i < len(prices); i++ {
		res[i] = prices[i]
		for j := i + 1; j < len(prices); j++ {
			if prices[j] <= prices[i] {
				res[i] -= prices[j]
				break
			}
		}
	}
	return res
}
```

## 1480 — Running Sum Of 1d Array

```go
package main

// LeetCode #1480: Running Sum of 1d Array
// https://leetcode.com/problems/running-sum-of-1d-array/
// Difficulty: Easy
//
// LeetCode submission: func runningSum(nums []int) []int

import "fmt"

func main() {
	fmt.Println(RunningSumOfOneDArray([]int{1, 2, 3, 4}))    // [1 3 6 10]
	fmt.Println(RunningSumOfOneDArray([]int{1, 1, 1, 1, 1})) // [1 2 3 4 5]
}

// Time: O(n), Space: O(1) excluding output
func RunningSumOfOneDArray(nums []int) []int {
	res := make([]int, len(nums))
	sum := 0
	for i, v := range nums {
		sum += v
		res[i] = sum
	}
	return res
}
```

## 1484 — Group Sold Products By The Date

```go
package main

// LeetCode #1484: Group Sold Products By The Date
// https://leetcode.com/problems/group-sold-products-by-the-date/
// Difficulty: Easy
//
// This is a SQL problem. The solution is the SQL query below.
// Table: Activities (sell_date, product)

import "fmt"

func main() {
	fmt.Println(GroupSoldProductsByTheDate())
}

// Time: N/A (SQL query), Space: N/A
func GroupSoldProductsByTheDate() string {
	return `SELECT
  sell_date,
  COUNT(DISTINCT product) AS num_sold,
  GROUP_CONCAT(DISTINCT product ORDER BY product SEPARATOR ',') AS products
FROM Activities
GROUP BY sell_date
ORDER BY sell_date;`
}
```

## 1486 — Xor Operation In An Array

```go
package main

// LeetCode #1486: XOR Operation in an Array
// https://leetcode.com/problems/xor-operation-in-an-array/
// Difficulty: Easy
//
// LeetCode submission: func xorOperation(n int, start int) int

import "fmt"

func main() {
	fmt.Println(XorOperationInAnArray(5, 0)) // 8
	fmt.Println(XorOperationInAnArray(4, 3)) // 8
	fmt.Println(XorOperationInAnArray(1, 7)) // 7
}

// Time: O(n), Space: O(1)
func XorOperationInAnArray(n int, start int) int {
	res := 0
	for i := 0; i < n; i++ {
		res ^= start + 2*i
	}
	return res
}
```

## 1491 — Average Salary Excluding The Minimum And Maximum Salary

```go
package main

// LeetCode #1491: Average Salary Excluding the Minimum and Maximum Salary
// https://leetcode.com/problems/average-salary-excluding-the-minimum-and-maximum-salary/
// Difficulty: Easy
//
// LeetCode submission: func average(salary []int) float64

import "fmt"

func main() {
	fmt.Println(AverageSalaryExcludingTheMinimumAndMaximumSalary([]int{4000, 3000, 1000, 2000})) // 2500
	fmt.Println(AverageSalaryExcludingTheMinimumAndMaximumSalary([]int{1000, 2000, 3000}))       // 2000
}

// Time: O(n), Space: O(1)
func AverageSalaryExcludingTheMinimumAndMaximumSalary(salary []int) float64 {
	min, max := salary[0], salary[0]
	sum := 0
	for _, v := range salary {
		sum += v
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return float64(sum-min-max) / float64(len(salary)-2)
}
```

## 1495 — Friendly Movies Streamed Last Month

```go
package main

// LeetCode #1495: Friendly Movies Streamed Last Month
// https://leetcode.com/problems/friendly-movies-streamed-last-month/
// Difficulty: Easy [Paid]
//
// This is a SQL problem. The solution is the SQL query below.
// Tables: TVProgram (program_date, content_id, channel), Content (content_id, title, Kids_content, content_type)

import "fmt"

func main() {
	fmt.Println(FriendlyMoviesStreamedLastMonth())
}

// Time: N/A (SQL query), Space: N/A
func FriendlyMoviesStreamedLastMonth() string {
	return `SELECT DISTINCT c.title
FROM TVProgram p
JOIN Content c ON p.content_id = c.content_id
WHERE c.Kids_content = 'Y'
  AND c.content_type = 'Movies'
  AND p.program_date BETWEEN '2020-06-01' AND '2020-06-30';`
}
```

## 1496 — Path Crossing

```go
package main

// LeetCode #1496: Path Crossing
// https://leetcode.com/problems/path-crossing/
// Difficulty: Easy
//
// LeetCode submission: func isPathCrossing(path string) bool

import "fmt"

func main() {
	fmt.Println(PathCrossing("NES"))   // false
	fmt.Println(PathCrossing("NESWW")) // true
}

// Time: O(n), Space: O(n)
func PathCrossing(path string) bool {
	visited := make(map[[2]int]bool)
	x, y := 0, 0
	visited[[2]int{0, 0}] = true
	for _, ch := range path {
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
		if visited[[2]int{x, y}] {
			return true
		}
		visited[[2]int{x, y}] = true
	}
	return false
}
```

## 1502 — Can Make Arithmetic Progression From Sequence

```go
package main

// LeetCode #1502: Can Make Arithmetic Progression From Sequence
// https://leetcode.com/problems/can-make-arithmetic-progression-from-sequence/
// Difficulty: Easy
//
// LeetCode submission: func canMakeArithmeticProgression(arr []int) bool

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(CanMakeArithmeticProgressionFromSequence([]int{3, 5, 1})) // true
	fmt.Println(CanMakeArithmeticProgressionFromSequence([]int{1, 2, 4})) // false
}

// Time: O(n log n), Space: O(1)
func CanMakeArithmeticProgressionFromSequence(arr []int) bool {
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

## 1507 — Reformat Date

```go
package main

// LeetCode #1507: Reformat Date
// https://leetcode.com/problems/reformat-date/
// Difficulty: Easy
//
// LeetCode submission: func reformatDate(date string) string

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(ReformatDate("20th Oct 2052")) // "2052-10-20"
	fmt.Println(ReformatDate("6th Jun 1933"))  // "1933-06-06"
	fmt.Println(ReformatDate("26th May 1960")) // "1960-05-26"
}

// Time: O(1), Space: O(1)
func ReformatDate(date string) string {
	months := map[string]string{
		"Jan": "01", "Feb": "02", "Mar": "03", "Apr": "04",
		"May": "05", "Jun": "06", "Jul": "07", "Aug": "08",
		"Sep": "09", "Oct": "10", "Nov": "11", "Dec": "12",
	}
	parts := strings.Split(date, " ")
	day := parts[0][:len(parts[0])-2]
	if len(day) == 1 {
		day = "0" + day
	}
	return parts[2] + "-" + months[parts[1]] + "-" + day
}
```

## 1511 — Customer Order Frequency

```go
package main

// LeetCode #1511: Customer Order Frequency
// https://leetcode.com/problems/customer-order-frequency/
// Difficulty: Easy [Paid]
//
// This is a SQL problem. The solution is the SQL query below.
// Tables: Customers (customer_id, name, country), Product (product_id, description, price), Orders (order_id, customer_id, product_id, order_date, quantity)

import "fmt"

func main() {
	fmt.Println(CustomerOrderFrequency())
}

// Time: N/A (SQL query), Space: N/A
func CustomerOrderFrequency() string {
	return `SELECT c.customer_id, c.name
FROM Customers c
JOIN Orders o ON c.customer_id = o.customer_id
JOIN Product p ON o.product_id = p.product_id
WHERE o.order_date BETWEEN '2020-06-01' AND '2020-07-31'
GROUP BY c.customer_id, c.name
HAVING SUM(CASE WHEN o.order_date BETWEEN '2020-06-01' AND '2020-06-30' THEN o.quantity * p.price END) >= 100
   AND SUM(CASE WHEN o.order_date BETWEEN '2020-07-01' AND '2020-07-31' THEN o.quantity * p.price END) >= 100;`
}
```

## 1512 — Number Of Good Pairs

```go
package main

// LeetCode #1512: Number of Good Pairs
// https://leetcode.com/problems/number-of-good-pairs/
// Difficulty: Easy
//
// LeetCode submission: func numIdenticalPairs(nums []int) int

import "fmt"

func main() {
	fmt.Println(NumberOfGoodPairs([]int{1, 2, 3, 1, 1, 3})) // 4
	fmt.Println(NumberOfGoodPairs([]int{1, 1, 1, 1}))        // 6
	fmt.Println(NumberOfGoodPairs([]int{1, 2, 3}))           // 0
}

// Time: O(n), Space: O(n)
func NumberOfGoodPairs(nums []int) int {
	freq := make(map[int]int)
	count := 0
	for _, v := range nums {
		count += freq[v]
		freq[v]++
	}
	return count
}
```

## 1517 — Find Users With Valid E Mails

```go
package main

// LeetCode #1517: Find Users With Valid E-Mails
// https://leetcode.com/problems/find-users-with-valid-e-mails/
// Difficulty: Easy
//
// This is a SQL problem. The solution is the SQL query below.
// Table: Users (user_id, name, mail)

import "fmt"

func main() {
	fmt.Println(FindUsersWithValidEMails())
}

// Time: N/A (SQL query), Space: N/A
func FindUsersWithValidEMails() string {
	return `SELECT user_id, name, mail
FROM Users
WHERE mail REGEXP '^[A-Za-z][A-Za-z0-9_.-]*@leetcode\\.com$';`
}
```

## 1518 — Water Bottles

```go
package main

// LeetCode #1518: Water Bottles
// https://leetcode.com/problems/water-bottles/
// Difficulty: Easy
//
// LeetCode submission: func numWaterBottles(numBottles int, numExchange int) int

import "fmt"

func main() {
	fmt.Println(WaterBottles(9, 3))  // 13
	fmt.Println(WaterBottles(15, 4)) // 19
	fmt.Println(WaterBottles(5, 5))  // 6
}

// Time: O(log n), Space: O(1)
func WaterBottles(numBottles int, numExchange int) int {
	total := numBottles
	empty := numBottles
	for empty >= numExchange {
		newBottles := empty / numExchange
		total += newBottles
		empty = newBottles + empty%numExchange
	}
	return total
}
```

## 1523 — Count Odd Numbers In An Interval Range

```go
package main

// LeetCode #1523: Count Odd Numbers in an Interval Range
// https://leetcode.com/problems/count-odd-numbers-in-an-interval-range/
// Difficulty: Easy
//
// LeetCode submission: func countOdds(low int, high int) int

import "fmt"

func main() {
	fmt.Println(CountOddNumbersInAnIntervalRange(3, 7)) // 3
	fmt.Println(CountOddNumbersInAnIntervalRange(8, 10)) // 1
	fmt.Println(CountOddNumbersInAnIntervalRange(0, 0)) // 0
}

// Time: O(1), Space: O(1)
func CountOddNumbersInAnIntervalRange(low int, high int) int {
	return (high + 1) / 2 - low / 2
}
```

## 1527 — Patients With A Condition

```go
package main

// LeetCode #1527: Patients With a Condition
// https://leetcode.com/problems/patients-with-a-condition/
// Difficulty: Easy
//
// This is a SQL problem. The solution is the SQL query below.
// Table: Patients (patient_id, patient_name, conditions)

import "fmt"

func main() {
	fmt.Println(PatientsWithACondition())
}

// Time: N/A (SQL query), Space: N/A
func PatientsWithACondition() string {
	return `SELECT patient_id, patient_name, conditions
FROM Patients
WHERE conditions LIKE 'DIAB1%'
   OR conditions LIKE '% DIAB1%';`
}
```

## 1528 — Shuffle String

```go
package main

// LeetCode #1528: Shuffle String
// https://leetcode.com/problems/shuffle-string/
// Difficulty: Easy
//
// LeetCode submission: func restoreString(s string, indices []int) string

import "fmt"

func main() {
	fmt.Println(ShuffleString("codeleet", []int{4, 5, 6, 7, 0, 2, 1, 3})) // "leetcode"
	fmt.Println(ShuffleString("abc", []int{0, 1, 2}))                     // "abc"
}

// Time: O(n), Space: O(n)
func ShuffleString(s string, indices []int) string {
	res := make([]byte, len(s))
	for i, idx := range indices {
		res[idx] = s[i]
	}
	return string(res)
}
```

## 1534 — Count Good Triplets

```go
package main

// LeetCode #1534: Count Good Triplets
// https://leetcode.com/problems/count-good-triplets/
// Difficulty: Easy
//
// LeetCode submission: func countGoodTriplets(arr []int, a int, b int, c int) int

import "fmt"

func main() {
	fmt.Println(CountGoodTriplets([]int{3, 0, 1, 1, 9, 7}, 7, 2, 3)) // 4
	fmt.Println(CountGoodTriplets([]int{1, 1, 2, 2, 3}, 0, 0, 1))   // 0
}

// Time: O(n^3), Space: O(1)
func CountGoodTriplets(arr []int, a int, b int, c int) int {
	n, count := len(arr), 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if abs(arr[i]-arr[j]) > a {
				continue
			}
			for k := j + 1; k < n; k++ {
				if abs(arr[j]-arr[k]) <= b && abs(arr[i]-arr[k]) <= c {
					count++
				}
			}
		}
	}
	return count
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

## 1539 — Kth Missing Positive Number

```go
package main

// LeetCode #1539: Kth Missing Positive Number
// https://leetcode.com/problems/kth-missing-positive-number/
// Difficulty: Easy
//
// LeetCode submission: func findKthPositive(arr []int, k int) int

import "fmt"

func main() {
	fmt.Println(KthMissingPositiveNumber([]int{2, 3, 4, 7, 11}, 5)) // 9
	fmt.Println(KthMissingPositiveNumber([]int{1, 2, 3, 4}, 2))     // 6
	fmt.Println(KthMissingPositiveNumber([]int{1, 3, 5}, 2))        // 4
}

// Time: O(log n), Space: O(1)
func KthMissingPositiveNumber(arr []int, k int) int {
	lo, hi := 0, len(arr)
	for lo < hi {
		mid := (lo + hi) / 2
		if arr[mid]-mid-1 < k {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo + k
}
```

## 1543 — Fix Product Name Format

```go
package main

// LeetCode #1543: Fix Product Name Format
// https://leetcode.com/problems/fix-product-name-format/
// Difficulty: Easy [Paid]
//
// This is a SQL problem. The solution is the SQL query below.
// Table: Products (product_id, product_name, price)

import "fmt"

func main() {
	fmt.Println(FixProductNameFormat())
}

// Time: N/A (SQL query), Space: N/A
func FixProductNameFormat() string {
	return `SELECT
  TRIM(LOWER(product_name)) AS product_name,
  DATE_FORMAT(sale_date, '%Y-%m') AS sale_date,
  COUNT(*) AS total
FROM Sales
GROUP BY TRIM(LOWER(product_name)), DATE_FORMAT(sale_date, '%Y-%m')
ORDER BY product_name, sale_date;`
}
```

## 1544 — Make The String Great

```go
package main

// LeetCode #1544: Make The String Great
// https://leetcode.com/problems/make-the-string-great/
// Difficulty: Easy
//
// LeetCode submission: func makeGood(s string) string

import "fmt"

func main() {
	fmt.Println(MakeTheStringGreat("leEeetcode")) // "leetcode"
	fmt.Println(MakeTheStringGreat("abBAcC"))     // ""
	fmt.Println(MakeTheStringGreat("s"))          // "s"
}

// Time: O(n), Space: O(n)
func MakeTheStringGreat(s string) string {
	stack := make([]byte, 0, len(s))
	for i := range s {
		stack = append(stack, s[i])
		n := len(stack)
		if n >= 2 {
			diff := int(stack[n-1]) - int(stack[n-2])
			if diff == 32 || diff == -32 {
				stack = stack[:n-2]
			}
		}
	}
	return string(stack)
}
```

## 1550 — Three Consecutive Odds

```go
package main

// LeetCode #1550: Three Consecutive Odds
// https://leetcode.com/problems/three-consecutive-odds/
// Difficulty: Easy
//
// LeetCode submission: func threeConsecutiveOdds(arr []int) bool

import "fmt"

func main() {
	fmt.Println(ThreeConsecutiveOdds([]int{2, 6, 4, 1}))  // false
	fmt.Println(ThreeConsecutiveOdds([]int{1, 2, 34, 3, 4, 5, 7, 23, 12})) // true
}

// Time: O(n), Space: O(1)
func ThreeConsecutiveOdds(arr []int) bool {
	count := 0
	for _, v := range arr {
		if v%2 == 1 {
			count++
			if count == 3 {
				return true
			}
		} else {
			count = 0
		}
	}
	return false
}
```

## 1556 — Thousand Separator

```go
package main

// LeetCode #1556: Thousand Separator
// https://leetcode.com/problems/thousand-separator/
// Difficulty: Easy
//
// LeetCode submission: func thousandSeparator(n int) string

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(ThousandSeparator(987))      // "987"
	fmt.Println(ThousandSeparator(1234))     // "1.234"
	fmt.Println(ThousandSeparator(1000000))  // "1.000.000"
}

// Time: O(log n), Space: O(log n)
func ThousandSeparator(n int) string {
	s := strconv.Itoa(n)
	res := make([]byte, 0, len(s)+len(s)/3)
	for i, ch := range s {
		if i > 0 && (len(s)-i)%3 == 0 {
			res = append(res, '.')
		}
		res = append(res, byte(ch))
	}
	return string(res)
}
```

## 1560 — Most Visited Sector In A Circular Track

```go
package main

// LeetCode #1560: Most Visited Sector in a Circular Track
// https://leetcode.com/problems/most-visited-sector-in-a-circular-track/
// Difficulty: Easy
//
// LeetCode submission: func mostVisited(n int, rounds []int) []int

import "fmt"

func main() {
	fmt.Println(MostVisitedSectorInACircularTrack(4, []int{1, 3, 1, 2})) // [1 2]
	fmt.Println(MostVisitedSectorInACircularTrack(2, []int{2, 1, 2, 1, 2, 1, 2, 1, 2})) // [2]
	fmt.Println(MostVisitedSectorInACircularTrack(7, []int{1, 3, 5, 7})) // [1 2 3 4 5 6 7]
}

// Time: O(n), Space: O(n)
func MostVisitedSectorInACircularTrack(n int, rounds []int) []int {
	start, end := rounds[0], rounds[len(rounds)-1]
	if start <= end {
		res := make([]int, end-start+1)
		for i := range res {
			res[i] = start + i
		}
		return res
	}
	res := make([]int, 0, n-end+start)
	for i := 1; i <= end; i++ {
		res = append(res, i)
	}
	for i := start; i <= n; i++ {
		res = append(res, i)
	}
	return res
}
```

## 1565 — Unique Orders And Customers Per Month

```go
package main

// LeetCode #1565: Unique Orders and Customers Per Month
// https://leetcode.com/problems/unique-orders-and-customers-per-month/
// Difficulty: Easy [Paid]
//
// This is a SQL problem. The solution is the SQL query below.
// Table: Orders (order_id, order_date, customer_id, invoice)

import "fmt"

func main() {
	fmt.Println(UniqueOrdersAndCustomersPerMonth())
}

// Time: N/A (SQL query), Space: N/A
func UniqueOrdersAndCustomersPerMonth() string {
	return `SELECT
  DATE_FORMAT(order_date, '%Y-%m') AS month,
  COUNT(DISTINCT order_id) AS order_count,
  COUNT(DISTINCT customer_id) AS customer_count
FROM Orders
WHERE invoice > 20
GROUP BY month
ORDER BY month;`
}
```

## 1566 — Detect Pattern Of Length M Repeated K Or More Times

```go
package main

// LeetCode #1566: Detect Pattern of Length M Repeated K or More Times
// https://leetcode.com/problems/detect-pattern-of-length-m-repeated-k-or-more-times/
// Difficulty: Easy
//
// LeetCode submission: func containsPattern(arr []int, m int, k int) bool

import "fmt"

func main() {
	fmt.Println(DetectPatternOfLengthMRepeatedKOrMoreTimes([]int{1, 2, 4, 4, 4, 4}, 1, 3)) // true
	fmt.Println(DetectPatternOfLengthMRepeatedKOrMoreTimes([]int{1, 2, 1, 2, 1, 1, 1, 3}, 2, 2)) // true
	fmt.Println(DetectPatternOfLengthMRepeatedKOrMoreTimes([]int{1, 2, 1, 2, 1, 3}, 2, 3)) // false
}

// Time: O(n * m), Space: O(1)
func DetectPatternOfLengthMRepeatedKOrMoreTimes(arr []int, m int, k int) bool {
	n := len(arr)
	if m*k > n {
		return false
	}
	for i := 0; i <= n-m*k; i++ {
		count := 1
		for j := i + m; j <= n-m; j += m {
			match := true
			for t := 0; t < m; t++ {
				if arr[j+t] != arr[i+t] {
					match = false
					break
				}
			}
			if match {
				count++
				if count >= k {
					return true
				}
			} else {
				break
			}
		}
	}
	return false
}
```

## 1571 — Warehouse Manager

```go
package main

// LeetCode #1571: Warehouse Manager
// https://leetcode.com/problems/warehouse-manager/
// Difficulty: Easy [Paid]
//
// This is a SQL problem. The solution is the SQL query below.
// Tables: Warehouse (name, product_id, units), Products (product_id, product_name, Width, Length, Height)

import "fmt"

func main() {
	fmt.Println(WarehouseManager())
}

// Time: N/A (SQL query), Space: N/A
func WarehouseManager() string {
	return `SELECT w.name AS warehouse_name, SUM(w.units * p.Width * p.Length * p.Height) AS volume
FROM Warehouse w
JOIN Products p ON w.product_id = p.product_id
GROUP BY w.name
ORDER BY warehouse_name;`
}
```

## 1572 — Matrix Diagonal Sum

```go
package main

// LeetCode #1572: Matrix Diagonal Sum
// https://leetcode.com/problems/matrix-diagonal-sum/
// Difficulty: Easy
//
// LeetCode submission: func diagonalSum(mat [][]int) int

import "fmt"

func main() {
	mat1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(MatrixDiagonalSum(mat1)) // 25

	mat2 := [][]int{
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
	}
	fmt.Println(MatrixDiagonalSum(mat2)) // 8
}

// Time: O(n), Space: O(1)
func MatrixDiagonalSum(mat [][]int) int {
	n := len(mat)
	sum := 0
	for i := 0; i < n; i++ {
		sum += mat[i][i]
		sum += mat[i][n-1-i]
	}
	if n%2 == 1 {
		mid := n / 2
		sum -= mat[mid][mid]
	}
	return sum
}
```

## 1576 — Replace All S To Avoid Consecutive Repeating Characters

```go
package main

// LeetCode #1576: Replace All ?'s to Avoid Consecutive Repeating Characters
// https://leetcode.com/problems/replace-all-s-to-avoid-consecutive-repeating-characters/
// Difficulty: Easy
//
// LeetCode submission: func modifyString(s string) string

import "fmt"

func main() {
	fmt.Println(ReplaceAllSToAvoidConsecutiveRepeatingCharacters("?zs")) // "azs"
	fmt.Println(ReplaceAllSToAvoidConsecutiveRepeatingCharacters("ubv?w")) // "ubvaw"
	fmt.Println(ReplaceAllSToAvoidConsecutiveRepeatingCharacters("??yw?ipkj?")) // "abywcipkja"
}

// Time: O(n), Space: O(n)
func ReplaceAllSToAvoidConsecutiveRepeatingCharacters(s string) string {
	res := []byte(s)
	for i, ch := range res {
		if ch == '?' {
			for c := byte('a'); c <= 'z'; c++ {
				if (i == 0 || res[i-1] != c) && (i == len(res)-1 || res[i+1] != c) {
					res[i] = c
					break
				}
			}
		}
	}
	return string(res)
}
```

## 1581 — Customer Who Visited But Did Not Make Any Transactions

```go
package main

// LeetCode #1581: Customer Who Visited but Did Not Make Any Transactions
// https://leetcode.com/problems/customer-who-visited-but-did-not-make-any-transactions/
// Difficulty: Easy
//
// This is a SQL problem. The solution is the SQL query below.
// Tables: Visits (visit_id, customer_id), Transactions (transaction_id, visit_id, amount)

import "fmt"

func main() {
	fmt.Println(CustomerWhoVisitedButDidNotMakeAnyTransactions())
}

// Time: N/A (SQL query), Space: N/A
func CustomerWhoVisitedButDidNotMakeAnyTransactions() string {
	return `SELECT v.customer_id, COUNT(*) AS count_no_trans
FROM Visits v
LEFT JOIN Transactions t ON v.visit_id = t.visit_id
WHERE t.visit_id IS NULL
GROUP BY v.customer_id;`
}
```

## 1582 — Special Positions In A Binary Matrix

```go
package main

// LeetCode #1582: Special Positions in a Binary Matrix
// https://leetcode.com/problems/special-positions-in-a-binary-matrix/
// Difficulty: Easy
//
// LeetCode submission: func numSpecial(mat [][]int) int

import "fmt"

func main() {
	mat1 := [][]int{
		{1, 0, 0},
		{0, 0, 1},
		{1, 0, 0},
	}
	fmt.Println(SpecialPositionsInABinaryMatrix(mat1)) // 1

	mat2 := [][]int{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
	fmt.Println(SpecialPositionsInABinaryMatrix(mat2)) // 3
}

// Time: O(m * n), Space: O(m + n)
func SpecialPositionsInABinaryMatrix(mat [][]int) int {
	m, n := len(mat), len(mat[0])
	rows := make([]int, m)
	cols := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 1 {
				rows[i]++
				cols[j]++
			}
		}
	}
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 1 && rows[i] == 1 && cols[j] == 1 {
				count++
			}
		}
	}
	return count
}
```

## 1587 — Bank Account Summary Ii

```go
package main

// LeetCode #1587: Bank Account Summary II
// https://leetcode.com/problems/bank-account-summary-ii/
// Difficulty: Easy
//
// This is a SQL problem. The solution is the SQL query below.
// Tables: Users (account, name), Transactions (trans_id, account, amount, transacted_on)

import "fmt"

func main() {
	fmt.Println(BankAccountSummaryIi())
}

// Time: N/A (SQL query), Space: N/A
func BankAccountSummaryIi() string {
	return `SELECT u.name, SUM(t.amount) AS balance
FROM Users u
JOIN Transactions t ON u.account = t.account
GROUP BY u.account, u.name
HAVING SUM(t.amount) > 10000;`
}
```

## 1588 — Sum Of All Odd Length Subarrays

```go
package main

// LeetCode #1588: Sum of All Odd Length Subarrays
// https://leetcode.com/problems/sum-of-all-odd-length-subarrays/
// Difficulty: Easy
//
// LeetCode submission: func sumOddLengthSubarrays(arr []int) int

import "fmt"

func main() {
	fmt.Println(SumOfAllOddLengthSubarrays([]int{1, 4, 2, 5, 3})) // 58
	fmt.Println(SumOfAllOddLengthSubarrays([]int{1, 2}))          // 3
	fmt.Println(SumOfAllOddLengthSubarrays([]int{10, 11, 12}))    // 66
}

// Time: O(n), Space: O(1)
func SumOfAllOddLengthSubarrays(arr []int) int {
	n := len(arr)
	sum := 0
	for i, v := range arr {
		contribution := ((i+1)*(n-i) + 1) / 2
		sum += v * contribution
	}
	return sum
}
```

## 1592 — Rearrange Spaces Between Words

```go
package main

// LeetCode #1592: Rearrange Spaces Between Words
// https://leetcode.com/problems/rearrange-spaces-between-words/
// Difficulty: Easy
//
// LeetCode submission: func reorderSpaces(text string) string

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(RearrangeSpacesBetweenWords("  this   is  a sentence ")) // "this   is   a   sentence"
	fmt.Println(RearrangeSpacesBetweenWords(" practice   makes   perfect")) // "practice   makes   perfect "
}

// Time: O(n), Space: O(n)
func RearrangeSpacesBetweenWords(text string) string {
	words := strings.Fields(text)
	spaces := strings.Count(text, " ")
	if len(words) == 1 {
		return words[0] + strings.Repeat(" ", spaces)
	}
	between := spaces / (len(words) - 1)
	extra := spaces % (len(words) - 1)
	res := strings.Join(words, strings.Repeat(" ", between))
	res += strings.Repeat(" ", extra)
	return res
}
```

## 1598 — Crawler Log Folder

```go
package main

// LeetCode #1598: Crawler Log Folder
// https://leetcode.com/problems/crawler-log-folder/
// Difficulty: Easy
//
// LeetCode submission: func minOperations(logs []string) int

import "fmt"

func main() {
	fmt.Println(CrawlerLogFolder([]string{"d1/", "d2/", "../", "d21/", "./"}))            // 2
	fmt.Println(CrawlerLogFolder([]string{"d1/", "../", "../", "../"}))                   // 0
	fmt.Println(CrawlerLogFolder([]string{"./", "../", "./"}))                            // 0
}

// Time: O(n), Space: O(1)
func CrawlerLogFolder(logs []string) int {
	depth := 0
	for _, log := range logs {
		switch log {
		case "../":
			if depth > 0 {
				depth--
			}
		case "./":
			// do nothing
		default:
			depth++
		}
	}
	return depth
}
```

