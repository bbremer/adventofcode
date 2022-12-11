from functools import reduce
from itertools import groupby

from input_parser import parse_day

elf_sums = (sum(int(i) for i in g)
            for k, g in groupby(parse_day(1), bool) if k)
three_maxes = reduce(lambda x, y: sorted(x + [y])[-3:], elf_sums, [])
print('Part 1:', three_maxes[-1])
print('Part 2:', sum(three_maxes))
