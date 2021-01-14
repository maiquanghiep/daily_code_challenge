package maximumsquare


/*
Given an m x n binary matrix filled with 0's and 1's, 
find the largest square containing only 1's and return its area.
1 0 1 0 0
1 0 1 1 1
1 1 1 1 1
1 0 0 1 0

Output: 4

*/

/* To solve this problem, we first create a new matrix n+1, m+1
0 0 0 0 0 0
0 0 0 0 0 0
0 0 0 0 0 0
0 0 0 0 0 0
0 0 0 0 0 0


A 2*2 square can be formed at matrix[i],[j] if the previous surrounding is all 1. 
(matrix[i][j-1], matrix[i-1][j], matrix[i-1][j-1])


*/

func min(a,b int) int {
    if (a>b) {
        return b
    }
    return a
}

func max(a,b int) int {
    if (a>b) {
        return a
    }
    return b
}

func maximalSquare(matrix [][]byte) int {
    n:= len(matrix)
    m:= len(matrix[0])
    
    dp:= make([][]int, n +1)
    for i:= range(dp) {
        dp[i] = make([]int, m+1)
    }
    
    maxsquare := 0
    for i:= 1; i < len(dp); i++ {
        for j:= 1; j < len(dp[i]); j++ {
            if matrix[i-1][j-1] == '1' {
                dp[i][j] = 1 + min(min(dp[i][j-1], dp[i-1][j]) , dp[i-1][j-1])
            }
            maxsquare = max(maxsquare, dp[i][j])
        }
    }
    
    return maxsquare * maxsquare
}