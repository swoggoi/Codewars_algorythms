#include <iostream>
using namespace std;
int main(){
    const int n = 10;
    int arr[n]{4,2,67,1,55,3,88,9,0,33};
    for(int i = 0; i < n- 1; ++i){
        int min = i;
        for(int j = i + 1;j< n; ++j){
            if(arr[j]>arr[min]){
                min= j;
                swap(arr[i], arr[min]);
            }
        }
    }
    for(auto& x: arr){
        cout << x << " ";
    }
}