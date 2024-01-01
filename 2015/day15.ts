import * as fs from "node:fs";

const regex = /^.*: capacity (-?\d+), durability (-?\d+), flavor (-?\d+), texture (-?\d+), calories (-?\d+)$/

const ingredients = (
  fs.readFileSync("inputs/day15.txt", "ascii")
  .trim().split("\n")
  .map(line => regex.exec(line).slice(1, 6).map(x => parseInt(x))) 
)

let part1Max = 0;
let part2Max = 0;
for (const teaspoons of teaspoonGen()) {
  const scores = (
    ingredients
    .map((x, i) => x.map(y => y * teaspoons[i]))
    .reduce((s, x) => s.map((y, i) => y + x[i]))
    .map(x => x > 0 ? x : 0)
  )
  let score = scores.slice(0, -1).reduce((p, x) => p*x);
  if (score > part1Max) {
    part1Max = score;
  }
  if (scores.at(-1) == 500 && score > part2Max) {
    part2Max = score;
  }
}
console.log("Part 1:", part1Max)
console.log("Part 2:", part2Max)

function* teaspoonGen() {
  const lim = 100;
  for (const i of range(0, lim)) {
    for (const j of range(0, lim-i)) {
      for (const k of range(0, lim-i-j)) {
        yield [i, j, k, lim-i-j-k]
      }
    }
  }
}

function* range(min, max) {
  for (let i = min; i < max; i++) {
    yield i
  }
}
