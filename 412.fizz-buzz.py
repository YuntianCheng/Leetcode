#
# @lc app=leetcode.cn id=412 lang=python3
#
# [412] Fizz Buzz
#
from typing import List
# @lc code=start
class Solution:
    def fizzBuzz(self, n: int) -> List[str]:
        j = 3
        k = 5
        l = 15
        answer = [""] * (n+1)
        while True:
            if j <= n:
                if j % 15 != 0:
                    answer[j] = 'Fizz' 
                j += 3
            if k <= n:
                if k % 15 != 0:
                    answer[k] = 'Buzz'
                k += 5
            if l <= n:
                answer[l] = "FizzBuzz"
                l += 15
            if j >= n and k >= n and l >= n:
                break
        if j == n and j % 15 != 0:
            answer[j] = 'Fizz'
        if k == n and k % 15 != 0:
            answer[k] = 'Buzz' 
        if l == n:
            answer[l] = 'FizzBuzz'
        for i in range(1,n+1):
            if answer[i] == "":
                answer[i] = str(i)
        return answer[1:]




# @lc code=end

s = Solution()
print(s.fizzBuzz(9999))