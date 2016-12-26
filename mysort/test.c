#include <stdio.h>
#include <time.h>
#include <stdlib.h>
#include <string.h>

static int cmp(const void *p1, const void *p2, void * priv)
{
	//return: < -1, == 0, > 1
	return *(int *)p1 < *(int *)p2 ? -1:(*(int *)p1 == *(int *)p2 ? 0:1);
}

int main(int argc, char *argv[])
{
	#define MAX 10000
	int j;
	int arr[MAX];
	int arr1[MAX];
	int arr2[MAX];
	int arr3[MAX];
	int temp=0;            
	if (argc < 2) {
		fprintf(stderr, "Usage: %s <string>...\n", argv[0]);
		//exit(EXIT_FAILURE);
	}

	for (j=0; j<MAX; j++){
		arr[j] = arr2[j] = arr3[j] = 99;//rand()%10000;
	}

	qsort_r(arr, MAX, sizeof(int), cmp, NULL);
	my_qsort(arr2, 0, MAX-1); 
	
	if (memcmp(arr, arr2, sizeof(int) *MAX))
		printf("---failed:%d\n", __LINE__);

    //--------------
	srand(time(NULL));
	for (j=0; j<MAX; j++){
		arr[j] = arr1[j] = arr2[j] = arr3[j] = rand()%100000000;
	}



	qsort_r(arr, MAX, sizeof(int), cmp, NULL);
	
	printf("\n\n: -- compare time spend:\n");
	struct timeval start, end;
	gettimeofday( &start, NULL );

	my_qsort(arr1, 0, MAX-1); 
	if (memcmp(arr, arr1, sizeof(int) *MAX))
		printf("---failed:%d\n", __LINE__);
	gettimeofday( &end, NULL );
	int timeuse = 1000000 * ( end.tv_sec - start.tv_sec ) 
			+ end.tv_usec -start.tv_usec;
	double timespan = (double)timeuse/1000;
	printf("-- quick qsort verified successed! time spend:%fms\n", timespan);



	gettimeofday( &start, NULL );

	heapsort(arr2, MAX);
	if (memcmp(arr, arr2, sizeof(int) *MAX))
		printf("---failed:%d\n", __LINE__);

	gettimeofday( &end, NULL );
	timeuse = 1000000 * ( end.tv_sec - start.tv_sec ) 
			+ end.tv_usec -start.tv_usec;
	timespan = (double)timeuse/1000;
	printf("-- heapsort verified successed! time spend:%fms\n", timespan);
//	printf("\n\n: -- compare time spend:\n");
//	clock_t before = clock();

	
	
	gettimeofday( &start, NULL );

	biblesort(arr3, MAX);
	if (memcmp(arr, arr3, sizeof(int) *MAX))
		printf("---failed:%d\n", __LINE__);

	gettimeofday( &end, NULL );
	timeuse = 1000000 * ( end.tv_sec - start.tv_sec ) 
			+ end.tv_usec -start.tv_usec;
	timespan = (double)timeuse/1000;
	printf("-- bible sort verified successed! time spend:%fms\n", timespan);


	

	return 0;	
}

