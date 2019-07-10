notminisign: *.go cmd/*.go Makefile
	docker run --rm -it -v $$PWD:$$PWD -w $$PWD -u $(shell id -u):$(shell id -g) brimstone/golang go build
