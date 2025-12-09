#include<iostream>
using namespace std;

int main(){
    const int n = 9;
    int arr[n]{1,2,3,4,5,77,111,222, 223};
    int target;
    cout << "Введите число которое хотите найти: " << endl;
    cin >> target;
    int right = n - 1;
    int left = 0;
    bool found = false;
    while(left<=right){
        int mid = (left+right) / 2;
        if(arr[mid] == target){
            cout << "Число найдено на позиции:" << mid << endl;
            found = true;
            break;
        }
            else if(arr[mid] >target){
                right = mid-1;
        }  else if(arr[mid  ] < target){
            left = mid + 1;
        }
    }
    if (!found) {
        cout << "Число не найдено." << endl;
    }
    return 0;
}