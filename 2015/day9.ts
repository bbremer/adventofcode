import * as fs from "node:fs"

let part1Total = 0;
let part2Total = 0;

let cities = {};

fs.createReadStream("inputs/day9.txt", "ascii").on("data", d => {
  if (Buffer.isBuffer(d)) {
    d = d.toString()
  }

  d.trim().split('\n').forEach(line => {
    let [c1, _, c2, _2, dRaw] = line.split(" ");
    let d = parseInt(dRaw);
    [c1, c2].forEach(c => {if (!(c in cities)) cities[c] = []})
    cities[c1].push({dest: c2, dist: d})
    cities[c2].push({dest: c1, dist: d})
  })
}).on("end", () => {
  let [min, max] = dfs();
  console.log("Part 1:", min)
  console.log("Part 2:", max)
})

function dfs() {
  let min = 100000000
  let max = 0
  let cityIs = Object.keys(cities).reduce((o, n, i) => ({...o, [n]: 1 << i}), {})
  let target = (1 << Object.keys(cities).length) - 1

  let stack = Object.keys(cities).map(n => ({c: n, d: 0, h: cityIs[n]}))

  while (stack.length > 0) {
    let x = stack.pop()

    if (x.h === target) {
      min = Math.min(x.d, min)
      max = Math.max(x.d, max)
    }

    stack.push(...cities[x.c].filter(e => (cityIs[e.dest] & x.h) === 0)
               .map(e => ({c: e.dest, d: x.d+e.dist, h: x.h | cityIs[e.dest]}))
    )
  }
  return [min, max]
}
