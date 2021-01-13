package findthejudge

/*
In a town, there are N people labelled from 1 to N.  There is a rumor that one of these people is secretly the town judge.

If the town judge exists, then:

The town judge trusts nobody.
Everybody (except for the town judge) trusts the town judge.
There is exactly one person that satisfies properties 1 and 2.
You are given trust, an array of pairs trust[i] = [a, b] representing that the person labelled a trusts the person labelled b.

If the town judge exists and can be identified, return the label of the town judge.  Otherwise, return -1.
*/

/* This can be solved using graph represented by an N*N metrix (M)
M(i,j): Represent the trust of person i to person j.
Traverse each column j, if total trust of person j is N-1 meaning the person j is a candidate of judge
Check each column of person j, if there is no 0 val then the candidate is the judge
Complexity O(n^2): traver n to n
Space O(n^2): Create new N*N matrix
=> I Implemented this solution to understand representing Graph as Adjacency Matrix
*/

func findJudge(N int, trust [][]int) int {
	trustGraph := make([][]int, N)
	for i := range trustGraph {
		trustGraph[i] = make([]int, N)
	}

	for i := range trust {
		trustGraph[trust[i][0]-1][trust[i][1]-1] = 1
	}

	for col := range trustGraph {
		count := 0
		for row := range trustGraph {
			if trustGraph[row][col] == 1 {
				count++
			}
		}
		isJudge := true
		if count == N-1 {
			for k := range trustGraph {
				if trustGraph[col][k] == 1 {
					isJudge = false
					break
				}
			}
			if isJudge {
				return col + 1
			}
		}
	}

	return -1
}
