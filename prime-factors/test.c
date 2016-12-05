/*
    to get prime factors 
    with TDD method.
*/

#include <stdio.h>
#include <stdlib.h>
#include <math.h>

typedef void (*test_fn_type)(int *);


static void prime_factor_test_1_is_null(int * factors)
{   
    int n = get_prime_factors(1, factors);
    if (0 !=  n)
    {   
        printf("verify failed: %s\n", __func__);
        exit(0);
    }
    
    return;
}

int main()
{
    int factors[1000] = {0};
    int res_factors[1000] = {0};
    int i=0;

    prime_factor_test_1_is_null(factors);


	printf("passed successfully!\n");    
    return 0;

}
