const pathTest = "test.txt";
const pathFinal = "inputs.txt";

const file = Bun.file(pathFinal);
const data = await file.text();

const parseGrid = data
  .trim()
  .split("\n")
  .map((line) => line.split(""));

let Part1 = 0;
const seen = new Set();

const directions = [
  [-1, -1],
  [-1, 0],
  [-1, 1],
  [0, -1],
  [0, 1],
  [1, -1],
  [1, 0],
  [1, 1],
];

// now just go through all the values and 8 directions to check if i have a number near
for (let r = 0; r < parseGrid.length; r++) {
  for (let c = 0; c < parseGrid[r].length; c++) {
    const cell = parseGrid[r][c];
    if (cell === "." || /\d/.test(cell)) continue;

    for (const [dr, dc] of directions) {
      const nr = r + dr;
      const nc = c + dc;

      if (
        nr < 0 ||
        nr >= parseGrid.length ||
        nc < 0 ||
        nc >= parseGrid[0].length
      ) {
        continue;
      }

      if (!/\d/.test(parseGrid[nr][nc])) continue;

      // lets find the start number
      let start = nc;
      while (start > 0 && /\d/.test(parseGrid[nr][start - 1])) {
        start--;
      }

      // avoid counting the same number twice
      const key = `${nr},${start}`;
      //   console.log(key);
      if (seen.has(key)) continue;
      seen.add(key);

      // read the whole number starting from the end
      let end = start;
      let num = "";
      while (end < parseGrid[0].length && /\d/.test(parseGrid[nr][end])) {
        num += parseGrid[nr][end];
        end++;
      }

      Part1 += Number(num);
    }
  }
}

console.log(Part1);

let Part2 = 0;
for (let r = 0; r < parseGrid.length; r++) {
  for (let c = 0; c < parseGrid[r].length; c++) {
    // parse only when we saw a gear
    const current_char = parseGrid[r][c];
    if (current_char !== "*") {
      continue;
    }

    const nums = [];
    const seen = new Set();

    for (const [dr, dc] of directions) {
      const nr = r + dr;
      const nc = c + dc;

      if (
        nr < 0 ||
        nr >= parseGrid.length ||
        nc < 0 ||
        nc >= parseGrid[0].length
      ) {
        continue;
      }

      // only process if the neighbour is a digit
      if (!/\d/.test(parseGrid[nr][nc])) continue;

      // go to the start of the number
      let start = nc;
      while (start > 0 && /\d/.test(parseGrid[nr][start - 1])) {
        start--;
      }

      const key = `${nr},${start}`;
      if (seen.has(key)) continue;
      seen.add(key);

      let end = start;
      let num = "";

      while (end < parseGrid[0].length && /\d/.test(parseGrid[nr][end])) {
        num += parseGrid[nr][end];
        end++;
      }
      nums.push(Number(num));
    }
    if (nums.length == 2) {
      Part2 += nums[0] * nums[1];
    }
  }
}

console.log(Part2);
