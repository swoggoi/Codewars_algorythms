#include<cmath>
bool narcissistic( int value ){
  int count = 0;
  int temp = value;
  while (temp > 0) {
        temp /= 10;
        count++;
    }
    int sum = 0;
    temp = value;
    while(temp > 0){
        int digit = temp % 10;
        sum += pow(digit, count);
        temp /= 10;
    }
    return sum==value;
}