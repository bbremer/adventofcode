import os

import requests

input_url_template = "https://adventofcode.com/2022/day/{}/input"
inputs_path = os.path.dirname(__file__) + '/inputs'
for i in range(1, 26):
    input_fn = f"{inputs_path}/day{i}.txt"
    if os.path.exists(input_fn):
        continue

    input_url = input_url_template.format(i)

    with open('cookie.txt') as f:
        cookie = f.read().strip()
    headers = {"Cookie": cookie}
    r = requests.get(input_url, headers=headers)

    if r.status_code == 404:
        print(f"Day {i} input not found. Breaking.")
        break

    r.raise_for_status()

    with open(input_fn, 'w') as f:
        f.write(r.text)
