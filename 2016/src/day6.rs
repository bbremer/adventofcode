use std::io::BufRead;

pub fn run(filename: &str) -> (String, String) {
    let f = std::fs::File::open(filename).unwrap();
    let mut counters: [[u8; 26]; 8] = [[0; 26]; 8];
    let lines = std::io::BufReader::new(f).lines();
    for line in lines {
        for (i, c) in line.unwrap().bytes().enumerate() {
            counters[i][(c - b'a') as usize] += 1;
        }
    }

    let mut part1: Vec<u8> = Vec::with_capacity(8);
    for counter in counters.iter() {
        let max_index = counter
            .iter()
            .enumerate()
            .max_by(|(_, a), (_, b)| a.cmp(b))
            .map(|(index, _)| index as u8)
            .unwrap();
        part1.push(max_index + b'a');
    }

    let part2: Vec<u8> = counters
        .iter()
        .map(|counter| {
            counter
                .iter()
                .enumerate()
                .min_by(|(_, a), (_, b)| a.cmp(b))
                .map(|(index, _)| index as u8)
                .unwrap()
                + b'a'
        })
        .collect();
    (
        String::from_utf8(part1).unwrap(),
        String::from_utf8(part2).unwrap(),
    )
}
