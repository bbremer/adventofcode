from dataclasses import dataclass
from itertools import pairwise
from operator import attrgetter

from input_parser import parse_day


@dataclass
class Number:
    value: int
    fwd = None
    rev = None


def init_numbers(test, multiple=1):
    if test:
        test_data = """1
        2
        -3
        3
        -2
        0
        4"""
        numbers = [Number(int(x)) for x in test_data.split('\n')]
    else:
        numbers = [Number(int(x)) for x in parse_day(20)]

    for n in numbers:
        n.value *= multiple

    numbers[0].rev = numbers[-1]
    numbers[-1].fwd = numbers[0]

    for n1, n2 in pairwise(numbers):
        n1.fwd = n2
        n2.rev = n1

    return numbers


def iterate(numbers):
    curr = next(x for x in numbers if x.value == 0)
    while True:
        yield curr
        curr = curr.fwd
        if not curr.value:
            break


def mix(numbers):
    for n in numbers:
        abs_value = abs(n.value) % (len(numbers) - 1)
        if not abs_value:
            continue

        next_attr, prev_attr = (('fwd', 'rev') if n.value > 0
                                else ('rev', 'fwd'))

        next_get = attrgetter(next_attr)
        prev_get = attrgetter(prev_attr)
        def next_set(obj, value): return setattr(obj, next_attr, value)
        def prev_set(obj, value): return setattr(obj, prev_attr, value)

        next_set(prev_get(n), next_get(n))
        prev_set(next_get(n), prev_get(n))

        target_prev = n
        for i in range(abs_value):
            target_prev = next_get(target_prev)
        target_next = next_get(target_prev)

        next_set(n, target_next)
        prev_set(n, target_prev)

        next_set(target_prev, n)
        prev_set(target_next, n)

        '''
        print(n.value, abs_value)
        print(', '.join([str(x.value) for x in iterate(numbers)]))
        print()
        '''


TEST = False

numbers = init_numbers(TEST)
# print(', '.join([str(x.value) for x in iterate(numbers)]))
# print()
mix(numbers)
final_numbers = list(iterate(numbers))

# print(*[final_numbers[i % len(final_numbers)].value
#       for i in range(1000, 3001, 1000)])
print("Part 1:", sum(final_numbers[i % len(final_numbers)].value
                     for i in range(1000, 3001, 1000)))

numbers = init_numbers(TEST, 811589153)
for i in range(10):
    mix(numbers)
final_numbers = list(iterate(numbers))

print("Part 2:", sum(final_numbers[i % len(final_numbers)].value
                     for i in range(1000, 3001, 1000)))
