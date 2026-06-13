class Solution:
    def checkInclusion(self, s1: str, s2: str) -> bool:
        left, right = 0, len(s1) - 1
        s1map = {}
        for char in s1:
            s1map[char] = s1map.get(char, 0) + 1
        while right < len(s2):
            s2map = {}
            for i in range(left, right + 1):
                s2map[s2[i]] = s2map.get(s2[i], 0) + 1
                if s1map == s2map:
                    return True
            left += 1
            right += 1
        return False
            