import { md5 } from 'js-md5';

let part1Found = false

const key = "ckczppom";
for (let i = 1; i < 1000000000; i++) {
  if (!part1Found && md5(key+i).slice(0, 5) === '00000') {
    console.log("Part 1:", i)
    part1Found = true
  }
  if (md5(key+i).slice(0, 6) === '000000') {
    console.log("Part 2:", i)
    break
  }
}
