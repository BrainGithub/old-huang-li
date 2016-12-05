/*
    to get prime factors 
    with TDD method.
*/

#include <stdio.h>
#include <stdlib.h>
#include <math.h>

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



