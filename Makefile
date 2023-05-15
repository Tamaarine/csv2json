exec := csv2json
sources := $(shell find -name "*.go")

all: setup bin/$(exec)

setup:
	mkdir -p bin

bin/$(exec):
	go build -o bin/$(exec) $(sources)