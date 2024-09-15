use std::io::prelude::*;
use std::sync::atomic::{AtomicBool, AtomicUsize, Ordering};
use std::sync::Mutex;

use md5::{Digest, Md5};

const CHUNK_SIZE: usize = 1_000;

struct Shared<'a> {
    prefix: &'a [u8],
    done: AtomicBool,
    counter: AtomicUsize,
    found: Mutex<Found>,
}

struct Found {
    entries: Vec<(usize, u8, u8)>,
    part2_mask: u8,
}

pub fn run(filename: &str) -> (String, String) {
    let mut file = std::fs::File::open(filename).unwrap();
    let mut buffer = String::new();
    file.read_to_string(&mut buffer).unwrap();
    let door_id = buffer.trim();

    let shared = Shared {
        prefix: &door_id.to_string().into_bytes(),
        done: AtomicBool::new(false),
        counter: AtomicUsize::new(0),
        found: Mutex::new(Found {
            entries: vec![],
            part2_mask: 0,
        }),
    };

    let num_threads = std::thread::available_parallelism().map_or(8, |n| n.get());
    std::thread::scope(|scope| {
        for _ in 0..num_threads {
            scope.spawn(|| worker(&shared));
        }
    });

    let found = shared.found.into_inner().unwrap();
    let mut result = found.entries;
    result.sort_unstable();

    let part1_v: Vec<u8> = result.iter().map(|(_, sixth, _)| *sixth).collect();
    let mut part2_v: [Option<u8>; 8] = [None; 8];
    for (_, sixth, seventh) in result {
        if sixth < 8 && part2_v[sixth as usize].is_none() {
            part2_v[sixth as usize] = Some(seventh);
        }
    }

    let part1: String = part1_v.iter().map(|b| format!("{:x}", b)).take(8).collect();
    let part2: String = part2_v
        .iter()
        .map(|b| format!("{:x}", b.unwrap()))
        .collect();

    (part1, part2)
}

fn worker(shared: &Shared) {
    while !shared.done.load(Ordering::Relaxed) {
        let start = shared.counter.fetch_add(CHUNK_SIZE, Ordering::Relaxed);
        let r = start..start + CHUNK_SIZE;
        for (sixth, seventh) in hash_range(r, shared.prefix) {
            let mut found = shared.found.lock().unwrap();

            found.entries.push((start, sixth, seventh));
            if sixth < 8 {
                found.part2_mask |= 1 << sixth;
            }

            if found.part2_mask == 0xff {
                shared.done.store(true, Ordering::Relaxed);
            }
        }
    }
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
