#
# @lc app=leetcode.cn id=696 lang=python3
#
# [696] 计数二进制子串
#

# @lc code=start
class Solution:
    def countBinarySubstrings(self, s: str) -> int:
        s_list = list(s)
        zero_count = 0
        one_count = 0
        count = 0
        for i in range(0,len(s_list)-1):
            if s_list[i] == '0':
                zero_count += 1
            else:
                one_count += 1
            if s_list[i] != s_list[i+1]:
                if zero_count > 0 and one_count > 0:
                    count += min(zero_count, one_count)
                    if s_list[i+1] == '0':
                        zero_count = 0
                    else:
                        one_count = 0
        if s_list[len(s_list)-1] == '0':
            zero_count += 1
        else:
            one_count += 1
        count += min(zero_count, one_count)
        return count

# @lc code=end
print(Solution().countBinarySubstrings('00110011'))
