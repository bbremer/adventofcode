import * as fs from "node:fs";

const standard = new Map(Object.entries({
  children: 3,
  cats: 7,
  samoyeds: 2,
  pomeranians: 3,
  akitas: 0,
  vizslas: 0,
  goldfish: 5,
  trees: 3,
  cars: 2,
  perfumes: 1,
}));

const regex = /^Sue (\d+): (.*): (\d+), (.*): (\d+), (.*): (\d+)$/;

let sues = (
  fs.readFileSync("inputs/day16.txt", "ascii")
  .trim().split("\n")
  .map(m => regex.exec(m).slice(1, 8))
  .map(m => ({i: m[0], j: m.slice(1).reduce((r, v, i, arr) => i % 2 === 0 ? r.concat([arr.slice(i, i+2)]) : r, [])}))
  // .forEach((_, i) => console.log("Part 1:", i+1))
)

console.log("Part 1:", sues.filter(m => (m.j.every(x => standard.get(x[0]) === parseInt(x[1]))))[0].i)
console.log("Part 2:", sues.filter(m => (m.j.every(compare)))[0].i)

function compare([key, x]) {
  switch (key) {
    case "cats":
    case "trees":
      return standard.get(key) < parseInt(x)
    case "pomeranians":
    case "goldfish":
      return standard.get(key) > parseInt(x)
    default:
      return standard.get(key) === parseInt(x)
  }
}
