games_mapper = {}
with open("inputs.txt", "r") as file:
    for line in file:
        if not line.strip():
            continue
        
        game_title, all_reveals = line.split(":", 1)

        game_id = int(game_title.split()[1])
        rounds = all_reveals.split(";")

        game_lists = []
        for r in rounds:
            round_dict = {}
            items = r.split(",")
            for item in items:
                count, color = item.strip().split()
                round_dict[color] = count
            game_lists.append(round_dict)
        games_mapper[game_id] = game_lists

red = 12
green = 13
blue = 14

ans = 0

for game_id, rounds in games_mapper.items():
    can_do = True
    for round in rounds:
        r, b, g = 0, 0, 0
        for color, count in round.items():
            if color == "blue":
                b = b + int(count)
            elif color == "red":
                r = r + int(count)
            elif color == "green":
                g = g + int(count)

            if not(r <= red and b <= blue and g <= green):
                can_do = False

    if can_do:
        ans = ans + game_id

print(ans)
                
ans1 = 0
for game_id, rounds in games_mapper.items():
    max_r, max_b, max_g = 0, 0, 0
    for round in rounds:
        for color, count in round.items():
            count = int(count)
            if color == "red":
                max_r = max(max_r, count)
            elif color == "green":
                max_g = max(max_g, count)
            elif color == "blue":
                max_b = max(max_b, count)

    ans1 += max_r * max_g * max_b
        
print(ans1)

