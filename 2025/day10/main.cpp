#include <bits/stdc++.h>
#include <cctype>
#include <cstddef>
#include <queue>
#include <sstream>
#include <string>
#include <unordered_set>
#include <vector>
using namespace std;

struct State
{
  vector<bool> current;
  int presses;
};

// for part2 we need to keep track of the switches as well as joltages
struct StatePart2
{
  vector<bool> target;
  vector<int> curr_joltages;
  int presses;
};

struct VectorBoolHash
{
  size_t operator()(const vector<bool> &values) const noexcept
  {
    size_t hash = 0;
    for (bool value : values)
    {
      hash = (hash << 1) ^ static_cast<size_t>(value);
    }
    return hash;
  }
};

struct Part2Hash
{
  size_t operator()(const pair<vector<bool>, vector<int>> &state) const noexcept
  {
    std::size_t hash = 0;
    for (bool value : state.first)
    {
      hash = (hash << 1) ^ static_cast<size_t>(value);
    }
    for (int val : state.second)
    {
      hash ^= hash << 13;
      hash ^= hash >> 7;
      hash ^= hash << 17;
      hash += static_cast<size_t>(val);
    }
    return hash;
  }
};

int solve(const vector<bool> &target, const vector<vector<int>> &buttons)
{
  vector<bool> start(target.size(),
                     false); // initially all the switches will be off
  queue<State> q;
  q.push({start, 0});
  unordered_set<vector<bool>, VectorBoolHash> visited;
  visited.insert(start);
  while (!q.empty())
  {
    State cur = q.front();
    q.pop();
    if (cur.current == target)
    {
      return cur.presses;
    }

    for (auto button : buttons)
    {
      vector<bool> next = cur.current;
      for (int i : button)
      {
        // simply toggle the state;
        if (i >= 0 && static_cast<size_t>(i) < next.size())
        {
          next[i] = !next[i];
        }
      }

      if (!visited.count(next))
      {
        q.push({next, cur.presses + 1});
        visited.insert(next);
      }
    }
  }
  return 0;
}

int solvePart2(const vector<bool> &target, const vector<vector<int>> &buttons,
               const vector<vector<int>> &joltages)
{
  if (joltages.empty() || buttons.empty())
    return 0;
  size_t max_joltSize = 0;
  for (const auto &row : joltages)
  {
    max_joltSize = max(max_joltSize, row.size());
  }
  vector<int> start_joltages(max_joltSize, 0);
  vector<bool> start_switches(target.size(), false);

  queue<StatePart2> q;
  q.push({start_switches, start_joltages, 0});

  unordered_set<pair<vector<bool>, vector<int>>, Part2Hash> visited;
  visited.insert({start_switches, start_joltages});

  while (!q.empty())
  {
    StatePart2 curr = q.front();
    q.pop();

    // if (curr.curr_joltages //?)
    // {
    //   return curr.presses;
    // }

    // update the switches similar to the part 1
    for (size_t i = 0; i < buttons.size(); i++)
    {
      vector<bool> next_switches = curr.target;
      for (int idx : buttons[i])
      {
        if (idx >= 0 && static_cast<size_t>(idx) < next_switches.size())
        {
          next_switches[idx] = !next_switches[idx];
        }
      }

      // update the joltages according the usage of the buttons
      vector<int> next_joltages = curr.curr_joltages;
      if (i < joltages.size())
      {

        for (size_t j = 0; j < next_joltages.size() && j < joltages[i].size(); ++j)
        {
          next_joltages[j] += joltages[i][j];
        }
      }

      pair<vector<bool>, vector<int>> next_state = {next_switches, next_joltages};
      if (!visited.count(next_state))
      {
        q.push({next_switches, next_joltages, curr.presses + 1});
        visited.insert(next_state);
      }
    }
  }
  return 0;
}

int main()
{
  const string filename = "test.txt";

  ifstream file(filename);
  if (!file.is_open())
  {
    cerr << "Error Unable to open the file\n";
    return 0;
  }

  string line;
  int total1 = 0;
  int total2 = 0;

  while (getline(file, line))
  {
    vector<bool> target;         // target holds the boolean of the final switches
    vector<vector<int>> buttons; // the buttons consists of an array of acceptable buttons we can press
    vector<vector<int>> joltages;
    stringstream ss(line);
    char ch; // helps us to go char by char and put values into the array

    while (ss >> ch)
    {
      // take the starting occurance of the [ and go on until u find ] and push
      // them all in to the array once
      if (ch == '[')
      {
        while (ss >> ch && ch != ']')
        {
          if (ch == '#' || ch == '.')
          {
            target.push_back(ch == '#');
          }
        }
      }
      // handling the buttons
      else if (ch == '(')
      {
        vector<int> currGroup;
        int num;
        while (ss >> ch)
        {
          if (isdigit(ch))
          {
            ss.putback(ch);
            ss >> num;
            currGroup.push_back(num);
          }
          else if (ch == ')')
          {
            break;
          }
        }
        buttons.push_back(currGroup);
      }
      else if (ch == '{')
      {
        vector<int> current;
        int num;
        while (ss >> ch)
        {
          if (isdigit(static_cast<unsigned char>(ch)))
          {
            ss.putback(ch);
            ss >> num;
            current.push_back(num);
          }
          else if (ch == '}')
          {
            break;
          }
        }
        joltages.push_back(current);
      }
    }
    auto print = [](const vector<bool> &target,
                    const vector<vector<int>> &buttons,
                    const vector<vector<int>> &joltages)
    {
      cout << "Target: [ ";
      for (bool t : target)
        cout << t << " ";
      cout << " ]\n";

      cout << " Buttons: [ ";
      for (auto &v : buttons)
      {
        cout << "( ";
        for (auto x : v)
          cout << x << ",";
        cout << " )";
      }
      cout << " ]\n";

      cout << "Joltages:[ ";
      for (const auto &v : joltages)
      {
        cout << "{";
        for (auto x : v)
        {
          cout << x << ", ";
        }
        cout << "}";
      }
      cout << " ]\n";
    };

    // print(target, buttons, joltages);

    int ans = solve(target, buttons);
    total1 += ans;

    ans = solvePart2(target, buttons, joltages);
    total2 += ans;
  }
  cout << total1 << "\n";
  cout << total2 << "\n";
  return 0;
}
