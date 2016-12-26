#include <stdio.h>
#include <stdlib.h>

static void swap(int * a, int * b){
	int tmp = *a;
	*a = *b;
	*b = tmp;
}

void biblesort(int * arr, int count){
	int i,j,flag;
	for(i=count-1; i>0; i--){
		flag = 0;
		for (j=0;j<i;j++){
			if (arr[j] > arr[j+1]){
				swap(&arr[j], &arr[j+1]);
				flag = 1;
			}
		}
		if (0 == flag)
			break;
	}
}

