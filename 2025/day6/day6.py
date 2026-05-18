from math import prod

with open("inputs.txt") as f:
    lines = [line.strip() for line in f if line.strip()]

symbols = lines[-1].split()
matrix = [list(map(int, line.split())) for line in lines[:-1]]
mappings = {}

for col in range(len(symbols)):
    symbol = symbols[col]
    values = [row[col] for row in matrix]

    if symbol not in mappings:
        mappings[symbol] = []

    mappings[symbol].append(values)

part_total1 = 0

for symbol, value_lists in mappings.items():
    for value in value_lists:
        if symbol == "+":
            part_total1 += sum(value)
        elif symbol == "*":
            part_total1+= prod(value)
        else:
            raise ValueError(f"Unknown symbol: {symbol}")

print(part_total1)