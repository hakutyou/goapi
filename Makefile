SUBDIRS = core account consumer moonlight web

all: init $(SUBDIRS)

init:
	go mod tidy
	go mod vendor


$(SUBDIRS): FORCE
	make -C $@

FORCE:
.PHONY: FORCE init all

