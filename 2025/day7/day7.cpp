#include <bits/stdc++.h>

using namespace std;

int main() {
  ifstream file("day7_inputs.txt");
  vector<string> grid;
  string line;

  while (getline(file, line)) {
    grid.push_back(line);
  }

  int rows = grid.size();
  int cols = grid[0].size();

  int start_col = -1;
  for (int j = 0; j < cols; j++) {
    if (grid[0][j] == 'S') {
      start_col = j;
      break;
    }
  }

  set<pair<int, int>> visited;
  int count = 0;

  // min-heap priority queue: smallest (row, col) comes out first
  priority_queue<pair<int, int>, vector<pair<int, int>>,
                 greater<pair<int, int>>>
      pq;

  pq.push({1, start_col});

  while (!pq.empty()) {
    auto [row, col] = pq.top();
    pq.pop();

    // bounds check
    if (row < 0 || row >= rows || col < 0 || col >= cols)
      continue;

    // visited check
    if (visited.count({row, col}))
      continue;

    visited.insert({row, col});
    char curr = grid[row][col];

    if (curr == '^') {
      count++;
      pq.push({row + 1, col - 1}); // down-left
      pq.push({row + 1, col + 1}); // down-right
    } else if (curr == '.') {
      pq.push({row + 1, col}); // straight down
    }
  }

  cout << "Part ones solution is " << count << '\n';
  return 0;
}
