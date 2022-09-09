# you can write to stdout for debugging purposes, e.g.
# print("this is a debug message")

def solution(S:str):
    # write your code in Python (Python 3.6)
    length = len(S)
    i = 0
    start = 0
    end = -1
    sub = ''
    count = 0
    while i < length:
        tmp = S[i]
        if S[i+1:].count(tmp) > 0:
            end = end if S[i+1:].index(tmp)+i+1 < end else S[i+1:].index(tmp)+i+1
        else:
            if end > start:
                sub = S[start:end+1]
                for j in range(len(sub)):
                    t2 = sub[j]
                    real = True
                    if sub.count(t2) % 2 != 0:
                        end = -1
                        start += 1
                        i = start
                        real = False
                        break
                count = count if count > len(sub) or not real else len(sub)
        i += 1
    return count


print(solution('bdaaadadb'))