class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        count = {}
        freq = [[] for i in range(len(nums) + 1)]
        for num in nums:
            count[num] = 1 + count.get(num, 0)
        for num, cout in count.items():
            freq[cout].append(num)
        res = []
        for arr in range(len(freq) - 1, 0, -1):
            for digit in freq[arr]:
                res.append(digit)
                if len(res) == k:
                    return res
        return []