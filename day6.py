grid = [list(line.strip())for line in open("day6_inputs.txt")]

rows = len(grid)
cols = len(grid[0])

#fining the position of the soldier
for i in range(rows):
    for j in range(cols):
        if grid[i][j] in "^>v<":
            sr, sc = i, j
            direction = grid[i][j]

directions = {
    "^": (-1, 0),
    ">": (0, 1),
    "v": (1, 0),
    "<": (0, -1)
}

turn_right = {
    "^": ">",
    "v": "<",
    ">": "v",
    "<": "^"
}


visited = set()
r, c = sr, sc

while True:
    visited.add((r, c))

    dr, dc = directions[direction]
    nr, nc = r + dr, c + dc

    if nr < 0 or nr >= rows or nc < 0 or nc >= cols:
        break

    if (grid[nr][nc] == "#"):
        direction = turn_right[direction]
        continue

    r, c = nr, nc

print(len(visited))