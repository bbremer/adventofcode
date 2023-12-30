import * as fs from "node:fs";

let stories = 0;
let basementFound = false;

fs.createReadStream('inputs/day1.txt', 'ascii').on('data', d => {
  if (Buffer.isBuffer(d)) {
    d = d.toString()
  }
  [...d].forEach((e, i) => {
    switch (e) {
      case '(':
        stories++;
        break;
      case ')':
        stories--;
        break;
      default:
        throw new Error(`'${e}' is not ( or )`)
    }
    if (!basementFound && stories === -1) {
      basementFound = true
      console.log(`Part 2: ${i+1}`)
    }
  });
}).on('end', () => console.log(`Part 1: ${stories}`))
