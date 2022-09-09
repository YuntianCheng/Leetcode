#
# @lc app=leetcode.cn id=481 lang=python3
#
# [481] 神奇字符串
#

# @lc code=start
class Solution:
    def magicalString(self, n: int) -> int:
        amazing = [1]
        should_index = 1
        should_length = 2
        should = 2
        count = 0
        count_one = 1
        while len(amazing) < n:
            if count < should_length:
                amazing.append(should)
                count += 1
                if should == 1:
                    count_one += 1
            else:
                should = 3 - should
                count = 0
                should_index += 1
                should_length = amazing[should_index]
        return count_one


# @lc code=end

print(Solution().magicalString(6))