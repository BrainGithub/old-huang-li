#include <stdio.h>
#include <stdlib.h>
#include <string.h>

static int cmp(const void *p1, const void *p2, void * priv)
{
	//return: < -1, == 0, > 1
	return *(int *)p1 < *(int *)p2 ? -1:(*(int *)p1 == *(int *)p2 ? 0:1);
}

int main(int argc, char *argv[])
{
	#define MAX 100
	int j;
	int arr[MAX];
	int arr2[MAX];
	int arr3[MAX];
	int temp=0;            
	if (argc < 2) {
		fprintf(stderr, "Usage: %s <string>...\n", argv[0]);
		//exit(EXIT_FAILURE);
	}

	srand(time(NULL));
	for (j=0; j<MAX; j++){
		arr[j] = arr2[j] = arr3[j] = rand()%10000;
	}
	printf("\n-------------------------------%d\n",__LINE__);
	for (j = 0; j < MAX; j++)
		printf(" %d", arr3[j]);

	printf("\n-------------------------------%d\n",__LINE__);
	qsort_r(arr, MAX, sizeof(int), cmp, NULL);
	my_qsort(arr2, 0, MAX-1); 
	for (j = 0; j < MAX; j++)
		printf(" %d", arr[j]);
	printf("\n-------------------------------%d\n",__LINE__);
	for (j = 0; j < MAX; j++)
		printf(" %d", arr2[j]);
	printf("\n");

	if (memcmp(arr, arr2, sizeof(int) *MAX))
		printf("---failed\n");

	return 0;	
}

