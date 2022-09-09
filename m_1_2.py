

def solution(S:str):
    # write your code in Python (Python 3.6)
    S = S.strip('W')
    tmp = []
    count_R = 0
    count_W = 0
    for s in S:
        if s == 'W':
            count_W +=1
            if count_R > 0:
                tmp.append(-count_R)
            count_R = 0
        else:
            count_R +=1
            if count_W > 0:
                tmp.append(count_W)
            count_W = 0
    if count_R > 0:
        tmp.append(-count_R)
    if len(tmp) <= 1:
        return 0
    result = 0
    left_length = 0
    right_length = 0
    i = 0
    j = len(tmp) - 1
    while i<j:
        if tmp[i] < 0 and tmp[j] < 0:
            left_length -= tmp[i]
            right_length -= tmp[j]
        else:
            result += tmp[i] + tmp[j]
            for k in range(1,left_length):
                result += tmp[i] + k
            for k in range(1,right_length):
                result += tmp[j] + k
            if result > 1000000000:
                return -1
        i+=1
        j-=1
    if tmp[i] > 0:
        short_length = min(left_length, right_length)
        result += tmp[i]
        for k in range(1,short_length):
            result += tmp[i]+k
            if result > 1000000000:
                return -1
    return result

result = solution('WRRWWR')
print(result)
            





