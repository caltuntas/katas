AS := nasm
ASFLAGS := -f elf64 -g -F dwarf
CC := gcc
CFLAGS := -no-pie 
SRC := $(wildcard *.asm)
TARGET := $(SRC:%.asm=bin/%)
OBJS = $(SRC:.asm=.o)

all: | bin $(TARGET)

%.o : %.asm
		$(AS) $(ASFLAGS) $<

bin/% : %.o
		$(CC) $(CFLAGS) $< -o $@

bin:
		mkdir bin

clean:
		rm -fr bin/

test:
		$(foreach var,$(.VARIABLES),$(info $(var) = $($(var))))
