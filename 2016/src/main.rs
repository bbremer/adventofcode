mod day1;
mod day2;
mod day3;
mod day4;
mod day5;
mod day6;

fn run_day(n: u8, f: fn(&str) -> (String, String)) {
    let now = std::time::Instant::now();
    let path = &format!("../../inputs/day{}.txt", n);
    let (part1, part2) = f(path);
    println!("Day {}: {} {} {:.2?}", n, part1, part2, now.elapsed());
}

fn main() {
    run_day(1, day1::run);
    run_day(2, day2::run);
    run_day(3, day3::run);
    run_day(4, day4::run);
    run_day(5, day5::run);
    run_day(6, day6::run);
}
