/*
 *heap sort test procedure
 *author: zxp
 *
 */

#include <stdio.h>
#include <stdlib.h>

#define LEN 10000

int main()
{
	int arr[LEN];
	int i;

	for (i=0; i<LEN; i++){
		printf(" %d", arr[i]);
	}
	printf("\n");

	heapsort(arr, LEN);

	for (i=0; i<LEN; i++){
		printf(" %d", arr[i]);
	}
	printf("\n");
	return 0;
}
