# you can write to stdout for debugging purposes, e.g.
# print("this is a debug message")

def solution(A, M):
    # write your code in Python (Python 3.6)
    A.sort()
    max = A[-1]
    count = 0
    for i in range(len(A)):
        tmp = A[i]
        tmp_index = i + 1
        tmp_count = 1
        while tmp <= max:
            if A[tmp_index:].count(tmp) > 0:
                tmp_count += 1
                tmp_index = A[tmp_index:].index(tmp) + tmp_index + 1
            else:
                tmp += M
        count = tmp_count if tmp_count > count else count
    return count


print(solution([-3,-2,1,0,8,7,1],3))