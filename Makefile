SUBDIRS = core account moonlight

all: init $(SUBDIRS) web

init:
	go mod tidy
	go mod vendor

clean:
	for d in $(SUBDIRS); do \
		make clean -C $$d; \
	done
	make clean -C web

$(SUBDIRS): $(patsubst %, .%.yaml, $(SUBDIRS)) FORCE
	make -C $@

web: .zap.yaml .config.yaml FORCE
	make -C $@

%.yaml:
	@echo "$@ 不存在"; exit 1

FORCE:
.PHONY: FORCE init all clean

