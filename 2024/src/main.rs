mod day1;

use std::io::BufRead;

fn main() {
    // day1::run();

    let (part1, part2) =
        std::io::BufReader::new(std::fs::File::open("../inputs/day2.txt").unwrap())
            .lines()
            .map(|line| analyze_line(line.unwrap()))
            .fold((0, 0), |(part1, part2), (part1_safety, part2_safety)| {
                (part1 + part1_safety, part2 + part2_safety)
            });
    println!("Part 1: {part1}");
    println!("Part 2: {part2}");
}

fn analyze_line(line: String) -> (i32, i32) {
    let iter: LineSafetyIter<_> = line
        .as_str()
        .split(" ")
        .map(|num_str| num_str.parse::<i32>().unwrap())
        .into();
    match iter.last() {
        Some(LineSafety::Safe) => (1, 1),
        Some(LineSafety::Part2Safe) => (0, 1),
        Some(LineSafety::Unsafe) => (0, 0),
        None => panic!(),
    }
}

#[derive(Copy, Clone)]
enum LineSafety {
    Safe,
    Part2Safe,
    Unsafe,
}

enum Variance {
    Initialized,
    SeenOne(i32),
    SeenTwo(i32, i32),
    SeenThree(LevelDifference, LevelDifference, i32),
    Decreasing(i32, LineSafety),
    Increasing(i32, LineSafety),
    Unsafe,
}

#[derive(Copy, Clone, PartialEq)]
enum LevelDifference {
    Decreasing,
    Increasing,
    Invalid,
}

impl LevelDifference {
    fn new(curr: i32, next: i32) -> Self {
        let diff = next - curr;
        if -3 <= diff && diff <= -1 {
            LevelDifference::Decreasing
        } else if 1 <= diff && diff <= 3 {
            LevelDifference::Increasing
        } else {
            LevelDifference::Invalid
        }
    }

    fn into_variance(&self, next: i32, safety: LineSafety) -> Variance {
        match self {
            LevelDifference::Decreasing => Variance::Decreasing(next, safety),
            LevelDifference::Increasing => Variance::Increasing(next, safety),
            _ => panic!(),
        }
    }
}

struct LineSafetyIter<I> {
    variance: Variance,
    iter: I,
}

impl<I: Iterator<Item = i32>> Iterator for LineSafetyIter<I> {
    type Item = LineSafety;

    fn next(&mut self) -> Option<LineSafety> {
        match self.iter.next() {
            None => None,
            Some(next) => {
                let (variance, ret) = match self.variance {
                    Variance::Initialized => (Variance::SeenOne(next), LineSafety::Safe),
                    Variance::SeenOne(curr) => match LevelDifference::new(curr, next) {
                        LevelDifference::Decreasing => (
                            Variance::Decreasing(next, LineSafety::Safe),
                            LineSafety::Safe,
                        ),
                        LevelDifference::Increasing => (
                            Variance::Increasing(next, LineSafety::Safe),
                            LineSafety::Safe,
                        ),
                        LevelDifference::Invalid => {
                            (Variance::SeenTwo(curr, next), LineSafety::Part2Safe)
                        }
                    },
                    Variance::SeenTwo(l1, l2) => {
                        let diff1 = LevelDifference::new(l1, next);
                        let diff2 = LevelDifference::new(l2, next);
                        match (diff1, diff2) {
                            (LevelDifference::Invalid, LevelDifference::Invalid) => {
                                (Variance::Unsafe, LineSafety::Unsafe)
                            }
                            (l1_diff, LevelDifference::Invalid) => (
                                l1_diff.into_variance(next, LineSafety::Part2Safe),
                                LineSafety::Part2Safe,
                            ),
                            (LevelDifference::Invalid, l2_diff) => (
                                l2_diff.into_variance(next, LineSafety::Part2Safe),
                                LineSafety::Part2Safe,
                            ),
                            _ => (
                                Variance::SeenThree(diff1, diff2, next),
                                LineSafety::Part2Safe,
                            ),
                        }
                    }
                    Variance::SeenThree(diff1, diff2, l3) => {
                        let diff3 = LevelDifference::new(l3, next);
                        if diff3 == diff1 || diff3 == diff2 {
                            (
                                diff3.into_variance(next, LineSafety::Part2Safe),
                                LineSafety::Part2Safe,
                            )
                        } else {
                            (Variance::Unsafe, LineSafety::Unsafe)
                        }
                    }
                    Variance::Decreasing(curr, safety) => {
                        let diff = LevelDifference::new(curr, next);
                        match diff {
                            LevelDifference::Decreasing => {
                                (diff.into_variance(next, safety), safety)
                            }
                            _ => match safety {
                                LineSafety::Safe => (
                                    Variance::Decreasing(curr, LineSafety::Part2Safe),
                                    LineSafety::Part2Safe,
                                ),
                                LineSafety::Part2Safe => (Variance::Unsafe, LineSafety::Unsafe),
                                _ => panic!(),
                            },
                        }
                    }
                    Variance::Increasing(curr, safety) => {
                        let diff = LevelDifference::new(curr, next);
                        match diff {
                            LevelDifference::Increasing => {
                                (diff.into_variance(next, safety), safety)
                            }
                            _ => match safety {
                                LineSafety::Safe => (
                                    Variance::Increasing(curr, LineSafety::Part2Safe),
                                    LineSafety::Part2Safe,
                                ),
                                LineSafety::Part2Safe => (Variance::Unsafe, LineSafety::Unsafe),
                                _ => panic!(),
                            },
                        }
                    }
                    Variance::Unsafe => {
                        return None;
                    }
                };
                self.variance = variance;
                Some(ret)
            }
        }
    }
}

impl<I: Iterator<Item = i32>> From<I> for LineSafetyIter<I> {
    fn from(iter: I) -> Self {
        LineSafetyIter {
            variance: Variance::Initialized,
            iter,
        }
    }
}
