    #include<iostream>
    using namespace std;

    int main(){
        const int n = 11;
        int arr[n]{4,2,67,1,55,3,88,9,0,33,333};
        int l =0;
        int r = n- 1;
        int pivot = arr[(l + r) / 2];
        while(true){
            while(arr[l] < pivot){ 
            l++;
        }   
            while(arr[r] > pivot){ 
            r--;
            }
            if( l>=r) {
                break;
            }
            swap(arr[l], arr[r]);
            l++;
            r--;
        }
        for(auto& x: arr)
        cout << x << " ";
    }
