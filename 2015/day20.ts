const target = 33100000;
const target_part1 = target / 10;
const target_part2 = target / 11;

function* count(limit: number) {
	for (let i = 1; i <= limit; i++) {
		yield i;
	}
}

for (const i of count(100000000)) {
	let presents = 0;
	for (const j of count(Math.sqrt(i))) {
		const k = i / j;
		if (Number.isInteger(k)) {
			presents += j + k;
		}
	}
	if (presents >= target_part1) {
		console.log("Part 1:", i);
		break;
	}
}

const active = new Map();

for (const i of count(100000000)) {
	let presents = 0;
	for (const j of count(Math.sqrt(i))) {
		const k = i / j;
		if (Number.isInteger(k)) {
			if (active.has(k)) {
				const n = active.get(k);
				if (n > 0) {
					presents += k;
					active.set(k, n - 1);
				}
			} else {
				presents += k;
				active.set(k, 49);
			}
			if (active.has(j)) {
				const n = active.get(j);
				if (n > 0) {
					presents += j;
					active.set(j, n - 1);
				}
			} else {
				presents += j;
				active.set(j, 49);
			}
		}
	}
	if (presents >= target_part2) {
		console.log("Part 2:", i);
		break;
	}
}
