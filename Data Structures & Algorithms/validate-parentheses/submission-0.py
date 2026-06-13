class Solution:
    def isValid(self, s: str) -> bool:
        if len(s)%2 != 0:
            return False
        stack = []
        for char in s:
            if char == "(":
                stack.append(")")
            elif char == "[":
                stack.append("]")
            elif char == "{":
                stack.append("}")
            else:
                if not stack or stack.pop() != char:
                    return False
        return not stack