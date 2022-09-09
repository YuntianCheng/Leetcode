#include<iostream>
#include<vector>
using namespace std;
class Solution {
public:

    vector<int> twoSum(vector<int>& nums, int target) {
        sort(nums.begin(),nums.end());
        int i=0, j=nums.size()-1;
        while(i<j){
            if(nums[i]+nums[j]==target)
                break;
            else if(nums[i]+nums[j]>target)
                j--;
            else
                i++;
        }
        int arr[2]={i,j};
        return vector<int>(arr,arr+2);
    }
};

int main(){
    vector<int> temp;
    temp.push_back(3);
    temp.push_back(2);
    temp.push_back(4);
    Solution s;
    vector<int> result=s.twoSum(temp,6);
    for(vector<int>::const_iterator iter=result.begin();iter!=result.end();iter++)
        cout<<(*iter)<<endl;

    return 0;
}