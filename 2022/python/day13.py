from functools import cmp_to_key, reduce
from itertools import groupby, zip_longest
import json

from input_parser import parse_day


def compare(val1, val2, i=0):
    if val1 == val2:
        return 0

    if val1 is None:
        return -1
    elif val2 is None:
        return 1

    if isinstance(val1, int) and isinstance(val2, int):
        if val1 < val2:
            return -1
        else:
            return 1

    if isinstance(val1, int):
        val1 = [val1]
    elif isinstance(val2, int):
        val2 = [val2]

    for v1, v2 in zip_longest(val1, val2):
        comp = compare(v1, v2, i+1)
        if comp != 0:
            return comp

    return 0


pairs = [(json.loads(next(g)), json.loads(next(g))) for k, g
         in groupby(parse_day(13), bool) if k]

valid_indices = [i for i, (p1, p2) in enumerate(pairs, 1)
                 if compare(p1, p2) == -1]
print("Part 1:", sum(valid_indices))

packets = reduce(lambda x, y: x + list(y), pairs, []) + [[[2]], [[6]]]
sorted_packets = sorted(packets, key=cmp_to_key(compare))
print("Part 2:",
      (sorted_packets.index([[2]])+1)*(sorted_packets.index([[6]])+1))
