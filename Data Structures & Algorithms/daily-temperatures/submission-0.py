class Solution:
    def dailyTemperatures(self, temps: List[int]) -> List[int]:
        result = []
        left, right = 0, 1
        while left < len(temps):
            if right < len(temps):
                if temps[left] < temps[right]:
                    result.append(right - left)
                    left += 1
                    right = left + 1
                else:
                    right += 1
            else:
                result.append(0)
                left += 1
                right = left + 1
        return result
            
