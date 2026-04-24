ranges = []
numbers = []

reading_ranges = True

# get the data
with open("inputs.txt") as f:
    for line in f:
        line = line.strip()

        if not line:
            reading_ranges = False
            continue

        if reading_ranges:
            s, e = map(int, line.split("-"))
            ranges.append((s, e))
        else:
            numbers.append(int(line))

# Part 1
count = 0

for n in numbers:
    for s, e in ranges:
        if s <= n <= e:
            count += 1
            break
print("Part 1: ", count)


# Part 2
ranges.sort()

merged = []
for s, e in ranges:
    if not merged or s > merged[-1][1]:
        merged.append([s, e])
    else:
        merged[-1][1] = max(merged[-1][1], e)

total = sum(e - s + 1 for s, e in merged)
print("Part 2: ", total)
