#
# @lc app=leetcode.cn id=299 lang=python3
#
# [299] 猜数字游戏
#

# @lc code=start
class Solution:
    def getHint(self, secret: str, guess: str) -> str:
        bullPosition = set()
        mayCow = {}
        length = len(secret)
        for i in range(length):
            if secret[i] == guess[i]:
                bullPosition.add(i)
            else:
                if secret[i] in mayCow:
                    mayCow[secret[i]] += 1
                else:
                    mayCow[secret[i]] = 1
        count = 0
        for i in range(length):
            if i not in bullPosition and guess[i] in mayCow:
                if mayCow[guess[i]] > 0:
                    count+=1
                    mayCow[guess[i]] -= 1
        return "".join([str(len(bullPosition)), 'A', str(count), 'B'])        
# @lc code=end

