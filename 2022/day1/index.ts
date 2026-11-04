import * as fs from "node:fs";

const filePath: string = "./inputs.txt";
const contents: string = fs.readFileSync(filePath, 'utf-8');
const lines: string[] = contents.split(/\r?\n/);
let data: number[] = [];

let part1: number[] = [];
let sum: number = 0;
for(const num of lines){
  if (num === ""){
    part1.push(Number(sum));
    sum = 0;
    continue;
  }
  sum += Number(num);
}
part1.sort((a, b) => b - a);
console.log(part1[0]);
let part2: number = 0;
for(let i = 0; i < 3; i++){
  part2 += Number(part1[i]);
}
console.log(part2)
