# Easy (Mudah) — Problem ��1979

## 1603 — Design Parking System

```go
package main

// LeetCode #1603: Design Parking System
// https://leetcode.com/problems/design-parking-system/
// Difficulty: Easy

import "fmt"

type ParkingSystem struct {
	spots [3]int
}

func NewParkingSystem(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{spots: [3]int{big, medium, small}}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if this.spots[carType-1] > 0 {
		this.spots[carType-1]--
		return true
	}
	return false
}

func main() {
	ps := NewParkingSystem(1, 1, 0)
	fmt.Println(ps.AddCar(1))
	fmt.Println(ps.AddCar(2))
	fmt.Println(ps.AddCar(3))
	fmt.Println(ps.AddCar(1))
}
```

## 1607 — Sellers With No Sales

```go
package main

// LeetCode #1607: Sellers With No Sales
// https://leetcode.com/problems/sellers-with-no-sales/
// Difficulty: Easy [Paid]

import "fmt"

func main() {
	fmt.Println(SellersWithNoSales())
}

func SellersWithNoSales() any {
	// TODO: implement
	return nil
}
```

## 1608 — Special Array With X Elements Greater Than Or Equal X

```go
package main

// LeetCode #1608: Special Array With X Elements Greater Than or Equal X
// https://leetcode.com/problems/special-array-with-x-elements-greater-than-or-equal-x/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

// Time: O(n log n), Space: O(log n) (for sorting)
func SpecialArray(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	for i := 0; i < len(nums); i++ {
		if nums[i] >= i+1 {
			if i == len(nums)-1 || nums[i+1] < i+1 {
				return i + 1
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(SpecialArray([]int{3, 5}))
	fmt.Println(SpecialArray([]int{0, 0}))
	fmt.Println(SpecialArray([]int{0, 4, 3, 0, 4}))
}
```

## 1614 — Maximum Nesting Depth Of The Parentheses

```go
package main

// LeetCode #1614: Maximum Nesting Depth of the Parentheses
// https://leetcode.com/problems/maximum-nesting-depth-of-the-parentheses/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func MaximumNestingDepthOfTheParentheses(s string) int {
	maxDepth, currentDepth := 0, 0
	for _, ch := range s {
		if ch == '(' {
			currentDepth++
			if currentDepth > maxDepth {
				maxDepth = currentDepth
			}
		} else if ch == ')' {
			currentDepth--
		}
	}
	return maxDepth
}

func main() {
	fmt.Println(MaximumNestingDepthOfTheParentheses("(1+(2*3)+((8)/4))+1"))
	fmt.Println(MaximumNestingDepthOfTheParentheses("(1)+((2))+(((3)))"))
	fmt.Println(MaximumNestingDepthOfTheParentheses(""))
}
```

## 1619 — Mean Of Array After Removing Some Elements

```go
package main

// LeetCode #1619: Mean of Array After Removing Some Elements
// https://leetcode.com/problems/mean-of-array-after-removing-some-elements/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

// Time: O(n log n), Space: O(log n) (for sorting)
func TrimMean(arr []int) float64 {
	sort.Ints(arr)
	n := len(arr)
	remove := n / 20
	sum := 0
	for i := remove; i < n-remove; i++ {
		sum += arr[i]
	}
	return float64(sum) / float64(n-2*remove)
}

func main() {
	fmt.Println(TrimMean([]int{1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3}))
	fmt.Println(TrimMean([]int{6, 2, 7, 5, 1, 2, 0, 3, 10, 2, 5, 0, 5, 5, 0, 8, 7, 6, 8, 0}))
	fmt.Println(TrimMean([]int{6, 0, 7, 0, 7, 5, 7, 8, 3, 4, 0, 7, 8, 1, 6, 8, 1, 1, 2, 4, 8, 1, 9, 5, 4, 3, 8, 5, 10, 8, 6, 6, 1, 0, 6, 10, 8, 2, 3, 4}))
}
```

## 1623 — All Valid Triplets That Can Represent A Country

```go
package main

// LeetCode #1623: All Valid Triplets That Can Represent a Country
// https://leetcode.com/problems/all-valid-triplets-that-can-represent-a-country/
// Difficulty: Easy [Paid]

import "fmt"

func main() {
	fmt.Println(AllValidTripletsThatCanRepresentACountry())
}

func AllValidTripletsThatCanRepresentACountry() any {
	// TODO: implement
	return nil
}
```

## 1624 — Largest Substring Between Two Equal Characters

```go
package main

// LeetCode #1624: Largest Substring Between Two Equal Characters
// https://leetcode.com/problems/largest-substring-between-two-equal-characters/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1) (since only 26 letters)
func MaxLengthBetweenEqualCharacters(s string) int {
	firstIndex := make(map[rune]int)
	maxLen := -1
	for i, ch := range s {
		if idx, exists := firstIndex[ch]; exists {
			if i-idx-1 > maxLen {
				maxLen = i - idx - 1
			}
		} else {
			firstIndex[ch] = i
		}
	}
	return maxLen
}

func main() {
	fmt.Println(MaxLengthBetweenEqualCharacters("aa"))
	fmt.Println(MaxLengthBetweenEqualCharacters("abca"))
	fmt.Println(MaxLengthBetweenEqualCharacters("cbzxy"))
}
```

## 1629 — Slowest Key

```go
package main

// LeetCode #1629: Slowest Key
// https://leetcode.com/problems/slowest-key/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func SlowestKey(releaseTimes []int, keysPressed string) byte {
	maxDuration := releaseTimes[0]
	result := keysPressed[0]
	for i := 1; i < len(releaseTimes); i++ {
		duration := releaseTimes[i] - releaseTimes[i-1]
		if duration > maxDuration || (duration == maxDuration && keysPressed[i] > result) {
			maxDuration = duration
			result = keysPressed[i]
		}
	}
	return result
}

func main() {
	fmt.Printf("%c\n", SlowestKey([]int{9, 29, 49, 50}, "cbcd"))
	fmt.Printf("%c\n", SlowestKey([]int{12, 23, 36, 46, 62}, "spuda"))
}
```

## 1633 — Percentage Of Users Attended A Contest

```go
package main

// LeetCode #1633: Percentage of Users Attended a Contest
// https://leetcode.com/problems/percentage-of-users-attended-a-contest/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(PercentageOfUsersAttendedAContest())
}

func PercentageOfUsersAttendedAContest() any {
	// TODO: implement
	return nil
}
```

## 1636 — Sort Array By Increasing Frequency

```go
package main

// LeetCode #1636: Sort Array by Increasing Frequency
// https://leetcode.com/problems/sort-array-by-increasing-frequency/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

// Time: O(n log n), Space: O(n)
func FrequencySort(nums []int) []int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}
	sort.Slice(nums, func(i, j int) bool {
		if freq[nums[i]] == freq[nums[j]] {
			return nums[i] > nums[j]
		}
		return freq[nums[i]] < freq[nums[j]]
	})
	return nums
}

func main() {
	fmt.Println(FrequencySort([]int{1, 1, 2, 2, 2, 3}))
	fmt.Println(FrequencySort([]int{2, 3, 1, 3, 2}))
	fmt.Println(FrequencySort([]int{-1, 1, -6, 4, 5, -6, 1, 4, 1}))
}
```

## 1637 — Widest Vertical Area Between Two Points Containing No Points

```go
package main

// LeetCode #1637: Widest Vertical Area Between Two Points Containing No Points
// https://leetcode.com/problems/widest-vertical-area-between-two-points-containing-no-points/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

// Time: O(n log n), Space: O(log n) (for sorting)
func MaxWidthOfVerticalArea(points [][]int) int {
	xs := make([]int, len(points))
	for i, p := range points {
		xs[i] = p[0]
	}
	sort.Ints(xs)
	maxWidth := 0
	for i := 1; i < len(xs); i++ {
		if xs[i]-xs[i-1] > maxWidth {
			maxWidth = xs[i] - xs[i-1]
		}
	}
	return maxWidth
}

func main() {
	fmt.Println(MaxWidthOfVerticalArea([][]int{{8, 7}, {9, 9}, {7, 4}, {9, 7}}))
	fmt.Println(MaxWidthOfVerticalArea([][]int{{3, 1}, {9, 0}, {1, 0}, {1, 4}, {5, 3}, {8, 8}}))
}
```

## 1640 — Check Array Formation Through Concatenation

```go
package main

// LeetCode #1640: Check Array Formation Through Concatenation
// https://leetcode.com/problems/check-array-formation-through-concatenation/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(n)
func CanFormArray(arr []int, pieces [][]int) bool {
	pos := make(map[int]int)
	for i, num := range arr {
		pos[num] = i
	}
	for _, piece := range pieces {
		first := piece[0]
		idx, ok := pos[first]
		if !ok {
			return false
		}
		for j := 1; j < len(piece); j++ {
			if idx+j >= len(arr) || arr[idx+j] != piece[j] {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(CanFormArray([]int{15, 88}, [][]int{{88}, {15}}))
	fmt.Println(CanFormArray([]int{49, 18, 16}, [][]int{{16, 18, 49}}))
	fmt.Println(CanFormArray([]int{91, 4, 64, 78}, [][]int{{78}, {4, 64}, {91}}))
}
```

## 1646 — Get Maximum In Generated Array

```go
package main

// LeetCode #1646: Get Maximum in Generated Array
// https://leetcode.com/problems/get-maximum-in-generated-array/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(n)
func GetMaximumGenerated(n int) int {
	if n == 0 {
		return 0
	}
	nums := make([]int, n+1)
	nums[0] = 0
	nums[1] = 1
	maxVal := 1
	for i := 2; i <= n; i++ {
		if i%2 == 0 {
			nums[i] = nums[i/2]
		} else {
			nums[i] = nums[i/2] + nums[i/2+1]
		}
		if nums[i] > maxVal {
			maxVal = nums[i]
		}
	}
	return maxVal
}

func main() {
	fmt.Println(GetMaximumGenerated(7))
	fmt.Println(GetMaximumGenerated(2))
	fmt.Println(GetMaximumGenerated(3))
}
```

## 1652 — Defuse The Bomb

```go
package main

// LeetCode #1652: Defuse the Bomb
// https://leetcode.com/problems/defuse-the-bomb/
// Difficulty: Easy

import "fmt"

// Time: O(n*|k|), Space: O(n) (or O(1) excluding output)
func Decrypt(code []int, k int) []int {
	n := len(code)
	result := make([]int, n)
	if k == 0 {
		return result
	}
	for i := 0; i < n; i++ {
		sum := 0
		if k > 0 {
			for j := 1; j <= k; j++ {
				sum += code[(i+j)%n]
			}
		} else {
			for j := 1; j <= -k; j++ {
				sum += code[(i-j+n)%n]
			}
		}
		result[i] = sum
	}
	return result
}

func main() {
	fmt.Println(Decrypt([]int{5, 7, 1, 4}, 3))
	fmt.Println(Decrypt([]int{1, 2, 3, 4}, 0))
	fmt.Println(Decrypt([]int{2, 4, 9, 3}, -2))
}
```

## 1656 — Design An Ordered Stream

```go
package main

// LeetCode #1656: Design an Ordered Stream
// https://leetcode.com/problems/design-an-ordered-stream/
// Difficulty: Easy

import "fmt"

type OrderedStream struct {
	stream []string
	ptr    int
}

func Constructor(n int) OrderedStream {
	return OrderedStream{stream: make([]string, n+1), ptr: 1}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	this.stream[idKey] = value
	var result []string
	for this.ptr < len(this.stream) && this.stream[this.ptr] != "" {
		result = append(result, this.stream[this.ptr])
		this.ptr++
	}
	return result
}

func main() {
	os := Constructor(5)
	fmt.Println(os.Insert(3, "ccccc"))
	fmt.Println(os.Insert(1, "aaaaa"))
	fmt.Println(os.Insert(2, "bbbbb"))
	fmt.Println(os.Insert(5, "eeeee"))
	fmt.Println(os.Insert(4, "ddddd"))
}
```

## 1661 — Average Time Of Process Per Machine

```go
package main

// LeetCode #1661: Average Time of Process per Machine
// https://leetcode.com/problems/average-time-of-process-per-machine/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(AverageTimeOfProcessPerMachine())
}

func AverageTimeOfProcessPerMachine() any {
	// TODO: implement
	return nil
}
```

## 1662 — Check If Two String Arrays Are Equivalent

```go
package main

// LeetCode #1662: Check If Two String Arrays Are Equivalent
// https://leetcode.com/problems/check-if-two-string-arrays-are-equivalent/
// Difficulty: Easy

import "fmt"
import "strings"

// Time: O(n), Space: O(n) where n is total characters
func ArrayStringsAreEqual(word1 []string, word2 []string) bool {
	return strings.Join(word1, "") == strings.Join(word2, "")
}

func main() {
	fmt.Println(ArrayStringsAreEqual([]string{"ab", "c"}, []string{"a", "bc"}))
	fmt.Println(ArrayStringsAreEqual([]string{"a", "cb"}, []string{"ab", "c"}))
	fmt.Println(ArrayStringsAreEqual([]string{"abc", "d", "defg"}, []string{"abcddefg"}))
}
```

## 1667 — Fix Names In A Table

```go
package main

// LeetCode #1667: Fix Names in a Table
// https://leetcode.com/problems/fix-names-in-a-table/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FixNamesInATable())
}

func FixNamesInATable() any {
	// TODO: implement
	return nil
}
```

## 1668 — Maximum Repeating Substring

```go
package main

// LeetCode #1668: Maximum Repeating Substring
// https://leetcode.com/problems/maximum-repeating-substring/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

// Time: O(n * m), Space: O(n) where n = len(sequence), m = len(word)
func MaxRepeating(sequence string, word string) int {
	k := 0
	repeated := word
	for strings.Contains(sequence, repeated) {
		k++
		repeated += word
	}
	return k
}

func main() {
	fmt.Println(MaxRepeating("ababc", "ab"))
	fmt.Println(MaxRepeating("ababc", "ba"))
	fmt.Println(MaxRepeating("ababc", "ac"))
}
```

## 1672 — Richest Customer Wealth

```go
package main

// LeetCode #1672: Richest Customer Wealth
// https://leetcode.com/problems/richest-customer-wealth/
// Difficulty: Easy

import "fmt"

// Time: O(m*n), Space: O(1)
func MaximumWealth(accounts [][]int) int {
	maxWealth := 0
	for _, customer := range accounts {
		sum := 0
		for _, amount := range customer {
			sum += amount
		}
		if sum > maxWealth {
			maxWealth = sum
		}
	}
	return maxWealth
}

func main() {
	fmt.Println(MaximumWealth([][]int{{1, 2, 3}, {3, 2, 1}}))
	fmt.Println(MaximumWealth([][]int{{1, 5}, {7, 3}, {3, 5}}))
	fmt.Println(MaximumWealth([][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}))
}
```

## 1677 — Products Worth Over Invoices

```go
package main

// LeetCode #1677: Product's Worth Over Invoices
// https://leetcode.com/problems/products-worth-over-invoices/
// Difficulty: Easy [Paid]

import "fmt"

func main() {
	fmt.Println(ProductsWorthOverInvoices())
}

func ProductsWorthOverInvoices() any {
	// TODO: implement
	return nil
}
```

## 1678 — Goal Parser Interpretation

```go
package main

// LeetCode #1678: Goal Parser Interpretation
// https://leetcode.com/problems/goal-parser-interpretation/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

// Time: O(n), Space: O(n)
func Interpret(command string) string {
	command = strings.ReplaceAll(command, "()", "o")
	command = strings.ReplaceAll(command, "(al)", "al")
	return command
}

func main() {
	fmt.Println(Interpret("G()(al)"))
	fmt.Println(Interpret("G()()()()(al)"))
	fmt.Println(Interpret("(al)G(al)()()G"))
}
```

## 1683 — Invalid Tweets

```go
package main

// LeetCode #1683: Invalid Tweets
// https://leetcode.com/problems/invalid-tweets/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(InvalidTweets())
}

func InvalidTweets() any {
	// TODO: implement
	return nil
}
```

## 1684 — Count The Number Of Consistent Strings

```go
package main

// LeetCode #1684: Count the Number of Consistent Strings
// https://leetcode.com/problems/count-the-number-of-consistent-strings/
// Difficulty: Easy

import "fmt"

// Time: O(n + m*k), Space: O(1)
func CountConsistentStrings(allowed string, words []string) int {
	allowedSet := make(map[byte]bool)
	for i := 0; i < len(allowed); i++ {
		allowedSet[allowed[i]] = true
	}
	count := 0
	for _, word := range words {
		consistent := true
		for i := 0; i < len(word); i++ {
			if !allowedSet[word[i]] {
				consistent = false
				break
			}
		}
		if consistent {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(CountConsistentStrings("ab", []string{"ad", "bd", "aaab", "baa", "badab"}))
	fmt.Println(CountConsistentStrings("abc", []string{"a", "b", "c", "ab", "ac", "bc", "abc"}))
	fmt.Println(CountConsistentStrings("cad", []string{"cc", "acd", "b", "ba", "bac", "bad", "ac", "d"}))
}
```

## 1688 — Count Of Matches In Tournament

```go
package main

// LeetCode #1688: Count of Matches in Tournament
// https://leetcode.com/problems/count-of-matches-in-tournament/
// Difficulty: Easy

import "fmt"

// Time: O(log n), Space: O(1)
func NumberOfMatches(n int) int {
	matches := 0
	for n > 1 {
		if n%2 == 0 {
			matches += n / 2
			n /= 2
		} else {
			matches += (n - 1) / 2
			n = (n-1)/2 + 1
		}
	}
	return matches
}

// Alternative O(1): return n-1 (each match eliminates one team, champion is last)

func main() {
	fmt.Println(NumberOfMatches(7))
	fmt.Println(NumberOfMatches(14))
}
```

## 1693 — Daily Leads And Partners

```go
package main

// LeetCode #1693: Daily Leads and Partners
// https://leetcode.com/problems/daily-leads-and-partners/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(DailyLeadsAndPartners())
}

func DailyLeadsAndPartners() any {
	// TODO: implement
	return nil
}
```

## 1694 — Reformat Phone Number

```go
package main

// LeetCode #1694: Reformat Phone Number
// https://leetcode.com/problems/reformat-phone-number/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(n)
func ReformatNumber(number string) string {
	digits := make([]byte, 0, len(number))
	for i := 0; i < len(number); i++ {
		if number[i] >= '0' && number[i] <= '9' {
			digits = append(digits, number[i])
		}
	}
	var result []byte
	i := 0
	for len(digits)-i > 4 {
		result = append(result, digits[i], digits[i+1], digits[i+2], '-')
		i += 3
	}
	remaining := len(digits) - i
	if remaining == 4 {
		result = append(result, digits[i], digits[i+1], '-', digits[i+2], digits[i+3])
	} else {
		result = append(result, digits[i:]...)
	}
	return string(result)
}

func main() {
	fmt.Println(ReformatNumber("1-23-45 6"))
	fmt.Println(ReformatNumber("123 4-567"))
	fmt.Println(ReformatNumber("123 4-5678"))
}
```

## 1700 — Number Of Students Unable To Eat Lunch

```go
package main

// LeetCode #1700: Number of Students Unable to Eat Lunch
// https://leetcode.com/problems/number-of-students-unable-to-eat-lunch/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func CountStudents(students []int, sandwiches []int) int {
	count := [2]int{0, 0}
	for _, s := range students {
		count[s]++
	}
	for _, sandwich := range sandwiches {
		if count[sandwich] == 0 {
			break
		}
		count[sandwich]--
	}
	return count[0] + count[1]
}

func main() {
	fmt.Println(CountStudents([]int{1, 1, 0, 0}, []int{0, 1, 0, 1}))
	fmt.Println(CountStudents([]int{1, 1, 1, 0, 0, 1}, []int{1, 0, 0, 0, 1, 1}))
}
```

## 1704 — Determine If String Halves Are Alike

```go
package main

// LeetCode #1704: Determine if String Halves Are Alike
// https://leetcode.com/problems/determine-if-string-halves-are-alike/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func HalvesAreAlike(s string) bool {
	vowels := map[byte]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}
	mid := len(s) / 2
	count := 0
	for i := 0; i < mid; i++ {
		if vowels[s[i]] {
			count++
		}
		if vowels[s[i+mid]] {
			count--
		}
	}
	return count == 0
}

func main() {
	fmt.Println(HalvesAreAlike("book"))
	fmt.Println(HalvesAreAlike("textbook"))
}
```

## 1708 — Largest Subarray Length K

```go
package main

// LeetCode #1708: Largest Subarray Length K
// https://leetcode.com/problems/largest-subarray-length-k/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(k) (for output)
func LargestSubarray(nums []int, k int) []int {
	bestIdx := 0
	for i := 1; i <= len(nums)-k; i++ {
		if nums[i] > nums[bestIdx] {
			bestIdx = i
		}
	}
	return nums[bestIdx : bestIdx+k]
}

func main() {
	fmt.Println(LargestSubarray([]int{1, 4, 5, 2, 3}, 3))
	fmt.Println(LargestSubarray([]int{1, 4, 5, 2, 3}, 4))
	fmt.Println(LargestSubarray([]int{1, 2, 3, 4, 5}, 2))
}
```

## 1710 — Maximum Units On A Truck

```go
package main

// LeetCode #1710: Maximum Units on a Truck
// https://leetcode.com/problems/maximum-units-on-a-truck/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

// Time: O(n log n), Space: O(log n) (for sorting)
func MaximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})
	totalUnits := 0
	for _, box := range boxTypes {
		take := box[0]
		if truckSize < take {
			take = truckSize
		}
		totalUnits += take * box[1]
		truckSize -= take
		if truckSize == 0 {
			break
		}
	}
	return totalUnits
}

func main() {
	fmt.Println(MaximumUnits([][]int{{1, 3}, {2, 2}, {3, 1}}, 4))
	fmt.Println(MaximumUnits([][]int{{5, 10}, {2, 5}, {4, 7}, {3, 9}}, 10))
}
```

## 1716 — Calculate Money In Leetcode Bank

```go
package main

// LeetCode #1716: Calculate Money in Leetcode Bank
// https://leetcode.com/problems/calculate-money-in-leetcode-bank/
// Difficulty: Easy

import "fmt"

// Time: O(1), Space: O(1)
func TotalMoney(n int) int {
	weeks := n / 7
	days := n % 7
	// Sum of arithmetic progression: first week = 28, each week adds 7
	total := weeks*28 + 7*weeks*(weeks-1)/2
	// Remaining days
	total += days*(weeks+1) + days*(days-1)/2
	return total
}

func main() {
	fmt.Println(TotalMoney(4))
	fmt.Println(TotalMoney(10))
	fmt.Println(TotalMoney(20))
}
```

## 1720 — Decode Xored Array

```go
package main

// LeetCode #1720: Decode XORed Array
// https://leetcode.com/problems/decode-xored-array/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(n)
func Decode(encoded []int, first int) []int {
	result := make([]int, len(encoded)+1)
	result[0] = first
	for i := 0; i < len(encoded); i++ {
		result[i+1] = result[i] ^ encoded[i]
	}
	return result
}

func main() {
	fmt.Println(Decode([]int{1, 2, 3}, 1))
	fmt.Println(Decode([]int{6, 2, 7, 3}, 4))
}
```

## 1725 — Number Of Rectangles That Can Form The Largest Square

```go
package main

// LeetCode #1725: Number of Rectangles That Can Form the Largest Square
// https://leetcode.com/problems/number-of-rectangles-that-can-form-the-largest-square/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func CountGoodRectangles(rectangles [][]int) int {
	maxLen := 0
	count := 0
	for _, rect := range rectangles {
		side := rect[0]
		if rect[1] < side {
			side = rect[1]
		}
		if side > maxLen {
			maxLen = side
			count = 1
		} else if side == maxLen {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(CountGoodRectangles([][]int{{5, 8}, {3, 9}, {5, 12}, {16, 5}}))
	fmt.Println(CountGoodRectangles([][]int{{2, 3}, {3, 7}, {4, 3}, {3, 7}}))
}
```

## 1729 — Find Followers Count

```go
package main

// LeetCode #1729: Find Followers Count
// https://leetcode.com/problems/find-followers-count/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindFollowersCount())
}

func FindFollowersCount() any {
	// TODO: implement
	return nil
}
```

## 1731 — The Number Of Employees Which Report To Each Employee

```go
package main

// LeetCode #1731: The Number of Employees Which Report to Each Employee
// https://leetcode.com/problems/the-number-of-employees-which-report-to-each-employee/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(TheNumberOfEmployeesWhichReportToEachEmployee())
}

func TheNumberOfEmployeesWhichReportToEachEmployee() any {
	// TODO: implement
	return nil
}
```

## 1732 — Find The Highest Altitude

```go
package main

// LeetCode #1732: Find the Highest Altitude
// https://leetcode.com/problems/find-the-highest-altitude/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func LargestAltitude(gain []int) int {
	maxAlt := 0
	current := 0
	for _, g := range gain {
		current += g
		if current > maxAlt {
			maxAlt = current
		}
	}
	return maxAlt
}

func main() {
	fmt.Println(LargestAltitude([]int{-5, 1, 5, 0, -7}))
	fmt.Println(LargestAltitude([]int{-4, -3, -2, -1, 4, 3, 2}))
}
```

## 1736 — Latest Time By Replacing Hidden Digits

```go
package main

// LeetCode #1736: Latest Time by Replacing Hidden Digits
// https://leetcode.com/problems/latest-time-by-replacing-hidden-digits/
// Difficulty: Easy

import "fmt"

// Time: O(1), Space: O(1)
func MaximumTime(time string) string {
	t := []byte(time)
	if t[0] == '?' {
		if t[1] == '?' || t[1] <= '3' {
			t[0] = '2'
		} else {
			t[0] = '1'
		}
	}
	if t[1] == '?' {
		if t[0] == '2' {
			t[1] = '3'
		} else {
			t[1] = '9'
		}
	}
	if t[3] == '?' {
		t[3] = '5'
	}
	if t[4] == '?' {
		t[4] = '9'
	}
	return string(t)
}

func main() {
	fmt.Println(MaximumTime("2?:?0"))
	fmt.Println(MaximumTime("0?:3?"))
	fmt.Println(MaximumTime("1?:22"))
}
```

## 1741 — Find Total Time Spent By Each Employee

```go
package main

// LeetCode #1741: Find Total Time Spent by Each Employee
// https://leetcode.com/problems/find-total-time-spent-by-each-employee/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindTotalTimeSpentByEachEmployee())
}

func FindTotalTimeSpentByEachEmployee() any {
	// TODO: implement
	return nil
}
```

## 1742 — Maximum Number Of Balls In A Box

```go
package main

// LeetCode #1742: Maximum Number of Balls in a Box
// https://leetcode.com/problems/maximum-number-of-balls-in-a-box/
// Difficulty: Easy

import "fmt"

// Time: O(n log n), Space: O(n) - but n <= 10^5, fine
func CountBalls(lowLimit int, highLimit int) int {
	boxes := make(map[int]int)
	maxBalls := 0
	for i := lowLimit; i <= highLimit; i++ {
		sum := digitSum(i)
		boxes[sum]++
		if boxes[sum] > maxBalls {
			maxBalls = boxes[sum]
		}
	}
	return maxBalls
}

func digitSum(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

func main() {
	fmt.Println(CountBalls(1, 10))
	fmt.Println(CountBalls(5, 15))
	fmt.Println(CountBalls(19, 28))
}
```

## 1748 — Sum Of Unique Elements

```go
package main

// LeetCode #1748: Sum of Unique Elements
// https://leetcode.com/problems/sum-of-unique-elements/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(n)
func SumOfUnique(nums []int) int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}
	sum := 0
	for num, count := range freq {
		if count == 1 {
			sum += num
		}
	}
	return sum
}

func main() {
	fmt.Println(SumOfUnique([]int{1, 2, 3, 2}))
	fmt.Println(SumOfUnique([]int{1, 1, 1, 1, 1}))
	fmt.Println(SumOfUnique([]int{1, 2, 3, 4, 5}))
}
```

## 1752 — Check If Array Is Sorted And Rotated

```go
package main

// LeetCode #1752: Check if Array Is Sorted and Rotated
// https://leetcode.com/problems/check-if-array-is-sorted-and-rotated/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func Check(nums []int) bool {
	drops := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > nums[(i+1)%len(nums)] {
			drops++
		}
	}
	return drops <= 1
}

func main() {
	fmt.Println(Check([]int{3, 4, 5, 1, 2}))
	fmt.Println(Check([]int{2, 1, 3, 4}))
	fmt.Println(Check([]int{1, 2, 3}))
}
```

## 1757 — Recyclable And Low Fat Products

```go
package main

// LeetCode #1757: Recyclable and Low Fat Products
// https://leetcode.com/problems/recyclable-and-low-fat-products/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(RecyclableAndLowFatProducts())
}

func RecyclableAndLowFatProducts() any {
	// TODO: implement
	return nil
}
```

## 1758 — Minimum Changes To Make Alternating Binary String

```go
package main

// LeetCode #1758: Minimum Changes to Make Alternating Binary String
// https://leetcode.com/problems/minimum-changes-to-make-alternating-binary-string/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func MinOperations(s string) int {
	changes := 0
	for i := 0; i < len(s); i++ {
		expected := byte('0' + i%2)
		if s[i] != expected {
			changes++
		}
	}
	if changes < len(s)-changes {
		return changes
	}
	return len(s) - changes
}

func main() {
	fmt.Println(MinOperations("0100"))
	fmt.Println(MinOperations("10"))
	fmt.Println(MinOperations("1111"))
}
```

## 1763 — Longest Nice Substring

```go
package main

// LeetCode #1763: Longest Nice Substring
// https://leetcode.com/problems/longest-nice-substring/
// Difficulty: Easy

import "fmt"
import "unicode"

// Time: O(n^2), Space: O(n)
func LongestNiceSubstring(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		lower := 0
		upper := 0
		for j := i; j < len(s); j++ {
			ch := rune(s[j])
			if unicode.IsUpper(ch) {
				upper |= 1 << (unicode.ToLower(ch) - 'a')
			} else {
				lower |= 1 << (ch - 'a')
			}
			if lower == upper && j-i+1 > len(result) {
				result = s[i : j+1]
			}
		}
	}
	return result
}

func main() {
	fmt.Println(LongestNiceSubstring("YazaAay"))
	fmt.Println(LongestNiceSubstring("Bb"))
	fmt.Println(LongestNiceSubstring("c"))
}
```

## 1768 — Merge Strings Alternately

```go
package main

// LeetCode #1768: Merge Strings Alternately
// https://leetcode.com/problems/merge-strings-alternately/
// Difficulty: Easy

import "fmt"

// Time: O(n+m), Space: O(n+m)
func MergeAlternately(word1 string, word2 string) string {
	result := make([]byte, 0, len(word1)+len(word2))
	i, j := 0, 0
	for i < len(word1) && j < len(word2) {
		result = append(result, word1[i], word2[j])
		i++
		j++
	}
	result = append(result, word1[i:]...)
	result = append(result, word2[j:]...)
	return string(result)
}

func main() {
	fmt.Println(MergeAlternately("abc", "pqr"))
	fmt.Println(MergeAlternately("ab", "pqrs"))
	fmt.Println(MergeAlternately("abcd", "pq"))
}
```

## 1773 — Count Items Matching A Rule

```go
package main

// LeetCode #1773: Count Items Matching a Rule
// https://leetcode.com/problems/count-items-matching-a-rule/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func CountMatches(items [][]string, ruleKey string, ruleValue string) int {
	idx := 0
	switch ruleKey {
	case "color":
		idx = 1
	case "name":
		idx = 2
	}
	count := 0
	for _, item := range items {
		if item[idx] == ruleValue {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(CountMatches([][]string{{"phone", "blue", "pixel"}, {"computer", "silver", "lenovo"}, {"phone", "gold", "iphone"}}, "color", "silver"))
	fmt.Println(CountMatches([][]string{{"phone", "blue", "pixel"}, {"computer", "silver", "phone"}, {"phone", "gold", "iphone"}}, "type", "phone"))
}
```

## 1777 — Products Price For Each Store

```go
package main

// LeetCode #1777: Product's Price for Each Store
// https://leetcode.com/problems/products-price-for-each-store/
// Difficulty: Easy [Paid]

import "fmt"

func main() {
	fmt.Println(ProductsPriceForEachStore())
}

func ProductsPriceForEachStore() any {
	// TODO: implement
	return nil
}
```

## 1779 — Find Nearest Point That Has The Same X Or Y Coordinate

```go
package main

// LeetCode #1779: Find Nearest Point That Has the Same X or Y Coordinate
// https://leetcode.com/problems/find-nearest-point-that-has-the-same-x-or-y-coordinate/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func NearestValidPoint(x int, y int, points [][]int) int {
	minDist := -1
	bestIdx := -1
	for i, p := range points {
		if p[0] == x || p[1] == y {
			dist := abs(p[0]-x) + abs(p[1]-y)
			if minDist == -1 || dist < minDist {
				minDist = dist
				bestIdx = i
			}
		}
	}
	return bestIdx
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	fmt.Println(NearestValidPoint(3, 4, [][]int{{1, 2}, {3, 1}, {2, 4}, {2, 3}, {4, 4}}))
	fmt.Println(NearestValidPoint(3, 4, [][]int{{3, 4}}))
	fmt.Println(NearestValidPoint(3, 4, [][]int{{2, 3}}))
}
```

## 1784 — Check If Binary String Has At Most One Segment Of Ones

```go
package main

// LeetCode #1784: Check if Binary String Has at Most One Segment of Ones
// https://leetcode.com/problems/check-if-binary-string-has-at-most-one-segment-of-ones/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

// Time: O(n), Space: O(1)
func CheckOnesSegment(s string) bool {
	return !strings.Contains(s, "01")
}

func main() {
	fmt.Println(CheckOnesSegment("1001"))
	fmt.Println(CheckOnesSegment("110"))
	fmt.Println(CheckOnesSegment("1"))
}
```

## 1789 — Primary Department For Each Employee

```go
package main

// LeetCode #1789: Primary Department for Each Employee
// https://leetcode.com/problems/primary-department-for-each-employee/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(PrimaryDepartmentForEachEmployee())
}

func PrimaryDepartmentForEachEmployee() any {
	// TODO: implement
	return nil
}
```

## 1790 — Check If One String Swap Can Make Strings Equal

```go
package main

// LeetCode #1790: Check if One String Swap Can Make Strings Equal
// https://leetcode.com/problems/check-if-one-string-swap-can-make-strings-equal/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func AreAlmostEqual(s1 string, s2 string) bool {
	var diff []int
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			diff = append(diff, i)
			if len(diff) > 2 {
				return false
			}
		}
	}
	if len(diff) == 0 {
		return true
	}
	if len(diff) != 2 {
		return false
	}
	return s1[diff[0]] == s2[diff[1]] && s1[diff[1]] == s2[diff[0]]
}

func main() {
	fmt.Println(AreAlmostEqual("bank", "kanb"))
	fmt.Println(AreAlmostEqual("attack", "defend"))
	fmt.Println(AreAlmostEqual("kelb", "kelb"))
}
```

## 1791 — Find Center Of Star Graph

```go
package main

// LeetCode #1791: Find Center of Star Graph
// https://leetcode.com/problems/find-center-of-star-graph/
// Difficulty: Easy

import "fmt"

// Time: O(1), Space: O(1)
func FindCenter(edges [][]int) int {
	if edges[0][0] == edges[1][0] || edges[0][0] == edges[1][1] {
		return edges[0][0]
	}
	return edges[0][1]
}

func main() {
	fmt.Println(FindCenter([][]int{{1, 2}, {2, 3}, {4, 2}}))
	fmt.Println(FindCenter([][]int{{1, 2}, {5, 1}, {1, 3}, {1, 4}}))
}
```

## 1795 — Rearrange Products Table

```go
package main

// LeetCode #1795: Rearrange Products Table
// https://leetcode.com/problems/rearrange-products-table/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(RearrangeProductsTable())
}

func RearrangeProductsTable() any {
	// TODO: implement
	return nil
}
```

## 1796 — Second Largest Digit In A String

```go
package main

// LeetCode #1796: Second Largest Digit in a String
// https://leetcode.com/problems/second-largest-digit-in-a-string/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func SecondHighest(s string) int {
	largest := -1
	second := -1
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			digit := int(s[i] - '0')
			if digit > largest {
				second = largest
				largest = digit
			} else if digit < largest && digit > second {
				second = digit
			}
		}
	}
	return second
}

func main() {
	fmt.Println(SecondHighest("dfa12321afd"))
	fmt.Println(SecondHighest("abc1111"))
	fmt.Println(SecondHighest("ck077"))
}
```

## 1800 — Maximum Ascending Subarray Sum

```go
package main

// LeetCode #1800: Maximum Ascending Subarray Sum
// https://leetcode.com/problems/maximum-ascending-subarray-sum/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func MaxAscendingSum(nums []int) int {
	maxSum, currentSum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			currentSum += nums[i]
		} else {
			currentSum = nums[i]
		}
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}
	return maxSum
}

func main() {
	fmt.Println(MaxAscendingSum([]int{10, 20, 30, 5, 10, 50}))
	fmt.Println(MaxAscendingSum([]int{10, 20, 30, 40, 50}))
	fmt.Println(MaxAscendingSum([]int{12, 17, 15, 13, 10, 11, 12}))
}
```

## 1805 — Number Of Different Integers In A String

```go
package main

// LeetCode #1805: Number of Different Integers in a String
// https://leetcode.com/problems/number-of-different-integers-in-a-string/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(n)
func NumDifferentIntegers(word string) int {
	seen := make(map[string]bool)
	i := 0
	for i < len(word) {
		if word[i] >= '0' && word[i] <= '9' {
			j := i
			for j < len(word) && word[j] >= '0' && word[j] <= '9' {
				j++
			}
			for i < j && word[i] == '0' {
				i++
			}
			num := word[i:j]
			if !seen[num] {
				seen[num] = true
			}
			i = j
		} else {
			i++
		}
	}
	return len(seen)
}

func main() {
	fmt.Println(NumDifferentIntegers("a123bc34d8ef34"))
	fmt.Println(NumDifferentIntegers("leet1234code234"))
	fmt.Println(NumDifferentIntegers("a1b01c001"))
}
```

## 1809 — Ad Free Sessions

```go
package main

// LeetCode #1809: Ad-Free Sessions
// https://leetcode.com/problems/ad-free-sessions/
// Difficulty: Easy [Paid]

import "fmt"

func main() {
	fmt.Println(AdFreeSessions())
}

func AdFreeSessions() any {
	// TODO: implement
	return nil
}
```

## 1812 — Determine Color Of A Chessboard Square

```go
package main

// LeetCode #1812: Determine Color of a Chessboard Square
// https://leetcode.com/problems/determine-color-of-a-chessboard-square/
// Difficulty: Easy

import "fmt"

// Time: O(1), Space: O(1)
func SquareIsWhite(coordinates string) bool {
	// (col + row) % 2 == 0 means dark, == 1 means light (white)
	col := int(coordinates[0] - 'a' + 1)
	row := int(coordinates[1] - '0')
	return (col+row)%2 == 1
}

func main() {
	fmt.Println(SquareIsWhite("a1"))
	fmt.Println(SquareIsWhite("h3"))
	fmt.Println(SquareIsWhite("c7"))
}
```

## 1816 — Truncate Sentence

```go
package main

// LeetCode #1816: Truncate Sentence
// https://leetcode.com/problems/truncate-sentence/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

// Time: O(n), Space: O(n)
func TruncateSentence(s string, k int) string {
	words := strings.Split(s, " ")
	return strings.Join(words[:k], " ")
}

func main() {
	fmt.Println(TruncateSentence("Hello how are you Contestant", 4))
	fmt.Println(TruncateSentence("What is the solution to this problem", 4))
	fmt.Println(TruncateSentence("chopper is not a tanuki", 5))
}
```

## 1821 — Find Customers With Positive Revenue This Year

```go
package main

// LeetCode #1821: Find Customers With Positive Revenue this Year
// https://leetcode.com/problems/find-customers-with-positive-revenue-this-year/
// Difficulty: Easy [Paid]

import "fmt"

func main() {
	fmt.Println(FindCustomersWithPositiveRevenueThisYear())
}

func FindCustomersWithPositiveRevenueThisYear() any {
	// TODO: implement
	return nil
}
```

## 1822 — Sign Of The Product Of An Array

```go
package main

// LeetCode #1822: Sign of the Product of an Array
// https://leetcode.com/problems/sign-of-the-product-of-an-array/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func ArraySign(nums []int) int {
	negCount := 0
	for _, num := range nums {
		if num == 0 {
			return 0
		}
		if num < 0 {
			negCount++
		}
	}
	if negCount%2 == 0 {
		return 1
	}
	return -1
}

func main() {
	fmt.Println(ArraySign([]int{-1, -2, -3, -4, 3, 2, 1}))
	fmt.Println(ArraySign([]int{1, 5, 0, 2, -3}))
	fmt.Println(ArraySign([]int{-1, 1, -1, 1, -1}))
}
```

## 1826 — Faulty Sensor

```go
package main

// LeetCode #1826: Faulty Sensor
// https://leetcode.com/problems/faulty-sensor/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func BadSensor(sensor1 []int, sensor2 []int) int {
	n := len(sensor1)
	i := 0
	for i < n-1 && sensor1[i] == sensor2[i] {
		i++
	}
	if i == n-1 {
		return -1
	}
	// Try sensor1 as faulty (sensor1 drops value at i, rest shifted)
	ok1 := true
	for j := i; j < n-1; j++ {
		if sensor1[j] != sensor2[j+1] {
			ok1 = false
			break
		}
	}
	// Try sensor2 as faulty
	ok2 := true
	for j := i; j < n-1; j++ {
		if sensor2[j] != sensor1[j+1] {
			ok2 = false
			break
		}
	}
	if ok1 && !ok2 {
		return 1
	}
	if !ok1 && ok2 {
		return 2
	}
	return -1
}

func main() {
	fmt.Println(BadSensor([]int{2, 3, 4, 5}, []int{2, 1, 3, 4}))
	fmt.Println(BadSensor([]int{2, 2, 2, 2, 2}, []int{2, 2, 2, 2, 5}))
	fmt.Println(BadSensor([]int{2, 3, 2, 2, 3, 2}, []int{2, 3, 2, 3, 2, 7}))
}
```

## 1827 — Minimum Operations To Make The Array Increasing

```go
package main

// LeetCode #1827: Minimum Operations to Make the Array Increasing
// https://leetcode.com/problems/minimum-operations-to-make-the-array-increasing/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func MinOperations(nums []int) int {
	ops := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			needed := nums[i-1] - nums[i] + 1
			nums[i] += needed
			ops += needed
		}
	}
	return ops
}

func main() {
	fmt.Println(MinOperations([]int{1, 1, 1}))
	fmt.Println(MinOperations([]int{1, 5, 2, 4, 1}))
	fmt.Println(MinOperations([]int{8}))
}
```

## 1832 — Check If The Sentence Is Pangram

```go
package main

// LeetCode #1832: Check if the Sentence Is Pangram
// https://leetcode.com/problems/check-if-the-sentence-is-pangram/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func CheckIfPangram(sentence string) bool {
	seen := 0
	for i := 0; i < len(sentence); i++ {
		seen |= 1 << (sentence[i] - 'a')
	}
	return seen == (1<<26)-1
}

func main() {
	fmt.Println(CheckIfPangram("thequickbrownfoxjumpsoverthelazydog"))
	fmt.Println(CheckIfPangram("leetcode"))
}
```

## 1837 — Sum Of Digits In Base K

```go
package main

// LeetCode #1837: Sum of Digits in Base K
// https://leetcode.com/problems/sum-of-digits-in-base-k/
// Difficulty: Easy

import "fmt"

// Time: O(log_k(n)), Space: O(1)
func SumBase(n int, k int) int {
	sum := 0
	for n > 0 {
		sum += n % k
		n /= k
	}
	return sum
}

func main() {
	fmt.Println(SumBase(34, 6))
	fmt.Println(SumBase(10, 10))
	fmt.Println(SumBase(42, 2))
}
```

## 1844 — Replace All Digits With Characters

```go
package main

// LeetCode #1844: Replace All Digits with Characters
// https://leetcode.com/problems/replace-all-digits-with-characters/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(n)
func ReplaceDigits(s string) string {
	result := []byte(s)
	for i := 1; i < len(s); i += 2 {
		result[i] = s[i-1] + (s[i] - '0')
	}
	return string(result)
}

func main() {
	fmt.Println(ReplaceDigits("a1c1e1"))
	fmt.Println(ReplaceDigits("a1b2c3d4e"))
}
```

## 1848 — Minimum Distance To The Target Element

```go
package main

// LeetCode #1848: Minimum Distance to the Target Element
// https://leetcode.com/problems/minimum-distance-to-the-target-element/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func GetMinDistance(nums []int, target int, start int) int {
	minDist := len(nums)
	for i, num := range nums {
		if num == target {
			dist := start - i
			if dist < 0 {
				dist = -dist
			}
			if dist < minDist {
				minDist = dist
			}
		}
	}
	return minDist
}

func main() {
	fmt.Println(GetMinDistance([]int{1, 2, 3, 4, 5}, 5, 3))
	fmt.Println(GetMinDistance([]int{1}, 1, 0))
	fmt.Println(GetMinDistance([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 1, 0))
}
```

## 1853 — Convert Date Format

```go
package main

// LeetCode #1853: Convert Date Format
// https://leetcode.com/problems/convert-date-format/
// Difficulty: Easy [Paid]

import "fmt"

func main() {
	fmt.Println(ConvertDateFormat())
}

func ConvertDateFormat() any {
	// TODO: implement
	return nil
}
```

## 1854 — Maximum Population Year

```go
package main

// LeetCode #1854: Maximum Population Year
// https://leetcode.com/problems/maximum-population-year/
// Difficulty: Easy

import "fmt"

// Time: O(n + range), Space: O(range)
func MaximumPopulation(logs [][]int) int {
	delta := make([]int, 101) // 1950 to 2050
	for _, log := range logs {
		delta[log[0]-1950]++
		delta[log[1]-1950]--
	}
	maxPop := 0
	currentPop := 0
	bestYear := 1950
	for i := 0; i < 101; i++ {
		currentPop += delta[i]
		if currentPop > maxPop {
			maxPop = currentPop
			bestYear = 1950 + i
		}
	}
	return bestYear
}

func main() {
	fmt.Println(MaximumPopulation([][]int{{1993, 1999}, {2000, 2010}}))
	fmt.Println(MaximumPopulation([][]int{{1950, 1961}, {1960, 1971}, {1970, 1981}}))
}
```

## 1859 — Sorting The Sentence

```go
package main

// LeetCode #1859: Sorting the Sentence
// https://leetcode.com/problems/sorting-the-sentence/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

// Time: O(n), Space: O(n)
func SortSentence(s string) string {
	words := strings.Split(s, " ")
	result := make([]string, len(words))
	for _, w := range words {
		pos := int(w[len(w)-1] - '0') - 1
		result[pos] = w[:len(w)-1]
	}
	return strings.Join(result, " ")
}

func main() {
	fmt.Println(SortSentence("is2 sentence4 This1 a3"))
	fmt.Println(SortSentence("Myself2 Me1 I4 and3"))
}
```

## 1863 — Sum Of All Subset Xor Totals

```go
package main

// LeetCode #1863: Sum of All Subset XOR Totals
// https://leetcode.com/problems/sum-of-all-subset-xor-totals/
// Difficulty: Easy

import "fmt"

// Time: O(2^n), Space: O(n) (recursion stack)
func SubsetXORSum(nums []int) int {
	return dfs(nums, 0, 0)
}

func dfs(nums []int, idx int, currentXor int) int {
	if idx == len(nums) {
		return currentXor
	}
	// Include nums[idx] or skip it
	return dfs(nums, idx+1, currentXor^nums[idx]) + dfs(nums, idx+1, currentXor)
}

func main() {
	fmt.Println(SubsetXORSum([]int{1, 3}))
	fmt.Println(SubsetXORSum([]int{5, 1, 6}))
	fmt.Println(SubsetXORSum([]int{3, 4, 5, 6, 7, 8}))
}
```

## 1869 — Longer Contiguous Segments Of Ones Than Zeros

```go
package main

// LeetCode #1869: Longer Contiguous Segments of Ones Than Zeros
// https://leetcode.com/problems/longer-contiguous-segments-of-ones-than-zeros/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func CheckZeroOnes(s string) bool {
	maxOnes, maxZeros := 0, 0
	curOnes, curZeros := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			curOnes++
			curZeros = 0
			if curOnes > maxOnes {
				maxOnes = curOnes
			}
		} else {
			curZeros++
			curOnes = 0
			if curZeros > maxZeros {
				maxZeros = curZeros
			}
		}
	}
	return maxOnes > maxZeros
}

func main() {
	fmt.Println(CheckZeroOnes("1101"))
	fmt.Println(CheckZeroOnes("111000"))
	fmt.Println(CheckZeroOnes("110100010"))
}
```

## 1873 — Calculate Special Bonus

```go
package main

// LeetCode #1873: Calculate Special Bonus
// https://leetcode.com/problems/calculate-special-bonus/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CalculateSpecialBonus())
}

func CalculateSpecialBonus() any {
	// TODO: implement
	return nil
}
```

## 1876 — Substrings Of Size Three With Distinct Characters

```go
package main

// LeetCode #1876: Substrings of Size Three with Distinct Characters
// https://leetcode.com/problems/substrings-of-size-three-with-distinct-characters/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func CountGoodSubstrings(s string) int {
	count := 0
	for i := 0; i+2 < len(s); i++ {
		if s[i] != s[i+1] && s[i] != s[i+2] && s[i+1] != s[i+2] {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(CountGoodSubstrings("xyzzaz"))
	fmt.Println(CountGoodSubstrings("aababcabc"))
}
```

## 1880 — Check If Word Equals Summation Of Two Words

```go
package main

// LeetCode #1880: Check if Word Equals Summation of Two Words
// https://leetcode.com/problems/check-if-word-equals-summation-of-two-words/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func IsSumEqual(firstWord string, secondWord string, targetWord string) bool {
	return wordValue(firstWord)+wordValue(secondWord) == wordValue(targetWord)
}

func wordValue(s string) int {
	val := 0
	for i := 0; i < len(s); i++ {
		val = val*10 + int(s[i]-'a')
	}
	return val
}

func main() {
	fmt.Println(IsSumEqual("acb", "cba", "cdb"))
	fmt.Println(IsSumEqual("aaa", "a", "aab"))
	fmt.Println(IsSumEqual("aaa", "a", "aaaa"))
}
```

## 1886 — Determine Whether Matrix Can Be Obtained By Rotation

```go
package main

// LeetCode #1886: Determine Whether Matrix Can Be Obtained by Rotation
// https://leetcode.com/problems/determine-whether-matrix-can-be-obtained-by-rotation/
// Difficulty: Easy

import "fmt"

// Time: O(n^2), Space: O(1)
func FindRotation(mat [][]int, target [][]int) bool {
	for rotation := 0; rotation < 4; rotation++ {
		if equal(mat, target) {
			return true
		}
		mat = rotate(mat)
	}
	return false
}

func rotate(mat [][]int) [][]int {
	n := len(mat)
	rotated := make([][]int, n)
	for i := 0; i < n; i++ {
		rotated[i] = make([]int, n)
		for j := 0; j < n; j++ {
			rotated[i][j] = mat[n-1-j][i]
		}
	}
	return rotated
}

func equal(a, b [][]int) bool {
	n := len(a)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(FindRotation([][]int{{0, 1}, {1, 0}}, [][]int{{1, 0}, {0, 1}}))
	fmt.Println(FindRotation([][]int{{0, 1}, {1, 1}}, [][]int{{1, 0}, {0, 1}}))
	fmt.Println(FindRotation([][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}, [][]int{{1, 1, 1}, {0, 1, 0}, {0, 0, 0}}))
}
```

## 1890 — The Latest Login In 2020

```go
package main

// LeetCode #1890: The Latest Login in 2020
// https://leetcode.com/problems/the-latest-login-in-2020/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(TheLatestLoginInTwoZeroTwoZero())
}

func TheLatestLoginInTwoZeroTwoZero() any {
	// TODO: implement
	return nil
}
```

## 1893 — Check If All The Integers In A Range Are Covered

```go
package main

// LeetCode #1893: Check if All the Integers in a Range Are Covered
// https://leetcode.com/problems/check-if-all-the-integers-in-a-range-are-covered/
// Difficulty: Easy

import "fmt"

// Time: O(n * range), Space: O(1)
func IsCovered(ranges [][]int, left int, right int) bool {
	covered := make([]bool, 51)
	for _, r := range ranges {
		for i := r[0]; i <= r[1]; i++ {
			covered[i] = true
		}
	}
	for i := left; i <= right; i++ {
		if !covered[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsCovered([][]int{{1, 2}, {3, 4}, {5, 6}}, 2, 5))
	fmt.Println(IsCovered([][]int{{1, 10}, {10, 20}}, 21, 21))
}
```

## 1897 — Redistribute Characters To Make All Strings Equal

```go
package main

// LeetCode #1897: Redistribute Characters to Make All Strings Equal
// https://leetcode.com/problems/redistribute-characters-to-make-all-strings-equal/
// Difficulty: Easy

import "fmt"

// Time: O(n * len), Space: O(1)
func MakeEqual(words []string) bool {
	freq := make([]int, 26)
	for _, w := range words {
		for i := 0; i < len(w); i++ {
			freq[w[i]-'a']++
		}
	}
	n := len(words)
	for _, count := range freq {
		if count%n != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(MakeEqual([]string{"abc", "aabc", "bc"}))
	fmt.Println(MakeEqual([]string{"ab", "a"}))
}
```

## 1903 — Largest Odd Number In String

```go
package main

// LeetCode #1903: Largest Odd Number in String
// https://leetcode.com/problems/largest-odd-number-in-string/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(LargestOddNumberInString("52"))      // "5"
	fmt.Println(LargestOddNumberInString("4206"))     // ""
	fmt.Println(LargestOddNumberInString("35427"))    // "35427"
}

// Time: O(n), Space: O(1)
func LargestOddNumberInString(num string) string {
	for i := len(num) - 1; i >= 0; i-- {
		if (num[i]-'0')%2 == 1 {
			return num[:i+1]
		}
	}
	return ""
}
```

## 1909 — Remove One Element To Make The Array Strictly Increasing

```go
package main

// LeetCode #1909: Remove One Element to Make the Array Strictly Increasing
// https://leetcode.com/problems/remove-one-element-to-make-the-array-strictly-increasing/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(RemoveOneElementToMakeTheArrayStrictlyIncreasing([]int{1, 2, 10, 5, 7}))       // true
	fmt.Println(RemoveOneElementToMakeTheArrayStrictlyIncreasing([]int{2, 3, 1, 2}))           // false
	fmt.Println(RemoveOneElementToMakeTheArrayStrictlyIncreasing([]int{1, 1, 1}))              // false
}

// Time: O(n), Space: O(1)
func RemoveOneElementToMakeTheArrayStrictlyIncreasing(nums []int) bool {
	removed := false
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			if removed {
				return false
			}
			removed = true
			// Try removing nums[i-1] or nums[i]
			if i-2 < 0 || nums[i] > nums[i-2] {
				// Removing nums[i-1] works
			} else if i+1 >= len(nums) || nums[i+1] > nums[i-1] {
				// Removing nums[i] works, skip it
				nums[i] = nums[i-1] // adjust for next comparison
			} else {
				return false
			}
		}
	}
	return true
}
```

## 1913 — Maximum Product Difference Between Two Pairs

```go
package main

// LeetCode #1913: Maximum Product Difference Between Two Pairs
// https://leetcode.com/problems/maximum-product-difference-between-two-pairs/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MaximumProductDifferenceBetweenTwoPairs([]int{5, 6, 2, 7, 4}))      // 34
	fmt.Println(MaximumProductDifferenceBetweenTwoPairs([]int{4, 2, 5, 9, 7, 4, 8})) // 64
}

// Time: O(n log n), Space: O(1) ignoring sort
func MaximumProductDifferenceBetweenTwoPairs(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	return nums[n-1]*nums[n-2] - nums[0]*nums[1]
}
```

## 1920 — Build Array From Permutation

```go
package main

// LeetCode #1920: Build Array from Permutation
// https://leetcode.com/problems/build-array-from-permutation/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(BuildArrayFromPermutation([]int{0, 2, 1, 5, 3, 4}))    // [0,1,2,4,5,3]
	fmt.Println(BuildArrayFromPermutation([]int{5, 0, 1, 2, 3, 4}))    // [4,5,0,1,2,3]
}

// Time: O(n), Space: O(n)
func BuildArrayFromPermutation(nums []int) []int {
	ans := make([]int, len(nums))
	for i, v := range nums {
		ans[i] = nums[v]
	}
	return ans
}
```

## 1925 — Count Square Sum Triples

```go
package main

// LeetCode #1925: Count Square Sum Triples
// https://leetcode.com/problems/count-square-sum-triples/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountSquareSumTriples(5))  // 2
	fmt.Println(CountSquareSumTriples(10)) // 4
}

// Time: O(n^2), Space: O(1)
func CountSquareSumTriples(n int) int {
	count := 0
	for a := 1; a <= n; a++ {
		for b := 1; b <= n; b++ {
			c2 := a*a + b*b
			c := 1
			for c*c < c2 {
				c++
			}
			if c*c == c2 && c <= n {
				count++
			}
		}
	}
	return count
}
```

## 1929 — Concatenation Of Array

```go
package main

// LeetCode #1929: Concatenation of Array
// https://leetcode.com/problems/concatenation-of-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ConcatenationOfArray([]int{1, 2, 1}))       // [1,2,1,1,2,1]
	fmt.Println(ConcatenationOfArray([]int{1, 3, 2, 1}))    // [1,3,2,1,1,3,2,1]
}

// Time: O(n), Space: O(n)
func ConcatenationOfArray(nums []int) []int {
	n := len(nums)
	ans := make([]int, 2*n)
	for i, v := range nums {
		ans[i] = v
		ans[i+n] = v
	}
	return ans
}
```

## 1933 — Check If String Is Decomposable Into Value Equal Substrings

```go
package main

// LeetCode #1933: Check if String Is Decomposable Into Value-Equal Substrings
// https://leetcode.com/problems/check-if-string-is-decomposable-into-value-equal-substrings/
// Difficulty: Easy [Paid]

import "fmt"

func main() {
	fmt.Println(CheckIfStringIsDecomposableIntoValueEqualSubstrings("000111000"))   // false
	fmt.Println(CheckIfStringIsDecomposableIntoValueEqualSubstrings("00011111222"))  // true
	fmt.Println(CheckIfStringIsDecomposableIntoValueEqualSubstrings("011100022233")) // false
}

// Time: O(n), Space: O(1)
func CheckIfStringIsDecomposableIntoValueEqualSubstrings(s string) bool {
	hasGroupOfTwo := false
	i := 0
	for i < len(s) {
		j := i
		for j < len(s) && s[j] == s[i] {
			j++
		}
		count := j - i
		if count%3 == 1 {
			return false
		}
		if count%3 == 2 {
			if hasGroupOfTwo {
				return false
			}
			hasGroupOfTwo = true
		}
		i = j
	}
	return hasGroupOfTwo
}
```

## 1935 — Maximum Number Of Words You Can Type

```go
package main

// LeetCode #1935: Maximum Number of Words You Can Type
// https://leetcode.com/problems/maximum-number-of-words-you-can-type/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(MaximumNumberOfWordsYouCanType("hello world", "ad"))                   // 1
	fmt.Println(MaximumNumberOfWordsYouCanType("leet code", "e"))                      // 0
	fmt.Println(MaximumNumberOfWordsYouCanType("leet code", "lt"))                     // 1
}

// Time: O(n + m), Space: O(k) where k = len(brokenLetters)
func MaximumNumberOfWordsYouCanType(text string, brokenLetters string) int {
	broken := make(map[byte]bool)
	for i := 0; i < len(brokenLetters); i++ {
		broken[brokenLetters[i]] = true
	}

	words := strings.Fields(text)
	count := 0
	for _, word := range words {
		canType := true
		for i := 0; i < len(word); i++ {
			if broken[word[i]] {
				canType = false
				break
			}
		}
		if canType {
			count++
		}
	}
	return count
}
```

## 1939 — Users That Actively Request Confirmation Messages

```go
package main

// LeetCode #1939: Users That Actively Request Confirmation Messages
// https://leetcode.com/problems/users-that-actively-request-confirmation-messages/
// Difficulty: Easy [Paid] (SQL)

import (
	"fmt"
	"sort"
)

func main() {
	// user actions: each pair is (userID, action)
	actions := [][2]string{{"1", "confirmed"}, {"2", "timeout"}, {"1", "confirmed"}, {"3", "confirmed"}}
	fmt.Println(UsersThatActivelyRequestConfirmationMessages(actions)) // [1 3]
}

// Time: O(n log n), Space: O(n)
func UsersThatActivelyRequestConfirmationMessages(actions [][2]string) []int {
	count := make(map[int]int)
	for _, a := range actions {
		userID := 0
		for _, c := range a[0] {
			userID = userID*10 + int(c-'0')
		}
		if a[1] == "confirmed" {
			count[userID]++
		}
	}

	var result []int
	for uid, c := range count {
		if c >= 2 {
			result = append(result, uid)
		}
	}
	sort.Ints(result)
	return result
}
```

## 1941 — Check If All Characters Have Equal Number Of Occurrences

```go
package main

// LeetCode #1941: Check if All Characters Have Equal Number of Occurrences
// https://leetcode.com/problems/check-if-all-characters-have-equal-number-of-occurrences/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CheckIfAllCharactersHaveEqualNumberOfOccurrences("abacbc")) // true
	fmt.Println(CheckIfAllCharactersHaveEqualNumberOfOccurrences("aaabb"))  // false
}

// Time: O(n), Space: O(1) (max 26 chars)
func CheckIfAllCharactersHaveEqualNumberOfOccurrences(s string) bool {
	freq := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		freq[s[i]]++
	}

	var target int
	for _, v := range freq {
		target = v
		break
	}
	for _, v := range freq {
		if v != target {
			return false
		}
	}
	return true
}
```

## 1945 — Sum Of Digits Of String After Convert

```go
package main

// LeetCode #1945: Sum of Digits of String After Convert
// https://leetcode.com/problems/sum-of-digits-of-string-after-convert/
// Difficulty: Easy

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(SumOfDigitsOfStringAfterConvert("iiii", 1))  // 36
	fmt.Println(SumOfDigitsOfStringAfterConvert("leetcode", 2))  // 6
}

// Time: O(n), Space: O(n)
func SumOfDigitsOfStringAfterConvert(s string, k int) int {
	var digits string
	for i := 0; i < len(s); i++ {
		digits += strconv.Itoa(int(s[i] - 'a' + 1))
	}

	for t := 0; t < k; t++ {
		sum := 0
		for i := 0; i < len(digits); i++ {
			sum += int(digits[i] - '0')
		}
		digits = strconv.Itoa(sum)
	}

	result, _ := strconv.Atoi(digits)
	return result
}
```

## 1952 — Three Divisors

```go
package main

// LeetCode #1952: Three Divisors
// https://leetcode.com/problems/three-divisors/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ThreeDivisors(2))  // false
	fmt.Println(ThreeDivisors(4))  // true
	fmt.Println(ThreeDivisors(81)) // false
}

// Time: O(sqrt(n)), Space: O(1)
func ThreeDivisors(n int) bool {
	// n has exactly 3 divisors iff n is a perfect square of a prime
	if n < 4 {
		return false
	}

	// Check if sqrt(n) is integer
	root := 1
	for root*root < n {
		root++
	}
	if root*root != n {
		return false
	}

	// Check if root is prime
	for i := 2; i*i <= root; i++ {
		if root%i == 0 {
			return false
		}
	}
	return root > 1
}
```

## 1957 — Delete Characters To Make Fancy String

```go
package main

// LeetCode #1957: Delete Characters to Make Fancy String
// https://leetcode.com/problems/delete-characters-to-make-fancy-string/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(DeleteCharactersToMakeFancyString("leeetcode"))     // "leetcode"
	fmt.Println(DeleteCharactersToMakeFancyString("aaabaaaa"))      // "aabaa"
	fmt.Println(DeleteCharactersToMakeFancyString("aab"))           // "aab"
}

// Time: O(n), Space: O(n)
func DeleteCharactersToMakeFancyString(s string) string {
	result := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		n := len(result)
		if n >= 2 && result[n-1] == s[i] && result[n-2] == s[i] {
			continue
		}
		result = append(result, s[i])
	}
	return string(result)
}
```

## 1961 — Check If String Is A Prefix Of Array

```go
package main

// LeetCode #1961: Check If String Is a Prefix of Array
// https://leetcode.com/problems/check-if-string-is-a-prefix-of-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CheckIfStringIsAPrefixOfArray("iloveleetcode", []string{"i", "love", "leetcode", "apples"})) // true
	fmt.Println(CheckIfStringIsAPrefixOfArray("iloveleetcode", []string{"apples", "i", "love", "leetcode"})) // false
}

// Time: O(n), Space: O(1)
func CheckIfStringIsAPrefixOfArray(s string, words []string) bool {
	i := 0
	for _, w := range words {
		if i >= len(s) {
			break
		}
		for j := 0; j < len(w); j++ {
			if i >= len(s) || s[i] != w[j] {
				return false
			}
			i++
		}
	}
	return i == len(s)
}
```

## 1965 — Employees With Missing Information

```go
package main

// LeetCode #1965: Employees With Missing Information
// https://leetcode.com/problems/employees-with-missing-information/
// Difficulty: Easy (SQL)

import (
	"fmt"
	"sort"
)

func main() {
	// Demonstration: employees (ID, name) and salaries (ID, salary)
	// Find employees missing name or salary
	employees := [][2]string{{"1", "Alice"}, {"2", "Bob"}, {"3", "Charlie"}}
	salaries := [][2]string{{"1", "5000"}, {"3", "6000"}}
	fmt.Println(EmployeesWithMissingInformation(employees, salaries)) // [2]
}

// Time: O(n log n), Space: O(n)
func EmployeesWithMissingInformation(employees, salaries [][2]string) []int {
	present := make(map[int]bool)
	for _, e := range employees {
		id := 0
		for _, c := range e[0] {
			id = id*10 + int(c-'0')
		}
		present[id] = true
	}
	for _, s := range salaries {
		id := 0
		for _, c := range s[0] {
			id = id*10 + int(c-'0')
		}
		present[id] = true
	}

	// IDs that appear in only one table
	empSet := make(map[int]bool)
	for _, e := range employees {
		id := 0
		for _, c := range e[0] {
			id = id*10 + int(c-'0')
		}
		empSet[id] = true
	}
	salSet := make(map[int]bool)
	for _, s := range salaries {
		id := 0
		for _, c := range s[0] {
			id = id*10 + int(c-'0')
		}
		salSet[id] = true
	}

	var result []int
	for id := range present {
		if empSet[id] != salSet[id] {
			result = append(result, id)
		}
	}
	sort.Ints(result)
	return result
}
```

## 1967 — Number Of Strings That Appear As Substrings In Word

```go
package main

// LeetCode #1967: Number of Strings That Appear as Substrings in Word
// https://leetcode.com/problems/number-of-strings-that-appear-as-substrings-in-word/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(NumberOfStringsThatAppearAsSubstringsInWord([]string{"a", "abc", "bc", "d"}, "abc")) // 3
	fmt.Println(NumberOfStringsThatAppearAsSubstringsInWord([]string{"a", "b", "c"}, "aaaaabbbbb"))  // 2
}

// Time: O(n * m), Space: O(1)
func NumberOfStringsThatAppearAsSubstringsInWord(patterns []string, word string) int {
	count := 0
	for _, p := range patterns {
		if strings.Contains(word, p) {
			count++
		}
	}
	return count
}
```

## 1971 — Find If Path Exists In Graph

```go
package main

// LeetCode #1971: Find if Path Exists in Graph
// https://leetcode.com/problems/find-if-path-exists-in-graph/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindIfPathExistsInGraph(3, [][]int{{0, 1}, {1, 2}, {2, 0}}, 0, 2)) // true
	fmt.Println(FindIfPathExistsInGraph(6, [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}}, 0, 5)) // false
}

// Time: O(V + E), Space: O(V + E)
func FindIfPathExistsInGraph(n int, edges [][]int, source int, destination int) bool {
	if source == destination {
		return true
	}

	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	visited := make([]bool, n)
	queue := []int{source}
	visited[source] = true

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		for _, v := range adj[u] {
			if v == destination {
				return true
			}
			if !visited[v] {
				visited[v] = true
				queue = append(queue, v)
			}
		}
	}
	return false
}
```

## 1974 — Minimum Time To Type Word Using Special Typewriter

```go
package main

// LeetCode #1974: Minimum Time to Type Word Using Special Typewriter
// https://leetcode.com/problems/minimum-time-to-type-word-using-special-typewriter/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MinimumTimeToTypeWordUsingSpecialTypewriter("abc"))  // 5
	fmt.Println(MinimumTimeToTypeWordUsingSpecialTypewriter("bza"))  // 7
	fmt.Println(MinimumTimeToTypeWordUsingSpecialTypewriter("zjpc")) // 34
}

// Time: O(n), Space: O(1)
func MinimumTimeToTypeWordUsingSpecialTypewriter(word string) int {
	seconds := 0
	pos := 0 // 'a'
	for i := 0; i < len(word); i++ {
		target := int(word[i] - 'a')
		diff := target - pos
		if diff < 0 {
			diff = -diff
		}
		if diff > 13 {
			diff = 26 - diff
		}
		seconds += diff + 1 // move + type
		pos = target
	}
	return seconds
}
```

## 1978 — Employees Whose Manager Left The Company

```go
package main

// LeetCode #1978: Employees Whose Manager Left the Company
// https://leetcode.com/problems/employees-whose-manager-left-the-company/
// Difficulty: Easy (SQL)

import (
	"fmt"
	"sort"
)

func main() {
	// Employees: (employee_id, manager_id, salary)
	employees := [][3]string{{"1", "3", "1000"}, {"2", "3", "2000"}, {"3", "", "3000"}, {"4", "5", "4000"}}
	fmt.Println(EmployeesWhoseManagerLeftTheCompany(employees)) // [4]
}

// Time: O(n log n), Space: O(n)
func EmployeesWhoseManagerLeftTheCompany(employees [][3]string) []int {
	empSet := make(map[int]bool)
	for _, e := range employees {
		id := 0
		for _, c := range e[0] {
			id = id*10 + int(c-'0')
		}
		empSet[id] = true
	}

	var result []int
	for _, e := range employees {
		id := 0
		for _, c := range e[0] {
			id = id*10 + int(c-'0')
		}
		managerID := 0
		if e[1] != "" {
			for _, c := range e[1] {
				managerID = managerID*10 + int(c-'0')
			}
		}
		salary := 0
		for _, c := range e[2] {
			salary = salary*10 + int(c-'0')
		}
		if managerID > 0 && !empSet[managerID] && salary < 30000 {
			result = append(result, id)
		}
	}
	sort.Ints(result)
	return result
}
```

## 1979 — Find Greatest Common Divisor Of Array

```go
package main

// LeetCode #1979: Find Greatest Common Divisor of Array
// https://leetcode.com/problems/find-greatest-common-divisor-of-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindGreatestCommonDivisorOfArray([]int{2, 5, 6, 9, 10})) // 2
	fmt.Println(FindGreatestCommonDivisorOfArray([]int{7, 5, 6, 8, 3}))  // 1
	fmt.Println(FindGreatestCommonDivisorOfArray([]int{3, 3}))            // 3
}

// Time: O(n), Space: O(1)
func FindGreatestCommonDivisorOfArray(nums []int) int {
	min, max := nums[0], nums[0]
	for _, v := range nums {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	a, b := min, max
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
```

