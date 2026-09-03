const filename = "inputs.txt";

const rawData = await Deno.readTextFile(filename);
const cardValues = [
  "2",
  "3",
  "4",
  "5",
  "6",
  "7",
  "8",
  "9",
  "T",
  "J",
  "Q",
  "K",
  "A",
];

/*
    we have five cards labled: A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, or 2.
*/

const hands = rawData
  .trim()
  .split("\n")
  .map((line) => {
    const [cards, value] = line.split(" ");
    return { cards, bid: Number(value) };
  });

function getHandByStrength(cards) {
  // map all the values counts like a unordered_map in cpp
  const counts = {};
  for (const card of cards) {
    counts[card] = (counts[card] || 0) + 1;
  }

  // sort the frequencies in descending order
  const freq = Object.values(counts).sort((a, b) => b - a);

  // we only have five cards means we have typically need to check only the starting two chars counts
  if (freq[0] === 5) return 7; // a five kind of card;
  if (freq[0] === 4) return 6; // a four kind of card;
  if (freq[0] === 3 && freq[1] === 2) return 5; // three kind of card
  if (freq[0] === 3) return 4; // three face card
  if (freq[0] === 2 && freq[1] === 2) return 3; // two face kind of card
  if (freq[0] === 2) return 2; // One pair
  return 1; // high card
}

hands.sort((a, b) => {
  const typeA = getHandByStrength(a.cards);
  const typeB = getHandByStrength(b.cards);

  if (typeA !== typeB) {
    return typeA - typeB;
  }

  for (let i = 0; i < 5; i++) {
    const valA = cardValues.indexOf(a.cards[i]);
    const valB = cardValues.indexOf(b.cards[i]);

    if (valA !== valB) {
      return valA - valB;
    }
  }
  return 0;
});

const totalWinnings = hands.reduce((sum, hand, index) => {
  const rank = index + 1;
  return sum + hand.bid * rank;
}, 0);

console.log("Part 1: ", Number(totalWinnings));
