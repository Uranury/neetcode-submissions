class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        left, right, maxprof = 0, 1, 0
        while right < len(prices): 
            diff = prices[right] - prices[left]
            if diff < 0:
                left = right
            else:
                maxprof = max(maxprof, diff)
            right += 1
        return maxprof 