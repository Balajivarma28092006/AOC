import gleam/io
import gleam/int
import gleam/list
import gleam/string
import simplifile

pub fn main() {
  let assert Ok(input) = simplifile.read("input.txt")
  let lines = input |> string.trim |> string.split("\n")

  let part1 = lines |> list.map(best_number(_, 2)) |> int.sum
  let part2 = lines |> list.map(best_number(_, 12)) |> int.sum

  io.println("Part 1: " <> int.to_string(part1))
  io.println("Part 2: " <> int.to_string(part2))
}

fn best_number(line: String, count: Int) -> Int {
  let digits =
    line
    |> string.to_graphemes
    |> list.filter_map(int.parse)
    |> list.index_map(fn(d, i) { #(i, d) })

  combos(digits, count)
  |> list.map(fn(chosen) {
    list.fold(chosen, 0, fn(acc, pair) { acc * 10 + pair.1 })
  })
  |> list.fold(0, int.max)
}

fn combos(items: List(#(Int, Int)), k: Int) -> List(List(#(Int, Int))) {
  case k {
    0 -> [[]]
    _ ->
      case items {
        [] -> []
        [x, ..rest] -> {
          let with_x = combos(rest, k - 1) |> list.map(fn(c) { [x, ..c] })
          list.append(with_x, combos(rest, k))
        }
      }
  }
}