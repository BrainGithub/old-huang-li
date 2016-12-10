#include <stdio.h>
#include <stdlib.h>

int my_qsort(int * arr, int const left, int const right)
{
	int n;
	int i=left,j=right;
	int pivot = arr[i];
	
	if (left >= right)
		return;

	for(n=left; n<=right; n++)
		printf(" %d",arr[n]);
	printf("--before sort\n");
	
	while(i < j){
		for (;i<j && arr[j] > pivot; j--);
		if (i<j) arr[i++] = arr[j];
		for (;i<j && arr[i] < pivot; i++);
		if (i<j) arr[j] = arr[i];
	}
	arr[i] = pivot;
	printf("i:%d,j:%d\n",i,j);
	for(n=left; n<=right; n++)
		printf(" %d",arr[n]);
	printf("--after sort\n");

	sleep(1);
	my_qsort(arr, left, i-1);
	my_qsort(arr, i+1, right);
	return 0;
}
