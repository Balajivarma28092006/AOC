"""
Advent of Code 2025 - Day 4
Look at every @ in the grid.
Count how many neighboring cells also contain @.
If that count is less than 4, add it to the answer.
"""

directions = [(-1, -1), (-1, 0), (-1, 1), (0, -1), (0, 1), (1, -1), (1, 0), (1, 1)]

def solve_part1(grid):
    answer = 0
    for i in range(len(grid)):
        for j in range(len(grid[i])):
            if grid[i][j] == '@':
                count = 0
                for dx, dy in directions:
                    x, y = i + dx, j + dy
                    if 0 <= x < len(grid) and 0 <= y < len(grid[i]) and grid[x][y] == '@':
                        count += 1
                if count < 4:
                    answer += 1
    return answer

def solve_part2(grid):
    total_removed = 0
    while True:
        to_remove = []

        # Identify cells to remove
        for i in range(len(grid)):
            for j in range(len(grid[i])):
                if grid[i][j] == '@':
                    neighbor_count = 0
                    for dx, dy in directions:
                        x, y = i + dx, j + dy
                        if 0 <= x < len(grid) and 0 <= y < len(grid[i]) and grid[x][y] == '@':
                            neighbor_count += 1
                    if neighbor_count < 4:
                        to_remove.append((i, j))

        # stop if no cells to remove
        if not to_remove:
            break

        # Remove identified cells
        for g, f in to_remove:
            grid[g][f] = '.'
        
        total_removed += len(to_remove)

    return total_removed

grid = [list(line.strip()) for line in open("inputs.txt")]
print("Part 1:", solve_part1(grid))
print("Part 2:", solve_part2(grid))