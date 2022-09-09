# you can write to stdout for debugging purposes, e.g.
# print("this is a debug message")

from unittest import result


def solution(S):
    # write your code in Python (Python 3.6)
    count = [0]*10
    for s in S:
        count[int(s)] += 1
    i = 9
    result = []
    while i > -1: 
        if count[i] > 1:
            result.append(str(i))
            count[i] -= 2
        else:
            i -= 1
    if len(result)*2 < len(S):
        for i in range(9,-1,-1):
            if count[i] > 0:
                result.append(str(i))
                break
        result = result + [result[j] for j in range(len(result)-2, -1, -1)]
    else:
        result = result + [result[j] for j in range(len(result)-1, -1, -1)]

    tmp = ''.join(result).strip('0')
    if len(tmp) == 0:
        return '0'
    return tmp


print(solution('0000000'))
