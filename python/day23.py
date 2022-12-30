from collections import defaultdict
from itertools import count, cycle, islice, product

from input_parser import parse_day

TEST = False

TEST_DATA = """....#..
..###.#
#...#.#
.#...##
#.###..
##.#.##
.#..#..""".split('\n')

gen = iter(TEST_DATA) if TEST else parse_day(23)

elves = {complex(r, c) for r, row in enumerate(gen)
         for c, ele in enumerate(row) if ele == '#'}


def print_elves(elves):
    min_r = int(min(e.real for e in elves))
    max_r = int(max(e.real for e in elves))
    min_c = int(min(e.imag for e in elves))
    max_c = int(max(e.imag for e in elves))

    for r in range(min_r, max_r+1):
        for c in range(min_c, max_c+1):
            print('#' if complex(r, c) in elves else '.', end='')
        print()
    print()


directions = [
    (-1, (-1-1j, -1+0j, -1+1j)),
    (+1, (+1-1j, +1+0j, +1+1j)),
    (-1j, (-1-1j, +0-1j, +1-1j)),
    (+1j, (-1+1j, +0+1j, +1+1j)),
    ]
directions_iter = cycle(directions)

surrounding = [complex(i, j) for i, j in product((-1, 0, 1), repeat=2)]
surrounding = [x for x in surrounding if x != 0]

for i in count(1):
    moves = defaultdict(list)

    for elf in elves:
        if not any(elf + x in elves for x in surrounding):
            continue

        for direction, adjacents in list(islice(directions_iter, 4)):
            if not any(elf + x in elves for x in adjacents):
                moves[elf + direction].append(elf)
                break

    if not moves:
        break

    for position, orig_list in moves.items():
        if len(orig_list) == 1:
            elves.remove(orig_list[0])
            elves.add(position)

    # Change next direction.
    next(directions_iter)

    if i == 10:
        min_r = min(e.real for e in elves)
        max_r = max(e.real for e in elves)
        min_c = min(e.imag for e in elves)
        max_c = max(e.imag for e in elves)
        print("Part 1:", int((max_r - min_r + 1) * (max_c - min_c + 1)
                             - len(elves)))

print('Part 2:', i)
