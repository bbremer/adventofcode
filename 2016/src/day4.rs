use std::collections::HashMap;
use std::io::BufRead;

pub fn run(path: &str) -> (String, String) {
    let file = std::fs::File::open(path).unwrap();
    let buf_reader = std::io::BufReader::new(file);

    let (part1, part2): (u32, Option<u32>) = buf_reader
        .lines()
        .map(|line| {
            let mut unwrapped_line = line.unwrap();
            unwrapped_line.pop(); // remove last
            let string_chars = unwrapped_line.split("[").collect::<Vec<&str>>();
            let chars = string_chars[1];

            let strings_number = string_chars[0].split("-").collect::<Vec<&str>>();

            let strings = &strings_number[..strings_number.len() - 1];
            let string = strings.join("");

            let number = strings_number.last().unwrap().parse::<u32>().unwrap();

            let mut map: HashMap<char, u8> = HashMap::with_capacity(string.len());

            for c in string.chars() {
                *(map.entry(c).or_insert(0)) += 1;
            }

            let mut heap = map
                .iter()
                .map(|(k, v)| (v, k))
                .collect::<Vec<(&u8, &char)>>();

            heap.sort_by(
                |(count_a, char_a), (count_b, char_b)| match count_b.cmp(count_a) {
                    std::cmp::Ordering::Equal => char_a.cmp(char_b),
                    x => x,
                },
            );

            let common_letters = heap
                .iter()
                .take(chars.len())
                .map(|(_, letter)| *letter)
                .collect::<String>();

            if common_letters == chars {
                let decrypted = strings
                    .iter()
                    .map(|s| {
                        s.chars()
                            .map(|c| {
                                std::char::from_u32((c as u32 - 97 + number) % 26 + 97).unwrap()
                            })
                            .collect::<String>()
                    })
                    .collect::<Vec<String>>()
                    .join(" ");
                if decrypted == "northpole object storage" {
                    (number, Some(number))
                } else {
                    (number, None)
                }
            } else {
                (0, None)
            }
        })
        .fold((0, None), |(sum, part2_found), (number, part2)| {
            (
                sum + number,
                if let Some(x) = part2 {
                    if let Some(_) = part2_found {
                        panic!();
                    }
                    Some(x)
                } else {
                    part2_found
                },
            )
        });

    (part1.to_string(), part2.unwrap().to_string())
}
