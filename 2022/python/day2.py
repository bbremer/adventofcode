from enum import IntEnum
from functools import reduce

from input_parser import parse_day


Shape = IntEnum('Shape', 'ROCK PAPER SCISSORS')

str_to_shape = {
        'A': Shape.ROCK, 'X': Shape.ROCK,
        'B': Shape.PAPER, 'Y': Shape.PAPER,
        'C': Shape.SCISSORS, 'Z': Shape.SCISSORS}

Outcome = IntEnum('Outcome', {'LOSS': 0, 'DRAW': 3, 'WIN': 6})

str_to_outcome = {'X': Outcome.LOSS, 'Y': Outcome.DRAW, 'Z': Outcome.WIN}

winner_to_loser = {Shape.ROCK: Shape.SCISSORS, Shape.PAPER: Shape.ROCK,
                   Shape.SCISSORS: Shape.PAPER}
loser_to_winner = {v: k for k, v in winner_to_loser.items()}


def score_part1(their_entry: str, my_entry: str) -> int:
    their_shape = str_to_shape[their_entry]
    my_shape = str_to_shape[my_entry]
    if my_shape == their_shape:
        return my_shape + Outcome.DRAW
    if winner_to_loser[my_shape] == their_shape:
        return my_shape + Outcome.WIN
    return my_shape + Outcome.LOSS


def score_part2(their_entry: str, outcome_str: str) -> int:
    their_shape = str_to_shape[their_entry]
    outcome = str_to_outcome[outcome_str]
    match outcome:
        case Outcome.LOSS:
            my_shape = winner_to_loser[their_shape]
        case Outcome.DRAW:
            my_shape = their_shape
        case Outcome.WIN:
            my_shape = loser_to_winner[their_shape]
    return my_shape + outcome


split_lines = (line.split() for line in parse_day(2))
score_iters = ((score_part1(*sl), score_part2(*sl)) for sl in split_lines)
part1_score, part2_score = reduce(lambda x, y: (x[0]+y[0], x[1]+y[1]),
                                  score_iters, (0, 0))
print("Part 1:", part1_score)
print("Part 2:", part2_score)
