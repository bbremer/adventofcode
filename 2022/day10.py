from input_parser import parse_day

checked_cycles = iter([20, 60, 100, 140, 180, 220, None])
next_cycle = next(checked_cycles)
strengths_sum = 0
current_cycle = 1
x_register = 1

for instruction in parse_day(10):
    if current_cycle == next_cycle:
        strengths_sum += current_cycle * x_register
        next_cycle = next(checked_cycles)
    current_cycle += 1
    match instruction.split():
        case "addx", value:
            if current_cycle == next_cycle:
                strengths_sum += current_cycle * x_register
                next_cycle = next(checked_cycles)
            current_cycle += 1
            x_register += int(value)

print("Part 1:", strengths_sum)

rows = []
x_register = 1
current_position = 0
current_row = ''

for instruction in parse_day(10):
    if current_position in (x_register-1, x_register, x_register+1):
        current_row += '#'
    else:
        current_row += '.'
    current_position += 1
    if current_position == 40:
        current_position = 0
        rows.append(current_row)
        current_row = ''
    match instruction.split():
        case "addx", value:
            if current_position in (x_register-1, x_register, x_register+1):
                current_row += '#'
            else:
                current_row += '.'
            current_position += 1
            if current_position == 40:
                current_position = 0
                rows.append(current_row)
                current_row = ''
            x_register += int(value)
rows.append(current_row)

for row in rows:
    print(row)
