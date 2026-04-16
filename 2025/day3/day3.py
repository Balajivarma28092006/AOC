def best_number(line, k):
    digits = [int(ch) for ch in line if ch.isdigit()]
    drop = len(digits) - k
    stack = []

    for d in digits:
        while drop > 0 and stack and stack[-1] < d:
            stack.pop()
            drop -= 1
        stack.append(d)

    # take first k digits
    stack = stack[:k]

    val = 0
    for d in stack:
        val = val * 10 + d
    return val


p1 = p2 = 0

with open("inputs.txt") as f:
    for line in f:
        line = line.strip()
        p1 += best_number(line, 2)
        p2 += best_number(line, 12)

print("Part 1:", p1)
print("Part 2:", p2)