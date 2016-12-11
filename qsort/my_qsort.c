#include <stdio.h>
#include <stdlib.h>

int deep = 0;

int my_qsort(int * arr, int const left, int const right)
{
	int i,j,p,pivot;
	
	if (left >= right)
		return;

	i=left;
	j=right;
	p = (i+j)/2;
	pivot= arr[p];
	arr[p] = arr[i];

	while(i < j){
		for (;i<j && arr[j] >= pivot; j--);
		arr[i] = arr[j];
		for (;i<j && arr[i] <= pivot; i++);
		arr[j] = arr[i];
	}
	arr[i] = pivot;
	
	my_qsort(arr, left, i-1);
	my_qsort(arr, i+1, right);

	return 0;
}
