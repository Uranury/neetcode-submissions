class Solution:
    def maxAreaOfIsland(self, grid: List[List[int]]) -> int:
        maxarea = 0
        rows, cols = len(grid), len(grid[0])
        
        def bfs(i, j):
            q = collections.deque([(i, j)])
            grid[i][j] = 0
            area = 1
            directions = [(0, 1), (0, -1), (1, 0), (-1, 0)]
            while q:
                r, c = q.popleft()
                for dr, dc in directions:
                    nr, nc = r + dr, c + dc
                    if 0 <= nr < rows and 0 <= nc < cols and grid[nr][nc] == 1:
                        grid[nr][nc] = 0
                        area += 1
                        q.append((nr, nc))
            return area

        for r in range(rows):
            for c in range(cols):
                if grid[r][c] == 1:
                    maxarea = max(maxarea, bfs(r,c))

        return maxarea