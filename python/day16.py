"""Shamelessly derived from
https://github.com/juanplopes/advent-of-code-2022/blob/main/day16.py.
Thanks for the lesson Juan!"""

from collections import deque
from itertools import product
import re

from input_parser import parse_day


valves_to_rates = {}
valves_to_tunnels = {}
template = re.compile("Valve (..) has flow rate=(.+); "
                      "tunnels? leads? to valves? (.+)")
for line in parse_day(16):
    m = template.match(line)
    assert m is not None
    name, rate_str, tunnels_str = m.groups()
    valves_to_rates[name] = int(rate_str)
    valves_to_tunnels[name] = tunnels_str.split(', ')

valves_to_rates = {v: r for v, r in valves_to_rates.items() if r > 0}

# Floyd-Warshall: https://w.wiki/67rE
dist = {name: {name2: 1 if name2 in tunnels else float("inf")
               for name2 in valves_to_tunnels}
        for name, tunnels in valves_to_tunnels.items()}
for k in valves_to_tunnels:
    for i in valves_to_tunnels:
        for j in valves_to_tunnels:
            dist[i][j] = min(dist[i][j], dist[i][k] + dist[k][j])

valve_flags = {name: 1 << i for i, name in enumerate(valves_to_rates)}


# DFS.
def path_maximums(starting_minutes):
    path_maxs = {}
    stack = deque()
    stack.append(("AA", starting_minutes, 0, 0))
    while stack:
        name, minutes, rate, flags = stack.pop()
        path_maxs[flags] = max(path_maxs.get(flags, 0), rate)
        for valve, valve_rate in valves_to_rates.items():
            new_minutes = minutes - dist[name][valve] - 1
            if valve_flags[valve] & flags or new_minutes <= 0:
                continue
            new_rate = rate + new_minutes*valve_rate
            stack.append((valve, new_minutes, new_rate,
                          valve_flags[valve] | flags))
    return path_maxs


print("Part 1:", max(path_maximums(30).values()))

p2_path_mins_iter = product(path_maximums(26).items(), repeat=2)
print("Part 2:", max(m1+m2 for (p1, m1), (p2, m2) in p2_path_mins_iter
                     if not p1 & p2))
