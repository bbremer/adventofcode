from input_parser import parse_day

input_str = next(parse_day(6))


def find_marker(marker_len):
    start_iter = range(len(input_str))
    end_iter = range(marker_len, len(input_str))
    for start, end in zip(start_iter, end_iter):
        if len(set(input_str[start:end])) == marker_len:
            return end


if __name__ == "__main__":
    print("Part 1:", find_marker(4))
    print("Part 2:", find_marker(14))
