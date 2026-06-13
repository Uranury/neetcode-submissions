class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        count = 0
        ans = 0
        stack = collections.deque([])
        for char in s:
            while char in stack:
                stack.popleft()
                count -= 1
            else:
                count += 1
            stack.append(char)
            ans = max(ans, count)
        return ans
            
            