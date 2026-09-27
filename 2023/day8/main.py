mp = {} # string to tuple(string, string) represents left and right
with open("inputs.txt", "r", encoding="utf-8") as file:
    data = file.readlines()

directions = data[0].strip()

for i in range(1, len(data)):
    line = data[i].strip()
    if not line:
        continue
    if "=" not in line:
        continue

    parts = line.split("=")
    key = parts[0].strip()

    dirs = [d.strip() for d in parts[1].strip().strip("()").split(",")]
    mp[key] = dirs

count = 0
curr = "AAA"

i = 0
while True:
    dr = directions[i % len(directions)]
    if curr == "ZZZ":
        print(count)
        break
    elif dr == 'L':
        count += 1
        child = mp[curr]
        curr = child[0]
    else:
        count += 1
        child = mp[curr]
        curr = child[1]
    i += 1
