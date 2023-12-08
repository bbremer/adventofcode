from operator import add, eq, floordiv, mul, sub
import re

from input_parser import parse_day

int_line = re.compile(r"([a-z]{4}): (\d+)")
op_line = re.compile(r"([a-z]{4}): ([a-z]{4}) (.) ([a-z]{4})")
ops = {"+": add, "==": eq, "/": floordiv, "*": mul, "-": sub}
rev_ops = {v: k for k, v in ops.items()}
inverses = {"+": sub, "/": mul, "*": floordiv, "-": add}

exp_pattern1 = re.compile(r"(.*)(.)(\d+)")
exp_pattern2 = re.compile(r"(\d+)(.)(.*)")


# class Monkeys(Mapping):
class Monkeys:
    def __init__(self, monkeys_line_iter):
        self._monkeys = {}

        for line in monkeys_line_iter:
            m = re.match(int_line, line)
            if m is not None:
                k, v = m.groups()
                self._monkeys[k] = int(v)
                continue

            k, a, op_str, b = re.match(op_line, line).groups()
            self._monkeys[k] = a, ops[op_str], b

    def __getitem__(self, index):
        value = self._monkeys[index]

        if isinstance(value, int) or isinstance(value, str):
            return value

        a, op, b = value
        new_a, new_b = self[a], self[b]
        if isinstance(new_a, int) and isinstance(new_b, int):
            new_value = op(new_a, new_b)
            self._monkeys[index] = new_value
            return new_value

        return f"({new_a}{rev_ops[op]}{new_b})"


def eval2(a, exp):
    if exp[0] == 'x':
        return inverses[exp[1]](a, int(exp[2:]))
    if exp[-1] == 'x':
        op_str = exp[2]
        if op_str == '/':
            return int(exp[:-2]) // a
        if op_str == '-':
            return int(exp[:-2]) - a
        return inverses[exp[-2]](a, int(exp[:-2]))

    if exp[0] == '(':
        new_exp_end = exp.rfind(')')
        new_exp = exp[1:new_exp_end]
        op_str = exp[new_exp_end+1]
        new_a = int(exp[new_exp_end+2:])
    else:
        new_exp_start = exp.find('(')
        new_a = int(exp[:new_exp_start-1])
        op_str = exp[new_exp_start-1]
        new_exp = exp[new_exp_start+1:-1]

        if op_str == '/':
            a = new_a // a
            return eval2(a, new_exp)
        if op_str == '-':
            a = new_a - a
            return eval2(a, new_exp)

    op_inv = inverses[op_str]
    return eval2(op_inv(a, new_a), new_exp)


test_data = """root: pppw + sjmn
dbpl: 5
cczh: sllz + lgvd
zczc: 2
ptdq: humn - dvpt
dvpt: 3
lfqf: 4
humn: 5
ljgn: 2
sjmn: drzm * dbpl
sllz: 4
pppw: cczh / lfqf
lgvd: ljgn * ptdq
drzm: hmdt - zczc
hmdt: 32""".split('\n')

test_data2 = """root: juli + josi
juli: amee + alex
amee: buki * abby
buki: 5
abby: 4
alex: 4
josi: benj / mark
benj: 360
mark: emly - humn
emly: 34
humn: 0""".split('\n')

gen = list(parse_day(21))
# gen = test_data2

monkeys = Monkeys(gen)
print('Part 1:', monkeys['root'])

monkeys = Monkeys(gen)
a, _, b = monkeys._monkeys['root']
monkeys._monkeys['root'] = a, eq, b
monkeys._monkeys['humn'] = 'x'
root_expression = monkeys['root']

# Get a as the int.
a, b = root_expression[1:-1].split("==")
if b.isdigit():
    a, b = b, a
a = int(a)
exp = b[1:-1]
ans = eval2(a, exp)
print("Part 2:", ans)

# Check answer.
monkeys = Monkeys(gen)
a, _, b = monkeys._monkeys['root']
monkeys._monkeys['root'] = a, sub, b
monkeys._monkeys['humn'] = ans
assert monkeys['root'] == 0
