import * as fs from "node:fs"

let part1Total = 0;
let part2Total = 0;

fs.createReadStream('inputs/day2.txt', 'ascii').on('data', d => {
  if (Buffer.isBuffer(d)) {
    d = d.toString()
  }
  let lines = (
    d.split('\n')
    .filter(line => line.length > 0)
    .map(line => line.split('x').map(x => parseInt(x)))
  )

  part1Total += lines.reduce((total, line) => total + (
      line
      .flatMap((x, i) => line.slice(i+1).map(y => x*y))
      .flatMap((x, i, arr) => i === 0 ? [Math.min(...arr), x] : [x])
      .reduce((lineTotal, c) => lineTotal + 2*c)
    ), 0)

  part2Total += lines.reduce((total, line) => (
    total
    + line.reduce((a, c) => a*c, 1)
    + Math.min(...line.flatMap((x, i) => line.slice(i+1).map(y => 2*(x+y))))
  ), 0)

}).on('end', () => console.log(`Part 1: ${part1Total}\nPart 2: ${part2Total}`))
