package main

// LeetCode #3451: Find Invalid IP Addresses
// https://leetcode.com/problems/find-invalid-ip-addresses/
// Difficulty: Hard
//
// Given a list of IP addresses, find which ones are invalid.
// An IPv4 address is valid if it has exactly 4 parts separated by dots,
// each part is an integer between 0 and 255 without leading zeros.
//
// This is originally a SQL problem. Implemented as a Go function for
// consistency.

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Example 1
	fmt.Println(findInvalidIPAddresses([]string{"192.168.1.1", "256.1.2.3", "1.2.3.4"}))
	// Example 2
	fmt.Println(findInvalidIPAddresses([]string{"0.0.0.0", "01.2.3.4", "192.168.001.1"}))
	// Edge: all valid
	fmt.Println(findInvalidIPAddresses([]string{"10.0.0.1", "172.16.0.1"}))
	// Edge: empty
	fmt.Println(findInvalidIPAddresses([]string{}))
}

func findInvalidIPAddresses(ips []string) []string {
	var invalid []string

	for _, ip := range ips {
		parts := strings.Split(ip, ".")
		if len(parts) != 4 {
			invalid = append(invalid, ip)
			continue
		}

		valid := true
		for _, part := range parts {
			// Check for leading zeros
			if len(part) > 1 && part[0] == '0' {
				valid = false
				break
			}
			// Check range
			num, err := strconv.Atoi(part)
			if err != nil || num < 0 || num > 255 {
				valid = false
				break
			}
		}

		if !valid {
			invalid = append(invalid, ip)
		}
	}

	return invalid
}
