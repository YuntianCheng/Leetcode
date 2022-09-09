# you can write to stdout for debugging purposes, e.g.
# print("this is a debug message")

def solution(X, Y, W):
    # write your code in Python (Python 3.6)
    X.sort()
    count = 1
    start = X[0]
    for i in range(1, len(X)):
        if X[i]-start>W:
            count += 1
            start = X[i]
    return count
