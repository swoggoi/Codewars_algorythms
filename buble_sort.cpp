#include <iostream>
using namespace std;

int main(){
    const int n = 10;
    int arr[n];
    for(int i = 0; i < n; i++){
        cout << "Введите номер [" << i << "]: ";
        cin >> arr[i];
    }
    for(int i = 0; i < n; ++i){
        for(int j = 0; j < n - i -1; j++){
            if(arr[j] > arr[j + 1]){
            int zxc = arr[j];
            arr[j] = arr[j+1];
            arr[j + 1] = zxc;
        }
    }
    }
    for(auto& x: arr){
        cout << x << " ";
    }
    return 0;
}