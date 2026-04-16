import strutils, sequtils, math

proc combos(digits: seq[int], k: int): seq[seq[int]] =
  if k == 0: return @[@[]]
  if digits.len == 0: return @[]
  let rest = digits[1..^1]
  for c in combos(rest, k - 1): result.add(@[digits[0]] & c)
  result.add combos(rest, k)

proc bestNumber(line: string, count: int): int =
  let digits = line.toSeq.filterIt(it.isDigit).mapIt(it.ord - '0'.ord)
  for combo in combos(digits, count):
    let val = combo.foldl(a * 10 + b, 0)
    result = max(result, val)

let lines = readFile("inputs.txt").strip.splitLines
echo "Part 1: ", lines.mapIt(bestNumber(it, 2)).sum
echo "Part 2: ", lines.mapIt(bestNumber(it, 12)).sum