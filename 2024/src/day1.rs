use std::io::BufRead;

pub fn run() {
    let (mut v1, mut v2): (Vec<_>, Vec<_>) =
        std::io::BufReader::new(std::fs::File::open("../inputs/day1.txt").unwrap())
            .lines()
            .map(|line| {
                let v = line
                    .unwrap()
                    .as_str()
                    .split_whitespace()
                    .map(|x| x.parse::<i32>().unwrap())
                    .collect::<Vec<_>>();
                (v[0], v[1])
            })
            .unzip();
    v1.sort_unstable();
    v2.sort_unstable();

    let part1: i32 = v1.iter().zip(v2.iter()).map(|(x, y)| (x - y).abs()).sum();
    println!("Part 1: {part1}");

    let mut v2_counter = std::collections::HashMap::new();
    for x in v2 {
        v2_counter.insert(
            x,
            match v2_counter.get(&x) {
                Some(y) => y + 1,
                None => 1,
            },
        );
    }
    let part2: i32 = v1
        .iter()
        .map(|x| {
            x * {
                match v2_counter.get(x) {
                    Some(y) => *y,
                    None => 0,
                }
            }
        })
        .sum();
    println!("Part 2: {part2}");
}
