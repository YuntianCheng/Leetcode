# you can write to stdout for debugging purposes, e.g.
# print("this is a debug message")

from unittest import result


def solution(S, B):
    # write your code in Python (Python 3.6)
    count = 0
    poles_num = []
    for s in S:
        if s == 'X':
            count += 1
        else:
            if count > 0:
                poles_num.append(count)
                count = 0
    if count > 0:
        poles_num.append(count)
    poles_num.sort(reverse=True)
    result = 0
    for num in poles_num:
        if num + 1 < B:
            result += num
            B = B-num-1
        else:
            result += (B-1)
            break
    return result

result = solution("..XXXXX",4)
print(result)
