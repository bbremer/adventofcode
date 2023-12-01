from collections import namedtuple

from input_parser import parse_day

Tree = namedtuple("Tree", "height r c")


def edge_visibles(trees_iter):
    up_steps = [next(trees_iter)]
    down_steps = []

    for tree in trees_iter:
        if tree.height > up_steps[-1].height:
            up_steps.append(tree)
            down_steps = []
        elif not down_steps or tree.height < down_steps[-1].height:
            down_steps.append(tree)
        else:
            while down_steps:
                if down_steps[-1].height > tree.height:
                    break
                down_steps.pop()
            down_steps.append(tree)

    return set(up_steps + down_steps)  # Concave down.


trees = [[Tree(int(h), r, c) for c, h in enumerate(line)]
         for r, line in enumerate(parse_day(8))]
trees_t = list(zip(*trees))
edge_visible_trees = set.union(*[edge_visibles(iter(row)) for row in trees])
for col in trees_t:
    edge_visible_trees |= edge_visibles(iter(col))

print("Part 1:", len(edge_visible_trees))


def local_visibility(base_tree):
    r = base_tree.r
    c = base_tree.c
    directions = [
        c - next((tree.c for tree in reversed(trees[r][:c])
                  if tree.height >= base_tree.height), 0),
        -c + next((tree.c for tree in trees[r][c+1:]
                   if tree.height >= base_tree.height), len(trees[r])-1),
        r - next((tree.r for tree in reversed(trees_t[c][:r])
                  if tree.height >= base_tree.height), 0),
        -r + next((tree.r for tree in trees_t[c][r+1:]
                   if tree.height >= base_tree.height), len(trees[c])-1),
        ]
    return directions[0] * directions[1] * directions[2] * directions[3]


trees_iter = (t for line in trees for t in line)
print("Part 2:", max(local_visibility(t) for t in trees_iter))
