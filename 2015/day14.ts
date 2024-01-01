import * as fs from "node:fs"

let regex = /^.* can fly (\d*) km\/s for (\d*) seconds, but then must rest for (\d*) seconds.$/

const TOTALDURATION = 2503;

const reindeerLists = (
  fs.readFileSync("inputs/day14.txt", "ascii")
  .trim().split("\n")
  .map(x => regex.exec(x).slice(1, 4).map(y => parseInt(y)))
)

console.log("Part 1:", (
  reindeerLists.map(([s, d, r]) => {
    let quotient = Math.floor(TOTALDURATION / (d + r));
    let remainder = TOTALDURATION % (d + r);
    let movingRemainder = Math.min(remainder, d);
    return s * (quotient*d + movingRemainder);
  })
  .reduce((max, x) => x > max ? x : max)
))

class Reindeer {
  score
  currDistance
  moving
  remaining
  constructor(private s, private d, private r) {
    this.score = 0;
    this.currDistance = 0;
    this.moving = true;
    this.remaining = d;
  }

  update() {
    if (this.moving) {
      this.currDistance += this.s;
    }
    this.remaining--;
    if (this.remaining == 0) {
      this.moving = !this.moving;
      this.remaining = this.moving ? this.d : this.r;
    }
  }
}

let reindeer = reindeerLists.map(([s, d, r]) => new Reindeer(s, d, r))
for (let i = 0; i < TOTALDURATION; i++) {
  reindeer.forEach(r => r.update());
  reindeer.slice(1).reduce((maxRs, r) => r.currDistance > maxRs[0].currDistance ? [r] : (
    r.currDistance == maxRs[0].currDistance ? maxRs.concat(r) : maxRs), [reindeer[0]]).forEach(r => r.score++);
}
console.log("Part 2:", reindeer.reduce((max, r) => r.score > max.score ? r : max).score);
