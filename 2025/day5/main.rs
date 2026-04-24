use std::cmp::max;
use std::fs;

fn merge_intervals(mut ranges: Vec<(i64, i64)>) -> Vec<(i64, i64)> {
    ranges.sort_by_key(|r| r.0);

    let mut merged = Vec::new();

    for (start, end) in ranges {
        if merged.is_empty() {
            merged.push((start, end));
        } else {
            let last = merged.last_mut().unwrap();
            if start <= last.1 {
                last.1 = max(end, last.1);
            } else {
                merged.push((start, end));
            }
        }
    }
    merged
}

fn main() {
    let content = fs::read_to_string("inputs.txt").unwrap();

    // two mutable dynamic arrays
    let mut ranges = Vec::new();
    let mut numbers = Vec::new();

    // distinguish between ranges and numbers
    let mut reading_ranges = true;

    for line in content.lines() {
        if line.trim().is_empty() {
            reading_ranges = false;
            continue;
        }

        if reading_ranges {
            let parts: Vec<i64> = line.split('-').map(|x| x.parse().unwrap()).collect();

            ranges.push((parts[0], parts[1]));
        } else {
            numbers.push(line.parse::<i64>().unwrap());
        }
    }

    let mut count = 0;

    for n in numbers {
        for (start, end) in &ranges {
            if n >= *start && n <= *end {
                count += 1;
                break;
            }
        }
    }

    let merged = merge_intervals(ranges);

    let fresh_things: i64 = merged.iter().map(|(s, e)| e - s + 1).sum();

    println!("Part 1 = {}", count);
    println!("Part 2 = {}", fresh_things)
}
