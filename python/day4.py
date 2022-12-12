from itertools import tee

from input_parser import parse_day


single_split_lines = (line.split(',') for line in parse_day(4))
split_lines = ((a.split('-'), b.split('-')) for a, b in single_split_lines)
int_lines = (((int(a), int(b)), (int(c), int(d)))
             for (a, b), (c, d) in split_lines)

int_lines1, int_lines2 = tee(int_lines)
print("Part 1:", sum(1 for (a, b), (c, d) in int_lines1
                     if a >= c and b <= d or c >= a and d <= b))
print("Part 2:", sum(1 for (a, b), (c, d) in int_lines2
                     if any((c <= a <= d, c <= b <= d, a <= c <= b,
                             a <= d <= b))))
