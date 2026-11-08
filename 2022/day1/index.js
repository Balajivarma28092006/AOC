import * as fs from "node:fs";
const filePath = "./inputs.txt";
const contents = fs.readFileSync(filePath, 'utf-8');
const lines = contents.split(/\r?\n/);

let part1 = [];
let sum = 0;
for (const num of lines) {
    if (num === "") {
        part1.push(Number(sum));
        sum = 0;
        continue;
    }
    sum += Number(num);
}
part1.sort((a, b) => b - a);
console.log(part1[0]);
let part2 = 0;
for (let i = 0; i < 3; i++) {
    part2 += Number(part1[i]);
}
console.log(part2);
