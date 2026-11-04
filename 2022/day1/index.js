import * as fs from "node:fs";
function readFile(fileName) {
    const fileData = fs.readFileSync(fileName, 'utf-8');
    return fileData;
}
console.log(readFile("./test.txt"));
