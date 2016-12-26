#include <stdio.h>
#include <stdlib.h>
static void swap(int * a, int * b);
static void shift_down(int *arr, int start, int end);
static void heapify(int *arr, int count);

void heapsort(int *arr, int count){
	int end;
	heapify(arr, count);
	for (end = count -1; end>0; end--){
		swap(&arr[0], &arr[end]);
		shift_down(arr, 0, end-1);
	}
}

static void swap(int *a, int * b)
{
	int temp = *a;
	*a = *b;
	*b = temp;
}

static void heapify(int * arr, int count){
	int start;
	for (start = (count-1)/2; start>=0; start--){
		shift_down(arr, start, count-1);
	}
}

static void shift_down(int * arr, int start, int end){
	int tmp,left,right, root;
	tmp = root = start;
	left = 2*root +1;
	right = 2*root +2;
	if (left <= end && arr[left] > arr[tmp])
		tmp = left;
	if (right <= end && arr[right] > arr[tmp])
		tmp = right;
	if (tmp == root)
		return;

	swap(&arr[root], &arr[tmp]);
	shift_down(arr, tmp, end);
}

