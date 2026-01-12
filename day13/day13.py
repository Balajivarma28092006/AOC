import re

def solve(lines, part2=False):
    total_cost = 0

    lines = [ line for line in lines if line.strip()]

    for i in range(0, len(lines), 3):
        block = lines[i] + lines[i+1] + lines[i+2]
        nums = list(map(int, re.findall(f"-?\\d+", block)))
        print(nums)
        Ax, Ay, Bx, By, Px, Py = nums

        if part2:
            Px += 10**13
            Py += 10**13
        D = Ax*By - Bx*Ay
        if D == 0:
            continue

        a_nums = Px*By - Py*Bx
        b_nums = Ax*Py - Ay*Px

        if a_nums % D != 0 or b_nums % D != 0:
            continue

        a = a_nums // D
        b = b_nums // D

        if a < 0 or b < 0:
            continue

        total_cost += 3*a + b
    return total_cost



if __name__ == "__main__":
    with open("day13_inputs.txt") as f:
        lines = f.readlines();
    
print("Part 1: ", solve(lines, False))
print("Part 2:", solve(lines, True))