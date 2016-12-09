#include <stdio.h>
#include <stdlib.h>
#include <string.h>

static int cmp(const void *p1, const void *p2, void * priv)
{
/* The actual arguments to this function are "pointers to
*pointers to char", but strcmp(3) arguments are "pointers
*to char", hence the following cast plus dereference */
	return *(int *)p1 < *(int *)p2 ? -1:(*(int *)p1 == *(int *)p2 ? 0:1);
}

int main(int argc, char *argv[])
{
	int j;
	int arr[] = {1, -1, 5, 3};
		            
	if (argc < 2) {
		fprintf(stderr, "Usage: %s <string>...\n", argv[0]);
		//exit(EXIT_FAILURE);
	}


	qsort_r(arr, 4, sizeof(int), cmp, NULL);

	for (j = 0; j < 4; j++)
		printf(" %d", arr[j]);
	printf("\n");
	exit(EXIT_SUCCESS);
}

