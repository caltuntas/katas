#include <stdio.h>

int main() {
    size_t i = 0;
    size_t a = i - 1;

    printf("a = %zu\n", a); // Output: a = 18446744073709551615
    return 0;
}
