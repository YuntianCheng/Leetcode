#
# @lc app=leetcode.cn id=557 lang=python3
#
# [557] 反转字符串中的单词 III
#

# @lc code=start
class Solution:
    def reverseWords(self, s: str) -> str:
       words = s.split(' ')
       for i in range(len(words)):
        newWord = reversed(words[i])
        words[i] = ''.join(list(newWord))
       return ' '.join(words)

                
# @lc code=end

print(Solution().reverseWords("Let's take LeetCode contest"))