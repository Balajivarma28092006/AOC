grid = [list(line) for line in open("day7_inputs.txt", "r").read().splitlines()]

rows = len(grid)
cols = len(grid[0])

start_col = grid[0].index("S")

visited = set()
count = 0

# just backtrack for every direction
def backtrack1(row, col):
    global count

    if row < 0 or row >= rows or col < 0 or col >= cols:
        return

    if (row, col) in visited:
        return

    visited.add((row, col))
    curr = grid[row][col]

    if curr == "^":
        count += 1
        backtrack1(row + 1, col - 1)
        backtrack1(row + 1, col + 1)
    elif curr == ".":
        backtrack1(row + 1, col)

backtrack1(1, start_col)
print(f"Part ones solution is {count}")

count = 0
memo = [[-1] * cols for _ in range(rows)]


def backtrack2(row, col):
    global count

    # if we went out side the frame in horizantally
    if col < 0 or col >= cols:
        return 0

    # if the splitter reached the end of the grid
    if row == rows:
        return 1

    if memo[row][col] != -1:
        return memo[row][col]

    ans = 0
    curr = grid[row][col]

    # if there is no splitter just go normally
    if curr == "S"  or curr == "|" or curr == '.':
        ans = backtrack2(row + 1, col)
    elif curr == '^':
        ans = backtrack2(row + 1, col - 1) + backtrack2(row + 1, col + 1)

    memo[row][col] = ans
    return memo[row][col]

count = backtrack2(1, start_col)
print(f"Part2 solution is : {count}")
