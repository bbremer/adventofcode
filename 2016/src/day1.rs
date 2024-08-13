#[derive(Clone, Copy)]
enum Direction {
    North,
    East,
    South,
    West,
}

impl Direction {
    fn turn(&self, side: &str) -> Self {
        match (self, side) {
            (Direction::East, "L") | (Direction::West, "R") => Direction::North,
            (Direction::South, "L") | (Direction::North, "R") => Direction::East,
            (Direction::West, "L") | (Direction::East, "R") => Direction::South,
            (Direction::North, "L") | (Direction::South, "R") => Direction::West,
            _ => panic!(),
        }
    }

    fn travel(&self, x: i32, y: i32) -> (i32, i32) {
        match self {
            Direction::North => (x, y + 1),
            Direction::East => (x + 1, y),
            Direction::South => (x, y - 1),
            Direction::West => (x - 1, y),
        }
    }
}

pub fn run(input_path: &str) -> (String, String) {
    let ((x, y), _, part2): ((i32, i32), Vec<(i32, i32)>, Option<i32>) =
        std::fs::read_to_string(input_path)
            .unwrap()
            .trim()
            .split(", ")
            .scan(Direction::North, |state, x| {
                let side = &x[..1];
                *state = state.turn(side);
                let n = (&x[1..]).parse::<usize>().unwrap();
                Some(std::iter::repeat(*state).take(n))
            })
            .flatten()
            .fold(((0, 0), Vec::new(), None), |((x, y), mut v, part2), d| {
                let (new_x, new_y) = d.travel(x, y);
                let new_part2 = match part2 {
                    Some(_) => part2,
                    None => {
                        if v.contains(&(new_x, new_y)) {
                            Some(new_x.abs() + new_y.abs())
                        } else {
                            v.push((x, y));
                            None
                        }
                    }
                };
                ((new_x, new_y), v, new_part2)
            });

    let part1 = x.abs() + y.abs();
    (format!("{}", part1), format!("{}", part2.unwrap()))
}
