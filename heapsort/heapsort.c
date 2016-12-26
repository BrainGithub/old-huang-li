/*
 * Desc		: heap sort. 
 * Author	: zxp
 * Date		: Dec 23 07:18:07 2016
 * */

// 0 1 2 3 4 5 6 7
#include <stdio.h>
#include <stdlib.h>

#define PARENT(i) (i-1)/2;
#define LEFT_CHILD(i) 2*(i) + 1
#define RIGHT_CHILD(i) 2*(i) + 2

static void swap(int * a, int * b)
{
	int temp = *a;
	*a = *b;
	*b = temp;
	return;
}

static void shift_down(int * arr, int start, int end)
{
	int root, left, right, tmp;
	tmp = root = start;
	left = LEFT_CHILD(root);
	right = RIGHT_CHILD(root); 

	if (left > end)
		return;

	if (arr[tmp] < arr[left])
		tmp = left;
	if (right <= end && arr[tmp] < arr[right])
		tmp = right;

	if (tmp == root)
		return;
		
	swap(&arr[root], &arr[tmp]);
	shift_down(arr, tmp, end);
}

static void heapify(int * arr, int count)
{
	int start = PARENT(count -1);

	for (; start >= 0; start--){
		shift_down(arr, start, count - 1);
	}
}

int heapsort(int * arr, int count)
{
	int end;
	int i;

	heapify(arr, count);
	for (end=count-1; end>0; end--){
		swap(&arr[0], &arr[end]);
		shift_down(arr, 0, end-1);
	}
	
	return;
}
