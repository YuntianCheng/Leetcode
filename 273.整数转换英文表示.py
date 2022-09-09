#
# @lc app=leetcode.cn id=273 lang=python3
#
# [273] 整数转换英文表示
#

# @lc code=start
from math import fabs


numbers = {
    '0':'Zero',
    '1':'One',
    '2':'Two',
    '3':'Three',
    '4':'Four',
    '5':'Five',
    '6':'Six',
    '7':'Seven',
    '8':'Eight',
    '9':'Nine',
    '10':'Ten',
    '11':'Eleven',
    '12':'Twelve'
}
special_pre = {
    '2':['Twen','Twen'],
    '3':['Thir','Thir'],
    '4':['Four','For'],
    '5':['Fif','Fif'],
    '8':['Eigh','Eigh']
}
end = ['teen','ty']

bigs = [{
        'length':3,
        'word':'Hundred'
    }, 
    {
        'length':4,
        'word':'Thousand'
    },
    {
        'length':7,
        'word':'Million'
    },
    {
        'length':10,
        'word':'Billion'
    }]
    

class Solution:
    def results(self, num: int, zero: bool) -> str:
        num_str = str(num)
        if num <= 12:
            if num == 0 and not zero:
                return '' 
            return numbers[num_str]
        result = ''
        if num < 100:
            spec = num_str[0] if num >= 20 else num_str[1]
            if spec in special_pre:
                result += special_pre[spec][1] if num >= 20 else special_pre[spec][0]
            else:
                result += numbers[spec]
            if num < 20:
                result += end[0]
                return result
            else:
                result += end[1]
            result += (' ' + numbers[num_str[1]]) if num_str[1] != '0' else ''
            return result
        length = len(num_str)
        for b in reversed(bigs):
            if b['length'] <= length:
                index = length - b['length']
                return self.results(int(num_str[:index+1]), False).removesuffix(' ').removeprefix(' ') + ' ' + b['word'] + ' ' + self.results(int(num_str[index+1:]), False).removesuffix(' ').removeprefix(' ')
    def numberToWords(self, num: int) -> str:
        return self.results(num,True).removesuffix(' ').removeprefix(' ')
        



# @lc code=end

print(Solution().numberToWords(100))