import * as fs from "node:fs"

let part1Total = 0;
let part2Total = 0;

fs.createReadStream("inputs/day8.txt", "ascii").on("data", d => {
  if (Buffer.isBuffer(d)) {
    d = d.toString()
  }
  let lines = d.trim().split("\n")
  part1Total += lines.reduce((t, line) => t + line.length - escape(line), 0)
  part2Total += lines.reduce((t, line) => t + unescape(line) - line.length, 0)
}).on("end", () => console.log("Part 1:", part1Total, "Part 2:", part2Total))

function escape(line) {
  let count = 0;
  for (let i = 1; i < line.length-1; i++) {
    count++;
    if (line[i] != '\\') {
      continue;
    }
    if (line[i+1] == 'x') {
      i += 3
      continue
    }
    i += 1;
  }
  return count;
}

function unescape(line) {
  let count = 0;
  for (let i = 0; i < line.length; i++) {
    count++;
    if (line[i] === '"' || line[i] === '\\') {
      count++;
      continue;
    }
  }
  return count + 2;
}
