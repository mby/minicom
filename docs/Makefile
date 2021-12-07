
all: services.svg;

.PHONY: all services.svg
services.svg: services.dot
	dot -T svg services.dot > services.svg

clean:
	rm services.svg
