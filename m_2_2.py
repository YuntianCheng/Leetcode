import fractions
def solution(X, Y):
    fs = []
    count = 0
    for i in range(len(X)):
        fs.append(fractions.Fraction(X[i],Y[i]))
    fs.sort()
    i = 0
    j = len(X)-1
    while i < j:
        if fs[i] + fs[j] < 1:
            i+=1
        elif fs[i] +fs[j] == 1:
            count += 1
            k = j-1
            while i < k:
                if fs[k] == fs[j]:
                    count+=1
                    k -= 1
                else:
                    break
            i += 1
        else:
            j -= 1
    return count

print(solution([1,1,1],[2,2,2]))