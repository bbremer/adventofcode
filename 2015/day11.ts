const utf8Encode = new TextEncoder();
const utf8Decode = new TextDecoder();
let passwordStr = "cqjxjnds";
let password = utf8Encode.encode(passwordStr);
let part1Found = false;

while (true) {
  password = increment(password);
  if (valid(password)) {
    if (part1Found) {
      break
    }
    console.log("Part 1:", utf8Decode.decode(password));
    part1Found = true;
  }
}
console.log("Part 2:", utf8Decode.decode(password));

function increment(password) {
  const ZBYTE = "z".charCodeAt(0);
  const ABYTE = "a".charCodeAt(0);
  const INVALIDCHARS = utf8Encode.encode("iol");
  for (let i = password.length-1; i >= 0; i--) {
    if (password[i] === ZBYTE) {
      password[i] = ABYTE;
      continue;
    }
    password[i]++;
    if (password[i] in INVALIDCHARS) {
      password[i]++;
    }
    break;
  }
  return password;
}

function valid(passwordStr) {
  let password = [...passwordStr];
  return (
      password.slice(0, -2)
      .map((_, i) => password.slice(i, i+3))
      .some(x => x[0] + 2 === x[2] && x[1] + 1 === x[2])
    ) && (
      password.slice(0, -1)
      .map((_, i) => password.slice(i, i+2))
      .filter(x => x[0] === x[1])
      .map(x => x[0])
      .flatMap((x, i, arr) => arr.slice(i+1).map(y => [x, y]))
      .some(x => x[0] != x[1])
    )
}
