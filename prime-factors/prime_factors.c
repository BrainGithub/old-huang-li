/*
    to get prime factors 
    with TDD method.
*/

#include <stdio.h>
#include <stdlib.h>
#include <math.h>
/*
int get_prime_factors(int data, int * factors)
{
    int n;
    int candidate = 2;

    for (n=0; candidate <= data; candidate++){
        for (; data % candidate == 0; data /= candidate){
            factors[n++] = candidate;
        }
    }

    return n;
}

*/
int get_prime_factors_recursive(int data, int * factors, int count, int candidate);

int get_prime_factors(int data, int * factors)
{
        return get_prime_factors_recursive(data, factors, 0, 2);
}

int get_prime_factors_recursive(int data, int * factors, int count, int candidate)
{
    if (data < 2)
        return count;

    if (data % candidate == 0){
        factors[count++] = candidate;
        data /= candidate;
    }
    else
    {
        candidate++;
    }

    return get_prime_factors_recursive(data, factors, count, candidate);
}
