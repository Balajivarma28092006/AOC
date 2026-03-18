import re

WIDTH = 101
HEIGHT = 103

def parse_input(filename):
    robots = []
    with open(filename) as f:
        for line in f:
            nums = list(map(int, re.findall(r"-?\d+", line)))
            x, y, vx, vy = nums
            robots.append((x, y, vx, vy))
    return robots

def part1(robots):
    q1 = q2 = q3 = q4 = 0

    for x, y, vx, vy in robots:
        nx = (x + vx * 100) % WIDTH
        ny = (y + vy * 100) % HEIGHT

        #ignore center lines
        if nx == WIDTH // 2 and ny == HEIGHT // 2:
            continue

        if nx < WIDTH // 2 and ny < HEIGHT // 2:
            q1 += 1
        if nx > WIDTH // 2 and ny < HEIGHT // 2:
            q2 += 1
        if nx < WIDTH // 2 and ny > HEIGHT // 2:
            q3 += 1
        if nx > WIDTH // 2 and ny > HEIGHT // 2:
            q4 += 1

    return q1 * q2 * q3 * q4

def part2(robots):

    t = 0
    while True:
        t += 1
        positions = set()
        valid = True
        for x, y, vx, vy in robots:
            # Move robot t seconds and wrap around
            nx = (x + vx * t) % WIDTH
            ny = (y + vy * t) % HEIGHT

            # Check if two robots occupy the same spot
            if (nx, ny) in positions:
                valid = False
                break
            positions.add((nx, ny))

        if valid:
            return t


if __name__ == "__main__":
    inputs = parse_input("day14_inputs.txt")
    print(part1(inputs))
    print(part2(inputs))
