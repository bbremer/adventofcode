import * as fs from "node:fs";

const textLines = fs
	.readFileSync("inputs/day19.txt", "ascii")
	// .readFileSync("inputs/test.txt", "ascii")
	.trim()
	.split("\n");

const symbolMap = new Map();

for (const line of textLines) {
	if (line === "") break;
	const [src, dest] = line.split(" => ");
	symbolMap.has(src)
		? symbolMap.get(src).push(dest)
		: symbolMap.set(src, [dest]);
}

const baseString = textLines[textLines.length - 1];

const part1Set = new Set();

for (const [symbol, i, j] of toSymbols(baseString)) {
	if (!symbolMap.has(symbol)) {
		continue;
	}
	for (const newSymbol of symbolMap.get(symbol)) {
		part1Set.add(baseString.slice(0, i) + newSymbol + baseString.slice(j));
	}
}

console.log("Part 1:", part1Set.size);

function* toSymbols(
	s: string,
): Generator<[string, number, number], undefined, undefined> {
	let j = 0;
	for (let i = 0; i < baseString.length; i = j) {
		if (i + 1 < baseString.length && isLower(baseString[i + 1])) {
			j = i + 2;
		} else {
			j = i + 1;
		}
		yield [baseString.slice(i, j), i, j];
	}
}

function isLower(s: string): boolean {
	return s.toLowerCase() === s;
}

const reverseSymbolMap = new Map<string, string>(
	[...symbolMap.entries()]
		.filter(([x, y]) => x !== "e")
		.flatMap(([x, y]) => y.map((z) => [z, x])),
);

const tokenPattern = new RegExp([...symbolMap.keys()].join("|"), "g");

const s = baseString
	.replaceAll("Rn", "(")
	.replaceAll("Y", ",")
	.replaceAll("Ar", ")")
	.replaceAll(tokenPattern, "x")
	.replaceAll("C", "x");

console.log(
	"Part 2:",
	s.length - 2 * s.match(/\(/g).length - 2 * s.match(/,/g).length - 1,
);
