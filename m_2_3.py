def solution(A, X, Y):
    # write your code in Python (Python 3.6)
    min_cost = -1
    length = (X-1)*Y
    for i in range(len(A)-length):
        left_X = X
        tmp_cost = 0
        for j in range(i,len(A),Y):
            if left_X == 0:
                break
            tmp_cost += A[j]
            left_X -= 1
        if min_cost == -1 or tmp_cost < min_cost:
            min_cost = tmp_cost
    return min_cost

print(solution([4,2,3,7],2,2))