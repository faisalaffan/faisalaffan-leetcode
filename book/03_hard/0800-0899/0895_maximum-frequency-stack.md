# 0895 — Maximum Frequency Stack

## Deskripsi

**Soal:** [0895. Maximum Frequency Stack](https://leetcode.com/problems/maximum-frequency-stack/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Stack (tumpukan LIFO)

**Fungsi Solusi:** `func Constructor() FreqStack`

## Solusi Go

```go
package main

// LeetCode #895: Maximum Frequency Stack
// https://leetcode.com/problems/maximum-frequency-stack/
// Difficulty: Hard
//
// Maintain:
//   - freq: map from value to its frequency
//   - groups: map from frequency to stack of values with that frequency
//   - maxFreq: current maximum frequency
//
// Push: increment freq[val], add val to groups[freq], update maxFreq.
// Pop: pop from groups[maxFreq], decrement freq[val], if groups[maxFreq] empty, decrement maxFreq.

import "fmt"

type FreqStack struct {
	freq    map[int]int
	groups  map[int][]int
	maxFreq int
}

func Constructor() FreqStack {
	return FreqStack{
		freq:   make(map[int]int),
		groups: make(map[int][]int),
	}
}

func (fs *FreqStack) Push(val int) {
	f := fs.freq[val] + 1
	fs.freq[val] = f
	fs.groups[f] = append(fs.groups[f], val)
	if f > fs.maxFreq {
		fs.maxFreq = f
	}
}

func (fs *FreqStack) Pop() int {
	stack := fs.groups[fs.maxFreq]
	val := stack[len(stack)-1]
	fs.groups[fs.maxFreq] = stack[:len(stack)-1]
	fs.freq[val]--
	if len(fs.groups[fs.maxFreq]) == 0 {
		fs.maxFreq--
	}
	return val
}

func main() {
	// Example: push 5,7,5,7,4,5 -> pop -> 5, pop -> 7, pop -> 5, pop -> 4
	fs := Constructor()
	fs.Push(5)
	fs.Push(7)
	fs.Push(5)
	fs.Push(7)
	fs.Push(4)
	fs.Push(5)
	fmt.Println("Test 1 pop:", fs.Pop()) // 5 (freq 3)
	fmt.Println("Test 2 pop:", fs.Pop()) // 7 (freq 2, most recent)
	fmt.Println("Test 3 pop:", fs.Pop()) // 5 (freq 2)
	fmt.Println("Test 4 pop:", fs.Pop()) // 4 (freq 1)

	// Fresh stack: push 1,2,3,1,2,1 -> pop->1, pop->2, pop->1
	fs2 := Constructor()
	fs2.Push(1)
	fs2.Push(2)
	fs2.Push(3)
	fs2.Push(1)
	fs2.Push(2)
	fs2.Push(1)
	fmt.Println("Test 5 pop:", fs2.Pop()) // 1 (freq 3)
	fmt.Println("Test 6 pop:", fs2.Pop()) // 2 (freq 2, most recent among freq 2)
	fmt.Println("Test 7 pop:", fs2.Pop()) // 1 (freq 2, the other freq 2)
	fmt.Println("Test 8 pop:", fs2.Pop()) // 3 (freq 1)
}
```
