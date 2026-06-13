class MinStack:

    def __init__(self):
        self.array = collections.deque()
        self.minstack = collections.deque()
    def push(self, val: int) -> None:
        self.array.append(val)
        if not self.minstack or self.minstack[-1] >= val:
            self.minstack.append(val)

    def pop(self) -> None:
        popped = self.array.pop()
        if popped == self.minstack[-1]:
            self.minstack.pop()

    def top(self) -> int:
        return self.array[-1]

    def getMin(self) -> int:
        return self.minstack[-1]
        
