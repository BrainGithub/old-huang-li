/*
    to get prime factors 
    with TDD method.
*/

#include <stdio.h>
#include <stdlib.h>
#include <math.h>

//typedef void (*test_fn_type)(int *);

static void prime_factor_test(int data, int * factors, int * result, int count)
{
    int i;
    int n = get_prime_factors(data, factors);
        
    printf("data=%d, there are n=%d prime factors:\n", data, n);
    for (i=0; i<n; i++)
         printf(" %d", factors[i]);
    printf("\n");

    if (n != count || memcmp(factors, result, n) != 0)
    {
        printf("verify failed: %s\n", __func__);
        exit(0);
    }
}

int main()
{
    int factors[1000] = {0};
    int res_factors[1000] = {0};
    int i=0;
    int data=-1;

    i=0;
    prime_factor_test(-1, factors, res_factors, i);
    prime_factor_test(0, factors, res_factors, i);
   
    i = 0;
    res_factors[i++] = 2;
    prime_factor_test(2, factors, res_factors, i);

    i = 0;
    res_factors[i++] = 3;
    prime_factor_test(3, factors, res_factors, i);

    i = 0;
    res_factors[i++] = 2;
    res_factors[i++] = 2;
    prime_factor_test(4, factors, res_factors, i);
	
    data = 5;    
    
    i = 0;
    res_factors[i++] = 5;
    prime_factor_test(5, factors, res_factors, i);
   
   
    i = 0;
    res_factors[i++] = 2;
    res_factors[i++] = 3;
    prime_factor_test(6, factors, res_factors, i);
   
    i = 0;
    res_factors[i++] = 7;
    prime_factor_test(7, factors, res_factors, i);
   
    i = 0;
    res_factors[i++] = 2;
    res_factors[i++] = 2;
    res_factors[i++] = 2;
    prime_factor_test(8, factors, res_factors, i);
   
    i = 0;
    res_factors[i++] = 3;
    res_factors[i++] = 3;
    prime_factor_test(9, factors, res_factors, i);
   
    i = 0;
    res_factors[i++] = 2;
    res_factors[i++] = 3;
    res_factors[i++] = 5;
    res_factors[i++] = 7;
    prime_factor_test(210, factors, res_factors, i);
    
    i = 0;
    res_factors[i++] = 2;
    res_factors[i++] = 3;
    prime_factor_test(29280, factors, res_factors, i);
    
    printf("---- passed successfully!\n");    
    return 0;
}
