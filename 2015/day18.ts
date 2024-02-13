import * as fs from "node:fs";

const textLines = fs
	.readFileSync("inputs/day18.txt", "ascii")
	.trim()
	.split("\n");

let lines = textLines.map((l) => [...l]);
let lines2 = textLines.map((l) => [...l]);

const MAXR = lines.length;
const MAXC = lines[0].length;

const BASENEIGHBORINDICIES: Array<[number, number]> = [
	[-1, -1],
	[-1, 0],
	[-1, 1],
	[0, -1],
	[0, 1],
	[1, -1],
	[1, 0],
	[1, 1],
];

function run(grid: string[][], constOn: (r: number, c: number) => boolean) {
	return grid.map((line, r) =>
		line.map((e, c) => (constOn(r, c) ? "#" : onFromNeighbors(grid, e, r, c))),
	);
}

function onFromNeighbors(grid: string[][], e: string, r, c: number): string {
	const n = getNeighborIndicies(r, c).filter(
		([nr, nc]) => grid[nr][nc] === "#",
	).length;
	return e === "#" ? onTransform(n) : offTransform(n);
}

function cornerLight(r: number, c: number): boolean {
	return (r === 0 || r === MAXR - 1) && (c === 0 || c === MAXC - 1);
}

function onTransform(len: number): string {
	return len === 2 || len === 3 ? "#" : ".";
}

function offTransform(len: number): string {
	return len === 3 ? "#" : ".";
}

function getNeighborIndicies(br: number, bc: number) {
	return BASENEIGHBORINDICIES.map(([dr, dc]) => [br + dr, bc + dc]).filter(
		([r, c]) => 0 <= r && r < MAXR && 0 <= c && c < MAXC,
	);
}

for (let i = 0; i < 100; i++) {
	lines = run(lines, (r, c) => false);
}
console.log(
	"Part 1:",
	lines.reduce((a, l) => a + l.filter((e) => e === "#").length, 0),
);

for (let i = 0; i < 100; i++) {
	lines2 = run(lines2, cornerLight);
}
console.log(
	"Part 2:",
	lines2.reduce((a, l) => a + l.filter((e) => e === "#").length, 0),
);
