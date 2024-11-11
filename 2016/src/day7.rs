use std::io::BufRead;

pub fn run(filename: &str) -> (String, String) {
    let part1 = std::io::BufReader::new(std::fs::File::open(filename).unwrap())
        .lines()
        .map(|line| line.unwrap())
        .filter(|line| valid(&line))
        .count()
        .to_string();
    (part1, "world".to_string())
}

fn valid(line: &str) -> bool {
    let substrings = IPSubstrings::new(line);
    let mut valid_found = false;
    for (substring, state) in substrings {
        match state {
            SquareBracketPosition::Inside => {
                if has_abba(&substring) {
                    return false;
                };
            }
            SquareBracketPosition::Outside => {
                if !valid_found && has_abba(&substring) {
                    valid_found = true;
                }
            }
        }
    }
    valid_found
}

enum SquareBracketPosition {
    Outside,
    Inside,
}

struct IPSubstrings<'a> {
    bytes: std::str::Bytes<'a>,
    state: SquareBracketPosition,
}

impl<'a> IPSubstrings<'a> {
    fn new(ip: &'a str) -> Self {
        Self {
            bytes: ip.bytes(),
            state: SquareBracketPosition::Outside,
        }
    }
}

impl<'a> Iterator for IPSubstrings<'a> {
    type Item = (String, SquareBracketPosition);

    fn next(&mut self) -> Option<Self::Item> {
        let mut substring_vec = Vec::new();

        for b in self.bytes.by_ref() {
            match (b, &self.state) {
                (b'[', SquareBracketPosition::Outside) => {
                    self.state = SquareBracketPosition::Inside;
                    return Some((
                        String::from_utf8(substring_vec).unwrap(),
                        SquareBracketPosition::Outside,
                    ));
                }
                (b']', SquareBracketPosition::Inside) => {
                    self.state = SquareBracketPosition::Outside;
                    return Some((
                        String::from_utf8(substring_vec).unwrap(),
                        SquareBracketPosition::Inside,
                    ));
                }
                (b'[', SquareBracketPosition::Inside) | (b']', SquareBracketPosition::Outside) => {
                    panic!();
                }
                _ => {
                    substring_vec.push(b);
                }
            }
        }

        if substring_vec.len() > 0 {
            Some((
                String::from_utf8(substring_vec).unwrap(),
                match self.state {
                    SquareBracketPosition::Outside => SquareBracketPosition::Outside,
                    SquareBracketPosition::Inside => panic!(),
                },
            ))
        } else {
            None
        }
    }
}

fn has_abba(substring: &str) -> bool {
    AbbaFinder::new(substring).any(|x| x)
}

struct AbbaFinder<'a> {
    bytes: std::str::Bytes<'a>,
    prev: std::collections::VecDeque<u8>,
}

impl<'a> AbbaFinder<'a> {
    fn new(string: &'a str) -> Self {
        let mut bytes = string.bytes();
        let prev: std::collections::VecDeque<u8> = bytes.by_ref().take(3).collect();
        assert_eq!(prev.len(), 3);
        AbbaFinder { bytes, prev }
    }
}

impl<'a> Iterator for AbbaFinder<'a> {
    type Item = bool;

    fn next(&mut self) -> Option<Self::Item> {
        let next2 = self.bytes.next();
        let next = match next2 {
            Some(c) => c,
            None => return None,
        };

        self.prev.make_contiguous();
        if let ([a, b, c], &[]) = &self.prev.as_slices() {
            let ret = *a != *b && *a == next && *b == *c;
            self.prev.pop_front();
            self.prev.push_back(next);
            Some(ret)
        } else {
            panic!();
        }
    }
}
