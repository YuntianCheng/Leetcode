# you can write to stdout for debugging purposes, e.g.
# print("this is a debug message")

def nums(key, d:dict):
    count = 1
    for k in d[key]:
        count += nums(k,d)
    return count


def solution(A, B):
    # write your code in Python (Python 3.6)
    d = {}
    q = []
    q.append(0)
    length = len(A)
    while len(q) > 0:
        root = q[0]
        d[root] = []
        for i in range(length):
            if A[i] == root:
                q.append(B[i])
                d[root].append(B[i])
                A[i] = -1
                B[i] = -1
            if B[i] == root:
                q.append(A[i])
                d[root].append(A[i])
                A[i] = -1
                B[i] = -1
        q.remove(root)
    count = 0
    q = []
    q = q+d[0]
    while len(q) > 0:
        root = q[0]
        num = nums(root,d)
        tmp = 0
        while num > 0:
            num -= 5
            tmp += 1
        count += tmp
        q.remove(root)
        q = q+d[root]
    return count




print(solution([0,1,1],[1,2,3]))