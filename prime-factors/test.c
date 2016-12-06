/*
    to get prime factors 
    with TDD method.
*/

#include <stdio.h>
#include <stdlib.h>
#include <math.h>

typedef void (*test_fn_type)(int *);

static void prime_factor_test(int data, int * factors, int * result, int count)
{
    int i;
    int n = get_prime_factors(data, factors);
    if (n != count || memcmp(factors, result, n) != 0)
    {
        printf("verify failed: %s\n", __func__);
        printf("data=%d,n=%d\n", data, n);
        for (i=0; i<n; i++)
            printf(" %d", factors[i]);
        printf("\n");
        exit(0);
    }
}

int main()
{
    int factors[1000] = {0};
    int res_factors[1000] = {0};
    int i=0;
    int data=0;

    i=0;
    prime_factor_test(-1, factors, res_factors, i);
   
/*    i = 0;
    res_factors[i++] = 2;
    prime_factor_test_2_is_2(factors, res_factors, i);

    i = 0;
    res_factors[i++] = 3;
    prime_factor_test_3_is_3(factors, res_factors, i);


    i = 0;
    res_factors[i++] = 2;
    res_factors[i++] = 2;
    prime_factor_test_4_is_2_2(factors, res_factors, i);
	
    data = 5;    
    
    i = 0;
    res_factors[i++] = 5;
    prime_factor_test(data++, factors, res_factors, i);
   
   
    i = 0;
    res_factors[i++] = 2;
    res_factors[i++] = 3;
    prime_factor_test(data++, factors, res_factors, i);
   
    i = 0;
    res_factors[i++] = 7;
    prime_factor_test(data++, factors, res_factors, i);
   
    i = 0;
    res_factors[i++] = 2;
    res_factors[i++] = 2;
    res_factors[i++] = 2;
    prime_factor_test(data++, factors, res_factors, i);
   
    
    i = 0;
    res_factors[i++] = 3;
    res_factors[i++] = 3;
    prime_factor_test(data++, factors, res_factors, i);
   
    i = 0;
    res_factors[i++] = 3;
    res_factors[i++] = 7;
    prime_factor_test(21, factors, res_factors, i);
*/   
    
    printf("passed successfully!\n");    
    return 0;

}
