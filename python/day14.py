from itertools import count, pairwise

from input_parser import parse_day


test_data = """498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9"""


def parse_rock_paths(test=False):
    rock_list = []

    if test:
        gen = iter(test_data.split('\n'))
    else:
        gen = parse_day(14)

    for rock_path_str in gen:
        rock_path_vertices = []
        for rock_line_vertex_str in rock_path_str.split(' -> '):
            x, y = rock_line_vertex_str.split(',')
            rock_path_vertices.append((int(x), int(y)))

        for (x0, y0), (x1, y1) in pairwise(rock_path_vertices):
            if x0 == x1:
                for y in range(y0, y1, 1 if y0 < y1 else -1):
                    rock_list.append((x0, y))
            else:
                for x in range(x0, x1, 1 if x0 < x1 else -1):
                    rock_list.append((x, y0))

            rock_list.append((x1, y1))

    return set(rock_list)


rocks = parse_rock_paths(False)
max_y_rocks = max(rock[1] for rock in rocks)

occupied = rocks.copy()
part1_sands = -1
for i in count(0):
    print(i, '\r', end='', flush=True)

    if (500, 0) in occupied:
        break

    x, y = 500, 0
    while True:
        if part1_sands == -1 and y > max_y_rocks:
            part1_sands = i

        if y >= max_y_rocks + 1:
            occupied.add((x, y))
            break

        y = y+1

        if (x, y) not in occupied:
            continue
        elif (x-1, y) not in occupied:
            x = x-1
        elif (x+1, y) not in occupied:
            x = x+1
        else:
            occupied.add((x, y-1))
            break

'''
min_x = min(occ[0] for occ in occupied)
max_x = max(occ[0] for occ in occupied)
min_y = min(occ[1] for occ in occupied)
max_y = max(occ[1] for occ in occupied)

occupied_grid = [['.' for _ in range(max_x+1)] for _ in range(max_y+1)]
for (x, y) in occupied - rocks:
    occupied_grid[y][x] = 'o'
for (x, y) in rocks:
    occupied_grid[y][x] = '#'
for row in occupied_grid[min_y:max_y+1]:
    print(''.join(row[min_x:max_x+1]))
'''

print("Part 1:", part1_sands)
print("Part 2:", i)
