import * as fs from "node:fs";

const indicies: Map<string, number> = new Map([..."ABCDEFGMZ"].map((x, i) => [x, i]));
const l = [...indicies.keys()].length;
let happy = new Array(l);
for (let i = 0; i < l; i++) {
  happy[i] = new Array(l);
  for (let j = 0; j < l; j++) {
    happy[i][j] = 0;
  }
}

(
  fs.readFileSync("inputs/day13.txt", "ascii")
  .trim().split("\n")
  .map(line => line.replace(".", "").split(" "))
  .map(x => [x[0], x[2] === "gain" ? "" : "-", x[3], x.at(-1)])
  .map(x => ({s: x[0][0], x: parseInt(x[1] + x[2]), t: x[3][0]}))
  .map(x => [indicies.get(x.s), x.x, indicies.get(x.t)])
  .forEach(x => {
    happy[x[0]][x[2]] += x[1];
    happy[x[2]][x[0]] += x[1];
  })
);

function run(lim) {
  let stack = [{x: "A", s: 0, h: 1}];
  let max = 0;
  while (stack.length > 0) {
    let a = stack.pop()
    const i = indicies.get(a.x);
    if (a.h == lim) {
      let m = a.s + happy[i][0]
      if (m > max) {
        max = m;
      }
      continue;
    }

    indicies.forEach((j, x) => {
      if (i != j && (a.h & 1 << j) == 0) {
        stack.push({x: x, s: happy[i][j]+a.s, h: a.h | (1 << j)})
      }
    })
  }
  return max
}

console.log("Part 1:", run(255));
console.log("Part 2:", run(511));
