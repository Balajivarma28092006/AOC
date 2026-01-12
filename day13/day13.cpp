#include <iostream>
#include <fstream>
#include <vector>
#include <string>
#include <regex>


using namespace std;

vector<long long> extract_nums(const string& s){
    static regex num_re("-?\\d+");
    vector<long long> nums;
    for(auto it = sregex_iterator(s.begin(), s.end(), num_re);
            it != sregex_iterator(); ++it){
                nums.push_back(stoll(it->str()));
            }
            return nums;
}

long long solve(const vector<string>& lines, bool part2){
    long long totalcost = 0;
    for(int i = 0; i < (int)lines.size(); i+= 3){
        string block = lines[i] + lines[i+1] + lines[i+2];
        auto v = extract_nums(block);

        long long Ax = v[0], Ay = v[1];
        long long Bx = v[2], By = v[3];
        long long Px = v[4], Py = v[5];

         if (part2) {
            Px += 10000000000000LL;
            Py += 10000000000000LL;
        }

        long long D = Ax * By - Bx * Ay;
        if ( D == 0 ) continue;

        long long aNum = Px * By - Py * Bx;
        long long bNum = Ax * Py - Px * Ay;

        if ( aNum % D != 0 || bNum % D != 0 ) continue;

        long long a = aNum / D;
        long long b = bNum / D;

        if( a < 0 || b < 0) continue;

        totalcost += 3 * a + b;
    }
    return totalcost;
}

int main() {
    ifstream fin("day13_inputs.txt");
    vector<string> lines;
    string line;

    while(getline(fin, line)){
        if (!line.empty()){
            lines.push_back(line);
        }
    }

    cout << "Part 1: " << solve(lines, false) << '\n';
    cout << "Part 2: " << solve(lines, true) << '\n';
    return 0;
}