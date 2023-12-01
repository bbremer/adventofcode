from itertools import cycle

from input_parser import parse_day
instruction_str = next(parse_day(17))
# instruction_str = ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"

# num_instructions = len(instruction_str)

rock_lists = [
        [15],
        [2, 7, 2],
        [7, 4, 4],
        [1, 1, 1, 1],
        [3, 3],
        ]


def byte_to_img(byte):
    return ''.join(['#' if byte & 2**i else '.' for i in range(7)])


def int_to_imgs(data, length=4):
    byte_list = [(data >> 8*i) & 255 for i in range(length)]
    return [byte_to_img(byte) for byte in byte_list if byte]


rocks = [(int.from_bytes(i, 'little'), len(i)) for i in rock_lists]

divisor = len(rock_lists) * len(instruction_str)

x_lim = int.from_bytes([128]*4, 'little')


def tower(num_rocks):
    jet_iter = cycle(instruction_str)
    FLOOR = -1
    starting_y = 3
    settled_positions = bytearray(1000000000)

    floor_checks = {}
    ys = []
    new_floor = FLOOR

    for i, (rock, rock_height) in enumerate(cycle(rocks)):

        if i >= num_rocks:
            return starting_y - 3

        if not i % divisor and i:
            new_floor = starting_y - 30
            last_floors = bytes(settled_positions[starting_y-30:starting_y+1])
            if last_floors in floor_checks:
                prev_i, prev_y = floor_checks[last_floors]
                y = ys[-1]

                target = num_rocks - i
                cycle_size = i - prev_i
                num_cycles = target // cycle_size
                body_height = num_cycles * (y - prev_y)

                cycle_excess = target % cycle_size
                return (y + body_height + ys[prev_i+cycle_excess-1]
                        - ys[prev_i-1])
            floor_checks[last_floors] = (i, starting_y-3)

        if not i % 1000000:
            print(i, '\r', flush=True, end='')

        x, y = 4, starting_y

        while True:
            new_x = x * 2 if next(jet_iter) == '>' else x >> 1
            new_rock = new_x * rock

            if (1 <= new_x and not new_rock & x_lim
                    and not new_rock & int.from_bytes(settled_positions[y:y+4],
                                                      'little')):
                x = new_x

            new_rock = x * rock

            new_y = y - 1

            if (new_y == FLOOR
                or new_rock & int.from_bytes(settled_positions[new_y:new_y+4],
                                             'little')):
                if y <= new_floor:
                    print("new floor hit:", y, new_floor)
                    raise ValueError(f"new floor hit {y=} {new_floor=}")
                for bi in range(4):
                    settled_positions[y+bi] += (new_rock >> 8*bi) & 255
                break
            y = new_y

        starting_y = max(starting_y, y + rock_height + 3)
        if starting_y + 4 > len(settled_positions):
            settled_positions.extend([0]*4)
        ys.append(starting_y - 3)


print("Part 1:", tower(2022))
print("Part 2:", tower(1000000000000))
