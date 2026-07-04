#include <bits/stdc++.h>

using namespace std;
using ll = long long;

ll dfs2(int row, int col, vector<string>& grid, vector<vector<ll>>& memo, int rows, int cols){
    if (col < 0 || col >= cols) return 0;
    if(row == rows) return 1;
    if(memo[row][col] != -1) return memo[row][col];

    char curr = grid[row][col];
    ll ans;
    if (curr == 'S' || curr == '|' || curr == '.'){
        ans = dfs2(row + 1, col, grid, memo, rows, cols);
    }else if(curr == '^'){
        ans = dfs2(row + 1, col + 1, grid, memo, rows, cols) + dfs2(row + 1, col - 1, grid, memo, rows, cols);
    }else{
        ans = 0;
    }
    return memo[row][col] = ans;
}


int main() {
  ifstream file("day7_inputs.txt");
  vector<string> grid;
  string line;

  while (getline(file, line)) {
    grid.push_back(line);
  }

  int rows = grid.size();
  int cols = grid[0].size();
  vector<vector<ll>> memo;
  memo.assign(rows, vector<ll>(cols, -1));

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
  ll count1= dfs2(1, start_col, grid, memo, rows, cols);
  cout << "Part two solution is " << count1 << '\n';
  return 0;
}
