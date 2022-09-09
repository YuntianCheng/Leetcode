# you can write to stdout for debugging purposes, e.g.
# print("this is a debug message")

def solution(S:str):
    # write your code in Python (Python 3.6)
    start = 0
    d = {0:-1}
    ans = 0
    for i in range(len(S)):
        cur = 1 << ord(S[i]) - ord('a')
        start ^=cur
        if start not in d:
            d[start] = i
        else:
            ans = max(ans, i - d[start])
    return ans


