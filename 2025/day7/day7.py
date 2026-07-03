grid = [list(line) for line in open("day7_inputs.txt", "r").read().splitlines()]

rows = len(grid)
cols = len(grid[0])

start_col = grid[0].index("S")

visited = set()
count = 0

# just backtrack for every direction
def backtrack(row, col):
    global count

    if row < 0 or row >= rows or col < 0 or col >= cols:
        return

    if (row, col) in visited:
        return

    visited.add((row, col))
    curr = grid[row][col]

    if curr == "^":
        count += 1
        backtrack(row + 1, col - 1)
        backtrack(row + 1, col + 1)
    elif curr == ".":
        backtrack(row + 1, col)

backtrack(1, start_col)
print(f"Part ones solution is {count}")
