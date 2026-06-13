class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False
        smap, tmap = {}, {}
        for char_s, char_t in zip(s, t):
            smap[char_s] = smap.get(char_s, 0) + 1
            tmap[char_t] = tmap.get(char_t, 0) + 1
        return smap == tmap