#include <bits/stdc++.h>
#include <cctype>
#include <queue>
#include <sstream>
#include <unordered_set>
#include <vector>
using namespace std;

struct State {
  vector<bool> current;
  int presses;
};

int solve(vector<bool> &target, vector<vector<int>> &buttons) {
  vector<bool> start(target.size(),
                     false); // initially all the switches will be off
  queue<State> q;
  q.push({start, 0});
  unordered_set<vector<bool>> visited;

  while (!q.empty()) {
    State cur = q.front();
    q.pop();
    visited.insert(cur.current);
    if (cur.current == target) {
      return cur.presses;
    }

    for (auto button : buttons) {
      vector<bool> next = cur.current;
      for (auto i : button) {
        // simply toggle the state;
        next[i] = !next[i];
      }

      if (!visited.count(next)) {
        q.push({next, cur.presses + 1});
        visited.insert(next);
      }
    }
  }
  return 0;
}

int main() {
  const string filename = "inputs.txt";

  ifstream file(filename);
  if (!file.is_open()) {
    cerr << "Error Unable to open the file\n";
    return 0;
  }

  string line;
  int total = 0;

  while (getline(file, line)) {
    vector<bool> target; // target holds the boolean of the final switches
    vector<vector<int>> buttons; // the buttons consists of an array of
                                 // acceptable buttons we can press

    stringstream ss(line);
    char ch; // helps us to go char by char and put values into the array

    while (ss >> ch) {
      // take the starting occurance of the [ and go on until u find ] and push
      // them all in to the array once
      if (ch == '[') {
        while (ss >> ch && ch != ']') {
          ch == '#' ? target.push_back(true) : target.push_back(false);
        }
      }
      // handling the buttons
      else if (ch == '(') {
        vector<int> currGroup;
        int num;
        while (ss >> ch) {
          if (isdigit(ch)) {
            ss.putback(ch);
            ss >> num;
            currGroup.push_back(num);
          } else if (ch == ')') {
            break;
          }
        }
        buttons.push_back(currGroup);
      } else if (ch == '{') {
        break;
      }
    }
    auto print = [&]() -> void {
      cout << "Target: [ ";
      for (auto t : target)
        cout << t << " ";
      cout << " ]\n";

      cout << " Buttons: [ ";
      for (auto &v : buttons) {
        cout << "( ";
        for (auto x : v)
          cout << x << ",";
        cout << " )";
      }
      cout << " ]\n";
    };
    int ans = solve(target, buttons);
    total += ans;
  }
  cout << total << "\n";
  return 0;
}
