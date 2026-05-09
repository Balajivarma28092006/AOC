class Warehouse:
    def __init__(self):
        self.grid = self.read_grid()
        self.moves = self.read_moves()
        self.robot = self.find_robot()

    def read_grid(self):
        with open("grid.txt") as f:
            return [list(line.strip()) for line in f if line.strip()]

    def read_moves(self):
        with open("directions.txt") as f:
            return ''.join(ch for ch in f.read() if ch in "<>^v")

    def find_robot(self):
        for r in range(len(self.grid)):
            for c in range(len(self.grid[r])):
                if self.grid[r][c] == '@':
                    return r, c

        raise ValueError("Robot '@' not found in grid")

    def move(self, direction):
        directions = {
            '^': (-1, 0),
            'v': (1, 0),
            '<': (0, -1),
            '>': (0, 1)
        }

        dr, dc = directions[direction]

        r, c = self.robot
        nr, nc = r + dr, c + dc

        # Wall
        if self.grid[nr][nc] == '#':
            return

        # Empty space
        if self.grid[nr][nc] == '.':
            self.grid[r][c] = '.'
            self.grid[nr][nc] = '@'
            self.robot = (nr, nc)
            return

        # Push boxes
        if self.grid[nr][nc] == 'O':
            br, bc = nr, nc

            while self.grid[br][bc] == 'O':
                br += dr
                bc += dc

            if self.grid[br][bc] == '.':
                self.grid[br][bc] = 'O'
                self.grid[nr][nc] = '@'
                self.grid[r][c] = '.'
                self.robot = (nr, nc)

    def run(self):
        for move in self.moves:
            self.move(move)

    def gps_sum(self):
        total = 0

        for r in range(len(self.grid)):
            for c in range(len(self.grid[r])):
                if self.grid[r][c] == 'O':
                    total += 100 * r + c

        return total


warehouse = Warehouse()
warehouse.run()

print(warehouse.gps_sum())