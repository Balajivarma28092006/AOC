use std::fs;
use std::io;

fn part1(contents: &str) -> i32 {
    let mut position: i32 = 50;
    let mut count = 0;

    for line in contents.lines() {
        if line.trim().is_empty() {
            continue;
        }

        let (dir, dist_str) = line.split_at(1);
        let dist: i32 = dist_str.parse().unwrap();

        if dir == "L" {
            position -= dist;
        } else {
            position += dist;
        }

        position = position.rem_euclid(100);

        if position == 0 {
            count += 1;
        }
    }

    count
}

fn part2(contents: &str) -> i32 {
    let mut position: i32 = 50;
    let mut count = 0;

    for line in contents.lines() {
        if line.trim().is_empty() {
            continue;
        }

        let (dir, dist_str) = line.split_at(1);
        let mut dist: i32 = dist_str.parse().unwrap();

        let left = dir == "L";

        while dist > 0 {
            if left {
                position -= 1;
            } else {
                position += 1;
            }

            if position == 100 {
                position = 0;
            }
            if position == -1 {
                position = 99;
            }

            if position == 0 {
                count += 1;
            }

            dist -= 1;
        }
    }

    count
}

fn main() -> io::Result<()> {
    let contents = fs::read_to_string("day1.txt").expect("something's missing");

    println!("Part1: {}", part1(&contents));
    println!("Part2: {}", part2(&contents));

    Ok(())
}
