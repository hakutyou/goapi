SUBDIRS = core account moonlight

all: init $(SUBDIRS) web
	docker exec -u root goapi_api_1 kill -USR2 `cat web.pid`

init:
	go mod tidy
	go mod vendor

fmt:
	for file in `find . -path ./vendor -prune -o -name '*.go' -print`; do \
		gofmt -w $${file}; \
	done

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
.PHONY: FORCE init all clean fmt

