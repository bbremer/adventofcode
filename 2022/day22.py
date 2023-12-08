from itertools import groupby, pairwise, product, takewhile

from input_parser import parse_day

test_data = """        ...#
        .#..
        #...
        ....
...#.......#
........#...
..#....#....
..........#.
        ...#....
        .....#..
        .#......
        ......#.

10R5L5R10L4R5L5""".split('\n')

TEST = False

gen = iter(test_data) if TEST else parse_day(22, False)

board = list(takewhile(bool, gen))
instructions = next(gen)

directions = 'eswn'
right_turn = dict(pairwise(directions[-1] + directions))
left_turn = dict(pairwise(directions[0] + directions[::-1]))

max_r = len(board)-1
max_c = max(len(row) for row in board)-1
board = [row.ljust(max_c+1) for row in board]

facing = 'e'
start_r = 0
start_c = next(i for i, x in enumerate(board[0]) if x == '.')

funcs = {
        'n': lambda r, c: (r-1, c),
        'e': lambda r, c: (r, c+1),
        's': lambda r, c: (r+1, c),
        'w': lambda r, c: (r, c-1),
        }


def wrap_p1(direction, r, c):
    match direction:
        case 'n':
            r = max_r
            while board[r][c] == ' ':
                r -= 1
        case 'e':
            c = 0
            while board[r][c] == ' ':
                c += 1
        case 's':
            r = 0
            while board[r][c] == ' ':
                r += 1
        case 'w':
            c = max_c
            while board[r][c] == ' ':
                c -= 1
    return direction, r, c


def move_forward(direction, r, c, wrap_f):
    next_direction = direction
    next_r, next_c = funcs[direction](r, c)
    if not (0 <= next_r <= max_r and 0 <= next_c <= max_c
            and board[next_r][next_c] != ' '):
        next_direction, next_r, next_c = wrap_f(direction, r, c)

    if board[next_r][next_c] == '#':
        return direction, r, c

    return next_direction, next_r, next_c


def run(wrap, r, c, facing):
    for numeric, g in groupby(instructions, str.isdigit):
        instruction = ''.join(g)
        if numeric:
            val = int(instruction)
            for i in range(val):
                facing, r, c = move_forward(facing, r, c, wrap)
        else:
            if len(instruction) > 1:
                raise ValueError(g)
            if instruction == 'R':
                facing = right_turn[facing]
            else:
                facing = left_turn[facing]

    return r, c


side_len = 4 if TEST else 50


final_r, final_c = run(wrap_p1, start_r, start_c, facing)
print("Part 1:", 1000*(final_r+1) + 4*(final_c+1) + directions.index(facing))


gen = iter(test_data) if TEST else parse_day(22, False)
board2 = {}
for row, line in enumerate(takewhile(bool, gen)):
    for column, element in enumerate(line):
        if element != ' ':
            board2[complex(row, column)] = element

boundaries = {}
for position in board2:
    if sum(position + complex(dr, dc) in board2 for dr, dc
           in product((-1, 0, 1), repeat=2)) != 8:
        continue

    (dr, dc), = [(dr, dc) for dr, dc in product((-1, 1), repeat=2)
                 if position+complex(dr, dc) not in board2]
    direction1 = normal2 = complex(dr, 0)
    direction2 = normal1 = complex(0, dc)
    position1 = position + direction1
    position2 = position + direction2

    while True:
        for _ in range(side_len):
            boundaries[(position1, normal1)] = position2, -normal2
            boundaries[(position2, normal2)] = position1, -normal1
            position1 += direction1
            position2 += direction2

        if (position1 not in board2 and position2 not in board2):
            # Next edges are not zipped.
            break

        # Turn.
        if position1 not in board2:
            position1 -= direction1
            normal1 = direction1
            direction1 *= 1j
            direction1 = (direction1 if position1 + direction1 in board2
                          else -direction1)
        if position2 not in board2:
            position2 -= direction2
            normal2 = direction2
            direction2 *= 1j
            direction2 = (direction2 if position2 + direction2 in board2
                          else -direction2)

# Might have boundaries missing if we have one of nets 5, 6, 8 and 10 from
# https://w.wiki/6AkN. Handle nets 5, 6 and 8, too lazy to handle 10.
if len(boundaries) == 12 * side_len:
    r_zeros = {p for p in board2 if p.real == 0
               and (p, -1+0j) not in boundaries}
    r_maxes = {p for p in board2 if p.real == 4*side_len-1
               and (p, 1+0j) not in boundaries}
    c_zeros = {p for p in board2 if p.imag == 0
               and (p, -1j) not in boundaries}
    c_maxes = {p for p in board2 if p.imag == 4*side_len-1
               and (p, 1j) not in boundaries}

    if (r_zeros and r_maxes) or (c_zeros and c_maxes):
        # Nets 5 or 6.
        if r_zeros and r_maxes:
            matches = ((x0, x1) for x0 in r_zeros for x1 in r_maxes
                       if x0.imag == x1.imag)
            direction = -1 + 0j
        else:
            matches = ((x0, x1) for x0 in c_zeros for x1 in c_maxes
                       if x0.real == x1.real)
            direction = -1j

        for x0, x1 in matches:
            boundaries[(x0, +direction)] = x1, +direction
            boundaries[(x1, -direction)] = x0, -direction
    else:
        # Net 8.
        r_bounds, r_direction = (r_zeros, -1+0j) if r_zeros else (r_maxes,
                                                                  1+0j)
        c_bounds, c_direction = (c_zeros, -1j) if c_zeros else (c_maxes, 1j)

        # r_bounds and c_bounds' elements are zipped closest to furthest
        # relative to quadrant r_direction + c_direction.
        from operator import attrgetter
        sorted_r_bounds = sorted(r_bounds, key=attrgetter('imag'),
                                 reverse=(c_direction == 0+1j))
        sorted_c_bounds = sorted(c_bounds, key=attrgetter('real'),
                                 reverse=(r_direction == 1+0j))
        for x0, x1 in zip(sorted_r_bounds, sorted_c_bounds):
            boundaries[(x0, r_direction)] = x1, -c_direction
            boundaries[(x1, c_direction)] = x0, -r_direction

elif len(boundaries) == 10 * side_len:
    raise ValueError("Input is https://w.wiki/6AkN net 10 (not implemented).")

assert len(boundaries) == 14*side_len


def move_forward_complex(direction, pos):
    next_pos = pos + direction
    if next_pos not in board2:
        next_pos, next_direction = boundaries[(pos, direction)]
    else:
        next_direction = direction

    if board2[next_pos] == '#':
        return direction, pos

    return next_direction, next_pos


def run_complex(direction, pos):
    for numeric, g in groupby(instructions, str.isdigit):
        instruction = ''.join(g)
        if numeric:
            val = int(instruction)
            for i in range(val):
                direction, pos = move_forward_complex(direction, pos)
        else:
            if len(instruction) > 1:
                raise ValueError(g)
            direction *= 1j if instruction == 'L' else -1j

    return direction, pos


directions_complex = [1j, 1+0j, -1j, -1+0j]

start_pos = complex(start_r, start_c)
direction, pos = run_complex(1j, start_pos)
print("Part 2:", (1000*(int(pos.real)+1) + 4*(int(pos.imag)+1)
                  + directions_complex.index(direction)))
