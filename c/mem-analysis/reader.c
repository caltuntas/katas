#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <fcntl.h>

typedef struct {
    int id;
    float value;
    char name[32];
} ExampleStruct;

int main(int argc, char *argv[]) {
    if (argc != 3) {
        fprintf(stderr, "Usage: %s <pid> <address>\n", argv[0]);
        return 1;
    }

    int pid = atoi(argv[1]);
    unsigned long address = strtoul(argv[2], NULL, 16);

    // Open /proc/<pid>/mem
    char mem_path[64];
    snprintf(mem_path, sizeof(mem_path), "/proc/%d/mem", pid);
    int fd = open(mem_path, O_RDONLY);
    if (fd == -1) {
        perror("Failed to open memory file");
        return 1;
    }

    // Seek to the address
    if (lseek(fd, address, SEEK_SET) == -1) {
        perror("Failed to seek to address");
        close(fd);
        return 1;
    }

    // Read the struct
    ExampleStruct example;
    if (read(fd, &example, sizeof(ExampleStruct)) != sizeof(ExampleStruct)) {
        perror("Failed to read memory");
        close(fd);
        return 1;
    }

    // Print the struct contents
    printf("Struct at 0x%lx:\n", address);
    printf("  id: %d\n", example.id);
    printf("  value: %.2f\n", example.value);
    printf("  name: %s\n", example.name);

    close(fd);
    return 0;
}
