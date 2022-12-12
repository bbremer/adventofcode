from dataclasses import dataclass
from itertools import pairwise

from input_parser import parse_day


@dataclass
class Knot:
    x: int
    y: int

    def move(self, direction):
        match direction:
            case 'U':
                self.y += 1
            case 'R':
                self.x += 1
            case 'D':
                self.y -= 1
            case 'L':
                self.x -= 1

    def catchup(self, other: 'Knot'):
        x_diff = other.x - self.x
        y_diff = other.y - self.y

        if x_diff > 1:
            self.move('R')
            if y_diff:
                self.move('U' if y_diff > 0 else 'D')
        elif x_diff < -1:
            self.move('L')
            if y_diff:
                self.move('U' if y_diff > 0 else 'D')
        elif y_diff > 1:
            self.move('U')
            if x_diff:
                self.move('R' if x_diff > 0 else 'L')
        elif y_diff < -1:
            self.move('D')
            if x_diff:
                self.move('R' if x_diff > 0 else 'L')


def sim_positions(num_knots):
    knots = [Knot(0, 0) for _ in range(num_knots)]
    head = knots[0]
    tail = knots[-1]
    history = {(tail.x, tail.y)}

    for line in parse_day(9):
        direction, times = line.split()
        for _ in range(int(times)):
            head.move(direction)
            for k1, k2 in pairwise(knots):
                k2.catchup(k1)
            history.add((tail.x, tail.y))

    return len(history)


print("Part 1:", sim_positions(2))
print("Part 2:", sim_positions(10))
