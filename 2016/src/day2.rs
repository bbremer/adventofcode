use std::fs::File;
use std::io::BufRead;

enum Direction {
    Up,
    Down,
    Left,
    Right,
}

impl Direction {
    fn new(c: char) -> Self {
        match c {
            'U' => Direction::Up,
            'D' => Direction::Down,
            'L' => Direction::Left,
            'R' => Direction::Right,
            _ => panic!(),
        }
    }
}

#[derive(Clone, Copy)]
enum Part1Key {
    One,
    Two,
    Three,
    Four,
    Five,
    Six,
    Seven,
    Eight,
    Nine,
}

impl Part1Key {
    fn travel(&self, direction: Direction) -> Self {
        match (self, direction) {
            (Part1Key::One, Direction::Up) => Part1Key::One,
            (Part1Key::One, Direction::Down) => Part1Key::Four,
            (Part1Key::One, Direction::Left) => Part1Key::One,
            (Part1Key::One, Direction::Right) => Part1Key::Two,
            (Part1Key::Two, Direction::Up) => Part1Key::Two,
            (Part1Key::Two, Direction::Down) => Part1Key::Five,
            (Part1Key::Two, Direction::Left) => Part1Key::One,
            (Part1Key::Two, Direction::Right) => Part1Key::Three,
            (Part1Key::Three, Direction::Up) => Part1Key::Three,
            (Part1Key::Three, Direction::Down) => Part1Key::Six,
            (Part1Key::Three, Direction::Left) => Part1Key::Two,
            (Part1Key::Three, Direction::Right) => Part1Key::Three,
            (Part1Key::Four, Direction::Up) => Part1Key::One,
            (Part1Key::Four, Direction::Down) => Part1Key::Seven,
            (Part1Key::Four, Direction::Left) => Part1Key::Four,
            (Part1Key::Four, Direction::Right) => Part1Key::Five,
            (Part1Key::Five, Direction::Up) => Part1Key::Two,
            (Part1Key::Five, Direction::Down) => Part1Key::Eight,
            (Part1Key::Five, Direction::Left) => Part1Key::Four,
            (Part1Key::Five, Direction::Right) => Part1Key::Six,
            (Part1Key::Six, Direction::Up) => Part1Key::Three,
            (Part1Key::Six, Direction::Down) => Part1Key::Nine,
            (Part1Key::Six, Direction::Left) => Part1Key::Five,
            (Part1Key::Six, Direction::Right) => Part1Key::Six,
            (Part1Key::Seven, Direction::Up) => Part1Key::Four,
            (Part1Key::Seven, Direction::Down) => Part1Key::Seven,
            (Part1Key::Seven, Direction::Left) => Part1Key::Seven,
            (Part1Key::Seven, Direction::Right) => Part1Key::Eight,
            (Part1Key::Eight, Direction::Up) => Part1Key::Five,
            (Part1Key::Eight, Direction::Down) => Part1Key::Eight,
            (Part1Key::Eight, Direction::Left) => Part1Key::Seven,
            (Part1Key::Eight, Direction::Right) => Part1Key::Nine,
            (Part1Key::Nine, Direction::Up) => Part1Key::Six,
            (Part1Key::Nine, Direction::Down) => Part1Key::Nine,
            (Part1Key::Nine, Direction::Left) => Part1Key::Eight,
            (Part1Key::Nine, Direction::Right) => Part1Key::Nine,
        }
    }

    fn value(&self) -> String {
        (match self {
            Part1Key::One => "1",
            Part1Key::Two => "2",
            Part1Key::Three => "3",
            Part1Key::Four => "4",
            Part1Key::Five => "5",
            Part1Key::Six => "6",
            Part1Key::Seven => "7",
            Part1Key::Eight => "8",
            Part1Key::Nine => "9",
        })
        .to_string()
    }
}

#[derive(Clone, Copy)]
enum Part2Key {
    One,
    Two,
    Three,
    Four,
    Five,
    Six,
    Seven,
    Eight,
    Nine,
    A,
    B,
    C,
    D,
}

impl Part2Key {
    fn travel(&self, direction: Direction) -> Self {
        match (self, direction) {
            (Part2Key::One, Direction::Up) => Part2Key::One,
            (Part2Key::One, Direction::Down) => Part2Key::Three,
            (Part2Key::One, Direction::Left) => Part2Key::One,
            (Part2Key::One, Direction::Right) => Part2Key::One,
            (Part2Key::Two, Direction::Up) => Part2Key::Two,
            (Part2Key::Two, Direction::Down) => Part2Key::Six,
            (Part2Key::Two, Direction::Left) => Part2Key::Two,
            (Part2Key::Two, Direction::Right) => Part2Key::Three,
            (Part2Key::Three, Direction::Up) => Part2Key::One,
            (Part2Key::Three, Direction::Down) => Part2Key::Seven,
            (Part2Key::Three, Direction::Left) => Part2Key::Two,
            (Part2Key::Three, Direction::Right) => Part2Key::Four,
            (Part2Key::Four, Direction::Up) => Part2Key::Four,
            (Part2Key::Four, Direction::Down) => Part2Key::Eight,
            (Part2Key::Four, Direction::Left) => Part2Key::Three,
            (Part2Key::Four, Direction::Right) => Part2Key::Four,
            (Part2Key::Five, Direction::Up) => Part2Key::Five,
            (Part2Key::Five, Direction::Down) => Part2Key::Five,
            (Part2Key::Five, Direction::Left) => Part2Key::Five,
            (Part2Key::Five, Direction::Right) => Part2Key::Six,
            (Part2Key::Six, Direction::Up) => Part2Key::Two,
            (Part2Key::Six, Direction::Down) => Part2Key::A,
            (Part2Key::Six, Direction::Left) => Part2Key::Five,
            (Part2Key::Six, Direction::Right) => Part2Key::Seven,
            (Part2Key::Seven, Direction::Up) => Part2Key::Three,
            (Part2Key::Seven, Direction::Down) => Part2Key::B,
            (Part2Key::Seven, Direction::Left) => Part2Key::Six,
            (Part2Key::Seven, Direction::Right) => Part2Key::Eight,
            (Part2Key::Eight, Direction::Up) => Part2Key::Four,
            (Part2Key::Eight, Direction::Down) => Part2Key::C,
            (Part2Key::Eight, Direction::Left) => Part2Key::Seven,
            (Part2Key::Eight, Direction::Right) => Part2Key::Nine,
            (Part2Key::Nine, Direction::Up) => Part2Key::Nine,
            (Part2Key::Nine, Direction::Down) => Part2Key::Nine,
            (Part2Key::Nine, Direction::Left) => Part2Key::Eight,
            (Part2Key::Nine, Direction::Right) => Part2Key::Nine,
            (Part2Key::A, Direction::Up) => Part2Key::Six,
            (Part2Key::A, Direction::Down) => Part2Key::A,
            (Part2Key::A, Direction::Left) => Part2Key::A,
            (Part2Key::A, Direction::Right) => Part2Key::B,
            (Part2Key::B, Direction::Up) => Part2Key::Seven,
            (Part2Key::B, Direction::Down) => Part2Key::D,
            (Part2Key::B, Direction::Left) => Part2Key::A,
            (Part2Key::B, Direction::Right) => Part2Key::C,
            (Part2Key::C, Direction::Up) => Part2Key::Eight,
            (Part2Key::C, Direction::Down) => Part2Key::C,
            (Part2Key::C, Direction::Left) => Part2Key::B,
            (Part2Key::C, Direction::Right) => Part2Key::C,
            (Part2Key::D, Direction::Up) => Part2Key::B,
            (Part2Key::D, Direction::Down) => Part2Key::D,
            (Part2Key::D, Direction::Left) => Part2Key::D,
            (Part2Key::D, Direction::Right) => Part2Key::D,
        }
    }

    fn value(&self) -> String {
        (match self {
            Part2Key::One => "1",
            Part2Key::Two => "2",
            Part2Key::Three => "3",
            Part2Key::Four => "4",
            Part2Key::Five => "5",
            Part2Key::Six => "6",
            Part2Key::Seven => "7",
            Part2Key::Eight => "8",
            Part2Key::Nine => "9",
            Part2Key::A => "A",
            Part2Key::B => "B",
            Part2Key::C => "C",
            Part2Key::D => "D",
        })
        .to_string()
    }
}
pub fn run(path: &str) -> (String, String) {
    let data = std::io::BufReader::new(File::open(path).unwrap())
        .lines()
        .map(|x| x.unwrap());
    let part1 = data
        .scan(Part1Key::Five, |state, x| {
            *state = x.chars().fold(*state, |y, d| y.travel(Direction::new(d)));
            Some((*state).value())
        })
        .collect::<String>();

    let data = std::io::BufReader::new(File::open(path).unwrap())
        .lines()
        .map(|x| x.unwrap());
    let part2 = data
        .scan(Part2Key::Five, |state, x| {
            *state = x.chars().fold(*state, |y, d| y.travel(Direction::new(d)));
            Some((*state).value())
        })
        .collect::<String>();

    (part1, part2)
}
