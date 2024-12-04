mod day1;

use std::io::BufRead;

fn main() {
    // day1::run();

    let (part1, part2) =
        std::io::BufReader::new(std::fs::File::open("../inputs/day2.txt").unwrap())
            .lines()
            .map(|line| analyze_line(line.unwrap()))
            .fold((0, 0), |(part1, part2), (part1_safety, part2_safety)| {
                (
                    part1 + matches!(part1_safety, LineSafety::Safe) as i32,
                    part2 + matches!(part2_safety, LineSafety::Safe) as i32,
                )
            });
    println!("Part 1: {part1}");
    println!("Part 2: {part2}");
}

fn analyze_line(line: String) -> (LineSafety, LineSafety) {
    let iter: LineSafetyIter<_> = line
        .as_str()
        .split(" ")
        .map(|num_str| num_str.parse::<i32>().unwrap())
        .into();
    let unsafe_count = iter
        .filter(|safety| matches!(safety, LineSafety::Unsafe))
        .take(2)
        .count();
    match unsafe_count {
        0 => (LineSafety::Safe, LineSafety::Safe),
        1 => (LineSafety::Unsafe, LineSafety::Safe),
        2 => (LineSafety::Unsafe, LineSafety::Unsafe),
        _ => {
            panic!();
        }
    }
}

enum LineSafety {
    Safe,
    Unsafe,
}

enum Variance {
    Initialized,
    Increasing,
    Decreasing,
}

struct LineSafetyIter<I> {
    curr: Option<i32>,
    variance: Variance,
    iter: I,
}

impl<I: Iterator<Item = i32>> Iterator for LineSafetyIter<I> {
    type Item = LineSafety;

    fn next(&mut self) -> Option<LineSafety> {
        if self.curr.is_none() {
            self.curr = Some(self.iter.next().unwrap());
            return Some(LineSafety::Safe);
        }
        let curr = self.curr.unwrap();

        self.iter.next().map(|next| {
            let diff = next - curr;
            match self.variance {
                Variance::Initialized => {
                    self.variance = if -3 <= diff && diff <= -1 {
                        Variance::Decreasing
                    } else if 1 <= diff && diff <= 3 {
                        Variance::Increasing
                    } else {
                        return LineSafety::Unsafe;
                    };
                }
                Variance::Increasing => {
                    if diff < 1 || 3 < diff {
                        return LineSafety::Unsafe;
                    }
                }
                Variance::Decreasing => {
                    if diff < -3 || -1 < diff {
                        return LineSafety::Unsafe;
                    }
                }
            }
            self.curr = Some(next);

            LineSafety::Safe
        })
    }
}

impl<I: Iterator<Item = i32>> From<I> for LineSafetyIter<I> {
    fn from(iter: I) -> Self {
        LineSafetyIter {
            curr: None,
            variance: Variance::Initialized,
            iter,
        }
    }
}
