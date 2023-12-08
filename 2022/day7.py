from itertools import groupby
from pathlib import PurePosixPath

from input_parser import parse_day

ROOT = PurePosixPath("/")
current_directory = ROOT
directory_to_size: dict[PurePosixPath, int] = {}


def line_parser():
    for key, group in groupby(parse_day(7), lambda s: s[0] == '$'):
        if key:
            yield from group
        else:
            item1s = (line.split()[0] for line in group)
            yield sum(int(item1) for item1 in item1s if item1 != 'dir')


def cd_back(current_directory):
    current_size = directory_to_size[current_directory]
    directory_to_size[current_directory.parent] += current_size
    return current_directory.parent


instrs_or_fsizes = line_parser()
for line in instrs_or_fsizes:
    match line.split():
        case '$', 'cd', '..':
            # Lazily evaluate update parent.
            current_directory = cd_back(current_directory)
        case '$', 'cd', directory:
            current_directory /= directory
        case '$', 'ls':
            directory_to_size[current_directory] = next(instrs_or_fsizes)


# Need to finish traversal at root, so simulate multiple "$ cd ..".
while current_directory != ROOT:
    current_directory = cd_back(current_directory)


print("Part 1:", sum(v for v in directory_to_size.values() if v <= 100000))

deleted_size_needed = directory_to_size[ROOT] - 4E7
print("Part 2:", min(v for v in directory_to_size.values()
                     if v >= deleted_size_needed))
