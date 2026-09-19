/* 0110 -- target
 * 0001
 * 0101
 * 0010
 * 0011
 * 1010
 * 1100
 * Think of toggling the bit what we use is OR but we should toggle things even when they
 * are the same means one thing which is XOR thats the main trick
 * assuming the given switches as binary positions and we need to find the things
 * that produce the target when XORed and they should be in less steps
 */

/*
 * But think of even more cooler way is to use BFS
 */

use std::{
    collections::{HashSet, VecDeque},
    fs,
};

fn main() {
    let input = fs::read_to_string("../test.txt").unwrap();
    let mut total = 0;

    for line in input.lines() {
        let parts: Vec<&str> = line.split_whitespace().collect();
        // First part is the target
        let target_str = parts[0];
        let target: Vec<bool> = target_str
            .trim_matches(['[', ']'])
            .chars()
            .map(|c| c == '#')
            .collect();
        // println!("{:?}", target)
        // basically converts # to true and . to false
        let mut buttons = Vec::new();
        for part in &parts[1..] {
            if part.starts_with('{') {
                break; // ignore the last {} for part 1
            }

            let nums = part
                .trim_matches(['(', ')'])
                .split(',')
                .map(|x| x.parse::<usize>().unwrap())
                .collect::<Vec<_>>();
            buttons.push(nums);
        }
        // println!("{:?}", buttons); get the values as an array
        let anwser = solve(&target, &buttons);
        println!("{} -> {}", line, anwser);
        total += anwser;
    }
    println!("Total: {}", total)
}

fn solve(target: &[bool], buttons: &[Vec<usize>]) -> usize {
    let start = vec![false; target.len()];
    println!("{:?}", start);
    let mut queue = VecDeque::new();
    queue.push_back((start.clone(), 0)); // current state and number of presses

    let mut visited = HashSet::new();
    visited.insert(start);

    while let Some((state, presses)) = queue.pop_front() {
        if state == target {
            return presses;
        }

        for button in buttons {
            let mut next = state.clone();
            for &index in button {
                next[index] = !next[index];
            }

            if visited.insert(next.clone()) {
                queue.push_back((next, presses + 1));
            }
        }
    }
    unreachable!("Target cannot be reached");
}
