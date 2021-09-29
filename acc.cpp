 #include <iostream>
 #include <bits/stdc++.h>
 using namespace std;

string reverseStr(string str)
{
    int n = str.length();
 
    // Swap character starting from two
    // corners
    for (int i = 0; i < n / 2; i++)
        swap(str[i], str[n - i - 1]);

    return str;
}
 

 int main() {
     string str, input;
     int len, i, j, temp, flag = 0;

     cin >> str;

    len = str.size();
    //str.copy(input,len);  
    for(int i=0; i < len; i++) 
    { 
        if(! (str[i]>='A'&& str[i]<='Z'|| str[i]>='a'&& str[i]<='z')) 
        { 
            flag=1; 
            break;   //if this condition is true then no need to check any other characters, so I break from for loop here 
        } 
    } 
 
    if(flag == 0) {
        // now reverse a string
           input =  reverseStr(str);

    }
    cout << input << endl;
    cout << str << endl;
    if(str == input) {
        cout << "you inputted a strong string.";
    }
    else if(str != input) {
        cout << "you inputted a weak string.";
    }
    return 0;
    
 }