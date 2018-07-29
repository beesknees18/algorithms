package main

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func editDistanceDP(s, t string) int {
	maxLen := 100
	matchingSubstitutionCost := 0
	insertionCost := 0
	deletionCost := 0

	// println("Running edit distance for pattern: ", s, " and text: ", t)

	dp := make([][]int, maxLen)

	// initializing the matrix
	for i := 0; i < maxLen; i++ {
		dp[i] = make([]int, maxLen)

		if i == 0 {
			for j := 0; j < maxLen; j++ {
				dp[i][j] = j
			}
		} else {
			dp[i][0] = i
		}
	}

	// dp algoritm
	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(t); j++ {
			if string(s[i-1]) == string(t[j-1]) {
				matchingSubstitutionCost = dp[i-1][j-1]
			} else {
				matchingSubstitutionCost = dp[i-1][j-1] + 1
			}

			insertionCost = dp[i-1][j] + 1
			deletionCost = dp[i][j-1] + 1

			// fmt.Printf("For i: %v and j: %v, matching-substitution cost is %v", i, j, matchingSubstitutionCost)
			// fmt.Printf("For i: %v and j: %v, insertion cost is %v", i, j, insertionCost)
			// fmt.Printf("For i: %v and j: %v, deletion cost is %v", i, j, deletionCost)

			dp[i][j] = min(matchingSubstitutionCost, min(insertionCost, deletionCost))

		}
	}

	return dp[len(s)][len(t)]
}
