#
# @lc app=leetcode.cn id=506 lang=python3
#
# [506] 相对名次
#
from typing import List
# @lc code=start
class Solution:
    def findRelativeRanks(self, score: List[int]) -> List[str]:
        orgin_seq = {}
        for i in range(len(score)):
            orgin_seq[score[i]] = i
        score.sort(reverse=True)
        answer = [""] * len(score)
        for i in range(len(score)):
            if i == 0:
                answer[orgin_seq[score[i]]] = "Gold Medal"
            elif i == 1:
                answer[orgin_seq[score[i]]] = "Silver Medal"
            elif i == 2:
                answer[orgin_seq[score[i]]] = "Bronze Medal"
            else:
                answer[orgin_seq[score[i]]] = str(i+1)
        return answer


# @lc code=end

