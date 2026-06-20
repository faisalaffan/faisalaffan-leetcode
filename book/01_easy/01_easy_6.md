# Easy (Mudah) — Problem 1984–2432

## 1984 — Minimum Difference Between Highest And Lowest Of K Scores

```go
package main

// LeetCode #1984: Minimum Difference Between Highest and Lowest of K Scores
// https://leetcode.com/problems/minimum-difference-between-highest-and-lowest-of-k-scores/
// Difficulty: Easy

import (
	"fmt"
	"sort"
	"math"
)

func main() {
	fmt.Println(MinimumDifferenceBetweenHighestAndLowestOfKScores([]int{90}, 1))            // 0
	fmt.Println(MinimumDifferenceBetweenHighestAndLowestOfKScores([]int{9, 4, 1, 7}, 2))    // 2
	fmt.Println(MinimumDifferenceBetweenHighestAndLowestOfKScores([]int{9, 4, 1, 7, 5, 3}, 3)) // 3
}

// Time: O(n log n), Space: O(1) ignoring sort
func MinimumDifferenceBetweenHighestAndLowestOfKScores(nums []int, k int) int {
	if k == 1 {
		return 0
	}
	sort.Ints(nums)
	minDiff := math.MaxInt32
	for i := 0; i <= len(nums)-k; i++ {
		diff := nums[i+k-1] - nums[i]
		if diff < minDiff {
			minDiff = diff
		}
	}
	return minDiff
}
```

## 1991 — Find The Middle Index In Array

```go
package main

// LeetCode #1991: Find the Middle Index in Array
// https://leetcode.com/problems/find-the-middle-index-in-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindTheMiddleIndexInArray([]int{2, 3, -1, 8, 4}))   // 3
	fmt.Println(FindTheMiddleIndexInArray([]int{1, -1, 4}))          // 2
	fmt.Println(FindTheMiddleIndexInArray([]int{2, 5}))              // -1
}

// Time: O(n), Space: O(1)
func FindTheMiddleIndexInArray(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}

	leftSum := 0
	for i, v := range nums {
		if leftSum == total-leftSum-v {
			return i
		}
		leftSum += v
	}
	return -1
}
```

## 1995 — Count Special Quadruplets

```go
package main

// LeetCode #1995: Count Special Quadruplets
// https://leetcode.com/problems/count-special-quadruplets/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountSpecialQuadruplets([]int{1, 2, 3, 6}))   // 1
	fmt.Println(CountSpecialQuadruplets([]int{3, 3, 6, 4, 5})) // 0
	fmt.Println(CountSpecialQuadruplets([]int{1, 1, 1, 3, 5})) // 4
}

// Time: O(n^3), Space: O(1)
func CountSpecialQuadruplets(nums []int) int {
	n := len(nums)
	count := 0
	for a := 0; a < n; a++ {
		for b := a + 1; b < n; b++ {
			for c := b + 1; c < n; c++ {
				for d := c + 1; d < n; d++ {
					if nums[a]+nums[b]+nums[c] == nums[d] {
						count++
					}
				}
			}
		}
	}
	return count
}
```

## 2000 — Reverse Prefix Of Word

```go
package main

// LeetCode #2000: Reverse Prefix of Word
// https://leetcode.com/problems/reverse-prefix-of-word/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ReversePrefixOfWord("abcdefd", 'd')) // "dcbaefd"
	fmt.Println(ReversePrefixOfWord("xyxzxe", 'z'))  // "zxyxxe"
	fmt.Println(ReversePrefixOfWord("abcd", 'z'))    // "abcd"
}

// Time: O(n), Space: O(n)
func ReversePrefixOfWord(word string, ch byte) string {
	idx := -1
	for i := 0; i < len(word); i++ {
		if word[i] == ch {
			idx = i
			break
		}
	}
	if idx == -1 {
		return word
	}

	result := make([]byte, len(word))
	for i := 0; i <= idx; i++ {
		result[i] = word[idx-i]
	}
	for i := idx + 1; i < len(word); i++ {
		result[i] = word[i]
	}
	return string(result)
}
```

## 2006 — Count Number Of Pairs With Absolute Difference K

```go
package main

// LeetCode #2006: Count Number of Pairs With Absolute Difference K
// https://leetcode.com/problems/count-number-of-pairs-with-absolute-difference-k/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountNumberOfPairsWithAbsoluteDifferenceK([]int{1, 2, 2, 1}, 1))   // 4
	fmt.Println(CountNumberOfPairsWithAbsoluteDifferenceK([]int{1, 3}, 3))          // 0
	fmt.Println(CountNumberOfPairsWithAbsoluteDifferenceK([]int{3, 2, 1, 5, 4}, 2)) // 3
}

// Time: O(n), Space: O(n)
func CountNumberOfPairsWithAbsoluteDifferenceK(nums []int, k int) int {
	freq := make(map[int]int)
	count := 0
	for _, v := range nums {
		count += freq[v-k] + freq[v+k]
		freq[v]++
	}
	return count
}
```

## 2011 — Final Value Of Variable After Performing Operations

```go
package main

// LeetCode #2011: Final Value of Variable After Performing Operations
// https://leetcode.com/problems/final-value-of-variable-after-performing-operations/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FinalValueOfVariableAfterPerformingOperations([]string{"--X", "X++", "X++"}))     // 1
	fmt.Println(FinalValueOfVariableAfterPerformingOperations([]string{"++X", "++X", "X++"}))     // 3
	fmt.Println(FinalValueOfVariableAfterPerformingOperations([]string{"X++", "++X", "--X", "X--"})) // 0
}

// Time: O(n), Space: O(1)
func FinalValueOfVariableAfterPerformingOperations(operations []string) int {
	x := 0
	for _, op := range operations {
		if op[1] == '+' {
			x++
		} else {
			x--
		}
	}
	return x
}
```

## 2016 — Maximum Difference Between Increasing Elements

```go
package main

// LeetCode #2016: Maximum Difference Between Increasing Elements
// https://leetcode.com/problems/maximum-difference-between-increasing-elements/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MaximumDifferenceBetweenIncreasingElements([]int{7, 1, 5, 4}))   // 4
	fmt.Println(MaximumDifferenceBetweenIncreasingElements([]int{9, 4, 3, 2}))   // -1
	fmt.Println(MaximumDifferenceBetweenIncreasingElements([]int{1, 5, 2, 10}))  // 9
}

// Time: O(n), Space: O(1)
func MaximumDifferenceBetweenIncreasingElements(nums []int) int {
	minSoFar := nums[0]
	maxDiff := -1
	for i := 1; i < len(nums); i++ {
		if nums[i] > minSoFar {
			diff := nums[i] - minSoFar
			if diff > maxDiff {
				maxDiff = diff
			}
		}
		if nums[i] < minSoFar {
			minSoFar = nums[i]
		}
	}
	return maxDiff
}
```

## 2022 — Convert 1D Array Into 2D Array

```go
package main

// LeetCode #2022: Convert 1D Array Into 2D Array
// https://leetcode.com/problems/convert-1d-array-into-2d-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ConvertOneDArrayIntoTwoDArray([]int{1, 2, 3, 4}, 2, 2)) // [[1 2] [3 4]]
	fmt.Println(ConvertOneDArrayIntoTwoDArray([]int{1, 2, 3}, 1, 3))    // [[1 2 3]]
	fmt.Println(ConvertOneDArrayIntoTwoDArray([]int{1, 2}, 1, 1))       // []
}

// Time: O(m*n), Space: O(m*n)
func ConvertOneDArrayIntoTwoDArray(original []int, m int, n int) [][]int {
	if len(original) != m*n {
		return [][]int{}
	}

	result := make([][]int, m)
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
		for j := 0; j < n; j++ {
			result[i][j] = original[i*n+j]
		}
	}
	return result
}
```

## 2026 — Low Quality Problems

```go
package main

// LeetCode #2026: Low-Quality Problems
// https://leetcode.com/problems/low-quality-problems/
// Difficulty: Easy [Paid] (SQL)

import (
	"fmt"
	"sort"
)

func main() {
	// Problems: (problem_id, likes, dislikes)
	problems := [][3]string{{"1", "10", "5"}, {"2", "3", "10"}, {"3", "2", "2"}}
	fmt.Println(LowQualityProblems(problems)) // [2]
}

// Time: O(n log n), Space: O(n)
func LowQualityProblems(problems [][3]string) []int {
	var result []int
	for _, p := range problems {
		id := 0
		for _, c := range p[0] {
			id = id*10 + int(c-'0')
		}
		likes := 0
		for _, c := range p[1] {
			likes = likes*10 + int(c-'0')
		}
		dislikes := 0
		for _, c := range p[2] {
			dislikes = dislikes*10 + int(c-'0')
		}

		// Low-quality: likes / (likes + dislikes) < 0.6
		total := likes + dislikes
		if total > 0 && likes*5 < total*3 {
			result = append(result, id)
		}
	}
	sort.Ints(result)
	return result
}
```

## 2027 — Minimum Moves To Convert String

```go
package main

// LeetCode #2027: Minimum Moves to Convert String
// https://leetcode.com/problems/minimum-moves-to-convert-string/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MinimumMovesToConvertString("XXX"))       // 1
	fmt.Println(MinimumMovesToConvertString("XXOX"))      // 2
	fmt.Println(MinimumMovesToConvertString("OOOO"))      // 0
}

// Time: O(n), Space: O(1)
func MinimumMovesToConvertString(s string) int {
	moves := 0
	i := 0
	for i < len(s) {
		if s[i] == 'X' {
			moves++
			i += 3
		} else {
			i++
		}
	}
	return moves
}
```

## 2032 — Two Out Of Three

```go
package main

// LeetCode #2032: Two Out of Three
// https://leetcode.com/problems/two-out-of-three/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(TwoOutOfThree([]int{1, 1, 3, 2}, []int{2, 3}, []int{3}))    // [3 2]
	fmt.Println(TwoOutOfThree([]int{3, 1}, []int{2, 3}, []int{1, 2}))       // [2 3 1]
	fmt.Println(TwoOutOfThree([]int{1, 2, 2}, []int{4, 3, 3}, []int{5}))    // []
}

// Time: O(n), Space: O(n)
func TwoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	set1 := make(map[int]bool)
	set2 := make(map[int]bool)
	set3 := make(map[int]bool)

	for _, v := range nums1 {
		set1[v] = true
	}
	for _, v := range nums2 {
		set2[v] = true
	}
	for _, v := range nums3 {
		set3[v] = true
	}

	count := make(map[int]int)
	for v := range set1 {
		count[v]++
	}
	for v := range set2 {
		count[v]++
	}
	for v := range set3 {
		count[v]++
	}

	var result []int
	for v, c := range count {
		if c >= 2 {
			result = append(result, v)
		}
	}
	sort.Ints(result)
	return result
}
```

## 2037 — Minimum Number Of Moves To Seat Everyone

```go
package main

// LeetCode #2037: Minimum Number of Moves to Seat Everyone
// https://leetcode.com/problems/minimum-number-of-moves-to-seat-everyone/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MinimumNumberOfMovesToSeatEveryone([]int{3, 1, 5}, []int{2, 7, 4}))   // 4
	fmt.Println(MinimumNumberOfMovesToSeatEveryone([]int{4, 1, 5, 9}, []int{1, 3, 2, 6})) // 7
}

// Time: O(n log n), Space: O(1)
func MinimumNumberOfMovesToSeatEveryone(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)
	moves := 0
	for i := 0; i < len(seats); i++ {
		diff := seats[i] - students[i]
		if diff < 0 {
			diff = -diff
		}
		moves += diff
	}
	return moves
}
```

## 2042 — Check If Numbers Are Ascending In A Sentence

```go
package main

// LeetCode #2042: Check if Numbers Are Ascending in a Sentence
// https://leetcode.com/problems/check-if-numbers-are-ascending-in-a-sentence/
// Difficulty: Easy

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(CheckIfNumbersAreAscendingInASentence("1 box has 3 blue 4 red 6 green and 12 yellow marbles")) // true
	fmt.Println(CheckIfNumbersAreAscendingInASentence("hello world 5 x 5"))                                     // false
	fmt.Println(CheckIfNumbersAreAscendingInASentence("sunset is at 7 11 pm overnight"))                        // false
}

// Time: O(n), Space: O(1)
func CheckIfNumbersAreAscendingInASentence(s string) bool {
	prev := 0
	num := 0
	hasNum := false

	for i := 0; i < len(s); i++ {
		if unicode.IsDigit(rune(s[i])) {
			num = num*10 + int(s[i]-'0')
			hasNum = true
		} else {
			if hasNum {
				if num <= prev {
					return false
				}
				prev = num
				num = 0
				hasNum = false
			}
		}
	}
	if hasNum && num <= prev {
		return false
	}
	return true
}
```

## 2047 — Number Of Valid Words In A Sentence

```go
package main

// LeetCode #2047: Number of Valid Words in a Sentence
// https://leetcode.com/problems/number-of-valid-words-in-a-sentence/
// Difficulty: Easy

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(NumberOfValidWordsInASentence("cat and  dog"))       // 3
	fmt.Println(NumberOfValidWordsInASentence("!this  1-s b8d!"))     // 0
	fmt.Println(NumberOfValidWordsInASentence("alice and  bob are playing stone-game10")) // 5
}

// Time: O(n), Space: O(1)
func NumberOfValidWordsInASentence(sentence string) int {
	count := 0
	i := 0
	for i < len(sentence) {
		// Skip spaces
		if sentence[i] == ' ' {
			i++
			continue
		}

		// Extract word
		start := i
		for i < len(sentence) && sentence[i] != ' ' {
			i++
		}
		word := sentence[start:i]

		if isValidWord(word) {
			count++
		}
	}
	return count
}

func isValidWord(word string) bool {
	hasHyphen := false
	for i, ch := range word {
		if unicode.IsDigit(ch) {
			return false
		}
		if ch == '-' {
			if hasHyphen || i == 0 || i == len(word)-1 {
				return false
			}
			if !unicode.IsLower(rune(word[i-1])) || !unicode.IsLower(rune(word[i+1])) {
				return false
			}
			hasHyphen = true
		}
		if ch == '!' || ch == '.' || ch == ',' {
			if i != len(word)-1 {
				return false
			}
		}
	}
	return true
}
```

## 2053 — Kth Distinct String In An Array

```go
package main

// LeetCode #2053: Kth Distinct String in an Array
// https://leetcode.com/problems/kth-distinct-string-in-an-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(KthDistinctStringInAnArray([]string{"d", "b", "c", "b", "c", "a"}, 2)) // "a"
	fmt.Println(KthDistinctStringInAnArray([]string{"aaa", "aa", "a"}, 1))              // "aaa"
	fmt.Println(KthDistinctStringInAnArray([]string{"a", "b", "a"}, 3))                 // ""
}

// Time: O(n), Space: O(n)
func KthDistinctStringInAnArray(arr []string, k int) string {
	freq := make(map[string]int)
	for _, s := range arr {
		freq[s]++
	}

	idx := 1
	for _, s := range arr {
		if freq[s] == 1 {
			if idx == k {
				return s
			}
			idx++
		}
	}
	return ""
}
```

## 2057 — Smallest Index With Equal Value

```go
package main

// LeetCode #2057: Smallest Index With Equal Value
// https://leetcode.com/problems/smallest-index-with-equal-value/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(SmallestIndexWithEqualValue([]int{0, 1, 2}))          // 0
	fmt.Println(SmallestIndexWithEqualValue([]int{4, 3, 2, 1}))       // 2
	fmt.Println(SmallestIndexWithEqualValue([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})) // -1
}

// Time: O(n), Space: O(1)
func SmallestIndexWithEqualValue(nums []int) int {
	for i, v := range nums {
		if i%10 == v {
			return i
		}
	}
	return -1
}
```

## 2062 — Count Vowel Substrings Of A String

```go
package main

// LeetCode #2062: Count Vowel Substrings of a String
// https://leetcode.com/problems/count-vowel-substrings-of-a-string/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountVowelSubstringsOfAString("aeiouu"))    // 2
	fmt.Println(CountVowelSubstringsOfAString("unicornarihan")) // 0
	fmt.Println(CountVowelSubstringsOfAString("cuaieuouac"))    // 7
}

// Time: O(n^2), Space: O(1)
func CountVowelSubstringsOfAString(word string) int {
	isVowel := func(ch byte) bool {
		return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u'
	}

	count := 0
	for i := 0; i < len(word); i++ {
		vowelSet := make(map[byte]bool)
		for j := i; j < len(word); j++ {
			if !isVowel(word[j]) {
				break
			}
			vowelSet[word[j]] = true
			if len(vowelSet) == 5 {
				count++
			}
		}
	}
	return count
}
```

## 2068 — Check Whether Two Strings Are Almost Equivalent

```go
package main

// LeetCode #2068: Check Whether Two Strings are Almost Equivalent
// https://leetcode.com/problems/check-whether-two-strings-are-almost-equivalent/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CheckWhetherTwoStringsAreAlmostEquivalent("aaaa", "bccb"))   // false
	fmt.Println(CheckWhetherTwoStringsAreAlmostEquivalent("abcdeef", "abaaacc")) // true
	fmt.Println(CheckWhetherTwoStringsAreAlmostEquivalent("cccddabba", "babababab")) // true
}

// Time: O(n), Space: O(1)
func CheckWhetherTwoStringsAreAlmostEquivalent(word1 string, word2 string) bool {
	freq := make([]int, 26)
	for i := 0; i < len(word1); i++ {
		freq[word1[i]-'a']++
	}
	for i := 0; i < len(word2); i++ {
		freq[word2[i]-'a']--
	}
	for _, v := range freq {
		if v < 0 {
			v = -v
		}
		if v > 3 {
			return false
		}
	}
	return true
}
```

## 2072 — The Winner University

```go
package main

// LeetCode #2072: The Winner University
// https://leetcode.com/problems/the-winner-university/
// Difficulty: Easy [Paid] (SQL)

import "fmt"

func main() {
	// New York students: (score)
	newYork := []int{90, 85, 78}
	// California students: (score)
	california := []int{92, 88, 70}
	fmt.Println(TheWinnerUniversity(newYork, california)) // "California University"
}

// Time: O(n), Space: O(1)
func TheWinnerUniversity(newYork []int, california []int) string {
	nyScore := 0
	for _, s := range newYork {
		nyScore += s
	}
	caScore := 0
	for _, s := range california {
		caScore += s
	}

	if nyScore > caScore {
		return "New York University"
	} else if caScore > nyScore {
		return "California University"
	}
	return "No Winner"
}
```

## 2073 — Time Needed To Buy Tickets

```go
package main

// LeetCode #2073: Time Needed to Buy Tickets
// https://leetcode.com/problems/time-needed-to-buy-tickets/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(TimeNeededToBuyTickets([]int{2, 3, 2}, 2)) // 6
	fmt.Println(TimeNeededToBuyTickets([]int{5, 1, 1, 1}, 0)) // 8
}

// Time: O(n), Space: O(1)
func TimeNeededToBuyTickets(tickets []int, k int) int {
	time := 0
	for i, t := range tickets {
		if i <= k {
			if t <= tickets[k] {
				time += t
			} else {
				time += tickets[k]
			}
		} else {
			if t < tickets[k] {
				time += t
			} else {
				time += tickets[k] - 1
			}
		}
	}
	return time
}
```

## 2078 — Two Furthest Houses With Different Colors

```go
package main

// LeetCode #2078: Two Furthest Houses With Different Colors
// https://leetcode.com/problems/two-furthest-houses-with-different-colors/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(TwoFurthestHousesWithDifferentColors([]int{1, 1, 1, 6, 1, 1, 1})) // 3
	fmt.Println(TwoFurthestHousesWithDifferentColors([]int{1, 2, 3, 4, 5}))       // 4
	fmt.Println(TwoFurthestHousesWithDifferentColors([]int{0, 1}))                // 1
}

// Time: O(n), Space: O(1)
func TwoFurthestHousesWithDifferentColors(colors []int) int {
	n := len(colors)
	maxDist := 0

	// Check from leftmost with rightmost
	if colors[0] != colors[n-1] {
		return n - 1
	}

	// If ends are same, find furthest different color from either end
	for i := 1; i < n-1; i++ {
		if colors[i] != colors[0] {
			dist := n - 1 - i
			if dist > maxDist {
				maxDist = dist
			}
			dist = i
			if dist > maxDist {
				maxDist = dist
			}
			break
		}
	}
	return maxDist
}
```

## 2082 — The Number Of Rich Customers

```go
package main

// LeetCode #2082: The Number of Rich Customers
// https://leetcode.com/problems/the-number-of-rich-customers/
// Difficulty: Easy [Paid] (SQL)

import "fmt"

func main() {
	// (customer_id, amount)
	transactions := [][2]string{{"1", "500"}, {"2", "300"}, {"1", "600"}, {"3", "200"}}
	fmt.Println(TheNumberOfRichCustomers(transactions)) // 1 (only customer 1 has amount > 500)
}

// Time: O(n), Space: O(n)
func TheNumberOfRichCustomers(transactions [][2]string) int {
	customers := make(map[int]bool)
	for _, t := range transactions {
		amount := 0
		for _, c := range t[1] {
			amount = amount*10 + int(c-'0')
		}
		if amount > 500 {
			id := 0
			for _, c := range t[0] {
				id = id*10 + int(c-'0')
			}
			customers[id] = true
		}
	}
	return len(customers)
}
```

## 2085 — Count Common Words With One Occurrence

```go
package main

// LeetCode #2085: Count Common Words With One Occurrence
// https://leetcode.com/problems/count-common-words-with-one-occurrence/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountCommonWordsWithOneOccurrence(
		[]string{"leetcode", "is", "amazing", "as", "is"},
		[]string{"amazing", "leetcode", "is"},
	)) // 2
	fmt.Println(CountCommonWordsWithOneOccurrence(
		[]string{"a", "ab"},
		[]string{"a", "a", "a", "ab"},
	)) // 1
}

// Time: O(n + m), Space: O(n + m)
func CountCommonWordsWithOneOccurrence(words1 []string, words2 []string) int {
	freq1 := make(map[string]int)
	freq2 := make(map[string]int)

	for _, w := range words1 {
		freq1[w]++
	}
	for _, w := range words2 {
		freq2[w]++
	}

	count := 0
	for w, c := range freq1 {
		if c == 1 && freq2[w] == 1 {
			count++
		}
	}
	return count
}
```

## 2089 — Find Target Indices After Sorting Array

```go
package main

// LeetCode #2089: Find Target Indices After Sorting Array
// https://leetcode.com/problems/find-target-indices-after-sorting-array/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(FindTargetIndicesAfterSortingArray([]int{1, 2, 5, 2, 3}, 2)) // [1 2]
	fmt.Println(FindTargetIndicesAfterSortingArray([]int{1, 2, 5, 2, 3}, 3)) // [3]
	fmt.Println(FindTargetIndicesAfterSortingArray([]int{1, 2, 5, 2, 3}, 5)) // [4]
}

// Time: O(n log n), Space: O(1) ignoring sort
func FindTargetIndicesAfterSortingArray(nums []int, target int) []int {
	sort.Ints(nums)
	var result []int
	for i, v := range nums {
		if v == target {
			result = append(result, i)
		}
	}
	return result
}
```

## 2094 — Finding 3 Digit Even Numbers

```go
package main

// LeetCode #2094: Finding 3-Digit Even Numbers
// https://leetcode.com/problems/finding-3-digit-even-numbers/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(FindingThreeDigitEvenNumbers([]int{2, 1, 3, 0}))       // [102 120 130 132 210 230 302 310 312 320]
	fmt.Println(FindingThreeDigitEvenNumbers([]int{2, 2, 8, 8, 2}))    // [222 228 282 288 822 828 882]
	fmt.Println(FindingThreeDigitEvenNumbers([]int{0, 0, 0}))          // []
}

// Time: O(n^3), Space: O(1)
func FindingThreeDigitEvenNumbers(digits []int) []int {
	set := make(map[int]bool)
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
					set[num] = true
				}
			}
		}
	}

	result := make([]int, 0, len(set))
	for v := range set {
		result = append(result, v)
	}
	sort.Ints(result)
	return result
}
```

## 2099 — Find Subsequence Of Length K With The Largest Sum

```go
package main

// LeetCode #2099: Find Subsequence of Length K With the Largest Sum
// https://leetcode.com/problems/find-subsequence-of-length-k-with-the-largest-sum/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(FindSubsequenceOfLengthKWithTheLargestSum([]int{2, 1, 3, 3}, 2))       // [3 3]
	fmt.Println(FindSubsequenceOfLengthKWithTheLargestSum([]int{-1, -2, 3, 4}, 3))    // [-1 3 4]
	fmt.Println(FindSubsequenceOfLengthKWithTheLargestSum([]int{3, 4, 3, 3}, 2))      // [3 4]
}

// Time: O(n log n), Space: O(n)
func FindSubsequenceOfLengthKWithTheLargestSum(nums []int, k int) []int {
	type pair struct {
		val int
		idx int
	}

	pairs := make([]pair, len(nums))
	for i, v := range nums {
		pairs[i] = pair{v, i}
	}

	// Sort by value descending
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].val > pairs[j].val
	})

	// Take top k
	selected := pairs[:k]

	// Sort by original index to preserve order
	sort.Slice(selected, func(i, j int) bool {
		return selected[i].idx < selected[j].idx
	})

	result := make([]int, k)
	for i, p := range selected {
		result[i] = p.val
	}
	return result
}
```

## 2103 — Rings And Rods

```go
package main

// LeetCode #2103: Rings and Rods
// https://leetcode.com/problems/rings-and-rods/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(RingsAndRods("B0B6G0R6R0R6G9"))    // 1
	fmt.Println(RingsAndRods("B0R0G0R9R0B0G0"))    // 1
	fmt.Println(RingsAndRods("G4"))                 // 0
}

// Time: O(n), Space: O(1)
func RingsAndRods(rings string) int {
	rods := make([]int, 10)
	for i := 0; i < len(rings); i += 2 {
		color := rings[i]
		rod := rings[i+1] - '0'
		switch color {
		case 'R':
			rods[rod] |= 1
		case 'G':
			rods[rod] |= 2
		case 'B':
			rods[rod] |= 4
		}
	}

	count := 0
	for _, v := range rods {
		if v == 7 { // R|G|B = 1|2|4 = 7
			count++
		}
	}
	return count
}
```

## 2108 — Find First Palindromic String In The Array

```go
package main

// LeetCode #2108: Find First Palindromic String in the Array
// https://leetcode.com/problems/find-first-palindromic-string-in-the-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindFirstPalindromicStringInTheArray([]string{"abc", "car", "ada", "racecar", "cool"})) // "ada"
	fmt.Println(FindFirstPalindromicStringInTheArray([]string{"notapalindrome", "racecar"}))             // "racecar"
	fmt.Println(FindFirstPalindromicStringInTheArray([]string{"def", "ghi"}))                             // ""
}

// Time: O(n * m), Space: O(1)
func FindFirstPalindromicStringInTheArray(words []string) string {
	for _, w := range words {
		if isPalindrome(w) {
			return w
		}
	}
	return ""
}

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
```

## 2114 — Maximum Number Of Words Found In Sentences

```go
package main

// LeetCode #2114: Maximum Number of Words Found in Sentences
// https://leetcode.com/problems/maximum-number-of-words-found-in-sentences/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(MaximumNumberOfWordsFoundInSentences([]string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"})) // 6
	fmt.Println(MaximumNumberOfWordsFoundInSentences([]string{"please wait", "continue to fight", "continue to win"}))                              // 3
}

// Time: O(n * m), Space: O(1)
func MaximumNumberOfWordsFoundInSentences(sentences []string) int {
	maxWords := 0
	for _, s := range sentences {
		count := strings.Count(s, " ") + 1
		if count > maxWords {
			maxWords = count
		}
	}
	return maxWords
}
```

## 2119 — A Number After A Double Reversal

```go
package main

// LeetCode #2119: A Number After a Double Reversal
// https://leetcode.com/problems/a-number-after-a-double-reversal/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ANumberAfterADoubleReversal(526))  // true
	fmt.Println(ANumberAfterADoubleReversal(1800)) // false
	fmt.Println(ANumberAfterADoubleReversal(0))    // true
}

// Time: O(1), Space: O(1)
func ANumberAfterADoubleReversal(num int) bool {
	// Reversing twice yields the same iff num has no trailing zeros
	return num == 0 || num%10 != 0
}
```

## 2124 — Check If All As Appears Before All Bs

```go
package main

// LeetCode #2124: Check if All A's Appears Before All B's
// https://leetcode.com/problems/check-if-all-as-appears-before-all-bs/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CheckIfAllAsAppearsBeforeAllBs("aaabbb")) // true
	fmt.Println(CheckIfAllAsAppearsBeforeAllBs("abab"))   // false
	fmt.Println(CheckIfAllAsAppearsBeforeAllBs("bbb"))    // true
}

// Time: O(n), Space: O(1)
func CheckIfAllAsAppearsBeforeAllBs(s string) bool {
	foundB := false
	for i := 0; i < len(s); i++ {
		if s[i] == 'b' {
			foundB = true
		} else if s[i] == 'a' && foundB {
			return false
		}
	}
	return true
}
```

## 2129 — Capitalize The Title

```go
package main

// LeetCode #2129: Capitalize the Title
// https://leetcode.com/problems/capitalize-the-title/
// Difficulty: Easy

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(CapitalizeTheTitle("capiTalIze tHe titLe")) // "Capitalize The Title"
	fmt.Println(CapitalizeTheTitle("First leTTER of EACH Word")) // "First Letter of Each Word"
	fmt.Println(CapitalizeTheTitle("i lOve leetcode"))           // "i Love Leetcode"
}

// Time: O(n), Space: O(n)
func CapitalizeTheTitle(title string) string {
	words := strings.Fields(title)
	for i, w := range words {
		lower := strings.ToLower(w)
		if len(lower) > 2 {
			runes := []rune(lower)
			runes[0] = unicode.ToUpper(runes[0])
			words[i] = string(runes)
		} else {
			words[i] = lower
		}
	}
	return strings.Join(words, " ")
}
```

## 2133 — Check If Every Row And Column Contains All Numbers

```go
package main

// LeetCode #2133: Check if Every Row and Column Contains All Numbers
// https://leetcode.com/problems/check-if-every-row-and-column-contains-all-numbers/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CheckIfEveryRowAndColumnContainsAllNumbers([][]int{{1, 2, 3}, {3, 1, 2}, {2, 3, 1}})) // true
	fmt.Println(CheckIfEveryRowAndColumnContainsAllNumbers([][]int{{1, 1, 1}, {1, 2, 3}, {1, 2, 3}})) // false
}

// Time: O(n^2), Space: O(n)
func CheckIfEveryRowAndColumnContainsAllNumbers(matrix [][]int) bool {
	n := len(matrix)

	for i := 0; i < n; i++ {
		rowSet := make([]bool, n+1)
		colSet := make([]bool, n+1)
		for j := 0; j < n; j++ {
			if matrix[i][j] < 1 || matrix[i][j] > n || rowSet[matrix[i][j]] {
				return false
			}
			rowSet[matrix[i][j]] = true

			if matrix[j][i] < 1 || matrix[j][i] > n || colSet[matrix[j][i]] {
				return false
			}
			colSet[matrix[j][i]] = true
		}
	}
	return true
}
```

## 2138 — Divide A String Into Groups Of Size K

```go
package main

// LeetCode #2138: Divide a String Into Groups of Size k
// https://leetcode.com/problems/divide-a-string-into-groups-of-size-k/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(DivideAStringIntoGroupsOfSizeK("abcdefghi", 3, 'x')) // ["abc" "def" "ghi"]
	fmt.Println(DivideAStringIntoGroupsOfSizeK("abcdefghij", 3, 'x')) // ["abc" "def" "ghi" "jxx"]
}

// Time: O(n), Space: O(n)
func DivideAStringIntoGroupsOfSizeK(s string, k int, fill byte) []string {
	var result []string
	for i := 0; i < len(s); i += k {
		end := i + k
		if end > len(s) {
			end = len(s)
		}
		group := s[i:end]
		if len(group) < k {
			for len(group) < k {
				group += string(fill)
			}
		}
		result = append(result, group)
	}
	return result
}
```

## 2144 — Minimum Cost Of Buying Candies With Discount

```go
package main

// LeetCode #2144: Minimum Cost of Buying Candies With Discount
// https://leetcode.com/problems/minimum-cost-of-buying-candies-with-discount/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MinimumCostOfBuyingCandiesWithDiscount([]int{1, 2, 3}))    // 5
	fmt.Println(MinimumCostOfBuyingCandiesWithDiscount([]int{6, 5, 7, 9, 2, 2})) // 23
	fmt.Println(MinimumCostOfBuyingCandiesWithDiscount([]int{5, 5}))        // 10
}

// Time: O(n log n), Space: O(1) ignoring sort
func MinimumCostOfBuyingCandiesWithDiscount(cost []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(cost)))
	total := 0
	for i, c := range cost {
		if i%3 != 2 { // buy 2, get 1 free (the cheapest = every 3rd item)
			total += c
		}
	}
	return total
}
```

## 2148 — Count Elements With Strictly Smaller And Greater Elements

```go
package main

// LeetCode #2148: Count Elements With Strictly Smaller and Greater Elements
// https://leetcode.com/problems/count-elements-with-strictly-smaller-and-greater-elements/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountElementsWithStrictlySmallerAndGreaterElements([]int{11, 7, 2, 15}))   // 2
	fmt.Println(CountElementsWithStrictlySmallerAndGreaterElements([]int{-3, 3, 3, 90}))   // 2
	fmt.Println(CountElementsWithStrictlySmallerAndGreaterElements([]int{1, 2, 3}))         // 1
}

// Time: O(n), Space: O(1)
func CountElementsWithStrictlySmallerAndGreaterElements(nums []int) int {
	if len(nums) < 3 {
		return 0
	}

	min, max := nums[0], nums[0]
	for _, v := range nums {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	if min == max {
		return 0
	}

	count := 0
	for _, v := range nums {
		if v > min && v < max {
			count++
		}
	}
	return count
}
```

## 2154 — Keep Multiplying Found Values By Two

```go
package main

// LeetCode #2154: Keep Multiplying Found Values by Two
// https://leetcode.com/problems/keep-multiplying-found-values-by-two/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(KeepMultiplyingFoundValuesByTwo([]int{5, 3, 6, 1, 12}, 3))  // 24
	fmt.Println(KeepMultiplyingFoundValuesByTwo([]int{2, 7, 9}, 4))          // 4
}

// Time: O(n), Space: O(n)
func KeepMultiplyingFoundValuesByTwo(nums []int, original int) int {
	set := make(map[int]bool)
	for _, v := range nums {
		set[v] = true
	}

	for set[original] {
		original *= 2
	}
	return original
}
```

## 2160 — Minimum Sum Of Four Digit Number After Splitting Digits

```go
package main

// LeetCode #2160: Minimum Sum of Four Digit Number After Splitting Digits
// https://leetcode.com/problems/minimum-sum-of-four-digit-number-after-splitting-digits/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MinimumSumOfFourDigitNumberAfterSplittingDigits(2932)) // 52
	fmt.Println(MinimumSumOfFourDigitNumberAfterSplittingDigits(4009)) // 13
}

// Time: O(1), Space: O(1)
func MinimumSumOfFourDigitNumberAfterSplittingDigits(num int) int {
	digits := make([]int, 4)
	for i := 0; i < 4; i++ {
		digits[i] = num % 10
		num /= 10
	}
	sort.Ints(digits)
	// Smallest sum: smallest digit and second smallest as tens, rest as ones
	return (digits[0]*10 + digits[2]) + (digits[1]*10 + digits[3])
}
```

## 2164 — Sort Even And Odd Indices Independently

```go
package main

// LeetCode #2164: Sort Even and Odd Indices Independently
// https://leetcode.com/problems/sort-even-and-odd-indices-independently/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(SortEvenAndOddIndicesIndependently([]int{4, 1, 2, 3})) // [2 3 4 1]
	fmt.Println(SortEvenAndOddIndicesIndependently([]int{2, 1}))       // [2 1]
}

// Time: O(n log n), Space: O(n)
func SortEvenAndOddIndicesIndependently(nums []int) []int {
	n := len(nums)
	even := make([]int, 0, (n+1)/2)
	odd := make([]int, 0, n/2)

	for i, v := range nums {
		if i%2 == 0 {
			even = append(even, v)
		} else {
			odd = append(odd, v)
		}
	}

	sort.Ints(even)
	sort.Sort(sort.Reverse(sort.IntSlice(odd)))

	result := make([]int, n)
	ei, oi := 0, 0
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			result[i] = even[ei]
			ei++
		} else {
			result[i] = odd[oi]
			oi++
		}
	}
	return result
}
```

## 2169 — Count Operations To Obtain Zero

```go
package main

// LeetCode #2169: Count Operations to Obtain Zero
// https://leetcode.com/problems/count-operations-to-obtain-zero/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountOperationsToObtainZero(2, 3))  // 3
	fmt.Println(CountOperationsToObtainZero(10, 10)) // 1
}

// Time: O(log max(num1, num2)), Space: O(1)
func CountOperationsToObtainZero(num1 int, num2 int) int {
	ops := 0
	for num1 > 0 && num2 > 0 {
		if num1 >= num2 {
			num1 -= num2
		} else {
			num2 -= num1
		}
		ops++
	}
	return ops
}
```

## 2176 — Count Equal And Divisible Pairs In An Array

```go
package main

// LeetCode #2176: Count Equal and Divisible Pairs in an Array
// https://leetcode.com/problems/count-equal-and-divisible-pairs-in-an-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountEqualAndDivisiblePairsInAnArray([]int{3, 1, 2, 2, 2, 1, 3}, 2)) // 4
	fmt.Println(CountEqualAndDivisiblePairsInAnArray([]int{1, 2, 3, 4}, 1))           // 0
}

// Time: O(n^2), Space: O(1)
func CountEqualAndDivisiblePairsInAnArray(nums []int, k int) int {
	n := len(nums)
	count := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i] == nums[j] && (i*j)%k == 0 {
				count++
			}
		}
	}
	return count
}
```

## 2180 — Count Integers With Even Digit Sum

```go
package main

// LeetCode #2180: Count Integers With Even Digit Sum
// https://leetcode.com/problems/count-integers-with-even-digit-sum/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountIntegersWithEvenDigitSum(4))  // 2
	fmt.Println(CountIntegersWithEvenDigitSum(30))  // 14
}

// Time: O(n log n), Space: O(1)
func CountIntegersWithEvenDigitSum(num int) int {
	count := 0
	for i := 1; i <= num; i++ {
		if digitSumEven(i) {
			count++
		}
	}
	return count
}

func digitSumEven(n int) bool {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum%2 == 0
}
```

## 2185 — Counting Words With A Given Prefix

```go
package main

// LeetCode #2185: Counting Words With a Given Prefix
// https://leetcode.com/problems/counting-words-with-a-given-prefix/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(CountingWordsWithAGivenPrefix([]string{"pay", "attention", "practice", "attend"}, "at")) // 2
	fmt.Println(CountingWordsWithAGivenPrefix([]string{"leetcode", "win", "loops", "success"}, "code"))   // 0
}

// Time: O(n * m), Space: O(1)
func CountingWordsWithAGivenPrefix(words []string, pref string) int {
	count := 0
	for _, w := range words {
		if strings.HasPrefix(w, pref) {
			count++
		}
	}
	return count
}
```

## 2190 — Most Frequent Number Following Key In An Array

```go
package main

// LeetCode #2190: Most Frequent Number Following Key In an Array
// https://leetcode.com/problems/most-frequent-number-following-key-in-an-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MostFrequentNumberFollowingKeyInAnArray([]int{1, 100, 200, 1, 100}, 1))    // 100
	fmt.Println(MostFrequentNumberFollowingKeyInAnArray([]int{2, 2, 2, 2, 3}, 2))          // 2
}

// Time: O(n), Space: O(n)
func MostFrequentNumberFollowingKeyInAnArray(nums []int, key int) int {
	freq := make(map[int]int)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == key {
			freq[nums[i+1]]++
		}
	}

	maxCount := 0
	result := 0
	for num, count := range freq {
		if count > maxCount {
			maxCount = count
			result = num
		}
	}
	return result
}
```

## 2194 — Cells In A Range On An Excel Sheet

```go
package main

// LeetCode #2194: Cells in a Range on an Excel Sheet
// https://leetcode.com/problems/cells-in-a-range-on-an-excel-sheet/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CellsInARangeOnAnExcelSheet("K1:L2")) // [K1 K2 L1 L2]
	fmt.Println(CellsInARangeOnAnExcelSheet("A1:F1")) // [A1 B1 C1 D1 E1 F1]
}

// Time: O((colDiff+1) * (rowDiff+1)), Space: O((colDiff+1) * (rowDiff+1))
func CellsInARangeOnAnExcelSheet(s string) []string {
	col1 := s[0]
	row1 := s[1]
	col2 := s[3]
	row2 := s[4]

	var result []string
	for c := col1; c <= col2; c++ {
		for r := row1; r <= row2; r++ {
			result = append(result, string(c)+string(r))
		}
	}
	return result
}
```

## 2200 — Find All K Distant Indices In An Array

```go
package main

// LeetCode #2200: Find All K-Distant Indices in an Array
// https://leetcode.com/problems/find-all-k-distant-indices-in-an-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindAllKDistantIndicesInAnArray([]int{3, 4, 9, 1, 3, 9, 5}, 9, 1)) // [1 2 3 4 5 6]
	fmt.Println(FindAllKDistantIndicesInAnArray([]int{2, 2, 2, 2, 2}, 2, 2))       // [0 1 2 3 4]
}

// Time: O(n), Space: O(n)
func FindAllKDistantIndicesInAnArray(nums []int, key int, k int) []int {
	n := len(nums)
	marked := make([]bool, n)

	farthest := -1
	for i, v := range nums {
		if v == key {
			start := i - k
			if start < 0 {
				start = 0
			}
			if start <= farthest {
				start = farthest + 1
			}
			end := i + k
			if end >= n {
				end = n - 1
			}
			for j := start; j <= end; j++ {
				marked[j] = true
			}
			farthest = end
		}
	}

	var result []int
	for i, m := range marked {
		if m {
			result = append(result, i)
		}
	}
	return result
}
```

## 2205 — The Number Of Users That Are Eligible For Discount

```go
package main

// LeetCode #2205: The Number of Users That Are Eligible for Discount
// https://leetcode.com/problems/the-number-of-users-that-are-eligible-for-discount/
// Difficulty: Easy [Paid] (SQL)

import "fmt"

func main() {
	// User purchases: (user_id, amount, date)
	purchases := [][3]string{
		{"1", "120", "2024-01-15"},
		{"2", "50", "2024-02-01"},
		{"1", "80", "2024-03-01"},
		{"3", "200", "2024-01-20"},
	}
	fmt.Println(TheNumberOfUsersThatAreEligibleForDiscount(purchases)) // 2 (users with total > 100)
}

// Time: O(n), Space: O(n)
func TheNumberOfUsersThatAreEligibleForDiscount(purchases [][3]string) int {
	totals := make(map[int]int)
	for _, p := range purchases {
		id := 0
		for _, c := range p[0] {
			id = id*10 + int(c-'0')
		}
		amount := 0
		for _, c := range p[1] {
			amount = amount*10 + int(c-'0')
		}
		totals[id] += amount
	}

	count := 0
	for _, total := range totals {
		if total >= 100 {
			count++
		}
	}
	return count
}
```

## 2206 — Divide Array Into Equal Pairs

```go
package main

// LeetCode #2206: Divide Array Into Equal Pairs
// https://leetcode.com/problems/divide-array-into-equal-pairs/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(DivideArrayIntoEqualPairs([]int{3, 2, 3, 2, 2, 2})) // true
	fmt.Println(DivideArrayIntoEqualPairs([]int{1, 2, 3, 4}))       // false
}

// Time: O(n), Space: O(n)
func DivideArrayIntoEqualPairs(nums []int) bool {
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}
	for _, c := range freq {
		if c%2 != 0 {
			return false
		}
	}
	return true
}
```

## 2210 — Count Hills And Valleys In An Array

```go
package main

// LeetCode #2210: Count Hills and Valleys in an Array
// https://leetcode.com/problems/count-hills-and-valleys-in-an-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountHillsAndValleysInAnArray([]int{2, 4, 1, 1, 6, 5})) // 3
	fmt.Println(CountHillsAndValleysInAnArray([]int{6, 6, 5, 5, 4, 1})) // 0
}

// Time: O(n), Space: O(1)
func CountHillsAndValleysInAnArray(nums []int) int {
	// Build flattened array (remove consecutive duplicates)
	flat := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			flat = append(flat, nums[i])
		}
	}

	count := 0
	for i := 1; i < len(flat)-1; i++ {
		if (flat[i] > flat[i-1] && flat[i] > flat[i+1]) || // hill
			(flat[i] < flat[i-1] && flat[i] < flat[i+1]) { // valley
			count++
		}
	}
	return count
}
```

## 2215 — Find The Difference Of Two Arrays

```go
package main

// LeetCode #2215: Find the Difference of Two Arrays
// https://leetcode.com/problems/find-the-difference-of-two-arrays/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindTheDifferenceOfTwoArrays([]int{1, 2, 3}, []int{2, 4, 6}))       // [[1 3] [4 6]]
	fmt.Println(FindTheDifferenceOfTwoArrays([]int{1, 2, 3, 3}, []int{1, 1, 2, 2})) // [[3] []]
}

// Time: O(n + m), Space: O(n + m)
func FindTheDifferenceOfTwoArrays(nums1 []int, nums2 []int) [][]int {
	set1 := make(map[int]bool)
	set2 := make(map[int]bool)

	for _, v := range nums1 {
		set1[v] = true
	}
	for _, v := range nums2 {
		set2[v] = true
	}

	var diff1 []int
	for v := range set1 {
		if !set2[v] {
			diff1 = append(diff1, v)
		}
	}

	var diff2 []int
	for v := range set2 {
		if !set1[v] {
			diff2 = append(diff2, v)
		}
	}

	return [][]int{diff1, diff2}
}
```

## 2220 — Minimum Bit Flips To Convert Number

```go
package main

// LeetCode #2220: Minimum Bit Flips to Convert Number
// https://leetcode.com/problems/minimum-bit-flips-to-convert-number/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MinimumBitFlipsToConvertNumber(10, 7))  // 3
	fmt.Println(MinimumBitFlipsToConvertNumber(3, 4))   // 3
}

// Time: O(1), Space: O(1)
func MinimumBitFlipsToConvertNumber(start int, goal int) int {
	xor := start ^ goal
	count := 0
	for xor > 0 {
		count += xor & 1
		xor >>= 1
	}
	return count
}
```

## 2224 — Minimum Number Of Operations To Convert Time

```go
package main

// LeetCode #2224: Minimum Number of Operations to Convert Time
// https://leetcode.com/problems/minimum-number-of-operations-to-convert-time/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MinimumNumberOfOperationsToConvertTime("02:30", "04:35")) // 3
	fmt.Println(MinimumNumberOfOperationsToConvertTime("11:00", "11:01")) // 1
}

// Time: O(1), Space: O(1)
func MinimumNumberOfOperationsToConvertTime(current string, correct string) int {
	cur := int(current[0]-'0')*600 + int(current[1]-'0')*60 + int(current[3]-'0')*10 + int(current[4]-'0')
	cor := int(correct[0]-'0')*600 + int(correct[1]-'0')*60 + int(correct[3]-'0')*10 + int(correct[4]-'0')

	diff := cor - cur
	ops := 0

	ops += diff / 60
	diff %= 60
	ops += diff / 15
	diff %= 15
	ops += diff / 5
	diff %= 5
	ops += diff

	return ops
}
```

## 2229 — Check If An Array Is Consecutive

```go
package main

// LeetCode #2229: Check if an Array Is Consecutive
// https://leetcode.com/problems/check-if-an-array-is-consecutive/
// Difficulty: Easy [Paid]

import "fmt"

func main() {
	fmt.Println(CheckIfAnArrayIsConsecutive([]int{1, 3, 4, 2})) // true
	fmt.Println(CheckIfAnArrayIsConsecutive([]int{1, 3, 5}))    // false
	fmt.Println(CheckIfAnArrayIsConsecutive([]int{1, 4}))       // false
}

// Time: O(n), Space: O(n)
func CheckIfAnArrayIsConsecutive(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	set := make(map[int]bool)
	min, max := nums[0], nums[0]

	for _, v := range nums {
		set[v] = true
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	if max-min+1 != len(nums) {
		return false
	}

	for i := min; i <= max; i++ {
		if !set[i] {
			return false
		}
	}
	return true
}
```

## 2230 — The Users That Are Eligible For Discount

```go
package main

// LeetCode #2230: The Users That Are Eligible for Discount
// https://leetcode.com/problems/the-users-that-are-eligible-for-discount/
// Difficulty: Easy [Paid] (SQL)

import "fmt"

func main() {
	// User purchases: (user_id, amount, date)
	purchases := [][3]string{
		{"1", "120", "2024-01-15"},
		{"2", "50", "2024-02-01"},
		{"1", "80", "2024-03-01"},
	}
	fmt.Println(TheUsersThatAreEligibleForDiscount(purchases)) // [1]
}

// Time: O(n), Space: O(n)
func TheUsersThatAreEligibleForDiscount(purchases [][3]string) []int {
	totals := make(map[int]int)
	for _, p := range purchases {
		id := 0
		for _, c := range p[0] {
			id = id*10 + int(c-'0')
		}
		amount := 0
		for _, c := range p[1] {
			amount = amount*10 + int(c-'0')
		}
		totals[id] += amount
	}

	var result []int
	for id, total := range totals {
		if total >= 100 {
			result = append(result, id)
		}
	}
	return result
}
```

## 2231 — Largest Number After Digit Swaps By Parity

```go
package main

// LeetCode #2231: Largest Number After Digit Swaps by Parity
// https://leetcode.com/problems/largest-number-after-digit-swaps-by-parity/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(LargestNumberAfterDigitSwapsByParity(1234)) // 3412
	fmt.Println(LargestNumberAfterDigitSwapsByParity(65875)) // 87655
}

// Time: O(log n * log(log n)), Space: O(log n)
func LargestNumberAfterDigitSwapsByParity(num int) int {
	var digits []int
	for n := num; n > 0; n /= 10 {
		digits = append(digits, n%10)
	}

	// Reverse to get original order
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}

	// Collect even and odd digits
	var evens, odds []int
	for _, d := range digits {
		if d%2 == 0 {
			evens = append(evens, d)
		} else {
			odds = append(odds, d)
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(evens)))
	sort.Sort(sort.Reverse(sort.IntSlice(odds)))

	// Reconstruct
	result := 0
	ei, oi := 0, 0
	for _, d := range digits {
		result *= 10
		if d%2 == 0 {
			result += evens[ei]
			ei++
		} else {
			result += odds[oi]
			oi++
		}
	}
	return result
}
```

## 2235 — Add Two Integers

```go
package main

// LeetCode #2235: Add Two Integers
// https://leetcode.com/problems/add-two-integers/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(AddTwoIntegers(12, 5))   // 17
	fmt.Println(AddTwoIntegers(-10, 4))  // -6
}

// Time: O(1), Space: O(1)
func AddTwoIntegers(num1 int, num2 int) int {
	return num1 + num2
}
```

## 2236 — Root Equals Sum Of Children

```go
package main

// LeetCode #2236: Root Equals Sum of Children
// https://leetcode.com/problems/root-equals-sum-of-children/
// Difficulty: Easy

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{10, &TreeNode{4, nil, nil}, &TreeNode{6, nil, nil}}
	fmt.Println(RootEqualsSumOfChildren(root)) // true

	root2 := &TreeNode{5, &TreeNode{3, nil, nil}, &TreeNode{1, nil, nil}}
	fmt.Println(RootEqualsSumOfChildren(root2)) // false
}

// Time: O(1), Space: O(1)
func RootEqualsSumOfChildren(root *TreeNode) bool {
	return root.Val == root.Left.Val+root.Right.Val
}
```

## 2239 — Find Closest Number To Zero

```go
package main

// LeetCode #2239: Find Closest Number to Zero
// https://leetcode.com/problems/find-closest-number-to-zero/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindClosestNumberToZero([]int{-4, -2, 1, 4, 8})) // 1
	fmt.Println(FindClosestNumberToZero([]int{2, -1, 1}))         // 1
}

// Time: O(n), Space: O(1)
func FindClosestNumberToZero(nums []int) int {
	closest := nums[0]
	for _, v := range nums[1:] {
		absV := v
		if absV < 0 {
			absV = -absV
		}
		absClosest := closest
		if absClosest < 0 {
			absClosest = -absClosest
		}
		if absV < absClosest || (absV == absClosest && v > closest) {
			closest = v
		}
	}
	return closest
}
```

## 2243 — Calculate Digit Sum Of A String

```go
package main

// LeetCode #2243: Calculate Digit Sum of a String
// https://leetcode.com/problems/calculate-digit-sum-of-a-string/
// Difficulty: Easy

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(CalculateDigitSumOfAString("11111222223", 3)) // "135"
	fmt.Println(CalculateDigitSumOfAString("00000000", 3))    // "000"
}

// Time: O(n), Space: O(n)
func CalculateDigitSumOfAString(s string, k int) string {
	for len(s) > k {
		var next string
		for i := 0; i < len(s); i += k {
			end := i + k
			if end > len(s) {
				end = len(s)
			}
			sum := 0
			for j := i; j < end; j++ {
				sum += int(s[j] - '0')
			}
			next += strconv.Itoa(sum)
		}
		s = next
	}
	return s
}
```

## 2248 — Intersection Of Multiple Arrays

```go
package main

// LeetCode #2248: Intersection of Multiple Arrays
// https://leetcode.com/problems/intersection-of-multiple-arrays/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(IntersectionOfMultipleArrays([][]int{{3, 1, 2, 4, 5}, {1, 2, 3, 4}, {3, 4, 5, 6}})) // [3 4]
	fmt.Println(IntersectionOfMultipleArrays([][]int{{1, 2, 3}, {4, 5, 6}}))                        // []
}

// Time: O(n * m), Space: O(n)
func IntersectionOfMultipleArrays(nums [][]int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	freq := make(map[int]int)
	for _, v := range nums[0] {
		freq[v] = 1
	}

	for i := 1; i < len(nums); i++ {
		seen := make(map[int]bool)
		for _, v := range nums[i] {
			if !seen[v] {
				freq[v]++
				seen[v] = true
			}
		}
	}

	var result []int
	for v, c := range freq {
		if c == len(nums) {
			result = append(result, v)
		}
	}
	sort.Ints(result)
	return result
}
```

## 2255 — Count Prefixes Of A Given String

```go
package main

// LeetCode #2255: Count Prefixes of a Given String
// https://leetcode.com/problems/count-prefixes-of-a-given-string/
// Difficulty: Easy
// Time O(n * m) | Space O(1)

import "fmt"

func main() {
	fmt.Println(CountPrefixesOfAGivenString([]string{"a", "b", "c", "ab", "bc", "abc"}, "abc"))           // 3
	fmt.Println(CountPrefixesOfAGivenString([]string{"a", "a"}, "aa"))                                      // 2
	fmt.Println(CountPrefixesOfAGivenString([]string{"feh", "w", "w", "l", "w", "o", "w", "o", "w"}, "w")) // 0
}

func CountPrefixesOfAGivenString(words []string, s string) int {
	count := 0
	for _, w := range words {
		if len(w) <= len(s) && s[:len(w)] == w {
			count++
		}
	}
	return count
}
```

## 2259 — Remove Digit From Number To Maximize Result

```go
package main

// LeetCode #2259: Remove Digit From Number to Maximize Result
// https://leetcode.com/problems/remove-digit-from-number-to-maximize-result/
// Difficulty: Easy
// Time O(n) | Space O(n)

import "fmt"

func main() {
	fmt.Println(RemoveDigitFromNumberToMaximizeResult("123", '3'))  // "12"
	fmt.Println(RemoveDigitFromNumberToMaximizeResult("1231", '1')) // "231"
	fmt.Println(RemoveDigitFromNumberToMaximizeResult("551", '5'))  // "51"
}

func RemoveDigitFromNumberToMaximizeResult(number string, digit byte) string {
	best := ""
	for i := 0; i < len(number); i++ {
		if number[i] == digit {
			candidate := number[:i] + number[i+1:]
			if candidate > best {
				best = candidate
			}
		}
	}
	return best
}
```

## 2264 — Largest 3 Same Digit Number In String

```go
package main

// LeetCode #2264: Largest 3-Same-Digit Number in String
// https://leetcode.com/problems/largest-3-same-digit-number-in-string/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(LargestThreeSameDigitNumberInString("6777133339")) // "777"
	fmt.Println(LargestThreeSameDigitNumberInString("2300011114")) // "000"
	fmt.Println(LargestThreeSameDigitNumberInString("42352338"))   // ""
}

func LargestThreeSameDigitNumberInString(num string) string {
	best := byte(0)
	for i := 2; i < len(num); i++ {
		if num[i] == num[i-1] && num[i] == num[i-2] && num[i] > best {
			best = num[i]
		}
	}
	if best == 0 {
		return ""
	}
	return string([]byte{best, best, best})
}
```

## 2269 — Find The K Beauty Of A Number

```go
package main

// LeetCode #2269: Find the K-Beauty of a Number
// https://leetcode.com/problems/find-the-k-beauty-of-a-number/
// Difficulty: Easy
// Time O(n * k) | Space O(1)

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(FindTheKBeautyOfANumber(240, 2))  // 2
	fmt.Println(FindTheKBeautyOfANumber(430043, 2)) // 2
}

func FindTheKBeautyOfANumber(num int, k int) int {
	s := strconv.Itoa(num)
	count := 0
	for i := 0; i <= len(s)-k; i++ {
		sub, _ := strconv.Atoi(s[i : i+k])
		if sub != 0 && num%sub == 0 {
			count++
		}
	}
	return count
}
```

## 2273 — Find Resultant Array After Removing Anagrams

```go
package main

// LeetCode #2273: Find Resultant Array After Removing Anagrams
// https://leetcode.com/problems/find-resultant-array-after-removing-anagrams/
// Difficulty: Easy
// Time O(n * m log m) | Space O(n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(FindResultantArrayAfterRemovingAnagrams([]string{"abba", "baba", "bbaa", "cd", "cd"})) // ["abba","cd"]
	fmt.Println(FindResultantArrayAfterRemovingAnagrams([]string{"a", "b", "c", "d", "e"}))            // ["a","b","c","d","e"]
}

func FindResultantArrayAfterRemovingAnagrams(words []string) []string {
	result := []string{words[0]}
	prev := sortWord(words[0])

	for i := 1; i < len(words); i++ {
		curr := sortWord(words[i])
		if curr != prev {
			result = append(result, words[i])
			prev = curr
		}
	}
	return result
}

func sortWord(w string) string {
	b := []byte(w)
	sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
	return string(b)
}
```

## 2278 — Percentage Of Letter In String

```go
package main

// LeetCode #2278: Percentage of Letter in String
// https://leetcode.com/problems/percentage-of-letter-in-string/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(PercentageOfLetterInString("foobar", 'o')) // 33
	fmt.Println(PercentageOfLetterInString("jjjj", 'k'))   // 0
	fmt.Println(PercentageOfLetterInString("sgawtb", 's')) // 16
}

func PercentageOfLetterInString(s string, letter byte) int {
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == letter {
			count++
		}
	}
	return count * 100 / len(s)
}
```

## 2283 — Check If Number Has Equal Digit Count And Digit Value

```go
package main

// LeetCode #2283: Check if Number Has Equal Digit Count and Digit Value
// https://leetcode.com/problems/check-if-number-has-equal-digit-count-and-digit-value/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(CheckIfNumberHasEqualDigitCountAndDigitValue("1210")) // true
	fmt.Println(CheckIfNumberHasEqualDigitCountAndDigitValue("030"))  // false
}

func CheckIfNumberHasEqualDigitCountAndDigitValue(num string) bool {
	count := [10]int{}
	for i := 0; i < len(num); i++ {
		count[num[i]-'0']++
	}
	for i := 0; i < len(num); i++ {
		if count[i] != int(num[i]-'0') {
			return false
		}
	}
	return true
}
```

## 2287 — Rearrange Characters To Make Target String

```go
package main

// LeetCode #2287: Rearrange Characters to Make Target String
// https://leetcode.com/problems/rearrange-characters-to-make-target-string/
// Difficulty: Easy
// Time O(n + m) | Space O(1)

import "fmt"

func main() {
	fmt.Println(RearrangeCharactersToMakeTargetString("ilovecodingonleetcode", "code")) // 2
	fmt.Println(RearrangeCharactersToMakeTargetString("abcba", "abc"))                  // 1
}

func RearrangeCharactersToMakeTargetString(s string, target string) int {
	sCount := [26]int{}
	for i := 0; i < len(s); i++ {
		sCount[s[i]-'a']++
	}

	tCount := [26]int{}
	for i := 0; i < len(target); i++ {
		tCount[target[i]-'a']++
	}

	maxCopies := len(s)
	for i := 0; i < 26; i++ {
		if tCount[i] > 0 {
			copies := sCount[i] / tCount[i]
			if copies < maxCopies {
				maxCopies = copies
			}
		}
	}
	return maxCopies
}
```

## 2293 — Min Max Game

```go
package main

// LeetCode #2293: Min Max Game
// https://leetcode.com/problems/min-max-game/
// Difficulty: Easy
// Time O(n) | Space O(n)

import "fmt"

func main() {
	fmt.Println(MinMaxGame([]int{1, 3, 5, 2, 4, 8, 2, 2})) // 1
	fmt.Println(MinMaxGame([]int{3}))                        // 3
}

func MinMaxGame(nums []int) int {
	for len(nums) > 1 {
		next := make([]int, len(nums)/2)
		for i := 0; i < len(next); i++ {
			if i%2 == 0 {
				next[i] = min(nums[2*i], nums[2*i+1])
			} else {
				next[i] = max(nums[2*i], nums[2*i+1])
			}
		}
		nums = next
	}
	return nums[0]
}
```

## 2299 — Strong Password Checker Ii

```go
package main

// LeetCode #2299: Strong Password Checker II
// https://leetcode.com/problems/strong-password-checker-ii/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(StrongPasswordCheckerIi("IloveLe3tcode!")) // true
	fmt.Println(StrongPasswordCheckerIi("Me+You--IsMyDream")) // false
	fmt.Println(StrongPasswordCheckerIi("1aB!")) // false
}

func StrongPasswordCheckerIi(password string) bool {
	if len(password) < 8 {
		return false
	}

	hasLower, hasUpper, hasDigit, hasSpecial := false, false, false, false
	special := "!@#$%^&*()-+"

	for i := 0; i < len(password); i++ {
		if i > 0 && password[i] == password[i-1] {
			return false
		}
		ch := password[i]
		switch {
		case ch >= 'a' && ch <= 'z':
			hasLower = true
		case ch >= 'A' && ch <= 'Z':
			hasUpper = true
		case ch >= '0' && ch <= '9':
			hasDigit = true
		default:
			for j := 0; j < len(special); j++ {
				if ch == special[j] {
					hasSpecial = true
					break
				}
			}
		}
	}
	return hasLower && hasUpper && hasDigit && hasSpecial
}
```

## 2303 — Calculate Amount Paid In Taxes

```go
package main

// LeetCode #2303: Calculate Amount Paid in Taxes
// https://leetcode.com/problems/calculate-amount-paid-in-taxes/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(CalculateAmountPaidInTaxes([][]int{{3, 50}, {7, 10}, {12, 25}}, 10)) // 2.65
	fmt.Println(CalculateAmountPaidInTaxes([][]int{{1, 0}, {4, 25}, {5, 50}}, 2))    // 0.25
}

func CalculateAmountPaidInTaxes(brackets [][]int, income int) float64 {
	tax := 0.0
	prev := 0
	for _, b := range brackets {
		upper := b[0]
		percent := float64(b[1]) / 100.0
		taxable := min(upper, income) - prev
		if taxable > 0 {
			tax += float64(taxable) * percent
		}
		prev = upper
		if income <= upper {
			break
		}
	}
	return tax
}
```

## 2309 — Greatest English Letter In Upper And Lower Case

```go
package main

// LeetCode #2309: Greatest English Letter in Upper and Lower Case
// https://leetcode.com/problems/greatest-english-letter-in-upper-and-lower-case/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(GreatestEnglishLetterInUpperAndLowerCase("lEeTcOdE")) // "E"
	fmt.Println(GreatestEnglishLetterInUpperAndLowerCase("arRAzFif")) // "R"
	fmt.Println(GreatestEnglishLetterInUpperAndLowerCase("AbCdEfGhIjK")) // ""
}

func GreatestEnglishLetterInUpperAndLowerCase(s string) string {
	seen := [26]bool{}
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			seen[s[i]-'a'] = true
		}
	}
	best := byte(0)
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			idx := s[i] - 'A'
			if seen[idx] && s[i] > best {
				best = s[i]
			}
		}
	}
	if best == 0 {
		return ""
	}
	return string(best)
}
```

## 2315 — Count Asterisks

```go
package main

// LeetCode #2315: Count Asterisks
// https://leetcode.com/problems/count-asterisks/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(CountAsterisks("l|*e*et|c**o|*de|")) // 2
	fmt.Println(CountAsterisks("iamprogrammer"))      // 0
}

func CountAsterisks(s string) int {
	count := 0
	bar := false
	for i := 0; i < len(s); i++ {
		if s[i] == '|' {
			bar = !bar
		} else if s[i] == '*' && !bar {
			count++
		}
	}
	return count
}
```

## 2319 — Check If Matrix Is X Matrix

```go
package main

// LeetCode #2319: Check if Matrix Is X-Matrix
// https://leetcode.com/problems/check-if-matrix-is-x-matrix/
// Difficulty: Easy
// Time O(n^2) | Space O(1)

import "fmt"

func main() {
	fmt.Println(CheckIfMatrixIsXMatrix([][]int{{2, 0, 0, 1}, {0, 3, 1, 0}, {0, 5, 2, 0}, {4, 0, 0, 2}})) // true
	fmt.Println(CheckIfMatrixIsXMatrix([][]int{{5, 7, 0}, {0, 3, 1}, {0, 5, 0}}))                        // false
}

func CheckIfMatrixIsXMatrix(grid [][]int) bool {
	n := len(grid)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j || i+j == n-1 {
				if grid[i][j] == 0 {
					return false
				}
			} else if grid[i][j] != 0 {
				return false
			}
		}
	}
	return true
}
```

## 2325 — Decode The Message

```go
package main

// LeetCode #2325: Decode the Message
// https://leetcode.com/problems/decode-the-message/
// Difficulty: Easy
// Time O(n + m) | Space O(1)

import "fmt"

func main() {
	fmt.Println(DecodeTheMessage("the quick brown fox jumps over the lazy dog", "vkbs bs t suepuv")) // "this is a secret"
	fmt.Println(DecodeTheMessage("eljuxhpwnyrdgtqkviszcfmabo", "zwx hnfx lqantp mnoeius ycgk vcnjrdb")) // "the five boxing wizards jump quickly"
}

func DecodeTheMessage(key string, message string) string {
	mapping := make([]byte, 26)
	idx := byte(0)

	for i := 0; i < len(key); i++ {
		if key[i] != ' ' && mapping[key[i]-'a'] == 0 {
			mapping[key[i]-'a'] = 'a' + idx
			idx++
		}
	}

	res := make([]byte, len(message))
	for i := 0; i < len(message); i++ {
		if message[i] == ' ' {
			res[i] = ' '
		} else {
			res[i] = mapping[message[i]-'a']
		}
	}
	return string(res)
}
```

## 2329 — Product Sales Analysis V

```go
package main

// LeetCode #2329: Product Sales Analysis V
// https://leetcode.com/problems/product-sales-analysis-v/
// Difficulty: Easy [Paid]

import "fmt"

func main() {
	fmt.Println(ProductSalesAnalysisV())
}

func ProductSalesAnalysisV() any {
	// TODO: implement
	return nil
}
```

## 2331 — Evaluate Boolean Binary Tree

```go
package main

// LeetCode #2331: Evaluate Boolean Binary Tree
// https://leetcode.com/problems/evaluate-boolean-binary-tree/
// Difficulty: Easy
// Time O(n) | Space O(h)

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// full binary tree: OR(2, AND(3,1)) = true
	root1 := &TreeNode{2,
		&TreeNode{1, nil, nil},
		&TreeNode{3, &TreeNode{0, nil, nil}, &TreeNode{1, nil, nil}},
	}
	fmt.Println(EvaluateBooleanBinaryTree(root1)) // true

	root2 := &TreeNode{0, nil, nil}
	fmt.Println(EvaluateBooleanBinaryTree(root2)) // false
}

func EvaluateBooleanBinaryTree(root *TreeNode) bool {
	switch root.Val {
	case 0:
		return false
	case 1:
		return true
	case 2:
		return EvaluateBooleanBinaryTree(root.Left) || EvaluateBooleanBinaryTree(root.Right)
	default: // 3
		return EvaluateBooleanBinaryTree(root.Left) && EvaluateBooleanBinaryTree(root.Right)
	}
}
```

## 2335 — Minimum Amount Of Time To Fill Cups

```go
package main

// LeetCode #2335: Minimum Amount of Time to Fill Cups
// https://leetcode.com/problems/minimum-amount-of-time-to-fill-cups/
// Difficulty: Easy
// Time O(1) | Space O(1)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MinimumAmountOfTimeToFillCups([]int{1, 4, 2})) // 4
	fmt.Println(MinimumAmountOfTimeToFillCups([]int{5, 4, 4})) // 5
}

func MinimumAmountOfTimeToFillCups(amount []int) int {
	sort.Ints(amount)
	a, b, c := amount[0], amount[1], amount[2]
	if a+b <= c {
		return c
	}
	return (a+b+c+1)/2
}
```

## 2339 — All The Matches Of The League

```go
package main

// LeetCode #2339: All the Matches of the League
// https://leetcode.com/problems/all-the-matches-of-the-league/
// Difficulty: Easy [Paid]

import "fmt"

func main() {
	fmt.Println(AllTheMatchesOfTheLeague())
}

func AllTheMatchesOfTheLeague() any {
	// TODO: implement
	return nil
}
```

## 2341 — Maximum Number Of Pairs In Array

```go
package main

// LeetCode #2341: Maximum Number of Pairs in Array
// https://leetcode.com/problems/maximum-number-of-pairs-in-array/
// Difficulty: Easy
// Time O(n) | Space O(n)

import "fmt"

func main() {
	fmt.Println(MaximumNumberOfPairsInArray([]int{1, 3, 2, 1, 3, 2, 2})) // [3,1]
	fmt.Println(MaximumNumberOfPairsInArray([]int{1, 1}))                  // [1,0]
}

func MaximumNumberOfPairsInArray(nums []int) []int {
	freq := map[int]int{}
	for _, n := range nums {
		freq[n]++
	}
	pairs, leftovers := 0, 0
	for _, c := range freq {
		pairs += c / 2
		leftovers += c % 2
	}
	return []int{pairs, leftovers}
}
```

## 2347 — Best Poker Hand

```go
package main

// LeetCode #2347: Best Poker Hand
// https://leetcode.com/problems/best-poker-hand/
// Difficulty: Easy
// Time O(1) | Space O(1)

import "fmt"

func main() {
	fmt.Println(BestPokerHand([]int{13, 2, 3, 1, 9}, []byte{'a', 'a', 'a', 'a', 'a'})) // "Flush"
	fmt.Println(BestPokerHand([]int{4, 4, 2, 4, 4}, []byte{'d', 'a', 'a', 'b', 'c'})) // "Three of a Kind"
	fmt.Println(BestPokerHand([]int{10, 10, 2, 12, 9}, []byte{'a', 'b', 'c', 'a', 'd'})) // "Pair"
}

func BestPokerHand(ranks []int, suits []byte) string {
	// Check flush
	if suits[0] == suits[1] && suits[1] == suits[2] && suits[2] == suits[3] && suits[3] == suits[4] {
		return "Flush"
	}

	// Check three of a kind or pair
	rankCount := [14]int{}
	for _, r := range ranks {
		rankCount[r]++
		if rankCount[r] == 3 {
			return "Three of a Kind"
		}
	}
	for _, c := range rankCount {
		if c == 2 {
			return "Pair"
		}
	}
	return "High Card"
}
```

## 2351 — First Letter To Appear Twice

```go
package main

// LeetCode #2351: First Letter to Appear Twice
// https://leetcode.com/problems/first-letter-to-appear-twice/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(string(FirstLetterToAppearTwice("abccbaacz"))) // "c"
	fmt.Println(string(FirstLetterToAppearTwice("abcdd")))      // "d"
}

func FirstLetterToAppearTwice(s string) byte {
	seen := [26]bool{}
	for i := 0; i < len(s); i++ {
		idx := s[i] - 'a'
		if seen[idx] {
			return s[i]
		}
		seen[idx] = true
	}
	return 0
}
```

## 2356 — Number Of Unique Subjects Taught By Each Teacher

```go
package main

// LeetCode #2356: Number of Unique Subjects Taught by Each Teacher
// https://leetcode.com/problems/number-of-unique-subjects-taught-by-each-teacher/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(NumberOfUniqueSubjectsTaughtByEachTeacher())
}

func NumberOfUniqueSubjectsTaughtByEachTeacher() any {
	// TODO: implement
	return nil
}
```

## 2357 — Make Array Zero By Subtracting Equal Amounts

```go
package main

// LeetCode #2357: Make Array Zero by Subtracting Equal Amounts
// https://leetcode.com/problems/make-array-zero-by-subtracting-equal-amounts/
// Difficulty: Easy
// Time O(n log n) | Space O(1)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MakeArrayZeroBySubtractingEqualAmounts([]int{1, 5, 0, 3, 5})) // 3
	fmt.Println(MakeArrayZeroBySubtractingEqualAmounts([]int{0}))              // 0
}

func MakeArrayZeroBySubtractingEqualAmounts(nums []int) int {
	sort.Ints(nums)
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			count++
			sub := nums[i]
			for j := i; j < len(nums); j++ {
				nums[j] -= sub
			}
		}
	}
	return count
}
```

## 2363 — Merge Similar Items

```go
package main

// LeetCode #2363: Merge Similar Items
// https://leetcode.com/problems/merge-similar-items/
// Difficulty: Easy
// Time O((n+m) log(n+m)) | Space O(n+m)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MergeSimilarItems([][]int{{1, 1}, {4, 5}, {3, 8}}, [][]int{{3, 1}, {1, 5}})) // [[1,6],[3,9],[4,5]]
	fmt.Println(MergeSimilarItems([][]int{{1, 1}, {3, 2}, {2, 3}}, [][]int{{2, 1}, {3, 2}, {1, 3}})) // [[1,4],[2,4],[3,4]]
}

func MergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	valueMap := map[int]int{}
	for _, item := range items1 {
		valueMap[item[0]] += item[1]
	}
	for _, item := range items2 {
		valueMap[item[0]] += item[1]
	}

	result := make([][]int, 0, len(valueMap))
	for v, w := range valueMap {
		result = append(result, []int{v, w})
	}
	sort.Slice(result, func(i, j int) bool { return result[i][0] < result[j][0] })
	return result
}
```

## 2367 — Number Of Arithmetic Triplets

```go
package main

// LeetCode #2367: Number of Arithmetic Triplets
// https://leetcode.com/problems/number-of-arithmetic-triplets/
// Difficulty: Easy
// Time O(n) | Space O(n)

import "fmt"

func main() {
	fmt.Println(NumberOfArithmeticTriplets([]int{0, 1, 4, 6, 7, 10}, 3)) // 2
	fmt.Println(NumberOfArithmeticTriplets([]int{4, 5, 6, 7, 8, 9}, 2))  // 2
}

func NumberOfArithmeticTriplets(nums []int, diff int) int {
	seen := make(map[int]bool, len(nums))
	for _, n := range nums {
		seen[n] = true
	}
	count := 0
	for _, n := range nums {
		if seen[n+diff] && seen[n+2*diff] {
			count++
		}
	}
	return count
}
```

## 2373 — Largest Local Values In A Matrix

```go
package main

// LeetCode #2373: Largest Local Values in a Matrix
// https://leetcode.com/problems/largest-local-values-in-a-matrix/
// Difficulty: Easy
// Time O(n^2) | Space O(n^2)

import "fmt"

func main() {
	fmt.Println(LargestLocalValuesInAMatrix([][]int{{9, 9, 8, 1}, {5, 6, 2, 6}, {8, 2, 6, 4}, {6, 2, 2, 2}})) // [[9,9],[8,6]]
	fmt.Println(LargestLocalValuesInAMatrix([][]int{{1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 2, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}})) // [[2,2,2],[2,2,2],[2,2,2]]
}

func LargestLocalValuesInAMatrix(grid [][]int) [][]int {
	n := len(grid)
	res := make([][]int, n-2)
	for i := 0; i < n-2; i++ {
		res[i] = make([]int, n-2)
		for j := 0; j < n-2; j++ {
			maxVal := 0
			for r := i; r < i+3; r++ {
				for c := j; c < j+3; c++ {
					if grid[r][c] > maxVal {
						maxVal = grid[r][c]
					}
				}
			}
			res[i][j] = maxVal
		}
	}
	return res
}
```

## 2377 — Sort The Olympic Table

```go
package main

// LeetCode #2377: Sort the Olympic Table
// https://leetcode.com/problems/sort-the-olympic-table/
// Difficulty: Easy [Paid]

import "fmt"

func main() {
	fmt.Println(SortTheOlympicTable())
}

func SortTheOlympicTable() any {
	// TODO: implement
	return nil
}
```

## 2379 — Minimum Recolors To Get K Consecutive Black Blocks

```go
package main

// LeetCode #2379: Minimum Recolors to Get K Consecutive Black Blocks
// https://leetcode.com/problems/minimum-recolors-to-get-k-consecutive-black-blocks/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(MinimumRecolorsToGetKConsecutiveBlackBlocks("WBBWWBBWBW", 7)) // 3
	fmt.Println(MinimumRecolorsToGetKConsecutiveBlackBlocks("WBWBBBW", 2))    // 0
}

func MinimumRecolorsToGetKConsecutiveBlackBlocks(blocks string, k int) int {
	wCount := 0
	for i := 0; i < k; i++ {
		if blocks[i] == 'W' {
			wCount++
		}
	}
	minOps := wCount
	for i := k; i < len(blocks); i++ {
		if blocks[i-k] == 'W' {
			wCount--
		}
		if blocks[i] == 'W' {
			wCount++
		}
		if wCount < minOps {
			minOps = wCount
		}
	}
	return minOps
}
```

## 2383 — Minimum Hours Of Training To Win A Competition

```go
package main

// LeetCode #2383: Minimum Hours of Training to Win a Competition
// https://leetcode.com/problems/minimum-hours-of-training-to-win-a-competition/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(MinimumHoursOfTrainingToWinACompetition(5, 3, []int{1, 4, 3, 2}, []int{2, 6, 3, 1})) // 8
	fmt.Println(MinimumHoursOfTrainingToWinACompetition(2, 4, []int{1}, []int{3}))                   // 0
}

func MinimumHoursOfTrainingToWinACompetition(initialEnergy int, initialExperience int, energy []int, experience []int) int {
	totalHours := 0
	curEnergy := initialEnergy
	curExp := initialExperience

	for i := 0; i < len(energy); i++ {
		// Energy
		if curEnergy <= energy[i] {
			needed := energy[i] - curEnergy + 1
			totalHours += needed
			curEnergy += needed
		}
		curEnergy -= energy[i]

		// Experience
		if curExp <= experience[i] {
			needed := experience[i] - curExp + 1
			totalHours += needed
			curExp += needed
		}
		curExp += experience[i]
	}
	return totalHours
}
```

## 2389 — Longest Subsequence With Limited Sum

```go
package main

// LeetCode #2389: Longest Subsequence With Limited Sum
// https://leetcode.com/problems/longest-subsequence-with-limited-sum/
// Difficulty: Easy
// Time O((n+m) log n) | Space O(1)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(LongestSubsequenceWithLimitedSum([]int{4, 5, 2, 1}, []int{3, 10, 21})) // [2,3,4]
	fmt.Println(LongestSubsequenceWithLimitedSum([]int{2, 3, 4, 5}, []int{1}))          // [0]
}

func LongestSubsequenceWithLimitedSum(nums []int, queries []int) []int {
	sort.Ints(nums)
	prefix := make([]int, len(nums))
	sum := 0
	for i, n := range nums {
		sum += n
		prefix[i] = sum
	}

	res := make([]int, len(queries))
	for i, q := range queries {
		// Binary search for rightmost index where prefix <= q
		res[i] = sort.SearchInts(prefix, q+1)
	}
	return res
}
```

## 2395 — Find Subarrays With Equal Sum

```go
package main

// LeetCode #2395: Find Subarrays With Equal Sum
// https://leetcode.com/problems/find-subarrays-with-equal-sum/
// Difficulty: Easy
// Time O(n) | Space O(n)

import "fmt"

func main() {
	fmt.Println(FindSubarraysWithEqualSum([]int{4, 2, 4}))   // true
	fmt.Println(FindSubarraysWithEqualSum([]int{1, 2, 3, 4, 5})) // false
}

func FindSubarraysWithEqualSum(nums []int) bool {
	seen := map[int]bool{}
	for i := 1; i < len(nums); i++ {
		sum := nums[i-1] + nums[i]
		if seen[sum] {
			return true
		}
		seen[sum] = true
	}
	return false
}
```

## 2399 — Check Distances Between Same Letters

```go
package main

// LeetCode #2399: Check Distances Between Same Letters
// https://leetcode.com/problems/check-distances-between-same-letters/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(CheckDistancesBetweenSameLetters("abaccb", []int{1, 3, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})) // true
	fmt.Println(CheckDistancesBetweenSameLetters("aa", []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}))    // false
}

func CheckDistancesBetweenSameLetters(s string, distance []int) bool {
	first := [26]int{}
	for i := 0; i < 26; i++ {
		first[i] = -1
	}
	for i := 0; i < len(s); i++ {
		idx := s[i] - 'a'
		if first[idx] == -1 {
			first[idx] = i
		} else if i-first[idx]-1 != distance[idx] {
			return false
		}
	}
	return true
}
```

## 2404 — Most Frequent Even Element

```go
package main

// LeetCode #2404: Most Frequent Even Element
// https://leetcode.com/problems/most-frequent-even-element/
// Difficulty: Easy
// Time O(n) | Space O(n)

import "fmt"

func main() {
	fmt.Println(MostFrequentEvenElement([]int{0, 1, 2, 2, 4, 4, 1})) // 2
	fmt.Println(MostFrequentEvenElement([]int{4, 4, 4, 9, 2, 4}))    // 4
	fmt.Println(MostFrequentEvenElement([]int{1, 3, 5, 7}))           // -1
}

func MostFrequentEvenElement(nums []int) int {
	freq := map[int]int{}
	for _, n := range nums {
		if n%2 == 0 {
			freq[n]++
		}
	}
	if len(freq) == 0 {
		return -1
	}
	bestNum := -1
	bestCount := 0
	for n, c := range freq {
		if c > bestCount || (c == bestCount && n < bestNum) {
			bestNum = n
			bestCount = c
		}
	}
	return bestNum
}
```

## 2409 — Count Days Spent Together

```go
package main

// LeetCode #2409: Count Days Spent Together
// https://leetcode.com/problems/count-days-spent-together/
// Difficulty: Easy
// Time O(1) | Space O(1)

import "fmt"

var daysInMonth = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func main() {
	fmt.Println(CountDaysSpentTogether("08-15", "08-18", "08-16", "08-19")) // 3
	fmt.Println(CountDaysSpentTogether("10-01", "10-31", "11-01", "12-31")) // 0
}

func dayOfYear(date string) int {
	mm := int(date[0]-'0')*10 + int(date[1]-'0')
	dd := int(date[3]-'0')*10 + int(date[4]-'0')
	days := 0
	for m := 0; m < mm-1; m++ {
		days += daysInMonth[m]
	}
	return days + dd
}

func CountDaysSpentTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
	aStart := dayOfYear(arriveAlice)
	aEnd := dayOfYear(leaveAlice)
	bStart := dayOfYear(arriveBob)
	bEnd := dayOfYear(leaveBob)

	start := max(aStart, bStart)
	end := min(aEnd, bEnd)

	if start > end {
		return 0
	}
	return end - start + 1
}
```

## 2413 — Smallest Even Multiple

```go
package main

// LeetCode #2413: Smallest Even Multiple
// https://leetcode.com/problems/smallest-even-multiple/
// Difficulty: Easy
// Time O(1) | Space O(1)

import "fmt"

func main() {
	fmt.Println(SmallestEvenMultiple(5)) // 10
	fmt.Println(SmallestEvenMultiple(6)) // 6
}

func SmallestEvenMultiple(n int) int {
	if n%2 == 0 {
		return n
	}
	return n * 2
}
```

## 2418 — Sort The People

```go
package main

// LeetCode #2418: Sort the People
// https://leetcode.com/problems/sort-the-people/
// Difficulty: Easy
// Time O(n log n) | Space O(n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(SortThePeople([]string{"Mary", "John", "Emma"}, []int{180, 165, 170})) // ["Mary","Emma","John"]
	fmt.Println(SortThePeople([]string{"Alice", "Bob", "Bob"}, []int{155, 185, 150}))   // ["Bob","Alice","Bob"]
}

func SortThePeople(names []string, heights []int) []string {
	n := len(names)
	idx := make([]int, n)
	for i := 0; i < n; i++ {
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool {
		return heights[idx[i]] > heights[idx[j]]
	})
	res := make([]string, n)
	for i, id := range idx {
		res[i] = names[id]
	}
	return res
}
```

## 2423 — Remove Letter To Equalize Frequency

```go
package main

// LeetCode #2423: Remove Letter To Equalize Frequency
// https://leetcode.com/problems/remove-letter-to-equalize-frequency/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(RemoveLetterToEqualizeFrequency("abcc")) // true
	fmt.Println(RemoveLetterToEqualizeFrequency("aazz")) // false
	fmt.Println(RemoveLetterToEqualizeFrequency("bac"))  // true
}

func RemoveLetterToEqualizeFrequency(word string) bool {
	freq := make([]int, 26)
	for i := 0; i < len(word); i++ {
		freq[word[i]-'a']++
	}

	// Try removing one occurrence of each letter
	for i := 0; i < 26; i++ {
		if freq[i] == 0 {
			continue
		}
		freq[i]--
		if allSameFreq(freq) {
			return true
		}
		freq[i]++
	}
	return false
}

func allSameFreq(freq []int) bool {
	target := 0
	for _, f := range freq {
		if f == 0 {
			continue
		}
		if target == 0 {
			target = f
		} else if f != target {
			return false
		}
	}
	return true
}
```

## 2427 — Number Of Common Factors

```go
package main

// LeetCode #2427: Number of Common Factors
// https://leetcode.com/problems/number-of-common-factors/
// Difficulty: Easy
// Time O(min(a,b)) | Space O(1)

import "fmt"

func main() {
	fmt.Println(NumberOfCommonFactors(12, 6)) // 4
	fmt.Println(NumberOfCommonFactors(25, 30)) // 2
}

func NumberOfCommonFactors(a int, b int) int {
	count := 0
	n := a
	if b < n {
		n = b
	}
	for i := 1; i <= n; i++ {
		if a%i == 0 && b%i == 0 {
			count++
		}
	}
	return count
}
```

## 2432 — The Employee That Worked On The Longest Task

```go
package main

// LeetCode #2432: The Employee That Worked on the Longest Task
// https://leetcode.com/problems/the-employee-that-worked-on-the-longest-task/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(TheEmployeeThatWorkedOnTheLongestTask(10, [][]int{{0, 3}, {2, 5}, {0, 9}, {1, 15}}))   // 1
	fmt.Println(TheEmployeeThatWorkedOnTheLongestTask(26, [][]int{{1, 1}, {3, 7}, {2, 12}, {7, 17}})) // 3
}

func TheEmployeeThatWorkedOnTheLongestTask(n int, logs [][]int) int {
	bestID := logs[0][0]
	bestTime := logs[0][1]
	prevEnd := logs[0][1]

	for i := 1; i < len(logs); i++ {
		id := logs[i][0]
		start := prevEnd
		end := logs[i][1]
		duration := end - start
		if duration > bestTime || (duration == bestTime && id < bestID) {
			bestID = id
			bestTime = duration
		}
		prevEnd = end
	}
	return bestID
}
```

