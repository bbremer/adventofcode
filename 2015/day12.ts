import * as fs from "node:fs";

const json = JSON.parse(fs.readFileSync("inputs/day12.txt", "ascii"));

let part1Total = 0;
let part2Total = 0;
let stack = [[json, false]];
while (stack.length > 0) {
  let [x, red] = stack.pop();
  switch (typeof x) {
    case "number": 
      part1Total += x
      if (!red) {
        part2Total += x
      }
      break;
    case "string":
      break;
    case "object":
      let newRed = red || !Array.isArray(x) && Object.values(x).some(x => x === "red");
      for (const y of Object.values(x)) {
        stack.push([y, newRed]);
      }
      break;
    default:
      throw new Error();
  }
}
console.log("Part 1:", part1Total);
console.log("Part 2:", part2Total);
