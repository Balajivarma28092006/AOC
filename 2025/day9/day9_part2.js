const data = await Bun.file("inputs.txt").text();
const pairStrings = data.trim().split(/\s+/);

// make pairs of strings
const pairs = pairStrings.map(str => {
	const [x, y] = str.split(",").map(Number);
	return [x, y];
});

let max_area = 0;
for (let i = 0; i < pairs.length; i++){
	for(let j = i + 1; j < pairs.length; j++){
		let area = (Math.abs(pairs[i][0] - pairs[j][0]) + 1 ) 
			* (Math.abs(pairs[j][1] - pairs[i][1]) + 1);
		if (max_area < area) max_area = area;
	}
}

console.log(max_area);