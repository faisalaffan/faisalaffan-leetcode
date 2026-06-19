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
