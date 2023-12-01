from dataclasses import dataclass

from input_parser import parse_day


@dataclass
class Node:
    height: int
    shortest_path: list | None = None


se_pos = {}
nodes = []
for r, line in enumerate(parse_day(12)):
    node_line = []
    for c, letter in enumerate(line):
        if 'a' <= letter <= 'z':
            node_line.append(Node(ord(letter)))
        else:
            se_pos[letter] = (r, c)
            node_line.append(Node(ord('a') if letter == 'S' else ord('z')))
    nodes.append(node_line)

max_r = len(nodes) - 1
max_c = len(nodes[0]) - 1


def conditional_push(curr_node, next_r, next_c):
    next_node = nodes[next_r][next_c]
    if (not next_node.shortest_path
            and curr_node.height - next_node.height <= 1):
        next_node.shortest_path = curr_node.shortest_path + [(next_r, next_c)]
        next_nodes.append((next_r, next_c))


e_r, e_c = se_pos['E']
nodes[e_r][e_c].shortest_path = [(e_r, e_c)]

curr_nodes = [(e_r, e_c)]
next_nodes = []
i = 0
part2_finished = False
while curr_nodes:
    for last_r, last_c in curr_nodes:
        curr_node = nodes[last_r][last_c]
        if not part2_finished and curr_node.height == ord('a'):
            print("Part 2:", len(curr_node.shortest_path)-1)
            part2_finished = True
        if (last_r, last_c) == se_pos['S']:
            print("Part 1:", len(curr_node.shortest_path)-1)  # includes E, S.
            import sys
            sys.exit()

        if last_r > 0:
            conditional_push(curr_node, last_r-1, last_c)
        if last_r < max_r:
            conditional_push(curr_node, last_r+1, last_c)
        if last_c > 0:
            conditional_push(curr_node, last_r, last_c-1)
        if last_c < max_c:
            conditional_push(curr_node, last_r, last_c+1)

    curr_nodes = next_nodes
    next_nodes = []
    i += 1

print()
for r, line in enumerate(nodes):
    for c, n in enumerate(line):
        if n.shortest_path:
            print('X', end='')
        else:
            print('.', end='')
    print()
