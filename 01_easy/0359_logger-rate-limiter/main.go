package main

// LeetCode #359: Logger Rate Limiter
// https://leetcode.com/problems/logger-rate-limiter/
// Difficulty: Easy [Paid]

import "fmt"

// Logger tracks messages and their last printed timestamp.
type Logger struct {
	lastPrinted map[string]int
}

// Constructor creates a new Logger.
func Constructor() Logger {
	return Logger{lastPrinted: make(map[string]int)}
}

// ShouldPrintMessage returns true if the message should be printed at the given timestamp.
// A message can be printed if it hasn't been printed in the last 10 seconds.
// Time: O(1), Space: O(n)
func (l *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	if lastTs, ok := l.lastPrinted[message]; ok && timestamp-lastTs < 10 {
		return false
	}
	l.lastPrinted[message] = timestamp
	return true
}

func main() {
	obj := Constructor()
	fmt.Println(obj.ShouldPrintMessage(1, "foo"))
	fmt.Println(obj.ShouldPrintMessage(2, "bar"))
	fmt.Println(obj.ShouldPrintMessage(3, "foo"))
	fmt.Println(obj.ShouldPrintMessage(8, "bar"))
	fmt.Println(obj.ShouldPrintMessage(10, "foo"))
	fmt.Println(obj.ShouldPrintMessage(11, "foo"))
}
