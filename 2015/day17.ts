import * as fs from "node:fs"

let numbers = fs.readFileSync("inputs/day17.txt", "ascii").trim().split('\n').map(x => parseInt(x))

const target = 150;

// const target = 25;
// numbers = [20, 15, 10, 5, 5];
// numbers = [20, 5, 5];
let stack = numbers.map((x, i) => ({currSum: x, options: numbers.slice(i+1), history: [x]}))

let part1Total = 0;
let min = 1000000;
let minWays = 0;
while (stack.length > 0) {
  let {currSum, options, history} = stack.pop();
  if (currSum === target) {
    part1Total++;
    if (history.length === min) {
      minWays++;
    } else if (history.length < min) {
      min = history.length;
      minWays = 1;
    }
    continue;
  }
  if (currSum > target || options.length === 0) {
    continue;
  }
  options.forEach((x, i) => {
    stack.push({currSum: currSum+x, options: options.slice(i+1), history: history.concat(x)});
  })
}
console.log("Part 1:", part1Total);
console.log("Part 2:", minWays);
