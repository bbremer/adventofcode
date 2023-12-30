import * as fs from "node:fs";

let part1Total = 0;

let funcs = {};
fs.createReadStream("inputs/day7.txt", "ascii").on("data", d => {
  if (Buffer.isBuffer(d)) {
    d = d.toString()
  }

  let lines = d.trim().split('\n')
  lines.forEach(line => {
    let splitLine = line.split(' ')
    funcs[splitLine.at(-1)] = lineToFunc(splitLine)
  })
}).on("end", () => {
  let p1 = source("a", {});
  console.log("Part 1:", p1, "\nPart 2:", source("a", {b: p1}));
})

function lineToFunc(splitLine: Array<string>) {
    if (splitLine[0] === "NOT") {
      return (cs) => ~source(splitLine[1], cs);
    }

    switch (splitLine[1]) {
      case "->":
        return (cs) => source(splitLine[0], cs)
      case "AND":
        return (cs) => source(splitLine[0], cs) & source(splitLine[2], cs)
      case "OR":
        return (cs) => source(splitLine[0], cs) | source(splitLine[2], cs)
      case "LSHIFT":
        return (cs) => source(splitLine[0], cs) << parseInt(splitLine[2])
      case "RSHIFT":
        return (cs) => source(splitLine[0], cs) >> parseInt(splitLine[2])
    }
}

function source(src, circuits) {
  let i = parseInt(src);
  if (!isNaN(i)) {
    return i
  }

  if (!(src in circuits)) {
    circuits[src] = funcs[src](circuits)
  }
  return circuits[src]
}
