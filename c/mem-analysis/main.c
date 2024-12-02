#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

typedef struct {
    int id;
    float value;
    char name[32];
} ExampleStruct;

int main() {
    // Allocate struct on the heap
    ExampleStruct *example = (ExampleStruct *)malloc(sizeof(ExampleStruct));
    if (example == NULL) {
        perror("malloc failed");
        return 1;
    }

    // Populate the struct
    example->id = 42;
    example->value = 3.14f;
    strncpy(example->name, "HeapStruct", sizeof(example->name) - 1);

    printf("Struct is allocated at: %p\n", (void *)example);
    printf("PID: %d\n", getpid());
    printf("Press Enter to exit...");
    getchar(); // Keeps the program running

    // Free the memory
    free(example);
    return 0;
}
