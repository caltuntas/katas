TARGET = lab
FILES = lab.asm

OBJS = $(FILES:.asm=.o)
NASM = nasm
ASM_FLAGS = -f elf64
LD = ld

all : $(TARGET)

lab: $(OBJS)
		$(LD) -o $(TARGET) $(OBJS)

%.o : %.asm
		$(NASM) $(ASM_FLAGS) -o $@ $<

test:
		$(foreach var,$(.VARIABLES),$(info $(var) = $($(var))))
