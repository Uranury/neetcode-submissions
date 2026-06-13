class Solution:
    def dailyTemperatures(self, temps: List[int]) -> List[int]:
        result = [0] * len(temps)
        stack = []      #pair [Temp, Index]
        for index, temp in enumerate(temps):
            while stack and temp > stack[-1][0]: #first element of last tuple(Temp)
                stackT, stackInd = stack.pop()
                result[stackInd] = index - stackInd
            stack.append((temp, index))
        return result
            
