#include<iostream>
using namespace std;
int main(){
    const int n = 11;
    int arr[n]{4,2,67,1,55,3,88,9,0,33,333};
    for(int gap = n/2; gap>0; gap/= 2){
        for(int i = gap; i  < n; i++){
            int temp = arr[i];
            int j = i;
            while(j>=gap && arr[j - gap] > temp){
                arr[j] = arr[j-gap];
                j-=gap;
            }
            arr[j] = temp;
        }
    }
    for(auto& x: arr){
        cout << x << " ";
    }

    return 0 ;
}