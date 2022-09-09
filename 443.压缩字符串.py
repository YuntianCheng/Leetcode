#
# @lc app=leetcode.cn id=443 lang=python3
#
# [443] 压缩字符串
#
from typing import List
# @lc code=start
class Solution:
    def compress(self, chars: List[str]) -> int:
        i = 0
        k = 0
        count = 1
        for j in range(1,len(chars)):
            if chars[i] == chars[j]:
                count+=1
            else:
                chars[k] = chars[i]
                k += 1
                if count > 1:
                    count_str = str(count)
                    for c in count_str:
                        chars[k] = c
                        k += 1
                count = 1
                i = j
        chars[k] = chars[i]
        k += 1
        if count > 1:
            count_str = str(count)
            for c in count_str:
                chars[k] = c
                k += 1
        return len(chars[:k])
        
        

# @lc code=end

