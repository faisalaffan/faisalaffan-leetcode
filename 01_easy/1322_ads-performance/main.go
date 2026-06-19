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
