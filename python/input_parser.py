import os

PYTHON_DIR = os.path.dirname(__file__)
INPUTS_DIR = os.path.dirname(PYTHON_DIR) + "/inputs"


def parse_day(day: int):
    with open(f"{INPUTS_DIR}/day{day}.txt") as f:
        for line in f:
            yield line.strip()
