package main

// LeetCode #1364: Number of Trusted Contacts of a Customer
// https://leetcode.com/problems/number-of-trusted-contacts-of-a-customer/
// Difficulty: Medium

import "fmt"
import "sort"

func main() {
	// SQL problem - simulating in Go
	type invoice struct {
		invoiceID    int
		customerID   int
		price        int
	}
	type contact struct {
		userID        int
		contactEmail  string
		trusted       bool
	}
	type customer struct {
		customerID int
		customerName string
		email       string
	}

	invoices := []invoice{
		{44, 2, 100},
		{55, 8, 120},
		{66, 9, 150},
	}
	customers := []customer{
		{1, "Alice", "alice@leetcode.com"},
		{2, "Bob", "bob@leetcode.com"},
		{13, "John", "john@leetcode.com"},
		{6, "Alex", "alex@leetcode.com"},
	}
	contacts := []contact{
		{1, "bob@leetcode.com", true},
		{1, "john@leetcode.com", true},
		{1, "jane@leetcode.com", false},
		{2, "alice@leetcode.com", true},
		{2, "john@leetcode.com", false},
	}

	result := trustedContacts(invoices, customers, contacts)
	for _, r := range result {
		fmt.Printf("%d %s %d %d\n", r.invoiceID, r.customerName, r.price, r.trustedCount)
	}
}

type invoiceResult struct {
	invoiceID    int
	customerName string
	price        int
	trustedCount int
}

// Time: O(n log n) for sorting
// Space: O(n)
func trustedContacts(invoices []struct {
	invoiceID  int
	customerID int
	price      int
}, customers []struct {
	customerID   int
	customerName string
	email        string
}, contacts []struct {
	userID       int
	contactEmail string
	trusted      bool
}) []invoiceResult {
	// Map customer emails to IDs for trusted contact lookup
	emailToCustomerID := make(map[string]int)
	for _, c := range customers {
		emailToCustomerID[c.email] = c.customerID
	}

	// Count trusted contacts per customer
	trustedCounts := make(map[int]int)
	for _, c := range contacts {
		if c.trusted {
			if custID, ok := emailToCustomerID[c.contactEmail]; ok {
				trustedCounts[custID]++
			}
		}
	}

	// Map customer IDs to names
	customerNames := make(map[int]string)
	for _, c := range customers {
		customerNames[c.customerID] = c.customerName
	}

	// Sort invoices by invoice ID
	sort.Slice(invoices, func(i, j int) bool {
		return invoices[i].invoiceID < invoices[j].invoiceID
	})

	var result []invoiceResult
	for _, inv := range invoices {
		name := customerNames[inv.customerID]
		if name == "" {
			name = ""
		}
		result = append(result, invoiceResult{
			inv.invoiceID,
			name,
			inv.price,
			trustedCounts[inv.customerID],
		})
	}

	return result
}
