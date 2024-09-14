use std::io::prelude::*;

use md5::{Digest, Md5};

pub fn run(filename: &str) -> (String, String) {
    let mut file = std::fs::File::open(filename).unwrap();
    let mut buffer = String::new();
    file.read_to_string(&mut buffer).unwrap();
    let door_id = buffer.trim();

    const CHUNK_SIZE: usize = 50_000;
    let mut chunker = (0..).step_by(CHUNK_SIZE).map(|i| i..i + CHUNK_SIZE);

    let call = |r| {
        let v = door_id.to_string().into_bytes();
        std::thread::spawn(move || hash_range(r, &v))
    };

    let num_threads = std::thread::available_parallelism().map_or(8, |n| n.get());
    let mut join_handles: std::collections::VecDeque<std::thread::JoinHandle<_>> = chunker
        .by_ref()
        .take(num_threads)
        .map(|r| call(r))
        .collect();

    let mut part1_v = Vec::new();
    let mut part2_v: [Option<u8>; 8] = [None; 8];
    let valid_sixths_and_sevenths = chunker.flat_map(|r| {
        join_handles.push_back(call(r));
        join_handles.pop_front().unwrap().join().unwrap()
    });

    for (sixth, seventh) in valid_sixths_and_sevenths {
        if part1_v.len() < 8 {
            part1_v.push(sixth)
        }

        if sixth < 8 && part2_v[sixth as usize].is_none() {
            part2_v[sixth as usize] = Some(seventh);
            if part2_v.into_iter().all(|x| x.is_some()) {
                break;
            }
        }
    }

    let part1: String = part1_v.iter().map(|b| format!("{:x}", b)).collect();
    let part2: String = part2_v
        .iter()
        .map(|b| format!("{:x}", b.unwrap()))
        .collect();

    (part1, part2)
}

fn hash_range(r: std::ops::Range<usize>, value: &[u8]) -> Vec<(u8, u8)> {
    let value_size = value.len();
    let num_zeros = r.start.to_string().as_bytes().len() - 1;
    let mut input: Vec<u8> = [value, r.start.to_string().as_bytes()].concat();

    r.filter_map(|i| {
        // Manually format string for performance. Update the digits in reverse.
        for j in (1..num_zeros).rev() {
            let a = 10_usize.pow((num_zeros - j) as u32);
            let index = value_size + j;
            let new_val = b'0' + (i / a % 10) as u8;
            if input[index] == new_val {
                break;
            }
            input[index] = new_val;
        }
        input[value_size + num_zeros] = b'0' + (i % 10) as u8;
        let result = Md5::digest(&input);

        if result[0..2] == [0; 2] && result[2] < 16 {
            Some((result[2], result[3] >> 4))
        } else {
            None
        }
    })
    .collect()
}
