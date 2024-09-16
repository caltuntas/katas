#include <stdio.h>

extern unsigned int fib (unsigned int);

int main(void)
{
    printf("%d\n", fib (40));
    return 0;
}
