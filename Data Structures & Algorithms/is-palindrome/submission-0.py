class Solution:
    def isPalindrome(self, s: str) -> bool:
        t = ""
        for char in s:
            if char.isalnum():
                t += char.lower()
        return t == t[::-1]