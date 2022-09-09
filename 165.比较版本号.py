#
# @lc app=leetcode.cn id=165 lang=python3
#
# [165] 比较版本号
#

# @lc code=start
class Solution:
    def compareVersion(self, version1: str, version2: str) -> int:
        version1_list = version1.split('.')
        version2_list = version2.split('.')
        length = len(version1_list) if len(version1_list) > len(version2_list) else len(version2_list)
        for i in range(length):
            v1 = int(version1_list[i]) if i < len(version1_list) else 0
            v2 = int(version2_list[i]) if i < len(version2_list) else 0
            if v1 > v2:
                return 1
            if v1 < v2:
                return -1
        return 0
                
# @lc code=end

print(Solution().compareVersion('1.01','1.001'))