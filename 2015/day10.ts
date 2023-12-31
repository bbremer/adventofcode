var input = "1113122113";
var output = "";

for (let i = 0; i < 50; i++) {
  if (i == 40) {
    console.log("Part 1:", input.length);
  }

  let count = 1;
  let state = input[0];
  [...input.slice(1)].forEach(e => {
    if (e == state) {
      count++;
      return;
    }
    output += `${count}${state}`;
    count = 1;
    state = e;
  })
  output += `${count}${state}`;
  input = output;
  output = "";
}
console.log("Part 2:", input.length);
