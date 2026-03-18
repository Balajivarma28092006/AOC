#include <bits/stdc++.h>
using namespace std;

static const int BLINK_PART1 = 25;
static const int BLINK_PART2 = 75;

void add(unordered_map<long long, long long>& freq,
         long long key, long long value) {
    freq[key] += value;
}

unordered_map<long long, long long>
blink(const unordered_map<long long, long long>& freq) {

    unordered_map<long long, long long> next;

    for (auto& [stone, count] : freq) {

        if (stone == 0) {
            add(next, 1, count);
            continue;
        }

        string s = to_string(stone);

        if (s.size() % 2 == 0) {
            int mid = s.size() / 2;
            long long left = stoll(s.substr(0, mid));
            long long right = stoll(s.substr(mid));

            add(next, left, count);
            add(next, right, count);
        } else {
            add(next, stone * 2024LL, count);
        }
    }
    return next;
}

long long solve(unordered_map<long long, long long> freq, int blinks) {
    for (int i = 0; i < blinks; i++) {
        freq = blink(freq);
    }

    long long ans = 0;
    for (auto& [_, count] : freq) {
        ans += count;
    }
    return ans;
}

int main() {
    ios::sync_with_stdio(false);

    ifstream file("day11.txt");
    if (!file.is_open()) {
        cerr << "Error opening day11.txt\n";
        return 1;
    }

    unordered_map<long long, long long> freq;
    long long x;
    while (file >> x) {
        freq[x]++;
    }

    int part;
    cout << "Enter which part (1 or 2): ";
    cin >> part;

    if (part == 1) {
        cout << "Answer: " << solve(freq, BLINK_PART1) << "\n";
    } else {
        cout << "Answer: " << solve(freq, BLINK_PART2) << "\n";
    }

    return 0;
}
