from collections import defaultdict
from itertools import count

from input_parser import parse_day


TEST_DATA = """#.######
#>>.<^<#
#.<..<<#
#>v.><>#
#<^v^^>#
######.#""".split('\n')

TEST = False

gen = TEST_DATA if TEST else list(parse_day(24))

start = 0 + 1j
end = complex(len(gen)-1, len(gen[-1])-2)

min_minutes = 1000000000

directions = {'>': 1j, 'v': 1, '<': -1j, '^': -1}

init_state = defaultdict(list)
for r, line in enumerate(gen):
    for c, element in enumerate(line):
        if element == '.':
            continue
        init_state[complex(r, c)].append(element if element == '#'
                                         else directions[element])
states = {0: init_state}

denominator = (len(gen) - 2) * (len(gen[0]) - 2)
for i in range(1, denominator):
    state = defaultdict(list)
    old_state = states[i-1]
    for p, ds in dict(old_state).items():
        if ds == ['#']:
            state[p].append('#')
            continue
        for d in ds:
            new_pos = p + d
            if old_state.get(new_pos, None) == ['#']:
                # Luckily no verticals in start and end columns.
                new_pos = p
                while old_state.get(new_pos, None) != ['#']:
                    new_pos -= d
                new_pos += d
            state[new_pos].append(d)
    states[i] = state


def get_state(minutes):
    return states[minutes % denominator]


def bfs(start_position, end_position, start_time):
    curr_positions = {start}
    next_positions = set()
    for minutes in count(start_time):
        state = get_state(minutes+1)
        for curr_pos in curr_positions:
            if curr_pos == end:
                return minutes

            candidate_positions = (curr_pos + d for d in (1j, 1, -1j, -1, 0))
            for p in candidate_positions:
                if (p not in next_positions and p not in state
                        and 0 <= p.real < len(gen)):
                    next_positions.add(p)

        curr_positions = next_positions
        next_positions = set()
        continue


first_bfs_minutes = bfs(start, end, 0)
print("Part 1:", first_bfs_minutes)

second_bfs_minutes = bfs(end, start, first_bfs_minutes)
third_bfs_minutes = bfs(start, end, second_bfs_minutes)
print("Part 2:", third_bfs_minutes)
