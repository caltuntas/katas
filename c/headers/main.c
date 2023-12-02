#include <stdio.h>
#include "myfunctions.h"

int main() {
    int number = 5;
    int result_square = square(number);

    printf("Square of %d is: %d\n", number, result_square);

    int num1 = 3;
    int num2 = 4;
    int result_multiply = multiply(num1, num2);

    printf("Multiplication of %d and %d is: %d\n", num1, num2, result_multiply);

    return 0;
}
