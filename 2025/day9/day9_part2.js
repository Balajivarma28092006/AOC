const data = await Bun.file("inputs.txt").text();
const pairStrings = data.trim().split(/\s+/);

// make pairs of strings
const pairs = pairStrings.map((str) => {
  const [x, y] = str.split(",").map(Number);
  return [x, y];
});

let max_area = 0;
for (let i = 0; i < pairs.length; i++) {
  for (let j = i + 1; j < pairs.length; j++) {
    let area =
      (Math.abs(pairs[i][0] - pairs[j][0]) + 1) *
      (Math.abs(pairs[j][1] - pairs[i][1]) + 1);
    if (max_area < area) max_area = area;
  }
}

console.log(max_area);

// part 2
max_area = 0;
for (let i = 0; i < pairs.length; i++) {
  for (let j = i + 1; j < pairs.length; j++) {
    const p1 = pairs[i];
    const p2 = pairs[j];

    // check if the points donot form a straight line
    if (p1[0] == p2[0] || p1[1] == p2[1]) continue;

    // find the left right top bottom boundaries for the given points
    const rMinX = Math.min(p1[0], p2[0]);
    const rMaxX = Math.max(p1[0], p2[0]);
    const rMinY = Math.min(p1[1], p2[1]);
    const rMaxY = Math.max(p1[1], p2[1]);

    let area = (rMaxX - rMinX + 1) * (rMaxY - rMinY + 1);
    if (area <= max_area) continue;

    // the rule says that we should not have any other points inside our rectangle
    // so loop through all the points and check for our boundaries
    let hasInternalVertex = false;
    for (let k = 0; k < pairs.length; k++) {
      const pk = pairs[k];
      if (pk[0] > rMinX && pk[0] < rMaxX && pk[1] > rMinY && pk[1] < rMaxY) {
        hasInternalVertex = true;
        break;
      }
    }
    if(hasInternalVertex) continue;

    // check for polygon edges
    let crossesBoundary = false;
    for (let k = 0; k < pairs.length; k++) {
      const a = pairs[k];
      // next point will be
      const b = pairs[(k + 1) % pairs.length];

      if (a[1] == b[1]) {
        const y = a[1];

        // check if this horizantal edge is inside the vertical range of our rectangle
        if (y > rMinY && y < rMaxY) {
          const edgeMinX = Math.min(a[0], b[0]);
          const edgeMaxX = Math.max(a[0], b[0]);
          if (edgeMaxX > rMinX && edgeMinX < rMaxX) {
            crossesBoundary = true;
            break;
          }
        }
      } else if (a[0] == b[0]) {
        const x = a[0];
        if (x > rMinX && x < rMaxX) {
          const edgeMinY = Math.min(a[1], b[1]);
          const edgeMaxY = Math.max(a[1], b[1]);

          // Does the edge overlap the rectangle's
          if (edgeMaxY > rMinY && edgeMinY < rMaxY) {
            crossesBoundary = true;
            break;
          }
        }
      }
    }
    if (crossesBoundary) continue;

    const centerX = (rMinX + rMaxX) / 2;
    const centerY = (rMinY + rMaxY) / 2;

    let inside = false;

    for (let k = 0, l = pairs.length - 1; k < pairs.length; l = k++) {
      const x1 = pairs[k][0];
      const y1 = pairs[k][1];

      const x2 = pairs[l][0];
      const y2 = pairs[l][1];

      if (
        y1 > centerY !== y2 > centerY &&
        centerX < ((x2 - x1) * (centerY - y1)) / (y2 - y1) + x1
      ) {
        inside = !inside;
      }
    }
    if (!inside) continue;
    max_area = area;
  }
}

console.log(max_area);
