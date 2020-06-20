ELF = moonlight.elf

all: $(ELF)

%.elf: FORCE
	go build -o $@

clean:
	-rm -f $(ELF)

FORCE:
.PHONY: FORCE all clean

