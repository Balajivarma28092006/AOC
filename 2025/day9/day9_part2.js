const fs = require('fs');
const path = require('path');

const filePath = path.join(__dirname, 'test.txt');
const rawContent = fs.readFileSync(filePath, 'utf-8');

const pairs = rawContent
	.trim()
	.split(/\s+/)
	.map(pair => {
		const [x, y] = pair.split(',').map(Number);
		return {x, y, area: x * y};
	});

const maxArea = pairs.reduce((max, current) => {
	return current.area > max.area ? current : max;
}, pairs[0]);

console.log('All Pairs with areas: ', pairs)
console.log('Pair with maximum area: ', maxArea);

