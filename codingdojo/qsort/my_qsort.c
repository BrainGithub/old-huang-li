#include <stdio.h>
#include <stdlib.h>

void my_qsort(int * arr, int left, int right)
{
	int i,j,pivot;
	int temp;

	if (left >= right)
		return;

	i = left;
	j = right;
	temp = (i+j)/2;
	pivot = arr[temp];
	arr[temp] = arr[i];

	for (;i<j;){
		for(;i<j && arr[j] >= pivot; j--);
		arr[i] = arr[j];
		for(;i<j && arr[i] <= pivot; i++);
		arr[j] = arr[i];
	}
	arr[i] = pivot;

	my_qsort(arr, left, i-1);
	my_qsort(arr, i+1, right);

	return;
}
