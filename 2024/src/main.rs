mod day1;

use std::io::BufRead;

fn main() {
    // day1::run();

    let (part1, part2) =
        // std::io::BufReader::new(std::fs::File::open("../inputs/test.txt").unwrap())
        std::io::BufReader::new(std::fs::File::open("../inputs/day2.txt").unwrap())
            .lines()
            .map(|line| {
                analyze_line(line.unwrap())
            })
            .fold(
                (0, 0),
                |(part1, part2), (part1_safety, part2_safety)| {
                    (part1 + part1_safety, part2 + part2_safety)
                },
            );
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

#[derive(Copy, Clone, Debug)]
enum LineSafety {
    Safe,
    Part2Safe,
    Unsafe,
}

#[derive(Copy, Clone, Debug)]
enum LineSafetyIterState {
    Initialized,
    SeenOne(i32),
    SeenTwo(i32, i32),
    SeenThreeConflicting(i32, i32, i32),
    Decreasing(i32, i32, LineSafety),
    Increasing(i32, i32, LineSafety),
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
}

struct LineSafetyIter<I> {
    state: LineSafetyIterState,
    iter: I,
}

impl<I: Iterator<Item = i32>> Iterator for LineSafetyIter<I> {
    type Item = LineSafety;

    fn next(&mut self) -> Option<LineSafety> {
        match (self.iter.next(), self.state) {
            (None, _) | (_, LineSafetyIterState::Unsafe) => None,
            (Some(next), current_state) => match current_state {
                LineSafetyIterState::Initialized => {
                    self.state = LineSafetyIterState::SeenOne(next);
                    Some(LineSafety::Safe)
                }
                LineSafetyIterState::SeenOne(curr) => {
                    let diff = LevelDifference::new(curr, next);
                    let safety = match diff {
                        LevelDifference::Decreasing | LevelDifference::Increasing => {
                            LineSafety::Safe
                        }
                        _ => LineSafety::Part2Safe,
                    };
                    self.state = LineSafetyIterState::SeenTwo(curr, next);
                    Some(safety)
                }
                LineSafetyIterState::SeenTwo(prev, curr) => {
                    let diff1 = LevelDifference::new(prev, curr);
                    let diff2 = LevelDifference::new(curr, next);
                    let (state, safety) = match (diff1, diff2) {
                        (LevelDifference::Decreasing, LevelDifference::Decreasing) => (
                            LineSafetyIterState::Decreasing(curr, next, LineSafety::Safe),
                            LineSafety::Safe,
                        ),
                        (LevelDifference::Increasing, LevelDifference::Increasing) => (
                            LineSafetyIterState::Increasing(curr, next, LineSafety::Safe),
                            LineSafety::Safe,
                        ),
                        _ => (
                            LineSafetyIterState::SeenThreeConflicting(prev, curr, next),
                            LineSafety::Part2Safe,
                        ),
                    };
                    self.state = state;
                    Some(safety)
                }
                LineSafetyIterState::SeenThreeConflicting(p0, p1, p2) => {
                    match triple_safety(p1, p2, next) {
                        LevelDifference::Decreasing => {
                            self.state =
                                LineSafetyIterState::Decreasing(p2, next, LineSafety::Part2Safe);
                            return Some(LineSafety::Part2Safe);
                        }
                        LevelDifference::Increasing => {
                            self.state =
                                LineSafetyIterState::Increasing(p2, next, LineSafety::Part2Safe);
                            return Some(LineSafety::Part2Safe);
                        }
                        LevelDifference::Invalid => {}
                    }

                    match triple_safety(p0, p2, next) {
                        LevelDifference::Decreasing => {
                            self.state =
                                LineSafetyIterState::Decreasing(p2, next, LineSafety::Part2Safe);
                            return Some(LineSafety::Part2Safe);
                        }
                        LevelDifference::Increasing => {
                            self.state =
                                LineSafetyIterState::Increasing(p2, next, LineSafety::Part2Safe);
                            return Some(LineSafety::Part2Safe);
                        }
                        LevelDifference::Invalid => {}
                    }

                    match triple_safety(p0, p1, next) {
                        LevelDifference::Decreasing => {
                            self.state =
                                LineSafetyIterState::Decreasing(p1, next, LineSafety::Part2Safe);
                            return Some(LineSafety::Part2Safe);
                        }
                        LevelDifference::Increasing => {
                            self.state =
                                LineSafetyIterState::Increasing(p1, next, LineSafety::Part2Safe);
                            return Some(LineSafety::Part2Safe);
                        }
                        LevelDifference::Invalid => {}
                    }

                    self.state = LineSafetyIterState::Unsafe;
                    return Some(LineSafety::Unsafe);
                }
                LineSafetyIterState::Decreasing(prev, curr, current_safety) => {
                    let diff = LevelDifference::new(curr, next);
                    match diff {
                        LevelDifference::Decreasing => {
                            self.state =
                                LineSafetyIterState::Decreasing(curr, next, current_safety);
                            Some(current_safety)
                        }
                        _ => match current_safety {
                            LineSafety::Safe => {
                                let new_curr = match LevelDifference::new(prev, next) {
                                    LevelDifference::Decreasing => next,
                                    _ => curr,
                                };
                                self.state = LineSafetyIterState::Decreasing(
                                    prev,
                                    new_curr,
                                    LineSafety::Part2Safe,
                                );
                                Some(LineSafety::Part2Safe)
                            }
                            LineSafety::Part2Safe => {
                                self.state = LineSafetyIterState::Unsafe;
                                Some(LineSafety::Unsafe)
                            }
                            _ => panic!(),
                        },
                    }
                }
                LineSafetyIterState::Increasing(prev, curr, current_safety) => {
                    let diff = LevelDifference::new(curr, next);
                    match diff {
                        LevelDifference::Increasing => {
                            self.state =
                                LineSafetyIterState::Increasing(curr, next, current_safety);
                            Some(current_safety)
                        }
                        _ => match current_safety {
                            LineSafety::Safe => {
                                let new_curr = match LevelDifference::new(prev, next) {
                                    LevelDifference::Increasing => next,
                                    _ => curr,
                                };
                                self.state = LineSafetyIterState::Increasing(
                                    prev,
                                    new_curr,
                                    LineSafety::Part2Safe,
                                );
                                Some(LineSafety::Part2Safe)
                            }
                            LineSafety::Part2Safe => {
                                self.state = LineSafetyIterState::Unsafe;
                                Some(LineSafety::Unsafe)
                            }
                            _ => panic!(),
                        },
                    }
                }
                LineSafetyIterState::Unsafe => panic!(),
            },
        }
    }
}

impl<I: Iterator<Item = i32>> From<I> for LineSafetyIter<I> {
    fn from(iter: I) -> Self {
        LineSafetyIter {
            state: LineSafetyIterState::Initialized,
            iter,
        }
    }
}

fn triple_safety(x: i32, y: i32, z: i32) -> LevelDifference {
    let diff1 = LevelDifference::new(x, y);
    let diff2 = LevelDifference::new(y, z);
    match (diff1, diff2) {
        (LevelDifference::Increasing, LevelDifference::Increasing) => LevelDifference::Increasing,
        (LevelDifference::Decreasing, LevelDifference::Decreasing) => LevelDifference::Decreasing,
        _ => LevelDifference::Invalid,
    }
}
