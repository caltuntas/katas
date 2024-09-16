OBJ = fib.o main.o
CFLAGS = -Wall -g
TARGET = macho64
NASM_FLAGS += -Dfib=_fib

testt: t
	./t

clean:
	rm -f t *.o

t: $(OBJ)
	$(CC) $(LDFLAGS) -o $@ $* $(OBJ) $(LIBS)

fib.o: fib.asm

main.o: main.c

%.o: %.asm
	nasm $(NASM_FLAGS) -f$(TARGET) $<

%.o: %.c
	$(CC) $(CFLAGS) -c $<
