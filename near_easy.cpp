#include <iostream>
#include <bits/stdc++.h>
using namespace std;


bool KBeautiful(int x, int k) {
    // check count <= k
    long long int  rem;
    bool ans;
    set<long long int> Set;
    while(x > 0) {
        rem = x % 10;
        Set.insert(rem);
        x = x / 10;
    }
    if(Set.size() <= k) {
        ans = true;
    } 
    else if(Set.size() > k) {
        ans = false;

    }
    return ans;

}

int main() {
    long long int T, n, i, x, k, flag, res;
    cin >> T;

    while(T--) {
        cin >> n >> k;
        flag = 1;
        for(x = n; flag != 0; x++) {
            res = KBeautiful(x, k);
            if(res) {
                flag = 0;
                break;
            }
            else if(res) {
                flag = 1;
            }
        }
        cout << x << endl;
    }
    return 0;

}