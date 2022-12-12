from itertools import takewhile

from input_parser import parse_day


def init_stack():
    day_gen = parse_day(5)
    init_lines = list(takewhile(bool, day_gen))
    column_to_stack_index = [(col, index) for col, index
                             in enumerate(init_lines[-1], 1) if index != ' ']
    # ^ enumerate from 1 because of strip() in parse_day.

    stacks = {index: [] for _, index in column_to_stack_index}
    for line in init_lines[-2::-1]:
        for col, index in column_to_stack_index:
            if line[col] != ' ':
                stacks[index].append(line[col])

    return stacks, day_gen


# Start Part 1.
p1_stacks, p1_instructions = init_stack()
for line in p1_instructions:
    _, amount_str, _, i0, _, i1 = line.split()
    for _ in range(int(amount_str)):
        p1_stacks[i1].append(p1_stacks[i0].pop())

p1_sorted_stacks = [p1_stacks[i] for i in sorted(p1_stacks, key=int)]
print("Part 1:", ''.join([stack[-1] for stack in p1_sorted_stacks]))


# Start Part 2.
p2_stacks, p2_instructions = init_stack()
for line in p2_instructions:
    _, amount_str, _, i0, _, i1 = line.split()
    # moved_crates = [p2_stacks[i0].pop() for _ in range(int(amount_str))]
    # p2_stacks[i1].extend(moved_crates[::-1])
    src_stack = p2_stacks[i0]
    p2_stacks[i1].extend(src_stack[-int(amount_str):])
    p2_stacks[i0] = src_stack[:-int(amount_str)]

p2_sorted_stacks = [p2_stacks[i] for i in sorted(p2_stacks, key=int)]
print("Part 1:", ''.join([stack[-1] for stack in p2_sorted_stacks]))
