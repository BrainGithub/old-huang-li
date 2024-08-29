/*
    to get prime factors 
    with TDD method.
*/

#include <stdio.h>
#include <stdlib.h>
#include <math.h>

int get_prime_factors(int data, int * factors)
{
    return recursive_prime_factors(data, factors, 0, 2);
}

int recursive_prime_factors(int data, int * factors, int n, int candidate)
{
    if (data < 2)
        return n;
    for (; data % candidate == 0; data /= candidate){
        factors[n++] = candidate;
    }
    return recursive_prime_factors(data, factors, n, ++candidate);
}





