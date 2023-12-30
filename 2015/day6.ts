import * as fs from "node:fs";

let part1Total = 0;
let part2Total = 0;

let lights = new Array(1000);
for (let i = 0; i < 1000; i++) {
  lights[i] = new Array(1000);
  for (let j = 0; j < 1000; j++) {
    lights[i][j] = false
  }
}
let p2Lights = new Array(1000);
for (let i = 0; i < 1000; i++) {
  p2Lights[i] = new Array(1000);
  for (let j = 0; j < 1000; j++) {
    p2Lights[i][j] = 0;
  }
}

fs.createReadStream("inputs/day6.txt", "ascii").on("data", d => {
  if (Buffer.isBuffer(d)) {
    d = d.toString()
  }

  let lines = d.trim().split("\n")
  // lines = ["turn on 0,0 through 999,999", "turn off 499,499 through 500,500",
  //         "toggle 0,0 through 999,0"]

  lines.map(line => line.split(' ')).forEach(splitLine => {
    let [x0, y0] = splitLine.at(-3).split(',').map(x => parseInt(x))
    let [x1, y1] = splitLine.at(-1).split(',').map(x => parseInt(x))
    let f, f2;
    if (splitLine[0] === "toggle") {
      f = l => !l;
      f2 = l => l+2;
    } else if (splitLine[1] === "on") {
      f = l => true;
      f2 = l => l+1;
    } else if (splitLine[1] === "off") {
      f = l => false;
      f2 = l => l === 0 ? 0 : l-1;
    }
    for (let i = x0; i <= x1; i++ ) {
      for (let j = y0; j <= y1; j++) {
        lights[i][j] = f(lights[i][j])
        p2Lights[i][j] = f2(p2Lights[i][j])
      }
    }
  })

  part1Total += lights.reduce((a, ls) => a + ls.reduce((b, l) => b+l, 0), 0)
  part2Total += p2Lights.reduce((a, ls) => a + ls.reduce((b, l) => b+l, 0), 0)
}).on("end", () => console.log("Part 1", part1Total, "\nPart 2", part2Total));
