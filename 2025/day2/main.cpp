#include <bits/stdc++.h>
using namespace std;

typedef long long ll;
typedef pair<ll, ll> pll;
typedef vector<pll> vl;

#define fastio ios::sync_with_stdio(false); cin.tie(nullptr);
#define endl '\n'

bool checkForCondition(ll nums) {
    string data = to_string(nums);
    int n = data.size();
    if(n%2) return false;

    return data.substr(0, n/2) == data.substr(n/2);
}

vl getInputs(string filename) {
    vl inputs;
    ifstream file(filename);

    if(!file){
        cout << "File not found\n";
        return inputs;
    }

    string input;
    getline(file, input);

    stringstream ss(input);
    string range;

    while(getline(ss, range, ',')) {
        int dash = range.find('-');

        ll L = stoll(range.substr(0, dash));
        ll R = stoll(range.substr(dash+1));

        inputs.push_back({L, R});
    }
    return inputs;
}

ll partOne(){
    vl data = getInputs("inputs.txt");
    ll count = 0;
    for(auto [L, R] : data) {
        for(ll x = L; x <= R; x++){
            if(checkForCondition(x))
                count += x;
        }
    }

    return count;
}

bool checkForCondition2(ll num){
    string data = to_string(num);
    int n = data.size();

    for(int k = 1; k <= n/2; k++){
        if(n%k) continue;

        bool ok = true;
        for(int i = k; i < n; i++){
            if(data[i] != data[i % k]){
                ok = false;
                break;
            }
        }
        if(ok)return true;
    }    
    return false;
}

ll partTwo() {
    vl data = getInputs("inputs.txt");
    ll count = 0;

    for(auto [L, R] : data) {
        for(ll x = L; x <= R; x++){
            if(checkForCondition2(x))
                count += x;
        }
    }
    return count;
}

int main() {
    fastio;
    cout << "Part 1: " << partOne() << "\n";
    cout << "Part 2: " << partTwo() << "\n";
    return 0;
}