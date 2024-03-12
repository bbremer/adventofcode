import * as fs from "node:fs";

const textLines = fs
	.readFileSync("inputs/day21.txt", "ascii")
	.trim()
	.split("\n");

const [hp, d, a] = textLines.map((x) => Number(x.split(": ")[1]));
const playerHp = 100;

const weapons: Array<[string, number, number, number]> = [
	["Dagger", 8, 4, 0],
	["Shortsword", 10, 5, 0],
	["Warhammer", 25, 6, 0],
	["Longsword", 40, 7, 0],
	["Greataxe", 74, 8, 0],
];

const armor: Array<[string, number, number, number]> = [
	["Nothing", 0, 0, 0],
	["Leather", 13, 0, 1],
	["Chainmail", 31, 0, 2],
	["Splintmail", 53, 0, 3],
	["Bandedmail", 75, 0, 4],
	["Platemail", 102, 0, 5],
];

const rings: Array<[string, number, number, number]> = [
	["Nothing", 0, 0, 0],
	["Damage +1", 25, 1, 0],
	["Damage +2", 50, 2, 0],
	["Damage +3", 100, 3, 0],
	["Defense +1", 20, 0, 1],
	["Defense +2", 40, 0, 2],
	["Defense +3", 80, 0, 3],
];

function* ringsIter() {
	for (const x of rings) {
		for (const y of rings) {
			if (x !== y) {
				yield [x, y];
			}
		}
	}
}

let minCost = 100000;
let maxCost = 0;

function sim(playerD: number, playerA: number): boolean {
	let hp_ = hp;
	let playerHp_ = playerHp;
	for (let i = 0; i < 1000; i++) {
		hp_ -= playerD - a;
		if (hp_ <= 0) {
			return true;
		}
		playerHp_ -= d - playerA;
		if (playerHp_ <= 0) {
			return false;
		}
	}
	throw new Error();
}

for (const [w, wc, wd, wa] of weapons) {
	for (const [a, ac, ad, aa] of armor) {
		for (const [[r1, r1c, r1d, r1a], [r2, r2c, r2d, r2a]] of ringsIter()) {
			const cost = wc + ac + r1c + r2c;
			const win = sim(wd + r1d + r2d, aa + r1a + r2a);
			if (cost < minCost && win) {
				minCost = cost;
			}
			if (cost > maxCost && !win) {
				maxCost = cost;
			}
		}
	}
}

console.log("Part 1:", minCost);
console.log("Part 2:", maxCost);
