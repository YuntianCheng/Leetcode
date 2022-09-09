from operator import mod


def solution(A, X, Y):
    min_cost = -1
    for i in range(Y):
        tmp = 0
        tem_X = X
        for j in range(len(A)):
            if j%Y==i:
                if tem_X > 0:
                    tmp += A[j]
                    tem_X-=1
                    continue
                else:
                    index = j - (X-1)*Y
                    tmp = tmp - A[index] + A[j]
                if min_cost == -1 or tmp < min_cost:
                    min_cost = tmp
    return min_cost

print(solution([3,5,1,4.2,7],3,2))