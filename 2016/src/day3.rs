use std::fs::File;
use std::io::BufRead;

pub fn run(path: &str) -> (String, String) {
    let data = std::io::BufReader::new(File::open(path).unwrap())
        .lines()
        .map(|x| {
            let v = x
                .unwrap()
                .split_whitespace()
                .map(|y| y.parse::<usize>().unwrap())
                .collect::<Vec<usize>>();
            if let [x, y, z] = v[..] {
                (x, y, z)
            } else {
                panic!();
            }
        });
    let part1 = data
        .filter(|(x, y, z)| x + y > *z && x + z > *y && y + z > *x)
        .count();

    let data = std::io::BufReader::new(File::open(path).unwrap())
        .lines()
        .map(|x| {
            x.unwrap()
                .split_whitespace()
                .map(|y| y.parse::<usize>().unwrap())
                .collect::<Vec<usize>>()
        });
    let (_, part2) = data.fold((vec![], 0), |(mut v, c), x| {
        v.push(x);
        match v.len() {
            1 | 2 => (v, c),
            3 => {
                let d = (0..3)
                    .map(|col| (0..3).map(|row| v[row][col]).collect())
                    .map(|v2: Vec<usize>| {
                        if let [x, y, z] = v2[..] {
                            (x, y, z)
                        } else {
                            panic!();
                        }
                    })
                    .filter(|(x, y, z)| x + y > *z && x + z > *y && y + z > *x)
                    .count();
                (vec![], c + d)
            }
            _ => panic!(),
        }
    });

    (format!("{}", part1), format!("{}", part2))
}
