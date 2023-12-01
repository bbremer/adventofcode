from input_parser import parse_day


TEST_DATA = """1=-0-2
12111
2=0=
21
2=01
111
20012
112
1=-1=
1-12
12
1=
122""".split('\n')

TEST = False

gen = iter(TEST_DATA) if TEST else parse_day(25)

snafu_letter_to_decimal = {'2': 2, '1': 1, '0': 0, '-': -1, '=': -2}
base5_to_snafu = {4: '2', 3: '1', 2: '0', 1: '-', 0: '='}


def snafu_to_decimal(snafu):
    return sum(snafu_letter_to_decimal[letter] * 5**i
               for i, letter in enumerate(reversed(snafu)))


def decimal_to_snafu(decimal):
    """Black magic from Concrete Mathematics."""
    n = 0
    total = 0
    while True:
        total += 2 * 5**n
        if total >= decimal:
            break
        else:
            n += 1

    adjusted_decimal = decimal + (5**(n+1) - 1) // 2
    base5 = []  # Little endian.
    while adjusted_decimal:
        adjusted_decimal, b = divmod(adjusted_decimal, 5)
        base5.append(b)

    return ''.join([base5_to_snafu[x] for x in reversed(base5)])


decimal = sum(snafu_to_decimal(x) for x in gen)
snafu = decimal_to_snafu(decimal)
print("Part 1:", snafu)
