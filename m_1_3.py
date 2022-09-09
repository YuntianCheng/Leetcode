

def solution(A:list, B:list, S:int):
    # write your code in Python (Python 3.6)
    if S < len(A):
        return False
    choice = set()
    nums = {}
    for num in A+B:
        choice.add(num)
    if len(choice) < len(A):
        return False
    for i in range(len(A)):
        key = str(A[i])+','+str(B[i]) if A[i] > B[i] else str(B[i])+','+str(A[i])
        if key not in nums:
            nums[key] = 1
        else:
            nums[key] += 1
            if nums[key] > 2:
                return False
    return True

result = solution([1,2,3,1,1,1],[2,3,1,4,4,5],6)
print(result)