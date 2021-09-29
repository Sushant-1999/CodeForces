#include <iostream>
#include <bits/stdc++.h>
using namespace std;

int main() {
    long long int up_part, lo_part, num, k, count, temp, i, j, n, flag, ptr;
    up_part = 1; lo_part = 1; num = 0; k = 0; count = 1;
    cin >> n;
    if(up_part == 1) {
        for(j = 0; j <= n; j++) {
        // Printing Spaces

            for(i = 0; i < n-j; i++) {
                cout << "  ";
            }
            // Printing Middle Part
            temp = 0; count = j + 1;
            while(count--) {
                cout << temp << " ";
                temp = temp + 1;
            }
            
            // Printing last part
            flag = j; ptr = j - 1;
            while(flag--) {
                cout << ptr << " ";
                ptr = ptr - 1;
            }
            cout << endl;
        }

    }

    // For Lower Part
    if(lo_part == 1) {
        for(j = 0; j < n; j++) {
            // Printing Spaces
            for(i = 0; i <= j; i++) {
                cout << "  ";
            }

            // Printing Middle part
            count = n - j; temp = 0;
            while(count--) {
                cout << temp << " ";
                temp = temp + 1;
            }

            // Printing lower part
            flag = n - (j + 1); ptr = n - (j + 2);
            while(flag--) {
                cout << ptr << " ";
                ptr = ptr - 1;
            }
            cout << endl;
        }
    }

    return 0;

}