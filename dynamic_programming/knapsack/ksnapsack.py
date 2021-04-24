'''
Given weights and values of n items and a value k. 
We need to choose a subset of these items in such a way that ratio of 
the sum of weight and sum of values of chosen items is K and sum of weight is maximum among all possible subset choices.

Input : val = [60, 50, 70, 30]
wt = [5, 3 , 4 , 2]
        W = 5
 Output : 80      
'''

val = [60, 50, 70, 30]
wt = [5, 3 , 4 , 2]
W = 5
n = len(val)

def knapSack(W, wt, val, n):
    K = [[0 for x in range(W + 1)] for x in range(n + 1)]

    for i in range(n + 1):
        for w in range(W + 1):
            if i == 0 or w == 0:
                K[i][w] = 0
            elif wt[i-1] <= w:
                K[i][w] = max(val[i-1]
                          + K[i-1][w-wt[i-1]], 
                              K[i-1][w])
            else:
                K[i][w] = K[i-1][w]
    print(K)

    return K[n][W]
    
print(knapSack(W, wt, val, n))