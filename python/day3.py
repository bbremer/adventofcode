from functools import reduce

from input_parser import parse_day

priorities = ({chr(n): n-ord('a')+1 for n in range(ord('a'), ord('z')+1)}
              | {chr(n): n-ord('A')+27 for n in range(ord('A'), ord('Z')+1)})


'''
def duplicate_letter(line):
    half = len(line) // 2
    return set(line[:half]).intersection(line[half:]).pop()
'''


def duplicate_letter(strs):
    return set.intersection(*str).pop()


def group_priorities(lines):
    line_list = list(lines)
    part1_strs = ((line[:len(line)//2], line[len(line)//2:])
                  for line in line_list)
    part1_dups = (set(strs[0]).intersection(*strs[1:]).pop()
                  for strs in part1_strs)
    part1_sum = sum(priorities[letter] for letter in part1_dups)

    part2_dup = set(line_list[0]).intersection(*line_list[1:]).pop()
    part2_priority = priorities[part2_dup]

    return part1_sum, part2_priority


gps = (group_priorities(lines) for lines in zip(*[parse_day(3)]*3))
part1_sum, part2_sum = reduce(lambda x, y: (x[0]+y[0], x[1]+y[1]), gps)
print("Part 1:", part1_sum)
print("Part 2:", part2_sum)
