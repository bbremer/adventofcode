from operator import itemgetter

from tqdm import tqdm

from input_parser import parse_day


class Sensor:
    def __init__(self, line):
        sensor_str, beacon_str = line.split(': ')

        x, y = sensor_str.replace(',', '').split(' ')[-2:]
        self.x, self.y = int(x[2:]), int(y[2:])

        beacon_x, beacon_y = beacon_str.replace(',', '').split(' ')[-2:]
        self.bx, self.by = int(beacon_x[2:]), int(beacon_y[2:])
        self.distance = abs(self.x - self.bx) + abs(self.y - self.by)

    def row_bounds(self, y):
        y_dis = abs(self.y - y)
        if y_dis > self.distance:
            return
        start = self.x - self.distance + y_dis
        stop = self.x + self.distance - y_dis + 1
        return start, stop

    def relevant(self, y):
        return abs(self.y - y) <= self.distance


test_data = """Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3"""


# REAL DATA
PART1_Y = 2 * 10**6
PART2_MIN = 0
PART2_MAX = 4 * 10**6
sensors = [Sensor(line) for line in parse_day(15)]

# TEST DATA
'''
PART1_Y = 10
PART2_MIN = 0
PART2_MAX = 20
sensors = [Sensor(line) for line in test_data.split('\n')]
'''


def split_range(start, stop, position):
    if not start <= position < stop:
        yield start, stop
    else:
        if position > start:
            yield start, position
        if position < stop-1:
            yield position+1, stop


def sorted_bounds(y, sensor_beacon_positions=[]):
    relevant_sensors = (s for s in sensors if s.relevant(y))
    bounds = sorted([s.row_bounds(y) for s in relevant_sensors],
                    key=itemgetter(0), reverse=True)
    while bounds:
        start, stop = bounds.pop()
        if not sensor_beacon_positions:
            yield start, stop
        else:
            ranges = [(start, stop)]
            for sp in sensor_beacon_positions:
                new_ranges = []
                for start, stop in ranges:
                    new_ranges.extend(list(split_range(start, stop, sp)))
                ranges = sorted(new_ranges, key=itemgetter(0), reverse=True)

            yield from ranges


def possible_sensor_positions(y):
    sensor_positions = [s.x for s in sensors if s.y == y]
    beacon_positions = [s.bx for s in sensors if s.by == y]
    bounds = sorted_bounds(y, sensor_positions+beacon_positions)

    start, stop = next(bounds)
    current_index = stop
    count = stop - start

    for start, stop in bounds:
        if current_index >= stop:
            continue
        if current_index < start:
            count += stop - start
        else:
            count += stop - current_index
        current_index = stop

    return count


def unoccupied_position(y, min_x, max_x):
    current_index = 0
    for start, stop in sorted_bounds(y):
        if stop <= current_index:
            continue
        if current_index < start:
            return current_index
        current_index = stop

    if current_index <= max_x:
        return current_index


print("Part 1:", possible_sensor_positions(PART1_Y))


for y in tqdm(range(PART2_MIN, PART2_MAX+1)):
    unocc = unoccupied_position(y, PART2_MIN, PART2_MAX)
    if unocc is not None:
        print("Part 2:", unocc * 4 * 10**6 + y)
        break
