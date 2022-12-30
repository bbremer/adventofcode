from itertools import combinations, groupby, pairwise, product, takewhile

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

TEST = True

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

        print(r, c, facing)

    return r, c


side_len = 4 if TEST else 50


# final_r, final_c = run(wrap_p1, start_r, start_c, facing)
# print("Part 1:", 1000*(final_r+1) + 4*(final_c+1) + directions.index(facing))


inside_corners = []
for r, row in enumerate(board[1:-1], 1):
    for c, e in enumerate(row[1:-1], 1):
        if e != ' ' and sum(board[r+dr][c+dc] == ' ' for dr, dc
                            in product((-1, 0, 1), repeat=2)) == 1:
            inside_corners.append((r, c))

for inside_r, inside_c in inside_corners:
    dr, dy = [prod[] for dr, dy in ]


'''
# Get the square upper left corners.
inside_edges = set()
for sq_r, sq_c in product(range(0, max_r, side_len),
                          range(0, max_c, side_len)):
    for add_r, add_c in ((0, 0), (side_len-1, 0), (0, side_len-1),
                         (side_len-1, side_len-1)):
        r, c = sq_r + add_r, sq_c + add_c
        if 0 < r < max_r and 0 < c < max_c:
            inside_edges.add((r, c))

print(max_r)
print(max_c)
for r, c in inside_edges:
    print(r, c)
    # if 0 < r < max_r and 0 < c < max_c:
'''


def wrap_p2(direction, r, c):
    pass


import sys
sys.exit()

r, c = run(wrap_p2, r, c, facing)
print("Part 2:", 1000*(r+1) + 4*(c+1) + directions.index(facing))
