import * as fs from "node:fs";

let part1Total = 0;
let part2Total = 0;

const vowels = new Set("aeiou");

fs.createReadStream("inputs/day5.txt", "ascii").on("data", e => {
  if (Buffer.isBuffer(e)) {
    e = e.toString();
  }
  
  let lines = e.trim().split("\n")

  part1Total += lines.reduce((total, line) => total + ([
      [...line].reduce((c, e) => vowels.has(e) ? c+1 : c, 0) >= 3,
      [...line.slice(0, -1)].some((e, i) => e == line[i+1]),
      !["ab", "cd", "pq", "xy"].some(ss => line.includes(ss)),
    ].every(x => x) ? 1 : 0), 0)

  part2Total += lines.reduce((total, line) => total + (
    overlap(line) && [...line.slice(0, -2)].some((e, i) => e == line[i+2]) ? 1 : 0
  ), 0)
}).on("end", () => console.log("Part 1:", part1Total, "\nPart 2:", part2Total));

function overlap(line: string) : boolean  {
  let pairs = {};
  return [...line.slice(0, -1)].map((_, i) => line.slice(i, i+2)).some((e, i) => {
    if (e in pairs) {
      if (pairs[e] < i-1) {
        return true
      }
    } else {
      pairs[e] = i;
    }
    return false
  })
}
