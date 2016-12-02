/*
    to get prime factors 
    with TDD method.
*/

#include <stdio.h>
#include <stdlib.h>
#include <math.h>

typedef void (*test_fn_type)(int *);


int get_prime_factors(int val, int * factors)
{
    int i = 0;
    int fact = 2;

    for (; fact <= val; fact++){
        for (; val % fact == 0; val /= fact){
            factors[i++] = fact;
        }
    }

    return i;
}


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

static void prime_factor_test_2_is_2(int * factors, int * result)
{   
    int n = get_prime_factors(2, factors);
    if (1 != n || 0 != memcmp(factors, result, n))
    {   
        printf("verify failed: %s\n", __func__);
        exit(0);
    }
    
    return;
}

static void prime_factor_test_3_is_3(int * factors, int * result)
{   
    int n = get_prime_factors(3, factors);
    if (1 != n || 0 != memcmp(factors, result, n))
    {   
        printf("verify failed: %s\n", __func__);
        exit(0);
    }
    
    return;
}

static void prime_factor_test_4_is_2_2(int * factors, int * result, int len)
{   
    int i;
    int n = get_prime_factors(4, factors);
    if (n != len || 0 != memcmp(factors, result, n))
    {   
        printf("verify failed: %s\n", __func__);
        for (i=0; i<len; i++){
            printf(" %d", factors[i]);
        }
        exit(0);
    }
    
    return;
}

static void prime_factor_test_6_is_2_3(int * factors, int * result, int len)
{   
    int i;
    int n = get_prime_factors(6, factors);
    if (n != len || 0 != memcmp(factors, result, n))
    {   
        printf("verify failed: %s\n", __func__);
        for (i=0; i<len; i++){
            printf(" %d", factors[i]);
        }
        exit(0);
    }
    
    return;
}

static void prime_factor_test_7_is_7(int * factors, int * result, int len)
{

    int i;
    int n = get_prime_factors(7, factors);
    if (n != len || 0 != memcmp(factors, result, n))
    {   
        printf("verify failed: %s\n", __func__);
        for (i=0; i<len; i++){
            printf(" %d", factors[i]);
        }
        exit(0);
    }
    return;
}

static void prime_factor_test_924_is_2_2_3_7_11(int * factors, int * result, int len)
{
    int i;
    int n = get_prime_factors(924, factors);
    if (n != len || 0 != memcmp(factors, result, n))
    {   
        printf("verify failed: %s\n", __func__);
        for (i=0; i<len; i++){
            printf(" %d", factors[i]);
        }
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

    res_factors[i] = 2;
    prime_factor_test_2_is_2(factors, res_factors);

    i = 0;
    res_factors[i] = 3;
    prime_factor_test_3_is_3(factors, res_factors);
 
    i=0;
    res_factors[i++] =2;
    res_factors[i++] =2;
    prime_factor_test_4_is_2_2(factors, res_factors, 2);
    

    i = 0;
    res_factors[i++] =2;
    res_factors[i++] =3;
    prime_factor_test_6_is_2_3(factors, res_factors, 2);

    i=0;
    res_factors[i++] = 7;
    prime_factor_test_7_is_7(factors, res_factors, 1);

    i=0;
    res_factors[i++] = 2;
    res_factors[i++] = 2;
    res_factors[i++] = 3;
    res_factors[i++] = 7;
    res_factors[i++] = 11;
    prime_factor_test_924_is_2_2_3_7_11(factors, res_factors, 5);
    
    return 0;

}
