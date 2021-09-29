#include <iostream>
#include <bits/stdc++.h>
using namespace std;

int main() {
    long long int a, b, n, res = 0, i, j;
    cin >> a >> b >> n;
    res = a;
    for(i = 0; i < n; i++) {
        for(j = 0; j < 9; j++) {
            long long temp = res * 10 + j;
            if(temp % b == 0) {
                res = temp;
                break;
            }
            
        }
    }
    if(res == a) {
        cout << "-1" << endl;
    }
    else {
        cout << res << endl;
    }
    return 0;
    
}