import * as fs from "node:fs";

const textLines = fs
	.readFileSync("inputs/day19.txt", "ascii")
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

const reverseSymbolMapRaw = [...symbolMap.entries()].flatMap(([k, v]) =>
	v.map((x) => [x, k]),
);
const reverseSymbolMap = reverseSymbolMapRaw.filter(([k, v]) => v !== "e");
const compoundsToE = new Set(
	reverseSymbolMapRaw.filter(([k, v]) => v === "e").map(([k, _]) => k),
);

const cache = new Map();

function dfs(s: string, count: number): number {
	if (compoundsToE.has(s)) {
		return count + 1;
	}
	if (cache.has(s)) {
		return cache.get(s);
	}
	let localMin = 1000000000 + 1;
	for (let i = 0; i < s.length; i++) {
		for (const [k, v] of reverseSymbolMap) {
			const j = i + k.length;
			if (s.slice(i, j) === k) {
				const newS = s.slice(0, i) + v + s.slice(j);
				localMin = Math.min(localMin, dfs(newS, count + 1));
			}
		}
	}
	cache.set(s, localMin);
	return localMin;
}

class CachedString {
	constructor(
		public s: string,
		public n: number,
	) {}
}

const symbols = [...toSymbols(baseString)].map((x) => x[0]);
let currOptions: CachedString[] = [new CachedString(symbols[0], 0)];
const MAXSIZE = 16;
for (const symbol of symbols.slice(1)) {
	const nextOptions: Set<CachedString> = new Set();
	for (const opt of currOptions) {
		for (const newOpt of combine(opt, symbol)) {
			nextOptions.add(newOpt);
		}
	}
	currOptions = [...nextOptions];
	console.log(symbol, currOptions);
}

function combine(cs: CachedString, symbol: string): CachedString[] {
	const newString = cs.s + symbol;
	const ret = [new CachedString(newString, cs.n)];
	for (const [k, v] of reverseSymbolMap) {
		if (newString.slice(-k.length) === k) {
			ret.push(new CachedString(newString.slice(0, -k.length) + v, cs.n + 1));
		}
	}
	return ret.filter((cs) => cs.s.length <= MAXSIZE);
}

console.log();
console.log(reverseSymbolMap);
console.log(currOptions.filter((x) => compoundsToE.has(x)));
console.log(compoundsToE);

// console.log("Part 2:", dfs(baseString, 0));
