package hackerrank

import (
	"fmt"
	"sort"
)

/*

	HackerRank only has a few languages available to resolve this problem.
	I went with C++, and the solution is:

	#include <bits/stdc++.h>
	using namespace std;

	void findZigZagSequence(vector < int > a, int n){
	    sort(a.begin(), a.end());

	    int mid = (n - 1)/2;
	    swap(a[mid], a[n-1]);

	    int st = mid + 1;
	    int ed = n - 2;
	    while(st <= ed){
	        swap(a[st], a[ed]);
	        st = st + 1;
	        ed = ed - 1;
	    }

	    for(int i = 0; i < n; i++){
	        if(i > 0) cout << " ";
	        cout << a[i];
	    }
	    cout << endl;
	}

	And the Go translation:
*/

func findZigZagSequence(arr []int32, n int32) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	last := n - 1
	mid := (n + 1) / 2

	arr[mid], arr[last] = arr[last], arr[mid]

	st := mid + 1
	end := last - 1

	for st <= end {
		arr[st], arr[end] = arr[end], arr[st]
		st++
		end--
	}

	for _, v := range arr {
		fmt.Printf("%d ", v)
	}
}
