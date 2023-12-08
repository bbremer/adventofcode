from collections import deque

from input_parser import parse_day


test_data = """2,2,2
1,2,2
3,2,2
2,1,2
2,3,2
2,2,1
2,2,3
2,2,4
2,2,6
1,2,5
3,2,5
2,1,5
2,3,5""".split('\n')


def neighbors(x, y, z):
    return ((x-1, y, z), (x+1, y, z), (x, y-1, z), (x, y+1, z), (x, y, z-1),
            (x, y, z+1))


points = {tuple(int(i) for i in line.split(','))
          for line in parse_day(18)}
# points = {tuple(int(i) for i in line.split(','))
#           for line in test_data}

p1 = len(points) * 6 - sum(sum(1 for n in neighbors(*p) if n in points)
                           for p in points)
print("Part 1:", p1)

max_x = max(p[0] for p in points)
min_x = min(p[0] for p in points)
max_y = max(p[1] for p in points)
min_y = min(p[1] for p in points)
max_z = max(p[2] for p in points)
min_z = min(p[2] for p in points)

stack = deque()
stack.append((0, 0, 0))
outside = set()
p2 = 0
while stack:
    p = stack.pop()
    if any((not min_x-1 <= p[0] <= max_x+1,
            not min_y-1 <= p[1] <= max_y+1,
            not min_z-1 <= p[2] <= max_z+1,)):
        continue

    if p not in outside and p not in points:
        outside.add(p)
        for n in neighbors(*p):
            if n in points:
                p2 += 1
            else:
                stack.append(n)

print("Part 2:", p2)
