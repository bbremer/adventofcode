import * as fs from 'node:fs';

let part1Total = 0;
let part2Total = 0

fs.createReadStream('inputs/day3.txt', 'ascii').on('data', rawLine => {
  if (Buffer.isBuffer(rawLine)) {
    rawLine = rawLine.toString();
  }

  let line = [...rawLine.trim()];

  let s = new Set([[0, 0].toString()]);
  run(line, s);
  part1Total = s.size;

  let s2 = new Set([[0, 0].toString()]);
  let line1 = [];
  let line2 = [];
  line.forEach((e, i) => i % 2 == 0 ? line1.push(e) : line2.push(e))
  run(line1, s2);
  run(line2, s2);
  part2Total = s2.size;
}).on('end', () => console.log(`Part 1: ${part1Total}\nPart 2: ${part2Total}`));

function run(line: Array<string>, s: Set<string>) {
  let x = 0;
  let y = 0;
  line.forEach(c => {
    switch (c) {
      case '^':
        x--;
        break;
      case 'v':
        x++;
        break;
      case '>':
        y++;
        break;
      case '<':
        y--;
        break;
      default:
        console.log(c)
        throw new Error();
    }
    let e = [x, y].toString()
    if (!s.has(e)) {
      s.add(e)
    }
  })
}
