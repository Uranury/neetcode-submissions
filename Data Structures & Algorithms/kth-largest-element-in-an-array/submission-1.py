class Solution:
    def findKthLargest(self, nums: List[int], k: int) -> int:
        nums = [-num for num in nums]
        heapq.heapify(nums)
        res = []
        for _ in range(k):
            res.append(-nums[0])
            heapq.heappop(nums)
        return res[-1]
