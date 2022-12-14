from dataclasses import dataclass
from functools import partial, reduce
from itertools import groupby
from operator import add, mul
from typing import Callable


from input_parser import parse_day


@dataclass
class Monkey:
    items: list
    operation: Callable[[int], int]
    test_divisor: int
    test_true_destination: int
    test_false_destination: int
    num_items_inspected = 0

    def inspect_and_throw_items(self):
        self.num_items_inspected += len(self.items)
        while self.items:
            item = self.items.pop()
            thrown_item = self.operation(item)
            if not div_three:
                thrown_item %= divisor_pi
            destination = (self.test_true_destination
                           if not thrown_item % self.test_divisor
                           else self.test_false_destination)
            yield thrown_item, destination


def parse_monkey(monkey_lines):
    next(monkey_lines)  # "Monkey {i}:".

    items_line = next(monkey_lines)
    items = [int(i.strip()) for i in items_line.split(': ')[1].split(',')]

    op_str, op_arg = next(monkey_lines).strip().split()[-2:]
    if op_arg == 'old':
        operation = partial(pow, exp=2)
    else:
        operation = partial({'*': mul, '+': add}[op_str], int(op_arg))

    if div_three:
        def operation2(x): return operation(x) // 3
    else:
        operation2 = operation

    test_divisor = int(next(monkey_lines).split()[-1])
    test_true_destination = int(next(monkey_lines).split()[-1])
    test_false_destination = int(next(monkey_lines).split()[-1])

    return Monkey(items, operation2, test_divisor, test_true_destination,
                  test_false_destination)


div_three = True
file = parse_day(11)
monkeys = [parse_monkey(g) for k, g in groupby(parse_day(11), bool) if k]
for _ in range(20):
    for monkey in monkeys:
        for item, destination in monkey.inspect_and_throw_items():
            monkeys[destination].items.append(item)
print("Part 1:", mul(*sorted(m.num_items_inspected for m in monkeys)[-2:]))

div_three = False
file = parse_day(11)
monkeys = [parse_monkey(g) for k, g in groupby(parse_day(11), bool) if k]
divisor_pi = reduce((lambda x, y: x*y.test_divisor), monkeys, 1)
for i in range(10000):
    for monkey in monkeys:
        for item, destination in monkey.inspect_and_throw_items():
            monkeys[destination].items.append(item)
print("Part 1:", mul(*sorted(m.num_items_inspected for m in monkeys)[-2:]))
