# you can write to stdout for debugging purposes, e.g.
# print("this is a debug message")

def solution(A, B):
    # write your code in Python (Python 3.6)
    length = len(A)
    # C = []
    # for i in range(length):
    #     C.append(A[i] if A[i] > B[i] else B[i])
    for i in range(length):
        tmp = abs(A[i]) if abs(A[i]) > B[i] else B[i]
        if tmp <= length and A[tmp-1] > 0:
            A[tmp-1] = - A[tmp-1]
    for j in range(length):
        if A[j] > 0:
            return j+1
    return length + 1
