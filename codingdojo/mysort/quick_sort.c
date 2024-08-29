#include <stdio.h>
#include <stdlib.h>

void my_qsort(int * arr, int start, int end) {
	int i, j, pivot;
	
	if (start >= end)
		return;
	i = start;
	j = end;
	pivot = arr[(i+j)/2];
	arr[(i+j)/2] = arr[i];
	while(i<j){
		for (;i<j && arr[j] >= pivot; j--);
		arr[i] = arr[j];
		for (;i<j && arr[i] <= pivot; i++);
		arr[j] = arr[i];
	}
	arr[i] = pivot;

	my_qsort(arr, start, i-1);
	my_qsort(arr, i+1, end);
}

