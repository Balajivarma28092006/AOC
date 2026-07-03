from math import prod

def solve(filename="inputs.txt"):
    with open(filename, "r") as f:
        lines = f.read().splitlines()

    width = max(len(line) for line in lines)
    grid = [line.ljust(width) for line in lines]

    rows = len(grid)
    op_row = rows - 1

    blank_col = [
        all(grid[r][c] == " " for r in range(rows))
        for c in range(width)
    ]

    part1 = 0
    part2 = 0

    c = 0
    while c < width:
        if blank_col[c]:
            c += 1
            continue

        start = c
        while c < width and not blank_col[c]:
            c += 1
        end = c

        op = grid[op_row][start:end].strip()

        # ---------- Part 1 ----------
        nums1 = []
        for row in range(op_row):
            s = grid[row][start:end].strip()
            if s:
                nums1.append(int(s))

        if op == "+":
            part1 += sum(nums1)
        else:
            part1 += prod(nums1)

        # ---------- Part 2 ----------
        nums2 = []
        for col in range(end - 1, start - 1, -1):
            digits = []

            for row in range(op_row):
                ch = grid[row][col]
                if ch.isdigit():
                    digits.append(ch)

            if digits:
                nums2.append(int("".join(digits)))

        if op == "+":
            part2 += sum(nums2)
        else:
            part2 += prod(nums2)

    print("Part 1:", part1)
    print("Part 2:", part2)


if __name__ == "__main__":
    solve("inputs.txt")
