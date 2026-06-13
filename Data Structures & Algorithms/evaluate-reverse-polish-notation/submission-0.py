class Solution:
    def evalRPN(self, tokens: List[str]) -> int:
        dq = collections.deque()
        for token in tokens:
            if token.lstrip('-').isdigit():
                dq.append(int(token))
            else:
                b = dq.pop()
                a = dq.pop()
                if token == "-":
                    dq.append(a - b)
                elif token == "+":
                    dq.append(a + b)
                elif token == "/":
                    dq.append(int(a/b))
                elif token == "*":
                    dq.append(a*b)
        return dq[0]