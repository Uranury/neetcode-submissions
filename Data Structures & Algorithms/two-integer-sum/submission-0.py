class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        dic = {}
        for index, num in enumerate(nums):
            compl = target - num
            if compl in dic:
                return [dic[compl], index]
            dic[num] = index
        return []